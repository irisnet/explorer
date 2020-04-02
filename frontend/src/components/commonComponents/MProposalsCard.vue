<template>
  <div :class="['propsals_card_container', $store.state.isMobile ? 'mobile_propsals_card_container' : '']">
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
    <div class="text_content">
      <div>
        <img v-if="data.level === 'Important'" src="../../assets/important.png" />
        <img v-if="data.level === 'Normal'" src="../../assets/normal.png" />
        <img v-if="data.level === 'Critical'" src="../../assets/critical.png" />
        <span>{{data.type}}</span>
      </div>
      <div>
        <img src="../../assets/deposit_period.png" />
        <span>DepositPeriod</span>
      </div>
      <div v-if="flShowTime" style="margin-left: 0.15rem;display: flex;align-items: center">
        <i style="color: rgb(90, 200, 250);padding-right: 0.05rem" class="iconfont iconHoursLeft"></i> {{depositHourLeft}} Left
      </div>
    </div>
    <div class="content" ref="content">
      <span class="start" :style="{marginLeft: `${startMarginLeft + 30}px`}">0</span>
      <div class="progress">
        <div class="current" :style="{width: `${data.total_deposit_number_per}`}">
          <div class="current_value">
            <span ref="totalDeposit">CurrentDeposit {{data.total_deposit_format}}</span>
          </div>
        </div>
        <div class="intital" :style="{left: `${data.intial_deposit_number_per}`}">
          <div class="intital_value">
            <span ref="intialDeposit">InitialDeposit {{data.intial_deposit_format}}</span>
          </div>
        </div>
      </div>
      <span class="end">
        <span>
          {{data.min_deposit_format}}
          <span class="end_value">(MinDeposit)</span>
        </span>
      </span>
    </div>
  </div>
</template>

<script>
  import Tools from "../../util/Tools"
export default {
  name: 'MProposalsCard',
  props: {
    data: {
      type: Object,
      default: function() {return {}}
    }
  },
  watch:{
    data(data){
      this.getVotingEndTime(data.deposit_end_time)
    }
  },
  data() {
    return {
      titleValueTipShow: false,
      startMarginLeft: 0,
      depositTimer: null,
      depositHourLeft:'',
      flShowTime:false
    }
  },
  methods:{
    getDepositEndTime(time){
      if(time){
        let that = this;
        let currentServerTime = new Date().getTime() + this.diffMilliseconds;
        if(new Date(time).getTime() >  currentServerTime){
          that.flShowTime = true;
          that.depositHourLeft =  Tools.formatAge(new Date(time).getTime(),currentServerTime);
        }
        clearInterval(this.depositTimer);
        this.depositTimer = setInterval(() => {
          currentServerTime = new Date().getTime() + this.diffMilliseconds;
          if(new Date(time).getTime() >  currentServerTime){
            that.flShowTime = true;
            that.depositHourLeft =  Tools.formatAge(new Date(time).getTime(),currentServerTime);
          }else {
          }
        },1000);
      }
    },
  },
  mounted() {
    this.getDepositEndTime(this.data.deposit_end_time)
    this.$nextTick(() => {
      let title_value_a_w = this.$refs.titleValue.querySelector('a').getBoundingClientRect().width;
      let w = this.$refs.titleValue.getBoundingClientRect().width;
      this.titleValueTipShow = title_value_a_w > w;
      let contentLeft = this.$refs.content.getBoundingClientRect().left;
      this.startMarginLeft = Math.ceil(Math.max(contentLeft - this.$refs.totalDeposit.getBoundingClientRect().left, contentLeft - this.$refs.intialDeposit.getBoundingClientRect().left));
      this.startMarginLeft = Math.max(0, this.startMarginLeft);
    });
  }
}
</script>

<style lang="scss" scoped>
  .propsals_card_container {
    width: 6.3rem;
    height: 100%;
    border-radius: 1px;
    border: 1px solid #D7D9E0;
    position: relative;
    display: flex;
    flex-direction: column;
    background: #fff;
    .top {
      padding: 0.3rem 0.3rem 0.22rem;
      display: flex;
      font-size: 18px;
      position: relative;
      z-index: 1;
      .title {
        color: var(--contentColor);
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
    .text_content {
      padding: 0 0.3rem;
      font-size: 12px;
      display: flex;
      align-items: center;
      div {
        img {
          width: 0.14rem;
          vertical-align: middle;
        }
        span {
          vertical-align: middle;
          margin-left: 8px;
        }
      }
      & > div:nth-child(2) {
        margin-left: 0.15rem;
      }
    }
    .content {
      width: 100%;
      display: flex;
      flex: 1;
      align-items: center;
      span.start {
        text-align: right;
        margin-left: 30px;
        margin-right: 16px;
      }
      span.end {
        margin-left: 16px;
        margin-right: 30px;
        color: var(--contentColor);
        span.end_value {
          color: var(--contentColor);
        }
      }
      .progress {
        width: 4.6rem;
        height: 0.12rem;
        background: #E5E9FB;
        border-radius: 0.17rem;
        position: relative;
        overflow: visible;
        white-space: nowrap;
        flex: 1;
        .current {
          height: 100%;
          background: var(--bgColor);
          border-radius: 0.17rem;
          .current_value {
            position: relative;
            &::before {
              display: block;
              content: '';
              width: 0;
              height: 0;
              border: 4.5px solid transparent;
              border-top-color: #FF9942;
              margin-left: 100%;
              margin-top: -2px;
              transform: translateX(-50%) translateY(-50%);
            }
            span {
              position: absolute;
              top: -100%;
              left: 100%;
              margin-top: -20px;
              font-size: 12px;
              transform: translateX(-50%);
              color: #FF9942;
            }
          }
        }
        .intital {
          position: absolute;
          width: 0px;
          height: 100%;
          overflow: visible;
          .intital_value {
            text-align: center;
            position: absolute;
            top: 100%;
            font-size: 12px;
            transform: translateX(-50%);
            color: #3698DB;
            .intital_value_span {
              display: block;
              color: var(--contentColor);
            }
            span {
              display: inline-block;
              padding-top: 2px;
            }
            &::before {
              display: block;
              content: '';
              width: 0;
              height: 0;
              border: 4.5px solid transparent;
              border-bottom-color: var(--bgColor);
              margin: auto;
              margin-top: 2px;
              transform: translateX(1px) translateY(-50%);
            }
          }
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
  .mobile_propsals_card_container {
    .top {
      padding: 0.15rem;
    }
    .text_content {
      padding: 0 0.15rem;
    }
  }
</style>
