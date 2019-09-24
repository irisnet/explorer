package service

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/orm/document"
	"testing"
)

func TestAssetsService_GetGatewayAsset(t *testing.T) {

	ret, err := (&AssetsService{}).GetGatewayAsset(document.Tx_Asset_TxType_Issue, "", "", 0, 10, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	byteret, _ := json.Marshal(ret)
	t.Log(string(byteret))
}

func TestAssetsService_GetNativeAsset(t *testing.T) {
	ret, err := (&AssetsService{}).GetNativeAsset(document.Tx_Asset_TxType_Edit, "", 0, 10, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	byteret, _ := json.Marshal(ret)
	t.Log(string(byteret))
}

//func TestLoadModelFromCommonTx(t *testing.T) {
//	txcommon := document.CommonTx{
//		Height:880,
//		TxHash:"9E7BF9B28DE7156627439651D46AB41DBA7BA41E1E76737F2BA62333C5E9B824",
//		Fee:document.Fee{
//			Amount:document.Coins{
//				document.Coin{Denom:utils.CoinTypeAtto,Amount:327542754275.32423424000},
//			},
//			Gas:200000000,
//		},
//		Amount:document.Coins{
//			document.Coin{Denom:utils.CoinTypeAtto,Amount:327542754275.32423424000},
//		},
//		Time:time.Now(),
//		Status:"success",
//		Msgs:[]document.MsgItem{
//			{Type: document.Tx_Asset_TxType_Issue,
//	         MsgData:document.MsgData{
//	         	TokenId:          "",
//				 MinUnitAlias:    "",
//	         	SrcOwner:         "",
//	         	Symbol:           "lc",
//	         	Source:           "native",
//				 CanonicalSymbol: "",
//	         	Amount:           0,
//	         	DstOwner:         "",
//	         	Decimal:          3333,
//	         	Name:             "lc",
//	         	Mintable:         false,
//	         	MaxSupply:        10000000000,
//	         	InitialSupply:    100000000,
//	         	Gateway:          "",
//	         	Owner:            "faa1f7e8cfc7cnny7vxtshe6tc0mtqhg0k64wg3e65",
//			 },
//	        },
//		},
//	}
//	assets := LoadModelFromCommonTx(txcommon)
//	byteret ,_ := json.Marshal(assets)
//	t.Log(string(byteret))
//}
func TestAssetsServiceisFieldTokenType(t *testing.T) {
	if isFieldTokenType(document.Tx_Asset_TxType_Edit) {
		t.Log("ok")
	} else {
		t.Failed()
	}

}

func TestAssetService_QueryAssetTokens(t *testing.T) {
	source := "gateway"
	res, _ := new(AssetsService).QueryAssetTokens(source)

	for k, v := range res {
		t.Logf("k: %v \nv: %v\n", k, v)
	}
}

func TestAssetsService_QueryAssetTokenDetail(t *testing.T) {
	tokenid := "shark.fly"
	res, _ := new(AssetsService).QueryAssetTokenDetail(tokenid)

	bytesmsg, _ := json.Marshal(res)
	t.Log(string(bytesmsg))
}

func TestAssetService_UpdateAssetTokens(t *testing.T) {
	dst := []document.AssetToken{}
	err := new(AssetsService).UpdateAssetTokens(dst)

	if err != nil {
		t.Fatal(err)
	}
}
