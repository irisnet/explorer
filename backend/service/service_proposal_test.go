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
