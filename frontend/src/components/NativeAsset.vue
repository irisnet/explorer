<template>
    <div class="native_asset_list_page_container">
        <div class="native_asset_list_title_wrap">
            <div class="native_asset_list_title_content">
                <span class="native_asset_list_title">{{listTitleName}}</span>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-show="issueToken.length !== 0">
            <div style="padding: 0.2rem 0">Issue Token</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="issueToken" name="issueToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="issueTokenTotalPageNum" v-model="issueTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-show="editToken.length !== 0">
            <div style="padding: 0.2rem 0">Edit Token</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="editToken" name="editToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="editTokenTotalPageNum" v-model="editTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-show="mintToken.length !== 0">
            <div style="padding: 0.2rem 0">Mint Token</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="mintToken" name="mintToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="mintTokenTotalPageNum" v-model="mintTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-show="transferToken.length !== 0">
            <div style="padding: 0.2rem 0">Transfer Owner</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="transferToken" name="transferToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="transferTokenTotalPageNum" v-model="transferTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
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
				totalPageNum: 100,
				issueTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
				editTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
				mintTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
				transferTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
				issueTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
				editTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
				mintTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
				transferTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
				pageSize: 30,
				showNoData: false,
				flShowLoading: false,
				listTitleName:'NativeAsset',
				issueTokenType:'IssueToken',
                editTokenLType:'EditToken',
                mintTokenType:'MintToken',
                transferTokenType:'TransferTokenOwner'
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
	        	this.flShowLoading = true;
	            Service.commonInterface({nativeAsset:{
	            	pageNumber:this.issueTokenCurrentPageNum,
	            	pageSize:this.pageSize,
                    tokenType:this.issueTokenType
                    }}, (issueToken) => {
		            this.flShowLoading = false;
                    try {
                        if(issueToken.data){
	                        sessionStorage.setItem('issueTokenTotalPageNum',issueToken.data.total)
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
	                        sessionStorage.setItem('issueTokenTotalPageNum',editToken.data.total)
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
					        sessionStorage.setItem('issueTokenTotalPageNum',mintToken.data.total)
	        				this.mintToken = mintToken.data.txs.map( item => {
	        					return {
							        Token: item.token_id,
							        Owner: item.owner,
							        MintTo: item.min_to,
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
					        sessionStorage.setItem('issueTokenTotalPageNum',transferToken.data.total)
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