package indexer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/taikochain/taiko-mono/packages/relayer"
	"golang.org/x/sync/errgroup"
)

var (
	eventName = relayer.EventNameMessageSent
)

// FilterThenSubscribe gets the most recent block height that has been indexed, and works it's way
// up to the latest block. As it goes, it tries to process messages.
// When it catches up, it then starts to Subscribe to latest events as they come in.
func (svc *Service) FilterThenSubscribe(ctx context.Context, mode relayer.Mode, watchMode relayer.WatchMode) error {
	chainID, err := svc.ethClient.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "svc.ethClient.ChainID()")
	}

	// if subscribing to new events, skip filtering and subscribe
	if watchMode == relayer.SubscribeWatchMode {
		return svc.subscribe(ctx, chainID)
	}

	if err := svc.setInitialProcessingBlockByMode(ctx, mode, chainID); err != nil {
		return errors.Wrap(err, "svc.setInitialProcessingBlockByMode")
	}

	header, err := svc.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "svc.ethClient.HeaderByNumber")
	}

	if svc.processingBlock.Height == header.Number.Uint64() {
		log.Info("caught up, subscribing to new incoming events")
		return svc.subscribe(ctx, chainID)
	}

	log.Infof("getting events between %v and %v in batches of %v",
		svc.processingBlock.Height,
		header.Number.Int64(),
		svc.blockBatchSize,
	)

	for i := svc.processingBlock.Height; i < header.Number.Uint64(); i += svc.blockBatchSize {
		end := svc.processingBlock.Height + svc.blockBatchSize
		// if the end of the batch is greater than the latest block number, set end
		// to the latest block number
		if end > header.Number.Uint64() {
			end = header.Number.Uint64()
		}

		log.Infof("batch from %v to %v", i, end)

		events, err := svc.bridge.FilterMessageSent(&bind.FilterOpts{
			Start:   svc.processingBlock.Height,
			End:     &end,
			Context: ctx,
		}, nil)
		if err != nil {
			return errors.Wrap(err, "bridge.FilterMessageSent")
		}

		if !events.Next() || events.Event == nil {
			if err := svc.handleNoEventsInBatch(ctx, chainID, int64(end)); err != nil {
				return errors.Wrap(err, "s.handleNoEventsInBatch")
			}

			continue
		}

		group, ctx := errgroup.WithContext(ctx)

		group.SetLimit(svc.numGoroutines)

		for {
			group.Go(func() error {
				err := svc.handleEvent(ctx, chainID, events.Event)
				if err != nil {
					// log error but always return nil to keep other goroutines active
					log.Error(err.Error())
				}

				return nil
			})

			if !events.Next() {
				if err := group.Wait(); err != nil {
					return errors.Wrap(err, "group.Wait")
				}

				if err := svc.handleNoEventsRemaining(ctx, chainID, events); err != nil {
					return errors.Wrap(err, "svc.handleNoEventsRemaining")
				}

				break
			}
		}
	}

	log.Info("indexer fully caught up, checking latest block number to see if it's advanced")

	latestBlock, err := svc.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "svc.ethclient.HeaderByNumber")
	}

	if svc.processingBlock.Height < latestBlock.Number.Uint64() {
		return svc.FilterThenSubscribe(ctx, relayer.SyncMode, watchMode)
	}

	// we are caught up and specified not to subscribe, we can return now
	if watchMode == relayer.FilterWatchMode {
		return nil
	}

	return svc.subscribe(ctx, chainID)
}
