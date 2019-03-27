<template>
  <div class="home_wrap">
    <div :class="pageClassName">
      <div class="information_preview">
        <div class="information_preview_module latest_block_content" @click="toBlockDetail(currentBlockHeight)">
          <span :class="flFadeInBlockHeight ? 'animation' : '' ">{{currentBlockHeight}}</span>
          <span class="information_module_key">Latest Block</span>
        </div>
        <div class="information_preview_module mobile_age_content">
          <span :class="flFadeInBlockHeight ? 'animation' : '' " :style="{color:diffSeconds > 120 ? '#ff001b' : ''}">{{lastBlockAge}}</span>
          <span class="information_module_key">Age</span>
        </div>
        <div class="information_preview_module mobile_transaction_content">
          <span :class="flFadeInTransaction ? 'animation' : '' ">{{transactionValue}}</span>
          <span class="information_module_key">Transactions</span>
        </div>
        <div class="information_preview_module mobile_validator"
             style="cursor:pointer;" @click="skipValidators"
        >
          <span style="text-align:center;" :class="flFadeInValidator ? 'animation' : '' ">{{validatorValue}}</span>
          <span class="information_module_key">Validators</span>
        </div>
        <div class="information_preview_module">
          <span :class="flFadeInVotingPower ? 'animation' : '' ">{{votingPowerValue}}</span>
          <span class="information_module_key">Online Voting Power</span>
        </div>

      </div>
      <div :class="module_item_wrap">
        <div class="home_module_item home_module_item_pie" style="overflow-x: hidden;">
          <echarts-pie :information="information"></echarts-pie>
        </div>
        <div class="home_module_item home_module_item_pie">
          <echarts-line :informationLine="informationLine"></echarts-line>
        </div>
      </div>
      <div :class="module_item_wrap">
        <div class="home_module_item fixed_item_height">
          <home-block-module :moduleName="'Blocks'" :information="blocksInformation"></home-block-module>
        </div>
        <div class="home_module_item fixed_item_height">
          <home-block-module :moduleName="'Transactions'" :information="transactionInformation"></home-block-module>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import Tools from '../util/Tools';
