<template>
    <div class="address_detail_table" style="background:#fff;min-height: 2.34rem;">
        <m-table :columns="fields" :data="items" :width="width">
            <template slot-scope="{ row }"
                      slot="address">
                <router-link  v-if="row.moniker" class="address_link" :to="`/validators/${row.address}`">{{formatAddress(row.moniker)}}</router-link>
                <router-link v-if="!row.moniker" class="address_link" :to="`/validators/${row.address}`">{{formatAddress(row.address)}}</router-link>
            </template>
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
            <template slot-scope="{ row }" slot="token">
                <div>
                   <span v-if="row.token === 'IRIS'">
                       <img style="width: 0.18rem;height:0.18rem;margin-right: 0.1rem" src="../../assets/iris_token_logo.png">
                   </span>
                    <span v-if="row.token !== 'IRIS'">
                        <i style="margin-right: 0.1rem" class="iconfont iconqitabizhongwuiconshishiyongdemorenicon"></i>
                    </span>
                    <span>{{row.token}}</span>
                </div>
            </template>
        </m-table>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MAddressInformationTable",
		props: {
			items: {
				type: Array,
				default: []
			},
			listName: {
				type: String,
				default: ""
			},
			width: {
				type: Number,
				default: 1280
			}
        },
        data (){
			return {
				fields: [],
				assets: [
					{
						title: "Token",
                        slot:'token',
					},
					{
						title: "Total Amount",
						key:'totalAmount',
					},
					{
						title: "Balance",
						key:'balance',
					},
					{
						title: "Delegated",
						key:'delegated',
					},
					{
						title: "Unbonding",
						key:'unBonding',
					},
					{
						title: "Rewards",
						key:'reward',
					},
				],
				delegations: [
					{
						title: "Address",
						slot:'address',

					},
					{
						title: "Amount",
						key:'amount',
						className: 'text_right'
					},
					{
						title: "Shares",
						key:'shares'
					},
					{
						title: "Block",
						slot:'block'
					},

				],
				unBondingDelegations: [
					{
						title: "Address",
                        slot:'address',
					},
					{
						title: "Amount",
                        key:'amount',
						className: 'text_right'
					},
					{
						title: "Block",
                        slot: "block"
					},
					{
						title: "End Time",
                        key:'endTime'
					},

				],
				rewards: [
					{
						title: "Address",
                        slot:'address',
						tooltip: true
					},
					{
						title: "Amount",
                        key:'amount',
						className: 'text_right'
					},

				],
				transactions: [
					{
						title: "TxHash",
                        slot:'txHash',
						tooltip: true
					},
					{
						title: "Block",
                        slot:'block',
					},
					{
						title: "Amount",
                        key:'amount',
						className: 'text_right'
					},
					{
						title: "TxType",
                        key:'txType',
					},
					{
						title: "Fee",
                        key:'fee'
					},
					{
						title: "Signer",
                        slot:'signer'
					},
					{
						title: "status",
                        key:'status'
					},

					{
						title: "Timestamp",
                        key:'timestamp'
					}
				],
            }
        },
		mounted() {
			this.fields = this[this.listName] || [];
		},
        methods:{
	        formatAddress(address) {
		        return Tools.formatValidatorAddress(address);
	        },
	        formatTxHash(TxHash){
		        if(TxHash){
			        return Tools.formatTxHash(TxHash)
		        }
	        },
        }
	}
</script>

<style lang="scss">
    .address_link{
        color:var(--bgColor) !important;
    }
    .address_information_delegation_list_content{
        .address_detail_table {
            .m-table-header {
                width: 100%;
                border-left: 0.01rem solid #E7E9EB;
                border-right: 0.01rem solid #E7E9EB;
                border-bottom: 0.01rem solid #E7E9EB;
                border-radius: 0.04rem;
                table.m_table {
                    width: auto;
                }
                table thead tr{
                    border-top: none;
                    border-left: none;
                    border-right: none;
                }
            }
            table.m_table {
                min-width: 100%;
            }
        }
    }
    .address_information_unbonding_delegation_list_content{
        .address_detail_table {
            .m-table-header {
                width: 100%;
                border-left: 0.01rem solid #E7E9EB;
                border-right: 0.01rem solid #E7E9EB;
                border-bottom: 0.01rem solid #E7E9EB;
                border-radius: 0.04rem;
                table.m_table {
                    width: auto;
                }
                table thead tr{
                    border-top: none;
                    border-left: none;
                    border-right: none;
                }
            }
            table.m_table {
                min-width: 100%;
            }
        }
    }
    .address_information_list_content{
        .address_detail_table {
            .m-table-header {
                width: 100%;
                border-left: 0.01rem solid #E7E9EB;
                border-right: 0.01rem solid #E7E9EB;
                border-bottom: 0.01rem solid #E7E9EB;
                border-radius: 0.04rem;
                table.m_table {
                    width: auto;
                }
                table thead tr{
                    border-top: none;
                    border-left: none;
                    border-right: none;
                }
            }
            table.m_table {
                min-width: 100%;
            }
        }
    }
</style>
