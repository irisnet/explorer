package service

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"gopkg.in/mgo.v2/txn"
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
)

func (service *AssetsService) UpdateAssetGateway(vs []document.AssetGateways) error {

	var vMap = make(map[string]document.AssetGateways)
	for _, v := range vs {
		vMap[v.Moniker] = v
	}

	dstAssetTokens := buildAssetGateway()
	var txs []txn.Op
	for _, v := range dstAssetTokens {
		if it, ok := vMap[v.Moniker]; ok {
			if isDiffAssetGateway(it, v) {
				v.ID = it.ID
				txs = append(txs, txn.Op{
					C:  document.CollectionNmAssetGatways,
					Id: v.ID,
					Update: bson.M{
						"$set": v,
					},
				})
			}
			delete(vMap, v.Moniker)
		} else {
			v.ID = bson.NewObjectId()
			txs = append(txs, txn.Op{
				C:      document.CollectionNmAssetGatways,
				Id:     bson.NewObjectId(),
				Insert: v,
			})
		}
	}

	if len(vMap) > 0 {
		for symbol := range vMap {
			v := vMap[symbol]
			txs = append(txs, txn.Op{
				C:      document.CollectionNmAssetGatways,
				Id:     v.ID,
				Remove: true,
			})
		}
	}
	return document.AssetGateways{}.Batch(txs)

}

func buildAssetGateway() []document.AssetGateways {
	res := lcd.GetAssetGateways()
	var result []document.AssetGateways
	var buildAssetGateways = func(v lcd.AssetGateways) (document.AssetGateways, error) {
		var assetGateways document.AssetGateways
		if err := utils.Copy(v, &assetGateways); err != nil {
			logger.Error("utils.copy assetGateways failed")
			return assetGateways, err
		}
		return assetGateways, nil
	}

	for _, v := range res {
		var genAssetTokens = func(va lcd.AssetGateways, result *[]document.AssetGateways) {
			assetGateways, err := buildAssetGateways(va)
			if err != nil {
				logger.Error("utils.copy assetGateways failed")
				panic(err)
			}
			*result = append(*result, assetGateways)
		}
		genAssetTokens(v, &result)
	}

	return result
}

func isDiffAssetGateway(src, dst document.AssetGateways) bool {
	if src.Website != dst.Website ||
		src.Details != dst.Details ||
		src.Moniker != dst.Moniker ||
		src.Identity != dst.Identity ||
		src.Owner != dst.Owner {
		return true
	}
	return false
}
