<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="blocksListPageWrap">
        <span class="blocks_list_title">{{listTitleName}}</span>
      </p>
    </div>
    <div :class="blocksListPageWrap">
      <div class="pagination total_num">
        <span class="blocks_list_page_wrap_hash_var" v-show="count">{{count}} total</span>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>

      <div class="list_wrap">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="this.$route.params.type"
                           :minWidth="tableMinWidth"
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
  import Tools from '../common/Tools';
  import axios from 'axios';
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';

  export default {
    name:"blockListPage",
    components:{
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        this.$router.push({
          path: this.$route.path,
          query:{
            pagenum: this.currentPage
          }
        });
        new Promise((resolve)=>{
          this.getDataList(currentPage, this.pageSize, this.$route.params.type);
          resolve();
        }).then(()=>{
          document.body.scrollTop = 0;
        })
      },
      $route() {
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = Number(this.$route.query.pagenum);
        this.getDataList(Number(this.$route.query.pagenum), this.pageSize);
        this.showNoData = false;
        this.listTitleName = 'Transfer';
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',
        currentPage: Number(this.$route.query.pagenum),
        pageSize: 30,
        count: localStorage.getItem("transferListCount") ? Number(localStorage.getItem("transferListCount")) : 0,
        items: [],
        type: 'list',
        showNoData: false,
        showLoading: false,
        tableMinWidth: '',
        listTitleName: "",
        minWidth: 12,
      }

    },
    beforeMount() {
      this.type = this.$route.params.type;
      this.blocksListPageWrap = Tools.getClassByWindowInnerWidth(window.innerWidth);
    },
    mounted() {
      this.getDataList(Number(this.$route.query.pagenum), this.pageSize, this.$route.params.type);
      this.listTitleName = 'Transfer';
      this.tableMinWidth = this.minWidth;
    },
    beforeDestroy() {
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
      getDataList(currentPage, pageSize) {
        this.showLoading = true;
        let url;
        let that = this;
        if(this.$route.params.type === 'transfer'){
          this.listTitleName = "Transfer";
          url = `/api/tx/trans/${currentPage}/${pageSize}`
        }
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          this.count = data.Count;
          localStorage.setItem("transferListCount",this.count);
          if(data.Data){
            this.items = data.Data.map(item => {
              let [Amount,Fee] = ['--','--'],objList;
              Amount = Tools.formatTxAmount(this.$route.params.type,item.Amount);
              if(item.Fee.amount && item.Fee.denom){
                Fee = item.Fee.amount = `${Tools.formatFeeToFixedNumber(item.Fee.amount)} ${Tools.formatDenom(item.Fee.denom).toUpperCase()}`;
              }
              objList = {
                TxHash: item.Hash,
                Block:item.BlockHeight,
                From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
                To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
                Amount,
                Fee,
                Timestamp: Tools.conversionTimeToUTCToYYMMDD(item.Timestamp),
              };
              return objList;
            })
          }else{
            if(that.$route.params.param === 'transfer'){
              this.items = [{
                TxHash: '',
                Block:'',
                From:'',
                To:'',
                Amount:'',
                Fee:'',
                Timestamp:'',
              }];
            }
            that.showNoData = true;
          }
          that.showLoading = false;
        })
        .catch(e => {
          that.showLoading = false;
          console.log(e)
        })
      }
    }
  }
</script>
<style lang="scss" >
  @import "../style/tabListCommonStyle.scss";
</style>
