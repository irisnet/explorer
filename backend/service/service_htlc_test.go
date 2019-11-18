package service

import (
	"testing"
	"encoding/json"
)

func TestHtlcService_QueryHtlcByHashLock(t *testing.T) {
	htlcdata := new(HtlcService).QueryHtlcByHashLock("55df15b5d46e1ae185561b93b3395a1a787f7d45c04a1724ed0409599f11bca8")
	bytedata, _ := json.Marshal(htlcdata)
	t.Log(string(bytedata))
}
