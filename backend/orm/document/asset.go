package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/logger"
)

const (
	CollectionNmAsset      = "ex_asset_tokens"
	AssetFieldOwnerAddress = "owner"
	AssetFieldSource       = "source"
	AssetFieldFamily       = "family"
	FungibleFamily         = "fungible"
)

type Asset struct {
	ID              bson.ObjectId `bson:"_id"`
	TokenId         string        `bson:"token_id" json:"id"`
	Family          string        `bson:"family" json:"family"`
	Source          string        `bson:"source" json:"source"`
	Gateway         string        `bson:"gateway" json:"gateway"`
	Symbol          string        `bson:"symbol" json:"symbol"`
	Name            string        `bson:"name" json:"name"`
	Decimal         int           `bson:"decimal" json:"decimal"`
	CanonicalSymbol string        `bson:"canonical_symbol" json:"canonical_symbol"`
	MinUnitAlias    string        `bson:"min_unit_alias" json:"min_unit_alias"`
	InitialSupply   string        `bson:"initial_supply" json:"initial_supply"`
	MaxSupply       string        `bson:"max_supply" json:"max_supply"`
	TotalSupply     string        `bson:"total_supply" json:"total_supply"`
	Mintable        bool          `bson:"mintable" json:"mintable"`
	Owner           string        `bson:"owner" json:"owner"`
}

func (_ Asset) GetAllAssets() ([]Asset, error) {
	var assets []Asset
	var qeury = orm.NewQuery()
	defer qeury.Release()
	qeury.SetCollection(CollectionNmAsset).
		SetResult(&assets)
	err := qeury.Exec()

	return assets, err
}

func (_ Asset) GetAssetByAddr(address, assettype string) ([]Asset, error) {
	var assets []Asset
	cond := bson.M{AssetFieldOwnerAddress: address, AssetFieldSource: assettype}
	cond[AssetFieldFamily] = FungibleFamily
	err := queryOne(CollectionNmAsset, nil, cond, &assets)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return assets, err
	}
	return assets, nil
}

func (_ Asset) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
