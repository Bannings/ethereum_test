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

// ZrlSpliterABI is the input ABI used to generate the binding from.
const ZrlSpliterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"zrlStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxKidCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zrlLocker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ARCHIVE_FLAG\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_kidIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"_kidValues\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"_kidStates\",\"type\":\"uint256[]\"}],\"name\":\"SplitToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_root\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_createBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"KidMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_erc721token\",\"type\":\"address\"},{\"name\":\"_zrlStorage\",\"type\":\"address\"},{\"name\":\"_zrlLocker\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_count\",\"type\":\"uint16\"}],\"name\":\"setMaxKidCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_kidIds\",\"type\":\"uint256[]\"},{\"name\":\"_kidValues\",\"type\":\"uint256[]\"},{\"name\":\"_kidStates\",\"type\":\"uint256[]\"}],\"name\":\"split\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZrlSpliter is an auto generated Go binding around an Ethereum contract.
type ZrlSpliter struct {
	ZrlSpliterCaller     // Read-only binding to the contract
	ZrlSpliterTransactor // Write-only binding to the contract
}

// ZrlSpliterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrlSpliterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlSpliterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrlSpliterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlSpliterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrlSpliterSession struct {
	Contract     *ZrlSpliter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrlSpliterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrlSpliterCallerSession struct {
	Contract *ZrlSpliterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZrlSpliterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrlSpliterTransactorSession struct {
	Contract     *ZrlSpliterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZrlSpliterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrlSpliterRaw struct {
	Contract *ZrlSpliter // Generic contract binding to access the raw methods on
}

// ZrlSpliterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrlSpliterCallerRaw struct {
	Contract *ZrlSpliterCaller // Generic read-only contract binding to access the raw methods on
}

// ZrlSpliterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrlSpliterTransactorRaw struct {
	Contract *ZrlSpliterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrlSpliter creates a new instance of ZrlSpliter, bound to a specific deployed contract.
func NewZrlSpliter(address common.Address, backend bind.ContractBackend) (*ZrlSpliter, error) {
	contract, err := bindZrlSpliter(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZrlSpliter{ZrlSpliterCaller: ZrlSpliterCaller{contract: contract}, ZrlSpliterTransactor: ZrlSpliterTransactor{contract: contract}}, nil
}

// NewZrlSpliterCaller creates a new read-only instance of ZrlSpliter, bound to a specific deployed contract.
func NewZrlSpliterCaller(address common.Address, caller bind.ContractCaller) (*ZrlSpliterCaller, error) {
	contract, err := bindZrlSpliter(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlSpliterCaller{contract: contract}, nil
}

// NewZrlSpliterTransactor creates a new write-only instance of ZrlSpliter, bound to a specific deployed contract.
func NewZrlSpliterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrlSpliterTransactor, error) {
	contract, err := bindZrlSpliter(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ZrlSpliterTransactor{contract: contract}, nil
}

// bindZrlSpliter binds a generic wrapper to an already deployed contract.
func bindZrlSpliter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrlSpliterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlSpliter *ZrlSpliterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlSpliter.Contract.ZrlSpliterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlSpliter *ZrlSpliterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.ZrlSpliterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlSpliter *ZrlSpliterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.ZrlSpliterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlSpliter *ZrlSpliterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlSpliter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlSpliter *ZrlSpliterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlSpliter *ZrlSpliterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.contract.Transact(opts, method, params...)
}

// ARCHIVE_FLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_ZrlSpliter *ZrlSpliterCaller) ARCHIVE_FLAG(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "ARCHIVE_FLAG")
	return *ret0, err
}

// ARCHIVE_FLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_ZrlSpliter *ZrlSpliterSession) ARCHIVE_FLAG() (bool, error) {
	return _ZrlSpliter.Contract.ARCHIVE_FLAG(&_ZrlSpliter.CallOpts)
}

