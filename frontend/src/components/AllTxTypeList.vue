<template>
    <div>
        <div class="all_type_list_title_container">
            <div class="all_type_list_title_wrap">
                <div class="all_type_list_filter_content">
                    <h1 class="all_type_list_title">{{countNum}} txs</h1>
                    <div class="filter_content">
                        <div class="tx_type_content">
                            <div class="tx_type_mobile_content">
                                <i-select :model.sync="value" v-model="value" :on-change="filterTxByTxType(value)">
                                    <i-option v-for="item in txTypeOption"
                                              :value="item.value"
                                    >{{item.label}}</i-option>
                                </i-select>
                                <i-select :model.sync="statusValue" v-model="statusValue" :on-change="filterTxByStatus(statusValue)">
                                    <i-option v-for="item in status"
                                              :value="item.value"
                                    >{{item.label}}</i-option>
                                </i-select>
                            </div>
                            <div class="tx_type_mobile_content">
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
                            <div class="tx_type_mobile_content">
                                <div class="search_btn" @click="getFilterTxs">Search</div>
                                <div class="reset_btn" @click="resetFilterCondition"><i class="iconfont iconzhongzhi"></i></div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="pagination_content mobile_style">
                    <m-pagination
                            :page-size="pageSize"
                            :total="countNum"
                            :page="currentPageNum"
                            :page-change="pageChange"
                    ></m-pagination>
                </div>
            </div>
        </div>

        <div class="all_type_list_table_container">
            <div class="all_type_list_table_wrap">

                <m-all-tx-type-list-table :items="allTxTypeList"></m-all-tx-type-list-table>
                <div class="no_data_img_content" v-if="allTxTypeList.length === 0">
                    <img src="../assets/no_data.svg" >
                </div>
            </div>

            <div class="pagination_content">
                <keep-alive>
                    <m-pagination
                            :page-size="pageSize"
                            :total="countNum"
                            :page="currentPageNum"
                            :page-change="pageChange"
                    ></m-pagination>
                </keep-alive>
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
                txTypeOption:[
	                {
		                value:'allTxType',
		                label:'All TXType',
		                slot:'allTXType'
	                }
                ],
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
                statusValue:'allStatus',
				value: 'allTxType',
                firstEntry:false,
				startTime: '',
                endTime: '',
                filterStartTime: '',
                filterEndTime: '',
                TxType: '',
                txStatus: '',
            }
        },
        created(){
            this.getTxListByFilterCondition();
	        this.getAllTxType();
        },
        methods:{
	        getFilterTxs(){
		        this.currentPageNum = 1;
		        this.resetUrl();
		        sessionStorage.setItem('txpagenum',1);
		        this.getTxListByFilterCondition();
	        },
			filterTxByTxType(e){
				if (e === 'allTxType') {
					this.TxType = ''
                }else {
					this.TxType = e
                }
            },
	        resetUrl(){
	        	if(this.$route.query.page){
			        this.$router.push({
				        path: this.$route.path,
				        query:{
					        page:1
				        }
			        });
                }
	        },
	        getStartTime(time){
				this.filterStartTime = this.formatStartTime(time)
            },
	        getEndTime(time){
		        this.filterEndTime = this.formatTime(time)
            },
            formatTime(time){
	            let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
	            let oneDaySeconds = 24 * 60 *60;
	            return Number(new Date(utcTime).getTime()/1000) + Number(oneDaySeconds)
            },
	        formatStartTime(time){
		        let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
		        return Number(new Date(utcTime).getTime()/1000)
            },
	        filterTxByStatus(e){
		        if(e === 'allStatus'){
			        this.txStatus = ''
		        }else {
			        this.txStatus = e
		        }
	        },
            getAllTxType(){
			    Service.commonInterface({allTxType:{
			    	    type: 'all'
                    }},(res) => {
			    	try {
                        if(res){
                        	let txType;
	                        txType = res.map(item => {
	                        	return {
			                        value : item,
                                    label : item
                                }
                            });
	                        this.txTypeOption = this.txTypeOption.concat(txType);
                        }
				    }catch (e) {
                        console.error(e)
				    }
                })
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
                this.resetUrl();
                this.getTxListByFilterCondition()
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
			        let path = "/txs";
			        history.pushState(null, null, `/#${path}?page=${pageNum}`);
			        this.getTxListByFilterCondition();
	        },
            formatFee(Fee){
	            if(Fee.amount && Fee.denom){
		            return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(Fee.amount)),4)} ${Tools.formatDenom(Fee.denom).toUpperCase()}`;
	            }
            },
	        getTxListByFilterCondition(){
                let param = {};
                param.getTxListByFilterCondition = {};
		        param.getTxListByFilterCondition.pageNumber = this.currentPageNum;
		        param.getTxListByFilterCondition.pageSize = this.pageSize;
		        param.getTxListByFilterCondition.txType = this.TxType;
		        param.getTxListByFilterCondition.status = this.txStatus;
		        param.getTxListByFilterCondition.beginTime = this.filterStartTime;
		        param.getTxListByFilterCondition.endTime = this.filterEndTime;
                Service.commonInterface(param, (res) => {
                	try {
		                this.countNum = res.Count;
		                if(res && res.Data) {
			                sessionStorage.setItem('txsTotal',res.Count);
			                this.allTxTypeList = res.Data.map( item => {
				                return {
					                txHash:item.hash,
					                block: item.block_height,
					                type: item.type,
					                fee: this.formatFee(item.fee),
					                signer: item.signer,
					                status: item.status,
					                timestamp: Tools.format2UTC(item.timestamp)
				                }
			                })
                        }else {
			                this.allTxTypeList = []

                        }
                    }catch (e) {
                		console.error(e)
	                }
                })
            },
        },
		watch: {
			$route(newVal) {
				// 有时候 mounted 方法不起作用，为此添加该 watch
				this.currentPageNum = Number(this.$route.query.page || 1);
				this.getTxListByFilterCondition();
			},
		},
	}
</script>

<style scoped lang="scss">
    .all_type_list_title_container{
        width: 100%;
        position: fixed;
        z-index: 3;
        background: #F5F7FD;
        .all_type_list_title_wrap{
            max-width: 12.8rem;
            margin: 0 auto;
            top:0;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .all_type_list_filter_content{
                display: flex;
                align-items: center;
                .all_type_list_title{
                    height: 0.7rem;
                    margin: 0;
                    line-height: 0.7rem;
                    font-size: 0.18rem;
                    padding-left: 0.2rem;
                }
                .filter_content{
                    margin-left: 0.1rem;
                    .tx_type_content{
                        display: flex;
                        align-items: center;
                        .tx_type_mobile_content{
                            display: flex;
                            align-items: center;
                            .ivu-select-visible{
                                /deep/ .ivu-select-selection{
                                    border-color: var(--bgColor) !important;
                                }

                            }
                            .ivu-select{
                                margin-right: 0.1rem;
                                width: 1.3rem;
                                /deep/ .ivu-select-selection:hover{
                                    border-color: var(--bgColor) !important;
                                }
                                .ivu-select-item{
                                    text-indent: 0.1rem;
                                    font-size: 0.14rem;
                                    line-height: 0.18rem;
                                    padding: 0.07rem 0.1rem 0.07rem 0;
                                    color: var(--bgColor);
                                }
                            }
                            .joint_mark{
                                margin: 0 0.1rem;
                            }
                            .ivu-date-picker{
                                width: 1.3rem;
                                /deep/ .ivu-date-picker-rel{
                                    .ivu-input-wrapper{
                                        .ivu-input:hover{
                                            border-color: var(--bgColor) !important;
                                        }
                                        .ivu-input:focus{
                                             border-color: var(--bgColor) !important;
                                         }
                                    }
                                }
                            }
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
    }
    .all_type_list_table_container{
        padding-top: 0.7rem;
        .all_type_list_table_wrap{
            max-width: 12.8rem;
            margin: 0 auto;
            overflow-x: auto;
            .no_data_img_content{
                display: flex;
                justify-content: center;
                border-top: 0.01rem solid #eee;
                border-bottom: 0.01rem solid #eee;
                font-size: 0.14rem;
                height: 2.8rem;
                align-items: center;
            }
        }
        .pagination_content{
            max-width: 12.8rem;
            display: flex;
            margin: 0.2rem auto 0.2rem auto;
            justify-content:flex-end;
        }
    }
    .el-select-dropdown__item{
        padding-left: 0.15rem;
    }
    @media screen and (max-width: 910px){
        .all_type_list_title_container{
            position: static;
            padding-left: 0.1rem;
            .all_type_list_title_wrap{
                .all_type_list_filter_content{
                    flex-direction: column;
                    align-items: flex-start;
                    width: 100%;
                    .filter_content{
                        width: 100%;
                        margin-left: 0;
                        display: flex;
                        .tx_type_content{
                            width: 100%;
                            display: flex;
                            flex-direction: column;
                            align-items: flex-start;
                            .tx_type_mobile_content{
                                width: 3.45rem;
                                display: flex;
                                justify-content: space-between;
                                margin-bottom: 0.1rem;
                                .ivu-select{
                                    margin-right: 0;
                                    width: 1.6rem;
                                }
                                .ivu-select-visible{
                                    .ivu-select-selection{
                                        border-color: var(--bgColor);
                                    }
                                }
                                .ivu-date-picker{
                                    width: 1.6rem;
                                }
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
        }
        .all_type_list_table_container{
            padding-top: 0;
            padding-left: 0.1rem;
        }
        .pagination_content{
            padding-right: 0.1rem;
        }
        .mobile_style{
            display: none;
        }
    }
</style>
