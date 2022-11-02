<template>
	<div class="goz_container">
		<div class="goz_content_wrap">
			<div class="goz_content_left_content">
				<div class="goz_content_logo_content">
					<div><img src="../../assets/IOBSCAN_logo.png" alt=""></div>
				</div>
				<div class="goz_content_network_state" :class="changeRouter ? 'active_border_bottom' :'unActive_border_bottom'">
					<span><router-link :to="`/`">Network State Visualizer </router-link></span>
				</div>
				<div class="menu_content" @click="flShowMobileMenu">
					<img src="../../assets/menu.png" alt="">
				</div>
			</div>
			<div class="goz_content_right_content">
				<!--<div class="goz_content_link_content">
					<a href="https://map-of-zones-web.herokuapp.com/?period=24" target="_blank">Map of zones</a>
				</div>-->
				<!--<div class="goz_content_link_content">
					<a href="https://goz.cosmos.network/3d" target="_blank">Visualizer from Cosmos Team</a>
				</div>-->
				<div class="goz_content_link_content">
					<span>Network</span>
				</div>
				<div class="goz_network_links_content">
					<div class="network_container" @mouseenter="showNetWorkLogo()" @mouseleave="hideNetWorkLogo()">
                            <span style="color: #fff">
                                <i style="font-size: 0.24rem;padding-right: 0.02rem;" :class="currentNetworkClass"></i>
                                <i style="font-size: 0.08rem" class="iconfont icon-zhankai"></i>
                            </span>
						<ul class="network_list_container"
						    v-show="flShowNetworkLogo"
						    @mouseenter="showNetWorkLogo()"
						    @mouseleave="hideNetWorkLogo()">
							<li class="network_list_item"
							    v-for="item in netWorkArray"
							    @click="windowOpenUrl(item.host)"><i :class="item.icon"></i>{{item.netWorkSelectOption}}</li>
						</ul>
					</div>
				</div>
			</div>
			<ul class="goz_mobile_content" v-show="isShowMobileMenu">
				<li class="goz_mobile_link_item">
					<router-link :to="`/`">Network State Visualizer</router-link>
				</li>
				<!--<li class="goz_mobile_link_item">
					<router-link :to="`/download-rainbowgoz`">IBC-GoZ  Wallet</router-link>
				</li>-->
			<!--	<li class="goz_mobile_link_item">
					<a href="https://goz.cosmosnetwork.dev/2020/04/14/a-whole-new-world-testnets-in-the-ibc-era/" target="_blank">GoZ Website</a>
				</li>-->
			<!--	<li class="goz_mobile_link_item">
					<a href="https://map-of-zones-web.herokuapp.com/?period=24" target="_blank">Map of zones</a>
				</li>-->
				<!--<li class="goz_mobile_link_item">
					<a href="https://goz.cosmos.network/3d" target="_blank">Visualizer from Cosmos Team</a>
				</li>-->
				<li class="goz_mobile_net_work_content" @click="showNetworkToggle">
					<div class="goz_mobile_net_work_wrap">
						<span class="net_work_name">Network</span>
						<span class="iconfont icon-zhankai" :class="flShowNetwork ? 'defalut_rotate_style' : 'rotate_style'"></span>
					</div>
				</li>
				<li class="network_list_content network_up_style"
				    v-show="flShowNetwork"
				    >
					<div class="network_list_container_content_wrap">
						<div class="network_list_content_wrap" v-for="item in netWorkArray" @click="windowOpenUrl(item.host)">
							{{item.netWorkSelectOption}}<i :class="item.icon"></i>
						</div>
					</div>
				</li>
			</ul>
		</div>
	</div>
</template>

