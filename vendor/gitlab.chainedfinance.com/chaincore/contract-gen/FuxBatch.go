// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_gen

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// FuxBatchABI is the input ABI used to generate the binding from.
const FuxBatchABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cursorGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isFuxHub\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"runMaxCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInFuxGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_router\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_len\",\"type\":\"uint256\"}],\"name\":\"setMaxLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cnt\",\"type\":\"uint256\"}],\"name\":\"setRunMaxLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gas\",\"type\":\"uint256\"}],\"name\":\"setTransferGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokens\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokens\",\"type\":\"uint256[]\"}],\"name\":\"addJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"runJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isCompile\",\"outputs\":[{\"name\":\"compile\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxBatch is an auto generated Go binding around an Ethereum contract.
type FuxBatch struct {
	FuxBatchCaller     // Read-only binding to the contract
	FuxBatchTransactor // Write-only binding to the contract
	FuxBatchFilterer   // Log filterer for contract events
}

// FuxBatchCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxBatchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxBatchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxBatchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxBatchSession struct {
	Contract     *FuxBatch         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxBatchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxBatchCallerSession struct {
	Contract *FuxBatchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxBatchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxBatchTransactorSession struct {
	Contract     *FuxBatchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxBatchRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxBatchRaw struct {
	Contract *FuxBatch // Generic contract binding to access the raw methods on
}

// FuxBatchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxBatchCallerRaw struct {
	Contract *FuxBatchCaller // Generic read-only contract binding to access the raw methods on
}

// FuxBatchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxBatchTransactorRaw struct {
	Contract *FuxBatchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxBatch creates a new instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatch(address common.Address, backend bind.ContractBackend) (*FuxBatch, error) {
	contract, err := bindFuxBatch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}, FuxBatchFilterer: FuxBatchFilterer{contract: contract}}, nil
}

// NewFuxBatchCaller creates a new read-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchCaller(address common.Address, caller bind.ContractCaller) (*FuxBatchCaller, error) {
	contract, err := bindFuxBatch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchCaller{contract: contract}, nil
}

// NewFuxBatchTransactor creates a new write-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxBatchTransactor, error) {
	contract, err := bindFuxBatch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchTransactor{contract: contract}, nil
}

// NewFuxBatchFilterer creates a new log filterer instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxBatchFilterer, error) {
	contract, err := bindFuxBatch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxBatchFilterer{contract: contract}, nil
}

