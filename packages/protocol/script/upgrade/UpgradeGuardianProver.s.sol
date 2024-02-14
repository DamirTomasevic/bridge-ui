// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/
//
//   Email: security@taiko.xyz
//   Website: https://taiko.xyz
//   GitHub: https://github.com/taikoxyz
//   Discord: https://discord.gg/taikoxyz
//   Twitter: https://twitter.com/taikoxyz
//   Blog: https://mirror.xyz/labs.taiko.eth
//   Youtube: https://www.youtube.com/@taikoxyz

pragma solidity ^0.8.24;

import "forge-std/src/Script.sol";
import "forge-std/src/console2.sol";
import "../../contracts/L1/provers/GuardianProver.sol";
import "./UpgradeScript.s.sol";

contract UpgradeGuardianProver is UpgradeScript {
    function run() external setUp {
        console2.log("upgrading GuardianProver");
        GuardianProver newGuardianProver = new GuardianProver();
        upgrade(address(newGuardianProver));

        console2.log("upgraded GuardianProver to", address(newGuardianProver));
    }
}
