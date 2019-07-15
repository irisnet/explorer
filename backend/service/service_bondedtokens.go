package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/conf"
)

type BondedTokensService struct {
	BaseService
}

func (service *BondedTokensService) QueryBondedTokensValidator(vtype string) ([]*model.BondedTokensVo, error) {

	_, validators, err := document.Validator{}.GetValidatorListByPage(vtype, 0, 0, false)
	if err != nil {
		return nil, err
	}
	result := make([]*model.BondedTokensVo, 0, len(validators))
	blackList := service.QueryBlackList()
	for _, val := range validators {
		bondedtoken := &model.BondedTokensVo{
			Moniker:         val.Description.Moniker,
			Identity:        val.Description.Identity,
			VotingPower:     val.VotingPower,
			OperatorAddress: val.OperatorAddress,
		}
		if item, ok := blackList[val.OperatorAddress]; ok {
			bondedtoken.Moniker = item.Moniker
			bondedtoken.Identity = item.Identity
		}
		bondedtoken.OwnerAddress = utils.Convert(conf.Get().Hub.Prefix.AccAddr, val.OperatorAddress)

		result = append(result, bondedtoken)
	}

	return result, nil
}
