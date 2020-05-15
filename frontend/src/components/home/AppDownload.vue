<template>
	<div class="app_download_container">
		<div class="app_download_wrap">
			<h2 class="app_download_title">Rainbow Wallet</h2>
			<div class="app_download_content_wrap">
				<div class="app_download_goz_content">
					<h4 class="app_download_goz_title">IBC-GoZ Version</h4>
					<p class="app_download_goz_doc">Come experience the future IBC cross-chain wallet</p>
					<div class="app_download_link_content">
<!--						<p class="app_download_goz_link_title">COMING SOON.</p>-->
						<a @click="$uMeng.push('RainbowIBC-GoZ','下载','安卓','统计从GoZ浏览器下载的安卓')" href="https://rb-app-stage.oss-cn-shanghai.aliyuncs.com/rainbow-latest.apk"  target="_blank"  v-if="isShowAndroid || !isMobile">
							<div class="app_download_android_content">
								<img src="../../assets/android_apk.png" alt="">
							</div>
						</a>
						<a @click="$uMeng.push('RainbowIBC-GoZ','下载','iOS','统计从GoZ浏览器下载的iOS')" href="https://testflight.apple.com/join/4K1dJKxe" target="_blank"  v-if="isShowIphone || !isMobile">
							<div class="app_download_ios_content">
								<img src="../../assets/ios_testflight.png" alt="">
							</div>
						</a>
						<div class="app_download_qr_content" v-if="!isMobile">
							<img src="../../assets/scan_icon.png" alt="">
							<div class="qr_content">
								<img src="../../assets/goz_rainbow_qr.png" alt="">
								<span class="tooltip_content">Scan to download</span>
							</div>
						</div>
<!--						<p class="app_download_goz_link_content">Expected to be online on May 1</p>-->
					</div>
				</div>
				<div class="app_download_mainnet_content">
					<h4 class="app_download_mainnet_title">Mainnet Version</h4>
					<p class="app_download_mainnet_doc">If you want to experience delegate/Transfer functions from mainnet version in Cosmos and IRIS Hub</p>
					<div class="app_download_img_content">
						<a @click="$uMeng.push('RainbowMainnet_Android','click')" href="https://play.google.com/store/apps/details?id=rainbow.wallet"  target="_blank"  v-if="isShowAndroid || !isMobile">
							<div class="app_download_android_content">
								<img src="../../assets/rainbow_google_plays.png" alt="">
							</div>
						</a>
						<a @click="$uMeng.push('RainbowMainnet_iOS','click')" href="https://apps.apple.com/us/app/rainbow-irisnet-cosmos/id1466725641" target="_blank"  v-if="isShowIphone || !isMobile">
							<div class="app_download_ios_content">
								<img src="../../assets/app_store.png" alt="">
							</div>
						</a>
						<div class="app_download_qr_content" v-if="!isMobile">
							<img src="../../assets/scan_icon.png" alt="">
							<div class="qr_content">
								<img src="../../assets/rainbow_qr.png" alt="">
								<span class="tooltip_content">Scan to download</span>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		name: "AppDownload",
		data() {
			return {
				isShowAndroid: true,
				isShowIphone: true,
				isMobile: false
			}
		},
		mounted () {
			let ua = navigator.userAgent;
			let ipad = ua.match(/(iPad).*OS\s([\d_]+)/),
				isIphone = ipad || ua.match(/(iPhone\sOS)\s([\d_]+)/),
				isAndroid = ua.match(/(Android)\s+([\d.]+)/),
				isMobile = isIphone || isAndroid;
			if(isMobile){
				this.isMobile = true
			}
			if(isIphone){
				this.isShowIphone = true;
				this.isShowAndroid = false;
			}else if(isAndroid){
				this.isShowAndroid = true;
				this.isShowIphone = false;
			}
		}
	}
</script>

