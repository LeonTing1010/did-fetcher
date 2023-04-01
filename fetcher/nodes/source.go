package nodes

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"fetcher/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	ZERO    = "0x0000000000000000000000000000000000000000"
	TIMEOUT = 100 * time.Second
)

var (
	_            *mongo.Collection = config.GetCollection("Sync_Contract_Metadata_Collection")
	nfts         *mongo.Collection = config.GetCollection("Sync_NFT_Collection")
	nftMetadatas *mongo.Collection = config.GetCollection("Sync_NFT_Metadata_Collection")
	nftEvents    *mongo.Collection = config.GetCollection("Sync_NFT_Event_Collection")
	accountInfo  *mongo.Collection = config.GetCollection("Sync_Bit_Account_Info_Collection")
)

func SyncNFTs(metas []interface{}, events []NewURIEvent) {
	if len(metas) == 0 {
		return
	}
	uris := make([]interface{}, 0)
	for _, event := range events {
		uris = append(uris, NFT{
			BlockNumber:      event.BlockNumber,
			TransactionHash:  event.TransactionHash,
			TransactionIndex: event.TransactionIndex,
			BlockHash:        event.BlockHash,
			LogIndex:         event.LogIndex,
			Contract:         event.Contract,
			Uri:              event.Uri,
			TokenId:          event.TokenId.String(),
		})
		// if !hasNFTRecords(event.Uri) {
		// 	insertUniqueInfo(project, event.Uri)
		// }
	}
	opts := options.InsertMany().SetOrdered(false)
	ctx := context.Background()
	nfts.InsertMany(ctx, uris, opts)
	nftMetadatas.InsertMany(ctx, metas, opts)
}

func countNFTmeta(id string) (int64, error) {
	return nftMetadatas.CountDocuments(context.Background(), bson.M{"_id": id})
}

func SaveNFTmeta(meta Metadata) {
	nftMetadatas.InsertOne(context.Background(), meta)
}

func SyncNFTmeta(metas map[string]Metadata) {
	l := len(metas)
	if l == 0 {
		return
	}
	ms := make([]interface{}, 0)
	for _, m := range metas {
		if c, err := countNFTmeta(m.Id()); err == nil && c == 0 {
			ms = append(ms, m)
		}
	}
	if len(ms) > 0 {
		opts := options.InsertMany().SetOrdered(false)
		ctx := context.Background()
		result, err := nftMetadatas.InsertMany(ctx, ms, opts)
		if err != nil {
			log.Println("SyncNFTmeta", err)
		}
		log.Println("SyncNFTmeta", len(result.InsertedIDs))
	}
}

func SyncNFT(project string, meta Metadata, event *NewURIEvent) {
	if len(event.Uri) == 0 {
		event.Uri = meta.Name()
	}
	// log.Printf("SyncNFT DID:%s URI:%s TLD:%s tx:%s", project, event.Uri, meta.TLD(), event.TransactionHash)
	insertNFT(event.Contract, event.Uri, NFT{
		BlockNumber:      event.BlockNumber,
		TransactionHash:  event.TransactionHash,
		TransactionIndex: event.TransactionIndex,
		BlockHash:        event.BlockHash,
		LogIndex:         event.LogIndex,
		Contract:         event.Contract,
		Uri:              event.Uri,
		TokenId:          event.TokenId.String(),
	})
	insertNFTMetadata(meta.ContractAddress(), meta.TokenId(), meta)
	// if !hasNFTRecords(event.Uri) {
	// 	insertUniqueInfo(project, event.Uri)
	// }
}

func insertUniqueInfo(project, uri string) {
	if project == "Unstoppable Domains" {

		domain, err := Domain(uri)
		if err != nil {
			log.Println(err.Error())
		}
		insertNFTRecords(domain.Meta.Domain, domain)

	} else if project == "dotbit" {
		account, err := Account(uri)
		if err != nil {
			log.Println(err.Error())
		}
		records, err := Records(uri)
		if err != nil {
			log.Println(err.Error())
		}
		if len(records.Records) > 0 {
			account.Records = records.Records
		}
		insertAccountInfo(account.Account, account)
	}
}

