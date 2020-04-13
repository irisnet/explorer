<template>
    <div class="transaction_list_page_container">
        <page-title :title="pageTitle" :content="`${count} Txs`"></page-title>
        <div class="transaction_list_title_wrap">
            <div class="transaction_list_title_content">
                <div class="filter_container">
                    <div class="filter_tx_type_statue_content">
                        <el-select v-model="value" filterable :change="filterTxByTxType(value)">
                            <el-option v-for="(item, index) in txTypeListArray"
                                       :key="index"
                                       :label="item.label"
                                       :value="item.value">
                            </el-option>
                        </el-select>

                        <el-select v-model="statusValue" :change="filterTxByStatus(statusValue)">
                            <el-option v-for="(item, index) in status"
                                       :key="index"
                                       :label="item.label"
                                       :value="item.value"></el-option>
                        </el-select>
                    </div>
                    <div class="select_date_content">
                        <el-date-picker  type="date"
                                         v-model="startTime"
                                         @change="getStartTime(startTime)"
                                         :picker-options="PickerOptions"
                                         :editable="false"
                                         value-format="yyyy-MM-dd"
                                         placeholder="Select Date">
                        </el-date-picker>
                        <span class="joint_mark">~</span>
                        <el-date-picker  type="date"
                                         v-model="endTime"
                                         :picker-options="PickerOptions"
                                         value-format="yyyy-MM-dd"
                                         @change="getEndTime(endTime)"
                                         :editable="false"
                                         placeholder="Select Date">
                        </el-date-picker>
                        <date-tooltip></date-tooltip>
                    </div>
                    <div class="reset_search_content">
                        <div class="tx_search_btn" @click="getFilterTxs">Search</div>
                        <div class="reset_btn" @click="resetFilterCondition"><i class="iconfont iconzhongzhi"></i></div>
                    </div>
                </div>
            </div>
        </div>
        <div class="transaction_list_table_container">
            <div class="transaction_list_table_content">
                <div class="table_list_content">
                    <m-tx-list-page-table :items="txList"></m-tx-list-page-table>
                    <div v-if="txList.length === 0" class="no_data_show">
                        <img src="../../assets/no_data.svg" alt="">
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
	import Service from "../../service";
	import Tools from "../../util/Tools";
	import MTxListPageTable from "./MTxListPageTable";
	import MPagination from "../commontables/MPagination";
    import DateTooltip from "../dateToolTip/DateTooltip";
    import PageTitle from "../pageTitle/PageTitle";
    import pageTitleConfig from "../pageTitle/pageTitleConfig";
	export default {
		name: "TransactionListPage",
		components: {PageTitle, DateTooltip, MPagination, MTxListPageTable},
		data() {
			return {
                pageTitle:'',
				totalPageNum: sessionStorage.getItem("txpagenum") ? JSON.parse(sessionStorage.getItem("txpagenum")) : 1,
				currentPageNum: this.forCurrentPageNum(),
                pickerStartTime:sessionStorage.getItem('firstBlockTime') ? sessionStorage.getItem('firstBlockTime') : '',
                PickerOptions: {
                    disabledDate: (time) => {
                        return time.getTime() < new Date(this.pickerStartTime).getTime() || time.getTime() > Date.now()
                    }
                },
				txList: [],
				txTypeListArray:[
					{
						value:'allTxType',
						label:'All TxType',
						slot:'allTXType'
					}
                ],
				listTitleName: "",
				count: sessionStorage.getItem("txTotal") ? Number(sessionStorage.getItem("txTotal")) : 0,
				pageSize: 30,
				showNoData: false,
				flShowLoading: false,
				value: this.getParamsByUrlHash().txType ? this.getParamsByUrlHash().txType : 'allTxType',
				txStatus: '',
				statusValue: this.getParamsByUrlHash().txStatus ? this.getParamsByUrlHash().txStatus : 'allStatus',
				status:[],
				startTime: this.getParamsByUrlHash().urlParamShowStartTime ? this.getParamsByUrlHash().urlParamShowStartTime : '',
				endTime: this.getParamsByUrlHash().urlParamShowEndTime ? this.getParamsByUrlHash().urlParamShowEndTime : '',
				filterStartTime: '',
				filterEndTime: '',
                urlParamsShowStartTime:this.getParamsByUrlHash().urlParamShowStartTime ? this.getParamsByUrlHash().urlParamShowStartTime : '',
                urlParamsShowEndTime:this.getParamsByUrlHash().urlParamShowEndTime ? this.getParamsByUrlHash().urlParamShowEndTime : '',
				type:'',
			}
		},
		mounted(){
                let statusArray = [
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
                    label:'Failed'
                }
            ]
            statusArray.forEach( item => {
                this.status.push(item)
            })

            this.getType();
			this.getTxListByFilterCondition();
		},
		methods: {
            getParamsByUrlHash(){
                let txType,
                    txStatus,
                    urlParamShowStartTime,
                    urlParamShowEndTime,
                    filterStartTime ,
                    filterEndTime ;
                let path = window.location.hash;
                if(path.includes("?")){
                    let urlHash = path.split('?')[1];
                    let params =  urlHash.split("&");
                    params.forEach( item => {
                        if(item.includes('txType')){
                            txType =  item.split("=")[1]
                        }else if (item.includes('status')){
                            txStatus = item.split("=")[1]
                        }else if(item.includes('startTime')){
                            urlParamShowStartTime = item.split("=")[1]
                            filterStartTime = this.formatStartTime(item.split("=")[1])
                        }else if(item.includes('endTime')){
                            urlParamShowEndTime = item.split("=")[1]
                            filterEndTime = this.formatEndTime(item.split("=")[1])
                        }
                    })
                }
                return {txType,txStatus,filterStartTime,filterEndTime,urlParamShowStartTime,urlParamShowEndTime}
            },
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
				let path = this.$route.path , urlParams = this.getParamsByUrlHash();
                this.statusValue = urlParams.txStatus ? urlParams.txStatus : 'allStatus';
                this.value = urlParams.txType ? urlParams.txType : 'allTxType';
                this.startTime = urlParams.urlParamShowStartTime ? urlParams.urlParamShowStartTime : '';
                this.endTime = urlParams.urlParamShowEndTime ? urlParams.urlParamShowEndTime : '';
                history.pushState(null, null, `/#${path}?txType=${urlParams.txType ? urlParams.txType : ''}&status=${urlParams.txStatus ? urlParams.txStatus : ''}&startTime=${urlParams.urlParamShowStartTime ? urlParams.urlParamShowStartTime : ''}&endTime=${urlParams.urlParamShowEndTime ? urlParams.urlParamShowEndTime : ''}&page=${pageNum}`);
				this.getTxListByFilterCondition();
			},
			getFilterTxs(){
				this.currentPageNum = 1;
				sessionStorage.setItem('txpagenum',1);
                history.pushState(null, null, `/#${this.$route.path}?txType=${this.TxType}&status=${this.txStatus}&startTime=${this.urlParamsShowStartTime}&endTime=${this.urlParamsShowEndTime}&page=1`);
                this.getTxListByFilterCondition();
                this.$uMeng.push('Transactions_Search','click')
            },
            resetUrl(){
                this.value = 'allTxType';
                this.statusValue = 'allStatus';
                this.startTime = '';
                this.endTime = '';
                this.currentPageNum = 1;
                this.urlParamsShowStartTime = '';
                this.urlParamsShowEndTime = '';
                history.pushState(null, null, `/#${this.$route.path}?txType=&status=&startTime=&endTime=&page=1`);
            },
			filterTxByTxType(e){
				if (e === 'allTxType' || e === undefined ) {
					this.TxType = '';
                    this.$uMeng.push('Transactions_All Type','click')
				}else {
					this.TxType = e
				}
			},
			filterTxByStatus(e){
				if(e === 'allStatus'){
					this.txStatus = '';
                    this.$uMeng.push('Transactions_All Status','click')
                    
				}else {
					this.txStatus = e
				}
			},
			getStartTime(time){
				this.filterStartTime = this.formatStartTime(time)
			},
			getEndTime(time){
				this.filterEndTime = this.formatEndTime(time)
			},
			formatStartTime(time){
                this.urlParamsShowStartTime =  time
				// let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
				return Number(new Date(time).getTime()/1000)
			},
			formatEndTime(time){
                this.urlParamsShowEndTime = time
				// let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
				let oneDaySeconds = 24 * 60 *60;
				return Number(new Date(time).getTime()/1000) + Number(oneDaySeconds)
			},
			resetFilterCondition(){
				this.getType();
                this.resetUrl()
                this.getTxListByFilterCondition();
                this.$uMeng.push('Transactions_Refresh','click')
			},
            getType(){
	            switch (this.$route.params.txType) {
		            case 'delegations' :
			            this.type = 'stake';
			            this.pageTitle = pageTitleConfig.StakingDelegationTxs;
			            break;
		            case 'validations':
			            this.type = 'declaration';
			            this.pageTitle = pageTitleConfig.StakingValidationTxs;
			            break;
		            case 'transfers':
			            this.type = 'trans';
			            this.pageTitle = pageTitleConfig.Transfer;
			            break;
		            case 'governance':
			            this.type = 'gov';
			            this.pageTitle = pageTitleConfig.GovGovTxs;
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
                            this.txTypeListArray = [
                                {
                                    value:'allTxType',
                                    label:'All TxType',
                                    slot:'allTXType'
                                }
                            ];
							this.txTypeListArray = this.txTypeListArray.concat(typeArray)
						}
					}catch (e) {
						console.error(e)
					}
				})
			},
			getTxListByFilterCondition(){
                let urlParams = this.getParamsByUrlHash(), param = {};
				param.getTxListByTypeAndTxType = {};
				param.getTxListByTypeAndTxType.type = this.type;
				param.getTxListByTypeAndTxType.pageNumber = this.currentPageNum;
				param.getTxListByTypeAndTxType.pageSize = this.pageSize;
				param.getTxListByTypeAndTxType.txType = urlParams.txType ? urlParams.txType: '';
				param.getTxListByTypeAndTxType.status = urlParams.txStatus ? urlParams.txStatus: '';
				param.getTxListByTypeAndTxType.beginTime = urlParams.filterStartTime ? urlParams.filterStartTime: '';
				param.getTxListByTypeAndTxType.endTime =  urlParams.filterEndTime ? urlParams.filterEndTime: '';
				Service.commonInterface(param, (txList) => {
					try {
                        this.count = txList.Count;
                        if(txList && txList.Data){
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
		},
    }
</script>

<style scoped lang="scss">
    .transaction_list_page_container{
        .transaction_list_title_wrap{
            width: 100%;
            position: fixed;
            z-index: 3;
            background-color: #F5F7FD;
            padding-top: 0.54rem;
            .transaction_list_title_content{
                height:0.7rem;
                display: flex;
                align-items: center;
                max-width: 12.8rem;
                margin: 0 auto;
                .filter_container{
                    display: flex;
                    .filter_tx_type_statue_content{
                        display: flex;
                        align-items: center;
                        /deep/.el-select{
                            width: 1.3rem;
                            margin-right: 0.1rem;
                            .el-input{
                                .el-input__inner{
                                    padding-left: 0.07rem;
                                    height: 0.32rem;
                                    font-size: 0.14rem !important;
                                    line-height: 0.32rem;
                                    &::-webkit-input-placeholder{
                                        font-size: 0.1rem !important;
                                    }
                                }
                                .el-input__inner:focus{
                                    border-color: var(--bgColor) !important;
                                }
                                .el-input__suffix{
                                   .el-input__suffix-inner{
                                       .el-input__icon{
                                           line-height: 0.32rem;
                                       }
                                   }
                                }
                            }
                            .is-focus{
                                .el-input__inner{
                                    border-color: var(--bgColor) !important;
                                }
                            }

                        }
                    }
                    .select_date_content{
                        display: flex;
                        align-items: center;
                        /deep/.el-date-editor{
                            width: 1.3rem;
                            .el-icon-circle-close{
                                display: none !important;
                            }
                            .el-input__inner{
                                height:0.32rem;
                                padding-left: 0.07rem;
                                padding-right: 0;
                                line-height: 0.32rem;
                                &::-webkit-input-placeholder{
                                    font-size: 0.14rem !important;
                                }
                                &:focus{
                                    border-color: var(--bgColor);
                                }
                            }
                            .el-input__prefix{
                                right: 5px;
                                left: 1rem;
                                .el-input__icon{
                                    line-height: 0.32rem;
                                }
                            }
                        }
                        .joint_mark{
                            margin:  0 0.08rem;
                        }
                    }
                    .reset_search_content{
                        display: flex;
                        align-items: center;
                        .reset_btn{
                            background: var(--bgColor);
                            color: #fff;
                            border-radius: 0.04rem;
                            margin-left: 0.1rem;
                            cursor: pointer;
                            height: 0.3rem;
                            i{
                                padding: 0.08rem;
                                font-size: 0.14rem;
                                line-height: 1;
                                display: inline-block;
                            }
                        }
                        .tx_search_btn{
                            height: 0.2rem;
                            cursor: pointer;
                            background: var(--bgColor);
                            margin-left: 0.1rem;
                            color: #fff;
                            border-radius: 0.04rem;
                            padding: 0.05rem 0.18rem;
                            font-size: 0.14rem;
                            line-height: 0.2rem;
                        }
                    }
                }

            }
        }
    }
    .transaction_list_table_container{
        max-width: 12.8rem;
        padding: 1.24rem 0 0.2rem 0;
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
                padding-top: 0.15rem;
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
                        .filter_tx_type_statue_content{
                            width: 3.45rem;
                            display: flex;
                            justify-content: space-between;
                            margin-bottom: 0.1rem;
                            .el-select{
                                margin-right: 0;
                                width: 1.6rem;
                            }
                        }
                        .select_date_content{
                            width: 3.45rem;
                            display: flex;
                            justify-content: space-between;
                            margin-bottom: 0.1rem;
                            .el-date-editor{
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
                            .tx_search_btn{
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
