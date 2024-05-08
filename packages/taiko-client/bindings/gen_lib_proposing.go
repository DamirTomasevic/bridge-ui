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

// LibProposingMetaData contains all meta data concerning the LibProposing contract.
var LibProposingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"BlockProposed\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"assignedProver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"livenessBond\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"meta\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"components\":[{\"name\":\"l1Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"difficulty\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"coinbase\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l1Height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minTier\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"blobUsed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"parentMetaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"depositsProcessed\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit[]\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"L1_BLOB_NOT_AVAILABLE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOB_NOT_FOUND\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_HOOK\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_PROVER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_SIG\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_LIVENESS_BOND_NOT_RECEIVED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TOO_MANY_BLOCKS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNAUTHORIZED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNEXPECTED_PARENT\",\"inputs\":[]}]",
}

// LibProposingABI is the input ABI used to generate the binding from.
// Deprecated: Use LibProposingMetaData.ABI instead.
var LibProposingABI = LibProposingMetaData.ABI

// LibProposing is an auto generated Go binding around an Ethereum contract.
type LibProposing struct {
	LibProposingCaller     // Read-only binding to the contract
	LibProposingTransactor // Write-only binding to the contract
	LibProposingFilterer   // Log filterer for contract events
}

// LibProposingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibProposingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibProposingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibProposingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibProposingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibProposingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibProposingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibProposingSession struct {
	Contract     *LibProposing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibProposingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibProposingCallerSession struct {
	Contract *LibProposingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LibProposingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibProposingTransactorSession struct {
	Contract     *LibProposingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LibProposingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibProposingRaw struct {
	Contract *LibProposing // Generic contract binding to access the raw methods on
}

// LibProposingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibProposingCallerRaw struct {
	Contract *LibProposingCaller // Generic read-only contract binding to access the raw methods on
}

// LibProposingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibProposingTransactorRaw struct {
	Contract *LibProposingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibProposing creates a new instance of LibProposing, bound to a specific deployed contract.
func NewLibProposing(address common.Address, backend bind.ContractBackend) (*LibProposing, error) {
	contract, err := bindLibProposing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibProposing{LibProposingCaller: LibProposingCaller{contract: contract}, LibProposingTransactor: LibProposingTransactor{contract: contract}, LibProposingFilterer: LibProposingFilterer{contract: contract}}, nil
}

// NewLibProposingCaller creates a new read-only instance of LibProposing, bound to a specific deployed contract.
func NewLibProposingCaller(address common.Address, caller bind.ContractCaller) (*LibProposingCaller, error) {
	contract, err := bindLibProposing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibProposingCaller{contract: contract}, nil
}

// NewLibProposingTransactor creates a new write-only instance of LibProposing, bound to a specific deployed contract.
func NewLibProposingTransactor(address common.Address, transactor bind.ContractTransactor) (*LibProposingTransactor, error) {
	contract, err := bindLibProposing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibProposingTransactor{contract: contract}, nil
}

// NewLibProposingFilterer creates a new log filterer instance of LibProposing, bound to a specific deployed contract.
func NewLibProposingFilterer(address common.Address, filterer bind.ContractFilterer) (*LibProposingFilterer, error) {
	contract, err := bindLibProposing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibProposingFilterer{contract: contract}, nil
}

// bindLibProposing binds a generic wrapper to an already deployed contract.
func bindLibProposing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LibProposingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibProposing *LibProposingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibProposing.Contract.LibProposingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibProposing *LibProposingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibProposing.Contract.LibProposingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibProposing *LibProposingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibProposing.Contract.LibProposingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibProposing *LibProposingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibProposing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibProposing *LibProposingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibProposing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibProposing *LibProposingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibProposing.Contract.contract.Transact(opts, method, params...)
}

// LibProposingBlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the LibProposing contract.
type LibProposingBlockProposedIterator struct {
	Event *LibProposingBlockProposed // Event containing the contract specifics and raw log

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
func (it *LibProposingBlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibProposingBlockProposed)
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
		it.Event = new(LibProposingBlockProposed)
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
func (it *LibProposingBlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibProposingBlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibProposingBlockProposed represents a BlockProposed event raised by the LibProposing contract.
type LibProposingBlockProposed struct {
	BlockId           *big.Int
	AssignedProver    common.Address
	LivenessBond      *big.Int
	Meta              TaikoDataBlockMetadata
	DepositsProcessed []TaikoDataEthDeposit
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0xcda4e564245eb15494bc6da29f6a42e1941cf57f5314bf35bab8a1fca0a9c60a.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, (bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint16,bool,bytes32,address) meta, (address,uint96,uint64)[] depositsProcessed)
func (_LibProposing *LibProposingFilterer) FilterBlockProposed(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address) (*LibProposingBlockProposedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _LibProposing.contract.FilterLogs(opts, "BlockProposed", blockIdRule, assignedProverRule)
	if err != nil {
		return nil, err
	}
	return &LibProposingBlockProposedIterator{contract: _LibProposing.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0xcda4e564245eb15494bc6da29f6a42e1941cf57f5314bf35bab8a1fca0a9c60a.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, (bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint16,bool,bytes32,address) meta, (address,uint96,uint64)[] depositsProcessed)
func (_LibProposing *LibProposingFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *LibProposingBlockProposed, blockId []*big.Int, assignedProver []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _LibProposing.contract.WatchLogs(opts, "BlockProposed", blockIdRule, assignedProverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibProposingBlockProposed)
				if err := _LibProposing.contract.UnpackLog(event, "BlockProposed", log); err != nil {
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

// ParseBlockProposed is a log parse operation binding the contract event 0xcda4e564245eb15494bc6da29f6a42e1941cf57f5314bf35bab8a1fca0a9c60a.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, (bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint16,bool,bytes32,address) meta, (address,uint96,uint64)[] depositsProcessed)
func (_LibProposing *LibProposingFilterer) ParseBlockProposed(log types.Log) (*LibProposingBlockProposed, error) {
	event := new(LibProposingBlockProposed)
	if err := _LibProposing.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
