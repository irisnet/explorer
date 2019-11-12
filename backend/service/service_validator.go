package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"math/big"
	"strconv"
	"sort"
)

type ValidatorService struct {
	BaseService
	monikerMap map[string]string
}

func (service *ValidatorService) GetModule() Module {
	return Validator
}

func (service *ValidatorService) GetValidators(typ, origin string, page, size int, istotal bool) interface{} {
	if origin == "browser" {
		var result []lcd.ValidatorVo
		var blackList = service.QueryBlackList()

		total, validatorList, err := document.Validator{}.GetValidatorListByPage(typ, page, size, true, istotal)
		if err != nil || total <= 0 {
			if err != nil {
				logger.Error("GetValidatorListByPage have error", logger.String("error", err.Error()))
			}
			panic(types.CodeNotFound)
		}

		var totalVotingPower = getTotalVotingPower()
		for i, v := range validatorList {
			if desc, ok := blackList[v.OperatorAddress]; ok {
				validatorList[i].Description.Moniker = desc.Moniker
				validatorList[i].Description.Identity = desc.Identity
				validatorList[i].Description.Website = desc.Website
				validatorList[i].Description.Details = desc.Details
				validatorList[i].Icons = ""
			}
			var validator lcd.ValidatorVo
			if err := utils.Copy(validatorList[i], &validator); err != nil {
				logger.Error("utils.Copy have error", logger.String("error", err.Error()))
			}
			validator.VotingRate = float32(v.VotingPower) / float32(totalVotingPower)
			selfbond := ComputeSelfBonded(v.Tokens, v.DelegatorShares, v.SelfBond)
			validator.SelfBond = selfbond
			result = append(result, validator)
		}

		return vo.PageVo{
			Data:  result,
			Count: total,
		}
	}

	return service.queryValForRainbow(page, size)
}

func (service *ValidatorService) GetVoteTxsByValidatorAddr(validatorAddr string, page, size int, istotal bool) vo.ValidatorVotePage {

	validatorAcc := utils.Convert(conf.Get().Hub.Prefix.AccAddr, validatorAddr)
	total, proposalsAsDoc, err := document.Proposal{}.QueryIdTitleStatusVotedTxhashByValidatorAcc(validatorAcc, page, size, istotal)

	if err != nil {
		logger.Error("QueryIdTitleStatusVotedTxhashByValidatorAcc", logger.String("err", err.Error()))
		return vo.ValidatorVotePage{}
	}

	items := make([]vo.ValidatorVote, 0, size)

	for _, v := range proposalsAsDoc {
		votedOption, txhash := "", ""

		for _, vote := range v.Votes {
			if vote.Voter == validatorAcc {
				votedOption = vote.Option
				txhash = vote.TxHash
			}
		}

		tmp := vo.ValidatorVote{
			Title:      v.Title,
			ProposalId: v.ProposalId,
			Status:     v.Status,
			Voted:      votedOption,
			TxHash:     txhash,
		}

		items = append(items, tmp)
	}

	return vo.ValidatorVotePage{
		Total: total,
		Items: items,
	}
}

func (service *ValidatorService) GetDepositedTxByValidatorAddr(validatorAddr string, page, size int, istotal bool) vo.ValidatorDepositTxPage {

	validatorAcc := utils.Convert(conf.Get().Hub.Prefix.AccAddr, validatorAddr)
	total, txs, err := document.CommonTx{}.QueryDepositedProposalTxByValidatorWithSubmitOrDepositType(validatorAcc, page, size, istotal)

	if err != nil {
		if err.Error() != "not found" {
			logger.Error("QueryDepositedProposalTxByValidatorWithSubmitOrDepositType", logger.String("err", err.Error()))
		}
		return vo.ValidatorDepositTxPage{}
	}

	proposalIds := make([]uint64, 0, len(txs))
	for _, v := range txs {
		proposalIds = append(proposalIds, v.ProposalId)
	}

	proposerByIdMap, err := document.CommonTx{}.QueryProposalTxFromById(proposalIds)

	if err != nil {
		logger.Error("QueryProposalTxFromById", logger.String("err", err.Error()))
	}

	addrArr := make([]string, 0, len(txs))
	for _, v := range proposerByIdMap {
		addrArr = append(addrArr, utils.Convert(conf.Get().Hub.Prefix.ValAddr, v))
	}

	addrArr = utils.RemoveDuplicationStrArr(addrArr)
	validatorMonikerMap, err := document.Validator{}.QueryValidatorMonikerByAddrArr(addrArr)

	if err != nil {
		logger.Error("QueryValidatorMonikerByAddrArr", logger.String("err", err.Error()))
	}

	blackList := service.QueryBlackList()
	items := make([]vo.ValidatorDepositTx, 0, size)
	for _, v := range txs {
		submited := false
		if v.Type == types.TxTypeSubmitProposal {
			submited = true
		}

		amount := make(utils.Coins, 0, len(v.Amount))

		for _, coin := range v.Amount {
			tmp := utils.Coin{
				Denom:  coin.Denom,
				Amount: coin.Amount,
			}
			amount = append(amount, tmp)
		}

		moniker := ""
		proposer := ""
		if from, ok := proposerByIdMap[v.ProposalId]; ok {
			proposer = from
			valaddr := utils.Convert(conf.Get().Hub.Prefix.ValAddr, from)
			if m, ok := validatorMonikerMap[valaddr]; ok {
				moniker = m
			}
			if blackone, ok := blackList[valaddr]; ok {
				moniker = blackone.Moniker
			}
		}

		tmp := vo.ValidatorDepositTx{
			ProposalId:      v.ProposalId,
			Proposer:        proposer,
			Moniker:         moniker,
			DepositedAmount: amount,
			Submited:        submited,
			TxHash:          v.TxHash,
		}
		items = append(items, tmp)
	}

	return vo.ValidatorDepositTxPage{
		Total: total,
		Items: items,
	}
}

