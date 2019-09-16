<template>
    <div class="transaction_list_page_container">
        <div class="transaction_list_title_wrap">
            <div class="transaction_list_title_content">
                <span class="transaction_list_title">{{count}} Txs</span>
                <div class="filter_container">
                    <div class="filter_tx_type_statue_content">
                        <i-select :model.sync="value" v-model="value" :on-change="filterTxByTxType(value)">
                            <i-option v-for="item in txTypeListArray" :value="item.value">{{item.label}}</i-option>
                        </i-select>
                        <i-select :model.sync="statusValue" v-model="statusValue" :on-change="filterTxByStatus(statusValue)">
                            <i-option v-for="item in status"
                                      :value="item.value"
                            >{{item.label}}</i-option>
                        </i-select>
                    </div>
                    <div class="select_date_content">
                        <Date-picker type="date"
                                     v-model="startTime"
                                     :value.sync="startTime"
                                     :clearable = false
                                     @on-change='getStartTime'
                                     placeholder="Select Date"
                        ></Date-picker>
                        <span class="joint_mark">~</span>
                        <Date-picker type="date"
                                     v-model="endTime"
                                     :value.sync="endTime"
                                     :clearable = false
                                     @on-change="getEndTime"
                                     placeholder="Select Date"
                        ></Date-picker>
                    </div>
                    <div class="reset_search_content">
                        <div class="search_btn" @click="getFilterTxs">Search</div>
                        <div class="reset_btn" @click="resetFilterCondition"><i class="iconfont iconzhongzhi"></i></div>
                    </div>
                </div>
            </div>
        </div>
        <div class="transaction_list_table_container">
            <div class="transaction_list_table_content">
                <div class="table_list_content">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <m-tx-list-page-table :items="txList"></m-tx-list-page-table>
                    <div v-if="txList.length === 0" class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
                <div class="pagination_nav_footer_content">
                    <keep-alive>
                        <m-pagination :total="count" :page="currentPageNum" :page-size="pageSize" :page-change="pageChange"></m-pagination>
                    </keep-alive>
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
	import MPagination from "./commonComponents/MPagination";
	export default {
		name: "TransactionListPage",
		components: {MPagination, MTxListPageTable, SpinComponent},
		data() {
			return {
				totalPageNum: sessionStorage.getItem("txpagenum") ? JSON.parse(sessionStorage.getItem("txpagenum")) : 1,
				currentPageNum: this.forCurrentPageNum(),
				txList: [],
				txTypeListArray:[
					{
						value:'allTxType',
						label:'All TXType',
						slot:'allTXType'
					}
                ],
				listTitleName: "",
				count: sessionStorage.getItem("txTotal") ? Number(sessionStorage.getItem("txTotal")) : 0,
				pageSize: 30,
				showNoData: false,
				flShowLoading: false,
				value: 'allTxType',
				txStatus: '',
				statusValue:'allStatus',
				status:[
					{
						value:'allStatus',
						label:'All Status'
					},
					{
						value:'success',
						label:'Success'
					},
					{
						value:'fail',
						label:'Fail'
					}
				],
				startTime: '',
				endTime: '',
				filterStartTime: '',
				filterEndTime: '',
				type:'',
				TxType:''
			}
		},
		created(){
			this.getType();
			this.getTxListByFilterCondition();
		},
		methods: {
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
				let path = "/txs/transfers";
				history.pushState(null, null, `/#${path}?page=${pageNum}`);
				this.getTxListByFilterCondition();
			},
			getFilterTxs(){
				this.currentPageNum = 1;
				this.resetUrl();
				sessionStorage.setItem('txpagenum',1);
				this.getTxListByFilterCondition();
            },
            resetUrl(){
				if(this.$route.query.page)
	            this.$router.push({
		            path: this.$route.path,
		            query:{
			            page:1
		            }
	            });
            },
			filterTxByTxType(e){
				if (e === 'allTxType') {
					this.TxType = ''
				}else {
					this.TxType = e
				}
			},
			filterTxByStatus(e){
				if(e === 'allStatus'){
					this.txStatus = ''
				}else {
					this.txStatus = e
				}
			},
			getStartTime(time){
				this.filterStartTime = this.formatStartTime(time)
			},
			getEndTime(time){
				this.filterEndTime = this.formatTime(time)
			},
			formatStartTime(time){
				let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
				return Number(new Date(utcTime).getTime()/1000)
			},
			formatTime(time){
				let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
				let oneDaySeconds = 24 * 60 *60;
				return Number(new Date(utcTime).getTime()/1000) + Number(oneDaySeconds)
			},
			resetFilterCondition(){
				this.value = 'allTxType';
				this.statusValue = 'allStatus';
				this.startTime = '';
				this.endTime = '';
				this.filterStartTime= '';
				this.filterEndTime = '';
				this.TxType = '';
				this.txStatus = '';
				this.currentPageNum = 1;
				this.getType();
				this.getTxListByFilterCondition();
				this.resetUrl()
			},
			linkGen(pageNum) {
                return pageNum === 1 ? '?' : `?page=${pageNum}`
            },
            getType(){
	            switch (this.$route.params.txType) {
		            case 'delegations' :
			            this.type = 'stake';
			            break;
		            case 'validations':
			            this.type = 'declaration';
			            break;
		            case 'transfers':
			            this.type = 'trans';
			            break;
		            case 'governance':
			            this.type = 'gov';
	            }
	            this.getAllTxType();
            },
			getAllTxType(){
				Service.commonInterface({allTxType:{
						type: this.type
					}},(res) => {
					try {
						if(res){
							let typeArray;
							typeArray = res.map(item => {
								return {
									value : item,
									label : item
								}
							});
							this.txTypeListArray = this.txTypeListArray.concat(typeArray)
						}
					}catch (e) {
						console.error(e)
					}
				})
			},
			getTxListByFilterCondition(){

				this.flShowLoading = true;
				let param = {};
				param.getTxListByTypeAndTxType = {};
				param.getTxListByTypeAndTxType.type = this.type;
				param.getTxListByTypeAndTxType.pageNumber = this.currentPageNum;
				param.getTxListByTypeAndTxType.pageSize = this.pageSize;
				param.getTxListByTypeAndTxType.txType = this.TxType;
				param.getTxListByTypeAndTxType.status = this.txStatus;
				param.getTxListByTypeAndTxType.beginTime = this.filterStartTime;
				param.getTxListByTypeAndTxType.endTime = this.filterEndTime;
				Service.commonInterface(param, (txList) => {
					try {
						if(txList && txList.Data){
							this.count = txList.Count;
							sessionStorage.setItem('txTotal',txList.Count);
							this.totalPageNum =  Math.ceil((txList.Count/this.pageSize) === 0 ? 1 : (txList.Count/this.pageSize));
							sessionStorage.setItem('txpagenum',JSON.stringify(this.totalPageNum));
							if(txList.Data){
								this.txList = Tools.formatTxList(txList.Data,this.$route.params.txType)
							}else{
								this.txList = [];
								this.showNoData = true;
							}
							this.flShowLoading = false;
                        }else {
							this.txList = [];
							this.showNoData = true;
							this.flShowLoading = false;
                        }
					}catch (e) {
						this.showNoData = true;
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
            z-index: 3;
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
                .filter_container{
                    display: flex;
                    margin-left: 0.1rem;
                    .filter_tx_type_statue_content{
                        display: flex;
                        align-items: center;
                        .ivu-select{
                            width: 1.3rem;
                            margin-right: 0.1rem;
                            .ivu-select-item{
                                text-indent: 0.1rem;
                                font-size: 0.14rem;
                                line-height: 0.18rem;
                                padding: 0.07rem 0.1rem 0.07rem 0;
                            }
                        }
                    }
                    .select_date_content{
                        display: flex;
                        align-items: center;
                        .ivu-date-picker{
                            width: 1.3rem;
                        }
                        .joint_mark{
                            margin:  0 0.1rem;
                        }
                    }
                    .reset_search_content{
                        display: flex;
                        .reset_btn{
                            background: var(--bgColor);
                            color: #fff;
                            border-radius: 0.04rem;
                            margin-left: 0.1rem;
                            cursor: pointer;
                            i{
                                padding: 0.08rem;
                                font-size: 0.14rem;
                                line-height: 1;
                                display: inline-block;
                            }
                        }
                        .search_btn{
                            cursor: pointer;
                            background: var(--bgColor);
                            margin-left: 0.1rem;
                            color: #fff;
                            border-radius: 0.04rem;
                            padding: 0.05rem 0.18rem;
                            font-size: 0.14rem;
                        }
                    }
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
            }

        }
    }
    @media screen and (max-width: 910px){
        .transaction_list_page_container{
            .transaction_list_title_wrap{
                position: static;
                .transaction_list_title_content{
                    display: flex;
                    flex-direction: column;
                    height: auto;
                    align-items: flex-start;
                    .transaction_list_title{
                        height: 0.7rem;
                        line-height: 0.7rem;
                    }
                    .filter_container{
                        flex-direction: column;
                        margin-left: 0.1rem;
                        width: 100%;
                        .filter_tx_type_statue_content{
                            width: 3.45rem;
                            display: flex;
                            justify-content: space-between;
                            margin-bottom: 0.1rem;
                            .ivu-select{
                                margin-right: 0;
                                width: 1.6rem;
                            }
                        }
                        .select_date_content{
                            width: 3.45rem;
                            display: flex;
                            justify-content: space-between;
                            margin-bottom: 0.1rem;
                            .ivu-date-picker{
                                width: 1.6rem;
                            }
                        }
                        .reset_search_content{
                            width: 3.45rem;
                            display: flex;
                            justify-content: space-between;
                            margin-bottom: 0.1rem;
                            .reset_btn{
                                margin-left: 0;
                            }
                            .search_btn{
                                flex: 1;
                                margin-left: 0;
                                margin-right: 0.1rem;
                                text-align: center;
                            }
                        }
                    }
                }
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
