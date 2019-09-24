<template>
    <div class="deposit_card_content">
        <div class="title_container" v-if="showTitle">
            <div class="title_content">
                <router-link :to="`/ProposalsDetail/${proposalId}`">
                    <span class="proposal_id_content">#{{proposalId}}</span>
                    <span class="proposal_title">{{title}}</span>
                </router-link>
            </div>
            <div class="view_all_content">
                <span><router-link :to="`/gov/proposals`">View All</router-link></span>
            </div>
        </div>
        <div class="deposit_title_container">
            <span><i class="iconfont iconParameterChange"></i>{{proposalType}}</span>
            <span v-show="flShowVotingPeriod"><i class="iconfont iconDepositPeriod"></i>Voting Period</span>
            <span v-show="!flShowPass && !flShowReject && !flShowVotingPeriod"><i style="color:#0580d3" class="iconfont iconDepositPeriod-liebiao"></i>DepositPeriod</span>
            <span v-show="flShowPass" ><i class="iconfont iconPass"></i>Passed</span>
            <span v-show="flShowReject"><i class="iconfont iconVeto"></i>Rejected</span>
            <span v-show="hourLeft > 1"><i class="iconfont iconHoursLeft"></i>{{hourLeft === 1 ? `${hourLeft} Hour Left` : `${hourLeft} Hours Left` }}</span>
            <span v-show="hourLeft < 1 && hourLeft > 0"><i class="iconfont iconHoursLeft"></i>{{ `${hourLeft} Min Left` }}</span>
        </div>
        <div class="deposit_content">
            <div class="content">
                <span class="initial_value_content">0</span>
                <div class="progress_bar_content">
                    <span v-show="!flShowBurnAll" class="min_value_content" :style="minValueStyleObj">
                        <span class="min_value_title">Total Deposit {{totalDeposit}}</span>
                    </span>
                    <span v-show="flShowBurnAll" class="min_value_content" :style="burnTipStyle0bj">
                        <span class="min_value_title">Burned Deposit {{burnValue}}</span>
                    </span>
                    <div class="default_progress_bar_content" v-show="!flBurnAll" :class="flShowDiffStyle ? 'diff_blue' : ''"></div>
                    <div class="default_progress_bar_content" v-show="flBurnAll" :class="flBurnAll ? 'diff_burn_red' : ''"></div>
                    <div class="burned_progress_bar_content" :style="burnStyle0bj" :class="flShowDiffStyle ? 'show_burn_style' : ''"></div>
                    <div class="min_deposit_bar_content" v-if="flShowBlue && flShowDiffStyle" :style="minDepositStyleObject" :class="flShowBlue ? 'showBlue' : ''"></div>
                    <div class="min_deposit_bar_content" v-if="flHideBlue && !flShowDiffStyle" :style="minDepositStyleObject" :class="flHideBlue ? 'hideBlue' : ''"></div>
                    <div class="total_deposit_bar_content"></div>
                    <div class="init_content"></div>
                    <span v-show="!flShowBurnAll" class="max_value_content" :style="initialDepositStyleObj">
                        <span class="max_value_title">Initial Deposit {{initialDeposit}}</span>
                    </span>
                    <span v-show="flShowBurnAll" class="max_value_content no_tool_tip_style" :style="minValueStyleObj">
                        <span class="max_value_title">Total Deposit {{totalDeposit}}</span>
                    </span>
                </div>
                <span class="total_value_content">{{minDepositToken}} {{minDepositTokenDenom}} MinDeposit</span>
            </div>
        </div>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MDepositCard",
        props:{
			depositObj:{
				type: Object
            },
	        burnPercent:{
		        type: Number
            },
	        status:{
		        type: String
	        },
	        showTitle:{
		        type: Boolean
            }
        },
        data () {
			return {
				proposalType:'',
                title:'',
				proposalId:'',
				hourLeft: '',
                minDepositToken:'',
                minDepositTokenDenom:'',
                totalDeposit: "",
                initialDeposit:'',
                flShowBlue: false,
                flShowDepositPeriod: false,
                flHideBlue: false,
				flShowDiffStyle:false,
                blueWidth: '',
                flBurnAll: false,
                burnValue:'',
				flTotalGreaterThanMin: false,
                flShowBurnAll:false,
                localBurnPercent: '',
                flShowPass: false,
                flShowReject:false,
				flShowVotingPeriod: false,
				minDepositStyleObject:{
					width: ''
                },
				minValueStyleObj: {
					left:''
                },
                burnStyle0bj:{
					width:''
                },
				initialDepositStyleObj:{
					left:''
                },
				burnTipStyle0bj:{
					left:''
                }
            }
        },
        watch:{
            depositObj(depositObj){
                this.formatDepositObj(depositObj);
                this.flShowProgressBar();
                this.flShowInitial();
                if(this.localBurnPercent){
                	this.flShowBurn(this.localBurnPercent)
                }
            },
	        burnPercent(burnPercent){
            	this.localBurnPercent = burnPercent;
            	this.flShowBurn(burnPercent);
            },
	        status(status){
            	if(status === 'Passed'){
            		this.flShowPass = true
                }else if(status === 'Rejected'){
		            this.flShowReject = true
                }
            }

        },
        methods:{
	        formatDepositObj(depositObj){
	        	if(depositObj){
	        		if(depositObj.status === 'VotingPeriod'){
	        			this.flShowVotingPeriod = true
                    }
	        		this.title = depositObj.title;
	        		this.proposalId = depositObj.proposal_id;
                    this.proposalType = depositObj.type;
			        this.getAgeHour(depositObj.deposit_end_time);
			        this.setMinDepositToken(depositObj);
			        this.setTotalDepositToken(depositObj);
			        this.setInitialDeposit(depositObj);
                }


	        },
            setMinDepositToken(depositObj){
	            if(depositObj.level.gov_param.min_deposit.amount === 0){
		            this.minDepositToken = '0 IRIS'
	            }else {
		            this.minDepositToken = `${Tools.numberMoveDecimal(depositObj.level.gov_param.min_deposit.amount)}`;
		            this.minDepositTokenDenom = `${Tools.formatDenom(depositObj.level.gov_param.min_deposit.denom)}`

	            }
            },
            setTotalDepositToken(depositObj){
	            if(depositObj.total_deposit.amount === 0){
		            this.totalDeposit = '0'
	            }else {
		            this.totalDeposit = `${Tools.numberMoveDecimal(depositObj.total_deposit.amount)}`;
               }
            },
            setInitialDeposit(depositObj){
                if(depositObj.intial_deposit.amount === 0){
	                this.initialDeposit = '0'
                }else {
	                this.initialDeposit = `${Tools.numberMoveDecimal(depositObj.intial_deposit.amount)}`
                }
            },
            getAgeHour(time){
	        	if(time){
	        		let localTime = new Date().getTimezoneOffset() * 60 * 1000;
			        let localUtcTime;
	        		if(localTime < 0){
				        localUtcTime = Math.floor(new Date().getTime()) + localTime;
			        }else {
				        localUtcTime = Math.floor(new Date().getTime()) - localTime;
                    }
                    let depositEndUTCTime = new Date(Tools.conversionTimeToUTCByValidatorsLine(time)).getTime();
			        let diffTime = (depositEndUTCTime - localUtcTime) / 1000;
	        		let hourLeft = Math.ceil(diffTime/ 3600);
	        		if(hourLeft < 1){
				        this.hourLeft = Math.ceil(diffTime/ 60)
                    }else {
				        this.hourLeft = hourLeft
                    }

                }

            },
	        flShowProgressBar(){
		        if(Number(this.totalDeposit) === Number(this.minDepositToken)){
			        this.flHideBlue = true;
			        this.flShowDepositPeriod = true;
			        this.$set(this.minDepositStyleObject,'width','100%');
			        this.$set(this.minValueStyleObj,'left','100%');
			        this.$set(this.initialDepositStyleObj,'left','100%');
		        }else if(Number(this.totalDeposit) > Number(this.minDepositToken)){
			        this.flShowBlue = true;
			        this.flShowDiffStyle = true;
			        this.flShowDepositPeriod = true;
			        // let diffNumber = this.totalDeposit - this.minDepositToken;
			        // let diffContent = ((diffNumber / this.minDepositToken) * 100).toFixed(1);
			        // this.$set(this.minDepositStyleObject,'width',`${(100 - diffContent === 0 ? 50 : 100 - diffContent)}%`);
			        this.$set(this.minDepositStyleObject,'width',`${((this.minDepositToken/this.totalDeposit)* 100).toFixed(2)}%`);
			        this.$set(this.minValueStyleObj,'left',"100%");
		        }else if(Number(this.totalDeposit) < Number(this.minDepositToken)){
			        this.flHideBlue = true;
			        this.$set(this.minDepositStyleObject,'width',`${((this.totalDeposit / this.minDepositToken)*100).toFixed(1)}%`);
			        this.$set(this.minValueStyleObj,'left',`${((this.totalDeposit / this.minDepositToken)*100).toFixed(1)}%`);
		        }
	        },
	        flShowInitial(){
		        if(Number(this.initialDeposit) === Number(this.minDepositToken)){
			        this.$set(this.initialDepositStyleObj,'left','100%');
		        }else if(Number(this.initialDeposit) > Number(this.minDepositToken)){
			        let diffNumber = this.initialDeposit - this.minDepositToken;
			        // let diffContent = ((diffNumber / this.minDepositToken) * 100).toFixed(1);
			        this.$set(this.minDepositStyleObject,'width',`${((this.minDepositToken / this.initialDeposit) *100).toFixed(2)}%`);
			        this.$set(this.minValueStyleObj,'left','100%');
			        this.$set(this.initialDepositStyleObj,'left','100%');
		        }else if(Number(this.initialDeposit) < Number(this.minDepositToken)){
			        this.$set(this.initialDepositStyleObj,'left',`${((this.initialDeposit / this.minDepositToken)*100).toFixed(1)}%`);
		        }
            },
	        flShowBurn(burnPercent){
		        if(burnPercent === 0){

		        }else if(burnPercent === 0.2){
	        		this.burnValue = (this.totalDeposit * burnPercent).toFixed(2);
			        this.flShowDiffStyle = true;
			        let diffNumber = this.totalDeposit - this.minDepositToken;
			        let diffContent = ((diffNumber / this.minDepositToken) * 100).toFixed(1);
			        this.$set(this.burnStyle0bj,'width',`${Math.floor((100 - diffContent === 0 ? 50 : 100 - diffContent) * burnPercent)}%`);
			        this.$set(this.burnTipStyle0bj,'left',`${Math.floor((100 - diffContent === 0 ? 50 : 100 - diffContent) * burnPercent)}%`);
			        this.flShowBurnAll = true
		        }else if(burnPercent === 1){
			        this.burnValue = this.totalDeposit;
			  /*      let diffNumber = this.totalDeposit - this.minDepositToken;
			        let diffContent = ((diffNumber / this.minDepositToken) * 100).toFixed(1);
			        this.$set(this.burnStyle0bj,'width',`${(100 - diffContent === 0 ? 50 : 100 - diffContent)}%`);*/
			        this.$set(this.burnStyle0bj,'width',`${((this.minDepositToken / this.totalDeposit) *100).toFixed(2)}%`);
			        this.$set(this.burnTipStyle0bj,'left','100%');
			        this.flShowDiffStyle = true;
			        this.flBurnAll = true;
			        this.flShowBurnAll = true
		        }
            }
        },
        mounted () {
	        this.formatDepositObj(this.depositObj);
	        this.flShowProgressBar();
	        this.flShowInitial()
	        if(this.localBurnPercent){
		        this.flShowBurn(this.localBurnPercent)
	        }
        }
	}
