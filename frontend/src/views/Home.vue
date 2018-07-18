<template>
  <div class="home_wrap">
    <div :class="pageClassName">
      <div class="information_preview">
        <div class="information_preview_module">
          <span>¥ {{marketCapValue}}</span>
          <span>Market Cap :</span>
        </div>
        <div class="information_preview_module">
          <span>¥ {{priceValue}}</span>
          <span>Price :</span>
        </div>
        <div class="information_preview_module">
          <span>{{validatorsValue}}</span>
          <span>Validators :</span>
        </div>
        <div class="information_preview_module">
          <span>{{votingPowerValue}}</span>
          <span>Voting Power :</span>
        </div>
        <div class="information_preview_module">
          <span>{{transactionValue}}</span>
          <span>Transactions :</span>
        </div>
      </div>
      <div :class="module_item_wrap">
        <div class="home_module_item home_module_item_pie">
          <echarts-pie :information="information"></echarts-pie>
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
  import EchartsPie from "../components/EchartsPie";

  export default {
    name: 'app-header',
    components: {EchartsPie},
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
        information:{},//饼图的所有信息

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
      document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
      setTimeout(()=>this.information = {name:'lsc'},3000)
    },

    methods: {}
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .home_wrap {
    @include flex();
    @include pcContainer;
    //pc端和移动端公共样式
    .personal_computer_home_wrap, .mobile_home_wrap{
      margin-top:1.5rem;
      .information_preview{
        @include flex;
        margin-bottom:1.5rem;

        .information_preview_module{
          border-right:0.1rem solid #eee;
          @include flex;
          flex-direction:column;
          align-items:center;
          &:last-child{
            border-right:none;
          }
          span{
            &:first-child{
              font-size:1.8rem;
              font-weight:500;
            }
          }
        }
      }
      //饼状图
      .home_module_item_pie{
        height:25rem;
      }
    }



    .personal_computer_home_wrap {
      @include pcCenter;
      .information_preview{
        .information_preview_module{
          flex:1;
        }
      }



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
      .information_preview{
        overflow-x: scroll;
        .information_preview_module{
          min-width:16rem;
        }
      }
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
