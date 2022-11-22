import { expect } from "chai"
import { BigNumber, Signer } from "ethers"
import { ethers } from "hardhat"
import RLP from "rlp"
import { AddressManager, Bridge, EtherVault } from "../../typechain"
import { Message } from "../utils/message"
import { Block, BlockHeader, EthGetProofResponse } from "../utils/rpc"

async function deployBridge(
    signer: Signer,
    addressManager: AddressManager,
    destChain: number,
    srcChain: number
): Promise<{ bridge: Bridge; etherVault: EtherVault }> {
    const libTrieProof = await (await ethers.getContractFactory("LibTrieProof"))
        .connect(signer)
        .deploy()

    const libBridgeProcess = await (
        await ethers.getContractFactory("LibBridgeProcess", {
            libraries: {
                LibTrieProof: libTrieProof.address,
            },
        })
    )
        .connect(signer)
        .deploy()

    const libBridgeRetry = await (
        await ethers.getContractFactory("LibBridgeRetry")
    )
        .connect(signer)
        .deploy()

    const BridgeFactory = await ethers.getContractFactory("Bridge", {
        libraries: {
            LibBridgeProcess: libBridgeProcess.address,
            LibBridgeRetry: libBridgeRetry.address,
            LibTrieProof: libTrieProof.address,
        },
    })

    const bridge: Bridge = await BridgeFactory.connect(signer).deploy()

    await bridge.connect(signer).init(addressManager.address)

    await bridge.connect(signer).enableDestChain(destChain, true)

    const etherVault: EtherVault = await (
        await ethers.getContractFactory("EtherVault")
    )
        .connect(signer)
        .deploy()

    await etherVault.connect(signer).init(addressManager.address)

    await etherVault.connect(signer).authorize(bridge.address, true)

    await etherVault.connect(signer).authorize(await signer.getAddress(), true)

    await addressManager.setAddress(
        `${srcChain}.ether_vault`,
        etherVault.address
    )

    await signer.sendTransaction({
        to: etherVault.address,
        value: BigNumber.from(100000000),
        gasLimit: 1000000,
    })

    return { bridge, etherVault }
}
describe("Bridge", function () {
    async function deployBridgeFixture() {
        const [owner, nonOwner] = await ethers.getSigners()

        const { chainId } = await ethers.provider.getNetwork()

        const srcChainId = chainId

        const enabledDestChainId = srcChainId + 1

        const addressManager: AddressManager = await (
            await ethers.getContractFactory("AddressManager")
        ).deploy()
        await addressManager.init()

        const { bridge: l1Bridge, etherVault: l1EtherVault } =
            await deployBridge(
                owner,
                addressManager,
                enabledDestChainId,
                srcChainId
            )

        // deploy protocol contract
        return {
            owner,
            nonOwner,
            l1Bridge,
            addressManager,
            enabledDestChainId,
            l1EtherVault,
            srcChainId,
        }
    }

    describe("sendMessage()", function () {
        it("throws when owner is the zero address", async () => {
            const { owner, nonOwner, l1Bridge } = await deployBridgeFixture()

            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: 5,
                owner: ethers.constants.AddressZero,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 100,
                data: ethers.constants.HashZero,
                memo: "",
            }

            await expect(l1Bridge.sendMessage(message)).to.be.revertedWith(
                "B:owner"
            )
        })

        it("throws when dest chain id is same as block.chainid", async () => {
            const { owner, nonOwner, l1Bridge } = await deployBridgeFixture()

            const network = await ethers.provider.getNetwork()
            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: network.chainId,
                owner: owner.address,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 100,
                data: ethers.constants.HashZero,
                memo: "",
            }

            await expect(l1Bridge.sendMessage(message)).to.be.revertedWith(
                "B:destChainId"
            )
        })

        it("throws when dest chain id is not enabled", async () => {
            const { owner, nonOwner, l1Bridge } = await deployBridgeFixture()

            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: 5,
                owner: owner.address,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 100,
                data: ethers.constants.HashZero,
                memo: "",
            }

            await expect(l1Bridge.sendMessage(message)).to.be.revertedWith(
                "B:destChainId"
            )
        })

        it("throws when msg.value is not the same as expected amount", async () => {
            const { owner, nonOwner, l1Bridge, enabledDestChainId } =
                await deployBridgeFixture()

            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 100,
                data: ethers.constants.HashZero,
                memo: "",
            }

            await expect(l1Bridge.sendMessage(message)).to.be.revertedWith(
                "B:value"
            )
        })

        it("emits event and is successful when message is valid, ether_vault receives the expectedAmount", async () => {
            const {
                owner,
                nonOwner,
                l1EtherVault,
                l1Bridge,
                enabledDestChainId,
            } = await deployBridgeFixture()

            const etherVaultOriginalBalance = await ethers.provider.getBalance(
                l1EtherVault.address
            )

            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 100,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const expectedAmount =
                message.depositValue + message.callValue + message.processingFee
            await expect(
                l1Bridge.sendMessage(message, {
                    value: expectedAmount,
                })
            ).to.emit(l1Bridge, "MessageSent")

            const etherVaultUpdatedBalance = await ethers.provider.getBalance(
                l1EtherVault.address
            )

            expect(etherVaultUpdatedBalance).to.be.eq(
                etherVaultOriginalBalance.add(expectedAmount)
            )
        })
    })

    describe("sendSignal()", async function () {
        it("throws when signal is empty", async function () {
            const { owner, l1Bridge } = await deployBridgeFixture()

            await expect(
                l1Bridge.connect(owner).sendSignal(ethers.constants.HashZero)
            ).to.be.revertedWith("B:signal")
        })

        it("sends signal, confirms it was sent", async function () {
            const { owner, l1Bridge } = await deployBridgeFixture()

            const hash =
                "0xf2e08f6b93d8cf4f37a3b38f91a8c37198095dde8697463ca3789e25218a8e9d"
            await expect(l1Bridge.connect(owner).sendSignal(hash))
                .to.emit(l1Bridge, "SignalSent")
                .withArgs(owner.address, hash)

            const isSignalSent = await l1Bridge.isSignalSent(
                owner.address,
                hash
            )
            expect(isSignalSent).to.be.eq(true)
        })
    })

    describe("isDestChainEnabled()", function () {
        it("is disabled for unabled chainIds", async () => {
            const { l1Bridge } = await deployBridgeFixture()

            const enabled = await l1Bridge.isDestChainEnabled(68)
            expect(enabled).to.be.eq(false)
        })

        it("is enabled for enabled chainId", async () => {
            const { l1Bridge, enabledDestChainId } = await deployBridgeFixture()

            const enabled = await l1Bridge.isDestChainEnabled(
                enabledDestChainId
            )
            expect(enabled).to.be.eq(true)
        })
    })

    describe("context()", function () {
        it("returns unitialized context", async () => {
            const { l1Bridge } = await deployBridgeFixture()

            const ctx = await l1Bridge.context()
            expect(ctx[0]).to.be.eq(ethers.constants.HashZero)
            expect(ctx[1]).to.be.eq(ethers.constants.AddressZero)
            expect(ctx[2]).to.be.eq(BigNumber.from(0))
        })
    })

    describe("getMessageStatus()", function () {
        it("returns new for uninitialized signal", async () => {
            const { l1Bridge } = await deployBridgeFixture()

            const messageStatus = await l1Bridge.getMessageStatus(
                ethers.constants.HashZero
            )

            expect(messageStatus).to.be.eq(0)
        })

        it("returns for initiaized signal", async () => {
            const { owner, nonOwner, enabledDestChainId, l1Bridge } =
                await deployBridgeFixture()

            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 100,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const expectedAmount =
                message.depositValue + message.callValue + message.processingFee

            const tx = await l1Bridge.sendMessage(message, {
                value: expectedAmount,
            })

            const receipt = await tx.wait()

            const [messageSentEvent] = receipt.events as any as Event[]

            const { signal } = (messageSentEvent as any).args

            expect(signal).not.to.be.eq(ethers.constants.HashZero)

            const messageStatus = await l1Bridge.getMessageStatus(signal)

            expect(messageStatus).to.be.eq(0)
        })
    })

    describe("processMessage()", async function () {
        it("throws when message.gasLimit is 0 and msg.sender is not the message.owner", async () => {
            const { owner, nonOwner, l1Bridge, enabledDestChainId } =
                await deployBridgeFixture()

            const message: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: 1,
                destChainId: enabledDestChainId,
                owner: nonOwner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 0,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const proof = ethers.constants.HashZero

            await expect(
                l1Bridge.processMessage(message, proof)
            ).to.be.revertedWith("B:forbidden")
        })

        it("throws message.destChainId is not block.chainId", async () => {
            const { owner, nonOwner, l1Bridge } = await deployBridgeFixture()

            const message: Message = {
                id: 1,
                sender: nonOwner.address,
                srcChainId: 1,
                destChainId: 5,
                owner: owner.address,
                to: nonOwner.address,
                refundAddress: owner.address,
                depositValue: 1,
                callValue: 1,
                processingFee: 1,
                gasLimit: 0,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const proof = ethers.constants.HashZero

            await expect(
                l1Bridge.processMessage(message, proof)
            ).to.be.revertedWith("B:destChainId")
        })
    })
})

