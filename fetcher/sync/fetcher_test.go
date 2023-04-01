package sync

import (
	"log"
	"os"
	"sync"
	"testing"

	"fetcher/config"
	"fetcher/contracts"
	"fetcher/models"
	"fetcher/nodes"
	"fetcher/util"

	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/stretchr/testify/assert"
)

func TestFetchExpireAndOwner(t *testing.T) {
	t.Log("TestFetchMBNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "ETHEREUM")
	did := "DOTBIT"
	config.SetUp()
	models.Setup(did)
	FetchExpireAndOwner(did)
	select {}
}

func TestFetchNNSNFTs(t *testing.T) {
	t.Log("TestFetchMBNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "BSC")
	config.SetUp()
	did := "NNS"
	models.Setup(did)
	FetchHistory(did)
	// c, _ := models.FindContractByAddress("0xb10f5664a65513cca275cfcd4242ea8dc4010efc")
	// node := nodes.New(c.DID(), c.Blockchain())
	// c.LatestBlockNumber = 25516131
	// node.FetchNFTsFromLatest(c, 25516131, 25516131)
	select {}
}

func TestFetchMBNFTs(t *testing.T) {
	t.Log("TestFetchMBNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "Moonbeam")
	did := "HKDID"
	models.Setup(did)
	config.SetUp()
	FetchHistory(did)
	// FetchNewBlock("HKDID")
	select {}
}

func TestFetchHKNFTs(t *testing.T) {
	t.Log("TestFetchHKNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "PlatON")
	did := "HKDID"
	models.Setup(did)
	config.SetUp()
	FetchHistory(did)
	FetchNewBlock(did)
	select {}
}

func TestFetchCNSNFTs(t *testing.T) {
	t.Log("TestFetchCNSNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "Conflux")
	did := "CNS"
	models.Setup(did)
	config.SetUp()
	// FetchHistory(did)
	FetchNewBlock(did)
	select {}
}

func TestFetchUDNFTs(t *testing.T) {
	t.Log("TestFetchUDNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "POLYGON")
	config.SetUp()
	models.Setup("UD")
	// DID := contracts.UD
	// FetchContractMetadatas(DID)
	// FetchHistory(DID)
	c, _ := models.FindContractByAddress("0xa9a6a3626993d487d2dbda3173cf58ca1a9d9e9f")
	node := nodes.New(c.DID(), c.Blockchain())
	c.LatestBlockNumber = 21632176
	node.FetchNFTsFromLatest(c, 19345077, 21632176)
	select {}
}

func TestFetchBitNFTs(t *testing.T) {
	t.Log("TestFetchBitNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "ETHEREUM")
	models.Setup("DOTBIT")
	c, _ := models.FindContractByAddress("0x60eb332bd4a0e2a9eeb3212cfdd6ef03ce4cb3b5")
	node := nodes.New(c.DID(), c.Blockchain())
	node.FetchNFTsFromLatest(c, 14875690, 16309169)
	select {}
}

// func TestSync(t *testing.T) {
// 	t.Log("Test Sync")
// 	models.Setup()
// 	config.SetUp()
// 	UD := contracts.UD
// 	chain := "POLYGON"
// 	// node := nodes.New(UD, chain)
// 	var limit int64 = 100
// 	all := models.CountMintEvents(UD, chain)
// 	var page int64 = all / limit
// 	var p int64 = page / 2
// 	for ; p <= page; p++ {
// 		log.Println(limit, p, page, all)
// 		es := models.FindMintEvents("Unstoppable Domains", chain, limit, p)
// 		SyncEvent2URI(UD, chain, es)
// 	}
// }

// func TestSyncEvent2Event(t *testing.T) {
// 	os.Setenv("NETWORK", "main")
// 	os.Setenv("BLOCKCHAIN", "POLYGON")
// 	models.Setup()
// 	config.SetUp()
// 	UD := "Unstoppable Domains"
// 	SyncEvent2Event(UD)
// }

func TestFetchNewBlock(t *testing.T) {
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "Conflux")
	did := "CNS"
	models.Setup(did)
	config.SetUp()
	FetchNewBlock(did)
	select {}
}

func TestFetchKeyInfo(t *testing.T) {
	t.Log("TestFetchBitNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "ETHEREUM")
	models.Setup("DOTBIT")
	c, _ := models.FindContractByAddress("0x60eb332bd4a0e2a9eeb3212cfdd6ef03ce4cb3b5")
	node := nodes.New(c.DID(), c.Blockchain())
	tokenId, _ := contracts.Str2BigInt("636089999433181234785303117851547307074501179544")
	key := node.FetchKeyInfo(c.Address, tokenId)
	expire := key.Expire()
	log.Println("expire", expire)
	assert.True(t, expire.Uint64() > 0)
}

func TestFetchNameInfo(t *testing.T) {
	t.Log("TestFetchBitNFTs")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "ETHEREUM")
	models.Setup("DOTBIT")
	contract := "0x60eb332bd4a0e2a9eeb3212cfdd6ef03ce4cb3b5"
	c, _ := models.FindContractByAddress(contract)
	node := nodes.New(c.DID(), c.Blockchain())
	nfts := models.FindBitDomains()
	log.Println("Total=>", len(nfts))
	nftChunks := util.ChunkSlice(nfts, 100)
	var wg sync.WaitGroup
	for i := range nftChunks {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			uris := cmap.New[models.NFT]()
			var wg sync.WaitGroup
			for j, nft := range nftChunks[i] {
				wg.Add(1)
				go func(j int, nft models.NFT) {
					defer wg.Done()
					tokenId, _ := contracts.Str2BigInt(nft.TokenIdDec)
					meta, err := node.NFTmeta("ETHEREUM", contract, tokenId)
					if err != nil {
						log.Println(j, err, tokenId)
						return
					}
					nft.Name = meta.Name()
					uris.Set(nft.Id, nft)
					log.Println(i, j, nft.Name)
				}(j, nft)
			}
			wg.Wait()
			models.UpdateNFTByMap(uris.Items())
		}(i)
	}
	wg.Wait()
}

func TestMoveOldToNew(t *testing.T) {
	t.Log("TestMoveOldToNew")
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", "POLYGON")
	did := "UD"
	config.SetUp()
	models.Setup(did)

	MoveOldToNew(did)
}

func TestFetchNFTsForCollection(t *testing.T) {
	os.Setenv("BLOCKCHAIN", "POLYGON")
	did := "UD"
	models.Setup(did)
	config.SetUp()
	FetchNFTsForCollection(did)
}
