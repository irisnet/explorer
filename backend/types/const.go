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

	//Htlc
	UrlRegisterQueryHtlc    = "/htlcs/{hash_lock}"
	UrlRegisterQueryHtlcTxs = "/htlcs/{hash_lock}/txs"

	//SearchBox
	UrlRegisterQueryText    = "/search/{text}"
	UrlRegisterQuerySysDate = "/sysdate"
	UrlRegisterQueryConfig  = "/config"

	//faucet
	UrlRegisterQueryFaucet = "/faucet/account"
	UrlRegisterApply       = "/faucet/apply/{address}"

	UrlFaucetAccountService = "%s/account"
	UrlFaucetApplyService   = "%s/apply"

	//Service
	UrlRegisterQueryServiceBindings = "/service/bindings/{servicename}"
	UrlRegisterQueryServiceRequest  = "/service/contexts/{contextid}"

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
	UrlRegisterQueryCommissionInfo                   = "/stake/commission_info"

	//Tx
	UrlRegisterQueryTxList       = "/txs"
	UrlRegisterQueryTxListByType = "/txs/{type}/{page}/{size}"
	UrlRegisterQueryRecentTx     = "/txs/recent"
	//UrlRegisterQueryTxsCounter   = "/txs/statistics"
	UrlRegisterQueryTxsByAccount = "/txsByAddress/{address}/{page}/{size}"
	UrlRegisterQueryTxsByDay     = "/txsByDay"
	UrlRegisterQueryTx           = "/tx/{hash}"
	UrlRegisterQueryTxType       = "/tx_types/{type}"
	//tokenstats
	UrlRegisterQueryTokenStats    = "/tokenstats"
	UrlRegisterQueryBaseDenom     = "/unit_info"
	UrlRegisterTokensAccountTotal = "/tokenstats/account_total"
	//bondedtokens
	UrlRegisterBondedTokensValidators = "/bondedtokens/validators"
	//assetTokens
	UrlRegisterAssetTokens = "/asset/tokens"
	//assetToken Detail
	UrlRegisterAssetTokenInfo = "/asset/tokens/{symbol}"
	//assetGateway
	//UrlRegisterAssetGatewayInfo = "/asset/gateways/{moniker}"
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

	TimeLayout             = "2006-01-02T15:04:05"
	TimeLayout1            = "2006-01-02 15:04:05 +0000 UTC"
	DelegatorRewardTag     = "DelegatorReward"
	ValidatorRewardTag     = "ValidatorReward"
	ValidatorCommissionTag = "ValidatorCommission"
	Change                 = "powerChange"
	Slash                  = "slash"
	Recover                = "recover"

	TxTag_WithDrawRewardFromValidator = "withdraw-reward-from-validator-"
	//TxTag_WithDrawAddress             = "withdraw-address"

	StakeUint     = "stake"
	AssetMinDenom = "-min"
	Unknown       = "unknown"

	Success            = "success"
	TaskConTrolMonitor = "task_control_monitor"
)

