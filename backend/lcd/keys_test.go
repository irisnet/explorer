package lcd

import (
	"encoding/json"
	"testing"
)

func TestAccount(t *testing.T) {
	address := "faa192vef4442d07lqde59mx35dvmfv9v72wrsu84a"

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