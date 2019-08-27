package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/model/msgvo"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type AssetsService struct {
	BaseService
}

func (asset AssetsService) GetModule() Module {
	return Asset
}

func (assets *AssetsService) GetNativeAsset(txtype string, page, size int, istotal bool) (model.AssetsRespond, error) {

	if !isFieldTokenType(txtype) {
		txtype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(document.Tx_AssetType_Native, txtype, page, size, istotal)
	if err != nil {
		logger.Error("GetNativeAsset have error", logger.String("error", err.Error()))
		return model.AssetsRespond{}, err
	}

	result := make([]model.AssetsVo, 0, len(retassets))
	for _, asset := range retassets {
		result = append(result, LoadModelFromCommonTx(asset))
	}

	return model.AssetsRespond{
		Total:     total,
		Txs:       result,
		AssetType: document.Tx_AssetType_Native,
	}, nil
}

func (assets *AssetsService) GetGatewayAsset(txtype string, page, size int, istotal bool) (model.AssetsRespond, error) {

	if !isFieldTokenType(txtype) {
		txtype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(document.Tx_AssetType_Gateway, txtype, page, size, istotal)
	if err != nil {
		logger.Error("GetNativeAsset have error", logger.String("error", err.Error()))
		return model.AssetsRespond{}, err
	}
	result := make([]model.AssetsVo, 0, len(retassets))
	for _, asset := range retassets {
		result = append(result, LoadModelFromCommonTx(asset))
	}

	return model.AssetsRespond{
		Total:     total,
		Txs:       result,
		AssetType: document.Tx_AssetType_Gateway,
	}, nil
}

func isFieldTokenType(tokentype string) bool {
	return tokentype == document.Tx_Asset_TxType_Mint ||
		tokentype == document.Tx_Asset_TxType_Issue ||
		tokentype == document.Tx_Asset_TxType_Edit ||
		tokentype == document.Tx_Asset_TxType_TransferOwner
}

func LoadModelFromCommonTx(src document.CommonTx) (dst model.AssetsVo) {

	dst.Height = src.Height
	dst.TxHash = src.TxHash
	dst.TxStatus = src.Status
	dst.Timestamp = src.Time

	dst.TxFee = convertModelFee(src.Fee)

	dst.Type = src.Msgs[0].Type
	switch dst.Type {
	case types.TxTypeIssueToken:
		msgData := msgvo.TxMsgIssueToken{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Gateway = msgData.Gateway
			dst.Symbol = msgData.Symbol
			dst.CanonicalSymbol = msgData.CanonicalSymbol
			dst.Name = msgData.Name
			dst.Decimal = msgData.Decimal
			dst.InitialSupply = msgData.InitialSupply
			dst.MaxSupply = msgData.MaxSupply
			dst.Mintable = msgData.Mintable
			dst.Owner = msgData.Owner
			source := msgData.UdInfo.Source
			if source == document.Tx_AssetType_Native {
				dst.SymbolMin = fmt.Sprintf("%s-min", msgData.UdInfo.Symbol)
			} else if source == document.Tx_AssetType_Gateway {
				dst.SymbolMin = fmt.Sprintf("%s.%s-min", msgData.UdInfo.Gateway, msgData.UdInfo.Symbol)
			}
		}

	case types.TxTypeEditToken:
		msgData := msgvo.TxMsgEditToken{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.TokenId = msgData.TokenId
			dst.Owner = msgData.Owner
			dst.CanonicalSymbol = msgData.CanonicalSymbol
			dst.MaxSupply = msgData.MaxSupply
			dst.Mintable = msgData.Mintable
			dst.Name = msgData.Name
			dst.SymbolMin = msgData.TokenId
		}

	case types.TxTypeMintToken:
		msgData := msgvo.TxMsgMintToken{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.TokenId = msgData.TokenId
			dst.Owner = msgData.Owner
			dst.Amount = msgData.Amount
			dst.SymbolMin = msgData.TokenId
			dst.MintTo = msgData.To
		}

	case types.TxTypeTransferTokenOwner:
		msgData := msgvo.TxMsgTransferTokenOwner{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.SrcOwner = msgData.SrcOwner
			dst.DstOwner = msgData.DstOwner
			dst.TokenId = msgData.TokenId
			dst.SymbolMin = msgData.TokenId
		}

	case types.TxTypeCreateGateway:
		msgData := msgvo.TxMsgCreateGateway{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Owner = msgData.Owner
		}

	case types.TxTypeEditGateway:
		msgData := msgvo.TxMsgEditGateway{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Owner = msgData.Owner
		}

	case types.TxTypeTransferGatewayOwner:
		msgData := msgvo.TxMsgTransferGatewayOwner{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Owner = msgData.Owner
			dst.MintTo = msgData.To
		}

	case types.TxTypeSetMemoRegexp:
		msgData := msgvo.TxMsgSetMemoRegexp{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Owner = msgData.Owner
		}

	}

	return
}

func convertModelCoin(coin document.Coin) model.Coin {
	return model.Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount,
	}
}
func convertModelCoins(coins document.Coins) (dst model.Coins) {

	for _, coin := range coins {
		dst = append(dst, convertModelCoin(coin))
	}
	return
}

func convertModelFee(fee document.Fee) model.Fee {
	return model.Fee{
		Gas:    fee.Gas,
		Amount: convertModelCoins(fee.Amount),
	}
}

func (service *AssetsService) UpdateAssetTokens(vs []document.Asset) error {
	var vMap = make(map[string]document.Asset)
	for _, v := range vs {
		vMap[v.Id] = v
	}

	dstAssetTokens := buildAssetTokens()
	var txs []txn.Op
	for _, v := range dstAssetTokens {
		if it, ok := vMap[v.Id]; ok {
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
			delete(vMap, v.Id)
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

	for _, v := range res {
		var genAssetTokens = func(va lcd.AssetTokens, result *[]document.Asset) {
			assetToken, err := buildAssetToken(va)
			if err != nil {
				logger.Error("utils.copy assetToken failed")
				panic(err)
			}
			*result = append(*result, assetToken)
		}
		genAssetTokens(v, &result)
	}
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

func (service *AssetsService) QueryAssetToken() []model.AssetTokens {
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
