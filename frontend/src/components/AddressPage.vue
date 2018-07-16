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
        <div class="information_props_wrap" style="border-bottom:0.1rem solid #eee">
          <span class="information_props">Description:</span>
          <span class="information_value">{{descriptionValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Commission Rate:</span>
          <span class="information_value">{{commissionRateValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Announcement:</span>
          <span class="information_value">{{AnnouncementValue}}</span>
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
            <span class="information_props">Uptime(in last 100 blocks):</span>
            <span class="information_value">{{uptimeValue}}</span>
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
            <span>Commission Rate</span>
            <div class="progress_wrap_background">
              <div class="progress_value" :style="`width:${secondPercent}`">{{secondPercent}}</div>
            </div>
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
        <p class="table_instruction">
          <span>Total blocks:</span>
          <span>{{}}</span>
          <span>Total Fee:</span>
          <span>{{}}</span>
        </p>
        <div class="transaction_table">
          <blocks-list-table :items="items" :type="type"></blocks-list-table>
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
      }
    },

    data() {

      return {
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
        activeBtn:0,
        firstPercent:'33%',
        secondPercent:'50%',
        hashValue:'',
        balanceValue:'',
        depositsValue:'',
        transactionsValue:'',
        nameValue:'',
        pubKeyValue:'',
        websiteValue:'',
        descriptionValue:'',
        commissionRateValue:'',
        AnnouncementValue:'',
        bondHeightValue:'',
        votingPowerValue:'',
        uptimeValue:'',
        precommitedBlocksValue:'',
        returnsValue:'',
        items:[],
        type:this.$route.params.type,


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
      /*this.getAddressInformation();
      this.getProfileInformation();
      this.getCurrentTenureInformation();
      this.getTableListInforamtion(1,10,this.$route.params.type);*/

    },
    methods: {
      getAddressInformation(){
        let url = ``;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          console.log(data)

        })
      },
      getProfileInformation(){
        let url = ``;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          console.log(data)

        })
      },
      getCurrentTenureInformation(){
        let url = ``;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          console.log(data)

        })
      },
      getTableListInforamtion(currentPage,pageSize,type){
        let url = ``;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          console.log(data)

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
    font-size:1.4rem;
    .transactions_title_wrap {
      width: 100%;
      border-bottom: 0.1rem solid #eee;
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
        height:4rem;
        line-height:4rem;
        font-size:1.8rem;
        color:#555;
        margin-bottom:0;
        border-bottom:0.1rem solid #efefef;
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
        .information_props_wrap{
          @include flex;
          .information_props{
            width:15rem;
          }
        }
      }






      .transactions_detail_title {
        height: 4rem;
        line-height: 4rem;
        font-size: 1.8rem;
        color: #555;
        margin-right: 2rem;
        font-weight:500;
      }
      .transactions_detail_wrap_hash_var {
        height: 4rem;
        line-height: 4rem;
        font-size: 1.4rem;
        color: #ccc;
      }
    }

    .mobile_transactions_detail_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.5rem;
      .transaction_information_content_title{
        height:4rem;
        line-height:4rem;
        font-size:1.8rem;
        color:#555;
        margin-bottom:0;
      }
      .transactions_detail_information_wrap{

        .information_props_wrap{
          @include flex;
          flex-direction:column;
          border-bottom:none !important;
          margin-bottom:0.5rem;
          .information_value{
            overflow-x:scroll;
          }

        }
      }
      .transactions_detail_title {
        height: 3rem;
        line-height: 3rem;
        font-size: 1.8rem;
        color: #555;
        margin-right: 2rem;
        font-weight:500;
      }
      .transactions_detail_wrap_hash_var {
        overflow-x: scroll;
        height: 3rem;
        line-height: 3rem;
        font-size: 1.4rem;
        color: #ccc;
      }

    }
    .address_profile{
      @include borderRadius(5px);
    }

    //current tenure部分
    .current_tenure{
      margin-top:1rem;
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
            margin-bottom:1.5rem;
            .progress_wrap_background{
              height:2.4rem;
              background: #efeff1;
              margin-top:1rem;
              .progress_value{
                background:#a4d7f4;
                height:100%;
                line-height:2.4rem;
                padding-left:2rem;
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
          margin-bottom:1.5rem;
          .progress_wrap_background{
            height:2.4rem;
            background: #efeff1;
            margin-top:1rem;
            .progress_value{
              background:#a4d7f4;
              height:100%;
              line-height:2.4rem;
              padding-left:2rem;
            }
          }
        }

      }
    }
    //底部表格部分
    .transaction_precommit_table{
      margin-top:1rem;
      border-bottom:0.1rem solid #eee;
      .tab_wrap{

        span{
          height:3rem;
          line-height:3rem;
          display:inline-block;
          border:0.1rem solid #eee;
          text-align: center;
          padding:0 1rem;
          border-bottom:none;
          cursor:pointer;
        }
        .transactions_btn_active{
          background: #3190e8;
          color:#fff;
        }
      }
      .table_wrap{
        border:0.1rem solid #eee;
      }
    }
  }

</style>
