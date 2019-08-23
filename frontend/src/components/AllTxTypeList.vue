<template>
    <div>
        <div class="all_type_list_title_container">
            <div class="all_type_list_title_wrap">
                <h1 class="all_type_list_title">Transactions</h1>
                <div class="pagination_content">
                    <m-pagination
                            :page-size="pageSize"
                            :total="countNum"
                            :page="currentPageNum"
                            :showInfo="false"
                            :page-change="pageChange"
                    ></m-pagination>
                </div>
            </div>
        </div>

        <div class="all_type_list_table_container">
            <div class="all_type_list_table_wrap">
                <m-all-tx-type-list-table :items="allTxTypeList"></m-all-tx-type-list-table>
            </div>
            <div class="pagination_content">
                <m-pagination
                        :page-size="pageSize"
                        :total="countNum"
                        :page="currentPageNum"
                        :showInfo="false"
                        :page-change="pageChange"
                ></m-pagination>
            </div>
        </div>

    </div>
</template>

<script>
    import Service from "../service"
    import Tools from "../util/Tools"
    import MAllTxTypeListTable from "./table/MAllTxTypeListTable";
    import MPagination from "./commonComponents/MPagination";
	export default {
		name: "AllTxTypeList",
		components: {MAllTxTypeListTable,MPagination},
		data() {
			return {
				allTxTypeList: [],
                pageSize: 30,
				countNum: sessionStorage.getItem("txsTotal") ? Number(sessionStorage.getItem("txsTotal")) : 0,
				currentPageNum: this.forCurrentPageNum(),
				currentPageNumCache: 0,
            }
        },
        created(){
	        this.getAllTxTypeList();
        },
        methods:{
	        forCurrentPageNum() {
		        let currentPageNum = 1;
		        let urlPageSize = this.$route.query.page && Number(this.$route.query.page);
		        currentPageNum = urlPageSize ? urlPageSize : 1;
		        return currentPageNum;
	        },
	        pageChange(pageNum) {
		        this.currentPageNum = pageNum;
		        if (this.currentPageNumCache === this.currentPageNum) {
			        return;
		        }
		        this.currentPageNumCache = this.currentPageNum;
			        let path = "/txs";
			        history.pushState(null, null, `/#${path}?page=${pageNum}`);
			        this.getAllTxTypeList();
	        },
            formatFee(Fee){
	            if(Fee.amount && Fee.denom){
		            return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(Fee.amount)),4)} ${Tools.formatDenom(Fee.denom).toUpperCase()}`;
	            }
            },
            getAllTxTypeList(){
	            Service.commonInterface({allTypeList:{
			            pageNumber:this.currentPageNum,
			            pageSize: this.pageSize,
		            }},(res) => {
		            try {
			            if(res){
				            this.countNum = res.Count;
				            sessionStorage.setItem('txsTotal',res.Count);
				            this.allTxTypeList = res.Data.map( item => {
					            return {
						            txHash:item.Hash,
						            block: item.BlockHeight,
						            type: item.Type,
						            fee: this.formatFee(item.Fee),
						            signer: item.Signer,
						            status: item.Status,
						            timestamp: Tools.format2UTC(item.Timestamp)
					            }
				            })

			            }
		            }catch (e) {
			            console.error(e)
		            }
	            })
            }
        },
		watch: {
			$route(newVal) {
				// 有时候 mounted 方法不起作用，为此添加该 watch
				this.currentPageNum = Number(this.$route.query.page || 1);
				this.getAllTxTypeList();
			}
		},
	}
</script>

<style scoped lang="scss">
    .all_type_list_title_container{
        width: 100%;
        position: fixed;
        z-index: 1;
        background: #fff;
        .all_type_list_title_wrap{
            max-width: 12.8rem;
            margin: 0 auto;
            top:0;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .all_type_list_title{
                height: 0.7rem;
                margin: 0;
                line-height: 0.7rem;
                font-size: 0.18rem;
            }
        }
    }
    .all_type_list_table_container{
        padding-top: 0.7rem;
        .all_type_list_table_wrap{
            max-width: 12.8rem;
            margin: 0 auto;
            overflow-x: auto;
        }
        .pagination_content{
            max-width: 12.8rem;
            margin: 0.2rem auto 0.2rem auto;
            text-align: right;
        }
    }
    @media screen and (max-width: 910px){
        .all_type_list_title_container{
            position: static;
            padding-left: 0.1rem;
        }
        .all_type_list_table_container{
            padding-top: 0;
            padding-left: 0.1rem;
        }
        .pagination_content{
            padding-right: 0.1rem;
        }
    }
    .pagination_content{
        .info{
            display: none !important;
        }
    }

</style>
