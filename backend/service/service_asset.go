package service

import (
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"sync"
)

type AssetService struct {
	BaseService
}

func (service *AssetService) GetModule() Module {
	return Validator
}

func (service *AssetService) UpdateAssetTokens(vs []document.Asset) error {
	var vMap = make(map[string]document.Asset)
	for _, v := range vs {
		vMap[v.Symbol] = v
	}

	dstAssetTokens := buildAssetTokens()
	var txs []txn.Op
	for _, v := range dstAssetTokens {
		if it, ok := vMap[v.Symbol]; ok {
			if isDiffAssetToken(it, v) {
				v.ID = it.ID
				txs = append(txs, txn.Op{
					C:  document.CollectionNmAsset,
					Id: it.ID,
					Update: bson.M{
						"$set": v,
					},
				})
			}
			delete(vMap, v.Symbol)
		} else {
			v.ID = bson.NewObjectId()
			txs = append(txs, txn.Op{
				C:      document.CollectionNmAsset,
				Id:     bson.NewObjectId(),
				Insert: v,
			})
		}
	}

	if len(vMap) > 0 {
		for symbol := range vMap {
			v := vMap[symbol]
			txs = append(txs, txn.Op{
				C:      document.CollectionNmAsset,
				Id:     v.ID,
				Remove: true,
			})
		}
	}
	return document.Asset{}.Batch(txs)
}

func buildAssetTokens() []document.Asset {
	res := lcd.GetAssetTokens()
	var result []document.Asset
	var buildAssetToken = func(v lcd.AssetTokens) (document.Asset, error) {
		var assetToken document.Asset
		if err := utils.Copy(v.BaseToken, &assetToken); err != nil {
			logger.Error("utils.copy assetToken failed")
			return assetToken, err
		}
		return assetToken, nil
	}

	var group sync.WaitGroup
	var mutex sync.Mutex
	group.Add(len(res))
	for _, v := range res {
		var genAssetTokens = func(va lcd.AssetTokens, result *[]document.Asset) {
			defer group.Done()
			assetToken, err := buildAssetToken(va)
			if err != nil {
				logger.Error("utils.copy assetToken failed")
				panic(err)
			}
			mutex.Lock()
			*result = append(*result, assetToken)
			mutex.Unlock()
		}
		go genAssetTokens(v, &result)
	}
	group.Wait()
	return result
}

func isDiffAssetToken(src, dst document.Asset) bool {
	if src.Id != dst.Id ||
		src.Family != dst.Family ||
		src.Source != dst.Source ||
		src.Gateway != dst.Gateway ||
		src.Symbol != dst.Symbol ||
		src.Name != dst.Name ||
		src.Decimal != dst.Decimal ||
		src.CanonicalSymbol != dst.CanonicalSymbol ||
		src.MinUnitAlias != dst.MinUnitAlias ||
		src.InitialSupply != dst.InitialSupply ||
		src.MaxSupply != dst.MaxSupply ||
		src.Mintable != dst.Mintable ||
		src.Owner != dst.Owner {
		return true
	}
	return false
}

func (service *AssetService) QueryAssetToken() []model.AssetTokens {
	res, err := document.Asset{}.GetAllAssets()
	if err != nil {
		logger.Error("GetAllAssets", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	assetTokens := make([]model.AssetTokens, 0, len(res))

	for _, v := range res {
		tmp := model.AssetTokens{
			Symbol:  v.Symbol,
			Decimal: v.Decimal,
		}
		assetTokens = append(assetTokens, tmp)
	}
	if len(assetTokens) > 1 {
		return assetTokens
	}
	return []model.AssetTokens{}
}
