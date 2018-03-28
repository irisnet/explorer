package sync

import (
	"github.com/irisnet/iris-explorer/modules/store"
	"github.com/irisnet/iris-explorer/modules/rpc"
	"log"
)

func handle(txType string, tx store.DocsHander){
	for _,fun := range handler[txType]{
		fun(tx)
	}
}

var handler = map[string][]func(tx store.DocsHander){
	"coin": {saveTx,handleCoinTx},
	"stake":{saveTx,handleStakeTx},
}

func saveTx(tx store.DocsHander){
	store.Save(tx)
}
// 处理转账类交易
func handleCoinTx(tx store.DocsHander){
	coinTx, _ := tx.(store.CoinTx)
	fun := func(address string) {
		account,err := store.QueryAccount(address)
		ac := rpc.QueryBalance(address,true)
		if err == nil {
			account.Amount = ac.Coins
			account.Height = coinTx.Height
			account.Time = coinTx.Time
			if err := store.Update(account); err != nil{
				log.Printf("account:[%q] balance update failed,%s",account.Address,err)
			}
		}else {
			account = store.Account{
				Address:address,
				Amount:ac.Coins,
				Time:coinTx.Time,
				Height:coinTx.Height,
			}
			if err := store.Save(account); err != nil{
				log.Printf("account:[%q] balance save failed,%s",account.Address,err)
			}

		}
	}
	fun(coinTx.From)
	fun(coinTx.To)

}

// 处理代理类交易
func handleStakeTx(tx store.DocsHander){
	stakeTx, _ := tx.(store.StakeTx)
	fun := func(address string) {
		account,err := store.QueryAccount(address)
		//查询账户余额
		ac := rpc.QueryBalance(address,true)

		if err == nil {
			account.Amount = ac.Coins
			if err := store.Update(account); err != nil{
				log.Printf("account:[%q] balance update failed,%s",account.Address,err)
			}
		}else {
			account = store.Account{
				Address:address,
				Amount:ac.Coins,
				Time:stakeTx.Time,
				Height:stakeTx.Height,
			}
			if err := store.Save(account); err != nil{
				log.Printf("account:[%q] balance save failed,%s",account.Address,err)
			}
		}
	}
	fun(stakeTx.From)
}
