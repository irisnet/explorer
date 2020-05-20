<template>
	<div class="advertising_space_container" :class="pageClass === 'txDetail' ? 'tx_detail_class' :'' ">
		<div class="swiper_container">
			<swiper ref="mySwiper" :options="swiperOption"  class="my-swipe" v-if="true">
				<swiper-slide v-for="item in swiperList" :key="item.id" class="swipe_slider">
					<a :href="item.href" target="_blank" class="swipe_slider_link">
						<img :src="item.src" :class="item.className">
					</a>
				</swiper-slide>
			</swiper>
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
			},
		data() {
			 return {
				 swiperOption:{
					 loop:true,
					 autoplay: {
						 delay: 5000,
						 disableOnInteraction:false,
					 },
					 speed: 300,
					 // direction : 'vertical',
				 },
				 swiperList:[
					 {
					 	 id:0,
						 href:'https://ibc-goz.irisplorer.io/#/',
						 src:require("../../assets/advertising_space_img/irisnet-goz-1280.jpg"),
						 className:'goz_explorer_class'
					 },
					 {
						 id:1,
						 href:'https://ibc-goz.irisplorer.io/#/download-rainbowgoz',
						 src:require("../../assets/advertising_space_img/goz_wallet_1280.jpg"),
						 className:'goz_wallet_class'
					 }
				 ]
			 }
		},
		computed: {
			mySwiper() {
				return this.$refs.mySwiper.swiper
			},
			
		},
	}
</script>

<style scoped lang="scss">
	.advertising_space_container{
		padding-top: 0.2rem;
		.swiper_container{
			max-width: 12.4rem;
			margin: 0 auto;
			.my-swipe{
				.swipe_slider{
					.swipe_slider_link{
						.goz_explorer_class{
							width: 100%;
							content:url("../../assets/advertising_space_img/irisnet-goz-1280.jpg")
						}
						.goz_wallet_class{
							width: 100%;
							content:url("../../assets/advertising_space_img/goz_wallet_1280.jpg")
						}
					}
				}
			}
		}
	}
	.tx_detail_class{
		.swiper_container{
			max-width: 12.8rem;
		}
	}
	@media screen and (max-width: 1280px){
		.advertising_space_container{
			width: 100%;
			.swiper_container{
				width: 100% !important;
				max-width: none;
				.my-swipe{
					margin: 0 0.2rem;
					.swipe_slider{
					
					}
				}
			}
		}
	}
	@media screen and (max-width: 910px){
		.advertising_space_container{
			.swiper_container{
				.my-swipe{
					margin: 0 0.1rem;
					.swipe_slider{
						.swipe_slider_link{
							.goz_explorer_class{
								content:url("../../assets/advertising_space_img/irisnet-goz-768.jpg")
							}
							.goz_wallet_class{
								content:url("../../assets/advertising_space_img/goz_wallet_768.jpg")
							}
						}
					}
				}
			}
		}
	}
	@media screen and (max-width: 705px){
		.advertising_space_container{
			.swiper_container{
				.my-swipe{
					.swipe_slider{
						.swipe_slider_link{
							.goz_explorer_class{
								width: 100%;
								content:url("../../assets/advertising_space_img/irisnet-goz-375.jpg")
							}
							.goz_wallet_class{
								width: 100%;
								content:url("../../assets/advertising_space_img/goz_wallet_375.jpg")
							}
						}
					}
				}
			}
		}
	}
</style>
