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
          <span class="information_value">{{timestampValue}}</span>
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
              @click="tabTxList(index,item.txTabName,1,20)">{{item.txTabName}}
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
    components: {
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        new Promise((resolve)=>{
          this.getDataList(this.currentTabIndex,this.currentTxTabName,currentPage, this.pageSize);
          resolve();
        }).then(()=>{
          document.getElementById('router_wrap').scrollTop = 0;
        })
      },
      $route() {
        this.getBlockInformation();
        if (Number(this.$route.params.height) <= 0) {
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
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
        hashValue: '',
        heightValue: '',
        timestampValue: '',
        blockHashValue: '',
        transactionsValue: '',
        // feeValue: '',
        lastBlockHashValue: '',
        precommitValidatorsValue: '',
        votingPowerValue: '',
        items: [],
        showNoData: false,//列表无数据的时候显示
        acitve: true,
        activeNext: true,
        maxBlock: 0,
        ProposalsTransactionsNum: 0,
        txTabName:"Transfers",
        tabTxListIndex:0,
        count: 0,
        showLoading:false,
        currentPage: 1,
        pageSize: 20,
        addressTxList: "",
        txTab:[
          {
            "txTabName":"Transfers",
            "active": true
          },
          {
            "txTabName":"Stakes",
            "active": false

          },
          {
            "txTabName":"Declarations",
            "active": false
          },
          {
            "txTabName":"Governance",
            "active": false
          }
        ]
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      this.tabTxList(this.tabTxListIndex,this.txTabName,this.currentPage,this.pageSize);
      this.getBlockInformation();
      if (Number(this.$route.params.height) <= 0) {
        this.acitve = false;
      } else {
        this.acitve = true;
      }
      this.getMaxBlock();
    },
    methods: {
      tabTxList(index,txTabName,currentPage,pageSize){
        this.currentPage = currentPage;
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
        axios.get(url).then((data) => {
          if(data.status === 200){
            return data.data;
          }
        }).then((data) => {
          this.showLoading = false;
          this.showNoData = false;
          if(data.Data){
            that.items = data.Data.map(item => {
              let [Amount,Fee] = ['',''];
              if(txTabName === 'Transfers' || txTabName === 'Stakes' || txTabName === 'Governance'){
                if(item.Amount !== null && item.Amount.length > 0){
                  item.Amount[0].amount = Tools.dealWithFees(item.Amount[0].amount);
                }
                if(item.Amount instanceof Array){
                  Amount = item.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
                  if(item.Type === 'CompleteUnbonding' || item.Type === 'BeginUnbonding'){
                    Amount = item.Amount.map(listItem => `${listItem.amount}shares`).join(',');
                  }
                }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
                  Amount = `${item.Amount.amount} ${item.Amount.denom.toUpperCase()}`;
                  if(item.Type === 'CompleteUnbonding' || item.Type === 'BeginUnbonding'){
                    Amount = `${item.Amount.amount}shares`;
                  }
                }else if(item.Amount === null){
                  Amount = '';
                }
                if(item.Fee.amount && item.Fee.denom){
                  Fee = item.Fee.amount = Tools.formatFeeToFixedNumber(item.Fee.amount) + item.Fee.denom.toUpperCase();

                }
              }
              let objList;
              if(txTabName === 'Transfers'){
                objList = {
                  TxHash: item.Hash,
                  Block:item.BlockHeight,
                  From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
                  To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
                  Amount,
                  Fee,
                  Timestamp: Tools.conversionTimeToUTC(item.Timestamp),
                }
              }else if(txTabName === 'Stakes'){
                objList = {
                  TxHash: item.Hash,
                  Block:item.BlockHeight,
                  From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
                  To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
                  Type:item.Type === 'coin'?'transfer':item.Type,
                  Amount,
                  Fee,
                  Timestamp: Tools.conversionTimeToUTC(item.Timestamp),
                };
              }else if(txTabName === 'Declarations'){
                objList = {
                  TxHash: item.Hash,
                  Block:item.BlockHeight,
                  Owner:item.Owner,
                  Moniker: item.Moniker,
                  'Self-Bond': item.SelfBond,
                  Type: item.Type,
                  Fee,
                  Timestamp: Tools.conversionTimeToUTC(item.Timestamp),
                };
              }else if(txTabName === 'Governance'){
                objList = {
                  TxHash: item.Hash,
                  Block:item.BlockHeight,
                  From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
                  "Proposal_ID": item.ProposalId,
                  Type:item.Type === 'coin'?'transfer':item.Type,
                  Fee,
                  Timestamp: Tools.conversionTimeToUTC(item.Timestamp),
                };
              }
              return objList
            })

          }else {
            this.showNoData = true;
            if(txTabName === 'Transfers'){
              this.items = [{
                'Tx Hash': "",
                Block: "",
                From: "",
                To: "",
                Amount: "",
                Fee: "",
                Timestamp: "",
              }]
            }else if(txTabName === 'Stakes'){
              this.items = [{
                'Tx Hash': "",
                Block: "",
                From: "",
                To: "",
                Type: "",
                Amount: "",
                Fee: "",
                Timestamp: "",
              }]
            }else if(txTabName === 'Declarations'){
              this.items = [{
                'Tx Hash': "",
                Block: "",
                Owner: "",
                Moniker: "",
                'Self-Bond': "",
                Type: "",
                Fee: "",
                Timestamp: "",
              }]
            }else if(txTabName === 'Governance'){
              this.items = [{
                'Tx Hash': "",
                Block: "",
                From: "",
                'Proposal_ID': "",
                Type: "",
                Fee: "",
                Timestamp: "",
              }]
            }
          }
        })
      },
      getBlockInformation() {
        let url = `/api/block/${this.$route.params.height}`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if (data) {
            let denominator = 0;
            data.Validators.forEach(item => denominator += item.VotingPower);
            let numerator = 0;
            for (let i = 0; i < data.Block.LastCommit.Precommits.length; i++) {
              for (let j = 0; j < data.Validators.length; j++) {
                if (data.Block.LastCommit.Precommits[i].ValidatorAddress === data.Validators[j].Address) {
                  numerator += data.Validators[j].VotingPower;
                  break;
                }
              }
            }
            if (data) {
              this.transactionsValue = data.NumTxs;
              this.hashValue = data.Height;
              this.heightValue = data.Height;
              this.timestampValue = Tools.conversionTimeToUTC(data.Time);
              this.blockHashValue = data.Hash;
              // this.feeValue = '0 IRIS';
              this.lastBlockHashValue = data.Block.LastCommit.BlockID.Hash;
              this.precommitValidatorsValue = data.Validators.length !== 0 ? `${data.Block.LastCommit.Precommits.length}/${data.Validators.length}` : '';
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
            this.feeValue = '0 IRIS';
            this.lastBlockHashValue = '';
            this.precommitValidatorsValue = '';
            this.votingPowerValue = '';
          }
        }).catch(e => {
          console.log(e)
        })
      },
      skipNext(num) {
        if (Number(this.$route.params.height) <= 0) {
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
        let url = `/api/blocks/1/1`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if (data && typeof data === "object") {
            if(data.Data && data.Data.length !==0){
              this.maxBlock = data.Data[0].Height;
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
      .transaction_information_content_title {
        height: 0.5rem !important;
        line-height: 0.5rem !important;
        font-size: 0.18rem !important;
        color: #000000;
        margin-bottom: 0;
        border-bottom:1px solid #d6d9e0 !important;
      }
      @include pcCenter;
      .transactions_detail_information_wrap {
        .information_props_wrap {
          @include flex;
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
        font-weight: 500;
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
        font-weight: 500;
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
          z-index: 5;
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
