package nodes

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"fetcher/config"
	"fetcher/contracts"
	"fetcher/models"

	"github.com/stretchr/testify/assert"
)

var (
	UD_MAIN = "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f" // main
	UD      = "0x2a93C52E7B6E7054870758e15A1446E769EdfB93" // test
)

func TestDomain(t *testing.T) {
	t.Log("Test Domain")
	domain, err := Domain("mdripmom.blockchain")
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, len(domain.Records) > 0)
}

func TestRegExp(t *testing.T) {
	source := "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response. Based on your parameters and the response size limit, this block range should work: [0x1183f37, 0x16c7daf]"
	reg := regexp.MustCompile(`\[(?P<From>.+?), (?P<To>.+?)\]`)
	if reg != nil {
		results := reg.FindStringSubmatch(source)
		log.Println(results)
		fromHex := strings.TrimPrefix(results[1], "0x")
		toHex := strings.TrimPrefix(results[2], "0x")
		log.Println(fromHex, toHex)
		from, _ := strconv.ParseUint(fromHex, 16, 64)
		to, _ := strconv.ParseUint(toHex, 16, 64)
		log.Println(from, to)
	}
}

// alex011221a2.coin
func TestUDGetNftMetadata(t *testing.T) {
	t.Log("Test GetNftMetadata")
	// tokenId, err := FindNFTTokenId("alex011221a2.coin")
	// if err != nil {
	// 	t.Log(err.Error())
	// 	return
	// }
	tokenId := "31657805781633231781251673101568514227656771479804165504185953547215780006445"
	// Initialize Alchemy node
	// alchemy := New(contracts.UD, "POLYGON")
	// tid, _ := contracts.Str2BigInt(tokenId)
	// meta, err := alchemy.Metadata(UD_MAIN, tid)
	meta, err := UDMeta(tokenId)
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, meta.TokenId() == tokenId)
}

func TestGetContractMetadata(t *testing.T) {
	alchemy := New(contracts.UD, "POLYGON")
	meta, err := alchemy.ContractMetadata("0x7b29dac7b5d3a90696c8b059cbf1fae9af32fa28")
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, meta.DeployedBlockNumber() > 0)
}

func TestGetNFTsForCollection(t *testing.T) {
	project := "Unstoppable Domains"
	chain := "POLYGON"
	models.Setup("UD")
	config.SetUp()
	node := NewAlchemy(contracts.UD, chain)
	chainId := 137
	contract := "0xa9a6a3626993d487d2dbda3173cf58ca1a9d9e9f"
	var limit uint64 = 100
	collection, err := node.GetNFTsForCollection(contract, "", limit)
	for err == nil {
		nfts := collection.Nfts
		next := collection.NextToken
		uris := make([]models.NFT, 0)
		metas := make(map[string]Metadata, 0)
		for _, nft := range nfts {
			tokenId, ok := contracts.Hex2BigInt(nft.Id.TokenId)
			if !ok {
				log.Println("can't Hex2BigInt", nft.Id.TokenId)
				continue
			}
			ID := models.NFTID(chain, contract, tokenId.String())
			meta, err := node.Metadata(contract, tokenId)
			var uri string
			if err == nil {
				if count, err := countNFTmeta(ID); err == nil && count == 0 {
					metas[meta.Id()] = meta
				}
				uri = meta.Name()
			}
			nft := models.NewNFT(project, uri, contract, tokenId, chain, chainId, contract, 0)
			keyInfo := node.KeyInfo(contract, tokenId)
			nft.SetExpires(keyInfo.Expire().Uint64())
			nft.SetOwner(keyInfo.Owner())
			uris = append(uris, nft)
		}
		if len(uris) > 0 {
			models.SaveNFTs(uris)
		}
		if len(metas) > 0 {
			SyncNFTmeta(metas)
		}
		nextCollection, err := node.GetNFTsForCollection(contract, next, limit)
		if err != nil {
			log.Println("No more data", err.Error())
		}
		collection = nextCollection
	}
}

func TestOwnerOf(t *testing.T) {
	t.Log("Test OwnerOf")
	nft, err := contracts.New(contracts.UD, "0xa9a6a3626993d487d2dbda3173cf58ca1a9d9e9f", NewAlchemy(UD, "POLYGON").EthClient())
	if err != nil {
		t.Log(err.Error())
		return
	}
	tokenId, ok := contracts.Str2BigInt("40103477672410470392349689981673008939926702360484281818524236634979959681317")
	assert.True(t, ok)
	owner, _ := nft.Owner(tokenId)
	log.Println(owner.Hex())
	assert.True(t, owner.Hex() != contracts.ZERO)
}
