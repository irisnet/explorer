package service

import (
	"testing"
)

func TestQueryDelegation(t *testing.T) {

	selfBond, otherBond := new(DelegatorService).QueryDelegation("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx")

	t.Logf("selfBond: %v  otherBond: %v \n", selfBond, otherBond)

}

func TestGetDeposits(t *testing.T) {

	depositCoin := new(DelegatorService).GetDeposits("faa1k4xyra6qac7dwgngq8t6w8ra2qa9hj95mwth42")

	t.Logf("deposit Coin: %v \n", depositCoin)
}
