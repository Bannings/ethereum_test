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
const FuxTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TRANSFER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_BC_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_RBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_ISSUER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staticAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_COIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_BATCH\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isUsingProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_SPLIT\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_PAYBOX_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_ERC721_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_coin\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"fuxBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxTokenBin is the compiled bytecode used for deploying new contracts.
const FuxTokenBin = `0x608060405260008060146101000a81548160ff0219169083151502179055503480156200002b57600080fd5b5060405160208062002510833981018060405281019080805190602001909291905050506200009e6040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000bf640100000000026401000000009004565b620000b881620000db640100000000026401000000009004565b506200020b565b8060029080519060200190620000d79291906200015c565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200011857600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019f57805160ff1916838001178555620001d0565b82800160010185558215620001d0579182015b82811115620001cf578251825591602001919060010190620001b2565b5b509050620001df9190620001e3565b5090565b6200020891905b8082111562000204576000816000905550600101620001ea565b5090565b90565b6122f5806200021b6000396000f300608060405260043610610133576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063071db5611461013857806309c4b0fa146101c85780630c7bba78146102585780630daaf6e4146102e85780632bc3a44c1461037857806332061899146103a757806333e653e8146103fe57806343a73d9a1461048e5780634aa03a35146104e5578063510c1b231461057557806370a082311461060557806375ae7a8d1461065c578063781bed44146106b357806388f7748d146106e25780639dc29fac14610772578063beabacc8146107bf578063c0d193251461082c578063cf21977c146108bc578063d6b419fb1461096d578063d75ab616146109fd578063e941ed6114610a54578063f43a6d2414610ae4578063f5ac8fb914610b74575b600080fd5b34801561014457600080fd5b5061014d610c04565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561018d578082015181840152602081019050610172565b50505050905090810190601f1680156101ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101d457600080fd5b506101dd610c3d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561021d578082015181840152602081019050610202565b50505050905090810190601f16801561024a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026457600080fd5b5061026d610c76565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102ad578082015181840152602081019050610292565b50505050905090810190601f1680156102da5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102f457600080fd5b506102fd610caf565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561033d578082015181840152602081019050610322565b50505050905090810190601f16801561036a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038457600080fd5b5061038d610ce8565b604051808215151515815260200191505060405180910390f35b3480156103b357600080fd5b506103bc610cfb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561040a57600080fd5b50610413610d21565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610453578082015181840152602081019050610438565b50505050905090810190601f1680156104805780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561049a57600080fd5b506104a3610d5a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f157600080fd5b506104fa610d84565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561053a57808201518184015260208101905061051f565b50505050905090810190601f1680156105675780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058157600080fd5b5061058a610dbd565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105ca5780820151818401526020810190506105af565b50505050905090810190601f1680156105f75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561061157600080fd5b50610646600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610df6565b6040518082815260200191505060405180910390f35b34801561066857600080fd5b50610671610edc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106bf57600080fd5b506106c8610f01565b604051808215151515815260200191505060405180910390f35b3480156106ee57600080fd5b506106f7610f18565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561073757808201518184015260208101905061071c565b50505050905090810190601f1680156107645780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561077e57600080fd5b506107bd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f51565b005b3480156107cb57600080fd5b5061082a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611241565b005b34801561083857600080fd5b5061084161158f565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610881578082015181840152602081019050610866565b50505050905090810190601f1680156108ae5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156108c857600080fd5b5061096b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506115c8565b005b34801561097957600080fd5b5061098261188b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109c25780820151818401526020810190506109a7565b50505050905090810190601f1680156109ef5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610a0957600080fd5b50610a3e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506118c4565b6040518082815260200191505060405180910390f35b348015610a6057600080fd5b50610a696119aa565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610aa9578082015181840152602081019050610a8e565b50505050905090810190601f168015610ad65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610af057600080fd5b50610af96119e3565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b39578082015181840152602081019050610b1e565b50505050905090810190601f168015610b665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610b8057600080fd5b50610b89611a1c565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610bc9578082015181840152602081019050610bae565b50505050905090810190601f168015610bf65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805190810160405280601381526020017f5472616e73666572526f6c65576974684675780000000000000000000000000081525081565b6040805190810160405280600a81526020017f667578426347726f75700000000000000000000000000000000000000000000081525081565b6040805190810160405280600b81526020017f524241435769746846757800000000000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f667578497373756572000000000000000000000000000000000000000000000081525081565b600060149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600781526020017f467578436f696e0000000000000000000000000000000000000000000000000081525081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525081565b6040805190810160405280600881526020017f467578426174636800000000000000000000000000000000000000000000000081525081565b6000610e00611a55565b73ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610e9a57600080fd5b505af1158015610eae573d6000803e3d6000fd5b505050506040513d6020811015610ec457600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615905090565b6040805190810160405280600881526020017f46757853706c697400000000000000000000000000000000000000000000000081525081565b6000610f5c33611a9a565b80610f6c5750610f6b33611b80565b5b1515610f7757600080fd5b610f7f611c66565b73ffffffffffffffffffffffffffffffffffffffff16632639508d836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015610fed57600080fd5b505af1158015611001573d6000803e3d6000fd5b505050506040513d602081101561101757600080fd5b81019080805190602001909291905050509050611032611cab565b73ffffffffffffffffffffffffffffffffffffffff166326ffee0884836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156110d457600080fd5b505af11580156110e8573d6000803e3d6000fd5b505050506110f4611a55565b73ffffffffffffffffffffffffffffffffffffffff16639dc29fac84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561119657600080fd5b505af11580156111aa573d6000803e3d6000fd5b505050506111b6611c66565b73ffffffffffffffffffffffffffffffffffffffff1663e234f61d836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561122457600080fd5b505af1158015611238573d6000803e3d6000fd5b50505050505050565b6000838361124f8282611cf0565b151561125a57600080fd5b611262611c66565b73ffffffffffffffffffffffffffffffffffffffff16632639508d856040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156112d057600080fd5b505af11580156112e4573d6000803e3d6000fd5b505050506040513d60208110156112fa57600080fd5b81019080805190602001909291905050509250611315611cab565b73ffffffffffffffffffffffffffffffffffffffff1663f5d82b6b86856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156113b757600080fd5b505af11580156113cb573d6000803e3d6000fd5b505050506113d7611cab565b73ffffffffffffffffffffffffffffffffffffffff166326ffee0887856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561147957600080fd5b505af115801561148d573d6000803e3d6000fd5b50505050611499611a55565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc88787876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561156f57600080fd5b505af1158015611583573d6000803e3d6000fd5b50505050505050505050565b6040805190810160405280600a81526020017f46757853746f726167650000000000000000000000000000000000000000000081525081565b6115d133611a9a565b806115e157506115e033611b80565b5b15156115ec57600080fd5b6115f4611c66565b73ffffffffffffffffffffffffffffffffffffffff166386b27f4d86868686866040518663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018086815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561169757808201518184015260208101905061167c565b50505050905090810190601f1680156116c45780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b1580156116e757600080fd5b505af11580156116fb573d6000803e3d6000fd5b50505050611707611a55565b73ffffffffffffffffffffffffffffffffffffffff166340c10f1987876040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156117a957600080fd5b505af11580156117bd573d6000803e3d6000fd5b505050506117c9611cab565b73ffffffffffffffffffffffffffffffffffffffff1663f5d82b6b87866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561186b57600080fd5b505af115801561187f573d6000803e3d6000fd5b50505050505050505050565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b60006118ce611cab565b73ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561196857600080fd5b505af115801561197c573d6000803e3d6000fd5b505050506040513d602081101561199257600080fd5b81019080805190602001909291905050509050919050565b6040805190810160405280601081526020017f467578506179426f78466163746f72790000000000000000000000000000000081525081565b6040805190810160405280600e81526020017f467578455243373231546f6b656e00000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f467578436f6e666967000000000000000000000000000000000000000000000081525081565b6000611a956040805190810160405280600e81526020017f467578455243373231546f6b656e000000000000000000000000000000000000815250611e3b565b905090565b6000611aa4611f75565b73ffffffffffffffffffffffffffffffffffffffff1663abd7712b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611b3e57600080fd5b505af1158015611b52573d6000803e3d6000fd5b505050506040513d6020811015611b6857600080fd5b81019080805190602001909291905050509050919050565b6000611b8a611f75565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611c2457600080fd5b505af1158015611c38573d6000803e3d6000fd5b505050506040513d6020811015611c4e57600080fd5b81019080805190602001909291905050509050919050565b6000611ca66040805190810160405280600a81526020017f46757853746f7261676500000000000000000000000000000000000000000000815250611e3b565b905090565b6000611ceb6040805190810160405280600781526020017f467578436f696e00000000000000000000000000000000000000000000000000815250611e3b565b905090565b6000611cfa611f84565b73ffffffffffffffffffffffffffffffffffffffff16631de1ed6484846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611dc857600080fd5b505af1158015611ddc573d6000803e3d6000fd5b505050506040513d6020811015611df257600080fd5b810190808051906020019092919050505080611e135750611e1283611fc9565b5b80611e235750611e2282611fc9565b5b80611e335750611e3233611fc9565b5b905092915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ee7578082015181840152602081019050611ecc565b50505050905090810190601f168015611f145780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015611f3357600080fd5b505af1158015611f47573d6000803e3d6000fd5b505050506040513d6020811015611f5d57600080fd5b81019080805190602001909291905050509050919050565b6000611f7f611feb565b905090565b6000611fc46040805190810160405280601381526020017f5472616e73666572526f6c655769746846757800000000000000000000000000815250611e3b565b905090565b6000611fd482611b80565b80611fe45750611fe382612141565b5b9050919050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560026040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156120e15780601f106120b6576101008083540402835291602001916120e1565b820191906000526020600020905b8154815290600101906020018083116120c457829003601f168201915b505092505050602060405180830381600087803b15801561210157600080fd5b505af1158015612115573d6000803e3d6000fd5b505050506040513d602081101561212b57600080fd5b8101908080519060200190929190505050905090565b600061214b611f75565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561223a57808201518184015260208101905061221f565b50505050905090810190601f1680156122675780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561228757600080fd5b505af115801561229b573d6000803e3d6000fd5b505050506040513d60208110156122b157600080fd5b810190808051906020019092919050505090509190505600a165627a7a7230582032aba3787778b6daecb8674f2e57b633cb313b963ffe70ece74af902eebb3ce40029`

