import { expect } from "chai"
import * as hre from "hardhat"

const ethers = hre.ethers
const action = process.env.TEST_L2_GENESIS ? describe : describe.skip

action("Generate Genesis", function () {
    let alloc: any = null

    if (process.env.TEST_L2_GENESIS) {
        alloc = require("../../deployments/genesis_alloc.json")
    }

    const provider = new hre.ethers.providers.JsonRpcProvider(
        "http://localhost:18545"
    )

    const signer = new hre.ethers.Wallet(
        "2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200",
        provider
    )

    const testConfig = require("./test_config")

    const seedAccounts = testConfig.seedAccounts

    before(async () => {
        let retry = 0

        while (true) {
            try {
                const network = await provider.getNetwork()
                if (network.chainId) break
            } catch (_) {}

            if (++retry > 10) {
                throw new Error("geth initializing timeout")
            }

            await sleep(1000)
        }

        console.log("geth initialized")
    })

    it("contracts should be deployed", async function () {
        for (const address of Object.keys(alloc)) {
            if (
                seedAccounts
                    .map((seedAccount: any) => {
                        const accountAddress = Object.keys(seedAccount)[0]
                        return accountAddress
                    })
                    .includes(address)
            ) {
                continue
            }
            const code: string = await provider.getCode(address)
            const expectCode: string = alloc[address].code

            expect(code.toLowerCase()).to.be.equal(expectCode.toLowerCase())
        }
    })

    it("premint ETH should be allocated", async function () {
        let etherVaultBalance = hre.ethers.BigNumber.from("2").pow(128).sub(1) // MaxUint128

        for (const seedAccount of seedAccounts) {
            const accountAddress = Object.keys(seedAccount)[0]
            const balance = hre.ethers.utils.parseEther(
                `${Object.values(seedAccount)[0]}`
            )
            expect(await provider.getBalance(accountAddress)).to.be.equal(
                balance.toHexString()
            )

            etherVaultBalance = etherVaultBalance.sub(balance)
        }

        const etherVaultAddress = getContractAlloc("EtherVault").address

        expect(await provider.getBalance(etherVaultAddress)).to.be.equal(
            etherVaultBalance.toHexString()
        )
    })

    describe("contracts can be called normally", function () {
        it("AddressManager", async function () {
            const addressManagerAlloc = getContractAlloc("AddressManager")

            const addressManager = new hre.ethers.Contract(
                addressManagerAlloc.address,
                require("../../artifacts/contracts/thirdparty/AddressManager.sol/AddressManager.json").abi,
                provider
            )

            const owner = await addressManager.owner()

            expect(owner).to.be.equal(testConfig.contractOwner)

            const bridge = await addressManager.getAddress(
                `${testConfig.chainId}.bridge`
            )

            expect(bridge).to.be.equal(getContractAlloc("Bridge").address)

            const tokenValut = await addressManager.getAddress(
                `${testConfig.chainId}.token_vault`
            )

            expect(tokenValut).to.be.equal(
                getContractAlloc("TokenVault").address
            )

            const etherVault = await addressManager.getAddress(
                `${testConfig.chainId}.ether_vault`
            )

            expect(etherVault).to.be.equal(
                getContractAlloc("EtherVault").address
            )

            const v1TaikoL2 = await addressManager.getAddress(
                `${testConfig.chainId}.taiko`
            )

            expect(v1TaikoL2).to.be.equal(getContractAlloc("V1TaikoL2").address)
        })

        it("LibTxDecoder", async function () {
            const LibTxDecoderAlloc = getContractAlloc("LibTxDecoder")

            const LibTxDecoder = new hre.ethers.Contract(
                LibTxDecoderAlloc.address,
                require("../../artifacts/contracts/libs/LibTxDecoder.sol/LibTxDecoder.json").abi,
                signer
            )

            await expect(
                LibTxDecoder.decodeTxList(ethers.utils.RLP.encode([]))
            ).to.be.revertedWith("empty txList")
        })

        it("V1TaikoL2", async function () {
            const V1TaikoL2Alloc = getContractAlloc("V1TaikoL2")

            const V1TaikoL2 = new hre.ethers.Contract(
                V1TaikoL2Alloc.address,
                require("../../artifacts/contracts/L2/V1TaikoL2.sol/V1TaikoL2.json").abi,
                signer
            )

            let latestL1Height = 1
            for (let i = 0; i < 300; i++) {
                const tx = await V1TaikoL2.anchor(
                    latestL1Height++,
                    ethers.utils.hexlify(ethers.utils.randomBytes(32)),
                    { gasLimit: 1000000 }
                )

                const receipt = await tx.wait()

                expect(receipt.status).to.be.equal(1)

                if (i === 299) {
                    console.log({
                        message:
                            "V1TaikoL2.anchor gas cost after 256 L2 blocks",
                        gasUsed: receipt.gasUsed,
                    })
                }
            }
        })

        it("Bridge", async function () {
            const BridgeAlloc = getContractAlloc("Bridge")
            const Bridge = new hre.ethers.Contract(
                BridgeAlloc.address,
                require("../../artifacts/contracts/bridge/Bridge.sol/Bridge.json").abi,
                signer
            )

            const owner = await Bridge.owner()

            expect(owner).to.be.equal(testConfig.contractOwner)

            await expect(Bridge.enableDestChain(1, true)).not.to.reverted
        })

        it("TokenVault", async function () {
            const TokenVaultAlloc = getContractAlloc("TokenVault")
            const TokenVault = new hre.ethers.Contract(
                TokenVaultAlloc.address,
                require("../../artifacts/contracts/bridge/TokenVault.sol/TokenVault.json").abi,
                signer
            )

            const owner = await TokenVault.owner()

            expect(owner).to.be.equal(testConfig.contractOwner)

            await expect(
                TokenVault.sendEther(
                    1,
                    ethers.Wallet.createRandom().address,
                    100,
                    0,
                    ethers.Wallet.createRandom().address,
                    "memo",
                    {
                        gasLimit: 10000000,
                        value: hre.ethers.utils.parseEther("100"),
                    }
                )
            ).to.emit(TokenVault, "EtherSent")
        })

        it("EtherVault", async function () {
            const EtherVault = new hre.ethers.Contract(
                getContractAlloc("EtherVault").address,
                require("../../artifacts/contracts/bridge/EtherVault.sol/EtherVault.json").abi,
                signer
            )

            const owner = await EtherVault.owner()

            expect(owner).to.be.equal(testConfig.contractOwner)

            expect(
                await EtherVault.isAuthorized(
                    getContractAlloc("Bridge").address
                )
            ).to.be.true

            expect(
                await EtherVault.isAuthorized(
                    ethers.Wallet.createRandom().address
                )
            ).to.be.false
        })

        it("ERC20", async function () {
            const ERC20 = new hre.ethers.Contract(
                getContractAlloc("TestERC20").address,
                require("../../artifacts/contracts/test/TestERC20.sol/TestERC20.json").abi,
                signer
            )

            const {
                TOKEN_NAME,
                TOKEN_SYMBOL,
                PREMINT_SEED_ACCOUNT_BALANCE,
            } = require("../../utils/generate_genesis/erc20")

            expect(await ERC20.name()).to.be.equal(TOKEN_NAME)
            expect(await ERC20.symbol()).to.be.equal(TOKEN_SYMBOL)

            for (const seedAccount of seedAccounts) {
                const accountAddress = Object.keys(seedAccount)[0]

                expect(await ERC20.balanceOf(accountAddress)).to.be.equal(
                    PREMINT_SEED_ACCOUNT_BALANCE
                )
            }

            expect(await ERC20.totalSupply()).to.be.equal(
                seedAccounts.length * PREMINT_SEED_ACCOUNT_BALANCE
            )

            await expect(
                ERC20.transfer(ethers.Wallet.createRandom().address, 1)
            ).to.emit(ERC20, "Transfer")
        })
    })

    function getContractAlloc(name: string): any {
        for (const address of Object.keys(alloc)) {
            if (alloc[address].contractName === name) {
                return Object.assign(alloc[address], { address })
            }
        }

        throw new Error(`contract alloc: ${name} not found`)
    }
})

function sleep(ms: number) {
    return new Promise((resolve) => {
        setTimeout(resolve, ms)
    })
}
