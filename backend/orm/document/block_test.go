package document

import (
	"testing"
	"time"
)

func TestGetBlockListByOffsetAndSize(t *testing.T) {

	blocks, err := Block{}.GetBlockListByOffsetAndSize(100, 50)

	if err != nil {
		t.Error(err)
	}

	for k, v := range blocks {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryBlockHeightTimeHashByHeight(t *testing.T) {

	block, err := Block{}.QueryBlockHeightTimeHashByHeight(100)
	if err != nil {
		t.Error(err)
	}

	t.Logf("height: 100 block: %v\n", block)
}

func TestGetBlockListByPage(t *testing.T) {

	total, blockPage, err := Block{}.GetBlockListByPage(100, 10, true)
	if err != nil {
		t.Error(err)
	}

	t.Logf("total: %v \n", total)

	for k, v := range blockPage {
		t.Logf("k: %v  v: %v\n", k, v)
	}

}

func TestGetRecentBlockList(t *testing.T) {

	blockList, err := Block{}.GetRecentBlockList()
	if err != nil {
		t.Error(err)
	}

	for k, v := range blockList {
		t.Logf("k: %v  v: %v\n", k, v)
	}

}

func TestQueryOneBlockOrderByHeightAsc(t *testing.T) {

	block, err := Block{}.QueryOneBlockOrderByHeightAsc()

	if err != nil {
		t.Error(err)
	}

	t.Logf("block:  %v \n", block)

}

func TestQueryOneBlockOrderByHeightDesc(t *testing.T) {

	block, err := Block{}.QueryOneBlockOrderByHeightDesc()

	if err != nil {
		t.Error(err)
	}

	t.Logf("block:  %v \n", block)

}

func TestQueryBlocksByDurationWithHeightAsc(t *testing.T) {

	blocks, err := Block{}.QueryBlocksByDurationWithHeightAsc(time.Now().Add(time.Hour*(-10)), time.Now().Add(time.Hour*(-9)))

	if err != nil {
		t.Error(err)
	}

	for k, v := range blocks {
		t.Logf("k: %v  v: %v\n", k, v)
	}

}
