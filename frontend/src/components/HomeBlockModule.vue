<template>
  <div class="home_module_block">
    <div class="home_module_block_title_wrap">
      <span class="home_module_block_title">{{title}}</span>
      <span class="view_all_btn">View All</span>
    </div>
    <div class="home_module_block_content">
      <div class="home_module_block_content_item" >
        <div>
          1
        </div>
        <div>
          <span></span>
          <span></span>
          <span></span>
        </div>
      </div>
    </div>
  </div>


</template>

<script>
  import Tools from '../common/Tools';

  export default {
    name: 'home-block-module',
    watch:{
      blockInformation(blockInformation){
        console.log(this.titleBlock)

      },
      transactionInformation(transactionInformation){
        console.log(this.titleTransaction)

      },

    },
    data() {
      let isPersonalComputer = Tools.currentDeviceIsPersonComputer();
      return {
        deviceType:window.innerWidth > 500 ? 1 : 0,

      }
    },
    props:['title','blockInformation','transactionInformation'],
    beforeMount() {

    },
    mounted() {




    },

    methods: {}
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .home_module_block{
    width:100%;
    height:100%;
    @include flex;
    flex-direction:column;
    .home_module_block_title_wrap{
      height:5rem;
      @include flex;
      padding:1.2rem 1rem 0 1rem;
      justify-content: space-between;
      background: #efeff1;
      border-bottom:1px solid #e4e4e4;
      .home_module_block_title{
        font-size:1.8rem;
        text-indent:2.5rem;
        font-weight:600;
        background: url('../assets/people.svg') no-repeat 0 -0.2rem;
      }
      .view_all_btn{
        @include viewBtn;
      }
    }
    #echarts_pie{
      width:100%;
      flex:1;
    }
    .home_module_block_content{

    }
  }



</style>
