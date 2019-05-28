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

func (service *TxService) QueryTxList(query bson.M, page, pageSize int) model.PageVo {

	logger.Info("QueryStakeList start", service.GetTraceLog())
	var data []document.CommonTx

	total, err := pageQuery(document.CollectionNmCommonTx, nil, query, desc(document.Tx_Field_Time), page, pageSize, &data)

	if err != nil {
		logger.Error("query stake list ", logger.String("err", err.Error()))
		return model.PageVo{}
	}

	items := service.buildData(data)

	forwardTxHashs := make([]string, 0, len(items))

	for _, v := range items {
		if stakeTx, ok := v.(model.StakeTx); ok {
			if service.isForwardTxByType(stakeTx.Type) {
				forwardTxHashs = append(forwardTxHashs, stakeTx.Hash)
			}
		}
	}

	if len(forwardTxHashs) == 0 {
		for i := 0; i < len(items); i++ {
			if stakeTx, ok := items[i].(model.StakeTx); ok {
				stakeTx.From, stakeTx.To = service.ParseCoinFlowFromAndTo(stakeTx.Type, stakeTx.From, stakeTx.To)
				items[i] = stakeTx
				continue
			}

			if DeclarationTx, ok := items[i].(model.DeclarationTx); ok {
				DeclarationTx.From, DeclarationTx.To = service.ParseCoinFlowFromAndTo(DeclarationTx.Type, DeclarationTx.From, DeclarationTx.To)
				items[i] = DeclarationTx
				continue
			}

			if TransTx, ok := items[i].(model.TransTx); ok {
				TransTx.From, TransTx.To = service.ParseCoinFlowFromAndTo(TransTx.Type, TransTx.From, TransTx.To)
				items[i] = TransTx
				continue
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

	for _, vMsg := range txMsgs {
		for k, stakeTx := range items {

			if vTx, ok := stakeTx.(model.StakeTx); ok {
				if vMsg.Hash == vTx.Hash {
					forwardAddr, err := service.getForwardAddr(vMsg.Type, vMsg.Content)
					if err != nil {
						logger.Error("get forward addr ", logger.String("err", err.Error()))
						continue
					}
					vTx.From = forwardAddr
					items[k] = vTx
				}
			}
		}
	}

	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(model.StakeTx); ok {
			stakeTx.From, stakeTx.To = service.ParseCoinFlowFromAndTo(stakeTx.Type, stakeTx.From, stakeTx.To)
			items[i] = stakeTx
			continue
		}

		if DeclarationTx, ok := items[i].(model.DeclarationTx); ok {
			DeclarationTx.From, DeclarationTx.To = service.ParseCoinFlowFromAndTo(DeclarationTx.Type, DeclarationTx.From, DeclarationTx.To)
			items[i] = DeclarationTx
			continue
		}

		if TransTx, ok := items[i].(model.TransTx); ok {
			TransTx.From, TransTx.To = service.ParseCoinFlowFromAndTo(TransTx.Type, TransTx.From, TransTx.To)
			items[i] = TransTx
			continue
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

		pageInfo.Data = service.buildData(data)
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

	blackList := map[string]document.BlackList{}
	candidateAddrMap := map[string]document.Candidate{}
	govTxMsgHashMap := map[string]document.TxMsg{}
	govProposalIdMap := map[uint64]document.Proposal{}

	switch types.Convert(result.Type) {
	case types.Declaration:
		blackList = service.QueryBlackList(dbm)
		candidateAddrMap[result.From] = document.Candidate{}
		service.GetTxAttachedFields(&candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)
	case types.Gov:
		govTxMsgHashMap[result.TxHash] = document.TxMsg{}
		if result.Type == types.TxTypeVote || result.Type == types.TxTypeDeposit {
			govProposalIdMap[result.ProposalId] = document.Proposal{}
		}
		service.GetTxAttachedFields(&candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)

	}

	tx := service.buildTx(result, &blackList, &candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)

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

func (service *TxService) ParseCoinFlowFromAndTo(txType, from, to string) (string, string) {
	switch txType {
	case types.TxTypeTransfer:
		return from, to
	case types.TxTypeBurn:
		return from, ""
	case types.TxTypeStakeEditValidator:
		return "", ""
	case types.TxTypeStakeDelegate:
		return from, to
	case types.TxTypeUnjail:
		return "", ""
	case types.TxTypeSetWithdrawAddress:
		return "", ""
	case types.TxTypeStakeCreateValidator:
		return from, to
		//exchange
	case types.TxTypeBeginRedelegate:
		return to, from
	case types.TxTypeWithdrawDelegatorReward:
		return to, ""
	case types.TxTypeWithdrawDelegatorRewardsAll:
		return "", ""
	case types.TxTypeWithdrawValidatorRewardsAll:
		return "", ""
	case types.TxTypeStakeBeginUnbonding:
		return to, from
	default:
		return from, to
	}
}

func (service *TxService) QueryTxMsgByHashArr(hashArr []string) []document.TxMsg {

	selector := bson.M{
		document.TxMsg_Field_Hash:    1,
		document.TxMsg_Field_Content: 1,
		document.TxMsg_Field_Type:    1,
	}
	condition := bson.M{
		document.TxMsg_Field_Hash: bson.M{"$in": hashArr},
	}

	txMsgs := []document.TxMsg{}
	if err := queryAll(document.CollectionNmTxMsg, selector, condition, "", 0, &txMsgs); err != nil {
		logger.Error("query tx msg", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	return txMsgs
}

func (service *TxService) getForwardAddr(txType, content string) (string, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(content), &m)
	if err != nil {
		return "", err
	}

	switch txType {
	case types.TxTypeBeginRedelegate:
		if v, ok := m["validator_src_addr"].(string); ok {
			return v, nil
		}
	}
	return "", nil
}

func (service *TxService) isForwardTxByType(t string) bool {
	for _, v := range types.ForwardList {
		if v == t {
			return true
		}
	}
	return false
}

func (service *TxService) GetTxAttachedFields(candidateAddrMap *map[string]document.Candidate, govTxMsgHashMap *map[string]document.TxMsg, govProposalIdMap *map[uint64]document.Proposal) {
	candidateAddrs := make([]string, 0, len(*candidateAddrMap))
	govHashArr := make([]string, 0, len(*govTxMsgHashMap))
	govProposalIdArr := make([]uint64, 0, len(*govProposalIdMap))
	for k, _ := range *candidateAddrMap {
		candidateAddrs = append(candidateAddrs, k)
	}
	for k, _ := range *govTxMsgHashMap {
		govHashArr = append(govHashArr, k)
	}
	for k, _ := range *govProposalIdMap {
		govProposalIdArr = append(govProposalIdArr, k)
	}

	candidateArr := []document.Candidate{}

	if len(candidateAddrs) > 0 {
		canCondition := bson.M{
			document.Candidate_Field_Address: bson.M{"$in": candidateAddrs},
		}

		if err := queryAll(document.CollectionNmStakeRoleCandidate, nil, canCondition, "", 0, &candidateArr); err != nil {
			logger.Error(fmt.Sprintf("query collection(%v) with dondition: %v err: %v", document.CollectionNmStakeRoleCandidate, canCondition, err.Error()))
		}

		for k, _ := range *candidateAddrMap {
			for _, v := range candidateArr {
				if k == v.Address {
					(*candidateAddrMap)[k] = v
					break
				}
			}
		}
	}

	govTxMsgArr := []document.TxMsg{}
	if len(govHashArr) > 0 {
		txMsgCondition := bson.M{
			document.TxMsg_Field_Hash: bson.M{"$in": govHashArr},
		}
		if err := queryAll(document.CollectionNmTxMsg, nil, txMsgCondition, "", 0, &govTxMsgArr); err != nil {
			logger.Error(fmt.Sprintf("query collection(%v) with dondition: %v err: %v", document.CollectionNmStakeRoleCandidate, txMsgCondition, err.Error()))
		}

		for k, _ := range *govTxMsgHashMap {
			for _, v := range govTxMsgArr {
				if k == v.Hash {
					(*govTxMsgHashMap)[k] = v
					break
				}
			}
		}
	}

	proposalArr := []document.Proposal{}

	if len(govProposalIdArr) > 0 {
		depositVoteCondition := bson.M{
			document.Proposal_Field_ProposalId: bson.M{"$in": govProposalIdArr},
		}

		if err := queryAll(document.CollectionNmProposal, nil, depositVoteCondition, "", 0, &proposalArr); err != nil {
			logger.Error(fmt.Sprintf("query collection(%v) with dondition: %v err: %v", document.CollectionNmStakeRoleCandidate, depositVoteCondition, err.Error()))
		}
		for k, _ := range *govProposalIdMap {
			for _, v := range proposalArr {
				if k == v.ProposalId {
					(*govProposalIdMap)[k] = v
					break
				}
			}
		}
	}
}

func (service *TxService) buildData(txs []document.CommonTx) []interface{} {
	var txList []interface{}

	if len(txs) == 0 {
		return txList
	}

	db := getDb()
	defer db.Session.Close()

	blackList := map[string]document.BlackList{}
	candidateAddrMap := map[string]document.Candidate{}
	govTxMsgHashMap := map[string]document.TxMsg{}
	govProposalIdMap := map[uint64]document.Proposal{}

	for _, v := range txs {
		switch types.Convert(v.Type) {
		case types.Declaration:
			if len(blackList) == 0 {
				blackList = service.QueryBlackList(db)
			}
			candidateAddrMap[v.From] = document.Candidate{}
		case types.Gov:
			govTxMsgHashMap[v.TxHash] = document.TxMsg{}
			if v.Type == types.TxTypeVote || v.Type == types.TxTypeDeposit {
				govProposalIdMap[v.ProposalId] = document.Proposal{}
			}
		}
	}

	service.GetTxAttachedFields(&candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)

	for _, tx := range txs {
		txResp := txService.buildTx(tx, &blackList, &candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)
		txList = append(txList, txResp)
	}

	return txList
}

func (service *TxService) buildTx(tx document.CommonTx, blackListP *map[string]document.BlackList,
	candidateAddrMapP *map[string]document.Candidate, govTxMsgHashMapP *map[string]document.TxMsg, govProposalIdMapP *map[uint64]document.Proposal) interface{} {

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
			Amount:   tx.Amount,
		}
		if tx.Type == types.TxTypeStakeCreateValidator {
			dtx.From = tx.From
			dtx.To = tx.To
			dtx.OperatorAddr = tx.To
			var moniker = tx.StakeCreateValidator.Description.Moniker
			var identity = tx.StakeCreateValidator.Description.Identity
			var website = tx.StakeCreateValidator.Description.Website
			var details = tx.StakeCreateValidator.Description.Details
			if desc, ok := (*blackListP)[tx.To]; ok {
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
			dtx.From = dtx.Signer
			dtx.To = tx.From
			dtx.OperatorAddr = tx.From
			var moniker = tx.StakeEditValidator.Description.Moniker
			var identity = tx.StakeEditValidator.Description.Identity
			var website = tx.StakeEditValidator.Description.Website
			var details = tx.StakeEditValidator.Description.Details
			if desc, ok := (*blackListP)[tx.From]; ok {
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
			dtx.From = dtx.Signer
			dtx.To = tx.From
			dtx.OperatorAddr = tx.From
			can := (*candidateAddrMapP)[dtx.Owner]

			var moniker = can.Description.Moniker
			var identity = can.Description.Identity
			var website = can.Description.Website
			var details = can.Description.Details
			if desc, ok := (*blackListP)[tx.From]; ok {
				moniker = desc.Moniker
				identity = desc.Identity
				website = desc.Website
				details = desc.Details
			}
			dtx.Moniker = moniker
			dtx.Details = details
			dtx.Website = website
			dtx.Identity = identity
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

		if govTx.Type == types.TxTypeSubmitProposal {
			if v, ok := (*govTxMsgHashMapP)[govTx.Hash]; ok {
				var msg model.MsgSubmitProposal
				if err := json.Unmarshal([]byte(v.Content), &msg); err != nil {
					logger.Error("unmarshal gov tx msg ", logger.String("tx hash", govTx.Hash), logger.String("content", v.Content), logger.Any("err", err.Error()))
				}
				govTx.Title = msg.Title
				govTx.Description = msg.Description
				govTx.ProposalType = msg.ProposalType
			}
		} else if govTx.Type == types.TxTypeDeposit {

			if v, ok := (*govTxMsgHashMapP)[govTx.Hash]; ok {
				var msg model.MsgDeposit
				if err := json.Unmarshal([]byte(v.Content), &msg); err != nil {
					logger.Error("unmarshal gov tx msg ", logger.String("tx hash", govTx.Hash), logger.String("content", v.Content), logger.Any("err", err.Error()))
				}
				govTx.Amount = msg.Amount
			}

			if v, ok := (*govProposalIdMapP)[govTx.ProposalId]; ok {
				govTx.Title = v.Title
				govTx.ProposalType = v.Type
				govTx.Description = v.Description
			}

		} else if govTx.Type == types.TxTypeVote {

			if v, ok := (*govTxMsgHashMapP)[govTx.Hash]; ok {
				var msg model.MsgVote
				if err := json.Unmarshal([]byte(v.Content), &msg); err != nil {
					logger.Error("unmarshal gov tx msg ", logger.String("tx hash", govTx.Hash), logger.String("content", v.Content), logger.Any("err", err.Error()))
				}
				govTx.Option = msg.Option
			}

			if v, ok := (*govProposalIdMapP)[govTx.ProposalId]; ok {
				govTx.Title = v.Title
				govTx.ProposalType = v.Type
				govTx.Description = v.Description
			}

		}
		return govTx
	}
	return nil
}

func buildBaseTx(tx document.CommonTx) model.BaseTx {
	res := model.BaseTx{
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

	if len(tx.Signers) > 0 {
		res.Signer = tx.Signers[0].AddrBech32
	}
	return res
}
