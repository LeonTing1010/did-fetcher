package nodes

type ContractMeta interface {
	Name() string
	Address() string
	Symbol() string
	TokenType() string
	DeployedBlockNumber() uint64
	SetDeployedBlockNumber(blocknumber uint64)
}

type CMetadata struct {
	Csymbol              string `json:"symbol"`
	Cname                string `json:"name"`
	Contract             string
	CtokenType           string `json:"tokenType"`
	CdeployedBlockNumber uint64 `json:"deployedBlockNumber"`
	ChainId              uint64
}

func (meta CMetadata) Name() string {
	return meta.Cname
}

func (meta CMetadata) Address() string {
	return meta.Contract
}

func (meta CMetadata) Symbol() string {
	return meta.Csymbol
}

func (meta CMetadata) TokenType() string {
	return meta.CtokenType
}

func (meta CMetadata) DeployedBlockNumber() uint64 {
	return meta.CdeployedBlockNumber
}

func (meta *CMetadata) SetDeployedBlockNumber(blocknumber uint64) {
	meta.CdeployedBlockNumber = blocknumber
}
