package nodes

import (
	"context"
	"log"
	"math/big"
	"time"

	"fetcher/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	MOONBEAM = "Moonbeam" // 1284
)

var (
	MBWS  = config.Env("MB_WSS_URI")
	MBURL = config.Env("MB_API_URI")
)

type Moonbeam struct {
	PlatON
}

func NewMoonbeam(chain string) *Moonbeam {
	moonbeam := Moonbeam{
		PlatON: PlatON{
			ws:  MBWS,
			url: MBURL,
		},
	}
	moonbeam.chain = MOONBEAM
	moonbeam.step = 500
	moonbeam.rawclient = mbClient()
	blockNumber, err := moonbeam.rawclient.BlockNumber(context.Background())
	if err != nil {
		log.Panic("can't get blocknumber: ", err.Error())
	}
	moonbeam.blockNumber = blockNumber
	return &moonbeam
}

func mbClient() *ethclient.Client {
	url := MBURL
	c, err := ethclient.Dial(url)
	log.Println("Dial ethclient", url)
	if err != nil {
		log.Panicln("Can't Dial Moonbeam", url)
	}
	return c
}

func (mb *Moonbeam) Subscribe(handle Handle) {
	wss := mb.ws
	log.Println("Open Websocket:", mb.Info(), wss)
	client, err := ethclient.Dial(wss)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ticker := time.NewTicker(mb.WaitTime())
	defer ticker.Stop()

	bn := mb.BlockNumber()
	for {
		<-ticker.C
		if bn%10 == 0 {
			handle(big.NewInt(int64(bn)))
		}
		bn, _ = client.BlockNumber(context.Background())
	}
}