func hasNFTRecords(name string) bool {
	ctx := context.Background()

	filter := bson.M{"meta.domain": name}
	if count, err := nfts.CountDocuments(ctx, filter); err == nil && count > 0 {
		return true
	}
	return false
}

func insertNFTRecords(name string, domain any) (string, error) {
	ctx := context.Background()

	filter := bson.M{"meta.domain": name}
	if count, err := nfts.CountDocuments(ctx, filter); err == nil && count > 0 {
		nfts.DeleteMany(ctx, filter)
	}
	result, err := nfts.InsertOne(ctx, domain)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("sync: InsertNFTRecords Not objectid.ObjectID")
}

// func FindContractMetadata(address string) (*ALCMeta, error) {
// 	ctx := context.Background()

// 	filter := bson.M{"address": strings.ToLower(address)}
// 	var contractMetadata ALCMeta
// 	err := cMetadatas.FindOne(ctx, filter).Decode(&contractMetadata)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 		return nil, err
// 	}
// 	return &contractMetadata, nil
// }

// func SaveContractMetadata(address string, contract any) (string, error) {
// 	ctx := context.Background()

// 	filter := bson.M{"address": address}
// 	if count, err := cMetadatas.CountDocuments(ctx, filter); err == nil && count > 0 {
// 		cMetadatas.DeleteMany(ctx, filter)
// 	}
// 	result, err := cMetadatas.InsertOne(ctx, contract)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 		return "", err
// 	}
// 	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
// 		return oid.Hex(), nil
// 	}
// 	return "", errors.New("sync: InsertContractMetadata Not objectid.ObjectID")
// }

type NFT struct {
	BlockNumber       uint64
	TransactionHash   string
	TransactionIndex  uint
	BlockHash         string
	LogIndex          uint
	Contract          string
	Uri               string
	TokenId           string
	LatestBlockNumber uint64
}

func (nft NFT) TLD() string {
	ds := strings.Split(nft.Uri, ".")
	return ds[len(ds)-1]
}

// func updateOneNFT(contract string, tokenId string, latest uint64) (string, error) {
// 	ctx := context.Background()

// 	filter := bson.M{"tokenId": tokenId, "contract": contract}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"LatestBlockNumber": latest,
// 		},
// 	}
// 	_, err := nfts.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return tokenId, err
// 	}
// 	// log.Printf("update nft contract: %s latest: %d", contract, latest)
// 	return tokenId, nil
// }

func insertNFT(contract, uri string, nft any) (string, error) {
	ctx := context.Background()

	filter := bson.M{"uri": uri, "contract": contract}
	if count, err := nfts.CountDocuments(ctx, filter); err == nil && count > 0 {
		nfts.DeleteMany(ctx, filter)
	}
	result, err := nfts.InsertOne(ctx, nft)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("sync: InsertNFT Not objectid.ObjectID")
}

// func FindNFTs(contract string) []NFT {
// 	ctx := context.Background()

// 	cursor, err := nfts.Find(ctx, bson.M{"contract": contract})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var nfts []NFT
// 	if err = cursor.All(ctx, &nfts); err != nil {
// 		log.Fatal(err)
// 	}
// 	return nfts
// }

func FindNFTTokenId(uri string) (string, error) {
	ctx := context.Background()

	filter := bson.M{"uri": uri}
	var nft TransferEvent
	err := nfts.FindOne(ctx, filter).Decode(&nft)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return nft.TokenId.String(), nil
}

func FindNFTMetaById(id string) (Metadata, error) {
	ctx := context.Background()
	filter := bson.M{"_id": id}
	var meta NFTmeta
	if err := nftMetadatas.FindOne(ctx, filter).Decode(&meta); err != nil {
		return nil, err
	}
	return meta, nil
}

func insertNFTMetadata(contract, tokenId string, nftMeta any) (string, error) {
	ctx := context.Background()

	filter := bson.M{"id.tokenid": tokenId, "contract.address": contract}
	if count, err := nftMetadatas.CountDocuments(ctx, filter); err == nil && count > 0 {
		nftMetadatas.DeleteMany(ctx, filter)
	}
	result, err := nftMetadatas.InsertOne(ctx, nftMeta)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("sync: InsertNFTMetadata Not objectid.ObjectID")
}

