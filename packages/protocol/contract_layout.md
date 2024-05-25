## TaikoL1
| Name           | Type                   | Slot | Offset | Bytes | Contract                         |
|----------------|------------------------|------|--------|-------|----------------------------------|
| _initialized   | uint8                  | 0    | 0      | 1     | contracts/L1/TaikoL1.sol:TaikoL1 |
| _initializing  | bool                   | 0    | 1      | 1     | contracts/L1/TaikoL1.sol:TaikoL1 |
| __gap          | uint256[50]            | 1    | 0      | 1600  | contracts/L1/TaikoL1.sol:TaikoL1 |
| _owner         | address                | 51   | 0      | 20    | contracts/L1/TaikoL1.sol:TaikoL1 |
| __gap          | uint256[49]            | 52   | 0      | 1568  | contracts/L1/TaikoL1.sol:TaikoL1 |
| _pendingOwner  | address                | 101  | 0      | 20    | contracts/L1/TaikoL1.sol:TaikoL1 |
| __gap          | uint256[49]            | 102  | 0      | 1568  | contracts/L1/TaikoL1.sol:TaikoL1 |
| addressManager | address                | 151  | 0      | 20    | contracts/L1/TaikoL1.sol:TaikoL1 |
| __gap          | uint256[49]            | 152  | 0      | 1568  | contracts/L1/TaikoL1.sol:TaikoL1 |
| __reentry      | uint8                  | 201  | 0      | 1     | contracts/L1/TaikoL1.sol:TaikoL1 |
| __paused       | uint8                  | 201  | 1      | 1     | contracts/L1/TaikoL1.sol:TaikoL1 |
| lastUnpausedAt | uint64                 | 201  | 2      | 8     | contracts/L1/TaikoL1.sol:TaikoL1 |
| __gap          | uint256[49]            | 202  | 0      | 1568  | contracts/L1/TaikoL1.sol:TaikoL1 |
| state          | struct TaikoData.State | 251  | 0      | 1600  | contracts/L1/TaikoL1.sol:TaikoL1 |
| __gap          | uint256[50]            | 301  | 0      | 1600  | contracts/L1/TaikoL1.sol:TaikoL1 |

## TaikoL2
| Name            | Type                        | Slot | Offset | Bytes | Contract                         |
|-----------------|-----------------------------|------|--------|-------|----------------------------------|
| _initialized    | uint8                       | 0    | 0      | 1     | contracts/L2/TaikoL2.sol:TaikoL2 |
| _initializing   | bool                        | 0    | 1      | 1     | contracts/L2/TaikoL2.sol:TaikoL2 |
| __gap           | uint256[50]                 | 1    | 0      | 1600  | contracts/L2/TaikoL2.sol:TaikoL2 |
| _owner          | address                     | 51   | 0      | 20    | contracts/L2/TaikoL2.sol:TaikoL2 |
| __gap           | uint256[49]                 | 52   | 0      | 1568  | contracts/L2/TaikoL2.sol:TaikoL2 |
| _pendingOwner   | address                     | 101  | 0      | 20    | contracts/L2/TaikoL2.sol:TaikoL2 |
| __gap           | uint256[49]                 | 102  | 0      | 1568  | contracts/L2/TaikoL2.sol:TaikoL2 |
| addressManager  | address                     | 151  | 0      | 20    | contracts/L2/TaikoL2.sol:TaikoL2 |
| __gap           | uint256[49]                 | 152  | 0      | 1568  | contracts/L2/TaikoL2.sol:TaikoL2 |
| __reentry       | uint8                       | 201  | 0      | 1     | contracts/L2/TaikoL2.sol:TaikoL2 |
| __paused        | uint8                       | 201  | 1      | 1     | contracts/L2/TaikoL2.sol:TaikoL2 |
| lastUnpausedAt  | uint64                      | 201  | 2      | 8     | contracts/L2/TaikoL2.sol:TaikoL2 |
| __gap           | uint256[49]                 | 202  | 0      | 1568  | contracts/L2/TaikoL2.sol:TaikoL2 |
| l2Hashes        | mapping(uint256 => bytes32) | 251  | 0      | 32    | contracts/L2/TaikoL2.sol:TaikoL2 |
| publicInputHash | bytes32                     | 252  | 0      | 32    | contracts/L2/TaikoL2.sol:TaikoL2 |
| gasExcess       | uint64                      | 253  | 0      | 8     | contracts/L2/TaikoL2.sol:TaikoL2 |
| lastSyncedBlock | uint64                      | 253  | 8      | 8     | contracts/L2/TaikoL2.sol:TaikoL2 |
| __deprecated1   | uint64                      | 253  | 16     | 8     | contracts/L2/TaikoL2.sol:TaikoL2 |
| __deprecated2   | uint64                      | 253  | 24     | 8     | contracts/L2/TaikoL2.sol:TaikoL2 |
| l1ChainId       | uint64                      | 254  | 0      | 8     | contracts/L2/TaikoL2.sol:TaikoL2 |
| __gap           | uint256[46]                 | 255  | 0      | 1472  | contracts/L2/TaikoL2.sol:TaikoL2 |

