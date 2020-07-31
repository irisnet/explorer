package service

import (
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/vo/msgvo"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
)

type AssetsService struct {
	BaseService
}

func (assets *AssetsService) GetNativeAsset(symbol, txtype string, page, size int, istotal bool) (vo.AssetsRespond, error) {

	if !isFieldTokenType(txtype) {
		txtype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(txtype, symbol, page, size, istotal)
	if err != nil {
		logger.Error("GetNativeAsset have error", logger.String("error", err.Error()))
		return vo.AssetsRespond{}, err
	}

	result := make([]vo.AssetsVo, 0, len(retassets))
	for _, asset := range retassets {
		result = append(result, LoadModelFromCommonTx(asset))
	}

	return vo.AssetsRespond{
		Total:     total,
		Txs:       result,
		AssetType: document.Tx_AssetType_Native,
	}, nil
}

func isFieldTokenType(tokentype string) bool {
	return tokentype == document.Tx_Asset_TxType_Mint ||
		tokentype == document.Tx_Asset_TxType_Issue ||
		tokentype == document.Tx_Asset_TxType_Edit ||
		tokentype == document.Tx_Asset_TxType_TransferOwner
}

func LoadModelFromCommonTx(src document.CommonTx) (dst vo.AssetsVo) {

	dst.Height = src.Height
	dst.TxHash = src.TxHash
	dst.TxStatus = utils.FailtoFailed(src.Status)
	dst.Timestamp = src.Time.UTC()

	dst.TxFee = convertModelActualFee(src.ActualFee)

	dst.Type = src.Msgs[0].Type
	switch dst.Type {
	case types.TxTypeIssueToken:
		msgData := msgvo.TxMsgIssueToken{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Symbol = msgData.Symbol
			dst.Name = msgData.Name
			dst.Decimal = msgData.Scale
			dst.InitialSupply = msgData.InitialSupply
			dst.MaxSupply = msgData.MaxSupply
			dst.Mintable = msgData.Mintable
			dst.Owner = msgData.Owner
			dst.SymbolMin = msgData.MinUnit
		}

	case types.TxTypeEditToken:
		msgData := msgvo.TxMsgEditToken{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Owner = msgData.Owner
			dst.MaxSupply = msgData.MaxSupply
			dst.Mintable = msgData.Mintable
			dst.Name = msgData.Name
			dst.Symbol = msgData.Symbol
		}

	case types.TxTypeMintToken:
		msgData := msgvo.TxMsgMintToken{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.Owner = msgData.Owner
			dst.Amount = msgData.Amount
			dst.Symbol = msgData.Symbol
			dst.MintTo = checkMintToAddress(msgData.Owner, msgData.To)
		}

	case types.TxTypeTransferTokenOwner:
		msgData := msgvo.TxMsgTransferTokenOwner{}
		if err := msgData.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(src.Msgs[0].MsgData)); err != nil {
			logger.Error("LoadModelFromCommonTx have error", logger.String("err", err.Error()))
		} else {
			dst.SrcOwner = msgData.SrcOwner
			dst.DstOwner = msgData.DstOwner
			dst.Symbol = msgData.Symbol
		}

	}

	return
}

func checkMintToAddress(owner, address string) string {
	if len(owner) != len(address) {
		return owner
	}
	return address
}

func convertModelCoin(coin document.Coin) vo.Coin {
	return vo.Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount,
	}
}
func convertModelCoins(coins document.Coins) (dst vo.Coins) {

	for _, coin := range coins {
		dst = append(dst, convertModelCoin(coin))
	}
	return
}

func convertModelFee(fee document.Fee) vo.Fee {
	return vo.Fee{
		Gas:    fee.Gas,
		Amount: convertModelCoins(fee.Amount),
	}
}

func convertModelActualFee(actfee document.ActualFee) vo.ActualFee {
	return vo.ActualFee{
		Amount: actfee.Amount,
		Denom:  actfee.Denom,
	}
}

