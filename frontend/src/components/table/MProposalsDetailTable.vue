<template>
    <div>
        <m-table :columns="columns" :data="items">
            <template slot-scope="{ row }" slot="Voter">
                <div class="skip_route">
                    <div class="name_address">
                        <span class="remove_default_style">
                            <router-link
                                :to="addressRoute(row.Voter)"
                                class="link_style justify"
                            >{{row.moniker}}</router-link>
                        </span>
                    </div>
                </div>
            </template>
            <template slot="Tx_Hash" slot-scope="{ row }">
                <router-link
                    class="skip_route"
                    :to="`/tx?txHash=${row.Tx_Hash}`"
                >{{row.Tx_Hash ? `${formatTxHash(String(row.Tx_Hash))}` : ''}}</router-link>
            </template>
            <template slot="Depositor" slot-scope="{ row }">
                <span v-if="(/^[1-9]\d*$/).test(row.Depositor)" class="skip_route">
                    <router-link :to="`/tx?txHash=${row.Tx_Hash}`">{{row.Depositor}} Validators</router-link>
                </span>
                <div
                    class="name_address"
                    v-show="!(/^[0-9]\d*$/).test(row.Depositor) && row.Depositor && row.Depositor !== '--'"
                >
                    <span
                        class="remove_default_style"
                        :class="row.Depositor === $route.params.param?'no_skip':''"
                    >
                        <router-link
                            :to="addressRoute(row.Depositor)"
                            class="link_style"
                        >{{formatMoniker(row.moniker) || formatAddress(row.Depositor)}}</router-link>
                    </span>
                </div>
                <span
                    class="no_skip"
                    v-show="(/^[0]\d*$/).test(row.Depositor) || row.Depositor === '--'"
                >--</span>
            </template>
        </m-table>
    </div>
</template>
<script>
import Tools from "../../util/Tools";

export default {
    name: "MProposalsDetailTable",
    props: {
        items: {
            type: Array,
            default: function() {
                return [];
            }
        },
        fields: {
            type: String,
            default: "votersFields"
        }
    },
    data() {
        return {
            columns: [],
            votersFields: [
                {
                    title: "Voter",
                    slot: "Voter"
                },
                {
                    title: "Vote Option",
                    key: "Vote_Option"
                },
                {
                    title: "Tx Hash",
                    slot: "Tx_Hash",
                    tooltip: true
                },
                {
                    title: "Time",
                    key: "Time",
                    className: "text_right"
                }
            ],
            depositorsFields: [
                {
                    title: "Depositor",
                    slot: "Depositor"
                },
                {
                    title: "Amount",
                    key: "Amount"
                },
                {
                    title: "Type",
                    key: "Type"
                },
                {
                    title: "Tx Hash",
                    slot: "Tx_Hash",
                    tooltip: true
                },
                {
                    title: "Time",
                    key: "Time",
                    className: "text_right"
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
        },
        formatMoniker(moniker) {
            if (!moniker) {
                return "";
            }
            return Tools.formatString(moniker, 15, "...");
        }
    },
    mounted() {
        this.columns = this[this.fields];
    }
};
</script>

<style lang="scss" scoped>
.skip_route {
    color: #3598db !important;
}
.name_address {
    .remove_default_style {
        line-height: 16px;
        a {
            line-height: 16px;
        }
    }
}
</style>
