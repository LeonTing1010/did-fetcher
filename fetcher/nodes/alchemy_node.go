package nodes

import (
	"context"
	"encoding/json"
	"fmt"
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
	POLYGON  = "POLYGON"
	ETHEREUM = "ETHEREUM"
	Arbitrum = "Arbitrum"
)

type Alchemy struct {
	Node
	ws  string
	url string
	key string
}

func NewAlchemy(project, chain string) *Alchemy {
	alchemy := Alchemy{
		Node: Node{
			chain:     chain,
			project:   project,
			step:      config.MaxBlock(chain),
			rawclient: rawClient(chain),
		},
		ws:  config.AlchemyWS(chain),
		url: config.AlchemyURI(chain),
		key: config.AlchemyKey(chain),
	}
	blockNumber, err := alchemy.rawclient.BlockNumber(context.Background())
	if err != nil {
		log.Panic("can't get blocknumber: ", err.Error())
	}
	alchemy.Node.blockNumber = blockNumber
	return &alchemy
}

func (alchemy *Alchemy) WaitTime() time.Duration {
	if alchemy.chain == config.POLYGON {
		return 100 * time.Millisecond
	}
	return time.Duration(7) * time.Second
}

func (alchemy *Alchemy) BlockNumber() uint64 {
	return alchemy.blockNumber
}

func (alchemy *Alchemy) Chain() string {
	return alchemy.chain
}

func (alchemy *Alchemy) Step() uint64 {
	return alchemy.step
}

func (alchemy *Alchemy) Info() string {
	return alchemy.project + ":" + alchemy.chain + ":" + strconv.FormatUint(alchemy.step, 10) + ":" + alchemy.WaitTime().String()
}

func rawClient(chain string) *ethclient.Client {
	url := config.PublicURI(chain)
	// if chain == Arbitrum {
	// 	url = config.AlchemyURI(chain) + "/v2/" + config.AlchemyKey(chain)
	// }
	cc, err := ethclient.Dial(url)
	log.Println("Dial alchemy client", url)
	if err != nil {
		log.Panicln("Can't Dial ethclient", url)
	}
	return cc
}

func (alchemy *Alchemy) EthClient() *ethclient.Client {
	return alchemy.rawclient
}

// func (proxy *Alchemy) GetNewURIs(project string, from uint64, to uint64, contract string) ([]NewURIEvent, []TransferEvent, error) {
// 	logs := proxy.GetTokenLogs(contract, from, to)
// 	token := contracts.New(project, contract, proxy.EthClient())
// 	var transferEvents []TransferEvent
// 	var newURIEvents []NewURIEvent
// 	for _, vLog := range logs {
// 		newTransferEvent, err := token.ParseTransfer(vLog)
// 		if err == nil {
// 			transferEvents = append(transferEvents, TransferEvent{
// 				BlockNumber:      vLog.BlockNumber,
// 				TransactionHash:  vLog.TxHash.Hex(),
// 				TransactionIndex: vLog.TxIndex,
// 				BlockHash:        vLog.BlockHash.Hex(),
// 				LogIndex:         vLog.Index,
// 				Contract:         contract,
// 				From:             newTransferEvent.From.Hex(),
// 				To:               newTransferEvent.To.Hex(),
// 				TokenId:          newTransferEvent.TokenId,
// 				Event:            "Transfer",
// 			})
// 			// log.Println("Event:", newTransferEvent.TokenId.String())
// 		}
// 		newURIEvent, err := token.ParseNewURI(vLog)
// 		if err == nil && newURIEvent.TokenId != nil {
// 			newURIEvents = append(newURIEvents, NewURIEvent{
// 				BlockNumber:      vLog.BlockNumber,
// 				TransactionHash:  vLog.TxHash.Hex(),
// 				TransactionIndex: vLog.TxIndex,
// 				BlockHash:        vLog.BlockHash.Hex(),
// 				LogIndex:         vLog.Index,
// 				Contract:         contract,
// 				Uri:              newURIEvent.Uri,
// 				TokenId:          newURIEvent.TokenId,
// 				From:             newURIEvent.From.Hex(),
// 				To:               newURIEvent.To.Hex(),
// 			})
// 			// log.Println("URI:", newURIEvent.TokenId.String())
// 		}

