<template>
  <div class="transactions_detail_wrap">
    <div class="transactions_title_wrap">
      <p :class="transactionsDetailWrap" style="margin-bottom:0;">
        <span class="transactions_detail_title">Search Results</span>
        <span class="transactions_detail_wrap_hash_var"> for {{this.$route.params.searchContent}}</span>
      </p>
    </div>
    <div :class="transactionsDetailWrap">
      <div v-show="!flshowResult">
        <p class="transaction_information_content_title">Block</p>
        <div class="block_content_container">
          <p  class="block_height_container">
            <span>Height:</span>
            <span @click="toBlockDetail(blockHeight)">{{blockHeight}}</span>
          </p>
          <p class="block_time_container">
            <span>Timestamp</span>
            <span>{{blockTime}}</span>
          </p>
          <p class="block_hash_container">
            <span>Block Hash</span>
            <span>{{blockHash}}</span>
          </p>
        </div>
        <p class="transaction_information_content_title proposal_title">Proposal</p>
        <div class="proposal_content_container">
          <p class="proposal_id_container">
            <span>Proposal id :</span>
            <span @click="toProposalsDetail(proposalId)">{{proposalId}}</span>
          </p>
          <p class="proposal_title_container">
            <span>Title :</span>
            <span>{{proposalTitle}}</span>
          </p>
          <p class="proposal_type_container">
            <span>Type :</span>
            <span>{{proposalType}}</span>
          </p>
          <p class="proposal_status_container">
            <span>Status :</span>
            <span>{{proposalStatus}}</span>
          </p>
          <p class="proposa_time_container">
            <span>Submit Time :</span>
            <span>{{proposalTime}}</span>
          </p>
        </div>
      </div>
      <div class="resultless_container" v-show="flshowResult">
        <div class="resultless_content_container">
          <div class="result_img">
            <img src="../assets/resultless.png">
          </div>
          <p class="resultless_title">There is no valid result</p>
          <p class="try_info">Try to search with Address, Transaction, Block Number, Proposal ID.</p>
          <div class="back_home_btn" @click="backHome">
            <span>Back Home</span>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>

    import Tools from '../common/Tools';
    import axios from 'axios';
    export default {
      name: "searchResult",
      data() {
        return {
          flshowResult: true,
          blockHeight: "",
          blockTime: "",
          blockHash: "",
          proposalId: "",
          proposalTitle: "",
          proposalType: "",
          proposalStatus: "",
          proposalTime: "",
        }
      },
      mounted(){
        if(/^\+?[1-9][0-9]*$/.test(this.$route.params.searchContent)){
          this.flshowResult = false;
          this.searchResult(this.$route.params.searchContent)
        }
      },
      methods:{
        backHome(){
          this.$router.push("/home")
        },
        toBlockDetail(blockHeight){
          this.$router.push(`/blocks_detail/${blockHeight}`)
        },
        toProposalsDetail(proposalId){
          this.$router.push(`/ProposalsDetail/${proposalId}`)
        },
        searchResult(searchValue){
          let searchUrl = `/api/search/${searchValue}`;
          let that = this;
          axios.get(searchUrl).then((data) => {
            if(data.status === 200){
              return data.data
            }
          }).then((searchResult) =>{
            that.flshowResult = false;
              if(searchResult){
                searchResult.forEach((item) => {
                  if(item.Type == "block"){
                    that.blockHeight = item.Data.Height;
                    that.blockTime = Tools.formatDateYearAndMinutesAndSeconds(item.Data.Timestamp);
                    that.blockHash = item.Data.Hash;
                  }else if(item.Type == "proposal"){
                    that.proposalId = item.Data.ProposalId;
                    that.proposalTitle = item.Data.Title;
                    that.proposalType = item.Data.Type;
                    that.proposalStatus = item.Data.Status;
                    that.proposalTime = Tools.formatDateYearAndMinutesAndSeconds(item.Data.Timestamp)
                  }
                })
              }else {
                that.flshowResult = true;
              }
          })
        }
      },
      beforeMount() {
        if (Tools.currentDeviceIsPersonComputer()) {
          this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
        } else {
          this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
        }
      },

    }
