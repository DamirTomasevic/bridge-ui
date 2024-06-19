// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../../signal/ISignalService.sol";
import "./LibUtils.sol";

/// @title LibVerifying
/// @notice A library for handling block verification in the Taiko protocol.
/// @custom:security-contact security@taiko.xyz
library LibVerifying {
    using LibMath for uint256;

    struct Local {
        TaikoData.SlotB b;
        uint64 blockId;
        uint64 slot;
        uint64 numBlocksVerified;
        uint32 tid;
        uint32 lastVerifiedTransitionId;
        uint16 tier;
        bytes32 blockHash;
        bytes32 stateRoot;
        uint64 syncBlockId;
        address prover;
        ITierRouter tierRouter;
    }

    // Warning: Any errors defined here must also be defined in TaikoErrors.sol.
    error L1_BATCH_TRANSFER_FAILED();
    error L1_BLOCK_MISMATCH();
    error L1_INVALID_CONFIG();
    error L1_TRANSITION_ID_ZERO();
    error L1_TOO_LATE();

    /// @dev Verifies up to N blocks.
    function verifyBlocks(
        TaikoData.State storage _state,
        TaikoToken _tko,
        TaikoData.Config memory _config,
        IAddressResolver _resolver,
        uint64 _maxBlocksToVerify
    )
        internal
    {
        if (_maxBlocksToVerify == 0) {
            return;
        }

        Local memory local;
        local.b = _state.slotB;
        local.blockId = local.b.lastVerifiedBlockId;
        local.slot = local.blockId % _config.blockRingBufferSize;

        TaikoData.Block storage blk = _state.blocks[local.slot];
        if (blk.blockId != local.blockId) revert L1_BLOCK_MISMATCH();

        local.lastVerifiedTransitionId = blk.verifiedTransitionId;
        local.tid = local.lastVerifiedTransitionId;

        // The following scenario should never occur but is included as a
        // precaution.
        if (local.tid == 0) revert L1_TRANSITION_ID_ZERO();

        // The `blockHash` variable represents the most recently trusted
        // blockHash on L2.
        local.blockHash = _state.transitions[local.slot][local.tid].blockHash;

        // Unchecked is safe:
        // - assignment is within ranges
        // - blockId and numBlocksVerified values incremented will still be OK in the
        // next 584K years if we verifying one block per every second

        address[] memory provers = new address[](_maxBlocksToVerify);
        uint256[] memory bonds = new uint256[](_maxBlocksToVerify);

        unchecked {
            ++local.blockId;

            while (
                local.blockId < local.b.numBlocks && local.numBlocksVerified < _maxBlocksToVerify
            ) {
                local.slot = local.blockId % _config.blockRingBufferSize;

                blk = _state.blocks[local.slot];
                if (blk.blockId != local.blockId) revert L1_BLOCK_MISMATCH();

                local.tid = LibUtils.getTransitionId(_state, blk, local.slot, local.blockHash);
                // When `tid` is 0, it indicates that there is no proven
                // transition with its parentHash equal to the blockHash of the
                // most recently verified block.
                if (local.tid == 0) break;

                // A transition with the correct `parentHash` has been located.
                TaikoData.TransitionState storage ts = _state.transitions[local.slot][local.tid];

                // It's not possible to verify this block if either the
                // transition is contested and awaiting higher-tier proof or if
                // the transition is still within its cooldown period.
                local.tier = ts.tier;

                if (ts.contester != address(0)) {
                    break;
                } else {
                    if (local.tierRouter == ITierRouter(address(0))) {
                        local.tierRouter =
                            ITierRouter(_resolver.resolve(LibStrings.B_TIER_ROUTER, false));
                    }

                    uint24 cooldown = ITierProvider(local.tierRouter.getProvider(local.blockId))
                        .getTier(local.tier).cooldownWindow;

                    if (!LibUtils.isPostDeadline(ts.timestamp, local.b.lastUnpausedAt, cooldown)) {
                        // If cooldownWindow is 0, the block can theoretically
                        // be proved and verified within the same L1 block.
                        break;
                    }
                }

                // Update variables
                local.lastVerifiedTransitionId = local.tid;
                local.blockHash = ts.blockHash;
                local.prover = ts.prover;

                provers[local.numBlocksVerified] = local.prover;
                bonds[local.numBlocksVerified] = ts.validityBond;

                // Note: We exclusively address the bonds linked to the
                // transition used for verification. While there may exist
                // other transitions for this block, we disregard them entirely.
                // The bonds for these other transitions are burned (more precisely held in custody)
                // either when the transitions are generated or proven. In such cases, both the
                // provers and contesters of those transitions forfeit their bonds.

                emit LibUtils.BlockVerified({
                    blockId: local.blockId,
                    prover: local.prover,
                    blockHash: local.blockHash,
                    stateRoot: 0, // DEPRECATED and is always zero.
                    tier: local.tier
                });

                if (LibUtils.shouldSyncStateRoot(_config.stateRootSyncInternal, local.blockId)) {
                    local.stateRoot = ts.stateRoot;
                    local.syncBlockId = local.blockId;
                }

                ++local.blockId;
                ++local.numBlocksVerified;
            }

            if (local.numBlocksVerified != 0) {
                uint64 lastVerifiedBlockId = local.b.lastVerifiedBlockId + local.numBlocksVerified;
                local.slot = lastVerifiedBlockId % _config.blockRingBufferSize;

                _state.slotB.lastVerifiedBlockId = lastVerifiedBlockId;
                _state.blocks[local.slot].verifiedTransitionId = local.lastVerifiedTransitionId;

                // Resize the provers and bonds array
                uint256 newLen = local.numBlocksVerified;
                assembly {
                    mstore(provers, newLen)
                    mstore(bonds, newLen)
                }
                if (!_tko.batchTransfer(provers, bonds)) revert L1_BATCH_TRANSFER_FAILED();

                if (local.stateRoot != 0) {
                    _state.slotA.lastSyncedBlockId = local.syncBlockId;
                    _state.slotA.lastSynecdAt = uint64(block.timestamp);

                    ISignalService(_resolver.resolve(LibStrings.B_SIGNAL_SERVICE, false))
                        .syncChainData(
                        _config.chainId, LibStrings.H_STATE_ROOT, local.syncBlockId, local.stateRoot
                    );
                }
            }
        }
    }
}
