// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arb

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

// ArbMetaData contains all meta data concerning the Arb contract.
var ArbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractARBID\",\"name\":\"_arbid\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_baseNode\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"ControllerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"ControllerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbid\",\"outputs\":[{\"internalType\":\"contractARBID\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"nameExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"registerOnly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArbABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbMetaData.ABI instead.
var ArbABI = ArbMetaData.ABI

// Arb is an auto generated Go binding around an Ethereum contract.
type Arb struct {
	ArbCaller     // Read-only binding to the contract
	ArbTransactor // Write-only binding to the contract
	ArbFilterer   // Log filterer for contract events
}

// ArbCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSession struct {
	Contract     *Arb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbCallerSession struct {
	Contract *ArbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbTransactorSession struct {
	Contract     *ArbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRaw struct {
	Contract *Arb // Generic contract binding to access the raw methods on
}

// ArbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbCallerRaw struct {
	Contract *ArbCaller // Generic read-only contract binding to access the raw methods on
}

// ArbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbTransactorRaw struct {
	Contract *ArbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArb creates a new instance of Arb, bound to a specific deployed contract.
func NewArb(address common.Address, backend bind.ContractBackend) (*Arb, error) {
	contract, err := bindArb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arb{ArbCaller: ArbCaller{contract: contract}, ArbTransactor: ArbTransactor{contract: contract}, ArbFilterer: ArbFilterer{contract: contract}}, nil
}

// NewArbCaller creates a new read-only instance of Arb, bound to a specific deployed contract.
func NewArbCaller(address common.Address, caller bind.ContractCaller) (*ArbCaller, error) {
	contract, err := bindArb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbCaller{contract: contract}, nil
}

// NewArbTransactor creates a new write-only instance of Arb, bound to a specific deployed contract.
func NewArbTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbTransactor, error) {
	contract, err := bindArb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTransactor{contract: contract}, nil
}

// NewArbFilterer creates a new log filterer instance of Arb, bound to a specific deployed contract.
func NewArbFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFilterer, error) {
	contract, err := bindArb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFilterer{contract: contract}, nil
}

// bindArb binds a generic wrapper to an already deployed contract.
func bindArb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.ArbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transact(opts, method, params...)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Arb *ArbCaller) GRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Arb *ArbSession) GRACEPERIOD() (*big.Int, error) {
	return _Arb.Contract.GRACEPERIOD(&_Arb.CallOpts)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Arb *ArbCallerSession) GRACEPERIOD() (*big.Int, error) {
	return _Arb.Contract.GRACEPERIOD(&_Arb.CallOpts)
}

// Arbid is a free data retrieval call binding the contract method 0x75dc1764.
//
// Solidity: function arbid() view returns(address)
func (_Arb *ArbCaller) Arbid(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "arbid")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbid is a free data retrieval call binding the contract method 0x75dc1764.
//
// Solidity: function arbid() view returns(address)
func (_Arb *ArbSession) Arbid() (common.Address, error) {
	return _Arb.Contract.Arbid(&_Arb.CallOpts)
}

