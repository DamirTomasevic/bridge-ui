// SPDX-License-Identifier: MIT
//
// ╭━━━━╮╱╱╭╮╱╱╱╱╱╭╮╱╱╱╱╱╭╮
// ┃╭╮╭╮┃╱╱┃┃╱╱╱╱╱┃┃╱╱╱╱╱┃┃
// ╰╯┃┃┣┻━┳┫┃╭┳━━╮┃┃╱╱╭━━┫╰━┳━━╮
// ╱╱┃┃┃╭╮┣┫╰╯┫╭╮┃┃┃╱╭┫╭╮┃╭╮┃━━┫
// ╱╱┃┃┃╭╮┃┃╭╮┫╰╯┃┃╰━╯┃╭╮┃╰╯┣━━┃
// ╱╱╰╯╰╯╰┻┻╯╰┻━━╯╰━━━┻╯╰┻━━┻━━╯
pragma solidity ^0.8.9;

import "../common/EssentialContract.sol";
import "./IBridge.sol";
import "./libs/LibBridgeData.sol";
import "./libs/LibBridgeProcess.sol";
import "./libs/LibBridgeRead.sol";
import "./libs/LibBridgeRetry.sol";
import "./libs/LibBridgeSend.sol";

/// @author dantaik <dan@taiko.xyz>
/// @dev The code hash for the same address on L1 and L2 may be different.
contract Bridge is EssentialContract, IBridge {
    using LibBridgeData for Message;

    /*********************
     * State Variables   *
     *********************/

    LibBridgeData.State private state; // 50 slots reserved
    uint256[50] private __gap;

    /*********************
     * Events            *
     *********************/

    event MessageSent(bytes32 indexed mhash, IBridge.Message message);

    event MessageStatusChanged(
        bytes32 indexed mhash,
        IBridge.MessageStatus status
    );

    event DestChainEnabled(uint256 indexed chainId, bool enabled);

    /*********************
     * External Functions*
     *********************/

    /// allow Bridge to receive ETH from EtherVault.
    receive() external payable {}

    /// @dev Initializer to be called after being deployed behind a proxy.
    function init(address _addressManager) external initializer {
        EssentialContract._init(_addressManager);
    }

    function sendMessage(Message calldata message)
        external
        payable
        nonReentrant
        returns (bytes32 mhash)
    {
        return LibBridgeSend.sendMessage(state, AddressResolver(this), message);
    }

    function processMessage(Message calldata message, bytes calldata proof)
        external
        nonReentrant
    {
        return
            LibBridgeProcess.processMessage(
                state,
                AddressResolver(this),
                message,
                proof
            );
    }

    function retryMessage(Message calldata message, bool lastAttempt)
        external
        nonReentrant
    {
        return
            LibBridgeRetry.retryMessage(
                state,
                AddressResolver(this),
                message,
                lastAttempt
            );
    }

    function enableDestChain(uint256 _chainId, bool enabled)
        external
        nonReentrant
    {
        LibBridgeSend.enableDestChain(state, _chainId, enabled);
    }

    /*********************
     * Public Functions  *
     *********************/

    function isMessageSent(bytes32 mhash) public view virtual returns (bool) {
        return LibBridgeRead.isMessageSent(mhash);
    }

    function isMessageReceived(
        bytes32 mhash,
        uint256 srcChainId,
        bytes calldata proof
    ) public view virtual returns (bool) {
        return
            LibBridgeRead.isMessageReceived(
                AddressResolver(this),
                mhash,
                srcChainId,
                proof
            );
    }

    function getMessageStatus(bytes32 mhash)
        public
        view
        virtual
        returns (MessageStatus)
    {
        return state.messageStatus[mhash];
    }

    function context() public view returns (Context memory) {
        return state.ctx;
    }

    function isDestChainEnabled(uint256 _chainId) public view returns (bool) {
        return LibBridgeRead.isDestChainEnabled(state, _chainId);
    }
}
