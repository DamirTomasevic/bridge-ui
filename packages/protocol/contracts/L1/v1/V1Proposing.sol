// SPDX-License-Identifier: MIT
//
// ╭━━━━╮╱╱╭╮╱╱╱╱╱╭╮╱╱╱╱╱╭╮
// ┃╭╮╭╮┃╱╱┃┃╱╱╱╱╱┃┃╱╱╱╱╱┃┃
// ╰╯┃┃┣┻━┳┫┃╭┳━━╮┃┃╱╱╭━━┫╰━┳━━╮
// ╱╱┃┃┃╭╮┣┫╰╯┫╭╮┃┃┃╱╭┫╭╮┃╭╮┃━━┫
// ╱╱┃┃┃╭╮┃┃╭╮┫╰╯┃┃╰━╯┃╭╮┃╰╯┣━━┃
// ╱╱╰╯╰╯╰┻┻╯╰┻━━╯╰━━━┻╯╰┻━━┻━━╯
pragma solidity ^0.8.9;

import "../../common/ConfigManager.sol";
import "../../libs/LibConstants.sol";
import "../../libs/LibTxDecoder.sol";
import "../TkoToken.sol";
import "./V1Utils.sol";

/// @author dantaik <dan@taiko.xyz>
library V1Proposing {
    using LibTxDecoder for bytes;
    using SafeCastUpgradeable for uint256;
    using LibData for LibData.State;

    event BlockCommitted(
        uint64 commitSlot,
        uint64 commitHeight,
        bytes32 commitHash
    );
    event BlockProposed(uint256 indexed id, LibData.BlockMetadata meta);

    modifier onlyWhitelistedProposer(LibData.TentativeState storage tentative) {
        if (tentative.whitelistProposers) {
            require(tentative.proposers[msg.sender], "L1:whitelist");
        }
        _;
    }

    function commitBlock(
        LibData.State storage state,
        uint64 commitSlot,
        bytes32 commitHash
    ) public {
        assert(LibConstants.K_COMMIT_DELAY_CONFIRMS > 0);
        // It's OK to allow committing block when the system is halt.
        // By not checking the halt status, this method will be cheaper.
        //
        // assert(!V1Utils.isHalted(state));

        bytes32 hash = _aggregateCommitHash(block.number, commitHash);

        require(state.commits[msg.sender][commitSlot] != hash, "L1:committed");
        state.commits[msg.sender][commitSlot] = hash;

        emit BlockCommitted({
            commitSlot: commitSlot,
            commitHeight: uint64(block.number),
            commitHash: commitHash
        });
    }

    function proposeBlock(
        LibData.State storage state,
        LibData.TentativeState storage tentative,
        AddressResolver resolver,
        bytes[] calldata inputs
    ) public onlyWhitelistedProposer(tentative) {
        assert(!V1Utils.isHalted(state));

        require(inputs.length == 2, "L1:inputs:size");
        LibData.BlockMetadata memory meta = abi.decode(
            inputs[0],
            (LibData.BlockMetadata)
        );
        _verifyBlockCommit(state, meta);
        _validateMetadata(meta);

        {
            bytes calldata txList = inputs[1];
            // perform validation and populate some fields
            require(
                txList.length > 0 &&
                    txList.length <= LibConstants.K_TXLIST_MAX_BYTES &&
                    meta.txListHash == txList.hashTxList(),
                "L1:txList"
            );
            require(
                state.nextBlockId <
                    state.latestVerifiedId + LibConstants.K_MAX_NUM_BLOCKS,
                "L1:tooMany"
            );

            meta.id = state.nextBlockId;
            meta.l1Height = block.number - 1;
            meta.l1Hash = blockhash(block.number - 1);
            meta.timestamp = uint64(block.timestamp);

            // if multiple L2 blocks included in the same L1 block,
            // their block.mixHash fields for randomness will be the same.
            meta.mixHash = bytes32(block.difficulty);
        }

        state.saveProposedBlock(
            state.nextBlockId,
            LibData.ProposedBlock({
                metaHash: LibData.hashMetadata(meta),
                proposer: msg.sender,
                proposedAt: meta.timestamp
            })
        );

        state.lastProposedAt = meta.timestamp;
        emit BlockProposed(state.nextBlockId++, meta);
    }

    function isCommitValid(
        LibData.State storage state,
        uint256 commitSlot,
        uint256 commitHeight,
        bytes32 commitHash
    ) public view returns (bool) {
        assert(LibConstants.K_COMMIT_DELAY_CONFIRMS > 0);
        bytes32 hash = _aggregateCommitHash(commitHeight, commitHash);
        return
            state.commits[msg.sender][commitSlot] == hash &&
            block.number >= commitHeight + LibConstants.K_COMMIT_DELAY_CONFIRMS;
    }

    function _verifyBlockCommit(
        LibData.State storage state,
        LibData.BlockMetadata memory meta
    ) private {
        if (LibConstants.K_COMMIT_DELAY_CONFIRMS == 0) {
            return;
        }
        bytes32 commitHash = _calculateCommitHash(
            meta.beneficiary,
            meta.txListHash
        );

        require(
            isCommitValid({
                state: state,
                commitSlot: meta.commitSlot,
                commitHeight: meta.commitHeight,
                commitHash: commitHash
            }),
            "L1:notCommitted"
        );

        if (meta.commitSlot == 0) {
            // Special handling of slot 0 for refund; non-zero slots
            // are supposed to managed by node software for reuse.
            delete state.commits[msg.sender][meta.commitSlot];
        }
    }

    function _validateMetadata(LibData.BlockMetadata memory meta) private pure {
        require(
            meta.id == 0 &&
                meta.l1Height == 0 &&
                meta.l1Hash == 0 &&
                meta.mixHash == 0 &&
                meta.timestamp == 0 &&
                meta.beneficiary != address(0) &&
                meta.txListHash != 0,
            "L1:placeholder"
        );

        require(
            meta.gasLimit <= LibConstants.K_BLOCK_MAX_GAS_LIMIT,
            "L1:gasLimit"
        );
        require(meta.extraData.length <= 32, "L1:extraData");
    }

    function _calculateCommitHash(
        address beneficiary,
        bytes32 txListHash
    ) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(beneficiary, txListHash));
    }

    function _aggregateCommitHash(
        uint256 commitHeight,
        bytes32 commitHash
    ) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(commitHash, commitHeight));
    }
}
