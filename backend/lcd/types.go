package lcd

import (
	"time"
)

const (
	UrlAccount            = "%s/auth/accounts/%s"
	UrlValidator          = "%s/stake/validators/%s"
	UrlValidators         = "%s/stake/validators?page=%d&size=%d"
	UrlDelegationByVal    = "%s/stake/validators/%s/delegations"
	UrlSignInfo           = "%s/slashing/validators/%s/signing_info"
	UrlNodeInfo           = "%s/node_info"
	UrlGenesis            = "%s/genesis"
	UrlWithdrawAddress    = "%s/distribution/%s/withdrawAddress"
	UrlBlockLatest        = "%s/blocks/latest"
	UrlBlock              = "%s/blocks/%d"
	UrlValidatorSet       = "%s/validatorsets/%d"
	UrlValidatorSetLatest = "%s/validatorsets/latest"
	UrlStakePool          = "%s/stake/pool"
	UrlBlocksResult       = "%s/blocks-result/%d"
)

type AccountVo struct {
	Address   string   `json:"address"`
	Coins     []string `json:"coins"`
	PublicKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"public_key"`
	AccountNumber string `json:"account_number"`
	Sequence      string `json:"sequence"`
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
	ProtocolVersion struct {
		P2P   string `json:"p2p"`
		Block string `json:"block"`
		App   string `json:"app"`
	} `json:"protocol_version"`
	ID         string `json:"id"`
	ListenAddr string `json:"listen_addr"`
	Network    string `json:"network"`
	Version    string `json:"version"`
	Channels   string `json:"channels"`
	Moniker    string `json:"moniker"`
	Other      struct {
		TxIndex    string `json:"tx_index"`
		RPCAddress string `json:"rpc_address"`
	} `json:"other"`
}

type GenesisVo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Genesis struct {
			GenesisTime     time.Time `json:"genesis_time"`
			ChainID         string    `json:"chain_id"`
			ConsensusParams struct {
				BlockSize struct {
					MaxBytes string `json:"max_bytes"`
					MaxGas   string `json:"max_gas"`
				} `json:"block_size"`
				Evidence struct {
					MaxAge string `json:"max_age"`
				} `json:"evidence"`
				Validator struct {
					PubKeyTypes []string `json:"pub_key_types"`
				} `json:"validator"`
			} `json:"consensus_params"`
			AppHash  string `json:"app_hash"`
			AppState struct {
				Accounts []struct {
					Address        string   `json:"address"`
					Coins          []string `json:"coins"`
					SequenceNumber string   `json:"sequence_number"`
					AccountNumber  string   `json:"account_number"`
				} `json:"accounts"`
				Auth struct {
					CollectedFee interface{} `json:"collected_fee"`
					Data         struct {
						NativeFeeDenom string `json:"native_fee_denom"`
					} `json:"data"`
					Params struct {
						GasPriceThreshold string `json:"gas_price_threshold"`
					} `json:"params"`
				} `json:"auth"`
				Stake struct {
					Pool struct {
						BondedTokens string `json:"bonded_tokens"`
					} `json:"pool"`
					Params struct {
						UnbondingTime string `json:"unbonding_time"`
						MaxValidators int    `json:"max_validators"`
					} `json:"params"`
					LastTotalPower       string      `json:"last_total_power"`
					LastValidatorPowers  interface{} `json:"last_validator_powers"`
					Validators           interface{} `json:"validators"`
					Bonds                interface{} `json:"bonds"`
					UnbondingDelegations interface{} `json:"unbonding_delegations"`
					Redelegations        interface{} `json:"redelegations"`
					Exported             bool        `json:"exported"`
				} `json:"stake"`
				Mint struct {
					Minter struct {
						LastUpdate        time.Time `json:"last_update"`
						MintDenom         string    `json:"mint_denom"`
						InflationBasement string    `json:"inflation_basement"`
					} `json:"minter"`
					Params struct {
						Inflation string `json:"inflation"`
					} `json:"params"`
				} `json:"mint"`
				Distr struct {
					Params struct {
						CommunityTax        string `json:"community_tax"`
						BaseProposerReward  string `json:"base_proposer_reward"`
						BonusProposerReward string `json:"bonus_proposer_reward"`
					} `json:"params"`
					FeePool struct {
						ValAccum struct {
							UpdateHeight string `json:"update_height"`
							Accum        string `json:"accum"`
						} `json:"val_accum"`
						ValPool       interface{} `json:"val_pool"`
						CommunityPool interface{} `json:"community_pool"`
					} `json:"fee_pool"`
					ValidatorDistInfos     interface{} `json:"validator_dist_infos"`
					DelegatorDistInfos     interface{} `json:"delegator_dist_infos"`
					DelegatorWithdrawInfos interface{} `json:"delegator_withdraw_infos"`
					PreviousProposer       string      `json:"previous_proposer"`
				} `json:"distr"`
				Gov struct {
					Params struct {
						CriticalDepositPeriod string `json:"critical_deposit_period"`
						CriticalMinDeposit    []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"critical_min_deposit"`
						CriticalVotingPeriod   string `json:"critical_voting_period"`
						CriticalMaxNum         string `json:"critical_max_num"`
						CriticalThreshold      string `json:"critical_threshold"`
						CriticalVeto           string `json:"critical_veto"`
						CriticalParticipation  string `json:"critical_participation"`
						CriticalPenalty        string `json:"critical_penalty"`
						ImportantDepositPeriod string `json:"important_deposit_period"`
						ImportantMinDeposit    []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"important_min_deposit"`
						ImportantVotingPeriod  string `json:"important_voting_period"`
						ImportantMaxNum        string `json:"important_max_num"`
						ImportantThreshold     string `json:"important_threshold"`
						ImportantVeto          string `json:"important_veto"`
						ImportantParticipation string `json:"important_participation"`
						ImportantPenalty       string `json:"important_penalty"`
						NormalDepositPeriod    string `json:"normal_deposit_period"`
						NormalMinDeposit       []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"normal_min_deposit"`
						NormalVotingPeriod  string `json:"normal_voting_period"`
						NormalMaxNum        string `json:"normal_max_num"`
						NormalThreshold     string `json:"normal_threshold"`
						NormalVeto          string `json:"normal_veto"`
						NormalParticipation string `json:"normal_participation"`
						NormalPenalty       string `json:"normal_penalty"`
						SystemHaltPeriod    string `json:"system_halt_period"`
					} `json:"params"`
				} `json:"gov"`
				Upgrade struct {
					GenesisVersion struct {
						UpgradeInfo struct {
							ProposalID string `json:"ProposalID"`
							Protocol   struct {
								Version  string `json:"version"`
								Software string `json:"software"`
								Height   string `json:"height"`
							} `json:"Protocol"`
						} `json:"UpgradeInfo"`
						Success bool `json:"Success"`
					} `json:"GenesisVersion"`
					UpgradeParams struct {
						Threshold string `json:"threshold"`
					} `json:"UpgradeParams"`
				} `json:"upgrade"`
				Slashing struct {
					Params struct {
						MaxEvidenceAge           string `json:"max-evidence-age"`
						SignedBlocksWindow       string `json:"signed-blocks-window"`
						MinSignedPerWindow       string `json:"min-signed-per-window"`
						DoubleSignUnbondDuration string `json:"double-sign-unbond-duration"`
						DowntimeUnbondDuration   string `json:"downtime-unbond-duration"`
						SlashFractionDoubleSign  string `json:"slash-fraction-double-sign"`
						SlashFractionDowntime    string `json:"slash-fraction-downtime"`
					} `json:"params"`
					SigningInfos struct {
					} `json:"signing_infos"`
					MissedBlocks struct {
					} `json:"missed_blocks"`
					SlashingPeriods interface{} `json:"slashing_periods"`
				} `json:"slashing"`
				Service struct {
					Params struct {
						MaxRequestTimeout    string `json:"max_request_timeout"`
						MinDepositMultiple   string `json:"min_deposit_multiple"`
						ServiceFeeTax        string `json:"service_fee_tax"`
						SlashFraction        string `json:"slash_fraction"`
						ComplaintRetrospect  string `json:"complaint_retrospect"`
						ArbitrationTimeLimit string `json:"arbitration_time_limit"`
					} `json:"params"`
				} `json:"service"`
				Guardian struct {
					Profilers []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Address     string `json:"address"`
						AddedBy     string `json:"added_by"`
					} `json:"profilers"`
					Trustees []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Address     string `json:"address"`
						AddedBy     string `json:"added_by"`
					} `json:"trustees"`
				} `json:"guardian"`
				Gentxs []struct {
					Type  string `json:"type"`
					Value struct {
						Msg []struct {
							Type  string `json:"type"`
							Value struct {
								Description struct {
									Moniker  string `json:"moniker"`
									Identity string `json:"identity"`
									Website  string `json:"website"`
									Details  string `json:"details"`
								} `json:"Description"`
								Commission struct {
									Rate          string `json:"rate"`
									MaxRate       string `json:"max_rate"`
									MaxChangeRate string `json:"max_change_rate"`
								} `json:"Commission"`
								DelegatorAddress string `json:"delegator_address"`
								ValidatorAddress string `json:"validator_address"`
								Pubkey           struct {
									Type  string `json:"type"`
									Value string `json:"value"`
								} `json:"pubkey"`
								Delegation struct {
									Denom  string `json:"denom"`
									Amount string `json:"amount"`
								} `json:"delegation"`
							} `json:"value"`
						} `json:"msg"`
						Fee struct {
							Amount interface{} `json:"amount"`
							Gas    string      `json:"gas"`
						} `json:"fee"`
						Signatures []struct {
							PubKey struct {
								Type  string `json:"type"`
								Value string `json:"value"`
							} `json:"pub_key"`
							Signature     string `json:"signature"`
							AccountNumber string `json:"account_number"`
							Sequence      string `json:"sequence"`
						} `json:"signatures"`
						Memo string `json:"memo"`
					} `json:"value"`
				} `json:"gentxs"`
			} `json:"app_state"`
		} `json:"genesis"`
	} `json:"result"`
}

