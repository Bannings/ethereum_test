// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CfNodesABI is the input ABI used to generate the binding from.
const CfNodesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"nodeName\",\"type\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\"},{\"name\":\"otherNodeInfo\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeName\",\"type\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteAllNodes\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeArrayCounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeName\",\"type\":\"string\"}],\"name\":\"getNodeKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeName\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeInfoByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeName\",\"type\":\"string\"}],\"name\":\"isNodeExisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"adminAccount\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherNodeInfo\",\"type\":\"string\"}],\"name\":\"AddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"}],\"name\":\"DeleteNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"UpdateNode\",\"type\":\"event\"}]"

// CfNodesBin is the compiled bytecode used for deploying new contracts.
const CfNodesBin = `0x6060604052341561000f57600080fd5b6040516020806112b3833981016040528080519150505b600160a060020a038082166000908152602081905260408082206001908190553390931682529020555b505b611252806100616000396000f3006060604052361561008b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166303f1bc4f8114610090578063517fb203146100c857806362d8c667146100f457806391d9d9421461010957806393d5badf1461012e578063a0c15b7714610262578063d53da5b814610282578063f99101f5146103e5575b600080fd5b341561009b57600080fd5b6100c6602460048035828101929082013591813580830192908201359160443591820191013561044a565b005b34156100d357600080fd5b6100c660246004803582810192908201359181359182019101356106be565b005b34156100ff57600080fd5b6100c66107dd565b005b341561011457600080fd5b61011c6108e0565b60405190815260200160405180910390f35b341561013957600080fd5b61017f60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506108e795505050505050565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156101c15780820151818401525b6020016101a8565b50505050905090810190601f1680156101ee5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156102255780820151818401525b60200161020c565b50505050905090810190601f1680156102525780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b341561026d57600080fd5b6100c66004803560248101910135610b09565b005b341561028d57600080fd5b610298600435610c1c565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156102de5780820151818401525b6020016102c5565b50505050905090810190601f16801561030b5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156103425780820151818401525b602001610329565b50505050905090810190601f16801561036f5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156103a65780820151818401525b60200161038d565b50505050905090810190601f1680156103d35780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34156103f057600080fd5b61043660046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610f5395505050505050565b604051901515815260200160405180910390f35b73ffffffffffffffffffffffffffffffffffffffff331660009081526020819052604090205460011461047c57600080fd5b60018686604051808383808284378201915050925050509081526020016040519081900390205460026000196101006001841615020190911604156104c057600080fd5b60028686604051808383808284378201915050925050509081526020016040519081900390205460ff1615156001146105595760038054600181016105058382610fda565b916000526020600020900160005b5061051f908888611004565b505060016002878760405180838380828437820191505092505050908152602001604051908190039020805460ff19169115159190911790555b604080519081016040528085858080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050815260200183838080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050815250600187876040518083838082843782019150509250505090815260200160405190819003902081518190805161060a929160200190611083565b50602082015181600101908051610625929160200190611083565b509050507f92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd86868686868660405160608082528101869052806020810160408201608083018a8a8082843790910185810384528881526020019050888880828437909101858103835286815260200190508686808284378201915050995050505050505050505060405180910390a15b5b505050505050565b73ffffffffffffffffffffffffffffffffffffffff33166000908152602081905260409020546001146106f057600080fd5b60018484604051808383808284378201915050925050509081526020016040519081900390205460026000196101006001841615020190911604151561073557600080fd5b81816001868660405180838380828437820191505092505050908152602001604051908190039020610768929091611004565b507fddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a8484848460405160408082528101849052806020810160608201878780828437909101848103835285815260200190508585808284378201915050965050505050505060405180910390a15b5b50505050565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260208190526040812054819060011461081157600080fd5b505060035460005b818110156108da5760206040519081016040526000815260038054600191908490811061084257fe5b906000526020600020900160005b5060405180828054600181600116156101000203166002900480156108ac5780601f1061088a5761010080835404028352918201916108ac565b820191906000526020600020905b815481529060010190602001808311610898575b505092835250506020016040519081900390209080516108d0929160200190611083565b505b600101610819565b5b5b5050565b6003545b90565b6108ef611181565b6108f7611181565b6001836040518082805190602001908083835b6020831061092a57805182525b601f19909201916020918201910161090a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206001846040518082805190602001908083835b6020831061099157805182525b601f199092019160209182019101610971565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600101818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a5c5780601f10610a3157610100808354040283529160200191610a5c565b820191906000526020600020905b815481529060010190602001808311610a3f57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610af85780601f10610acd57610100808354040283529160200191610af8565b820191906000526020600020905b815481529060010190602001808311610adb57829003601f168201915b50505050509050915091505b915091565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260208190526040902054600114610b3b57600080fd5b600182826040518083838082843782019150509250505090815260200160405190819003902054600260001961010060018416150201909116041515610b8057600080fd5b60206040519081016040528060008152506001838360405180838380828437820191505092505050908152602001604051908190039020908051610bc8929160200190611083565b507fba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a82826040516020808252810182905280604081018484808284378201915050935050505060405180910390a15b5b5050565b610c24611181565b610c2c611181565b610c34611181565b6003805485908110610c4257fe5b906000526020600020900160005b506001600386815481101515610c6257fe5b906000526020600020900160005b506040518082805460018160011615610100020316600290048015610ccc5780601f10610caa576101008083540402835291820191610ccc565b820191906000526020600020905b815481529060010190602001808311610cb8575b5050928352505060200160405180910390206000016001600387815481101515610cf257fe5b906000526020600020900160005b506040518082805460018160011615610100020316600290048015610d5c5780601f10610d3a576101008083540402835291820191610d5c565b820191906000526020600020905b815481529060010190602001808311610d48575b505092835250506020016040518091039020600101828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e065780601f10610ddb57610100808354040283529160200191610e06565b820191906000526020600020905b815481529060010190602001808311610de957829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ea25780601f10610e7757610100808354040283529160200191610ea2565b820191906000526020600020905b815481529060010190602001808311610e8557829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f3e5780601f10610f1357610100808354040283529160200191610f3e565b820191906000526020600020905b815481529060010190602001808311610f2157829003601f168201915b505050505090509250925092505b9193909250565b60006001826040518082805190602001908083835b60208310610f8857805182525b601f199092019160209182019101610f68565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205460026000196101006001841615020190911604151590505b919050565b815481835581811511610ffe57600083815260209020610ffe918101908301611193565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110455782800160ff19823516178555611072565b82800160010185558215611072579182015b82811115611072578235825591602001919060010190611057565b5b5061107f9291506111bd565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110c457805160ff1916838001178555611072565b82800160010185558215611072579182015b828111156110725782518255916020019190600101906110d6565b5b5061107f9291506111bd565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110c457805160ff1916838001178555611072565b82800160010185558215611072579182015b828111156110725782518255916020019190600101906110d6565b5b5061107f9291506111bd565b5090565b60206040519081016040526000815290565b6108e491905b8082111561107f5760006111ad82826111de565b50600101611199565b5090565b90565b6108e491905b8082111561107f57600081556001016111c3565b5090565b90565b50805460018160011615610100020316600290046000825580601f106112045750611222565b601f01602090049060005260206000209081019061122291906111bd565b5b505600a165627a7a723058207c433f079635bc0fdf9b0383d05e89e4052b056e047daae55482c3cec7e2cc310029`

