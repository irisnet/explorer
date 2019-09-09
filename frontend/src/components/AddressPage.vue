<template>
    <div class="transactions_detail_wrap">
        <!-- <div class="transactions_title_wrap">
            <p :class="transactionsDetailWrap"
               style="margin-bottom:0;">
                <span class="transactions_detail_title">Address</span>
                <span class="transactions_detail_wrap_hash_var">
                    {{address}}
                    <i v-if="flValidator"
                       :style="{background:validatorStatusColor}">v</i>
                    <img v-show="flShowProfileLogo"
                         class="profile_logo_img"
                         src="../assets/profiler_logo.png">
                    <span v-show="flShowValidatorCandidate && flValidator"
                          class="candidate_validator">(This Validator is a Candidate)</span>
                    <span v-show="flShowValidatorJailed && flValidator"
                          class="jailed_validator">(This Validator is jailed!)</span>
                </span>
            </p>
        </div> -->

        <div :class="transactionsDetailWrap">
            <p class="transaction_information_content_title">Address Information</p>
            <div class="transactions_detail_information_wrap">
                <div class="information_props_wrap">
                    <span class="information_props">Address :</span>
                    <span class="information_value information_adress_value">
                        <span>
                            {{address?address:'--'}}
                        </span>
                        <i v-if="flValidator"
                        :style="{background:validatorStatusColor}">v</i>
                        <img v-show="flShowProfileLogo"
                            class="profile_logo_img"
                            src="../assets/profiler_logo.png">
                        <span v-show="flShowValidatorCandidate && flValidator"
                            class="candidate_validator">(This Validator is a Candidate)</span>
                        <span v-show="flShowValidatorJailed && flValidator"
                            class="jailed_validator">(This Validator is jailed!)</span>
                    </span>
                </div>
                <div class="information_props_wrap"
                     v-show="flValidator">
                    <span class="information_props">Self-Bonded :</span>
                    <span class="information_value">{{balanceValue?balanceValue:'--'}}</span>
                </div>
                <div class="information_props_wrap"
                     v-show="!flValidator">
                    <span class="information_props">Balance :</span>
                    <span class="information_value">{{balanceValue?balanceValue:'--'}}</span>
                </div>
                <div class="information_props_wrap"
                     v-show="flValidator">
                    <span class="information_props">Bonded Stake :</span>
                    <span class="information_value information_show_trim">{{depositsValue?depositsValue:'--'}}</span>
                </div>
                <div class="information_props_wrap"
                     v-show="!flValidator">
                    <span class="information_props">Delegated :</span>
                    <span class="information_value information_show_trim">{{depositsValue?depositsValue:'--'}}</span>
                </div>
                <div class="information_props_wrap"
                     v-show="!flValidator">
                    <span class="information_props">Withdraw To :</span>
                    <span class="information_value information_show_trim jump_link_style"
                          v-show="withdrawAddress">
                        <router-link :to="addressRoute(withdrawAddress)">{{withdrawAddress}}</router-link>
                    </span>
                    <span class="information_value information_show_trim"
                          v-show="!withdrawAddress">--</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Transactions :</span>
                    <span class="information_value">{{transactionsValue?transactionsValue:'--'}}</span>
                </div>
            </div>
        </div>
        <div :class="transactionsDetailWrap"
             class="address_profile"
             v-if="flValidator">
            <p class="transaction_information_content_title">Validator Profile</p>
            <div class="transactions_detail_information_wrap">
                <div class="information_props_wrap">
                    <span class="information_props">Name :</span>
                    <span class="information_value information_show_trim">{{nameValue?nameValue:'--'}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Pub Key :</span>
                    <span class="information_value">{{pubKeyValue?pubKeyValue:'--'}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Owner :</span>
                    <span class="information_value operator_value"
                          v-show="operatorValue">
                        <router-link :to="addressRoute(operatorValue)">{{operatorValue}}</router-link>
                    </span>
                    <span class="information_value"
                          v-show="!operatorValue">--</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Commission Rate :</span>
                    <span class="information_value">{{rateValue}}</span>
                </div>

                <div class="information_props_wrap">
                    <span class="information_props">Website :</span>
                    <span class="information_value"
                          :class="websiteValue && websiteValue !== '--' ? 'link_style' : ''">
                        <pre class="information_pre"
                             @click="openUrl(websiteValue)">{{websiteValue}}</pre>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Identity :</span>
                    <span class="information_value">
                        <pre class="information_pre">{{identity}}</pre></span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Details :</span>
                    <span class="information_value">
                        <pre class="information_pre">{{descriptionValue}}</pre></span>
                </div>
            </div>
        </div>
        <div :class="transactionsDetailWrap"
             class="current_tenure"
             v-show="flActiveValidator">
            <p class="transaction_information_content_title"
               style="border-bottom:1px solid #eee">Current Tenure</p>
            <div class="current_tenure_wrap">
                <div class="transactions_detail_information_wrap">
                    <div class="information_props_wrap">
                        <span class="information_props">Bond Height :</span>
                        <span class="information_value">{{bondHeightValue}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Voting Power :</span>
                        <span class="information_value">{{votingPowerValue}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Precommited Blocks :</span>
                        <span class="information_value">{{precommitedBlocksValue}}</span>
                    </div>
                </div>
                <div class="canvas_voting_power"
                     v-show="flActiveValidator">
                    <div class="progress_wrap">
                        <span>Uptime(in last 100)</span>
                        <div class="progress_wrap_background">
                            <div class="progress_value"
                                 :style="`width:${firstPercent}`">{{firstPercent}}</div>
                        </div>
                    </div>
                    <div class="progress_wrap">
                    </div>
                </div>
            </div>
        </div>
        <div class="line_container_wrap"
             v-show="flActiveValidator">
            <div class="line_container"
                 :class="transactionsDetailWrap">
                <p class="line_history_title">History</p>
                <div class="line_content">
                    <div class="line_echarts_content">
                        <div class="line_left_container"
                             style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
                            <echarts-validators-line :informationValidatorsLine="informationValidatorsLine"></echarts-validators-line>
                        </div>
                        <div class="line_tab_content">
                            <div v-for="(item,index) in tabVotingPower"
                                 :key="index"
                                 @click="getValidatorHistory(item.title,index)"
                                 :class="item.active ? 'border-none' : 'border-block' ">{{item.title}}</div>
                        </div>
                    </div>
                    <div class="line_echarts_content "
                         :class="transactionsDetailWrap === 'personal_computer_transactions_detail_wrap' ?
           'content_right' : 'model_content_right' ">
                        <div class="line_right_container"
                             style="overflow-x: auto;-webkit-overflow-scrolling:touch;">
                            <echarts-validators-uptime-line :informationUptimeLine="informationUptimeLine"></echarts-validators-uptime-line>
                        </div>
                        <div class="line_tab_content">
                            <div v-for="(item,index) in tabUptime"
                                 :key="index"
                                 @click="getValidatorUptimeHistory(item.title,index)"
                                 :class="item.active ? 'border-none' : 'border-block' ">{{item.title}}</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="list_tab_wrap"
             :class="transactionsDetailWrap">
            <div class="list_tab_content">
                <ul class="list_tab_container">
                    <li class="list_tab_item"
                        :class="item.active ? 'activeStyle' : 'unactiveStyle'"
                        v-for="(item,index) in txTab"
                        :key="index"
                        @click="tabTxList(index,item.txTabName,1,20)">{{item.txTabName}} ({{item.txTotal}})
                    </li>
                </ul>
            </div>
        </div>
        <div :class="transactionsDetailWrap">
            <div class="pagination total_num">
                <b-pagination size="md"
                              :total-rows="count"
                              v-model="currentPage"
                              :per-page="pageSize">
                </b-pagination>
            </div>
            <div class="blocks_list_table_container">
                <spin-component :showLoading="showLoading" />
                <div class="address_tx_list_table">
                    <blocks-list-table :items="items"
                                       :type="'addressTxList'"
                                       :showNoData="showNoData"
                                       class="address_tx_list_table"
                                       style="border-left: 1px solid #dee2e6; border-right: 1px solid #dee2e6;"></blocks-list-table>
                    <div v-show="showNoData"
                         class="no_data_show">
                        <img src="../assets/no_data.svg" alt="">
                    </div>
                </div>
            </div>
            <div class="pagination"
                 style='margin:0.2rem 0 0.4rem;'>
                <b-pagination size="md"
                              :total-rows="count"
                              v-model="currentPage"
                              :per-page="pageSize">
                </b-pagination>
            </div>
        </div>
    </div>
</template>

<script>
import Tools from '../util/Tools';
import BlocksListTable from './table/BlocksListTable';
import EchartsLine from "./EchartsLine";
import EchartsValidatorsLine from "./EchartsValidatorsLine";
import EchartsValidatorsUptimeLine from "./EchartsValidatorsUpTimeLine";
import SpinComponent from './commonComponents/SpinComponent';
import Service from "../service";
import Constants from "../constant/Constant"
export default {
    watch: {
        currentPage (currentPage) {
            this.currentPage = currentPage;
            new Promise((resolve) => {
                this.tabTxList(this.currentTabIndex, this.currentTxTabName, currentPage, this.pageSize);
                resolve();
            }).then(() => {
                document.getElementById('router_wrap').scrollTop = 0;
            })

        },
        $route () {
            Tools.scrollToTop();
            this.type = this.$route.params.type;
            this.tabTxList(this.tabTxListIndex, this.txTabName, this.currentPage, this.pageSize);
            this.getAddressTxStatistics();
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList(1, 10, this.$route.params.type);
            this.getProfileInformation();
            this.getCurrentTenureInformation();
            this.getValidatorHistory('14days');
            this.getValidatorUptimeHistory('24hours');
        },
        "$store.state.isMobile"(newVal) {
            this.isMobileFunc(newVal);
        }
    },
    data () {

        return {
            rateValue: '',
            devicesWidth: window.innerWidth,
            transactionsDetailWrap: 'personal_computer_transactions_detail',
            activeBtn: 0,
            currentPage: 1,
            pageSize: 20,
            addressTxList: "",
            firstPercent: '',
            secondPercent: '%',
            address: this.$route.params.param,
            balanceValue: '',
            depositsValue: '',
            transactionsValue: '',
            nameValue: '',
            pubKeyValue: '',
            websiteValue: '',
            descriptionValue: '',
            commissionRateValue: '',
            announcementValue: '',
            bondHeightValue: '',
            votingPowerValue: '',
            uptimeValue: '',
            precommitedBlocksValue: '',
            returnsValue: '',
            operatorValue: '',
            items: [],
            itemsPre: [],
            type: this.$route.params.type,
            totalBlocks: 0,
            totalFee: 0,
            TransactionsShowNoData: false,
            PrecommitBlocksshowNoData: false,
            transactionsCount: 0,
            flValidator: false,
            showNoData: false,
            showLoading: false,
            informationValidatorsLine: {},
            informationUptimeLine: {},
            transactionsTitle: "",
            count: 0,
            txTabName: "Transfers",
            tabTxListIndex: 0,
            currentTabIndex: "",
            currentTxTabName: "",
            identity: "",
            withdrawAddress: "",
            flShowValidatorJailed: false,
            flShowValidatorCandidate: false,
            flActiveValidator: true,
            flShowProfileLogo: false,
            validatorStatusColor: "var(--bgColor)",
            tabVotingPower: [
                {
                    "title": "14days",
                    "active": true
                },
                {
                    "title": "30days",
                    "active": false
                },
                {
                    "title": "60days",
                    "active": false
                }
            ],
            tabUptime: [
                {
                    "title": "24hours",
                    "active": true
                },
                {
                    "title": "14days",
                    "active": false
                },
                {
                    "title": "30days",
                    "active": false
                }
            ],
            txTab: [
                {
                    "txTabName": "Transfers",
                    "active": true,
                    txTotal: "",
                },
                {
                    "txTabName": "Stakes",
                    "active": false,
                    txTotal: "",

                },
                {
                    "txTabName": "Declarations",
                    "active": false,
                    txTotal: "",
                },
                {
                    "txTabName": "Governance",
                    "active": false,
                    txTotal: ""
                }
            ]
        }
    },
    components: {
        EchartsValidatorsUptimeLine,
        EchartsValidatorsLine,
        EchartsLine,
        BlocksListTable,
        SpinComponent,
    },
    beforeMount () {
        this.isMobileFunc(this.$store.state.isMobile);
    },
    mounted () {
        Tools.scrollToTop();
        if (sessionStorage.getItem('currentEnv') !== Constants.ENVCONFIG.MAINNET) {
            this.$Crypto.getCrypto(Constants.CHAINNAME, Constants.ENVCONFIG.TESTNET);
        }
        if (this.$route.params.param.substring(0, 3) === this.$Crypto.config.iris.bech32.valAddr) {
            this.tabTxList(this.tabTxListIndex, this.txTabName, this.currentPage, this.pageSize);
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList(1, 10, this.$route.params.type);
            this.getProfileInformation();
            this.getCurrentTenureInformation();
            this.getValidatorHistory('14days');
            this.getValidatorUptimeHistory('24hours');
            this.getAddressTxStatistics();
        } else {
            this.tabTxList(this.tabTxListIndex, this.txTabName, this.currentPage, this.pageSize);
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList(1, 10, this.$route.params.type);
            this.getProfileInformation();
            this.getAddressTxStatistics();
        }
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
        getAddressTxStatistics () {
            Service.commonInterface({                address: {
                    address: this.$route.params.param
                }            }, (data) => {
                try {
                    if (data) {
                        this.txTab[0].txTotal = data.trans_cnt;
                        this.txTab[1].txTotal = data.stake_cnt;
                        this.txTab[2].txTotal = data.declaration_cnt;
                        this.txTab[3].txTotal = data.gov_cnt;
                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },
        tabTxList (index, txTabName, currentPage, pageSize) {
            txTabName = Tools.firstWordLowerCase(txTabName)
            this.currentPage = currentPage;
            this.currentTabIndex = index;
            this.currentTxTabName = txTabName;
            this.showLoading = true;
            for (let txTabIndex = 0; txTabIndex < this.txTab.length; txTabIndex++) {
                this.txTab[txTabIndex].active = false;
                this.txTab[index].active = true;
            }
            let that = this, params;
            if(txTabName === 'stakes'){
	            txTabName = 'delegations'
            }else if(txTabName === 'declarations'){
	            txTabName = 'validations'
            }
            if (txTabName === 'transfers') {
                params = { addressTxTrans: { pageNumber: currentPage, pageSize: pageSize, address: this.$route.params.param } }
            } else if (txTabName === 'delegations') {
                params = { addressTxStake: { pageNumber: currentPage, pageSize: pageSize, address: this.$route.params.param } }
            } else if (txTabName === 'validations') {
                params = { addressTxDeclaration: { pageNumber: currentPage, pageSize: pageSize, address: this.$route.params.param } }
            } else if (txTabName === 'governance') {
                params = { addressTxGov: { pageNumber: currentPage, pageSize: pageSize, address: this.$route.params.param } }
            }
            Service.commonInterface(params, (txList) => {
                try {
                    this.showLoading = false;
                    this.showNoData = false;
                    this.count = txList.Count;
                    if (txList.Data) {
                        this.items = Tools.formatTxList(txList.Data, txTabName)
                    } else {
                        this.items = Tools.formatTxList(null, txTabName);
                        this.showNoData = true;
                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },
        getAddressInformation (address) {
            this.address = address;
            Service.commonInterface({ addressAccount: { address: this.$route.params.param } }, (result) => {
                try {
                    let Amount = '--';
                    if (result) {
                        if (result.amount && result.amount instanceof Array) {
	                        result.amount.forEach( (item,index) => {
                                if(item.denom === Constants.Denom.IRISATTO){
                                    result.amount[index].amount = Tools.formatNumber(result.amount[index].amount);
	                                Amount = `${Tools.formatPriceToFixed(result.amount[index].amount)} ${Tools.formatDenom(result.amount[index].denom).toUpperCase()}`
                                }
                            });
                            // Amount = result.amount.map(listItem => `${Tools.formatPriceToFixed(listItem.amount)} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',')
                        } else if (result.amount && result.amount instanceof Array && result.amount[0].denom === Constants.Denom.IRIS) {
                            Amount = result.amount.map(listItem => `${Tools.formatPriceToFixed(listItem.amount)} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',')
                        } else if (result.amount && Object.keys(result.amount).includes('amount') && Object.keys(result.amount).includes('denom') && result.amount.denom === Constants.Denom.IRISATTO) {
                            Amount = `${Tools.formatPriceToFixed(Tools.formatNumber(result.amount))} ${result.denom.toUpperCase()}`
                        } else if (result.amount && Object.keys(result.amount).includes('amount') && Object.keys(result.amount).includes('denom') && result.amount.denom === Constants.Denom.IRIS) {
                            Amount = `${Tools.formatPriceToFixed(result.amount)} ${result.denom.toUpperCase()}`
                        } else {
                            Amount = ''
                        }
                    }
                    this.flShowProfileLogo = result && result.isProfiler;
                    if (result && result.deposits) {
                        let deposits = Tools.formatToken(result.deposits);
	                    if(String(Tools.FormatScientificNotationToNumber(deposits.amount)).length > 18){
		                    deposits.amount =`${String(Tools.FormatScientificNotationToNumber(deposits.amount)).split('.')[0]}.${String(Tools.FormatScientificNotationToNumber(deposits.amount)).split('.')[1].substring(0,18)}`
	                    }
                        this.depositsValue = result.deposits.amount && result.deposits.amount !== 0 && result.deposits.amount !== '' ? `${Tools.formatPriceToFixed(deposits.amount)} ${deposits.denom}` : '';
                    }
                    this.withdrawAddress = (result && result.withdrawAddress) ? result.withdrawAddress : '--';
                    this.balanceValue = Amount;

                } catch (e) {
                    console.error(e)
                }
            })
        },
        getProfileInformation () {
            Service.commonInterface({                addressStakeCandidate: {
                    address: this.$route.params.param
                }            }, (result) => {
                try {
                    if (result) {
                        let validator = result.validator;
                        this.flValidator = true;
                        if (validator.jailed === true) {
                            this.flActiveValidator = false;
                            this.flShowValidatorJailed = true;
                            this.validatorStatusColor = "#f00";
                        } else {
                            if (validator.status === 'Unbonded' || validator.status === 'Unbonding') {
                                this.flShowValidatorCandidate = true;
                                this.validatorStatusColor = "#45B035";
                                this.flActiveValidator = false;
                            } else if (validator.status === "Bonded") {
                                this.bondHeightValue = validator.bond_height;
                                this.votingPowerValue = validator.voting_power;
                            }
                        }
                        this.rateValue = validator.rate ? `${Tools.formatRate(validator.rate.toString())}%` : '--';
                        this.identity = validator.description && validator.description.identity ? validator.description.identity : "--";
                        this.nameValue = validator.description && validator.description.moniker ? validator.description.moniker : '--';
                        this.pubKeyValue = validator.pub_key ? validator.pub_key : "--";
                        this.websiteValue = validator.description && validator.description.website ? validator.description.website : '--';
                        this.descriptionValue = validator.description && validator.description.details ? validator.description.details : "--";
                        this.commissionRateValue = '';
                        this.announcementValue = '';
                        this.operatorValue = validator.owner;
                    } else {
                        this.flValidator = false;
                        this.flActiveValidator = false;
                    }
                } catch (e) {
                    this.flActiveValidator = false;
                    console.error(e)
                }
            })
        },
        getCurrentTenureInformation () {
            Service.commonInterface({ addressStakeCandidateStatus: { address: this.$route.params.param } }, (data) => {
                try {
                    if (data) {
                        this.precommitedBlocksValue = data.precommit_cnt ? data.precommit_cnt : '--';
                        this.returnsValue = '';
                        this.firstPercent = data.up_time ? `${data.up_time}%` : "--";
                    }
                } catch (e) {
                    console.error(err)
                }
            })
        },
        getTransactionsList () {
            Service.commonInterface({                addressTxByAddress: {
                    address: this.$route.params.param,
                    pageNumber: 1,
                    pageSize: 30,
                }            }, (data) => {
                try {
                    if (data) {
                        this.transactionsCount = data.Count;
                        this.transactionsValue = data.Count;
                        this.TransactionsShowNoData = true;
                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },
        getValidatorHistory (tabTime, index) {
            if (index !== undefined) {
                for (var i = 0; i < this.tabVotingPower.length; i++) {
                    this.tabVotingPower[i].active = false;
                    this.tabVotingPower[index].active = true
                }
            }
            let url, params;
            if (tabTime == "14days") {
                params = { addressStakeCandidateWeek: { address: this.$route.params.param } };
            } else if (tabTime == "30days") {
                params = { addressStakeCandidateMonth: { address: this.$route.params.param } };
            } else if (tabTime == "60days") {
                params = { addressStakeCandidateMonths: { address: this.$route.params.param } };
            }
            Service.commonInterface(params, (validatorVotingPowerList) => {
                try {
                    if (validatorVotingPowerList) {
                        let seriesData = [], noDatayAxisDefaultMaxByValidators;
                        let maxPowerValue = 0;
                        validatorVotingPowerList.forEach(item => {
                            if (item.power > maxPowerValue) {
                                maxPowerValue = item.power
                            }
                            if (item.power == 0) {
                                item.power = ""
                            }
                            let obj = [];
                            obj[0] = Tools.conversionTimeToUTCByValidatorsLine(item.time);
                            obj[1] = item.power;
                            seriesData.push(obj);
                        });
                        if (maxPowerValue < 100) {
                            noDatayAxisDefaultMaxByValidators = "100"
                        }
                        this.informationValidatorsLine = { seriesData, noDatayAxisDefaultMaxByValidators };

                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },
        getValidatorUptimeHistory (tabTime, index) {
            if (index != undefined) {
                for (var i = 0; i < this.tabUptime.length; i++) {
                    this.tabUptime[i].active = false;
                    this.tabUptime[index].active = true
                }
            }

            let url, params;
            if (tabTime == "24hours") {
                params = { addressStakeCandidateUptimeHour: { address: this.$route.params.param } };
            } else if (tabTime == "14days") {
                params = { addressStakeCandidateUptimeWeek: { address: this.$route.params.param } };
            } else if (tabTime == "30days") {
                params = { addressStakeCandidateUptimeMonth: { address: this.$route.params.param } };
            }
            Service.commonInterface(params, (data) => {
                try {
                    if (data) {
                        data.forEach(item => {
                            let notValidatorTag = -1;
                            if (item.uptime === notValidatorTag) {
                                item.uptime = ""
                            }
                        });
                        let xData, currayDate;
                        if (tabTime == "24hours") {
                            data.forEach(item => {
                                //兼容火狐
                                let hourseconds = item.time.replace("-", "/").replace("-", "/") + ":00:00";

                                let changeMilliseconds = new Date(hourseconds).getTime();
                                //展示需加一小时
                                changeMilliseconds = changeMilliseconds + 60 * 60 * 1000;
                                item.time = Tools.formatDateYearAndMinutesAndSeconds(changeMilliseconds).split(":")[0];
                            });
                            if (data.length !== 0) {
                                currayDate = data[0].time;
                            } else {
                                currayDate = new Date().toISOString().substr(0, 13).replace("T", " ");
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
                                    let hoursDate = yearAndDayTime + " " + hourTime;
                                    data.unshift({ AddressL: data.address, time: hoursDate, uptime: "" })
                                }
                            }
                            data.forEach((item) => {
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
                                    dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                                //补全日期的逻辑
                                for (var lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                                    millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                                    let complementdate = Tools.formatDateYearToDate(millisecondstime);

                                    data.unshift({ time: complementdate, uptime: "" });
                                }

                            } else if (tabTime == "30days") {

                                let dataDateLength = data.length,
                                    complementdateLength = 30 - dataDateLength,
                                    monthDate = new Date(currayDate),
                                    millisecondstime = monthDate.getTime(),
                                    dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                                for (var lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                                    millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                                    let complementdate = Tools.formatDateYearToDate(millisecondstime);

                                    data.unshift({ time: complementdate, uptime: "" });
                                }
                            }

                            xData = data.map(item => `${String(item.time).substr(5, 2)}/${String(item.time).substr(8, 2)}`);
                        }
                        let seriesData = data.map(item => item.uptime.toString().split(".")[0]);
                        let noDatayAxisDefaultMax;

                        for (let seriesDataIndex = 0; seriesDataIndex < seriesData.length; seriesDataIndex++) {
                            if (seriesData[seriesDataIndex] === "") {
                                noDatayAxisDefaultMax = "100"
                            }
                        }
                        this.informationUptimeLine = { xData, seriesData, noDatayAxisDefaultMax };
                    } else {
                        let xData, currayDate, data = [];
                        if (tabTime == "24hours") {
                            currayDate = new Date().toISOString().substr(0, 13).replace("T", " ");
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
                                data.unshift({ AddressL: data.address, time: hoursDate, uptime: "" })
                            }

                            data.forEach((item) => {
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
                            for (let lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                                millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                                let complementdate = Tools.formatDateYearToDate(millisecondstime);

                                data.unshift({ time: complementdate, uptime: "" });
                            }
                            xData = data.map(item => `${String(item.time).substr(5, 2)}/${String(item.time).substr(8, 2)}`);
                        } else if (tabTime == "30days") {
                            currayDate = new Date().toISOString();
                            let complementdateLength = 30,
                                monthDate = new Date(currayDate),
                                millisecondstime = monthDate.getTime(),
                                dayNumberOfMilliseconds = 60 * 60 * 1000 * 24;
                            for (let lackOfDateNum = 0; lackOfDateNum < complementdateLength; lackOfDateNum++) {
                                millisecondstime = millisecondstime - dayNumberOfMilliseconds;
                                let complementdate = Tools.formatDateYearToDate(millisecondstime);
                                data.unshift({ time: complementdate, uptime: "" });
                            }
                            xData = data.map(item => `${String(item.time).substr(5, 2)}/${String(item.time).substr(8, 2)}`);
                        }
                        let noDatayAxisDefaultMax = "100";
                        let seriesData = data.map(item => item.uptime.toString().split(".")[0]);
                        this.informationUptimeLine = { xData, seriesData, noDatayAxisDefaultMax };
                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },

        openUrl (url) {
            if (url && url !== '--') {
                if (!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)) {
                    url = `http://${url}`
                }
                window.open(url)
            }
        }
    }
}
</script>
<style lang="scss" scoped>
@import "../style/mixin.scss";
@import "../style/address_detail.scss";
</style>
