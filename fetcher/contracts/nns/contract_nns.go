// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nns

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

// NnsMetaData contains all meta data concerning the Nns contract.
var NnsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"ControllerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"ControllerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNNS\",\"name\":\"_nns\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_baseNode\",\"type\":\"bytes32\"}],\"name\":\"__BaseRegistrarImplementation_i\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"nameExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nns\",\"outputs\":[{\"internalType\":\"contractNNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"registerOnly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerWithName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_baseNode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setBaseNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NnsABI is the input ABI used to generate the binding from.
// Deprecated: Use NnsMetaData.ABI instead.
var NnsABI = NnsMetaData.ABI

// Nns is an auto generated Go binding around an Ethereum contract.
type Nns struct {
	NnsCaller     // Read-only binding to the contract
	NnsTransactor // Write-only binding to the contract
	NnsFilterer   // Log filterer for contract events
}

// NnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NnsSession struct {
	Contract     *Nns              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NnsCallerSession struct {
	Contract *NnsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NnsTransactorSession struct {
	Contract     *NnsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NnsRaw struct {
	Contract *Nns // Generic contract binding to access the raw methods on
}

// NnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NnsCallerRaw struct {
	Contract *NnsCaller // Generic read-only contract binding to access the raw methods on
}

// NnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NnsTransactorRaw struct {
	Contract *NnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNns creates a new instance of Nns, bound to a specific deployed contract.
func NewNns(address common.Address, backend bind.ContractBackend) (*Nns, error) {
	contract, err := bindNns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nns{NnsCaller: NnsCaller{contract: contract}, NnsTransactor: NnsTransactor{contract: contract}, NnsFilterer: NnsFilterer{contract: contract}}, nil
}

// NewNnsCaller creates a new read-only instance of Nns, bound to a specific deployed contract.
func NewNnsCaller(address common.Address, caller bind.ContractCaller) (*NnsCaller, error) {
	contract, err := bindNns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NnsCaller{contract: contract}, nil
}

// NewNnsTransactor creates a new write-only instance of Nns, bound to a specific deployed contract.
func NewNnsTransactor(address common.Address, transactor bind.ContractTransactor) (*NnsTransactor, error) {
	contract, err := bindNns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NnsTransactor{contract: contract}, nil
}

// NewNnsFilterer creates a new log filterer instance of Nns, bound to a specific deployed contract.
func NewNnsFilterer(address common.Address, filterer bind.ContractFilterer) (*NnsFilterer, error) {
	contract, err := bindNns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NnsFilterer{contract: contract}, nil
}

// bindNns binds a generic wrapper to an already deployed contract.
func bindNns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NnsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nns *NnsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nns.Contract.NnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nns *NnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nns.Contract.NnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nns *NnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nns.Contract.NnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nns *NnsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nns *NnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nns *NnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nns.Contract.contract.Transact(opts, method, params...)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Nns *NnsCaller) GRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Nns *NnsSession) GRACEPERIOD() (*big.Int, error) {
	return _Nns.Contract.GRACEPERIOD(&_Nns.CallOpts)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Nns *NnsCallerSession) GRACEPERIOD() (*big.Int, error) {
	return _Nns.Contract.GRACEPERIOD(&_Nns.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xcfc86f7b.
//
// Solidity: function _baseTokenURI() view returns(string)
func (_Nns *NnsCaller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "_baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xcfc86f7b.
//
// Solidity: function _baseTokenURI() view returns(string)
func (_Nns *NnsSession) BaseTokenURI() (string, error) {
	return _Nns.Contract.BaseTokenURI(&_Nns.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xcfc86f7b.
//
// Solidity: function _baseTokenURI() view returns(string)
func (_Nns *NnsCallerSession) BaseTokenURI() (string, error) {
	return _Nns.Contract.BaseTokenURI(&_Nns.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Nns *NnsCaller) Available(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "available", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Nns *NnsSession) Available(id *big.Int) (bool, error) {
	return _Nns.Contract.Available(&_Nns.CallOpts, id)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_Nns *NnsCallerSession) Available(id *big.Int) (bool, error) {
	return _Nns.Contract.Available(&_Nns.CallOpts, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nns *NnsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nns *NnsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nns.Contract.BalanceOf(&_Nns.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nns *NnsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nns.Contract.BalanceOf(&_Nns.CallOpts, owner)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Nns *NnsCaller) BaseNode(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "baseNode")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Nns *NnsSession) BaseNode() ([32]byte, error) {
	return _Nns.Contract.BaseNode(&_Nns.CallOpts)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Nns *NnsCallerSession) BaseNode() ([32]byte, error) {
	return _Nns.Contract.BaseNode(&_Nns.CallOpts)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Nns *NnsCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Nns *NnsSession) Controllers(arg0 common.Address) (bool, error) {
	return _Nns.Contract.Controllers(&_Nns.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Nns *NnsCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _Nns.Contract.Controllers(&_Nns.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nns *NnsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nns *NnsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nns.Contract.GetApproved(&_Nns.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nns *NnsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nns.Contract.GetApproved(&_Nns.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nns *NnsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nns *NnsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nns.Contract.IsApprovedForAll(&_Nns.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nns *NnsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nns.Contract.IsApprovedForAll(&_Nns.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nns *NnsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nns *NnsSession) Name() (string, error) {
	return _Nns.Contract.Name(&_Nns.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nns *NnsCallerSession) Name() (string, error) {
	return _Nns.Contract.Name(&_Nns.CallOpts)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Nns *NnsCaller) NameExpires(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "nameExpires", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Nns *NnsSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _Nns.Contract.NameExpires(&_Nns.CallOpts, id)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_Nns *NnsCallerSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _Nns.Contract.NameExpires(&_Nns.CallOpts, id)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_Nns *NnsCaller) Names(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_Nns *NnsSession) Names(arg0 *big.Int) (string, error) {
	return _Nns.Contract.Names(&_Nns.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_Nns *NnsCallerSession) Names(arg0 *big.Int) (string, error) {
	return _Nns.Contract.Names(&_Nns.CallOpts, arg0)
}

// Nns is a free data retrieval call binding the contract method 0xc358133b.
//
// Solidity: function nns() view returns(address)
func (_Nns *NnsCaller) Nns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "nns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nns is a free data retrieval call binding the contract method 0xc358133b.
//
// Solidity: function nns() view returns(address)
func (_Nns *NnsSession) Nns() (common.Address, error) {
	return _Nns.Contract.Nns(&_Nns.CallOpts)
}

// Nns is a free data retrieval call binding the contract method 0xc358133b.
//
// Solidity: function nns() view returns(address)
func (_Nns *NnsCallerSession) Nns() (common.Address, error) {
	return _Nns.Contract.Nns(&_Nns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nns *NnsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nns *NnsSession) Owner() (common.Address, error) {
	return _Nns.Contract.Owner(&_Nns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nns *NnsCallerSession) Owner() (common.Address, error) {
	return _Nns.Contract.Owner(&_Nns.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nns *NnsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nns *NnsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nns.Contract.OwnerOf(&_Nns.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nns *NnsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nns.Contract.OwnerOf(&_Nns.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Nns *NnsCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Nns *NnsSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Nns.Contract.SupportsInterface(&_Nns.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Nns *NnsCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Nns.Contract.SupportsInterface(&_Nns.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nns *NnsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nns *NnsSession) Symbol() (string, error) {
	return _Nns.Contract.Symbol(&_Nns.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nns *NnsCallerSession) Symbol() (string, error) {
	return _Nns.Contract.Symbol(&_Nns.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Nns *NnsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Nns *NnsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Nns.Contract.TokenByIndex(&_Nns.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Nns *NnsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Nns.Contract.TokenByIndex(&_Nns.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Nns *NnsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Nns *NnsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Nns.Contract.TokenOfOwnerByIndex(&_Nns.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Nns *NnsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Nns.Contract.TokenOfOwnerByIndex(&_Nns.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nns *NnsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nns *NnsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nns.Contract.TokenURI(&_Nns.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nns *NnsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nns.Contract.TokenURI(&_Nns.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nns *NnsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nns.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nns *NnsSession) TotalSupply() (*big.Int, error) {
	return _Nns.Contract.TotalSupply(&_Nns.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nns *NnsCallerSession) TotalSupply() (*big.Int, error) {
	return _Nns.Contract.TotalSupply(&_Nns.CallOpts)
}

// BaseRegistrarImplementationI is a paid mutator transaction binding the contract method 0x1e34ac0f.
//
// Solidity: function __BaseRegistrarImplementation_i(address _nns, bytes32 _baseNode) returns()
func (_Nns *NnsTransactor) BaseRegistrarImplementationI(opts *bind.TransactOpts, _nns common.Address, _baseNode [32]byte) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "__BaseRegistrarImplementation_i", _nns, _baseNode)
}

// BaseRegistrarImplementationI is a paid mutator transaction binding the contract method 0x1e34ac0f.
//
// Solidity: function __BaseRegistrarImplementation_i(address _nns, bytes32 _baseNode) returns()
func (_Nns *NnsSession) BaseRegistrarImplementationI(_nns common.Address, _baseNode [32]byte) (*types.Transaction, error) {
	return _Nns.Contract.BaseRegistrarImplementationI(&_Nns.TransactOpts, _nns, _baseNode)
}

// BaseRegistrarImplementationI is a paid mutator transaction binding the contract method 0x1e34ac0f.
//
// Solidity: function __BaseRegistrarImplementation_i(address _nns, bytes32 _baseNode) returns()
func (_Nns *NnsTransactorSession) BaseRegistrarImplementationI(_nns common.Address, _baseNode [32]byte) (*types.Transaction, error) {
	return _Nns.Contract.BaseRegistrarImplementationI(&_Nns.TransactOpts, _nns, _baseNode)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Nns *NnsTransactor) AddController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "addController", controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Nns *NnsSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _Nns.Contract.AddController(&_Nns.TransactOpts, controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_Nns *NnsTransactorSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _Nns.Contract.AddController(&_Nns.TransactOpts, controller)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nns *NnsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nns *NnsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Approve(&_Nns.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nns *NnsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Approve(&_Nns.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Nns *NnsTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Nns *NnsSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Burn(&_Nns.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Nns *NnsTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Burn(&_Nns.TransactOpts, tokenId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Nns *NnsTransactor) Reclaim(opts *bind.TransactOpts, id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "reclaim", id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Nns *NnsSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Nns.Contract.Reclaim(&_Nns.TransactOpts, id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_Nns *NnsTransactorSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Nns.Contract.Reclaim(&_Nns.TransactOpts, id, owner)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Nns *NnsTransactor) Register(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "register", id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Nns *NnsSession) Register(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Register(&_Nns.TransactOpts, id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0xfca247ac.
//
// Solidity: function register(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Nns *NnsTransactorSession) Register(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Register(&_Nns.TransactOpts, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Nns *NnsTransactor) RegisterOnly(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "registerOnly", id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Nns *NnsSession) RegisterOnly(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.RegisterOnly(&_Nns.TransactOpts, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x0e297b45.
//
// Solidity: function registerOnly(uint256 id, address owner, uint256 duration) returns(uint256)
func (_Nns *NnsTransactorSession) RegisterOnly(id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.RegisterOnly(&_Nns.TransactOpts, id, owner, duration)
}

// RegisterWithName is a paid mutator transaction binding the contract method 0xc4671ac2.
//
// Solidity: function registerWithName(uint256 id, address owner, uint256 duration, string name) returns(uint256)
func (_Nns *NnsTransactor) RegisterWithName(opts *bind.TransactOpts, id *big.Int, owner common.Address, duration *big.Int, name string) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "registerWithName", id, owner, duration, name)
}

// RegisterWithName is a paid mutator transaction binding the contract method 0xc4671ac2.
//
// Solidity: function registerWithName(uint256 id, address owner, uint256 duration, string name) returns(uint256)
func (_Nns *NnsSession) RegisterWithName(id *big.Int, owner common.Address, duration *big.Int, name string) (*types.Transaction, error) {
	return _Nns.Contract.RegisterWithName(&_Nns.TransactOpts, id, owner, duration, name)
}

// RegisterWithName is a paid mutator transaction binding the contract method 0xc4671ac2.
//
// Solidity: function registerWithName(uint256 id, address owner, uint256 duration, string name) returns(uint256)
func (_Nns *NnsTransactorSession) RegisterWithName(id *big.Int, owner common.Address, duration *big.Int, name string) (*types.Transaction, error) {
	return _Nns.Contract.RegisterWithName(&_Nns.TransactOpts, id, owner, duration, name)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Nns *NnsTransactor) RemoveController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "removeController", controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Nns *NnsSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _Nns.Contract.RemoveController(&_Nns.TransactOpts, controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_Nns *NnsTransactorSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _Nns.Contract.RemoveController(&_Nns.TransactOpts, controller)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Nns *NnsTransactor) Renew(opts *bind.TransactOpts, id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "renew", id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Nns *NnsSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Renew(&_Nns.TransactOpts, id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_Nns *NnsTransactorSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.Renew(&_Nns.TransactOpts, id, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nns *NnsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nns *NnsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nns.Contract.RenounceOwnership(&_Nns.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nns *NnsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nns.Contract.RenounceOwnership(&_Nns.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nns *NnsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nns *NnsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.SafeTransferFrom(&_Nns.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nns *NnsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.SafeTransferFrom(&_Nns.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nns *NnsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nns *NnsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nns.Contract.SafeTransferFrom0(&_Nns.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nns *NnsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nns.Contract.SafeTransferFrom0(&_Nns.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nns *NnsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nns *NnsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nns.Contract.SetApprovalForAll(&_Nns.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nns *NnsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nns.Contract.SetApprovalForAll(&_Nns.TransactOpts, operator, approved)
}

// SetBaseNode is a paid mutator transaction binding the contract method 0x2611dd5b.
//
// Solidity: function setBaseNode(bytes32 _baseNode, address resolver) returns()
func (_Nns *NnsTransactor) SetBaseNode(opts *bind.TransactOpts, _baseNode [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "setBaseNode", _baseNode, resolver)
}

// SetBaseNode is a paid mutator transaction binding the contract method 0x2611dd5b.
//
// Solidity: function setBaseNode(bytes32 _baseNode, address resolver) returns()
func (_Nns *NnsSession) SetBaseNode(_baseNode [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Nns.Contract.SetBaseNode(&_Nns.TransactOpts, _baseNode, resolver)
}

// SetBaseNode is a paid mutator transaction binding the contract method 0x2611dd5b.
//
// Solidity: function setBaseNode(bytes32 _baseNode, address resolver) returns()
func (_Nns *NnsTransactorSession) SetBaseNode(_baseNode [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Nns.Contract.SetBaseNode(&_Nns.TransactOpts, _baseNode, resolver)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Nns *NnsTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Nns *NnsSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Nns.Contract.SetBaseURI(&_Nns.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Nns *NnsTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Nns.Contract.SetBaseURI(&_Nns.TransactOpts, baseURI)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 id, string name) returns()
func (_Nns *NnsTransactor) SetName(opts *bind.TransactOpts, id *big.Int, name string) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "setName", id, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 id, string name) returns()
func (_Nns *NnsSession) SetName(id *big.Int, name string) (*types.Transaction, error) {
	return _Nns.Contract.SetName(&_Nns.TransactOpts, id, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 id, string name) returns()
func (_Nns *NnsTransactorSession) SetName(id *big.Int, name string) (*types.Transaction, error) {
	return _Nns.Contract.SetName(&_Nns.TransactOpts, id, name)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Nns *NnsTransactor) SetResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "setResolver", resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Nns *NnsSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Nns.Contract.SetResolver(&_Nns.TransactOpts, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Nns *NnsTransactorSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Nns.Contract.SetResolver(&_Nns.TransactOpts, resolver)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nns *NnsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nns *NnsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.TransferFrom(&_Nns.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nns *NnsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nns.Contract.TransferFrom(&_Nns.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nns *NnsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nns.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nns *NnsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nns.Contract.TransferOwnership(&_Nns.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nns *NnsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nns.Contract.TransferOwnership(&_Nns.TransactOpts, newOwner)
}

// NnsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nns contract.
type NnsApprovalIterator struct {
	Event *NnsApproval // Event containing the contract specifics and raw log

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
func (it *NnsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsApproval)
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
		it.Event = new(NnsApproval)
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
func (it *NnsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsApproval represents a Approval event raised by the Nns contract.
type NnsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nns *NnsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NnsApprovalIterator, error) {

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

	logs, sub, err := _Nns.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NnsApprovalIterator{contract: _Nns.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nns *NnsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NnsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Nns.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsApproval)
				if err := _Nns.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Nns *NnsFilterer) ParseApproval(log types.Log) (*NnsApproval, error) {
	event := new(NnsApproval)
	if err := _Nns.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nns contract.
type NnsApprovalForAllIterator struct {
	Event *NnsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NnsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsApprovalForAll)
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
		it.Event = new(NnsApprovalForAll)
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
func (it *NnsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsApprovalForAll represents a ApprovalForAll event raised by the Nns contract.
type NnsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nns *NnsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NnsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NnsApprovalForAllIterator{contract: _Nns.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nns *NnsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NnsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsApprovalForAll)
				if err := _Nns.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Nns *NnsFilterer) ParseApprovalForAll(log types.Log) (*NnsApprovalForAll, error) {
	event := new(NnsApprovalForAll)
	if err := _Nns.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsControllerAddedIterator is returned from FilterControllerAdded and is used to iterate over the raw logs and unpacked data for ControllerAdded events raised by the Nns contract.
type NnsControllerAddedIterator struct {
	Event *NnsControllerAdded // Event containing the contract specifics and raw log

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
func (it *NnsControllerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsControllerAdded)
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
		it.Event = new(NnsControllerAdded)
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
func (it *NnsControllerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsControllerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsControllerAdded represents a ControllerAdded event raised by the Nns contract.
type NnsControllerAdded struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerAdded is a free log retrieval operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Nns *NnsFilterer) FilterControllerAdded(opts *bind.FilterOpts, controller []common.Address) (*NnsControllerAddedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "ControllerAdded", controllerRule)
	if err != nil {
		return nil, err
	}
	return &NnsControllerAddedIterator{contract: _Nns.contract, event: "ControllerAdded", logs: logs, sub: sub}, nil
}

// WatchControllerAdded is a free log subscription operation binding the contract event 0x0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474.
//
// Solidity: event ControllerAdded(address indexed controller)
func (_Nns *NnsFilterer) WatchControllerAdded(opts *bind.WatchOpts, sink chan<- *NnsControllerAdded, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "ControllerAdded", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsControllerAdded)
				if err := _Nns.contract.UnpackLog(event, "ControllerAdded", log); err != nil {
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
func (_Nns *NnsFilterer) ParseControllerAdded(log types.Log) (*NnsControllerAdded, error) {
	event := new(NnsControllerAdded)
	if err := _Nns.contract.UnpackLog(event, "ControllerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsControllerRemovedIterator is returned from FilterControllerRemoved and is used to iterate over the raw logs and unpacked data for ControllerRemoved events raised by the Nns contract.
type NnsControllerRemovedIterator struct {
	Event *NnsControllerRemoved // Event containing the contract specifics and raw log

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
func (it *NnsControllerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsControllerRemoved)
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
		it.Event = new(NnsControllerRemoved)
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
func (it *NnsControllerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsControllerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsControllerRemoved represents a ControllerRemoved event raised by the Nns contract.
type NnsControllerRemoved struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerRemoved is a free log retrieval operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Nns *NnsFilterer) FilterControllerRemoved(opts *bind.FilterOpts, controller []common.Address) (*NnsControllerRemovedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "ControllerRemoved", controllerRule)
	if err != nil {
		return nil, err
	}
	return &NnsControllerRemovedIterator{contract: _Nns.contract, event: "ControllerRemoved", logs: logs, sub: sub}, nil
}

// WatchControllerRemoved is a free log subscription operation binding the contract event 0x33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113.
//
// Solidity: event ControllerRemoved(address indexed controller)
func (_Nns *NnsFilterer) WatchControllerRemoved(opts *bind.WatchOpts, sink chan<- *NnsControllerRemoved, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "ControllerRemoved", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsControllerRemoved)
				if err := _Nns.contract.UnpackLog(event, "ControllerRemoved", log); err != nil {
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
func (_Nns *NnsFilterer) ParseControllerRemoved(log types.Log) (*NnsControllerRemoved, error) {
	event := new(NnsControllerRemoved)
	if err := _Nns.contract.UnpackLog(event, "ControllerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Nns contract.
type NnsInitializedIterator struct {
	Event *NnsInitialized // Event containing the contract specifics and raw log

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
func (it *NnsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsInitialized)
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
		it.Event = new(NnsInitialized)
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
func (it *NnsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsInitialized represents a Initialized event raised by the Nns contract.
type NnsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nns *NnsFilterer) FilterInitialized(opts *bind.FilterOpts) (*NnsInitializedIterator, error) {

	logs, sub, err := _Nns.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NnsInitializedIterator{contract: _Nns.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nns *NnsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NnsInitialized) (event.Subscription, error) {

	logs, sub, err := _Nns.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsInitialized)
				if err := _Nns.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nns *NnsFilterer) ParseInitialized(log types.Log) (*NnsInitialized, error) {
	event := new(NnsInitialized)
	if err := _Nns.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsNameMigratedIterator is returned from FilterNameMigrated and is used to iterate over the raw logs and unpacked data for NameMigrated events raised by the Nns contract.
type NnsNameMigratedIterator struct {
	Event *NnsNameMigrated // Event containing the contract specifics and raw log

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
func (it *NnsNameMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsNameMigrated)
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
		it.Event = new(NnsNameMigrated)
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
func (it *NnsNameMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsNameMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsNameMigrated represents a NameMigrated event raised by the Nns contract.
type NnsNameMigrated struct {
	Id      *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameMigrated is a free log retrieval operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Nns *NnsFilterer) FilterNameMigrated(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*NnsNameMigratedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "NameMigrated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &NnsNameMigratedIterator{contract: _Nns.contract, event: "NameMigrated", logs: logs, sub: sub}, nil
}

// WatchNameMigrated is a free log subscription operation binding the contract event 0xea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717.
//
// Solidity: event NameMigrated(uint256 indexed id, address indexed owner, uint256 expires)
func (_Nns *NnsFilterer) WatchNameMigrated(opts *bind.WatchOpts, sink chan<- *NnsNameMigrated, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "NameMigrated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsNameMigrated)
				if err := _Nns.contract.UnpackLog(event, "NameMigrated", log); err != nil {
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
func (_Nns *NnsFilterer) ParseNameMigrated(log types.Log) (*NnsNameMigrated, error) {
	event := new(NnsNameMigrated)
	if err := _Nns.contract.UnpackLog(event, "NameMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Nns contract.
type NnsNameRegisteredIterator struct {
	Event *NnsNameRegistered // Event containing the contract specifics and raw log

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
func (it *NnsNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsNameRegistered)
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
		it.Event = new(NnsNameRegistered)
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
func (it *NnsNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsNameRegistered represents a NameRegistered event raised by the Nns contract.
type NnsNameRegistered struct {
	Id      *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Nns *NnsFilterer) FilterNameRegistered(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*NnsNameRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "NameRegistered", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &NnsNameRegisteredIterator{contract: _Nns.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9.
//
// Solidity: event NameRegistered(uint256 indexed id, address indexed owner, uint256 expires)
func (_Nns *NnsFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *NnsNameRegistered, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "NameRegistered", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsNameRegistered)
				if err := _Nns.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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
func (_Nns *NnsFilterer) ParseNameRegistered(log types.Log) (*NnsNameRegistered, error) {
	event := new(NnsNameRegistered)
	if err := _Nns.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Nns contract.
type NnsNameRenewedIterator struct {
	Event *NnsNameRenewed // Event containing the contract specifics and raw log

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
func (it *NnsNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsNameRenewed)
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
		it.Event = new(NnsNameRenewed)
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
func (it *NnsNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsNameRenewed represents a NameRenewed event raised by the Nns contract.
type NnsNameRenewed struct {
	Id      *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Nns *NnsFilterer) FilterNameRenewed(opts *bind.FilterOpts, id []*big.Int) (*NnsNameRenewedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "NameRenewed", idRule)
	if err != nil {
		return nil, err
	}
	return &NnsNameRenewedIterator{contract: _Nns.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6.
//
// Solidity: event NameRenewed(uint256 indexed id, uint256 expires)
func (_Nns *NnsFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *NnsNameRenewed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "NameRenewed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsNameRenewed)
				if err := _Nns.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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
func (_Nns *NnsFilterer) ParseNameRenewed(log types.Log) (*NnsNameRenewed, error) {
	event := new(NnsNameRenewed)
	if err := _Nns.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nns contract.
type NnsOwnershipTransferredIterator struct {
	Event *NnsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NnsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsOwnershipTransferred)
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
		it.Event = new(NnsOwnershipTransferred)
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
func (it *NnsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsOwnershipTransferred represents a OwnershipTransferred event raised by the Nns contract.
type NnsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nns *NnsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NnsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nns.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NnsOwnershipTransferredIterator{contract: _Nns.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nns *NnsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NnsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nns.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsOwnershipTransferred)
				if err := _Nns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nns *NnsFilterer) ParseOwnershipTransferred(log types.Log) (*NnsOwnershipTransferred, error) {
	event := new(NnsOwnershipTransferred)
	if err := _Nns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NnsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nns contract.
type NnsTransferIterator struct {
	Event *NnsTransfer // Event containing the contract specifics and raw log

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
func (it *NnsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NnsTransfer)
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
		it.Event = new(NnsTransfer)
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
func (it *NnsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NnsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NnsTransfer represents a Transfer event raised by the Nns contract.
type NnsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nns *NnsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NnsTransferIterator, error) {

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

	logs, sub, err := _Nns.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NnsTransferIterator{contract: _Nns.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nns *NnsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NnsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Nns.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NnsTransfer)
				if err := _Nns.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Nns *NnsFilterer) ParseTransfer(log types.Log) (*NnsTransfer, error) {
	event := new(NnsTransfer)
	if err := _Nns.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
