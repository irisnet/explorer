<template>
  <div class="propsals_echart_container">
    <div class="text">
      <div class="top">
        <span class="title">ID:</span>
        <span class="value">10</span>
        <span class="title" style="margin-left: 40px;">Title:</span>
        <span class="value">Title:</span>
      </div>
      <div class="content">
        <div class="content_div">
          <div>
            <img src="../../assets/no_pass.png"/>
            <span>Threshold > 50%</span>
          </div>
          <div style="margin-top: 16px;">
            <img src="../../assets/pass.png"/>
            <span>Threshold > 50%</span>
          </div>
        </div>
      </div>
    </div>
    <div class="propsals_echart" ref="propsalsEchart"></div>
  </div>
</template>

<script>
import echarts from 'echarts';

export default {
  name: 'MProposalsEchart',
  data() {
    return {
      chart: null
    }
  },
  methods: {
    configuration() {
      var data = [
        {
          name: 'Vote',
          value: 1,
          itemStyle: {
            color: '#3598DB',
            borderColor: '#CCDCFF',
            borderWidth: 1
          },
          children: [
            {
              name: '4',
              value: 0.3,
              zlevel: 4,
              nodeClick: false,
              itemStyle: {
                color: '#FE8A8A',
                borderColor: '#CCDCFF',
                borderWidth: 0
              }
            },
            {
              name: '3',
              value: 0.3,
              zlevel: 3,
              nodeClick: false,
              itemStyle: {
                color: '#FFCF65',
                borderColor: '#CCDCFF',
                borderWidth: 0
              }
            },
            {
              name: '2',
              value: 0.3,
              zlevel: 2,
              nodeClick: false,
              itemStyle: {
                color: '#CCDCFF',
                borderColor: '#CCDCFF',
                borderWidth: 0
              }
            },
            {
              name: '1',
              value: 0.1,
              nodeClick: false,
              zlevel: 1,
              itemStyle: {
                color: '#45B4FF',
                borderColor: '#CCDCFF',
                borderWidth: 0
              }
            }
          ]
        },
        {
          name: '',
          value: 2,
          nodeClick: false,
          itemStyle: {
            color: '#E5E9FB',
            borderColor: '#CCDCFF',
            borderWidth: 1
          },
          children: [
            {
              name: '',
              value: 2,
              nodeClick: false,
              itemStyle: {
                color: '#E5E9FB',
                borderColor: '#CCDCFF',
                borderWidth: 0
              }
            }
          ]
        }
      ];
      let option = {
        series: {
          type: 'sunburst',
          data: data,
          radius: ['0%', '100%'],
          label: {
            rotate: '0'
          },
          sort: function(a, b) {
            if (!a.name && a.depth === 1) {
              return -1;
            }
          },
          startAngle: 145
        },
        color: ['#51A9FF']
      };
      this.chart.setOption(option);
    }
  },
  mounted() {
    this.chart = echarts.init(this.$refs.propsalsEchart);
    this.configuration();
  }
}
</script>

<style lang="scss" scoped>
  .propsals_echart_container {
    width: 6.3rem;
    height: 2.7rem;
    border-radius: 1px;
    border: 1px solid #D7D9E0;
    display: flex;
    justify-content: space-between;
    .text {
      flex: 1;
      position: relative;
      .top {
        padding: 0.3rem;
        font-size: 12px;
        .title {
          color: #A2A2AE;
        }
        .value {
          margin-left: 6px;
          color: #3598DB;
        }
      }
      .content {
        width: 100%;
        height: 100%;
        margin: auto;
        position: absolute;
        top: 0;
        display: flex;
        justify-content: center;
        div.content_div {
          margin: auto;
          display: inline-block;
          & > div {
            font-size: 12px;
            color: #22252A;
            vertical-align: middle;
            img {
              width: 0.14rem;
              height: 0.14rem;
            }
            span {
              margin-left: 8px;
              vertical-align: middle;
            }
          }
        }
        
      }
    }
    .propsals_echart {
      width: 2rem;
      height: 2rem;
      margin-right: 0.9rem;
      align-self: center;
    }
  }
</style>
