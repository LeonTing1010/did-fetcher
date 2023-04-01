package sync

import (
	"log"
	"strings"
	"sync"

	"fetcher/config"
	"fetcher/contracts"
	"fetcher/models"
	"fetcher/nodes"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func FetchNFTsForCollection(did string) {
	project := did
	chain := config.Blockchain()
	log.Printf("FetchNFTsForCollection => %s %s", project, chain)
	// Finds all contracts related to the project
	cts, err := models.FindContractsBySymbol(project)
	if err != nil {
		log.Fatalln(err.Error())
	}
	l := len(cts)
	if l == 0 {
		log.Fatalln("Can't find " + project + " in Collection")
	}
	// Loops through all the contracts related to the project
	for _, c := range cts {
		if c.Chain != chain {
			continue
		}
		start := config.Env("START_TOKEN")
		log.Println("START_TOKEN", start)
		contract := c.Contract()
		node := nodes.New(c.DID(), c.Blockchain())
		var limit uint64 = 100
		collection, err := node.GetNFTsForCollection(contract, start, limit)
		for err == nil {
			nfts := collection.Nfts
			next := collection.NextToken
			uris := cmap.New[models.NFT]()
			// uris := make(map[string]models.NFT, 0)
			// metas := make([]nodes.Metadata, 0)
			var wg sync.WaitGroup
			for _, token := range nfts {
				wg.Add(1)
				go func(token nodes.Token) {
					defer wg.Done()
					tokenId, ok := contracts.Hex2BigInt(token.Id.TokenId)
					if !ok {
						log.Println("can't Hex2BigInt", token.Id.TokenId)
						return
					}
					// ID := models.NFTID(chain, contract, tokenId.String())
					var uri string
					// meta, err := node.NFTmeta(chain, contract, tokenId)
					// if err == nil {
					// 	uri = meta.Name()
					// }
					nft := models.NewNFT(c.Project, uri, contract, tokenId, chain, c.ChainId, contract, c.DeployedBlockNumber)
					// keyInfo := node.FetchKeyInfo(contract, tokenId)
					// nft.SetExpires(keyInfo.Expire().Uint64())
					// nft.SetOwner(keyInfo.Owner())
					uris.Set(nft.Id, nft)
					log.Println(nft.TokenId, nft.Name)
				}(token)
			}
			wg.Wait()
			ul := uris.Count()
			if ul > 0 {
				models.UpdateNFTByMap(uris.Items())
			}
			log.Printf("SaveNFTs nfts: %d next: %s", ul, next)
			nextCollection, err := node.GetNFTsForCollection(contract, next, limit)
			if err != nil {
				log.Println("No more data", err.Error())
			}
			collection = nextCollection
		}

	}
}

func SyncEvent2URI(did, chain string, es []models.Event) {
	// log.Printf("len:%d ", len(es))
	node := nodes.New(did, chain)
	uris := make([]models.NFT, 0, len(es))
	for _, e := range es {
		count, _ := models.CountNFTs(models.NFTID(e.Chain, e.ContractAddress, e.TokenIdDec))
		if count > 0 {
			// log.Println("NFT Count:", count)
			continue
		}
		// c, _ := models.FindContractByAddress(e.ContractAddress)
		tid, _ := contracts.Str2BigInt(e.TokenIdDec)
		meta, _ := node.NFTmeta(chain, e.ContractAddress, tid)
		if meta == nil {
			continue
		}
		nft := models.NewNFT(e.Project, meta.Name(), e.ContractAddress, tid, e.Chain, e.ChainId, e.To, e.BlockNumber)
		uris = append(uris, nft)
	}
	if len(uris) > 0 {
		models.SaveNFTs(uris)
	}
	log.Printf("SyncURIs len:%d ", len(uris))
}

// func SyncEvent2Event(project string) {
// 	chain := config.Blockchain()
// 	log.Printf("SyncEvent2Event: %s %s", project, chain)
// 	var limit int64 = 10
// 	all := models.CountAllEvents(project, chain)
// 	var page int64 = all / limit
// 	var p int64 = 0
// 	for ; p <= page; p++ {
// 		log.Println(limit, p, page, p*limit, all)
// 		nodes.SyncEvent2Event(project, chain, limit, p)
// 	}
// }

// func SyncEvent2NFT(did string) {
// 	chain := config.Blockchain()
// 	log.Printf("SyncEvent2NFT: %s %s", did, chain)
// 	var limit int64 = 10000
// 	all := models.CountMintEvents(did, chain)
// 	// var all int64 = 2545345
// 	var page int64 = all / limit
// 	var p int64 = 0

// 	node := nodes.New(contracts.UD, chain)
// 	for ; p <= page; p++ {
// 		log.Println(limit, p, page, p*limit, all)
// 		es := models.FindMintEvents(project, chain, limit, p)
// 		node.SyncEvent2URI(es)
// 	}
// }

func MoveOldToNew(did string) {
	chain := config.Blockchain()
	var total int64 = 854869
	var limit int64 = 100
	var page int64 = total / limit
	log.Println("Total,Page", total, page)
	var p int64 = 0
	cs, _ := models.FindContractsBySymbol(did)
	for _, c := range cs {
		if c.Chain != chain {
			continue
		}
		node := nodes.New(did, chain)
		for ; p <= page; p++ {
			es := models.FindUDDomains(chain, limit, p)
			l := len(es)
			log.Println("p", p, l, page-p)
			if l == 0 {
				break
			}
			var wg sync.WaitGroup
			mn := cmap.New[models.NFT]()
			for _, nft := range es {
				wg.Add(1)
				go func(nft models.NFT) {
					defer wg.Done()
					if len(nft.Name) == 0 {
						tokenId, ok := contracts.Hex2BigInt(nft.TokenId)
						if !ok {
							return
						}
						meta, _ := node.NFTmeta(chain, c.Address, tokenId)
						nft.Name = meta.Name()
					} else {
						if !strings.Contains(nft.Name, ".") {
							nft.Name = nft.Name + "." + nft.Parent
						}
					}
					mn.Set(nft.Id, nft)
				}(nft)
			}
			wg.Wait()
			if mn.Count() > 0 {
				models.UpdateNFTByMap(mn.Items())
			}
		}
	}
}
