package models

import (
	"context"
	"errors"
	"log"
	"time"

	"fetcher/config"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const TIMEOUT = 100 * time.Second

var (
	_         *mongo.Collection = config.GetCollection("Project_Collection")
	contracts *mongo.Collection = config.GetCollection("Collection_Collection")
	events    *mongo.Collection
	nfts      *mongo.Collection
)

func Setup(did string) {
	var domain, activity string
	if config.Network() == "main" {
		domain = did + "_Domain_Collection"
		activity = did + "_Activity_Collection"
	} else {
		domain = "Test_Domain_Collection"
		activity = "Test_Activity_Collection"
	}
	log.Println("Collection Setup:", config.Network(), domain, activity)
	nfts = config.GetCollection(domain)
	events = config.GetCollection(activity)
}

func SaveOneEvent(event Event) (string, error) {
	ctx := context.Background()

	count, err := events.CountDocuments(ctx, bson.M{"_id": event.Id})
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	if count == 0 {
		result, err := events.InsertOne(ctx, event)
		if err != nil {
			log.Fatalln(err.Error())
			return "", err
		}
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex(), nil
		}
	}
	return event.Id.Hex(), nil
}

func UpdateEvent(event Event) (string, error) {
	ctx := context.Background()

	update := bson.M{
		"$set": bson.M{
			"transactionHash": event.TransactionHash,
			"logIndex":        event.LogIndex,
		},
	}
	_, err := events.UpdateByID(ctx, event.Id, update)
	if err != nil {
		log.Println(err.Error())
		return event.Id.Hex(), err
	}
	return event.Id.Hex(), nil
}

func FindContractByAddress(address string) (Contract, error) {
	ctx := context.Background()
	// Clean up when this function returns
	// Create the filter object to search by project
	filter := bson.M{"address": address}
	var contract Contract
	// Search the collection for the given project
	if err := contracts.FindOne(ctx, filter).Decode(&contract); err != nil {
		return Contract{}, err
	}
	return contract, nil
}

// func FindContractsByName(project string) ([]Contract, error) {
// 	ctx := context.Background()
// 	// Clean up when this function returns
// 	// Create the filter object to search by project
// 	filter := bson.M{"project": project}
// 	// Search the collection for the given project
// 	cursor, err := contracts.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Create an empty results array
// 	var contracts []Contract
// 	// Populate the results array with the data from the cursor
// 	if err = cursor.All(ctx, &contracts); err != nil {
// 		return nil, err
// 	}
// 	var fcs []Contract
// 	for _, c := range contracts {
// 		if config.Network() == c.Network() {
// 			fcs = append(fcs, c)
// 		}
// 	}
// 	return fcs, nil
// }

func FindContractsBySymbol(symbol string) ([]Contract, error) {
	ctx := context.Background()
	// Clean up when this function returns
	// Create the filter object to search by project
	filter := bson.M{"symbol": symbol}
	// Search the collection for the given project
	cursor, err := contracts.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Create an empty results array
	var contracts []Contract
	// Populate the results array with the data from the cursor
	if err = cursor.All(ctx, &contracts); err != nil {
		return nil, err
	}
	var fcs []Contract
	fcs = append(fcs, contracts...)
	return fcs, nil
}

func SaveOneContract(contract Contract) (string, error) {
	ctx := context.Background()

	result, err := contracts.InsertOne(ctx, contract)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("Contract: Not objectid.ObjectID")
}

func UpdateLastestBlock(c Contract, to uint64) {
	if c.LatestBlockNumber >= to {
		return
	}
	ctx := context.Background()
	_, err := contracts.UpdateByID(ctx, c.Id, bson.M{
		"$set": bson.M{
			"latestBlockNumber": to,
		},
	})
	if err != nil {
		log.Println("UpdateLastestBlock", err, to)
	}
}

