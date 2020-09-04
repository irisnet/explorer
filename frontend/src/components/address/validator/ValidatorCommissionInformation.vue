<template>
	<div class="validator_commission_information_container">
		<p class="validator_commission_information_title">
			<span>Commission Info</span>
		</p>
		<div class="validator_commission_information_wrap">
			<div class="validator_commission_information_content">
				<div class="validator_commission_information_scatter_content">
					<p class="validator_commission_information_scatter_title">Commission Rate & Bonded Tokens Distribution</p>
					<validator-detail-scatter :jailedData="jailedData" :validatorStatus="validatorStatus"></validator-detail-scatter>
				</div>
				<div class="validator_commission_bonded_container">
					<ul class="validator_commission_bonded_list">
						<li class="validator_commission_bonded_item" v-for="(item,index) in bondedAndCommissionArr" :key="index">
							<p class="validator_commission_parent_content">
								<span>{{item.label}} <i class="iconfont" style="cursor: pointer;color: var(--bgColor)" @click="showChildren(index)" v-if="item.flShowSelectIcon" :class="item.flShowChildren ? 'iconInitialDeposit' : 'iconwangluoqiehuanjiantou'"></i></span>
								<span>{{item.value}}</span>
							</p>
							<ul class="validator_commission_children_content" v-if="item.flShowChildren">
								<li class="validator_commission_children_item" v-for="(childrenItem,childrenIndex) in item.children" :key="childrenIndex">
									<span>{{childrenItem.label}}</span>
									<span>{{childrenItem.value}}</span>
								</li>
							</ul>
						</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import Constants from "../../../constant/Constant";
	import Tools from "../../../util/Tools";
	import Service from "../../../service";
	import ValidatorDetailScatter from "./ValidatorDetailScatter";
	export default {
		name: "ValidatorCommissionInformation",
		components: {ValidatorDetailScatter},
		props:{
			validationCommissionInfo:{
				type:Object
			}
		},
		data(){
			return {
				informationData:'',
				validatorStatus:'',
				jailedData:{},
				irisTokenFixedNumber:6,
				irisTokenMaxFixedNumber:18,
				bondedAndCommissionArr:[
					{
						label:'Commission Rate:',
						dataName:'commission_rate',
						value:'',
						flShowSelectIcon:false,
						flShowChildren:false,
						children:[
						
						]
					},
					{
						label:'Bonded Tokens:',
						dataName:'bonded_tokens',
						value:'',
						flShowSelectIcon:true,
						flShowChildren:false,
						children:[
						
						]
					},
					{
						label:'Delegators:',
						dataName:'delegator_count',
						flShowSelectIcon:false,
						flShowChildren:false,
						value:'',
						children:[
						
						]
					},
					{
						label:'Total Shares:',
						dataName:'delegator_shares',
						flShowSelectIcon:false,
						flShowChildren:false,
						value:'',
						children:[
						
						]
					},
					{
						label:'Commission Rewards:',
						dataName:'commissionRewards',
						flShowSelectIcon:false,
						flShowChildren:false,
						value:'',
						children:[
						
						]
					},
					{
						label:'Commission Rate Range:',
						dataName:'commissionRateRange',
						flShowSelectIcon:true,
						flShowChildren:false,
						value:'',
						children:[
						
						]
					},
				]
			}
		},
		watch:{
			validationCommissionInfo(){
				this.informationData = this.validationCommissionInfo;
				this.handlePropsData()
			}
		},
		mounted(){
			this.informationData = this.validationCommissionInfo;
		},
		methods:{
			showChildren(index){
				if(index === 1 || index === 5){
					this.bondedAndCommissionArr[index].flShowChildren = !this.bondedAndCommissionArr[index].flShowChildren;
				}
			},
			handlePropsData(){
				let dataInfomation = this.informationData;
				this.jailedData.bonded_tokens = dataInfomation.bonded_tokens;
				this.jailedData.commission_rate = dataInfomation.commission_rate;
				this.jailedData.operator_address = dataInfomation.operator_addr;
				this.jailedData.moniker = dataInfomation.description.moniker;
				this.validatorStatus = dataInfomation.status;
				this.bondedAndCommissionArr.forEach( item => {
					if(item.label === 'Commission Rate Range:'){
						item.value = `0 ~ ${Number(dataInfomation.commission_max_rate) * 100} %`;
						item.children.push({
							label:'Max Change Rate Everytime:',
							value: `0 ~ ${Number(dataInfomation.commision_max_change_rate) * 100} %`
						})
					}else if(item.label === 'Bonded Tokens:'){
						item.value =`${this.$options.filters.amountFromat(dataInfomation.bonded_tokens, Constants.Denom.IRIS.toUpperCase(), this.irisTokenFixedNumber)}`;
						let selfBonded = {
							label:'Self-Bonded:',
							value: `${this.$options.filters.amountFromat(
								dataInfomation.self_bonded, Constants.Denom.IRIS.toUpperCase(),
								this.irisTokenFixedNumber)}
								(${this.formatPerNumber((dataInfomation.self_bonded / Number(dataInfomation.bonded_tokens)) * 100)} %)`
						};
						let delegatorBonded = {
							label:'Delegator Bonded:',
							value:`${this.$options.filters.amountFromat(
								dataInfomation.bonded_stake, Constants.Denom.IRIS.toUpperCase(), this.irisTokenFixedNumber)}
								 (${this.formatPerNumber((Number(dataInfomation.bonded_stake) / Number(dataInfomation.bonded_tokens)) * 100)} %)`
						};
						item.children.unshift(selfBonded,delegatorBonded)
					}else if(item.label === 'Total Shares:'){
						item.value = `${this.$options.filters.amountFromat(dataInfomation.delegator_shares, "", this.irisTokenFixedNumber)}`
					}else if(item.label === 'Commission Rate:'){
						item.value = dataInfomation.commission_update !== "0001-01-01 00:00:00 +0000 UTC"
							? dataInfomation.status === 'active' ? `${this.formatPerNumber(Number(dataInfomation.commission_rate) * 100)} % (${Tools.format2UTC(dataInfomation.commission_update).substr(0,10)} Updated)` : `${this.formatPerNumber(Number(dataInfomation.commission_rate) * 100)} %`
							: `${this.formatPerNumber(Number(dataInfomation.commission_rate) * 100)} %`
					}else {
						item.value = dataInfomation[item.dataName]
					}
				});
				this.getValidatorRewards();
			},
			formatPerNumber(num) {
				if (typeof num === "number" && !Object.is(num, NaN)) {
					let afterPoint = String(num).split(".")[1];
					let afterPointLong = (afterPoint && afterPoint.length) || 0;
					if (afterPointLong > 2 && num !== 0) {
						return num.toFixed(4);
					} else {
						return num.toFixed(2);
					}
				}
				return num;
			},
			getValidatorRewards() {
				Service.commonInterface(
					{
						validatorCommissionRewards: {
							validatorAddr: this.$route.params.param
						}
					},
					data => {
						try {
							if (data) {
								if (Array.isArray(data) && data[0]) {
									this.bondedAndCommissionArr.map(item => {
										if(item.label === 'Commission Rewards:'){
											return item.value = Tools.formatAmount2(data[0]) || '--'
										}
									})
								}else {
									this.bondedAndCommissionArr.map(item => {
										if(item.label === 'Commission Rewards:'){
											return item.value = '--'
										}
									})
								}
							}
						} catch (e) {
							console.error(e)
						}
					}
				);
			}
		}
	}
