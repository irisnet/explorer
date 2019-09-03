<template>
    <div class="asset_info_container">
        <div class="asset_info_container">
            <div class="asset_info_wrap">
                <h1 class="asset_title">{{headerTitle}}</h1>
            </div>
        </div>
        <div class="asset_info_list_container">
            <div class="asset_info_list_content">
                <div class="asset_info_kflower_contennt">
                    <div class="asset_info_kflower_title">
                        <span class="kflower_title">KFLOWER</span>
                        <span class="native_blue_style">NATIVE</span>
                        <span class="native_fungible_style">FUNGIBLE</span>
                    </div>
                    <div class="kflower_content">
                        <ul class="kflower_left_content">
                            <li class="kflower_list_item" v-for="item in leftInfoContentArray">
                                <span class="kflower_item_name">{{item.key}}</span>
                                <span class="kflower_item_value">
                                    <router-link v-if="item.address" :to="addressRoute(item.address)">{{item.address}}</router-link>
                                    <span v-if="item.value">{{item.value}}</span>
                                </span>
                            </li>
                        </ul>
                        <ul class="kflower_right_content">
                            <li class="kflower_list_item"  v-for="item in rightInfoContentArray">
                                <span class="kflower_item_name">{{item.key}}</span>
                                <span class="kflower_item_value">{{item.value}}</span>
                            </li>
                        </ul>
                    </div>
                    <div class="asset_list_token_content" v-if="assetTransferList.length > 0">
                        <h3 class="asset_list_token_title">AssetTransferTxs 100</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="assetTransferList" name="assetTransferTxs"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <b-pagination  :total-rows="issueTokenTotalPageNum" v-model="issueTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="issueTokenList.length > 0">
                        <h3 class="asset_list_token_title">Issue Token Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="issueTokenList" :name="nativeOrGatewayToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <b-pagination  :total-rows="issueTokenTotalPageNum" v-model="issueTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="editTokenList.length > 0">
                        <h3 class="asset_list_token_title">Edit Token Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="editTokenList" name="editToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <b-pagination :total-rows="editTokenTotalPageNum" v-model="editTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="mintTokenList.length > 0">
                        <h3 class="asset_list_token_title">Mint Token Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="mintTokenList" name="mintToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <b-pagination  :total-rows="mintTokenTotalPageNum" v-model="mintTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="transferTokenList.length > 0">
                        <h3 class="asset_list_token_title">Transfer Owner Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="transferTokenList" name="transferToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <b-pagination :total-rows="transferTokenTotalPageNum" v-model="transferTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<script>
    import Server from "../service"
    import Tools from "../util/Tools"
	import NativeAsset from "./table/MNativeAssetTxListTable"
	export default {
		name: "AssetInformation",
		components: {NativeAsset},
        data () {
			return {
				showNoData:false,
                headerTitle: "NativeAssetInfo",
                issueTokenList: [],
                editTokenList: [],
                mintTokenList: [],
                transferTokenList: [],
				assetTransferList: [],
                leftInfoContentArray:[
                    {
                    	key:'Owner:',
                        address:'faa1x292qss22x4rls6ygr7hhnp0et94vwwrchaklp'
                    },
	                {
		                key:'Total Supply:',
		                value:'2000'
	                },
	                {
		                key:'Initial Supply:',
		                value:'2000'
	                },
	                {
		                key:'Max Supply:',
		                value:'2000'
	                },
	                {
		                key:'Mintable:',
		                value:'2000'
	                }
                ],
                rightInfoContentArray:[
	                {
		                key:'Name:',
		                value:'KBoyflower Token'
	                },
	                {
		                key:'Canonical Symbol:',
		                value:'2000'
	                },
	                {
		                key:'Decimal:',
		                value:'2000'
	                },
	                {
		                key:'Min Unit Alias:',
		                value:'2000'
	                }
                ],
				nativeOrGatewayToken: 'nativeIssueToken',//nativeIssueToken or gateWayIssueToken
				issueTokenTotalPageNum: 0,
				editTokenTotalPageNum: 0,
				mintTokenTotalPageNum: 0,
				transferTokenTotalPageNum: 0,
				issueTokenCurrentPageNum: 1,
				editTokenCurrentPageNum: 1,
				mintTokenCurrentPageNum: 1,
				transferTokenCurrentPageNum: 1,
				pageSize: 10,
				issueTokenType:'IssueToken',
				editTokenLType:'EditToken',
				mintTokenType:'MintToken',
				transferTokenType:'TransferTokenOwner'
            }
        },
		watch:{
			issueTokenCurrentPageNum(issueTokenCurrentPageNum){
				this.issueTokenCurrentPageNum = issueTokenCurrentPageNum;
				this.getIssueToken()
			},
			editTokenCurrentPageNum(editTokenCurrentPageNum){
				this.editTokenCurrentPageNum = editTokenCurrentPageNum;
				this.getEditToken()
			},
			mintTokenCurrentPageNum(mintTokenCurrentPageNum){
				this.mintTokenCurrentPageNum = mintTokenCurrentPageNum;
				this.getMintToken()
			},
			transferTokenCurrentPageNum(transferTokenCurrentPageNum){
				this.transferTokenCurrentPageNum = transferTokenCurrentPageNum;
				this.getTransferToken()
			}
		},
        mounted(){
            this.getIssueToken();
            this.getEditToken();
            this.getMintToken();
            this.getMintToken();
            this.getAssetInfo();
        },
        methods:{
	        getAssetInfo () {
	            Server.commonInterface( {}, (info) => {
	            	try {
                        if(info){

                        }
		            }catch (e) {
                        console.error(e)
		            }
                })
            },
			getIssueToken () {
                Server.commonInterface( {nativeAsset:{
		                pageNumber:this.issueTokenCurrentPageNum,
		                pageSize:this.pageSize,
		                tokenType:this.issueTokenType
	                }},(issueTxs) => {
                	try {
		                if(issueTxs.data){
		                	console.log(issueTxs.data,">?>>>>>")
			                this.issueTokenTotalPageNum = issueTxs.data.total;
			                this.issueTokenList = issueTxs.data.txs.map(item => {
				                return {
					                Owner: item.owner,
					                Symbol: item.symbol,
					                InitialSupply: item.initial_supply,
					                Mintable: Tools.firstWordUpperCase(item.mintable),
					                Block: item.height,
					                TxHash: item.tx_hash,
					                TxFee: this.formatFee(item.tx_fee),
					                TxStatus: Tools.firstWordUpperCase(item.tx_status),
					                Timestamp: Tools.format2UTC(item.timestamp),
				                }
			                })
		                }
	                }catch (e) {
                        console.err(e)
	                }
                })
            },
            getEditToken () {
				Server.commonInterface( {nativeAsset: {
						pageNumber: this.editTokenCurrentPageNum,
						pageSize: this.pageSize,
						tokenType: this.editTokenLType,
					}},(editTxs) => {
					try {
						if(editTxs.data){
							this.editTokenTotalPageNum = editTxs.data.total
							this.editTokenList = editTxs.data.txs.map( item => {
								return {
									Token: item.token_id,
									Owner: item.owner,
									Block: item.height,
									TxHash: item.tx_hash,
									TxFee: this.formatFee(item.tx_fee),
									TxStatus:Tools.firstWordUpperCase(item.tx_status),
									Timestamp: Tools.format2UTC(item.timestamp),
								}
							})
						}
					}catch (e) {
                        console.error(e)
					}
                })
            },
            getMintToken () {
				Server.commonInterface( {nativeAsset:{
						pageNumber: this.mintTokenCurrentPageNum,
						pageSize: this.pageSize,
						tokenType: this.mintTokenType
					}}, (mintTxs) => {
					try {
						if(mintTxs.data){
							this.mintTokenTotalPageNum = mintTxs.data.total
							this.mintTokenList = mintTxs.data.txs.map( item => {
								return {
									Token: item.token_id,
									Owner: item.owner,
									MintTo: item.mint_to,
									Amount: item.amount,
									Block: item.height,
									TxHash: item.tx_hash,
									TxFee: this.formatFee(item.tx_fee),
									TxStatus: Tools.firstWordUpperCase(item.tx_status),
									Timestamp: Tools.format2UTC(item.timestamp),
								}
							})
						}
					}catch (e) {
                        console.error(e)
					}
                })
            },
            getTransferToken () {
				Server.commonInterface( {nativeAsset:{
						pageNumber: this.transferTokenCurrentPageNum,
						pageSize: this.pageSize,
						tokenType: this.transferTokenType
					}}, (transferTxs) => {
				    try {
					    if(transferTxs.data){
						    this.transferTokenTotalPageNum = transferTxs.data.total
						    this.transferTokenList = transferTxs.data.txs.map( item => {
							    return {
								    Token: item.token_id,
								    SrcOwner: item.src_owner,
								    DstOwner: item.dst_owner,
								    Block: item.height,
								    TxHash: item.tx_hash,
								    TxFee: this.formatFee(item.tx_fee),
								    TxStatus: Tools.firstWordUpperCase(item.tx_status),
								    Timestamp: Tools.format2UTC(item.timestamp),
							    }
						    })
					    }
				    }catch (e) {
                        console.error(e)
				    }
                })
            },
	        formatFee(fee){
		        if(fee){
			        return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(fee.amount)),4)} ${Tools.formatDenom(fee.denom).toUpperCase()}`;
		        }
	        }
        }
	}
</script>

<style lang="scss">
    .asset_info_container{
        width: 100%;
        .asset_info_container{
            max-width: 12.8rem;
            margin: 0 auto;
            .asset_info_wrap{
                padding-left: 0.2rem;
                height: 0.7rem;
                display: flex;
                align-items: center;
                .asset_title{
                    margin: 0;
                }
            }
        }
        .asset_info_list_container{
            width: 100%;
            margin-bottom: 0.2rem;
            .asset_info_list_content{
                max-width: 12.8rem;
                margin: 0 auto;
                .asset_info_kflower_contennt{
                    .asset_info_kflower_title{
                        margin: 0;
                        padding-left: 0.2rem;
                        .kflower_title{
                            font-size: 0.18rem;
                            color:var(--titleColor);
                        }
                        .native_blue_style{
                            font-size: 0.14rem;
                            padding: 0.03rem 0.14rem;
                            background: #0580d3;
                            border-radius: 0.1rem;
                            color: #fff;
                            margin-left: 0.1rem;
                        }
                        .native_fungible_style{
                            font-size: 0.14rem;
                            padding: 0.03rem 0.14rem;
                            background: #00C276;
                            border-radius: 0.1rem;
                            color: #fff;
                            margin-left: 0.1rem;
                        }
                    }
                    .kflower_content{
                        box-sizing: border-box;
                        padding: 0.2rem;
                        display: flex;
                        border: 0.01rem solid #e7e9eb;
                        margin-top: 0.1rem;
                        .kflower_left_content{
                            flex: 1;
                            margin-right: 0.2rem;
                            .kflower_list_item{
                                display: flex;
                                margin-bottom: 0.12rem;
                                .kflower_item_name{
                                    font-size: 0.14rem;
                                    width: 1.3rem;
                                    color: var(--contentColor);
                                }
                                .kflower_item_value{
                                    flex: 1;
                                    font-size: 0.14rem;
                                    color: var(--titleColor);
                                    a{
                                        color: var(--bgColor) !important;
                                    }
                                }
                            }
                        }
                        .kflower_right_content{
                            flex: 1;
                            .kflower_list_item{
                                display: flex;
                                margin-bottom: 0.12rem;
                                .kflower_item_name{
                                    width: 1.6rem;
                                    font-size: 0.14rem;
                                    color: var(--contentColor);
                                }
                                .kflower_item_value{
                                    flex: 1;
                                    font-size: 0.14rem;
                                    color: var(--titleColor);
                                }
                            }
                        }
                    }
                    .asset_transfer_txs_content{

                    }
                    .asset_list_token_content{
                        width: 100%;
                        .asset_list_token_title{
                            margin-left: 0.2rem;
                            font-size: 0.18rem;
                            color:var(--titleColor);
                            height: 0.6rem;
                            line-height: 0.6rem;
                        }
                        .asset_list_table_content{
                            padding-top: 0;
                            width: 100%;
                            overflow-x: auto;
                            -webkit-overflow-scrolling: touch;
                        }
                        .native_asset_nav_footer_content{
                            display: flex;
                            justify-content: flex-end;
                            height: 0.7rem;
                            align-items: center;
                        }
                    }
                    .asset_list_edit_token_content{
                        .asset_list_edit_token_title{
                            margin-left: 0.2rem;
                            font-size: 0.18rem;
                            color:var(--titleColor);
                            height: 0.6rem;
                            line-height: 0.6rem;
                        }
                    }
                    .asset_list_mint_token_content{
                        .asset_list_mint_token_title{
                            margin-left: 0.2rem;
                            font-size: 0.18rem;
                            color:var(--titleColor);
                            height: 0.6rem;
                            line-height: 0.6rem;
                        }
                    }
                    .asset_list_transfer_token_content{
                        .asset_list_transfer_token_title{
                            margin-left: 0.2rem;
                            font-size: 0.18rem;
                            color:var(--titleColor);
                            height: 0.6rem;
                            line-height: 0.6rem;
                        }
                    }
                }
            }
        }
    }
    @media screen and (max-width: 910px) {
        .asset_info_container{
            .asset_info_list_container{
                .asset_info_list_content{
                    .asset_info_kflower_contennt{
                        box-sizing: border-box;
                        padding:  0 0.1rem;
                        .asset_info_kflower_title{
                            padding-left: 0.1rem;
                        }
                        .kflower_content{
                            flex-direction: column;
                            .kflower_left_content{
                                .kflower_list_item{
                                    flex-direction: column;
                                    margin-bottom: 0;
                                }
                            }
                            .kflower_right_content{
                                .kflower_list_item{
                                    flex-direction: column;
                                    margin-bottom: 0;
                                }
                            }
                        }
                        .asset_list_token_content{
                            .asset_list_token_title{
                                margin-left: 0.1rem;
                            }
                        }
                    }
                }
            }
        }
    }
</style>
