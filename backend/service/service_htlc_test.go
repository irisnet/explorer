package service

import (
	"testing"
	"encoding/json"
)

func TestHtlcService_QueryHtlcByHashLock(t *testing.T) {
	htlcdata := new(HtlcService).QueryHtlcByHashLock("d99c06c3531c6751d8ecfebe33ea4eafc09a5acbbb691b5fe35a9a1eaec21278")
	bytedata, _ := json.Marshal(htlcdata)
	t.Log(string(bytedata))
}
