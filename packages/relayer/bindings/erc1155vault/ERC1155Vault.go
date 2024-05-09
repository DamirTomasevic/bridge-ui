// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1155vault

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

// BaseNFTVaultBridgeTransferOp is an auto generated low-level Go binding around an user-defined struct.
type BaseNFTVaultBridgeTransferOp struct {
	DestChainId uint64
	DestOwner   common.Address
	To          common.Address
	Fee         uint64
	Token       common.Address
	GasLimit    uint32
	TokenIds    []*big.Int
	Amounts     []*big.Int
}

// IBridgeMessage is an auto generated low-level Go binding around an user-defined struct.
type IBridgeMessage struct {
	Id          uint64
	Fee         uint64
	GasLimit    uint32
	From        common.Address
	SrcChainId  uint64
	SrcOwner    common.Address
	DestChainId uint64
	DestOwner   common.Address
	To          common.Address
	Value       *big.Int
	Data        []byte
}

// ERC1155VaultMetaData contains all meta data concerning the ERC1155Vault contract.
var ERC1155VaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_TOKEN_PER_TXN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addressManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgedToCanonical\",\"inputs\":[{\"name\":\"btoken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canonicalToBridged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ctoken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"btoken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inNonReentrant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addressManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastUnpausedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onMessageInvocation\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onMessageRecalled\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structIBridge.Message\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"srcOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendToken\",\"inputs\":[{\"name\":\"_op\",\"type\":\"tuple\",\"internalType\":\"structBaseNFTVault.BridgeTransferOp\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"message_\",\"type\":\"tuple\",\"internalType\":\"structIBridge.Message\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"srcOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgedTokenDeployed\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"ctoken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"btoken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ctokenSymbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ctokenName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenReceived\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"ctoken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenReleased\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ctoken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSent\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"ctoken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ETH_TRANSFER_FAILED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FUNC_NOT_IMPLEMENTED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_PAUSE_STATUS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"REENTRANT_CALL\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_DENIED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_INVALID_MANAGER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_UNEXPECTED_CHAINID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_ZERO_ADDR\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"VAULT_INTERFACE_NOT_SUPPORTED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VAULT_INVALID_AMOUNT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VAULT_INVALID_TOKEN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VAULT_INVALID_TO_ADDR\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VAULT_MAX_TOKEN_PER_TXN_EXCEEDED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VAULT_PERMISSION_DENIED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VAULT_TOKEN_ARRAY_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZERO_ADDR_MANAGER\",\"inputs\":[]}]",
}

// ERC1155VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155VaultMetaData.ABI instead.
var ERC1155VaultABI = ERC1155VaultMetaData.ABI

// ERC1155Vault is an auto generated Go binding around an Ethereum contract.
type ERC1155Vault struct {
	ERC1155VaultCaller     // Read-only binding to the contract
	ERC1155VaultTransactor // Write-only binding to the contract
	ERC1155VaultFilterer   // Log filterer for contract events
}

// ERC1155VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155VaultSession struct {
	Contract     *ERC1155Vault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155VaultCallerSession struct {
	Contract *ERC1155VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC1155VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155VaultTransactorSession struct {
	Contract     *ERC1155VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC1155VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155VaultRaw struct {
	Contract *ERC1155Vault // Generic contract binding to access the raw methods on
}

// ERC1155VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155VaultCallerRaw struct {
	Contract *ERC1155VaultCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155VaultTransactorRaw struct {
	Contract *ERC1155VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Vault creates a new instance of ERC1155Vault, bound to a specific deployed contract.
func NewERC1155Vault(address common.Address, backend bind.ContractBackend) (*ERC1155Vault, error) {
	contract, err := bindERC1155Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Vault{ERC1155VaultCaller: ERC1155VaultCaller{contract: contract}, ERC1155VaultTransactor: ERC1155VaultTransactor{contract: contract}, ERC1155VaultFilterer: ERC1155VaultFilterer{contract: contract}}, nil
}

// NewERC1155VaultCaller creates a new read-only instance of ERC1155Vault, bound to a specific deployed contract.
func NewERC1155VaultCaller(address common.Address, caller bind.ContractCaller) (*ERC1155VaultCaller, error) {
	contract, err := bindERC1155Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultCaller{contract: contract}, nil
}

// NewERC1155VaultTransactor creates a new write-only instance of ERC1155Vault, bound to a specific deployed contract.
func NewERC1155VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155VaultTransactor, error) {
	contract, err := bindERC1155Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultTransactor{contract: contract}, nil
}

// NewERC1155VaultFilterer creates a new log filterer instance of ERC1155Vault, bound to a specific deployed contract.
func NewERC1155VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155VaultFilterer, error) {
	contract, err := bindERC1155Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultFilterer{contract: contract}, nil
}

// bindERC1155Vault binds a generic wrapper to an already deployed contract.
func bindERC1155Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Vault *ERC1155VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Vault.Contract.ERC1155VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Vault *ERC1155VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.ERC1155VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Vault *ERC1155VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.ERC1155VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Vault *ERC1155VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Vault *ERC1155VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Vault *ERC1155VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.contract.Transact(opts, method, params...)
}

