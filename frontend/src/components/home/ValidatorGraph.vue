<template>
	<div class="validator_graph_container" :class="switchValue ? 'start_sky_img' : 'default_bg_img'">
		<div class="graph_content_wrap">
			<div class="graph_content_title">
				<div>Interchain Network
					<span class="beat_content" v-if="!flShowTestTooltip">Beta </span>
				</div>
				<div class="tooltip" v-if="flShowNetwork">
					<div class="graph_tooltip"><p><span></span><span>Connection Opened</span></p>
						<p><span></span><span>Connection Unopened</span></p></div>
					<!--<div class="background_toggle_content">
						 <el-switch v-model="switchValue"
									active-color="rgba(55, 124, 248, 1)"
									inactive-color="rgba(16, 19, 55, 1)"
									@change="switchBgColor(switchValue)"></el-switch>
						<span style="color: rgba(115, 122, 174, 1);font-size: 0.14rem;margin-left:0.1rem;line-height: 0.16rem">Starry-sky Mode</span>
					</div>-->
				</div>
			</div>
			<div class="graph_charts_container" :class="flShowNetwork ? '' : 'show_error_content'">
				<div class="graph_content_container" v-if="flShowNetwork">
					<span class="iconfont iconfuwei" v-if="flShowRevertIcon" @click="revertGraph()"></span>
					<div class="graph_container_content_wrap"
					     :style="{background: switchValue ? 'transparent' : '#2d325a'}">
						<!--						<p class="graph_charts_title" v-if="flShowNetwork && !flShowTestTooltip"></p>-->
						<div id="validator_graph_content" ref="chart_content"></div>
					</div>
					<div class="graph_list_container"
					     id="graph_list_container"
					     :style="{background: switchValue ? '' : '#2d325a'}"
					     :class="[flShowDefaultPickUp ? 'hide_pick_up_style' : '', showDefaultAnimation ? 'show_pick_up_style' : '' ]"
					     v-show="colorDataArray.length > 0">
						<div class="sort_content">
							<span :class="flConnectionActiveStyle ? 'active_style' : ''"
							      @click="sortByConnection(sortByConnectionSwitchIcon)">Connections <i class="iconfont"
							                                                                    v-show="flShowConnectionSortIcon"
							                                                                    :class="sortByConnectionSwitchIcon ? 'iconxiangxia' : 'iconxiangshang'"></i></span>
							<span :class="flLetterActiveStyle ? 'active_style' : ''"
							      @click="sortByLetter(sortByLetterSwitchIcon)">Chains <i class="iconfont"
							                                                           v-show="flShowLetterSortIcon"
							                                                           :class="sortByLetterSwitchIcon ? 'iconxiangxia' : 'iconxiangshang'"></i></span>
						</div>
						<div class="sort_scroll_content" ref="graph_list_content">
							<div class="graph_list_item_all" @click="selectAll()">
								<div class="legend_all_block">
									<img v-show="flAllCheckout" src="../../assets/select_all.svg" alt="">
									<img v-show="!flAllCheckout" src="../../assets/unselect_all.svg" alt=""></div>
								<span class="legend_name" :class=" flAllCheckout ? 'hide_style' : ''">All</span>
							</div>
							<ul class="graph_list_content">
								<li class="graph_list_item" v-for="(item,index) in colorDataArray"
								    @click="changeChart(item,index)">
									<div class="legend_block"
									     :style="{background: item.color}"
									     :class="item.isActive ? '' : 'hide_style'"
									     v-if="item.name !== 'Cosmos Hub (cosmoshub-4)'
									      && item.name !== 'IRIS Hub (irishub-1)'
									      && item.name !== 'Akash (Akashnet-2)'
									      && item.name !== 'Crypto.com (Crypto-Org-Chain-Mainnet-1)'">
									</div>
									<div class="legend_block_img" :class="item.isActive ? '' : 'img_hide_style'" v-if=" item.name === 'IRIS Hub (irishub-1)'">
										<img src="../../assets/IRISnet_LOGO_Small.png" alt="">
									</div>
									<div class="legend_block_img" :class="item.isActive ? '' : 'img_hide_style'"  v-if="item.name === 'Cosmos Hub (cosmoshub-4)'">
										<img src="../../assets/Cosmos_LOGO_Small.png" alt="">
									</div>
                                    <div class="legend_block_img" :class="item.isActive ? '' : 'img_hide_style'"  v-if="item.name === 'Akash (Akashnet-2)'">
                                        <img src="../../assets/logo_akash_menu.png" alt="">
                                    </div>
                                    <div class="legend_block_img" :class="item.isActive ? '' : 'img_hide_style'"  v-if="item.name === 'Crypto.com (Crypto-Org-Chain-Mainnet-1)'">
                                        <img src="../../assets/logo_crypto_menu.png" alt="">
                                    </div>
									<div class="legend_name_content">
										<p class="legend_name" :class="item.isActive ? '' : 'hide_style_color'">
											<span v-show="item.name !== 'Crypto.com (Crypto-Org-Chain-Mainnet-1)'
											 && item.name !== 'Cosmos Hub (cosmoshub-4)'
											 && item.name !== 'IRIS Hub (irishub-1)'
											 && item.name !== 'Akash (Akashnet-2)' ">{{item.name}}</span>
                                            <span class="legend_name_style" style="display: block">
                                               <span v-show="item.name === 'Crypto.com (Crypto-Org-Chain-Mainnet-1)'">Crypto.com</span>
                                                <el-tooltip effect="dark" content="(Crypto-Org-Chain-Mainnet-1)">
                                                   <span class="display_ellipsis" v-show="item.name === 'Crypto.com (Crypto-Org-Chain-Mainnet-1)'">
                                                       (Crypto-Org-Chain-Mainnet-1)
                                                   </span>
                                                </el-tooltip>
                                            </span>
                                            <span class="legend_name_style">
                                               <span v-show="item.name === 'Cosmos Hub (cosmoshub-4)'">Cosmos Hub</span>
											   <span class="display_ellipsis" v-show="item.name === 'Cosmos Hub (cosmoshub-4)'"> (cosmoshub-4)</span>
                                            </span>
                                            <span class="legend_name_style">
                                               <span v-show="item.name === 'IRIS Hub (irishub-1)'">IRIS Hub</span>
											   <span class="display_ellipsis" v-show="item.name === 'IRIS Hub (irishub-1)'"> (irishub-1)</span>
                                            </span>
                                            <span class="legend_name_style">
                                               <span v-show="item.name === 'Akash (Akashnet-2)'">Akash</span>
											   <span class="display_ellipsis" v-show="item.name === 'Akash (Akashnet-2)'"> (Akashnet-2)</span>
                                            </span>
                                        </p>
										<p class="legend_name" :class="item.isActive ? '' : 'hide_style_color'"
										   v-show="item.connection !== 0"><span
												style="display: flex;white-space: nowrap">{{item.connection}} {{`${item.connection > 1 ? 'Connections' :'Connection'}`}}</span>
										</p>
									</div>
								</li>
							</ul>
						</div>
					</div>
					<div class="pick_up_content" v-if="colorDataArray.length > 0">
						<div class="pick_up_img_content" @click="flPickUp()">
							<img v-show="!flShowDefaultPickUp" :src=" switchValue ? starLeftImg : defaultLeftImg "
							     alt="">
						</div>
					</div>
					<div class="pick_up_show_content">
						<div class="pick_up_show_img_content" @click="flShowPickUp()">
							<img v-show="flShowDefaultRightPickUp" :src=" switchValue ? starRightImg : defaultRightImg"
							     alt="">
						</div>
					</div>
				</div>
				<div class="error_content" v-if="!flShowNetwork">
					<p>Sorry, failed to obtain information Check if you are using vpn</p>
					<span @click="refreshPage()">Please refresh <i class="iconfont iconshuaxin"></i></span>
				</div>
				<div class="hover_up_content" v-if="flShowRevertIcon" v-show="flShowHoverUp" @click="scrollBottom()">
					<img style="width: 0.16rem;height:0.22rem;" src="../../assets/hover_up.gif" alt="">
				</div>
			</div>
