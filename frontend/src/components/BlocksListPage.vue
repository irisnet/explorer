<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">

      <p :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">Blocks</span>
        <span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>
      </p>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize"
          @click="(data)=>console.log(data)"
        >
        </b-pagination>
      </div>
      <template>
        <b-table :fields='fields' :items='items' striped>
          <template slot='height' slot-scope='data'>
            <a href="http://www.baidu.com">
              {{data.item.height}}
            </a>
          </template>
          <template slot='txn' slot-scope='data'>
            <a href="http://www.baidu.com">
              {{data.item.txn}}
            </a>
          </template>
        </b-table>
      </template>
      <div class="pagination">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';

  export default {
    watch:{
      currentPage(currentPage){
        this.currentPage = currentPage;
        this.getDataList(currentPage,10,this.$route.params.type);
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',//1是显示pc端，0是移动端
        blocksValue: '',
        currentPage:1,
        pageSize:10,
        count:0,
        fields:[

        ],
        items:[]


      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
      } else {
        this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
      }
    },
    mounted() {
      //组件页面根据路由类型判断是哪个页面
      console.log(this.$route.params)



      this.getDataList(1,10,this.$route.params.type);
    },
    methods: {
      getDataList(currentPage,pageSize,type){
        if(type === '1'){
          let url = `/api/blocks/${currentPage}/${pageSize}`;
          console.log(url)
          axios.get(url).then((data)=>{
            if(data.status === 200){
              return data.data;
            }
          }).then((data)=>{
            this.count = data.Count;
            this.items = data.Data.map(item=>{
              let txn = item.NumTxs;
              let precommit = item.Block.LastCommit.Precommits.length;
              return {
                height:item.Height,
                txn,
                fee:'',
                timestamp:item.Time,
                precommit,
                voting:'',
              };
            })
          })
        }else if(type === '2'){
          let url = `/api/txs/coin/${currentPage}/${pageSize}`;
          console.log(url)
          axios.get(url).then((data)=>{
            if(data.status === 200){
              return data.data;
            }
          }).then((data)=>{
            console.log(data)
            this.count = data.Count;
            /*this.items = data.Data.map(item=>{
              let txn = item.NumTxs;
              let precommit = item.Block.LastCommit.Precommits.length;
              return {
                height:item.Height,
                txn,
                fee:'',
                timestamp:item.Time,
                precommit,
                voting:'',
              };
            })*/
          })
        }

      }
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .blocks_list_page_wrap {
    @include flex;
    @include pcContainer;
    font-size: 1.4rem;
    .pagination{
      margin-top:0.5rem;
      @include flex;
      justify-content: flex-end;
    }
    .b-table{
      border-bottom:1px solid #eeeeee;
      min-width:50rem;


      a{
        text-decoration: none;
      }
    }
    .blocks_list_title_wrap {
      width: 100%;
      border-bottom: 0.1rem solid #eee;
      @include flex;
      @include pcContainer;
      .personal_computer_blocks_list_page_wrap {
        @include flex;
      }
      .mobile_blocks_list_page_wrap {
        @include flex;
        flex-direction: column;
      }

    }
    .personal_computer_blocks_list_page_wrap {
      .transaction_information_content_title {
        height: 4rem;
        line-height: 4rem;
        font-size: 1.8rem;
        color: #555;
        margin-bottom: 0;
      }
      @include pcCenter;
      .transactions_detail_information_wrap {
        .information_props_wrap {
          @include flex;
          .information_props {
            width: 15rem;
          }
        }
      }

      .blocks_list_title {
        height: 4rem;
        line-height: 4rem;
        font-size: 1.8rem;
        color: #555;
        margin-right: 2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        height: 4rem;
        line-height: 4rem;
        font-size: 1.4rem;
        color: #ccc;
      }
    }

    .mobile_blocks_list_page_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.5rem;
      overflow-x: scroll;
      .transaction_information_content_title {
        height: 4rem;
        line-height: 4rem;
        font-size: 1.8rem;
        color: #555;
        margin-bottom: 0;
      }
      .transactions_detail_information_wrap {

        .information_props_wrap {
          @include flex;
          flex-direction: column;
          border-bottom: 0.1rem solid #eee;
          margin-bottom: 0.5rem;
          .information_value {
            overflow-x: scroll;
          }

        }
      }

      .blocks_list_title {
        height: 3rem;
        line-height: 3rem;
        font-size: 1.8rem;
        color: #555;
        margin-right: 2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        overflow-x: scroll;
        height: 3rem;
        line-height: 3rem;
        font-size: 1.4rem;
        color: #ccc;
      }

    }
  }

</style>