import EchartsPie from "../components/EchartsPie";
import EchartsLine from "../components/EchartsLine";
import HomeBlockModule from "../components/HomeBlockModule";
import Service from '../util/axios';
import Constant from "../constant/Constant"
  export default {
      name: 'app-header',
      components: {EchartsPie, EchartsLine, HomeBlockModule},
          data() {
              return {
                  devicesWidth: window.innerWidth,
                  pageClassName: 'personal_computer_home_wrap',
                  module_item_wrap: 'module_item_wrap_computer',
                  currentBlockHeight: '--',
                  validatorValue: '--',
                  transactionValue: '--',
                  lastBlockAge: '--',
                  votingPowerValue: '--',
                  blockHeightValue: '--',
                  timestampValue: '--',
                  tpsValue:'',
                  information: {},
                  informationLine: {},
                  blocksInformation: [],
                  transactionInformation: [],
                  innerWidth : window.innerWidth,
                  blocksTimer:null,
                  transfersTimer:null,
                  diffSeconds: 0,
                  flFadeInBlockHeight:false,
                  flFadeInTransaction: false,
                  flFadeInValidator:false,
                  flFadeInVotingPower:false,
                  timer: null
              }
          },

          beforeMount() {
              this.getBlocksStatus();
              this.getBlocksList();
              this.getTransactionHistory();
              this.getTransactionList();
              this.getValidatorsList();
          },
          mounted() {
              document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
              let that =this;
              that.getBlocksList();
              that.getTransactionHistory();
              that.getTransactionList();
              that.getValidatorsList();
              this.timer = setInterval(function () {
                  that.getBlocksList();
                  that.getTransactionHistory();
                  that.getTransactionList();
                  that.getValidatorsList();
              },10000);
              window.addEventListener('resize',this.onresize);
              if (window.innerWidth > 910) {
                  this.pageClassName = 'personal_computer_home_wrap';
                  this.module_item_wrap = 'module_item_wrap_computer';
                  if(document.getElementsByClassName('fixed_item_height').length > 0){
                      document.getElementsByClassName('fixed_item_height')[0].style.height = '6.55rem';
                      document.getElementsByClassName('fixed_item_height')[1].style.height = '6.55rem';
                  }
              } else {
                  this.pageClassName = 'mobile_home_wrap';
                  this.module_item_wrap = 'module_item_wrap_mobile';
              }
          },
          beforeDestroy(){
            window.removeEventListener('resize',this.onWindowResize);
            clearInterval(this.timer)
          },
      methods: {
          onresize(){
              this.innerWidth = window.innerWidth;
              if(window.innerWidth > 910){
                  this.pageClassName = 'personal_computer_home_wrap';
                  this.module_item_wrap = 'module_item_wrap_computer';
                  if(document.getElementsByClassName('fixed_item_height').length > 0) {
                      document.getElementsByClassName('fixed_item_height')[0].style.height = '6.55rem';
                      document.getElementsByClassName('fixed_item_height')[1].style.height = '6.55rem';
                  }
              }else {
                  this.pageClassName = 'mobile_home_wrap';
                  this.module_item_wrap = 'module_item_wrap_mobile';
                  if(document.getElementsByClassName('fixed_item_height').length > 0) {
                      document.getElementsByClassName('fixed_item_height')[0].style.height = 'auto';
                      document.getElementsByClassName('fixed_item_height')[1].style.height = 'auto';
                  }
              }
          },
          getBlocksStatus() {
              this.flFadeInTransaction = false;
              let url = `/api/chain/status`;
              let lastTransfer =  {};
              let that = this;
              Service.http(url).then((data) => {
                setTimeout(function () {
                  let storedLastTransfer  = localStorage.getItem('lastTransfer') ? JSON.parse(localStorage.getItem('lastTransfer')) : '';
                  if(storedLastTransfer){
                    if(storedLastTransfer.txCount !== data.TxCount
                      || storedLastTransfer.tps !== data.Tps){
                      that.flFadeInTransaction = true
                    }
                  }
                  that.tpsValue = `(${data.Tps.toFixed(2)} TPS)`;
                  lastTransfer.txCount = data.TxCount;
                  lastTransfer.tps = data.Tps;
                  localStorage.setItem('lastTransfer',JSON.stringify(lastTransfer))
                },1000)

              }).catch(e => {
                  console.log(e)
              })
          },
          getValidatorsList() {
              let url = `/api/stake/candidatesTop`;
              Service.http(url).then((data) => {
                  let colors = ['#3498db', '#47a2df', '#59ade3', '#6cb7e7', '#7fc2eb', '#91ccef', '#a4d7f3', '#b7e1f7', '#c9ecfb', '#dcf6ff', '#DADDE3',];
                  let [seriesData, legendData] = [[], []];
                  if (data.validators instanceof Array) {
                      let totalCount = 0;
                      data.validators.forEach(item=>totalCount += item.voting_power);
                      let others = data.power_all - totalCount;
                      let monikerReserveLength = 10;
                      let addressReserveLength = 6;
                      let powerAll = data.power_all;
                      for (let i = 0; i < data.validators.length; i++) {
                        seriesData.push({
                          value: data.validators[i].voting_power,
                          name: data.validators[i].description && data.validators[i].description.moniker ? `${Tools.formatString(data.validators[i].description.moniker,monikerReserveLength,"...")} (${Tools.formatString(data.validators[i].address,addressReserveLength,"...")})` : (data.validators[i].address ? data.validators[i].address : ''),
                          itemStyle: {color: colors[i]},
                          emphasis : {itemStyle:{color: colors[i]}},
                          upTime:`${data.validators[i].up_time}%`,
                          address:data.validators[i].address,
                          powerAll,
                        });
                        legendData.push(data.validators[i].description && data.validators[i].description.moniker ? `${Tools.formatString(data.validators[i].description.moniker,monikerReserveLength,"...")} (${Tools.formatString(data.validators[i].address,addressReserveLength,"...")})` : (data.validators[i].address ? data.validators[i].address : ''))
                      }

                    if(others > 0 ){
                      seriesData.push({
                        value: others,
                        name:'others',
                        powerAll,
                        itemStyle:{color:colors[10]},
                      });
                    }

                  }
                  this.information = {legendData, seriesData}

              }).catch(e => {
                console.log(e)
              })
          },
          getTransactionHistory() {
          let url = `/api/txsByDay`;
          Service.http(url).then((data) => {
            let maxValue = 0;
            if(data){
              data.forEach(item=>{
                if(item.Count > maxValue){
                  maxValue = item.Count;
                }
              });
              let xData = data.map(item=>`${String(item.Time).substr(5,2)}/${String(item.Time).substr(8,2)}/${String(item.Time).substr(0,4)}`);
              let seriesData = data.map(item=>item.Count);
              this.informationLine = {maxValue,xData,seriesData};
            }
          }).catch(e => {
            console.log(e)
          })
        },
        hideFadeinAnimation(){
          this.flFadeInBlockHeight = false;
          this.flFadeInValidator = false;
          this.flFadeInVotingPower = false;
        },
        showFadeinAnimation(blockList,numerator,denominator){
          let storedLastBlock = localStorage.getItem('lastBlock') ? JSON.parse(localStorage.getItem('lastBlock')) : '';
          if(storedLastBlock){
            if(storedLastBlock.activeValidator !== blockList[0].last_commit.length || storedLastBlock.totalValidator !== blockList[0].validators.length){
              this.flFadeInValidator = true;
            }
            if(storedLastBlock.numerator !== numerator || storedLastBlock.denominator !== denominator){
              this.flFadeInVotingPower = true
            }
          }
        },
        showBlockFadeinAnimation(blockList){
          let storedLastBlockHeight = localStorage.getItem('lastBlockHeight') ? localStorage.getItem('lastBlockHeight') : '';
          if(storedLastBlockHeight){
            for(let index = 0; index < blockList.length; index++){
              if(blockList[index].height > storedLastBlockHeight){
                blockList[index].showAnimation = "show"
              }
            }
          }
        },
        handelTotalTxs(totalTxs){
          let thousand = 1000,million = 1000000,billion = 1000000000;
          if(totalTxs < thousand){
            return totalTxs;
          }else if(totalTxs > thousand){
            return `${(totalTxs/thousand).toFixed(2)} K`;
          }else if(totalTxs > million){
            return `${(totalTxs/million).toFixed(2)} M`;
          }else if(totalTxs > billion){
            return `${(totalTxs/billion).toFixed(2)} B`;
          }
        },
        getBlocksList() {
          let url = `/api/blocks/recent`;
          Service.http(url).then((blockList) => {
            this.getBlocksStatus();
            this.hideFadeinAnimation();
            if(blockList){
              let denominator = 0,lastBlock = {};
              blockList[0].validators.forEach(item=>denominator += item.voting_power);
              let numerator = 0;
              for(let i = 0; i < blockList[0].last_commit.length; i++){
                for (let j = 0; j < blockList[0].validators.length; j++){
                  if(blockList[0].last_commit[i] === blockList[0].validators[j].address){
                    numerator += blockList[0].validators[j].voting_power;
                    break;
                  }
                }
              }
              lastBlock.lastBlockHeight = blockList[0].height;
              lastBlock.numerator = numerator;
              lastBlock.denominator = denominator;
              lastBlock.activeValidator = blockList.Data[0].Block.LastCommit.Precommits.length;
              lastBlock.totalValidator = blockList.Data[0].Validators.length;
              this.validatorValue = `${blockList.Data[0].Block.LastCommit.Precommits.length} Voting / ${blockList.Data[0].Validators.length} Active`;
              this.votingPowerValue = denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'';
              this.showFadeinAnimation(blockList,numerator,denominator);
              this.showBlockFadeinAnimation(blockList);
              let that = this;
              clearInterval(this.blocksTimer);
                let storedLastBlockHeight = localStorage.getItem('lastBlockHeight');
                if(storedLastBlockHeight){
                  if(Number(storedLastBlockHeight) !== blockList[0].height){
                    that.flFadeInBlockHeight = true;
                  }
                }
                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                that.lastBlockAge = Tools.formatAge(currentServerTime,blockList[0].time);
                that.diffSeconds = Math.floor(Tools.getDiffMilliseconds(currentServerTime,blockList[0].time)/1000);
                localStorage.setItem('lastBlock',JSON.stringify(lastBlock));
                localStorage.setItem("lastBlockHeight",blockList[0].height);
                that.currentBlockHeight = blockList[0].height;
                that.blocksInformation = blockList.map(item => {
                  if(!item.num_txs){
                    item.num_txs = 0
                  }
                  return {
                    showAnimation: item.showAnimation ? item.showAnimation : "",
                    Height: item.height,
                    Proposer: item.hash,
                    Txn: item.num_txs,
                    Time: Tools.format2UTC(item.time),
                    Fee: '0 IRIS',
                    time:item.time,
                    age: Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX)
                  };
                });
              this.blocksTimer = setInterval(function () {
                that.transactionValue = that.handelTotalTxs(blockList[0].total_txs);
                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                that.lastBlockAge = Tools.formatAge(currentServerTime,blockList[0].time);
                that.blocksInformation = that.blocksInformation.map(item => {
                  item.age = Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX);
                  return item
                })
              },1000);
            }
          }).catch(e => {
            console.log(e)
          })
        },
        skipValidators(){
          this.$router.push('/validators/3/active');
        },
        getTransactionList() {
          let url = `/api/txs/recent`;
          Service.http(url).then((transactionList) => {
            if(transactionList){
              let that = this;
              for (let txIndex = 0; txIndex < transactionList.length; txIndex++){
                if(new Date(transactionList[txIndex].time).getTime() > localStorage.getItem("lastTxTime")){
                  transactionList[txIndex].showAnimation = "show"
                }
              }
              let lastTxTime = new Date(transactionList[0].time).getTime();
              clearInterval(this.transfersTimer);
              this.transfersTimer = setInterval(function () {
                localStorage.setItem('lastTxTime',lastTxTime);
                that.transactionInformation = transactionList.Data.map(item => {
                  let [Amount, Fee] = ['--', '--'];
                  if(item.Amount){
                    if (item.Amount instanceof Array) {
                      if(item.Amount.length > 0){
                        item.Amount[0].amount = Tools.formatAmount(item.Amount[0].amount);
                        if(!item.Amount[0].denom){
                          Amount = item.Amount.map(listItem => `${listItem.amount} SHARES`).join(',');
                        }else {
                          Amount = item.Amount.map(listItem => `${listItem.amount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',');
                        }
                      }
                    } else if (item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')) {
                      if(!item.Amount.denom){
                        Amount = `${item.Amount.amount} SHARES`;
                      }else {
                        Amount = `${item.Amount.amount} ${Tools.formatDenom(item.Amount.denom).toUpperCase()}`;
                      }
                    }
                  }else {
                    Amount = '';
                  }
                  if(item.ActualFee.amount && item.ActualFee.denom){
                    Fee =  `${Tools.formatFeeToFixedNumber(item.ActualFee.amount)} ${Tools.formatDenom(item.ActualFee.denom).toUpperCase()}`;
                  }
                  let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                  return {
                    showAnimation: item.showAnimation ? item.showAnimation : '',
                    TxHash: item.TxHash,
                    From: item.From,
                    To: item.To,
                    Type: item.Type === 'coin'?'transfer':item.Type,
                    Fee,
                    Amount,
                    Time: Tools.format2UTC(item.Time),
                    age: Tools.formatAge(currentServerTime,item.Time,Constant.SUFFIX,Constant.PREFIX)
                  };
                })
              },1000)
            }
          }).catch(e => {
            console.log(e)
          })
        },
        toBlockDetail(blockHeight){
            if(blockHeight && blockHeight !== '--'){
              let path = `/blocks_detail/${blockHeight}`;
              this.$router.push(path)
            }
        }
      }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';
  .home_wrap {
    @include flex();
    @include pcContainer;
    //pc端和移动端公共样式
    .personal_computer_home_wrap, .mobile_home_wrap {
      margin-top: 0.35rem;
      .information_preview {
        @include flex;
        margin-bottom: 0.35rem;

        .information_preview_module {
          border-right: 1px solid #d6d9e0;
          @include flex;
          flex-direction: column;
          align-items: center;
          span:nth-child(1){
            height: 0.27rem;
          }
          &:last-child {
            border-right: none;
          }
          span {
            &:first-child {
              font-size: 0.18rem;
              @include fontWeight;
            }
          }
          .information_module_key {
            font-size: 0.14rem;
            color: #a2a2ae;
          }
        }
      }
      //饼状图
      .home_module_item_pie {
        height: 3.5rem;
      }
    }

    .personal_computer_home_wrap {
      width: 100%!important;
      @include pcCenter;
      .information_preview {
        .information_preview_module {
          flex: 1;
        }
      }

      .module_item_wrap_computer {
        width: 100%;
        @include flex();
        justify-content: space-between;
        .home_module_item {
          flex:1;
          border: 0.01rem solid #d6d9e0;
          margin-bottom: 0.4rem;
          height: 3.54rem;
          &:nth-child(2n+1){
            margin-right:0.4rem;
          }
        }
      }

    }
    .mobile_home_wrap {
      flex-direction: column;
      justify-content: space-between;
      width: 100%;
      padding: 0.1rem;
      .information_preview {
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
        .information_preview_module {
          min-width: 1.6rem;
        }
        .mobile_transaction_content{
          min-width: 2rem !important;
        }
        .mobile_age_content{
          min-width: 1.2rem !important;
        }
        .mobile_validator{
          min-width: 2rem !important;
        }
      }
      .module_item_wrap_mobile {
        @include flex();
        flex-direction: column;
        align-items: center;
        .home_module_item {
          width: 98%;
          margin-bottom: 0.4rem;
          border: 0.01rem solid #d6d9e0;
        }
        .home_module_item_pie {
          height: 100%;
          overflow-x: auto;
          -webkit-overflow-scrolling:touch;
        }
      }
    }
    .home_module_item_status {
      padding: 0.1rem;
      background: #3190e8;

      .current_block {
        @include fontWeight;
        color: #fff;
        display: inline-block;
        height: 0.28rem;
        line-height: 0.28rem;
        border-bottom: 0.01rem solid #ffffff;
        width: 100%;

      }
      .home_status_bottom {
        @include flex();

        .home_status_bottom_content {
          flex: 1;
          @include flex();
          flex-direction: column;
          span {
            font-size: 0.14rem;
            color: #ffffff;
            font-weight: normal;
          }
        }

      }
    }
  }
  .blocks_background_type{
    background: url('../assets/block.png') no-repeat 0 0.02rem;
    text-indent:0.2rem;
    color:#3598db;
  }
  .home_module_transaction_title_wrap{
    @include flex;
    padding:0.2rem;
    height:0.64rem;
    justify-content: space-between;
    background: #efeff1;
    border-bottom:1px solid #e4e4e4;
    align-items: center;
    .home_module_transaction_title{
      font-size:0.18rem;
      @include fontWeight;
    }
    .blocks_background{
      background: url('../assets/blocks.png') no-repeat 0 0.02rem;
      text-indent:0.35rem;
    }
    .transactions_background{
      background: url('../assets/transactions.png') no-repeat 0 0.02rem;
      text-indent:0.3rem;
    }
    .view_all_btn{
      @include viewBtn;
    }
  }
  .transaction_title_name{
    padding-left: 0.1rem;
  }
  .transactions_background_type{
    background: url('../assets/transactions.png') no-repeat 0 0.02rem;
    text-indent:0.2rem;
  }
  pre{
    margin: 0!important;
  }
  .animation{
    @include fadeInAnimation
  }
  .latest_block_content:hover{
    cursor: pointer;
  }
</style>
