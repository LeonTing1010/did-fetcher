package nodes

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"fetcher/config"
	"fetcher/contracts"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CONFLUX = "Conflux"
	XYZ     = "https://api.cnsdomains.xyz/v1/cns/"
)

var (
	CWS  = config.Env("CNS_WSS_URI")
	CURL = config.Env("CNS_API_URI")
)

type Conflux struct {
	Node
	ws  string
	url string
}

func NewConflux(chain string) *Conflux {
	conflux := Conflux{
		ws:  CWS,
		url: CURL,
	}
	conflux.chain = CONFLUX
	conflux.step = 1024
	conflux.rawclient = cfxClient()
	blockNumber, err := conflux.rawclient.BlockNumber(context.Background())
	if err != nil {
		log.Panic("can't get blocknumber: ", err.Error())
	}
	conflux.blockNumber = blockNumber
	return &conflux
}

func cfxClient() *ethclient.Client {
	url := CURL
	c, err := ethclient.Dial(url)
	// log.Println("Dial ethclient", url)
	if err != nil {
		log.Panicln("Can't Dial Conflux", url)
	}
	return c
}

func (cfx *Conflux) EthClient() *ethclient.Client {
	return cfx.rawclient
}

func (cfx *Conflux) WaitTime() time.Duration {
	return 1 * time.Second
}

func (cfx *Conflux) BlockNumber() uint64 {
	return cfx.blockNumber
}

func (cfx *Conflux) Chain() string {
	return cfx.chain
}

func (cfx *Conflux) Step() uint64 {
	return cfx.step
}

func (cfx *Conflux) Info() string {
	return cfx.chain + ":" + strconv.FormatUint(cfx.step, 10) + ":" + cfx.WaitTime().String()
}

func (cfx *Conflux) KeyInfo(contract string, tokenId *big.Int) KeyInfo {
	token, err := contracts.New(contracts.CNS, contract, cfx.rawclient)
	if err != nil {
		log.Panicln("not found token", contract)
	}
	owner, err := token.Owner(tokenId)
	if err != nil {
		log.Println("not found owner tokenId", contract, tokenId)
	}
	expire := token.Expire(tokenId)

	return &CKeyInfo{
		contract: contract,
		tokenId:  tokenId,
		owner:    owner.Hex(),
		expire:   expire,
		ttl:      token.TTL(tokenId),
	}
}

type CNMeta struct {
	Cid         string `bson:"_id"`
	Mname       string `json:"name" bson:"name"`
	Description string
	Image       string
	MtokenId    string `json:"tokenId" bson:"tokenId"`
	Contract    string
	Mowner      string `json:"owner" bson:"owner"`
	Mttl        uint64 `json:"ttl" bson:"ttl"`
	Mexpire     uint64 `json:"expire" bson:"expire"`
}

func (meta CNMeta) Id() string {
	return meta.Cid
}

func (meta CNMeta) TokenId() string {
	return meta.MtokenId
}

func (meta CNMeta) Name() string {
	return meta.Mname
}

func (meta CNMeta) ContractAddress() string {
	return meta.Contract
}

func (cfx *Conflux) Metadata(contract string, tokenId *big.Int) (Metadata, error) {
	url := XYZ + tokenId.String()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36")
	res, err := cfx.RLClient().Do(req)
	if err != nil {
		// log.Println(tokenId, err.Error())
		return nil, err
	}
	response, err := io.ReadAll(res.Body)
	if err != nil {
		// log.Println(tokenId, err.Error())
		return nil, err
	}
	var meta CNMeta
	err = json.Unmarshal(response, &meta)
	if err != nil {
		// log.Println(tokenId, err.Error(), string(response))
		return nil, err
	}
	return &NFTmeta{
		MId:      ID(cfx.chain, meta.ContractAddress(), meta.TokenId()),
		MName:    meta.Name(),
		MTokenId: tokenId.String(),
		Contract: contract,
		meta:     meta,
	}, nil
}

func (cfx *Conflux) ContractMetadata(contract string) (ContractMeta, error) {
	token, err := contracts.New(contracts.CNS, contract, cfx.rawclient)
	if err != nil {
		log.Panicln("ContractMetadata:", err.Error())
	}
	return &CMetadata{
		Csymbol:              contracts.CNS,
		Cname:                token.Name(),
		Contract:             contract,
		CtokenType:           "ERC721",
		CdeployedBlockNumber: 0,
	}, nil
}

func (cfx *Conflux) Subscribe(handle Handle) {
	log.Println("Open Websocket", cfx.Info())
	wss := cfx.ws
	client, err := ethclient.Dial(wss)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Panicln(err.Error())
			return
		case header := <-headers:
			latest := header.Number
			log.Println(latest)
			if latest.Uint64()%1000 == 0 {
				handle(header.Number)
			}

		}
	}
}

func (cfx *Conflux) GetNFTsForCollection(contract, startToken string, limit uint64) (*Collection, error) {
	return nil, errors.New("not support")
}

// func (cfx *Conflux) FetchLogs(contract string, from uint64, to uint64) ([]types.Log, error) {
// 	time.Sleep(1 * time.Second)
// 	return cfx.Node.FetchLogs(contract, from, to)
// }
