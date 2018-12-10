<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="transferListPageWrap">
        <span class="blocks_list_title">{{listTitleName}}</span>
      </p>
    </div>
    <div :class="transferListPageWrap">
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
      <div class="pagination">
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
          this.getDataList(currentPage, this.pageSize);
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
        this.listTitleName = 'Transfers';
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        transferListPageWrap: 'personal_computer_blocks_list_page',
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
      this.transferListPageWrap = Tools.getClassByWindowInnerWidth(window.innerWidth);
    },
    mounted() {
      this.getDataList(Number(this.$route.query.pagenum), this.pageSize);
      this.listTitleName = 'Transfers';
      this.tableMinWidth = this.minWidth;
    },
    methods: {
      getDataList(currentPage, pageSize) {
        this.showLoading = true;
        let url = `/api/tx/trans/${currentPage}/${pageSize}`;
        let that = this;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          if(data.code == 0){
            this.count = data.data.Count;
            localStorage.setItem("transferListCount",this.count);
            if(data.data.Data){
              this.items = data.data.Data.map(item => {
                let [Amount,Fee] = ['--','--'],objList;
                Amount = Tools.formatTxAmount(this.$route.params.type,item.Amount);
                if(item.Fee.amount && item.Fee.denom){
                  Fee = item.Fee.amount = `${Tools.formatFeeToFixedNumber(item.Fee.amount)} ${Tools.formatDenom(item.Fee.denom).toUpperCase()}`;
                }
                objList = {
                  TxHash: item.Hash,
                  Block: item.BlockHeight,
                  From: item.From ? item.From : (item.DelegatorAddr ? item.DelegatorAddr : ''),
                  To: item.To ? item.To : (item.ValidatorAddr ? item.ValidatorAddr : ''),
                  Amount,
                  Fee,
                  Timestamp: Tools.conversionTimeToUTCToYYMMDD(item.Timestamp),
                };
                return objList;
              })
            }else{
              this.items = [{
                TxHash: '',
                Block: '',
                From: '',
                To: '',
                Amount: '',
                Fee: '',
                Timestamp: '',
              }];
              that.showNoData = true;
            }
            that.showLoading = false;
          }else {
            localStorage.setItem("transferListCount",0);
          }
        })
        .catch(e => {
          localStorage.setItem("transferListCount",0);
          that.showLoading = false;
          console.log(e)
        })
      }
    }
  }
</script>
<style lang="scss" scoped>
  @import "../style/tabListCommonStyle.scss";
</style>
