package service

import (
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/vo/msgvo"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/types"
)

type HtlcService struct {
	BaseService
}

func (service *HtlcService) GetModule() Module {
	return Htlc
}

func (service *HtlcService) QueryHtlcByHashLock(hashlock string) vo.HtlcInfo {

	var resp vo.HtlcInfo
	htlcinfo, err := lcd.HtlcInfo(hashlock)
	if err != nil {
		logger.Error("HtlcInfo from lcd have error", logger.String("err", err.Error()))
		return resp
	}
	resp.From = htlcinfo.Value.Sender
	resp.HashLock = hashlock
	resp.To = htlcinfo.Value.To
	resp.ExpireHeight = htlcinfo.Value.ExpireHeight
	resp.Timestamp = htlcinfo.Value.Timestamp
	resp.CrossChainReceiver = htlcinfo.Value.ReceiverOnOtherChain
	resp.State = htlcinfo.Value.State
	for _, val := range htlcinfo.Value.Amount {
		resp.Amount = append(resp.Amount, LoadCoinVoFromLcdCoin(val))
	}
	query := bson.M{
		document.Tx_Field_Type:          types.TxTypeCreateHTLC,
		document.Tx_Field_Msgs_Hashcode: hashlock,
		document.Tx_Field_Status:        "success",
	}


	txAsDoc, err := document.CommonTx{}.QueryHtlcTx(query)
	if err != nil {
		logger.Error("get HtlcInfo from db have error", logger.String("err", err.Error()))
	}
	msgVO := msgvo.TxMsgCreateHTLC{}
	if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(txAsDoc.Msgs[0].MsgData)); err != nil {
		logger.Error("BuildTxMsgRequestRandByUnmarshalJson", logger.String("err", err.Error()))
	}
	resp.TimeLock = int64(msgVO.TimeLock)

	resp.FromMoniker, resp.ToMoniker = service.BuildFTMoniker(resp.From, resp.To)
	return resp
}
