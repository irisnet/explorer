<template>
  <div class="transactions_detail_wrap">
    <div class="transactions_title_wrap">
      <p :class="transactionsDetailWrap" style="margin-bottom:0;">
        <span class="transactions_detail_title">Address</span>
        <span class="transactions_detail_wrap_hash_var">
          {{address}} <i v-if="showProfile">v</i></span>
      </p>
    </div>

    <div :class="transactionsDetailWrap">
      <p class="transaction_information_content_title">Address Information</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Balance:</span>
          <span class="information_value">{{balanceValue?balanceValue:'--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Deposits:</span>
          <span class="information_value">{{depositsValue?depositsValue:'--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Transactions:</span>
          <span class="information_value">{{transactionsValue?transactionsValue:'--'}}</span>
        </div>
      </div>
    </div>
    <div :class="transactionsDetailWrap" class="address_profile" v-if="showProfile">
      <p class="transaction_information_content_title">Profile</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Name:</span>
          <span class="information_value">{{nameValue?nameValue:'--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Pub Key:</span>
          <span class="information_value">{{pubKeyValue?pubKeyValue:'--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Website:</span>
          <span class="information_value"
             v-show="websiteValue !== '--'"
             style="color:#a2a2ae;">{{websiteValue}}</span>
          <i v-show="websiteValue === '--'" style="font-style:normal;color:#a2a2ae">--</i>
        </div>
        <div class="information_props_wrap" style="border-bottom:0.01rem solid #eee">
          <span class="information_props">Description:</span>
          <span class="information_value">{{descriptionValue}}</span>
        </div>
        <!--<div class="information_props_wrap">-->
          <!--<span class="information_props">Commission Rate:</span>-->
          <!--<span class="information_value">{{commissionRateValue}}</span>-->
        <!--</div>-->
        <!--<div class="information_props_wrap">-->
          <!--<span class="information_props">Announcement:</span>-->
          <!--<span class="information_value">{{announcementValue}}</span>-->
        <!--</div>-->


      </div>
    </div>
    <div :class="transactionsDetailWrap" class="current_tenure" v-show="showProfile">
      <p class="transaction_information_content_title" style="border-bottom:1px solid #eee">Current Tenure</p>
      <div class="current_tenure_wrap">
        <div class="transactions_detail_information_wrap">
          <div class="information_props_wrap">
            <span class="information_props">Bond Height:</span>
            <span class="information_value">{{bondHeightValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Voting Power:</span>
            <span class="information_value">{{votingPowerValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Precommited Blocks:</span>
            <span class="information_value">{{precommitedBlocksValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Returns:</span>
            <span class="information_value">{{returnsValue?returnsValue:'--'}}</span>
          </div>

        </div>
        <div class="canvas_voting_power">
          <div class="progress_wrap">
            <span>Uptime(in last 100)</span>
            <div class="progress_wrap_background">
              <div class="progress_value" :style="`width:${firstPercent}`">{{firstPercent}}</div>
            </div>
          </div>
          <div class="progress_wrap">

          </div>
        </div>
      </div>

    </div>
    <div class="line_container_wrap" v-if="showProfile">
      <div class="line_container" :class="transactionsDetailWrap">
        <p class="line_history_title">History</p>
        <div class="line_content">
          <div class="line_echarts_content">
            <div  class="line_left_container" style="overflow-x: auto;">
              <echarts-validators-line :informationValidatorsLine="informationValidatorsLine" ></echarts-validators-line>
            </div>
            <div class="line_tab_content">
              <div v-for="(item,index) in tabVotingPower" @click="getValidatorHistory(item.title,index)"
                   :class="item.active ? 'border-none' : 'border-block' " >{{item.title}}</div>
            </div>
          </div>
          <div class="line_echarts_content " :class="transactionsDetailWrap === 'personal_computer_transactions_detail_wrap' ?
           'content_right' : 'model_content_right' ">
            <div class="line_right_container" style="overflow-x: auto;">
              <echarts-validators-uptime-line :informationUptimeLine="informationUptimeLine" ></echarts-validators-uptime-line>
            </div>
            <div class="line_tab_content">
              <div v-for="(item,index) in tabUptime" @click="getValidatorUptimeHistory(item.title,index)"
                   :class="item.active ? 'border-none' : 'border-block' ">{{item.title}}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div :class="transactionsDetailWrap" class="transaction_precommit_table">
      <div class="tab_wrap">
        <span @click="activeBtn = 0" :class="activeBtn === 0?'transactions_btn_active':''">Transactions</span>
        <span v-show="showProfile" @click="activeBtn = 1" :class="activeBtn === 1?'transactions_btn_active':''">Precommit Blocks</span>
      </div>
      <div class="table_wrap">
        <div class="transactions_view_all" v-show="activeBtn === 0">
          <span>{{transactionsTitle}}</span>
          <span @click="viewAllClick(2)">View All</span>
        </div>
        <div class="precommit_view_all" v-show="activeBtn === 1">
          <p class="table_instruction">
            <span>Total blocks:</span>
            <span>{{totalBlocks}}</span>
            <!--<span>Total Fees:</span>-->
            <!--<span>{{totalFee}}</span>-->
          </p>
          <span class="view_all_btn" @click="viewAllClick(1)">View All</span>
        </div>

        <div class="transaction_table">
          <blocks-list-table :items="items" :type="'6'" v-show="activeBtn === 0" :showNoData="showNoData1"></blocks-list-table>
          <blocks-list-table :items="itemsPre" :type="'7'" v-show="activeBtn === 1" :showNoData="showNoData2" :minWidth="5.4"></blocks-list-table>
          <div v-show="(activeBtn === 0 && showNoData1) || (activeBtn === 1 && showNoData2)" class="no_data_show">
            No Data
          </div>
        </div>
      </div>
    </div>



  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';
  import BlocksListTable from './table/BlocksListTable';
  import EchartsLine from "./EchartsLine";
  import EchartsValidatorsLine from "./EchartsValidatorsLine";
  import EchartsValidatorsUptimeLine from "./EchartsValidatorsUpTimeLine";

  export default {
    watch:{
      $route(){
        this.type = this.$route.params.type;
        this.getAddressInformation(this.$route.params.param);
        this.getTransactionsList(1,10,this.$route.params.type);
        this.getProfileInformation();
        this.getPrecommitBlocksList();
        this.getCurrentTenureInformation();
        this.getValidatorHistory('14days');
        this.getValidatorUptimeHistory('24hours');
      },
      activeBtn(activeBtn){
        //0是Transactions List 1是Precommit Blocks List
      },

    },

    data() {

      return {
        devicesWidth: window.innerWidth,
        transactionsDetailWrap: 'personal_computer_transactions_detail',//1是显示pc端，0是移动端
        activeBtn:0,
        firstPercent:'',
        secondPercent:'%',
        address:this.$route.params.param,
        balanceValue:'',
        depositsValue:'',
        transactionsValue:'',
        nameValue:'',
        pubKeyValue:'',
        websiteValue:'',
        descriptionValue:'',
        commissionRateValue:'',
        announcementValue:'',
        bondHeightValue:'',
        votingPowerValue:'',
        uptimeValue:'',
        precommitedBlocksValue:'',
        returnsValue:'',
        items:[],
        itemsPre:[],
        type:this.$route.params.type,
        totalBlocks:0,
        totalFee:0,
        showNoData1:false,//无数据的时候显示无数据状态
        showNoData2:false,
        transactionsCount:0,
        showProfile:true,
        informationValidatorsLine: {},//折线图端所有信息
        informationUptimeLine:{},
        transactionsTitle: "",
        tabVotingPower:[
          {
            "title":"14days",
            "active":true
          },
          {
            "title":"30days",
            "active":false
          },
          {
            "title":"60days",
            "active":false
          }
        ],
        tabUptime:[
          {
            "title":"24hours",
            "active":true
          },
          {
            "title":"14days",
            "active":false
          },
          {
            "title":"30days",
            "active":false
          }
        ],
      }
    },
    components:{
      EchartsValidatorsUptimeLine,
      EchartsValidatorsLine,
      EchartsLine,
      BlocksListTable,
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      this.getAddressInformation(this.$route.params.param);
      this.getTransactionsList(1,10,this.$route.params.type);
      this.getProfileInformation();
      this.getPrecommitBlocksList();
      this.getCurrentTenureInformation();
      this.getValidatorHistory('14days');
      this.getValidatorUptimeHistory('24hours')
    },
    methods: {
      getAddressInformation(address){
        this.address = address;
        let url = `/api/account/${this.$route.params.param}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          let Amount = '';
          if(data !== null && data !=="" && data && typeof data === "object"){
            if(data.Amount.length > 0 ){
              data.Amount[0].amount = Tools.formatNumber(data.Amount[0].amount);
            }
          }

          if(data.Amount instanceof Array){
            Amount = data.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
          }else if(data.Amount && Object.keys(data.Amount).includes('amount') && Object.keys(data.Amount).includes('denom')){
            Amount = `${data.Amount.amount} ${data.Amount.denom.toUpperCase()}`;
          }else if(data.Amount === null){
            Amount = '';
          }
          this.balanceValue = Amount;

        })
      },
      //点击view all跳转页面
      viewAllClick(type){
        if(type === 1){
          this.$router.push(`/block/${type}/0?address=${this.$route.params.param}`);
        }else if(type === 2){
          this.$router.push(`/recent_transactions/2/recent?address=${this.$route.params.param}`)
        }

      },
      getProfileInformation(){
        let url = `/api/stake/candidate/${this.$route.params.param}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          if(data && typeof data === "object"){
            this.nameValue = data.Description.Moniker;
            this.pubKeyValue = data.PubKey;
            this.websiteValue = data.Description.Website?data.Description.Website:'--';
            this.descriptionValue= data.Description.Details;
            this.commissionRateValue = '';
            this.announcementValue = '';
            this.votingPowerValue = `${(data.VotingPower/data.PowerAll*100).toFixed(2)}%`;
            this.showProfile = true;
            this.bondHeightValue = data.BondHeight;
          }else{
            this.showProfile = false;
          }

        })
      },
      getCurrentTenureInformation(){
        let url = `/api/stake/candidate/${this.$route.params.param}/status`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          if(data && typeof data === "object"){
            this.precommitedBlocksValue = data.PrecommitCount;
            this.returnsValue = '';
            this.firstPercent = `${data.Uptime}%`;
          }

        })
      },
      getTransactionsList(){
        let url = `/api/txsByAddress/${this.$route.params.param}/1/30`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          if(data.Count > 30 && typeof data === "object"){
            this.transactionsTitle = "Last 30 txn from a total of " + data.Count + " transactions"
          }else {
            this.transactionsTitle = "Last 30 txn"
          }
          this.transactionsCount = data.Count;
          this.transactionsValue = data.Count;
          if(data.Data && typeof data === "object"){
            this.items = data.Data.map(item=>{

              if(item.Amount.length > 0){
                item.Amount[0].amount = Tools.dealWithFees(item.Amount[0].amount);
              }
              let [Amount,Fees] = ['',''];
              if(item.Amount instanceof Array){
                Amount = item.Amount.map(listItem=>`${listItem.amount} ${listItem.denom.toUpperCase()}`).join(',');
                if(item.Type === 'CompleteUnbonding' || item.Type === 'BeginUnbonding'){
                  Amount = item.Amount.map(listItem => `${listItem.amount} shares`).join(',');
                }
              }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
                Amount = `${item.Amount.amount} ${item.Amount.denom.toUpperCase()}`;
                if(item.Type === 'CompleteUnbonding' || item.Type === 'BeginUnbonding'){
                  Amount = `${item.Amount.amount} shares`;
                }
              }else if(item.Amount === null){
                Amount = '';
              }
              if(item.ActualFee.amount && item.ActualFee.denom){
                Fees = item.ActualFee.amount = Tools.formatFeeToFixedNumber(item.ActualFee.amount) + ' ' + item.ActualFee.denom.toUpperCase();

              }
              let type = '';
              if(item.Type === 'Transfer'){
                if(this.$route.params.param === item.From){
                  type = 'Send';
                }else{
                  type = 'Receive';
                }
              }else{
                type = item.Type;
              }
              return {
                TxHash:item.TxHash,
                Block:item.Height,
                From:item.From,
                To:item.To,
                Type:type,
                Amount,
                Fees,
                Timestamp:Tools.conversionTimeToUTC(item.Time),
              }
            })
          }else{
            this.items = [{
              TxHash:'',
              Block:'',
              From:'',
              To:'',
              Type:'',
              Amount:'',
              Fees:'',
              Timestamp:'',
            }];
            this.showNoData1 = true;
          }

        })
      },
      getPrecommitBlocksList(){
        let url = `/api/blocks/precommits/${this.$route.params.param}/1/6`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          this.totalBlocks = data.Count;
          if(data.Data && typeof data === "object"){
            this.itemsPre = data.Data.map(item=>{
              return{
                'Block Height':item.Height,
                Txn:item.NumTxs,
                // Fees:'0 IRIS',
                Timestamp:Tools.conversionTimeToUTC(item.Time),
                'Precommit Validators':item.Validators.length !== 0?`${item.Block.LastCommit.Precommits.length}/${item.Validators.length}`:'',
              }
            });
          }else{
            this.itemsPre = [{
              'Block Height':'',
              Txn:'',
              // Fees:'',
              Timestamp:'',
              'Precommit Validators':'',
            }];
            this.showNoData2 = true;
          }
        })
      },
      getValidatorHistory(tabTime,index){
        if(index !== undefined){
          for(var i = 0; i < this.tabVotingPower.length; i++){
            this.tabVotingPower[i].active = false;
            this.tabVotingPower[index].active = true
          }
        }
        let url;
        if(tabTime == "14days"){
          url = `/api/stake/candidate/${this.$route.params.param}/power/week`;
        }else if(tabTime == "30days"){
          url = `/api/stake/candidate/${this.$route.params.param}/power/month`;
        }else if(tabTime == "60days"){
          url = `/api/stake/candidate/${this.$route.params.param}/power/months`;
        }
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          let maxValue = 0;
          let array = [];
          if(data && typeof data === "object"){
            data.forEach(item=>{
              if(item.Power == 0){
                item.Power = ""
              }
              let obj =[];
              obj[0] = Tools.conversionTimeToUTCByValidatorsLine(item.Time);
              obj[1] = item.Power;
              array.push(obj);
              if(item.Power > maxValue){
                maxValue = item.Power;
              }
            });
          }
          let seriesData = array, noDatayAxisDefaultMax;
          //如果没有votingPower，返回的数据中会默认带两条数据
          if(seriesData.length < 3){
            noDatayAxisDefaultMax = "100"
          }
          this.informationValidatorsLine = {maxValue,seriesData,noDatayAxisDefaultMax};
        })
      },
      getValidatorUptimeHistory(tabTime,index){
        if(index != undefined){
          for(var i = 0; i < this.tabUptime.length; i++){
            this.tabUptime[i].active = false;
            this.tabUptime[index].active = true
          }
        }

        let url;
        if(tabTime == "24hours"){
          url = `/api/stake/candidate/${this.$route.params.param}/uptime/hour `;
        }else if(tabTime == "14days"){
          url = `/api/stake/candidate/${this.$route.params.param}/uptime/week `;
        }else if(tabTime == "30days"){
          url = `/api/stake/candidate/${this.$route.params.param}/uptime/month `;
        }
        axios.get(url).then((data)=>{

          if(data.status === 200){
            return data.data;
          }
        }).then((data)=>{
          if(data && typeof data === "object") {
            let xData , currayDate;
            if (tabTime == "24hours") {
              if(data.length !== 0){
                currayDate = data[0].Time;
              }else {
                currayDate = new Date().toISOString().substr(0,13).replace("T", " ");
              }
              if(data.length < 24){
                let complementHourLength = 24 - data.length;
                let hourTime = currayDate.split(" ")[1];
                let yearAndDayTime = currayDate.split(" ")[0];

                for (let k = 0; k < complementHourLength; k++){
                  hourTime--;
                  //当hourTime的数值为负数的时候，+24格式化成24小时显示
                  if(hourTime < 0){
                    hourTime = 24 + hourTime;
                  }
                  //当小时数为一位的时候补零
                  if(String(hourTime).length < 2){
                    hourTime = "0" + hourTime;
                  }
                  let hoursDate = yearAndDayTime + " " + hourTime;
                  data.unshift({AddressL:data.Address,Time: hoursDate ,Uptime: ""})
                }
              }

              data.forEach((item) => {
                item.Time = item.Time.substr(10, 12)+ ":00";
              });

              xData = data.map(item => item.Time);
            } else {

              let currayDate;
              if(data.length > 2){
                currayDate = data[0].Time;
              }else {
                currayDate = new Date().toISOString();
              }

              if (tabTime == "14days") {
                let dataDateLength = data.length,

                //获取需要补全的天数
                  complementdateLength = 14 - dataDateLength,
                  //从那天需要补全的日期
                  weekDate = new Date(currayDate),
                  millisecondstime = weekDate.getTime(),
                  //24小时的时间戳（毫秒数）
                  dayNumberOfMilliseconds = 60 * 60 * 1000 * 24 ;
                //补全日期的逻辑
                for (let lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                  millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                  let complementdate = Tools.formatDateYearToDate(millisecondstime);

                  data.unshift({Time: complementdate, Uptime: ""});
                }

              } else if (tabTime == "30days") {

                let dataDateLength = data.length,
                  complementdateLength = 30 - dataDateLength,
                  monthDate = new Date(currayDate),
                  millisecondstime = monthDate.getTime(),
                  dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                for (let lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                  millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                  let complementdate = Tools.formatDateYearToDate(millisecondstime);

                  data.unshift({Time: complementdate, Uptime: ""});
                }
              }

              xData = data.map(item => `${String(item.Time).substr(5, 2)}/${String(item.Time).substr(8, 2)}`);
            }
            let seriesData = data.map(item => item.Uptime.toString().split(".")[0]);
            this.informationUptimeLine = {xData, seriesData};
          }else {
            let xData , currayDate  , data = [];
            if (tabTime == "24hours") {
                 currayDate = new Date().toISOString().substr(0, 13).replace("T", " ");
                let complementHourLength = 24;
                let hourTime = currayDate.split(" ")[1];
                let yearAndDayTime = currayDate.split(" ")[0];

                for (let k = 0; k < complementHourLength; k++) {
                  hourTime--;
                  //当hourTime的数值为负数的时候，+24格式化成24小时显示
                  if (hourTime < 0) {
                    hourTime = 24 + hourTime;
                  }
                  //当小时数为一位的时候补零
                  if (String(hourTime).length < 2) {
                    hourTime = "0" + hourTime;
                  }
                  let hoursDate = yearAndDayTime + " " + hourTime;
                  data.unshift({AddressL: data.Address, Time: hoursDate, Uptime: ""})
                }

              data.forEach((item) => {
                item.Time = item.Time.substr(10, 12) + ":00";
              });
              xData = data.map(item => item.Time);

            } else if (tabTime == "14days") {
              currayDate = new Date().toISOString();
                //获取需要补全的天数
              let complementdateLength = 14 ,
                //从那天需要补全的日期
                weekDate = new Date(currayDate),
                millisecondstime = weekDate.getTime(),
                //24小时的时间戳（毫秒数）
                dayNumberOfMilliseconds = 60 * 60 * 1000 * 24 ;
              //补全日期的逻辑
              for (let lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                let complementdate = Tools.formatDateYearToDate(millisecondstime);

                data.unshift({Time: complementdate, Uptime: ""});
              }
              xData = data.map(item => `${String(item.Time).substr(5, 2)}/${String(item.Time).substr(8, 2)}`);
            } else if (tabTime == "30days") {
              currayDate = new Date().toISOString();
              let complementdateLength = 30 ,
                  monthDate = new Date(currayDate),
                  millisecondstime = monthDate.getTime(),
                  dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                  for (let lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                    millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                    let complementdate = Tools.formatDateYearToDate(millisecondstime);
                data.unshift({Time: complementdate, Uptime: ""});
              }
              xData = data.map(item => `${String(item.Time).substr(5, 2)}/${String(item.Time).substr(8, 2)}`);
            }
            let noDatayAxisDefaultMax = "100";
            let seriesData = data.map(item => item.Uptime.toString().split(".")[0]);
            this.informationUptimeLine = {xData, seriesData,noDatayAxisDefaultMax};
          }

        })
      },
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
        color:#000000;
        margin-bottom:0;
        border-bottom:0.01rem solid #efefef;
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
        .information_props_wrap{
          @include flex;
          .information_props{
            width:1.5rem;
            font-size:0.14rem;
            color:#000000;
          }
          .information_value{
            color: #a2a2ae;
            font-size:0.14rem;
            flex:1;
          }
        }
      }
      .transactions_detail_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
        font-weight:500;
      }
      .transactions_detail_wrap_hash_var {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.22rem;
        color: #a2a2ae;
        i{
          font-style:normal;
          padding:0.02rem 0.07rem;
          background:#3598db;
          border-radius:0.04rem;
          color:#ffffff;
          font-size:0.18rem;
        }
      }
    }

    .mobile_transactions_detail_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.1rem;
      .transaction_information_content_title{
        height:0.4rem;
        line-height:0.4rem;
        font-size:0.18rem;
        color:#000000;
        margin-bottom:0;
      }
      .transactions_detail_information_wrap{

        .information_props_wrap{
          @include flex;
          flex-direction:column;
          border-bottom:none !important;
          margin-bottom:0.05rem;
          .information_value{
            overflow-x:auto;
            color: #a2a2ae;
          }
          .information_props{
            font-size:0.14rem;
            color:#000000;
          }


        }
      }
      .transactions_detail_title {
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #000000;
        margin-right: 0.2rem;
        font-weight:500;
      }
      .transactions_detail_wrap_hash_var {
        overflow-x: auto;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #a2a2ae;
        i{
          font-style:normal;
          padding:0.02rem 0.07rem;
          background:#3598db;
          border-radius:0.04rem;
          color:#ffffff;
          font-size:0.18rem;
        }
      }

    }
    .address_profile{
      @include borderRadius(5px);
    }

    //current tenure部分
    .current_tenure{
      margin-top:0.1rem;
      .current_tenure_wrap{
        width:100%;
      }
    }
    .personal_computer_transactions_detail_wrap{
      width:80%;
      .current_tenure_wrap{
        @include flex;
        flex-direction:row;
        .transactions_detail_information_wrap{
          flex:3;
        }
        .canvas_voting_power{
          flex:2;
          padding:0.2rem 0;
          .progress_wrap{
            margin-bottom:0.15rem;
            .progress_wrap_background{
              height:0.3rem;
              background: #efeff1;
              margin-top:0.12rem;
              .progress_value{
                background:#a4d7f4;
                height:100%;
                line-height:0.3rem;
                text-indent:0.2rem;
              }
            }
          }

        }
      }
    }
    .mobile_transactions_detail_wrap{
      width:100%;
      .current_tenure_wrap{
        @include flex;
        flex-direction:column;
      }

      .canvas_voting_power{
        flex:2;
        .progress_wrap{
          margin-bottom:0.15rem;
          .progress_wrap_background{
            height:0.3rem;
            background: #efeff1;
            margin-top:0.12rem;
            .progress_value{
              background:#a4d7f4;
              height:100%;
              line-height:0.3rem;
              padding-left:0.2rem;
            }
          }
        }

      }
    }
    //底部表格部分
    .transaction_precommit_table{
      margin-top:0.28rem;
      margin-bottom:0.4rem;
      .tab_wrap{
        border-bottom:1px solid #d6d9e0;
        span{
          height:0.38rem;
          line-height:0.38rem;
          width:1.5rem;
          display:inline-block;
          color:#fff;
          text-align: center;
          background:rgba(214,217,224,1);
          cursor:pointer;
          margin-bottom:0.2rem;

        }
        .transactions_btn_active{
          background: #3598db;

        }
      }
      .table_wrap{
        width:100%;
        .pagination{
          @include flex;
          justify-content: flex-end;
        }
        .transactions_view_all{
          padding:0.1rem;
          @include flex;
          justify-content:space-between;
          height:0.62rem;
          align-items: center;
          span:nth-child(1){
            color: #a2a2ae;
            font-size: 0.14rem;
          }
          span:nth-child(2){
            @include viewBtn;
          }
        }
        .precommit_view_all{
          padding:0.1rem;
          height:0.62rem;
          @include flex;
          justify-content:space-between;
          align-items: center;
          span{
            font-size:0.14rem;
            color:#a2a2ae;
            &:nth-child(3){
              display:inline-block;
              margin-left:0.1rem;
            }
          }
          .view_all_btn{
            @include viewBtn;
          }
        }
        .transaction_table{
          overflow-x:auto;
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
      }
    }
  }
  .line_container_wrap{
    width: 100%;
    .line_container{
      width: 80%;
      min-width: 3.2rem;
      max-width: 12.8rem;
      margin: 0 auto;
      .line_history_title{
        height:0.5rem;
        line-height: 0.5rem;
        border-bottom: 1px solid #d6d9e0 !important;
        font-size: 0.18rem;
        color: #000;
      }
      .line_content{
        @include flex;
        .line_echarts_content{
          margin-top: 0.2rem;
          flex: 1;
          border: 0.01rem solid #e4e4e4;
          .line_left_container{
            max-width: 6.2rem;
            height: 2.76rem;
          }
          .line_right_container{
            max-width: 6.2rem;
            height: 2.76rem;
          }
        }

      }

    }
  }
  .line_tab_content{
    @include flex;
    width: 100%;
    margin: 0 auto;
    height: 0.42rem;
    div{
      display: inline-block;
      flex: 1;
      text-align: center;
      line-height: 0.4rem;
      font-size: 0.16rem;
      color: #000;
      border-top: 0.01rem solid #e4e4e4;
      border-right: 0.01rem solid #e4e4e4;
    }
    div:nth-child(3){
      border-right: none;
    }
  }
  .content_right{
    margin-left: 0.4rem;
  }
  .border-none{
    color: #000!important;
    border-top: 0.01rem solid #fff !important;
  }
  .border-block{
    color: #a2a2ae!important;
    border-top: 0.01rem solid #e4e4e4 !important;
  }
  .mobile_transactions_detail_wrap{
    width: 100%!important;
    .line_content {
      @include flex;
      flex-direction: column;
      .line_echarts_content {
        margin-top: 0.2rem;
        padding: 0 0.1rem;
        flex: 1;
        border: 0.01rem solid #e4e4e4;

        .line_left_container {
          width: 98%;
          max-width: 98%!important;
          height: 2.76rem;
        }

        .line_right_container {
          width: 98%;
          max-width: 98%;
          height: 2.76rem;
        }

      }
    }
  }
  .model_content_right{
    margin-left: 0 ;
  }
</style>
