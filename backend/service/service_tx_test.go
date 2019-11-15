package service

import (
	"github.com/irisnet/explorer/backend/utils"
	"testing"

	"encoding/json"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/vo/msgvo"
)

func TestQueryTxList(t *testing.T) {
	txPage := new(TxService).QueryTxList(nil, 0, 100, true)

	t.Logf("total: %v \n", txPage.Count)

	bytestr, _ := json.Marshal(txPage.Data)
	t.Logf("items: %v \n", string(bytestr))
}

func TestQueryRecentTx(t *testing.T) {

	txList := new(TxService).QueryRecentTx()

	for k, v := range txList {
		t.Logf("idx: %v v: %v \n", k, v)
	}
}

func TestQueryTxByHash(t *testing.T) {
	tx := new(TxService).Query("3EB9D229189C84A0A11FC4A19154FEE5CDC0E2C10E6695E316F6888BE2677C3C")
	t.Logf("tx: %v\n", string(utils.MarshalJsonIgnoreErr(tx)))
}

func TestServiceTxfetchLogMessage(t *testing.T) {
	log := "Msg 0 failed: {\"codespace\":\"sdk\",\"code\":10,\"message\":\"12097471760000000000000iris-atto is less than 100000000000000000000000iris-atto\"}"
	ret := fetchLogMessage(log)
	t.Log(ret)
}

func TestQueryByAcc(t *testing.T) {

	txPage := new(TxService).QueryByAcc("faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn", 0, 10, true)

	t.Logf("total: %v \n", txPage.Count)
	if modelV, ok := txPage.Data.([]vo.CommonTx); ok {
		for k, v := range modelV {
			t.Logf("idx: %v  v: %v \n", k, v)
		}
	}
}

//func TestCountByType(t *testing.T) {
//
//	statistic := new(TxService).CountByType(bson.M{})
//	t.Logf("tx statistic by type: %v \n", statistic)
//}

func TestQueryTxNumGroupByDay(t *testing.T) {

	txCountByDay := new(TxService).QueryTxNumGroupByDay()

	for k, v := range txCountByDay {
		t.Logf("idx: %v  txCountByDay: %v \n", k, v)
	}
}

func TestTxService_checkTags(t *testing.T) {
	tags := map[string]string{
		"voting-period-start": "41",
		"action":              "submit_proposal",
		"proposer":            "faa1x292qss22x4rls6ygr7hhnp0et94vwwrchaklp",
		"proposal-id":         "41",
		"param":               "[{\"subspace\":\"stake\",\"key\":\"UnbondingTime\",\"value\":\"123m\"}]",
	}
	tags1 := map[string]string{
		"voting-period-start": "41",
		"action":              "submit_proposal",
		"proposer":            "faa1x292qss22x4rls6ygr7hhnp0et94vwwrchaklp",
		"proposal-id":         "41",
	}
	submitprodata := msgvo.TxMsgSubmitProposal{
		Params: []msgvo.Param{
			{Subspace: "stake", Key: "UnbondingTime", Value: "12m"},
		},
	}
	data := checkTags(tags, submitprodata.Params)
	t.Log(tags)
	t.Log(data)
	data1 := checkTags(tags1, submitprodata.Params)
	t.Log(tags1)
	t.Log(data1)

}

func TestTxService_QueryBaseList(t *testing.T) {
	res := new(TxService).QueryBaseList(bson.M{"from": "faa174qyl02cupyqq77cqqtdl0frda6dl3rpjcrgnp"}, 1, 10, false)
	bytestr, _ := json.Marshal(res)
	t.Logf(" %v \n", string(bytestr))
}