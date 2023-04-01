// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cns

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

// CnsMetaData contains all meta data concerning the Cns contract.
var CnsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCNS\",\"name\":\"_cns\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_baseNode\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"ControllerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"ControllerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cns\",\"outputs\":[{\"internalType\":\"contractCNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"nameExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"registerOnly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_base\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CnsABI is the input ABI used to generate the binding from.
// Deprecated: Use CnsMetaData.ABI instead.
var CnsABI = CnsMetaData.ABI

// Cns is an auto generated Go binding around an Ethereum contract.
type Cns struct {
	CnsCaller     // Read-only binding to the contract
	CnsTransactor // Write-only binding to the contract
	CnsFilterer   // Log filterer for contract events
}

// CnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CnsSession struct {
	Contract     *Cns              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CnsCallerSession struct {
	Contract *CnsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CnsTransactorSession struct {
	Contract     *CnsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CnsRaw struct {
	Contract *Cns // Generic contract binding to access the raw methods on
}

// CnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CnsCallerRaw struct {
	Contract *CnsCaller // Generic read-only contract binding to access the raw methods on
}

// CnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CnsTransactorRaw struct {
	Contract *CnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCns creates a new instance of Cns, bound to a specific deployed contract.
func NewCns(address common.Address, backend bind.ContractBackend) (*Cns, error) {
	contract, err := bindCns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cns{CnsCaller: CnsCaller{contract: contract}, CnsTransactor: CnsTransactor{contract: contract}, CnsFilterer: CnsFilterer{contract: contract}}, nil
}

// NewCnsCaller creates a new read-only instance of Cns, bound to a specific deployed contract.
func NewCnsCaller(address common.Address, caller bind.ContractCaller) (*CnsCaller, error) {
	contract, err := bindCns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CnsCaller{contract: contract}, nil
}

// NewCnsTransactor creates a new write-only instance of Cns, bound to a specific deployed contract.
func NewCnsTransactor(address common.Address, transactor bind.ContractTransactor) (*CnsTransactor, error) {
	contract, err := bindCns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CnsTransactor{contract: contract}, nil
}

// NewCnsFilterer creates a new log filterer instance of Cns, bound to a specific deployed contract.
func NewCnsFilterer(address common.Address, filterer bind.ContractFilterer) (*CnsFilterer, error) {
	contract, err := bindCns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CnsFilterer{contract: contract}, nil
}

// bindCns binds a generic wrapper to an already deployed contract.
func bindCns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CnsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cns *CnsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cns.Contract.CnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cns *CnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cns.Contract.CnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cns *CnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cns.Contract.CnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cns *CnsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cns *CnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cns *CnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cns.Contract.contract.Transact(opts, method, params...)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Cns *CnsCaller) GRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Cns *CnsSession) GRACEPERIOD() (*big.Int, error) {
	return _Cns.Contract.GRACEPERIOD(&_Cns.CallOpts)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Cns *CnsCallerSession) GRACEPERIOD() (*big.Int, error) {
	return _Cns.Contract.GRACEPERIOD(&_Cns.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Cns *CnsCaller) Available(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "available", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Cns *CnsSession) Available(id *big.Int) (bool, error) {
	return _Cns.Contract.Available(&_Cns.CallOpts, id)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Cns *CnsCallerSession) Available(id *big.Int) (bool, error) {
	return _Cns.Contract.Available(&_Cns.CallOpts, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cns *CnsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cns *CnsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cns.Contract.BalanceOf(&_Cns.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cns *CnsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cns.Contract.BalanceOf(&_Cns.CallOpts, owner)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Cns *CnsCaller) BaseNode(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "baseNode")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Cns *CnsSession) BaseNode() ([32]byte, error) {
	return _Cns.Contract.BaseNode(&_Cns.CallOpts)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Cns *CnsCallerSession) BaseNode() ([32]byte, error) {
	return _Cns.Contract.BaseNode(&_Cns.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Cns *CnsCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Cns *CnsSession) BaseURI() (string, error) {
	return _Cns.Contract.BaseURI(&_Cns.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Cns *CnsCallerSession) BaseURI() (string, error) {
	return _Cns.Contract.BaseURI(&_Cns.CallOpts)
}

// Cns is a free data retrieval call binding the contract method 0xf1c248b1.
//
// Solidity: function cns() view returns(address)
func (_Cns *CnsCaller) Cns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "cns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cns is a free data retrieval call binding the contract method 0xf1c248b1.
//
// Solidity: function cns() view returns(address)
func (_Cns *CnsSession) Cns() (common.Address, error) {
	return _Cns.Contract.Cns(&_Cns.CallOpts)
}

// Cns is a free data retrieval call binding the contract method 0xf1c248b1.
//
// Solidity: function cns() view returns(address)
func (_Cns *CnsCallerSession) Cns() (common.Address, error) {
	return _Cns.Contract.Cns(&_Cns.CallOpts)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Cns *CnsCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Cns *CnsSession) Controllers(arg0 common.Address) (bool, error) {
	return _Cns.Contract.Controllers(&_Cns.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Cns *CnsCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _Cns.Contract.Controllers(&_Cns.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cns *CnsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cns *CnsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cns.Contract.GetApproved(&_Cns.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cns *CnsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cns.Contract.GetApproved(&_Cns.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cns *CnsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cns *CnsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cns.Contract.IsApprovedForAll(&_Cns.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cns *CnsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cns.Contract.IsApprovedForAll(&_Cns.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cns *CnsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cns *CnsSession) Name() (string, error) {
	return _Cns.Contract.Name(&_Cns.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cns *CnsCallerSession) Name() (string, error) {
	return _Cns.Contract.Name(&_Cns.CallOpts)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Cns *CnsCaller) NameExpires(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "nameExpires", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Cns *CnsSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _Cns.Contract.NameExpires(&_Cns.CallOpts, id)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Cns *CnsCallerSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _Cns.Contract.NameExpires(&_Cns.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cns *CnsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cns *CnsSession) Owner() (common.Address, error) {
	return _Cns.Contract.Owner(&_Cns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cns *CnsCallerSession) Owner() (common.Address, error) {
	return _Cns.Contract.Owner(&_Cns.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cns *CnsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cns *CnsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cns.Contract.OwnerOf(&_Cns.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cns *CnsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cns.Contract.OwnerOf(&_Cns.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cns *CnsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cns *CnsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cns.Contract.SupportsInterface(&_Cns.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cns *CnsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cns.Contract.SupportsInterface(&_Cns.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cns *CnsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cns *CnsSession) Symbol() (string, error) {
	return _Cns.Contract.Symbol(&_Cns.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cns *CnsCallerSession) Symbol() (string, error) {
	return _Cns.Contract.Symbol(&_Cns.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cns *CnsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cns *CnsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Cns.Contract.TokenByIndex(&_Cns.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cns *CnsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Cns.Contract.TokenByIndex(&_Cns.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cns *CnsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cns *CnsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Cns.Contract.TokenOfOwnerByIndex(&_Cns.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cns *CnsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Cns.Contract.TokenOfOwnerByIndex(&_Cns.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cns *CnsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cns *CnsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Cns.Contract.TokenURI(&_Cns.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cns *CnsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Cns.Contract.TokenURI(&_Cns.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cns *CnsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cns.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cns *CnsSession) TotalSupply() (*big.Int, error) {
	return _Cns.Contract.TotalSupply(&_Cns.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cns *CnsCallerSession) TotalSupply() (*big.Int, error) {
	return _Cns.Contract.TotalSupply(&_Cns.CallOpts)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Cns *CnsTransactor) AddController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "addController", controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Cns *CnsSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _Cns.Contract.AddController(&_Cns.TransactOpts, controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Cns *CnsTransactorSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _Cns.Contract.AddController(&_Cns.TransactOpts, controller)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cns *CnsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cns *CnsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.Approve(&_Cns.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cns *CnsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.Approve(&_Cns.TransactOpts, to, tokenId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Cns *CnsTransactor) Reclaim(opts *bind.TransactOpts, id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "reclaim", id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Cns *CnsSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Cns.Contract.Reclaim(&_Cns.TransactOpts, id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Cns *CnsTransactorSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Cns.Contract.Reclaim(&_Cns.TransactOpts, id, owner)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Cns *CnsTransactor) Register(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "register", id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Cns *CnsSession) Register(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.Register(&_Cns.TransactOpts, id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Cns *CnsTransactorSession) Register(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.Register(&_Cns.TransactOpts, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Cns *CnsTransactor) RegisterOnly(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "registerOnly", id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Cns *CnsSession) RegisterOnly(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.RegisterOnly(&_Cns.TransactOpts, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Cns *CnsTransactorSession) RegisterOnly(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.RegisterOnly(&_Cns.TransactOpts, id, owner, duration)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Cns *CnsTransactor) RemoveController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "removeController", controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Cns *CnsSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _Cns.Contract.RemoveController(&_Cns.TransactOpts, controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Cns *CnsTransactorSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _Cns.Contract.RemoveController(&_Cns.TransactOpts, controller)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Cns *CnsTransactor) Renew(opts *bind.TransactOpts, id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "renew", id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Cns *CnsSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.Renew(&_Cns.TransactOpts, id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Cns *CnsTransactorSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.Renew(&_Cns.TransactOpts, id, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cns *CnsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cns *CnsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cns.Contract.RenounceOwnership(&_Cns.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cns *CnsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cns.Contract.RenounceOwnership(&_Cns.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cns *CnsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cns *CnsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.SafeTransferFrom(&_Cns.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cns *CnsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.SafeTransferFrom(&_Cns.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cns *CnsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cns *CnsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cns.Contract.SafeTransferFrom0(&_Cns.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cns *CnsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cns.Contract.SafeTransferFrom0(&_Cns.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cns *CnsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cns *CnsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cns.Contract.SetApprovalForAll(&_Cns.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cns *CnsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cns.Contract.SetApprovalForAll(&_Cns.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _base) returns()
func (_Cns *CnsTransactor) SetBaseURI(opts *bind.TransactOpts, _base string) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "setBaseURI", _base)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _base) returns()
func (_Cns *CnsSession) SetBaseURI(_base string) (*types.Transaction, error) {
	return _Cns.Contract.SetBaseURI(&_Cns.TransactOpts, _base)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _base) returns()
func (_Cns *CnsTransactorSession) SetBaseURI(_base string) (*types.Transaction, error) {
	return _Cns.Contract.SetBaseURI(&_Cns.TransactOpts, _base)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Cns *CnsTransactor) SetResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "setResolver", resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Cns *CnsSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Cns.Contract.SetResolver(&_Cns.TransactOpts, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Cns *CnsTransactorSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Cns.Contract.SetResolver(&_Cns.TransactOpts, resolver)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cns *CnsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cns *CnsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.TransferFrom(&_Cns.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cns *CnsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cns.Contract.TransferFrom(&_Cns.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cns *CnsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cns.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cns *CnsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cns.Contract.TransferOwnership(&_Cns.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cns *CnsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cns.Contract.TransferOwnership(&_Cns.TransactOpts, newOwner)
}

// CnsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cns contract.
type CnsApprovalIterator struct {
	Event *CnsApproval // Event containing the contract specifics and raw log

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
func (it *CnsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsApproval)
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
		it.Event = new(CnsApproval)
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
func (it *CnsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsApproval represents a Approval event raised by the Cns contract.
type CnsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cns *CnsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CnsApprovalIterator, error) {

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

	logs, sub, err := _Cns.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CnsApprovalIterator{contract: _Cns.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cns *CnsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CnsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Cns.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsApproval)
				if err := _Cns.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Cns *CnsFilterer) ParseApproval(log types.Log) (*CnsApproval, error) {
	event := new(CnsApproval)
	if err := _Cns.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Cns contract.
type CnsApprovalForAllIterator struct {
	Event *CnsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CnsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsApprovalForAll)
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
		it.Event = new(CnsApprovalForAll)
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
func (it *CnsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsApprovalForAll represents a ApprovalForAll event raised by the Cns contract.
type CnsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cns *CnsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CnsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CnsApprovalForAllIterator{contract: _Cns.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cns *CnsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CnsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsApprovalForAll)
				if err := _Cns.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Cns *CnsFilterer) ParseApprovalForAll(log types.Log) (*CnsApprovalForAll, error) {
	event := new(CnsApprovalForAll)
	if err := _Cns.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsControllerAddedIterator is returned from FilterControllerAdded and is used to iterate over the raw logs and unpacked data for ControllerAdded events raised by the Cns contract.
type CnsControllerAddedIterator struct {
	Event *CnsControllerAdded // Event containing the contract specifics and raw log

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
func (it *CnsControllerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsControllerAdded)
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
		it.Event = new(CnsControllerAdded)
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
func (it *CnsControllerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsControllerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsControllerAdded represents a ControllerAdded event raised by the Cns contract.
type CnsControllerAdded struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerAdded is a free log retrieval operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Cns *CnsFilterer) FilterControllerAdded(opts *bind.FilterOpts, controller []common.Address) (*CnsControllerAddedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "ControllerAdded", controllerRule)
	if err != nil {
		return nil, err
	}
	return &CnsControllerAddedIterator{contract: _Cns.contract, event: "ControllerAdded", logs: logs, sub: sub}, nil
}

// WatchControllerAdded is a free log subscription operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Cns *CnsFilterer) WatchControllerAdded(opts *bind.WatchOpts, sink chan<- *CnsControllerAdded, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "ControllerAdded", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsControllerAdded)
				if err := _Cns.contract.UnpackLog(event, "ControllerAdded", log); err != nil {
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
func (_Cns *CnsFilterer) ParseControllerAdded(log types.Log) (*CnsControllerAdded, error) {
	event := new(CnsControllerAdded)
	if err := _Cns.contract.UnpackLog(event, "ControllerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsControllerRemovedIterator is returned from FilterControllerRemoved and is used to iterate over the raw logs and unpacked data for ControllerRemoved events raised by the Cns contract.
type CnsControllerRemovedIterator struct {
	Event *CnsControllerRemoved // Event containing the contract specifics and raw log

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
func (it *CnsControllerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsControllerRemoved)
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
		it.Event = new(CnsControllerRemoved)
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
func (it *CnsControllerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsControllerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsControllerRemoved represents a ControllerRemoved event raised by the Cns contract.
type CnsControllerRemoved struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerRemoved is a free log retrieval operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Cns *CnsFilterer) FilterControllerRemoved(opts *bind.FilterOpts, controller []common.Address) (*CnsControllerRemovedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "ControllerRemoved", controllerRule)
	if err != nil {
		return nil, err
	}
	return &CnsControllerRemovedIterator{contract: _Cns.contract, event: "ControllerRemoved", logs: logs, sub: sub}, nil
}

// WatchControllerRemoved is a free log subscription operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Cns *CnsFilterer) WatchControllerRemoved(opts *bind.WatchOpts, sink chan<- *CnsControllerRemoved, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "ControllerRemoved", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsControllerRemoved)
				if err := _Cns.contract.UnpackLog(event, "ControllerRemoved", log); err != nil {
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
func (_Cns *CnsFilterer) ParseControllerRemoved(log types.Log) (*CnsControllerRemoved, error) {
	event := new(CnsControllerRemoved)
	if err := _Cns.contract.UnpackLog(event, "ControllerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsNameMigratedIterator is returned from FilterNameMigrated and is used to iterate over the raw logs and unpacked data for NameMigrated events raised by the Cns contract.
type CnsNameMigratedIterator struct {
	Event *CnsNameMigrated // Event containing the contract specifics and raw log

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
func (it *CnsNameMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsNameMigrated)
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
		it.Event = new(CnsNameMigrated)
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
func (it *CnsNameMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsNameMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsNameMigrated represents a NameMigrated event raised by the Cns contract.
type CnsNameMigrated struct {
	Id      *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameMigrated is a free log retrieval operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Cns *CnsFilterer) FilterNameMigrated(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*CnsNameMigratedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "NameMigrated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CnsNameMigratedIterator{contract: _Cns.contract, event: "NameMigrated", logs: logs, sub: sub}, nil
}

// WatchNameMigrated is a free log subscription operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Cns *CnsFilterer) WatchNameMigrated(opts *bind.WatchOpts, sink chan<- *CnsNameMigrated, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "NameMigrated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsNameMigrated)
				if err := _Cns.contract.UnpackLog(event, "NameMigrated", log); err != nil {
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
func (_Cns *CnsFilterer) ParseNameMigrated(log types.Log) (*CnsNameMigrated, error) {
	event := new(CnsNameMigrated)
	if err := _Cns.contract.UnpackLog(event, "NameMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Cns contract.
type CnsNameRegisteredIterator struct {
	Event *CnsNameRegistered // Event containing the contract specifics and raw log

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
func (it *CnsNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsNameRegistered)
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
		it.Event = new(CnsNameRegistered)
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
func (it *CnsNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsNameRegistered represents a NameRegistered event raised by the Cns contract.
type CnsNameRegistered struct {
	Id      *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Cns *CnsFilterer) FilterNameRegistered(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*CnsNameRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "NameRegistered", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CnsNameRegisteredIterator{contract: _Cns.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Cns *CnsFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *CnsNameRegistered, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "NameRegistered", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsNameRegistered)
				if err := _Cns.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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
func (_Cns *CnsFilterer) ParseNameRegistered(log types.Log) (*CnsNameRegistered, error) {
	event := new(CnsNameRegistered)
	if err := _Cns.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Cns contract.
type CnsNameRenewedIterator struct {
	Event *CnsNameRenewed // Event containing the contract specifics and raw log

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
func (it *CnsNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsNameRenewed)
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
		it.Event = new(CnsNameRenewed)
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
func (it *CnsNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsNameRenewed represents a NameRenewed event raised by the Cns contract.
type CnsNameRenewed struct {
	Id      *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Cns *CnsFilterer) FilterNameRenewed(opts *bind.FilterOpts, id []*big.Int) (*CnsNameRenewedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "NameRenewed", idRule)
	if err != nil {
		return nil, err
	}
	return &CnsNameRenewedIterator{contract: _Cns.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Cns *CnsFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *CnsNameRenewed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "NameRenewed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsNameRenewed)
				if err := _Cns.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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
func (_Cns *CnsFilterer) ParseNameRenewed(log types.Log) (*CnsNameRenewed, error) {
	event := new(CnsNameRenewed)
	if err := _Cns.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cns contract.
type CnsOwnershipTransferredIterator struct {
	Event *CnsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CnsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsOwnershipTransferred)
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
		it.Event = new(CnsOwnershipTransferred)
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
func (it *CnsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsOwnershipTransferred represents a OwnershipTransferred event raised by the Cns contract.
type CnsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cns *CnsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CnsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cns.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CnsOwnershipTransferredIterator{contract: _Cns.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cns *CnsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CnsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cns.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsOwnershipTransferred)
				if err := _Cns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cns *CnsFilterer) ParseOwnershipTransferred(log types.Log) (*CnsOwnershipTransferred, error) {
	event := new(CnsOwnershipTransferred)
	if err := _Cns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CnsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cns contract.
type CnsTransferIterator struct {
	Event *CnsTransfer // Event containing the contract specifics and raw log

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
func (it *CnsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CnsTransfer)
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
		it.Event = new(CnsTransfer)
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
func (it *CnsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CnsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CnsTransfer represents a Transfer event raised by the Cns contract.
type CnsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cns *CnsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CnsTransferIterator, error) {

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

	logs, sub, err := _Cns.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CnsTransferIterator{contract: _Cns.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cns *CnsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CnsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Cns.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CnsTransfer)
				if err := _Cns.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Cns *CnsFilterer) ParseTransfer(log types.Log) (*CnsTransfer, error) {
	event := new(CnsTransfer)
	if err := _Cns.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
