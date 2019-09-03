<template>
    <div>
        <m-table :class="showNoData ? 'show_no_data' : ''" class="tx_container_table" :columns="fields" :data="items">
            <template slot-scope="{ row }" slot="TxHash">
                <!--<router-link :to="`/tx?txHash=${row.TxHash}`" class="link_style">{{formatTxHash(row.TxHash)}}</router-link>-->
                <span>{{formatTxHash(row.TxHash)}}</span>
            </template>
            <template slot-scope="{ row }" slot="Owner">
                <router-link :to="`/address/${row.Owner}`" class="link_style">{{formatAddress(row.Owner)}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="Block">
                <router-link :to="`/block/${row.Block}`" class="link_style">{{row.Block}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="TxHash">
                <router-link :to="`/tx?txHash=${row.TxHash}`" class="link_style">{{formatTxHash(row.TxHash)}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="SrcOwner">
                <router-link :to="`/address/${row.SrcOwner}`" class="link_style">{{formatAddress(row.SrcOwner)}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="DstOwner">
                <router-link :to="`/address/${row.DstOwner}`" class="link_style">{{formatAddress(row.DstOwner)}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="MintTo">
                <router-link :to="`/address/${row.MintTo}`" class="link_style">{{formatAddress(row.MintTo)}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="Amount">
                <div class="name_address">
                    <div>
                            <span>
                                <span>{{substrAmount(row.Amount)}}</span>
                            </span>
                    </div>
                    <span class="address" v-if="row.Amount.toString().length > 12">{{row.Amount}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="InitialSupply" v-if="row.InitialSupply">
                <div class="name_address">
                    <span>{{substrAmount(row.InitialSupply)}}</span>
                    <span class="address" v-if="row.InitialSupply.toString().length > 12">{{row.InitialSupply}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="Gateway">
                <div class="Gateway">
                    <span>{{row.Gateway}}</span>
                </div>
            </template>
        </m-table>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MNativeAssetTxListTable",
		props:{
			items: {
				type: Array,
				default: []
			},
			name: {
				type: String,
				default: ""
			},
			showNoData: {
				type: Boolean,
				default: true
			},
		},
        data(){
			return {
				fields:null,
				assetTransferTxs:[
                    {
                    	title: 'TxHash',
                        slot: 'TxHash',
	                    tooltip: true,
                    },
					{
						title: 'Block',
						slot: 'Block',
						width: 100,
						tooltipClassName: 'tooltip_left'
					},
					{
						title: 'From',
						slot: 'From',
					},
					{
						title: 'Amount',
						key: 'Amount',
						tooltipClassName: 'tooltip_left',
					},
					{
						title: 'Token',
						key: 'Token',
						tooltipClassName: 'tooltip_left',
					},
					{
						title: 'To',
						slot: 'To',
					},
					{
						title: 'TxType',
						key: 'TxType',
					},
					{
						title: 'Fee',
						key: 'Fee',
					},
					{
						title: 'Signer',
						slot: 'Signer',
					},
					{
						title: 'Status',
						key: 'Status',
					},
					{
						title: 'Timestamp',
						key: 'Timestamp',
					},

                ],
				gateWayIssueToken: [
                    {
	                    title:'Owner',
	                    slot: 'Owner',
	                    tooltip: true,
	                    tooltipClassName: 'tooltip_left'
                    },
	                {
		                title:'Gateway',
		                slot: 'Gateway',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Symbol',
		                key:'Symbol',
		                slot: 'Symbol',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'InitialSupply',
		                // key:'InitialSupply',
		                slot: 'InitialSupply',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Mintable',
		                key:'Mintable',
		                slot: 'Mintable',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Block',
		                slot: 'Block',
		                width: 100,
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'TxHash',
		                slot: 'TxHash',
		                tooltip: true,
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Fee',
		                key:'TxFee',
		                slot: 'TxFee(IRIS)',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Status',
		                key:'TxStatus',
		                slot: 'TxStatus',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Timestamp',
		                key:'Timestamp',
		                slot: 'Timestamp',
		                tooltipClassName: 'tooltip_left'
	                }
                ],
				nativeIssueToken: [
					{
						title:'Owner',
						slot: 'Owner',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},

					{
						title:'Symbol',
						key:'Symbol',
						slot: 'Symbol',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'InitialSupply',
						// key:'InitialSupply',
						slot: 'InitialSupply',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Mintable',
						key:'Mintable',
						slot: 'Mintable',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Block',
						slot: 'Block',
						width: 100,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxHash',
						slot: 'TxHash',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Fee',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Status',
						key:'TxStatus',
						slot: 'TxStatus',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Timestamp',
						key:'Timestamp',
						slot: 'Timestamp',
						tooltipClassName: 'tooltip_left'
					}
				],
				editToken: [
					{
						title:'Token',
						key:'Token',
						slot: 'Token',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Owner',
						slot: 'Owner',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Block',
						slot: 'Block',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxHash',
						slot: 'TxHash',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Fee',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Status',
						key:'TxStatus',
						slot: 'TxStatus',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Timestamp',
						key:'Timestamp',
						slot: 'Timestamp',
						tooltipClassName: 'tooltip_left'
					},
				],
				mintToken: [
					{
						title:'Token',
						key:'Token',
						slot: 'Token',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Owner',
						slot: 'Owner',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'MintTo',
						slot: 'MintTo',
                        tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Amount',
						// key:'Amount',
						slot: 'Amount',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Block',
						slot: 'Block',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxHash',
						slot: 'TxHash',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Fee',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Status',
						key:'TxStatus',
						slot: 'TxStatus',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Timestamp',
						key:'Timestamp',
						slot: 'Timestamp',
						tooltipClassName: 'tooltip_left'
					},
				],
				transferToken: [
					{
						title:'Token',
						key:'Token',
						slot: 'Token',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'SrcOwner',
						tooltip: true,
						slot: 'SrcOwner',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'DstOwner',
						tooltip: true,
						slot: 'DstOwner',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Block',
						slot: 'Block',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxHash',
						slot: 'TxHash',
						tooltip: true,
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Fee',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Status',
						key:'TxStatus',
						slot: 'TxStatus',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Timestamp',
						key:'Timestamp',
						slot: 'Timestamp',
						tooltipClassName: 'tooltip_left'
					},
				]
            }
        },
        mounted(){
	        this.fields = this[this.name] || [];
        },
		methods:{
			formatTxHash(TxHash){
				if(TxHash){
					return Tools.formatTxHash(TxHash)
				}
            },
			substrAmount(amount){
				if(amount && amount.toString().length > 12){
					return Tools.formatString(amount.toString(),12,'...')
				}else {
					return amount
				}
			},
			formatAddress(address){
				return Tools.formatValidatorAddress(address)
			},
        }
	}
</script>

<style lang="scss" >
.tx_container_table{
    .m-table-header{
        overflow: hidden;
    }
}
</style>
