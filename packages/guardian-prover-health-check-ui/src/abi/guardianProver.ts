export const guardianProverABI = [
  {
    inputs: [],
    name: "INVALID_GUARDIAN",
    type: "error",
  },
  {
    inputs: [],
    name: "INVALID_GUARDIAN_SET",
    type: "error",
  },
  {
    inputs: [],
    name: "INVALID_MIN_GUARDIANS",
    type: "error",
  },
  {
    inputs: [],
    name: "INVALID_PAUSE_STATUS",
    type: "error",
  },
  {
    inputs: [],
    name: "INVALID_PROOF",
    type: "error",
  },
  {
    inputs: [],
    name: "PROVING_FAILED",
    type: "error",
  },
  {
    inputs: [],
    name: "REENTRANT_CALL",
    type: "error",
  },
  {
    inputs: [],
    name: "RESOLVER_DENIED",
    type: "error",
  },
  {
    inputs: [],
    name: "RESOLVER_INVALID_MANAGER",
    type: "error",
  },
  {
    inputs: [],
    name: "RESOLVER_UNEXPECTED_CHAINID",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint64",
        name: "chainId",
        type: "uint64",
      },
      {
        internalType: "string",
        name: "name",
        type: "string",
      },
    ],
    name: "RESOLVER_ZERO_ADDR",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "previousAdmin",
        type: "address",
      },
      {
        indexed: false,
        internalType: "address",
        name: "newAdmin",
        type: "address",
      },
    ],
    name: "AdminChanged",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "uint256",
        name: "operationId",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "approvalBits",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "proofSubmitted",
        type: "bool",
      },
    ],
    name: "Approved",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "beacon",
        type: "address",
      },
    ],
    name: "BeaconUpgraded",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint32",
        name: "version",
        type: "uint32",
      },
      {
        indexed: false,
        internalType: "address[]",
        name: "guardians",
        type: "address[]",
      },
    ],
    name: "GuardiansUpdated",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint8",
        name: "version",
        type: "uint8",
      },
    ],
    name: "Initialized",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "Paused",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "Unpaused",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "implementation",
        type: "address",
      },
    ],
    name: "Upgraded",
    type: "event",
  },
  {
    inputs: [],
    name: "MIN_NUM_GUARDIANS",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "addressManager",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "bytes32",
            name: "l1Hash",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "difficulty",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "blobHash",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "extraData",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "depositsHash",
            type: "bytes32",
          },
          {
            internalType: "address",
            name: "coinbase",
            type: "address",
          },
          {
            internalType: "uint64",
            name: "id",
            type: "uint64",
          },
          {
            internalType: "uint32",
            name: "gasLimit",
            type: "uint32",
          },
          {
            internalType: "uint64",
            name: "timestamp",
            type: "uint64",
          },
          {
            internalType: "uint64",
            name: "l1Height",
            type: "uint64",
          },
          {
            internalType: "uint24",
            name: "txListByteOffset",
            type: "uint24",
          },
          {
            internalType: "uint24",
            name: "txListByteSize",
            type: "uint24",
          },
          {
            internalType: "uint16",
            name: "minTier",
            type: "uint16",
          },
          {
            internalType: "bool",
            name: "blobUsed",
            type: "bool",
          },
          {
            internalType: "bytes32",
            name: "parentMetaHash",
            type: "bytes32",
          },
        ],
        internalType: "struct TaikoData.BlockMetadata",
        name: "meta",
        type: "tuple",
      },
      {
        components: [
          {
            internalType: "bytes32",
            name: "parentHash",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "blockHash",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "signalRoot",
            type: "bytes32",
          },
          {
            internalType: "bytes32",
            name: "graffiti",
            type: "bytes32",
          },
        ],
        internalType: "struct TaikoData.Transition",
        name: "tran",
        type: "tuple",
      },
      {
        components: [
          {
            internalType: "uint16",
            name: "tier",
            type: "uint16",
          },
          {
            internalType: "bytes",
            name: "data",
            type: "bytes",
          },
        ],
        internalType: "struct TaikoData.TierProof",
        name: "proof",
        type: "tuple",
      },
    ],
    name: "approve",
    outputs: [
      {
        internalType: "bool",
        name: "approved",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "guardian",
        type: "address",
      },
    ],
    name: "guardianIds",
    outputs: [
      {
        internalType: "uint256",
        name: "id",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    name: "guardians",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_addressManager",
        type: "address",
      },
    ],
    name: "init",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "hash",
        type: "bytes32",
      },
    ],
    name: "isApproved",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "minGuardians",
    outputs: [
      {
        internalType: "uint32",
        name: "",
        type: "uint32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "numGuardians",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "pause",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "paused",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "proxiableUUID",
    outputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint64",
        name: "chainId",
        type: "uint64",
      },
      {
        internalType: "bytes32",
        name: "name",
        type: "bytes32",
      },
      {
        internalType: "bool",
        name: "allowZeroAddress",
        type: "bool",
      },
    ],
    name: "resolve",
    outputs: [
      {
        internalType: "address payable",
        name: "addr",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "name",
        type: "bytes32",
      },
      {
        internalType: "bool",
        name: "allowZeroAddress",
        type: "bool",
      },
    ],
    name: "resolve",
    outputs: [
      {
        internalType: "address payable",
        name: "addr",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address[]",
        name: "_guardians",
        type: "address[]",
      },
      {
        internalType: "uint8",
        name: "_minGuardians",
        type: "uint8",
      },
    ],
    name: "setGuardians",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "unpause",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation",
        type: "address",
      },
    ],
    name: "upgradeTo",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "upgradeToAndCall",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [],
    name: "version",
    outputs: [
      {
        internalType: "uint32",
        name: "",
        type: "uint32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];
