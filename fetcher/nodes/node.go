package nodes

import (
	"context"
	"log"
	"math/big"
	"runtime/debug"
	"time"

	"fetcher/client"
	"fetcher/contracts"
	"fetcher/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/time/rate"
)

var rlclient *client.RLHTTPClient

func init() {
	rl := rate.NewLimiter(rate.Every(10*time.Second), 1000) // 1000 request every 10 seconds
	rlclient = client.NewClient(rl)
}

type Node struct {
	step        uint64
	blockNumber uint64
	waitTime    time.Duration
	project     string
	chain       string
	rawclient   *ethclient.Client
	rpc         Client
}

func New(did, chain string) *Node {
	var rpc Client
	if chain == CONFLUX {
		rpc = NewConflux(chain)
	} else if chain == PLATON {
		rpc = NewPlatON(chain)
	} else if chain == MOONBEAM {
		rpc = NewMoonbeam(chain)
	} else if chain == BSC {
		rpc = NewBSC(chain)
	} else {
		rpc = NewAlchemy(did, chain)
	}
	return &Node{
		project:     did,
		step:        rpc.Step(),
		blockNumber: rpc.BlockNumber(),
		chain:       chain,
		rawclient:   rpc.EthClient(),
		rpc:         rpc,
		waitTime:    rpc.WaitTime(),
	}
}

func (node *Node) RLClient() *client.RLHTTPClient {
	return rlclient
}

func (node *Node) WaitTime() time.Duration {
	return node.waitTime
}

func (node *Node) Porject() string {
	return node.project
}

func (node *Node) Info() string {
	return node.rpc.Info()
}

func (node *Node) Chain() string {
	return node.chain
}

func (node *Node) Step() uint64 {
	return node.step
}

func (node *Node) SetStep(step uint64) {
	node.step = step
}

func (node *Node) BlockNumber() uint64 {
	return node.blockNumber
}

func (node *Node) ContractMetadata(contract string) (ContractMeta, error) {
	return node.rpc.ContractMetadata(contract)
}

func (node *Node) Subscribe(handle Handle) {
	node.rpc.Subscribe(handle)
}

func (node *Node) filterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return node.rawclient.FilterLogs(context.Background(), query)
}

func (node *Node) FetchNFTsByContract(c models.Contract) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic error:", err)
			log.Panicln("stacktrace from panic:", string(debug.Stack()))
		}
	}()
	to := node.BlockNumber()
	log.Printf("Fetching => %s chain:%s from:%d to:%d", c.DID(), c.Chain, c.LatestBlockNumber, to)
	node.FetchNFTsFromLatest(c, c.LatestBlockNumber, to)
}

func (node *Node) FetchNFTsFromLatest(c models.Contract, from, latest uint64) {
	log.Println("FetchNFTsFromLatest => ", node.Info(), from, latest)
	logCh := make(chan []types.Log, 1024)
	step := node.Step()
	to := from + step
	if to > latest {
		to = latest
	}
	go func() {
		for ; to <= latest; to = to + step {
			logs, err := node.FetchLogs(c.Contract(), from, to)
			if err != nil {
				models.UpdateLastestBlock(c, to)
				time.Sleep(100 * node.WaitTime())
				log.Panicln("FetchNFTsFromLatest", err, from, to)
				continue
			}
			size := len(logs)
			if size > 0 {
				logCh <- logs
				log.Println("N", from, to, size, latest-to)
			} else {
				models.UpdateLastestBlock(c, to)
			}
			from = to
			time.Sleep(node.WaitTime())
		}
		close(logCh)
	}()
	go parseEvents(node, c, logCh)
}

func (node *Node) NFTmeta(chain, contract string, tokenId *big.Int) (Metadata, error) {
	mId := ID(chain, contract, tokenId.String())
	meta, err := FindNFTMetaById(mId)
	if err != nil {
		// log.Println("FindNFTMetaById ", err, mId)
		// if node.project == contracts.UD {
		// 	return UDMeta(tokenId.String())
		// }
		meta, err = node.rpc.Metadata(contract, tokenId)
		if err != nil {
			log.Println(tokenId, err.Error())
			return nil, err
		}
	}
	return meta, nil
}

func (node *Node) FetchKeyInfo(contract string, tokenId *big.Int) KeyInfo {
	return node.rpc.KeyInfo(contract, tokenId)
}

func (node *Node) FetchLogs(contract string, from uint64, to uint64) ([]types.Log, error) {
	// log.Printf("FetchLogs => %s form:%d to:%d", node.Info(), from, to)
	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := node.filterLogs(context.Background(), query)
	if err != nil {
		log.Printf("FetchLogs Error => %s form:%d to:%d %s", node.Info(), from, to, err.Error())
		return nil, err
	}
	return logs, nil
}

// func blockRange(err string) (from, to int64) {
// 	log.Fatalln(err)
// 	reg := regexp.MustCompile(`\[(?P<From>.+?), (?P<To>.+?)\]`)
// 	if reg != nil {
// 		results := reg.FindStringSubmatch(err)
// 		fromHex := strings.TrimPrefix(results[1], "0x")
// 		toHex := strings.TrimPrefix(results[2], "0x")
// 		from, _ = strconv.ParseInt(fromHex, 16, 64)
// 		to, _ = strconv.ParseInt(toHex, 16, 64)
// 		return from, to
// 	}
// 	return 0, 0
// }

func (node *Node) GetNFTsForCollection(contract, startToken string, limit uint64) (*Collection, error) {
	return node.rpc.GetNFTsForCollection(contract, startToken, limit)
}

func ID(chain, contract, tokenId string) string {
	return models.NFTID(chain, contract, tokenId)
}

func (node *Node) UpdateExpireAndOwner(contract string) {
	total := models.NFTCount()
	log.Println("Total", total)
	var limit int64 = 1000
	page := total / limit
	var i int64 = 0
	log.Println("page", page)
	for ; i <= page; i++ {
		nfts := models.FindNFTs(limit, i)
		log.Println("i", i, len(nfts), page-i)
		nm := make(map[string]models.NFT)
		for _, nft := range nfts {
			tokenId, ok := contracts.Hex2BigInt(nft.TokenId)
			if !ok {
				continue
			}
			keyInfo := node.FetchKeyInfo(contract, tokenId)
			nft.SetExpires(keyInfo.Expire().Uint64())
			nft.SetOwner(keyInfo.Owner())
			nm[nft.Id] = nft
			// log.Println("nft", nft.Id)
			time.Sleep(1 * time.Second)
		}
		models.UpdateNFTByMap(nm)
	}
	log.Println("Finish UpdateExpireAndOwner", i)
}
