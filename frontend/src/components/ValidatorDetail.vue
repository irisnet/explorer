<template>
    <div class="transactions_detail_wrap">
        <div :class="[transactionsDetailWrap, 'validator_title']">
            <span class="title">{{validatorName}}</span>
            <div class="status_btn" v-if="validatorStatus === 'Active'">Active</div>
            <div
                class="status_btn"
                style="background-color: #3DA87E;"
                v-if="validatorStatus === 'Candidate'"
            >Candidates</div>
            <div
                class="status_btn"
                style="background-color: #FA7373;"
                v-if="validatorStatus === 'Jailed'"
            >Jailed</div>
        </div>
        <div :class="transactionsDetailWrap">
            <div class="validator_detail_information_wrap">
                <div
                    class="information_props_wrap"
                    v-for="v in Object.entries(validatorInfo)"
                    :key="v[0]"
                >
                    <span class="information_props">{{v[0]}}:</span>
                    <template v-if="v[1]">
                        <span class="information_value">{{v[1]}}</span>
                    </template>
                    <template v-else>
                        <span class="information_value">--</span>
                    </template>
                </div>
            </div>
        </div>
        <div :class="transactionsDetailWrap" class="address_profile">
            <p class="validator_information_content_title">Validator Profile</p>
            <div class="validator_detail_information_wrap">
                <div
                    class="information_props_wrap"
                    v-for="v in Object.entries(validatorProfile)"
                    :key="v[0]"
                >
                    <span class="information_props">{{v[0]}}:</span>
                    <template v-if="v[1] && validatorProfileLinks.indexOf(v[0]) > -1">
                        <span class="information_value skip_route">
                            <router-link
                                v-if="v[0] !== 'Website' && v[0] !== 'Identity'"
                                :to="addressRoute(v[1])"
                            >{{v[1]}}</router-link>
                            <span @click="openUrl(v[1])" v-if="v[0] === 'Website'">{{v[1]}}</span>
                            <span v-if="v[0] === 'Identity'">{{v[1]}}</span>
                        </span>
                    </template>
                    <template v-else>
                        <span v-if="!v[1]" class="information_value">--</span>
                        <span v-if="v[1]" class="information_value">{{v[1]}}</span>
                    </template>
                </div>
            </div>
        </div>
        <div class="line_container_wrap" v-show="flActiveValidator">
            <div class="line_container" :class="transactionsDetailWrap">
                <p class="validator_information_content_title">History</p>
                <div class="line_content">
                    <div class="line_echarts_content">
                        <div
                            class="line_left_container"
                            style="overflow-x: auto;-webkit-overflow-scrolling:touch;"
                        >
                            <echarts-validators-line
                                :informationValidatorsLine="informationValidatorsLine"
                            >
                                <div slot="tabs" class="line_tab_content_new">
                                    <div
                                        v-for="(item,index) in tabVotingPower"
                                        :key="index"
                                        @click="getValidatorHistory(item.title,index)"
                                        :class="item.active ? 'active' : ''"
                                    >{{item.title}}</div>
                                </div>
                            </echarts-validators-line>
                        </div>
                    </div>
                    <div
                        class="line_echarts_content"
                        :class="transactionsDetailWrap === 'personal_computer_transactions_detail_wrap' ?
           'content_right' : 'model_content_right' "
                    >
                        <div
                            class="line_right_container"
                            style="overflow-x: auto;-webkit-overflow-scrolling:touch;"
                        >
                            <echarts-validators-uptime-line
                                :informationUptimeLine="informationUptimeLine"
                            >
                                <div slot="tabs" class="line_tab_content_new">
                                    <div
                                        v-for="(item,index) in tabUptime"
                                        :key="index"
                                        @click="getValidatorUptimeHistory(item.title,index)"
                                        :class="item.active ? 'active' : ''"
                                    >{{item.title}}</div>
                                </div>
                            </echarts-validators-uptime-line>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div :class="transactionsDetailWrap">
            <div
                :class="[!$store.state.isMobile && txTableList.delegations.total && txTableList.unbondingDelegations.total ? 'delegations_two_container': '']"
                class="delegations_container"
            >
                <div v-if="txTableList.delegations.total > 0">
                    <p class="validator_information_content_title">Delegations</p>
                    <div class="delegations_table_container">
                        <m-validator-detail-table
                            :items="txTableList.delegations.items"
                            listName="delegations"
                            :width="txTableList.delegations.total && txTableList.unbondingDelegations.total ? 630 : 1280"
                        ></m-validator-detail-table>
                    </div>
                    <m-pagination
                        v-if="txTableList.delegations.total > pageSize"
                        :page-size="pageSize"
                        :total="txTableList.delegations.total"
                        :page-change="pageChange('getDelegations')"
                    ></m-pagination>
                </div>
                <div
                    v-if="txTableList.unbondingDelegations.total > 0"
                    class="second_table_container"
                >
                    <p class="validator_information_content_title">Unbonding Delegations</p>
                    <div class="delegations_table_container">
                        <m-validator-detail-table
                            :items="txTableList.unbondingDelegations.items"
                            listName="unbondingDelegations"
                            :width="txTableList.delegations.total && txTableList.unbondingDelegations.total ? 630 : 1280"
                        ></m-validator-detail-table>
                    </div>
                    <m-pagination
                        v-if="txTableList.unbondingDelegations.total > pageSize"
                        :page-size="pageSize"
                        :total="txTableList.unbondingDelegations.total"
                        :page-change="pageChange('getUnbondingDelegations')"
                    ></m-pagination>
                </div>
            </div>
        </div>

        <div
            :class="transactionsDetailWrap"
            class="validator_table_container"
            v-if="txTableList.redelegations.total > 0"
        >
            <p class="validator_information_content_title">Redelegations</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validator-detail-table
                        :items="txTableList.redelegations.items"
                        listName="redelegations"
                    ></m-validator-detail-table>
                </div>
            </div>
            <m-pagination
                v-if="txTableList.redelegations.total > pageSize"
                :page-size="pageSize"
                :total="txTableList.redelegations.total"
                :page-change="pageChange('getRedelegations')"
            ></m-pagination>
        </div>
        <div :class="transactionsDetailWrap">
            <div
                :class="[!$store.state.isMobile && txTableList.depositedProposals.total && txTableList.votedProposals.total ? 'delegations_two_container': '']"
                class="delegations_container"
            >
                <div v-if="txTableList.depositedProposals.total > 0">
                    <p class="validator_information_content_title">Deposited Proposals</p>
                    <div class="delegations_table_container">
                        <m-validator-detail-table
                            :items="txTableList.depositedProposals.items"
                            listName="depositedProposals"
                            :width="txTableList.depositedProposals.total && txTableList.votedProposals.total ? 630 : 1280"
                        ></m-validator-detail-table>
                    </div>
                    <m-pagination
                        v-if="txTableList.depositedProposals.total > pageSize"
                        :page-size="pageSize"
                        :total="txTableList.depositedProposals.total"
                        :page-change="pageChange('getDepositedProposals')"
                    ></m-pagination>
                </div>
                <div v-if="txTableList.votedProposals.total > 0" class="second_table_container">
                    <p class="validator_information_content_title">Voted Proposals</p>
                    <div class="delegations_table_container">
                        <m-validator-detail-table
                            :items="txTableList.votedProposals.items"
                            listName="votedProposals"
                            :width="txTableList.depositedProposals.total && txTableList.votedProposals.total ? 630 : 1280"
                        ></m-validator-detail-table>
                    </div>
                    <m-pagination
                        v-if="txTableList.votedProposals.total > pageSize"
                        :page-size="pageSize"
                        :total="txTableList.votedProposals.total"
                        :page-change="pageChange('getVotedProposals')"
                    ></m-pagination>
                </div>
            </div>
        </div>

        <div
            :class="transactionsDetailWrap"
            class="validator_table_container"
            v-if="txTableList.transfers.total > 0"
        >
            <p class="validator_information_content_title">Transfers</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validator-detail-table
                        :items="txTableList.transfers.items"
                        listName="transfer"
                    ></m-validator-detail-table>
                </div>
            </div>
            <m-pagination
                v-if="txTableList.transfers.total > pageSize"
                :page-size="pageSize"
                :total="txTableList.transfers.total"
                :page-change="pageChange('getTransfers')"
            ></m-pagination>
        </div>
        <div
            :class="transactionsDetailWrap"
            class="validator_table_container"
            v-if="txTableList.stakes.total > 0"
        >
            <p class="validator_information_content_title">Stakes</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validator-detail-table :items="txTableList.stakes.items" listName="stakes"></m-validator-detail-table>
                </div>
            </div>
            <m-pagination
                v-if="txTableList.stakes.total > pageSize"
                :page-size="pageSize"
                :total="txTableList.stakes.total"
                :page-change="pageChange('getStakes')"
            ></m-pagination>
        </div>
        <div
            :class="transactionsDetailWrap"
            class="validator_table_container"
            v-if="txTableList.declarations.total > 0"
        >
            <p class="validator_information_content_title">Declarations</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validator-detail-table
                        :items="txTableList.declarations.items"
                        listName="declarations"
                    ></m-validator-detail-table>
                </div>
            </div>
            <m-pagination
                v-if="txTableList.declarations.total > pageSize"
                :page-size="pageSize"
                :total="txTableList.declarations.total"
                :page-change="pageChange('getDeclarations')"
            ></m-pagination>
        </div>
        <div
            :class="transactionsDetailWrap"
            class="validator_table_container"
            v-if="txTableList.governance.total > 0"
        >
            <p class="validator_information_content_title">Governance</p>
            <div class="blocks_list_table_container">
                <!-- <spin-component :showLoading="showLoading" /> -->
                <div class="address_tx_list_table">
                    <m-validator-detail-table :items="txTableList.governance.items" listName="gov"></m-validator-detail-table>
                </div>
            </div>
            <m-pagination
                v-if="txTableList.governance.total > pageSize"
                :page-size="pageSize"
                :total="txTableList.governance.total"
                :page-change="pageChange('getGovernance')"
            ></m-pagination>
        </div>
    </div>
