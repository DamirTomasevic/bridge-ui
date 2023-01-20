// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.9;

import "../L1/TaikoData.sol";
import "../libs/LibTxDecoder.sol";
import "../libs/LibTxUtils.sol";
import "../thirdparty/LibRLPReader.sol";
import "../thirdparty/LibRLPWriter.sol";

/**
 * A library to invalidate a txList using the following rules:
 *
 * A txList is valid if and only if:
 * 1. The txList's length is no more than `maxBytesPerTxList`.
 * 2. The txList is well-formed RLP, with no additional trailing bytes.
 * 3. The total number of transactions is no more than
 *    `maxTransactionsPerBlock`.
 * 4. The sum of all transaction gas limit is no more than
 *    `blockMaxGasLimit`.
 *
 * A transaction is valid if and only if:
 * 1. The transaction is well-formed RLP, with no additional trailing bytes
 *    (rule #1 in Ethereum yellow paper).
 * 2. The transaction's signature is valid (rule #2 in Ethereum yellow paper).
 * 3. The transaction's the gas limit is no smaller than the intrinsic gas
 *    `minTxGasLimit` (rule #5 in Ethereum yellow paper).
 *
 * @title LibInvalidTxList
 * @author david <david@taiko.xyz>
 */
library LibInvalidTxList {
    // NOTE: If the order of this enum changes, then some test cases that using
    // this enum in generate_genesis.test.ts may also needs to be
    // modified accordingly.
    enum Reason {
        OK,
        BINARY_TOO_LARGE,
        BINARY_NOT_DECODABLE,
        BLOCK_TOO_MANY_TXS,
        BLOCK_GAS_LIMIT_TOO_LARGE,
        TX_INVALID_SIG,
        TX_GAS_LIMIT_TOO_SMALL
    }

    function isTxListInvalid(
        TaikoData.Config memory config,
        bytes calldata encoded,
        Reason hint,
        uint256 txIdx
    ) internal pure returns (Reason) {
        if (encoded.length > config.maxBytesPerTxList) {
            return Reason.BINARY_TOO_LARGE;
        }

        try LibTxDecoder.decodeTxList(config.chainId, encoded) returns (
            LibTxDecoder.TxList memory txList
        ) {
            if (txList.items.length > config.maxTransactionsPerBlock) {
                return Reason.BLOCK_TOO_MANY_TXS;
            }

            if (LibTxDecoder.sumGasLimit(txList) > config.blockMaxGasLimit) {
                return Reason.BLOCK_GAS_LIMIT_TOO_LARGE;
            }

            require(txIdx < txList.items.length, "invalid txIdx");
            LibTxDecoder.Tx memory _tx = txList.items[txIdx];

            if (hint == Reason.TX_INVALID_SIG) {
                require(
                    LibTxUtils.recoverSender(config.chainId, _tx) == address(0),
                    "bad hint TX_INVALID_SIG"
                );
                return Reason.TX_INVALID_SIG;
            }

            if (hint == Reason.TX_GAS_LIMIT_TOO_SMALL) {
                require(_tx.gasLimit >= config.minTxGasLimit, "bad hint");
                return Reason.TX_GAS_LIMIT_TOO_SMALL;
            }

            revert("failed to prove txlist invalid");
        } catch (bytes memory) {
            return Reason.BINARY_NOT_DECODABLE;
        }
    }
}
