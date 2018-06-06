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

// FuxTokenABI is the input ABI used to generate the binding from.
const FuxTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isFuxHub\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInFuxGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_coin\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"fuxBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEFUXGROUP() (string, error) {
	return _FuxToken.Contract.ROLEFUXGROUP(&_FuxToken.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxToken.Contract.ROLEFUXGROUP(&_FuxToken.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEFUXHUB() (string, error) {
	return _FuxToken.Contract.ROLEFUXHUB(&_FuxToken.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxToken.Contract.ROLEFUXHUB(&_FuxToken.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTECFRBAC() (string, error) {
	return _FuxToken.Contract.ROUTECFRBAC(&_FuxToken.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxToken.Contract.ROUTECFRBAC(&_FuxToken.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTECONFIG() (string, error) {
	return _FuxToken.Contract.ROUTECONFIG(&_FuxToken.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTECONFIG() (string, error) {
	return _FuxToken.Contract.ROUTECONFIG(&_FuxToken.CallOpts)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsFuxHub(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "_isFuxHub", _addr)
	return *ret0, err
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsFuxHub(&_FuxToken.CallOpts, _addr)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsFuxHub(&_FuxToken.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsInFuxGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "_isInFuxGroup", _addr)
	return *ret0, err
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsInFuxGroup(&_FuxToken.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsInFuxGroup(&_FuxToken.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.BalanceOf(&_FuxToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.BalanceOf(&_FuxToken.CallOpts, _owner)
}

// FuxBalanceOf is a free data retrieval call binding the contract method 0xd75ab616.
//
// Solidity: function fuxBalanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCaller) FuxBalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "fuxBalanceOf", _owner)
	return *ret0, err
}

// FuxBalanceOf is a free data retrieval call binding the contract method 0xd75ab616.
//
// Solidity: function fuxBalanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenSession) FuxBalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.FuxBalanceOf(&_FuxToken.CallOpts, _owner)
}

// FuxBalanceOf is a free data retrieval call binding the contract method 0xd75ab616.
//
// Solidity: function fuxBalanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCallerSession) FuxBalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.FuxBalanceOf(&_FuxToken.CallOpts, _owner)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxToken *FuxTokenCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxToken *FuxTokenSession) GetRouterAddress() (common.Address, error) {
	return _FuxToken.Contract.GetRouterAddress(&_FuxToken.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxToken.Contract.GetRouterAddress(&_FuxToken.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxToken *FuxTokenSession) IsUsingRouter() (bool, error) {
	return _FuxToken.Contract.IsUsingRouter(&_FuxToken.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsUsingRouter() (bool, error) {
	return _FuxToken.Contract.IsUsingRouter(&_FuxToken.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxToken *FuxTokenCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxToken *FuxTokenSession) SysRouter() (common.Address, error) {
	return _FuxToken.Contract.SysRouter(&_FuxToken.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) SysRouter() (common.Address, error) {
	return _FuxToken.Contract.SysRouter(&_FuxToken.CallOpts)
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

// Mint is a paid mutator transaction binding the contract method 0xcf21977c.
//
// Solidity: function mint(_to address, _tokenId uint256, _coin uint256, _expire uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _coin *big.Int, _expire *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "mint", _to, _tokenId, _coin, _expire, _state, _otherInfo)
}

// Mint is a paid mutator transaction binding the contract method 0xcf21977c.
//
// Solidity: function mint(_to address, _tokenId uint256, _coin uint256, _expire uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenSession) Mint(_to common.Address, _tokenId *big.Int, _coin *big.Int, _expire *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.Mint(&_FuxToken.TransactOpts, _to, _tokenId, _coin, _expire, _state, _otherInfo)
}

// Mint is a paid mutator transaction binding the contract method 0xcf21977c.
//
// Solidity: function mint(_to address, _tokenId uint256, _coin uint256, _expire uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactorSession) Mint(_to common.Address, _tokenId *big.Int, _coin *big.Int, _expire *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.Mint(&_FuxToken.TransactOpts, _to, _tokenId, _coin, _expire, _state, _otherInfo)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "transfer", _from, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) Transfer(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Transfer(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Transfer(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Transfer(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}
