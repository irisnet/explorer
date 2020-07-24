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
	"gopkg.in/mgo.v2/txn"
)

type AssetsService struct {
	BaseService
}

func (asset AssetsService) GetModule() Module {
	return Asset
}

func (assets *AssetsService) GetNativeAsset(symbol, txtype string, page, size int, istotal bool) (vo.AssetsRespond, error) {

	if !isFieldTokenType(txtype) {
		txtype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(document.Tx_AssetType_Native, txtype, symbol, "", page, size, istotal)
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
	//tokentype == document.Tx_Asset_TxType_TransferGatewayOwner
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

func (service *AssetsService) UpdateAssetTokens(vs []document.AssetToken) error {
	var vMap = make(map[string]document.AssetToken)
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
	return document.AssetToken{}.Batch(txs)
}

func buildAssetTokens() []document.AssetToken {
	res := lcd.GetAssetTokens()
	var result []document.AssetToken
	var buildAssetToken = func(v lcd.AssetToken) (document.AssetToken, error) {
		var assetToken document.AssetToken
		if err := utils.Copy(v.BaseToken, &assetToken); err != nil {
			logger.Error("utils.copy assetToken failed")
			return assetToken, err
		}
		return assetToken, nil
	}

	tokenstats, err := lcd.GetBankTokenStats()
	if err != nil {
		logger.Error("buildAssetTokens GetBankTokenStats failed")
		return nil
	}
	asset_totalsupplyMap := make(map[string]string, len(tokenstats.TotalSupply))
	for _, val := range tokenstats.TotalSupply {
		asset_totalsupplyMap[val.Denom] = val.Amount
	}

	var genAssetTokens = func(va lcd.AssetToken, result *[]document.AssetToken) {
		assetToken, err := buildAssetToken(va)
		if err != nil {
			logger.Error("utils.copy assetToken failed")
			panic(err)
		}
		denome := assetToken.Symbol + "-min"
		if assetToken.Symbol == "stake" {
			return
		}

		if totalsupply, ok := asset_totalsupplyMap[denome]; ok {
			assetToken.TotalSupply = totalsupply
		}
		*result = append(*result, assetToken)
	}
	for _, v := range res {

		genAssetTokens(v, &result)
	}
	return result
}

func isDiffAssetToken(src, dst document.AssetToken) bool {
	if src.Symbol != dst.Symbol ||
		src.Name != dst.Name ||
		src.Scale != dst.Scale ||
		src.MinUnitAlias != dst.MinUnitAlias ||
		src.InitialSupply != dst.InitialSupply ||
		src.MaxSupply != dst.MaxSupply ||
		src.TotalSupply != dst.TotalSupply ||
		src.Mintable != dst.Mintable ||
		src.Owner != dst.Owner {
		return true
	}
	return false
}


func (service *AssetsService) QueryAssetTokens(source string) (vo.AssetTokensRespond, error) {
	res, err := document.AssetToken{}.GetAssetTokens(source)
	if err != nil {
		logger.Error("GetAssetByAddr", logger.String("err", err.Error()))
		return []vo.AssetTokens{}, err
	}

	assetinfos := make([]vo.AssetTokens, 0, len(res))

	for _, v := range res {
		tmp := vo.AssetTokens{
			Owner:         v.Owner,
			TotalSupply:   utils.CovertAssetUnit(v.TotalSupply, v.Scale),
			InitialSupply: utils.CovertAssetUnit(v.InitialSupply, v.Scale),
			MaxSupply:     utils.CovertAssetUnit(v.MaxSupply, v.Scale),
			MinUnitAlias:  v.MinUnitAlias,
			Mintable:      v.Mintable,
			Name:          v.Name,
			Decimal:       v.Scale,
			Symbol:        v.Symbol,
		}
		assetinfos = append(assetinfos, tmp)
	}

	return assetinfos, nil
}

func (service *AssetsService) QueryAssetTokenDetail(symbol string) (vo.AssetTokens, error) {

	res, err := document.AssetToken{}.GetAssetTokenDetail(symbol)
	if err != nil {
		logger.Error("GetAssetByAddr", logger.String("err", err.Error()))
		return vo.AssetTokens{}, err
	}
	ret := vo.AssetTokens{
		Owner:         res.Owner,
		TotalSupply:   utils.CovertAssetUnit(res.TotalSupply, res.Scale),
		InitialSupply: utils.CovertAssetUnit(res.InitialSupply, res.Scale),
		MaxSupply:     utils.CovertAssetUnit(res.MaxSupply, res.Scale),
		MinUnitAlias:  res.MinUnitAlias,
		Mintable:      res.Mintable,
		Name:          res.Name,
		Decimal:       res.Scale,
		Symbol:        res.Symbol,
	}


	return ret, nil
}
