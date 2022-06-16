// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rch_state

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	//_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RchMgrInterfaceABI is the input ABI used to generate the binding from.
const RchMgrInterfaceABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner2\",\"type\":\"address\"}],\"name\":\"ChangeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"GetContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"SetAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"SetContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"contractAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RchMgrInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var RchMgrInterfaceFuncSigs = map[string]string{
	"f2853292": "ChangeOwner(address)",
	"d4527406": "GetContractAddr(string)",
	"b11e3c1a": "IsAdmin(address)",
	"55a5194b": "SetAdmin(address,bool)",
	"87ae4ac2": "SetContract(address,string)",
	"27304dfd": "adminList(address)",
	"98ef5e3e": "contractAddrs(string)",
	"54fd4d50": "version()",
}

// RchMgrInterfaceBin is the compiled bytecode used for deploying new contracts.
var RchMgrInterfaceBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905560408051808201909152601b8082527f5263684d6772496e746572666163652076657273696f6e20312e300000000000602090920191825261006791600191610090565b50600080546001600160a01b03168152600260205260409020805460ff1916600117905561012b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100d157805160ff19168380011785556100fe565b828001600101855582156100fe579182015b828111156100fe5782518255916020019190600101906100e3565b5061010a92915061010e565b5090565b61012891905b8082111561010a5760008155600101610114565b90565b6107bd8061013a6000396000f3fe60806040526004361061007b5760003560e01c806398ef5e3e1161004e57806398ef5e3e14610251578063b11e3c1a14610320578063d452740614610353578063f2853292146104065761007b565b806327304dfd1461008057806354fd4d50146100c757806355a5194b1461015157806387ae4ac21461018e575b600080fd5b34801561008c57600080fd5b506100b3600480360360208110156100a357600080fd5b50356001600160a01b0316610439565b604080519115158252519081900360200190f35b3480156100d357600080fd5b506100dc61044e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101165781810151838201526020016100fe565b50505050905090810190601f1680156101435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015d57600080fd5b5061018c6004803603604081101561017457600080fd5b506001600160a01b03813516906020013515156104db565b005b34801561019a57600080fd5b5061018c600480360360408110156101b157600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156101dc57600080fd5b8201836020820111156101ee57600080fd5b8035906020019184600183028401116401000000008311171561021057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061055e945050505050565b34801561025d57600080fd5b506103046004803603602081101561027457600080fd5b81019060208101813564010000000081111561028f57600080fd5b8201836020820111156102a157600080fd5b803590602001918460018302840111640100000000831117156102c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610634945050505050565b604080516001600160a01b039092168252519081900360200190f35b34801561032c57600080fd5b506100b36004803603602081101561034357600080fd5b50356001600160a01b031661065a565b34801561035f57600080fd5b506103046004803603602081101561037657600080fd5b81019060208101813564010000000081111561039157600080fd5b8201836020820111156103a357600080fd5b803590602001918460018302840111640100000000831117156103c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610678945050505050565b34801561041257600080fd5b5061018c6004803603602081101561042957600080fd5b50356001600160a01b03166106e9565b60026020526000908152604090205460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104d35780601f106104a8576101008083540402835291602001916104d3565b820191906000526020600020905b8154815290600101906020018083116104b657829003601f168201915b505050505081565b6000546001600160a01b03163314610533576040805162461bcd60e51b81526020600482015260166024820152756f6e6c79206f776e65722063616e20646f206974202160501b604482015290519081900360640190fd5b6001600160a01b03919091166000908152600260205260409020805460ff1916911515919091179055565b6105673361065a565b6105b1576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e20646f206974202160501b604482015290519081900360640190fd5b816003826040518082805190602001908083835b602083106105e45780518252601f1990920191602091820191016105c5565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160a01b0319166001600160a01b03949094169390931790925550505050565b80516020818301810180516003825292820191909301209152546001600160a01b031681565b6001600160a01b031660009081526002602052604090205460ff1690565b60006003826040518082805190602001908083835b602083106106ac5780518252601f19909201916020918201910161068d565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b0316949350505050565b6000546001600160a01b03163314610741576040805162461bcd60e51b81526020600482015260166024820152756f6e6c79206f776e65722063616e20646f206974202160501b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820251849b89ccba4294027156a46f768976a8384e93480e98553f22808efe42d7664736f6c637827302e352e31362d646576656c6f702e323032302e332e352b636f6d6d69742e39633332323663650057"

// DeployRchMgrInterface deploys a new Ethereum contract, binding an instance of RchMgrInterface to it.
func DeployRchMgrInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RchMgrInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(RchMgrInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RchMgrInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RchMgrInterface{RchMgrInterfaceCaller: RchMgrInterfaceCaller{contract: contract}, RchMgrInterfaceTransactor: RchMgrInterfaceTransactor{contract: contract}, RchMgrInterfaceFilterer: RchMgrInterfaceFilterer{contract: contract}}, nil
}

// RchMgrInterface is an auto generated Go binding around an Ethereum contract.
type RchMgrInterface struct {
	RchMgrInterfaceCaller     // Read-only binding to the contract
	RchMgrInterfaceTransactor // Write-only binding to the contract
	RchMgrInterfaceFilterer   // Log filterer for contract events
}

// RchMgrInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type RchMgrInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RchMgrInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RchMgrInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RchMgrInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RchMgrInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RchMgrInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RchMgrInterfaceSession struct {
	Contract     *RchMgrInterface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RchMgrInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RchMgrInterfaceCallerSession struct {
	Contract *RchMgrInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RchMgrInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RchMgrInterfaceTransactorSession struct {
	Contract     *RchMgrInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RchMgrInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type RchMgrInterfaceRaw struct {
	Contract *RchMgrInterface // Generic contract binding to access the raw methods on
}

// RchMgrInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RchMgrInterfaceCallerRaw struct {
	Contract *RchMgrInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// RchMgrInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RchMgrInterfaceTransactorRaw struct {
	Contract *RchMgrInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRchMgrInterface creates a new instance of RchMgrInterface, bound to a specific deployed contract.
func NewRchMgrInterface(address common.Address, backend bind.ContractBackend) (*RchMgrInterface, error) {
	contract, err := bindRchMgrInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RchMgrInterface{RchMgrInterfaceCaller: RchMgrInterfaceCaller{contract: contract}, RchMgrInterfaceTransactor: RchMgrInterfaceTransactor{contract: contract}, RchMgrInterfaceFilterer: RchMgrInterfaceFilterer{contract: contract}}, nil
}

// NewRchMgrInterfaceCaller creates a new read-only instance of RchMgrInterface, bound to a specific deployed contract.
func NewRchMgrInterfaceCaller(address common.Address, caller bind.ContractCaller) (*RchMgrInterfaceCaller, error) {
	contract, err := bindRchMgrInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RchMgrInterfaceCaller{contract: contract}, nil
}

// NewRchMgrInterfaceTransactor creates a new write-only instance of RchMgrInterface, bound to a specific deployed contract.
func NewRchMgrInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*RchMgrInterfaceTransactor, error) {
	contract, err := bindRchMgrInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RchMgrInterfaceTransactor{contract: contract}, nil
}

// NewRchMgrInterfaceFilterer creates a new log filterer instance of RchMgrInterface, bound to a specific deployed contract.
func NewRchMgrInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*RchMgrInterfaceFilterer, error) {
	contract, err := bindRchMgrInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RchMgrInterfaceFilterer{contract: contract}, nil
}

// bindRchMgrInterface binds a generic wrapper to an already deployed contract.
func bindRchMgrInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RchMgrInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RchMgrInterface *RchMgrInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RchMgrInterface.Contract.RchMgrInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RchMgrInterface *RchMgrInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.RchMgrInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RchMgrInterface *RchMgrInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.RchMgrInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RchMgrInterface *RchMgrInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RchMgrInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RchMgrInterface *RchMgrInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RchMgrInterface *RchMgrInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.contract.Transact(opts, method, params...)
}

// GetContractAddr is a free data retrieval call binding the contract method 0xd4527406.
//
// Solidity: function GetContractAddr(string name) constant returns(address ret)
func (_RchMgrInterface *RchMgrInterfaceCaller) GetContractAddr(opts *bind.CallOpts, name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RchMgrInterface.contract.Call(opts, out, "GetContractAddr", name)
	return *ret0, err
}

// GetContractAddr is a free data retrieval call binding the contract method 0xd4527406.
//
// Solidity: function GetContractAddr(string name) constant returns(address ret)
func (_RchMgrInterface *RchMgrInterfaceSession) GetContractAddr(name string) (common.Address, error) {
	return _RchMgrInterface.Contract.GetContractAddr(&_RchMgrInterface.CallOpts, name)
}