## SignalService
| Name           | Type                                          | Slot | Offset | Bytes | Contract                                         |
|----------------|-----------------------------------------------|------|--------|-------|--------------------------------------------------|
| _initialized   | uint8                                         | 0    | 0      | 1     | contracts/signal/SignalService.sol:SignalService |
| _initializing  | bool                                          | 0    | 1      | 1     | contracts/signal/SignalService.sol:SignalService |
| __gap          | uint256[50]                                   | 1    | 0      | 1600  | contracts/signal/SignalService.sol:SignalService |
| _owner         | address                                       | 51   | 0      | 20    | contracts/signal/SignalService.sol:SignalService |
| __gap          | uint256[49]                                   | 52   | 0      | 1568  | contracts/signal/SignalService.sol:SignalService |
| _pendingOwner  | address                                       | 101  | 0      | 20    | contracts/signal/SignalService.sol:SignalService |
| __gap          | uint256[49]                                   | 102  | 0      | 1568  | contracts/signal/SignalService.sol:SignalService |
| addressManager | address                                       | 151  | 0      | 20    | contracts/signal/SignalService.sol:SignalService |
| __gap          | uint256[49]                                   | 152  | 0      | 1568  | contracts/signal/SignalService.sol:SignalService |
| __reentry      | uint8                                         | 201  | 0      | 1     | contracts/signal/SignalService.sol:SignalService |
| __paused       | uint8                                         | 201  | 1      | 1     | contracts/signal/SignalService.sol:SignalService |
| lastUnpausedAt | uint64                                        | 201  | 2      | 8     | contracts/signal/SignalService.sol:SignalService |
| __gap          | uint256[49]                                   | 202  | 0      | 1568  | contracts/signal/SignalService.sol:SignalService |
| topBlockId     | mapping(uint64 => mapping(bytes32 => uint64)) | 251  | 0      | 32    | contracts/signal/SignalService.sol:SignalService |
| isAuthorized   | mapping(address => bool)                      | 252  | 0      | 32    | contracts/signal/SignalService.sol:SignalService |
| __gap          | uint256[48]                                   | 253  | 0      | 1536  | contracts/signal/SignalService.sol:SignalService |

## Bridge
| Name           | Type                                    | Slot | Offset | Bytes | Contract                           |
|----------------|-----------------------------------------|------|--------|-------|------------------------------------|
| _initialized   | uint8                                   | 0    | 0      | 1     | contracts/bridge/Bridge.sol:Bridge |
| _initializing  | bool                                    | 0    | 1      | 1     | contracts/bridge/Bridge.sol:Bridge |
| __gap          | uint256[50]                             | 1    | 0      | 1600  | contracts/bridge/Bridge.sol:Bridge |
| _owner         | address                                 | 51   | 0      | 20    | contracts/bridge/Bridge.sol:Bridge |
| __gap          | uint256[49]                             | 52   | 0      | 1568  | contracts/bridge/Bridge.sol:Bridge |
| _pendingOwner  | address                                 | 101  | 0      | 20    | contracts/bridge/Bridge.sol:Bridge |
| __gap          | uint256[49]                             | 102  | 0      | 1568  | contracts/bridge/Bridge.sol:Bridge |
| addressManager | address                                 | 151  | 0      | 20    | contracts/bridge/Bridge.sol:Bridge |
| __gap          | uint256[49]                             | 152  | 0      | 1568  | contracts/bridge/Bridge.sol:Bridge |
| __reentry      | uint8                                   | 201  | 0      | 1     | contracts/bridge/Bridge.sol:Bridge |
| __paused       | uint8                                   | 201  | 1      | 1     | contracts/bridge/Bridge.sol:Bridge |
| lastUnpausedAt | uint64                                  | 201  | 2      | 8     | contracts/bridge/Bridge.sol:Bridge |
| __gap          | uint256[49]                             | 202  | 0      | 1568  | contracts/bridge/Bridge.sol:Bridge |
| __reserved1    | uint64                                  | 251  | 0      | 8     | contracts/bridge/Bridge.sol:Bridge |
| nextMessageId  | uint64                                  | 251  | 8      | 8     | contracts/bridge/Bridge.sol:Bridge |
| messageStatus  | mapping(bytes32 => enum IBridge.Status) | 252  | 0      | 32    | contracts/bridge/Bridge.sol:Bridge |
| __ctx          | struct IBridge.Context                  | 253  | 0      | 64    | contracts/bridge/Bridge.sol:Bridge |
| __reserved2    | uint256                                 | 255  | 0      | 32    | contracts/bridge/Bridge.sol:Bridge |
| __reserved3    | uint256                                 | 256  | 0      | 32    | contracts/bridge/Bridge.sol:Bridge |
| __gap          | uint256[44]                             | 257  | 0      | 1408  | contracts/bridge/Bridge.sol:Bridge |

