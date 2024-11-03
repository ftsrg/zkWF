// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecdh

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EcdhMetaData contains all meta data concerning the Ecdh contract.
var EcdhMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"combo\",\"type\":\"string\"}],\"name\":\"getIntermediateValue\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numParticipants\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"combo\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"sharedValue\",\"type\":\"bytes\"}],\"name\":\"uploadIntermediateValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610c4f380380610c4f83398101604081905261002e9161015d565b80515f805460ff191660ff909216918217905560021180159061005857505f54608060ff90911611155b6100b75760405162461bcd60e51b815260206004820152602660248201527f5061727469636970616e7473206d757374206265206265747765656e203220616044820152650dcc8406264760d31b606482015260840160405180910390fd5b5f5b5f5460ff908116908216101561011257818160ff16815181106100de576100de610281565b602002602001015160015f8360ff1660ff1681526020019081526020015f2090816101099190610319565b506001016100b9565b50506103d3565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561015557610155610119565b604052919050565b5f6020828403121561016d575f5ffd5b81516001600160401b03811115610182575f5ffd5b8201601f81018413610192575f5ffd5b80516001600160401b038111156101ab576101ab610119565b8060051b6101bb6020820161012d565b918252602081840181019290810190878411156101d6575f5ffd5b6020850192505b838310156102765782516001600160401b038111156101fa575f5ffd5b8501603f8101891361020a575f5ffd5b60208101516001600160401b0381111561022657610226610119565b610239601f8201601f191660200161012d565b8181526040838301018b101561024d575f5ffd5b8160408401602083015e5f602083830101528085525050506020820191506020830192506101dd565b979650505050505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c908216806102a957607f821691505b6020821081036102c757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561031457805f5260205f20601f840160051c810160208510156102f25750805b601f840160051c820191505b81811015610311575f81556001016102fe565b50505b505050565b81516001600160401b0381111561033257610332610119565b610346816103408454610295565b846102cd565b6020601f821160018114610378575f83156103615750848201515b5f19600385901b1c1916600184901b178455610311565b5f84815260208120601f198516915b828110156103a75787850151825560209485019460019092019101610387565b50848210156103c457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b61086f806103e05f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806315285fed146100595780634e76a846146100775780635990ebb61461009557806395a9548e146100aa578063bea53dc3146100ca575b5f5ffd5b6100616100dd565b60405161006e91906104df565b60405180910390f35b5f546100839060ff1681565b60405160ff909116815260200161006e565b6100a86100a33660046105f2565b61020d565b005b6100bd6100b836600461066a565b6102f3565b60405161006e919061068a565b6100bd6100d836600461069c565b61038a565b5f80546060919060ff1667ffffffffffffffff8111156100ff576100ff610542565b60405190808252806020026020018201604052801561013257816020015b606081526020019060019003908161011d5790505b5090505f5b5f5460ff90811690821610156102075760ff81165f9081526001602052604090208054610163906106d6565b80601f016020809104026020016040519081016040528092919081815260200182805461018f906106d6565b80156101da5780601f106101b1576101008083540402835291602001916101da565b820191905f5260205f20905b8154815290600101906020018083116101bd57829003601f168201915b5050505050828260ff16815181106101f4576101f4610708565b6020908102919091010152600101610137565b50919050565b60038260405161021d919061071c565b9081526040519081900360200190205460ff16156102915760405162461bcd60e51b815260206004820152602660248201527f56616c756520616c72656164792073657420666f72207468697320636f6d62696044820152653730ba34b7b760d11b60648201526084015b60405180910390fd5b806002836040516102a2919061071c565b908152602001604051809103902090816102bc919061077e565b5060016003836040516102cf919061071c565b908152604051908190036020019020805491151560ff199092169190911790555050565b60016020525f90815260409020805461030b906106d6565b80601f0160208091040260200160405190810160405280929190818152602001828054610337906106d6565b80156103825780601f1061035957610100808354040283529160200191610382565b820191905f5260205f20905b81548152906001019060200180831161036557829003601f168201915b505050505081565b606060038260405161039c919061071c565b9081526040519081900360200190205460ff166104055760405162461bcd60e51b815260206004820152602160248201527f4e6f2076616c75652073657420666f72207468697320636f6d62696e6174696f6044820152603760f91b6064820152608401610288565b600282604051610415919061071c565b9081526020016040518091039020805461042e906106d6565b80601f016020809104026020016040519081016040528092919081815260200182805461045a906106d6565b80156104a55780601f1061047c576101008083540402835291602001916104a5565b820191905f5260205f20905b81548152906001019060200180831161048857829003601f168201915b50505050509050919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561053657603f198786030184526105218583516104b1565b94506020938401939190910190600101610505565b50929695505050505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f67ffffffffffffffff84111561057057610570610542565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561059f5761059f610542565b6040528381529050808284018510156105b6575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f8301126105dc575f5ffd5b6105eb83833560208501610556565b9392505050565b5f5f60408385031215610603575f5ffd5b823567ffffffffffffffff811115610619575f5ffd5b610625858286016105cd565b925050602083013567ffffffffffffffff811115610641575f5ffd5b8301601f81018513610651575f5ffd5b61066085823560208401610556565b9150509250929050565b5f6020828403121561067a575f5ffd5b813560ff811681146105eb575f5ffd5b602081525f6105eb60208301846104b1565b5f602082840312156106ac575f5ffd5b813567ffffffffffffffff8111156106c2575f5ffd5b6106ce848285016105cd565b949350505050565b600181811c908216806106ea57607f821691505b60208210810361020757634e487b7160e01b5f52602260045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f82518060208501845e5f920191825250919050565b601f82111561077957805f5260205f20601f840160051c810160208510156107575750805b601f840160051c820191505b81811015610776575f8155600101610763565b50505b505050565b815167ffffffffffffffff81111561079857610798610542565b6107ac816107a684546106d6565b84610732565b6020601f8211600181146107de575f83156107c75750848201515b5f19600385901b1c1916600184901b178455610776565b5f84815260208120601f198516915b8281101561080d57878501518255602094850194600190920191016107ed565b508482101561082a57868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea264697066735822122008675362f52dacd687c03c0590e59b90e82934000d72a7c4460fe9ac640914e964736f6c634300081c0033",
}

