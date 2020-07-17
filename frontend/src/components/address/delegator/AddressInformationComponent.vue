<template>
	<div class="address_information_component_container">
		<div class="address_information_component_wrap">
			<div class="address_information_component_content">
				<div class="address_information_asset_content">
					<div class="address_information_asset_total_content">
						<img src="../../../assets/iris_token_logo.svg" alt="">
						<ul class="address_information_content">
							<li class="address_information_item">
								<span class="address_information_label">Address:</span>
								<p>
									<span>{{address}}<m-clip :text="address" style="margin-left: 0.09rem"></m-clip><span class="profiler_content" v-if="isProfiler">Profiler</span></span>
								</p>
							</li>
							<li class="address_information_item">
								<span class="address_information_label">Token:</span>
								<span class="address_information_value">IRIS</span>
							</li>
							<li class="address_information_item">
								<span class="address_information_label">Total Amount:</span>
								<span class="address_information_value">{{totalAmount || 0}}</span>
							</li>
						</ul>
						<!--<div class="address_information_asset_logo">-->
							<!---->
							<!--<span>IRIS</span>-->
						<!--</div>-->
						<!--<p class="address_information_title">Total Amount</p>-->
						<!--<p class="address_information_total_amount_content">{{totalAmount}}</p>-->
					</div>
					<ul class="address_information_asset_constitute_content" v-show="flShowAssetInfo(assetConstitute)">
						<li class="address_information_asset_constitute_item" v-for="(item,index) in assetConstitute" :key="index">
							<span :style="{background:item.color}" ></span>
							<span >{{item.label}}</span>
							<span >{{item.value}}</span>
						</li>
					</ul>
					<div class="address_information_asset_pie_content"  v-show="flShowAssetInfo(assetConstitute)">
						<address-information-pie :echartData="assetConstitute"></address-information-pie>
					</div>
				</div>
				<div class="address_information_asset_list_container" v-if="otherTokenList.length !== 0">
					<div class="address_information_asset_header_content">
						<span>Token</span>
						<span>Balance</span>
					</div>
					<ul class="address_information_asset_list_content">
						<li class="address_information_asset_list_item" v-for="(item,index) in otherTokenList" :key="index">
							<div class="address_information_token_img_content">
								<span class="iconfont iconqitabizhongwuiconshishiyongdemorenicon"></span>
								<span class="address_information_token_name">{{item.token}}</span>
							</div>
							<div class="address_information_list_item_value">{{item.balance}}</div>
						</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import MClip from "../../commontables/MClip";
	import Tools from "../../../util/Tools"
	import AddressInformationPie from "./AddressInformationPie";
	import moveDecimal from 'move-decimal-point'
	export default {
		name: "AddressInformationComponent",
		components: {AddressInformationPie, MClip},
		props:{
			address:{
				type:String
			},
			data:{
				type: Array
			},
			isProfiler:{
				type:Boolean,
			},
		},
		data(){
			return {
				assetInformation:'',
				assetConstitute:[
					{
						label:'Balance',
						value:'',
						color:'var(--bgColor)'
					},
					{
						label:'Delegated',
						value:'',
						color:'#FFA300'
					},
					{
						label:'UnBonding',
						value:'',
						color:'#67E523'
					},
					{
						label:'Rewards',
						value:'',
						color:'#8E66FF'
					},
					
				],
				totalAmount: '',
				balance:'',
				delegated:'',
				unbonding:'',
				rewards:'',
				otherTokenList:[]
			}
		},
		watch:{
			data(){
				this.assetInformation = this.data;
				this.formatAssetInformation(this.assetInformation)
			}
		},
		mounted () {
			this.assetInformation = this.data;
		},
		methods:{
			flShowAssetInfo(data){
				let res = data.every( (item) => {
					 return item.value !== ''
				 });
				return res
			},
			formatAssetInformation(assetInformation){
				assetInformation.forEach( item => {
					if(item && item.token === 'IRIS'){
						this.totalAmount = item.totalAmount;
						this.assetConstitute.forEach( res => {
							 if(res.label === "UnBonding"){
								res.value = item['unBonding'] || "--";
								res.numberValue = item['unBonding'] ? item['unBonding'].replace(/[^\d.]/g,"") : 0;
								res.percent = this.formatDecimalNumberToFixedNumber(item.totalAmount.replace(/[^\d.]/g,""),res.numberValue)
							}else {
								res.value = item[Tools.firstWordLowerCase(res.label)] && item[Tools.firstWordLowerCase(res.label)] !== 0 ? item[Tools.firstWordLowerCase(res.label)] : "--";
								res.numberValue = item[Tools.firstWordLowerCase(res.label)] ?
									item[Tools.firstWordLowerCase(res.label)].replace(/[^\d.]/g,"") : 0;
								res.percent = this.formatDecimalNumberToFixedNumber(item.totalAmount.replace(/[^\d.]/g,""),res.numberValue)
							}
						})
					}
				});
				this.otherTokenList = assetInformation.filter((item) => {
					return item.token !== 'IRIS'
				})
			},
			formatDecimalNumberToFixedNumber(total,data) {
				let percentNumber = (Number(data) / Number(total)).toString();
				let num;
				if(percentNumber !== 'Infinity'){
					num = Tools.subStrings(moveDecimal(Tools.FormatScientificNotationToNumber(percentNumber),2),2) ;
				}else {
					//数字太小赋值为0.00
					num = '0.00'
				}
				return num;
			}
		}
	}
</script>