var (
	TxTypeTransfer  = "Transfer"
	TxTypeMultiSend = "MultiSend"
	//TxTypeBurn          = "Burn"
	//TxTypeSetMemoRegexp = "SetMemoRegexp"

	TxTypeStakeCreateValidator    = "CreateValidator"
	TxTypeStakeEditValidator      = "EditValidator"
	TxTypeStakeDelegate           = "Delegate"
	TxTypeStakeBeginUnbonding     = "BeginUnbonding"
	TxTypeBeginRedelegate         = "BeginRedelegate"
	TxTypeUnjail                  = "Unjail"
	TxTypeSetWithdrawAddress      = "SetWithdrawAddress"
	TxTypeWithdrawDelegatorReward = "WithdrawDelegatorReward"
	//TxTypeWithdrawDelegatorRewardsAll      = "WithdrawDelegatorRewardsAll"
	//TxTypeWithdrawValidatorRewardsAll      = "WithdrawValidatorRewardsAll"
	TxTypeMsgFundCommunityPool           = "FundCommunityPool"
	TxTypeMsgWithdrawValidatorCommission = "WithdrawValidatorCommission"
	TxTypeSubmitProposal                 = "SubmitProposal"
	//TxMsgTypeSubmitSoftwareUpgradeProposal = "SubmitSoftwareUpgradeProposal"
	//TxMsgTypeSubmitTaxUsageProposal        = "SubmitTaxUsageProposal"
	//TxMsgTypeSubmitTokenAdditionProposal   = "SubmitTokenAdditionProposal"
	TxTypeDeposit = "Deposit"
	TxTypeVote    = "Vote"

	TxTypeIssueToken         = "IssueToken"
	TxTypeEditToken          = "EditToken"
	TxTypeMintToken          = "MintToken"
	TxTypeTransferTokenOwner = "TransferTokenOwner"
	//TxTypeCreateGateway        = "CreateGateway"
	//TxTypeEditGateway          = "EditGateway"
	//TxTypeTransferGatewayOwner = "TransferGatewayOwner"

	TxTypeAddProfiler    = "AddProfiler"
	TxTypeAddTrustee     = "AddTrustee"
	TxTypeDeleteTrustee  = "DeleteTrustee"
	TxTypeDeleteProfiler = "DeleteProfiler"

	TxTypeCreateHTLC = "CreateHTLC"
	TxTypeClaimHTLC  = "ClaimHTLC"
	TxTypeRefundHTLC = "RefundHTLC"

	TxTypeAddLiquidity    = "AddLiquidity"
	TxTypeRemoveLiquidity = "RemoveLiquidity"
	TxTypeSwapOrder       = "SwapOrder"

	TxTypeRequestRand = "RequestRand"

	TypeValStatusUnbonded  = "Unbonded"
	TypeValStatusUnbonding = "Unbonding"
	TypeValStatusBonded    = "Bonded"

	TxTypeNFTMint     = "NFTMint"
	TxTypeNFTEdit     = "NFTEdit"
	TxTypeNFTTransfer = "NFTTransfer"
	TxTypeNFTBurn     = "NFTBurn"
	TxTypeIssueDenom  = "IssueDenom"

	TxTypeDefineService             = "DefineService"              // type for MsgDefineService
	TxTypeBindService               = "BindService"                // type for MsgBindService
	TxTypeUpdateServiceBinding      = "UpdateServiceBinding"       // type for MsgUpdateServiceBinding
	TxTypeServiceSetWithdrawAddress = "service/SetWithdrawAddress" // type for SetWithdrawAddress
	TxTypeDisableServiceBinding     = "DisableServiceBinding"      // type for MsgDisableServiceBinding
	TxTypeEnableServiceBinding      = "EnableServiceBinding"       // type for MsgEnableServiceBinding
	TxTypeRefundServiceDeposit      = "RefundServiceDeposit"       // type for MsgRefundServiceDeposit
	TxTypeCallService               = "CallService"                // type for MsgCallService
	TxTypeRespondService            = "RespondService"             // type for MsgRespondService
	TxTypePauseRequestContext       = "PauseRequestContext"        // type for MsgPauseRequestContext
	TxTypeStartRequestContext       = "StartRequestContext"        // type for MsgStartRequestContext
	TxTypeKillRequestContext        = "KillRequestContext"         // type for MsgKillRequestContext
	TxTypeUpdateRequestContext      = "UpdateRequestContext"       // type for MsgUpdateRequestContext
	TxTypeWithdrawEarnedFees        = "WithdrawEarnedFees"         // type for MsgWithdrawEarnedFees

	TxTypeCreateFeed = "CreateFeed"
	TxTypeEditFeed   = "EditFeed"
	TxTypePauseFeed  = "PauseFeed"
	TxTypeStartFeed  = "StartFeed"

	TxTypeSubmitEvidence  = "SubmitEvidence"
	TxTypeVerifyInvariant = "VerifyInvariant"

	TxTypeStatus = "success"

	Unbonded  = 0x00
	Unbonding = 0x01
	Bonded    = 0x02

	RoleValidator = "validator"
	RoleCandidate = "candidate"
	RoleJailed    = "jailed"

	//RoleList = []string{RoleValidator, RoleCandidate, RoleJailed}

	BankList        = []string{TxTypeTransfer, TxTypeMultiSend}
	DeclarationList = []string{TxTypeStakeCreateValidator, TxTypeStakeEditValidator, TxTypeUnjail}
	StakeList       = []string{TxTypeStakeDelegate, TxTypeBeginRedelegate, TxTypeSetWithdrawAddress, TxTypeStakeBeginUnbonding, TxTypeWithdrawDelegatorReward, TxTypeMsgFundCommunityPool, TxTypeMsgWithdrawValidatorCommission}
	GovernanceList  = []string{TxTypeSubmitProposal, TxTypeDeposit, TxTypeVote}
	GuardianList    = []string{TxTypeAddProfiler, TxTypeAddTrustee, TxTypeDeleteProfiler, TxTypeDeleteTrustee}
	HTLCList        = []string{TxTypeClaimHTLC, TxTypeCreateHTLC, TxTypeRefundHTLC}
	CoinswapList    = []string{TxTypeAddLiquidity, TxTypeRemoveLiquidity, TxTypeSwapOrder}
	SlashingList    = []string{TxTypeVerifyInvariant}

	//ForwardList = []string{TxTypeBeginRedelegate}
	//TxTypeExcludeGov = append(append(DeclarationList, StakeList...), BankList...)
	AssetList   = []string{TxTypeIssueToken, TxTypeEditToken, TxTypeMintToken, TxTypeTransferTokenOwner}
	RandList    = []string{TxTypeRequestRand}
	OrcaleList  = []string{TxTypeCreateFeed, TxTypeStartFeed, TxTypeEditFeed, TxTypePauseFeed}
	NftList     = []string{TxTypeNFTBurn, TxTypeNFTEdit, TxTypeNFTMint, TxTypeNFTTransfer, TxTypeIssueDenom}
	ServiceList = []string{TxTypeDefineService, TxTypeBindService, TxTypeCallService, TxTypeRespondService,
		TxTypeDisableServiceBinding, TxTypeEnableServiceBinding, TxTypeUpdateServiceBinding,
		TxTypeStartRequestContext, TxTypeKillRequestContext, TxTypePauseRequestContext, TxTypeUpdateRequestContext,
		TxTypeServiceSetWithdrawAddress, TxTypeRefundServiceDeposit, TxTypeWithdrawEarnedFees}

	ProfilerAddrList = map[string]string{
		"iaa1v6c3sa76s3grss3xu64tvn9nd556jlcw6azc85": "Genesis Profiler 1",
		"iaa1us4dhfyq66hzeu2tu06tvvr90cy33vg3yn2uld": "Genesis Profiler 2",
	}

	FoundationDelegatorAddr = "iaa1w7ewedr57z6p7f8nknmdvukfxwkwlsvfjumdts"
	MainnetAccPrefix        = "iaa"
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

func IsCoinswapType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range CoinswapList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsHTLCType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range HTLCList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsGuardianType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range GuardianList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsNftType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range NftList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsServiceType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range ServiceList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsOrcaleType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range OrcaleList {
		if t == typ {
			return true
		}
	}
	return false
}
func IsSlashingType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range SlashingList {
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
	Guardian
	Htlc
	Coinswap
	Nft
	Service
	Orcale
	Slashing
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
	} else if IsCoinswapType(typ) {
		return Coinswap
	} else if IsHTLCType(typ) {
		return Htlc
	} else if IsGuardianType(typ) {
		return Guardian
	} else if IsNftType(typ) {
		return Nft
	} else if IsOrcaleType(typ) {
		return Orcale
	} else if IsServiceType(typ) {
		return Service
	} else if IsSlashingType(typ) {
		return Slashing
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
	} else if typ == "nft" {
		return Nft
	} else if typ == "service" {
		return Service
	} else if typ == "orale" {
		return Orcale
	}
	logger.Error("TxTypeFromString UnSupportTx Type", logger.String("txtype", typ))
	panic(CodeUnSupportTx)
}
