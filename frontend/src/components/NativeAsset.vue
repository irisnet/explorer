<template>
    <div class="native_asset_list_page_container">
        <div class="native_asset_list_title_wrap" :class="$store.state.isMobile ? 'mobile_asset_title' : ''">
            <div class="native_asset_list_title_content">
                <span class="native_asset_list_title">{{listTitleName}}</span>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="issueToken.length !== 0">
            <div style="padding: 0.2rem 0">Issue Token</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flIssueTokenShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="issueToken" name="issueToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination  :total-rows="issueTokenTotalPageNum" v-model="issueTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="editToken.length !== 0">
            <div style="padding: 0.2rem 0">Edit Token</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <!--<spin-component :showLoading="flEditTokenShowLoading"></spin-component>-->
                    <native-asset :showNoData="showNoData" :items="editToken" name="editToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination :total-rows="editTokenTotalPageNum" v-model="editTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="mintToken.length !== 0">
            <div style="padding: 0.2rem 0">Mint Token</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <!--<spin-component :showLoading="flMinTokenShowLoading"></spin-component>-->
                    <native-asset :showNoData="showNoData" :items="mintToken" name="mintToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination  :total-rows="mintTokenTotalPageNum" v-model="mintTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="transferToken.length !== 0">
            <div style="padding: 0.2rem 0">Transfer Owner</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <!--<spin-component :showLoading="flTransferTokenShowLoading"></spin-component>-->
                    <native-asset :showNoData="showNoData" :items="transferToken" name="transferToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination :total-rows="transferTokenTotalPageNum" v-model="transferTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div v-show="issueToken.length === 0 && editToken.length === 0
        && mintToken.length === 0 && transferToken.length === 0">
            <img class="no_data_img" src="../assets/no_data.svg">
        </div>
    </div>
</template>

