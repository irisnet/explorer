<template>
    <div class="transaction_list_page_container">
        <div class="transaction_list_title_wrap">
            <div class="transaction_list_title_content">
                <span class="transaction_list_title">{{count}} {{listTitleName}}</span>
            </div>
        </div>
        <div class="transaction_list_table_container">
            <div class="transaction_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <m-tx-list-page-table :items="txList"></m-tx-list-page-table>
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
	import Service from "../service";
	import Tools from "../util/Tools";
	import SpinComponent from './commonComponents/SpinComponent';
	import MTxListPageTable from "./table/MTxListPageTable";
	export default {
		name: "TransactionListPage",
		components: {MTxListPageTable, SpinComponent},
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
				let that = this, parmas;
				if(this.$route.params.txType === 'transfers'){
					this.listTitleName = "Txs (Transfer or Burn)";
					parmas = {txListTransfer: {pageNumber: currentPage,pageSize: pageSize}};
				}else if(this.$route.params.txType === 'delegations'){
					this.listTitleName = "Txs (Staking or Distribution)";
					parmas = {txListStake: {pageNumber: currentPage,pageSize: pageSize}};
				}else if(this.$route.params.txType === 'validations'){
					this.listTitleName = "Txs (CreateValidator, EditValidator or Unjail)";
					parmas = {txListDeclaration: {pageNumber: currentPage,pageSize: pageSize}};
				}else if(this.$route.params.txType === 'governance'){
					this.listTitleName = "Txs (SubmitProposal, Deposit or Vote)";
					parmas = {txListGov: {pageNumber: currentPage,pageSize: pageSize}};
				}
				this.flShowLoading = true;
				Service.commonInterface(parmas,(txList) => {
					try {
						this.count = txList.Count;
						this.totalPageNum =  Math.ceil((txList.Count/this.pageSize) === 0 ? 1 : (txList.Count/this.pageSize));
						sessionStorage.setItem('txpagenum',JSON.stringify(this.totalPageNum));
						if(txList.Data){
							this.txList = Tools.formatTxList(txList.Data,that.$route.params.txType)
						}else{
							this.txList = [];
							this.showNoData = true;
						}
						this.flShowLoading = false;
					}catch (e) {
						console.error(e)
					}
                })
			},
		}
	}
</script>

<style scoped lang="scss">
    .transaction_list_page_container{
        .transaction_list_title_wrap{
            width: 100%;
            position: fixed;
            z-index: 1;
            background-color: #ffffff;
            .transaction_list_title_content{
                height:0.7rem;
                display: flex;
                align-items: center;
                max-width: 12.8rem;
                margin: 0 auto;
                background-color: #ffffff;
                .transaction_list_title{
                    font-size: 0.18rem;
                    font-weight: 500;
                    padding-left: 0.2rem;
                }
            }
        }
    }
    .transaction_list_table_container{
        max-width: 12.8rem;
        padding-top: 0.7rem;
        margin: 0 auto;
        .transaction_list_table_content{
            .table_list_content{
                width: 100%;
                overflow-x: auto;
                padding-top: 0rem;
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
    @media screen and (max-width: 910px){
        .transaction_list_page_container{
            .transaction_list_title_wrap{
                position: static;
            }
        }

        .transaction_list_table_container{
            padding-top: 0;
            padding-left: 0.1rem;
            padding-right: 0.1rem;
            .transaction_list_table_content{
                .table_list_content{
                    padding-top: 0;
                }
            }
        }
    }
</style>
