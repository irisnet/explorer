package tools

import (
	"fmt"
	"log"
	"strings"
	sdk "github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/cosmos/cosmos-sdk/modules/nonce"
	"github.com/tendermint/go-wire/data"
	"github.com/irisnet/irisplorer.io/modules/store/m"
	"github.com/irisnet/irisplorer.io/modules/stake"
	"github.com/tendermint/tendermint/types"
	"encoding/hex"
)

func ParseTx(txByte types.Tx) (string, interface{}) {
	txb, _ := sdk.LoadTx(txByte)
	txl, ok := txb.Unwrap().(sdk.TxLayer)
	var txi sdk.Tx
	var coinTx m.CoinTx
	var stakeTx m.StakeTx
	var nonceAddr data.Bytes
	for ok {
		txi = txl.Next()
		switch txi.Unwrap().(type) {
		case coin.SendTx:
			ctx, _ := txi.Unwrap().(coin.SendTx)
			coinTx.From = fmt.Sprintf("%s", ctx.Inputs[0].Address.Address)
			coinTx.To = fmt.Sprintf("%s", ctx.Outputs[0].Address.Address)
			coinTx.Amount = ctx.Inputs[0].Coins
			coinTx.TxHash = strings.ToUpper(hex.EncodeToString(txByte.Hash()))
			return "coin", coinTx
		case nonce.Tx:
			ctx, _ := txi.Unwrap().(nonce.Tx)
			nonceAddr = ctx.Signers[0].Address
			break
		case stake.TxUnbond, stake.TxDelegate, stake.TxDeclareCandidacy:
			kind, _ := txi.GetKind()
			stakeTx.From = fmt.Sprintf("%s", nonceAddr)
			stakeTx.Type = strings.Replace(kind, "stake/", "", -1)
			stakeTx.TxHash = strings.ToUpper(hex.EncodeToString(txByte.Hash()))
			switch kind {
			case stake.TypeTxUnbond:
				ctx, _ := txi.Unwrap().(stake.TxUnbond)
				stakeTx.Amount.Denom = "fermion"
				stakeTx.Amount.Amount = int64(ctx.Shares)
				stakeTx.PubKey = fmt.Sprintf("%s", ctx.PubKey.KeyString())
				break
			case stake.TypeTxDelegate:
				ctx, _ := txi.Unwrap().(stake.TxDelegate)
				stakeTx.Amount.Denom = ctx.Bond.Denom
				stakeTx.Amount.Amount = ctx.Bond.Amount
				stakeTx.PubKey = fmt.Sprintf("%s", ctx.PubKey.KeyString())
				break
			case stake.TypeTxDeclareCandidacy:
				ctx, _ := txi.Unwrap().(stake.TxDeclareCandidacy)
				stakeTx.Amount.Denom = ctx.BondUpdate.Bond.Denom
				stakeTx.Amount.Amount = ctx.BondUpdate.Bond.Amount
				stakeTx.PubKey = fmt.Sprintf("%s", ctx.PubKey.KeyString())
				break
			}
			return "stake", stakeTx
		default:
			log.Printf("upsupproted tx type")
		}
		txl, ok = txi.Unwrap().(sdk.TxLayer)
	}
	return "", nil
}
