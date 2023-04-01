package nodes

import (
	"log"
	"strconv"
	"sync"

	"fetcher/models"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func bufferSyncMint(node *Node, c models.Contract, uriCh chan NewURIEvent, to uint64) {
	// log.Printf("bufferSyncMint from:%d to:%d", c.DeployedBlockNumber, to)
	// var to uint64 = c.LatestBlockNumber
	uris := cmap.New[models.NFT]()
	metas := cmap.New[Metadata]()
	var wg sync.WaitGroup
	for e := range uriCh {
		wg.Add(1)
		go func(e NewURIEvent) {
			defer wg.Done()
			ID := models.NFTID(c.Chain, c.Address, e.TokenId.String())
			_, ok := uris.Get(ID)
			if ok {
				return
			}
			uri := e.Uri
			if len(uri) == 0 {
				meta, err := node.NFTmeta(node.Chain(), c.Address, e.TokenId)
				if err != nil {
					log.Println("not found metadata of nft ", err, e.TokenId)
				} else {
					metas.Set(meta.Id(), meta)
					uri = meta.Name()
				}
			}
			nft := models.NewNFT(c.Project, uri, c.Address, e.TokenId, c.Chain, c.ChainId, e.To, e.BlockNumber)
			keyInfo := node.FetchKeyInfo(e.Contract, e.TokenId)
			nft.SetExpires(keyInfo.Expire().Uint64())
			// nft.SetOwner(keyInfo.Owner())
			uris.Set(ID, nft)
		}(e)
	}
	wg.Wait()
	if uris.Count() > 0 {
		models.UpdateNFTByMap(uris.Items())
	}
	if metas.Count() > 0 {
		SyncNFTmeta(metas.Items())
	}
	// models.UpdateLastestBlock(c, to)
	// log.Printf("bufferSyncMint len:%d blocks:%d from:%d ", len(uris), to-c.LatestBlockNumber, c.LatestBlockNumber)
}

func bufferSyncTransfer(node *Node, c models.Contract, tfCh chan TransferEvent, to uint64) {
	// log.Printf("bufferSyncTransfer from:%d to:%d", c.DeployedBlockNumber, to)
	// var to uint64 = c.LatestBlockNumber
	// nodeEvents := make([]TransferEvent, 0)
	events := cmap.New[models.Event]()
	uris := cmap.New[models.NFT]()
	var wg sync.WaitGroup
	for event := range tfCh {
		wg.Add(1)
		go func(event TransferEvent) {
			defer wg.Done()
			eID := event.TransactionHash + "_" + strconv.FormatUint(uint64(event.LogIndex), 10)
			ne := models.NewEvent(c.Project, event.From, event.To, event.BlockNumber, event.TokenId, c.Chain, c.ChainId, c.Address, c.Tld, event.Event, event.TransactionHash, event.LogIndex)
			_, ok := events.Get(eID)
			if ok {
				return
			}
			events.Set(eID, ne)
			// nodeEvents = append(nodeEvents, event)
			ID := models.NFTID(c.Chain, c.Address, event.TokenId.String())
			if nft, err := models.FindNFTById(ID); err == nil {
				_, ok := uris.Get(ID)
				if ok {
					return
				}
				keyInfo := node.FetchKeyInfo(event.Contract, event.TokenId)
				nft.SetExpires(keyInfo.Expire().Uint64())
				nft.SetOwner(keyInfo.Owner())
				uris.Set(ID, nft)
			}
		}(event)
	}
	wg.Wait()
	if uris.Count() > 0 {
		models.UpdateNFTByMap(uris.Items())
	}
	if events.Count() > 0 {
		models.SaveEvents(events.Items())
	}
	// if len(nodeEvents) > 0 {
	// 	SaveEvent(nodeEvents)
	// }
	models.UpdateLastestBlock(c, to)
	// log.Printf("bufferSyncTransfer len:%d blocks:%d from:%d ", len(events), to-c.LatestBlockNumber, c.LatestBlockNumber)
}
