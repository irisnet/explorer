package lcd_test

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/lcd"
	"testing"
)

func TestGetAssetTokens(t *testing.T) {
	res := lcd.GetAssetTokens()

	bytesData, _ := json.Marshal(res)
	t.Log(string(bytesData))
}

//func TestGetAssetGateways(t *testing.T) {
//	res := lcd.GetAssetGateways()
//
//	bytesData, _ := json.Marshal(res)
//	t.Log(string(bytesData))
//}