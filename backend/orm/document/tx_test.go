package document

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
)

func TestQueryByAddr(t *testing.T) {
	total, txList, err := CommonTx{}.QueryByAddr("faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn", 0, 100)
	if err != nil {
		t.Error(err)
	}
	t.Logf("total: %v \n", total)

	for k, v := range txList {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryTxByPage(t *testing.T) {
	total, txList, err := CommonTx{}.QueryByPage(nil, 0, 20)

	if err != nil {
		t.Error(err)
	}
	t.Logf("total: %v \n", total)

	for k, v := range txList {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryHashActualFeeType(t *testing.T) {
	txList, err := CommonTx{}.QueryHashActualFeeType()
	if err != nil {
		t.Error(err)
	}

	for k, v := range txList {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestQueryTxByHash(t *testing.T) {

	tx, err := CommonTx{}.QueryTxByHash("D128AA72A3FF465A6292FED7C69507C5224655C86DA1EA67AE316479D07A1191")

	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %v \n", tx)

}

func TestCountByType(t *testing.T) {

	countByType, err := CommonTx{}.CountByType(bson.M{})

	if err != nil {
		t.Error(err)
	}
	t.Logf("count by type: %v  \n", countByType)
}

func TestGetTxlistByDuration(t *testing.T) {

	txList, err := CommonTx{}.GetTxlistByDuration(time.Now().Add(time.Hour*-240000).String(), time.Now().String())
	if err != nil {
		t.Error(err)
	}

	for k, v := range txList {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestGetTxCountByDuration(t *testing.T) {

	starttime,_ := time.Parse(utils.DateFmtYYYYMMDDHHmmss,"2019-07-01 00:00:00")
	endtime,_ := time.Parse(utils.DateFmtYYYYMMDDHHmmss,"2019-07-02 00:00:00")
	cnt,err := CommonTx{}.GetTxCountByDuration(starttime,endtime)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cnt)

}

func TestCommonTx_QueryTxAsset(t *testing.T) {
	total,ret,err := CommonTx{}.QueryTxAsset("gateway",Tx_Asset_TokenType_Edit,0,10)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(total)
	bytedata, _ := json.Marshal(ret)
	t.Log(string(bytedata))
}