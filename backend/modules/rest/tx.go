package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

type TxDay struct {
	Time  string `bson:"_id,omitempty"`
	Count int
}

func registerQueryTxList(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxList, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		query := bson.M{}
		request.ParseForm()
		if len(request.Form["address"]) > 0 {
			address := request.Form["address"][0]
			query["$or"] = []bson.M{{"from": address}, {"to": address}}
		}

		if len(request.Form["height"]) > 0 {
			height, _ := strconv.Atoi(request.Form["height"][0])
			query["height"] = height
		}

		txType := args["type"]
		switch types.TxTypeFromString(txType) {
		case types.Trans:
			queryTransferTx(writer, request, query)
			break
		case types.Declaration:
			queryDeclarationTx(writer, request, query)
			break
		case types.Stake:
			queryStakeTx(writer, request, query)
			break
		case types.Gov:
			queryGovTx(writer, request, query)
			break
		}

	}).Methods("GET")
	return nil
}

func registerQueryAllStakeTxByPage(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAllStakeTxByPage, queryAllStakeTxByPage).Methods("GET")
	return nil
}

func registerQueryStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryStakeTxByAccount, queryStakeTxByAccount).Methods("GET")
	return nil
}

func registerQueryPageStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryPageStakeTxByAccount, queryPageStakeTxByAccount).Methods("GET")
	return nil
}

func registerQueryTxs(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxs, queryTxs).Methods("GET")
	return nil
}

func registerQueryTxsCounter(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxsCounter, countTxs).Methods("GET")
	return nil
}

func registerQueryTxsByAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxsByAccount, queryTxsByAccount).Methods("GET")
	return nil
}

func registerQueryTxsByDay(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxsByDay, queryTxsByDay).Methods("GET")
	return nil
}

func queryTxs(w http.ResponseWriter, r *http.Request) {
	var data []document.CommonTx
	query := bson.M{}
	var typeArr []string
	typeArr = append(typeArr, types.TypeTransfer)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query["type"] = bson.M{
		"$in": typeArr,
	}

	utils.QueryByPage("tx_common", &data, query, "-time", r)
	var result = countTxByType(r)
	result.Data = data
	res, _ := json.Marshal(result)
	w.Write(res)
}

func queryAllCoinTxByPage(w http.ResponseWriter, r *http.Request) {
	var data []document.CommonTx
	w.Write(utils.QueryList("tx_common", &data, bson.M{"type": "Transfer"}, "-time", r))
}

func queryCoinTxByAccount(w http.ResponseWriter, r *http.Request) {
}

func queryCoinTxPageByAccount(w http.ResponseWriter, r *http.Request) {
	var data []document.CommonTx
	args := mux.Vars(r)
	address := args["address"]
	w.Write(utils.QueryList("tx_common", &data, bson.M{"type": "Transfer", "address": address}, "-time", r))
}

func queryTxsByAccount(w http.ResponseWriter, r *http.Request) {
	var data []document.CommonTx
	args := mux.Vars(r)
	address := args["address"]
	query := bson.M{}
	query["$or"] = []bson.M{{"from": address}, {"to": address}}
	var typeArr []string
	typeArr = append(typeArr, types.TypeTransfer)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query["type"] = bson.M{
		"$in": typeArr,
	}
	w.Write(utils.QueryList("tx_common", &data, query, "-time", r))
}

func queryTxsByBlock(w http.ResponseWriter, r *http.Request) {
	var data []document.CommonTx
	args := mux.Vars(r)
	height := args["height"]
	iHeight, _ := strconv.Atoi(height)
	query := bson.M{}
	query["height"] = iHeight
	var typeArr []string
	typeArr = append(typeArr, types.TypeTransfer)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query["type"] = bson.M{
		"$in": typeArr,
	}
	w.Write(utils.QueryList("tx_common", &data, query, "-time", r))
}

func queryTxsByDay(w http.ResponseWriter, r *http.Request) {
	c := utils.GetDatabase().C("tx_common")
	defer c.Database.Session.Close()

	now := time.Now()
	d, _ := time.ParseDuration("-336h") //14 days ago
	start := now.Add(d)
	pipe := c.Pipe(
		[]bson.M{
			{"$match": bson.M{
				"time": bson.M{
					"$gte": time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()),
					"$lt":  time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, start.Location()),
				},
			}},
			{"$project": bson.M{
				"day": bson.M{"$substr": []interface{}{"$time", 0, 10}},
			}},
			{"$group": bson.M{
				"_id":   "$day",
				"count": bson.M{"$sum": 1},
			}},
			{"$sort": bson.M{
				"_id": 1,
			}},
		},
	)
	var txDays []TxDay
	var result []TxDay
	pipe.All(&txDays)
	var i time.Duration
	var j int
	day := start
	for day.Unix() < time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, start.Location()).Unix() {
		key := day.Format("2006-01-02")
		if len(txDays) > j && txDays[j].Time == key {
			result = append(result, txDays[j])
			j++
		} else {
			var txDay TxDay
			txDay.Time = key
			txDay.Count = 0
			result = append(result, txDay)
		}
		i++
		day = start.Add(i * 24 * time.Hour)
	}
	resultByte, _ := json.Marshal(result)
	w.Write(resultByte)
}

