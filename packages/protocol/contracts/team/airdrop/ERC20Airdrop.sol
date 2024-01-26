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

import "lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "./MerkleClaimable.sol";

/// @title ERC20Airdrop
/// Contract for managing Taiko token airdrop for eligible users
contract ERC20Airdrop is MerkleClaimable {
    address public token;
    address public vault;
    uint256[48] private __gap;

    function init(
        uint64 _claimStarts,
        uint64 _claimEnds,
        bytes32 _merkleRoot,
        address _token,
        address _vault
    )
        external
        initializer
    {
        __Essential_init();
        _setConfig(_claimStarts, _claimEnds, _merkleRoot);

        token = _token;
        vault = _vault;
    }

    function _claimWithData(bytes calldata data) internal override {
        (address user, uint256 amount) = abi.decode(data, (address, uint256));
        IERC20(token).transferFrom(vault, user, amount);
    }
}
