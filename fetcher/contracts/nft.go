package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type NFT interface {
	ParseNewURI(vlog types.Log) (*NewURI, error)
	ParseTransfer(vlog types.Log) (*TokenTransfer, error)
	ParseNameRenewed(vlog types.Log) (*NameRenewed, error)
	Owner(tokenId *big.Int) (common.Address, error)
	Name() string
	Symbol() string
	ChainId() uint64
	Expire(tokenId *big.Int) *big.Int
	TTL(tokenId *big.Int) uint64
	Metadata(tokenId *big.Int) (Metadata, error)
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type NewURI struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Uri     string
}

type NameRenewed struct {
	TokenId  *big.Int
	Expiries *big.Int
}
