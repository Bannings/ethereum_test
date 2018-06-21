// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dr_contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ARStoreABI is the input ABI used to generate the binding from.
const ARStoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInDRGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_router\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"bcAccount\",\"type\":\"string\"}],\"name\":\"RegisterUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherARPayInfo\",\"type\":\"string\"}],\"name\":\"ARPay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherContractInfo\",\"type\":\"string\"}],\"name\":\"AddContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"discountId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherDiscountInfo\",\"type\":\"string\"}],\"name\":\"ARDiscountTxUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"arId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherDiscountId\",\"type\":\"string\"}],\"name\":\"ARInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"companyId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherCompanyInfo\",\"type\":\"string\"}],\"name\":\"CompanyInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherContractInfo\",\"type\":\"string\"}],\"name\":\"ContractInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherARPayInfo\",\"type\":\"string\"}],\"name\":\"ARPaymentTxUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"companyId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherCompanyInfo\",\"type\":\"string\"}],\"name\":\"AddCompany\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"discountId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherDiscountInfo\",\"type\":\"string\"}],\"name\":\"ARDiscount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"arId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherARInfo\",\"type\":\"string\"}],\"name\":\"CreateAR\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_arId\",\"type\":\"string\"}],\"name\":\"queryARInfo\",\"outputs\":[{\"name\":\"otherInfo\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_arId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"updateARInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_arId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"createAR\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_companyId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"updateCompanyInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_companyId\",\"type\":\"string\"}],\"name\":\"queryCompanyInfo\",\"outputs\":[{\"name\":\"otherInfo\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_companyId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"addCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractNo\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"updateContractInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contractNo\",\"type\":\"string\"}],\"name\":\"queryContractInfo\",\"outputs\":[{\"name\":\"otherInfo\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractNo\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"addContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"updatePaymentInfoOfAR\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_discountId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"updateDiscountInfoOfAR\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"payByAR\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_payId\",\"type\":\"string\"}],\"name\":\"queryPaymentTx\",\"outputs\":[{\"name\":\"otherInfo\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_discountId\",\"type\":\"string\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"discountByAR\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_discountId\",\"type\":\"string\"}],\"name\":\"queryDiscountTx\",\"outputs\":[{\"name\":\"otherInfo\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"hasAR\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"hasCompany\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"hasContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"hasPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"hasDiscount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ARStore is an auto generated Go binding around an Ethereum contract.
type ARStore struct {
	ARStoreCaller     // Read-only binding to the contract
	ARStoreTransactor // Write-only binding to the contract
	ARStoreFilterer   // Log filterer for contract events
}

// ARStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ARStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARStoreSession struct {
	Contract     *ARStore          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARStoreCallerSession struct {
	Contract *ARStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ARStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARStoreTransactorSession struct {
	Contract     *ARStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ARStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ARStoreRaw struct {
	Contract *ARStore // Generic contract binding to access the raw methods on
}

// ARStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARStoreCallerRaw struct {
	Contract *ARStoreCaller // Generic read-only contract binding to access the raw methods on
}

// ARStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARStoreTransactorRaw struct {
	Contract *ARStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewARStore creates a new instance of ARStore, bound to a specific deployed contract.
func NewARStore(address common.Address, backend bind.ContractBackend) (*ARStore, error) {
	contract, err := bindARStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARStore{ARStoreCaller: ARStoreCaller{contract: contract}, ARStoreTransactor: ARStoreTransactor{contract: contract}, ARStoreFilterer: ARStoreFilterer{contract: contract}}, nil
}

// NewARStoreCaller creates a new read-only instance of ARStore, bound to a specific deployed contract.
func NewARStoreCaller(address common.Address, caller bind.ContractCaller) (*ARStoreCaller, error) {
	contract, err := bindARStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARStoreCaller{contract: contract}, nil
}

// NewARStoreTransactor creates a new write-only instance of ARStore, bound to a specific deployed contract.
func NewARStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ARStoreTransactor, error) {
	contract, err := bindARStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARStoreTransactor{contract: contract}, nil
}

// NewARStoreFilterer creates a new log filterer instance of ARStore, bound to a specific deployed contract.
func NewARStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ARStoreFilterer, error) {
	contract, err := bindARStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARStoreFilterer{contract: contract}, nil
}

