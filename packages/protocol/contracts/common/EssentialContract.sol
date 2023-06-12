// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import { IAddressManager } from "./AddressManager.sol";
import { OwnableUpgradeable } from
    "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from
    "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import { AddressResolver } from "./AddressResolver.sol";

/**
 * @dev This abstract contract serves as the base contract for many core
 *      components in this package.
 */
abstract contract EssentialContract is
    ReentrancyGuardUpgradeable,
    OwnableUpgradeable,
    AddressResolver
{
    /**
     * Sets a new AddressManager's address.
     *
     * @param newAddressManager New address manager contract address
     */
    function setAddressManager(address newAddressManager) external onlyOwner {
        if (newAddressManager == address(0)) revert RESOLVER_INVALID_ADDR();
        _addressManager = IAddressManager(newAddressManager);

        emit AddressManagerChanged(newAddressManager);
    }

    function _init(address _addressManager) internal virtual override {
        ReentrancyGuardUpgradeable.__ReentrancyGuard_init();
        OwnableUpgradeable.__Ownable_init();
        AddressResolver._init(_addressManager);
    }
}
