<template>
    <div :style="{'min-width': minWidth + 'rem'}">
        <m-table v-if="!flTableFixed" v-table-head-fixed :columns="columns"
                 :data="items">
            <template slot-scope="{ row }" slot="txHash">
                <div class="skip_route" v-if="row.txHash">
                    <router-link :to="`/tx?txHash=${row.txHash}`" style="font-family: Consolas,Menlo ;" class="link_style common_font_style">{{formatTxHash(row.txHash)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="block">
                <span class="skip_route">
                    <router-link :to="`/block/${row.block}`" class="link_style">{{row.block}}</router-link>
                </span>
            </template>
            <template slot-scope="{ row }" slot="signer">
                <div class="skip_route" v-if="row.signer">
                    <router-link :to="addressRoute(row.signer)" style="font-family: Consolas,Menlo ;" class="link_style common_font_style">{{formatAddress(row.signer)}}
                    </router-link>
                </div>
            </template>
        </m-table>
        <m-table v-if="flTableFixed" :columns="columns"
                 :data="items">
            <template slot-scope="{ row }" slot="txHash">
                <div class="skip_route" v-if="row.txHash">
                    <router-link :to="`/tx?txHash=${row.txHash}`" class="link_style common_font_style">{{formatTxHash(row.txHash)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="block">
                <span class="skip_route">
                    <router-link :to="`/block/${row.block}`">{{row.block}}</router-link>
                </span>
            </template>
            <template slot-scope="{ row }" slot="signer">
                <div class="skip_route" v-if="row.signer">
                    <router-link :to="addressRoute(row.signer)" class="link_style common_font_style">{{formatAddress(row.signer)}}
                    </router-link>
                </div>
            </template>
        </m-table>
        <!-- TODO  新增无数据展示-->
        <no-data :fl-show-no-data="true" :no-data-doc="'block detail transaction list'"></no-data>
    </div>
</template>

<script>
    import Tools from '../../util/Tools'
    import NoData from "../noDataComponent/NoData";
	export default {
		name: "MAllTxTypeListTable",
        components: {NoData},
        props: {
			items: {
				type: Array,
				default: function() {return[]}
			},
			minWidth: {
				type: Number,
				default: 12.8
			},
			flTableFixed: {
				type: Boolean,
            }
		},
        data(){
			return {
				columns:[
					{
						title: 'TxHash',
						slot: 'txHash',
						tooltip: true,
					},
					{
						title: 'Block',
						slot: 'block',
					},
					{
						title: 'TxType',
						slot: 'type',
						key:"type",
					},
					{
						title: 'Fee',
						key:"fee"
					},
					{
						title: 'Signer',
						slot: 'signer',
						tooltip: true,
					},
					{
						title: 'Status',
						slot: 'status',
						key:"status"
					},
					{
						title: 'Timestamp',
						slot: 'timestamp',
						key:"timestamp"
					},
                ]
            }
        },
        methods:{
	        formatTxHash(TxHash){
		        if(TxHash){
			        return Tools.formatTxHash(TxHash)
		        }
	        },
	        formatAddress(address){
		        return Tools.formatValidatorAddress(address)
	        },
        }

	}
</script>

<style scoped lang="scss">

</style>
