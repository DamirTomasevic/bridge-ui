# Taiko Mainnet Contract Logs - L1

## Notes

1. Code used on mainnet must correspond to a commit on the main branch of the official repo: https://github.com/taikoxyz/taiko-mono.

## Shared

#### shared_address_manager

- ens: `sam.based.taiko.eth`
- proxy: `0xEf9EaA1dd30a9AA1df01c36411b5F082aA65fBaa`
- impl: `0xF1cA1F1A068468E1dcF90dA6add185467de80943`
- owner: `admin.taiko.eth`
- names:
  - taiko_token: `0x10dea67478c5F8C5E2D90e5E9B26dBe60c54d800`
  - signal_service: `0x9e0a24964e5397B566c1ed39258e21aB5E35C77C`
  - signal_service@167000: `0x1670000000000000000000000000000000000005`
  - bridge: `0xd60247c6848B7Ca29eDdF63AA924E53dB6Ddd8EC`
  - bridge@167000: `0x1670000000000000000000000000000000000001`
  - erc20_vault: `0x996282cA11E5DEb6B5D122CC3B9A1FcAAD4415Ab`
  - erc20_vault@167000: `0x1670000000000000000000000000000000000002`
  - erc721_vault: `0x0b470dd3A0e1C41228856Fb319649E7c08f419Aa`
  - erc721_vault@167000: `0x1670000000000000000000000000000000000003`
  - erc1155_vault: `0xaf145913EA4a56BE22E120ED9C24589659881702`
  - erc1155_vault@167000: `0x1670000000000000000000000000000000000004`
  - bridged_erc1155: `0x3c90963cFBa436400B0F9C46Aa9224cB379c2c40`
  - bridged_erc721: `0xC3310905E2BC9Cfb198695B75EF3e5B69C6A1Bf7`
  - bridged_erc20: `0x79BC0Aada00fcF6E7AB514Bfeb093b5Fae3653e3`
  - bridge_watchdog: `0x00000291ab79c55dc4fcd97dfba4880df4b93624`
  - quota_manager: `0x91f67118DD47d502B1f0C354D0611997B022f29E`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - admin.taiko.eth accepted the ownership @tx`0x0ed114fee6de4e3e2206cea44e6632ec0c4588f73648d98d8df5dc0183b07885`
  - Upgraded from `0x9cA1Ab10c9fAc5153F8b78E67f03aAa69C9c6A15` to `0xF1cA1F1A068468E1dcF90dA6add185467de80943` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - `Init2()` called @tx`0x7311fee56f87294e336393b55939489bc1e810c402f304013475d04c90ca32a9`

#### taiko_token

- ens: `token.taiko.eth`
- proxy: `0x10dea67478c5F8C5E2D90e5E9B26dBe60c54d800`
- impl: `0xea53c0f4b129Cf3f3FBA896F9f23ca18246e9B3c`
- owner: `admin.taiko.eth`
- logs:
  - deployed on April 25, 2024 @commit`2f6d3c62e`
  - upgraded impl from `0x9ae1a067f9655dd0511390e3d70bb25933ae61eb` to `0xea53c0f4b129Cf3f3FBA896F9f23ca18246e9B3c` @commit`b90b932` and,
  - Changed owner from `labs.taiko.eth` to `admin.taiko.eth` @tx`0x7d82794932540ed9edd259e58f6ef8ae21a49beada7f0224638f888f7149c01c`
  - Accept owner @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### signal_service

- ens: `signals.based.taiko.eth`
- proxy: `0x9e0a24964e5397B566c1ed39258e21aB5E35C77C`
- impl: `0xB11Cd7bA46a12F238b4Ad831f6F296262C1e652d`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - admin.taiko.eth accepted the ownership @tx`0x0ed114fee6de4e3e2206cea44e6632ec0c4588f73648d98d8df5dc0183b07885`
  - upgraded from `0xE1d91bAE44B70bD66e8b688B8421fD62dcC33c72` to `0xB11Cd7bA46a12F238b4Ad831f6F296262C1e652d` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### bridge

