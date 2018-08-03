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

// ZrlTokenABI is the input ABI used to generate the binding from.
const ZrlTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"autoTransferLock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zrlStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_LOCK_STATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_TRANSFER_FUNC_LOCK\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zrlLocker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zrlCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_erc721token\",\"type\":\"address\"},{\"name\":\"_zrlCoin\",\"type\":\"address\"},{\"name\":\"_zrlStorage\",\"type\":\"address\"},{\"name\":\"_zrlLocker\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"open\",\"type\":\"bool\"}],\"name\":\"setAutoTransferLock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_caller\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFromInCluster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZrlToken is an auto generated Go binding around an Ethereum contract.
type ZrlToken struct {
	ZrlTokenCaller     // Read-only binding to the contract
	ZrlTokenTransactor // Write-only binding to the contract
	ZrlTokenFilterer   // Log filterer for contract events
}

// ZrlTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrlTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrlTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZrlTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrlTokenSession struct {
	Contract     *ZrlToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrlTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrlTokenCallerSession struct {
	Contract *ZrlTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZrlTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrlTokenTransactorSession struct {
	Contract     *ZrlTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZrlTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrlTokenRaw struct {
	Contract *ZrlToken // Generic contract binding to access the raw methods on
}

// ZrlTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrlTokenCallerRaw struct {
	Contract *ZrlTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZrlTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrlTokenTransactorRaw struct {
	Contract *ZrlTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrlToken creates a new instance of ZrlToken, bound to a specific deployed contract.
func NewZrlToken(address common.Address, backend bind.ContractBackend) (*ZrlToken, error) {
	contract, err := bindZrlToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZrlToken{ZrlTokenCaller: ZrlTokenCaller{contract: contract}, ZrlTokenTransactor: ZrlTokenTransactor{contract: contract}, ZrlTokenFilterer: ZrlTokenFilterer{contract: contract}}, nil
}

// NewZrlTokenCaller creates a new read-only instance of ZrlToken, bound to a specific deployed contract.
func NewZrlTokenCaller(address common.Address, caller bind.ContractCaller) (*ZrlTokenCaller, error) {
	contract, err := bindZrlToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlTokenCaller{contract: contract}, nil
}

// NewZrlTokenTransactor creates a new write-only instance of ZrlToken, bound to a specific deployed contract.
func NewZrlTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrlTokenTransactor, error) {
	contract, err := bindZrlToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlTokenTransactor{contract: contract}, nil
}

// NewZrlTokenFilterer creates a new log filterer instance of ZrlToken, bound to a specific deployed contract.
func NewZrlTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZrlTokenFilterer, error) {
	contract, err := bindZrlToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZrlTokenFilterer{contract: contract}, nil
}

// bindZrlToken binds a generic wrapper to an already deployed contract.
func bindZrlToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrlTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlToken *ZrlTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlToken.Contract.ZrlTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlToken *ZrlTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlToken.Contract.ZrlTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlToken *ZrlTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlToken.Contract.ZrlTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlToken *ZrlTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlToken *ZrlTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlToken *ZrlTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlToken.Contract.contract.Transact(opts, method, params...)
}

// IDXTRANSFERFUNCLOCK is a free data retrieval call binding the contract method 0xc8f3a5f1.
//
// Solidity: function IDX_TRANSFER_FUNC_LOCK() constant returns(uint256)
func (_ZrlToken *ZrlTokenCaller) IDXTRANSFERFUNCLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "IDX_TRANSFER_FUNC_LOCK")
	return *ret0, err
}

// IDXTRANSFERFUNCLOCK is a free data retrieval call binding the contract method 0xc8f3a5f1.
//
// Solidity: function IDX_TRANSFER_FUNC_LOCK() constant returns(uint256)
func (_ZrlToken *ZrlTokenSession) IDXTRANSFERFUNCLOCK() (*big.Int, error) {
	return _ZrlToken.Contract.IDXTRANSFERFUNCLOCK(&_ZrlToken.CallOpts)
}

