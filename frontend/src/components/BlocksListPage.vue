<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">{{listTitleName}}</span>
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
      <div style="position:relative;overflow-x: auto;-webkit-overflow-scrolling:touch;">
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
  import Tools from '../util/Tools';
  import axios from 'axios';
  import Service from "../util/axios"
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
        clearInterval(this.timer);
        clearInterval(this.transactionTimer);
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = 1;
        this.getDataList(1, 30, this.$route.params.type);
        this.showNoData = false;
        switch (this.$route.params.type) {
          case '1': this.listTitleName = 'Blocks';
                  break;
        }
        this.computeMinWidth();
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',
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
        listTitleName:"",
        timer: null,
        transactionTimer: null,
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
        case '1': this.listTitleName = 'Blocks';
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
          this.tableMinWidth = 8.5;
        }else if(this.$route.params.type === '2' && this.$route.params.param === 'transfer'){
          this.tableMinWidth = 12;
        }else if(this.$route.params.type === '3' || this.$route.params.type === '4'){
          this.tableMinWidth = 8.1;
        }
      },
      getBlockList(currentPage, pageSize){
        let url = `/api/blocks/${currentPage}/${pageSize}`;
        Service.http(url).then((data) => {
          if(data){
            let that = this;
            clearInterval(this.timer);
            this.timer = setInterval(function () {
              that.count = data.Count;
              that.items = data.Data.map(item => {
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
                  Age: Tools.formatAge(item.Time),
                  'Precommit Validators':precommit,
                  'Voting Power': denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'',
                };
              })
            },1000)

          }else{
            this.items = [{Height:'',Txn:'',Fee:'',Age:'','Precommit Validators':'','Voting Power':''}];
            this.showNoData = true;
          }
          this.showLoading = false;
        }).catch(e => {
          console.log(e)
        })
      },
      getTransactionList(currentPage, pageSize){
        let url;
        let that = this;
        if(this.$route.params.param === 'Transfers'){
          this.listTitleName = "Transfers";
          url = `/api/tx/trans/${currentPage}/${pageSize}`
        }else if(this.$route.params.param === 'Stakes'){
          this.listTitleName = "Stakes";
          url = `/api/tx/stake/${currentPage}/${pageSize}`

        }else if(this.$route.params.param === 'Declarations'){
          this.listTitleName = "Declarations";
          url = `/api/tx/declaration/${currentPage}/${pageSize}`
        }else if(this.$route.params.param === 'Governance'){
          this.listTitleName = "Governance";
          url = `/api/tx/gov/${currentPage}/${pageSize}`
        }
        Service.http(url).then((data) => {
          that.count = data.Count;
          clearInterval(this.transactionTimer);
          if(data){
            that.transactionTimer = setInterval(function () {
              that.items = Tools.formatTxList(data.Data,that.$route.params.param)
            },1000);
          }else{
            that.items = Tools.formatTxList(null,that.$route.params.param);
            that.showNoData = true;
          }
          that.showLoading = false;
        })
      },
      getValidatorList(currentPage, pageSize){
        let url;
        if(this.$route.params.param === "active"){
          this.listTitleName = "Active Validators";
          url = `/api/stake/validators/${currentPage}/${pageSize}`;
        }else if(this.$route.params.param === "jailed"){
          this.listTitleName = "Jailed Validators";
          url = `/api/stake/revokedVal/${currentPage}/${pageSize}`;
        }else if(this.$route.params.param === "candidates"){
          this.listTitleName = "Validator Candidates";
          url = `/api/stake/candidates/${currentPage}/${pageSize}`;
        }
        Service.http(url).then((data) => {
          if(data){
            this.count = data.Count;
            if(this.$route.params.param === "active"){
              if(data.Candidates){
                this.items = data.Candidates.map(item => {
                  return {
                    Address: item.Address,
                    Name:Tools.formatString(item.Description.Moniker,20,"..."),
                    'Voting Power':`${Tools.formatStringToFixedNumber(Tools.formatStringToNumber(item.OriginalTokens),2)} (${(item.OriginalTokens/data.PowerAll*100).toFixed(2)}%)`,
                    'Uptime':`${item.Uptime}%`,
                    'Bond Height': item.BondHeight
                  };
                })
              }else{
                this.showNoData = true;
                this.items = [{
                  Address: '',
                  Name:'',
                  'Voting Power':'',
                  'Uptime':'',
                }]
              }
            }else if(this.$route.params.param === "jailed"){
              if(data.Candidates){
                this.items = data.Candidates.map(item => {
                  return {
                    Address: item.Address,
                    Name:Tools.formatString(item.Description.Moniker,20,"..."),
                    'Voting Power':Tools.formatStringToFixedNumber(Tools.formatStringToNumber(item.OriginalTokens),2),
                  };

                })
              }else{
                this.showNoData = true;
                this.items = [{
                  Address: '',
                  Name:'',
                  'Voting Power':'',
                }]
              }
            }else if(this.$route.params.param === "candidates"){
              if(data.Candidates){
                this.items = data.Candidates.map(item => {
                  return {
                    Address: item.Address,
                    Name: Tools.formatString(item.Description.Moniker,20,"..."),
                    'Voting Power': Tools.formatStringToFixedNumber(Tools.formatStringToNumber(item.OriginalTokens),2),
                    'Declare Height' : item.BondHeight
                  };
                })
              }else{
                this.showNoData = true;
                this.items = [{
                  Address: '',
                  Name:'',
                  'Voting Power':'',
                  'Declare Height' :""
                }]
              }
            }
          }
          this.showLoading = false;
        }).catch(e => {
          this.showLoading = false;
          this.showNoData = true;
          console.log(e)
        })
      },
      getDataList(currentPage, pageSize, type) {
        this.showLoading = true;
        if (type === '1') {
          this.getBlockList(currentPage, pageSize);
        }else if (type === '2') {
          this.getTransactionList(currentPage, pageSize)
        }else if (type === '3' || type === '4') {
          this.getValidatorList(currentPage, pageSize)
        }

      }
    }
  }
