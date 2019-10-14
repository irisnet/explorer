<template>
    <div class="address_information_container">
        <div class="address_information_content">
            <div class="address_information_header_container">
                <div class="address_information_header_content">
                    <span class="address_information_header">{{headerAddress}}</span>
                </div>
            </div>
            <div class="address_information_assets_container">
                <div class="address_information_assets_title">Assets</div>
                <div class="address_information_assets_list_content">
                    <m-address-information-table
                            :items="assetsItems"
                            listName="assets">
                    </m-address-information-table>
                </div>
            </div>
            <div class="address_information_delegation_tx_container">
                <div class="address_information_delegation_tx_content">
                    <div class="address_information_delegation_title">Delegations</div>
                    <div class="address_information_delegation_list_content">
                        <m-address-information-table
                                :items="delegationsItems"
                                listName="delegations"
                                :width="630">

                        </m-address-information-table>
                    </div>
                </div>
                <div class="address_information_unbonding_delegation_tx_content">
                    <div class="address_information_unbonding_delegation_title">Unbonding Delegations</div>
                    <div class="address_information_unbonding_delegation_list_content">
                        <m-address-information-table
                                :items="unBondingDelegationsItems"
                                listName="unBondingDelegations"
                                :width="630">

                        </m-address-information-table>
                    </div>
                </div>
            </div>
            <div class="address_information_redelegation_header_title">Delegator Rewards</div>
            <div class="address_information_redelegation_tx_container">
                <div class="address_information_delegator_rewards_content">
                    <div class="address_information_detail_option">
                        <span class="address_information_detail_option_name">Withdraw To:</span>
                        <span class="address_information_detail_option_value">
                            <router-link :to="`/address/${withdrewToAddress}`">{{withdrewToAddress}}</router-link></span>
                    </div>
                    <div class="address_information_list_content">
                        <m-address-information-table
                                :items="rewardsItems"
                                listName="rewards"
                                :width="630">

                        </m-address-information-table>
                    </div>
                </div>
                <div class="address_information_detail_container">
                    <div class="address_information_redelegation_title">Validator Rewards</div>
                    <ul class="address_information_detail_content">
                        <li class="address_information_detail_option">
                            <span class="address_information_detail_option_name">Validator Moniker:</span>
                            <span class="address_information_detail_option_value">{{validatorMoniker}}</span>
                            <span class="address_information_address_status_active" v-if="validatorStatus === 'active'">Active</span>
                            <span class="address_information_address_status_candidate" v-if="validatorStatus === 'candidate'">Candidate</span>
                            <span class="address_information_address_status_jailed" v-if="validatorStatus === 'jailed'">Jailed</span>
                            <span class="address_information_address_status_jailed" v-if="isProfiler">
                                 <img v-show="flShowProfileLogo"
                                      class="profile_logo_img"
                                      src="../assets/profiler_logo.png">
                            </span>
                        </li>
                        <li class="address_information_detail_option">
                            <span class="address_information_detail_option_name">Operator Address:</span>
                            <span class="address_information_detail_option_value">
                                <router-link v-show="OperatorAddress !== '--'" :to="`/validators/${OperatorAddress}`">{{OperatorAddress}}</router-link>
                                <span v-show="OperatorAddress === '--'" >{{OperatorAddress}}</span>
                            </span>
                        </li>
                    </ul>
                </div>
            </div>
            <div class="address_information_transaction_container">
                <div class="address_information_transaction_header_content">
                    <span class="address_information_transaction_title">Transactions</span>
                    <div class="address_information_list_filter_content">
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
                </div>
                <div class="address_information_list_content">
                    <m-address-information-table
                            :items="transactionsItems"
                            listName="transactions">

                    </m-address-information-table>
                </div>
                <!--<div class="pagination_content">
                    <keep-alive>
                        <m-pagination
                                :page-size="pageSize"
                                :total="countNum"
                                :page="currentPageNum"
                                :page-change="pageChange"
                        ></m-pagination>
                    </keep-alive>
                </div>-->
            </div>
        </div>
    </div>
</template>

