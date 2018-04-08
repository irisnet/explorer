package tools

import (
	"net/http"
	"fmt"
	"github.com/tendermint/go-wire/data"
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/modules/fee"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/cosmos/cosmos-sdk/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"strings"
	"github.com/irisnet/irisplorer.io/server/modules/stake"
	"github.com/spf13/viper"
	"sort"
	"encoding/hex"
	wire "github.com/tendermint/go-wire"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/iavl"
	"github.com/gorilla/mux"
	"errors"
	"github.com/spf13/cast"
)

type TxResponse struct {
	Height int64       `json:"height"`
	Tx     interface{} `json:"tx"`
	TxHash string      `json:"txhash"`
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
	loop:
		for ok {
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

	return out, nil
}

// argument (so pass in a pointer to the appropriate struct)
func GetParsed(key []byte, data interface{}, height int64, prove bool) (int64, error) {
	bs, h, err := Get(key, height, prove)
	if err != nil {
		return 0, err
	}
	err = wire.ReadBinaryBytes(bs, data)
	if err != nil {
		return 0, err
	}
	return h, nil
}

// Get queries the given key and returns the value stored there and the
// height we checked at.
//
// If prove is true (and why shouldn't it be?),
// the data is fully verified before returning.  If prove is false,
// we just repeat whatever any (potentially malicious) node gives us.
// Only use that if you are running the full node yourself,
// and it is localhost or you have a secure connection (not HTTP)
func Get(key []byte, height int64, prove bool) (data.Bytes, int64, error) {
	if height < 0 {
		return nil, 0, fmt.Errorf("Height cannot be negative")
	}

	if !prove {
		node := GetNode()
		defer node.Release()
		resp, err := node.Client.ABCIQueryWithOptions("/key", key,
			rpcclient.ABCIQueryOptions{Trusted: true, Height: int64(height)})
		if resp == nil {
			return nil, height, err
		}
		return data.Bytes(resp.Response.Value), resp.Response.Height, err
	}
	val, h, _, err := GetWithProof(key, height)
	return val, h, err
}

// GetWithProof returns the values stored under a given key at the named
// height as in Get.  Additionally, it will return a validated merkle
// proof for the key-value pair if it exists, and all checks pass.
func GetWithProof(key []byte, height int64) (data.Bytes, int64, iavl.KeyProof, error) {
	node := GetNode()
	defer node.Release()
	cert, err := commands.GetCertifier()
	if err != nil {
		return nil, 0, nil, err
	}
	return client.GetWithProof(key, height, node.Client, cert)
}

func ValidateReq(w http.ResponseWriter,r *http.Request) error{
	args := mux.Vars(r)
	for ag := range args{
		if len(args[ag]) < 0 {
			err := errors.New(fmt.Sprintf("%c is not empty",ag))
			sdk.WriteError(w, err)
			return err
		}

		if ag == "page" || ag == "size" || ag == "height" {
			if _,err := cast.ToIntE(args[ag]);err != nil {
				sdk.WriteError(w, err)
				return err
			}

		}
	}
	return nil
}
