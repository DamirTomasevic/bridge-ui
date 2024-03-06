package indexer

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"math/big"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cyberhorsey/errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/taiko-mono/packages/relayer"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/bridge"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/signalservice"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/taikol1"
	"github.com/taikoxyz/taiko-mono/packages/relayer/pkg/queue"
	"github.com/taikoxyz/taiko-mono/packages/relayer/pkg/repo"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// WatchMode is a type that determines how the indexer will operate.
type WatchMode string

var (
	// Filter will filter past blocks, but when catches up to latest block,
	// will stop.
	Filter WatchMode = "filter"
	// Subscribe ignores all past blocks, only subscibes to new events from latest block.
	Subscribe WatchMode = "subscribe"
	// FilterAndSubscribe filters up til latest block, then subscribes to new events. This is the
	// default mode.
	FilterAndSubscribe WatchMode = "filter-and-subscribe"
	// CrawlPastBlocks filters through the past N blocks on a loop, when it reaches `latestBlock - N`,
	// it will recursively start the loop again, filtering for missed events, or ones the
	// processor failed to process.
	CrawlPastBlocks WatchMode = "crawl-past-blocks"
	WatchModes                = []WatchMode{Filter, Subscribe, FilterAndSubscribe, CrawlPastBlocks}
)

// SyncMode is a type which determines how the indexer will start indexing.
type SyncMode string

var (
	// Sync grabs the latest processed block in the DB and starts from there.
	Sync SyncMode = "sync"
	// Resync starts from genesis, ignoring the database.
	Resync SyncMode = "resync"
	Modes           = []SyncMode{Sync, Resync}
)

// ethClient is a local interface that lets us narrow the large ethclient.Client type
// from go-ethereum down to a mockable interface for testing.
type ethClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// DB is a local interface that lets us narrow down a database type for testing.
type DB interface {
	DB() (*sql.DB, error)
	GormDB() *gorm.DB
}

// Indexer is the main struct of this package, containing all dependencies necessary for indexing
// relayer-related chain data. All database repositories, contract implementations,
// and configurations will be injected here.
// An indexer should be configured and deployed for all possible combinations of bridging.
// IE: an indexer for L1-L2, and another for L2-L1. an L1-L2 indexer will have the L1 configurations
// as its source, and vice versa for the L2-L1 indexer. They will add messages to a queue
// specifically for a processor of the same configuration.
type Indexer struct {
	eventRepo    relayer.EventRepository
	blockRepo    relayer.BlockRepository
	srcEthClient ethClient

	processingBlockHeight uint64

	bridge     relayer.Bridge
	destBridge relayer.Bridge

	signalService relayer.SignalService

	blockBatchSize      uint64
	numGoroutines       int
	subscriptionBackoff time.Duration

	taikol1 *taikol1.TaikoL1

	queue queue.Queue

	srcChainId  *big.Int
	destChainId *big.Int

	watchMode WatchMode
	syncMode  SyncMode

	ethClientTimeout time.Duration

	wg *sync.WaitGroup

	numLatestBlocksToIgnoreWhenCrawling uint64

	targetBlockNumber *uint64

	ctx context.Context

	mu *sync.Mutex

	eventName string
}

// InitFromCli inits a new Indexer from command line or environment variables.
func (i *Indexer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, i, cfg)
}

