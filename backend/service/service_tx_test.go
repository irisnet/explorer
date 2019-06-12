package service

import (
	"testing"

	"github.com/irisnet/explorer/backend/model"
	"gopkg.in/mgo.v2/bson"
)

func TestQueryTxList(t *testing.T) {
	txPage := new(TxService).QueryTxList(nil, 0, 100)

	t.Logf("total: %v \n", txPage.Count)

	t.Logf("items: %v \n", txPage.Data)
}

func TestTxQueryList(t *testing.T) {

	txPage := new(TxService).QueryList(nil, 0, 100)
	t.Logf("total: %v \n", txPage.Count)
	t.Logf("items: %v \n", txPage.Data)
}

func TestQueryRecentTx(t *testing.T) {

	txList := new(TxService).QueryRecentTx()

	for k, v := range txList {
		t.Logf("idx: %v v: %v \n", k, v)
	}
}

func TestQueryTxByHash(t *testing.T) {

	tx := new(TxService).Query("C8C9090F3504A39403F243F1B83F75734E328A4742BC904BC1CBAD471D4FF66F")
	t.Logf("tx: %v\n", tx)
}

func TestQueryByAcc(t *testing.T) {

	txPage := new(TxService).QueryByAcc("faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn", 0, 10)

	t.Logf("total: %v \n", txPage.Count)
	if modelV, ok := txPage.Data.([]model.CommonTx); ok {
		for k, v := range modelV {
			t.Logf("idx: %v  v: %v \n", k, v)
		}
	}
}

func TestCountByType(t *testing.T) {

	statistic := new(TxService).CountByType(bson.M{})
	t.Logf("tx statistic by type: %v \n", statistic)
}

func TestQueryTxNumGroupByDay(t *testing.T) {

	txCountByDay := new(TxService).QueryTxNumGroupByDay()

	for k, v := range txCountByDay {
		t.Logf("idx: %v  txCountByDay: %v \n", k, v)
	}
}
