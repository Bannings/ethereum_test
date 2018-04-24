// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dr_contracts

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

// CFNodesABI is the input ABI used to generate the binding from.
const CFNodesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherNodeInfo\",\"type\":\"string\"}],\"name\":\"AddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"UpdateNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"}],\"name\":\"DeleteNode\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"},{\"name\":\"_publicKey\",\"type\":\"string\"},{\"name\":\"_otherNodeInfo\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"}],\"name\":\"isNodeExisted\",\"outputs\":[{\"name\":\"existed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"}],\"name\":\"getNodeKey\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"name\":\"totalNodes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"},{\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CFNodesBin is the compiled bytecode used for deploying new contracts.
const CFNodesBin = `0x6060604052341561000f57600080fd5b610f168061001e6000396000f300606060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806303f1bc4f1461007d57806339bf397e14610160578063517fb2031461018957806393d5badf14610229578063a0c15b77146102ff578063f99101f51461035c575b600080fd5b341561008857600080fd5b61015e600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506103d1565b005b341561016b57600080fd5b6101736106f7565b6040518082815260200191505060405180910390f35b341561019457600080fd5b610227600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610701565b005b341561023457600080fd5b610284600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610915565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102c45780820151818401526020810190506102a9565b50505050905090810190601f1680156102f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561030a57600080fd5b61035a600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610aaf565b005b341561036757600080fd5b6103b7600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610c94565b604051808215151515815260200191505060405180910390f35b6103d9610d18565b6000846040518082805190602001908083835b60208310151561041157805182526020820191506020810190506020830392506103ec565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff1615151561045e57fe5b608060405190810160405280858152602001600115158152602001848152602001838152509050806000856040518082805190602001908083835b6020831015156104be5780518252602082019150602081019050602083039250610499565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082015181600001908051906020019061050d929190610d55565b5060208201518160010160006101000a81548160ff021916908315150217905550604082015181600201908051906020019061054a929190610d55565b506060820151816003019080519060200190610567929190610d55565b50905050600180600082825401925050819055507f92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd84848460405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156105e75780820151818401526020810190506105cc565b50505050905090810190601f1680156106145780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561064d578082015181840152602081019050610632565b50505050905090810190601f16801561067a5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156106b3578082015181840152602081019050610698565b50505050905090810190601f1680156106e05780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a150505050565b6000600154905090565b6000826040518082805190602001908083835b6020831015156107395780518252602082019150602081019050602083039250610714565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16151561078557fe5b806000836040518082805190602001908083835b6020831015156107be5780518252602082019150602081019050602083039250610799565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206002019080519060200190610807929190610dd5565b507fddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a8282604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561086f578082015181840152602081019050610854565b50505050905090810190601f16801561089c5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156108d55780820151818401526020810190506108ba565b50505050905090810190601f1680156109025780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15050565b61091d610e55565b6000826040518082805190602001908083835b6020831015156109555780518252602082019150602081019050602083039250610930565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff1615156109a157fe5b6000826040518082805190602001908083835b6020831015156109d957805182526020820191506020810190506020830392506109b4565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aa35780601f10610a7857610100808354040283529160200191610aa3565b820191906000526020600020905b815481529060010190602001808311610a8657829003601f168201915b50505050509050919050565b6000816040518082805190602001908083835b602083101515610ae75780518252602082019150602081019050602083039250610ac2565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff161515610b3357fe5b6000816040518082805190602001908083835b602083101515610b6b5780518252602082019150602081019050602083039250610b46565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008082016000610baf9190610e69565b6001820160006101000a81549060ff0219169055600282016000610bd39190610e69565b600382016000610be39190610e69565b5050600180600082825403925050819055507fba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a816040518080602001828103825283818151815260200191508051906020019080838360005b83811015610c57578082015181840152602081019050610c3c565b50505050905090810190601f168015610c845780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b600080826040518082805190602001908083835b602083101515610ccd5780518252602082019150602081019050602083039250610ca8565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff169050919050565b608060405190810160405280610d2c610eb1565b8152602001600015158152602001610d42610eb1565b8152602001610d4f610eb1565b81525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610d9657805160ff1916838001178555610dc4565b82800160010185558215610dc4579182015b82811115610dc3578251825591602001919060010190610da8565b5b509050610dd19190610ec5565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e1657805160ff1916838001178555610e44565b82800160010185558215610e44579182015b82811115610e43578251825591602001919060010190610e28565b5b509050610e519190610ec5565b5090565b602060405190810160405280600081525090565b50805460018160011615610100020316600290046000825580601f10610e8f5750610eae565b601f016020900490600052602060002090810190610ead9190610ec5565b5b50565b602060405190810160405280600081525090565b610ee791905b80821115610ee3576000816000905550600101610ecb565b5090565b905600a165627a7a72305820afbdd424e3e0db615cd28bf1f4dcdceaf799a12a9cbce07a5ce04fd36937c14a0029`

// DeployCFNodes deploys a new Ethereum contract, binding an instance of CFNodes to it.
func DeployCFNodes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CFNodes, error) {
	parsed, err := abi.JSON(strings.NewReader(CFNodesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CFNodesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CFNodes{CFNodesCaller: CFNodesCaller{contract: contract}, CFNodesTransactor: CFNodesTransactor{contract: contract}, CFNodesFilterer: CFNodesFilterer{contract: contract}}, nil
}

// CFNodes is an auto generated Go binding around an Ethereum contract.
type CFNodes struct {
	CFNodesCaller     // Read-only binding to the contract
	CFNodesTransactor // Write-only binding to the contract
	CFNodesFilterer   // Log filterer for contract events
}

// CFNodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CFNodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFNodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CFNodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFNodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CFNodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFNodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CFNodesSession struct {
	Contract     *CFNodes          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CFNodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CFNodesCallerSession struct {
	Contract *CFNodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CFNodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CFNodesTransactorSession struct {
	Contract     *CFNodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CFNodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CFNodesRaw struct {
	Contract *CFNodes // Generic contract binding to access the raw methods on
}

// CFNodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CFNodesCallerRaw struct {
	Contract *CFNodesCaller // Generic read-only contract binding to access the raw methods on
}

// CFNodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CFNodesTransactorRaw struct {
	Contract *CFNodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCFNodes creates a new instance of CFNodes, bound to a specific deployed contract.
func NewCFNodes(address common.Address, backend bind.ContractBackend) (*CFNodes, error) {
	contract, err := bindCFNodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CFNodes{CFNodesCaller: CFNodesCaller{contract: contract}, CFNodesTransactor: CFNodesTransactor{contract: contract}, CFNodesFilterer: CFNodesFilterer{contract: contract}}, nil
}

// NewCFNodesCaller creates a new read-only instance of CFNodes, bound to a specific deployed contract.
func NewCFNodesCaller(address common.Address, caller bind.ContractCaller) (*CFNodesCaller, error) {
	contract, err := bindCFNodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CFNodesCaller{contract: contract}, nil
}

// NewCFNodesTransactor creates a new write-only instance of CFNodes, bound to a specific deployed contract.
func NewCFNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*CFNodesTransactor, error) {
	contract, err := bindCFNodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CFNodesTransactor{contract: contract}, nil
}

// NewCFNodesFilterer creates a new log filterer instance of CFNodes, bound to a specific deployed contract.
func NewCFNodesFilterer(address common.Address, filterer bind.ContractFilterer) (*CFNodesFilterer, error) {
	contract, err := bindCFNodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CFNodesFilterer{contract: contract}, nil
}

// bindCFNodes binds a generic wrapper to an already deployed contract.
func bindCFNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CFNodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CFNodes *CFNodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CFNodes.Contract.CFNodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CFNodes *CFNodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFNodes.Contract.CFNodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CFNodes *CFNodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CFNodes.Contract.CFNodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CFNodes *CFNodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CFNodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CFNodes *CFNodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFNodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CFNodes *CFNodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CFNodes.Contract.contract.Transact(opts, method, params...)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(totalNodes uint256)
func (_CFNodes *CFNodesCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "getNodeCount")
	return *ret0, err
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(totalNodes uint256)
func (_CFNodes *CFNodesSession) GetNodeCount() (*big.Int, error) {
	return _CFNodes.Contract.GetNodeCount(&_CFNodes.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(totalNodes uint256)
func (_CFNodes *CFNodesCallerSession) GetNodeCount() (*big.Int, error) {
	return _CFNodes.Contract.GetNodeCount(&_CFNodes.CallOpts)
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(_nodeName string) constant returns(publicKey string)
func (_CFNodes *CFNodesCaller) GetNodeKey(opts *bind.CallOpts, _nodeName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "getNodeKey", _nodeName)
	return *ret0, err
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(_nodeName string) constant returns(publicKey string)
func (_CFNodes *CFNodesSession) GetNodeKey(_nodeName string) (string, error) {
	return _CFNodes.Contract.GetNodeKey(&_CFNodes.CallOpts, _nodeName)
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(_nodeName string) constant returns(publicKey string)
func (_CFNodes *CFNodesCallerSession) GetNodeKey(_nodeName string) (string, error) {
	return _CFNodes.Contract.GetNodeKey(&_CFNodes.CallOpts, _nodeName)
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(_nodeName string) constant returns(existed bool)
func (_CFNodes *CFNodesCaller) IsNodeExisted(opts *bind.CallOpts, _nodeName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "isNodeExisted", _nodeName)
	return *ret0, err
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(_nodeName string) constant returns(existed bool)
func (_CFNodes *CFNodesSession) IsNodeExisted(_nodeName string) (bool, error) {
	return _CFNodes.Contract.IsNodeExisted(&_CFNodes.CallOpts, _nodeName)
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(_nodeName string) constant returns(existed bool)
func (_CFNodes *CFNodesCallerSession) IsNodeExisted(_nodeName string) (bool, error) {
	return _CFNodes.Contract.IsNodeExisted(&_CFNodes.CallOpts, _nodeName)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(_nodeName string, _publicKey string, _otherNodeInfo string) returns()
func (_CFNodes *CFNodesTransactor) AddNode(opts *bind.TransactOpts, _nodeName string, _publicKey string, _otherNodeInfo string) (*types.Transaction, error) {
	return _CFNodes.contract.Transact(opts, "addNode", _nodeName, _publicKey, _otherNodeInfo)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(_nodeName string, _publicKey string, _otherNodeInfo string) returns()
func (_CFNodes *CFNodesSession) AddNode(_nodeName string, _publicKey string, _otherNodeInfo string) (*types.Transaction, error) {
	return _CFNodes.Contract.AddNode(&_CFNodes.TransactOpts, _nodeName, _publicKey, _otherNodeInfo)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(_nodeName string, _publicKey string, _otherNodeInfo string) returns()
func (_CFNodes *CFNodesTransactorSession) AddNode(_nodeName string, _publicKey string, _otherNodeInfo string) (*types.Transaction, error) {
	return _CFNodes.Contract.AddNode(&_CFNodes.TransactOpts, _nodeName, _publicKey, _otherNodeInfo)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(_nodeName string) returns()
func (_CFNodes *CFNodesTransactor) DeleteNode(opts *bind.TransactOpts, _nodeName string) (*types.Transaction, error) {
	return _CFNodes.contract.Transact(opts, "deleteNode", _nodeName)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(_nodeName string) returns()
func (_CFNodes *CFNodesSession) DeleteNode(_nodeName string) (*types.Transaction, error) {
	return _CFNodes.Contract.DeleteNode(&_CFNodes.TransactOpts, _nodeName)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(_nodeName string) returns()
func (_CFNodes *CFNodesTransactorSession) DeleteNode(_nodeName string) (*types.Transaction, error) {
	return _CFNodes.Contract.DeleteNode(&_CFNodes.TransactOpts, _nodeName)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(_nodeName string, _publicKey string) returns()
func (_CFNodes *CFNodesTransactor) UpdateNode(opts *bind.TransactOpts, _nodeName string, _publicKey string) (*types.Transaction, error) {
	return _CFNodes.contract.Transact(opts, "updateNode", _nodeName, _publicKey)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(_nodeName string, _publicKey string) returns()
func (_CFNodes *CFNodesSession) UpdateNode(_nodeName string, _publicKey string) (*types.Transaction, error) {
	return _CFNodes.Contract.UpdateNode(&_CFNodes.TransactOpts, _nodeName, _publicKey)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(_nodeName string, _publicKey string) returns()
func (_CFNodes *CFNodesTransactorSession) UpdateNode(_nodeName string, _publicKey string) (*types.Transaction, error) {
	return _CFNodes.Contract.UpdateNode(&_CFNodes.TransactOpts, _nodeName, _publicKey)
}

// CFNodesAddNodeIterator is returned from FilterAddNode and is used to iterate over the raw logs and unpacked data for AddNode events raised by the CFNodes contract.
type CFNodesAddNodeIterator struct {
	Event *CFNodesAddNode // Event containing the contract specifics and raw log

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
func (it *CFNodesAddNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFNodesAddNode)
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
		it.Event = new(CFNodesAddNode)
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
func (it *CFNodesAddNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFNodesAddNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFNodesAddNode represents a AddNode event raised by the CFNodes contract.
type CFNodesAddNode struct {
	NodeName      string
	PublicKey     string
	OtherNodeInfo string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddNode is a free log retrieval operation binding the contract event 0x92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd.
//
// Solidity: event AddNode(nodeName string, publicKey string, otherNodeInfo string)
func (_CFNodes *CFNodesFilterer) FilterAddNode(opts *bind.FilterOpts) (*CFNodesAddNodeIterator, error) {

	logs, sub, err := _CFNodes.contract.FilterLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return &CFNodesAddNodeIterator{contract: _CFNodes.contract, event: "AddNode", logs: logs, sub: sub}, nil
}

// WatchAddNode is a free log subscription operation binding the contract event 0x92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd.
//
// Solidity: event AddNode(nodeName string, publicKey string, otherNodeInfo string)
func (_CFNodes *CFNodesFilterer) WatchAddNode(opts *bind.WatchOpts, sink chan<- *CFNodesAddNode) (event.Subscription, error) {

	logs, sub, err := _CFNodes.contract.WatchLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFNodesAddNode)
				if err := _CFNodes.contract.UnpackLog(event, "AddNode", log); err != nil {
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

// CFNodesDeleteNodeIterator is returned from FilterDeleteNode and is used to iterate over the raw logs and unpacked data for DeleteNode events raised by the CFNodes contract.
type CFNodesDeleteNodeIterator struct {
	Event *CFNodesDeleteNode // Event containing the contract specifics and raw log

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
func (it *CFNodesDeleteNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFNodesDeleteNode)
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
		it.Event = new(CFNodesDeleteNode)
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
func (it *CFNodesDeleteNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFNodesDeleteNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFNodesDeleteNode represents a DeleteNode event raised by the CFNodes contract.
type CFNodesDeleteNode struct {
	NodeName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteNode is a free log retrieval operation binding the contract event 0xba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a.
//
// Solidity: event DeleteNode(nodeName string)
func (_CFNodes *CFNodesFilterer) FilterDeleteNode(opts *bind.FilterOpts) (*CFNodesDeleteNodeIterator, error) {

	logs, sub, err := _CFNodes.contract.FilterLogs(opts, "DeleteNode")
	if err != nil {
		return nil, err
	}
	return &CFNodesDeleteNodeIterator{contract: _CFNodes.contract, event: "DeleteNode", logs: logs, sub: sub}, nil
}

// WatchDeleteNode is a free log subscription operation binding the contract event 0xba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a.
//
// Solidity: event DeleteNode(nodeName string)
func (_CFNodes *CFNodesFilterer) WatchDeleteNode(opts *bind.WatchOpts, sink chan<- *CFNodesDeleteNode) (event.Subscription, error) {

	logs, sub, err := _CFNodes.contract.WatchLogs(opts, "DeleteNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFNodesDeleteNode)
				if err := _CFNodes.contract.UnpackLog(event, "DeleteNode", log); err != nil {
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

// CFNodesUpdateNodeIterator is returned from FilterUpdateNode and is used to iterate over the raw logs and unpacked data for UpdateNode events raised by the CFNodes contract.
type CFNodesUpdateNodeIterator struct {
	Event *CFNodesUpdateNode // Event containing the contract specifics and raw log

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
func (it *CFNodesUpdateNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFNodesUpdateNode)
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
		it.Event = new(CFNodesUpdateNode)
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
func (it *CFNodesUpdateNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFNodesUpdateNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFNodesUpdateNode represents a UpdateNode event raised by the CFNodes contract.
type CFNodesUpdateNode struct {
	NodeName  string
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateNode is a free log retrieval operation binding the contract event 0xddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a.
//
// Solidity: event UpdateNode(nodeName string, publicKey string)
func (_CFNodes *CFNodesFilterer) FilterUpdateNode(opts *bind.FilterOpts) (*CFNodesUpdateNodeIterator, error) {

	logs, sub, err := _CFNodes.contract.FilterLogs(opts, "UpdateNode")
	if err != nil {
		return nil, err
	}
	return &CFNodesUpdateNodeIterator{contract: _CFNodes.contract, event: "UpdateNode", logs: logs, sub: sub}, nil
}

// WatchUpdateNode is a free log subscription operation binding the contract event 0xddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a.
//
// Solidity: event UpdateNode(nodeName string, publicKey string)
func (_CFNodes *CFNodesFilterer) WatchUpdateNode(opts *bind.WatchOpts, sink chan<- *CFNodesUpdateNode) (event.Subscription, error) {

	logs, sub, err := _CFNodes.contract.WatchLogs(opts, "UpdateNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFNodesUpdateNode)
				if err := _CFNodes.contract.UnpackLog(event, "UpdateNode", log); err != nil {
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