// DeployCfNodes deploys a new Ethereum contract, binding an instance of CfNodes to it.
func DeployCfNodes(auth *bind.TransactOpts, backend bind.ContractBackend, adminAccount common.Address) (common.Address, *types.Transaction, *CfNodes, error) {
	parsed, err := abi.JSON(strings.NewReader(CfNodesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CfNodesBin), backend, adminAccount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CfNodes{CfNodesCaller: CfNodesCaller{contract: contract}, CfNodesTransactor: CfNodesTransactor{contract: contract}, CfNodesFilterer: CfNodesFilterer{contract: contract}}, nil
}

// CfNodes is an auto generated Go binding around an Ethereum contract.
type CfNodes struct {
	CfNodesCaller     // Read-only binding to the contract
	CfNodesTransactor // Write-only binding to the contract
	CfNodesFilterer   // Log filterer for contract events
}

// CfNodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CfNodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CfNodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CfNodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CfNodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CfNodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CfNodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CfNodesSession struct {
	Contract     *CfNodes          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CfNodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CfNodesCallerSession struct {
	Contract *CfNodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CfNodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CfNodesTransactorSession struct {
	Contract     *CfNodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CfNodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CfNodesRaw struct {
	Contract *CfNodes // Generic contract binding to access the raw methods on
}

// CfNodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CfNodesCallerRaw struct {
	Contract *CfNodesCaller // Generic read-only contract binding to access the raw methods on
}

// CfNodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CfNodesTransactorRaw struct {
	Contract *CfNodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCfNodes creates a new instance of CfNodes, bound to a specific deployed contract.
func NewCfNodes(address common.Address, backend bind.ContractBackend) (*CfNodes, error) {
	contract, err := bindCfNodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CfNodes{CfNodesCaller: CfNodesCaller{contract: contract}, CfNodesTransactor: CfNodesTransactor{contract: contract}, CfNodesFilterer: CfNodesFilterer{contract: contract}}, nil
}

// NewCfNodesCaller creates a new read-only instance of CfNodes, bound to a specific deployed contract.
func NewCfNodesCaller(address common.Address, caller bind.ContractCaller) (*CfNodesCaller, error) {
	contract, err := bindCfNodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CfNodesCaller{contract: contract}, nil
}

// NewCfNodesTransactor creates a new write-only instance of CfNodes, bound to a specific deployed contract.
func NewCfNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*CfNodesTransactor, error) {
	contract, err := bindCfNodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CfNodesTransactor{contract: contract}, nil
}

// NewCfNodesFilterer creates a new log filterer instance of CfNodes, bound to a specific deployed contract.
func NewCfNodesFilterer(address common.Address, filterer bind.ContractFilterer) (*CfNodesFilterer, error) {
	contract, err := bindCfNodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CfNodesFilterer{contract: contract}, nil
}

// bindCfNodes binds a generic wrapper to an already deployed contract.
func bindCfNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CfNodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CfNodes *CfNodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CfNodes.Contract.CfNodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CfNodes *CfNodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CfNodes.Contract.CfNodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CfNodes *CfNodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CfNodes.Contract.CfNodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CfNodes *CfNodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CfNodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CfNodes *CfNodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CfNodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CfNodes *CfNodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CfNodes.Contract.contract.Transact(opts, method, params...)
}

