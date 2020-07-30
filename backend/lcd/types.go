package lcd

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/utils"
)

const (

	UrlGovParam                                  = "%s/params?module=%s"
	UrlDistributionRewardsByValidatorAcc         = "%s/distribution/%s/rewards"
	UrlValidatorsSigningInfoByConsensuPublicKey  = "%s/slashing/validators/%s/signing-info"
	UrlTokenStatsSupply                          = "https://rpc.irisnet.org/token-stats/supply"
	UrlTokenStatsCirculation                     = "https://rpc.irisnet.org/token-stats/circulation"
	UrlLookupIconsByKeySuffix                    = "https://keybase.io/_/api/1.0/user/lookup.json?fields=pictures&key_suffix=%s"
	CommunityTaxAddr                             = "iaa18rtw90hxz4jsgydcusakz6q245jh59kfma3e5h"
)

type AccountVo struct {
	Address       string   `json:"address"`
	Coins         []string `json:"coins"`
	PublicKey     string   `json:"public_key"`
	AccountNumber uint64   `json:"account_number"`
	Sequence      string   `json:"sequence"`
}

type ValidatorVo struct {
	OperatorAddress string      `json:"operator_address"`
	ConsensusPubkey string      `json:"consensus_pubkey"`
	Jailed          bool        `json:"jailed"`
	Status          int         `json:"status"`
	Tokens          string      `json:"tokens"`
	DelegatorShares string      `json:"delegator_shares"`
	Description     Description `json:"description"`
	BondHeight      string      `json:"bond_height"`
	UnbondingHeight string      `json:"unbonding_height"`
	UnbondingTime   time.Time   `json:"unbonding_time"`
	Commission      Commission  `json:"commission"`
	Uptime          float32     `json:"uptime"`
	SelfBond        string      `json:"self_bond"`
	DelegatorNum    int         `json:"delegator_num"`
	ProposerAddr    string      `json:"proposer_addr"`
	VotingRate      float32     `json:"voting_rate"`
	Icons           string      `json:"icons"`
}

type ValidatorsVoRespond []ValidatorVo

func (v ValidatorVo) String() string {
	return fmt.Sprintf(`
		OperatorAddress :%v
		ConsensusPubkey :%v
		Jailed          :%v
		Status          :%v
		Tokens          :%v
		DelegatorShares :%v
		Description     :%v
		BondHeight      :%v
		UnbondingHeight :%v
		UnbondingTime   :%v
		Commission      :%v
		Uptime          :%v
		SelfBond        :%v
		DelegatorNum    :%v
		ProposerAddr    :%v
		VotingRate      :%v
		Icons           :%v
		`, v.OperatorAddress, v.ConsensusPubkey, v.Jailed, v.Status, v.Tokens, v.DelegatorShares, v.Description, v.BondHeight, v.UnbondingHeight, v.UnbondingTime,
		v.Commission, v.Uptime, v.SelfBond, v.DelegatorNum, v.ProposerAddr, v.VotingRate, v.Icons)
}

type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
}
type Commission struct {
	Rate          string    `json:"rate"`
	MaxRate       string    `json:"max_rate"`
	MaxChangeRate string    `json:"max_change_rate"`
	UpdateTime    time.Time `json:"update_time"`
}

type NodeInfoVo struct {
	//ProtocolVersion struct {
	//	P2P   string `json:"p2p"`
	//	Block string `json:"block"`
	//	App   string `json:"app"`
	//} `json:"protocol_version"`
	//ID         string `json:"id"`
	//ListenAddr string `json:"listen_addr"`
	Network string `json:"network"`
	Version string `json:"version"`
	//Channels   string `json:"channels"`
	Moniker string `json:"moniker"`
	//Other      struct {
	//	TxIndex    string `json:"tx_index"`
	//	RPCAddress string `json:"rpc_address"`
	//} `json:"other"`
}



type BlockVo struct {
	BlockMeta BlockMeta `json:"block_meta"`
	Block     BlockI    `json:"block"`
}

