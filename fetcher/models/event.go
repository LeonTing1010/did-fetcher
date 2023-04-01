package models

import (
	"fmt"
	"math/big"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Activity_Collection
*/
type Event struct {
	Id              primitive.ObjectID `bson:"_id"`
	Time            time.Time          `bson:"time,omitempty" ` // datetime of the action
	To              string             `bson:"to,omitempty" `
	From            string             `bson:"from,omitempty" `
	Action          string             `bson:"action,omitempty" ` // Actions such as Mint, Transfer, Burn (currently omitted), Register, Renew
	BlockTime       int64              `bson:"blockTime,omitempty" `
	BlockNumber     uint64             `bson:"blockNumber,omitempty" `
	Project         string             `bson:"project,omitempty" ` // project name (ref: Project_Collection _id)
	Tld             string             `bson:"tld,omitempty" `     // the top level domain, i.e. for any domains in ENS, eth is the tld
	ChainId         int                `bson:"chainId,omitempty" `
	Chain           string             `bson:"chain,omitempty" `
	TokenId         string             `bson:"tokenId,omitempty" `    // hex string of tokenId
	TokenIdDec      string             `bson:"tokenIdDec,omitempty" ` // a decimal string of tokens
	ContractAddress string             `bson:"contractAddress,omitempty" `
	TransactionHash string             `bson:"transactionHash,omitempty" `
	DomainInfo      string             `bson:"domainInfo,omitempty" `    //(ref: Domain_Collection _id)
	Fee             string             `bson:"payment.fee,omitempty" `   // amount of tokens paid during register/renew, exists when the action is register/renew
	Token           string             `bson:"payment.token,omitempty" ` // contractaddress of token (ref: Token_Collection _id), exists when the action is register/renew
	Expires         uint64             `bson:"expires,omitempty" `       // blocktime of the expiry. exists when the action is renew action
	ExpiresTime     string             `bson:"expiresTime,omitempty" `   // timestamp of the expiry. exists when the action is renew action
	LogIndex        uint               `bson:"logIndex,omitempty" `
}

func NewEvent(project string, from string, to string, blockNumber uint64, tokenId *big.Int, chain string, chainId int, contractAddress string, tld string, action string, transactionHash string, logIndex uint) Event {
	event := Event{
		Time:            time.Now(),
		Project:         project,
		From:            from,
		To:              to,
		BlockNumber:     blockNumber,
		ContractAddress: contractAddress,
		DomainInfo:      chain + ":" + contractAddress + ":" + tokenId.String(),
		TokenId:         fmt.Sprintf("0x%x", tokenId),
		TokenIdDec:      tokenId.String(),
		Chain:           chain,
		Tld:             tld,
		ChainId:         chainId,
		Action:          action,
		TransactionHash: transactionHash,
		LogIndex:        logIndex,
	}
	event.Id = primitive.NewObjectID()
	return event
}
