package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/logger"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/utils"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type ValidatorStaticByDayTask struct {
	validator document.Validator
}

func (task ValidatorStaticByDayTask) Name() string {
	return "validator_static_task"
}
func (task ValidatorStaticByDayTask) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeTxNumByDay

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task ValidatorStaticByDayTask) DoTask() error {
	ops, err := task.getAllValidatorTokens()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return document.ExValidatorStatic{}.Batch(ops)
}

func (task ValidatorStaticByDayTask) getAllValidatorTokens() ([]txn.Op, error) {

	validators, err := task.getValidatorFromDb()
	if err != nil {
		return nil, err
	}

	return task.saveExValidatorStaticOps(validators)
}

func (task ValidatorStaticByDayTask) getValidatorFromDb() ([]document.Validator, error) {
	return task.validator.GetAllValidator()
}

func (task ValidatorStaticByDayTask) saveExValidatorStaticOps(validators []document.Validator) ([]txn.Op, error) {
	today := utils.TruncateTime(time.Now(), utils.Day)
	ops := make([]txn.Op, 0, len(validators))
	for _, addr := range validators {
		item, err := task.loadValidatorTokens(addr, today)
		if err != nil {
			continue
		}
		op := txn.Op{
			C:      document.CollectionNameExValidatorStatic,
			Id:     bson.NewObjectId(),
			Insert: item,
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (task ValidatorStaticByDayTask) loadValidatorTokens(validator document.Validator, today time.Time) (document.ExValidatorStatic, error) {

	item := document.ExValidatorStatic{
		Id:              bson.NewObjectId(),
		OperatorAddress: validator.OperatorAddress,
		Date:            today,
		SelfBond:        validator.SelfBond,
		DelegatorShares: validator.DelegatorShares,
		Tokens:          validator.Tokens,
		Commission:      validator.Commission,
	}
	item.Delegations = funcSubStr(item.Tokens, item.SelfBond).FloatString(18)
	return item, nil
}