</script>

<style scoped lang="scss">
	.validator_commission_information_container{
		width: 100%;
		.validator_commission_information_title{
			max-width: 12.8rem;
			margin: 0 auto;
			span{
				display: inline-block;
				margin: 0.35rem 0 0.12rem 0;
				color: var(--titleColor);
				font-size: 0.18rem;
				line-height: 0.21rem;
				padding-left: 0.2rem;
			}
		}
		.validator_commission_information_wrap{
			max-width: 12.8rem;
			margin: 0 auto;
			background: #fff;
			border: 0.01rem solid #E7E9EB;
			.validator_commission_information_content{
				display: grid;
				grid-template-columns: repeat(1,50% 50%);
				box-sizing: border-box;
				padding: 0.3rem 0.2rem;
				.validator_commission_information_scatter_content{
					width: 100%;
					.validator_commission_information_scatter_title{
						font-size: 0.14rem;
						line-height: 0.16rem;
						color: var(--titleColor);
					}
				}
				.validator_commission_bonded_container{
					.validator_commission_bonded_list{
						box-sizing: border-box;
						padding-left: 0.7rem;
						border-left: 0.01rem dashed #D7DCE0;
						.validator_commission_bonded_item{
							display: flex;
							flex-direction: column;
							margin-top: 0.2rem;
							width: 100%;
							.validator_commission_parent_content{
								display: flex;
								span:nth-of-type(1){
									width: 1.9rem;
									color: var(--contentColor);
								}
								span:nth-of-type(2){
									color: var(--titleColor)
								}
							}
							.validator_commission_children_content{
								display: flex;
								flex-direction: column;
								margin-top: 0.1rem;
								border: 0.01rem dashed #D7DCE0;
								box-sizing: border-box;
								padding: 0.1rem;
								.validator_commission_children_item{
									margin-top: 0.05rem;
									span:nth-of-type(1){
										display: inline-block;
										width: 1.79rem;
										color: var(--contentColor);
										font-size: 0.12rem;
									}
									span:nth-of-type(2){
										display: inline-block;
										color: var(--titleColor);
										font-size: 0.12rem;
									}
								}
								.validator_commission_children_item:first-child{
									margin-top: 0;
								}
							}
						}
						.validator_commission_bonded_item:first-child{
							margin-top: 0;
						}
					}
				}
			}
		}
	}
	@media screen and (max-width: 1050px){
		.validator_commission_information_container{
			.validator_commission_information_wrap{
				margin: 0 0.2rem;
				
				.validator_commission_information_content{
					width: 100%;
					grid-template-columns: repeat(1,auto);
					grid-template-rows: repeat(2,auto);
					.validator_commission_bonded_container{
						margin-top: 0.3rem;
						.validator_commission_bonded_list{
							border-left: 0;
							padding-left: 0;
							padding-top: 0.3rem;
							border-top: 0.01rem solid #D7DCE0;
							.validator_commission_bonded_item{
								.validator_commission_children_content{
									margin-right: 0.2rem;
								}
							}
						}
					}
				}
			}
		}
	}
	@media screen and (max-width: 750px){
		.validator_commission_information_container{
			.validator_commission_information_wrap{
				margin: 0 0.1rem;
				overflow-x: auto;
				.validator_commission_information_content{
					width: 100%;
					.validator_commission_information_scatter_content{
						overflow-x: auto;
					}
				}
			}
		}
	}
</style>
