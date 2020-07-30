package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmAsset = "ex_asset_tokens"
	//AssetFieldTokenId = "token_id"
	AssetFieldSymbol = "symbol"
	//AssetFieldSource = "source"
	//AssetFieldFamily = "family"
	//FungibleFamily   = "fungible"
)

type AssetToken struct {
	ID bson.ObjectId `bson:"_id"`
	//TokenId         string        `bson:"token_id" json:"id"`
	//Family          string        `bson:"family" json:"family"`
	//Source          string        `bson:"source" json:"source"`
	//Gateway         string        `bson:"gateway" json:"gateway"`
	Symbol    string `bson:"symbol" json:"symbol"`
	AssetName string `bson:"name" json:"name"`
	Scale     int    `bson:"scale" json:"scale"`
	//CanonicalSymbol string        `bson:"canonical_symbol" json:"canonical_symbol"`
	MinUnitAlias  string `bson:"min_unit_alias" json:"min_unit_alias"`
	InitialSupply string `bson:"initial_supply" json:"initial_supply"`
	MaxSupply     string `bson:"max_supply" json:"max_supply"`
	//TotalSupply   string `bson:"total_supply" json:"total_supply"`
	Mintable bool   `bson:"mintable" json:"mintable"`
	Owner    string `bson:"owner" json:"owner"`
}

func (a AssetToken) Name() string {
	return CollectionNmAsset
}

func (a AssetToken) PkKvPair() map[string]interface{} {
	return bson.M{AssetFieldSymbol: a.Symbol}
}

func (a AssetToken) GetAllAssets() ([]AssetToken, error) {
	var assets []AssetToken
	var qeury = orm.NewQuery()
	defer qeury.Release()
	qeury.SetCollection(CollectionNmAsset).
		SetResult(&assets)
	err := qeury.Exec()

	return assets, err
}

func (a AssetToken) GetAssetTokens() ([]AssetToken, error) {
	var assets []AssetToken
	err := queryAll(CollectionNmAsset, nil, nil, "", 0, &assets)
	if err != nil {
		//logger.Error("validator not found", logger.Any("err", err.Error()))
		return assets, err
	}
	return assets, nil
}

func (a AssetToken) GetAssetTokenDetail(symbol string) (AssetToken, error) {
	var asset AssetToken
	cond := bson.M{}
	if symbol != "" {
		cond[AssetFieldSymbol] = symbol
	}
	err := queryOne(CollectionNmAsset, nil, cond, &asset)
	if err != nil {
		//logger.Error("validator not found", logger.Any("err", err.Error()))
		return asset, err
	}
	return asset, nil
}

func (a AssetToken) GetAssetTokenDetailBySymbols(symbol []string) ([]AssetToken, error) {
	var asset []AssetToken
	cond := bson.M{}
	if len(symbol) > 0 {
		cond[AssetFieldSymbol] = bson.M{"$in": symbol}
	}
	err := queryOne(CollectionNmAsset, nil, cond, &asset)
	if err != nil {
		//logger.Error("validator not found", logger.Any("err", err.Error()))
		return asset, err
	}
	return asset, nil
}

func (a AssetToken) Update(token AssetToken) error {
	q := orm.NewQuery()
	defer q.Release()

	c := q.GetDb().C(CollectionNmAsset)

	if err := c.Update(token.PkKvPair(), token); err != nil {
		return err
	}
	return nil
}

func (a AssetToken) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{AssetFieldSymbol},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (a AssetToken) Save(token AssetToken) error {
	q := orm.NewQuery()
	defer q.Release()

	c := q.GetDb().C(CollectionNmAsset)
	if err := c.Insert(token); err != nil {
		return err
	}
	return nil
}

func (a AssetToken) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
