---
title: TokenVault
---

## TokenVault

This vault holds all ERC20 tokens (but not Ether) that users have deposited.
It also manages the mapping between canonical ERC20 tokens and their bridged
tokens.

### CanonicalERC20

```solidity
struct CanonicalERC20 {
  uint256 chainId;
  address addr;
  uint8 decimals;
  string symbol;
  string name;
}
```

### isBridgedToken

```solidity
mapping(address => bool) isBridgedToken
```

### bridgedToCanonical

```solidity
mapping(address => struct TokenVault.CanonicalERC20) bridgedToCanonical
```

### canonicalToBridged

```solidity
mapping(uint256 => mapping(address => address)) canonicalToBridged
```

### BridgedERC20Deployed

```solidity
event BridgedERC20Deployed(uint256 srcChainId, address canonicalToken, address bridgedToken, string canonicalTokenSymbol, string canonicalTokenName, uint8 canonicalTokenDecimal)
```

### EtherSent

```solidity
event EtherSent(address to, uint256 destChainId, uint256 amount, bytes32 signal)
```

### EtherReceived

```solidity
event EtherReceived(address from, uint256 amount)
```

### ERC20Sent

```solidity
event ERC20Sent(address to, uint256 destChainId, address token, uint256 amount, bytes32 signal)
```

### ERC20Received

```solidity
event ERC20Received(address to, address from, uint256 srcChainId, address token, uint256 amount)
```

### init

```solidity
function init(address addressManager) external
```

### sendEther

```solidity
function sendEther(uint256 destChainId, address to, uint256 gasLimit, uint256 processingFee, address refundAddress, string memo) external payable
```

Transfers Ether to this vault and sends a message to the destination
chain so the user can receive Ether.

_Ether is held by Bridges on L1 and by the EtherVault on L2,
not TokenVaults._

#### Parameters

| Name          | Type    | Description                                            |
| ------------- | ------- | ------------------------------------------------------ |
| destChainId   | uint256 | The destination chain ID where the `to` address lives. |
| to            | address | The destination address.                               |
| gasLimit      | uint256 |                                                        |
| processingFee | uint256 | @custom:see Bridge                                     |
| refundAddress | address |                                                        |
| memo          | string  |                                                        |

### sendERC20

```solidity
function sendERC20(uint256 destChainId, address to, address token, uint256 amount, uint256 gasLimit, uint256 processingFee, address refundAddress, string memo) external payable
```

Transfers ERC20 tokens to this vault and sends a message to the
destination chain so the user can receive the same amount of tokens
by invoking the message call.

#### Parameters

| Name          | Type    | Description                                                                                                  |
| ------------- | ------- | ------------------------------------------------------------------------------------------------------------ |
| destChainId   | uint256 | The destination chain ID where the `to` address lives.                                                       |
| to            | address | The destination address.                                                                                     |
| token         | address | The address of the token to be sent.                                                                         |
| amount        | uint256 | The amount of token to be transferred.                                                                       |
| gasLimit      | uint256 | @custom:see Bridge                                                                                           |
| processingFee | uint256 | @custom:see Bridge                                                                                           |
| refundAddress | address | The fee refund address. If this address is address(0), extra fees will be refunded back to the `to` address. |
| memo          | string  |                                                                                                              |

### receiveERC20

```solidity
function receiveERC20(struct TokenVault.CanonicalERC20 canonicalToken, address from, address to, uint256 amount) external
```

_This function can only be called by the bridge contract while
invoking a message call._

#### Parameters

| Name           | Type                             | Description                                                                                                          |
| -------------- | -------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| canonicalToken | struct TokenVault.CanonicalERC20 | The canonical ERC20 token which may or may not live on this chain. If not, a BridgedERC20 contract will be deployed. |
| from           | address                          | The source address.                                                                                                  |
| to             | address                          | The destination address.                                                                                             |
| amount         | uint256                          | The amount of tokens to be sent. 0 is a valid value.                                                                 |
