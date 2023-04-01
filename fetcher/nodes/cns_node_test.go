package nodes

import (
	"os"
	"testing"

	"fetcher/config"
	"fetcher/contracts"
	"fetcher/models"

	"github.com/stretchr/testify/assert"
)

func TestCNSGetNftMetadata(t *testing.T) {
	t.Log("Test TestCNSGetNftMetadata")
	// tokenId, err := FindNFTTokenId("alex011221a2.coin")
	// if err != nil {
	// 	t.Log(err.Error())
	// 	return
	// }
	contract := "0xcab10736e8aa9498ad772f68350ce501746caabe"
	tokenId := "67471485830746893924574146644674603542414160274959919404633156362225377486826"
	conflex := New("CNS", CONFLUX)
	tid, _ := contracts.Str2BigInt(tokenId)
	meta, err := conflex.NFTmeta(conflex.Chain(), contract, tid)
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, meta.TokenId() == tokenId)
}

func TestCNSGetContractMetadata(t *testing.T) {
	c := "0x7b29dac7b5d3a90696c8b059cbf1fae9af32fa28"
	conflex := New(contracts.CNS, "Conflex")
	meta, err := conflex.rpc.ContractMetadata(c)
	if err != nil {
		t.Log(err.Error())
		return
	}
	assert.True(t, meta.DeployedBlockNumber() > 0)
}

func TestFetchNFTsByBlockNumber(t *testing.T) {
	os.Setenv("NETWORK", "main")
	os.Setenv("BLOCKCHAIN", CONFLUX)
	config.SetUp()
	did := contracts.CNS
	models.Setup(did)
	c := "0xcab10736e8aa9498ad772f68350ce501746caabe"
	contract, _ := models.FindContractByAddress(c)
	node := New(did, CONFLUX)
	node.FetchNFTsFromLatest(contract, 55922005, 66377205)
	select {}
}
