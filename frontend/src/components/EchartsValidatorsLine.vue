<template>
  <div :class="echartsComponentWrapLine">
    <div class="echarts_title_wrap_line">
      Voting Power
    </div>
    <div id="echarts_validator_line">

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
        immediate:true,
        deep: true,
        informationValidatorsLine:function () {
          let radius = this.deviceType === 1 ? '85%' : '65%';
          this.configuration();
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
    created(){
    },
    mounted() {
      line = echarts.init(document.getElementById('echarts_validator_line'));
      window.addEventListener('resize',this.onWindowResize);
      //为解决有时候watch不触发导致图表展示不出来的问题，如果没有触发0.5秒再次渲染图表
      let that = this;
      setTimeout(function () {
        that.configuration()
      },500)
    },

    methods: {
      onWindowResize(){
        line.resize();
      },
      configuration(){
        let option = {
          tooltip: {
            trigger: 'axis',
            formatter(params) {
              params[0].value[0] = Tools.formatDateYearAndMinutesAndSeconds(params[0].value[0]);
              let res = `<span style="display:block;">${params[0].value[0].substr(5, 11)} +UTC</span>`;
              res += `<span style="display:block;">Voting Power:${params[0].value[1]}</span>`;
              return res;
            },
            axisPointer: {
              show: true,
              axis: "x",
              type: "line",
              lineStyle: {
                color: "#a2a2ae",
              },
            },
          },
          xAxis: {
            axisLabel: {
              color: "#000",
              show: true,
              rotate: 45,
              interval: '0',
              formatter: (value) => {
                value = Tools.formatDateYearToDate(value).substr(5, 6);
                return `${this.month[this.monthNum.findIndex(item => value.substr(0, 2) === item)]}${value.substr(3, 2)}`;
              }
            },
            type: 'time',
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },

            splitLine: {
              show: false
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
            axisLabel: {
              color: "#000",
            },
            min: 0,
            splitNumber: 5,
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
            top: '5%',
            containLabel: true
          },
          series: [
            {
              data: [],
              type: 'line',
              step: "end",
              areaStyle: {
                normal: {
                  color: new echarts.graphic.LinearGradient(//设置渐变颜色
                    0, 0, 0, 1,
                    [
                      {offset: 0, color: '#a4d7f4'},
                      {offset: 1, color: '#dcf6ff'}
                    ]
                  )
                }
              },
              smooth: false,//曲线平滑
              itemStyle: {
                normal: {
                  color: '#3598db',
                  borderColor: '#3598db',  //拐点边框颜色
                  opacity: 0,//拐点显示隐藏 0为隐藏
                }
              },
            }
          ]
        };
        if (line) {
          if (this.informationValidatorsLine.noDatayAxisDefaultMaxByValidators) {
            option.yAxis.max = "100"
          } else {
            option.yAxis.max = null
          }
          option.series[0].data = this.informationValidatorsLine.seriesData;
          line.setOption(option, true);
        }
      }
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.onWindowResize);
    },
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
    #echarts_validator_line {
      width: 100%;
      height: 85%;
    }
  }
  .echarts_component_wrap_line_mobile{
    min-width:4rem;
  }


</style>
