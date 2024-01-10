# Taiko native vs. wrapped token bridging

![Wrapped_vs_Native](./native_bridge/native_support.png "Wrapped vs. Native bridging")

Taiko's briding concept is a lock-and-mint type. It simply means (the red path above) on the canonical chain we take custody of the assets and on the destination chain we mint the wrapped counterpart. When someone wants to bridge back (from destination to canonical) it will first burn the tokens, then release the funds on the canonical chain.

But there might be some incentives (e.g.: adoption, liquidity fragmentation, etc.) when deploying a native token on the destination chain is beneficial. For this reason Taiko introduced the possibilty of deploying the canonical assets (together with all their sub/parent/proxy contracts) and plug it into our ERC20Vault via adaptors (green path).

Important to note that while wrapped asset briding is 'automatical', the native one requires the willingness and efforts from Taiko side (and maybe also original token issuer green light to recognise officially as "native"), to support that type of asset-transfer.

## Howto

There are some steps to do in order to facilitate native token bridging. In the next steps, here is a TLDR breakdown how we do it with USDC.

1. Deploy the same (bytecode equivalent) ERC-20 token on L2. An example of the contracts + deployments can be found in our [USDC repo](https://github.com/taikoxyz/USDC).
2. Deploy adaptor (e.g.: [USDC adaptor](../contracts/tokenvault/adaptors/USDCAdaptor.sol)). As this will serve as the plug-in to our `ERC20Vault` for custom (native) tokens. This adaptor serves multiple purposes. It is also a wrapper around the native token in a way - that it matches our conform `ERC20Vault` interfaces so that we can be sure any kind of native ERC-20 can be supported on L2. Also can handle custom logic required by the native asset (roles handling, specific logic, etc.).
3. Transfer the ownership (if not already owned by) to `ERC20Vault` owner since those 2 have to be owned by the same address. (!IMPORTANT! Not the token owned by the same owner, but the token adpter! USDC owner will still be Circle on L2.)
4. Since our bridge is permissionless, there might have been some USDC bridge operations in the past. It would mean, there is already an existing `BridgedUSDC` on our L2. To overcome liquidity fragmentation, we (Taiko) need to call `ERC20Vault` `changeBridgedToken()` function with the appropriate parameters. This way the "old" `BridgedUSDC` can be migrated to this new native token and the bridging operation will mint into the new token frm that point on.

The above steps (2. - 4.) is incorporated into the script [DeployUSDCAdaptor.s.sol](../script/DeployUSDCAdaptor.s.sol).
