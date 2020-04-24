<template>
	<div class="validator_graph_container">
		<div class="graph_content_wrap">
			<div class="graph_content_title"><div>GoZ Global Network State*</div> <span class="beat_content">Beta</span> </div>
			<div class="tooltip" v-if="flShowNetwork"><p><span></span><span>Connection Opened</span></p> <p><span></span><span>Connection Unopened</span></p></div>
			<div class="graph_charts_container" :class="flShowNetwork ? '' : 'show_error_content'">
				<p class="graph_charts_title" v-if="flShowNetwork">*This demo is using simulated data. Please stay tuned for the grand GoZ Opening at May 1st</p>
				<div class="graph_content_container" v-if="flShowNetwork">
					<div id="validator_graph_content"></div>
					<div class="graph_list_container">
						<div class="graph_list_item_all" @click="selectAll()">
							<div class="legend_all_block">
								<img v-show="flAllCheckout" src="../../assets/select_all.svg" alt="">
								<img v-show="!flAllCheckout" src="../../assets/unselect_all.svg" alt=""></div>
							<span class="legend_name" :class=" flAllCheckout ? 'hide_style' : ''">All Zones</span>
						</div>
						<ul class="graph_list_content">
							
							<li class="graph_list_item" v-for="(item,index) in colorDataArray" @click="changeChart(item,index)">
								<div class="legend_block"
								     :style="{background: item.color}"
								     :class="item.isActive ? '' : 'hide_style'"></div>
								<span class="legend_name" :class="item.isActive ? '' : 'hide_style_color'">{{item.name}}</span>
							</li>
						</ul>
					</div>
				</div>
				<div class="error_content" v-if="!flShowNetwork">
					<p>Sorry, failed to obtain information Check if you are using vpn</p>
					<span @click="refreshPage()">Please refresh <i class="iconfont iconshuaxin" ></i></span>
				</div>
				
			</div>
		</div>
	</div>
</template>