// bindARStore binds a generic wrapper to an already deployed contract.
func bindARStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARStore *ARStoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ARStore.Contract.ARStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARStore *ARStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARStore.Contract.ARStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARStore *ARStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARStore.Contract.ARStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARStore *ARStoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ARStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARStore *ARStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARStore *ARStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARStore.Contract.contract.Transact(opts, method, params...)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_ARStore *ARStoreCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_ARStore *ARStoreSession) ROUTECFRBAC() (string, error) {
	return _ARStore.Contract.ROUTECFRBAC(&_ARStore.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_ARStore *ARStoreCallerSession) ROUTECFRBAC() (string, error) {
	return _ARStore.Contract.ROUTECFRBAC(&_ARStore.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_ARStore *ARStoreCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_ARStore *ARStoreSession) ROUTECONFIG() (string, error) {
	return _ARStore.Contract.ROUTECONFIG(&_ARStore.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_ARStore *ARStoreCallerSession) ROUTECONFIG() (string, error) {
	return _ARStore.Contract.ROUTECONFIG(&_ARStore.CallOpts)
}

// IsInDRGroup is a free data retrieval call binding the contract method 0xacd11b60.
//
// Solidity: function _isInDRGroup(_addr address) constant returns(bool)
func (_ARStore *ARStoreCaller) IsInDRGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "_isInDRGroup", _addr)
	return *ret0, err
}

// IsInDRGroup is a free data retrieval call binding the contract method 0xacd11b60.
//
// Solidity: function _isInDRGroup(_addr address) constant returns(bool)
func (_ARStore *ARStoreSession) IsInDRGroup(_addr common.Address) (bool, error) {
	return _ARStore.Contract.IsInDRGroup(&_ARStore.CallOpts, _addr)
}

// IsInDRGroup is a free data retrieval call binding the contract method 0xacd11b60.
//
// Solidity: function _isInDRGroup(_addr address) constant returns(bool)
func (_ARStore *ARStoreCallerSession) IsInDRGroup(_addr common.Address) (bool, error) {
	return _ARStore.Contract.IsInDRGroup(&_ARStore.CallOpts, _addr)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_ARStore *ARStoreCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_ARStore *ARStoreSession) GetRouterAddress() (common.Address, error) {
	return _ARStore.Contract.GetRouterAddress(&_ARStore.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_ARStore *ARStoreCallerSession) GetRouterAddress() (common.Address, error) {
	return _ARStore.Contract.GetRouterAddress(&_ARStore.CallOpts)
}

// HasAR is a free data retrieval call binding the contract method 0x8c1fdea5.
//
// Solidity: function hasAR(_id string) constant returns(bool)
func (_ARStore *ARStoreCaller) HasAR(opts *bind.CallOpts, _id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "hasAR", _id)
	return *ret0, err
}

// HasAR is a free data retrieval call binding the contract method 0x8c1fdea5.
//
// Solidity: function hasAR(_id string) constant returns(bool)
func (_ARStore *ARStoreSession) HasAR(_id string) (bool, error) {
	return _ARStore.Contract.HasAR(&_ARStore.CallOpts, _id)
}

// HasAR is a free data retrieval call binding the contract method 0x8c1fdea5.
//
// Solidity: function hasAR(_id string) constant returns(bool)
func (_ARStore *ARStoreCallerSession) HasAR(_id string) (bool, error) {
	return _ARStore.Contract.HasAR(&_ARStore.CallOpts, _id)
}

// HasCompany is a free data retrieval call binding the contract method 0x24f55e29.
//
// Solidity: function hasCompany(_id string) constant returns(bool)
func (_ARStore *ARStoreCaller) HasCompany(opts *bind.CallOpts, _id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "hasCompany", _id)
	return *ret0, err
}

// HasCompany is a free data retrieval call binding the contract method 0x24f55e29.
//
// Solidity: function hasCompany(_id string) constant returns(bool)
func (_ARStore *ARStoreSession) HasCompany(_id string) (bool, error) {
	return _ARStore.Contract.HasCompany(&_ARStore.CallOpts, _id)
}

// HasCompany is a free data retrieval call binding the contract method 0x24f55e29.
//
// Solidity: function hasCompany(_id string) constant returns(bool)
func (_ARStore *ARStoreCallerSession) HasCompany(_id string) (bool, error) {
	return _ARStore.Contract.HasCompany(&_ARStore.CallOpts, _id)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(_id string) constant returns(bool)
func (_ARStore *ARStoreCaller) HasContract(opts *bind.CallOpts, _id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "hasContract", _id)
	return *ret0, err
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(_id string) constant returns(bool)
func (_ARStore *ARStoreSession) HasContract(_id string) (bool, error) {
	return _ARStore.Contract.HasContract(&_ARStore.CallOpts, _id)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(_id string) constant returns(bool)
func (_ARStore *ARStoreCallerSession) HasContract(_id string) (bool, error) {
	return _ARStore.Contract.HasContract(&_ARStore.CallOpts, _id)
}

// HasDiscount is a free data retrieval call binding the contract method 0x1d9801c1.
//
// Solidity: function hasDiscount(_id string) constant returns(bool)
func (_ARStore *ARStoreCaller) HasDiscount(opts *bind.CallOpts, _id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "hasDiscount", _id)
	return *ret0, err
}

// HasDiscount is a free data retrieval call binding the contract method 0x1d9801c1.
//
// Solidity: function hasDiscount(_id string) constant returns(bool)
func (_ARStore *ARStoreSession) HasDiscount(_id string) (bool, error) {
	return _ARStore.Contract.HasDiscount(&_ARStore.CallOpts, _id)
}

// HasDiscount is a free data retrieval call binding the contract method 0x1d9801c1.
//
// Solidity: function hasDiscount(_id string) constant returns(bool)
func (_ARStore *ARStoreCallerSession) HasDiscount(_id string) (bool, error) {
	return _ARStore.Contract.HasDiscount(&_ARStore.CallOpts, _id)
}

// HasPayment is a free data retrieval call binding the contract method 0xd9ac2b52.
//
// Solidity: function hasPayment(_id string) constant returns(bool)
func (_ARStore *ARStoreCaller) HasPayment(opts *bind.CallOpts, _id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "hasPayment", _id)
	return *ret0, err
}

// HasPayment is a free data retrieval call binding the contract method 0xd9ac2b52.
//
// Solidity: function hasPayment(_id string) constant returns(bool)
func (_ARStore *ARStoreSession) HasPayment(_id string) (bool, error) {
	return _ARStore.Contract.HasPayment(&_ARStore.CallOpts, _id)
}

// HasPayment is a free data retrieval call binding the contract method 0xd9ac2b52.
//
// Solidity: function hasPayment(_id string) constant returns(bool)
func (_ARStore *ARStoreCallerSession) HasPayment(_id string) (bool, error) {
	return _ARStore.Contract.HasPayment(&_ARStore.CallOpts, _id)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_ARStore *ARStoreCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_ARStore *ARStoreSession) IsUsingRouter() (bool, error) {
	return _ARStore.Contract.IsUsingRouter(&_ARStore.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_ARStore *ARStoreCallerSession) IsUsingRouter() (bool, error) {
	return _ARStore.Contract.IsUsingRouter(&_ARStore.CallOpts)
}

// QueryARInfo is a free data retrieval call binding the contract method 0x4e32904e.
//
// Solidity: function queryARInfo(_arId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCaller) QueryARInfo(opts *bind.CallOpts, _arId string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "queryARInfo", _arId)
	return *ret0, err
}

