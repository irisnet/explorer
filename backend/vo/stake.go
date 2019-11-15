package vo

import (
	"fmt"
	"time"
)

type ValidatorForDetail struct {
	TotalPower              int64       `json:"total_power"`
	SelfPower               int64       `json:"self_power"`
	Status                  string      `json:"status"`
	BondedTokens            string      `json:"bonded_tokens"`
	SelfBonded              string      `json:"self_bonded"`
	BondedStake             string      `json:"bonded_stake"`
	DelegatorShares         string      `json:"delegator_shares"`
	DelegatorCount          int         `json:"delegator_count"`
	CommissionRate          string      `json:"commission_rate"`
	CommissionUpdate        string      `json:"commission_update"`
	CommissionMaxRate       string      `json:"commission_max_rate"`
	CommissionMaxChangeRate string      `json:"commision_max_change_rate"`
	BondHeight              string      `json:"bond_height"`
	UnbondingHeight         string      `json:"unbond_height"`
	JailedUntil             string      `json:"jailed_until"`
	MissedBlocksCount       string      `json:"missed_blocks_count"`
	OperatorAddr            string      `json:"operator_addr"`
	OwnerAddr               string      `json:"owner_addr"`
	ConsensusAddr           string      `json:"consensus_addr"`
	Description             Description `json:"description"`
	Icons                   string      `json:"icons"`
	Uptime                  float32     `json:"uptime"`
	StatsBlocksWindow       string      `json:"stats_blocks_window"`
}

type WithdrawAddr struct {
	Address string `json:"address"`
}

type RedelegationPage struct {
	Total int            `json:"total"`
	Items []Redelegation `json:"items"`
}

func (re RedelegationPage) String() string {
	return fmt.Sprintf(`
Total  : %v
Items  : %v
		`, re.Total, re.Items)
}

type Redelegation struct {
	Address   string `json:"address"`
	Amount    string `json:"amount"`
	To        string `json:"to"`
	ToMoniker string `json:"to_moniker"`
	Block     string `json:"block"`
}

type DelegationsPage struct {
	Total int          `json:"total"`
	Items []Delegation `json:"items"`
}

func (d DelegationsPage) String() string {
	return fmt.Sprintf(`
total: %v
items:
%v
		`, d.Total, d.Items)
}

type UnbondingDelegationsPage struct {
	Total int                    `json:"total"`
	Items []UnbondingDelegations `json:"items"`
}

func (un UnbondingDelegationsPage) String() string {
	return fmt.Sprintf(`
total: %v
items:
%v
		`, un.Total, un.Items)
}

type UnbondingDelegations struct {
	Address string `json:"address"`
	Moniker string `json:"moniker"`
	Amount  string `json:"amount"`
	Block   int64  `json:"block,string"`
	Until   string `json:"until"`
}

func (un UnbondingDelegations) String() string {
	return fmt.Sprintf(`
		Address :%v
		Amount  :%v
		Block   :%v
		Until   :%v
		`, un.Address, un.Amount, un.Block, un.Until)
}

type Delegation struct {
	Address     string  `json:"address"`
	Moniker     string  `json:"moniker"`
	Amount      string  `json:"amount"`
	SelfShares  string  `json:"self_shares"`
	TotalShares float64 `json:"total_shares"`
	Block       int64   `json:"block,string"`
}

func (d Delegation) String() string {
	return fmt.Sprintf(`
	 Address     :%v
 	 Amount      :%v
 	 Shares      :%v
 	 Block       :%v
	 TotalShares :%v
		`, d.Address, d.Amount, d.SelfShares, d.Block, d.Block)
}

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
	Icons       string      `json:"icons"`
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
	PowerAll int64 `json:"power_all"`
	Validator      `json:"validator"`
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
