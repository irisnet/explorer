package service

import (
	"testing"
)

func TestGetValidatorSet(t *testing.T) {

	validatorPage := blockService.GetValidatorSet(1000, 0, 100)

	t.Logf("total: %v \n", validatorPage.Total)

	for k, v := range validatorPage.Items {
		t.Logf("k: %v   v: %v \n", k, v)
	}
}

func TestQueryBlockInfo(t *testing.T) {

	blockInfo := blockService.QueryBlockInfo(100)
	t.Logf("block info: %v \n", blockInfo)

}

func TestQueryRecent(t *testing.T) {

	blockList := blockService.QueryRecent()

	for k, v := range blockList {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestBlockService_QueryLatestHeight(t *testing.T) {
	block := blockService.QueryLatestHeight()
	t.Log(block)
}