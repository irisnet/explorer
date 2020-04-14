<template>
	<div class="statistical_bar_container">
		<div class="statistical_bar_wrap">
			<div class="statistical_bar_content">
				<div class="statistical_validator_content">
					<div class="statistical_validator_top_content">
						<p class="statistical_validator_header">
							<span class="iconfont iconBlocks"></span>
							<span class="statistical_validator_header_label">Block Height</span>
						</p>
						<p class="statistical_current_block skip_route">
							<router-link :to="`/block/${currentBlockHeight}`">{{currentBlockHeight}}</router-link>
						</p>
					</div>
					<div  class="statistical_validator_bottom_content">
						<div class="statistical_img_content">
							<router-link :to="addressRoute(proposerAddress)">
								<div class="statistical_validator_header_img_content">
									<span >{{validatorHeaderImgSrc}}</span>
									<img v-show="validatorHeaderImgHref" :src="validatorHeaderImgHref"  @error="imgLoadError()">
								</div>
							</router-link>
						</div>
						<p class="statistical_moniker_content skip_route">
							<router-link :to="addressRoute(proposerAddress)">{{moniker}}</router-link>
						</p>
					</div>
				</div>
				<div class="statistical_right_content">
					<ul class="statistical_list_content">
						<li class="statistical_item_content" v-for="(item,index) in navigationArray" :key="index">
							<p class="statistical_header_content">
								<span class="statistical_icon" :class="item.iconClass"></span>
								<span class="statistical_content">{{item.label}}</span>
								<el-tooltip v-if="item.flShowTip" :content="'Proportion of community bonded token in total circulation.'">
									<span class="iconfont icontishi"></span>
								</el-tooltip>
							</p>
							<p class="statistical_center_content skip_route">
								<router-link v-if="item.isLink" :to="`/txs`">{{item.value}}</router-link>
								<span v-else>{{item.value}}</span>
							</p>
							<p class="statistical_footer_content">
								{{item.footerLabel}}
							</p>
						</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import Service  from "../../service"
	import Tools from "../../util/Tools"
	export default {
		name: "StatisticalBar",
		data () {
			return {
				navigationArray:[
					{
						iconClass:'iconfont iconTransactions',
						label:'Transactions',
						footerLabel:'--',
						value:'--',
						flShowTip:false,
						isLink: true,
					},
					{
						iconClass:'iconfont iconVotingPower',
						label:'Voting Power',
						footerLabel:'--',
						value:'--',
						flShowTip:false,
						isLink: false,
					},
					{
						iconClass:'iconfont iconAvgBlockTime',
						label:'Avg Block Time',
						footerLabel:'Last 100 Blocks',
						value:'--',
						flShowTip:false,
						isLink: false,
					},
					{
						iconClass:'iconfont iconBondedTokens',
						label:'Bonded Tokens',
						footerLabel:'--',
						value:'--',
						flShowTip:false,
						isLink: false,
					},
					// {
					// 	iconClass:'iconfont iconCirculationBondedTokens',
					// 	label:'Circulation Bonded Tokens',
					// 	footerLabel:'--',
					// 	value:'--',
					// 	flShowTip:true,
					// 	isLink: false,
					// }
				],
				moniker:'--',
				proposerAddress:'--',
				diffSeconds:'--',
				currentBlockHeight:'--',
				validatorHeaderImgSrc:'',
				validatorHeaderImgHref:'',
				timer:null,
			}
		},
		mounted(){
			this.getNavigation();
			clearInterval(this.timer);
			setInterval(() =>{
				this.getNavigation()
			},5000)
		},
		methods:{
			imgLoadError(){
				this.validatorHeaderImgHref = ""
			},
			getNavigation(){
				Service.commonInterface({navigation:{}},(res) => {
					try {
						if(res){
							//先通过正则剔除符号空格及表情，只保留数字字母汉字
							let regex =  /[^\w\u4e00-\u9fa50-9a-zA-Z]/g;
							let replaceMoniker = res.moniker.replace(regex,'');
							let reservedStringLength = 12;
							this.moniker = Tools.formatString(res.moniker,reservedStringLength,'...');
							this.validatorHeaderImgHref = res.validator_icon ? res.validator_icon : replaceMoniker ? '' : require('../../assets/default_validator_icon.svg');
							this.validatorHeaderImgSrc = replaceMoniker ? Tools.firstWordUpperCase(replaceMoniker.match(/^[0-9a-zA-Z\u4E00-\u9FA5]/g)[0]) : '';
							this.proposerAddress = res.operator_addr;
							this.diffSeconds = Number(res.avg_block_time);
							this.currentBlockHeight = res.block_height;
							this.navigationArray[0].value = this.formatTransactions(res.total_txs);
							this.navigationArray[0].footerLabel = Tools.format2UTC(res.block_time);
							this.navigationArray[1].value = `${Number(res.voting_ratio * 100).toFixed(2)} %`;
							this.navigationArray[1].footerLabel = `${res.vote_val_num} / ${res.active_val_num} Validators`;
							this.navigationArray[2].value = `${Number(res.avg_block_time).toFixed(2)} s`;
							this.navigationArray[3].value = `${(res.bonded_ratio * 100).toFixed(2)} %`;
							this.navigationArray[3].footerLabel = this.formatBondedTokens(res.bonded_tokens,res.total_supply);
							// this.navigationArray[4].value = `${((Number(res.bonded_tokens) - Number(res.foundation_bonded)) / Number(res.circulation)*100).toFixed(2)} %`
							// this.navigationArray[4].footerLabel = this.formatBondedTokens((Number(res.bonded_tokens) - Number(res.foundation_bonded)),res.circulation);
						}
					}catch (e) {
						console.error(e)
					}
				});
			},
			formatTransactions(totalTxs){
				let num, thousand = 1000,million = 1000000,billion = 1000000000;
				if(totalTxs > billion){
					num = `${(totalTxs/billion).toFixed(2)} B`;
				}else if(totalTxs > million){
					num = `${(totalTxs/million).toFixed(2)} M`;
				}else if(totalTxs > thousand){
					num = `${(totalTxs/thousand).toFixed(2)} K`;
				}else {
					num = totalTxs
				}
				return num
			},
			formatBondedTokens(bondedTokens,totalTokens){
				let tokens,allTokens,thousand = 1000,million = 1000000,billion = 1000000000;
				if(bondedTokens >= billion){
					tokens = `${(Number(bondedTokens) / billion).toFixed(2)}B`
				}else if(bondedTokens >= million){
					tokens = `${(Number(bondedTokens) / million).toFixed(2)}M`
				}else if(bondedTokens >= thousand){
					tokens = `${(Number(bondedTokens) / thousand).toFixed(2)}k`
				}else {
					tokens = `${Number(bondedTokens).toFixed(2)}`
				}
				
				if(totalTokens >= billion){
					allTokens = `${(Number(totalTokens) / billion).toFixed(2)}B`
				}else if(totalTokens >= million){
					allTokens = `${(Number(totalTokens) / million).toFixed(2)}M`
				}else if(totalTokens >= thousand){
					allTokens = `${(Number(totalTokens) / thousand).toFixed(2)}k`
				}else {
					allTokens = `${Number(totalTokens).toFixed(2)}`
				}
				return `${tokens} / ${allTokens}`
			},
		},
		destroyed () {
			clearInterval(this.timer);
		}
	}