// QueryARInfo is a free data retrieval call binding the contract method 0x4e32904e.
//
// Solidity: function queryARInfo(_arId string) constant returns(otherInfo string)
func (_ARStore *ARStoreSession) QueryARInfo(_arId string) (string, error) {
	return _ARStore.Contract.QueryARInfo(&_ARStore.CallOpts, _arId)
}

// QueryARInfo is a free data retrieval call binding the contract method 0x4e32904e.
//
// Solidity: function queryARInfo(_arId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCallerSession) QueryARInfo(_arId string) (string, error) {
	return _ARStore.Contract.QueryARInfo(&_ARStore.CallOpts, _arId)
}

// QueryCompanyInfo is a free data retrieval call binding the contract method 0x3dc904a1.
//
// Solidity: function queryCompanyInfo(_companyId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCaller) QueryCompanyInfo(opts *bind.CallOpts, _companyId string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "queryCompanyInfo", _companyId)
	return *ret0, err
}

// QueryCompanyInfo is a free data retrieval call binding the contract method 0x3dc904a1.
//
// Solidity: function queryCompanyInfo(_companyId string) constant returns(otherInfo string)
func (_ARStore *ARStoreSession) QueryCompanyInfo(_companyId string) (string, error) {
	return _ARStore.Contract.QueryCompanyInfo(&_ARStore.CallOpts, _companyId)
}

// QueryCompanyInfo is a free data retrieval call binding the contract method 0x3dc904a1.
//
// Solidity: function queryCompanyInfo(_companyId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCallerSession) QueryCompanyInfo(_companyId string) (string, error) {
	return _ARStore.Contract.QueryCompanyInfo(&_ARStore.CallOpts, _companyId)
}

// QueryContractInfo is a free data retrieval call binding the contract method 0xadbe6459.
//
// Solidity: function queryContractInfo(_contractNo string) constant returns(otherInfo string)
func (_ARStore *ARStoreCaller) QueryContractInfo(opts *bind.CallOpts, _contractNo string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "queryContractInfo", _contractNo)
	return *ret0, err
}

// QueryContractInfo is a free data retrieval call binding the contract method 0xadbe6459.
//
// Solidity: function queryContractInfo(_contractNo string) constant returns(otherInfo string)
func (_ARStore *ARStoreSession) QueryContractInfo(_contractNo string) (string, error) {
	return _ARStore.Contract.QueryContractInfo(&_ARStore.CallOpts, _contractNo)
}

// QueryContractInfo is a free data retrieval call binding the contract method 0xadbe6459.
//
// Solidity: function queryContractInfo(_contractNo string) constant returns(otherInfo string)
func (_ARStore *ARStoreCallerSession) QueryContractInfo(_contractNo string) (string, error) {
	return _ARStore.Contract.QueryContractInfo(&_ARStore.CallOpts, _contractNo)
}

// QueryDiscountTx is a free data retrieval call binding the contract method 0xafac4086.
//
// Solidity: function queryDiscountTx(_discountId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCaller) QueryDiscountTx(opts *bind.CallOpts, _discountId string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "queryDiscountTx", _discountId)
	return *ret0, err
}

// QueryDiscountTx is a free data retrieval call binding the contract method 0xafac4086.
//
// Solidity: function queryDiscountTx(_discountId string) constant returns(otherInfo string)
func (_ARStore *ARStoreSession) QueryDiscountTx(_discountId string) (string, error) {
	return _ARStore.Contract.QueryDiscountTx(&_ARStore.CallOpts, _discountId)
}

// QueryDiscountTx is a free data retrieval call binding the contract method 0xafac4086.
//
// Solidity: function queryDiscountTx(_discountId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCallerSession) QueryDiscountTx(_discountId string) (string, error) {
	return _ARStore.Contract.QueryDiscountTx(&_ARStore.CallOpts, _discountId)
}

// QueryPaymentTx is a free data retrieval call binding the contract method 0x74cd21c4.
//
// Solidity: function queryPaymentTx(_payId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCaller) QueryPaymentTx(opts *bind.CallOpts, _payId string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "queryPaymentTx", _payId)
	return *ret0, err
}

// QueryPaymentTx is a free data retrieval call binding the contract method 0x74cd21c4.
//
// Solidity: function queryPaymentTx(_payId string) constant returns(otherInfo string)
func (_ARStore *ARStoreSession) QueryPaymentTx(_payId string) (string, error) {
	return _ARStore.Contract.QueryPaymentTx(&_ARStore.CallOpts, _payId)
}

// QueryPaymentTx is a free data retrieval call binding the contract method 0x74cd21c4.
//
// Solidity: function queryPaymentTx(_payId string) constant returns(otherInfo string)
func (_ARStore *ARStoreCallerSession) QueryPaymentTx(_payId string) (string, error) {
	return _ARStore.Contract.QueryPaymentTx(&_ARStore.CallOpts, _payId)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_ARStore *ARStoreCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ARStore.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_ARStore *ARStoreSession) SysRouter() (common.Address, error) {
	return _ARStore.Contract.SysRouter(&_ARStore.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_ARStore *ARStoreCallerSession) SysRouter() (common.Address, error) {
	return _ARStore.Contract.SysRouter(&_ARStore.CallOpts)
}

// AddCompany is a paid mutator transaction binding the contract method 0x6559c310.
//
// Solidity: function addCompany(_companyId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) AddCompany(opts *bind.TransactOpts, _companyId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "addCompany", _companyId, _otherInfo)
}

