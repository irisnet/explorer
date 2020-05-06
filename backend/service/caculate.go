package service

import (
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"math"
	"strconv"
)

type CaculateService struct {
	BaseService
}

func (s CaculateService) GetDelegatorCaculateMonth(cond bson.M, page, size int, istotal bool) ([]vo.ExStaticDelegatorMonthVo, int, error) {
	datas, total, err := delegatorMonthModel.List(cond, page, size, istotal)
	res := make([]vo.ExStaticDelegatorMonthVo, 0, len(datas))
	for _, val := range datas {
		item := LoadFromDelModel(val)
		res = append(res, item)
	}
	return res, total, err
}

func (s CaculateService) GetValidatorCaculateMonth(cond bson.M, page, size int, istotal bool) ([]vo.ExStaticValidatorMonthVo, int, error) {
	datas, total, err := validatorMonthModel.List(cond, page, size, istotal)
	res := make([]vo.ExStaticValidatorMonthVo, 0, len(datas))
	for _, val := range datas {
		item := LoadFromValModel(val)
		res = append(res, item)
	}
	return res, total, err
}

func LoadFromDelModel(model document.ExStaticDelegatorMonth) vo.ExStaticDelegatorMonthVo {
	ret := vo.ExStaticDelegatorMonthVo{
		Address:                model.Address,
		Date:                   model.Date,
		CaculateDate:           model.CaculateDate,
		TerminalRewards:        model.TerminalRewards.Iris,
		PeriodIncrementRewards: model.PeriodIncrementRewards.Iris,
		PeriodWithdrawRewards:  model.PeriodWithdrawRewards.Iris,
		TerminalDelegation:     coverAmount(model.TerminalDelegation),
		IncrementDelegation:    coverAmount(model.IncrementDelegation),
		PeriodDelegationTimes:  model.PeriodDelegationTimes,
	}
	return ret
}

func coverAmount(coin document.Coin) float64 {
	if coin.Denom == types.IRISAttoUint {
		coin.Amount = coin.Amount / math.Pow10(18)
	}
	return coin.Amount
}
func LoadFromValModel(model document.ExStaticValidatorMonth) vo.ExStaticValidatorMonthVo {
	coverFloat64 := func(data string) float64 {
		ret, _ := strconv.ParseFloat(data, 64)
		return ret
	}

	ret := vo.ExStaticValidatorMonthVo{
		Address:               model.Address,
		OperatorAddress:       model.OperatorAddress,
		Status:                model.Status,
		CreateValidatorHeight: model.CreateValidatorHeight,
		Date:                  model.Date,
		CaculateDate:          model.CaculateDate,
		TerminalCommission:    coverAmount(model.TerminalCommission),
		PeriodCommission:      coverAmount(model.PeriodCommission),
		IncrementCommission:   coverAmount(model.IncrementCommission),
		TerminalDelegation:    coverFloat64(model.TerminalDelegation),
		IncrementDelegation:   coverFloat64(model.IncrementDelegation),
		Tokens:                coverFloat64(model.Tokens),
		TerminalDelegatorN:      model.TerminalDelegatorN,
		IncrementDelegatorN:     model.IncrementDelegatorN,
		TerminalSelfBond:        coverFloat64(model.TerminalSelfBond),
		IncrementSelfBond:       coverFloat64(model.IncrementSelfBond),
		CommissionRateMax:       coverFloat64(model.CommissionRateMax),
		CommissionRateMin:       coverFloat64(model.CommissionRateMin),
		FoundationDelegateT:     coverFloat64(model.FoundationDelegateT),
		FoundationDelegateIncre: coverFloat64(model.FoundationDelegateIncre),
		Moniker:                 model.Moniker,
	}
	return ret
}
