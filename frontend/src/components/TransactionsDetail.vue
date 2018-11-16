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
          <span class="information_value link_active_style" @click="skipRoute(`/blocks_detail/${blockValue}`)">{{blockValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Type :</span>
          <span class="information_value">{{typeValue}}</span>
        </div>
        <div class="information_props_wrap" v-if="showProposer">
          <span class="information_props">Proposer :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${proposer}`)">{{proposer}}</span>
        </div>
        <div class="information_props_wrap" v-if="title">
          <span class="information_props">ProposalTitle :</span>
          <span class="information_value">{{title}}</span>
        </div>
        <div class="information_props_wrap" v-if="proposalType">
          <span class="information_props">ProposalType :</span>
          <span class="information_value">{{proposalType}}</span>
        </div>
        <div class="information_props_wrap" v-if="showInitialDeposit">
          <span class="information_props">InitialDeposit :</span>
          <span class="information_value">{{initialDeposit}}</span>
        </div>
        <div class="information_props_wrap" v-if="description">
          <span class="information_props">Description :</span>
          <span class="information_value">{{description}}</span>
        </div>
        <div class="information_props_wrap" v-if="depositer">
          <span class="information_props">Depositer :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${depositer}`)">{{depositer}}</span>
        </div>
        <div class="information_props_wrap" v-if="showProposalId">
          <span class="information_props">Proposal ID :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/ProposalsDetail/${proposalId}`)">{{proposalId}}</span>
        </div>
        <div class="information_props_wrap" v-if="showVoter">
          <span class="information_props">Voter :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${voter}`)">{{voter}}</span>
        </div>
        <div class="information_props_wrap" v-if="showTypeTransfer">
          <span class="information_props">From :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${fromValue}`)">{{fromValue}}</span>
        </div>
        <div class="information_props_wrap" v-if="showSource">
          <span class="information_props">Source :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${source}`)">{{source}}</span>
        </div>
        <div class="information_props_wrap" v-if="showTypeTransfer">
          <span class="information_props">To :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${toValue}`)">{{toValue}}</span>
        </div>
        <div class="information_props_wrap" v-if="moniker">
          <span class="information_props">Moniker :</span>
          <span class="information_value"><pre class="information_pre">{{moniker}}</pre></span>
        </div>
        <div class="information_props_wrap" v-if="identity">
          <span class="information_props">Identity :</span>
          <span class="information_value">{{identity}}</span>
        </div>
        <div class="information_props_wrap" v-if="owner">
          <span class="information_props">Owner :</span>
          <span class="information_value link_active_style" @click="skipRoute(`/address/1/${owner}`)">{{owner}}</span>
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
          <span class="information_props">Self-Bond :</span>
          <span class="information_value">{{selfBond}}</span>
        </div>
        <div class="information_props_wrap" v-if="details">
          <span class="information_props">Details :</span>
          <span class="information_value">{{details}}</span>
        </div>
        <div class="information_props_wrap" v-if="showVoter">
          <span class="information_props">Option :</span>
          <span class="information_value">{{option}}</span>
        </div>
        <div class="information_props_wrap" v-if="showTypeTransfer || showTypeDeposit">
          <span class="information_props">Amount :</span>
          <span class="information_value">{{amountValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Timestamp :</span>
          <span class="information_value">{{timestampValue}}</span>
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
          <span class="information_value">{{memo}}</span>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';
  export default {

    data() {
      return {
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
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
        showProposalId: false,
        showProposer: false,
        showInitialDeposit: false,
        showTypeTransfer:false,
        showVoter: false,
        showTypeDeposit: false,
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
      let url = `/api/tx/${this.$route.query.txHash}`;
      if(!this.$route.query.txHash){
        return;
      }
      axios.get(url).then((data)=>{
        if(data.status === 200){
          return data.data;
        }
      }).then((data)=>{
        if(data.code === "0"){
          if(data.data && typeof data.data === "object"){
            this.hashValue = data.data.Hash;
            this.blockValue = data.data.BlockHeight;
            this.typeValue = data.data.Type === 'coin'?'transfer':data.data.Type;
            this.timestampValue = Tools.conversionTimeToUTCToYYMMDD(data.data.Timestamp);
            this.gasPrice = Tools.convertScientificNotation2Number(Tools.formaNumberAboutGasPrice(data.data.GasPrice));
            this.gasLimit = data.data.GasLimit;
            this.gasUsedByTxn = data.data.GasUsed;
            this.memo = data.data.Memo ? data.data.Memo : '--';

            if(data.data.Amount && data.data.Amount.length !==0){
              this.amountValue = data.data.Amount.map(item=>{
                item.amount = Tools.convertScientificNotation2Number(Tools.formatNumber(item.amount));
                if(Tools.flTxType(data.data.Type)){
                  return `${item.amount} SHARES`;
                }else{
                  return `${item.amount} ${Tools.formatDenom(item.denom).toUpperCase()}`;
                }
              }).join(',') ;
            }else if(data.data.Amount && Object.keys(data.data.Amount).includes('amount') && Object.keys(data.data.Amount).includes('denom')){
              data.data.Amount =  Tools.convertScientificNotation2Number(Tools.formatNumber(data.data.Amount.amount));
              this.amountValue = `${data.data.Amount.amount} ${Tools.formatDenom(data.data.Amount.denom).toUpperCase()}`
            }else {
              this.amountValue = "--"
            }
            this.actualTxFee = `${Tools.convertScientificNotation2Number(Tools.formatNumber(data.data.Fee.amount))} ${Tools.formatDenom(data.data.Fee.denom).toUpperCase()}`;


            if(data.data.Type === "Transfer" || data.data.Type === "Delegate" || data.data.Type === "BeginUnbonding" || data.data.Type === "CompleteUnbonding"){
              this.showTypeTransfer = true;
              this.fromValue = data.data.From;
              this.toValue = data.data.To;
            }else if(data.data.Type === "CreateValidator" || data.data.Type === "EditValidator"){
              this.owner = data.data.Owner ? data.data.Owner : '--';
              this.moniker = data.data.Moniker ? data.data.Moniker : '--';
              this.pubkey = data.data.Pubkey ? data.data.Pubkey : "--";
              this.identity = data.data.Identity ? data.data.Identity : '--';
              this.website = data.data.Website ? data.data.Website : '--';
              this.details = data.data.Details ? data.data.Details : '--';
              if(data.data.SelfBond && data.data.SelfBond.length !== 0){
                this.selfBond = `${Tools.convertScientificNotation2Number(Tools.formatNumber(data.data.SelfBond[0].amount))} ${Tools.formatDenom(data.data.SelfBond[0].denom).toUpperCase()}`;
              }else {
                this.selfBond = "--"
              }
            }else if(data.data.Type === "BeginRedelegate" || data.data.Type === "CompleteRedelegate" ){
              this.showTypeTransfer = true;
              this.showSource = true;
              this.fromValue = data.data.From ? data.data.From : '';
              this.toValue = data.data.To ? data.data.To : "";
              this.source = data.data.Source ? data.data.Source : "";
            }else if(data.data.Type === "SubmitProposal"){
              this.showProposer = true;
              this.showInitialDeposit = true;
              this.title = data.data.Title ? data.data.Title : '--';
              this.proposer = data.data.From;
              this.proposalType = data.data.ProposalType;
              if(data.data.Amount && data.data.Amount.length !==0){
                this.initialDeposit = data.data.Amount.map(item=>{
                  return `${item.amount} ${Tools.formatDenom(item.denom).toUpperCase()}`;
                }).join(',') ;
              }else {
                this.initialDeposit = "--"
              }
              this.description = data.data.Description ? data.data.Description : '--';
            }else if(data.data.Type === "Deposit"){
              this.showProposalId = true;
              this.showTypeDeposit = true;
              this.proposalId = data.data.ProposalId === 0 ? "--" : data.data.ProposalId;
              this.depositer = data.data.From ? data.data.From : "--";
            }else if(data.data.Type === "Vote"){
              this.showProposalId = true;
              this.showVoter = true;
              this.proposalId = data.data.ProposalId === 0 ? "--" : data.data.ProposalId;
              this.voter = data.data.From ? data.data.From : '--';
              this.option = data.data.Option ? data.data.Option : "--";
            }
          }
        }else {
          console.log(data.msg)
        }


      }).catch(e => {
        console.log(e)
      })
    },
    methods: {
      skipRoute(path) {
        this.$router.push(path);
      }
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';
  .information_pre{
    color: #a2a2ae;
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
            .information_props{
              width:1.5rem;
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
          border-bottom:0.01rem solid #eee;
          margin-bottom:0.05rem;
          .information_value{
            overflow-x:auto;
            -webkit-overflow-scrolling:touch;

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
        font-size: 0.22rem;
        color: #a2a2ae;
      }
    }
  }
  .link_active_style{
    color:#3598db !important;
    cursor:pointer;
  }
</style>
