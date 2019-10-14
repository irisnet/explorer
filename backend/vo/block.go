package vo

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/utils"
)

type BlockForList struct {
	Height                  int64     `json:"height"`
	ProposerMoniker         string    `json:"proposer_moniker"`
	ProposerAsValidatorAddr string    `json:"proposer_as_validator_addr"`
	TxnNum                  int64     `json:"txs_num"`
	PrecommitValidatorNum   int       `json:"precommit_validator_num"`
	ValidatorNumForHeight   int       `json:"validator_num_for_height"`
	PrecommitVotingPower    int64     `json:"precommit_voting_power"`
	VotingPowerForHeight    int64     `json:"voting_power_for_height"`
	Timestamp               time.Time `json:"timestamp"`
}

type BlockForListRespond []BlockForList

func (b BlockForList) String() string {

	return fmt.Sprintf(`
		Height                  :%v
		ProposerMoniker         :%v
		ProposerAsValidatorAddr :%v
		TxnNum                  :%v
		PrecommitValidatorNum   :%v
		ValidatorNumForHeight   :%v
		PrecommitVotingPower    :%v
		VotingPowerForHeight    :%v
		Timestamp               :%v
		`, b.Height, b.ProposerMoniker, b.ProposerAsValidatorAddr, b.TxnNum, b.PrecommitValidatorNum, b.ValidatorNumForHeight, b.PrecommitVotingPower, b.VotingPowerForHeight, b.Timestamp)
}

type BlockInfoVo struct {
	Height        int64     `json:"height"`
	Hash          string    `json:"hash,omitempty"`
	Time          time.Time `json:"time"`
	NumTxs        int64     `json:"num_txs"`
	Validators    []ValInfo `json:"validators,omitempty"`
	LastCommit    []string  `json:"last_commit,omitempty"`
	TotalTxs      int64     `json:"total_txs,omitempty"`
	LastBlockHash string    `json:"last_block_hash,omitempty"`
}

type BlockInfoVoRespond []BlockInfoVo

func (b BlockInfoVo) String() string {
	return fmt.Sprintf(`
		Height        :%v
		Hash          :%v
		Time          :%v
		NumTxs        :%v
		Validators    :%v
		LastCommit    :%v
		TotalTxs      :%v
		LastBlockHash :%v
		`, b.Height, b.Hash, b.Time, b.NumTxs, b.Validators, b.LastCommit, b.TotalTxs, b.LastBlockHash)
}

type ValInfo struct {
	Address     string `json:"address"`
	VotingPower int64  `json:"voting_power"`
}

func (v ValInfo) String() string {
	return fmt.Sprintf("address: %v   voting power: %v \n", v.Address, v.VotingPower)
}

type BlockValidator struct {
	Moniker          string `json:"moniker"`
	OperatorAddress  string `json:"operator_address"`
	Consensus        string `json:"consensus"`
	ProposerPriority string `json:"proposer_priority"`
	VotingPower      string `json:"voting_power"`
	IsProposer       bool   `json:"is_proposer"`
}

func (v BlockValidator) String() string {
	return fmt.Sprintf(`
		Moniker          :%v
		OperatorAddress  :%v
		Consensus        :%v
		ProposerPriority :%v
		VotingPower      :%v
		IsProposer       :%v
		`, v.Moniker, v.OperatorAddress, v.Consensus, v.ProposerPriority, v.VotingPower, v.IsProposer)
}

type ValidatorSet struct {
	Total int              `json:"total"`
	Items []BlockValidator `json:"items"`
}

type Block struct {
	BlockInfo    BlockInfo    `json:"block_info"`
	ValidatorSet ValidatorSet `json:"validator_set"`
	TxPage       TxPage       `json:"tx_page"`
	ProposalPage ProposalPage `json:"proposals_page"`
}

type BlockInfo struct {
	BlockHeight           string      `json:"block_height"`
	Timestamp             time.Time   `json:"timestamp"`
	BlockHash             string      `json:"block_hash"`
	Transactions          string      `json:"transactions"`
	PropopserMoniker      string      `json:"propopser_moniker"`
	PropoperAddr          string      `json:"propopser_addr"`
	PrecommitValidatorNum interface{} `json:"precommit_validator_num"`
	TotalValidatorNum     int         `json:"total_validator_num"`
	PrecommitVotingPower  interface{} `json:"precommit_voting_power"`
	TotalVotingPower      int         `json:"total_voting_power"`
	LatestHeight          string      `json:"latest_height"`
	MintCoin              utils.Coin  `json:"mint_coin"`
}

func (info BlockInfo) String() string {
	return fmt.Sprintf(`
		BlockHeight           :%v
		Timestamp             :%v
		BlockHash             :%v
		Transactions          :%v
		PropopserMoniker      :%v
		PropoperAddr          :%v
		PrecommitValidatorNum :%v
		TotalValidatorNum     :%v
		PrecommitVotingPower  :%v
		TotalVotingPower      :%v
		LatestHeight          :%v
		MintCoin              :%v
		`, info.BlockHeight, info.Timestamp, info.BlockHash, info.Transactions, info.PropopserMoniker, info.PropoperAddr, info.PrecommitValidatorNum, info.TotalValidatorNum, info.PrecommitVotingPower,
		info.TotalVotingPower, info.LatestHeight, info.MintCoin)
}
