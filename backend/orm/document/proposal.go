package document

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmProposal = "proposal"

	Proposal_Field_ProposalId        = "proposal_id"
	Proposal_Field_Title             = "title"
	Proposal_Field_Type              = "type"
	Proposal_Field_Description       = "description"
	Proposal_Field_Status            = "status"
	Proposal_Field_SubmitTime        = "submit_time"
	Proposal_Field_DepositEndTime    = "deposit_end_time"
	Proposal_Field_VotingStartTime   = "voting_start_time"
	Proposal_Field_VotingEndTime     = "voting_end_time"
	Proposal_Field_TotalDeposit      = "total_deposit"
	Proposal_Field_Votes             = "votes"
	Proposal_Field_Votes_voter       = "votes.voter"
	Proposal_Field_Votes_Option      = "votes.option"
	Proposal_Field_Votes_TxHash      = "votes.txhash"
	Proposal_Field_VotingStartHeight = "voting_start_height"
	Proposal_Field_Final_Votes       = "tally_result"

	ProposalStatusDeposit  = "DepositPeriod"
	ProposalStatusVoting   = "VotingPeriod"
	ProposalStatusPassed   = "Passed"
	ProposalStatusRejected = "Rejected"
)

type FinalVoteAsPower struct {
	Yes               string `bson:"yes"`
	No                string `bson:"no"`
	NoWithVeto        string `bson:"nowithveto"`
	Abstain           string `bson:"abstain"`
	SystemVotingPower string `bson:"system_voting_power"`
}

type Proposal struct {
	ProposalId        uint64           `bson:"proposal_id"`
	Title             string           `bson:"title"`
	Type              string           `bson:"type"`
	Description       string           `bson:"description"`
	Status            string           `bson:"status"`
	SubmitTime        time.Time        `bson:"submit_time"`
	DepositEndTime    time.Time        `bson:"deposit_end_time"`
	VotingStartTime   time.Time        `bson:"voting_start_time"`
	VotingEndTime     time.Time        `bson:"voting_end_time"`
	TotalDeposit      Coins            `bson:"total_deposit"`
	Votes             []PVote          `bson:"votes"`
	VotingStartHeight int64            `bson:"voting_start_height"`
	FinalVotes        FinalVoteAsPower `bson:"tally_result"`
}

func (p Proposal) String() string {
	return fmt.Sprintf(`
		ProposalId        :%v
		Title             :%v
		Type              :%v
		Description       :%v
		Status            :%v
		SubmitTime        :%v
		DepositEndTime    :%v
		VotingStartTime   :%v
		VotingEndTime     :%v
		TotalDeposit      :%v
		Votes             :%v
		VotingStartHeight :%v
		FinalVotes        :%v
		`, p.ProposalId, p.Title, p.Type, p.Description, p.Status, p.SubmitTime, p.DepositEndTime, p.VotingStartTime, p.VotingEndTime, p.TotalDeposit, p.Votes, p.VotingStartHeight, p.FinalVotes)
}

type PVote struct {
	Voter  string    `json:"voter"`
	Option string    `json:"option"`
	Time   time.Time `json:"time"`
	TxHash string    `json:"txhash"`
}

func (p PVote) String() string {
	return fmt.Sprintf(`Voter: %v  option: %v  time: %v`, p.Voter, p.Option, p.Time)
}

func (m Proposal) Name() string {
	return CollectionNmProposal
}

func (m Proposal) PkKvPair() map[string]interface{} {
	return bson.M{Proposal_Field_ProposalId: m.ProposalId}
}

func (_ Proposal) QueryByPage(page, size int, total bool) (int, []Proposal, error) {

	var data []Proposal
	sort := desc(Proposal_Field_SubmitTime)
	cnt, err := pageQuery(CollectionNmProposal, nil, nil, sort, page, size, total, &data)
	return cnt, data, err
}

func (_ Proposal) QueryByIdList(idList []uint64) ([]Proposal, error) {

	proposalArr := []Proposal{}

	depositVoteCondition := bson.M{
		Proposal_Field_ProposalId: bson.M{"$in": idList},
	}

	err := queryAll(CollectionNmProposal, nil, depositVoteCondition, "", 0, &proposalArr)

	return proposalArr, err
}