- ens: `bridge.based.taiko.eth`
- proxy: `0xd60247c6848B7Ca29eDdF63AA924E53dB6Ddd8EC`
- impl: `0xc71CC3B0a47149878fad337fb2ca54E546A645ba`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - admin.taiko.eth accepted the ownership @tx`0x0ed114fee6de4e3e2206cea44e6632ec0c4588f73648d98d8df5dc0183b07885`
  - upgraded from `0x91d593d34f2E1904cDCe3D5290a74563F87bCF6f` to `0x4A1091c2fb37D9C4a661c2384Ff539d94CCF853D` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - upgraded from `0x4A1091c2fb37D9C4a661c2384Ff539d94CCF853D` to `0xc71CC3B0a47149878fad337fb2ca54E546A645ba` @commit`b955e0e` @tx`0x5a60c5815947a199cc84e1bc75539e01a202597b20c1f87bd9d02f8be6453abd`
  - called `selfDelegate` for Taiko Token @tx`0x740c255322873b3feb62ad1de71b51417053787328eae3aa84557c953463d55f`

#### quota_manager

- proxy: `0x91f67118DD47d502B1f0C354D0611997B022f29E`
- impl: `0x49c5e5F131314Bb24b17E249960F8B12F925ef22`
- owner: `admin.taiko.eth`
- quota:
  - ETH: `64516129032258064516` (`200_000 * 1 ether / 3100`)
  - WETH(`0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2`): `64516129032258064516` (`200_000 * 1 ether / 3100`)
  - TKO(`0x10dea67478c5F8C5E2D90e5E9B26dBe60c54d800`): `40000000000000000000000` (`200_000 * 1e18 / 5`)
  - USDT(`0xdAC17F958D2ee523a2206206994597C13D831ec7`): `200000000000` (`200_000 * 1e6`)
  - USDC(`0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48`): `200000000000` (`200_000 * 1e6`)
- logs:
  - deployed on May 13, 2024 at commit `b90b932`
  - admin.taiko.eth accepted the ownership @tx`0x2d6ce1781137899f65c1810e42f556c27caa4e9bd13077ba5bc7a9a0975eefcb`

#### erc20_vault

- ens: `v20.based.taiko.eth`
- proxy: `0x996282cA11E5DEb6B5D122CC3B9A1FcAAD4415Ab`
- impl: `0xC722d9f3f8D60288589F7f67a9CFAd34d3B9bf8E`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - upgraded from `0x15D9F7e12aEa18DAEF5c651fBf97567CAd4a4BEc` to `0xC722d9f3f8D60288589F7f67a9CFAd34d3B9bf8E` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### erc721_vault

- ens: `v721.based.taiko.eth`
- proxy: `0x0b470dd3A0e1C41228856Fb319649E7c08f419Aa`
- impl: `0x41A7BDD153a5AfFb10Ed1AD3D6a4e5ad001495FA`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - upgraded from `0xEC04849E7722Fd69797a155796Db75aC8F94f692` to `0x41A7BDD153a5AfFb10Ed1AD3D6a4e5ad001495FA` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### erc1155_vault

- ens: `v1155.based.taiko.eth`
- proxy: `0xaf145913EA4a56BE22E120ED9C24589659881702`
- impl: `0xd90b5fcf8d00d333d107E4Ab7F94c0c0A41CDcfE`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - upgraded from `0x7748dA086A2e6EDd8Db97eD236840910013c6396` to `0xd90b5fcf8d00d333d107E4Ab7F94c0c0A41CDcfE` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### bridged_erc20

- impl: `0xcc5d488073FA918cBbd73B9A523F3858C4de7372`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`

#### bridged_erc721

- impl: `0xc4096E9ff1526Bd1840B65e9f45695135aC12De7`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`

#### bridged_erc1155

