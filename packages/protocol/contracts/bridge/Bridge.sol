// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../common/EssentialContract.sol";
import "../libs/LibAddress.sol";
import "../libs/LibMath.sol";
import "../signal/ISignalService.sol";
import "./IBridge.sol";

/// @title Bridge
/// @notice See the documentation for {IBridge}.
/// @dev Labeled in AddressResolver as "bridge". Additionally, the code hash for the same address on
/// L1 and L2 may be different.
/// @custom:security-contact security@taiko.xyz
contract Bridge is EssentialContract, IBridge {
    using Address for address;
    using LibMath for uint256;
    using LibAddress for address;
    using LibAddress for address payable;

    /// @dev The slot in transient storage of the call context. This is the keccak256 hash
    /// of "bridge.ctx_slot"
    bytes32 private constant _CTX_SLOT =
        0xe4ece82196de19aabe639620d7f716c433d1348f96ce727c9989a982dbadc2b9;

    /// @dev Place holder value when not using transient storage
    uint256 internal constant PLACEHOLDER = type(uint256).max;

    /// @notice The next message ID.
    /// @dev Slot 1.
    uint128 public nextMessageId;

    /// @notice Mapping to store the status of a message from its hash.
    /// @dev Slot 2.
    mapping(bytes32 msgHash => Status status) public messageStatus;

    /// @dev Slots 3 and 4
    Context private __ctx;

    /// @notice Mapping to store banned addresses.
    /// @dev Slot 5.
    uint256 private __reserved1;

    /// @notice Mapping to store the proof receipt of a message from its hash.
    /// @dev Slot 6.
    mapping(bytes32 msgHash => ProofReceipt receipt) public proofReceipt;

    uint256[44] private __gap;

    error B_INVALID_CHAINID();
    error B_INVALID_CONTEXT();
    error B_INVALID_GAS_LIMIT();
    error B_INVALID_STATUS();
    error B_INVALID_USER();
    error B_INVALID_VALUE();
    error B_MESSAGE_NOT_PROVEN();
    error B_MESSAGE_NOT_SENT();
    error B_MESSAGE_NOT_SUSPENDED();
    error B_MESSAGE_SUSPENDED();
    error B_NON_RETRIABLE();
    error B_NOT_ENOUGH_GASLEFT();
    error B_NOT_FAILED();
    error B_NOT_RECEIVED();
    error B_PERMISSION_DENIED();
    error B_STATUS_MISMATCH();
    error B_INVOCATION_TOO_EARLY();

    modifier sameChain(uint64 _chainId) {
        if (_chainId != block.chainid) revert B_INVALID_CHAINID();
        _;
    }

    /// @notice Function to receive Ether.
    receive() external payable { }

    /// @notice Initializes the contract.
    /// @param _owner The owner of this contract. msg.sender will be used if this value is zero.
    /// @param _addressManager The address of the {AddressManager} contract.
    function init(address _owner, address _addressManager) external initializer {
        __Essential_init(_owner, _addressManager);
    }

    /// @notice Suspend or unsuspend invocation for a list of messages.
    /// @param _msgHashes The array of msgHashes to be suspended.
    /// @param _suspend True if suspend, false if unsuspend.
    function suspendMessages(
        bytes32[] calldata _msgHashes,
        bool _suspend
    )
        external
        onlyFromOwnerOrNamed("bridge_watchdog")
    {
        for (uint256 i; i < _msgHashes.length; ++i) {
            bytes32 msgHash = _msgHashes[i];

            ProofReceipt storage receipt = proofReceipt[msgHash];
            uint64 _receivedAt = receipt.receivedAt;

            if (_suspend) {
                if (_receivedAt == 0) revert B_MESSAGE_NOT_PROVEN();
                if (_receivedAt == type(uint64).max) revert B_MESSAGE_SUSPENDED();

                receipt.receivedAt = type(uint64).max;
                emit MessageSuspended(msgHash, true, 0);
            } else {
                // Note before we set the receivedAt to current timestamp, we have to be really
                // careful that this message must have been proven then suspended.
                if (_receivedAt != type(uint64).max) revert B_MESSAGE_NOT_SUSPENDED();
                receipt.receivedAt = uint64(block.timestamp);
                emit MessageSuspended(msgHash, false, uint64(block.timestamp));
            }
        }
    }

    /// @inheritdoc IBridge
    function sendMessage(Message calldata _message)
        external
        payable
        override
        whenNotPaused
        nonReentrant
        returns (bytes32 msgHash_, Message memory message_)
    {
        // Ensure the message owner is not null.
        if (_message.srcOwner == address(0) || _message.destOwner == address(0)) {
            revert B_INVALID_USER();
        }

        // Check if the destination chain is enabled.
        (bool destChainEnabled,) = isDestChainEnabled(_message.destChainId);

        // Verify destination chain and to address.
        if (!destChainEnabled) revert B_INVALID_CHAINID();
        if (_message.destChainId == block.chainid) {
            revert B_INVALID_CHAINID();
        }

        // Ensure the sent value matches the expected amount.
        uint256 expectedAmount = _message.value + _message.fee;
        if (expectedAmount != msg.value) revert B_INVALID_VALUE();

        message_ = _message;

        // Configure message details and send signal to indicate message sending.
        message_.id = nextMessageId++;
        message_.from = msg.sender;
        message_.srcChainId = uint64(block.chainid);

        msgHash_ = hashMessage(message_);

        emit MessageSent(msgHash_, message_);
        ISignalService(resolve("signal_service", false)).sendSignal(msgHash_);
    }

    /// @inheritdoc IBridge
    function recallMessage(
        Message calldata _message,
        bytes calldata _proof
    )
        external
        whenNotPaused
        sameChain(_message.srcChainId)
        nonReentrant
    {
        bytes32 msgHash = hashMessage(_message);

        if (messageStatus[msgHash] != Status.NEW) revert B_STATUS_MISMATCH();

        uint64 receivedAt = proofReceipt[msgHash].receivedAt;
        if (receivedAt == type(uint64).max) revert B_MESSAGE_SUSPENDED();

        (uint256 invocationDelay,) = getInvocationDelays();

        bool isNewlyProven;
        if (receivedAt == 0) {
            address signalService = resolve("signal_service", false);

            if (!ISignalService(signalService).isSignalSent(address(this), msgHash)) {
                revert B_MESSAGE_NOT_SENT();
            }

            bytes32 failureSignal = signalForFailedMessage(msgHash);
            if (!_proveSignalReceived(signalService, failureSignal, _message.destChainId, _proof)) {
                revert B_NOT_FAILED();
            }

            receivedAt = uint64(block.timestamp);
            isNewlyProven = true;

            if (invocationDelay != 0) {
                proofReceipt[msgHash].receivedAt = receivedAt;
            }
        }

        if (_isPostInvocationDelay(receivedAt, invocationDelay)) {
            delete proofReceipt[msgHash];
            messageStatus[msgHash] = Status.RECALLED;

            // Execute the recall logic based on the contract's support for the
            // IRecallableSender interface
            if (_message.from.supportsInterface(type(IRecallableSender).interfaceId)) {
                _storeContext(msgHash, address(this), _message.srcChainId);

                // Perform recall
                IRecallableSender(_message.from).onMessageRecalled{ value: _message.value }(
                    _message, msgHash
                );

                // Must reset the context after the message call
                _resetContext();
            } else {
                _message.srcOwner.sendEtherAndVerify(_message.value);
            }
            emit MessageRecalled(msgHash);
        } else if (isNewlyProven) {
            emit MessageReceived(msgHash, _message, true);
        } else {
            revert B_INVOCATION_TOO_EARLY();
        }
    }

    /// @inheritdoc IBridge
    function processMessage(
        Message calldata _message,
        bytes calldata _proof
    )
        external
        whenNotPaused
        sameChain(_message.destChainId)
        nonReentrant
    {
        bytes32 msgHash = hashMessage(_message);
        if (messageStatus[msgHash] != Status.NEW) revert B_STATUS_MISMATCH();

        address signalService = resolve("signal_service", false);

        uint64 receivedAt = proofReceipt[msgHash].receivedAt;
        if (receivedAt == type(uint64).max) revert B_MESSAGE_SUSPENDED();

        (uint256 invocationDelay, uint256 invocationExtraDelay) = getInvocationDelays();

        bool isNewlyProven;
        if (receivedAt == 0) {
            if (!_proveSignalReceived(signalService, msgHash, _message.srcChainId, _proof)) {
                revert B_NOT_RECEIVED();
            }

            receivedAt = uint64(block.timestamp);
            isNewlyProven = true;

            if (invocationDelay != 0) {
                proofReceipt[msgHash] = ProofReceipt({
                    receivedAt: receivedAt,
                    preferredExecutor: _message.gasLimit == 0 ? _message.destOwner : msg.sender
                });
            }
        }

        if (invocationDelay != 0 && msg.sender != proofReceipt[msgHash].preferredExecutor) {
            // If msg.sender is not the one that proved the message, then there
            // is an extra delay.
            unchecked {
                invocationDelay += invocationExtraDelay;
            }
        }

        if (_isPostInvocationDelay(receivedAt, invocationDelay)) {
            // If the gas limit is set to zero, only the owner can process the message.
            if (_message.gasLimit == 0 && msg.sender != _message.destOwner) {
                revert B_PERMISSION_DENIED();
            }

            delete proofReceipt[msgHash];

            uint256 refundAmount;

            // Process message differently based on the target address
            if (
                _message.to == address(0) || _message.to == address(this)
                    || _message.to == signalService
            ) {
                // Handle special addresses that don't require actual invocation but
                // mark message as DONE
                refundAmount = _message.value;
                _updateMessageStatus(msgHash, Status.DONE);
            } else {
                uint256 gasLimit;
                if (msg.sender == _message.destOwner) {
                    // Use the remaining gas if called by a the destOwner, else
                    // use the specified gas limit.
                    gasLimit = gasleft();
                } else {
                    // The "1/64th rule" refers to the gasleft at the time the call is made. When a
                    // contract makes a call to another contract, it can only forward 63/64 of the
                    // gas remaining (gasleft) at that moment, ensuring that there is always some
                    // gas reserved for the calling contract to complete its execution after the
                    // called contract finishes. This does not necessarily relate to the gas amount
                    // specified in the call itself, but rather to the actual remaining gas at the
                    // time of the call.
                    //
                    // See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-150.md
                    if (_message.gasLimit > (gasleft() * 63) >> 6) revert B_NOT_ENOUGH_GASLEFT();

                    gasLimit = _message.gasLimit;
                }

                if (_invokeMessageCall(_message, msgHash, gasLimit)) {
                    _updateMessageStatus(msgHash, Status.DONE);
                } else {
                    _updateMessageStatus(msgHash, Status.RETRIABLE);
                }
            }

            // Determine the refund recipient
            address refundTo =
                _message.refundTo == address(0) ? _message.destOwner : _message.refundTo;

            // Refund the processing fee
            if (msg.sender == refundTo) {
                refundTo.sendEtherAndVerify(_message.fee + refundAmount);
            } else {
                // If sender is another address, reward it and refund the rest
                msg.sender.sendEtherAndVerify(_message.fee);
                refundTo.sendEtherAndVerify(refundAmount);
            }
            emit MessageExecuted(msgHash);
        } else if (isNewlyProven) {
            emit MessageReceived(msgHash, _message, false);
        } else {
            revert B_INVOCATION_TOO_EARLY();
        }
    }

    /// @inheritdoc IBridge
    function retryMessage(
        Message calldata _message,
        bool _isLastAttempt
    )
        external
        whenNotPaused
        sameChain(_message.destChainId)
        nonReentrant
    {
        // If the gasLimit is set to 0 or isLastAttempt is true, the caller must
        // be the message.destOwner.
        if (_message.gasLimit == 0 || _isLastAttempt) {
            if (msg.sender != _message.destOwner) revert B_PERMISSION_DENIED();
        }

        bytes32 msgHash = hashMessage(_message);
        if (messageStatus[msgHash] != Status.RETRIABLE) {
            revert B_NON_RETRIABLE();
        }

        // Attempt to invoke the messageCall.
        if (_invokeMessageCall(_message, msgHash, gasleft())) {
            _updateMessageStatus(msgHash, Status.DONE);
        } else if (_isLastAttempt) {
            _updateMessageStatus(msgHash, Status.FAILED);
        }
        emit MessageRetried(msgHash);
    }

    /// @inheritdoc IBridge
    function isMessageSent(Message calldata _message) external view returns (bool) {
        if (_message.srcChainId != block.chainid) return false;
        return ISignalService(resolve("signal_service", false)).isSignalSent({
            _app: address(this),
            _signal: hashMessage(_message)
        });
    }

    /// @notice Checks if a msgHash has failed on its destination chain and caches cross-chain data
    /// if requested.
    /// @param _message The message.
    /// @param _proof The merkle inclusion proof.
    /// @return true if the message has failed, false otherwise.
    function proveMessageFailed(
        Message calldata _message,
        bytes calldata _proof
    )
        external
        returns (bool)
    {
        if (_message.srcChainId != block.chainid) return false;

        return _proveSignalReceived(
            resolve("signal_service", false),
            signalForFailedMessage(hashMessage(_message)),
            _message.destChainId,
            _proof
        );
    }

    /// @notice Checks if a msgHash has failed on its destination chain and caches cross-chain data
    /// if requested.
    /// @param _message The message.
    /// @param _proof The merkle inclusion proof.
    /// @return true if the message has been received, false otherwise.
    function proveMessageReceived(
        Message calldata _message,
        bytes calldata _proof
    )
        external
        returns (bool)
    {
        if (_message.destChainId != block.chainid) return false;
        return _proveSignalReceived(
            resolve("signal_service", false), hashMessage(_message), _message.srcChainId, _proof
        );
    }

    /// @notice Checks if a msgHash has failed on its destination chain.
    /// This is the 'readonly' version of proveMessageFailed.
    /// @param _message The message.
    /// @param _proof The merkle inclusion proof.
    /// @return true if the message has failed, false otherwise.
    function isMessageFailed(
        Message calldata _message,
        bytes calldata _proof
    )
        external
        view
        returns (bool)
    {
        if (_message.srcChainId != block.chainid) return false;

        return _isSignalReceived(
            resolve("signal_service", false),
            signalForFailedMessage(hashMessage(_message)),
            _message.destChainId,
            _proof
        );
    }

    /// @notice Checks if a msgHash has failed on its destination chain.
    /// This is the 'readonly' version of proveMessageReceived.
    /// @param _message The message.
    /// @param _proof The merkle inclusion proof.
    /// @return true if the message has failed, false otherwise.
    function isMessageReceived(
        Message calldata _message,
        bytes calldata _proof
    )
        external
        view
        returns (bool)
    {
        if (_message.destChainId != block.chainid) return false;
        return _isSignalReceived(
            resolve("signal_service", false), hashMessage(_message), _message.srcChainId, _proof
        );
    }

    /// @notice Checks if the destination chain is enabled.
    /// @param _chainId The destination chain ID.
    /// @return enabled_ True if the destination chain is enabled.
    /// @return destBridge_ The bridge of the destination chain.
    function isDestChainEnabled(uint64 _chainId)
        public
        view
        returns (bool enabled_, address destBridge_)
    {
        destBridge_ = resolve(_chainId, "bridge", true);
        enabled_ = destBridge_ != address(0);
    }

    /// @notice Gets the current context.
    /// @inheritdoc IBridge
    function context() external view returns (Context memory ctx_) {
        ctx_ = _loadContext();
        if (ctx_.msgHash == 0 || ctx_.msgHash == bytes32(PLACEHOLDER)) {
            revert B_INVALID_CONTEXT();
        }
    }

    /// @notice Returns invocation delay values.
    /// @dev Bridge contract deployed on L1 shall use a non-zero value for better
    /// security.
    /// @return The minimal delay in second before a message can be executed since and the time it
    /// was received on the this chain.
    /// @return The extra delay in second (to be added to invocationDelay) if the transactor is not
    /// the preferredExecutor who proved this message.
    function getInvocationDelays() public view virtual returns (uint256, uint256) {
        if (LibNetwork.isEthereumMainnetOrTestnet(block.chainid)) {
            // For Taiko mainnet and public testnets
            // 384 seconds = 6.4 minutes = one ethereum epoch
            return (1 hours, 384 seconds);
        } else if (LibNetwork.isTaikoDevnet(block.chainid)) {
            return (5 minutes, 384 seconds);
        } else {
            // This is a Taiko L2 chain where no deleys are applied.
            return (0, 0);
        }
    }

    /// @inheritdoc IBridge
    function hashMessage(Message memory _message) public pure returns (bytes32) {
        return keccak256(abi.encode("TAIKO_MESSAGE", _message));
    }

    /// @notice Returns a signal representing a failed/recalled message.
    /// @param _msgHash The message hash.
    /// @return The failed representation of it as bytes32.
    function signalForFailedMessage(bytes32 _msgHash) public pure returns (bytes32) {
        return _msgHash ^ bytes32(uint256(Status.FAILED));
    }

    /// @notice Checks if the given address can pause and/or unpause the bridge.
    /// @dev Considering that the watchdog is a hot wallet, in case its private key is leaked, we
    /// only allow watchdog to pause the bridge, but does not allow it to unpause the bridge.
    function _authorizePause(address addr, bool toPause) internal view override {
        // Owenr and chain_pauser can pause/unpause the bridge.
        if (addr == owner() || addr == resolve("chain_pauser", true)) return;

        // bridge_watchdog can pause the bridge, but cannot unpause it.
        if (toPause && addr == resolve("bridge_watchdog", true)) return;

        revert RESOLVER_DENIED();
    }

    /// @notice Invokes a call message on the Bridge.
    /// @param _message The call message to be invoked.
    /// @param _msgHash The hash of the message.
    /// @param _gasLimit The gas limit for the message call.
    /// @return success_ A boolean value indicating whether the message call was
    /// successful.
    /// @dev This function updates the context in the state before and after the
    /// message call.
    function _invokeMessageCall(
        Message calldata _message,
        bytes32 _msgHash,
        uint256 _gasLimit
    )
        private
        returns (bool success_)
    {
        if (_gasLimit == 0) revert B_INVALID_GAS_LIMIT();
        assert(_message.from != address(this));

        _storeContext(_msgHash, _message.from, _message.srcChainId);

        if (
            _message.data.length >= 4 // msg can be empty
                && bytes4(_message.data) != IMessageInvocable.onMessageInvocation.selector
                && _message.to.isContract()
        ) {
            success_ = false;
        } else {
            success_ = _message.to.sendEther(_message.value, _gasLimit, _message.data);
        }

        // Must reset the context after the message call
        _resetContext();
    }

    /// @notice Updates the status of a bridge message.
    /// @dev If the new status is different from the current status in the
    /// mapping, the status is updated and an event is emitted.
    /// @param _msgHash The hash of the message.
    /// @param _status The new status of the message.
    function _updateMessageStatus(bytes32 _msgHash, Status _status) private {
        if (messageStatus[_msgHash] == _status) return;

        messageStatus[_msgHash] = _status;
        emit MessageStatusChanged(_msgHash, _status);

        if (_status == Status.FAILED) {
            ISignalService(resolve("signal_service", false)).sendSignal(
                signalForFailedMessage(_msgHash)
            );
        }
    }

    /// @notice Resets the call context
    function _resetContext() private {
        if (LibNetwork.isDencunSupported(block.chainid)) {
            _storeContext(bytes32(0), address(0), uint64(0));
        } else {
            _storeContext(bytes32(PLACEHOLDER), address(uint160(PLACEHOLDER)), uint64(PLACEHOLDER));
        }
    }

    /// @notice Stores the call context
    /// @param _msgHash The message hash.
    /// @param _from The sender's address.
    /// @param _srcChainId The source chain ID.
    function _storeContext(bytes32 _msgHash, address _from, uint64 _srcChainId) private {
        if (LibNetwork.isDencunSupported(block.chainid)) {
            assembly {
                tstore(_CTX_SLOT, _msgHash)
                tstore(add(_CTX_SLOT, 1), _from)
                tstore(add(_CTX_SLOT, 2), _srcChainId)
            }
        } else {
            __ctx = Context(_msgHash, _from, _srcChainId);
        }
    }

    /// @notice Loads and returns the call context.
    /// @return ctx_ The call context.
    function _loadContext() private view returns (Context memory) {
        if (LibNetwork.isDencunSupported(block.chainid)) {
            bytes32 msgHash;
            address from;
            uint64 srcChainId;
            assembly {
                msgHash := tload(_CTX_SLOT)
                from := tload(add(_CTX_SLOT, 1))
                srcChainId := tload(add(_CTX_SLOT, 2))
            }
            return Context(msgHash, from, srcChainId);
        } else {
            return __ctx;
        }
    }

    /// @notice Checks if the signal was received and caches cross-chain data if requested.
    /// @param _signalService The signal service address.
    /// @param _signal The signal.
    /// @param _chainId The ID of the chain the signal is stored on.
    /// @param _proof The merkle inclusion proof.
    /// @return true if the message was received.
    function _proveSignalReceived(
        address _signalService,
        bytes32 _signal,
        uint64 _chainId,
        bytes calldata _proof
    )
        private
        returns (bool)
    {
        try ISignalService(_signalService).proveSignalReceived(
            _chainId, resolve(_chainId, "bridge", false), _signal, _proof
        ) {
            return true;
        } catch {
            return false;
        }
    }

    /// @notice Checks if the signal was received.
    /// This is the 'readonly' version of _proveSignalReceived.
    /// @param _signalService The signal service address.
    /// @param _signal The signal.
    /// @param _chainId The ID of the chain the signal is stored on.
    /// @param _proof The merkle inclusion proof.
    /// @return true if the message was received.
    function _isSignalReceived(
        address _signalService,
        bytes32 _signal,
        uint64 _chainId,
        bytes calldata _proof
    )
        private
        view
        returns (bool)
    {
        try ISignalService(_signalService).verifySignalReceived(
            _chainId, resolve(_chainId, "bridge", false), _signal, _proof
        ) {
            return true;
        } catch {
            return false;
        }
    }

    function _isPostInvocationDelay(
        uint256 _receivedAt,
        uint256 _invocationDelay
    )
        private
        view
        returns (bool)
    {
        unchecked {
            return block.timestamp >= _receivedAt.max(lastUnpausedAt) + _invocationDelay;
        }
    }
}
