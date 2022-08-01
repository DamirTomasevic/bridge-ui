// SPDX-License-Identifier: MIT
//
// ╭━━━━╮╱╱╭╮╱╱╱╱╱╭╮╱╱╱╱╱╭╮
// ┃╭╮╭╮┃╱╱┃┃╱╱╱╱╱┃┃╱╱╱╱╱┃┃
// ╰╯┃┃┣┻━┳┫┃╭┳━━╮┃┃╱╱╭━━┫╰━┳━━╮
// ╱╱┃┃┃╭╮┣┫╰╯┫╭╮┃┃┃╱╭┫╭╮┃╭╮┃━━┫
// ╱╱┃┃┃╭╮┃┃╭╮┫╰╯┃┃╰━╯┃╭╮┃╰╯┣━━┃
// ╱╱╰╯╰╯╰┻┻╯╰┻━━╯╰━━━┻╯╰┻━━┻━━╯
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";

import "../libs/LibStorageProof.sol";
import "../libs/LibMerkleProof.sol";
import "../libs/LibTxList.sol";
import "../libs/LibConstants.sol";
import "./LibBlockHeader.sol";
import "./LibZKP.sol";
import "./KeyManager.sol";

struct BlockContext {
    uint256 id;
    uint256 anchorHeight;
    bytes32 anchorHash;
    address beneficiary;
    uint64 gasLimit;
    bytes extraData;
    bytes32 txListHash;
    bytes32 mixHash;
    uint64 timestamp;
}

struct Snippet {
    bytes32 blockHash;
    bytes32 stateRoot;
}

struct ProofRecord {
    address prover;
    Snippet snippet;
}