<!--			<app-download></app-download>-->
<!--			<validator-bianjie-information></validator-bianjie-information>-->
		</div>
	</div>
</template>

<script>
	import axios from 'axios';
	import apiUrlConfig from "../../../config/config"
	import ValidatorBianjieInformation from "./ValidatorBianjieInformation";
	import AppDownload from "@/components/home/AppDownload";
	import Store from "@/store";
	var echarts = require('echarts/lib/echarts');
	require('echarts/lib/component/legend');
	require('echarts/lib/component/tooltip');
	require('echarts/lib/component/title');
	require('echarts/lib/chart/graph');
	require('echarts/extension/dataTool')
	require('echarts/lib/component/legendScroll');

	export default {
		name: "ValidatorGraph",
		components: {AppDownload, ValidatorBianjieInformation},
		data () {
			return {
				testData:null,
				graphContent: '',
				flConnectionActiveStyle: true,
				flLetterActiveStyle: false,
				sortByConnectionSwitchIcon: true,
				sortByLetterSwitchIcon: false,
				flShowConnectionSortIcon: true,
				flShowLetterSortIcon: false,
				flAllCheckout: true,
				flShowDefaultPickUp: false,
				flShowDefaultRightPickUp: false,
				showDefaultAnimation: false,
				defaultLeftImg: require('../../assets/default_pick_up_left.png'),
				defaultRightImg: require('../../assets/default_pick_up_right.png'),
				starLeftImg: require('../../assets/star_pick_up_left.png'),
				starRightImg: require('../../assets/star_pick_up_right.png'),
				flShowNetwork: true,
				data: null,
				flShowRevertIcon: false,
				colorDataArray: [],
				sortCopyData: [],
				copyData: '',
				flShowTestTooltip: true,
				dataTimer: null,
				flShowHoverUp: true,
				switchValue: sessionStorage.getItem('starSky') ? sessionStorage.getItem('starSky') === 'show' ? true : false : true,
				colorUseCopyData: '',
				timer: null,
				dataIndex: 2,
				maxLinks: 0,
				graphEcharts: null,
                linearColor:[
                    '#35B1EF',
                    '#DA6457',
                    '#125EEB',

                ],
				color: [
					'#CAA4FF',
					'#35B1EF',
					'#DA6457',
					'#125EEB',
					'#FBB27A',
					'#FECCB3',
					'#CFC2EC',
					'#F2D4EE',
					'#44C8ED',
					'#FF7CAD',
					'#AECFEB',
					'#FA9084',
					'#DF94AC',
					'#F2B8A9',
					'#BAC3FB',
					'#F3EFE9',
					'#DEE7EC',
					'#AFD3FC',
					'#F3D3EE',
					'#B8F28C',
					'#F27FCB',
					'#F7F6F7',
					'#4BB5FD',
					'#D3FCFC',
					'#FDAB74',
					'#FBDC94',
					'#75C6FF',
					'#ADDFF6',
					'#E8BDAF',
					'#ECE1D8',
					'#90B5EF',
					'#B2A9FF',
					'#18D6E4',
					'#CED8E7',
					'#E6C8F4',
					'#D3A8F4',
					'#DBD9D4',
					'#C4F3C9',
					'#77E9D6',
					'#4FDBC1',
					'#F19A6C',
					'#BED7F3',
					'#FDB2AD',
					'#FD8F8C',
					'#55B6E9',
					'#97BCFE',
					'#CF82E5',
					'#6ED1E7',
					'#FF9689',
					'#B7C6DA',
					'#66E9CE',
					'#E6FFC4',
					'#98EEFF',
					'#FFBEEE',
					'#58E2FF',
					'#B0FFB9',
					'#60F5BF',
					'#BBDDFF',
					'#EA98FF',
					'#AACCFF',
					'#D69EFD',
					'#FFCF9E',
					'#FFAB9F',
					'#FF839B',
					'#57D6FF',
					'#8CC2FF',
					'#CDBBFF',
					'#EB89DF',
					'#E6FFFE',
					'#F7FEFF',
					'#F9EBF9',
					'#6FDAFF',
					'#B5C9FF',
					'#AEFFE9',
					'#71E7FF',
					'#D3E2FF',
					'#C4DFFF',
					'#88FFFF',
					'#59DCFF',
					'#D293FF',
					'#EBE1FF',
					'#ADFAE1',
					'#72F4D2',
					'#29B5FF',
					'#85F4DF',
					'#4CDDF8',
					'#729EFF',
					'#E0CFFF',
					'#FF94EF',
					'#FF79CC',
					'#8FBEFF',
					'#64FFEE',
					'#D0B2FF',
					'#FFAA93',
					'#B2B8FF',
					'#33D4E0',
					'#DBE2FF',
					'#F193F8',
					'#FACE29',
					'#FF9A53',
					'#FDB6B2',
					'#DBAEF9',
					'#FBD0D6',
					'#FFA7B1',
					'#FBB27A',
					'#FECCB3',
					'#CFC2EC',
					'#F2D4EE',
					'#44C8ED',
					'#FF7CAD',
					'#AECFEB',
					'#FA9084',
					'#DF94AC',
					'#F2B8A9',
					'#BAC3FB',
					'#F3EFE9',
					'#DEE7EC',
					'#AFD3FC',
					'#F3D3EE',
					'#B8F28C',
					'#F27FCB',
					'#F7F6F7',
					'#4BB5FD',
					'#D3FCFC',
					'#FDAB74',
					'#FBDC94',
					'#75C6FF',
					'#ADDFF6',
					'#E8BDAF',
					'#ECE1D8',
					'#90B5EF',
					'#B2A9FF',
					'#18D6E4',
					'#CED8E7',
					'#E6C8F4',
					'#D3A8F4',
					'#DBD9D4',
					'#C4F3C9',
					'#77E9D6',
					'#4FDBC1',
					'#F19A6C',
					'#BED7F3',
					'#FDB2AD',
					'#FD8F8C',
					'#55B6E9',
					'#97BCFE',
					'#CF82E5',
					'#6ED1E7',
					'#FF9689',
					'#B7C6DA',
					'#66E9CE',
					'#E6FFC4',
					'#98EEFF',
					'#FFBEEE',
					'#58E2FF',
					'#B0FFB9',
					'#60F5BF',
					'#BBDDFF',
					'#EA98FF',
					'#AACCFF',
					'#D69EFD',
					'#FFCF9E',
					'#FFAB9F',
					'#FF839B',
					'#57D6FF',
					'#8CC2FF',
					'#CDBBFF',
					'#EB89DF',
					'#E6FFFE',
					'#F7FEFF',
					'#F9EBF9',
					'#6FDAFF',
					'#B5C9FF',
					'#AEFFE9',
					'#71E7FF',
					'#D3E2FF',
					'#C4DFFF',
					'#88FFFF',
					'#59DCFF',
					'#D293FF',
					'#EBE1FF',
					'#ADFAE1',
					'#72F4D2',
					'#29B5FF',
					'#85F4DF',
					'#4CDDF8',
					'#729EFF',
					'#E0CFFF',
					'#FF94EF',
					'#FF79CC',
					'#8FBEFF',
					'#64FFEE',
					'#D0B2FF',
					'#FFAA93',
					'#B2B8FF',
					'#33D4E0',
					'#DBE2FF',
					'#F193F8',
					'#FACE29',
					'#FF9A53',
					'#FDB6B2',
					'#DBAEF9',
					'#FBD0D6',
					'#FFA7B1',
					'#FBB27A',
					'#FECCB3',
					'#CFC2EC',
					'#F2D4EE',
					'#44C8ED',
					'#FF7CAD',
					'#AECFEB',
					'#FA9084',
					'#DF94AC',
					'#F2B8A9',
					'#BAC3FB',
					'#F3EFE9',
					'#DEE7EC',
					'#AFD3FC',
					'#F3D3EE',
					'#B8F28C',
					'#F27FCB',
					'#F7F6F7',
					'#4BB5FD',
					'#D3FCFC',
					'#FDAB74',
					'#FBDC94',
					'#75C6FF',
					'#ADDFF6',
					'#E8BDAF',
					'#ECE1D8',
					'#90B5EF',
					'#B2A9FF',
					'#18D6E4',
					'#CED8E7',
					'#E6C8F4',
					'#D3A8F4',
					'#DBD9D4',
					'#C4F3C9',
					'#77E9D6',
					'#4FDBC1',
					'#F19A6C',
					'#BED7F3',
					'#FDB2AD',
					'#FD8F8C',
					'#55B6E9',
					'#97BCFE',
					'#CF82E5',
					'#6ED1E7',
					'#FF9689',
					'#B7C6DA',
					'#66E9CE',
					'#E6FFC4',
					'#98EEFF',
					'#FFBEEE',
					'#58E2FF',
					'#B0FFB9',
					'#60F5BF',
					'#BBDDFF',
					'#EA98FF',
					'#AACCFF',
					'#D69EFD',
					'#FFCF9E',
					'#FFAB9F',
					'#FF839B',
					'#57D6FF',
					'#8CC2FF',
					'#CDBBFF',
					'#EB89DF',
					'#E6FFFE',
					'#F7FEFF',
					'#F9EBF9',
					'#6FDAFF',
					'#B5C9FF',
					'#AEFFE9',
					'#71E7FF',
					'#D3E2FF',
					'#C4DFFF',
					'#88FFFF',
					'#59DCFF',
					'#D293FF',
					'#EBE1FF',
					'#ADFAE1',
					'#72F4D2',
					'#29B5FF',
					'#85F4DF',
					'#4CDDF8',
					'#729EFF',
					'#E0CFFF',
					'#FF94EF',
					'#FF79CC',
					'#8FBEFF',
					'#64FFEE',
					'#D0B2FF',
					'#FFAA93',
					'#B2B8FF',
					'#33D4E0',
					'#DBE2FF',
					'#F193F8',
					'#FACE29',
					'#FF9A53',
				],

			}
		},
		watch: {
			colorUseCopyData: {
				deep: true,
				handler: function () {
					//判断是否点击了选项
					let AllIsActive = this.colorDataArray.every((item) => {
						return item.isActive !== false
					});
					//控制all的样式
					if (!AllIsActive) {
						this.flAllCheckout = false;
					} else {
						this.flAllCheckout = true;
					}
					// 处理选中节点的显示
					let nameArray = [];
					this.colorDataArray.forEach(item => {
						if (item.isActive) {
							nameArray.push({
								name: item.name,
							})
						}
					});

					let needShowNodeArray = [];
					this.data.nodes.forEach(item => {
						nameArray.forEach(value => {
							if (item['chain-id'] === value.name) {
								needShowNodeArray.push(item)
							}
						})
					});
					this.copyData.nodes = needShowNodeArray;
					this.initChartsGraph();
				}
			}
		},
		mounted () {
			this.testData  = require('../../../config/ibc-node.json')
			/*
						let container = new PerfectScrollbar('#graph_list_container',{
							wheelSpeed: 2,
							wheelPropagation: true,
							minScrollbarLength: 20
						});
						container.update()*/
			this.$refs.chart_content.style.height = (window.innerHeight - 164) + "px"
			this.$refs.graph_list_content.style.height = (window.innerHeight - 288) + "px"
			window.addEventListener('resize', this.onresize)
			clearTimeout(this.timer);
			this.getData();
			this.timer = setInterval(() => {
				this.getData();
			}, 300000);
			window.addEventListener("scroll", this.handleScroll, true)
		},
		methods: {
			handleScroll (e) {
				let scrollTop = document.documentElement.scrollTop || document.body.scrollTop
				if (scrollTop > 5) {
					this.flShowHoverUp = false
				} else {
					this.flShowHoverUp = true
				}

			},
			scrollBottom () {
				window.pageYOffset = 10000000;
				document.documentElement.scrollTop = 10000000;
				document.body.scrollTop = 10000000;
			},
			revertGraph () {
				window.location.reload();
			},
			sortByConnection (flSwitchValue) {
				this.flShowLetterSortIcon = false;
				this.flShowConnectionSortIcon = true;
				let copyData = JSON.parse(JSON.stringify(this.colorDataArray))
				this.flConnectionActiveStyle = true;
				this.flLetterActiveStyle = false;
				this.sortByConnectionSwitchIcon = !this.sortByConnectionSwitchIcon
				if (flSwitchValue) {
					this.colorDataArray = copyData.sort((a, b) => {
						return a.connection - b.connection
					})
				} else {
					this.colorDataArray = copyData.sort((a, b) => {
						return b.connection - a.connection
					})
				}
			},
			sortByLetter (sortRule) {
				this.flShowLetterSortIcon = true;
				this.flShowConnectionSortIcon = false;
				let copyData = JSON.parse(JSON.stringify(this.colorDataArray))
				this.flConnectionActiveStyle = false;
				this.flLetterActiveStyle = true;
				this.sortByLetterSwitchIcon = !this.sortByLetterSwitchIcon;
				this.colorDataArray = copyData.sort((a, b) => {
					let value1 = a.name.toLowerCase(), value2 = b.name.toLowerCase();
					if (sortRule) {
						if (value1 > value2) {
							return -1
						} else {
							return 1
						}
					} else {
						if (value1 < value2) {
							return -1
						} else {
							return 1
						}
					}

				})
			},
			onresize () {
				this.$refs.chart_content.style.height = (window.innerHeight - 164) + "px"
				// this.$refs.graph_list_content.style.height = (window.innerHeight - 288) + "px"
				this.graphEcharts.resize()
			},
			flPickUp () {
				this.showDefaultAnimation = false
				this.flShowDefaultPickUp = true
				setTimeout(() => {
					this.flShowDefaultRightPickUp = true
				}, 1000)
			},
			flShowPickUp () {
				this.showDefaultAnimation = true
				this.flShowDefaultRightPickUp = false
				setTimeout(() => {
					this.flShowDefaultPickUp = false
				}, 1000)
			},
			/*switchBgColor(switchValue){
				// true 为星空背景
				if(switchValue){
					sessionStorage.setItem('starSky','show')
					this.$uMeng.push('TurnOn_starry-skymode','click')
				}else {
					this.$uMeng.push('TurnOff_starry-skymode','click')
					sessionStorage.setItem('starSky','hide')
				}
			},*/
			getData () {
				this.$store.commit('flShowLoading', true)
				let Url = apiUrlConfig.app.url;
				/*axios.get(`${Url}`).then(data => {
					if (data.status === 200) {
						return data.data
					} else {
						return 'network is error'
					}
				}).then(res => {*/
				const testData = this.testData;
				this.$store.commit('flShowLoading', true)
				let res = testData
				if (res && JSON.stringify(res) !== '{}') {
					this.data = res;
					this.flShowTestTooltip = res.production
					//数据先排序
					this.data.nodes.sort((a, b) => {
						return b.connections - a.connections
					});
					this.maxLinks = this.data.nodes[0].connections
					this.data.nodes.forEach((item, index) => {
						item.isDelete = false;
						item.color = this.color[index]
					});
					this.copyData = JSON.parse(JSON.stringify(this.data));
					this.colorUseCopyData = JSON.parse(JSON.stringify(this.data));
					this.initLegend();
					this.initChartsGraph();
				} else {
					this.flShowNetwork = false
					this.$store.commit('flShowLoading', false)
				}
				/*}).catch(e => {
					console.log(e);
					this.$store.commit('flShowLoading', false)
					if (e) {
						this.flShowNetwork = false
					}
				})*/
			},
			refreshPage () {
				window.location.reload();
			},
			selectAll () {
				this.flAllCheckout = !this.flAllCheckout;
				let flIsActive = this.colorDataArray.every((item) => {
					return item.isActive !== false
				});
				if (flIsActive) {
					this.colorDataArray.forEach(item => {
						item.isActive = false
					});
					this.copyData.nodes = [];
					this.initChartsGraph();
				} else {
					this.colorDataArray.forEach(item => {
						item.isActive = true;
						this.copyData = JSON.parse(JSON.stringify(this.data));
					})
					this.initChartsGraph();
				}
			},
			initLegend () {
				this.colorDataArray = [];
				this.copyData.nodes.forEach((item) => {
					this.colorDataArray.push({
						name: item['chain-id'],
						color: item.color,
						isActive: true,
						connection: item.connections,
					})
				});
				this.sortCopyData = JSON.parse(JSON.stringify(this.colorDataArray));
			},
			changeChart (value, index) {
				this.colorDataArray[index].isActive = !this.colorDataArray[index].isActive;
				this.colorUseCopyData.nodes.forEach((item) => {
					if (item['chain-id'] === value.name) {
						item.isDelete = !item.isDelete;
					}
				});
				this.copyData.nodes = this.colorUseCopyData.nodes.filter(item => {
					return item.isDelete === false
				});
			},
            setNodeImage(nodeName){
                switch (nodeName){
                    case 'IRIS Hub (irishub-1)' :
                        return 'image://data:image/png;base64,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'
                    case 'Cosmos Hub (cosmoshub-4)':
                        return 'image://data:image/png;base64,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'
                    case 'Akash (Akashnet-2)':
                        return 'image://data:image/png;base64,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'
                    case 'Crypto.com (Crypto-Org-Chain-Mainnet-1)':
                        return 'image://data:image/png;base64,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'
                }
            },
            formatName(nodeName){
                let formatNameArr  = nodeName.split(' ')
                const name = formatNameArr.filter( (item,index) => {
                    if(index >= 1){
                        return  index !== (formatNameArr.length - 1)
                    }
                    return  item
                }).join(' ')
                return name
            },
			initChartsGraph () {
				this.flShowRevertIcon = false
				this.graphEcharts = echarts.init(document.getElementById('validator_graph_content'));
				let nodeLinksArray = [], nodeArray = [];
				//最大像素点与最小像素点的差值66  最小的symbolSize 为 8 * 节点递增的比例
				let symbolSizeRule = 50;
				//数据结果展示
				let minSymbolSizeRule = Math.floor(20 / (Number(symbolSizeRule) / this.data.nodes.length))
				for (let i in this.copyData.nodes) {
					let connectionValue = this.copyData.nodes[i].connections;
					const name = this.formatName(this.copyData.nodes[i]['chain-id'])
					nodeArray.push({
						name: name,
                        nodeName: this.copyData.nodes[i]['chain-id'],
						symbolSize: this.copyData.nodes[i].connections === 1 ? 50 : this.copyData.nodes[i].connections * (symbolSizeRule / this.maxLinks) + 20,
						label: {
							show: true,//默认展示信息
							position: 'bottom',
						},
                        /*symbol: this.copyData.nodes[i]['chain-id'] === 'IRIS Hub (irishub-1)' ? 'image://data:image/png;base64,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'
                            : this.copyData.nodes[i]['chain-id'] === 'Cosmos Hub (cosmoshub-4)' ? 'image://data:image/png;base64,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' : '' ,*/
						symbol: this.setNodeImage(this.copyData.nodes[i]['chain-id']),
						itemStyle: {
							color: this.copyData.nodes[i].color,
						},
						tooltip: {
							backgroundColor:'rgba(34,33,78,0.5)',
							borderColor: this.copyData.nodes[i].color,
							borderWidth:1,
							borderRadius: 15,
							padding: [5, 20],
                            formatter:function (value){
                                let name = value.data.nodeName.split(' ')
                                return  name[name.length -1].replace(/[(|)]/g,'')
                            }
						}
					});
				}
				let categories = [];
				this.copyData.nodes.forEach(item => {
					categories.push({
						name: item['chain-id']
					})
				})

				this.copyData.paths.forEach((item, index) => {
                    const sourceName = this.formatName(item['src-chain-id'])
                    const targetName = this.formatName(item['dst-chain-id'])
					if (item.state === 'OPEN') {
						nodeLinksArray.push({
							source: sourceName,
							target: targetName,
                            lineStyle:{
                                color: 'rgba(112, 198, 199, 1)',
                                width:1,
                            },
							emphasis: {
								lineStyle: {
									width: 1,
									type: 'solid',
								}
							}
						})
					} else {
						nodeLinksArray.push({
                            source: sourceName,
                            target: targetName,
							//连接线的样式设置
							lineStyle: {
                                color: 'rgba(112, 198, 199, 1)',
								type: 'dashed',
							},
							//鼠标滑到连接线上的样式配置
							emphasis: {
								lineStyle: {
									width: 1,
									type: 'dashed',
								}
							}
						})
					}

				});
				let graphOption = {
					name: '',
					tooltip: {
						trigger:'item',
						formatter: function (data) {
							return `<span>${data.name.toString().replace('>', '<->')}</span>`
						}
					},
					series: [
						{
							type: 'graph',
							layout: 'force',
							animation: false,
							// animationDuration: 10000,
							// animationThreshold: 1000,
							animationEasingUpdate: 'exponentialOut', //缓动效果
							// animationDurationUpdate:300,
							// edgeSymbol: ['arrow','arrow'], //指示线的箭头
							data: nodeArray,
							// links: nodeLinksArray,
							nodeScaleRatio: 0.6, //鼠标每次缩放的整体缩放比例
							force: {
								repulsion: [800, 1000], //斥力因子
								gravity: 0.8, //是否向中心靠拢 值越大越接近于中心
								edgeLength: [100, 200], //链接线的长度范围
								layoutAnimation: true,
							},
							// zoom:0.1, //设置整体视图缩放的比例
							zoom: this.dataIndex,
							// symbol:'',//图形形状
							roam: true, //平移和缩放  move 和 scale true 包含所有
							draggable: true, //是否拖拽
							hoverAnimation: true,
							focusNodeAdjacency: true,
							/*itemStyle: {
								shadowColor: 'rgba(255, 255, 255, 0.8)',
								shadowOffsetY:2,
								shadowBlur: 3,
								opacity:1
							},*/
							emphasis: {
								lineStyle: {
									width: 1,
								},
								label: {
									show: false,
								}
							}
						}
					]
				};
				this.graphEcharts.setOption(graphOption)
				clearInterval(this.dataTimer)
				// 根据节点数使用不同的缩放规则
				let zoomRule = 1, zoomSpeedRule = 0.05, setTime = 950;
				if (nodeArray.length > 150) {
					zoomRule = 0.3;
					zoomSpeedRule = 0.2;
					setTime = 1450
				} else if (nodeArray.length > 100 && nodeArray.length < 150) {
					zoomRule = 0.6;
					zoomSpeedRule = 0.2;
					setTime = 1250
				} else if (nodeArray.length > 50 && nodeArray.length < 100) {
					zoomRule = 0.7;
					zoomSpeedRule = 0.04
					setTime = 950
				} else if (nodeArray.length < 50) {
					zoomRule = 1;
					zoomSpeedRule = 0.07
					setTime = 950
				}
				this.$store.commit('flShowLoading', false)
				// 渲染完成收缩图形让其在区域内全部展示

				this.dataTimer = setInterval(() => {
					graphOption.series[0].draggable = false
					this.dataIndex = this.dataIndex - 0.07;
					if (this.dataIndex < zoomRule) {
						this.dataIndex = zoomRule
						clearInterval(this.dataTimer);
						graphOption.series[0].draggable = true
					}
					graphOption.series[0].zoom = this.dataIndex;
					this.graphEcharts.setOption(graphOption)
				}, 20)

				//最后一次渲染
				setTimeout(() => {
					graphOption.series[0].zoom = zoomRule;
					graphOption.series[0].links = nodeLinksArray;
					graphOption.series[0].force.gravity = 0.3
					this.graphEcharts.setOption(graphOption);
					this.flShowRevertIcon = true
				}, setTime)
			},
		}
	}
