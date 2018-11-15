package service

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type TxService struct {
	*BaseService
}

func GetTx() *TxService {
	return txService
}

func (service *TxService) QueryList(query bson.M, page, pageSize int) model.Page {
	var data []document.CommonTx
	pageInfo := service.QueryPage(document.CollectionNmCommonTx, &data, query, "-time", page, pageSize)
	pageInfo.Data = buildData(data)
	return pageInfo
}

func (service *TxService) QueryLatest(query bson.M, page, pageSize int) model.Page {
	var data []document.CommonTx
	pageInfo := service.QueryPage(document.CollectionNmCommonTx, &data, query, "-time", page, pageSize)
	return pageInfo
}

func (service *TxService) Query(hash string) interface{} {
	dbm := service.GetDb()
	defer dbm.Session.Close()

	var result document.CommonTx
	query := bson.M{}
	query["tx_hash"] = hash
	err := dbm.C(document.CollectionNmCommonTx).Find(query).Sort("-time").One(&result)
	if err != nil {
		panic(types.ErrorCodeNotFound)
	}

	tx := service.buildTx(result)

	switch tx.(type) {
	case model.GovTx:
		govTx := tx.(model.GovTx)
		return govTx
	case model.StakeTx:
		stakeTx := tx.(model.StakeTx)
		if stakeTx.Type == types.TypeBeginRedelegation || stakeTx.Type == types.TypeCompleteRedelegation {
			var res document.TxMsg
			err := dbm.C(document.CollectionNmTxMsg).Find(bson.M{"hash": stakeTx.Hash}).One(&res)
			if err != nil {
				break
			}
			var msg model.MsgBeginRedelegate
			if err = json.Unmarshal([]byte(res.Content), &msg); err == nil {
				stakeTx.From = msg.DelegatorAddr
				stakeTx.To = msg.ValidatorDstAddr
				stakeTx.Source = msg.ValidatorSrcAddr
			}
		}
		return stakeTx
	}
	return tx
}

func (service *TxService) QueryByAcc(address string, page, size int) model.Page {
	var data []document.CommonTx
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
	return service.QueryPage(document.CollectionNmCommonTx, &data, query, "-time", page, size)
}

func (service *TxService) CountByType(query bson.M) model.TxCounter {
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

	c := service.GetDb().C(document.CollectionNmCommonTx)
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

	var result model.TxCounter
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

func (service *TxService) CountByDay() []model.TxDay {
	c := service.GetDb().C(document.CollectionNmCommonTx)
	defer c.Database.Session.Close()

	now := time.Now()
	d, _ := time.ParseDuration("-336h") //14 days ago
	start := now.Add(d)

	fromDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, start.Location())

	log.Println(fmt.Sprintf("from:%s,to:%s", fromDate.String(), endDate.String()))

	pipe := c.Pipe(
		[]bson.M{
			{"$match": bson.M{
				"time": bson.M{
					"$gte": fromDate,
					"$lt":  endDate,
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
	var txDays []model.TxDay
	var result []model.TxDay
	pipe.All(&txDays)
	var i time.Duration
	var j int
	day := start
	for day.Unix() < endDate.Unix() {
		key := day.Format("2006-01-02")
		if len(txDays) > j && txDays[j].Time == key {
			result = append(result, txDays[j])
			j++
		} else {
			var txDay model.TxDay
			txDay.Time = key
			txDay.Count = 0
			result = append(result, txDay)
		}
		i++
		day = start.Add(i * 24 * time.Hour)
	}
	return result
}

func buildData(txs []document.CommonTx) []interface{} {
	var txList []interface{}

	if len(txs) == 0 {
		return txList
	}
	for _, tx := range txs {
		txResp := txService.buildTx(tx)
		txList = append(txList, txResp)
	}

	return txList
}

func (service *TxService) buildTx(tx document.CommonTx) interface{} {
	db := service.GetDb()
	defer db.Session.Close()

	switch types.Convert(tx.Type) {
	case types.Trans:
		return model.TransTx{
			BaseTx: buildBaseTx(tx),
			From:   tx.From,
			To:     tx.To,
			Amount: tx.Amount,
		}
	case types.Declaration:
		dtx := model.DeclarationTx{
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
			candidateDb := db.C(document.CollectionNmStakeRoleCandidate)
			var can document.Candidate
			candidateDb.Find(bson.M{"address": dtx.Owner}).One(&can)
			dtx.Moniker = can.Description.Moniker
			dtx.Details = can.Description.Details
			dtx.Website = can.Description.Website
			dtx.Identity = can.Description.Identity
		}
		return dtx
	case types.Stake:
		return model.StakeTx{
			TransTx: model.TransTx{
				BaseTx: buildBaseTx(tx),
				From:   tx.From,
				To:     tx.To,
				Amount: tx.Amount,
			},
		}
	case types.Gov:
		govTx := model.GovTx{
			BaseTx:     buildBaseTx(tx),
			Amount:     tx.Amount,
			From:       tx.From,
			ProposalId: tx.ProposalId,
		}

		txMsgDb := db.C(document.CollectionNmTxMsg)
		var res document.TxMsg
		err := txMsgDb.Find(bson.M{"hash": govTx.Hash}).One(&res)
		if err != nil {
			return govTx
		}

		if govTx.Type == types.TypeSubmitProposal {
			var msg model.MsgSubmitProposal
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.Title = msg.Title
			govTx.Description = msg.Description
			govTx.ProposalType = msg.ProposalType
		} else if govTx.Type == types.TypeDeposit {
			var msg model.MsgDeposit
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.ProposalId = msg.ProposalID
			govTx.Amount = msg.Amount
		} else if govTx.Type == types.TypeVote {
			var msg model.MsgVote
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.ProposalId = msg.ProposalID
			govTx.Option = msg.Option
		}

		return govTx
	}
	return nil
}

func buildBaseTx(tx document.CommonTx) model.BaseTx {
	return model.BaseTx{
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
