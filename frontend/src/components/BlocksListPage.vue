<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <div :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">{{listTitleName}}</span>
        <!--<span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>-->
        <span class="blocks_list_page_wrap_hash_var for_block"
              v-show="$route.query.block || $route.query.address">
          {{blockVar}}
        </span>
      </div>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination total_num">
        <span class="blocks_list_page_wrap_hash_var" v-show="['1','2','3','4'].includes(type)">{{count}} total</span>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
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
      <div class="pagination" style='margin:0.2rem 0;'>
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
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  export default {
    components:{
      BlocksListTable,
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
      getDataList(currentPage, pageSize, type) {
        this.showLoading = true;
        if (type === '1') {
          this.getBlockList(currentPage, pageSize);
        }else if (type === '2') {
          this.getTransactionList(currentPage, pageSize)
        }
      },
    },
  }
</script>
<style lang="scss" >
  @import '../style/mixin.scss';
  @import '../style/block_list_page.scss';
</style>
