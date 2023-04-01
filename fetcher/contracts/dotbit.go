package contracts

import (
	"context"
	"errors"
	"log"
	"math/big"

	"fetcher/contracts/dotbit"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const BIT = "DOTBIT"

type Dotbit struct {
	name     string
	contract *dotbit.Token
	client   *ethclient.Client
	chainId  *big.Int
	address  string
}

func NewBIT(contract string, client *ethclient.Client) NFT {
	token, err := dotbit.NewToken(common.HexToAddress(contract), client)
	if err != nil {
		log.Fatal(err)
	}
	chainId, _ := client.ChainID(context.Background())
	return &Dotbit{
		name:     BIT,
		contract: token,
		client:   client,
		chainId:  chainId,
		address:  contract,
	}
}

func (token *Dotbit) ParseNewURI(vlog types.Log) (*NewURI, error) {
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

func (token *Dotbit) ParseTransfer(vlog types.Log) (*TokenTransfer, error) {
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

func (token *Dotbit) Owner(tokenId *big.Int) (common.Address, error) {
	owner, err := token.contract.OwnerOf(nil, tokenId)
	if err != nil {
		// log.Println("Owner:", err.Error(), tokenId, token.address)
		return *new(common.Address), err
	}
	return owner, nil
}

func (token *Dotbit) Name() string {
	name, err := token.contract.Name(nil)
	if err != nil {
		log.Println("Name:", err.Error())
		return ""
	}
	return name
}

func (token *Dotbit) Symbol() string {
	symbol, err := token.contract.Symbol(nil)
	if err != nil {
		log.Println("Symbol:", err.Error())
		return ""
	}
	return symbol
}

func (token *Dotbit) Expire(tokenId *big.Int) *big.Int {
	limit, err := token.contract.GetExpires(nil, tokenId)
	if err != nil {
		log.Println("Expire:", err.Error(), tokenId)
		return common.Big0
	}
	return limit
}

func (token *Dotbit) ChainId() uint64 {
	return token.chainId.Uint64()
}

func (token *Dotbit) TTL(tokenId *big.Int) uint64 {
	return 0
}

func (token *Dotbit) ParseNameRenewed(vlog types.Log) (*NameRenewed, error) {
	return nil, errors.New("not support function")
}

func (token Dotbit) Metadata(tokenId *big.Int) (Metadata, error) {
	url, err := token.contract.TokenURI(nil, tokenId)
	if err != nil {
		return nil, err
	}
	return getMetadata(url)
}
