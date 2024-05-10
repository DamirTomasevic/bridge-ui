package handler

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/utils"
	eventIterator "github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	guardianProverHeartbeater "github.com/taikoxyz/taiko-mono/packages/taiko-client/prover/guardian_prover_heartbeater"
	proofProducer "github.com/taikoxyz/taiko-mono/packages/taiko-client/prover/proof_producer"
	state "github.com/taikoxyz/taiko-mono/packages/taiko-client/prover/shared_state"
)

var (
	errL1Reorged                           = errors.New("L1 reorged")
	proofExpirationDelay                   = 6 * 12 * time.Second // 6 ethereum blocks
	submissionDelayRandomBumpRange float64 = 20
)

// BlockProposedEventHandler is responsible for handling the BlockProposed event as a prover.
type BlockProposedEventHandler struct {
	sharedState           *state.SharedState
	proverAddress         common.Address
	genesisHeightL1       uint64
	rpc                   *rpc.Client
	proofGenerationCh     chan<- *proofProducer.ProofWithHeader
	assignmentExpiredCh   chan<- *bindings.TaikoL1ClientBlockProposed
	proofSubmissionCh     chan<- *proofProducer.ProofRequestBody
	proofContestCh        chan<- *proofProducer.ContestRequestBody
	backOffRetryInterval  time.Duration
	backOffMaxRetrys      uint64
	contesterMode         bool
	proveUnassignedBlocks bool
	// Guardian prover related.
	isGuardian      bool
	submissionDelay time.Duration
}

// NewBlockProposedEventHandlerOps is the options for creating a new BlockProposedEventHandler.
type NewBlockProposedEventHandlerOps struct {
	SharedState           *state.SharedState
	ProverAddress         common.Address
	GenesisHeightL1       uint64
	RPC                   *rpc.Client
	ProofGenerationCh     chan *proofProducer.ProofWithHeader
	AssignmentExpiredCh   chan *bindings.TaikoL1ClientBlockProposed
	ProofSubmissionCh     chan *proofProducer.ProofRequestBody
	ProofContestCh        chan *proofProducer.ContestRequestBody
	BackOffRetryInterval  time.Duration
	BackOffMaxRetrys      uint64
	ContesterMode         bool
	ProveUnassignedBlocks bool
	SubmissionDelay       time.Duration
}

// NewBlockProposedEventHandler creates a new BlockProposedEventHandler instance.
func NewBlockProposedEventHandler(opts *NewBlockProposedEventHandlerOps) *BlockProposedEventHandler {
	return &BlockProposedEventHandler{
		opts.SharedState,
		opts.ProverAddress,
		opts.GenesisHeightL1,
		opts.RPC,
		opts.ProofGenerationCh,
		opts.AssignmentExpiredCh,
		opts.ProofSubmissionCh,
		opts.ProofContestCh,
		opts.BackOffRetryInterval,
		opts.BackOffMaxRetrys,
		opts.ContesterMode,
		opts.ProveUnassignedBlocks,
		false,
		opts.SubmissionDelay,
	}
}

// Handle implements the BlockProposedHandler interface.
func (h *BlockProposedEventHandler) Handle(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposedEventIterFunc,
) error {
	// If there are newly generated proofs, we need to submit them as soon as possible,
	// to avoid proof submission timeout.
	if len(h.proofGenerationCh) > 0 {
		log.Info("onBlockProposed callback early return", "proofGenerationChannelLength", len(h.proofGenerationCh))
		end()
		return nil
	}

	// Wait for the corresponding L2 block being mined in node.
	if _, err := h.rpc.WaitL2Header(ctx, e.BlockId); err != nil {
		return fmt.Errorf("failed to wait L2 header (eventID %d): %w", e.BlockId, err)
	}

	// Check if the L1 chain has reorged at first.
	if err := h.checkL1Reorg(ctx, e); err != nil {
		if err.Error() == errL1Reorged.Error() {
			end()
			return nil
		}

		return err
	}

	// If the current block is handled, just skip it.
	if e.BlockId.Uint64() <= h.sharedState.GetLastHandledBlockID() {
		return nil
	}

	log.Info(
		"New BlockProposed event",
		"l1Height", e.Raw.BlockNumber,
		"l1Hash", e.Raw.BlockHash,
		"blockID", e.BlockId,
		"removed", e.Raw.Removed,
		"assignedProver", e.AssignedProver,
		"blobHash", common.Bytes2Hex(e.Meta.BlobHash[:]),
		"livenessBond", utils.WeiToEther(e.LivenessBond),
		"minTier", e.Meta.MinTier,
		"blobUsed", e.Meta.BlobUsed,
	)
	metrics.ProverReceivedProposedBlockGauge.Set(float64(e.BlockId.Uint64()))

	// Move l1Current cursor.
	newL1Current, err := h.rpc.L1.HeaderByHash(ctx, e.Raw.BlockHash)
	if err != nil {
		return err
	}
	h.sharedState.SetL1Current(newL1Current)
	h.sharedState.SetLastHandledBlockID(e.BlockId.Uint64())

	// Try generating a proof for the proposed block with the given backoff policy.
	go func() {
		if err := backoff.Retry(
			func() error {
				if err := h.checkExpirationAndSubmitProof(ctx, e); err != nil {
					log.Error(
						"Failed to check proof status and submit proof",
						"error", err,
						"blockID", e.BlockId,
						"minTier", e.Meta.MinTier,
						"maxRetrys", h.backOffMaxRetrys,
					)
					return err
				}
				return nil
			},
			backoff.WithContext(
				backoff.WithMaxRetries(backoff.NewConstantBackOff(h.backOffRetryInterval), h.backOffMaxRetrys),
				ctx,
			),
		); err != nil {
			log.Error("Handle new BlockProposed event error", "error", err)
		}
	}()

	return nil
}

