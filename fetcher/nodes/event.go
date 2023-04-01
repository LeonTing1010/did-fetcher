package nodes

import "math/big"

const (
	Mint     = "Mint"
	Burn     = "Burn"
	Transfer = "Transfer"
)

type UserRecycleEvent struct {
	BlockNumber      uint64
	TransactionHash  string
	TransactionIndex uint
	BlockHash        string
	LogIndex         uint
	Contract         string
	UUID             string
	Event            string `default:"UserRecycle"`
}

type TransferEvent struct {
	BlockNumber      uint64
	TransactionHash  string
	TransactionIndex uint
	BlockHash        string
	LogIndex         uint
	Contract         string
	From             string
	To               string
	TokenId          *big.Int
	Event            string `default:"Transfer"`
}

type NewURIEvent struct {
	BlockNumber      uint64
	TransactionHash  string
	TransactionIndex uint
	BlockHash        string
	LogIndex         uint
	Contract         string
	Uri              string
	TokenId          *big.Int
	From             string
	To               string
}
