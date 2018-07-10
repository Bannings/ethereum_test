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

// FuxBatchABI is the input ABI used to generate the binding from.
const FuxBatchABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fuxToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_fuxToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"updateMaxLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEADMIN() (string, error) {
	return _FuxBatch.Contract.ROLEADMIN(&_FuxBatch.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEADMIN() (string, error) {
	return _FuxBatch.Contract.ROLEADMIN(&_FuxBatch.CallOpts)
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

// FuxToken is a free data retrieval call binding the contract method 0xd91f5a6a.
//
// Solidity: function fuxToken() constant returns(address)
func (_FuxBatch *FuxBatchCaller) FuxToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "fuxToken")
	return *ret0, err
}

// FuxToken is a free data retrieval call binding the contract method 0xd91f5a6a.
//
// Solidity: function fuxToken() constant returns(address)
func (_FuxBatch *FuxBatchSession) FuxToken() (common.Address, error) {
	return _FuxBatch.Contract.FuxToken(&_FuxBatch.CallOpts)
}

// FuxToken is a free data retrieval call binding the contract method 0xd91f5a6a.
//
// Solidity: function fuxToken() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) FuxToken() (common.Address, error) {
	return _FuxBatch.Contract.FuxToken(&_FuxBatch.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxBatch.Contract.IsMigrated(&_FuxBatch.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxBatch.Contract.IsMigrated(&_FuxBatch.CallOpts, contractName, migrationId)
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

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxBatch *FuxBatchCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxBatch *FuxBatchSession) Rbac() (common.Address, error) {
	return _FuxBatch.Contract.Rbac(&_FuxBatch.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) Rbac() (common.Address, error) {
	return _FuxBatch.Contract.Rbac(&_FuxBatch.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxBatch *FuxBatchTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxBatch *FuxBatchSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxBatch.Contract.Initialize(&_FuxBatch.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxBatch *FuxBatchTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxBatch.Contract.Initialize(&_FuxBatch.TransactOpts, _rbac)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x24cf6f0f.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenIds uint256[]) returns()
func (_FuxBatch *FuxBatchTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenIds)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x24cf6f0f.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenIds uint256[]) returns()
func (_FuxBatch *FuxBatchSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SafeTransferFrom(&_FuxBatch.TransactOpts, _from, _to, _tokenIds)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x24cf6f0f.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenIds uint256[]) returns()
func (_FuxBatch *FuxBatchTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SafeTransferFrom(&_FuxBatch.TransactOpts, _from, _to, _tokenIds)
}

// UpdateMaxLength is a paid mutator transaction binding the contract method 0x23969d51.
//
// Solidity: function updateMaxLength(_length uint256) returns()
func (_FuxBatch *FuxBatchTransactor) UpdateMaxLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "updateMaxLength", _length)
}

// UpdateMaxLength is a paid mutator transaction binding the contract method 0x23969d51.
//
// Solidity: function updateMaxLength(_length uint256) returns()
func (_FuxBatch *FuxBatchSession) UpdateMaxLength(_length *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.UpdateMaxLength(&_FuxBatch.TransactOpts, _length)
}

// UpdateMaxLength is a paid mutator transaction binding the contract method 0x23969d51.
//
// Solidity: function updateMaxLength(_length uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) UpdateMaxLength(_length *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.UpdateMaxLength(&_FuxBatch.TransactOpts, _length)
}

// FuxBatchMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FuxBatch contract.
type FuxBatchMigratedIterator struct {
	Event *FuxBatchMigrated // Event containing the contract specifics and raw log

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
func (it *FuxBatchMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxBatchMigrated)
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
		it.Event = new(FuxBatchMigrated)
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
func (it *FuxBatchMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxBatchMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxBatchMigrated represents a Migrated event raised by the FuxBatch contract.
type FuxBatchMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: e Migrated(contractName string, migrationId string)
func (_FuxBatch *FuxBatchFilterer) FilterMigrated(opts *bind.FilterOpts) (*FuxBatchMigratedIterator, error) {

	logs, sub, err := _FuxBatch.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &FuxBatchMigratedIterator{contract: _FuxBatch.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: e Migrated(contractName string, migrationId string)
func (_FuxBatch *FuxBatchFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FuxBatchMigrated) (event.Subscription, error) {

	logs, sub, err := _FuxBatch.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxBatchMigrated)
				if err := _FuxBatch.contract.UnpackLog(event, "Migrated", log); err != nil {
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
