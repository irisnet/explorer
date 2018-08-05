<template>
  <div class="home_wrap">
    <div :class="pageClassName">
      <div class="information_preview">
        <div class="information_preview_module">
          <span>{{marketCapValue}}</span>
          <span class="information_module_key">Market Cap</span>
        </div>
        <div class="information_preview_module">
          <span>{{priceValue}}</span>
          <span class="information_module_key">Price</span>
        </div>
        <div class="information_preview_module"
             style="cursor:pointer;" @click="skipValidators"
        >
          <span style="text-align:center;">{{validatorsValue}}</span>
          <span class="information_module_key">Validators</span>
        </div>
        <div class="information_preview_module">
          <span>{{votingPowerValue}}</span>
          <span class="information_module_key">Voting Power</span>
        </div>
        <div class="information_preview_module">
          <span>{{transactionValue}}</span>
          <span class="information_module_key">Transactions</span>
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
          <home-block-module :title="'Blocks'" :information="blocksInformation"></home-block-module>
        </div>
        <div class="home_module_item fixed_item_height">
          <home-block-module :title="'Transactions'" :information="transactionInformation"></home-block-module>
        </div>
      </div>

    </div>
  </div>

</template>

<script>

  import Tools from '../common/Tools';
  import EchartsPie from "../components/EchartsPie";
  import EchartsLine from "../components/EchartsLine";
  import HomeBlockModule from "../components/HomeBlockModule";
  import axios from 'axios';

  export default {
    name: 'app-header',
    components: {EchartsPie, EchartsLine, HomeBlockModule},
    data() {
      return {
        devicesWidth: window.innerWidth,
        pageClassName: 'personal_computer_home_wrap',//1是显示pc端，0是移动端
        module_item_wrap: 'module_item_wrap_computer',//module_item_wrap_computer是pc端，module_item_wrap_mobile是移动端
        marketCapValue: '--',
        validatorsValue: '',
        transactionValue: '',
        priceValue: '--',
        votingPowerValue: '',
        blockHeightValue: '',
        timestampValue: '',
        information: {},//饼图的所有信息
        informationLine: {},//折线图端所有信息
        blocksInformation: [],
        transactionInformation: [],
        innerWidth : window.innerWidth,

      }
    },
    beforeMount() {

    },
    mounted() {
      document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
      //请求各个列表数据
      this.getBlocksStatusData();
      this.getBlocksList();
      this.getTransactionHistory();
      this.getTransactionList();
      this.getValidatorsList();
      window.addEventListener('resize',this.onresize);
      if (window.innerWidth > 910) {
        this.pageClassName = 'personal_computer_home_wrap';
        this.module_item_wrap = 'module_item_wrap_computer';
        if(document.getElementsByClassName('fixed_item_height').length > 0){
          document.getElementsByClassName('fixed_item_height')[0].style.height = '7.19rem';
          document.getElementsByClassName('fixed_item_height')[1].style.height = '7.19rem';
        }
      } else {
        this.pageClassName = 'mobile_home_wrap';
        this.module_item_wrap = 'module_item_wrap_mobile';
      }
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
      onresize(){
        this.innerWidth = window.innerWidth;
        if(window.innerWidth > 910){
          this.pageClassName = 'personal_computer_home_wrap';
          this.module_item_wrap = 'module_item_wrap_computer';
          if(document.getElementsByClassName('fixed_item_height').length > 0) {
            document.getElementsByClassName('fixed_item_height')[0].style.height = '7.19rem';
            document.getElementsByClassName('fixed_item_height')[1].style.height = '7.19rem';
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
        let url = `/api/chain/status`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          let num = data.TxCount;
          if(data.TxCount > 1000000000){
            num = `${(data.TxCount/1000000000).toFixed(2)} B`;
          }else if(data.TxCount > 1000000){
            num = `${(data.TxCount/1000000).toFixed(2)} M`;
          }else if(data.TxCount > 1000){
            num = `${(data.TxCount/1000).toFixed(2)} K`;
          }
          this.marketCapValue = '--';
          this.priceValue = '--';
          this.transactionValue = `${num}(${data.Tps.toFixed(2)} TPS)`;

        })
      },
      getValidatorsList() {
        let url = `/api/stake/candidatesTop`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          let colors = ['#3498db', '#47a2df', '#59ade3', '#6cb7e7', '#7fc2eb', '#91ccef', '#a4d7f3', '#b7e1f7', '#c9ecfb', '#dcf6ff', '#efeff1',];
          //跟series的name匹配
          let [seriesData, legendData] = [[], []];
          if (data.Candidates instanceof Array) {
            //计算前十的voting power总数；
            let totalCount = 0;
            data.Candidates.forEach(item=>totalCount += item.VotingPower);
            //其他部分的
            let others = data.PowerAll - totalCount;
            for (let i = 0; i < data.Candidates.length; i++) {
              seriesData.push({
                value: data.Candidates[i].VotingPower,
                name: data.Candidates[i].Description.Moniker ? data.Candidates[i].Description.Moniker : (data.Candidates[i].Address ? data.Candidates[i].Address : ''),
                itemStyle: {color: colors[i]},
                upTime:`${data.Candidates[i].Uptime}%`,
                address:data.Candidates[i].Address,
                totalCount,
              });
              legendData.push(data.Candidates[i].Description.Moniker ? data.Candidates[i].Description.Moniker : (data.Candidates[i].Address ? data.Candidates[i].Address : ''))
            }

            if(data.Candidates.length > 10){
              legendData.push('others');
              seriesData.push({
                value:others,
                name:'others',
                itemStyle:{color:colors[10]},
              });
            }

          }
          this.information = {legendData, seriesData}

        })
      },
      getTransactionHistory() {
        let url = `/api/txsByDay`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          //找出最大的数据
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
        })

      },
      getBlocksList() {
        let url = `/api/blocks/1/10`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if(data.Data){
            let denominator = 0;
            data.Data[0].Validators.forEach(item=>denominator += item.VotingPower);
            let numerator = 0;
            for(let i = 0; i < data.Data[0].Block.LastCommit.Precommits.length; i++){
              for (let j = 0; j < data.Data[0].Validators.length; j++){
                if(data.Data[0].Block.LastCommit.Precommits[i].ValidatorAddress === data.Data[0].Validators[j].Address){
                  numerator += data.Data[0].Validators[j].VotingPower;
                  break;
                }
              }
            }
            this.votingPowerValue = denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'';
            this.validatorsValue = `${data.Data[0].Block.LastCommit.Precommits.length} voting / ${data.Data[0].Validators.length} total`;
            this.blocksInformation = data.Data.map(item => {
              return {
                Height: item.Height,
                Proposer: item.Hash,
                Txn: item.NumTxs,
                Time: Tools.conversionTimeToUTC(item.Time),
                Fee: '0 IRIS',
              }
            })
          }
        })
      },
      skipValidators(){
        this.$router.push('/validators/3/0');
      },
      getTransactionList() {
        let url = `/api/txs/1/10`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if(data.Data){
            this.transactionInformation = data.Data.map(item => {
              let [Amount, Fee] = ['', ''];
              if (item.Amount instanceof Array) {
                Amount = item.Amount.map(listItem => `${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
              } else if (item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')) {
                Amount = `${item.Amount.amount} ${item.Amount.denom.toUpperCase()}`;
              } else if (item.Amount === null) {
                Amount = '';
              }
              if (item.Fee.Amount instanceof Array) {
                Fee = item.Fee.Amount.map(listItem => `${listItem.amount} ${listItem.amount === 0?'IRIS':listItem.denom.toUpperCase()}`).join(',');
              } else if (item.Fee.Amount && Object.keys(item.Fee.Amount).includes('amount') && Object.keys(item.Fee.Amount).includes('denom')) {
                Fee = `${item.Fee.Amount} ${item.Fee.Amount}`;
              } else if (item.Fee.Amount === null) {
                Fee = '0 IRIS';
              }
              return {
                TxHash: item.TxHash,
                From: item.From,
                To: item.To,
                Type: item.Type === 'coin'?'transfer':item.Type,
                Fee,
                Amount,
                Time: Tools.conversionTimeToUTC(item.Time),
              };
            })
          }

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
      margin-top: 0.15rem;
      .information_preview {
        @include flex;
        margin-bottom: 0.4rem;

        .information_preview_module {
          border-right: 0.01rem solid #eee;
          @include flex;
          flex-direction: column;
          align-items: center;
          &:last-child {
            border-right: none;
          }
          span {
            &:first-child {
              font-size: 0.18rem;
              font-weight: 500;
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
          border: 0.01rem solid #eee;
        }
        .home_module_item_pie {
          overflow-x: auto;
        }
      }
    }
    .home_module_item_status {
      padding: 0.1rem;
      background: #3190e8;

      .current_block {
        font-weight: 500;
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
</style>