/*
	{
	  "_id": {
	    "$oid": "63fdaddc6879b3b23c305653"
	  },
	  "blocknumber": 14884034,
	  "transactionhash": "0x5e47536168842bbf492b39173aca4e2b678bf7f73e66cfcbeb33b264596371da",
	  "transactionindex": 3,
	  "blockhash": "0x08d67ae3abfabeeeccad0a00920c96c6d82cf16daaf6cc5dcc9d3eb4b8bac561",
	  "logindex": 3,
	  "contract": "0x60eb332bd4a0e2a9eeb3212cfdd6ef03ce4cb3b5",
	  "from": "0x0000000000000000000000000000000000000000",
	  "to": "0x1D643FAc9a463c9d544506006a6348c234dA485f",
	  "tokenid": "574639461609516808003167560251625784695169817307",
	  "event": "Mint"
	}
*/
// func FindNFTEvent(contract, tokenid string) TransferEvent {
// 	var tf TransferEvent
// 	nftEvents.FindOne(context.Background(), bson.M{"contract": contract, "tokenid": tokenid}).Decode(&tf)
// 	return tf
// }

// func SyncEvent2Event(project, chain string, limit, page int64) {
// 	es := models.FindAllEvents(project, chain, limit, page)
// 	for _, e := range es {
// 		tf := FindNFTEvent(e.ContractAddress, e.TokenIdDec)
// 		e.TransactionHash = tf.TransactionHash
// 		e.LogIndex = tf.LogIndex
// 		models.UpdateEvent(e)
// 		log.Println("Update TokenId=>", e.TokenIdDec)
// 	}
// }

func SaveNFTEvents(event *TransferEvent) {
	// Set the event type
	if event.From == ZERO {
		event.Event = "Mint"
	} else if event.To == ZERO {
		event.Event = "Burn"
	} else {
		event.Event = "Transfer"
	}
	// log.Printf("SyncNFTEvents: %s  block: %d contract: %s txhash:%s", event.Event, event.BlockNumber, event.Contract, event.TransactionHash)
	insertNFTEvent(event.TransactionHash, event)
	// updateOneNFT(event.Contract, event.TokenId.String(), event.BlockNumber)
}

func SaveEvent(events []TransferEvent) error {
	ctx := context.Background()
	tfs := make([]interface{}, 0)
	for _, event := range events {
		filter := bson.M{"transactionhash": event.TransactionHash, "logindex": event.LogIndex}
		if count, err := nftEvents.CountDocuments(ctx, filter); err != nil || count == 0 {
			tfs = append(tfs, event)
		}
	}
	if len(tfs) > 0 {
		opts := options.InsertMany().SetOrdered(false)
		_, err := nftEvents.InsertMany(ctx, tfs, opts)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
	}
	return nil
}

// func CountEvents(transactionHash string, logIndex uint) (int64, error) {
// 	return nftEvents.CountDocuments(context.Background(), bson.M{"transactionHash": transactionHash, "logIndex": logIndex})
// }

// func SaveNFT(nfts []any) error {
// 	ctx := context.Background()
// 	// for _, event := range events {
// 	// 	filter := bson.M{"hash": event.TransactionHash}
// 	// 	if count, err := nftEvents.CountDocuments(ctx, filter); err == nil && count > 0 {
// 	// 		nftEvents.DeleteMany(ctx, filter)
// 	// 	}
// 	// }
// 	_, err := nftEvents.InsertMany(ctx, nfts)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 		return err
// 	}
// 	return nil
// }

func insertNFTEvent(hash string, event any) (string, error) {
	ctx := context.Background()

	filter := bson.M{"hash": hash}
	if count, err := nftEvents.CountDocuments(ctx, filter); err == nil && count > 0 {
		nftEvents.DeleteMany(ctx, filter)
	}
	result, err := nftEvents.InsertOne(ctx, event)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("sync: InsertNFTEvent Not objectid.ObjectID")
}

func insertAccountInfo(account string, event any) (string, error) {
	ctx := context.Background()

	filter := bson.M{"account": account}
	if count, err := accountInfo.CountDocuments(ctx, filter); err == nil && count > 0 {
		accountInfo.DeleteMany(ctx, filter)
	}
	result, err := accountInfo.InsertOne(ctx, event)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("sync: InsertAccountInfo Not objectid.ObjectID")
}
