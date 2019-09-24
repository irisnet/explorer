package document

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmAsset = "ex_asset_tokens"
	AssetFieldTokenId = "token_id"
	AssetFieldSource  = "source"
	AssetFieldFamily  = "family"
	FungibleFamily    = "fungible"
)

type AssetToken struct {
	ID              bson.ObjectId `bson:"_id"`
	TokenId         string        `bson:"token_id"`
	Family          string        `bson:"family"`
	Source          string        `bson:"source"`
	Gateway         string        `bson:"gateway"`
	Symbol          string        `bson:"symbol"`
	Name            string        `bson:"name"`
	Decimal         int           `bson:"decimal"`
	CanonicalSymbol string        `bson:"canonical_symbol"`
	MinUnitAlias    string        `bson:"min_unit_alias"`
	InitialSupply   string        `bson:"initial_supply"`
	MaxSupply       string        `bson:"max_supply"`
	TotalSupply     string        `bson:"total_supply"`
	Mintable        bool          `bson:"mintable"`
	Owner           string        `bson:"owner"`
}

func (_ AssetToken) GetAllAssets() ([]AssetToken, error) {
	var assets []AssetToken
	var qeury = orm.NewQuery()
	defer qeury.Release()
	qeury.SetCollection(CollectionNmAsset).
		SetResult(&assets)
	err := qeury.Exec()

	return assets, err
}

func (_ AssetToken) GetAssetTokens(source string) ([]AssetToken, error) {
	var assets []AssetToken
	cond := bson.M{}
	if source != "" {
		cond[AssetFieldSource] = source
	}
	cond[AssetFieldFamily] = FungibleFamily
	err := queryAll(CollectionNmAsset, nil, cond, "", 0, &assets)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return assets, err
	}
	return assets, nil
}

func (_ AssetToken) GetAssetTokenDetail(tokenid string) (AssetToken, error) {
	var asset AssetToken
	cond := bson.M{}
	if tokenid != "" {
		cond[AssetFieldTokenId] = tokenid
	}
	err := queryOne(CollectionNmAsset, nil, cond, &asset)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return asset, err
	}
	return asset, nil
}

func (_ AssetToken) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