func (service *ValidatorService) GetUnbondingDelegationsFromLcd(valAddr string, page, size int) vo.UnbondingDelegationsPage {

	lcdUnbondingDelegations := lcd.GetUnbondingDelegationsByValidatorAddr(valAddr)

	items := make([]vo.UnbondingDelegations, 0, size)
	var valaddrlist []string
	for _, v := range lcdUnbondingDelegations {
		valaddrlist = append(valaddrlist, v.ValidatorAddr)
	}

	validatorMap := getValidators(valaddrlist)
	for k, v := range lcdUnbondingDelegations {
		if k >= page*size && k < (page+1)*size {

			tmp := vo.UnbondingDelegations{
				Address: v.DelegatorAddr,
				Amount:  v.Balance,
				Block:   v.CreationHeight,
				Until:   v.MinTime,
			}
			if validatorMap != nil {
				valaddr := utils.Convert(conf.Get().Hub.Prefix.ValAddr, v.DelegatorAddr)
				if valdator, ok := validatorMap[valaddr]; ok {
					tmp.Moniker = valdator.Description.Moniker
				}
			}

			items = append(items, tmp)
		}
	}

	return vo.UnbondingDelegationsPage{
		Total: len(lcdUnbondingDelegations),
		Items: items,
	}
}

func (service *ValidatorService) GetDelegationsFromLcd(valAddr string, page, size int, needpage bool, istotal bool) vo.DelegationsPage {

	var lcdDelegations lcd.ValidatorDelegations
	lcdDelegations = lcd.GetDelegationsByValidatorAddr(valAddr)
	var valaddrlist []string
	totalShareAsRat := new(big.Rat)
	for _, v := range lcdDelegations {
		valaddrlist = append(valaddrlist, v.ValidatorAddr)
		sharesAsRat, ok := new(big.Rat).SetString(v.Shares)
		if !ok {
			logger.Error("convert delegation shares type (string -> big.Rat) err", logger.String("shares str", v.Shares))
			continue
		}
		totalShareAsRat = totalShareAsRat.Add(totalShareAsRat, sharesAsRat)
	}
	validatorMap := getValidators(valaddrlist)

	addrArr := []string{valAddr}

	tokenShareRatioByValidatorAddr, err := document.Validator{}.QueryTokensAndShareRatioByValidatorAddrs(addrArr)
	if err != nil {
		logger.Debug("QueryTokensAndShareRatioByValidatorAddrs", logger.String("err", err.Error()))
	}
	sort.Sort(lcdDelegations)
	var items []vo.Delegation
	if needpage {
		items = GetDalegationbyPageSize(validatorMap, lcdDelegations, totalShareAsRat, tokenShareRatioByValidatorAddr, page, size)
	} else {
		items = GetDalegation(validatorMap, lcdDelegations, totalShareAsRat, tokenShareRatioByValidatorAddr)
	}

	return vo.DelegationsPage{
		Total: len(lcdDelegations),
		Items: items,
	}
}

