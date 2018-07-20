<template>
  <div class="home_module_block">
    <div class="home_module_block_title_wrap">
      <span class="home_module_block_title">{{title}}</span>
      <span class="view_all_btn" @click="viewAllClick()">View All</span>
    </div>
    <div class="home_module_block_content">
      <div class="home_module_block_content_item" v-for="item in information">
        <div class="blocks_module_left">
          <div class="key_value_wrap">
            <span class="blocks_module_value"
                  @click="skipRouter(item.Height?`/blocks_detail/${item.Height}`:`/tx?txHash=${item.TxHash}`)">
              {{item.Height?item.Height:`TX# ${item.TxHash}`}}</span>
          </div>
          <div class="key_value_wrap" v-show="item.TxHash">
            <span class="blocks_module_props">From:</span>
            <span class="blocks_module_value" @click="skipRouter(`/address/1/${item.From}`)">{{item.From}}</span>
          </div>
          <div class="key_value_wrap" v-show="item.TxHash">
            <span class="blocks_module_props">To:</span>
            <span class="blocks_module_value" @click="skipRouter(`/address/1/${item.To}`)">{{item.To}}</span>
          </div>

          <div class="key_value_wrap">
            <span class="blocks_module_props">{{item.Height?'Txn:':'Amount:'}}</span>
            <span class="blocks_module_value" style="color:#555;">{{item.Height?item.Txn:item.Amount}}</span>
          </div>
        </div>
        <div class="blocks_module_right">
          <span>{{item.Time}}</span>
          <span>Fee: {{item.Fee}}</span>
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
        console.log(information)

      },

    },
    data() {
      let isPersonalComputer = Tools.currentDeviceIsPersonComputer();
      return {
        deviceType:window.innerWidth > 500 ? 1 : 0,

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
      flex:1;
      overflow-y:scroll;
      .home_module_block_content_item{
        @include flex;
        justify-content:space-between;
        border-bottom:0.1rem solid #eee;
        padding:1.2rem;
        .blocks_module_left{
          @include flex;
          flex-direction:column;

          .key_value_wrap{
            width:15rem;
            @include flex;
            flex-direction:row;
            .blocks_module_value{
              color:#3598db;
              cursor:pointer;
              font-size:1.4rem;
              display:inline-block;
              @include overflowEllipsis;
            }
            .blocks_module_props{
              font-size:1.4rem;
              color:#555;
            }
          }

        }
        .blocks_module_right{
          @include flex;
          flex-direction:column;
          align-items: flex-end;
          span{
            font-size:1.4rem;
            color:#a2a2ae;
          }
        }
      }
    }
  }



</style>
