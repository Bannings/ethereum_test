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

// FuxStorageABI is the input ABI used to generate the binding from.
const FuxStorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"everExisted\",\"outputs\":[{\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"existedArchive\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"existed\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_val\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_createBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_root\",\"type\":\"uint256\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_val\",\"type\":\"uint256\"},{\"name\":\"_createBy\",\"type\":\"uint256\"},{\"name\":\"_root\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"archive\",\"type\":\"bool\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"archive\",\"type\":\"bool\"}],\"name\":\"getProperties\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"createBy\",\"type\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxStorage is an auto generated Go binding around an Ethereum contract.
type FuxStorage struct {
	FuxStorageCaller     // Read-only binding to the contract
	FuxStorageTransactor // Write-only binding to the contract
	FuxStorageFilterer   // Log filterer for contract events
}

// FuxStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxStorageSession struct {
	Contract     *FuxStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxStorageCallerSession struct {
	Contract *FuxStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FuxStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxStorageTransactorSession struct {
	Contract     *FuxStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FuxStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxStorageRaw struct {
	Contract *FuxStorage // Generic contract binding to access the raw methods on
}

// FuxStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxStorageCallerRaw struct {
	Contract *FuxStorageCaller // Generic read-only contract binding to access the raw methods on
}

// FuxStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxStorageTransactorRaw struct {
	Contract *FuxStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxStorage creates a new instance of FuxStorage, bound to a specific deployed contract.
func NewFuxStorage(address common.Address, backend bind.ContractBackend) (*FuxStorage, error) {
	contract, err := bindFuxStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxStorage{FuxStorageCaller: FuxStorageCaller{contract: contract}, FuxStorageTransactor: FuxStorageTransactor{contract: contract}, FuxStorageFilterer: FuxStorageFilterer{contract: contract}}, nil
}

// NewFuxStorageCaller creates a new read-only instance of FuxStorage, bound to a specific deployed contract.
func NewFuxStorageCaller(address common.Address, caller bind.ContractCaller) (*FuxStorageCaller, error) {
	contract, err := bindFuxStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxStorageCaller{contract: contract}, nil
}

// NewFuxStorageTransactor creates a new write-only instance of FuxStorage, bound to a specific deployed contract.
func NewFuxStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxStorageTransactor, error) {
	contract, err := bindFuxStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxStorageTransactor{contract: contract}, nil
}

// NewFuxStorageFilterer creates a new log filterer instance of FuxStorage, bound to a specific deployed contract.
func NewFuxStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxStorageFilterer, error) {
	contract, err := bindFuxStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxStorageFilterer{contract: contract}, nil
}

// bindFuxStorage binds a generic wrapper to an already deployed contract.
func bindFuxStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxStorage *FuxStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxStorage.Contract.FuxStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxStorage *FuxStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxStorage.Contract.FuxStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxStorage *FuxStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxStorage.Contract.FuxStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxStorage *FuxStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxStorage *FuxStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxStorage *FuxStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxStorage.Contract.contract.Transact(opts, method, params...)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxStorage *FuxStorageCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxStorage *FuxStorageSession) ROLEADMIN() (string, error) {
	return _FuxStorage.Contract.ROLEADMIN(&_FuxStorage.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxStorage *FuxStorageCallerSession) ROLEADMIN() (string, error) {
	return _FuxStorage.Contract.ROLEADMIN(&_FuxStorage.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxStorage *FuxStorageCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxStorage *FuxStorageSession) ROLECLUSTER() (string, error) {
	return _FuxStorage.Contract.ROLECLUSTER(&_FuxStorage.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxStorage *FuxStorageCallerSession) ROLECLUSTER() (string, error) {
	return _FuxStorage.Contract.ROLECLUSTER(&_FuxStorage.CallOpts)
}

// EverExisted is a free data retrieval call binding the contract method 0x258a0eea.
//
// Solidity: function everExisted(_tokenId uint256) constant returns(exist bool)
func (_FuxStorage *FuxStorageCaller) EverExisted(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "everExisted", _tokenId)
	return *ret0, err
}

// EverExisted is a free data retrieval call binding the contract method 0x258a0eea.
//
// Solidity: function everExisted(_tokenId uint256) constant returns(exist bool)
func (_FuxStorage *FuxStorageSession) EverExisted(_tokenId *big.Int) (bool, error) {
	return _FuxStorage.Contract.EverExisted(&_FuxStorage.CallOpts, _tokenId)
}

// EverExisted is a free data retrieval call binding the contract method 0x258a0eea.
//
// Solidity: function everExisted(_tokenId uint256) constant returns(exist bool)
func (_FuxStorage *FuxStorageCallerSession) EverExisted(_tokenId *big.Int) (bool, error) {
	return _FuxStorage.Contract.EverExisted(&_FuxStorage.CallOpts, _tokenId)
}

// Existed is a free data retrieval call binding the contract method 0xb8840616.
//
// Solidity: function existed(_tokenId uint256) constant returns(ret bool)
func (_FuxStorage *FuxStorageCaller) Existed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "existed", _tokenId)
	return *ret0, err
}

