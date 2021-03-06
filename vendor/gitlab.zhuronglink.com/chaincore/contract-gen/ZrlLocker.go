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

// ZrlLockerABI is the input ABI used to generate the binding from.
const ZrlLockerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"zrlStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_FUNC_LOCK_MAX\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_EXPIRE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_STATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_EXISTED\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDX_FUNC_LOCK_MIN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\"},{\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"isMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deltaT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_deltaT\",\"type\":\"uint256\"}],\"name\":\"DeltaTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_lock\",\"type\":\"bool\"}],\"name\":\"LockStateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"StateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"ExpireSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_funcId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"FunctionLockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"migrationId\",\"type\":\"string\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rbac\",\"type\":\"address\"},{\"name\":\"_zrlStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_lock\",\"type\":\"bool\"}],\"name\":\"setLockState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"setState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"}],\"name\":\"setExpire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getExpire\",\"outputs\":[{\"name\":\"expire\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_funcId\",\"type\":\"uint256\"},{\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"setFunctionLock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_funcId\",\"type\":\"uint256\"}],\"name\":\"getFunctionLock\",\"outputs\":[{\"name\":\"lock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isLock\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_deltaT\",\"type\":\"uint256\"}],\"name\":\"setDeltaT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZrlLocker is an auto generated Go binding around an Ethereum contract.
type ZrlLocker struct {
	ZrlLockerCaller     // Read-only binding to the contract
	ZrlLockerTransactor // Write-only binding to the contract
	ZrlLockerFilterer   // Log filterer for contract events
}

// ZrlLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrlLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrlLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZrlLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrlLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrlLockerSession struct {
	Contract     *ZrlLocker        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrlLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrlLockerCallerSession struct {
	Contract *ZrlLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZrlLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrlLockerTransactorSession struct {
	Contract     *ZrlLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZrlLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrlLockerRaw struct {
	Contract *ZrlLocker // Generic contract binding to access the raw methods on
}

// ZrlLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrlLockerCallerRaw struct {
	Contract *ZrlLockerCaller // Generic read-only contract binding to access the raw methods on
}

// ZrlLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrlLockerTransactorRaw struct {
	Contract *ZrlLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrlLocker creates a new instance of ZrlLocker, bound to a specific deployed contract.
func NewZrlLocker(address common.Address, backend bind.ContractBackend) (*ZrlLocker, error) {
	contract, err := bindZrlLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZrlLocker{ZrlLockerCaller: ZrlLockerCaller{contract: contract}, ZrlLockerTransactor: ZrlLockerTransactor{contract: contract}, ZrlLockerFilterer: ZrlLockerFilterer{contract: contract}}, nil
}

// NewZrlLockerCaller creates a new read-only instance of ZrlLocker, bound to a specific deployed contract.
func NewZrlLockerCaller(address common.Address, caller bind.ContractCaller) (*ZrlLockerCaller, error) {
	contract, err := bindZrlLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlLockerCaller{contract: contract}, nil
}

// NewZrlLockerTransactor creates a new write-only instance of ZrlLocker, bound to a specific deployed contract.
func NewZrlLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrlLockerTransactor, error) {
	contract, err := bindZrlLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZrlLockerTransactor{contract: contract}, nil
}

// NewZrlLockerFilterer creates a new log filterer instance of ZrlLocker, bound to a specific deployed contract.
func NewZrlLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*ZrlLockerFilterer, error) {
	contract, err := bindZrlLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZrlLockerFilterer{contract: contract}, nil
}

// bindZrlLocker binds a generic wrapper to an already deployed contract.
func bindZrlLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrlLockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlLocker *ZrlLockerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlLocker.Contract.ZrlLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlLocker *ZrlLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlLocker.Contract.ZrlLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlLocker *ZrlLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlLocker.Contract.ZrlLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZrlLocker *ZrlLockerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZrlLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZrlLocker *ZrlLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZrlLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZrlLocker *ZrlLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZrlLocker.Contract.contract.Transact(opts, method, params...)
}

// IDXEXISTED is a free data retrieval call binding the contract method 0x82185f8c.
//
// Solidity: function IDX_EXISTED() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCaller) IDXEXISTED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "IDX_EXISTED")
	return *ret0, err
}

// IDXEXISTED is a free data retrieval call binding the contract method 0x82185f8c.
//
// Solidity: function IDX_EXISTED() constant returns(uint256)
func (_ZrlLocker *ZrlLockerSession) IDXEXISTED() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXEXISTED(&_ZrlLocker.CallOpts)
}