/// @dev We have the following design assumptions:
/// - Assumption 1: the `difficulty` and `nonce` fields in Taiko block header
//                  will always be zeros, and this will be checked by zkEVM.
///
/// - Assumption 2: Taiko L2 allows block.timestamp >= parent.timestamp.
///
/// - Assumption 3: mixHash will be used by Taiko L2 for randomness, see:
///                 https://blog.ethereum.org/2021/11/29/how-the-merge-impacts-app-layer
///
/// - Assumption 4: Taiko zkEVM will check `sum(tx_i.gasLimit) <= header.gasLimit`
///                 and `header.gasLimit <= MAX_TAIKO_BLOCK_GAS_LIMIT`
///
/// This contract shall be deployed as the initial implementation of a
/// https://docs.openzeppelin.com/contracts/4.x/api/proxy#UpgradeableBeacon contract,
/// then a https://docs.openzeppelin.com/contracts/4.x/api/proxy#BeaconProxy contract
/// shall be deployed infront of it.
contract TaikoL1 is ReentrancyGuardUpgradeable {
    using LibBlockHeader for BlockHeader;
    using LibTxList for bytes;
    /**********************
     * Constants   *
     **********************/
    uint256 public constant MAX_ANCHOR_HEIGHT_DIFF = 128;
    uint256 public constant MAX_PENDING_BLOCKS = 1024;
    uint256 public constant MAX_THROW_AWAY_PARENT_DIFF = 1024;
    uint256 public constant MAX_FINALIZATION_WRITES_PER_TX = 5;
    uint256 public constant MAX_FINALIZATION_READS_PER_TX = 50;
    string public constant ZKP_VKEY = "TAIKO_ZKP_VKEY";
    bytes32 private constant JUMP_MARKER = bytes32(uint256(1));

    /**********************
     * State Variables    *
     **********************/

    // Finalized taiko block headers
    mapping(uint256 => Snippet) public finalizedBlocks;

    // block id => block context hash
    mapping(uint256 => bytes32) public pendingBlocks;

    mapping(uint256 => mapping(bytes32 => ProofRecord)) public proofRecords;

    address public taikoL2Address;
    uint64 public genesisHeight;
    uint64 public lastFinalizedHeight;
    uint64 public lastFinalizedId;
    uint64 public nextPendingId;
    KeyManager public keyManager;

    uint256[44] private __gap;

    /**********************
     * Events             *
     **********************/

    event BlockProposed(uint256 indexed id, BlockContext context);
    event BlockProven(
        uint256 indexed id,
        bytes32 parentHash,
        ProofRecord record
    );
    event BlockProvenInvalid(uint256 indexed id);
    event BlockFinalized(uint256 indexed id, Snippet snippet);

    /**********************
     * Modifiers          *
     **********************/

    modifier whenBlockIsPending(BlockContext calldata context) {
        _checkContextPending(context);
        _;
        finalizeBlocks();
    }

    /**********************
     * External Functions *
     **********************/

    function init(Snippet calldata genesis, address keyManagerAddr)
        external
        initializer
    {
        ReentrancyGuardUpgradeable.__ReentrancyGuard_init();

        require(
            AddressUpgradeable.isContract(keyManagerAddr),
            "invalid keyManager"
        );

        finalizedBlocks[0] = genesis;
        nextPendingId = 1;

        genesisHeight = uint64(block.number);
        keyManager = KeyManager(keyManagerAddr);

        emit BlockFinalized(0, genesis);
    }

    /// @notice Propose a Taiko L2 block.
    /// @param context The context that the actual L2 block header must satisfy.
    ///        Note the following fields in the provided context object must
    ///        be zeros, and their actual values will be provisioned by Ethereum.
    ///        - txListHash
    ///        - mixHash
    ///        - timestamp
    /// @param txList A list of transactions in this block, encoded with RLP.
    ///
    function proposeBlock(BlockContext memory context, bytes calldata txList)
        external
        nonReentrant
    {
        // Try to finalize blocks first to make room
        finalizeBlocks();

        require(txList.length > 0, "empty txList");
        require(
            nextPendingId <= lastFinalizedId + MAX_PENDING_BLOCKS,
            "too many pending blocks"
        );
        validateContext(context);

        context.id = nextPendingId;
        context.timestamp = uint64(block.timestamp);
        context.txListHash = txList.hashTxList();

        // if multiple L2 blocks included in the same L1 block,
        // their block.mixHash fields for randomness will be the same.
        context.mixHash = bytes32(block.difficulty);

        _savePendingBlock(nextPendingId, _hashContext(context));
        emit BlockProposed(nextPendingId, context);

        nextPendingId += 1;
    }

    function proveBlock(
        bool anchored,
        BlockHeader calldata header,
        BlockContext calldata context,
        bytes[2] calldata proofs
    ) external nonReentrant whenBlockIsPending(context) {
        _validateHeaderForContext(header, context);
        bytes32 blockHash = header.hashBlockHeader();

        LibZKP.verify(
            keyManager.getKey(ZKP_VKEY),
            header.parentHash,
            blockHash,
            context.txListHash,
            proofs[0]
        );

        (bytes32 proofKey, bytes32 proofVal) = LibStorageProof
            .computeAnchorProofKV(
                header.height,
                context.anchorHeight,
                context.anchorHash
            );

        if (!anchored) {
            proofVal = 0x0;
        }

        LibMerkleProof.verify(
            header.stateRoot,
            taikoL2Address,
            proofKey,
            proofVal,
            proofs[1]
        );

        ProofRecord memory record = ProofRecord({
            prover: msg.sender,
            snippet: Snippet({
                blockHash: blockHash,
                stateRoot: header.stateRoot
            })
        });

        proofRecords[context.id][header.parentHash] = record;

        emit BlockProven(context.id, header.parentHash, record);
    }

    function proveBlockInvalid(
        bytes32 throwAwayTxListHash, // hash of a txList that contains a verifyBlockInvalid tx on L2.
        BlockHeader calldata throwAwayHeader,
        BlockContext calldata context,
        bytes[2] calldata proofs
    ) external nonReentrant whenBlockIsPending(context) {
        require(
            throwAwayHeader.isPartiallyValidForTaiko(),
            "throwAwayHeader invalid"
        );

        require(
            lastFinalizedHeight <=
                throwAwayHeader.height + MAX_THROW_AWAY_PARENT_DIFF,
            "parent too old"
        );
        require(
            throwAwayHeader.parentHash ==
                finalizedBlocks[throwAwayHeader.height - 1].blockHash,
            "parent mismatch"
        );

        LibZKP.verify(
            keyManager.getKey(ZKP_VKEY),
            throwAwayHeader.parentHash,
            throwAwayHeader.hashBlockHeader(),
            throwAwayTxListHash,
            proofs[0]
        );

        (bytes32 key, bytes32 value) = LibStorageProof
            .computeInvalidTxListProofKV(context.txListHash);

        LibMerkleProof.verify(
            throwAwayHeader.stateRoot,
            taikoL2Address,
            key,
            value,
            proofs[1]
        );

        _invalidateBlock(context.id);
    }

    function verifyBlockInvalid(
        BlockContext calldata context,
        bytes calldata txList
    ) external nonReentrant whenBlockIsPending(context) {
        require(txList.hashTxList() == context.txListHash, "txList mismatch");
        require(!LibTxListValidator.isTxListValid(txList), "txList decoded");

        _invalidateBlock(context.id);
    }

    /**********************
     * Public Functions   *
     **********************/

    function finalizeBlocks() public {
        Snippet memory parent = finalizedBlocks[lastFinalizedHeight];
        uint256 nextId = lastFinalizedId + 1;
        uint256 reads = 0;
        uint256 writes = 0;
        while (
            nextId < nextPendingId &&
            reads <= MAX_FINALIZATION_READS_PER_TX &&
            writes <= MAX_FINALIZATION_WRITES_PER_TX
        ) {
            Snippet storage snippet = proofRecords[nextId][parent.blockHash]
                .snippet;

            if (snippet.blockHash != 0x0) {
                delete proofRecords[nextId][parent.blockHash];
                lastFinalizedHeight += 1;

                finalizedBlocks[lastFinalizedHeight] = snippet;
                emit BlockFinalized(lastFinalizedHeight, snippet);

                parent = snippet;
                writes += 1;
            } else if (
                proofRecords[nextId][JUMP_MARKER].snippet.blockHash ==
                JUMP_MARKER
            ) {
                delete proofRecords[nextId][JUMP_MARKER];
            } else {
                break;
            }

            lastFinalizedId += 1;
            nextId += 1;
            reads += 1;
        }
    }

    function validateContext(BlockContext memory context) public view {
        require(
            context.id == 0 &&
                context.txListHash == 0x0 &&
                context.mixHash == 0x0 &&
                context.timestamp == 0,
            "nonzero placeholder fields"
        );

        require(
            block.number <= context.anchorHeight + MAX_ANCHOR_HEIGHT_DIFF &&
                context.anchorHash == blockhash(context.anchorHeight) &&
                context.anchorHash != 0x0,
            "invalid anchor"
        );

        require(context.beneficiary != address(0), "null beneficiary");
        require(
            context.gasLimit <= LibConstants.MAX_TAIKO_BLOCK_GAS_LIMIT,
            "invalid gasLimit"
        );
        require(context.extraData.length <= 32, "extraData too large");
    }

    /**********************
     * Private Functions  *
     **********************/

    function _invalidateBlock(uint256 id) private {
        require(
            proofRecords[id][JUMP_MARKER].snippet.blockHash == 0x0,
            "already invalidated"
        );
        proofRecords[id][JUMP_MARKER] = ProofRecord({
            prover: msg.sender,
            snippet: Snippet({blockHash: JUMP_MARKER, stateRoot: JUMP_MARKER})
        });
        emit BlockProvenInvalid(id);
    }

    function _savePendingBlock(uint256 id, bytes32 contextHash)
        private
        returns (bytes32)
    {
        return pendingBlocks[id % MAX_PENDING_BLOCKS] = contextHash;
    }

    function _getPendingBlock(uint256 id) private view returns (bytes32) {
        return pendingBlocks[id % MAX_PENDING_BLOCKS];
    }

    function _checkContextPending(BlockContext calldata context) private view {
        require(
            context.id > lastFinalizedId && context.id < nextPendingId,
            "invalid id"
        );
        require(
            _getPendingBlock(context.id) == _hashContext(context),
            "context mismatch"
        );
    }

    function _validateHeader(BlockHeader calldata header) private pure {
        require(
            header.parentHash != 0x0 &&
                header.gasLimit <= LibConstants.MAX_TAIKO_BLOCK_GAS_LIMIT &&
                header.extraData.length <= 32 &&
                header.difficulty == 0 &&
                header.nonce == 0,
            "header mismatch"
        );
    }

    function _validateHeaderForContext(
        BlockHeader calldata header,
        BlockContext memory context
    ) private pure {
        require(
            header.beneficiary == context.beneficiary &&
                header.gasLimit == context.gasLimit &&
                header.timestamp == context.timestamp &&
                keccak256(header.extraData) == keccak256(context.extraData) && // TODO: direct compare
                header.mixHash == context.mixHash,
            "header mismatch"
        );
    }

    function _hashContext(BlockContext memory context)
        private
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(context));
    }
}
