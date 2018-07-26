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
          <i :class="acitve?'flag_item_left':'flag_item_left_disabled'" @click="skipNext(-1)"></i>
          <span class="information_value" style="flex:none;">{{heightValue}}</span>
          <i :class="activeNext?'flag_item_right':'flag_item_right_disabled'" @click="skipNext(1)"></i>
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
        <blocks-list-table :items="items" :type="'5'" :showNoData="showNoData"></blocks-list-table>
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

  export default {
    components:{
      BlocksListTable,
    },
    watch: {
      $route(){
        this.getBlockInformation();
        if(Number(this.$route.params.height) <= 0){
          this.acitve = false;
        }else{
          this.acitve = true;
        }
        if(this.maxBlock !== 0){
          if(Number(this.$route.params.height) >= this.maxBlock){
            this.activeNext = false;
          }else{
            this.activeNext = true;
          }
        }


      }
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
        showNoData:false,//列表无数据的时候显示
        acitve:true,
        activeNext:true,
        maxBlock:0,

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
      if(Number(this.$route.params.height) <= 0){
        this.acitve = false;
      }else{
        this.acitve = true;
      }
      this.getMaxBlock();
    },
    updated(){
      console.log(this.activeNext)
    },
    methods: {
      getBlockInformation(){
        let url = `/api/block/${this.$route.params.height}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          if(data){
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
            if(data.Block.LastCommit.Precommits && data.Block.LastCommit.Precommits.length > 0){
              this.items = data.Block.LastCommit.Precommits.map(item=>{
                return {
                  Address:item.ValidatorAddress,
                  Index:item.ValidatorIndex,
                  Round:item.Round,
                  Signature:item.Signature.Type,
                }
              });
            }else{
              this.items = [{
                Address:'',
                Index:'',
                Round:'',
                Signature:'',
              }];
              this.showNoData = true;
            }

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
          }else{
            this.items = [{
              Address:'',
              Index:'',
              Round:'',
              Signature:'',
            }];
            this.showNoData = true;
            this.hashValue = this.$route.params.height;
            this.heightValue = this.$route.params.height;
            this.timestampValue = '';
            this.blockHashValue = '';
            this.transactionsValue = '';
            this.feeValue = '';
            this.lastBlockHashValue = '';
            this.precommitValidatorsValue = '';
            this.votingPowerValue = '';
          }


        })
      },
      skipNext(num){
        if(Number(this.$route.params.height) <= 0){
          this.acitve = false;
          if(num !== -1){
            this.$router.push(`/blocks_detail/${Number(this.$route.params.height)+num}`)
          }
        }else if(Number(this.$route.params.height) >= this.maxBlock){
          if(num !== 1){
            this.$router.push(`/blocks_detail/${Number(this.$route.params.height)+num}`)
          }
        } else{
          this.acitve = true;
          this.$router.push(`/blocks_detail/${Number(this.$route.params.height)+num}`)
        }
      },
      getMaxBlock(){
        let url = `/api/blocks/1/1`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if(data){
            this.maxBlock = data.Data[0].Height;
            if(Number(this.$route.params.height) >= this.maxBlock){
              this.activeNext = false;
            }else{
              this.activeNext = true;
            }
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
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
        .information_props_wrap{
          @include flex;
          .information_props{
            width:1.5rem;
          }
          .flag_item_left{
            display:inline-block;
            width:0.2rem;
            height:0.2rem;
            background: url('../assets/left.png') no-repeat 0 1px;
            margin-right:0.02rem;
            cursor:pointer;
          }
          .flag_item_left_disabled{
            display:inline-block;
            width:0.2rem;
            height:0.2rem;
            margin-right:0.02rem;
            cursor:pointer;
            background: url('../assets/left_disabled.png') no-repeat 0 1px;
          }
          .flag_item_right{
            display:inline-block;
            width:0.2rem;
            height:0.2rem;
            background: url('../assets/right.png') no-repeat 0 0;
            margin-left:0.02rem;
            cursor:pointer;
          }
          .flag_item_right_disabled{
            display:inline-block;
            width:0.2rem;
            height:0.2rem;
            background: url('../assets/right_disabled.png') no-repeat 0 0;
            margin-left:0.02rem;
            cursor:pointer;
          }

        }
      }
      .block_detail_table_wrap{
        width:100%;
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
      .block_detail_table_wrap{
        width:100%;
        overflow-x: auto;
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
        margin-right: 0.02rem;
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
