---
title: Mainnet Addresses
description: Network reference page describing various important addresses on Taiko.
---

## Ethereum contracts

| Contract Name (Shared) | Address                                      | ENS                     |
| ---------------------- | -------------------------------------------- | ----------------------- |
| SharedAddressManager   | `0xEf9EaA1dd30a9AA1df01c36411b5F082aA65fBaa` | sam.based.taiko.eth     |
| TaikoToken             | `0x10dea67478c5F8C5E2D90e5E9B26dBe60c54d800` | token.taiko.eth         |
| SignalService          | `0x9e0a24964e5397B566c1ed39258e21aB5E35C77C` | signals.based.taiko.eth |
| Bridge                 | `0xd60247c6848B7Ca29eDdF63AA924E53dB6Ddd8EC` | bridge.based.taiko.eth  |
| QuotaManager           | `0x91f67118DD47d502B1f0C354D0611997B022f29E` | N/A                     |
| ERC20Vault             | `0x996282cA11E5DEb6B5D122CC3B9A1FcAAD4415Ab` | v20.based.taiko.eth     |
| ERC721Vault            | `0x0b470dd3A0e1C41228856Fb319649E7c08f419Aa` | v721.based.taiko.eth    |
| ERC1155Vault           | `0xaf145913EA4a56BE22E120ED9C24589659881702` | v1155.based.taiko.eth   |
| BridgedERC20           | `0xcc5d488073FA918cBbd73B9A523F3858C4de7372` | N/A                     |
| BridgedERC721          | `0xc4096E9ff1526Bd1840B65e9f45695135aC12De7` | N/A                     |
| BridgedERC1155         | `0x39E4C1214e733639d059979079A151911e42791d` | N/A                     |

| Contract Name (Rollup-Specific) | Address                                      | ENS                        |
| ------------------------------- | -------------------------------------------- | -------------------------- |
| TaikoL1                         | `0x06a9Ab27c7e2255df1815E6CC0168d7755Feb19a` | based.taiko.eth            |
| RollupAddressManager            | `0x579f40D0BE111b823962043702cabe6Aaa290780` | ram.based.taiko.eth        |
| GuardianProver                  | `0xE3D777143Ea25A6E031d1e921F396750885f43aC` | guardians.based.taiko.eth  |
| GuardianProverMinority          | `0x579A8d63a2Db646284CBFE31FE5082c9989E985c` | guardians1.based.taiko.eth |
| AssignmentHook                  | `0x537a2f0D3a5879b41BCb5A2afE2EA5c4961796F6` | N/A                        |
| TierProvider                    | `0x4cffe56C947E26D07C14020499776DB3e9AE3a23` | N/A                        |
| SgxVerifier                     | `0xb0f3186FC1963f774f52ff455DC86aEdD0b31F81` | N/A                        |
| AutomataDcapAttestation         | `0x8d7C954960a36a7596d7eA4945dDf891967ca8A3` | N/A                        |
| PemCertChainLib                 | `0x02772b7B3a5Bea0141C993Dbb8D0733C19F46169` | N/A                        |
| P256Verifier                    | `0x11A9ebA17EbF92b40fcf9a640Ebbc47Db6fBeab0` | N/A                        |
| SigVerifyLib                    | `0x47bB416ee947fE4a4b655011aF7d6E3A1B80E6e9` | N/A                        |
| TokenUnlock                     | `0x035AFfC82612de31E9Db2259B9482D0Dd53B7819` | N/A                        |
| ProverSet                       | `0x34f2B21107AfE3584949c184A1E6236FFDAC4f6F` | N/A                        |
| labprover                       | `0x68d30f47F19c07bCCEf4Ac7FAE2Dc12FCa3e0dC9` | labprover.taiko.eth        |
| labcontester                    | `0xa01d464ca3982DAa97B19fa7F8a232eB11A9DDb3` | labcontester.taiko.eth     |

## Taiko (Mainnet) contracts

| Contract Name (Shared) | Address                                      |
| ---------------------- | -------------------------------------------- |
| Bridge                 | `0x1670000000000000000000000000000000000001` |
| ERC20Vault             | `0x1670000000000000000000000000000000000002` |
| ERC721Vault            | `0x1670000000000000000000000000000000000003` |
| ERC1155Vault           | `0x1670000000000000000000000000000000000004` |
| SignalService          | `0x1670000000000000000000000000000000000005` |
| SharedAddressManager   | `0x1670000000000000000000000000000000000006` |

| Contract Name (Rollup-Specific) | Address                                      |
| ------------------------------- | -------------------------------------------- |
| TaikoL2                         | `0x1670000000000000000000000000000000010001` |
| RollupAddressManager            | `0x1670000000000000000000000000000000010002` |
| WETH                            | `0xA51894664A773981C6C112C43ce576f315d5b1B6` |

## Rollup contracts owner

The rollup contracts owner is `admin.taiko.eth`, public address is `0x9CBeE534B5D8a6280e01a14844Ee8aF350399C7F`.

## Taiko Labs' proposers and provers addresses

| Name         | Address                                      |
| ------------ | -------------------------------------------- |
| Proposer #1  | `0x000000633b68f5d8d3a86593ebb815b4663bcbe0` |
| Prover #1    | `0x000000629FBCf27A347d1AEbA658435230D74a5f` |
| Contester #1 | `0x00000027F51a57E7FcBC4b481d15fcE5BE68b30B` |

## Taiko Labs' bootnode addresses

Find the latest bootnodes here in [simple-taiko-node](https://github.com/taikoxyz/simple-taiko-node/blob/mainnet/.env.sample).
