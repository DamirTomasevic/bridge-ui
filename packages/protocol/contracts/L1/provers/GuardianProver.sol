// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../tiers/ITierProvider.sol";
import "../ITaikoL1.sol";
import "./Guardians.sol";
import "../../common/LibStrings.sol";

/// @title GuardianProver
/// @custom:security-contact security@taiko.xyz
contract GuardianProver is Guardians {
    uint256[50] private __gap;

    /// @notice Emitted when a guardian proof is approved.
    /// @param addr The address of the guardian.
    /// @param blockId The block ID.
    /// @param blockHash The block hash.
    /// @param approved If the proof is approved.
    /// @param proofData The proof data.
    event GuardianApproval(
        address indexed addr,
        uint256 indexed blockId,
        bytes32 indexed blockHash,
        bool approved,
        bytes proofData
    );

    /// @notice Initializes the contract.
    /// @param _owner The owner of this contract. msg.sender will be used if this value is zero.
    /// @param _addressManager The address of the {AddressManager} contract.
    function init(address _owner, address _addressManager) external initializer {
        __Essential_init(_owner, _addressManager);
    }

    /// @dev Called by guardians to approve a guardian proof
    /// @param _meta The block's metadata.
    /// @param _tran The valid transition.
    /// @param _proof The tier proof.
    /// @return approved_ True if the minimum number of approval is acquired, false otherwise.
    function approve(
        TaikoData.BlockMetadata calldata _meta,
        TaikoData.Transition calldata _tran,
        TaikoData.TierProof calldata _proof
    )
        external
        whenNotPaused
        nonReentrant
        returns (bool approved_)
    {
        if (_proof.tier != LibTiers.TIER_GUARDIAN) {
            revert INVALID_PROOF();
        }

        bytes32 hash = keccak256(abi.encode(_meta, _tran, _proof.data));
        approved_ = approve(_meta.id, hash);

        emit GuardianApproval(msg.sender, _meta.id, _tran.blockHash, approved_, _proof.data);

        if (approved_) {
            deleteApproval(hash);
            ITaikoL1(resolve(LibStrings.B_TAIKO, false)).proveBlock(
                _meta.id, abi.encode(_meta, _tran, _proof)
            );
        }
    }
}
