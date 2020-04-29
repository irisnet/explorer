<template>
	<div class="san_key_container">
		<div class="san_key_content_wrap">
			<div id="san_key_content"></div>
		</div>
	</div>
</template>

<script>
	var echarts = require('echarts/lib/echarts');
	require('echarts/lib/component/legend');
	require('echarts/lib/component/tooltip');
	require('echarts/lib/component/title');
	require('echarts/lib/chart/sankey');
	require('echarts/extension/dataTool')
	require('echarts/lib/component/legendScroll');
	export default {
		name: "Sankey",
		data() {
			return {
			
			}
		},
		mounted () {
			this.getSankey()
		},
		methods:{
			getSankey(){
				let graphEcharts = echarts.init(document.getElementById('san_key_content'));
				let option = {
					
					color: [
						'#f4a582', '#d1e5f0', '#92c5de', '#4393c3', '#2166ac', '#053061'
					],
					tooltip: {
						trigger: 'item',
						position: 'top',
						extraCssText: 'transform:rotate(90deg)', //设置提示框的旋转角度
						formatter:function (data) {
							if(data.data.source && data.data.target){
								return `<div>${data.data.source} -- ${data.data.target} : ${data.data.value}</div>    `
							}else if(data.data.name){
								return `<div>${data.data.name} : ${Number(data.value).toFixed(2)} </div>`
							}
						}
					},
					animation: false,
					series: [
						{
							type: 'sankey',
							bottom: '10%',
							focusNodeAdjacency: 'allEdges',
							data: [
								{"name":"日报 2020.02.08","depth":3},
								{"name":"周报 2020.02.08","depth":3},
								{"name":"月报 2020.02","depth":3},
								{"name":"债券发行量","depth":1},
								{"name":"债券兑付量","depth":1},
								{"name":"债券交易总量","depth":1},
								{"name":"债券结算量","depth":1},
								{"name":"期末投资者数量","depth":1},
								{"name":"中国结算","depth":2},
								{"name":"中央结算","depth":2},
								{"name":"上清所","depth":2},
								{"name":"上交所","depth":2},
								{"name":"深交所","depth":2},
								{"name":"外汇交易中心","depth":2},
								
								
							],
							links: [
								{"target": "中国结算", "source": "债券发行量", "value": 0.16},
								{"target": "中央结算", "source": "债券发行量", "value": 0.16},
								{"target": "上清所", "source": "债券发行量", "value": 0.16},
								{"target": "上交所", "source": "债券发行量", "value": 0.16},
								{"target": "深交所", "source": "债券发行量", "value": 0.16},
								{"target": "外汇交易中心", "source": "债券发行量", "value": 0.16},

								{"target": "中国结算", "source": "债券兑付量", "value": 0.16},
								{"target": "中央结算", "source": "债券兑付量", "value": 0.16},
								{"target": "上清所", "source": "债券兑付量", "value": 0.16},
								{"target": "上交所", "source": "债券兑付量", "value": 0.16},
								{"target": "深交所", "source": "债券兑付量", "value": 0.16},
								{"target": "外汇交易中心", "source": "债券兑付量", "value": 0.16},

								{"target": "上交所", "source": "债券交易总量", "value": 0.16},
								{"target": "深交所", "source": "债券交易总量", "value": 0.16},
								{"target": "外汇交易中心", "source": "债券交易总量", "value": 0.16},

								{"target": "中国结算", "source": "债券结算量", "value": 0.16},
								{"target": "中央结算", "source": "债券结算量", "value": 0.16},
								{"target": "上清所", "source": "债券结算量", "value": 0.16},
								{"target": "外汇交易中心", "source": "债券结算量", "value": 0.16},

								{"target": "中国结算", "source": "期末投资者数量", "value": 0.16},
								{"target": "中央结算", "source": "期末投资者数量", "value": 0.16},
								{"target": "上清所", "source": "期末投资者数量", "value": 0.16},
								{"target": "上交所", "source": "期末投资者数量", "value": 0.16},
								{"target": "深交所", "source": "期末投资者数量", "value": 0.16},
								{"target": "外汇交易中心", "source": "期末投资者数量", "value": 0.16},

								{"target": "债券发行量", "source": "周报 2020.02.08", "value": 0.3},
								{"target": "债券兑付量", "source": "周报 2020.02.08", "value": 0.3},
								{"target": "债券交易总量", "source": "周报 2020.02.08", "value": 0.3},
								{"target": "债券结算量", "source": "周报 2020.02.08", "value": 0.3},
								{"target": "期末投资者数量", "source": "周报 2020.02.08", "value": 0.3},

								{"target": "债券发行量", "source": "月报 2020.02", "value": 0.3},
								{"target": "债券兑付量", "source": "月报 2020.02", "value": 0.3},
								{"target": "债券交易总量", "source": "月报 2020.02", "value": 0.3},
								{"target": "债券结算量", "source": "月报 2020.02", "value": 0.3},
								{"target": "期末投资者数量", "source": "月报 2020.02", "value": 0.3},

								{"target": "债券发行量", "source": "日报 2020.02.08", "value": 0.3},
								{"target": "债券兑付量", "source": "日报 2020.02.08", "value": 0.3},
								{"target": "债券交易总量", "source": "日报 2020.02.08", "value": 0.3},
								{"target": "债券结算量", "source": "日报 2020.02.08", "value": 0.3},
								{"target": "期末投资者数量", "source": "日报 2020.02.08", "value": 0.3}
							],
							layout:'none',
							orient: 'horizontal',
							label: {
								position: [-10,10],
								rotate: -90
							},
							animation: true,
							lineStyle: {
								color: 'source',
								curveness: 0.5
							}
						}
					]
				};
				graphEcharts.setOption(option)
			}
		}
	}
</script>

<style scoped lang="scss">
	.san_key_container{
		width: 100%;
		.san_key_content_wrap{
			max-width: 12.8rem;
			margin: 0 auto;
			#san_key_content{
				margin: 0 auto;
				max-width: 8rem;
				height: 8rem;
				background: #fefefe;
				transform: rotate(-90deg);
			}
		}
	}

</style>
