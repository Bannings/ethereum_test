// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_gen

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// FuxLockerABI is the input ABI used to generate the binding from.
const FuxLockerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"fuxStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_EXPIRE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_STATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_EXISTED\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deltaT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_deltaT\",\"type\":\"uint256\"}],\"name\":\"DeltaTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_lock\",\"type\":\"bool\"}],\"name\":\"LockStateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"StateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"ExpireSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_fuxStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_lock\",\"type\":\"bool\"}],\"name\":\"setLockState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"setState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"setExpire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getExpire\",\"outputs\":[{\"name\":\"expire\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isLock\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_deltaT\",\"type\":\"uint256\"}],\"name\":\"setDeltaT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxLocker is an auto generated Go binding around an Ethereum contract.
type FuxLocker struct {
	FuxLockerCaller     // Read-only binding to the contract
	FuxLockerTransactor // Write-only binding to the contract
	FuxLockerFilterer   // Log filterer for contract events
}

// FuxLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxLockerSession struct {
	Contract     *FuxLocker        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxLockerCallerSession struct {
	Contract *FuxLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FuxLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxLockerTransactorSession struct {
	Contract     *FuxLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FuxLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxLockerRaw struct {
	Contract *FuxLocker // Generic contract binding to access the raw methods on
}

// FuxLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxLockerCallerRaw struct {
	Contract *FuxLockerCaller // Generic read-only contract binding to access the raw methods on
}

// FuxLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxLockerTransactorRaw struct {
	Contract *FuxLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxLocker creates a new instance of FuxLocker, bound to a specific deployed contract.
func NewFuxLocker(address common.Address, backend bind.ContractBackend) (*FuxLocker, error) {
	contract, err := bindFuxLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxLocker{FuxLockerCaller: FuxLockerCaller{contract: contract}, FuxLockerTransactor: FuxLockerTransactor{contract: contract}, FuxLockerFilterer: FuxLockerFilterer{contract: contract}}, nil
}

// NewFuxLockerCaller creates a new read-only instance of FuxLocker, bound to a specific deployed contract.
func NewFuxLockerCaller(address common.Address, caller bind.ContractCaller) (*FuxLockerCaller, error) {
	contract, err := bindFuxLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxLockerCaller{contract: contract}, nil
}

// NewFuxLockerTransactor creates a new write-only instance of FuxLocker, bound to a specific deployed contract.
func NewFuxLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxLockerTransactor, error) {
	contract, err := bindFuxLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxLockerTransactor{contract: contract}, nil
}

// NewFuxLockerFilterer creates a new log filterer instance of FuxLocker, bound to a specific deployed contract.
func NewFuxLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxLockerFilterer, error) {
	contract, err := bindFuxLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxLockerFilterer{contract: contract}, nil
}

// bindFuxLocker binds a generic wrapper to an already deployed contract.
func bindFuxLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxLockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxLocker *FuxLockerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxLocker.Contract.FuxLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxLocker *FuxLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxLocker.Contract.FuxLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxLocker *FuxLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxLocker.Contract.FuxLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxLocker *FuxLockerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxLocker *FuxLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxLocker *FuxLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxLocker.Contract.contract.Transact(opts, method, params...)
}

// IDXEXISTED is a free data retrieval call binding the contract method 0x82185f8c.
//
// Solidity: function IDX_EXISTED() constant returns(uint256)
func (_FuxLocker *FuxLockerCaller) IDXEXISTED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "IDX_EXISTED")
	return *ret0, err
}

// IDXEXISTED is a free data retrieval call binding the contract method 0x82185f8c.
//
// Solidity: function IDX_EXISTED() constant returns(uint256)
func (_FuxLocker *FuxLockerSession) IDXEXISTED() (*big.Int, error) {
	return _FuxLocker.Contract.IDXEXISTED(&_FuxLocker.CallOpts)
}

// IDXEXISTED is a free data retrieval call binding the contract method 0x82185f8c.
//
// Solidity: function IDX_EXISTED() constant returns(uint256)
func (_FuxLocker *FuxLockerCallerSession) IDXEXISTED() (*big.Int, error) {
	return _FuxLocker.Contract.IDXEXISTED(&_FuxLocker.CallOpts)
}

// IDXEXPIRE is a free data retrieval call binding the contract method 0x66a41a70.
//
// Solidity: function IDX_EXPIRE() constant returns(uint256)
func (_FuxLocker *FuxLockerCaller) IDXEXPIRE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "IDX_EXPIRE")
	return *ret0, err
}

