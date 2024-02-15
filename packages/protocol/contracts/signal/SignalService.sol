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

import "@openzeppelin/contracts/utils/math/SafeCast.sol";
import "../common/EssentialContract.sol";
import "../libs/LibTrieProof.sol";
import "./ISignalService.sol";
import "./LibSignals.sol";

/// @title SignalService
/// @dev Labeled in AddressResolver as "signal_service"
/// @notice See the documentation in {ISignalService} for more details.
contract SignalService is EssentialContract, ISignalService {
    enum CacheOption {
        CACHE_NOTHING,
        CACHE_SIGNAL_ROOT,
        CACHE_STATE_ROOT,
        CACHE_BOTH
    }

    struct HopProof {
        uint64 chainId;
        CacheOption cacheOption;
        bytes32 rootHash;
        bytes[] accountProof;
        bytes[] storageProof;
    }

    uint256[50] private __gap;

    event SnippetRelayed(
        uint64 indexed chainid, bytes32 indexed kind, bytes32 data, bytes32 signal
    );

    error SS_EMPTY_PROOF();
    error SS_INVALID_APP();
    error SS_INVALID_LAST_HOP_CHAINID();
    error SS_INVALID_MID_HOP_CHAINID();
    error SS_INVALID_PARAMS();
    error SS_INVALID_SIGNAL();
    error SS_LOCAL_CHAIN_DATA_NOT_FOUND();
    error SS_UNSUPPORTED();

    /// @dev Initializer to be called after being deployed behind a proxy.
    function init(address _addressManager) external initializer {
        __Essential_init(_addressManager);
    }

    /// @inheritdoc ISignalService
    function relayChainData(
        uint64 chainId,
        bytes32 kind,
        bytes32 data
    )
        external
        onlyFromNamed("taiko")
        returns (bytes32 slot)
    {
        return _relayChainData(chainId, kind, data);
    }

    /// @inheritdoc ISignalService
    function sendSignal(bytes32 signal) public returns (bytes32 slot) {
        return _sendSignal(msg.sender, signal);
    }

    /// @inheritdoc ISignalService
    /// @dev This function may revert.
    function proveSignalReceived(
        uint64 chainId,
        address app,
        bytes32 signal,
        bytes calldata proof
    )
        public
        virtual
    {
        if (app == address(0) || signal == 0) revert SS_INVALID_PARAMS();

        HopProof[] memory _hopProofs = abi.decode(proof, (HopProof[]));
        if (_hopProofs.length == 0) revert SS_EMPTY_PROOF();

        uint64 _chainId = chainId;
        address _app = app;
        bytes32 _signal = signal;
        address _signalService = resolve(_chainId, "signal_service", false);

        for (uint256 i; i < _hopProofs.length; ++i) {
            HopProof memory hop = _hopProofs[i];

            bytes32 signalRoot = _verifyHopProof(_chainId, _app, _signal, hop, _signalService);

            bool isLastHop = i == _hopProofs.length - 1;
            if (isLastHop) {
                if (hop.chainId != block.chainid) revert SS_INVALID_LAST_HOP_CHAINID();
                _signalService = address(this);
            } else {
                if (hop.chainId == 0 || hop.chainId == block.chainid) {
                    revert SS_INVALID_MID_HOP_CHAINID();
                }
                _signalService = resolve(hop.chainId, "signal_service", false);
            }

            bool isFullProof = hop.accountProof.length > 0;

            _cacheChainData(hop, _chainId, signalRoot, isFullProof, isLastHop);

            bytes32 kind = isFullProof ? LibSignals.STATE_ROOT : LibSignals.SIGNAL_ROOT;
            _signal = signalForChainData(_chainId, kind, hop.rootHash);
            _chainId = hop.chainId;
            _app = _signalService;
        }

        if (!isSignalSent(address(this), _signal)) revert SS_LOCAL_CHAIN_DATA_NOT_FOUND();
    }

    /// @inheritdoc ISignalService
    function isChainDataRelayed(
        uint64 chainId,
        bytes32 kind,
        bytes32 data
    )
        public
        view
        returns (bool)
    {
        return isSignalSent(address(this), signalForChainData(chainId, kind, data));
    }

    /// @inheritdoc ISignalService
    function isSignalSent(address app, bytes32 signal) public view returns (bool) {
        if (signal == 0) revert SS_INVALID_SIGNAL();
        if (app == address(0)) revert SS_INVALID_APP();
        bytes32 slot = getSignalSlot(uint64(block.chainid), app, signal);
        uint256 value;
        assembly {
            value := sload(slot)
        }
        return value == 1;
    }

    /// @notice Get the storage slot of the signal.
    /// @param chainId The address's chainId.
    /// @param app The address that initiated the signal.
    /// @param signal The signal to get the storage slot of.
    /// @return The unique storage slot of the signal which is
    /// created by encoding the sender address with the signal (message).
    function getSignalSlot(
        uint64 chainId,
        address app,
        bytes32 signal
    )
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked("SIGNAL", chainId, app, signal));
    }

    function signalForChainData(
        uint64 chainId,
        bytes32 kind,
        bytes32 data
    )
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(chainId, kind, data));
    }

    function _relayChainData(
        uint64 chainId,
        bytes32 kind,
        bytes32 data
    )
        internal
        returns (bytes32 slot)
    {
        bytes32 signal = signalForChainData(chainId, kind, data);
        emit SnippetRelayed(chainId, kind, data, signal);
        return _sendSignal(address(this), signal);
    }

    function _sendSignal(address sender, bytes32 signal) internal returns (bytes32 slot) {
        if (signal == 0) revert SS_INVALID_SIGNAL();
        slot = getSignalSlot(uint64(block.chainid), sender, signal);
        assembly {
            sstore(slot, 1)
        }
    }

    function _verifyHopProof(
        uint64 chainId,
        address app,
        bytes32 signal,
        HopProof memory hop,
        address relay
    )
        internal
        virtual
        returns (bytes32 signalRoot)
    {
        return LibTrieProof.verifyMerkleProof(
            hop.rootHash,
            relay,
            getSignalSlot(chainId, app, signal),
            hex"01",
            hop.accountProof,
            hop.storageProof
        );
    }

    function _authorizePause(address) internal pure override {
        revert SS_UNSUPPORTED();
    }

    function _cacheChainData(
        HopProof memory hop,
        uint64 chainId,
        bytes32 signalRoot,
        bool isFullProof,
        bool isLastHop
    )
        private
    {
        // cache state root
        bool cacheStateRoot = hop.cacheOption == CacheOption.CACHE_BOTH
            || hop.cacheOption == CacheOption.CACHE_STATE_ROOT;

        if (cacheStateRoot && isFullProof && !isLastHop) {
            _relayChainData(chainId, LibSignals.STATE_ROOT, hop.rootHash);
        }

        // cache signal root
        bool cacheSignalRoot = hop.cacheOption == CacheOption.CACHE_BOTH
            || hop.cacheOption == CacheOption.CACHE_SIGNAL_ROOT;

        if (cacheSignalRoot && (!isLastHop || isFullProof)) {
            _relayChainData(chainId, LibSignals.SIGNAL_ROOT, signalRoot);
        }
    }
}
