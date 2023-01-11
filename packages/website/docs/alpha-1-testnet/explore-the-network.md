# 🔍 Explore the network

Taiko's Alpha-1 testnet consists of L1 / L2 nodes with all [Taiko protocol contracts](/docs/category/contract-documentation) deployed. The mining interval of the L1 node is set to 12 seconds.

## Endpoints

### L1

- **Block Explorer:** <https://l1explorer.a1.taiko.xyz>
- **HTTP RPC Endpoint:** <https://l1rpc.a1.taiko.xyz>
- **Web Socket RPC Endpoint:** <wss://l1ws.a1.taiko.xyz>
- **ETH faucet:** <https://l1faucet.a1.taiko.xyz>
- **Chain ID:** `31338`

### L2

- **Block Explorer:** <https://l2explorer.a1.taiko.xyz>
- **HTTP RPC Endpoint:** <https://l2rpc.a1.taiko.xyz>
- **Web Socket RPC Endpoint:** <wss://l2ws.a1.taiko.xyz>
- **ETH faucet:** <https://l2faucet.a1.taiko.xyz>
- **Chain ID:** `167003`

## Contract addresses

### L1

- **TaikoL1:** `0x7B3AF414448ba906f02a1CA307C56c4ADFF27ce7`
- **TokenVault:** `0xD0dfd5baCf160B97C8eE3ecb463F18c08673160c`
- **Bridge:** `0x3612E284D763f42f5E4CB72B1602b23DAEC3cA60`

### L2

- **TaikoL2:** `0x0000777700000000000000000000000000000001`
- **TokenVault:** `0x0000777700000000000000000000000000000002`
- **EtherVault:** `0x0000777700000000000000000000000000000003`
- **Bridge:** `0x0000777700000000000000000000000000000004`

## Cron job

There will be a cron job service that proposes empty blocks periodically (every 2 minutes).
