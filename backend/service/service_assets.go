package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	dmsg "github.com/irisnet/explorer/backend/orm/document/msg"
	"github.com/irisnet/explorer/backend/types"
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
	switch dst.Type {
	case types.TxTypeIssueToken:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgIssueToken)
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

	case types.TxTypeEditToken:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgEditToken)
		dst.TokenId = msgData.TokenId
		dst.Owner = msgData.Owner
		dst.CanonicalSymbol = msgData.CanonicalSymbol
		dst.MaxSupply = msgData.MaxSupply
		dst.Mintable = msgData.Mintable
		dst.Name = msgData.Name
		dst.SymbolMin = msgData.TokenId

	case types.TxTypeMintToken:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgMintToken)
		dst.TokenId = msgData.TokenId
		dst.Owner = msgData.Owner
		dst.Amount = msgData.Amount
		dst.SymbolMin = msgData.TokenId
		dst.MintTo = msgData.To

	case types.TxTypeTransferTokenOwner:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgTransferTokenOwner)
		dst.SrcOwner = msgData.SrcOwner
		dst.DstOwner = msgData.DstOwner
		dst.TokenId = msgData.TokenId
		dst.SymbolMin = msgData.TokenId

	case types.TxTypeCreateGateway:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgCreateGateway)
		dst.Owner = msgData.Owner

	case types.TxTypeEditGateway:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgEditGateway)
		dst.Owner = msgData.Owner

	case types.TxTypeTransferGatewayOwner:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgTransferGatewayOwner)
		dst.Owner = msgData.Owner
		dst.MintTo = msgData.To

	case types.TxTypeSetMemoRegexp:
		msgData := src.Msgs[0].MsgData.(dmsg.TxMsgSetMemoRegexp)
		dst.Owner = msgData.Owner

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
