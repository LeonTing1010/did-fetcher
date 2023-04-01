package models

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

/*
Token Sets in NFT Contracts
Domain_Collection
{"_id":"BNB:0xe3b1d32e43ce8d658368e2cbff95d57ef39be8a6:24293655359084398983762148957951887367734741115283469135310733710692111186993",
"collectionInfo":"BNB:0xe3b1d32e43ce8d658368e2cbff95d57ef39be8a6",
"project":"SpaceId",
"contractAddress":"0xe3b1d32e43ce8d658368e2cbff95d57ef39be8a6",
"tokenId":"0x35b5b8bece53958bb309db665734c38515f37439f69bfdbc64808f1af9a97c31",
"tokenIdDec":"24293655359084398983762148957951887367734741115283469135310733710692111186993",
"chainId":{"$numberInt":"56"},
"chain":"BNB",
"name":"000",
"tld":"bnb",
"parent":"bnb",
"level":{"$numberInt":"2"},
"owner":"0xe0a9e5b59701a776575fdd6257c3f89ae362629a",
"registeredAtBlock":{"$numberInt":"19673241"},
"registeredAt":{"$numberInt":"1658191859"},
"expires":{"$numberInt":"1689748811"},
"length":{"$numberInt":"3"},
"segment":{"$numberInt":"3"},
"characterSet":"Digit",
"indexed":true,
"valid":true,
"createdAt":{"$date":{"$numberLong":"1665338273063"}},
"updatedAt":{"$date":{"$numberLong":"1674243295743"}},
"__v":{"$numberInt":"0"}}
*/
type NFT struct {
	Id                string    `bson:"_id"`
	CollectionInfo    string    `bson:"collectionInfo,omitempty"`
	Project           string    `bson:"project,omitempty"`
	ContractAddress   string    `bson:"contractAddress,omitempty"`
	TokenId           string    `bson:"tokenId,omitempty"`
	TokenIdDec        string    `bson:"tokenIdDec"`
	ChainId           int       `bson:"chainId,omitempty"`
	Chain             string    `bson:"chain,omitempty"`
	Name              string    `bson:"name,omitempty"`
	Tld               string    `bson:"tld,omitempty"`
	Parent            string    `bson:"parent,omitempty"`
	Level             int       `bson:"level,omitempty"`
	Owner             string    `bson:"owner,omitempty"`
	RegisteredAtBlock uint64    `bson:"registeredAtBlock,omitempty"`
	LatestBlockNumber uint64    `bson:"latestBlockNumber,omitempty"`
	RegisteredAt      string    `bson:"registeredAt,omitempty"`
	Expires           uint64    `bson:"expires"`
	Length            int       `bson:"length"`
	Segment           string    `bson:"segment"`
	CharacterSet      string    `bson:"characterSet"`
	CreatedAt         time.Time `bson:"createdAt"`
	UpdatedAt         time.Time `bson:"updatedAt"`
}

func NewNFT(project string, uri string, contract string, tokenId *big.Int, chain string, chainId int, owner string, registeredAtBlock uint64) NFT {
	nft := NFT{
		Project:           project,
		ContractAddress:   contract,
		CollectionInfo:    chain + ":" + contract,
		TokenId:           fmt.Sprintf("0x%x", tokenId),
		TokenIdDec:        tokenId.String(),
		Chain:             chain,
		Level:             Level(uri),
		Owner:             owner,
		ChainId:           chainId,
		Length:            Length(uri),
		RegisteredAtBlock: registeredAtBlock,
		LatestBlockNumber: registeredAtBlock,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	nft.Tld = TLD(uri)
	nft.Name = uri
	nft.Parent = Parent(uri)
	nft.Id = nft.CollectionInfo + ":" + nft.TokenIdDec
	return nft
}

func NFTID(chain, contract, tokenId string) string {
	return chain + ":" + contract + ":" + tokenId
}

func (nft *NFT) SetExpires(expires uint64) {
	nft.Expires = expires
}

func (nft *NFT) SetOwner(owner string) {
	nft.Owner = owner
}

func (nft NFT) TLD() string {
	return nft.Tld
}

func Level(name string) int {
	ds := strings.Split(name, ".")
	return len(ds)
}

func Length(name string) int {
	return len(name)
}

func Parent(name string) string {
	ds := strings.Split(name, ".")
	if len(ds) > 1 {
		var parent string
		if len(ds) > 1 {
			parent = strings.TrimPrefix(name, ds[0]+".")
		}
		return parent
	}
	return name
}

func TLD(name string) string {
	ds := strings.Split(name, ".")
	return ds[len(ds)-1]
}
