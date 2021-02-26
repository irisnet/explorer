<template>
	<div class="validator_graph_container" :class="switchValue ? 'start_sky_img' : 'default_bg_img'">
		<div class="graph_content_wrap">
			<div class="graph_content_title">
				<div>Inter-Chain Network
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
							      @click="sortByConnection(sortByConnectionSwitchIcon)">Conns <i class="iconfont"
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
									     v-if="item.name !== 'Cosmos Hub (cosmoshub-4)' && item.name !== 'IRIS Hub (irishub-1)'">
									</div>
									<div class="legend_block_img" v-if=" item.name === 'IRIS Hub (irishub-1)'">
										<img src="../../assets/IRISnet_LOGO_Small.png" alt="">
									</div>
									<div class="legend_block_img"  v-if="item.name === 'Cosmos Hub (cosmoshub-4)'">
										<img src="../../assets/Cosmos_LOGO_Small.png" alt="">
									</div>
									<div class="legend_name_content">
										<p class="legend_name" :class="item.isActive ? '' : 'hide_style_color'">
											{{item.name}} </p>
										<p class="legend_name" :class="item.isActive ? '' : 'hide_style_color'"
										   v-show="item.connection !== 0"><span
												style="display: flex;white-space: nowrap">{{item.connection}} {{`${item.connection > 1 ? 'Conns' :'Conn'}`}}</span>
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
				color: [
					'#CAA4FF',
					'#95B8FF',
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
			initChartsGraph () {
				this.flShowRevertIcon = false
				this.graphEcharts = echarts.init(document.getElementById('validator_graph_content'));
				let nodeLinksArray = [], nodeArray = [];
				//最大像素点与最小像素点的差值66  最小的symbolSize 为 8 * 节点递增的比例
				let symbolSizeRule = 30;
				//数据结果展示
				let minSymbolSizeRule = Math.floor(20 / (Number(symbolSizeRule) / this.data.nodes.length))
				for (let i in this.copyData.nodes) {
					let connectionValue = this.copyData.nodes[i].connections;
					nodeArray.push({
						name: this.copyData.nodes[i]['chain-id'],
						symbolSize: this.copyData.nodes[i].connections === 1 ? 50 : this.copyData.nodes[i].connections * (symbolSizeRule / this.maxLinks) + 20,
						label: {
							show: true,//默认展示信息
							position: 'bottom',
						},
						symbol: this.copyData.nodes[i]['chain-id'] === 'IRIS Hub (irishub-1)' ? 'image://data:image/png;base64,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'
							: this.copyData.nodes[i]['chain-id'] === 'Cosmos Hub (cosmoshub-4)' ? 'image://data:image/png;base64,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' : '' ,
						itemStyle: {
							color: this.copyData.nodes[i].color,
						},
						tooltip: {
							backgroundColor:'rgba(34,33,78,0.5)',
							borderColor: this.copyData.nodes[i].color,
							borderWidth:1,
							borderRadius: 15,
							padding: [5, 20]
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
					if (item.state === 'OPEN') {
						nodeLinksArray.push({
							source: item['src-chain-id'],
							target: item['dst-chain-id'],
							lineStyle: {
								color: 'rgba(112, 198, 199, 1)',
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
							source: item['src-chain-id'],
							target: item['dst-chain-id'],
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
						background: rgba(45, 50, 90, 0.2);
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
										width: 0.32rem;
										height: 0.32rem;
										border-radius: 0.07rem;
										cursor: pointer;
										img{
											width: 100%;
										}
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

