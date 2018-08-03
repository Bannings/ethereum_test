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

// ZrlStorageABI is the input ABI used to generate the binding from.
const ZrlStorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"everExisted\",\"outputs\":[{\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"existedArchive\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"existed\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_val\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_createBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_root\",\"type\":\"uint256\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_val\",\"type\":\"uint256\"},{\"name\":\"_createBy\",\"type\":\"uint256\"},{\"name\":\"_root\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"archive\",\"type\":\"bool\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"archive\",\"type\":\"bool\"}],\"name\":\"getProperties\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"createBy\",\"type\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZrlStorage is an auto generated Go binding around an Ethereum contract.
type ZrlStorage struct {
	ZrlStorageCaller     // Read-only binding to the contract
	ZrlStorageTransactor // Write-only binding to the contract
	ZrlStorageFilterer   // Log filterer for contract events
}

// ZrlStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrlStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrlStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZrlStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrlStorageSession struct {
	Contract     *ZrlStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrlStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrlStorageCallerSession struct {
	Contract *ZrlStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZrlStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrlStorageTransactorSession struct {
	Contract     *ZrlStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZrlStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrlStorageRaw struct {
	Contract *ZrlStorage // Generic contract binding to access the raw methods on
}

// ZrlStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrlStorageCallerRaw struct {
	Contract *ZrlStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ZrlStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrlStorageTransactorRaw struct {
	Contract *ZrlStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrlStorage creates a new instance of ZrlStorage, bound to a specific deployed contract.
func NewZrlStorage(address common.Address, backend bind.ContractBackend) (*ZrlStorage, error) {
	contract, err := bindZrlStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZrlStorage{ZrlStorageCaller: ZrlStorageCaller{contract: contract}, ZrlStorageTransactor: ZrlStorageTransactor{contract: contract}, ZrlStorageFilterer: ZrlStorageFilterer{contract: contract}}, nil
}

// NewZrlStorageCaller creates a new read-only instance of ZrlStorage, bound to a specific deployed contract.
func NewZrlStorageCaller(address common.Address, caller bind.ContractCaller) (*ZrlStorageCaller, error) {
	contract, err := bindZrlStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlStorageCaller{contract: contract}, nil
}

// NewZrlStorageTransactor creates a new write-only instance of ZrlStorage, bound to a specific deployed contract.
func NewZrlStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrlStorageTransactor, error) {
	contract, err := bindZrlStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlStorageTransactor{contract: contract}, nil
}

// NewZrlStorageFilterer creates a new log filterer instance of ZrlStorage, bound to a specific deployed contract.
func NewZrlStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ZrlStorageFilterer, error) {
	contract, err := bindZrlStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZrlStorageFilterer{contract: contract}, nil
}

// bindZrlStorage binds a generic wrapper to an already deployed contract.
func bindZrlStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrlStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlStorage *ZrlStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlStorage.Contract.ZrlStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlStorage *ZrlStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlStorage.Contract.ZrlStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlStorage *ZrlStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlStorage.Contract.ZrlStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlStorage *ZrlStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlStorage *ZrlStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlStorage *ZrlStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlStorage.Contract.contract.Transact(opts, method, params...)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlStorage *ZrlStorageCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlStorage *ZrlStorageSession) ROLEADMIN() (string, error) {
	return _ZrlStorage.Contract.ROLEADMIN(&_ZrlStorage.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlStorage *ZrlStorageCallerSession) ROLEADMIN() (string, error) {
	return _ZrlStorage.Contract.ROLEADMIN(&_ZrlStorage.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlStorage *ZrlStorageCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlStorage *ZrlStorageSession) ROLECLUSTER() (string, error) {
	return _ZrlStorage.Contract.ROLECLUSTER(&_ZrlStorage.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlStorage *ZrlStorageCallerSession) ROLECLUSTER() (string, error) {
	return _ZrlStorage.Contract.ROLECLUSTER(&_ZrlStorage.CallOpts)
}

// EverExisted is a free data retrieval call binding the contract method 0x258a0eea.
//
// Solidity: function everExisted(_tokenId uint256) constant returns(exist bool)
func (_ZrlStorage *ZrlStorageCaller) EverExisted(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "everExisted", _tokenId)
	return *ret0, err
}

// EverExisted is a free data retrieval call binding the contract method 0x258a0eea.
//
// Solidity: function everExisted(_tokenId uint256) constant returns(exist bool)
func (_ZrlStorage *ZrlStorageSession) EverExisted(_tokenId *big.Int) (bool, error) {
	return _ZrlStorage.Contract.EverExisted(&_ZrlStorage.CallOpts, _tokenId)
}

// EverExisted is a free data retrieval call binding the contract method 0x258a0eea.
//
// Solidity: function everExisted(_tokenId uint256) constant returns(exist bool)
func (_ZrlStorage *ZrlStorageCallerSession) EverExisted(_tokenId *big.Int) (bool, error) {
	return _ZrlStorage.Contract.EverExisted(&_ZrlStorage.CallOpts, _tokenId)
}

// Existed is a free data retrieval call binding the contract method 0xb8840616.
//
// Solidity: function existed(_tokenId uint256) constant returns(ret bool)
func (_ZrlStorage *ZrlStorageCaller) Existed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "existed", _tokenId)
	return *ret0, err
}

