package document_test

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"testing"
)

func TestAsset_GetAllAssets(t *testing.T) {
	allAsset, err := document.AssetToken{}.GetAllAssets()
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range allAsset {
		t.Logf("k: %v \nv: %v\n", k, v)
	}
}
