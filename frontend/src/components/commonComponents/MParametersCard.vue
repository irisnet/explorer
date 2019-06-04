<template>
  <div class="m_parameters_container">
    <div class="top">
      <span class="title">
        {{data.key}}
      </span>
      <span class="desc" v-show="data.description">（{{data.description}}）</span>
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
      equal: false
    }
  },
  mounted() {
    if ((typeof this.data.current_per === 'number') && (typeof this.data.genesis_per === 'number')) {
      let w = this.$refs.progress.offsetWidth;
      this.equal = Math.abs((this.data.current_per - this.data.genesis_per) * w) < 1;
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
.Stake {
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
    .desc {
      font-size: 0.12rem;
      color: #A2A2AE;
    }
  }
  .content {
    display: flex;
    justify-content: center;
    font-size: 0.12rem;
    position: relative;
    padding-bottom: 0.2rem;
    @mixin min_max {
      width: 100px;
      color: #A2A2AE;
      word-break: break-all;
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
      flex: 1 0 2.4rem;
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
