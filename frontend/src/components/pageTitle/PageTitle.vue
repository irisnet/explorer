<template>
	<div class="page_title_container"  style="height: 0.54rem;background: #F5F7FD">
		<div class="page_title_container_content_wrap">
			<div :class="flShowMobileStyle ? 'mobile_page_title_wrap ' :'page_title_content_wrap'">
				<div class="mobile_page_title_content_wrap">
					<div class="page_title">{{title}}</div>
					<span class="page_link" v-if="flShowPageLink">|</span>
				</div>
				<div class="page_title_content" v-if="reversal">
					<div class="title_content">{{content}} <m-clip :text="copyText" v-if="flShowClip"></m-clip></div>
					<span class="number_content" v-if="!flLink">{{number}}</span>
					<router-link :to="`${href}${number}`" v-else>
						<span class="number_content">{{number}}</span>
					</router-link>
					
					<slot></slot>
				</div>
				<div class="page_title_content" v-if="!reversal">
					<span class="number_content" style="color: #515A6E;padding-right: 0.03rem" v-if="!flLink">{{number}}</span>
					<span class="title_content"> {{content}}</span>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import MClip from "../commonComponents/MClip";
	export default {
		name: "PageTitle",
		components: {MClip},
		props:{
			title:{
				type:String
			},
			content:{
				type: String
			},
			number:{
				type:Number
			},
			href:{
				type:String,
			},
			flLink:{
				type: Boolean,
				default: false,
			},
			flShowPageLink:{
				type: Boolean,
				default: true
			},
			copyText:{
				type:String
			},
			flShowClip:{
				type:Boolean,
				default:false
			},
			reversal:{
				type:Boolean,
				default: true
			},
			flShowMobileStyle:{
				type:Boolean,
				default: false
			}
		}
	}
</script>

