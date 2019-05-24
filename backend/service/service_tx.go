package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type TxService struct {
	BaseService
}

func (service *TxService) GetModule() Module {
	return Tx
}

func (service *TxService) QueryStakeTxList(query bson.M, page, pageSize int) model.PageVo {

	logger.Info("QueryStakeList start", service.GetTraceLog())
	var data []document.CommonTx

	total, err := pageQuery(document.CollectionNmCommonTx, nil, query, desc(document.Tx_Field_Time), page, pageSize, &data)

	if err != nil {
		logger.Error("query stake list ", logger.String("err", err.Error()))
		return model.PageVo{}
	}
	items := buildData(data)

	forwardTxHashs := make([]string, 0, len(items))
	for k, v := range items {
		if stakeTx, ok := v.(model.StakeTx); ok {
			logger.Info("query stake list", logger.String(fmt.Sprintf("k: %v ", k), stakeTx.PrintHashFromToAmount()))

			stakeTx.Signer = stakeTx.From
			items[k] = stakeTx
			forwardTxHashs = append(forwardTxHashs, stakeTx.Hash)
		}
	}

	if len(forwardTxHashs) == 0 {
		logger.Info("Result")
		for i := 0; i < len(items); i++ {
			if stakeTx, ok := items[i].(model.StakeTx); ok {
				stakeTx.From, stakeTx.To = Get(Block).(*BlockService).ParseCoinFlowFromAndTo(stakeTx.Type, stakeTx.From, stakeTx.To)
				items[i] = stakeTx
			}
		}
		return model.PageVo{
			Data:  items,
			Count: total,
		}
	}

	selector := bson.M{
		document.TxMsg_Field_Hash:    1,
		document.TxMsg_Field_Content: 1,
		document.TxMsg_Field_Type:    1,
	}
	condition := bson.M{
		document.TxMsg_Field_Hash: bson.M{"$in": forwardTxHashs},
	}

	txMsgs := make([]document.TxMsg, 0, len(forwardTxHashs))

	if err := queryAll(document.CollectionNmTxMsg, selector, condition, "", 0, &txMsgs); err != nil {
		logger.Error("query tx msg", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	logger.Info("TxMsg")
	for k, v := range txMsgs {
		logger.Info(fmt.Sprintf("K: %v \n % v", k, v))
	}

	for _, vMsg := range txMsgs {
		for k, stakeTx := range items {

			if vTx, ok := stakeTx.(model.StakeTx); ok {
				if vMsg.Hash == vTx.Hash {
					forwarAddr, err := Get(Block).(*BlockService).getForwardAddr(vMsg.Type, vMsg.Content)
					if err != nil {
						logger.Error("get forward addr ", logger.String("err", err.Error()))
						continue
					}
					vTx.From = forwarAddr
					items[k] = vTx
				}
			}
		}
	}
	logger.Info("Result")
	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(model.StakeTx); ok {
			stakeTx.From, stakeTx.To = Get(Block).(*BlockService).ParseCoinFlowFromAndTo(stakeTx.Type, stakeTx.From, stakeTx.To)
			items[i] = stakeTx
		}
	}
	return model.PageVo{
		Data:  items,
		Count: total,
	}
}

func (service *TxService) QueryList(query bson.M, page, pageSize int) (pageInfo model.PageVo) {
	logger.Debug("QueryList start", service.GetTraceLog())
	var data []document.CommonTx

	if cnt, err := pageQuery(document.CollectionNmCommonTx, nil,
		query, desc(document.Tx_Field_Time), page, pageSize, &data); err == nil {
		pageInfo.Data = buildData(data)
		pageInfo.Count = cnt
	}
	logger.Debug("QueryList end", service.GetTraceLog())
	return pageInfo
}

func (service *TxService) QueryRecentTx() []model.RecentTx {
	logger.Debug("QueryRecentTx start", service.GetTraceLog())
	var selector = bson.M{"time": 1, "tx_hash": 1, "actual_fee": 1, "type": 1}
	var txs []document.CommonTx

	err := queryAll(document.CollectionNmCommonTx, selector, nil, desc(document.Tx_Field_Time), 10, &txs)
	if err != nil {
		panic(err)
	}
	var txList []model.RecentTx
	for _, tx := range txs {
		var recentTx = model.RecentTx{
			Fee: model.Coin{
				Amount: tx.ActualFee.Amount,
				Denom:  tx.ActualFee.Denom,
			},
			Time:   tx.Time,
			TxHash: tx.TxHash,
			Type:   tx.Type,
		}
		txList = append(txList, recentTx)
	}
	logger.Debug("QueryRecentTx end", service.GetTraceLog())
	return txList
}

func (service *TxService) Query(hash string) interface{} {
	logger.Debug("Query start", service.GetTraceLog())
	dbm := getDb()
	defer dbm.Session.Close()

	var result document.CommonTx
	query := bson.M{}
	query[document.Tx_Field_Hash] = hash
	err := dbm.C(document.CollectionNmCommonTx).Find(query).Sort(desc(document.Tx_Field_Time)).One(&result)
	if err != nil {
		panic(types.CodeNotFound)
	}

	tx := service.buildTx(result)

	switch tx.(type) {
	case model.GovTx:
		govTx := tx.(model.GovTx)
		return govTx
	case model.StakeTx:
		stakeTx := tx.(model.StakeTx)
		if stakeTx.Type == types.TxTypeBeginRedelegate {
			var res document.TxMsg
			err := dbm.C(document.CollectionNmTxMsg).Find(bson.M{document.TxMsg_Field_Hash: stakeTx.Hash}).One(&res)
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
	logger.Debug("QueryList end", service.GetTraceLog())
	return tx
}

func (service *TxService) QueryByAcc(address string, page, size int) (result model.PageVo) {
	var data []document.CommonTx
	query := bson.M{}
	query["$or"] = []bson.M{{document.Tx_Field_From: address}, {document.Tx_Field_To: address}}
	var typeArr []string
	typeArr = append(typeArr, types.TxTypeTransfer)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query[document.Tx_Field_Type] = bson.M{
		"$in": typeArr,
	}
	cnt, err := pageQuery(document.CollectionNmCommonTx, nil, query, desc(document.Tx_Field_Time), page, size, &data)
	if err == nil {
		result.Count = cnt
		result.Data = data
	}
	return
}

func (service *TxService) CountByType(query bson.M) model.TxStatisticsVo {
	logger.Debug("CountByType start", service.GetTraceLog())
	var typeArr []string
	typeArr = append(typeArr, types.TxTypeTransfer)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query[document.Tx_Field_Type] = bson.M{
		"$in": typeArr,
	}

	var counter []struct {
		Type  string `bson:"_id,omitempty"`
		Count int
	}

	c := getDb().C(document.CollectionNmCommonTx)
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

	var result model.TxStatisticsVo
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
	logger.Debug("CountByType end", service.GetTraceLog())
	return result
}

func (service *TxService) QueryTxNumGroupByDay() []model.TxNumGroupByDayVo {
	logger.Debug("QueryTxNumGroupByDay start", service.GetTraceLog())

	now := time.Now()
	start := now.Add(-336 * time.Hour)

	fromDate := utils.FmtTime(utils.TruncateTime(start, utils.Day), utils.DateFmtYYYYMMDD)
	endDate := utils.FmtTime(utils.TruncateTime(now, utils.Day), utils.DateFmtYYYYMMDD)

	query := bson.M{}
	query["date"] = bson.M{"$gte": fromDate, "$lt": endDate}

	var selector = bson.M{"date": 1, "num": 1}
	var txNumStatList []document.TxNumStat

	q := orm.NewQuery()
	q.SetCollection(document.CollectionTxNumStat)
	q.SetCondition(query)
	q.SetSelector(selector).SetSort("date")
	q.SetResult(&txNumStatList)

	defer q.Release()

	var result []model.TxNumGroupByDayVo

	if err := q.Exec(); err == nil {
		for _, t := range txNumStatList {
			result = append(result, model.TxNumGroupByDayVo{
				Date: t.Date,
				Num:  t.Num,
			})
		}
	}

	logger.Debug("QueryTxNumGroupByDay end", service.GetTraceLog())
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
	db := getDb()
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
		var blackList = service.QueryBlackList(db)
		if tx.Type == types.TxTypeStakeCreateValidator {
			var moniker = tx.StakeCreateValidator.Description.Moniker
			var identity = tx.StakeCreateValidator.Description.Identity
			var website = tx.StakeCreateValidator.Description.Website
			var details = tx.StakeCreateValidator.Description.Details
			if desc, ok := blackList[tx.To]; ok {
				moniker = desc.Moniker
				identity = desc.Identity
				website = desc.Website
				details = desc.Details
			}
			dtx.Moniker = moniker
			dtx.Details = details
			dtx.Website = website
			dtx.Identity = identity
		} else if tx.Type == types.TxTypeStakeEditValidator {
			var moniker = tx.StakeEditValidator.Description.Moniker
			var identity = tx.StakeEditValidator.Description.Identity
			var website = tx.StakeEditValidator.Description.Website
			var details = tx.StakeEditValidator.Description.Details
			if desc, ok := blackList[tx.From]; ok {
				moniker = desc.Moniker
				identity = desc.Identity
				website = desc.Website
				details = desc.Details
			}
			dtx.Moniker = moniker
			dtx.Details = details
			dtx.Website = website
			dtx.Identity = identity
		} else if tx.Type == types.TxTypeUnjail {
			candidateDb := db.C(document.CollectionNmStakeRoleCandidate)
			var can document.Candidate
			candidateDb.Find(bson.M{document.Candidate_Field_Address: dtx.Owner}).One(&can)
			var moniker = can.Description.Moniker
			var identity = can.Description.Identity
			var website = can.Description.Website
			var details = can.Description.Details
			if desc, ok := blackList[tx.From]; ok {
				moniker = desc.Moniker
				identity = desc.Identity
				website = desc.Website
				details = desc.Details
			}
			dtx.Moniker = moniker
			dtx.Details = identity
			dtx.Website = website
			dtx.Identity = details
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
		err := txMsgDb.Find(bson.M{document.TxMsg_Field_Hash: govTx.Hash}).One(&res)
		if err != nil {
			return govTx
		}

		if govTx.Type == types.TxTypeSubmitProposal {
			var msg model.MsgSubmitProposal
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.Title = msg.Title
			govTx.Description = msg.Description
			govTx.ProposalType = msg.ProposalType
		} else if govTx.Type == types.TxTypeDeposit {
			var msg model.MsgDeposit
			json.Unmarshal([]byte(res.Content), &msg)
			govTx.ProposalId = msg.ProposalID
			govTx.Amount = msg.Amount
		} else if govTx.Type == types.TxTypeVote {
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
		Status:      tx.Status,
		GasLimit:    tx.Fee.Gas,
		GasUsed:     tx.GasUsed,
		GasPrice:    tx.GasPrice,
		Memo:        tx.Memo,
		Timestamp:   tx.Time,
	}
}
