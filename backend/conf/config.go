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

	KeyServerPort  = "PORT"
	KeyAddrHubLcd  = "ADDR_NODE_SERVER"
	KeyAddrHubNode = "ADDR_HUB_RPC"
	KeyAddrFaucet  = "FAUCET_URL"
	KeyApiVersion  = "API_VERSION"
	KeyMaxDrawCnt  = "MAX_DRAW_CNT"
	KeyShowFaucet  = "SHOW_FAUCET"

	KeyPrefixAccAddr  = "PrefixAccAddr"
	KeyPrefixAccPub   = "PrefixAccPub"
	KeyPrefixValAddr  = "PrefixValAddr"
	KeyPrefixValPub   = "PrefixValPub"
	KeyPrefixConsAddr = "PrefixConsAddr"
	KeyPrefixConsPub  = "PrefixConsPub"

	KeyCurEnv = "CUR_ENV"

	EnvironmentDevelop = "dev"
	EnvironmentLocal   = "local"
	EnvironmentQa      = "qa"
	EnvironmentStage   = "stage"
	EnvironmentProd    = "prod"
)

var (
	config        Config
	defaultConfig = map[string]map[string]string{}
)

func init() {
	logger.Info("==================================load config start==================================")
	loadDefault()

	env := getEnv(KeyCurEnv, EnvironmentDevelop)

	addrs := strings.Split(getEnv(KeyDbAddr, env), ",")
	db := dbConf{
		Addrs:     addrs,
		Database:  getEnv(KeyDATABASE, env),
		UserName:  getEnv(KeyDbUser, env),
		Password:  getEnv(KeyDbPwd, env),
		PoolLimit: getEnvInt(KeyDbPoolLimit, env),
	}
	config.Db = db

	server := serverConf{
		ServerPort: getEnvInt(KeyServerPort, env),
		FaucetUrl:  getEnv(KeyAddrFaucet, env),
		ApiVersion: getEnv(KeyApiVersion, env),
		MaxDrawCnt: getEnvInt(KeyMaxDrawCnt, env),
		CurEnv:     getEnv(KeyCurEnv, env),
	}
	config.Server = server

	hubcf := hubConf{
		Prefix: bech32Prefix{
			AccAddr:  getEnv(KeyPrefixAccAddr, env),
			AccPub:   getEnv(KeyPrefixAccPub, env),
			ValAddr:  getEnv(KeyPrefixValAddr, env),
			ValPub:   getEnv(KeyPrefixValPub, env),
			ConsAddr: getEnv(KeyPrefixConsAddr, env),
			ConsPub:  getEnv(KeyPrefixConsPub, env),
		},
		LcdUrl:  getEnv(KeyAddrHubLcd, env),
		NodeUrl: getEnv(KeyAddrHubNode, env),
	}
	config.Hub = hubcf

	logger.Info("==================================load config end==================================")
}

func loadDefault() {
	defaultConfig[EnvironmentDevelop] = map[string]string{
		KeyDbAddr:         "192.168.150.7:30000",
		KeyDATABASE:       "sync-iris",
		KeyDbUser:         "iris",
		KeyDbPwd:          "irispassword",
		KeyDbPoolLimit:    "4096",
		KeyServerPort:     "8080",
		KeyAddrHubLcd:     "http://192.168.150.7:30317",
		KeyAddrHubNode:    "http://192.168.150.7:30657",
		KeyAddrFaucet:     "http://192.168.150.7:30200",
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
		KeyApiVersion:     "v0.6.5",
		KeyMaxDrawCnt:     "10",
		KeyPrefixAccAddr:  "faa",
		KeyPrefixAccPub:   "fap",
		KeyPrefixValAddr:  "fva",
		KeyPrefixValPub:   "fvp",
		KeyPrefixConsAddr: "fca",
		KeyPrefixConsPub:  "fcp",
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
	CurEnv     string
}

type hubConf struct {
	Prefix  bech32Prefix
	LcdUrl  string
	NodeUrl string
}

type bech32Prefix struct {
	AccAddr  string
	AccPub   string
	ValAddr  string
	ValPub   string
	ConsAddr string
	ConsPub  string
}

func getEnv(key string, environment ...string) string {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if len(environment) == 0 {
			logger.Panic("config must be not empty", logger.String("key", key))
		}
		if environment[0] == EnvironmentStage || environment[0] == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", environment[0]))
		}
		value = defaultConfig[environment[0]][key]
	}
	if value == "" {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	logger.Info("config", logger.String(key, value))
	return value
}

func getEnvInt(key string, environment ...string) int {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if len(environment) == 0 {
			logger.Panic("config must be not empty", logger.String("key", key))
		}
		if environment[0] == EnvironmentStage || environment[0] == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", environment[0]))
		}
		value = defaultConfig[environment[0]][key]
	}

	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	logger.Info("config", logger.Int64(key, i))
	return int(i)
}
