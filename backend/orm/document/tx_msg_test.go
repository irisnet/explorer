package document

import (
	"testing"
)

func TestQueryTxMsgListByHashList(t *testing.T) {

	txMsgList, err := TxMsg{}.QueryTxMsgListByHashList([]string{"74C018BF53EC501A9B00A0403A0584CB6FE382D1B53A26F0A856A692A891620F"})
	if err != nil {
		t.Error(err)
	}

	for k, v := range txMsgList {
		t.Logf("k: %v v: %v \n", k, v)
	}
}

func TestQueryTxMsgByHash(t *testing.T) {
	txMsg, err := TxMsg{}.QueryTxMsgByHash("74C018BF53EC501A9B00A0403A0584CB6FE382D1B53A26F0A856A692A891620F")

	if err != nil {
		t.Error(err)
	}

	t.Logf("txMsg: %v", txMsg)
}