// InitFromConfig inits a new Indexer from a provided Config struct
func InitFromConfig(ctx context.Context, i *Indexer, cfg *Config) (err error) {
	db, err := cfg.OpenDBFunc()
	if err != nil {
		return err
	}

	eventRepository, err := repo.NewEventRepository(db)
	if err != nil {
		return err
	}

	blockRepository, err := repo.NewBlockRepository(db)
	if err != nil {
		return err
	}

	srcEthClient, err := ethclient.Dial(cfg.SrcRPCUrl)
	if err != nil {
		return err
	}

	destEthClient, err := ethclient.Dial(cfg.DestRPCUrl)
	if err != nil {
		return err
	}

	q, err := cfg.OpenQueueFunc()
	if err != nil {
		return err
	}

	srcBridge, err := bridge.NewBridge(cfg.SrcBridgeAddress, srcEthClient)
	if err != nil {
		return errors.Wrap(err, "bridge.NewBridge")
	}

	destBridge, err := bridge.NewBridge(cfg.DestBridgeAddress, destEthClient)
	if err != nil {
		return errors.Wrap(err, "bridge.NewBridge")
	}

	// taikoL1 will only be set when initializing a L1 - L2 indexer
	var taikoL1 *taikol1.TaikoL1
	if cfg.SrcTaikoAddress != ZeroAddress {
		taikoL1, err = taikol1.NewTaikoL1(cfg.SrcTaikoAddress, srcEthClient)
		if err != nil {
			return errors.Wrap(err, "taikol1.NewTaikoL1")
		}
	}

	var signalService relayer.SignalService
	if cfg.SrcSignalServiceAddress != ZeroAddress {
		signalService, err = signalservice.NewSignalService(cfg.SrcSignalServiceAddress, srcEthClient)
		if err != nil {
			return errors.Wrap(err, "signalservice.NewSignalService")
		}
	}

	srcChainID, err := srcEthClient.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "srcEthClient.ChainID")
	}

	destChainID, err := destEthClient.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "destEthClient.ChainID")
	}

	i.blockRepo = blockRepository
	i.eventRepo = eventRepository
	i.srcEthClient = srcEthClient

	i.bridge = srcBridge
	i.destBridge = destBridge
	i.signalService = signalService
	i.taikol1 = taikoL1

	i.blockBatchSize = cfg.BlockBatchSize
	i.numGoroutines = int(cfg.NumGoroutines)
	i.subscriptionBackoff = time.Duration(cfg.SubscriptionBackoff) * time.Second

	i.queue = q

	i.srcChainId = srcChainID
	i.destChainId = destChainID

	i.syncMode = cfg.SyncMode
	i.watchMode = cfg.WatchMode

	i.wg = &sync.WaitGroup{}

	i.ethClientTimeout = time.Duration(cfg.ETHClientTimeout) * time.Second

	i.numLatestBlocksToIgnoreWhenCrawling = cfg.NumLatestBlocksToIgnoreWhenCrawling

	i.targetBlockNumber = cfg.TargetBlockNumber

	i.mu = &sync.Mutex{}

	i.eventName = cfg.EventName

	return nil
}

// Name implements the SubcommandAction interface
func (i *Indexer) Name() string {
	return "indexer"
}

// Close waits for the wait groups internally to be stopped ,which will be done when the
// context is stopped externally by cmd/main.go shutdown.
func (i *Indexer) Close(ctx context.Context) {
	i.wg.Wait()
}

// Start starts the indexer, which should initialize the queue, add to wait groups,
// and start filtering or subscribing depending on the WatchMode provided.
// nolint: funlen
func (i *Indexer) Start() error {
	i.ctx = context.Background()

	if err := i.queue.Start(i.ctx, i.queueName()); err != nil {
		return err
	}

	i.wg.Add(1)

	go func() {
		defer func() {
			i.wg.Done()
		}()

		if err := i.filter(i.ctx); err != nil {
			slog.Error("error filtering blocks", "error", err.Error())
		}
	}()

	go func() {
		if err := backoff.Retry(func() error {
			return scanBlocks(i.ctx, i.srcEthClient, i.srcChainId, i.wg)
		}, backoff.NewConstantBackOff(5*time.Second)); err != nil {
			slog.Error("scan blocks backoff retry", "error", err)
		}
	}()

	go func() {
		if err := backoff.Retry(func() error {
			return i.queue.Notify(i.ctx, i.wg)
		}, backoff.NewConstantBackOff(5*time.Second)); err != nil {
			slog.Error("queue notify backoff retry", "error", err)
		}
	}()

	return nil
}

