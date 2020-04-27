<template>
	<div class="test_chart_container">
		<div class="test_chart_content_wrap">
			<div id="test_chart_content"></div>
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
		name: "TestChart",
		data() {
			return {
			
			}
		},
		mounted(){
			this.testInitChart();
		},
		methods:{
			testInitChart() {
				let graphEcharts = echarts.init(document.getElementById('test_chart_content'));
				// 球型关系图
				let option = {
					title: {
						text: '关系图的层次'
					},
					tooltip: {},
					animationDurationUpdate: 1500,
					animationEasingUpdate: 'quinticInOut',
					series: [
						{
							type: 'graph',//图形类型
							layout: 'none',
							
							symbolSize: 20,//控制圆的大小
							roam: true,
							label: { //是否显示圆圈里面的文字
								show: true,
								formatter: function (params) {
									return params.name;
								},
								fontSize: 10,
								fontStyle: 'normal',//'normal italic oblique'文字样式
								fontWeight: 'normal',//文字加粗
								position: 'insideRight',
								distance: -20,//距离圆圈的距离
								color: '#red'//字体颜色
							},
							edgeSymbol: ['circle', 'arrow'],//全局线的箭头
							edgeSymbolSize: [8, 8],//全局线箭头大小
							edgeLabel: {
								fontSize: 10//全局线上字体大小
							},
							data: [{//名称节点坐标
								name: 'A3',
								x: 10,
								y: 50,
								symbol:'arrow',
							}, {
								name: 'F5',
								x: 42,
								y: 52,
								symbol:'diamond',
							}, {
								name: 'F4A',
								x: 72,
								y: 54,
								symbol:'circle',
							}, {
								name: 'G4',
								x: 72,
								y: 57,
								symbol:'rect',
							}, {
								name: 'G5',
								x: 42,
								y: 60,
								symbol:'roundRect',
							}, {
								name: 'A4',
								x: 10,
								y: 65,
								symbol:'triangle',
							}, {
								name: 'F3',
								x: 42,
								y: 67,
								symbol:'pin',
							}, {
								name: 'E4',
								x: 72,
								y: 69
							}, {
								name: 'E1',
								x: 10,
								y: 77
							}, {
								name: 'E2',
								x: 10,
								y: 80
							}, {
								name: 'E8',
								x: 42,
								y: 82
							}, {
								name: 'F1',
								x: 42,
								y: 85
							}, {
								name: 'E3',
								x: 10,
								y: 89
							}, {
								name: 'E5',
								x: 10,
								y: 92
							}, {
								name: 'E6',
								x: 42,
								y: 96
							}, {
								name: 'E7',
								x: 72,
								y: 99
							}, {
								name: 'F4B',
								x: 42,
								y: 106
							}, {
								name: 'G1',
								x: 42,
								y: 109
							}, {
								name: 'G2',
								x: 42,
								y: 112
							}, {
								name: 'G3',
								x: 42,
								y: 115
							}, {
								name: 'F2',
								x: 10,
								y: 119
							}, {
								name: 'F6',
								x: 10,
								y: 122
							}, {
								name: 'G6',
								x: 10,
								y: 125
							}],
							links: [{//线的链接
								source: 'A3',//源头
								target: 'F5',//指向源终
								//symbolSize: [5, 20], 个性线
								label: {//是否显示线上的文字
									show: false,
									formatter: 'A3'//显示线上的文字
								}
								// lineStyle: {//设置线的大小、样式等
								// 	width: 5,
								// 	curveness: 0.2
								// }
							}, {
								source: 'A3',
								target: 'G5',
								label: {
									show: false
								}
							}, {
								source: 'F5',
								target: 'F4A',
								label: {
									show: false
								}
							}, {
								source: 'F5',
								target: 'G4',
								label: {
									show: false
								}
							}, {
								source: 'G5',
								target: 'G4',
								label: {
									show: false
								}
							}, {
								source: 'A4',
								target: 'F3',
								label: {
									show: false
								}
							}, {
								source: 'F3',
								target: 'F4A',
								label: {
									show: false
								}
							}, {
								source: 'F3',
								target: 'E4',
								label: {
									show: false
								}
							}, {
								source: 'E2',
								target: 'E8',
								label: {
									show: false
								}
							}, {
								source: 'E2',
								target: 'F1',
								label: {
									show: false
								}
							}, {
								source: 'E8',
								target: 'E4',
								label: {
									show: false
								}
							}, {
								source: 'E5',
								target: 'E6',
								label: {
									show: false
								}
							}, {
								source: 'E5',
								target: 'F4B',
								label: {
									show: false
								}
							}, {
								source: 'E5',
								target: 'G1',
								label: {
									show: false
								}
							}, {
								source: 'E5',
								target: 'G2',
								label: {
									show: false
								}
							}, {
								source: 'E5',
								target: 'G3',
								label: {
									show: false
								}
							}, {
								source: 'E6',
								target: 'E7',
								label: {
									show: false
								}
							},
								{
									source: 'G1',
									target: 'G2',
									label: {
										show: false
									}
								}],
							lineStyle: {//如果links里面没有单独设置线的特色，默认取这里
								opacity: 0.9,
								width: 2,
								curveness: 0
							}
						}
					]
				};
				
				//桑基图
				/*let option = {
					title: {
						text: '桑基图'//标题
					},
					//颜色 color: [
					// 	'#67001f', '#b2182b', '#d6604d', '#f4a582', '#fddbc7', '#d1e5f0', '#92c5de', '#4393c3', '#2166ac', '#053061'
					// ],
					tooltip: {
						trigger: 'item',
							triggerOn: 'mousemove'//鼠标移动
					},
					// animation: false,
					series: [
						{
							type: 'sankey',//桑基图类型
							focusNodeAdjacency: 'allEdges',
							itemStyle: {//控制边框样式
								borderWidth: 1,
								borderColor: '#aaa'
							},
							data: [//名称
								{name: 'Total'},
								{name: 'Environment'},
								{name: 'Land use'},
								{name: 'Cocoa butter (Organic)'},
								{name: 'Cocoa mass (Organic)'},
								{name: 'Hazelnuts (Organic)'},
								{name: 'Cane sugar (Organic)'},
								{name: 'Vegetables (Organic)'},
								{name: 'Climate change'},
								{name: 'Harmful substances'},
								{name: 'Water use'},
								{name: 'Resource depletion'},
								{name: 'Refrigeration'},
								{name: 'Packaging'},
								{name: 'Human rights'},
								{name: 'Child labour'},
								{name: 'Coconut oil (Organic)'},
								{name: 'Forced labour'},
								{name: 'Health safety'},
								{name: 'Access to water'},
								{name: 'Freedom of association'},
								{name: 'Access to land'},
								{name: 'Sufficient wage'},
								{name: 'Equal rights migrants'},
								{name: 'Discrimination'},
								{name: 'Working hours'}
							],
							links: [//对应关系
								{source:'Total', target:'Environment',value: 0.342284047256003},
								{source:'Environment', target:'Land use',value: 0.32322870366987},
								{source:'Land use', target:'Cocoa butter (Organic)',value: 0.177682517071359},
								{source:'Land use', target:'Cocoa mass (Organic)',value: 0.137241325342711},
								{source:'Land use', target:'Hazelnuts (Organic)',value: 0.00433076373512774},
								{source:'Land use', target:'Cane sugar (Organic)',value: 0.00296956039863467},
								{source:'Land use', target:'Vegetables (Organic)',value: 0.00100453712203756},
								{source:'Environment', target:'Climate change',value: 0.0112886157414413},
								{source:'Climate change', target:'Cocoa butter (Organic)',value: 0.00676852971933996},
								{source:'Climate change', target:'Cocoa mass (Organic)',value: 0.00394686874786743},
								{source:'Climate change', target:'Cane sugar (Organic)',value: 0.000315972058711838},
								{source:'Climate change', target:'Hazelnuts (Organic)',value: 0.000218969462265292},
								{source:'Climate change', target:'Vegetables (Organic)',value: 3.82757532567656e-05},
								{source:'Environment', target:'Harmful substances',value: 0.00604275542495656},
								{source:'Harmful substances', target:'Cocoa mass (Organic)',value: 0.0055125989240741},
								{source:'Harmful substances', target:'Cocoa butter (Organic)',value: 0.000330017607892127},
								{source:'Harmful substances', target:'Cane sugar (Organic)',value: 0.000200138892990337},
								{source:'Harmful substances', target:'Hazelnuts (Organic)',value: 0},
								{source:'Harmful substances', target:'Vegetables (Organic)',value: 0},
								{source:'Environment', target:'Water use',value: 0.00148345269044703},
								{source:'Water use', target:'Cocoa butter (Organic)',value: 0.00135309891304186},
								{source:'Water use', target:'Cocoa mass (Organic)',value: 0.000105714137908639},
								{source:'Water use', target:'Hazelnuts (Organic)',value: 1.33452642581887e-05},
								{source:'Water use', target:'Cane sugar (Organic)',value: 8.78074837009238e-06},
								{source:'Water use', target:'Vegetables (Organic)',value: 2.5136268682477e-06},
								{source:'Environment', target:'Resource depletion',value: 0.000240519729288764},
								{source:'Resource depletion', target:'Cane sugar (Organic)',value: 0.000226237279345084},
								{source:'Resource depletion', target:'Vegetables (Organic)',value: 1.42824499436793e-05},
								{source:'Resource depletion', target:'Hazelnuts (Organic)',value: 0},
								{source:'Resource depletion', target:'Cocoa mass (Organic)',value: 0},
								{source:'Resource depletion', target:'Cocoa butter (Organic)',value: 0},
								{source:'Environment', target:'Refrigeration',value: 0},
								{source:'Environment', target:'Packaging',value: 0},
								{source:'Total', target:'Human rights',value: 0.307574096993239},
								{source:'Human rights', target:'Child labour',value: 0.0410641202645833},
								{source:'Child labour', target:'Hazelnuts (Organic)',value: 0.0105339381639722},
								{source:'Child labour', target:'Cocoa mass (Organic)',value: 0.0105},
								{source:'Child labour', target:'Cocoa butter (Organic)',value: 0.0087294420777},
								{source:'Child labour', target:'Coconut oil (Organic)',value: 0.00474399974233333},
								{source:'Child labour', target:'Cane sugar (Organic)',value: 0.00388226450884445},
								{source:'Child labour', target:'Vegetables (Organic)',value: 0.00267447577173333},
								{source:'Human rights', target:'Forced labour',value: 0.0365458590642445},
								{source:'Forced labour', target:'Hazelnuts (Organic)',value: 0.0114913076376389},
								{source:'Forced labour', target:'Cocoa butter (Organic)',value: 0.0081134807347},
								{source:'Forced labour', target:'Cocoa mass (Organic)',value: 0.00765230236575},
								{source:'Forced labour', target:'Cane sugar (Organic)',value: 0.004},
								{source:'Forced labour', target:'Vegetables (Organic)',value: 0.00296668823626667},
								{source:'Forced labour', target:'Coconut oil (Organic)',value: 0.00232208008988889},
								{source:'Human rights', target:'Health safety',value: 0.0345435327843611},
								{source:'Health safety', target:'Hazelnuts (Organic)',value: 0.0121419536385},
								{source:'Health safety', target:'Cocoa mass (Organic)',value: 0.00766772850678333},
								{source:'Health safety', target:'Cocoa butter (Organic)',value: 0.0056245892061},
								{source:'Health safety', target:'Coconut oil (Organic)',value: 0.00361616847688889},
								{source:'Health safety', target:'Cane sugar (Organic)',value: 0.00277374682533333},
								{source:'Health safety', target:'Vegetables (Organic)',value: 0.00271934613075556},
								{source:'Human rights', target:'Access to water',value: 0.0340206659360667},
								{source:'Access to water', target:'Cocoa mass (Organic)',value: 0.0105},
								{source:'Access to water', target:'Cocoa butter (Organic)',value: 0.0089274160792},
								{source:'Access to water', target:'Hazelnuts (Organic)',value: 0.0054148022845},
								{source:'Access to water', target:'Cane sugar (Organic)',value: 0.00333938149786667},
								{source:'Access to water', target:'Vegetables (Organic)',value: 0.00314663377488889},
								{source:'Access to water', target:'Coconut oil (Organic)',value: 0.00269243229961111},
								{source:'Human rights', target:'Freedom of association',value: 0.0320571523941667},
								{source:'Freedom of association', target:'Hazelnuts (Organic)',value: 0.0132312483463611},
								{source:'Freedom of association', target:'Cocoa butter (Organic)',value: 0.0077695200707},
								{source:'Freedom of association', target:'Cocoa mass (Organic)',value: 0.00510606573995},
								{source:'Freedom of association', target:'Vegetables (Organic)',value: 0.00354321156324444},
								{source:'Freedom of association', target:'Cane sugar (Organic)',value: 0.00240710667391111},
								{source:'Freedom of association', target:'Coconut oil (Organic)',value: 0},
								{source:'Human rights', target:'Access to land',value: 0.0315022209894056},
								{source:'Access to land', target:'Hazelnuts (Organic)',value: 0.00964970063322223},
								{source:'Access to land', target:'Cocoa mass (Organic)',value: 0.00938530207965},
								{source:'Access to land', target:'Cocoa butter (Organic)',value: 0.0060110791848},
								{source:'Access to land', target:'Cane sugar (Organic)',value: 0.00380818314608889},
								{source:'Access to land', target:'Vegetables (Organic)',value: 0.00264795594564445},
								{source:'Access to land', target:'Coconut oil (Organic)',value: 0},
								{source:'Human rights', target:'Sufficient wage',value: 0.0287776757227333},
								{source:'Sufficient wage', target:'Cocoa mass (Organic)',value: 0.00883512456493333},
								{source:'Sufficient wage', target:'Cocoa butter (Organic)',value: 0.0078343367268},
								{source:'Sufficient wage', target:'Coconut oil (Organic)',value: 0.00347879026511111},
								{source:'Sufficient wage', target:'Hazelnuts (Organic)',value: 0.00316254211388889},
								{source:'Sufficient wage', target:'Vegetables (Organic)',value: 0.00281013722808889},
								{source:'Sufficient wage', target:'Cane sugar (Organic)',value: 0.00265674482391111},
								{source:'Human rights', target:'Equal rights migrants',value: 0.0271146645119444},
								{source:'Equal rights migrants', target:'Cocoa butter (Organic)',value: 0.0071042315061},
								{source:'Equal rights migrants', target:'Cocoa mass (Organic)',value: 0.00636673210005},
								{source:'Equal rights migrants', target:'Hazelnuts (Organic)',value: 0.00601459775836111},
								{source:'Equal rights migrants', target:'Coconut oil (Organic)',value: 0.00429185583138889},
								{source:'Equal rights migrants', target:'Cane sugar (Organic)',value: 0.00182647471915556},
								{source:'Equal rights migrants', target:'Vegetables (Organic)',value: 0.00151077259688889},
								{source:'Human rights', target:'Discrimination',value: 0.0211217763359833},
								{source:'Discrimination', target:'Cocoa mass (Organic)',value: 0.00609671700306667},
								{source:'Discrimination', target:'Cocoa butter (Organic)',value: 0.0047738806325},
								{source:'Discrimination', target:'Coconut oil (Organic)',value: 0.00368119084494444},
								{source:'Discrimination', target:'Vegetables (Organic)',value: 0.00286009813604444},
								{source:'Discrimination', target:'Cane sugar (Organic)',value: 0.00283706180951111},
								{source:'Discrimination', target:'Hazelnuts (Organic)',value: 0.000872827909916666},
								{source:'Human rights', target:'Working hours',value: 0.02082642898975},
								{source:'Working hours', target:'Hazelnuts (Organic)',value: 0.0107216773848333},
								{source:'Working hours', target:'Coconut oil (Organic)',value: 0.00359009052944444},
								{source:'Working hours', target:'Vegetables (Organic)',value: 0.00212300379075556},
								{source:'Working hours', target:'Cocoa butter (Organic)',value: 0.0018518584356},
								{source:'Working hours', target:'Cocoa mass (Organic)',value: 0.00158227069058333},
								{source:'Working hours', target:'Cane sugar (Organic)',value: 0.000957528158533333}
							],
							lineStyle: {//控制线条样式
								color: 'source',
								curveness: 0.5
							}
						}
					]
				};*/

				
				graphEcharts.setOption(option)
			}
		}
	}
</script>

<style scoped lang="scss">
	.test_chart_container{
		width: 100%;
		.test_chart_content_wrap{
			max-width: 12.8rem;
			margin: 0 auto;
			#test_chart_content{
				margin: 0 auto;
				max-width: 12.8rem;
				height: 8rem;
				background: #fefefe;
			}
		}
	}
</style>