// IDXEXISTED is a free data retrieval call binding the contract method 0x82185f8c.
//
// Solidity: function IDX_EXISTED() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCallerSession) IDXEXISTED() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXEXISTED(&_ZrlLocker.CallOpts)
}

// IDXEXPIRE is a free data retrieval call binding the contract method 0x66a41a70.
//
// Solidity: function IDX_EXPIRE() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCaller) IDXEXPIRE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "IDX_EXPIRE")
	return *ret0, err
}

// IDXEXPIRE is a free data retrieval call binding the contract method 0x66a41a70.
//
// Solidity: function IDX_EXPIRE() constant returns(uint256)
func (_ZrlLocker *ZrlLockerSession) IDXEXPIRE() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXEXPIRE(&_ZrlLocker.CallOpts)
}

// IDXEXPIRE is a free data retrieval call binding the contract method 0x66a41a70.
//
// Solidity: function IDX_EXPIRE() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCallerSession) IDXEXPIRE() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXEXPIRE(&_ZrlLocker.CallOpts)
}

// IDXFUNCLOCKMAX is a free data retrieval call binding the contract method 0x2cb902cf.
//
// Solidity: function IDX_FUNC_LOCK_MAX() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCaller) IDXFUNCLOCKMAX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "IDX_FUNC_LOCK_MAX")
	return *ret0, err
}

// IDXFUNCLOCKMAX is a free data retrieval call binding the contract method 0x2cb902cf.
//
// Solidity: function IDX_FUNC_LOCK_MAX() constant returns(uint256)
func (_ZrlLocker *ZrlLockerSession) IDXFUNCLOCKMAX() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXFUNCLOCKMAX(&_ZrlLocker.CallOpts)
}

// IDXFUNCLOCKMAX is a free data retrieval call binding the contract method 0x2cb902cf.
//
// Solidity: function IDX_FUNC_LOCK_MAX() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCallerSession) IDXFUNCLOCKMAX() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXFUNCLOCKMAX(&_ZrlLocker.CallOpts)
}

// IDXFUNCLOCKMIN is a free data retrieval call binding the contract method 0xa3e09cfb.
//
// Solidity: function IDX_FUNC_LOCK_MIN() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCaller) IDXFUNCLOCKMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "IDX_FUNC_LOCK_MIN")
	return *ret0, err
}

// IDXFUNCLOCKMIN is a free data retrieval call binding the contract method 0xa3e09cfb.
//
// Solidity: function IDX_FUNC_LOCK_MIN() constant returns(uint256)
func (_ZrlLocker *ZrlLockerSession) IDXFUNCLOCKMIN() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXFUNCLOCKMIN(&_ZrlLocker.CallOpts)
}

// IDXFUNCLOCKMIN is a free data retrieval call binding the contract method 0xa3e09cfb.
//
// Solidity: function IDX_FUNC_LOCK_MIN() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCallerSession) IDXFUNCLOCKMIN() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXFUNCLOCKMIN(&_ZrlLocker.CallOpts)
}

// IDXSTATE is a free data retrieval call binding the contract method 0x79aa0b41.
//
// Solidity: function IDX_STATE() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCaller) IDXSTATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "IDX_STATE")
	return *ret0, err
}

// IDXSTATE is a free data retrieval call binding the contract method 0x79aa0b41.
//
// Solidity: function IDX_STATE() constant returns(uint256)
func (_ZrlLocker *ZrlLockerSession) IDXSTATE() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXSTATE(&_ZrlLocker.CallOpts)
}

// IDXSTATE is a free data retrieval call binding the contract method 0x79aa0b41.
//
// Solidity: function IDX_STATE() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCallerSession) IDXSTATE() (*big.Int, error) {
	return _ZrlLocker.Contract.IDXSTATE(&_ZrlLocker.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlLocker *ZrlLockerCaller) ROLEADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlLocker *ZrlLockerSession) ROLEADMIN() (string, error) {
	return _ZrlLocker.Contract.ROLEADMIN(&_ZrlLocker.CallOpts)
}

// ROLEADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_ZrlLocker *ZrlLockerCallerSession) ROLEADMIN() (string, error) {
	return _ZrlLocker.Contract.ROLEADMIN(&_ZrlLocker.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlLocker *ZrlLockerCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlLocker *ZrlLockerSession) ROLECLUSTER() (string, error) {
	return _ZrlLocker.Contract.ROLECLUSTER(&_ZrlLocker.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_ZrlLocker *ZrlLockerCallerSession) ROLECLUSTER() (string, error) {
	return _ZrlLocker.Contract.ROLECLUSTER(&_ZrlLocker.CallOpts)
}

// DeltaT is a free data retrieval call binding the contract method 0xc7a0f8f7.
//
// Solidity: function deltaT() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCaller) DeltaT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "deltaT")
	return *ret0, err
}

