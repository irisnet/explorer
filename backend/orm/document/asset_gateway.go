package document

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/logger"
)

const (
	CollectionNmAssetGatways = "ex_asset_gateways"
	GatewayFieldOwnerAddress = "owner"
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

func (_ AssetGateways) GetGatewayInfoByOwner(owneraddr string) (gateway []AssetGateways, err error) {
	cond := bson.M{}
	if owneraddr != "" {
		cond[GatewayFieldOwnerAddress] = owneraddr
	}
	err = queryAll(CollectionNmAssetGatways, nil, cond, "", 0, &gateway)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return
	}
	return
}

func (_ AssetGateways) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
