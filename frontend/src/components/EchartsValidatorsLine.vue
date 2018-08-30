<template>
  <div :class="echartsComponentWrapLine">
    <div class="echarts_title_wrap_line">
      Voting Power
    </div>
    <div id="echarts_line">

    </div>
  </div>


</template>

<script>
  import echarts from 'echarts';

  let line = null;
  export default {
    name: 'echarts-validators-line',
    watch: {
      informationValidatorsLine(informationValidatorsLine) {
        //根据设备大小显示饼图的大小
        let radius = this.deviceType === 1 ? '85%' : '65%';
        let option = {
          tooltip : {
            trigger: 'item',
            formatter(params){
              console.log(params,'线图数据');
              let res =  `<span style="display:block;">${params.name}</span>`;
              res += `<span style="display:block;">Transactions: ${params.value}</span>`;
              return res;
            }
          },
          xAxis: {
            type: 'time',
            boundaryGap: ["0%","0%"],
            realtime: false,
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },
            splitNumber:1,
          },
          yAxis: {
            type: 'value',
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },
            min:0,
            max:0,
            splitNumber:5,
            splitLine: {
              lineStyle: {
                // 使用深浅的间隔色
                color: ['#eee']
              }
            }
          },
          grid: {
            left: '3%',   //图表距边框的距离
            right: '5%',
            bottom: '3%',
            top:'5%',
            containLabel: true
          },
          series: [
            {
              data: [],
              type: 'line',
              step:"end",
              areaStyle: {
                normal: {
                  color: new echarts.graphic.LinearGradient(//设置渐变颜色
                    0, 0, 0, 1,
                    [
                      {offset: 0, color: '#3498db'},
                      {offset: 0.5, color: '#91ccef'},
                      {offset: 1, color: '#dcf6ff'}
                    ]
                  )
                }
              },
              smooth:false,//曲线平滑
              itemStyle:{
                normal:{
                  color:'#3598db',
                  borderColor:'#3598db',  //拐点边框颜色
                }
              },
            }
          ]
        };
        if (line) {
          let test = ["08/15/2018", "08/21/2018", "08/22/2018", "08/22/2018", "08/24/2018", "08/24/2018", "08/24/2018",
            "08/24/2018", "08/24/2018", "08/25/2018", "08/25/2018", "08/25/2018", "08/25/2018", "08/25/2018", "08/25/2018",
            "08/25/2018", "08/25/2018", "08/25/2018", "08/27/2018", "08/27/2018", "08/27/2018", "08/27/2018",
            "08/27/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018",
            "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018",
            "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018",
            "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018", "08/28/2018",
            "08/29/2018",];
          let testa = [0, 100, 150, 110, 112, 114, 116, 150, 152, 155, 187, 219, 220, 227, 229, 233, 229, 233, 236, 239, 240, 242, 248, 250, 251, 252, 250, 251, 252, 254,
            259, 260, 261, 262, 263, 264, 265, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 279];
          var data = [
            ["2018-08-21 02:54:20",
              2,
              "翻身"
            ],
            [
              '2018-08-22 02:54:20',
              0.5,
              "没翻身"
            ],
            ];
          console.log(informationValidatorsLine.seriesData);
          option.yAxis.max = informationValidatorsLine.maxValue;
          option.series[0].data = informationValidatorsLine.seriesData;
          line.setOption(option)
        }
      }
    },
    data() {
      return {
        deviceType: window.innerWidth > 500 ? 1 : 0,
        echartsComponentWrapLine:window.innerWidth > 500 ?'echarts_component_wrap_line_personal_computer':'echarts_component_wrap_line_mobile',
        month:['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'],
        monthNum:['01','02','03','04','05','06','07','08','09','10','11','12'],

      }
    },
    props: ['informationValidatorsLine'],
    beforeMount() {

    },
    mounted() {
      line = echarts.init(document.getElementById('echarts_line'));
      window.addEventListener('resize',this.onWindowResize)
    },

    methods: {
      onWindowResize(){
        line.resize();
      }
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.onWindowResize);
    }
  }
</script>
<style scoped lang="scss">
  @import '../style/mixin.scss';

  .echarts_component_wrap_line_personal_computer, .echarts_component_wrap_line_mobile {
    width: 100%;
    height: 100%;
    padding: 0!important;
    margin: 0!important;
    .echarts_title_wrap_line {
      height: 15%;
      font-size:0.18rem;
      font-weight:600;
      line-height: 0.53rem;
      padding-left: 0.2rem;
    }
    #echarts_line {
      width: 100%;
      height: 85%;
    }
  }
  .echarts_component_wrap_line_mobile{
    min-width:4rem;
  }


</style>
