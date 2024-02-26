// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/
//
//   Email: security@taiko.xyz
//   Website: https://taiko.xyz
//   GitHub: https://github.com/taikoxyz
//   Discord: https://discord.gg/taikoxyz
//   Twitter: https://twitter.com/taikoxyz
//   Blog: https://mirror.xyz/labs.taiko.eth
//   Youtube: https://www.youtube.com/@taikoxyz

pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../../common/AddressResolver.sol";
import "../../libs/LibMath.sol";
import "../../verifiers/IVerifier.sol";
import "../tiers/ITierProvider.sol";
import "../TaikoData.sol";
import "./LibUtils.sol";

/// @title LibProving
/// @notice A library for handling block contestation and proving in the Taiko
/// protocol.
library LibProving {
    using LibMath for uint256;

    bytes32 public constant RETURN_LIVENESS_BOND = keccak256("RETURN_LIVENESS_BOND");
    bytes32 public constant TIER_OP = bytes32("tier_optimistic");

    // Warning: Any events defined here must also be defined in TaikoEvents.sol.
    event TransitionProved(
        uint256 indexed blockId,
        TaikoData.Transition tran,
        address prover,
        uint96 validityBond,
        uint16 tier
    );

    event TransitionContested(
        uint256 indexed blockId,
        TaikoData.Transition tran,
        address contester,
        uint96 contestBond,
        uint16 tier
    );

    event ProvingPaused(bool paused);

    // Warning: Any errors defined here must also be defined in TaikoErrors.sol.
    error L1_ALREADY_CONTESTED();
    error L1_ALREADY_PROVED();
    error L1_ASSIGNED_PROVER_NOT_ALLOWED();
    error L1_BLOCK_MISMATCH();
    error L1_INVALID_BLOCK_ID();
    error L1_INVALID_PAUSE_STATUS();
    error L1_INVALID_TIER();
    error L1_INVALID_TRANSITION();
    error L1_MISSING_VERIFIER();
    error L1_NOT_ASSIGNED_PROVER();
    error L1_UNEXPECTED_TRANSITION_TIER();

    function pauseProving(TaikoData.State storage state, bool toPause) external {
        if (state.slotB.provingPaused == toPause) revert L1_INVALID_PAUSE_STATUS();

        state.slotB.provingPaused = toPause;

        if (!toPause) {
            state.slotB.lastUnpausedAt = uint64(block.timestamp);
        }
        emit ProvingPaused(toPause);
    }

    /// @dev Proves or contests a block transition.
    function proveBlock(
        TaikoData.State storage state,
        TaikoData.Config memory config,
        AddressResolver resolver,
        TaikoData.BlockMetadata memory meta,
        TaikoData.Transition memory tran,
        TaikoData.TierProof memory proof
    )
        external
        returns (uint8 maxBlocksToVerify)
    {
        // Make sure parentHash is not zero
        // To contest an existing transition, simply use any non-zero value as
        // the blockHash and stateRoot.
        if (tran.parentHash == 0 || tran.blockHash == 0 || tran.stateRoot == 0) {
            revert L1_INVALID_TRANSITION();
        }

        // Check that the block has been proposed but has not yet been verified.
        TaikoData.SlotB memory b = state.slotB;
        if (meta.id <= b.lastVerifiedBlockId || meta.id >= b.numBlocks) {
            revert L1_INVALID_BLOCK_ID();
        }

        uint64 slot = meta.id % config.blockRingBufferSize;
        TaikoData.Block storage blk = state.blocks[slot];

        // Check the integrity of the block data. It's worth noting that in
        // theory, this check may be skipped, but it's included for added
        // caution.
        if (blk.blockId != meta.id || blk.metaHash != keccak256(abi.encode(meta))) {
            revert L1_BLOCK_MISMATCH();
        }

        // Each transition is uniquely identified by the parentHash, with the
        // blockHash and stateRoot open for later updates as higher-tier proofs
        // become available. In cases where a transition with the specified
        // parentHash does not exist, the transition ID (tid) will be set to 0.
        (uint32 tid, TaikoData.TransitionState storage ts) =
            _createTransition(state, blk, tran, slot);

        // The new proof must meet or exceed the minimum tier required by the
        // block or the previous proof; it cannot be on a lower tier.
        if (proof.tier == 0 || proof.tier < meta.minTier || proof.tier < ts.tier) {
            revert L1_INVALID_TIER();
        }

        // Retrieve the tier configurations. If the tier is not supported, the
        // subsequent action will result in a revert.
        ITierProvider.Tier memory tier =
            ITierProvider(resolver.resolve("tier_provider", false)).getTier(proof.tier);

        // Check if this prover is allowed to submit a proof for this block
        _checkProverPermission(state, blk, ts, tid, tier);

        // We must verify the proof, and any failure in proof verification will
        // result in a revert.
        //
        // It's crucial to emphasize that the proof can be assessed in two
        // potential modes: "proving mode" and "contesting mode." However, the
        // precise verification logic is defined within each tier's IVerifier
        // contract implementation. We simply specify to the verifier contract
        // which mode it should utilize - if the new tier is higher than the
        // previous tier, we employ the proving mode; otherwise, we employ the
        // contesting mode (the new tier cannot be lower than the previous tier,
        // this has been checked above).
        //
        // It's obvious that proof verification is entirely decoupled from
        // Taiko's core protocol.
        {
            address verifier = resolver.resolve(tier.verifierName, true);

            if (verifier != address(0)) {
                bool isContesting = proof.tier == ts.tier && tier.contestBond != 0;

                IVerifier.Context memory ctx = IVerifier.Context({
                    metaHash: blk.metaHash,
                    blobHash: meta.blobHash,
                    // Separate msgSender to allow the prover to be any address in the future.
                    prover: msg.sender,
                    msgSender: msg.sender,
                    blockId: blk.blockId,
                    isContesting: isContesting,
                    blobUsed: meta.blobUsed
                });

                IVerifier(verifier).verifyProof(ctx, tran, proof);
            } else if (tier.verifierName != TIER_OP) {
                // The verifier can be address-zero, signifying that there are no
                // proof checks for the tier. In practice, this only applies to
                // optimistic proofs.
                revert L1_MISSING_VERIFIER();
            }
        }

        bool isTopTier = tier.contestBond == 0;
        IERC20 tko = IERC20(resolver.resolve("taiko_token", false));

        if (isTopTier) {
            // A special return value from the top tier prover can signal this
            // contract to return all liveness bond.
            bool returnLivenessBond = blk.livenessBond > 0 && proof.data.length == 32
                && bytes32(proof.data) == RETURN_LIVENESS_BOND;

            if (returnLivenessBond) {
                tko.transfer(blk.assignedProver, blk.livenessBond);
                blk.livenessBond = 0;
            }
        }

        bool sameTransition = tran.blockHash == ts.blockHash && tran.stateRoot == ts.stateRoot;

        if (proof.tier > ts.tier) {
            // Handles the case when an incoming tier is higher than the current transition's tier.
            // Reverts when the incoming proof tries to prove the same transition
            // (L1_ALREADY_PROVED).
            _overrideWithHigherProof(ts, tran, proof, tier, tko, sameTransition);

            emit TransitionProved({
                blockId: blk.blockId,
                tran: tran,
                prover: msg.sender,
                validityBond: tier.validityBond,
                tier: proof.tier
            });
        } else {
            // New transition and old transition on the same tier - and if this transaction tries to
            // prove the same, it reverts
            if (sameTransition) revert L1_ALREADY_PROVED();

            if (isTopTier) {
                // The top tier prover re-proves.
                assert(tier.validityBond == 0);
                assert(ts.validityBond == 0 && ts.contestBond == 0 && ts.contester == address(0));

                ts.prover = msg.sender;
                ts.blockHash = tran.blockHash;
                ts.stateRoot = tran.stateRoot;

                emit TransitionProved({
                    blockId: blk.blockId,
                    tran: tran,
                    prover: msg.sender,
                    validityBond: 0,
                    tier: proof.tier
                });
            } else {
                // Contesting but not on the highest tier
                if (ts.contester != address(0)) revert L1_ALREADY_CONTESTED();

                // Burn the contest bond from the prover.
                tko.transferFrom(msg.sender, address(this), tier.contestBond);

                // We retain the contest bond within the transition, just in
                // case this configuration is altered to a different value
                // before the contest is resolved.
                //
                // It's worth noting that the previous value of ts.contestBond
                // doesn't have any significance.
                ts.contestBond = tier.contestBond;
                ts.contester = msg.sender;
                ts.contestations += 1;

                emit TransitionContested({
                    blockId: blk.blockId,
                    tran: tran,
                    contester: msg.sender,
                    contestBond: tier.contestBond,
                    tier: proof.tier
                });
            }
        }

        ts.timestamp = uint64(block.timestamp);
        return tier.maxBlocksToVerifyPerProof;
    }

    /// @dev Handle the transition initialization logic
    function _createTransition(
        TaikoData.State storage state,
        TaikoData.Block storage blk,
        TaikoData.Transition memory tran,
        uint64 slot
    )
        private
        returns (uint32 tid, TaikoData.TransitionState storage ts)
    {
        tid = LibUtils.getTransitionId(state, blk, slot, tran.parentHash);

        if (tid == 0) {
            // In cases where a transition with the provided parentHash is not
            // found, we must essentially "create" one and set it to its initial
            // state. This initial state can be viewed as a special transition
            // on tier-0.
            //
            // Subsequently, we transform this tier-0 transition into a
            // non-zero-tier transition with a proof. This approach ensures that
            // the same logic is applicable for both 0-to-non-zero transition
            // updates and non-zero-to-non-zero transition updates.
            unchecked {
                // Unchecked is safe:  Not realistic 2**32 different fork choice
                // per block will be proven and none of them is valid
                tid = blk.nextTransitionId++;
            }

            // Keep in mind that state.transitions are also reusable storage
            // slots, so it's necessary to reinitialize all transition fields
            // below.
            ts = state.transitions[slot][tid];
            ts.blockHash = 0;
            ts.stateRoot = 0;
            ts.validityBond = 0;
            ts.contester = address(0);
            ts.contestBond = 1; // to save gas
            ts.timestamp = blk.proposedAt;
            ts.tier = 0;
            ts.contestations = 0;

            if (tid == 1) {
                // This approach serves as a cost-saving technique for the
                // majority of blocks, where the first transition is expected to
                // be the correct one. Writing to `tran` is more economical
                // since it resides in the ring buffer, whereas writing to
                // `transitionIds` is not as cost-effective.
                ts.key = tran.parentHash;

                // In the case of this first transition, the block's assigned
                // prover has the privilege to re-prove it, but only when the
                // assigned prover matches the previous prover. To ensure this,
                // we establish the transition's prover as the block's assigned
                // prover. Consequently, when we carry out a 0-to-non-zero
                // transition update, the previous prover will consistently be
                // the block's assigned prover.
                //
                // While alternative implementations are possible, introducing
                // such changes would require additional if-else logic.
                ts.prover = blk.assignedProver;
            } else {
                // In scenarios where this transition is not the first one, we
                // straightforwardly reset the transition prover to address
                // zero.
                ts.prover = address(0);

                // Furthermore, we index the transition for future retrieval.
                // It's worth emphasizing that this mapping for indexing is not
                // reusable. However, given that the majority of blocks will
                // only possess one transition — the correct one — we don't need
                // to be concerned about the cost in this case.
                state.transitionIds[blk.blockId][tran.parentHash] = tid;

                // There is no need to initialize ts.key here because it's only used when tid == 1
            }
        } else {
            // A transition with the provided parentHash has been located.
            ts = state.transitions[slot][tid];
        }
    }

    /// @dev Handles what happens when there is a higher proof incoming
    function _overrideWithHigherProof(
        TaikoData.TransitionState storage ts,
        TaikoData.Transition memory tran,
        TaikoData.TierProof memory proof,
        ITierProvider.Tier memory tier,
        IERC20 tko,
        bool sameTransition
    )
        private
    {
        // Higher tier proof overwriting lower tier proof
        uint256 reward;

        if (ts.contester != address(0)) {
            if (sameTransition) {
                // The contested transition is proven to be valid, contestor loses the game
                reward = ts.contestBond >> 2;
                tko.transfer(ts.prover, ts.validityBond + reward);
            } else {
                // The contested transition is proven to be invalid, contestor wins the game
                reward = ts.validityBond >> 2;
                tko.transfer(ts.contester, ts.contestBond + reward);
            }
        } else {
            if (sameTransition) revert L1_ALREADY_PROVED();
            // Contest the existing transition and prove it to be invalid
            reward = ts.validityBond >> 1;
            ts.contestations += 1;
        }

        unchecked {
            if (reward > tier.validityBond) {
                tko.transfer(msg.sender, reward - tier.validityBond);
            } else {
                tko.transferFrom(msg.sender, address(this), tier.validityBond - reward);
            }
        }

        ts.validityBond = tier.validityBond;
        ts.contestBond = 1; // to save gas
        ts.contester = address(0);
        ts.prover = msg.sender;
        ts.tier = proof.tier;

        if (!sameTransition) {
            ts.blockHash = tran.blockHash;
            ts.stateRoot = tran.stateRoot;
        }
    }

    /// @dev Check the msg.sender (the new prover) against the block's assigned prover.
    function _checkProverPermission(
        TaikoData.State storage state,
        TaikoData.Block storage blk,
        TaikoData.TransitionState storage ts,
        uint32 tid,
        ITierProvider.Tier memory tier
    )
        private
        view
    {
        // The highest tier proof can always submit new proofs
        if (tier.contestBond == 0) return;

        bool inProvingWindow = uint256(ts.timestamp).max(state.slotB.lastUnpausedAt)
            + tier.provingWindow * 60 >= block.timestamp;
        bool isAssignedPover = msg.sender == blk.assignedProver;

        // The assigned prover can only submit the very first transition.
        if (tid == 1 && ts.tier == 0 && inProvingWindow) {
            if (!isAssignedPover) revert L1_NOT_ASSIGNED_PROVER();
        } else {
            // Disallow the same address to prove the block so that we can detect that the
            // assigned prover should not receive his liveness bond back
            if (isAssignedPover) revert L1_ASSIGNED_PROVER_NOT_ALLOWED();
        }
    }
}