## DelegateOwner
| Name           | Type        | Slot | Offset | Bytes | Contract                                     |
|----------------|-------------|------|--------|-------|----------------------------------------------|
| _initialized   | uint8       | 0    | 0      | 1     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| _initializing  | bool        | 0    | 1      | 1     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __gap          | uint256[50] | 1    | 0      | 1600  | contracts/L2/DelegateOwner.sol:DelegateOwner |
| _owner         | address     | 51   | 0      | 20    | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __gap          | uint256[49] | 52   | 0      | 1568  | contracts/L2/DelegateOwner.sol:DelegateOwner |
| _pendingOwner  | address     | 101  | 0      | 20    | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __gap          | uint256[49] | 102  | 0      | 1568  | contracts/L2/DelegateOwner.sol:DelegateOwner |
| addressManager | address     | 151  | 0      | 20    | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __gap          | uint256[49] | 152  | 0      | 1568  | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __reentry      | uint8       | 201  | 0      | 1     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __paused       | uint8       | 201  | 1      | 1     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| lastUnpausedAt | uint64      | 201  | 2      | 8     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __gap          | uint256[49] | 202  | 0      | 1568  | contracts/L2/DelegateOwner.sol:DelegateOwner |
| l1ChainId      | uint64      | 251  | 0      | 8     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| nextTxId       | uint64      | 251  | 8      | 8     | contracts/L2/DelegateOwner.sol:DelegateOwner |
| realOwner      | address     | 252  | 0      | 20    | contracts/L2/DelegateOwner.sol:DelegateOwner |
| __gap          | uint256[48] | 253  | 0      | 1536  | contracts/L2/DelegateOwner.sol:DelegateOwner |

## GuardianProver
| Name                    | Type                                            | Slot | Offset | Bytes | Contract                                               |
|-------------------------|-------------------------------------------------|------|--------|-------|--------------------------------------------------------|
| _initialized            | uint8                                           | 0    | 0      | 1     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| _initializing           | bool                                            | 0    | 1      | 1     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __gap                   | uint256[50]                                     | 1    | 0      | 1600  | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| _owner                  | address                                         | 51   | 0      | 20    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __gap                   | uint256[49]                                     | 52   | 0      | 1568  | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| _pendingOwner           | address                                         | 101  | 0      | 20    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __gap                   | uint256[49]                                     | 102  | 0      | 1568  | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| addressManager          | address                                         | 151  | 0      | 20    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __gap                   | uint256[49]                                     | 152  | 0      | 1568  | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __reentry               | uint8                                           | 201  | 0      | 1     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __paused                | uint8                                           | 201  | 1      | 1     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| lastUnpausedAt          | uint64                                          | 201  | 2      | 8     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __gap                   | uint256[49]                                     | 202  | 0      | 1568  | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| guardianIds             | mapping(address => uint256)                     | 251  | 0      | 32    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| approvals               | mapping(uint256 => mapping(bytes32 => uint256)) | 252  | 0      | 32    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| guardians               | address[]                                       | 253  | 0      | 32    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| version                 | uint32                                          | 254  | 0      | 4     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| minGuardians            | uint32                                          | 254  | 4      | 4     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| provingAutoPauseEnabled | bool                                            | 254  | 8      | 1     | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| latestProofHash         | mapping(uint256 => mapping(uint256 => bytes32)) | 255  | 0      | 32    | contracts/L1/provers/GuardianProver.sol:GuardianProver |
| __gap                   | uint256[45]                                     | 256  | 0      | 1440  | contracts/L1/provers/GuardianProver.sol:GuardianProver |

## TaikoToken
| Name                                                | Type                                                          | Slot | Offset | Bytes | Contract                                |
|-----------------------------------------------------|---------------------------------------------------------------|------|--------|-------|-----------------------------------------|
| _initialized                                        | uint8                                                         | 0    | 0      | 1     | contracts/tko/TaikoToken.sol:TaikoToken |
| _initializing                                       | bool                                                          | 0    | 1      | 1     | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[50]                                                   | 1    | 0      | 1600  | contracts/tko/TaikoToken.sol:TaikoToken |
| _owner                                              | address                                                       | 51   | 0      | 20    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[49]                                                   | 52   | 0      | 1568  | contracts/tko/TaikoToken.sol:TaikoToken |
| _pendingOwner                                       | address                                                       | 101  | 0      | 20    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[49]                                                   | 102  | 0      | 1568  | contracts/tko/TaikoToken.sol:TaikoToken |
| addressManager                                      | address                                                       | 151  | 0      | 20    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[49]                                                   | 152  | 0      | 1568  | contracts/tko/TaikoToken.sol:TaikoToken |
| __reentry                                           | uint8                                                         | 201  | 0      | 1     | contracts/tko/TaikoToken.sol:TaikoToken |
| __paused                                            | uint8                                                         | 201  | 1      | 1     | contracts/tko/TaikoToken.sol:TaikoToken |
| lastUnpausedAt                                      | uint64                                                        | 201  | 2      | 8     | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[49]                                                   | 202  | 0      | 1568  | contracts/tko/TaikoToken.sol:TaikoToken |
| __slots_previously_used_by_ERC20SnapshotUpgradeable | uint256[50]                                                   | 251  | 0      | 1600  | contracts/tko/TaikoToken.sol:TaikoToken |
| _balances                                           | mapping(address => uint256)                                   | 301  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _allowances                                         | mapping(address => mapping(address => uint256))               | 302  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _totalSupply                                        | uint256                                                       | 303  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _name                                               | string                                                        | 304  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _symbol                                             | string                                                        | 305  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[45]                                                   | 306  | 0      | 1440  | contracts/tko/TaikoToken.sol:TaikoToken |
| _hashedName                                         | bytes32                                                       | 351  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _hashedVersion                                      | bytes32                                                       | 352  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _name                                               | string                                                        | 353  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _version                                            | string                                                        | 354  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[48]                                                   | 355  | 0      | 1536  | contracts/tko/TaikoToken.sol:TaikoToken |
| _nonces                                             | mapping(address => struct CountersUpgradeable.Counter)        | 403  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _PERMIT_TYPEHASH_DEPRECATED_SLOT                    | bytes32                                                       | 404  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[49]                                                   | 405  | 0      | 1568  | contracts/tko/TaikoToken.sol:TaikoToken |
| _delegates                                          | mapping(address => address)                                   | 454  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _checkpoints                                        | mapping(address => struct ERC20VotesUpgradeable.Checkpoint[]) | 455  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| _totalSupplyCheckpoints                             | struct ERC20VotesUpgradeable.Checkpoint[]                     | 456  | 0      | 32    | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[47]                                                   | 457  | 0      | 1504  | contracts/tko/TaikoToken.sol:TaikoToken |
| __gap                                               | uint256[50]                                                   | 504  | 0      | 1600  | contracts/tko/TaikoToken.sol:TaikoToken |