// Arbid is a free data retrieval call binding the contract method 0x75dc1764.
//
// Solidity: function arbid() view returns(address)
func (_Arb *ArbCallerSession) Arbid() (common.Address, error) {
	return _Arb.Contract.Arbid(&_Arb.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Arb *ArbCaller) Available(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "available", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Arb *ArbSession) Available(id *big.Int) (bool, error) {
	return _Arb.Contract.Available(&_Arb.CallOpts, id)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Arb *ArbCallerSession) Available(id *big.Int) (bool, error) {
	return _Arb.Contract.Available(&_Arb.CallOpts, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Arb *ArbCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Arb *ArbSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Arb.Contract.BalanceOf(&_Arb.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Arb *ArbCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Arb.Contract.BalanceOf(&_Arb.CallOpts, owner)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Arb *ArbCaller) BaseNode(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "baseNode")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Arb *ArbSession) BaseNode() ([32]byte, error) {
	return _Arb.Contract.BaseNode(&_Arb.CallOpts)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Arb *ArbCallerSession) BaseNode() ([32]byte, error) {
	return _Arb.Contract.BaseNode(&_Arb.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Arb *ArbCaller) BaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Arb *ArbSession) BaseUri() (string, error) {
	return _Arb.Contract.BaseUri(&_Arb.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Arb *ArbCallerSession) BaseUri() (string, error) {
	return _Arb.Contract.BaseUri(&_Arb.CallOpts)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Arb *ArbCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Arb *ArbSession) Controllers(arg0 common.Address) (bool, error) {
	return _Arb.Contract.Controllers(&_Arb.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Arb *ArbCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _Arb.Contract.Controllers(&_Arb.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Arb *ArbCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Arb *ArbSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Arb.Contract.GetApproved(&_Arb.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Arb *ArbCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Arb.Contract.GetApproved(&_Arb.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Arb *ArbCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Arb *ArbSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Arb.Contract.IsApprovedForAll(&_Arb.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Arb *ArbCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Arb.Contract.IsApprovedForAll(&_Arb.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Arb *ArbCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Arb *ArbSession) Name() (string, error) {
	return _Arb.Contract.Name(&_Arb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Arb *ArbCallerSession) Name() (string, error) {
	return _Arb.Contract.Name(&_Arb.CallOpts)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Arb *ArbCaller) NameExpires(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "nameExpires", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Arb *ArbSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _Arb.Contract.NameExpires(&_Arb.CallOpts, id)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Arb *ArbCallerSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _Arb.Contract.NameExpires(&_Arb.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCallerSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Arb *ArbCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Arb *ArbSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Arb.Contract.OwnerOf(&_Arb.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Arb *ArbCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Arb.Contract.OwnerOf(&_Arb.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Arb *ArbCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Arb *ArbSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Arb.Contract.SupportsInterface(&_Arb.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Arb *ArbCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Arb.Contract.SupportsInterface(&_Arb.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Arb *ArbCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Arb *ArbSession) Symbol() (string, error) {
	return _Arb.Contract.Symbol(&_Arb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Arb *ArbCallerSession) Symbol() (string, error) {
	return _Arb.Contract.Symbol(&_Arb.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Arb *ArbCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Arb *ArbSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Arb.Contract.TokenURI(&_Arb.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Arb *ArbCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Arb.Contract.TokenURI(&_Arb.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Arb *ArbCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Arb *ArbSession) TotalSupply() (*big.Int, error) {
	return _Arb.Contract.TotalSupply(&_Arb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Arb *ArbCallerSession) TotalSupply() (*big.Int, error) {
	return _Arb.Contract.TotalSupply(&_Arb.CallOpts)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Arb *ArbTransactor) AddController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "addController", controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Arb *ArbSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _Arb.Contract.AddController(&_Arb.TransactOpts, controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Arb *ArbTransactorSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _Arb.Contract.AddController(&_Arb.TransactOpts, controller)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Arb *ArbTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Arb *ArbSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Approve(&_Arb.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Arb *ArbTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Approve(&_Arb.TransactOpts, to, tokenId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Arb *ArbTransactor) Reclaim(opts *bind.TransactOpts, id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "reclaim", id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Arb *ArbSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.Reclaim(&_Arb.TransactOpts, id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Arb *ArbTransactorSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.Reclaim(&_Arb.TransactOpts, id, owner)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Arb *ArbTransactor) Register(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "register", id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Arb *ArbSession) Register(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Register(&_Arb.TransactOpts, id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Arb *ArbTransactorSession) Register(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Register(&_Arb.TransactOpts, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Arb *ArbTransactor) RegisterOnly(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "registerOnly", id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Arb *ArbSession) RegisterOnly(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.RegisterOnly(&_Arb.TransactOpts, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Arb *ArbTransactorSession) RegisterOnly(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.RegisterOnly(&_Arb.TransactOpts, id, owner, duration)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Arb *ArbTransactor) RemoveController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "removeController", controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Arb *ArbSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _Arb.Contract.RemoveController(&_Arb.TransactOpts, controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Arb *ArbTransactorSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _Arb.Contract.RemoveController(&_Arb.TransactOpts, controller)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Arb *ArbTransactor) Renew(opts *bind.TransactOpts, id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "renew", id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Arb *ArbSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Renew(&_Arb.TransactOpts, id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Arb *ArbTransactorSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Renew(&_Arb.TransactOpts, id, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Arb *ArbTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Arb *ArbSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.SafeTransferFrom(&_Arb.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Arb *ArbTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.SafeTransferFrom(&_Arb.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Arb *ArbTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Arb *ArbSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arb.Contract.SafeTransferFrom0(&_Arb.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Arb *ArbTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arb.Contract.SafeTransferFrom0(&_Arb.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Arb *ArbTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Arb *ArbSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Arb.Contract.SetApprovalForAll(&_Arb.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Arb *ArbTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Arb.Contract.SetApprovalForAll(&_Arb.TransactOpts, operator, approved)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Arb *ArbTransactor) SetResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "setResolver", resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Arb *ArbSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Arb.Contract.SetResolver(&_Arb.TransactOpts, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Arb *ArbTransactorSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Arb.Contract.SetResolver(&_Arb.TransactOpts, resolver)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_Arb *ArbTransactor) SetURI(opts *bind.TransactOpts, newURI string) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "setURI", newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_Arb *ArbSession) SetURI(newURI string) (*types.Transaction, error) {
	return _Arb.Contract.SetURI(&_Arb.TransactOpts, newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_Arb *ArbTransactorSession) SetURI(newURI string) (*types.Transaction, error) {
	return _Arb.Contract.SetURI(&_Arb.TransactOpts, newURI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Arb *ArbTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Arb *ArbSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.TransferFrom(&_Arb.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Arb *ArbTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.TransferFrom(&_Arb.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// ArbApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Arb contract.
type ArbApprovalIterator struct {
	Event *ArbApproval // Event containing the contract specifics and raw log

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
func (it *ArbApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbApproval)
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
		it.Event = new(ArbApproval)
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
func (it *ArbApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbApproval represents a Approval event raised by the Arb contract.
type ArbApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Arb *ArbFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ArbApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbApprovalIterator{contract: _Arb.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Arb *ArbFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArbApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbApproval)
				if err := _Arb.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Arb *ArbFilterer) ParseApproval(log types.Log) (*ArbApproval, error) {
	event := new(ArbApproval)
	if err := _Arb.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Arb contract.
type ArbApprovalForAllIterator struct {
	Event *ArbApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArbApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbApprovalForAll)
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
		it.Event = new(ArbApprovalForAll)
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
func (it *ArbApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbApprovalForAll represents a ApprovalForAll event raised by the Arb contract.
type ArbApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Arb *ArbFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ArbApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArbApprovalForAllIterator{contract: _Arb.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Arb *ArbFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArbApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbApprovalForAll)
				if err := _Arb.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Arb *ArbFilterer) ParseApprovalForAll(log types.Log) (*ArbApprovalForAll, error) {
	event := new(ArbApprovalForAll)
	if err := _Arb.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbControllerAddedIterator is returned from FilterControllerAdded and is used to iterate over the raw logs and unpacked data for ControllerAdded events raised by the Arb contract.
type ArbControllerAddedIterator struct {
	Event *ArbControllerAdded // Event containing the contract specifics and raw log

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
func (it *ArbControllerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbControllerAdded)
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
		it.Event = new(ArbControllerAdded)
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
func (it *ArbControllerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbControllerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbControllerAdded represents a ControllerAdded event raised by the Arb contract.
type ArbControllerAdded struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerAdded is a free log retrieval operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Arb *ArbFilterer) FilterControllerAdded(opts *bind.FilterOpts, controller []common.Address) (*ArbControllerAddedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "ControllerAdded", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ArbControllerAddedIterator{contract: _Arb.contract, event: "ControllerAdded", logs: logs, sub: sub}, nil
}

// WatchControllerAdded is a free log subscription operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Arb *ArbFilterer) WatchControllerAdded(opts *bind.WatchOpts, sink chan<- *ArbControllerAdded, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "ControllerAdded", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbControllerAdded)
				if err := _Arb.contract.UnpackLog(event, "ControllerAdded", log); err != nil {
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

// ParseControllerAdded is a log parse operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Arb *ArbFilterer) ParseControllerAdded(log types.Log) (*ArbControllerAdded, error) {
	event := new(ArbControllerAdded)
	if err := _Arb.contract.UnpackLog(event, "ControllerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbControllerRemovedIterator is returned from FilterControllerRemoved and is used to iterate over the raw logs and unpacked data for ControllerRemoved events raised by the Arb contract.
type ArbControllerRemovedIterator struct {
	Event *ArbControllerRemoved // Event containing the contract specifics and raw log

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
func (it *ArbControllerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbControllerRemoved)
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
		it.Event = new(ArbControllerRemoved)
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
func (it *ArbControllerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbControllerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbControllerRemoved represents a ControllerRemoved event raised by the Arb contract.
type ArbControllerRemoved struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerRemoved is a free log retrieval operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Arb *ArbFilterer) FilterControllerRemoved(opts *bind.FilterOpts, controller []common.Address) (*ArbControllerRemovedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "ControllerRemoved", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ArbControllerRemovedIterator{contract: _Arb.contract, event: "ControllerRemoved", logs: logs, sub: sub}, nil
}

// WatchControllerRemoved is a free log subscription operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Arb *ArbFilterer) WatchControllerRemoved(opts *bind.WatchOpts, sink chan<- *ArbControllerRemoved, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "ControllerRemoved", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbControllerRemoved)
				if err := _Arb.contract.UnpackLog(event, "ControllerRemoved", log); err != nil {
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

// ParseControllerRemoved is a log parse operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Arb *ArbFilterer) ParseControllerRemoved(log types.Log) (*ArbControllerRemoved, error) {
	event := new(ArbControllerRemoved)
	if err := _Arb.contract.UnpackLog(event, "ControllerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbNameMigratedIterator is returned from FilterNameMigrated and is used to iterate over the raw logs and unpacked data for NameMigrated events raised by the Arb contract.
type ArbNameMigratedIterator struct {
	Event *ArbNameMigrated // Event containing the contract specifics and raw log

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
func (it *ArbNameMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbNameMigrated)
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
		it.Event = new(ArbNameMigrated)
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
func (it *ArbNameMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbNameMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbNameMigrated represents a NameMigrated event raised by the Arb contract.
type ArbNameMigrated struct {
	Id      *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameMigrated is a free log retrieval operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Arb *ArbFilterer) FilterNameMigrated(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*ArbNameMigratedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "NameMigrated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ArbNameMigratedIterator{contract: _Arb.contract, event: "NameMigrated", logs: logs, sub: sub}, nil
}

// WatchNameMigrated is a free log subscription operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Arb *ArbFilterer) WatchNameMigrated(opts *bind.WatchOpts, sink chan<- *ArbNameMigrated, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "NameMigrated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbNameMigrated)
				if err := _Arb.contract.UnpackLog(event, "NameMigrated", log); err != nil {
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

// ParseNameMigrated is a log parse operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Arb *ArbFilterer) ParseNameMigrated(log types.Log) (*ArbNameMigrated, error) {
	event := new(ArbNameMigrated)
	if err := _Arb.contract.UnpackLog(event, "NameMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Arb contract.
type ArbNameRegisteredIterator struct {
	Event *ArbNameRegistered // Event containing the contract specifics and raw log

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
func (it *ArbNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbNameRegistered)
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
		it.Event = new(ArbNameRegistered)
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
func (it *ArbNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbNameRegistered represents a NameRegistered event raised by the Arb contract.
type ArbNameRegistered struct {
	Id      *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Arb *ArbFilterer) FilterNameRegistered(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*ArbNameRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "NameRegistered", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ArbNameRegisteredIterator{contract: _Arb.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Arb *ArbFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *ArbNameRegistered, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "NameRegistered", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbNameRegistered)
				if err := _Arb.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Arb *ArbFilterer) ParseNameRegistered(log types.Log) (*ArbNameRegistered, error) {
	event := new(ArbNameRegistered)
	if err := _Arb.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Arb contract.
type ArbNameRenewedIterator struct {
	Event *ArbNameRenewed // Event containing the contract specifics and raw log

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
func (it *ArbNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbNameRenewed)
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
		it.Event = new(ArbNameRenewed)
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
func (it *ArbNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbNameRenewed represents a NameRenewed event raised by the Arb contract.
type ArbNameRenewed struct {
	Id      *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Arb *ArbFilterer) FilterNameRenewed(opts *bind.FilterOpts, id []*big.Int) (*ArbNameRenewedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "NameRenewed", idRule)
	if err != nil {
		return nil, err
	}
	return &ArbNameRenewedIterator{contract: _Arb.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Arb *ArbFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *ArbNameRenewed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "NameRenewed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbNameRenewed)
				if err := _Arb.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Arb *ArbFilterer) ParseNameRenewed(log types.Log) (*ArbNameRenewed, error) {
	event := new(ArbNameRenewed)
	if err := _Arb.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Arb contract.
type ArbOwnershipTransferredIterator struct {
	Event *ArbOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbOwnershipTransferred)
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
		it.Event = new(ArbOwnershipTransferred)
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
func (it *ArbOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbOwnershipTransferred represents a OwnershipTransferred event raised by the Arb contract.
type ArbOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbOwnershipTransferredIterator{contract: _Arb.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbOwnershipTransferred)
				if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) ParseOwnershipTransferred(log types.Log) (*ArbOwnershipTransferred, error) {
	event := new(ArbOwnershipTransferred)
	if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Arb contract.
type ArbTransferIterator struct {
	Event *ArbTransfer // Event containing the contract specifics and raw log

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
func (it *ArbTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTransfer)
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
		it.Event = new(ArbTransfer)
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
func (it *ArbTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTransfer represents a Transfer event raised by the Arb contract.
type ArbTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Arb *ArbFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ArbTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbTransferIterator{contract: _Arb.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Arb *ArbFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArbTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTransfer)
				if err := _Arb.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Arb *ArbFilterer) ParseTransfer(log types.Log) (*ArbTransfer, error) {
	event := new(ArbTransfer)
	if err := _Arb.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
