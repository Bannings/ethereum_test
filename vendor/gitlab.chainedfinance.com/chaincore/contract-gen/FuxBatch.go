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

// FuxBatchABI is the input ABI used to generate the binding from.
const FuxBatchABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TRANSFER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_BC_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_RBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_ISSUER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staticAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_COIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_BATCH\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isUsingProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_SPLIT\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_PAYBOX_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_ERC721_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_len\",\"type\":\"uint256\"}],\"name\":\"setMaxLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokens\",\"type\":\"uint256[]\"}],\"name\":\"addJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"runJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isCompile\",\"outputs\":[{\"name\":\"compile\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxBatchBin is the compiled bytecode used for deploying new contracts.
const FuxBatchBin = `0x608060405260008060146101000a81548160ff0219169083151502179055503480156200002b57600080fd5b5060405160208062001acb833981018060405281019080805190602001909291905050506200009e6040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000c7640100000000026401000000009004565b620000b881620000e3640100000000026401000000009004565b601e6003819055505062000213565b8060029080519060200190620000df92919062000164565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200012057600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a757805160ff1916838001178555620001d8565b82800160010185558215620001d8579182015b82811115620001d7578251825591602001919060010190620001ba565b5b509050620001e79190620001eb565b5090565b6200021091905b808211156200020c576000816000905550600101620001f2565b5090565b90565b6118a880620002236000396000f300608060405260043610610133576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063071db5611461013857806309c4b0fa146101c85780630c7bba78146102585780630daaf6e4146102e85780632bc3a44c1461037857806332061899146103a757806333e653e8146103fe57806343a73d9a1461048e5780634aa03a35146104e5578063510c1b231461057557806375ae7a8d14610605578063781bed441461065c57806388f7748d1461068b5780639a1e48221461071b578063a3c332ab14610748578063abf836fc1461078d578063c0d193251461083d578063d06a89a4146108cd578063d6b419fb146108f8578063dc2f786714610988578063e941ed61146109b5578063f43a6d2414610a45578063f5ac8fb914610ad5575b600080fd5b34801561014457600080fd5b5061014d610b65565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561018d578082015181840152602081019050610172565b50505050905090810190601f1680156101ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101d457600080fd5b506101dd610b9e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561021d578082015181840152602081019050610202565b50505050905090810190601f16801561024a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026457600080fd5b5061026d610bd7565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102ad578082015181840152602081019050610292565b50505050905090810190601f1680156102da5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102f457600080fd5b506102fd610c10565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561033d578082015181840152602081019050610322565b50505050905090810190601f16801561036a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038457600080fd5b5061038d610c49565b604051808215151515815260200191505060405180910390f35b3480156103b357600080fd5b506103bc610c5c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561040a57600080fd5b50610413610c82565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610453578082015181840152602081019050610438565b50505050905090810190601f1680156104805780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561049a57600080fd5b506104a3610cbb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f157600080fd5b506104fa610ce5565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561053a57808201518184015260208101905061051f565b50505050905090810190601f1680156105675780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058157600080fd5b5061058a610d1e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105ca5780820151818401526020810190506105af565b50505050905090810190601f1680156105f75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561061157600080fd5b5061061a610d57565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561066857600080fd5b50610671610d7c565b604051808215151515815260200191505060405180910390f35b34801561069757600080fd5b506106a0610d93565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106e05780820151818401526020810190506106c5565b50505050905090810190601f16801561070d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561072757600080fd5b5061074660048036038101908080359060200190929190505050610dcc565b005b34801561075457600080fd5b5061077360048036038101908080359060200190929190505050610ff5565b604051808215151515815260200191505060405180910390f35b34801561079957600080fd5b5061083b60048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061101f565b005b34801561084957600080fd5b50610852611219565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610892578082015181840152602081019050610877565b50505050905090810190601f1680156108bf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156108d957600080fd5b506108e2611252565b6040518082815260200191505060405180910390f35b34801561090457600080fd5b5061090d611258565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561094d578082015181840152602081019050610932565b50505050905090810190601f16801561097a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561099457600080fd5b506109b360048036038101908080359060200190929190505050611291565b005b3480156109c157600080fd5b506109ca6112af565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a0a5780820151818401526020810190506109ef565b50505050905090810190601f168015610a375780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610a5157600080fd5b50610a5a6112e8565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a9a578082015181840152602081019050610a7f565b50505050905090810190601f168015610ac75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610ae157600080fd5b50610aea611321565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b2a578082015181840152602081019050610b0f565b50505050905090810190601f168015610b575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805190810160405280601381526020017f5472616e73666572526f6c65576974684675780000000000000000000000000081525081565b6040805190810160405280600a81526020017f667578426347726f75700000000000000000000000000000000000000000000081525081565b6040805190810160405280600b81526020017f524241435769746846757800000000000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f667578497373756572000000000000000000000000000000000000000000000081525081565b600060149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600781526020017f467578436f696e0000000000000000000000000000000000000000000000000081525081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525081565b6040805190810160405280600881526020017f467578426174636800000000000000000000000000000000000000000000000081525081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615905090565b6040805190810160405280600881526020017f46757853706c697400000000000000000000000000000000000000000000000081525081565b6000806000610dda3361135a565b80610dea5750610de933611440565b5b1515610df557600080fd5b610dfe84610ff5565b15610e0857610fef565b6005600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692506006600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915060016007600086815260200190815260200160002060006101000a81548160ff021916908315150217905550600090505b6004600085815260200190815260200160002080549050811015610fee57610ecb611526565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc884846004600089815260200190815260200160002085815481101515610f0857fe5b90600052602060002001546040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610fcb57600080fd5b505af1158015610fdf573d6000803e3d6000fd5b50505050806001019050610ea5565b5b50505050565b60006007600083815260200190815260200160002060009054906101000a900460ff169050919050565b6110283361135a565b80611038575061103733611440565b5b151561104357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561107f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156110bb57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166005600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561112957600080fd5b600354815110801561113c575060008151115b151561114757600080fd5b8060046000868152602001908152602001600020908051906020019061116e92919061180a565b50826005600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816006600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6040805190810160405280600a81526020017f46757853746f726167650000000000000000000000000000000000000000000081525081565b60035481565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b61129a33611440565b15156112a557600080fd5b8060038190555050565b6040805190810160405280601081526020017f467578506179426f78466163746f72790000000000000000000000000000000081525081565b6040805190810160405280600e81526020017f467578455243373231546f6b656e00000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f467578436f6e666967000000000000000000000000000000000000000000000081525081565b600061136461156b565b73ffffffffffffffffffffffffffffffffffffffff1663abd7712b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156113fe57600080fd5b505af1158015611412573d6000803e3d6000fd5b505050506040513d602081101561142857600080fd5b81019080805190602001909291905050509050919050565b600061144a61156b565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156114e457600080fd5b505af11580156114f8573d6000803e3d6000fd5b505050506040513d602081101561150e57600080fd5b81019080805190602001909291905050509050919050565b60006115666040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525061157a565b905090565b60006115756116b4565b905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561162657808201518184015260208101905061160b565b50505050905090810190601f1680156116535780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561167257600080fd5b505af1158015611686573d6000803e3d6000fd5b505050506040513d602081101561169c57600080fd5b81019080805190602001909291905050509050919050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560026040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156117aa5780601f1061177f576101008083540402835291602001916117aa565b820191906000526020600020905b81548152906001019060200180831161178d57829003601f168201915b505092505050602060405180830381600087803b1580156117ca57600080fd5b505af11580156117de573d6000803e3d6000fd5b505050506040513d60208110156117f457600080fd5b8101908080519060200190929190505050905090565b828054828255906000526020600020908101928215611846579160200282015b8281111561184557825182559160200191906001019061182a565b5b5090506118539190611857565b5090565b61187991905b8082111561187557600081600090555060010161185d565b5090565b905600a165627a7a723058208fde18b4ef84abfda436453bb377d1fe22ac40a9a947fd7a8beb352aa407f1ab0029`

