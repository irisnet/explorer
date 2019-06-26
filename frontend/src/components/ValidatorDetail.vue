<template>
    <div class="transactions_detail_wrap">
        <div :class="[transactionsDetailWrap, 'validator_title']">
            <span class="title">HashQuark</span>
            <div class="status_btn">Active</div>
        </div>
        <div :class="transactionsDetailWrap">
            <div class="validator_detail_information_wrap">
                <div class="information_props_wrap"
                     v-for="v in Object.entries(validatorInfo)"
                     :key="v[0]">
                    <span class="information_props">{{v[0]}}: </span>
                    <span class="information_value">{{v[1]}}</span>
                </div>
            </div>
        </div>
        <div :class="transactionsDetailWrap"
             class="address_profile">
            <p class="validator_information_content_title">Validator Profile</p>
            <div class="validator_detail_information_wrap">
                <div class="information_props_wrap"
                     v-for="v in Object.entries(validatorProfile)"
                     :key="v[0]">
                    <span class="information_props">{{v[0]}}: </span>
                    <span class="information_value">{{v[1]}}</span>
                </div>
            </div>
        </div>
        <div class="line_container_wrap"
             v-show="flActiveValidator">
            <div class="line_container"
                 :class="transactionsDetailWrap">
                <p class="validator_information_content_title">History</p>
                <div class="line_content">
                    <div class="line_echarts_content">
                        <div class="line_left_container"
                             style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
                            <echarts-validators-line :informationValidatorsLine="informationValidatorsLine">
                                <div slot="tabs"
                                     class="line_tab_content_new">
                                    <div v-for="(item,index) in tabVotingPower"
                                         :key="index"
                                         @click="getValidatorHistory(item.title,index)"
                                         :class="item.active ? 'active' : ''">{{item.title}}</div>
                                </div>
                            </echarts-validators-line>
                        </div>
                    </div>
                    <div class="line_echarts_content"
                         :class="transactionsDetailWrap === 'personal_computer_transactions_detail_wrap' ?
           'content_right' : 'model_content_right' ">
                        <div class="line_right_container"
                             style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
                            <echarts-validators-uptime-line :informationUptimeLine="informationUptimeLine">
                                <div slot="tabs"
                                     class="line_tab_content_new">
                                    <div v-for="(item,index) in tabUptime"
                                         :key="index"
                                         @click="getValidatorUptimeHistory(item.title,index)"
                                         :class="item.active ? 'active' : ''">{{item.title}}</div>
                                </div>
                            </echarts-validators-uptime-line>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div :class="transactionsDetailWrap"
             class="delegations_container">
            <div>
                <p class="validator_information_content_title">Delegations</p>
                <m-validatorDetail-table :items="txTableList.delegations.items"
                                         listName="delegations"></m-validatorDetail-table>
                <m-pagination :page-size="pageSize"
                              :total="txTableList.delegations.total"
                              :page-change="pageChange"></m-pagination>

            </div>
            <div>
                <p class="validator_information_content_title">Unbonding Delegations</p>
                <m-validatorDetail-table :items="txTableList.unbondingDelegations.items"
                                         listName="unbondingDelegations"></m-validatorDetail-table>
                <m-pagination :page-size="pageSize"
                              :total="txTableList.unbondingDelegations.total"
                              :page-change="pageChange"></m-pagination>
            </div>
        </div>
        <div :class="transactionsDetailWrap"
             class="validator_table_container">
            <p class="validator_information_content_title">Redelegations</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validatorDetail-table :items="txTableList.redelegations.items"
                                             listName="redelegations"></m-validatorDetail-table>
                </div>
            </div>
            <m-pagination :page-size="pageSize"
                          :total="txTableList.delegations.total"
                          :page-change="pageChange"></m-pagination>
        </div>
        <div :class="transactionsDetailWrap"
             class="delegations_container">
            <div>
                <p class="validator_information_content_title">Deposited Proposals</p>
                <m-validatorDetail-table :items="txTableList.depositedProposals.items"
                                         listName="depositedProposals"></m-validatorDetail-table>
                <m-pagination :page-size="pageSize"
                              :total="txTableList.delegations.total"
                              :page-change="pageChange"></m-pagination>
            </div>
            <div>
                <p class="validator_information_content_title">Voted Proposals</p>
                <m-validatorDetail-table :items="txTableList.votedProposals.items"
                                         listName="votedProposals"></m-validatorDetail-table>
                <m-pagination :page-size="pageSize"
                              :total="txTableList.delegations.total"
                              :page-change="pageChange"></m-pagination>
            </div>
        </div>
        <div :class="transactionsDetailWrap"
             class="validator_table_container">
            <p class="validator_information_content_title">Transfers</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validatorDetail-table :items="txTableList.transfers.items"
                                             listName="transfer"></m-validatorDetail-table>
                </div>
            </div>
            <m-pagination :page-size="pageSize"
                          :total="txTableList.delegations.total"
                          :page-change="pageChange"></m-pagination>
        </div>
        <div :class="transactionsDetailWrap"
             class="validator_table_container">
            <p class="validator_information_content_title">Stakes</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validatorDetail-table :items="txTableList.stakes.items"
                                             listName="stakes"></m-validatorDetail-table>
                </div>
            </div>
            <m-pagination :page-size="pageSize"
                          :total="txTableList.delegations.total"
                          :page-change="pageChange"></m-pagination>
        </div>
        <div :class="transactionsDetailWrap"
             class="validator_table_container">
            <p class="validator_information_content_title">Declarations</p>
            <div class="blocks_list_table_container">
                <div class="address_tx_list_table">
                    <m-validatorDetail-table :items="txTableList.declarations.items"
                                             listName="declarations"></m-validatorDetail-table>
                </div>
            </div>
            <m-pagination :page-size="pageSize"
                          :total="txTableList.delegations.total"
                          :page-change="pageChange"></m-pagination>
        </div>
        <div :class="transactionsDetailWrap"
             class="validator_table_container"
             style="margin-bottom: 0.4rem;">
            <p class="validator_information_content_title">Governance</p>
            <div class="blocks_list_table_container">
                <!-- <spin-component :showLoading="showLoading" /> -->
                <div class="address_tx_list_table">
                    <m-validatorDetail-table :items="txTableList.governance.items"
                                             listName="gov"></m-validatorDetail-table>
                </div>
            </div>
            <m-pagination :page-size="pageSize"
                          :total="txTableList.delegations.total"
                          :page-change="pageChange"></m-pagination>
        </div>
    </div>
