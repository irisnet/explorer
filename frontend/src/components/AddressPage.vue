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
            this.getAddressTxStatistics();
        } else {
            this.tabTxList(this.tabTxListIndex, this.txTabName, this.currentPage, this.pageSize);
            this.getAddressInformation(this.$route.params.param);
            this.getTransactionsList(1, 10, this.$route.params.type);
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
