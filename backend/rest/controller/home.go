package controller

import (
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterHome(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerNavigationBar,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// @Summary navigation
// @Description get navigation
// @Tags home
// @Accept  json
// @Produce  json
// @Success 200 {object} NavigationData   "success"
// @Router /api/home/navigation [get]
func registerNavigationBar(r *mux.Router) error {
	doApi(r, types.UrlRegisterNavigationBar, "GET", func(request vo.IrisReq) interface{} {
		var block = lcd.BlockLatest()
		if block.BlockMeta.Header.Height == "" {
			return nil
		}
		var height, ok = utils.ParseInt(block.BlockMeta.Header.Height)
		if !ok {
			panic(types.CodeNotFound)
		}
		var result = NavigationData{
			BlockHeight: height,
			TotalTxs:    utils.ParseIntWithDefault(block.BlockMeta.Header.TotalTxs, 0),
			BlockTime:   block.BlockMeta.Header.Time,
		}
		var proposer = block.BlockMeta.Header.ProposerAddress
		var funGroup []func()
		var group sync.WaitGroup
		var queryVoteInfo = func() {
			defer func() {
				group.Done()
				if r := recover(); r != nil {
					logger.Error("queryVoteInfo error", logger.Any("err", r))
				}
			}()
			var validatorSet = lcd.ValidatorSet(height - 1)

			var totalVp = int64(0)
			var voteVp = int64(0)
			var voteValNum = 0
			for i, commit := range block.Block.LastCommit.Precommits {
				var vp = utils.ParseIntWithDefault(validatorSet.Validators[i].VotingPower, 0)
				if len(commit.Height) > 0 {
					voteVp += vp
					voteValNum++
				}
				totalVp += vp
			}
			result.VotingRatio = float32(voteVp) / float32(totalVp)
			result.VotingTokens = strconv.FormatInt(voteVp, 10)
			result.VoteValNum = voteValNum
			result.ActiveValNum = len(validatorSet.Validators)

			validator, err := stake.QueryValidatorMonikerAndValidatorAddrByHashAddr(proposer)
			if err != nil {
				logger.Error("query validator moniker and addr ", logger.String("err", err.Error()))
			}

			result.Moniker = validator.Description.Moniker
			result.OperatorAddr = validator.OperatorAddress
			result.ValidatorIcon = validator.Icons
		}
		funGroup = append(funGroup, queryVoteInfo)
		group.Add(1)

		var queryBondedInfo = func() {
			defer func() {
				group.Done()
				if r := recover(); r != nil {
					logger.Error("queryBondedInfo error", logger.Any("err", r))
				}
			}()
			var poolStake = lcd.StakePool()
			result.BondedTokens = poolStake.BondedTokens
			result.BondedRatio = poolStake.BondedRatio
			result.TotalSupply = poolStake.TotalSupply
		}
		funGroup = append(funGroup, queryBondedInfo)
		group.Add(1)

		//var getTokenStatCirculation = func() {
		//	defer func() {
		//		group.Done()
		//		if r := recover(); r != nil {
		//			logger.Error("getTokenStatCirculation error", logger.Any("err", r))
		//		}
		//	}()
		//	if v, err := lcd.GetTokenStatsCirculation(); err != nil {
		//		logger.Error("GetTokenStatsCirculation fail", logger.String("err", err.Error()))
		//		result.Circulation = "0"
		//	} else {
		//		result.Circulation = v.Amount
		//	}
		//}
		//funGroup = append(funGroup, getTokenStatCirculation)
		//group.Add(1)

		//var getTokenStatFoundationBonded = func() {
		//defer func() {
		//	group.Done()
		//	if r := recover(); r != nil {
		//		logger.Error("getTokenStatFoundationBonded error", logger.Any("err", r))
		//	}
		//}()

		//accService := service.AccountService{}
		//if conf.Get().Hub.Prefix.AccAddr == types.MainnetAccPrefix {
		//res := accService.QueryDelegations(types.FoundationDelegatorAddr)
		//foundationBondAmt := float64(0)
		//for _, v := range res {
		//	foundationBondAmt += v.Amount.Amount
		//}
		//result.FoundationBonded = utils.ParseStringFromFloat64(foundationBondAmt)
		//}
		//}
		//funGroup = append(funGroup, getTokenStatFoundationBonded)
		//group.Add(1)

		var calculateAvgBlockTime = func() {
			defer func() {
				group.Done()
				if r := recover(); r != nil {
					logger.Error("calculateAvgBlockTime error", logger.Any("err", r))
				}
			}()
			var startHeight = int64(1)
			var period = int64(100)
			if height > 100 {
				startHeight = height - 100
			} else {
				period = height
			}
			var tmpBlock = lcd.Block(startHeight)
			var bwtSeconds = block.BlockMeta.Header.Time.Sub(tmpBlock.BlockMeta.Header.Time).Seconds()
			result.AvgBlockTime = float32(bwtSeconds) / float32(period)
		}
		funGroup = append(funGroup, calculateAvgBlockTime)
		group.Add(1)

		for _, fun := range funGroup {
			go fun()
		}

		group.Wait()
		return result
	})
	return nil
}

type NavigationData struct {
	BlockHeight      int64     `json:"block_height"`
	BlockTime        time.Time `json:"block_time"`
	TotalTxs         int64     `json:"total_txs"`
	Tps              float32   `json:"tps"`
	AvgBlockTime     float32   `json:"avg_block_time"`
	VotingRatio      float32   `json:"voting_ratio"`
	VotingTokens     string    `json:"voting_tokens"`
	VoteValNum       int       `json:"vote_val_num"`
	ActiveValNum     int       `json:"active_val_num"`
	BondedRatio      string    `json:"bonded_ratio"`
	BondedTokens     string    `json:"bonded_tokens"`
	TotalSupply      string    `json:"total_supply"`
	Circulation      string    `json:"circulation"`
	FoundationBonded string    `json:"foundation_bonded"`
	Moniker          string    `json:"moniker"`
	OperatorAddr     string    `json:"operator_addr"`
	ValidatorIcon    string    `json:"validator_icon"`
}
