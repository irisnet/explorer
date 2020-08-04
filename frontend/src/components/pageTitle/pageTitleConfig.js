import store from "../../store";

export default {
	BlockchainBlocks:'Block List',
	BlockchainBlocksBlockDetails:'Block Details',
	BlockchainTransactions:'Transaction List',
	BlockchainTransactionsTransactionDetails:'Transaction Details',
	StakingValidators:'Validator List',
	StakingValidatorsValidatorDetails:'Validator Details',
	StakingDelegationTxs:'Delegation Txs List',
	StakingValidationTxs:'Validation Txs List',
	Transfer:'Transfer List',
	AssetNativeAsset:'Native Asset List',
	AssetNativeAssetTxs:'Native Asset Txs List',
	GovParameters:'Parameter List',
	GovProposals:'Proposal List',
	GovProposalsProposalDetails:'Proposal Details',
	GovGovTxs:'Gov Txs List',
	StatsIRISRichList:`${store.state.nativeToken} Rich List`,
	StatsIRISRichListAddress:'Address',
	StatsIRISStats:`${store.state.displayToken} Stats Analyse`,
}
