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
				Title:            propo.Title,
				ProposalId:       propo.ProposalId,
				Type:             propo.Type,
				Description:      propo.Description,
				Status:           propo.Status,
				SubmitBlock:      propo.SubmitBlock,
				SubmitTime:       propo.SubmitTime,
				VotingStartBlock: propo.VotingStartBlock,
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
		c := utils.GetDatabase().C("proposal")
		defer c.Database.Session.Close()
		args := mux.Vars(request)
		pid, _ := strconv.Atoi(args["pid"])
		if err := c.Find(bson.M{"proposal_id": pid}).One(&data); err != nil {
			writer.Write([]byte("error"))
			return
		}
		var info ProposalInfo
		proposal := Proposal{
			Title:            data.Title,
			ProposalId:       data.ProposalId,
			Type:             data.Type,
			Description:      data.Description,
			Status:           data.Status,
			SubmitBlock:      data.SubmitBlock,
			SubmitTime:       data.SubmitTime,
			TotalDeposit:     data.TotalDeposit,
			VotingStartBlock: data.VotingStartBlock,
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
	ProposalId       int64       `json:"proposal_id"`
	Type             string      `json:"type"`
	Description      string      `json:"description"`
	Status           string      `json:"status"`
	SubmitBlock      int64       `json:"submit_block"`
	SubmitTime       time.Time   `json:"submit_time"`
	TotalDeposit     store.Coins `json:"total_deposit"`
	VotingStartBlock int64       `json:"voting_start_block"`
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
