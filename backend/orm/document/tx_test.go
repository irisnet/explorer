package document

import (
	"testing"
	"time"

	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

func TestQueryByAddr(t *testing.T) {
	total, txList, err := CommonTx{}.QueryByAddr("faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn", 0, 100, true)
	if err != nil {
		t.Error(err)
	}
	t.Logf("total: %v \n", total)

	for k, v := range txList {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryTxByPage(t *testing.T) {
	total, txList, err := CommonTx{}.QueryByPage(nil, 0, 20, true)

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
	tx, err := CommonTx{}.QueryTxByHash("11E414CD9DAEEA1CF8BA11B862C643D801EBB2E3BCE704A9AE8EFA47EFA8BD0F")

	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v \n", tx)

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
	total,ret,err := CommonTx{}.QueryTxAsset("gateway","",0,10, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(total)
	bytedata, _ := json.Marshal(ret)
	t.Log(string(bytedata))
}