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

pragma solidity 0.8.24;

import "../tiers/ITierProvider.sol";
import "../ITaikoL1.sol";
import "./Guardians.sol";

/// @title GuardianProver
contract GuardianProver is Guardians {
    uint256[50] private __gap;

    event GuardianApproval(
        address indexed addr, uint256 indexed blockId, bytes32 blockHash, bool approved
    );

    /// @notice Initializes the contract with the provided address manager.
    /// @param _addressManager The address of the address manager contract.
    function init(address _addressManager) external initializer {
        __Essential_init(_addressManager);
    }

    /// @dev Called by guardians to approve a guardian proof
    function approve(
        TaikoData.BlockMetadata calldata meta,
        TaikoData.Transition calldata tran,
        TaikoData.TierProof calldata proof
    )
        external
        whenNotPaused
        nonReentrant
        returns (bool approved)
    {
        if (proof.tier != LibTiers.TIER_GUARDIAN) revert INVALID_PROOF();
        bytes32 hash = keccak256(abi.encode(meta, tran));
        approved = approve(meta.id, hash);

        if (approved) {
            deleteApproval(hash);
            ITaikoL1(resolve("taiko", false)).proveBlock(meta.id, abi.encode(meta, tran, proof));
        }

        emit GuardianApproval(msg.sender, meta.id, tran.blockHash, approved);
    }
}
