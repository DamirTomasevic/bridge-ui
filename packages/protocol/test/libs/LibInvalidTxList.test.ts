import { expect } from "chai"
import { UnsignedTransaction } from "ethers"
import { ethers } from "hardhat"

describe("LibInvalidTxList", function () {
    let libInvalidTxList: any
    let libRLPWriter: any
    let libRLPReader: any
    let LibConstants: any
    let testUnsignedTxs: Array<UnsignedTransaction>
    let chainId: any

    const signingKey = new ethers.utils.SigningKey(ethers.utils.randomBytes(32))
    const signerAddress = new ethers.Wallet(signingKey.privateKey).address

    before(async function () {
        LibConstants = await (
            await ethers.getContractFactory("LibConstants")
        ).deploy()

        libInvalidTxList = await (
            await ethers.getContractFactory("TestLibInvalidTxList")
        ).deploy()

        libRLPReader = await (
            await ethers.getContractFactory("TestLib_RLPReader")
        ).deploy()

        libRLPWriter = await (
            await ethers.getContractFactory("TestLib_RLPWriter")
        ).deploy()

        chainId = (await LibConstants.TAIKO_CHAIN_ID()).toNumber()

        const unsignedLegacyTx: UnsignedTransaction = {
            type: 0,
            // if chainId is defined, ether.js will automatically use EIP-155
            // signature
            chainId,
            nonce: Math.floor(Math.random() * 1024),
            gasPrice: randomBigInt(),
            gasLimit: randomBigInt(),
            to: ethers.Wallet.createRandom().address,
            value: randomBigInt(),
            data: ethers.utils.randomBytes(32),
        }

        const unsigned2930Tx: UnsignedTransaction = {
            type: 1,
            chainId,
            nonce: Math.floor(Math.random() * 1024),
            gasPrice: randomBigInt(),
            gasLimit: randomBigInt(),
            to: ethers.Wallet.createRandom().address,
            value: randomBigInt(),
            accessList: [
                [
                    ethers.Wallet.createRandom().address,
                    [ethers.utils.hexlify(ethers.utils.randomBytes(32))],
                ],
            ],
            data: ethers.utils.randomBytes(32),
        }

        const unsigned1559Tx: UnsignedTransaction = {
            type: 2,
            chainId,
            nonce: Math.floor(Math.random() * 1024),
            maxPriorityFeePerGas: randomBigInt(),
            maxFeePerGas: randomBigInt(),
            gasLimit: randomBigInt(),
            to: ethers.Wallet.createRandom().address,
            value: randomBigInt(),
            accessList: [
                [
                    ethers.Wallet.createRandom().address,
                    [ethers.utils.hexlify(ethers.utils.randomBytes(32))],
                ],
            ],
            data: ethers.utils.randomBytes(32),
        }

        testUnsignedTxs = [unsignedLegacyTx, unsigned2930Tx, unsigned1559Tx]
    })

    it("should parse the recover payloads correctly", async function () {
        for (const unsignedTx of testUnsignedTxs) {
            const expectedHash = ethers.utils.keccak256(
                ethers.utils.serializeTransaction(unsignedTx)
            )

            const signature = signingKey.signDigest(expectedHash)
            const { v: expectedV, r: expectedR, s: expectedS } = signature

            const [hash, v, r, s] = await libInvalidTxList.parseRecoverPayloads(
                {
                    txType: unsignedTx.type,
                    destination: unsignedTx.to,
                    data: unsignedTx.data,
                    gasLimit: unsignedTx.gasLimit,
                    txData: ethers.utils.serializeTransaction(
                        unsignedTx,
                        signature
                    ),
                }
            )

            expect(hash).to.be.equal(expectedHash)
            expect(v).to.be.equal(expectedV)
            expect(r).to.be.equal(expectedR)
            expect(s).to.be.equal(expectedS)
        }
    })

    it("should verify valid transaction signatures", async function () {
        for (const unsignedTx of testUnsignedTxs) {
            const expectedHash = ethers.utils.keccak256(
                ethers.utils.serializeTransaction(unsignedTx)
            )
            const signature = signingKey.signDigest(expectedHash)

            expect(
                await libInvalidTxList.verifySignature({
                    txType: unsignedTx.type,
                    destination: unsignedTx.to,
                    data: unsignedTx.data,
                    gasLimit: unsignedTx.gasLimit,
                    txData: ethers.utils.serializeTransaction(
                        unsignedTx,
                        signature
                    ),
                })
            ).to.be.equal(signerAddress)
        }
    })

    it("should verify invalid transaction signatures", async function () {
        for (const unsignedTx of testUnsignedTxs) {
            const expectedHash = ethers.utils.keccak256(
                ethers.utils.serializeTransaction(unsignedTx)
            )
            const signature = signingKey.signDigest(expectedHash)

            const randomV =
                unsignedTx.type === 0
                    ? Math.floor(Math.random() * Math.pow(2, 7)) +
                      2 * chainId +
                      35
                    : Math.floor(Math.random() * Math.pow(2, 7))

            const randomSignature = {
                v: randomV,
                r: ethers.utils.hexlify(ethers.utils.randomBytes(32)),
                s: ethers.utils.hexlify(ethers.utils.randomBytes(32)),
            }

            const txData = await changeSignature(
                unsignedTx.type,
                ethers.utils.arrayify(
                    ethers.utils.serializeTransaction(unsignedTx, signature)
                ),
                randomSignature
            )

            expect(
                await libInvalidTxList.verifySignature({
                    txType: unsignedTx.type,
                    destination: unsignedTx.to,
                    data: unsignedTx.data,
                    gasLimit: unsignedTx.gasLimit,
                    txData,
                })
            ).to.be.equal(ethers.constants.AddressZero)
        }
    })

    async function changeSignature(
        type: any,
        encoded: Uint8Array,
        signature: any
    ) {
        if (type !== 0) encoded = encoded.slice(1)

        const rlpItemsList = (await libRLPReader.readList(encoded)).slice(0, -3)

        let result = await libRLPWriter.writeList(
            rlpItemsList.concat([
                await libRLPWriter.writeUint(signature.v),
                await libRLPWriter.writeBytes(signature.r),
                await libRLPWriter.writeBytes(signature.s),
            ])
        )

        if (type !== 0) result = ethers.utils.concat([[type], result])

        return ethers.utils.hexlify(result)
    }

    function randomBigInt() {
        return ethers.BigNumber.from(ethers.utils.randomBytes(32))
    }
})
