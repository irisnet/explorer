<template>
    <div class="transactions_detail_wrap">
        <!-- <div class="transactions_title_wrap">
            <p :class="transactionsDetailWrap" style="margin-bottom:0;">
                <span class="transactions_detail_title">Transaction</span>
                <span class="transactions_detail_wrap_hash_var">{{hashValue}}</span>
            </p>
        </div> -->
        <div :class="transactionsDetailWrap">
            <p class="transaction_information_content_title" style="height: 0.7rem; line-height: 0.7rem;">Transaction Information</p>
            <div class="transactions_detail_information_wrap" ref="valueInformation">
                <div class="information_props_wrap">
                    <span class="information_props">TxHash :</span>
                    <span class="information_value">{{hashValue}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Block Height :</span>
                    <span class="information_value link_active_style">
            <router-link :to="`/block/${blockValue}`">{{blockValue}}</router-link>
          </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Type :</span>
                    <span class="information_value">{{typeValue}}</span>
                </div>
                <div class="information_props_wrap" v-if="flShowProposer">
                    <span class="information_props">Proposer :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(proposer)">{{proposer}}</router-link></span>
                </div>
                <div class="information_props_wrap" v-if="title">
                    <span class="information_props">ProposalTitle :</span>
                    <span class="information_value">{{title}}</span>
                </div>
                <div class="information_props_wrap" v-if="proposalType">
                    <span class="information_props">ProposalType :</span>
                    <span class="information_value">{{proposalType}}</span>
                </div>
                <div class="information_props_wrap" v-if="flShowInitialDeposit">
                    <span class="information_props">InitialDeposit :</span>
                    <span class="information_value">{{initialDeposit}}</span>
                </div>
                <div class="information_props_wrap" v-if="description">
                    <span class="information_props">Description :</span>
                    <span class="information_value">{{description}}</span>
                </div>
                <div class="information_props_wrap" v-if="depositer">
                    <span class="information_props">Depositer :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(depositer)">{{depositer}}</router-link>
          </span>
                </div>
                <div class="information_props_wrap" v-if="flShowProposalId">
                    <span class="information_props">Proposal ID :</span>
                    <span v-show="proposalId !== '--' " class="information_value link_active_style">
            <router-link :to="`/ProposalsDetail/${proposalId}`">{{proposalId}}</router-link>
          </span>
                    <span v-show="proposalId === '--' " class="information_value link_active_style">{{proposalId}}</span>
                </div>
                <div class="information_props_wrap" v-if="flShowVoter">
                    <span class="information_props">Voter :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(voter)">{{voter}}</router-link></span>
                </div>
                <div class="information_props_wrap" v-if="flShowTypeTransfer || flShowWithdrawAddress || flShowTypeBurn">
                    <span class="information_props">From :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(fromValue)">{{fromValue}}</router-link>
          </span>
                </div>
                <div class="information_props_wrap" v-if="flShowWithdrawAddress">
                    <span class="information_props">Withdraw To :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(withdrawAddress)">{{withdrawAddress}}</router-link></span>
                </div>
                <div class="information_props_wrap" v-if="flShowDelegatorAddress">
                    <span class="information_props">Delegator Address :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(delegatorAddress)">{{delegatorAddress}}</router-link></span>
                </div>
                <div class="information_props_wrap" v-if="flShowValidatorAddress">
                    <span class="information_props">Validator Address :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(validatorAddress)">{{validatorAddress}}</router-link>
          </span>
                </div>
                <div class="information_props_wrap" v-if="showSource">
                    <span class="information_props">Source :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(source)">{{source}}</router-link>
          </span>
                </div>
                <div class="information_props_wrap" v-if="flShowTypeTransfer">
                    <span class="information_props">To :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(toValue)">{{toValue}}</router-link></span>
                </div>
                <div class="information_props_wrap" v-if="moniker">
                    <span class="information_props">Moniker :</span>
                    <span class="information_value"><pre class="information_pre">{{moniker}}</pre></span>
                </div>
                <div class="information_props_wrap" v-if="identity">
                    <span class="information_props">Identity :</span>
                    <span class="information_value"><pre class="information_pre">{{identity}}</pre></span>
                </div>
                <div class="information_props_wrap" v-if="owner">
                    <span class="information_props">From :</span>
                    <span class="information_value link_active_style">
            <router-link :to="addressRoute(owner)">{{owner}}</router-link></span>
                </div>
                <div class="information_props_wrap" v-if="pubkey">
                    <span class="information_props">Pub key :</span>
                    <span class="information_value">{{pubkey}}</span>
                </div>
                <div class="information_props_wrap" v-if="website">
                    <span class="information_props">Website :</span>
                    <span class="information_value"><pre class="information_pre">{{website}}</pre></span>
                </div>
                <div class="information_props_wrap" v-if="selfBond">
                    <span class="information_props">Self-Bonded :</span>
                    <span class="information_value">{{selfBond}}</span>
                </div>
                <div class="information_props_wrap" v-if="details">
                    <span class="information_props">Details :</span>
                    <span class="information_value"><pre class="information_pre">{{details}}</pre></span>
                </div>
                <div class="information_props_wrap" v-if="flShowVoter">
                    <span class="information_props">Option :</span>
                    <span class="information_value">{{option}}</span>
                </div>
                <div class="information_props_wrap" v-if="flShowTypeTransfer || flShowTypeDeposit || flShowTypeBurn">
                    <span class="information_props">Amount :</span>
                    <span class="information_value">{{amountValue}}</span>
                </div>
                <div class="information_props_wrap" v-if="flShowReceivedRewardsValue">
                    <span class="information_props">Received Rewards :</span>
                    <span class="information_value">{{amountValue}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Status :</span>
                    <span class="information_value information_value_fixed">
                        <span :class="{'fail_status': status === 'Fail' }">{{status}}</span>
                        <div class="info_icon_div question_icon_div" v-if="status === 'Fail' && failInfo" v-table-tooltip="{show: true, container: $refs.valueInformation}">
                            <div class="tooltip_span">
                                <div>{{failInfo}}</div>
                                <i></i>
                            </div>
                        </div>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Age(Timestamp) :</span>
                    <span class="information_value" v-show="ageValue">{{ageValue}} ({{timestampValue}})</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Actual Tx Fee :</span>
                    <span class="information_value">{{actualTxFee}}</span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Gas Used :</span>
                    <span class="information_value information_value_fixed">
                        <span>{{gasUsedByTxn}}</span>
                        <div class="info_icon_div" v-if="gasPrice || gasUsedByTxn || gasWanted || gasLimit" v-table-tooltip="{show: true, container: $refs.valueInformation}">
                            <div class="tooltip_span">
                                <div>
                                    <p>Gas Used : {{gasUsedByTxn || '--'}}</p>
                                    <p>Gas Wanted : {{gasWanted || '--'}}</p>
                                    <p>Gas Limit : {{gasLimit || '--'}}</p>
                                </div>
                                <i></i>
                            </div>
                        </div>
                    </span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Gas Price :</span>
                    <span class="information_value">{{gasPrice}} <span v-if="gasPrice">Nano</span></span>
                </div>
                <div class="information_props_wrap">
                    <span class="information_props">Memo :</span>
                    <span class="information_value"><pre class="information_pre">{{memo}}</pre></span>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
	import Tools from '../util/Tools';
	import Service from "../service";
	import Constant from "../constant/Constant"
	export default {
		data() {
			return {
				devicesWidth: window.innerWidth,
				transactionsDetailWrap: 'personal_computer_transactions_detail',
				hashValue: '',
				blockValue: '',
				typeValue: '',
				fromValue: '',
				toValue: '',
				timestampValue: '',
				amountValue: '',
				actualTxFee: '',
				gasLimit:'',
				gasUsedByTxn:'',
                gasPrice:'',
                gasWanted: '',
				memo: "",
				owner: "",
				moniker: "",
				selfBond: "",
				pubkey: "",
				identity: "",
				website: "",
				details: "",
				source: "",
				showSource: "",
				proposer: "",
				initialDeposit: "",
				title: "",
				description: "",
				proposalId: "",
				proposalType:"",
				depositer: "",
				voter: "",
				option: "",
                status: "",
                failInfo: "",
				withdrawAddress: "",
				delegatorAddress: "",
				validatorAddress: "",
				flShowProposalId: false,
				flShowProposer: false,
				flShowInitialDeposit: false,
				flShowTypeTransfer:false,
				flShowVoter: false,
				flShowTypeDeposit: false,
				flShowWithdrawAddress: false,
				flShowDelegatorAddress: false,
				flShowValidatorAddress: false,
				flShowReceivedRewardsValue: false,
				flShowTypeBurn: false,
				ageValue: '',
				transactionDetailTimer: null,
			}
		},
		watch:{
			$route(){
				this.getTransactionInfo();
				Tools.scrollToTop();
            },
            "$store.state.isMobile"(newVal) {
                this.isMobileFunc(newVal);
            }
		},
		beforeMount() {
            this.isMobileFunc(this.$store.state.isMobile);
		},
		mounted() {
            this.getTransactionInfo();
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
			getTransactionInfo(){
				if(this.$route.query.txHash){
					Service.commonInterface({txDetail:{txHash:this.$route.query.txHash}}, (data) => {
						try {
							clearInterval(this.transactionDetailTimer);
							if(data){
								let that = this;
								let currentServerTime = new Date().getTime() + that.diffMilliseconds;
								this.transactionDetailTimer = setInterval(function () {
									currentServerTime = new Date().getTime() + that.diffMilliseconds;
									that.ageValue = Tools.formatAge(currentServerTime,data.Timestamp,Constant.SUFFIX);
								},1000);
								this.ageValue = Tools.formatAge(currentServerTime,data.Timestamp,Constant.SUFFIX);
								this.timestampValue = Tools.format2UTC(data.Timestamp);
								this.hashValue = data.Hash;
								this.blockValue = data.BlockHeight;
								this.typeValue = data.Type === 'coin'?'transfer':data.Type;
								this.gasPrice = Tools.convertScientificNotation2Number(Tools.formaNumberAboutGasPrice(data.GasPrice));
								this.gasLimit = data.GasLimit;
                                this.gasUsedByTxn = data.GasUsed;
                                this.gasWanted = data.GasWanted;
								this.memo = data.Memo ? data.Memo : '--';
                                this.status = data.Status ? Tools.firstWordUpperCase(data.Status): '--';
                                this.failInfo = data.Log;
								if(data.Amount && data.Amount.length !==0){
									this.amountValue = data.Amount.map(item=>{
										if(item.denom === Constant.Denom.IRISATTO){
											return item.amount = `${Tools.formatPriceToFixed(Tools.numberMoveDecimal(item.amount))} ${Tools.formatDenom(item.denom).toLocaleUpperCase()}`;
										}else if (item.denom !== Constant.Denom.IRISATTO && item.denom !== ''){
											return item.amount = `${Tools.FormatScientificNotationToNumber(item.amount)} ${Tools.formatDenom(item.denom).toUpperCase()}`
                                        }else {
											if(data.Type === "BeginUnbonding" || data.Type === "BeginRedelegate"){
												return item.amount = `${Tools.formatPriceToFixed(Tools.numberMoveDecimal(item.amount))} SHARES`;
                                            }
                                        }
									}).join(',') ;
								}else if(data.Amount && Object.keys(data.Amount).includes('amount') && Object.keys(data.Amount).includes('denom')){
									data.Amount =  Tools.formatPriceToFixed(Tools.numberMoveDecimal(item.amount));
									this.amountValue = `${data.Amount.amount} ${Tools.formatDenom(data.Amount.denom).toUpperCase()}`
								}else {
									this.amountValue = "--"
								}
								this.actualTxFee = `${Tools.convertScientificNotation2Number(Tools.formatNumber(data.Fee.amount))} ${Tools.formatDenom(data.Fee.denom).toUpperCase()}`;
								if(data.Type === "Transfer" || data.Type === "Delegate" || data.Type === "BeginUnbonding"){
									this.flShowTypeTransfer = true;
									this.fromValue = data.From;
									this.toValue = data.To;
								}else if(data.Type === "CreateValidator" || data.Type === "EditValidator" || data.Type === "Unjail"){
									this.owner = data.Owner ? data.Owner : '--';
									this.moniker = data.Moniker ? data.Moniker : '--';
									this.pubkey = data.Pubkey ? data.Pubkey : "--";
									this.identity = data.Identity ? data.Identity : '--';
									this.website = data.Website ? data.Website : '--';
									this.details = data.Details ? data.Details : '--';
									if(data.SelfBond && data.SelfBond.length !== 0){
										this.selfBond = `${Tools.formatPriceToFixed(Tools.convertScientificNotation2Number(Tools.formatNumber(data.SelfBond[0].amount)))} ${Tools.formatDenom(data.SelfBond[0].denom).toUpperCase()}`;
									}else {
										this.selfBond = "--"
									}
								}else if(data.Type === "BeginRedelegate"){
									this.flShowTypeTransfer = true;
									this.showSource = true;
									this.fromValue = data.From ? data.From : '';
									this.toValue = data.To ? data.To : "";
									this.source = data.Source ? data.Source : "";
								}else if(data.Type === "SubmitProposal"){
									this.flShowProposer = true;
									this.flShowInitialDeposit = true;
									this.title = data.Title ? data.Title : '--';
									this.proposer = data.From;
									this.proposalType = data.ProposalType;
									if(data.Amount && data.Amount.length !==0){
										this.initialDeposit = data.Amount.map(item=>{
											return `${item.amount} ${Tools.formatDenom(item.denom).toUpperCase()}`;
										}).join(',') ;
									}else {
										this.initialDeposit = "--"
									}
									this.description = data.Description ? data.Description : '--';
								}else if(data.Type === "Deposit"){
									this.flShowProposalId = true;
									this.flShowTypeDeposit = true;
									this.proposalId = data.ProposalId === 0 ? "--" : data.ProposalId;
									this.depositer = data.From ? data.From : "--";
								}else if(data.Type === "Vote"){
									this.flShowProposalId = true;
									this.flShowVoter = true;
									this.proposalId = data.ProposalId === 0 ? "--" : data.ProposalId;
									this.voter = data.From ? data.From : '--';
									this.option = data.Option ? data.Option : "--";
								}else if(data.Type === "SetWithdrawAddress"){
									this.flShowWithdrawAddress = true;
									this.fromValue = data.From ? data.From : '';
									this.withdrawAddress = data.To ? data.To : '';
								}else if(data.Type === "WithdrawDelegatorRewardsAll"){
									this.flShowReceivedRewardsValue = true;
									this.flShowDelegatorAddress = true;
									this.delegatorAddress = data.From ? data.From : '';
								}else if(data.Type === "WithdrawDelegatorReward"){
									this.flShowReceivedRewardsValue = true;
									this.flShowDelegatorAddress = true;
									this.flShowValidatorAddress = true;
									this.delegatorAddress = data.From ? data.From : '';
									this.validatorAddress = data.To ? data.To : "";
								} else if(data.Type === "WithdrawValidatorRewardsAll"){
									this.flShowReceivedRewardsValue = true;
									this.flShowValidatorAddress = true;
									this.validatorAddress = data.From ? data.From : "";
								} else if(data.Type === 'Burn'){
									this.flShowTypeBurn = true;
									this.fromValue = data.From;
								}
							}
						}catch (e) {
                            console.error(e)
						}
                    })
				}
			}
		}
	}
</script>
<style lang="scss" scoped>
    @import '../style/mixin.scss';
    .information_pre{
        color: #a2a2ae;
        white-space: pre-wrap;
    }
    .transactions_detail_wrap {
        @include flex;
        @include pcContainer;
        font-size:0.14rem;
        .transactions_title_wrap {
            width: 100%;
            border-bottom: 1px solid #d6d9e0 !important;
            height:0.62rem;
            background:#efeff1;
            line-height:0.62rem;
            @include flex;
            @include pcContainer;
            .personal_computer_transactions_detail_wrap {
                height:0.62rem;
                @include flex;
                span{
                    line-height:0.62rem;
                    height:0.62rem;
                }
            }
            .mobile_transactions_detail_wrap {
                @include flex;
                flex-direction: column;
            }

        }
        .personal_computer_transactions_detail_wrap {
            .transaction_information_content_title{
                height:0.5rem;
                line-height:0.5rem;
                font-size:0.18rem;
                color:#000000;
                @include fontWeight;
                margin-left: 0.2rem;
                margin-bottom:0;
                font-family:ArialMT;
            }
            @include pcCenter;
            .transactions_detail_information_wrap{
                padding: 0.2rem 0.2rem 0.08rem;
                border: 1px solid #d7d9e0;
                margin-bottom: 0.4rem;
                .information_props_wrap{
                    @include flex;
                    line-height: 0.2rem;
                    margin-bottom: 0.12rem;
                    .information_props{
                        width:1.5rem;
                    }
                    .information_value{
                        color: #a2a2ae;
                        flex:1;
                    }
                }
            }
            .transactions_detail_title {
                height: 0.4rem;
                line-height: 0.4rem;
                font-size: 0.22rem;
                color: #000000;
                margin-right: 0.2rem;
                @include fontWeight;
            }
            .transactions_detail_wrap_hash_var {
                height: 0.4rem;
                line-height: 0.4rem;
                font-size: 0.22rem;
                color: #a2a2ae;
            }
        }

        .mobile_transactions_detail_wrap {
            width: 100%;
            @include flex;
            flex-direction: column;
            padding: 0 0.1rem;
            .transaction_information_content_title{
                height: 0.5rem;
                line-height: 0.5rem;
                font-size: 0.18rem;
                color: #000000;
                margin-left: 0.1rem;
                margin-bottom:0;
                @include fontWeight;
            }
            .transactions_detail_information_wrap{
                padding: 0.1rem;
                border: 1px solid #d7d9e0;
                margin-bottom: 0.4rem;
                .information_props_wrap{
                    @include flex;
                    flex-direction:column;
                    margin-bottom: 0;
                    .information_value{
                        overflow-x:auto;
                        -webkit-overflow-scrolling:touch;
                        overflow-y:hidden;
                        color: #a2a2ae;
                    }
                }
            }
            .transactions_detail_title {
                height: 0.3rem;
                line-height: 0.3rem;
                font-size: 0.22rem;
                color: #000000;
                margin-right: 0.2rem;
                @include fontWeight;
            }
            .transactions_detail_wrap_hash_var {
                overflow-x: auto;
                -webkit-overflow-scrolling:touch;
                height: 0.3rem;
                line-height: 0.3rem;
                font-size: 0.18rem;
                color: #a2a2ae;
            }
        }
    }
    .link_active_style{
        a{
            color:var(--bgColor) !important;
        }
        color:var(--bgColor) !important;
        cursor:pointer;
    }
    .information_value_fixed {
        display: flex;
        align-items: center;
        & > span {
            margin-right: 0.06rem;
        }
        .fail_status {
            color: #fa8593;
        }
        .question_icon_div {
            background-image: url(../assets/question_icon.png) !important;
        }
        .info_icon_div {
            width: 0.14rem;
            height: 0.14rem;
            position: relative;
            background: url(../assets/info_icon.png) no-repeat top left / 14px 14px;
            cursor: pointer;
            &:hover {
                .tooltip_span {
                    display: block;
                    position: fixed;
                    opacity: 0;
                }
            }
            .tooltip_span {
                display: none;
                z-index: 1000;
                color: #ffffff;
                background-color: #000000;
                border-radius: 0.04rem;
                line-height: 16px;
                div {
                    padding: 8px 15px;
                    & > p {
                        white-space: nowrap;
                    }
                }
                & > i {
                    width: 0;
                    height: 0;
                    border: 0.06rem solid transparent;
                    content: "";
                    display: block;
                    position: absolute;
                    border-top-color: #000000;
                    margin-left: -6px;
                }
            }
            .tooltip_span_word_warp {
                word-break: break-all;
                word-wrap: break-word;
                white-space: normal;
            }
        }
    }
</style>
