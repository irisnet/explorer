package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func RegisterTextSearch(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryText,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryText(r *mux.Router) error {
	r.HandleFunc("/api/search/{text}", func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		text := args["text"]

		db := utils.GetDatabase()
		defer db.Session.Close()

		var result []ResultVo
		i, isUint := utils.ParseUint(text)
		if isUint {
			//查询block信息
			var block document.Block
			blockStore := db.C("block")
			err := blockStore.Find(bson.M{"height": i}).One(&block)
			if err == nil {
				vo := ResultVo{
					Type: "block",
					Data: SimpleBlock{
						Height:    block.Height,
						Timestamp: block.Time,
						Hash:      block.Hash,
					},
				}
				result = append(result, vo)
			}

			//查询proposal信息
			var proposal document.Proposal
			proposalStore := db.C("proposal")
			err = proposalStore.Find(bson.M{"proposal_id": i}).One(&proposal)
			if err == nil {
				vo := ResultVo{
					Type: "proposal",
					Data: SimpleProposal{
						ProposalId: proposal.ProposalId,
						Title:      proposal.Title,
						Type:       proposal.Type,
						Status:     proposal.Status,
						SubmitTime: proposal.SubmitTime,
					},
				}
				result = append(result, vo)
			}
		}

		res, _ := json.Marshal(result)
		writer.Write(res)
		return
	}).Methods("GET")
	return nil
}

type ResultVo struct {
	Type string      `json:"Type,omitempty"`
	Data interface{} `json:"Data,omitempty"`
}

type SimpleBlock struct {
	Height    int64
	Timestamp time.Time
	Hash      string
}

type SimpleProposal struct {
	ProposalId uint64
	Title      string
	Type       string
	Status     string
	SubmitTime time.Time
}
