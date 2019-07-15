package service

import (
	"testing"
	"encoding/json"
)

func TestBondedTokensService_QueryBondedTokensValidator(t *testing.T) {
	res,err := (&BondedTokensService{}).QueryBondedTokensValidator()
	if err != nil {
		t.Fatal(err)
	}
	byteData,_ := json.Marshal(res)
	t.Log(string(byteData))
}
