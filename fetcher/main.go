package main

import (
	"log"
	"os"

	"fetcher/config"
	"fetcher/models"
	"fetcher/sync"
)

// This code sets up the environment to sync blockchain data from the given  DID.
// It sets the two environment variables 'CHAIN' and 'LOG_FILE_NAME' from the command line and configures the model.
// It then begins to fetch the contract metadata, NFTs and events for the given DID,
// and then proceeds to fetch new blocks from the blockchain.
func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please enter the DID project symbol & chain")
	}
	DID := os.Args[1]        // project symbol
	BLOCKCHAIN := os.Args[2] // main or test
	os.Setenv("BLOCKCHAIN", BLOCKCHAIN)
	os.Setenv("LOG_FILE_NAME", DID+"_"+BLOCKCHAIN+".log")
	config.SetUp()
	models.Setup(DID)
	// sync.FetchContractMetadata(DID)
	// sync.FetchHistory(DID)
	// sync.FetchNewBlock(DID)
	// select {}
	sync.FetchNFTsForCollection(DID)
	// sync.FetchExpireAndOwner(DID)
	// sync.MoveOldToNew(DID)
}
