package sync

import (
	"strings"
	"time"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"context"
	"encoding/hex"
	"github.com/tendermint/tendermint/types"
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"log"
)


func watch(c rpcclient.Client) {
	log.Printf("watched Transactions start")

	funcChain := []func(tx store.Docs){
		saveTx,saveOrUpdateAccount,updateAccountBalance,
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bk := make(chan interface{})

	c.Start()
	err := c.Subscribe(ctx, tools.SubscriberTx, types.EventQueryNewBlock, bk)

	if err != nil {
		log.Println("got ", err)
	}

	go func() {
		log.Println("listening block begin")
		setDelay(true) //延迟1秒钟
		for e := range bk {
			b, _ := m.QuerySyncTask()
			blockData := e.(types.TMEventData).Unwrap().(types.EventDataNewBlock)

			if len(blockData.Block.Data.Txs) > 0 {
				txs := blockData.Block.Data.Txs
				for _, txByte := range txs {

					txType, tx := tools.ParseTx(txByte)
					txHash := strings.ToUpper(hex.EncodeToString(txByte.Hash()))
					log.Printf("find tx,txType=%s;txHash%s", txType, txHash)
					if txType == "coin" {
						coinTx, _ := tx.(m.CoinTx)
						coinTx.TxHash = txHash
						coinTx.Height = blockData.Block.Height
						coinTx.Time = blockData.Block.Time
						handle(coinTx, funcChain)
					} else if txType == "stake" {
						stakeTx, _ := tx.(m.StakeTx)
						stakeTx.TxHash = txHash
						stakeTx.Height = blockData.Block.Height
						stakeTx.Time = blockData.Block.Time
						handle(stakeTx, funcChain)
					}
				}
			}
			//保存区块信息
			bk := m.Block{
				Height: blockData.Block.Height,
				Time:   blockData.Block.Time,
				TxNum:  int64(len(blockData.Block.Data.Txs)),
			}

			b.Height = blockData.Block.Height
			b.Time = blockData.Block.Time

			store.Update(b)
			store.SaveOrUpdate(bk)
		}
	}()
}
