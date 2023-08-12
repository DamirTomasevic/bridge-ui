// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import { AddressResolver } from "../common/AddressResolver.sol";
import { EssentialContract } from "../common/EssentialContract.sol";
import { Proxied } from "../common/Proxied.sol";
import { IBridge } from "./IBridge.sol";
import { BridgeErrors } from "./BridgeErrors.sol";
import { LibBridgeData } from "./libs/LibBridgeData.sol";
import { LibBridgeProcess } from "./libs/LibBridgeProcess.sol";
import { LibBridgeRelease } from "./libs/LibBridgeRelease.sol";
import { LibBridgeRetry } from "./libs/LibBridgeRetry.sol";
import { LibBridgeSend } from "./libs/LibBridgeSend.sol";
import { LibBridgeStatus } from "./libs/LibBridgeStatus.sol";

/// @title Bridge
/// @notice See the documentation for {IBridge}.
/// @dev The code hash for the same address on L1 and L2 may be different.
contract Bridge is EssentialContract, IBridge, BridgeErrors {
    using LibBridgeData for Message;

    LibBridgeData.State private _state; // 50 slots reserved

    event MessageStatusChanged(
        bytes32 indexed msgHash,
        LibBridgeStatus.MessageStatus status,
        address transactor
    );

    event DestChainEnabled(uint256 indexed chainId, bool enabled);

    receive() external payable { }

    /// @notice Initializes the contract.
    /// @param _addressManager The address of the {AddressManager} contract.
    function init(address _addressManager) external initializer {
        EssentialContract._init(_addressManager);
    }

    /// @notice Sends a message from the current chain to the destination chain
    /// specified in the message.
    /// @param message The message to send. (See {IBridge})
    /// @return msgHash The hash of the message that was sent.
    function sendMessage(Message calldata message)
        external
        payable
        nonReentrant
        returns (bytes32 msgHash)
    {
        return LibBridgeSend.sendMessage({
            state: _state,
            resolver: AddressResolver(this),
            message: message
        });
    }

    /// @notice Releases the Ether locked in the bridge as part of a cross-chain
    /// transfer.
    /// @param message The message containing the details of the Ether transfer.
    /// (See {IBridge})
    /// @param proof The proof of the cross-chain transfer.
    function releaseEther(
        IBridge.Message calldata message,
        bytes calldata proof
    )
        external
        nonReentrant
    {
        return LibBridgeRelease.releaseEther({
            state: _state,
            resolver: AddressResolver(this),
            message: message,
            proof: proof
        });
    }

    /// @notice Processes a message received from another chain.
    /// @param message The message to process.
    /// @param proof The proof of the cross-chain transfer.
    function processMessage(
        Message calldata message,
        bytes calldata proof
    )
        external
        nonReentrant
    {
        return LibBridgeProcess.processMessage({
            state: _state,
            resolver: AddressResolver(this),
            message: message,
            proof: proof
        });
    }

    /// @notice Retries sending a message that previously failed to send.
    /// @param message The message to retry.
    /// @param isLastAttempt Specifies whether this is the last attempt to send
    /// the message.
    function retryMessage(
        Message calldata message,
        bool isLastAttempt
    )
        external
        nonReentrant
    {
        return LibBridgeRetry.retryMessage({
            state: _state,
            resolver: AddressResolver(this),
            message: message,
            isLastAttempt: isLastAttempt
        });
    }

    /// @notice Check if the message with the given hash has been sent.
    /// @param msgHash The hash of the message.
    /// @return Returns true if the message has been sent, false otherwise.
    function isMessageSent(bytes32 msgHash)
        public
        view
        virtual
        returns (bool)
    {
        return LibBridgeSend.isMessageSent(AddressResolver(this), msgHash);
    }

    /// @notice Check if the message with the given hash has been received.
    /// @param msgHash The hash of the message.
    /// @param srcChainId The source chain ID.
    /// @param proof The proof of message receipt.
    /// @return Returns true if the message has been received, false otherwise.
    function isMessageReceived(
        bytes32 msgHash,
        uint256 srcChainId,
        bytes calldata proof
    )
        public
        view
        virtual
        override
        returns (bool)
    {
        return LibBridgeSend.isMessageReceived({
            resolver: AddressResolver(this),
            msgHash: msgHash,
            srcChainId: srcChainId,
            proof: proof
        });
    }

    /// @notice Check if the message with the given hash has failed.
    /// @param msgHash The hash of the message.
    /// @param destChainId The destination chain ID.
    /// @param proof The proof of message failure.
    /// @return Returns true if the message has failed, false otherwise.
    function isMessageFailed(
        bytes32 msgHash,
        uint256 destChainId,
        bytes calldata proof
    )
        public
        view
        virtual
        override
        returns (bool)
    {
        return LibBridgeStatus.isMessageFailed({
            resolver: AddressResolver(this),
            msgHash: msgHash,
            destChainId: destChainId,
            proof: proof
        });
    }

    /// @notice Get the status of the message with the given hash.
    /// @param msgHash The hash of the message.
    /// @return Returns the status of the message.
    function getMessageStatus(bytes32 msgHash)
        public
        view
        virtual
        returns (LibBridgeStatus.MessageStatus)
    {
        return LibBridgeStatus.getMessageStatus(msgHash);
    }

    /// @notice Get the current context.
    /// @return Returns the current context.
    function context() public view returns (Context memory) {
        return _state.ctx;
    }

    /// @notice Check if the Ether associated with the given message hash has
    /// been released.
    /// @param msgHash The hash of the message.
    /// @return Returns true if the Ether has been released, false otherwise.
    function isEtherReleased(bytes32 msgHash) public view returns (bool) {
        return _state.etherReleased[msgHash];
    }

    /// @notice Check if the destination chain with the given ID is enabled.
    /// @param _chainId The ID of the chain.
    /// @return enabled Returns true if the destination chain is enabled, false
    /// otherwise.
    function isDestChainEnabled(uint256 _chainId)
        public
        view
        returns (bool enabled)
    {
        (enabled,) =
            LibBridgeSend.isDestChainEnabled(AddressResolver(this), _chainId);
    }

    /// @notice Compute the hash of a given message.
    /// @param message The message to compute the hash for.
    /// @return Returns the hash of the message.
    function hashMessage(Message calldata message)
        public
        pure
        override
        returns (bytes32)
    {
        return LibBridgeData.hashMessage(message);
    }

    /// @notice Get the slot associated with a given message hash status.
    /// @param msgHash The hash of the message.
    /// @return Returns the slot associated with the given message hash status.
    function getMessageStatusSlot(bytes32 msgHash)
        public
        pure
        returns (bytes32)
    {
        return LibBridgeStatus.getMessageStatusSlot(msgHash);
    }
}

/// @title ProxiedBridge
/// @notice Proxied version of the Bridge contract.
contract ProxiedBridge is Proxied, Bridge { }
