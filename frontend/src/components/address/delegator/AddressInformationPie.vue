<template>
	<div id="address_information_chart"></div>

</template>

<script>
	import Constant from "../../../constant/Constant"
	var echarts = require('echarts/lib/echarts')
	require('echarts/lib/component/legend')
	require('echarts/lib/component/tooltip')
	require('echarts/lib/component/title')
	require('echarts/lib/chart/pie');
	require('echarts/lib/component/legendScroll')
	let pie = null;
	export default {
		name: "AddressInformationPie",
		data (){
			return {
				addressInformationCharts:null,
				themeStyleArray:'',
				mainnetThemeStyle:["#3264FD","#FFA300","#67E523","#8E66FF"],
				testnetFuXiThemeStyle:["#0C4282","#FFA300","#67E523","#8E66FF"],
				testnetNyancatThemeStyle:["#0D9388","#FFA300","#67E523","#8E66FF"],
				defaultThemeStyle:["#0580D3","#FFA300","#67E523","#8E66FF"],
			}
		},
		watch:{
			echartData(){
				this.initCharts();
			}
		},
		props:{
			echartData:{
				type: Array
			}
		},
		mounted () {
			setTimeout(() => {
				this.initCharts();
			},400)
		},
		methods:{
			initCharts(){
				this.addressInformationCharts = echarts.init(document.getElementById('address_information_chart'));
				let echartsOption = {
					tooltip: {
						trigger: 'item',
						formatter: function (data) {
							console.log(data)
							return `${data.name}: ${data.value} IRIS (${data.percent}%)`
						}
					},
					legend: {
						orient: 'vertical',
						left: 10,
						data: []
					},
					series : [
						{
							hoverAnimation :false,
							type: 'pie',
							radius: ['50%', '70%'],
							avoidLabelOverlap: false,
							label: {
								show: false,
								position: 'center'
							},
						
							emphasis: {
								label: {
									show: false,
									fontSize: '30',
									fontWeight: 'bold'
								}
							},
							labelLine: {
								show: false
							},
						}
					]
				};
				if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.IRISHUB){
					this.themeStyleArray = this.mainnetThemeStyle;
				}else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.FUXI){
					this.themeStyleArray = this.testnetFuXiThemeStyle;
				}else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.NYANCAT){
					this.themeStyleArray = this.testnetNyancatThemeStyle;
				}else {
					this.themeStyleArray = this.defaultThemeStyle;
				}
				let seriesData = this.echartData.map( (item,index )=> {
					return {
						value: item.numberValue,
						name: item.label,
						itemStyle:{
							color: this.themeStyleArray[index]
						}
					}
				});
				echartsOption.series[0].data = seriesData;
				
				this.addressInformationCharts.setOption(echartsOption)
			},
			
		}
	}
</script>

<style scoped lang="scss">
	#address_information_chart{
		height: 1rem;
	}
</style>
