package service

//func (service *AssetsService) UpdateAssetGateway(vs []document.AssetGateways) error {
//
//	//var vMap = make(map[string]document.AssetGateways)
//	//for _, v := range vs {
//	//	vMap[v.Moniker] = v
//	//}
//
//	//dstAssetTokens := buildAssetGateway()
//	var txs []txn.Op
//	//for _, v := range dstAssetTokens {
//	//	if it, ok := vMap[v.Moniker]; ok {
//	//		if isDiffAssetGateway(it, v) {
//	//			v.ID = it.ID
//	//			txs = append(txs, txn.Op{
//	//				C:  document.CollectionNmAssetGatways,
//	//				Id: v.ID,
//	//				Update: bson.M{
//	//					"$set": v,
//	//				},
//	//			})
//	//		}
//	//		delete(vMap, v.Moniker)
//	//	} else {
//	//		v.ID = bson.NewObjectId()
//	//		txs = append(txs, txn.Op{
//	//			C:      document.CollectionNmAssetGatways,
//	//			Id:     bson.NewObjectId(),
//	//			Insert: v,
//	//		})
//	//	}
//	//}
//	//
//	//if len(vMap) > 0 {
//	//	for symbol := range vMap {
//	//		v := vMap[symbol]
//	//		txs = append(txs, txn.Op{
//	//			C:      document.CollectionNmAssetGatways,
//	//			Id:     v.ID,
//	//			Remove: true,
//	//		})
//	//	}
//	//}
//	return document.AssetGateways{}.Batch(txs)
//
//}
//
//func buildAssetGateway() []document.AssetGateways {
//	res := lcd.GetAssetGateways()
//	var result []document.AssetGateways
//	var buildAssetGateways = func(v lcd.AssetGateways) (document.AssetGateways, error) {
//		var assetGateways document.AssetGateways
//		if err := utils.Copy(v, &assetGateways); err != nil {
//			logger.Error("utils.copy assetGateways failed")
//			return assetGateways, err
//		}
//		return assetGateways, nil
//	}
//
//	var genAssetTokens = func(va lcd.AssetGateways, result *[]document.AssetGateways) {
//		assetGateways, err := buildAssetGateways(va)
//		if err != nil {
//			logger.Error("utils.copy assetGateways failed")
//			panic(err)
//		}
//		//if identity := va.Identity; identity != "" {
//		//	urlicons, err := lcd.GetIconsByKey(identity)
//		//	if err != nil || len(urlicons) == 0 {
//		//		if err != nil {
//		//			logger.Error("GetIconsByKey have error", logger.String("error", err.Error()))
//		//		}
//		//	} else {
//		//		assetGateways.Icons = urlicons
//		//	}
//		//}
//		*result = append(*result, assetGateways)
//	}
//	for _, v := range res {
//		genAssetTokens(v, &result)
//	}
//
//	return result
//}
//
//func isDiffAssetGateway(src, dst document.AssetGateways) bool {
//	if src.Website != dst.Website ||
//		src.Details != dst.Details ||
//		src.Moniker != dst.Moniker ||
//		src.Identity != dst.Identity ||
//		src.Owner != dst.Owner {
//		return true
//	}
//	return false
//}