// AddCompany is a paid mutator transaction binding the contract method 0x6559c310.
//
// Solidity: function addCompany(_companyId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) AddCompany(_companyId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.AddCompany(&_ARStore.TransactOpts, _companyId, _otherInfo)
}

// AddCompany is a paid mutator transaction binding the contract method 0x6559c310.
//
// Solidity: function addCompany(_companyId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) AddCompany(_companyId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.AddCompany(&_ARStore.TransactOpts, _companyId, _otherInfo)
}

// AddContract is a paid mutator transaction binding the contract method 0x49061757.
//
// Solidity: function addContract(_contractNo string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) AddContract(opts *bind.TransactOpts, _contractNo string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "addContract", _contractNo, _otherInfo)
}

// AddContract is a paid mutator transaction binding the contract method 0x49061757.
//
// Solidity: function addContract(_contractNo string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) AddContract(_contractNo string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.AddContract(&_ARStore.TransactOpts, _contractNo, _otherInfo)
}

// AddContract is a paid mutator transaction binding the contract method 0x49061757.
//
// Solidity: function addContract(_contractNo string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) AddContract(_contractNo string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.AddContract(&_ARStore.TransactOpts, _contractNo, _otherInfo)
}

// CreateAR is a paid mutator transaction binding the contract method 0xd627325e.
//
// Solidity: function createAR(_arId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) CreateAR(opts *bind.TransactOpts, _arId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "createAR", _arId, _otherInfo)
}

// CreateAR is a paid mutator transaction binding the contract method 0xd627325e.
//
// Solidity: function createAR(_arId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) CreateAR(_arId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.CreateAR(&_ARStore.TransactOpts, _arId, _otherInfo)
}

// CreateAR is a paid mutator transaction binding the contract method 0xd627325e.
//
// Solidity: function createAR(_arId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) CreateAR(_arId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.CreateAR(&_ARStore.TransactOpts, _arId, _otherInfo)
}

// DiscountByAR is a paid mutator transaction binding the contract method 0x86672ae0.
//
// Solidity: function discountByAR(_discountId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) DiscountByAR(opts *bind.TransactOpts, _discountId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "discountByAR", _discountId, _otherInfo)
}

// DiscountByAR is a paid mutator transaction binding the contract method 0x86672ae0.
//
// Solidity: function discountByAR(_discountId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) DiscountByAR(_discountId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.DiscountByAR(&_ARStore.TransactOpts, _discountId, _otherInfo)
}

// DiscountByAR is a paid mutator transaction binding the contract method 0x86672ae0.
//
// Solidity: function discountByAR(_discountId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) DiscountByAR(_discountId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.DiscountByAR(&_ARStore.TransactOpts, _discountId, _otherInfo)
}

// PayByAR is a paid mutator transaction binding the contract method 0x849a84de.
//
// Solidity: function payByAR(_payId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) PayByAR(opts *bind.TransactOpts, _payId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "payByAR", _payId, _otherInfo)
}

// PayByAR is a paid mutator transaction binding the contract method 0x849a84de.
//
// Solidity: function payByAR(_payId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) PayByAR(_payId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.PayByAR(&_ARStore.TransactOpts, _payId, _otherInfo)
}

// PayByAR is a paid mutator transaction binding the contract method 0x849a84de.
//
// Solidity: function payByAR(_payId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) PayByAR(_payId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.PayByAR(&_ARStore.TransactOpts, _payId, _otherInfo)
}

// UpdateARInfo is a paid mutator transaction binding the contract method 0xdc527b1a.
//
// Solidity: function updateARInfo(_arId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) UpdateARInfo(opts *bind.TransactOpts, _arId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "updateARInfo", _arId, _otherInfo)
}

// UpdateARInfo is a paid mutator transaction binding the contract method 0xdc527b1a.
//
// Solidity: function updateARInfo(_arId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) UpdateARInfo(_arId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateARInfo(&_ARStore.TransactOpts, _arId, _otherInfo)
}

// UpdateARInfo is a paid mutator transaction binding the contract method 0xdc527b1a.
//
// Solidity: function updateARInfo(_arId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) UpdateARInfo(_arId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateARInfo(&_ARStore.TransactOpts, _arId, _otherInfo)
}

// UpdateCompanyInfo is a paid mutator transaction binding the contract method 0x1a0a6c68.
//
// Solidity: function updateCompanyInfo(_companyId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) UpdateCompanyInfo(opts *bind.TransactOpts, _companyId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "updateCompanyInfo", _companyId, _otherInfo)
}

// UpdateCompanyInfo is a paid mutator transaction binding the contract method 0x1a0a6c68.
//
// Solidity: function updateCompanyInfo(_companyId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) UpdateCompanyInfo(_companyId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateCompanyInfo(&_ARStore.TransactOpts, _companyId, _otherInfo)
}

// UpdateCompanyInfo is a paid mutator transaction binding the contract method 0x1a0a6c68.
//
// Solidity: function updateCompanyInfo(_companyId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) UpdateCompanyInfo(_companyId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateCompanyInfo(&_ARStore.TransactOpts, _companyId, _otherInfo)
}

// UpdateContractInfo is a paid mutator transaction binding the contract method 0xfe5d552d.
//
// Solidity: function updateContractInfo(_contractNo string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) UpdateContractInfo(opts *bind.TransactOpts, _contractNo string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "updateContractInfo", _contractNo, _otherInfo)
}

// UpdateContractInfo is a paid mutator transaction binding the contract method 0xfe5d552d.
//
// Solidity: function updateContractInfo(_contractNo string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) UpdateContractInfo(_contractNo string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateContractInfo(&_ARStore.TransactOpts, _contractNo, _otherInfo)
}