// IDXTRANSFERFUNCLOCK is a free data retrieval call binding the contract method 0xc8f3a5f1.
//
// Solidity: function IDX_TRANSFER_FUNC_LOCK() constant returns(uint256)
func (_ZrlToken *ZrlTokenCallerSession) IDXTRANSFERFUNCLOCK() (*big.Int, error) {
	return _ZrlToken.Contract.IDXTRANSFERFUNCLOCK(&_ZrlToken.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlToken *ZrlTokenCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlToken *ZrlTokenSession) ROLEADMIN() (string, error) {
	return _ZrlToken.Contract.ROLEADMIN(&_ZrlToken.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlToken *ZrlTokenCallerSession) ROLEADMIN() (string, error) {
	return _ZrlToken.Contract.ROLEADMIN(&_ZrlToken.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlToken *ZrlTokenCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlToken *ZrlTokenSession) ROLECLUSTER() (string, error) {
	return _ZrlToken.Contract.ROLECLUSTER(&_ZrlToken.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlToken *ZrlTokenCallerSession) ROLECLUSTER() (string, error) {
	return _ZrlToken.Contract.ROLECLUSTER(&_ZrlToken.CallOpts)
}

// TRANSFERLOCKSTATE is a free data retrieval call binding the contract method 0x337dd622.
//
// Solidity: function TRANSFER_LOCK_STATE() constant returns(uint256)
func (_ZrlToken *ZrlTokenCaller) TRANSFERLOCKSTATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "TRANSFER_LOCK_STATE")
	return *ret0, err
}

// TRANSFERLOCKSTATE is a free data retrieval call binding the contract method 0x337dd622.
//
// Solidity: function TRANSFER_LOCK_STATE() constant returns(uint256)
func (_ZrlToken *ZrlTokenSession) TRANSFERLOCKSTATE() (*big.Int, error) {
	return _ZrlToken.Contract.TRANSFERLOCKSTATE(&_ZrlToken.CallOpts)
}

// TRANSFERLOCKSTATE is a free data retrieval call binding the contract method 0x337dd622.
//
// Solidity: function TRANSFER_LOCK_STATE() constant returns(uint256)
func (_ZrlToken *ZrlTokenCallerSession) TRANSFERLOCKSTATE() (*big.Int, error) {
	return _ZrlToken.Contract.TRANSFERLOCKSTATE(&_ZrlToken.CallOpts)
}

// AutoTransferLock is a free data retrieval call binding the contract method 0x07564737.
//
// Solidity: function autoTransferLock() constant returns(bool)
func (_ZrlToken *ZrlTokenCaller) AutoTransferLock(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "autoTransferLock")
	return *ret0, err
}

// AutoTransferLock is a free data retrieval call binding the contract method 0x07564737.
//
// Solidity: function autoTransferLock() constant returns(bool)
func (_ZrlToken *ZrlTokenSession) AutoTransferLock() (bool, error) {
	return _ZrlToken.Contract.AutoTransferLock(&_ZrlToken.CallOpts)
}

// AutoTransferLock is a free data retrieval call binding the contract method 0x07564737.
//
// Solidity: function autoTransferLock() constant returns(bool)
func (_ZrlToken *ZrlTokenCallerSession) AutoTransferLock() (bool, error) {
	return _ZrlToken.Contract.AutoTransferLock(&_ZrlToken.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_ZrlToken *ZrlTokenCaller) Erc721token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "erc721token")
	return *ret0, err
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_ZrlToken *ZrlTokenSession) Erc721token() (common.Address, error) {
	return _ZrlToken.Contract.Erc721token(&_ZrlToken.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_ZrlToken *ZrlTokenCallerSession) Erc721token() (common.Address, error) {
	return _ZrlToken.Contract.Erc721token(&_ZrlToken.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlToken *ZrlTokenCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlToken *ZrlTokenSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlToken.Contract.IsMigrated(&_ZrlToken.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlToken *ZrlTokenCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlToken.Contract.IsMigrated(&_ZrlToken.CallOpts, contractName, migrationId)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlToken *ZrlTokenCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlToken *ZrlTokenSession) Rbac() (common.Address, error) {
	return _ZrlToken.Contract.Rbac(&_ZrlToken.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlToken *ZrlTokenCallerSession) Rbac() (common.Address, error) {
	return _ZrlToken.Contract.Rbac(&_ZrlToken.CallOpts)
}

// ZrlCoin is a free data retrieval call binding the contract method 0xdd356845.
//
// Solidity: function zrlCoin() constant returns(address)
func (_ZrlToken *ZrlTokenCaller) ZrlCoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "zrlCoin")
	return *ret0, err
}

// ZrlCoin is a free data retrieval call binding the contract method 0xdd356845.
//
// Solidity: function zrlCoin() constant returns(address)
func (_ZrlToken *ZrlTokenSession) ZrlCoin() (common.Address, error) {
	return _ZrlToken.Contract.ZrlCoin(&_ZrlToken.CallOpts)
}

// ZrlCoin is a free data retrieval call binding the contract method 0xdd356845.
//
// Solidity: function zrlCoin() constant returns(address)
func (_ZrlToken *ZrlTokenCallerSession) ZrlCoin() (common.Address, error) {
	return _ZrlToken.Contract.ZrlCoin(&_ZrlToken.CallOpts)
}

// ZrlLocker is a free data retrieval call binding the contract method 0xda20e984.
//
// Solidity: function zrlLocker() constant returns(address)
func (_ZrlToken *ZrlTokenCaller) ZrlLocker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "zrlLocker")
	return *ret0, err
}

