package document

import (
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleCandidate = "stake_role_candidate"

	Candidate_Field_Address         = "address"
	Candidate_Field_PubKey          = "pub_key"
	Candidate_Field_PubKeyAddr      = "pub_key_addr"
	Candidate_Field_Jailed          = "jailed"
	Candidate_Field_Tokens          = "tokens"
	Candidate_Field_OriginalTokens  = "original_tokens"
	Candidate_Field_DelegatorShares = "delegator_shares"
	Candidate_Field_VotingPower     = "voting_power"
	Candidate_Field_Description     = "description"
	Candidate_Field_BondHeight      = "bond_height"
	Candidate_Field_Status          = "status"
)

type (
	Candidate struct {
		Address         string               `bson:"address"` // owner, identity key
		PubKey          string               `bson:"pub_key"`
		PubKeyAddr      string               `bson:"pub_key_addr"`
		Jailed          bool                 `bson:"jailed"` // has the validator been revoked from bonded status
		Tokens          float64              `bson:"tokens"`
		OriginalTokens  string               `bson:"original_tokens"`
		DelegatorShares float64              `bson:"delegator_shares"`
		VotingPower     float64              `bson:"voting_power"` // Voting power if pubKey is a considered a validator
		Description     utils.ValDescription `bson:"description"`  // Description terms for the candidate
		BondHeight      int64                `bson:"bond_height"`
		Status          string               `bson:"status"`
		Rank            int                  `bson:"rank,omitempty"`
	}
)

func (d Candidate) Name() string {
	return CollectionNmStakeRoleCandidate
}

func (d Candidate) PkKvPair() map[string]interface{} {
	return bson.M{Candidate_Field_Address: d.Address}
}
