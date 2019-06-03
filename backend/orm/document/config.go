package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmConfig = "ex_config"

	Account_Field_EnvNm      = "env_nm"
	Account_Field_Host       = "host"
	Account_Field_ChainId    = "chain_id"
	Account_Field_ShowFaucet = "show_faucet"
)

type Config struct {
	EnvNm      string `bson:"env_nm" json:"env_nm",omitempty`
	Host       string `bson:"host" json:"host,omitempty"`
	ChainId    string `bson:"chain_id" json:"chain_id,omitempty"`
	ShowFaucet int    `bson:"show_faucet" json:"show_faucet,omitempty"`
}

func (a Config) Name() string {
	return CollectionNmConfig
}

func (a Config) PkKvPair() map[string]interface{} {
	return bson.M{Account_Field_EnvNm: a.EnvNm}
}

func (_ Config) GetConfig() ([]Config, error) {

	var configs []Config
	var query = orm.NewQuery().
		SetCollection(CollectionNmConfig).
		SetResult(&configs)
	err := query.Exec()
	return configs, err
}
