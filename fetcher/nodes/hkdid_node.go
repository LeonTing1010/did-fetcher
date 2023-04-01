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

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PLATON = "PlatON"
	HK_XYZ = "https://api.hashkey.id/did/api/nft/metadata/"
)

var (
	HKWS  = config.Env("HK_WSS_URI")
	HKURL = config.Env("HK_API_URI")
)

type PlatON struct {
	Node
	ws  string
	url string
}

func NewPlatON(chain string) *PlatON {
	platON := PlatON{
		ws:  HKWS,
		url: HKURL,
	}
	platON.chain = PLATON
	platON.step = 1024
	platON.rawclient = PlatONClient()
	blockNumber, err := platON.rawclient.BlockNumber(context.Background())
	if err != nil {
		log.Panic("can't get blocknumber: ", err.Error())
	}
	platON.blockNumber = blockNumber
	return &platON
}

func PlatONClient() *ethclient.Client {
	url := HKURL
	c, err := ethclient.Dial(url)
	// log.Println("Dial ethclient", url)
	if err != nil {
		log.Panicln("Can't Dial PlatON", url)
	}
	return c
}

func (hk *PlatON) EthClient() *ethclient.Client {
	return hk.rawclient
}

func (hk *PlatON) WaitTime() time.Duration {
	return 1 * time.Second
}

func (hk *PlatON) Chain() string {
	return hk.chain
}

func (hk *PlatON) BlockNumber() uint64 {
	return hk.blockNumber
}

func (hk *PlatON) Step() uint64 {
	return hk.step
}

func (hk *PlatON) Info() string {
	return hk.chain + ":" + strconv.FormatUint(hk.step, 10) + ":" + hk.WaitTime().String()
}

type HKMeta struct {
	Hid         string `bson:"_id"`
	Mname       string `json:"name" bson:"name"`
	Description string
	Image       string
	MtokenId    string `json:"tokenId" bson:"tokenId"`
	Contract    string
	Mowner      string `json:"owner" bson:"owner"`
	Ipfs        string `json:"ipfs_url"`
}

func (meta HKMeta) Id() string {
	return meta.Hid
}

func (meta HKMeta) TokenId() string {
	return meta.MtokenId
}

func (meta HKMeta) Name() string {
	return meta.Mname
}

func (meta HKMeta) ContractAddress() string {
	return meta.Contract
}

/*
https://api.hashkey.id/did/api/nft/metadata/7725944446619810869152252710829725001783725220341070053556238396080733066424

	{
		"description": "Your Passport in the Metaverse",
		"image": "https://api.hashkey.id/did/api/file/avatar_c375a663-2943-41bc-a431-48084218f2da.png",
		"name": "phong1368.key",
		"ipfs_url": ""
	}
*/
func (hk *PlatON) Metadata(contract string, tokenId *big.Int) (Metadata, error) {
	req, _ := http.NewRequest("GET", HK_XYZ+tokenId.String(), nil)
	if hk.RLClient() == nil {
		return nil, errors.New("not found http client")
	}
	getNFTMetadata, err := hk.RLClient().Do(req)
	if err != nil {
		// log.Println(tokenId, err.Error())
		return nil, err
	}
	response, err := io.ReadAll(getNFTMetadata.Body)
	if err != nil {
		// log.Println(tokenId, err.Error())
		return nil, err
	}
	var meta HKMeta
	err = json.Unmarshal(response, &meta)
	if err != nil {
		// log.Println(tokenId, err.Error(), string(response))
		return nil, err
	}
	return &NFTmeta{
		MId:      ID(hk.Chain(), contract, meta.TokenId()),
		MName:    meta.Name(),
		MTokenId: tokenId.String(),
		Contract: contract,
		meta:     meta,
	}, nil
}

func (hk *PlatON) ContractMetadata(contract string) (ContractMeta, error) {
	token, err := contracts.New(contracts.HK, contract, hk.rawclient)
	if err != nil {
		log.Panicln("ContractMetadata:", err.Error())
	}
	return &CMetadata{
		Csymbol:              contracts.HK,
		Cname:                token.Name(),
		Contract:             contract,
		CtokenType:           "ERC721",
		CdeployedBlockNumber: 0,
	}, nil
}

func (hk *PlatON) KeyInfo(contract string, tokenId *big.Int) KeyInfo {
	token, err := contracts.New(contracts.HK, contract, hk.rawclient)
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

func (hk *PlatON) Subscribe(handle Handle) {
	log.Println("Open Websocket:", hk.Info())
	wss := hk.ws
	client, err := ethclient.Dial(wss)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ticker := time.NewTicker(hk.WaitTime())
	defer ticker.Stop()

	bn := hk.BlockNumber()
	for {
		<-ticker.C
		if bn%100 == 0 {
			handle(big.NewInt(int64(bn)))
		}
		bn, _ = client.BlockNumber(context.Background())
	}
}

func (hk *PlatON) GetNFTsForCollection(contract, startToken string, limit uint64) (*Collection, error) {
	return nil, errors.New("not support")
}
