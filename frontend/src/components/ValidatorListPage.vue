<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <div :class="blocksListPageWrap" style="margin-bottom:0;">
        <div class="validators_status_tab">
          <span class="validators_status_title" v-for="(item,index) in validatorStatusTitleList" :class="item.isActive ? 'active_title' : '' " @click="selectValidatorStatus(index)">{{item.title}}</span>
        </div>
      </div>
    </div>
    <div :class="blocksListPageWrap" style="margin-top: 0.62rem">
      <div style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
        <spin-component :showLoading="showLoading"/>
        <validator-list-table :items="items" :minWidth="tableMinWidth" :showNoData="showNoData"></validator-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination total_num" style='margin:0.2rem 0;'>
        <span class="blocks_list_page_wrap_hash_var">{{count}} total</span>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
    </div>
  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import Constant from "../constant/Constant"
  import Service from "../util/axios"
  import SpinComponent from './commonComponents/SpinComponent';
  import ValidatorListTable from "./table/ValidatorListTable";
  export default {
    components:{
      ValidatorListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        new Promise((resolve)=>{
          this.getDataList(currentPage, 30, this.$route.params.type);
          resolve();
        }).then(()=>{
          document.getElementById('router_wrap').scrollTop = 0;
        })
      },
      $route() {
        clearInterval(this.timer);
        clearInterval(this.transactionTimer);
        this.items = [];
        this.currentPage = 1;
        this.showNoData = false;
        this.computeMinWidth();
        this.getDataList();
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',
        blocksValue: '',
        currentPage: 1,
        pageSize: 30,
        validatorPageSize: 100,
        defaultValidatorPageNumber:1,
        count: 0,
        fields: [],
        items: [],
        type: '1',
        showNoData:false,//是否显示列表的无数据
        showLoading:false,
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
            title:'Candidate',
            isActive: false,
          },
          {
            title:'Jailed',
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
      window.addEventListener('resize',this.onresize);
      this.computeMinWidth();
    },
    beforeDestroy() {
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
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
        this.tableMinWidth = 12;
      },
      getValidatorList(currentPage, pageSize,status){
        this.pageSize = this.validatorPageSize;
        let url;
        url = `/api/stake/validators?page=${currentPage}&size=${pageSize}&type=${status}&origin=browser`;
        Service.http(url).then((result) => {
          this.count = result && result.Count ? result.Count : 0;
          result = result && result.Data ? result.Data : null;
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
                selfBond: `${Number(item.self_bond.match(/\d*(\.\d{0,4})?/)[0])} ${Constant.CHAINNAME.toLocaleUpperCase()}`,
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
            validatorStatus = 'candidate';
            break;
          case 2 :
            validatorStatus = 'jailed'
        }
        return validatorStatus
      },
      getDataList() {
        this.showLoading = true;
        this.selectValidatorStatus(Number(this.validatorTabIndex))
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
  @import '../style/validator_list_page.scss';
</style>
