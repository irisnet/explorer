<template>
    <div>
        <m-table v-table-head-fixed :class="showNoData ? 'show_no_data' : ''" class="tx_container_table" :columns="fields" :data="items">
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
                        <div class="hover_content">
                            {{row.Tx_Hash}}
                        </div>
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="From">
                <div class="common_hover_address_parent">
                    <span v-if="(/^[1-9]\d*$/).test(row.From)" class="skip_route common_font_style">
                        <router-link :to="`/tx?txHash=${row.Tx_Hash}`">{{row.From}} Validators</router-link>
                    </span>
                    <div v-if="!(/^[0-9]\d*$/).test(row.From) && row.From !== '--'">
                        <div>
                            <span class="remove_default_style">
                                <router-link :to="`/address/1/${row.From}`" class="link_style justify common_font_style">{{formatAddress(row.From)}}
                                    <span class="address_hover_content">{{row.From}}</span>
                                </router-link>
                            </span>
                        </div>
                    </div>
                    <span class="no_skip" v-show="(/^[0]\d*$/).test(row.From) || row.From === '--'">--</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="To">
                <div class="common_hover_address_parent">
                    <div class="skip_route" style="display: flex" v-if="row.To !== '--'">
                        <div>
                            <span>
                                <router-link :to="`/address/1/${row.To}`" class="link_style common_font_style">{{formatAddress(row.To)}}
                                    <span class="address_hover_content">{{row.To}}</span>
                                </router-link>
                            </span>
                        </div>
                    </div>
                    <span class="no_skip" v-show="row.To === '--'">--</span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="Tx_Signer">
                <div class="common_hover_address_parent" v-if="row.Tx_Signer">
                    <router-link :to="`/address/1/${row.Tx_Signer}`" class="link_style common_font_style">{{formatAddress(row.Tx_Signer)}}
                        <div class="address_hover_content">
                            {{row.Tx_Signer}}
                        </div>
                    </router-link>
                </div>
            </template>
            <template slot-scope="{ row }" slot="OperatorAddr">
                <div class="common_hover_address_parent">
                    <router-link :to="`/address/1/${row.OperatorAddr}`" class="link_style common_font_style">{{formatAddress(row.OperatorAddr)}}
                        <div class="address_hover_content">
                            {{row.Tx_Signer}}
                        </div>
                    </router-link>
                </div>
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
                    },
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right'
					},
					{
						title:'From',
						slot: 'From',
					},
					{
						title:'Amount',
						slot: 'Amount',
						key: 'Amount',
						className: 'text_right'
					},
					{
						title:'To',
						slot: 'To',
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
                    },
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right',
					},
					{
						title:'Moniker',
						slot: 'Moniker',
						key: 'Moniker',
					},
					{
						title:'OperatorAddr',
						slot: 'OperatorAddr',
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
					},
					{
						title:'Block',
						slot: 'Block',
						className: 'text_right',
					},
					{
						title:'From',
						slot: 'From',
					},
					{
						title:'Amount',
						slot: 'Amount',
						key: 'Amount',
						className: 'text_right',
					},
					{
						title:'To',
						slot: 'To',
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
        .m-table-body{
            display: none;
        }
    }
    table{
        thead{
            tr{
                border-top: 0.01rem solid #d7d9e0;
            }
        }
    }

    .m-table-body tr{
        border: none !important;
    }
    .common_hover_parent{
        a{
            display: inline-block;
            position: relative;
            .hover_content{
                position: absolute;
                left: 0;
                height: 0.4rem;
                color:#fff;
                background: #000;
                border-radius: 0.04rem;
                display: none;
                padding: 0.1rem 0.15rem 0.06rem 0.15rem;
                z-index: 10;
            }
            &:hover{
                .hover_content{
                    position: absolute;
                    top: -0.4rem;
                    height: 0.4rem;
                    color:#fff;
                    background: #000;
                    border-radius: 0.04rem;
                    display: block;
                    &::after{
                        content: '';
                        display: block;
                        background: rgba(0,0,0,1);
                        transform: rotate(45deg);
                        width: 0;
                        height: 0;
                        border: 0.05rem solid transparent;
                        position: relative;
                        top: 0.04rem;
                        z-index: 1;
                        left: 0.14rem;
                    }
                }
            }
        }
    }
    .common_hover_address_parent{
        a{
            position: relative;
            .address_hover_content{
                display: none;
                position: absolute;
                height: 0.4rem;
                bottom: 0.19rem;
                left: -100%;
                color: #3598db;
                background: rgba(0,0,0,0.8);
                border-radius:0.04rem;
                z-index: 10;
            }
            &:hover{
                .address_hover_content{
                    background: rgba(0,0,0,1);
                    color: #fff;
                    padding: 0.1rem 0.15rem 0.06rem 0.15rem;
                    display: block;
                    border-radius:0.04rem;
                    font-size: 0.14rem;
                    &::after{
                        content: '';
                        display: block;
                        background: rgba(0,0,0,1);
                        transform: rotate(45deg);
                        width: 0;
                        height: 0;
                        border: 0.05rem solid transparent;
                        position: relative;
                        top: 0.04rem;
                        z-index: 1;
                        left: 1.23rem;
                    }
                }
            }
        }

    }
.common_font_style{
    font-family: Consolas;
}
</style>