// DeployFuxBatch deploys a new Ethereum contract, binding an instance of FuxBatch to it.
func DeployFuxBatch(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *FuxBatch, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxBatchABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxBatchBin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}, FuxBatchFilterer: FuxBatchFilterer{contract: contract}}, nil
}

// FuxBatch is an auto generated Go binding around an Ethereum contract.
type FuxBatch struct {
	FuxBatchCaller     // Read-only binding to the contract
	FuxBatchTransactor // Write-only binding to the contract
	FuxBatchFilterer   // Log filterer for contract events
}

// FuxBatchCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxBatchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxBatchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxBatchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxBatchSession struct {
	Contract     *FuxBatch         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxBatchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxBatchCallerSession struct {
	Contract *FuxBatchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxBatchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxBatchTransactorSession struct {
	Contract     *FuxBatchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxBatchRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxBatchRaw struct {
	Contract *FuxBatch // Generic contract binding to access the raw methods on
}

// FuxBatchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxBatchCallerRaw struct {
	Contract *FuxBatchCaller // Generic read-only contract binding to access the raw methods on
}

// FuxBatchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxBatchTransactorRaw struct {
	Contract *FuxBatchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxBatch creates a new instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatch(address common.Address, backend bind.ContractBackend) (*FuxBatch, error) {
	contract, err := bindFuxBatch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}, FuxBatchFilterer: FuxBatchFilterer{contract: contract}}, nil
}

