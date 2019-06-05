package service

import (
	"testing"
)

func TestGetValidatorPublicKeyMonikerFromProposalVoter(t *testing.T) {
	aa := "faa17cjdg63thy2vfqvvgj5lfv5dp339t0lr99wc8p"

	multiType, err := ProposalService{}.GetValidatorPublicKeyMonikerFromProposalVoter([]string{aa})
	if err != nil {
		t.Error(err)
	}

	for k, v := range multiType {
		t.Logf("k: %v   aa: %v  va: %v  consensusPub: %v  consensusHex: %v \n", k, v.AA, v.Va, v.ConsensusPubKey, v.ConsensusHex)
	}

}
