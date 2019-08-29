<template>
  <div class="home_module_block">
    <div class="home_module_block_title_wrap">
      <span class="home_module_block_title"
            :class="homeModuleBlockTitle"
      ><i :class="this.moduleName === 'Blocks' ? 'iconfont iconBlockHeight': 'iconfont iconshishijiaoyijilu\n'"></i>{{moduleName}}</span>
      <router-link class="view_all_btn" :to="moduleName === 'Blocks' ? `/blocks` : `/txs`">
        <span class="view_all_Content">View All</span>
      </router-link>
    </div>
    <div class="home_module_block_content">
      <div class="content" >
        <div class="home_module_block_content_item" :class="item.flShowTranslationalAnimation ? 'animation ' : ''"  v-for="item in information" :style="innerWidth<500?'padding:0.1rem;':''">
          <div class="blocks_module_left" :class="item.showAnimation === 'show' ? 'fadeIn_animation ' : ''" :style="`${moduleName === 'Blocks'?'flex:1;':''}`">
            <div class="key_value_wrap">
            <span class="blocks_module_value" :class="moduleTitle">
              <span class="transactions_tx" v-if="item.TxHash">TX# </span>
             <span style="cursor:pointer;">
               <router-link :to="item.Height?`/block/${item.Height}`:`/tx?txHash=${item.TxHash}`" style="color: var(--bgColor) !important;">{{item.Height?item.Height:`${item.TxHash && item.TxHash.substr(0,16)}...`}}</router-link>
               </span>
              </span>
              <span class="key_value_transfers_age" v-show="moduleName !== 'Blocks'">{{item.age}}</span>
              <span class="key_value_blocks_age" v-show="moduleName == 'Blocks'">{{item.age}}</span>
            </div>
            <div class="key_value_wrap_bottom">
              <span class="blocks_module_props">{{item.Height?'Txn:':''}}</span>
              <span class="blocks_module_Amount">{{item.Height?item.Txn:''}}</span>
              <span class="blocks_module_type" v-show="item.TxHash">{{item.Type}}</span>
              <div class="blocks_module_right" :style="`${moduleName === 'Blocks'?'flex:2;':''}`">
                <span :class="`${moduleName === 'Blocks' ? 'hide_fee' : 'show_fee'}`">Fee: {{item.Fee}}</span>
                <span v-show="moduleName === 'Blocks'">{{item.Time}}</span>
              </div>
            </div>
          </div>

        </div>
        <div class="none_data_img_container" v-if="information.length === 0">
          <div class="nodata_img_content">
            <div>
              <img src="../assets/no_data.svg">
            </div>
            <span v-show="moduleTitle !== 'blocks'">No Transaction</span>
            <span v-show="moduleTitle === 'blocks'">No Block</span>
          </div>
        </div>
      </div>
    </div>
  </div>


</template>

