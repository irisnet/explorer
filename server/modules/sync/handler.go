package sync

import (
	"github.com/irisnet/irisplorer.io/server/modules/rest"
	"github.com/irisnet/irisplorer.io/server/modules/stake"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"log"
	"strings"
)

var delay = false

func setDelay(d bool) {
	delay = d
}

func handle(tx store.Docs, funChains []func(tx store.Docs)) {
	for _, fun := range funChains {
		fun(tx)
	}
}

func saveTx(tx store.Docs) {
	store.Save(tx)

	if tx.Name() == m.DocsNmStakeTx {
		stakeTx, _ := tx.(m.StakeTx)

		switch stakeTx.Type {
		case strings.TrimLeft(stake.TypeTxUnbond,"stake/"):
			de, err := m.QueryDelegatorByAddressAndPubkey(stakeTx.From, stakeTx.PubKey)
			if err != nil {
				de2, err2 := m.QueryCandidateByAddressAndPubkey(stakeTx.From, stakeTx.PubKey)
				if err2 != nil {
					log.Printf("error:delegator is lost,add = %s,pub_key=%s", stakeTx.From, stakeTx.PubKey)
					return
				}
				de2.Shares -= stakeTx.Amount.Amount
				store.Update(de)
				return
			}
			de.Shares -= stakeTx.Amount.Amount
			store.Update(de)
		case strings.TrimLeft(stake.TypeTxDelegate,"stake/"):
			de, err := m.QueryDelegatorByAddressAndPubkey(stakeTx.From, stakeTx.PubKey)
			if err != nil {
				de = m.Delegator{
					Address: stakeTx.From,
					PubKey:  stakeTx.PubKey,
				}
			}
			de.Shares += stakeTx.Amount.Amount
			store.SaveOrUpdate(de)
		case strings.TrimLeft(stake.TypeTxDeclareCandidacy,"stake/"):
			de, err := m.QueryCandidateByAddressAndPubkey(stakeTx.From, stakeTx.PubKey)
			if err != nil {
				de = m.Candidate{
					Address: stakeTx.From,
					PubKey:  stakeTx.PubKey,
				}
			}
			de.Shares += stakeTx.Amount.Amount
			store.SaveOrUpdate(de)
		}

	}
}

func saveOrUpdateAccount(tx store.Docs) {
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

func updateAccountBalance(tx store.Docs) {
	fun := func(address string) {
		account, _ := m.QueryAccount(address)
		//查询账户余额
		ac := rest.QueryBalance(address, delay)
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
		ac := rest.QueryBalance(account.Address, delay)
		account.Amount = ac.Coins
		if err := store.Update(account); err != nil {
			log.Printf("account:[%q] balance update failed,%s", account.Address, err)
		}
	}

}