// filter is the main function run by Start in the indexer, which should filter on a loop,
// then if desired to subscribe, start subscriptions to events when done filtering.
func (i *Indexer) filter(ctx context.Context) error {
	// if subscribing to new events, skip filtering and subscribe only.
	if i.watchMode == Subscribe {
		return i.subscribe(ctx, i.srcChainId, i.destChainId)
	}

	syncMode := i.syncMode

	// always use Resync when crawling past blocks
	if i.watchMode == CrawlPastBlocks {
		syncMode = Resync
	}

	// set the initial processing block, which will vary by sync mode.
	if err := i.setInitialProcessingBlockByMode(ctx, syncMode, i.srcChainId); err != nil {
		return errors.Wrap(err, "i.setInitialProcessingBlockByMode")
	}

	// get the latest header
	header, err := i.srcEthClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "i.srcEthClient.HeaderByNumber")
	}

	// compare it to the processing block height, set above. if its equal to the latest block,
	// we dont actually need to filter - we should just subscribe, given our watch mode is not "Filter".
	if i.processingBlockHeight == header.Number.Uint64() {
		if i.watchMode == Filter {
			slog.Info("indexing caught up", "chainID", i.srcChainId.Uint64())

			return nil
		}

		slog.Info("indexing caught up, subscribing to new incoming events", "chainID", i.srcChainId.Uint64())

		return i.subscribe(ctx, i.srcChainId, i.destChainId)
	}

	// the end block is the latest header.
	endBlockID := header.Number.Uint64()

	// ignore latest N blocks if we are crawling past blocks, they are probably in queue already
	// and are not "missed", have just not been processed.
	if i.watchMode == CrawlPastBlocks {
		// if targetBlockNumber is not nil, we are just going to process a singular block.
		if i.targetBlockNumber != nil {
			slog.Info("targetBlockNumber is set", "targetBlockNumber", *i.targetBlockNumber)

			i.processingBlockHeight = *i.targetBlockNumber

			endBlockID = i.processingBlockHeight + 1
		} else if endBlockID > i.numLatestBlocksToIgnoreWhenCrawling {
			// otherwise, we need to set the endBlockID as the greater of the two:
			// either the endBlockID minus the number of latest blocks to ignore,
			// or endBlockID.
			endBlockID -= i.numLatestBlocksToIgnoreWhenCrawling
		}
	}

	slog.Info("fetching batch block events",
		"chainID", i.srcChainId.Uint64(),
		"processingBlockHeight", i.processingBlockHeight,
		"endblock", endBlockID,
		"batchsize", i.blockBatchSize,
		"watchMode", i.watchMode,
	)

	// iterate through from the starting block (i.processingBlockHeight) through the
	// latest block (endBlockID) in batches of i.blockBatchSize until we are finished.
	for j := i.processingBlockHeight; j < endBlockID; j += i.blockBatchSize {
		end := i.processingBlockHeight + i.blockBatchSize
		// if the end of the batch is greater than the latest block number, set end
		// to the latest block number
		if end > endBlockID {
			end = endBlockID
		}

		// filter exclusive of the end block.
		// we use "end" as the next starting point of the batch, and
		// process up to end - 1 for this batch.
		filterEnd := end - 1

		slog.Info("block batch", "start", j, "end", filterEnd)

		filterOpts := &bind.FilterOpts{
			Start:   i.processingBlockHeight,
			End:     &filterEnd,
			Context: ctx,
		}

		switch i.eventName {
		case relayer.EventNameMessageSent:
			if err := i.indexMessageSentEvents(ctx, filterOpts); err != nil {
				return errors.Wrap(err, "i.indexMessageSentEvents")
			}

			// we dont want to watch for message status changed events
			// when crawling past blocks on a loop. but otherwise,
			// we want to index all three event types when indexing MessageSent events,
			// since they are related.
			if i.watchMode != CrawlPastBlocks {
				if err := i.indexMessageStatusChangedEvents(ctx, filterOpts); err != nil {
					return errors.Wrap(err, "i.indexMessageStatusChangedEvents")
				}

				// we also want to index chain data synced events.
				if err := i.indexChainDataSyncedEvents(ctx, filterOpts); err != nil {
					return errors.Wrap(err, "i.indexChainDataSyncedEvents")
				}
			}
		case relayer.EventNameMessageReceived:
			if err := i.indexMessageReceivedEvents(ctx, filterOpts); err != nil {
				return errors.Wrap(err, "i.indexMessageReceivedEvents")
			}
		}

		// handle no events remaining, saving the processing block and continuing on
		// to the next batch.
		if err := i.handleNoEventsInBatch(ctx, i.srcChainId, int64(end)); err != nil {
			return errors.Wrap(err, "i.handleNoEventsInBatch")
		}
	}

	slog.Info(
		"indexer fully caught up",
	)

	// if we are crawling past blocks, we dont want to continue, we want to repeat the loop above
	// recursively.
	if i.watchMode == CrawlPastBlocks {
		slog.Info("restarting filtering from genesis")
		return i.filter(ctx)
	}

	slog.Info("getting latest block to see if header has advanced")

	latestBlock, err := i.srcEthClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "i.srcEthClient.HeaderByNumber")
	}

	if i.processingBlockHeight < latestBlock.Number.Uint64() {
		slog.Info("header has advanced",
			"processingBlockHeight", i.processingBlockHeight,
			"latestBlock", latestBlock.Number.Uint64(),
		)

		return i.filter(ctx)
	}

	// we are caught up and specified not to subscribe, we can return now and the indexer
	// is finished it's job.
	if i.watchMode == Filter {
		return nil
	}

	// otherwise, we subscribe to new events

	slog.Info("processing is caught up to latest block, subscribing to new blocks")

	return i.subscribe(ctx, i.srcChainId, i.destChainId)
}