// DeltaT is a free data retrieval call binding the contract method 0xc7a0f8f7.
//
// Solidity: function deltaT() constant returns(uint256)
func (_ZrlLocker *ZrlLockerSession) DeltaT() (*big.Int, error) {
	return _ZrlLocker.Contract.DeltaT(&_ZrlLocker.CallOpts)
}

// DeltaT is a free data retrieval call binding the contract method 0xc7a0f8f7.
//
// Solidity: function deltaT() constant returns(uint256)
func (_ZrlLocker *ZrlLockerCallerSession) DeltaT() (*big.Int, error) {
	return _ZrlLocker.Contract.DeltaT(&_ZrlLocker.CallOpts)
}

// GetExpire is a free data retrieval call binding the contract method 0x95c78f82.
//
// Solidity: function getExpire(_tokenId uint256) constant returns(expire uint256)
func (_ZrlLocker *ZrlLockerCaller) GetExpire(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "getExpire", _tokenId)
	return *ret0, err
}

// GetExpire is a free data retrieval call binding the contract method 0x95c78f82.
//
// Solidity: function getExpire(_tokenId uint256) constant returns(expire uint256)
func (_ZrlLocker *ZrlLockerSession) GetExpire(_tokenId *big.Int) (*big.Int, error) {
	return _ZrlLocker.Contract.GetExpire(&_ZrlLocker.CallOpts, _tokenId)
}

// GetExpire is a free data retrieval call binding the contract method 0x95c78f82.
//
// Solidity: function getExpire(_tokenId uint256) constant returns(expire uint256)
func (_ZrlLocker *ZrlLockerCallerSession) GetExpire(_tokenId *big.Int) (*big.Int, error) {
	return _ZrlLocker.Contract.GetExpire(&_ZrlLocker.CallOpts, _tokenId)
}

// GetFunctionLock is a free data retrieval call binding the contract method 0x7d24993a.
//
// Solidity: function getFunctionLock(_tokenId uint256, _funcId uint256) constant returns(lock uint256)
func (_ZrlLocker *ZrlLockerCaller) GetFunctionLock(opts *bind.CallOpts, _tokenId *big.Int, _funcId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "getFunctionLock", _tokenId, _funcId)
	return *ret0, err
}

// GetFunctionLock is a free data retrieval call binding the contract method 0x7d24993a.
//
// Solidity: function getFunctionLock(_tokenId uint256, _funcId uint256) constant returns(lock uint256)
func (_ZrlLocker *ZrlLockerSession) GetFunctionLock(_tokenId *big.Int, _funcId *big.Int) (*big.Int, error) {
	return _ZrlLocker.Contract.GetFunctionLock(&_ZrlLocker.CallOpts, _tokenId, _funcId)
}

