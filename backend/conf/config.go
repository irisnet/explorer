package conf

import (
	"github.com/irisnet/irishub-sync/logger"
	"os"
	"strconv"
)

const (
	KeyDbAddr      = "DB_ADDR"
	KeyDATABASE    = "DB_DATABASE"
	KeyDbUser      = "DB_USER"
	KeyDbPwd       = "DB_PASSWORD"
	KeyDbPoolLimit = "DB_POOL_LIMIT"

	KeyServerPort  = "PORT"
	KeyAddrHubLcd  = "ADDR_NODE_SERVER"
	KeyAddrHubNode = "ADDR_HUB_RPC"
	KeyAddrFaucet  = "FAUCET_URL"
	KeyChainId     = "CHAIN_ID"
	KeyApiVersion  = "API_VERSION"
	KeyMaxDrawCnt  = "MAX_DRAW_CNT"
	KeyShowFaucet = "SHOW_FAUCET"

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

	db := dbConf{
		Addr:      getEnv(KeyDbAddr, DefaultEnvironment),
		Database:  getEnv(KeyDATABASE, DefaultEnvironment),
		UserName:  getEnv(KeyDbUser, DefaultEnvironment),
		Password:  getEnv(KeyDbPwd, DefaultEnvironment),
		PoolLimit: getEnvInt(KeyDbPoolLimit, DefaultEnvironment),
	}

	config.Db = db

	server := serverConf{
		ServerPort:   getEnvInt(KeyServerPort, DefaultEnvironment),
		HubLcdUrl:    getEnv(KeyAddrHubLcd, DefaultEnvironment),
		HubNodeUrl:   getEnv(KeyAddrHubNode, DefaultEnvironment),
		HubFaucetUrl: getEnv(KeyAddrFaucet, DefaultEnvironment),
		ChainId:      getEnv(KeyChainId, DefaultEnvironment),
		ApiVersion:   getEnv(KeyApiVersion, DefaultEnvironment),
		MaxDrawCnt:   getEnvInt(KeyMaxDrawCnt, DefaultEnvironment),
		ShowFaucet:   getEnvInt(KeyShowFaucet,"1"),
	}

	config.Server = server
	logger.Info("==================================load config end==================================")
}

func loadDefault() {
	defaultConfig[EnvironmentDevelop] = map[string]string{
		KeyDbAddr:      "192.168.150.7:30000",
		KeyDATABASE:    "sync-iris",
		KeyDbUser:      "iris",
		KeyDbPwd:       "irispassword",
		KeyDbPoolLimit: "4096",
		KeyServerPort:  "8080",
		KeyAddrHubLcd:  "http://192.168.150.7:30317",
		KeyAddrHubNode: "http://192.168.150.7:26657",
		KeyAddrFaucet:  "http://192.168.150.7:30200",
		KeyChainId:     "rainbow-dev",
		KeyApiVersion:  "v0.6.5",
		KeyMaxDrawCnt:  "10",
	}

	defaultConfig[EnvironmentLocal] = map[string]string{
		KeyDbAddr:      "127.0.0.1:27017",
		KeyDATABASE:    "sync-iris",
		KeyDbUser:      "iris",
		KeyDbPwd:       "irispassword",
		KeyDbPoolLimit: "4096",
		KeyServerPort:  "8080",
		KeyAddrHubLcd:  "http://127.0.0.1:1317",
		KeyAddrHubNode: "http://127.0.0.1:26657",
		KeyAddrFaucet:  "http://192.168.150.7:30200",
		KeyChainId:     "rainbow-dev",
		KeyApiVersion:  "v0.6.5",
		KeyMaxDrawCnt:  "10",
	}
}

func Get() Config {
	return config
}

type Config struct {
	Db     dbConf
	Server serverConf
}

type dbConf struct {
	Addr      string
	Database  string
	UserName  string
	Password  string
	PoolLimit int
}

type serverConf struct {
	ServerPort   int
	HubLcdUrl    string
	HubNodeUrl   string
	HubFaucetUrl string
	ChainId      string
	ApiVersion   string
	MaxDrawCnt   int
	ShowFaucet   int
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
