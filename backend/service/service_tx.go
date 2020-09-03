package service

import (
	"fmt"
	"strings"
	"time"

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

// get tx list
func (service *TxService) QueryTxList(query bson.M, page, pageSize int, istotal bool) vo.PageVo {

	query = document.FilterUnknownTxs(query)
	total, data, err := document.CommonTx{}.QueryByPage(query, page, pageSize, istotal)

	if err != nil {
		logger.Error("query stake list ", logger.String("err", err.Error()))
		return vo.PageVo{}
	}

	commonTxUtils := buildTxVOsFromDoc(data)
	items := service.buildTxs(commonTxUtils)

	// get tx from and to base amount coin flow direction
	//items = parseFromAndToByAmountCoinFlow(items, false)

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

	query = document.FilterUnknownTxs(query)
	total, txList, err := document.CommonTx{}.QueryByPage(query, page, pageSize, istotal)

	if err != nil {
		logger.Error("query tx list by page", logger.String("err", err.Error()))
		return
	}
	var unit []string
	for _, val := range txList {
		if len(val.Amount) > 0 && val.Amount[0].Denom != types.StakeUint {
			unit = append(unit, strings.Split(val.Amount[0].Denom, types.AssetMinDenom)[0])
		}
	}

	decimalMap := getDecimalMapData(unit)

	for i, val := range txList {
		if len(val.Amount) == 0 {
			continue
		}
		denom := strings.Split(val.Amount[0].Denom, types.AssetMinDenom)[0]
		if dem, ok := decimalMap[denom]; ok && dem > 0 {
			txList[i].Amount[0].Denom = denom
			txList[i].Amount[0].Amount = txList[i].Amount[0].Amount / float64(dem)
		}
	}

	data := buildTxVOsFromDoc(txList)
	var baseData []interface{}
	for _, tx := range data {
		txResp := buildBaseTx(tx)

		monikers := make(map[string]string, 1)

		baseData = append(baseData, vo.StakeTx{
			TransTx: vo.TransTx{
				BaseTx:   txResp,
				Monikers: monikers,
				Msgs:     tx.Msgs,
			},
		})
	}
	baseData = service.getValidatorMonikerByAddress(baseData)

	pageInfo.Data = baseData
	pageInfo.Count = total

	logger.Debug("QueryBaseList end", service.GetTraceLog())
	return pageInfo
}

func (service *TxService) QueryHtlcTx(query bson.M, page, pageSize int, istotal bool) (pageInfo vo.PageVo) {
	query = document.FilterUnknownTxs(query)
	total, txList, err := document.CommonTx{}.QueryByPage(query, page, pageSize, istotal)

	if err != nil {
		logger.Error("query tx list by page", logger.String("err", err.Error()))
		return
	}
	var unit []string
	for _, val := range txList {
		if len(val.Amount) > 0 && val.Amount[0].Denom != types.StakeUint {
			unit = append(unit, strings.Split(val.Amount[0].Denom, types.AssetMinDenom)[0])
		}
	}

	decimalMap := getDecimalMapData(unit)

	for i, val := range txList {
		if len(val.Amount) == 0 {
			continue
		}
		denom := strings.Split(val.Amount[0].Denom, types.AssetMinDenom)[0]
		if dem, ok := decimalMap[denom]; ok && dem > 0 {
			txList[i].Amount[0].Denom = denom
			txList[i].Amount[0].Amount = txList[i].Amount[0].Amount / float64(dem)
		}
	}

	data := buildTxVOsFromDoc(txList)
	var baseData []vo.HtlcTx
	for _, tx := range data {
		monikers := make(map[string]string, 1)
		expireheight := int64(-1)
		if tx.Status == types.Success {
			switch tx.Type {
			case types.TxTypeCreateHTLC:
				msgData := tx.Msgs[0].MsgData.(msgvo.TxMsgCreateHTLC)
				expireheight = tx.Height + int64(msgData.TimeLock)
			}
		}

		baseData = append(baseData, vo.HtlcTx{
			BaseTx:       buildBaseTx(tx),
			ExpireHeight: expireheight,
			Monikers:     monikers,
			Events:       tx.Events,
			Msgs:         tx.Msgs,
		})
	}

	pageInfo.Data = baseData
	pageInfo.Count = total
	return pageInfo
}

// get recent txs
func (service *TxService) QueryRecentTx() vo.RecentTxRespond {
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
func (service *TxService) QueryTxDetail(hash string) interface{} {
	logger.Debug("Query start", service.GetTraceLog())

	txAsDoc, err := document.CommonTx{}.QueryTxByHash(hash)
	if err != nil {
		logger.Error("QueryTxByHash have error", logger.String("err", err.Error()))
		return nil
	}
	txList := buildTxVOsFromDoc([]document.CommonTx{txAsDoc})
	txVOs := service.buildTxs(txList)

	items := service.getValidatorMonikerByAddress(txVOs)
	//add events in tx detail api
	if serviceTx, ok := items[0].(vo.ServiceTx); ok {
		serviceTx.Events = serviceTx.BaseTx.Events
		return serviceTx
	}

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

func (service *TxService) QueryTxNumGroupByDay(intervalDays int64) vo.TxNumGroupByDayVoRespond {
	logger.Debug("QueryTxNumGroupByDay start", service.GetTraceLog())

	now := time.Now()
	today := utils.TruncateTime(time.Now(), utils.Day)
	start := today.Add(time.Duration(-intervalDays*24) * time.Hour)

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
			Date:         t.Date,
			Num:          t.Num,
			TotalAccNum:  t.TotalAccNum,
			DelegatorNum: t.DelegatorNum,

			TokenStat: vo.TokenStatStr{
				TotalSupply:      t.TokenStat.TotalSupply,
				Circulation:      t.TokenStat.Circulation,
				Bonded:           t.TokenStat.Bonded,
				FoundationBonded: t.TokenStat.FoundationBonded,
			},
		})
	}

	logger.Debug("QueryTxNumGroupByDay end", service.GetTraceLog())
	return result
}

