package store

import (
	"testing"
	"fmt"
	"time"
	"github.com/cosmos/cosmos-sdk/modules/coin"
)

func TestSaveCoinTxs(t *testing.T) {
	Init("localhost:27017")

	tx := CoinTx{
		TxHash:"3DFA30A625DB0B4EF7138D0655156223AF29DFA6",
		Time:time.Now(),
		Height:1020814,
		From:"F75D0FD7036A3BB9092DD0938BBF0A9323224832",
		To:"37703E80926E2977167BB073C100E98C511FDCBC",
		Amount:[]coin.Coin{
			coin.Coin{
				Amount:100,
				Denom:"fermion",
			},
		},
	}
	Save(tx)

}

func TestQueryCoinTxsByFrom(t *testing.T) {
	Init("localhost:27017")
	result := QueryAllCoinTx()
	fmt.Printf("size = %d",len(result))

}
