// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ud

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

// IForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type IForwarderForwardRequest struct {
	From    common.Address
	Nonce   *big.Int
	TokenId *big.Int
	Data    []byte
}

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"NewKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"NewURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"NewURIPrefix\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RemoveReverse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ResetRecords\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyIndex\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"valueIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SetReverse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BATCH_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"addKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addProxyReader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"domains\",\"type\":\"string[][]\"}],\"name\":\"backfillReverseNames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositToPolygon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"hashes\",\"type\":\"uint256[]\"}],\"name\":\"getKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMany\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getManyByHash\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintingManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cnsRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rootChainManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childChainManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintTLD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"labels\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"withReverse\",\"type\":\"bool\"}],\"name\":\"mintWithRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"labels\",\"type\":\"string[]\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reconfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeReverse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"resolverOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"reverseNameOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"reverseUri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"reverseOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reverse\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setByHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setManyByHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setReverse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"labels\",\"type\":\"string[]\"}],\"name\":\"setReverse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"labels\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"withReverse\",\"type\":\"bool\"}],\"name\":\"unlockWithRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"withReverse\",\"type\":\"bool\"}],\"name\":\"unlockWithRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inputData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"withdrawFromPolygon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

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
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// BATCHLIMIT is a free data retrieval call binding the contract method 0x9559c0bd.
//
// Solidity: function BATCH_LIMIT() view returns(uint256)
func (_Token *TokenCaller) BATCHLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BATCH_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BATCHLIMIT is a free data retrieval call binding the contract method 0x9559c0bd.
//
// Solidity: function BATCH_LIMIT() view returns(uint256)
func (_Token *TokenSession) BATCHLIMIT() (*big.Int, error) {
	return _Token.Contract.BATCHLIMIT(&_Token.CallOpts)
}

