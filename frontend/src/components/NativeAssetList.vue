<template>
    <div class="native_asset_list_container">
        <div class="native_asset_list_header_content">
            <div class="native_asset_list_header_wrap">
                <h2 class="native_header_title_content">
                    Native Assets
                </h2>
            </div>
        </div>
        <div class="native_asset_list_content">
            <div class="native_asset_list_wrap">
                <m-asset-list-table :showNoData="showNoData" :items="nativeAssetList" name="nativeAssetList"></m-asset-list-table>
            </div>
            <div v-show="nativeAssetList.length === 0 && !showLoading">
                <img class="no_data_img" src="../assets/no_data.svg">
            </div>
        </div>
        <spin-component :show-loading="showLoading"></spin-component>
    </div>
</template>

<script>
	import SpinComponent from "./commonComponents/SpinComponent";
	import Service from "../service"
	import Tools from "../util/Tools"
	import MAssetListTable from "./table/MAssetListTable";
	export default {
		name: "NativeAssetList",
		components: {SpinComponent, MAssetListTable},
        data () {
			return {
				showNoData:false,
				nativeAssetList: [],
				showLoading: false,
            }
        },
        mounted(){
		    this.getNativeAssetList()
        },
        methods:{
			getNativeAssetList(){
				this.showLoading = true;
				Service.commonInterface({nativeAssetList:{}}, (res)=> {
					try {
						if(res){
							this.showLoading = false;
							this.nativeAssetList = res.data.map( item => {
								return {
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
							this.showLoading = false;
                        }
					}catch (e) {
						this.showLoading = false;
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
    .native_asset_list_container{
        width: 100%;
        position: relative;
        .no_data_img{
            position: absolute;
            left: 50%;
            top: 50%;
            transform: translate(-50%,-50%);
        }
        .native_asset_list_header_content{
            width: 100%;
            .native_asset_list_header_wrap{
                max-width: 12.8rem;
                margin: 0 auto;
                .native_header_title_content{
                    height: 0.7rem;
                    padding-left: 0.2rem;
                    font-size: 0.18rem;
                    color:var(--titleColor);
                    line-height: 0.7rem;
                }
            }
        }
        .native_asset_list_content{
            max-width: 12.8rem;
            margin: 0 auto 0.4rem auto;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            .native_asset_list_wrap{
                width: 100%;
                overflow-x: auto;
            }
        }
    }
</style>
