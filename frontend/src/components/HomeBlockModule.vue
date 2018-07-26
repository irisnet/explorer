<template>
  <div class="home_module_block">
    <div class="home_module_block_title_wrap">
      <span class="home_module_block_title"
            :class="homeModuleBlockTitleVar"
      >{{title}}</span>
      <span class="view_all_btn" @click="viewAllClick()">View All</span>
    </div>
    <div class="home_module_block_content">
      <div class="home_module_block_content_item" v-for="item in information">
        <div class="blocks_module_left">
          <div class="key_value_wrap">
            <span class="blocks_module_value"
                  :class="blockModuleTypeVar"
                  @click="skipRouter(item.Height?`/blocks_detail/${item.Height}`:`/tx?txHash=${item.TxHash}`)">
              {{item.Height?item.Height:`TX# ${item.TxHash.substr(0,16)}...`}}</span>
          </div>
          <div class="key_value_wrap" v-show="item.TxHash">
            <span class="blocks_module_props">From:</span>
            <span class="blocks_module_value" @click="skipRouter(`/address/1/${item.From}`)">
              {{item.From?`${String(item.From).substr(0,16)}...`:''}}
            </span>
          </div>
          <div class="key_value_wrap" v-show="item.TxHash">
            <span class="blocks_module_props">To:</span>
            <span class="blocks_module_value" @click="skipRouter(`/address/1/${item.To}`)">
              {{item.To?`${String(item.To).substr(0,16)}...`:''}}
            </span>
          </div>

          <div class="key_value_wrap">
            <span class="blocks_module_props">{{item.Height?'Txn:':'Amount:'}}</span>
            <span class="blocks_module_value" style="color:#555;">{{item.Height?item.Txn:item.Amount}}</span>
          </div>
        </div>
        <div class="blocks_module_right">
          <span>{{item.Time}}</span>
          <span>Fees: {{item.Fee}}</span>
          <span v-show="item.TxHash">{{item.Type}}</span>
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
      information(information){

      },

    },
    data() {
      let isPersonalComputer = Tools.currentDeviceIsPersonComputer();
      return {
        deviceType:window.innerWidth > 500 ? 1 : 0,
        homeModuleBlockTitleVar:this.title === 'Blocks'?'blocks_background':'transactions_background',
        blockModuleTypeVar:this.title === 'Blocks'?'blocks_background_type':'transactions_background_type',
      }
    },
    props:['title','information',],
    beforeMount() {

    },
    mounted() {

    },

    methods: {
      skipRouter(path){
        this.$router.push(path);
      },
      viewAllClick(){
        if(this.title === 'Blocks'){
          this.$router.push('/block/1/0')
        }else if(this.title === 'Transaction'){
          this.$router.push('/recent_transactions/2/0')
        }
      }
    }
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
      height:0.5rem;
      @include flex;
      padding:0.12rem 0.1rem 0 0.1rem;
      justify-content: space-between;
      background: #efeff1;
      border-bottom:1px solid #e4e4e4;
      .home_module_block_title{
        font-size:0.18rem;
        text-indent:0.3rem;
        font-weight:600;
      }
      .blocks_background{
        background: url('../assets/blocks.png') no-repeat 0 0.02rem;
      }
      .transactions_background{
        background: url('../assets/transactions.png') no-repeat 0 0.02rem;
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
      flex:1;
      overflow-y:auto;
      .home_module_block_content_item{
        @include flex;
        justify-content:space-between;
        border-bottom:1px solid #eee;
        padding:0.12rem;
        .blocks_module_left{
          @include flex;
          flex-direction:column;

          .key_value_wrap{
            @include flex;
            flex-direction:row;
            .blocks_module_value{
              color:#3598db;
              cursor:pointer;
              font-size:0.14rem;
              display:inline-block;
            }
            .blocks_module_props{
              font-size:0.14rem;
              color:#555;
              font-weight:600;
            }
            .blocks_background_type{
              background: url('../assets/block.png') no-repeat 0 0.02rem;
              text-indent:0.2rem;
              color:#3598db;
            }
            .transactions_background_type{
              background: url('../assets/transaction.png') no-repeat 0 0.02rem;
              text-indent:0.2rem;
            }
          }

        }
        .blocks_module_right{
          @include flex;
          flex-direction:column;
          align-items: flex-end;
          span{
            font-size:0.14rem;
            color:#a2a2ae;
            text-align: right;
          }
        }
      }
    }
  }



</style>