// GetContractAddr is a free data retrieval call binding the contract method 0xd4527406.
//
// Solidity: function GetContractAddr(string name) constant returns(address ret)
func (_RchMgrInterface *RchMgrInterfaceCallerSession) GetContractAddr(name string) (common.Address, error) {
	return _RchMgrInterface.Contract.GetContractAddr(&_RchMgrInterface.CallOpts, name)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb11e3c1a.
//
// Solidity: function IsAdmin(address addr) constant returns(bool ret)
func (_RchMgrInterface *RchMgrInterfaceCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RchMgrInterface.contract.Call(opts, out, "IsAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0xb11e3c1a.
//
// Solidity: function IsAdmin(address addr) constant returns(bool ret)
func (_RchMgrInterface *RchMgrInterfaceSession) IsAdmin(addr common.Address) (bool, error) {
	return _RchMgrInterface.Contract.IsAdmin(&_RchMgrInterface.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb11e3c1a.
//
// Solidity: function IsAdmin(address addr) constant returns(bool ret)
func (_RchMgrInterface *RchMgrInterfaceCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _RchMgrInterface.Contract.IsAdmin(&_RchMgrInterface.CallOpts, addr)
}

// AdminList is a free data retrieval call binding the contract method 0x27304dfd.
//
// Solidity: function adminList(address ) constant returns(bool)
func (_RchMgrInterface *RchMgrInterfaceCaller) AdminList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RchMgrInterface.contract.Call(opts, out, "adminList", arg0)
	return *ret0, err
}

// AdminList is a free data retrieval call binding the contract method 0x27304dfd.
//
// Solidity: function adminList(address ) constant returns(bool)
func (_RchMgrInterface *RchMgrInterfaceSession) AdminList(arg0 common.Address) (bool, error) {
	return _RchMgrInterface.Contract.AdminList(&_RchMgrInterface.CallOpts, arg0)
}

// AdminList is a free data retrieval call binding the contract method 0x27304dfd.
//
// Solidity: function adminList(address ) constant returns(bool)
func (_RchMgrInterface *RchMgrInterfaceCallerSession) AdminList(arg0 common.Address) (bool, error) {
	return _RchMgrInterface.Contract.AdminList(&_RchMgrInterface.CallOpts, arg0)
}

// ContractAddrs is a free data retrieval call binding the contract method 0x98ef5e3e.
//
// Solidity: function contractAddrs(string ) constant returns(address)
func (_RchMgrInterface *RchMgrInterfaceCaller) ContractAddrs(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RchMgrInterface.contract.Call(opts, out, "contractAddrs", arg0)
	return *ret0, err
}

// ContractAddrs is a free data retrieval call binding the contract method 0x98ef5e3e.
//
// Solidity: function contractAddrs(string ) constant returns(address)
func (_RchMgrInterface *RchMgrInterfaceSession) ContractAddrs(arg0 string) (common.Address, error) {
	return _RchMgrInterface.Contract.ContractAddrs(&_RchMgrInterface.CallOpts, arg0)
}

// ContractAddrs is a free data retrieval call binding the contract method 0x98ef5e3e.
//
// Solidity: function contractAddrs(string ) constant returns(address)
func (_RchMgrInterface *RchMgrInterfaceCallerSession) ContractAddrs(arg0 string) (common.Address, error) {
	return _RchMgrInterface.Contract.ContractAddrs(&_RchMgrInterface.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_RchMgrInterface *RchMgrInterfaceCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RchMgrInterface.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_RchMgrInterface *RchMgrInterfaceSession) Version() (string, error) {
	return _RchMgrInterface.Contract.Version(&_RchMgrInterface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_RchMgrInterface *RchMgrInterfaceCallerSession) Version() (string, error) {
	return _RchMgrInterface.Contract.Version(&_RchMgrInterface.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf2853292.
//
// Solidity: function ChangeOwner(address owner2) returns()
func (_RchMgrInterface *RchMgrInterfaceTransactor) ChangeOwner(opts *bind.TransactOpts, owner2 common.Address) (*types.Transaction, error) {
	return _RchMgrInterface.contract.Transact(opts, "ChangeOwner", owner2)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf2853292.
//
// Solidity: function ChangeOwner(address owner2) returns()
func (_RchMgrInterface *RchMgrInterfaceSession) ChangeOwner(owner2 common.Address) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.ChangeOwner(&_RchMgrInterface.TransactOpts, owner2)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf2853292.
//
// Solidity: function ChangeOwner(address owner2) returns()
func (_RchMgrInterface *RchMgrInterfaceTransactorSession) ChangeOwner(owner2 common.Address) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.ChangeOwner(&_RchMgrInterface.TransactOpts, owner2)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x55a5194b.
//
// Solidity: function SetAdmin(address addr, bool flag) returns()
func (_RchMgrInterface *RchMgrInterfaceTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, flag bool) (*types.Transaction, error) {
	return _RchMgrInterface.contract.Transact(opts, "SetAdmin", addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x55a5194b.
//
// Solidity: function SetAdmin(address addr, bool flag) returns()
func (_RchMgrInterface *RchMgrInterfaceSession) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.SetAdmin(&_RchMgrInterface.TransactOpts, addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x55a5194b.
//
// Solidity: function SetAdmin(address addr, bool flag) returns()
func (_RchMgrInterface *RchMgrInterfaceTransactorSession) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.SetAdmin(&_RchMgrInterface.TransactOpts, addr, flag)
}

// SetContract is a paid mutator transaction binding the contract method 0x87ae4ac2.
//
// Solidity: function SetContract(address contractAddr, string name) returns()
func (_RchMgrInterface *RchMgrInterfaceTransactor) SetContract(opts *bind.TransactOpts, contractAddr common.Address, name string) (*types.Transaction, error) {
	return _RchMgrInterface.contract.Transact(opts, "SetContract", contractAddr, name)
}

// SetContract is a paid mutator transaction binding the contract method 0x87ae4ac2.
//
// Solidity: function SetContract(address contractAddr, string name) returns()
func (_RchMgrInterface *RchMgrInterfaceSession) SetContract(contractAddr common.Address, name string) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.SetContract(&_RchMgrInterface.TransactOpts, contractAddr, name)
}

// SetContract is a paid mutator transaction binding the contract method 0x87ae4ac2.
//
// Solidity: function SetContract(address contractAddr, string name) returns()
func (_RchMgrInterface *RchMgrInterfaceTransactorSession) SetContract(contractAddr common.Address, name string) (*types.Transaction, error) {
	return _RchMgrInterface.Contract.SetContract(&_RchMgrInterface.TransactOpts, contractAddr, name)
}

// RchRegPoolMgrABI is the input ABI used to generate the binding from.
const RchRegPoolMgrABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RchTransfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner2\",\"type\":\"address\"}],\"name\":\"ChangeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetStakeGateValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ret\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ret\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"RchStakePoolValue\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"RchUnStakePoolValue\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"SetAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"}],\"name\":\"SetStakeTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolAddrList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolStakeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userBalane\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"unstakeing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"unstakeGate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unstake_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RchRegPoolMgrFuncSigs maps the 4-byte function signature to its string representation.
var RchRegPoolMgrFuncSigs = map[string]string{
	"f2853292": "ChangeOwner(address)",
	"2617493e": "GetStakeGateValue()",
	"b11e3c1a": "IsAdmin(address)",
	"28dbe185": "RchStakePoolValue()",
	"f443e244": "RchUnStakePoolValue()",
	"55a5194b": "SetAdmin(address,bool)",
	"8694f12b": "SetStakeTime(uint256)",
	"27304dfd": "adminList(address)",
	"41c0e1b5": "kill()",
	"009f6a0b": "poolAddrList(uint256)",
	"88cbacf9": "poolListCount()",
	"974abfda": "poolStakeList(address)",
	"e0928bdb": "unstake_time()",
	"54fd4d50": "version()",
}

// RchRegPoolMgrBin is the compiled bytecode used for deploying new contracts.
var RchRegPoolMgrBin = "0x608060405262278d0060065534801561001757600080fd5b50600080546001600160a01b031916331790556040805180820190915260198082527f526368526567506f6f6c4d67722076657273696f6e20312e3500000000000000602090920191825261006e91600191610097565b50600080546001600160a01b03168152600360205260409020805460ff19166001179055610132565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100d857805160ff1916838001178555610105565b82800160010185558215610105579182015b828111156101055782518255916020019190600101906100ea565b50610111929150610115565b5090565b61012f91905b80821115610111576000815560010161011b565b90565b6109bb806101416000396000f3fe6080604052600436106100dc5760003560e01c80638694f12b1161007f578063b11e3c1a11610059578063b11e3c1a14610315578063e0928bdb14610348578063f28532921461035d578063f443e24414610390576100dc565b80638694f12b1461026c57806388cbacf914610296578063974abfda146102ab576100dc565b806328dbe185116100bb57806328dbe1851461019557806341c0e1b51461019f57806354fd4d50146101a757806355a5194b14610231576100dc565b80629f6a0b146100e15780632617493e1461012757806327304dfd1461014e575b600080fd5b3480156100ed57600080fd5b5061010b6004803603602081101561010457600080fd5b5035610398565b604080516001600160a01b039092168252519081900360200190f35b34801561013357600080fd5b5061013c6103bf565b60408051918252519081900360200190f35b34801561015a57600080fd5b506101816004803603602081101561017157600080fd5b50356001600160a01b03166103cd565b604080519115158252519081900360200190f35b61019d6103e2565b005b61019d61057e565b3480156101b357600080fd5b506101bc6105aa565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f65781810151838201526020016101de565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023d57600080fd5b5061019d6004803603604081101561025457600080fd5b506001600160a01b0381351690602001351515610637565b34801561027857600080fd5b5061019d6004803603602081101561028f57600080fd5b50356106ba565b3480156102a257600080fd5b5061013c610762565b3480156102b757600080fd5b506102de600480360360208110156102ce57600080fd5b50356001600160a01b0316610768565b604080516001600160a01b039096168652602086019490945284840192909252151560608401526080830152519081900360a00190f35b34801561032157600080fd5b506101816004803603602081101561033857600080fd5b50356001600160a01b03166107a7565b34801561035457600080fd5b5061013c6107c5565b34801561036957600080fd5b5061019d6004803603602081101561038057600080fd5b50356001600160a01b03166107cb565b61019d610845565b600481815481106103a557fe5b6000918252602090912001546001600160a01b0316905081565b69021e19e0c9bab240000090565b60036020526000908152604090205460ff1681565b6103ea6103bf565b34101561043e576040805162461bcd60e51b815260206004820152601f60248201527f6e65656420313030302072636820666f722063726561746520746f6b656e2e00604482015290519081900360640190fd5b3360009081526002602052604090206003015460ff16156104a6576040805162461bcd60e51b815260206004820152601f60248201527f6e65656420313030302072636820666f722063726561746520746f6b656e2e00604482015290519081900360640190fd5b336000818152600260208190526040822080546001600160a01b03191690931783556001830180543401905542920191909155805b60045481101561052857336001600160a01b0316600482815481106104fc57fe5b6000918252602090912001546001600160a01b031614156105205760019150610528565b6001016104db565b508061057b5760048054600181810183556000929092527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b031916331790556005805490910190555b50565b6000546001600160a01b03163314156105a857737a620f1ad467487dd6b6353224f3ff7f5771a248ff5b565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561062f5780601f106106045761010080835404028352916020019161062f565b820191906000526020600020905b81548152906001019060200180831161061257829003601f168201915b505050505081565b6000546001600160a01b0316331461068f576040805162461bcd60e51b81526020600482015260166024820152756f6e6c79206f776e65722063616e20646f206974202160501b604482015290519081900360640190fd5b6001600160a01b03919091166000908152600360205260409020805460ff1916911515919091179055565b6106c3336107a7565b61070d576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e20646f206974202160501b604482015290519081900360640190fd5b61025881101561075d576040805162461bcd60e51b8152602060048201526016602482015275030ba103632b0b9ba101b18181039b2b1b7b73232b9960551b604482015290519081900360640190fd5b600655565b60055481565b6002602081905260009182526040909120805460018201549282015460038301546004909301546001600160a01b039092169392909160ff9091169085565b6001600160a01b031660009081526003602052604090205460ff1690565b60065481565b6000546001600160a01b03163314610823576040805162461bcd60e51b81526020600482015260166024820152756f6e6c79206f776e65722063616e20646f206974202160501b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b33600090815260026020526040902060010154156105a85733600090815260026020819052600654604090922001544203116108c8576040805162461bcd60e51b815260206004820181905260248201527f49742077696c6c2074616b65203330206461797320746f20756e7374616b652e604482015290519081900360640190fd5b3360008181526002602052604080822060010154905181156108fc0292818181858888f19350505050158015610902573d6000803e3d6000fd5b503360008181526002602052604080822060010154905190929130917f9a2833e7502e6ea7dc729986eef8d1f258c15585c55395efe46ecf2f9010652d9190a4336000908152600260208190526040822060018101929092554291015556fea265627a7a72315820ef94e158b2416bfa1a12ef56aca67d1d277863a1dcd054a3fa9f25f81cf06b5464736f6c637827302e352e31362d646576656c6f702e323032302e332e352b636f6d6d69742e39633332323663650057"

// DeployRchRegPoolMgr deploys a new Ethereum contract, binding an instance of RchRegPoolMgr to it.
func DeployRchRegPoolMgr(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RchRegPoolMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(RchRegPoolMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RchRegPoolMgrBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RchRegPoolMgr{RchRegPoolMgrCaller: RchRegPoolMgrCaller{contract: contract}, RchRegPoolMgrTransactor: RchRegPoolMgrTransactor{contract: contract}, RchRegPoolMgrFilterer: RchRegPoolMgrFilterer{contract: contract}}, nil
}

// RchRegPoolMgr is an auto generated Go binding around an Ethereum contract.
type RchRegPoolMgr struct {
	RchRegPoolMgrCaller     // Read-only binding to the contract
	RchRegPoolMgrTransactor // Write-only binding to the contract
	RchRegPoolMgrFilterer   // Log filterer for contract events
}

// RchRegPoolMgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type RchRegPoolMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RchRegPoolMgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RchRegPoolMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RchRegPoolMgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RchRegPoolMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RchRegPoolMgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RchRegPoolMgrSession struct {
	Contract     *RchRegPoolMgr    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RchRegPoolMgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RchRegPoolMgrCallerSession struct {
	Contract *RchRegPoolMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RchRegPoolMgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RchRegPoolMgrTransactorSession struct {
	Contract     *RchRegPoolMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RchRegPoolMgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type RchRegPoolMgrRaw struct {
	Contract *RchRegPoolMgr // Generic contract binding to access the raw methods on
}

// RchRegPoolMgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RchRegPoolMgrCallerRaw struct {
	Contract *RchRegPoolMgrCaller // Generic read-only contract binding to access the raw methods on
}

// RchRegPoolMgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RchRegPoolMgrTransactorRaw struct {
	Contract *RchRegPoolMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRchRegPoolMgr creates a new instance of RchRegPoolMgr, bound to a specific deployed contract.
func NewRchRegPoolMgr(address common.Address, backend bind.ContractBackend) (*RchRegPoolMgr, error) {
	contract, err := bindRchRegPoolMgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RchRegPoolMgr{RchRegPoolMgrCaller: RchRegPoolMgrCaller{contract: contract}, RchRegPoolMgrTransactor: RchRegPoolMgrTransactor{contract: contract}, RchRegPoolMgrFilterer: RchRegPoolMgrFilterer{contract: contract}}, nil
}

// NewRchRegPoolMgrCaller creates a new read-only instance of RchRegPoolMgr, bound to a specific deployed contract.
func NewRchRegPoolMgrCaller(address common.Address, caller bind.ContractCaller) (*RchRegPoolMgrCaller, error) {
	contract, err := bindRchRegPoolMgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RchRegPoolMgrCaller{contract: contract}, nil
}

// NewRchRegPoolMgrTransactor creates a new write-only instance of RchRegPoolMgr, bound to a specific deployed contract.
func NewRchRegPoolMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*RchRegPoolMgrTransactor, error) {
	contract, err := bindRchRegPoolMgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RchRegPoolMgrTransactor{contract: contract}, nil
}

// NewRchRegPoolMgrFilterer creates a new log filterer instance of RchRegPoolMgr, bound to a specific deployed contract.
func NewRchRegPoolMgrFilterer(address common.Address, filterer bind.ContractFilterer) (*RchRegPoolMgrFilterer, error) {
	contract, err := bindRchRegPoolMgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RchRegPoolMgrFilterer{contract: contract}, nil
}

// bindRchRegPoolMgr binds a generic wrapper to an already deployed contract.
func bindRchRegPoolMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RchRegPoolMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RchRegPoolMgr *RchRegPoolMgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RchRegPoolMgr.Contract.RchRegPoolMgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RchRegPoolMgr *RchRegPoolMgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.RchRegPoolMgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RchRegPoolMgr *RchRegPoolMgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.RchRegPoolMgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RchRegPoolMgr *RchRegPoolMgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RchRegPoolMgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RchRegPoolMgr *RchRegPoolMgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RchRegPoolMgr *RchRegPoolMgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.contract.Transact(opts, method, params...)
}

// GetStakeGateValue is a free data retrieval call binding the contract method 0x2617493e.
//
// Solidity: function GetStakeGateValue() constant returns(uint256 ret)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) GetStakeGateValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "GetStakeGateValue")
	return *ret0, err
}

// GetStakeGateValue is a free data retrieval call binding the contract method 0x2617493e.
//
// Solidity: function GetStakeGateValue() constant returns(uint256 ret)
func (_RchRegPoolMgr *RchRegPoolMgrSession) GetStakeGateValue() (*big.Int, error) {
	return _RchRegPoolMgr.Contract.GetStakeGateValue(&_RchRegPoolMgr.CallOpts)
}

// GetStakeGateValue is a free data retrieval call binding the contract method 0x2617493e.
//
// Solidity: function GetStakeGateValue() constant returns(uint256 ret)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) GetStakeGateValue() (*big.Int, error) {
	return _RchRegPoolMgr.Contract.GetStakeGateValue(&_RchRegPoolMgr.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb11e3c1a.
//
// Solidity: function IsAdmin(address addr) constant returns(bool ret)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "IsAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0xb11e3c1a.
//
// Solidity: function IsAdmin(address addr) constant returns(bool ret)
func (_RchRegPoolMgr *RchRegPoolMgrSession) IsAdmin(addr common.Address) (bool, error) {
	return _RchRegPoolMgr.Contract.IsAdmin(&_RchRegPoolMgr.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb11e3c1a.
//
// Solidity: function IsAdmin(address addr) constant returns(bool ret)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _RchRegPoolMgr.Contract.IsAdmin(&_RchRegPoolMgr.CallOpts, addr)
}

// AdminList is a free data retrieval call binding the contract method 0x27304dfd.
//
// Solidity: function adminList(address ) constant returns(bool)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) AdminList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "adminList", arg0)
	return *ret0, err
}

// AdminList is a free data retrieval call binding the contract method 0x27304dfd.
//
// Solidity: function adminList(address ) constant returns(bool)
func (_RchRegPoolMgr *RchRegPoolMgrSession) AdminList(arg0 common.Address) (bool, error) {
	return _RchRegPoolMgr.Contract.AdminList(&_RchRegPoolMgr.CallOpts, arg0)
}

// AdminList is a free data retrieval call binding the contract method 0x27304dfd.
//
// Solidity: function adminList(address ) constant returns(bool)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) AdminList(arg0 common.Address) (bool, error) {
	return _RchRegPoolMgr.Contract.AdminList(&_RchRegPoolMgr.CallOpts, arg0)
}

// PoolAddrList is a free data retrieval call binding the contract method 0x009f6a0b.
//
// Solidity: function poolAddrList(uint256 ) constant returns(address)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) PoolAddrList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "poolAddrList", arg0)
	return *ret0, err
}

// PoolAddrList is a free data retrieval call binding the contract method 0x009f6a0b.
//
// Solidity: function poolAddrList(uint256 ) constant returns(address)
func (_RchRegPoolMgr *RchRegPoolMgrSession) PoolAddrList(arg0 *big.Int) (common.Address, error) {
	return _RchRegPoolMgr.Contract.PoolAddrList(&_RchRegPoolMgr.CallOpts, arg0)
}

// PoolAddrList is a free data retrieval call binding the contract method 0x009f6a0b.
//
// Solidity: function poolAddrList(uint256 ) constant returns(address)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) PoolAddrList(arg0 *big.Int) (common.Address, error) {
	return _RchRegPoolMgr.Contract.PoolAddrList(&_RchRegPoolMgr.CallOpts, arg0)
}

// PoolListCount is a free data retrieval call binding the contract method 0x88cbacf9.
//
// Solidity: function poolListCount() constant returns(uint256)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) PoolListCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "poolListCount")
	return *ret0, err
}

// PoolListCount is a free data retrieval call binding the contract method 0x88cbacf9.
//
// Solidity: function poolListCount() constant returns(uint256)
func (_RchRegPoolMgr *RchRegPoolMgrSession) PoolListCount() (*big.Int, error) {
	return _RchRegPoolMgr.Contract.PoolListCount(&_RchRegPoolMgr.CallOpts)
}

// PoolListCount is a free data retrieval call binding the contract method 0x88cbacf9.
//
// Solidity: function poolListCount() constant returns(uint256)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) PoolListCount() (*big.Int, error) {
	return _RchRegPoolMgr.Contract.PoolListCount(&_RchRegPoolMgr.CallOpts)
}

// PoolStakeList is a free data retrieval call binding the contract method 0x974abfda.
//
// Solidity: function poolStakeList(address ) constant returns(address userAddr, uint256 userBalane, uint256 updateTime, bool unstakeing, uint256 unstakeGate)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) PoolStakeList(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserAddr    common.Address
	UserBalane  *big.Int
	UpdateTime  *big.Int
	Unstakeing  bool
	UnstakeGate *big.Int
}, error) {
	ret := new(struct {
		UserAddr    common.Address
		UserBalane  *big.Int
		UpdateTime  *big.Int
		Unstakeing  bool
		UnstakeGate *big.Int
	})
	out := ret
	err := _RchRegPoolMgr.contract.Call(opts, out, "poolStakeList", arg0)
	return *ret, err
}

// PoolStakeList is a free data retrieval call binding the contract method 0x974abfda.
//
// Solidity: function poolStakeList(address ) constant returns(address userAddr, uint256 userBalane, uint256 updateTime, bool unstakeing, uint256 unstakeGate)
func (_RchRegPoolMgr *RchRegPoolMgrSession) PoolStakeList(arg0 common.Address) (struct {
	UserAddr    common.Address
	UserBalane  *big.Int
	UpdateTime  *big.Int
	Unstakeing  bool
	UnstakeGate *big.Int
}, error) {
	return _RchRegPoolMgr.Contract.PoolStakeList(&_RchRegPoolMgr.CallOpts, arg0)
}

// PoolStakeList is a free data retrieval call binding the contract method 0x974abfda.
//
// Solidity: function poolStakeList(address ) constant returns(address userAddr, uint256 userBalane, uint256 updateTime, bool unstakeing, uint256 unstakeGate)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) PoolStakeList(arg0 common.Address) (struct {
	UserAddr    common.Address
	UserBalane  *big.Int
	UpdateTime  *big.Int
	Unstakeing  bool
	UnstakeGate *big.Int
}, error) {
	return _RchRegPoolMgr.Contract.PoolStakeList(&_RchRegPoolMgr.CallOpts, arg0)
}

// UnstakeTime is a free data retrieval call binding the contract method 0xe0928bdb.
//
// Solidity: function unstake_time() constant returns(uint256)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) UnstakeTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "unstake_time")
	return *ret0, err
}