func GetDalegation(validatorMap map[string]document.Validator, lcdDelegations []lcd.DelegationVo, totalShareAsRat *big.Rat, tokenShareRatio map[string]*big.Rat) []vo.Delegation {
	items := make([]vo.Delegation, 0, len(lcdDelegations))
	for _, v := range lcdDelegations {

		////amountAsFloat64 := float64(0)
		//var amountAsFloat64 string
		//if ratio, ok := tokenShareRatio[v.ValidatorAddr]; ok {
		//	if shareAsRat, ok := new(big.Rat).SetString(v.Shares); ok {
		//		amountAsRat := new(big.Rat).Mul(shareAsRat, ratio)
		//		amountAsFloat64 = amountAsRat.FloatString(18)
		//		//exact := false
		//		//amountAsFloat64, exact = amountAsRat.Float64()
		//		//if !exact {
		//		//	logger.Info("convert new(big.Rat).Mul(shareAsRat, ratio)  (big.Rat to float64) ",
		//		//		logger.Any("exact", exact),
		//		//		logger.Any("amountAsRat", amountAsRat))
		//		//}
		//	} else {
		//		logger.Error("convert validator share  type (string -> big.Rat) err", logger.String("str", v.Shares))
		//	}
		//} else {
		//	logger.Error("can not fond the validator addr from the validator collection in db", logger.String("validator addr", v.ValidatorAddr))
		//}
		//
		//totalShareAsFloat64, exact := totalShareAsRat.Float64()
		//
		//if !exact {
		//	logger.Info("convert totalShareAsFloat64  (big.Rat to float64) ",
		//		logger.Any("exact", exact),
		//		logger.Any("totalShareAsFloat64", totalShareAsFloat64))
		//}
		//
		//tmp := vo.Delegation{
		//	Address:     v.DelegatorAddr,
		//	Block:       v.Height,
		//	SelfShares:  v.Shares,
		//	TotalShares: totalShareAsFloat64,
		//	Amount:      amountAsFloat64,
		//}
		//if validatorMap != nil {
		//	if valdator, ok := validatorMap[v.ValidatorAddr]; ok {
		//		tmp.Moniker = valdator.Description.Moniker
		//	}
		//}
		tmp := getVoDelegation(v, validatorMap, totalShareAsRat, tokenShareRatio)
		items = append(items, tmp)

	}
	return items
}

func getVoDelegation(v lcd.DelegationVo, validatorMap map[string]document.Validator, totalShareAsRat *big.Rat,
	tokenShareRatio map[string]*big.Rat) vo.Delegation {
	//amountAsFloat64 := float64(0)
	var amountAsFloat64 string
	if ratio, ok := tokenShareRatio[v.ValidatorAddr]; ok {
		if shareAsRat, ok := new(big.Rat).SetString(v.Shares); ok {
			amountAsRat := new(big.Rat).Mul(shareAsRat, ratio)
			amountAsFloat64 = amountAsRat.FloatString(18)
			//exact := false
			//amountAsFloat64, exact = amountAsRat.Float64()
			//if !exact {
			//	logger.Info("convert new(big.Rat).Mul(shareAsRat, ratio)  (big.Rat to float64) ",
			//		logger.Any("exact", exact),
			//		logger.Any("amountAsRat", amountAsRat))
			//}
		} else {
			logger.Error("convert validator share  type (string -> big.Rat) err", logger.String("str", v.Shares))
		}
	} else {
		logger.Error("can not fond the validator addr from the validator collection in db", logger.String("validator addr", v.ValidatorAddr))
	}

	totalShareAsFloat64, exact := totalShareAsRat.Float64()

	if !exact {
		logger.Info("convert totalShareAsFloat64  (big.Rat to float64) ",
			logger.Any("exact", exact),
			logger.Any("totalShareAsFloat64", totalShareAsFloat64))
	}

	tmp := vo.Delegation{
		Address:     v.DelegatorAddr,
		Block:       v.Height,
		SelfShares:  v.Shares,
		TotalShares: totalShareAsFloat64,
		Amount:      amountAsFloat64,
	}
	if validatorMap != nil {
		valaddr := utils.Convert(conf.Get().Hub.Prefix.ValAddr, v.DelegatorAddr)
		if valdator, ok := validatorMap[valaddr]; ok {
			tmp.Moniker = valdator.Description.Moniker
		}
	}

	return tmp
}

