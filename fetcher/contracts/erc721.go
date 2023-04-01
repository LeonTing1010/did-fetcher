package contracts

import (
	"errors"

	"github.com/ethereum/go-ethereum/ethclient"
	cmap "github.com/orcaman/concurrent-map/v2"
)

const ZERO = "0x0000000000000000000000000000000000000000"

var nfts = cmap.New[NFT]()

func New(did, contract string, client *ethclient.Client) (NFT, error) {
	_, ok := nfts.Get(did)
	if !ok {
		if did == UD {
			nfts.Set(UD, NewUD(contract, client))
		} else if did == BIT {
			nfts.Set(BIT, NewBIT(contract, client))
		} else if did == CNS {
			nfts.Set(CNS, NewCNS(contract, client))
		} else if did == HK {
			nfts.Set(HK, NewKey(contract, client))
		} else if did == NNS {
			nfts.Set(NNS, NewNNS(contract, client))
		} else if did == ARB {
			nfts.Set(ARB, NewArb(contract, client))
		} else {
			return nil, errors.New("not found nft contract")
		}
	}
	nft, ok := nfts.Get(did)
	if ok {
		return nft, nil
	}
	return nil, errors.New("not found nft contract")
}