## BridgedTaikoToken
| Name                                                | Type                                                          | Slot | Offset | Bytes | Contract                                              |
|-----------------------------------------------------|---------------------------------------------------------------|------|--------|-------|-------------------------------------------------------|
| _initialized                                        | uint8                                                         | 0    | 0      | 1     | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _initializing                                       | bool                                                          | 0    | 1      | 1     | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[50]                                                   | 1    | 0      | 1600  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _owner                                              | address                                                       | 51   | 0      | 20    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[49]                                                   | 52   | 0      | 1568  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _pendingOwner                                       | address                                                       | 101  | 0      | 20    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[49]                                                   | 102  | 0      | 1568  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| addressManager                                      | address                                                       | 151  | 0      | 20    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[49]                                                   | 152  | 0      | 1568  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __reentry                                           | uint8                                                         | 201  | 0      | 1     | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __paused                                            | uint8                                                         | 201  | 1      | 1     | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| lastUnpausedAt                                      | uint64                                                        | 201  | 2      | 8     | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[49]                                                   | 202  | 0      | 1568  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __slots_previously_used_by_ERC20SnapshotUpgradeable | uint256[50]                                                   | 251  | 0      | 1600  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _balances                                           | mapping(address => uint256)                                   | 301  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _allowances                                         | mapping(address => mapping(address => uint256))               | 302  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _totalSupply                                        | uint256                                                       | 303  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _name                                               | string                                                        | 304  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _symbol                                             | string                                                        | 305  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[45]                                                   | 306  | 0      | 1440  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _hashedName                                         | bytes32                                                       | 351  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _hashedVersion                                      | bytes32                                                       | 352  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _name                                               | string                                                        | 353  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _version                                            | string                                                        | 354  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[48]                                                   | 355  | 0      | 1536  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _nonces                                             | mapping(address => struct CountersUpgradeable.Counter)        | 403  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _PERMIT_TYPEHASH_DEPRECATED_SLOT                    | bytes32                                                       | 404  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[49]                                                   | 405  | 0      | 1568  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _delegates                                          | mapping(address => address)                                   | 454  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _checkpoints                                        | mapping(address => struct ERC20VotesUpgradeable.Checkpoint[]) | 455  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| _totalSupplyCheckpoints                             | struct ERC20VotesUpgradeable.Checkpoint[]                     | 456  | 0      | 32    | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[47]                                                   | 457  | 0      | 1504  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |
| __gap                                               | uint256[50]                                                   | 504  | 0      | 1600  | contracts/tko/BridgedTaikoToken.sol:BridgedTaikoToken |

## ERC20Vault
| Name               | Type                                                 | Slot | Offset | Bytes | Contract                                       |
|--------------------|------------------------------------------------------|------|--------|-------|------------------------------------------------|
| _initialized       | uint8                                                | 0    | 0      | 1     | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| _initializing      | bool                                                 | 0    | 1      | 1     | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[50]                                          | 1    | 0      | 1600  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| _owner             | address                                              | 51   | 0      | 20    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[49]                                          | 52   | 0      | 1568  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| _pendingOwner      | address                                              | 101  | 0      | 20    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[49]                                          | 102  | 0      | 1568  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| addressManager     | address                                              | 151  | 0      | 20    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[49]                                          | 152  | 0      | 1568  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __reentry          | uint8                                                | 201  | 0      | 1     | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __paused           | uint8                                                | 201  | 1      | 1     | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| lastUnpausedAt     | uint64                                               | 201  | 2      | 8     | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[49]                                          | 202  | 0      | 1568  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[50]                                          | 251  | 0      | 1600  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| bridgedToCanonical | mapping(address => struct ERC20Vault.CanonicalERC20) | 301  | 0      | 32    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| canonicalToBridged | mapping(uint256 => mapping(address => address))      | 302  | 0      | 32    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| btokenDenylist     | mapping(address => bool)                             | 303  | 0      | 32    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| lastMigrationStart | mapping(uint256 => mapping(address => uint256))      | 304  | 0      | 32    | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |
| __gap              | uint256[46]                                          | 305  | 0      | 1472  | contracts/tokenvault/ERC20Vault.sol:ERC20Vault |

