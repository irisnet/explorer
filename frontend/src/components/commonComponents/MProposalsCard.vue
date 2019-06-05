<template>
  <div class="propsals_card_container">
    <div class="top">
      <span class="title">ID:</span>
      <span class="value">10</span>
      <span class="title" style="margin-left: 40px;">Title:</span>
      <span class="value">Title:</span>
    </div>
    <div class="content">
      <span class="start">o</span>
      <div class="progress">
        <div class="current">
          <div class="current_value">
            <span class="current_value_span">1000</span>
            <span>CurrentDeposit</span>
          </div>
        </div>
        <div class="intital">
          <div class="intital_value">
            <span>IntitalDeposit</span>  
            <span class="intital_value_span">10000</span>
          </div>
        </div>
      </div>
      <span class="end">
        <span>MinDeposit</span>
        <span class="end_value">100000</span>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MProposalsCard'
}
</script>

<style lang="scss" scoped>
  .propsals_card_container {
    width: 6.3rem;
    height: 2.7rem;
    border-radius: 1px;
    border: 1px solid #D7D9E0;
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
      display: flex;
      margin-top: 0.6rem;
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
        flex: 1;
        .current {
          width: 2.64rem;
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
          left: 100px;
          overflow: visible;
          .intital_value {
            position: absolute;
            top: 100%;
            font-size: 12px;
            transform: translateX(-50%);
            color: #22252A;
            .intital_value_span {
              display: block;
              color: #A2A2AE;
            }
          }
        }
      }
    }
  }
</style>
