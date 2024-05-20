// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../common/AddressManager.sol";
import "../common/LibStrings.sol";

/// @title L1SharedAddressManager
/// @notice See the documentation in {IAddressManager}.
/// @dev This contract shall NOT be used to upgrade existing implementation unless the name-address
/// registration becomes stable in 0xEf9EaA1dd30a9AA1df01c36411b5F082aA65fBaa.
/// @custom:security-contact security@taiko.xyz
contract L1SharedAddressManager is AddressManager {
    /// @notice Gets the address mapped to a specific chainId-name pair.
    /// @dev Sub-contracts can override this method to avoid reading from storage.
    /// The following names are not cached:
    /// - B_BRIDGE_WATCHDOG
    function _getOverride(
        uint64 _chainId,
        bytes32 _name
    )
        internal
        pure
        override
        returns (address addr_)
    {
        if (_chainId == 1) {
            if (_name == LibStrings.B_TAIKO_TOKEN) {
                return 0x10dea67478c5F8C5E2D90e5E9B26dBe60c54d800;
            }
            if (_name == LibStrings.B_SIGNAL_SERVICE) {
                return 0x9e0a24964e5397B566c1ed39258e21aB5E35C77C;
            }
            if (_name == LibStrings.B_BRIDGE) {
                return 0xd60247c6848B7Ca29eDdF63AA924E53dB6Ddd8EC;
            }
            if (_name == LibStrings.B_ERC20_VAULT) {
                return 0x996282cA11E5DEb6B5D122CC3B9A1FcAAD4415Ab;
            }
            if (_name == LibStrings.B_ERC721_VAULT) {
                return 0x0b470dd3A0e1C41228856Fb319649E7c08f419Aa;
            }
            if (_name == LibStrings.B_ERC1155_VAULT) {
                return 0xaf145913EA4a56BE22E120ED9C24589659881702;
            }
            if (_name == LibStrings.B_BRIDGED_ERC20) {
                return 0x79BC0Aada00fcF6E7AB514Bfeb093b5Fae3653e3;
            }
            if (_name == LibStrings.B_BRIDGED_ERC721) {
                return 0xC3310905E2BC9Cfb198695B75EF3e5B69C6A1Bf7;
            }
            if (_name == LibStrings.B_BRIDGED_ERC1155) {
                return 0x3c90963cFBa436400B0F9C46Aa9224cB379c2c40;
            }
            if (_name == LibStrings.B_QUOTA_MANAGER) {
                return 0x91f67118DD47d502B1f0C354D0611997B022f29E;
            }
        } else if (_chainId == 167_000) {
            if (_name == LibStrings.B_BRIDGE) {
                return 0x1670000000000000000000000000000000000001;
            }
            if (_name == LibStrings.B_ERC20_VAULT) {
                return 0x1670000000000000000000000000000000000002;
            }
            if (_name == LibStrings.B_ERC721_VAULT) {
                return 0x1670000000000000000000000000000000000003;
            }
            if (_name == LibStrings.B_ERC1155_VAULT) {
                return 0x1670000000000000000000000000000000000004;
            }
            if (_name == LibStrings.B_SIGNAL_SERVICE) {
                return 0x1670000000000000000000000000000000000005;
            }
        }

        return address(0);
    }
}
