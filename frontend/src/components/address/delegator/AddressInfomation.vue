<template>
    <div class="address_information_container">
        <page-title :title="pageTitle" :flShowPageLink="false"></page-title>
        <address-information-component :address="headerAddress" :data="assetsItems" :isProfiler="isProfiler"></address-information-component>
        <div class="address_information_content">
            <div class="address_information_delegation_tx_container">
                <div class="address_information_delegation_tx_content">
                    <div class="address_information_delegation_title">Delegations
                        <span class="address_information_delegation_value" v-show="totalDelegatorValue">{{totalDelegatorValue}}</span>
                    </div>
                    <div class="address_information_delegation_list_content">
                        <div>
                            <m-address-information-table
                                    :items="delegationsItems"
                                    :showNoData="delegationsItems.length === 0"
                                    :showNoDataDoc = "'No Delegations'"
                                    listName="delegations"
                                    :width="630">
                            </m-address-information-table>
                        </div>
                    </div>
                    <div class="pagination_content" v-if="flDelegationNextPage">
                        <keep-alive>
                            <m-pagination
                                    :page-size="pageSize"
                                    :total="delegationCountNum"
                                    :page="delegationCurrentPage"
                                    :page-change="delegationPageChange"
                            ></m-pagination>
                        </keep-alive>
                    </div>
                </div>
                <div class="address_information_unbonding_delegation_tx_content">
                    <div class="address_information_unbonding_delegation_title">Unbonding Delegations
                        <span class="address_information_unbonding_delegation_value" v-show="totalUnBondingDelegatorValue">{{totalUnBondingDelegatorValue}}</span>
                    </div>
                    <div class="address_information_unbonding_delegation_list_content">
                        <div>
                            <m-address-information-table
                                    :items="unBondingDelegationsItems"
                                    listName="unBondingDelegations"
                                    :showNoData="unBondingDelegationsItems.length === 0"
                                    :showNoDataDoc="'No Unbonding Delegations'"
                                    :width="630">
                            </m-address-information-table>
                        </div>
                    </div>
                    <div class="pagination_content" v-if="flUnBondingDelegationNextPage">
                        <keep-alive>
                            <m-pagination
                                    :page-size="pageSize"
                                    :total="unBondingDelegationCountNum"
                                    :page="unBondingDelegationCurrentPage"
                                    :page-change="unBondingDelegationPageChange"
                            ></m-pagination>
                        </keep-alive>
                    </div>
                </div>
            </div>
            <div class="address_information_redelegation_header_title">Delegator Rewards
                <span class="address_information_redelegation_rewards_value" v-show="totalDelegatorRewardValue">{{totalDelegatorRewardValue}}</span>
            </div>
            <div class="address_information_redelegation_tx_container">
                <div class="address_information_delegator_rewards_content">
                    <div class="address_information_detail_option">
                        <span class="address_information_detail_option_name">Withdraw To:</span>
                        <span class="address_information_detail_option_value">
                            <router-link :to="`/address/${withdrewToAddress}`">{{withdrewToAddress}}</router-link></span>
                    </div>
                    <div class="address_information_list_content">
                        <div>
                            <m-address-information-table
                                    :items="rewardsItems"
                                    :showNoData="rewardsItems.length === 0"
                                    :showNoDataDoc="'No Delegator Rewards'"
                                    listName="rewards"
                                    :width="630">
                            </m-address-information-table>
                        </div>
                    </div>
                    <div class="pagination_content" v-if="flRewardsDelegationNextPage">
                        <keep-alive>
                            <m-pagination
                                    :page-size="pageSize"
                                    :total="rewardsDelegationCountNum"
                                    :page="rewardsDelegationCurrentPage"
                                    :page-change="rewardsDelegationPageChange"
                            ></m-pagination>
                        </keep-alive>
                    </div>
                </div>
                <div class="address_information_detail_container" :class="OperatorAddress !== '--' ? '' :'hide_style'" :style="{visibility:OperatorAddress && OperatorAddress !== '--' ? 'visible':'hidden'}">
                    <div class="address_information_redelegation_title">Validator Rewards
                        <span class="address_information_validator_rewards_value" v-show="totalValidatorRewards">{{totalValidatorRewards}}</span>
                    </div>
                    <ul class="address_information_detail_content">
                        <li class="address_information_detail_option">
                            <span class="address_information_detail_option_name">Validator Moniker:</span>
                            <div class="validator_status_content">
                                <span class="address_information_detail_option_value">
                                     <router-link v-show="OperatorAddress !== '--' && validatorMoniker !== '--'" :to="`/validators/${OperatorAddress}`">{{validatorMoniker}}</router-link>
                                    <span v-show="OperatorAddress === '--' || validatorMoniker === '--'">{{validatorMoniker}}</span>
                                </span>
                                <span class="address_information_address_status_active" v-if="validatorStatus === 'Active'">Active</span>
                                <span class="address_information_address_status_candidate" v-if="validatorStatus === 'Candidate'">Candidate</span>
                                <span class="address_information_address_status_jailed" v-if="validatorStatus === 'Jailed'">Jailed</span>
                            </div>
                        </li>
                        <li class="address_information_detail_option" style="margin-top: 0.05rem">
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
                    <span class="address_information_transaction_title">{{allTxCountNum}} Txs</span>
                    <div class="address_information_list_filter_content">
                        <div class="filter_content">
                            <div class="tx_type_content">
                                <div class="tx_type_mobile_content">
                                    <el-cascader v-model="value"
                                                 :options="txTypeOption"
                                                 :props="{ expandTrigger: 'hover' }"
                                                 :show-all-levels="false"
                                                 :filterable="true"
                                                 :filter-method="filter"
                                                 @change="filterTxByTxType(value)"></el-cascader>

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
                                <div class="tx_type_mobile_content">
                                    <div class="search_btn" @click="getFilterTxs">Search</div>
                                    <div class="reset_btn" @click="resetFilterCondition"><i class="iconfont iconzhongzhi"></i></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="address_information_list_content">
                    <div>
                        <m-address-information-table
                                :items="transactionsItems"
                                :showNoData="transactionsItems.length === 0"
                                :showNoDataDoc="'No Transactions'"
                                listName="transactions">
                        </m-address-information-table>
                    </div>
                </div>
                <div class="pagination_content" v-if="flAllTxNextPage">
                    <keep-alive>
                        <m-pagination
                                :page-size="addressTxPageSize"
                                :total="allTxCountNum"
                                :page="allTxCurrentPage"
                                :page-change="allTxPageChange"
                        ></m-pagination>
                    </keep-alive>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
	import MAddressInformationTable from "./MAddressInformationTable";
	import Tools from "../../../util/Tools"
	import Server from '../../../service'
	import Constant from "../../../constant/Constant";
	import MPagination from "../../commontables/MPagination";
	import BigNumber from "bignumber.js"
    import moveDecimal from "move-decimal-point"
    import DateTooltip from "../../dateToolTip/DateTooltip";
    import FormatTxType from "../../../util/formatTxType"
    import PageTitle from "../../pageTitle/PageTitle";
	import pageTitleConfig from "../../pageTitle/pageTitleConfig";

    import AddressInformationComponent from "./AddressInformationComponent";
	export default {
		name: "AddressInfomation",
		components: {AddressInformationComponent, PageTitle, DateTooltip, MPagination, MAddressInformationTable},
		data(){
			return {
                pageTitle:pageTitleConfig.StatsIRISRichListAddress,
				withdrewToAddress: '',
				validatorMoniker: '',
                validatorStatus:'',
				OperatorAddress: '',
				isProfiler:false,
                pickerStartTime:sessionStorage.getItem('firstBlockTime') ? sessionStorage.getItem('firstBlockTime') : '',
                PickerOptions: {
                    disabledDate: (time) => {
                        return time.getTime() <= new Date(this.pickerStartTime).getTime() || time.getTime() > Date.now()
                    }
                },
				txTypeOption:[
					{
						value:'allTxType',
						label:'All TxType',
						slot:'allTXType'
					}
				],
				statusValue:JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).status ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).status : 'allStatus',
				startTime: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime : '',
				endTime: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime : '',
				filterStartTime: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).beginTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).beginTime : '',
				filterEndTime: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).endTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).endTime : '',
				TxType: '',
				txStatus: '',
				status:[],
				value: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType  ? this.getRefUrlTxType(JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType) : 'allTxType',
                assetsItems:[],
                delegationsItems:[],
                unBondingDelegationsItems:[],
                rewardsItems:[],
                transactionsItems:[],
				headerAddress:'',
                totalDelegatorReward: 0,
                totalDelegatorRewardValue: 0,
                totalUnBondingDelegator:0,
                totalUnBondingDelegatorValue:0,
                totalDelegator:0,
				totalDelegatorValue:0,
				totalValidatorRewards:0,
                allRewardsValue: 0,
                allRewardsAmountValue:0,
                assetList: [],
				fixedNumber: 2,
				delegationCountNum: 0,
				unBondingDelegationCountNum:0,
				rewardsDelegationCountNum:0,
				delegationCurrentPage:1,
				unBondingDelegationCurrentPage:1,
				rewardsDelegationCurrentPage:1,
				pageSize:5,
                addressTxPageSize: 10,
                delegationPageNationArrayData:[],
                unBondingDelegationPageNationArrayData:[],
				rewardsDelegationPageNationArrayData:[],
                flDelegationNextPage:false,
                flUnBondingDelegationNextPage:false,
                flRewardsDelegationNextPage:false,
                flAllTxNextPage: false,
				allTxCountNum:0,
				allTxCurrentPage: JSON.parse(sessionStorage.getItem('addressPageNum')) ? JSON.parse(sessionStorage.getItem('addressPageNum')) : 1,
                pageShowStartTime: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime : '',
                pageShowEndTime: JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime : ''
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
            this.getRewardsItems();
            this.getAddressInformation();
            this.getAssetList();
            this.getDelegationList();
            this.getUnBondingDelegationList();
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
            statusArray.forEach(item => {
                this.status.push(item)
            })
            this.getTxListByFilterCondition();
        },
        methods:{
            filter(v,iptValue){
                if(Tools.firstWordLowerCase(v.text).includes(Tools.firstWordLowerCase(iptValue))){
                    return true
                }else {
                    return false
                }
            },
            getRefUrlTxType(txType){
                return FormatTxType.getRefUrlTxType(txType);
            },
	        getFilterTxs(){
                let searchCondition = {
                    txType: this.TxType,
                    status: this.txStatus,
                    beginTime: this.filterStartTime,
                    endTime: this.filterEndTime,
                    pageShowStartTime: this.pageShowStartTime,
                    pageShowEndTime: this.pageShowEndTime,
                };
                sessionStorage.setItem('searchResultByTxTypeAndAddress',JSON.stringify(searchCondition))
		        this.allTxCurrentPage = 1;
		        this.resetUrl();
		        sessionStorage.setItem('addressTxPageNum',1);
		        this.getTxListByFilterCondition();
                this.$uMeng.push('Transactions_Search','click')
	        },
			getAddressInformation(){
                Server.commonInterface({addressInformation:{
		                address:this.$route.params.param
                    }},(res) => {
                	try {
                        if(res){
                        	let arrayIndexOneData;
                        	if(res.amount){
                                res.amount.forEach( item => {
                                    if(item.denom === 'iris-atto'){
                                        arrayIndexOneData = item
                                    }
                                });
                                res.amount.unshift(arrayIndexOneData);
                                res.amount = Array.from(new Set(res.amount));
                                this.assetList = res.amount;
                            }
	                        this.validatorMoniker = res.moniker ? res.moniker : '--';
	                        this.OperatorAddress = res.operator_address ? res.operator_address : '--';
	                        this.validatorStatus = res.status;
	                        this.withdrewToAddress = res.withdrawAddress ? res.withdrawAddress : '--';
	                        this.isProfiler = res.isProfiler;
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
                            balanceNumber: item.amount,
				            delegatedValue: this.totalDelegator ? this.totalDelegator : 0,
				            delegated: this.totalDelegator ? `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(this.totalDelegator.toString(),-2)).toFormat(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`: 0,
				            unBondingValue: this.totalUnBondingDelegator ? this.totalUnBondingDelegator : 0,
				            unBonding: this.totalUnBondingDelegator ?`${new BigNumber(Tools.formatStringToFixedNumber(moveDecimal(this.totalUnBondingDelegator.toString(),-2),this.fixedNumber)).toFormat()} ${Constant.Denom.IRIS.toUpperCase()}`  : 0,
				            reward: this.allRewardsValue ? this.allRewardsValue : 0,
                            rewards:this.allRewardsValue ? this.allRewardsValue : 0,
				            totalAmount:`${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal((Number(Tools.formatStringToFixedNumber(Tools.numberMoveDecimal(item.amount.toString(),-18),this.fixedNumber))*100 +
					            Number(Tools.formatStringToFixedNumber(this.totalDelegator.toString(),this.fixedNumber)) +
					            Number(Tools.formatStringToFixedNumber(this.totalUnBondingDelegator.toString(),this.fixedNumber))+
					            Number(Tools.formatStringToFixedNumber(this.allRewardsAmountValue.toString(),this.fixedNumber)) * 100).toString(),-2)).toFormat(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}` ,
			            }
                    }else {
			            return {
				            token: item.denom,
				            balance: item.amount ? `${new BigNumber(item.amount).toFormat()} ${item.denom.toUpperCase()}`: 0,
				            delegated: 0,
				            unBonding: 0,
				            reward: 0,
				            totalAmount: item.amount ? `${new BigNumber(item.amount).toFormat()} ${item.denom.toUpperCase()}`: 0
			            }
                    }
                });
            },
            getDelegationList(){
	            Server.commonInterface({delegationList:{
	            	    address: this.$route.params.param
                    }},(res) => {
		            try {
		            	if(res && res.length > 0){
				            let copyResult = JSON.parse(JSON.stringify(res));
				            this.delegationPageNationArrayData = this.pageNation(copyResult);
				            if(res.length > this.pageSize){
					            this.flDelegationNextPage = true;
                            }else {
					            this.flDelegationNextPage = false;
                            }
				            this.delegationCountNum = res.length;
				            this.delegationPageChange(this.delegationCurrentPage);
				            if(res.length > 0){
                                res.forEach( item => {
                                    if(item.amount && item.amount.amount){
                                        if(item.amount.amount.toString().indexOf('.') !== -1){
                                            let splitNumber = item.amount.amount.toString().split('.')[1].substr(0,2);
                                            item.amount.amount =  Number(`${item.amount.amount.toString().split('.')[0]}.${splitNumber}`) * 100
                                        }else {
                                            item.amount.amount = item.amount.amount * 100
                                        }
                                    }
                                });
                                this.totalDelegator = res.reduce( (total,item) => {
                                    return Number(item.amount.amount) + Number(total)
					            },0)
                            }
                            this.totalDelegatorValue = `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(this.totalDelegator.toString(),-2)).toFormat(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`
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
		            		let copyResult = JSON.parse(JSON.stringify(res));
				            this.unBondingDelegationPageNationArrayData = this.pageNation(copyResult);
				            if(res.length > this.pageSize){
					            this.flUnBondingDelegationNextPage = true
				            }else {
					            this.flUnBondingDelegationNextPage = false
				            }
				            this.unBondingDelegationCountNum = res.length;
                            this.unBondingDelegationPageChange(this.unBondingDelegationCurrentPage);
		            		if(res.length > 0){
                                res.forEach( item => {
                                    if(item.amount && item.amount.amount){
                                        if(item.amount.amount.toString().indexOf('.') !== -1){
                                            let splitNumber = item.amount.amount.toString().split('.')[1].substr(0,2);
                                            item.amount.amount =  Number(`${item.amount.amount.toString().split('.')[0]}.${splitNumber}`) * 100
                                        }else {
                                            item.amount.amount = item.amount.amount * 100
                                        }
                                    }
                                });
					            this.totalUnBondingDelegator = res.reduce( (total,item) => {
						            return Number(item.amount.amount) + Number(total)
					            },0)
                            }
				            this.totalUnBondingDelegatorValue = `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(this.totalUnBondingDelegator.toString(),-2)).toFormat(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`

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
		            	if(res && res.delagations_rewards && res.delagations_rewards.length > 0) {
                            res.delagations_rewards.map( item => {
				            	if(item.amount.length === 0){
						            item.amount.push({
							            amount:0,
                                        denom:'iris-atto'
                                    })
                                }
                            });
                            let copyResult = JSON.parse(JSON.stringify(res));
				            this.totalValidatorRewards = res.commission_rewards ? Tools.formatAmount2(res.commission_rewards,this.fixedNumber) : 0;
		            		this.allRewardsValue = res.total_rewards ? Tools.formatAmount2(res.total_rewards,this.fixedNumber) : 0;
				            this.rewardsDelegationPageNationArrayData = this.pageNation(copyResult.delagations_rewards);
				            if(res.delagations_rewards.length > this.pageSize){
					            this.flRewardsDelegationNextPage = true
				            }else {
					            this.flRewardsDelegationNextPage = false
				            }
                            this.rewardsDelegationCountNum = res.delagations_rewards.length;
				            this.rewardsDelegationPageChange(this.rewardsDelegationCurrentPage);
                            if(res.delagations_rewards.length > 0){
                                res.delagations_rewards.forEach( item => {
                                    if(item.amount && item.amount.length > 0){
                                        item.amount[0].amount = (Tools.formatStringToFixedNumber(Tools.numberMoveDecimal(item.amount[0].amount,-18),this.fixedNumber)) * 100
                                    }
                                })
                                this.totalDelegatorReward = res.delagations_rewards.reduce( (total,item) => {
                                    return Number(item.amount[0].amount) + Number(total)
                                },0);
                            }
                            this.allRewardsAmountValue = res.total_rewards ? Tools.formatStringToFixedNumber(Tools.numberMoveDecimal(res.total_rewards[0].amount,-18),this.fixedNumber) : 0;
                            this.totalDelegatorRewardValue = `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(this.totalDelegatorReward.toString(),-2)).toFormat(),this.fixedNumber)} ${Constant.Denom.IRIS.toUpperCase()}`
                            this.getAssetList()
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
                            let txType = FormatTxType.formatTxType(res);
                            this.txTypeOption = this.txTypeOption.concat(txType);
				        }
			        }catch (e) {
				        console.error(e)
			        }
		        })
	        },
	        getTxListByFilterCondition(){
                let txType = sessionStorage.getItem('searchResultByTxTypeAndAddress') ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType : "",
                    txStatus = sessionStorage.getItem('searchResultByTxTypeAndAddress') ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).status : "",
                    filterStartTime = sessionStorage.getItem('searchResultByTxTypeAndAddress') ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).beginTime : "",
                    filterEndTime = sessionStorage.getItem('searchResultByTxTypeAndAddress') ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).endTime : "",
                    param = {};
                param.getTxListByAddress = {};
		        param.getTxListByAddress.pageNumber = this.allTxCurrentPage;
		        param.getTxListByAddress.pageSize = this.addressTxPageSize;
		        param.getTxListByAddress.txType = txType;
		        param.getTxListByAddress.status = txStatus;
		        param.getTxListByAddress.beginTime = filterStartTime;
		        param.getTxListByAddress.endTime = filterEndTime;
		        param.getTxListByAddress.address = this.$route.params.param;
		        Server.commonInterface(param, (res) => {
			        try {
                        this.allTxCountNum = res.Count;
                        if(res && res.Data) {
					        sessionStorage.setItem('txsTotal',res.Count);
					        if(res.Count > this.addressTxPageSize){
					        	this.flAllTxNextPage = true
                            }else {
						        this.flAllTxNextPage = false
                            }
					        this.transactionsItems = res.Data.map( item => {
					        	let Amount = '--',fromInformation,toInformation;
                                fromInformation = Tools.formatListAmount(item).fromAddressAndMoniker;
                                toInformation = Tools.formatListAmount(item).toAddressAndMoniker;
						        if(item.type === 'BeginUnbonding' || item.type === 'BeginRedelegate'){
							        if(item.status === 'success'){
								        if(item.tags.balance){
									        let tokenValue = Tools.formatAccountCoinsAmount(item.tags.balance);
									        let tokenStr = String(Tools.numberMoveDecimal(tokenValue[0],18));
									        item.amount[0].formatAmount =  new BigNumber(Tools.formatStringToFixedNumber(tokenStr,2)).toFormat();
									        Amount = item.amount.map(listItem => `${listItem.formatAmount} IRIS`).join(',');
                                        }
							        }
						        }else {
						            if(item.amount && item.amount.length > 0){
						                if(item.amount.length > 1){
                                            Amount = `${item.amount.length} Tokens`
                                        }else {
                                            Amount = item.amount && item.amount.length > 0 ? Tools.formatAmount2(item.amount,this.fixedNumber) : '--';
                                        }
                                    }
						        }
						        return {
							        txHash: item.hash,
							        block: item.block_height,
							        amount: Amount,
                                    from: fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].address : '--',
                                    fromMoniker: fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].moniker :'',
                                    to: toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].address : '--',
                                    toMoniker: toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].moniker :'',
							        txType: item.type,
							        fee: this.formatFee(item.fee),
							        signer: item.signer,
							        status: Tools.firstWordUpperCase(item.status),
							        timestamp: Tools.format2UTC(item.timestamp),
                                    isSkipRouter: item.signer === this.$route.params.param,
                                    isFromSkipRouter: fromInformation.length === 1 ? fromInformation[0].address === this.$route.params.param : false,
                                    isToSkipRouter: toInformation.length === 1 ? toInformation[0].address === this.$route.params.param : false
						        }
					        })
				        }else {
					        this.allTxCountNum = 0;
					        this.flAllTxNextPage = false;
					        this.transactionsItems = []

				        }
			        }catch (e) {
				        console.error(e)
			        }
		        })
	        },
	        getStartTime(time){
	            this.pageShowStartTime= time;
		        this.filterStartTime = this.formatStartTime(time)

	        },
	        getEndTime(time){
	            this.pageShowEndTime = time;
		        this.filterEndTime = this.formatEndTime(time)
	        },

	        filterTxByTxType(e){
		        if (Array.isArray(e) && e[e.length-1] === 'allTxType' || e === undefined ) {
			        this.TxType = '';
                    this.$uMeng.push('Transactions_All Type','click')
		        }else {
			        this.TxType = Tools.onlyFirstWordUpperCase(e[e.length-1]);
		        }
	        },
	        filterTxByStatus(e){
		        if(e === 'allStatus'){
			        this.txStatus = ''
                    this.$uMeng.push('Transactions_All Status','click')
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
                let searchCondition = {
                    txType: '',
                    status: '',
                    beginTime: '',
                    endTime: '',
                };
                sessionStorage.setItem('searchResultByTxTypeAndAddress',JSON.stringify(searchCondition))
                this.value = 'allTxType';
		        this.statusValue = 'allStatus';
                this.filterStartTime = '';
                this.filterEndTime = '';
                this.pageShowStartTime = '';
                this.pageShowEndTime = '';
                this.TxType = '';
                this.startTime = '';
		        this.endTime = '';
		        this.allTxCurrentPage = 1;
		        this.resetUrl();
		        this.getTxListByFilterCondition()
                this.$uMeng.push('Transactions_Refresh','click')
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
	        formatFee(Fee){
		        if(Fee.amount && Fee.denom){
			        return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(Fee.amount)),4)} ${Tools.formatDenom(Fee.denom).toUpperCase()}`;
		        }
	        },
            pageNation(dataArray){
	            let index = 0;
	        	let newArray  = [];
	        	if(dataArray.length > this.pageSize){
			        while(index < dataArray.length) {
				        newArray.push(dataArray.slice(index, index += this.pageSize));
			        }
                }else {
			        newArray = dataArray
                }
	            return newArray
            },
	        delegationPageChange(pageNum){
		        pageNum = pageNum - 1
	        	if(this.flDelegationNextPage){
			        this.delegationsItems = this.delegationPageNationArrayData[pageNum].map(item => {
				        return {
					        address: item.address,
					        amount: `${new BigNumber(Tools.formatStringToFixedNumber(item.amount.amount.toString(),this.fixedNumber)).toFormat()} ${item.amount.denom.toUpperCase()}`,
					        shares: new BigNumber((Number(item.shares)).toFixed(2)).toFormat(),
					        block: item.height,
					        moniker: item.moniker
				        }
			        });
                }else {
			        this.delegationsItems = this.delegationPageNationArrayData.map(item => {
				        return {
					        address: item.address,
					        amount: `${new BigNumber(Tools.formatStringToFixedNumber(item.amount.amount.toString(),this.fixedNumber)).toFormat()} ${item.amount.denom.toUpperCase()}`,
					        shares: new BigNumber((Number(item.shares)).toFixed(2)).toFormat(),
					        block: item.height,
					        moniker: item.moniker
				        }
			        });
                }

            },
	        unBondingDelegationPageChange(pageNum){
		        pageNum = pageNum - 1;
		        if(this.flUnBondingDelegationNextPage){
			        this.unBondingDelegationsItems = this.unBondingDelegationPageNationArrayData[pageNum].map( item =>{
				        return {
					        address: item.address,
					        amount: `${new BigNumber(Tools.formatStringToFixedNumber(item.amount.amount.toString(),this.fixedNumber)).toFormat()} ${item.amount.denom.toUpperCase()}`,
					        block: item.height,
					        endTime: Tools.format2UTC(item.end_time),
					        moniker: item.moniker
				        }
			        });
                }else {
			        this.unBondingDelegationsItems = this.unBondingDelegationPageNationArrayData.map( item =>{
				        return {
					        address: item.address,
					        amount: `${new BigNumber(Tools.formatStringToFixedNumber(item.amount.amount.toString(),this.fixedNumber)).toFormat()} ${item.amount.denom.toUpperCase()}`,
					        block: item.height,
					        endTime: Tools.format2UTC(item.end_time),
					        moniker: item.moniker
				        }
			        });
                }

            },
	        rewardsDelegationPageChange(pageNum){
		        pageNum = pageNum - 1;
	        	if(this.flRewardsDelegationNextPage){
			        this.rewardsItems = this.rewardsDelegationPageNationArrayData[pageNum].map( item => {
				        return {
					        address: item.address,
					        amount: Tools.formatAmount2(item.amount,this.fixedNumber),
					        moniker: item.moniker
				        }
			        });
                }else {
			        this.rewardsItems = this.rewardsDelegationPageNationArrayData.map( item => {
				        return {
					        address: item.address,
					        amount: item.amount && item.amount.length > 0 ? Tools.formatAmount2(item.amount,this.fixedNumber) : 0,
					        moniker: item.moniker
				        }
			        });
                }

            },
	        allTxPageChange(pageNum){
                let filterTxType = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType  ?  this.getRefUrlTxType(JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType) : ''
                this.TxType = filterTxType ?  Tools.onlyFirstWordUpperCase(filterTxType[filterTxType.length-1]) : '';
	            this.value = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType  ?  this.getRefUrlTxType(JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).txType) : 'allTxType';
	            this.statusValue = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).status ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).status : 'allStatus';
	            this.startTime = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime : '';
	            this.endTime = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime : '';
	        	this.allTxCurrentPage = pageNum;
                this.pageShowStartTime = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowStartTime : '';
                this.filterStartTime = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).beginTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).beginTime : '';
                this.pageShowEndTime = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).pageShowEndTime : '';
                this.filterEndTime = JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')) && JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).endTime ? JSON.parse(sessionStorage.getItem('searchResultByTxTypeAndAddress')).endTime : '';
	        	sessionStorage.setItem('addressPageNum',pageNum);
	        	this.getTxListByFilterCondition()
            }
        },
        beforeDestroy() {
            let searchCondition = {
                txType: '',
                status: '',
                beginTime: '',
                endTime: ''
            };
            sessionStorage.setItem('searchResultByTxTypeAndAddress',JSON.stringify(searchCondition))
            sessionStorage.setItem('addressPageNum',1)
        }
	}
</script>

<style scoped lang="scss">

    .address_information_container{
        width: 100%;
        .address_information_content{
            max-width: 12.8rem;
            margin: 0 auto;
            padding-bottom: 0.4rem;
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
                        color: #515a6e;
                        font-weight: bold;
                        .address_information_address_status_profiler{
                            background: var(--bgColor);
                            font-size: 0.12rem;
                            color: #fff;
                            padding: 0.02rem 0.14rem;
                            border-radius: 0.22rem;
                            line-height: 0.22rem;
                            height: 0.22rem;
                            display: inline-block;
                        }
                    }
                }
            }
            .address_information_assets_container{
                .address_information_assets_title{
                    padding: 0.12rem 0 0.12rem 0.2rem;
                    font-size: 0.18rem;
                    font-weight: 500;
                    color: #171D44;
                    line-height: 0.21rem;
                }
                .address_information_assets_list_content{
                    &>div{
                        max-width: 12.8rem;
                        overflow-x: auto;
                        background: #fff;
                    }
                }
            }
            .address_information_delegation_tx_container{
                margin-top: 0.3rem;
                display: flex;
                overflow-x: auto;
                overflow-y: hidden;
                .address_information_delegation_tx_content{
                    flex: 1;
                    margin-right: 0.2rem;
                    .address_information_delegation_title{
                        font-size: 0.18rem;
                        color: #171D44;
                        padding: 0 0 0.12rem 0.2rem;
                        .address_information_delegation_value{
                            font-size: 0.14rem;
                            color: #787C99;
                            line-height: 0.16rem;
                            margin-left: 0.15rem;
                        }
                    }
                    .pagination_content{
                        margin-top: 0.2rem;
                        display: flex;
                        justify-content: flex-end;
                    }
                    .address_information_delegation_list_content{
                        width: 100%;
                        overflow-x: auto;
                        &>div{
                            min-width: 6.3rem;
                        }
                    }
                }
                .address_information_unbonding_delegation_tx_content{
                    flex: 1;
                    .address_information_unbonding_delegation_title{
                        font-size: 0.18rem;
                        color: #171D44;
                        padding: 0 0 0.12rem 0.2rem;
                        .address_information_unbonding_delegation_value{
                            font-size: 0.14rem;
                            color: #787C99;
                            line-height: 0.16rem;
                            margin-left: 0.15rem;
                        }
                    }
                    .address_information_unbonding_delegation_list_content{
                        width: 100%;
                        overflow-x: auto;
                        &>div{
                            min-width: 6.3rem;
                        }
                    }
                    .pagination_content{
                        margin-top: 0.2rem;
                        display: flex;
                        justify-content: flex-end;
                    }
                }
            }
            .address_information_redelegation_header_title{
                font-size: 0.18rem;
                color: #171D44;
                line-height: 0.21rem;
                margin: 0.27rem 0 0 0.2rem;
                .address_information_redelegation_rewards_value{
                    font-size: 0.14rem;
                    color: #787C99;
                    line-height: 0.16rem;
                    margin-left: 0.15rem;
                }
            }
            .address_information_redelegation_tx_container{
                display: flex;
                .address_information_detail_container{
                    flex: 1;
                    .address_information_redelegation_title{
                        font-size: 0.18rem;
                        color: #171D44;
                        padding: 0 0 0.06rem 0.2rem;
                        .address_information_validator_rewards_value{
                            font-size: 0.14rem;
                            color: #787C99;
                            line-height: 0.16rem;
                            margin-left: 0.15rem;
                        }
                    }
                    .address_information_detail_content{
                        border: 0.01rem solid #E7E9EB;
                        background: #fff;
                        box-sizing: border-box;
                        padding: 0.2rem;
                        min-height: 2.34rem;
                        .address_information_detail_option{
                            display: flex;
                            align-items: center;
                            .address_information_detail_option_name{
                                width: 1.3rem;
                                font-size: 0.14rem;
                                color: #787c99;
                                line-height: 0.16rem;
                                margin-right: 0.3rem;
                            }
                            .address_information_detail_option_value{
                                font-size: 0.14rem;
                                color: #171D44;
                                margin-right: 0.1rem;
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
                                margin-right: 0.1rem;
                            }
                            .address_information_address_status_candidate{
                                background: #3DA87E;
                                font-size: 0.12rem;
                                color: #fff;
                                padding: 0.02rem 0.14rem;
                                border-radius: 0.22rem;
                                margin-right: 0.1rem;
                            }
                            .address_information_address_status_jailed{
                                background: #FA7373;
                                font-size: 0.12rem;
                                color: #fff;
                                padding: 0.02rem 0.14rem;
                                border-radius: 0.22rem;
                                margin-right: 0.1rem;
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
                        overflow-x: auto;
                        &>div{
                            min-width: 6.3rem;
                        }
                    }
                    .pagination_content{
                        margin-top: 0.2rem;
                        display: flex;
                        justify-content: flex-end;
                    }
                }
                .address_information_detail_container{
                    flex: 1;
                }

            }
            .address_information_transaction_container{
                margin: 0.3rem 0 0 0;
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
                                    /deep/.el-select{
                                        width: 1.3rem;
                                        margin-right: 0.1rem;
                                        .el-input{
                                            .el-input__inner{
                                                padding-left: 0.07rem;
                                                height: 0.32rem;
                                                line-height: 0.32rem;
                                                font-size: 0.14rem !important;
                                                &::-webkit-input-placeholder{
                                                    font-size: 0.14rem !important;
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
                                    /deep/.el-cascader{
                                        width: 1.3rem;
                                        margin-right: 0.1rem;
                                        .el-input{
                                            .el-input__inner{
                                                padding-left: 0.07rem;
                                                height: 0.32rem;
                                                line-height: 0.32rem;
                                                font-size: 0.14rem !important;
                                                &::-webkit-input-placeholder{
                                                    font-size: 0.14rem !important;
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
                                        margin: 0 0.08rem;
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
                .address_information_list_content{
                    margin-top: 0.07rem;
                    overflow-x: auto;
                    background: #fff;
                    &>div{
                        width: 12.8rem;
                    }
                }
                .pagination_content{
                    margin-top: 0.2rem;
                    display: flex;
                    justify-content: flex-end;
                }
            }
        }
    }
    @media screen and (max-width: 910px){
        .address_information_container{
            .address_information_content{
                padding-top: 0;
                .address_information_header_container{
                    .address_information_header_content{
                        .address_information_header{
                            word-break: break-all;
                        }
                    }
                }
                .address_information_transaction_container{
                    .address_information_transaction_header_content{
                        flex-direction: column;
                        align-items: flex-start;
                        box-sizing: border-box;
                        padding: 0 0.1rem;
                        .address_information_transaction_title{
                            padding-left: 0;
                            margin-bottom: 0.1rem;
                        }
                        .address_information_list_filter_content{
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
                                        /deep/ .el-cascader{
                                            margin-right: 0;
                                            width: 1.6rem !important;
                                        }
                                        /deep/ .el-select{
                                            margin-right: 0;
                                            width: 1.6rem !important;
                                        }

                                        /deep/ .el-date-editor{
                                            width: 1.6rem !important;
                                            .el-input__prefix{
                                            }
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
                .address_information_assets_container{
                    .address_information_assets_list_content{
                        margin: 0 0.1rem;
                    }
                }
                .address_information_transaction_container{
                    .address_information_list_content{
                        margin: 0 0.1rem;
                    }
                    .pagination_content{
                        margin-right: 0.1rem;
                    }
                }
                .address_information_redelegation_header_title{
                    margin-left: 0.1rem;
                }
            }

        }
        .address_information_content{
            .address_information_assets_container{
                padding-top: 0 !important;
            }
        }
    }
    @media screen and (max-width: 1280px){
        .address_information_container{
            .address_information_content{
                .address_information_delegation_tx_container{
                    flex-direction: column;
                    margin: 0 0.1rem;
                    .address_information_delegation_tx_content{
                        width: 100%;
                        .address_information_delegation_title{
                            padding: 0.2rem 0 0.2rem;
                        }
                        .address_information_delegation_list_content{
                            overflow-x: auto;
                        }
                    }
                    .address_information_unbonding_delegation_tx_content{
                        .address_information_unbonding_delegation_title{
                            padding: 0.2rem 0 0.2rem;
                        }
                        .address_information_unbonding_delegation_value{
                            overflow-x: auto;
                        }
                    }
                }
                .address_information_redelegation_tx_container{
                    flex-direction: column;
                    margin: 0 0.1rem;
                    .address_information_delegator_rewards_content{
                        margin-right: 0;
                        .address_information_detail_option{
                            padding-left: 0;
                            display: flex;
                            flex-direction: column;
                        }
                        .address_information_list_content{
                            overflow-x: auto;
                        }
                    }
                    .address_information_detail_container{
                        .address_information_redelegation_title{
                            padding: 0.2rem 0;
                        }
                        .address_information_detail_content{
                            .address_information_detail_option{
                                display: flex;
                                align-items: flex-start;
                                flex-direction: column;
                                .validator_status_content{
                                    display: flex;
                                    margin: 0.05rem 0;
                                }
                                .address_information_detail_option_value{
                                }
                            }
                        }
                    }
                    .hide_style{
                        display: none;
                    }
                }
            }
        }
    }
</style>
