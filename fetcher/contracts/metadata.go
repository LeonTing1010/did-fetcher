package contracts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"fetcher/client"

	"golang.org/x/time/rate"
)

var c *client.RLHTTPClient

func init() {
	rl := rate.NewLimiter(rate.Every(10*time.Second), 10000) // 1000 request every 10 seconds
	c = client.NewClient(rl)
}

/*
	{
	  "name": "5207.bit",
	  "description": "5207.bit, Web3 identity for you and your community.\n https://did.id\n More about 5207.bit: https://bit.ly/3te6SOP",
	  "image": "https://display.did.id/erc721/card/774314209262442299961399623206984268714100117420",
	  "external_url": "https://did.id",
	  "attributes": [
	    {
	      "trait_type": "Expiration Date",
	      "display_type": "date",
	      "value": 1696515240000
	    },
	    {
	      "trait_type": "Registration Date",
	      "display_type": "date",
	      "value": 1664979240000
	    },
	    {
	      "trait_type": "Length",
	      "display_type": "number",
	      "value": 4
	    },
	    {
	      "trait_type": "Character Set",
	      "value": "digit"
	    },
	    {
	      "trait_type": "4D",
	      "value": "10K Club"
	    }
	  ]
	}
*/
type Meta struct {
	MName        string `json:"name"`
	Description  string
	Image        string
	External_url string
	Attributes   []Attribute
}

func (m Meta) Name() string {
	return m.MName
}

type Attribute struct {
	Value        any
	Trait_type   string
	Display_type string
}

type Metadata interface {
	Name() string
}

func getMetadata(tokenURI string) (Metadata, error) {
	req, _ := http.NewRequest("GET", tokenURI, nil)
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36")
	req.Header.Add("cookie", "_ga=GA1.1.998144915.1679562141; __cuid=1ef55001081d436287228b5488d623c7; amp_fef1e8=e6e06346-f560-4343-a5d9-9e29e6d43413R...1gscie1uo.1gscie1ve.q.6.10; _ga_3NLSMD7JJ6=GS1.1.1680061182.3.0.1680061182.0.0.0")
	res, err := c.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var meta Meta

	err = json.Unmarshal(data, &meta)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}
