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
          <span class="information_value">{{blockValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Type:</span>
          <span class="information_value">{{typeValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">From:</span>
          <span class="information_value">{{fromValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">To:</span>
          <span class="information_value">{{toValue}}</span>
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
          <span class="information_props">Fee:</span>
          <span class="information_value">{{feeValue}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Message:</span>
          <span class="information_value">{{messageValue}}</span>
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
        hashValue: '0x636693f1a303aafb4d1d8a19b5a6427c101f7797fa9f0d4e7a35fc837617e86e',
        blockValue: '4832791',
        typeValue: 'delegate',
        fromValue: '0x11a5aa0d30a834cae76d5e0b404222222werwrwerwe33424242423423e0c880e698eb',
        toValue: '0x11a5aa0d30a834cae76d5e0b4042342423423423423423423424234e0c880e698eb',
        timestampValue: '12313213321',
        amountValue: '1231312',
        feeValue: '12313123',
        messageValue: '3242353454'


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
      console.log(this.$route.query);
      let url = `/api/tx/4CC31317C59D81B32937A19FE81E006F59F249F3`;

      axios.get(url).then((data)=>{
        if(data.status === 200){
          /*return data.json();*/
        }
        console.log(data)
      })
    },
    methods: {}
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
          border-bottom:0.1rem solid #eee;
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

      /*.transactions_detail_title{
        height:4rem;
        line-height:4rem;
        font-size:1.8rem;
        color:#555;
        margin-right:2rem;
      }
      .transactions_detail_wrap_hash_var{
        height:4rem;
        line-height:4rem;
        font-size:1.4rem;
        color:#ccc;
      }*/
    }
  }

</style>
