package nodes

import (
	"log"
	"runtime/debug"
	"time"

	"fetcher/contracts"
	"fetcher/models"

	"github.com/ethereum/go-ethereum/core/types"
)

func parseEvents(node *Node, c models.Contract, logCh chan []types.Log) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic error:", err)
			log.Panicln("stacktrace from panic:", string(debug.Stack()))
		}
	}()

	token, err := contracts.New(c.DID(), c.Contract(), node.rawclient)
	if err != nil {
		log.Panicln("parseEvents:", err.Error(), c.DID())
	}
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop() // release resources
	data := make([]types.Log, 0)
	for {
		select {
		case <-ticker.C:
			length := len(data)
			if length > 0 {
				to := data[len(data)-1].BlockNumber
				// log.Printf("parse logs ticker %ds size:%d", waitTime, length)
				uriCh, tfCh := batchParseLog(token, c.Contract(), data)
				go bufferSyncMint(node, c, uriCh, to)
				go bufferSyncTransfer(node, c, tfCh, to)
				data = make([]types.Log, 0)
			}
		case logs := <-logCh:
			// log.Printf("log size: %d", len(logs))
			data = append(data, logs...)
		}
	}
}

func batchParseLog(token contracts.NFT, contract string, logs []types.Log) (chan NewURIEvent, chan TransferEvent) {
	cap := len(logs) / 2
	uriCh := make(chan NewURIEvent, cap)
	tfCh := make(chan TransferEvent, cap)
	go func() {
		for _, vLog := range logs {
			// log.Println(vLog.TxHash.Hex())
			newURIEvent, err := token.ParseNewURI(vLog)
			if err == nil && newURIEvent != nil {
				uriCh <- NewURIEvent{
					BlockNumber:      vLog.BlockNumber,
					TransactionHash:  vLog.TxHash.Hex(),
					TransactionIndex: vLog.TxIndex,
					BlockHash:        vLog.BlockHash.Hex(),
					LogIndex:         vLog.Index,
					Contract:         contract,
					Uri:              newURIEvent.Uri,
					TokenId:          newURIEvent.TokenId,
					From:             newURIEvent.From.Hex(),
					To:               newURIEvent.To.Hex(),
				}
			}
		}
		close(uriCh)
	}()
	go func() {
		for _, vLog := range logs {
			newTransferEvent, err := token.ParseTransfer(vLog)
			if err == nil {
				e := TransferEvent{
					BlockNumber:      vLog.BlockNumber,
					TransactionHash:  vLog.TxHash.Hex(),
					TransactionIndex: vLog.TxIndex,
					BlockHash:        vLog.BlockHash.Hex(),
					LogIndex:         vLog.Index,
					Contract:         contract,
					From:             newTransferEvent.From.Hex(),
					To:               newTransferEvent.To.Hex(),
					TokenId:          newTransferEvent.TokenId,
					Event:            Transfer,
				}

				if e.From == contracts.ZERO {
					e.Event = Mint
				} else if e.To == contracts.ZERO {
					e.Event = Burn
				} else {
					e.Event = Transfer
				}
				tfCh <- e
			}
		}
		close(tfCh)
	}()

	return uriCh, tfCh
}
