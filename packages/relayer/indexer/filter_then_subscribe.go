package indexer

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/taikoxyz/taiko-mono/packages/relayer"
	"golang.org/x/sync/errgroup"
)

var (
	eventName = relayer.EventNameMessageSent
)

// FilterThenSubscribe gets the most recent block height that has been indexed, and works it's way
// up to the latest block. As it goes, it tries to process messages.
// When it catches up, it then starts to Subscribe to latest events as they come in.
func (svc *Service) FilterThenSubscribe(
	ctx context.Context,
	mode relayer.Mode,
	watchMode relayer.WatchMode,
) error {
	chainID, err := svc.ethClient.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "svc.ethClient.ChainID()")
	}

	go scanBlocks(ctx, svc.ethClient, chainID)

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

	if svc.processingBlockHeight == header.Number.Uint64() {
		slog.Info("indexing caught up, subscribing to new incoming events", "chainID", chainID.Uint64())
		return svc.subscribe(ctx, chainID)
	}

	slog.Info("fetching batch block events",
		"chainID", chainID.Uint64(),
		"startblock", svc.processingBlockHeight,
		"endblock", header.Number.Int64(),
		"batchsize", svc.blockBatchSize,
	)

	for i := svc.processingBlockHeight; i < header.Number.Uint64(); i += svc.blockBatchSize {
		end := svc.processingBlockHeight + svc.blockBatchSize
		// if the end of the batch is greater than the latest block number, set end
		// to the latest block number
		if end > header.Number.Uint64() {
			end = header.Number.Uint64()
		}

		// filter exclusive of the end block.
		// we use "end" as the next starting point of the batch, and
		// process up to end - 1 for this batch.
		filterEnd := end - 1

		fmt.Printf("block batch from %v to %v", i, filterEnd)
		fmt.Println()

		filterOpts := &bind.FilterOpts{
			Start:   svc.processingBlockHeight,
			End:     &filterEnd,
			Context: ctx,
		}

		messageStatusChangedEvents, err := svc.bridge.FilterMessageStatusChanged(filterOpts, nil)
		if err != nil {
			return errors.Wrap(err, "bridge.FilterMessageStatusChanged")
		}

		// we dont need to do anything with msgStatus events except save them to the DB.
		// we dont need to process them. they are for exposing via the API.

		err = svc.saveMessageStatusChangedEvents(ctx, chainID, messageStatusChangedEvents)
		if err != nil {
			return errors.Wrap(err, "bridge.saveMessageStatusChangedEvents")
		}

		messageSentEvents, err := svc.bridge.FilterMessageSent(filterOpts, nil)
		if err != nil {
			return errors.Wrap(err, "bridge.FilterMessageSent")
		}

		if !messageSentEvents.Next() || messageSentEvents.Event == nil {
			// use "end" not "filterEnd" here, because it will be used as the start
			// of the next batch.
			if err := svc.handleNoEventsInBatch(ctx, chainID, int64(end)); err != nil {
				return errors.Wrap(err, "svc.handleNoEventsInBatch")
			}

			continue
		}

		group, groupCtx := errgroup.WithContext(ctx)

		group.SetLimit(svc.numGoroutines)

		for {
			event := messageSentEvents.Event

			group.Go(func() error {
				err := svc.handleEvent(groupCtx, chainID, event)
				if err != nil {
					relayer.ErrorEvents.Inc()
					// log error but always return nil to keep other goroutines active
					log.Error(err.Error())
				}

				return nil
			})

			// if there are no more events
			if !messageSentEvents.Next() {
				// wait for the last of the goroutines to finish
				if err := group.Wait(); err != nil {
					return errors.Wrap(err, "group.Wait")
				}
				// handle no events remaining, saving the processing block and restarting the for
				// loop
				if err := svc.handleNoEventsInBatch(ctx, chainID, int64(end)); err != nil {
					return errors.Wrap(err, "svc.handleNoEventsInBatch")
				}

				break
			}
		}
	}

	log.Infof(
		"chain id %v indexer fully caught up, checking latest block number to see if it's advanced",
		chainID.Uint64(),
	)

	latestBlock, err := svc.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "svc.ethclient.HeaderByNumber")
	}

	if svc.processingBlockHeight < latestBlock.Number.Uint64() {
		return svc.FilterThenSubscribe(ctx, relayer.SyncMode, watchMode)
	}

	// we are caught up and specified not to subscribe, we can return now
	if watchMode == relayer.FilterWatchMode {
		return nil
	}

	return svc.subscribe(ctx, chainID)
}
