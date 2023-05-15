// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.18;

import {AddressResolver} from "../../common/AddressResolver.sol";
import {BlockHeader} from "../../libs/LibBlockHeader.sol";
import {IBridge} from "../IBridge.sol";
import {LibAddress} from "../../libs/LibAddress.sol";
import {LibMath} from "../../libs/LibMath.sol";

/**
 * Stores message metadata on the Bridge.
 */
library LibBridgeData {
    struct State {
        uint256 nextMessageId;
        IBridge.Context ctx; // 3 slots
        mapping(bytes32 msgHash => bool released) etherReleased;
        uint256[45] __gap;
    }

    struct StatusProof {
        BlockHeader header;
        bytes proof;
    }

    bytes32 internal constant MESSAGE_HASH_PLACEHOLDER = bytes32(uint256(1));
    uint256 internal constant CHAINID_PLACEHOLDER = type(uint256).max;
    address internal constant SRC_CHAIN_SENDER_PLACEHOLDER = address(uint160(uint256(1)));

    // Note: These events must match the ones defined in Bridge.sol.
    event MessageSent(bytes32 indexed msgHash, IBridge.Message message);
    event DestChainEnabled(uint256 indexed chainId, bool enabled);

    /**
     * @return msgHash The keccak256 hash of the message.
     */
    function hashMessage(IBridge.Message memory message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }
}
