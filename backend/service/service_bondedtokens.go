package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
)

type BondedTokensService struct {
	BaseService
}

func (service *BondedTokensService) QueryBondedTokensValidator() ([]*model.BondedTokensVo, error) {

	validators, err := document.Validator{}.GetAllValidator()
	if err != nil {
		return nil, err
	}
	result := make([]*model.BondedTokensVo, 0, len(validators))
	if blackList := service.QueryBlackList(); len(blackList) > 0 {
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

			result = append(result, bondedtoken)
		}
	}

	return result, nil
}
