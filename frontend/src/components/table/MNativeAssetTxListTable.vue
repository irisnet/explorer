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
                <!--<span>{{row.Block}}</span>-->
            </template>
            <template slot-scope="{ row }" slot="SrcOwner">
                <router-link :to="`/address/${row.SrcOwner}`" class="link_style">{{formatAddress(row.SrcOwner)}}</router-link>
                <!--<span>{{row.Block}}</span>-->
            </template>
            <template slot-scope="{ row }" slot="DstOwner">
                <router-link :to="`/address/${row.DstOwner}`" class="link_style">{{formatAddress(row.DstOwner)}}</router-link>
                <!--<span>{{row.Block}}</span>-->
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
                issueToken: [
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
		                key:'InitialSupply',
		                slot: 'InitialSupply',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'MaxSupply',
		                key:'MaxSupply',
		                slot: 'MaxSupply',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Mintable',
		                key:'Mintable',
		                slot: 'Mintable',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Decimal',
		                key:'Decimal',
		                slot: 'Decimal',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'SymbolMin',
		                key:'SymbolMin',
		                slot: 'SymbolMin',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'Name',
		                key:'Name',
		                slot: 'Name',
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
		                title:'TxFee(IRIS)',
		                key:'TxFee',
		                slot: 'TxFee(IRIS)',
		                tooltipClassName: 'tooltip_left'
	                },
	                {
		                title:'TxStatus',
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
						title:'MaxSupply',
						key:'MaxSupply',
						slot: 'MaxSupply',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Mintable',
						key:'Mintable',
						slot: 'Mintable',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Decimal',
						key:'Decimal',
						slot: 'Decimal',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'SymbolMin',
						key:'SymbolMin',
						slot: 'SymbolMin',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Name',
						key:'Name',
						slot: 'Name',
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
						title:'TxFee(IRIS)',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxStatus',
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
						key:'MintTo',
						slot: 'MintTo',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'Amount',
						key:'Amount',
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
						title:'TxFee(IRIS)',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxStatus',
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
						slot: 'SrcOwner',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'DstOwner',
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
						title:'TxFee(IRIS)',
						key:'TxFee',
						slot: 'TxFee(IRIS)',
						tooltipClassName: 'tooltip_left'
					},
					{
						title:'TxStatus',
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
			formatAddress(address){
				return Tools.formatValidatorAddress(address)
			},
        }
	}
</script>

<style lang="scss" scoped>

</style>