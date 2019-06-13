<template>
  <div class="propsals_card_container">
    <div class="top">
      <span class="title">#{{data.proposal_id}}:</span>
      <div class="title_value_content">
        <span class="value title_value" ref="titleValue">
          <router-link :to="`/ProposalsDetail/${data.proposal_id}`"
                      class="link_style">{{data.title}}</router-link>
        </span>
        <div v-if="titleValueTipShow" class="tooltip_span">{{data.title}}</div>
      </div>
    </div>
    <div class="content">
      <span class="start">0</span>
      <div class="progress">
        <div class="current" :style="{width: `${data.total_deposit_number_per}`}">
          <div class="current_value">
            <span class="current_value_span">{{data.total_deposit_format}}</span>
            <span>CurrentDeposit</span>
          </div>
        </div>
        <div class="intital" :style="{left: `${data.intial_deposit_number_per}`}">
          <div class="intital_value">
            <span>IntitalDeposit</span>
            <span class="intital_value_span">{{data.intial_deposit_format}}</span>
          </div>
        </div>
      </div>
      <span class="end">
        <span>MinDeposit</span>
        <span class="end_value">{{data.min_deposit_format}}</span>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MProposalsCard',
  props: {
    data: {
      type: Object,
      default: function() {return {}}
    }
  },
  data() {
    return {
      titleValueTipShow: false
    }
  },
  mounted() {
    this.$nextTick(() => {
      let title_value_a_w = this.$refs.titleValue.querySelector('a').getBoundingClientRect().width;
      let w = this.$refs.titleValue.getBoundingClientRect().width;
      this.titleValueTipShow = title_value_a_w > w;
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
    .top {
      padding: 0.3rem;
      font-size: 12px;
      display: flex;
      position: relative;
      z-index: 1;
      .title {
        color: #A2A2AE;
      }
      .value {
        margin-left: 6px;
        color: #3598DB;
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
      height: 100%;
      position: absolute;
      display: flex;
      top: 0;
      align-items: center;
      span.start {
        text-align: right;
        margin-left: 30px;
        margin-right: 16px;
      }
      span.end {
        margin-left: 16px;
        margin-right: 30px;
        span.end_value {
          display: block;
          font-size: 12px;
          color: #A2A2AE;
        }
      }
      .progress {
        width: 4.6rem;
        height: 0.2rem;
        background: #E5E9FB;
        border-radius: 0.17rem;
        position: relative;
        overflow: visible;
        white-space: nowrap;
        flex: 1;
        .current {
          height: 0.2rem;
          background: #3598DB;
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
              transform: translateX(-50%) translateY(-50%);
            }
            span {
              position: absolute;
              top: -100%;
              left: 100%;
              margin-top: -15px;
              font-size: 12px;
              transform: translateX(-50%);
              color: #FF9942;
            }
            span.current_value_span {
              margin-top: -32px;
            }
          }
        }
        .intital {
          position: absolute;
          width: 2px;
          height: 100%;
          background: #D7D9E0;
          overflow: visible;
          .intital_value {
            text-align: center;
            position: absolute;
            top: 100%;
            font-size: 12px;
            transform: translateX(-50%);
            color: #22252A;
            .intital_value_span {
              display: block;
              color: #A2A2AE;
            }
            &::before {
              display: block;
              content: '';
              width: 0;
              height: 0;
              border: 4.5px solid transparent;
              border-bottom-color: #3598DB;
              margin: auto;
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
    padding: 5px 15px;
    color: #ffffff;
    background-color: #000000;
    line-height: 35px;
    border-radius: 0.04rem;
    word-wrap: break-word;
    white-space: normal;
    line-height: 1.7;
    &::after {
      width: 0;
      height: 0;
      border: 0.04rem solid transparent;
      content: "";
      display: block;
      position: absolute;
      border-top-color: #000000;
      left: 50%;
      margin-left: -4px;
      bottom: -8px;
    }
  }
</style>