// BATCHLIMIT is a free data retrieval call binding the contract method 0x9559c0bd.
//
// Solidity: function BATCH_LIMIT() view returns(uint256)
func (_Token *TokenCallerSession) BATCHLIMIT() (*big.Int, error) {
	return _Token.Contract.BATCHLIMIT(&_Token.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Token *TokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Token *TokenSession) NAME() (string, error) {
	return _Token.Contract.NAME(&_Token.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Token *TokenCallerSession) NAME() (string, error) {
	return _Token.Contract.NAME(&_Token.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Token *TokenCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Token *TokenSession) VERSION() (string, error) {
	return _Token.Contract.VERSION(&_Token.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Token *TokenCallerSession) VERSION() (string, error) {
	return _Token.Contract.VERSION(&_Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Token *TokenCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Token *TokenSession) Exists(tokenId *big.Int) (bool, error) {
	return _Token.Contract.Exists(&_Token.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Token *TokenCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Token.Contract.Exists(&_Token.CallOpts, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Token *TokenCaller) Get(opts *bind.CallOpts, key string, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "get", key, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Token *TokenSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Token.Contract.Get(&_Token.CallOpts, key, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Token *TokenCallerSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Token.Contract.Get(&_Token.CallOpts, key, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.GetApproved(&_Token.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.GetApproved(&_Token.CallOpts, tokenId)
}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Token *TokenCaller) GetByHash(opts *bind.CallOpts, keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getByHash", keyHash, tokenId)

	outstruct := new(struct {
		Key   string
		Value string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Key = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Token *TokenSession) GetByHash(keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	return _Token.Contract.GetByHash(&_Token.CallOpts, keyHash, tokenId)
}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Token *TokenCallerSession) GetByHash(keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	return _Token.Contract.GetByHash(&_Token.CallOpts, keyHash, tokenId)
}

// GetKey is a free data retrieval call binding the contract method 0xbb5b27e1.
//
// Solidity: function getKey(uint256 keyHash) view returns(string)
func (_Token *TokenCaller) GetKey(opts *bind.CallOpts, keyHash *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getKey", keyHash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbb5b27e1.
//
// Solidity: function getKey(uint256 keyHash) view returns(string)
func (_Token *TokenSession) GetKey(keyHash *big.Int) (string, error) {
	return _Token.Contract.GetKey(&_Token.CallOpts, keyHash)
}

// GetKey is a free data retrieval call binding the contract method 0xbb5b27e1.
//
// Solidity: function getKey(uint256 keyHash) view returns(string)
func (_Token *TokenCallerSession) GetKey(keyHash *big.Int) (string, error) {
	return _Token.Contract.GetKey(&_Token.CallOpts, keyHash)
}

// GetKeys is a free data retrieval call binding the contract method 0xf5c1f76e.
//
// Solidity: function getKeys(uint256[] hashes) view returns(string[] values)
func (_Token *TokenCaller) GetKeys(opts *bind.CallOpts, hashes []*big.Int) ([]string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getKeys", hashes)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetKeys is a free data retrieval call binding the contract method 0xf5c1f76e.
//
// Solidity: function getKeys(uint256[] hashes) view returns(string[] values)
func (_Token *TokenSession) GetKeys(hashes []*big.Int) ([]string, error) {
	return _Token.Contract.GetKeys(&_Token.CallOpts, hashes)
}

// GetKeys is a free data retrieval call binding the contract method 0xf5c1f76e.
//
// Solidity: function getKeys(uint256[] hashes) view returns(string[] values)
func (_Token *TokenCallerSession) GetKeys(hashes []*big.Int) ([]string, error) {
	return _Token.Contract.GetKeys(&_Token.CallOpts, hashes)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Token *TokenCaller) GetMany(opts *bind.CallOpts, keys []string, tokenId *big.Int) ([]string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getMany", keys, tokenId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Token *TokenSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Token.Contract.GetMany(&_Token.CallOpts, keys, tokenId)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Token *TokenCallerSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Token.Contract.GetMany(&_Token.CallOpts, keys, tokenId)
}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Token *TokenCaller) GetManyByHash(opts *bind.CallOpts, keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getManyByHash", keyHashes, tokenId)

	outstruct := new(struct {
		Keys   []string
		Values []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[1], new([]string)).(*[]string)

	return *outstruct, err

}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Token *TokenSession) GetManyByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	return _Token.Contract.GetManyByHash(&_Token.CallOpts, keyHashes, tokenId)
}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Token *TokenCallerSession) GetManyByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	return _Token.Contract.GetManyByHash(&_Token.CallOpts, keyHashes, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Token *TokenCaller) IsApprovedOrOwner(opts *bind.CallOpts, spender common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isApprovedOrOwner", spender, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Token *TokenSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _Token.Contract.IsApprovedOrOwner(&_Token.CallOpts, spender, tokenId)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Token *TokenCallerSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _Token.Contract.IsApprovedOrOwner(&_Token.CallOpts, spender, tokenId)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Token *TokenCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Token *TokenSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Token.Contract.IsTrustedForwarder(&_Token.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Token *TokenCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Token.Contract.IsTrustedForwarder(&_Token.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Namehash is a free data retrieval call binding the contract method 0x276fabb1.
//
// Solidity: function namehash(string[] labels) pure returns(uint256)
func (_Token *TokenCaller) Namehash(opts *bind.CallOpts, labels []string) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "namehash", labels)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Namehash is a free data retrieval call binding the contract method 0x276fabb1.
//
// Solidity: function namehash(string[] labels) pure returns(uint256)
func (_Token *TokenSession) Namehash(labels []string) (*big.Int, error) {
	return _Token.Contract.Namehash(&_Token.CallOpts, labels)
}

// Namehash is a free data retrieval call binding the contract method 0x276fabb1.
//
// Solidity: function namehash(string[] labels) pure returns(uint256)
func (_Token *TokenCallerSession) Namehash(labels []string) (*big.Int, error) {
	return _Token.Contract.Namehash(&_Token.CallOpts, labels)
}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Token *TokenCaller) NonceOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "nonceOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Token *TokenSession) NonceOf(tokenId *big.Int) (*big.Int, error) {
	return _Token.Contract.NonceOf(&_Token.CallOpts, tokenId)
}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Token *TokenCallerSession) NonceOf(tokenId *big.Int) (*big.Int, error) {
	return _Token.Contract.NonceOf(&_Token.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, tokenId)
}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) ResolverOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "resolverOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Token *TokenSession) ResolverOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.ResolverOf(&_Token.CallOpts, tokenId)
}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) ResolverOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.ResolverOf(&_Token.CallOpts, tokenId)
}

// ReverseNameOf is a free data retrieval call binding the contract method 0xbebec6b4.
//
// Solidity: function reverseNameOf(address addr) view returns(string reverseUri)
func (_Token *TokenCaller) ReverseNameOf(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "reverseNameOf", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReverseNameOf is a free data retrieval call binding the contract method 0xbebec6b4.
//
// Solidity: function reverseNameOf(address addr) view returns(string reverseUri)
func (_Token *TokenSession) ReverseNameOf(addr common.Address) (string, error) {
	return _Token.Contract.ReverseNameOf(&_Token.CallOpts, addr)
}

// ReverseNameOf is a free data retrieval call binding the contract method 0xbebec6b4.
//
// Solidity: function reverseNameOf(address addr) view returns(string reverseUri)
func (_Token *TokenCallerSession) ReverseNameOf(addr common.Address) (string, error) {
	return _Token.Contract.ReverseNameOf(&_Token.CallOpts, addr)
}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256 reverse)
func (_Token *TokenCaller) ReverseOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "reverseOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256 reverse)
func (_Token *TokenSession) ReverseOf(addr common.Address) (*big.Int, error) {
	return _Token.Contract.ReverseOf(&_Token.CallOpts, addr)
}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256 reverse)
func (_Token *TokenCallerSession) ReverseOf(addr common.Address) (*big.Int, error) {
	return _Token.Contract.ReverseOf(&_Token.CallOpts, addr)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() pure returns(uint256)
func (_Token *TokenCaller) Root(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() pure returns(uint256)
func (_Token *TokenSession) Root() (*big.Int, error) {
	return _Token.Contract.Root(&_Token.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() pure returns(uint256)
func (_Token *TokenCallerSession) Root() (*big.Int, error) {
	return _Token.Contract.Root(&_Token.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Token.Contract.TokenURI(&_Token.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Token.Contract.TokenURI(&_Token.CallOpts, tokenId)
}

// Verify is a free data retrieval call binding the contract method 0xa4247400.
//
// Solidity: function verify((address,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Token *TokenCaller) Verify(opts *bind.CallOpts, req IForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xa4247400.
//
// Solidity: function verify((address,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Token *TokenSession) Verify(req IForwarderForwardRequest, signature []byte) (bool, error) {
	return _Token.Contract.Verify(&_Token.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xa4247400.
//
// Solidity: function verify((address,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Token *TokenCallerSession) Verify(req IForwarderForwardRequest, signature []byte) (bool, error) {
	return _Token.Contract.Verify(&_Token.CallOpts, req, signature)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Token *TokenTransactor) AddKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addKey", key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Token *TokenSession) AddKey(key string) (*types.Transaction, error) {
	return _Token.Contract.AddKey(&_Token.TransactOpts, key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Token *TokenTransactorSession) AddKey(key string) (*types.Transaction, error) {
	return _Token.Contract.AddKey(&_Token.TransactOpts, key)
}

// AddProxyReader is a paid mutator transaction binding the contract method 0x50960239.
//
// Solidity: function addProxyReader(address addr) returns()
func (_Token *TokenTransactor) AddProxyReader(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addProxyReader", addr)
}

// AddProxyReader is a paid mutator transaction binding the contract method 0x50960239.
//
// Solidity: function addProxyReader(address addr) returns()
func (_Token *TokenSession) AddProxyReader(addr common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddProxyReader(&_Token.TransactOpts, addr)
}

// AddProxyReader is a paid mutator transaction binding the contract method 0x50960239.
//
// Solidity: function addProxyReader(address addr) returns()
func (_Token *TokenTransactorSession) AddProxyReader(addr common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddProxyReader(&_Token.TransactOpts, addr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token *TokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, to, tokenId)
}

// BackfillReverseNames is a paid mutator transaction binding the contract method 0x44d5f66c.
//
// Solidity: function backfillReverseNames(string[][] domains) returns()
func (_Token *TokenTransactor) BackfillReverseNames(opts *bind.TransactOpts, domains [][]string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "backfillReverseNames", domains)
}

// BackfillReverseNames is a paid mutator transaction binding the contract method 0x44d5f66c.
//
// Solidity: function backfillReverseNames(string[][] domains) returns()
func (_Token *TokenSession) BackfillReverseNames(domains [][]string) (*types.Transaction, error) {
	return _Token.Contract.BackfillReverseNames(&_Token.TransactOpts, domains)
}

// BackfillReverseNames is a paid mutator transaction binding the contract method 0x44d5f66c.
//
// Solidity: function backfillReverseNames(string[][] domains) returns()
func (_Token *TokenTransactorSession) BackfillReverseNames(domains [][]string) (*types.Transaction, error) {
	return _Token.Contract.BackfillReverseNames(&_Token.TransactOpts, domains)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Token *TokenSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_Token *TokenTransactor) Deposit(opts *bind.TransactOpts, user common.Address, depositData []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deposit", user, depositData)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_Token *TokenSession) Deposit(user common.Address, depositData []byte) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, user, depositData)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_Token *TokenTransactorSession) Deposit(user common.Address, depositData []byte) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, user, depositData)
}

// DepositToPolygon is a paid mutator transaction binding the contract method 0x638e5c78.
//
// Solidity: function depositToPolygon(uint256 tokenId) returns()
func (_Token *TokenTransactor) DepositToPolygon(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "depositToPolygon", tokenId)
}

// DepositToPolygon is a paid mutator transaction binding the contract method 0x638e5c78.
//
// Solidity: function depositToPolygon(uint256 tokenId) returns()
func (_Token *TokenSession) DepositToPolygon(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DepositToPolygon(&_Token.TransactOpts, tokenId)
}

// DepositToPolygon is a paid mutator transaction binding the contract method 0x638e5c78.
//
// Solidity: function depositToPolygon(uint256 tokenId) returns()
func (_Token *TokenTransactorSession) DepositToPolygon(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DepositToPolygon(&_Token.TransactOpts, tokenId)
}

// Execute is a paid mutator transaction binding the contract method 0x1bf7e13e.
//
// Solidity: function execute((address,uint256,uint256,bytes) req, bytes signature) returns(bytes)
func (_Token *TokenTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1bf7e13e.
//
// Solidity: function execute((address,uint256,uint256,bytes) req, bytes signature) returns(bytes)
func (_Token *TokenSession) Execute(req IForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.Execute(&_Token.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1bf7e13e.
//
// Solidity: function execute((address,uint256,uint256,bytes) req, bytes signature) returns(bytes)
func (_Token *TokenTransactorSession) Execute(req IForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.Execute(&_Token.TransactOpts, req, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address mintingManager, address cnsRegistry, address rootChainManager, address childChainManager) returns()
func (_Token *TokenTransactor) Initialize(opts *bind.TransactOpts, mintingManager common.Address, cnsRegistry common.Address, rootChainManager common.Address, childChainManager common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initialize", mintingManager, cnsRegistry, rootChainManager, childChainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address mintingManager, address cnsRegistry, address rootChainManager, address childChainManager) returns()
func (_Token *TokenSession) Initialize(mintingManager common.Address, cnsRegistry common.Address, rootChainManager common.Address, childChainManager common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, mintingManager, cnsRegistry, rootChainManager, childChainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address mintingManager, address cnsRegistry, address rootChainManager, address childChainManager) returns()
func (_Token *TokenTransactorSession) Initialize(mintingManager common.Address, cnsRegistry common.Address, rootChainManager common.Address, childChainManager common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, mintingManager, cnsRegistry, rootChainManager, childChainManager)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 tokenId) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", user, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 tokenId) returns()
func (_Token *TokenSession) Mint(user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, user, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Mint(user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, user, tokenId)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address user, uint256 tokenId, bytes ) returns()
func (_Token *TokenTransactor) Mint0(opts *bind.TransactOpts, user common.Address, tokenId *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint0", user, tokenId, arg2)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address user, uint256 tokenId, bytes ) returns()
func (_Token *TokenSession) Mint0(user common.Address, tokenId *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Token.Contract.Mint0(&_Token.TransactOpts, user, tokenId, arg2)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address user, uint256 tokenId, bytes ) returns()
func (_Token *TokenTransactorSession) Mint0(user common.Address, tokenId *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Token.Contract.Mint0(&_Token.TransactOpts, user, tokenId, arg2)
}

// MintTLD is a paid mutator transaction binding the contract method 0xf7df5c60.
//
// Solidity: function mintTLD(uint256 tokenId, string uri) returns()
func (_Token *TokenTransactor) MintTLD(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintTLD", tokenId, uri)
}

// MintTLD is a paid mutator transaction binding the contract method 0xf7df5c60.
//
// Solidity: function mintTLD(uint256 tokenId, string uri) returns()
func (_Token *TokenSession) MintTLD(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Token.Contract.MintTLD(&_Token.TransactOpts, tokenId, uri)
}

// MintTLD is a paid mutator transaction binding the contract method 0xf7df5c60.
//
// Solidity: function mintTLD(uint256 tokenId, string uri) returns()
func (_Token *TokenTransactorSession) MintTLD(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Token.Contract.MintTLD(&_Token.TransactOpts, tokenId, uri)
}

// MintWithRecords is a paid mutator transaction binding the contract method 0xba5d40b7.
//
// Solidity: function mintWithRecords(address to, string[] labels, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenTransactor) MintWithRecords(opts *bind.TransactOpts, to common.Address, labels []string, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintWithRecords", to, labels, keys, values, withReverse)
}

// MintWithRecords is a paid mutator transaction binding the contract method 0xba5d40b7.
//
// Solidity: function mintWithRecords(address to, string[] labels, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenSession) MintWithRecords(to common.Address, labels []string, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.Contract.MintWithRecords(&_Token.TransactOpts, to, labels, keys, values, withReverse)
}

// MintWithRecords is a paid mutator transaction binding the contract method 0xba5d40b7.
//
// Solidity: function mintWithRecords(address to, string[] labels, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenTransactorSession) MintWithRecords(to common.Address, labels []string, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.Contract.MintWithRecords(&_Token.TransactOpts, to, labels, keys, values, withReverse)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Token *TokenTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "onERC721Received", arg0, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Token *TokenSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.OnERC721Received(&_Token.TransactOpts, arg0, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Token *TokenTransactorSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.OnERC721Received(&_Token.TransactOpts, arg0, from, tokenId, data)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Token *TokenTransactor) Reconfigure(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "reconfigure", keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Token *TokenSession) Reconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Reconfigure(&_Token.TransactOpts, keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Reconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Reconfigure(&_Token.TransactOpts, keys, values, tokenId)
}

// RemoveReverse is a paid mutator transaction binding the contract method 0xf25eb5c1.
//
// Solidity: function removeReverse() returns()
func (_Token *TokenTransactor) RemoveReverse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeReverse")
}

// RemoveReverse is a paid mutator transaction binding the contract method 0xf25eb5c1.
//
// Solidity: function removeReverse() returns()
func (_Token *TokenSession) RemoveReverse() (*types.Transaction, error) {
	return _Token.Contract.RemoveReverse(&_Token.TransactOpts)
}

// RemoveReverse is a paid mutator transaction binding the contract method 0xf25eb5c1.
//
// Solidity: function removeReverse() returns()
func (_Token *TokenTransactorSession) RemoveReverse() (*types.Transaction, error) {
	return _Token.Contract.RemoveReverse(&_Token.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Token *TokenTransactor) Reset(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "reset", tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Token *TokenSession) Reset(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Reset(&_Token.TransactOpts, tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Reset(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Reset(&_Token.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Token *TokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Token *TokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom0(&_Token.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom0(&_Token.TransactOpts, from, to, tokenId, data)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Token *TokenTransactor) Set(opts *bind.TransactOpts, key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "set", key, value, tokenId)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Token *TokenSession) Set(key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Set(&_Token.TransactOpts, key, value, tokenId)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Set(key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Set(&_Token.TransactOpts, key, value, tokenId)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetByHash is a paid mutator transaction binding the contract method 0x4a72584d.
//
// Solidity: function setByHash(uint256 keyHash, string value, uint256 tokenId) returns()
func (_Token *TokenTransactor) SetByHash(opts *bind.TransactOpts, keyHash *big.Int, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setByHash", keyHash, value, tokenId)
}

// SetByHash is a paid mutator transaction binding the contract method 0x4a72584d.
//
// Solidity: function setByHash(uint256 keyHash, string value, uint256 tokenId) returns()
func (_Token *TokenSession) SetByHash(keyHash *big.Int, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetByHash(&_Token.TransactOpts, keyHash, value, tokenId)
}

// SetByHash is a paid mutator transaction binding the contract method 0x4a72584d.
//
// Solidity: function setByHash(uint256 keyHash, string value, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SetByHash(keyHash *big.Int, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetByHash(&_Token.TransactOpts, keyHash, value, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Token *TokenTransactor) SetMany(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMany", keys, values, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Token *TokenSession) SetMany(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMany(&_Token.TransactOpts, keys, values, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SetMany(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMany(&_Token.TransactOpts, keys, values, tokenId)
}

// SetManyByHash is a paid mutator transaction binding the contract method 0x27f18975.
//
// Solidity: function setManyByHash(uint256[] keyHashes, string[] values, uint256 tokenId) returns()
func (_Token *TokenTransactor) SetManyByHash(opts *bind.TransactOpts, keyHashes []*big.Int, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setManyByHash", keyHashes, values, tokenId)
}

// SetManyByHash is a paid mutator transaction binding the contract method 0x27f18975.
//
// Solidity: function setManyByHash(uint256[] keyHashes, string[] values, uint256 tokenId) returns()
func (_Token *TokenSession) SetManyByHash(keyHashes []*big.Int, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetManyByHash(&_Token.TransactOpts, keyHashes, values, tokenId)
}

// SetManyByHash is a paid mutator transaction binding the contract method 0x27f18975.
//
// Solidity: function setManyByHash(uint256[] keyHashes, string[] values, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SetManyByHash(keyHashes []*big.Int, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetManyByHash(&_Token.TransactOpts, keyHashes, values, tokenId)
}

// SetOwner is a paid mutator transaction binding the contract method 0xab3b87fe.
//
// Solidity: function setOwner(address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) SetOwner(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setOwner", to, tokenId)
}

// SetOwner is a paid mutator transaction binding the contract method 0xab3b87fe.
//
// Solidity: function setOwner(address to, uint256 tokenId) returns()
func (_Token *TokenSession) SetOwner(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetOwner(&_Token.TransactOpts, to, tokenId)
}

// SetOwner is a paid mutator transaction binding the contract method 0xab3b87fe.
//
// Solidity: function setOwner(address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SetOwner(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetOwner(&_Token.TransactOpts, to, tokenId)
}

// SetReverse is a paid mutator transaction binding the contract method 0x384e9a55.
//
// Solidity: function setReverse(uint256 tokenId) returns()
func (_Token *TokenTransactor) SetReverse(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setReverse", tokenId)
}

// SetReverse is a paid mutator transaction binding the contract method 0x384e9a55.
//
// Solidity: function setReverse(uint256 tokenId) returns()
func (_Token *TokenSession) SetReverse(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetReverse(&_Token.TransactOpts, tokenId)
}

// SetReverse is a paid mutator transaction binding the contract method 0x384e9a55.
//
// Solidity: function setReverse(uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SetReverse(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetReverse(&_Token.TransactOpts, tokenId)
}

// SetReverse0 is a paid mutator transaction binding the contract method 0x663f7b2a.
//
// Solidity: function setReverse(string[] labels) returns()
func (_Token *TokenTransactor) SetReverse0(opts *bind.TransactOpts, labels []string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setReverse0", labels)
}

// SetReverse0 is a paid mutator transaction binding the contract method 0x663f7b2a.
//
// Solidity: function setReverse(string[] labels) returns()
func (_Token *TokenSession) SetReverse0(labels []string) (*types.Transaction, error) {
	return _Token.Contract.SetReverse0(&_Token.TransactOpts, labels)
}

// SetReverse0 is a paid mutator transaction binding the contract method 0x663f7b2a.
//
// Solidity: function setReverse(string[] labels) returns()
func (_Token *TokenTransactorSession) SetReverse0(labels []string) (*types.Transaction, error) {
	return _Token.Contract.SetReverse0(&_Token.TransactOpts, labels)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_Token *TokenTransactor) SetTokenURIPrefix(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTokenURIPrefix", prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_Token *TokenSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _Token.Contract.SetTokenURIPrefix(&_Token.TransactOpts, prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_Token *TokenTransactorSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _Token.Contract.SetTokenURIPrefix(&_Token.TransactOpts, prefix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// UnlockWithRecords is a paid mutator transaction binding the contract method 0x1f71be06.
//
// Solidity: function unlockWithRecords(address to, string[] labels, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenTransactor) UnlockWithRecords(opts *bind.TransactOpts, to common.Address, labels []string, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unlockWithRecords", to, labels, keys, values, withReverse)
}

// UnlockWithRecords is a paid mutator transaction binding the contract method 0x1f71be06.
//
// Solidity: function unlockWithRecords(address to, string[] labels, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenSession) UnlockWithRecords(to common.Address, labels []string, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.Contract.UnlockWithRecords(&_Token.TransactOpts, to, labels, keys, values, withReverse)
}

// UnlockWithRecords is a paid mutator transaction binding the contract method 0x1f71be06.
//
// Solidity: function unlockWithRecords(address to, string[] labels, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenTransactorSession) UnlockWithRecords(to common.Address, labels []string, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.Contract.UnlockWithRecords(&_Token.TransactOpts, to, labels, keys, values, withReverse)
}

// UnlockWithRecords0 is a paid mutator transaction binding the contract method 0xd106353f.
//
// Solidity: function unlockWithRecords(address to, uint256 tokenId, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenTransactor) UnlockWithRecords0(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unlockWithRecords0", to, tokenId, keys, values, withReverse)
}

// UnlockWithRecords0 is a paid mutator transaction binding the contract method 0xd106353f.
//
// Solidity: function unlockWithRecords(address to, uint256 tokenId, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenSession) UnlockWithRecords0(to common.Address, tokenId *big.Int, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.Contract.UnlockWithRecords0(&_Token.TransactOpts, to, tokenId, keys, values, withReverse)
}

// UnlockWithRecords0 is a paid mutator transaction binding the contract method 0xd106353f.
//
// Solidity: function unlockWithRecords(address to, uint256 tokenId, string[] keys, string[] values, bool withReverse) returns()
func (_Token *TokenTransactorSession) UnlockWithRecords0(to common.Address, tokenId *big.Int, keys []string, values []string, withReverse bool) (*types.Transaction, error) {
	return _Token.Contract.UnlockWithRecords0(&_Token.TransactOpts, to, tokenId, keys, values, withReverse)
}

// WithdrawFromPolygon is a paid mutator transaction binding the contract method 0x9508b1c4.
//
// Solidity: function withdrawFromPolygon(bytes inputData, uint256 tokenId, string[] keys, string[] values) returns()
func (_Token *TokenTransactor) WithdrawFromPolygon(opts *bind.TransactOpts, inputData []byte, tokenId *big.Int, keys []string, values []string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawFromPolygon", inputData, tokenId, keys, values)
}

// WithdrawFromPolygon is a paid mutator transaction binding the contract method 0x9508b1c4.
//
// Solidity: function withdrawFromPolygon(bytes inputData, uint256 tokenId, string[] keys, string[] values) returns()
func (_Token *TokenSession) WithdrawFromPolygon(inputData []byte, tokenId *big.Int, keys []string, values []string) (*types.Transaction, error) {
	return _Token.Contract.WithdrawFromPolygon(&_Token.TransactOpts, inputData, tokenId, keys, values)
}

// WithdrawFromPolygon is a paid mutator transaction binding the contract method 0x9508b1c4.
//
// Solidity: function withdrawFromPolygon(bytes inputData, uint256 tokenId, string[] keys, string[] values) returns()
func (_Token *TokenTransactorSession) WithdrawFromPolygon(inputData []byte, tokenId *big.Int, keys []string, values []string) (*types.Transaction, error) {
	return _Token.Contract.WithdrawFromPolygon(&_Token.TransactOpts, inputData, tokenId, keys, values)
}

// TokenAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Token contract.
type TokenAdminChangedIterator struct {
	Event *TokenAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminChanged)
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
		it.Event = new(TokenAdminChanged)
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
func (it *TokenAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdminChanged represents a AdminChanged event raised by the Token contract.
type TokenAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Token *TokenFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TokenAdminChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TokenAdminChangedIterator{contract: _Token.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Token *TokenFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdminChanged)
				if err := _Token.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Token *TokenFilterer) ParseAdminChanged(log types.Log) (*TokenAdminChanged, error) {
	event := new(TokenAdminChanged)
	if err := _Token.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TokenApprovalIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Token contract.
type TokenApprovalForAllIterator struct {
	Event *TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApprovalForAll)
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
		it.Event = new(TokenApprovalForAll)
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
func (it *TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApprovalForAll represents a ApprovalForAll event raised by the Token contract.
type TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalForAllIterator{contract: _Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApprovalForAll)
				if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Token *TokenFilterer) ParseApprovalForAll(log types.Log) (*TokenApprovalForAll, error) {
	event := new(TokenApprovalForAll)
	if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Token contract.
type TokenInitializedIterator struct {
	Event *TokenInitialized // Event containing the contract specifics and raw log

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
func (it *TokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenInitialized)
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
		it.Event = new(TokenInitialized)
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
func (it *TokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenInitialized represents a Initialized event raised by the Token contract.
type TokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Token *TokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenInitializedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenInitializedIterator{contract: _Token.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Token *TokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenInitialized) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenInitialized)
				if err := _Token.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Token *TokenFilterer) ParseInitialized(log types.Log) (*TokenInitialized, error) {
	event := new(TokenInitialized)
	if err := _Token.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenNewKeyIterator is returned from FilterNewKey and is used to iterate over the raw logs and unpacked data for NewKey events raised by the Token contract.
type TokenNewKeyIterator struct {
	Event *TokenNewKey // Event containing the contract specifics and raw log

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
func (it *TokenNewKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNewKey)
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
		it.Event = new(TokenNewKey)
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
func (it *TokenNewKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNewKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNewKey represents a NewKey event raised by the Token contract.
type TokenNewKey struct {
	TokenId  *big.Int
	KeyIndex common.Hash
	Key      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewKey is a free log retrieval operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Token *TokenFilterer) FilterNewKey(opts *bind.FilterOpts, tokenId []*big.Int, keyIndex []string) (*TokenNewKeyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "NewKey", tokenIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return &TokenNewKeyIterator{contract: _Token.contract, event: "NewKey", logs: logs, sub: sub}, nil
}

// WatchNewKey is a free log subscription operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Token *TokenFilterer) WatchNewKey(opts *bind.WatchOpts, sink chan<- *TokenNewKey, tokenId []*big.Int, keyIndex []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "NewKey", tokenIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNewKey)
				if err := _Token.contract.UnpackLog(event, "NewKey", log); err != nil {
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

// ParseNewKey is a log parse operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Token *TokenFilterer) ParseNewKey(log types.Log) (*TokenNewKey, error) {
	event := new(TokenNewKey)
	if err := _Token.contract.UnpackLog(event, "NewKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenNewURIIterator is returned from FilterNewURI and is used to iterate over the raw logs and unpacked data for NewURI events raised by the Token contract.
type TokenNewURIIterator struct {
	Event *TokenNewURI // Event containing the contract specifics and raw log

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
func (it *TokenNewURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNewURI)
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
		it.Event = new(TokenNewURI)
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
func (it *TokenNewURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNewURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNewURI represents a NewURI event raised by the Token contract.
type TokenNewURI struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewURI is a free log retrieval operation binding the contract event 0xc5beef08f693b11c316c0c8394a377a0033c9cf701b8cd8afd79cecef60c3952.
//
// Solidity: event NewURI(uint256 indexed tokenId, string uri)
func (_Token *TokenFilterer) FilterNewURI(opts *bind.FilterOpts, tokenId []*big.Int) (*TokenNewURIIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "NewURI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenNewURIIterator{contract: _Token.contract, event: "NewURI", logs: logs, sub: sub}, nil
}

// WatchNewURI is a free log subscription operation binding the contract event 0xc5beef08f693b11c316c0c8394a377a0033c9cf701b8cd8afd79cecef60c3952.
//
// Solidity: event NewURI(uint256 indexed tokenId, string uri)
func (_Token *TokenFilterer) WatchNewURI(opts *bind.WatchOpts, sink chan<- *TokenNewURI, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "NewURI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNewURI)
				if err := _Token.contract.UnpackLog(event, "NewURI", log); err != nil {
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

// ParseNewURI is a log parse operation binding the contract event 0xc5beef08f693b11c316c0c8394a377a0033c9cf701b8cd8afd79cecef60c3952.
//
// Solidity: event NewURI(uint256 indexed tokenId, string uri)
func (_Token *TokenFilterer) ParseNewURI(log types.Log) (*TokenNewURI, error) {
	event := new(TokenNewURI)
	if err := _Token.contract.UnpackLog(event, "NewURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenNewURIPrefixIterator is returned from FilterNewURIPrefix and is used to iterate over the raw logs and unpacked data for NewURIPrefix events raised by the Token contract.
type TokenNewURIPrefixIterator struct {
	Event *TokenNewURIPrefix // Event containing the contract specifics and raw log

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
func (it *TokenNewURIPrefixIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNewURIPrefix)
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
		it.Event = new(TokenNewURIPrefix)
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
func (it *TokenNewURIPrefixIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNewURIPrefixIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNewURIPrefix represents a NewURIPrefix event raised by the Token contract.
type TokenNewURIPrefix struct {
	Prefix string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewURIPrefix is a free log retrieval operation binding the contract event 0x4b120d6a959a84a520fa48f5f937cca0e79129423487af7901213b5d2e89313b.
//
// Solidity: event NewURIPrefix(string prefix)
func (_Token *TokenFilterer) FilterNewURIPrefix(opts *bind.FilterOpts) (*TokenNewURIPrefixIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "NewURIPrefix")
	if err != nil {
		return nil, err
	}
	return &TokenNewURIPrefixIterator{contract: _Token.contract, event: "NewURIPrefix", logs: logs, sub: sub}, nil
}

// WatchNewURIPrefix is a free log subscription operation binding the contract event 0x4b120d6a959a84a520fa48f5f937cca0e79129423487af7901213b5d2e89313b.
//
// Solidity: event NewURIPrefix(string prefix)
func (_Token *TokenFilterer) WatchNewURIPrefix(opts *bind.WatchOpts, sink chan<- *TokenNewURIPrefix) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "NewURIPrefix")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNewURIPrefix)
				if err := _Token.contract.UnpackLog(event, "NewURIPrefix", log); err != nil {
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

// ParseNewURIPrefix is a log parse operation binding the contract event 0x4b120d6a959a84a520fa48f5f937cca0e79129423487af7901213b5d2e89313b.
//
// Solidity: event NewURIPrefix(string prefix)
func (_Token *TokenFilterer) ParseNewURIPrefix(log types.Log) (*TokenNewURIPrefix, error) {
	event := new(TokenNewURIPrefix)
	if err := _Token.contract.UnpackLog(event, "NewURIPrefix", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRemoveReverseIterator is returned from FilterRemoveReverse and is used to iterate over the raw logs and unpacked data for RemoveReverse events raised by the Token contract.
type TokenRemoveReverseIterator struct {
	Event *TokenRemoveReverse // Event containing the contract specifics and raw log

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
func (it *TokenRemoveReverseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRemoveReverse)
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
		it.Event = new(TokenRemoveReverse)
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
func (it *TokenRemoveReverseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRemoveReverseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRemoveReverse represents a RemoveReverse event raised by the Token contract.
type TokenRemoveReverse struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemoveReverse is a free log retrieval operation binding the contract event 0xfcf5eec0cfa3e6332f5f0e63ec242d71f866a61d121d6cdf5c2eb3b668a26c4f.
//
// Solidity: event RemoveReverse(address indexed addr)
func (_Token *TokenFilterer) FilterRemoveReverse(opts *bind.FilterOpts, addr []common.Address) (*TokenRemoveReverseIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RemoveReverse", addrRule)
	if err != nil {
		return nil, err
	}
	return &TokenRemoveReverseIterator{contract: _Token.contract, event: "RemoveReverse", logs: logs, sub: sub}, nil
}

// WatchRemoveReverse is a free log subscription operation binding the contract event 0xfcf5eec0cfa3e6332f5f0e63ec242d71f866a61d121d6cdf5c2eb3b668a26c4f.
//
// Solidity: event RemoveReverse(address indexed addr)
func (_Token *TokenFilterer) WatchRemoveReverse(opts *bind.WatchOpts, sink chan<- *TokenRemoveReverse, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RemoveReverse", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRemoveReverse)
				if err := _Token.contract.UnpackLog(event, "RemoveReverse", log); err != nil {
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

// ParseRemoveReverse is a log parse operation binding the contract event 0xfcf5eec0cfa3e6332f5f0e63ec242d71f866a61d121d6cdf5c2eb3b668a26c4f.
//
// Solidity: event RemoveReverse(address indexed addr)
func (_Token *TokenFilterer) ParseRemoveReverse(log types.Log) (*TokenRemoveReverse, error) {
	event := new(TokenRemoveReverse)
	if err := _Token.contract.UnpackLog(event, "RemoveReverse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenResetRecordsIterator is returned from FilterResetRecords and is used to iterate over the raw logs and unpacked data for ResetRecords events raised by the Token contract.
type TokenResetRecordsIterator struct {
	Event *TokenResetRecords // Event containing the contract specifics and raw log

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
func (it *TokenResetRecordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenResetRecords)
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
		it.Event = new(TokenResetRecords)
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
func (it *TokenResetRecordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenResetRecordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenResetRecords represents a ResetRecords event raised by the Token contract.
type TokenResetRecords struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResetRecords is a free log retrieval operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterResetRecords(opts *bind.FilterOpts, tokenId []*big.Int) (*TokenResetRecordsIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ResetRecords", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenResetRecordsIterator{contract: _Token.contract, event: "ResetRecords", logs: logs, sub: sub}, nil
}

// WatchResetRecords is a free log subscription operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchResetRecords(opts *bind.WatchOpts, sink chan<- *TokenResetRecords, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ResetRecords", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenResetRecords)
				if err := _Token.contract.UnpackLog(event, "ResetRecords", log); err != nil {
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

// ParseResetRecords is a log parse operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Token *TokenFilterer) ParseResetRecords(log types.Log) (*TokenResetRecords, error) {
	event := new(TokenResetRecords)
	if err := _Token.contract.UnpackLog(event, "ResetRecords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Token contract.
type TokenSetIterator struct {
	Event *TokenSet // Event containing the contract specifics and raw log

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
func (it *TokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSet)
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
		it.Event = new(TokenSet)
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
func (it *TokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSet represents a Set event raised by the Token contract.
type TokenSet struct {
	TokenId    *big.Int
	KeyIndex   common.Hash
	ValueIndex common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Token *TokenFilterer) FilterSet(opts *bind.FilterOpts, tokenId []*big.Int, keyIndex []string, valueIndex []string) (*TokenSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}
	var valueIndexRule []interface{}
	for _, valueIndexItem := range valueIndex {
		valueIndexRule = append(valueIndexRule, valueIndexItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Set", tokenIdRule, keyIndexRule, valueIndexRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetIterator{contract: _Token.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Token *TokenFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *TokenSet, tokenId []*big.Int, keyIndex []string, valueIndex []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}
	var valueIndexRule []interface{}
	for _, valueIndexItem := range valueIndex {
		valueIndexRule = append(valueIndexRule, valueIndexItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Set", tokenIdRule, keyIndexRule, valueIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSet)
				if err := _Token.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Token *TokenFilterer) ParseSet(log types.Log) (*TokenSet, error) {
	event := new(TokenSet)
	if err := _Token.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetReverseIterator is returned from FilterSetReverse and is used to iterate over the raw logs and unpacked data for SetReverse events raised by the Token contract.
type TokenSetReverseIterator struct {
	Event *TokenSetReverse // Event containing the contract specifics and raw log

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
func (it *TokenSetReverseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetReverse)
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
		it.Event = new(TokenSetReverse)
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
func (it *TokenSetReverseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetReverseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetReverse represents a SetReverse event raised by the Token contract.
type TokenSetReverse struct {
	Addr    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetReverse is a free log retrieval operation binding the contract event 0xeb76a21470988c474a21f690cc28fee1ed511bd812dc3c21fd0f49c5e5d4708a.
//
// Solidity: event SetReverse(address indexed addr, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterSetReverse(opts *bind.FilterOpts, addr []common.Address, tokenId []*big.Int) (*TokenSetReverseIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetReverse", addrRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetReverseIterator{contract: _Token.contract, event: "SetReverse", logs: logs, sub: sub}, nil
}

// WatchSetReverse is a free log subscription operation binding the contract event 0xeb76a21470988c474a21f690cc28fee1ed511bd812dc3c21fd0f49c5e5d4708a.
//
// Solidity: event SetReverse(address indexed addr, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchSetReverse(opts *bind.WatchOpts, sink chan<- *TokenSetReverse, addr []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetReverse", addrRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetReverse)
				if err := _Token.contract.UnpackLog(event, "SetReverse", log); err != nil {
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

// ParseSetReverse is a log parse operation binding the contract event 0xeb76a21470988c474a21f690cc28fee1ed511bd812dc3c21fd0f49c5e5d4708a.
//
// Solidity: event SetReverse(address indexed addr, uint256 indexed tokenId)
func (_Token *TokenFilterer) ParseSetReverse(log types.Log) (*TokenSetReverse, error) {
	event := new(TokenSetReverse)
	if err := _Token.contract.UnpackLog(event, "SetReverse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TokenTransferIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Token contract.
type TokenUpgradedIterator struct {
	Event *TokenUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUpgraded)
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
		it.Event = new(TokenUpgraded)
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
func (it *TokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUpgraded represents a Upgraded event raised by the Token contract.
type TokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Token *TokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TokenUpgradedIterator{contract: _Token.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Token *TokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUpgraded)
				if err := _Token.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Token *TokenFilterer) ParseUpgraded(log types.Log) (*TokenUpgraded, error) {
	event := new(TokenUpgraded)
	if err := _Token.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