// ZrlLocker is a free data retrieval call binding the contract method 0xda20e984.
//
// Solidity: function zrlLocker() constant returns(address)
func (_ZrlToken *ZrlTokenSession) ZrlLocker() (common.Address, error) {
	return _ZrlToken.Contract.ZrlLocker(&_ZrlToken.CallOpts)
}

// ZrlLocker is a free data retrieval call binding the contract method 0xda20e984.
//
// Solidity: function zrlLocker() constant returns(address)
func (_ZrlToken *ZrlTokenCallerSession) ZrlLocker() (common.Address, error) {
	return _ZrlToken.Contract.ZrlLocker(&_ZrlToken.CallOpts)
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlToken *ZrlTokenCaller) ZrlStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlToken.contract.Call(opts, out, "zrlStorage")
	return *ret0, err
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlToken *ZrlTokenSession) ZrlStorage() (common.Address, error) {
	return _ZrlToken.Contract.ZrlStorage(&_ZrlToken.CallOpts)
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlToken *ZrlTokenCallerSession) ZrlStorage() (common.Address, error) {
	return _ZrlToken.Contract.ZrlStorage(&_ZrlToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenTransactor) Burn(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.contract.Transact(opts, "burn", _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.Burn(&_ZrlToken.TransactOpts, _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenTransactorSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.Burn(&_ZrlToken.TransactOpts, _owner, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlToken *ZrlTokenTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _ZrlToken.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlToken *ZrlTokenSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _ZrlToken.Contract.Initialize(&_ZrlToken.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlToken *ZrlTokenTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _ZrlToken.Contract.Initialize(&_ZrlToken.TransactOpts, _rbac)
}

// Mint is a paid mutator transaction binding the contract method 0xf92883a2.
//
// Solidity: function mint(_to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256) returns()
func (_ZrlToken *ZrlTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _value *big.Int, _state *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _ZrlToken.contract.Transact(opts, "mint", _to, _tokenId, _value, _state, _expire)
}

// Mint is a paid mutator transaction binding the contract method 0xf92883a2.
//
// Solidity: function mint(_to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256) returns()
func (_ZrlToken *ZrlTokenSession) Mint(_to common.Address, _tokenId *big.Int, _value *big.Int, _state *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.Mint(&_ZrlToken.TransactOpts, _to, _tokenId, _value, _state, _expire)
}

// Mint is a paid mutator transaction binding the contract method 0xf92883a2.
//
// Solidity: function mint(_to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256) returns()
func (_ZrlToken *ZrlTokenTransactorSession) Mint(_to common.Address, _tokenId *big.Int, _value *big.Int, _state *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.Mint(&_ZrlToken.TransactOpts, _to, _tokenId, _value, _state, _expire)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.SafeTransferFrom(&_ZrlToken.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.SafeTransferFrom(&_ZrlToken.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFromInCluster is a paid mutator transaction binding the contract method 0xe9ce758c.
//
// Solidity: function safeTransferFromInCluster(_caller address, _from address, _to address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenTransactor) SafeTransferFromInCluster(opts *bind.TransactOpts, _caller common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.contract.Transact(opts, "safeTransferFromInCluster", _caller, _from, _to, _tokenId)
}

// SafeTransferFromInCluster is a paid mutator transaction binding the contract method 0xe9ce758c.
//
// Solidity: function safeTransferFromInCluster(_caller address, _from address, _to address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenSession) SafeTransferFromInCluster(_caller common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.SafeTransferFromInCluster(&_ZrlToken.TransactOpts, _caller, _from, _to, _tokenId)
}

// SafeTransferFromInCluster is a paid mutator transaction binding the contract method 0xe9ce758c.
//
// Solidity: function safeTransferFromInCluster(_caller address, _from address, _to address, _tokenId uint256) returns()
func (_ZrlToken *ZrlTokenTransactorSession) SafeTransferFromInCluster(_caller common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ZrlToken.Contract.SafeTransferFromInCluster(&_ZrlToken.TransactOpts, _caller, _from, _to, _tokenId)
}

// SetAutoTransferLock is a paid mutator transaction binding the contract method 0x84a29a28.
//
// Solidity: function setAutoTransferLock(open bool) returns()
func (_ZrlToken *ZrlTokenTransactor) SetAutoTransferLock(opts *bind.TransactOpts, open bool) (*types.Transaction, error) {
	return _ZrlToken.contract.Transact(opts, "setAutoTransferLock", open)
}

// SetAutoTransferLock is a paid mutator transaction binding the contract method 0x84a29a28.
//
// Solidity: function setAutoTransferLock(open bool) returns()
func (_ZrlToken *ZrlTokenSession) SetAutoTransferLock(open bool) (*types.Transaction, error) {
	return _ZrlToken.Contract.SetAutoTransferLock(&_ZrlToken.TransactOpts, open)
}

// SetAutoTransferLock is a paid mutator transaction binding the contract method 0x84a29a28.
//
// Solidity: function setAutoTransferLock(open bool) returns()
func (_ZrlToken *ZrlTokenTransactorSession) SetAutoTransferLock(open bool) (*types.Transaction, error) {
	return _ZrlToken.Contract.SetAutoTransferLock(&_ZrlToken.TransactOpts, open)
}

// ZrlTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the ZrlToken contract.
type ZrlTokenBurnedIterator struct {
	Event *ZrlTokenBurned // Event containing the contract specifics and raw log

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
func (it *ZrlTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlTokenBurned)
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
		it.Event = new(ZrlTokenBurned)
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
func (it *ZrlTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlTokenBurned represents a Burned event raised by the ZrlToken contract.
type ZrlTokenBurned struct {
	Caller  common.Address
	Owner   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x6ab368f832c266c8eb942b84fbcaa20aedc24a699d2a05fae2568028733b1d09.
//
// Solidity: event Burned(_caller address, _owner address, _tokenId uint256)
func (_ZrlToken *ZrlTokenFilterer) FilterBurned(opts *bind.FilterOpts) (*ZrlTokenBurnedIterator, error) {

	logs, sub, err := _ZrlToken.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &ZrlTokenBurnedIterator{contract: _ZrlToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x6ab368f832c266c8eb942b84fbcaa20aedc24a699d2a05fae2568028733b1d09.
//
// Solidity: event Burned(_caller address, _owner address, _tokenId uint256)
func (_ZrlToken *ZrlTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *ZrlTokenBurned) (event.Subscription, error) {

	logs, sub, err := _ZrlToken.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlTokenBurned)
				if err := _ZrlToken.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ZrlTokenMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the ZrlToken contract.
type ZrlTokenMigratedIterator struct {
	Event *ZrlTokenMigrated // Event containing the contract specifics and raw log

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
func (it *ZrlTokenMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlTokenMigrated)
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
		it.Event = new(ZrlTokenMigrated)
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
func (it *ZrlTokenMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlTokenMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlTokenMigrated represents a Migrated event raised by the ZrlToken contract.
type ZrlTokenMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ZrlToken *ZrlTokenFilterer) FilterMigrated(opts *bind.FilterOpts) (*ZrlTokenMigratedIterator, error) {

	logs, sub, err := _ZrlToken.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &ZrlTokenMigratedIterator{contract: _ZrlToken.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ZrlToken *ZrlTokenFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *ZrlTokenMigrated) (event.Subscription, error) {

	logs, sub, err := _ZrlToken.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlTokenMigrated)
				if err := _ZrlToken.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ZrlTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the ZrlToken contract.
type ZrlTokenMintedIterator struct {
	Event *ZrlTokenMinted // Event containing the contract specifics and raw log

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
func (it *ZrlTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlTokenMinted)
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
		it.Event = new(ZrlTokenMinted)
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
func (it *ZrlTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlTokenMinted represents a Minted event raised by the ZrlToken contract.
type ZrlTokenMinted struct {
	Caller  common.Address
	To      common.Address
	TokenId *big.Int
	Value   *big.Int
	State   *big.Int
	Expire  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x8f9cc53f42d35118d48a1cb6cff6d69109ae0898e83402061ffe1e4d05748be0.
//
// Solidity: event Minted(_caller address, _to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256)
func (_ZrlToken *ZrlTokenFilterer) FilterMinted(opts *bind.FilterOpts) (*ZrlTokenMintedIterator, error) {

	logs, sub, err := _ZrlToken.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &ZrlTokenMintedIterator{contract: _ZrlToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x8f9cc53f42d35118d48a1cb6cff6d69109ae0898e83402061ffe1e4d05748be0.
//
// Solidity: event Minted(_caller address, _to address, _tokenId uint256, _value uint256, _state uint256, _expire uint256)
func (_ZrlToken *ZrlTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ZrlTokenMinted) (event.Subscription, error) {

	logs, sub, err := _ZrlToken.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlTokenMinted)
				if err := _ZrlToken.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ZrlTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZrlToken contract.
type ZrlTokenTransferIterator struct {
	Event *ZrlTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ZrlTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlTokenTransfer)
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
		it.Event = new(ZrlTokenTransfer)
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
func (it *ZrlTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlTokenTransfer represents a Transfer event raised by the ZrlToken contract.
type ZrlTokenTransfer struct {
	Caller  common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(_caller address, _from address, _to address, _tokenId uint256)
func (_ZrlToken *ZrlTokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*ZrlTokenTransferIterator, error) {

	logs, sub, err := _ZrlToken.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &ZrlTokenTransferIterator{contract: _ZrlToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(_caller address, _from address, _to address, _tokenId uint256)
func (_ZrlToken *ZrlTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZrlTokenTransfer) (event.Subscription, error) {

	logs, sub, err := _ZrlToken.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlTokenTransfer)
				if err := _ZrlToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
