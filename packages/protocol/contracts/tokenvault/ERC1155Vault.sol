// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/utils/ERC1155ReceiverUpgradeable.sol";
import "../libs/LibAddress.sol";
import "./BaseNFTVault.sol";
import "./BridgedERC1155.sol";

/// @title IERC1155NameAndSymbol
/// @notice Interface for ERC1155 contracts that provide name() and symbol()
/// functions. These functions may not be part of the official interface but are
/// used by some contracts.
/// @custom:security-contact security@taiko.xyz
interface IERC1155NameAndSymbol {
    /// @notice Returns the name of the token.
    function name() external view returns (string memory);

    /// @notice Returns the symbol of the token.
    function symbol() external view returns (string memory);
}

/// @title ERC1155Vault
/// @dev Labeled in AddressResolver as "erc1155_vault"
/// @notice This vault holds all ERC1155 tokens that users have deposited.
/// It also manages the mapping between canonical tokens and their bridged
/// tokens.
contract ERC1155Vault is BaseNFTVault, ERC1155ReceiverUpgradeable {
    using LibAddress for address;

    uint256[50] private __gap;

    /// @notice Initializes the contract.
    /// @param _owner The owner of this contract. msg.sender will be used if this value is zero.
    /// @param _addressManager The address of the {AddressManager} contract.
    function init(address _owner, address _addressManager) external initializer {
        __Essential_init(_owner, _addressManager);
        __ERC1155Receiver_init_unchained();
    }
    /// @notice Transfers ERC1155 tokens to this vault and sends a message to
    /// the destination chain so the user can receive the same (bridged) tokens
    /// by invoking the message call.
    /// @param _op Option for sending the ERC1155 token.
    /// @return message_ The constructed message.

    function sendToken(BridgeTransferOp memory _op)
        external
        payable
        whenNotPaused
        withValidOperation(_op)
        nonReentrant
        returns (IBridge.Message memory message_)
    {
        for (uint256 i; i < _op.amounts.length; ++i) {
            if (_op.amounts[i] == 0) revert VAULT_INVALID_AMOUNT();
        }
        // Check token interface support
        if (!_op.token.supportsInterface(ERC1155_INTERFACE_ID)) {
            revert VAULT_INTERFACE_NOT_SUPPORTED();
        }

        (bytes memory data, CanonicalNFT memory ctoken) = _handleMessage(msg.sender, _op);

        // Create a message to send to the destination chain
        IBridge.Message memory message = IBridge.Message({
            id: 0, // will receive a new value
            from: address(0), // will receive a new value
            srcChainId: 0, // will receive a new value
            destChainId: _op.destChainId,
            srcOwner: msg.sender,
            destOwner: _op.destOwner != address(0) ? _op.destOwner : msg.sender,
            to: resolve(_op.destChainId, name(), false),
            refundTo: _op.refundTo,
            value: msg.value - _op.fee,
            fee: _op.fee,
            gasLimit: _op.gasLimit,
            data: data,
            memo: _op.memo
        });

        // Send the message and obtain the message hash
        bytes32 msgHash;
        (msgHash, message_) =
            IBridge(resolve("bridge", false)).sendMessage{ value: msg.value }(message);

        // Emit TokenSent event
        emit TokenSent({
            msgHash: msgHash,
            from: message_.srcOwner,
            to: _op.to,
            destChainId: message_.destChainId,
            ctoken: ctoken.addr,
            token: _op.token,
            tokenIds: _op.tokenIds,
            amounts: _op.amounts
        });
    }

    /// @inheritdoc IMessageInvocable
    function onMessageInvocation(bytes calldata data) external payable whenNotPaused nonReentrant {
        (
            CanonicalNFT memory ctoken,
            address from,
            address to,
            uint256[] memory tokenIds,
            uint256[] memory amounts
        ) = abi.decode(data, (CanonicalNFT, address, address, uint256[], uint256[]));

        // Check context validity
        // `onlyFromBridge` checked in checkProcessMessageContext
        IBridge.Context memory ctx = checkProcessMessageContext();

        // Don't allow sending to disallowed addresses.
        // Don't send the tokens back to `from` because `from` is on the source chain.
        if (to == address(0) || to == address(this)) revert VAULT_INVALID_TO();

        // Transfer the ETH and the tokens to the `to` address
        address token = _transferTokens(ctoken, to, tokenIds, amounts);
        to.sendEtherAndVerify(msg.value);

        emit TokenReceived({
            msgHash: ctx.msgHash,
            from: from,
            to: to,
            srcChainId: ctx.srcChainId,
            ctoken: ctoken.addr,
            token: token,
            tokenIds: tokenIds,
            amounts: amounts
        });
    }

    /// @inheritdoc IRecallableSender
    function onMessageRecalled(
        IBridge.Message calldata message,
        bytes32 msgHash
    )
        external
        payable
        override
        whenNotPaused
        nonReentrant
    {
        // `onlyFromBridge` checked in checkRecallMessageContext
        checkRecallMessageContext();

        (bytes memory data) = abi.decode(message.data[4:], (bytes));
        (CanonicalNFT memory ctoken,,, uint256[] memory tokenIds, uint256[] memory amounts) =
            abi.decode(data, (CanonicalNFT, address, address, uint256[], uint256[]));

        // Transfer the ETH and tokens back to the owner
        address token = _transferTokens(ctoken, message.srcOwner, tokenIds, amounts);
        message.srcOwner.sendEtherAndVerify(message.value);

        // Emit TokenReleased event
        emit TokenReleased({
            msgHash: msgHash,
            from: message.srcOwner,
            ctoken: ctoken.addr,
            token: token,
            tokenIds: tokenIds,
            amounts: amounts
        });
    }

    /// @notice See {ERC1155ReceiverUpgradeable-onERC1155BatchReceived}.
    function onERC1155BatchReceived(
        address,
        address,
        uint256[] calldata,
        uint256[] calldata,
        bytes calldata
    )
        external
        pure
        returns (bytes4)
    {
        return IERC1155ReceiverUpgradeable.onERC1155BatchReceived.selector;
    }

    /// @notice See {ERC1155ReceiverUpgradeable-onERC1155Received}.
    function onERC1155Received(
        address,
        address,
        uint256,
        uint256,
        bytes calldata
    )
        external
        pure
        returns (bytes4)
    {
        return IERC1155ReceiverUpgradeable.onERC1155Received.selector;
    }

    /// @dev See {BaseVault-supportsInterface}.
    /// @param interfaceId The interface identifier.
    /// @return true if supports, else otherwise.
    function supportsInterface(bytes4 interfaceId)
        public
        pure
        override(BaseVault, ERC1155ReceiverUpgradeable)
        returns (bool)
    {
        return interfaceId == type(ERC1155ReceiverUpgradeable).interfaceId
            || BaseVault.supportsInterface(interfaceId);
    }

    /// @inheritdoc BaseVault
    function name() public pure override returns (bytes32) {
        return "erc1155_vault";
    }

    /// @dev Transfers ERC1155 tokens to the `to` address.
    /// @param ctoken CanonicalNFT data.
    /// @param to The address to transfer the tokens to.
    /// @param tokenIds The token IDs to transfer.
    /// @param amounts The amounts to transfer.
    /// @return token The address of the token.
    function _transferTokens(
        CanonicalNFT memory ctoken,
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts
    )
        private
        returns (address token)
    {
        if (ctoken.chainId == block.chainid) {
            // Token lives on this chain
            token = ctoken.addr;
            IERC1155(token).safeBatchTransferFrom(address(this), to, tokenIds, amounts, "");
        } else {
            // Token does not live on this chain
            token = _getOrDeployBridgedToken(ctoken);
            BridgedERC1155(token).mintBatch(to, tokenIds, amounts);
        }
    }

    /// @dev Handles the message on the source chain and returns the encoded
    /// call on the destination call.
    /// @param _user The user's address.
    /// @param _op BridgeTransferOp data.
    /// @return msgData_ Encoded message data.
    /// @return ctoken_ The canonical token.
    function _handleMessage(
        address _user,
        BridgeTransferOp memory _op
    )
        private
        returns (bytes memory msgData_, CanonicalNFT memory ctoken_)
    {
        unchecked {
            // is a btoken, meaning, it does not live on this chain
            if (bridgedToCanonical[_op.token].addr != address(0)) {
                ctoken_ = bridgedToCanonical[_op.token];
                for (uint256 i; i < _op.tokenIds.length; ++i) {
                    BridgedERC1155(_op.token).burn(_user, _op.tokenIds[i], _op.amounts[i]);
                }
            } else {
                // is a ctoken token, meaning, it lives on this chain
                ctoken_ = CanonicalNFT({
                    chainId: uint64(block.chainid),
                    addr: _op.token,
                    symbol: "",
                    name: ""
                });
                IERC1155NameAndSymbol t = IERC1155NameAndSymbol(_op.token);
                try t.name() returns (string memory _name) {
                    ctoken_.name = _name;
                } catch { }
                try t.symbol() returns (string memory _symbol) {
                    ctoken_.symbol = _symbol;
                } catch { }
                for (uint256 i; i < _op.tokenIds.length; ++i) {
                    IERC1155(_op.token).safeTransferFrom({
                        from: msg.sender,
                        to: address(this),
                        id: _op.tokenIds[i],
                        amount: _op.amounts[i],
                        data: ""
                    });
                }
            }
        }
        msgData_ = abi.encodeCall(
            this.onMessageInvocation, abi.encode(ctoken_, _user, _op.to, _op.tokenIds, _op.amounts)
        );
    }

    /// @dev Retrieve or deploy a bridged ERC1155 token contract.
    /// @param _ctoken CanonicalNFT data.
    /// @return btoken_ Address of the bridged token contract.
    function _getOrDeployBridgedToken(CanonicalNFT memory _ctoken)
        private
        returns (address btoken_)
    {
        btoken_ = canonicalToBridged[_ctoken.chainId][_ctoken.addr];
        if (btoken_ == address(0)) {
            btoken_ = _deployBridgedToken(_ctoken);
        }
    }

    /// @dev Deploy a new BridgedNFT contract and initialize it.
    /// This must be called before the first time a bridged token is sent to
    /// this chain.
    /// @param _ctoken CanonicalNFT data.
    /// @return btoken_ Address of the deployed bridged token contract.
    function _deployBridgedToken(CanonicalNFT memory _ctoken) private returns (address btoken_) {
        bytes memory data = abi.encodeCall(
            BridgedERC1155.init,
            (owner(), addressManager, _ctoken.addr, _ctoken.chainId, _ctoken.symbol, _ctoken.name)
        );

        btoken_ = address(new ERC1967Proxy(resolve("bridged_erc1155", false), data));

        bridgedToCanonical[btoken_] = _ctoken;
        canonicalToBridged[_ctoken.chainId][_ctoken.addr] = btoken_;

        emit BridgedTokenDeployed({
            chainId: _ctoken.chainId,
            ctoken: _ctoken.addr,
            btoken: btoken_,
            ctokenSymbol: _ctoken.symbol,
            ctokenName: _ctoken.name
        });
    }
}
