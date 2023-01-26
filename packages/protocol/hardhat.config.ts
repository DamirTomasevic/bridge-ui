import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@openzeppelin/hardhat-upgrades";
import "@typechain/hardhat";
import "hardhat-abi-exporter";
import "hardhat-gas-reporter";
import { HardhatUserConfig } from "hardhat/config";
import "solidity-coverage";
import "solidity-docgen";
import "./tasks/deploy_L1";
import "./tasks/compile_yul";

const hardhatMnemonic =
    "test test test test test test test test test test test taik";
const config: HardhatUserConfig = {
    docgen: {
        exclude: [
            "bridge/libs/",
            "L1/libs/",
            "libs/",
            "test/",
            "thirdparty/",
            "common/EssentialContract.sol",
        ],
        outputDir: "../website/pages/docs/contract-documentation/",
        pages: "files",
        templates: "./solidity-docgen/templates",
    },
    gasReporter: {
        currency: "USD",
        enabled: process.env.REPORT_GAS === "true",
    },
    mocha: {
        timeout: 300000,
    },
    networks: {
        goerli: {
            accounts:
                process.env.PRIVATE_KEY !== undefined
                    ? [process.env.PRIVATE_KEY]
                    : [],
            url: process.env.GOERLI_URL || "",
        },
        hardhat: {
            accounts: {
                mnemonic: hardhatMnemonic,
            },
            gas: 8000000,
        },
        mainnet: {
            accounts:
                process.env.PRIVATE_KEY !== undefined
                    ? [process.env.PRIVATE_KEY]
                    : [],
            url: process.env.MAINNET_URL || "",
        },
        l1_test: {
            accounts:
                process.env.PRIVATE_KEY !== undefined
                    ? [process.env.PRIVATE_KEY]
                    : { mnemonic: hardhatMnemonic },
            url: "http://127.0.0.1:18545" || "",
        },
        l2_test: {
            accounts:
                process.env.PRIVATE_KEY !== undefined
                    ? [process.env.PRIVATE_KEY]
                    : { mnemonic: hardhatMnemonic },
            url: "http://127.0.0.1:28545" || "",
        },
        localhost: {
            accounts:
                process.env.PRIVATE_KEY !== undefined
                    ? [process.env.PRIVATE_KEY]
                    : [],
            url: "http://127.0.0.1:8545" || "",
        },
        ropsten: {
            accounts:
                process.env.PRIVATE_KEY !== undefined
                    ? [process.env.PRIVATE_KEY]
                    : [],
            url: process.env.ROPSTEN_URL || "",
        },
    },
    solidity: {
        settings: {
            optimizer: {
                enabled: true,
                runs: 10000,
            },
            outputSelection: {
                "*": {
                    "*": ["storageLayout"],
                },
            },
        },
        version: "0.8.9",
    },
};

export default config;
