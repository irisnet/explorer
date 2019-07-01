package service

import (
	"encoding/json"
	"testing"
)

func TestProposalService_GetVoteTxs(t *testing.T) {
	res := proposalService.GetVoteTxs(1, 1, 10)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestProposalService_GetDepositTxs(t *testing.T) {
	res := proposalService.GetDepositTxs(1, 1, 10)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestProposalService_QueryDepositAndVotingProposalList(t *testing.T) {
	res := proposalService.QueryDepositAndVotingProposalList()
	resBytes, _ := json.Marshal(res)
	t.Log(string(resBytes))
}

func TestProposalService_QueryList(t *testing.T) {
	res := proposalService.QueryList(1, 20)
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
