package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNmProposal = "proposal"

	Proposal_Field_ProposalId      = "proposal_id"
	Proposal_Field_Title           = "title"
	Proposal_Field_Type            = "type"
	Proposal_Field_Description     = "description"
	Proposal_Field_Status          = "status"
	Proposal_Field_SubmitTime      = "submit_time"
	Proposal_Field_DepositEndTime  = "deposit_end_time"
	Proposal_Field_VotingStartTime = "voting_start_time"
	Proposal_Field_VotingEndTime   = "voting_end_time"
	Proposal_Field_TotalDeposit    = "total_deposit"
	Proposal_Field_Votes           = "votes"
)

type Proposal struct {
	ProposalId      uint64    `bson:"proposal_id"`
	Title           string    `bson:"title"`
	Type            string    `bson:"type"`
	Description     string    `bson:"description"`
	Status          string    `bson:"status"`
	SubmitTime      time.Time `bson:"submit_time"`
	DepositEndTime  time.Time `bson:"deposit_end_time"`
	VotingStartTime time.Time `bson:"voting_start_time"`
	VotingEndTime   time.Time `bson:"voting_end_time"`
	TotalDeposit    Coins     `bson:"total_deposit"`
	Votes           []PVote   `bson:"votes"`
}

type PVote struct {
	Voter  string    `json:"voter"`
	Option string    `json:"option"`
	Time   time.Time `json:"time"`
}

func (m Proposal) Name() string {
	return CollectionNmProposal
}

func (m Proposal) PkKvPair() map[string]interface{} {
	return bson.M{Proposal_Field_ProposalId: m.ProposalId}
}
