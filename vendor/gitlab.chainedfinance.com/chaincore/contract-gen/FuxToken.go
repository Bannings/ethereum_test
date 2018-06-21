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

// FuxTokenABI is the input ABI used to generate the binding from.
const FuxTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"fuxCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fuxStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fuxLocker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_erc721token\",\"type\":\"address\"},{\"name\":\"_fuxCoin\",\"type\":\"address\"},{\"name\":\"_fuxStorage\",\"type\":\"address\"},{\"name\":\"_fuxLocker\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxToken is an auto generated Go binding around an Ethereum contract.
type FuxToken struct {
	FuxTokenCaller     // Read-only binding to the contract
	FuxTokenTransactor // Write-only binding to the contract
	FuxTokenFilterer   // Log filterer for contract events
}

// FuxTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxTokenSession struct {
	Contract     *FuxToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxTokenCallerSession struct {
	Contract *FuxTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxTokenTransactorSession struct {
	Contract     *FuxTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxTokenRaw struct {
	Contract *FuxToken // Generic contract binding to access the raw methods on
}

// FuxTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxTokenCallerRaw struct {
	Contract *FuxTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FuxTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxTokenTransactorRaw struct {
	Contract *FuxTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxToken creates a new instance of FuxToken, bound to a specific deployed contract.
func NewFuxToken(address common.Address, backend bind.ContractBackend) (*FuxToken, error) {
	contract, err := bindFuxToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxToken{FuxTokenCaller: FuxTokenCaller{contract: contract}, FuxTokenTransactor: FuxTokenTransactor{contract: contract}, FuxTokenFilterer: FuxTokenFilterer{contract: contract}}, nil
}

// NewFuxTokenCaller creates a new read-only instance of FuxToken, bound to a specific deployed contract.
func NewFuxTokenCaller(address common.Address, caller bind.ContractCaller) (*FuxTokenCaller, error) {
	contract, err := bindFuxToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxTokenCaller{contract: contract}, nil
}

// NewFuxTokenTransactor creates a new write-only instance of FuxToken, bound to a specific deployed contract.
func NewFuxTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxTokenTransactor, error) {
	contract, err := bindFuxToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxTokenTransactor{contract: contract}, nil
}

// NewFuxTokenFilterer creates a new log filterer instance of FuxToken, bound to a specific deployed contract.
func NewFuxTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxTokenFilterer, error) {
	contract, err := bindFuxToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxTokenFilterer{contract: contract}, nil
}

// bindFuxToken binds a generic wrapper to an already deployed contract.
func bindFuxToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxToken *FuxTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxToken.Contract.FuxTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxToken *FuxTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxToken.Contract.FuxTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxToken *FuxTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxToken.Contract.FuxTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxToken *FuxTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxToken *FuxTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxToken *FuxTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxToken.Contract.contract.Transact(opts, method, params...)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEADMIN() (string, error) {
	return _FuxToken.Contract.ROLEADMIN(&_FuxToken.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEADMIN() (string, error) {
	return _FuxToken.Contract.ROLEADMIN(&_FuxToken.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLECLUSTER() (string, error) {
	return _FuxToken.Contract.ROLECLUSTER(&_FuxToken.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLECLUSTER() (string, error) {
	return _FuxToken.Contract.ROLECLUSTER(&_FuxToken.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxToken *FuxTokenCaller) Erc721token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "erc721token")
	return *ret0, err
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxToken *FuxTokenSession) Erc721token() (common.Address, error) {
	return _FuxToken.Contract.Erc721token(&_FuxToken.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) Erc721token() (common.Address, error) {
	return _FuxToken.Contract.Erc721token(&_FuxToken.CallOpts)
}

// FuxCoin is a free data retrieval call binding the contract method 0x015c3e3d.
//
// Solidity: function fuxCoin() constant returns(address)
func (_FuxToken *FuxTokenCaller) FuxCoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "fuxCoin")
	return *ret0, err
}

// FuxCoin is a free data retrieval call binding the contract method 0x015c3e3d.
//
// Solidity: function fuxCoin() constant returns(address)
func (_FuxToken *FuxTokenSession) FuxCoin() (common.Address, error) {
	return _FuxToken.Contract.FuxCoin(&_FuxToken.CallOpts)
}

// FuxCoin is a free data retrieval call binding the contract method 0x015c3e3d.
//
// Solidity: function fuxCoin() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) FuxCoin() (common.Address, error) {
	return _FuxToken.Contract.FuxCoin(&_FuxToken.CallOpts)
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxToken *FuxTokenCaller) FuxLocker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "fuxLocker")
	return *ret0, err
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxToken *FuxTokenSession) FuxLocker() (common.Address, error) {
	return _FuxToken.Contract.FuxLocker(&_FuxToken.CallOpts)
}

// FuxLocker is a free data retrieval call binding the contract method 0xfd06bc9a.
//
// Solidity: function fuxLocker() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) FuxLocker() (common.Address, error) {
	return _FuxToken.Contract.FuxLocker(&_FuxToken.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxToken *FuxTokenCaller) FuxStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "fuxStorage")
	return *ret0, err
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxToken *FuxTokenSession) FuxStorage() (common.Address, error) {
	return _FuxToken.Contract.FuxStorage(&_FuxToken.CallOpts)
}