// GetFunctionLock is a free data retrieval call binding the contract method 0x7d24993a.
//
// Solidity: function getFunctionLock(_tokenId uint256, _funcId uint256) constant returns(lock uint256)
func (_ZrlLocker *ZrlLockerCallerSession) GetFunctionLock(_tokenId *big.Int, _funcId *big.Int) (*big.Int, error) {
	return _ZrlLocker.Contract.GetFunctionLock(&_ZrlLocker.CallOpts, _tokenId, _funcId)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(_tokenId uint256) constant returns(state uint256)
func (_ZrlLocker *ZrlLockerCaller) GetState(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "getState", _tokenId)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(_tokenId uint256) constant returns(state uint256)
func (_ZrlLocker *ZrlLockerSession) GetState(_tokenId *big.Int) (*big.Int, error) {
	return _ZrlLocker.Contract.GetState(&_ZrlLocker.CallOpts, _tokenId)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(_tokenId uint256) constant returns(state uint256)
func (_ZrlLocker *ZrlLockerCallerSession) GetState(_tokenId *big.Int) (*big.Int, error) {
	return _ZrlLocker.Contract.GetState(&_ZrlLocker.CallOpts, _tokenId)
}

// IsLock is a free data retrieval call binding the contract method 0x603cc0b8.
//
// Solidity: function isLock(_tokenId uint256) constant returns(ret bool)
func (_ZrlLocker *ZrlLockerCaller) IsLock(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "isLock", _tokenId)
	return *ret0, err
}

// IsLock is a free data retrieval call binding the contract method 0x603cc0b8.
//
// Solidity: function isLock(_tokenId uint256) constant returns(ret bool)
func (_ZrlLocker *ZrlLockerSession) IsLock(_tokenId *big.Int) (bool, error) {
	return _ZrlLocker.Contract.IsLock(&_ZrlLocker.CallOpts, _tokenId)
}

// IsLock is a free data retrieval call binding the contract method 0x603cc0b8.
//
// Solidity: function isLock(_tokenId uint256) constant returns(ret bool)
func (_ZrlLocker *ZrlLockerCallerSession) IsLock(_tokenId *big.Int) (bool, error) {
	return _ZrlLocker.Contract.IsLock(&_ZrlLocker.CallOpts, _tokenId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlLocker *ZrlLockerCaller) IsMigrated(opts *bind.CallOpts, contractName string, migrationId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "isMigrated", contractName, migrationId)
	return *ret0, err
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlLocker *ZrlLockerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlLocker.Contract.IsMigrated(&_ZrlLocker.CallOpts, contractName, migrationId)
}

// IsMigrated is a free data retrieval call binding the contract method 0xc0bac1a8.
//
// Solidity: function isMigrated(contractName string, migrationId string) constant returns(bool)
func (_ZrlLocker *ZrlLockerCallerSession) IsMigrated(contractName string, migrationId string) (bool, error) {
	return _ZrlLocker.Contract.IsMigrated(&_ZrlLocker.CallOpts, contractName, migrationId)
}

// LockList is a free data retrieval call binding the contract method 0xa41d222b.
//
// Solidity: function lockList( uint256) constant returns(bool)
func (_ZrlLocker *ZrlLockerCaller) LockList(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "lockList", arg0)
	return *ret0, err
}

// LockList is a free data retrieval call binding the contract method 0xa41d222b.
//
// Solidity: function lockList( uint256) constant returns(bool)
func (_ZrlLocker *ZrlLockerSession) LockList(arg0 *big.Int) (bool, error) {
	return _ZrlLocker.Contract.LockList(&_ZrlLocker.CallOpts, arg0)
}

// LockList is a free data retrieval call binding the contract method 0xa41d222b.
//
// Solidity: function lockList( uint256) constant returns(bool)
func (_ZrlLocker *ZrlLockerCallerSession) LockList(arg0 *big.Int) (bool, error) {
	return _ZrlLocker.Contract.LockList(&_ZrlLocker.CallOpts, arg0)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlLocker *ZrlLockerCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlLocker *ZrlLockerSession) Rbac() (common.Address, error) {
	return _ZrlLocker.Contract.Rbac(&_ZrlLocker.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_ZrlLocker *ZrlLockerCallerSession) Rbac() (common.Address, error) {
	return _ZrlLocker.Contract.Rbac(&_ZrlLocker.CallOpts)
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlLocker *ZrlLockerCaller) ZrlStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZrlLocker.contract.Call(opts, out, "zrlStorage")
	return *ret0, err
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlLocker *ZrlLockerSession) ZrlStorage() (common.Address, error) {
	return _ZrlLocker.Contract.ZrlStorage(&_ZrlLocker.CallOpts)
}

