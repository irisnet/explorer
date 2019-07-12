package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
)

type AssetsService struct {
	BaseService
}

func (assets *AssetsService) GetNativeAsset(tokentype string, page, size int) (model.AssetsRespond, error) {

	if !isFieldTokenType(tokentype) {
		tokentype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(document.Tx_AssetType_Native, tokentype, page, size)
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
		Assets:    result,
		AssetType: document.Tx_AssetType_Native,
	}, nil
}

func (assets *AssetsService) GetGatewayAsset(tokentype string, page, size int) (model.AssetsRespond, error) {

	if !isFieldTokenType(tokentype) {
		tokentype = ""
	}

	total, retassets, err := document.CommonTx{}.QueryTxAsset(document.Tx_AssetType_Gateway, tokentype, page, size)
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
		Assets:    result,
		AssetType: document.Tx_AssetType_Gateway,
	}, nil
}

func isFieldTokenType(tokentype string) bool {
	return tokentype == document.Tx_Asset_TokenType_Issue ||
		tokentype == document.Tx_Asset_TokenType_TransferOwner ||
		tokentype == document.Tx_Asset_TokenType_Edit ||
		tokentype == document.Tx_Asset_TokenType_Mint
}

func LoadModelFromCommonTx(src document.CommonTx) (dst model.AssetsVo) {

	dst.Height = src.Height
	dst.TxHash = src.TxHash
	dst.TxStatus = src.Status
	dst.Timestamp = src.Time

	dst.Amount = convertModelCoins(src.Amount)
	dst.TxFee = convertModelFee(src.Fee)

	dst.Type = src.Msgs[0].Type
	dst.TokenId = src.Msgs[0].MsgData.TokenId
	dst.SymbolAtSource = src.Msgs[0].MsgData.SymbolAtSource
	dst.Symbol = src.Msgs[0].MsgData.Symbol
	dst.SymbolMin = src.Msgs[0].MsgData.SymbolMinAlias

	dst.InitialSupply = src.Msgs[0].MsgData.InitialSupply
	dst.MaxSupply = src.Msgs[0].MsgData.MaxSupply

	dst.Name = src.Msgs[0].MsgData.Name
	dst.Decimal = src.Msgs[0].MsgData.Decimal

	dst.DstOwner = src.Msgs[0].MsgData.DstOwner
	dst.SrcOwner = src.Msgs[0].MsgData.SrcOwner
	dst.Owner = src.Msgs[0].MsgData.Owner

	dst.Gateway = src.Msgs[0].MsgData.Gateway
	dst.MintTo = src.Msgs[0].MsgData.To
	dst.Mintable = src.Msgs[0].MsgData.Mintable

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
