# Taiko protocol

This package contains rollup contracts on both L1 and L2, along with other assisting code. Taiko L2's chain ID is [167](https://github.com/ethereum-lists/chains/pull/1611).

## Compile

To compile smart contracts, run:

```sh
pnpm compile
```

If you run into `Error: Unknown version provided`, you should upgrade your foundry installation by running `curl -L https://foundry.paradigm.xyz | bash`.

## Deploy

Deploy TaikoL1 on foundry network:

```sh
pnpm deploy:foundry
```

## Test

Run test cases on foundry network:

```sh
pnpm test
```

Run test cases that require a running go-ethereum node:

```sh
pnpm test:integration
```

## Generate L2 genesis JSON's `alloc` field

Start by creating a `config.js`, for example:

```javascript
module.exports = {
  // Owner address of the pre-deployed L2 contracts.
  contractOwner: "0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39",
  // Chain ID of the Taiko L2 network.
  chainId: 167,
  // Account address and pre-mint ETH amount as key-value pairs.
  seedAccounts: [
    { "0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39": 1024 },
    { "0x79fcdef22feed20eddacbb2587640e45491b757f": 1024 },
  ],
  // L2 EIP-1559 baseFee calculation related fields.
  param1559: {
    gasExcess: 1,
  },
  // Option to pre-deploy an ERC-20 token.
  predeployERC20: true,
};
```

Next, run the generation script:

```sh
pnpm compile && pnpm generate:genesis config.js
```

The script will output two JSON files under `./deployments`:

- `l2_genesis_alloc.json`: the `alloc` field which will be used in L2 genesis JSON file
- `l2_genesis_storage_layout.json`: the storage layout of those pre-deployed contracts

## Using Foundry

This project also integrates with Foundry for building and testing contracts.

- To compile using foundry: `forge build` or `pnpm compile`
- To run foundry tests: `forge test --gas-report -vvv` or `pnpm test:foundry`

Note that compiling with foundry uses dependencies inside the `lib` dir (instead of `node_modules`).