// MAXTOKENPERTXN is a free data retrieval call binding the contract method 0x634da63a.
//
// Solidity: function MAX_TOKEN_PER_TXN() view returns(uint256)
func (_ERC1155Vault *ERC1155VaultCaller) MAXTOKENPERTXN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "MAX_TOKEN_PER_TXN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKENPERTXN is a free data retrieval call binding the contract method 0x634da63a.
//
// Solidity: function MAX_TOKEN_PER_TXN() view returns(uint256)
func (_ERC1155Vault *ERC1155VaultSession) MAXTOKENPERTXN() (*big.Int, error) {
	return _ERC1155Vault.Contract.MAXTOKENPERTXN(&_ERC1155Vault.CallOpts)
}

// MAXTOKENPERTXN is a free data retrieval call binding the contract method 0x634da63a.
//
// Solidity: function MAX_TOKEN_PER_TXN() view returns(uint256)
func (_ERC1155Vault *ERC1155VaultCallerSession) MAXTOKENPERTXN() (*big.Int, error) {
	return _ERC1155Vault.Contract.MAXTOKENPERTXN(&_ERC1155Vault.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ERC1155Vault *ERC1155VaultCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ERC1155Vault *ERC1155VaultSession) AddressManager() (common.Address, error) {
	return _ERC1155Vault.Contract.AddressManager(&_ERC1155Vault.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ERC1155Vault *ERC1155VaultCallerSession) AddressManager() (common.Address, error) {
	return _ERC1155Vault.Contract.AddressManager(&_ERC1155Vault.CallOpts)
}

// BridgedToCanonical is a free data retrieval call binding the contract method 0x9aa8605c.
//
// Solidity: function bridgedToCanonical(address btoken) view returns(uint64 chainId, address addr, string symbol, string name)
func (_ERC1155Vault *ERC1155VaultCaller) BridgedToCanonical(opts *bind.CallOpts, btoken common.Address) (struct {
	ChainId uint64
	Addr    common.Address
	Symbol  string
	Name    string
}, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "bridgedToCanonical", btoken)

	outstruct := new(struct {
		ChainId uint64
		Addr    common.Address
		Symbol  string
		Name    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Symbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// BridgedToCanonical is a free data retrieval call binding the contract method 0x9aa8605c.
//
// Solidity: function bridgedToCanonical(address btoken) view returns(uint64 chainId, address addr, string symbol, string name)
func (_ERC1155Vault *ERC1155VaultSession) BridgedToCanonical(btoken common.Address) (struct {
	ChainId uint64
	Addr    common.Address
	Symbol  string
	Name    string
}, error) {
	return _ERC1155Vault.Contract.BridgedToCanonical(&_ERC1155Vault.CallOpts, btoken)
}

// BridgedToCanonical is a free data retrieval call binding the contract method 0x9aa8605c.
//
// Solidity: function bridgedToCanonical(address btoken) view returns(uint64 chainId, address addr, string symbol, string name)
func (_ERC1155Vault *ERC1155VaultCallerSession) BridgedToCanonical(btoken common.Address) (struct {
	ChainId uint64
	Addr    common.Address
	Symbol  string
	Name    string
}, error) {
	return _ERC1155Vault.Contract.BridgedToCanonical(&_ERC1155Vault.CallOpts, btoken)
}

// CanonicalToBridged is a free data retrieval call binding the contract method 0x67090ccf.
//
// Solidity: function canonicalToBridged(uint256 chainId, address ctoken) view returns(address btoken)
func (_ERC1155Vault *ERC1155VaultCaller) CanonicalToBridged(opts *bind.CallOpts, chainId *big.Int, ctoken common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "canonicalToBridged", chainId, ctoken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanonicalToBridged is a free data retrieval call binding the contract method 0x67090ccf.
//
// Solidity: function canonicalToBridged(uint256 chainId, address ctoken) view returns(address btoken)
func (_ERC1155Vault *ERC1155VaultSession) CanonicalToBridged(chainId *big.Int, ctoken common.Address) (common.Address, error) {
	return _ERC1155Vault.Contract.CanonicalToBridged(&_ERC1155Vault.CallOpts, chainId, ctoken)
}

// CanonicalToBridged is a free data retrieval call binding the contract method 0x67090ccf.
//
// Solidity: function canonicalToBridged(uint256 chainId, address ctoken) view returns(address btoken)
func (_ERC1155Vault *ERC1155VaultCallerSession) CanonicalToBridged(chainId *big.Int, ctoken common.Address) (common.Address, error) {
	return _ERC1155Vault.Contract.CanonicalToBridged(&_ERC1155Vault.CallOpts, chainId, ctoken)
}

// InNonReentrant is a free data retrieval call binding the contract method 0x3075db56.
//
// Solidity: function inNonReentrant() view returns(bool)
func (_ERC1155Vault *ERC1155VaultCaller) InNonReentrant(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "inNonReentrant")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InNonReentrant is a free data retrieval call binding the contract method 0x3075db56.
//
// Solidity: function inNonReentrant() view returns(bool)
func (_ERC1155Vault *ERC1155VaultSession) InNonReentrant() (bool, error) {
	return _ERC1155Vault.Contract.InNonReentrant(&_ERC1155Vault.CallOpts)
}

// InNonReentrant is a free data retrieval call binding the contract method 0x3075db56.
//
// Solidity: function inNonReentrant() view returns(bool)
func (_ERC1155Vault *ERC1155VaultCallerSession) InNonReentrant() (bool, error) {
	return _ERC1155Vault.Contract.InNonReentrant(&_ERC1155Vault.CallOpts)
}

// LastUnpausedAt is a free data retrieval call binding the contract method 0xe07baba6.
//
// Solidity: function lastUnpausedAt() view returns(uint64)
func (_ERC1155Vault *ERC1155VaultCaller) LastUnpausedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "lastUnpausedAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastUnpausedAt is a free data retrieval call binding the contract method 0xe07baba6.
//
// Solidity: function lastUnpausedAt() view returns(uint64)
func (_ERC1155Vault *ERC1155VaultSession) LastUnpausedAt() (uint64, error) {
	return _ERC1155Vault.Contract.LastUnpausedAt(&_ERC1155Vault.CallOpts)
}

// LastUnpausedAt is a free data retrieval call binding the contract method 0xe07baba6.
//
// Solidity: function lastUnpausedAt() view returns(uint64)
func (_ERC1155Vault *ERC1155VaultCallerSession) LastUnpausedAt() (uint64, error) {
	return _ERC1155Vault.Contract.LastUnpausedAt(&_ERC1155Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(bytes32)
func (_ERC1155Vault *ERC1155VaultCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(bytes32)
func (_ERC1155Vault *ERC1155VaultSession) Name() ([32]byte, error) {
	return _ERC1155Vault.Contract.Name(&_ERC1155Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(bytes32)
func (_ERC1155Vault *ERC1155VaultCallerSession) Name() ([32]byte, error) {
	return _ERC1155Vault.Contract.Name(&_ERC1155Vault.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_ERC1155Vault *ERC1155VaultCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_ERC1155Vault *ERC1155VaultSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _ERC1155Vault.Contract.OnERC1155BatchReceived(&_ERC1155Vault.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_ERC1155Vault *ERC1155VaultCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _ERC1155Vault.Contract.OnERC1155BatchReceived(&_ERC1155Vault.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_ERC1155Vault *ERC1155VaultCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_ERC1155Vault *ERC1155VaultSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _ERC1155Vault.Contract.OnERC1155Received(&_ERC1155Vault.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_ERC1155Vault *ERC1155VaultCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _ERC1155Vault.Contract.OnERC1155Received(&_ERC1155Vault.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC1155Vault *ERC1155VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC1155Vault *ERC1155VaultSession) Owner() (common.Address, error) {
	return _ERC1155Vault.Contract.Owner(&_ERC1155Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC1155Vault *ERC1155VaultCallerSession) Owner() (common.Address, error) {
	return _ERC1155Vault.Contract.Owner(&_ERC1155Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC1155Vault *ERC1155VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC1155Vault *ERC1155VaultSession) Paused() (bool, error) {
	return _ERC1155Vault.Contract.Paused(&_ERC1155Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC1155Vault *ERC1155VaultCallerSession) Paused() (bool, error) {
	return _ERC1155Vault.Contract.Paused(&_ERC1155Vault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC1155Vault *ERC1155VaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC1155Vault *ERC1155VaultSession) PendingOwner() (common.Address, error) {
	return _ERC1155Vault.Contract.PendingOwner(&_ERC1155Vault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC1155Vault *ERC1155VaultCallerSession) PendingOwner() (common.Address, error) {
	return _ERC1155Vault.Contract.PendingOwner(&_ERC1155Vault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC1155Vault *ERC1155VaultCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC1155Vault *ERC1155VaultSession) ProxiableUUID() ([32]byte, error) {
	return _ERC1155Vault.Contract.ProxiableUUID(&_ERC1155Vault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC1155Vault *ERC1155VaultCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ERC1155Vault.Contract.ProxiableUUID(&_ERC1155Vault.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 _chainId, bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_ERC1155Vault *ERC1155VaultCaller) Resolve(opts *bind.CallOpts, _chainId uint64, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "resolve", _chainId, _name, _allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 _chainId, bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_ERC1155Vault *ERC1155VaultSession) Resolve(_chainId uint64, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _ERC1155Vault.Contract.Resolve(&_ERC1155Vault.CallOpts, _chainId, _name, _allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 _chainId, bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_ERC1155Vault *ERC1155VaultCallerSession) Resolve(_chainId uint64, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _ERC1155Vault.Contract.Resolve(&_ERC1155Vault.CallOpts, _chainId, _name, _allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_ERC1155Vault *ERC1155VaultCaller) Resolve0(opts *bind.CallOpts, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "resolve0", _name, _allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_ERC1155Vault *ERC1155VaultSession) Resolve0(_name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _ERC1155Vault.Contract.Resolve0(&_ERC1155Vault.CallOpts, _name, _allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_ERC1155Vault *ERC1155VaultCallerSession) Resolve0(_name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _ERC1155Vault.Contract.Resolve0(&_ERC1155Vault.CallOpts, _name, _allowZeroAddress)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ERC1155Vault *ERC1155VaultCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155Vault.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ERC1155Vault *ERC1155VaultSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC1155Vault.Contract.SupportsInterface(&_ERC1155Vault.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ERC1155Vault *ERC1155VaultCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC1155Vault.Contract.SupportsInterface(&_ERC1155Vault.CallOpts, _interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC1155Vault *ERC1155VaultTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC1155Vault *ERC1155VaultSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.AcceptOwnership(&_ERC1155Vault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.AcceptOwnership(&_ERC1155Vault.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _owner, address _addressManager) returns()
func (_ERC1155Vault *ERC1155VaultTransactor) Init(opts *bind.TransactOpts, _owner common.Address, _addressManager common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "init", _owner, _addressManager)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _owner, address _addressManager) returns()
func (_ERC1155Vault *ERC1155VaultSession) Init(_owner common.Address, _addressManager common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.Init(&_ERC1155Vault.TransactOpts, _owner, _addressManager)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _owner, address _addressManager) returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) Init(_owner common.Address, _addressManager common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.Init(&_ERC1155Vault.TransactOpts, _owner, _addressManager)
}

// OnMessageInvocation is a paid mutator transaction binding the contract method 0x7f07c947.
//
// Solidity: function onMessageInvocation(bytes data) payable returns()
func (_ERC1155Vault *ERC1155VaultTransactor) OnMessageInvocation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "onMessageInvocation", data)
}

// OnMessageInvocation is a paid mutator transaction binding the contract method 0x7f07c947.
//
// Solidity: function onMessageInvocation(bytes data) payable returns()
func (_ERC1155Vault *ERC1155VaultSession) OnMessageInvocation(data []byte) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.OnMessageInvocation(&_ERC1155Vault.TransactOpts, data)
}

// OnMessageInvocation is a paid mutator transaction binding the contract method 0x7f07c947.
//
// Solidity: function onMessageInvocation(bytes data) payable returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) OnMessageInvocation(data []byte) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.OnMessageInvocation(&_ERC1155Vault.TransactOpts, data)
}

// OnMessageRecalled is a paid mutator transaction binding the contract method 0x0178733a.
//
// Solidity: function onMessageRecalled((uint64,uint64,uint32,address,uint64,address,uint64,address,address,uint256,bytes) message, bytes32 msgHash) payable returns()
func (_ERC1155Vault *ERC1155VaultTransactor) OnMessageRecalled(opts *bind.TransactOpts, message IBridgeMessage, msgHash [32]byte) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "onMessageRecalled", message, msgHash)
}

// OnMessageRecalled is a paid mutator transaction binding the contract method 0x0178733a.
//
// Solidity: function onMessageRecalled((uint64,uint64,uint32,address,uint64,address,uint64,address,address,uint256,bytes) message, bytes32 msgHash) payable returns()
func (_ERC1155Vault *ERC1155VaultSession) OnMessageRecalled(message IBridgeMessage, msgHash [32]byte) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.OnMessageRecalled(&_ERC1155Vault.TransactOpts, message, msgHash)
}

// OnMessageRecalled is a paid mutator transaction binding the contract method 0x0178733a.
//
// Solidity: function onMessageRecalled((uint64,uint64,uint32,address,uint64,address,uint64,address,address,uint256,bytes) message, bytes32 msgHash) payable returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) OnMessageRecalled(message IBridgeMessage, msgHash [32]byte) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.OnMessageRecalled(&_ERC1155Vault.TransactOpts, message, msgHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC1155Vault *ERC1155VaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC1155Vault *ERC1155VaultSession) Pause() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.Pause(&_ERC1155Vault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.Pause(&_ERC1155Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC1155Vault *ERC1155VaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC1155Vault *ERC1155VaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.RenounceOwnership(&_ERC1155Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.RenounceOwnership(&_ERC1155Vault.TransactOpts)
}

// SendToken is a paid mutator transaction binding the contract method 0x1f59a830.
//
// Solidity: function sendToken((uint64,address,address,uint64,address,uint32,uint256[],uint256[]) _op) payable returns((uint64,uint64,uint32,address,uint64,address,uint64,address,address,uint256,bytes) message_)
func (_ERC1155Vault *ERC1155VaultTransactor) SendToken(opts *bind.TransactOpts, _op BaseNFTVaultBridgeTransferOp) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "sendToken", _op)
}

// SendToken is a paid mutator transaction binding the contract method 0x1f59a830.
//
// Solidity: function sendToken((uint64,address,address,uint64,address,uint32,uint256[],uint256[]) _op) payable returns((uint64,uint64,uint32,address,uint64,address,uint64,address,address,uint256,bytes) message_)
func (_ERC1155Vault *ERC1155VaultSession) SendToken(_op BaseNFTVaultBridgeTransferOp) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.SendToken(&_ERC1155Vault.TransactOpts, _op)
}

// SendToken is a paid mutator transaction binding the contract method 0x1f59a830.
//
// Solidity: function sendToken((uint64,address,address,uint64,address,uint32,uint256[],uint256[]) _op) payable returns((uint64,uint64,uint32,address,uint64,address,uint64,address,address,uint256,bytes) message_)
func (_ERC1155Vault *ERC1155VaultTransactorSession) SendToken(_op BaseNFTVaultBridgeTransferOp) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.SendToken(&_ERC1155Vault.TransactOpts, _op)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC1155Vault *ERC1155VaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC1155Vault *ERC1155VaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.TransferOwnership(&_ERC1155Vault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.TransferOwnership(&_ERC1155Vault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC1155Vault *ERC1155VaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC1155Vault *ERC1155VaultSession) Unpause() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.Unpause(&_ERC1155Vault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC1155Vault.Contract.Unpause(&_ERC1155Vault.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ERC1155Vault *ERC1155VaultTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ERC1155Vault *ERC1155VaultSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.UpgradeTo(&_ERC1155Vault.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.UpgradeTo(&_ERC1155Vault.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC1155Vault *ERC1155VaultTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC1155Vault.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC1155Vault *ERC1155VaultSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.UpgradeToAndCall(&_ERC1155Vault.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC1155Vault *ERC1155VaultTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC1155Vault.Contract.UpgradeToAndCall(&_ERC1155Vault.TransactOpts, newImplementation, data)
}

// ERC1155VaultAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1155Vault contract.
type ERC1155VaultAdminChangedIterator struct {
	Event *ERC1155VaultAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultAdminChanged)
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
		it.Event = new(ERC1155VaultAdminChanged)
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
func (it *ERC1155VaultAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultAdminChanged represents a AdminChanged event raised by the ERC1155Vault contract.
type ERC1155VaultAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1155VaultAdminChangedIterator, error) {

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultAdminChangedIterator{contract: _ERC1155Vault.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1155VaultAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultAdminChanged)
				if err := _ERC1155Vault.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseAdminChanged(log types.Log) (*ERC1155VaultAdminChanged, error) {
	event := new(ERC1155VaultAdminChanged)
	if err := _ERC1155Vault.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1155Vault contract.
type ERC1155VaultBeaconUpgradedIterator struct {
	Event *ERC1155VaultBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultBeaconUpgraded)
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
		it.Event = new(ERC1155VaultBeaconUpgraded)
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
func (it *ERC1155VaultBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1155Vault contract.
type ERC1155VaultBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1155VaultBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultBeaconUpgradedIterator{contract: _ERC1155Vault.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1155VaultBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultBeaconUpgraded)
				if err := _ERC1155Vault.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1155VaultBeaconUpgraded, error) {
	event := new(ERC1155VaultBeaconUpgraded)
	if err := _ERC1155Vault.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultBridgedTokenDeployedIterator is returned from FilterBridgedTokenDeployed and is used to iterate over the raw logs and unpacked data for BridgedTokenDeployed events raised by the ERC1155Vault contract.
type ERC1155VaultBridgedTokenDeployedIterator struct {
	Event *ERC1155VaultBridgedTokenDeployed // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultBridgedTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultBridgedTokenDeployed)
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
		it.Event = new(ERC1155VaultBridgedTokenDeployed)
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
func (it *ERC1155VaultBridgedTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultBridgedTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultBridgedTokenDeployed represents a BridgedTokenDeployed event raised by the ERC1155Vault contract.
type ERC1155VaultBridgedTokenDeployed struct {
	ChainId      uint64
	Ctoken       common.Address
	Btoken       common.Address
	CtokenSymbol string
	CtokenName   string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgedTokenDeployed is a free log retrieval operation binding the contract event 0x44977f2d30fe1e3aee2c1476f2f95aaacaf34e44b9359c403da01fcc93fd751b.
//
// Solidity: event BridgedTokenDeployed(uint64 indexed chainId, address indexed ctoken, address indexed btoken, string ctokenSymbol, string ctokenName)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterBridgedTokenDeployed(opts *bind.FilterOpts, chainId []uint64, ctoken []common.Address, btoken []common.Address) (*ERC1155VaultBridgedTokenDeployedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var ctokenRule []interface{}
	for _, ctokenItem := range ctoken {
		ctokenRule = append(ctokenRule, ctokenItem)
	}
	var btokenRule []interface{}
	for _, btokenItem := range btoken {
		btokenRule = append(btokenRule, btokenItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "BridgedTokenDeployed", chainIdRule, ctokenRule, btokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultBridgedTokenDeployedIterator{contract: _ERC1155Vault.contract, event: "BridgedTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchBridgedTokenDeployed is a free log subscription operation binding the contract event 0x44977f2d30fe1e3aee2c1476f2f95aaacaf34e44b9359c403da01fcc93fd751b.
//
// Solidity: event BridgedTokenDeployed(uint64 indexed chainId, address indexed ctoken, address indexed btoken, string ctokenSymbol, string ctokenName)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchBridgedTokenDeployed(opts *bind.WatchOpts, sink chan<- *ERC1155VaultBridgedTokenDeployed, chainId []uint64, ctoken []common.Address, btoken []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var ctokenRule []interface{}
	for _, ctokenItem := range ctoken {
		ctokenRule = append(ctokenRule, ctokenItem)
	}
	var btokenRule []interface{}
	for _, btokenItem := range btoken {
		btokenRule = append(btokenRule, btokenItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "BridgedTokenDeployed", chainIdRule, ctokenRule, btokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultBridgedTokenDeployed)
				if err := _ERC1155Vault.contract.UnpackLog(event, "BridgedTokenDeployed", log); err != nil {
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

// ParseBridgedTokenDeployed is a log parse operation binding the contract event 0x44977f2d30fe1e3aee2c1476f2f95aaacaf34e44b9359c403da01fcc93fd751b.
//
// Solidity: event BridgedTokenDeployed(uint64 indexed chainId, address indexed ctoken, address indexed btoken, string ctokenSymbol, string ctokenName)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseBridgedTokenDeployed(log types.Log) (*ERC1155VaultBridgedTokenDeployed, error) {
	event := new(ERC1155VaultBridgedTokenDeployed)
	if err := _ERC1155Vault.contract.UnpackLog(event, "BridgedTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1155Vault contract.
type ERC1155VaultInitializedIterator struct {
	Event *ERC1155VaultInitialized // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultInitialized)
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
		it.Event = new(ERC1155VaultInitialized)
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
func (it *ERC1155VaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultInitialized represents a Initialized event raised by the ERC1155Vault contract.
type ERC1155VaultInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1155VaultInitializedIterator, error) {

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultInitializedIterator{contract: _ERC1155Vault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1155VaultInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultInitialized)
				if err := _ERC1155Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseInitialized(log types.Log) (*ERC1155VaultInitialized, error) {
	event := new(ERC1155VaultInitialized)
	if err := _ERC1155Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ERC1155Vault contract.
type ERC1155VaultOwnershipTransferStartedIterator struct {
	Event *ERC1155VaultOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultOwnershipTransferStarted)
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
		it.Event = new(ERC1155VaultOwnershipTransferStarted)
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
func (it *ERC1155VaultOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ERC1155Vault contract.
type ERC1155VaultOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC1155VaultOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultOwnershipTransferStartedIterator{contract: _ERC1155Vault.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ERC1155VaultOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultOwnershipTransferStarted)
				if err := _ERC1155Vault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseOwnershipTransferStarted(log types.Log) (*ERC1155VaultOwnershipTransferStarted, error) {
	event := new(ERC1155VaultOwnershipTransferStarted)
	if err := _ERC1155Vault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC1155Vault contract.
type ERC1155VaultOwnershipTransferredIterator struct {
	Event *ERC1155VaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultOwnershipTransferred)
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
		it.Event = new(ERC1155VaultOwnershipTransferred)
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
func (it *ERC1155VaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultOwnershipTransferred represents a OwnershipTransferred event raised by the ERC1155Vault contract.
type ERC1155VaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC1155VaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultOwnershipTransferredIterator{contract: _ERC1155Vault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC1155VaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultOwnershipTransferred)
				if err := _ERC1155Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseOwnershipTransferred(log types.Log) (*ERC1155VaultOwnershipTransferred, error) {
	event := new(ERC1155VaultOwnershipTransferred)
	if err := _ERC1155Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC1155Vault contract.
type ERC1155VaultPausedIterator struct {
	Event *ERC1155VaultPaused // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultPaused)
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
		it.Event = new(ERC1155VaultPaused)
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
func (it *ERC1155VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultPaused represents a Paused event raised by the ERC1155Vault contract.
type ERC1155VaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC1155VaultPausedIterator, error) {

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultPausedIterator{contract: _ERC1155Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC1155VaultPaused) (event.Subscription, error) {

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultPaused)
				if err := _ERC1155Vault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC1155Vault *ERC1155VaultFilterer) ParsePaused(log types.Log) (*ERC1155VaultPaused, error) {
	event := new(ERC1155VaultPaused)
	if err := _ERC1155Vault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultTokenReceivedIterator is returned from FilterTokenReceived and is used to iterate over the raw logs and unpacked data for TokenReceived events raised by the ERC1155Vault contract.
type ERC1155VaultTokenReceivedIterator struct {
	Event *ERC1155VaultTokenReceived // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultTokenReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultTokenReceived)
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
		it.Event = new(ERC1155VaultTokenReceived)
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
func (it *ERC1155VaultTokenReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultTokenReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultTokenReceived represents a TokenReceived event raised by the ERC1155Vault contract.
type ERC1155VaultTokenReceived struct {
	MsgHash    [32]byte
	From       common.Address
	To         common.Address
	SrcChainId uint64
	Ctoken     common.Address
	Token      common.Address
	TokenIds   []*big.Int
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenReceived is a free log retrieval operation binding the contract event 0x895f73e418d1bbbad2a311d085fad00e5d98a960e9f2afa4b942071d39bec43a.
//
// Solidity: event TokenReceived(bytes32 indexed msgHash, address indexed from, address indexed to, uint64 srcChainId, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterTokenReceived(opts *bind.FilterOpts, msgHash [][32]byte, from []common.Address, to []common.Address) (*ERC1155VaultTokenReceivedIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "TokenReceived", msgHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultTokenReceivedIterator{contract: _ERC1155Vault.contract, event: "TokenReceived", logs: logs, sub: sub}, nil
}

// WatchTokenReceived is a free log subscription operation binding the contract event 0x895f73e418d1bbbad2a311d085fad00e5d98a960e9f2afa4b942071d39bec43a.
//
// Solidity: event TokenReceived(bytes32 indexed msgHash, address indexed from, address indexed to, uint64 srcChainId, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchTokenReceived(opts *bind.WatchOpts, sink chan<- *ERC1155VaultTokenReceived, msgHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "TokenReceived", msgHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultTokenReceived)
				if err := _ERC1155Vault.contract.UnpackLog(event, "TokenReceived", log); err != nil {
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

// ParseTokenReceived is a log parse operation binding the contract event 0x895f73e418d1bbbad2a311d085fad00e5d98a960e9f2afa4b942071d39bec43a.
//
// Solidity: event TokenReceived(bytes32 indexed msgHash, address indexed from, address indexed to, uint64 srcChainId, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseTokenReceived(log types.Log) (*ERC1155VaultTokenReceived, error) {
	event := new(ERC1155VaultTokenReceived)
	if err := _ERC1155Vault.contract.UnpackLog(event, "TokenReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultTokenReleasedIterator is returned from FilterTokenReleased and is used to iterate over the raw logs and unpacked data for TokenReleased events raised by the ERC1155Vault contract.
type ERC1155VaultTokenReleasedIterator struct {
	Event *ERC1155VaultTokenReleased // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultTokenReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultTokenReleased)
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
		it.Event = new(ERC1155VaultTokenReleased)
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
func (it *ERC1155VaultTokenReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultTokenReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultTokenReleased represents a TokenReleased event raised by the ERC1155Vault contract.
type ERC1155VaultTokenReleased struct {
	MsgHash  [32]byte
	From     common.Address
	Ctoken   common.Address
	Token    common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenReleased is a free log retrieval operation binding the contract event 0xe48bef18455e47bca14864ab6e82dffa29df148b051c09de95aec44ecf13598c.
//
// Solidity: event TokenReleased(bytes32 indexed msgHash, address indexed from, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterTokenReleased(opts *bind.FilterOpts, msgHash [][32]byte, from []common.Address) (*ERC1155VaultTokenReleasedIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "TokenReleased", msgHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultTokenReleasedIterator{contract: _ERC1155Vault.contract, event: "TokenReleased", logs: logs, sub: sub}, nil
}

// WatchTokenReleased is a free log subscription operation binding the contract event 0xe48bef18455e47bca14864ab6e82dffa29df148b051c09de95aec44ecf13598c.
//
// Solidity: event TokenReleased(bytes32 indexed msgHash, address indexed from, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchTokenReleased(opts *bind.WatchOpts, sink chan<- *ERC1155VaultTokenReleased, msgHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "TokenReleased", msgHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultTokenReleased)
				if err := _ERC1155Vault.contract.UnpackLog(event, "TokenReleased", log); err != nil {
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

// ParseTokenReleased is a log parse operation binding the contract event 0xe48bef18455e47bca14864ab6e82dffa29df148b051c09de95aec44ecf13598c.
//
// Solidity: event TokenReleased(bytes32 indexed msgHash, address indexed from, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseTokenReleased(log types.Log) (*ERC1155VaultTokenReleased, error) {
	event := new(ERC1155VaultTokenReleased)
	if err := _ERC1155Vault.contract.UnpackLog(event, "TokenReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the ERC1155Vault contract.
type ERC1155VaultTokenSentIterator struct {
	Event *ERC1155VaultTokenSent // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultTokenSent)
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
		it.Event = new(ERC1155VaultTokenSent)
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
func (it *ERC1155VaultTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultTokenSent represents a TokenSent event raised by the ERC1155Vault contract.
type ERC1155VaultTokenSent struct {
	MsgHash     [32]byte
	From        common.Address
	To          common.Address
	DestChainId uint64
	Ctoken      common.Address
	Token       common.Address
	TokenIds    []*big.Int
	Amounts     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0xabbf62a1459339f9ac59136d313a5ccd83d2706cc6d4c04d90642520169144dc.
//
// Solidity: event TokenSent(bytes32 indexed msgHash, address indexed from, address indexed to, uint64 destChainId, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterTokenSent(opts *bind.FilterOpts, msgHash [][32]byte, from []common.Address, to []common.Address) (*ERC1155VaultTokenSentIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "TokenSent", msgHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultTokenSentIterator{contract: _ERC1155Vault.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0xabbf62a1459339f9ac59136d313a5ccd83d2706cc6d4c04d90642520169144dc.
//
// Solidity: event TokenSent(bytes32 indexed msgHash, address indexed from, address indexed to, uint64 destChainId, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *ERC1155VaultTokenSent, msgHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "TokenSent", msgHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultTokenSent)
				if err := _ERC1155Vault.contract.UnpackLog(event, "TokenSent", log); err != nil {
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

// ParseTokenSent is a log parse operation binding the contract event 0xabbf62a1459339f9ac59136d313a5ccd83d2706cc6d4c04d90642520169144dc.
//
// Solidity: event TokenSent(bytes32 indexed msgHash, address indexed from, address indexed to, uint64 destChainId, address ctoken, address token, uint256[] tokenIds, uint256[] amounts)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseTokenSent(log types.Log) (*ERC1155VaultTokenSent, error) {
	event := new(ERC1155VaultTokenSent)
	if err := _ERC1155Vault.contract.UnpackLog(event, "TokenSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC1155Vault contract.
type ERC1155VaultUnpausedIterator struct {
	Event *ERC1155VaultUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultUnpaused)
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
		it.Event = new(ERC1155VaultUnpaused)
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
func (it *ERC1155VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultUnpaused represents a Unpaused event raised by the ERC1155Vault contract.
type ERC1155VaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC1155VaultUnpausedIterator, error) {

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultUnpausedIterator{contract: _ERC1155Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC1155VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultUnpaused)
				if err := _ERC1155Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseUnpaused(log types.Log) (*ERC1155VaultUnpaused, error) {
	event := new(ERC1155VaultUnpaused)
	if err := _ERC1155Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155VaultUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1155Vault contract.
type ERC1155VaultUpgradedIterator struct {
	Event *ERC1155VaultUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1155VaultUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155VaultUpgraded)
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
		it.Event = new(ERC1155VaultUpgraded)
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
func (it *ERC1155VaultUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155VaultUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155VaultUpgraded represents a Upgraded event raised by the ERC1155Vault contract.
type ERC1155VaultUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1155Vault *ERC1155VaultFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1155VaultUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1155Vault.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155VaultUpgradedIterator{contract: _ERC1155Vault.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1155Vault *ERC1155VaultFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1155VaultUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1155Vault.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155VaultUpgraded)
				if err := _ERC1155Vault.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1155Vault *ERC1155VaultFilterer) ParseUpgraded(log types.Log) (*ERC1155VaultUpgraded, error) {
	event := new(ERC1155VaultUpgraded)
	if err := _ERC1155Vault.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
