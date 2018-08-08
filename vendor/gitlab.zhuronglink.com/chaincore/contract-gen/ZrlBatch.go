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

// ZrlBatchABI is the input ABI used to generate the binding from.
const ZrlBatchABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"zrlToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_zrlToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"updateMaxLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tos\",\"type\":\"address[]\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"safeTransferFromToMulti\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZrlBatch is an auto generated Go binding around an Ethereum contract.
type ZrlBatch struct {
	ZrlBatchCaller     // Read-only binding to the contract
	ZrlBatchTransactor // Write-only binding to the contract
}

// ZrlBatchCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrlBatchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlBatchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrlBatchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlBatchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrlBatchSession struct {
	Contract     *ZrlBatch         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrlBatchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrlBatchCallerSession struct {
	Contract *ZrlBatchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZrlBatchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrlBatchTransactorSession struct {
	Contract     *ZrlBatchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZrlBatchRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrlBatchRaw struct {
	Contract *ZrlBatch // Generic contract binding to access the raw methods on
}

// ZrlBatchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrlBatchCallerRaw struct {
	Contract *ZrlBatchCaller // Generic read-only contract binding to access the raw methods on
}

// ZrlBatchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrlBatchTransactorRaw struct {
	Contract *ZrlBatchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrlBatch creates a new instance of ZrlBatch, bound to a specific deployed contract.
func NewZrlBatch(address common.Address, backend bind.ContractBackend) (*ZrlBatch, error) {
	contract, err := bindZrlBatch(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZrlBatch{ZrlBatchCaller: ZrlBatchCaller{contract: contract}, ZrlBatchTransactor: ZrlBatchTransactor{contract: contract}}, nil
}

// NewZrlBatchCaller creates a new read-only instance of ZrlBatch, bound to a specific deployed contract.
func NewZrlBatchCaller(address common.Address, caller bind.ContractCaller) (*ZrlBatchCaller, error) {
	contract, err := bindZrlBatch(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlBatchCaller{contract: contract}, nil
}

// NewZrlBatchTransactor creates a new write-only instance of ZrlBatch, bound to a specific deployed contract.
func NewZrlBatchTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrlBatchTransactor, error) {
	contract, err := bindZrlBatch(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ZrlBatchTransactor{contract: contract}, nil
}

// bindZrlBatch binds a generic wrapper to an already deployed contract.
func bindZrlBatch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrlBatchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlBatch *ZrlBatchRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlBatch.Contract.ZrlBatchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlBatch *ZrlBatchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlBatch.Contract.ZrlBatchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlBatch *ZrlBatchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlBatch.Contract.ZrlBatchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlBatch *ZrlBatchCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlBatch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlBatch *ZrlBatchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlBatch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlBatch *ZrlBatchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlBatch.Contract.contract.Transact(opts, method, params...)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlBatch *ZrlBatchCaller) ROLE_ADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlBatch.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlBatch *ZrlBatchSession) ROLE_ADMIN() (string, error) {
	return _ZrlBatch.Contract.ROLE_ADMIN(&_ZrlBatch.CallOpts)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlBatch *ZrlBatchCallerSession) ROLE_ADMIN() (string, error) {
	return _ZrlBatch.Contract.ROLE_ADMIN(&_ZrlBatch.CallOpts)
}

// ROLE_CLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlBatch *ZrlBatchCaller) ROLE_CLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlBatch.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLE_CLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlBatch *ZrlBatchSession) ROLE_CLUSTER() (string, error) {
	return _ZrlBatch.Contract.ROLE_CLUSTER(&_ZrlBatch.CallOpts)
}

