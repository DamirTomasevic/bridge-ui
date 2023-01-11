# 🚀 Deploy a contract

These steps will show you how to deploy a smart contract to Taiko A1 using Foundry. You can find the latest Foundry docs at the Foundry Book: https://book.getfoundry.sh/getting-started/first-steps. This guide uses snippets / examples from there.

## Prerequisites

- Have the private key to an account that has some ETH on Taiko A1. This is to pay the small transaction fee for deploying the contract. You can check the following articles if you are unsure how to do this:
  - [With MetaMask](https://www.herongyang.com/Ethereum/Ethereum-Account-Public-Private-Key-Example.html)
  - [With OpenSSL](https://gist.github.com/miguelmota/3793b160992b4ea0b616497b8e5aee2f)
  - [With ethereumjs-wallet](https://piyopiyo.medium.com/how-to-generate-ethereum-private-key-and-address-in-local-offline-environment-90294308593c)

## Steps

1. [Install Foundry](https://book.getfoundry.sh/getting-started/installation)
2. Create a project with Foundry, and `cd` into it:
   ```sh
   forge init hello_foundry && cd hello_foundry
   ```
3. Deploy the contract from your project, located at `src/Counter.sol`. Replace `<YOUR_PRIVATE_KEY>` with your private key, mentioned in the previous prerequisites section.
   ```sh
   forge create --legacy --rpc-url https://l2rpc.a1.taiko.xyz --private-key <YOUR_PRIVATE_KEY> src/Counter.sol:Counter
   ```
   Note: Remove "<" and ">" from <YOUR_PRIVATE_KEY>

We're using the `--legacy` flag because EIP-1559 is currently disabled on Taiko. We've plans to re-enable it in the future.

You can use the block explorer to verify that the contract was deployed: https://l2explorer.a1.taiko.xyz/
