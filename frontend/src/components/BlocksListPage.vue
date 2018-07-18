<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">

      <p :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">{{titleVar}}</span>
        <span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>
      </p>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
      <blocks-list-table :items="items" :type="this.$route.params.type"></blocks-list-table>
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
  import BlocksListTable from './table/BlocksListTable.vue';

  export default {
    components:{
      BlocksListTable,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        this.getDataList(currentPage, 30, this.$route.params.type);
      },
      $route() {
        this.items = [];
        this.getDataList(1, 30, this.$route.params.type);
        switch (this.$route.params.type) {
          case '1': this.titleVar = 'Blocks';
                  break;
          case '2': this.titleVar = 'Transactions';
                  break;
          case '3': this.titleVar = 'Validators';
                  break;
          case '4': this.titleVar = 'Candidates';
                  break;

        }
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',//1是显示pc端，0是移动端
        blocksValue: '',
        currentPage: 1,
        pageSize: 30,
        count: 0,
        fields: [],
        items: [],
        type: '1',
        titleVar: '',


      }
    },
    beforeMount() {
      this.type = this.$route.params.type;
      if (Tools.currentDeviceIsPersonComputer()) {
        this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
      } else {
        this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
      }
    },
    mounted() {
      //组件页面根据路由类型判断是哪个页面
      console.log(this.$route.params)


      this.getDataList(1, 30, this.$route.params.type);
      switch (this.$route.params.type) {
        case '1': this.titleVar = 'Blocks';
          break;
        case '2': this.titleVar = 'Transactions';
          break;
        case '3': this.titleVar = 'Validators';
          break;
        case '4': this.titleVar = 'Candidates';
          break;

      }
    },
    methods: {
      getDataList(currentPage, pageSize, type) {
        console.log(type)
        if (type === '1') {
          let url = `/api/blocks/${currentPage}/${pageSize}`;
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            this.count = data.Count;
            this.items = data.Data.map(item => {
              let txn = item.NumTxs;
              let precommit = item.Block.LastCommit.Precommits.length;
              return {
                Height: item.Height,
                Txn:txn,
                Fee: '',
                Timestamp: item.Time,
                'Precommit Validators':precommit,
                'Voting Power': '',
              };
            })
          })
        } else if (type === '2') {
          let url = `/api/txs/${currentPage}/${pageSize}`;
          if(this.$route.params.param === 'transfer'){
            url = `/api/txs/coin/${currentPage}/${pageSize}`
          }else if(this.$route.params.param === 'stake'){
            url = `/api/txs/stake/${currentPage}/${pageSize}`
          }
          console.log(url)
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            this.count = data.Count;
            this.items = data.Data.map(item => {
              let [Amount,Fees] = ['',''];
              if(item.Amount instanceof Array){
                Amount = item.Amount.map(listItem=>`${listItem.amount} ${listItem.denom}`).join(',');
              }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
                Amount = `${item.Amount.amount} ${item.Amount.denom}`;
              }else if(item.Amount === null){
                Amount = '';
              }
              if(item.Fee.Amount instanceof Array){
                Fees = item.Fee.Amount.map(listItem=>`${listItem.amount} ${listItem.denom}`).join(',');
              }else if(item.Fee.Amount && Object.keys(item.Fee.Amount).includes('amount') && Object.keys(item.Fee.Amount).includes('denom')){
                Fees = `${item.Fee.Amount} ${item.Fee.Amount}`;
              }else if(item.Fee.Amount === null){
                Fees = '';
              }
              return {
                TxHash: item.TxHash,
                Block:item.Height,
                From:item.From,
                To:item.To,
                Type:item.Type,
                Amount,
                Fees,
                Timestamp: item.Time,
              };
            })
          })
        }else if (type === '3' || type === '4') {
          let url = `/api/stake/candidates/${currentPage}/${pageSize}`;
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            console.log(data)
            this.count = data.Count;
            this.items = data.Data.map(item => {
              return {
                Address: item.Address,
                Name:item.Description.Moniker,
                'Voting Power':item.VotingPower,
                'Uptime':item.UpdateTime,
                'Commission Rate':'',
                Returns:'',
              };
            })
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
    .pagination {
      margin-top: 0.5rem;
      @include flex;
      justify-content: flex-end;
    }
    .b-table {
      border-bottom: 1px solid #eeeeee;
      min-width: 50rem;

      a {
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

  //重置bootstrap-vue的表格样式
  table{

    td{
      max-width:20rem !important;
      overflow-wrap: break-word !important;
    }
  }

</style>
