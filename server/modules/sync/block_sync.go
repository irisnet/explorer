package sync

import (
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"strings"
	"encoding/hex"
	"log"
	"github.com/spf13/viper"
	"github.com/robfig/cron"
)

func InitServer() {
	url := viper.GetString(tools.MgoUrl)
	chainId := viper.GetString(tools.ChainId)
	store.Init(url)
	block, err := m.QuerySyncTask()

	if err != nil {
		if chainId == "" {
			log.Fatalf("sync process start failed,chainId is empty")
		}
		//初始化配置表
		block = m.SyncTask{
			Height:  0,
			ChainID: chainId,
		}
		store.Save(block)
	}
	//初始化连接池
	tools.Init()
}

func startCron(c rpcclient.Client) {
	spec := viper.GetString(tools.SyncCron)
	cron := cron.New()
	cron.AddFunc(spec, func() {
		fastSync(c)
	})
	go cron.Start()
}

func Start() {
	InitServer()
	c := tools.GetNode().Client
	if err := fastSync(c); err != nil {
		log.Fatalf("sync block failed,%v", err)
	}
	startCron(c)
	//go watch(c) 监控的方式在启动同步过程中容易丢失区块
}

func fastSync(c rpcclient.Client) error {
	b, _ := m.QuerySyncTask()
	status, _ := c.Status()
	latestBlockHeight := status.LatestBlockHeight
	for i := b.Height + 1; i <= latestBlockHeight; i++ {
		log.Printf("===========sync block,height:%d===========", i)
		block, _ := c.Block(&i)
		if block.BlockMeta.Header.NumTxs > 0 {
			txs := block.Block.Data.Txs
			for _, txByte := range txs {
				txType, tx := tools.ParseTx(txByte)
				txHash := strings.ToUpper(hex.EncodeToString(txByte.Hash()))
				log.Printf("find tx,txType=%s;txHash=%s", txType, txHash)
				if txType == "coin" {
					coinTx, _ := tx.(m.CoinTx)
					coinTx.TxHash = txHash
					coinTx.Height = block.Block.Height
					coinTx.Time = block.Block.Time
					handle(txType, coinTx)
				} else if txType == "stake" {
					stakeTx, _ := tx.(m.StakeTx)
					stakeTx.TxHash = txHash
					stakeTx.Height = block.Block.Height
					stakeTx.Time = block.Block.Time
					handle(txType, stakeTx)
				}
			}
		}
		b.Height = block.Block.Height
		b.Time = block.Block.Time

		//保存区块信息
		bk := m.Block{
			Height: b.Height,
			Time:   b.Time,
			TxNum:  block.BlockMeta.Header.NumTxs,
		}
		store.SaveOrUpdate(bk)
	}
	return store.Update(b)
}