// ZrlStorage is a free data retrieval call binding the contract method 0x16dfd88a.
//
// Solidity: function zrlStorage() constant returns(address)
func (_ZrlLocker *ZrlLockerCallerSession) ZrlStorage() (common.Address, error) {
	return _ZrlLocker.Contract.ZrlStorage(&_ZrlLocker.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_rbac address, _zrlStorage address) returns()
func (_ZrlLocker *ZrlLockerTransactor) Initialize(opts *bind.TransactOpts, _rbac common.Address, _zrlStorage common.Address) (*types.Transaction, error) {
	return _ZrlLocker.contract.Transact(opts, "initialize", _rbac, _zrlStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_rbac address, _zrlStorage address) returns()
func (_ZrlLocker *ZrlLockerSession) Initialize(_rbac common.Address, _zrlStorage common.Address) (*types.Transaction, error) {
	return _ZrlLocker.Contract.Initialize(&_ZrlLocker.TransactOpts, _rbac, _zrlStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_rbac address, _zrlStorage address) returns()
func (_ZrlLocker *ZrlLockerTransactorSession) Initialize(_rbac common.Address, _zrlStorage common.Address) (*types.Transaction, error) {
	return _ZrlLocker.Contract.Initialize(&_ZrlLocker.TransactOpts, _rbac, _zrlStorage)
}

// SetDeltaT is a paid mutator transaction binding the contract method 0x52fc1d58.
//
// Solidity: function setDeltaT(_deltaT uint256) returns()
func (_ZrlLocker *ZrlLockerTransactor) SetDeltaT(opts *bind.TransactOpts, _deltaT *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.contract.Transact(opts, "setDeltaT", _deltaT)
}

// SetDeltaT is a paid mutator transaction binding the contract method 0x52fc1d58.
//
// Solidity: function setDeltaT(_deltaT uint256) returns()
func (_ZrlLocker *ZrlLockerSession) SetDeltaT(_deltaT *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetDeltaT(&_ZrlLocker.TransactOpts, _deltaT)
}

// SetDeltaT is a paid mutator transaction binding the contract method 0x52fc1d58.
//
// Solidity: function setDeltaT(_deltaT uint256) returns()
func (_ZrlLocker *ZrlLockerTransactorSession) SetDeltaT(_deltaT *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetDeltaT(&_ZrlLocker.TransactOpts, _deltaT)
}

// SetExpire is a paid mutator transaction binding the contract method 0xa8c5cf27.
//
// Solidity: function setExpire(_tokenId uint256, _expire uint256) returns()
func (_ZrlLocker *ZrlLockerTransactor) SetExpire(opts *bind.TransactOpts, _tokenId *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.contract.Transact(opts, "setExpire", _tokenId, _expire)
}

// SetExpire is a paid mutator transaction binding the contract method 0xa8c5cf27.
//
// Solidity: function setExpire(_tokenId uint256, _expire uint256) returns()
func (_ZrlLocker *ZrlLockerSession) SetExpire(_tokenId *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetExpire(&_ZrlLocker.TransactOpts, _tokenId, _expire)
}

// SetExpire is a paid mutator transaction binding the contract method 0xa8c5cf27.
//
// Solidity: function setExpire(_tokenId uint256, _expire uint256) returns()
func (_ZrlLocker *ZrlLockerTransactorSession) SetExpire(_tokenId *big.Int, _expire *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetExpire(&_ZrlLocker.TransactOpts, _tokenId, _expire)
}

// SetFunctionLock is a paid mutator transaction binding the contract method 0x60a0bd7c.
//
// Solidity: function setFunctionLock(_tokenId uint256, _funcId uint256, _lock uint256) returns()
func (_ZrlLocker *ZrlLockerTransactor) SetFunctionLock(opts *bind.TransactOpts, _tokenId *big.Int, _funcId *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.contract.Transact(opts, "setFunctionLock", _tokenId, _funcId, _lock)
}

// SetFunctionLock is a paid mutator transaction binding the contract method 0x60a0bd7c.
//
// Solidity: function setFunctionLock(_tokenId uint256, _funcId uint256, _lock uint256) returns()
func (_ZrlLocker *ZrlLockerSession) SetFunctionLock(_tokenId *big.Int, _funcId *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetFunctionLock(&_ZrlLocker.TransactOpts, _tokenId, _funcId, _lock)
}

// SetFunctionLock is a paid mutator transaction binding the contract method 0x60a0bd7c.
//
// Solidity: function setFunctionLock(_tokenId uint256, _funcId uint256, _lock uint256) returns()
func (_ZrlLocker *ZrlLockerTransactorSession) SetFunctionLock(_tokenId *big.Int, _funcId *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetFunctionLock(&_ZrlLocker.TransactOpts, _tokenId, _funcId, _lock)
}

// SetLockState is a paid mutator transaction binding the contract method 0x2a291e11.
//
// Solidity: function setLockState(_state uint256, _lock bool) returns()
func (_ZrlLocker *ZrlLockerTransactor) SetLockState(opts *bind.TransactOpts, _state *big.Int, _lock bool) (*types.Transaction, error) {
	return _ZrlLocker.contract.Transact(opts, "setLockState", _state, _lock)
}

// SetLockState is a paid mutator transaction binding the contract method 0x2a291e11.
//
// Solidity: function setLockState(_state uint256, _lock bool) returns()
func (_ZrlLocker *ZrlLockerSession) SetLockState(_state *big.Int, _lock bool) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetLockState(&_ZrlLocker.TransactOpts, _state, _lock)
}

// SetLockState is a paid mutator transaction binding the contract method 0x2a291e11.
//
// Solidity: function setLockState(_state uint256, _lock bool) returns()
func (_ZrlLocker *ZrlLockerTransactorSession) SetLockState(_state *big.Int, _lock bool) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetLockState(&_ZrlLocker.TransactOpts, _state, _lock)
}

// SetState is a paid mutator transaction binding the contract method 0xb9d77bfc.
//
// Solidity: function setState(_tokenId uint256, _state uint256) returns()
func (_ZrlLocker *ZrlLockerTransactor) SetState(opts *bind.TransactOpts, _tokenId *big.Int, _state *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.contract.Transact(opts, "setState", _tokenId, _state)
}

// SetState is a paid mutator transaction binding the contract method 0xb9d77bfc.
//
// Solidity: function setState(_tokenId uint256, _state uint256) returns()
func (_ZrlLocker *ZrlLockerSession) SetState(_tokenId *big.Int, _state *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetState(&_ZrlLocker.TransactOpts, _tokenId, _state)
}

// SetState is a paid mutator transaction binding the contract method 0xb9d77bfc.
//
// Solidity: function setState(_tokenId uint256, _state uint256) returns()
func (_ZrlLocker *ZrlLockerTransactorSession) SetState(_tokenId *big.Int, _state *big.Int) (*types.Transaction, error) {
	return _ZrlLocker.Contract.SetState(&_ZrlLocker.TransactOpts, _tokenId, _state)
}

// ZrlLockerDeltaTimeSetIterator is returned from FilterDeltaTimeSet and is used to iterate over the raw logs and unpacked data for DeltaTimeSet events raised by the ZrlLocker contract.
type ZrlLockerDeltaTimeSetIterator struct {
	Event *ZrlLockerDeltaTimeSet // Event containing the contract specifics and raw log

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
func (it *ZrlLockerDeltaTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlLockerDeltaTimeSet)
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
		it.Event = new(ZrlLockerDeltaTimeSet)
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
func (it *ZrlLockerDeltaTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlLockerDeltaTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlLockerDeltaTimeSet represents a DeltaTimeSet event raised by the ZrlLocker contract.
type ZrlLockerDeltaTimeSet struct {
	Caller common.Address
	DeltaT *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeltaTimeSet is a free log retrieval operation binding the contract event 0xacd17757a3133f061ef65177bb74cd35cf1e005becb93c65f73f9c95d4a98794.
//
// Solidity: event DeltaTimeSet(_caller address, _deltaT uint256)
func (_ZrlLocker *ZrlLockerFilterer) FilterDeltaTimeSet(opts *bind.FilterOpts) (*ZrlLockerDeltaTimeSetIterator, error) {

	logs, sub, err := _ZrlLocker.contract.FilterLogs(opts, "DeltaTimeSet")
	if err != nil {
		return nil, err
	}
	return &ZrlLockerDeltaTimeSetIterator{contract: _ZrlLocker.contract, event: "DeltaTimeSet", logs: logs, sub: sub}, nil
}

// WatchDeltaTimeSet is a free log subscription operation binding the contract event 0xacd17757a3133f061ef65177bb74cd35cf1e005becb93c65f73f9c95d4a98794.
//
// Solidity: event DeltaTimeSet(_caller address, _deltaT uint256)
func (_ZrlLocker *ZrlLockerFilterer) WatchDeltaTimeSet(opts *bind.WatchOpts, sink chan<- *ZrlLockerDeltaTimeSet) (event.Subscription, error) {

	logs, sub, err := _ZrlLocker.contract.WatchLogs(opts, "DeltaTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlLockerDeltaTimeSet)
				if err := _ZrlLocker.contract.UnpackLog(event, "DeltaTimeSet", log); err != nil {
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

// ZrlLockerExpireSetIterator is returned from FilterExpireSet and is used to iterate over the raw logs and unpacked data for ExpireSet events raised by the ZrlLocker contract.
type ZrlLockerExpireSetIterator struct {
	Event *ZrlLockerExpireSet // Event containing the contract specifics and raw log

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
func (it *ZrlLockerExpireSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlLockerExpireSet)
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
		it.Event = new(ZrlLockerExpireSet)
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
func (it *ZrlLockerExpireSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlLockerExpireSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlLockerExpireSet represents a ExpireSet event raised by the ZrlLocker contract.
type ZrlLockerExpireSet struct {
	Caller  common.Address
	TokenId *big.Int
	Expire  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExpireSet is a free log retrieval operation binding the contract event 0x78cd799f5d099bdd3586a6ab5a947eb80e932459410dc30146b47705e952de57.
//
// Solidity: event ExpireSet(_caller address, _tokenId uint256, _expire uint256)
func (_ZrlLocker *ZrlLockerFilterer) FilterExpireSet(opts *bind.FilterOpts) (*ZrlLockerExpireSetIterator, error) {

	logs, sub, err := _ZrlLocker.contract.FilterLogs(opts, "ExpireSet")
	if err != nil {
		return nil, err
	}
	return &ZrlLockerExpireSetIterator{contract: _ZrlLocker.contract, event: "ExpireSet", logs: logs, sub: sub}, nil
}

// WatchExpireSet is a free log subscription operation binding the contract event 0x78cd799f5d099bdd3586a6ab5a947eb80e932459410dc30146b47705e952de57.
//
// Solidity: event ExpireSet(_caller address, _tokenId uint256, _expire uint256)
func (_ZrlLocker *ZrlLockerFilterer) WatchExpireSet(opts *bind.WatchOpts, sink chan<- *ZrlLockerExpireSet) (event.Subscription, error) {

	logs, sub, err := _ZrlLocker.contract.WatchLogs(opts, "ExpireSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlLockerExpireSet)
				if err := _ZrlLocker.contract.UnpackLog(event, "ExpireSet", log); err != nil {
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

// ZrlLockerFunctionLockSetIterator is returned from FilterFunctionLockSet and is used to iterate over the raw logs and unpacked data for FunctionLockSet events raised by the ZrlLocker contract.
type ZrlLockerFunctionLockSetIterator struct {
	Event *ZrlLockerFunctionLockSet // Event containing the contract specifics and raw log

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
func (it *ZrlLockerFunctionLockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlLockerFunctionLockSet)
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
		it.Event = new(ZrlLockerFunctionLockSet)
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
func (it *ZrlLockerFunctionLockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlLockerFunctionLockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlLockerFunctionLockSet represents a FunctionLockSet event raised by the ZrlLocker contract.
type ZrlLockerFunctionLockSet struct {
	Caller  common.Address
	TokenId *big.Int
	FuncId  *big.Int
	Lock    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFunctionLockSet is a free log retrieval operation binding the contract event 0xc93e147af259a8b65d8f09a4df8a9ce187f29e04e41bf69310536743578196ad.
//
// Solidity: event FunctionLockSet(_caller address, _tokenId uint256, _funcId uint256, _lock uint256)
func (_ZrlLocker *ZrlLockerFilterer) FilterFunctionLockSet(opts *bind.FilterOpts) (*ZrlLockerFunctionLockSetIterator, error) {

	logs, sub, err := _ZrlLocker.contract.FilterLogs(opts, "FunctionLockSet")
	if err != nil {
		return nil, err
	}
	return &ZrlLockerFunctionLockSetIterator{contract: _ZrlLocker.contract, event: "FunctionLockSet", logs: logs, sub: sub}, nil
}

// WatchFunctionLockSet is a free log subscription operation binding the contract event 0xc93e147af259a8b65d8f09a4df8a9ce187f29e04e41bf69310536743578196ad.
//
// Solidity: event FunctionLockSet(_caller address, _tokenId uint256, _funcId uint256, _lock uint256)
func (_ZrlLocker *ZrlLockerFilterer) WatchFunctionLockSet(opts *bind.WatchOpts, sink chan<- *ZrlLockerFunctionLockSet) (event.Subscription, error) {

	logs, sub, err := _ZrlLocker.contract.WatchLogs(opts, "FunctionLockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlLockerFunctionLockSet)
				if err := _ZrlLocker.contract.UnpackLog(event, "FunctionLockSet", log); err != nil {
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

// ZrlLockerLockStateSetIterator is returned from FilterLockStateSet and is used to iterate over the raw logs and unpacked data for LockStateSet events raised by the ZrlLocker contract.
type ZrlLockerLockStateSetIterator struct {
	Event *ZrlLockerLockStateSet // Event containing the contract specifics and raw log

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
func (it *ZrlLockerLockStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlLockerLockStateSet)
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
		it.Event = new(ZrlLockerLockStateSet)
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
func (it *ZrlLockerLockStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlLockerLockStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlLockerLockStateSet represents a LockStateSet event raised by the ZrlLocker contract.
type ZrlLockerLockStateSet struct {
	Caller common.Address
	State  *big.Int
	Lock   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockStateSet is a free log retrieval operation binding the contract event 0xb8640e2eb4b26293f0742ae8635ee5e79149eda9dbc2d72a22730d96cdcae63b.
//
// Solidity: event LockStateSet(_caller address, _state uint256, _lock bool)
func (_ZrlLocker *ZrlLockerFilterer) FilterLockStateSet(opts *bind.FilterOpts) (*ZrlLockerLockStateSetIterator, error) {

	logs, sub, err := _ZrlLocker.contract.FilterLogs(opts, "LockStateSet")
	if err != nil {
		return nil, err
	}
	return &ZrlLockerLockStateSetIterator{contract: _ZrlLocker.contract, event: "LockStateSet", logs: logs, sub: sub}, nil
}

// WatchLockStateSet is a free log subscription operation binding the contract event 0xb8640e2eb4b26293f0742ae8635ee5e79149eda9dbc2d72a22730d96cdcae63b.
//
// Solidity: event LockStateSet(_caller address, _state uint256, _lock bool)
func (_ZrlLocker *ZrlLockerFilterer) WatchLockStateSet(opts *bind.WatchOpts, sink chan<- *ZrlLockerLockStateSet) (event.Subscription, error) {

	logs, sub, err := _ZrlLocker.contract.WatchLogs(opts, "LockStateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlLockerLockStateSet)
				if err := _ZrlLocker.contract.UnpackLog(event, "LockStateSet", log); err != nil {
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

// ZrlLockerMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the ZrlLocker contract.
type ZrlLockerMigratedIterator struct {
	Event *ZrlLockerMigrated // Event containing the contract specifics and raw log

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
func (it *ZrlLockerMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlLockerMigrated)
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
		it.Event = new(ZrlLockerMigrated)
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
func (it *ZrlLockerMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlLockerMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlLockerMigrated represents a Migrated event raised by the ZrlLocker contract.
type ZrlLockerMigrated struct {
	ContractName string
	MigrationId  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ZrlLocker *ZrlLockerFilterer) FilterMigrated(opts *bind.FilterOpts) (*ZrlLockerMigratedIterator, error) {

	logs, sub, err := _ZrlLocker.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &ZrlLockerMigratedIterator{contract: _ZrlLocker.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xdd117a11c22118c9dee4b5a67ce578bc44529dce21ee0ccc439588fbb9fb4ea3.
//
// Solidity: event Migrated(contractName string, migrationId string)
func (_ZrlLocker *ZrlLockerFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *ZrlLockerMigrated) (event.Subscription, error) {

	logs, sub, err := _ZrlLocker.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlLockerMigrated)
				if err := _ZrlLocker.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ZrlLockerStateSetIterator is returned from FilterStateSet and is used to iterate over the raw logs and unpacked data for StateSet events raised by the ZrlLocker contract.
type ZrlLockerStateSetIterator struct {
	Event *ZrlLockerStateSet // Event containing the contract specifics and raw log

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
func (it *ZrlLockerStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZrlLockerStateSet)
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
		it.Event = new(ZrlLockerStateSet)
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
func (it *ZrlLockerStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZrlLockerStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZrlLockerStateSet represents a StateSet event raised by the ZrlLocker contract.
type ZrlLockerStateSet struct {
	Caller  common.Address
	TokenId *big.Int
	State   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStateSet is a free log retrieval operation binding the contract event 0x1d79a6e0d9992f96b2facb02f3f2430911c472085429997f740dd1c9fbcbc445.
//
// Solidity: event StateSet(_caller address, _tokenId uint256, _state uint256)
func (_ZrlLocker *ZrlLockerFilterer) FilterStateSet(opts *bind.FilterOpts) (*ZrlLockerStateSetIterator, error) {

	logs, sub, err := _ZrlLocker.contract.FilterLogs(opts, "StateSet")
	if err != nil {
		return nil, err
	}
	return &ZrlLockerStateSetIterator{contract: _ZrlLocker.contract, event: "StateSet", logs: logs, sub: sub}, nil
}

// WatchStateSet is a free log subscription operation binding the contract event 0x1d79a6e0d9992f96b2facb02f3f2430911c472085429997f740dd1c9fbcbc445.
//
// Solidity: event StateSet(_caller address, _tokenId uint256, _state uint256)
func (_ZrlLocker *ZrlLockerFilterer) WatchStateSet(opts *bind.WatchOpts, sink chan<- *ZrlLockerStateSet) (event.Subscription, error) {

	logs, sub, err := _ZrlLocker.contract.WatchLogs(opts, "StateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZrlLockerStateSet)
				if err := _ZrlLocker.contract.UnpackLog(event, "StateSet", log); err != nil {
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
