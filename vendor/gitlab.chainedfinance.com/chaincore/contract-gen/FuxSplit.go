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

// FuxSplitABI is the input ABI used to generate the binding from.
const FuxSplitABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TRANSFER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_BC_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_RBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_ISSUER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staticAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_COIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addrFuxERC721\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isUsingProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_SPLIT\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addrFuxStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_PAYBOX_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_ERC721_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_newIds\",\"type\":\"uint256[2]\"},{\"name\":\"_newCoins\",\"type\":\"uint256[2]\"},{\"name\":\"_states\",\"type\":\"uint256[2]\"}],\"name\":\"splitFux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxSplitBin is the compiled bytecode used for deploying new contracts.
const FuxSplitBin = `0x606060405260008060146101000a81548160ff02191690831515021790555034156200002a57600080fd5b60405160208062001d4f833981016040528080519060200190919050506200009b6040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000c164010000000002620019c7176401000000009004565b620000ba81620000dd64010000000002620019e1176401000000009004565b506200020d565b8060029080519060200190620000d99291906200015e565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200011a57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a157805160ff1916838001178555620001d2565b82800160010185558215620001d2579182015b82811115620001d1578251825591602001919060010190620001b4565b5b509050620001e19190620001e5565b5090565b6200020a91905b8082111562000206576000816000905550600101620001ec565b5090565b90565b611b32806200021d6000396000f300606060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063071db5611461011757806309c4b0fa146101a55780630c7bba78146102335780630daaf6e4146102c15780632bc3a44c1461034f578063320618991461037c57806333e653e8146103d157806341a88a1b1461045f57806343a73d9a146104b45780634aa03a35146105095780635209e9841461059757806375ae7a8d1461063e578063781bed441461069357806388f7748d146106c057806395c8757e1461074e578063c0d19325146107a3578063d6b419fb14610831578063e941ed61146108bf578063f43a6d241461094d578063f5ac8fb9146109db575b600080fd5b341561012257600080fd5b61012a610a69565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561016a57808201518184015260208101905061014f565b50505050905090810190601f1680156101975780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101b057600080fd5b6101b8610aa2565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f85780820151818401526020810190506101dd565b50505050905090810190601f1680156102255780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561023e57600080fd5b610246610adb565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561028657808201518184015260208101905061026b565b50505050905090810190601f1680156102b35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102cc57600080fd5b6102d4610b14565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103145780820151818401526020810190506102f9565b50505050905090810190601f1680156103415780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561035a57600080fd5b610362610b4d565b604051808215151515815260200191505060405180910390f35b341561038757600080fd5b61038f610b60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103dc57600080fd5b6103e4610b86565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610424578082015181840152602081019050610409565b50505050905090810190601f1680156104515780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561046a57600080fd5b610472610bbf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104bf57600080fd5b6104c7610be5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561051457600080fd5b61051c610c0f565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561055c578082015181840152602081019050610541565b50505050905090810190601f1680156105895780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156105a257600080fd5b61063c6004808035906020019091908060400190600280602002604051908101604052809291908260026020028082843782019150505050509190806040019060028060200260405190810160405280929190826002602002808284378201915050505050919080604001906002806020026040519081016040528092919082600260200280828437820191505050505091905050610c48565b005b341561064957600080fd5b6106516114f1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561069e57600080fd5b6106a6611516565b604051808215151515815260200191505060405180910390f35b34156106cb57600080fd5b6106d361152d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107135780820151818401526020810190506106f8565b50505050905090810190601f1680156107405780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561075957600080fd5b610761611566565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156107ae57600080fd5b6107b661158c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107f65780820151818401526020810190506107db565b50505050905090810190601f1680156108235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561083c57600080fd5b6108446115c5565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610884578082015181840152602081019050610869565b50505050905090810190601f1680156108b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156108ca57600080fd5b6108d26115fe565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109125780820151818401526020810190506108f7565b50505050905090810190601f16801561093f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561095857600080fd5b610960611637565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109a0578082015181840152602081019050610985565b50505050905090810190601f1680156109cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156109e657600080fd5b6109ee611670565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a2e578082015181840152602081019050610a13565b50505050905090810190601f168015610a5b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805190810160405280601381526020017f5472616e73666572526f6c65576974684675780000000000000000000000000081525081565b6040805190810160405280600a81526020017f667578426347726f75700000000000000000000000000000000000000000000081525081565b6040805190810160405280600b81526020017f524241435769746846757800000000000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f667578497373756572000000000000000000000000000000000000000000000081525081565b600060149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600781526020017f467578436f696e0000000000000000000000000000000000000000000000000081525081565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525081565b600080600080600080610c5a336116a9565b80610c6a5750610c693361176a565b5b1515610c7557600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c25110256040805190810160405280600e81526020017f467578455243373231546f6b656e0000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d54578082015181840152602081019050610d39565b50505050905090810190601f168015610d815780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1515610d9f57600080fd5b5af11515610dac57600080fd5b505050604051805190509550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c25110256040805190810160405280600a81526020017f46757853746f72616765000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e97578082015181840152602081019050610e7c565b50505050905090810190601f168015610ec45780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1515610ee257600080fd5b5af11515610eef57600080fd5b5050506040518051905094508573ffffffffffffffffffffffffffffffffffffffff16636352211e8b6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1515610f6957600080fd5b5af11515610f7657600080fd5b5050506040518051905093508473ffffffffffffffffffffffffffffffffffffffff16631c61984b8b6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050606060405180830381600087803b1515610ff057600080fd5b5af11515610ffd57600080fd5b505050604051805190602001805190602001805190508094508195508293505050508061105989600160028110151561103257fe5b60200201518a600060028110151561104657fe5b602002015161182b90919063ffffffff16565b14151561106557600080fd5b8573ffffffffffffffffffffffffffffffffffffffff16639dc29fac858c6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b151561110757600080fd5b5af1151561111457600080fd5b5050508473ffffffffffffffffffffffffffffffffffffffff1663e234f61d8b6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b151561118557600080fd5b5af1151561119257600080fd5b5050508473ffffffffffffffffffffffffffffffffffffffff166386b27f4d8a60006002811015156111c057fe5b60200201518a60006002811015156111d457fe5b6020020151868b60006002811015156111e957fe5b60200201516040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808581526020018481526020018381526020018281526020018060200182810382526000815260200160200195505050505050600060405180830381600087803b151561126857600080fd5b5af1151561127557600080fd5b5050508473ffffffffffffffffffffffffffffffffffffffff166386b27f4d8a60016002811015156112a357fe5b60200201518a60016002811015156112b757fe5b6020020151868b60016002811015156112cc57fe5b60200201516040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808581526020018481526020018381526020018281526020018060200182810382526000815260200160200195505050505050600060405180830381600087803b151561134b57600080fd5b5af1151561135857600080fd5b5050508573ffffffffffffffffffffffffffffffffffffffff166340c10f19858b600060028110151561138757fe5b60200201516040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b151561141057600080fd5b5af1151561141d57600080fd5b5050508573ffffffffffffffffffffffffffffffffffffffff166340c10f19858b600160028110151561144c57fe5b60200201516040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15156114d557600080fd5b5af115156114e257600080fd5b50505050505050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615905090565b6040805190810160405280600881526020017f46757853706c697400000000000000000000000000000000000000000000000081525081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600a81526020017f46757853746f726167650000000000000000000000000000000000000000000081525081565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6040805190810160405280601081526020017f467578506179426f78466163746f72790000000000000000000000000000000081525081565b6040805190810160405280600e81526020017f467578455243373231546f6b656e00000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f467578436f6e666967000000000000000000000000000000000000000000000081525081565b60006116b3611849565b73ffffffffffffffffffffffffffffffffffffffff1663abd7712b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561174c57600080fd5b5af1151561175957600080fd5b505050604051805190509050919050565b6000611774611849565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561180d57600080fd5b5af1151561181a57600080fd5b505050604051805190509050919050565b600080828401905083811015151561183f57fe5b8091505092915050565b6000611853611858565b905090565b60008060149054906101000a900460ff1615611897576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506119c4565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560026040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561198b5780601f106119605761010080835404028352916020019161198b565b820191906000526020600020905b81548152906001019060200180831161196e57829003601f168201915b505092505050602060405180830381600087803b15156119aa57600080fd5b5af115156119b757600080fd5b5050506040518051905090505b90565b80600290805190602001906119dd929190611a61565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611a1d57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611aa257805160ff1916838001178555611ad0565b82800160010185558215611ad0579182015b82811115611acf578251825591602001919060010190611ab4565b5b509050611add9190611ae1565b5090565b611b0391905b80821115611aff576000816000905550600101611ae7565b5090565b905600a165627a7a72305820eb922d7b667411eaec03d979a2bc6301db9920419548ab68fe8b03d5aacd3ae60029`

