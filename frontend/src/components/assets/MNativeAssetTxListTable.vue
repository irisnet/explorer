<template>
    <div>
        <m-table :class="showNoData ? 'show_no_data' : ''" class="tx_container_table" :columns="fields" :data="items">
            <template slot-scope="{ row }" slot="TxHash">
                <!--<router-link :to="`/tx?txHash=${row.TxHash}`" class="link_style">{{formatTxHash(row.TxHash)}}</router-link>-->
                <span>{{formatTxHash(row.TxHash)}}</span>
            </template>
            <template slot-scope="{ row }" slot="Owner">
                <div class="skip_route">
                    <router-link :to="`/address/${row.Owner}`" style="font-family: Consolas,Menlo " class="link_style">{{formatAddress(row.Owner)}}</router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="Block">
                <div class="skip_route">
                    <router-link :to="`/block/${row.Block}`" class="link_style">{{row.Block}}</router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="TxHash">
                <div class="skip_route">
                    <router-link :to="`/tx?txHash=${row.TxHash}`" style="font-family: Consolas,Menlo " class="link_style">{{formatTxHash(row.TxHash)}}</router-link>
                </div>
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
            <template slot-scope="{ row }" slot="Symbol">
                <div class="skip_route">
                    <router-link v-if="row.flShowLink" :to="`/asset/${row.TokenID}`" class="link_style">{{row.Symbol}}</router-link>
                </div>
                <span v-if="!row.flShowLink">{{row.Symbol}}</span>
            </template>
            <template slot-scope="{ row }" slot="Token">
                <router-link v-if="row.flShowLink" :to="`/asset/${row.TokenID}`" class="link_style">{{row.Token}}</router-link>
                <span v-if="!row.flShowLink">{{row.Token}}</span>
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
						slot: 'Token',
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
                    },
	                {
		                title:'Gateway',
		                slot: 'Gateway',
	                },
	                {
		                title:'Symbol',
		                slot: 'Symbol',
	                },
	                {
		                title:'InitialSupply',
		                // key:'InitialSupply',
		                slot: 'InitialSupply',
	                },
	                {
		                title:'Mintable',
		                key:'Mintable',
		                slot: 'Mintable',
	                },
	                {
		                title:'Block',
		                slot: 'Block',
		                width: 100,
	                },
	                {
		                title:'TxHash',
		                slot: 'TxHash',
		                tooltip: true,
	                },
	                {
		                title:'Fee',
		                key:'TxFee',
		                slot: 'TxFee(IRIS)',
	                },
	                {
		                title:'Status',
		                key:'TxStatus',
		                slot: 'TxStatus',
	                },
	                {
		                title:'Timestamp',
		                key:'Timestamp',
		                slot: 'Timestamp',
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
				],
				transferGatewayOwnerTxs: [
					{
						title:'Gateway',
						key: 'Gateway',
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
    .name_address {
        display: inline-block;
        position: relative;
        .address {
            display: none;
            position: absolute;
            padding: 0rem 0.15rem;
            top: -0.36rem;
            left: 50%;
            transform: translateX(-50%);
            color: #fff;
            background: rgba(0, 0, 0, 1);
            border-radius: 0.04rem;
            z-index: 10;
            line-height: 32px;
            font-size: 0.14rem;
            &::after {
                width: 0;
                height: 0;
                border: 0.06rem solid transparent;
                content: "";
                display: block;
                position: absolute;
                border-top-color: #000000;
                left: 50%;
                margin-left: -6px;
            }
        }
        &:hover {
            .address {
                display: block;
            }
        }
    }
.tx_container_table{
    .m-table-header{
        overflow: hidden;
    }
}
</style>
