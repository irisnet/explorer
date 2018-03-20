package sync

import (
	"fmt"
	"log"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/cosmos/cosmos-sdk/modules/nonce"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/tendermint/go-wire/data"
	"github.com/irisnet/iris-explorer/modules/store"
	"github.com/irisnet/iris-explorer/modules/stake"
	"github.com/irisnet/iris-explorer/modules/tools"
	"time"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"context"
	"encoding/hex"
	"github.com/tendermint/tendermint/types"
	"github.com/spf13/cast"
)

const (
	MgoUrl = "mgo-url"
)

var (
	SyncCmd = &cobra.Command{
		Use:  "sync",
		Long: `sync iris-hub tx data to mongodb`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartWatch()
		},
	}
)

func processSync(c rpcclient.Client) {
	url := viper.GetString(MgoUrl)
	store.Mgo.Init(url)
	block, err := store.Mgo.QueryLastedBlock()

	if err != nil {
		//初始化配置表
		tx := store.SyncBlock{
			CurrentPos:    1,
			TotalCoinTxs:  0,
			TotalStakeTxs: 0,
		}
		store.Mgo.Save(tx)
		block = tx
	}

	//开始漏单查询
	sync(block, c)
}

func StartWatch() error {
	c := tools.GetNode().Client
	processSync(c)
	processWatch(c)
	return nil
}

func processWatch(c rpcclient.Client) {
	log.Printf("watched Transactions start")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	txs := make(chan interface{})

	c.Start()
	err := c.Subscribe(ctx, "iris-explorer", types.EventQueryTx, txs)

	if err != nil {
		log.Println("got ", err)
	}

	go func() {
		log.Println("listening tx begin")
		for e := range txs {
			block, _ := store.Mgo.QueryLastedBlock()
			deliverTxRes := e.(types.TMEventData).Unwrap().(types.EventDataTx)
			height := deliverTxRes.Height

			txb, _ := sdk.LoadTx(deliverTxRes.Tx)
			txType, tx := parseTx(txb)
			if txType == "coin" {
				coinTx, _ := tx.(store.CoinTx)

				coinTx.TxHash = strings.ToUpper(hex.EncodeToString(deliverTxRes.Tx.Hash()))
				coinTx.Time = queryBlockTime(c, height)
				coinTx.Height = height
				if store.Mgo.Save(coinTx) != nil {
					break
				}
				block.TotalCoinTxs += 1

				log.Printf("watched coinTx,tx_hash=%s", coinTx.TxHash)
			} else if txType == "stake" {
				stakeTx, _ := tx.(store.StakeTx)
				stakeTx.TxHash = strings.ToUpper(hex.EncodeToString(deliverTxRes.Tx.Hash()))
				stakeTx.Time = queryBlockTime(c, height)
				stakeTx.Height = height
				if store.Mgo.Save(stakeTx) != nil {
					break
				}
				block.TotalStakeTxs += 1
				log.Printf("watched stakeTx,tx_hash=%s", stakeTx.TxHash)
			}
			block.CurrentPos = height
			store.Mgo.UpdateBlock(block)
		}
	}()
}

func parseTx(tx sdk.Tx) (string, interface{}) {
	txl, ok := tx.Unwrap().(sdk.TxLayer)
	var txi sdk.Tx
	var coinTx store.CoinTx
	var stakeTx store.StakeTx
	var nonceAddr data.Bytes
	for ok {
		txi = txl.Next()
		switch txi.Unwrap().(type) {
		case coin.SendTx:
			ctx, _ := txi.Unwrap().(coin.SendTx)
			coinTx.From = fmt.Sprintf("%s", ctx.Inputs[0].Address.Address)
			coinTx.To = fmt.Sprintf("%s", ctx.Outputs[0].Address.Address)
			coinTx.Amount = ctx.Inputs[0].Coins
			return "coin", coinTx
		case nonce.Tx:
			ctx, _ := txi.Unwrap().(nonce.Tx)
			nonceAddr = ctx.Signers[0].Address
			break
		case stake.TxUnbond, stake.TxDelegate, stake.TxDeclareCandidacy:
			kind, _ := txi.GetKind()
			stakeTx.From = fmt.Sprintf("%s", nonceAddr)
			stakeTx.Type = strings.Replace(kind, "stake/", "", -1)
			switch kind {
			case "stake/unbond":
				ctx, _ := txi.Unwrap().(stake.TxUnbond)
				stakeTx.Amount.Denom = "fermion"
				stakeTx.Amount.Amount = int64(ctx.Shares)
				break
			case "stake/delegate":
				ctx, _ := txi.Unwrap().(stake.TxDelegate)
				stakeTx.Amount.Denom = ctx.Bond.Denom
				stakeTx.Amount.Amount = ctx.Bond.Amount
				break
			case "stake/declareCandidacy":
				ctx, _ := txi.Unwrap().(stake.TxDeclareCandidacy)
				stakeTx.Amount.Denom = ctx.BondUpdate.Bond.Denom
				stakeTx.Amount.Amount = ctx.BondUpdate.Bond.Amount
				break
			}
			return "stake", stakeTx
		}
		txl, ok = txi.Unwrap().(sdk.TxLayer)
	}
	return "", nil
}

func queryBlockTime(c rpcclient.Client, height int64) time.Time {
	h := cast.ToInt64(height)
	block, err := c.Block(&h)
	if err != nil {
		log.Printf("query block fail ,%d", height)
	}
	return block.BlockMeta.Header.Time
}

func sync(curBlock store.SyncBlock, c rpcclient.Client) {
	log.Printf("sync Transactions start")

	current := curBlock.CurrentPos
	latest := int64(0)

	log.Printf("last block heigth：%d", current)

	for ok := true; ok; ok = current < latest {
		blocks, err := c.BlockchainInfo(current, current+20)
		if err != nil {
			log.Fatal(err)
		}
		for _, block := range blocks.BlockMetas {
			if block.Header.NumTxs > 0 {
				txhash := block.Header.DataHash
				prove := !viper.GetBool(commands.FlagTrustNode)
				res, _ := c.Tx(txhash, prove)
				txs, _ := sdk.LoadTx(res.Proof.Data)
				txType, tx := parseTx(txs)
				if txType == "coin" {
					coinTx, _ := tx.(store.CoinTx)
					coinTx.TxHash = strings.ToUpper(fmt.Sprintf("%s", txhash))
					coinTx.Time = block.Header.Time
					coinTx.Height = block.Header.Height
					if store.Mgo.Save(coinTx) == nil {
						curBlock.TotalCoinTxs += 1
						log.Printf("sync coinTx,tx_hash=%s", coinTx.TxHash)
					}

				} else if txType == "stake" {
					stakeTx, _ := tx.(store.StakeTx)
					stakeTx.TxHash = strings.ToUpper(fmt.Sprintf("%s", txhash))
					stakeTx.Time = block.Header.Time
					stakeTx.Height = block.Header.Height
					if store.Mgo.Save(stakeTx) == nil {
						curBlock.TotalStakeTxs += 1
						log.Printf("sync stakeTx,tx_hash=%s", stakeTx.TxHash)
					}
				}
			}
		}
		current = blocks.BlockMetas[0].Header.Height + 1
		latest = blocks.LastHeight
		log.Printf("current block height:%d", current-1)
	}

	curBlock.CurrentPos = current
	store.Mgo.UpdateBlock(curBlock)
	log.Printf("sync Transactions end,current block height:%d", current)
}
