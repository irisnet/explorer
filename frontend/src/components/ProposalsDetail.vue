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
          <span class="information_props">Title :</span>
          <span class="information_value information_show_trim">
            <pre class="information_pre">{{title}}</pre>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Proposer :</span>
          <span v-show="proposer !== '--'" class="information_value information_show_trim jump_route" @click="jumpRoute(`/address/1/${proposer}`)">{{proposer}}</span>
          <span v-show="proposer == '--'" class="information_value information_show_trim ">{{proposer}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Submit Hash :</span>
          <span v-show="submitHash !== '--'" class="information_value information_show_trim jump_route" @click="jumpRoute(`/tx?txHash=${submitHash}`)">{{submitHash}}</span>
          <span v-show="submitHash == '--'" class="information_value information_show_trim ">{{submitHash}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Type :</span>
          <span class="information_value">{{type}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Status :</span>
          <span class="information_value">{{status}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Submit Time :</span>
          <span class="information_value">{{submitAge}} <span v-show="submitAge">(</span>{{submitTime}}<span v-show="submitAge">)</span></span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Deposit End Time :</span>
          <span class="information_value">{{depositEndAge}} <span v-show="depositEndAge">(</span>{{depositEndTime}}<span v-show="depositEndAge">)</span>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Total Deposit :</span>
          <span class="information_value">{{totalDeposit}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Voting Start Time :</span>
          <span class="information_value">{{votingStartAge}} <span v-show="votingStartAge">(</span>{{votingStartTime}}<span v-show="votingStartAge">)</span>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Voting End Time :</span>
          <span class="information_value">{{votingEndAge}} <span v-show="votingEndAge">(</span>{{votingEndTime}}<span v-show="votingEndAge">)</span></span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Description :</span>
          <span class="information_value information_show_trim">
            <pre class="information_pre">{{description}}</pre>
          </span>
        </div>

        <div v-show="type === 'SoftwareUpgrade'">
          <div class="information_props_wrap">
            <span class="information_props">Software :</span>
            <span class="information_value information_show_trim">
            <a class="information_link" :href="software" target="_blank">{{software}}</a>
          </span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Version :</span>
            <span class="information_value information_show_trim">
            <span class="information_pre">{{version}}</span>
          </span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Switch Height :</span>
            <span class="information_value information_show_trim">
            <span class="information_pre">{{switchHeight}}</span>
          </span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Threshold :</span>
            <span class="information_value information_show_trim">
            <span class="information_pre">{{threshold}}</span>
          </span>
          </div>
        </div>
        <div class="parameter_container" v-show="type === 'ParameterChange'">
          <div class="information_props_wrap">
            <span class="information_props">Parameter Details :</span>
            <textarea :rows="textareaRows" readonly spellcheck="false" class="parameter_detail_content">{{parameterValue}}
            </textarea>
          </div>
        </div>
      </div>
    </div>
    <div :class="proposalsDetailWrap">
      <p class="proposals_information_content_title" style='border-bottom:none !important;'>Vote Details</p>
      <div class="vote-details-content">
        <div class="vote_content_container">
          <div class="total_num">
            <span>{{count}} Total</span>
          </div>
          <div class="voting_options">
            <span>Yes : {{voteDetailsYes}}</span>|<span>No : {{voteDetailsNo}}</span>|<span>NoWithVeto : {{voteDetailsNoWithVeto}}</span>|<span>Abstain : {{voteDetailsAbstain}}</span>
          </div>
        </div>
      </div>
    </div>
    <div :class="proposalsDetailWrap">
      <div class="proposals_detail_table_wrap">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="'ProposalsDetail'" :showNoData="showNoData" :min-width="tableMinWidth"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
    </div>


  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import Service from "../util/axios"
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  import Constant from "../constant/Constant"
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
        proposalsDetailWrap: 'personal_computer_transactions_detail',
        items: [],
        showLoading:false,
        showNoData: false,
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
        voteDetailsYes: "",
        voteDetailsNo: "",
        voteDetailsNoWithVeto: "",
        voteDetailsAbstain: "",
        proposer: "",
        submitHash: "",
        tableMinWidth: "",
        depositEndTime: "",
        votingStartTime: "",
        votingEndTime: "",
        submitAge: '',
        depositEndAge: '',
        votingStartAge:'',
        votingEndAge: '',
        votingTimer: null,
        proposalTimer: null,
        software: ' ',
        version: ' ',
        switchHeight: ' ',
        threshold: ' ',
        textareaRows: '2',
        parameterValue: ' '
      }
    },
    beforeMount() {
      Tools.scrollToTop();
      if (Tools.currentDeviceIsPersonComputer()) {
        this.proposalsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.proposalsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      this.getProposalsInformation();
      this.computeMinWidth();
    },
    methods: {
      computeMinWidth(){
        if(this.$route.params.proposal_id){
          this.tableMinWidth = 7.5;
        }
      },
      flShowProposalTime(proposalTimeName,status){
       if(status === 'Rejected' || status === 'Passed' || status === 'VotingPeriod'){
         return true
       }else{
         switch (proposalTimeName){
           case proposalTimeName === 'depositEndTime' && status === 'DepositPeriod' : return true ;
           case proposalTimeName === 'votingStartTime' && status === 'VotingPeriod' : return true ;
           case proposalTimeName === 'votingEndTime' && status === 'VotingPeriod' : return true ;
         }
       }
      },
      formatProposalTime(time){
        let currentServerTime  = new Date().getTime() + this.diffMilliseconds;
        if(time && new Date(time).getTime() < currentServerTime){
          return Tools.formatAge(currentServerTime,time,Constant.SUFFIX);
        }
      },
      getProposalsInformation() {
        this.showLoading = true;
        let url = `/api/proposal/${this.$route.params.proposal_id}`;
        Service.http(url).then((data) => {
          this.showLoading = false;
          if(data){
            this.showNoData = false;
            if(data.proposal.proposal_id === 0){
              this.proposalsId = '--';
              this.title = '--';
              this.type = '--';
              this.status = '--';
              this.proposer = '--';
              this.submitHash = '--';
              this.submitTime = '--';
              this.depositEndTime = '--';
              this.votingStartTime = '--';
              this.votingEndTime = "--";
              this.description = '--';
              this.voteDetailsYes = '--';
              this.voteDetailsNo = '--';
              this.voteDetailsNoWithVeto = '--';
              this.voteDetailsAbstain = '--';
              this.voteDetailsAbstain = '--';
              this.totalDeposit = '--';
              this.count = 0;
            }else {
              let that = this;
              clearInterval(this.proposalTimer);
              this.proposalTimer = setInterval(function () {
                that.submitAge = that.formatProposalTime(data.proposal.submit_time ? data.proposal.submit_time : '');
                that.depositEndAge = that.formatProposalTime(that.flShowProposalTime('depositEndTime',data.proposal.status) ? data.proposal.deposit_end_time : '');
                that.votingStartAge = that.formatProposalTime(that.flShowProposalTime('votingStartTime',data.proposal.status) ? data.proposal.voting_start_time : '');
                that.votingEndAge = that.formatProposalTime(that.flShowProposalTime('votingEndTime',data.proposal.status) ? data.proposal.voting_end_time : '' );
              },1000);
              this.software = data.proposal.software;
              this.version = data.proposal.version;
              this.switchHeight = data.proposal.switch_height;
              this.threshold = data.proposal.threshold ? `${Number(data.proposal.threshold) * 100}%` : '';
              this.proposalsId = data.proposal.proposal_id === 0 ? "--" : data.proposal.proposal_id;
              this.title = data.proposal.title;
              this.type = data.proposal.type;
              this.status = data.proposal.status;
              this.proposer = data.proposal.proposer ? data.proposal.proposer : "--";
              this.submitHash = data.proposal.tx_hash ? data.proposal.tx_hash : "--";
              this.submitTime = data.proposal.submit_time ? Tools.format2UTC(data.proposal.submit_time) : '--';
              this.depositEndTime = that.flShowProposalTime('depositEndTime',data.proposal.status) ? Tools.format2UTC(data.proposal.deposit_end_time) : '--';
              this.votingStartTime = that.flShowProposalTime('votingStartTime',data.proposal.status) ? Tools.format2UTC(data.proposal.voting_start_time) : '--';
              this.votingEndTime = that.flShowProposalTime('votingEndTime',data.proposal.status) ? Tools.format2UTC(data.proposal.voting_end_time) : '--';
              this.description = data.proposal.description ? data.proposal.description : " -- ";
              this.voteDetailsYes = data.proposal.status === "DepositPeriod" ? "--" : data.result.Yes;
              this.voteDetailsNo = data.proposal.status === "DepositPeriod" ? "--" : data.result.No;
              this.voteDetailsNoWithVeto = data.proposal.status === "DepositPeriod" ? "--" : data.result.NoWithVeto;
              this.voteDetailsAbstain = data.proposal.status === "DepositPeriod" ? "--" : data.result.Abstain;
              if(data.proposal && data.proposal.total_deposit.length !==0){
                this.totalDeposit = `${Tools.convertScientificNotation2Number(Tools.formatNumber(data.proposal.total_deposit[0].amount))} ${Tools.formatDenom(data.proposal.total_deposit[0].denom).toUpperCase()}`;
              }else {
                this.totalDeposit = "";
              }
              if(data.proposal.status === "DepositPeriod"){
                this.count = "--"
              }else {
                this.count = 0;
              }
              if(data.proposal.type === 'ParameterChange'){
                for (let index = 0; index < data.proposal.parameters.length; index++){
                  this.parameterValue+= `${data.proposal.parameters[index].subspace}/${data.proposal.parameters[index].key} = ${data.proposal.parameters[index].value}\n`
                }
                let defaultTextareaRows = 2;
                this.textareaRows = data.proposal.parameters.length + defaultTextareaRows;
              }
              if(data.votes){
                this.count = data.votes.length;
                let that = this;
                clearInterval(this.votingTimer);
                this.votingTimer = setInterval(function () {
                  let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                  that.items = data.votes.map(item =>{
                    let votingListItemTime = Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX);
                    return {
                      Voter: item.voter,
                      "Vote Option": item.option,
                      VoteTime: votingListItemTime
                    }
                  })
                },1000);
              }else {
                this.items = [{
                  Voter: "",
                  "Vote Option": "",
                  "Vote Time": ""
                }];
              }
            }
          }else {
              this.showNoData = false;
              this.items = [{
                Voter: "",
                "Vote Option": "",
                "Vote Time": ""
              }];
              this.showNoData = true
            }
        }).catch(e => {
          this.showNoData = false;
          console.log(e)
        })

      },
      jumpRoute(path) {
        this.$router.push(path);
      }

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
    .proposals_detail_information_wrap{
      .parameter_container{
        .information_props_wrap{
          .parameter_detail_content{
            width: 90%;
            margin-right:20%;
            background: #EEE;
          }
        }
      }
    }
    }
  }
  .personal_computer_transactions_detail_wrap {
    width: 100%!important;
    .proposals_information_content_title {
      padding-left: 0.2rem !important;
      height: 0.5rem !important;
      line-height: 0.5rem !important;
      font-size: 0.18rem !important;
      color: #000000;
      margin-bottom: 0;
      @include fontWeight;
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
      min-width: 1.5rem;
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
      .parameter_detail_content{
        width: 100%;
        margin-right:20%;
        background: #eee;
      }
    }
  }
  .proposals_detail_table_wrap {
    margin-bottom: 0.2rem;
    width: 100%;
    overflow-x: auto;
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
      @include fontWeight;
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
    padding-left: 0.1rem;
    .proposals_detail_wrap_hash_var{
      color: #a2a2ae;
    }
    .proposals_information_content_title {
      height: 0.5rem !important;
      line-height: 0.5rem !important;
      font-size: 0.18rem !important;
      color: #000000;
      margin-bottom: 0;
      @include fontWeight;
      border-bottom: 1px solid #d6d9e0 !important;
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
      .vote-details-content{
        width: 100%;
        overflow-x: auto;
        border-top: 1px solid #d6d9e0;
        display: flex;
        justify-content: space-between;
        height: 0.62rem;
        line-height: 0.62rem;
        .vote_content_container{
          min-width: 150%;
          display: flex;
          justify-content: space-between;
        }
      }
    }
  }
  .information_show_trim{
    color: #a2a2ae;
  }
  .information_value{
    color: #a2a2ae;
    word-break: break-all;
  }
  .vote-details-content{
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
      @include fontWeight;
      padding: 0 0.18rem;
    }
  }
  .information_show_trim{
    white-space: pre-wrap ;
  }
  .jump_route {
    color: #3598db;
    cursor: pointer;
  }
  .vote_content_container{
    min-width: 100%;
    display: flex;
    justify-content: space-between;
  }
  pre{
    font-family: Arial !important;
  }
  .information_link{
    color: #3498db !important;
  }
</style>