// indexMessageSentEvents indexes `MessageSent` events on the bridge contract
// and stores them to the database, and adds the message to the queue if it has not been
// seen before.
func (i *Indexer) indexMessageSentEvents(ctx context.Context,
	filterOpts *bind.FilterOpts) error {
	events, err := i.bridge.FilterMessageSent(filterOpts, nil)
	if err != nil {
		return errors.Wrap(err, "bridge.FilterMessageSent")
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(i.numGoroutines)

	for events.Next() {
		event := events.Event

		group.Go(func() error {
			err := i.handleMessageSentEvent(groupCtx, i.srcChainId, event, false)
			if err != nil {
				relayer.ErrorEvents.Inc()
				// log error but always return nil to keep other goroutines active
				slog.Error("error handling event", "err", err.Error())
			} else {
				slog.Info("handled messagesent event successfully")
			}

			return nil
		})
	}

	// wait for the last of the goroutines to finish
	if err := group.Wait(); err != nil {
		return errors.Wrap(err, "group.Wait")
	}

	return nil
}

// indexMessageReceivedEvents indexes `MessageReceived` events on the bridge contract
// and stores them to the database, and adds the message to the queue if it has not been
// seen before.
func (i *Indexer) indexMessageReceivedEvents(ctx context.Context,
	filterOpts *bind.FilterOpts,
) error {
	events, err := i.bridge.FilterMessageReceived(filterOpts, nil)
	if err != nil {
		return errors.Wrap(err, "bridge.FilterMessageReceived")
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(i.numGoroutines)

	for events.Next() {
		event := events.Event

		group.Go(func() error {
			err := i.handleMessageReceivedEvent(groupCtx, i.srcChainId, event, false)
			if err != nil {
				relayer.MessageReceivedEventsIndexingErrors.Inc()
				// log error but always return nil to keep other goroutines active
				slog.Error("error handling event", "err", err.Error())
			} else {
				slog.Info("handled message received event successfully")
			}

			return nil
		})
	}

	// wait for the last of the goroutines to finish
	if err := group.Wait(); err != nil {
		return errors.Wrap(err, "group.Wait")
	}

	return nil
}

// indexMessageStatusChangedEvents indexes `MessageStatusChanged` events on the bridge contract
// and stores them to the database. It does not add them to any queue.
func (i *Indexer) indexMessageStatusChangedEvents(ctx context.Context,
	filterOpts *bind.FilterOpts) error {
	slog.Info("indexing messageStatusChanged events")

	events, err := i.bridge.FilterMessageStatusChanged(filterOpts, nil)
	if err != nil {
		return errors.Wrap(err, "bridge.FilterMessageStatusChanged")
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(i.numGoroutines)

	for events.Next() {
		event := events.Event

		group.Go(func() error {
			err := i.handleMessageStatusChangedEvent(groupCtx, i.srcChainId, event)
			if err != nil {
				relayer.MessageStatusChangedEventsIndexingErrors.Inc()
				// log error but always return nil to keep other goroutines active
				slog.Error("error handling messageStatusChanged", "err", err.Error())
			} else {
				slog.Info("handled messageStatusChanged event successfully")
			}

			return nil
		})
	}

	// wait for the last of the goroutines to finish
	if err := group.Wait(); err != nil {
		return errors.Wrap(err, "group.Wait")
	}

	slog.Info("done indexing messageStatusChanged events")

	return nil
}

// indexChainDataSyncedEvents indexes `ChainDataSynced` events on the bridge contract
// and stores them to the database. It does not add them to any queue. It only indexes
// the "STATE_ROOT" kind, not the "SIGNAL_ROOT" kind.
func (i *Indexer) indexChainDataSyncedEvents(ctx context.Context,
	filterOpts *bind.FilterOpts,
) error {
	slog.Info("indexing chainDataSynced events")

	chainDataSyncedEvents, err := i.signalService.FilterChainDataSynced(
		filterOpts,
		[]uint64{i.destChainId.Uint64()}, // only index intended events destination chain
		nil,
		[][32]byte{crypto.Keccak256Hash([]byte("STATE_ROOT"))}, // only index state root
	)
	if err != nil {
		return errors.Wrap(err, "bridge.FilterChainDataSynced")
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(i.numGoroutines)

	for chainDataSyncedEvents.Next() {
		event := chainDataSyncedEvents.Event

		group.Go(func() error {
			err := i.handleChainDataSyncedEvent(groupCtx, i.srcChainId, event, false)
			if err != nil {
				relayer.MessageStatusChangedEventsIndexingErrors.Inc()

				// log error but always return nil to keep other goroutines active
				slog.Error("error handling chainDataSynced", "err", err.Error())
			} else {
				slog.Info("handled chainDataSynced event successfully")
			}

			return nil
		})
	}

	// wait for the last of the goroutines to finish
	if err := group.Wait(); err != nil {
		return errors.Wrap(err, "group.Wait")
	}

	slog.Info("done indexing chainDataSynced events")

	return nil
}

// queueName builds out the name of a queue, in the format the processor will also
// use to listen to events.
func (i *Indexer) queueName() string {
	return fmt.Sprintf("%v-%v-%v-queue", i.srcChainId.String(), i.destChainId.String(), i.eventName)
}
