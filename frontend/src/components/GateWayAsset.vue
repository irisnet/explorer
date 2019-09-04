<template>
    <div class="gateway_asset_list_page_container">
        <div class="gateway_asset_list_title_wrap" :class="$store.state.isMobile ? 'mobile_asset_title' : ''">
            <div class="gateway_asset_list_title_content">
                <span class="gateway_asset_list_title">{{listTitleName}}</span>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-if="gateWayIssueTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Issue Token Txs</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="gateWayIssueTokenList" name="gateWayIssueToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination :total-rows="gateWayIssueTokenTotalPageNum" v-model="gateWayIssueTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-if="gateWayEditTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Edit Token Txs</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <!--<spin-component :showLoading="flShowLoading"></spin-component>-->
                    <native-asset :showNoData="showNoData" :items="gateWayEditTokenList" name="editToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination  :total-rows="gateWayEditTokenTotalPageNum" v-model="gateWayEditTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-if="gateWayMintTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Mint Token Txs</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <!--<spin-component :showLoading="flShowLoading"></spin-component>-->
                    <native-asset :showNoData="showNoData" :items="gateWayMintTokenList" name="mintToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination  :total-rows="gateWayMintTokenTotalPageNum" v-model="gateWayMintTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-if="gateWayTransferOwnerTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Transfer Gateway Owner Txs</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <!--<spin-component :showLoading="flShowLoading"></spin-component>-->
                    <native-asset :showNoData="showNoData" :items="gateWayTransferOwnerTokenList" name="transferToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination  :total-rows="gateWayTransferTokenTotalPageNum" v-model="gateWayTransferTokenCurrentPageNum" :per-page="pageSize"></b-pagination>
                </div>
            </div>
        </div>
        <div v-show="gateWayIssueTokenList.length === 0 && gateWayEditTokenList.length === 0
                    && gateWayMintTokenList.length === 0 && gateWayTransferOwnerTokenList.length === 0">
            <img class="no_data_img" src="../assets/no_data.svg">
        </div>
    </div>
</template>

