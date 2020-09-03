import Tools from "../util/Tools"
export default class FormatTxType {
	static formatTxType(txTypeArray){
		
		let allTxType = [],
			tansferObj = {
				value:'transfer',
				label:'Transfer',
				children:[]
			},
			delegationObj = {
				value:'delegation',
				label: 'Delegation',
				children:[]
			},
			validationObj = {
				value:'validation',
				label:'Validation',
				children:[]
			},
			govObj = {
				value:'gov',
				label:'Gov',
				children:[]
			},
            nftObj = {
				value:'nft',
				label:'NFT',
				children:[]
			},
            oracleObj = {
				value:'oracle',
				label:'Oracle',
				children:[]
			},
            coinswapObj = {
				value:'coinswap',
				label:'Coinswap',
				children:[]
			},
            assetObj = {
				value:'asset',
				label:'Asset',
				children:[]
			},
            iServiceObj = {
				value:'iService',
				label:'iService',
				children:[]
			},



			othersObj = {
				value:'others',
				label:'Others',
				children:[]
			};
		txTypeArray.forEach( item => {

			switch (item) {
				case "Transfer":
					tansferObj.children.push({
						value:'transfer',
						label:'Transfer'
					});
					break;
				case "MultiSend":
					// tansferObj.children.push({
					// 	value:'multiSend',
					// 	label:'MultiSend'
					// });
					break;
				case "Burn":
					tansferObj.children.push({
						value:'burn',
						label:'Burn'
					});
					break;
				case "SetMemoRegexp":
					tansferObj.children.push({
						value:'setMemoRegexp',
						label:'SetMemoRegexp'
					});
					break;
				case "Delegate":
					delegationObj.children.push({
						value:'delegate',
						label:'Delegate'
					});
					break;
				case "BeginRedelegate":
					delegationObj.children.push({
						value:'beginRedelegate',
						label:'BeginRedelegate'
					});
					break;
				case "SetWithdrawAddress":
					delegationObj.children.push({
						value:'setWithdrawAddress',
						label:'SetWithdrawAddress'
					});
					break;
				case "BeginUnbonding":
					delegationObj.children.push({
						value:'beginUnbonding',
						label:'BeginUnbonding'
					});
					break;
				case "WithdrawDelegatorReward":
					delegationObj.children.push({
						value:'withdrawDelegatorReward',
						label:'WithdrawDelegatorReward'
					});
					break;
				case "WithdrawDelegatorRewardsAll":
					delegationObj.children.push({
						value:'withdrawDelegatorRewardsAll',
						label:'WithdrawDelegatorRewardsAll'
					});
					break;
                case "FundCommunityPool":
                    delegationObj.children.push({
                        value:'FundCommunityPool',
                        label:'FundCommunityPool'
                    });
                    break;
                // case "WithdrawValidatorCommission":
                //     delegationObj.children.push({
                //         value:'WithdrawValidatorCommission',
                //         label:'WithdrawValidatorCommission'
                //     });
                //     break;

				case "CreateValidator":
					validationObj.children.push({
						value:'createValidator',
						label:'CreateValidator'
					});
					break;
				case "EditValidator":
					validationObj.children.push({
						value:'editValidator',
						label:'EditValidator'
					});
					break;
				case "Unjail":
					validationObj.children.push({
						value:'unjail',
						label:'Unjail'
					});
					break;
				case "WithdrawValidatorRewardsAll":
					validationObj.children.push({
						value:'withdrawValidatorRewardsAll',
						label:'WithdrawValidatorRewardsAll'
					});
					break;
				case "SubmitProposal":
					govObj.children.push({
						value:'submitProposal',
						label:'SubmitProposal'
					});
					break;
				case "Deposit":
					govObj.children.push({
						value:'deposit',
						label:'Deposit'
					});
					break;
				case "Vote":
					govObj.children.push({
						value:'vote',
						label:'Vote'
					});
					break;

				case "IssueDenom":
					nftObj.children.push({
						value:'IssueDenom',
						label:'IssueDenom'
					});
					break;
				case "NFTEdit":
					nftObj.children.push({
						value:'NFTEdit',
						label:'NFTEdit'
					});
					break;
				case "NFTTransfer":
					nftObj.children.push({
						value:'NFTTransfer',
						label:'NFTTransfer'
					});
					break;
				case "NFTMint":
					nftObj.children.push({
						value:'NFTMint',
						label:'NFTMint'
					});
					break;
				case "NFTBurn":
					nftObj.children.push({
						value:'NFTBurn',
						label:'NFTBurn'
					});
					break;
				case "CreateFeed":
                    oracleObj.children.push({
						value:'CreateFeed',
						label:'CreateFeed'
					});
					break;
				case "StartFeed":
                    oracleObj.children.push({
						value:'StartFeed',
						label:'StartFeed'
					});
					break;
				case "PauseFeed":
                    oracleObj.children.push({
						value:'PauseFeed',
						label:'PauseFeed'
					});
					break;
				case "EditFeed":
                    oracleObj.children.push({
						value:'EditFeed',
						label:'EditFeed'
					});
					break;
				case "IssueToken":
                    assetObj.children.push({
						value:'IssueToken',
						label:'IssueToken'
					});
					break;
				case "EditToken":
                    assetObj.children.push({
						value:'EditToken',
						label:'EditToken'
					});
					break;
				case "MintToken":
                    assetObj.children.push({
						value:'MintToken',
						label:'MintToken'
					});
					break;
				case "TransferTokenOwner":
                    assetObj.children.push({
						value:'TransferTokenOwner',
						label:'TransferTokenOwner'
					});
					break;
				case "CreateGateway":
					othersObj.children.push({
						value:'createGateway',
						label:'CreateGateway'
					});
					break;
				case "EditGateway":
					othersObj.children.push({
						value:'editGateway',
						label:'EditGateway'
					});
					break;
				case "TransferGatewayOwner":
					othersObj.children.push({
						value:'transferGatewayOwner',
						label:'TransferGatewayOwner'
					});
					break;
				case "RequestRand":
					othersObj.children.push({
						value:'requestRand',
						label:'RequestRand'
					});
					break;
				case "AddTrustee":
					othersObj.children.push({
						value:'addTrustee',
						label:'AddTrustee'
					});
					break;
				case "DeleteProfiler":
					othersObj.children.push({
						value:'deleteProfiler',
						label:'DeleteProfiler'
					});
					break;
				case "DeleteTrustee":
					othersObj.children.push({
						value:'deleteTrustee',
						label:'DeleteTrustee'
					});
					break;
				case "ClaimHTLC":
					othersObj.children.push({
						value:'claimHTLC',
						label:'ClaimHTLC'
					});
					break;
				case "CreateHTLC":
					othersObj.children.push({
						value:'createHTLC',
						label:'CreateHTLC'
					});
					break;
				case "RefundHTLC":
					othersObj.children.push({
						value:'refundHTLC',
						label:'RefundHTLC'
					});
					break;
				case "VerifyInvariant":
					othersObj.children.push({
						value:'verifyInvariant',
						label:'VerifyInvariant'
					});
					break;

				case "AddLiquidity":
					coinswapObj.children.push({
						value:'addLiquidity',
						label:'AddLiquidity'
					});
					break;
				case "RemoveLiquidity":
					coinswapObj.children.push({
						value:'removeLiquidity',
						label:'RemoveLiquidity'
					});
					break;
				case "SwapOrder":
					coinswapObj.children.push({
						value:'swapOrder',
						label:'SwapOrder'
					});
					break;
				case "AddProfiler":
					othersObj.children.push({
						value:'AddProfiler',
						label:'AddProfiler'
					});
					break;

				case "DefineService":
                    iServiceObj.children.push({
						value:'DefineService',
						label:'DefineService'
					});
					break;
				case "BindService":
                    iServiceObj.children.push({
						value:'BindService',
						label:'BindService'
					});
					break;
				case "CallService":
                    iServiceObj.children.push({
						value:'CallService',
						label:'CallService'
					});
					break;
                case "service/SetWithdrawAddress":
                    iServiceObj.children.push({
                        value:'service/SetWithdrawAddress',
                        label:'service/SetWithdrawAddress'
                    });
                    break;
				case "RespondService":
                    iServiceObj.children.push({
						value:'RespondService',
						label:'RespondService'
					});
					break;
				case "DisableServiceBinding":
                    iServiceObj.children.push({
						value:'DisableServiceBinding',
						label:'DisableServiceBinding'
					});
					break;
				case "EnableServiceBinding":
                    iServiceObj.children.push({
						value:'EnableServiceBinding',
						label:'EnableServiceBinding'
					});
					break;
				case "UpdateServiceBinding":
                    iServiceObj.children.push({
						value:'updateServiceBinding',
						label:'UpdateServiceBinding'
					});
					break;
				case "StartRequestContext":
                    iServiceObj.children.push({
						value:'StartRequestContext',
						label:'StartRequestContext'
					});
					break;
				case "KillRequestContext":
                    iServiceObj.children.push({
						value:'KillRequestContext',
						label:'KillRequestContext'
					});
					break;
				case "PauseRequestContext":
                    iServiceObj.children.push({
						value:'PauseRequestContext',
						label:'PauseRequestContext'
					});
					break;
				case "UpdateRequestContext":
                    iServiceObj.children.push({
						value:'UpdateRequestContext',
						label:'UpdateRequestContext'
					});
					break;
				case "RefundServiceDeposit":
                    iServiceObj.children.push({
						value:'RefundServiceDeposit',
						label:'RefundServiceDeposit'
					});
					break;
				case "WithdrawEarnedFees":
                    iServiceObj.children.push({
						value:'WithdrawEarnedFees',
						label:'WithdrawEarnedFees'
					});
					break;






			}
			
		});
		allTxType.push(tansferObj,delegationObj,validationObj,
            govObj,nftObj, oracleObj,coinswapObj,assetObj,iServiceObj,othersObj);
		return allTxType
	}
	static getRefUrlTxType(txType){
		let refUrlTxType;
		//下拉框回显项数据结构
		let UrlTxType = [
			['transfer','transfer'],
			['transfer','MultiSend'],
			['transfer','burn'],
			['transfer','setMemoRegexp'],
			['delegation','delegate'],
			['delegation','beginRedelegate'],
			['delegation','setWithdrawAddress'],
			['delegation','beginUnbonding'],
			['delegation','withdrawDelegatorReward'],
			['delegation','withdrawDelegatorRewardsAll'],
			['delegation','WithdrawValidatorCommission'],
			['validation','createValidator'],
			['validation','editValidator'],
			['validation','unjail'],
			['validation','withdrawValidatorRewardsAll'],
			['gov','submitProposal'],
			['gov','deposit'],
			['gov','vote'],
			['nft','IssueDenom'],
			['nft','NFTEdit'],
			['nft','NFTTransfer'],
			['nft','NFTMint'],
			['nft','NFTBurn'],
			['oracle','CreateFeed'],
			['oracle','StartFeed'],
			['oracle','PauseFeed'],
			['oracle','EditFeed'],
			['asset','IssueToken'],
			['asset','EditToken'],
			['asset','MintToken'],
			['asset','TransferTokenOwner'],
			['others','createGateway'],
			['others','editGateway'],
			['others','transferGatewayOwner'],
			['others','requestRand'],
			['others','addTrustee'],
			['others','deleteProfiler'],
			['others','deleteTrustee'],
			['others','verifyInvariant'],
			['others','claimHTLC'],
			['others','AddProfiler'],
			['others','createHTLC'],
			['others','refundHTLC'],


			['iService','DefineService'],
			['iService','BindService'],
			['iService','CallService'],
			['iService','RespondService'],
			['iService','DisableServiceBinding'],
			['iService','EnableServiceBinding'],
			['iService','UpdateServiceBinding'],
			['iService','StartRequestContext'],
			['iService','KillRequestContext'],
			['iService','PauseRequestContext'],
			['iService','UpdateRequestContext'],
			['iService','RefundServiceDeposit'],
			['iService','WithdrawEarnedFees'],
			['iService','service/SetWithdrawAddress'],




			['coinswap','addLiquidity'],
			['coinswap','removeLiquidity'],
			['coinswap','swapOrder'],
			['others','FundCommunityPool'],
		];
		UrlTxType.forEach( item => {
			if(Tools.onlyFirstWordUpperCase(item[item.length -1]) === txType){
				refUrlTxType = item
			}
		});
		return refUrlTxType
	}
}