</script>

<style scoped lang="scss">
	.default_bg_img {
		background: rgba(32, 35, 66, 1);
	}

	.start_sky_img {
		background-image: url("../../assets/GOZ_bg.png");
		background-repeat: no-repeat;
		background-position: center;
		background-size: cover;
	}

	.validator_graph_container {
		width: 100%;
		height: 100%;
		//线上数据
		.graph_content_wrap {
			width: 100%;
			box-sizing: border-box;
			padding: 0 0.4rem;
			margin: 0 auto;
			margin-bottom: 0.2rem;
			.graph_content_title {
				height: 0.6rem;
				font-size: 0.18rem;
				color: rgba(255, 255, 255, 1);
				display: flex;
				align-items: center;
				.beat_content {
					margin-left: 0.18rem;
					background: rgba(55, 124, 248, 1);
					color: #fff;
					font-size: 0.14rem;
					border-radius: 0.1rem;
					padding: 0.02rem 0.1rem;

				}

				.tooltip {
					display: flex;
					justify-content: space-between;

					.graph_tooltip {
						display: flex;

						p:nth-of-type(1) {
							display: flex;
							align-items: center;

							span:nth-of-type(1) {
								margin-left: 0.2rem;
								display: inline-block;
								width: 0.2rem;
								height: 0.02rem;
								border-top: 0.02rem solid #70C6C7;
							}

							span:nth-of-type(2) {
								margin-left: 0.1rem;
								font-size: 0.14rem;
								color: rgba(115, 122, 174, 1);
							}
						}

						p:nth-of-type(2) {
							display: flex;
							align-items: center;

							span:nth-of-type(1) {
								width: 0.2rem;
								margin-left: 0.41rem;
								display: inline-block;
								border-top: 0.02rem dashed #70C6C7;
							}

							span:nth-of-type(2) {
								margin-left: 0.1rem;
								font-size: 0.14rem;
								color: rgba(115, 122, 174, 1);
							}
						}
					}

					.background_toggle_content {
						margin-left: 0.35rem;
					}
				}

			}

			.graph_charts_container {
				height: 100%;

				.graph_charts_title {
					padding: 0.2rem 0 0.2rem 0.2rem;
					font-size: 0.12rem;
					line-height: 0.14rem;
					color: rgba(115, 122, 174, 1);
				}

				.graph_content_container {
					height: 100%;
					display: flex;
					position: relative;

					.iconfuwei {
						position: absolute;
						left: 0;
						top: 0;
						cursor: pointer;
						color: #377CF8;
						font-size: 0.24rem;
						z-index: 9;
					}

					.graph_container_content_wrap {
						flex: 1;
						background: #2d325a;
						border-radius: 0.04rem;

						#validator_graph_content {
							/*max-width: 8rem;*/
							width: 100%;
							background: transparent;
							//线上
							/*background: #2D325A;*/
						}
					}

					.sort_content {
						box-sizing: border-box;
						padding-left: 0.38rem;
						margin-bottom: 0.22rem;

						.active_style {
							color: #377CF8;
						}

						.un_active_style {

						}

						span {
							cursor: pointer;
							margin-right: 0.2rem;
							font-size: 0.14rem;
							color: #868FD3;
							white-space: nowrap;
						}
					}

					.sort_scroll_content::-webkit-scrollbar {
						width: 0.05rem;
						height: 0.05rem;
					}

					.sort_scroll_content::-webkit-scrollbar-thumb {
						background: rgba(63, 67, 105, 1);
						border-radius: 0.05rem;
					}

					/*.sort_scroll_content::-webkit-scrollbar-thumb:hover{
						background: #333;
					}
					.sort_scroll_content::-webkit-scrollbar-corner{
						background: #179a16;
					}*/
					.graph_list_container {
						scrollbar-width: none;
					}

					.graph_list_container {
						margin-left: 0.02rem;
						padding-top: 0.2rem;
						border-radius: 0.04rem;
						position: absolute;
						overflow: hidden;
						right: 0;
						top: 0;
						background: rgba(95, 115, 255, 0.15);
						backdrop-filter: blur(2px);
						.sort_scroll_content {
							overflow-x: hidden;
							overflow-y: auto;
							margin-top: 0.02rem;
							@media(max-width:1040px ){
								height: 2rem !important;
							}
							.graph_list_item_all {
								/*position: fixed;*/
								display: flex;
								justify-content: flex-start;
								margin: 0 0.6rem 0.2rem 0.38rem;
								background: transparent;

								.legend_all_block {
									width: 0.14rem;
									height: 0.14rem;
									border-radius: 0.07rem;
									cursor: pointer;
									background: transparent;
									img {
										width: 0.14rem;
									}
								}

								.legend_block {
									box-sizing: border-box;
									width: 0.14rem;
									height: 0.14rem;
									border-radius: 0.07rem;
									cursor: pointer;
								}

								.hide_style {
									color: rgba(134, 143, 211, 1) !important;
								}

								.legend_name {
									cursor: pointer;
									height: 0.14rem;
									color: rgba(134, 143, 211, 0.5);
									margin-left: 0.1rem;
									width: 0.9rem;
									white-space: nowrap;
								}

								.legend_name_content {
									.legend_name {
										cursor: pointer;
										height: 0.14rem;
										color: rgba(134, 143, 211, 0.5);
										margin-left: 0.1rem;
										width: 0.9rem;
										white-space: nowrap;

									}
								}
							}

							.graph_list_content {
								flex: 1;
								max-width: 2.2rem;
								/*margin-top: 0.34rem;*/
								min-height: 0.14rem;
								display: flex;
								flex-wrap: wrap;
								margin-left: 0.18rem;
								padding-bottom: 0.4rem;

								.graph_list_item {
									width: 2rem;
									align-items: center;
									margin-left: 0.1rem;
									display: flex;
									padding: 0.05rem 0 0.1rem 0;
									cursor: pointer;

									.legend_all_block {
										width: 0.14rem;
										height: 0.14rem;
										border-radius: 0.07rem;
										cursor: pointer;

										img {
											width: 0.14rem;
										}
									}

									.legend_name_content {
										max-width: 1.6rem;

										.legend_name {
											line-height: 0.16rem;
											color: rgba(134, 143, 211, 1);
											margin-left: 0.1rem;
										}
										.legend_name:first-child{
											font-weight: 600;
											margin-bottom: 0.03rem;
                                            width: 1.6rem;
                                            word-break: keep-all;
                                            display:-webkit-box;
                                            -webkit-box-orient:vertical;
                                            -webkit-line-clamp:2;
                                            overflow:hidden;
                                            .legend_name_style{
                                                display: flex;
                                                flex-direction: column;
                                                .display_ellipsis{
                                                    display:-webkit-box;
                                                    -webkit-box-orient:vertical;
                                                    -webkit-line-clamp:1;
                                                    overflow:hidden;
                                                }
                                            }
										}
										.hide_style_color {
											color: rgba(134, 143, 211, 0.5) !important;
										}
									}

									.hide_style {
										background: transparent !important;
										border: 0.01rem solid rgba(134, 143, 211, 0.5);
									}

									.legend_block {
										box-sizing: border-box;
										width: 0.14rem;
										height: 0.14rem;
										border-radius: 0.07rem;
									}
									.legend_block_img{
										box-sizing: border-box;
										width: 0.28rem;
										height: 0.28rem;
										border-radius: 0.07rem;
										cursor: pointer;
										opacity: 1;
										img{
											width: 100%;
										}
									}
									.img_hide_style{
										opacity: 0.4;
									}
								}
							}
						}
					}

					.pick_up_content {
						position: absolute;
						right: 2.29rem;
						top: 50%;
						z-index: 10;

						.pick_up_img_content {
							width: 0.14rem;
							height: 0.6rem;
							cursor: pointer;

							img {
								width: 100%;
							}
						}
					}

					.pick_up_show_content {
						position: absolute;
						top: 50%;
						right: 0;

						.pick_up_show_img_content {
							width: 0.14rem;
							height: 0.6rem;
							cursor: pointer;

							img {
								width: 100%;
							}
						}
					}

					.hide_pick_up_style {
						animation: hideGraphOptionList 1s ease-in;
						animation-fill-mode: forwards;
					}

					.show_pick_up_style {
						animation: showGraphOptionList 1s ease-in;
						animation-fill-mode: forwards;
					}
				}
			}

			.hover_up_content {
				position: absolute;
				bottom: 0;
				left: 50%;
				cursor: pointer;
				display: flex;
				justify-content: center;
				margin-top: 0.05rem;
			}

			.show_error_content {
				min-height: 6rem;
				display: flex;
				align-items: center;
				justify-content: center;

				.error_content {
					text-align: center;

					p {
						max-width: 2.05rem;
						color: rgba(115, 122, 174, 1);
						font-size: 0.14rem;
						margin-bottom: 0.2rem;
					}

					span {
						cursor: pointer;
						text-align: center;
						color: rgba(55, 124, 248, 1);
						font-size: 0.14rem;
						line-height: 0.16rem;
						border-bottom: 0.01rem solid rgba(55, 124, 248, 1);
					}
				}

			}
		}
	}

	@media (max-width: 1040px) {
		.validator_graph_container {
			.graph_content_wrap {
				box-sizing: border-box;
				padding: 0 0.3rem;

				.graph_content_title {
					height: auto;
					padding: 0.15rem 0;
					align-items: flex-start;
				}

				.tooltip {
					justify-content: space-between;

					.graph_tooltip {
						flex: 1;
						display: flex;
						align-items: flex-start;
					}

					.background_toggle_content {
						text-align: right;
					}
				}

				.graph_charts_container {
					.graph_content_container {
						flex-direction: column;

						.sort_content {
							padding-left: 0.2rem;
						}

						.graph_container_content_wrap {
							border-radius: 0;
						}

						.graph_list_container {
							max-width: none;
							height: 2rem !important;
							margin-bottom: 0.2rem;
							border-radius: 0;
							margin-left: 0;
							position: static;
							.sort_scroll_content {
								.graph_list_item_all {
									width: 1.2rem;
									margin-left: 0.2rem;
									justify-content: flex-start;
								}

								.graph_list_content {
									height: auto;
									max-width: none;

									.graph_list_item {
										margin-left: 0;
									}
								}
							}

						}

						.pick_up_content {
							display: none;
						}
					}
				}
			}
		}
	}

	@keyframes rotate {
		from {
			transform: rotate(0deg);
		}
		to {
			transform: rotate(360deg);
		}
	}

	@media (max-width: 900px) {
		.validator_graph_container {
			.graph_content_wrap {
				.graph_content_title {
					flex-direction: column;

					.beat_content {
						/deep/ .tooltip_content {
							width: 3rem;
						}
					}

					.tooltip {
						margin-top: 0.1rem;
						flex-direction: row;

						.graph_tooltip {
							flex-direction: row;

							p:nth-of-type(1) {
								span:nth-of-type(1) {
									margin-left: 0;
								}
							}

							p:nth-of-type(2) {
								span:nth-of-type(1) {
									margin-left: 0;
								}
							}
						}

						.background_toggle_content {
							margin-right: 0.1rem;
						}
					}
				}
			}
		}
	}

	@media (max-width: 610px) {
		.validator_graph_container {
			.graph_content_wrap {
				.graph_content_title {
					.tooltip {
						flex-direction: column;

						.graph_tooltip {
							flex-direction: column;

							p:nth-of-type(1) {
								span:nth-of-type(1) {
									margin-left: 0;
								}
							}

							p:nth-of-type(2) {
								span:nth-of-type(1) {
									margin-left: 0;
								}
							}
						}

						.background_toggle_content {
							margin-top: 0.1rem;
							margin-left: 0;
						}
					}
				}

				.graph_charts_container {
					.graph_charts_title {
						padding-left: 0;
						padding-top: 0;
					}

					.graph_content_container {
						.iconfuwei {
							padding-left: 0;
						}

						.sort_content {
							padding-left: 0.2rem;
						}

						.graph_list_container {
							height: 2rem !important;
							margin-bottom: 0.2rem;
						}
					}
				}
			}
		}
	}

	@keyframes hideGraphOptionList {
		from {
			width: 2.43rem;
		}
		to {
			width: 0;
		}
	}

	@keyframes showGraphOptionList {
		from {
			width: 0;
		}
		to {
			width: 2.43rem;
		}
	}
</style>

