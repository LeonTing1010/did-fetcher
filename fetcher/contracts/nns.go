package contracts

import (
	"context"
	"log"
	"math/big"

	"fetcher/contracts/nns"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 0xb10F5664A65513CcA275cfcD4242eA8dC4010Efc BSC
const NNS = "NNS"

type Nns struct {
	name     string
	contract *nns.Nns
	client   *ethclient.Client
	chainId  *big.Int
}

func NewNNS(contract string, client *ethclient.Client) NFT {
	token, err := nns.NewNns(common.HexToAddress(contract), client)
	if err != nil {
		log.Fatal(err)
	}
	chainId, _ := client.ChainID(context.Background())
	return &Nns{
		name:     UD,
		contract: token,
		client:   client,
		chainId:  chainId,
	}
}

func (token *Nns) ParseNewURI(vlog types.Log) (*NewURI, error) {
	newURI, err := token.contract.ParseNameRegistered(vlog)
	if err != nil {
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
	return &NewURI{
		From:    common.Address{},
		To:      newURI.Owner,
		TokenId: newURI.Id,
		Uri:     "",
	}, nil
}

func (token *Nns) ParseTransfer(vlog types.Log) (*TokenTransfer, error) {
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

func (token *Nns) ParseNameRenewed(vlog types.Log) (*NameRenewed, error) {
	event, err := token.contract.ParseNameRenewed(vlog)
	if err != nil {
		return nil, err
	}
	return &NameRenewed{
		Expiries: event.Expires,
		TokenId:  event.Id,
	}, nil
}

func (token *Nns) Owner(tokenId *big.Int) (common.Address, error) {
	owner, err := token.contract.OwnerOf(nil, tokenId)
	if err != nil {
		log.Println("Owner", tokenId, err.Error())
		return *new(common.Address), err
	}
	return owner, nil
}

func (token *Nns) Name() string {
	name, err := token.contract.Name(nil)
	if err != nil {
		log.Println("Name", err.Error())
		return ""
	}
	return name
}

func (token *Nns) Symbol() string {
	symbol, err := token.contract.Symbol(nil)
	if err != nil {
		log.Println("Symbol", err.Error())
		return ""
	}
	return symbol
}

func (token *Nns) Expire(tokenId *big.Int) *big.Int {
	expires, err := token.contract.NameExpires(nil, tokenId)
	if err != nil {
		log.Println("Expire", err.Error())
		return common.Big0
	}
	return expires
}

func (token *Nns) ChainId() uint64 {
	return token.chainId.Uint64()
}

func (token *Nns) TTL(tokenId *big.Int) uint64 {
	return 0
}

func (token Nns) Metadata(tokenId *big.Int) (Metadata, error) {
	url, err := token.contract.TokenURI(nil, tokenId)
	if err != nil {
		return nil, err
	}
	return getMetadata(url)
}