<script>
	import NativeAsset from "./table/MNativeAssetTxListTable"
	import SpinComponent from "./commonComponents/SpinComponent"
	import Service from "../service"
	import Tools from "../util/Tools"
	export default {
		name: "GateWayAsset",
		components: {NativeAsset, SpinComponent},
        data() {
            return {
	            gateWayIssueTokenList:[],
	            gateWayEditTokenList:[],
	            gateWayMintTokenList:[],
	            gateWayTransferOwnerTokenList:[],
	            // totalPageNum: sessionStorage.getItem("txpagenum") ? JSON.parse(sessionStorage.getItem("txpagenum")) : 1,
	            gateWayIssueTokenTotalPageNum: 0,
	            gateWayEditTokenTotalPageNum: 0,
	            gateWayMintTokenTotalPageNum: 0,
	            gateWayTransferTokenTotalPageNum: 0,
	            gateWayIssueTokenCurrentPageNum: 1,
	            gateWayEditTokenCurrentPageNum: 1,
	            gateWayMintTokenCurrentPageNum: 1,
	            gateWayTransferTokenCurrentPageNum: 1,
	            pageSize: 10,
	            showNoData: false,
	            flShowLoading: false,
	            listTitleName:'Gateway Asset Txs',
	            issueTokenType:'IssueToken',
	            editTokenLType:'EditToken',
	            mintTokenType:'MintToken',
	            transferTokenType:'TransferTokenOwner'
            }
        },
        watch:{
            gateWayIssueTokenCurrentPageNum(gateWayIssueTokenCurrentPageNum){
                this.gateWayIssueTokenCurrentPageNum = gateWayIssueTokenCurrentPageNum;
                this.getIssueToken()
            },
            gateWayEditTokenCurrentPageNum(gateWayEditTokenCurrentPageNum){
                this.gateWayEditTokenCurrentPageNum = gateWayEditTokenCurrentPageNum
                this.getEditToken()
            },
            gateWayMintTokenCurrentPageNum(gateWayMintTokenCurrentPageNum){
                this.gateWayMintTokenCurrentPageNum = gateWayMintTokenCurrentPageNum
                this.getMintToken()
            },
            gateWayTransferTokenCurrentPageNum(gateWayTransferTokenCurrentPageNum){
                this.gateWayTransferTokenCurrentPageNum = gateWayTransferTokenCurrentPageNum
                this.getTransferToken()
            }
        },
        mounted () {
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
		        Service.commonInterface({gatewayAsset:{
				        pageNumber:this.gateWayIssueTokenCurrentPageNum,
				        pageSize:this.pageSize,
				        tokenType:this.issueTokenType
			        }}, (issueToken) => {
			        this.flShowLoading = false;
			        try {
				        if(issueToken.data){
				            this.gateWayIssueTokenTotalPageNum = issueToken.data.total
					        sessionStorage.setItem('gateWayIssueTokenTotalPageNum',issueToken.data.total)
					        this.gateWayIssueTokenList = issueToken.data.txs.map(item => {
						        return {
							        Owner: item.owner,
							        Symbol: item.symbol,
							        InitialSupply: item.initial_supply,
							        Gateway: item.gateway,
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
				        }
			        }catch (e) {
				        console.error(e)
			        }
		        })
	        },
	        getEditToken(){
		        Service.commonInterface({gatewayAsset: {
				        pageNumber: this.gateWayEditTokenCurrentPageNum,
				        pageSize: this.pageSize,
				        tokenType: this.editTokenLType,
			        }}, (editToken) => {
			        try {
				        if(editToken.data){
                            this.gateWayEditTokenTotalPageNum = editToken.data.total
					        this.gateWayEditTokenList = editToken.data.txs.map( item => {
						        return {
							        Token: item.token_id,
							        Owner: item.owner,
							        Block: item.height,
							        TxHash: item.tx_hash,
							        TxFee: this.formatFee(item.tx_fee),
							        TxStatus: Tools.firstWordUpperCase(item.tx_status),
							        Timestamp: Tools.format2UTC(item.timestamp),
							        TokenID: item.token_id,
							        flShowLink: true,
						        }
					        })
				        }
			        }catch (e) {
				        console.error(e)
			        }
		        })
	        },
	        getMintToken(){
		        Service.commonInterface({gatewayAsset:{
				        pageNumber: this.gateWayMintTokenCurrentPageNum,
				        pageSize: this.pageSize,
				        tokenType: this.mintTokenType
			        }},(mintToken) => {
			        try {
				        if(mintToken.data){
                            this.gateWayMintTokenTotalPageNum = mintToken.data.total
					        this.gateWayMintTokenList = mintToken.data.txs.map( item => {
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
							        TokenID: item.token_id,
							        flShowLink: true,
						        }
					        })
				        }

			        }catch(e){
				        console.error(e)
			        }
		        })
	        },
	        getTransferToken(){
		        Service.commonInterface({gatewayAsset:{
				        pageNumber: this.gateWayTransferTokenCurrentPageNum,
				        pageSize: this.pageSize,
				        tokenType: this.transferTokenType
			        }},(transferToken) => {
			        try {
				        if(transferToken.data){
                            this.gateWayTransferTokenTotalPageNum = transferToken.data.total
					        this.gateWayTransferOwnerTokenList = transferToken.data.txs.map( item => {
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

<style scoped lang="scss">
    .gateway_asset_list_page_container{
        position: relative;
        .no_data_img{
            position: absolute;
            left: 50%;
            top: 50%;
            transform: translate(-50%,-50%);
        }
        .mobile_asset_title{
            position: static !important;
            .gateway_asset_list_title_content{
                .gateway_asset_list_title{
                    padding-left: 0.1rem;
                }
            }
        }
        .gateway_asset_list_title_wrap{
        width: 100%;
        position: fixed;
        z-index: 10;
        background-color: #ffffff;
            .gateway_asset_list_title_content{
                height:0.7rem;
                display: flex;
                align-items: center;
                max-width: 12.8rem;
                margin: 0 auto;
                background-color: #ffffff;
                .gateway_asset_list_title{
                    font-size: 0.18rem;
                    font-weight: 500;
                }
            }
        }
    }
    .gateway_asset_list_table_container{
        max-width: 12.8rem;
        padding-top: 0.7rem;
        margin: 0 auto;
    .gateway_asset_list_table_content{
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
    .pagination_nav_footer_content{
        display: flex;
        justify-content: flex-end;
        height: 0.7rem;
        align-items: center;
    }

    }
    }
    .gateway_asset_list_table_container:nth-of-type(3){
        padding-top: 0;
    }
    .gateway_asset_list_table_container:nth-of-type(4){
        padding-top: 0;
    }
    .gateway_asset_list_table_container:nth-of-type(5){
        padding-top: 0;
    }
    @media screen and (max-width: 910px){
        .gateway_asset_list_page_container{
            .transaction_list_title_wrap{
                position: static;
            }
        }
        .gateway_asset_list_table_container{
        padding-top: 0;
        padding-left: 0.1rem;
        padding-right: 0.1rem;
        .gateway_asset_list_table_content{
        .table_list_content{
        padding-top: 0;
    }
    }
    }
    }
</style>
