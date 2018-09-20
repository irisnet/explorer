<template>
  <div class="proposals_detail_wrap">
    <div class="proposals_title_wrap">
      <p :class="proposalsDetailWrap" style="margin-bottom:0;">
        <span class="proposals_detail_title">Proposals</span>
        <span class="proposals_detail_wrap_hash_var">{{`#${proposalsId}`}}</span>
      </p>
    </div>

    <div :class="proposalsDetailWrap">
      <p class="proposals_information_content_title">Proposals Information</p>
      <div class="proposals_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Proposal ID:</span>
          <span class="information_value">{{proposalsId}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Title:</span>
          <span class="information_value">{{title}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Type:</span>
          <span class="information_value">{{type}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Status:</span>
          <span class="information_value">{{status}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Submit Block:</span>
          <span class="information_value">{{submitBlock}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Submit Time:</span>
          <span class="information_value">{{submitTime}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Total Deposit:</span>
          <span class="information_value">{{totalDeposit}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Voting Start Block:</span>
          <span class="information_value">{{votingStartBlock}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Description:</span>
          <span class="information_value">{{description}}</span>
        </div>
      </div>
    </div>
    <div :class="proposalsDetailWrap">
      <p class="proposals_information_content_title" style='border-bottom:none !important;'>Vote Detals</p>
      <div class="vote-detals-content">
        <div class="total_num">
          <span>Total:{{count}}</span>
        </div>
        <div class="voting_options">
          <span>Yes:{{voteDetalsYes}}</span>|<span>No: {{voteDetalsNo}}</span>|<span>NoWithVeto:{{voteDetalsNoWithVeto}}</span>|<span>Abstain:{{voteDetalsAbstain}}</span>
        </div>
      </div>
    </div>
    <div :class="proposalsDetailWrap">
      <div class="proposals_detail_table_wrap">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="'ProposalsDetail'" :showNoData="showNoData"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
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
      SpinComponent
    },
    watch: {
      $route() {
        this.getProposalsInformation();
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        proposalsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
        items: [],
        showLoading:false,
        showNoData: false,//列表无数据的时候显示
        count: "",
        proposalsId: "",
        title: "",
        type: "",
        status: "",
        submitBlock: "",
        submitTime: "",
        totalDeposit: "",
        votingStartBlock: "",
        description: "",
        voteDetalsYes: "",
        voteDetalsNo: "",
        voteDetalsNoWithVeto: "",
        voteDetalsAbstain: "",
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.proposalsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.proposalsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      this.getProposalsInformation();
    },
    methods: {
      getProposalsInformation() {
        this.showLoading = true;
        let url = `/api/proposal/${this.$route.params.proposal_id}`;
        axios.get(url).then((data) => {
          if(data.status === 200) {
            return data.data
          }
        }).then((data) => {
          if(data && typeof  data === "object" ){
            // this.count = data.votes || data.proposal.status === "DepositPeriod" ? data.votes.length : "--";
            this.proposalsId = data.proposal.proposal_id;
            this.title = data.proposal.title;
            this.type = data.proposal.type;
            this.status = data.proposal.status;
            this.submitBlock = data.proposal.submit_block;
            this.submitTime = Tools.conversionTimeToUTCToYYMMDD(data.proposal.submit_time);
            this.votingStartBlock = data.proposal.voting_start_block ? data.proposal.voting_start_block : " -- ";
            this.description = data.proposal.description ? data.proposal.description : " -- ";
            this.voteDetalsYes = data.proposal.status === "DepositPeriod" ? "--" : data.result.Yes;
            this.voteDetalsNo = data.proposal.status === "DepositPeriod" ? "--" : data.result.No;
            this.voteDetalsNoWithVeto = data.proposal.status === "DepositPeriod" ? "--" : data.result.NoWithVeto;
            this.voteDetalsAbstain = data.proposal.status === "DepositPeriod" ? "--" : data.result.Abstain;

            if(data.proposal && data.proposal.total_deposit.length !==0){
              this.totalDeposit = Tools.scientificToNumber(Tools.formatNumber(data.proposal.total_deposit[0].amount)) + " " +data.proposal.total_deposit[0].denom.toUpperCase();
            }else {
              this.totalDeposit = "";
            }
            if(data.proposal.status === "DepositPeriod"){
              this.count = "--"
            }
            if(data.votes){
              this.count = data.votes.length;
              this.items = data.votes.map(item =>{
                item.time = Tools.conversionTimeToUTCToYYMMDD(item.time);
                return {
                  Voter: item.voter,
                  "Voter Option": item.option,
                  "Vote Time": item.time
                }
              })
            }else {
              this.items = [{
                Voter: "",
                "Vote Option": "",
                "Vote Time": ""
              }]
            }
          }
          this.showLoading = false;
        })

      },

    }
  }
</script>

<style scoped lang="scss">
  @import '../style/mixin.scss';

  .proposals_detail_wrap {
    @include flex;
    @include pcContainer;
      font-size: 0.14rem;
    .proposals_title_wrap {
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
    .proposals_information_content_title {
      margin-left: 0.2rem !important;
      height: 0.5rem !important;
      line-height: 0.5rem !important;
      font-size: 0.18rem !important;
      color: #000000;
      margin-bottom: 0;
      font-weight: 500!important;
      border-bottom:1px solid #d6d9e0 !important;
    }
    @include pcCenter;
      .proposals_detail_information_wrap {
        margin-top: 0.21rem;
        margin-left: 0.2rem;
      .information_props_wrap {
        @include flex;
        margin-bottom:0.08rem;
    .information_props {
      width: 1.5rem;
      font-weight: 500;
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
  .proposals_detail_table_wrap {
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

    .proposals_detail_title {
      height: 0.61rem;
      line-height: 0.61rem;
      font-size: 0.22rem;
      color: #000000;
      margin-right: 0.2rem;
      font-weight: 500;
      margin-left: 0.2rem;
    }
    .proposals_detail_wrap_hash_var {
      height: 0.61rem;
      line-height: 0.61rem;
      font-size: 0.22rem;
      color: #a2a2ae;
    }
  }

  .mobile_transactions_detail_wrap {
    width: 100%;
    @include flex;
    flex-direction: column;
    padding: 0 0.1rem;
    .proposals_information_content_title {
      height: 0.5rem !important;
      line-height: 0.5rem !important;
      font-size: 0.18rem !important;
      color: #000000;
      margin-bottom: 0;
      font-weight: 500!important;
    }
    .proposals_detail_table_wrap {
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
    .proposals_detail_information_wrap {

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
      .proposals_detail_title {
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
  .vote-detals-content{
    border-top: 1px solid #d6d9e0;
    display: flex;
    justify-content: space-between;
    height: 0.62rem;
    line-height: 0.62rem;
  }
  .total_num{
    font-size: 0.14rem;
    color: #a2a2ae;
    margin-left: 0.2rem;
  }
  .voting_options{
    display: flex;
    color: #a2a2ae;
    span{
      font-size: 0.14rem;
      color: #000;
      font-weight: 500;
      padding: 0 0.18rem;
    }
  }
</style>