// FuxStorage is a free data retrieval call binding the contract method 0x3203e7a9.
//
// Solidity: function fuxStorage() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) FuxStorage() (common.Address, error) {
	return _FuxToken.Contract.FuxStorage(&_FuxToken.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxToken *FuxTokenSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxToken.Contract.IsMigrated(&_FuxToken.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _FuxToken.Contract.IsMigrated(&_FuxToken.CallOpts, contractName, migrationId)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxToken *FuxTokenCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxToken *FuxTokenSession) Rbac() (common.Address, error) {
	return _FuxToken.Contract.Rbac(&_FuxToken.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) Rbac() (common.Address, error) {
	return _FuxToken.Contract.Rbac(&_FuxToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) Burn(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "burn", _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Burn(&_FuxToken.TransactOpts, _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Burn(&_FuxToken.TransactOpts, _owner, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxToken *FuxTokenTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxToken *FuxTokenSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxToken.Contract.Initialize(&_FuxToken.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_FuxToken *FuxTokenTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _FuxToken.Contract.Initialize(&_FuxToken.TransactOpts, _rbac)
}

// Mint is a paid mutator transaction binding the contract method 0xf92883a2.
//
// Solidity: function mint(_to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256) returns()
func (_FuxToken *FuxTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _value *big.Int, _state *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "mint", _to, _tokenId, _value, _state, _expire)
}

// Mint is a paid mutator transaction binding the contract method 0xf92883a2.
//
// Solidity: function mint(_to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256) returns()
func (_FuxToken *FuxTokenSession) Mint(_to common.Address, _tokenId *big.Int, _value *big.Int, _state *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Mint(&_FuxToken.TransactOpts, _to, _tokenId, _value, _state, _expire)
}

// Mint is a paid mutator transaction binding the contract method 0xf92883a2.
//
// Solidity: function mint(_to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Mint(_to common.Address, _tokenId *big.Int, _value *big.Int, _state *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Mint(&_FuxToken.TransactOpts, _to, _tokenId, _value, _state, _expire)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.SafeTransferFrom(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.SafeTransferFrom(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}

// FuxTokenMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FuxToken contract.
type FuxTokenMigratedIterator struct {
	Event *FuxTokenMigrated // Event containing the contract specifics and raw log

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
func (it *FuxTokenMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxTokenMigrated)
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
		it.Event = new(FuxTokenMigrated)
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
func (it *FuxTokenMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxTokenMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxTokenMigrated represents a Migrated event raised by the FuxToken contract.
type FuxTokenMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxToken *FuxTokenFilterer) FilterMigrated(opts *bind.FilterOpts) (*FuxTokenMigratedIterator, error) {

	logs, sub, err := _FuxToken.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &FuxTokenMigratedIterator{contract: _FuxToken.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_FuxToken *FuxTokenFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FuxTokenMigrated) (event.Subscription, error) {

	logs, sub, err := _FuxToken.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxTokenMigrated)
				if err := _FuxToken.contract.UnpackLog(event, "Migrated", log); err != nil {
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
