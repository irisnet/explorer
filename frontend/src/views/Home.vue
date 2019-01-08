<template>
  <div class="home_wrap">
    <div :class="pageClassName">
      <div class="information_preview">
        <div class="information_preview_module">
          <span :class="flFadeInBlockHeight ? 'animation' : '' ">{{currentBlockHeight}}</span>
          <span class="information_module_key">Current Block</span>
        </div>
        <div class="information_preview_module">
          <span :class="flFadeInBlockHeight ? 'animation' : '' " :style="{color:diffSeconds > 120 ? '#ff001b' : ''}">{{lastBlockAge}}</span>
          <span class="information_module_key">Last Block</span>
        </div>
        <div class="information_preview_module">
          <span :class="flFadeInTransaction ? 'animation' : '' ">{{transactionValue}}</span>
          <span class="information_module_key">Transactions</span>
        </div>
        <div class="information_preview_module"
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
                  blockHeightValue: '',
                  timestampValue: '',
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
              this.getBlocksStatusData();
              this.getBlocksList();
              this.getTransactionHistory();
              this.getTransactionList();
              this.getValidatorsList();
          },
          mounted() {
              document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
              let that =this;
              this.timer = setInterval(function () {
                  that.getBlocksStatusData();
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
          getBlocksStatusData() {
              this.flFadeInTransaction = false;
              let url = `/api/chain/status`;
              let lastTransfer =  {};
              Service.http(url).then((data) => {
                let storedLastTransfer  = localStorage.getItem('lastTransfer');
                  if(storedLastTransfer){
                    if(storedLastTransfer.txCount !== data.TxCount
                      || storedLastTransfer.tps !== data.Tps){
                      this.flFadeInTransaction = true
                    }
                  }

                  let num = data.TxCount;
                  if(data) {
                      if(data.TxCount > 1000000000){
                          num = `${(data.TxCount/1000000000).toFixed(2)} B`;
                      }else if(data.TxCount > 1000000){
                          num = `${(data.TxCount/1000000).toFixed(2)} M`;
                      }else if(data.TxCount > 1000){
                          num = `${(data.TxCount/1000).toFixed(2)} K`;
                  }
                  this.transactionValue = `${num}(${data.Tps.toFixed(2)} TPS)`;
                }
                lastTransfer.txCount = data.TxCount;
                lastTransfer.tps = data.Tps;
                localStorage.setItem('lastTransfer',JSON.stringify(lastTransfer))
              }).catch(e => {
                  console.log(e)
              })
          },
          getValidatorsList() {
              let url = `/api/stake/candidatesTop`;
              Service.http(url).then((data) => {
                  let colors = ['#3498db', '#47a2df', '#59ade3', '#6cb7e7', '#7fc2eb', '#91ccef', '#a4d7f3', '#b7e1f7', '#c9ecfb', '#dcf6ff', '#DADDE3',];
                  let [seriesData, legendData] = [[], []];
                  if (data.Candidates instanceof Array) {
                      let totalCount = 0;
                      data.Candidates.forEach(item=>totalCount += item.VotingPower);
                      let others = data.PowerAll - totalCount;
                      let monikerReserveLength = 10;
                      let addressReserveLength = 6;
                      let powerAll = data.PowerAll;
                      for (let i = 0; i < data.Candidates.length; i++) {
                        seriesData.push({
                          value: data.Candidates[i].VotingPower,
                          name: data.Candidates[i].Description.Moniker ? `${Tools.formatString(data.Candidates[i].Description.Moniker,monikerReserveLength,"...")} (${Tools.formatString(data.Candidates[i].Address,addressReserveLength,"...")})` : (data.Candidates[i].Address ? data.Candidates[i].Address : ''),
                          itemStyle: {color: colors[i]},
                          emphasis : {itemStyle:{color: colors[i]}},
                          upTime:`${data.Candidates[i].Uptime}%`,
                          address:data.Candidates[i].Address,
                          powerAll,
                        });
                        legendData.push(data.Candidates[i].Description.Moniker ? `${Tools.formatString(data.Candidates[i].Description.Moniker,monikerReserveLength,"...")} (${Tools.formatString(data.Candidates[i].Address,addressReserveLength,"...")})` : (data.Candidates[i].Address ? data.Candidates[i].Address : ''))
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
          let storedLastBlock = localStorage.getItem('lastBlock');
          if(storedLastBlock){
            if(storedLastBlock.activeValidator !== blockList.Data[0].Block.LastCommit.Precommits.length || storedLastBlock.totalValidator !== blockList.Data[0].Validators.length){
              this.flFadeInValidator = true;
            }
            if(storedLastBlock.numerator !== numerator || storedLastBlock.denominator !== denominator){
              this.flFadeInVotingPower = true
            }
          }
        },
        showBlockFadeinAnimation(blockList){
          let storedLastBlockHeight = localStorage.getItem('lastBlockHeight');
          if(storedLastBlockHeight){
            for(let index = 0; index < blockList.Data.length; index++){
              if(blockList.Data[index].Height > storedLastBlockHeight){
                blockList.Data[index].showAnimation = "show"
              }
            }
          }
        },
        getBlocksList() {
          let url = `/api/blocks/1/10`;
          Service.http(url).then((blockList) => {
            this.hideFadeinAnimation();
            if(blockList.Data){
              let denominator = 0,lastBlock = {};
              blockList.Data[0].Validators.forEach(item=>denominator += item.VotingPower);
              let numerator = 0;
              for(let i = 0; i < blockList.Data[0].Block.LastCommit.Precommits.length; i++){
                for (let j = 0; j < blockList.Data[0].Validators.length; j++){
                  if(blockList.Data[0].Block.LastCommit.Precommits[i].ValidatorAddress === blockList.Data[0].Validators[j].Address){
                    numerator += blockList.Data[0].Validators[j].VotingPower;
                    break;
                  }
                }
              }
              lastBlock.lastBlockHeight = blockList.Data[0].Height;
              lastBlock.numerator = numerator;
              lastBlock.denominator = denominator;
              lastBlock.activeValidator = blockList.Data[0].Block.LastCommit.Precommits.length;
              lastBlock.totalValidator = blockList.Data[0].Validators.length;
              this.validatorValue = `${blockList.Data[0].Block.LastCommit.Precommits.length} voting / ${blockList.Data[0].Validators.length} total`;
              this.votingPowerValue = denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'';
              this.showFadeinAnimation(blockList,numerator,denominator);
              this.showBlockFadeinAnimation(blockList);
              let that = this;
              clearInterval(this.blocksTimer);
              this.blocksTimer = setInterval(function () {
                let storedLastBlockHeight = localStorage.getItem('lastBlockHeight');
                if(storedLastBlockHeight){
                  if(storedLastBlockHeight !== blockList.Data[0].Height){
                    that.flFadeInBlockHeight = true;
                  }
                }
                let currentTime = new Date().getTime() + that.diffMilliseconds;
                that.lastBlockAge = Tools.formatAge(currentTime,blockList.Data[0].Time);
                that.diffSeconds = Math.floor(Tools.getDiffMilliseconds(currentTime,blockList.Data[0].Time)/1000);
                localStorage.setItem('lastBlock',JSON.stringify(lastBlock));
                localStorage.setItem("lastBlockHeight",blockList.Data[0].Height);
                that.currentBlockHeight = blockList.Data[0].Height;
                that.blocksInformation = blockList.Data.map(item => {
                  return {
                    showAnimation: item.showAnimation ? item.showAnimation : "",
                    Height: item.Height,
                    Proposer: item.Hash,
                    Txn: item.NumTxs,
                    Time: Tools.format2UTC(item.Time),
                    Fee: '0 IRIS',
                    age: Tools.formatAge(currentTime,item.Time,Constant.SUFFIX,Constant.PREFIX)
                  };
                });
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
          let url = `/api/txs/1/10`;
          Service.http(url).then((transactionList) => {
            if(transactionList.Data){
              let that = this;
              for (let txIndex = 0; txIndex < transactionList.Data.length; txIndex++){
                if(new Date(transactionList.Data[txIndex].Time).getTime() > localStorage.getItem("lastTxTime")){
                  transactionList.Data[txIndex].showAnimation = "show"
                }
              }
              let lastTxTime = new Date(transactionList.Data[0].Time).getTime();
              clearInterval(this.transfersTimer);
              this.transfersTimer = setInterval(function () {
                localStorage.setItem('lastTxTime',lastTxTime);
                that.transactionInformation = transactionList.Data.map(item => {
                  let [Amount, Fee] = ['--', '--'];
                  if(item.Amount){
                    if (item.Amount instanceof Array) {
                      if(item.Amount.length > 0){
                        item.Amount[0].amount = Tools.formatAmount(item.Amount[0].amount);
                      }
                      if(Tools.flTxType(item.Type)){
                        Amount = item.Amount.map(listItem => `${listItem.amount} SHARES`).join(',');
                      }else {
                        Amount = item.Amount.map(listItem => `${listItem.amount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',');
                      }
                    } else if (item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')) {
                      Amount = `${item.Amount.amount} ${Tools.formatDenom(item.Amount.denom).toUpperCase()}`;
                      if(Tools.flTxType(item.Type)){
                        Amount = `${item.Amount.amount} SHARES`;
                      }
                    }
                  }else {
                    Amount = '';
                  }
                  if(item.ActualFee.amount && item.ActualFee.denom){
                    Fee =  `${Tools.formatFeeToFixedNumber(item.ActualFee.amount)} ${Tools.formatDenom(item.ActualFee.denom).toUpperCase()}`;
                  }
                  let currentTime = new Date().getTime() + that.diffMilliseconds;
                  return {
                    showAnimation: item.showAnimation ? item.showAnimation : '',
                    TxHash: item.TxHash,
                    From: item.From,
                    To: item.To,
                    Type: item.Type === 'coin'?'transfer':item.Type,
                    Fee,
                    Amount,
                    Time: Tools.format2UTC(item.Time),
                    age: Tools.formatAge(currentTime,item.Time,Constant.SUFFIX,Constant.PREFIX)
                  };
                })
              },1000)
            }
          }).catch(e => {
            console.log(e)
          })
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
</style>
