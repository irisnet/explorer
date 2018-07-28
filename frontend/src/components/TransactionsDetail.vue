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
          <span class="information_props">TxHash:</span>
          <span class="information_value">{{hashValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Block Height:</span>
          <span class="information_value" style="color:#3598db;cursor:pointer;" @click="skipRoute(`/blocks_detail/${blockValue}`)">{{blockValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Type:</span>
          <span class="information_value">{{typeValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">From:</span>
          <span class="information_value" style="color:#3598db;cursor:pointer;" @click="skipRoute(`/address/1/${fromValue}`)">{{fromValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">To:</span>
          <span class="information_value" style="color:#3598db;cursor:pointer;" @click="skipRoute(`/address/1/${toValue}`)">{{toValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Timestamp:</span>
          <span class="information_value">{{timestampValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Amount:</span>
          <span class="information_value">{{amountValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Fees:</span>
          <span class="information_value">{{feeValue}}</span>
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
        feeValue: '',


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
        if(data){
          this.hashValue = data.TxHash;
          this.blockValue = data.Height;
          this.typeValue = data.Type === 'coin'?'transfer':data.Type;
          this.fromValue = data.From;
          this.toValue = data.To;
          this.timestampValue = Tools.conversionTimeToUTC(data.Time);
          this.amountValue = data.Amount.map(item=>{
            return `${item.amount} ${item.denom.toUpperCase()}`;
          }).join(',');
          this.feeValue = data.Fee.Amount.map(item=>{
            return `${item.amount} ${item.amount === 0?'IRIS':item.denom.toUpperCase()}`;
          }).join(',');
        }

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
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
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
          border-bottom:0.01rem solid #eee;
          margin-bottom:0.05rem;
          .information_value{
            overflow-x:auto;
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
  }

</style>
