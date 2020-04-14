<template>
	<div class="validator_detail_scatter_container">
		<div class="validator_detail_scatter_content_wrap">
			<div id="validator_detail_scatter_content"></div>
		</div>
		<div class="validator_chart_tip">
			<p class="current_chart_tip">
				<span class="current_chart_color_block" :style="{background:tipColor}"></span>
				<span class="current_moniker">{{currentMoniker}}</span>
			</p>
			<p class="other_chart_tip">
				<span class="other_chart_color_block"></span>
				<span>Other Validator</span>
			</p>
		</div>
	</div>
</template>

<script>
	import Service from "../../../service"
	import Tools from "../../../util/Tools"
	import Constant from "../../../constant/Constant"
	import bigNumber from "bignumber.js"
	var echarts = require('echarts/lib/echarts');
	require('echarts/lib/component/legend');
	require('echarts/lib/component/tooltip');
	require('echarts/lib/component/title');
	require('echarts/lib/chart/scatter');
	require('echarts/lib/component/legendScroll');
	export default {
		name: "ValidatorDetailScatter",
		props:{
			jailedData:{
				type:Object
			},
			validatorStatus:{
				type:String
			}
		},
		data () {
			return {
				otherValidatorBondedTokenArr:[],
				currentValidatorBondedTokenArr:[],
				mainnetThemeStyle:["#3264FD","#FF5C01"],
				testnetFuXiThemeStyle:["#0C4282","#FF5C01"],
				testnetNyancatThemeStyle:["#0D9388","#FF5C01"],
				defaultThemeStyle:["#0580D3","#FF5C01"],
				defaultJailedThemeStyle:["#0580D3","#101149"],
				mainnetJailedThemeStyle:['#3264FD',"#101149"],
				testnetFuXiJailedThemeStyle:['#0C4282',"#101149"],
				testnetNyancatJailedThemeStyle:['#0D9388',"#101149"],
				chartOptionColor:'',
				currentMoniker:'',
				tipColor:''
			}
		},
		watch:{
			validatorStatus(){
				if(this.validatorStatus === 'Jailed' || this.validatorStatus === 'Candidate'){
					this.tipColor = "#101149"
				}else {
					this.tipColor = "#FF5C01"
				}
			}
		},
		mounted () {
			let getJailedData = new Promise((resolve,reject)=>{
				resolve (this.jailedData);
				reject('props data is empty')
			});
			getJailedData.then((data) => {
				if(data){
					this.getValidatorCommissionInfo()
				}
			}).catch(e => {
				console.error(e)
			});
			setTimeout(() => {
				this.getValidatorCommissionInfo()
			},300)
		},
		methods:{
			initCharts(){
				let itemStyle = {
					opacity: 0.5,
					shadowOffsetX: 0,
					shadowOffsetY: 0,
				};
				let echartsData = echarts.init(document.getElementById('validator_detail_scatter_content'));
				let echartsOption = {
					backgroundColor: '#fff',
					legendHoverLink:true,
					emphasis:{
						itemStyle:{
							opacity: 1,
						}
					},
					grid: {
						left: '19%',
						right: 150,
						top: '18%',
						bottom: '10%'
					},
					tooltip: {
						padding: 10,
						backgroundColor: '#222',
						borderColor: '#777',
						borderWidth: 1,
						formatter: function (obj) {
							let value = obj.value;
							return `<div>
										<p>${value[2]}</p>
										<p>Commission Rate:${value[0]}%</p>
										<p>Bonded Tokens:<br/>${new bigNumber(value[1]).toFormat()} IRIS</p>
										</div>`
						}
					},
					xAxis: {
						type: 'value',
						name:'Commission Rate(%)',
						nameLocation:'end',
						nameGap: 6,
						nameTextStyle: {
							color: '#787C99',
							fontSize: 14
						},
						splitLine: {
							show: false
						},
						axisLine: {
							lineStyle: {
								color: '#a2a2ae'
							}
						}
					},
					yAxis: {
						type: 'value',
						name:'Bonded_Token(IRIS)',
						nameLocation: 'end',
						nameGap: 20,
						nameTextStyle: {
							color: '#787C99',
							fontSize: 14
						},
						axisLine: {
							lineStyle: {
								color: '#a2a2ae'
							}
						},
						splitLine: {
							show: false
						}
					},
					series: [
						{
							type: 'scatter',
							itemStyle: itemStyle,
							data: this.otherValidatorBondedTokenArr
						},
						{
							type: 'scatter',
							itemStyle: itemStyle,
							data: this.currentValidatorBondedTokenArr
						}
					]
				};
				if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.IRISHUB){
					if(this.validatorStatus === 'Jailed' || this.validatorStatus === 'Candidate'){
						this.chartOptionColor = this.mainnetJailedThemeStyle;
					}else {
						this.chartOptionColor = this.mainnetThemeStyle;
					}
				}else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.FUXI){
					if(this.validatorStatus === 'Jailed' || this.validatorStatus === 'Candidate'){
						this.chartOptionColor = this.testnetFuXiJailedThemeStyle;
					}else {
						this.chartOptionColor = this.testnetFuXiThemeStyle;
					}
				}else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.NYANCAT){
					if(this.validatorStatus === 'Jailed' || this.validatorStatus === 'Candidate'){
						this.chartOptionColor = this.testnetNyancatJailedThemeStyle;
					}else {
						this.chartOptionColor = this.testnetNyancatThemeStyle;
					}
					
				}else {
					if(this.validatorStatus === 'Jailed' || this.validatorStatus === 'Candidate'){
						this.chartOptionColor = this.defaultJailedThemeStyle;
					}else {
						this.chartOptionColor = this.defaultThemeStyle;
					}
					
				}
				echartsOption.color = this.chartOptionColor
				echartsData.setOption(echartsOption)
			},
			getValidatorCommissionInfo(){
				Service.commonInterface({validatorScatterList:{
						validatorAddr:this.$route.params.param
					}},(res) => {
					try {
						if(res && res.commission_data && res.commission_data.length > 0){
							let copyData = JSON.parse(JSON.stringify(this.jailedData));
							res.commission_data.push(copyData);
							let allValidatorBondedTokenArr = res.commission_data;
							this.currentValidatorBondedTokenArr = [];
							this.otherValidatorBondedTokenArr = [];
							allValidatorBondedTokenArr.forEach( item => {
								item.formatCommissionRote = (Number(item.commission_rate) * 100).toFixed(0);
								item.formatBondedToken = Tools.subStrings(item.bonded_tokens,6)
								if(item.operator_address === this.$route.params.param){
									this.currentMoniker = item.moniker;
									this.currentValidatorBondedTokenArr.push([item.formatCommissionRote,item.formatBondedToken,item.moniker]);
								}else {
									this.otherValidatorBondedTokenArr.push([item.formatCommissionRote,item.formatBondedToken,item.moniker]);
									
								}
							});
							this.initCharts();
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
	.validator_detail_scatter_container{
		min-height: 2.2rem;
		width: 100%;
		.validator_detail_scatter_content_wrap{
			width: 100%;
			#validator_detail_scatter_content{
				height: 2.2rem;
				margin-right: 0.7rem;
				margin-top: 0.2rem;
				min-width: 5rem;
			}
		}
		.validator_chart_tip{
			display: flex;
			padding-left: 0.2rem;
			margin-top: 0.2rem;
			.current_chart_tip{
				display: flex;
				align-items: center;
				.current_chart_color_block{
					display: inline-block;
					width: 0.14rem;
					height: 0.14rem;
					border-radius: 0.07rem;
					opacity: 0.5;
				}
				.current_moniker{
					margin-left: 0.05rem;
				}
			}
			.other_chart_tip{
				margin-left: 0.4rem;
				display: flex;
				align-items: center;
				.other_chart_color_block{
					margin-right: 0.05rem;
					display: inline-block;
					width: 0.14rem;
					height: 0.14rem;
					border-radius: 0.07rem;
					opacity: 0.5;
					background: var(--bgColor);
				}
			}
		}
	}
	@media screen and (max-width: 750px){
		.validator_detail_scatter_container{
			width: 100%;
			.validator_detail_scatter_content_wrap{
				width: 100%;
				overflow-x: auto;
				#validator_detail_scatter_content{
					margin-right: 0;
				}
			}
		}
	}
</style>
