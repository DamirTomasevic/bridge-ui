// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import { AddressResolver } from "../../common/AddressResolver.sol";
import { IMintableERC20 } from "../../common/IMintableERC20.sol";
import { IProverPool } from "../ProverPool.sol";
import { ISignalService } from "../../signal/ISignalService.sol";
import { LibUtils } from "./LibUtils.sol";
import { LibMath } from "../../libs/LibMath.sol";
import { SafeCastUpgradeable } from
    "@openzeppelin/contracts-upgradeable/utils/math/SafeCastUpgradeable.sol";
import { TaikoData } from "../../L1/TaikoData.sol";
import { TaikoToken } from "../TaikoToken.sol";
import { LibL2Consts } from "../../L2/LibL2Consts.sol";

library LibVerifying {
    using SafeCastUpgradeable for uint256;
    using LibUtils for TaikoData.State;
    using LibMath for uint256;

    event BlockVerified(
        uint256 indexed blockId,
        bytes32 blockHash,
        address prover,
        uint64 blockFee,
        uint64 proofReward
    );
    event CrossChainSynced(
        uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot
    );

    error L1_INVALID_CONFIG();

    function init(
        TaikoData.State storage state,
        TaikoData.Config memory config,
        bytes32 genesisBlockHash,
        uint32 initFeePerGas,
        uint16 initAvgProofDelay
    )
        internal
    {
        if (
            config.chainId <= 1 //
                || config.blockMaxProposals == 1
                || config.blockRingBufferSize <= config.blockMaxProposals + 1
                || config.blockMaxGasLimit == 0 || config.blockMaxTransactions == 0
                || config.blockMaxTxListBytes == 0
                || config.blockTxListExpiry > 30 * 24 hours
                || config.blockMaxTxListBytes > 128 * 1024 //blob up to 128K
                || config.proofRegularCooldown < config.proofOracleCooldown
                || config.proofMinWindow == 0
                || config.proofMaxWindow < config.proofMinWindow
                || config.proofWindowMultiplier <= 100
                || config.ethDepositRingBufferSize <= 1
                || config.ethDepositMinCountPerBlock == 0
                || config.ethDepositMaxCountPerBlock
                    < config.ethDepositMinCountPerBlock
                || config.ethDepositMinAmount == 0
                || config.ethDepositMaxAmount <= config.ethDepositMinAmount
                || config.ethDepositMaxAmount >= type(uint96).max
                || config.ethDepositGas == 0 || config.ethDepositMaxFee == 0
                || config.ethDepositMaxFee >= type(uint96).max
                || config.ethDepositMaxFee
                    >= type(uint96).max / config.ethDepositMaxCountPerBlock
                || config.rewardOpenMultipler < 100
                || config.rewardMaxDelayPenalty >= 10_000
        ) revert L1_INVALID_CONFIG();

        unchecked {
            uint64 timeNow = uint64(block.timestamp);

            // Init state
            state.genesisHeight = uint64(block.number);
            state.genesisTimestamp = timeNow;
            state.numBlocks = 1;
            state.lastVerifiedAt = uint64(block.timestamp);
            state.feePerGas = initFeePerGas;
            state.avgProofDelay = initAvgProofDelay;

            // Init the genesis block
            TaikoData.Block storage blk = state.blocks[0];
            blk.nextForkChoiceId = 2;
            blk.verifiedForkChoiceId = 1;
            blk.proposedAt = timeNow;

            // Init the first fork choice
            TaikoData.ForkChoice storage fc = state.blocks[0].forkChoices[1];
            fc.blockHash = genesisBlockHash;
            fc.provenAt = timeNow;
        }

        emit BlockVerified({
            blockId: 0,
            blockHash: genesisBlockHash,
            prover: address(0),
            blockFee: 0,
            proofReward: 0
        });
    }

    function verifyBlocks(
        TaikoData.State storage state,
        TaikoData.Config memory config,
        AddressResolver resolver,
        uint256 maxBlocks
    )
        internal
    {
        uint256 i = state.lastVerifiedBlockId;
        TaikoData.Block storage blk =
            state.blocks[i % config.blockRingBufferSize];

        uint24 fcId = blk.verifiedForkChoiceId;
        assert(fcId > 0);

        bytes32 blockHash = blk.forkChoices[fcId].blockHash;
        uint32 gasUsed = blk.forkChoices[fcId].gasUsed;
        bytes32 signalRoot;

        uint64 processed;
        unchecked {
            ++i;
        }

        while (i < state.numBlocks && processed < maxBlocks) {
            blk = state.blocks[i % config.blockRingBufferSize];
            assert(blk.blockId == i);

            fcId = LibUtils.getForkChoiceId(state, blk, blockHash, gasUsed);
            if (fcId == 0) break;

            TaikoData.ForkChoice memory fc = blk.forkChoices[fcId];
            if (fc.prover == address(0)) break;

            uint256 proofRegularCooldown = fc.prover == address(1)
                ? config.proofOracleCooldown
                : config.proofRegularCooldown;

            if (block.timestamp <= fc.provenAt + proofRegularCooldown) break;

            blockHash = fc.blockHash;
            gasUsed = fc.gasUsed;
            signalRoot = fc.signalRoot;

            _verifyBlock({
                state: state,
                config: config,
                resolver: resolver,
                blk: blk,
                fcId: fcId,
                fc: fc
            });

            unchecked {
                ++i;
                ++processed;
            }
        }

        if (processed > 0) {
            unchecked {
                state.lastVerifiedAt = uint64(block.timestamp);
                state.lastVerifiedBlockId += processed;
            }

            if (config.relaySignalRoot) {
                // Send the L2's signal root to the signal service so other
                // TaikoL1  deployments, if they share the same signal
                // service, can relay the signal to their corresponding
                // TaikoL2 contract.
                ISignalService(resolver.resolve("signal_service", false))
                    .sendSignal(signalRoot);
            }
            emit CrossChainSynced(
                state.lastVerifiedBlockId, blockHash, signalRoot
            );
        }
    }

    function _verifyBlock(
        TaikoData.State storage state,
        TaikoData.Config memory config,
        AddressResolver resolver,
        TaikoData.Block storage blk,
        TaikoData.ForkChoice memory fc,
        uint24 fcId
    )
        private
    {
        // the actually mined L2 block's gasLimit is blk.gasLimit +
        // LibL2Consts.ANCHOR_GAS_COST, so fc.gasUsed may greater than
        // blk.gasLimit here.
        uint32 _gasLimit = blk.gasLimit + LibL2Consts.ANCHOR_GAS_COST;
        assert(fc.gasUsed <= _gasLimit);

        IProverPool proverPool =
            IProverPool(resolver.resolve("prover_pool", false));

        if (blk.assignedProver == address(0)) {
            --state.numOpenBlocks;
        } else if (!blk.proverReleased) {
            proverPool.releaseProver(blk.assignedProver);
        }

        // Reward the prover (including the oracle prover)
        uint64 proofReward =
            (config.blockFeeBaseGas + fc.gasUsed) * blk.rewardPerGas;

        if (fc.prover == address(1)) {
            // system prover is rewarded with `proofReward`.
        } else if (blk.assignedProver == address(0)) {
            // open prover is rewarded with more tokens
            proofReward = proofReward * config.rewardOpenMultipler / 100;
        } else if (blk.assignedProver != fc.prover) {
            // proving out side of the proof window, by a prover other
            // than the assigned prover
            proofReward = proofReward * config.rewardOpenMultipler / 100;
            proverPool.slashProver(blk.blockId, blk.assignedProver, proofReward);
        } else if (fc.provenAt <= blk.proposedAt + blk.proofWindow) {
            // proving inside the window, by the assigned prover
            uint64 proofDelay;
            unchecked {
                proofDelay = fc.provenAt - blk.proposedAt;

                if (config.rewardMaxDelayPenalty > 0) {
                    // Give the reward a penalty up to a small percentage.
                    // This will encourage prover to submit proof ASAP.
                    proofReward -= proofReward * proofDelay
                        * config.rewardMaxDelayPenalty / 10_000 / blk.proofWindow;
                }
            }

            // The selected prover managed to prove the block in time
            state.avgProofDelay = uint16(
                LibUtils.movingAverage({
                    maValue: state.avgProofDelay,
                    newValue: proofDelay,
                    maf: 7200
                })
            );

            state.feePerGas = uint32(
                LibUtils.movingAverage({
                    maValue: state.feePerGas,
                    newValue: blk.rewardPerGas,
                    maf: 7200
                })
            );
        } else {
            // proving out side of the proof window, by the assigned prover
            proverPool.slashProver(blk.blockId, blk.assignedProver, proofReward);
            proofReward = 0;
        }

        blk.verifiedForkChoiceId = fcId;

        // refund the proposer
        state.taikoTokenBalances[blk.proposer] +=
            (_gasLimit - fc.gasUsed) * blk.feePerGas;

        // Reward the prover
        state.taikoTokenBalances[fc.prover] += proofReward;

        emit BlockVerified({
            blockId: blk.blockId,
            blockHash: fc.blockHash,
            prover: fc.prover,
            blockFee: LibUtils.getBlockFee(state, config, fc.gasUsed),
            proofReward: proofReward
        });
    }
}
