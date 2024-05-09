// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../common/TestMulticall3.sol";
import "../TaikoTest.sol";

contract Target is EssentialContract {
    function init(address _owner) external initializer {
        __Essential_init(_owner);
    }
}

contract TestDelegateOwner is TaikoTest {
    address public owner;
    address public remoteOwner;
    Bridge public bridge;
    SignalService public signalService;
    AddressManager public addressManager;
    DelegateOwner public delegateOwner;
    TestMulticall3 public multicall;

    uint64 remoteChainId = uint64(block.chainid + 1);
    address remoteBridge = vm.addr(0x2000);

    function setUp() public {
        owner = vm.addr(0x1000);
        vm.deal(owner, 100 ether);

        remoteOwner = vm.addr(0x2000);

        vm.startPrank(owner);

        multicall = new TestMulticall3();

        addressManager = AddressManager(
            deployProxy({
                name: "address_manager",
                impl: address(new AddressManager()),
                data: abi.encodeCall(AddressManager.init, (address(0)))
            })
        );

        delegateOwner = DelegateOwner(
            deployProxy({
                name: "delegate_owner",
                impl: address(new DelegateOwner()),
                data: abi.encodeCall(
                    DelegateOwner.init, (remoteOwner, address(addressManager), remoteChainId)
                ),
                registerTo: address(addressManager)
            })
        );

        signalService = SkipProofCheckSignal(
            deployProxy({
                name: "signal_service",
                impl: address(new SkipProofCheckSignal()),
                data: abi.encodeCall(SignalService.init, (address(0), address(addressManager))),
                registerTo: address(addressManager)
            })
        );

        bridge = Bridge(
            payable(
                deployProxy({
                    name: "bridge",
                    impl: address(new Bridge()),
                    data: abi.encodeCall(Bridge.init, (address(0), address(addressManager))),
                    registerTo: address(addressManager)
                })
            )
        );

        addressManager.setAddress(remoteChainId, "bridge", remoteBridge);
        vm.stopPrank();
    }

    function test_delegate_owner_single_non_delegatecall() public {
        Target target1 = Target(
            deployProxy({
                name: "target1",
                impl: address(new Target()),
                data: abi.encodeCall(Target.init, (address(delegateOwner)))
            })
        );

        bytes memory data = abi.encode(
            DelegateOwner.Call(
                uint64(0),
                address(target1),
                false, // CALL
                abi.encodeCall(EssentialContract.pause, ())
            )
        );

        vm.expectRevert(DelegateOwner.DO_DRYRUN_SUCCEEDED.selector);
        delegateOwner.dryrunMessageInvocation(data);

        IBridge.Message memory message;
        message.from = remoteOwner;
        message.destChainId = uint64(block.chainid);
        message.srcChainId = remoteChainId;
        message.destOwner = Bob;
        message.data = abi.encodeCall(DelegateOwner.onMessageInvocation, (data));
        message.to = address(delegateOwner);

        vm.prank(Bob);
        bridge.processMessage(message, "");

        bytes32 hash = bridge.hashMessage(message);
        assertTrue(bridge.messageStatus(hash) == IBridge.Status.DONE);

        assertEq(delegateOwner.nextTxId(), 1);
        assertTrue(target1.paused());
    }

    function test_delegate_owner_single_non_delegatecall_self() public {
        address delegateOwnerImpl2 = address(new DelegateOwner());

        bytes memory data = abi.encode(
            DelegateOwner.Call(
                uint64(0),
                address(delegateOwner),
                false, // CALL
                abi.encodeCall(UUPSUpgradeable.upgradeTo, (delegateOwnerImpl2))
            )
        );

        vm.expectRevert(DelegateOwner.DO_DRYRUN_SUCCEEDED.selector);
        delegateOwner.dryrunMessageInvocation(data);

        IBridge.Message memory message;
        message.from = remoteOwner;
        message.destChainId = uint64(block.chainid);
        message.srcChainId = remoteChainId;
        message.destOwner = Bob;
        message.data = abi.encodeCall(DelegateOwner.onMessageInvocation, (data));
        message.to = address(delegateOwner);

        vm.prank(Bob);
        bridge.processMessage(message, "");

        bytes32 hash = bridge.hashMessage(message);
        assertTrue(bridge.messageStatus(hash) == IBridge.Status.DONE);

        assertEq(delegateOwner.nextTxId(), 1);
        assertEq(delegateOwner.impl(), delegateOwnerImpl2);
    }

    function test_delegate_owner_delegate_multicall() public {
        address impl1 = address(new Target());
        address impl2 = address(new Target());

        address delegateOwnerImpl2 = address(new DelegateOwner());

        Target target1 = Target(
            deployProxy({
                name: "target1",
                impl: impl1,
                data: abi.encodeCall(Target.init, (address(delegateOwner)))
            })
        );
        Target target2 = Target(
            deployProxy({
                name: "target2",
                impl: impl1,
                data: abi.encodeCall(Target.init, (address(delegateOwner)))
            })
        );

        TestMulticall3.Call3[] memory calls = new TestMulticall3.Call3[](3);
        calls[0].target = address(target1);
        calls[0].allowFailure = false;
        calls[0].callData = abi.encodeCall(EssentialContract.pause, ());

        calls[1].target = address(target2);
        calls[1].allowFailure = false;
        calls[1].callData = abi.encodeCall(UUPSUpgradeable.upgradeTo, (impl2));

        calls[2].target = address(delegateOwner);
        calls[2].allowFailure = false;
        calls[2].callData = abi.encodeCall(UUPSUpgradeable.upgradeTo, (delegateOwnerImpl2));

        bytes memory data = abi.encode(
            DelegateOwner.Call(
                uint64(0),
                address(multicall),
                true, // DELEGATECALL
                abi.encodeCall(TestMulticall3.aggregate3, (calls))
            )
        );

        vm.expectRevert(DelegateOwner.DO_DRYRUN_SUCCEEDED.selector);
        delegateOwner.dryrunMessageInvocation(data);

        IBridge.Message memory message;
        message.from = remoteOwner;
        message.destChainId = uint64(block.chainid);
        message.srcChainId = remoteChainId;
        message.destOwner = Bob;
        message.data = abi.encodeCall(DelegateOwner.onMessageInvocation, (data));
        message.to = address(delegateOwner);

        vm.prank(Bob);
        bridge.processMessage(message, "");

        bytes32 hash = bridge.hashMessage(message);
        assertTrue(bridge.messageStatus(hash) == IBridge.Status.DONE);

        assertEq(delegateOwner.nextTxId(), 1);
        assertTrue(target1.paused());
        assertEq(target2.impl(), impl2);
        assertEq(delegateOwner.impl(), delegateOwnerImpl2);
    }
}
