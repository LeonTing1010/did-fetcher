package contracts

import (
	"context"
	"errors"
	"log"
	"math/big"

	"fetcher/contracts/cns"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const CNS = "CNS"

type Conflux struct {
	name     string
	contract *cns.Cns
	client   *ethclient.Client
	chainId  *big.Int
}

func NewCNS(contract string, client *ethclient.Client) NFT {
	token, err := cns.NewCns(common.HexToAddress(contract), client)
	if err != nil {
		log.Fatal(err)
	}
	chainId, _ := client.ChainID(context.Background())
	return &Conflux{
		name:     CNS,
		contract: token,
		client:   client,
		chainId:  chainId,
	}
}

func (token *Conflux) ParseNewURI(vlog types.Log) (*NewURI, error) {
	e, err := token.contract.ParseTransfer(vlog)
	if err != nil {
		return nil, err
	}
	if e.From.Hex() == ZERO { // Mint
		return &NewURI{
			From:    e.From,
			To:      e.To,
			TokenId: e.TokenId,
			Uri:     "",
		}, nil
	} else {
		return nil, err
	}
}

func (token *Conflux) ParseTransfer(vlog types.Log) (*TokenTransfer, error) {
	event, err := token.contract.ParseTransfer(vlog)
	if err != nil {
		return nil, err
	}
	return &TokenTransfer{
		From:    event.From,
		To:      event.To,
		TokenId: event.TokenId,
	}, nil
}

func (token *Conflux) Owner(tokenId *big.Int) (common.Address, error) {
	owner, err := token.contract.OwnerOf(nil, tokenId)
	if err != nil {
		log.Println("Owner", tokenId, err.Error())
		return *new(common.Address), err
	}
	return owner, nil
}

func (token *Conflux) Name() string {
	name, err := token.contract.Name(nil)
	if err != nil {
		log.Println("Name", err.Error())
		return ""
	}
	return name
}

func (token *Conflux) Symbol() string {
	symbol, err := token.contract.Symbol(nil)
	if err != nil {
		log.Println("Symbol", err.Error())
		return ""
	}
	return symbol
}

func (token *Conflux) Expire(tokenId *big.Int) *big.Int {
	date, err := token.contract.NameExpires(nil, tokenId)
	if err != nil {
		log.Println("Expire", err.Error())
		return common.Big0
	}
	return date
}

func (token *Conflux) ChainId() uint64 {
	return token.chainId.Uint64()
}

func (token *Conflux) TTL(tokenId *big.Int) uint64 {
	return 0
}

func (token *Conflux) ParseNameRenewed(vlog types.Log) (*NameRenewed, error) {
	return nil, errors.New("not support function")
}

func (token Conflux) Metadata(tokenId *big.Int) (Metadata, error) {
	url, err := token.contract.TokenURI(nil, tokenId)
	if err != nil {
		return nil, err
	}
	return getMetadata(url)
}
