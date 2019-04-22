<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <div :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title" v-show="!$store.state.flShowValidatorStatus">{{listTitleName}}</span>
        <!--<span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>-->
        <span class="blocks_list_page_wrap_hash_var for_block"
              v-show="$route.query.block || $route.query.address">
          {{blockVar}}
        </span>
        <div class="validators_status_tab" v-show="$store.state.flShowValidatorStatus">
          <span class="validators_status_title" v-for="(item,index) in validatorStatusTitleList" :class="item.isActive ? 'active_title' : '' " @click="selectValidatorStatus(index)">{{item.title}}</span>
        </div>
      </div>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination total_num" v-if="!$store.state.flShowValidatorStatus">
        <span class="blocks_list_page_wrap_hash_var" v-show="['1','2','3','4'].includes(type)">{{count}} total</span>
        <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPageNum" v-model="navCurrentPage" use-router></b-pagination-nav>
    </div>
      <div style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="this.$route.params.type"
                           :minWidth="tableMinWidth"
                           :showNoData="showNoData" :status="this.$route.params.param"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination" :class="$store.state.flShowValidatorStatus ? 'total_num' : '' " style='margin:0.2rem 0;'>
        <span v-if="$store.state.flShowValidatorStatus" class="blocks_list_page_wrap_hash_var" v-show="['1','2','3','4'].includes(type)">{{count}} total</span>
        <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPageNum" v-model="navCurrentPage" use-router></b-pagination-nav>
      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import Constant from "../constant/Constant"
  import Service from "../util/axios"
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  export default {
    components:{
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      $route() {
        clearInterval(this.timer);
        clearInterval(this.transactionTimer);
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = 1;
        this.showNoData = false;
        switch (this.$route.params.type) {
          case '1': this.listTitleName = 'Blocks';
                  break;
        }
        this.computeMinWidth();
        this.getDataList(this.$route.query.page ? this.$route.query.page : this.defaultValidatorPageNumber, this.pageSize, this.$route.params.type);
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',
        blocksValue: '',
        currentPage: 1,
        navCurrentPage: this.$route.query.page ? this.$route.query.page : null,
        pageSize: 30,
        validatorPageSize: 100,
        defaultValidatorPageNumber:1,
        totalPageNum:localStorage.getItem("pagenum") ? Number(localStorage.getItem("pagenum")) : 1,
        count: 0,
        fields: [],
        items: [],
        type: '1',
        titleVar: '',
        showNoData:false,//是否显示列表的无数据
        showLoading:false,
        blockVar:'',
        innerWidth : window.innerWidth,
        tableMinWidth:'',
        listTitleName:"",
        timer: null,
        transactionTimer: null,
        validatorTabIndex: localStorage.getItem('validatorTabIndex') ? localStorage.getItem('validatorTabIndex') : 0,
        validatorStatusTitleList:[
          {
            title:'Active',
            isActive: true,
          },
          {
            title:'Jailed',
            isActive: false,
          },
          {
            title:'Candidate',
            isActive: false,
          }
        ]
      }
    },
    beforeMount() {
      this.type = this.$route.params.type;
      if (window.innerWidth > 910) {
        this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
      } else {
        this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
      }
    },
    mounted() {
      this.getDataList(this.$route.query.page ? this.$route.query.page : this.defaultValidatorPageNumber, this.pageSize, this.$route.params.type);
      switch (this.$route.params.type) {
        case '1': this.listTitleName = 'Blocks';
          break;
      }
      window.addEventListener('resize',this.onresize);
      this.computeMinWidth();
    },
    beforeDestroy() {
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
      linkGen(pageNum){
        return pageNum === 1 ? '?' : `?page=${pageNum}`
      },
      selectValidatorStatus(index){
        this.validatorStatusTitleList.forEach( item => {
          item.isActive = false
        });
        localStorage.setItem('validatorTabIndex',index);
        this.validatorStatusTitleList[index].isActive = true;
        this.getValidatorList(this.defaultValidatorPageNumber,this.validatorPageSize,this.getValidatorStatus(index))
      },
      onresize(){
        this.innerWidth = window.innerWidth;
        if(window.innerWidth > 910){
          this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
        } else {
          this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
        }
      },
      computeMinWidth(){
        if(this.$route.params.type === '1'){
          this.tableMinWidth = 8.5;
        }else if(this.$route.params.type === '2' && this.$route.params.param === 'transfer'){
          this.tableMinWidth = 12;
        }else if(this.$route.params.type === '3' || this.$route.params.type === '4'){
          this.tableMinWidth = 12;
        }
      },
      getBlockList(currentPage, pageSize){
        let url = `/api/blocks?page=${currentPage}&size=${pageSize}`;
        Service.http(url).then((data) => {
          if(data){
            let that = this;
            clearInterval(this.timer);
              this.count = data.Count;
              this.setTotalPageNum(this.count,this.pageSize);
              this.items = data.Data.map(item => {
                let txn = item.num_txs;
                let precommit = item.last_commit && item.last_commit.length !== 0  ? item.last_commit.length : 0;
                let [votingPower,denominator,numerator] = [0,0,0];
                item.validators.forEach(listItem=>votingPower += listItem.voting_power);
                item.validators.forEach(item=>denominator += item.voting_power);
                if(item.last_commit && item.last_commit.length !== 0){
                  for(let i = 0; i < item.last_commit.length; i++){
                    for (let j = 0; j < item.validators.length; j++){
                      if(item.last_commit[i] === item.validators[j].address){
                        numerator += item.validators[j].voting_power;
                        break;
                      }
                    }
                  }
                }
                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                return {
                  Height: item.height,
                  Txn:txn,
                  Time:item.time,
                  Age: Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX),
                  'Precommit Validators':precommit,
                  'Voting Power': denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'',
                };
              })
            this.timer = setInterval(function () {
              that.items = that.items.map(item => {
                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                item.Age = Tools.formatAge(currentServerTime,item.Time,Constant.SUFFIX,Constant.PREFIX)
                return item
              })
            },1000)

          }else{
            this.items = [{Height:'',Txn:'',Fee:'',Age:'','Precommit Validators':'','Voting Power':''}];
            this.showNoData = true;
          }
          this.showLoading = false;
        }).catch(e => {
          console.log(e)
        })
      },
      getTransactionList(currentPage, pageSize){
        let url;
        let that = this;
        if(this.$route.params.param === 'Transfers'){
          this.listTitleName = "Transfers";
          url = `/api/tx/trans/${currentPage}/${pageSize}`
        }else if(this.$route.params.param === 'Stakes'){
          this.listTitleName = "Stakes";
          url = `/api/tx/stake/${currentPage}/${pageSize}`

        }else if(this.$route.params.param === 'Declarations'){
          this.listTitleName = "Declarations";
          url = `/api/tx/declaration/${currentPage}/${pageSize}`
        }else if(this.$route.params.param === 'Governance'){
          this.listTitleName = "Governance";
          url = `/api/tx/gov/${currentPage}/${pageSize}`
        }
        Service.http(url).then((txList) => {
          that.count = txList.Count;
          this.setTotalPageNum(this.count,this.pageSize);
          clearInterval(this.transactionTimer);
          if(txList){
            that.transactionTimer = setInterval(function () {
              let currentServerTime = new Date().getTime() + that.diffMilliseconds;
              that.items = Tools.formatTxList(txList.Data,that.$route.params.param,currentServerTime)
            },1000);
          }else{
            let currentServerTime = new Date().getTime() + that.diffMilliseconds;
            that.items = Tools.formatTxList(null,that.$route.params.param,currentServerTime);
            that.showNoData = true;
          }
          that.showLoading = false;
        })
      },
      getValidatorList(currentPage, pageSize,status){
        let url;
        url = `/api/stake/validators?page=${currentPage}&size=${pageSize}&type=${status}&origin=browser`;
        Service.http(url).then((result) => {
          this.count = result && result.Count ? result.Count : 0;
          result = result && result.Data ? result.Data : null;
          this.setTotalPageNum(this.count,pageSize);
          if(result){
            this.items = result.map((item) => {
              return {
                validatorStatus: status,
                url:require('../assets/header_img.png'),
                moniker: Tools.formatString(item.description.moniker,15,'...'),
                operatorAddress: item.operator_address,
                commission: `${(item.commission.rate * 100).toFixed(2)} %`,
                bondedToken: `${Number(item.tokens).toFixed(2)} ${Constant.CHAINNAME.toLocaleUpperCase()}`,
                uptime: `${(item.uptime * 100).toFixed(2)}%`,
                votingPower: `${(item.voting_rate * 100).toFixed(2)}%`,
                selfBond: `${Number(item.self_bond)} ${Constant.CHAINNAME.toLocaleUpperCase()}`,
                delegatorNum: item.delegator_num,
                bondHeight: item.bond_height,
                identity: item.description.identity,
              }
            });
            this.items = this.getValidatorHeaderImg(this.items);
            this.showNoData = false;
          }else {
            this.showNoData = true;
            this.items = [{
              validatorStatus: status,
              moniker: "",
              operatorAddress: "",
              commission: "",
              bondedToken: "",
              votingPower: "",
              uptime: "",
              selfBond: "",
              delegatorNum:"",
              bondHeight: "",
              identity: "",
            }]
          }
          this.showLoading = false;
        }).catch(e =>{
          this.showLoading = false;
          this.showNoData = true;
          console.log(e)
        });
      },
      getValidatorHeaderImg(data){
        let url = 'https://keybase.io/_/api/1.0/user/lookup.json?fields=pictures&key_suffix=';
        for(let i = 0; i < data.length; i++){
          if(data[i].identity){
            Service.http(`${url}${data[i].identity}`).then(res =>{
              if(res && res.them && res.them[0].pictures && res.them[0].pictures.primary && res.them[0].pictures.primary.url){
                data[i].url = res.them[0].pictures.primary.url;
              }else {
                data[i].url = require('../assets/header_img.png');
              }
            })
          }else {
            data[i].url = require('../assets/header_img.png');
          }
        }
        return data
      },
      getValidatorStatus(index){
        let validatorStatus;
        switch (index) {
          case 0 :
            validatorStatus = 'validator';
            break;
          case 1 :
            validatorStatus = 'jailed';
            break;
          case 2 :
            validatorStatus = 'candidate'
        }
        return validatorStatus
      },
      getDataList(currentPage, pageSize, type) {
        this.showLoading = true;
        if (type === '1') {
          this.getBlockList(currentPage, pageSize);
        }else if (type === '2') {
          this.getTransactionList(currentPage, pageSize)
        }else if (type === '3' || type === '4') {
          this.$store.commit('flShowValidatorStatus',true)
          this.selectValidatorStatus(Number(this.validatorTabIndex))
        }
      },
      setTotalPageNum(count,pageSize){
        if(count === 0 ){
          this.totalPageNum = 1
        }else {
          this.totalPageNum = Math.ceil(count / pageSize);
        }
        localStorage.setItem("pagenum",this.totalPageNum);
      },
    },
    beforeDestroy () {
      let defaultValidatorTabIndex = 0;
      localStorage.setItem('validatorTabIndex',defaultValidatorTabIndex);
    }
  }
</script>
<style lang="scss" >
  @import '../style/mixin.scss';
  @import '../style/block_list_page.scss';
</style>
