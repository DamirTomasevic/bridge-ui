// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./TierProviderBase.sol";

/// @title DevnetTierProvider
/// @dev Labeled in AddressResolver as "tier_provider"
/// @custom:security-contact security@taiko.xyz
contract DevnetTierProvider is TierProviderBase {
    /// @inheritdoc ITierProvider
    function getTierIds() public pure override returns (uint16[] memory tiers_) {
        tiers_ = new uint16[](3);
        tiers_[0] = LibTiers.TIER_OPTIMISTIC;
        tiers_[1] = LibTiers.TIER_GUARDIAN_MINORITY;
        tiers_[2] = LibTiers.TIER_GUARDIAN;
    }

    /// @inheritdoc ITierProvider
    function getMinTier(uint256) public pure override returns (uint16) {
        return LibTiers.TIER_OPTIMISTIC;
    }
}
