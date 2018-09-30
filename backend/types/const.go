package types

var (
	TypeTransfer             = "Transfer"
	TypeCreateValidator      = "CreateValidator"
	TypeEditValidator        = "EditValidator"
	TypeDelegate             = "Delegate"
	TypeBeginRedelegation    = "BeginRedelegation"
	TypeCompleteRedelegation = "CompleteRedelegation"
	TypeBeginUnbonding       = "BeginUnbonding"
	TypeCompleteUnbonding    = "CompleteUnbonding"
	TypeSubmitProposal       = "SubmitProposal"
	TypeDeposit              = "Deposit"
	TypeVote                 = "Vote"

	DeclarationList = []string{TypeCreateValidator, TypeEditValidator}
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
	} else if typ == "declar" {
		return Declaration
	} else if typ == "gov" {
		return Gov
	}
	panic("invalid tx type")
}