// ARCHIVE_FLAG is a free data retrieval call binding the contract method 0xf8e782af.
//
// Solidity: function ARCHIVE_FLAG() constant returns(bool)
func (_ZrlSpliter *ZrlSpliterCallerSession) ARCHIVE_FLAG() (bool, error) {
	return _ZrlSpliter.Contract.ARCHIVE_FLAG(&_ZrlSpliter.CallOpts)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlSpliter *ZrlSpliterCaller) ROLE_ADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlSpliter *ZrlSpliterSession) ROLE_ADMIN() (string, error) {
	return _ZrlSpliter.Contract.ROLE_ADMIN(&_ZrlSpliter.CallOpts)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlSpliter *ZrlSpliterCallerSession) ROLE_ADMIN() (string, error) {
	return _ZrlSpliter.Contract.ROLE_ADMIN(&_ZrlSpliter.CallOpts)
}

// ROLE_CLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlSpliter *ZrlSpliterCaller) ROLE_CLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLE_CLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlSpliter *ZrlSpliterSession) ROLE_CLUSTER() (string, error) {
	return _ZrlSpliter.Contract.ROLE_CLUSTER(&_ZrlSpliter.CallOpts)
}

// ROLE_CLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlSpliter *ZrlSpliterCallerSession) ROLE_CLUSTER() (string, error) {
	return _ZrlSpliter.Contract.ROLE_CLUSTER(&_ZrlSpliter.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCaller) Erc721token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "erc721token")
	return *ret0, err
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_ZrlSpliter *ZrlSpliterSession) Erc721token() (common.Address, error) {
	return _ZrlSpliter.Contract.Erc721token(&_ZrlSpliter.CallOpts)
}

// Erc721token is a free data retrieval call binding the contract method 0xe5ebccf4.
//
// Solidity: function erc721token() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCallerSession) Erc721token() (common.Address, error) {
	return _ZrlSpliter.Contract.Erc721token(&_ZrlSpliter.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlSpliter *ZrlSpliterCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlSpliter *ZrlSpliterSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlSpliter.Contract.IsMigrated(&_ZrlSpliter.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlSpliter *ZrlSpliterCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlSpliter.Contract.IsMigrated(&_ZrlSpliter.CallOpts, contractName, migrationId)
}

// MaxKidCount is a free data retrieval call binding the contract method 0xaa58969e.
//
// Solidity: function maxKidCount() constant returns(uint16)
func (_ZrlSpliter *ZrlSpliterCaller) MaxKidCount(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "maxKidCount")
	return *ret0, err
}

// MaxKidCount is a free data retrieval call binding the contract method 0xaa58969e.
//
// Solidity: function maxKidCount() constant returns(uint16)
func (_ZrlSpliter *ZrlSpliterSession) MaxKidCount() (uint16, error) {
	return _ZrlSpliter.Contract.MaxKidCount(&_ZrlSpliter.CallOpts)
}

