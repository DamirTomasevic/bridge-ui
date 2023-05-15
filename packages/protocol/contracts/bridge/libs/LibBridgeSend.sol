// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.18;

import {AddressResolver} from "../../common/AddressResolver.sol";
import {IBridge} from "../IBridge.sol";
import {ISignalService} from "../../signal/ISignalService.sol";
import {LibAddress} from "../../libs/LibAddress.sol";
import {LibBridgeData} from "./LibBridgeData.sol";

/**
 * Entry point for starting a bridge transaction.
 *
 * @title LibBridgeSend
 */
library LibBridgeSend {
    using LibAddress for address;
    using LibBridgeData for IBridge.Message;

    error B_INCORRECT_VALUE();
    error B_OWNER_IS_NULL();
    error B_WRONG_CHAIN_ID();
    error B_WRONG_TO_ADDRESS();

    /**
     * Send a message to the Bridge with the details of the request. The Bridge
     * takes custody of the funds, unless the source chain is Taiko, in which
     * the funds are sent to and managed by the EtherVault.
     *
     * @param message Specifies the `depositValue`, `callValue`,
     * and `processingFee`. These must sum to `msg.value`. It also specifies the
     * `destChainId` which must have a `bridge` address set on the
     * AddressResolver and differ from the current chain ID.
     *
     * @return msgHash The hash of message sent.
     * This is picked up by an off-chain relayer which indicates a
     * bridge message has been sent and is ready to be processed on the
     * destination chain.
     */
    function sendMessage(
        LibBridgeData.State storage state,
        AddressResolver resolver,
        IBridge.Message memory message
    ) internal returns (bytes32 msgHash) {
        if (message.owner == address(0)) {
            revert B_OWNER_IS_NULL();
        }

        (bool destChainEnabled, address destChain) =
            isDestChainEnabled(resolver, message.destChainId);

        if (!destChainEnabled || message.destChainId == block.chainid) {
            revert B_WRONG_CHAIN_ID();
        }
        if (message.to == address(0) || message.to == destChain) {
            revert B_WRONG_TO_ADDRESS();
        }

        uint256 expectedAmount = message.depositValue + message.callValue + message.processingFee;

        if (expectedAmount != msg.value) {
            revert B_INCORRECT_VALUE();
        }

        // If on Taiko, send the expectedAmount to the EtherVault. Otherwise,
        // store it here on the Bridge. Processing will release Ether from the
        // EtherVault or the Bridge on the destination chain.
        address ethVault = resolver.resolve("ether_vault", true);
        ethVault.sendEther(expectedAmount);

        message.id = state.nextMessageId++;
        message.sender = msg.sender;
        message.srcChainId = block.chainid;

        msgHash = message.hashMessage();
        // Store a key which is the hash of this contract address and the
        // msgHash, with a value of 1.
        ISignalService(resolver.resolve("signal_service", false)).sendSignal(msgHash);
        emit LibBridgeData.MessageSent(msgHash, message);
    }

    function isDestChainEnabled(AddressResolver resolver, uint256 chainId)
        internal
        view
        returns (bool enabled, address destBridge)
    {
        destBridge = resolver.resolve(chainId, "bridge", true);
        enabled = destBridge != address(0);
    }

    function isMessageSent(AddressResolver resolver, bytes32 msgHash)
        internal
        view
        returns (bool)
    {
        return ISignalService(resolver.resolve("signal_service", false)).isSignalSent({
            app: address(this),
            signal: msgHash
        });
    }

    function isMessageReceived(
        AddressResolver resolver,
        bytes32 msgHash,
        uint256 srcChainId,
        bytes calldata proof
    ) internal view returns (bool) {
        address srcBridge = resolver.resolve(srcChainId, "bridge", false);
        return ISignalService(resolver.resolve("signal_service", false)).isSignalReceived({
            srcChainId: srcChainId,
            app: srcBridge,
            signal: msgHash,
            proof: proof
        });
    }
}
