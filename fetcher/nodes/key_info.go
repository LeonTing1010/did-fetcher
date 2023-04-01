package nodes

import "math/big"

type KeyInfo interface {
	Expire() *big.Int
	TokenId() *big.Int
	ContractAddress() string
	Owner() string
	TTL() uint64
}

type CKeyInfo struct {
	expire   *big.Int
	tokenId  *big.Int
	contract string
	owner    string
	ttl      uint64
}

func (key CKeyInfo) Expire() *big.Int {
	return key.expire
}

func (key CKeyInfo) TokenId() *big.Int {
	return key.tokenId
}

func (key CKeyInfo) ContractAddress() string {
	return key.contract
}

func (key CKeyInfo) Owner() string {
	return key.owner
}

func (key CKeyInfo) TTL() uint64 {
	return key.ttl
}
