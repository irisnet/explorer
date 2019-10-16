package types

import (
	"github.com/irisnet/explorer/backend/logger"
)

const (
	UrlRoot = "/api"

	//Account
	UrlRegisterQueryAccount                     = "/account/{address}"
	UrlRegisterQueryAccounts                    = "/accounts"
	UrlRegisterQueryAccountDelegations          = "/account/{address}/delegations"
	UrlRegisterQueryAccountUnbondingDelegations = "/account/{address}/unbonding_delegations"
	UrlRegisterQueryAccountRewards              = "/account/{address}/rewards"

	//home
	UrlRegisterNavigationBar = "/home/navigation"

	//Assets
	UrlRegisterQueryNativeAsset  = "/assets/native/txs"
	UrlRegisterQueryGatewayAsset = "/assets/gateway/txs"

	//Block
	UrlRegisterQueryBlockLatestHeight = "/block/latestheight"
	UrlRegisterQueryRecentBlocks      = "/blocks/recent"
	UrlRegisterQueryBlocks            = "/blocks"
	UrlRegisterQueryBlockValidatorSet = "/block/validatorset/{height}"
	UrlRegisterQueryBlockInfo         = "/block/blockinfo/{height}"

	//Governance
	UrlRegisterQueryProposals              = "/gov/proposals"
	UrlRegisterQueryDepositVotingProposals = "/gov/deposit_voting_proposals"
	UrlRegisterQueryProposal               = "/gov/proposals/{pid}"
	UrlRegisterQueryProposalsVoterTxs      = "/gov/proposals/{id}/voter_txs"
	UrlRegisterQueryProposalsDepositorTxs  = "/gov/proposals/{id}/depositor_txs"
	UrlRegisterQueryProposalDeposit        = "/gov/proposal_deposit/{id}"
	UrlRegisterQueryProposalVoting         = "/gov/proposal_voting/{id}"

	UrlRegisterQueryParams = "/params"

	//SearchBox
	UrlRegisterQueryText    = "/search/{text}"
	UrlRegisterQuerySysDate = "/sysdate"
	UrlRegisterQueryConfig  = "/config"

	//faucet
	UrlRegisterQueryFaucet = "/faucet/account"
	UrlRegisterApply       = "/faucet/apply/{address}"

	UrlFaucetAccountService = "%s/account"
	UrlFaucetApplyService   = "%s/apply"

	//Stake
	UrlRegisterGetValidators                         = "/stake/validators"
	UrlRegisterUpdateIcons                           = "/stake/validators/update_icons"
	UrlRegisterGetValidator                          = "/stake/validators/{address}"
	UrlRegisterQueryCandidatesTop                    = "/stake/candidatesTop"
	UrlRegisterQueryCandidate                        = "/stake/candidate/{address}"
	UrlRegisterQueryCandidateStatus                  = "/stake/candidate/{address}/status"
	UrlRegisterQueryValidatorsDelegations            = "/stake/validators/{validatorAddr}/delegations"
	UrlRegisterQueryValidatorUnbondingDelegations    = "/stake/validators/{validatorAddr}/unbonding-delegations"
	UrlRegisterQueryValidatorRedelegations           = "/stake/validators/{validatorAddr}/redelegations"
	UrlRegisterQueryValidatorVoteByValidatorAddr     = "/stake/validators/{validatorAddr}/vote"
	UrlRegisterQueryDepositorTxsByValidatorAddr      = "/stake/validators/{validatorAddr}/depositor_txs"
	UrlRegisterQueryWithdrawAddrByValidatorAddr      = "/stake/validators/{validatorAddr}/withdraw-addr"
	UrlRegisterQueryCommissionRewardsByValidatorAddr = "/stake/validators/{validatorAddr}/commission-rewards"

	//Tx
	UrlRegisterQueryTxList       = "/txs"
	UrlRegisterQueryTxListByType = "/txs/{type}/{page}/{size}"
	UrlRegisterQueryRecentTx     = "/txs/recent"
	UrlRegisterQueryTxsCounter   = "/txs/statistics"
	UrlRegisterQueryTxsByAccount = "/txsByAddress/{address}/{page}/{size}"
	UrlRegisterQueryTxsByDay     = "/txsByDay"
	UrlRegisterQueryTx           = "/tx/{hash}"
	UrlRegisterQueryTxType       = "/tx_types/{type}"
	//tokenstats
	UrlRegisterQueryTokenStats    = "/tokenstats"
	UrlRegisterTokensAccountTotal = "/tokenstats/account_total"
	//bondedtokens
	UrlRegisterBondedTokensValidators = "/bondedtokens/validators"
	//assetTokens
	UrlRegisterAssetTokens = "/asset/tokens"
	//assetToken Detail
	UrlRegisterAssetTokenInfo = "/asset/tokens/{token_id}"
	//assetGateway
	UrlRegisterAssetGatewayInfo = "/asset/gateways/{moniker}"
	//version
	UrlRegisterQueryApiVersion = "/version"

	// cron task
	UrlRegisterDoCronTaskAssetGateways   = "/task/asset_gateways"
	UrlRegisterDoCronTaskAssetTokens     = "/task/asset_tokens"
	UrlRegisterDoCronTaskGovParams       = "/task/gov_params"
	UrlRegisterDoCronTaskTxNumByDay      = "/task/tx_num_by_day"
	UrlRegisterDoCronTaskValidators      = "/task/validators"
	UrlRegisterDoCronTaskValidatorIcons  = "/task/validator_icons"
	UrlRegisterTaskUpdateBlackValidators = "/task/black_validators"

	Format = "2006/01/02T15:04:05Z07:00"

	Change  = "powerChange"
	Slash   = "slash"
	Recover = "recover"

	TxTag_WithDrawRewardFromValidator = "withdraw-reward-from-validator-"
	TxTag_WithDrawAddress             = "withdraw-address"

	IRISUint   = "iris"
	IRISAttoUint   = "iris-atto"
	AssetMinDenom = "-min"
)

