<template>
  <div class="propsals_echart_container">
    <div class="text">
      <div class="top">
        <span class="title">ID:</span>
        <span class="value">{{data.proposal_id}}</span>
        <span class="title" style="margin-left: 40px;">Title:</span>
        <span class="value">{{data.title}}</span>
      </div>
      <div class="content">
        <div class="content_div">
          <div>
            <img src="../../assets/pass.png"/>
            <span>Participation > {{data.participation}} %</span>
          </div>
          <div style="margin-top: 16px;">
            <img src="../../assets/no_pass.png"/>
            <span>Pass Threshold > {{data.threshold}} %</span>
          </div>
        </div>
      </div>
    </div>
    <div class="propsals_echart_content">
      <div class="propsals_echart" ref="propsalsEchart">
        <div class="propsals_echart_center"></div>
      </div>
    </div>
  </div>
</template>

<script>
import echarts from 'echarts';

export default {
  name: 'MProposalsEchart',
  data() {
    return {
      chart: null,
      level: 4,
      levelName: '',
      levels: [
        {
          r0: '0%',
          r: '25%',
          itemStyle: {
            borderWidth: 0,
            color: '#51A9FF'
          }
        }, 
        {
          r0: '25%',
          r: '50%',
          itemStyle: {
            borderWidth: 0
          }
        },
        {
          r0: '50%',
          r: '75%'
        },
        {
          r0: '75%',
          r: '100%'
        }
      ],
      bundClickEventFun: function() {}
    }
  },
  props: {
    data: {
      type: Object,
      default: function() {return {}}
    }
  },
  methods: {
    configuration(data, levels) {
      let that = this;
      let option = {
        series: {
          type: 'sunburst',
          data: data,
          radius: ['0%', '100%'],
          label: {
            rotate: '0'
          },
          sort: null,
          startAngle: 145,
          nodeClick: false,
          levels: levels
        },
        color: ['rgba(254, 138, 138, 0.6)'],
        tooltip: {
          formatter: function(v) {
            let info = v.data.info;
            that.levelName = v.name;
            if (v.data.info) {
              return `${v.marker}voter: ${info.voter_moniker}<br><span style="margin-left: 15px;">voting_power: ${info.voting_power}</span>`;
            } else {
              return `${v.marker}${v.name}: ${v.value}`;
            }
          }
        }
      };
      this.chart.setOption(option);
    },
    bundClickFun(w, h) {
      return (e) => {
        let x = (e.offsetX - w);
        let y = (e.offsetY - h);
        let l = Math.sqrt(x*x + y*y);
        let per = w / this.level;
        if (this.level === 4) {
          let levels = [];
          if (l > per && l <= per * 2) {
            this.level = 3;
            let arr = this.data.data;
            let data = arr.filter(v => v.name === this.levelName);
            this.configuration(data, levels);
          }
          if (l > per * 2 && l <= per * 3) {
            if (this.levelName) {
              this.level = 2;
              let arr = this.data.data[0].children;
              let data = arr.filter(v => v.name === this.levelName);
              this.configuration(data, levels);
            }
          }
        } else if (this.level === 3){
          if (l >= 0 && l <= per) {
            this.level = 4;
            this.configuration(this.data.data, this.levels);
          }
          if (l > per && l <= per * 2) {
            if (this.levelName) {
              this.level = 2;
              let levels = [];
              let arr = this.data.data[0].children;
              let data = arr.filter(v => v.name === this.levelName);
              this.configuration(data, levels);
            }
          }
        } else if (this.level === 2) {
          if (l >= 0 && l <= per) {
            this.level = 3;
            let levels = [];
            let arr = [this.data.data[0]];
            this.configuration(arr, levels);
          }
        }
      }
    },
    bundClick() {
      let w = this.$refs.propsalsEchart.getBoundingClientRect().width / 2;
      let h = this.$refs.propsalsEchart.getBoundingClientRect().height / 2;
      this.bundClickEventFun = this.bundClickFun(w, h);
      this.$refs.propsalsEchart && this.$refs.propsalsEchart.addEventListener('click', this.bundClickEventFun);
    },
    unBundClick() {
      this.$refs.propsalsEchart && this.$refs.propsalsEchart.removeEventListener('click', this.bundClickEventFun);
    }
  },
  mounted() {
    this.chart = echarts.init(this.$refs.propsalsEchart);
    this.bundClick();
    this.configuration(this.data.data, this.levels);
  },
  destroyed() {
    this.unBundClick();
  }
}
</script>

<style lang="scss" scoped>
  .propsals_echart_container {
    width: 6.3rem;
    height: 4.6rem;
    border-radius: 1px;
    border: 1px solid #D7D9E0;
    display: flex;
    justify-content: space-between;
    .text {
      width: 200px;
      position: relative;
      .top {
        padding: 0.3rem;
        font-size: 12px;
        white-space: nowrap;
        .title {
          color: #A2A2AE;
          white-space: nowrap;
        }
        .value {
          margin-left: 6px;
          color: #3598DB;
          white-space: nowrap;
        }
      }
      .content {
        width: 100%;
        margin: auto;
        top: 0;
        display: flex;
        div.content_div {
          margin-left: 0.3rem;
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
    .propsals_echart_content {
      flex: 1;
      display: flex;
      justify-content: center;
      .propsals_echart {
        width: calc(100% - 20px);
        max-width: 3.2rem;
        height: calc(100% - 20px);
        max-height: 3.2rem;
        align-self: center;
        .propsals_echart_center {
          width: 0.8rem;
          height: 0.8rem;
        }
      }
    }
  }
</style>