// UnstakeTime is a free data retrieval call binding the contract method 0xe0928bdb.
//
// Solidity: function unstake_time() constant returns(uint256)
func (_RchRegPoolMgr *RchRegPoolMgrSession) UnstakeTime() (*big.Int, error) {
	return _RchRegPoolMgr.Contract.UnstakeTime(&_RchRegPoolMgr.CallOpts)
}

// UnstakeTime is a free data retrieval call binding the contract method 0xe0928bdb.
//
// Solidity: function unstake_time() constant returns(uint256)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) UnstakeTime() (*big.Int, error) {
	return _RchRegPoolMgr.Contract.UnstakeTime(&_RchRegPoolMgr.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_RchRegPoolMgr *RchRegPoolMgrCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RchRegPoolMgr.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_RchRegPoolMgr *RchRegPoolMgrSession) Version() (string, error) {
	return _RchRegPoolMgr.Contract.Version(&_RchRegPoolMgr.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_RchRegPoolMgr *RchRegPoolMgrCallerSession) Version() (string, error) {
	return _RchRegPoolMgr.Contract.Version(&_RchRegPoolMgr.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf2853292.
//
// Solidity: function ChangeOwner(address owner2) returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactor) ChangeOwner(opts *bind.TransactOpts, owner2 common.Address) (*types.Transaction, error) {
	return _RchRegPoolMgr.contract.Transact(opts, "ChangeOwner", owner2)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf2853292.
//
// Solidity: function ChangeOwner(address owner2) returns()
func (_RchRegPoolMgr *RchRegPoolMgrSession) ChangeOwner(owner2 common.Address) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.ChangeOwner(&_RchRegPoolMgr.TransactOpts, owner2)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf2853292.
//
// Solidity: function ChangeOwner(address owner2) returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactorSession) ChangeOwner(owner2 common.Address) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.ChangeOwner(&_RchRegPoolMgr.TransactOpts, owner2)
}

// RchStakePoolValue is a paid mutator transaction binding the contract method 0x28dbe185.
//
// Solidity: function RchStakePoolValue() returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactor) RchStakePoolValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchRegPoolMgr.contract.Transact(opts, "RchStakePoolValue")
}

// RchStakePoolValue is a paid mutator transaction binding the contract method 0x28dbe185.
//
// Solidity: function RchStakePoolValue() returns()
func (_RchRegPoolMgr *RchRegPoolMgrSession) RchStakePoolValue() (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.RchStakePoolValue(&_RchRegPoolMgr.TransactOpts)
}

// RchStakePoolValue is a paid mutator transaction binding the contract method 0x28dbe185.
//
// Solidity: function RchStakePoolValue() returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactorSession) RchStakePoolValue() (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.RchStakePoolValue(&_RchRegPoolMgr.TransactOpts)
}