</template>

<script>
import Tools from "../util/Tools";
import EchartsLine from "./EchartsLine";
import EchartsValidatorsLine from "./EchartsValidatorsLine";
import EchartsValidatorsUptimeLine from "./EchartsValidatorsUpTimeLine";
import SpinComponent from "./commonComponents/SpinComponent";
import Services from "../util/axios";
import Service from "../service";
import Constants from "../constant/Constant";
import MValidatorDetailTable from "./table/MValidatorDetailTable";
import MPagination from "./commonComponents/MPagination";
export default {
    watch: {
        "$store.state.isMobile"(newVal) {
            this.isMobileFunc(newVal);
        }
    },
    data() {
        return {
            rateValue: "",
            transactionsDetailWrap: "personal_computer_transactions_detail",
            currentPage: 1,
            pageSize: 5,
            firstPercent: "",
            address: this.$route.params.param,
            balanceValue: "",
            depositsValue: "",
            transactionsValue: "",
            nameValue: "",
            websiteValue: "",
            commissionRateValue: "",
            announcementValue: "",
            bondHeightValue: "",
            votingPowerValue: "",
            precommitedBlocksValue: "",
            returnsValue: "",
            operatorValue: "",
            TransactionsShowNoData: false,
            PrecommitBlocksshowNoData: false,
            transactionsCount: 0,
            flValidator: false,
            showLoading: false,
            informationValidatorsLine: {},
            informationUptimeLine: {},
            txTabName: "Transfers",
            currentTxTabName: "",
            withdrawAddress: "",
            flShowValidatorJailed: false,
            flShowValidatorCandidate: false,
            flActiveValidator: true,
            flShowProfileLogo: false,
            validatorStatusColor: "#3598db",
            tabVotingPower: [
                {
                    title: "14days",
                    active: true
                },
                {
                    title: "30days",
                    active: false
                },
                {
                    title: "60days",
                    active: false
                }
            ],
            tabUptime: [
                {
                    title: "24hours",
                    active: true
                },
                {
                    title: "14days",
                    active: false
                },
                {
                    title: "30days",
                    active: false
                }
            ],
            validatorInfo: {
                "Voting Power": "",
                "Bond Height": "",
                "Bonded Tokens": "",
                "Unbonding Height": "",
                "Self Bonded": "",
                "Jailed Until": "",
                "Delegator Bonded": "",
                "Missed Blocks": "",
                Delegators: "",
                "Commission Rate": "",
                "Delegator Shares": "",
                "Max Rate": "",
                Rewards: "",
                "Max Change Rate": ""
            },
            validatorProfile: {
                "Operator Address": "",
                Website: "",
                "Owner Address": "",
                Identity: "",
                "Withdraw Address": "",
                Details: "",
                "Consensus Pubkey": ""
            },
            validatorProfileLinks: [
                "Website",
                "Identity",
                "Owner Address",
                "Withdraw Address"
            ],
            txTableList: {
                delegations: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                unbondingDelegations: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                redelegations: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                depositedProposals: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                votedProposals: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                transfers: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                stakes: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                declarations: {
                    items: [],
                    currentPage: 1,
                    total: 0
                },
                governance: {
                    items: [],
                    currentPage: 1,
                    total: 0
                }
            },
            validatorName: "",
            validatorStatus: ""
        };
    },
    components: {
        EchartsValidatorsUptimeLine,
        EchartsValidatorsLine,
        EchartsLine,
        SpinComponent,
        MValidatorDetailTable,
        MPagination
    },
    beforeMount() {
        this.isMobileFunc(this.$store.state.isMobile);
    },
    mounted() {
        Tools.scrollToTop();
        if (
            sessionStorage.getItem("currentEnv") !== Constants.ENVCONFIG.MAINNET
        ) {
            this.$Crypto.getCrypto(
                Constants.CHAINNAME,
                Constants.ENVCONFIG.TESTNET
            );
        }
        this.getValidatorsInfo();
        this.getDelegations();
        this.getUnbondingDelegations();
        this.getRedelegations();
        this.getDepositedProposals();
        this.getVotedProposals();

        this.getStakes();
        this.getDeclarations();
        this.getGovernance();
        this.getValidatorHistory("14days");
        this.getValidatorUptimeHistory("24hours");
        // if (
        //     this.$route.params.param.substring(0, 3) ===
        //     this.$Crypto.config.iris.bech32.valAddr
        // ) {
        // this.getAddressInformation(this.$route.params.param);
        // this.getTransactionsList();
        // this.getProfileInformation();
        // this.getCurrentTenureInformation();
        // } else {
        // this.getAddressInformation(this.$route.params.param);
        // this.getTransactionsList();
        // this.getProfileInformation();
        // }
    },
    methods: {
        isMobileFunc(isMobile) {
            if (isMobile) {
                this.transactionsDetailWrap = "mobile_transactions_detail_wrap";
            } else {
                this.transactionsDetailWrap =
                    "personal_computer_transactions_detail_wrap";
            }
        },
        pageChange(key) {
            return page => {
                this[key](page);
            };
        },
        getValidatorsInfo() {
            Service.commonInterface(
                {
                    validatorsInfo: {
                        address: this.$route.params.param
                    }
                },
                data => {
                    try {
                        if (data) {
                            this.validatorStatus = data.status;
                            this.validatorInfo["Voting Power"] =
                                data.status === "Active"
                                    ? `${Tools.formatDecimalNumberToFixedNumber(
                                          (data.self_power / data.total_power) *
                                              100
                                      )} %`
                                    : data.status;
                            this.validatorInfo["Bond Height"] =
                                data.bond_height;
                            this.validatorInfo[
                                "Bonded Tokens"
                            ] = `${Tools.formatPriceToFixed(
                                Number(data.bonded_tokens),
                                2
                            )} IRIS`;
                            this.validatorInfo["Unbonding Height"] =
                                data.unbond_height;
                            this.validatorInfo[
                                "Self Bonded"
                            ] = `${Tools.formatPriceToFixed(
                                Number(data.self_bonded),
                                2
                            )} IRIS`;
                            this.validatorInfo["Jailed Until"] = new Date(
                                data.jailed_until
                            ).getTime()
                                ? Tools.format2UTC(data.jailed_until)
                                : "";
                            let delegator_tokens =
                                Number(data.bonded_tokens) -
                                Number(data.self_bonded);
                            this.validatorInfo[
                                "Delegator Bonded"
                            ] = `${Tools.formatPriceToFixed(
                                Number(delegator_tokens),
                                2
                            )} IRIS (${Tools.formatDecimalNumberToFixedNumber(
                                (delegator_tokens /
                                    Number(data.bonded_tokens)) *
                                    100
                            )} %)`;
                            this.validatorInfo["Missed Blocks"] =
                                data.missed_blocks_count;
                            this.validatorInfo["Delegators"] =
                                data.delegator_count;
                            this.validatorInfo["Commission Rate"] =
                                data.commission_update !==
                                "0001-01-01 00:00:00 +0000 UTC"
                                    ? `${Tools.formatDecimalNumberToFixedNumber(
                                          Number(data.commission_rate) * 100
                                      )} % (${Tools.format2UTC(
                                          data.commission_update
                                      )})`
                                    : `${Tools.formatDecimalNumberToFixedNumber(
                                          Number(data.commission_rate) * 100
                                      )} %`;
                            this.validatorInfo[
                                "Delegator Shares"
                            ] = `${Tools.formatPriceToFixed(
                                Number(data.delegator_shares),
                                2
                            )} IRIS`;
                            this.validatorInfo[
                                "Max Rate"
                            ] = `${Tools.formatDecimalNumberToFixedNumber(
                                Number(data.commission_max_rate) * 100
                            )} %`;
                            this.validatorInfo["Rewards"] = this.formatAmount(
                                data.reward
                            );
                            this.validatorInfo[
                                "Max Change Rate"
                            ] = `${Tools.formatDecimalNumberToFixedNumber(
                                Number(data.commision_max_change_rate) * 100
                            )} %`;
                            this.validatorProfile["Withdraw Address"] =
                                data.withdraw_addr;
                            this.validatorProfile["Operator Address"] =
                                data.operator_addr;
                            this.validatorProfile["Owner Address"] =
                                data.owner_addr;
                            this.validatorProfile["Consensus Pubkey"] =
                                data.consensus_addr;
                            if (data.description) {
                                this.validatorProfile["Website"] =
                                    data.description.website;

                                this.validatorProfile["Identity"] =
                                    data.description.identity;

                                this.validatorProfile["Details"] =
                                    data.description.details;
                                this.validatorName =
                                    data.description.moniker ||
                                    this.$route.params.param;
                            }
                        }
                    } catch (e) {}
                }
            );
        },
        getDelegations(page) {
            Service.commonInterface(
                {
                    validatorDelegations: {
                        validatorAddr: this.$route.params.param,
                        currentPage: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (Array.isArray(data.items)) {
                            for (let it of data.items) {
                                it.amount =
                                    this.formatAmount(it).split(" ")[0] +
                                    " IRIS";
                                let selfShares = Number(it.self_shares);
                                it.shares = `${selfShares} (${Tools.formatDecimalNumberToFixedNumber(
                                    (selfShares / Number(it.total_shares)) * 100
                                )}%)`;
                            }
                            this.txTableList.delegations.items = data.items;
                        } else {
                            this.txTableList.delegations.items = [];
                        }
                        this.txTableList.delegations.total = data.total || 0;
                    } catch (e) {}
                }
            );
        },
        getUnbondingDelegations(page) {
            Service.commonInterface(
                {
                    validatorUnbondingDelegations: {
                        validatorAddr: this.$route.params.param,
                        currentPage: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (Array.isArray(data.items)) {
                            for (let it of data.items) {
                                let u = it.amount.match(/[a-zA-Z]+/g, "")[0];
                                it.amount = u
                                    ? `${it.amount.replace(
                                          u,
                                          ""
                                      )} ${u.toUpperCase()}`
                                    : it.amount;
                                it.Timestamp = Tools.format2UTC(it.until);
                                it.Block = it.block;
                            }
                            this.txTableList.unbondingDelegations.items =
                                data.items;
                        } else {
                            this.txTableList.unbondingDelegations.items = [];
                        }
                        this.txTableList.unbondingDelegations.total =
                            data.total || 0;
                    } catch (e) {}
                }
            );
        },
        getRedelegations(page) {
            Service.commonInterface(
                {
                    validatorRedelegations: {
                        validatorAddr: this.$route.params.param,
                        currentPage: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (Array.isArray(data.items)) {
                            for (let it of data.items) {
                                let u = it.amount.match(/[a-zA-Z]+/g, "")[0];
                                it.amount = u
                                    ? `${it.amount.replace(
                                          u,
                                          ""
                                      )} ${u.toUpperCase()}`
                                    : it.amount;
                                it.To = it.to;
                            }
                            this.txTableList.redelegations.items = data.items;
                        } else {
                            this.txTableList.redelegations.items = [];
                        }
                        this.txTableList.redelegations.total = data.total || 0;
                    } catch (e) {}
                }
            );
        },
        getDepositedProposals(page) {
            Service.commonInterface(
                {
                    validatorDepositorTxs: {
                        validatorAddr: this.$route.params.param,
                        currentPage: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (Array.isArray(data.items)) {
                            for (let it of data.items) {
                                it.Tx_Hash = it.tx_hash;
                                it.submited = it.submited ? "Yes" : "No";
                                it.address = it.proposer;

                                it.depositedAmount = this.formatAmount(
                                    it.deposited_amount
                                );
                            }
                            this.txTableList.depositedProposals.items =
                                data.items;
                        } else {
                            this.txTableList.depositedProposals.items = [];
                        }

                        this.txTableList.depositedProposals.total =
                            data.total || 0;
                    } catch (e) {}
                }
            );
        },
        getVotedProposals(page) {
            Service.commonInterface(
                {
                    validatorVote: {
                        validatorAddr: this.$route.params.param,
                        currentPage: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (Array.isArray(data.items)) {
                            for (let it of data.items) {
                                it.Tx_Hash = it.tx_hash;
                                it.submited = it.submited ? "Yes" : "No";
                            }
                            this.txTableList.votedProposals.items = data.items;
                        } else {
                            this.txTableList.votedProposals.items = [];
                        }
                        this.txTableList.votedProposals.total = data.total || 0;
                    } catch (e) {}
                }
            );
        },
        getTransfers(page) {
            Service.commonInterface(
                {
                    addressTxTrans: {
                        address: this.$route.params.param,
                        pageNumber: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (Array.isArray(data.Data)) {
                            for (let it of data.Data) {
                                it.Tx_Hash = it.Hash;
                            }
                            this.txTableList.transfers.items = Tools.formatTxList(
                                data.Data,
                                "transfers"
                            );
                        } else {
                            this.txTableList.transfers.items = [];
                        }
                        this.txTableList.transfers.total = data.Count || 0;
                    } catch (e) {}
                }
            );
        },
        getStakes(page) {
            Service.commonInterface(
                {
                    addressTxStake: {
                        address: this.$route.params.param,
                        pageNumber: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (data) {
                            if (Array.isArray(data.Data)) {
                                for (let it of data.Data) {
                                    it.Tx_Hash = it.Hash;
                                }
                                this.txTableList.stakes.items = Tools.formatTxList(
                                    data.Data,
                                    "stakes"
                                );
                            } else {
                                this.txTableList.stakes.items = [];
                            }
                            this.txTableList.stakes.total = data.Count || 0;
                        }
                    } catch (e) {}
                }
            );
        },
        getDeclarations(page) {
            Service.commonInterface(
                {
                    addressTxDeclaration: {
                        address: this.$route.params.param,
                        pageNumber: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (data) {
                            if (Array.isArray(data.Data)) {
                                for (let it of data.Data) {
                                    it.Tx_Hash = it.Hash;
                                }
                                this.txTableList.declarations.items = Tools.formatTxList(
                                    data.Data,
                                    "declarations"
                                );
                            } else {
                                this.txTableList.declarations.items = [];
                            }
                            this.txTableList.declarations.total =
                                data.Count || 0;
                        }
                    } catch (e) {}
                }
            );
        },
        getGovernance(page) {
            Service.commonInterface(
                {
                    addressTxGov: {
                        address: this.$route.params.param,
                        pageNumber: page || 1,
                        pageSize: this.pageSize
                    }
                },
                data => {
                    try {
                        if (data) {
                            if (Array.isArray(data.Data)) {
                                for (let it of data.Data) {
                                    it.Tx_Hash = it.Hash;
                                }
                                this.txTableList.governance.items = Tools.formatTxList(
                                    data.Data,
                                    "governance"
                                );
                            } else {
                                this.txTableList.governance.items = [];
                            }
                            this.txTableList.governance.total = data.Count || 0;
                        }
                    } catch (e) {}
                }
            );
        },
        formatAmount(item) {
            let amount = "--";
            if (item) {
                if (item instanceof Array) {
                    if (item.length > 0) {
                        amount = Tools.formatAmount(item[0].amount);
                        if (!item[0].denom) {
                            amount = item
                                .map(
                                    listItem =>
                                        `${listItem.formatAmount} SHARES`
                                )
                                .join(",");
                        } else {
                            amount =
                                amount + "" + Tools.formatDenom(item[0].denom);
                        }
                    }
                } else if (item.amount) {
                    if (!item.denom) {
                        amount = `${Tools.formatPriceToFixed(
                            item.amount
                        )} SHARES`;
                    } else {
                        amount = `${Tools.formatPriceToFixed(
                            item.amount
                        )}  ${Tools.formatDenom(item.denom).toUpperCase()}`;
                    }
                }
            }
            return amount;
        },
        getAddressInformation(address) {
            this.address = address;
            let url = `/api/account/${this.$route.params.param}`;
            Services.http(url)
                .then(result => {
                    let Amount = "--";
                    if (result) {
                        if (
                            result.amount &&
                            result.amount instanceof Array &&
                            result.amount[0].denom === Constants.Denom.IRISATTO
                        ) {
                            result.amount[0].amount = Tools.formatNumber(
                                result.amount[0].amount
                            );
                            Amount = result.amount
                                .map(
                                    listItem =>
                                        `${Tools.formatPriceToFixed(
                                            listItem.amount
                                        )} ${Tools.formatDenom(
                                            listItem.denom
                                        ).toUpperCase()}`
                                )
                                .join(",");
                        } else if (
                            result.amount &&
                            result.amount instanceof Array &&
                            result.amount[0].denom === Constants.Denom.IRIS
                        ) {
                            Amount = result.amount
                                .map(
                                    listItem =>
                                        `${Tools.formatPriceToFixed(
                                            listItem.amount
                                        )} ${Tools.formatDenom(
                                            listItem.denom
                                        ).toUpperCase()}`
                                )
                                .join(",");
                        } else if (
                            result.amount &&
                            Object.keys(result.amount).includes("amount") &&
                            Object.keys(result.amount).includes("denom") &&
                            result.amount.denom === Constants.Denom.IRISATTO
                        ) {
                            Amount = `${Tools.formatPriceToFixed(
                                Tools.formatNumber(result.amount)
                            )} ${result.denom.toUpperCase()}`;
                        } else if (
                            result.amount &&
                            Object.keys(result.amount).includes("amount") &&
                            Object.keys(result.amount).includes("denom") &&
                            result.amount.denom === Constants.Denom.IRIS
                        ) {
                            Amount = `${Tools.formatPriceToFixed(
                                result.amount
                            )} ${result.denom.toUpperCase()}`;
                        } else {
                            Amount = "";
                        }
                    }
                    this.flShowProfileLogo = result && result.isProfiler;
                    if (result && result.deposits) {
                        let deposits = Tools.formatToken(result.deposits);
                        this.depositsValue =
                            result.deposits.amount &&
                            result.deposits.amount !== 0 &&
                            result.deposits.amount !== ""
                                ? `${Tools.formatPriceToFixed(
                                      deposits.amount
                                  )} ${deposits.denom}`
                                : "";
                    }
                    this.withdrawAddress =
                        result && result.withdrawAddress
                            ? result.withdrawAddress
                            : "";
                    this.balanceValue = Amount;
                })
                .catch(e => {});
        },
        getProfileInformation() {
            let url = `/api/stake/candidate/${this.$route.params.param}`;
            Services.http(url)
                .then(result => {
                    if (result) {
                        let validator = result.validator;
                        this.flValidator = true;
                        if (validator.jailed === true) {
                            this.flActiveValidator = false;
                            this.flShowValidatorJailed = true;
                            this.validatorStatusColor = "#f00";
                        } else {
                            if (
                                validator.status === "Unbonded" ||
                                validator.status === "Unbonding"
                            ) {
                                this.flShowValidatorCandidate = true;
                                this.validatorStatusColor = "#45B035";
                                this.flActiveValidator = false;
                            } else if (validator.status === "Bonded") {
                                this.bondHeightValue = validator.bond_height;
                                this.votingPowerValue = validator.voting_power;
                            }
                        }
                        this.rateValue = validator.rate
                            ? `${Tools.formatRate(validator.rate.toString())}%`
                            : "";
                        this.validatorProfile["Identity"] =
                            validator.description &&
                            validator.description.identity
                                ? validator.description.identity
                                : "";
                        this.nameValue =
                            validator.description &&
                            validator.description.moniker
                                ? validator.description.moniker
                                : "";
                        this.validatorProfile[
                            "Consensus Pubkey"
                        ] = validator.pub_key ? validator.pub_key : "";
                        this.validatorProfile["Website"] =
                            validator.description &&
                            validator.description.website
                                ? validator.description.website
                                : "";
                        this.validatorProfile["Details"] =
                            validator.description &&
                            validator.description.details
                                ? validator.description.details
                                : "";
                        this.commissionRateValue = "";
                        this.announcementValue = "";
                        this.validatorProfile["Owner Address"] =
                            validator.owner;
                    } else {
                        this.flValidator = false;
                        this.flActiveValidator = false;
                    }
                })
                .catch(err => {
                    this.flActiveValidator = false;
                });
        },
        getCurrentTenureInformation() {
            let url = `/api/stake/candidate/${this.$route.params.param}/status`;
            Services.http(url)
                .then(data => {
                    if (data) {
                        this.precommitedBlocksValue = data.precommit_cnt
                            ? data.precommit_cnt
                            : "";
                        this.returnsValue = "";
                        this.firstPercent = data.up_time
                            ? `${data.up_time}%`
                            : "";
                    }
                })
                .catch(err => {});
        },
        getTransactionsList() {
            let url = `/api/txsByAddress/${this.$route.params.param}/1/30`;
            Services.http(url)
                .then(data => {
                    if (data) {
                        this.transactionsCount = data.Count;
                        this.transactionsValue = data.Count;
                        this.TransactionsShowNoData = true;
                    }
                })
                .catch(e => {});
        },
        getValidatorHistory(tabTime, index) {
            if (index !== undefined) {
                for (var i = 0; i < this.tabVotingPower.length; i++) {
                    this.tabVotingPower[i].active = false;
                    this.tabVotingPower[index].active = true;
                }
            }
            let url;
            if (tabTime == "14days") {
                url = `/api/stake/candidate/${
                    this.$route.params.param
                }/power/week`;
            } else if (tabTime == "30days") {
                url = `/api/stake/candidate/${
                    this.$route.params.param
                }/power/month`;
            } else if (tabTime == "60days") {
                url = `/api/stake/candidate/${
                    this.$route.params.param
                }/power/months`;
            }
            Services.http(url).then(validatorVotingPowerList => {
                if (validatorVotingPowerList) {
                    let seriesData = [],
                        noDatayAxisDefaultMaxByValidators;
                    let maxPowerValue = 0;
                    validatorVotingPowerList.forEach(item => {
                        if (item.power > maxPowerValue) {
                            maxPowerValue = item.power;
                        }
                        if (item.power == 0) {
                            item.power = "";
                        }
                        let obj = [];
                        obj[0] = Tools.conversionTimeToUTCByValidatorsLine(
                            item.time
                        );
                        obj[1] = item.power;
                        seriesData.push(obj);
                    });
                    if (maxPowerValue < 100) {
                        noDatayAxisDefaultMaxByValidators = "100";
                    }
                    this.informationValidatorsLine = {
                        seriesData,
                        noDatayAxisDefaultMaxByValidators
                    };
                }
            });
        },
        getValidatorUptimeHistory(tabTime, index) {
            if (index != undefined) {
                for (var i = 0; i < this.tabUptime.length; i++) {
                    this.tabUptime[i].active = false;
                    this.tabUptime[index].active = true;
                }
            }
            let url;
            if (tabTime == "24hours") {
                url = `/api/stake/candidate/${
                    this.$route.params.param
                }/uptime/hour `;
            } else if (tabTime == "14days") {
                url = `/api/stake/candidate/${
                    this.$route.params.param
                }/uptime/week `;
            } else if (tabTime == "30days") {
                url = `/api/stake/candidate/${
                    this.$route.params.param
                }/uptime/month `;
            }
            Services.http(url)
                .then(data => {
                    if (data) {
                        data.forEach(item => {
                            let notValidatorTag = -1;
                            if (item.uptime === notValidatorTag) {
                                item.uptime = "";
                            }
                        });
                        let xData, currayDate;
                        if (tabTime == "24hours") {
                            data.forEach(item => {
                                //兼容火狐
                                let hourseconds =
                                    item.time
                                        .replace("-", "/")
                                        .replace("-", "/") + ":00:00";

                                let changeMilliseconds = new Date(
                                    hourseconds
                                ).getTime();
                                //展示需加一小时
                                changeMilliseconds =
                                    changeMilliseconds + 60 * 60 * 1000;
                                item.time = Tools.formatDateYearAndMinutesAndSeconds(
                                    changeMilliseconds
                                ).split(":")[0];
                            });
                            if (data.length !== 0) {
                                currayDate = data[0].time;
                            } else {
                                currayDate = new Date()
                                    .toISOString()
                                    .substr(0, 13)
                                    .replace("T", " ");
                            }
                            if (data.length < 24) {
                                let complementHourLength = 24 - data.length;
                                let hourTime = currayDate.split(" ")[1];
                                let yearAndDayTime = currayDate.split(" ")[0];

                                for (let k = 0; k < complementHourLength; k++) {
                                    hourTime--;
                                    //当hourTime的数值为负数的时候，+24格式化成24小时显示
                                    if (hourTime < 0) {
                                        hourTime = 24 + hourTime;
                                    }
                                    //当小时数为一位的时候补零
                                    if (String(hourTime).length < 2) {
                                        hourTime = "0" + hourTime;
                                    }
                                    let hoursDate =
                                        yearAndDayTime + " " + hourTime;
                                    data.unshift({
                                        AddressL: data.address,
                                        time: hoursDate,
                                        uptime: ""
                                    });
                                }
                            }
                            data.forEach(item => {
                                item.time = item.time.substr(10, 12) + ":00";
                            });
                            xData = data.map(item => item.time);
                        } else {
                            let currayDate;
                            if (data.length > 2) {
                                currayDate = data[0].time;
                            } else {
                                currayDate = new Date().toISOString();
                            }
                            if (tabTime == "14days") {
                                let dataDateLength = data.length,
                                    //获取需要补全的天数
                                    complementdateLength = 14 - dataDateLength,
                                    //从那天需要补全的日期
                                    weekDate = new Date(currayDate),
                                    millisecondstime = weekDate.getTime(),
                                    //24小时的时间戳（毫秒数）
                                    dayNumberOfMilliseconds =
                                        60 * 60 * 1000 * 24;
                                //补全日期的逻辑
                                for (
                                    var lackOfDateNum = 0;
                                    lackOfDateNum < complementdateLength;
                                    lackOfDateNum++
                                ) {
                                    millisecondstime =
                                        millisecondstime -
                                        dayNumberOfMilliseconds;
                                    let complementdate = Tools.formatDateYearToDate(
                                        millisecondstime
                                    );

                                    data.unshift({
                                        time: complementdate,
                                        uptime: ""
                                    });
                                }
                            } else if (tabTime == "30days") {
                                let dataDateLength = data.length,
                                    complementdateLength = 30 - dataDateLength,
                                    monthDate = new Date(currayDate),
                                    millisecondstime = monthDate.getTime(),
                                    dayNumberOfMilliseconds =
                                        60 * 60 * 1000 * 24;
                                for (
                                    var lackOfDateNum = 0;
                                    lackOfDateNum < complementdateLength;
                                    lackOfDateNum++
                                ) {
                                    millisecondstime =
                                        millisecondstime -
                                        dayNumberOfMilliseconds;
                                    let complementdate = Tools.formatDateYearToDate(
                                        millisecondstime
                                    );

                                    data.unshift({
                                        time: complementdate,
                                        uptime: ""
                                    });
                                }
                            }

                            xData = data.map(
                                item =>
                                    `${String(item.time).substr(5, 2)}/${String(
                                        item.time
                                    ).substr(8, 2)}`
                            );
                        }
                        let seriesData = data.map(
                            item => item.uptime.toString().split(".")[0]
                        );
                        let noDatayAxisDefaultMax;

                        for (
                            let seriesDataIndex = 0;
                            seriesDataIndex < seriesData.length;
                            seriesDataIndex++
                        ) {
                            if (seriesData[seriesDataIndex] === "") {
                                noDatayAxisDefaultMax = "100";
                            }
                        }
                        this.informationUptimeLine = {
                            xData,
                            seriesData,
                            noDatayAxisDefaultMax
                        };
                    } else {
                        let xData,
                            currayDate,
                            data = [];
                        if (tabTime == "24hours") {
                            currayDate = new Date()
                                .toISOString()
                                .substr(0, 13)
                                .replace("T", " ");
                            let complementHourLength = 24;
                            let hourTime = currayDate.split(" ")[1];
                            let yearAndDayTime = currayDate.split(" ")[0];

                            for (let k = 0; k < complementHourLength; k++) {
                                hourTime--;
                                //当hourTime的数值为负数的时候，+24格式化成24小时显示
                                if (hourTime < 0) {
                                    hourTime = 24 + hourTime;
                                }
                                //当小时数为一位的时候补零
                                if (String(hourTime).length < 2) {
                                    hourTime = "0" + hourTime;
                                }
                                let hoursDate = yearAndDayTime + " " + hourTime;
                                data.unshift({
                                    AddressL: data.address,
                                    time: hoursDate,
                                    uptime: ""
                                });
                            }

                            data.forEach(item => {
                                item.time = item.time.substr(10, 12) + ":00";
                            });
                            xData = data.map(item => item.time);
                        } else if (tabTime == "14days") {
                            currayDate = new Date().toISOString();
                            //获取需要补全的天数
                            let complementdateLength = 14,
                                //从那天需要补全的日期
                                weekDate = new Date(currayDate),
                                millisecondstime = weekDate.getTime(),
                                //24小时的时间戳（毫秒数）
                                dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                            //补全日期的逻辑
                            for (
                                let lackOfDateNum = 0;
                                lackOfDateNum < complementdateLength;
                                lackOfDateNum++
                            ) {
                                millisecondstime =
                                    millisecondstime - dayNumberOfMilliseconds;
                                let complementdate = Tools.formatDateYearToDate(
                                    millisecondstime
                                );

                                data.unshift({
                                    time: complementdate,
                                    uptime: ""
                                });
                            }
                            xData = data.map(
                                item =>
                                    `${String(item.time).substr(5, 2)}/${String(
                                        item.time
                                    ).substr(8, 2)}`
                            );
                        } else if (tabTime == "30days") {
                            currayDate = new Date().toISOString();
                            let complementdateLength = 30,
                                monthDate = new Date(currayDate),
                                millisecondstime = monthDate.getTime(),
                                dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                            for (
                                let lackOfDateNum = 0;
                                lackOfDateNum < complementdateLength;
                                lackOfDateNum++
                            ) {
                                millisecondstime =
                                    millisecondstime - dayNumberOfMilliseconds;
                                let complementdate = Tools.formatDateYearToDate(
                                    millisecondstime
                                );
                                data.unshift({
                                    time: complementdate,
                                    uptime: ""
                                });
                            }
                            xData = data.map(
                                item =>
                                    `${String(item.time).substr(5, 2)}/${String(
                                        item.time
                                    ).substr(8, 2)}`
                            );
                        }
                        let noDatayAxisDefaultMax = "100";
                        let seriesData = data.map(
                            item => item.uptime.toString().split(".")[0]
                        );
                        this.informationUptimeLine = {
                            xData,
                            seriesData,
                            noDatayAxisDefaultMax
                        };
                    }
                })
                .catch(e => {});
        },
        openUrl(url) {
            if (url) {
                if (!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)) {
                    url = `http://${url}`;
                }
                window.open(url);
            }
        }
    }
};
</script>
<style lang="scss" scoped>
@import "../style/mixin.scss";
@import "../style/address_detail.scss";
</style>

