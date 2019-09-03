<template>
    <div class="gateway_asset_list_container">
        <div class="gateway_asset_list_header_content">
            <div class="gateway_asset_list_header_wrap">
                <h2 class="gateway_header_title_content">
                    GatewayAssets
                </h2>
            </div>
        </div>
        <div class="gateway_asset_list_content">
            <div class="gateway_asset_list_wrap">
                <m-asset-list-table :showNoData="showNoData" :items="gatewayAssetList" name="gatewayAssetList"></m-asset-list-table>
            </div>
        </div>
        <div v-show="gatewayAssetList.length === 0">
            <img class="no_data_img" src="../assets/no_data.svg">
        </div>
    </div>
</template>

<script>
	import Service from "../service"
	import Tools from "../util/Tools"
	import MAssetListTable from "./table/MAssetListTable";
	export default {
		name: "GatewayAssetList",
		components: {MAssetListTable},
        data () {
			return {
				gatewayAssetList:[],
				showNoData:false
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
				                    Symbol: item.symbol,
				                    Owner: item.owner,
				                    TotalSupply: item.total_supply,
				                    InitialSupply: item.initial_supply,
				                    MaxSupply: item.max_supply,
				                    Mintable: Tools.firstWordUpperCase(item.mintable)
                                }
                            })
                        }
                    }catch (e) {
                        console.error(e)
                    }
                })
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
                    color:var(--titleColor);
                    line-height: 0.7rem;
                }
            }
        }
        .gateway_asset_list_content{
            max-width: 12.8rem;
            margin: 0 auto;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            .gateway_asset_list_wrap{
                width: 100%;
                overflow-x: auto;
            }
        }
    }
</style>