// MaxKidCount is a free data retrieval call binding the contract method 0xaa58969e.
//
// Solidity: function maxKidCount() constant returns(uint16)
func (_ZrlSpliter *ZrlSpliterCallerSession) MaxKidCount() (uint16, error) {
	return _ZrlSpliter.Contract.MaxKidCount(&_ZrlSpliter.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlSpliter *ZrlSpliterSession) Rbac() (common.Address, error) {
	return _ZrlSpliter.Contract.Rbac(&_ZrlSpliter.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCallerSession) Rbac() (common.Address, error) {
	return _ZrlSpliter.Contract.Rbac(&_ZrlSpliter.CallOpts)
}

// ZrlLocker is a free data retrieval call binding the contract method 0xda20e984.
//
// Solidity: function zrlLocker() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCaller) ZrlLocker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "zrlLocker")
	return *ret0, err
}

// ZrlLocker is a free data retrieval call binding the contract method 0xda20e984.
//
// Solidity: function zrlLocker() constant returns(address)
func (_ZrlSpliter *ZrlSpliterSession) ZrlLocker() (common.Address, error) {
	return _ZrlSpliter.Contract.ZrlLocker(&_ZrlSpliter.CallOpts)
}

// ZrlLocker is a free data retrieval call binding the contract method 0xda20e984.
//
// Solidity: function zrlLocker() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCallerSession) ZrlLocker() (common.Address, error) {
	return _ZrlSpliter.Contract.ZrlLocker(&_ZrlSpliter.CallOpts)
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCaller) ZrlStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlSpliter.contract.Call(opts, out, "zrlStorage")
	return *ret0, err
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlSpliter *ZrlSpliterSession) ZrlStorage() (common.Address, error) {
	return _ZrlSpliter.Contract.ZrlStorage(&_ZrlSpliter.CallOpts)
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlSpliter *ZrlSpliterCallerSession) ZrlStorage() (common.Address, error) {
	return _ZrlSpliter.Contract.ZrlStorage(&_ZrlSpliter.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _zrlStorage address, _zrlLocker address) returns()
func (_ZrlSpliter *ZrlSpliterTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address, _erc721token common.Address, _zrlStorage common.Address, _zrlLocker common.Address) (*types.Transaction, error) {
	return _ZrlSpliter.contract.Transact(opts, "initialize", _rbac, _erc721token, _zrlStorage, _zrlLocker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _zrlStorage address, _zrlLocker address) returns()
func (_ZrlSpliter *ZrlSpliterSession) Initialize(_rbac common.Address, _erc721token common.Address, _zrlStorage common.Address, _zrlLocker common.Address) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.Initialize(&_ZrlSpliter.TransactOpts, _rbac, _erc721token, _zrlStorage, _zrlLocker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_rbac address, _erc721token address, _zrlStorage address, _zrlLocker address) returns()
func (_ZrlSpliter *ZrlSpliterTransactorSession) Initialize(_rbac common.Address, _erc721token common.Address, _zrlStorage common.Address, _zrlLocker common.Address) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.Initialize(&_ZrlSpliter.TransactOpts, _rbac, _erc721token, _zrlStorage, _zrlLocker)
}

// SetMaxKidCount is a paid mutator transaction binding the contract method 0xba4c2cb0.
//
// Solidity: function setMaxKidCount(_count uint16) returns()
func (_ZrlSpliter *ZrlSpliterTransactor) SetMaxKidCount(opts *bind.TransactOpts, _count uint16) (*types.Transaction, error) {
	return _ZrlSpliter.contract.Transact(opts, "setMaxKidCount", _count)
}

// SetMaxKidCount is a paid mutator transaction binding the contract method 0xba4c2cb0.
//
// Solidity: function setMaxKidCount(_count uint16) returns()
func (_ZrlSpliter *ZrlSpliterSession) SetMaxKidCount(_count uint16) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.SetMaxKidCount(&_ZrlSpliter.TransactOpts, _count)
}

// SetMaxKidCount is a paid mutator transaction binding the contract method 0xba4c2cb0.
//
// Solidity: function setMaxKidCount(_count uint16) returns()
func (_ZrlSpliter *ZrlSpliterTransactorSession) SetMaxKidCount(_count uint16) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.SetMaxKidCount(&_ZrlSpliter.TransactOpts, _count)
}

// Split is a paid mutator transaction binding the contract method 0x44d4b528.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[]) returns()
func (_ZrlSpliter *ZrlSpliterTransactor) Split(opts *bind.TransactOpts, _tokenId *big.Int, _kidIds []*big.Int, _kidValues []*big.Int, _kidStates []*big.Int) (*types.Transaction, error) {
	return _ZrlSpliter.contract.Transact(opts, "split", _tokenId, _kidIds, _kidValues, _kidStates)
}

// Split is a paid mutator transaction binding the contract method 0x44d4b528.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[]) returns()
func (_ZrlSpliter *ZrlSpliterSession) Split(_tokenId *big.Int, _kidIds []*big.Int, _kidValues []*big.Int, _kidStates []*big.Int) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.Split(&_ZrlSpliter.TransactOpts, _tokenId, _kidIds, _kidValues, _kidStates)
}

// Split is a paid mutator transaction binding the contract method 0x44d4b528.
//
// Solidity: function split(_tokenId uint256, _kidIds uint256[], _kidValues uint256[], _kidStates uint256[]) returns()
func (_ZrlSpliter *ZrlSpliterTransactorSession) Split(_tokenId *big.Int, _kidIds []*big.Int, _kidValues []*big.Int, _kidStates []*big.Int) (*types.Transaction, error) {
	return _ZrlSpliter.Contract.Split(&_ZrlSpliter.TransactOpts, _tokenId, _kidIds, _kidValues, _kidStates)
}