// UpdateContractInfo is a paid mutator transaction binding the contract method 0xfe5d552d.
//
// Solidity: function updateContractInfo(_contractNo string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) UpdateContractInfo(_contractNo string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateContractInfo(&_ARStore.TransactOpts, _contractNo, _otherInfo)
}

// UpdateDiscountInfoOfAR is a paid mutator transaction binding the contract method 0xa090a0a8.
//
// Solidity: function updateDiscountInfoOfAR(_discountId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) UpdateDiscountInfoOfAR(opts *bind.TransactOpts, _discountId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "updateDiscountInfoOfAR", _discountId, _otherInfo)
}

// UpdateDiscountInfoOfAR is a paid mutator transaction binding the contract method 0xa090a0a8.
//
// Solidity: function updateDiscountInfoOfAR(_discountId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) UpdateDiscountInfoOfAR(_discountId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateDiscountInfoOfAR(&_ARStore.TransactOpts, _discountId, _otherInfo)
}

// UpdateDiscountInfoOfAR is a paid mutator transaction binding the contract method 0xa090a0a8.
//
// Solidity: function updateDiscountInfoOfAR(_discountId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) UpdateDiscountInfoOfAR(_discountId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdateDiscountInfoOfAR(&_ARStore.TransactOpts, _discountId, _otherInfo)
}

// UpdatePaymentInfoOfAR is a paid mutator transaction binding the contract method 0xaa86c09e.
//
// Solidity: function updatePaymentInfoOfAR(_payId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactor) UpdatePaymentInfoOfAR(opts *bind.TransactOpts, _payId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.contract.Transact(opts, "updatePaymentInfoOfAR", _payId, _otherInfo)
}

// UpdatePaymentInfoOfAR is a paid mutator transaction binding the contract method 0xaa86c09e.
//
// Solidity: function updatePaymentInfoOfAR(_payId string, _otherInfo string) returns()
func (_ARStore *ARStoreSession) UpdatePaymentInfoOfAR(_payId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdatePaymentInfoOfAR(&_ARStore.TransactOpts, _payId, _otherInfo)
}

// UpdatePaymentInfoOfAR is a paid mutator transaction binding the contract method 0xaa86c09e.
//
// Solidity: function updatePaymentInfoOfAR(_payId string, _otherInfo string) returns()
func (_ARStore *ARStoreTransactorSession) UpdatePaymentInfoOfAR(_payId string, _otherInfo string) (*types.Transaction, error) {
	return _ARStore.Contract.UpdatePaymentInfoOfAR(&_ARStore.TransactOpts, _payId, _otherInfo)
}