// DeployFuxToken deploys a new Ethereum contract, binding an instance of FuxToken to it.
func DeployFuxToken(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *FuxToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxTokenBin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxToken{FuxTokenCaller: FuxTokenCaller{contract: contract}, FuxTokenTransactor: FuxTokenTransactor{contract: contract}, FuxTokenFilterer: FuxTokenFilterer{contract: contract}}, nil
}

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

// ROLEFUXBCGROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEFUXBCGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_FUX_BC_GROUP")
	return *ret0, err
}

// ROLEFUXBCGROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEFUXBCGROUP() (string, error) {
	return _FuxToken.Contract.ROLEFUXBCGROUP(&_FuxToken.CallOpts)
}

// ROLEFUXBCGROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEFUXBCGROUP() (string, error) {
	return _FuxToken.Contract.ROLEFUXBCGROUP(&_FuxToken.CallOpts)
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

// ROLEFUXISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEFUXISSUER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_FUX_ISSUER")
	return *ret0, err
}

// ROLEFUXISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEFUXISSUER() (string, error) {
	return _FuxToken.Contract.ROLEFUXISSUER(&_FuxToken.CallOpts)
}

// ROLEFUXISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEFUXISSUER() (string, error) {
	return _FuxToken.Contract.ROLEFUXISSUER(&_FuxToken.CallOpts)
}

// ROUTEFUXBATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXBATCH(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_BATCH")
	return *ret0, err
}

// ROUTEFUXBATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXBATCH() (string, error) {
	return _FuxToken.Contract.ROUTEFUXBATCH(&_FuxToken.CallOpts)
}

// ROUTEFUXBATCH is a free data retrieval call binding the contract method 0x510c1b23.
//
// Solidity: function ROUTE_FUX_BATCH() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXBATCH() (string, error) {
	return _FuxToken.Contract.ROUTEFUXBATCH(&_FuxToken.CallOpts)
}

// ROUTEFUXCOIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXCOIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_COIN")
	return *ret0, err
}

// ROUTEFUXCOIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXCOIN() (string, error) {
	return _FuxToken.Contract.ROUTEFUXCOIN(&_FuxToken.CallOpts)
}

// ROUTEFUXCOIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXCOIN() (string, error) {
	return _FuxToken.Contract.ROUTEFUXCOIN(&_FuxToken.CallOpts)
}

// ROUTEFUXCONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXCONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_CONFIG")
	return *ret0, err
}

// ROUTEFUXCONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXCONFIG() (string, error) {
	return _FuxToken.Contract.ROUTEFUXCONFIG(&_FuxToken.CallOpts)
}

// ROUTEFUXCONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXCONFIG() (string, error) {
	return _FuxToken.Contract.ROUTEFUXCONFIG(&_FuxToken.CallOpts)
}

// ROUTEFUXERC721TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXERC721TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_ERC721_TOKEN")
	return *ret0, err
}

// ROUTEFUXERC721TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXERC721TOKEN() (string, error) {
	return _FuxToken.Contract.ROUTEFUXERC721TOKEN(&_FuxToken.CallOpts)
}

// ROUTEFUXERC721TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXERC721TOKEN() (string, error) {
	return _FuxToken.Contract.ROUTEFUXERC721TOKEN(&_FuxToken.CallOpts)
}

// ROUTEFUXPAYBOXFACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXPAYBOXFACTORY(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_PAYBOX_FACTORY")
	return *ret0, err
}

// ROUTEFUXPAYBOXFACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXPAYBOXFACTORY() (string, error) {
	return _FuxToken.Contract.ROUTEFUXPAYBOXFACTORY(&_FuxToken.CallOpts)
}

// ROUTEFUXPAYBOXFACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXPAYBOXFACTORY() (string, error) {
	return _FuxToken.Contract.ROUTEFUXPAYBOXFACTORY(&_FuxToken.CallOpts)
}

// ROUTEFUXRBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_RBAC")
	return *ret0, err
}

// ROUTEFUXRBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXRBAC() (string, error) {
	return _FuxToken.Contract.ROUTEFUXRBAC(&_FuxToken.CallOpts)
}

// ROUTEFUXRBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXRBAC() (string, error) {
	return _FuxToken.Contract.ROUTEFUXRBAC(&_FuxToken.CallOpts)
}

// ROUTEFUXSPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXSPLIT(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_SPLIT")
	return *ret0, err
}

// ROUTEFUXSPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXSPLIT() (string, error) {
	return _FuxToken.Contract.ROUTEFUXSPLIT(&_FuxToken.CallOpts)
}

// ROUTEFUXSPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXSPLIT() (string, error) {
	return _FuxToken.Contract.ROUTEFUXSPLIT(&_FuxToken.CallOpts)
}

// ROUTEFUXSTORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXSTORAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_STORAGE")
	return *ret0, err
}

// ROUTEFUXSTORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXSTORAGE() (string, error) {
	return _FuxToken.Contract.ROUTEFUXSTORAGE(&_FuxToken.CallOpts)
}

// ROUTEFUXSTORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXSTORAGE() (string, error) {
	return _FuxToken.Contract.ROUTEFUXSTORAGE(&_FuxToken.CallOpts)
}

// ROUTEFUXTOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXTOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_TOKEN")
	return *ret0, err
}

// ROUTEFUXTOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXTOKEN() (string, error) {
	return _FuxToken.Contract.ROUTEFUXTOKEN(&_FuxToken.CallOpts)
}

// ROUTEFUXTOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXTOKEN() (string, error) {
	return _FuxToken.Contract.ROUTEFUXTOKEN(&_FuxToken.CallOpts)
}

// ROUTEFUXTRANSFERROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTEFUXTRANSFERROLE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_FUX_TRANSFER_ROLE")
	return *ret0, err
}

// ROUTEFUXTRANSFERROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTEFUXTRANSFERROLE() (string, error) {
	return _FuxToken.Contract.ROUTEFUXTRANSFERROLE(&_FuxToken.CallOpts)
}

// ROUTEFUXTRANSFERROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTEFUXTRANSFERROLE() (string, error) {
	return _FuxToken.Contract.ROUTEFUXTRANSFERROLE(&_FuxToken.CallOpts)
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

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxToken *FuxTokenCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "getProxyAddress")
	return *ret0, err
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxToken *FuxTokenSession) GetProxyAddress() (common.Address, error) {
	return _FuxToken.Contract.GetProxyAddress(&_FuxToken.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) GetProxyAddress() (common.Address, error) {
	return _FuxToken.Contract.GetProxyAddress(&_FuxToken.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsUsingProxy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "isUsingProxy")
	return *ret0, err
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxToken *FuxTokenSession) IsUsingProxy() (bool, error) {
	return _FuxToken.Contract.IsUsingProxy(&_FuxToken.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsUsingProxy() (bool, error) {
	return _FuxToken.Contract.IsUsingProxy(&_FuxToken.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxToken *FuxTokenCaller) SAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "sAddr")
	return *ret0, err
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxToken *FuxTokenSession) SAddr() (common.Address, error) {
	return _FuxToken.Contract.SAddr(&_FuxToken.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) SAddr() (common.Address, error) {
	return _FuxToken.Contract.SAddr(&_FuxToken.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxToken *FuxTokenCaller) StaticAddr(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "staticAddr")
	return *ret0, err
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxToken *FuxTokenSession) StaticAddr() (bool, error) {
	return _FuxToken.Contract.StaticAddr(&_FuxToken.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) StaticAddr() (bool, error) {
	return _FuxToken.Contract.StaticAddr(&_FuxToken.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxToken *FuxTokenCaller) SysProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "sysProxy")
	return *ret0, err
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxToken *FuxTokenSession) SysProxy() (common.Address, error) {
	return _FuxToken.Contract.SysProxy(&_FuxToken.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) SysProxy() (common.Address, error) {
	return _FuxToken.Contract.SysProxy(&_FuxToken.CallOpts)
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
