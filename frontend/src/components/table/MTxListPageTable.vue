<template>
    <div>
        <m-table v-table-head-fixed :class="showNoData ? 'show_no_data' : ''" class="tx_container_table override_mtable" :columns="fields" :data="items">
            <template slot-scope="{ row }" slot="Block">
                    <router-link :to="`/block/${row.Block}`" class="link_style">{{row.Block}}</router-link>
            </template>
            <template slot-scope="{ row }" slot="Proposal_ID">
                <router-link v-if="row.Proposal_ID !== '--' && row.Proposal_Title && row.Proposal_Title !== '--'" :to="`/ProposalsDetail/${row.Proposal_ID}`" class="link_style">{{row.Proposal_ID}}</router-link>
                <span v-if="row.Proposal_ID === '--' && row.Proposal_Title && row.Proposal_Title !== '--'">{{row.Proposal_ID}}</span>
                <span v-if="row.Proposal_Title && row.Proposal_Title === '--'">--</span>
            </template>
            <template slot-scope="{ row }" slot="Proposal_Title">
                <router-link v-if="row.Proposal_ID !== '--' && row.Proposal_Title && row.Proposal_Title !== '--'" :to="`/ProposalsDetail/${row.Proposal_ID}`" class="link_style">{{row.Proposal_Title}}</router-link>
                <span v-if="row.Proposal_ID === '--' && row.Proposal_Title && row.Proposal_Title !== '--'">{{row.Proposal_Title}}</span>
                <span v-if="row.Proposal_Title && row.Proposal_Title === '--'">--</span>
            </template>
            <template slot-scope="{ row }" slot="Tx_Hash">
                <div class="common_hover_parent" v-if="row.Tx_Hash">
                    <router-link :to="`/tx?txHash=${row.Tx_Hash}`" class="link_style common_font_style">{{formatTxHash(row.Tx_Hash)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="From">
                <div class="common_hover_address_parent">
                    <span v-if="(/^[1-9]\d*$/).test(row.From)" class="skip_route common_font_style">
                        <router-link :to="`/tx?txHash=${row.Tx_Hash}`">{{row.From}} Validators</router-link>
                    </span>
                    <div class="name_address" v-if="!(/^[0-9]\d*$/).test(row.From) && row.From !== '--'">
                        <div>
                            <span class="remove_default_style">
                                <router-link :to="addressRoute(row.From)" class="link_style justify common_font_style">{{formatMoniker(row.fromMoniker) || formatAddress(row.From)}}
                                </router-link>
                            </span>
                        </div>
                        <span v-if="!row.fromMoniker" class="address">{{row.From}}</span>
                    </div>
                    <span class="no_skip" v-show="(/^[0]\d*$/).test(row.From) || row.From === '--'">--</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="To">
                <div class="common_hover_address_parent">
                    <span v-if="(/^[1-9]\d*$/).test(row.To)" class="skip_route common_font_style">
                        <router-link :to="`/tx?txHash=${row.Tx_Hash}`">{{row.To}} Validators</router-link>
                    </span>
                    <div class="name_address" v-if="!(/^[0-9]\d*$/).test(row.To) && row.To !== '--'">
                        <div>
                            <span class="remove_default_style">
                                <router-link :to="addressRoute(row.To)" class="link_style justify common_font_style">{{formatMoniker(row.toMoniker) || formatAddress(row.To)}}
                                </router-link>
                            </span>
                        </div>
                        <span v-if="!row.toMoniker" class="address">{{row.To}}</span>
                    </div>
                    <span class="no_skip" v-show="(/^[0]\d*$/).test(row.To) || row.To === '--'">--</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="Tx_Signer">
                <div class="common_hover_address_parent" v-if="row.Tx_Signer">
                    <router-link :to="addressRoute(row.Tx_Signer)" class="link_style common_font_style">{{formatAddress(row.Tx_Signer)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="OperatorAddr">
                <div class="common_hover_address_parent">
                    <router-link :to="addressRoute(row.OperatorAddr)" class="link_style common_font_style">{{formatAddress(row.OperatorAddr)}}
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="Moniker">
                <div class="common_hover_address_parent">
                    <router-link v-if="row.Moniker && row.Moniker !== '--'" :to="addressRoute(row.OperatorAddr)" class="link_style">{{row.Moniker}}</router-link>
                </div>
                <span v-if="row.Moniker && row.Moniker === '--'">--</span>
            </template>
        </m-table>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MTxListPageTable",
        props:{
	        items: {
		        type: Array,
		        default: []
	        },
	        showNoData: {
		        type: Boolean,
		        default: true
	        },
        },
        data () {
			return {
				fields:null,
				transferFields: [
          {
            title:'Tx_Hash',
            slot: 'Tx_Hash',
            width: 100,
            tooltip: true,
            tooltipClassName: 'tooltip_left'
          },
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right'
					},
					{
						title:'From',
            slot: 'From'
					},
					{
						title:'Amount',
						slot: 'Amount',
						key: 'Amount',
						className: 'text_right'
					},
					{
						title:'To',
            slot: 'To'
					},
					{
						title:'Tx_Type',
						slot: 'Tx_Type',
						key: 'Tx_Type',
					},
					{
						title:'Tx_Fee',
						slot: 'Tx_Fee',
						key: 'Tx_Fee',
						className: 'text_right'
					},
					{
						title:'Tx_Signer',
            slot: 'Tx_Signer',
            tooltip: true
					},
					{
						title:'Tx_Status',
						slot: 'Tx_Status',
						key: 'Tx_Status',
					},
					{
						title:'Timestamp',
						slot: 'Timestamp',
						key: 'Timestamp',
					},
				],
				declarationFields:[
              {
                title:'Tx_Hash',
                slot: 'Tx_Hash',
                width: 100,
                tooltip: true,
                tooltipClassName: 'tooltip_left'
              },
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right',
					},
					{
						title:'Moniker',
						slot: 'Moniker',
					},
					{
						title:'Operator_Address',
                        slot: 'OperatorAddr',
                        tooltip: true
					},
					{
						title:'Self_Bonded',
						slot: 'Self_Bonded',
						key: 'Amount',
						className: 'text_right',
					},
					{
						title:'Tx_Type',
						slot: 'Tx_Type',
						key: 'Tx_Type',
					},
					{
						title:'Tx_Fee',
						slot: 'Tx_Fee',
						key: 'Tx_Fee',
						className: 'text_right',
					},
					{
						title:'Tx_Signer',
            slot: 'Tx_Signer',
            tooltip: true
					},
					{
						title:'Tx_Status',
						slot: 'Tx_Status',
						key: 'Tx_Status',
					},
					{
						title:'Timestamp',
						slot: 'Timestamp',
						key: 'Timestamp',
					},
                ],
				stakeFields:[
					{
						title:'Tx_Hash',
						slot: 'Tx_Hash',
            width: 100,
            tooltip: true,
            tooltipClassName: 'tooltip_left'
					},
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right',
					},
					{
						title:'From',
            slot: 'From'
					},
					{
						title:'Amount',
						slot: 'Amount',
						key: 'Amount',
						className: 'text_right',
					},
					{
						title:'To',
            slot: 'To'
					},
					{
						title:'Tx_Type',
						slot: 'Tx_Type',
						key: 'Tx_Type',
					},
					{
						title:'Tx_Fee',
						slot: 'Tx_Fee',
						key: 'Tx_Fee',
						className: 'text_right',
					},
					{
						title:'Tx_Signer',
            slot: 'Tx_Signer',
            tooltip: true
					},
					{
						title:'Tx_Status',
						slot: 'Tx_Status',
						key: 'Tx_Status',
					},
					{
						title:'Timestamp',
						slot: 'Timestamp',
						key: 'Timestamp',
					},
				],
				govFields:[
					{
						title:'Tx_Hash',
						slot: 'Tx_Hash',
            width: 100,
            tooltip: true,
            tooltipClassName: 'tooltip_left'
					},
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right',
					},
					{
						title:'Proposal_Type',
						slot: 'Proposal_Type',
						key: 'Proposal_Type',
					},
					{
						title:'Proposal_ID',
						slot: 'Proposal_ID',
						className: 'text_right',
					},
					{
						title:'Proposal_Title',
						slot: 'Proposal_Title',
					},
					{
						title:'Amount',
						slot: 'Amount',
						key: 'Amount',
						className: 'text_right',
					},
					{
						title:'Tx_Type',
						slot: 'Tx_Type',
						key: 'Tx_Type',
					},
					{
						title:'Tx_Fee',
						slot: 'Tx_Fee',
                        key: 'Tx_Fee',
						className: 'text_right',
					},
					{
						title:'Tx_Signer',
            slot: 'Tx_Signer',
            tooltip: true
					},
					{
						title:'Tx_Status',
						slot: 'Tx_Status',
						key: 'Tx_Status',
					},
					{
						title:'Timestamp',
						slot: 'Timestamp',
						key: 'Timestamp',
					},
				],
            }
        },
        methods:{
	        formatAddress(address){
		        return Tools.formatValidatorAddress(address)
	        },
	        formatTxHash(TxHash){
		        if(TxHash){
			        return Tools.formatTxHash(TxHash)
		        }
	        },
	        setTxFields(items){
		        items.forEach( (tx) => {
			        if(tx.listName === 'transfer'){
				        this.fields = this.transferFields
			        }else if(tx.listName === 'declarations') {
				        this.fields = this.declarationFields
			        }else  if(tx.listName === 'stakes'){
				        this.fields = this.stakeFields
			        }else if(tx.listName === 'gov'){
				        this.fields = this.govFields
			        }else {
				        this.fields = []
			        }
		        })
            },
            formatMoniker (moniker) {
                if (!moniker) {
                    return '';
                }
                return Tools.formatString(moniker,15,"...");
            }
        },
		watch:{
			showNoData(showNoData){
			},
			items(items){
				this.setTxFields(items)
			}
		},
	}
</script>

<style lang="scss">
    .show_no_data{
        .m-table-body {
            .m_table tbody{
                display: none;
            }
        }
    }
    table{
        thead{
            tr{
                border-top: 0.01rem solid #d7d9e0;
            }
        }
    }
    .common_hover_parent{
        a{
            display: inline-block;
            position: relative;
        }
    }
    .common_hover_address_parent{
        a{
            position: relative;
        }
    }
.common_font_style{
    font-family: Consolas;
}
</style>
