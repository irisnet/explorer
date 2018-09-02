<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">{{titleVar}}</span>
        <!--<span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>-->

        <span class="blocks_list_page_wrap_hash_var for_block"
              v-show="$route.query.block || $route.query.address">
          {{blockVar}}
        </span>
      </p>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination total_num">
        <span class="blocks_list_page_wrap_hash_var" v-show="['1','2','3','4'].includes(type)">{{count}} total</span>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
      <div style="position:relative;overflow-x: auto;">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="this.$route.params.type"
                           :minWidth="tableMinWidth"
                           :showNoData="showNoData"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination" style='margin:0.2rem 0;'>
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
        new Promise((resolve)=>{
          this.getDataList(currentPage, 30, this.$route.params.type);
          resolve();
        }).then(()=>{
          document.getElementById('router_wrap').scrollTop = 0;
        })

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

        };
        this.computeMinWidth();
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
        innerWidth : window.innerWidth,
        tableMinWidth:'',
      }
    },
    beforeMount() {
      this.type = this.$route.params.type;
      if (window.innerWidth > 910) {
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
      window.addEventListener('resize',this.onresize);
      this.computeMinWidth();
    },
    beforeDestroy() {
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
      onresize(){
        this.innerWidth = window.innerWidth;
        if(window.innerWidth > 910){
          this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
        } else {
          this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
        }
      },
      //根绝页面的不同展示最小宽度,不换行显示列表
      computeMinWidth(){
        if(this.$route.params.type === '1'){
          this.tableMinWidth = 6.5;
        }else if(this.$route.params.type === '2' && this.$route.params.param === 'transfer'){
          this.tableMinWidth = 4.7;
        }else if(this.$route.params.type === '3' || this.$route.params.type === '4'){
          this.tableMinWidth = 6.1;
        }
      },
      getDataList(currentPage, pageSize, type) {
        this.showLoading = true;
        if (type === '1') {
          let url = `/api/blocks/${currentPage}/${pageSize}`;
          if(this.$route.query.address){
            url = `/api/blocks/precommits/${this.$route.query.address}/${currentPage}/${pageSize}`;
            this.blockVar = `Precommit by ${this.$route.query.address}`;
          }
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
                  // Fees: '0 IRIS',
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
            if(this.$route.query.block){
              url = `/api/txsByBlock/${this.$route.query.block}/${currentPage}/${pageSize}`;
              this.blockVar = `for block ${this.$route.query.block}`;
            }else if(this.$route.query.address){
              url = `/api/txsByAddress/${this.$route.query.address}/${currentPage}/${pageSize}`;
              this.blockVar = `for ${this.$route.query.address}`;
            } else{
              url = `/api/txs/${currentPage}/${pageSize}`;
            }
          }
          axios.get(url).then((data) => {
            if (data.status === 200) {
              return data.data;
            }
          }).then((data) => {
            this.count = data.Count;
            if(data.Data){
              this.items = data.Data.map(item => {
                if(item.Amount.length > 0){
                  item.Amount[0].amount = Tools.dealWithFees(item.Amount[0].amount);
                }
                let [Amount,Fees] = ['',''];
                if(item.Amount instanceof Array){
                  Amount = item.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
                  if(item.Type === 'CompleteUnbonding' || item.Type === 'BeginUnbonding'){
                    Amount = item.Amount.map(listItem => `${listItem.amount}...shares`).join(',');
                  }
                }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
                  Amount = `${item.Amount.amount} ${item.Amount.denom.toUpperCase()}`;
                  if(item.Type === 'CompleteUnbonding' || item.Type === 'BeginUnbonding'){
                    Amount = `${item.Amount.amount}...shares`;
                  }
                }else if(item.Amount === null){
                  Amount = '';
                }
                if(item.ActualFee.amount && item.ActualFee.denom){
                  Fees = item.ActualFee.amount = Tools.dealWithFees(item.ActualFee.amount) + item.ActualFee.denom.toUpperCase();
                }
                return {
                  TxHash: item.TxHash,
                  Block:item.Height,
                  From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
                  To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
                  Type:item.Type === 'coin'?'transfer':item.Type,
                  Amount,
                  // Fees,
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
      @include flex;
      justify-content: flex-end;
      @include borderRadius(0.025rem);
      height:0.3rem;
      li{
        height:0.3rem !important;
        a{
          box-shadow: none;
        }
      }
    }
    .total_num{
      @include flex;
      justify-content: space-between;
      height:0.7rem;
      align-items: center;
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
      //min-width: 8rem;

      a {
        text-decoration: none;
      }
    }
    .blocks_list_title_wrap {
      width: 100%;
      border-bottom: 1px solid #d6d9e0 !important;
      @include flex;
      @include pcContainer;
      height:0.62rem;
      background:#efeff1;
      p{
        height:0.62rem;
        span{
          height:0.62rem;
      line-height:0.62rem;
        }
      }
      .personal_computer_blocks_list_page_wrap {
        @include flex;

      }
      .mobile_blocks_list_page_wrap {
        @include flex;
        flex-direction: column;
        overflow-x: auto;
        width:100%;
        .blocks_list_page_wrap_hash_var{
          min-width:7rem;
        }
      }

    }
    .personal_computer_blocks_list_page_wrap {
      .transaction_information_content_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #000000;
        margin-bottom: 0;
      }
      @include pcCenter;
      min-height:4.6rem;
      .transactions_detail_information_wrap {
        .information_props_wrap {
          @include flex;
          .information_props {
            width: 15rem;
          }
        }
      }

      .blocks_list_title {
        height: 0.62rem;
        line-height: 0.62rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        height:  0.62rem;
        line-height: 0.62rem;
        font-size: 0.22rem;
        color: #a2a2ae;
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
      padding: 0 0.1rem;
      .transaction_information_content_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #000000;
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
        color: #000000;
        margin-right: 0.2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        overflow-x: auto;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #a2a2ae;
      }
      .for_block{
        display:inline-block;
        //margin-left:0.1rem;
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
