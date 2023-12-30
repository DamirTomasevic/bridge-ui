// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/
//
//   Email: security@taiko.xyz
//   Website: https://taiko.xyz
//   GitHub: https://github.com/taikoxyz
//   Discord: https://discord.gg/taikoxyz
//   Twitter: https://twitter.com/taikoxyz
//   Blog: https://mirror.xyz/labs.taiko.eth
//   Youtube: https://www.youtube.com/@taikoxyz

pragma solidity 0.8.20;

import "lib/openzeppelin-contracts/contracts/utils/Strings.sol";
import "lib/openzeppelin-contracts-upgradeable/contracts/token/ERC1155/ERC1155Upgradeable.sol";
import
    "lib/openzeppelin-contracts-upgradeable/contracts/token/ERC1155/extensions/IERC1155MetadataURIUpgradeable.sol";
import "lib/openzeppelin-contracts-upgradeable/contracts/token/ERC1155/IERC1155Upgradeable.sol";
import "../common/EssentialContract.sol";
import "./LibBridgedToken.sol";

/// @title BridgedERC1155
/// @notice Contract for bridging ERC1155 tokens across different chains.
contract BridgedERC1155 is
    EssentialContract,
    IERC1155Upgradeable,
    IERC1155MetadataURIUpgradeable,
    ERC1155Upgradeable
{
    address public srcToken; // Address of the source token contract.
    uint256 public srcChainId; // Source chain ID where the token originates.
    string private symbol_; // Symbol of the bridged token.
    string private name_; // Name of the bridged token.

    uint256[46] private __gap;

    error BTOKEN_CANNOT_RECEIVE();
    error BTOKEN_INVALID_PARAMS();

    /// @dev Initializer function to be called after deployment.
    /// @param _addressManager The address of the address manager.
    /// @param _srcToken Address of the source token.
    /// @param _srcChainId Source chain ID.
    /// @param _symbol Symbol of the bridged token.
    /// @param _name Name of the bridged token.
    function init(
        address _addressManager,
        address _srcToken,
        uint256 _srcChainId,
        string memory _symbol,
        string memory _name
    )
        external
        initializer
    {
        if (_srcToken == address(0) || _srcChainId == 0 || _srcChainId == block.chainid) {
            revert BTOKEN_INVALID_PARAMS();
        }
        __Essential_init(_addressManager);
        __ERC1155_init("");
        srcToken = _srcToken;
        srcChainId = _srcChainId;
        symbol_ = _symbol;
        name_ = _name;
    }

    /// @dev Mints tokens.
    /// @param account Address to receive the minted tokens.
    /// @param tokenId ID of the token to mint.
    /// @param amount Amount of tokens to mint.
    function mint(
        address account,
        uint256 tokenId,
        uint256 amount
    )
        public
        nonReentrant
        whenNotPaused
        onlyFromNamed("erc1155_vault")
    {
        _mint(account, tokenId, amount, "");
    }

    /// @dev Burns tokens.
    /// @param account Address from which tokens are burned.
    /// @param tokenId ID of the token to burn.
    /// @param amount Amount of tokens to burn.
    function burn(
        address account,
        uint256 tokenId,
        uint256 amount
    )
        public
        nonReentrant
        whenNotPaused
        onlyFromNamed("erc1155_vault")
    {
        _burn(account, tokenId, amount);
    }

    /// @notice Gets the name of the bridged token.
    /// @return The name.
    function name() public view returns (string memory) {
        return LibBridgedToken.buildName(name_, srcChainId);
    }

    /// @notice Gets the symbol of the bridged token.
    /// @return The symbol.
    function symbol() public view returns (string memory) {
        return LibBridgedToken.buildSymbol(symbol_);
    }

    function _beforeTokenTransfer(
        address, /*operator*/
        address, /*from*/
        address to,
        uint256[] memory, /*ids*/
        uint256[] memory, /*amounts*/
        bytes memory /*data*/
    )
        internal
        virtual
        override
    {
        if (to == address(this)) revert BTOKEN_CANNOT_RECEIVE();
        if (paused()) revert INVALID_PAUSE_STATUS();
    }
}