// checkL1Reorg checks whether the L1 chain has been reorged.
func (h *BlockProposedEventHandler) checkL1Reorg(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
) error {
	// Check whether the L2 EE's anchored L1 info, to see if the L1 chain has been reorged.
	reorgCheckResult, err := h.rpc.CheckL1Reorg(
		ctx,
		new(big.Int).Sub(e.BlockId, common.Big1),
	)
	if err != nil {
		return fmt.Errorf("failed to check whether L1 chain was reorged from L2EE (eventID %d): %w", e.BlockId, err)
	}

	if reorgCheckResult.IsReorged {
		log.Info(
			"Reset L1Current cursor due to reorg",
			"l1CurrentHeightOld", h.sharedState.GetL1Current().Number,
			"l1CurrentHeightNew", reorgCheckResult.L1CurrentToReset.Number,
			"lastHandledBlockIDOld", h.sharedState.GetLastHandledBlockID(),
			"lastHandledBlockIDNew", reorgCheckResult.LastHandledBlockIDToReset,
		)
		h.sharedState.SetL1Current(reorgCheckResult.L1CurrentToReset)
		if reorgCheckResult.LastHandledBlockIDToReset == nil {
			h.sharedState.SetLastHandledBlockID(0)
		} else {
			h.sharedState.SetLastHandledBlockID(reorgCheckResult.LastHandledBlockIDToReset.Uint64())
		}
		return errL1Reorged
	}

	lastL1OriginHeader, err := h.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(e.Meta.L1Height))
	if err != nil {
		return fmt.Errorf("failed to get L1 header, height %d: %w", e.Meta.L1Height, err)
	}

	if lastL1OriginHeader.Hash() != e.Meta.L1Hash {
		log.Warn(
			"L1 block hash mismatch due to L1 reorg",
			"height", e.Meta.L1Height,
			"lastL1OriginHeader", lastL1OriginHeader.Hash(),
			"l1HashInEvent", e.Meta.L1Hash,
		)

		return fmt.Errorf(
			"L1 block hash mismatch due to L1 reorg: %s != %s",
			lastL1OriginHeader.Hash(),
			e.Meta.L1Hash,
		)
	}

	return nil
}

// getRandomBumpedSubmissionDelay returns a random bumped submission delay.
func (h *BlockProposedEventHandler) getRandomBumpedSubmissionDelay(expiredAt time.Time) (time.Duration, error) {
	if h.submissionDelay == 0 {
		return h.submissionDelay, nil
	}

	randomBump, err := rand.Int(
		rand.Reader,
		new(big.Int).SetUint64(uint64(h.submissionDelay.Seconds()*submissionDelayRandomBumpRange/100)),
	)
	if err != nil {
		return 0, err
	}

	delay := time.Duration(h.submissionDelay.Seconds()+float64(randomBump.Uint64())) * time.Second

	if time.Since(expiredAt) >= delay {
		return 0, nil
	}

	return delay - time.Since(expiredAt), nil
}

