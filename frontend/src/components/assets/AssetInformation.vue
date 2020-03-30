<template>
    <div class="asset_info_container">
        <div class="asset_info_list_container" v-show="flShowInformation">
            <page-title :flShowPageLink="false">
                <span style="color: #515a6e;font-weight: bold">{{tokenID}} {{contentDoc}}</span>
            </page-title>
            <div class="asset_info_list_content">
                <div class="asset_info_kflower_contennt">
                    <div class="asset_info_kflower_title">
                        <div>
                            <span class="native_blue_style" :class="assetType === 'NATIVE' ? 'blue_style' : 'yellow_style'">{{assetType}}</span>
                            <span class="native_fungible_style">{{family}}</span>
                        </div>
                    </div>
                    <div class="kflower_content" v-if="leftInfoContentArray.length > 0">
                        <ul class="kflower_left_content">
                            <li class="kflower_list_item" v-for="item in leftInfoContentArray">
                                <span class="kflower_item_name">{{item.key}}</span>
                                <span class="kflower_item_value">
                                    <router-link v-if="item.address" :to="addressRoute(item.address)">{{item.address}}</router-link>
                                    <span v-if="item.value">{{item.value}}</span>
                                </span>
                            </li>
                        </ul>
                        <ul class="kflower_right_content" v-show="!flShowGatewayInfo">
                            <li class="kflower_list_item"  v-for="item in rightInfoContentArray">
                                <span class="kflower_item_name">{{item.key}}</span>
                                <span class="kflower_item_value">{{item.value}}</span>
                            </li>
                        </ul>
                        <ul class="kflower_right_content" v-show="flShowGatewayInfo">
                            <li class="kflower_list_item"  v-for="item in gatewayRightInfoContentArray">
                                <span class="kflower_item_name">{{item.key}}</span>
                                <span class="kflower_item_value">{{item.value}}</span>
                            </li>
                        </ul>
                    </div>
                    <div class="gateway_content" v-show="flShowGatewayInfo">
                        <h3 class="gateway_title">
                            <img v-show="iconImg" :src="iconImg">Gateway Info</h3>
                        <ul class="gateway_left_content">
                            <li class="gateway_list_item" v-for="i in gatewayLeftContentArray">
                                <span class="gateway_item_name">{{i.key}}</span>
                                <span class="gateway_item_value">
                                    <router-link v-if="i.address" :to="addressRoute(i.address)">{{i.address}}</router-link>
                                    <a v-if="i.key ==='Identity:' && i.value !== '--'" :href="i.href" target="_blank">{{i.value}}</a>
                                    <span v-if="i.key ==='Identity:' && i.value === '--'">--</span>
                                    <span v-if="i.value && i.key !=='Identity:'" :class="i.key === 'Website:' && i.value !=='--' ? 'link_style' : '' " @click="openUrl(i.key,i.value)">{{i.value}}</span>
                                </span>
                            </li>
                        </ul>
                    </div>
                    <div class="asset_list_token_content" v-if="issueTokenList.length > 0">
                        <h3 class="asset_list_token_title">Issue Token Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="issueTokenList" :name="nativeOrGatewayToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <m-pagination :total="issueTokenTotalPageNum"
                                          :page-size="pageSize"
                                          :page="issueTokenCurrentPageNum"
                                          :page-change="issueTokenPageChange"></m-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="editTokenList.length > 0">
                        <h3 class="asset_list_token_title">Edit Token Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="editTokenList" name="editToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <m-pagination :total="editTokenTotalPageNum"
                                          :page-size="pageSize"
                                          :page="editTokenCurrentPageNum"
                                          :page-change="editTokenPageChange"></m-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="mintTokenList.length > 0">
                        <h3 class="asset_list_token_title">Mint Token Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="mintTokenList" name="mintToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <m-pagination :total="mintTokenTotalPageNum"
                                          :page-size="pageSize"
                                          :page="mintTokenCurrentPageNum"
                                          :page-change="mintTokenPageChange"></m-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="transferTokenList.length > 0">
                        <h3 class="asset_list_token_title">Transfer Owner Txs</h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="transferTokenList" name="transferToken"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <m-pagination :total="transferTokenTotalPageNum"
                                          :page-size="pageSize"
                                          :page="transferTokenCurrentPageNum"
                                          :page-change="transferTokenPageChange"></m-pagination>
                        </div>
                    </div>
                    <div class="asset_list_token_content" v-if="transferGatewayOwnerList.length > 0">
                        <h3 class="asset_list_token_title">Transfer Gateway Owner Txs </h3>
                        <div class="asset_list_table_content">
                            <native-asset :showNoData="showNoData" :items="transferGatewayOwnerList" name="transferGatewayOwnerTxs"></native-asset>
                        </div>
                        <div class="native_asset_nav_footer_content">
                            <m-pagination></m-pagination>
                            <m-pagination :total="transferGatewayTokenTotalPageNum"
                                          :page-size="pageSize"
                                          :page="transferGatewayTokenCurrentPageNum"
                                          :page-change="transferGatewayTokenPageChange"></m-pagination>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-show="flShowNoDataImg">
            <img class="no_data_img" src="../../assets/no_data.svg">
        </div>
    </div>
