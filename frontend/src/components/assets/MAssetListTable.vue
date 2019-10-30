<template>
    <div>
        <m-table v-if="items.length > 0" :class="showNoData ? 'show_no_data' : ''" class="tx_container_table" :columns="fields" :data="items">
            <template slot-scope="{ row }" slot="Owner">
                <div class="skip_route">
                    <router-link :to="`/address/${row.Owner}`" class="link_style">{{row.Owner}}</router-link>
                </div>
            </template>
 <!--           <template slot-scope="{ row }" slot="Owner">
                <router-link :to="`/address/${row.Owner}`" class="link_style">{{row.Owner}}</router-link>
            </template>-->
            <template slot-scope="{ row }" slot="Symbol">
                <router-link :to="`/asset/${row.TokenID}`" class="link_style">{{row.Symbol}}</router-link>
                <!--<router-link v-if="row.Gateway" :to="`/asset/${row.Gateway}/${row.Symbol}`" class="link_style">{{row.Symbol}}</router-link>-->
            </template>
        </m-table>
    </div>
</template>

<script>
	export default {
		name: "MAssetListTable",
        props:{
	        items:{
		        type: Array,
		        default: []
            },
	        showNoData:{
		        type: Boolean,
		        default: true
            },
	        name: {
		        type: String,
		        default: ""
	        },
        },
        data () {
            return {
	            fields: null,
	            nativeAssetList:[
                    {
                    	title:'Symbol',
	                    slot:'Symbol',
                    },
	                {
		                title:'Owner',
		                slot:'Owner',
	                },
	                {
		                title:'TotalSupply',
		                key:'TotalSupply',
	                },
	                {
		                title:'InitialSupply',
		                key:'InitialSupply',
	                },
	                {
		                title:'MaxSupply',
		                key:'MaxSupply',
	                },
	                {
		                title:'Mintable',
		                key:'Mintable',
	                },
                ],
	            gatewayAssetList:[
		            {
			            title:'Gateway',
			            key:'Gateway',
		            },
		            {
			            title:'Symbol',
			            slot:'Symbol',
		            },
		            {
			            title:'Owner',
			            slot:'Owner',
		            },
		            {
			            title:'TotalSupply',
			            key:'TotalSupply',
		            },
		            {
			            title:'InitialSupply',
			            key:'InitialSupply',
		            },
		            {
			            title:'MaxSupply',
			            key:'MaxSupply',
		            },
		            {
			            title:'Mintable',
			            key:'Mintable',
		            },
	            ]
            }
        },
		mounted(){
			this.fields = this[this.name] || [];

		},
	}
</script>

<style scoped lang="scss">
.link_style{
    color: var(--bgColor) !important;
}
</style>