type BlockVo struct {
	BlockMeta struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total string `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			ChainID     string    `json:"chain_id"`
			Height      string    `json:"height"`
			Time        time.Time `json:"time"`
			NumTxs      string    `json:"num_txs"`
			TotalTxs    string    `json:"total_txs"`
			LastBlockID struct {
				Hash  string `json:"hash"`
				Parts struct {
					Total string `json:"total"`
					Hash  string `json:"hash"`
				} `json:"parts"`
			} `json:"last_block_id"`
			LastCommitHash     string `json:"last_commit_hash"`
			DataHash           string `json:"data_hash"`
			ValidatorsHash     string `json:"validators_hash"`
			NextValidatorsHash string `json:"next_validators_hash"`
			ConsensusHash      string `json:"consensus_hash"`
			AppHash            string `json:"app_hash"`
			LastResultsHash    string `json:"last_results_hash"`
			EvidenceHash       string `json:"evidence_hash"`
			ProposerAddress    string `json:"proposer_address"`
		} `json:"header"`
	} `json:"block_meta"`
	Block struct {
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			ChainID     string    `json:"chain_id"`
			Height      string    `json:"height"`
			Time        time.Time `json:"time"`
			NumTxs      string    `json:"num_txs"`
			TotalTxs    string    `json:"total_txs"`
			LastBlockID struct {
				Hash  string `json:"hash"`
				Parts struct {
					Total string `json:"total"`
					Hash  string `json:"hash"`
				} `json:"parts"`
			} `json:"last_block_id"`
			LastCommitHash     string `json:"last_commit_hash"`
			DataHash           string `json:"data_hash"`
			ValidatorsHash     string `json:"validators_hash"`
			NextValidatorsHash string `json:"next_validators_hash"`
			ConsensusHash      string `json:"consensus_hash"`
			AppHash            string `json:"app_hash"`
			LastResultsHash    string `json:"last_results_hash"`
			EvidenceHash       string `json:"evidence_hash"`
			ProposerAddress    string `json:"proposer_address"`
		} `json:"header"`
		Data struct {
			Txs interface{} `json:"txs"`
		} `json:"data"`
		Evidence struct {
			Evidence interface{} `json:"evidence"`
		} `json:"evidence"`
		LastCommit struct {
			BlockID struct {
				Hash  string `json:"hash"`
				Parts struct {
					Total string `json:"total"`
					Hash  string `json:"hash"`
				} `json:"parts"`
			} `json:"block_id"`
			Precommits []struct {
				Type    int    `json:"type"`
				Height  string `json:"height"`
				Round   string `json:"round"`
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total string `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"block_id"`
				Timestamp        time.Time `json:"timestamp"`
				ValidatorAddress string    `json:"validator_address"`
				ValidatorIndex   string    `json:"validator_index"`
				Signature        string    `json:"signature"`
			} `json:"precommits"`
		} `json:"last_commit"`
	} `json:"block"`
}

type ValidatorSetVo struct {
	BlockHeight string `json:"block_height"`
	Validators  []struct {
		Address          string `json:"address"`
		PubKey           string `json:"pub_key"`
		ProposerPriority string `json:"proposer_priority"`
		VotingPower      string `json:"voting_power"`
	} `json:"validators"`
}

type StakePoolVo struct {
	LooseTokens  string `json:"loose_tokens"`
	BondedTokens string `json:"bonded_tokens"`
	TotalSupply  string `json:"total_supply"`
	BondedRatio  string `json:"bonded_ratio"`
}

type DelegationVo struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
	Shares        string `json:"shares"`
	Height        string `json:"height"`
}

type SignInfoVo struct {
	StartHeight         string    `json:"start_height"`
	IndexOffset         string    `json:"index_offset"`
	JailedUntil         time.Time `json:"jailed_until"`
	MissedBlocksCounter string    `json:"missed_blocks_counter"`
}

type BlockResultVo struct {
	Height  string `json:"height"`
	Results struct {
		DeliverTx []struct {
			Code      int         `json:"code"`
			Data      interface{} `json:"data"`
			Log       string      `json:"log"`
			Info      string      `json:"info"`
			GasWanted string      `json:"gas_wanted"`
			GasUsed   string      `json:"gas_used"`
			Tags      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"deliver_tx"`
		EndBlock struct {
			ValidatorUpdates []struct {
				PubKey struct {
					Type string `json:"type"`
					Data string `json:"data"`
				} `json:"pub_key"`
				Power string `json:"power"`
			} `json:"validator_updates"`
			ConsensusParamUpdates interface{} `json:"consensus_param_updates"`
			Tags                  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"end_block"`
		BeginBlock struct {
			Tags []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"begin_block"`
	} `json:"results"`
}
