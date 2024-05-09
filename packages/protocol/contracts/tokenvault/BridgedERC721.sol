// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "../common/EssentialContract.sol";
import "../common/LibStrings.sol";
import "./LibBridgedToken.sol";

/// @title BridgedERC721
/// @notice Contract for bridging ERC721 tokens across different chains.
/// @custom:security-contact security@taiko.xyz
contract BridgedERC721 is EssentialContract, ERC721Upgradeable {
    /// @notice Address of the source token contract.
    address public srcToken;

    /// @notice Source chain ID where the token originates.
    uint256 public srcChainId;

    uint256[48] private __gap;

    error BTOKEN_INVALID_PARAMS();
    error BTOKEN_INVALID_TO_ADDR();
    error BTOKEN_INVALID_BURN();

    /// @notice Initializes the contract.
    /// @param _owner The owner of this contract. msg.sender will be used if this value is zero.
    /// @param _addressManager The address of the {AddressManager} contract.
    /// @param _srcToken Address of the source token.
    /// @param _srcChainId Source chain ID.
    /// @param _symbol Symbol of the bridged token.
    /// @param _name Name of the bridged token.
    function init(
        address _owner,
        address _addressManager,
        address _srcToken,
        uint256 _srcChainId,
        string calldata _symbol,
        string calldata _name
    )
        external
        initializer
    {
        // Check if provided parameters are valid
        LibBridgedToken.validateInputs(_srcToken, _srcChainId);
        __Essential_init(_owner, _addressManager);
        __ERC721_init(_name, _symbol);

        srcToken = _srcToken;
        srcChainId = _srcChainId;
    }

    /// @dev Mints tokens.
    /// @param _account Address to receive the minted token.
    /// @param _tokenIds IDs of the tokens to mint.
    function batchMint(
        address _account,
        uint256[] memory _tokenIds
    )
        external
        whenNotPaused
        onlyFromNamed(LibStrings.B_ERC721_VAULT)
        nonReentrant
    {
        for (uint256 i; i < _tokenIds.length; ++i) {
            _safeMint(_account, _tokenIds[i]);
        }
    }

    /// @dev Burns tokens.
    /// @param _tokenIds IDs of the tokens to burn.
    function batchBurn(uint256[] memory _tokenIds)
        external
        whenNotPaused
        onlyFromNamed(LibStrings.B_ERC721_VAULT)
        nonReentrant
    {
        for (uint256 i; i < _tokenIds.length; ++i) {
            // Check if the caller is the owner of the token. Somehow this is not done inside the
            // _burn() function below.
            if (ownerOf(_tokenIds[i]) != msg.sender) revert BTOKEN_INVALID_BURN();
            _burn(_tokenIds[i]);
        }
    }

    function safeBatchTransferFrom(address _from, address _to, uint256[] memory _tokenIds) public {
        for (uint256 i; i < _tokenIds.length; ++i) {
            safeTransferFrom(_from, _to, _tokenIds[i], "");
        }
    }

    /// @notice Gets the source token and source chain ID being bridged.
    /// @return The source token's address.
    /// @return The source token's chain ID.
    function source() public view returns (address, uint256) {
        return (srcToken, srcChainId);
    }

    /// @notice Returns the token URI.
    /// @param _tokenId The token id.
    /// @return The token URI following EIP-681.
    function tokenURI(uint256 _tokenId) public view override returns (string memory) {
        // https://github.com/crytic/slither/wiki/Detector-Documentation#abi-encodePacked-collision
        // The abi.encodePacked() call below takes multiple dynamic arguments. This is known and
        // considered acceptable in terms of risk.
        return LibBridgedToken.buildURI(srcToken, srcChainId, Strings.toString(_tokenId));
    }

    /// @notice Gets the canonical token's address and chain ID.
    /// @return The canonical token's address.
    /// @return The canonical token's chain ID.
    function canonical() public view returns (address, uint256) {
        return (srcToken, srcChainId);
    }

    function _beforeTokenTransfer(
        address _from,
        address _to,
        uint256 _firstTokenId,
        uint256 _batchSize
    )
        internal
        override
        whenNotPaused
    {
        LibBridgedToken.checkToAddress(_to);
        super._beforeTokenTransfer(_from, _to, _firstTokenId, _batchSize);
    }
}