// ROLE_CLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlBatch *ZrlBatchCallerSession) ROLE_CLUSTER() (string, error) {
	return _ZrlBatch.Contract.ROLE_CLUSTER(&_ZrlBatch.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlBatch *ZrlBatchCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlBatch.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlBatch *ZrlBatchSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlBatch.Contract.IsMigrated(&_ZrlBatch.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlBatch *ZrlBatchCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlBatch.Contract.IsMigrated(&_ZrlBatch.CallOpts, contractName, migrationId)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_ZrlBatch *ZrlBatchCaller) MaxLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlBatch.contract.Call(opts, out, "maxLength")
	return *ret0, err
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_ZrlBatch *ZrlBatchSession) MaxLength() (*big.Int, error) {
	return _ZrlBatch.Contract.MaxLength(&_ZrlBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_ZrlBatch *ZrlBatchCallerSession) MaxLength() (*big.Int, error) {
	return _ZrlBatch.Contract.MaxLength(&_ZrlBatch.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlBatch *ZrlBatchCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlBatch.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlBatch *ZrlBatchSession) Rbac() (common.Address, error) {
	return _ZrlBatch.Contract.Rbac(&_ZrlBatch.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlBatch *ZrlBatchCallerSession) Rbac() (common.Address, error) {
	return _ZrlBatch.Contract.Rbac(&_ZrlBatch.CallOpts)
}

// ZrlToken is a free data retrieval call binding the contract method 0x448900ba.
//
// Solidity: function zrlToken() constant returns(address)
func (_ZrlBatch *ZrlBatchCaller) ZrlToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlBatch.contract.Call(opts, out, "zrlToken")
	return *ret0, err
}

// ZrlToken is a free data retrieval call binding the contract method 0x448900ba.
//
// Solidity: function zrlToken() constant returns(address)
func (_ZrlBatch *ZrlBatchSession) ZrlToken() (common.Address, error) {
	return _ZrlBatch.Contract.ZrlToken(&_ZrlBatch.CallOpts)
}

// ZrlToken is a free data retrieval call binding the contract method 0x448900ba.
//
// Solidity: function zrlToken() constant returns(address)
func (_ZrlBatch *ZrlBatchCallerSession) ZrlToken() (common.Address, error) {
	return _ZrlBatch.Contract.ZrlToken(&_ZrlBatch.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlBatch *ZrlBatchTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address) (*types.Transaction, error) {
	return _ZrlBatch.contract.Transact(opts, "initialize", _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlBatch *ZrlBatchSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _ZrlBatch.Contract.Initialize(&_ZrlBatch.TransactOpts, _rbac)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_rbac address) returns()
func (_ZrlBatch *ZrlBatchTransactorSession) Initialize(_rbac common.Address) (*types.Transaction, error) {
	return _ZrlBatch.Contract.Initialize(&_ZrlBatch.TransactOpts, _rbac)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x24cf6f0f.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenIds uint256[]) returns()
func (_ZrlBatch *ZrlBatchTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ZrlBatch.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenIds)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x24cf6f0f.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenIds uint256[]) returns()
func (_ZrlBatch *ZrlBatchSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ZrlBatch.Contract.SafeTransferFrom(&_ZrlBatch.TransactOpts, _from, _to, _tokenIds)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x24cf6f0f.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenIds uint256[]) returns()
func (_ZrlBatch *ZrlBatchTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ZrlBatch.Contract.SafeTransferFrom(&_ZrlBatch.TransactOpts, _from, _to, _tokenIds)
}

// SafeTransferFromToMulti is a paid mutator transaction binding the contract method 0x7009e6bd.
//
// Solidity: function safeTransferFromToMulti(_from address, _tos address[], _tokenIds uint256[]) returns()
func (_ZrlBatch *ZrlBatchTransactor) SafeTransferFromToMulti(opts *bind.TransactOpts, _from common.Address, _tos []common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ZrlBatch.contract.Transact(opts, "safeTransferFromToMulti", _from, _tos, _tokenIds)
}

// SafeTransferFromToMulti is a paid mutator transaction binding the contract method 0x7009e6bd.
//
// Solidity: function safeTransferFromToMulti(_from address, _tos address[], _tokenIds uint256[]) returns()
func (_ZrlBatch *ZrlBatchSession) SafeTransferFromToMulti(_from common.Address, _tos []common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ZrlBatch.Contract.SafeTransferFromToMulti(&_ZrlBatch.TransactOpts, _from, _tos, _tokenIds)
}

// SafeTransferFromToMulti is a paid mutator transaction binding the contract method 0x7009e6bd.
//
// Solidity: function safeTransferFromToMulti(_from address, _tos address[], _tokenIds uint256[]) returns()
func (_ZrlBatch *ZrlBatchTransactorSession) SafeTransferFromToMulti(_from common.Address, _tos []common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ZrlBatch.Contract.SafeTransferFromToMulti(&_ZrlBatch.TransactOpts, _from, _tos, _tokenIds)
}

// UpdateMaxLength is a paid mutator transaction binding the contract method 0x23969d51.
//
// Solidity: function updateMaxLength(_length uint256) returns()
func (_ZrlBatch *ZrlBatchTransactor) UpdateMaxLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _ZrlBatch.contract.Transact(opts, "updateMaxLength", _length)
}

// UpdateMaxLength is a paid mutator transaction binding the contract method 0x23969d51.
//
// Solidity: function updateMaxLength(_length uint256) returns()
func (_ZrlBatch *ZrlBatchSession) UpdateMaxLength(_length *big.Int) (*types.Transaction, error) {
	return _ZrlBatch.Contract.UpdateMaxLength(&_ZrlBatch.TransactOpts, _length)
}

// UpdateMaxLength is a paid mutator transaction binding the contract method 0x23969d51.
//
// Solidity: function updateMaxLength(_length uint256) returns()
func (_ZrlBatch *ZrlBatchTransactorSession) UpdateMaxLength(_length *big.Int) (*types.Transaction, error) {
	return _ZrlBatch.Contract.UpdateMaxLength(&_ZrlBatch.TransactOpts, _length)
}
