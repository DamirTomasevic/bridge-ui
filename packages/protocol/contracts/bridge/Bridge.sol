// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.18;

import {AddressResolver} from "../common/AddressResolver.sol";
import {EssentialContract} from "../common/EssentialContract.sol";
import {Proxied} from "../common/Proxied.sol";
import {IBridge} from "./IBridge.sol";
import {BridgeErrors} from "./BridgeErrors.sol";
import {LibBridgeData} from "./libs/LibBridgeData.sol";
import {LibBridgeProcess} from "./libs/LibBridgeProcess.sol";
import {LibBridgeRelease} from "./libs/LibBridgeRelease.sol";
import {LibBridgeRetry} from "./libs/LibBridgeRetry.sol";
import {LibBridgeSend} from "./libs/LibBridgeSend.sol";
import {LibBridgeStatus} from "./libs/LibBridgeStatus.sol";

/**
 * Bridge contract which is deployed on both L1 and L2. Mostly a thin wrapper
 * which calls the library implementations. See _IBridge_ for more details.
 * @dev The code hash for the same address on L1 and L2 may be different.
 * @custom:security-contact hello@taiko.xyz
 */
contract Bridge is EssentialContract, IBridge, BridgeErrors {
    using LibBridgeData for Message;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/

    LibBridgeData.State private _state; // 50 slots reserved

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/

    event MessageStatusChanged(
        bytes32 indexed msgHash, LibBridgeStatus.MessageStatus status, address transactor
    );

    event DestChainEnabled(uint256 indexed chainId, bool enabled);

    /*//////////////////////////////////////////////////////////////
                         USER-FACING FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// Allow Bridge to receive ETH from the TaikoL1, TokenVault or EtherVault.
    receive() external payable {
        if (
            msg.sender != resolve("token_vault", true) && msg.sender != resolve("ether_vault", true)
                && msg.sender != resolve("taiko", true) && msg.sender != owner()
        ) {
            revert B_CANNOT_RECEIVE();
        }
    }

    /// @dev Initializer to be called after being deployed behind a proxy.
    function init(address _addressManager) external initializer {
        EssentialContract._init(_addressManager);
    }

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

    function releaseEther(IBridge.Message calldata message, bytes calldata proof)
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

    function processMessage(Message calldata message, bytes calldata proof) external nonReentrant {
        return LibBridgeProcess.processMessage({
            state: _state,
            resolver: AddressResolver(this),
            message: message,
            proof: proof
        });
    }

    function retryMessage(Message calldata message, bool isLastAttempt) external nonReentrant {
        return LibBridgeRetry.retryMessage({
            state: _state,
            resolver: AddressResolver(this),
            message: message,
            isLastAttempt: isLastAttempt
        });
    }

    function isMessageSent(bytes32 msgHash) public view virtual returns (bool) {
        return LibBridgeSend.isMessageSent(AddressResolver(this), msgHash);
    }

    function isMessageReceived(bytes32 msgHash, uint256 srcChainId, bytes calldata proof)
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

    function isMessageFailed(bytes32 msgHash, uint256 destChainId, bytes calldata proof)
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

    function getMessageStatus(bytes32 msgHash)
        public
        view
        virtual
        returns (LibBridgeStatus.MessageStatus)
    {
        return LibBridgeStatus.getMessageStatus(msgHash);
    }

    function context() public view returns (Context memory) {
        return _state.ctx;
    }

    function isEtherReleased(bytes32 msgHash) public view returns (bool) {
        return _state.etherReleased[msgHash];
    }

    function isDestChainEnabled(uint256 _chainId) public view returns (bool enabled) {
        (enabled,) = LibBridgeSend.isDestChainEnabled(AddressResolver(this), _chainId);
    }

    function hashMessage(Message calldata message) public pure override returns (bytes32) {
        return LibBridgeData.hashMessage(message);
    }

    function getMessageStatusSlot(bytes32 msgHash) public pure returns (bytes32) {
        return LibBridgeStatus.getMessageStatusSlot(msgHash);
    }
}

contract ProxiedBridge is Proxied, Bridge {}