<script>
  export default {
    name: 'home-block-module',
    watch:{
      information(information){
        setTimeout(function () {
         information.forEach(item => {
           if(item.showAnimation){
             item.showAnimation = " "
           }
         })
        },1000)
      },
    },
    data() {
      return {
        deviceType:window.innerWidth,
        homeModuleBlockTitle:this.moduleName === 'Blocks'?'blocks_background_img':'transactions_background_img',
        moduleTitle:this.moduleName === 'Blocks'?'blocks':'transactions',
        innerWidth : window.innerWidth
      }
    },
    props:['moduleName','information'],
    beforeMount() {

    },
    mounted() {
      window.addEventListener('resize',this.onresize);
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.onresize);
    },

    methods: {
      onresize(){
        this.innerWidth = window.innerWidth;
      },
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';
  .module_item_wrap_mobile
  .home_module_item .home_module_block
  .home_module_block_content
  .home_module_block_content_item{
    height: auto!important;
  }
  .home_module_block_content{
    .content{
      .animation{
        animation:translational_animation 1.1s infinite;
      }
    }
  }
  @keyframes translational_animation{
    from{transform: translateY(-0.59rem);}
    to{transform: translateY(0)}
  }
  .none_data_img_container{
    width: 100%;
    height: 100%;
    min-height: 6.95rem;
    position: relative;
    top: -0.2rem;
    .nodata_img_content{
      @include center;
      display: flex;
      flex-direction: column;

      div{
        margin: 0 auto;
        width: 0.84rem;
        height: 0.84rem;
        img{
          width: 100%;
        }
      }
      span{
        text-align: center;
        padding-top: 0.16rem;
        @include fontSize;
        color: var(--contentColor);
      }
    }

  }
  .home_module_block{
    width:100%;
    height:100%;
    @include flex;
    flex-direction:column;
    .home_module_block_title_wrap{
      @include flex;
      padding:0.2rem;
      height:0.64rem;
      justify-content: space-between;
      background: #efeff1;
      border-bottom:1px solid #e4e4e4;
      align-items: center;
      .home_module_block_title{
        font-size:0.18rem;
        @include fontWeight;
        text-align: left;
        color:var(--moduleColor);
        i{
          padding-right: 0.1rem;
          font-size: 0.2rem;
          color: var(--contentColor);
        }
      }
      .view_all_btn{
      }
    }
    #echarts_pie{
      width:100%;
      flex:1;
    }
    .home_module_block_content{
      flex:1;
      overflow-y:hidden;
      -webkit-overflow-scrolling:touch;
      .home_module_block_content_item{
        @include flex;
        justify-content:space-between;
        border-bottom:1px solid #eee;
        padding:0 0.2rem;
        height: 0.59rem;
        &:last-child{
          border-bottom:none;
        }
        .blocks_module_left{
          @include flex;
          flex-direction:column;
          flex:2;
          justify-content:center;
          .from_to_wrap{
            @include flex;
            flex-wrap: wrap;
          }
          .key_value_wrap_bottom{
            @include flex;
            justify-content: space-between;
            @include fontSize;
            .blocks_module_Amount{
              color:var(--contentColor);
              display:inline-block;
            }
            .blocks_module_type{
              @include fontSize;
              color:var(--contentColor);
              display:inline-block;
            }
            .blocks_module_props{
              @include fontSize;
              color:var(--contentColor);
            }
          }
          .key_value_wrap{
            @include flex;
            justify-content: space-between;
            .key_value_transfers_age{
              display: inline-block;
              @include fontSize;
              color: var(--contentColor);
              text-align: right;
            }
            .key_value_blocks_age{
              display: inline-block;
              @include fontSize;
              color: var(--contentColor);
              text-align: right;
            }
            .blocks_module_value{
              color:var(--bgColor);
              @include fontSize;
              display:inline-block;
              .transactions_tx{
                color: #000;
              }
              .transactions{
                display: inline-block;
                width: 0.18rem;
                height: 0.18rem;
                img{
                  width: 100%;
                }
              }
            }
            .blocks{
              color:var(--bgColor);
            }
            .transactions{

            }
          }

        }
        .blocks_module_right{
          @include flex;
          flex-direction:column;
          align-items: flex-end;
          flex:1;
          justify-content:center;
          span{
            @include fontSize;
            color:var(--contentColor);
            text-align: right;
          }
        }
      }
    }
  }
  .hide_fee{
    display: none;
  }
  .show_fee{
    display: block;
    line-height: 1;
  }
  .fadeIn_animation{
   animation: fadeIn 1s 1 0s;
  }
  @-webkit-keyframes fadeIn {
    0% {
      opacity: 0;
    }
    50% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }
  .view_all_Content{
    color: var(--bgColor);
    font-size: 0.14rem;
    border-bottom: 0.01rem solid var(--bgColor);
    i{
      display: inline-block;
      transform: rotate(-90deg);
      font-size: 0.14rem;
      padding-top: 0.02rem;
    }
  }
</style>
