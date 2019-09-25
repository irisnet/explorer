package service

import (
	"testing"

	"github.com/irisnet/explorer/backend/vo"
)

//func TestGetValidatorPublicKeyMonikerFromProposalVoter(t *testing.T) {
//
//	aa := "faa17cjdg63thy2vfqvvgj5lfv5dp339t0lr99wc8p"
//
//	multiType, err := ProposalService{}.GetValidatorPublicKeyMonikerFromProposalVoter([]string{aa})
//	if err != nil {
//		t.Error(err)
//	}
//
//	for k, v := range multiType {
//		t.Logf("k: %v  va: %v  consensusPub: %v  consensusHex: %v \n", k, v.Va, v.ConsensusPubKey, v.ConsensusHex)
//	}
//
//}

//func TestQueryDepositAndVotingProposalList(t *testing.T) {
//
//	proposalList := new(ProposalService).QueryDepositAndVotingProposalList()
//
//	for k, v := range proposalList {
//		t.Logf("k: %v id: %v  proposal vote end time: %v   deposit end time: %v \n", k, v.ProposalId, v.VotingEndTime, v.DepositEndTime)
//	}
//
//}

func TestQueryList(t *testing.T) {

	proposalPage := new(ProposalService).QueryList(0, 10, true)

	t.Logf("total: %v \n", proposalPage.Count)

	if pV, ok := proposalPage.Data.([]vo.ProposalNewStyle); ok {
		for k, v := range pV {
			t.Logf("k: %v  v: %v \n", k, v)
		}
	}

}
