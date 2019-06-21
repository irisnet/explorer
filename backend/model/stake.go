package model

import (
	"fmt"
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

func (vStatus ValStatus) String() string {
	return fmt.Sprintf(`
		Uptime         :%v
		PrecommitCount :%v
		`, vStatus.Uptime, vStatus.PrecommitCount)
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
	Address     string      `json:"address"` // operator
	Owner       string      `json:"owner"`   // owner, identity key
	PubKey      string      `json:"pub_key"`
	Jailed      bool        `json:"jailed"`                // has the validator been revoked from bonded status
	VotingPower int64       `json:"voting_power"`          // Voting power if pubKey is a considered a validator
	Description Description `json:"description,omitempty"` // Description terms for the candidate
	BondHeight  int64       `json:"bond_height"`
	Status      string      `json:"status"`
	Rate        string      `json:"rate"`
	ValStatus
}

func (v Validator) String() string {
	return fmt.Sprintf(
		`
		Address        :%v
		Owner          :%v
		PubKey         :%v
		Jailed         :%v
		VotingPower    :%v
		Description    :%v
		BondHeight     :%v
		Status         :%v
		Rate           :%v
		`, v.Address, v.Owner, v.PubKey, v.Jailed, v.VotingPower, v.Description, v.BondHeight, v.Status, v.Rate)
}

type Description struct {
	Moniker  string `json:"moniker,omitempty"`
	Identity string `json:"identity,omitempty"`
	Website  string `json:"website,omitempty"`
	Details  string `json:"details,omitempty"`
}

func (d Description) String() string {
	return fmt.Sprintf(`
		Moniker  :%v
		Identity :%v
		Website  :%v
		Details  :%v
		`, d.Moniker, d.Identity, d.Website, d.Details)
}

type CandidatesInfoVo struct {
	PowerAll  int64 `json:"power_all"`
	Validator `json:"validator"`
}

func (c CandidatesInfoVo) String() string {

	return fmt.Sprintf(`
		PowerAll  : %v
    Validator : %v
		`, c.PowerAll, c.Validator)
}

type ChainStatusVo struct {
	ValidatorsCount int
	TxCount         int
	VotingPower     int64
	Tps             float64
}

type UptimeChangeVo struct {
	Address string  `json:"address"`
	Time    string  `json:"time"`
	Uptime  float64 `json:"uptime"`
}

func (uptime UptimeChangeVo) String() string {
	return fmt.Sprintf(`
		Address :%v
		Time    :%v
		Uptime  :%v
		`, uptime.Address, uptime.Time, uptime.Uptime)
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
	Height  int64     `json:"height"`
	Address string    `json:"address"`
	Power   int64     `json:"power"`
	Time    time.Time `json:"time"`
	Change  string    `json:"change"`
}
