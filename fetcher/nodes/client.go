package nodes

import (
	"math/big"
	"time"

	"fetcher/contracts"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	Chain() string
	Step() uint64
	Info() string
	WaitTime() time.Duration
	BlockNumber() uint64
	ContractMetadata(contract string) (ContractMeta, error)
	Metadata(contract string, tokenId *big.Int) (Metadata, error)
	KeyInfo(contract string, tokenId *big.Int) KeyInfo
	EthClient() *ethclient.Client
	Subscribe(handle Handle)
	FetchLogs(contract string, from uint64, to uint64) ([]types.Log, error)
	GetNFTsForCollection(contract, startToken string, limit uint64) (*Collection, error)
}
type Handle func(*big.Int)

type Metadata interface {
	contracts.Metadata
	Id() string
	TokenId() string
	ContractAddress() string
}

type NFTmeta struct {
	MId      string `bson:"_id"`
	MName    string `bson:"name"`
	MTokenId string `bson:"tokenId"`
	Contract string
	meta     any
}

func (nft NFTmeta) Name() string {
	return nft.MName
}

func (nft NFTmeta) Id() string {
	return nft.MId
}

func (nft NFTmeta) TokenId() string {
	return nft.MTokenId
}

func (nft NFTmeta) ContractAddress() string {
	return nft.Contract
}