func (service *TxService) QueryTxType(txType string) vo.QueryTxTypeRespond {
	if txType == "all" {
		length := len(types.BankList) + len(types.DeclarationList) + len(types.StakeList) + len(types.GovernanceList) + len(types.AssetList) + len(types.RandList) + len(types.GuardianList)
		typeList := make([]string, 0, length)
		res := append(typeList, types.BankList...)
		res = append(res, types.DeclarationList...)
		res = append(res, types.StakeList...)
		res = append(res, types.GovernanceList...)
		res = append(res, types.AssetList...)
		res = append(res, types.RandList...)
		res = append(res, types.GuardianList...)
		res = append(res, types.HTLCList...)
		res = append(res, types.CoinswapList...)
		res = append(res, types.OrcaleList...)
		res = append(res, types.NftList...)
		res = append(res, types.ServiceList...)
		res = append(res, types.CrisisList...)
		return res
	}
	switch txType {
	case "trans":
		return types.BankList
	case "declaration":
		return types.DeclarationList
	case "stake":
		return types.StakeList
	case "gov":
		return types.GovernanceList
	case "asset":
		return types.AssetList
	case "rand":
		return types.RandList
	case "guardian":
		return types.GuardianList
	case "htlc":
		return types.HTLCList
	case "coinswap":
		return types.CoinswapList
	case "orcale":
		return types.OrcaleList
	case "nft":
		return types.NftList
	case "service":
		return types.ServiceList
	case "slashing":
		return types.CrisisList
	}
	return nil
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

		var createvalidator vo.StakeCreateValidator
		var editvalidator vo.StakeEditValidator
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
			case types.TxTypeRequestRand:
				msgVO := msgvo.TxMsgRequestRand{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgRequestRandByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeCreateHTLC:
				msgVO := msgvo.TxMsgCreateHTLC{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgCreateHTLCByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeClaimHTLC:
				msgVO := msgvo.TxMsgClaimHTLC{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgClaimHTLCyUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeRefundHTLC:
				msgVO := msgvo.TxMsgRefundHTLC{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("BuildTxMsgRefundHTLCByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}

				break
			case types.TxTypeStakeEditValidator:
				msgVO := msgvo.TxMsgStakeEdit{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgStakeEdit ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
					editvalidator.Description = vo.ValDescription{
						Moniker:  msgVO.ValDescription.Moniker,
						Identity: msgVO.ValDescription.Identity,
						Website:  msgVO.ValDescription.Website,
						Details:  msgVO.ValDescription.Details,
					}
					editvalidator.CommissionRate = msgVO.CommissionRate
				}
				break
			case types.TxTypeStakeCreateValidator:
				msgVO := msgvo.TxMsgStakeCreate{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgStakeCreate ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
					createvalidator.Description = vo.ValDescription{
						Moniker:  msgVO.ValDescription.Moniker,
						Identity: msgVO.ValDescription.Identity,
						Website:  msgVO.ValDescription.Website,
						Details:  msgVO.ValDescription.Details,
					}
					createvalidator.Commission = vo.CommissionMsg{
						Rate:          msgVO.Commission.Rate,
						MaxChangeRate: msgVO.Commission.MaxChangeRate,
						MaxRate:       msgVO.Commission.MaxRate,
					}
					createvalidator.PubKey = msgVO.PubKey
				}
				break
			case types.TxTypeUnjail:
				msgVO := msgvo.TxMsgUnjail{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgUnjail ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeBeginRedelegate:
				msgVO := msgvo.TxMsgBeginRedelegate{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgBeginRedelegate ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeStakeBeginUnbonding:
				msgVO := msgvo.TxMsgBeginUnbonding{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgBeginUnbonding ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeStakeDelegate:
				msgVO := msgvo.TxMsgDelegate{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgDelegate ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeDeleteProfiler:
				msgVO := msgvo.TxMsgDeleteProfiler{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgDeleteProfiler ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeAddProfiler:
				msgVO := msgvo.TxMsgAddProfiler{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgAddProfiler ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeDeleteTrustee:
				msgVO := msgvo.TxMsgDeleteTrustee{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgDeleteTrustee ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeAddTrustee:
				msgVO := msgvo.TxMsgAddTrustee{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgAddTrustee ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeAddLiquidity:
				msgVO := msgvo.TxMsgAddLiquidity{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgAddLiquidity ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeRemoveLiquidity:
				msgVO := msgvo.TxMsgRemoveLiquidity{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgRemsoveLiquidity ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeSwapOrder:
				msgVO := msgvo.TxMsgSwapOrder{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgSwapOrder ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break

			case types.TxTypeSetWithdrawAddress:
				msgVO := msgvo.TxMsgSetWithdrawAddress{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgSetWithdrawAddress ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeWithdrawDelegatorReward:
				msgVO := msgvo.TxMsgWithdrawDelegatorReward{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxMsgWithdrawDelegatorReward ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeMsgWithdrawValidatorCommission:
				msgVO := msgvo.TxTypeMsgWithdrawValidatorCommission{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxTypeMsgWithdrawValidatorCommission  ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			case types.TxTypeMsgFundCommunityPool:
				msgVO := msgvo.TxTypeMsgFundCommunityPool{}
				if err := msgVO.BuildMsgByUnmarshalJson(utils.MarshalJsonIgnoreErr(m.MsgData)); err != nil {
					logger.Error("Build TxTypeMsgFundCommunityPool ByUnmarshalJson", logger.String("err", err.Error()))
				} else {
					msgDataVO = msgVO
				}
				break
			default:
				msgDataVO = m.MsgData
			}

			tmpMsgsArr = append(tmpMsgsArr, vo.MsgItem{
				Type:    m.Type,
				MsgData: msgDataVO,
			})
		}

		tmpTx := vo.CommonTx{
			Time:   v.Time,
			Height: v.Height,
			TxHash: v.TxHash,
			//From:       v.From,
			//To:         v.To,
			Type:       v.Type,
			Memo:       v.Memo,
			Status:     utils.FailtoFailed(v.Status),
			Code:       v.Code,
			Log:        v.Log,
			GasUsed:    v.GasUsed,
			GasWanted:  v.GasWanted,
			GasPrice:   v.GasPrice,
			ProposalId: v.ProposalId,
			Events:     loadEvents(v.Events),
			Signers:    tmpSignerArr,
			Amount:     tmpCoinArr,
			Fee:        tmpFee,
			ActualFee: utils.ActualFee{
				Denom:  v.ActualFee.Denom,
				Amount: v.ActualFee.Amount,
			},
			Msgs: tmpMsgsArr,
		}

		commonTxUtils = append(commonTxUtils, tmpTx)
	}

	return commonTxUtils
}

// get validator moniker by address which in stake tx
func (service *TxService) getValidatorMonikerByAddress(items []interface{}) []interface{} {

	// set moniker value
	for i := 0; i < len(items); i++ {
		if stakeTx, ok := items[i].(vo.StakeTx); ok {
			for _, msg := range stakeTx.Msgs {
				switch msg.Type {
				case types.TxTypeBeginRedelegate:
					msgData := msg.MsgData.(msgvo.TxMsgBeginRedelegate)
					FromMoniker, ToMoniker := service.BuildFTMoniker(msgData.ValidatorSrcAddr, msgData.ValidatorDstAddr)
					if stakeTx.Monikers == nil {
						stakeTx.Monikers = make(map[string]string, 2)
					}
					if FromMoniker != "" && msgData.ValidatorSrcAddr != "" {
						stakeTx.Monikers[msgData.ValidatorSrcAddr] = FromMoniker
					}
					if ToMoniker != "" && msgData.ValidatorDstAddr != "" {
						stakeTx.Monikers[msgData.ValidatorDstAddr] = ToMoniker
					}
					//todo
				case types.TxTypeStakeDelegate:
					msgData := msg.MsgData.(msgvo.TxMsgDelegate)
					FromMoniker, _ := service.BuildFTMoniker(msgData.ValidatorAddr, "")
					if stakeTx.Monikers == nil {
						stakeTx.Monikers = make(map[string]string, 1)
					}
					if FromMoniker != "" && msgData.ValidatorAddr != "" {
						stakeTx.Monikers[msgData.ValidatorAddr] = FromMoniker
					}

					//case types.TxTypeSetWithdrawAddress:
					//	msgData := msg.MsgData.(msgvo.TxMsgSetWithdrawAddress)
				case types.TxTypeStakeBeginUnbonding:
					msgData := msg.MsgData.(msgvo.TxMsgBeginUnbonding)
					FromMoniker, _ := service.BuildFTMoniker(msgData.ValidatorAddr, "")
					if stakeTx.Monikers == nil {
						stakeTx.Monikers = make(map[string]string, 1)
					}
					if FromMoniker != "" && msgData.ValidatorAddr != "" {
						stakeTx.Monikers[msgData.ValidatorAddr] = FromMoniker
					}
				case types.TxTypeWithdrawDelegatorReward:
					msgData := msg.MsgData.(msgvo.TxMsgWithdrawDelegatorReward)
					FromMoniker, _ := service.BuildFTMoniker(msgData.ValidatorAddr, "")
					if stakeTx.Monikers == nil {
						stakeTx.Monikers = make(map[string]string, 1)
					}
					if FromMoniker != "" && msgData.ValidatorAddr != "" {
						stakeTx.Monikers[msgData.ValidatorAddr] = FromMoniker
					}
				}
			}
			items[i] = stakeTx
		}
		if declarationTx, ok := items[i].(vo.DeclarationTx); ok {
			for _, msg := range declarationTx.Msgs {
				switch declarationTx.Type {
				case types.TxTypeStakeEditValidator:
					msgData := msg.MsgData.(msgvo.TxMsgStakeEdit)
					FromMoniker, _ := service.BuildFTMoniker(msgData.ValidatorAddr, "")
					if declarationTx.Monikers == nil {
						declarationTx.Monikers = make(map[string]string, 1)
					}
					if FromMoniker != "" && msgData.ValidatorAddr != "" {
						declarationTx.Monikers[msgData.ValidatorAddr] = FromMoniker
					}
					declarationTx.OperatorAddr = msgData.ValidatorAddr
				case types.TxTypeStakeCreateValidator:
					msgData := msg.MsgData.(msgvo.TxMsgStakeCreate)
					FromMoniker, _ := service.BuildFTMoniker(msgData.ValidatorAddr, "")
					if declarationTx.Monikers == nil {
						declarationTx.Monikers = make(map[string]string, 1)
					}
					if FromMoniker != "" && msgData.ValidatorAddr != "" {
						declarationTx.Monikers[msgData.ValidatorAddr] = FromMoniker
					}
					declarationTx.OperatorAddr = msgData.ValidatorAddr
				case types.TxTypeUnjail:
					msgData := msg.MsgData.(msgvo.TxMsgUnjail)
					FromMoniker, _ := service.BuildFTMoniker(msgData.ValidatorAddr, "")
					if declarationTx.Monikers == nil {
						declarationTx.Monikers = make(map[string]string, 1)
					}
					if FromMoniker != "" && msgData.ValidatorAddr != "" {
						declarationTx.Monikers[msgData.ValidatorAddr] = FromMoniker
					}
					declarationTx.OperatorAddr = msgData.ValidatorAddr
				}
			}
			items[i] = declarationTx
		}
	}

	return items
}
func (service *TxService) TxCount() int {

	if num, err := (document.CommonTx{}).GetTxCount(); err == nil {
		return num
	}
	return 0
}
func (service *TxService) buildTxs(txs []vo.CommonTx) []interface{} {
	var txList []interface{}

	if len(txs) == 0 {
		return txList
	}

	blackList := map[string]document.BlackList{}
	govProposalIdMap := map[uint64]document.Proposal{}

	onlyOnce := true
	for _, v := range txs {
		switch types.Convert(v.Type) {
		case types.Declaration:
			if onlyOnce {
				blackList = service.QueryBlackList()
				onlyOnce = false
			}
		case types.Gov:
			govProposalIdMap[v.ProposalId] = document.Proposal{}

		}
	}

	service.getTxAttachedFields(&govProposalIdMap)

	for _, tx := range txs {
		txResp := txService.buildTx(tx, &blackList, &govProposalIdMap)
		txList = append(txList, txResp)
	}

	return txList
}

func (service *TxService) getTxAttachedFields(govProposalIdMap *map[uint64]document.Proposal) {
	govProposalIdArr := make([]uint64, 0, len(*govProposalIdMap))
	for k, _ := range *govProposalIdMap {
		govProposalIdArr = append(govProposalIdArr, k)
	}

	if len(govProposalIdArr) > 0 {
		proposalArr, err := document.Proposal{}.QueryByIdList(govProposalIdArr)
		if err != nil {
			logger.Error(fmt.Sprintf("query  Proposal collection with govProposalId err: %v", err.Error()))
			return
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

func (service *TxService) buildTx(tx vo.CommonTx, blackListP *map[string]document.BlackList,
	govProposalIdMapP *map[uint64]document.Proposal) interface{} {
	monikersMap := make(map[string]string)
	switch types.Convert(tx.Type) {
	case types.Trans:
		return vo.TransTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Declaration:
		dtx := vo.DeclarationTx{
			BaseTx: buildBaseTx(tx),
			//OperatorAddr: tx.From,
			SelfBond: tx.Amount,
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
		return dtx
	case types.Stake:
		return vo.StakeTx{
			TransTx: vo.TransTx{
				BaseTx:   buildBaseTx(tx),
				Msgs:     tx.Msgs,
				Monikers: monikersMap,
			},
		}

	case types.Gov:
		govTx := vo.GovTx{
			BaseTx:     buildBaseTx(tx),
			Amount:     tx.Amount,
			ProposalId: tx.ProposalId,
			Msgs:       tx.Msgs,
			Monikers:   monikersMap,
		}

		if v, ok := (*govProposalIdMapP)[govTx.ProposalId]; ok {
			govTx.Title = v.Title
			govTx.ProposalType = v.Type
		}
		return govTx
	case types.Asset:
		return vo.AssetTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Rand:
		return vo.AssetTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Htlc:
		expireheight := int64(-1)
		if tx.Status == types.Success {
			switch tx.Type {
			case types.TxTypeCreateHTLC:
				msgData := tx.Msgs[0].MsgData.(msgvo.TxMsgCreateHTLC)
				expireheight = tx.Height + int64(msgData.TimeLock)
			}
		}
		return vo.HtlcTx{
			BaseTx:       buildBaseTx(tx),
			Msgs:         tx.Msgs,
			Monikers:     monikersMap,
			ExpireHeight: expireheight,
		}
	case types.Coinswap:
		return vo.CoinswapTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Guardian:
		return vo.GuardianTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Orcale:
		return vo.OrcaleTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Service:
		return vo.ServiceTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Nft:
		return vo.NftTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	case types.Crisis:
		return vo.CrisisTx{
			BaseTx:   buildBaseTx(tx),
			Msgs:     tx.Msgs,
			Monikers: monikersMap,
		}
	}
	return nil
}

//func checkTags(tags map[string]string, param msgvo.Params) map[string]string {
//	if _, ok := tags["param"]; !ok {
//		bytesData, _ := json.Marshal(param)
//		tags["param"] = string(bytesData)
//	}
//	return tags
//}

func loadEvents(events []document.Event) (ret []vo.Event) {

	for _, val := range events {
		event := vo.Event{Type: val.Type}
		event.Attributes = make(map[string]string, len(val.Attributes))
		for _, one := range val.Attributes {
			event.Attributes[one.Key] = one.Value
		}
		ret = append(ret, event)
	}
	return
}

func buildBaseTx(tx vo.CommonTx) vo.BaseTx {
	res := vo.BaseTx{
		Hash:        tx.TxHash,
		BlockHeight: tx.Height,
		Type:        tx.Type,
		Fee:         tx.ActualFee,
		Amount:      tx.Amount,
		Status:      utils.FailtoFailed(tx.Status),
		GasLimit:    tx.Fee.Gas,
		GasUsed:     tx.GasUsed,
		GasWanted:   tx.GasWanted,
		GasPrice:    tx.GasPrice,
		Memo:        tx.Memo,
		Log:         fetchLogMessage(tx.Log),
		Timestamp:   tx.Time.UTC(),
		Events:      tx.Events,
	}
	if len(tx.Signers) > 0 {
		res.Signer = tx.Signers[0].AddrBech32
	}
	return res
}

func fetchLogMessage(logmsg string) string {
	TagMsg := []string{"out of gas in location", "failed to execute message"}
	if strings.Contains(logmsg, TagMsg[0]) || strings.Contains(logmsg, TagMsg[1]) {
		msg_begin := strings.Index(logmsg, ":")
		if msg_begin > 0 {
			return string([]byte(logmsg)[msg_begin:])
		}
	}

	return ""
}
