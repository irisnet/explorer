package service

import (
	"testing"

	"github.com/irisnet/explorer/backend/model"
)

func TestQueryProposalsByHeight(t *testing.T) {

	proposalList := new(ProposalService).QueryProposalsByHeight(426520)

	for k, v := range proposalList {
		t.Logf("idx: %v  v: %v \n", k, v)
	}
}

func TestProposalQueryList(t *testing.T) {

	proposalPage := new(ProposalService).QueryList(0, 100)

	t.Logf("total: %v \n", proposalPage.Count)

	if modelV, ok := proposalPage.Data.([]model.Proposal); ok {
		for k, v := range modelV {
			t.Logf("k: %v   v: %v \n", k, v)
		}
	}
}

func TestProposalQuery(t *testing.T) {

	proposal := new(ProposalService).Query(15)

	t.Logf("proposal: %v \n", proposal)
}
