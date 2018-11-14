package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	KeyDbUrl       = "DB_URL"
	KeyDATABASE    = "DB_DATABASE"
	KeyDbUser      = "DB_USER"
	KeyDbPwd       = "DB_PASSWORD"
	KeyDbPoolLimit = "DB_POOL_LIMIT"

	KeyServerPort  = "PORT"
	KeyAddrHubLcd  = "ADDR_NODE_SERVER"
	KeyAddrHubNode = "ADDR_HUB_RPC"
	KeyAddrFaucet  = "FAUCET_URL"
	KeyChainId     = "CHAIN_ID"

	EnvironmentDevelop = "develop"
	EnvironmentLocal   = "local"

	DefaultEnvironment = EnvironmentDevelop
)

var (
	config        Config
	defaultConfig = map[string]map[string]string{}
)

func init() {
	log.Printf("==================================load config start==================================")
	loadDefault()

	db := dbConf{
		Url:       getEnv(KeyDbUrl, DefaultEnvironment),
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
	}

	config.Server = server
	log.Printf("==================================load config end==================================")
}

func loadDefault() {
	defaultConfig[EnvironmentDevelop] = map[string]string{
		KeyDbUrl:       "192.168.150.7:30000",
		KeyDATABASE:    "sync-iris",
		KeyDbUser:      "iris",
		KeyDbPwd:       "irispassword",
		KeyDbPoolLimit: "4096",
		KeyServerPort:  "8080",
		KeyAddrHubLcd:  "http://192.168.150.7:1317",
		KeyAddrHubNode: "http://192.168.150.7:26657",
		KeyAddrFaucet:  "http://dev.faucet.irisplorer.io",
		KeyChainId:     "rainbow-dev",
	}

	defaultConfig[EnvironmentLocal] = map[string]string{
		KeyDbUrl:       "127.0.0.1:27017",
		KeyDATABASE:    "sync-iris",
		KeyDbUser:      "iris",
		KeyDbPwd:       "irispassword",
		KeyDbPoolLimit: "4096",
		KeyServerPort:  "8080",
		KeyAddrHubLcd:  "http://127.0.0.1:1317",
		KeyAddrHubNode: "http://127.0.0.1:26657",
		KeyAddrFaucet:  "http://dev.faucet.irisplorer.io",
		KeyChainId:     "rainbow-dev",
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
	Url       string
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
}

func getEnv(key string, environment string) string {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		value = defaultConfig[environment][key]
	}
	if value == "" {
		panic(fmt.Sprintf("key [%s] must be not empty", key))
	}
	log.Printf(fmt.Sprintf("%s=%s", key, value))
	return value
}

func getEnvInt(key string, environment string) int {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		value = defaultConfig[environment][key]
	}

	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		panic(fmt.Sprintf("key [%s] must be int ", key))
	}
	log.Printf(fmt.Sprintf("%s=%d", key, i))
	return int(i)
}
