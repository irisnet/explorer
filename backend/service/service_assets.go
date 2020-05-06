package service

import (
	"fmt"
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

func (assets *AssetsService) GetGatewayAsset(symbol, gateway, txtype string, page, size int, istotal bool) (vo.AssetsRespond, error) {

	if !isFieldTokenType(txtype) {
		txtype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(document.Tx_AssetType_Gateway, txtype, symbol, gateway, page, size, istotal)
	if err != nil {
		logger.Error("GetGatewayAsset have error", logger.String("error", err.Error()))
		return vo.AssetsRespond{}, err
	}
	result := make([]vo.AssetsVo, 0, len(retassets))
	for _, asset := range retassets {
		result = append(result, LoadModelFromCommonTx(asset))
	}

	return vo.AssetsRespond{
		Total:     total,
		Txs:       result,
		AssetType: document.Tx_AssetType_Gateway,
	}, nil
}

func (assets *AssetsService) GetTransferGatewayOwner(moniker string, page, size int, istotal bool) (vo.AssetsRespond, error) {

	total, retassets, err := document.CommonTx{}.QueryTxTransferGatewayOwner(moniker, page, size, istotal)
	if err != nil {
		logger.Error("GetTransferGatewayOwner have error", logger.String("error", err.Error()))
		return vo.AssetsRespond{}, err
	}
	result := make([]vo.AssetsVo, 0, len(retassets))
	for _, asset := range retassets {
		result = append(result, LoadModelFromCommonTx(asset))
	}

	return vo.AssetsRespond{
		Total: total,
		Txs:   result,
	}, nil
}

func isFieldTokenType(tokentype string) bool {
	return tokentype == document.Tx_Asset_TxType_Mint ||
		tokentype == document.Tx_Asset_TxType_Issue ||
		tokentype == document.Tx_Asset_TxType_Edit ||
		tokentype == document.Tx_Asset_TxType_TransferOwner ||
		tokentype == document.Tx_Asset_TxType_TransferGatewayOwner
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
			dst = converModelTokenId(dst)
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
			dst.MintTo = checkMintToAddress(msgData.Owner, msgData.To)
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
			dst.SrcOwner = msgData.Owner
			dst.DstOwner = msgData.To
			dst.Gateway = msgData.Moniker
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

func converModelTokenId(dst vo.AssetsVo) vo.AssetsVo {
	dst.TokenId = dst.Symbol
	if dst.Gateway != "" {
		dst.TokenId = dst.Gateway + "." + dst.Symbol
	}
	return dst
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
		vMap[v.TokenId] = v
	}

	dstAssetTokens := buildAssetTokens()
	var txs []txn.Op
	for _, v := range dstAssetTokens {
		if it, ok := vMap[v.TokenId]; ok {
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
			delete(vMap, v.TokenId)
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
	var buildAssetToken = func(v lcd.AssetTokens) (document.AssetToken, error) {
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

	var genAssetTokens = func(va lcd.AssetTokens, result *[]document.AssetToken) {
		assetToken, err := buildAssetToken(va)
		if err != nil {
			logger.Error("utils.copy assetToken failed")
			panic(err)
		}
		denome := assetToken.Symbol + "-min"
		if assetToken.Symbol == "iris" {
			return
		} else {
			if len(assetToken.Gateway) > 0 {
				denome = assetToken.Gateway + "." + denome
			}
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
	if src.TokenId != dst.TokenId ||
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
		src.TotalSupply != dst.TotalSupply ||
		src.Mintable != dst.Mintable ||
		src.Owner != dst.Owner {
		return true
	}
	return false
}

func (service *AssetsService) QueryAssetGatewayDetail(moniker string) (vo.AssetGateways, error) {
	res, err := document.AssetGateways{}.GetGatewayInfo(moniker)
	if err != nil {
		logger.Error("QueryAssetGateways", logger.String("err", err.Error()))
		return vo.AssetGateways{}, err
	}

	return vo.AssetGateways{
		Owner:    res.Owner,
		Identity: res.Identity,
		Website:  res.Website,
		Details:  res.Details,
		Moniker:  res.Moniker,
		Icons:    res.Icons,
	}, nil
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
			TokenId:         v.TokenId,
			Owner:           v.Owner,
			Gateway:         v.Gateway,
			Family:          v.Family,
			TotalSupply:     utils.CovertAssetUnit(v.TotalSupply, v.Decimal),
			InitialSupply:   utils.CovertAssetUnit(v.InitialSupply, v.Decimal),
			MaxSupply:       utils.CovertAssetUnit(v.MaxSupply, v.Decimal),
			MinUnitAlias:    v.MinUnitAlias,
			Mintable:        v.Mintable,
			Name:            v.Name,
			CanonicalSymbol: v.CanonicalSymbol,
			Decimal:         v.Decimal,
			Symbol:          v.Symbol,
			Source:          v.Source,
		}
		assetinfos = append(assetinfos, tmp)
	}

	return assetinfos, nil
}

func (service *AssetsService) QueryAssetTokenDetail(tokenid string) (vo.AssetTokens, error) {

	res, err := document.AssetToken{}.GetAssetTokenDetail(tokenid)
	if err != nil {
		logger.Error("GetAssetByAddr", logger.String("err", err.Error()))
		return vo.AssetTokens{}, err
	}
	ret := vo.AssetTokens{
		TokenId:         res.TokenId,
		Owner:           res.Owner,
		Gateway:         res.Gateway,
		Family:          res.Family,
		TotalSupply:     utils.CovertAssetUnit(res.TotalSupply, res.Decimal),
		InitialSupply:   utils.CovertAssetUnit(res.InitialSupply, res.Decimal),
		MaxSupply:       utils.CovertAssetUnit(res.MaxSupply, res.Decimal),
		MinUnitAlias:    res.MinUnitAlias,
		Mintable:        res.Mintable,
		Name:            res.Name,
		CanonicalSymbol: res.CanonicalSymbol,
		Decimal:         res.Decimal,
		Symbol:          res.Symbol,
		Source:          res.Source,
	}
	if res.Source == document.Tx_AssetType_Gateway {
		gatewaydata, err := document.AssetGateways{}.GetGatewayInfo(res.Gateway)
		if err != nil {
			logger.Warn("GetGatewayInfo have error", logger.String("err", err.Error()))
		}
		ret.AssetGateway = &vo.AssetGateways{
			Owner:    gatewaydata.Owner,
			Identity: gatewaydata.Identity,
			Website:  gatewaydata.Website,
			Details:  gatewaydata.Details,
			Moniker:  gatewaydata.Moniker,
			Icons:    gatewaydata.Icons,
		}

	}

	return ret, nil
}
