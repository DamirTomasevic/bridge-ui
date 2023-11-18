// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../TestBase.sol";
import "../../contracts/common/AddressManager.sol";
import "../../contracts/common/AddressResolver.sol";
import "../../contracts/L1/TaikoToken.sol";

contract TaikoTokenTest is TaikoTest {
    bytes32 GENESIS_BLOCK_HASH;

    address public tokenOwner;

    AddressManager public addressManager;
    TransparentUpgradeableProxy public tokenProxy;
    TaikoToken public tko;
    TaikoToken public tkoUpgradedImpl;

    function setUp() public {
        GENESIS_BLOCK_HASH = getRandomBytes32();

        tokenOwner = getRandomAddress();

        addressManager = new AddressManager();
        addressManager.init();
        tko = new TaikoToken();

        tokenProxy = _deployViaProxy(
            address(tko),
            bytes.concat(
                tko.init.selector,
                abi.encode(address(addressManager), "Taiko Token", "TKO", address(this))
            )
        );

        tko = TaikoToken(address(tokenProxy));
        tko.transfer(Yasmine, 5 ether);
        tko.transfer(Zachary, 5 ether);
    }

    function test_TaikoToken_upgrade() public {
        tkoUpgradedImpl = new TaikoToken();

        vm.prank(tokenOwner);
        tokenProxy.upgradeTo(address(tkoUpgradedImpl));

        // Check if balance is still same
        assertEq(tko.balanceOf(Yasmine), 5 ether);
        assertEq(tko.balanceOf(Zachary), 5 ether);
    }

    function test_TaikoToken_upgrade_without_admin_rights() public {
        tkoUpgradedImpl = new TaikoToken();

        vm.expectRevert();
        tokenProxy.upgradeTo(address(tkoUpgradedImpl));
    }

    function _registerAddress(bytes32 nameHash, address addr) private {
        addressManager.setAddress(uint64(block.chainid), nameHash, addr);
    }

    function _deployViaProxy(
        address implementation,
        bytes memory data
    )
        private
        returns (TransparentUpgradeableProxy)
    {
        return new TransparentUpgradeableProxy(
            implementation,
            tokenOwner,
            data
        );
    }
}