<script>
    import NativeAsset from "./table/MNativeAssetTxListTable"
    import Service from "../service"
    import Tools from "../util/Tools"
    import SpinComponent from "./commonComponents/SpinComponent"
	export default {
		name: "NativeAssetPage",
		components: {NativeAsset, SpinComponent},
        data(){
			return {
				issueToken:[],
				editToken:[],
				mintToken:[],
				transferToken:[],
				issueTokenTotalPageNum: 0,
				editTokenTotalPageNum: 0,
				mintTokenTotalPageNum: 0,
				transferTokenTotalPageNum: 0,
				issueTokenCurrentPageNum:  1,
				editTokenCurrentPageNum:  1,
				mintTokenCurrentPageNum:  1,
				transferTokenCurrentPageNum:  1,
				pageSize: 10,
				showNoData: false,
                flIssueTokenShowLoading: false,
                flEditTokenShowLoading: false,
                flMinTokenShowLoading: false,
                flTransferTokenShowLoading: false,
				listTitleName:'Native Asset Txs',
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
                this.editTokenCurrentPageNum = editTokenCurrentPageNum
                this.getEditToken()
            },
            mintTokenCurrentPageNum(mintTokenCurrentPageNum){
                this.mintTokenCurrentPageNum = mintTokenCurrentPageNum
                this.getMintToken()
            },
            transferTokenCurrentPageNum(transferTokenCurrentPageNum){
                this.transferTokenCurrentPageNum = transferTokenCurrentPageNum
                this.getTransferToken()
            }
        },
        mounted(){
			this.getIssueToken();
			this.getEditToken();
			this.getMintToken();
			this.getTransferToken()
        },
        methods:{
	        linkGen(pageNum) {
		        return pageNum === 1 ? '?' : `?page=${pageNum}`
	        },
            getIssueToken(){
	            this.flIssueTokenShowLoading = true;
	            Service.commonInterface({nativeAsset:{
	            	pageNumber:this.issueTokenCurrentPageNum,
	            	pageSize:this.pageSize,
                    tokenType:this.issueTokenType
                    }}, (issueToken) => {
                    this.flIssueTokenShowLoading = false;
                    try {
                        if(issueToken.data){
                            this.issueTokenTotalPageNum = issueToken.data.total
	                        this.issueToken = issueToken.data.txs.map(item => {
		                       return {
			                       Owner: item.owner,
			                       Symbol: item.symbol,
                                   InitialSupply: item.initial_supply,
                                   MaxSupply: item.max_supply,
                                   Mintable: item.mintable,
                                   Decimal: item.decimal,
                                   SymbolMin: item.symbol_min,
                                   Name: item.name,
                                   Block: item.height,
                                   TxHash: item.tx_hash,
                                   TxFee: this.formatFee(item.tx_fee),
                                   TxStatus: item.tx_status,
                                   Timestamp: Tools.format2UTC(item.timestamp),
                               }
                            })
                        }
                    }catch (e) {
                        console.error(e)
                    }
                })
            },
	        getEditToken(){
                Service.commonInterface({nativeAsset: {
                	pageNumber: this.editTokenCurrentPageNum,
                    pageSize: this.pageSize,
                    tokenType: this.editTokenLType,
                    }}, (editToken) => {
                	try {
                        if(editToken.data){
                            this.editTokenTotalPageNum = editToken.data.total
	                        this.editToken = editToken.data.txs.map( item => {
                        		return {
                        			Token: item.token_id,
                                    Owner: item.owner,
                                    MaxSupply: item.max_supply,
                                    Mintable: item.mintable,
                                    Decimal: item.decimal,
                                    Symbolmin: item.symbol_min,
			                        Name: item.name,
			                        Block: item.height,
			                        TxHash: item.tx_hash,
			                        TxFee: this.formatFee(item.tx_fee),
			                        TxStatus: item.tx_status,
			                        Timestamp: Tools.format2UTC(item.timestamp),
                                }
                            })
                        }
	                }catch (e) {
                        console.error(e)
	                }
                })
            },
	        getMintToken(){
	        	Service.commonInterface({nativeAsset:{
	        		pageNumber: this.mintTokenCurrentPageNum,
                    pageSize: this.pageSize,
                    tokenType: this.mintTokenType
                    }},(mintToken) => {
	        		try {
	        			if(mintToken.data){
                            this.mintTokenTotalPageNum = mintToken.data.total
	        				this.mintToken = mintToken.data.txs.map( item => {
	        					return {
							        Token: item.token_id,
							        Owner: item.owner,
							        MintTo: item.mint_to,
                                    Amount: item.amount,
							        Block: item.height,
							        TxHash: item.tx_hash,
							        TxFee: this.formatFee(item.tx_fee),
							        TxStatus: item.tx_status,
							        Timestamp: Tools.format2UTC(item.timestamp),
                                }
                            })
                        }

			        }catch(e){
	        			console.error(e)
                    }
                })
            },
	        getTransferToken(){
	        	Service.commonInterface({nativeAsset:{
	        		pageNumber: this.transferTokenCurrentPageNum,
                    pageSize: this.pageSize,
                    tokenType: this.transferTokenType
                    }},(transferToken) => {
	        		try {
	        			if(transferToken.data){
	        			    this.transferTokenTotalPageNum = transferToken.data.total
	        				this.transferToken = transferToken.data.txs.map( item => {
	        					return {
							        Token: item.token_id,
                                    SrcOwner: item.src_owner,
                                    DstOwner: item.dst_owner,
							        Block: item.height,
							        TxHash: item.tx_hash,
							        TxFee: this.formatFee(item.tx_fee),
							        TxStatus: item.tx_status,
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
	        		return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(fee.amount[0].amount)),4)} ${Tools.formatDenom(fee.amount[0].denom).toUpperCase()}`;
                }
            }
        }
	}
</script>

<style scoped lang="scss">
    .native_asset_list_page_container{
        position: relative;
        .no_data_img{
            position: absolute;
            left: 50%;
            top: 50%;
            transform: translate(-50%,-50%);
        }
        .mobile_asset_title{
            position: static;
            .native_asset_list_title_content{
                .native_asset_list_title{
                    padding-left: 0.1rem;
                }
            }

        }
        .native_asset_list_title_wrap{
            width: 100%;
            position: fixed;
            z-index: 10;
            background-color: #ffffff;
            .native_asset_list_title_content{
                height:0.7rem;
                display: flex;
                align-items: center;
                max-width: 12.8rem;
                margin: 0 auto;
                background-color: #ffffff;
                .native_asset_list_title{
                    font-size: 0.18rem;
                    font-weight: 500;
                }
            }
        }
    }
    .native_asset_list_table_container{
        max-width: 12.8rem;
        padding-top: 0.7rem;
        margin: 0 auto;
        .native_asset_list_table_content{
            .table_list_content{
                width: 100%;
                overflow-x: auto;
                padding-top: 0rem;
                .no_data_show{
                    display: flex;
                    justify-content: center;
                    border-top:0.01rem solid #eee;
                    border-bottom:0.01rem solid #eee;
                    font-size:0.14rem;
                    height:3rem;
                    align-items: center;
                }
            }
            .native_asset_nav_footer_content{
                display: flex;
                justify-content: flex-end;
                height: 0.7rem;
                align-items: center;
                margin-bottom: 0.2rem;
            }

        }
    }
    .native_asset_list_table_container:nth-of-type(3){
        padding-top: 0;
    }
    .native_asset_list_table_container:nth-of-type(4){
        padding-top: 0;
    }
    .native_asset_list_table_container:nth-of-type(5){
        padding-top: 0;
    }
    @media screen and (max-width: 910px){
        .native_asset_list_page_container{
            .native_asset_list_title_wrap{
                position: static;
            }
        }
        .native_asset_list_table_container{
            padding-top: 0;
            padding-left: 0.1rem;
            padding-right: 0.1rem;
            .native_asset_list_table_content{
                .table_list_content{
                    padding-top: 0;
                }
            }
        }
    }
</style>
