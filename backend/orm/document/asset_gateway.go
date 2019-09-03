package document

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmAssetGatways = "ex_asset_gateways"
)

type AssetGateways struct {
	ID       bson.ObjectId `bson:"_id"`
	Owner    string        `bson:"owner" json:"owner"`
	Moniker  string        `bson:"moniker" json:"moniker"`
	Identity string        `bson:"identity" json:"identity"`
	Details  string        `bson:"details" json:"details"`
	Website  string        `bson:"website" json:"website"`
}

func (_ AssetGateways) GetAllAssetGateways() ([]AssetGateways, error) {
	var assets []AssetGateways
	var qeury = orm.NewQuery()
	defer qeury.Release()
	qeury.SetCollection(CollectionNmAssetGatways).
		SetResult(&assets)
	err := qeury.Exec()

	return assets, err
}

func (_ AssetGateways) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
