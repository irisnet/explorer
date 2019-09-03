<template>
    <div>
        <m-table v-if="items.length > 0" :class="showNoData ? 'show_no_data' : ''" class="tx_container_table" :columns="fields" :data="items">
            <template slot-scope="{ row }" slot="Owner">
                <router-link :to="`/address/${row.Owner}`" class="link_style">{{row.Owner}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="Owner">
                <router-link :to="`/address/${row.Owner}`" class="link_style">{{row.Owner}}</router-link>
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
                        key:'Symbal',
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
			            slot:'Gateway',
		            },
		            {
			            title:'Symbal',
			            slot:'Symbal',
		            },
		            {
			            title:'Owner',
			            key:'Owner',
		            },
		            {
			            title:'CurrentSupply',
			            key:'CurrentSupply',
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
			console.log(this.name,"/////////////////")

		},
	}
</script>

<style scoped>

</style>
