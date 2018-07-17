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
          <span class="information_value">{{heightValue}}</span>
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
          <a class="information_value">{{transactionsValue}}</a>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Fee:</span>
          <span class="information_value">{{feeValue}}</span>
        </div>
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
    <div :class="transactionsDetailWrap">
      <p class="transaction_information_content_title">Precommit Details</p>
      <div class="block_detail_table_wrap">
        <blocks-list-table :items="items" :type="'5'"></blocks-list-table>
      </div>
    </div>



  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';
  import BlocksListTable from './table/BlocksListTable.vue';

  export default {
    components:{
      BlocksListTable,
    },
    watch: {

    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
        hashValue: '',
        heightValue: '',
        timestampValue: '',
        blockHashValue:'',
        transactionsValue:'',
        feeValue: '',
        lastBlockHashValue:'',
        precommitValidatorsValue:'',
        votingPowerValue:'',
        items:[],

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
      this.getBlockInformation();
    },
    methods: {
      getBlockInformation(){
        let url = `/api/block/${this.$route.params.height}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          console.log(data)

          let denominator = 0;
          data.Validators.forEach(item=>denominator += item.VotingPower);
          let numerator = 0;
          for(let i = 0; i < data.Block.LastCommit.Precommits.length; i++){
            for (let j = 0; j < data.Validators.length; j++){
              if(data.Block.LastCommit.Precommits[i].ValidatorAddress === data.Validators[j].Address){
                numerator += data.Validators[j].VotingPower;
                break;
              }
            }
          }
          this.items = data.Block.LastCommit.Precommits.map(item=>{
              return {
                Address:item.ValidatorAddress,
                Index:item.ValidatorIndex,
                Round:item.Round,
                Signature:item.Signature.Type,
              }
          });
          if(data){
            this.hashValue = data.Height;
            this.heightValue = data.Height;
            this.timestampValue = data.Time;
            this.blockHashValue = data.Hash;
            this.transactionsValue = data.NumTxs;
            this.feeValue = '';
            this.lastBlockHashValue = data.Block.LastCommit.BlockID.Hash;
            this.precommitValidatorsValue = data.Validators.length !== 0?`${data.Block.LastCommit.Precommits.length}/${data.Validators.length}`:'';
            this.votingPowerValue = denominator !== 0? `${numerator/denominator*100}%`:'';
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
      .block_detail_table_wrap{
        width:100%;
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
      .block_detail_table_wrap{
        width:100%;
        overflow-x: scroll;
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