type (
	BlockMeta struct {
		BlockID BlockID     `json:"block_id"`
		Header  BlockHeader `json:"header"`
	}
	BlockID struct {
		Hash string `json:"hash"`
		//Parts struct {
		//	Total string `json:"total"`
		//	Hash  string `json:"hash"`
		//} `json:"parts"`
	}
	BlockHeader struct {
		//	Version struct {
		//		Block string `json:"block"`
		//		App   string `json:"app"`
		//	} `json:"version"`
		//	ChainID     string    `json:"chain_id"`
		Height int64     `json:"height"`
		Time   time.Time `json:"time"`
		NumTxs string    `json:"num_txs"`
		//TotalTxs string    `json:"total_txs"`
		//	LastBlockID struct {
		//		Hash  string `json:"hash"`
		//		Parts struct {
		//			Total string `json:"total"`
		//			Hash  string `json:"hash"`
		//		} `json:"parts"`
		//	} `json:"last_block_id"`
		//	LastCommitHash     string `json:"last_commit_hash"`
		//	DataHash           string `json:"data_hash"`
		//	ValidatorsHash     string `json:"validators_hash"`
		//	NextValidatorsHash string `json:"next_validators_hash"`
		//	ConsensusHash      string `json:"consensus_hash"`
		//	AppHash            string `json:"app_hash"`
		//	LastResultsHash    string `json:"last_results_hash"`
		//	EvidenceHash       string `json:"evidence_hash"`
		ProposerAddress string `json:"proposer_address"`
	}
	BlockI struct {
		Header     Header     `json:"header"`
		LastCommit LastCommit `json:"last_commit"`
	}
	Header struct {
		//	Version struct {
		//		Block string `json:"block"`
		//		App   string `json:"app"`
		//	} `json:"version"`
		//	ChainID     string    `json:"chain_id"`
		Height int64 `json:"height"`
		//	Time        time.Time `json:"time"`
		//	NumTxs      string    `json:"num_txs"`
		//	TotalTxs    string    `json:"total_txs"`
		//	LastBlockID struct {
		//		Hash  string `json:"hash"`
		//		Parts struct {
		//			Total string `json:"total"`
		//			Hash  string `json:"hash"`
		//		} `json:"parts"`
		//	} `json:"last_block_id"`
		//	LastCommitHash     string `json:"last_commit_hash"`
		//	DataHash           string `json:"data_hash"`
		//	ValidatorsHash     string `json:"validators_hash"`
		//	NextValidatorsHash string `json:"next_validators_hash"`
		//	ConsensusHash      string `json:"consensus_hash"`
		//	AppHash            string `json:"app_hash"`
		//	LastResultsHash    string `json:"last_results_hash"`
		//	EvidenceHash       string `json:"evidence_hash"`
		//	ProposerAddress    string `json:"proposer_address"`
	}
	LastCommit struct {
		//BlockID struct {
		//	Hash  string `json:"hash"`
		//	Parts struct {
		//		Total string `json:"total"`
		//		Hash  string `json:"hash"`
		//	} `json:"parts"`
		//} `json:"block_id"`
		Precommits []Precommit `json:"precommits"`
	}
	Precommit struct {
		//	Type    int    `json:"type"`
		Height string `json:"height"`
		//	Round   string `json:"round"`
		//	BlockID struct {
		//		Hash  string `json:"hash"`
		//		Parts struct {
		//			Total string `json:"total"`
		//			Hash  string `json:"hash"`
		//		} `json:"parts"`
		//	} `json:"block_id"`
		//Timestamp        time.Time `json:"timestamp"`
		ValidatorAddress string `json:"validator_address"`
		ValidatorIndex   string `json:"validator_index"`
		//Signature        string    `json:"signature"`
	}
)

type ValidatorSetVo struct {
	BlockHeight string             `json:"block_height"`
	Validators  []StakeValidatorVo `json:"validators"`
}

type StakeValidatorVo struct {
	Address          string `json:"address"`
	PubKey           string `json:"pub_key"`
	ProposerPriority int64  `json:"proposer_priority"`
	VotingPower      int64  `json:"voting_power"`
}
type StakePoolVo struct {
	LooseTokens  string `json:"loose_tokens"`
	BondedTokens string `json:"bonded_tokens"`
	//TotalSupply  string `json:"total_supply"`
	//BondedRatio  string `json:"bonded_ratio"`
}

type DelegationVo struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
	Shares        string `json:"shares"`
	Height        int64  `json:"height,string"`
}

type DelegationFromVal struct {
	Tokens          string `json:"tokens"`
	DelegatorShares string `json:"delegator_shares"`
	OperatorAddress string `json:"operator_address"`
	BondHeight      int64  `json:"bond_height,string"`
}

type ValidatorDelegations []DelegationVo

func (sort ValidatorDelegations) Len() int {
	return len(sort)
}
func (sort ValidatorDelegations) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}
func (sort ValidatorDelegations) Less(i, j int) bool {
	return sort[i].Height > sort[j].Height
}

type DistributionRewards struct {
	Total       utils.CoinsAsStr         `json:"total"`
	Delegations []RewardsFromDelegations `json:"delegations"`
	Commission  utils.CoinsAsStr         `json:"commission"`
}

type RewardsFromDelegations struct {
	Validator string           `json:"validator"`
	Reward    utils.CoinsAsStr `json:"reward"`
}

type ValidatorSigningInfo struct {
	StartHeight       string `json:"start_height"`
	IndexOffset       string `json:"index_offset"`
	JailedUntil       string `json:"jailed_until"`
	MissedBlocksCount string `json:"missed_blocks_counter"`
}

