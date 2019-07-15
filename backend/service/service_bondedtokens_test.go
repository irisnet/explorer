package service

import (
	"testing"
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
)

func TestBondedTokensService_QueryBondedTokensValidator(t *testing.T) {
	res, err := (&BondedTokensService{}).QueryBondedTokensValidator(types.RoleValidator)
	if err != nil {
		t.Fatal(err)
	}
	byteData,_ := json.Marshal(res)
	t.Log(string(byteData))
}
