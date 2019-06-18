package types

const (
	UrlRoot = "/api"

	//Account
	UrlRegisterQueryAccount  = "/account/{address}"
	UrlRegisterQueryAccounts = "/accounts"

	//home
	UrlRegisterNavigationBar = "/home/navigation"

	//Block
	UrlRegisterQueryBlockLatestHeight = "/block/latestheight"
	UrlRegisterQueryBlock             = "/block/{height}"
	UrlRegisterQueryRecentBlocks      = "/blocks/recent"
	UrlRegisterQueryBlocks            = "/blocks"
	UrlRegisterQueryBlockTxs          = "/block/txs/{height}"
	UrlRegisterQueryBlockTxGov        = "/block/txgov/{height}"
	UrlRegisterQueryBlockValidatorSet = "/block/validatorset/{height}"
	UrlRegisterQueryBlockInfo         = "/block/blockinfo/{height}"

	//Governance
	UrlRegisterQueryProposals              = "/gov/proposals"
	UrlRegisterQueryDepositVotingProposals = "/gov/deposit_voting_proposals"
	UrlRegisterQueryProposal               = "/gov/proposals/{pid}"
	UrlRegisterQueryGovParams              = "/gov/params"
	UrlRegisterQueryProposalsVoterTxs      = "/gov/proposals/{id}/voter_txs"
	UrlRegisterQueryProposalsDepositorTxs  = "/gov/proposals/{id}/depositor_txs"

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
	UrlRegisterGetValidators        = "/stake/validators"
	UrlRegisterGetValidator         = "/stake/validators/{address}"
	UrlRegisterQueryCandidatesTop   = "/stake/candidatesTop"
	UrlRegisterQueryCandidate       = "/stake/candidate/{address}"
	UrlRegisterQueryCandidateUptime = "/stake/candidate/{address}/uptime/{category}"
	UrlRegisterQueryCandidatePower  = "/stake/candidate/{address}/power/{category}"
	UrlRegisterQueryCandidateStatus = "/stake/candidate/{address}/status"

	//Tx
	UrlRegisterQueryTxList       = "/tx/{type}/{page}/{size}"
	UrlRegisterQueryRecentTx     = "/txs/recent"
	UrlRegisterQueryTxsCounter   = "/txs/statistics"
	UrlRegisterQueryTxsByAccount = "/txsByAddress/{address}/{page}/{size}"
	UrlRegisterQueryTxsByDay     = "/txsByDay"
	UrlRegisterQueryTx           = "/tx/{hash}"
	//version
	UrlRegisterQueryApiVersion = "/version"

	Format = "2006/01/02T15:04:05Z07:00"

	Change  = "powerChange"
	Slash   = "slash"
	Recover = "recover"
)

var (
	TxTypeTransfer                    = "Transfer"
	TxTypeBurn                        = "Burn"
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

	RoleList = []string{RoleValidator, RoleCandidate, RoleJailed}

	BankList        = []string{TxTypeTransfer, TxTypeBurn}
	DeclarationList = []string{TxTypeStakeCreateValidator, TxTypeStakeEditValidator, TxTypeUnjail}
	StakeList       = []string{TxTypeStakeDelegate, TxTypeBeginRedelegate, TxTypeSetWithdrawAddress, TxTypeStakeBeginUnbonding, TxTypeWithdrawDelegatorReward, TxTypeWithdrawDelegatorRewardsAll, TxTypeWithdrawValidatorRewardsAll}
	GovernanceList  = []string{TxTypeSubmitProposal, TxTypeDeposit, TxTypeVote}

	ForwardList      = []string{TxTypeBeginRedelegate}
	TxTypeExcludeGov = append(append(DeclarationList, StakeList...), BankList...)
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

type TxType int

const (
	_ TxType = iota
	Trans
	Declaration
	Stake
	Gov
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
	}
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
	panic(CodeUnSupportTx)
}
