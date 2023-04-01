package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	POLYGON  = "POLYGON"
	ETHEREUM = "ETHEREUM"
	Mumbai   = "Mumbai"
	Arbitrum = "Arbitrum"
)

var (
	Profile = Env("ENV")

	PROJECT_NAME = "fetcher"
)

func EnvMongoURI() string {
	if Profile == "prod" {
		return Env("MONGOURI_PROD")
	}
	return Env("MONGOURI_DEV")
}

func AlchemyURI(chain string) string {
	if chain == Mumbai || chain == POLYGON {
		return polygonURI(chain)
	} else if chain == Arbitrum {
		return Env("ALCHEMY_Arbitrum_URI")
	} else {
		return ethereumURI(chain)
	}
}

func PublicURI(chain string) string {
	if chain == ETHEREUM {
		return Env("PUBLIC_ETHEREUM_URI")
	}
	if chain == POLYGON {
		return Env("PUBLIC_POLYGON_URI")
	}
	if chain == Arbitrum {
		return Env("PUBLIC_Arbitrum_URI")
	}
	return AlchemyURI(chain) + "/v2/" + AlchemyKey(chain)
}

func polygonURI(chain string) string {
	if chain == POLYGON {
		return Env("ALCHEMY_POLYGON_URI")
	}
	return Env("ALCHEMY_MUMBAI_URI")
}

func ethereumURI(chain string) string {
	if chain == ETHEREUM {
		return Env("ALCHEMY_ETHEREUM_URI")
	}
	return Env("ALCHEMY_GOERLI_URI")
}

func AlchemyWS(chain string) string {
	if chain == Mumbai || chain == POLYGON {
		return polygonWS(chain)
	} else if chain == Arbitrum {
		return Env("ALCHEMY_Arbitrum_WS")
	} else {
		return ethereumWS(chain)
	}
}

func polygonWS(chain string) string {
	if chain == POLYGON {
		return Env("ALCHEMY_POLYGON_WS")
	}
	return Env("ALCHEMY_MUMBAI_WS")
}

func ethereumWS(chain string) string {
	if chain == ETHEREUM {
		return Env("ALCHEMY_ETHEREUM_WS")
	}
	return Env("ALCHEMY_GOERLI_WS")
}

func AlchemyKey(chain string) string {
	if chain == Mumbai || chain == POLYGON {
		return polygonKey(chain)
	} else {
		return ethereumKey(chain)
	}
}

func polygonKey(chain string) string {
	if chain == POLYGON {
		return Env("ALCHEMY_POLYGON_API_KEY")
	}
	return Env("ALCHEMY_MUMBAI_API_KEY")
}

func ethereumKey(chain string) string {
	if chain == ETHEREUM {
		return Env("ALCHEMY_ETHEREUM_API_KEY")
	}
	return Env("ALCHEMY_GOERLI_API_KEY")
}

func MongoDB() string {
	if Profile == "prod" {
		return Env("MONGODB_PROD")
	}
	return Env("MONGODB_DEV")
}

func UnstoppableURI() string {
	return Env("UNSTOPPABLE_URI")
}

func MaxBlock(chain string) uint64 {
	if chain == Mumbai || chain == POLYGON {
		return polygonMaxBlock()
	} else {
		return ethereumMaxBlock()
	}
}

func polygonMaxBlock() uint64 {
	max, err := strconv.ParseUint(Env("POLYGON_MAX_BLOCK"), 10, 64)
	if err != nil {
		log.Fatalf("Error ParseUint POLYGON_MAX_BLOCK")
	}
	return max
}

func ethereumMaxBlock() uint64 {
	max, err := strconv.ParseUint(Env("ETHEREUM_MAX_BLOCK"), 10, 64)
	if err != nil {
		log.Fatalf("Error ParseUint ETHEREUM_MAX_BLOCK")
	}
	return max
}

func Env(key string) string {
	path, _ := os.Getwd()
	err := godotenv.Load(strings.Split(path, PROJECT_NAME)[0] + PROJECT_NAME + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func DotbitApiURI() string {
	return Env("DOTBIT_API_URI")
}

func Network() string {
	return Env("NETWORK")
}

func Blockchain() string {
	return Env("BLOCKCHAIN")
}

func NeedRefresh() bool {
	b, err := strconv.ParseBool(Env("REFRESH"))
	if err != nil {
		return false
	}
	return b
}