// 	}
// 	return newURIEvents, transferEvents, nil
// }

/*
{
  "contract": {
    "address": "0x2a93c52e7b6e7054870758e15a1446e769edfb93"
  },
  "id": {
    "tokenId": "0x00020e2dd0720f37f7d153208ab3ed2d214c9795700417014ff7774282c7691d",
    "tokenMetadata": {
      "tokenType": "ERC721"
    }
  },
  "title": "legalparrot66.nft",
  "description": "A CNS or UNS blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
  "tokenUri": {
    "gateway": "https://metadata.ud-staging.com/metadata/3631553727889860662261680140427215369061523740011538286026797230179445021",
    "raw": "https://metadata.ud-staging.com/metadata/3631553727889860662261680140427215369061523740011538286026797230179445021"
  },
  "media": [
    {
      "gateway": "https://nft-cdn.alchemy.com/matic-mumbai/fd69376a5557f7b808a462732c6f3e65",
      "thumbnail": "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/matic-mumbai/fd69376a5557f7b808a462732c6f3e65",
      "raw": "https://storage.googleapis.com/dot-crypto-metadata-api/images/animals/parrot.svg",
      "format": "svg+xml",
      "bytes": 264839
    }
  ],
  "metadata": {
    "image": "https://storage.googleapis.com/dot-crypto-metadata-api/images/animals/parrot.svg",
    "external_url": "https://unstoppabledomains.com/search?searchTerm=legalparrot66.nft",
    "tokenId": "3631553727889860662261680140427215369061523740011538286026797230179445021",
    "namehash": "0x00020e2dd0720f37f7d153208ab3ed2d214c9795700417014ff7774282c7691d",
    "background_color": "4C47F7",
    "image_url": "https://storage.googleapis.com/dot-crypto-metadata-api/images/animals/parrot.svg",
    "name": "legalparrot66.nft",
    "description": "A CNS or UNS blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
    "attributes": [
      {
        "value": "nft",
        "trait_type": "Ending"
      },
      {
        "value": 2,
        "trait_type": "Level"
      },
      {
        "value": 13,
        "trait_type": "Length"
      },
      {
        "value": 0,
        "trait_type": "Subdomains"
      },
      {
        "value": "animal",
        "trait_type": "Type"
      },
      {
        "value": "alphanumeric",
        "trait_type": "Character Set"
      }
    ]
  },
  "timeLastUpdated": "2023-02-07T03:44:43.508Z",
  "contractMetadata": {
    "name": "Unstoppable Domains",
    "symbol": "UD",
    "tokenType": "ERC721",
    "contractDeployer": "0x6f8b4b73c48627debaf6b6f552f1b8e77db7a667",
    "deployedBlockNumber": 18956099,
    "openSea": {
      "lastIngestedAt": "2023-02-13T13:24:30.000Z"
    }
  }
}
*/

