package sync

import (
	"encoding/hex"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"log"
	"strings"
)

func InitServer() {
	store.Init()

	chainId := viper.GetString(tools.ChainId)
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
		watchBlock(c)
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

func watchBlock(c rpcclient.Client) error {
	b, _ := m.QuerySyncTask()
	status, _ := c.Status()
	latestBlockHeight := status.LatestBlockHeight

	funcChain := []func(tx store.Docs){
		saveTx, saveOrUpdateAccount, updateAccountBalance,
	}

	ch := make(chan int64)

	go syncBlock(b.Height+1, latestBlockHeight, funcChain, ch, 0)
	select {
	case <-ch:
		//更新同步记录
		block, _ := c.Block(&latestBlockHeight)
		b.Height = block.Block.Height
		b.Time = block.Block.Time
		return store.Update(b)
	}
}

var maxBatchNum = int64(10000)
var minBatchSize = int64(20)

func fastSync(c rpcclient.Client) error {
	b, _ := m.QuerySyncTask()
	status, _ := c.Status()
	latestBlockHeight := status.LatestBlockHeight
	funcChain := []func(tx store.Docs){
		saveTx, saveOrUpdateAccount,
	}
	ch := make(chan int64)
	threadNumAdd := int64(0)

	threadNum := (latestBlockHeight - b.Height) / maxBatchNum
	//单线程处理
	if threadNum == 0 {
		go syncBlock(b.Height, latestBlockHeight, funcChain, ch, 0)
	} else {
		//开启多线程处理
		for i := int64(1); i <= threadNum; i++ {
			threadNumAdd += i
			var start = b.Height + (i-1)*maxBatchNum + 1
			var end = b.Height + i*maxBatchNum
			if i == threadNum {
				end = latestBlockHeight
			}

			go syncBlock(start, end, funcChain, ch, i)
		}

	}

	for {
		select {
		case threadNo := <-ch:
			log.Printf("threadNo[%d] is over", threadNo)
			threadNumAdd -= threadNo
			if threadNumAdd == 0 {
				goto end
			}
		}
	}

end:
	{
		//更新同步记录
		block, _ := c.Block(&latestBlockHeight)
		b.Height = block.Block.Height
		b.Time = block.Block.Time
		store.Update(b)

		//同步账户余额
		accounts := m.QueryAll()
		for _, account := range accounts {
			updateAccountBalance(account)
		}
		log.Printf("update account balance over")

		return nil
	}
}
func syncBlock(start int64, end int64, funcChain []func(tx store.Docs), ch chan int64, threadNum int64) {
	log.Printf("threadNo[%d] begin sync block from %d to %d", threadNum, start, end)
	node := tools.GetNode()
	defer node.Release()

	for j := start; j <= end; j++ {
		log.Printf("===========threadNo[%d] sync block,height:%d===========", threadNum, j)

		//TODO 使用node.Client.BlockchainInfo
		block, err := node.Client.Block(&j)
		if err != nil {
			//重新尝试一次
			node2 := tools.GetNode()
			block, err = node2.Client.Block(&j)
			if err != nil {
				log.Fatalf("invalide block height %d", j)
			}
		}
		if block.BlockMeta.Header.NumTxs > 0 {
			txs := block.Block.Data.Txs
			for _, txByte := range txs {
				txType, tx := tools.ParseTx(txByte)
				txHash := strings.ToUpper(hex.EncodeToString(txByte.Hash()))
				log.Printf("===========threadNo[%d] find tx,txType=%s;txHash=%s", threadNum, txType, txHash)
				if txType == "coin" {
					coinTx, _ := tx.(m.CoinTx)
					coinTx.TxHash = txHash
					coinTx.Height = block.Block.Height
					coinTx.Time = block.Block.Time
					handle(coinTx, funcChain)
				} else if txType == "stake" {
					stakeTx, _ := tx.(m.StakeTx)
					stakeTx.TxHash = txHash
					stakeTx.Height = block.Block.Height
					stakeTx.Time = block.Block.Time
					handle(stakeTx, funcChain)
				}
			}
		}

		//保存区块信息
		bk := m.Block{
			Height: block.Block.Height,
			Time:   block.Block.Time,
			TxNum:  block.BlockMeta.Header.NumTxs,
		}
		store.SaveOrUpdate(bk)

	}
	ch <- threadNum
}
