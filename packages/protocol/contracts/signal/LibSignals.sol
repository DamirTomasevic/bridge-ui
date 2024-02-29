// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

/// @title LibSignals
/// @custom:security-contact security@taiko.xyz
library LibSignals {
    bytes32 public constant STATE_ROOT = keccak256("STATE_ROOT");
    bytes32 public constant SIGNAL_ROOT = keccak256("SIGNAL_ROOT");
}
