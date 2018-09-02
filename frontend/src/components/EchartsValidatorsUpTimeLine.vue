<template>
  <div :class="echartsComponentWrapLine">
    <div class="echarts_title_wrap_line">
      Uptime
    </div>
    <div id="echarts_uptime_line">

    </div>
  </div>


</template>

<script>
  import echarts from 'echarts';

  let line = null;
  export default {
    name: 'echarts-uptime-line',
    watch: {
      informationUptimeLine(informationUptimeLine) {
        //根据设备大小显示饼图的大小
        let radius = this.deviceType === 1 ? '85%' : '65%';
        let option = {
          tooltip : {
            trigger: 'item',
            formatter(params){
              let res =  `<span style="display:block;">${params.name}</span>`;
              res += `<span style="display:block;">Average Uptime: ${params.value}%</span>`;
              return res;
            }
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: [],
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },
            axisLabel:{
              interval:5,
              rotate:0,
              margin:12,
              formatter:(value)=>{
                if(value.split(":")[1] == "00"){
                  return value
                }else {
                  return `${this.month[this.monthNum.findIndex(item=>value.substr(0,2) === item)]}${value.substr(3,2)}`;
                }
              }
            }
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
              smooth:true,//曲线平滑
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
          option.xAxis.data = informationUptimeLine.xData;
          option.yAxis.max = informationUptimeLine.maxValue;
          option.series[0].data = informationUptimeLine.seriesData;
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
    props: ['informationUptimeLine'],
    beforeMount() {

    },
    mounted() {
      line = echarts.init(document.getElementById('echarts_uptime_line'));
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
<style lang="scss">
  @import '../style/mixin.scss';

  .echarts_component_wrap_line_personal_computer, .echarts_component_wrap_line_mobile {
    width: 100%;
    height: 100%;
    padding:0.12rem 0.2rem 0 0.2rem;

    .echarts_title_wrap_line {
      height: 15%;
      font-size:0.18rem;
      font-weight:600;
    }
    #echarts_uptime_line {
      width: 100%;
      height: 85%;
    }
  }
  .echarts_component_wrap_line_mobile{
    min-width:4rem;
  }


</style>
