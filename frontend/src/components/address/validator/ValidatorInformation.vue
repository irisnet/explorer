<template>
	<div class="validation_information_container">
		<div class="validation_information_wrap">
			<div class="validation_information_content">
				<div class="validation_information_header_content">
					<div class="validation_information_header_img">
						<img v-if="validatorIconHref" :src="validatorIconHref" alt="" @error="imgLoadError()">
						<span v-else>{{validatorIconSrc}}</span>
					</div>
					<p class="validation_information_moniker">{{moniker}}</p>
					<p class="validation_information_website" v-if="website">
						<span class="validation_website_link" @click="openUrl(website)">{{website}}</span>
					</p>
					<p class="validation_information_identity" v-if="identity">
						<a class="validation_information_link" v-if="keyBaseName" :href="keyBaseName" target="_blank">{{identity}}</a>
						<span class="validation_information_not_link" v-else>{{identity}}</span>
					</p>
					<p class="validation_information_details" v-if="details">{{details}}</p>
					<p class="validation_information_no_more" v-if="!website && !identity && !details">~ The validator has no more information ~</p>
				</div>
				<div class="validation_information_asset_information_content">
					<div class="validation_information_status_content">
						<span class="status_btn" v-if="validatorStatus === 'Active'">Active</span>
						<span
								class="status_btn"
								style="background-color: #3DA87E;"
								v-if="validatorStatus === 'Candidate'"
						>Candidate</span>
						<span
								class="status_btn"
								style="background-color: #FA7373;"
								v-if="validatorStatus === 'Jailed'"
						>Jailed</span>
					</div>
					<ul class="validation_information_list_content">
						<li class="validation_information_item" v-for="(item, index) in validationAssetInfoArr" :key="index" v-show="showVotingPower(validatorStatus,item.label)">
							<div class="validation_information_item_label_content">
								<span>{{item.label}}</span>
								<el-tooltip v-if="item.isToolTip" class="tip_content" :content="'The address you used to Create a Validator , Delegate or Withdraw DelegatorReward, etc.'">
									<span class="iconfont icontishi"></span>
								</el-tooltip>
							</div>
							<div class="validation_information_item_value_content">
								<p>
									<span class="validation_information_item_value">
										<el-tooltip v-if="item.isCopyIcon" :content="`${item.value}`">
											<span v-if="item.flAddressLink" style="font-family: Consolas,Menlo;">{{formatAddress(item.value)}}</span>
											<router-link style="font-family: Consolas,Menlo;" v-if="!item.flAddressLink" :to="addressRoute(item.value)" :style="{color: item.isValidatorAddress ? '' :'var(--bgColor) !important'}">{{formatAddress(item.value)}}</router-link>
										</el-tooltip>
										<span v-if="!item.isCopyIcon">
											{{item.value}}
										</span>
										<m-clip v-if="item.isCopyIcon" style="margin-left: 0.06rem;cursor: pointer;" :text="item.value ? item.value : ''" ></m-clip>
									</span>
								</p>
							</div>
						</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import MClip from "../../commontables/MClip";
	import Tools from "../../../util/Tools";
	import axios from "../../../util/axios";
	import Service from "../../../service"
	export default {
		name: "ValidationInformation",
		components: {MClip},
		props:{
			validationInformation:{
				type: Object
			},
			validatorStatus:{
				type: String
			}
		},
		data () {
			return {
				informationData:"",
				validatorIconSrc:'',
				validatorIconHref:"",
				moniker:'',
				website:'',
				identity:'',
				details:'',
				validationStatus:'',
				keyBaseName:'',
				validationAssetInfoArr:[
					{
						label:'Operator Address:',
						dataName:'operator_addr',
						value:'',
						isToolTip:true,
						isCopyIcon:true,
						isValidatorAddress:true,
						flAddressLink:false,
					},
					{
						label:'Owner Address:',
						dataName:'owner_addr',
						value:'',
						isToolTip:false,
						isCopyIcon:true,
						flAddressLink:false,
					},
					{
						label:'Withdraw Address:',
						dataName:'address',
						value:'',
						isToolTip:false,
						isCopyIcon:true,
						flAddressLink:false,
					},
					{
						label:'Voting Power:',
						dataName:'self_power',
						value:'',
						isToolTip:false,
						isCopyIcon:false,
						flAddressLink:false,
					},
					{
						label:'Uptime:',
						dataName:'uptime',
						value:'',
						isToolTip:false,
						isCopyIcon:false,
						flAddressLink:false,
					},
					{
						label:'Missed Blocks:',
						dataName:'missed_blocks_count',
						value:'',
						isToolTip:false,
						isCopyIcon:false,
						flAddressLink:false,
					},
					{
						label:'Bond Height:',
						dataName:'bond_height',
						value:'',
						isToolTip:false,
						isCopyIcon:false,
						flAddressLink:false,
					},
					{
						label:'Unbonding Height:',
						dataName:'unbond_height',
						value:'',
						isToolTip:false,
						isCopyIcon:false,
						flAddressLink:false,
					},
					{
						label:'Consensus Pubkey:',
						dataName:'consensus_addr',
						value:'',
						isToolTip:false,
						isCopyIcon:true,
						flAddressLink:true,
					},
					{
						label:'Jailed Until:',
						dataName:'jailed_until',
						value:'',
						isToolTip:false,
						isCopyIcon:false,
						flAddressLink:false,
					}
				]
			}
		},
		watch:{
			validationInformation(){
				this.informationData = this.validationInformation;
				this.handlePropsData()
			}
		},
		mounted () {
			this.informationData = this.validationInformation;
			this.getValidatorWithdrawAddr()
		},
		methods:{
			imgLoadError(){
				this.validatorIconHref = ""
			},
			showVotingPower(validatorStatus,labelName){
				if(validatorStatus === 'Candidate' || validatorStatus === 'Jailed'){
					if(labelName === 'Voting Power:' || labelName === 'Bond Height:'){
						return false
					}else {
						return true
					}
				}else {
					if(labelName === 'Jailed Until:' || labelName === 'Unbonding Height:'){
						return false
					}
					return true
				}
				
			},
			handlePropsData(){
				let information = this.informationData;
				//匹配第一个字符用于头像展示
				let regex =  /[^\w\u4e00-\u9fa50-9a-zA-Z]/g;
				let replaceMoniker = information.description.moniker.replace(regex,'');
				this.validatorIconSrc = replaceMoniker ? Tools.firstWordUpperCase(replaceMoniker.match(/^[0-9a-zA-Z\u4E00-\u9FA5]/g)[0]) : '';
				this.validatorIconHref = information.icons ? information.icons : replaceMoniker ? '' : require('../../../assets/default_validator_icon.svg');
				this.moniker = information.description.moniker;
				this.website = information.description.website ? information.description.website : '';
				if(information.description.identity){
					this.getKeyBaseName(information.description.identity)
				}
				this.identity = information.description.identity ? information.description.identity : '';
				this.details = information.description.details ? information.description.details : '';
				this.validationAssetInfoArr.forEach( item => {
					if(item.dataName !== 'address'){
						if(item.dataName === 'uptime'){
							item.value = Tools.FormatUptime(information[item.dataName]);
						}else if(item.dataName === 'self_power') {
							item.value = information.status === "Active"
								? `${information.self_power} (${this.formatPerNumber((information.self_power / information.total_power) * 100)} %)`
								: "";
						}else if(item.dataName === 'missed_blocks_count'){
							item.value = `${information.missed_blocks_count} in ${information.stats_blocks_window} blocks`;
						}else if(item.dataName === 'jailed_until'){
							item.value = new Date(information[item.dataName]).getTime() ? Tools.format2UTC(information[item.dataName]) : "--";
						}else {
							item.value = information[item.dataName];
						}
					}
				})
			},
			openUrl(url) {
				url = url.trim();
				if (url) {
					if (!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)) {
						url = `http://${url}`;
					}
					window.open(url);
				}
			},
			getKeyBaseName(identity) {
				let url = `https://keybase.io/_/api/1.0/user/lookup.json?fields=basics&key_suffix=${identity}`;
				if (identity) {
					axios.http(url).then(res => {
						if (res.them && res.them.length > 0 && res.them[0].basics && res.them[0].basics.username) {
							this.keyBaseName = `https://keybase.io/${res.them[0].basics.username}`;
						}else {
							this.keyBaseName = ''
						}
					});
				}
			},
			formatAddress (address) {
				return Tools.formatValidatorAddress(address);
			},
			getValidatorWithdrawAddr() {
				Service.commonInterface(
					{
						validatorWithdrawAddr: {
							validatorAddr: this.$route.params.param
						}
					},
					(data) => {
						try {
							if (data) {
								this.validationAssetInfoArr.forEach( item => {
									if(item.dataName === 'address')
									item.value = data[item.dataName];
								})
							}
						} catch (e) {
							console.error(e)
						}
					}
				);
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
		}
	}