## ERC721Vault
| Name               | Type                                                 | Slot | Offset | Bytes | Contract                                         |
|--------------------|------------------------------------------------------|------|--------|-------|--------------------------------------------------|
| _initialized       | uint8                                                | 0    | 0      | 1     | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| _initializing      | bool                                                 | 0    | 1      | 1     | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[50]                                          | 1    | 0      | 1600  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| _owner             | address                                              | 51   | 0      | 20    | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[49]                                          | 52   | 0      | 1568  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| _pendingOwner      | address                                              | 101  | 0      | 20    | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[49]                                          | 102  | 0      | 1568  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| addressManager     | address                                              | 151  | 0      | 20    | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[49]                                          | 152  | 0      | 1568  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __reentry          | uint8                                                | 201  | 0      | 1     | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __paused           | uint8                                                | 201  | 1      | 1     | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| lastUnpausedAt     | uint64                                               | 201  | 2      | 8     | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[49]                                          | 202  | 0      | 1568  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[50]                                          | 251  | 0      | 1600  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| bridgedToCanonical | mapping(address => struct BaseNFTVault.CanonicalNFT) | 301  | 0      | 32    | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| canonicalToBridged | mapping(uint256 => mapping(address => address))      | 302  | 0      | 32    | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[48]                                          | 303  | 0      | 1536  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |
| __gap              | uint256[50]                                          | 351  | 0      | 1600  | contracts/tokenvault/ERC721Vault.sol:ERC721Vault |

## ERC1155Vault
| Name               | Type                                                 | Slot | Offset | Bytes | Contract                                           |
|--------------------|------------------------------------------------------|------|--------|-------|----------------------------------------------------|
| _initialized       | uint8                                                | 0    | 0      | 1     | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| _initializing      | bool                                                 | 0    | 1      | 1     | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[50]                                          | 1    | 0      | 1600  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| _owner             | address                                              | 51   | 0      | 20    | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[49]                                          | 52   | 0      | 1568  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| _pendingOwner      | address                                              | 101  | 0      | 20    | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[49]                                          | 102  | 0      | 1568  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| addressManager     | address                                              | 151  | 0      | 20    | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[49]                                          | 152  | 0      | 1568  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __reentry          | uint8                                                | 201  | 0      | 1     | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __paused           | uint8                                                | 201  | 1      | 1     | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| lastUnpausedAt     | uint64                                               | 201  | 2      | 8     | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[49]                                          | 202  | 0      | 1568  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[50]                                          | 251  | 0      | 1600  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| bridgedToCanonical | mapping(address => struct BaseNFTVault.CanonicalNFT) | 301  | 0      | 32    | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| canonicalToBridged | mapping(uint256 => mapping(address => address))      | 302  | 0      | 32    | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[48]                                          | 303  | 0      | 1536  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[50]                                          | 351  | 0      | 1600  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[50]                                          | 401  | 0      | 1600  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |
| __gap              | uint256[50]                                          | 451  | 0      | 1600  | contracts/tokenvault/ERC1155Vault.sol:ERC1155Vault |

## BridgedERC20
| Name             | Type                                            | Slot | Offset | Bytes | Contract                                           |
|------------------|-------------------------------------------------|------|--------|-------|----------------------------------------------------|
| _initialized     | uint8                                           | 0    | 0      | 1     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _initializing    | bool                                            | 0    | 1      | 1     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[50]                                     | 1    | 0      | 1600  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _owner           | address                                         | 51   | 0      | 20    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[49]                                     | 52   | 0      | 1568  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _pendingOwner    | address                                         | 101  | 0      | 20    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[49]                                     | 102  | 0      | 1568  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| addressManager   | address                                         | 151  | 0      | 20    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[49]                                     | 152  | 0      | 1568  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __reentry        | uint8                                           | 201  | 0      | 1     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __paused         | uint8                                           | 201  | 1      | 1     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| lastUnpausedAt   | uint64                                          | 201  | 2      | 8     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[49]                                     | 202  | 0      | 1568  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _balances        | mapping(address => uint256)                     | 251  | 0      | 32    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _allowances      | mapping(address => mapping(address => uint256)) | 252  | 0      | 32    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _totalSupply     | uint256                                         | 253  | 0      | 32    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _name            | string                                          | 254  | 0      | 32    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| _symbol          | string                                          | 255  | 0      | 32    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[45]                                     | 256  | 0      | 1440  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| srcToken         | address                                         | 301  | 0      | 20    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __srcDecimals    | uint8                                           | 301  | 20     | 1     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| srcChainId       | uint256                                         | 302  | 0      | 32    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| migratingAddress | address                                         | 303  | 0      | 20    | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| migratingInbound | bool                                            | 303  | 20     | 1     | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |
| __gap            | uint256[47]                                     | 304  | 0      | 1504  | contracts/tokenvault/BridgedERC20.sol:BridgedERC20 |

## BridgedERC721
| Name               | Type                                         | Slot | Offset | Bytes | Contract                                             |
|--------------------|----------------------------------------------|------|--------|-------|------------------------------------------------------|
| _initialized       | uint8                                        | 0    | 0      | 1     | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _initializing      | bool                                         | 0    | 1      | 1     | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[50]                                  | 1    | 0      | 1600  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _owner             | address                                      | 51   | 0      | 20    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[49]                                  | 52   | 0      | 1568  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _pendingOwner      | address                                      | 101  | 0      | 20    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[49]                                  | 102  | 0      | 1568  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| addressManager     | address                                      | 151  | 0      | 20    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[49]                                  | 152  | 0      | 1568  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __reentry          | uint8                                        | 201  | 0      | 1     | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __paused           | uint8                                        | 201  | 1      | 1     | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| lastUnpausedAt     | uint64                                       | 201  | 2      | 8     | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[49]                                  | 202  | 0      | 1568  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[50]                                  | 251  | 0      | 1600  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _name              | string                                       | 301  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _symbol            | string                                       | 302  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _owners            | mapping(uint256 => address)                  | 303  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _balances          | mapping(address => uint256)                  | 304  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _tokenApprovals    | mapping(uint256 => address)                  | 305  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| _operatorApprovals | mapping(address => mapping(address => bool)) | 306  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[44]                                  | 307  | 0      | 1408  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| srcToken           | address                                      | 351  | 0      | 20    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| srcChainId         | uint256                                      | 352  | 0      | 32    | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |
| __gap              | uint256[48]                                  | 353  | 0      | 1536  | contracts/tokenvault/BridgedERC721.sol:BridgedERC721 |