// EcdhABI is the input ABI used to generate the binding from.
// Deprecated: Use EcdhMetaData.ABI instead.
var EcdhABI = EcdhMetaData.ABI

// EcdhBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EcdhMetaData.Bin instead.
var EcdhBin = EcdhMetaData.Bin

// DeployEcdh deploys a new Ethereum contract, binding an instance of Ecdh to it.
func DeployEcdh(auth *bind.TransactOpts, backend bind.ContractBackend, _publicKeys [][]byte) (common.Address, *types.Transaction, *Ecdh, error) {
	parsed, err := EcdhMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EcdhBin), backend, _publicKeys)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ecdh{EcdhCaller: EcdhCaller{contract: contract}, EcdhTransactor: EcdhTransactor{contract: contract}, EcdhFilterer: EcdhFilterer{contract: contract}}, nil
}

// Ecdh is an auto generated Go binding around an Ethereum contract.
type Ecdh struct {
	EcdhCaller     // Read-only binding to the contract
	EcdhTransactor // Write-only binding to the contract
	EcdhFilterer   // Log filterer for contract events
}

// EcdhCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcdhCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcdhTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcdhTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcdhFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcdhFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcdhSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcdhSession struct {
	Contract     *Ecdh             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcdhCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcdhCallerSession struct {
	Contract *EcdhCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EcdhTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcdhTransactorSession struct {
	Contract     *EcdhTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcdhRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcdhRaw struct {
	Contract *Ecdh // Generic contract binding to access the raw methods on
}

// EcdhCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcdhCallerRaw struct {
	Contract *EcdhCaller // Generic read-only contract binding to access the raw methods on
}

// EcdhTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcdhTransactorRaw struct {
	Contract *EcdhTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcdh creates a new instance of Ecdh, bound to a specific deployed contract.
func NewEcdh(address common.Address, backend bind.ContractBackend) (*Ecdh, error) {
	contract, err := bindEcdh(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecdh{EcdhCaller: EcdhCaller{contract: contract}, EcdhTransactor: EcdhTransactor{contract: contract}, EcdhFilterer: EcdhFilterer{contract: contract}}, nil
}

// NewEcdhCaller creates a new read-only instance of Ecdh, bound to a specific deployed contract.
func NewEcdhCaller(address common.Address, caller bind.ContractCaller) (*EcdhCaller, error) {
	contract, err := bindEcdh(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcdhCaller{contract: contract}, nil
}

// NewEcdhTransactor creates a new write-only instance of Ecdh, bound to a specific deployed contract.
func NewEcdhTransactor(address common.Address, transactor bind.ContractTransactor) (*EcdhTransactor, error) {
	contract, err := bindEcdh(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcdhTransactor{contract: contract}, nil
}

// NewEcdhFilterer creates a new log filterer instance of Ecdh, bound to a specific deployed contract.
func NewEcdhFilterer(address common.Address, filterer bind.ContractFilterer) (*EcdhFilterer, error) {
	contract, err := bindEcdh(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcdhFilterer{contract: contract}, nil
}

// bindEcdh binds a generic wrapper to an already deployed contract.
func bindEcdh(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcdhMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecdh *EcdhRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecdh.Contract.EcdhCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecdh *EcdhRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecdh.Contract.EcdhTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecdh *EcdhRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecdh.Contract.EcdhTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecdh *EcdhCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecdh.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecdh *EcdhTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecdh.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecdh *EcdhTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecdh.Contract.contract.Transact(opts, method, params...)
}

// GetIntermediateValue is a free data retrieval call binding the contract method 0xbea53dc3.
//
// Solidity: function getIntermediateValue(string combo) view returns(bytes)
func (_Ecdh *EcdhCaller) GetIntermediateValue(opts *bind.CallOpts, combo string) ([]byte, error) {
	var out []interface{}
	err := _Ecdh.contract.Call(opts, &out, "getIntermediateValue", combo)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetIntermediateValue is a free data retrieval call binding the contract method 0xbea53dc3.
//
// Solidity: function getIntermediateValue(string combo) view returns(bytes)
func (_Ecdh *EcdhSession) GetIntermediateValue(combo string) ([]byte, error) {
	return _Ecdh.Contract.GetIntermediateValue(&_Ecdh.CallOpts, combo)
}

// GetIntermediateValue is a free data retrieval call binding the contract method 0xbea53dc3.
//
// Solidity: function getIntermediateValue(string combo) view returns(bytes)
func (_Ecdh *EcdhCallerSession) GetIntermediateValue(combo string) ([]byte, error) {
	return _Ecdh.Contract.GetIntermediateValue(&_Ecdh.CallOpts, combo)
}

// GetPublicKeys is a free data retrieval call binding the contract method 0x15285fed.
//
// Solidity: function getPublicKeys() view returns(bytes[])
func (_Ecdh *EcdhCaller) GetPublicKeys(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Ecdh.contract.Call(opts, &out, "getPublicKeys")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetPublicKeys is a free data retrieval call binding the contract method 0x15285fed.
//
// Solidity: function getPublicKeys() view returns(bytes[])
func (_Ecdh *EcdhSession) GetPublicKeys() ([][]byte, error) {
	return _Ecdh.Contract.GetPublicKeys(&_Ecdh.CallOpts)
}

// GetPublicKeys is a free data retrieval call binding the contract method 0x15285fed.
//
// Solidity: function getPublicKeys() view returns(bytes[])
func (_Ecdh *EcdhCallerSession) GetPublicKeys() ([][]byte, error) {
	return _Ecdh.Contract.GetPublicKeys(&_Ecdh.CallOpts)
}

// NumParticipants is a free data retrieval call binding the contract method 0x4e76a846.
//
// Solidity: function numParticipants() view returns(uint8)
func (_Ecdh *EcdhCaller) NumParticipants(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ecdh.contract.Call(opts, &out, "numParticipants")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NumParticipants is a free data retrieval call binding the contract method 0x4e76a846.
//
// Solidity: function numParticipants() view returns(uint8)
func (_Ecdh *EcdhSession) NumParticipants() (uint8, error) {
	return _Ecdh.Contract.NumParticipants(&_Ecdh.CallOpts)
}

// NumParticipants is a free data retrieval call binding the contract method 0x4e76a846.
//
// Solidity: function numParticipants() view returns(uint8)
func (_Ecdh *EcdhCallerSession) NumParticipants() (uint8, error) {
	return _Ecdh.Contract.NumParticipants(&_Ecdh.CallOpts)
}

// PublicKeys is a free data retrieval call binding the contract method 0x95a9548e.
//
// Solidity: function publicKeys(uint8 ) view returns(bytes)
func (_Ecdh *EcdhCaller) PublicKeys(opts *bind.CallOpts, arg0 uint8) ([]byte, error) {
	var out []interface{}
	err := _Ecdh.contract.Call(opts, &out, "publicKeys", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PublicKeys is a free data retrieval call binding the contract method 0x95a9548e.
//
// Solidity: function publicKeys(uint8 ) view returns(bytes)
func (_Ecdh *EcdhSession) PublicKeys(arg0 uint8) ([]byte, error) {
	return _Ecdh.Contract.PublicKeys(&_Ecdh.CallOpts, arg0)
}

// PublicKeys is a free data retrieval call binding the contract method 0x95a9548e.
//
// Solidity: function publicKeys(uint8 ) view returns(bytes)
func (_Ecdh *EcdhCallerSession) PublicKeys(arg0 uint8) ([]byte, error) {
	return _Ecdh.Contract.PublicKeys(&_Ecdh.CallOpts, arg0)
}

// UploadIntermediateValue is a paid mutator transaction binding the contract method 0x5990ebb6.
//
// Solidity: function uploadIntermediateValue(string combo, bytes sharedValue) returns()
func (_Ecdh *EcdhTransactor) UploadIntermediateValue(opts *bind.TransactOpts, combo string, sharedValue []byte) (*types.Transaction, error) {
	return _Ecdh.contract.Transact(opts, "uploadIntermediateValue", combo, sharedValue)
}

// UploadIntermediateValue is a paid mutator transaction binding the contract method 0x5990ebb6.
//
// Solidity: function uploadIntermediateValue(string combo, bytes sharedValue) returns()
func (_Ecdh *EcdhSession) UploadIntermediateValue(combo string, sharedValue []byte) (*types.Transaction, error) {
	return _Ecdh.Contract.UploadIntermediateValue(&_Ecdh.TransactOpts, combo, sharedValue)
}

// UploadIntermediateValue is a paid mutator transaction binding the contract method 0x5990ebb6.
//
// Solidity: function uploadIntermediateValue(string combo, bytes sharedValue) returns()
func (_Ecdh *EcdhTransactorSession) UploadIntermediateValue(combo string, sharedValue []byte) (*types.Transaction, error) {
	return _Ecdh.Contract.UploadIntermediateValue(&_Ecdh.TransactOpts, combo, sharedValue)
}
