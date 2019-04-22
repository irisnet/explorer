<template>
  <div class="transactions_detail_wrap">
    <div class="transactions_title_wrap">
      <p :class="transactionsDetailWrap" style="margin-bottom:0;">
        <span class="transactions_detail_title">Address</span>
        <span class="transactions_detail_wrap_hash_var">
          {{address}}
          <i v-if="flValidator" :style="{background:validatorStatusColor}">v</i>
            <img v-show="flShowProfileLogo" class="profile_logo_img" src="../assets/profiler_logo.png">
            <span v-show="flShowValidatorCandidate && flValidator" class="candidate_validator">(This Validator is a Candidate)</span>
            <span v-show="flShowValidatorJailed && flValidator" class="jailed_validator">(This Validator is jailed!)</span>
        </span>
      </p>
    </div>

    <div :class="transactionsDetailWrap">
      <p class="transaction_information_content_title">Address Information</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap" v-show="flValidator">
          <span class="information_props">Self-Bond :</span>
          <span class="information_value">{{balanceValue?balanceValue:'--'}}</span>
        </div>
        <div class="information_props_wrap" v-show="!flValidator">
          <span class="information_props">Balance :</span>
          <span class="information_value">{{balanceValue?balanceValue:'--'}}</span>
        </div>
        <div class="information_props_wrap" v-show="flValidator">
          <span class="information_props">Bonded Stake :</span>
          <span class="information_value information_show_trim">{{depositsValue?depositsValue:'--'}}</span>
        </div>
        <div class="information_props_wrap" v-show="!flValidator">
          <span class="information_props">Delegated :</span>
          <span class="information_value information_show_trim">{{depositsValue?depositsValue:'--'}}</span>
        </div>
        <div class="information_props_wrap" v-show="!flValidator">
          <span class="information_props">Withdraw To :</span>
          <span class="information_value information_show_trim jump_link_style" v-show="withdrawAddress"><router-link :to="`/address/1/${withdrawAddress}`">{{withdrawAddress}}</router-link></span>
          <span class="information_value information_show_trim" v-show="!withdrawAddress">--</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Transactions :</span>
          <span class="information_value">{{transactionsValue?transactionsValue:'--'}}</span>
        </div>
      </div>
    </div>
    <div :class="transactionsDetailWrap" class="address_profile" v-if="flValidator">
      <p class="transaction_information_content_title">Validator Profile</p>
      <div class="transactions_detail_information_wrap">
        <div class="information_props_wrap">
          <span class="information_props">Name :</span>
          <span class="information_value information_show_trim">{{nameValue?nameValue:'--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Pub Key :</span>
          <span class="information_value">{{pubKeyValue?pubKeyValue:'--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Owner :</span>
          <span class="information_value operator_value" v-show="operatorValue"><router-link :to="`/address/1/${operatorValue}`">{{operatorValue}}</router-link></span>
          <span class="information_value" v-show="!operatorValue">--</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Commission Rate :</span>
          <span class="information_value">{{rateValue}}</span>
        </div>

        <div class="information_props_wrap">
          <span class="information_props">Website :</span>
          <span class="information_value" :class="websiteValue && websiteValue !== '--' ? 'link_style' : ''">
            <pre class="information_pre" @click="openUrl(websiteValue)">{{websiteValue}}</pre>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Identity :</span>
          <span class="information_value">
            <pre class="information_pre">{{identity}}</pre></span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Details :</span>
          <span class="information_value"><pre class="information_pre">{{descriptionValue}}</pre></span>
        </div>
      </div>
    </div>
    <div :class="transactionsDetailWrap" class="current_tenure" v-show="flActiveValidator">
      <p class="transaction_information_content_title" style="border-bottom:1px solid #eee">Current Tenure</p>
      <div class="current_tenure_wrap">
        <div class="transactions_detail_information_wrap">
          <div class="information_props_wrap">
            <span class="information_props">Bond Height :</span>
            <span class="information_value">{{bondHeightValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Voting Power :</span>
            <span class="information_value">{{votingPowerValue}}</span>
          </div>
          <div class="information_props_wrap">
            <span class="information_props">Precommited Blocks :</span>
            <span class="information_value">{{precommitedBlocksValue}}</span>
          </div>
        </div>
        <div class="canvas_voting_power" v-show="flActiveValidator">
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
    <div class="line_container_wrap"  v-show="flActiveValidator">
      <div class="line_container" :class="transactionsDetailWrap">
        <p class="line_history_title">History</p>
        <div class="line_content">
          <div class="line_echarts_content">
            <div  class="line_left_container" style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
              <echarts-validators-line :informationValidatorsLine="informationValidatorsLine" ></echarts-validators-line>
            </div>
            <div class="line_tab_content">
              <div v-for="(item,index) in tabVotingPower" @click="getValidatorHistory(item.title,index)"
                   :class="item.active ? 'border-none' : 'border-block' " >{{item.title}}</div>
            </div>
          </div>
          <div class="line_echarts_content " :class="transactionsDetailWrap === 'personal_computer_transactions_detail_wrap' ?
           'content_right' : 'model_content_right' ">
            <div class="line_right_container" style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
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
    <div class="list_tab_wrap" :class="transactionsDetailWrap">
      <div class="list_tab_content">
        <ul class="list_tab_container">
          <li class="list_tab_item"
              :class="item.active ? 'activeStyle' : 'unactiveStyle'" v-for="(item,index) in txTab"
              @click="tabTxList(index,item.txTabName,1,20)">{{item.txTabName}} ({{item.txTotal}})
          </li>
        </ul>
      </div>
    </div>
    <div :class="transactionsDetailWrap">
      <div class="pagination total_num">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
      <div class="blocks_list_table_container">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="'addressTxList'"
                           :showNoData="showNoData"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination" style='margin:0.2rem 0;'>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
  </div>

  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import BlocksListTable from './table/BlocksListTable';
  import EchartsLine from "./EchartsLine";
  import EchartsValidatorsLine from "./EchartsValidatorsLine";
  import EchartsValidatorsUptimeLine from "./EchartsValidatorsUpTimeLine";
  import SpinComponent from './commonComponents/SpinComponent';
  import Service from "../util/axios";
  import Constants from "../constant/Constant"
  export default {
      watch:{
          currentPage(currentPage) {
              this.currentPage = currentPage;
              new Promise((resolve)=>{
                  this.tabTxList(this.currentTabIndex,this.currentTxTabName,currentPage, this.pageSize);
                  resolve();
              }).then(()=>{
                  document.getElementById('router_wrap').scrollTop = 0;
              })

          },
          $route(){
              Tools.scrollToTop();
              this.type = this.$route.params.type;
              this.tabTxList(this.tabTxListIndex,this.txTabName,this.currentPage,this.pageSize);
              this.getAddressTxStatistics();
              this.getAddressInformation(this.$route.params.param);
              this.getTransactionsList(1,10,this.$route.params.type);
              this.getProfileInformation();
              this.getCurrentTenureInformation();
              this.getValidatorHistory('14days');
              this.getValidatorUptimeHistory('24hours');
          },


      },
      data() {

          return {
              rateValue: '',
              transactionTimer: null,
              devicesWidth: window.innerWidth,
              transactionsDetailWrap: 'personal_computer_transactions_detail',
              activeBtn:0,
              currentPage: 1,
              pageSize: 20,
              addressTxList: "",
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
              operatorValue:'',
              items:[],
              itemsPre:[],
              type:this.$route.params.type,
              totalBlocks:0,
              totalFee:0,
              TransactionsShowNoData:false,
              PrecommitBlocksshowNoData:false,
              transactionsCount:0,
              flValidator:false,
              showNoData:false,
              showLoading:false,
              informationValidatorsLine: {},
              informationUptimeLine:{},
              transactionsTitle: "",
              count: 0,
              txTabName:"Transfers",
              tabTxListIndex:0,
              currentTabIndex:"",
              currentTxTabName:"",
              identity: "",
              withdrawAddress:"",
              flShowValidatorJailed: false,
              flShowValidatorCandidate: false,
              flActiveValidator: true,
              flShowProfileLogo:false,
              validatorStatusColor:"#3598db",
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
              txTab:[
                {
                  "txTabName":"Transfers",
                  "active": true,
                  txTotal:"",
                },
                {
                  "txTabName":"Stakes",
                  "active": false,
                  txTotal:"",

                },
                {
                  "txTabName":"Declarations",
                  "active": false,
                  txTotal: "",
                },
                {
                  "txTabName":"Governance",
                  "active": false,
                  txTotal: ""
                }
              ]
          }
      },
    components:{
      EchartsValidatorsUptimeLine,
      EchartsValidatorsLine,
      EchartsLine,
      BlocksListTable,
      SpinComponent,
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
      } else {
        this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
      }
    },
    mounted() {
      Tools.scrollToTop();
      if(this.$route.params.param.substring(0,3) === this.$Crypto.config.iris.bech32.valAddr){
        this.tabTxList(this.tabTxListIndex,this.txTabName,this.currentPage,this.pageSize);
        this.getAddressInformation(this.$route.params.param);
        this.getTransactionsList(1,10,this.$route.params.type);
        this.getProfileInformation();
        this.getCurrentTenureInformation();
        this.getValidatorHistory('14days');
        this.getValidatorUptimeHistory('24hours');
        this.getAddressTxStatistics();
      }else {
        this.tabTxList(this.tabTxListIndex,this.txTabName,this.currentPage,this.pageSize);
        this.getAddressInformation(this.$route.params.param);
        this.getTransactionsList(1,10,this.$route.params.type);
        this.getProfileInformation();
        this.getAddressTxStatistics();
      }
    },
    methods: {
      getAddressTxStatistics(){
        let url = `/api/txs/statistics?address=${this.$route.params.param}`;
        Service.http(url).then((data) => {
          if(data){
            this.txTab[0].txTotal= data.TransCnt;
            this.txTab[1].txTotal = data.StakeCnt;
            this.txTab[2].txTotal = data.DeclarationCnt;
            this.txTab[3].txTotal = data.GovCnt;
          }
        }).catch((e) => {
          console.error(e)
        })
      },
      tabTxList(index,txTabName,currentPage,pageSize){
        this.currentPage = currentPage;
        this.currentTabIndex = index;
        this.currentTxTabName = txTabName;
        this.showLoading = true;
        for (let txTabIndex = 0; txTabIndex < this.txTab.length; txTabIndex++){
          this.txTab[txTabIndex].active = false;
          this.txTab[index].active = true;
        }
        let url;
        let that = this;
        if(txTabName === 'Transfers'){
          url = `/api/tx/trans/${currentPage}/${pageSize}?address=${this.$route.params.param}`
        }else if(txTabName === 'Stakes'){
          url = `/api/tx/stake/${currentPage}/${pageSize}?address=${this.$route.params.param}`

        }else if(txTabName === 'Declarations'){
          url = `/api/tx/declaration/${currentPage}/${pageSize}?address=${this.$route.params.param}`
        }else if(txTabName === 'Governance'){
          url = `/api/tx/gov/${currentPage}/${pageSize}?address=${this.$route.params.param}`
        }
        Service.http(url).then((txList) => {
          this.showLoading = false;
          this.showNoData = false;
          this.count = txList.Count;
          clearInterval(this.transactionTimer);
          if(txList.Data){
            this.transactionTimer = setInterval(function () {
              let currentServerTime = new Date().getTime() + that.diffMilliseconds;
              that.items = Tools.formatTxList(txList.Data,txTabName,currentServerTime)
            },1000);
          }else {
            let currentServerTime = new Date().getTime() + that.diffMilliseconds;
            that.items = Tools.formatTxList(null,txTabName,currentServerTime);
            that.showNoData = true;
          }
        })
      },
      getAddressInformation(address){
        this.address = address;
        let url = `/api/account/${this.$route.params.param}`;
        Service.http(url).then((result)=>{
          let Amount = '--';
          if(result){
            if (result.amount && result.amount instanceof Array && result.amount[0].denom === Constants.Denom.IRISATTO){
              result.amount[0].amount = Tools.formatNumber(result.amount[0].amount);
              Amount = result.amount.map(listItem => `${listItem.amount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',')
            }else if(result.amount && result.amount instanceof Array && result.amount[0].denom === Constants.Denom.IRIS){
              Amount = result.amount.map(listItem => `${listItem.amount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',')
            }else if(result.amount && Object.keys(result.amount).includes('amount') && Object.keys(result.amount).includes('denom') && result.amount.denom === Constants.Denom.IRISATTO){
              Amount = `${Tools.formatNumber(result.amount)} ${result.denom.toUpperCase()}`
            }else if(result.amount && Object.keys(result.amount).includes('amount') && Object.keys(result.amount).includes('denom') && result.amount.denom === Constants.Denom.IRIS){
              Amount = `${result.amount}${result.denom.toUpperCase()}`
            }else {
              Amount = ''
            }
          }
          this.flShowProfileLogo = result.isProfiler;
          if(result.deposits){
            let deposits = Tools.formatToken(result.deposits);
            this.depositsValue = result.deposits.amount && result.deposits.amount !== 0 && result.deposits.amount !== '' ?`${deposits.amount} ${deposits.denom}`  : '';
          }
          this.withdrawAddress = result.withdrawAddress ? result.withdrawAddress : '--';
          this.balanceValue = Amount;

        }).catch(e =>{
          console.error(e)
        })
      },
      getProfileInformation(){
        let url = `/api/stake/candidate/${this.$route.params.param}`;
        Service.http(url).then((result)=>{
          if(result){
            let validator = result.validator;
            this.flValidator = true;
            if(validator.jailed === true){
              this.flActiveValidator = false;
              this.flShowValidatorJailed = true;
              this.validatorStatusColor = "#f00";
              this.votingPowerValue = Tools.formatStringToNumber(validator.original_tokens);
            }else{
              if(validator.status === 'Unbonded' || validator.status === 'Unbonding' ){
                this.flShowValidatorCandidate = true;
                this.validatorStatusColor = "#45B035";
                this.flActiveValidator = false;
              }else if(validator.status === "Bonded"){
                this.bondHeightValue = validator.bond_height;
                this.votingPowerValue = validator.voting_power;
              }
            }
            this.rateValue = validator.rate ? `${Tools.formatRate(validator.rate.toString())}%`  : '--';
            this.identity = validator.description && validator.description.identity ? validator.description.identity : "--";
            this.nameValue = validator.description && validator.description.moniker ? validator.description.moniker : '--';
            this.pubKeyValue = validator.pub_key ? validator.pub_key : "--";
            this.websiteValue = validator.description && validator.description.website?validator.description.website:'--';
            this.descriptionValue= validator.description && validator.description.details ? validator.description.details : "--";
            this.commissionRateValue = '';
            this.announcementValue = '';
            this.operatorValue = validator.owner;
          }else{
            this.flValidator = false;
            this.flActiveValidator = false;
          }

        }).catch(err => {
          this.flActiveValidator = false;
          console.error(err)
        })
      },
      getCurrentTenureInformation(){
        let url = `/api/stake/candidate/${this.$route.params.param}/status`;
        Service.http(url).then((data)=>{
          if(data){
            this.precommitedBlocksValue = data.precommit_cnt ? data.precommit_cnt : '--';
            this.returnsValue = '';
            this.firstPercent = data.up_time ? `${data.up_time}%` : "--";
          }
        }).catch(err => {
          console.error(err)
        })
      },
      getTransactionsList(){
        let url = `/api/txsByAddress/${this.$route.params.param}/1/30`;
        Service.http(url).then((data)=>{
          if(data){
            this.transactionsCount = data.Count;
            this.transactionsValue = data.Count;
            this.TransactionsShowNoData = true;
          }
        }).catch(e => {
          console.error(e)
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
        Service.http(url).then((validatorVotingPowerList)=>{
          if(validatorVotingPowerList){
            let seriesData = [], noDatayAxisDefaultMaxByValidators;
            let maxPowerValue = 0;
            validatorVotingPowerList.forEach(item=>{
              if(item.Power > maxPowerValue){
                maxPowerValue = item.Power
              }
              if(item.Power == 0){
                item.Power = ""
              }
              let obj =[];
              obj[0] = Tools.conversionTimeToUTCByValidatorsLine(item.Time);
              obj[1] = item.Power;
              seriesData.push(obj);
            });
            if(maxPowerValue < 100){
              noDatayAxisDefaultMaxByValidators = "100"
            }
            this.informationValidatorsLine = {seriesData,noDatayAxisDefaultMaxByValidators};

          }
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
        Service.http(url).then((data)=>{
          if(data) {
            data.forEach(item => {
              let notValidatorTag = -1;
              if(item.Uptime === notValidatorTag){
                item.Uptime = ""
              }
            });
            let xData , currayDate;
            if (tabTime == "24hours") {
              data.forEach(item => {
                //兼容火狐
                let hourseconds = item.Time.replace("-", "/").replace("-", "/")+":00:00";

                let changeMilliseconds = new Date(hourseconds).getTime();
                //展示需加一小时
                changeMilliseconds = changeMilliseconds + 60*60*1000;
                item.Time = Tools.formatDateYearAndMinutesAndSeconds(changeMilliseconds).split(":")[0];
              });
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
                for (var lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
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
                for (var lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                  millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                  let complementdate = Tools.formatDateYearToDate(millisecondstime);

                  data.unshift({Time: complementdate, Uptime: ""});
                }
              }

              xData = data.map(item => `${String(item.Time).substr(5, 2)}/${String(item.Time).substr(8, 2)}`);
            }
            let seriesData = data.map(item => item.Uptime.toString().split(".")[0]);
            let noDatayAxisDefaultMax ;

            for(let seriesDataIndex = 0 ;seriesDataIndex < seriesData.length; seriesDataIndex++){
              if(seriesData[seriesDataIndex] === ""){
                noDatayAxisDefaultMax = "100"
              }
            }
            this.informationUptimeLine = {xData, seriesData,noDatayAxisDefaultMax};

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

        }).catch(e => {
          console.error(e)
        })
      },

      openUrl(url){
        if(url && url !== '--'){
          if(!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)){
            url = `http://${url}`
          }
          window.open(url)
        }
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
    .pagination {
      @include flex;
      justify-content: flex-end;
      @include borderRadius(0.025rem);
      height:0.3rem;
      li{
        height:0.3rem !important;
      }
    }
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
        @include fontWeight;
        margin-bottom:0;
        border-bottom:0.01rem solid #efefef;
      }
      @include pcCenter;
      .transactions_detail_information_wrap{
        .information_props_wrap{
          @include flex;
          .information_props{
            min-width:1.5rem;
            font-size:0.14rem;
            color:#000000;
          }
          .operator_value{
            cursor: pointer;
            color: #3598db !important;
            a{
              color: #3598db !important;
            }
          }
          .information_value{
            word-break: break-all;
            color: #a2a2ae;
            font-size:0.14rem;
            /*flex:1;*/
          }
          .link_style{
            pre{
              color: #3598db !important;
              cursor: pointer;
            }
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
        i{
          font-style:normal;
          padding: 0 0.07rem;
          background:#3598db;
          border-radius:0.04rem;
          color:#ffffff;
          font-size:0.22rem;
        }
        .profile_logo_img{
          margin-bottom:0.03rem;
          border-radius: 0.04rem;
          margin-left: 0.05rem;
        }
        span{
          display: inline-block;
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
        @include fontWeight;
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
            -webkit-overflow-scrolling:touch;
            color: #a2a2ae;
          }
          .link_style{
            pre{
              color: #3598db !important;
              cursor: pointer;
            }
          }
          .operator_value{
            cursor: pointer;
            color: #3598db !important;
            a{
              color: #3598db !important;
            }
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
        @include fontWeight;
      }
      .transactions_detail_wrap_hash_var {
        overflow-x: auto;
        -webkit-overflow-scrolling:touch;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #a2a2ae;
        display: flex;
        i{
          font-style:normal;
          padding:0.02rem 0.07rem;
          background:#3598db;
          border-radius:0.04rem;
          color:#ffffff;
          font-size:0.18rem;
          margin: 0 0.11rem;
        }
        .profile_logo_img{
          border-radius: 0.04rem;
          margin-left: 0.03rem;
          margin-right: 0.03rem;
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
      width:100%!important;
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
          -webkit-overflow-scrolling:touch;
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
    padding-bottom: 0.2rem;
    .line_container{
      width: 80%;
      min-width: 3.2rem;
      max-width: 12.8rem;
      margin: 0 auto;
      .line_history_title{
        height:0.5rem;
        line-height: 0.5rem;
        font-size: 0.18rem;
        padding-left: 0.2rem;
        color: #000;
        @include fontWeight;
      }
      .line_content{
        @include flex;
        .line_echarts_content{
          max-width: 50%;
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
      font-size: 0.12rem !important;
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
    &:hover{
      cursor: pointer;
    }
  }
  .border-block{
    color: #a2a2ae!important;
    border-top: 0.01rem solid #e4e4e4 !important;
    &:hover{
      cursor: pointer;
    }
  }
  .mobile_transactions_detail_wrap{
    width: 100%!important;
    .line_content {
      @include flex;
      flex-direction: column;
      .line_echarts_content {
        max-width: 100%!important;
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
  .information_show_trim{
    white-space: pre-wrap ;
  }
  .jump_link_style{
    cursor: pointer;
    color: #3598db !important;
    a{
      color: #3598db !important;
    }
  }
  .list_tab_wrap{
    padding: 0!important;
  }
  .mobile_transactions_detail_wrap{
    .list_tab_content{
      width: 100%;
      margin-bottom: 0.2rem;
      border-bottom: 0.01rem solid #fff!important;
      overflow-x: auto;
      .list_tab_container{
        @include flex;
        height: 0.38rem;
        min-width: 4rem;
        max-width: 12.8rem;
        padding-left: 0!important;
        margin-left: 0!important;
        .list_tab_item{
          text-align: center;
          line-height: 0.38rem;
          width: 1.54rem;
          color: #A2A2AE;
          border-left: none!important;
          border-top: 0.01rem solid #e4e4e4;
          border-right: none!important;
          border-bottom: 0.02rem solid #e4e4e4;
        }
        }
        .activeStyle{
          color: #3598db!important;
          border-top: 0.01rem solid #e4e4e4 !important;
          border-bottom: 0.02rem solid #3598db !important;
          z-index: 5;
        }
      }
    .no_data_show{
      @include flex;
      justify-content: center;
      border-top:0.01rem solid #eee;
      border-bottom:0.01rem solid #eee;
      font-size:0.14rem;
      height:3rem;
      align-items: center;
    }
    .blocks_list_table_container{
      position:relative;
      overflow-x: auto;
      -webkit-overflow-scrolling:touch;
      top: 0.2rem;
      padding-bottom: 0.2rem;
    }
  }
  .candidate_validator{
    width: 2rem;
    white-space: nowrap;
    color: #45B035;
    @include fontSize;
    padding-left: 0.09rem;
  }
  .jailed_validator{
    width: 2rem;
    color: #f00;
    white-space: nowrap;
    @include fontSize;
    padding-left: 0.09rem;
  }
  .candidate_color{
    background: #45B035!important;
  }
  .jailed_color{
    background: #f00!important;
  }
.personal_computer_transactions_detail_wrap{
  .list_tab_content{
   margin-bottom: 0.2rem;
  }
  .list_tab_wrap{
    width: 100%;
    padding-top: 0.44rem;
    padding-bottom: 0.2rem;
    overflow-x: auto;
    .list_tab_content{
      width: 100%;
      border-bottom: 0.01rem solid #3598db;
      .list_tab_container{
        @include flex;
        height: 0.38rem;
        min-width: 4rem;
        max-width: 12.8rem;
        padding-left: 0.2rem;
        .list_tab_item{
          position: relative;
          top: 0.01rem;
          text-align: center;
          line-height: 0.38rem;
          width: 1.54rem;
          color: #A2A2AE;
          border-left: 0.01rem solid #e4e4e4;
          border-top: 0.01rem solid #e4e4e4;
          border-right: 0.01rem solid #e4e4e4;
          border-bottom: 0.01rem solid #3598db;
          cursor: pointer;
        }
      }
    }
    .activeStyle{
      color: #3598db!important;
      border-top: 0.01rem solid #3598db !important;
      border-right: 0.01rem solid #3598db !important;
      border-left: 0.01rem solid #3598db !important;
      border-bottom: 0.01rem solid #fff !important;
    }
  }
  .no_data_show{
    @include flex;
    justify-content: center;
    border-top:0.01rem solid #eee;
    border-bottom:0.01rem solid #eee;
    font-size:0.14rem;
    height:3rem;
    align-items: center;
  }
  .blocks_list_table_container{
    position:relative;
    overflow-x: auto;
    -webkit-overflow-scrolling:touch;
    top: 0.2rem;
    padding-bottom: 0.2rem;
  }
}
  .information_pre{
    a{
      color: #3598db !important;
    }
  }
</style>
