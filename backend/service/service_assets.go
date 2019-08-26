package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"fmt"
)

type AssetsService struct {
	BaseService
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
	dst.Amount = src.Msgs[0].MsgData.Amount
	dst.TokenId = src.Msgs[0].MsgData.TokenId
	dst.CanonicalSymbol = src.Msgs[0].MsgData.CanonicalSymbol
	dst.Symbol = src.Msgs[0].MsgData.Symbol
	if dst.Type == document.Tx_Asset_TxType_Issue {
		source := src.Msgs[0].MsgData.UdInfo.Source
		if source == document.Tx_AssetType_Native {
			dst.SymbolMin = fmt.Sprintf("%s-min", src.Msgs[0].MsgData.UdInfo.Symbol)
		} else if source == document.Tx_AssetType_Gateway {
			dst.SymbolMin = fmt.Sprintf("%s.%s-min", src.Msgs[0].MsgData.UdInfo.Gateway, src.Msgs[0].MsgData.UdInfo.Symbol)
		}

	} else {
		dst.SymbolMin = fmt.Sprintf("%s-min", src.Msgs[0].MsgData.TokenId)
	}

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
