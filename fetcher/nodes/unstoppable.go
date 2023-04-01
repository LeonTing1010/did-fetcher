package nodes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"fetcher/config"
)

// Create a Bearer string by appending string access token
var bearer = "Bearer " + config.Env("UNSTOPPABLE_API_KEY")

/*
{
  "meta": {
    "domain": "halfcharge.wallet",
    "tokenId": "191287042669689698626981458362724906155142764155885548527830900898316",
    "namehash": "0x0000000718610ae62730f5a3526399878fa39bebba77a2e988e8bc7551d77a0c",
    "blockchain": "MATIC",
    "networkId": 137,
    "owner": "0x5295c84e2f133957c6eb732ab156f0b92a449f96",
    "resolver": "0xa9a6a3626993d487d2dbda3173cf58ca1a9d9e9f",
    "registry": "0xa9a6a3626993d487d2dbda3173cf58ca1a9d9e9f",
    "reverse": false
  },
  "records": {
    "crypto.ETH.address": "0x5295c84e2f133957c6eb732ab156f0b92a449f96",
    "crypto.MATIC.version.ERC20.address": "0x5295c84e2f133957c6eb732ab156f0b92a449f96",
    "crypto.MATIC.version.MATIC.address": "0x5295c84e2f133957c6eb732ab156f0b92a449f96"
  }
}
*/

type DomainInfo struct {
	Meta struct {
		Domain     string
		TokenId    string
		Namehash   string
		Blockchain string // MATIC
		NetworkId  int    // 137
		Owner      string
		Resolver   string
		Registry   string
		Reverse    bool // false
	}
	Records  map[string]string
	Contract string
}

// https://resolve.unstoppabledomains.com/domains/{domainName}
func Domain(domainName string) (*DomainInfo, error) {
	// Create a new request using http
	req, _ := http.NewRequest("GET", config.UnstoppableURI()+"/domains/"+domainName, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var domain DomainInfo
	json.Unmarshal(responseData, &domain)
	return &domain, nil
}

func (meta UDMetadata) Level() int {
	for _, attr := range meta.Attributes {
		if attr.Trait_type == "Level" {
			return int(attr.Value.(float64))
		}
	}
	return 1
}

func (meta UDMetadata) Length() int {
	for _, attr := range meta.Attributes {
		if attr.Trait_type == "Length" {
			return int(attr.Value.(float64))
		}
	}
	return len(meta.Name())
}

func (meta UDMetadata) Parent() string {
	ds := strings.Split(meta.URI(), ".")
	if len(ds) > 1 {
		var parent string
		for i := 1; i < len(ds); i++ {
			parent = parent + ds[i]
		}
		return parent
	}
	return meta.URI()
}

func (meta UDMetadata) TLD() string {
	ds := strings.Split(meta.URI(), ".")
	return ds[len(ds)-1]
}

func (meta UDMetadata) Expire() uint64 {
	return 0
}

func (meta UDMetadata) URI() string {
	return meta.Mname
}

func (meta UDMetadata) TokenId() string {
	return meta.MtokenId
}

func (meta UDMetadata) Name() string {
	return meta.Mname
}

func (meta UDMetadata) ContractAddress() string {
	return meta.Mcontract
}

func (meta UDMetadata) TTL() uint64 {
	return 0
}

func (meta UDMetadata) Owner() string {
	return meta.Mowner
}

func (meta UDMetadata) Id() string {
	return meta.Uid
}

func (meta *UDMetadata) SetOwner(owner string) {
	meta.Mowner = owner
}

func (meta *UDMetadata) SetId(id string) {
	meta.Uid = id
}

type UDMetadata struct {
	Uid          string `bson:"_id"`
	Mname        string `json:"name" bson:"name"`
	MtokenId     string `json:"tokenId" bson:"tokenId"`
	Namehash     string
	Description  string
	External_url string
	Image        string
	Image_url    string
	Mcontract    string `json:"contract" bson:"contract"`
	Mowner       string `json:"owner"  bson:"owner"`
	Attributes   []Attribute
}

/*
{
  "name": "mycashinminutes.crypto",
  "tokenId": "92596372731794458187264453803238509625972004707291899261635368063849062894708",
  "namehash": "0xccb7b090ae612d8a91ea697823029613f4485e497d2fccb349ab1389eb3a8c74",
  "description": "A CNS or UNS blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
  "external_url": "https://unstoppabledomains.com/search?searchTerm=mycashinminutes.crypto",
  "image": "https://metadata.unstoppabledomains.com/image-src/mycashinminutes.crypto.svg",
  "image_url": "https://metadata.unstoppabledomains.com/image-src/mycashinminutes.crypto.svg",
  "attributes": [
    {
      "trait_type": "Ending",
      "value": "crypto"
    },
    {
      "trait_type": "Level",
      "value": 2
    },
    {
      "trait_type": "Length",
      "value": 15
    },
    {
      "trait_type": "Subdomains",
      "value": 0
    },
    {
      "trait_type": "Type",
      "value": "standard"
    },
    {
      "trait_type": "Character Set",
      "value": "letter"
    }
  ],
  "background_color": "4C47F7"
}
*/
// https://resolve.unstoppabledomains.com/metadata/alex011221a2.coin
func UDMeta(domainOrToken string) (Metadata, error) {
	// Create a new request using http
	req, _ := http.NewRequest("GET", config.UnstoppableURI()+"/metadata/"+domainOrToken, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var meta UDMetadata
	json.Unmarshal(responseData, &meta)
	return &meta, nil
}

/*
	{
	  "data": [
	    {
	      "domain": "halfcharge.wallet",
	      "from": "0x0000000000000000000000000000000000000000",
	      "to": "0x5295C84E2f133957c6EB732aB156F0B92a449f96",
	      "networkId": 137,
	      "blockNumber": 21489466,
	      "blockchain": "MATIC"
	    }
	  ]
	}
*/
type TransferData struct {
	Data []Event
}
type Event struct {
	Domain      string
	From        string
	To          string
	NetworkId   int
	BlockNumber uint64
	Blockchain  string
}

// https://resolve.unstoppabledomains.com/domains/{domainName}/transfers/latest
func LatestTransfer(domainName string) (*TransferData, error) {
	// Create a new request using http
	req, _ := http.NewRequest("GET", config.UnstoppableURI()+"/domains/"+domainName+"/transfers/latest", nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var nsObject TransferData
	json.Unmarshal(responseData, &nsObject)
	return &nsObject, nil
}
