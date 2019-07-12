<template>
    <div class="gateway_asset_list_page_container">
        <div class="gateway_asset_list_title_wrap">
            <div class="gateway_asset_list_title_content">
                <span class="gateway_asset_list_title">{{listTitleName}}</span>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-show="gateWayIssueTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Issue Token</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="gateWayIssueTokenList" name="issueToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="gateWayIssueTokenTotalPageNum" v-model="gateWayIssueTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-show="gateWayEditTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Edit Token</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="gateWayEditTokenList" name="editToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="gateWayEditTokenTotalPageNum" v-model="gateWayEditTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-show="gateWayMintTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Mint Token</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="gateWayMintTokenList" name="mintToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="gateWayMintTokenTotalPageNum" v-model="gateWayMintTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
        </div>
        <div class="gateway_asset_list_table_container" v-show="gateWayTransferOwnerTokenList.length !== 0">
            <div style="padding: 0.2rem 0">Transfer Owner</div>
            <div class="gateway_asset_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <native-asset :showNoData="showNoData" :items="gateWayTransferOwnerTokenList" name="transferToken"></native-asset>
                    <!--<m-tx-list-page-table :showNoData="showNoData" :items="txList"></m-tx-list-page-table>-->
                    <div v-show="showNoData" class="no_data_show">
                        No Data
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <b-pagination-nav :link-gen="linkGen" :number-of-pages="gateWayTransferTokenTotalPageNum" v-model="gateWayTransferTokenCurrentPageNum" use-router></b-pagination-nav>
                </div>
            </div>
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
	            currentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
	            gateWayIssueTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
	            gateWayEditTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
	            gateWayMintTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
	            gateWayTransferTokenTotalPageNum: sessionStorage.getItem("issueTokenTotalPageNum") ? JSON.parse(sessionStorage.getItem("issueTokenTotalPageNum")) : 1,
	            gateWayIssueTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
	            gateWayEditTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
	            gateWayMintTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
	            gateWayTransferTokenCurrentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
	            pageSize: 30,
	            showNoData: false,
	            flShowLoading: false,
	            listTitleName:'GatewayAsset',
	            issueTokenType:'IssueToken',
	            editTokenLType:'EditToken',
	            mintTokenType:'MintToken',
	            transferTokenType:'TransferTokenOwner'
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
					        sessionStorage.setItem('issueTokenTotalPageNum',issueToken.data.total)
					        this.gateWayIssueTokenList = issueToken.data.txs.map(item => {
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
		        Service.commonInterface({gatewayAsset: {
				        pageNumber: this.gateWayEditTokenCurrentPageNum,
				        pageSize: this.pageSize,
				        tokenType: this.editTokenLType,
			        }}, (editToken) => {
			        try {
				        if(editToken.data){
					        sessionStorage.setItem('issueTokenTotalPageNum',editToken.data.total)
					        this.gateWayEditTokenList = editToken.data.txs.map( item => {
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
		        Service.commonInterface({gatewayAsset:{
				        pageNumber: this.gateWayMintTokenCurrentPageNum,
				        pageSize: this.pageSize,
				        tokenType: this.mintTokenType
			        }},(mintToken) => {
			        try {
				        if(mintToken.data){
					        sessionStorage.setItem('issueTokenTotalPageNum',mintToken.data.total)
					        this.gateWayMintTokenList = mintToken.data.txs.map( item => {
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
		        Service.commonInterface({gatewayAsset:{
				        pageNumber: this.gateWayTransferTokenCurrentPageNum,
				        pageSize: this.pageSize,
				        tokenType: this.transferTokenType
			        }},(transferToken) => {
			        try {
				        if(transferToken.data){
					        sessionStorage.setItem('issueTokenTotalPageNum',transferToken.data.total)
					        this.gateWayTransferOwnerTokenList = transferToken.data.txs.map( item => {
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
    .gateway_asset_list_page_container{
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
        margin-bottom: 0.2rem;
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