package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func registerQueryTx(r *mux.Router) error {
	r.HandleFunc("/api/tx/{hash}", queryTx).Methods("GET")
	return nil
}

// queryTx is to query transaction by hash
func queryTx(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	hash := args["hash"]

	c := utils.GetDatabase().C("tx_common")
	defer c.Database.Session.Close()

	var result document.CommonTx
	query := bson.M{}
	query["tx_hash"] = hash
	err := c.Find(query).Sort("-time").One(&result)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("not found"))
	}

	tx := buildTx(result)

	switch tx.(type) {
	case types.GovTx:
		govTx := tx.(types.GovTx)
		resultByte, _ := json.Marshal(govTx)
		w.Write(resultByte)
		return
	case types.StakeTx:
		stakeTx := tx.(types.StakeTx)
		if stakeTx.Type == types.TypeBeginRedelegation || stakeTx.Type == types.TypeCompleteRedelegation {
			c := utils.GetDatabase().C("tx_msg")
			var res document.TxMsg
			err := c.Find(bson.M{"hash": stakeTx.Hash}).One(&res)
			if err != nil {
				break
			}
			var msg MsgBeginRedelegate
			if err = json.Unmarshal([]byte(res.Content), &msg); err == nil {
				stakeTx.From = msg.DelegatorAddr
				stakeTx.To = msg.ValidatorDstAddr
				stakeTx.Source = msg.ValidatorSrcAddr
			}
		}
		resultByte, _ := json.Marshal(stakeTx)
		w.Write(resultByte)
		return
	}
	resultByte, _ := json.Marshal(tx)
	w.Write(resultByte)
}

type MsgSubmitProposal struct {
	Title          string      `json:"title"`          //  Title of the proposal
	Description    string      `json:"description"`    //  Description of the proposal
	ProposalType   string      `json:"proposalType"`   //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Proposer       string      `json:"proposer"`       //  Address of the proposer
	InitialDeposit store.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Params         []Param     `json:"params"`
}

type Param struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
	Op    string `bson:"op"`
}

type MsgDeposit struct {
	ProposalID int64       `json:"proposal_id"` // ID of the proposal
	Depositer  string      `json:"depositer"`   // Address of the depositer
	Amount     store.Coins `json:"amount"`      // Coins to add to the proposal's deposit
}

type MsgVote struct {
	ProposalID int64  `json:"proposal_id"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
}

type MsgBeginRedelegate struct {
	DelegatorAddr    string  `json:"delegator_addr"`
	ValidatorSrcAddr string  `json:"validator_src_addr"`
	ValidatorDstAddr string  `json:"validator_dst_addr"`
	SharesAmount     float64 `json:"shares_amount"`
}
