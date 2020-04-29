<template>
	<div class="validator_bianjie_information_container">
		<h3 class="validator_bianjie_node_title">Delegate for IRISnet-Bianjie</h3>
		<p class="validator_operator_address_content"><span>{{operatorAddress}}<m-clip :text="operatorAddress"></m-clip></span></p>
		<ul class="validator_information_list_content">
			<li class="validator_information_item">
				<p class="validator_information_value">{{atomToken}}</p>
				<p class="validator_information_name">Bonded ATOMs</p>
			</li>
			<li class="validator_information_item">
				<p class="validator_information_value">{{votingPower}}</p>
				<p class="validator_information_name">Voting Power</p>
			</li>
			<li class="validator_information_item">
				<p class="validator_information_value">{{commissionRate}}</p>
				<p class="validator_information_name">Commission Rate</p>
			</li>
			<li class="validator_information_item">
				<p class="validator_information_value">{{uptime}}</p>
				<p class="validator_information_name">Uptime</p>
			</li>
			<li class="validator_information_item">
				<span>
					<a @click="$uMeng.push('Delegate_more','click')" href="https://irisnet.org/irisnet-bianjie" target="_blank">More</a>
				</span>
			</li>
		</ul>
<!--		<div class="validator_bianjie_information_content" style="color:#fff;">{{atomToken}} {{votingPower}} {{commissionRate}} {{uptime}} {{operatorAddress}} <m-clip :text="operatorAddress"></m-clip> </div>-->
	
	</div>
</template>

<script>
	import axios from "axios"
	import moveDecimalPoint from "move-decimal-point"
	import bigNumber from "bignumber.js"
	import MClip from "../commontables/MClip";
	export default {
		name: "ValidatorBianjieInformation",
		components: {MClip},
		data() {
			return {
				url:'https://stage.irisnet.org',
				bondedToken:'/bondedTokens',
				signedBlocksWindow:	'/signedBlocksWindow',
				missedBlockCounter:'/missedBlockCounter',
				validators:'/validators',
				commission:'',
				atomToken:'',
				votingPower:'',
				sourceData:'',
				commissionRate: '',
				uptime:'',
				signedBlocks:'',
				operatorAddress:'',
			}
		},
		mounted(){
			this.getValidatorInformation();
			this.getBondedTokens();
			this.getMissedBlockCounter();
			this.getSignedBlocksWindow();
		},
		methods:{
			//TODO  先处理 bondedTokens
			
			getValidatorInformation(){
				axios.get(`${this.url}${this.validators}`).then(data => {
					if(data.status === 200){
						return data.data
					}
				}).then(result => {
					if(result && JSON.stringify(result) !== '{}' && result.result && result.result.tokens){
						this.atomToken = this.formatAtomTokens(result.result.tokens);
						this.operatorAddress = result.result.operator_address;
						//计算votingPower 源数据
						this.sourceData = result.result.tokens;
						let commissionNumber = ( Number(result.result.commission.commission_rates.rate) * 100 ).toFixed(2)
						this.commissionRate = `${commissionNumber} %`
					}else {
						this.atomToken = 0
					}
				}).catch(e => {
					console.error(e)
				})
			},
			
			getBondedTokens(){
				axios.get(`${this.url}${this.bondedToken}`).then(data => {
					if(data.status === 200){
						return data.data
					}
				}).then(result => {
					if(result && JSON.stringify(result) !== '{}' && result.result){
						let promiseAtomToken = new Promise(resolve => {
							if(this.sourceData !== ''){
								resolve(this.sourceData)
							}else {
								this.getBondedTokens();
							}
						});
						promiseAtomToken.then( res =>{
							if(res){
								let votingPower = ( Number(res) / Number(result.result.bonded_tokens) * 100 ) .toFixed(2)
								this.votingPower = `${votingPower} %`
							}else {
								this.votingPower = '0.00 %'
							}
						}).catch(e => {
						
						})
					}else {
						this.votingPower = ''
					}
				}).catch(e => {
					console.error(e)
				})
			},
			
			getSignedBlocksWindow(){
				axios.get(`${this.url}${this.signedBlocksWindow}`).then(data => {
					if(data.status === 200){
						return data.data
					}
				}).then(result => {
					if(result && JSON.stringify(result) !== '{}' && result.result && result.result.signed_blocks_window){
						this.signedBlocks = result.result.signed_blocks_window
					}else {
						this.signedBlocks = 0
					}
				}).catch(e => {
					console.error(e)
				})
			},
			
			getMissedBlockCounter(){
				axios.get(`${this.url}${this.missedBlockCounter}`).then(data => {
					if(data.status === 200){
						return data.data
					}
				}).then(result => {
					if(result && JSON.stringify(result) !== '{}' && result.result ){
						let totalSignedBlocksWindow = new Promise(resolve => {
							if(this.signedBlocks !== ''){
								resolve (this.signedBlocks)
							}else {
								this.getMissedBlockCounter()
							}
						});
						totalSignedBlocksWindow.then( value => {
							let uptimeNumber = ( ( ( Number(value) - Number(result.result.missed_blocks_counter) ) / Number(value) ) * 100 ).toFixed(4)
							this.uptime = `${uptimeNumber} %`
						})
					}else {
						this.uptime = 0
					}
					
				}).catch(e => {
					console.error(e)
				})
			},
			formatAtomTokens(atomTokens){
				let formatedData ='';
				//除以10的6次方
				let leftLength = -6;
				formatedData = new bigNumber( moveDecimalPoint(atomTokens.toString(),leftLength) ).toFormat()
				return formatedData
			}
		}
	}
</script>

<style scoped lang="scss">
	.validator_bianjie_information_container{
		margin-top: 0.4rem;
		.validator_bianjie_node_title{
			font-size: 0.18rem;
			line-height: 0.21rem;
			color: #ffff;
			padding-left: 0.2rem;
			font-weight: lighter;
		}
		.validator_operator_address_content{
			margin-top: 0.12rem;
			padding-left: 0.2rem;
			word-break: break-all;
			word-wrap: break-word;
			span{
				color: rgba(115, 122, 174, 1);
			}
			div{
				padding-left: 0.1rem;
			}
		}
		.validator_information_list_content{
			display: flex;
			margin-top: 0.15rem;
			margin-bottom: 0.4rem;
			.validator_information_item{
				box-sizing: border-box;
				padding: 0.29rem 0;
				flex: 1;
				background: rgba(45, 50, 90, 0.64);
				border-radius: 0.04rem;
				text-align: center;
				margin-right: 0.02rem;
				display: flex;
				align-content: center;
				justify-content: center;
				flex-direction: column;
				.validator_information_value{
					color: #fff;
					font-size: 0.16rem;
					line-height: 0.18rem;
					margin-bottom: 0.1rem;
				}
				.validator_information_name{
					font-size: 0.14rem;
					line-height: 0.16rem;
					color: rgba(134, 143, 211, 1);
				}
			}
			.validator_information_item:last-child{
				max-width: 0.6rem;
				span{
					a{
						border-bottom: 0.01rem solid rgba(55, 124, 248, 1);
						color: rgba(55, 124, 248, 1) !important;
					}
				}
			
			}
		}
	}
	@media (max-width: 768px) {
		.validator_bianjie_information_container{
			.validator_bianjie_node_title{
				padding-left: 0;
			}
			.validator_operator_address_content{
				padding-left: 0;
			}
			.validator_information_list_content{
				flex-direction: column;
				.validator_information_item{
					margin-top: 0.1rem;
				}
				.validator_information_item:last-child{
					max-width: none !important;
				}
			}
		}
	}
</style>
