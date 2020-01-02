<template>
  <div :class="['propsals_echart_container', $store.state.isMobile ? 'mobile_propsals_echart_container' : '']">
    <div class="text">
      <div class="top">
        <span class="title">#{{data.proposal_id}}</span>
        <div class="title_value_content">
          <span class="value title_value" ref="titleValue">
            <router-link :to="`/ProposalsDetail/${data.proposal_id}`"
                        class="link_style">{{data.title}}</router-link>
          </span>
          <div v-if="titleValueTipShow" class="tooltip_span"><div>{{data.title}}</div></div>
        </div>
      </div>
      <div class="content">
        <div class="content_div">
          <div class="content_header_content">
            <div>
              <img v-if="data.level === 'Important'" src="../../assets/important.png" />
              <img v-if="data.level === 'Normal'" src="../../assets/normal.png" />
              <img v-if="data.level === 'Critical'" src="../../assets/critical.png" />
              <span>{{data.type}}</span>
            </div>
            <div class="deposit_period_content" style="margin-top: 12px;display: flex;align-items: center;" v-show="data.status === 'DepositPeriod'">
              <i class="iconfont iconDepositPeriod-liebiao" style="color: var(--bgColor)"></i>
              <span>DepositPeriod</span>
            </div>
            <div class="voting_period_content" style="margin-top: 12px;display: flex;align-items: center;" v-show="data.status === 'VotingPeriod'">
              <i class="iconfont iconDepositPeriod" style="color: var(--bgColor)"></i>
              <span>VotingPeriod</span>
            </div>
            <div class="voting_period_content" style="margin-top: 12px;display: flex;align-items: center;" v-show="flShowTime">
              <i class="iconfont iconHoursLeft" style="color: rgb(90, 200, 250)"></i>
              <span>{{votingHourLeft}} Left</span>
            </div>
          </div>
          <div class="per_div">
            <div class="per_title">Gov Tallying</div>
            <div style="margin-top: 16px;">
              <p>
                <img v-if="data.participation > data.participationNum" src="../../assets/participant.png"/>
                <img v-if="data.participation <= data.participationNum" src="../../assets/no_threshold.png"/>
                <span>Participation</span>
              </p>
              <span style="margin-left: 20px;">{{data.participationNum}} %</span>
            </div>
            <div style="margin-top: 16px;">
              <p>
                <img v-if="data.passThreshold > data.passThresholdNum" src="../../assets/pass_threshold.png"/>
                <img v-if="data.passThreshold <= data.passThresholdNum" src="../../assets/no_threshold.png"/>
                <span>Pass Threshold</span>
              </p>
              <span style="margin-left: 20px;">{{data.passThresholdNum}} %</span>
            </div>
            <div style="margin-top: 16px;">
              <p>
                <img v-if="data.vetoThreshold > data.vetoThresholdNum" src="../../assets/veto_threshold.png"/>
                <img v-if="data.vetoThreshold <= data.vetoThresholdNum" src="../../assets/no_threshold.png"/>
                <span>Veto Threshold</span>
              </p>
              <span style="margin-left: 20px;">{{data.vetoThresholdNum}} %</span>
            </div>
          </div>
        </div>
        <div class="propsals_echart_content">
          <div class="propsals_echart" ref="propsalsEchart">
            <div class="propsals_echart_center"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
var echarts = require('echarts')
import Tools from '../../util/Tools';