// IDXEXPIRE is a free data retrieval call binding the contract method 0x66a41a70.
//
// Solidity: function IDX_EXPIRE() constant returns(uint256)
func (_FuxLocker *FuxLockerSession) IDXEXPIRE() (*big.Int, error) {
	return _FuxLocker.Contract.IDXEXPIRE(&_FuxLocker.CallOpts)
}

// IDXEXPIRE is a free data retrieval call binding the contract method 0x66a41a70.
//
// Solidity: function IDX_EXPIRE() constant returns(uint256)
func (_FuxLocker *FuxLockerCallerSession) IDXEXPIRE() (*big.Int, error) {
	return _FuxLocker.Contract.IDXEXPIRE(&_FuxLocker.CallOpts)
}

// IDXSTATE is a free data retrieval call binding the contract method 0x79aa0b41.
//
// Solidity: function IDX_STATE() constant returns(uint256)
func (_FuxLocker *FuxLockerCaller) IDXSTATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "IDX_STATE")
	return *ret0, err
}

// IDXSTATE is a free data retrieval call binding the contract method 0x79aa0b41.
//
// Solidity: function IDX_STATE() constant returns(uint256)
func (_FuxLocker *FuxLockerSession) IDXSTATE() (*big.Int, error) {
	return _FuxLocker.Contract.IDXSTATE(&_FuxLocker.CallOpts)
}

