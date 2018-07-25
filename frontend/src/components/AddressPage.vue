<template>
  <div class="transactions_detail_wrap">
    <div class="transactions_title_wrap">
      <p :class="transactionsDetailWrap" style="margin-bottom:0;">
        <span class="transactions_detail_title">Address</span>
        <span class="transactions_detail_wrap_hash_var">{{hashValue}}</span>
      </p>
    </div>

    <div :class="transactionsDetailWrap">
      <p class="transaction_information_content_title">Address Information</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Balance:</span>
          <span class="information_value">{{balanceValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Deposits:</span>
          <span class="information_value">{{depositsValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Transactions:</span>
          <span class="information_value">{{transactionsValue}}</span>
        </div>
      </div>
    </div>
    <div :class="transactionsDetailWrap" class="address_profile">
      <p class="transaction_information_content_title">Profile</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Name:</span>
          <span class="information_value">{{nameValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Pub Key:</span>
          <span class="information_value">{{pubKeyValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Website:</span>
          <span class="information_value">{{websiteValue}}</span>
        </div>
        <div class="information_props_wrap" style="border-bottom:0.01rem solid #eee">
          <span class="information_props">Description:</span>
          <span class="information_value">{{descriptionValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Commission Rate:</span>
          <span class="information_value">{{commissionRateValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Announcement:</span>
          <span class="information_value">{{announcementValue}}</span>
        </div>


      </div>
    </div>
    <div :class="transactionsDetailWrap" class="current_tenure">
      <p class="transaction_information_content_title" style="border-bottom:1px solid #eee">Current Tenure</p>
      <div class="current_tenure_wrap">
        <div class="transactions_detail_information_wrap">
          <div class="information_props_wrap">
            <span class="information_props">Bond Height:</span>
            <span class="information_value">{{bondHeightValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Voting Power:</span>
            <span class="information_value">{{votingPowerValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Precommited Blocks:</span>
            <span class="information_value">{{precommitedBlocksValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Returns:</span>
            <span class="information_value">{{returnsValue}}</span>
          </div>

        </div>
        <div class="canvas_voting_power">
          <div class="progress_wrap">
            <span>Uptime(in last 100)</span>
            <div class="progress_wrap_background">
              <div class="progress_value" :style="`width:${firstPercent}`">{{firstPercent}}</div>
            </div>
          </div>
          <div class="progress_wrap">

          </div>
        </div>
      </div>

    </div>
    <div :class="transactionsDetailWrap" class="transaction_precommit_table">
      <div class="tab_wrap">
        <span @click="activeBtn = 0" :class="activeBtn === 0?'transactions_btn_active':''">Transactions</span>
        <span @click="activeBtn = 1" :class="activeBtn === 1?'transactions_btn_active':''">Precommit Blocks</span>
      </div>
      <div class="table_wrap">
        <div class="transactions_view_all" v-show="activeBtn === 0">
          <p style="flex:1;">Lastest 6 txns from a total Of {{transactionsCount}} transactions</p>
          <span @click="viewAllClick(2)">View All</span>
        </div>
        <div class="precommit_view_all" v-show="activeBtn === 1">
          <p class="table_instruction">
            <span>Total blocks:</span>
            <span>{{totalBlocks}}</span>
            <span>Total Fee:</span>
            <span>{{totalFee}}</span>
          </p>
          <span class="view_all_btn" @click="viewAllClick(1)">View All</span>
        </div>

        <div class="transaction_table">
          <blocks-list-table :items="items" :type="'6'" v-show="activeBtn === 0" :showNoData="showNoData1"></blocks-list-table>
          <blocks-list-table :items="itemsPre" :type="'7'" v-show="activeBtn === 1" :showNoData="showNoData2"></blocks-list-table>
          <div v-show="(activeBtn === 0 && showNoData1) || (activeBtn === 1 && showNoData2)" class="no_data_show">
            No Data
          </div>
        </div>
      </div>
    </div>



  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';
  import BlocksListTable from './table/BlocksListTable';

  export default {
    watch:{
      $route(){
        this.type = this.$route.params.type;
      },
      activeBtn(activeBtn){
        //0是Transactions List 1是Precommit Blocks List
      },

    },

    data() {

      return {
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
        activeBtn:0,
        firstPercent:'',
        secondPercent:'%',
        hashValue:this.$route.params.param,
        balanceValue:'',
        depositsValue:'',
        transactionsValue:'',
        nameValue:'',
        pubKeyValue:'',
        websiteValue:'',
        descriptionValue:'',
        commissionRateValue:'',
        announcementValue:'',
        bondHeightValue:'',
        votingPowerValue:'',
        uptimeValue:'',
        precommitedBlocksValue:'',
        returnsValue:'',
        items:[],
        itemsPre:[],
        type:this.$route.params.type,
        totalBlocks:0,
        totalFee:0,
        showNoData1:false,//无数据的时候显示无数据状态
        showNoData2:false,
        transactionsCount:0,


      }
    },
    components:{
      BlocksListTable,
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      this.getAddressInformation(this.$route.params.param);
      this.getTransactionsList(1,10,this.$route.params.type);
      this.getProfileInformation();
      this.getPrecommitBlocksList();
      this.getCurrentTenureInformation();
    },
    methods: {
      getAddressInformation(address){
        let url = `/api/account/${this.$route.params.param}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          let Amount = '';
          if(data.Amount instanceof Array){
            Amount = data.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
          }else if(data.Amount && Object.keys(data.Amount).includes('amount') && Object.keys(data.Amount).includes('denom')){
            Amount = `${data.Amount.amount} ${data.Amount.denom.toUpperCase()}`;
          }else if(data.Amount === null){
            Amount = '';
          }
          this.balanceValue = Amount;

        })
      },
      //点击view all跳转页面
      viewAllClick(type){
        if(type === 1){
          this.$router.push(`/block/${type}/0`);
        }else if(type === 2){
          this.$router.push(`/recent_transactions/2/0`)
        }

      },
      getProfileInformation(){
        let url = `/api/stake/candidate/${this.$route.params.param}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          if(data){
            this.nameValue = data.Description.Moniker;
            this.pubKeyValue = data.PubKey;
            this.websiteValue = data.Description.Website;
            this.descriptionValue= data.Description.Details;
            this.commissionRateValue = '';
            this.announcementValue = '';
            this.votingPowerValue = data.VotingPower;
          }

        })
      },
      getCurrentTenureInformation(){
        let url = `/api/stake/candidate/${this.$route.params.param}/status`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          console.log(data)
          if(data){
            this.bondHeightValue = data.TotalBlock;
            this.precommitedBlocksValue = data.PrecommitCount;
            this.returnsValue = '';
            this.firstPercent = (data.TotalBlock !== 0 && data.PrecommitCount !== 0) ?`${data.Uptime/data.TotalBlock*100}%`:'0%';
          }

        })
      },
      getTransactionsList(){
        let url = `/api/txsByAddress/${this.$route.params.param}/1/6`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          this.transactionsCount = data.Count;
          if(data.Data){
            this.items = data.Data.map(item=>{
              let [Amount,Fees] = ['',''];
              if(item.Amount instanceof Array){
                Amount = item.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
              }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
                Amount = `${item.Amount.amount} ${item.Amount.denom.toUpperCase()}`;
              }else if(item.Amount === null){
                Amount = '';
              }
              if(item.Fee.Amount instanceof Array){
                Fees = item.Fee.Amount.map(listItem=>`${listItem.amount} ${listItem.denom?listItem.denom.toUpperCase():'IRIS'}`).join(',');
              }else if(item.Fee.Amount && Object.keys(item.Fee.Amount).includes('amount') && Object.keys(item.Fee.Amount).includes('denom')){
                Fees = `${item.Fee.Amount} ${item.Fee.Amount}`;
              }else if(item.Fee.Amount === null){
                Fees = '';
              }
              return {
                TxHash:item.TxHash,
                Block:item.Height,
                From:item.From,
                To:item.To,
                Type:item.Type,
                Amount,
                Fees,
                Timestamp:Tools.conversionTimeToUTC(item.Time),
              }
            })
          }else{
            this.items = [{
              TxHash:'',
              Block:'',
              From:'',
              To:'',
              Type:'',
              Amount:'',
              Fees:'',
              Timestamp:'',
            }];
            this.showNoData1 = true;
          }


        })
      },
      getPrecommitBlocksList(){
        let url = `/api/blocks/precommits/${this.$route.params.param}/1/6`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          this.totalBlocks = data.Count;
          if(data.Data){
            this.itemsPre = data.Data.map(item=>{
              return{
                'Block Height':item.Height,
                Txn:item.NumTxs,
                Fees:'0 IRIS',
                Timestamp:Tools.conversionTimeToUTC(item.Time),
                'Precommit Validators':item.Validators.length !== 0?`${item.Block.LastCommit.Precommits.length}/${item.Validators.length}`:'',
              }
            })
          }else{
            this.itemsPre = [{
              'Block Height':'',
              Txn:'',
              Fees:'',
              Timestamp:'',
              'Precommit Validators':'',
            }];
            this.showNoData2 = true;
          }
        })
      }
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .transactions_detail_wrap {
    @include flex;
    @include pcContainer;
    font-size:0.14rem;
    .transactions_title_wrap {
      width: 100%;
      border-bottom: 0.01rem solid #eee;
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
      .transaction_information_content_title{
        height:0.4rem;
        line-height:0.4rem;
        font-size:0.18rem;
        color:#555;
        margin-bottom:0;
        border-bottom:0.01rem solid #efefef;
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
        .information_props_wrap{
          @include flex;
          .information_props{
            width:1.5rem;
            font-size:0.14rem;
          }
          .information_value{
            color: #a2a2ae;
            font-size:0.14rem;
          }
        }
      }






      .transactions_detail_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #555;
        margin-right: 0.2rem;
        font-weight:500;
      }
      .transactions_detail_wrap_hash_var {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.14rem;
        color: #ccc;
      }
    }

    .mobile_transactions_detail_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.05rem;
      .transaction_information_content_title{
        height:0.4rem;
        line-height:0.4rem;
        font-size:0.18rem;
        color:#555;
        margin-bottom:0;
      }
      .transactions_detail_information_wrap{

        .information_props_wrap{
          @include flex;
          flex-direction:column;
          border-bottom:none !important;
          margin-bottom:0.05rem;
          .information_value{
            overflow-x:auto;
            color: #a2a2ae;
          }
          .information_props{
            font-size:0.14rem;
          }


        }
      }
      .transactions_detail_title {
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #555;
        margin-right: 0.2rem;
        font-weight:500;
      }
      .transactions_detail_wrap_hash_var {
        overflow-x: auto;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.14rem;
        color: #ccc;
      }

    }
    .address_profile{
      @include borderRadius(5px);
    }

    //current tenure部分
    .current_tenure{
      margin-top:0.1rem;
      .current_tenure_wrap{
        width:100%;
      }
    }
    .personal_computer_transactions_detail_wrap{
      width:80%;
      .current_tenure_wrap{
        @include flex;
        flex-direction:row;
        .transactions_detail_information_wrap{
          flex:3;
        }
        .canvas_voting_power{
          flex:2;
          .progress_wrap{
            margin-bottom:0.15rem;
            .progress_wrap_background{
              height:0.24rem;
              background: #efeff1;
              margin-top:0.1rem;
              .progress_value{
                background:#a4d7f4;
                height:100%;
                line-height:0.24rem;
                text-indent:0.2rem;
              }
            }
          }

        }
      }
    }
    .mobile_transactions_detail_wrap{
      width:100%;
      .current_tenure_wrap{
        @include flex;
        flex-direction:column;
      }

      .canvas_voting_power{
        flex:2;
        .progress_wrap{
          margin-bottom:0.15rem;
          .progress_wrap_background{
            height:0.24rem;
            background: #efeff1;
            margin-top:0.1rem;
            .progress_value{
              background:#a4d7f4;
              height:100%;
              line-height:0.24rem;
              padding-left:0.2rem;
            }
          }
        }

      }
    }
    //底部表格部分
    .transaction_precommit_table{
      margin-top:0.1rem;
      border-bottom:0.01rem solid #eee;
      margin-bottom:0.2rem;
      .tab_wrap{

        span{
          height:0.3rem;
          line-height:0.3rem;
          display:inline-block;
          border:0.01rem solid #eee;
          text-align: center;
          padding:0 0.1rem;
          border-bottom:none;
          cursor:pointer;
        }
        .transactions_btn_active{
          background: #3190e8;
          color:#fff;
        }
      }
      .table_wrap{
        overflow-x:auto;
        .pagination{
          @include flex;
          justify-content: flex-end;
        }
        .transactions_view_all{
          padding:0.1rem;
          @include flex;
          justify-content:space-between;
          span{
            @include viewBtn;
          }
        }
        .precommit_view_all{
          padding:0.1rem;
          @include flex;
          justify-content:space-between;
          span{
            font-size:0.14rem;
            &:nth-child(3){
              display:inline-block;
              margin-left:0.1rem;
            }
          }
          .view_all_btn{
            @include viewBtn;
          }
        }
        .transaction_table{
          .no_data_show{
            @include flex;
            justify-content: center;
            border-top:0.01rem solid #eee;
            border-bottom:0.01rem solid #eee;
            font-size:0.14rem;
            height:3rem;
            align-items: center;
          }
        }
      }
    }
  }

</style>
