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
  import Tools from "../common/Tools";

  let line = null;
  export default {
    name: 'echarts-validators-line',
    watch: {
      informationValidatorsLine(informationValidatorsLine) {
        let radius = this.deviceType === 1 ? '85%' : '65%';
        let option = {
          tooltip : {
            trigger: 'item',
            formatter(params){
              params.value[0] = Tools.formatDateYearAndMinutesAndSeconds(params.value[0]);
              let res =  `<span style="display:block;">${params.value[0].substr(5,11)} +UTC</span>`;
              res += `<span style="display:block;">Voting Power:${params.value[1]}</span>`;
              return res;
            }
          },
          xAxis: {
            axisLabel:{
              show:true,
              formatter:(value)=>{
                value = Tools.formatDateYearToDate(value).substr(5,6);
                return `${this.month[this.monthNum.findIndex(item=>value.substr(0,2) === item)]}${value.substr(3,2)}`;
              }
            },
            type: 'time',
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },

            splitLine:{
              show:false
            },
            // splitNumber:0,
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