// GetNodeArrayCounts is a free data retrieval call binding the contract method 0x91d9d942.
//
// Solidity: function getNodeArrayCounts() constant returns(uint256)
func (_CfNodes *CfNodesCaller) GetNodeArrayCounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CfNodes.contract.Call(opts, out, "getNodeArrayCounts")
	return *ret0, err
}

// GetNodeArrayCounts is a free data retrieval call binding the contract method 0x91d9d942.
//
// Solidity: function getNodeArrayCounts() constant returns(uint256)
func (_CfNodes *CfNodesSession) GetNodeArrayCounts() (*big.Int, error) {
	return _CfNodes.Contract.GetNodeArrayCounts(&_CfNodes.CallOpts)
}

// GetNodeArrayCounts is a free data retrieval call binding the contract method 0x91d9d942.
//
// Solidity: function getNodeArrayCounts() constant returns(uint256)
func (_CfNodes *CfNodesCallerSession) GetNodeArrayCounts() (*big.Int, error) {
	return _CfNodes.Contract.GetNodeArrayCounts(&_CfNodes.CallOpts)
}

// GetNodeInfoByIndex is a free data retrieval call binding the contract method 0xd53da5b8.
//
// Solidity: function getNodeInfoByIndex(index uint256) constant returns(string, string, string)
func (_CfNodes *CfNodesCaller) GetNodeInfoByIndex(opts *bind.CallOpts, index *big.Int) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _CfNodes.contract.Call(opts, out, "getNodeInfoByIndex", index)
	return *ret0, *ret1, *ret2, err
}

