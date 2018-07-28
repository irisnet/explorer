<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">{{titleVar}}</span>
        <span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>
        <span class="blocks_list_page_wrap_hash_var" v-show="['1','2','3','4'].includes(type)">{{count}} total</span>
        <span class="blocks_list_page_wrap_hash_var for_block"
              v-show="this.$route.params.param.includes('address') || this.$route.params.param.includes('block')">
          for {{blockVar}}
        </span>
      </p>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
      <div style="position:relative;min-height:3.36rem;">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="this.$route.params.type" :showNoData="showNoData"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
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
  import SpinComponent from './commonComponents/SpinComponent';

  export default {
    components:{
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        this.getDataList(currentPage, 30, this.$route.params.type);
      },
      $route() {
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = 1;
        this.getDataList(1, 30, this.$route.params.type);
        this.showNoData = false;
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
        showNoData:false,//是否显示列表的无数据
        showLoading:false,
        blockVar:'',

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
        this.showLoading = true;
        if (type === '1') {
          let url = `/api/blocks/${currentPage}/${pageSize}`;
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            this.count = data.Count;
            if(data.Data){
              this.items = data.Data.map(item => {
                let txn = item.NumTxs;
                let precommit = item.Block.LastCommit.Precommits.length;
                let [votingPower,denominator,numerator] = [0,0,0];
                item.Validators.forEach(listItem=>votingPower += listItem.VotingPower);
                item.Validators.forEach(item=>denominator += item.VotingPower);
                for(let i = 0; i < item.Block.LastCommit.Precommits.length; i++){
                  for (let j = 0; j < item.Validators.length; j++){
                    if(item.Block.LastCommit.Precommits[i].ValidatorAddress === item.Validators[j].Address){
                      numerator += item.Validators[j].VotingPower;
                      break;
                    }
                  }
                }
                return {
                  Height: item.Height,
                  Txn:txn,
                  Fees: '0 IRIS',
                  Timestamp: Tools.conversionTimeToUTC(item.Time),
                  'Precommit Validators':precommit,
                  'Voting Power': denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'',
                };
              })
            }else{
              this.items = [{Height:'',Txn:'',Fee:'',Timestamp:'','Precommit Validators':'','Voting Power':''}]
              this.showNoData = true;
            }
            this.showLoading = false;
          })
        } else if (type === '2') {
          let url;
          if(this.$route.params.param === 'transfer'){
            url = `/api/txs/coin/${currentPage}/${pageSize}`
          }else if(this.$route.params.param === 'stake'){
            url = `/api/txs/stake/${currentPage}/${pageSize}`
          }else if(this.$route.params.param === 'recent'){
            url = `/api/txs/${currentPage}/${pageSize}`;
          }else if(this.$route.params.param.includes('block')){
            url = `/api/txsByBlock/${this.$route.params.param.split(':')[1]}/${currentPage}/${pageSize}`;
            this.blockVar = this.$route.params.param.split(':')[1];
          }else if(this.$route.params.param.includes('address')){
            url = `/api/txsByAddress/${this.$route.params.param.split(':')[1]}/${currentPage}/${pageSize}`;
            this.blockVar = this.$route.params.param.split(':')[1];
          }
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            this.count = data.Count;
            if(data.Data){
              this.items = data.Data.map(item => {
                let [Amount,Fees] = ['',''];
                if(item.Amount instanceof Array){
                  Amount = item.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
                }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
                  Amount = `${item.Amount.amount} ${item.Amount.denom.toUpperCase()}`;
                }else if(item.Amount === null){
                  Amount = '';
                }
                if(item.Fee.Amount instanceof Array){
                  Fees = item.Fee.Amount.map(listItem=>`${listItem.amount} ${listItem.amount === 0?'IRIS':listItem.denom.toUpperCase()}`).join(',');
                }else if(item.Fee.Amount && Object.keys(item.Fee.Amount).includes('amount') && Object.keys(item.Fee.Amount).includes('denom')){
                  Fees = `${item.Fee.Amount} ${item.Fee.Amount.toUpperCase()}`;
                }else if(item.Fee.Amount === null){
                  Fees = '';
                }
                return {
                  TxHash: item.TxHash,
                  Block:item.Height,
                  From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
                  To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
                  Type:item.Type === 'coin'?'transfer':item.Type,
                  Amount,
                  Fees,
                  Timestamp: Tools.conversionTimeToUTC(item.Time),
                };
              })
            }else{
              this.items = [{
                TxHash: '',
                Block:'',
                From:'',
                To:'',
                Type:'',
                Amount:'',
                Fees:'',
                Timestamp:'',
              }];
              this.showNoData = true;
            }
            this.showLoading = false;
          })
        }else if (type === '3' || type === '4') {
          let url = `/api/stake/candidates/${currentPage}/${pageSize}`;
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            this.count = data.Count;
            if(data.Candidates){
              this.items = data.Candidates.map(item => {
                return {
                  Address: item.Address,
                  Name:item.Description.Moniker,
                  'Voting Power':`${(item.VotingPower/data.PowerAll*100).toFixed(2)}%`,
                  'Uptime':`${item.Uptime}%`,
                  'Commission Rate':'',
                  Returns:'',
                };
              })
            }else{
              this.showNoData = true;
              this.items = [{
                Address: '',
                Name:'',
                'Voting Power':'',
                'Uptime':'',
                'Commission Rate':'',
                Returns:'',
              }]
            }
            this.showLoading = false;
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
    font-size: 0.14rem;
    .pagination {
      margin-top: 0.05rem;
      margin-bottom: 0.05rem;
      @include flex;
      justify-content: flex-end;
      @include borderRadius(0.025rem);
    }
    .no_data_show{
      @include flex;
      justify-content: center;
      border-top:0.01rem solid #eee;
      border-bottom:0.01rem solid #eee;
      font-size:0.14rem;
      height:3rem;
      align-items: center;
    }
    .b-table {
      min-width: 7rem;

      a {
        text-decoration: none;
      }
    }
    .blocks_list_title_wrap {
      width: 100%;
      border-bottom: 0.01rem solid #eee;
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
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
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
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #555;
        margin-right: 0.2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.14rem;
        color: #ccc;
      }
      .for_block{
        display:inline-block;
        margin-left:0.1rem;
      }
    }

    .mobile_blocks_list_page_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.05rem;
      overflow-x: auto;
      .transaction_information_content_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #555;
        margin-bottom: 0;
      }
      .transactions_detail_information_wrap {

        .information_props_wrap {
          @include flex;
          flex-direction: column;
          border-bottom: 0.01rem solid #eee;
          margin-bottom: 0.05rem;
          .information_value {
            overflow-x: auto;
          }

        }
      }

      .blocks_list_title {
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #555;
        margin-right: 0.2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        overflow-x: auto;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.14rem;
        color: #ccc;
      }
      .for_block{
        display:inline-block;
        margin-left:0.1rem;
      }

    }
  }

  //重置bootstrap-vue的表格样式
  table{

    td{
      max-width:2.2rem !important;
      overflow-wrap: break-word !important;
    }
  }
  .page-item{
    &:first-child, &:last-child{
      .page-link{
        @include borderRadius(0.025rem);
      }
    }
  }

</style>