// RchUnStakePoolValue is a paid mutator transaction binding the contract method 0xf443e244.
//
// Solidity: function RchUnStakePoolValue() returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactor) RchUnStakePoolValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchRegPoolMgr.contract.Transact(opts, "RchUnStakePoolValue")
}

// RchUnStakePoolValue is a paid mutator transaction binding the contract method 0xf443e244.
//
// Solidity: function RchUnStakePoolValue() returns()
func (_RchRegPoolMgr *RchRegPoolMgrSession) RchUnStakePoolValue() (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.RchUnStakePoolValue(&_RchRegPoolMgr.TransactOpts)
}

// RchUnStakePoolValue is a paid mutator transaction binding the contract method 0xf443e244.
//
// Solidity: function RchUnStakePoolValue() returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactorSession) RchUnStakePoolValue() (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.RchUnStakePoolValue(&_RchRegPoolMgr.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x55a5194b.
//
// Solidity: function SetAdmin(address addr, bool flag) returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, flag bool) (*types.Transaction, error) {
	return _RchRegPoolMgr.contract.Transact(opts, "SetAdmin", addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x55a5194b.
//
// Solidity: function SetAdmin(address addr, bool flag) returns()
func (_RchRegPoolMgr *RchRegPoolMgrSession) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.SetAdmin(&_RchRegPoolMgr.TransactOpts, addr, flag)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x55a5194b.
//
// Solidity: function SetAdmin(address addr, bool flag) returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactorSession) SetAdmin(addr common.Address, flag bool) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.SetAdmin(&_RchRegPoolMgr.TransactOpts, addr, flag)
}