// IDXSTATE is a free data retrieval call binding the contract method 0x79aa0b41.
//
// Solidity: function IDX_STATE() constant returns(uint256)
func (_FuxLocker *FuxLockerCallerSession) IDXSTATE() (*big.Int, error) {
	return _FuxLocker.Contract.IDXSTATE(&_FuxLocker.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxLocker *FuxLockerCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxLocker *FuxLockerSession) ROLEADMIN() (string, error) {
	return _FuxLocker.Contract.ROLEADMIN(&_FuxLocker.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxLocker *FuxLockerCallerSession) ROLEADMIN() (string, error) {
	return _FuxLocker.Contract.ROLEADMIN(&_FuxLocker.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxLocker *FuxLockerCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxLocker *FuxLockerSession) ROLECLUSTER() (string, error) {
	return _FuxLocker.Contract.ROLECLUSTER(&_FuxLocker.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxLocker *FuxLockerCallerSession) ROLECLUSTER() (string, error) {
	return _FuxLocker.Contract.ROLECLUSTER(&_FuxLocker.CallOpts)
}

// DeltaT is a free data retrieval call binding the contract method 0xc7a0f8f7.
//
// Solidity: function deltaT() constant returns(uint256)
func (_FuxLocker *FuxLockerCaller) DeltaT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "deltaT")
	return *ret0, err
}

// DeltaT is a free data retrieval call binding the contract method 0xc7a0f8f7.
//
// Solidity: function deltaT() constant returns(uint256)
func (_FuxLocker *FuxLockerSession) DeltaT() (*big.Int, error) {
	return _FuxLocker.Contract.DeltaT(&_FuxLocker.CallOpts)
}

// DeltaT is a free data retrieval call binding the contract method 0xc7a0f8f7.
//
// Solidity: function deltaT() constant returns(uint256)
func (_FuxLocker *FuxLockerCallerSession) DeltaT() (*big.Int, error) {
	return _FuxLocker.Contract.DeltaT(&_FuxLocker.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxLocker *FuxLockerCaller) FuxStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "fuxStorage")
	return *ret0, err
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxLocker *FuxLockerSession) FuxStorage() (common.Address, error) {
	return _FuxLocker.Contract.FuxStorage(&_FuxLocker.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxLocker *FuxLockerCallerSession) FuxStorage() (common.Address, error) {
	return _FuxLocker.Contract.FuxStorage(&_FuxLocker.CallOpts)
}

// GetExpire is a free data retrieval call binding the contract method 0x95c78f82.
//
// Solidity: function getExpire(_tokenId uint256) constant returns(expire uint256)
func (_FuxLocker *FuxLockerCaller) GetExpire(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "getExpire", _tokenId)
	return *ret0, err
}

// GetExpire is a free data retrieval call binding the contract method 0x95c78f82.
//
// Solidity: function getExpire(_tokenId uint256) constant returns(expire uint256)
func (_FuxLocker *FuxLockerSession) GetExpire(_tokenId *big.Int) (*big.Int, error) {
	return _FuxLocker.Contract.GetExpire(&_FuxLocker.CallOpts, _tokenId)
}

// GetExpire is a free data retrieval call binding the contract method 0x95c78f82.
//
// Solidity: function getExpire(_tokenId uint256) constant returns(expire uint256)
func (_FuxLocker *FuxLockerCallerSession) GetExpire(_tokenId *big.Int) (*big.Int, error) {
	return _FuxLocker.Contract.GetExpire(&_FuxLocker.CallOpts, _tokenId)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(_tokenId uint256) constant returns(state uint256)
func (_FuxLocker *FuxLockerCaller) GetState(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "getState", _tokenId)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(_tokenId uint256) constant returns(state uint256)
func (_FuxLocker *FuxLockerSession) GetState(_tokenId *big.Int) (*big.Int, error) {
	return _FuxLocker.Contract.GetState(&_FuxLocker.CallOpts, _tokenId)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(_tokenId uint256) constant returns(state uint256)
func (_FuxLocker *FuxLockerCallerSession) GetState(_tokenId *big.Int) (*big.Int, error) {
	return _FuxLocker.Contract.GetState(&_FuxLocker.CallOpts, _tokenId)
}

// IsLock is a free data retrieval call binding the contract method 0x603cc0b8.
//
// Solidity: function isLock(_tokenId uint256) constant returns(ret bool)
func (_FuxLocker *FuxLockerCaller) IsLock(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "isLock", _tokenId)
	return *ret0, err
}

// IsLock is a free data retrieval call binding the contract method 0x603cc0b8.
//
// Solidity: function isLock(_tokenId uint256) constant returns(ret bool)
func (_FuxLocker *FuxLockerSession) IsLock(_tokenId *big.Int) (bool, error) {
	return _FuxLocker.Contract.IsLock(&_FuxLocker.CallOpts, _tokenId)
}

// IsLock is a free data retrieval call binding the contract method 0x603cc0b8.
//
// Solidity: function isLock(_tokenId uint256) constant returns(ret bool)
func (_FuxLocker *FuxLockerCallerSession) IsLock(_tokenId *big.Int) (bool, error) {
	return _FuxLocker.Contract.IsLock(&_FuxLocker.CallOpts, _tokenId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxLocker *FuxLockerCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxLocker *FuxLockerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxLocker.Contract.IsMigrated(&_FuxLocker.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxLocker *FuxLockerCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxLocker.Contract.IsMigrated(&_FuxLocker.CallOpts, contractName, migrationId)
}

// LockList is a free data retrieval call binding the contract method 0xa41d222b.
//
// Solidity: function lockList( uint256) constant returns(bool)
func (_FuxLocker *FuxLockerCaller) LockList(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "lockList", arg0)
	return *ret0, err
}

// LockList is a free data retrieval call binding the contract method 0xa41d222b.
//
// Solidity: function lockList( uint256) constant returns(bool)
func (_FuxLocker *FuxLockerSession) LockList(arg0 *big.Int) (bool, error) {
	return _FuxLocker.Contract.LockList(&_FuxLocker.CallOpts, arg0)
}

// LockList is a free data retrieval call binding the contract method 0xa41d222b.
//
// Solidity: function lockList( uint256) constant returns(bool)
func (_FuxLocker *FuxLockerCallerSession) LockList(arg0 *big.Int) (bool, error) {
	return _FuxLocker.Contract.LockList(&_FuxLocker.CallOpts, arg0)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxLocker *FuxLockerCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxLocker.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxLocker *FuxLockerSession) Rbac() (common.Address, error) {
	return _FuxLocker.Contract.Rbac(&_FuxLocker.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxLocker *FuxLockerCallerSession) Rbac() (common.Address, error) {
	return _FuxLocker.Contract.Rbac(&_FuxLocker.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxLocker *FuxLockerTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _FuxLocker.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxLocker *FuxLockerSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxLocker.Contract.Initialize(&_FuxLocker.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxLocker *FuxLockerTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxLocker.Contract.Initialize(&_FuxLocker.TransactOpts, _rbac)
}

// SetDeltaT is a paid mutator transaction binding the contract method 0x52fc1d58.
//
// Solidity: function setDeltaT(_deltaT uint256) returns()
func (_FuxLocker *FuxLockerTransactor) SetDeltaT(opts *bind.TransactOpts, _deltaT *big.Int) (*types.Transaction, error) {
	return _FuxLocker.contract.Transact(opts, "setDeltaT", _deltaT)
}

// SetDeltaT is a paid mutator transaction binding the contract method 0x52fc1d58.
//
// Solidity: function setDeltaT(_deltaT uint256) returns()
func (_FuxLocker *FuxLockerSession) SetDeltaT(_deltaT *big.Int) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetDeltaT(&_FuxLocker.TransactOpts, _deltaT)
}

// SetDeltaT is a paid mutator transaction binding the contract method 0x52fc1d58.
//
// Solidity: function setDeltaT(_deltaT uint256) returns()
func (_FuxLocker *FuxLockerTransactorSession) SetDeltaT(_deltaT *big.Int) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetDeltaT(&_FuxLocker.TransactOpts, _deltaT)
}

// SetExpire is a paid mutator transaction binding the contract method 0xa8c5cf27.
//
// Solidity: function setExpire(_tokenId uint256, _expire uint256) returns()
func (_FuxLocker *FuxLockerTransactor) SetExpire(opts *bind.TransactOpts, _tokenId *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _FuxLocker.contract.Transact(opts, "setExpire", _tokenId, _expire)
}

// SetExpire is a paid mutator transaction binding the contract method 0xa8c5cf27.
//
// Solidity: function setExpire(_tokenId uint256, _expire uint256) returns()
func (_FuxLocker *FuxLockerSession) SetExpire(_tokenId *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetExpire(&_FuxLocker.TransactOpts, _tokenId, _expire)
}

// SetExpire is a paid mutator transaction binding the contract method 0xa8c5cf27.
//
// Solidity: function setExpire(_tokenId uint256, _expire uint256) returns()
func (_FuxLocker *FuxLockerTransactorSession) SetExpire(_tokenId *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetExpire(&_FuxLocker.TransactOpts, _tokenId, _expire)
}

// SetLockState is a paid mutator transaction binding the contract method 0x2a291e11.
//
// Solidity: function setLockState(_state uint256, _lock bool) returns()
func (_FuxLocker *FuxLockerTransactor) SetLockState(opts *bind.TransactOpts, _state *big.Int, _lock bool) (*types.Transaction, error) {
	return _FuxLocker.contract.Transact(opts, "setLockState", _state, _lock)
}

// SetLockState is a paid mutator transaction binding the contract method 0x2a291e11.
//
// Solidity: function setLockState(_state uint256, _lock bool) returns()
func (_FuxLocker *FuxLockerSession) SetLockState(_state *big.Int, _lock bool) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetLockState(&_FuxLocker.TransactOpts, _state, _lock)
}

// SetLockState is a paid mutator transaction binding the contract method 0x2a291e11.
//
// Solidity: function setLockState(_state uint256, _lock bool) returns()
func (_FuxLocker *FuxLockerTransactorSession) SetLockState(_state *big.Int, _lock bool) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetLockState(&_FuxLocker.TransactOpts, _state, _lock)
}

// SetState is a paid mutator transaction binding the contract method 0xb9d77bfc.
//
// Solidity: function setState(_tokenId uint256, _state uint256) returns()
func (_FuxLocker *FuxLockerTransactor) SetState(opts *bind.TransactOpts, _tokenId *big.Int, _state *big.Int) (*types.Transaction, error) {
	return _FuxLocker.contract.Transact(opts, "setState", _tokenId, _state)
}

// SetState is a paid mutator transaction binding the contract method 0xb9d77bfc.
//
// Solidity: function setState(_tokenId uint256, _state uint256) returns()
func (_FuxLocker *FuxLockerSession) SetState(_tokenId *big.Int, _state *big.Int) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetState(&_FuxLocker.TransactOpts, _tokenId, _state)
}

// SetState is a paid mutator transaction binding the contract method 0xb9d77bfc.
//
// Solidity: function setState(_tokenId uint256, _state uint256) returns()
func (_FuxLocker *FuxLockerTransactorSession) SetState(_tokenId *big.Int, _state *big.Int) (*types.Transaction, error) {
	return _FuxLocker.Contract.SetState(&_FuxLocker.TransactOpts, _tokenId, _state)
}

// FuxLockerDeltaTimeSetIterator is returned from FilterDeltaTimeSet and is used to iterate over the raw logs and unpacked data for DeltaTimeSet events raised by the FuxLocker contract.
type FuxLockerDeltaTimeSetIterator struct {
	Event *FuxLockerDeltaTimeSet // Event containing the contract specifics and raw log

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
func (it *FuxLockerDeltaTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxLockerDeltaTimeSet)
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
		it.Event = new(FuxLockerDeltaTimeSet)
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
func (it *FuxLockerDeltaTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxLockerDeltaTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxLockerDeltaTimeSet represents a DeltaTimeSet event raised by the FuxLocker contract.
type FuxLockerDeltaTimeSet struct {
	Caller common.Address
	DeltaT *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeltaTimeSet is a free log retrieval operation binding the contract event 0xacd17757a3133f061ef65177bb74cd35cf1e005becb93c65f73f9c95d4a98794.
//
// Solidity: event DeltaTimeSet(_caller address, _deltaT uint256)
func (_FuxLocker *FuxLockerFilterer) FilterDeltaTimeSet(opts *bind.FilterOpts) (*FuxLockerDeltaTimeSetIterator, error) {

	logs, sub, err := _FuxLocker.contract.FilterLogs(opts, "DeltaTimeSet")
	if err != nil {
		return nil, err
	}
	return &FuxLockerDeltaTimeSetIterator{contract: _FuxLocker.contract, event: "DeltaTimeSet", logs: logs, sub: sub}, nil
}

// WatchDeltaTimeSet is a free log subscription operation binding the contract event 0xacd17757a3133f061ef65177bb74cd35cf1e005becb93c65f73f9c95d4a98794.
//
// Solidity: event DeltaTimeSet(_caller address, _deltaT uint256)
func (_FuxLocker *FuxLockerFilterer) WatchDeltaTimeSet(opts *bind.WatchOpts, sink chan<- *FuxLockerDeltaTimeSet) (event.Subscription, error) {

	logs, sub, err := _FuxLocker.contract.WatchLogs(opts, "DeltaTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxLockerDeltaTimeSet)
				if err := _FuxLocker.contract.UnpackLog(event, "DeltaTimeSet", log); err != nil {
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

// FuxLockerExpireSetIterator is returned from FilterExpireSet and is used to iterate over the raw logs and unpacked data for ExpireSet events raised by the FuxLocker contract.
type FuxLockerExpireSetIterator struct {
	Event *FuxLockerExpireSet // Event containing the contract specifics and raw log

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
func (it *FuxLockerExpireSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxLockerExpireSet)
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
		it.Event = new(FuxLockerExpireSet)
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
func (it *FuxLockerExpireSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxLockerExpireSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxLockerExpireSet represents a ExpireSet event raised by the FuxLocker contract.
type FuxLockerExpireSet struct {
	Caller  common.Address
	TokenId *big.Int
	Expire  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExpireSet is a free log retrieval operation binding the contract event 0x78cd799f5d099bdd3586a6ab5a947eb80e932459410dc30146b47705e952de57.
//
// Solidity: event ExpireSet(_caller address, _tokenId uint256, _expire uint256)
func (_FuxLocker *FuxLockerFilterer) FilterExpireSet(opts *bind.FilterOpts) (*FuxLockerExpireSetIterator, error) {

	logs, sub, err := _FuxLocker.contract.FilterLogs(opts, "ExpireSet")
	if err != nil {
		return nil, err
	}
	return &FuxLockerExpireSetIterator{contract: _FuxLocker.contract, event: "ExpireSet", logs: logs, sub: sub}, nil
}

// WatchExpireSet is a free log subscription operation binding the contract event 0x78cd799f5d099bdd3586a6ab5a947eb80e932459410dc30146b47705e952de57.
//
// Solidity: event ExpireSet(_caller address, _tokenId uint256, _expire uint256)
func (_FuxLocker *FuxLockerFilterer) WatchExpireSet(opts *bind.WatchOpts, sink chan<- *FuxLockerExpireSet) (event.Subscription, error) {

	logs, sub, err := _FuxLocker.contract.WatchLogs(opts, "ExpireSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxLockerExpireSet)
				if err := _FuxLocker.contract.UnpackLog(event, "ExpireSet", log); err != nil {
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

// FuxLockerLockStateSetIterator is returned from FilterLockStateSet and is used to iterate over the raw logs and unpacked data for LockStateSet events raised by the FuxLocker contract.
type FuxLockerLockStateSetIterator struct {
	Event *FuxLockerLockStateSet // Event containing the contract specifics and raw log

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
func (it *FuxLockerLockStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxLockerLockStateSet)
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
		it.Event = new(FuxLockerLockStateSet)
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
func (it *FuxLockerLockStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxLockerLockStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxLockerLockStateSet represents a LockStateSet event raised by the FuxLocker contract.
type FuxLockerLockStateSet struct {
	Caller common.Address
	State  *big.Int
	Lock   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockStateSet is a free log retrieval operation binding the contract event 0xb8640e2eb4b26293f0742ae8635ee5e79149eda9dbc2d72a22730d96cdcae63b.
//
// Solidity: event LockStateSet(_caller address, _state uint256, _lock bool)
func (_FuxLocker *FuxLockerFilterer) FilterLockStateSet(opts *bind.FilterOpts) (*FuxLockerLockStateSetIterator, error) {

	logs, sub, err := _FuxLocker.contract.FilterLogs(opts, "LockStateSet")
	if err != nil {
		return nil, err
	}
	return &FuxLockerLockStateSetIterator{contract: _FuxLocker.contract, event: "LockStateSet", logs: logs, sub: sub}, nil
}

// WatchLockStateSet is a free log subscription operation binding the contract event 0xb8640e2eb4b26293f0742ae8635ee5e79149eda9dbc2d72a22730d96cdcae63b.
//
// Solidity: event LockStateSet(_caller address, _state uint256, _lock bool)
func (_FuxLocker *FuxLockerFilterer) WatchLockStateSet(opts *bind.WatchOpts, sink chan<- *FuxLockerLockStateSet) (event.Subscription, error) {

	logs, sub, err := _FuxLocker.contract.WatchLogs(opts, "LockStateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxLockerLockStateSet)
				if err := _FuxLocker.contract.UnpackLog(event, "LockStateSet", log); err != nil {
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

// FuxLockerMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FuxLocker contract.
type FuxLockerMigratedIterator struct {
	Event *FuxLockerMigrated // Event containing the contract specifics and raw log

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
func (it *FuxLockerMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxLockerMigrated)
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
		it.Event = new(FuxLockerMigrated)
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
func (it *FuxLockerMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxLockerMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxLockerMigrated represents a Migrated event raised by the FuxLocker contract.
type FuxLockerMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxLocker *FuxLockerFilterer) FilterMigrated(opts *bind.FilterOpts) (*FuxLockerMigratedIterator, error) {

	logs, sub, err := _FuxLocker.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &FuxLockerMigratedIterator{contract: _FuxLocker.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxLocker *FuxLockerFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FuxLockerMigrated) (event.Subscription, error) {

	logs, sub, err := _FuxLocker.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxLockerMigrated)
				if err := _FuxLocker.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// FuxLockerStateSetIterator is returned from FilterStateSet and is used to iterate over the raw logs and unpacked data for StateSet events raised by the FuxLocker contract.
type FuxLockerStateSetIterator struct {
	Event *FuxLockerStateSet // Event containing the contract specifics and raw log

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
func (it *FuxLockerStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxLockerStateSet)
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
		it.Event = new(FuxLockerStateSet)
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
func (it *FuxLockerStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxLockerStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxLockerStateSet represents a StateSet event raised by the FuxLocker contract.
type FuxLockerStateSet struct {
	Caller  common.Address
	TokenId *big.Int
	State   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStateSet is a free log retrieval operation binding the contract event 0x1d79a6e0d9992f96b2facb02f3f2430911c472085429997f740dd1c9fbcbc445.
//
// Solidity: event StateSet(_caller address, _tokenId uint256, _state uint256)
func (_FuxLocker *FuxLockerFilterer) FilterStateSet(opts *bind.FilterOpts) (*FuxLockerStateSetIterator, error) {

	logs, sub, err := _FuxLocker.contract.FilterLogs(opts, "StateSet")
	if err != nil {
		return nil, err
	}
	return &FuxLockerStateSetIterator{contract: _FuxLocker.contract, event: "StateSet", logs: logs, sub: sub}, nil
}

// WatchStateSet is a free log subscription operation binding the contract event 0x1d79a6e0d9992f96b2facb02f3f2430911c472085429997f740dd1c9fbcbc445.
//
// Solidity: event StateSet(_caller address, _tokenId uint256, _state uint256)
func (_FuxLocker *FuxLockerFilterer) WatchStateSet(opts *bind.WatchOpts, sink chan<- *FuxLockerStateSet) (event.Subscription, error) {

	logs, sub, err := _FuxLocker.contract.WatchLogs(opts, "StateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxLockerStateSet)
				if err := _FuxLocker.contract.UnpackLog(event, "StateSet", log); err != nil {
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
