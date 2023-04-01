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
	BSC     = "BSC"
	nns_XYZ = "https://assets.nft.space/bsc/0xb10f5664a65513cca275cfcd4242ea8dc4010efc/"
)

var nnsURL = config.Env("NNS_API_URI")

type Bsc struct {
	Node
	url string
}

func NewBSC(chain string) *Bsc {
	bsc := Bsc{
		url: nnsURL,
	}
	bsc.chain = chain
	bsc.step = 256
	bsc.rawclient = bscClient()
	blockNumber, err := bsc.rawclient.BlockNumber(context.Background())
	if err != nil {
		log.Panic("can't get blocknumber: ", err.Error())
	}
	bsc.blockNumber = blockNumber
	return &bsc
}

func bscClient() *ethclient.Client {
	url := nnsURL
	c, err := ethclient.Dial(url)
	// log.Println("Dial ethclient", url)
	if err != nil {
		log.Panicln("Can't Dial Bsc", url)
	}
	return c
}

func (node *Bsc) EthClient() *ethclient.Client {
	return node.rawclient
}

func (node *Bsc) WaitTime() time.Duration {
	return 1 * time.Second
}

func (node *Bsc) Chain() string {
	return node.chain
}

func (node *Bsc) BlockNumber() uint64 {
	return node.blockNumber
}

func (node *Bsc) Step() uint64 {
	return node.step
}

func (node *Bsc) Info() string {
	return node.chain + ":" + strconv.FormatUint(node.step, 10) + ":" + node.WaitTime().String()
}

type NNSMeta struct {
	Nid         string `bson:"_id"`
	Mname       string `json:"name" bson:"name"`
	Description string
	Image       string
	MtokenId    string `json:"tokenId" bson:"tokenId"`
	Contract    string
	Mowner      string `json:"owner" bson:"owner"`
	Url         string `json:"url"`
}

func (meta NNSMeta) Id() string {
	return meta.Nid
}

func (meta NNSMeta) TokenId() string {
	return meta.MtokenId
}

func (meta NNSMeta) Name() string {
	return meta.Mname
}

func (meta NNSMeta) ContractAddress() string {
	return meta.Contract
}

/*
https://assets.nft.space/bsc/0xb10f5664a65513cca275cfcd4242ea8dc4010efc/48493857913400469574466960623845223681749249524334853482317768219357308951429

	{
	  "is_normalized": true,
	  "name": "sushiswap.meta",
	  "description": "sushiswap.meta, a NFT name.",
	  "attributes": [
	    {
	      "trait_type": "Created Date",
	      "display_type": "date",
	      "value": 1678901909
	    },
	    {
	      "trait_type": "Length",
	      "display_type": "number",
	      "value": 9
	    },
	    {
	      "trait_type": "Segment Length",
	      "display_type": "number",
	      "value": 9
	    },
	    {
	      "trait_type": "Character Set",
	      "display_type": "string",
	      "value": "letter"
	    },
	    {
	      "trait_type": "Registration Date",
	      "display_type": "date",
	      "value": 1678901909
	    },
	    {
	      "trait_type": "Expiration Date",
	      "display_type": "date",
	      "value": 1710437909
	    }
	  ],
	  "nameLength": 9,
	  "segment_length": 9,
	  "url": "https://nft.space/sushiswap.meta/details",
	  "version": 0,
	  "background_image": "https://assets.nft.space/bsc/0xb10f5664a65513cca275cfcd4242ea8dc4010efc/0x6b368d086df1aee085e2153c3563b8277337bab4b2aa409a33bf9238030e9b85/image",
	  "image": "https://assets.nft.space/bsc/0xb10f5664a65513cca275cfcd4242ea8dc4010efc/0x6b368d086df1aee085e2153c3563b8277337bab4b2aa409a33bf9238030e9b85/image",
	  "image_url": "https://assets.nft.space/bsc/0xb10f5664a65513cca275cfcd4242ea8dc4010efc/0x6b368d086df1aee085e2153c3563b8277337bab4b2aa409a33bf9238030e9b85/image"
	}
*/
func (node *Bsc) Metadata(contract string, tokenId *big.Int) (Metadata, error) {
	req, _ := http.NewRequest("GET", nns_XYZ+tokenId.String(), nil)
	getNFTMetadata, err := node.RLClient().Do(req)
	if err != nil {
		// log.Println(tokenId, err.Error())
		return nil, err
	}
	response, err := io.ReadAll(getNFTMetadata.Body)
	if err != nil {
		// log.Println(tokenId, err.Error())
		return nil, err
	}
	var meta NNSMeta
	err = json.Unmarshal(response, &meta)
	if err != nil {
		// log.Println(tokenId, err.Error(), string(response))
		return nil, err
	}
	return &NFTmeta{
		MId:      ID(node.Chain(), contract, meta.TokenId()),
		MName:    meta.Name(),
		MTokenId: tokenId.String(),
		Contract: contract,
		meta:     meta,
	}, nil
}

func (node *Bsc) ContractMetadata(contract string) (ContractMeta, error) {
	token, err := contracts.New(contracts.NNS, contract, node.rawclient)
	if err != nil {
		log.Panicln("ContractMetadata:", err.Error())
	}
	return &CMetadata{
		Csymbol:              contracts.NNS,
		Cname:                token.Name(),
		Contract:             contract,
		CtokenType:           "ERC721",
		CdeployedBlockNumber: 0,
	}, nil
}

func (node *Bsc) KeyInfo(contract string, tokenId *big.Int) KeyInfo {
	token, err := contracts.New(contracts.HK, contract, node.rawclient)
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

func (node *Bsc) Subscribe(handle Handle) {
	log.Println("Open Websocket:", node.Info())
	wss := node.url
	client, err := ethclient.Dial(wss)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ticker := time.NewTicker(node.WaitTime())
	defer ticker.Stop()

	bn := node.BlockNumber()
	for {
		<-ticker.C
		if bn%100 == 0 {
			handle(big.NewInt(int64(bn)))
		}
		bn, _ = client.BlockNumber(context.Background())
	}
}

func (node *Bsc) GetNFTsForCollection(contract, startToken string, limit uint64) (*Collection, error) {
	return nil, errors.New("not support")
}
