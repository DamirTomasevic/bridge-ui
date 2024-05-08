// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LibProvingMetaData contains all meta data concerning the LibProving contract.
var LibProvingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"ProvingPaused\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransitionContested\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tran\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.Transition\",\"components\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"graffiti\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"contester\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"contestBond\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransitionProved\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tran\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.Transition\",\"components\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"graffiti\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"validityBond\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"L1_ALREADY_CONTESTED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_ALREADY_PROVED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOCK_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_CANNOT_CONTEST\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_BLOCK_ID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_PAUSE_STATUS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_TIER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_TRANSITION\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_NOT_ASSIGNED_PROVER\",\"inputs\":[]}]",
}

// LibProvingABI is the input ABI used to generate the binding from.
// Deprecated: Use LibProvingMetaData.ABI instead.
var LibProvingABI = LibProvingMetaData.ABI

// LibProving is an auto generated Go binding around an Ethereum contract.
type LibProving struct {
	LibProvingCaller     // Read-only binding to the contract
	LibProvingTransactor // Write-only binding to the contract
	LibProvingFilterer   // Log filterer for contract events
}

// LibProvingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibProvingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibProvingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibProvingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibProvingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibProvingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibProvingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibProvingSession struct {
	Contract     *LibProving       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibProvingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibProvingCallerSession struct {
	Contract *LibProvingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LibProvingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibProvingTransactorSession struct {
	Contract     *LibProvingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LibProvingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibProvingRaw struct {
	Contract *LibProving // Generic contract binding to access the raw methods on
}

// LibProvingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibProvingCallerRaw struct {
	Contract *LibProvingCaller // Generic read-only contract binding to access the raw methods on
}

// LibProvingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibProvingTransactorRaw struct {
	Contract *LibProvingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibProving creates a new instance of LibProving, bound to a specific deployed contract.
func NewLibProving(address common.Address, backend bind.ContractBackend) (*LibProving, error) {
	contract, err := bindLibProving(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibProving{LibProvingCaller: LibProvingCaller{contract: contract}, LibProvingTransactor: LibProvingTransactor{contract: contract}, LibProvingFilterer: LibProvingFilterer{contract: contract}}, nil
}

// NewLibProvingCaller creates a new read-only instance of LibProving, bound to a specific deployed contract.
func NewLibProvingCaller(address common.Address, caller bind.ContractCaller) (*LibProvingCaller, error) {
	contract, err := bindLibProving(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibProvingCaller{contract: contract}, nil
}

// NewLibProvingTransactor creates a new write-only instance of LibProving, bound to a specific deployed contract.
func NewLibProvingTransactor(address common.Address, transactor bind.ContractTransactor) (*LibProvingTransactor, error) {
	contract, err := bindLibProving(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibProvingTransactor{contract: contract}, nil
}

// NewLibProvingFilterer creates a new log filterer instance of LibProving, bound to a specific deployed contract.
func NewLibProvingFilterer(address common.Address, filterer bind.ContractFilterer) (*LibProvingFilterer, error) {
	contract, err := bindLibProving(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibProvingFilterer{contract: contract}, nil
}

// bindLibProving binds a generic wrapper to an already deployed contract.
func bindLibProving(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LibProvingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibProving *LibProvingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibProving.Contract.LibProvingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibProving *LibProvingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibProving.Contract.LibProvingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibProving *LibProvingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibProving.Contract.LibProvingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibProving *LibProvingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibProving.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibProving *LibProvingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibProving.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibProving *LibProvingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibProving.Contract.contract.Transact(opts, method, params...)
}

// LibProvingProvingPausedIterator is returned from FilterProvingPaused and is used to iterate over the raw logs and unpacked data for ProvingPaused events raised by the LibProving contract.
type LibProvingProvingPausedIterator struct {
	Event *LibProvingProvingPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LibProvingProvingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibProvingProvingPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LibProvingProvingPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LibProvingProvingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibProvingProvingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibProvingProvingPaused represents a ProvingPaused event raised by the LibProving contract.
type LibProvingProvingPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProvingPaused is a free log retrieval operation binding the contract event 0xed64db85835d07c3c990b8ebdd55e32d64e5ed53143b6ef2179e7bfaf17ddc3b.
//
// Solidity: event ProvingPaused(bool paused)
func (_LibProving *LibProvingFilterer) FilterProvingPaused(opts *bind.FilterOpts) (*LibProvingProvingPausedIterator, error) {

	logs, sub, err := _LibProving.contract.FilterLogs(opts, "ProvingPaused")
	if err != nil {
		return nil, err
	}
	return &LibProvingProvingPausedIterator{contract: _LibProving.contract, event: "ProvingPaused", logs: logs, sub: sub}, nil
}

// WatchProvingPaused is a free log subscription operation binding the contract event 0xed64db85835d07c3c990b8ebdd55e32d64e5ed53143b6ef2179e7bfaf17ddc3b.
//
// Solidity: event ProvingPaused(bool paused)
func (_LibProving *LibProvingFilterer) WatchProvingPaused(opts *bind.WatchOpts, sink chan<- *LibProvingProvingPaused) (event.Subscription, error) {

	logs, sub, err := _LibProving.contract.WatchLogs(opts, "ProvingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibProvingProvingPaused)
				if err := _LibProving.contract.UnpackLog(event, "ProvingPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvingPaused is a log parse operation binding the contract event 0xed64db85835d07c3c990b8ebdd55e32d64e5ed53143b6ef2179e7bfaf17ddc3b.
//
// Solidity: event ProvingPaused(bool paused)
func (_LibProving *LibProvingFilterer) ParseProvingPaused(log types.Log) (*LibProvingProvingPaused, error) {
	event := new(LibProvingProvingPaused)
	if err := _LibProving.contract.UnpackLog(event, "ProvingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibProvingTransitionContestedIterator is returned from FilterTransitionContested and is used to iterate over the raw logs and unpacked data for TransitionContested events raised by the LibProving contract.
type LibProvingTransitionContestedIterator struct {
	Event *LibProvingTransitionContested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LibProvingTransitionContestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibProvingTransitionContested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LibProvingTransitionContested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LibProvingTransitionContestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibProvingTransitionContestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibProvingTransitionContested represents a TransitionContested event raised by the LibProving contract.
type LibProvingTransitionContested struct {
	BlockId     *big.Int
	Tran        TaikoDataTransition
	Contester   common.Address
	ContestBond *big.Int
	Tier        uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransitionContested is a free log retrieval operation binding the contract event 0xb4c0a86c1ff239277697775b1e91d3375fd3a5ef6b345aa4e2f6001c890558f6.
//
// Solidity: event TransitionContested(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address contester, uint96 contestBond, uint16 tier)
func (_LibProving *LibProvingFilterer) FilterTransitionContested(opts *bind.FilterOpts, blockId []*big.Int) (*LibProvingTransitionContestedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _LibProving.contract.FilterLogs(opts, "TransitionContested", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &LibProvingTransitionContestedIterator{contract: _LibProving.contract, event: "TransitionContested", logs: logs, sub: sub}, nil
}

// WatchTransitionContested is a free log subscription operation binding the contract event 0xb4c0a86c1ff239277697775b1e91d3375fd3a5ef6b345aa4e2f6001c890558f6.
//
// Solidity: event TransitionContested(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address contester, uint96 contestBond, uint16 tier)
func (_LibProving *LibProvingFilterer) WatchTransitionContested(opts *bind.WatchOpts, sink chan<- *LibProvingTransitionContested, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _LibProving.contract.WatchLogs(opts, "TransitionContested", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibProvingTransitionContested)
				if err := _LibProving.contract.UnpackLog(event, "TransitionContested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransitionContested is a log parse operation binding the contract event 0xb4c0a86c1ff239277697775b1e91d3375fd3a5ef6b345aa4e2f6001c890558f6.
//
// Solidity: event TransitionContested(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address contester, uint96 contestBond, uint16 tier)
func (_LibProving *LibProvingFilterer) ParseTransitionContested(log types.Log) (*LibProvingTransitionContested, error) {
	event := new(LibProvingTransitionContested)
	if err := _LibProving.contract.UnpackLog(event, "TransitionContested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibProvingTransitionProvedIterator is returned from FilterTransitionProved and is used to iterate over the raw logs and unpacked data for TransitionProved events raised by the LibProving contract.
type LibProvingTransitionProvedIterator struct {
	Event *LibProvingTransitionProved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LibProvingTransitionProvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibProvingTransitionProved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LibProvingTransitionProved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LibProvingTransitionProvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibProvingTransitionProvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibProvingTransitionProved represents a TransitionProved event raised by the LibProving contract.
type LibProvingTransitionProved struct {
	BlockId      *big.Int
	Tran         TaikoDataTransition
	Prover       common.Address
	ValidityBond *big.Int
	Tier         uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransitionProved is a free log retrieval operation binding the contract event 0xc195e4be3b936845492b8be4b1cf604db687a4d79ad84d979499c136f8e6701f.
//
// Solidity: event TransitionProved(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address prover, uint96 validityBond, uint16 tier)
func (_LibProving *LibProvingFilterer) FilterTransitionProved(opts *bind.FilterOpts, blockId []*big.Int) (*LibProvingTransitionProvedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _LibProving.contract.FilterLogs(opts, "TransitionProved", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &LibProvingTransitionProvedIterator{contract: _LibProving.contract, event: "TransitionProved", logs: logs, sub: sub}, nil
}

// WatchTransitionProved is a free log subscription operation binding the contract event 0xc195e4be3b936845492b8be4b1cf604db687a4d79ad84d979499c136f8e6701f.
//
// Solidity: event TransitionProved(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address prover, uint96 validityBond, uint16 tier)
func (_LibProving *LibProvingFilterer) WatchTransitionProved(opts *bind.WatchOpts, sink chan<- *LibProvingTransitionProved, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _LibProving.contract.WatchLogs(opts, "TransitionProved", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibProvingTransitionProved)
				if err := _LibProving.contract.UnpackLog(event, "TransitionProved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransitionProved is a log parse operation binding the contract event 0xc195e4be3b936845492b8be4b1cf604db687a4d79ad84d979499c136f8e6701f.
//
// Solidity: event TransitionProved(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address prover, uint96 validityBond, uint16 tier)
func (_LibProving *LibProvingFilterer) ParseTransitionProved(log types.Log) (*LibProvingTransitionProved, error) {
	event := new(LibProvingTransitionProved)
	if err := _LibProving.contract.UnpackLog(event, "TransitionProved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
