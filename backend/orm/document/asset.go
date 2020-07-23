package document

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmAsset = "ex_asset_tokens"
	//AssetFieldTokenId = "token_id"
	AssetFieldSymbol = "symbol"
	AssetFieldSource = "source"
	AssetFieldFamily = "family"
	FungibleFamily   = "fungible"
)

type AssetToken struct {
	ID bson.ObjectId `bson:"_id"`
	//TokenId         string        `bson:"token_id" json:"id"`
	//Family          string        `bson:"family" json:"family"`
	//Source          string        `bson:"source" json:"source"`
	//Gateway         string        `bson:"gateway" json:"gateway"`
	Symbol string `bson:"symbol" json:"symbol"`
	Name   string `bson:"name" json:"name"`
	//Decimal         int           `bson:"decimal" json:"decimal"`
	//CanonicalSymbol string        `bson:"canonical_symbol" json:"canonical_symbol"`
	MinUnitAlias  string `bson:"min_unit_alias" json:"min_unit_alias"`
	InitialSupply string `bson:"initial_supply" json:"initial_supply"`
	MaxSupply     string `bson:"max_supply" json:"max_supply"`
	TotalSupply   string `bson:"total_supply" json:"total_supply"`
	Mintable      bool   `bson:"mintable" json:"mintable"`
	Owner         string `bson:"owner" json:"owner"`
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

func (_ AssetToken) GetAssetTokenDetail(symbol string) (AssetToken, error) {
	var asset AssetToken
	cond := bson.M{}
	if symbol != "" {
		cond[AssetFieldSymbol] = symbol
	}
	err := queryOne(CollectionNmAsset, nil, cond, &asset)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return asset, err
	}
	return asset, nil
}

func (_ AssetToken) GetAssetTokenDetailBySymbols(symbol []string) ([]AssetToken, error) {
	var asset []AssetToken
	cond := bson.M{}
	if len(symbol) > 0 {
		cond[AssetFieldSymbol] = bson.M{"$in": symbol}
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
