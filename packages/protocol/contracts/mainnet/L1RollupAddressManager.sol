// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../common/AddressManager.sol";
import "../common/LibStrings.sol";

/// @title L1RollupAddressManager
/// @notice See the documentation in {IAddressManager}.
/// @dev This contract shall NOT be used to upgrade existing implementation unless the name-address
/// registration becomes stable in 0x579f40D0BE111b823962043702cabe6Aaa290780.
/// @custom:security-contact security@taiko.xyz
contract L1RollupAddressManager is AddressManager {
    /// @notice Gets the address mapped to a specific chainId-name pair.
    /// @dev Sub-contracts can override this method to avoid reading from storage.
    /// The following names are not cached:
    /// - B_PROPOSER
    /// - B_PROPOSER_ONE
    /// - B_TIER_PROVIDER
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
            if (_name == LibStrings.B_TAIKO) {
                return 0x06a9Ab27c7e2255df1815E6CC0168d7755Feb19a;
            }
            if (_name == LibStrings.B_TIER_SGX) {
                return 0xb0f3186FC1963f774f52ff455DC86aEdD0b31F81;
            }
            if (_name == LibStrings.B_TIER_GUARDIAN_MINORITY) {
                return 0x579A8d63a2Db646284CBFE31FE5082c9989E985c;
            }
            if (_name == LibStrings.B_TIER_GUARDIAN) {
                return 0xE3D777143Ea25A6E031d1e921F396750885f43aC;
            }
            if (_name == LibStrings.B_AUTOMATA_DCAP_ATTESTATION) {
                return 0x8d7C954960a36a7596d7eA4945dDf891967ca8A3;
            }
            if (_name == LibStrings.B_ASSIGNMENT_HOOK) {
                return 0x537a2f0D3a5879b41BCb5A2afE2EA5c4961796F6;
            }
        }
        return address(0);
    }
}