describe("integration:Bridge", function () {
    async function deployBridgeFixture() {
        const [owner, nonOwner] = await ethers.getSigners()

        const { chainId } = await ethers.provider.getNetwork()

        const srcChainId = chainId

        // seondary node to deploy L2 on
        const l2Provider = new ethers.providers.JsonRpcProvider(
            "http://localhost:28545"
        )

        const l2Signer = await l2Provider.getSigner(
            "0x4D9E82AC620246f6782EAaBaC3E3c86895f3f0F8"
        )

        const l2NonOwner = await l2Provider.getSigner()

        const l2Network = await l2Provider.getNetwork()
        const enabledDestChainId = l2Network.chainId

        const addressManager: AddressManager = await (
            await ethers.getContractFactory("AddressManager")
        ).deploy()
        await addressManager.init()

        const l2AddressManager: AddressManager = await (
            await ethers.getContractFactory("AddressManager")
        )
            .connect(l2Signer)
            .deploy()
        await l2AddressManager.init()

        const { bridge: l1Bridge, etherVault: l1EtherVault } =
            await deployBridge(
                owner,
                addressManager,
                enabledDestChainId,
                srcChainId
            )

        const { bridge: l2Bridge, etherVault: l2EtherVault } =
            await deployBridge(
                l2Signer,
                l2AddressManager,
                srcChainId,
                enabledDestChainId
            )

        await addressManager.setAddress(
            `${enabledDestChainId}.bridge`,
            l2Bridge.address
        )

        await l2AddressManager
            .connect(l2Signer)
            .setAddress(`${srcChainId}.bridge`, l1Bridge.address)

        const headerSync = await (
            await ethers.getContractFactory("TestHeaderSync")
        )
            .connect(l2Signer)
            .deploy()

        await l2AddressManager
            .connect(l2Signer)
            .setAddress(`${enabledDestChainId}.taiko`, headerSync.address)

        return {
            owner,
            l2Signer,
            nonOwner,
            l2NonOwner,
            l1Bridge,
            l2Bridge,
            addressManager,
            enabledDestChainId,
            l1EtherVault,
            l2EtherVault,
            srcChainId,
            headerSync,
        }
    }

    describe("processMessage()", function () {
        it("should throw if message.gasLimit == 0 & msg.sender is not message.owner", async function () {
            const {
                owner,
                l2NonOwner,
                srcChainId,
                enabledDestChainId,
                l2Bridge,
            } = await deployBridgeFixture()

            const m: Message = {
                id: 1,
                sender: await l2NonOwner.getAddress(),
                srcChainId: srcChainId,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1000,
                callValue: 1000,
                processingFee: 1000,
                gasLimit: 0,
                data: ethers.constants.HashZero,
                memo: "",
            }

            await expect(
                l2Bridge.processMessage(m, ethers.constants.HashZero)
            ).to.be.revertedWith("B:forbidden")
        })

        it("should throw if message.destChainId is not equal to current block.chainId", async function () {
            const { owner, srcChainId, enabledDestChainId, l2Bridge } =
                await deployBridgeFixture()

            const m: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: srcChainId,
                destChainId: enabledDestChainId + 1,
                owner: owner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1000,
                callValue: 1000,
                processingFee: 1000,
                gasLimit: 10000,
                data: ethers.constants.HashZero,
                memo: "",
            }

            await expect(
                l2Bridge.processMessage(m, ethers.constants.HashZero)
            ).to.be.revertedWith("B:destChainId")
        })

        it("should throw if messageStatus of message is != NEW", async function () {
            const {
                owner,
                l1Bridge,
                srcChainId,
                enabledDestChainId,
                l2Bridge,
                headerSync,
            } = await deployBridgeFixture()

            const m: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: srcChainId,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1000,
                callValue: 1000,
                processingFee: 1000,
                gasLimit: 10000,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const expectedAmount =
                m.depositValue + m.callValue + m.processingFee
            const tx = await l1Bridge.sendMessage(m, {
                value: expectedAmount,
            })

            const receipt = await tx.wait()

            const [messageSentEvent] = receipt.events as any as Event[]

            const { signal, message } = (messageSentEvent as any).args

            const sender = l1Bridge.address

            const key = ethers.utils.keccak256(
                ethers.utils.solidityPack(
                    ["address", "bytes32"],
                    [sender, signal]
                )
            )

            // use this instead of ethers.provider.getBlock() beccause it doesnt have stateRoot
            // in the response
            const block: Block = await ethers.provider.send(
                "eth_getBlockByNumber",
                ["latest", false]
            )

            await headerSync.setSyncedHeader(block.hash)

            const logsBloom = block.logsBloom.toString().substring(2)

            const blockHeader: BlockHeader = {
                parentHash: block.parentHash,
                ommersHash: block.sha3Uncles,
                beneficiary: block.miner,
                stateRoot: block.stateRoot,
                transactionsRoot: block.transactionsRoot,
                receiptsRoot: block.receiptsRoot,
                logsBloom: logsBloom
                    .match(/.{1,64}/g)!
                    .map((s: string) => "0x" + s),
                difficulty: block.difficulty,
                height: block.number,
                gasLimit: block.gasLimit,
                gasUsed: block.gasUsed,
                timestamp: block.timestamp,
                extraData: block.extraData,
                mixHash: block.mixHash,
                nonce: block.nonce,
                baseFeePerGas: block.baseFeePerGas
                    ? parseInt(block.baseFeePerGas)
                    : 0,
            }

            // rpc call to get the merkle proof what value is at key on the bridge contract
            const proof: EthGetProofResponse = await ethers.provider.send(
                "eth_getProof",
                [l1Bridge.address, [key], block.hash]
            )

            // RLP encode the proof together for LibTrieProof to decode
            const encodedProof = ethers.utils.defaultAbiCoder.encode(
                ["bytes", "bytes"],
                [
                    RLP.encode(proof.accountProof),
                    RLP.encode(proof.storageProof[0].proof),
                ]
            )
            // encode the SignalProof struct from LibBridgeSignal
            const signalProof = ethers.utils.defaultAbiCoder.encode(
                [
                    "tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas) header, bytes proof)",
                ],
                [{ header: blockHeader, proof: encodedProof }]
            )

            // upon successful processing, this immediately gets marked as DONE
            await l2Bridge.processMessage(message, signalProof)

            // recalling this process should be prevented as it's status is no longer NEW
            await expect(
                l2Bridge.processMessage(message, signalProof)
            ).to.be.revertedWith("B:status")
        })

        it("should throw if message signalproof is not valid", async function () {
            const {
                owner,
                l1Bridge,
                srcChainId,
                enabledDestChainId,
                l2Bridge,
                headerSync,
            } = await deployBridgeFixture()

            const m: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: srcChainId,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1000,
                callValue: 1000,
                processingFee: 1000,
                gasLimit: 10000,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const block: Block = await ethers.provider.send(
                "eth_getBlockByNumber",
                ["latest", false]
            )

            const libData = await (
                await ethers.getContractFactory("TestLibBridgeData")
            ).deploy()

            const signal = await libData.hashMessage(m)

            const sender = l1Bridge.address

            const key = ethers.utils.keccak256(
                ethers.utils.solidityPack(
                    ["address", "bytes32"],
                    [sender, signal]
                )
            )

            await headerSync.setSyncedHeader(ethers.constants.HashZero)

            const logsBloom = block.logsBloom.toString().substring(2)

            const blockHeader: BlockHeader = {
                parentHash: block.parentHash,
                ommersHash: block.sha3Uncles,
                beneficiary: block.miner,
                stateRoot: block.stateRoot,
                transactionsRoot: block.transactionsRoot,
                receiptsRoot: block.receiptsRoot,
                logsBloom: logsBloom
                    .match(/.{1,64}/g)!
                    .map((s: string) => "0x" + s),
                difficulty: block.difficulty,
                height: block.number,
                gasLimit: block.gasLimit,
                gasUsed: block.gasUsed,
                timestamp: block.timestamp,
                extraData: block.extraData,
                mixHash: block.mixHash,
                nonce: block.nonce,
                baseFeePerGas: block.baseFeePerGas
                    ? parseInt(block.baseFeePerGas)
                    : 0,
            }

            const proof: EthGetProofResponse = await ethers.provider.send(
                "eth_getProof",
                [l1Bridge.address, [key], block.hash]
            )

            // RLP encode the proof together for LibTrieProof to decode
            const encodedProof = ethers.utils.defaultAbiCoder.encode(
                ["bytes", "bytes"],
                [
                    RLP.encode(proof.accountProof),
                    RLP.encode(proof.storageProof[0].proof),
                ]
            )
            // encode the SignalProof struct from LibBridgeSignal
            const signalProof = ethers.utils.defaultAbiCoder.encode(
                [
                    "tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas) header, bytes proof)",
                ],
                [{ header: blockHeader, proof: encodedProof }]
            )

            await expect(
                l2Bridge.processMessage(m, signalProof)
            ).to.be.revertedWith("LTP:invalid storage proof")
        })

        it("should throw if message has not been received", async function () {
            const {
                owner,
                l1Bridge,
                srcChainId,
                enabledDestChainId,
                l2Bridge,
                headerSync,
            } = await deployBridgeFixture()

            const m: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: srcChainId,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1000,
                callValue: 1000,
                processingFee: 1000,
                gasLimit: 10000,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const expectedAmount =
                m.depositValue + m.callValue + m.processingFee
            const tx = await l1Bridge.sendMessage(m, {
                value: expectedAmount,
            })

            const receipt = await tx.wait()

            const [messageSentEvent] = receipt.events as any as Event[]

            const { signal, message } = (messageSentEvent as any).args

            expect(signal).not.to.be.eq(ethers.constants.HashZero)

            const messageStatus = await l1Bridge.getMessageStatus(signal)

            expect(messageStatus).to.be.eq(0)

            const sender = l1Bridge.address

            const key = ethers.utils.keccak256(
                ethers.utils.solidityPack(
                    ["address", "bytes32"],
                    [sender, signal]
                )
            )

            // use this instead of ethers.provider.getBlock() beccause it doesnt have stateRoot
            // in the response
            const block: Block = await ethers.provider.send(
                "eth_getBlockByNumber",
                ["latest", false]
            )

            await headerSync.setSyncedHeader(ethers.constants.HashZero)

            const logsBloom = block.logsBloom.toString().substring(2)

            const blockHeader: BlockHeader = {
                parentHash: block.parentHash,
                ommersHash: block.sha3Uncles,
                beneficiary: block.miner,
                stateRoot: block.stateRoot,
                transactionsRoot: block.transactionsRoot,
                receiptsRoot: block.receiptsRoot,
                logsBloom: logsBloom
                    .match(/.{1,64}/g)!
                    .map((s: string) => "0x" + s),
                difficulty: block.difficulty,
                height: block.number,
                gasLimit: block.gasLimit,
                gasUsed: block.gasUsed,
                timestamp: block.timestamp,
                extraData: block.extraData,
                mixHash: block.mixHash,
                nonce: block.nonce,
                baseFeePerGas: block.baseFeePerGas
                    ? parseInt(block.baseFeePerGas)
                    : 0,
            }

            // get storageValue for the key
            const storageValue = await ethers.provider.getStorageAt(
                l1Bridge.address,
                key,
                block.number
            )
            // make sure it equals 1 so our proof will pass
            expect(storageValue).to.be.eq(
                "0x0000000000000000000000000000000000000000000000000000000000000001"
            )
            // rpc call to get the merkle proof what value is at key on the bridge contract
            const proof: EthGetProofResponse = await ethers.provider.send(
                "eth_getProof",
                [l1Bridge.address, [key], block.hash]
            )

            // RLP encode the proof together for LibTrieProof to decode
            const encodedProof = ethers.utils.defaultAbiCoder.encode(
                ["bytes", "bytes"],
                [
                    RLP.encode(proof.accountProof),
                    RLP.encode(proof.storageProof[0].proof),
                ]
            )
            // encode the SignalProof struct from LibBridgeSignal
            const signalProof = ethers.utils.defaultAbiCoder.encode(
                [
                    "tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas) header, bytes proof)",
                ],
                [{ header: blockHeader, proof: encodedProof }]
            )

            await expect(
                l2Bridge.processMessage(message, signalProof)
            ).to.be.revertedWith("B:notReceived")
        })

        it("processes a message when the signal has been verified from the sending chain", async () => {
            const {
                owner,
                l1Bridge,
                srcChainId,
                enabledDestChainId,
                l2Bridge,
                headerSync,
            } = await deployBridgeFixture()

            const m: Message = {
                id: 1,
                sender: owner.address,
                srcChainId: srcChainId,
                destChainId: enabledDestChainId,
                owner: owner.address,
                to: owner.address,
                refundAddress: owner.address,
                depositValue: 1000,
                callValue: 1000,
                processingFee: 1000,
                gasLimit: 10000,
                data: ethers.constants.HashZero,
                memo: "",
            }

            const expectedAmount =
                m.depositValue + m.callValue + m.processingFee
            const tx = await l1Bridge.sendMessage(m, {
                value: expectedAmount,
            })

            const receipt = await tx.wait()

            const [messageSentEvent] = receipt.events as any as Event[]

            const { signal, message } = (messageSentEvent as any).args

            expect(signal).not.to.be.eq(ethers.constants.HashZero)

            const messageStatus = await l1Bridge.getMessageStatus(signal)

            expect(messageStatus).to.be.eq(0)

            const sender = l1Bridge.address

            const key = ethers.utils.keccak256(
                ethers.utils.solidityPack(
                    ["address", "bytes32"],
                    [sender, signal]
                )
            )

            // use this instead of ethers.provider.getBlock() beccause it doesnt have stateRoot
            // in the response
            const block: Block = await ethers.provider.send(
                "eth_getBlockByNumber",
                ["latest", false]
            )

            await headerSync.setSyncedHeader(block.hash)

            const logsBloom = block.logsBloom.toString().substring(2)

            const blockHeader: BlockHeader = {
                parentHash: block.parentHash,
                ommersHash: block.sha3Uncles,
                beneficiary: block.miner,
                stateRoot: block.stateRoot,
                transactionsRoot: block.transactionsRoot,
                receiptsRoot: block.receiptsRoot,
                logsBloom: logsBloom
                    .match(/.{1,64}/g)!
                    .map((s: string) => "0x" + s),
                difficulty: block.difficulty,
                height: block.number,
                gasLimit: block.gasLimit,
                gasUsed: block.gasUsed,
                timestamp: block.timestamp,
                extraData: block.extraData,
                mixHash: block.mixHash,
                nonce: block.nonce,
                baseFeePerGas: block.baseFeePerGas
                    ? parseInt(block.baseFeePerGas)
                    : 0,
            }

            // get storageValue for the key
            const storageValue = await ethers.provider.getStorageAt(
                l1Bridge.address,
                key,
                block.number
            )
            // make sure it equals 1 so our proof will pass
            expect(storageValue).to.be.eq(
                "0x0000000000000000000000000000000000000000000000000000000000000001"
            )
            // rpc call to get the merkle proof what value is at key on the bridge contract
            const proof: EthGetProofResponse = await ethers.provider.send(
                "eth_getProof",
                [l1Bridge.address, [key], block.hash]
            )

            // RLP encode the proof together for LibTrieProof to decode
            const encodedProof = ethers.utils.defaultAbiCoder.encode(
                ["bytes", "bytes"],
                [
                    RLP.encode(proof.accountProof),
                    RLP.encode(proof.storageProof[0].proof),
                ]
            )
            // encode the SignalProof struct from LibBridgeSignal
            const signalProof = ethers.utils.defaultAbiCoder.encode(
                [
                    "tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas) header, bytes proof)",
                ],
                [{ header: blockHeader, proof: encodedProof }]
            )

            // ROGER: this is where we at, we need now to deploy a custom TestHeaderSync that implements
            // IHeaderSync where we can manually save synced headers.
            await l2Bridge.processMessage(message, signalProof, {
                gasLimit: BigNumber.from(2000000),
            })
        })
    })
})