func GetDalegationbyPageSize(validatorMap map[string]document.Validator, lcdDelegations []lcd.DelegationVo, totalShareAsRat *big.Rat,
	tokenShareRatio map[string]*big.Rat, page, size int) []vo.Delegation {
	items := make([]vo.Delegation, 0, size)
	for k, v := range lcdDelegations {
		if k >= page*size && k < (page+1)*size {

			////amountAsFloat64 := float64(0)
			//var amountAsFloat64 string
			//if ratio, ok := tokenShareRatio[v.ValidatorAddr]; ok {
			//	if shareAsRat, ok := new(big.Rat).SetString(v.Shares); ok {
			//		amountAsRat := new(big.Rat).Mul(shareAsRat, ratio)
			//		amountAsFloat64 = amountAsRat.FloatString(18)
			//		//exact := false
			//		//amountAsFloat64, exact = amountAsRat.Float64()
			//		//if !exact {
			//		//	logger.Info("convert new(big.Rat).Mul(shareAsRat, ratio)  (big.Rat to float64) ",
			//		//		logger.Any("exact", exact),
			//		//		logger.Any("amountAsRat", amountAsRat))
			//		//}
			//	} else {
			//		logger.Error("convert validator share  type (string -> big.Rat) err", logger.String("str", v.Shares))
			//	}
			//} else {
			//	logger.Error("can not fond the validator addr from the validator collection in db", logger.String("validator addr", v.ValidatorAddr))
			//}
			//
			//totalShareAsFloat64, exact := totalShareAsRat.Float64()
			//
			//if !exact {
			//	logger.Info("convert totalShareAsFloat64  (big.Rat to float64) ",
			//		logger.Any("exact", exact),
			//		logger.Any("totalShareAsFloat64", totalShareAsFloat64))
			//}
			//
			//tmp := vo.Delegation{
			//	Address:     v.DelegatorAddr,
			//	Block:       v.Height,
			//	SelfShares:  v.Shares,
			//	TotalShares: totalShareAsFloat64,
			//	Amount:      amountAsFloat64,
			//}
			//if validatorMap != nil {
			//	if valdator, ok := validatorMap[v.ValidatorAddr]; ok {
			//		tmp.Moniker = valdator.Description.Moniker
			//	}
			//}
			tmp := getVoDelegation(v, validatorMap, totalShareAsRat, tokenShareRatio)
			items = append(items, tmp)
		}
	}
	return items
}

func (service *ValidatorService) GetRedelegationsFromLcd(valAddr string, page, size int) vo.RedelegationPage {

	lcdReDelegations := lcd.GetRedelegationsByValidatorAddr(valAddr)
	blacklist := service.QueryBlackList()

	items := make([]vo.Redelegation, 0, size)

	for k, v := range lcdReDelegations {
		if k >= page*size && k < (page+1)*size {

			tomoniker := ""
			if validator, err := document.GetValidatorByAddr(v.ValidatorDstAddr); err == nil {
				tomoniker = validator.Description.Moniker
			}
			if blockone, ok := blacklist[v.ValidatorDstAddr]; ok {
				tomoniker = blockone.Moniker
			}
			tmp := vo.Redelegation{
				Address:   v.DelegatorAddr,
				Amount:    v.Balance,
				To:        v.ValidatorDstAddr,
				ToMoniker: tomoniker,
				Block:     v.CreationHeight,
			}

			items = append(items, tmp)
		}
	}

	return vo.RedelegationPage{
		Total: len(lcdReDelegations),
		Items: items,
	}
}

func (service *ValidatorService) GetWithdrawAddrByValidatorAddr(valAddr string) vo.WithdrawAddr {

	withdrawAddr, err := lcd.GetWithdrawAddressByValidatorAcc(utils.Convert(conf.Get().Hub.Prefix.AccAddr, valAddr))
	if err != nil {
		logger.Error("GetWithdrawAddressByValidatorAcc", logger.String("validator", valAddr), logger.String("err", err.Error()))
	}

	return vo.WithdrawAddr{
		Address: withdrawAddr,
	}
}

func (service *ValidatorService) GetDistributionRewardsByValidatorAddr(valAddr string) utils.CoinsAsStr {

	rewardsCoins, _, _, err := lcd.GetDistributionRewardsByValidatorAcc(utils.Convert(conf.Get().Hub.Prefix.AccAddr, valAddr))
	if err != nil {
		logger.Error("GetDistributionRewardsByValidatorAcc", logger.String("validator", valAddr), logger.String("err", err.Error()))
	}

	return rewardsCoins
}

func (service *ValidatorService) UpdateValidatorIcons() error {

	validatorsDocArr, err := document.Validator{}.GetAllValidator()
	if err != nil {
		return err
	}
	var txs []txn.Op
	for _, validator := range validatorsDocArr {
		if identity := validator.Description.Identity; identity != "" {
			urlicons, err := lcd.GetIconsByKey(identity)
			if err != nil || len(urlicons) == 0 {
				if err != nil {
					logger.Error("GetIconsByKey have error", logger.String("error", err.Error()))
				}
				continue
			}
			validator.Icons = urlicons
			txs = append(txs, txn.Op{
				C:  document.CollectionNmValidator,
				Id: validator.ID,
				Update: bson.M{
					"$set": validator,
				},
			})

		}

	}
	return document.Validator{}.Batch(txs)
}

