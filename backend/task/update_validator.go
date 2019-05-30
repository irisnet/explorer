package task

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateValidator struct{}

func (task UpdateValidator) Name() string {
	return "update_validator"
}
func (task UpdateValidator) Start() {
	utils.RunTimer(30, utils.Sec, func() {
		var validators []document.Validator
		var query = orm.NewQuery()
		defer query.Release()
		query.SetCollection(document.CollectionNmValidator).
			SetResult(&validators)
		if err := query.Exec(); err != nil {
			logger.Error("queryValidators failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}

		candidateService := service.Get(service.Candidate).(*service.CandidateService)
		err := candidateService.UpdateValidators(validators)
		if err != nil {
			logger.Error("UpdateValidators task failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
		}
	})

}