// SetStakeTime is a paid mutator transaction binding the contract method 0x8694f12b.
//
// Solidity: function SetStakeTime(uint256 unstakeTime) returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactor) SetStakeTime(opts *bind.TransactOpts, unstakeTime *big.Int) (*types.Transaction, error) {
	return _RchRegPoolMgr.contract.Transact(opts, "SetStakeTime", unstakeTime)
}

// SetStakeTime is a paid mutator transaction binding the contract method 0x8694f12b.
//
// Solidity: function SetStakeTime(uint256 unstakeTime) returns()
func (_RchRegPoolMgr *RchRegPoolMgrSession) SetStakeTime(unstakeTime *big.Int) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.SetStakeTime(&_RchRegPoolMgr.TransactOpts, unstakeTime)
}

// SetStakeTime is a paid mutator transaction binding the contract method 0x8694f12b.
//
// Solidity: function SetStakeTime(uint256 unstakeTime) returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactorSession) SetStakeTime(unstakeTime *big.Int) (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.SetStakeTime(&_RchRegPoolMgr.TransactOpts, unstakeTime)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RchRegPoolMgr.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RchRegPoolMgr *RchRegPoolMgrSession) Kill() (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.Kill(&_RchRegPoolMgr.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RchRegPoolMgr *RchRegPoolMgrTransactorSession) Kill() (*types.Transaction, error) {
	return _RchRegPoolMgr.Contract.Kill(&_RchRegPoolMgr.TransactOpts)
}

