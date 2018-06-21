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

// FuxSplitABI is the input ABI used to generate the binding from.
const FuxSplitABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"fuxStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ARCHIVE_FLAG\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fuxLocker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_kidIds\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"name\":\"_kidValues\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"name\":\"_kidStates\",\"type\":\"uint256[2]\"}],\"name\":\"SplitToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_erc721token\",\"type\":\"address\"},{\"name\":\"_fuxStorage\",\"type\":\"address\"},{\"name\":\"_fuxLocker\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_kidIds\",\"type\":\"uint256[2]\"},{\"name\":\"_kidValues\",\"type\":\"uint256[2]\"},{\"name\":\"_kidStates\",\"type\":\"uint256[2]\"}],\"name\":\"split\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxSplit is an auto generated Go binding around an Ethereum contract.
type FuxSplit struct {
	FuxSplitCaller     // Read-only binding to the contract
	FuxSplitTransactor // Write-only binding to the contract
	FuxSplitFilterer   // Log filterer for contract events
}

// FuxSplitCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxSplitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxSplitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxSplitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxSplitSession struct {
	Contract     *FuxSplit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxSplitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxSplitCallerSession struct {
	Contract *FuxSplitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxSplitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxSplitTransactorSession struct {
	Contract     *FuxSplitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxSplitRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxSplitRaw struct {
	Contract *FuxSplit // Generic contract binding to access the raw methods on
}

// FuxSplitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxSplitCallerRaw struct {
	Contract *FuxSplitCaller // Generic read-only contract binding to access the raw methods on
}

// FuxSplitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxSplitTransactorRaw struct {
	Contract *FuxSplitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxSplit creates a new instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplit(address common.Address, backend bind.ContractBackend) (*FuxSplit, error) {
	contract, err := bindFuxSplit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxSplit{FuxSplitCaller: FuxSplitCaller{contract: contract}, FuxSplitTransactor: FuxSplitTransactor{contract: contract}, FuxSplitFilterer: FuxSplitFilterer{contract: contract}}, nil
}

// NewFuxSplitCaller creates a new read-only instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitCaller(address common.Address, caller bind.ContractCaller) (*FuxSplitCaller, error) {
	contract, err := bindFuxSplit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSplitCaller{contract: contract}, nil
}

// NewFuxSplitTransactor creates a new write-only instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxSplitTransactor, error) {
	contract, err := bindFuxSplit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSplitTransactor{contract: contract}, nil
}

// NewFuxSplitFilterer creates a new log filterer instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxSplitFilterer, error) {
	contract, err := bindFuxSplit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxSplitFilterer{contract: contract}, nil
}

// bindFuxSplit binds a generic wrapper to an already deployed contract.
func bindFuxSplit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxSplitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSplit *FuxSplitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSplit.Contract.FuxSplitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSplit *FuxSplitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSplit.Contract.FuxSplitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSplit *FuxSplitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSplit.Contract.FuxSplitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSplit *FuxSplitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSplit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSplit *FuxSplitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSplit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSplit *FuxSplitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSplit.Contract.contract.Transact(opts, method, params...)
}

// ARCHIVEFLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_FuxSplit *FuxSplitCaller) ARCHIVEFLAG(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ARCHIVE_FLAG")
	return *ret0, err
}

// ARCHIVEFLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_FuxSplit *FuxSplitSession) ARCHIVEFLAG() (bool, error) {
	return _FuxSplit.Contract.ARCHIVEFLAG(&_FuxSplit.CallOpts)
}

// ARCHIVEFLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) ARCHIVEFLAG() (bool, error) {
	return _FuxSplit.Contract.ARCHIVEFLAG(&_FuxSplit.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLEADMIN() (string, error) {
	return _FuxSplit.Contract.ROLEADMIN(&_FuxSplit.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLEADMIN() (string, error) {
	return _FuxSplit.Contract.ROLEADMIN(&_FuxSplit.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLECLUSTER() (string, error) {
	return _FuxSplit.Contract.ROLECLUSTER(&_FuxSplit.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLECLUSTER() (string, error) {
	return _FuxSplit.Contract.ROLECLUSTER(&_FuxSplit.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxSplit *FuxSplitCaller) Erc721token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "erc721token")
	return *ret0, err
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxSplit *FuxSplitSession) Erc721token() (common.Address, error) {
	return _FuxSplit.Contract.Erc721token(&_FuxSplit.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) Erc721token() (common.Address, error) {
	return _FuxSplit.Contract.Erc721token(&_FuxSplit.CallOpts)
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxSplit *FuxSplitCaller) FuxLocker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "fuxLocker")
	return *ret0, err
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxSplit *FuxSplitSession) FuxLocker() (common.Address, error) {
	return _FuxSplit.Contract.FuxLocker(&_FuxSplit.CallOpts)
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) FuxLocker() (common.Address, error) {
	return _FuxSplit.Contract.FuxLocker(&_FuxSplit.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxSplit *FuxSplitCaller) FuxStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "fuxStorage")
	return *ret0, err
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxSplit *FuxSplitSession) FuxStorage() (common.Address, error) {
	return _FuxSplit.Contract.FuxStorage(&_FuxSplit.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) FuxStorage() (common.Address, error) {
	return _FuxSplit.Contract.FuxStorage(&_FuxSplit.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxSplit *FuxSplitCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxSplit *FuxSplitSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxSplit.Contract.IsMigrated(&_FuxSplit.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxSplit.Contract.IsMigrated(&_FuxSplit.CallOpts, contractName, migrationId)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxSplit *FuxSplitCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxSplit *FuxSplitSession) Rbac() (common.Address, error) {
	return _FuxSplit.Contract.Rbac(&_FuxSplit.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) Rbac() (common.Address, error) {
	return _FuxSplit.Contract.Rbac(&_FuxSplit.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _fuxStorage address, _fuxLocker address) returns()
func (_FuxSplit *FuxSplitTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address, _erc721token common.Address, _fuxStorage common.Address, _fuxLocker common.Address) (*types.Transaction, error) {
	return _FuxSplit.contract.Transact(opts, "initialize", _rbac, _erc721token, _fuxStorage, _fuxLocker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _fuxStorage address, _fuxLocker address) returns()
func (_FuxSplit *FuxSplitSession) Initialize(_rbac common.Address, _erc721token common.Address, _fuxStorage common.Address, _fuxLocker common.Address) (*types.Transaction, error) {
	return _FuxSplit.Contract.Initialize(&_FuxSplit.TransactOpts, _rbac, _erc721token, _fuxStorage, _fuxLocker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _fuxStorage address, _fuxLocker address) returns()
func (_FuxSplit *FuxSplitTransactorSession) Initialize(_rbac common.Address, _erc721token common.Address, _fuxStorage common.Address, _fuxLocker common.Address) (*types.Transaction, error) {
	return _FuxSplit.Contract.Initialize(&_FuxSplit.TransactOpts, _rbac, _erc721token, _fuxStorage, _fuxLocker)
}

// Split is a paid mutator transaction binding the contract method 0x7f35d2f3.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[2], _kidValues uint256[2], _kidStates uint256[2]) returns()
func (_FuxSplit *FuxSplitTransactor) Split(opts *bind.TransactOpts, _tokenId *big.Int, _kidIds [2]*big.Int, _kidValues [2]*big.Int, _kidStates [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.contract.Transact(opts, "split", _tokenId, _kidIds, _kidValues, _kidStates)
}

// Split is a paid mutator transaction binding the contract method 0x7f35d2f3.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[2], _kidValues uint256[2], _kidStates uint256[2]) returns()
func (_FuxSplit *FuxSplitSession) Split(_tokenId *big.Int, _kidIds [2]*big.Int, _kidValues [2]*big.Int, _kidStates [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.Contract.Split(&_FuxSplit.TransactOpts, _tokenId, _kidIds, _kidValues, _kidStates)
}

// Split is a paid mutator transaction binding the contract method 0x7f35d2f3.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[2], _kidValues uint256[2], _kidStates uint256[2]) returns()
func (_FuxSplit *FuxSplitTransactorSession) Split(_tokenId *big.Int, _kidIds [2]*big.Int, _kidValues [2]*big.Int, _kidStates [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.Contract.Split(&_FuxSplit.TransactOpts, _tokenId, _kidIds, _kidValues, _kidStates)
}

// FuxSplitMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FuxSplit contract.
type FuxSplitMigratedIterator struct {
	Event *FuxSplitMigrated // Event containing the contract specifics and raw log

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
func (it *FuxSplitMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxSplitMigrated)
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
		it.Event = new(FuxSplitMigrated)
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
func (it *FuxSplitMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxSplitMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxSplitMigrated represents a Migrated event raised by the FuxSplit contract.
type FuxSplitMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxSplit *FuxSplitFilterer) FilterMigrated(opts *bind.FilterOpts) (*FuxSplitMigratedIterator, error) {

	logs, sub, err := _FuxSplit.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &FuxSplitMigratedIterator{contract: _FuxSplit.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxSplit *FuxSplitFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FuxSplitMigrated) (event.Subscription, error) {

	logs, sub, err := _FuxSplit.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxSplitMigrated)
				if err := _FuxSplit.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// FuxSplitSplitTokenIterator is returned from FilterSplitToken and is used to iterate over the raw logs and unpacked data for SplitToken events raised by the FuxSplit contract.
type FuxSplitSplitTokenIterator struct {
	Event *FuxSplitSplitToken // Event containing the contract specifics and raw log

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
func (it *FuxSplitSplitTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxSplitSplitToken)
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
		it.Event = new(FuxSplitSplitToken)
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
func (it *FuxSplitSplitTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxSplitSplitTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxSplitSplitToken represents a SplitToken event raised by the FuxSplit contract.
type FuxSplitSplitToken struct {
	Caller    common.Address
	TokenId   *big.Int
	Owner     common.Address
	KidIds    [2]*big.Int
	KidValues [2]*big.Int
	KidStates [2]*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSplitToken is a free log retrieval operation binding the contract event 0x267a65f439bfb39128d989ebb745250d8b7d0ccc4841369588f3bce5f3ff9316.
//
// Solidity: event SplitToken(_caller address, _tokenId uint256, _owner address, _kidIds uint256[2], _kidValues uint256[2], _kidStates uint256[2])
func (_FuxSplit *FuxSplitFilterer) FilterSplitToken(opts *bind.FilterOpts) (*FuxSplitSplitTokenIterator, error) {

	logs, sub, err := _FuxSplit.contract.FilterLogs(opts, "SplitToken")
	if err != nil {
		return nil, err
	}
	return &FuxSplitSplitTokenIterator{contract: _FuxSplit.contract, event: "SplitToken", logs: logs, sub: sub}, nil
}

// WatchSplitToken is a free log subscription operation binding the contract event 0x267a65f439bfb39128d989ebb745250d8b7d0ccc4841369588f3bce5f3ff9316.
//
// Solidity: event SplitToken(_caller address, _tokenId uint256, _owner address, _kidIds uint256[2], _kidValues uint256[2], _kidStates uint256[2])
func (_FuxSplit *FuxSplitFilterer) WatchSplitToken(opts *bind.WatchOpts, sink chan<- *FuxSplitSplitToken) (event.Subscription, error) {

	logs, sub, err := _FuxSplit.contract.WatchLogs(opts, "SplitToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxSplitSplitToken)
				if err := _FuxSplit.contract.UnpackLog(event, "SplitToken", log); err != nil {
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
