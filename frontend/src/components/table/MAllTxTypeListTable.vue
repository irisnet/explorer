<template>
    <div :style="{'min-width': minWidth + 'rem'}">
        <m-table v-table-head-fixed :columns="columns"
                 :data="items">
            <template slot-scope="{ row }" slot="txHash">
                <div class="common_hover_parent" v-if="row.txHash">
                    <router-link :to="`/tx?txHash=${row.txHash}`" class="link_style common_font_style">{{formatTxHash(row.txHash)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="block">
                <router-link :to="`/block/${row.block}`" class="link_style">{{row.block}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="signer">
                <div class="common_hover_address_parent" v-if="row.signer">
                    <router-link :to="addressRoute(row.signer)" class="link_style common_font_style">{{formatAddress(row.signer)}}
                    </router-link>
                </div>
            </template>
        </m-table>
    </div>
</template>

<script>
    import Tools from '../../util/Tools'
	export default {
		name: "MAllTxTypeListTable",
		props: {
			items: {
				type: Array,
				default: function() {return[]}
			},
			minWidth: {
				type: Number,
				default: 12.8
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
						title: 'Type',
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
    .common_hover_address_parent{
        a{
            position: relative;
        }
    }
</style>
