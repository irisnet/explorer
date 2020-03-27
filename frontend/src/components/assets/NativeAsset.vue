<template>
    <div class="native_asset_list_page_container">
        <page-title :title="listTitleName" :flShowPageLink="false"></page-title>
        <div class="native_asset_list_table_container" v-if="issueToken.length !== 0">
            <div style="padding: 0.2rem 0;color: #171d44">Issue Token Txs</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <native-asset :showNoData="showNoData" :items="issueToken" name="nativeIssueToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <m-pagination
                            :total="issueTokenTotalPageNum"
                            :page-size="pageSize"
                            :page="issueTokenCurrentPageNum"
                            :page-change="issueTokenPageChange"
                    ></m-pagination>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="editToken.length !== 0">
            <div style="padding: 0.2rem 0;color: #171d44">Edit Token Txs</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <native-asset :showNoData="showNoData" :items="editToken" name="editToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <m-pagination
                            :total="editTokenTotalPageNum"
                            :page-size="pageSize"
                            :page="editTokenCurrentPageNum"
                            :page-change="editTokenPageChange"
                    ></m-pagination>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="mintToken.length !== 0">
            <div style="padding: 0.2rem 0;color: #171d44">Mint Token Txs</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <native-asset :showNoData="showNoData" :items="mintToken" name="mintToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <m-pagination
                            :total="mintTokenTotalPageNum"
                            :page-size="pageSize"
                            :page="mintTokenCurrentPageNum"
                            :page-change="mintTokenPageChange"
                    ></m-pagination>
                </div>
            </div>
        </div>
        <div class="native_asset_list_table_container" v-if="transferToken.length !== 0">
            <div style="padding: 0.2rem 0;color: #171d44">Transfer Owner Txs</div>
            <div class="native_asset_list_table_content">
                <div class="table_list_content">
                    <native-asset :showNoData="showNoData" :items="transferToken" name="transferToken"></native-asset>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="native_asset_nav_footer_content">
                    <m-pagination
                            :total="transferTokenTotalPageNum"
                            :page-size="pageSize"
                            :page="transferTokenCurrentPageNum"
                            :page-change="transferTokenPageChange"
                    ></m-pagination>
                </div>
            </div>
        </div>
        <div v-show="issueToken.length === 0 && editToken.length === 0
        && mintToken.length === 0 && transferToken.length === 0 && flShowContent">
            <img class="no_data_img"  src="../../assets/no_data.svg">
        </div>
    </div>
</template>

