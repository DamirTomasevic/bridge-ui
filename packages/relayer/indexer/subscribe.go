package indexer

import (
	"context"
	"math/big"

	"log/slog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/taikoxyz/taiko-mono/packages/relayer"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/bridge"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/signalservice"
)

// subscribe subscribes to latest events
func (i *Indexer) subscribe(ctx context.Context, chainID *big.Int, destChainID *big.Int) error {
	slog.Info("subscribing to new events")

	errChan := make(chan error)

	if i.eventName == relayer.EventNameMessageSent {
		go i.subscribeMessageSent(ctx, chainID, destChainID, errChan)

		go i.subscribeMessageStatusChanged(ctx, chainID, destChainID, errChan)

		go i.subscribeChainDataSynced(ctx, chainID, destChainID, errChan)
	} else if i.eventName == relayer.EventNameMessageReceived {
		go i.subscribeMessageReceived(ctx, chainID, destChainID, errChan)
	}

	// nolint: gosimple
	for {
		select {
		case <-ctx.Done():
			slog.Info("context finished")
			return nil
		case err := <-errChan:
			slog.Info("error encountered durign subscription", "error", err)

			relayer.ErrorsEncounteredDuringSubscription.Inc()

			return errors.Wrap(err, "errChan")
		}
	}
}

func (i *Indexer) subscribeMessageSent(
	ctx context.Context,
	chainID *big.Int,
	destChainID *big.Int,
	errChan chan error,
) {
	sink := make(chan *bridge.BridgeMessageSent)

	sub := event.ResubscribeErr(i.subscriptionBackoff, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			slog.Error("i.bridge.WatchMessageSent", "error", err)
		}

		slog.Info("resubscribing to WatchMessageSent events")

		return i.bridge.WatchMessageSent(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil)
	})

	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			slog.Info("context finished")
			return
		case err := <-sub.Err():
			errChan <- errors.Wrap(err, "sub.Err()")
		case event := <-sink:
			go func() {
				slog.Info("new message sent event", "msgHash", common.Hash(event.MsgHash).Hex(), "chainID", chainID.String())
				err := i.handleMessageSentEvent(ctx, chainID, event, true)

				if err != nil {
					slog.Error("i.subscribe, i.handleMessageSentEvent", "error", err)
					return
				}

				i.mu.Lock()

				defer i.mu.Unlock()

				block, err := i.blockRepo.GetLatestBlockProcessedForEvent(
					relayer.EventNameMessageSent,
					chainID,
					destChainID,
				)
				if err != nil {
					slog.Error("i.subscribe, blockRepo.GetLatestBlockProcessedForEvent", "error", err)
					return
				}

				if block.Height < event.Raw.BlockNumber {
					err = i.blockRepo.Save(relayer.SaveBlockOpts{
						Height:      event.Raw.BlockNumber,
						Hash:        event.Raw.BlockHash,
						ChainID:     chainID,
						DestChainID: destChainID,
						EventName:   relayer.EventNameMessageSent,
					})
					if err != nil {
						slog.Error("i.subscribe, i.blockRepo.Save", "error", err)
						return
					}

					relayer.BlocksProcessed.Inc()
				}
			}()
		}
	}
}

func (i *Indexer) subscribeMessageReceived(
	ctx context.Context,
	chainID *big.Int,
	destChainID *big.Int,
	errChan chan error,
) {
	sink := make(chan *bridge.BridgeMessageReceived)

	sub := event.ResubscribeErr(i.subscriptionBackoff, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			slog.Error("i.bridge.WatchMessageReceived", "error", err)
		}

		slog.Info("resubscribing to WatchMessageReceived events")

		return i.bridge.WatchMessageReceived(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil)
	})

	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			slog.Info("context finished")
			return
		case err := <-sub.Err():
			errChan <- errors.Wrap(err, "sub.Err()")
		case event := <-sink:
			go func() {
				slog.Info("new messageReceived event", "msgHash", common.Hash(event.MsgHash).Hex(), "chainID", chainID.String())
				err := i.handleMessageReceivedEvent(ctx, chainID, event, true)

				if err != nil {
					slog.Error("i.subscribe, i.handleMessageReceived", "error", err)
					return
				}

				i.mu.Lock()

				defer i.mu.Unlock()

				block, err := i.blockRepo.GetLatestBlockProcessedForEvent(
					relayer.EventNameMessageReceived,
					chainID,
					destChainID,
				)
				if err != nil {
					slog.Error("i.subscribe, blockRepo.GetLatestBlockProcessedForEvent", "error", err)
					return
				}

				if block.Height < event.Raw.BlockNumber {
					err = i.blockRepo.Save(relayer.SaveBlockOpts{
						Height:      event.Raw.BlockNumber,
						Hash:        event.Raw.BlockHash,
						ChainID:     chainID,
						DestChainID: destChainID,
						EventName:   relayer.EventNameMessageReceived,
					})
					if err != nil {
						slog.Error("i.subscribe, i.blockRepo.Save", "error", err)
						return
					}

					relayer.BlocksProcessed.Inc()
				}
			}()
		}
	}
}

func (i *Indexer) subscribeMessageStatusChanged(
	ctx context.Context,
	chainID *big.Int,
	destChainID *big.Int,
	errChan chan error) {
	sink := make(chan *bridge.BridgeMessageStatusChanged)

	sub := event.ResubscribeErr(i.subscriptionBackoff, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			slog.Error("i.bridge.WatchMessageStatusChanged", "error", err)
		}

		slog.Info("resubscribing to WatchMessageStatusChanged events")

		return i.bridge.WatchMessageStatusChanged(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil)
	})

	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			slog.Info("context finished")
			return
		case err := <-sub.Err():
			errChan <- errors.Wrap(err, "sub.Err()")
		case event := <-sink:
			slog.Info("new messageStatusChanged event",
				"msgHash", common.Hash(event.MsgHash).Hex(),
				"chainID", chainID.String(),
			)

			if err := i.saveMessageStatusChangedEvent(ctx, chainID, event); err != nil {
				slog.Error("i.subscribe, i.saveMessageStatusChangedEvent", "error", err)
			}
		}
	}
}

func (i *Indexer) subscribeChainDataSynced(
	ctx context.Context,
	chainID *big.Int,
	destChainID *big.Int,
	errChan chan error) {
	sink := make(chan *signalservice.SignalServiceChainDataSynced)

	sub := event.ResubscribeErr(i.subscriptionBackoff, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			slog.Error("i.signalService.WatchChainDataSynced", "error", err)
		}

		slog.Info("resubscribing to WatchChainDataSynced events", "destChainID", destChainID.Uint64())

		return i.signalService.WatchChainDataSynced(&bind.WatchOpts{
			Context: ctx,
		}, sink, []uint64{destChainID.Uint64()}, nil, nil)
	})

	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			slog.Info("context finished")
			return
		case err := <-sub.Err():
			errChan <- errors.Wrap(err, "sub.Err()")
		case event := <-sink:
			slog.Info("new chainDataSynced event",
				"signal", common.Hash(event.Signal).Hex(),
				"chainID", event.ChainId,
				"blockID", event.BlockId,
				"syncedInBlock", event.Raw.BlockNumber,
			)

			if err := i.handleChainDataSyncedEvent(ctx, i.srcChainId, event, true); err != nil {
				slog.Error("error handling chainDataSynced event", "error", err)
				continue
			}

			slog.Info("chainDataSynced event saved")
		}
	}
}
