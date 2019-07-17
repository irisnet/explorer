package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
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

func (t *TxService) CopyTxFromDoc(tx document.CommonTx) model.CommonTx {

	txList := t.CopyTxListFromDoc([]document.CommonTx{tx})

	if len(txList) == 1 {
		return txList[0]
	}

	return model.CommonTx{}
}

func (_ *TxService) CopyTxListFromDoc(data []document.CommonTx) []model.CommonTx {

	commonTxUtils := make([]model.CommonTx, 0, len(data))

	for _, v := range data {
		tmpSignerArr := make([]model.Signer, 0, len(v.Signers))
		for _, v := range v.Signers {
			tmpSignerArr = append(tmpSignerArr, model.Signer{AddrHex: v.AddrHex, AddrBech32: v.AddrBech32})
		}

		tmpCoinArr := make([]utils.Coin, 0, len(v.Amount))
		for _, v := range v.Amount {
			tmpCoinArr = append(tmpCoinArr, utils.Coin{Denom: v.Denom, Amount: v.Amount})
		}

		tmpFee := utils.Fee{}
		tmpFee.Gas = v.Fee.Gas
		tmpFee.Amount = make([]utils.Coin, 0, len(v.Fee.Amount))

		for _, v := range v.Fee.Amount {
			tmpFee.Amount = append(tmpFee.Amount, utils.Coin{Denom: v.Denom, Amount: v.Amount})
		}

		tmpTx := model.CommonTx{
			Time:       v.Time,
			Height:     v.Height,
			TxHash:     v.TxHash,
			From:       v.From,
			To:         v.To,
			Type:       v.Type,
			Memo:       v.Memo,
			Status:     v.Status,
			Code:       v.Code,
			Log:        v.Log,
			GasUsed:    v.GasUsed,
			GasPrice:   v.GasPrice,
			ProposalId: v.ProposalId,
			Tags:       v.Tags,
			Signers:    tmpSignerArr,
			Amount:     tmpCoinArr,
			Fee:        tmpFee,
			ActualFee: utils.ActualFee{
				Denom:  v.ActualFee.Denom,
				Amount: v.ActualFee.Amount,
			},
			Msg: v.Msg,
			StakeCreateValidator: model.StakeCreateValidator{
				PubKey: v.StakeCreateValidator.PubKey,
				Description: model.ValDescription{
					Moniker:  v.StakeCreateValidator.Description.Moniker,
					Identity: v.StakeCreateValidator.Description.Identity,
					Website:  v.StakeCreateValidator.Description.Website,
					Details:  v.StakeCreateValidator.Description.Details,
				},
			},
			StakeEditValidator: model.StakeEditValidator{
				Description: model.ValDescription{
					Moniker:  v.StakeEditValidator.Description.Moniker,
					Identity: v.StakeEditValidator.Description.Identity,
					Website:  v.StakeEditValidator.Description.Website,
					Details:  v.StakeEditValidator.Description.Details,
				},
			},
		}

		commonTxUtils = append(commonTxUtils, tmpTx)
	}

	return commonTxUtils
}

