<template>
    <div class="validator_detail_table">
        <m-table :columns="fields" :data="items" :width="width">
            <template slot="Tx_Hash" slot-scope="{ row }">
                <div class="common_hover_parent" v-if="row.Tx_Hash">
                    <router-link
                        :to="`/tx?txHash=${row.Tx_Hash}`"
                        class="link_style common_font_style"
                    >{{formatTxHash(row.Tx_Hash)}}</router-link>
                </div>
            </template>
            <template slot="Block" slot-scope="{ row }">
                <span class="skip_route">
                    <router-link :to="`/block/${row.Block || row.block}`">{{row.Block || row.block}}</router-link>
                </span>
            </template>
            <template slot="From" slot-scope="{ row }">
                <span v-if="(/^[1-9]\d*$/).test(row.From)" class="skip_route">
                    <router-link :to="`/tx?txHash=${row.Tx_Hash}`">{{row.From}} Validators</router-link>
                </span>
                <div
                    class="name_address"
                    v-show="!(/^[0-9]\d*$/).test(row.From) && row.From && row.From !== '--'"
                >
                    <span
                        class="remove_default_style"
                        :class="row.From === $route.params.param?'no_skip':''"
                    >
                        <router-link
                            :to="addressRoute(row.From)"
                            class="link_style"
                        >{{formatAddress(row.From)}}</router-link>
                    </span>
                </div>
                <span class="no_skip" v-show="(/^[0]\d*$/).test(row.From) || row.From === '--'">--</span>
            </template>
            <template slot="address" slot-scope="{ row }">
                <span v-if="(/^[1-9]\d*$/).test(row.address)" class="skip_route">
                    <router-link :to="`/tx?txHash=${row.address}`">{{row.address}} Validators</router-link>
                </span>
                <div
                    class="name_address"
                    v-show="!(/^[0-9]\d*$/).test(row.address) && row.address && row.address !== '--'"
                >
                    <span
                        class="remove_default_style"
                        :class="row.address === $route.params.param?'no_skip':''"
                    >
                        <router-link
                            :to="addressRoute(row.address)"
                            class="link_style"
                        >{{formatAddress(row.address)}}</router-link>
                    </span>
                </div>
                <span
                    class="no_skip"
                    v-show="(/^[0]\d*$/).test(row.address) || row.From === '--'"
                >--</span>
            </template>
            <template slot="To" slot-scope="{ row }">
                <div class="name_address" v-show="row.To && row.To !== '--'">
                    <span
                        class="remove_default_style"
                        :class="row.To === $route.params.param?'no_skip':''"
                    >
                        <router-link
                            :to="addressRoute(row.To)"
                            class="link_style"
                        >{{formatAddress(row.To)}}</router-link>
                    </span>
                </div>
                <span class="no_skip" v-show="row.To == '--'">--</span>
            </template>
            <template slot="Tx_Signer" slot-scope="{ row }">
                <span class="skip_route" style="display: flex" v-if="row.Tx_Signer">
                    <div class="name_address">
                        <span
                            class="remove_default_style"
                            :class="row.Tx_Signer === $route.params.param?'no_skip':''"
                        >
                            <router-link
                                :to="addressRoute(row.Tx_Signer)"
                                class="link_style justify"
                            >{{formatAddress(row.Tx_Signer)}}</router-link>
                        </span>
                    </div>
                </span>
            </template>
        </m-table>
    </div>
</template>

