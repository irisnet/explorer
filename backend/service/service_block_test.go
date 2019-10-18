package service

import (
	"testing"
)

func TestGetValidatorSet(t *testing.T) {

	validatorPage := new(BlockService).GetValidatorSet(1000, 0, 100)

	t.Logf("total: %v \n", validatorPage.Total)

	for k, v := range validatorPage.Items {
		t.Logf("k: %v   v: %v \n", k, v)
	}
}

func TestQueryBlockInfo(t *testing.T) {

	blockInfo := new(BlockService).QueryBlockInfo(100)
	t.Logf("block info: %v \n", blockInfo)

}

func TestQueryRecent(t *testing.T) {

	blockList := new(BlockService).QueryRecent()

	for k, v := range blockList {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestBlockService_QueryLatestHeight(t *testing.T) {
	block := new(BlockService).QueryLatestHeight()
	t.Log(block)
}