export default {
  name: 'MProposalsEchart',
  data() {
    return {
      titleValueTipShow: false,
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
      bundClickEventFun: function() {},
      votingTimer:null,
      votingHourLeft:'',
      flShowTime:false
    }
  },
  props: {
    data: {
      type: Object,
      default: function() {return {}}
    }
  },
  methods: {
    getVotingEndTime(time){
      if(time){
        let that = this;
        let currentServerTime = new Date().getTime() + this.diffMilliseconds;
        if(new Date(time).getTime() >  currentServerTime){
          that.flShowTime = true;
          that.votingHourLeft =  Tools.formatAge(new Date(time).getTime(),currentServerTime);
        }else {
        }
        clearInterval(this.votingTimer);
        this.votingTimer = setInterval(() => {
          currentServerTime = new Date().getTime() + this.diffMilliseconds;
          if(new Date(time).getTime() >  currentServerTime){
            that.flShowTime = true;
            that.votingHourLeft =  Tools.formatAge(new Date(time).getTime(),currentServerTime);
          }else {
          }
        },1000);
      }
    },
    configuration(data, levels) {
      this.forChildrenBorderWidth(data);
      let that = this;
      let option = {
        series: {
          type: 'sunburst',
          data: data,
          radius: ['0%', '100%'],
          label: {
            rotate: '0',
            textBorderColor: 'var(--contentColor)',
            textBorderWidth: 1,
            fontWeight: 600,
            fontFamily: 'Arial'
          },
          sort: null,
          startAngle: 145,
          nodeClick: false,
          levels: levels
        },
        color: ['rgba(254, 138, 138, 0.6)'],
        tooltip: {
          confine: true,
          formatter: function(v) {
            let info = v.data.info;
            that.levelName = v.name;
            if (v.data.info) {
              return `${v.marker}${info.voter_moniker || info.voter}: ${info.voting_power} (${v.data.per} %)`;
            } else {
              if (v.dataIndex === 1 || !v.name) {
                return `${v.marker}${v.data.tipName || v.name}: ${v.value} (${v.data.perData} %)`;
              }
              return `${v.marker}${v.data.tipName || v.name}: ${v.value} (${v.data.per} %)`;
            }
          },
          textStyle: {
            fontFamily: 'Arial'
          }
        }
      };
      this.chart.setOption(option);
    },
    forChildrenBorderWidth(data, all) {
      all = all || data.reduce((init, v) => v.value + init, 0);
      if (all > 0) {
        data.forEach(v => {
          let per = v.value / all;
          v.itemStyle.borderWidth = per < 1 ? 0.5 : 0;
          v.per = Tools.formatDecimalNumberToFixedNumber(per * 100);
          v.children && this.forChildrenBorderWidth(v.children, all);
        })
      }
    },
    bundClickFun(w, h) {
      return (e) => {
        let x = (e.offsetX - w);
        let y = (e.offsetY - h);
        let l = Math.sqrt(x*x + y*y);
        let per = Math.min(w, h) / this.level;
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
    this.getVotingEndTime(this.data.votingEndTime);
    this.chart = echarts.init(this.$refs.propsalsEchart);
    this.bundClick();
    this.configuration(this.data.data, this.levels);
    this.$nextTick(() => {
      let title_value_a_w = this.$refs.titleValue.querySelector('a').getBoundingClientRect().width;
      let w = this.$refs.titleValue.getBoundingClientRect().width;
      this.titleValueTipShow = title_value_a_w > w;
    });
  },
  destroyed() {
    this.unBundClick();
  }
}
</script>

<style lang="scss" scoped>
  .propsals_echart_container {
    width: 6.3rem;
    height: 4.2rem;
    border-radius: 1px;
    border: 1px solid #D7D9E0;
    display: flex;
    flex-direction: column;
      background: #fff;
    .text {
      width: 100%;
      height: 100%;
      position: relative;
      display: flex;
      flex-direction: column;
      .top {
        padding: 0.3rem 0.3rem 0.22rem;
        font-size: 18px;
        white-space: nowrap;
        display: flex;
        .title {
          color: var(--contentColor);
          white-space: nowrap;
        }
        .value {
          margin-left: 6px;
          color: var(--bgColor);
          white-space: nowrap;
        }
        .title_value_content {
          width: 1px;
          flex: 1;
          display: flex;
          position: relative;
          cursor: pointer;
          .title_value {
            width: 1px;
            flex: 1;
            display: block;
            overflow: hidden;
            text-overflow: ellipsis;
          }
          &:hover .tooltip_span{
            display: block;
          }
        }
      }
      .content {
        width: 100%;
        margin: auto;
        flex: 1;
        top: 0;
        display: flex;
        div.content_div {
          margin-left: 0.3rem;
          display: inline-block;

          div.per_title {
            margin-top: 30px;
            font-size: 14px;
            color: var(--contentColor);
          }
          & > div {
            font-size: 12px;
            color: var(--contentColor);
            vertical-align: middle;
            img {
              width: 0.14rem;
              vertical-align: middle;
            }
            span {
              vertical-align: middle;
              margin-left: 8px;
            }
            p {
              margin-bottom: 6px !important;
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
        .propsals_echart_center {
          width: 0.8rem;
          height: 0.8rem;
        }
      }
    }
  }
  .tooltip_span {
    width: 100%;
    max-width: calc(100% + 40px);
    display: none;
    position: absolute;
    z-index: 1000;
    bottom: calc(100% + 4px);
    left: 50%;
    transform: translateX(-50%);
    color: #ffffff;
    background-color: #000000;
    border-radius: 0.04rem;
    word-wrap: break-word;
    white-space: normal;
    font-size: 12px;
    line-height: 16px;
    div {
      padding: 8px 15px;
    }
    &::after {
      width: 0;
      height: 0;
      border: 0.06rem solid transparent;
      content: "";
      display: block;
      position: absolute;
      border-top-color: #000000;
      left: 50%;
      margin-left: -6px;
    }
  }
  .mobile_propsals_echart_container {
    .text {
      .top {
        padding: 0.15rem;
      }
    }
    .content {
      flex-direction: column;
      div.content_div {
        margin-left: 0.15rem!important;
        margin-right: 0.15rem!important;
        & > div:nth-child(1) {
          display: inline;
        }
        & > div:nth-child(2) {
          display: inline;
          margin-top: 0!important;
          margin-left: 0.38rem;
        }
        .per_div {
          display: flex;
          flex-wrap: wrap;
          justify-content: space-between;
          .per_title {
            width: 100%;
            flex: 1 0 100%;
            margin-top: 0.15rem!important;
          }
        }
      }
      .propsals_echart_content {
        align-items: center;
        .propsals_echart {
          height: 100%;
          width: 100%;
        }
      }
    }
  }
  @media screen and (max-width: 910px){
    .mobile_propsals_echart_container{
        background: #fff;
      .text{
        .content{
          .content_div{
            .content_header_content{
              display: flex;
                flex-direction: column;

              align-items: flex-start;
              .deposit_period_content{
                margin-top: 0!important;
              }
              .voting_period_content{
                  padding: 0.04rem 0;
                margin-top: 0!important;
              }
              .voting_period_content:last-child{
                  padding: 0;
              }
            }
            .per_div{
              margin-left: 0!important;
            }
          }
        }
      }
    }
  }
</style>
