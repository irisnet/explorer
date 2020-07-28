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
	total, txList, err := CommonTx{}.QueryByPage(bson.M{"$or": []bson.M{{"signers.addr_bech32": "iaa1x98k5n7xj0h3udnf5dcdzw85tsfa75qm0kqak0"},
		{"to": "iaa1x98k5n7xj0h3udnf5dcdzw85tsfa75qm0kqak0"}}}, 1, 10, true)

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
	tx, err := CommonTx{}.QueryTxByHash("3EB9D229189C84A0A11FC4A19154FEE5CDC0E2C10E6695E316F6888BE2677C3C")

	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v \n", string(utils.MarshalJsonIgnoreErr(tx)))

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

	starttime, _ := time.Parse(utils.DateFmtYYYYMMDDHHmmss, "2019-07-01 00:00:00")
	endtime, _ := time.Parse(utils.DateFmtYYYYMMDDHHmmss, "2019-07-02 00:00:00")
	cnt, err := CommonTx{}.GetTxCountByDuration(starttime, endtime)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cnt)

}

func TestCommonTx_QueryTxAsset(t *testing.T) {
	total, ret, err := CommonTx{}.QueryTxAsset("native", "MintToken", "h002", "", 0, 10, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(total)
	bytedata, _ := json.Marshal(ret)
	t.Log(string(bytedata))
}

func TestCommonTx_QueryTxTransferGatewayOwner(t *testing.T) {
	total, ret, err := CommonTx{}.QueryTxTransferGatewayOwner("shark", 0, 10, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(total)
	bytedata, _ := json.Marshal(ret)
	t.Log(string(bytedata))
}

func TestCommonTx_QueryHashTimeByProposalIdVoters(t *testing.T) {
	txList, err := CommonTx{}.QueryHashTimeByProposalIdVoters(44, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(len(txList))
	bytedata, _ := json.Marshal(txList)
	t.Log(string(bytedata))
}