<style scoped lang="scss">
	.page_title_container{
		width: 100%;
		position: fixed;
		height: 0.54rem;
		line-height: 0.54rem;
		z-index: 10;
		.page_title_container_content_wrap{
			height: 0.5rem;
			line-height: 0.5rem;
			width: 100%;
			box-shadow:0 0.02rem 0.04rem 0 rgba(0,0,0,0.1);
			.page_title_content_wrap{
				max-width: 12.8rem;
				margin: 0 auto;
				display: flex;
				justify-content: flex-start;
				background: #F5F7FD;
				.mobile_page_title_content_wrap{
					display: flex;
					.page_link{
						padding: 0 0.2rem;
					}
					.page_title{
						height: 0.5rem;
						line-height: 0.5rem;
						font-size: 0.14rem;
						font-weight: bold;
						color: rgba(120,124,153,1);
						white-space: nowrap;
					}
				}
				.page_title_wrap{
					.page_link{
						padding: 0 0.2rem;
					}
					.page_title{
						height: 0.5rem;
						line-height: 0.5rem;
						font-size: 0.14rem;
						font-weight: bold;
						color: rgba(120,124,153,1);
						white-space: nowrap;
					}
				}
				.page_title_content{
					display: flex;
					align-items: center;
					.title_content{
						color: #515A6E;
						font-size: 0.14rem;
						font-weight: bold;
						padding-right: 0.05rem;
					}
					.number_content{
						color: var(--bgColor);
						font-size: 0.14rem;
						font-weight: bold;
					}
					.tooltip_icon {
						width: 0.16rem;
						height: 0.16rem;
						margin-left: 0.1rem;
						display: flex;
						justify-content: center;
						align-items: center;
						border: 1px solid var(--bgColor);
						border-radius: 50%;
						font-size: 0.1rem;
						color: var(--bgColor);
						position: relative;
						cursor: pointer;
						&:hover .tooltip_span {
							display: block;
						}
						.tooltip_span {
							width: 2.4rem;
							display: none;
							position: absolute;
							z-index: 5;
							top:0.2rem;
							left: 0.25rem;
							transform: translateY(-50%);
							color: #ffffff;
							background-color: #000000;
							border-radius: 0.04rem;
							word-wrap: break-word;
							white-space: normal;
							line-height: 16px;
							div {
								padding: 8px 15px;
							}
							&::after {
								width: 0;
								height: 0;
								border: 0.06rem solid transparent;
								content: "";
								display: block;
								position: absolute;
								border-top-color: #000000;
								transform: rotate(90deg);
								top: 25%;
								margin-left: -0.12rem;
							}
						}
					}
					.address_information_address_status_profiler{
						background: var(--bgColor);
						font-size: 0.12rem;
						color: #fff;
						padding: 0.02rem 0.14rem;
						border-radius: 0.22rem;
					}
				}
			}
			.mobile_page_title_wrap{
				max-width: 12.8rem;
				margin: 0 auto;
				display: flex;
				justify-content: flex-start;
				background: #F5F7FD;
				.mobile_page_title_content_wrap{
					display: flex;
					.page_link{
						padding: 0 0.2rem;
					}
					.page_title{
						height: 0.5rem;
						line-height: 0.5rem;
						font-size: 0.14rem;
						font-weight: bold;
						color: rgba(120,124,153,1);
						white-space: nowrap;
					}
				}
				.page_title_wrap{
					.page_link{
						padding: 0 0.2rem;
					}
					.page_title{
						height: 0.5rem;
						line-height: 0.5rem;
						font-size: 0.14rem;
						font-weight: bold;
						color: rgba(120,124,153,1);
						white-space: nowrap;
					}
				}
				.page_title_content{
					display: flex;
					align-items: center;
					.title_content{
						color: #515A6E;
						font-size: 0.14rem;
						font-weight: bold;
						padding-right: 0.05rem;
					}
					.number_content{
						color: var(--bgColor);
						font-size: 0.14rem;
						font-weight: bold;
					}
					.tooltip_icon {
						width: 0.16rem;
						height: 0.16rem;
						margin-left: 0.1rem;
						display: flex;
						justify-content: center;
						align-items: center;
						border: 1px solid var(--bgColor);
						border-radius: 50%;
						font-size: 0.1rem;
						color: var(--bgColor);
						position: relative;
						cursor: pointer;
						&:hover .tooltip_span {
							display: block;
						}
						.tooltip_span {
							width: 2.4rem;
							display: none;
							position: absolute;
							z-index: 5;
							top:0.2rem;
							left: 0.25rem;
							transform: translateY(-50%);
							color: #ffffff;
							background-color: #000000;
							border-radius: 0.04rem;
							word-wrap: break-word;
							white-space: normal;
							line-height: 16px;
							div {
								padding: 8px 15px;
							}
							&::after {
								width: 0;
								height: 0;
								border: 0.06rem solid transparent;
								content: "";
								display: block;
								position: absolute;
								border-top-color: #000000;
								transform: rotate(90deg);
								top: 25%;
								margin-left: -0.12rem;
							}
						}
					}
					.address_information_address_status_profiler{
						background: var(--bgColor);
						font-size: 0.12rem;
						color: #fff;
						padding: 0.02rem 0.14rem;
						border-radius: 0.22rem;
						line-height: 0.22rem;
						height: 0.22rem;
						display: inline-block;
						min-width: 0.4rem;
					}
				}
			}
			
		}
		
	}
	@media screen and (max-width:1280px){
		.page_title_container{
			.page_title_container_content_wrap{
				.page_title_content_wrap{
					box-sizing: border-box;
					padding: 0 0.1rem;
				}
			}
		}
	}
	@media screen and (max-width: 910px){
		.page_title_container{
			position: static;
			.page_title_container_content_wrap{
				height: auto;
				.mobile_page_title_wrap{
					box-sizing: border-box;
					padding: 0 0.1rem;
					flex-direction: column;
					.mobile_page_title_content_wrap{
						.page_title{
							height: 0.3rem;
							line-height: 0.3rem;
						}
						.page_link{
							line-height: 0.3rem;
						}
					}
					.page_title_content{
						word-break: break-all;
						line-height: 0.16rem;
						padding-bottom: 0.05rem;
					}
				}
				.page_title_content_wrap{
					.page_title_content{
						.tooltip_icon {
							width: 0.16rem;
							height: 0.16rem;
							margin-left: 0.1rem;
							display: flex;
							justify-content: center;
							align-items: center;
							border: 1px solid var(--bgColor);
							border-radius: 50%;
							font-size: 0.1rem;
							color: var(--bgColor);
							position: relative;
							cursor: pointer;
							&:hover .tooltip_span {
								display: block;
							}
							.tooltip_span {
								width: 2.4rem;
								display: none;
								position: absolute;
								z-index: 5;
								top:0.6rem;
								left: auto;
								right: -0.1rem;
								transform: translateY(-50%);
								color: #ffffff;
								background-color: #000000;
								border-radius: 0.04rem;
								word-wrap: break-word;
								white-space: normal;
								line-height: 16px;
								div {
									padding: 8px 15px;
								}
								&::after {
									width: 0;
									height: 0;
									border: 0.06rem solid transparent;
									content: "";
									display: block;
									position: absolute;
									border-right-color: #000000;
									transform: rotate(90deg);
									right: 5%;
									top: -18%;
									margin-left: -0.12rem;
								}
							}
						}
					}
				}
			}
		}
	}
</style>