func queryAllStakeTxByPage(w http.ResponseWriter, r *http.Request) {
	var data []document.CommonTx
	w.Write(utils.QueryList("tx_common", &data, bson.M{"type": bson.M{"$in": []string{"Delegate", "CompleteUnbonding", "BeginUnbonding"}}}, "-time", r))
}

func queryStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
}

func queryPageStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
}

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryTxs,
		registerQueryTxsByAccount,
		registerQueryStakeTxByAccount,
		registerQueryPageStakeTxByAccount,
		registerQueryAllStakeTxByPage,
		registerQueryTxsByDay,
		//new
		registerQueryTxList,
		registerQueryTxsCounter,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func buildTx(tx document.CommonTx) interface{} {
	db := utils.GetDatabase()
	defer db.Session.Close()

	switch types.Convert(tx.Type) {
	case types.Trans:
		return types.TransTx{
			BaseTx: buildBaseTx(tx),
			From:   tx.From,
			To:     tx.To,
			Amount: tx.Amount,
		}
	case types.Declaration:
		dtx := types.DeclarationTx{
			BaseTx:   buildBaseTx(tx),
			SelfBond: tx.Amount,
			Owner:    tx.From,
			Pubkey:   tx.StakeCreateValidator.PubKey,
		}
		if tx.Type == types.TypeCreateValidator {
			dtx.Moniker = tx.StakeCreateValidator.Description.Moniker
			dtx.Details = tx.StakeCreateValidator.Description.Details
			dtx.Website = tx.StakeCreateValidator.Description.Website
			dtx.Identity = tx.StakeCreateValidator.Description.Identity
		} else if tx.Type == types.TypeEditValidator {
			dtx.Moniker = tx.StakeEditValidator.Description.Moniker
			dtx.Details = tx.StakeEditValidator.Description.Details
			dtx.Website = tx.StakeEditValidator.Description.Website
			dtx.Identity = tx.StakeEditValidator.Description.Identity
		} else if tx.Type == types.TypeUnRevoke {
			candidateDb := db.C("stake_role_candidate")
			var can document.Candidate
			candidateDb.Find(bson.M{"address": dtx.Owner}).One(&can)
			dtx.Moniker = can.Description.Moniker
			dtx.Details = can.Description.Details
			dtx.Website = can.Description.Website
			dtx.Identity = can.Description.Identity
		}
		return dtx
	case types.Stake:
		return types.StakeTx{
			TransTx: types.TransTx{
				BaseTx: buildBaseTx(tx),
				From:   tx.From,
				To:     tx.To,
				Amount: tx.Amount,
			},
		}
	case types.Gov:
		govTx := types.GovTx{
			BaseTx:     buildBaseTx(tx),
			Amount:     tx.Amount,
			From:       tx.From,
			ProposalId: tx.ProposalId,
		}

		txMsgDb := db.C("tx_msg")
		var res document.TxMsg
		err := txMsgDb.Find(bson.M{"hash": govTx.Hash}).One(&res)
		if err != nil {
			return govTx
		}

		if govTx.Type == types.TypeSubmitProposal {
			var msg MsgSubmitProposal
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.Title = msg.Title
			govTx.Description = msg.Description
			govTx.ProposalType = msg.ProposalType
		} else if govTx.Type == types.TypeDeposit {
			var msg MsgDeposit
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.ProposalId = msg.ProposalID
			govTx.Amount = msg.Amount
		} else if govTx.Type == types.TypeVote {
			var msg MsgVote
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.ProposalId = msg.ProposalID
			govTx.Option = msg.Option
		}

		return govTx
	}
	return nil
}

func buildBaseTx(tx document.CommonTx) types.BaseTx {
	return types.BaseTx{
		Hash:        tx.TxHash,
		BlockHeight: tx.Height,
		Type:        tx.Type,
		Fee:         tx.ActualFee,
		GasLimit:    tx.Fee.Gas,
		GasUsed:     tx.GasUsed,
		GasPrice:    tx.GasPrice,
		Memo:        tx.Memo,
		Timestamp:   tx.Time,
	}
}
