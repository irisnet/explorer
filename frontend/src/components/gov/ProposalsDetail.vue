<template>
    <div class="proposals_detail_wrap">
        <page-title :title="pageTitle"
                    :content="title"
                    
                    :flShowMobileStyle="true"></page-title>
        <div :class="proposalsDetailWrap"
                style="margin-bottom: 0.3rem; flex-direction: column; align-items: flex-start;">
            <div class="level_container">
                <div class="proposals_detail_level">
                    <i v-if="levelValue === 'Critical'" style="color:#FF5569;font-size: 0.16rem;" class="iconfont iconCritical"></i>
                    <i v-if="levelValue === 'Important'" style="color: #FF8000;font-size: 0.16rem;" class="iconfont iconImportant"></i>
                    <i v-if="levelValue === 'Normal'" style="color:#45B4FF;font-size: 0.16rem;" class="iconfont iconNormal"></i>
                    <span >{{type}}</span>
                </div>
                <div class="step_content">
                    <span v-if="status === 'VotingPeriod'">
                        <i style="color:rgb(5, 128, 211);" class="iconfont iconDepositPeriod"></i>VotingPeriod
                    </span>
                    <span v-if="status === 'DepositPeriod'">
                        <i style="color:rgb(5, 128, 211);" class="iconfont iconDepositPeriod-liebiao"></i>DepositPeriod
                    </span>
                    <span v-if="status === 'Passed'" ><i style="color:#44C190;" class="iconfont iconPass"></i>Passed</span>
                    <span v-if="status === 'Rejected'"><i style="color:rgb(254, 138, 138);" class="iconfont iconVeto"></i>Rejected</span>
                </div>
                <div class="time_content">
                    <span v-if="status === 'VotingPeriod' && flShowVotingHourLeft">
                        <i style="color:#5AC8FA;" class="iconfont iconHoursLeft"></i>{{votingHourLeft}} Left
                    </span>
                    <span v-if="status === 'DepositPeriod' && flShowDepositHourLeft">
                        <i style="color:#5AC8FA;" class="iconfont iconHoursLeft"></i>{{depositHourLeft}} Left
                    </span>
                </div>
            </div>
            <div class="proposals_detail_information_wrap" style="background: #fff">
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
               <div v-show="type === 'CommunityTaxUsage' ">
                   <div class="information_props_wrap">
                       <span class="information_props">Usage :</span>
                       <span class="information_value">{{usageValue}}</span>
                   </div>
                   <div class="information_props_wrap">
                       <span class="information_props">DestAddress :</span>
                       <span class="information_value">{{destAddressValue}}</span>
                   </div>
                   <div class="information_props_wrap">
                       <span class="information_props">Percent :</span>
                       <span class="information_value">{{percentValue}}</span>
                   </div>
               </div>
                <div v-show="type === 'TokenAddition'">
                    <div class="information_props_wrap">
                        <span class="information_props">Symbol :</span>
                        <span class="information_value">{{symbolValue}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">CanonicalSymbol :</span>
                        <span class="information_value">{{canonicalSymbolValue}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Name :</span>
                        <span class="information_value">{{nameValue}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Decimal :</span>
                        <span class="information_value">{{decimalValue}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">MinUnitAlias :</span>
                        <span class="information_value">{{minUnitAliasValue}}</span>
                    </div>
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
                     v-show="type === 'Parameter'">
                    <div class="information_props_wrap">
                        <span class="information_props">Parameter Details :</span>
                        <span v-model="parameterValue" v-html="parameterValue"></span>
                    </div>
                </div>
            </div>
            <div class="proposal_detail_content">
                <div class="proposals_detail_information_wrap">
                    <div class="information_props_wrap">
                        <span class="information_props">Submit Time :</span>
                        <span class="information_value">{{submitAge}} <span v-show="submitAge">(</span>{{submitTime}}<span v-show="submitAge">)</span></span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Deposit End Time :</span>
                        <span class="information_value">{{depositEndAge}}
                            <span v-show="flShowDepositHourLeft">{{depositHourLeft}} left </span>
                            <span v-show="depositEndAge || flShowDepositHourLeft">(</span>{{depositEndTime}}<span v-show="depositEndAge || flShowDepositHourLeft">)</span>
                        </span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Total Deposit :</span>
                        <span class="information_value">{{totalDeposit}} <span v-show="burnValue">({{burnValue}}% Burned)</span></span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Voting Start Time :</span>
                        <span class="information_value">{{votingStartAge}} <span v-show="votingStartAge">(</span>{{votingStartTime}}<span v-show="votingStartAge">)</span>
                        </span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Voting End Time :</span>
                        <span class="information_value">{{votingEndAge}}
                            <span v-show="flShowVotingHourLeft">{{votingHourLeft}} left </span>
                            <span v-show="votingEndAge || flShowVotingHourLeft">(</span>{{votingEndTime}}<span v-show="votingEndAge || flShowVotingHourLeft">)</span>
                        </span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Penalty :</span>
                        <span class="information_value">{{penaltyValue}}</span>
                    </div>
                </div>
                <div class="proposals_detail_information_wrap">
                    <div class="information_props_wrap">
                        <span class="information_props">Status :</span>
                        <span class="information_value">{{status}}</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Participation :</span>
                        <span class="information_value">{{$store.state.currentParticipationValue && $store.state.currentParticipationValue !== 'NaN' ? $store.state.currentParticipationValue : '0.00'}}% (Threshold {{participationValue}})</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Yes :</span>
                        <span class="information_value">{{$store.state.currentYesValue && $store.state.currentYesValue !== 'NaN' ? $store.state.currentYesValue : '0.00'}}% (Threshold {{yesThresholdValue}})</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">No :</span>
                        <span class="information_value">{{$store.state.currentNoValue && $store.state.currentNoValue !== 'NaN' ? $store.state.currentNoValue : '0.00'}}%</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">NoWithVeto :</span>
                        <span class="information_value">{{$store.state.currentNoWithVetoValue && $store.state.currentNoWithVetoValue !== 'NaN' ? $store.state.currentNoWithVetoValue : '0.00'}}% (Threshold {{vetoThresholdValue}})</span>
                    </div>
                    <div class="information_props_wrap">
                        <span class="information_props">Abstain :</span>
                        <span class="information_value">{{$store.state.currentAbstainValue && $store.state.currentAbstainValue !== 'NaN' ? $store.state.currentAbstainValue : '0.00'}}%</span>
                    </div>

                </div>
            </div>
        </div>
        <div class="card_container">
            <div v-show="depositorObj" class="deposit_card_content_wrap">
                <keep-alive>
                    <m-deposit-card :levelValue="levelValue" :depositObj="depositorObj" :burnPercent="burnPercent" :status="status"></m-deposit-card>
                </keep-alive>
            </div>
            <div  v-show="votingObj" class="voting_mobile_content">
                <div v-show="votingObj" class="voting_proposal_card_content">
                    <keep-alive>
                        <m-voting-card :votingBarObj="votingObj"></m-voting-card>
                    </keep-alive>
                </div>
            </div>
        </div>
        <div :class="['proposal_table', proposalsDetailWrap, $store.state.isMobile ? 'mobile_proposals_table_container' : '']"
        >
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <div class="proposals_table_title_div"
                     style="margin-top: 0;">Voters</div>
                <ul class="filter_content">
                    <li class="tab_option"
                        v-for="(item,index) in filterTabArr"
                        :class="item.isActive ? 'blue_style' : ''"
                        @click="filterVoteTx(item.key,index)"
                    ><span>{{item.label}} {{item.value}}</span> <span>|</span></li>
                </ul>
                <div class="voting_options">
                    <span><i class="yes_option_style"></i>Yes:
                        <span>{{voteDetailsYes}}</span>
                    </span>|<span><i class="no_option_style"></i>No:
                        <span>{{voteDetailsNo}}</span>
                    </span>|<span><i class="no_with_veto_option_style"></i>NoWithVeto:
                        <span>{{voteDetailsNoWithVeto}}</span>
                    </span>|<span><i class="abstain_option_style"></i>Abstain:
                        <span>{{voteDetailsAbstain}}</span>
                    </span>
                </div>
            </div>
            <div class="proposals_detail_table_wrap">
                <m-proposals-detail-table :items="items"
                                          fields="votersFields"></m-proposals-detail-table>
                <div v-show="showNoData"
                     class="no_data_show">
                    <img src="../../assets/no_data.svg" alt="">
                </div>
            </div>
            <div class="table_pagination">
                <!--<b-pagination v-model="currentPage"
                              :total-rows="itemTotal"
                              :per-page="perPage"></b-pagination>-->
                <keep-alive>
                    <m-pagination
                            :pageSize="perPage"
                            :total="itemTotal"
                            :page="currentPageNum"
                            :page-change="pageChange"></m-pagination>
                </keep-alive>
            </div>
        </div>
        <div :class="['proposal_table', proposalsDetailWrap]"
             v-if="!depositorShowNoData">
            <div class="proposals_table_title_div"
                 style="margin-top: 0;">Depositors</div>
            <div class="proposals_detail_table_wrap">
                <m-proposals-detail-table :items="depositorItems"
                                          fields="depositorsFields"></m-proposals-detail-table>
                <div v-show="depositorShowNoData"
                     class="no_data_show">
                    <img src="../../assets/no_data.svg" alt="">
                </div>
            </div>
            <div class="table_pagination">
                <!--<m-pagination v-model="depositorCurrentPage"
                              :total-rows="depositorItemsTotal"
                              :per-page="perPage"></m-pagination>-->
                <keep-alive>
                    <m-pagination
                            :pageSize="perPage"
                            :total="depositorItemsTotal"
                            :page="currentDepositorPageNum"
                            :page-change="pageChangeDepositor"></m-pagination>
                </keep-alive>
            </div>
        </div>
    </div>
</template>

<script>
import Tools from '../../util/Tools';
import Service from "../../service";
import BlocksListTable from '../block/BlocksListTable.vue';
import Constant from "../../constant/Constant";
import MProposalsDetailTable from './MProposalsDetailTable.vue';
import MDepositCard from "../commontables/MDepositCard";
import MVotingCard from "../commontables/MVotingCard";
import MPagination from "../commontables/MPagination";
import PageTitle from "../pageTitle/PageTitle";
import pageTitleConfig from "../pageTitle/pageTitleConfig";

export default {
    components: {
        PageTitle,
	    MPagination,
	    MVotingCard,
	    MDepositCard,
        BlocksListTable,
        MProposalsDetailTable
    },
    data () {
        return {
            pageTitle:pageTitleConfig.GovProposalsProposalDetails,
	        symbolValue:'',
	        canonicalSymbolValue:'',
	        nameValue:'',
	        decimalValue:'',
	        minUnitAliasValue: '',
	        destAddressValue:'',
	        percentValue: '',
            devicesWidth: window.innerWidth,
            proposalsDetailWrap: 'personal_computer_transactions_detail',
            items: [],
            depositorItems: [],
            showNoData: false,
            depositorShowNoData: false,
	        levelValue:'',
	        participationValue:'',
	        yesThresholdValue:'',
	        vetoThresholdValue:'',
	        penaltyValue:'',
	        usageValue:'',
            burnValue: '',
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
	        currentPageNum:1,
	        currentDepositorPageNum:1,
            depositorCurrentPage: 1,
            itemTotal: 0,
            depositorItemsTotal: 0,
            perPage: 10,
            depositorObj: null,
	        votingObj: null,
            firstFilter:true,
	        burnPercent: 0,
            filterTabArr:[
                {
                	label:'All: ',
                    key:'all',
                    value:'',
                    isActive:true
                },
	            {
		            label:'Validator: ',
                    key:'validator',
		            value:'',
		            isActive:false
	            },
	            {
		            label:'Delegator: ',
                    key:'delegator',
		            value:'',
		            isActive:false
	            }
            ],
            filterTab:'all',
	        votingHourLeft:'',
	        depositHourLeft:'',
	        flShowVotingHourLeft:false,
	        flShowDepositHourLeft:false,
            depositHourTimer:null,
            votingHourTimer:null,
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
	        this.currentPage = newVal;
            this.getVoter();
        },
        depositorCurrentPage (newVal) {
            this.getDepositor();
        },
        '$store.state.isMobile' (newVal) {
            this.computedProposalsDetailWrap();
        },
	    flShowVotingHourLeft(flShowVotingHourLeft){
        	if(!flShowVotingHourLeft){
                this.getVoter();
		        this.getDepositor();
		        this.getProposalsInformation();
		        this.getDepositorInformation();
		        this.getVotingBarInformation();
            }
        },
    },
    methods: {
	    pageChange(pageNum){
	    	this.currentPageNum = pageNum;
	    	this.getVoter()
        },
	    pageChangeDepositor(pageNum){
		    this.currentDepositorPageNum = pageNum;
		    this.getDepositor()
	    },
	    filterVoteTx(item,index){
		    this.currentPageNum = 1;
		    this.filterTab = item;
            this.resetActiveStyle();
            this.filterTabArr[index].isActive = true;
            this.getVoter();
        },
	    resetActiveStyle(){
	    	this.filterTabArr.map( item => {
	    		return item.isActive = false
            })
        },
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
	    setStats(stats){
		    if(stats){
			    this.voteDetailsYes =  stats.yes;
			    this.voteDetailsNo = stats.no;
			    this.voteDetailsNoWithVeto = stats.no_with_veto;
			    this.voteDetailsAbstain = stats.abstain;
		    }
		    this.filterTabArr.map( item => {
		    	return item.value = stats[item.key]
            })

        },

        getVoter () {
	        Service.commonInterface({proposalDetailVoterTxByFilter:{
			        proposalId: this.$route.params.proposal_id,
			        pageNumber: this.currentPageNum,
			        perPageSize: this.perPage,
			        voterType:this.filterTab
		        }},(data) => {
		        try {
		            if(data){
                        this.setStats(data.stats)
                        if (data.items && data.items.length > 0) {
                            this.showNoData = false;
                            this.itemTotal = data.total;
                            this.items = data.items.map(item => {
                                let votingListItemTime = (new Date(item.timestamp).getTime()) > 0 ? Tools.format2UTC(item.timestamp) : '--';
                                return {
                                    moniker: item.moniker,
                                    Block:item.height,
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
                    }

		        }catch (e) {
			        console.error(e)
		        }

	        })
        },
        getDepositor () {
            this.depositorShowNoData = false;
            this.depositorItems = [];
            Service.commonInterface({                proposalDetailDepositorTx: {
                    proposalId: this.$route.params.proposal_id,
                    pageNumber: this.currentDepositorPageNum,
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
            Service.commonInterface({                proposalDetail: {
                    proposalId: this.$route.params.proposal_id
                }            }, (data) => {
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
                            this.symbolValue = '--';
                            this.canonicalSymbolValue = '--';
                            this.nameValue = '--';
                            this.decimalValue = '--';
                            this.minUnitAliasValue = '--';
                            this.destAddressValue = '--';
                            this.percentValue = '--'
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
                            this.participationValue = `${Tools.formatPercent(data.proposal.participation)}%`;
                            this.yesThresholdValue = `${Tools.formatPercent(data.proposal.yes_threshold)}%`;
                            this.vetoThresholdValue = `${Tools.formatPercent(data.proposal.veto_threshold)}%`;
                            this.penaltyValue = `${(Number(data.proposal.penalty) * 100).toFixed(2)}%`;
	                        this.usageValue = data.proposal.usage ? data.proposal.usage : '--';
	                        this.burnValue = data.proposal.burn_percent ? (Number(data.proposal.burn_percent) *100).toFixed(2) : '';
	                        this.version = data.proposal.version;
                            this.switchHeight = data.proposal.switch_height;
                            this.threshold = data.proposal.threshold ? `${Number(data.proposal.threshold) * 100}%` : '';
                            this.proposalsId = data.proposal.proposal_id === 0 ? "--" : data.proposal.proposal_id;
                            this.title = `#${data.proposal.proposal_id} ${data.proposal.title}`;
                            this.type = data.proposal.type;
                            this.status = data.proposal.status;
                            this.levelValue = data.proposal.level;
                            this.proposer = data.proposal.proposer ? data.proposal.proposer : "--";
                            this.submitHash = data.proposal.tx_hash ? data.proposal.tx_hash : "--";
                            this.submitTime = data.proposal.submit_time ? Tools.format2UTC(data.proposal.submit_time) : '--';
                            this.depositEndTime = that.flShowProposalTime('depositEndTime', data.proposal.status) ? Tools.format2UTC(data.proposal.deposit_end_time) : '--';
                            this.votingStartTime = that.flShowProposalTime('votingStartTime', data.proposal.status) ? Tools.format2UTC(data.proposal.voting_start_time) : '--';
                            this.votingEndTime = that.flShowProposalTime('votingEndTime', data.proposal.status) ? Tools.format2UTC(data.proposal.voting_end_time) : '--';
                            this.description = data.proposal.description ? data.proposal.description : "--";
	                        this.percentValue = data.proposal.percent ? `${(Number(data.proposal.percent) * 100).toFixed(2)}%` :'';
	                        this.symbolValue = data.proposal.symbol ? data.proposal.symbol : "--";
	                        this.canonicalSymbolValue = data.proposal.canonical_symbol ? data.proposal.canonical_symbol : "--";
	                        this.nameValue = data.proposal.name ? data.proposal.name : "--";
	                        this.decimalValue = data.proposal.decimal ? data.proposal.decimal : "--";
	                        this.minUnitAliasValue = data.proposal.min_unit_alias ? data.proposal.min_unit_alias : "--";
	                        this.destAddressValue = data.proposal.dest_address ? data.proposal.dest_address : '--';
	                        if (data.proposal && data.proposal.total_deposit.length !== 0) {
                                this.totalDeposit = `${Tools.formatPriceToFixed(Tools.convertScientificNotation2Number(Tools.formatNumber(data.proposal.total_deposit[0].amount)))} ${Tools.formatDenom(data.proposal.total_deposit[0].denom).toUpperCase()}`;
                            } else {
                                this.totalDeposit = "";
                            }
                            this.burnPercent = data.proposal.burn_percent;
	                        this.parameterValue = '';
                            if (data.proposal.type === 'Parameter') {
                                for (let index = 0; index < data.proposal.parameters.length; index++) {
                                    this.parameterValue += `${data.proposal.parameters[index].subspace}/${data.proposal.parameters[index].key} = ${data.proposal.parameters[index].value}\n`
                                }
                            }
                        }
                    }
                } catch (e) {
                    console.error(e)
                }
            })
        },
        getDepositorInformation(){
        	Service.commonInterface({proposalDetailDepositor:{
			        proposalId: this.$route.params.proposal_id
                }},(res) => {
        		try {
                    if(res){
	                    this.depositorObj = res;
	                    this.getDepositTime(res.deposit_end_time)
                    }
		        }catch (e) {
                    console.error(e)
		        }
            })
        },
	    getVotingEndTime(time){
		    if(time){
		    	clearInterval(this.votingHourTimer);
		    	let that = this;
			    let currentServerTime = new Date().getTime() + this.diffMilliseconds;
			    if(new Date(time).getTime() >  currentServerTime){
				    that.votingHourLeft = Tools.formatAge(new Date(time).getTime(),currentServerTime).trim();
				    that.flShowVotingHourLeft = true;
			    }else {
				    that.flShowVotingHourLeft = false;
			    }
			    this.votingHourTimer = setInterval(() => {
				    currentServerTime = new Date().getTime() + this.diffMilliseconds;
				    if(new Date(time).getTime() >  currentServerTime){
					    that.votingHourLeft = Tools.formatAge(new Date(time).getTime(),currentServerTime).trim();
					    that.flShowVotingHourLeft = true;
				    }else {
					    that.flShowVotingHourLeft = false;
				    }
                },1000)

		    }
	    },
	    getDepositTime(time){
		    if(time){
		    	clearInterval(this.depositHourTimer);
			    let that = this;
			    this.depositHourTimer = setInterval(() => {
				    let currentServerTime = new Date().getTime() + this.diffMilliseconds;
				    if(new Date(time).getTime() >  currentServerTime){
					    that.depositHourLeft = Tools.formatAge(new Date(time).getTime(),currentServerTime);
					    that.flShowDepositHourLeft = true;
				    }else {
					    that.flShowDepositHourLeft = false;
				    }
                },1000)
		    }
	    },
	    getVotingBarInformation(){
        	Service.commonInterface({proposalDetailVotingBar:{
			        proposalId: this.$route.params.proposal_id
                }},(res) => {
        		try {

			        if(res){
				        this.votingObj = res;
				        this.getVotingEndTime(res.voting_end_time);
			        }else {
                        this.$store.commit('currentParticipationValue',0);
                        this.$store.commit('currentYesValue',0);
                        this.$store.commit('currentNoWithVetoValue',0);
                        this.$store.commit('currentAbstainValue',0);
                        this.$store.commit('currentNoValue',0);
                    }
		        }catch (e) {
                    console.error(e)
		        }

            })
        }
    },
    mounted () {
        this.getProposalsInformation();
        this.getVoter();
        this.getDepositor();
        this.getDepositorInformation();
        this.getVotingBarInformation();
    }
}
</script>

<style scoped lang="scss">
@import "../../style/mixin";
.card_container{
    display: flex;
    width: 100%;
    margin: 0 auto;
    max-width: 12.8rem;
    div{
        flex: 1;
    }
    .deposit_card_content_wrap{
        padding: 0;
    }
    .voting_proposal_card_content{
        min-width: 5rem;
    }
}
.card_container{
    .voting_mobile_content{
        margin-left: 0.1rem;
    }
}

@media screen and (max-width: 910px){
    .card_container{
        display: flex;
        width: 100%;
        margin: 0 auto;
        max-width: 12.8rem;
        flex-direction: column;
        .deposit_card_content_wrap{
            padding: 0 0.1rem;
        }
        .voting_mobile_content{
            margin: 0 0.1rem;
            overflow-x: auto;
        }
    }

}
.proposals_detail_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    .personal_computer_transactions_detail_wrap {
        width: 100% !important;
        margin-top: 0.64rem;
        .proposal_detail_content{
            display: flex;
            margin-top: 0.2rem;
            .proposals_detail_information_wrap{
                flex: 1;
                border: 0.01rem solid #d7d9e0;
                background: #fff;
            }
            .proposals_detail_information_wrap:first-child{
                margin-right: 0.1rem;
                background: #fff;
            }
        }
        .proposals_information_content_title {
            padding-left: 0.2rem !important;
            height: 0.7rem !important;
            line-height: 0.7rem !important;
            font-size: 0.18rem !important;
            color: #515a6e;
            font-weight: bold;
            margin-bottom: 0;
            .proposal_title{
                padding-left: 0.1rem;
            }
        }
        @include pcCenter;
        .proposals_detail_information_wrap {
            padding: 0.2rem 0.2rem 0.08rem;
            border-right: 1px solid rgba(215, 217, 224, 1) ;
            border-left: 1px solid rgba(215, 217, 224, 1) ;
            border-top: 1px solid rgba(215, 217, 224, 1) ;
            border-bottom: 1px solid rgba(215, 217, 224, 1) ;
            background: #fff;
            .proposals_detail_level{
                padding: 0 0 0.2rem 0;
                i{
                    img{
                        width: 0.16rem;
                    }
                }
                span{
                    color:#22252A;
                    font-size: 0.14rem;
                    line-height: 0.2rem;
                    margin-left: 0.1rem;
                }
            }

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
                    background: url("../../assets/left.png") no-repeat 0 1px;
                    margin-right: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_left_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    margin-right: 0.05rem;
                    cursor: pointer;
                    background: url("../../assets/left_disabled.png") no-repeat 0
                        1px;
                }
                .flag_item_right {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../../assets/right.png") no-repeat 0 0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_right_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../../assets/right_disabled.png") no-repeat 0
                        0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
            }
        }
        .level_container{
            display: flex;
            margin: 0 0 0.1rem 0;
            .proposals_detail_level{
                margin: 0 0.4rem 0 0.2rem;
                i{
                    padding-right: 0.1rem;
                }
            }
            .step_content{
                margin-right: 0.4rem;
                span{
                    i{
                        padding-right: 0.1rem;
                    }
                }

            }
            .time_content{
                margin-right: 0.4rem;
                span{
                    i{
                        padding-right: 0.1rem;
                    }
                }
            }

        }
        .proposals_detail_table_wrap {
            margin-bottom: 0.2rem;
            width: 100%;
            overflow-x: auto;
            overflow-y: hidden;
            background: #fff;
            .table_wrap {
                min-width: 9.6rem;
            }
            .no_data_show {
                background: #fff;
                width: 100%;
                min-height: 3rem;
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
        box-sizing: border-box;
        &:nth-child(1) {
            display: flex;
            flex-direction: row;
            align-items: center;
        }
        .proposals_table_title_div {
            margin-left: 0.1rem !important;
        }

        .proposals_information_content_title {
            font-size: 0.18rem !important;
            color: #515a6e;
            font-weight: bold;
            padding: 0.2rem 0 0.2rem 0.1rem;
                span:first-child{
                    padding-right: 0.1rem;
                }
        }
        .proposals_detail_table_wrap {
            width: 100%;
            overflow-x: auto;
            -webkit-overflow-scrolling: touch;
            margin-bottom: 0.2rem;
            background: #fff;
            .no_data_show {
                background: #fff;
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
            border-right: 1px solid rgba(215, 217, 224, 1) ;
            border-left: 1px solid rgba(215, 217, 224, 1) ;
            border-top: 1px solid rgba(215, 217, 224, 1) ;
            border-bottom: 1px solid rgba(215, 217, 224, 1) ;
            padding: 10px;
            width: 100%;
            background: #fff;
            box-sizing: border-box;
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
                    background: url("../../assets/left.png") no-repeat 0 1px;
                    margin-right: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_left_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    margin-right: 0.05rem;
                    cursor: pointer;
                    background: url("../../assets/left_disabled.png") no-repeat 0
                        1px;
                }
                .flag_item_right {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../../assets/right.png") no-repeat 0 0;
                    margin-left: 0.05rem;
                    cursor: pointer;
                }
                .flag_item_right_disabled {
                    display: inline-block;
                    width: 0.2rem;
                    height: 0.17rem;
                    background: url("../../assets/right_disabled.png") no-repeat 0
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
.filter_content{
    flex: 1;
    display: flex;
    .tab_option{
        font-size: 0.12rem;
        margin: 0 0 0.1rem;
        cursor: pointer;
        span{
            padding-right: 0.1rem;
        }
    }
    .tab_option:last-child{
        span:last-child{
            display: none;
        }
    }
    .blue_style{
        color: var(--bgColor);
    }
}
.voting_options {
    display: flex;
    color: var(--contentColor);
    margin-bottom: 10px;
    flex-wrap: wrap;
    align-items: center;
    .yes_option_style{
        display: inline-block;
        width: 0.12rem;
        height:0.12rem;
        border-radius: 0.02rem;
        background: var(--bgColor);
        margin-right: 0.1rem;
    }
    .no_option_style{
        display: inline-block;
        width: 0.12rem;
        height:0.12rem;
        border-radius: 0.02rem;
        background: #FFCF65;
        margin-right: 0.1rem;
    }
    .no_with_veto_option_style{
        display: inline-block;
        width: 0.12rem;
        height:0.12rem;
        border-radius: 0.02rem;
        background: #FE8A8A;
        margin-right: 0.1rem;
    }
    .abstain_option_style{
        display: inline-block;
        width: 0.12rem;
        height:0.12rem;
        border-radius: 0.02rem;
        background: #CCDCFF;
        margin-right: 0.1rem;
    }
    & > span {
        font-size: 0.14rem;
        color: var(--contentColor);
        @include fontWeight;
        padding: 0 0.18rem;
        white-space: nowrap;
        display: flex;
        align-items: center;
        line-height: 0.3rem;
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
    margin-top: 0.2rem;
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
    @media screen and (max-width: 910px){
        .proposals_detail_level{
            span{
                padding-left: 0.1rem !important;
            }
        }
        .information_props{
            color: #787c99 !important;
        }
        .mobile_transactions_detail_wrap{
            .proposal_detail_content{
                width: 100%;
                margin-top: 0.2rem;
                .proposals_detail_information_wrap:last-child{
                    margin-top: 0.2rem;
                }
            }
            .level_container{
                display: flex;
                margin: 0.1rem 0;
                .step_content{
                    margin-left: 0.2rem;
                    span{
                        i{
                            padding-right: 0.1rem;
                        }
                    }
                }
                .time_content{
                    margin-left: 0.2rem;
                    span{
                        i{
                            padding-right: 0.1rem;
                        }
                    }
                }
            }
        }
    }
</style>
