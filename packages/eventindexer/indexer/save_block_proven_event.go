package indexer

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer/contracts/taikol1"
)

var (
	systemProver = common.HexToAddress("0x0000000000000000000000000000000000000001")
	oracleProver = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func (svc *Service) saveBlockProvenEvents(
	ctx context.Context,
	chainID *big.Int,
	events *taikol1.TaikoL1BlockProvenIterator,
) error {
	if !events.Next() || events.Event == nil {
		log.Infof("no blockProven events")
		return nil
	}

	for {
		event := events.Event

		if event.Raw.Removed {
			continue
		}

		log.Infof("blockProven by: %v", event.Prover.Hex())

		if err := svc.saveBlockProvenEvent(ctx, chainID, event); err != nil {
			eventindexer.BlockProvenEventsProcessedError.Inc()

			return errors.Wrap(err, "svc.saveBlockProvenEvent")
		}

		if !events.Next() {
			return nil
		}
	}
}

func (svc *Service) saveBlockProvenEvent(
	ctx context.Context,
	chainID *big.Int,
	event *taikol1.TaikoL1BlockProven,
) error {
	log.Infof("blockProven event found, id: %v", event.Id.Int64())

	marshaled, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "json.Marshal(event)")
	}

	_, err = svc.eventRepo.Save(ctx, eventindexer.SaveEventOpts{
		Name:    eventindexer.EventNameBlockProven,
		Data:    string(marshaled),
		ChainID: chainID,
		Event:   eventindexer.EventNameBlockProven,
		Address: event.Prover.Hex(),
	})
	if err != nil {
		return errors.Wrap(err, "svc.eventRepo.Save")
	}

	eventindexer.BlockProvenEventsProcessed.Inc()

	if event.Prover.Hex() != systemProver.Hex() && event.Prover.Hex() != oracleProver.Hex() {
		if err := svc.updateAverageBlockTime(ctx, event); err != nil {
			return errors.Wrap(err, "svc.updateAverageBlockTime")
		}
	}

	return nil
}

func (svc *Service) updateAverageBlockTime(ctx context.Context, event *taikol1.TaikoL1BlockProven) error {
	block, err := svc.taikol1.GetBlock(nil, event.Id)
	if err != nil {
		return errors.Wrap(err, "svc.taikoL1.GetBlock")
	}

	stat, err := svc.statRepo.Find(ctx)
	if err != nil {
		return errors.Wrap(err, "svc.statRepo.Find")
	}

	proposedAt := block.ProposedAt

	provenAt := time.Now().Unix()

	proofTime := uint64(provenAt) - proposedAt

	newAverageProofTime := calcNewAverage(stat.AverageProofTime, stat.NumProofs, proofTime)

	_, err = svc.statRepo.Save(ctx, eventindexer.SaveStatOpts{
		ProofTime: &newAverageProofTime,
	})
	if err != nil {
		return errors.Wrap(err, "svc.statRepo.Save")
	}

	return nil
}

func calcNewAverage(a, t, new uint64) uint64 {
	return ((a * t) + new) / (t + 1)
}