<style scoped lang="scss">
	.app_download_container{
		width: 100%;
		padding-top: 0.4rem;
		.app_download_wrap{
			width: 100%;
			.app_download_title{
				font-size: 0.18rem;
				line-height: 0.21rem;
				color: #fff;
				padding-left: 0.2rem;
				font-weight: lighter;
			}
			.app_download_content_wrap{
				display: grid;
				grid-template-columns: repeat(2,50%);
				grid-template-rows: repeat(1,auto);
				margin-top: 0.15rem;
				.app_download_goz_content{
					margin-right: 0.01rem;
					box-sizing: border-box;
					padding: 0.2rem 0 0.29rem 0.2rem;
					border-radius: 0.04rem;
					background: rgba(45, 50, 90, 0.64);
					.app_download_goz_title{
						color: #fff;
						font-size: 0.16rem;
						line-height: 0.18rem;
						font-weight: 300;
					}
					.app_download_goz_doc{
						margin-top: 0.06rem;
						color: rgba(134, 143, 211, 1);
						font-size: 0.14rem;
						line-height: 0.16rem;
					}
					.app_download_link_content{
						margin-top: 0.29rem;
						.app_download_goz_link_title{
							color: rgba(55, 124, 248, 1);
							font-size: 0.16rem;
							line-height: 0.18rem;
						}
						.app_download_goz_link_content{
							margin-top: 0.06rem;
							font-size: 0.16rem;
							line-height: 0.18rem;
							color: rgba(55, 124, 248, 1);
						}
						
						display: flex;
						margin-top: 0.2rem;
						a{
							.app_download_android_content{
								margin-right: 0.2rem;
								background: #fff;
								width: 1.6rem;
								height: 0.35rem;
								box-sizing: border-box;
								padding: 0.07rem 0;
								display: flex;
								align-content: center;
								justify-content: center;
								border-radius: 0.04rem;
								img{
									width: 0.91rem;
									height:0.21rem;
								}
							}
						}
						a{
							.app_download_ios_content{
								background: #fff;
								width: 1.6rem;
								height: 0.35rem;
								box-sizing: border-box;
								padding: 0.07rem 0;
								display: flex;
								align-content: center;
								justify-content: center;
								border-radius: 0.04rem;
								img{
									width: 0.82rem;
									height:0.21rem;
								}
							}
						}
						.app_download_qr_content{
							cursor: pointer;
							width: 0.35rem;
							height: 0.35rem;
							background: #fff;
							margin-left: 0.2rem;
							border-radius: 0.04rem;
							box-sizing: border-box;
							padding: 0.05rem;
							position: relative;
							z-index: 5;
							img{
								width: 100%;
								height: 100%;
							}
							&:hover{
								.qr_content{
									display: block;
								}
							}
							.qr_content{
								display: none;
								background: #fff;
								margin-top: 0.1rem;
								position: absolute;
								width: 1.4rem;
								height: 1.55rem;
								border-radius: 0.04rem;
								margin-left: 50%;
								transform: translateX(-50%);
								text-align: center;
								img{
									margin-top: 0.1rem;
									width: 1.2rem;
									height:1.2rem;
								}
								.tooltip_content{
									font-size: 0.14rem;
								}
							}
							.qr_content::before {
								width: 0;
								content: '';
								display: block;
								border: 6px solid transparent;
								border-bottom-color: #ffffff;
								margin-top: -12px;
								margin-left: 48%;
								transform: translateX(-50%);
							}
						}
					}
				}
				.app_download_mainnet_content{
					margin-left: 0.01rem;
					box-sizing: border-box;
					padding: 0.2rem  0.2rem 0.2rem 0.22rem;
					background: rgba(45, 50, 90, 0.64);
					border-radius: 0.04rem;
					.app_download_mainnet_title{
						font-size: 0.16rem;
						color: #fff;
						line-height: 0.18rem;
						font-weight: lighter;
					}
					.app_download_mainnet_doc{
						margin-top: 0.06rem;
						color: rgba(134, 143, 211, 1);
						font-size: 0.14rem;
						line-height: 0.16rem;
					}
					.app_download_img_content{
						display: flex;
						margin-top: 0.2rem;
						a{
							.app_download_android_content{
							margin-right: 0.2rem;
							background: #fff;
							width: 1.6rem;
							height: 0.35rem;
							box-sizing: border-box;
							padding: 0.07rem 0;
							display: flex;
							align-content: center;
							justify-content: center;
							border-radius: 0.04rem;
								img{
									width: 0.91rem;
									height:0.21rem;
								}
							}
						}
						a{
							.app_download_ios_content{
							background: #fff;
							width: 1.6rem;
							height: 0.35rem;
							box-sizing: border-box;
							padding: 0.07rem 0;
							display: flex;
							align-content: center;
							justify-content: center;
							border-radius: 0.04rem;
								img{
									width: 0.82rem;
									height:0.21rem;
								}
							}
						}
						.app_download_qr_content{
							cursor: pointer;
							width: 0.35rem;
							height: 0.35rem;
							background: #fff;
							margin-left: 0.2rem;
							border-radius: 0.04rem;
							box-sizing: border-box;
							padding: 0.05rem;
							position: relative;
							img{
								width: 100%;
								height: 100%;
							}
							&:hover{
								.qr_content{
									display: block;
								}
							}
							.qr_content{
								display: none;
								background: #fff;
								margin-top: 0.1rem;
								position: absolute;
								width: 1.4rem;
								height: 1.55rem;
								border-radius: 0.04rem;
								margin-left: 50%;
								transform: translateX(-50%);
								text-align: center;
								img{
									margin-top: 0.1rem;
									width: 1.2rem;
									height:1.2rem;
								}
								.tooltip_content{
									font-size: 0.14rem;
								}
							}
							.qr_content::before {
								width: 0;
								content: '';
								display: block;
								border: 6px solid transparent;
								border-bottom-color: #ffffff;
								margin-top: -12px;
								margin-left: 48%;
								transform: translateX(-50%);
							}
						}
					}
				}
			}
		}
	}
	@media (max-width: 1040px) {
		.app_download_container{
			.app_download_wrap{
				.app_download_content_wrap{
					grid-template-rows: repeat(2,auto);
					grid-template-columns: repeat(1,auto);
					.app_download_goz_content{
						margin-right: 0;
					}
					.app_download_mainnet_content{
						margin-left: 0;
						margin-top: 0.2rem;
					}
				}
			}
		}
	}
	@media (max-width: 768px) {
		.app_download_container{
			.app_download_wrap{
				.app_download_title{
					padding-left: 0;
				}
			}
		}
	}
	@media (max-width: 520px) {
		.app_download_container{
			.app_download_wrap{
				.app_download_content_wrap{
					.app_download_goz_content{
					
					}
					.app_download_mainnet_content{
						.app_download_img_content{
							flex-direction: column;
							a{
								.app_download_android_content{
									width: auto;
									margin-right: 0;
								}
								.app_download_ios_content{
									width: auto;
									margin-top: 0.1rem;
								}
							}
							.app_download_qr_content{
								display: none;
							}
						}
					}
				}
			}
		}
	}
</style>