## BridgedERC1155
| Name               | Type                                            | Slot | Offset | Bytes | Contract                                               |
|--------------------|-------------------------------------------------|------|--------|-------|--------------------------------------------------------|
| _initialized       | uint8                                           | 0    | 0      | 1     | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| _initializing      | bool                                            | 0    | 1      | 1     | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[50]                                     | 1    | 0      | 1600  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| _owner             | address                                         | 51   | 0      | 20    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[49]                                     | 52   | 0      | 1568  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| _pendingOwner      | address                                         | 101  | 0      | 20    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[49]                                     | 102  | 0      | 1568  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| addressManager     | address                                         | 151  | 0      | 20    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[49]                                     | 152  | 0      | 1568  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __reentry          | uint8                                           | 201  | 0      | 1     | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __paused           | uint8                                           | 201  | 1      | 1     | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| lastUnpausedAt     | uint64                                          | 201  | 2      | 8     | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[49]                                     | 202  | 0      | 1568  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[50]                                     | 251  | 0      | 1600  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| _balances          | mapping(uint256 => mapping(address => uint256)) | 301  | 0      | 32    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| _operatorApprovals | mapping(address => mapping(address => bool))    | 302  | 0      | 32    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| _uri               | string                                          | 303  | 0      | 32    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[47]                                     | 304  | 0      | 1504  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| srcToken           | address                                         | 351  | 0      | 20    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| srcChainId         | uint256                                         | 352  | 0      | 32    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| symbol             | string                                          | 353  | 0      | 32    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| name               | string                                          | 354  | 0      | 32    | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |
| __gap              | uint256[46]                                     | 355  | 0      | 1472  | contracts/tokenvault/BridgedERC1155.sol:BridgedERC1155 |

## AssignmentHook
| Name           | Type        | Slot | Offset | Bytes | Contract                                             |
|----------------|-------------|------|--------|-------|------------------------------------------------------|
| _initialized   | uint8       | 0    | 0      | 1     | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| _initializing  | bool        | 0    | 1      | 1     | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __gap          | uint256[50] | 1    | 0      | 1600  | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| _owner         | address     | 51   | 0      | 20    | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __gap          | uint256[49] | 52   | 0      | 1568  | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| _pendingOwner  | address     | 101  | 0      | 20    | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __gap          | uint256[49] | 102  | 0      | 1568  | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| addressManager | address     | 151  | 0      | 20    | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __gap          | uint256[49] | 152  | 0      | 1568  | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __reentry      | uint8       | 201  | 0      | 1     | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __paused       | uint8       | 201  | 1      | 1     | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| lastUnpausedAt | uint64      | 201  | 2      | 8     | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __gap          | uint256[49] | 202  | 0      | 1568  | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |
| __gap          | uint256[50] | 251  | 0      | 1600  | contracts/L1/hooks/AssignmentHook.sol:AssignmentHook |

## ERC20Airdrop
| Name           | Type                     | Slot | Offset | Bytes | Contract                                             |
|----------------|--------------------------|------|--------|-------|------------------------------------------------------|
| _initialized   | uint8                    | 0    | 0      | 1     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| _initializing  | bool                     | 0    | 1      | 1     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[50]              | 1    | 0      | 1600  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| _owner         | address                  | 51   | 0      | 20    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[49]              | 52   | 0      | 1568  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| _pendingOwner  | address                  | 101  | 0      | 20    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[49]              | 102  | 0      | 1568  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| addressManager | address                  | 151  | 0      | 20    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[49]              | 152  | 0      | 1568  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __reentry      | uint8                    | 201  | 0      | 1     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __paused       | uint8                    | 201  | 1      | 1     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| lastUnpausedAt | uint64                   | 201  | 2      | 8     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[49]              | 202  | 0      | 1568  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| isClaimed      | mapping(bytes32 => bool) | 251  | 0      | 32    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| merkleRoot     | bytes32                  | 252  | 0      | 32    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| claimStart     | uint64                   | 253  | 0      | 8     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| claimEnd       | uint64                   | 253  | 8      | 8     | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[47]              | 254  | 0      | 1504  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| token          | address                  | 301  | 0      | 20    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| vault          | address                  | 302  | 0      | 20    | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |
| __gap          | uint256[48]              | 303  | 0      | 1536  | contracts/team/airdrop/ERC20Airdrop.sol:ERC20Airdrop |