var (
	TxTypeTransfer      = "Transfer"
	TxTypeBurn          = "Burn"
	TxTypeSetMemoRegexp = "SetMemoRegexp"

	TxTypeStakeCreateValidator        = "CreateValidator"
	TxTypeStakeEditValidator          = "EditValidator"
	TxTypeStakeDelegate               = "Delegate"
	TxTypeStakeBeginUnbonding         = "BeginUnbonding"
	TxTypeBeginRedelegate             = "BeginRedelegate"
	TxTypeUnjail                      = "Unjail"
	TxTypeSetWithdrawAddress          = "SetWithdrawAddress"
	TxTypeWithdrawDelegatorReward     = "WithdrawDelegatorReward"
	TxTypeWithdrawDelegatorRewardsAll = "WithdrawDelegatorRewardsAll"
	TxTypeWithdrawValidatorRewardsAll = "WithdrawValidatorRewardsAll"
	TxTypeSubmitProposal              = "SubmitProposal"
	TxTypeDeposit                     = "Deposit"
	TxTypeVote                        = "Vote"

	TxTypeIssueToken           = "IssueToken"
	TxTypeEditToken            = "EditToken"
	TxTypeMintToken            = "MintToken"
	TxTypeTransferTokenOwner   = "TransferTokenOwner"
	TxTypeCreateGateway        = "CreateGateway"
	TxTypeEditGateway          = "EditGateway"
	TxTypeTransferGatewayOwner = "TransferGatewayOwner"

	TxTypeRequestRand = "RequestRand"

	TypeValStatusUnbonded  = "Unbonded"
	TypeValStatusUnbonding = "Unbonding"
	TypeValStatusBonded    = "Bonded"

	TxTypeStatus = "success"

	Unbonded  = 0x00
	Unbonding = 0x01
	Bonded    = 0x02

	RoleValidator = "validator"
	RoleCandidate = "candidate"
	RoleJailed    = "jailed"

	//RoleList = []string{RoleValidator, RoleCandidate, RoleJailed}

	BankList        = []string{TxTypeTransfer, TxTypeBurn, TxTypeSetMemoRegexp}
	DeclarationList = []string{TxTypeStakeCreateValidator, TxTypeStakeEditValidator, TxTypeUnjail}
	StakeList       = []string{TxTypeStakeDelegate, TxTypeBeginRedelegate, TxTypeSetWithdrawAddress, TxTypeStakeBeginUnbonding, TxTypeWithdrawDelegatorReward, TxTypeWithdrawDelegatorRewardsAll, TxTypeWithdrawValidatorRewardsAll}
	GovernanceList  = []string{TxTypeSubmitProposal, TxTypeDeposit, TxTypeVote}

	ForwardList = []string{TxTypeBeginRedelegate}
	//TxTypeExcludeGov = append(append(DeclarationList, StakeList...), BankList...)
	AssetList = []string{TxTypeIssueToken, TxTypeEditToken, TxTypeMintToken, TxTypeTransferTokenOwner, TxTypeCreateGateway, TxTypeEditGateway, TxTypeTransferGatewayOwner}
	RandList  = []string{TxTypeRequestRand}
)

func IsDeclarationType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range DeclarationList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsBankType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range BankList {
		if t == typ {
			return true
		}
	}
	return false
}
func IsStakeType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range StakeList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsGovernanceType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range GovernanceList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsAssetType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range AssetList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsRandType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range RandList {
		if t == typ {
			return true
		}
	}
	return false
}

type TxType int

const (
	_           TxType = iota
	Trans
	Declaration
	Stake
	Gov
	Asset
	Rand
)

func Convert(typ string) TxType {
	if IsBankType(typ) {
		return Trans
	} else if IsStakeType(typ) {
		return Stake
	} else if IsDeclarationType(typ) {
		return Declaration
	} else if IsGovernanceType(typ) {
		return Gov
	} else if IsAssetType(typ) {
		return Asset
	} else if IsRandType(typ) {
		return Rand
	}
	logger.Error("Convert UnSupportTx Type", logger.String("txtype", typ))
	panic(CodeUnSupportTx)
}
func TxTypeFromString(typ string) TxType {
	if typ == "trans" {
		return Trans
	} else if typ == "stake" {
		return Stake
	} else if typ == "declaration" {
		return Declaration
	} else if typ == "gov" {
		return Gov
	}
	logger.Error("TxTypeFromString UnSupportTx Type", logger.String("txtype", typ))
	panic(CodeUnSupportTx)
}