// Existed is a free data retrieval call binding the contract method 0xb8840616.
//
// Solidity: function existed(_tokenId uint256) constant returns(ret bool)
func (_FuxStorage *FuxStorageSession) Existed(_tokenId *big.Int) (bool, error) {
	return _FuxStorage.Contract.Existed(&_FuxStorage.CallOpts, _tokenId)
}

// Existed is a free data retrieval call binding the contract method 0xb8840616.
//
// Solidity: function existed(_tokenId uint256) constant returns(ret bool)
func (_FuxStorage *FuxStorageCallerSession) Existed(_tokenId *big.Int) (bool, error) {
	return _FuxStorage.Contract.Existed(&_FuxStorage.CallOpts, _tokenId)
}

// ExistedArchive is a free data retrieval call binding the contract method 0x5fccab8a.
//
// Solidity: function existedArchive(_tokenId uint256) constant returns(ret bool)
func (_FuxStorage *FuxStorageCaller) ExistedArchive(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "existedArchive", _tokenId)
	return *ret0, err
}

// ExistedArchive is a free data retrieval call binding the contract method 0x5fccab8a.
//
// Solidity: function existedArchive(_tokenId uint256) constant returns(ret bool)
func (_FuxStorage *FuxStorageSession) ExistedArchive(_tokenId *big.Int) (bool, error) {
	return _FuxStorage.Contract.ExistedArchive(&_FuxStorage.CallOpts, _tokenId)
}

// ExistedArchive is a free data retrieval call binding the contract method 0x5fccab8a.
//
// Solidity: function existedArchive(_tokenId uint256) constant returns(ret bool)
func (_FuxStorage *FuxStorageCallerSession) ExistedArchive(_tokenId *big.Int) (bool, error) {
	return _FuxStorage.Contract.ExistedArchive(&_FuxStorage.CallOpts, _tokenId)
}

// GetProperties is a free data retrieval call binding the contract method 0x091bcc9d.
//
// Solidity: function getProperties(_tokenId uint256, archive bool) constant returns(value uint256, createBy uint256, root uint256)
func (_FuxStorage *FuxStorageCaller) GetProperties(opts *bind.CallOpts, _tokenId *big.Int, archive bool) (struct {
	Value    *big.Int
	CreateBy *big.Int
	Root     *big.Int
}, error) {
	ret := new(struct {
		Value    *big.Int
		CreateBy *big.Int
		Root     *big.Int
	})
	out := ret
	err := _FuxStorage.contract.Call(opts, out, "getProperties", _tokenId, archive)
	return *ret, err
}

// GetProperties is a free data retrieval call binding the contract method 0x091bcc9d.
//
// Solidity: function getProperties(_tokenId uint256, archive bool) constant returns(value uint256, createBy uint256, root uint256)
func (_FuxStorage *FuxStorageSession) GetProperties(_tokenId *big.Int, archive bool) (struct {
	Value    *big.Int
	CreateBy *big.Int
	Root     *big.Int
}, error) {
	return _FuxStorage.Contract.GetProperties(&_FuxStorage.CallOpts, _tokenId, archive)
}

// GetProperties is a free data retrieval call binding the contract method 0x091bcc9d.
//
// Solidity: function getProperties(_tokenId uint256, archive bool) constant returns(value uint256, createBy uint256, root uint256)
func (_FuxStorage *FuxStorageCallerSession) GetProperties(_tokenId *big.Int, archive bool) (struct {
	Value    *big.Int
	CreateBy *big.Int
	Root     *big.Int
}, error) {
	return _FuxStorage.Contract.GetProperties(&_FuxStorage.CallOpts, _tokenId, archive)
}

