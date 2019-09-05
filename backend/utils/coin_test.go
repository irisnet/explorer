package utils

import "testing"

func TestCovertAssetUnit(t *testing.T) {
	supply := "1000000000000000000"
	data := CovertAssetUnit(supply, 2)
	t.Log(supply)
	t.Log(data)
}
