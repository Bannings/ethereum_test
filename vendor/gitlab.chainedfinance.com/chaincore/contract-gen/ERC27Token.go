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

// ERC27TokenABI is the input ABI used to generate the binding from.
const ERC27TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFromInCluster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_caller\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"checkApprovedOrOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC27Token is an auto generated Go binding around an Ethereum contract.
type ERC27Token struct {
	ERC27TokenCaller     // Read-only binding to the contract
	ERC27TokenTransactor // Write-only binding to the contract
	ERC27TokenFilterer   // Log filterer for contract events
}

// ERC27TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC27TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC27TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC27TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC27TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC27TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC27TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC27TokenSession struct {
	Contract     *ERC27Token       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC27TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC27TokenCallerSession struct {
	Contract *ERC27TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC27TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC27TokenTransactorSession struct {
	Contract     *ERC27TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC27TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC27TokenRaw struct {
	Contract *ERC27Token // Generic contract binding to access the raw methods on
}

// ERC27TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC27TokenCallerRaw struct {
	Contract *ERC27TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC27TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC27TokenTransactorRaw struct {
	Contract *ERC27TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC27Token creates a new instance of ERC27Token, bound to a specific deployed contract.
func NewERC27Token(address common.Address, backend bind.ContractBackend) (*ERC27Token, error) {
	contract, err := bindERC27Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC27Token{ERC27TokenCaller: ERC27TokenCaller{contract: contract}, ERC27TokenTransactor: ERC27TokenTransactor{contract: contract}, ERC27TokenFilterer: ERC27TokenFilterer{contract: contract}}, nil
}

// NewERC27TokenCaller creates a new read-only instance of ERC27Token, bound to a specific deployed contract.
func NewERC27TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC27TokenCaller, error) {
	contract, err := bindERC27Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC27TokenCaller{contract: contract}, nil
}

// NewERC27TokenTransactor creates a new write-only instance of ERC27Token, bound to a specific deployed contract.
func NewERC27TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC27TokenTransactor, error) {
	contract, err := bindERC27Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC27TokenTransactor{contract: contract}, nil
}

// NewERC27TokenFilterer creates a new log filterer instance of ERC27Token, bound to a specific deployed contract.
func NewERC27TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC27TokenFilterer, error) {
	contract, err := bindERC27Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC27TokenFilterer{contract: contract}, nil
}

