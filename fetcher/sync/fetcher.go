package sync

import (
	"log"
	"math/big"

	"fetcher/config"
	"fetcher/models"
	"fetcher/nodes"
)

// It fetches Non-Fungible Tokens (NFTs) from the given project.
// It finds the contracts associated with the project, then fetches the metadata associated with the contract.
// It then gets the block number and compares it to the latest block number stored in the database.
// If the block number is greater than the stored number,
// it fetches the token transfers from that block number range and stores the information in the database.
// It also inserts the account info and records associated with the URI.
func FetchHistory(did string) {
	project := did
	chain := config.Blockchain()
	log.Printf("FetchNFTs => %s %s", project, chain)
	// Finds all contracts related to the project
	contracts, err := models.FindContractsBySymbol(project)
	if err != nil {
		log.Fatalln(err.Error())
	}
	l := len(contracts)
	if l == 0 {
		log.Fatalln("Can't find " + project + " in Collection")
	}
	// Loops through all the contracts related to the project
	for _, c := range contracts {
		if c.Chain != chain {
			continue
		}
		node := nodes.New(c.DID(), c.Blockchain())
		// Fetches the new URIs
		log.Println("FetchNFTsByContract", node.Info(), c.Address, node.BlockNumber())
		go node.FetchNFTsByContract(c)
	}
}

func FetchExpireAndOwner(did string) {
	project := did
	chain := config.Blockchain()
	// Finds all contracts related to the project
	contracts, err := models.FindContractsBySymbol(project)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Loops through all the contracts related to the project
	for _, c := range contracts {
		if c.Chain != chain {
			continue
		}
		node := nodes.New(c.DID(), c.Blockchain())
		// Fetches the new URIs
		log.Println("FetchExpireAndOwner", node.Info(), c.Address, node.BlockNumber())
		node.UpdateExpireAndOwner(c.Address)
	}
}

// It fetches contract metadata from a given project.
// It fetches the contract data from the models and then attempts to get the contract metadata from the nodes.
// If successful, it inserts the contract metadata and updates the contract with the new data.
func FetchContractMetadata(did string) {
	project := did
	chain := config.Blockchain()
	// This code prints information about contracted projects, creates a new alchemy from the contract, and inserts relevant metadata.
	log.Printf("FetchContractMetadata: %s %s", project, chain)
	contracts, err := models.FindContractsBySymbol(project)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(contracts) == 0 {
		log.Fatalln("Can't find " + project + " in Collection")
	}

	for _, c := range contracts {
		if c.Chain != chain {
			continue
		}
		// Skip if contract was not deployed or updated
		if c.DeployedBlockNumber == 0 || c.DeployedBlockNumber < c.LatestBlockNumber {
			continue
		}
		node := nodes.New(c.DID(), c.Blockchain())
		cMetadata, err := node.ContractMetadata(c.Address)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if c.LatestBlockNumber == 0 {
			c.LatestBlockNumber = cMetadata.DeployedBlockNumber()
		}
		// log.Println(cMetadata)
		// cMetadata.SetDeployedBlockNumber(c.DeployedBlockNumber)
		// Update the contract model with data from the metadata.
		c.CopyFrom(models.Contract{
			Name:                cMetadata.Name(),
			TokenType:           cMetadata.TokenType(),
			Symbol:              cMetadata.Symbol(),
			DeployedBlockNumber: c.DeployedBlockNumber,
			LatestBlockNumber:   c.LatestBlockNumber,
		})
		models.UpdateOneContract(c)
	}
}

func FetchNewBlock(did string) {
	project := did
	chain := config.Blockchain()
	log.Printf("FetchNewBlock: %s %s", project, chain)
	contracts, err := models.FindContractsBySymbol(project)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(contracts) == 0 {
		log.Fatalln("Can't find " + project + " in Collection")
	}
	for _, c := range contracts {
		if c.Chain != chain {
			continue
		}
		node := nodes.New(c.DID(), c.Blockchain())
		from := node.BlockNumber()
		step := node.Step()
		node.Subscribe(func(latest *big.Int) {
			to := latest.Uint64()
			if to-from > step/10 {
				log.Printf("Sync New block %s from:%d to:%d", node.Info(), from, to)
				node.FetchNFTsFromLatest(c, from, to)
				from = to
			}
		})
	}
}

// // It fetches all token transfer events for a given project and inserts them into the database
// func fetchEvents(project string) {
// 	contracts, err := models.FindContractsByName(project)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	if len(contracts) == 0 {
// 		log.Fatalln("Can't find " + project + " in Collection")
// 	}

// 	for _, c := range contracts {
// 		// Get start and end block for token transfers
// 		// Get all token transfer events between given blocks
// 		fetchEventsByContract(c)
// 	}
// }

// func fetchEventsByContract(c models.Contract) {
// 	alchemy := nodes.New(c.Chain)
// 	blockNumber := alchemy.BlockNumber()
// 	if blockNumber <= 0 {
// 		return
// 	}

// 	from := c.LatestBlockNumber
// 	log.Printf("FetchEvents contract:%s chain:%s from:%d block:%d", c.Symbol, c.Chain, from, blockNumber)

// 	if err := fetchEventsByBlockNumber(alchemy, c, from, blockNumber); err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// }

// func fetchEventsByBlockNumber(alchemy *nodes.Alchemy, c models.Contract, from uint64, blockNumber uint64) error {
// 	to := from + alchemy.Step()
// 	if to > blockNumber {
// 		to = blockNumber
// 	}
// 	events, err := alchemy.GetTokenTransfers(c.Project, from, to, c.Address)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	for len(events) >= 0 && to <= blockNumber {
// 		for _, event := range events {
// 			nodes.SaveNFTEvents(&event)
// 			meta, err := alchemy.Metadata(c.Address, event.TokenId.String())
// 			if err != nil {
// 				log.Println(err.Error())
// 				continue
// 			}
// 			log.Printf("FetchEventsByBlockNumber block:%d  uri:%s tokenId:%s", blockNumber, meta.URI(), meta.TokenId())
// 			models.SaveOneEvent(models.NewEvent(c.Project, event.From, event.To, event.BlockNumber, event.TokenId, c.Chain, c.ChainId, c.Address, meta.TLD(), event.Event))
// 			// update nft latest blocknumber
// 			if _, err := models.UpdateOneNFT(
// 				&models.NFT{
// 					Id:                models.NFTID(c.Chain, c.Address, event.TokenId.String()),
// 					LatestBlockNumber: to,
// 				},
// 			); err != nil {
// 				log.Println(err.Error())
// 			}
// 		}

// 		from = to
// 		to = to + alchemy.Step()
// 		events, err = alchemy.GetTokenTransfers(c.Project, from, to, c.Address)
// 		if err != nil {
// 			log.Println(err.Error())
// 			continue
// 		}
// 	}
// 	return nil
// }
