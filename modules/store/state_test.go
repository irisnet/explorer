package store

import (
	"testing"
	//"time"
	//"github.com/cosmos/cosmos-sdk/modules/coin"
	"fmt"
)

func TestSaveCoinTxs(t *testing.T) {
	Mgo.Init("localhost:27017")

	//tx := CoinTx{
	//	TxHash:"3DFA30A625DB0B4EF7138D0655156223AF29DFA6",
	//	Time:time.Now(),
	//	Height:1020814,
	//	From:"F75D0FD7036A3BB9092DD0938BBF0A9323224832",
	//	To:"37703E80926E2977167BB073C100E98C511FDCBC",
	//	Amount:[]coin.Coin{
	//		coin.Coin{
	//			Amount:100,
	//			Denom:"fermion",
	//		},
	//	},
	//}
	tx := SyncBlock{
		CurrentPos:1,
		TotalCoinTxs:0,
		TotalStakeTxs:0,
	}
	Mgo.Save(tx)

}

func TestQueryCoinTxsByFrom(t *testing.T) {
	Mgo.Init("localhost:27017")
	restlt := Mgo.QueryCoinTxsByAccount("0D7ACAD5C3F3EE3DBFB972F52D652509437E0044")
	fmt.Println("%s",restlt)
	////var s = fmt.Sprintf("%s",[]byte("D5EF35C78C86F8750FDAE1FED11AE8FA811E9095"))
	////fmt.Println("%s",s)
	//block,err := Mgo.QueryLastedBlock()
	//fmt.Println("%s",block.CurrentPos)
}
