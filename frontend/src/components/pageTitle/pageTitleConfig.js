import store from "../../store";
import Tools from "../../util/Tools";

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
	StatsIRISRichList:`${Tools.firstWordUpperCase(store.state.displayToken)} Rich List`,
	StatsIRISRichListAddress:'Address',
	StatsIRISStats:`${Tools.firstWordUpperCase(store.state.displayToken)} Stats Analyse`,
}