type ReDelegation struct {
	DelegatorAddr    string `json:"delegator_addr"`
	ValidatorSrcAddr string `json:"validator_src_addr"`
	ValidatorDstAddr string `json:"validator_dst_addr"`
	CreationHeight   string `json:"creation_height"`
	MinTime          int64  `json:"min_time"`
	InitialBalance   string `json:"initial_balance"`
	Balance          string `json:"balance"`
	SharesSrc        string `json:"shares_src"`
	SharesDst        string `json:"shares_dst"`
}

type UnbondingDelegation struct {
	DelegatorAddr  string `json:"delegator_addr"`
	ValidatorAddr  string `json:"validator_addr"`
	InitialBalance string `json:"initial_balance"`
	Balance        string `json:"balance"`
	CreationHeight int64  `json:"creation_height,string"`
	MinTime        string `json:"min_time"`
}

func (un UnbondingDelegation) String() string {
	return fmt.Sprintf(`
		DelegatorAddr  :%v
		ValidatorAddr  :%v
		InitialBalance :%v
		Balance        :%v
		CreationHeight :%v
		MinTime        :%v
		`, un.DelegatorAddr, un.ValidatorAddr, un.InitialBalance, un.Balance, un.CreationHeight, un.MinTime)

}

func (d DelegationVo) String() string {
	return fmt.Sprintf(`
		DelegatorAddr :%v
		ValidatorAddr :%v
		Shares        :%v
		Height        :%v
		`, d.DelegatorAddr, d.ValidatorAddr, d.Shares, d.Height)
}

type SignInfoVo struct {
	StartHeight         int64     `json:"start_height"`
	IndexOffset         int64     `json:"index_offset"`
	JailedUntil         time.Time `json:"jailed_until"`
	MissedBlocksCounter int64     `json:"missed_blocks_counter"`
}

type (
	BlockResultVo struct {
		//Height string `json:"height"`
		Results Results `json:"results"`
		//TxsResults []struct {
		//	Code      int          `json:"code"`
		//	Data      interface{}  `json:"data"`
		//	Log       string       `json:"log"`
		//	Info      string       `json:"info"`
		//	GasWanted int64        `json:"gas_wanted,string"`
		//	GasUsed   int64        `json:"gas_used,string"`
		//	Events    []BlockEvent `json:"events"`
		//} `json:"txs_results"`
	}
	Results struct {
		//DeliverTx []struct {
		//	Code      int         `json:"code"`
		//	Data      interface{} `json:"data"`
		//	Log       string      `json:"log"`
		//	Info      string      `json:"info"`
		//	GasWanted string      `json:"gas_wanted"`
		//	GasUsed   string      `json:"gas_used"`
		//	Tags []struct {
		//		Key   string `json:"key"`
		//		Value string `json:"value"`
		//	} `json:"tags"`
		//} `json:"deliver_tx"`
		//EndBlock struct {
		//	ValidatorUpdates []struct {
		//		PubKey struct {
		//			Type string `json:"type"`
		//			Data string `json:"data"`
		//		} `json:"pub_key"`
		//		Power string `json:"power"`
		//	} `json:"validator_updates"`
		//	ConsensusParamUpdates interface{} `json:"consensus_param_updates"`
		//	Tags []struct {
		//		Key   string `json:"key"`
		//		Value string `json:"value"`
		//	} `json:"tags"`
		//} `json:"end_block"`
		BeginBlock BeginBlockEvents `json:"begin_block"`
	}
	BeginBlockEvents []BlockEvent
	BlockEvent struct {
		Type       string            `json:"type"`
		Attributes map[string]string `json:"attributes"`
	}
)

type BlockCoinFlowVo struct {
	Height   string   `json:"height"`
	CoinFlow []string `json:"coin_flow"`
	Tx       Tx       `json:"tx"`
}

type Tx struct {
	Value struct {
		Msg []struct {
			Type string `json:"type"`
			Value struct {
				DelegatorAddr string `json:"delegator_addr"`
			} `json:"value"`
		} `json:"msg"`
	} `json:"value"`
}

type (
	AssetToken struct {
		BaseToken BaseToken `json:"base_token"`
	}
	BaseToken struct {
		Id string `json:"id"`
		//Family          string `json:"family"`
		//Source          string `json:"source"`
		//Gateway         string `json:"gateway"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Scale  int    `json:"scale"`
		//CanonicalSymbol string `json:"canonical_symbol"`
		MinUnitAlias  string `json:"min_unit_alias"`
		InitialSupply string `json:"initial_supply"`
		MaxSupply     string `json:"max_supply"`
		Mintable      bool   `json:"mintable"`
		Owner         string `json:"owner"`
	}
)

type AssetGateways struct {
	Owner    string `json:"owner"`
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Details  string `json:"details"`
	Website  string `json:"website"`
}