</script>

<style lang="scss">
  @import "../style/mixin";
  .resultless_container{
    width: 100%;
    .resultless_content_container{
      width: 100%;
      text-align: center;
      margin-top: 1.6rem;
      .result_img{
        margin: 0 auto;
        width: 1rem;
        height: 1rem;
        img{
          width: 100%;
        }
      }
      .resultless_title{
        margin-top: 0.2rem;
        color: #000;
        font-size: 0.18rem;
      }
      .try_info{
        font-size: 0.14rem;
        color: #A2A2AE;
        margin-top: 19px;
        margin-bottom: 0.44rem !important;
      }
      .back_home_btn{
        width: 1.58rem;
        height: 0.3rem;
        margin: 0 auto;
        background: #3498DB;
        border-radius: 5px;
        color: #fff;
        font-size: 0.14rem;
        line-height: 0.3rem
      }
    }
  }
  .block_content_container{
    margin-left: 0.2rem;
    .block_height_container{
      margin-top: 0.2rem;
      line-height: 1;
      span:nth-child(1){
        display: inline-block;
        text-align: left;
        width: 1.3rem;
        @include fontSize;
      }
      span:nth-child(2){
        text-align: left;
        color: #3598db;
        cursor: pointer;
      }
    }
    .block_time_container{
      margin-top: 0.08rem;
      span:nth-child(1){
        display: inline-block;
        line-height: 1;
        text-align: left;
        width: 1.3rem;
        @include fontSize;

      }
      span:nth-child(2){
        text-align: left;
        color: #A2A2AE;
      }
    }
    .block_hash_container{
      line-height: 1;
      margin-top: 0.08rem;
      span:nth-child(1){
        display: inline-block;
        text-align: left;
        width: 1.3rem;
        @include fontSize;
      }
      span:nth-child(2){
        text-align: left;
        color: #A2A2AE;
      }
    }
  }
  .proposal_content_container{
    margin-top: 0.21rem;
    margin-left: 0.2rem;
    .proposal_id_container{
      line-height: 1;
      span:nth-child(1){
        display: inline-block;
        width: 1.3rem;
        text-align: left;
        @include fontSize;
      }
      span:nth-child(2){
        color: #3598db;
        cursor: pointer;
      }
    }
    .proposal_title_container{
      line-height: 1;
      margin-top: 0.08rem;
      span:nth-child(1){
        display: inline-block;
        line-height: 1;
        text-align: left;
        width: 1.3rem;
        @include fontSize;

      }
      span:nth-child(2){
        text-align: left;
        color: #A2A2AE;
      }
    }
    .proposal_type_container{
      line-height: 1;
      margin-top: 0.08rem;
      span:nth-child(1){
        display: inline-block;
        line-height: 1;
        text-align: left;
        width: 1.3rem;
        @include fontSize;

      }
      span:nth-child(2){
        text-align: left;
        color: #A2A2AE;
      }
    }
    .proposal_status_container{
      line-height: 1;
      margin-top: 0.08rem;
      span:nth-child(1){
        display: inline-block;
        line-height: 1;
        text-align: left;
        width: 1.3rem;
        @include fontSize;

      }
      span:nth-child(2){
        text-align: left;
        color: #A2A2AE;
      }
    }
    .proposa_time_container{
      line-height: 1;
      margin-top: 0.08rem;
      span:nth-child(1){
        display: inline-block;
        line-height: 1;
        text-align: left;
        width: 1.3rem;
        @include fontSize;

      }
      span:nth-child(2){
        text-align: left;
        color: #A2A2AE;
      }
    }
  }
  .proposal_title{
    margin-top: 0.27rem;
  }
  .mobile_transactions_detail_wrap{
    .resultless_content_container{
      margin-top: 0.8rem!important;
    }
  }
</style>
