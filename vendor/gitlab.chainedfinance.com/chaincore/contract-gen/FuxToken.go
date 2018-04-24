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
const FuxTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TRANSFER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_BC_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_RBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_ISSUER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staticAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_COIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isUsingProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_SPLIT\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_PAYBOX_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_ERC721_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_coin\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"fuxBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxTokenBin is the compiled bytecode used for deploying new contracts.
const FuxTokenBin = `0x606060405260008060146101000a81548160ff02191690831515021790555034156200002a57600080fd5b60405160208062002374833981016040528080519060200190919050506200009b6040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000c16401000000000262001fec176401000000009004565b620000ba81620000dd6401000000000262002006176401000000009004565b506200020d565b8060029080519060200190620000d99291906200015e565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200011a57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a157805160ff1916838001178555620001d2565b82800160010185558215620001d2579182015b82811115620001d1578251825591602001919060010190620001b4565b5b509050620001e19190620001e5565b5090565b6200020a91905b8082111562000206576000816000905550600101620001ec565b5090565b90565b612157806200021d6000396000f300606060405260043610610128576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063071db5611461012d57806309c4b0fa146101bb5780630c7bba78146102495780630daaf6e4146102d75780632bc3a44c14610365578063320618991461039257806333e653e8146103e757806343a73d9a146104755780634aa03a35146104ca57806370a082311461055857806375ae7a8d146105a5578063781bed44146105fa57806388f7748d146106275780639dc29fac146106b5578063beabacc8146106f7578063c0d1932514610758578063cf21977c146107e6578063d6b419fb14610886578063d75ab61614610914578063e941ed6114610961578063f43a6d24146109ef578063f5ac8fb914610a7d575b600080fd5b341561013857600080fd5b610140610b0b565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610180578082015181840152602081019050610165565b50505050905090810190601f1680156101ad5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101c657600080fd5b6101ce610b44565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561020e5780820151818401526020810190506101f3565b50505050905090810190601f16801561023b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561025457600080fd5b61025c610b7d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561029c578082015181840152602081019050610281565b50505050905090810190601f1680156102c95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e257600080fd5b6102ea610bb6565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561032a57808201518184015260208101905061030f565b50505050905090810190601f1680156103575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561037057600080fd5b610378610bef565b604051808215151515815260200191505060405180910390f35b341561039d57600080fd5b6103a5610c02565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103f257600080fd5b6103fa610c28565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561043a57808201518184015260208101905061041f565b50505050905090810190601f1680156104675780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561048057600080fd5b610488610c61565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104d557600080fd5b6104dd610c8b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561051d578082015181840152602081019050610502565b50505050905090810190601f16801561054a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561056357600080fd5b61058f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610cc4565b6040518082815260200191505060405180910390f35b34156105b057600080fd5b6105b8610d85565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561060557600080fd5b61060d610daa565b604051808215151515815260200191505060405180910390f35b341561063257600080fd5b61063a610dc1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561067a57808201518184015260208101905061065f565b50505050905090810190601f1680156106a75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156106c057600080fd5b6106f5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610dfa565b005b341561070257600080fd5b610756600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506110aa565b005b341561076357600080fd5b61076b6114da565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107ab578082015181840152602081019050610790565b50505050905090810190601f1680156107d85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156107f157600080fd5b610884600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001909190803590602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050611513565b005b341561089157600080fd5b6108996117bb565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108d95780820151818401526020810190506108be565b50505050905090810190601f1680156109065780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561091f57600080fd5b61094b600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506117f4565b6040518082815260200191505060405180910390f35b341561096c57600080fd5b6109746118b5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109b4578082015181840152602081019050610999565b50505050905090810190601f1680156109e15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156109fa57600080fd5b610a026118ee565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a42578082015181840152602081019050610a27565b50505050905090810190601f168015610a6f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3415610a8857600080fd5b610a90611927565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610ad0578082015181840152602081019050610ab5565b50505050905090810190601f168015610afd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805190810160405280601381526020017f5472616e73666572526f6c65576974684675780000000000000000000000000081525081565b6040805190810160405280600a81526020017f667578426347726f75700000000000000000000000000000000000000000000081525081565b6040805190810160405280600b81526020017f524241435769746846757800000000000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f667578497373756572000000000000000000000000000000000000000000000081525081565b600060149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600781526020017f467578436f696e0000000000000000000000000000000000000000000000000081525081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525081565b6000610cce611960565b73ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515610d6757600080fd5b5af11515610d7457600080fd5b505050604051805190509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615905090565b6040805190810160405280600881526020017f46757853706c697400000000000000000000000000000000000000000000000081525081565b6000610e05336119a5565b80610e155750610e1433611a66565b5b1515610e2057600080fd5b610e28611b27565b73ffffffffffffffffffffffffffffffffffffffff16632639508d836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1515610e9557600080fd5b5af11515610ea257600080fd5b505050604051805190509050610eb6611b6c565b73ffffffffffffffffffffffffffffffffffffffff166326ffee0884836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1515610f5757600080fd5b5af11515610f6457600080fd5b505050610f6f611960565b73ffffffffffffffffffffffffffffffffffffffff16639dc29fac84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b151561101057600080fd5b5af1151561101d57600080fd5b505050611028611b27565b73ffffffffffffffffffffffffffffffffffffffff1663e234f61d836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b151561109557600080fd5b5af115156110a257600080fd5b505050505050565b600033836110b6611bb1565b73ffffffffffffffffffffffffffffffffffffffff16631de1ed6483836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b151561118357600080fd5b5af1151561119057600080fd5b50505060405180519050806111aa57506111a982611a66565b5b806111ba57506111b981611a66565b5b806111ca57506111c982611bf6565b5b806111da57506111d981611bf6565b5b15156111e557600080fd5b6111ed611b27565b73ffffffffffffffffffffffffffffffffffffffff16632639508d856040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b151561125a57600080fd5b5af1151561126757600080fd5b50505060405180519050925061127b611b6c565b73ffffffffffffffffffffffffffffffffffffffff1663f5d82b6b86856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b151561131c57600080fd5b5af1151561132957600080fd5b505050611334611b6c565b73ffffffffffffffffffffffffffffffffffffffff166326ffee0887856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15156113d557600080fd5b5af115156113e257600080fd5b5050506113ed611960565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc88787876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15156114c257600080fd5b5af115156114cf57600080fd5b505050505050505050565b6040805190810160405280600a81526020017f46757853746f726167650000000000000000000000000000000000000000000081525081565b61151c336119a5565b8061152c575061152b33611a66565b5b151561153757600080fd5b61153f611b27565b73ffffffffffffffffffffffffffffffffffffffff166386b27f4d86868686866040518663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018086815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156115e25780820151818401526020810190506115c7565b50505050905090810190601f16801561160f5780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b151561163157600080fd5b5af1151561163e57600080fd5b505050611649611960565b73ffffffffffffffffffffffffffffffffffffffff166340c10f1987876040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15156116ea57600080fd5b5af115156116f757600080fd5b505050611702611b6c565b73ffffffffffffffffffffffffffffffffffffffff1663f5d82b6b87866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15156117a357600080fd5b5af115156117b057600080fd5b505050505050505050565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b60006117fe611b6c565b73ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561189757600080fd5b5af115156118a457600080fd5b505050604051805190509050919050565b6040805190810160405280601081526020017f467578506179426f78466163746f72790000000000000000000000000000000081525081565b6040805190810160405280600e81526020017f467578455243373231546f6b656e00000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f467578436f6e666967000000000000000000000000000000000000000000000081525081565b60006119a06040805190810160405280600e81526020017f467578455243373231546f6b656e000000000000000000000000000000000000815250611d59565b905090565b60006119af611e6e565b73ffffffffffffffffffffffffffffffffffffffff1663abd7712b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515611a4857600080fd5b5af11515611a5557600080fd5b505050604051805190509050919050565b6000611a70611e6e565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515611b0957600080fd5b5af11515611b1657600080fd5b505050604051805190509050919050565b6000611b676040805190810160405280600a81526020017f46757853746f7261676500000000000000000000000000000000000000000000815250611d59565b905090565b6000611bac6040805190810160405280600781526020017f467578436f696e00000000000000000000000000000000000000000000000000815250611d59565b905090565b6000611bf16040805190810160405280601381526020017f5472616e73666572526f6c655769746846757800000000000000000000000000815250611d59565b905090565b6000611c00611e6e565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611cef578082015181840152602081019050611cd4565b50505050905090810190601f168015611d1c5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1515611d3b57600080fd5b5af11515611d4857600080fd5b505050604051805190509050919050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611e05578082015181840152602081019050611dea565b50505050905090810190601f168015611e325780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1515611e5057600080fd5b5af11515611e5d57600080fd5b505050604051805190509050919050565b6000611e78611e7d565b905090565b60008060149054906101000a900460ff1615611ebc576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050611fe9565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560026040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611fb05780601f10611f8557610100808354040283529160200191611fb0565b820191906000526020600020905b815481529060010190602001808311611f9357829003601f168201915b505092505050602060405180830381600087803b1515611fcf57600080fd5b5af11515611fdc57600080fd5b5050506040518051905090505b90565b8060029080519060200190612002929190612086565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561204257600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106120c757805160ff19168380011785556120f5565b828001600101855582156120f5579182015b828111156120f45782518255916020019190600101906120d9565b5b5090506121029190612106565b5090565b61212891905b8082111561212457600081600090555060010161210c565b5090565b905600a165627a7a72305820b472f7831c0b75126d5fee19f53d7ad7f1a4998675840c3d2f45c8f466c93c130029`

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
