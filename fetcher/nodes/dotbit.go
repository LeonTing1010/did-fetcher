package nodes

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"fetcher/config"
)

// coin_type：https://github.com/satoshilabs/slips/blob/master/slip-0044.md
// chain_id： https://github.com/ethereum-lists/chains
// type	Filled with "blockchain" for now
// coin_type 60: eth, 195: trx, 9006: bsc, 966: matic
// chain_id 	1: eth, 56: bsc, 137: polygon; 5: goerli, 97: bsct, 80001: mumbai
// account	Contains the suffix .bit in it
// key	Generally refers to the blockchain address for now

/*
	{
	  "errno": 0,
	  "errmsg": "",
	  "data": {
	    "out_point": {
	      "tx_hash": "0x36aa229e20134008dd688e59b955d6674c81016f6bda65375a8ef7712bc3f802",
	      "index": 0
	    },
	    "account_info": {
	        "account": "phone.bit",
			"account_alias": "phone.bit",
			"account_id_hex": "0x5f560ec1edc638d7dab7c7a1ca8c3b0f6ed1848b",
			"next_account_id_hex": "0x5f573a7184abd43632a7ff34b2a2dae413cc1f7f",
			"create_at_unix": 1666268687,
			"expired_at_unix": 1697804687,
			"status": 0,
			"das_lock_arg_hex": "0x049e628c5c452c56af9468fae59208aeb70971d1ad049e628c5c452c56af9468fae59208aeb70971d1ad",
			"owner_algorithm_id": 4,
			"owner_key": "TQQfpd5qZAEQJs74XBR3EiQfqPy6a2qDxJ",
			"manager_algorithm_id": 4,
			"manager_key": "TQQfpd5qZAEQJs74XBR3EiQfqPy6a2qDxJ"
	    }
	  }
	}
*/
type AccountInfoRpc struct {
	Errno  int
	Errmsg string
	Data   struct {
		Out_point struct {
			Tx_hash string
			Index   int
		}
		AccountInfo AccountInfo `json:"account_info"`
	}
}
type AccountInfo struct {
	Account              string
	Account_alias        string
	Account_id_hex       string
	Next_account_id_hex  string
	Create_at_unix       int64
	Expired_at_unix      int64
	Status               int
	Das_lock_arg_hex     string
	Owner_algorithm_id   int
	Owner_key            string
	Manager_algorithm_id int
	Manager_key          string
	Records              []Record
}

// curl -X POST https://indexer-v1.did.id/v1/account/info -d'{"account":"phone.bit"}'
func Account(account string) (*AccountInfo, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"account": account,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// Create a new request using http
	req, _ := http.NewRequest("POST", config.DotbitApiURI()+"/v1/account/info", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// // add authorization header to the req
	// req.Header.Add("Authorization", bearer)

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
	var rpc AccountInfoRpc
	json.Unmarshal(responseData, &rpc)

	return &rpc.Data.AccountInfo, nil
}

/*
	{
	  "errno": 0,
	  "errmsg": "",
	  "data": {
	    "account": "phone.bit",
	    "records": [
	      {
	        "key": "address.btc",
	        "label": "Personal account",
	        "value": "3EbtqPeAZbX6wmP6idySu4jc2URT8LG2aa",
	        "ttl": "300"
	      },
	      {
	        "key": "address.eth",
	        "label": "Personal account",
	        "value": "0x59724739940777947c56C4f2f2C9211cd5130FEf",
	        "ttl": "300"
	      }
	      // ...
	    ]
	  }
	}
*/
type RecordsRpc struct {
	Errno  int
	Errmsg string
	Data   AccountRecords `json:"data"`
}

type AccountRecords struct {
	Account string
	Records []Record
}
type Record struct {
	Key   string
	Label string
	Value string
	Ttl   string
}

func Records(account string) (*AccountRecords, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"account": account,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// Create a new request using http
	req, _ := http.NewRequest("POST", config.DotbitApiURI()+"/v1/account/records", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// // add authorization header to the req
	// req.Header.Add("Authorization", bearer)

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
	var rpc RecordsRpc
	json.Unmarshal(responseData, &rpc)

	return &rpc.Data, nil
}
