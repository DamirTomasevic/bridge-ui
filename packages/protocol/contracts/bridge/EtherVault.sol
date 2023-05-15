// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.18;

import {SafeERC20Upgradeable} from
    "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";
import {Create2Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/Create2Upgradeable.sol";
import {EssentialContract} from "../common/EssentialContract.sol";
import {Proxied} from "../common/Proxied.sol";
import {LibAddress} from "../libs/LibAddress.sol";
import {BridgeErrors} from "./BridgeErrors.sol";

/**
 * @custom:security-contact hello@taiko.xyz
 * EtherVault is a special vault contract that:
 * - Is initialized with 2^128 Ether.
 * - Allows the contract owner to authorize addresses.
 * - Allows authorized addresses to send/release Ether.
 */
contract EtherVault is EssentialContract, BridgeErrors {
    using LibAddress for address;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/

    mapping(address addr => bool isAuthorized) private _authorizedAddrs;
    uint256[49] private __gap;

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/

    event Authorized(address indexed addr, bool authorized);

    event EtherReleased(address indexed to, uint256 amount);

    /*//////////////////////////////////////////////////////////////
                               MODIFIERS
    //////////////////////////////////////////////////////////////*/

    modifier onlyAuthorized() {
        if (!isAuthorized(msg.sender)) {
            revert B_EV_NOT_AUTHORIZED();
        }
        _;
    }

    /*//////////////////////////////////////////////////////////////
                         USER-FACING FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    receive() external payable {
        // EthVault's balance must == 0 OR the sender isAuthorized.
        if (address(this).balance != 0 && !isAuthorized(msg.sender)) {
            revert B_EV_NOT_AUTHORIZED();
        }
    }

    function init(address addressManager) external initializer {
        EssentialContract._init(addressManager);
    }

    /**
     * Transfer Ether from EtherVault to the sender, checking that the sender
     * is authorized.
     * @param amount Amount of Ether to send.
     */
    function releaseEther(uint256 amount) public onlyAuthorized nonReentrant {
        msg.sender.sendEther(amount);
        emit EtherReleased(msg.sender, amount);
    }

    /**
     * Transfer Ether from EtherVault to a designated address, checking that the
     * sender is authorized.
     * @param recipient Address to receive Ether.
     * @param amount Amount of ether to send.
     */
    function releaseEther(address recipient, uint256 amount) public onlyAuthorized nonReentrant {
        if (recipient == address(0)) {
            revert B_EV_DO_NOT_BURN();
        }
        recipient.sendEther(amount);
        emit EtherReleased(recipient, amount);
    }

    /**
     * Set the authorized status of an address, only the owner can call this.
     * @param addr Address to set the authorized status of.
     * @param authorized Authorized status to set.
     */
    function authorize(address addr, bool authorized) public onlyOwner {
        if (addr == address(0) || _authorizedAddrs[addr] == authorized) {
            revert B_EV_PARAM();
        }
        _authorizedAddrs[addr] = authorized;
        emit Authorized(addr, authorized);
    }

    /**
     * Get the authorized status of an address.
     * @param addr Address to get the authorized status of.
     */
    function isAuthorized(address addr) public view returns (bool) {
        return _authorizedAddrs[addr];
    }
}

contract ProxiedEtherVault is Proxied, EtherVault {}
