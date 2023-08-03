// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import { TaikoData } from "../L1/TaikoData.sol";

/// @title TaikoConfig - Library for retrieving Taiko configuration parameters
library TaikoConfig {
    /// @dev Retrieves the Taiko configuration parameters
    /// @return config The Taiko configuration object

    function getConfig() internal pure returns (TaikoData.Config memory) {
        return TaikoData.Config({
            // Group 1: general configs
            chainId: 167_006,
            relaySignalRoot: false,
            // Group 2: block level configs
            // Two weeks if avg block time is 3 seconds
            blockMaxProposals: 403_200,
            blockRingBufferSize: 403_200 + 10,
            // Each time one more block is verified, there will be ~20k
            // more gas cost.
            blockMaxVerificationsPerTx: 10,
            // Set it to 6M, since its the upper limit of the Alpha-2
            // testnet's circuits.
            blockMaxGasLimit: 6_000_000,
            blockFeeBaseGas: 20_000,
            // Set it to 79  (+1 TaikoL2.anchor transaction = 80),
            // and 80 is the upper limit of the Alpha-2 testnet's circuits.
            blockMaxTransactions: 79,
            // Set it to 120KB, since 128KB is the upper size limit
            // of a geth transaction, so using 120KB for the proposed
            // transactions list calldata, 8K for the remaining tx fields.
            blockMaxTxListBytes: 120_000,
            blockTxListExpiry: 0,
            // Group 3: proof related configs
            proofRegularCooldown: 30 minutes,
            proofOracleCooldown: 15 minutes,
            proofMinWindow: 10 minutes,
            proofMaxWindow: 90 minutes,
            proofWindowMultiplier: 200, // 200%
            // Group 4: eth deposit related configs
            ethDepositRingBufferSize: 1024,
            ethDepositMinCountPerBlock: 8,
            ethDepositMaxCountPerBlock: 32,
            ethDepositMinAmount: 1 ether,
            ethDepositMaxAmount: 10_000 ether,
            ethDepositGas: 21_000,
            ethDepositMaxFee: 1 ether / 10,
            // Group 5: tokenomics
            rewardOpenMultipler: 150, // percentage
            rewardOpenMaxCount: 201_600, // blockMaxProposals / 2,
            rewardMaxDelayPenalty: 250 // bps
         });
    }
}