<style lang="scss" scoped>
.validator_title {
    height: 0.25rem;
    padding-left: 0.2rem;
    display: flex;
    align-items: center;
    margin-top: 0.3rem;
    flex-direction: row !important;
    .title {
        font-size: 22px;
        color: #22252a;
    }
    .status_btn {
        margin-left: 0.1rem;
    }
}
.validator_table_container {
    display: flex;
    flex-direction: column;
    .blocks_list_table_container + div {
        align-self: flex-end;
        margin-top: 0.2rem;
    }
}
.transactions_detail_wrap {
    & > div:nth-last-of-type(1) {
        margin-bottom: 0.4rem;
    }
}
.status_btn {
    height: 0.2rem;
    padding: 0 0.14rem;
    font-size: 12px;
    line-height: 0.2rem;
    border-radius: 0.1rem;
    color: #ffffff;
    background-color: #3598db;
}
.validator_detail_information_wrap {
    margin-top: 0.2rem;
    border: 1px solid #d7d9e0;
    border-radius: 1px;
    padding: 0.2rem 0.2rem 0.06rem;
    display: flex;
    flex-wrap: wrap;
    .information_props_wrap {
        width: 50%;
        font-size: 14px;
        line-height: 20px;
        margin-bottom: 12px;
        flex: 0 0 50%;
        display: flex;
        justify-content: space-between;
        .information_props {
            color: #22252a;
        }
        .information_value {
            color: #a2a2ae;
            width: calc(100% - 1.7rem);
            margin-right: 0.2rem;
            word-break: break-all;
            word-wrap: break-word;
        }
        .skip_route {
            a,
            span {
                cursor: pointer;
                color: #3598db !important;
            }
        }
    }
}
.validator_information_content_title {
    height: 0.2rem;
    line-height: 0.2rem;
    color: #22252a;
    font-size: 18px;
    margin-top: 0.3rem;
    padding-left: 0.2rem;
}
.line_container_wrap {
    padding-bottom: 0 !important;
}
.delegations_container {
    width: 100%;
    display: flex;
    justify-content: space-between;
    .validator_information_content_title {
        margin-bottom: 0.2rem !important;
    }
    & > div {
        flex: 1;
        display: flex;
        flex-direction: column;
        .delegations_table_container + div {
            float: right;
            align-self: flex-end;
            margin-top: 0.2rem;
        }
    }
    & > div:nth-child(2n) {
        margin-left: 20px;
    }
}
.delegations_two_container {
    display: block;
    & > div {
        display: inline-block;
        width: calc(50% - 0.1rem);
        vertical-align: top;
    }
}
.line_tab_content_new {
    width: 180px;
    height: 0.3rem;
    display: flex;
    font-size: 12px;
    text-align: center;
    line-height: 28px;
    cursor: pointer;
    div {
        width: 60px;
        box-sizing: border-box;
        border: 1px solid rgba(215, 217, 224, 1);
        color: #22252a;
        &:nth-of-type(1) {
            border-right-width: 0;
            border-top-left-radius: 2px;
            border-bottom-left-radius: 2px;
        }

        &:nth-of-type(3) {
            border-left-width: 0;
            border-top-right-radius: 2px;
            border-bottom-right-radius: 2px;
        }
    }
    .active {
        background-color: #3598db;
        border-color: #3598db;
        color: #ffffff;
    }
}
.mobile_transactions_detail_wrap {
    .information_props_wrap {
        flex: 0 0 100%;
    }
    & > .second_table_container:nth-child(2n) {
        margin-left: 0px;
    }
    .delegations_container {
        flex-flow: column;
        & > div {
            width: 100%;
        }
        .delegations_table_container {
            width: 100%;
            overflow-x: auto;
        }
        .second_table_container {
            margin-left: 0;
        }
    }
}
.personal_computer_transactions_detail_wrap {
    .delegations_container {
        overflow: hidden;
        .delegations_table_container {
            overflow-x: auto;
        }
    }
}
@media screen and (min-width: 1280px) {
    .personal_computer_transactions_detail_wrap {
        .delegations_container {
            overflow: visible !important;
            .delegations_table_container {
                overflow: visible !important;
            }
        }
    }
}
</style>
