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
const FuxBatchBin = `0x606060405260008060146101000a81548160ff02191690831515021790555034156200002a57600080fd5b60405160208062001ab0833981016040528080519060200190919050506200009b6040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000c96401000000000262001745176401000000009004565b620000ba81620000e5640100000000026200175f176401000000009004565b601e6003819055505062000215565b8060029080519060200190620000e192919062000166565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200012257600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a957805160ff1916838001178555620001da565b82800160010185558215620001da579182015b82811115620001d9578251825591602001919060010190620001bc565b5b509050620001e99190620001ed565b5090565b6200021291905b808211156200020e576000816000905550600101620001f4565b5090565b90565b61188b80620002256000396000f300606060405260043610610133576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063071db5611461013857806309c4b0fa146101c65780630c7bba78146102545780630daaf6e4146102e25780632bc3a44c14610370578063320618991461039d57806333e653e8146103f257806343a73d9a146104805780634aa03a35146104d5578063510c1b231461056357806375ae7a8d146105f1578063781bed441461064657806388f7748d146106735780639a1e482214610701578063a3c332ab14610724578063abf836fc1461075f578063c0d1932514610800578063d06a89a41461088e578063d6b419fb146108b7578063dc2f786714610945578063e941ed6114610968578063f43a6d24146109f6578063f5ac8fb914610a84575b600080fd5b341561014357600080fd5b61014b610b12565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561018b578082015181840152602081019050610170565b50505050905090810190601f1680156101b85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101d157600080fd5b6101d9610b4b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102195780820151818401526020810190506101fe565b50505050905090810190601f1680156102465780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561025f57600080fd5b610267610b84565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102a757808201518184015260208101905061028c565b50505050905090810190601f1680156102d45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102ed57600080fd5b6102f5610bbd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561033557808201518184015260208101905061031a565b50505050905090810190601f1680156103625780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561037b57600080fd5b610383610bf6565b604051808215151515815260200191505060405180910390f35b34156103a857600080fd5b6103b0610c09565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103fd57600080fd5b610405610c2f565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561044557808201518184015260208101905061042a565b50505050905090810190601f1680156104725780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561048b57600080fd5b610493610c68565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104e057600080fd5b6104e8610c92565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561052857808201518184015260208101905061050d565b50505050905090810190601f1680156105555780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561056e57600080fd5b610576610ccb565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105b657808201518184015260208101905061059b565b50505050905090810190601f1680156105e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156105fc57600080fd5b610604610d04565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561065157600080fd5b610659610d29565b604051808215151515815260200191505060405180910390f35b341561067e57600080fd5b610686610d40565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106c65780820151818401526020810190506106ab565b50505050905090810190601f1680156106f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561070c57600080fd5b6107226004808035906020019091905050610d79565b005b341561072f57600080fd5b6107456004808035906020019091905050610f9a565b604051808215151515815260200191505060405180910390f35b341561076a57600080fd5b6107fe600480803590602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091905050610fc4565b005b341561080b57600080fd5b610813611138565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610853578082015181840152602081019050610838565b50505050905090810190601f1680156108805780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561089957600080fd5b6108a1611171565b6040518082815260200191505060405180910390f35b34156108c257600080fd5b6108ca611177565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561090a5780820151818401526020810190506108ef565b50505050905090810190601f1680156109375780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561095057600080fd5b61096660048080359060200190919050506111b0565b005b341561097357600080fd5b61097b6111ce565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109bb5780820151818401526020810190506109a0565b50505050905090810190601f1680156109e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3415610a0157600080fd5b610a09611207565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a49578082015181840152602081019050610a2e565b50505050905090810190601f168015610a765780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3415610a8f57600080fd5b610a97611240565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610ad7578082015181840152602081019050610abc565b50505050905090810190601f168015610b045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805190810160405280601381526020017f5472616e73666572526f6c65576974684675780000000000000000000000000081525081565b6040805190810160405280600a81526020017f667578426347726f75700000000000000000000000000000000000000000000081525081565b6040805190810160405280600b81526020017f524241435769746846757800000000000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f667578497373756572000000000000000000000000000000000000000000000081525081565b600060149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600781526020017f467578436f696e0000000000000000000000000000000000000000000000000081525081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525081565b6040805190810160405280600881526020017f467578426174636800000000000000000000000000000000000000000000000081525081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615905090565b6040805190810160405280600881526020017f46757853706c697400000000000000000000000000000000000000000000000081525081565b6000806000610d8733611279565b80610d975750610d963361133a565b5b1515610da257600080fd5b610dab84610f9a565b15610db557610f94565b6005600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692506006600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915060016007600086815260200190815260200160002060006101000a81548160ff021916908315150217905550600090505b6004600085815260200190815260200160002080549050811015610f9357610e786113fb565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc884846004600089815260200190815260200160002085815481101515610eb557fe5b9060005260206000209001546040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1515610f7857600080fd5b5af11515610f8557600080fd5b505050806001019050610e52565b5b50505050565b60006007600083815260200190815260200160002060009054906101000a900460ff169050919050565b610fcd33611279565b80610fdd5750610fdc3361133a565b5b1515610fe857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166005600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561105657600080fd5b601e815110151561106657600080fd5b8060046000868152602001908152602001600020908051906020019061108d9291906116d3565b50826005600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816006600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6040805190810160405280600a81526020017f46757853746f726167650000000000000000000000000000000000000000000081525081565b60035481565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6111b93361133a565b15156111c457600080fd5b8060038190555050565b6040805190810160405280601081526020017f467578506179426f78466163746f72790000000000000000000000000000000081525081565b6040805190810160405280600e81526020017f467578455243373231546f6b656e00000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f467578436f6e666967000000000000000000000000000000000000000000000081525081565b6000611283611440565b73ffffffffffffffffffffffffffffffffffffffff1663abd7712b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561131c57600080fd5b5af1151561132957600080fd5b505050604051805190509050919050565b6000611344611440565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15156113dd57600080fd5b5af115156113ea57600080fd5b505050604051805190509050919050565b600061143b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525061144f565b905090565b600061144a611564565b905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156114fb5780820151818401526020810190506114e0565b50505050905090810190601f1680156115285780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b151561154657600080fd5b5af1151561155357600080fd5b505050604051805190509050919050565b60008060149054906101000a900460ff16156115a3576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506116d0565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560026040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156116975780601f1061166c57610100808354040283529160200191611697565b820191906000526020600020905b81548152906001019060200180831161167a57829003601f168201915b505092505050602060405180830381600087803b15156116b657600080fd5b5af115156116c357600080fd5b5050506040518051905090505b90565b82805482825590600052602060002090810192821561170f579160200282015b8281111561170e5782518255916020019190600101906116f3565b5b50905061171c9190611720565b5090565b61174291905b8082111561173e576000816000905550600101611726565b5090565b90565b806002908051906020019061175b9291906117df565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561179b57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061182057805160ff191683800117855561184e565b8280016001018555821561184e579182015b8281111561184d578251825591602001919060010190611832565b5b50905061185b9190611720565b50905600a165627a7a723058202cfa9131ed6a212973c32fd97b44031b1eae56c5e1f320067097d51d2574a0e90029`

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
	return address, tx, &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}}, nil
}

