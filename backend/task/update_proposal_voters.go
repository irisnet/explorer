package task

import (
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/logger"
	"fmt"
	"github.com/irisnet/explorer/backend/orm/document"
)

type UpdateProposalVoters struct{}

func (task UpdateProposalVoters) Name() string {
	return "update_proposal_voters"
}

func (task UpdateProposalVoters) Start() {
	utils.RunTimer(conf.Get().Server.CronTimeProposalVoters, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})

}

func (task UpdateProposalVoters) DoTask() error {
	status := []string{document.ProposalStatusVoting}
	sorts := []string{document.Proposal_Field_VotingEndTime}
	proposals, err := document.Proposal{}.GetProposalsByStatus(status, sorts)
	if err != nil {
		return err
	}
	proposalService := service.ProposalService{}
	proposalService.UpdateProposalVoters(proposals)

	return nil
}