## AutomataDcapV3Attestation
| Name                    | Type                                            | Slot | Offset | Bytes | Contract                                                                               |
|-------------------------|-------------------------------------------------|------|--------|-------|----------------------------------------------------------------------------------------|
| _initialized            | uint8                                           | 0    | 0      | 1     | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| _initializing           | bool                                            | 0    | 1      | 1     | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __gap                   | uint256[50]                                     | 1    | 0      | 1600  | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| _owner                  | address                                         | 51   | 0      | 20    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __gap                   | uint256[49]                                     | 52   | 0      | 1568  | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| _pendingOwner           | address                                         | 101  | 0      | 20    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __gap                   | uint256[49]                                     | 102  | 0      | 1568  | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| addressManager          | address                                         | 151  | 0      | 20    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __gap                   | uint256[49]                                     | 152  | 0      | 1568  | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __reentry               | uint8                                           | 201  | 0      | 1     | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __paused                | uint8                                           | 201  | 1      | 1     | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| lastUnpausedAt          | uint64                                          | 201  | 2      | 8     | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __gap                   | uint256[49]                                     | 202  | 0      | 1568  | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| sigVerifyLib            | contract ISigVerifyLib                          | 251  | 0      | 20    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| pemCertLib              | contract IPEMCertChainLib                       | 252  | 0      | 20    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| checkLocalEnclaveReport | bool                                            | 252  | 20     | 1     | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| trustedUserMrEnclave    | mapping(bytes32 => bool)                        | 253  | 0      | 32    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| trustedUserMrSigner     | mapping(bytes32 => bool)                        | 254  | 0      | 32    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| serialNumIsRevoked      | mapping(uint256 => mapping(bytes => bool))      | 255  | 0      | 32    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| tcbInfo                 | mapping(string => struct TCBInfoStruct.TCBInfo) | 256  | 0      | 32    | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| qeIdentity              | struct EnclaveIdStruct.EnclaveId                | 257  | 0      | 128   | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |
| __gap                   | uint256[39]                                     | 261  | 0      | 1248  | contracts/automata-attestation/AutomataDcapV3Attestation.sol:AutomataDcapV3Attestation |

## SgxVerifier
| Name              | Type                                            | Slot | Offset | Bytes | Contract                                        |
|-------------------|-------------------------------------------------|------|--------|-------|-------------------------------------------------|
| _initialized      | uint8                                           | 0    | 0      | 1     | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| _initializing     | bool                                            | 0    | 1      | 1     | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __gap             | uint256[50]                                     | 1    | 0      | 1600  | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| _owner            | address                                         | 51   | 0      | 20    | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __gap             | uint256[49]                                     | 52   | 0      | 1568  | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| _pendingOwner     | address                                         | 101  | 0      | 20    | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __gap             | uint256[49]                                     | 102  | 0      | 1568  | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| addressManager    | address                                         | 151  | 0      | 20    | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __gap             | uint256[49]                                     | 152  | 0      | 1568  | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __reentry         | uint8                                           | 201  | 0      | 1     | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __paused          | uint8                                           | 201  | 1      | 1     | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| lastUnpausedAt    | uint64                                          | 201  | 2      | 8     | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __gap             | uint256[49]                                     | 202  | 0      | 1568  | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| nextInstanceId    | uint256                                         | 251  | 0      | 32    | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| instances         | mapping(uint256 => struct SgxVerifier.Instance) | 252  | 0      | 32    | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| addressRegistered | mapping(address => bool)                        | 253  | 0      | 32    | contracts/verifiers/SgxVerifier.sol:SgxVerifier |
| __gap             | uint256[47]                                     | 254  | 0      | 1504  | contracts/verifiers/SgxVerifier.sol:SgxVerifier |

## RiscZeroVerifier
| Name            | Type                              | Slot | Offset | Bytes | Contract                                                  |
|-----------------|-----------------------------------|------|--------|-------|-----------------------------------------------------------|
| _initialized    | uint8                             | 0    | 0      | 1     | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| _initializing   | bool                              | 0    | 1      | 1     | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __gap           | uint256[50]                       | 1    | 0      | 1600  | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| _owner          | address                           | 51   | 0      | 20    | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __gap           | uint256[49]                       | 52   | 0      | 1568  | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| _pendingOwner   | address                           | 101  | 0      | 20    | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __gap           | uint256[49]                       | 102  | 0      | 1568  | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| addressManager  | address                           | 151  | 0      | 20    | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __gap           | uint256[49]                       | 152  | 0      | 1568  | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __reentry       | uint8                             | 201  | 0      | 1     | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __paused        | uint8                             | 201  | 1      | 1     | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| lastUnpausedAt  | uint64                            | 201  | 2      | 8     | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __gap           | uint256[49]                       | 202  | 0      | 1568  | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| receiptVerifier | contract IRiscZeroReceiptVerifier | 251  | 0      | 20    | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| isImageTrusted  | mapping(bytes32 => bool)          | 252  | 0      | 32    | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |
| __gap           | uint256[48]                       | 253  | 0      | 1536  | contracts/verifiers/RiscZeroVerifier.sol:RiscZeroVerifier |