</script>

<style scoped lang="scss">
    .deposit_card_content{
        flex: 1;
        justify-content: space-between;
        min-height: 2rem;
        height: auto;
        background: #fff;
        box-sizing: border-box;
        padding: 0.2rem;
        border: 0.01rem solid #d7d9e0;
        margin-bottom: 0.3rem;
        .title_container{
            margin-bottom: 0.2rem;
            display: flex;
            align-items: center;
            justify-content: space-between;
            .title_content{
                line-height: 1;
                font-size: 0.16rem;
                .proposal_id_content{
                    color: #787C99;
                    margin-right: 0.1rem;
                }
                .proposal_title{
                    color: #0580D3;
                }
            }
            .view_all_content{
                color: #0580D3;
                span{
                    line-height: 1;
                    border-bottom: 0.01rem solid #0580D3;
                    a{
                        font-size: 0.14rem;
                        color: #0580D3 !important;
                    }
                }
            }

        }
        .deposit_title_container{
            font-size: 0.12rem;
            i{
                margin-right: 0.1rem;
                font-size: 0.14rem;
            }
            .iconParameterChange{
                color: #FF8000;
            }
            .iconDepositPeriod{
                color: var(--bgColor);
            }
            .iconPass{
                color: #0580D3;
            }
            .iconVeto{
                color: #FFAAA6;
            }
            .iconHoursLeft{
                color: #787C99;
            }
            span{
                margin-right: 0.5rem;
                color: #22252A;
            }
        }
        .deposit_content{
            margin-top: 0.4rem;

            .content{
                display: flex;
                align-items: center;
                .initial_value_content{
                    font-size: 0.12rem;
                    color: #22252A;
                    margin-right: 0.2rem;
                }
                .progress_bar_content{
                    flex: 1;
                    width: 100%;
                    position: relative;
                    bottom: 0.07rem;
                    margin-right: 0.2rem;
                    .min_value_content{
                        color: #FF9942;
                        bottom:0.1rem;
                        position: absolute;
                        .min_value_title{
                            font-size: 0.12rem;
                            position: relative;
                            display: inline-block;
                            right: 30%;
                            width: 100%;
                            white-space: nowrap;
                            &::after{
                                width: 0;
                                height: 0;
                                border: 0.06rem solid transparent;
                                content: "";
                                display: block;
                                position: absolute;
                                border-top-color: #FF9942;
                                left: 24%;
                                bottom: -0.12rem;
                            }
                        }
                    }
                    .default_progress_bar_content{
                        width: 100%;
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                        background: #E5E9FB;
                        z-index: 1;
                    }
                    .diff_blue{
                        background: #0580D3 !important;
                    }
                    .diff_burn_red{
                        background: #FFAAA6 !important;
                    }
                    .burned_progress_bar_content{
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                        background: #D25F59;
                    }
                    .show_burn_style{
                        z-index: 3;
                    }
                    .min_deposit_bar_content{
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                    }
                    .showBlue{
                        background: #79C9FF;
                        z-index: 1;
                    }
                    .hideBlue{
                        background: #0580D3;
                        z-index: 1;
                    }
                    .show_max_blue{
                        background: #0580D3;
                    }
                    .total_deposit_bar_content{
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                    }
                    .init_content{
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                    }
                    .max_value_content{
                        position: absolute;
                        top: 0.21rem;
                        .max_value_title{
                            position: relative;
                            right: 30%;
                            width: 100%;
                            white-space: nowrap;
                            font-size: 0.12rem;
                            color: #0580D3;
                            &::after{
                                width: 0;
                                height: 0;
                                border: 0.06rem solid transparent;
                                content: "";
                                display: block;
                                position: absolute;
                                border-bottom-color: #0580D3;
                                left: 24%;
                                top: -0.14rem;
                            }
                        }
                    }
                    .no_tool_tip_style{
                        &::after{
                            width: 0;
                            height: 0;
                            border: none;
                            content: "";
                            display: block;
                            position: absolute;
                            border-bottom-color: transparent;
                            left: -0.06rem;
                            top: -50%;
                        }
                    }

                }
                .total_value_content{
                    color: #22252A;
                    font-size: 0.12rem;
                    margin-left: 0.02rem;
                }
            }
        }
    }
    @media screen and (max-width: 910px){
        .deposit_card_content{
            margin: 0 0 0.2rem 0;
        }
    }
</style>