// RchRegPoolMgrRchTransferIterator is returned from FilterRchTransfer and is used to iterate over the raw logs and unpacked data for RchTransfer events raised by the RchRegPoolMgr contract.
type RchRegPoolMgrRchTransferIterator struct {
	Event *RchRegPoolMgrRchTransfer // Event containing the contract specifics and raw log

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
func (it *RchRegPoolMgrRchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RchRegPoolMgrRchTransfer)
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
		it.Event = new(RchRegPoolMgrRchTransfer)
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
func (it *RchRegPoolMgrRchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RchRegPoolMgrRchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RchRegPoolMgrRchTransfer represents a RchTransfer event raised by the RchRegPoolMgr contract.
type RchRegPoolMgrRchTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRchTransfer is a free log retrieval operation binding the contract event 0x9a2833e7502e6ea7dc729986eef8d1f258c15585c55395efe46ecf2f9010652d.
//
// Solidity: event RchTransfer(address indexed _from, address indexed _to, uint256 indexed value)
func (_RchRegPoolMgr *RchRegPoolMgrFilterer) FilterRchTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, value []*big.Int) (*RchRegPoolMgrRchTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _RchRegPoolMgr.contract.FilterLogs(opts, "RchTransfer", _fromRule, _toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &RchRegPoolMgrRchTransferIterator{contract: _RchRegPoolMgr.contract, event: "RchTransfer", logs: logs, sub: sub}, nil
}

// WatchRchTransfer is a free log subscription operation binding the contract event 0x9a2833e7502e6ea7dc729986eef8d1f258c15585c55395efe46ecf2f9010652d.
//
// Solidity: event RchTransfer(address indexed _from, address indexed _to, uint256 indexed value)
func (_RchRegPoolMgr *RchRegPoolMgrFilterer) WatchRchTransfer(opts *bind.WatchOpts, sink chan<- *RchRegPoolMgrRchTransfer, _from []common.Address, _to []common.Address, value []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _RchRegPoolMgr.contract.WatchLogs(opts, "RchTransfer", _fromRule, _toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RchRegPoolMgrRchTransfer)
				if err := _RchRegPoolMgr.contract.UnpackLog(event, "RchTransfer", log); err != nil {
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

// ParseRchTransfer is a log parse operation binding the contract event 0x9a2833e7502e6ea7dc729986eef8d1f258c15585c55395efe46ecf2f9010652d.
//
// Solidity: event RchTransfer(address indexed _from, address indexed _to, uint256 indexed value)
func (_RchRegPoolMgr *RchRegPoolMgrFilterer) ParseRchTransfer(log types.Log) (*RchRegPoolMgrRchTransfer, error) {
	event := new(RchRegPoolMgrRchTransfer)
	if err := _RchRegPoolMgr.contract.UnpackLog(event, "RchTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582012c8f2cdaeb4b805e6003e1f9f33c9527a6a6c1cb1d185c40f8a3b4edb4f101f64736f6c637827302e352e31362d646576656c6f702e323032302e332e352b636f6d6d69742e39633332323663650057"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TokenFuncSigs maps the 4-byte function signature to its string representation.
var TokenFuncSigs = map[string]string{
	"a9059cbb": "transfer(address,uint256)",
}

// TokenBin is the compiled bytecode used for deploying new contracts.
var TokenBin = "0x6080604052348015600f57600080fd5b5060b68061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063a9059cbb14602d575b600080fd5b605660048036036040811015604157600080fd5b506001600160a01b0381351690602001356058565b005b505056fea265627a7a72315820bab4095bc682907e9c6c9a858c9b62b8a6752026cbfc9b48adf22c55f43237d964736f6c637827302e352e31362d646576656c6f702e323032302e332e352b636f6d6d69742e39633332323663650057"

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) constant returns()
func (_Token *TokenCaller) Transfer(opts *bind.CallOpts, receiver common.Address, amount *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _Token.contract.Call(opts, out, "transfer", receiver, amount)
	return err
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) constant returns()
func (_Token *TokenSession) Transfer(receiver common.Address, amount *big.Int) error {
	return _Token.Contract.Transfer(&_Token.CallOpts, receiver, amount)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) constant returns()
func (_Token *TokenCallerSession) Transfer(receiver common.Address, amount *big.Int) error {
	return _Token.Contract.Transfer(&_Token.CallOpts, receiver, amount)
}
