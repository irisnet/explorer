package lcd

import (
	"testing"
	"encoding/json"
)

func TestHtlcInfo(t *testing.T) {
	data, err := HtlcInfo("d99c06c3531c6751d8ecfebe33ea4eafc09a5acbbb691b5fe35a9a1eaec21278")
	if err != nil {
		t.Fatal(err)
	}

	msgbyte, _ := json.Marshal(data)
	t.Log(string(msgbyte))
}
