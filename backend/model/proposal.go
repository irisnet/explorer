package model

import (
	"github.com/irisnet/irishub-sync/store"
	"time"
)

type Proposal struct {
	Title           string      `json:"title"`
	ProposalId      uint64      `json:"proposal_id"`
	Type            string      `json:"type"`
	Description     string      `json:"description"`
	Status          string      `json:"status"`
	SubmitTime      string      `bson:"submit_time"`
	DepositEndTime  string      `bson:"deposit_end_time"`
	VotingStartTime string      `bson:"voting_start_time"`
	VotingEndTime   string      `bson:"voting_end_time"`
	TotalDeposit    store.Coins `json:"total_deposit"`
	Proposer        string      `json:"proposer"`
	TxHash          string      `json:"tx_hash"`
}

type Vote struct {
	Voter  string    `json:"voter"`
	Option string    `json:"option"`
	Time   time.Time `json:"time"`
}

type ProposalInfo struct {
	Proposal Proposal   `json:"proposal"`
	Votes    []Vote     `json:"votes"`
	Result   VoteResult `json:"result"`
}

type VoteResult struct {
	Yes        int
	No         int
	NoWithVeto int
	Abstain    int
}
