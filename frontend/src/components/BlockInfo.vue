<template>
    <div class="block_detail_container">
        <div class="block_detail_content">
            <div class="block_detail_title_content">
                <span class="block_height_content">
                    <i :class="active?'flag_item_left':'flag_item_left_disabled'" @click="skipNext(-1)"></i>
                        <span class="information_value" style="flex:none;">{{heightValue}}</span>
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
                        <span>Block Hash:</span>
                        <span>{{blockHashValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Proposer:</span>
                        <span v-if="proposerAddress !== '--'"><router-link class="common_link_style" :to="`/address/1/${proposerAddress}`">{{proposerValue}}</router-link></span>
                        <span v-if="proposerAddress === '--'">--</span>
                    </div>
                    <div class="block_information_item">
                        <span>Validators:</span>
                        <span>{{validatorValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Voting Power:</span>
                        <span>{{votingPowerValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Transactions:</span>
                        <span>{{transactionsValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Inflation:</span>
                        <span>{{inflationValue}}</span>
                    </div>
                    <div class="block_information_item">
                        <span>Timestamp:</span>
                        <span v-if="timestampValue">{{timestampValue}}</span>
                        <span v-if="!timestampValue">--</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="block_table_container">
            <div class="block_result_container">
                <div class="block_result_title">Transactions</div>
                <div class="block_result_table_content" style="overflow-x: auto;">
                    <blocks-list-table :items="items"
                                       :showNoData="flBlockResultNoData" :min-width="tableMinWidth"></blocks-list-table>
                    <div v-show="flBlockResultNoData" class="no_data_show">
                        No Data
                    </div>
                    <div class="pagination" style='margin-top:0.2rem;' v-if="flShowTxListPagination">
                        <b-pagination size="md" :total-rows="txListCount" v-model="txListCurrentPage" :per-page="pageSize">
                        </b-pagination>
                    </div>
                </div>
            </div>
            <div class="block_proposal_container">
                <div class="block_proposal_title">Governance</div>
                <div class="block_proposal_content" style="overflow-x: auto;">
                    <blocks-list-table :items="governanceList"
                                       :showNoData="flGovernanceNoData" :min-width="tableMinWidth"></blocks-list-table>
                    <div v-show="flGovernanceNoData" class="no_data_show">
                        No Data
                    </div>
                    <div class="pagination" style='margin-top:0.2rem;' v-if="flShowGovernanceListPagination">
                        <b-pagination size="md" :total-rows="governanceListCount" v-model="governanceListCurrentPage" :per-page="pageSize">
                        </b-pagination>
                    </div>
                </div>
            </div>
            <div class="block_validator_set_container">
                <div class="block_validator_set_title">Validator Set</div>
                <div class="block_validator_set_content" style="overflow-x: auto;">
                    <blocks-list-table :items="validatorSetList"
                                       :showNoData="flValidatorNoData" :min-width="tableMinWidth"></blocks-list-table>
                    <div v-show="flValidatorNoData" class="no_data_show">
                        No Data
                    </div>
                    <div class="pagination" style='margin-top:0.2rem;margin-bottom: 0.2rem;' v-if="flShowValidatorListSetPagination">
                        <b-pagination size="md" :total-rows="validatorSetListCount" v-model="validatorSetListCurrentPage" :per-page="pageSize">
                        </b-pagination>
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
	        governanceListCurrentPage(governanceListCurrentPage){
		        this.getGovList(governanceListCurrentPage,this.pageSize)
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
                proposerValue: "",
                proposerAddress:'',
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
                validatorValue: null,
	            inflationValue: null,
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
            this.getTxList(this.txListCurrentPage,this.pageSize);
	        this.getValidatorSetList(this.validatorSetListCurrentPage,this.pageSize);
	        this.getTxList(this.txListCurrentPage,this.pageSize);
            this.getGovList(this.governanceListCurrentPage,this.pageSize);
	        if (Number(this.$route.params.height) > 1) {
                this.active = true;
		        this.activeNext = true;
            } else {
                this.active = false;
                this.activeNext = false
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
                let url = `/api/block/blockinfo/${this.$route.params.height}`;
                Service.http(url).then((result) => {
                    if (result) {
                        this.transactionsValue = result.transactions;
                        this.heightValue = result.block_height;
                        this.validatorValue = `${result.precommit_validator_num !== null ? result.precommit_validator_num : '--'} / ${result.total_validator_num ? result.total_validator_num : '--'}`;
                        this.votingPowerValue = result.precommit_voting_power !== null ? `${((result.precommit_voting_power / result.total_voting_power) *100).toFixed(4)} %` : '--';
                        this.timestampValue = Tools.format2UTC(result.timestamp);
                        this.blockHashValue = result.block_hash;
                        this.proposerValue = result.propopser_moniker;
                        this.proposerAddress = result.propopser_addr;
                        this.inflationValue = result.mint_coin.denom !== '' ? `${Tools.convertScientificNotation2Number(Tools.formatNumber(result.mint_coin.amount))} ${Tools.formatDenom(result.mint_coin.denom)}` : '--';
                        this.precommitValidatorsValue = result.validator_num !== 0 ? result.validator_num : '--';
                        this.getMaxBlock(result.latest_height)
                    } else {
	                    this.validatorValue= '--'
                        this.proposerAddress = '--';
                        this.inflationValue = '--';
                        this.heightValue = '';
                        this.timestampValue = '';
                        this.blockHashValue = '--';
                        this.transactionsValue = '--';
                        this.precommitValidatorsValue = '--';
                        this.votingPowerValue = '--';
                    }
                }).catch(err => {
                    console.error(err)
	                this.validatorValue= '--'
	                this.proposerAddress = '--';
	                this.inflationValue = '--';
	                this.heightValue = '';
	                this.timestampValue = '';
	                this.blockHashValue = '--';
	                this.transactionsValue = '--';
	                this.precommitValidatorsValue = '--';
	                this.votingPowerValue = '--';
                })
            },
            getGovList(currentPage,pageSize){
	            let url = `/api/block/txgov/${this.$route.params.height}?page=${currentPage}&size=${pageSize}`;
	            Service.http(url).then((proposal) => {
                    this.handleGovernance(proposal)
	            }).catch(err => {
	            	console.error(err);
		            this.handleGovernance(null)
                })
            },
            getValidatorSetList(currentPage,pageSize){
                let url = `/api/block/validatorset/${this.$route.params.height}?page=${currentPage}&size=${pageSize}`;
                Service.http(url).then((validatorSetList) => {
                    this.handleValidatorSetList(validatorSetList)
                }).catch(err => {
	                console.error(err);
	                this.handleValidatorSetList(null)
                })
            },
            getTxList(currentPage,pageSize){
                let url = `/api/block/txs/${this.$route.params.height}?page=${currentPage}&size=${pageSize}`;
                Service.http(url).then((txList) => {
                    this.handleTxList(txList)
                }).catch(err => {
	                console.error(err);
	                this.handleTxList(null)
                })
            },
            handleTxList(txList){
                if(txList && txList.items && txList.items.length !== 0){
	                this.flBlockResultNoData = false;
                    this.txListCount  = txList.total;
                    if(txList.total > this.pageSize){
                        this.flShowTxListPagination = true
                    }else {
                        this.flShowTxListPagination = false
                    }
                    this.items = txList.items.map( item => {
	                    return{
                            'TxHash' : item.hash,
                            'From' : item.from,
                            'To' : item.to,
                            'Amount' : `${Tools.convertScientificNotation2Number(Tools.formatNumber(item.amount[0].amount))} ${Tools.formatDenom(item.amount[0].denom).toUpperCase()}`,
                            'Fee' : `${Tools.formatFeeToFixedNumber(item.actual_fee.amount)} ${Tools.formatDenom(item.actual_fee.denom).toUpperCase()}`,
                            'Tx_Initiator' : item.tx_initiator,
                            'Tx_Type' : item.type,
                            'Status' : item.status,
                            'listName':'tx'
                        }
                    });
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
                        'Status' : '',
	                    'listName':'tx'
                    }]
                }
            },
            handleGovernance(proposals){
	            if(proposals && proposals .items && proposals.items.length !== 0){
	                this.flGovernanceNoData = false;
                    this.governanceListCount = proposals.total;
                    if(proposals.total > this.pageSize){
                        this.flShowGovernanceListPagination = true
                    }else {
                        this.flShowGovernanceListPagination = false
                    }
                    this.governanceList = proposals.items.map(item => {
                        return{
                        	"TxHash": item.hash,
                            "TxFee": `${Tools.formatFeeToFixedNumber(item.actual_fee.amount)} ${Tools.formatDenom(item.actual_fee.denom).toUpperCase()}`,
                            "TxSigner": item.tx_initiator,
                            "TxType": item.tx_type,
                            "TxStatus": item.status,
                            "ProposalType": item.proposal_type,
                            "ProposalId": item.proposal_id,
                            "ProposalTitle": item.proposal_title,
                            "listName":"gov"
                        }
                    })
                }else {
                    this.flGovernanceNoData = true;
                    this.governanceList = [{
	                    "TxHash": '',
	                    "TxFee": '',
	                    "TxSigner": '',
	                    "TxType": '',
	                    "TxStatus": '',
	                    "ProposalType": '',
	                    "ProposalId": '',
	                    "ProposalTitle": '',
	                    "listName":"gov"
                    }]
                }
            },
            handleValidatorSetList(validatorList){
                if(validatorList && validatorList.items && validatorList.items.length !== 0){
	                this.flValidatorNoData = false;
                    this.validatorSetListCount = validatorList.total;
                    if(validatorList.total > this.pageSize){
                        this.flShowValidatorListSetPagination = true
                    }else {
                        this.flShowValidatorListSetPagination = false
                    }
                    this.validatorSetList = validatorList.items.map( validator => {
                        return{
                            'Moniker' : Tools.formatString(validator.moniker,15,'...'),
                            'OperatorAddress' : validator.operator_address,
                            'Consensus': validator.consensus,
                            'ProposerPriority': validator.proposer_priority,
                            'VotingPower' : validator.voting_power,
                            'flProposer' : validator.is_proposer
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
	                        'flProposer': ''
                        }
                    ]
                }
            },
            skipNext(num) {
            	if(Number(this.$route.params.height) >= 1){
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
                    padding: 0.2rem;
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
                    .block_information_item:last-child{
                        span:nth-of-type(1){
                            color: rgba(34, 37, 42, 1);
                            font-size: 0.14rem;
                            line-height: 0.2rem;
                            width: 1.5rem;
                            display: inline-block;
                            margin-bottom:0;
                        }
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
        border-bottom:0.01rem solid #eee;
        font-size:0.14rem;
        min-width: 12rem;
        padding-top: 0.1rem;
        height:1rem;
        align-items: center;
    }
    @media screen and (max-width:960px) {
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
            }
        }
    }
</style>
