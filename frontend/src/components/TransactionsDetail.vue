<template>
  <div class="transactions_detail_wrap">
    <div class="transactions_title_wrap">
      <p :class="transactionsDetailWrap" style="margin-bottom:0;">
        <span class="transactions_detail_title">Transaction</span>
        <span class="transactions_detail_wrap_hash_var">{{hashValue}}</span>
      </p>
    </div>
    <div :class="transactionsDetailWrap">
      <p class="transaction_information_content_title">Transaction Information</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">TxHash :</span>
          <span class="information_value">{{hashValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Block Height :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/block/${blockValue}`">{{blockValue}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Type :</span>
          <span class="information_value">{{typeValue}}</span>
        </div>
        <div class="information_props_wrap" v-if="flShowProposer">
          <span class="information_props">Proposer :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${proposer}`">{{proposer}}</router-link></span>
        </div>
        <div class="information_props_wrap" v-if="title">
          <span class="information_props">ProposalTitle :</span>
          <span class="information_value">{{title}}</span>
        </div>
        <div class="information_props_wrap" v-if="proposalType">
          <span class="information_props">ProposalType :</span>
          <span class="information_value">{{proposalType}}</span>
        </div>
        <div class="information_props_wrap" v-if="flShowInitialDeposit">
          <span class="information_props">InitialDeposit :</span>
          <span class="information_value">{{initialDeposit}}</span>
        </div>
        <div class="information_props_wrap" v-if="description">
          <span class="information_props">Description :</span>
          <span class="information_value">{{description}}</span>
        </div>
        <div class="information_props_wrap" v-if="depositer">
          <span class="information_props">Depositer :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${depositer}`">{{depositer}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap" v-if="flShowProposalId">
          <span class="information_props">Proposal ID :</span>
          <span v-show="proposalId !== '--' " class="information_value link_active_style">
            <router-link :to="`/ProposalsDetail/${proposalId}`">{{proposalId}}</router-link>
          </span>
          <span v-show="proposalId === '--' " class="information_value link_active_style">{{proposalId}}</span>
        </div>
        <div class="information_props_wrap" v-if="flShowVoter">
          <span class="information_props">Voter :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${voter}`">{{voter}}</router-link></span>
        </div>
        <div class="information_props_wrap" v-if="flShowTypeTransfer || flShowWithdrawAddress">
          <span class="information_props">From :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${fromValue}`">{{fromValue}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap" v-if="flShowWithdrawAddress">
          <span class="information_props">Withdraw To :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${withdrawAddress}`">{{withdrawAddress}}</router-link></span>
        </div>
        <div class="information_props_wrap" v-if="flShowDelegatorAddress">
          <span class="information_props">Delegator Address :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${delegatorAddress}`">{{delegatorAddress}}</router-link></span>
        </div>
        <div class="information_props_wrap" v-if="flShowValidatorAddress">
          <span class="information_props">Validator Address :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${validatorAddress}`">{{validatorAddress}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap" v-if="showSource">
          <span class="information_props">Source :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${source}`">{{source}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap" v-if="flShowTypeTransfer">
          <span class="information_props">To :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${toValue}`">{{toValue}}</router-link></span>
        </div>
        <div class="information_props_wrap" v-if="moniker">
          <span class="information_props">Moniker :</span>
          <span class="information_value"><pre class="information_pre">{{moniker}}</pre></span>
        </div>
        <div class="information_props_wrap" v-if="identity">
          <span class="information_props">Identity :</span>
          <span class="information_value"><pre class="information_pre">{{identity}}</pre></span>
        </div>
        <div class="information_props_wrap" v-if="owner">
          <span class="information_props">From :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/address/1/${owner}`">{{owner}}</router-link></span>
        </div>
        <div class="information_props_wrap" v-if="pubkey">
          <span class="information_props">Pub key :</span>
          <span class="information_value">{{pubkey}}</span>
        </div>
        <div class="information_props_wrap" v-if="website">
          <span class="information_props">Website :</span>
          <span class="information_value"><pre class="information_pre">{{website}}</pre></span>
        </div>
        <div class="information_props_wrap" v-if="selfBond">
          <span class="information_props">Self-Bonded :</span>
          <span class="information_value">{{selfBond}}</span>
        </div>
        <div class="information_props_wrap" v-if="details">
          <span class="information_props">Details :</span>
          <span class="information_value"><pre class="information_pre">{{details}}</pre></span>
        </div>
        <div class="information_props_wrap" v-if="flShowVoter">
          <span class="information_props">Option :</span>
          <span class="information_value">{{option}}</span>
        </div>
        <div class="information_props_wrap" v-if="flShowTypeTransfer || flShowTypeDeposit">
          <span class="information_props">Amount :</span>
          <span class="information_value">{{amountValue}}</span>
        </div>
        <div class="information_props_wrap" v-if="flShowReceivedRewardsValue">
          <span class="information_props">Received Rewards :</span>
          <span class="information_value">{{amountValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Status :</span>
          <span class="information_value">{{status}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Age(Timestamp) :</span>
          <span class="information_value" v-show="ageValue">{{ageValue}} ({{timestampValue}})</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Actual Tx Fee :</span>
          <span class="information_value">{{actualTxFee}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Gas Limit :</span>
          <span class="information_value">{{gasLimit}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Gas Used by Tx :</span>
          <span class="information_value">{{gasUsedByTxn}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Gas Price :</span>
          <span class="information_value">{{gasPrice}} <span v-if="gasPrice"></span>Nano</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Memo :</span>
          <span class="information_value"><pre class="information_pre">{{memo}}</pre></span>
        </div>

      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import Service from "../util/axios";
  import Constant from "../constant/Constant"
  export default {

    data() {
      return {
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',
        hashValue: '',
        blockValue: '',
        typeValue: '',
        fromValue: '',
        toValue: '',
        timestampValue: '',
        amountValue: '',
        actualTxFee: '',
        gasLimit:'',
        gasUsedByTxn:'',
        gasPrice:'',
        memo: "",
        owner: "",
        moniker: "",
        selfBond: "",
        pubkey: "",
        identity: "",
        website: "",
        details: "",
        source: "",
        showSource: "",
        proposer: "",
        initialDeposit: "",
        title: "",
        description: "",
        proposalId: "",
        proposalType:"",
        depositer: "",
        voter: "",
        option: "",
        status: "",
        withdrawAddress: "",
        delegatorAddress: "",
        validatorAddress: "",
        flShowProposalId: false,
        flShowProposer: false,
        flShowInitialDeposit: false,
        flShowTypeTransfer:false,
        flShowVoter: false,
        flShowTypeDeposit: false,
        flShowWithdrawAddress: false,
        flShowDelegatorAddress: false,
        flShowValidatorAddress: false,
        flShowReceivedRewardsValue: false,
        ageValue: '',
        transactionDetailTimer: null,
      }
    },
    watch:{
      $route(){
        this.getTransactionInfo();
        Tools.scrollToTop();
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
      this.getTransactionInfo()
    },
    methods: {
      getTransactionInfo(){
        if(this.$route.query.txHash){
          let url = `/api/tx/${this.$route.query.txHash}`;
          Service.http(url).then((data)=>{
            clearInterval(this.transactionDetailTimer);
            if(data){
              let that = this;
              let currentServerTime = new Date().getTime() + that.diffMilliseconds;
              this.transactionDetailTimer = setInterval(function () {
                currentServerTime = new Date().getTime() + that.diffMilliseconds;
                that.ageValue = Tools.formatAge(currentServerTime,data.Timestamp,Constant.SUFFIX);
              },1000);
              this.ageValue = Tools.formatAge(currentServerTime,data.Timestamp,Constant.SUFFIX);
              this.timestampValue = Tools.format2UTC(data.Timestamp);
              this.hashValue = data.Hash;
              this.blockValue = data.BlockHeight;
              this.typeValue = data.Type === 'coin'?'transfer':data.Type;
              this.gasPrice = Tools.convertScientificNotation2Number(Tools.formaNumberAboutGasPrice(data.GasPrice));
              this.gasLimit = data.GasLimit;
              this.gasUsedByTxn = data.GasUsed;
              this.memo = data.Memo ? data.Memo : '--';
              this.status = data.Status ? Tools.firstWordUpperCase(data.Status): '--';
              if(data.Amount && data.Amount.length !==0){
                this.amountValue = data.Amount.map(item=>{
                  item.amount =Tools.formatPriceToFixed(Tools.convertScientificNotation2Number(Tools.formatNumber(item.amount)));
                  if(!item.denom){
                    return `${item.amount} SHARES`;
                  }else{
                    return `${item.amount} ${Tools.formatDenom(item.denom).toUpperCase()}`;
                  }
                }).join(',') ;
              }else if(data.Amount && Object.keys(data.Amount).includes('amount') && Object.keys(data.Amount).includes('denom')){
                data.Amount =  Tools.formatPriceToFixed(Tools.convertScientificNotation2Number(Tools.formatNumber(data.Amount.amount)));
                this.amountValue = `${data.Amount.amount} ${Tools.formatDenom(data.Amount.denom).toUpperCase()}`
              }else {
                this.amountValue = "--"
              }
              this.actualTxFee = `${Tools.convertScientificNotation2Number(Tools.formatNumber(data.Fee.amount))} ${Tools.formatDenom(data.Fee.denom).toUpperCase()}`;
              if(data.Type === "Transfer" || data.Type === "Delegate" || data.Type === "BeginUnbonding"){
                this.flShowTypeTransfer = true;
                this.fromValue = data.From;
                this.toValue = data.To;
              }else if(data.Type === "CreateValidator" || data.Type === "EditValidator" || data.Type === "Unjail"){
                this.owner = data.Owner ? data.Owner : '--';
                this.moniker = data.Moniker ? data.Moniker : '--';
                this.pubkey = data.Pubkey ? data.Pubkey : "--";
                this.identity = data.Identity ? data.Identity : '--';
                this.website = data.Website ? data.Website : '--';
                this.details = data.Details ? data.Details : '--';
                if(data.SelfBond && data.SelfBond.length !== 0){
                  this.selfBond = `${Tools.formatPriceToFixed(Tools.convertScientificNotation2Number(Tools.formatNumber(data.SelfBond[0].amount)))} ${Tools.formatDenom(data.SelfBond[0].denom).toUpperCase()}`;
                }else {
                  this.selfBond = "--"
                }
              }else if(data.Type === "BeginRedelegate"){
                this.flShowTypeTransfer = true;
                this.showSource = true;
                this.fromValue = data.From ? data.From : '';
                this.toValue = data.To ? data.To : "";
                this.source = data.Source ? data.Source : "";
              }else if(data.Type === "SubmitProposal"){
                this.flShowProposer = true;
                this.flShowInitialDeposit = true;
                this.title = data.Title ? data.Title : '--';
                this.proposer = data.From;
                this.proposalType = data.ProposalType;
                if(data.Amount && data.Amount.length !==0){
                  this.initialDeposit = data.Amount.map(item=>{
                    return `${item.amount} ${Tools.formatDenom(item.denom).toUpperCase()}`;
                  }).join(',') ;
                }else {
                  this.initialDeposit = "--"
                }
                this.description = data.Description ? data.Description : '--';
              }else if(data.Type === "Deposit"){
                this.flShowProposalId = true;
                this.flShowTypeDeposit = true;
                this.proposalId = data.ProposalId === 0 ? "--" : data.ProposalId;
                this.depositer = data.From ? data.From : "--";
              }else if(data.Type === "Vote"){
                this.flShowProposalId = true;
                this.flShowVoter = true;
                this.proposalId = data.ProposalId === 0 ? "--" : data.ProposalId;
                this.voter = data.From ? data.From : '--';
                this.option = data.Option ? data.Option : "--";
              }else if(data.Type === "SetWithdrawAddress"){
                this.flShowWithdrawAddress = true;
                this.fromValue = data.From ? data.From : '';
                this.withdrawAddress = data.To ? data.To : '';
              }else if(data.Type === "WithdrawDelegatorRewardsAll"){
                this.flShowReceivedRewardsValue = true;
                this.flShowDelegatorAddress = true;
                this.delegatorAddress = data.From ? data.From : '';
              }else if(data.Type === "WithdrawDelegatorReward"){
                this.flShowReceivedRewardsValue = true;
                this.flShowDelegatorAddress = true;
                this.flShowValidatorAddress = true;
                this.delegatorAddress = data.From ? data.From : '';
                this.validatorAddress = data.To ? data.To : "";
              } else if(data.Type === "WithdrawValidatorRewardsAll"){
                this.flShowReceivedRewardsValue = true;
                this.flShowValidatorAddress = true;
                this.validatorAddress = data.From ? data.From : "";
              }
            }

          }).catch(e => {
            console.log(e)
          })
        }
      }
    }
  }
</script>
<style lang="scss" scoped>
  @import '../style/mixin.scss';
  .information_pre{
    color: #a2a2ae;
    white-space: pre-wrap;
  }
  .transactions_detail_wrap {
    @include flex;
    @include pcContainer;
    font-size:0.14rem;
    .transactions_title_wrap {
      width: 100%;
      border-bottom: 1px solid #d6d9e0 !important;
      height:0.62rem;
      background:#efeff1;
      line-height:0.62rem;
      @include flex;
      @include pcContainer;
      .personal_computer_transactions_detail_wrap {
        height:0.62rem;
        @include flex;
        span{
      line-height:0.62rem;
      height:0.62rem;
        }
      }
      .mobile_transactions_detail_wrap {
        @include flex;
        flex-direction: column;
      }

    }
    .personal_computer_transactions_detail_wrap {
      .transaction_information_content_title{
        height:0.5rem;
        line-height:0.5rem;
        font-size:0.18rem;
        color:#000000;
        @include fontWeight;
        margin-bottom:0;
        font-family:ArialMT;
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
        padding:0.16rem 0rem;
          .information_props_wrap{
            @include flex;
              margin-bottom: 0.08rem;
            .information_props{
              width:1.5rem;
            }
              .information_value{
                  color: #a2a2ae;
              }
          }
      }
      .transactions_detail_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
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
      .transaction_information_content_title{
        height: 0.5rem;
        line-height: 0.5rem;
        font-size: 0.18rem;
        color: #000000;
        margin-bottom:0;
        @include fontWeight;
      }
      .transactions_detail_information_wrap{
        padding:0.16rem 0rem;
        .information_props_wrap{
          @include flex;
          flex-direction:column;
          margin-bottom:0.05rem;
          .information_value{
            overflow-x:auto;
            -webkit-overflow-scrolling:touch;
            overflow-y:hidden;
          }
            .information_value{
                color: #a2a2ae;
            }
        }
      }
      .transactions_detail_title {
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
        @include fontWeight;
      }
      .transactions_detail_wrap_hash_var {
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #a2a2ae;
      }
    }
  }
  .link_active_style{
    a{
      color:#3598db !important;
    }
    color:#3598db !important;
    cursor:pointer;
  }
</style>
