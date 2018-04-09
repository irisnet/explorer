package sync

import (
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/rpc"
	"log"
)


var delay = false
func setDelay(d bool){
	delay = d
}

func saveTx(tx store.Docs) {
	store.Save(tx)
}

func handle(tx store.Docs,funChains []func(tx store.Docs)){
	for _,fun := range funChains {
		fun(tx)
	}
}

func saveOrUpdateAccount(tx store.Docs){
	switch tx.Name() {
	case m.DocsNmCoinTx:
		coinTx, _ := tx.(m.CoinTx)
		fun := func(address string) {
			account := m.Account{
				Address: address,
				Time:    coinTx.Time,
				Height:  coinTx.Height,
			}

			if err := store.SaveOrUpdate(account); err != nil {
				log.Printf("account:[%q] balance update failed,%s", account.Address, err)
			}
		}
		fun(coinTx.From)
		fun(coinTx.To)
	case m.DocsNmStakeTx:
		stakeTx, _ := tx.(m.StakeTx)
		fun := func(address string) {
			account := m.Account{
				Address: address,
				Time:    stakeTx.Time,
				Height:  stakeTx.Height,
			}

			if err := store.SaveOrUpdate(account); err != nil {
				log.Printf("account:[%q] balance update failed,%s", account.Address, err)
			}
		}
		fun(stakeTx.From)
	}
}

func updateAccountBalance(tx store.Docs){
	fun := func(address string) {
		account, _ := m.QueryAccount(address)
		//查询账户余额
		ac := rpc.QueryBalance(address, delay)
		account.Amount = ac.Coins
		if err := store.Update(account); err != nil {
			log.Printf("account:[%q] balance update failed,%s", account.Address, err)
		}
	}
	switch tx.Name() {
	case m.DocsNmCoinTx:
		coinTx, _ := tx.(m.CoinTx)
		fun(coinTx.From)
		fun(coinTx.To)
	case m.DocsNmStakeTx:
		stakeTx, _ := tx.(m.StakeTx)
		fun(stakeTx.From)
	case m.DocsNmAccount:
		account, _ := tx.(m.Account)
		ac := rpc.QueryBalance(account.Address, delay)
		account.Amount = ac.Coins
		if err := store.Update(account); err != nil {
			log.Printf("account:[%q] balance update failed,%s", account.Address, err)
		}
	}

}
