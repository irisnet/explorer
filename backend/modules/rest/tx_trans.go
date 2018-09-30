package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func queryTransferTx(w http.ResponseWriter, r *http.Request, query bson.M) {
	var data []document.CommonTx
	query["type"] = types.TypeTransfer
	count := utils.QueryByPage("tx_common", &data, query, "-time", r)

	var result types.Page
	if count > 0 {
		result = types.Page{
			Count: count,
			Data:  buildTransResp(data),
		}
	}

	resp, err := json.Marshal(result)
	if err == nil {
		w.Write(resp)
	}
}

func buildTransResp(txs []document.CommonTx) []types.TransTx {
	var txList []types.TransTx
	if len(txs) == 0 {
		return txList
	}
	for _, tx := range txs {
		txResp := buildTx(tx)
		txList = append(txList, txResp.(types.TransTx))
	}

	return txList
}