// NewFuxBatchCaller creates a new read-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchCaller(address common.Address, caller bind.ContractCaller) (*FuxBatchCaller, error) {
	contract, err := bindFuxBatch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchCaller{contract: contract}, nil
}

// NewFuxBatchTransactor creates a new write-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxBatchTransactor, error) {
	contract, err := bindFuxBatch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchTransactor{contract: contract}, nil
}

// NewFuxBatchFilterer creates a new log filterer instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxBatchFilterer, error) {
	contract, err := bindFuxBatch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxBatchFilterer{contract: contract}, nil
}

// bindFuxBatch binds a generic wrapper to an already deployed contract.
func bindFuxBatch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxBatchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxBatch *FuxBatchRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxBatch.Contract.FuxBatchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxBatch *FuxBatchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxBatch.Contract.FuxBatchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxBatch *FuxBatchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxBatch.Contract.FuxBatchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxBatch *FuxBatchCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxBatch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxBatch *FuxBatchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxBatch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxBatch *FuxBatchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxBatch.Contract.contract.Transact(opts, method, params...)
}

// ROLEFUXBCGROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXBCGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_BC_GROUP")
	return *ret0, err
}

// ROLEFUXBCGROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXBCGROUP() (string, error) {
	return _FuxBatch.Contract.ROLEFUXBCGROUP(&_FuxBatch.CallOpts)
}

// ROLEFUXBCGROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXBCGROUP() (string, error) {
	return _FuxBatch.Contract.ROLEFUXBCGROUP(&_FuxBatch.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXHUB() (string, error) {
	return _FuxBatch.Contract.ROLEFUXHUB(&_FuxBatch.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxBatch.Contract.ROLEFUXHUB(&_FuxBatch.CallOpts)
}

// ROLEFUXISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXISSUER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_ISSUER")
	return *ret0, err
}

// ROLEFUXISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXISSUER() (string, error) {
	return _FuxBatch.Contract.ROLEFUXISSUER(&_FuxBatch.CallOpts)
}

// ROLEFUXISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXISSUER() (string, error) {
	return _FuxBatch.Contract.ROLEFUXISSUER(&_FuxBatch.CallOpts)
}

// ROUTEFUXBATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXBATCH(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_BATCH")
	return *ret0, err
}

// ROUTEFUXBATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXBATCH() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXBATCH(&_FuxBatch.CallOpts)
}

// ROUTEFUXBATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXBATCH() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXBATCH(&_FuxBatch.CallOpts)
}

// ROUTEFUXCOIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXCOIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_COIN")
	return *ret0, err
}

// ROUTEFUXCOIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXCOIN() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXCOIN(&_FuxBatch.CallOpts)
}

// ROUTEFUXCOIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXCOIN() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXCOIN(&_FuxBatch.CallOpts)
}

// ROUTEFUXCONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXCONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_CONFIG")
	return *ret0, err
}

// ROUTEFUXCONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXCONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXCONFIG(&_FuxBatch.CallOpts)
}

// ROUTEFUXCONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXCONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXCONFIG(&_FuxBatch.CallOpts)
}

// ROUTEFUXERC721TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXERC721TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_ERC721_TOKEN")
	return *ret0, err
}

// ROUTEFUXERC721TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXERC721TOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXERC721TOKEN(&_FuxBatch.CallOpts)
}

// ROUTEFUXERC721TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXERC721TOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXERC721TOKEN(&_FuxBatch.CallOpts)
}

// ROUTEFUXPAYBOXFACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXPAYBOXFACTORY(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_PAYBOX_FACTORY")
	return *ret0, err
}

// ROUTEFUXPAYBOXFACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXPAYBOXFACTORY() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXPAYBOXFACTORY(&_FuxBatch.CallOpts)
}

// ROUTEFUXPAYBOXFACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXPAYBOXFACTORY() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXPAYBOXFACTORY(&_FuxBatch.CallOpts)
}

// ROUTEFUXRBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_RBAC")
	return *ret0, err
}

// ROUTEFUXRBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXRBAC() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXRBAC(&_FuxBatch.CallOpts)
}

// ROUTEFUXRBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXRBAC() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXRBAC(&_FuxBatch.CallOpts)
}

// ROUTEFUXSPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXSPLIT(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_SPLIT")
	return *ret0, err
}

// ROUTEFUXSPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXSPLIT() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXSPLIT(&_FuxBatch.CallOpts)
}

// ROUTEFUXSPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXSPLIT() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXSPLIT(&_FuxBatch.CallOpts)
}

// ROUTEFUXSTORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXSTORAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_STORAGE")
	return *ret0, err
}

// ROUTEFUXSTORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXSTORAGE() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXSTORAGE(&_FuxBatch.CallOpts)
}

// ROUTEFUXSTORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXSTORAGE() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXSTORAGE(&_FuxBatch.CallOpts)
}

// ROUTEFUXTOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXTOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_TOKEN")
	return *ret0, err
}

// ROUTEFUXTOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXTOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXTOKEN(&_FuxBatch.CallOpts)
}

// ROUTEFUXTOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXTOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXTOKEN(&_FuxBatch.CallOpts)
}

// ROUTEFUXTRANSFERROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTEFUXTRANSFERROLE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_TRANSFER_ROLE")
	return *ret0, err
}

// ROUTEFUXTRANSFERROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTEFUXTRANSFERROLE() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXTRANSFERROLE(&_FuxBatch.CallOpts)
}

// ROUTEFUXTRANSFERROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTEFUXTRANSFERROLE() (string, error) {
	return _FuxBatch.Contract.ROUTEFUXTRANSFERROLE(&_FuxBatch.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxBatch *FuxBatchCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "getProxyAddress")
	return *ret0, err
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxBatch *FuxBatchSession) GetProxyAddress() (common.Address, error) {
	return _FuxBatch.Contract.GetProxyAddress(&_FuxBatch.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) GetProxyAddress() (common.Address, error) {
	return _FuxBatch.Contract.GetProxyAddress(&_FuxBatch.CallOpts)
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchCaller) IsCompile(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isCompile", _id)
	return *ret0, err
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchSession) IsCompile(_id *big.Int) (bool, error) {
	return _FuxBatch.Contract.IsCompile(&_FuxBatch.CallOpts, _id)
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchCallerSession) IsCompile(_id *big.Int) (bool, error) {
	return _FuxBatch.Contract.IsCompile(&_FuxBatch.CallOpts, _id)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsUsingProxy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isUsingProxy")
	return *ret0, err
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsUsingProxy() (bool, error) {
	return _FuxBatch.Contract.IsUsingProxy(&_FuxBatch.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsUsingProxy() (bool, error) {
	return _FuxBatch.Contract.IsUsingProxy(&_FuxBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchCaller) MaxLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "maxLength")
	return *ret0, err
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchSession) MaxLength() (*big.Int, error) {
	return _FuxBatch.Contract.MaxLength(&_FuxBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchCallerSession) MaxLength() (*big.Int, error) {
	return _FuxBatch.Contract.MaxLength(&_FuxBatch.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxBatch *FuxBatchCaller) SAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "sAddr")
	return *ret0, err
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxBatch *FuxBatchSession) SAddr() (common.Address, error) {
	return _FuxBatch.Contract.SAddr(&_FuxBatch.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) SAddr() (common.Address, error) {
	return _FuxBatch.Contract.SAddr(&_FuxBatch.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxBatch *FuxBatchCaller) StaticAddr(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "staticAddr")
	return *ret0, err
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxBatch *FuxBatchSession) StaticAddr() (bool, error) {
	return _FuxBatch.Contract.StaticAddr(&_FuxBatch.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) StaticAddr() (bool, error) {
	return _FuxBatch.Contract.StaticAddr(&_FuxBatch.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxBatch *FuxBatchCaller) SysProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "sysProxy")
	return *ret0, err
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxBatch *FuxBatchSession) SysProxy() (common.Address, error) {
	return _FuxBatch.Contract.SysProxy(&_FuxBatch.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) SysProxy() (common.Address, error) {
	return _FuxBatch.Contract.SysProxy(&_FuxBatch.CallOpts)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactor) AddJob(opts *bind.TransactOpts, _id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "addJob", _id, _from, _to, _tokens)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchSession) AddJob(_id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.AddJob(&_FuxBatch.TransactOpts, _id, _from, _to, _tokens)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactorSession) AddJob(_id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.AddJob(&_FuxBatch.TransactOpts, _id, _from, _to, _tokens)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchTransactor) RunJob(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "runJob", _id)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchSession) RunJob(_id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.RunJob(&_FuxBatch.TransactOpts, _id)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) RunJob(_id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.RunJob(&_FuxBatch.TransactOpts, _id)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchTransactor) SetMaxLength(opts *bind.TransactOpts, _len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "setMaxLength", _len)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchSession) SetMaxLength(_len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetMaxLength(&_FuxBatch.TransactOpts, _len)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) SetMaxLength(_len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetMaxLength(&_FuxBatch.TransactOpts, _len)
}