func (service *TxService) QueryTxList(query bson.M, page, pageSize int) model.PageVo {

	total, data, err := document.CommonTx{}.QueryByPage(query, page, pageSize)

	if err != nil {
		logger.Error("query stake list ", logger.String("err", err.Error()))
		return model.PageVo{}
	}

	commonTxUtils := service.CopyTxListFromDoc(data)

	items := service.buildData(commonTxUtils)

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

			if TransTx, ok := items[i].(model.TransTx); ok {
				TransTx.From, TransTx.To = service.ParseCoinFlowFromAndTo(TransTx.Type, TransTx.From, TransTx.To)
				items[i] = TransTx
				continue
			}
		}

		valAddrArr := make([]string, 0, len(items))
		for i := 0; i < len(items); i++ {
			if stakeTx, ok := items[i].(model.StakeTx); ok {
				if service.IsValidatorAddrPrefix(stakeTx.From) {
					valAddrArr = append(valAddrArr, stakeTx.From)
				}

				if service.IsValidatorAddrPrefix(stakeTx.To) {
					valAddrArr = append(valAddrArr, stakeTx.To)
				}
			}
		}

		valAddrArr = utils.RemoveDuplicationStrArr(valAddrArr)

		monikerByAddrMap, err := document.Validator{}.QueryValidatorMonikerByAddrArr(valAddrArr)

		if err != nil {
			logger.Error("document.Validator{}.QueryValidatorMonikerByAddrArr(valAddrArr)", logger.String("err", err.Error()), logger.Any("params", valAddrArr))
		}

		for i := 0; i < len(items); i++ {
			if stakeTx, ok := items[i].(model.StakeTx); ok {
				if service.IsValidatorAddrPrefix(stakeTx.From) {
					if fromMoniker, ok := monikerByAddrMap[stakeTx.From]; ok {
						stakeTx.FromMoniker = fromMoniker
					}
				}

				if service.IsValidatorAddrPrefix(stakeTx.To) {
					if toMoniker, ok := monikerByAddrMap[stakeTx.To]; ok {
						stakeTx.ToMoniker = toMoniker
					}
				}
				items[i] = stakeTx
			}
		}

		return model.PageVo{
			Data:  items,
			Count: total,
		}
	}

	txMsgs, err := document.TxMsg{}.QueryTxMsgListByHashList(forwardTxHashs)
	if err != nil {
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

		if TransTx, ok := items[i].(model.TransTx); ok {
			TransTx.From, TransTx.To = service.ParseCoinFlowFromAndTo(TransTx.Type, TransTx.From, TransTx.To)
			items[i] = TransTx
			continue
		}
	}

	valAddrArr := make([]string, 0, len(items))
	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(model.StakeTx); ok {
			if service.IsValidatorAddrPrefix(stakeTx.From) {
				valAddrArr = append(valAddrArr, stakeTx.From)
			}

			if service.IsValidatorAddrPrefix(stakeTx.To) {
				valAddrArr = append(valAddrArr, stakeTx.To)
			}
		}
	}

	valAddrArr = utils.RemoveDuplicationStrArr(valAddrArr)

	monikerByAddrMap, err := document.Validator{}.QueryValidatorMonikerByAddrArr(valAddrArr)

	if err != nil {
		logger.Error("document.Validator{}.QueryValidatorMonikerByAddrArr(valAddrArr)", logger.String("err", err.Error()), logger.Any("params", valAddrArr))
	}

	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(model.StakeTx); ok {
			if service.IsValidatorAddrPrefix(stakeTx.From) {
				if fromMoniker, ok := monikerByAddrMap[stakeTx.From]; ok {
					stakeTx.FromMoniker = fromMoniker
				}
			}

			if service.IsValidatorAddrPrefix(stakeTx.To) {
				if toMoniker, ok := monikerByAddrMap[stakeTx.To]; ok {
					stakeTx.ToMoniker = toMoniker
				}
			}
			items[i] = stakeTx
		}
	}

	return model.PageVo{
		Data:  items,
		Count: total,
	}
}

func (service *TxService) IsValidatorAddrPrefix(addr string) bool {
	return strings.HasPrefix(addr, conf.Get().Hub.Prefix.ValAddr)
}

func (service *TxService) QueryList(query bson.M, page, pageSize int) (pageInfo model.PageVo) {
	logger.Debug("QueryList start", service.GetTraceLog())

	total, txList, err := document.CommonTx{}.QueryByPage(query, page, pageSize)

	if err != nil {
		logger.Error("query tx list by page", logger.String("err", err.Error()))
		return
	}

	data := service.CopyTxListFromDoc(txList)
	pageInfo.Data = service.buildData(data)
	pageInfo.Count = total

	logger.Debug("QueryList end", service.GetTraceLog())
	return pageInfo
}

