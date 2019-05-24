<template>
    <div class="block_detail_container">
        <div class="block_detail_content">
            <div class="block_detail_title_content">
                <span class="block_height_content">
                    <i :class="active?'flag_item_left':'flag_item_left_disabled'" @click="skipNext(-1)"></i>
                        <span class="information_value" style="flex:none;">{{`#${heightValue}`}}</span>
                    <i :class="activeNext?'flag_item_right':'flag_item_right_disabled'" @click="skipNext(1)"></i>
                </span>
            </div>
        </div>
        <div class="block_detail_information_container">
            <div class="block_information_title_content">
                <span>Block Information</span>
            </div>
            <div class="block_information_content">
                <div class="current_block_information_content">
                    <div class="block_information_item">
                        <span>Timestamp:</span>
                        <span v-if="timestampValue">{{ageValue}} ({{timestampValue}})</span>
                        <span v-if="!timestampValue">--</span>
                    </div>
                    <div class="block_information_item">
                        <span>Block Hash:</span>
                        <span>{{blockHashValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Transactions:</span>
                        <span>{{transactionsValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Proposer:</span>
                        <span v-if="proposerAddress !== '--'"><router-link class="common_link_style" :to="`/address/1/${proposerAddress}`">{{proposerValue}}</router-link></span>
                        <span v-if="proposerAddress === '--'">--</span>
                    </div>
                    <div class="block_information_item">
                        <span>Rewards:</span>
                        <span>{{rewardsValue}}</span>
                    </div>
                </div>
                <div class="last_block_information_content">
                    <div class="last_block_information_item">
                        <span>Last Block:</span>
                        <span v-if="lastBlockValue !== '--'"><router-link :to="`/block/${lastBlockValue}`" class="common_link_style">{{lastBlockValue}}</router-link></span>
                        <span v-if="lastBlockValue === '--'">--</span>
                    </div>
                    <div class="last_block_information_item">
                        <span>Last Block Hash:</span>
                        <span>{{lastBlockHashValue}}</span>
                    </div>
                    <div class="last_block_information_item">
                        <span>Precommit Validators:</span>
                        <span>{{precommitValidatorsValue}}</span>
                    </div>
                    <div class="last_block_information_item">
                        <span>Voting Power:</span>
                        <span>{{votingPowerValue}}</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="block_table_container">
            <div class="block_result_container">
                <div class="block_result_title">Block Result</div>
                <div class="block_result_table_content" style="overflow-x: auto;">
                    <blocks-list-table :items="items"
                                       :showNoData="flBlockResultNoData" :min-width="tableMinWidth"></blocks-list-table>
                    <div class="pagination" style='margin-top:0.2rem;' v-if="flShowTxListPagination">
                        <b-pagination size="md" :total-rows="txListCount" v-model="txListCurrentPage" :per-page="pageSize">
                        </b-pagination>
                    </div>
                    <div v-show="flBlockResultNoData" class="no_data_show">
                        No Data
                    </div>
                </div>

            </div>
            <div class="block_proposal_container">
                <div class="block_proposal_title">Proposal</div>
                <div class="block_proposal_content" style="overflow-x: auto;">
                    <blocks-list-table :items="governanceList"
                                       :showNoData="flGovernanceNoData" :min-width="tableMinWidth"></blocks-list-table>
                    <div class="pagination" style='margin-top:0.2rem;' v-if="flShowGovernanceListPagination">
                        <b-pagination size="md" :total-rows="governanceListCount" v-model="governanceListCurrentPage" :per-page="pageSize">
                        </b-pagination>
                    </div>
                    <div v-show="flGovernanceNoData" class="no_data_show">
                        No Data
                    </div>
                </div>

            </div>
            <div class="block_validator_set_container">
                <div class="block_validator_set_title">Validator Set</div>
                <div class="block_validator_set_content" style="overflow-x: auto;">
                    <blocks-list-table :items="validatorSetList"
                                       :showNoData="flValidatorNoData" :min-width="tableMinWidth"></blocks-list-table>
                    <div class="pagination" style='margin-top:0.2rem;margin-bottom: 0.2rem;' v-if="flShowValidatorListSetPagination">
                        <b-pagination size="md" :total-rows="validatorSetListCount" v-model="validatorSetListCurrentPage" :per-page="pageSize">
                        </b-pagination>
                    </div>
                    <div v-show="flValidatorNoData" class="no_data_show">
                        No Data
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<script>
    import Tools from '../util/Tools';
    import BlocksListTable from './table/BlockDetailListTable';
    import SpinComponent from './commonComponents/SpinComponent';
    import Service from "../util/axios";
    import Constant from "../constant/Constant"
    export default {
        components: {
            BlocksListTable,
            SpinComponent,
        },
        watch: {
            txListCurrentPage(txListCurrentPage){
                this.getTxList(txListCurrentPage,this.pageSize)
            },
            validatorSetListCurrentPage(validatorSetListCurrentPage){
                this.getValidatorSetList(validatorSetListCurrentPage,this.pageSize)
            },
            $route() {
                this.getBlockInformation();
                this.computeMinWidth();
                if (Number(this.$route.params.height) <= 1) {
                    this.active = false;
                } else {
                    this.active = true;
                }
                if (this.maxBlock !== 0) {
                    if (Number(this.$route.params.height) >= this.maxBlock) {
                        this.activeNext = false;
                    } else {
                        this.activeNext = true;
                    }
                }
            }
        },
        data() {
            return {
                transactionsDetailWrap: 'personal_computer_transactions_detail',
                heightValue: '',
                timestampValue: '',
                blockHashValue: '',
                transactionsValue: '',
                lastBlockHashValue: '',
                precommitValidatorsValue: '',
                votingPowerValue: '',
                items: [],
                governanceList: [],
                validatorSetList: [],
                showNoData: false,
                flBlockResultNoData: false,
                flGovernanceNoData: false,
                flValidatorNoData: false,
                active: true,
                activeNext: true,
                maxBlock: 0,
                showLoading:false,
                pageSize: 10,
                tableMinWidth:"",
                blockDetailTimer: null,
                ageValue: "",
                lastBlockValue:null,
                proposerValue: "",
                proposerAddress:'',
                rewardsValue: null,
                txListCount: 0,
                governanceListCount: 0,
                validatorSetListCount: 0,
                txListCurrentPage:1,
                validatorSetListCurrentPage:1,
                governanceListCurrentPage:1,
                blockListTxTimer: null,
                flShowTxListPagination: false,
                flShowGovernanceListPagination: false,
                flShowValidatorListSetPagination: false,
                flShowSampleIcon: false,
                sampleIconArray:[
                    {
                        color:'#3598DB',
                        name:'Yes'
                    },
                    {
                        color:'#73D1FF',
                        name:'Abstain'
                    },
                    {
                        color:'#F9CA18',
                        name:'No'
                    },
                    {
                        color:'#F27777',
                        name:'NoWithVeto'
                    }
                ]
            }
        },
        beforeMount() {
            Tools.scrollToTop();
            if (Tools.currentDeviceIsPersonComputer()) {
                this.transactionsDetailWrap = 'personal_computer_transactions_detail_wrap';
            } else {
                this.transactionsDetailWrap = 'mobile_transactions_detail_wrap';
            }
        },
        mounted() {
            this.getBlockInformation();
            if (Number(this.$route.params.height) <= 1) {
                this.active = false;
            } else {
                this.active = true;
            }
            this.computeMinWidth();
        },
        methods: {
            computeMinWidth(){
                if(this.$route.params.height){
                    this.tableMinWidth = 8.8;
                }
            },
            getBlockInformation() {
                let url = `/api/block/${this.$route.params.height}`;
                Service.http(url).then((data) => {
                    clearInterval(this.blockDetailTimer);
                    if (data) {
                        let that = this;
                        this.transactionsValue = data.block_info.transactions;
                        this.heightValue = data.block_info.block_height;
                        let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                        this.ageValue = Tools.formatAge(currentServerTime,data.block_info.timestamp,Constant.SUFFIX);
                        this.blockDetailTimer = setInterval(function () {
                            currentServerTime = new Date().getTime() + that.diffMilliseconds;
                            that.ageValue = Tools.formatAge(currentServerTime,data.block_info.timestamp,Constant.SUFFIX);
                        },1000);
                        this.timestampValue = Tools.format2UTC(data.block_info.timestamp);
                        this.blockHashValue = data.block_info.block_hash;
                        this.lastBlockValue = data.block_info.last_block !== 0 ? data.block_info.last_block : '--';
                        this.lastBlockHashValue = data.block_info.last_block_hash ? data.block_info.last_block_hash : '--';
                        this.proposerValue = data.block_info.propopser_moniker;
                        this.proposerAddress = data.block_info.propopser_addr;
                        this.rewardsValue = data.block_info.rewards ? `${Tools.convertScientificNotation2Number(Tools.formatNumber(data.block_info.rewards[0].amount))} ${Tools.formatDenom(data.block_info.rewards [0].denom)}` : '--';
                        this.precommitValidatorsValue = data.block_info.validator_num !== 0 ? data.block_info.validator_num : '--';
                        this.votingPowerValue = data.block_info.total_voting_power !== 0 ? data.block_info.total_voting_power : '--';
                        this.handleTxList(data.token_flows);
                        this.handleGovernance(data.proposals_info);
                        this.handleValidatorSetList(data.validator_set);
                        this.getMaxBlock(data.block_info.latest_height);
                    } else {
                        this.handleTxList(null);
                        this.handleGovernance(null);
                        this.handleValidatorSetList(null);
                        this.lastBlockValue = '--';
                        this.proposerAddress = '--';
                        this.rewardsValue = '--';
                        this.heightValue = '';
                        this.timestampValue = '';
                        this.blockHashValue = '--';
                        this.transactionsValue = '--';
                        this.lastBlockHashValue = '--';
                        this.precommitValidatorsValue = '--';
                        this.votingPowerValue = '--';
                    }
                }).catch(e => {
                    console.log(e)
                })
            },
            getValidatorSetList(currentPage,pageSize){
                let url = `/api/block/validatorset/${this.$route.params.height}&page=${currentPage}&size=${pageSize}`;
                Service.http(url).then((validatorSetList) => {
                    if(validatorSetList){
                        this.handleValidatorSetList(validatorSetList)
                    }
                })
            },
            getTxList(currentPage,pageSize){
                let url = `/api/block/coinflow/${this.$route.params.height}&page=${currentPage}&size=${pageSize}`;
                Service.http(url).then((txList) => {
                    if(txList){
                        this.handleTxList(txList)
                    }
                })
            },
            handleTxList(txList){
                if(txList && txList.items && txList.items.length !== 0){
                    this.txListCount  = txList.total;
                    if(txList.total > this.pageSize){
                        this.flShowTxListPagination = true
                    }else {
                        this.flShowTxListPagination = false
                    }
                    let that = this;
                    this.items = txList.items.map( item => {
                        let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                        return{
                            'TxHash' : item.tx_hash,
                            'From' : item.from,
                            'To' : item.to,
                            'Amount' : `${Tools.convertScientificNotation2Number(Tools.formatNumber(item.amount.amount))} ${Tools.formatDenom(item.amount.denom).toUpperCase()}`,
                            'Fee' : `${Tools.formatFeeToFixedNumber(item.fee.amount[0].amount)} ${Tools.formatDenom(item.fee.amount[0].denom).toUpperCase()}`,
                            'Tx_Initiator' : item.tx_initiator,
                            'Tx_Type' : item.tx_type,
                            'Tags' : item.flow_type,
                            'Status' : item.status,
                            'Timestamp' : Tools.formatAge(currentServerTime,item.timestamp,Constant.SUFFIX,Constant.PREFIX),
                            'txTime':item.timestamp,
                            'listName':"tx"
                        }
                    });
                    clearInterval(this.blockListTxTimer);
                    this.blockListTxTimer = setInterval(function () {
                        that.items = that.items.map( item => {
                            let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                            item.Timestamp = Tools.formatAge(currentServerTime,item.txTime,Constant.SUFFIX,Constant.PREFIX)
                            return item
                        })
                    },1000)
                }else {
                    this.flBlockResultNoData = true;
                    this.items =[{
                        'TxHash' : '',
                        'From' : '',
                        'To' : '',
                        'Amount' : '',
                        'Fee' : '',
                        'Tx_Initiator' : '',
                        'Tx_Type' : '',
                        'Tags' : '',
                        'Status' : '',
                        'Timestamp' : '',
                        'txTime': '',
                        'listName': ''
                    }]
                }
            },
            handleGovernance(proposals){
                if(proposals && proposals.length !== 0){
                    this.flShowSampleIcon = true;
                    this.governanceListCount = proposals.length;
                    if(proposals.length > this.pageSize){
                        this.flShowGovernanceListPagination = true
                    }else {
                        this.flShowGovernanceListPagination = false
                    }
                    this.governanceList = proposals.map(item => {
                        let voteTotal = item.result.Abstain + item.result.No + item.result.NoWithVeto + item.result.Yes;
                        return{
                            Radio: '',
                            abstainRadio: `${item.result.Abstain === 0 ? 0 :((item.result.Abstain / voteTotal) * 100 ).toFixed(0)}`,
                            noRadio: `${item.result.No === 0 ? 0 :((item.result.No / voteTotal) * 100 ).toFixed(0)}`,
                            noWithVetoRadio: `${item.result.NoWithVeto === 0 ? 0 :((item.result.NoWithVeto / voteTotal) * 100 ).toFixed(0)}`,
                            yesRadio: `${item.result.Yes === 0 ? 0 :((item.result.Yes / voteTotal) * 100 ).toFixed(0)}`,
                            Title: item.proposal.title,
                            'ProposalId': item.proposal.proposal_id,
                            'Type': item.proposal.type,
                            'Status': item.proposal.status,
                            'Proposer': item.proposal.proposer,
                            'Submit Time': Tools.format2UTC(item.proposal.submit_time),
                            'Voting Start Time': this.flShowProposalTime('votingStartTime',item.proposal.status) ? Tools.format2UTC(item.proposal.voting_start_time) : '--',
                            'Total Deposits': `${Tools.convertScientificNotation2Number(Tools.formatNumber(item.proposal.total_deposit[0].amount))} ${Tools.formatDenom(item.proposal.total_deposit[0].denom)}`,
                            'listName':"gov"
                        }
                    })
                }else {
                    this.flShowSampleIcon = false;
                    this.flGovernanceNoData = true;
                    this.governanceList = [{
                        Title: '',
                        Radio: '',
                        'Proposal ID': '',
                        'Type': '',
                        'Status': '',
                        'Proposer': '',
                        'Submit Time': '',
                        'Voting Start Time': '',
                        'Total Deposits': '',
                        'listName':"gov"
                    }]
                }
            },
            handleValidatorSetList(validatorList){
                if(validatorList && validatorList.items && validatorList.items.length !== 0){
                    this.validatorSetListCount = validatorList.total;
                    if(validatorList.total > this.pageSize){
                        this.flShowValidatorListSetPagination = true
                    }else {
                        this.flShowValidatorListSetPagination = false
                    }
                    this.validatorSetList = validatorList.items.map( validator => {
                        return{
                            'Moniker' : validator.moniker,
                            'OperatorAddress' : validator.operator_address,
                            'Consensus': validator.consensus,
                            'ProposerPriority': validator.proposer_priority,
                            'VotingPower' : validator.voting_power,
                        }
                    })
                }else {
                    this.flValidatorNoData = true;
                    this.validatorSetList = [
                        {
                            'Moniker' : '',
                            'OperatorAddress' : '',
                            'Consensus': '',
                            'ProposerPriority': '',
                            'VotingPower' : '',
                        }
                    ]
                }
            },
            skipNext(num) {
                if (Number(this.$route.params.height) <= 1) {
                    this.active = false;
                    if (num !== -1) {
                        this.$router.push(`/block/${Number(this.$route.params.height) + num}`)
                    }
                } else if (Number(this.$route.params.height) >= this.maxBlock) {
                    if (num !== 1) {
                        this.$router.push(`/block/${Number(this.$route.params.height) + num}`)
                    }
                } else {
                    this.active = true;
                    this.$router.push(`/block/${Number(this.$route.params.height) + num}`)
                }
            },
            getMaxBlock(latestHeight) {
                this.maxBlock = latestHeight;
                if (Number(this.$route.params.height) >= Number(latestHeight)) {
                    this.activeNext = false;
                } else {
                    this.activeNext = true;
                }
            },
            flShowProposalTime(proposalTimeName,status){
                if(status === 'Rejected' || status === 'Passed' || status === 'VotingPeriod'){
                    return true
                }else{
                    switch (proposalTimeName){
                        case proposalTimeName === 'depositEndTime' && status === 'DepositPeriod' : return true ;
                        case proposalTimeName === 'votingStartTime' && status === 'VotingPeriod' : return true ;
                        case proposalTimeName === 'votingEndTime' && status === 'VotingPeriod' : return true ;
                    }
                }
            },
        }
    }
</script>

<style scoped lang="scss">
    .block_detail_container{
        .block_detail_content{
            background: rgba(239, 239, 241, 1);
            .block_detail_title_content{
                height: 0.62rem;
                display: flex;
                align-items: center;
                max-width: 12.8rem;
                margin: 0 auto;
                padding-left: 0.2rem;
                .block_height_content{
                    padding-left: 0.2rem;
                    .information_value{
                        font-size: 0.22rem;
                        color: #22252a;
                        margin: 0 0.07rem;
                    }
                    .flag_item_left {
                        display: inline-block;
                        width: 0.2rem;
                        height: 0.17rem;
                        background: url('../assets/left.png') no-repeat 0 1px;
                        margin-right: 0.05rem;
                        cursor: pointer;
                    }
                    .flag_item_left_disabled {
                        display: inline-block;
                        width: 0.2rem;
                        height: 0.17rem;
                        margin-right: 0.05rem;
                        cursor: pointer;
                        background: url('../assets/left_disabled.png') no-repeat 0 1px;
                    }
                    .flag_item_right {
                        display: inline-block;
                        width: 0.2rem;
                        height: 0.17rem;
                        background: url('../assets/right.png') no-repeat 0 0;
                        margin-left: 0.05rem;
                        cursor: pointer;
                    }
                    .flag_item_right_disabled {
                        display: inline-block;
                        width: 0.2rem;
                        height: 0.17rem;
                        background: url('../assets/right_disabled.png') no-repeat 0 0;
                        margin-left: 0.05rem;
                        cursor: pointer;
                    }
                }
            }
        }
        .block_detail_information_container{
            max-width: 12.8rem;
            margin: 0 auto;
            border-radius: 0.01rem;
            .block_information_title_content{
                height:0.69rem;
                font-size: 0.18rem;
                display: flex;
                align-items: center;
                padding-left: 0.2rem;
            }
            .block_information_content{
                border: 0.01rem solid rgba(215, 217, 224, 1);
                border-radius: 0.01rem;
                .current_block_information_content{
                    box-sizing: border-box;
                    padding: 0.2rem 0.2rem 0 0.2rem;
                    .block_information_item{
                        display: flex;
                        span:nth-of-type(1){
                            color: rgba(34, 37, 42, 1);
                            font-size: 0.14rem;
                            line-height: 0.2rem;
                            width: 1.5rem;
                            display: inline-block;
                            margin-bottom: 0.12rem;
                        }
                        span:nth-of-type(2){
                            flex: 1;
                            color: rgba(162, 162, 174, 1);
                            font-size: 0.14rem;
                            line-height: 0.2rem;
                            overflow-x: auto;
                            .common_link_style{
                                color: rgba(53, 152, 219, 1) !important;
                            }
                        }
                    }
                }
                .last_block_information_content{
                    box-sizing: border-box;
                    background: rgba(243, 251, 255, 1);
                    padding: 0.16rem 0.2rem 0.2rem 0.2rem ;
                    .last_block_information_item{
                        margin-bottom: 0.12rem;
                        display: flex;
                        span:nth-of-type(1){
                            font-size: 0.14rem;
                            color: rgba(34, 37, 42, 1);
                            line-height: 0.2rem;
                            width: 1.5rem;
                            display: inline-block;
                        }
                        span:nth-of-type(2){
                            flex: 1;
                            color: rgba(162, 162, 174, 1);
                            font-size: 0.14rem;
                            line-height: 0.2rem;
                            overflow-x: auto;
                            .common_link_style{
                                color: rgba(53, 152, 219, 1) !important;
                            }
                        }
                    }
                    .last_block_information_item:last-child{
                        margin-bottom: 0;
                    }
                }
            }
        }
        .block_table_container{
            max-width: 12.8rem;
            margin: 0 auto;
            .block_result_container{
                .block_result_title{
                    height: 0.65rem;
                    display: flex;
                    align-items: center;
                    padding-left: 0.2rem;
                }
                .pagination{
                    display: flex;
                    justify-content: flex-end;
                }

            }
            .block_proposal_container{
                .block_proposal_title{
                    height: 0.65rem;
                    display: flex;
                    align-items: center;
                    padding-left: 0.2rem;
                }
                .pagination{
                    display: flex;
                    justify-content: flex-end;
                }
            }
            .block_validator_set_container{
                margin-bottom: 0.2rem;
                padding-bottom: 0.2rem;
                .block_validator_set_title{
                    height: 0.65rem;
                    display: flex;
                    align-items: center;
                    padding-left: 0.2rem;
                }
                .pagination{
                    display: flex;
                    justify-content: flex-end;
                }
            }
        }
    }
    .no_data_show{
        display: flex;
        justify-content: center;
        border-top:0.01rem solid #eee;
        border-bottom:0.01rem solid #eee;
        font-size:0.14rem;
        height:1rem;
        align-items: center;
    }
    @media screen and (max-width:910px) {
        .block_table_container{
            padding: 0 0.1rem;
            .block_detail_information_container{
                .block_information_content{
                    .current_block_information_content{
                        .block_information_item{
                            flex-direction: column;
                            span:nth-of-type(1){
                                margin-bottom: 0 !important;
                            }
                        }
                    }
                    .last_block_information_content{
                        .last_block_information_item{
                            flex-direction: column;
                            margin-bottom: 0 !important;
                            span:nth-of-type(1){
                                margin-bottom: 0 !important;
                            }
                        }
                    }
                }
            }

        }
    }
</style>
