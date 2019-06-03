package service

import (
	"fmt"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type GovParamsService struct {
	BaseService
}

func (service *GovParamsService) GetModule() Module {
	return GovParams
}

func (service *GovParamsService) QueryAll() []document.GovParams {
	params, err := document.GovParams{}.QueryAll()

	if err != nil {
		panic(types.CodeNotFound)
	}
	return params
}

func init() {
	var initParams = func() {
		var ops []txn.Op
		gen, _ := lcd.Genesis()

		//auth
		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "auth",
				Key:    "gas_price_threshold",
				Value:  gen.Result.Genesis.AppState.Auth.Params.GasPriceThreshold,
				Type:   "sdk.Int",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.EQ,
						Data: "0",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "1000000000000000000",
					},
				},
				Description: "the minimum of gas price",
			},
		})

		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "auth",
				Key:    "tx_size",
				Value:  gen.Result.Genesis.AppState.Auth.Params.TxSize,
				Type:   "int",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.EQ,
						Data: "500",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "1500",
					},
				},
				Description: "transaction size limit",
			},
		})

		//stake
		var maxValidators = fmt.Sprintf("%d", gen.Result.Genesis.AppState.Stake.Params.MaxValidators)
		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "stake",
				Key:    "max_validators",
				Value:  maxValidators,
				Type:   "int",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.EQ,
						Data: "100",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "200",
					},
				},
				Description: "the maximum number of validators",
			},
		})

		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "stake",
				Key:    "unbonding_time",
				Value:  gen.Result.Genesis.AppState.Stake.Params.UnbondingTime,
				Type:   "time.Duration",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.EQ,
						Data: "1209600000000000",
					},
				},
				Description: "unbonding time",
				Note:        "the locking time of unbonding and redelegation",
			},
		})

		//mint
		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "mint",
				Key:    "inflation",
				Value:  gen.Result.Genesis.AppState.Mint.Params.Inflation,
				Type:   "sdk.Dec",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.EQ,
						Data: "0",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "0.2",
					},
				},
				Description: "inflation rate",
			},
		})

		//distr
		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "distr",
				Key:    "community_tax",
				Value:  gen.Result.Genesis.AppState.Distr.Params.CommunityTax,
				Type:   "sdk.Dec",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.NEQ,
						Data: "0",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "0.2",
					},
				},
				Description: "community tax",
			},
		})

		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "distr",
				Key:    "base_proposer_reward",
				Value:  gen.Result.Genesis.AppState.Distr.Params.BaseProposerReward,
				Type:   "sdk.Dec",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.NEQ,
						Data: "0",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "0.02",
					},
				},
				Description: "base ratio of the block reward",
			},
		})

		ops = append(ops, txn.Op{
			C:  document.CollectionNmGovParams,
			Id: bson.NewObjectId(),
			Insert: document.GovParams{
				Module: "distr",
				Key:    "bonus_proposer_reward",
				Value:  gen.Result.Genesis.AppState.Distr.Params.BonusProposerReward,
				Type:   "sdk.Dec",
				Range: document.Range{
					Minimum: document.Op{
						Sign: document.NEQ,
						Data: "0",
					},
					Maximum: document.Op{
						Sign: document.EQ,
						Data: "0.08",
					},
				},
				Description: "maximum additional ratio bonus ratio",
			},
		})

		err := document.GovParams{}.Batch(ops)

		if err != nil {
			logger.Error("init gov_params data error", logger.String("err", err.Error()))
		}
	}
	if data := govParamsService.QueryAll(); len(data) == 0 {
		initParams()
	}
}
