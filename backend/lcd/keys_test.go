package lcd

import (
	"encoding/json"
	"testing"
)

func TestAccount(t *testing.T) {
	address := "iaa18e2e9fxxrr88k78gg7fhuuqgccfv8self9ye65"

	if res, err := Account(address); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.MarshalIndent(res, "", "\t")
		t.Log(string(resBytes))
	}
}

func TestAccountInfo(t *testing.T) {
	address := "iaa18e2e9fxxrr88k78gg7fhuuqgccfv8self9ye65"

	if res, err := AccountInfo(address); err != nil {
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