func (_ Proposal) QueryById(id int64) (Proposal, error) {
	var proposal Proposal
	selector := bson.M{"proposal_id": 1, "title": 1, "type": 1, "status": 1, "submit_time": 1}
	var condition = bson.M{Proposal_Field_ProposalId: id}

	var query = orm.NewQuery()
	query.Reset().
		SetCollection(CollectionNmProposal).
		SetSelector(selector).
		SetCondition(condition).
		SetResult(&proposal)

	err := query.Exec()

	return proposal, err
}

func (_ Proposal) QuerySubmitProposalByHeight(height int64) ([]CommonTx, error) {
	var query = orm.NewQuery()
	var data []CommonTx

	var selector = bson.M{"proposal_id": 1, "type": 1, "from": 1}

	query.SetCollection(CollectionNmCommonTx).
		SetSelector(selector).
		SetCondition(bson.M{Tx_Field_Height: height, Tx_Field_Type: "SubmitProposal"}).
		SetResult(&data)
	err := query.Exec()
	return data, err
}

func (_ Proposal) QueryProposalById(id int) (Proposal, error) {
	var query = orm.NewQuery()
	defer query.Release()

	var data Proposal
	query.SetCollection(CollectionNmProposal).
		SetCondition(bson.M{Proposal_Field_ProposalId: id}).
		SetResult(&data)

	err := query.Exec()

	return data, err
}

func (_ Proposal) QuerySubmitProposalTxByProposalId(id int) (string, string, error) {
	var query = orm.NewQuery()
	defer query.Release()
	var tx CommonTx
	query.Reset().SetCollection(CollectionNmCommonTx).
		SetCondition(bson.M{Tx_Field_Type: types.TxTypeSubmitProposal, Proposal_Field_ProposalId: id}).
		SetResult(&tx)
	err := query.Exec()
	return tx.From, tx.TxHash, err
}

func (_ Proposal) QueryProposalByPage(page, size int, total bool) (int, []Proposal, error) {

	var data []Proposal
	sort := desc(Proposal_Field_ProposalId)
	selector := bson.M{
		Proposal_Field_ProposalId:        1,
		Proposal_Field_Title:             1,
		Proposal_Field_Type:              1,
		Proposal_Field_Status:            1,
		Proposal_Field_SubmitTime:        1,
		Proposal_Field_DepositEndTime:    1,
		Proposal_Field_VotingEndTime:     1,
		Proposal_Field_Votes:             1,
		Proposal_Field_VotingStartHeight: 1,
		Proposal_Field_Final_Votes:       1,
	}

	cnt, err := pageQuery(CollectionNmProposal, selector, nil, sort, page, size, total, &data)

	return cnt, data, err
}

func (_ Proposal) QueryIdTitleStatusVotedTxhashByValidatorAcc(validatorAcc string, page, size int, total bool) (int, []Proposal, error) {

	var data []Proposal

	sort := desc(Proposal_Field_ProposalId)
	selector := bson.M{
		Proposal_Field_ProposalId: 1,
		Proposal_Field_Title:      1,
		Proposal_Field_Status:     1,
		Proposal_Field_Votes:      1,
	}

	condition := bson.M{Proposal_Field_Votes_voter: validatorAcc}

	cnt, err := pageQuery(CollectionNmProposal, selector, condition, sort, page, size, total, &data)

	if err != nil {
		return 0, nil, err
	}

	return cnt, data, err
}

func (_ Proposal) GetProposalsByStatus(status, sorts []string) ([]Proposal, error) {
	var (
		proposals []Proposal
		query     = orm.NewQuery()
	)
	defer query.Release()

	selector := bson.M{
		Proposal_Field_ProposalId:        1,
		Proposal_Field_Title:             1,
		Proposal_Field_Type:              1,
		Proposal_Field_Status:            1,
		Proposal_Field_Votes:             1,
		Proposal_Field_VotingStartHeight: 1,
		Proposal_Field_TotalDeposit:      1,
		Proposal_Field_DepositEndTime:    1,
	}

	condition := bson.M{
		"status": bson.M{
			"$in": status,
		},
	}
	sorts = append(sorts, desc(Proposal_Field_ProposalId))

	query.SetCollection(CollectionNmProposal).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sorts...).
		SetSize(0).
		SetResult(&proposals)

	err := query.Exec()

	return proposals, err
}
