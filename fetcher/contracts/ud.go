package contracts

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"fetcher/contracts/ud"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/time/rate"
)

const UD = "UD"

type Unstoppable struct {
	name     string
	contract *ud.Token
	client   *ethclient.Client
	chainId  *big.Int
}

var Limiter *rate.Limiter

func init() {
	Limiter = rate.NewLimiter(rate.Every(10*time.Second), 10000) // 5000 request every 10 seconds
}

func NewUD(contract string, client *ethclient.Client) NFT {
	token, err := ud.NewToken(common.HexToAddress(contract), client)
	if err != nil {
		log.Fatal(err)
	}
	chainId, _ := client.ChainID(context.Background())
	return &Unstoppable{
		name:     UD,
		contract: token,
		client:   client,
		chainId:  chainId,
	}
}

func (token *Unstoppable) ParseNewURI(vlog types.Log) (*NewURI, error) {
	newURI, err := token.contract.ParseNewURI(vlog)
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
		To:      newURI.Raw.Address,
		TokenId: newURI.TokenId,
		Uri:     newURI.Uri,
	}, nil
}

func (token *Unstoppable) ParseTransfer(vlog types.Log) (*TokenTransfer, error) {
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

func (token *Unstoppable) Owner(tokenId *big.Int) (common.Address, error) {
	err := Limiter.Wait(context.Background()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return *new(common.Address), err
	}
	owner, err := token.contract.OwnerOf(nil, tokenId)
	if err != nil {
		log.Println("Owner", tokenId, err.Error())
		return *new(common.Address), err
	}
	return owner, nil
}

func (token *Unstoppable) Name() string {
	name, err := token.contract.Name(nil)
	if err != nil {
		log.Println("Name", err.Error())
		return ""
	}
	return name
}

func (token *Unstoppable) Symbol() string {
	symbol, err := token.contract.Symbol(nil)
	if err != nil {
		log.Println("Symbol", err.Error())
		return ""
	}
	return symbol
}

func (token *Unstoppable) Expire(tokenId *big.Int) *big.Int {
	return common.Big0
}

func (token *Unstoppable) ChainId() uint64 {
	return token.chainId.Uint64()
}

func (token *Unstoppable) TTL(tokenId *big.Int) uint64 {
	return 0
}

func (token *Unstoppable) ParseNameRenewed(vlog types.Log) (*NameRenewed, error) {
	return nil, errors.New("not support function")
}

func (token Unstoppable) Metadata(tokenId *big.Int) (Metadata, error) {
	err := Limiter.Wait(context.Background()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	url, err := token.contract.TokenURI(nil, tokenId)
	if err != nil {
		return nil, err
	}
	return getMetadata(url)
}