func (service *TxService) QueryRecentTx() []model.RecentTx {
	logger.Debug("QueryRecentTx start", service.GetTraceLog())

	txListAsDoc, err := document.CommonTx{}.QueryHashActualFeeType()

	if err != nil {
		panic(err)
	}

	txs := service.CopyTxListFromDoc(txListAsDoc)

	var txList []model.RecentTx
	for _, tx := range txs {
		var recentTx = model.RecentTx{
			Fee: utils.Coin{
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

	txAsDoc, err := document.CommonTx{}.QueryTxByHash(hash)

	if err != nil {
		logger.Error("QueryTxByHash have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	blackList := map[string]document.BlackList{}

	validatorAddrMap := map[string]document.Validator{}
	govTxMsgHashMap := map[string]document.TxMsg{}
	govProposalIdMap := map[uint64]document.Proposal{}

	switch types.Convert(txAsDoc.Type) {
	case types.Declaration:
		blackList = service.QueryBlackList()
		validatorAddrMap[txAsDoc.From] = document.Validator{}
		service.GetTxAttachedFields(&validatorAddrMap, &govTxMsgHashMap, &govProposalIdMap)
	case types.Gov:
		govTxMsgHashMap[txAsDoc.TxHash] = document.TxMsg{}
		if txAsDoc.Type == types.TxTypeVote || txAsDoc.Type == types.TxTypeDeposit {
			govProposalIdMap[txAsDoc.ProposalId] = document.Proposal{}
		}
		service.GetTxAttachedFields(&validatorAddrMap, &govTxMsgHashMap, &govProposalIdMap)

	}

	tx := service.buildTx(service.CopyTxFromDoc(txAsDoc), &blackList, &validatorAddrMap, &govTxMsgHashMap, &govProposalIdMap)

	switch tx.(type) {
	case model.GovTx:
		govTx := tx.(model.GovTx)
		return govTx
	case model.StakeTx:
		stakeTx := tx.(model.StakeTx)
		if stakeTx.Type == types.TxTypeBeginRedelegate {
			res, err := document.TxMsg{}.QueryTxMsgByHash(stakeTx.Hash)

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

	total, txList, err := document.CommonTx{}.QueryByAddr(address, page, size)

	if err != nil {
		logger.Error("query tx list by address", logger.String("err", err.Error()))
	}

	if err == nil {
		result.Count = total
		result.Data = service.CopyTxListFromDoc(txList)
	}
	return
}

func (service *TxService) CountByType(query bson.M) model.TxStatisticsVo {
	logger.Debug("CountByType start", service.GetTraceLog())

	var result model.TxStatisticsVo
	counter, err := document.CommonTx{}.CountByType(query)
	if err != nil {
		logger.Error("tx count by Type ", logger.String("err", err.Error()), logger.Any("query", query))
		return result
	}

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
	var result []model.TxNumGroupByDayVo

	txNumStatList, err := document.CommonTx{}.GetTxlistByDuration(fromDate, endDate)
	if err != nil {
		logger.Error("get tx list by duration", logger.String("err", err.Error()))
		return result
	}

	for _, t := range txNumStatList {
		result = append(result, model.TxNumGroupByDayVo{
			Date: t.Date,
			Num:  t.Num,
		})
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
	case types.TxTypeBeginRedelegate:
		return from, to
	case types.TxTypeWithdrawDelegatorReward:
		return from, to
	case types.TxTypeWithdrawDelegatorRewardsAll:
		return from, to
	case types.TxTypeWithdrawValidatorRewardsAll:
		return from, to
	case types.TxTypeStakeBeginUnbonding:
		return to, from
	default:
		return from, to
	}
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

func (service *TxService) GetTxAttachedFields(candidateAddrMap *map[string]document.Validator, govTxMsgHashMap *map[string]document.TxMsg, govProposalIdMap *map[uint64]document.Proposal) {
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

	candidateArr := []document.Validator{}

	var err error
	if len(candidateAddrs) > 0 {

		candidateArr, err = document.Validator{}.QueryValidatorListByAddrList(candidateAddrs)

		if err != nil {
			logger.Error(fmt.Sprintf("query  candidator collection with condition: %v err: %v", candidateAddrs, err.Error()))
		}

		for k, _ := range *candidateAddrMap {
			for _, v := range candidateArr {
				if k == v.OperatorAddress {
					(*candidateAddrMap)[k] = v
					break
				}
			}
		}
	}

	govTxMsgArr := []document.TxMsg{}
	if len(govHashArr) > 0 {
		govTxMsgArr, err = document.TxMsg{}.QueryTxMsgListByHashList(govHashArr)

		if err != nil {
			logger.Error(fmt.Sprintf("query collection with dondition: %v err: %v", govHashArr, err.Error()))
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
		proposalArr, err = document.Proposal{}.QueryByIdList(govProposalIdArr)
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

func (service *TxService) buildData(txs []model.CommonTx) []interface{} {
	var txList []interface{}

	if len(txs) == 0 {
		return txList
	}

	blackList := map[string]document.BlackList{}
	candidateAddrMap := map[string]document.Validator{}
	govTxMsgHashMap := map[string]document.TxMsg{}
	govProposalIdMap := map[uint64]document.Proposal{}

	onlyOnce := true
	for _, v := range txs {
		switch types.Convert(v.Type) {
		case types.Declaration:
			if onlyOnce {
				blackList = service.QueryBlackList()
				onlyOnce = false
			}
			candidateAddrMap[v.From] = document.Validator{}
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

		if stakeTx, ok := txResp.(model.StakeTx); ok {
			switch stakeTx.Type {
			case types.TxTypeWithdrawDelegatorReward:
				stakeTx.From = tx.To
				stakeTx.To = tx.Tags["withdraw-address"]
				txResp = stakeTx
			case types.TxTypeWithdrawDelegatorRewardsAll, types.TxTypeWithdrawValidatorRewardsAll:
				stakeTx.To = tx.Tags["withdraw-address"]
				sourceTotal := 0
				validatorAddr := ""

				for k, _ := range tx.Tags {
					if strings.HasPrefix(k, "withdraw-reward-from-validator-") {
						sourceTotal++
						if sourceTotal == 1 {
							validatorAddr = strings.Trim(k, "withdraw-reward-from-validator-")
						}
					}
				}
				if sourceTotal == 1 {
					stakeTx.From = validatorAddr
				} else {
					stakeTx.From = strconv.Itoa(sourceTotal)
				}

				txResp = stakeTx
			}
		}

		txList = append(txList, txResp)
	}

	return txList
}

func (service *TxService) buildTx(tx model.CommonTx, blackListP *map[string]document.BlackList,
	candidateAddrMapP *map[string]document.Validator, govTxMsgHashMapP *map[string]document.TxMsg, govProposalIdMapP *map[uint64]document.Proposal) interface{} {

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
				Amount: tx.Amount,
				From:   tx.From,
				To:     tx.To,
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

func buildBaseTx(tx model.CommonTx) model.BaseTx {
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