func (service *ValidatorService) GetValidatorDetail(validatorAddr string) vo.ValidatorForDetail {

	validatorAsDoc, err := document.Validator{}.QueryValidatorDetailByOperatorAddr(validatorAddr)
	if err != nil {
		logger.Error("QueryValidatorDetailByOperatorAddr", logger.String("validator", validatorAddr), logger.String("err", err.Error()))
		return vo.ValidatorForDetail{}
	}

	desc := vo.Description{
		Moniker:  validatorAsDoc.Description.Moniker,
		Identity: validatorAsDoc.Description.Identity,
		Website:  validatorAsDoc.Description.Website,
		Details:  validatorAsDoc.Description.Details,
	}
	blackList := service.QueryBlackList()
	if blackone, ok := blackList[validatorAddr]; ok {
		desc.Moniker = blackone.Moniker
		desc.Identity = blackone.Identity
		desc.Website = blackone.Website
		desc.Details = blackone.Details
	}

	jailedUntil, missedBlockCount, startHeight, err := lcd.GetJailedUntilAndMissedBlocksCountByConsensusPublicKey(validatorAsDoc.ConsensusPubkey)

	var statsBlocksWindow string
	if err != nil {
		logger.Error("GetJailedUntilAndMissedBlocksCountByConsensusPublicKey", logger.String("consensus", validatorAsDoc.ConsensusPubkey), logger.String("err", err.Error()))
	} else {
		var startHeight, ok = utils.ParseInt(startHeight)
		if !ok {
			logger.Error("Format StartHeight", logger.String("err", err.Error()))
		} else {
			var lastBlock = lcd.BlockLatest()
			var currentHeight, ok = utils.ParseInt(lastBlock.BlockMeta.Header.Height)
			if !ok {
				logger.Error("Query CurrentHeight At LastBlock", logger.String("err", err.Error()))
			} else {
				signedBlocksWindow, err := document.GovParams{}.QueryOne("signed_blocks_window")
				if err != nil {
					logger.Error("Query signed_blocks_window", logger.String("err", err.Error()))
				} else {
					signedBlocksWindowCurrentValue, ok := utils.ParseInt(signedBlocksWindow.CurrentValue.(string))
					if ok {
						height := currentHeight - startHeight
						if height < signedBlocksWindowCurrentValue {
							statsBlocksWindow = strconv.FormatInt(height, 10)
						} else {
							statsBlocksWindow = strconv.FormatInt(signedBlocksWindowCurrentValue, 10)
						}
					}
				}
			}
		}
	}

	totalVotingPower, err := document.Validator{}.QueryTotalActiveValidatorVotingPower()

	if err != nil {
		logger.Error("QueryTotalActiveValidatorVotingPower", logger.String("err", err.Error()))
	}

	res := vo.ValidatorForDetail{
		TotalPower:              totalVotingPower,
		SelfPower:               validatorAsDoc.VotingPower,
		Status:                  validatorAsDoc.GetValidatorStatus(),
		BondedTokens:            validatorAsDoc.Tokens,
		SelfBonded:              ComputeSelfBonded(validatorAsDoc.Tokens, validatorAsDoc.DelegatorShares, validatorAsDoc.SelfBond),
		BondedStake:             ComputeBondStake(validatorAsDoc.Tokens, validatorAsDoc.DelegatorShares, validatorAsDoc.SelfBond),
		DelegatorShares:         validatorAsDoc.DelegatorShares,
		DelegatorCount:          validatorAsDoc.DelegatorNum,
		CommissionRate:          validatorAsDoc.Commission.Rate,
		CommissionUpdate:        validatorAsDoc.Commission.UpdateTime.String(),
		CommissionMaxRate:       validatorAsDoc.Commission.MaxRate,
		CommissionMaxChangeRate: validatorAsDoc.Commission.MaxChangeRate,
		BondHeight:              validatorAsDoc.BondHeight,
		MissedBlocksCount:       missedBlockCount,
		OperatorAddr:            validatorAsDoc.OperatorAddress,
		OwnerAddr:               utils.Convert(conf.Get().Hub.Prefix.AccAddr, validatorAsDoc.OperatorAddress),
		ConsensusAddr:           validatorAsDoc.ConsensusPubkey,
		Description:             desc,
		Icons:                   validatorAsDoc.Icons,
		Uptime:                  validatorAsDoc.Uptime,
		StatsBlocksWindow:       statsBlocksWindow,
	}
	if _, ok := blackList[validatorAddr]; ok {
		res.Icons = ""
	}

	if validatorAsDoc.Jailed {
		res.UnbondingHeight = validatorAsDoc.UnbondingHeight
		res.JailedUntil = jailedUntil
	}

	if validatorAsDoc.IsCandidatorWithStatus() {
		res.UnbondingHeight = validatorAsDoc.UnbondingHeight
		res.JailedUntil = jailedUntil
	}

	return res
}

func (service *ValidatorService) QueryCandidatesTopN() vo.ValDetailVo {

	validatorsList, power, upTimeMap, err := document.Validator{}.GetCandidatesTopN()

	if err != nil {
		logger.Error("GetCandidatesTopN have error", logger.String("error", err.Error()))
		panic(types.CodeNotFound)
	}

	var validators []vo.Validator

	for _, v := range validatorsList {

		validator := service.convert(v)
		validator.Uptime = upTimeMap[utils.GenHexAddrFromPubKey(v.ConsensusPubkey)]
		validators = append(validators, validator)
	}
	resp := vo.ValDetailVo{
		PowerAll:   power,
		Validators: validators,
	}

	return resp
}