// https://polygon-mumbai.g.alchemy.com/nft/v2/{key}/getNFTMetadata?contractAddress=0xe785E82358879F061BC3dcAC6f0444462D4b5330&tokenId=44&refreshCache=false
func (proxy *Alchemy) Metadata(contract string, tokenId *big.Int) (Metadata, error) {
	did := proxy.Porject()
	token, err := contracts.New(did, contract, proxy.rawclient)
	if err != nil {
		log.Panicln("not found token", contract)
	}
	meta, err := token.Metadata(tokenId)
	if err != nil {
		log.Println("Metadata not found meta", contract, tokenId)
		return nil, err
	}
	return &NFTmeta{
		MId:      ID(proxy.chain, contract, tokenId.String()),
		MName:    meta.Name(),
		MTokenId: tokenId.String(),
		Contract: contract,
		meta:     meta,
	}, nil
	// url := proxy.url + "/nft/v2/" + proxy.key + "/getNFTMetadata?refreshCache=false&contractAddress=" + contract + "&tokenId=" + tokenId.String()
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	// req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36")
	// req.Header.Add("cookie", "_ga=GA1.1.1761117347.1679475453; userty.core.p.4d9b8a=__2VySWQiOiI3MDA0YzFmMzM2YzNhZmM3NTIwNmQ4NzNlNDk5YjI3MCJ9eyJ1c; _ga_BX9X1NXC6Q=GS1.1.1679475452.1.1.1679475466.0.0.0; userty.core.s.4d9b8a=__WQiOiI1MDZhMmRhMjVmZDE1YWQ0N2ZkZDgxODk0MzM0MWE1MCIsInN0IjoxNjc5NDc1NDYxNTI4LCJwdiI6MiwicmVhZHkiOnRydWUsIndzIjoie1wid1wiOjEyNTQsXCJoXCI6Njk1fSJ9eyJza; _hp2_id.4099325469=%7B%22userId%22%3A%228746356497587035%22%2C%22pageviewId%22%3A%226271230504837616%22%2C%22sessionId%22%3A%224610458286028434%22%2C%22identity%22%3Anull%2C%22trackerVersion%22%3A%224.0%22%7D")
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return nil, err
	// }
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return nil, err
	// }
	// var meta ALMeta
	// json.Unmarshal(data, &meta)
	// meta.Aid = ID(proxy.Chain(), contract, tokenId.String())
	// return &meta, nil
}

func (proxy *Alchemy) KeyInfo(contract string, tokenId *big.Int) KeyInfo {
	did := proxy.Porject()
	token, err := contracts.New(did, contract, proxy.rawclient)
	if err != nil {
		log.Panicln("not found token", contract)
	}
	owner, err := token.Owner(tokenId)
	if err != nil {
		log.Println("Owner", err)
	}
	return &CKeyInfo{
		contract: contract,
		tokenId:  tokenId,
		owner:    owner.Hex(),
		expire:   token.Expire(tokenId),
		ttl:      token.TTL(tokenId),
	}
}

// func (proxy *Alchemy) GetTokenTransfers(project string, from uint64, to uint64, contract string) ([]TransferEvent, error) {
// 	logs := proxy.GetTokenLogs(contract, from, to)
// 	token := contracts.New(project, contract, proxy.EthClient())
// 	var events []TransferEvent
// 	for _, vLog := range logs {
// 		newTransferEvent, err := token.ParseTransfer(vLog)
// 		if err != nil {
// 			// log.Printf("Vlog tx: %s err:%s", vLog.TxHash.Hex(), err.Error())
// 			continue
// 		}
// 		if newTransferEvent.To.Hex() != contract {
// 			// log.Printf("Vlog tx: %s err:%s", vLog.TxHash.Hex(), "Contract Deploy Event")
// 			continue
// 		}
// 		events = append(events, TransferEvent{
// 			BlockNumber:      vLog.BlockNumber,
// 			TransactionHash:  vLog.TxHash.Hex(),
// 			TransactionIndex: vLog.TxIndex,
// 			BlockHash:        vLog.BlockHash.Hex(),
// 			LogIndex:         vLog.Index,
// 			Contract:         contract,
// 			From:             newTransferEvent.From.Hex(),
// 			To:               newTransferEvent.To.Hex(),
// 			TokenId:          newTransferEvent.TokenId,
// 			Event:            "Transfer",
// 		})

// 	}
// 	return events, nil
// }

/*
	{
	  "address": "0x2a93c52e7b6e7054870758e15a1446e769edfb93",
	  "contractMetadata": {
	    "name": "Unstoppable Domains",
	    "symbol": "UD",
	    "tokenType": "ERC721",
	    "contractDeployer": "0x6f8b4b73c48627debaf6b6f552f1b8e77db7a667",
	    "deployedBlockNumber": 18956099,
	    "openSea": {
	      "lastIngestedAt": "2023-02-21T13:54:08.000Z"
	    }
	  }
	}
*/

type ALCMeta struct {
	Contract string `json:"address"`
	Metadata struct {
		Name                string
		Symbol              string
		TokenType           string `default:"ERC721"`
		ContractDeployer    string
		DeployedBlockNumber uint64
		OpenSea             struct {
			LastIngestedAt string
		}
	} `json:"contractMetadata"`
}