## QuotaManager
| Name           | Type                                          | Slot | Offset | Bytes | Contract                                       |
|----------------|-----------------------------------------------|------|--------|-------|------------------------------------------------|
| _initialized   | uint8                                         | 0    | 0      | 1     | contracts/bridge/QuotaManager.sol:QuotaManager |
| _initializing  | bool                                          | 0    | 1      | 1     | contracts/bridge/QuotaManager.sol:QuotaManager |
| __gap          | uint256[50]                                   | 1    | 0      | 1600  | contracts/bridge/QuotaManager.sol:QuotaManager |
| _owner         | address                                       | 51   | 0      | 20    | contracts/bridge/QuotaManager.sol:QuotaManager |
| __gap          | uint256[49]                                   | 52   | 0      | 1568  | contracts/bridge/QuotaManager.sol:QuotaManager |
| _pendingOwner  | address                                       | 101  | 0      | 20    | contracts/bridge/QuotaManager.sol:QuotaManager |
| __gap          | uint256[49]                                   | 102  | 0      | 1568  | contracts/bridge/QuotaManager.sol:QuotaManager |
| addressManager | address                                       | 151  | 0      | 20    | contracts/bridge/QuotaManager.sol:QuotaManager |
| __gap          | uint256[49]                                   | 152  | 0      | 1568  | contracts/bridge/QuotaManager.sol:QuotaManager |
| __reentry      | uint8                                         | 201  | 0      | 1     | contracts/bridge/QuotaManager.sol:QuotaManager |
| __paused       | uint8                                         | 201  | 1      | 1     | contracts/bridge/QuotaManager.sol:QuotaManager |
| lastUnpausedAt | uint64                                        | 201  | 2      | 8     | contracts/bridge/QuotaManager.sol:QuotaManager |
| __gap          | uint256[49]                                   | 202  | 0      | 1568  | contracts/bridge/QuotaManager.sol:QuotaManager |
| tokenQuota     | mapping(address => struct QuotaManager.Quota) | 251  | 0      | 32    | contracts/bridge/QuotaManager.sol:QuotaManager |
| quotaPeriod    | uint24                                        | 252  | 0      | 3     | contracts/bridge/QuotaManager.sol:QuotaManager |
| __gap          | uint256[48]                                   | 253  | 0      | 1536  | contracts/bridge/QuotaManager.sol:QuotaManager |

## ProverSet
| Name           | Type                     | Slot | Offset | Bytes | Contract                                       |
|----------------|--------------------------|------|--------|-------|------------------------------------------------|
| _initialized   | uint8                    | 0    | 0      | 1     | contracts/team/proving/ProverSet.sol:ProverSet |
| _initializing  | bool                     | 0    | 1      | 1     | contracts/team/proving/ProverSet.sol:ProverSet |
| __gap          | uint256[50]              | 1    | 0      | 1600  | contracts/team/proving/ProverSet.sol:ProverSet |
| _owner         | address                  | 51   | 0      | 20    | contracts/team/proving/ProverSet.sol:ProverSet |
| __gap          | uint256[49]              | 52   | 0      | 1568  | contracts/team/proving/ProverSet.sol:ProverSet |
| _pendingOwner  | address                  | 101  | 0      | 20    | contracts/team/proving/ProverSet.sol:ProverSet |
| __gap          | uint256[49]              | 102  | 0      | 1568  | contracts/team/proving/ProverSet.sol:ProverSet |
| addressManager | address                  | 151  | 0      | 20    | contracts/team/proving/ProverSet.sol:ProverSet |
| __gap          | uint256[49]              | 152  | 0      | 1568  | contracts/team/proving/ProverSet.sol:ProverSet |
| __reentry      | uint8                    | 201  | 0      | 1     | contracts/team/proving/ProverSet.sol:ProverSet |
| __paused       | uint8                    | 201  | 1      | 1     | contracts/team/proving/ProverSet.sol:ProverSet |
| lastUnpausedAt | uint64                   | 201  | 2      | 8     | contracts/team/proving/ProverSet.sol:ProverSet |
| __gap          | uint256[49]              | 202  | 0      | 1568  | contracts/team/proving/ProverSet.sol:ProverSet |
| isProver       | mapping(address => bool) | 251  | 0      | 32    | contracts/team/proving/ProverSet.sol:ProverSet |
| admin          | address                  | 252  | 0      | 20    | contracts/team/proving/ProverSet.sol:ProverSet |
| __gap          | uint256[48]              | 253  | 0      | 1536  | contracts/team/proving/ProverSet.sol:ProverSet |

## TokenUnlock
| Name           | Type                     | Slot | Offset | Bytes | Contract                                               |
|----------------|--------------------------|------|--------|-------|--------------------------------------------------------|
| _initialized   | uint8                    | 0    | 0      | 1     | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| _initializing  | bool                     | 0    | 1      | 1     | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __gap          | uint256[50]              | 1    | 0      | 1600  | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| _owner         | address                  | 51   | 0      | 20    | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __gap          | uint256[49]              | 52   | 0      | 1568  | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| _pendingOwner  | address                  | 101  | 0      | 20    | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __gap          | uint256[49]              | 102  | 0      | 1568  | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| addressManager | address                  | 151  | 0      | 20    | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __gap          | uint256[49]              | 152  | 0      | 1568  | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __reentry      | uint8                    | 201  | 0      | 1     | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __paused       | uint8                    | 201  | 1      | 1     | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| lastUnpausedAt | uint64                   | 201  | 2      | 8     | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __gap          | uint256[49]              | 202  | 0      | 1568  | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| amountVested   | uint256                  | 251  | 0      | 32    | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| recipient      | address                  | 252  | 0      | 20    | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| tgeTimestamp   | uint64                   | 252  | 20     | 8     | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| isProverSet    | mapping(address => bool) | 253  | 0      | 32    | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |
| __gap          | uint256[47]              | 254  | 0      | 1504  | contracts/team/tokenunlock/TokenUnlock.sol:TokenUnlock |