<script>
	import Tools from '../../util/Tools';
	import Service from "../../service";
	import constant from '../../constant/Constant';
	import lang from "../../lang";
	import skinStyle from '../../skinStyle'
	export default {
		name: "GozHeader",
		data () {
			return {
				netWorkArray: [],
				flShowNetworkLogo: false,
				currentNetworkClass:'',
				isShowMobileMenu: false,
				flShowNetwork:false,
				networkList:[],
			}
		},
		mounted(){
			this.networkList = require('../../../config/network.json')
			this.getConfig();
			
		},
		computed:{
			changeRouter(){
				if(this.$route.path === '/'){
					return  true
				}else {
					return  false
				}
			}
		},
		methods:{
			showNetworkToggle(){
				this.flShowNetwork = !this.flShowNetwork
			},
			flShowMobileMenu(){
				this.isShowMobileMenu = !this.isShowMobileMenu
			},
			showNetWorkLogo(){
				this.flShowNetworkLogo = true;
			},
			hideNetWorkLogo(){
				this.flShowNetworkLogo = false;
			},
			windowOpenUrl (url) {
				window.open(url)
			},
			getConfig () {
				// Service.commonInterface({headerConfig:{}},(res) => {
					try {
						// TODO 测试展示
						// res.cur_env = 'testnet';
						// res.chain_id = 'goz';
						this.handleConfigs(this.networkList);
						this.setNetWorkLogo(this.networkList);
					}catch (e) {
						console.error(e);
					}
				// });
			},
			handleConfigs (configs) {
				
				this.netWorkArray = configs.map(item => {
					if(item.network_name === constant.NET_WORK_NAME.IRISHUB){
						item.icon = 'iconfont icon-irisnet'
					}else if(item.network_name === constant.NET_WORK_NAME.NYANCAT){
						item.icon = 'iconfont icon-caihongmao'
					}else if(item.network_name === constant.NET_WORK_NAME.GOZTESTNET){
						item.icon = 'iconfont iconGOZ'
					}else if(item.network_name === constant.NET_WORK_NAME.COSMOSHUB){
						item.icon = 'iconfont icon-cosmos'
					}else if(item.network_name === constant.NET_WORK_NAME.STARGATE){
						item.icon = 'iconfont icon-stargate'
					}
					item.netWorkSelectOption = `${item.chain_id.toLocaleUpperCase()}`;
					return item
				});
				this.netWorkArray = this.netWorkArray.filter(item => {
					return item.env !== constant.ENVCONFIG.DEV && item.env !== constant.ENVCONFIG.QA && item.env !== constant.ENVCONFIG.STAGE
				});
			},
			setNetWorkLogo (currentEnv) {
				let networkName = '';
				if(currentEnv.configs){
					currentEnv.configs.forEach(item => {
						if(currentEnv.cur_env === item.env && currentEnv.chain_id === item.chain_id){
							networkName = item.network_name;
							sessionStorage.setItem('UMengID',item.umeng_id)
						}
					})
				}
				if (networkName === constant.NET_WORK_NAME.IRISHUB) {
					this.currentNetworkClass = 'iconfont icon-irisnet'
				} else if (networkName === constant.NET_WORK_NAME.FUXI) {
					this.currentNetworkClass = 'iconfont iconfuxi'
				} else if (networkName === constant.NET_WORK_NAME.NYANCAT) {
					this.currentNetworkClass = 'iconfont icon-caihongmao'
				} else if (networkName === constant.NET_WORK_NAME.GOZTESTNET) {
					this.currentNetworkClass = 'iconfont iconGOZ'
				} else if(networkName === constant.NET_WORK_NAME.STARGATE){
					this.currentNetworkClass = 'iconfont icon-stargate'
				}else {
					this.currentNetworkClass = 'iconfont icon-irisnet';
				}
			},
		}
	}
</script>