<style scoped lang="scss">
.address_information_component_container{
	padding-top: 0.74rem;
	.address_information_component_wrap{
		max-width: 12.8rem;
		margin: 0 auto;
		.address_information_component_content{
			border: 0.01rem solid #E7E9EB;
			.address_information_component_title{
				color:var(--titleColor);
				font-size: 0.18rem;
				line-height: 0.21rem;
				margin-bottom: 0.1rem;
				.address{
					margin-right: 0.1rem;
				}
			}
			.address_information_asset_content{
				max-width: 12.8rem;
				box-sizing: border-box;
				background: #fff;
				padding: 0.3rem 0.2rem 0.2rem 0.2rem;
				display: grid;
				grid-template-columns: repeat(3,50% 30% 20%);
				.address_information_asset_total_content{
					display: grid;
					grid-template-columns: repeat(1,0.5rem auto);
					margin-right: 0.15rem;
					img{
						width: 0.4rem;
					}
					.address_information_content{
						.address_information_item{
							display: grid;
							grid-template-columns: repeat(1,1.1rem auto);
							margin-top: 0.14rem;
							.address_information_label{
								width: 1.1rem;
								color: var(--contentColor);
								font-size: 0.14rem;
								line-height: 0.16rem;
							}
							p{
								display: flex;
								color: var(--titleColor);
								font-size: 0.14rem;
								line-height: 0.2rem;
								padding-right: 0.01rem;
									.profiler_content{
										padding: 0 0.12rem;
										border-radius: 0.08rem;
										margin-left: 0.08rem;
										background: var(--bgColor);
										color: #fff;
										font-size: 0.14rem;
										font-weight: lighter;
										word-break: normal;
									}
							}
							.address_information_value{
								flex: 1;
								display: flex;
								color: var(--titleColor);
								font-size: 0.14rem;
								line-height: 0.16rem;
							}
						}
						.address_information_item:first-child{
							margin-top: 0;
							p{
								font-weight: bold;
							}
						}
					}
					.address_information_title{
						font-size: 0.14rem;
						line-height: 0.16rem;
						color: var(--contentColor);
						margin-bottom: 0.1rem;
					}
					.address_information_total_amount_content{
						font-size: 0.16rem;
						line-height: 0.18rem;
						font-weight: bold;
						color: var(--titleColor);
					}
				}
				.address_information_asset_constitute_content{
					display: grid;
					grid-template-columns: repeat(1,auto);
					grid-template-rows: repeat(4,auto);
					.address_information_asset_constitute_item{
						display: flex;
						align-items: center;
						margin-bottom: 0.1rem;
						span:nth-of-type(1){
							display: inline-block;
							width: 0.1rem;
							height: 0.1rem;
							margin-right: 0.09rem;
						}
						span:nth-of-type(2){
							font-size: 0.14rem;
							line-height: 0.16rem;
							width: 1rem;
						}
						span:nth-of-type(3){
							font-size: 0.14rem;
							line-height: 0.16rem;
						}
					}
				}
				.address_information_asset_pir_content{
				
				}
			}
			.address_information_asset_list_container{
				box-sizing: border-box;
				padding: 0 0.2rem;
				background: #fff;
				.address_information_asset_header_content{
					box-sizing: border-box;
					border-top: 0.01rem solid #D7DCE0;
					padding-top: 0.19rem;
					margin-bottom: 0.1rem;
					display: grid;
					grid-template-rows: repeat(1,100%);
					grid-template-columns: repeat(2,50%);
					span{
						color: var(--contentColor);
					}
					span:nth-of-type(1){
						padding-left: 0.2rem;
					}
				}
				.address_information_asset_list_content{
					display: flex;
					flex-direction: column;
					.address_information_asset_list_item{
						display: grid;
						grid-template-rows: repeat(1,100%);
						grid-template-columns: repeat(2,50%);
						background: #F0F3FB;
						box-sizing: border-box;
						padding: 0.08rem 0 0.08rem 0;
						margin-bottom: 0.05rem;
						.address_information_token_img_content{
							padding-left: 0.2rem;
							.iconqitabizhongwuiconshishiyongdemorenicon{
								font-size: 0.18rem;
								color: #B6BAD2;
								margin-right: 0.1rem;
							}
							.address_information_token_name{
								font-size: 0.14rem;
								line-height: 0.16rem;
								color: var(--titleColor);
							}
						}
						.address_information_list_item_value{
							font-size: 0.14rem;
							line-height: 0.16rem;
							color: var(--titleColor);
						}
					}
					.address_information_asset_list_item:last-child{
						margin-bottom: 0.18rem;
					}
				}
			}
		}
	}
}
	@media screen and (max-width: 1030px){
		
		.address_information_component_container{
			padding-top: 0.54rem;
			.address_information_component_wrap{
				margin: 0.2rem 0.1rem 0 0.1rem;
				.address_information_component_content{
					.address_information_asset_content{
						width: 100%;
						padding: 0.1rem;
						grid-template-columns: repeat(1,auto);
						.address_information_asset_total_content{
							width: 100%;
							grid-template-columns: repeat(1,auto);
							.address_information_content{
								margin-top: 0.15rem;
								.address_information_item{
									grid-template-columns: repeat(1,1.1rem auto);
									p{
										word-wrap: break-word;
										word-break: break-all;
									}
								}
							}
						}
						.address_information_asset_constitute_content{
							margin-top: 0.3rem;
							border-top: 0.01rem solid #D7DCE0;
							width: 100%;
							.address_information_asset_constitute_item:first-child{
								margin-top: 0.3rem;
							}
						}
					}
					.address_information_asset_list_container{
						padding: 0 0.1rem;
					}
				}
			}
		}
	}
	@media screen and (max-width: 910px){
		.address_information_component_container{
			padding-top: 0;
		}
	}
</style>
