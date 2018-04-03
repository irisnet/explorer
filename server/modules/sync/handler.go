package sync

import (
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/rpc"
	"log"
)


var delay = false
func handle(txType string, tx store.Docs) {
	for _, fun := range handler[txType] {
		fun(tx)
	}
}
func setDelay(d bool){
	delay = d
}
var handler = map[string][]func(tx store.Docs){
	"coin":  {saveTx, handleCoinTx},
	"stake": {saveTx, handleStakeTx},
}

func saveTx(tx store.Docs) {
	store.Save(tx)
}

// 处理转账类交易
func handleCoinTx(tx store.Docs) {
	coinTx, _ := tx.(m.CoinTx)
	fun := func(address string) {
		account, err := m.QueryAccount(address)
		ac := rpc.QueryBalance(address, delay)
		if err == nil {
			account.Amount = ac.Coins
			account.Height = coinTx.Height
			account.Time = coinTx.Time
			if err := store.Update(account); err != nil {
				log.Printf("account:[%q] balance update failed,%s", account.Address, err)
			}
		} else {
			account = m.Account{
				Address: address,
				Amount:  ac.Coins,
				Time:    coinTx.Time,
				Height:  coinTx.Height,
			}
			if err := store.Save(account); err != nil {
				log.Printf("account:[%q] balance save failed,%s", account.Address, err)
			}

		}
	}
	fun(coinTx.From)
	fun(coinTx.To)

}

// 处理代理类交易
func handleStakeTx(tx store.Docs) {
	stakeTx, _ := tx.(m.StakeTx)
	fun := func(address string) {
		account, err := m.QueryAccount(address)
		//查询账户余额
		ac := rpc.QueryBalance(address, delay)

		if err == nil {
			account.Amount = ac.Coins
			if err := store.Update(account); err != nil {
				log.Printf("account:[%q] balance update failed,%s", account.Address, err)
			}
		} else {
			account = m.Account{
				Address: address,
				Amount:  ac.Coins,
				Time:    stakeTx.Time,
				Height:  stakeTx.Height,
			}
			if err := store.Save(account); err != nil {
				log.Printf("account:[%q] balance save failed,%s", account.Address, err)
			}
		}
	}
	fun(stakeTx.From)
}
