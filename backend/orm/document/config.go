package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const (
	CollectionNmConfig = "ex_config"

	AccountFieldNetworkName  = "network_name"
	Account_Field_Host       = "host"
	Account_Field_ChainId    = "chain_id"
	Account_Field_ShowFaucet = "show_faucet"
)

type Config struct {
	NetworkName string `bson:"network_name"`
	Env         string `bson:"env"`
	Host        string `bson:"host"`
	ChainId     string `bson:"chain_id"`
	ShowFaucet  int    `bson:"show_faucet"`
	EnvLcd      string `bson:"env_lcd"`
	UmengId     int64  `bson:"umeng_id"`
}

func (c Config) String() string {
	return fmt.Sprintf(`
		NetworkName      :%v
		Env      :%v
		Host       :%v
		ChainId    :%v
		ShowFaucet :%v
		EnvLcd :%v
		UmengId :%v
 		`, c.NetworkName, c.Env, c.Host, c.ChainId, c.ShowFaucet, c.EnvLcd, c.UmengId)
}

func (a Config) Name() string {
	return CollectionNmConfig
}

func (a Config) PkKvPair() map[string]interface{} {
	return bson.M{AccountFieldNetworkName: a.NetworkName}
}

func (_ Config) GetConfig() ([]Config, error) {

	var configs []Config
	var query = orm.NewQuery().
		SetCollection(CollectionNmConfig).
		SetSort("+env").
		SetResult(&configs)
	err := query.Exec()
	return configs, err
}


func (a Config) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{"-network_name"},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}
