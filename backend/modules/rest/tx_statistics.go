package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

func countTxs(w http.ResponseWriter, r *http.Request) {
	result := countTxByType(r)
	res, _ := json.Marshal(result)
	w.Write(res)
}

type TxCounter struct {
	TransCnt       int
	StakeCnt       int
	DeclarationCnt int
	GovCnt         int
	Data           []document.CommonTx
}

func countTxByType(r *http.Request) TxCounter {
	query := bson.M{}
	r.ParseForm()
	if len(r.Form["address"]) > 0 {
		address := r.Form["address"][0]
		query["$or"] = []bson.M{{"from": address}, {"to": address}}
	}

	if len(r.Form["height"]) > 0 {
		height, _ := strconv.Atoi(r.Form["height"][0])
		query["height"] = height
	}

	var typeArr []string
	typeArr = append(typeArr, types.TypeTransfer)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query["type"] = bson.M{
		"$in": typeArr,
	}

	var counter []struct {
		Type  string `bson:"_id,omitempty"`
		Count int
	}

	c := utils.GetDatabase().C("tx_common")
	defer c.Database.Session.Close()

	pipe := c.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   "$type",
				"count": bson.M{"$sum": 1},
			}},
		},
	)

	pipe.All(&counter)

	var result TxCounter
	for _, cnt := range counter {
		switch types.Convert(cnt.Type) {
		case types.Trans:
			result.TransCnt = result.TransCnt + cnt.Count
		case types.Declaration:
			result.DeclarationCnt = result.DeclarationCnt + cnt.Count
		case types.Stake:
			result.StakeCnt = result.StakeCnt + cnt.Count
		case types.Gov:
			result.GovCnt = result.GovCnt + cnt.Count
		}
	}
	return result
}
