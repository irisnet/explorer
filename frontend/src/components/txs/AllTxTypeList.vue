<template>
    <div>
        <div class="all_type_list_title_container">
            <div class="all_type_list_title_wrap">
                <div class="all_type_list_filter_content">
                    <h1 class="all_type_list_title">{{countNum}} Txs</h1>
                    <div class="filter_content">
                        <div class="tx_type_content">
                            <div class="tx_type_mobile_content">
                                <el-select v-model="value" filterable :change="filterTxByTxType(value)">
                                    <el-option v-for="(item, index) in txTypeOption"
                                               :key="index"
                                               :label="item.label"
                                               :value="item.value"></el-option>
                                </el-select>

                                <el-select v-model="statusValue" :change="filterTxByStatus(statusValue)">
                                    <el-option v-for="(item, index) in status"
                                               :key="index"
                                               :label="item.label"
                                               :value="item.value"></el-option>
                                </el-select>
                            </div>
                            <div class="tx_type_mobile_content">
                                <el-date-picker  type="date"
                                                 v-model="startTime"
                                                 @change="getStartTime(startTime)"
                                                 value-format="yyyy-MM-dd"
                                                 placeholder="Select Date">
                                </el-date-picker>
                                <span class="joint_mark">~</span>
                                <el-date-picker  type="date"
                                                 v-model="endTime"
                                                 value-format="yyyy-MM-dd"
                                                 @change="getEndTime(endTime)"
                                                 placeholder="Select Date">
                                </el-date-picker>
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
                    <img src="../../assets/no_data.svg" >
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
    import Service from "../../service"
    import Tools from "../../util/Tools"
    import MAllTxTypeListTable from "./MAllTxTypeListTable";
    import MPagination from "../commontables/MPagination";
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
		                label:'All TxType',
		                slot:'allTxType'
	                }
                ],
                status:[],
                statusValue: this.getParamsByUrlHash().txStatus ? this.getParamsByUrlHash().txStatus : 'allStatus',
				value: this.getParamsByUrlHash().txType  ? this.getParamsByUrlHash().txType : 'allTxType',
                firstEntry:false,
				startTime: this.getParamsByUrlHash().urlParamShowStartTime ? this.getParamsByUrlHash().urlParamShowStartTime : '',
                endTime:  this.getParamsByUrlHash().urlParamShowEndTime ? this.getParamsByUrlHash().urlParamShowEndTime : '',
                filterStartTime: '',
                filterEndTime: '',
                urlParamsShowStartTime:'',
                urlParamsShowEndTime:'',
                txStatus: '',
            }
        },
        mounted(){
            this.getTxListByFilterCondition();
	        this.getAllTxType();
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
        },
        methods:{
	        getFilterTxs(){
                this.currentPageNum = 1;
		        sessionStorage.setItem('txpagenum',1);
                history.pushState(null, null, `/#/txs?txType=${this.TxType}&status=${this.txStatus}&startTime=${this.urlParamsShowStartTime}&endTime=${this.urlParamsShowEndTime}&page=1`);
                this.getTxListByFilterCondition();
	        },
			filterTxByTxType(e){
				if (e === 'allTxType' || e === undefined ) {
					this.TxType = ''
                }else {
                    this.TxType = e
                }
            },
	        resetUrl(){
	            this.startTime = '';
	            this.endTime = '';
	            this.value = 'allTxType';
	            this.statusValue = 'allStatus';
	            this.txStatus ='';
	            this.urlParamsShowStartTime = '';
	            this.urlParamsShowEndTime = ''
                history.pushState(null, null, `/#/txs?txType=&status=&startTime=&endTime=&page=1`);
	        },
	        getStartTime(time){
	            this.urlParamsShowStartTime = time;
				this.filterStartTime = this.formatStartTime(time)
            },
	        getEndTime(time){
	            this.urlParamsShowEndTime = time;
		        this.filterEndTime = this.formatEndTime(time)
            },
            formatEndTime(time){

	            // let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
	            let oneDaySeconds = 24 * 60 *60;
	            return Number(new Date(time).getTime()/1000) + Number(oneDaySeconds)
            },
	        formatStartTime(time){
		        // let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
		        return Number(new Date(time).getTime()/1000)
            },
	        filterTxByStatus(e){
		        if(e === 'allStatus' || e === undefined ){
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
                    let urlParams = this.getParamsByUrlHash();
                    this.statusValue = urlParams.txStatus ? urlParams.txStatus : 'allStatus';
                    this.value = urlParams.txType ? urlParams.txType : 'allTxType';
                    this.startTime = urlParams.urlParamShowStartTime ? urlParams.urlParamShowStartTime : '';
                    this.endTime = urlParams.urlParamShowEndTime ? urlParams.urlParamShowEndTime : '';
			        history.pushState(null, null, `/#/txs?txType=${urlParams.txType ? urlParams.txType : ''}&status=${urlParams.txStatus ? urlParams.txStatus : ''}&startTime=${urlParams.urlParamShowStartTime ? urlParams.urlParamShowStartTime : ''}&endTime=${urlParams.urlParamShowEndTime ? urlParams.urlParamShowEndTime : ''}&page=${pageNum}`);
			        this.getTxListByFilterCondition();
	        },
            formatFee(Fee){
	            if(Fee.amount && Fee.denom){
		            return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(Fee.amount)),4)} ${Tools.formatDenom(Fee.denom).toUpperCase()}`;
	            }
            },
            getParamsByUrlHash(){
                let txType,
                    txStatus,
                    filterStartTime ,
                    urlParamShowStartTime,
                    urlParamShowEndTime,
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
	        getTxListByFilterCondition(){
               let param = {},urlParams = this.getParamsByUrlHash();
                param.getTxListByFilterCondition = {};
		        param.getTxListByFilterCondition.pageNumber = this.currentPageNum;
		        param.getTxListByFilterCondition.pageSize = this.pageSize;
		        param.getTxListByFilterCondition.txType = urlParams.txType ? urlParams.txType: '';
		        param.getTxListByFilterCondition.status = urlParams.txStatus ? urlParams.txStatus: '';
		        param.getTxListByFilterCondition.beginTime = urlParams.filterStartTime ? urlParams.filterStartTime: '';
		        param.getTxListByFilterCondition.endTime = urlParams.filterEndTime ? urlParams.filterEndTime: '';
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
					                status: Tools.firstWordUpperCase(item.status),
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
        box-sizing: border-box;
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
                    color: #515a6e;
                }
                .filter_content{
                    margin-left: 0.1rem;
                    .tx_type_content{
                        display: flex;
                        align-items: center;
                        .tx_type_mobile_content{
                            display: flex;
                            align-items: center;

                            /deep/.el-select{
                                width: 1.3rem;
                                margin-right: 0.1rem;
                                .el-input{
                                    .el-input__inner{
                                        padding-left: 0.07rem;
                                        height: 0.32rem;
                                        font-size: 0.1rem !important;
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
                            /deep/.el-date-editor{
                                width: 1.3rem;
                                .el-icon-circle-close{
                                    display: none !important;
                                }
                                .el-input__inner{
                                    height:0.32rem;
                                    padding-left: 0.07rem;
                                    padding-right: 0;
                                    &::-webkit-input-placeholder{
                                        font-size: 0.1rem !important;
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
                                margin: 0 0.1rem;
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
                                line-height: 0.2rem;
                            }
                        }
                    }
                }
            }
        }
    }
    .all_type_list_table_container{
        padding: 0.7rem 0 0.01rem 0;
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
            margin: 0.2rem auto 0.4rem auto;
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
                                .el-select{
                                    margin-right: 0;
                                    width: 1.6rem;
                                }
                                .el-date-editor{
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
            margin: 0 0.1rem;
        }
        .mobile_style{
            display: none;
        }
    }
</style>