// bindFuxBatch binds a generic wrapper to an already deployed contract.
func bindFuxBatch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxBatchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxBatch *FuxBatchRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxBatch.Contract.FuxBatchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxBatch *FuxBatchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxBatch.Contract.FuxBatchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxBatch *FuxBatchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxBatch.Contract.FuxBatchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxBatch *FuxBatchCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxBatch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxBatch *FuxBatchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxBatch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxBatch *FuxBatchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxBatch.Contract.contract.Transact(opts, method, params...)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLECLUSTER() (string, error) {
	return _FuxBatch.Contract.ROLECLUSTER(&_FuxBatch.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLECLUSTER() (string, error) {
	return _FuxBatch.Contract.ROLECLUSTER(&_FuxBatch.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXGROUP() (string, error) {
	return _FuxBatch.Contract.ROLEFUXGROUP(&_FuxBatch.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxBatch.Contract.ROLEFUXGROUP(&_FuxBatch.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXHUB() (string, error) {
	return _FuxBatch.Contract.ROLEFUXHUB(&_FuxBatch.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxBatch.Contract.ROLEFUXHUB(&_FuxBatch.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTECFRBAC() (string, error) {
	return _FuxBatch.Contract.ROUTECFRBAC(&_FuxBatch.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxBatch.Contract.ROUTECFRBAC(&_FuxBatch.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTECONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTECONFIG(&_FuxBatch.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTECONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTECONFIG(&_FuxBatch.CallOpts)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsFuxHub(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "_isFuxHub", _addr)
	return *ret0, err
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsFuxHub(&_FuxBatch.CallOpts, _addr)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsFuxHub(&_FuxBatch.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsInFuxGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "_isInFuxGroup", _addr)
	return *ret0, err
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsInFuxGroup(&_FuxBatch.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsInFuxGroup(&_FuxBatch.CallOpts, _addr)
}

// CursorGas is a free data retrieval call binding the contract method 0x2133bfc6.
//
// Solidity: function cursorGas() constant returns(uint256)
func (_FuxBatch *FuxBatchCaller) CursorGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "cursorGas")
	return *ret0, err
}

// CursorGas is a free data retrieval call binding the contract method 0x2133bfc6.
//
// Solidity: function cursorGas() constant returns(uint256)
func (_FuxBatch *FuxBatchSession) CursorGas() (*big.Int, error) {
	return _FuxBatch.Contract.CursorGas(&_FuxBatch.CallOpts)
}

// CursorGas is a free data retrieval call binding the contract method 0x2133bfc6.
//
// Solidity: function cursorGas() constant returns(uint256)
func (_FuxBatch *FuxBatchCallerSession) CursorGas() (*big.Int, error) {
	return _FuxBatch.Contract.CursorGas(&_FuxBatch.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxBatch *FuxBatchCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxBatch *FuxBatchSession) GetRouterAddress() (common.Address, error) {
	return _FuxBatch.Contract.GetRouterAddress(&_FuxBatch.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxBatch.Contract.GetRouterAddress(&_FuxBatch.CallOpts)
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchCaller) IsCompile(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isCompile", _id)
	return *ret0, err
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchSession) IsCompile(_id *big.Int) (bool, error) {
	return _FuxBatch.Contract.IsCompile(&_FuxBatch.CallOpts, _id)
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchCallerSession) IsCompile(_id *big.Int) (bool, error) {
	return _FuxBatch.Contract.IsCompile(&_FuxBatch.CallOpts, _id)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsUsingRouter() (bool, error) {
	return _FuxBatch.Contract.IsUsingRouter(&_FuxBatch.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsUsingRouter() (bool, error) {
	return _FuxBatch.Contract.IsUsingRouter(&_FuxBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchCaller) MaxLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "maxLength")
	return *ret0, err
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchSession) MaxLength() (*big.Int, error) {
	return _FuxBatch.Contract.MaxLength(&_FuxBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchCallerSession) MaxLength() (*big.Int, error) {
	return _FuxBatch.Contract.MaxLength(&_FuxBatch.CallOpts)
}

// RunMaxCount is a free data retrieval call binding the contract method 0xe2bf84f8.
//
// Solidity: function runMaxCount() constant returns(uint256)
func (_FuxBatch *FuxBatchCaller) RunMaxCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "runMaxCount")
	return *ret0, err
}

// RunMaxCount is a free data retrieval call binding the contract method 0xe2bf84f8.
//
// Solidity: function runMaxCount() constant returns(uint256)
func (_FuxBatch *FuxBatchSession) RunMaxCount() (*big.Int, error) {
	return _FuxBatch.Contract.RunMaxCount(&_FuxBatch.CallOpts)
}

// RunMaxCount is a free data retrieval call binding the contract method 0xe2bf84f8.
//
// Solidity: function runMaxCount() constant returns(uint256)
func (_FuxBatch *FuxBatchCallerSession) RunMaxCount() (*big.Int, error) {
	return _FuxBatch.Contract.RunMaxCount(&_FuxBatch.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxBatch *FuxBatchCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxBatch *FuxBatchSession) SysRouter() (common.Address, error) {
	return _FuxBatch.Contract.SysRouter(&_FuxBatch.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) SysRouter() (common.Address, error) {
	return _FuxBatch.Contract.SysRouter(&_FuxBatch.CallOpts)
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() constant returns(uint256)
func (_FuxBatch *FuxBatchCaller) TransferGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "transferGas")
	return *ret0, err
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() constant returns(uint256)
func (_FuxBatch *FuxBatchSession) TransferGas() (*big.Int, error) {
	return _FuxBatch.Contract.TransferGas(&_FuxBatch.CallOpts)
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() constant returns(uint256)
func (_FuxBatch *FuxBatchCallerSession) TransferGas() (*big.Int, error) {
	return _FuxBatch.Contract.TransferGas(&_FuxBatch.CallOpts)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactor) AddJob(opts *bind.TransactOpts, _id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "addJob", _id, _from, _to, _tokens)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchSession) AddJob(_id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.AddJob(&_FuxBatch.TransactOpts, _id, _from, _to, _tokens)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactorSession) AddJob(_id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.AddJob(&_FuxBatch.TransactOpts, _id, _from, _to, _tokens)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchTransactor) RunJob(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "runJob", _id)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchSession) RunJob(_id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.RunJob(&_FuxBatch.TransactOpts, _id)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) RunJob(_id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.RunJob(&_FuxBatch.TransactOpts, _id)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchTransactor) SetMaxLength(opts *bind.TransactOpts, _len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "setMaxLength", _len)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchSession) SetMaxLength(_len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetMaxLength(&_FuxBatch.TransactOpts, _len)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) SetMaxLength(_len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetMaxLength(&_FuxBatch.TransactOpts, _len)
}

// SetRunMaxLength is a paid mutator transaction binding the contract method 0x5d2ddd15.
//
// Solidity: function setRunMaxLength(_cnt uint256) returns()
func (_FuxBatch *FuxBatchTransactor) SetRunMaxLength(opts *bind.TransactOpts, _cnt *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "setRunMaxLength", _cnt)
}

// SetRunMaxLength is a paid mutator transaction binding the contract method 0x5d2ddd15.
//
// Solidity: function setRunMaxLength(_cnt uint256) returns()
func (_FuxBatch *FuxBatchSession) SetRunMaxLength(_cnt *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetRunMaxLength(&_FuxBatch.TransactOpts, _cnt)
}

// SetRunMaxLength is a paid mutator transaction binding the contract method 0x5d2ddd15.
//
// Solidity: function setRunMaxLength(_cnt uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) SetRunMaxLength(_cnt *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetRunMaxLength(&_FuxBatch.TransactOpts, _cnt)
}

// SetTransferGas is a paid mutator transaction binding the contract method 0xc7e066ba.
//
// Solidity: function setTransferGas(_gas uint256) returns()
func (_FuxBatch *FuxBatchTransactor) SetTransferGas(opts *bind.TransactOpts, _gas *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "setTransferGas", _gas)
}

// SetTransferGas is a paid mutator transaction binding the contract method 0xc7e066ba.
//
// Solidity: function setTransferGas(_gas uint256) returns()
func (_FuxBatch *FuxBatchSession) SetTransferGas(_gas *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetTransferGas(&_FuxBatch.TransactOpts, _gas)
}

// SetTransferGas is a paid mutator transaction binding the contract method 0xc7e066ba.
//
// Solidity: function setTransferGas(_gas uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) SetTransferGas(_gas *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetTransferGas(&_FuxBatch.TransactOpts, _gas)
}

// Transfer is a paid mutator transaction binding the contract method 0x2b4e4e96.
//
// Solidity: function transfer(_to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "transfer", _to, _tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0x2b4e4e96.
//
// Solidity: function transfer(_to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchSession) Transfer(_to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.Transfer(&_FuxBatch.TransactOpts, _to, _tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0x2b4e4e96.
//
// Solidity: function transfer(_to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactorSession) Transfer(_to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.Transfer(&_FuxBatch.TransactOpts, _to, _tokens)
}