func (service *AssetsService) UpdateAssetTokens() {

	var vMap map[string]bson.ObjectId
	validators, err := document.AssetToken{}.GetAllAssets()
	if err == nil {
		vMap = make(map[string]bson.ObjectId, len(validators))
		for _, val := range validators {
			vMap[val.Symbol] = val.ID
		}
	}

	var assetModel document.AssetToken
	res := lcd.GetAssetTokens()
	for _, v := range res {
		item := document.AssetToken{
			Symbol:        v.BaseToken.Symbol,
			Scale:         v.BaseToken.Scale,
			AssetName:     v.BaseToken.Name,
			MinUnit:       v.BaseToken.MinUnitAlias,
			MaxSupply:     v.BaseToken.MaxSupply,
			Mintable:      v.BaseToken.Mintable,
			Owner:         v.BaseToken.Owner,
			InitialSupply: v.BaseToken.InitialSupply,
		}
		one, err := assetModel.GetAssetTokenDetail(v.BaseToken.Symbol)
		if err != nil && err == mgo.ErrNotFound {
			item.ID = bson.NewObjectId()
			if err := assetModel.Save(item); err != nil {
				logger.Error("asset token save failed",
					logger.String("err", err.Error()))
			}
		} else {
			if err == nil && isDiffAssetToken(one, item) {
				item.ID = one.ID
				if err := assetModel.Update(item); err != nil {
					logger.Error("asset token update failed",
						logger.String("err", err.Error()))
				}
			}
			delete(vMap, item.Symbol)
		}

	}
	//clear no exist asset
	if len(vMap) > 0 && len(res) > 0 {
		var txs []txn.Op
		for sym := range vMap {
			v := vMap[sym]
			txs = append(txs, txn.Op{
				C:      document.CollectionNmAsset,
				Id:     v,
				Remove: true,
			})
		}
		if len(txs) > 0 {
			err := document.AssetToken{}.Batch(txs)
			if err != nil {
				logger.Error("clean no exist tokens is failed", logger.String("err", err.Error()))
			}
		}
	}
}

func isDiffAssetToken(src, dst document.AssetToken) bool {
	if src.Symbol != dst.Symbol ||
		src.AssetName != dst.AssetName ||
		src.Scale != dst.Scale ||
		src.MinUnit != dst.MinUnit ||
		src.InitialSupply != dst.InitialSupply ||
		src.MaxSupply != dst.MaxSupply ||
		src.Owner != dst.Owner {
		return true
	}
	return false
}

func (service *AssetsService) QueryAssetTokens() (vo.AssetTokensRespond, error) {
	res, err := document.AssetToken{}.GetAssetTokens()
	if err != nil {
		logger.Error("Query AssetTokens", logger.String("err", err.Error()))
		return []vo.AssetTokens{}, err
	}

	assetinfos := make([]vo.AssetTokens, 0, len(res))

	for _, v := range res {
		tmp := vo.AssetTokens{
			Owner:         v.Owner,
			InitialSupply: v.InitialSupply,
			MaxSupply:     v.MaxSupply,
			MinUnit:       v.MinUnit,
			Mintable:      v.Mintable,
			Name:          v.AssetName,
			Scale:         v.Scale,
			Symbol:        v.Symbol,
		}
		assetinfos = append(assetinfos, tmp)
	}

	return assetinfos, nil
}

func (service *AssetsService) QueryAssetTokenDetail(symbol string) (vo.AssetTokens, error) {

	res, err := document.AssetToken{}.GetAssetTokenDetail(symbol)
	if err != nil {
		logger.Error("Query AssetTokenDetail", logger.String("err", err.Error()))
		return vo.AssetTokens{}, err
	}
	ret := vo.AssetTokens{
		Owner: res.Owner,
		//TotalSupply:   utils.CovertAssetUnit(res.TotalSupply, res.Scale),
		InitialSupply: res.InitialSupply,
		MaxSupply:     res.MaxSupply,
		MinUnit:       res.MinUnit,
		Mintable:      res.Mintable,
		Name:          res.AssetName,
		Scale:         res.Scale,
		Symbol:        res.Symbol,
	}

	return ret, nil
}
