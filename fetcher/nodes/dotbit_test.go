package nodes

import (
	"testing"

	"fetcher/contracts"

	"github.com/stretchr/testify/assert"
)

var Dotbit = "0x60eb332bd4a0e2a9eeb3212cfdd6ef03ce4cb3b5"

func TestDotbitGetContractMetadata(t *testing.T) {
	t.Log("Test GetContractMetadata")
	// Initialize Alchemy node
	alchemy := New(contracts.BIT, "ETHEREUM")
	meta, err := alchemy.ContractMetadata(Dotbit)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(meta)
	assert.True(t, meta.DeployedBlockNumber() > 0)
}

func TestDotbitGetNftMetadata(t *testing.T) {
	name := "coinexchange.bit"
	tokenId := "1379569504140102960579679260191560872791005714918"
	// Initialize Alchemy node
	alchemy := New(contracts.BIT, "ETHEREUM")
	tid, _ := contracts.Str2BigInt(tokenId)
	meta, err := alchemy.NFTmeta(alchemy.Chain(), Dotbit, tid)
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, meta.Name() == name)
}

func TestAccount(t *testing.T) {
	t.Log("Test Account")
	account := "phone.bit"
	accountInfo, err := Account(account)
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, accountInfo.Account == account)
}

func TestRecords(t *testing.T) {
	t.Log("Test Records")
	account := "justindrewbieber.bit"
	records, err := Records(account)
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, records.Account == account)
}
