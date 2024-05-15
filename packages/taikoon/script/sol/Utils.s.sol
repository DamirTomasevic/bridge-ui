// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { Script, console } from "forge-std/src/Script.sol";
import "forge-std/src/StdJson.sol";

contract UtilsScript is Script {
    using stdJson for string;

    address public nounsTokenAddress;

    uint256 public chainId;

    string public lowercaseNetworkKey;
    string public uppercaseNetworkKey;

    function setUp() public {
        // load all network configs
        chainId = block.chainid;

        if (chainId == 31_337) {
            lowercaseNetworkKey = "localhost";
            uppercaseNetworkKey = "LOCALHOST";
        } else if (chainId == 17_000) {
            lowercaseNetworkKey = "holesky";
            uppercaseNetworkKey = "HOLESKY";
        } else if (chainId == 167_001) {
            lowercaseNetworkKey = "devnet";
            uppercaseNetworkKey = "DEVNET";
        } else if (chainId == 167_008) {
            lowercaseNetworkKey = "katla";
            uppercaseNetworkKey = "KATLA";
        } else {
            revert("Unsupported chainId");
        }
    }

    function getPrivateKey() public view returns (uint256) {
        string memory lookupKey = string.concat(uppercaseNetworkKey, "_PRIVATE_KEY");
        return vm.envUint(lookupKey);
    }

    function getAddress() public view returns (address) {
        string memory lookupKey = string.concat(uppercaseNetworkKey, "_ADDRESS");
        return vm.envAddress(lookupKey);
    }

    function getContractJsonLocation() public view returns (string memory) {
        string memory root = vm.projectRoot();
        return string.concat(root, "/deployments/", lowercaseNetworkKey, ".json");
    }

    function getIpfsBaseURI() public view returns (string memory) {
        return vm.envString("IPFS_BASE_URI");
    }

    function run() public { }
}
