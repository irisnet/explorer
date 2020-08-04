package service

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/orm/document"
	"testing"
)

func TestAssetsService_GetGatewayAsset(t *testing.T) {

	//ret, err := assetsService.GetGatewayAsset(document.Tx_Asset_TxType_Issue, "", "", 0, 10, true)
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//byteret, _ := json.Marshal(ret)
	//t.Log(string(byteret))
}

func TestAssetsService_GetNativeAsset(t *testing.T) {
	ret, err := assetsService.GetNativeAsset(document.Tx_Asset_TxType_Edit, "", 0, 10, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	byteret, _ := json.Marshal(ret)
	t.Log(string(byteret))
}

func TestAssetsServiceisFieldTokenType(t *testing.T) {
	if isFieldTokenType(document.Tx_Asset_TxType_Edit) {
		t.Log("ok")
	} else {
		t.Failed()
	}

}

func TestAssetService_QueryAssetTokens(t *testing.T) {
	res, _ := assetsService.QueryAssetTokens()

	for k, v := range res {
		t.Logf("k: %v \nv: %v\n", k, v)
	}
}

func TestAssetsService_QueryAssetTokenDetail(t *testing.T) {
	tokenid := "skitty"
	res, _ := assetsService.QueryAssetTokenDetail(tokenid)

	bytesmsg, _ := json.Marshal(res)
	t.Log(string(bytesmsg))
}

func TestAssetService_UpdateAssetTokens(t *testing.T) {
	//err := assetsService.UpdateAssetTokens(dst)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
}
