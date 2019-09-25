package document

import (
	"testing"
)

func TestQueryByPage(t *testing.T) {

	total, proposalList, err := Proposal{}.QueryByPage(0, 100, false)
	if err != nil {
		t.Error(err)
	}

	t.Logf("total: %v \n", total)

	for k, v := range proposalList {
		t.Logf("k: %v  proposal: %v \n", k, v)
	}

}

func TestQueryByIdList(t *testing.T) {

	proposalList, err := Proposal{}.QueryByIdList([]uint64{5, 6, 7, 8, 9, 10})
	if err != nil {
		t.Error(err)
	}

	for k, v := range proposalList {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestQueryById(t *testing.T) {

	proposal, err := Proposal{}.QueryById(5)
	if err != nil {
		t.Error(err)
	}

	t.Logf("proposal: %v \n", proposal)
}

func TestQuerySubmitProposalByHeight(t *testing.T) {

	txList, err := Proposal{}.QuerySubmitProposalByHeight(411768)

	if err != nil {
		t.Error(err)
	}

	for k, v := range txList {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryProposalById(t *testing.T) {
	proposal, err := Proposal{}.QueryProposalById(6)
	if err != nil {
		t.Error(err)
	}

	t.Logf("proposal: %v \n", proposal)

}

func TestQueryTxFromToByTypeAndProposalId(t *testing.T) {

	from, to, err := Proposal{}.QuerySubmitProposalTxByProposalId(4)

	if err != nil {
		t.Error(err)
	}

	t.Logf("from: %v to: %v \n", from, to)

}
