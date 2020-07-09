package lcd

import (
	"encoding/json"
	"testing"
)

func TestAccount(t *testing.T) {
	address := "faa1dmnnl50sagq2f6x88mz62rcn2mv5pfe53madlz"

	if res, err := Account(address); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.MarshalIndent(res, "", "\t")
		t.Log(string(resBytes))
	}
}

func TestGetIconsByKey(t *testing.T) {
	pic, err := GetIconsByKey("8E15046A255C20CC")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pic)
}