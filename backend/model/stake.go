package model

import (
	"time"
)

type ValDetailVo struct {
	Count      int         `json:"count"`
	PowerAll   int64       `json:"power_all"`
	Validators []Validator `json:"validators"`
}

type ValStatus struct {
	Uptime         int     `json:"up_time"`
	PrecommitCount float64 `json:"precommit_cnt"`
}

type ValProfile struct {
	Description
	PubKey         string `json:"pub_key"`
	Owner          string `json:"owner"`
	BondHeight     int64  `json:"bond_height"`
	VotingPower    int64  `json:"voting_power"`
	CommitBlockNum int64  `json:"commit_block_num"`
	UpTime         int    `json:"up_time"`
}

type Validator struct {
	Address        string      `json:"address"` // owner, identity key
	PubKey         string      `json:"pub_key"`
	Jailed         bool        `json:"jailed"`                // has the validator been revoked from bonded status
	VotingPower    int64       `json:"voting_power"`          // Voting power if pubKey is a considered a validator
	Description    Description `json:"description,omitempty"` // Description terms for the candidate
	BondHeight     int64       `json:"bond_height"`
	Status         string      `json:"status"`
	OriginalTokens string      `json:"original_tokens"`
	Rate           string      `json:"rate"`
	ValStatus
}

type Description struct {
	Moniker  string `json:"moniker,omitempty"`
	Identity string `json:"identity,omitempty"`
	Website  string `json:"website,omitempty"`
	Details  string `json:"details,omitempty"`
}

type CandidatesInfoVo struct {
	PowerAll  int64 `json:"power_all"`
	Validator `json:"validator"`
}

type ChainStatusVo struct {
	ValidatorsCount int
	TxCount         int
	VotingPower     int64
	Tps             float64
}

type UptimeChangeVo struct {
	Address string
	Time    string
	Uptime  float64
}

type ValUpTimeVo struct {
	Time   string `bson:"_id,omitempty"`
	Uptime float64
}

type ValVotingPowerChangeInfo struct {
	Address string `bson:"_id,omitempty"`
	Power   int64
	Time    time.Time
	Change  string
}

type ValVotingPowerChangeVo struct {
	Height  int64
	Address string
	Power   int64
	Time    time.Time
	Change  string
}
