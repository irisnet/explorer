<template>
	<div class="advertising_space_container" :class="pageClass === 'txDetail' ? 'tx_detail_class' :'' ">
		<div class="advertising_space_content">
			<a href="https://www.binance.com/en/vote" target="_blank" @click="$uMeng.push('HomeOverview_Banner','click')">
				<img class="advertising_space_img_content" src="../../assets/advertising_space_img/advertising_space.png" alt="">
				<div class="advertising_vote_btn"></div>
			</a>
		</div>
	</div>
</template>

<script>
	export default {
		name: "AdvertisingSpace",
		props: {
				pageClass: {
					type: String
				},
			}
	}
</script>

<style scoped lang="scss">
	.advertising_space_container{
		padding-top: 0.2rem;
		.advertising_space_content{
			max-width: 12.4rem;
			margin: 0 auto;
			.advertising_space_img_content{
				width: 100%;
				height: 1rem;
				display: inline-block;
				content: url("../../assets/advertising_space_img/advertising_space.png");
			}
		}
	}
	.tx_detail_class{
		.advertising_space_content{
			max-width: 12.8rem;
			padding-top: 0.1rem;
		}
	}
	@media screen and (max-width: 1280px){
		.advertising_space_container{
			.advertising_space_content{
				padding: 0 0.2rem;
				.advertising_space_img_content{
					height: auto;
				}
			}
		}
	}
	@media screen and (max-width: 910px){
		.advertising_space_container{
			.advertising_space_content{
				padding: 0 0.1rem;
				height: auto;
				.advertising_space_img_content{
					content: url("../../assets/advertising_space_img/advertising_space_768.png");
				}
			}
		}
	}
	@media screen and (max-width: 705px){
		.advertising_space_container{
			.advertising_space_content{
				padding: 0 0.1rem;
				.advertising_space_img_content{
					content: url("../../assets/advertising_space_img/mobile_advertising_space.png");
				}
			}
		}
	}
</style>
