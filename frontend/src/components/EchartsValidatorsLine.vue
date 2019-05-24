<template>
  <div :class="echartsComponentWrapLine">
    <div class="echarts_title_wrap_line">
      <div>
        <strong>History</strong>
        <span style="color: #a2a2ae;"> - Voting Power</span>
      </div>
      <div>
        <slot name="tabs"></slot>
      </div>
    </div>
    <div id="echarts_validator_line">
    </div>
  </div>
</template>

<script>
  import echarts from 'echarts';
  import Tools from "../util/Tools";

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
        monthNum:['01','02','03','04','05','06','07','08','09','10','11','12']
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
              for (let it of params) {
                res += `<span style="display:block;">${it.seriesName}:${it.value[1]}</span>`;
              }
              return res;
            },
            axisPointer: {
              show: true,
              axis: "x",
              type: "line",
              lineStyle: {
                color: "#a2a2ae",
              }
            }
          },
          legend: {
            data: ['Voting Power', 'Delegators'],
            icon: 'circle',
            itemWidth: 6,
            left: '-4',
            textStyle: {
              color: '#839096'
            }
          },
          xAxis: {
            axisLabel: {
              color: "#839096",
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
            }
            // splitNumber:0,
          },
          yAxis: [
            {
              type: 'value',
              axisLine: {
                show: false,
                lineStyle: {
                  color: '#a2a2ae'
                }
              },
              axisTick: {
                show: false
              },
              axisLabel: {
                color: "#839096"
              },
              splitLine: {
                lineStyle: {
                  // 使用深浅的间隔色
                  color: ['#D7DCE0'],
                  type: 'dotted'
                }
              }
            },
            {
              type: 'value',
              axisLine: {
                show: false,
                lineStyle: {
                  color: '#a2a2ae'
                }
              },
              axisTick: {
                show: false
              },
              axisLabel: {
                color: "#839096",
              },
              splitLine: {
                lineStyle: {
                  // 使用深浅的间隔色
                  color: ['#D7DCE0'],
                  type: 'dotted'
                }
              }
            }
          ],
          grid: {
            left: '0',   //图表距边框的距离
            right: '0',
            bottom: '20',
            top: 40,
            containLabel: true
          },
          series: [
            {
              name:'Voting Power',
              data: [],
              type: 'line',
              smooth: true,//曲线平滑
              lineStyle: {
                color: '#6CC7FF',
                width: 2
              },
              symbol: 'circle',
              symbolSize: 8,
              itemStyle: {
                color: '#6CC7FF',
                borderColor: '#fff',
                borderWidth: 2
              }
            },
            {
              name:'Delegators',
              data: [],
              type: 'line',
              yAxisIndex: 1,
              smooth: true,//曲线平滑
              lineStyle: {
                color: '#F5A622',
                width: 2
              },
              symbol: 'circle',
              symbolSize: 8,
              itemStyle: {
                color: '#F5A622',
                borderColor: '#fff',
                borderWidth: 2
              }
            }
          ]
        };
        if (line) {
          if (this.informationValidatorsLine.noDatayAxisDefaultMaxByValidators === "100") {
            option.yAxis.max = "100"
          } else {
            option.yAxis.max = null
          }
          // option.series[0].data = this.informationValidatorsLine.seriesData;
          this.formatData(this.informationValidatorsLine.seriesData, 0, option);
          this.formatData(this.forDatas(), 1, option);
          line.setOption(option, true);
        }
      },
      forMaxMin(data, key = 'max', unit) {
        let method = key === 'max' ? 'ceil' : 'floor';
        if (unit) {
          return Math[method](data / unit ) * unit;
        }
        let n = 10;
        if (data < n) {
          return 0;
        }
        while(data > n) {
          n = n * 10
        }
        return Math[method](data * 10 / n ) * (n / 10);
      },
      forDatas() {
        if (!Array.isArray(this.informationValidatorsLine.seriesData)) {
          return;
        }
        let l = this.informationValidatorsLine.seriesData.length;
        let arr = new Array(l);
        for (let key of arr.keys()) {
          arr[key] = [this.informationValidatorsLine.seriesData[key][0], Math.ceil(Math.random() * 530)];
        }
        return arr;
      },
      formatData(data, k, option) {
        if (!Array.isArray(data)) {
          return;
        }
        let l = data.length;
        let powerArr = data.map(v => v[1]);
        let max = Math.max.apply(null, powerArr);
        let min = Math.min.apply(null, powerArr);
        let arr = [];
        for (let [index, it] of data.entries()) {
          let o = {
            value: it,
            symbol: 'none'
          }
          if (o.value[1] >= max || o.value[1] <= min) {
            o.symbol = "circle";
          }
          if (index === 0 || index === (l -1)) {
            o.symbol = "circle";
          }
          arr.push(o);
        }
        option.series[k].data = arr;
        let splitNumber = 5;
        let minN = k === 0 ? this.forMaxMin(min, 'min', 1000000) : this.forMaxMin(min, 'min');
        let maxN = k === 0 ? this.forMaxMin(max, 'max', 1000000) : this.forMaxMin(max, 'max');
        option.yAxis[k].min = minN;
        option.yAxis[k].max = maxN;
        option.yAxis[k].interval = (maxN - minN) / splitNumber;
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
      height: 0.3rem;
      font-size:0.16rem;
      padding: 0.25rem;
      padding-bottom: 0;
      @include fontWeight;
      @include flex;
      justify-content: space-between;
      line-height: 30px;
      box-sizing: content-box;
    }
    #echarts_validator_line {
      height: 3.85rem;
      margin: 0 0.25rem;
      padding-top: 0;
    }
  }
  .echarts_component_wrap_line_mobile{
    min-width:4rem;
  }
</style>