</template>

<script>
import Tools from "../util/Tools";
import EchartsLine from "./EchartsLine";
import EchartsValidatorsLine from "./EchartsValidatorsLine";
import EchartsValidatorsUptimeLine from "./EchartsValidatorsUpTimeLine";
import SpinComponent from "./commonComponents/SpinComponent";
import Service from "../util/axios";
import Constants from "../constant/Constant";
import MValidatorDetailTable from './table/MValidatorDetailTable';
import MPagination from "./commonComponents/MPagination";
export default {
    watch: {
        currentPage (currentPage) {
            this.currentPage = currentPage;
            new Promise(resolve => {
                // this.tabTxList(
                //     this.currentTxTabName,
                //     currentPage,
                //     this.pageSize
                // );
                resolve();
            }).then(() => {
                document.getElementById("router_wrap").scrollTop = 0;
            });
        },
        $route () {
            Tools.scrollToTop();
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList();
            this.getProfileInformation();
            this.getCurrentTenureInformation();
            this.getValidatorHistory("14days");
            this.getValidatorUptimeHistory("24hours");
        },
        '$store.state.isMobile' (newVal) {
            this.isMobileFunc(newVal);
        }
    },
    data () {
        return {
            rateValue: "",
            devicesWidth: window.innerWidth,
            transactionsDetailWrap: "personal_computer_transactions_detail",
            activeBtn: 0,
            currentPage: 1,
            pageSize: 5,
            addressTxList: "",
            firstPercent: "",
            secondPercent: "%",
            address: this.$route.params.param,
            balanceValue: "",
            depositsValue: "",
            transactionsValue: "",
            nameValue: "",
            pubKeyValue: "",
            websiteValue: "",
            descriptionValue: "",
            commissionRateValue: "",
            announcementValue: "",
            bondHeightValue: "",
            votingPowerValue: "",
            uptimeValue: "",
            precommitedBlocksValue: "",
            returnsValue: "",
            operatorValue: "",
            itemsPre: [],
            totalBlocks: 0,
            totalFee: 0,
            TransactionsShowNoData: false,
            PrecommitBlocksshowNoData: false,
            transactionsCount: 0,
            flValidator: false,
            showLoading: false,
            informationValidatorsLine: {},
            informationUptimeLine: {},
            transactionsTitle: "",
            txTabName: "Transfers",
            currentTxTabName: "",
            identity: "",
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
                'Voting Power': '',
                'Bond Height': '',
                'Bonded Tokens': '',
                'Unbonding Height': '',
                'Jailed Until': '',
                'Delegator Bonded': '',
                'Missed Blocks': '',
                'Delegators': '',
                'Commission Rate': '',
                'Delegator Shares': '',
                'Max Rate': '',
                'Rewards': '',
                'Max Change Rate': ''
            },
            validatorProfile: {
                'Validator Profile': '',
                'Website': '',
                'Owner Address': '',
                'Identity': '',
                'Withdraw Address': '',
                'Details': '',
                'Consensus Pubkey': ''
            },
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
            }
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
    beforeMount () {
        this.isMobileFunc(this.$store.state.isMobile);
    },
    mounted () {
        Tools.scrollToTop();
        if (
            sessionStorage.getItem("currentEnv") !== Constants.ENVCONFIG.MAINNET
        ) {
            this.$Crypto.getCrypto(
                Constants.CHAINNAME,
                Constants.ENVCONFIG.TESTNET
            );
        }
        this.queryTxList();
        if (
            this.$route.params.param.substring(0, 3) ===
            this.$Crypto.config.iris.bech32.valAddr
        ) {
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList();
            this.getProfileInformation();
            this.getCurrentTenureInformation();
            this.getValidatorHistory("14days");
            this.getValidatorUptimeHistory("24hours");
        } else {
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList();
            this.getProfileInformation();
        }
    },
    methods: {
        isMobileFunc (isMobile) {
            if (isMobile) {
                this.transactionsDetailWrap = "mobile_transactions_detail_wrap";
            } else {
                this.transactionsDetailWrap =
                    "personal_computer_transactions_detail_wrap";
            }
        },
        pageChange () {

        },
        getDelegations () {

        },
        getUnbondingDelegations () {

        },
        queryTxList () {
            let arr = Object.entries(this.txTableList).map(v => {
                this.tabTxList(
                    v[0],
                    v[1].currentPage,
                    this.pageSize
                );
            });
        },
        tabTxList (txTabName, currentPage, pageSize) {
            txTabName = Tools.firstWordLowerCase(txTabName);
            this.currentPage = currentPage;
            this.currentTxTabName = txTabName;
            this.showLoading = true;
            let url;
            let that = this;
            if (txTabName === "transfers") {
                url = `/api/tx/trans/${currentPage}/${pageSize}?address=${
                    this.$route.params.param
                    }`;
            } else if (txTabName === "stakes") {
                url = `/api/tx/stake/${currentPage}/${pageSize}?address=${
                    this.$route.params.param
                    }`;
            } else if (txTabName === "declarations") {
                url = `/api/tx/declaration/${currentPage}/${pageSize}?address=${
                    this.$route.params.param
                    }`;
            } else if (txTabName === "governance") {
                url = `/api/tx/gov/${currentPage}/${pageSize}?address=${
                    this.$route.params.param
                    }`;
            }
            Service.http(url).then(txList => {
                this.showLoading = false;
                if (txList.Data) {
                    this.txTableList[txTabName].items = Tools.formatTxList(txList.Data, txTabName);
                } else {
                    this.txTableList[txTabName].items = [];
                }
            });
        },
        getAddressInformation (address) {
            this.address = address;
            let url = `/api/account/${this.$route.params.param}`;
            Service.http(url)
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
                            : "--";
                    this.balanceValue = Amount;
                })
                .catch(e => {
                    console.error(e);
                });
        },
        getProfileInformation () {
            let url = `/api/stake/candidate/${this.$route.params.param}`;
            Service.http(url)
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
                            : "--";
                        this.identity =
                            validator.description &&
                                validator.description.identity
                                ? validator.description.identity
                                : "--";
                        this.nameValue =
                            validator.description &&
                                validator.description.moniker
                                ? validator.description.moniker
                                : "--";
                        this.pubKeyValue = validator.pub_key
                            ? validator.pub_key
                            : "--";
                        this.websiteValue =
                            validator.description &&
                                validator.description.website
                                ? validator.description.website
                                : "--";
                        this.descriptionValue =
                            validator.description &&
                                validator.description.details
                                ? validator.description.details
                                : "--";
                        this.commissionRateValue = "";
                        this.announcementValue = "";
                        this.operatorValue = validator.owner;
                    } else {
                        this.flValidator = false;
                        this.flActiveValidator = false;
                    }
                })
                .catch(err => {
                    this.flActiveValidator = false;
                    console.error(err);
                });
        },
        getCurrentTenureInformation () {
            let url = `/api/stake/candidate/${this.$route.params.param}/status`;
            Service.http(url)
                .then(data => {
                    if (data) {
                        this.precommitedBlocksValue = data.precommit_cnt
                            ? data.precommit_cnt
                            : "--";
                        this.returnsValue = "";
                        this.firstPercent = data.up_time
                            ? `${data.up_time}%`
                            : "--";
                    }
                })
                .catch(err => {
                    console.error(err);
                });
        },
        getTransactionsList () {
            let url = `/api/txsByAddress/${this.$route.params.param}/1/30`;
            Service.http(url)
                .then(data => {
                    if (data) {
                        this.transactionsCount = data.Count;
                        this.transactionsValue = data.Count;
                        this.TransactionsShowNoData = true;
                    }
                })
                .catch(e => {
                    console.error(e);
                });
        },
        getValidatorHistory (tabTime, index) {
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
            Service.http(url).then(validatorVotingPowerList => {
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
        getValidatorUptimeHistory (tabTime, index) {
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
            Service.http(url)
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
                .catch(e => {
                    console.error(e);
                });
        },

        openUrl (url) {
            if (url && url !== "--") {
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
        flex: 1 0 50%;
        .information_props {
            color: #22252a;
        }
        .information_value {
            color: #a2a2ae;
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
        .validator_detail_table + div {
            align-self: flex-end;
            margin-top: 0.2rem;
        }
    }
    & > div:nth-last-child(1) {
        margin-left: 20px;
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
</style>