- impl: `0x39E4C1214e733639d059979079A151911e42791d`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`

## Rollup Specific

#### rollup_address_manager

- ens: `ram.based.taiko.eth`
- proxy: `0x579f40D0BE111b823962043702cabe6Aaa290780`
- impl: `0xF1cA1F1A068468E1dcF90dA6add185467de80943`
- names:
  - taiko_token: `0x10dea67478c5F8C5E2D90e5E9B26dBe60c54d800`
  - signal_service: `0x9e0a24964e5397B566c1ed39258e21aB5E35C77C`
  - bridge: `0xd60247c6848B7Ca29eDdF63AA924E53dB6Ddd8EC`
  - taiko: `0x06a9Ab27c7e2255df1815E6CC0168d7755Feb19a`
  - tier_provider: `0x4cffe56C947E26D07C14020499776DB3e9AE3a23`
  - tier_sgx: `0xb0f3186FC1963f774f52ff455DC86aEdD0b31F81`
  - tier_guardian_minority: `0x579A8d63a2Db646284CBFE31FE5082c9989E985c`
  - tier_guardian: `0xE3D777143Ea25A6E031d1e921F396750885f43aC`
  - automata_dcap_attestation: `0x8d7C954960a36a7596d7eA4945dDf891967ca8A3`
  - assignment_hook: `0x537a2f0D3a5879b41BCb5A2afE2EA5c4961796F6`
  - prover_set: `0x34f2B21107AfE3584949c184A1E6236FFDAC4f6F`
  - proposer_one: `0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045` vitalik.eth
  - proposer: `0x000000633b68f5d8d3a86593ebb815b4663bcbe0`
  - chain_watchdog: `0xE3D777143Ea25A6E031d1e921F396750885f43aC`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - admin.taiko.eth accepted the ownership @tx`0x0ed114fee6de4e3e2206cea44e6632ec0c4588f73648d98d8df5dc0183b07885`
  - Upgraded from `0xd912aB787624c9eb96a37e658e9596e114360440` to `0xF1cA1F1A068468E1dcF90dA6add185467de80943` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - `Init2()` called @tx`0x7311fee56f87294e336393b55939489bc1e810c402f304013475d04c90ca32a9`
  - register `chain_watchdog` on May 21 @tx`0xaed098ad0c93113e401f61358f963501f40a046c5b5b659a1610f10120a9a86b`
  - register`prover_set` to `0x34f2B21107AfE3584949c184A1E6236FFDAC4f6F` @tx`0x252cd7fcb6e02a71c0770d00f2f2476d5dd469a4fb5df622fe7bf6280d8a4100`

#### taikoL1

- ens: `based.taiko.eth`
- proxy: `0x06a9Ab27c7e2255df1815E6CC0168d7755Feb19a`
- impl: `0xe0A5D394878723CEAEC8B993e04756DF1f4B44eF`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - Upgraded from `0x99Ba70E62cab0cB983e66F72330fBDDC11d85501` to `0x9fBBedBBcBb753E7214BE08381efE10d89D712fE` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - `Init2()` called and reset block hash to `0411D9F84A525864E0A7E8BB51667D49C6BF73820AF9E4BC76EA66ADB6BE8903` @tx`0x7311fee56f87294e336393b55939489bc1e810c402f304013475d04c90ca32a9`
  - Upgraded from `0x9fBBedBBcBb753E7214BE08381efE10d89D712fE` to `0xe0A5D394878723CEAEC8B993e04756DF1f4B44eF` on May 21 @commit`c817e76d9` @tx`0xaed098ad0c93113e401f61358f963501f40a046c5b5b659a1610f10120a9a86b`
  - `resetGenesisHash()` called to reset genesis block hash to `0x90bc60466882de9637e269e87abab53c9108cf9113188bc4f80bcfcb10e489b9` on May 22 @tx`0x5a60c5815947a199cc84e1bc75539e01a202597b20c1f87bd9d02f8be6453abd`

#### assignment_hook

- proxy: `0x537a2f0D3a5879b41BCb5A2afE2EA5c4961796F6`
- impl: `0xe226fAd08E2f0AE68C32Eb5d8210fFeDB736Fb0d`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - Upgraded from `0x4f664222C3fF6207558A745648B568D095dDA170` to `0xe226fAd08E2f0AE68C32Eb5d8210fFeDB736Fb0d` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### tier_provider

- impl: `0x4cffe56C947E26D07C14020499776DB3e9AE3a23`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - deployed on May 15, 2024 @commit`cd5144255`

#### tier_sgx

- proxy: `0xb0f3186FC1963f774f52ff455DC86aEdD0b31F81`
- impl: `0xf381868DD6B2aC8cca468D63B42F9040DE2257E9`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - Upgraded from `0x3f54067EF5d8B414Bdb1945cdF482BD158Aad175` to `0xf381868DD6B2aC8cca468D63B42F9040DE2257E9` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`