// DeployFuxSplit deploys a new Ethereum contract, binding an instance of FuxSplit to it.
func DeployFuxSplit(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address) (common.Address, *types.Transaction, *FuxSplit, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxSplitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxSplitBin), backend, _addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxSplit{FuxSplitCaller: FuxSplitCaller{contract: contract}, FuxSplitTransactor: FuxSplitTransactor{contract: contract}}, nil
}

// FuxSplit is an auto generated Go binding around an Ethereum contract.
type FuxSplit struct {
	FuxSplitCaller     // Read-only binding to the contract
	FuxSplitTransactor // Write-only binding to the contract
}

// FuxSplitCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxSplitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxSplitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxSplitSession struct {
	Contract     *FuxSplit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxSplitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxSplitCallerSession struct {
	Contract *FuxSplitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxSplitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxSplitTransactorSession struct {
	Contract     *FuxSplitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxSplitRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxSplitRaw struct {
	Contract *FuxSplit // Generic contract binding to access the raw methods on
}

// FuxSplitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxSplitCallerRaw struct {
	Contract *FuxSplitCaller // Generic read-only contract binding to access the raw methods on
}

// FuxSplitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxSplitTransactorRaw struct {
	Contract *FuxSplitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxSplit creates a new instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplit(address common.Address, backend bind.ContractBackend) (*FuxSplit, error) {
	contract, err := bindFuxSplit(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxSplit{FuxSplitCaller: FuxSplitCaller{contract: contract}, FuxSplitTransactor: FuxSplitTransactor{contract: contract}}, nil
}

// NewFuxSplitCaller creates a new read-only instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitCaller(address common.Address, caller bind.ContractCaller) (*FuxSplitCaller, error) {
	contract, err := bindFuxSplit(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSplitCaller{contract: contract}, nil
}

// NewFuxSplitTransactor creates a new write-only instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxSplitTransactor, error) {
	contract, err := bindFuxSplit(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FuxSplitTransactor{contract: contract}, nil
}

// bindFuxSplit binds a generic wrapper to an already deployed contract.
func bindFuxSplit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxSplitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSplit *FuxSplitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSplit.Contract.FuxSplitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSplit *FuxSplitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSplit.Contract.FuxSplitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSplit *FuxSplitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSplit.Contract.FuxSplitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSplit *FuxSplitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSplit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSplit *FuxSplitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSplit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSplit *FuxSplitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSplit.Contract.contract.Transact(opts, method, params...)
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLE_FUX_BC_GROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_FUX_BC_GROUP")
	return *ret0, err
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLE_FUX_BC_GROUP() (string, error) {
	return _FuxSplit.Contract.ROLE_FUX_BC_GROUP(&_FuxSplit.CallOpts)
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLE_FUX_BC_GROUP() (string, error) {
	return _FuxSplit.Contract.ROLE_FUX_BC_GROUP(&_FuxSplit.CallOpts)
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLE_FUX_HUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLE_FUX_HUB() (string, error) {
	return _FuxSplit.Contract.ROLE_FUX_HUB(&_FuxSplit.CallOpts)
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLE_FUX_HUB() (string, error) {
	return _FuxSplit.Contract.ROLE_FUX_HUB(&_FuxSplit.CallOpts)
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLE_FUX_ISSUER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_FUX_ISSUER")
	return *ret0, err
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLE_FUX_ISSUER() (string, error) {
	return _FuxSplit.Contract.ROLE_FUX_ISSUER(&_FuxSplit.CallOpts)
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLE_FUX_ISSUER() (string, error) {
	return _FuxSplit.Contract.ROLE_FUX_ISSUER(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_COIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_COIN")
	return *ret0, err
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_COIN() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_COIN(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_COIN() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_COIN(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_CONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_CONFIG")
	return *ret0, err
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_CONFIG() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_CONFIG(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_CONFIG() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_CONFIG(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_ERC721_TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_ERC721_TOKEN")
	return *ret0, err
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_ERC721_TOKEN() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_ERC721_TOKEN(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_ERC721_TOKEN() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_ERC721_TOKEN(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_PAYBOX_FACTORY(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_PAYBOX_FACTORY")
	return *ret0, err
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_PAYBOX_FACTORY() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_PAYBOX_FACTORY(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_PAYBOX_FACTORY() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_PAYBOX_FACTORY(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_RBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_RBAC")
	return *ret0, err
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_RBAC() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_RBAC(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_RBAC() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_RBAC(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_SPLIT(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_SPLIT")
	return *ret0, err
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_SPLIT() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_SPLIT(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_SPLIT() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_SPLIT(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_STORAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_STORAGE")
	return *ret0, err
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_STORAGE() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_STORAGE(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_STORAGE() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_STORAGE(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_TOKEN")
	return *ret0, err
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_TOKEN() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_TOKEN(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_TOKEN() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_TOKEN(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTE_FUX_TRANSFER_ROLE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_FUX_TRANSFER_ROLE")
	return *ret0, err
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTE_FUX_TRANSFER_ROLE() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_TRANSFER_ROLE(&_FuxSplit.CallOpts)
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTE_FUX_TRANSFER_ROLE() (string, error) {
	return _FuxSplit.Contract.ROUTE_FUX_TRANSFER_ROLE(&_FuxSplit.CallOpts)
}

// AddrFuxERC721 is a free data retrieval call binding the contract method 0x41a88a1b.
//
// Solidity: function addrFuxERC721() constant returns(address)
func (_FuxSplit *FuxSplitCaller) AddrFuxERC721(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "addrFuxERC721")
	return *ret0, err
}

// AddrFuxERC721 is a free data retrieval call binding the contract method 0x41a88a1b.
//
// Solidity: function addrFuxERC721() constant returns(address)
func (_FuxSplit *FuxSplitSession) AddrFuxERC721() (common.Address, error) {
	return _FuxSplit.Contract.AddrFuxERC721(&_FuxSplit.CallOpts)
}

// AddrFuxERC721 is a free data retrieval call binding the contract method 0x41a88a1b.
//
// Solidity: function addrFuxERC721() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) AddrFuxERC721() (common.Address, error) {
	return _FuxSplit.Contract.AddrFuxERC721(&_FuxSplit.CallOpts)
}

// AddrFuxStorage is a free data retrieval call binding the contract method 0x95c8757e.
//
// Solidity: function addrFuxStorage() constant returns(address)
func (_FuxSplit *FuxSplitCaller) AddrFuxStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "addrFuxStorage")
	return *ret0, err
}

// AddrFuxStorage is a free data retrieval call binding the contract method 0x95c8757e.
//
// Solidity: function addrFuxStorage() constant returns(address)
func (_FuxSplit *FuxSplitSession) AddrFuxStorage() (common.Address, error) {
	return _FuxSplit.Contract.AddrFuxStorage(&_FuxSplit.CallOpts)
}

// AddrFuxStorage is a free data retrieval call binding the contract method 0x95c8757e.
//
// Solidity: function addrFuxStorage() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) AddrFuxStorage() (common.Address, error) {
	return _FuxSplit.Contract.AddrFuxStorage(&_FuxSplit.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxSplit *FuxSplitCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "getProxyAddress")
	return *ret0, err
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxSplit *FuxSplitSession) GetProxyAddress() (common.Address, error) {
	return _FuxSplit.Contract.GetProxyAddress(&_FuxSplit.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) GetProxyAddress() (common.Address, error) {
	return _FuxSplit.Contract.GetProxyAddress(&_FuxSplit.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxSplit *FuxSplitCaller) IsUsingProxy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "isUsingProxy")
	return *ret0, err
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxSplit *FuxSplitSession) IsUsingProxy() (bool, error) {
	return _FuxSplit.Contract.IsUsingProxy(&_FuxSplit.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) IsUsingProxy() (bool, error) {
	return _FuxSplit.Contract.IsUsingProxy(&_FuxSplit.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxSplit *FuxSplitCaller) SAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "sAddr")
	return *ret0, err
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxSplit *FuxSplitSession) SAddr() (common.Address, error) {
	return _FuxSplit.Contract.SAddr(&_FuxSplit.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) SAddr() (common.Address, error) {
	return _FuxSplit.Contract.SAddr(&_FuxSplit.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxSplit *FuxSplitCaller) StaticAddr(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "staticAddr")
	return *ret0, err
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxSplit *FuxSplitSession) StaticAddr() (bool, error) {
	return _FuxSplit.Contract.StaticAddr(&_FuxSplit.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) StaticAddr() (bool, error) {
	return _FuxSplit.Contract.StaticAddr(&_FuxSplit.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxSplit *FuxSplitCaller) SysProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "sysProxy")
	return *ret0, err
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxSplit *FuxSplitSession) SysProxy() (common.Address, error) {
	return _FuxSplit.Contract.SysProxy(&_FuxSplit.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) SysProxy() (common.Address, error) {
	return _FuxSplit.Contract.SysProxy(&_FuxSplit.CallOpts)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _newIds uint256[2], _newCoins uint256[2], _states uint256[2]) returns()
func (_FuxSplit *FuxSplitTransactor) SplitFux(opts *bind.TransactOpts, _tokenId *big.Int, _newIds [2]*big.Int, _newCoins [2]*big.Int, _states [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.contract.Transact(opts, "splitFux", _tokenId, _newIds, _newCoins, _states)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _newIds uint256[2], _newCoins uint256[2], _states uint256[2]) returns()
func (_FuxSplit *FuxSplitSession) SplitFux(_tokenId *big.Int, _newIds [2]*big.Int, _newCoins [2]*big.Int, _states [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.Contract.SplitFux(&_FuxSplit.TransactOpts, _tokenId, _newIds, _newCoins, _states)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _newIds uint256[2], _newCoins uint256[2], _states uint256[2]) returns()
func (_FuxSplit *FuxSplitTransactorSession) SplitFux(_tokenId *big.Int, _newIds [2]*big.Int, _newCoins [2]*big.Int, _states [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.Contract.SplitFux(&_FuxSplit.TransactOpts, _tokenId, _newIds, _newCoins, _states)
}