</script>

<style scoped lang="scss">
.validation_information_container{
	width: 100%;
	margin-top: 0.74rem;
	.validation_information_wrap{
		max-width: 12.8rem;
		margin: 0 auto;
		.validation_information_content{
			width: 100%;
			background: #fff;
			display: grid;
			grid-template-columns: repeat(2,50%);
			.validation_information_header_content{
				display: flex;
				flex-direction: column;
				align-items: center;
				box-sizing: border-box;
				margin: 0.3rem 0 0.38rem 0;
				.validation_information_header_img{
					width: 1.2rem;
					height: 1.2rem;
					border-radius: 0.6rem;
					background: #E0E8FF;
					overflow: hidden;
					display: flex;
					align-items: center;
					justify-content: center;
					img{
						width: 100%;
					}
					span{
						font-size: 0.52rem;
					}
				}
				.validation_information_moniker{
					margin-top: 0.1rem;
					font-size: 0.2rem;
					line-height: 0.23rem;
					font-weight: bold;
				}
				.validation_information_website{
					margin-top: 0.15rem;
					.validation_website_link{
						font-size: 0.14rem;
						line-height: 0.16rem;
						color:var(--bgColor);
						cursor: pointer;
					}
				}
				.validation_information_identity{
					margin-top: 0.1rem;
					.validation_information_link{
						color:var(--bgColor) !important;
						font-size: 0.14rem;
						line-height: 0.16rem;
					}
					.validation_information_not_link{
						color:var(--titleColor);
						font-size: 0.14rem;
						line-height: 0.16rem;
					}
				}
				.validation_information_details{
					max-width: 5rem;
					margin-top: 0.1rem;
					font-size: 0.14rem;
					line-height: 0.21rem;
					color: var(--titleColor);
				}
				.validation_information_no_more{
					margin-top: 0.55rem;
					color: var(--contentColor);
					font-size: 0.14rem;
					line-height: 0.21rem;
				}
			}
			.validation_information_asset_information_content{
				box-sizing: border-box;
				margin: 0.3rem 0;
				border-left: 0.01rem dashed #D7DCE0;
				padding:0 0 0.3rem 0.7rem;
				.validation_information_status_content{
					.status_btn{
						height: 0.26rem;
						line-height: 0.26rem;
						padding: 0.05rem 0.16rem;
						background: var(--bgColor);
						color:#fff;
						border-radius: 0.13rem;
					}
				}
				.validation_information_list_content{
					.validation_information_item{
						display: grid;
						grid-template-columns: repeat(1,1.5rem auto);
						margin-top: 0.16rem;
						.validation_information_item_label_content{
							span{
								color: var(--contentColor);
							}
							.tip_content{
								margin-left: 0.05rem;
							}
						}
						.validation_information_item_value_content{
							display: flex;
							span{
								word-break: break-all;
							}
							.validation_information_item_value{
								margin-right: 0.06rem;
							}
						}
					}
				}
			}
		}
	}
}
	@media screen and (max-width: 1050px){
		.validation_information_container{
			.validation_information_wrap{
				margin: 0 0.2rem;
				.validation_information_content{
					width: 100%;
					grid-template-columns: repeat(1,auto);
					grid-template-rows: repeat(2,auto);
					.validation_information_header_content{
						align-items: flex-start;
						margin:0.3rem 0.15rem;
					}
					.validation_information_asset_information_content{
						border-top: 0.01rem solid #D7DCE0;
						margin: 0.15rem;
						padding: 0.3rem 0 0 0;
						border-left: none;
						.validation_information_item{
						
						}
					}
				}
			}
		}
	}
	@media screen and (max-width: 910px){
		.validation_information_container{
			margin-top: 0.2rem;
		}
	}
	@media screen and (max-width: 768px){
		.validation_information_container{
			margin-top: 0.2rem;
			.validation_information_wrap{
				margin: 0 0.1rem;
				.validation_information_content{
					.validation_information_header_content{
						align-items: center;
					}
				}
			}
		}
	}
</style>