// Existed is a free data retrieval call binding the contract method 0xb8840616.
//
// Solidity: function existed(_tokenId uint256) constant returns(ret bool)
func (_ZrlStorage *ZrlStorageSession) Existed(_tokenId *big.Int) (bool, error) {
	return _ZrlStorage.Contract.Existed(&_ZrlStorage.CallOpts, _tokenId)
}

// Existed is a free data retrieval call binding the contract method 0xb8840616.
//
// Solidity: function existed(_tokenId uint256) constant returns(ret bool)
func (_ZrlStorage *ZrlStorageCallerSession) Existed(_tokenId *big.Int) (bool, error) {
	return _ZrlStorage.Contract.Existed(&_ZrlStorage.CallOpts, _tokenId)
}

// ExistedArchive is a free data retrieval call binding the contract method 0x5fccab8a.
//
// Solidity: function existedArchive(_tokenId uint256) constant returns(ret bool)
func (_ZrlStorage *ZrlStorageCaller) ExistedArchive(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "existedArchive", _tokenId)
	return *ret0, err
}

// ExistedArchive is a free data retrieval call binding the contract method 0x5fccab8a.
//
// Solidity: function existedArchive(_tokenId uint256) constant returns(ret bool)
func (_ZrlStorage *ZrlStorageSession) ExistedArchive(_tokenId *big.Int) (bool, error) {
	return _ZrlStorage.Contract.ExistedArchive(&_ZrlStorage.CallOpts, _tokenId)
}

// ExistedArchive is a free data retrieval call binding the contract method 0x5fccab8a.
//
// Solidity: function existedArchive(_tokenId uint256) constant returns(ret bool)
func (_ZrlStorage *ZrlStorageCallerSession) ExistedArchive(_tokenId *big.Int) (bool, error) {
	return _ZrlStorage.Contract.ExistedArchive(&_ZrlStorage.CallOpts, _tokenId)
}

// GetProperties is a free data retrieval call binding the contract method 0x091bcc9d.
//
// Solidity: function getProperties(_tokenId uint256, archive bool) constant returns(value uint256, createBy uint256, root uint256)
func (_ZrlStorage *ZrlStorageCaller) GetProperties(opts *bind.CallOpts, _tokenId *big.Int, archive bool) (struct {
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
	err := _ZrlStorage.contract.Call(opts, out, "getProperties", _tokenId, archive)
	return *ret, err
}

// GetProperties is a free data retrieval call binding the contract method 0x091bcc9d.
//
// Solidity: function getProperties(_tokenId uint256, archive bool) constant returns(value uint256, createBy uint256, root uint256)
func (_ZrlStorage *ZrlStorageSession) GetProperties(_tokenId *big.Int, archive bool) (struct {
	Value    *big.Int
	CreateBy *big.Int
	Root     *big.Int
}, error) {
	return _ZrlStorage.Contract.GetProperties(&_ZrlStorage.CallOpts, _tokenId, archive)
}

// GetProperties is a free data retrieval call binding the contract method 0x091bcc9d.
//
// Solidity: function getProperties(_tokenId uint256, archive bool) constant returns(value uint256, createBy uint256, root uint256)
func (_ZrlStorage *ZrlStorageCallerSession) GetProperties(_tokenId *big.Int, archive bool) (struct {
	Value    *big.Int
	CreateBy *big.Int
	Root     *big.Int
}, error) {
	return _ZrlStorage.Contract.GetProperties(&_ZrlStorage.CallOpts, _tokenId, archive)
}

// GetValue is a free data retrieval call binding the contract method 0x9e846204.
//
// Solidity: function getValue(_tokenId uint256, archive bool) constant returns(val uint256)
func (_ZrlStorage *ZrlStorageCaller) GetValue(opts *bind.CallOpts, _tokenId *big.Int, archive bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "getValue", _tokenId, archive)
	return *ret0, err
}

