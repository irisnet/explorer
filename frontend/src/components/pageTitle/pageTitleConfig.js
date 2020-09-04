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
	StatsIRISRichList:`${sessionStorage.getItem('displayToken').toLocaleUpperCase()} Rich List`,
	StatsIRISRichListAddress:'Address',
	StatsIRISStats:`${sessionStorage.getItem('displayToken').toLocaleUpperCase()} Stats Analyse`,
}
