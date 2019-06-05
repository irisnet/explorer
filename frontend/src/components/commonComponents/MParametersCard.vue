<template>
  <div class="m_parameters_container">
    <div class="top">
      <span class="title">
        {{data.key}}
      </span>
      <span :class="['desc', descTipShow ?'desc_hover' : '']" ref="desc" v-show="data.description">
        <div class="desc_content" style="width: 100%;">
          <span class="desc_content_span">（{{data.description}}）</span>
        </div>
        <div v-if="descTipShow" class="tooltip_span">（{{data.description}}）</div>
      </span>
    </div>
    <div class="content">
      <span class="min">{{data.min}}</span>
      <div class="progress" ref="progress">
        <div class="genesis" v-show="data.genesis !== ''" :style="{left: `${data.genesis_per}%`}">
          <div class="arrow"></div>
          <span class="genesis_value">{{data.genesis}}</span>
        </div>
        <div class="current" v-show="data.current !== ''" :style="{left: `${data.current_per}%`, height: equal ? '50%' : '100%'}">
          <div class="arrow"></div>
          <span class="current_value">{{data.current}}</span>
        </div>
      </div>
      <span class="max">{{data.max}}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MParametersCard',
  props: {
    data: {
      type: Object,
      default: function(){return{}}
    }
  },
  data() {
    return {
      equal: false,
      descTipShow: false
    }
  },
  mounted() {
    if ((typeof this.data.current_per === 'number') && (typeof this.data.genesis_per === 'number')) {
      let w = this.$refs.progress.offsetWidth;
      this.equal = Math.abs((this.data.current_per - this.data.genesis_per) * w) < 1;
      this.$nextTick(() => {
        let desc_content_w = this.$refs.desc.querySelector('.desc_content_span').getBoundingClientRect().width;
        let w = this.$refs.desc.getBoundingClientRect().width;
        this.descTipShow = desc_content_w > w;
      });
    }
  }
}
</script>

<style lang="scss" scoped>
.General {
  .content {
    .progress {
      background-color: #E7F5FF!important;
    }
  }
}
.Staking {
  .content {
    .progress {
      background-color: #E8EFFF!important;
    }
  }
}
.Slashing {
  .content {
    .progress {
      background-color: #FFF1DB!important;
    }
  }
}
.m_parameters_container {
  width: 413px;
  height: 164px;
  border-radius: 1px;
  border: 1px solid #D7D9E0!important;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding-bottom: 0.2rem;
  .top {
    font-size: 0.14rem;
    color: #22252A;
    padding: 0.2rem;
    white-space: nowrap;
    display: flex;
    position: relative;
    .desc {
      font-size: 0.12rem;
      color: #A2A2AE;
      display: block;
      width: 1px;
      flex: 1;
      position: relative;
      &:hover .tooltip_span{
        display: block;
      }
      .desc_content {
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
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
        margin-top: -10px auto 0;
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
    }
    .desc_hover {
      cursor: pointer;
    }
  }
  .content {
    display: flex;
    justify-content: center;
    font-size: 0.12rem;
    position: relative;
    padding-bottom: 0.2rem;
    @mixin min_max {
      width: 80px;
      color: #A2A2AE;
      word-break: break-all;
      flex: 1 0 0.8rem;
    }
    .min {
      padding-right: 0.1rem;
      @include min_max;
      text-align: right;
    }
    .max {
      padding-left: 0.1rem;
      @include min_max;
    }
    .progress {
      width: 2.4rem;
      height: 0.2rem;
      background-color: #F0F7FD;
      border-radius: 0.16rem;
      position: relative;
      overflow: visible;
      font-size: 0.12rem;
      & > div {
        height: 100%;
        width: 0.04rem;
        position: relative;
        font-weight: 400;
        overflow: visible;
      }
      @mixin spanStyle{
        display: block;
        height: 100%;
        position: absolute;
        line-height: 0.2rem;
        white-space: nowrap;
        text-align: center;
        transform: translateX(-50%);
        left: 50%;
      }
      .current {
        color: #FFB779;
        background-color: #FFB779;
        position: absolute;
        margin-left: -2px;
        div.arrow {
          width: 8px;
          height: 8px;
          display: flex;
          justify-content: center;
          align-items: center;
          margin-top: -4px;
          margin-left: -2px;
          &::before {
            display: block;
            content: '';
            width: 0;
            height: 0;
            border: 3px solid transparent;
            border-top-color: #FFB779;
          }
        }
        span {
          @include spanStyle;
          margin-top: -0.06rem;
          top: -0.2rem;
        }
      }
      .genesis {
        color: #3698DB;
        background-color: #3698DB;
        position: absolute;
        margin-left: -2px;
        div.arrow {
          width: 8px;
          height: 8px;
          display: flex;
          justify-content: center;
          align-items: center;
          margin-top: 0.16rem;
          margin-left: -2px;
          &::after {
            display: block;
            content: '';
            width: 0;
            height: 0;
            border: 3px solid transparent;
            border-bottom-color: #3698DB;
          }
        }
        span {
          @include spanStyle;
          margin-top: 0.06rem;
          top: 0.2rem;
        }
      }
    }
  }
}
</style>