func UpdateOneContract(contract Contract) (string, error) {
	// log.Println("update contract =>", contract.Name, contract.LatestBlockNumber)
	ctx := context.Background()
	update := bson.M{
		"$set": bson.M{
			"tokenType":           contract.TokenType,
			"symbol":              contract.Symbol,
			"deployedBlockNumber": contract.DeployedBlockNumber,
			"name":                contract.Name,
			"latestBlockNumber":   contract.LatestBlockNumber,
		},
	}
	_, err := contracts.UpdateByID(ctx, contract.Id, update)
	if err != nil {
		log.Println(err.Error())
		return contract.Id, err
	}
	return contract.Id, nil
}

func DropAllNfts() error {
	ctx := context.Background()

	if err := nfts.Drop(ctx); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

// func FindAllNFTs() []NFT {
// 	ctx := context.Background()

// 	cursor, err := nfts.Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var nfts []NFT
// 	if err = cursor.All(ctx, &nfts); err != nil {
// 		log.Fatal(err)
// 	}
// 	return nfts
// }

// func FindNFTs(contract string) []NFT {
// 	ctx := context.Background()

// 	cursor, err := nfts.Find(ctx, bson.M{"contractAddress": contract})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var nfts []NFT
// 	if err = cursor.All(ctx, &nfts); err != nil {
// 		log.Fatal(err)
// 	}
// 	return nfts
// }

func FindNFTById(id string) (NFT, error) {
	ctx := context.Background()
	filter := bson.M{"_id": id}
	var nft NFT
	err := nfts.FindOne(ctx, filter).Decode(&nft)
	if err != nil {
		// log.Println("not found nft id ", id, err.Error())
		return NFT{}, err
	}
	return nft, nil
}

// func FindNFTByTokenId(contract string, tokenId *big.Int) (*NFT, error) {
// 	ctx := context.Background()

// 	tid := fmt.Sprintf("0x%x", tokenId)
// 	filter := bson.M{"contractAddress": contract, "tokenId": tid}
// 	var nft NFT
// 	err := nfts.FindOne(ctx, filter).Decode(&nft)
// 	if err != nil {
// 		log.Printf("%s,collection:%s,contract:%s,tokenId:%s", err.Error(), nfts.Name(), contract, tid)
// 		return nil, err
// 	}
// 	return &nft, nil
// }

func SaveOneNFT(nft NFT) (string, error) {
	ctx := context.Background()

	filter := bson.M{"_id": nft.Id}
	count, err := nfts.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	if count > 0 {
		nfts.DeleteOne(ctx, filter)
	}
	if _, err := nfts.InsertOne(ctx, nft); err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	return nft.Id, nil
}

func UpdateOneNFT(nft *NFT) (string, error) {
	// log.Printf("update nft: %s latest: %d", nft.Id, nft.LatestBlockNumber)
	ctx := context.Background()

	update := bson.M{
		"$set": bson.M{
			"latestBlockNumber": nft.LatestBlockNumber,
			// "owner":             nft.Owner,
			// "name":              nft.Name,
		},
	}
	_, err := nfts.UpdateByID(ctx, nft.Id, update)
	if err != nil {
		log.Println(err.Error())
		return nft.Id, err
	}
	return nft.Id, nil
}

// func FindProjectByName(name string) *Project {
// 	ctx := context.Background()

// 	// Reading a sigle document from a collection
// 	var project Project
// 	if err := projects.FindOne(ctx, bson.M{"name": name}).Decode(&project); err != nil {
// 		log.Fatal(err)
// 	}
// 	return &project
// }

// func FindAllProjects() []Project {
// 	ctx := context.Background()

// 	cursor, err := projects.Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var projects []Project
// 	if err = cursor.All(ctx, &projects); err != nil {
// 		log.Fatal(err)
// 	}
// 	return projects
// }

func FindMintEvents(project, chain string, limit, page int64) []Event {
	filter := bson.M{"project": project, "action": "Mint", "chain": chain}
	collection := events
	// projection := bson.M{"project": project, "action": "Mint", "chain": chain}
	// Querying paginated data
	// Sort and select are optional
	// Multiple Sort chaining is also allowed
	// If you want to do some complex sort like sort by score(weight) for full text search fields you can do it easily
	// sortValue := bson.M{
	//		"$meta" : "textScore",
	//	}
	// aggPaginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort("score", sortValue)...
	var es []Event
	_, err := paginate.New(collection).Context(context.Background()).Limit(limit).Page(page).Filter(filter).Decode(&es).Find()
	if err != nil {
		log.Panicln("FindMintEvents:", err, project)
		// panic(err)
	}
	return es
	// ctx := context.Background()

	// l := limit
	// skip := (page+1)*limit - limit
	// fOpt := options.FindOptions{Limit: &l, Skip: &skip}
	// cursor, err := events.Find(ctx, bson.M{"project": project, "action": "Mint", "chain": chain}, &fOpt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var es []Event
	// if err = cursor.All(ctx, &es); err != nil {
	// 	log.Fatal(err)
	// }
	// return es
}

// func FindAllEvents(project, chain string, limit, page int64) []Event {
// 	filter := bson.M{"project": project, "chain": chain}
// 	var es []Event
// 	_, err := paginate.New(events).Context(context.Background()).Limit(limit).Page(page).Filter(filter).Decode(&es).Find()
// 	if err != nil {
// 		log.Panicln("FindAllEvents:", err, project)
// 		// panic(err)
// 	}
// 	return es
// }

// func CountMintEvents(project, chain string) int64 {
// 	ctx := context.Background()
// 	count, err := events.CountDocuments(ctx, bson.M{"project": project, "action": "Mint", "chain": chain})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return count
// }

// func CountAllEvents(project, chain string) int64 {
// 	ctx := context.Background()
// 	count, err := events.CountDocuments(ctx, bson.M{"project": project, "chain": chain})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return count
// }

func countEvents(transactionHash string, logIndex uint) (int64, error) {
	return events.CountDocuments(context.Background(), bson.M{"transactionHash": transactionHash, "logIndex": logIndex})
}

func SaveEvents(es map[string]Event) {
	// ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Minute)
	// defer cancel()
	ctx := context.Background()
	ces := make([]mongo.WriteModel, 0)
	for _, e := range es {
		if count, err := countEvents(e.TransactionHash, e.LogIndex); err == nil && count == 0 {
			ces = append(ces, mongo.NewInsertOneModel().SetDocument(e))
		}
	}
	if len(ces) > 0 {
		opts := options.BulkWrite().SetOrdered(false)
		results, err := events.BulkWrite(ctx, ces, opts)
		if err != nil {
			log.Println("SaveEvents panic error", err)
		}
		log.Printf("SaveEvents i: %d ", results.InsertedCount)
	}
	log.Printf("SaveEvents s: %d ", len(es))
}

func CountNFTs(id string) (int64, error) {
	return nfts.CountDocuments(context.Background(), bson.M{"_id": id})
}

func SaveNFTs(events []NFT) {
	ctx := context.Background()
	uris := make([]interface{}, 0)
	for _, nft := range events {
		if count, err := nfts.CountDocuments(ctx, bson.M{"_id": nft.Id}); err == nil && count == 0 {
			// log.Println("not found nft", nft.Id)
			uris = append(uris, nft)
		} else {
			_, err := nfts.UpdateByID(ctx, nft.Id, bson.M{"$set": bson.M{
				"owner":             nft.Owner,
				"expires":           nft.Expires,
				"updatedAt":         time.Now(),
				"registeredAtBlock": nft.RegisteredAtBlock,
				"latestBlockNumber": nft.RegisteredAtBlock,
			}})
			if err != nil {
				log.Println("SaveNFTs UpdateByID panic error:", err, nft.Id)
			}
			// log.Println("SaveNFTs UpdateByID", u.ModifiedCount, nft.TokenIdDec)
		}
	}
	if len(uris) > 0 {
		_, err := nfts.InsertMany(ctx, uris, options.InsertMany().SetOrdered(false))
		if err != nil {
			log.Println("SaveNFTs InsertMany panic error:", err)
			for n := range uris {
				nfts.InsertOne(ctx, n)
			}
		}
	}
}

// func SaveNFTByMap(events map[string]NFT) {
// 	ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Second)
// 	defer cancel()
// 	uris := make([]interface{}, 0)
// 	for _, nft := range events {
// 		if count, err := nfts.CountDocuments(ctx, bson.M{"_id": nft.Id}); err == nil && count == 0 {
// 			uris = append(uris, nft)
// 		} else {
// 			_, err := nfts.UpdateByID(ctx, nft.Id, bson.M{"$set": bson.M{
// 				"owner":             nft.Owner,
// 				"expires":           nft.Expires,
// 				"updatedAt":         time.Now(),
// 				"registeredAtBlock": nft.RegisteredAtBlock,
// 				"latestBlockNumber": nft.RegisteredAtBlock,
// 			}})
// 			if err != nil {
// 				log.Println("SaveNFTs UpdateByID panic error:", err, nft.Id)
// 			}
// 		}
// 	}
// 	if len(uris) > 0 {
// 		_, err := nfts.InsertMany(ctx, uris, options.InsertMany().SetOrdered(false))
// 		if err != nil {
// 			log.Println("SaveNFTs InsertMany panic error:", err)
// 		}
// 	}
// }

func UpdateNFTByMap(events map[string]NFT) {
	// ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Minute)
	// defer cancel()
	ctx := context.Background()
	uris := make([]mongo.WriteModel, 0)
	for _, nft := range events {
		if count, err := nfts.CountDocuments(ctx, bson.M{"_id": nft.Id}); err == nil && count == 0 {
			// log.Println("not found nft", nft.Id)
			uris = append(uris, mongo.NewInsertOneModel().SetDocument(nft))
		} else {
			umm := mongo.NewUpdateManyModel().SetFilter(bson.M{"_id": nft.Id})
			update := umm.SetUpdate(bson.M{"$set": bson.M{
				"owner":             nft.Owner,
				"expires":           nft.Expires,
				"updatedAt":         time.Now(),
				"registeredAtBlock": nft.RegisteredAtBlock,
				"latestBlockNumber": nft.RegisteredAtBlock,
			}})
			if len(nft.Name) > 0 {
				update = umm.SetUpdate(bson.M{"$set": bson.M{
					"name":              nft.Name,
					"owner":             nft.Owner,
					"expires":           nft.Expires,
					"updatedAt":         time.Now(),
					"registeredAtBlock": nft.RegisteredAtBlock,
					"latestBlockNumber": nft.RegisteredAtBlock,
				}})
			}
			uris = append(uris, update)
		}
	}
	if len(uris) > 0 {
		opts := options.BulkWrite().SetOrdered(false)
		results, err := nfts.BulkWrite(ctx, uris, opts)
		if err != nil {
			log.Println("UpdateNFTByMap panic error", err)
		}
		log.Printf("UpdateNFTByMap l: %d i: %d u: %d", len(events), results.InsertedCount, results.ModifiedCount)
	}
}

func FindNFTs(limit, page int64) []NFT {
	filter := bson.M{}
	collection := nfts
	var es []NFT
	_, err := paginate.New(collection).Context(context.Background()).Limit(limit).Page(page).Filter(filter).Decode(&es).Find()
	if err != nil {
		log.Panicln("FindNFTs:", err, limit, page)
	}
	return es
}

func NFTCount() int64 {
	count, _ := nfts.EstimatedDocumentCount(context.Background())
	return count
}

func CountUDDomains(chain string) int64 {
	count, _ := config.GetCollection("Domain_Collection").CountDocuments(context.Background(), bson.M{"project": "Unstoppable Domains", "chain": chain})
	return count
}

func FindUDDomains(chain string, limit, page int64) []NFT {
	filter := bson.M{"project": "Unstoppable Domains", "chain": chain}
	collection := config.GetCollection("Domain_Collection")
	var es []NFT
	_, err := paginate.New(collection).Context(context.Background()).Limit(limit).Page(page).Filter(filter).Decode(&es).Find()
	if err != nil {
		log.Panicln("FindUDDomains:", err, page)
	}
	return es
}

func FindBitDomains() []NFT {
	filter := bson.M{"project": "dotbit", "chain": "ETHEREUM", "name": ""}
	var es []NFT
	_, err := paginate.New(nfts).Context(context.Background()).Limit(10000).Page(0).Filter(filter).Decode(&es).Find()
	if err != nil {
		log.Panicln("FindBitDomains:", err, 0)
	}
	return es
}
