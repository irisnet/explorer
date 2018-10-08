package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func queryDeclarationTx(w http.ResponseWriter, r *http.Request, query bson.M) {
	var data []document.CommonTx
	query["type"] = bson.M{
		"$in": types.DeclarationList,
	}
	var result types.Page
	count := utils.QueryByPage("tx_common", &data, query, "-time", r)
	if count > 0 {
		result = types.Page{
			Count: count,
			Data:  buildDeclarationTxResp(data),
		}
	}
	resp, err := json.Marshal(result)
	if err == nil {
		w.Write(resp)
	}
}

func buildDeclarationTxResp(txs []document.CommonTx) []types.DeclarationTx {
	var txList []types.DeclarationTx
	if len(txs) == 0 {
		return txList
	}
	for _, tx := range txs {
		txResp := buildTx(tx)
		txList = append(txList, txResp.(types.DeclarationTx))
	}

	return txList
}
