package conf

import (
	"github.com/irisnet/explorer/backend/logger"
	"os"
	"strconv"
	"strings"
)

const (
	KeyDbAddr      = "DB_ADDR"
	KeyDATABASE    = "DB_DATABASE"
	KeyDbUser      = "DB_USER"
	KeyDbPwd       = "DB_PASSWORD"
	KeyDbPoolLimit = "DB_POOL_LIMIT"

	KeyServerPort    = "PORT"
	KeyAddrHubLcd    = "ADDR_NODE_SERVER"
	KeyAddrHubNode   = "ADDR_HUB_RPC"
	KeyAddrFaucet    = "FAUCET_URL"
	KeyChainId       = "CHAIN_ID"
	KeyApiVersion    = "API_VERSION"
	KeyMaxDrawCnt    = "MAX_DRAW_CNT"
	KeyShowFaucet    = "SHOW_FAUCET"
	KeyCurEnv        = "CUR_ENV"
	KeyCosmosLcd     = "COSMOS_LCD"
	KeyCosmosValAddr = "COSMOS_VAL_ADDR"
	KeyCosmosValPub  = "COSMOS_VAL_PUB"

	KeyPrefixAccAddr  = "PrefixAccAddr"
	KeyPrefixAccPub   = "PrefixAccPub"
	KeyPrefixValAddr  = "PrefixValAddr"
	KeyPrefixValPub   = "PrefixValPub"
	KeyPrefixConsAddr = "PrefixConsAddr"
	KeyPrefixConsPub  = "PrefixConsPub"

	EnvironmentDevelop = "develop"
	EnvironmentLocal   = "local"
	EnvironmentQa      = "qa"
	EnvironmentStage   = "stage"
	EnvironmentProd    = "prod"

	DefaultEnvironment = EnvironmentDevelop
)

var (
	config        Config
	defaultConfig = map[string]map[string]string{}
)

func init() {
	logger.Info("==================================load config start==================================")
	loadDefault()

	addrs := strings.Split(getEnv(KeyDbAddr, DefaultEnvironment), ",")
	db := dbConf{
		Addrs:     addrs,
		Database:  getEnv(KeyDATABASE, DefaultEnvironment),
		UserName:  getEnv(KeyDbUser, DefaultEnvironment),
		Password:  getEnv(KeyDbPwd, DefaultEnvironment),
		PoolLimit: getEnvInt(KeyDbPoolLimit, DefaultEnvironment),
	}
	config.Db = db

	server := serverConf{
		ServerPort: getEnvInt(KeyServerPort, DefaultEnvironment),
		FaucetUrl:  getEnv(KeyAddrFaucet, DefaultEnvironment),
		ApiVersion: getEnv(KeyApiVersion, DefaultEnvironment),
		MaxDrawCnt: getEnvInt(KeyMaxDrawCnt, DefaultEnvironment),
		ShowFaucet: getEnv(KeyShowFaucet, DefaultEnvironment),
		CurEnv:     getEnv(KeyCurEnv, DefaultEnvironment),
	}
	config.Server = server

	hubcf := hubConf{
		Prefix: bech32Prefix{
			AccAddr:  getEnv(KeyPrefixAccAddr, DefaultEnvironment),
			AccPub:   getEnv(KeyPrefixAccPub, DefaultEnvironment),
			ValAddr:  getEnv(KeyPrefixValAddr, DefaultEnvironment),
			ValPub:   getEnv(KeyPrefixValPub, DefaultEnvironment),
			ConsAddr: getEnv(KeyPrefixConsAddr, DefaultEnvironment),
			ConsPub:  getEnv(KeyPrefixConsPub, DefaultEnvironment),
		},
		LcdUrl:        getEnv(KeyAddrHubLcd, DefaultEnvironment),
		NodeUrl:       getEnv(KeyAddrHubNode, DefaultEnvironment),
		ChainId:       getEnv(KeyChainId, DefaultEnvironment),
		CosmosLcd:     getEnv(KeyCosmosLcd, DefaultEnvironment),
		CosmosValAddr: getEnv(KeyCosmosValAddr, DefaultEnvironment),
		CosmosValPub:  getEnv(KeyCosmosValPub, DefaultEnvironment),
	}
	config.Hub = hubcf

	logger.Info("==================================load config end==================================")
}