// GetValue is a free data retrieval call binding the contract method 0x9e846204.
//
// Solidity: function getValue(_tokenId uint256, archive bool) constant returns(val uint256)
func (_FuxStorage *FuxStorageCaller) GetValue(opts *bind.CallOpts, _tokenId *big.Int, archive bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "getValue", _tokenId, archive)
	return *ret0, err
}

// GetValue is a free data retrieval call binding the contract method 0x9e846204.
//
// Solidity: function getValue(_tokenId uint256, archive bool) constant returns(val uint256)
func (_FuxStorage *FuxStorageSession) GetValue(_tokenId *big.Int, archive bool) (*big.Int, error) {
	return _FuxStorage.Contract.GetValue(&_FuxStorage.CallOpts, _tokenId, archive)
}

// GetValue is a free data retrieval call binding the contract method 0x9e846204.
//
// Solidity: function getValue(_tokenId uint256, archive bool) constant returns(val uint256)
func (_FuxStorage *FuxStorageCallerSession) GetValue(_tokenId *big.Int, archive bool) (*big.Int, error) {
	return _FuxStorage.Contract.GetValue(&_FuxStorage.CallOpts, _tokenId, archive)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxStorage *FuxStorageCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxStorage *FuxStorageSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxStorage.Contract.IsMigrated(&_FuxStorage.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxStorage *FuxStorageCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxStorage.Contract.IsMigrated(&_FuxStorage.CallOpts, contractName, migrationId)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxStorage *FuxStorageCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxStorage.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxStorage *FuxStorageSession) Rbac() (common.Address, error) {
	return _FuxStorage.Contract.Rbac(&_FuxStorage.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxStorage *FuxStorageCallerSession) Rbac() (common.Address, error) {
	return _FuxStorage.Contract.Rbac(&_FuxStorage.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x92848c9e.
//
// Solidity: function create(_tokenId uint256, _val uint256, _createBy uint256, _root uint256) returns()
func (_FuxStorage *FuxStorageTransactor) Create(opts *bind.TransactOpts, _tokenId *big.Int, _val *big.Int, _createBy *big.Int, _root *big.Int) (*types.Transaction, error) {
	return _FuxStorage.contract.Transact(opts, "create", _tokenId, _val, _createBy, _root)
}

// Create is a paid mutator transaction binding the contract method 0x92848c9e.
//
// Solidity: function create(_tokenId uint256, _val uint256, _createBy uint256, _root uint256) returns()
func (_FuxStorage *FuxStorageSession) Create(_tokenId *big.Int, _val *big.Int, _createBy *big.Int, _root *big.Int) (*types.Transaction, error) {
	return _FuxStorage.Contract.Create(&_FuxStorage.TransactOpts, _tokenId, _val, _createBy, _root)
}

// Create is a paid mutator transaction binding the contract method 0x92848c9e.
//
// Solidity: function create(_tokenId uint256, _val uint256, _createBy uint256, _root uint256) returns()
func (_FuxStorage *FuxStorageTransactorSession) Create(_tokenId *big.Int, _val *big.Int, _createBy *big.Int, _root *big.Int) (*types.Transaction, error) {
	return _FuxStorage.Contract.Create(&_FuxStorage.TransactOpts, _tokenId, _val, _createBy, _root)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxStorage *FuxStorageTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _FuxStorage.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxStorage *FuxStorageSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxStorage.Contract.Initialize(&_FuxStorage.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxStorage *FuxStorageTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxStorage.Contract.Initialize(&_FuxStorage.TransactOpts, _rbac)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(_tokenId uint256) returns()
func (_FuxStorage *FuxStorageTransactor) Remove(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxStorage.contract.Transact(opts, "remove", _tokenId)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(_tokenId uint256) returns()
func (_FuxStorage *FuxStorageSession) Remove(_tokenId *big.Int) (*types.Transaction, error) {
	return _FuxStorage.Contract.Remove(&_FuxStorage.TransactOpts, _tokenId)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(_tokenId uint256) returns()
func (_FuxStorage *FuxStorageTransactorSession) Remove(_tokenId *big.Int) (*types.Transaction, error) {
	return _FuxStorage.Contract.Remove(&_FuxStorage.TransactOpts, _tokenId)
}

// FuxStorageCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the FuxStorage contract.
type FuxStorageCreatedIterator struct {
	Event *FuxStorageCreated // Event containing the contract specifics and raw log

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
func (it *FuxStorageCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxStorageCreated)
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
		it.Event = new(FuxStorageCreated)
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
func (it *FuxStorageCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxStorageCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxStorageCreated represents a Created event raised by the FuxStorage contract.
type FuxStorageCreated struct {
	Caller   common.Address
	TokenId  *big.Int
	Val      *big.Int
	CreateBy *big.Int
	Root     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x7191c0804172ee8ac080b63e93fd6c2f26bbea256aed023d843c2a9f3dee0d3b.
//
// Solidity: event Created(_caller address, _tokenId uint256, _val uint256, _createBy uint256, _root uint256)
func (_FuxStorage *FuxStorageFilterer) FilterCreated(opts *bind.FilterOpts) (*FuxStorageCreatedIterator, error) {

	logs, sub, err := _FuxStorage.contract.FilterLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return &FuxStorageCreatedIterator{contract: _FuxStorage.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x7191c0804172ee8ac080b63e93fd6c2f26bbea256aed023d843c2a9f3dee0d3b.
//
// Solidity: event Created(_caller address, _tokenId uint256, _val uint256, _createBy uint256, _root uint256)
func (_FuxStorage *FuxStorageFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *FuxStorageCreated) (event.Subscription, error) {

	logs, sub, err := _FuxStorage.contract.WatchLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxStorageCreated)
				if err := _FuxStorage.contract.UnpackLog(event, "Created", log); err != nil {
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

// FuxStorageMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FuxStorage contract.
type FuxStorageMigratedIterator struct {
	Event *FuxStorageMigrated // Event containing the contract specifics and raw log

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
func (it *FuxStorageMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxStorageMigrated)
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
		it.Event = new(FuxStorageMigrated)
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
func (it *FuxStorageMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxStorageMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxStorageMigrated represents a Migrated event raised by the FuxStorage contract.
type FuxStorageMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxStorage *FuxStorageFilterer) FilterMigrated(opts *bind.FilterOpts) (*FuxStorageMigratedIterator, error) {

	logs, sub, err := _FuxStorage.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &FuxStorageMigratedIterator{contract: _FuxStorage.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxStorage *FuxStorageFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FuxStorageMigrated) (event.Subscription, error) {

	logs, sub, err := _FuxStorage.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxStorageMigrated)
				if err := _FuxStorage.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// FuxStorageRemovedIterator is returned from FilterRemoved and is used to iterate over the raw logs and unpacked data for Removed events raised by the FuxStorage contract.
type FuxStorageRemovedIterator struct {
	Event *FuxStorageRemoved // Event containing the contract specifics and raw log

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
func (it *FuxStorageRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxStorageRemoved)
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
		it.Event = new(FuxStorageRemoved)
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
func (it *FuxStorageRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxStorageRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxStorageRemoved represents a Removed event raised by the FuxStorage contract.
type FuxStorageRemoved struct {
	Caller  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoved is a free log retrieval operation binding the contract event 0xbe80a446a00b8794a7d05e8386915bdde937fe8b48da8d16175a5362b4c3f4f8.
//
// Solidity: event Removed(_caller address, _tokenId uint256)
func (_FuxStorage *FuxStorageFilterer) FilterRemoved(opts *bind.FilterOpts) (*FuxStorageRemovedIterator, error) {

	logs, sub, err := _FuxStorage.contract.FilterLogs(opts, "Removed")
	if err != nil {
		return nil, err
	}
	return &FuxStorageRemovedIterator{contract: _FuxStorage.contract, event: "Removed", logs: logs, sub: sub}, nil
}

// WatchRemoved is a free log subscription operation binding the contract event 0xbe80a446a00b8794a7d05e8386915bdde937fe8b48da8d16175a5362b4c3f4f8.
//
// Solidity: event Removed(_caller address, _tokenId uint256)
func (_FuxStorage *FuxStorageFilterer) WatchRemoved(opts *bind.WatchOpts, sink chan<- *FuxStorageRemoved) (event.Subscription, error) {

	logs, sub, err := _FuxStorage.contract.WatchLogs(opts, "Removed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxStorageRemoved)
				if err := _FuxStorage.contract.UnpackLog(event, "Removed", log); err != nil {
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