// bindERC27Token binds a generic wrapper to an already deployed contract.
func bindERC27Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC27TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC27Token *ERC27TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC27Token.Contract.ERC27TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC27Token *ERC27TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC27Token.Contract.ERC27TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC27Token *ERC27TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC27Token.Contract.ERC27TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC27Token *ERC27TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC27Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC27Token *ERC27TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC27Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC27Token *ERC27TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC27Token.Contract.contract.Transact(opts, method, params...)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ERC27Token *ERC27TokenCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ERC27Token *ERC27TokenSession) ROLEADMIN() (string, error) {
	return _ERC27Token.Contract.ROLEADMIN(&_ERC27Token.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ERC27Token *ERC27TokenCallerSession) ROLEADMIN() (string, error) {
	return _ERC27Token.Contract.ROLEADMIN(&_ERC27Token.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ERC27Token *ERC27TokenCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ERC27Token *ERC27TokenSession) ROLECLUSTER() (string, error) {
	return _ERC27Token.Contract.ROLECLUSTER(&_ERC27Token.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ERC27Token *ERC27TokenCallerSession) ROLECLUSTER() (string, error) {
	return _ERC27Token.Contract.ROLECLUSTER(&_ERC27Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC27Token *ERC27TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC27Token *ERC27TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC27Token.Contract.BalanceOf(&_ERC27Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC27Token *ERC27TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC27Token.Contract.BalanceOf(&_ERC27Token.CallOpts, _owner)
}

// CheckApprovedOrOwner is a free data retrieval call binding the contract method 0x7e2a0290.
//
// Solidity: function checkApprovedOrOwner(_caller address, _tokenId uint256) constant returns(bool)
func (_ERC27Token *ERC27TokenCaller) CheckApprovedOrOwner(opts *bind.CallOpts, _caller common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "checkApprovedOrOwner", _caller, _tokenId)
	return *ret0, err
}

// CheckApprovedOrOwner is a free data retrieval call binding the contract method 0x7e2a0290.
//
// Solidity: function checkApprovedOrOwner(_caller address, _tokenId uint256) constant returns(bool)
func (_ERC27Token *ERC27TokenSession) CheckApprovedOrOwner(_caller common.Address, _tokenId *big.Int) (bool, error) {
	return _ERC27Token.Contract.CheckApprovedOrOwner(&_ERC27Token.CallOpts, _caller, _tokenId)
}

// CheckApprovedOrOwner is a free data retrieval call binding the contract method 0x7e2a0290.
//
// Solidity: function checkApprovedOrOwner(_caller address, _tokenId uint256) constant returns(bool)
func (_ERC27Token *ERC27TokenCallerSession) CheckApprovedOrOwner(_caller common.Address, _tokenId *big.Int) (bool, error) {
	return _ERC27Token.Contract.CheckApprovedOrOwner(&_ERC27Token.CallOpts, _caller, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC27Token *ERC27TokenCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC27Token *ERC27TokenSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC27Token.Contract.Exists(&_ERC27Token.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC27Token *ERC27TokenCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC27Token.Contract.Exists(&_ERC27Token.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC27Token *ERC27TokenCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC27Token *ERC27TokenSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC27Token.Contract.GetApproved(&_ERC27Token.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC27Token *ERC27TokenCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC27Token.Contract.GetApproved(&_ERC27Token.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC27Token *ERC27TokenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC27Token *ERC27TokenSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC27Token.Contract.IsApprovedForAll(&_ERC27Token.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC27Token *ERC27TokenCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC27Token.Contract.IsApprovedForAll(&_ERC27Token.CallOpts, _owner, _operator)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ERC27Token *ERC27TokenCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ERC27Token *ERC27TokenSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ERC27Token.Contract.IsMigrated(&_ERC27Token.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ERC27Token *ERC27TokenCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ERC27Token.Contract.IsMigrated(&_ERC27Token.CallOpts, contractName, migrationId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC27Token *ERC27TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC27Token *ERC27TokenSession) Name() (string, error) {
	return _ERC27Token.Contract.Name(&_ERC27Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC27Token *ERC27TokenCallerSession) Name() (string, error) {
	return _ERC27Token.Contract.Name(&_ERC27Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC27Token *ERC27TokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC27Token *ERC27TokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC27Token.Contract.OwnerOf(&_ERC27Token.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC27Token *ERC27TokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC27Token.Contract.OwnerOf(&_ERC27Token.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ERC27Token *ERC27TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ERC27Token *ERC27TokenSession) Paused() (bool, error) {
	return _ERC27Token.Contract.Paused(&_ERC27Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ERC27Token *ERC27TokenCallerSession) Paused() (bool, error) {
	return _ERC27Token.Contract.Paused(&_ERC27Token.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ERC27Token *ERC27TokenCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ERC27Token *ERC27TokenSession) Rbac() (common.Address, error) {
	return _ERC27Token.Contract.Rbac(&_ERC27Token.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ERC27Token *ERC27TokenCallerSession) Rbac() (common.Address, error) {
	return _ERC27Token.Contract.Rbac(&_ERC27Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC27Token *ERC27TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC27Token *ERC27TokenSession) Symbol() (string, error) {
	return _ERC27Token.Contract.Symbol(&_ERC27Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC27Token *ERC27TokenCallerSession) Symbol() (string, error) {
	return _ERC27Token.Contract.Symbol(&_ERC27Token.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC27Token *ERC27TokenCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC27Token *ERC27TokenSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC27Token.Contract.TokenByIndex(&_ERC27Token.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC27Token *ERC27TokenCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC27Token.Contract.TokenByIndex(&_ERC27Token.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC27Token *ERC27TokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC27Token *ERC27TokenSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC27Token.Contract.TokenOfOwnerByIndex(&_ERC27Token.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC27Token *ERC27TokenCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC27Token.Contract.TokenOfOwnerByIndex(&_ERC27Token.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC27Token *ERC27TokenCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC27Token *ERC27TokenSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC27Token.Contract.TokenURI(&_ERC27Token.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC27Token *ERC27TokenCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC27Token.Contract.TokenURI(&_ERC27Token.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC27Token *ERC27TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC27Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC27Token *ERC27TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC27Token.Contract.TotalSupply(&_ERC27Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC27Token *ERC27TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC27Token.Contract.TotalSupply(&_ERC27Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.Approve(&_ERC27Token.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.Approve(&_ERC27Token.TransactOpts, _to, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactor) Burn(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "burn", _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.Burn(&_ERC27Token.TransactOpts, _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactorSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.Burn(&_ERC27Token.TransactOpts, _owner, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(_rbac address, _name string, _symbol string) returns()
func (_ERC27Token *ERC27TokenTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "initialize", _rbac, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(_rbac address, _name string, _symbol string) returns()
func (_ERC27Token *ERC27TokenSession) Initialize(_rbac common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _ERC27Token.Contract.Initialize(&_ERC27Token.TransactOpts, _rbac, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(_rbac address, _name string, _symbol string) returns()
func (_ERC27Token *ERC27TokenTransactorSession) Initialize(_rbac common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _ERC27Token.Contract.Initialize(&_ERC27Token.TransactOpts, _rbac, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "mint", _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.Mint(&_ERC27Token.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactorSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.Mint(&_ERC27Token.TransactOpts, _to, _tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC27Token *ERC27TokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC27Token *ERC27TokenSession) Pause() (*types.Transaction, error) {
	return _ERC27Token.Contract.Pause(&_ERC27Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC27Token *ERC27TokenTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC27Token.Contract.Pause(&_ERC27Token.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.SafeTransferFrom(&_ERC27Token.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.SafeTransferFrom(&_ERC27Token.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFromInCluster is a paid mutator transaction binding the contract method 0x8ffc4c12.
//
// Solidity: function safeTransferFromInCluster(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactor) SafeTransferFromInCluster(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "safeTransferFromInCluster", _from, _to, _tokenId)
}

// SafeTransferFromInCluster is a paid mutator transaction binding the contract method 0x8ffc4c12.
//
// Solidity: function safeTransferFromInCluster(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenSession) SafeTransferFromInCluster(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.SafeTransferFromInCluster(&_ERC27Token.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFromInCluster is a paid mutator transaction binding the contract method 0x8ffc4c12.
//
// Solidity: function safeTransferFromInCluster(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactorSession) SafeTransferFromInCluster(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.SafeTransferFromInCluster(&_ERC27Token.TransactOpts, _from, _to, _tokenId)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC27Token *ERC27TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC27Token *ERC27TokenSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC27Token.Contract.SetApprovalForAll(&_ERC27Token.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC27Token *ERC27TokenTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC27Token.Contract.SetApprovalForAll(&_ERC27Token.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.TransferFrom(&_ERC27Token.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC27Token *ERC27TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC27Token.Contract.TransferFrom(&_ERC27Token.TransactOpts, _from, _to, _tokenId)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_ERC27Token *ERC27TokenTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC27Token.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_ERC27Token *ERC27TokenSession) UnPause() (*types.Transaction, error) {
	return _ERC27Token.Contract.UnPause(&_ERC27Token.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_ERC27Token *ERC27TokenTransactorSession) UnPause() (*types.Transaction, error) {
	return _ERC27Token.Contract.UnPause(&_ERC27Token.TransactOpts)
}

// ERC27TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC27Token contract.
type ERC27TokenApprovalIterator struct {
	Event *ERC27TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC27TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC27TokenApproval)
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
		it.Event = new(ERC27TokenApproval)
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
func (it *ERC27TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC27TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC27TokenApproval represents a Approval event raised by the ERC27Token contract.
type ERC27TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC27Token *ERC27TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC27TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC27Token.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC27TokenApprovalIterator{contract: _ERC27Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC27Token *ERC27TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC27TokenApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC27Token.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC27TokenApproval)
				if err := _ERC27Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC27TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC27Token contract.
type ERC27TokenApprovalForAllIterator struct {
	Event *ERC27TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC27TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC27TokenApprovalForAll)
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
		it.Event = new(ERC27TokenApprovalForAll)
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
func (it *ERC27TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC27TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC27TokenApprovalForAll represents a ApprovalForAll event raised by the ERC27Token contract.
type ERC27TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC27Token *ERC27TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC27TokenApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC27Token.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC27TokenApprovalForAllIterator{contract: _ERC27Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC27Token *ERC27TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC27TokenApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC27Token.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC27TokenApprovalForAll)
				if err := _ERC27Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC27TokenMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the ERC27Token contract.
type ERC27TokenMigratedIterator struct {
	Event *ERC27TokenMigrated // Event containing the contract specifics and raw log

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
func (it *ERC27TokenMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC27TokenMigrated)
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
		it.Event = new(ERC27TokenMigrated)
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
func (it *ERC27TokenMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC27TokenMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC27TokenMigrated represents a Migrated event raised by the ERC27Token contract.
type ERC27TokenMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ERC27Token *ERC27TokenFilterer) FilterMigrated(opts *bind.FilterOpts) (*ERC27TokenMigratedIterator, error) {

	logs, sub, err := _ERC27Token.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &ERC27TokenMigratedIterator{contract: _ERC27Token.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ERC27Token *ERC27TokenFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *ERC27TokenMigrated) (event.Subscription, error) {

	logs, sub, err := _ERC27Token.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC27TokenMigrated)
				if err := _ERC27Token.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ERC27TokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC27Token contract.
type ERC27TokenPausedIterator struct {
	Event *ERC27TokenPaused // Event containing the contract specifics and raw log

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
func (it *ERC27TokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC27TokenPaused)
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
		it.Event = new(ERC27TokenPaused)
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
func (it *ERC27TokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC27TokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC27TokenPaused represents a Paused event raised by the ERC27Token contract.
type ERC27TokenPaused struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(_caller address)
func (_ERC27Token *ERC27TokenFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC27TokenPausedIterator, error) {

	logs, sub, err := _ERC27Token.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC27TokenPausedIterator{contract: _ERC27Token.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(_caller address)
func (_ERC27Token *ERC27TokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC27TokenPaused) (event.Subscription, error) {

	logs, sub, err := _ERC27Token.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC27TokenPaused)
				if err := _ERC27Token.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ERC27TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC27Token contract.
type ERC27TokenTransferIterator struct {
	Event *ERC27TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC27TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC27TokenTransfer)
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
		it.Event = new(ERC27TokenTransfer)
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
func (it *ERC27TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC27TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC27TokenTransfer represents a Transfer event raised by the ERC27Token contract.
type ERC27TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC27Token *ERC27TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC27TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC27Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC27TokenTransferIterator{contract: _ERC27Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC27Token *ERC27TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC27TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC27Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC27TokenTransfer)
				if err := _ERC27Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC27TokenUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the ERC27Token contract.
type ERC27TokenUnPausedIterator struct {
	Event *ERC27TokenUnPaused // Event containing the contract specifics and raw log

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
func (it *ERC27TokenUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC27TokenUnPaused)
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
		it.Event = new(ERC27TokenUnPaused)
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
func (it *ERC27TokenUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC27TokenUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC27TokenUnPaused represents a UnPaused event raised by the ERC27Token contract.
type ERC27TokenUnPaused struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(_caller address)
func (_ERC27Token *ERC27TokenFilterer) FilterUnPaused(opts *bind.FilterOpts) (*ERC27TokenUnPausedIterator, error) {

	logs, sub, err := _ERC27Token.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &ERC27TokenUnPausedIterator{contract: _ERC27Token.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(_caller address)
func (_ERC27Token *ERC27TokenFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *ERC27TokenUnPaused) (event.Subscription, error) {

	logs, sub, err := _ERC27Token.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC27TokenUnPaused)
				if err := _ERC27Token.contract.UnpackLog(event, "UnPaused", log); err != nil {
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