</template>

<script>
    import axios from 'axios'
    import Server from "../../service"
    import Tools from "../../util/Tools"
	import NativeAsset from "./MNativeAssetTxListTable"
    import MPagination from "../commontables/MPagination";
    import PageTitle from "../pageTitle/PageTitle";
	export default {
		name: "AssetInformation",
		components: {PageTitle, MPagination, NativeAsset},
        data () {
			return {
                contentDoc:'Detalils',
				showNoData:false,
                headerTitle: '',
				tokenID: '',
                assetType: '',
				source: '',
				family:'',
				symbol: '',
				gateway:'',
                transferOwnerTxsTitle:'',
                issueTokenList: [],
                editTokenList: [],
                mintTokenList: [],
                transferTokenList: [],
                transferGatewayOwnerList: [],
				assetTransferList: [],
                leftInfoContentArray:[
                    {
                    	id:'owner',
                    	key:'Owner:',
                        address:''
                    },
	                {
		                id:'total_supply',
		                key:'Total Supply:',
		                value:''
	                },
	                {
		                id:'initial_supply',
		                key:'Initial Supply:',
		                value:''
	                },
	                {
		                id:'max_supply',
		                key:'Max Supply:',
		                value:''
	                },
	                {
		                id:'mintable',
		                key:'Mintable:',
		                value:''
	                }
                ],
                rightInfoContentArray:[
	                {
	                	id:'name',
		                key:'Name:',
		                value:''
	                },
	                {
		                id:'decimal',
		                key:'Decimal:',
		                value:''
	                },
	                {
		                id:'min_unit_alias',
		                key:'Min Unit Alias:',
		                value:''
	                }
                ],
				gatewayRightInfoContentArray:[
					{
						id:'name',
						key:'Name:',
						value:''
					},
					{
                        id:'canonical_symbol',
                        key:'Canonical Symbol:',
                        value:''
                    },
					{
						id:'decimal',
						key:'Decimal:',
						value:''
					},
					{
						id:'min_unit_alias',
						key:'Min Unit Alias:',
						value:''
					}
				],
                gatewayLeftContentArray:[
	                {
		                id:'moniker',
		                key:'Moniker:',
		                value: ''
	                },
                    {
                    	id: 'owner',
                        key: 'Owner:',
                        address: ''
                    },
	                {
		                id:'identity',
		                key:'Identity:',
                        value: '',
                        href: ''
	                },
	                {
		                id:'website',
		                key:'Website:',
		                value: ''
	                },
	                {
		                id:'details',
		                key:'Details:',
		                value: ''
	                },
                ],
				nativeOrGatewayToken: 'nativeIssueToken',//nativeIssueToken or gateWayIssueToken
				issueTokenTotalPageNum: 0,
				editTokenTotalPageNum: 0,
				mintTokenTotalPageNum: 0,
				transferTokenTotalPageNum: 0,
				transferGatewayTokenTotalPageNum: 0,
				issueTokenCurrentPageNum: 1,
				editTokenCurrentPageNum: 1,
				mintTokenCurrentPageNum: 1,
				transferTokenCurrentPageNum: 1,
				transferGatewayTokenCurrentPageNum: 1,
				pageSize: 10,
				issueTokenType: 'IssueToken',
				editTokenLType: 'EditToken',
				mintTokenType: 'MintToken',
				transferTokenType: 'TransferTokenOwner',
                transferGatewayOwnerTxs: "TransferGatewayOwner",
                moniker: "",
                flShowGatewayInfo: false,
                iconImg:'',
				flShowNoDataImg: false,
                flShowInformation:false,
            }
        },
        mounted(){

            this.getAssetInfo();
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
                this.transferTokenCurrentPageNum = pageNum;
                this.getTransferToken()
            },
            transferGatewayTokenPageChange(pageNum){
                this.transferGatewayTokenTotalPageNum = pageNum;
                this.getTransferGatewayOwnerTxs()
            },
	        getAssetInfo () {
	        	let param;
	        	if(this.$route.params.assetType){
                    param = {
	                    assetTokenInfo:{
		                    tokenId : this.$route.params.assetType
                        }
                    }
                }else if(this.$route.params.assetType && this.$route.params.tokenName){
			        param = {
				        gatewayTokenInfo:{
					        moniker : this.$route.params.tokenName
				        }
			        }
                }
	            Server.commonInterface( param, (info) => {
	            	try {
                        if(info.data){
                        	this.flShowInformation = true;
                        	this.gateway = info.data.gateway;
	                        this.symbol = info.data.symbol;
	                        this.source = info.data.source;
	                        this.family = info.data.family.toLocaleUpperCase();
	                        this.leftInfoContentArray.forEach( item => {
	                        	if('address' in item ){
	                        		item.address = info.data[item.id]
                                }else {
	                        		if(item.id === 'total_supply' || item.id === 'initial_supply' || item.id === 'max_supply'){
				                        item.value = info.data[item.id] ? this.formatNumber(info.data[item.id]) : '--'
                                    }else if(item.id === 'mintable'){
				                        item.value = info.data[item.id] ? Tools.firstWordUpperCase(info.data[item.id]) : '--'
                                    } else {
				                        item.value = info.data[item.id] ? info.data[item.id] : '--'
			                        }
                                }
                            });
	                        this.getCommonRightContent(info);
                            this.getGatewayInfo(info);
                            this.getIssueToken();
                            this.getEditToken();
                            this.getMintToken();
                            if(info.data.source === 'gateway'){
                            	this.getTransferGatewayOwnerTxs()
                            }else {
	                            this.getTransferToken();
                            }
                        }else {
                        	this.flShowNoDataImg = true;
	                        this.flShowInformation = false;
                        }
		            }catch (e) {
                        console.error(e)
		            }
                })
            },
	        getCommonRightContent(info){
	        	this.transferOwnerTxsTitle= 'Transfer Owner Txs';
                this.flShowGatewayInfo = false;
                this.assetType = info.data.source.toLocaleUpperCase()
                this.headerTitle = 'NativeAssetInfo';
                this.tokenID = info.data.token_id;
                if(info.data.source === 'native'){
	                this.rightInfoContentArray.forEach( item => {
		                item.value = info.data[item.id] ? info.data[item.id] : '--'
	                })
                }else if(info.data.source === 'gateway') {
	                this.gatewayRightInfoContentArray.forEach( item => {
		                item.value = info.data[item.id] ? info.data[item.id] : '--'
	                })
                }

            },
	        getGatewayInfo(info){
	        	if(info.data.source === 'gateway'){
	        		this.iconImg = info.data.asset_gateway.icons ? info.data.asset_gateway.icons : '';
			        this.moniker = info.data.asset_gateway.moniker;
			        this.transferOwnerTxsTitle= 'Transfer Gateway Owner Txs';
	        		this.flShowGatewayInfo = true;
			        this.headerTitle = 'GatewayAssetInfo';
			        this.assetType = info.data.source.toLocaleUpperCase();
			        this.gatewayLeftContentArray.forEach( item => {
			        	if(item.id === 'owner'){
					        item.address = info.data[item.id] ? info.data[item.id] : '--'
                        }else if(item.id === 'identity'){
					        item.href = info.data.asset_gateway[item.id] ? this.getKeyBaseName(info.data.asset_gateway[item.id]) : '--';
                            item.value =  info.data.asset_gateway[item.id] ? info.data.asset_gateway[item.id] : '--'
                        } else {
					        item.value = info.data.asset_gateway[item.id] ? info.data.asset_gateway[item.id] : '--'
				        }
			        })
                }

            },
            handleParam(pageNumber,pageSize,tokenType,symbol,gateway){
	            let param;
	            if(this.source === 'gateway'){
		            param = {
			            gatewayAssetTxBySymbol:{
                            pageNumber:pageNumber,
                            pageSize:pageSize,
                            tokenType:tokenType,
                            symbol:symbol,
                            gateway:gateway
                        }
		            };

                }else if(this.source === 'native'){
		            param = {
		            	nativeAssetTxBySymbol:{
				            pageNumber:pageNumber,
				            pageSize:pageSize,
				            tokenType:tokenType,
				            symbol:symbol,
                        }
		            }
	            }
	            return param

            },
			getIssueToken () {
                Server.commonInterface(this.handleParam(
	                this.issueTokenCurrentPageNum,this.pageSize,this.issueTokenType,this.symbol,this.gateway
                ) ,(issueTxs) => {
                	try {
		                if(issueTxs.data){
			                this.issueTokenTotalPageNum = issueTxs.data.total;
			                this.issueTokenList = issueTxs.data.txs.map(item => {
			                	if(item.gateway){
					                return {
						                Owner: item.owner,
                                        Gateway: item.gateway,
						                Symbol: item.symbol,
						                InitialSupply: item.initial_supply,
						                Mintable: Tools.firstWordUpperCase(item.mintable),
						                Block: item.height,
						                TxHash: item.tx_hash,
						                TxFee: this.formatFee(item.tx_fee),
						                TxStatus: Tools.firstWordUpperCase(item.tx_status),
						                Timestamp: Tools.format2UTC(item.timestamp),
						                flShowLink: false,
					                }
                                }else {
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
						                flShowLink: false,
					                }
                                }

			                })
		                }
	                }catch (e) {
                        console.err(e)
	                }
                })
            },
            getEditToken () {
				Server.commonInterface( this.handleParam(
					this.editTokenCurrentPageNum,this.pageSize,this.editTokenLType,this.symbol,this.gateway
                ),(editTxs) => {
					try {
						if(editTxs.data){
							this.editTokenTotalPageNum = editTxs.data.total;
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
				Server.commonInterface( this.handleParam(
					this.mintTokenCurrentPageNum,this.pageSize,this.mintTokenType,this.symbol,this.gateway
                ), (mintTxs) => {
					try {
						if(mintTxs.data){
							this.mintTokenTotalPageNum = mintTxs.data.total;
							this.mintTokenList = mintTxs.data.txs.map( item => {
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
								}
							})
						}
					}catch (e) {
                        console.error(e)
					}
                })
            },
            getTransferToken () {
				Server.commonInterface( this.handleParam(
					this.transferTokenCurrentPageNum,this.pageSize,this.transferTokenType,this.symbol,this.gateway
                ), (transferTxs) => {
				    try {
					    if(transferTxs.data){
						    this.transferTokenTotalPageNum = transferTxs.data.total;
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
            getTransferGatewayOwnerTxs(){
                Server.commonInterface({transferGatewayOwnerTxs:{
		                pageNumber: this.pageNumber,
		                pageSize: this.pageSize,
		                tokenType: this.transferGatewayOwnerTxs,
		                moniker: this.moniker
                    }}, (res) => {
                	try {
                        if(res) {
	                        if(res.data){
		                        this.transferGatewayTokenTotalPageNum = res.data.total;
		                        this.transferGatewayOwnerList = res.data.txs.map( item => {
			                        return {
				                        Gateway: item.gateway,
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
	        },
	        formatNumber(value){
		        let million = 1000000;
		        if(value > million){
			        return `${value/million}M`
		        }else {
			        return value
		        }
	        },
	        getKeyBaseName(identity) {
		        let url = `https://keybase.io/_/api/1.0/user/lookup.json?fields=basics&key_suffix=${identity}`,that = this;
		        if (identity) {
			        axios.get(url).then(res => {
				        if (res.data.them && res.data.them[0].basics.username) {
					        that.gatewayLeftContentArray[2].href = `https://keybase.io/${res.data.them[0].basics.username}`;
				        }
			        });
		        }
	        },
	        openUrl(isUrl,url) {
	        	if(isUrl === 'Website:'){
			        url = url.trim();
			        if (url) {
				        if (!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)) {
					        url = `http://${url}`;
				        }
				        window.open(url);
			        }
                }
	        }
        }
	}
</script>

<style lang="scss">
    .asset_info_container{
        width: 100%;
        box-sizing: border-box;
        position: relative;
        .no_data_img{
            position: absolute;
            left: 50%;
            top: 50%;
            transform: translate(-50%,-50%);
        }
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
                    font-size: 0.22rem;
                }
            }
        }
        .asset_info_list_container{
            width: 100%;
            box-sizing: border-box;
            margin-bottom: 0.2rem;
            .asset_info_list_content{
                max-width: 12.8rem;
                margin: 0 auto;
                padding-top: 0.54rem;
                .asset_info_kflower_contennt{
                    padding-top: 0.2rem;
                    .asset_info_kflower_title{
                        margin: 0;
                        display: flex;
                        align-items: center;
                        .kflower_title{
                            font-size: 0.22rem;
                            color:var(--titleColor);
                        }
                        div{
                            .native_blue_style{
                                font-size: 0.14rem;
                                padding: 0.03rem 0.14rem;
                                border-radius: 0.1rem;
                                background: #E2F3FF;

                            }
                            .blue_style{
                                color: var(--bgColor);
                            }
                            .yellow_style{
                                color: #FF9500;
                            }
                            .native_fungible_style{
                                font-size: 0.14rem;
                                padding: 0.03rem 0.14rem;
                                background: #E2F3FF;
                                border-radius: 0.1rem;
                                color: #00C321;
                                margin-left: 0.1rem;
                            }
                        }

                    }
                    .kflower_content{
                        box-sizing: border-box;
                        padding: 0.2rem;
                        display: flex;
                        border: 0.01rem solid #e7e9eb;
                        margin-top: 0.1rem;
                        background: #fff;
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
                                    word-break: break-all;
                                    font-size: 0.14rem;
                                    color: var(--titleColor);
                                }
                            }
                        }
                    }
                    .gateway_content{
                        .gateway_title{
                            font-size: 0.18rem;
                            color: var(--titleCorlor);
                            margin-left: 0.2rem;
                            height: 0.6rem;
                            line-height: 0.6rem;
                            img{
                                width: 0.25rem;
                                height:0.25rem;
                                margin-right: 0.1rem;
                            }
                        }
                        .gateway_left_content{
                            box-sizing: border-box;
                            padding: 0.2rem;
                            border: 0.01rem solid #e7e9eb;
                            .gateway_list_item{
                                display: flex;
                                margin-bottom: 0.12rem;
                                .gateway_item_name{
                                    width: 1.3rem;
                                    font-size: 0.14rem;
                                    color: var(--contentColor);
                                }
                                .gateway_item_value{
                                    color: var(--titleColor);
                                    font-size: 0.14rem;
                                    flex: 1;
                                    word-break: break-all;
                                    a{
                                        color: var(--bgColor) !important;
                                        cursor: pointer;
                                    }
                                    .link_style{
                                        cursor: pointer;
                                        color: var(--bgColor)
                                    }
                                }
                            }
                        }
                    }
                    .asset_transfer_txs_content{

                    }
                    .asset_list_token_content{
                        width: 100%;
                        box-sizing: border-box;
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
                            box-sizing: border-box;
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
                    padding-top: 0.15rem;
                    .asset_info_kflower_contennt{
                        box-sizing: border-box;
                        padding:  0 0.1rem;
                        .asset_info_kflower_title{
                            flex-direction: column;
                            align-items: flex-start;
                            div{
                                .native_blue_style{
                                    margin-left: 0;
                                }
                            }
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
                        .gateway_content{
                            .gateway_left_content{
                                .gateway_list_item{
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
