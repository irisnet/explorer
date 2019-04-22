<template>
  <div class="transactions_detail_wrap">
    <div class="transactions_title_wrap">
      <p :class="transactionsDetailWrap" style="margin-bottom:0;">
        <span class="transactions_detail_title">Block</span>
        <span class="transactions_detail_wrap_hash_var">{{`#${hashValue}`}}</span>
      </p>
    </div>

    <div :class="transactionsDetailWrap">
      <p class="transaction_information_content_title">Block Information</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Height:</span>
          <span>
            <i :class="acitve?'flag_item_left':'flag_item_left_disabled'" @click="skipNext(-1)"></i>
            <span class="information_value" style="flex:none;">{{heightValue}}</span>
            <i :class="activeNext?'flag_item_right':'flag_item_right_disabled'" @click="skipNext(1)"></i>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Timestamp:</span>
          <span class="information_value">{{ageValue}} ({{timestampValue}})</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Block Hash:</span>
          <span class="information_value">{{blockHashValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Transactions:</span>
          <div class="information_value">
            <span>{{transactionsValue}}</span>
          </div>
        </div>
        <p class="transaction_information_content_title">Last Block</p>
        <div class="last_block_container">
          <div class="information_props_wrap">
            <span class="information_props">Last Block Hash:</span>
            <span class="information_value">{{lastBlockHashValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Precommit Validators:</span>
            <span class="information_value">{{precommitValidatorsValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Voting Power:</span>
            <span class="information_value">{{votingPowerValue}}</span>
          </div>
        </div>
      </div>
    </div>
    <div class="list_tab_wrap" :class="transactionsDetailWrap">
      <div class="list_tab_content">
        <ul class="list_tab_container">
          <li class="list_tab_item"
              :class="item.active ? 'activeStyle' : 'unactiveStyle'" v-for="(item,index) in txTab"
              @click="tabTxList(index,item.txTabName,1,20)">{{item.txTabName}} ({{item.BlockTxStatistics}})
          </li>
        </ul>
      </div>
    </div>
    <div :class="transactionsDetailWrap">
      <div class="pagination total_num">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
      <div class="blocks_list_table_contianer">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="'blockTxList'"
                           :showNoData="showNoData" :min-width="tableMinWidth"></blocks-list-table>
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
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  import Service from "../util/axios";
  import Constant from "../constant/Constant"
  export default {
    components: {
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        new Promise((resolve)=>{
          this.tabTxList(this.currentTabIndex,this.currentTxTabName,currentPage, this.pageSize);
          resolve();
        }).then(()=>{
          Tools.scrollToTop()
        })
      },
      $route() {
        this.getBlockInformation();
        this.getBlockTxStatistics();
        this.computeMinWidth();
        this.tabTxList(this.tabTxListIndex,this.txTabName,this.currentPage,this.pageSize);
        if (Number(this.$route.params.height) <= 1) {
          this.acitve = false;
        } else {
          this.acitve = true;
        }
        if (this.maxBlock !== 0) {
          if (Number(this.$route.params.height) >= this.maxBlock) {
            this.activeNext = false;
          } else {
            this.activeNext = true;
          }
        }
      }
    },
    data() {
      return {
        transactionsTimer:null,
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',
        hashValue: '',
        heightValue: '',
        timestampValue: '',
        blockHashValue: '',
        transactionsValue: '',
        lastBlockHashValue: '',
        precommitValidatorsValue: '',
        votingPowerValue: '',
        items: [],
        showNoData: false,
        acitve: true,
        activeNext: true,
        maxBlock: 0,
        ProposalsTransactionsNum: 0,
        txTabName:"Transfers",
        tabTxListIndex:0,
        currentTabIndex:"",
        currentTxTabName:"",
        count: 0,
        showLoading:false,
        currentPage: 1,
        pageSize: 20,
        addressTxList: "",
        tableMinWidth:"",
        blockDetailTimer: null,
        ageValue: "",
        txTab:[
          {
            "txTabName":"Transfers",
            "active": true,
            BlockTxStatistics:"",
          },
          {
            "txTabName":"Stakes",
            "active": false,
            BlockTxStatistics:"",
          },
          {
            "txTabName":"Declarations",
            "active": false,
            BlockTxStatistics:"",
          },
          {
            "txTabName":"Governance",
            "active": false,
            BlockTxStatistics:"",
          }
        ]
      }
    },
    beforeMount() {
      Tools.scrollToTop();
      if (Tools.currentDeviceIsPersonComputer()) {
        this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      this.tabTxList(this.tabTxListIndex,this.txTabName,this.currentPage,this.pageSize);
      this.getBlockInformation();
      if (Number(this.$route.params.height) <= 1) {
        this.acitve = false;
      } else {
        this.acitve = true;
      }
      this.getMaxBlock();
      this.getBlockTxStatistics();
      this.computeMinWidth();
    },
    methods: {
      computeMinWidth(){
        if(this.$route.params.height){
          this.tableMinWidth = 8.8;
        }
      },
      getBlockTxStatistics(){
        let url = `/api/txs/statistics?height=${this.$route.params.height}`;
        Service.http(url).then((data) => {
          if(data){
            this.txTab[0].BlockTxStatistics = data.TransCnt;
            this.txTab[1].BlockTxStatistics = data.StakeCnt;
            this.txTab[2].BlockTxStatistics = data.DeclarationCnt;
            this.txTab[3].BlockTxStatistics = data.GovCnt;
          }
        }).catch((e) => {
          console.log(e)
        })
      },
      tabTxList(index,txTabName,currentPage,pageSize){
        this.currentPage = currentPage;
        this.currentTabIndex = index;
        this.currentTxTabName = txTabName;
        this.showLoading = true;
        for (let txTabIndex = 0; txTabIndex < this.txTab.length; txTabIndex++){
          this.txTab[txTabIndex].active = false;
          this.txTab[index].active = true;
        }
        let url;
        let that = this;
        if(txTabName === 'Transfers'){
          url = `/api/tx/trans/${currentPage}/${pageSize}?height=${this.$route.params.height}`
        }else if(txTabName === 'Stakes'){
          url = `/api/tx/stake/${currentPage}/${pageSize}?height=${this.$route.params.height}`

        }else if(txTabName === 'Declarations'){
          url = `/api/tx/declaration/${currentPage}/${pageSize}?height=${this.$route.params.height}`
        }else if(txTabName === 'Governance'){
          url = `/api/tx/gov/${currentPage}/${pageSize}?height=${this.$route.params.height}`
        }
        Service.http(url).then((txList) => {
          that.showLoading = false;
          that.showNoData = false;
          that.count = txList.Count;
          clearInterval(this.transactionsTimer);
          if(txList.Data){
            this.transactionsTimer = setInterval(function () {
              let currentServerTime = new Date().getTime() + that.diffMilliseconds;
              that.items = Tools.formatTxList(txList.Data,txTabName,currentServerTime)
            },1000);
          }else {
            that.showNoData = true;
            let currentServerTime = new Date().getTime() + that.diffMilliseconds;
            that.items = Tools.formatTxList(null,txTabName,currentServerTime)
          }
        })
      },
      getBlockInformation() {
        let url = `/api/block/${this.$route.params.height}`;
        Service.http(url).then((data) => {
          clearInterval(this.blockDetailTimer);
          if (data) {
            let denominator = 0;
            data.validators.forEach(item => denominator += item.voting_power);
            let numerator = 0;
            if(data.last_commit && data.last_commit.length !== 0 ){
              for (let i = 0; i < data.last_commit.length; i++) {
                for (let j = 0; j < data.validators.length; j++) {
                  if (data.last_commit[i] === data.validators[j].address) {
                    numerator += data.validators[j].voting_power;
                    break;
                  }
                }
              }
            }
            if (data) {
              let that = this;
              this.transactionsValue = data.total_txs;
              this.hashValue = data.hash;
              this.heightValue = data.height;
              this.blockDetailTimer = setInterval(function () {
                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                that.ageValue = Tools.formatAge(currentServerTime,data.time,Constant.SUFFIX);
              },1000);
              this.timestampValue = Tools.format2UTC(data.time);
              this.blockHashValue = data.hash;
              this.lastBlockHashValue = data.last_block_hash;
              this.precommitValidatorsValue = data.validators.length !== 0 ? `${data.last_commit.length}/${data.validators.length}` : '';
              this.votingPowerValue = denominator !== 0 ? `${numerator / denominator * 100}%` : '';
            }
          } else {
            this.items = [{
              Address: '',
              Index: '',
              Round: '',
              Signature: '',
            }];
            this.showNoData = true;
            this.hashValue = this.$route.params.height;
            this.heightValue = this.$route.params.height;
            this.timestampValue = '';
            this.blockHashValue = '';
            this.transactionsValue = '';
            this.lastBlockHashValue = '';
            this.precommitValidatorsValue = '';
            this.votingPowerValue = '';
          }
        }).catch(e => {
          console.log(e)
        })
      },
      skipNext(num) {
        if (Number(this.$route.params.height) <= 1) {
          this.acitve = false;
          if (num !== -1) {
            this.$router.push(`/blocks_detail/${Number(this.$route.params.height) + num}`)
          }
        } else if (Number(this.$route.params.height) >= this.maxBlock) {
          if (num !== 1) {
            this.$router.push(`/blocks_detail/${Number(this.$route.params.height) + num}`)
          }
        } else {
          this.acitve = true;
          this.$router.push(`/blocks_detail/${Number(this.$route.params.height) + num}`)
        }
      },
      getMaxBlock() {
        let url = `/api/blocks?page=1&size=1`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if (data && typeof data === "object") {
            if(data.Data && data.Data.length !==0){
              this.maxBlock = data.Data[0].height;
            }
            if (Number(this.$route.params.height) >= this.maxBlock) {
              this.activeNext = false;
            } else {
              this.activeNext = true;
            }
          }
        }).catch(e => {
          console.log(e)
        })
      },
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .transactions_detail_wrap {
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
      }
    }
    .transactions_title_wrap {
      width: 100%;
      border-bottom: 1px solid #d6d9e0;
      @include flex;
      @include pcContainer;
      .personal_computer_transactions_detail_wrap {
        @include flex;
      }
      .mobile_transactions_detail_wrap {
        @include flex;
        flex-direction: column;
      }

    }
    .personal_computer_transactions_detail_wrap {
      padding-bottom: 0.2rem;
      .transaction_information_content_title {
        @include fontWeight;
        height: 0.5rem !important;
        line-height: 0.5rem !important;
        font-size: 0.18rem !important;
        color: #000000;
        margin-bottom: 0;
        padding-left: 0.2rem !important;
        border-bottom:1px solid #d6d9e0 !important;
      }
      @include pcCenter;
      .transactions_detail_information_wrap {
        .information_props_wrap {
          @include flex;
          padding-left: 0.2rem;
          margin-bottom:0.08rem;
          .information_props {
            width: 1.5rem;
          }
          .flag_item_left {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            background: url('../assets/left.png') no-repeat 0 1px;
            margin-right: 0.05rem;
            cursor: pointer;
          }
          .flag_item_left_disabled {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            margin-right: 0.05rem;
            cursor: pointer;
            background: url('../assets/left_disabled.png') no-repeat 0 1px;
          }
          .flag_item_right {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            background: url('../assets/right.png') no-repeat 0 0;
            margin-left: 0.05rem;
            cursor: pointer;
          }
          .flag_item_right_disabled {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            background: url('../assets/right_disabled.png') no-repeat 0 0;
            margin-left: 0.05rem;
            cursor: pointer;
          }

        }
      }
      .block_detail_table_wrap {
        width: 100%;
        margin-bottom:0.5rem;
        .no_data_show {
          @include flex;
          justify-content: center;
          border-top: 0.01rem solid #eee;
          border-bottom: 0.01rem solid #eee;
          font-size: 0.14rem;
          height: 3rem;
          align-items: center;
        }
      }

      .transactions_detail_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
        padding-left: 0.2rem;
        @include fontWeight;
      }
      .transactions_detail_wrap_hash_var {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.22rem;
        color: #a2a2ae;
      }
    }

    .mobile_transactions_detail_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.1rem;
      .transaction_information_content_title {
        height: 0.5rem !important;
        line-height: 0.5rem !important;
        font-size: 0.18rem !important;
        color: #000000;
        margin-bottom: 0;
        @include fontWeight;
        border-bottom: 0.01rem solid #d6d9e0;
      }
      .block_detail_table_wrap {
        width: 100%;
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
        margin-bottom:0.4rem;
        .no_data_show {
          @include flex;
          justify-content: center;
          border-top: 0.01rem solid #eee;
          border-bottom: 0.01rem solid #eee;
          font-size: 0.14rem;
          height: 3rem;
          align-items: center;
        }
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
          .flag_item_left {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            background: url('../assets/left.png') no-repeat 0 1px;
            margin-right: 0.05rem;
            cursor: pointer;
          }
          .flag_item_left_disabled {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            margin-right: 0.05rem;
            cursor: pointer;
            background: url('../assets/left_disabled.png') no-repeat 0 1px;
          }
          .flag_item_right {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            background: url('../assets/right.png') no-repeat 0 0;
            margin-left: 0.05rem;
            cursor: pointer;
          }
          .flag_item_right_disabled {
            display: inline-block;
            width: 0.2rem;
            height: 0.17rem;
            background: url('../assets/right_disabled.png') no-repeat 0 0;
            margin-left: 0.05rem;
            cursor: pointer;
          }
        }
      }
      .transactions_detail_title {
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.02rem;
        @include fontWeight;
      }
      .transactions_detail_wrap_hash_var {
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.22rem;
        color: #a2a2ae;
      }
    }

  }
  .last_block_container{
    padding-top: 0.2rem;
  }

  .list_tab_wrap{
    width: 100%;
    padding-top: 0.44rem;
    padding-bottom: 0.2rem;
    .list_tab_content{
      width: 100%;
      border-bottom: 0.01rem solid #3598db;
      .list_tab_container{
        @include flex;
        height: 0.38rem;
        min-width: 4rem;
        max-width: 12.8rem;
        margin-left: 0.2rem;
        .list_tab_item{
          position: relative;
          top: 0.01rem;
          text-align: center;
          line-height: 0.38rem;
          width: 1.54rem;
          color: #A2A2AE;
          border-left: 0.01rem solid #e4e4e4;
          border-top: 0.01rem solid #e4e4e4;
          border-right: 0.01rem solid #e4e4e4;
          border-bottom: 0.01rem solid #3598db;
          cursor: pointer;
        }
      }
    }
    .activeStyle{
      color: #3598db!important;
      border-top: 0.01rem solid #3598db !important;
      border-right: 0.01rem solid #3598db !important;
      border-left: 0.01rem solid #3598db !important;
      border-bottom: 0.01rem solid #fff !important;
    }
  }
  .blocks_list_table_contianer{
    position:relative;
    overflow-x: auto;
    -webkit-overflow-scrolling:touch;
    top: 0.2rem;
    padding-bottom: 0.2rem;
  }

</style>
