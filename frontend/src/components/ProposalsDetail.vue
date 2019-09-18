<template>
    <div class="proposals_detail_wrap">
        <div :class="proposalsDetailWrap"
                style="margin-bottom: 0.3rem; flex-direction: column; align-items: flex-start;">
            <p class="proposals_information_content_title">Proposal Information</p>
            <div class="proposals_detail_information_wrap">
                <div class="information_props_wrap">
                    <span class="information_props">Proposal ID :</span>
                    <span class="information_value information_show_trim">
                        <span class="information_pre">{{proposalsId}}</span>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Title :</span>
                    <span class="information_value information_show_trim">
                        <span class="information_pre">{{title}}</span>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Proposer :</span>
                    <span v-show="proposer !== '--'"
                          class="information_value information_show_trim jump_route">
                        <router-link :to="addressRoute(proposer)">{{proposer}}</router-link>

                    </span>
                    <span v-show="proposer == '--'"
                          class="information_value information_show_trim ">{{proposer}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Submit Hash :</span>
                    <span v-show="submitHash !== '--'"
                          class="information_value information_show_trim jump_route">
                        <router-link :to="`/tx?txHash=${submitHash}`">{{submitHash}}</router-link>
                    </span>
                    <span v-show="submitHash == '--'"
                          class="information_value information_show_trim ">{{submitHash}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Type :</span>
                    <span class="information_value">{{type}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Status :</span>
                    <span class="information_value">
                        {{status}}
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Submit Time :</span>
                    <span class="information_value">{{submitAge}} <span v-show="submitAge">(</span>{{submitTime}}<span v-show="submitAge">)</span></span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Deposit Endtime :</span>
                    <span class="information_value">{{depositEndAge}} <span v-show="depositEndAge">(</span>{{depositEndTime}}<span v-show="depositEndAge">)</span>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Total Deposit :</span>
                    <span class="information_value">{{totalDeposit}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Voting Starttime :</span>
                    <span class="information_value">{{votingStartAge}} <span v-show="votingStartAge">(</span>{{votingStartTime}}<span v-show="votingStartAge">)</span>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Voting Endtime :</span>
                    <span class="information_value">{{votingEndAge}} <span v-show="votingEndAge">(</span>{{votingEndTime}}<span v-show="votingEndAge">)</span></span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Description :</span>
                    <span class="information_value">
                        <pre class="information_pre information_show_trim" v-html="description"></pre>
                    </span>
                </div>

                <div v-show="type === 'SoftwareUpgrade'">
                    <div class="information_props_wrap">
                        <span class="information_props">Software :</span>
                        <span class="information_value information_show_trim">
                            <a class="information_link"
                               :href="software"
                               target="_blank">{{software}}</a>
                        </span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Version :</span>
                        <span class="information_value information_show_trim">
                            <span class="information_pre">{{version}}</span>
                        </span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Switch Height :</span>
                        <span class="information_value information_show_trim">
                            <span class="information_pre">{{switchHeight}}</span>
                        </span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Threshold :</span>
                        <span class="information_value information_show_trim">
                            <span class="information_pre">{{threshold}}</span>
                        </span>
                    </div>
                </div>
                <div class="parameter_container"
                     v-show="type === 'ParameterChange'">
                    <div class="information_props_wrap">
                        <span class="information_props">Parameter Details :</span>
                        <textarea :rows="textareaRows"
                                  v-model="parameterValue"
                                  readonly
                                  spellcheck="false"
                                  class="parameter_detail_content">
            </textarea>
                    </div>
                </div>
            </div>
        </div>
        <div class="card_container">
            <div v-show="depositorObj">
                <keep-alive>
                    <m-deposit-card :depositObj="depositorObj" :burnPercent="burnPercent"></m-deposit-card>
                </keep-alive>
            </div>
            <div v-show="votingObj">
                <keep-alive>
                    <m-voting-card></m-voting-card>
                </keep-alive>
            </div>
        </div>
        <div :class="['proposal_table', proposalsDetailWrap, $store.state.isMobile ? 'mobile_proposals_table_container' : '']"
             v-if="!showNoData">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <div class="proposals_table_title_div"
                     style="margin-top: 0;">Voters</div>
                <div class="voting_options">
                    <span>Yes:
                        <span>{{voteDetailsYes}}</span>
                    </span>|<span>No:
                        <span>{{voteDetailsNo}}</span>
                    </span>|<span>NoWithVeto:
                        <span>{{voteDetailsNoWithVeto}}</span>
                    </span>|<span>Abstain:
                        <span>{{voteDetailsAbstain}}</span>
                    </span>
                </div>
            </div>
            <div class="proposals_detail_table_wrap">
                <spin-component :showLoading="showLoading" />
                <m-proposals-detail-table :items="items"
                                          fields="votersFields"></m-proposals-detail-table>
                <div v-show="showNoData"
                     class="no_data_show">
                    <img src="../assets/no_data.svg" alt="">
                </div>
            </div>
            <div class="table_pagination">
                <b-pagination v-model="currentPage"
                              :total-rows="itemTotal"
                              :per-page="perPage"></b-pagination>
            </div>
        </div>
        <div :class="['proposal_table', proposalsDetailWrap]"
             v-if="!depositorShowNoData">
            <div class="proposals_table_title_div"
                 style="margin-top: 0;">Depositors</div>
            <div class="proposals_detail_table_wrap">
                <spin-component :showLoading="showLoading" />
                <m-proposals-detail-table :items="depositorItems"
                                          fields="depositorsFields"></m-proposals-detail-table>
                <div v-show="depositorShowNoData"
                     class="no_data_show">
                    <img src="../assets/no_data.svg" alt="">
                </div>
            </div>
            <div class="table_pagination">
                <b-pagination v-model="depositorCurrentPage"
                              :total-rows="depositorItemsTotal"
                              :per-page="perPage"></b-pagination>
            </div>
        </div>
    </div>
</template>

<script>
import Tools from '../util/Tools';
import Service from "../service";
import BlocksListTable from './table/BlocksListTable.vue';
import SpinComponent from './commonComponents/SpinComponent';
import Constant from "../constant/Constant";
import MProposalsDetailTable from './table/MProposalsDetailTable.vue';
import MDepositCard from "./commonComponents/MDepositCard";
import MVotingCard from "./commonComponents/MVotingCard";

export default {
    components: {
	    MVotingCard,
	    MDepositCard,
        BlocksListTable,
        SpinComponent,
        MProposalsDetailTable
    },
    data () {
        return {
            devicesWidth: window.innerWidth,
            proposalsDetailWrap: 'personal_computer_transactions_detail',
            items: [],
            depositorItems: [],
            showLoading: false,
            showNoData: false,
            depositorShowNoData: false,
            proposalsId: "",
            title: "",
            type: "",
            status: "",
            submitBlock: "",
            submitTime: "",
            totalDeposit: "",
            votingStartBlock: "",
            description: "",
            voteDetailsYes: "",
            voteDetailsNo: "",
            voteDetailsNoWithVeto: "",
            voteDetailsAbstain: "",
            proposer: "",
            submitHash: "",
            depositEndTime: "",
            votingStartTime: "",
            votingEndTime: "",
            submitAge: '',
            depositEndAge: '',
            votingStartAge: '',
            votingEndAge: '',
            proposalTimer: null,
            software: ' ',
            version: ' ',
            switchHeight: ' ',
            threshold: ' ',
            textareaRows: '2',
            parameterValue: '',
            currentPage: 1,
            depositorCurrentPage: 1,
            itemTotal: 0,
            depositorItemsTotal: 0,
            perPage: 10,
            depositorObj: null,
	        votingObj: null,
	        burnPercent: 0
        }
    },
    beforeMount () {
        Tools.scrollToTop();
        this.$nextTick(() => {
            this.computedProposalsDetailWrap();
        });
    },
    watch: {
        currentPage (newVal) {
            this.getVoter();
        },
        depositorCurrentPage (newVal) {
            this.getDepositor();
        },
        '$store.state.isMobile' (newVal) {
            this.computedProposalsDetailWrap();
        }
    },
    methods: {
        computedProposalsDetailWrap () {
            if (!this.$store.state.isMobile) {
                this.proposalsDetailWrap = 'personal_computer_transactions_detail_wrap';
            } else {
                this.proposalsDetailWrap = 'mobile_transactions_detail_wrap';
            }
        },
        flShowProposalTime (proposalTimeName, status) {
            if (status === 'Rejected' || status === 'Passed' || status === 'VotingPeriod') {
                return true
            } else {
                if (proposalTimeName === 'depositEndTime' && status === 'DepositPeriod') return true;
                if (proposalTimeName === 'votingStartTime' && status === 'VotingPeriod') return true;
                if (proposalTimeName === 'votingEndTime' && status === 'VotingPeriod') return true;
            }
        },
        formatProposalTime (time) {
            let currentServerTime = new Date().getTime() + this.diffMilliseconds;
            if (time && new Date(time).getTime() < currentServerTime) {
                return Tools.formatAge(currentServerTime, time, Constant.SUFFIX);
            }
        },
        getVoter () {
            this.showNoData = false;
            this.items = [];
            Service.commonInterface({                proposalDetailVoterTx: {
                    proposalId: this.$route.params.proposal_id,
                    pageNumber: this.currentPage,
                    perPageSize: this.perPage,
                }            }, (data) => {
                try {
                    if (data.items && data.items.length > 0) {
                        this.itemTotal = data.total;
                        this.items = data.items.map(item => {
                            let votingListItemTime = (new Date(item.timestamp).getTime()) > 0 ? Tools.format2UTC(item.timestamp) : '--';
                            return {
                                moniker: item.moniker,
                                Voter: item.voter,
                                Vote_Option: item.option,
                                Tx_Hash: item.tx_hash,
                                Time: votingListItemTime
                            }
                        });
                    } else {
                        this.items = [];
                        this.showNoData = true;
                    }
                } catch (e) {
                    this.items = [];
                    this.showNoData = true;
                }
            })
        },
        getDepositor () {
            this.depositorShowNoData = false;
            this.depositorItems = [];
            Service.commonInterface({                proposalDetailDepositorTx: {
                    proposalId: this.$route.params.proposal_id,
                    pageNumber: this.depositorCurrentPage,
                    perPageSize: this.perPage,
                }            }, (data) => {
                try {
                    if (data.items && data.items.length > 0) {
                        this.depositorItemsTotal = data.total;
                        this.depositorItems = data.items.map(item => {
                            let votingListItemTime = (new Date(item.timestamp).getTime()) > 0 ? Tools.format2UTC(item.timestamp) : '--';
                            let amount = item.amount && item.amount[0] ? `${this.$options.filters.amountFromat(item.amount[0])}` : '';
                            if (amount) {
                                let arr = amount.split(' ');
                                amount = `${Tools.subStrings(arr[0], 2)} ${arr[1]}`;
                            }
                            return {
                                Depositor: item.from,
                                Depositor_tooltip: item.moniker ? '' : item.from,
                                moniker: item.moniker,
                                Amount: amount || '--',
                                Type: item.type,
                                Tx_Hash: item.hash,
                                Time: votingListItemTime
                            }
                        });
                    } else {
                        this.depositorItems = [];
                        this.depositorShowNoData = true;
                    }
                } catch (e) {
                    this.depositorItems = [];
                    this.depositorShowNoData = true;
                }
            })
        },
        getProposalsInformation () {
            this.showLoading = true;
            Service.commonInterface({                proposalDetail: {
                    proposalId: this.$route.params.proposal_id
                }            }, (data) => {
                this.showLoading = false;
                try {
                    if (data) {
                        if (data.proposal.proposal_id === 0) {
                            this.proposalsId = '--';
                            this.title = '--';
                            this.type = '--';
                            this.status = '--';
                            this.proposer = '--';
                            this.submitHash = '--';
                            this.submitTime = '--';
                            this.depositEndTime = '--';
                            this.votingStartTime = '--';
                            this.votingEndTime = "--";
                            this.description = '--';
                            this.voteDetailsYes = '--';
                            this.voteDetailsNo = '--';
                            this.voteDetailsNoWithVeto = '--';
                            this.voteDetailsAbstain = '--';
                            this.voteDetailsAbstain = '--';
                            this.totalDeposit = '--';
                        } else {
                            let that = this;
                            clearInterval(this.proposalTimer);
                            this.proposalTimer = setInterval(function () {
                                that.submitAge = that.formatProposalTime(data.proposal.submit_time ? data.proposal.submit_time : '');
                                that.depositEndAge = that.formatProposalTime(that.flShowProposalTime('depositEndTime', data.proposal.status) ? data.proposal.deposit_end_time : '');
                                that.votingStartAge = that.formatProposalTime(that.flShowProposalTime('votingStartTime', data.proposal.status) ? data.proposal.voting_start_time : '');
                                that.votingEndAge = that.formatProposalTime(that.flShowProposalTime('votingEndTime', data.proposal.status) ? data.proposal.voting_end_time : '');
                            }, 1000);
                            this.submitAge = this.formatProposalTime(data.proposal.submit_time ? data.proposal.submit_time : '');
                            this.depositEndAge = this.formatProposalTime(this.flShowProposalTime('depositEndTime', data.proposal.status) ? data.proposal.deposit_end_time : '');
                            this.votingStartAge = this.formatProposalTime(this.flShowProposalTime('votingStartTime', data.proposal.status) ? data.proposal.voting_start_time : '');
                            this.votingEndAge = this.formatProposalTime(this.flShowProposalTime('votingEndTime', data.proposal.status) ? data.proposal.voting_end_time : '');
                            this.software = data.proposal.software;
                            this.version = data.proposal.version;
                            this.switchHeight = data.proposal.switch_height;
                            this.threshold = data.proposal.threshold ? `${Number(data.proposal.threshold) * 100}%` : '';
                            this.proposalsId = data.proposal.proposal_id === 0 ? "--" : data.proposal.proposal_id;
                            this.title = data.proposal.title;
                            this.type = data.proposal.type;
                            this.status = data.proposal.status;
                            this.proposer = data.proposal.proposer ? data.proposal.proposer : "--";
                            this.submitHash = data.proposal.tx_hash ? data.proposal.tx_hash : "--";
                            this.submitTime = data.proposal.submit_time ? Tools.format2UTC(data.proposal.submit_time) : '--';
                            this.depositEndTime = that.flShowProposalTime('depositEndTime', data.proposal.status) ? Tools.format2UTC(data.proposal.deposit_end_time) : '--';
                            this.votingStartTime = that.flShowProposalTime('votingStartTime', data.proposal.status) ? Tools.format2UTC(data.proposal.voting_start_time) : '--';
                            this.votingEndTime = that.flShowProposalTime('votingEndTime', data.proposal.status) ? Tools.format2UTC(data.proposal.voting_end_time) : '--';
                            this.description = data.proposal.description ? data.proposal.description : " -- ";
                            this.voteDetailsYes = data.proposal.status === "DepositPeriod" ? "--" : data.result.Yes;
                            this.voteDetailsNo = data.proposal.status === "DepositPeriod" ? "--" : data.result.No;
                            this.voteDetailsNoWithVeto = data.proposal.status === "DepositPeriod" ? "--" : data.result.NoWithVeto;
                            this.voteDetailsAbstain = data.proposal.status === "DepositPeriod" ? "--" : data.result.Abstain;
                            if (data.proposal && data.proposal.total_deposit.length !== 0) {
                                this.totalDeposit = `${Tools.formatPriceToFixed(Tools.convertScientificNotation2Number(Tools.formatNumber(data.proposal.total_deposit[0].amount)))} ${Tools.formatDenom(data.proposal.total_deposit[0].denom).toUpperCase()}`;
                            } else {
                                this.totalDeposit = "";
                            }
                            this.burnPercent = data.proposal.burn_percent;
                            if (data.proposal.type === 'ParameterChange') {
                                for (let index = 0; index < data.proposal.parameters.length; index++) {
                                    this.parameterValue += `${data.proposal.parameters[index].subspace}/${data.proposal.parameters[index].key} = ${data.proposal.parameters[index].value}\n`
                                }
                                let defaultTextareaRows = 2;
                                this.textareaRows = data.proposal.parameters.length + defaultTextareaRows;
                            }
                        }
                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },
        getDepositorInfomation(){
        	Service.commonInterface({proposalDetailDepositor:{
			        proposalId: this.$route.params.proposal_id
                }},(res) => {
                this.depositorObj = res
            })
        }
    },
    mounted () {
        this.getProposalsInformation();
        this.getVoter();
        this.getDepositor();
        this.getDepositorInfomation()
    }
}
</script>

<style scoped lang="scss">
@import "../style/mixin.scss";
.card_container{
    display: flex;
    width: 100%;
    margin: 0 auto;
    max-width: 12.8rem;
    div{
        flex: 1;
    }
}
@media screen and (max-width: 910px){
    .card_container{
        display: flex;
        width: 100%;
        margin: 0 auto;
        max-width: 12.8rem;
        flex-direction: column;
    }
}
.proposals_detail_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    .personal_computer_transactions_detail_wrap {
        width: 100% !important;
        .proposals_information_content_title {
            padding-left: 0.2rem !important;
            height: 0.7rem !important;
            line-height: 0.7rem !important;
            font-size: 0.18rem !important;
            color: #000000;
            margin-bottom: 0;
            @include fontWeight;
        }
        @include pcCenter;
        .proposals_detail_information_wrap {
            padding: 0.2rem 0.2rem 0.08rem;
            border: 1px solid rgba(215, 217, 224, 1) !important;
            .information_props_wrap {
                @include flex;
                line-height: 0.2rem;
                margin-bottom: 0.12rem;
                .information_props {
                    min-width: 1.5rem;
                    color: var(--contentColor);
                }
                .flag_item_left {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../assets/left.png") no-repeat 0 1px;
                    margin-right: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_left_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    margin-right: 0.05rem;
                    cursor: pointer;
                    background: url("../assets/left_disabled.png") no-repeat 0
                        1px;
                }
                .flag_item_right {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../assets/right.png") no-repeat 0 0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_right_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../assets/right_disabled.png") no-repeat 0
                        0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
            }
        }
        .proposals_detail_table_wrap {
            margin-bottom: 0.2rem;
            width: 100%;
            overflow-x: auto;
            overflow-y: hidden;
            .table_wrap {
                min-width: 9.6rem;
            }
            .no_data_show {
                width: 100%;
                @include flex;
                justify-content: center;
                border-top: 0.01rem solid #eee;
                border-bottom: 0.01rem solid #eee;
                font-size: 0.14rem;
                height: 1rem;
                align-items: center;
            }
        }
    }

    .mobile_transactions_detail_wrap {
        width: 100%;
        @include flex;
        flex-direction: column;
        padding: 0 0.1rem;
        &:nth-child(1) {
            display: flex;
            flex-direction: row;
            align-items: center;
        }
        .proposals_table_title_div {
            margin-left: 0.1rem !important;
        }
        .proposals_information_content_title {
            height: 0.7rem !important;
            line-height: 0.7rem !important;
            font-size: 0.18rem !important;
            color: #000000;
            margin-bottom: 0;
            padding-left: 0.1rem;
            @include fontWeight;
        }
        .proposals_detail_table_wrap {
            width: 100%;
            overflow-x: auto;
            -webkit-overflow-scrolling: touch;
            margin-bottom: 0.2rem;
            .no_data_show {
                @include flex;
                width: 100%;
                margin: auto;
                justify-content: center;
                border-top: 0.01rem solid #eee;
                border-bottom: 0.01rem solid #eee;
                font-size: 0.14rem;
                height: 3rem;
                align-items: center;
            }
        }
        .proposals_detail_information_wrap {
            border: 1px solid rgba(215, 217, 224, 1) !important;
            padding: 10px;
            width: 100%;
            .information_props_wrap {
                @include flex;
                flex-direction: column;
                line-height: 0.2rem;
                .information_value {
                    overflow-x: auto;
                    -webkit-overflow-scrolling: touch;
                }
                .flag_item_left {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../assets/left.png") no-repeat 0 1px;
                    margin-right: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_left_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    margin-right: 0.05rem;
                    cursor: pointer;
                    background: url("../assets/left_disabled.png") no-repeat 0
                        1px;
                }
                .flag_item_right {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../assets/right.png") no-repeat 0 0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_right_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../assets/right_disabled.png") no-repeat 0
                        0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
            }
        }
        .transactions_detail_wrap_hash_var {
            overflow-x: auto;
            -webkit-overflow-scrolling: touch;
            height: 0.3rem;
            line-height: 0.3rem;
            font-size: 0.22rem;
            color: var(--contentColor);
        }
        .vote-details-content {
            width: 100%;
            overflow-x: auto;
            border-top: 1px solid #d6d9e0;
            display: flex;
            justify-content: space-between;
            height: 0.62rem;
            line-height: 0.62rem;
            .vote_content_container {
                min-width: 150%;
                display: flex;
                justify-content: space-between;
            }
        }
    }
}
.information_show_trim {
    color: var(--contentColor);
}
.information_value {
    color: var(--titleColor);
    word-break: break-all;
}
.vote-details-content {
    border-top: 1px solid #d6d9e0;
    display: flex;
    justify-content: space-between;
    height: 0.62rem;
    line-height: 0.62rem;
}
.total_num {
    font-size: 0.14rem;
    color: var(--contentColor);
    margin-left: 0.2rem;
}
.voting_options {
    display: flex;
    color: var(--contentColor);
    margin-bottom: 10px;
    flex-wrap: wrap;
    & > span {
        font-size: 0.14rem;
        color: var(--contentColor);
        @include fontWeight;
        padding: 0 0.18rem;
        white-space: nowrap;
        & > span {
            color: var(--contentColor);
        }
    }
}
.information_show_trim {
    white-space: pre-wrap;
}
.jump_route {
    a {
        color: var(--bgColor) !important;
    }
    cursor: pointer;
}
.vote_content_container {
    min-width: 100%;
    display: flex;
    justify-content: space-between;
}
pre {
    font-family: Arial !important;
}
.information_link {
    color: var(--bgColor) !important;
}
.parameter_detail_content {
    box-sizing: border-box;
    padding: 0.1rem;
    width: 97%;
    margin-right: 20%;
    background-color: #fcfcfc !important;
    border-color: #d7d9e0;
    border-radius: 1px;
}
.information_pre {
    color: var(--titleColor);
    word-wrap: break-word;
    word-break: break-all;
}
.proposals_table_title_div {
    font-size: 18px;
    margin: 30px 20px 10px;
}
.mobile_proposals_table_container {
    & > div {
        display: flex;
        flex-wrap: wrap;
        justify-content: flex-start !important;
    }
    .table_wrap {
        width: 100%;
    }
    .table_pagination {
        justify-content: flex-end !important;
    }
    .voting_options {
        & > span {
            padding: 0 0.08rem;
        }
    }
}
.table_pagination {
    display: flex;
    justify-content: flex-end;
}
.proposal_table {
    &:nth-last-child(2n) {
        margin-bottom: 0.2rem;
    }
    &:nth-last-child(1) {
        margin-bottom: 0.4rem;
    }
}
</style>
