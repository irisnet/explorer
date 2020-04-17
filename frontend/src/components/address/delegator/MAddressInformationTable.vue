<template>
    <div class="address_detail_table" style="background:#fff;min-height: 2.34rem;">
        <m-table :columns="fields" :data="items" :width="width">
            <template slot-scope="{ row }"
                      slot="address">
                <router-link v-if="row.moniker" class="address_link" :to="`/validators/${row.address}`">{{formatMoniker(row.moniker)}}</router-link>
                <router-link v-if="!row.moniker" style="font-family:Consolas,Menlo" class="address_link" :to="`/validators/${row.address}`">{{formatAddress(row.address)}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="txHash">
                <div class="common_hover_parent  skip_route" v-if="row.txHash">
                    <router-link style="font-family: Consolas,Menlo " :to="`/tx?txHash=${row.txHash}`" class="link_style common_font_style">{{formatTxHash(row.txHash)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="block">
                <span class="skip_route">
                    <router-link v-if="row.block != 0" :to="`/block/${row.block}`">{{row.block}}</router-link>
                </span>
                <span v-if="row.block == 0">{{row.block}}</span>
            </template>
            <template slot-scope="{ row }" slot="signer">
                <div class="common_hover_address_parent" :class="row.isSkipRouter ? '' : 'skip_route'" v-if="row.signer"  style="font-family:Consolas,Menlo" >
                    <router-link :style="{cursor: row.isSkipRouter ? 'auto' : 'pointer'}" :to="addressRoute(row.signer)">{{formatAddress(row.signer)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="from">
                <div class="common_hover_address_parent">
                    <span v-if="(/^[1-9]\d*$/).test(row.from)" class="skip_route common_font_style">
                        <router-link :to="`/tx?txHash=${row.txHash}`">{{row.from}} Validators</router-link>
                    </span>
                    <div class="name_address" v-if="!(/^[0-9]\d*$/).test(row.from) && row.from !== '--'">
                        <div>
                            <span class="remove_default_style " :class="row.isFromSkipRouter ? '' : 'skip_route'">
                                <router-link :to="addressRoute(row.from)"
                                             :style="{cursor: row.isFromSkipRouter ? 'auto' : 'pointer'}"
                                             style="font-family: Consolas,Menlo;">{{formatMoniker(row.fromMoniker) || formatAddress(row.from)}}
                                </router-link>
                            </span>
                        </div>
                    </div>
                    <span class="no_skip" v-show="(/^[0]\d*$/).test(row.from) || row.from === '--'">--</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="to">
                <div class="common_hover_address_parent">
                    <span v-if="(/^[1-9]\d*$/).test(row.to)" class="skip_route common_font_style">
                        <router-link :to="`/tx?txHash=${row.txHash}`">{{row.to}} Validators</router-link>
                    </span>
                    <div class="name_address" v-if="!(/^[0-9]\d*$/).test(row.to) && row.to !== '--'">
                        <div>
                            <span class="remove_default_style"  :class="row.isToSkipRouter ? '' : 'skip_route'">
                                <router-link :to="addressRoute(row.to)"
                                             :style="{cursor: row.isToSkipRouter ? 'auto' : 'pointer'}"
                                             style="font-family: Consolas,Menlo;">{{formatMoniker(row.toMoniker) || formatAddress(row.to)}}
                                </router-link>
                            </span>
                        </div>
                    </div>
                    <span class="no_skip" v-show="(/^[0]\d*$/).test(row.to) || row.to === '--'">--</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="token">
                <div style="display: flex;align-items: center">
                   <span v-if="row.token === 'IRIS'" style="display: flex;align-items: center">
                       <img style="width: 0.18rem;height:0.18rem;margin-right: 0.1rem;" src="../../../assets/iris_token_logo.png">
                   </span>
                    <span v-if="row.token !== 'IRIS'">
                        <i style="margin-right: 0.1rem;color:#B6BAD2;font-size: 0.18rem" class="iconfont iconqitabizhongwuiconshishiyongdemorenicon"></i>
                    </span>
                    <span>{{row.token}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="totalAmount">
                <div>
                    <span>{{row.totalAmount === 0 ? "--" : row.totalAmount}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="balance">
                <div>
                    <span>{{row.balance === 0 ? "--" : row.balance}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="delegated">
                <div>
                    <span>{{row.delegated === 0 ? "--" : row.delegated}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="unBonding">
                <div>
                    <span>{{row.unBonding === 0 ? "--" : row.unBonding}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="reward">
                <div>
                    <span>{{row.reward === 0 ? "--" : row.reward}}</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="amount">
                <div class="name_address">
                    <div v-if="!row.amount.includes('Tokens') && row.amount.toString().length <= 12">
                            <span>
                                <span>{{row.amount}}</span>
                            </span>
                    </div>
                    <div  v-if="row.amount.includes('Tokens')"  class="skip_route">
                        <router-link :to="`/tx?txHash=${row.txHash}`">{{row.amount}}</router-link>
                    </div>
                </div>
            </template>
        </m-table>
        <div class="no_data_img_content" v-if="showNoData">
            <span>~ {{showNoDataDoc}} ~</span>
        </div>
    </div>
</template>

<script>
    import Tools from "../../../util/Tools"
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
            showNoDataDoc:{
			  type:String,
            },
			width: {
				type: Number,
				default: 1280
			},
			showNoData: {
				type: Boolean,
				default: true
			},
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
						slot:'totalAmount',
					},
					{
						title: "Balance",
						slot:'balance',
					},
					{
						title: "Delegated",
						slot:'delegated',
					},
					{
						title: "Unbonding",
						slot:'unBonding',
					},
					{
						title: "Rewards",
						slot:'reward',
					},
				],
				delegations: [
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
						tooltip: true
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
                        title:"From",
                        slot:'from'
                    },
					{
						title: "Amount",
                        slot:'amount',
						className: 'text_right'
					},
                    {
                        title: 'To',
                        slot: 'to'
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
						title: "Status",
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
	        formatMoniker(moniker) {
		        if (!moniker) {
			        return "";
		        }
		        return Tools.formatString(moniker, 15, "...");
	        },
            substrAmount(amount){
	            let tokenValue;
	            if(amount !== '--'){
                    tokenValue = Tools.formatAccountCoinsAmount(amount);
                }
                if(tokenValue && tokenValue.toString().length > 12){
                    return Tools.formatString(amount.toString(),12,'...')
                }else {
                    return amount
                }
            },
        }
	}
</script>

<style lang="scss">
    .address_link{
        color:var(--bgColor) !important;
    }
    .address_detail_table{
        .no_data_img_content{
            display: flex;
            justify-content: center;
            font-size: 0.14rem;
            height: 1.8rem;
            align-items: center;
            color: var(--contentColor);
            img{
                width: 1.2rem;
            }
        }
    }
    .address_information_delegation_list_content{
        .address_detail_table {
            .m-table-header {
                width: calc(100% - 0.02rem);
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
                width: calc(100% - 0.02rem);
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
                width: calc(100% - 0.02rem);
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