func (service *ValidatorService) QueryValidator(address string) vo.CandidatesInfoVo {

	validator, err := lcd.Validator(address)
	if err != nil {
		logger.Error("lcd.Validator have error", logger.String("error", err.Error()))
		panic(types.CodeNotFound)
	}

	var moniker = validator.Description.Moniker
	var identity = validator.Description.Identity
	var website = validator.Description.Website
	var details = validator.Description.Details
	var blackList = service.QueryBlackList()
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		moniker = desc.Moniker
		identity = desc.Identity
		website = desc.Website
		details = desc.Details
	}
	var tokenDec, _ = types.NewDecFromStr(validator.Tokens)
	var val = vo.Validator{
		Address:     validator.OperatorAddress,
		PubKey:      validator.ConsensusPubkey,
		Owner:       utils.Convert(conf.Get().Hub.Prefix.AccAddr, validator.OperatorAddress),
		Jailed:      validator.Jailed,
		Status:      BondStatusToString(validator.Status),
		BondHeight:  utils.ParseIntWithDefault(validator.BondHeight, 0),
		VotingPower: tokenDec.RoundInt64(),
		Description: vo.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
		Rate:  validator.Commission.Rate,
		Icons: validator.Icons,
	}

	result := vo.CandidatesInfoVo{
		Validator: val,
	}

	count, err := document.Validator{}.QueryPowerWithBonded()

	if err != nil {
		logger.Error("query candidate power with bonded ", logger.String("err", err.Error()))
		return result
	}

	result.PowerAll = count
	return result
}
func ComputeSelfBonded(tokens, shares, selfBond string) string {
	if selfBond == "" {
		return "0"
	}
	rate, err := utils.QuoByStr(tokens, shares)
	if err != nil {
		logger.Error("validator.Tokens / validator.DelegatorShares", logger.String("err", err.Error()))
		return ""
	}

	selfBondAsRat, ok := new(big.Rat).SetString(selfBond)
	if !ok {
		logger.Error("convert validator selfBond type (string -> big.Rat) err",
			logger.String("self bond str", selfBond))
		return ""

	}
	selfBondTokensAsRat := new(big.Rat).Mul(selfBondAsRat, rate)

	return selfBondTokensAsRat.FloatString(18)
}

func ComputeBondStake(tokens, shares, selfBond string) string {
	rate, err := utils.QuoByStr(tokens, shares)
	if err != nil {
		logger.Error("validator.Tokens / validator.DelegatorShares", logger.String("err", err.Error()))
		return ""
	}

	tokensAsRat, ok := new(big.Rat).SetString(tokens)
	if !ok {
		logger.Error("convert validator tokens type (string -> big.Rat) err", logger.String("token str", tokens))
		return ""
	}

	if selfBond == "" {
		return tokensAsRat.FloatString(18)
	}
	selfBondAsRat, ok := new(big.Rat).SetString(selfBond)
	if !ok {
		logger.Error("convert validator selfBond type (string -> big.Rat) err", logger.String("self bond str", selfBond))
		return ""

	}
	selfBondTokensAsRat := new(big.Rat).Mul(selfBondAsRat, rate)
	BondStakeAsRat := new(big.Rat).Sub(tokensAsRat, selfBondTokensAsRat)

	return BondStakeAsRat.FloatString(18)
}

