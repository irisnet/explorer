<template>
    <div class="gateway_asset_list_container">
        <page-title :title="pageTitle" :flShowPageLink="false"></page-title>
        <div class="gateway_asset_list_content">
            <div class="gateway_asset_list_wrap">
                <div>
                    <m-asset-list-table :showNoData="showNoData" :items="gatewayAssetList" name="gatewayAssetList"></m-asset-list-table>
                </div>
            </div>
            <div v-show="gatewayAssetList.length === 0 ">
                <img class="no_data_img" src="../../assets/no_data.svg">
            </div>
        </div>
    </div>
</template>

<script>
	import Service from "../../service"
	import Tools from "../../util/Tools"
	import MAssetListTable from "./MAssetListTable";
    import PageTitle from "../pageTitle/PageTitle";
	export default {
		name: "GatewayAssetList",
		components: {PageTitle, MAssetListTable},
        data () {
			return {
                pageTitle:'Gateway Assets List',
				gatewayAssetList:[],
				showNoData:false,
            }
        },
        mounted(){
			this.getGatewayAssetList()
        },
        methods:{
			getGatewayAssetList(){
				Service.commonInterface({gatewayAssetList:{}},(res) => {
                    try {
                    	if(res){
                    		this.gatewayAssetList = res.data.map(item => {
                    			return {
				                    Gateway: item.gateway,
				                    Symbol: item.symbol,
				                    Owner: item.owner,
				                    TotalSupply: this.formatNumber(item.total_supply),
				                    InitialSupply: this.formatNumber(item.initial_supply),
				                    MaxSupply: this.formatNumber(item.max_supply),
				                    Mintable: Tools.firstWordUpperCase(item.mintable),
                                    TokenID: item.token_id,
				                    flShowLink: true,
                                }
                            })
                        }else {
                        }
                    }catch (e) {
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
	        }
        }
	}
</script>

<style scoped lang="scss">
    .gateway_asset_list_container{
        width: 100%;
        position: relative;
        .no_data_img{
            position: absolute;
            left: 50%;
            top: 50%;
            transform: translate(-50%,-50%);
        }
        .gateway_asset_list_header_content{
            width: 100%;
            .gateway_asset_list_header_wrap{
                max-width: 12.8rem;
                margin: 0 auto;
                .gateway_header_title_content{
                    height: 0.7rem;
                    padding-left: 0.2rem;
                    font-size: 0.18rem;
                    color:#515a6e;
                    line-height: 0.7rem;
                }
            }
        }
        .gateway_asset_list_content{
            padding-top: 0.54rem;
            max-width: 12.8rem;
            margin: 0 auto 0.4rem auto;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            .gateway_asset_list_wrap{
                width: 100%;
                & > div{
                    overflow-x: auto;
                }
            }
        }
    }
    @media screen and (max-width: 910px){
        .gateway_asset_list_container{
            .gateway_asset_list_content{
                padding-top: 0;
                .gateway_asset_list_wrap{
                    & > div{
                        margin: 0 0.1rem;
                    }
                }
            }
        }
    }
</style>