#### guardian_prover_minority

- ens: `guardians1.based.taiko.eth`
- proxy: `0x579A8d63a2Db646284CBFE31FE5082c9989E985c`
- impl: `0x468F6A9C0ad2e9C8370687D2844A9e70fE942d5c`
- owner: `admin.taiko.eth`
- guardianProvers:
  - `0x000012dd12a6d9dd2045f5e2594f4996b99a5d33`
  - `0x0cAC6E2Fd10e92Bf798341Ad0A57b5Cb39DA8D0D`
  - `0xd6BB974bc47626E3547426efa4CA2A8d7DFCccdf`
  - `0xd26c4e85BC2fAAc27a320987e340971cF3b47d51`
  - `0xC384B679c028787166b9B3725aC14A60da205861`
  - `0x1602958A85494cd9C3e0D6672BA0eE42b95B4200`
  - `0x5CfEb9a72256B1b49dc2C98b1b7b99d172D50B68`
  - `0x1DB8Ac9f19AbdD60A6418383BfA56A4450aa80C6`
- minGuardiansReached: `1`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - admin.taiko.eth accepted the ownership @tx`0x0ed114fee6de4e3e2206cea44e6632ec0c4588f73648d98d8df5dc0183b07885`
  - Upgraded from `0x717DC5E3814591790BcB1fD9259eEdA7c14ce9CF` to `0x750221E951b77a2Cb4046De41Ec5F6d1aa7942D2` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - Upgraded from `0x750221E951b77a2Cb4046De41Ec5F6d1aa7942D2` to `0x253E47F2b1e91F2001d3578aeB24C0ccF464b65e` @commit`cd5144255` @tx`0x8030569e293baddbc4e8b26688a1ecf14a231d86c90e9d02dad1e919ea2f3964`
  - Upgraded from `0x253E47F2b1e91F2001d3578aeB24C0ccF464b65e` to `0x468F6A9C0ad2e9C8370687D2844A9e70fE942d5c` @commit`b955e0e` @tx`0x5a60c5815947a199cc84e1bc75539e01a202597b20c1f87bd9d02f8be6453abd`

#### guardian_prover

- ens: `guardians.based.taiko.eth`
- proxy: `0xE3D777143Ea25A6E031d1e921F396750885f43aC`
- impl: `0x468F6A9C0ad2e9C8370687D2844A9e70fE942d5c`
- owner: `admin.taiko.eth`
- guardianProvers:
  - `0x000012dd12a6d9dd2045f5e2594f4996b99a5d33`
  - `0x0cAC6E2Fd10e92Bf798341Ad0A57b5Cb39DA8D0D`
  - `0xd6BB974bc47626E3547426efa4CA2A8d7DFCccdf`
  - `0xd26c4e85BC2fAAc27a320987e340971cF3b47d51`
  - `0xC384B679c028787166b9B3725aC14A60da205861`
  - `0x1602958A85494cd9C3e0D6672BA0eE42b95B4200`
  - `0x5CfEb9a72256B1b49dc2C98b1b7b99d172D50B68`
  - `0x1DB8Ac9f19AbdD60A6418383BfA56A4450aa80C6`
