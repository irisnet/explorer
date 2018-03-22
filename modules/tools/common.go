package tools

import (
	"net/http"
	"fmt"
	"github.com/tendermint/go-wire/data"
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/modules/fee"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/cosmos/cosmos-sdk/client/commands"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"strings"
	"github.com/irisnet/iris-explorer/modules/stake"
	"github.com/spf13/viper"
	"sort"
	"encoding/hex"
)

type TxResponse struct {
	Height int64       `json:"height"`
	Tx   interface{} `json:"tx"`
	TxHash string   `json:"txhash"`
}


func FmtOutPutResult(w http.ResponseWriter, res interface{}) error {
	json, err := data.ToJSON(res)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%s\n", json)
	return err
}

func BuildTxResp(height int64, data []byte, raw bool, hash string) (interface{}, error) {
	tx, err := sdk.LoadTx(data)
	if err != nil {
		return tx, err
	}
	if (!raw) {
		txl, ok := tx.Unwrap().(sdk.TxLayer)
		var txi sdk.Tx
	loop: for ok {
		txi = txl.Next()
		switch txi.Unwrap().(type) {
		case fee.Fee, coin.SendTx, stake.TxDelegate, stake.TxDeclareCandidacy, stake.TxUnbond:
			tx = txi
			break loop
		}
		txl, ok = txi.Unwrap().(sdk.TxLayer)
	}
	}
	wrap := &TxResponse{height, tx, strings.ToUpper(hash)}
	return wrap, nil
}

func SearchTx(w http.ResponseWriter, queries ...string) ([]interface{}, error) {
	prove := !viper.GetBool(commands.FlagTrustNode)
	var all []*ctypes.ResultTx

	node := GetNode()
	defer node.Release()

	for _, q := range queries {

		txs, err := node.Client.TxSearch(q, prove)
		if err != nil {
			continue
		}
		all = append(all, txs...)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].Height > all[j].Height
	})


	out := make([]interface{}, 0, len(all))
	for _, r := range all {
		wrap, err := BuildTxResp(r.Height, r.Tx, false, hex.EncodeToString(r.Tx.Hash()))
		if err != nil {
			return nil, err
		}
		out = append(out, wrap)
	}

	return out,nil
}