<script>
	import axios from 'axios';
	import apiUrlConfig from "../../../config/config"
	var echarts = require('echarts/lib/echarts');
	require('echarts/lib/component/legend');
	require('echarts/lib/component/tooltip');
	require('echarts/lib/component/title');
	require('echarts/lib/chart/graph');
	require('echarts/extension/dataTool')
	require('echarts/lib/component/legendScroll');
	export default {
		name: "ValidatorGraph",
		data(){
			return {
				graphContent:'',
				flAllCheckout: true,
				flShowNetwork: true,
				data:null,
				colorDataArray:[],
				copyData:'',
				testData:{
					"nodes": [
						{
							"chain-id": "morpheus-ibc-3000",
							"connections": 2
						},
						{
							"chain-id": "pylonchain",
							"connections": 5
						},
						{
							"chain-id": "irishub",
							"connections": 8
						},
						{
							"chain-id": "achain",
							"connections": 1
						},
						{
							"chain-id": "ping-ibc",
							"connections": 2
						},
						{
							"chain-id": "kappa",
							"connections": 3
						},
						{
							"chain-id": "crusty",
							"connections": 2
						},
						{
							"chain-id": "vitwitchain",
							"connections": 2
						},
						{
							"chain-id": "umbrellachain",
							"connections": 6
						},
						{
							"chain-id": "ptpchain",
							"connections": 2
						},
						{
							"chain-id": "ibc-band-testnet1",
							"connections": 2
						},
						{
							"chain-id": "kava-ibc",
							"connections": 2
						},
						{
							"chain-id": "melea-11",
							"connections": 1
						},
						{
							"chain-id": "dokia",
							"connections": 0
						}
					],
					"paths": [
						{
							"src": "morpheus-ibc-3000",
							"dst": "pylonchain",
							"state": "I"
						},
						{
							"src": "irishub",
							"dst": "morpheus-ibc-3000",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "pylonchain",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "ping-ibc",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "kappa",
							"state": "OPEN"
						},
						{
							"src": "achain",
							"dst": "crusty",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "vitwitchain",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "umbrellachain",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "ptpchain",
							"state": "OPEN"
						},
						{
							"src": "irishub",
							"dst": "kava-ibc",
							"state": "OPEN"
						},
						{
							"src": "ping-ibc",
							"dst": "pylonchain",
							"state": "OPEN"
						},
						{
							"src": "kappa",
							"dst": "crusty",
							"state": "OPEN"
						},
						{
							"src": "vitwitchain",
							"dst": "umbrellachain",
							"state": "OPEN"
						},
						{
							"src": "ptpchain",
							"dst": "ibc-band-testnet1",
							"state": "OPEN"
						},
						{
							"src": "ibc-band-testnet1",
							"dst": "pylonchain",
							"state": "OPEN"
						},
						{
							"src": "umbrellachain",
							"dst": "pylonchain",
							"state": "OPEN"
						},
						{
							"src": "kava-ibc",
							"dst": "umbrellachain",
							"state": "OPEN"
						},
						{
							"src": "melea-11",
							"dst": "umbrellachain",
							"state": "OPEN"
						},
						{
							"src": "umbrellachain",
							"dst": "kappa",
							"state": "OPEN"
						}
					]
				},
				colorUseCopyData: '',
				timer: null,
				color:[
					'#4BB5FD',
					'#CFC2EC',
					'#4FDBC1',
					'#DBAEF9',
					'#ADDFF6',
					'#DF94AC',
					'#CF82E5',
					'#F27FCB',
					'#FBB27A',
					'#B2A9FF',
					'#90B5EF',
					'#D3A8F4',
					'#77E9D6',
					'#97BCFE',
					'#E6C8F4',
					'#AFD3FC',
					'#B8F28C',
					'#BED7F3',
					'#FDB2AD',
					'#55B6E9',
					'#B7C6DA',
					'#FF9689',
					'#6ED1E7',
					'#AECFEB',
					'#FF7CAD',
					'#44C8ED',
					'#B2A9FF',
					'#FDB6B2',
					'#FBDC94',
					'#D3FCFC',
					'#FDAB74',
				],
			}
		},
		watch:{
			colorUseCopyData:{
				deep:true,
				handler:function () {
					//判断是否点击了选项
					let AllIsActive = this.colorDataArray.every((item) => {
						return item.isActive !== false
					});
					//控制all的样式
					if(!AllIsActive){
						this.flAllCheckout = false;
					}else {
						this.flAllCheckout = true;
					}
					// 处理选中节点的显示
					let nameArray = [];
					this.colorDataArray.forEach(item => {
						if(item.isActive){
							nameArray.push({
								name: item.name
							})
						}
					});
					let needShowNodeArray = [];
					this.data.nodes.forEach( item => {
						nameArray.forEach(value => {
							if(item['chain-id'] === value.name){
								needShowNodeArray.push(item)
							}
						})
					});
					this.copyData.nodes = needShowNodeArray;
					this.initChartsGraph();
				}
			}
		},
		mounted(){
			clearTimeout(this.timer);
			this.getData();
			this.timer = setInterval(() => {
				this.getData();
			},300000)
		},
		methods:{
			getData(){
				let Url = apiUrlConfig.app.url;
				axios.get(`${Url}`).then(data => {
					if(data.status === 200){
						return data.data
					}else {
						return 'network is error'
					}
				}).then( res => {
					this.data = res;
					// TODO 演示专用数据
					this.data = this.testData;
					//数据先排序
					this.data.nodes.sort((a,b) => {
						return b.connections - a.connections
					});
					let maxDataArray = [],maxDataLength;
					//目前需求超过30个节点就不显示
					if(this.data.nodes.length > 30){
						maxDataLength = 30
					}else {
						maxDataLength = this.data.nodes.length
					}
					for (let i = 0 ; i < maxDataLength; i ++){
						maxDataArray.push(this.data.nodes[i])
					}
					
					this.data.nodes = maxDataArray;
					
					this.data.nodes.forEach((item,index) => {
						item.isDelete = false;
						item.color = this.color[index]
					});
					this.copyData = JSON.parse(JSON.stringify(this.data));
					this.colorUseCopyData = JSON.parse(JSON.stringify(this.data));
					this.initLegend();
					this.initChartsGraph();
				}).catch(e => {
					console.log(e);
					if(e){
						this.flShowNetwork = false
					}
				})
			},
			refreshPage(){
				window.location.reload();
			},
			selectAll(){
				this.flAllCheckout = !this.flAllCheckout;
				let flIsActive = this.colorDataArray.every( (item) =>{
					return item.isActive !== false
				});
				if(flIsActive){
					this.colorDataArray.forEach(item => {
						item.isActive = false
					});
					this.copyData.nodes = [];
					this.initChartsGraph();
				}else {
					this.colorDataArray.forEach(item => {
						item.isActive = true;
						this.copyData = JSON.parse(JSON.stringify(this.data));
						this.initChartsGraph();
					})
				}
			},
			initLegend(){
				this.colorDataArray = [];
				this.copyData.nodes.forEach( (item) => {
					this.colorDataArray.push({
						name: item['chain-id'],
						color: item.color,
						isActive: true
					})
				});
			},
			changeChart(value,index){
				this.colorDataArray[index].isActive = !this.colorDataArray[index].isActive;
				this.colorUseCopyData.nodes.forEach((item) => {
					if(item['chain-id'] === value.name){
						item.isDelete = !item.isDelete;
					}
				});
				this.copyData.nodes = this.colorUseCopyData.nodes.filter( item => {
					return item.isDelete === false
				});
				this.initChartsGraph();
			},
			initChartsGraph(){
				let graphEcharts = echarts.init(document.getElementById('validator_graph_content'));
				let nodeLinksArray = [],nodeArray = [];
				//最大像素点与最小像素点的差值106  最小的symbolSize 为 8 * 节点递增的比例
				let symbolSizeRule = 106;
				for(let i in this.copyData.nodes){
					nodeArray.push({
						name: this.copyData.nodes[i]['chain-id'],
						symbolSize: this.copyData.nodes[i].connections === 0 ? 8 : this.copyData.nodes[i].connections * symbolSizeRule / this.data.nodes.length ,
						label: {
							show: true,
							position:'right'
						},
						itemStyle: {
							color: this.copyData.nodes[i].color
						}
					});
				}
				
				let categories = [];
				this.copyData.nodes.forEach(item => {
					categories.push({
						name: item['chain-id']
					})
				})
				this.copyData.paths.forEach( (item,index) => {
					if(item.state === 'OPEN'){
						nodeLinksArray.push({
							source: item.src,
							target: item.dst,
							lineStyle:{
								color: '#70C6C7',
								curveness: 0.3,
							},
							emphasis: {
								lineStyle: {
									width: 1,
									type: 'solid',
								}
							}
						})
					}else {
						nodeLinksArray.push({
							source: item.src,
							target: item.dst,
							//连接线的样式设置
							lineStyle:{
								color: 'rgba(112, 198, 199, 1)',
								type: 'dashed',
								curveness: 0.3,
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
					name:'',
					tooltip:{
						formatter:function (data) {
							return `<span>${data.name.toString().replace('>','<>')}</span>`
						}
					},
					animationDuration: 1500,
					animationEasingUpdate: 'quinticInOut',
					series : [
						{
							name: '',
							type: 'graph',
							layout: 'force',
							data: nodeArray,
							links: nodeLinksArray,
							nodeScaleRatio: 0.6, //鼠标每次缩放的整体缩放比例
							force:{
								repulsion: [10,1500], //斥力因子
								gravity: 0.3, //是否向中心靠拢 值越大越接近于中心
								edgeLength: [10,300], //链接线的长度范围
								layoutAnimation: true
							},
							// zoom:0.9, //设置整体视图缩放的比例
							categories: categories, //类目数据
							roam: true, //平移和缩放  move 和 scale true 包含所有
							draggable: true, //是否拖拽
							hoverAnimation : true,
							focusNodeAdjacency: true,
							emphasis: {
								lineStyle: {
									width: 1,
								}
							}
						}
					]
				};
				graphEcharts.setOption(graphOption)
			},
		}
	}
</script>

<style scoped lang="scss">
	.validator_graph_container{
		width: 100%;
		height: 100%;
		.graph_content_wrap{
			max-width: 12.8rem;
			margin: 0 auto;
			.graph_content_title{
				height: 0.6rem;
				font-size: 0.18rem;
				color:rgba(255,255,255,1);
				display: flex;
				align-items: center;
				div:nth-of-type(1){
					margin-left: 0.2rem;
				}
				.beat_content{
					margin-left: 0.18rem;
					background: rgba(55, 124, 248, 1);
					color: #fff;
					font-size: 0.14rem;
					border-radius: 0.1rem;
					padding: 0.02rem 0.1rem;
				}
			}
			.tooltip{
				display: flex;
				align-items: center;
				margin-bottom: 0.1rem;
				p:nth-of-type(1){
					display: flex;
					align-items: center;
					span:nth-of-type(1){
						margin-left: 0.2rem;
						display: inline-block;
						width: 0.2rem;
						height: 0.02rem;
						border-top: 0.02rem solid #70C6C7;
					}
					span:nth-of-type(2){
						margin-left: 0.1rem;
						font-size: 0.14rem;
						color: rgba(115, 122, 174, 1);
					}
				}
				p:nth-of-type(2){
					display: flex;
					align-items: center;
					span:nth-of-type(1){
						width: 0.2rem;
						margin-left: 0.41rem;
						display: inline-block;
						border-top: 0.02rem dashed #70C6C7;
					}
					span:nth-of-type(2){
						margin-left: 0.1rem;
						font-size: 0.14rem;
						color: rgba(115, 122, 174, 1);
					}
				}
			}
			.graph_charts_container{
				background: #2D325A;
				margin-bottom: 0.6rem;
				min-height: 6.46rem;
				.graph_charts_title{
					padding: 0.2rem 0 0 0.2rem;
					font-size: 0.12rem;
					line-height: 0.14rem;
					color: rgba(115, 122, 174, 1);
				}
				.graph_content_container{
					display: flex;
					#validator_graph_content{
						max-width: 8rem;
						width: 100%;
						height: 6rem;
						background: #2D325A;
					}
					.graph_list_container{
						.graph_list_item_all{
							display: flex;
							justify-content: flex-start;
							margin:  0 0.6rem 0.2rem 0.7rem;
							.legend_all_block{
								width: 0.14rem;
								height: 0.14rem;
								border-radius: 0.07rem;
							}
							.legend_block{
								box-sizing: border-box;
								width: 0.14rem;
								height: 0.14rem;
								border-radius: 0.07rem;
							}
							.hide_style{
								color: rgba(134, 143, 211, 1) !important;
							}
							.legend_name{
								height: 0.14rem;
								color: rgba(134, 143, 211, 0.5);
								margin-left: 0.1rem;
							}
							
						}
						.graph_list_content{
							flex: 1;
							min-height: 0.14rem;
							display: flex;
							flex-wrap: wrap;
							margin-left: 0.2rem;
							.graph_list_item{
								width: 1.5rem;
								align-items: center;
								margin-left: 0.5rem;
								display: flex;
								padding: 0.05rem 0 0.1rem 0;
								cursor: pointer;
								.legend_all_block{
									width: 0.14rem;
									height: 0.14rem;
									border-radius: 0.07rem;
								}
								.legend_block{
									box-sizing: border-box;
									width: 0.14rem;
									height: 0.14rem;
									border-radius: 0.07rem;
								}
								.hide_style{
									background: transparent !important;
									border: 0.01rem solid rgba(134, 143, 211, 0.5);
								}
								.legend_name{
									height: 0.16rem;
									line-height: 0.16rem;
									color: rgba(134, 143, 211, 1);
									margin-left: 0.1rem;
								}
								.hide_style_color{
									color: rgba(134, 143, 211, 0.5) !important;
								}
							}
						}
					}
				}
			}
			.show_error_content{
				max-width: 12.8rem;
				min-height: 6rem;
				display: flex;
				align-items: center;
				justify-content: center;
				.error_content{
					text-align: center;
					p{
						max-width: 2.05rem;
						color: rgba(115, 122, 174, 1);
						font-size: 0.14rem;
						margin-bottom: 0.2rem;
					}
					span{
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
		.validator_graph_container{
			.graph_content_wrap{
				box-sizing: border-box;
				padding: 0 0.1rem;
				.tooltip{
					padding-top: 0.05rem;
					display: flex;
					flex-direction: column;
					align-items: flex-start;
					justify-content: center;
					p:nth-of-type(1){
						span:nth-of-type(1){
							margin-left: 0;
						}
					}
					p:nth-of-type(2){
						span:nth-of-type(1){
							margin-left: 0;
						}
					}
				}
				.graph_content_title{
					height: auto;
					padding: 0.15rem 0;
					align-items: flex-start;
					div:nth-of-type(1){
						margin-left: 0.1rem;
					}
				}
				.graph_charts_container{
					.tooltip{
						flex-direction: column;
						align-items: flex-start;
					}
					.graph_content_container{
						flex-direction: column;
						.graph_list_container{
							margin-bottom: 0.6rem;
							.graph_list_item_all{
								margin-left: 0.2rem;
								justify-content: flex-start;
							}
							.graph_list_content{
								height: auto;
								.graph_list_item{
									margin-left: 0;
								}
							}
						}
					}
				}
			}
		}
	}
</style>

