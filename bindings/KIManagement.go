// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kiAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"documentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"documentHash\",\"type\":\"string\"}],\"name\":\"DocumentUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kiAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"documentIndex\",\"type\":\"uint256\"}],\"name\":\"DocumentVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kiAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"KIRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kiAddress\",\"type\":\"address\"}],\"name\":\"KIVerified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_kiAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_documentIndex\",\"type\":\"uint256\"}],\"name\":\"getDocumentDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_kiAddress\",\"type\":\"address\"}],\"name\":\"getKIDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qualification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"licenseType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"experience\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"documentCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keyIndividuals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qualification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"licenseType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"experience\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"documentCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_qualification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_licenseType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_experience\",\"type\":\"string\"}],\"name\":\"registerKI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_documentHash\",\"type\":\"string\"}],\"name\":\"uploadDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_kiAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_documentIndex\",\"type\":\"uint256\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50600180546001600160a01b03191633179055610fb38061002e5f395ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c8063b7f5576d11610058578063b7f5576d146100e2578063bab1104e146100f5578063e2d915cd14610108578063f851a4401461011b575f5ffd5b806323b60fd91461007e5780633e9db283146100a85780639d5caf74146100bd575b5f5ffd5b61009161008c366004610b5e565b610146565b60405161009f929190610bb4565b60405180910390f35b6100bb6100b6366004610c76565b610265565b005b6100d06100cb366004610d2f565b61035f565b60405161009f96959493929190610d4f565b6100bb6100f0366004610db6565b6105ae565b6100bb610103366004610b5e565b6106b7565b6100d0610116366004610d2f565b6108be565b60015461012e906001600160a01b031681565b6040516001600160a01b03909116815260200161009f565b6001600160a01b0382165f9081526020819052604081206005810154606092919084106101b35760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840c8dec6eadacadce840d2dcc8caf60531b60448201526064015b60405180910390fd5b5f848152600482016020526040902060018101548154829160ff169082906101da90610df0565b80601f016020809104026020016040519081016040528092919081815260200182805461020690610df0565b80156102515780601f1061022857610100808354040283529160200191610251565b820191905f5260205f20905b81548152906001019060200180831161023457829003601f168201915b505050505091509350935050509250929050565b335f908152602081905260409020805461027e90610df0565b1590506102c55760405162461bcd60e51b815260206004820152601560248201527412d248185b1c9958591e481c9959da5cdd195c9959605a1b60448201526064016101aa565b335f908152602081905260409020806102de8682610e74565b50600181016102ed8582610e74565b50600281016102fc8482610e74565b506003810161030b8382610e74565b505f600582015560068101805460ff1916905560405133907f1bee31b285db69ab6eb80f05681ff4849ee8e4508c38ca8c0a7ea20b1df95adc90610350908890610f2f565b60405180910390a25050505050565b5f6020819052908152604090208054819061037990610df0565b80601f01602080910402602001604051908101604052809291908181526020018280546103a590610df0565b80156103f05780601f106103c7576101008083540402835291602001916103f0565b820191905f5260205f20905b8154815290600101906020018083116103d357829003601f168201915b50505050509080600101805461040590610df0565b80601f016020809104026020016040519081016040528092919081815260200182805461043190610df0565b801561047c5780601f106104535761010080835404028352916020019161047c565b820191905f5260205f20905b81548152906001019060200180831161045f57829003601f168201915b50505050509080600201805461049190610df0565b80601f01602080910402602001604051908101604052809291908181526020018280546104bd90610df0565b80156105085780601f106104df57610100808354040283529160200191610508565b820191905f5260205f20905b8154815290600101906020018083116104eb57829003601f168201915b50505050509080600301805461051d90610df0565b80601f016020809104026020016040519081016040528092919081815260200182805461054990610df0565b80156105945780601f1061056b57610100808354040283529160200191610594565b820191905f5260205f20905b81548152906001019060200180831161057757829003601f168201915b50505050600583015460069093015491929160ff16905086565b335f908152602081905260409020805481906105c990610df0565b90505f0361060d5760405162461bcd60e51b815260206004820152601160248201527012d2481b9bdd081c9959da5cdd195c9959607a1b60448201526064016101aa565b6040805180820182528381525f60208083018290526005850154825260048501905291909120815181906106419082610e74565b50602091909101516001909101805460ff1916911515919091179055600581015460405133917fdbf93cdefa5fd75a260713231f3452706be0547b64c93a47fef6198769398fff9161069591908690610f41565b60405180910390a2600581018054905f6106ae83610f59565b91905055505050565b6001546001600160a01b0316331461071c5760405162461bcd60e51b815260206004820152602260248201527f4f6e6c792061646d696e2063616e20706572666f726d2074686973206163746960448201526137b760f11b60648201526084016101aa565b6001600160a01b0382165f9081526020819052604090208054819061074090610df0565b90505f036107845760405162461bcd60e51b815260206004820152601160248201527012d2481b9bdd081c9959da5cdd195c9959607a1b60448201526064016101aa565b806005015482106107d05760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840c8dec6eadacadce840d2dcc8caf60531b60448201526064016101aa565b5f828152600482016020526040908190206001908101805460ff19169091179055516001600160a01b038416907f5e60b20b1e57e46b947931427afbd55ac1eca70b5b0590206f576cbc0da596449061082c9085815260200190565b60405180910390a260015f5b826005015481101561086e575f81815260048401602052604090206001015460ff16610866575f915061086e565b600101610838565b5080156108b85760068201805460ff191660011790556040516001600160a01b038516907f0b1d3de35deb6361a1ee7163fced02ebab2cd2507cad6286b57ab51e2717c327905f90a25b50505050565b6001600160a01b0381165f9081526020819052604081206005810154600682015482546060948594859485949293849383926001840192600285019260038601929160ff90911690869061091190610df0565b80601f016020809104026020016040519081016040528092919081815260200182805461093d90610df0565b80156109885780601f1061095f57610100808354040283529160200191610988565b820191905f5260205f20905b81548152906001019060200180831161096b57829003601f168201915b5050505050955084805461099b90610df0565b80601f01602080910402602001604051908101604052809291908181526020018280546109c790610df0565b8015610a125780601f106109e957610100808354040283529160200191610a12565b820191905f5260205f20905b8154815290600101906020018083116109f557829003601f168201915b50505050509450838054610a2590610df0565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5190610df0565b8015610a9c5780601f10610a7357610100808354040283529160200191610a9c565b820191905f5260205f20905b815481529060010190602001808311610a7f57829003601f168201915b50505050509350828054610aaf90610df0565b80601f0160208091040260200160405190810160405280929190818152602001828054610adb90610df0565b8015610b265780601f10610afd57610100808354040283529160200191610b26565b820191905f5260205f20905b815481529060010190602001808311610b0957829003601f168201915b505050505092509650965096509650965096505091939550919395565b80356001600160a01b0381168114610b59575f5ffd5b919050565b5f5f60408385031215610b6f575f5ffd5b610b7883610b43565b946020939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f610bc66040830185610b86565b905082151560208301529392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610bfa575f5ffd5b813567ffffffffffffffff811115610c1457610c14610bd7565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610c4357610c43610bd7565b604052818152838201602001851015610c5a575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f60808587031215610c89575f5ffd5b843567ffffffffffffffff811115610c9f575f5ffd5b610cab87828801610beb565b945050602085013567ffffffffffffffff811115610cc7575f5ffd5b610cd387828801610beb565b935050604085013567ffffffffffffffff811115610cef575f5ffd5b610cfb87828801610beb565b925050606085013567ffffffffffffffff811115610d17575f5ffd5b610d2387828801610beb565b91505092959194509250565b5f60208284031215610d3f575f5ffd5b610d4882610b43565b9392505050565b60c081525f610d6160c0830189610b86565b8281036020840152610d738189610b86565b90508281036040840152610d878188610b86565b90508281036060840152610d9b8187610b86565b6080840195909552505090151560a090910152949350505050565b5f60208284031215610dc6575f5ffd5b813567ffffffffffffffff811115610ddc575f5ffd5b610de884828501610beb565b949350505050565b600181811c90821680610e0457607f821691505b602082108103610e2257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610e6f57805f5260205f20601f840160051c81016020851015610e4d5750805b601f840160051c820191505b81811015610e6c575f8155600101610e59565b50505b505050565b815167ffffffffffffffff811115610e8e57610e8e610bd7565b610ea281610e9c8454610df0565b84610e28565b6020601f821160018114610ed4575f8315610ebd5750848201515b5f19600385901b1c1916600184901b178455610e6c565b5f84815260208120601f198516915b82811015610f035787850151825560209485019460019092019101610ee3565b5084821015610f2057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f610d486020830184610b86565b828152604060208201525f610de86040830184610b86565b5f60018201610f7657634e487b7160e01b5f52601160045260245ffd5b506001019056fea26469706673582212209b19f2f2d2622fb658a536d154e672eb2fae3728c2eb6e35641e9e4d117b87d064736f6c634300081b0033",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// BindingsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingsMetaData.Bin instead.
var BindingsBin = BindingsMetaData.Bin

