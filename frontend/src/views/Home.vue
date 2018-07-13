<template>
  <div class="home_wrap">
    <div :class="pageClassName">
      <div class="information_preview">
        <div>
          <span>¥ {{marketCapValue}}</span>
          <span>Market Cap :</span>
        </div>
        <div>
          <span>¥ {{priceValue}}</span>
          <span>Price :</span>
        </div>
        <div>
          <span>{{validatorsValue}}</span>
          <span>Validators :</span>
        </div>
        <div>
          <span>{{votingPowerValue}}</span>
          <span>Voting Power :</span>
        </div>
        <div>
          <span>{{transactionValue}}</span>
          <span>Transactions :</span>
        </div>
      </div>
      <div :class="module_item_wrap">
        <div class="home_module_item home_module_item_status">
          validators top 10
        </div>
        <div class="home_module_item">transaction History</div>
      </div>
      <div :class="module_item_wrap">
        <div class="home_module_item">
          <div>validators bottom 6</div>
          <div>candidate top 6</div>
        </div>
        <div class="home_module_item">Blocks</div>
      </div>

    </div>
  </div>

</template>

<script>

  import Tools from '../common/Tools';

  export default {
    name: 'app-header',
    data() {
      return {
        devicesWidth: window.innerWidth,
        pageClassName: 'personal_computer_home_wrap',//1是显示pc端，0是移动端
        module_item_wrap: 'module_item_wrap_computer',//module_item_wrap_computer是pc端，module_item_wrap_mobile是移动端
        marketCapValue: '3,222,333.02',
        validatorsValue: 'hello',
        transactionValue: 'world',
        priceValue: '23,222,444.21',
        votingPowerValue: 'hello world',
        blockHeightValue: 'aaa',
        timestampValue: 'bbb',

      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.pageClassName = 'personal_computer_home_wrap';
        this.module_item_wrap = 'module_item_wrap_computer';
      } else {
        this.pageClassName = 'mobile_home_wrap';
        this.module_item_wrap = 'module_item_wrap_mobile';
      }
    },
    mounted() {
      document.getElementById('router_wrap').addEventListener('click', this.hideFeature)
    },

    methods: {}
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .home_wrap {
    @include flex();
    @include pcContainer;
    .personal_computer_home_wrap {
      @include pcCenter;
      border: 0.1rem solid #3190e8;
      margin-top: 2rem;
      .module_item_wrap_computer {
        width: 100%;
        @include flex();
        justify-content: space-between;
        .home_module_item {
          width: 45%;
          margin-bottom: 1rem;
          border: 0.1rem solid #eee;
        }
      }

    }
    .mobile_home_wrap {
      flex-direction: column;
      justify-content: space-between;
      width: 100%;
      padding: 1rem;
      .module_item_wrap_mobile {
        @include flex();
        flex-direction: column;
        align-items: center;
        .home_module_item {
          width: 98%;
          margin-bottom: 1rem;
          border: 0.1rem solid #eee;
        }
      }
    }
    .home_module_item_status {
      padding: 1rem;
      background: #3190e8;

      .current_block {
        font-weight: 500;
        color: #fff;
        display: inline-block;
        height: 2.8rem;
        line-height: 2.8rem;
        border-bottom: 0.1rem solid #ffffff;
        width: 100%;

      }
      .home_status_bottom {
        @include flex();

        .home_status_bottom_content {
          flex: 1;
          @include flex();
          flex-direction: column;
          span {
            font-size: 1.4rem;
            color: #ffffff;
            font-weight: normal;
          }
        }

      }
    }
  }
</style>