// GetValue is a free data retrieval call binding the contract method 0x9e846204.
//
// Solidity: function getValue(_tokenId uint256, archive bool) constant returns(val uint256)
func (_ZrlStorage *ZrlStorageSession) GetValue(_tokenId *big.Int, archive bool) (*big.Int, error) {
	return _ZrlStorage.Contract.GetValue(&_ZrlStorage.CallOpts, _tokenId, archive)
}

// GetValue is a free data retrieval call binding the contract method 0x9e846204.
//
// Solidity: function getValue(_tokenId uint256, archive bool) constant returns(val uint256)
func (_ZrlStorage *ZrlStorageCallerSession) GetValue(_tokenId *big.Int, archive bool) (*big.Int, error) {
	return _ZrlStorage.Contract.GetValue(&_ZrlStorage.CallOpts, _tokenId, archive)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlStorage *ZrlStorageCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlStorage *ZrlStorageSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlStorage.Contract.IsMigrated(&_ZrlStorage.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlStorage *ZrlStorageCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlStorage.Contract.IsMigrated(&_ZrlStorage.CallOpts, contractName, migrationId)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlStorage *ZrlStorageCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlStorage.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlStorage *ZrlStorageSession) Rbac() (common.Address, error) {
	return _ZrlStorage.Contract.Rbac(&_ZrlStorage.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlStorage *ZrlStorageCallerSession) Rbac() (common.Address, error) {
	return _ZrlStorage.Contract.Rbac(&_ZrlStorage.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x92848c9e.
//
// Solidity: function create(_tokenId uint256, _val uint256, _createBy uint256, _root uint256) returns()
func (_ZrlStorage *ZrlStorageTransactor) Create(opts *bind.TransactOpts, _tokenId *big.Int, _val *big.Int, _createBy *big.Int, _root *big.Int) (*types.Transaction, error) {
	return _ZrlStorage.contract.Transact(opts, "create", _tokenId, _val, _createBy, _root)
}

// Create is a paid mutator transaction binding the contract method 0x92848c9e.
//
// Solidity: function create(_tokenId uint256, _val uint256, _createBy uint256, _root uint256) returns()
func (_ZrlStorage *ZrlStorageSession) Create(_tokenId *big.Int, _val *big.Int, _createBy *big.Int, _root *big.Int) (*types.Transaction, error) {
	return _ZrlStorage.Contract.Create(&_ZrlStorage.TransactOpts, _tokenId, _val, _createBy, _root)
}

// Create is a paid mutator transaction binding the contract method 0x92848c9e.
//
// Solidity: function create(_tokenId uint256, _val uint256, _createBy uint256, _root uint256) returns()
func (_ZrlStorage *ZrlStorageTransactorSession) Create(_tokenId *big.Int, _val *big.Int, _createBy *big.Int, _root *big.Int) (*types.Transaction, error) {
	return _ZrlStorage.Contract.Create(&_ZrlStorage.TransactOpts, _tokenId, _val, _createBy, _root)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlStorage *ZrlStorageTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _ZrlStorage.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlStorage *ZrlStorageSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _ZrlStorage.Contract.Initialize(&_ZrlStorage.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlStorage *ZrlStorageTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _ZrlStorage.Contract.Initialize(&_ZrlStorage.TransactOpts, _rbac)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(_tokenId uint256) returns()
func (_ZrlStorage *ZrlStorageTransactor) Remove(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlStorage.contract.Transact(opts, "remove", _tokenId)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(_tokenId uint256) returns()
func (_ZrlStorage *ZrlStorageSession) Remove(_tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlStorage.Contract.Remove(&_ZrlStorage.TransactOpts, _tokenId)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(_tokenId uint256) returns()
func (_ZrlStorage *ZrlStorageTransactorSession) Remove(_tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlStorage.Contract.Remove(&_ZrlStorage.TransactOpts, _tokenId)
}

// ZrlStorageCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the ZrlStorage contract.
type ZrlStorageCreatedIterator struct {
	Event *ZrlStorageCreated // Event containing the contract specifics and raw log

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
func (it *ZrlStorageCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlStorageCreated)
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
		it.Event = new(ZrlStorageCreated)
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
func (it *ZrlStorageCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlStorageCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlStorageCreated represents a Created event raised by the ZrlStorage contract.
type ZrlStorageCreated struct {
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
func (_ZrlStorage *ZrlStorageFilterer) FilterCreated(opts *bind.FilterOpts) (*ZrlStorageCreatedIterator, error) {

	logs, sub, err := _ZrlStorage.contract.FilterLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return &ZrlStorageCreatedIterator{contract: _ZrlStorage.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x7191c0804172ee8ac080b63e93fd6c2f26bbea256aed023d843c2a9f3dee0d3b.
//
// Solidity: event Created(_caller address, _tokenId uint256, _val uint256, _createBy uint256, _root uint256)
func (_ZrlStorage *ZrlStorageFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *ZrlStorageCreated) (event.Subscription, error) {

	logs, sub, err := _ZrlStorage.contract.WatchLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlStorageCreated)
				if err := _ZrlStorage.contract.UnpackLog(event, "Created", log); err != nil {
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

// ZrlStorageMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the ZrlStorage contract.
type ZrlStorageMigratedIterator struct {
	Event *ZrlStorageMigrated // Event containing the contract specifics and raw log

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
func (it *ZrlStorageMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlStorageMigrated)
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
		it.Event = new(ZrlStorageMigrated)
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
func (it *ZrlStorageMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlStorageMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlStorageMigrated represents a Migrated event raised by the ZrlStorage contract.
type ZrlStorageMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ZrlStorage *ZrlStorageFilterer) FilterMigrated(opts *bind.FilterOpts) (*ZrlStorageMigratedIterator, error) {

	logs, sub, err := _ZrlStorage.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &ZrlStorageMigratedIterator{contract: _ZrlStorage.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ZrlStorage *ZrlStorageFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *ZrlStorageMigrated) (event.Subscription, error) {

	logs, sub, err := _ZrlStorage.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlStorageMigrated)
				if err := _ZrlStorage.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ZrlStorageRemovedIterator is returned from FilterRemoved and is used to iterate over the raw logs and unpacked data for Removed events raised by the ZrlStorage contract.
type ZrlStorageRemovedIterator struct {
	Event *ZrlStorageRemoved // Event containing the contract specifics and raw log

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
func (it *ZrlStorageRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlStorageRemoved)
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
		it.Event = new(ZrlStorageRemoved)
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
func (it *ZrlStorageRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlStorageRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlStorageRemoved represents a Removed event raised by the ZrlStorage contract.
type ZrlStorageRemoved struct {
	Caller  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoved is a free log retrieval operation binding the contract event 0xbe80a446a00b8794a7d05e8386915bdde937fe8b48da8d16175a5362b4c3f4f8.
//
// Solidity: event Removed(_caller address, _tokenId uint256)
func (_ZrlStorage *ZrlStorageFilterer) FilterRemoved(opts *bind.FilterOpts) (*ZrlStorageRemovedIterator, error) {

	logs, sub, err := _ZrlStorage.contract.FilterLogs(opts, "Removed")
	if err != nil {
		return nil, err
	}
	return &ZrlStorageRemovedIterator{contract: _ZrlStorage.contract, event: "Removed", logs: logs, sub: sub}, nil
}

// WatchRemoved is a free log subscription operation binding the contract event 0xbe80a446a00b8794a7d05e8386915bdde937fe8b48da8d16175a5362b4c3f4f8.
//
// Solidity: event Removed(_caller address, _tokenId uint256)
func (_ZrlStorage *ZrlStorageFilterer) WatchRemoved(opts *bind.WatchOpts, sink chan<- *ZrlStorageRemoved) (event.Subscription, error) {

	logs, sub, err := _ZrlStorage.contract.WatchLogs(opts, "Removed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlStorageRemoved)
				if err := _ZrlStorage.contract.UnpackLog(event, "Removed", log); err != nil {
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