<script>
import Tools from "../../util/Tools";
export default {
    name: "MValidatorDetailTable",
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
    data() {
        return {
            fields: [],
            transfer: [
                {
                    title: "Tx_Hash",
                    slot: "Tx_Hash",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Block",
                    slot: "Block"
                },
                {
                    title: "From",
                    slot: "From",
                    tooltip: true
                },
                {
                    title: "Amount",
                    key: "Amount"
                },
                {
                    title: "To",
                    slot: "To",
                    tooltip: true
                },
                {
                    title: "Tx_Type",
                    key: "Tx_Type"
                },
                {
                    title: "Tx_Fee",
                    key: "Tx_Fee"
                },
                {
                    title: "Tx_Signer",
                    slot: "Tx_Signer",
                    tooltip: true
                },
                {
                    title: "Tx_Status",
                    key: "Tx_Status"
                },
                {
                    title: "Timestamp",
                    key: "Timestamp"
                }
            ],
            stakes: [
                {
                    title: "Tx_Hash",
                    slot: "Tx_Hash",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Block",
                    slot: "Block"
                },
                {
                    title: "From",
                    slot: "From",
                    tooltip: true
                },
                {
                    title: "Amount",
                    key: "Amount"
                },
                {
                    title: "To",
                    slot: "To",
                    tooltip: true
                },
                {
                    title: "Tx_Type",
                    key: "Tx_Type"
                },
                {
                    title: "Tx_Fee",
                    key: "Tx_Fee"
                },
                {
                    title: "Tx_Signer",
                    slot: "Tx_Signer",
                    tooltip: true
                },
                {
                    title: "Tx_Status",
                    key: "Tx_Status"
                },
                {
                    title: "Timestamp",
                    key: "Timestamp"
                }
            ],
            declarations: [
                {
                    title: "Tx_Hash",
                    slot: "Tx_Hash",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Block",
                    slot: "Block"
                },
                {
                    title: "Moniker",
                    slot: "Moniker"
                },
                {
                    title: "Operator_Address",
                    key: "Operator_Address"
                },
                {
                    title: "Self_Bonded",
                    slot: "Self_Bonded"
                },
                {
                    title: "Tx_Type",
                    key: "Tx_Type"
                },
                {
                    title: "Tx_Fee",
                    key: "Tx_Fee"
                },
                {
                    title: "Tx_Signer",
                    slot: "Tx_Signer",
                    tooltip: true
                },
                {
                    title: "Tx_Status",
                    key: "Tx_Status"
                },
                {
                    title: "Timestamp",
                    key: "Timestamp"
                }
            ],
            gov: [
                {
                    title: "Tx_Hash",
                    slot: "Tx_Hash",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Block",
                    slot: "Block"
                },
                {
                    title: "Proposal_Type",
                    slot: "Proposal_Type"
                },
                {
                    title: "Proposal_ID",
                    key: "Proposal_ID"
                },
                {
                    title: "Proposal_Title",
                    slot: "Proposal_Title"
                },
                {
                    title: "Amount",
                    key: "Amount"
                },
                {
                    title: "Tx_Type",
                    key: "Tx_Type"
                },
                {
                    title: "Tx_Fee",
                    slot: "Tx_Fee"
                },
                {
                    title: "Tx_Signer",
                    key: "Tx_Signer"
                },
                {
                    title: "Tx_Status",
                    key: "Tx_Status"
                },
                {
                    title: "Timestamp",
                    key: "Timestamp"
                }
            ],
            delegations: [
                {
                    title: "Address",
                    slot: "address",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Amount",
                    key: "amount"
                },
                {
                    title: "Shares",
                    key: "shares"
                },
                {
                    title: "Block",
                    slot: "Block"
                }
            ],
            unbondingDelegations: [
                {
                    title: "Address",
                    slot: "address",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Amount",
                    key: "amount"
                },
                {
                    title: "Block",
                    slot: "Block"
                },
                {
                    title: "Timestamp",
                    key: "Timestamp"
                }
            ],
            redelegations: [
                {
                    title: "Address",
                    slot: "address",
                    tooltip: true,
                    tooltipClassName: "tooltip_left"
                },
                {
                    title: "Amount",
                    key: "amount"
                },
                {
                    title: "To",
                    slot: "To",
                    tooltip: true
                },
                {
                    title: "Block",
                    slot: "Block"
                }
            ],
            depositedProposals: [
                {
                    title: "ID",
                    key: "proposal_id"
                },
                {
                    title: "Proposer",
                    slot: "address",
                    tooltip: true
                },
                {
                    title: "Deposit",
                    key: "depositedAmount"
                },
                {
                    title: "Submited",
                    key: "submited"
                },
                {
                    title: "Tx_Hash",
                    slot: "Tx_Hash",
                    tooltip: true
                }
            ],
            votedProposals: [
                {
                    title: "ID",
                    key: "proposal_id"
                },
                {
                    title: "Title",
                    key: "title",
                    className: "text_overflow"
                },
                {
                    title: "Status",
                    key: "status"
                },
                {
                    title: "Voted",
                    key: "voted"
                },
                {
                    title: "Tx_Hash",
                    slot: "Tx_Hash",
                    tooltip: true
                }
            ]
        };
    },
    methods: {
        formatTxHash(TxHash) {
            if (TxHash) {
                return Tools.formatTxHash(TxHash);
            }
        },
        formatAddress(address) {
            return Tools.formatValidatorAddress(address);
        }
    },
    mounted() {
        this.fields = this[this.listName] || [];
    }
};
</script>

<style lang="scss">
.personal_computer_transactions_detail_wrap {
    .validator_detail_table {
        .m-table-header {
            width: 100%;
            table.m_table {
                width: auto;
            }
        }
        table.m_table {
            min-width: 100%;
        }
    }
    .delegations_two_container {
        .m-table-header {
            width: 6.3rem;
        }
        table.m_table {
            min-width: 6.3rem;
        }
    }
}
</style>