// GetNodeInfoByIndex is a free data retrieval call binding the contract method 0xd53da5b8.
//
// Solidity: function getNodeInfoByIndex(index uint256) constant returns(string, string, string)
func (_CfNodes *CfNodesSession) GetNodeInfoByIndex(index *big.Int) (string, string, string, error) {
	return _CfNodes.Contract.GetNodeInfoByIndex(&_CfNodes.CallOpts, index)
}

// GetNodeInfoByIndex is a free data retrieval call binding the contract method 0xd53da5b8.
//
// Solidity: function getNodeInfoByIndex(index uint256) constant returns(string, string, string)
func (_CfNodes *CfNodesCallerSession) GetNodeInfoByIndex(index *big.Int) (string, string, string, error) {
	return _CfNodes.Contract.GetNodeInfoByIndex(&_CfNodes.CallOpts, index)
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(nodeName string) constant returns(string, string)
func (_CfNodes *CfNodesCaller) GetNodeKey(opts *bind.CallOpts, nodeName string) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _CfNodes.contract.Call(opts, out, "getNodeKey", nodeName)
	return *ret0, *ret1, err
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(nodeName string) constant returns(string, string)
func (_CfNodes *CfNodesSession) GetNodeKey(nodeName string) (string, string, error) {
	return _CfNodes.Contract.GetNodeKey(&_CfNodes.CallOpts, nodeName)
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(nodeName string) constant returns(string, string)
func (_CfNodes *CfNodesCallerSession) GetNodeKey(nodeName string) (string, string, error) {
	return _CfNodes.Contract.GetNodeKey(&_CfNodes.CallOpts, nodeName)
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(nodeName string) constant returns(bool)
func (_CfNodes *CfNodesCaller) IsNodeExisted(opts *bind.CallOpts, nodeName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CfNodes.contract.Call(opts, out, "isNodeExisted", nodeName)
	return *ret0, err
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(nodeName string) constant returns(bool)
func (_CfNodes *CfNodesSession) IsNodeExisted(nodeName string) (bool, error) {
	return _CfNodes.Contract.IsNodeExisted(&_CfNodes.CallOpts, nodeName)
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(nodeName string) constant returns(bool)
func (_CfNodes *CfNodesCallerSession) IsNodeExisted(nodeName string) (bool, error) {
	return _CfNodes.Contract.IsNodeExisted(&_CfNodes.CallOpts, nodeName)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(nodeName string, publicKey string, otherNodeInfo string) returns()
func (_CfNodes *CfNodesTransactor) AddNode(opts *bind.TransactOpts, nodeName string, publicKey string, otherNodeInfo string) (*types.Transaction, error) {
	return _CfNodes.contract.Transact(opts, "addNode", nodeName, publicKey, otherNodeInfo)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(nodeName string, publicKey string, otherNodeInfo string) returns()
func (_CfNodes *CfNodesSession) AddNode(nodeName string, publicKey string, otherNodeInfo string) (*types.Transaction, error) {
	return _CfNodes.Contract.AddNode(&_CfNodes.TransactOpts, nodeName, publicKey, otherNodeInfo)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(nodeName string, publicKey string, otherNodeInfo string) returns()
func (_CfNodes *CfNodesTransactorSession) AddNode(nodeName string, publicKey string, otherNodeInfo string) (*types.Transaction, error) {
	return _CfNodes.Contract.AddNode(&_CfNodes.TransactOpts, nodeName, publicKey, otherNodeInfo)
}

// DeleteAllNodes is a paid mutator transaction binding the contract method 0x62d8c667.
//
// Solidity: function deleteAllNodes() returns()
func (_CfNodes *CfNodesTransactor) DeleteAllNodes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CfNodes.contract.Transact(opts, "deleteAllNodes")
}

// DeleteAllNodes is a paid mutator transaction binding the contract method 0x62d8c667.
//
// Solidity: function deleteAllNodes() returns()
func (_CfNodes *CfNodesSession) DeleteAllNodes() (*types.Transaction, error) {
	return _CfNodes.Contract.DeleteAllNodes(&_CfNodes.TransactOpts)
}

// DeleteAllNodes is a paid mutator transaction binding the contract method 0x62d8c667.
//
// Solidity: function deleteAllNodes() returns()
func (_CfNodes *CfNodesTransactorSession) DeleteAllNodes() (*types.Transaction, error) {
	return _CfNodes.Contract.DeleteAllNodes(&_CfNodes.TransactOpts)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(nodeName string) returns()
func (_CfNodes *CfNodesTransactor) DeleteNode(opts *bind.TransactOpts, nodeName string) (*types.Transaction, error) {
	return _CfNodes.contract.Transact(opts, "deleteNode", nodeName)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(nodeName string) returns()
func (_CfNodes *CfNodesSession) DeleteNode(nodeName string) (*types.Transaction, error) {
	return _CfNodes.Contract.DeleteNode(&_CfNodes.TransactOpts, nodeName)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(nodeName string) returns()
func (_CfNodes *CfNodesTransactorSession) DeleteNode(nodeName string) (*types.Transaction, error) {
	return _CfNodes.Contract.DeleteNode(&_CfNodes.TransactOpts, nodeName)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(nodeName string, publicKey string) returns()
func (_CfNodes *CfNodesTransactor) UpdateNode(opts *bind.TransactOpts, nodeName string, publicKey string) (*types.Transaction, error) {
	return _CfNodes.contract.Transact(opts, "updateNode", nodeName, publicKey)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(nodeName string, publicKey string) returns()
func (_CfNodes *CfNodesSession) UpdateNode(nodeName string, publicKey string) (*types.Transaction, error) {
	return _CfNodes.Contract.UpdateNode(&_CfNodes.TransactOpts, nodeName, publicKey)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(nodeName string, publicKey string) returns()
func (_CfNodes *CfNodesTransactorSession) UpdateNode(nodeName string, publicKey string) (*types.Transaction, error) {
	return _CfNodes.Contract.UpdateNode(&_CfNodes.TransactOpts, nodeName, publicKey)
}

// CfNodesAddNodeIterator is returned from FilterAddNode and is used to iterate over the raw logs and unpacked data for AddNode events raised by the CfNodes contract.
type CfNodesAddNodeIterator struct {
	Event *CfNodesAddNode // Event containing the contract specifics and raw log

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
func (it *CfNodesAddNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CfNodesAddNode)
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
		it.Event = new(CfNodesAddNode)
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
func (it *CfNodesAddNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CfNodesAddNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CfNodesAddNode represents a AddNode event raised by the CfNodes contract.
type CfNodesAddNode struct {
	NodeName      string
	PublicKey     string
	OtherNodeInfo string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddNode is a free log retrieval operation binding the contract event 0x92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd.
//
// Solidity: event AddNode(nodeName string, publicKey string, otherNodeInfo string)
func (_CfNodes *CfNodesFilterer) FilterAddNode(opts *bind.FilterOpts) (*CfNodesAddNodeIterator, error) {

	logs, sub, err := _CfNodes.contract.FilterLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return &CfNodesAddNodeIterator{contract: _CfNodes.contract, event: "AddNode", logs: logs, sub: sub}, nil
}

// WatchAddNode is a free log subscription operation binding the contract event 0x92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd.
//
// Solidity: event AddNode(nodeName string, publicKey string, otherNodeInfo string)
func (_CfNodes *CfNodesFilterer) WatchAddNode(opts *bind.WatchOpts, sink chan<- *CfNodesAddNode) (event.Subscription, error) {

	logs, sub, err := _CfNodes.contract.WatchLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CfNodesAddNode)
				if err := _CfNodes.contract.UnpackLog(event, "AddNode", log); err != nil {
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

// CfNodesDeleteNodeIterator is returned from FilterDeleteNode and is used to iterate over the raw logs and unpacked data for DeleteNode events raised by the CfNodes contract.
type CfNodesDeleteNodeIterator struct {
	Event *CfNodesDeleteNode // Event containing the contract specifics and raw log

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
func (it *CfNodesDeleteNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CfNodesDeleteNode)
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
		it.Event = new(CfNodesDeleteNode)
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
func (it *CfNodesDeleteNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CfNodesDeleteNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CfNodesDeleteNode represents a DeleteNode event raised by the CfNodes contract.
type CfNodesDeleteNode struct {
	NodeName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteNode is a free log retrieval operation binding the contract event 0xba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a.
//
// Solidity: event DeleteNode(nodeName string)
func (_CfNodes *CfNodesFilterer) FilterDeleteNode(opts *bind.FilterOpts) (*CfNodesDeleteNodeIterator, error) {

	logs, sub, err := _CfNodes.contract.FilterLogs(opts, "DeleteNode")
	if err != nil {
		return nil, err
	}
	return &CfNodesDeleteNodeIterator{contract: _CfNodes.contract, event: "DeleteNode", logs: logs, sub: sub}, nil
}

// WatchDeleteNode is a free log subscription operation binding the contract event 0xba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a.
//
// Solidity: event DeleteNode(nodeName string)
func (_CfNodes *CfNodesFilterer) WatchDeleteNode(opts *bind.WatchOpts, sink chan<- *CfNodesDeleteNode) (event.Subscription, error) {

	logs, sub, err := _CfNodes.contract.WatchLogs(opts, "DeleteNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CfNodesDeleteNode)
				if err := _CfNodes.contract.UnpackLog(event, "DeleteNode", log); err != nil {
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

// CfNodesUpdateNodeIterator is returned from FilterUpdateNode and is used to iterate over the raw logs and unpacked data for UpdateNode events raised by the CfNodes contract.
type CfNodesUpdateNodeIterator struct {
	Event *CfNodesUpdateNode // Event containing the contract specifics and raw log

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
func (it *CfNodesUpdateNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CfNodesUpdateNode)
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
		it.Event = new(CfNodesUpdateNode)
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
func (it *CfNodesUpdateNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CfNodesUpdateNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CfNodesUpdateNode represents a UpdateNode event raised by the CfNodes contract.
type CfNodesUpdateNode struct {
	NodeName  string
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateNode is a free log retrieval operation binding the contract event 0xddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a.
//
// Solidity: event UpdateNode(nodeName string, publicKey string)
func (_CfNodes *CfNodesFilterer) FilterUpdateNode(opts *bind.FilterOpts) (*CfNodesUpdateNodeIterator, error) {

	logs, sub, err := _CfNodes.contract.FilterLogs(opts, "UpdateNode")
	if err != nil {
		return nil, err
	}
	return &CfNodesUpdateNodeIterator{contract: _CfNodes.contract, event: "UpdateNode", logs: logs, sub: sub}, nil
}

// WatchUpdateNode is a free log subscription operation binding the contract event 0xddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a.
//
// Solidity: event UpdateNode(nodeName string, publicKey string)
func (_CfNodes *CfNodesFilterer) WatchUpdateNode(opts *bind.WatchOpts, sink chan<- *CfNodesUpdateNode) (event.Subscription, error) {

	logs, sub, err := _CfNodes.contract.WatchLogs(opts, "UpdateNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CfNodesUpdateNode)
				if err := _CfNodes.contract.UnpackLog(event, "UpdateNode", log); err != nil {
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