<script>
    import NativeAsset from "./MNativeAssetTxListTable"
    import Service from "../../service"
    import Tools from "../../util/Tools"
    import MPagination from "../commontables/MPagination";
    import PageTitle from "../pageTitle/PageTitle";
    import pageTitleConfig from "../pageTitle/pageTitleConfig";
	export default {
		name: "NativeAssetPage",
		components: {PageTitle, MPagination, NativeAsset},
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
				flShowContent: false,
                flIssueTokenShowLoading: false,
                flEditTokenShowLoading: false,
                flMinTokenShowLoading: false,
                flTransferTokenShowLoading: false,
				listTitleName:pageTitleConfig.AssetNativeAssetTxs,
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
	        issueTokenPageChange(pageNum){
		        this.issueTokenCurrentPageNum = pageNum;
		        this.getIssueToken()
            },
	        editTokenPageChange(pageNum){
		        this.editTokenCurrentPageNum = pageNum;
		        this.getEditToken()
            },
	        mintTokenPageChange(pageNum){
		        this.mintTokenCurrentPageNum = pageNum;
		        this.getMintToken()
            },
	        transferTokenPageChange(pageNum){
		        this.transferTokenCurrentPageNum = pageNum
		        this.getTransferToken()
            },
            getIssueToken(){
	            this.flIssueTokenShowLoading = true;
	            this.flShowContent = false;
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
                                   InitialSupply: this.formatNumber(item.initial_supply),
                                   Mintable: Tools.firstWordUpperCase(item.mintable),
                                   Block: item.height,
                                   TxHash: item.tx_hash,
                                   TxFee: this.formatFee(item.tx_fee),
                                   TxStatus: Tools.firstWordUpperCase(item.tx_status),
                                   Timestamp: Tools.format2UTC(item.timestamp),
			                       TokenID: item.token_id,
			                       flShowLink: true,
                               }
                            })
                        }else {
	                        this.flShowContent = true;
                        }
                    }catch (e) {
	                    this.flShowContent = true;
                        console.error(e)
                    }
                })
            },
	        getEditToken(){
		        this.flShowContent = false;
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
			                        Block: item.height,
			                        TxHash: item.tx_hash,
			                        TxFee: this.formatFee(item.tx_fee),
			                        TxStatus:Tools.firstWordUpperCase(item.tx_status),
			                        Timestamp: Tools.format2UTC(item.timestamp),
			                        TokenID: item.token_id,
			                        flShowLink: true,
                                }
                            })
                        }else {
	                        this.flShowContent = true;
                        }
	                }catch (e) {
		                this.flShowContent = true;
                        console.error(e)
	                }
                })
            },
	        getMintToken(){
		        this.flShowContent = false;
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
                                    Amount: this.formatNumber(item.amount),
							        Block: item.height,
							        TxHash: item.tx_hash,
							        TxFee: this.formatFee(item.tx_fee),
							        TxStatus: Tools.firstWordUpperCase(item.tx_status),
							        Timestamp: Tools.format2UTC(item.timestamp),
							        TokenID: item.token_id,
							        flShowLink: true,
                                }
                            })
                        }else {
					        this.flShowContent = true;
                        }

			        }catch(e){
	        			console.error(e);
				        this.flShowContent = true;
                    }
                })
            },
	        getTransferToken(){
		        this.flShowContent = false;
	        	Service.commonInterface({nativeAsset:{
	        		pageNumber: this.transferTokenCurrentPageNum,
                    pageSize: this.pageSize,
                    tokenType: this.transferTokenType
                    }},(transferToken) => {
	        		try {
	        			if(transferToken.data){
	        			    this.transferTokenTotalPageNum = transferToken.data.total;
	        				this.transferToken = transferToken.data.txs.map( item => {
	        					return {
							        Token: item.token_id,
                                    SrcOwner: item.src_owner,
                                    DstOwner: item.dst_owner,
							        Block: item.height,
							        TxHash: item.tx_hash,
							        TxFee: this.formatFee(item.tx_fee),
							        TxStatus: Tools.firstWordUpperCase(item.tx_status),
							        Timestamp: Tools.format2UTC(item.timestamp),
							        TokenID: item.token_id,
							        flShowLink: true,
                                }
                            })
                        }else {
					        this.flShowContent = true;
                        }

			        }catch (e) {
				        this.flShowContent = true;
                        console.error(e)
			        }
                })
            },
	        formatNumber(value){
		        let million = 1000000;
		        if(value > million){
			        return `${value/million}M`
		        }else {
			        return value
		        }
	        },
            formatFee(fee){
	        	if(fee){
	        		return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(fee.amount)),4)} ${Tools.formatDenom(fee.denom).toUpperCase()}`;
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
                    color: #515a6e;
                }
            }

        }
        .native_asset_list_title_wrap{
            width: 100%;
            box-sizing: border-box;
            position: fixed;
            z-index: 10;
            background-color: #F5F7FD;
            .native_asset_list_title_content{
                height:0.7rem;
                display: flex;
                align-items: center;
                max-width: 12.8rem;
                margin: 0 auto;
                background-color: #F5F7FD;
                .native_asset_list_title{
                    font-size: 0.18rem;
                    color: #515a6e;
                    font-weight: bold;
                }
            }
        }
    }
    .native_asset_list_table_container{
        max-width: 12.8rem;
        padding-top: 0.54rem;
        margin: 0 auto;
        .native_asset_list_table_content{
            .table_list_content{
                width: 100%;
                box-sizing: border-box;
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
