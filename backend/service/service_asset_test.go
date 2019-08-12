package service_test

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"testing"
)

func TestAssetService_QueryAssetToken(t *testing.T) {
	res := new(service.AssetService).QueryAssetToken()

	for k, v := range res {
		t.Logf("k: %v \nv: %v\n", k, v)
	}
}

func TestAssetService_UpdateAssetTokens(t *testing.T) {
	dst := []document.Asset{}
	err := new(service.AssetService).UpdateAssetTokens(dst)

	if err != nil {
		t.Fatal(err)
	}
}
