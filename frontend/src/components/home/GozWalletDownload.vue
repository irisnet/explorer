<template>
	<div class="download_container">
		<div class="download_content">
			<app-download></app-download>
			<validator-bianjie-information></validator-bianjie-information>
		</div>
	</div>
</template>

<script>
	import AppDownload from "@/components/home/AppDownload";
	import ValidatorBianjieInformation from "@/components/home/ValidatorBianjieInformation";
	export default {
		name: "GozWalletDownload",
		components: {ValidatorBianjieInformation, AppDownload}
	}
</script>

<style scoped lang="scss">
	.download_container{
		width: 100%;
		height:100% !important;
		.download_content{
			width: 100%;
			height:100%;
			padding: 0 0.4rem;
			box-sizing: border-box;
			background-image: url("../../assets/start_bg.jpg");
			background-repeat: no-repeat;
			background-position: center;
			background-size: cover;
		}
	}
	@media (max-width: 768px) {
		.download_container{
			.download_content{
				padding: 0 0.2rem;
			}
		}
	}
</style>
