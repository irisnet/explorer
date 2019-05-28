<template>
    <div class="transaction_list_page_container">
      <div class="transaction_list_title_wrap">
        <div class="transaction_list_title_content">
          <span class="transaction_list_title">{{listTitleName}}</span>
        </div>
      </div>
      <div class="transaction_list_table_container">
        <div class="transaction_list_table_content">
          <div class="pagination_nav_content">
            <span>{{count}} Total</span>
            <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPageNum" v-model="currentPageNum" use-router></b-pagination-nav>
          </div>
          <div class="table_list_content">
            <spin-component :showLoading="flShowLoading"></spin-component>
            <tx-list-page-table :items="txList" :showNoData="showNoData"></tx-list-page-table>
            <!-- <i-view-tx-list-page-table :items="txList" :showNoData="showNoData"></i-view-tx-list-page-table> -->
            <div v-show="showNoData" class="no_data_show">
              No Data
            </div>
          </div>
          <div class="pagination_nav_footer_content">
            <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPageNum" v-model="currentPageNum" use-router></b-pagination-nav>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
  import IViewTxListPageTable from "./table/IViewTxListPageTable";
	import TxListPageTable from "./table/TxListPageTable";
	import Service from "../util/axios";
	import Tools from "../util/Tools";
	import SpinComponent from './commonComponents/SpinComponent';
    export default {
		  name: "TransactionListPage",
        components: {SpinComponent, TxListPageTable, IViewTxListPageTable},
        data() {
            return {
                totalPageNum: sessionStorage.getItem("txpagenum") ? JSON.parse(sessionStorage.getItem("txpagenum")) : 1,
                currentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
                txList: [],
                listTitleName: "",
                count: 0,
                pageSize: 30,
                showNoData: false,
                flShowLoading: false
            }
        },
        created(){
            this.getTransactionList(this.currentPageNum,this.pageSize)
        },
        methods: {
            linkGen(pageNum) {
                return pageNum === 1 ? '?' : `?page=${pageNum}`
            },
            getTransactionList(currentPage, pageSize){
                let url;
                let that = this;
                if(this.$route.params.txType === 'transfers'){
                    this.listTitleName = "Transfers";
                    url = `/api/tx/trans/${currentPage}/${pageSize}`
                }else if(this.$route.params.txType === 'stakes'){
                    this.listTitleName = "Stakes";
                    url = `/api/tx/stake/${currentPage}/${pageSize}`
                }else if(this.$route.params.txType === 'declarations'){
                    this.listTitleName = "Declarations";
                    url = `/api/tx/declaration/${currentPage}/${pageSize}`
                }else if(this.$route.params.txType === 'governance'){
                    this.listTitleName = "Governance";
                    url = `/api/tx/gov/${currentPage}/${pageSize}`
                }
                this.flShowLoading = true;
                Service.http(url).then((txList) => {
                    this.count = txList.Count;
                    this.totalPageNum =  Math.ceil((txList.Count/this.pageSize) === 0 ? 1 : (txList.Count/this.pageSize));
                    sessionStorage.setItem('txpagenum',JSON.stringify(this.totalPageNum));
                    if(txList.Data){
                        this.txList = Tools.formatTxList(txList.Data,that.$route.params.txType)
                    }else{
                        this.txList = Tools.formatTxList(null,that.$route.params.txType);
                        this.showNoData = true;
                    }
                    this.flShowLoading = false;
                }).catch(e => {
                    console.error(e)
                })
            },
        }
	}
</script>

<style scoped lang="scss">
  .transaction_list_page_container{
    .transaction_list_title_wrap{
      background: #efeff1;
      border-bottom: 0.01rem solid #d6d9e0;
      .transaction_list_title_content{
        height:0.62rem;
        display: flex;
        align-items: center;
        max-width: 12.8rem;
        margin: 0 auto;
        .transaction_list_title{
          font-size: 0.22rem;
          font-weight: 500;
          padding-left: 0.2rem;
        }
      }
    }
  }
  .transaction_list_table_container{
    max-width: 12.8rem;
    margin: 0 auto;
    .transaction_list_table_content{
      .pagination_nav_content{
        display: flex;
        justify-content: space-between;
        height: 0.7rem;
        align-items: center;
        span{
          color: #a2a2ae;
          padding-left: 0.2rem;
          font-size: 0.18rem;
        }
      }
      .table_list_content{
        width: 100%;
        overflow-x: auto;
        .no_data_show{
          display: flex;
          justify-content: center;
          border-top:0.01rem solid #eee;
          border-bottom:0.01rem solid #eee;
          font-size:0.14rem;
          height:3rem;
          align-items: center;
        }
      }
      .pagination_nav_footer_content{
        display: flex;
        justify-content: flex-end;
        height: 0.7rem;
        align-items: center;
        margin-bottom: 0.2rem;
      }

    }
  }
</style>