// FuxBatch is an auto generated Go binding around an Ethereum contract.
type FuxBatch struct {
	FuxBatchCaller     // Read-only binding to the contract
	FuxBatchTransactor // Write-only binding to the contract
}

// FuxBatchCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxBatchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxBatchTransactor struct {
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
	contract, err := bindFuxBatch(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}}, nil
}

// NewFuxBatchCaller creates a new read-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchCaller(address common.Address, caller bind.ContractCaller) (*FuxBatchCaller, error) {
	contract, err := bindFuxBatch(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchCaller{contract: contract}, nil
}

// NewFuxBatchTransactor creates a new write-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxBatchTransactor, error) {
	contract, err := bindFuxBatch(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FuxBatchTransactor{contract: contract}, nil
}

// bindFuxBatch binds a generic wrapper to an already deployed contract.
func bindFuxBatch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxBatchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLE_FUX_BC_GROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_BC_GROUP")
	return *ret0, err
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLE_FUX_BC_GROUP() (string, error) {
	return _FuxBatch.Contract.ROLE_FUX_BC_GROUP(&_FuxBatch.CallOpts)
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLE_FUX_BC_GROUP() (string, error) {
	return _FuxBatch.Contract.ROLE_FUX_BC_GROUP(&_FuxBatch.CallOpts)
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLE_FUX_HUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLE_FUX_HUB() (string, error) {
	return _FuxBatch.Contract.ROLE_FUX_HUB(&_FuxBatch.CallOpts)
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLE_FUX_HUB() (string, error) {
	return _FuxBatch.Contract.ROLE_FUX_HUB(&_FuxBatch.CallOpts)
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLE_FUX_ISSUER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_ISSUER")
	return *ret0, err
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLE_FUX_ISSUER() (string, error) {
	return _FuxBatch.Contract.ROLE_FUX_ISSUER(&_FuxBatch.CallOpts)
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLE_FUX_ISSUER() (string, error) {
	return _FuxBatch.Contract.ROLE_FUX_ISSUER(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_BATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_BATCH(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_BATCH")
	return *ret0, err
}

// ROUTE_FUX_BATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_BATCH() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_BATCH(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_BATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_BATCH() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_BATCH(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_COIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_COIN")
	return *ret0, err
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_COIN() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_COIN(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_COIN() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_COIN(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_CONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_CONFIG")
	return *ret0, err
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_CONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_CONFIG(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_CONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_CONFIG(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_ERC721_TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_ERC721_TOKEN")
	return *ret0, err
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_ERC721_TOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_ERC721_TOKEN(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_ERC721_TOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_ERC721_TOKEN(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_PAYBOX_FACTORY(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_PAYBOX_FACTORY")
	return *ret0, err
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_PAYBOX_FACTORY() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_PAYBOX_FACTORY(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_PAYBOX_FACTORY() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_PAYBOX_FACTORY(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_RBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_RBAC")
	return *ret0, err
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_RBAC() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_RBAC(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_RBAC() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_RBAC(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_SPLIT(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_SPLIT")
	return *ret0, err
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_SPLIT() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_SPLIT(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_SPLIT() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_SPLIT(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_STORAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_STORAGE")
	return *ret0, err
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_STORAGE() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_STORAGE(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_STORAGE() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_STORAGE(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_TOKEN")
	return *ret0, err
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_TOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_TOKEN(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_TOKEN() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_TOKEN(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTE_FUX_TRANSFER_ROLE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_FUX_TRANSFER_ROLE")
	return *ret0, err
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTE_FUX_TRANSFER_ROLE() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_TRANSFER_ROLE(&_FuxBatch.CallOpts)
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTE_FUX_TRANSFER_ROLE() (string, error) {
	return _FuxBatch.Contract.ROUTE_FUX_TRANSFER_ROLE(&_FuxBatch.CallOpts)
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