// ARStoreARDiscountIterator is returned from FilterARDiscount and is used to iterate over the raw logs and unpacked data for ARDiscount events raised by the ARStore contract.
type ARStoreARDiscountIterator struct {
	Event *ARStoreARDiscount // Event containing the contract specifics and raw log

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
func (it *ARStoreARDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreARDiscount)
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
		it.Event = new(ARStoreARDiscount)
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
func (it *ARStoreARDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreARDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreARDiscount represents a ARDiscount event raised by the ARStore contract.
type ARStoreARDiscount struct {
	DiscountId        string
	OtherDiscountInfo string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterARDiscount is a free log retrieval operation binding the contract event 0xee6188daa001b6adba3924f5aef1e43bbb3894d70c7e26c5376d54a9cdd6329f.
//
// Solidity: event ARDiscount(discountId string, otherDiscountInfo string)
func (_ARStore *ARStoreFilterer) FilterARDiscount(opts *bind.FilterOpts) (*ARStoreARDiscountIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "ARDiscount")
	if err != nil {
		return nil, err
	}
	return &ARStoreARDiscountIterator{contract: _ARStore.contract, event: "ARDiscount", logs: logs, sub: sub}, nil
}

// WatchARDiscount is a free log subscription operation binding the contract event 0xee6188daa001b6adba3924f5aef1e43bbb3894d70c7e26c5376d54a9cdd6329f.
//
// Solidity: event ARDiscount(discountId string, otherDiscountInfo string)
func (_ARStore *ARStoreFilterer) WatchARDiscount(opts *bind.WatchOpts, sink chan<- *ARStoreARDiscount) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "ARDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreARDiscount)
				if err := _ARStore.contract.UnpackLog(event, "ARDiscount", log); err != nil {
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

// ARStoreARDiscountTxUpdateIterator is returned from FilterARDiscountTxUpdate and is used to iterate over the raw logs and unpacked data for ARDiscountTxUpdate events raised by the ARStore contract.
type ARStoreARDiscountTxUpdateIterator struct {
	Event *ARStoreARDiscountTxUpdate // Event containing the contract specifics and raw log

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
func (it *ARStoreARDiscountTxUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreARDiscountTxUpdate)
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
		it.Event = new(ARStoreARDiscountTxUpdate)
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
func (it *ARStoreARDiscountTxUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreARDiscountTxUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreARDiscountTxUpdate represents a ARDiscountTxUpdate event raised by the ARStore contract.
type ARStoreARDiscountTxUpdate struct {
	DiscountId        string
	OtherDiscountInfo string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterARDiscountTxUpdate is a free log retrieval operation binding the contract event 0xa63efca52b1b290c061dbda267833fd5f7ab0e823c41b484dde98b5fafd15c88.
//
// Solidity: event ARDiscountTxUpdate(discountId string, otherDiscountInfo string)
func (_ARStore *ARStoreFilterer) FilterARDiscountTxUpdate(opts *bind.FilterOpts) (*ARStoreARDiscountTxUpdateIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "ARDiscountTxUpdate")
	if err != nil {
		return nil, err
	}
	return &ARStoreARDiscountTxUpdateIterator{contract: _ARStore.contract, event: "ARDiscountTxUpdate", logs: logs, sub: sub}, nil
}

// WatchARDiscountTxUpdate is a free log subscription operation binding the contract event 0xa63efca52b1b290c061dbda267833fd5f7ab0e823c41b484dde98b5fafd15c88.
//
// Solidity: event ARDiscountTxUpdate(discountId string, otherDiscountInfo string)
func (_ARStore *ARStoreFilterer) WatchARDiscountTxUpdate(opts *bind.WatchOpts, sink chan<- *ARStoreARDiscountTxUpdate) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "ARDiscountTxUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreARDiscountTxUpdate)
				if err := _ARStore.contract.UnpackLog(event, "ARDiscountTxUpdate", log); err != nil {
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

// ARStoreARInfoUpdateIterator is returned from FilterARInfoUpdate and is used to iterate over the raw logs and unpacked data for ARInfoUpdate events raised by the ARStore contract.
type ARStoreARInfoUpdateIterator struct {
	Event *ARStoreARInfoUpdate // Event containing the contract specifics and raw log

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
func (it *ARStoreARInfoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreARInfoUpdate)
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
		it.Event = new(ARStoreARInfoUpdate)
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
func (it *ARStoreARInfoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreARInfoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreARInfoUpdate represents a ARInfoUpdate event raised by the ARStore contract.
type ARStoreARInfoUpdate struct {
	ArId            string
	OtherDiscountId string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterARInfoUpdate is a free log retrieval operation binding the contract event 0x5825912fe71c9873e5908194054f98265d4c946e7b6486e4b2789010dff86540.
//
// Solidity: event ARInfoUpdate(arId string, otherDiscountId string)
func (_ARStore *ARStoreFilterer) FilterARInfoUpdate(opts *bind.FilterOpts) (*ARStoreARInfoUpdateIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "ARInfoUpdate")
	if err != nil {
		return nil, err
	}
	return &ARStoreARInfoUpdateIterator{contract: _ARStore.contract, event: "ARInfoUpdate", logs: logs, sub: sub}, nil
}

// WatchARInfoUpdate is a free log subscription operation binding the contract event 0x5825912fe71c9873e5908194054f98265d4c946e7b6486e4b2789010dff86540.
//
// Solidity: event ARInfoUpdate(arId string, otherDiscountId string)
func (_ARStore *ARStoreFilterer) WatchARInfoUpdate(opts *bind.WatchOpts, sink chan<- *ARStoreARInfoUpdate) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "ARInfoUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreARInfoUpdate)
				if err := _ARStore.contract.UnpackLog(event, "ARInfoUpdate", log); err != nil {
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

// ARStoreARPayIterator is returned from FilterARPay and is used to iterate over the raw logs and unpacked data for ARPay events raised by the ARStore contract.
type ARStoreARPayIterator struct {
	Event *ARStoreARPay // Event containing the contract specifics and raw log

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
func (it *ARStoreARPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreARPay)
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
		it.Event = new(ARStoreARPay)
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
func (it *ARStoreARPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreARPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreARPay represents a ARPay event raised by the ARStore contract.
type ARStoreARPay struct {
	PayId          string
	OtherARPayInfo string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterARPay is a free log retrieval operation binding the contract event 0xde6e0d80d761e7ca1ca83e4e4d8c7e2f4aa2318bf51dfbfe95d36b9aa2808c46.
//
// Solidity: event ARPay(payId string, otherARPayInfo string)
func (_ARStore *ARStoreFilterer) FilterARPay(opts *bind.FilterOpts) (*ARStoreARPayIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "ARPay")
	if err != nil {
		return nil, err
	}
	return &ARStoreARPayIterator{contract: _ARStore.contract, event: "ARPay", logs: logs, sub: sub}, nil
}

// WatchARPay is a free log subscription operation binding the contract event 0xde6e0d80d761e7ca1ca83e4e4d8c7e2f4aa2318bf51dfbfe95d36b9aa2808c46.
//
// Solidity: event ARPay(payId string, otherARPayInfo string)
func (_ARStore *ARStoreFilterer) WatchARPay(opts *bind.WatchOpts, sink chan<- *ARStoreARPay) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "ARPay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreARPay)
				if err := _ARStore.contract.UnpackLog(event, "ARPay", log); err != nil {
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

// ARStoreARPaymentTxUpdateIterator is returned from FilterARPaymentTxUpdate and is used to iterate over the raw logs and unpacked data for ARPaymentTxUpdate events raised by the ARStore contract.
type ARStoreARPaymentTxUpdateIterator struct {
	Event *ARStoreARPaymentTxUpdate // Event containing the contract specifics and raw log

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
func (it *ARStoreARPaymentTxUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreARPaymentTxUpdate)
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
		it.Event = new(ARStoreARPaymentTxUpdate)
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
func (it *ARStoreARPaymentTxUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreARPaymentTxUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreARPaymentTxUpdate represents a ARPaymentTxUpdate event raised by the ARStore contract.
type ARStoreARPaymentTxUpdate struct {
	PayId          string
	OtherARPayInfo string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterARPaymentTxUpdate is a free log retrieval operation binding the contract event 0xd3be6a47b550b0ca84c64b6de4a62e30572a14922fce33b89da9e777923e89d3.
//
// Solidity: event ARPaymentTxUpdate(payId string, otherARPayInfo string)
func (_ARStore *ARStoreFilterer) FilterARPaymentTxUpdate(opts *bind.FilterOpts) (*ARStoreARPaymentTxUpdateIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "ARPaymentTxUpdate")
	if err != nil {
		return nil, err
	}
	return &ARStoreARPaymentTxUpdateIterator{contract: _ARStore.contract, event: "ARPaymentTxUpdate", logs: logs, sub: sub}, nil
}

// WatchARPaymentTxUpdate is a free log subscription operation binding the contract event 0xd3be6a47b550b0ca84c64b6de4a62e30572a14922fce33b89da9e777923e89d3.
//
// Solidity: event ARPaymentTxUpdate(payId string, otherARPayInfo string)
func (_ARStore *ARStoreFilterer) WatchARPaymentTxUpdate(opts *bind.WatchOpts, sink chan<- *ARStoreARPaymentTxUpdate) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "ARPaymentTxUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreARPaymentTxUpdate)
				if err := _ARStore.contract.UnpackLog(event, "ARPaymentTxUpdate", log); err != nil {
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

// ARStoreAddCompanyIterator is returned from FilterAddCompany and is used to iterate over the raw logs and unpacked data for AddCompany events raised by the ARStore contract.
type ARStoreAddCompanyIterator struct {
	Event *ARStoreAddCompany // Event containing the contract specifics and raw log

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
func (it *ARStoreAddCompanyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreAddCompany)
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
		it.Event = new(ARStoreAddCompany)
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
func (it *ARStoreAddCompanyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreAddCompanyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreAddCompany represents a AddCompany event raised by the ARStore contract.
type ARStoreAddCompany struct {
	CompanyId        string
	OtherCompanyInfo string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddCompany is a free log retrieval operation binding the contract event 0x4da365beff783759f53b7fb0227cf09776d94a116be497d8c967223d38740c73.
//
// Solidity: event AddCompany(companyId string, otherCompanyInfo string)
func (_ARStore *ARStoreFilterer) FilterAddCompany(opts *bind.FilterOpts) (*ARStoreAddCompanyIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "AddCompany")
	if err != nil {
		return nil, err
	}
	return &ARStoreAddCompanyIterator{contract: _ARStore.contract, event: "AddCompany", logs: logs, sub: sub}, nil
}

// WatchAddCompany is a free log subscription operation binding the contract event 0x4da365beff783759f53b7fb0227cf09776d94a116be497d8c967223d38740c73.
//
// Solidity: event AddCompany(companyId string, otherCompanyInfo string)
func (_ARStore *ARStoreFilterer) WatchAddCompany(opts *bind.WatchOpts, sink chan<- *ARStoreAddCompany) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "AddCompany")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreAddCompany)
				if err := _ARStore.contract.UnpackLog(event, "AddCompany", log); err != nil {
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

// ARStoreAddContractIterator is returned from FilterAddContract and is used to iterate over the raw logs and unpacked data for AddContract events raised by the ARStore contract.
type ARStoreAddContractIterator struct {
	Event *ARStoreAddContract // Event containing the contract specifics and raw log

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
func (it *ARStoreAddContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreAddContract)
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
		it.Event = new(ARStoreAddContract)
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
func (it *ARStoreAddContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreAddContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreAddContract represents a AddContract event raised by the ARStore contract.
type ARStoreAddContract struct {
	ContractNo        string
	OtherContractInfo string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAddContract is a free log retrieval operation binding the contract event 0x516729973459e4298de976b6e109f87200dce20f85caa6cadb58afd5aca2a415.
//
// Solidity: event AddContract(contractNo string, otherContractInfo string)
func (_ARStore *ARStoreFilterer) FilterAddContract(opts *bind.FilterOpts) (*ARStoreAddContractIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "AddContract")
	if err != nil {
		return nil, err
	}
	return &ARStoreAddContractIterator{contract: _ARStore.contract, event: "AddContract", logs: logs, sub: sub}, nil
}

// WatchAddContract is a free log subscription operation binding the contract event 0x516729973459e4298de976b6e109f87200dce20f85caa6cadb58afd5aca2a415.
//
// Solidity: event AddContract(contractNo string, otherContractInfo string)
func (_ARStore *ARStoreFilterer) WatchAddContract(opts *bind.WatchOpts, sink chan<- *ARStoreAddContract) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "AddContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreAddContract)
				if err := _ARStore.contract.UnpackLog(event, "AddContract", log); err != nil {
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

// ARStoreCompanyInfoUpdateIterator is returned from FilterCompanyInfoUpdate and is used to iterate over the raw logs and unpacked data for CompanyInfoUpdate events raised by the ARStore contract.
type ARStoreCompanyInfoUpdateIterator struct {
	Event *ARStoreCompanyInfoUpdate // Event containing the contract specifics and raw log

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
func (it *ARStoreCompanyInfoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreCompanyInfoUpdate)
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
		it.Event = new(ARStoreCompanyInfoUpdate)
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
func (it *ARStoreCompanyInfoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreCompanyInfoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreCompanyInfoUpdate represents a CompanyInfoUpdate event raised by the ARStore contract.
type ARStoreCompanyInfoUpdate struct {
	CompanyId        string
	OtherCompanyInfo string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCompanyInfoUpdate is a free log retrieval operation binding the contract event 0xc5d4c08e7ddc547a2cc0c88249f55ca6f79db43f6b05cb9a9b426bea829384e6.
//
// Solidity: event CompanyInfoUpdate(companyId string, otherCompanyInfo string)
func (_ARStore *ARStoreFilterer) FilterCompanyInfoUpdate(opts *bind.FilterOpts) (*ARStoreCompanyInfoUpdateIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "CompanyInfoUpdate")
	if err != nil {
		return nil, err
	}
	return &ARStoreCompanyInfoUpdateIterator{contract: _ARStore.contract, event: "CompanyInfoUpdate", logs: logs, sub: sub}, nil
}

// WatchCompanyInfoUpdate is a free log subscription operation binding the contract event 0xc5d4c08e7ddc547a2cc0c88249f55ca6f79db43f6b05cb9a9b426bea829384e6.
//
// Solidity: event CompanyInfoUpdate(companyId string, otherCompanyInfo string)
func (_ARStore *ARStoreFilterer) WatchCompanyInfoUpdate(opts *bind.WatchOpts, sink chan<- *ARStoreCompanyInfoUpdate) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "CompanyInfoUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreCompanyInfoUpdate)
				if err := _ARStore.contract.UnpackLog(event, "CompanyInfoUpdate", log); err != nil {
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

// ARStoreContractInfoUpdateIterator is returned from FilterContractInfoUpdate and is used to iterate over the raw logs and unpacked data for ContractInfoUpdate events raised by the ARStore contract.
type ARStoreContractInfoUpdateIterator struct {
	Event *ARStoreContractInfoUpdate // Event containing the contract specifics and raw log

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
func (it *ARStoreContractInfoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreContractInfoUpdate)
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
		it.Event = new(ARStoreContractInfoUpdate)
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
func (it *ARStoreContractInfoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreContractInfoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreContractInfoUpdate represents a ContractInfoUpdate event raised by the ARStore contract.
type ARStoreContractInfoUpdate struct {
	ContractNo        string
	OtherContractInfo string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterContractInfoUpdate is a free log retrieval operation binding the contract event 0xf9d4226c2fb9d76fcdb373d47ab8e58acbd70b9c4c1f94792cad9d6de3fc612e.
//
// Solidity: event ContractInfoUpdate(contractNo string, otherContractInfo string)
func (_ARStore *ARStoreFilterer) FilterContractInfoUpdate(opts *bind.FilterOpts) (*ARStoreContractInfoUpdateIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "ContractInfoUpdate")
	if err != nil {
		return nil, err
	}
	return &ARStoreContractInfoUpdateIterator{contract: _ARStore.contract, event: "ContractInfoUpdate", logs: logs, sub: sub}, nil
}

// WatchContractInfoUpdate is a free log subscription operation binding the contract event 0xf9d4226c2fb9d76fcdb373d47ab8e58acbd70b9c4c1f94792cad9d6de3fc612e.
//
// Solidity: event ContractInfoUpdate(contractNo string, otherContractInfo string)
func (_ARStore *ARStoreFilterer) WatchContractInfoUpdate(opts *bind.WatchOpts, sink chan<- *ARStoreContractInfoUpdate) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "ContractInfoUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreContractInfoUpdate)
				if err := _ARStore.contract.UnpackLog(event, "ContractInfoUpdate", log); err != nil {
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

// ARStoreCreateARIterator is returned from FilterCreateAR and is used to iterate over the raw logs and unpacked data for CreateAR events raised by the ARStore contract.
type ARStoreCreateARIterator struct {
	Event *ARStoreCreateAR // Event containing the contract specifics and raw log

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
func (it *ARStoreCreateARIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreCreateAR)
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
		it.Event = new(ARStoreCreateAR)
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
func (it *ARStoreCreateARIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreCreateARIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreCreateAR represents a CreateAR event raised by the ARStore contract.
type ARStoreCreateAR struct {
	ArId        string
	OtherARInfo string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateAR is a free log retrieval operation binding the contract event 0x05d2736ab72b013950666611bf9f68c9f4736b6033f081a04c2f9ac7cb8d52b6.
//
// Solidity: event CreateAR(arId string, otherARInfo string)
func (_ARStore *ARStoreFilterer) FilterCreateAR(opts *bind.FilterOpts) (*ARStoreCreateARIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "CreateAR")
	if err != nil {
		return nil, err
	}
	return &ARStoreCreateARIterator{contract: _ARStore.contract, event: "CreateAR", logs: logs, sub: sub}, nil
}

// WatchCreateAR is a free log subscription operation binding the contract event 0x05d2736ab72b013950666611bf9f68c9f4736b6033f081a04c2f9ac7cb8d52b6.
//
// Solidity: event CreateAR(arId string, otherARInfo string)
func (_ARStore *ARStoreFilterer) WatchCreateAR(opts *bind.WatchOpts, sink chan<- *ARStoreCreateAR) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "CreateAR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreCreateAR)
				if err := _ARStore.contract.UnpackLog(event, "CreateAR", log); err != nil {
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

// ARStoreRegisterUserIterator is returned from FilterRegisterUser and is used to iterate over the raw logs and unpacked data for RegisterUser events raised by the ARStore contract.
type ARStoreRegisterUserIterator struct {
	Event *ARStoreRegisterUser // Event containing the contract specifics and raw log

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
func (it *ARStoreRegisterUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARStoreRegisterUser)
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
		it.Event = new(ARStoreRegisterUser)
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
func (it *ARStoreRegisterUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARStoreRegisterUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARStoreRegisterUser represents a RegisterUser event raised by the ARStore contract.
type ARStoreRegisterUser struct {
	UserId    string
	BcAccount string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegisterUser is a free log retrieval operation binding the contract event 0x8c98eb51f6ba14256776d52666f6b14b282f139267f36f8491d651a83b06752d.
//
// Solidity: event RegisterUser(userId string, bcAccount string)
func (_ARStore *ARStoreFilterer) FilterRegisterUser(opts *bind.FilterOpts) (*ARStoreRegisterUserIterator, error) {

	logs, sub, err := _ARStore.contract.FilterLogs(opts, "RegisterUser")
	if err != nil {
		return nil, err
	}
	return &ARStoreRegisterUserIterator{contract: _ARStore.contract, event: "RegisterUser", logs: logs, sub: sub}, nil
}

// WatchRegisterUser is a free log subscription operation binding the contract event 0x8c98eb51f6ba14256776d52666f6b14b282f139267f36f8491d651a83b06752d.
//
// Solidity: event RegisterUser(userId string, bcAccount string)
func (_ARStore *ARStoreFilterer) WatchRegisterUser(opts *bind.WatchOpts, sink chan<- *ARStoreRegisterUser) (event.Subscription, error) {

	logs, sub, err := _ARStore.contract.WatchLogs(opts, "RegisterUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARStoreRegisterUser)
				if err := _ARStore.contract.UnpackLog(event, "RegisterUser", log); err != nil {
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