<style scoped lang="scss">
	.goz_container{
		width: 100%;
		background: rgba(16, 19, 55, 1);
		.goz_content_wrap{
			box-sizing: border-box;
			padding: 0 0.4rem;
			margin: 0 auto;
			display: flex;
			justify-content: space-between;
			align-items: center;
			.goz_content_left_content{
				display: flex;
				color: #fff;
				align-items: center;
				.menu_content{
					display: none;
				}
				.goz_content_logo_content{
					width: 1.5rem;
					padding: 0.1rem 0;
					display: flex;
					align-items: center;
					.iconfont{
						display: none;
					}
					div{
						width: 1.5rem;
						box-sizing: border-box;
						margin-top: 0.1rem;
						img {
							width: 1.48rem;
							height: 0.3rem;
						}
					}
				}
				.goz_content_network_state{
					margin-left: 0.4rem;
					span{
						a{
							color: #fff !important;
							font-size: 0.14rem;
							line-height: 0.16rem;
							box-sizing: border-box;
							padding-bottom: 0.15rem;
						}
					}
				}
				.active_border_bottom{
					span{
						a{
							border-bottom: 0.02rem solid #fff;
						}
					}
				}
				.unActive_border_bottom{
					span{
						a{
							border-bottom: 0.02rem solid transparent;
						}
					}
				}
			}
			.goz_content_right_content{
				display: flex;
				align-items: center;
				.goz_content_link_content{
					color:#fff;
					a{
						color:rgba(255,255,255,0.75) !important;
						font-size: 0.14rem;
						line-height: 0.16rem;
						padding: 0 0 0 0.2rem;
						
					}
					a:last-child{
						border-right: none;
					}
				}
				.goz_content_link_content:nth-of-type(3){
					border-right: none !important;
					a{
						border-right: none!important;
					}
				}
				.goz_network_links_content{
					margin-right: 0.14rem;
					.network_container{
						position: relative;
						height:0.6rem;
						line-height: 0.6rem;
						margin-left: 0.2rem;
						cursor: pointer;
						.network_list_container{
							background: rgba(16, 15, 50, 1);
							box-shadow: 0 0.02rem 0.1rem 0 rgba(3,16,114,0.15);
							width: auto;
							position: absolute;
							right: 0;
							top: 0.6rem;
							z-index: 2;
							text-align: right;
							.network_list_item{
								height: 0.4rem;
								line-height: 0.4rem;
								white-space: nowrap;
								padding: 0 0.2rem;
								cursor: pointer;
								font-size: 0.14rem;
								display: flex;
								color: rgba(255,255,255,1);
								&:hover{
									background: rgba(41,39,90,1);
								}
								i{
									font-size: 0.18rem;
									color: var(--titleColor);
									padding-right: 0.2rem;
								}
							}
							.network_list_item:last-child{
								padding-bottom: 0.05rem;
							}
						}
					}
					.iconzhankai{
						font-size: 0.32rem;
					}
				}
			}
			.goz_mobile_content{
				display: none;
			}
		}
	}
	@media (max-width: 1330px) {
		.goz_container{
			.goz_content_wrap{
				flex-direction: column;
				width: 100%;
				position: relative;
				.goz_content_left_content{
					width: 100%;
					justify-content: space-between;
					.menu_content{
						cursor: pointer;
						display: block;
						width: 0.25rem;
						height: 0.25rem;
						top: 0.26rem;
						margin-right: 0.1rem;
						img{
							width: 100%;
						}
					}
					.goz_content_logo_content{
						.iconfont{
							display: block;
							font-size: 0.22rem;
						}
					}
					.goz_content_network_state{
						margin-right: 0.1rem;
						display: none;
					}
				}
				.goz_content_right_content{
					width: 100%;
					justify-content: center;
					display: none;
					.goz_content_link_content{
						width: 100%;
						text-align: center;
						padding: 0.12rem 0;
						margin: 0 0.2rem;
						border-top: 0.01rem solid rgba(255,255,255,0.19);
						font-size: 0.12rem;
					}
					.goz_network_links_content{
						display: none;
					}
				}
				.goz_mobile_content{
					display: block;
					width: 100%;
					box-sizing: border-box;
					position: absolute;
					top: 0.57rem;
					background:#2D325A;
					z-index: 10;
					.goz_mobile_link_item{
						background: rgba(255,255,255,0.1);
						padding: 0.05rem 0.2rem;
						font-size: 0.14rem;
						a{
							color: #fff !important;
						}
					}
					.goz_mobile_net_work_content{
						box-sizing: border-box;
						padding: 0.05rem 0.2rem;
						background: rgba(255,255,255,0.1);
						.goz_mobile_net_work_wrap{
							display: flex;
							justify-content: space-between;
							span{
								font-size: 0.14rem;
								color: #fff;
							}
							.defalut_rotate_style{
								transform: rotate(-180deg);
							}
							.rotate_style{
								transform: rotate(0deg);
							}
						}
					}
					.network_list_item{
						background: rgba(255,255,255,0.1);
						padding: 0.05rem 0.15rem;
						color: #fff ;
						font-size: 0.14rem;
						text-align: right;
						display: flex;
						justify-content: space-between;
					}
					.network_list_content{
						.network_list_container_content_wrap{
							padding-top: 0.1rem;
							background: rgba(0,0,0,0.05);
							.network_list_content_wrap{
								padding: 0.05rem 0.15rem;
								color: #fff ;
								font-size: 0.14rem;
								text-align: right;
								display: flex;
								justify-content: space-between;
							}
							.network_up_style{
								background-color: rgba(0,0,0,0.05) !important;
							}
						}
						
					}
				}
			}
		}
	}
	@media (max-width: 768px) {
		.goz_container{
			.goz_content_wrap{
				flex-direction: column;
				width: 100%;
				position: relative;
				padding: 0;
				.goz_content_left_content{
					width: 100%;
					justify-content: space-between;
					box-sizing: border-box;
					padding: 0 0.2rem;
					.menu_content{
						display: block;
						width: 0.25rem;
						height: 0.25rem;
						top: 0.26rem;
						margin-right: 0.1rem;
						img{
							width: 100%;
						}
					}
					.goz_content_logo_content{
						.iconfont{
							display: block;
							font-size: 0.22rem;
						}
					}
					.goz_content_network_state{
						margin-right: 0.1rem;
						display: none;
					}
				}
				.goz_content_right_content{
					width: 100%;
					justify-content: center;
					display: none;
					.goz_content_link_content{
						width: 100%;
						text-align: center;
						padding: 0.12rem 0;
						margin: 0 0.2rem;
						border-top: 0.01rem solid rgba(255,255,255,0.19);
						font-size: 0.12rem;
					}
					.goz_network_links_content{
						display: none;
					}
				}
				.goz_mobile_content{
					display: block;
					width: 100%;
					box-sizing: border-box;
					position: absolute;
					top: 0.57rem;
					background:#2D325A;
					z-index: 10;
					.goz_mobile_link_item{
						background: rgba(255,255,255,0.1);
						padding: 0.05rem 0.2rem;
						font-size: 0.14rem;
						a{
							color: #fff !important;
						}
					}
					.network_list_item{
						background: rgba(255,255,255,0.1);
						padding: 0.05rem 0.15rem;
						color: #fff ;
						font-size: 0.14rem;
						text-align: right;
						display: flex;
						justify-content: space-between;
					}
				}
			}
		}
	}
	@media (max-width: 450px) {
		.goz_container{
			.goz_content_wrap{
				.goz_content_right_content{
					.goz_content_link_content{
						margin: 0 0.1rem;
						display: flex;
						justify-content: space-between;
						a{
							font-size: 0.12rem;
							border-right: none;
							padding: 0;
						}
					}
				}
			}
		}
	}
</style>
