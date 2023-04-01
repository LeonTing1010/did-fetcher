package models

/*
NFT Contracts Collection
Collection_Collection
*/
type Contract struct {
	Id                  string `bson:"_id"`
	Address             string `bson:"address,omitempty"`
	Chain               string `bson:"chain,omitempty"`
	ChainId             int    `bson:"chainId,omitempty"`
	Parent              string `bson:"parent,omitempty"` // the parent of the domain, i.e. abc.def.eth  has parent def.eth
	Tld                 string `bson:"tld,omitempty"`    // the top level domain, i.e. for any domains in ENS, eth is the tld
	Project             string `bson:"project,omitempty"`
	TokenType           string `bson:"tokenType,omitempty"` // ERC-721 or ERC-1155
	Symbol              string `bson:"symbol,omitempty"`
	Name                string `bson:"name,omitempty"` // NFT contract name
	DeployedBlockNumber uint64 `bson:"deployedBlockNumber,omitempty"`
	LatestBlockNumber   uint64 `bson:"latestBlockNumber,omitempty"`
}

func (c Contract) Network() string {
	if c.ChainId == 1 || c.ChainId == 137 || c.ChainId == 1030 {
		return "main"
	}
	return "test"
}

func (c Contract) DID() string {
	return c.Symbol
}

func (c Contract) Blockchain() string {
	return c.Chain
}

func (c Contract) Contract() string {
	return c.Address
}

func (c *Contract) CopyFrom(contract Contract) {
	if len(contract.Name) > 0 {
		c.Name = contract.Name
	}
	if len(contract.TokenType) > 0 {
		c.TokenType = contract.TokenType
	}
	if contract.DeployedBlockNumber > 0 {
		c.DeployedBlockNumber = contract.DeployedBlockNumber
	}
	if len(c.Symbol) > 0 {
		c.Symbol = contract.Symbol
	}
}