// DeployBindings deploys a new Ethereum contract, binding an instance of Bindings to it.
func DeployBindings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bindings, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bindings *BindingsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bindings *BindingsSession) Admin() (common.Address, error) {
	return _Bindings.Contract.Admin(&_Bindings.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bindings *BindingsCallerSession) Admin() (common.Address, error) {
	return _Bindings.Contract.Admin(&_Bindings.CallOpts)
}

// GetDocumentDetails is a free data retrieval call binding the contract method 0x23b60fd9.
//
// Solidity: function getDocumentDetails(address _kiAddress, uint256 _documentIndex) view returns(string hash, bool isVerified)
func (_Bindings *BindingsCaller) GetDocumentDetails(opts *bind.CallOpts, _kiAddress common.Address, _documentIndex *big.Int) (struct {
	Hash       string
	IsVerified bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getDocumentDetails", _kiAddress, _documentIndex)

	outstruct := new(struct {
		Hash       string
		IsVerified bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IsVerified = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetDocumentDetails is a free data retrieval call binding the contract method 0x23b60fd9.
//
// Solidity: function getDocumentDetails(address _kiAddress, uint256 _documentIndex) view returns(string hash, bool isVerified)
func (_Bindings *BindingsSession) GetDocumentDetails(_kiAddress common.Address, _documentIndex *big.Int) (struct {
	Hash       string
	IsVerified bool
}, error) {
	return _Bindings.Contract.GetDocumentDetails(&_Bindings.CallOpts, _kiAddress, _documentIndex)
}

// GetDocumentDetails is a free data retrieval call binding the contract method 0x23b60fd9.
//
// Solidity: function getDocumentDetails(address _kiAddress, uint256 _documentIndex) view returns(string hash, bool isVerified)
func (_Bindings *BindingsCallerSession) GetDocumentDetails(_kiAddress common.Address, _documentIndex *big.Int) (struct {
	Hash       string
	IsVerified bool
}, error) {
	return _Bindings.Contract.GetDocumentDetails(&_Bindings.CallOpts, _kiAddress, _documentIndex)
}

// GetKIDetails is a free data retrieval call binding the contract method 0xe2d915cd.
//
// Solidity: function getKIDetails(address _kiAddress) view returns(string name, string qualification, string licenseType, string experience, uint256 documentCount, bool isVerified)
func (_Bindings *BindingsCaller) GetKIDetails(opts *bind.CallOpts, _kiAddress common.Address) (struct {
	Name          string
	Qualification string
	LicenseType   string
	Experience    string
	DocumentCount *big.Int
	IsVerified    bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getKIDetails", _kiAddress)

	outstruct := new(struct {
		Name          string
		Qualification string
		LicenseType   string
		Experience    string
		DocumentCount *big.Int
		IsVerified    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Qualification = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.LicenseType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Experience = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.DocumentCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsVerified = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetKIDetails is a free data retrieval call binding the contract method 0xe2d915cd.
//
// Solidity: function getKIDetails(address _kiAddress) view returns(string name, string qualification, string licenseType, string experience, uint256 documentCount, bool isVerified)
func (_Bindings *BindingsSession) GetKIDetails(_kiAddress common.Address) (struct {
	Name          string
	Qualification string
	LicenseType   string
	Experience    string
	DocumentCount *big.Int
	IsVerified    bool
}, error) {
	return _Bindings.Contract.GetKIDetails(&_Bindings.CallOpts, _kiAddress)
}

// GetKIDetails is a free data retrieval call binding the contract method 0xe2d915cd.
//
// Solidity: function getKIDetails(address _kiAddress) view returns(string name, string qualification, string licenseType, string experience, uint256 documentCount, bool isVerified)
func (_Bindings *BindingsCallerSession) GetKIDetails(_kiAddress common.Address) (struct {
	Name          string
	Qualification string
	LicenseType   string
	Experience    string
	DocumentCount *big.Int
	IsVerified    bool
}, error) {
	return _Bindings.Contract.GetKIDetails(&_Bindings.CallOpts, _kiAddress)
}

// KeyIndividuals is a free data retrieval call binding the contract method 0x9d5caf74.
//
// Solidity: function keyIndividuals(address ) view returns(string name, string qualification, string licenseType, string experience, uint256 documentCount, bool isVerified)
func (_Bindings *BindingsCaller) KeyIndividuals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name          string
	Qualification string
	LicenseType   string
	Experience    string
	DocumentCount *big.Int
	IsVerified    bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "keyIndividuals", arg0)

	outstruct := new(struct {
		Name          string
		Qualification string
		LicenseType   string
		Experience    string
		DocumentCount *big.Int
		IsVerified    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Qualification = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.LicenseType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Experience = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.DocumentCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsVerified = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// KeyIndividuals is a free data retrieval call binding the contract method 0x9d5caf74.
//
// Solidity: function keyIndividuals(address ) view returns(string name, string qualification, string licenseType, string experience, uint256 documentCount, bool isVerified)
func (_Bindings *BindingsSession) KeyIndividuals(arg0 common.Address) (struct {
	Name          string
	Qualification string
	LicenseType   string
	Experience    string
	DocumentCount *big.Int
	IsVerified    bool
}, error) {
	return _Bindings.Contract.KeyIndividuals(&_Bindings.CallOpts, arg0)
}

// KeyIndividuals is a free data retrieval call binding the contract method 0x9d5caf74.
//
// Solidity: function keyIndividuals(address ) view returns(string name, string qualification, string licenseType, string experience, uint256 documentCount, bool isVerified)
func (_Bindings *BindingsCallerSession) KeyIndividuals(arg0 common.Address) (struct {
	Name          string
	Qualification string
	LicenseType   string
	Experience    string
	DocumentCount *big.Int
	IsVerified    bool
}, error) {
	return _Bindings.Contract.KeyIndividuals(&_Bindings.CallOpts, arg0)
}

// RegisterKI is a paid mutator transaction binding the contract method 0x3e9db283.
//
// Solidity: function registerKI(string _name, string _qualification, string _licenseType, string _experience) returns()
func (_Bindings *BindingsTransactor) RegisterKI(opts *bind.TransactOpts, _name string, _qualification string, _licenseType string, _experience string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerKI", _name, _qualification, _licenseType, _experience)
}

// RegisterKI is a paid mutator transaction binding the contract method 0x3e9db283.
//
// Solidity: function registerKI(string _name, string _qualification, string _licenseType, string _experience) returns()
func (_Bindings *BindingsSession) RegisterKI(_name string, _qualification string, _licenseType string, _experience string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterKI(&_Bindings.TransactOpts, _name, _qualification, _licenseType, _experience)
}

// RegisterKI is a paid mutator transaction binding the contract method 0x3e9db283.
//
// Solidity: function registerKI(string _name, string _qualification, string _licenseType, string _experience) returns()
func (_Bindings *BindingsTransactorSession) RegisterKI(_name string, _qualification string, _licenseType string, _experience string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterKI(&_Bindings.TransactOpts, _name, _qualification, _licenseType, _experience)
}

// UploadDocument is a paid mutator transaction binding the contract method 0xb7f5576d.
//
// Solidity: function uploadDocument(string _documentHash) returns()
func (_Bindings *BindingsTransactor) UploadDocument(opts *bind.TransactOpts, _documentHash string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "uploadDocument", _documentHash)
}

// UploadDocument is a paid mutator transaction binding the contract method 0xb7f5576d.
//
// Solidity: function uploadDocument(string _documentHash) returns()
func (_Bindings *BindingsSession) UploadDocument(_documentHash string) (*types.Transaction, error) {
	return _Bindings.Contract.UploadDocument(&_Bindings.TransactOpts, _documentHash)
}

// UploadDocument is a paid mutator transaction binding the contract method 0xb7f5576d.
//
// Solidity: function uploadDocument(string _documentHash) returns()
func (_Bindings *BindingsTransactorSession) UploadDocument(_documentHash string) (*types.Transaction, error) {
	return _Bindings.Contract.UploadDocument(&_Bindings.TransactOpts, _documentHash)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0xbab1104e.
//
// Solidity: function verifyDocument(address _kiAddress, uint256 _documentIndex) returns()
func (_Bindings *BindingsTransactor) VerifyDocument(opts *bind.TransactOpts, _kiAddress common.Address, _documentIndex *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "verifyDocument", _kiAddress, _documentIndex)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0xbab1104e.
//
// Solidity: function verifyDocument(address _kiAddress, uint256 _documentIndex) returns()
func (_Bindings *BindingsSession) VerifyDocument(_kiAddress common.Address, _documentIndex *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.VerifyDocument(&_Bindings.TransactOpts, _kiAddress, _documentIndex)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0xbab1104e.
//
// Solidity: function verifyDocument(address _kiAddress, uint256 _documentIndex) returns()
func (_Bindings *BindingsTransactorSession) VerifyDocument(_kiAddress common.Address, _documentIndex *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.VerifyDocument(&_Bindings.TransactOpts, _kiAddress, _documentIndex)
}

// BindingsDocumentUploadedIterator is returned from FilterDocumentUploaded and is used to iterate over the raw logs and unpacked data for DocumentUploaded events raised by the Bindings contract.
type BindingsDocumentUploadedIterator struct {
	Event *BindingsDocumentUploaded // Event containing the contract specifics and raw log

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
func (it *BindingsDocumentUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDocumentUploaded)
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
		it.Event = new(BindingsDocumentUploaded)
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
func (it *BindingsDocumentUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDocumentUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDocumentUploaded represents a DocumentUploaded event raised by the Bindings contract.
type BindingsDocumentUploaded struct {
	KiAddress     common.Address
	DocumentIndex *big.Int
	DocumentHash  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDocumentUploaded is a free log retrieval operation binding the contract event 0xdbf93cdefa5fd75a260713231f3452706be0547b64c93a47fef6198769398fff.
//
// Solidity: event DocumentUploaded(address indexed kiAddress, uint256 documentIndex, string documentHash)
func (_Bindings *BindingsFilterer) FilterDocumentUploaded(opts *bind.FilterOpts, kiAddress []common.Address) (*BindingsDocumentUploadedIterator, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DocumentUploaded", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsDocumentUploadedIterator{contract: _Bindings.contract, event: "DocumentUploaded", logs: logs, sub: sub}, nil
}

// WatchDocumentUploaded is a free log subscription operation binding the contract event 0xdbf93cdefa5fd75a260713231f3452706be0547b64c93a47fef6198769398fff.
//
// Solidity: event DocumentUploaded(address indexed kiAddress, uint256 documentIndex, string documentHash)
func (_Bindings *BindingsFilterer) WatchDocumentUploaded(opts *bind.WatchOpts, sink chan<- *BindingsDocumentUploaded, kiAddress []common.Address) (event.Subscription, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DocumentUploaded", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDocumentUploaded)
				if err := _Bindings.contract.UnpackLog(event, "DocumentUploaded", log); err != nil {
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

// ParseDocumentUploaded is a log parse operation binding the contract event 0xdbf93cdefa5fd75a260713231f3452706be0547b64c93a47fef6198769398fff.
//
// Solidity: event DocumentUploaded(address indexed kiAddress, uint256 documentIndex, string documentHash)
func (_Bindings *BindingsFilterer) ParseDocumentUploaded(log types.Log) (*BindingsDocumentUploaded, error) {
	event := new(BindingsDocumentUploaded)
	if err := _Bindings.contract.UnpackLog(event, "DocumentUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDocumentVerifiedIterator is returned from FilterDocumentVerified and is used to iterate over the raw logs and unpacked data for DocumentVerified events raised by the Bindings contract.
type BindingsDocumentVerifiedIterator struct {
	Event *BindingsDocumentVerified // Event containing the contract specifics and raw log

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
func (it *BindingsDocumentVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDocumentVerified)
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
		it.Event = new(BindingsDocumentVerified)
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
func (it *BindingsDocumentVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDocumentVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDocumentVerified represents a DocumentVerified event raised by the Bindings contract.
type BindingsDocumentVerified struct {
	KiAddress     common.Address
	DocumentIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDocumentVerified is a free log retrieval operation binding the contract event 0x5e60b20b1e57e46b947931427afbd55ac1eca70b5b0590206f576cbc0da59644.
//
// Solidity: event DocumentVerified(address indexed kiAddress, uint256 documentIndex)
func (_Bindings *BindingsFilterer) FilterDocumentVerified(opts *bind.FilterOpts, kiAddress []common.Address) (*BindingsDocumentVerifiedIterator, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DocumentVerified", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsDocumentVerifiedIterator{contract: _Bindings.contract, event: "DocumentVerified", logs: logs, sub: sub}, nil
}

// WatchDocumentVerified is a free log subscription operation binding the contract event 0x5e60b20b1e57e46b947931427afbd55ac1eca70b5b0590206f576cbc0da59644.
//
// Solidity: event DocumentVerified(address indexed kiAddress, uint256 documentIndex)
func (_Bindings *BindingsFilterer) WatchDocumentVerified(opts *bind.WatchOpts, sink chan<- *BindingsDocumentVerified, kiAddress []common.Address) (event.Subscription, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DocumentVerified", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDocumentVerified)
				if err := _Bindings.contract.UnpackLog(event, "DocumentVerified", log); err != nil {
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

// ParseDocumentVerified is a log parse operation binding the contract event 0x5e60b20b1e57e46b947931427afbd55ac1eca70b5b0590206f576cbc0da59644.
//
// Solidity: event DocumentVerified(address indexed kiAddress, uint256 documentIndex)
func (_Bindings *BindingsFilterer) ParseDocumentVerified(log types.Log) (*BindingsDocumentVerified, error) {
	event := new(BindingsDocumentVerified)
	if err := _Bindings.contract.UnpackLog(event, "DocumentVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsKIRegisteredIterator is returned from FilterKIRegistered and is used to iterate over the raw logs and unpacked data for KIRegistered events raised by the Bindings contract.
type BindingsKIRegisteredIterator struct {
	Event *BindingsKIRegistered // Event containing the contract specifics and raw log

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
func (it *BindingsKIRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsKIRegistered)
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
		it.Event = new(BindingsKIRegistered)
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
func (it *BindingsKIRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsKIRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsKIRegistered represents a KIRegistered event raised by the Bindings contract.
type BindingsKIRegistered struct {
	KiAddress common.Address
	Name      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKIRegistered is a free log retrieval operation binding the contract event 0x1bee31b285db69ab6eb80f05681ff4849ee8e4508c38ca8c0a7ea20b1df95adc.
//
// Solidity: event KIRegistered(address indexed kiAddress, string name)
func (_Bindings *BindingsFilterer) FilterKIRegistered(opts *bind.FilterOpts, kiAddress []common.Address) (*BindingsKIRegisteredIterator, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "KIRegistered", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsKIRegisteredIterator{contract: _Bindings.contract, event: "KIRegistered", logs: logs, sub: sub}, nil
}

// WatchKIRegistered is a free log subscription operation binding the contract event 0x1bee31b285db69ab6eb80f05681ff4849ee8e4508c38ca8c0a7ea20b1df95adc.
//
// Solidity: event KIRegistered(address indexed kiAddress, string name)
func (_Bindings *BindingsFilterer) WatchKIRegistered(opts *bind.WatchOpts, sink chan<- *BindingsKIRegistered, kiAddress []common.Address) (event.Subscription, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "KIRegistered", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsKIRegistered)
				if err := _Bindings.contract.UnpackLog(event, "KIRegistered", log); err != nil {
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

// ParseKIRegistered is a log parse operation binding the contract event 0x1bee31b285db69ab6eb80f05681ff4849ee8e4508c38ca8c0a7ea20b1df95adc.
//
// Solidity: event KIRegistered(address indexed kiAddress, string name)
func (_Bindings *BindingsFilterer) ParseKIRegistered(log types.Log) (*BindingsKIRegistered, error) {
	event := new(BindingsKIRegistered)
	if err := _Bindings.contract.UnpackLog(event, "KIRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsKIVerifiedIterator is returned from FilterKIVerified and is used to iterate over the raw logs and unpacked data for KIVerified events raised by the Bindings contract.
type BindingsKIVerifiedIterator struct {
	Event *BindingsKIVerified // Event containing the contract specifics and raw log

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
func (it *BindingsKIVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsKIVerified)
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
		it.Event = new(BindingsKIVerified)
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
func (it *BindingsKIVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsKIVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsKIVerified represents a KIVerified event raised by the Bindings contract.
type BindingsKIVerified struct {
	KiAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKIVerified is a free log retrieval operation binding the contract event 0x0b1d3de35deb6361a1ee7163fced02ebab2cd2507cad6286b57ab51e2717c327.
//
// Solidity: event KIVerified(address indexed kiAddress)
func (_Bindings *BindingsFilterer) FilterKIVerified(opts *bind.FilterOpts, kiAddress []common.Address) (*BindingsKIVerifiedIterator, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "KIVerified", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsKIVerifiedIterator{contract: _Bindings.contract, event: "KIVerified", logs: logs, sub: sub}, nil
}

// WatchKIVerified is a free log subscription operation binding the contract event 0x0b1d3de35deb6361a1ee7163fced02ebab2cd2507cad6286b57ab51e2717c327.
//
// Solidity: event KIVerified(address indexed kiAddress)
func (_Bindings *BindingsFilterer) WatchKIVerified(opts *bind.WatchOpts, sink chan<- *BindingsKIVerified, kiAddress []common.Address) (event.Subscription, error) {

	var kiAddressRule []interface{}
	for _, kiAddressItem := range kiAddress {
		kiAddressRule = append(kiAddressRule, kiAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "KIVerified", kiAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsKIVerified)
				if err := _Bindings.contract.UnpackLog(event, "KIVerified", log); err != nil {
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

// ParseKIVerified is a log parse operation binding the contract event 0x0b1d3de35deb6361a1ee7163fced02ebab2cd2507cad6286b57ab51e2717c327.
//
// Solidity: event KIVerified(address indexed kiAddress)
func (_Bindings *BindingsFilterer) ParseKIVerified(log types.Log) (*BindingsKIVerified, error) {
	event := new(BindingsKIVerified)
	if err := _Bindings.contract.UnpackLog(event, "KIVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