// checkExpirationAndSubmitProof checks whether the proposed block's proving window is expired,
// and submits a new proof if necessary.
func (h *BlockProposedEventHandler) checkExpirationAndSubmitProof(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
) error {
	// Check whether the block has been verified.
	isVerified, err := isBlockVerified(ctx, h.rpc, e.BlockId)
	if err != nil {
		return fmt.Errorf("failed to check if the current L2 block is verified: %w", err)
	}
	if isVerified {
		log.Info("📋 Block has been verified", "blockID", e.BlockId)
		return nil
	}

	// Check whether the block's proof is still needed.
	proofStatus, err := rpc.GetBlockProofStatus(
		ctx,
		h.rpc,
		e.BlockId,
		h.proverAddress,
	)
	if err != nil {
		return fmt.Errorf("failed to check whether the L2 block needs a new proof: %w", err)
	}

	// If there is already a proof submitted on chain.
	if proofStatus.IsSubmitted {
		// If there is no need to contest the submitted proof, we skip proving this block here.
		if !proofStatus.Invalid {
			log.Info(
				"A valid proof has been submitted, skip proving",
				"blockID", e.BlockId,
				"parent", proofStatus.ParentHeader.Hash(),
			)
			return nil
		}

		// If there is an invalid proof, but current prover is not in contest mode, we skip proving this block.
		if !h.contesterMode {
			log.Info(
				"An invalid proof has been submitted, but current prover is not in contest mode, skip proving",
				"blockID", e.BlockId,
				"parent", proofStatus.ParentHeader.Hash(),
			)
			return nil
		}

		// If the current proof has not been contested, we should contest it at first.
		if proofStatus.CurrentTransitionState.Contester == rpc.ZeroAddress {
			h.proofContestCh <- &proofProducer.ContestRequestBody{
				BlockID:    e.BlockId,
				ProposedIn: new(big.Int).SetUint64(e.Raw.BlockNumber),
				ParentHash: proofStatus.ParentHeader.Hash(),
				Meta:       &e.Meta,
				Tier:       e.Meta.MinTier,
			}
		} else {
			// The invalid proof submitted to protocol is contested by another prover,
			// we need to submit a proof with a higher tier.
			h.proofSubmissionCh <- &proofProducer.ProofRequestBody{
				Tier:  proofStatus.CurrentTransitionState.Tier + 1,
				Event: e,
			}
		}

		return nil
	}

	windowExpired, expiredAt, timeToExpire, err := isProvingWindowExpired(e, h.sharedState.GetTiers())
	if err != nil {
		return fmt.Errorf("failed to check if the proving window is expired: %w", err)
	}

	// If the proving window is not expired, we need to check if the current prover is the assigned prover,
	// if no and the current prover wants to prove unassigned blocks, then we should wait for its expiration.
	if !windowExpired && e.AssignedProver != h.proverAddress {
		log.Info(
			"Proposed block is not provable by current prover at the moment",
			"blockID", e.BlockId,
			"prover", e.AssignedProver,
			"timeToExpire", timeToExpire,
		)

		if h.proveUnassignedBlocks {
			log.Info(
				"Add proposed block to wait for proof window expiration",
				"blockID", e.BlockId,
				"assignProver", e.AssignedProver,
				"timeToExpire", timeToExpire,
			)
			time.AfterFunc(
				// Add another 60 seconds, to ensure one more L1 block will be mined before the proof submission
				timeToExpire+proofExpirationDelay,
				func() { h.assignmentExpiredCh <- e },
			)
		}

		return nil
	}

	// The current prover is the assigned prover, or the proving window is expired,
	// try to submit a proof for this proposed block.
	tier := e.Meta.MinTier

	// Get a random bumped submission delay, if necessary.
	submissionDelay, err := h.getRandomBumpedSubmissionDelay(expiredAt)
	if err != nil {
		return err
	}

	if h.isGuardian {
		tier = encoding.TierGuardianMinorityID
	}

	log.Info(
		"Proposed block is provable",
		"blockID", e.BlockId,
		"assignProver", e.AssignedProver,
		"minTier", e.Meta.MinTier,
		"submissionDelay", submissionDelay,
		"tier", tier,
	)

	metrics.ProverProofsAssigned.Add(1)

	time.AfterFunc(submissionDelay, func() {
		h.proofSubmissionCh <- &proofProducer.ProofRequestBody{Tier: tier, Event: e}
	})

	return nil
}

// ========================= Guardian Prover =========================

// NewBlockProposedGuardianEventHandlerOps is the options for creating a new BlockProposedEventHandler.
type NewBlockProposedGuardianEventHandlerOps struct {
	*NewBlockProposedEventHandlerOps
	GuardianProverHeartbeater guardianProverHeartbeater.BlockSenderHeartbeater
}

// BlockProposedGuaridanEventHandler is responsible for handling the BlockProposed event as a guardian prover.
type BlockProposedGuaridanEventHandler struct {
	*BlockProposedEventHandler
	GuardianProverHeartbeater guardianProverHeartbeater.BlockSenderHeartbeater
}

// NewBlockProposedEventGuardianHandler creates a new BlockProposedEventHandler instance.
func NewBlockProposedEventGuardianHandler(
	opts *NewBlockProposedGuardianEventHandlerOps,
) *BlockProposedGuaridanEventHandler {
	blockProposedEventHandler := NewBlockProposedEventHandler(opts.NewBlockProposedEventHandlerOps)
	blockProposedEventHandler.isGuardian = true

	return &BlockProposedGuaridanEventHandler{
		BlockProposedEventHandler: blockProposedEventHandler,
		GuardianProverHeartbeater: opts.GuardianProverHeartbeater,
	}
}

// Handle implements the BlockProposedHandler interface.
func (h *BlockProposedGuaridanEventHandler) Handle(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposedEventIterFunc,
) error {
	// If we are operating as a guardian prover,
	// we should sign all seen proposed blocks as soon as possible.
	go func() {
		if err := h.GuardianProverHeartbeater.SignAndSendBlock(ctx, event.BlockId); err != nil {
			log.Error("Guardian prover unable to sign block", "blockID", event.BlockId, "error", err)
		}
	}()

	return h.BlockProposedEventHandler.Handle(ctx, event, end)
}