</script>

<style scoped lang="scss">
	.statistical_bar_container{
		width: 100%;
		.statistical_bar_wrap{
			max-width: 12.4rem;
			margin: 0.2rem auto 0 auto;
			.statistical_bar_content{
				box-sizing: border-box;
				padding: 0.2rem;
				background: #fff;
				display: grid;
				grid-template-columns: 30% 70%;
				border:0.01rem solid #d6d9e0;
				.statistical_validator_content{
					display: flex;
					flex-direction: column;
					align-items: center;
					.statistical_validator_top_content{
						display: flex;
						flex-direction: column;
						align-items: center;
						.statistical_validator_header{
							margin-bottom: 0.2rem;
							margin-top: 0.15rem;
							.iconBlocks{
								color: var(--bgColor);
								margin-right: 0.08rem;
							}
							.statistical_validator_header_label{
								color:var(--contentColor);
								font-size: 0.14rem;
								line-height: 0.16rem;
							}
						}
						.statistical_current_block{
							margin-bottom: 0.2rem;
							a{
								font-size: 0.2rem;
							}
							font-weight: bold;
						}
					}
					.statistical_validator_bottom_content{
						display: flex;
						flex-direction: column;
						align-items: center;
						.statistical_img_content{
							width: 1.2rem;
							height: 1.2rem;
							margin-bottom: 0.27rem;
							.statistical_validator_header_img_content{
								width: 100%;
								height: 100%;
								display: flex;
								align-items: center;
								justify-content: center;
								border-radius: 0.6rem;
								background: #E0E8FF;
								overflow: hidden;
								position: relative;
								img{
									position: absolute;
									width: 100%;
								}
								span{
									font-size: 0.52rem;
								}
							}
						}
					}
				}
				.statistical_right_content{
					.statistical_list_content{
						display: grid;
						grid-template-columns: repeat(2,1fr);
						grid-template-rows: repeat(2,auto);
						grid-row-gap: 0.2rem;
						grid-column-gap: 0.2rem;
						.statistical_item_content{
							box-sizing: border-box;
							padding: 0.14rem;
							border-radius: 0.04rem;
							border: 0.01rem solid rgba(231,233,235,1);
							.statistical_header_content{
								.statistical_icon{
									color: var(--bgColor);
									margin-right: 0.08rem;
								}
								.statistical_content{
									font-size: 0.14rem;
									color: var(--contentColor);
									margin-right: 0.08rem;
								}
							}
							.statistical_center_content{
								margin-top: 0.36rem;
								font-size: 0.2rem;
								color: var(--titleColor);
								a{
									font-size: 0.2rem;
									line-height: 0.23rem;
								}
							}
							.statistical_footer_content{
								margin-top: 0.1rem;
								color: var(--contentColor);
								font-size: 0.14rem;
								line-height: 0.16rem;
							}
						}
					}
				}
			}
		}
	}
	@media screen and (max-width: 1280px){
		.statistical_bar_container{
			.statistical_bar_wrap{
				margin: 0.2rem 0.2rem 0 0.2rem;
			}
		}
	}
	@media screen and (max-width: 1000px){
		.statistical_bar_container{
			.statistical_bar_wrap{
				margin: 0.2rem 0.1rem 0 0.1rem;
				.statistical_bar_content{
					margin: 0 0.1rem;
					padding: 0.1rem;
					grid-template-columns: repeat(1,auto);
					.statistical_validator_content{
						flex-direction: row;
						justify-content: space-between;
						align-items: center;
						
						.statistical_validator_top_content{
							.statistical_validator_header{
								margin-bottom: 0.15rem;
							}
						}
						.statistical_validator_bottom_content{
							display: flex;
							flex-direction: column;
							align-items: flex-end;
							.statistical_img_content{
								width: 0.3rem;
								height: 0.3rem;
								border-radius: 0.15rem;
								display: flex;
								align-items: center;
								justify-content: center;
								margin-bottom: 0.07rem;
								a{
									width: 100%;
									height: 100%;
									.statistical_validator_header_img_content{
										span{
											font-size: 0.1rem;
										}
									}
								}
							}
						}
					}
					.statistical_right_content{
						.statistical_list_content{
							grid-template-columns:repeat(1,auto);
							grid-template-rows: repeat(1,auto);
						}
					}
				}
			}
		}
	}
	@media screen and (max-width: 910px){
		.statistical_bar_container{
			.statistical_bar_wrap{
				margin: 0.2rem 0 0 0;
			}
		}
	}
</style>