- minGuardiansReached: `6`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - Upgraded from `0x717DC5E3814591790BcB1fD9259eEdA7c14ce9CF` to `0x750221E951b77a2Cb4046De41Ec5F6d1aa7942D2` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - Upgraded from `0x750221E951b77a2Cb4046De41Ec5F6d1aa7942D2` to `0x253E47F2b1e91F2001d3578aeB24C0ccF464b65e` @commit`cd5144255` @tx`0x8030569e293baddbc4e8b26688a1ecf14a231d86c90e9d02dad1e919ea2f3964`
  - Upgraded from `0x253E47F2b1e91F2001d3578aeB24C0ccF464b65e` to `0x468F6A9C0ad2e9C8370687D2844A9e70fE942d5c` @commit`b955e0e` @tx`0x5a60c5815947a199cc84e1bc75539e01a202597b20c1f87bd9d02f8be6453abd`

#### p256_verifier

- impl: `0x11A9ebA17EbF92b40fcf9a640Ebbc47Db6fBeab0`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`

#### sig_verify_lib

- impl: `0x47bB416ee947fE4a4b655011aF7d6E3A1B80E6e9`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`

#### pem_cert_chain_lib

- impl: `0x02772b7B3a5Bea0141C993Dbb8D0733C19F46169`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`

#### automata_dcap_attestation

- proxy: `0x8d7C954960a36a7596d7eA4945dDf891967ca8A3`
- impl: `0x5f73f0AdC7dAA6134Fe751C4a78d524f9384e0B5`
- owner: `admin.taiko.eth`
- logs:
  - deployed on May 1, 2024 @commit`56dddf2b6`
  - Upgraded from `0xEE8FC1dbb8D345f5bF35dFb939C6f9EdC5fCDAFc` to `0xde1b1FBe7D721af4A56651272ef91A59B7303323` @commit`b90b932` @tx`0x416560cd96dc75ccffebe889e8d1ab3e08b33f814dc4a2bf7c6f9555071d1f6f`
  - Called `configureTcbInfoJson` and `configureQeIdentityJson` @commit`b90b932` @tx`0x2d6ce1781137899f65c1810e42f556c27caa4e9bd13077ba5bc7a9a0975eefcb`
  - Called `configureTcbInfoJson` and `configureQeIdentityJson` @commit`cd5144255` @tx`0x8030569e293baddbc4e8b26688a1ecf14a231d86c90e9d02dad1e919ea2f3964`
  - Upgraded from `0xde1b1FBe7D721af4A56651272ef91A59B7303323` to `0x5f73f0AdC7dAA6134Fe751C4a78d524f9384e0B5` @commit`3740dc0` @tx`0x46a6d47c15505a1259c64d1e09353680e525b2706dd9e095e15019dda7c1b295`
  - Called `configureTcbInfoJson` @commit`3740dc0` @tx`0x46a6d47c15505a1259c64d1e09353680e525b2706dd9e095e15019dda7c1b295`

### token_unlock

- impl: `0x035AFfC82612de31E9Db2259B9482D0Dd53B7819.`
- logs:
  - deployed @commit`bca493f` @tx`0x0a4a63715257b766ca06e7e87ee25088d557c460e50120208b31666c83fc68bc`

### prover_set

- impl: `0x34f2B21107AfE3584949c184A1E6236FFDAC4f6F`
- logs:
  - deployed @commit`bca493f` @tx`0xfacd0f26e3ec4bf1f949637373483fcfe9a960dfc427d6fa62b116907bac3373`

### labprovers.taiko.eth

- ens: `labprovers.taiko.eth`
- proxy: `0x68d30f47F19c07bCCEf4Ac7FAE2Dc12FCa3e0dC9`
- impl: `0x34f2B21107AfE3584949c184A1E6236FFDAC4f6F`
- enabled provers:
  - `0x000000629FBCf27A347d1AEbA658435230D74a5f`
  - `0x00000027F51a57E7FcBC4b481d15fcE5BE68b30B`
- logs:
  - deployed @commit`bca493f`@tx`0xf3b6af477112d0a8209506c8f310f4eb0713beebb1911ef5d11162d36d93c0ff`
  - enabled two provers (`0x000000629FBCf27A347d1AEbA658435230D74a5f` and `0x00000027F51a57E7FcBC4b481d15fcE5BE68b30B`) @tx`0xa0b1565473849bc753d395abd982e6899ecdd9e754014eebed67b69edadb61c5`