func (cm ALCMeta) DeployedBlockNumber() uint64 {
	return cm.Metadata.DeployedBlockNumber
}

func (cm ALCMeta) TokenType() string {
	return cm.Metadata.TokenType
}

func (cm ALCMeta) Address() string {
	return cm.Contract
}

func (cm ALCMeta) Name() string {
	return cm.Metadata.Name
}

func (cm ALCMeta) Symbol() string {
	return cm.Metadata.Symbol
}

func (cm *ALCMeta) SetDeployedBlockNumber(blocknumber uint64) {
	cm.Metadata.DeployedBlockNumber = blocknumber
}

func (proxy *Alchemy) ContractMetadata(contract string) (ContractMeta, error) {
	// https://polygon-mumbai.g.alchemy.com/nft/v2/{apiKey}/getContractMetadata
	req, _ := http.NewRequest("GET", proxy.url+"/nft/v2/"+proxy.key+"/getContractMetadata?contractAddress="+contract, nil)

	// Send req using http Client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	meta := new(ALCMeta)
	err = json.Unmarshal(body, meta)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return meta, nil
}

/*
	{
	  "jsonrpc": "2.0",
	  "id": 1,
	  "result": "0x25d0b4d"
	}
*/
// type RPC struct {
// 	Jsonrpc string
// 	Id      int
// 	Result  string
// }

// func (proxy *Alchemy) EthClient() *ethclient.Client {
// 	return proxy.rawclient
// }

func (proxy *Alchemy) Subscribe(handle Handle) {
	log.Println("Open Websocket:", proxy.Info())
	wss := proxy.ws + "/v2/" + proxy.key
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
			if proxy.Chain() == config.ETHEREUM {
				if latest.Uint64()%10 == 0 {
					handle(latest)
				}
			} else {
				if latest.Uint64()%100 == 0 {
					handle(latest)
				}
			}
		}
	}
}

type ALMeta struct {
	Aid      string `bson:"_id"`
	Contract struct {
		Address string
	}
	Tid struct {
		TokenId       string
		TokenMetadata struct {
			TokenType string // ERC712
		}
	} `json:"id"`
	Title       string
	Description string
	TokenUri    struct {
		Gateway string
		Raw     string
	}
	Media    []Media
	Metadata struct {
		Image            string
		External_url     string
		TokenId          string
		Namehash         string
		Background_color string
		Image_url        string
		Name             string
		Description      string
		Attributes       []Attribute
	}
	TimeLastUpdated string
}

type Media struct {
	Gateway   string
	Thumbnail string
	Raw       string
	Format    string
	Bytes     int64
}

type Attribute struct {
	Value      any
	Trait_type string
}

func (meta ALMeta) Id() string {
	return meta.Aid
}

func (meta ALMeta) ContractAddress() string {
	return meta.Contract.Address
}

func (meta ALMeta) TokenId() string {
	return meta.Tid.TokenId
}

func (meta ALMeta) Name() string {
	return meta.Title
}

type TokenId struct {
	TokenId string
}

type Token struct {
	Id TokenId
}
type Collection struct {
	Nfts      []Token `json:"nfts"`
	NextToken string
}

/*
	{
	  "nfts": [
	    {
	      "id": {
	        "tokenId": "0x00005ad920a55754ba08270224ffc2250d478361bb0d8677b9d10bd2fe280806"
	      }
	    }
	  ],
	  "nextToken": "0x00005dacd88014bfd7ec3967110c4c7f80b68e8d65504da1d7ad62f8a9439328"
	}
*/
func (proxy *Alchemy) GetNFTsForCollection(contract, startToken string, limit uint64) (*Collection, error) {
	url := proxy.url + "/nft/v2/" + proxy.key + "/getNFTsForCollection?contractAddress=" + contract + "&withMetadata=false&startToken=" + startToken + "&limit=" + fmt.Sprint(limit)

	req, _ := http.NewRequest("GET", url, nil)
	// Send req using http Client
	res, err := proxy.RLClient().Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	collection := new(Collection)
	err = json.Unmarshal(body, collection)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return collection, nil
}
