package types

const (
	UrlRoot = "/api"

	//Account
	UrlRegisterQueryAccount = "/account/{address}"

	//home
	UrlRegisterNavigationBar = "/home/navigation"

	//Block
	UrlRegisterQueryBlock        = "/block/{height}"
	UrlRegisterQueryRecentBlocks = "/blocks/recent"
	UrlRegisterQueryBlocks       = "/blocks"

	//Proposal
	UrlRegisterQueryProposals = "/proposals/{page}/{size}"
	UrlRegisterQueryProposal  = "/proposal/{pid}"

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
	TypeTransfer                      = "Transfer"
	TypeCreateValidator               = "CreateValidator"
	TypeEditValidator                 = "EditValidator"
	TypeUnjail                        = "Unjail"
	TypeDelegate                      = "Delegate"
	TypeBeginRedelegation             = "BeginRedelegate"
	TypeBeginUnbonding                = "BeginUnbonding"
	TxTypeSetWithdrawAddress          = "SetWithdrawAddress"
	TxTypeWithdrawDelegatorReward     = "WithdrawDelegatorReward"
	TxTypeWithdrawDelegatorRewardsAll = "WithdrawDelegatorRewardsAll"
	TxTypeWithdrawValidatorRewardsAll = "WithdrawValidatorRewardsAll"
	TypeSubmitProposal                = "SubmitProposal"
	TypeDeposit                       = "Deposit"
	TypeVote                          = "Vote"

	TypeValStatusUnbonded  = "Unbonded"
	TypeValStatusUnbonding = "Unbonding"
	TypeValStatusBonded    = "Bonded"

	Unbonded  = 0x00
	Unbonding = 0x01
	Bonded    = 0x02

	RoleValidator = "validator"
	RoleCandidate = "candidate"
	RoleJailed    = "jailed"

	DeclarationList = []string{TypeCreateValidator, TypeEditValidator, TypeUnjail}
	StakeList       = []string{TypeDelegate, TypeBeginRedelegation, TxTypeSetWithdrawAddress, TypeBeginUnbonding, TxTypeWithdrawDelegatorReward, TxTypeWithdrawDelegatorRewardsAll, TxTypeWithdrawValidatorRewardsAll}
	GovernanceList  = []string{TypeSubmitProposal, TypeDeposit, TypeVote}
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
	if typ == TypeTransfer {
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