func (service *ValidatorService) QueryCandidateStatus(address string) (resp vo.ValStatus) {

	preCommitCount, uptime, err := document.Validator{}.QueryCandidateStatus(address)

	if err != nil {
		logger.Error("query candidate status", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	resp = vo.ValStatus{
		Uptime:         uptime,
		PrecommitCount: float64(preCommitCount),
	}

	return resp
}

func (service *ValidatorService) convert(validator document.Validator) vo.Validator {
	var moniker = validator.Description.Moniker
	var identity = validator.Description.Identity
	var website = validator.Description.Website
	var details = validator.Description.Details
	var blackList = service.QueryBlackList()
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		moniker = desc.Moniker
		identity = desc.Identity
		website = desc.Website
		details = desc.Details
	}

	bondHeightAsInt64, err := strconv.ParseInt(validator.BondHeight, 10, 64)

	if err != nil {
		logger.Error("convert string to int64", logger.String("err", err.Error()))
	}

	return vo.Validator{
		Address:     validator.OperatorAddress,
		PubKey:      utils.Convert(conf.Get().Hub.Prefix.ConsPub, validator.ConsensusPubkey),
		Owner:       utils.Convert(conf.Get().Hub.Prefix.AccAddr, validator.OperatorAddress),
		Jailed:      validator.Jailed,
		Status:      strconv.Itoa(validator.Status),
		BondHeight:  bondHeightAsInt64,
		VotingPower: validator.VotingPower,
		Description: vo.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
	}
}

func BondStatusToString(b int) string {
	switch b {
	case 0:
		return types.TypeValStatusUnbonded
	case 1:
		return types.TypeValStatusUnbonding
	case 2:
		return types.TypeValStatusBonded
	default:
		panic("improper use of BondStatusToString")
	}
}

func (service *ValidatorService) queryValForRainbow(page, size int) interface{} {
	var validators = lcd.Validators(page, size)

	var blackList = service.QueryBlackList()
	for i, v := range validators {
		if desc, ok := blackList[v.OperatorAddress]; ok {
			validators[i].Description.Moniker = desc.Moniker
			validators[i].Description.Identity = desc.Identity
			validators[i].Description.Website = desc.Website
			validators[i].Description.Details = desc.Details
		}
	}
	return validators
}
func (service *ValidatorService) UpdateMonikerMap(vs []document.Validator) {
	if service.monikerMap == nil {
		service.monikerMap = make(map[string]string, len(vs))
	}
	for _, v := range vs {
		if moniker, ok := service.monikerMap[v.OperatorAddress]; ok {
			if v.Description.Moniker != moniker {
				service.monikerMap[v.OperatorAddress] = v.Description.Moniker
			}
		} else {
			service.monikerMap[v.OperatorAddress] = v.Description.Moniker
		}
	}
}

func (service *ValidatorService) CleanMonikerMap() {
	for addr := range service.monikerMap {
		delete(service.monikerMap, addr)
	}
}

func (service *ValidatorService) UpdateValidators(vs []document.Validator) error {
	var vMap = make(map[string]document.Validator)
	for _, v := range vs {
		vMap[v.OperatorAddress] = v
	}

	var txs []txn.Op
	dstValidators := buildValidators()
	service.UpdateMonikerMap(dstValidators)
	for _, v := range dstValidators {
		if v1, ok := vMap[v.OperatorAddress]; ok {
			if isDiffValidator(v1, v) {
				v.ID = v1.ID
				v.Icons = v1.Icons
				// set staticInfo, see detail: buildValidatorStaticInfo
				v.Uptime = v1.Uptime
				v.SelfBond = v1.SelfBond
				v.DelegatorNum = v1.DelegatorNum
				txs = append(txs, txn.Op{
					C:  document.CollectionNmValidator,
					Id: v1.ID,
					Update: bson.M{
						"$set": v,
					},
				})
			}
			delete(vMap, v.OperatorAddress)
		} else {
			v.ID = bson.NewObjectId()
			txs = append(txs, txn.Op{
				C:      document.CollectionNmValidator,
				Id:     bson.NewObjectId(),
				Insert: v,
			})
		}
	}
	if len(vMap) > 0 {
		for addr := range vMap {
			v := vMap[addr]
			txs = append(txs, txn.Op{
				C:      document.CollectionNmValidator,
				Id:     v.ID,
				Remove: true,
			})
		}
	}
	return document.Validator{}.Batch(txs)
}

func (service *ValidatorService) UpdateValidatorStaticInfo() error {
	var validatorModel document.Validator
	validators, err := validatorModel.GetAllValidator()
	if err != nil {
		return err
	}
	validatorMap := make(map[string]document.Validator, len(validators))
	for _, v := range validators {
		validatorMap[v.OperatorAddress] = v
	}

	height := utils.ParseIntWithDefault(lcd.BlockLatest().BlockMeta.Header.Height, 0)
	updatedValidators := buildValidatorStaticInfo(validators, height)

	for _, v := range updatedValidators {
		if !isEqual(validatorMap[v.OperatorAddress], v) {
			if err := validatorModel.UpdateByPk(v); err != nil {
				logger.Error("update validator static data fail", logger.String("err", err.Error()),
					logger.String("validator", string(utils.MarshalJsonIgnoreErr(v))))
				continue
			}
		}
	}

	return nil
}

func (service *ValidatorService) QueryValidatorMonikerAndValidatorAddrByHashAddr(addr string) (document.Validator, error) {

	return document.Validator{}.QueryMonikerAndValidatorAddrByHashAddr(addr)
}

func (service *ValidatorService) QueryValidatorByConAddr(address string) document.Validator {

	validator, err := document.Validator{}.QueryValidatorByConsensusAddr(address)

	if err != nil {
		logger.Error("not found validator by conAddr", logger.String("conAddr", address))
	}
	return validator
}

func buildValidators() []document.Validator {

	res := lcd.Validators(1, 100)
	if res2 := lcd.Validators(2, 100); len(res2) > 0 {
		res = append(res, res2...)
	}

	var result []document.Validator
	var buildValidator = func(v lcd.ValidatorVo) (document.Validator, error) {
		var validator document.Validator
		if err := utils.Copy(v, &validator); err != nil {
			return validator, err
		}

		votingPower, err := types.NewDecFromStr(v.Tokens)
		if err == nil {
			validator.VotingPower = votingPower.RoundInt64()
		}
		validator.ProposerAddr = utils.GenHexAddrFromPubKey(validator.ConsensusPubkey)

		return validator, nil
	}

	for _, v := range res {
		if validator, err := buildValidator(v); err == nil {
			result = append(result, validator)
		} else {
			logger.Error("build validator fail", logger.String("err", err.Error()))
		}
	}
	return result
}

// update validator static info
// include uptime, selfBond, delegatorNum
func buildValidatorStaticInfo(validators []document.Validator, height int64) []document.Validator {
	if len(validators) == 0 {
		return nil
	}

	for i, v := range validators {
		v.Uptime = computeUptime(v.ConsensusPubkey, height)
		v.SelfBond, v.DelegatorNum = queryDelegationInfo(v.OperatorAddress)
		validators[i] = v
	}

	return validators
}

func computeUptime(valPub string, height int64) float32 {
	result := lcd.SignInfo(valPub)
	govSlashingParamMap, err := lcd.GetGovSlashingParam()
	if err != nil {
		logger.Error("GetGovSlashingParam have failed", logger.String("err", err.Error()))
	}
	startHeight := utils.ParseIntWithDefault(result.StartHeight, 0)

	var stats_blocks_window int64
	if _,ok := govSlashingParamMap["signed_blocks_window"]; ok {
		signed_blocks_window, ok := utils.ParseInt(govSlashingParamMap["signed_blocks_window"].(string))
		if !ok {
			stats_blocks_window = height - startHeight + 1
		} else {
			stats_blocks_window = Min(signed_blocks_window, height-startHeight+1)
		}
	}else{
		stats_blocks_window = height - startHeight + 1
	}


	missedBlocksCounter := utils.ParseIntWithDefault(result.MissedBlocksCounter, 0)

	tmp := float32(missedBlocksCounter) / float32(stats_blocks_window)
	return 1 - tmp
}

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func queryDelegationInfo(operatorAddress string) (string, int) {
	delegations := lcd.DelegationByValidator(operatorAddress)
	var selfBond string
	for _, d := range delegations {
		addr := utils.Convert(conf.Get().Hub.Prefix.AccAddr, operatorAddress)
		if d.DelegatorAddr == addr {
			selfBond = d.Shares
			break
		}
	}
	delegatorNum := len(delegations)
	return selfBond, delegatorNum
}

func isDiffValidator(src, dst document.Validator) bool {
	if src.OperatorAddress != dst.OperatorAddress ||
		src.ConsensusPubkey != dst.ConsensusPubkey ||
		src.Jailed != dst.Jailed ||
		src.Status != dst.Status ||
		src.Tokens != dst.Tokens ||
		src.DelegatorShares != dst.DelegatorShares ||
		src.BondHeight != dst.BondHeight ||
		src.UnbondingHeight != dst.UnbondingHeight ||
		src.UnbondingTime.Second() != dst.UnbondingTime.Second() ||
		src.VotingPower != dst.VotingPower ||
		src.ProposerAddr != dst.ProposerAddr ||
		src.Description.Moniker != dst.Description.Moniker ||
		src.Description.Identity != dst.Description.Identity ||
		src.Description.Website != dst.Description.Website ||
		src.Description.Details != dst.Description.Details ||
		src.Commission.Rate != dst.Commission.Rate ||
		src.Commission.MaxRate != dst.Commission.MaxRate ||
		src.Commission.MaxChangeRate != dst.Commission.MaxChangeRate ||
		src.Commission.UpdateTime.Second() != dst.Commission.UpdateTime.Second() {
		logger.Info("validator has changed", logger.String("OperatorAddress", src.OperatorAddress))
		return true
	}
	return false
}

func isEqual(srcValidator, dstValidator document.Validator) bool {
	srcBytes := utils.MarshalJsonIgnoreErr(srcValidator)
	dstBytes := utils.MarshalJsonIgnoreErr(dstValidator)
	if utils.Md5Encryption(srcBytes) == utils.Md5Encryption(dstBytes) {
		return true
	}

	return false
}

//func getVotingPowerFromToken(token string) int64 {
//	tokenPrecision := types.NewIntWithDecimal(1, 18)
//	power, err := types.NewDecFromStr(token)
//	if err != nil {
//		logger.Error("invalid token", logger.String("token", token))
//		return 0
//	}
//	return power.QuoInt(tokenPrecision).RoundInt64()
//}

func getTotalVotingPower() int64 {
	var total = int64(0)
	var set = lcd.LatestValidatorSet()
	for _, v := range set.Validators {
		votingPower := utils.ParseIntWithDefault(v.VotingPower, 0)
		total += votingPower
	}
	return total
}
