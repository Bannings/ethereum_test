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

// FuxSpliterABI is the input ABI used to generate the binding from.
const FuxSpliterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"fuxStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxKidCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ARCHIVE_FLAG\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fuxLocker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_kidIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"_kidValues\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"_kidStates\",\"type\":\"uint256[]\"}],\"name\":\"SplitToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_erc721token\",\"type\":\"address\"},{\"name\":\"_fuxStorage\",\"type\":\"address\"},{\"name\":\"_fuxLocker\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_count\",\"type\":\"uint16\"}],\"name\":\"setMaxKidCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_kidIds\",\"type\":\"uint256[]\"},{\"name\":\"_kidValues\",\"type\":\"uint256[]\"},{\"name\":\"_kidStates\",\"type\":\"uint256[]\"}],\"name\":\"split\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxSpliter is an auto generated Go binding around an Ethereum contract.
type FuxSpliter struct {
	FuxSpliterCaller     // Read-only binding to the contract
	FuxSpliterTransactor // Write-only binding to the contract
	FuxSpliterFilterer   // Log filterer for contract events
}

// FuxSpliterCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxSpliterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSpliterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxSpliterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSpliterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxSpliterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSpliterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxSpliterSession struct {
	Contract     *FuxSpliter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxSpliterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxSpliterCallerSession struct {
	Contract *FuxSpliterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FuxSpliterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxSpliterTransactorSession struct {
	Contract     *FuxSpliterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FuxSpliterRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxSpliterRaw struct {
	Contract *FuxSpliter // Generic contract binding to access the raw methods on
}

// FuxSpliterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxSpliterCallerRaw struct {
	Contract *FuxSpliterCaller // Generic read-only contract binding to access the raw methods on
}

// FuxSpliterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxSpliterTransactorRaw struct {
	Contract *FuxSpliterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxSpliter creates a new instance of FuxSpliter, bound to a specific deployed contract.
func NewFuxSpliter(address common.Address, backend bind.ContractBackend) (*FuxSpliter, error) {
	contract, err := bindFuxSpliter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxSpliter{FuxSpliterCaller: FuxSpliterCaller{contract: contract}, FuxSpliterTransactor: FuxSpliterTransactor{contract: contract}, FuxSpliterFilterer: FuxSpliterFilterer{contract: contract}}, nil
}

// NewFuxSpliterCaller creates a new read-only instance of FuxSpliter, bound to a specific deployed contract.
func NewFuxSpliterCaller(address common.Address, caller bind.ContractCaller) (*FuxSpliterCaller, error) {
	contract, err := bindFuxSpliter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSpliterCaller{contract: contract}, nil
}

// NewFuxSpliterTransactor creates a new write-only instance of FuxSpliter, bound to a specific deployed contract.
func NewFuxSpliterTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxSpliterTransactor, error) {
	contract, err := bindFuxSpliter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSpliterTransactor{contract: contract}, nil
}

// NewFuxSpliterFilterer creates a new log filterer instance of FuxSpliter, bound to a specific deployed contract.
func NewFuxSpliterFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxSpliterFilterer, error) {
	contract, err := bindFuxSpliter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxSpliterFilterer{contract: contract}, nil
}

// bindFuxSpliter binds a generic wrapper to an already deployed contract.
func bindFuxSpliter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxSpliterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSpliter *FuxSpliterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSpliter.Contract.FuxSpliterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSpliter *FuxSpliterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSpliter.Contract.FuxSpliterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSpliter *FuxSpliterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSpliter.Contract.FuxSpliterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSpliter *FuxSpliterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSpliter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSpliter *FuxSpliterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSpliter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSpliter *FuxSpliterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSpliter.Contract.contract.Transact(opts, method, params...)
}

// ARCHIVEFLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_FuxSpliter *FuxSpliterCaller) ARCHIVEFLAG(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "ARCHIVE_FLAG")
	return *ret0, err
}

// ARCHIVEFLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_FuxSpliter *FuxSpliterSession) ARCHIVEFLAG() (bool, error) {
	return _FuxSpliter.Contract.ARCHIVEFLAG(&_FuxSpliter.CallOpts)
}

// ARCHIVEFLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_FuxSpliter *FuxSpliterCallerSession) ARCHIVEFLAG() (bool, error) {
	return _FuxSpliter.Contract.ARCHIVEFLAG(&_FuxSpliter.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxSpliter *FuxSpliterCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxSpliter *FuxSpliterSession) ROLEADMIN() (string, error) {
	return _FuxSpliter.Contract.ROLEADMIN(&_FuxSpliter.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxSpliter *FuxSpliterCallerSession) ROLEADMIN() (string, error) {
	return _FuxSpliter.Contract.ROLEADMIN(&_FuxSpliter.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSpliter *FuxSpliterCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSpliter *FuxSpliterSession) ROLECLUSTER() (string, error) {
	return _FuxSpliter.Contract.ROLECLUSTER(&_FuxSpliter.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSpliter *FuxSpliterCallerSession) ROLECLUSTER() (string, error) {
	return _FuxSpliter.Contract.ROLECLUSTER(&_FuxSpliter.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxSpliter *FuxSpliterCaller) Erc721token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "erc721token")
	return *ret0, err
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxSpliter *FuxSpliterSession) Erc721token() (common.Address, error) {
	return _FuxSpliter.Contract.Erc721token(&_FuxSpliter.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxSpliter *FuxSpliterCallerSession) Erc721token() (common.Address, error) {
	return _FuxSpliter.Contract.Erc721token(&_FuxSpliter.CallOpts)
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxSpliter *FuxSpliterCaller) FuxLocker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "fuxLocker")
	return *ret0, err
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxSpliter *FuxSpliterSession) FuxLocker() (common.Address, error) {
	return _FuxSpliter.Contract.FuxLocker(&_FuxSpliter.CallOpts)
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxSpliter *FuxSpliterCallerSession) FuxLocker() (common.Address, error) {
	return _FuxSpliter.Contract.FuxLocker(&_FuxSpliter.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxSpliter *FuxSpliterCaller) FuxStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "fuxStorage")
	return *ret0, err
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxSpliter *FuxSpliterSession) FuxStorage() (common.Address, error) {
	return _FuxSpliter.Contract.FuxStorage(&_FuxSpliter.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxSpliter *FuxSpliterCallerSession) FuxStorage() (common.Address, error) {
	return _FuxSpliter.Contract.FuxStorage(&_FuxSpliter.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxSpliter *FuxSpliterCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxSpliter *FuxSpliterSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxSpliter.Contract.IsMigrated(&_FuxSpliter.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxSpliter *FuxSpliterCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxSpliter.Contract.IsMigrated(&_FuxSpliter.CallOpts, contractName, migrationId)
}

// MaxKidCount is a free data retrieval call binding the contract method 0xaa58969e.
//
// Solidity: function maxKidCount() constant returns(uint16)
func (_FuxSpliter *FuxSpliterCaller) MaxKidCount(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "maxKidCount")
	return *ret0, err
}

// MaxKidCount is a free data retrieval call binding the contract method 0xaa58969e.
//
// Solidity: function maxKidCount() constant returns(uint16)
func (_FuxSpliter *FuxSpliterSession) MaxKidCount() (uint16, error) {
	return _FuxSpliter.Contract.MaxKidCount(&_FuxSpliter.CallOpts)
}

// MaxKidCount is a free data retrieval call binding the contract method 0xaa58969e.
//
// Solidity: function maxKidCount() constant returns(uint16)
func (_FuxSpliter *FuxSpliterCallerSession) MaxKidCount() (uint16, error) {
	return _FuxSpliter.Contract.MaxKidCount(&_FuxSpliter.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxSpliter *FuxSpliterCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSpliter.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxSpliter *FuxSpliterSession) Rbac() (common.Address, error) {
	return _FuxSpliter.Contract.Rbac(&_FuxSpliter.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxSpliter *FuxSpliterCallerSession) Rbac() (common.Address, error) {
	return _FuxSpliter.Contract.Rbac(&_FuxSpliter.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _fuxStorage address, _fuxLocker address) returns()
func (_FuxSpliter *FuxSpliterTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address, _erc721token common.Address, _fuxStorage common.Address, _fuxLocker common.Address) (*types.Transaction, error) {
	return _FuxSpliter.contract.Transact(opts, "initialize", _rbac, _erc721token, _fuxStorage, _fuxLocker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _fuxStorage address, _fuxLocker address) returns()
func (_FuxSpliter *FuxSpliterSession) Initialize(_rbac common.Address, _erc721token common.Address, _fuxStorage common.Address, _fuxLocker common.Address) (*types.Transaction, error) {
	return _FuxSpliter.Contract.Initialize(&_FuxSpliter.TransactOpts, _rbac, _erc721token, _fuxStorage, _fuxLocker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _fuxStorage address, _fuxLocker address) returns()
func (_FuxSpliter *FuxSpliterTransactorSession) Initialize(_rbac common.Address, _erc721token common.Address, _fuxStorage common.Address, _fuxLocker common.Address) (*types.Transaction, error) {
	return _FuxSpliter.Contract.Initialize(&_FuxSpliter.TransactOpts, _rbac, _erc721token, _fuxStorage, _fuxLocker)
}

// SetMaxKidCount is a paid mutator transaction binding the contract method 0xba4c2cb0.
//
// Solidity: function setMaxKidCount(_count uint16) returns()
func (_FuxSpliter *FuxSpliterTransactor) SetMaxKidCount(opts *bind.TransactOpts, _count uint16) (*types.Transaction, error) {
	return _FuxSpliter.contract.Transact(opts, "setMaxKidCount", _count)
}

// SetMaxKidCount is a paid mutator transaction binding the contract method 0xba4c2cb0.
//
// Solidity: function setMaxKidCount(_count uint16) returns()
func (_FuxSpliter *FuxSpliterSession) SetMaxKidCount(_count uint16) (*types.Transaction, error) {
	return _FuxSpliter.Contract.SetMaxKidCount(&_FuxSpliter.TransactOpts, _count)
}

// SetMaxKidCount is a paid mutator transaction binding the contract method 0xba4c2cb0.
//
// Solidity: function setMaxKidCount(_count uint16) returns()
func (_FuxSpliter *FuxSpliterTransactorSession) SetMaxKidCount(_count uint16) (*types.Transaction, error) {
	return _FuxSpliter.Contract.SetMaxKidCount(&_FuxSpliter.TransactOpts, _count)
}

// Split is a paid mutator transaction binding the contract method 0x44d4b528.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[]) returns()
func (_FuxSpliter *FuxSpliterTransactor) Split(opts *bind.TransactOpts, _tokenId *big.Int, _kidIds []*big.Int, _kidValues []*big.Int, _kidStates []*big.Int) (*types.Transaction, error) {
	return _FuxSpliter.contract.Transact(opts, "split", _tokenId, _kidIds, _kidValues, _kidStates)
}

// Split is a paid mutator transaction binding the contract method 0x44d4b528.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[]) returns()
func (_FuxSpliter *FuxSpliterSession) Split(_tokenId *big.Int, _kidIds []*big.Int, _kidValues []*big.Int, _kidStates []*big.Int) (*types.Transaction, error) {
	return _FuxSpliter.Contract.Split(&_FuxSpliter.TransactOpts, _tokenId, _kidIds, _kidValues, _kidStates)
}

// Split is a paid mutator transaction binding the contract method 0x44d4b528.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[]) returns()
func (_FuxSpliter *FuxSpliterTransactorSession) Split(_tokenId *big.Int, _kidIds []*big.Int, _kidValues []*big.Int, _kidStates []*big.Int) (*types.Transaction, error) {
	return _FuxSpliter.Contract.Split(&_FuxSpliter.TransactOpts, _tokenId, _kidIds, _kidValues, _kidStates)
}

// FuxSpliterMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FuxSpliter contract.
type FuxSpliterMigratedIterator struct {
	Event *FuxSpliterMigrated // Event containing the contract specifics and raw log

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
func (it *FuxSpliterMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxSpliterMigrated)
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
		it.Event = new(FuxSpliterMigrated)
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
func (it *FuxSpliterMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxSpliterMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxSpliterMigrated represents a Migrated event raised by the FuxSpliter contract.
type FuxSpliterMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: e Migrated(contractName string, migrationId string)
func (_FuxSpliter *FuxSpliterFilterer) FilterMigrated(opts *bind.FilterOpts) (*FuxSpliterMigratedIterator, error) {

	logs, sub, err := _FuxSpliter.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &FuxSpliterMigratedIterator{contract: _FuxSpliter.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: e Migrated(contractName string, migrationId string)
func (_FuxSpliter *FuxSpliterFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FuxSpliterMigrated) (event.Subscription, error) {

	logs, sub, err := _FuxSpliter.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxSpliterMigrated)
				if err := _FuxSpliter.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// FuxSpliterSplitTokenIterator is returned from FilterSplitToken and is used to iterate over the raw logs and unpacked data for SplitToken events raised by the FuxSpliter contract.
type FuxSpliterSplitTokenIterator struct {
	Event *FuxSpliterSplitToken // Event containing the contract specifics and raw log

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
func (it *FuxSpliterSplitTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxSpliterSplitToken)
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
		it.Event = new(FuxSpliterSplitToken)
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
func (it *FuxSpliterSplitTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxSpliterSplitTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxSpliterSplitToken represents a SplitToken event raised by the FuxSpliter contract.
type FuxSpliterSplitToken struct {
	Caller    common.Address
	TokenId   *big.Int
	Owner     common.Address
	KidIds    []*big.Int
	KidValues []*big.Int
	KidStates []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSplitToken is a free log retrieval operation binding the contract event 0x90709a787fac1d8a6fe04f102351ddd5f293664af00bf9e6dce6968c3d33d4b6.
//
// Solidity: e SplitToken(_caller address, _tokenId uint256, _owner address, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[])
func (_FuxSpliter *FuxSpliterFilterer) FilterSplitToken(opts *bind.FilterOpts) (*FuxSpliterSplitTokenIterator, error) {

	logs, sub, err := _FuxSpliter.contract.FilterLogs(opts, "SplitToken")
	if err != nil {
		return nil, err
	}
	return &FuxSpliterSplitTokenIterator{contract: _FuxSpliter.contract, event: "SplitToken", logs: logs, sub: sub}, nil
}

// WatchSplitToken is a free log subscription operation binding the contract event 0x90709a787fac1d8a6fe04f102351ddd5f293664af00bf9e6dce6968c3d33d4b6.
//
// Solidity: e SplitToken(_caller address, _tokenId uint256, _owner address, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[])
func (_FuxSpliter *FuxSpliterFilterer) WatchSplitToken(opts *bind.WatchOpts, sink chan<- *FuxSpliterSplitToken) (event.Subscription, error) {

	logs, sub, err := _FuxSpliter.contract.WatchLogs(opts, "SplitToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxSpliterSplitToken)
				if err := _FuxSpliter.contract.UnpackLog(event, "SplitToken", log); err != nil {
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