<script>
	import MAddressInformationTable from "./table/MAddressInformationTable";
	import Tools from "../util/Tools"
	import Server from '../service'
	import Constant from "../constant/Constant";
	export default {
		name: "AddressInfomation",
		components: {MAddressInformationTable},
		data(){
			return {
				withdrewToAddress: '',
				validatorMoniker: '',
                validatorStatus:'',
				OperatorAddress: '',
				isProfiler:'',
				txTypeOption:[
					{
						value:'allTxType',
						label:'All TXType',
						slot:'allTXType'
					}
				],
				statusValue:'allStatus',
				startTime: '',
				endTime: '',
				filterStartTime: '',
				filterEndTime: '',
				TxType: '',
				txStatus: '',
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
				value: 'allTxType',
                assetsItems:[],
                delegationsItems:[],
                unBondingDelegationsItems:[],
                rewardsItems:[],
                transactionsItems:[],
				headerAddress:'',
                totalDelegatorReward: 0,
                totalUnBondingDelegator:0,
                totalDelegator:0,
                assetList: [],
				fixedNumber: 2
            }
        },
        watch:{
			$router(){
				this.getTxListByFilterCondition();
            },
	        totalDelegatorReward(totalDelegatorReward){
				this.getAssetList()
            },
	        totalUnBondingDelegator(totalDelegatorReward){
		        this.getAssetList()
            },
	        totalDelegator(){
		        this.getAssetList()
            }
        },
        mounted(){
			this.headerAddress = this.$route.params.param;
            this.getAddressInformation();
            this.getAssetList();
            this.getDelegationList();
            this.getUnBondingDelegationList();
            this.getRewardsItems();
	        this.getTxListByFilterCondition();
            this.getAllTxType();
        },
        methods:{
	        getFilterTxs(){
		        this.currentPageNum = 1;
		        this.resetUrl();
		        sessionStorage.setItem('addressTxPageNum',1);
		        this.getTxListByFilterCondition();
	        },
			getAddressInformation(){
                Server.commonInterface({addressInformation:{
		                address:this.$route.params.param
                    }},(res) => {
                	try {
                        if(res){
	                        this.validatorMoniker = res.moniker ? res.moniker : '--';
	                        this.OperatorAddress = res.operator_address ? res.operator_address : '--';
	                        this.validatorStatus = res.status;
	                        this.withdrewToAddress = res.withdrawAddress ? res.withdrawAddress : '--';
	                        this.isProfiler = res.isProfiler;
	                        this.assetList = res.amount;
	                        this.getAssetList()
                        }
	                }catch (e) {
                        console.error(e)
	                }
                })
            },
            getAssetList(){
	            this.assetsItems = this.assetList.map( item => {
	            	if(item.denom === 'iris-atto'){
			            return {
				            token: Tools.formatDenom(item.denom),
				            balance: item.amount ? Tools.formatAmount2(item,this.fixedNumber): 0,
				            delegatedValue: this.totalDelegator ? this.totalDelegator : 0,
				            delegated: this.totalDelegator ? `${Tools.formatStringToFixedNumber(this.totalDelegator.toString(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`: 0,
				            unBondingValue: this.totalUnBondingDelegator ? this.totalUnBondingDelegator : 0,
				            unBonding: this.totalUnBondingDelegator ?`${Tools.formatStringToFixedNumber(this.totalUnBondingDelegator.toString(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`  : 0,
				            rewardValue: this.totalDelegatorReward ? this.totalDelegatorReward : 0,
				            reward: this.totalDelegatorReward ? `${Tools.formatStringToFixedNumber(this.totalDelegatorReward.toString(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`: 0,
				            totalAmount:`${Tools.formatStringToFixedNumber((Number(Tools.formatStringToFixedNumber(Tools.numberMoveDecimal(item.amount.toString(),-18),this.fixedNumber)) +
					            Number(Tools.formatStringToFixedNumber(this.totalDelegator.toString(),this.fixedNumber)) +
					            Number(Tools.formatStringToFixedNumber(this.totalUnBondingDelegator.toString(),this.fixedNumber)) +
					            Number(Tools.formatStringToFixedNumber(this.totalDelegatorReward.toString(),this.fixedNumber))).toString(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}` ,
			            }
                    }else {

                    }
                });
            },
            getDelegationList(){
	            Server.commonInterface({delegationList:{
	            	    address: this.$route.params.param
                    }},(res) => {
		            try {
		            	if(res && res.length > 0){
				            this.delegationsItems = res.map(item => {
					            return {
						            address: item.address,
						            amount: `${Tools.formatStringToFixedNumber(item.amount.amount.toString(),this.fixedNumber)} ${item.amount.denom.toUpperCase()}`,
						            shares: (Number(item.shares)).toFixed(2),
						            block: item.height,
						            moniker: item.moniker
					            }
				            });
				            if(res.length > 1){
					            this.totalDelegator = res.reduce( (total,item) => {
						            return Number(item.amount.amount) + Number(total)
					            },0)
                            }else {
					            this.totalDelegator = res[0].amount.amount
                            }
                        }else {
				            this.delegationsItems = []
                        }
		            }catch (e) {
			            console.error(e)
		            }
	            })
            },
            getUnBondingDelegationList(){
	            Server.commonInterface({unDelegationList:{
	            	address: this.$route.params.param
                    }},(res) => {
		            try {
		            	if(res && res.length > 0){
		            		this.unBondingDelegationsItems = res.map( item =>{
		            			return {
		            				address: item.address,
                                    amount: `${Tools.formatStringToFixedNumber(item.amount.amount.toString(),this.fixedNumber)} ${item.amount.denom.toUpperCase()}`,
                                    block: item.height,
                                    endTime: Tools.format2UTC(item.end_time),
						            moniker: item.moniker
                                }
                            });
		            		if(res.length > 1){
					            this.totalUnBondingDelegator = res.reduce( (total,item) => {
						            return Number(item.amount.amount) + Number(total)
					            },0)
                            }else {
					            this.totalUnBondingDelegator = res[0].amount.amount
                            }
                        }
		            }catch (e) {
			            console.error(e)
		            }
	            })
            },
            getRewardsItems(){
	            Server.commonInterface({rewardList:{
	            	address: this.$route.params.param
                    }},(res) => {
		            try {
		            	if(res && res.delagations_rewards.length > 0) {
				            this.rewardsItems = res.delagations_rewards.map( item => {
					            return {
						            address: item.address,
                                    amount: Tools.formatAmount2(item.amount,this.fixedNumber),
						            moniker: item.moniker
				                }
				            });
				            if(res.delagations_rewards.length > 1){
					            this.totalDelegatorReward = Tools.numberMoveDecimal(res.delagations_rewards.reduce( (total,item) => {
						            return Number(item.amount[0].amount) + Number(total)
					            },0),-18);
                            }else {
					            this.totalDelegatorReward = Tools.numberMoveDecimal(res.delagations_rewards[0].amount[0].amount)
                            }

			            }
		            }catch (e) {
			            console.error(e)
		            }
	            })
            },
	        getAllTxType(){
		        Server.commonInterface({addressAllTxType:{
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
	        getTxListByFilterCondition(){
		        let param = {};
		        param.getTxListByAddress = {};
		        param.getTxListByAddress.pageNumber = this.currentPageNum;
		        param.getTxListByAddress.pageSize = this.pageSize;
		        param.getTxListByAddress.txType = this.TxType;
		        param.getTxListByAddress.status = this.txStatus;
		        param.getTxListByAddress.beginTime = this.filterStartTime;
		        param.getTxListByAddress.endTime = this.filterEndTime;
		        param.getTxListByAddress.address = this.$route.params.param;
		        Server.commonInterface(param, (res) => {
			        try {
				        if(res && res.Data) {
					        sessionStorage.setItem('txsTotal',res.Count);
					        this.transactionsItems = res.Data.map( item => {
					        	let Amount;
						        if(item.type === 'BeginUnbonding' || item.type === 'BeginRedelegate'){
							        if(item.status === 'success'){
								        let tokenValue = Tools.formatAccountCoinsAmount(item.tags.balance);
								        let tokenStr = String(Tools.numberMoveDecimal(tokenValue[0],18));
								        item.amount[0].formatAmount =  Tools.formatStringToFixedNumber(tokenStr,2);
								        Amount = item.amount.map(listItem => `${listItem.formatAmount} IRIS`).join(',');
							        }
						        }else {
							        Amount = item.amount && item.amount.length > 0 ? Tools.formatAmount2(item.amount,this.fixedNumber) : '--';
						        }
						        return {
							        txHash: item.hash,
							        block: item.block_height,
							        amount: Amount,
							        txType: item.type,
							        fee: this.formatFee(item.fee),
							        signer: item.signer,
							        status: item.status,
							        timestamp: Tools.format2UTC(item.timestamp),
						        }
					        })
				        }else {
					        this.transactionsItems = []

				        }
			        }catch (e) {
				        console.error(e)
			        }
		        })
	        },
	        getStartTime(time){
		        this.filterStartTime = this.formatStartTime(time)
	        },
	        getEndTime(time){
		        this.filterEndTime = this.formatTime(time)
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
	        formatTime(time){
		        let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
		        let oneDaySeconds = 24 * 60 *60;
		        return Number(new Date(utcTime).getTime()/1000) + Number(oneDaySeconds)
	        },
	        formatStartTime(time){
		        let utcTime = Tools.conversionTimeToUTCByValidatorsLine(new Date(time).toISOString());
		        return Number(new Date(utcTime).getTime()/1000)
	        },
	        formatFee(Fee){
		        if(Fee.amount && Fee.denom){
			        return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(Fee.amount)),4)} ${Tools.formatDenom(Fee.denom).toUpperCase()}`;
		        }
	        },
        }
	}
</script>

<style scoped lang="scss">
    .address_information_container{
        width: 100%;
        .address_information_content{
            max-width: 12.8rem;
            margin: 0 auto;
            padding-bottom: 0.5rem;
            .address_information_header_container{
                width: 100%;
                .address_information_header_content{
                    height: 0.7rem;
                    display: flex;
                    align-items: center;
                    font-weight: bold;
                    font-size: 0.22rem;
                    .address_information_header{
                        padding: 0 0.2rem;
                        color: #171D44;
                    }
                }
            }
            .address_information_assets_container{
                margin-top: 0.18rem;
                .address_information_assets_title{
                    padding: 0.12rem 0 0.12rem 0.2rem;
                    font-size: 0.18rem;
                    font-weight: 500;
                    color: #171D44;
                    line-height: 0.21rem;
                }
            }
            .address_information_delegation_tx_container{
                margin-top: 0.3rem;
                display: flex;
                width: 100%;
                .address_information_delegation_tx_content{
                    flex: 1;
                    margin-right: 0.2rem;
                    .address_information_delegation_title{
                        font-size: 0.18rem;
                        color: #171D44;
                        padding: 0 0 0.12rem 0.2rem;
                    }
                    .address_information_delegation_list_content{
                        width: 100%;
                    }
                }
                .address_information_unbonding_delegation_tx_content{
                    flex: 1;
                    .address_information_unbonding_delegation_title{
                        font-size: 0.18rem;
                        color: #171D44;
                        padding: 0 0 0.12rem 0.2rem;
                    }
                    .address_information_unbonding_delegation_list_content{
                        width: 100%;
                    }
                }
            }
            .address_information_redelegation_header_title{
                font-size: 0.18rem;
                color: #171D44;
                line-height: 0.21rem;
                margin: 0.27rem 0 0 0.2rem;
            }
            .address_information_redelegation_tx_container{
                display: flex;
                width: 100%;
                .address_information_detail_container{
                    flex: 1;
                    .address_information_redelegation_title{
                        font-size: 0.18rem;
                        color: #171D44;
                        padding: 0 0 0.06rem 0.2rem;
                    }
                    .address_information_detail_content{
                        border: 0.01rem solid #E7E9EB;
                        background: #fff;
                        box-sizing: border-box;
                        padding: 0.2rem;
                        height: 2.34rem;
                        .address_information_detail_option{
                            .address_information_detail_option_name{
                                font-size: 0.14rem;
                                color: #787c99;
                                line-height: 0.16rem;
                                margin-right: 0.3rem;
                            }
                            .address_information_detail_option_value{
                                font-size: 0.14rem;
                                color: #171D44;
                                a{
                                    color: var(--bgColor) !important;
                                }
                            }
                            .address_information_address_status_active{
                                background: var(--bgColor);
                                font-size: 0.12rem;
                                color: #fff;
                                padding: 0.02rem 0.14rem;
                                border-radius: 0.22rem;
                            }
                            .address_information_address_status_candidate{
                                background: #3DA87E;
                                font-size: 0.12rem;
                                color: #fff;
                                padding: 0.02rem 0.14rem;
                                border-radius: 0.22rem;
                            }
                            .address_information_address_status_jailed{
                                background: #FA7373;
                                font-size: 0.12rem;
                                color: #fff;
                                padding: 0.02rem 0.14rem;
                                border-radius: 0.22rem;
                            }
                        }
                    }
                }
                .address_information_delegator_rewards_content{
                    flex: 1;
                    margin-right: 0.2rem;
                    .address_information_detail_option{
                        padding: 0 0 0.1rem 0.2rem;
                        .address_information_detail_option_name{
                            font-size: 0.14rem;
                            color: #787C99;
                            margin-right: 0.1rem;
                        }
                        .address_information_detail_option_value{
                            font-size: 0.14rem;
                            a{
                                color: var(--bgColor) !important;
                            }
                        }
                    }
                    .address_information_list_content{
                        width: 100%;
                    }
                }
                .address_information_detail_container{
                    flex: 1;
                }

            }
            .address_information_transaction_container{
                margin: 0.3rem 0 0.5rem 0;
                display: flex;
                flex-direction: column;
                .address_information_transaction_header_content{
                    display: flex;
                    align-items: center;
                    .address_information_transaction_title{
                        font-size: 0.18rem;
                        line-height: 0.21rem;
                        color: #171D44;
                        padding:0 0.4rem 0 0.2rem;
                    }
                    .address_information_list_filter_content{
                        .filter_content{
                            .tx_type_content{
                                display: flex;
                                .tx_type_mobile_content{
                                    display: flex;
                                    align-items: center;
                                    .ivu-select{
                                        margin-right: 0.1rem;
                                        width: 1.3rem;
                                        .ivu-select-item{
                                            text-indent: 0.1rem;
                                            font-size: 0.14rem;
                                            line-height: 0.18rem;
                                            padding: 0.07rem 0.1rem 0.07rem 0;
                                        }
                                    }
                                    .joint_mark{
                                        margin: 0 0.1rem;
                                    }
                                    .ivu-date-picker{
                                        width: 1.3rem;
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
                .address_information_list_content{
                    margin-top: 0.07rem;
                }
            }
        }
    }
</style>
