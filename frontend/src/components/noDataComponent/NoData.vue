<template>
	<div class="no_data_container" v-if="flShowNoData">
		<div class="no_data_content">
			<span class="no_data_text">~ {{noDataDoc}} ~</span>
		</div>
	</div>
</template>

<script>
	export default {
		name: "NoData",
		props:{
			flShowNoData:{
				type:Boolean
			},
			noDataDoc:{
				type: String
			}
		}
	}
</script>

<style scoped lang="scss">
	.no_data_container{
		display: flex;
		align-items: center;
		justify-content: center;
		min-height: 1.2rem;
		background: #fff;
		color: var(--contentColor);
	}
</style>
