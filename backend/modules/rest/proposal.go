package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

func RegisterProposal(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryProposals,
		registerQueryProposal,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryProposals(r *mux.Router) error {
	r.HandleFunc("/api/proposals/{page}/{size}", func(writer http.ResponseWriter, request *http.Request) {
		var data []document.Proposal
		result := utils.QueryList("proposal", &data, nil, "-submit_block", request)
		var pageInfo types.Page
		json.Unmarshal(result, &pageInfo)
		var proposals []Proposal
		for _, propo := range data {
			mP := Proposal{
				Title:           propo.Title,
				ProposalId:      propo.ProposalId,
				Type:            propo.Type,
				Description:     propo.Description,
				Status:          propo.Status,
				SubmitTime:      propo.SubmitTime,
				DepositEndTime:  propo.DepositEndTime,
				VotingStartTime: propo.VotingStartTime,
				VotingEndTime:   propo.VotingEndTime,
				TotalDeposit:    propo.TotalDeposit,
			}
			proposals = append(proposals, mP)
		}
		pageInfo.Data = proposals
		res, _ := json.Marshal(pageInfo)
		writer.Write(res)
	}).Methods("GET")
	return nil
}

func registerQueryProposal(r *mux.Router) error {
	r.HandleFunc("/api/proposal/{pid}", func(writer http.ResponseWriter, request *http.Request) {
		var data document.Proposal
		db := utils.GetDatabase()
		defer db.Session.Close()
		propoStore := db.C("proposal")
		txStore := db.C("tx_common")
		args := mux.Vars(request)
		pid, _ := strconv.Atoi(args["pid"])

		var info ProposalInfo
		if err := propoStore.Find(bson.M{"proposal_id": pid}).One(&data); err != nil {
			res, _ := json.Marshal(info)
			writer.Write(res)
			return
		}

		proposal := Proposal{
			Title:           data.Title,
			ProposalId:      data.ProposalId,
			Type:            data.Type,
			Description:     data.Description,
			Status:          data.Status,
			SubmitTime:      data.SubmitTime,
			DepositEndTime:  data.DepositEndTime,
			VotingStartTime: data.VotingStartTime,
			VotingEndTime:   data.VotingEndTime,
			TotalDeposit:    data.TotalDeposit,
		}

		var tx document.CommonTx
		if err := txStore.Find(bson.M{"type": types.TypeSubmitProposal, "proposal_id": pid}).One(&tx); err == nil {
			proposal.Proposer = tx.From
			proposal.TxHash = tx.TxHash
		}

		info.Proposal = proposal

		var votes []Vote
		var result VoteResult
		for _, v := range data.Votes {
			vote := Vote{
				Voter:  v.Voter,
				Option: v.Option,
				Time:   v.Time,
			}
			votes = append(votes, vote)

			switch v.Option {
			case "Yes":
				result.Yes++
			case "Abstain":
				result.Abstain++
			case "No":
				result.No++
			case "NoWithVeto":
				result.NoWithVeto++
			}
		}
		info.Votes = votes
		info.Result = result
		res, _ := json.Marshal(info)
		writer.Write(res)

	}).Methods("GET")
	return nil
}

type Proposal struct {
	Title            string      `json:"title"`
	ProposalId       uint64      `json:"proposal_id"`
	Type             string      `json:"type"`
	Description      string      `json:"description"`
	Status           string      `json:"status"`
	SubmitTime       time.Time   `json:"submit_time"`
	DepositEndTime   time.Time   `json:"deposit_end_time"`
	VotingStartTime  time.Time   `json:"voting_start_time"`
	VotingEndTime    time.Time   `json:"voting_end_time"`
	TotalDeposit     store.Coins `json:"total_deposit"`
	VotingStartBlock int64       `json:"voting_start_block"`
	Proposer         string      `json:"proposer"`
	TxHash           string      `json:"tx_hash"`
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