</script>
<style lang="scss" >
  @import '../style/mixin.scss';

  .blocks_list_page_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    width: 100%!important;
    .pagination {
      @include flex;
      justify-content: flex-end;
      @include borderRadius(0.025rem);
      height:0.3rem;
      li{
        height:0.3rem !important;
      }
    }
    .total_num{
      @include flex;
      justify-content: space-between;
      height:0.7rem !important;
      align-items: center;
      .blocks_list_page_wrap_hash_var{
        padding-left: 0.2rem;
      }
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
        padding-bottom: 0.2rem;
      }
      .mobile_blocks_list_page_wrap {
        @include flex;
        flex-direction: column;
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
        width:100%;
        .blocks_list_page_wrap_hash_var{
          min-width:7rem;
          margin-left: 0.2rem!important;
        }
      }

    }
    .personal_computer_blocks_list_page_wrap {
      padding-bottom: 0.2rem;
      width: 100%!important;
      .transaction_information_content_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #000000;
        margin-bottom: 0;
        @include fontWeight;
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
        margin-left: 0.2rem;
        height: 0.62rem;
        line-height: 0.62rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
        @include fontWeight;
      }
      .blocks_list_page_wrap_hash_var {
        height:  0.62rem;
        line-height: 0.62rem;
        font-size: 0.18rem;
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
        @include fontWeight;
      }
      .transactions_detail_information_wrap {

        .information_props_wrap {
          @include flex;
          flex-direction: column;
          border-bottom: 0.01rem solid #eee;
          margin-bottom: 0.05rem;
          .information_value {
            overflow-x: auto;
            -webkit-overflow-scrolling:touch;
          }

        }
      }

      .blocks_list_title {
        height: 0.6rem !important;
        line-height: 0.6rem !important;
        font-size: 0.18rem;
        color: #000000;
        margin-right: 0.2rem;
        padding-left: 0.2rem;
        @include fontWeight;
      }
      .blocks_list_page_wrap_hash_var {
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
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
      word-wrap: break-word !important;
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
