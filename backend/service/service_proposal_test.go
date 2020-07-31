package service

import (
	"encoding/json"
	"testing"

	"github.com/irisnet/explorer/backend/vo"
)

func TestQueryProposalsByHeight(t *testing.T) {

	proposalList := new(ProposalService).QueryProposalsByHeight(426520)

	for k, v := range proposalList {
		t.Logf("idx: %v  v: %v \n", k, v)
	}
}

func TestProposalQueryList(t *testing.T) {

	proposalPage := new(ProposalService).QueryList(0, 100, true)

	t.Logf("total: %v \n", proposalPage.Count)

	if modelV, ok := proposalPage.Data.([]vo.Proposal); ok {
		for k, v := range modelV {
			t.Logf("k: %v   v: %v \n", k, v)
		}
	}
}

func TestProposalQuery(t *testing.T) {

	proposal := new(ProposalService).Query(2)

	t.Logf("proposal: %v \n", proposal)
}

func TestProposalService_GetVoteTxs(t *testing.T) {
	res := proposalService.GetVoteTxs(6, 1, 10, true,"all")
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestProposalService_GetDepositTxs(t *testing.T) {
	res := proposalService.GetDepositTxs(8, 1, 10, true)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))

}

//func TestProposalService_QueryDepositAndVotingProposalList(t *testing.T) {
//	res := proposalService.QueryDepositAndVotingProposalList()
//	resBytes, _ := json.Marshal(res)
//	t.Log(string(resBytes))
//}

func TestProposalService_QueryList(t *testing.T) {
	res := proposalService.QueryList(1, 20, true)
	resBytes, _ := json.Marshal(res)
	t.Log(string(resBytes))
}

func TestProposalService_GetSystemVotingPower(t *testing.T) {
	if res, err := proposalService.GetSystemVotingPower(); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.Marshal(res)
		t.Log(string(resBytes))
	}
}

func TestProposalService_QueryDepositAndVotingProposalList(t *testing.T) {
	res := proposalService.QueryDepositAndVotingProposalList(true)

	resBytes, _ := json.Marshal(res)
	t.Log(string(resBytes))
}
