package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/vo/msgvo"
	"gopkg.in/mgo.v2/bson"
)

type TxService struct {
	BaseService
}

func (service *TxService) GetModule() Module {
	return Tx
}

// get tx list
func (service *TxService) QueryTxList(query bson.M, page, pageSize int, istotal bool) vo.PageVo {

	total, data, err := document.CommonTx{}.QueryByPage(query, page, pageSize, istotal)

	if err != nil {
		logger.Error("query stake list ", logger.String("err", err.Error()))
		return vo.PageVo{}
	}

	commonTxUtils := buildTxVOsFromDoc(data)
	items := service.buildTxVOs(commonTxUtils, false)

	forwardTxHashs := make([]string, 0, len(items))

	for _, v := range items {
		if stakeTx, ok := v.(vo.StakeTx); ok {
			if service.isForwardTxByType(stakeTx.Type) {
				forwardTxHashs = append(forwardTxHashs, stakeTx.Hash)
			}
		}
	}

	if len(forwardTxHashs) != 0 {
		txMsgs, err := document.TxMsg{}.QueryTxMsgListByHashList(forwardTxHashs)
		if err != nil {
			logger.Error("query tx msg", logger.String("err", err.Error()))
			panic(types.CodeNotFound)
		}

		for _, vMsg := range txMsgs {
			for k, stakeTx := range items {

				if vTx, ok := stakeTx.(vo.StakeTx); ok {
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
	}

	// get tx from and to base amount coin flow direction
	items = parseFromAndToByAmountCoinFlow(items, false)

	// get validator moniker by address
	items = service.getValidatorMonikerByAddress(items)

	return vo.PageVo{
		Data:  items,
		Count: total,
	}
}

// get base tx list in home page
func (service *TxService) QueryBaseList(query bson.M, page, pageSize int, istotal bool) (pageInfo vo.PageVo) {
	logger.Debug("QueryBaseList start", service.GetTraceLog())

	total, txList, err := document.CommonTx{}.QueryByPage(query, page, pageSize, istotal)

	if err != nil {
		logger.Error("query tx list by page", logger.String("err", err.Error()))
		return
	}

	data := buildTxVOsFromDoc(txList)
	var baseData []vo.BaseTx
	for _, tx := range data {
		txResp := buildBaseTx(tx)
		baseData = append(baseData, txResp)
	}

	pageInfo.Data = baseData
	pageInfo.Count = total

	logger.Debug("QueryBaseList end", service.GetTraceLog())
	return pageInfo
}

// get recent txs
func (service *TxService) QueryRecentTx() []vo.RecentTx {
	logger.Debug("QueryRecentTx start", service.GetTraceLog())

	txListAsDoc, err := document.CommonTx{}.QueryHashActualFeeType()

	if err != nil {
		panic(err)
	}

	txs := buildTxVOsFromDoc(txListAsDoc)

	var txList []vo.RecentTx
	for _, tx := range txs {
		var recentTx = vo.RecentTx{
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

// query tx detail by hash
func (service *TxService) Query(hash string) interface{} {
	logger.Debug("Query start", service.GetTraceLog())
	var (
		forwardTxHashes []string
	)

	txAsDoc, err := document.CommonTx{}.QueryTxByHash(hash)
	if err != nil {
		logger.Error("QueryTxByHash have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	txVOs := service.buildTxVOs([]vo.CommonTx{buildTxVOFromDoc(txAsDoc)}, true)

	for _, v := range txVOs {
		if stakeTx, ok := v.(vo.StakeTx); ok {
			if service.isForwardTxByType(stakeTx.Type) {
				forwardTxHashes = append(forwardTxHashes, stakeTx.Hash)
			}
		}
	}
	if len(forwardTxHashes) != 0 {
		txMsgs, err := document.TxMsg{}.QueryTxMsgListByHashList(forwardTxHashes)
		if err != nil {
			logger.Error("query tx msg", logger.String("err", err.Error()))
			panic(types.CodeNotFound)
		}

		for _, vMsg := range txMsgs {
			for k, stakeTx := range txVOs {
				if vTx, ok := stakeTx.(vo.StakeTx); ok {
					if vMsg.Hash == vTx.Hash {
						forwardAddr, err := service.getForwardAddr(vMsg.Type, vMsg.Content)
						if err != nil {
							logger.Error("get forward addr ", logger.String("err", err.Error()))
							continue
						}
						vTx.From = forwardAddr
						txVOs[k] = vTx
					}
				}
			}
		}
	}

	items := parseFromAndToByAmountCoinFlow(txVOs, true)
	items = service.getValidatorMonikerByAddress(items)

	logger.Debug("getTxsByFilter end", service.GetTraceLog())
	return items[0]
}

func (service *TxService) QueryByAcc(address string, page, size int, istotal bool) (result vo.PageVo) {

	total, txList, err := document.CommonTx{}.QueryByAddr(address, page, size, istotal)

	if err != nil {
		logger.Error("query tx list by address", logger.String("err", err.Error()))
	}

	if err == nil {
		result.Count = total
		result.Data = buildTxVOsFromDoc(txList)
	}
	return
}

func (service *TxService) CountByType(query bson.M) vo.TxStatisticsVo {
	logger.Debug("CountByType start", service.GetTraceLog())

	var result vo.TxStatisticsVo
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

func (service *TxService) QueryTxNumGroupByDay() []vo.TxNumGroupByDayVo {
	logger.Debug("QueryTxNumGroupByDay start", service.GetTraceLog())

	now := time.Now()
	start := now.Add(-336 * time.Hour)

	fromDate := utils.FmtTime(utils.TruncateTime(start, utils.Day), utils.DateFmtYYYYMMDD)
	endDate := utils.FmtTime(utils.TruncateTime(now, utils.Day), utils.DateFmtYYYYMMDD)
	var result []vo.TxNumGroupByDayVo

	txNumStatList, err := document.CommonTx{}.GetTxlistByDuration(fromDate, endDate)
	if err != nil {
		logger.Error("get tx list by duration", logger.String("err", err.Error()))
		return result
	}

	for _, t := range txNumStatList {
		result = append(result, vo.TxNumGroupByDayVo{
			Date: t.Date,
			Num:  t.Num,
		})
	}

	logger.Debug("QueryTxNumGroupByDay end", service.GetTraceLog())
	return result
}

func (service *TxService) QueryTxType(txType string) interface{} {
	if txType == "" {
		length := len(types.BankList) + len(types.DeclarationList) + len(types.StakeList) + len(types.GovernanceList) + len(types.AssetList) + len(types.RandList)
		typeList := make([]string, 0, length)
		res := append(typeList, types.BankList...)
		res = append(res, types.DeclarationList...)
		res = append(res, types.StakeList...)
		res = append(res, types.GovernanceList...)
		res = append(res, types.AssetList...)
		res = append(res, types.RandList...)
		return res
	}
	switch txType {
	case "Trans":
		return types.BankList
	case "Declaration":
		return types.DeclarationList
	case "Stake":
		return types.StakeList
	case "Gov":
		return types.GovernanceList
	case "Asset":
		return types.AssetList
	case "Rand":
		return types.RandList
	}
	return nil
}

func buildTxVOFromDoc(tx document.CommonTx) vo.CommonTx {
	txList := buildTxVOsFromDoc([]document.CommonTx{tx})

	if len(txList) == 1 {
		return txList[0]
	}

	return vo.CommonTx{}
}

func buildTxVOsFromDoc(data []document.CommonTx) []vo.CommonTx {
	commonTxUtils := make([]vo.CommonTx, 0, len(data))

	for _, v := range data {
		tmpSignerArr := make([]vo.Signer, 0, len(v.Signers))
		for _, v := range v.Signers {
			tmpSignerArr = append(tmpSignerArr, vo.Signer{AddrHex: v.AddrHex, AddrBech32: v.AddrBech32})
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
		tmpMsgsArr := make([]vo.MsgItem, 0, len(v.Msgs))

		// build tx msgVO
		for _, m := range v.Msgs {
			var msgDataVO interface{}
			switch m.Type {
			case types.TxTypeIssueToken:
				msgVO := msgvo.TxMsgIssueToken{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgIssueTokenByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeEditToken:
				msgVO := msgvo.TxMsgEditToken{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgEditTokenByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeMintToken:
				msgVO := msgvo.TxMsgMintToken{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgMintTokenByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgVO.To = checkMintToAddress(msgVO.Owner, msgVO.To)
					msgDataVO = msgVO
				}
				break
			case types.TxTypeTransferTokenOwner:
				msgVO := msgvo.TxMsgTransferTokenOwner{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgTransferTokenOwnerByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeCreateGateway:
				msgVO := msgvo.TxMsgCreateGateway{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgCreateGatewayByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeEditGateway:
				msgVO := msgvo.TxMsgEditGateway{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgEditGatewayByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeTransferGatewayOwner:
				msgVO := msgvo.TxMsgTransferGatewayOwner{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgTransferGatewayOwnerByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeSetMemoRegexp:
				msgVO := msgvo.TxMsgSetMemoRegexp{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgSetMemoRegexpByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeRequestRand:
				msgVO := msgvo.TxMsgRequestRand{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgRequestRandByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			}

			tmpMsgsArr = append(tmpMsgsArr, vo.MsgItem{
				Type:    v.Type,
				MsgData: msgDataVO,
			})
		}

		tmpTx := vo.CommonTx{
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
			GasWanted:  v.GasWanted,
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
			StakeCreateValidator: vo.StakeCreateValidator{
				PubKey: v.StakeCreateValidator.PubKey,
				Description: vo.ValDescription{
					Moniker:  v.StakeCreateValidator.Description.Moniker,
					Identity: v.StakeCreateValidator.Description.Identity,
					Website:  v.StakeCreateValidator.Description.Website,
					Details:  v.StakeCreateValidator.Description.Details,
				},
				Commission: vo.CommissionMsg{
					Rate:          v.StakeCreateValidator.Commission.Rate,
					MaxChangeRate: v.StakeCreateValidator.Commission.MaxChangeRate,
					MaxRate:       v.StakeCreateValidator.Commission.MaxRate,
				},
			},
			StakeEditValidator: vo.StakeEditValidator{
				Description: vo.ValDescription{
					Moniker:  v.StakeEditValidator.Description.Moniker,
					Identity: v.StakeEditValidator.Description.Identity,
					Website:  v.StakeEditValidator.Description.Website,
					Details:  v.StakeEditValidator.Description.Details,
				},
				CommissionRate: v.StakeEditValidator.CommissionRate,
			},
			Msgs: tmpMsgsArr,
		}

		commonTxUtils = append(commonTxUtils, tmpTx)
	}

	return commonTxUtils
}

func parseCoinFlowFromAndToForTxList(txType, from, to string) (string, string) {
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

func parseCoinFlowFromAndToForTxDetail(txType, from, to string) (string, string) {
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
		return from, to
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

func parseFromAndToByAmountCoinFlow(items []interface{}, forTxDetail bool) []interface{} {
	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(vo.StakeTx); ok {
			if forTxDetail {
				stakeTx.From, stakeTx.To = parseCoinFlowFromAndToForTxDetail(stakeTx.Type, stakeTx.From, stakeTx.To)
			} else {
				stakeTx.From, stakeTx.To = parseCoinFlowFromAndToForTxList(stakeTx.Type, stakeTx.From, stakeTx.To)
			}
			items[i] = stakeTx
			continue
		}

		if TransTx, ok := items[i].(vo.TransTx); ok {
			if forTxDetail {
				TransTx.From, TransTx.To = parseCoinFlowFromAndToForTxDetail(TransTx.Type, TransTx.From, TransTx.To)
			} else {
				TransTx.From, TransTx.To = parseCoinFlowFromAndToForTxList(TransTx.Type, TransTx.From, TransTx.To)
			}

			items[i] = TransTx
			continue
		}
	}

	return items
}

// get validator moniker by address which in stake tx
func (service *TxService) getValidatorMonikerByAddress(items []interface{}) []interface{} {
	// get validator addresses
	valAddrArr := make([]string, 0, len(items))
	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(vo.StakeTx); ok {
			if isValidatorAddrPrefix(stakeTx.From) {
				valAddrArr = append(valAddrArr, stakeTx.From)
			}

			if isValidatorAddrPrefix(stakeTx.To) {
				valAddrArr = append(valAddrArr, stakeTx.To)
			}
		}
	}
	valAddrArr = utils.RemoveDuplicationStrArr(valAddrArr)

	// query moniker by addresses
	monikerByAddrMap, err := document.Validator{}.QueryValidatorMonikerByAddrArr(valAddrArr)
	if err != nil {
		logger.Error("document.Validator{}.QueryValidatorMonikerByAddrArr(valAddrArr)", logger.String("err", err.Error()), logger.Any("params", valAddrArr))
	}

	// set moniker value
	blacklist := service.QueryBlackList()
	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(vo.StakeTx); ok {
			if isValidatorAddrPrefix(stakeTx.From) {
				if fromMoniker, ok := monikerByAddrMap[stakeTx.From]; ok {
					stakeTx.FromMoniker = fromMoniker
				}
				if blackone, ok := blacklist[stakeTx.From]; ok {
					stakeTx.FromMoniker = blackone.Moniker
				}
			}

			if isValidatorAddrPrefix(stakeTx.To) {
				if toMoniker, ok := monikerByAddrMap[stakeTx.To]; ok {
					stakeTx.ToMoniker = toMoniker
				}
				if blackone, ok := blacklist[stakeTx.To]; ok {
					stakeTx.ToMoniker = blackone.Moniker
				}
			}
			items[i] = stakeTx
		}
	}

	return items
}

func isValidatorAddrPrefix(addr string) bool {
	return strings.HasPrefix(addr, conf.Get().Hub.Prefix.ValAddr)
}

func (service *TxService) buildTxVOs(txs []vo.CommonTx, isDetail bool) []interface{} {
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

	service.getTxAttachedFields(&candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)

	for _, tx := range txs {
		txResp := txService.buildTxVO(tx, &blackList, &candidateAddrMap, &govTxMsgHashMap, &govProposalIdMap)

		if stakeTx, ok := txResp.(vo.StakeTx); ok {
			switch stakeTx.Type {
			case types.TxTypeWithdrawDelegatorReward:
				stakeTx.From = tx.To
				stakeTx.To = tx.Tags[types.TxTag_WithDrawAddress]
				txResp = stakeTx
			case types.TxTypeWithdrawDelegatorRewardsAll, types.TxTypeWithdrawValidatorRewardsAll:
				stakeTx.To = tx.Tags[types.TxTag_WithDrawAddress]
				sourceTotal := 0
				//validatorAddr := ""
				if isDetail {
					var validatorAddr bytes.Buffer
					for k, _ := range tx.Tags {
						if strings.HasPrefix(k, types.TxTag_WithDrawRewardFromValidator) {
							sourceTotal++

							if sourceTotal > 1 {
								_, err := validatorAddr.WriteString(",")
								if err != nil {
									logger.Error("ValidatorAddr WriteString", logger.String("err", err.Error()))
								}
							}
							_, err := validatorAddr.WriteString(string([]byte(k)[len(types.TxTag_WithDrawRewardFromValidator):]))
							if err != nil {
								logger.Error("ValidatorAddr WriteString", logger.String("err", err.Error()))
							}
						}
					}
					stakeTx.From = validatorAddr.String()
				} else {
					validatorAddr := ""

					for k, _ := range tx.Tags {
						if strings.HasPrefix(k, types.TxTag_WithDrawRewardFromValidator) {
							sourceTotal++
							if sourceTotal == 1 {
								validatorAddr = string([]byte(k)[len(types.TxTag_WithDrawRewardFromValidator):])
							}
						}
					}
					if sourceTotal == 1 {
						stakeTx.From = validatorAddr
					} else {
						stakeTx.From = strconv.Itoa(sourceTotal)
					}
				}

				txResp = stakeTx
			}
		}

		txList = append(txList, txResp)
	}

	return txList
}

func (service *TxService) getTxAttachedFields(candidateAddrMap *map[string]document.Validator,
	govTxMsgHashMap *map[string]document.TxMsg,
	govProposalIdMap *map[uint64]document.Proposal) {
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

func (service *TxService) buildTxVO(tx vo.CommonTx, blackListP *map[string]document.BlackList,
	candidateAddrMapP *map[string]document.Validator, govTxMsgHashMapP *map[string]document.TxMsg,
	govProposalIdMapP *map[uint64]document.Proposal) interface{} {

	switch types.Convert(tx.Type) {
	case types.Trans:
		if tx.Type == types.TxTypeSetMemoRegexp {
			return vo.AssetTx{
				BaseTx: buildBaseTx(tx),
				From:   tx.From,
				To:     tx.To,
				Amount: tx.Amount,
				Msgs:   tx.Msgs,
			}
		}
		return vo.TransTx{
			BaseTx: buildBaseTx(tx),
			From:   tx.From,
			To:     tx.To,
			Amount: tx.Amount,
		}
	case types.Declaration:
		dtx := vo.DeclarationTx{
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
			dtx.Commission = vo.CommissionMsg{
				Rate:          tx.StakeCreateValidator.Commission.Rate,
				MaxRate:       tx.StakeCreateValidator.Commission.MaxRate,
				MaxChangeRate: tx.StakeCreateValidator.Commission.MaxChangeRate,
			}
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
			dtx.Commission = vo.CommissionMsg{
				Rate: tx.StakeEditValidator.CommissionRate,
			}
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
		return vo.StakeTx{
			TransTx: vo.TransTx{
				BaseTx: buildBaseTx(tx),
				Amount: tx.Amount,
				From:   tx.From,
				To:     tx.To,
			},
		}

	case types.Gov:
		govTx := vo.GovTx{
			BaseTx:     buildBaseTx(tx),
			Amount:     tx.Amount,
			From:       tx.From,
			ProposalId: tx.ProposalId,
		}

		if govTx.Type == types.TxTypeSubmitProposal {
			if v, ok := (*govTxMsgHashMapP)[govTx.Hash]; ok {
				var msg vo.MsgSubmitProposal
				if err := json.Unmarshal([]byte(v.Content), &msg); err != nil {
					logger.Error("unmarshal gov tx msg ", logger.String("tx hash", govTx.Hash), logger.String("content", v.Content), logger.Any("err", err.Error()))
				}
				govTx.Title = msg.Title
				govTx.Description = msg.Description
				govTx.ProposalType = msg.ProposalType
				govTx.Tags = tx.Tags
				govTx.Software = msg.Software
				govTx.Version = msg.Version
				govTx.SwitchHeight = msg.SwitchHeight
				govTx.Treshold = msg.Treshold
			}
		} else if govTx.Type == types.TxTypeDeposit {

			if v, ok := (*govTxMsgHashMapP)[govTx.Hash]; ok {
				var msg vo.MsgDeposit
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
				var msg vo.MsgVote
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
	case types.Asset:
		return vo.AssetTx{
			BaseTx: buildBaseTx(tx),
			From:   tx.From,
			To:     tx.To,
			Amount: tx.Amount,
			Msgs:   tx.Msgs,
		}
	case types.Rand:
		return vo.AssetTx{
			BaseTx: buildBaseTx(tx),
			From:   tx.From,
			To:     tx.To,
			Amount: tx.Amount,
			Tags:   tx.Tags,
			Msgs:   tx.Msgs,
		}
	}
	return nil
}

func buildBaseTx(tx vo.CommonTx) vo.BaseTx {
	res := vo.BaseTx{
		Hash:        tx.TxHash,
		BlockHeight: tx.Height,
		Type:        tx.Type,
		Fee:         tx.ActualFee,
		Status:      tx.Status,
		GasLimit:    tx.Fee.Gas,
		GasUsed:     tx.GasUsed,
		GasWanted:   tx.GasWanted,
		GasPrice:    tx.GasPrice,
		Memo:        tx.Memo,
		Log:         fetchLogMessage(tx.Log),
		Timestamp:   tx.Time,
	}
	if tx.Type != types.TxTypeWithdrawDelegatorRewardsAll && tx.Type != types.TxTypeWithdrawValidatorRewardsAll {
		res.Tags = tx.Tags
	}
	if len(tx.Signers) > 0 {
		res.Signer = tx.Signers[0].AddrBech32
	}
	return res
}

func fetchLogMessage(logmsg string) string {

	const TagMsg = "\"message\":\""
	msg_begin := strings.Index(logmsg, TagMsg)
	if msg_begin > 0 {
		data := strings.Split(logmsg, TagMsg)
		return data[1][:len(data[1])-2]
	}
	return ""
}
