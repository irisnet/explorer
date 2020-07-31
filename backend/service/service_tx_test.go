package service

import (
	"github.com/irisnet/explorer/backend/utils"
	"testing"

	"encoding/json"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
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
	tx := new(TxService).QueryTxDetail("48F166919DB0394BA1553C013C2CF26EFC2ED870C27017C016C651BCDF5F97B6")
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

	txCountByDay := new(TxService).QueryTxNumGroupByDay(3)

	t.Log(string(utils.MarshalJsonIgnoreErr(txCountByDay)))
}

func TestTxService_checkTags(t *testing.T) {
	//tags := map[string]string{
	//	"voting-period-start": "41",
	//	"action":              "submit_proposal",
	//	"proposer":            "faa1x292qss22x4rls6ygr7hhnp0et94vwwrchaklp",
	//	"proposal-id":         "41",
	//	"param":               "[{\"subspace\":\"stake\",\"key\":\"UnbondingTime\",\"value\":\"123m\"}]",
	//}
	//tags1 := map[string]string{
	//	"voting-period-start": "41",
	//	"action":              "submit_proposal",
	//	"proposer":            "faa1x292qss22x4rls6ygr7hhnp0et94vwwrchaklp",
	//	"proposal-id":         "41",
	//}
	//submitprodata := msgvo.TxMsgSubmitProposal{
	//	//Params: []msgvo.Param{
	//	//	{Subspace: "stake", Key: "UnbondingTime", Value: "12m"},
	//	//},
	//}
	//data := checkTags(tags, submitprodata.Params)
	//t.Log(tags)
	//t.Log(data)
	//data1 := checkTags(tags1, submitprodata.Params)
	//t.Log(tags1)
	//t.Log(data1)

}

func TestTxService_QueryBaseList(t *testing.T) {

	res := new(TxService).QueryBaseList(bson.M{"$or": []bson.M{{"signers.addr_bech32": "iaa1x98k5n7xj0h3udnf5dcdzw85tsfa75qm0kqak0"},
		{"to": "iaa1x98k5n7xj0h3udnf5dcdzw85tsfa75qm0kqak0"}}}, 1, 10, true)
	bytestr, _ := json.Marshal(res)
	t.Logf(" %v \n", string(bytestr))
}
