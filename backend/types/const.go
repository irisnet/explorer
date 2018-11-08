package types

const (
	//Account
	UrlRegisterQueryAccount    = "/account/{address}"
	UrlRegisterQueryAllAccount = "/accounts/{page}/{size}"

	//Block
	UrlRegisterQueryBlock            = "/block/{height}"
	UrlRegisterQueryBlocks           = "/blocks/{page}/{size}"
	UrlRegisterQueryBlocksPrecommits = "/blocks/precommits/{address}/{page}/{size}"

	//Chain
	UrlRegisterQueryChain = "/chain/status"

	//Location
	UrlRegisterQueryIp = "/ip/"

	//Node
	UrlRegisterQueryNodes = "/net_info"

	//Proposal
	UrlRegisterQueryProposals = "/proposals/{page}/{size}"
	UrlRegisterQueryProposal  = "/proposal/{pid}"

	//SearchBox
	UrlRegisterQueryText = "/search/{text}"

	//Stake
	UrlRegisterQueryValidator        = "/stake/validators/{page}/{size}"
	UrlRegisterQueryRevokedValidator = "/stake/revokedVal/{page}/{size}"
	UrlRegisterQueryCandidates       = "/stake/candidates/{page}/{size}"
	UrlRegisterQueryCandidatesTop    = "/stake/candidatesTop"
	UrlRegisterQueryCandidate        = "/stake/candidate/{address}"
	UrlRegisterQueryCandidateUptime  = "/stake/candidate/{address}/uptime/{category}"
	UrlRegisterQueryCandidatePower   = "/stake/candidate/{address}/power/{category}"
	UrlRegisterQueryCandidateStatus  = "/stake/candidate/{address}/status"

	//Tx
	UrlRegisterQueryTxList               = "/tx/{type}/{page}/{size}"
	UrlRegisterQueryAllStakeTxByPage     = "/txs/stake/{page}/{size}"
	UrlRegisterQueryStakeTxByAccount     = "/tx/stake/{address}"
	UrlRegisterQueryPageStakeTxByAccount = "/tx/stake/{address}/{page}/{size}"
	UrlRegisterQueryTxs                  = "/txs/{page}/{size}"
	UrlRegisterQueryTxsCounter           = "/txs/statistics"
	UrlRegisterQueryTxsByAccount         = "/txsByAddress/{address}/{page}/{size}"
	UrlRegisterQueryTxsByBlock           = "/txsByBlock/{height}/{page}/{size}"
	UrlRegisterQueryTxsByDay             = "/txsByDay"
	UrlRegisterQueryTx                   = "/tx/{hash}"
)

var (
	TypeTransfer             = "Transfer"
	TypeCreateValidator      = "CreateValidator"
	TypeEditValidator        = "EditValidator"
	TypeUnRevoke             = "Unrevoke"
	TypeDelegate             = "Delegate"
	TypeBeginRedelegation    = "BeginRedelegate"
	TypeCompleteRedelegation = "CompleteRedelegate"
	TypeBeginUnbonding       = "BeginUnbonding"
	TypeCompleteUnbonding    = "CompleteUnbonding"
	TypeSubmitProposal       = "SubmitProposal"
	TypeDeposit              = "Deposit"
	TypeVote                 = "Vote"

	TypeValStatusUnbonded  = "Unbonded"
	TypeValStatusUnbonding = "Unbonding"
	TypeValStatusBonded    = "Bonded"

	DeclarationList = []string{TypeCreateValidator, TypeEditValidator, TypeUnRevoke}
	StakeList       = []string{TypeDelegate, TypeBeginRedelegation, TypeCompleteRedelegation, TypeBeginUnbonding, TypeCompleteUnbonding}
	GovernanceList  = []string{TypeSubmitProposal, TypeDeposit, TypeVote}
)

func isDeclarationType(typ string) bool {
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

func isStakeType(typ string) bool {
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

func isGovernanceType(typ string) bool {
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
	} else if isStakeType(typ) {
		return Stake
	} else if isDeclarationType(typ) {
		return Declaration
	} else if isGovernanceType(typ) {
		return Gov
	}
	panic("invalid tx type")
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
	panic("invalid tx type")
}