func loadDefault() {
	defaultConfig[EnvironmentDevelop] = map[string]string{
		KeyDbAddr:         "192.168.150.31:27017",
		KeyDATABASE:       "sync-iris",
		KeyDbUser:         "iris",
		KeyDbPwd:          "irispassword",
		KeyDbPoolLimit:    "4096",
		KeyServerPort:     "8080",
		KeyAddrHubLcd:     "http://irisnet-lcd.dev.rainbow.one",
		KeyAddrHubNode:    "http://192.168.150.31:30657",
		KeyAddrFaucet:     "http://192.168.150.31:4000",
		KeyChainId:        "rainbow-dev",
		KeyApiVersion:     "v0.6.5",
		KeyMaxDrawCnt:     "10",
		KeyPrefixAccAddr:  "faa",
		KeyPrefixAccPub:   "fap",
		KeyPrefixValAddr:  "fva",
		KeyPrefixValPub:   "fvp",
		KeyPrefixConsAddr: "fca",
		KeyPrefixConsPub:  "fcp",
		KeyShowFaucet:     "1",
		KeyCurEnv:         "dev",
		KeyCosmosLcd:      "http://35.220.142.249:2317",
		KeyCosmosValAddr:  "cosmosvaloper1ssm0d433seakyak8kcf93yefhknjleeds4y3em",
		KeyCosmosValPub:   "cosmosvalconspub1zcjduepqrgyyjxpe0ujefxwnkpmqz9m0hj03y09tdz9lwc0s7mvy469hulfq69f8sd",
	}

	defaultConfig[EnvironmentLocal] = map[string]string{
		KeyDbAddr:         "127.0.0.1:27017",
		KeyDATABASE:       "sync-iris",
		KeyDbUser:         "iris",
		KeyDbPwd:          "irispassword",
		KeyDbPoolLimit:    "4096",
		KeyServerPort:     "8080",
		KeyAddrHubLcd:     "http://127.0.0.1:1317",
		KeyAddrHubNode:    "http://127.0.0.1:26657",
		KeyAddrFaucet:     "http://192.168.150.7:30200",
		KeyChainId:        "rainbow-dev",
		KeyApiVersion:     "v0.6.5",
		KeyMaxDrawCnt:     "10",
		KeyPrefixAccAddr:  "faa",
		KeyPrefixAccPub:   "fap",
		KeyPrefixValAddr:  "fva",
		KeyPrefixValPub:   "fvp",
		KeyPrefixConsAddr: "fca",
		KeyPrefixConsPub:  "fcp",
		KeyShowFaucet:     "1",
		KeyCurEnv:         "dev",
	}
}

func Get() Config {
	return config
}

type Config struct {
	Db     dbConf
	Server serverConf
	Hub    hubConf
}

type dbConf struct {
	Addrs     []string
	Database  string
	UserName  string
	Password  string
	PoolLimit int
}

type serverConf struct {
	ServerPort int
	FaucetUrl  string
	ApiVersion string
	MaxDrawCnt int
	ShowFaucet string
	CurEnv     string
}

type hubConf struct {
	Prefix        bech32Prefix
	LcdUrl        string
	NodeUrl       string
	ChainId       string
	CosmosLcd     string
	CosmosValAddr string
	CosmosValPub  string
}

type bech32Prefix struct {
	AccAddr  string
	AccPub   string
	ValAddr  string
	ValPub   string
	ConsAddr string
	ConsPub  string
}

func getEnv(key string, environment string) string {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if DefaultEnvironment == EnvironmentStage || DefaultEnvironment == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", DefaultEnvironment))
		}
		value = defaultConfig[environment][key]
	}
	if value == "" {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	logger.Info("config", logger.String(key, value))
	return value
}

func getEnvInt(key string, environment string) int {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if DefaultEnvironment == EnvironmentStage || DefaultEnvironment == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", DefaultEnvironment))
		}
		value = defaultConfig[environment][key]
	}

	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	logger.Info("config", logger.Int64(key, i))
	return int(i)
}
