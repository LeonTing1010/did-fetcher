package contracts

import (
	"context"
	"errors"
	"log"
	"math/big"

	"fetcher/contracts/hkdid"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const HK = "HKDID"

type Hkdid struct {
	name     string
	contract *hkdid.Hkdid
	client   *ethclient.Client
	chainId  *big.Int
}

func NewKey(contract string, client *ethclient.Client) NFT {
	token, err := hkdid.NewHkdid(common.HexToAddress(contract), client)
	if err != nil {
		log.Fatal(err)
	}
	chainId, _ := client.ChainID(context.Background())
	return &Hkdid{
		name:     HK,
		contract: token,
		client:   client,
		chainId:  chainId,
	}
}

func (token *Hkdid) ParseNewURI(vlog types.Log) (*NewURI, error) {
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

func (token *Hkdid) ParseTransfer(vlog types.Log) (*TokenTransfer, error) {
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

func (token *Hkdid) Owner(tokenId *big.Int) (common.Address, error) {
	owner, err := token.contract.OwnerOf(nil, tokenId)
	if err != nil {
		log.Println("Owner:", err.Error(), tokenId)
		return *new(common.Address), err
	}
	return owner, nil
}

func (token *Hkdid) Name() string {
	name, err := token.contract.Name(nil)
	if err != nil {
		log.Println("Name:", err.Error())
		return ""
	}
	return name
}

func (token *Hkdid) Symbol() string {
	symbol, err := token.contract.Symbol(nil)
	if err != nil {
		log.Println("Symbol:", err.Error())
		return ""
	}
	return symbol
}

func (token *Hkdid) Expire(tokenId *big.Int) *big.Int {
	return common.Big0
}

func (token *Hkdid) ChainId() uint64 {
	return token.chainId.Uint64()
}

func (token *Hkdid) TTL(tokenId *big.Int) uint64 {
	return 0
}

func (token *Hkdid) ParseNameRenewed(vlog types.Log) (*NameRenewed, error) {
	return nil, errors.New("not support function")
}

func (token Hkdid) Metadata(tokenId *big.Int) (Metadata, error) {
	url, err := token.contract.TokenURI(nil, tokenId)
	if err != nil {
		return nil, err
	}
	return getMetadata(url)
}
