<template>
    <div class="voting_card_content">
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
        <div class="voting_title_container" :class="showTitle ? 'home_style': ''">
            <div>
                <span><i :style="{color:flHighlightParticipation ? 'var(--bgColor)' : '#D7DCE0'}" class="iconfont iconBondedTokens"></i>Participation {{participationThreshold}}%</span>
                <span><i :style="{color:flShowPassThreshold ? 'var(--bgColor)' : '#D7DCE0'}" class="iconfont iconBondedTokens"></i>Pass Threshold {{passThreshold}}%</span>
            </div>
            <div class="voting_card_right_content">
                <span><i :style="{color:flShowVoteThreshold ? '#FE8A8A' : '#D7DCE0'}" class="iconfont iconBondedTokens"></i>Veto Threshold {{voteThreshold}}%</span>
                <div v-if="flShowHourLeft">
                    <span><i style="color:#5AC8FA;" class="iconfont iconHoursLeft"></i>{{hourLeft}} Left</span>
                </div>
            </div>
        </div>
        <div class="voting_content">
            <div class="voting_left_container">
                <span class="delegator_voted_content" v-show="flShowTotalVoted">0.00%</span>
                <span class="delegator_voted_content" v-show="!flShowTotalVoted">0.00%</span>
                <span class="yes_content">Yes {{yesVotingPowerWidth === 'NaN' ? '0.00' : yesVotingPowerWidth ? yesVotingPowerWidth : '0.00'}}%</span>
                <span class="no_content">No {{noVotingPowerWidth === 'NaN' ? '0.00' : noVotingPowerWidth ? noVotingPowerWidth : '0.00'}}%</span>
            </div>
            <div class="voting_center_container">
                <div class="voting_progress_bar_content">
                    <span class="min_value_content" :style="minTotalTipStyleNumber" v-show="flShowTotalVoted">
                        <span class="min_value_title">{{delegatorVoted}}% by Delegator</span>
                    </span>
                    <div class="default_progress_bar_content" :style="{background: totalVoted || delegatorVoted ? 'var(--bgColor)' : ''}"></div>
                    <div class="min_deposit_bar_content" :style="minVotingPowerStyleObj"></div>
                </div>
                <div class="voting_bottom_progress_bar_content" :style="minVotingPowerStyleObj">
                    <div class="voting_bottom_min_deposit_bar_content" :style="yesVotingPowerStyleObj"></div>
                    <div class="voting_bottom_total_deposit_bar_content" :style="noVotingPowerStyleObj"></div>
                    <div class="voting_bottom_default_progress_bar_content" :style="abstainVotingPowerStyleObj"></div>
                    <div class="voting_bottom_burned_progress_bar_content" :style="vetoVotingPowerStyleObj"></div>
                </div>
            </div>
            <div class="voting_right_container">
                <span class="participation_threshold_content" v-show="flShowTotalVoted">{{totalVoted ? totalVoted : '0.00'}}% Participation</span>
                <span class="participation_threshold_content" v-show="!flShowTotalVoted">{{delegatorVoted ? delegatorVoted : '0.00'}}% Participation</span>
                <span class="veto_content">{{vetoVotingPowerWidth === 'NaN' ? '0.00' : vetoVotingPowerWidth ? vetoVotingPowerWidth : '0.00'}}% NoWithVeto</span>
                <span class="abstain_content">{{abstainVotingPowerWidth === 'NaN' ? '0.00' : abstainVotingPowerWidth ? abstainVotingPowerWidth : '0.00'}}% Abstain</span>
            </div>
        </div>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MVotingCard",
        props:{
	        votingBarObj:{
				type: Object
            },
	        showTitle:{
		        type: Boolean
	        }
        },
        watch:{
	        votingBarObj(votingBarObj){
                let copyVotingBarObj = JSON.parse(JSON.stringify(votingBarObj));
	        	this.formatVotingBarObj(copyVotingBarObj)
            }
        },

        data(){
			return {
                totalVoted: '',
				proposalId:'',
				title:'',
                participationThreshold:'',
				hourLeft: '',
                validatorVotedPower: '',
                delegatorVotedPower: '',
                delegatorVoted:'',
                systemVotingPower:'',
				totalVotedNumber:'',
				yesVotingPowerWidth:'',
				noVotingPowerWidth:'',
				vetoVotingPowerWidth:'',
				abstainVotingPowerWidth:'',
                passThreshold:'',
                voteThreshold:'',
				flShowHourLeft: false,
                flShowPassThreshold:false,
                flShowVoteThreshold:false,
                totalVotedGreaterThan: false,
                flShowTotalVoted:false,
				flHighlightParticipation:false,
                votingTimer:null,
				minVotingPowerStyleObj:{
                	width:""
                },
				minVotingPowerStyleNumber:'',
                minTotalTipStyleNumber:{
	                left:""
                },
                yesVotingPowerStyleObj:{
	                width:""
                },
				noVotingPowerStyleObj:{
					width:""
				},
                vetoVotingPowerStyleObj:{
	                width:""
                },
				abstainVotingPowerStyleObj:{
					width:""
				}
            }
        },
		mounted(){
			let copyVotingBarObj = JSON.parse(JSON.stringify(this.votingBarObj));
			this.formatVotingBarObj(copyVotingBarObj)
        },
        methods:{
	        formatVotingBarObj(votingBarObj){
	        	if(votingBarObj){
			        this.title = votingBarObj.title;
			        this.proposalId = votingBarObj.proposal_id;
			        this.systemVotingPower = votingBarObj.voting_power_for_height;
			        this.participationThreshold = `${Tools.formatPercent(votingBarObj.level.gov_param.participation)}`;
			        this.passThreshold = `${Tools.formatPercent(votingBarObj.level.gov_param.pass_threshold)}`;
			        this.voteThreshold = `${Tools.formatPercent(votingBarObj.level.gov_param.veto_threshold)}`;
			        this.getVotingEndTime(votingBarObj.voting_end_time);
		        }
		        if(votingBarObj && Array.isArray(votingBarObj.votes) && votingBarObj.votes.length > 0){
		        	this.flShowTotalVoted = true;
			        this.getTotalVoted(votingBarObj.votes);
			        this.getYesVotingPower(votingBarObj.votes);
			        this.getNoVotingPower(votingBarObj.votes);
			        this.getVetoVotingPower(votingBarObj.votes);
			        this.getAbstainVotingPower(votingBarObj.votes);
                }
		        if(votingBarObj && votingBarObj.final_votes && JSON.stringify(votingBarObj.final_votes) !== "{}"){
                    this.getFinalVotes(votingBarObj.final_votes);
                }
            },
	        getFinalVotes(finalVotes){
		        this.flShowTotalVoted = false;
                let optionTotalNumber = Number(finalVotes.yes) + Number(finalVotes.no) + Number(finalVotes.abstain) + Number(finalVotes.no_with_veto);
		        this.delegatorVoted = ((Number(optionTotalNumber) / Number(finalVotes.system_voting_power)) * 100).toFixed(2);
		        this.yesVotingPowerWidth = ((Number(finalVotes.yes) / Number(optionTotalNumber)) * 100).toFixed(2);
		        this.$set(this.yesVotingPowerStyleObj,'width',`${this.yesVotingPowerWidth}%`);
		        this.abstainVotingPowerWidth = ((Number(finalVotes.abstain) / Number(optionTotalNumber)) * 100).toFixed(2);
		        this.$set(this.abstainVotingPowerStyleObj,'width',`${this.abstainVotingPowerWidth}%`);
		        this.vetoVotingPowerWidth = ((Number(finalVotes.no_with_veto) / Number(optionTotalNumber)) * 100).toFixed(2);
		        this.$set(this.vetoVotingPowerStyleObj,'width',`${this.vetoVotingPowerWidth}%`);
		        this.noVotingPowerWidth = ((Number(finalVotes.no) / Number(optionTotalNumber)) * 100).toFixed(2);
		        this.$set(this.noVotingPowerStyleObj,'width',`${this.noVotingPowerWidth}%`);
		        if(Number(this.yesVotingPowerWidth) > Number(this.passThreshold)){
			        this.flShowPassThreshold = true;
		        }else {
			        this.flShowPassThreshold = false
		        }
		        if(Number(this.vetoVotingPowerWidth) > Number(this.voteThreshold)){
			        this.flShowVoteThreshold = true;
		        }else {
			        this.flShowVoteThreshold = false
		        }
		        this.totalVoted = ((optionTotalNumber / finalVotes.system_voting_power) * 100).toFixed(2);
		        this.$store.commit('currentParticipationValue',this.totalVoted);
		        this.$store.commit('currentYesValue',this.yesVotingPowerWidth);
		        this.$store.commit('currentNoValue',this.noVotingPowerWidth);
		        this.$store.commit('currentNoWithVetoValue',this.vetoVotingPowerWidth);
		        this.$store.commit('currentAbstainValue',this.abstainVotingPowerWidth);
		        this.setStyleFunc()
	        },
	        getVotingEndTime(time){
		        if(time){
			        let that = this;
			        let currentServerTime = new Date().getTime() + this.diffMilliseconds;
			        if(new Date(time).getTime() >  currentServerTime){
				        that.hourLeft = Tools.formatAge(new Date(time).getTime(),currentServerTime);
				        that.flShowHourLeft = true;
			        }
		        	clearInterval(this.votingTimer);
			        this.votingTimer = setInterval(() => {
				        currentServerTime = new Date().getTime() + this.diffMilliseconds;
				        if(new Date(time).getTime() >  currentServerTime){
					        that.hourLeft = Tools.formatAge(new Date(time).getTime(),currentServerTime);
					        that.flShowHourLeft = true;
				        }else {
					        that.flShowHourLeft = false;
				        }
                    },1000)

		        }
	        },
            getTotalVoted(votTx){
	            let totalDelegatorPower =[];
                votTx.forEach( item => {
	                totalDelegatorPower.push(item.del_voting_power)
                });
                this.delegatorVotedPower = totalDelegatorPower.reduce( (a,b) => {
                    return a + b
                });
	            let totalValidatorPower = [];
                votTx.forEach( item => {
	                totalValidatorPower.push(item.val_voting_power)
                });
               this.validatorVotedPower = totalValidatorPower.reduce( (a,b) => {
                    return a + b
                });
               this.totalVotedNumber =  this.validatorVotedPower + this.delegatorVotedPower;
               this.delegatorVoted = (Number(this.delegatorVotedPower) / (Number(this.delegatorVotedPower) + Number(this.validatorVotedPower)) * 100).toFixed(2);
	           this.$set(this.minTotalTipStyleNumber,'left',`${this.delegatorVoted}%`);
               this.totalVoted = (((this.delegatorVotedPower + this.validatorVotedPower) / this.systemVotingPower) *100).toFixed(2);
               this.$store.commit('currentParticipationValue',this.totalVoted)
               this.setStyleFunc();
            },
            getYesVotingPower(votTx){
	        	let yesVotingPowerArr = [];
	            votTx.forEach((item) => {
	            	if(item.option === "Yes"){
			            yesVotingPowerArr.push((item.del_voting_power + item.val_voting_power))
                    }
                });
	            let yesVotingPower;
                if(yesVotingPowerArr.length > 0){
	                yesVotingPower = yesVotingPowerArr.reduce((a,b) => {
		                return a + b
	                });
                }
	            this.yesVotingPowerWidth = ((Number(yesVotingPower) / Number(this.totalVotedNumber)) * 100).toFixed(2);
                if(Number(this.yesVotingPowerWidth) > Number(this.passThreshold)){
                	this.flShowPassThreshold = true
                }else {
	                this.flShowPassThreshold = false
                }
	            this.$store.commit('currentYesValue',this.yesVotingPowerWidth)
                this.$set(this.yesVotingPowerStyleObj,'width',`${this.yesVotingPowerWidth}%`)
            },
	        getNoVotingPower(votTx){
		        let noVotingPowerArr = [];
		        votTx.forEach((item) => {
			        if(item.option === "No"){
				        noVotingPowerArr.push((item.del_voting_power + item.val_voting_power))
			        }
		        });
		        let yesVotingPower;
		        if(noVotingPowerArr.length > 0){
			        yesVotingPower = noVotingPowerArr.reduce((a,b) => {
				        return a + b
			        });
		        }
		        this.noVotingPowerWidth = ((Number(yesVotingPower) / Number(this.totalVotedNumber)) * 100).toFixed(2);
		        this.$store.commit('currentNoValue',this.noVotingPowerWidth);
		        this.$set(this.noVotingPowerStyleObj,'width',`${this.noVotingPowerWidth}%`)
	        },
	        getVetoVotingPower(votTx){
		        let vetoVotingPowerArr = [];
		        votTx.forEach((item) => {
			        if(item.option === "NoWithVeto"){
				        vetoVotingPowerArr.push((item.del_voting_power + item.val_voting_power))
			        }
		        });
		        let yesVotingPower;
		        if(vetoVotingPowerArr.length > 0){
			        yesVotingPower = vetoVotingPowerArr.reduce((a,b) => {
				        return a + b
			        });
		        }
		        this.vetoVotingPowerWidth = ((Number(yesVotingPower) / Number(this.totalVotedNumber)) * 100).toFixed(2);
		        if(Number(this.vetoVotingPowerWidth) > Number(this.voteThreshold)){
		        	this.flShowVoteThreshold = true
                }else {
			        this.flShowVoteThreshold = false
                }
		        this.$store.commit('currentNoWithVetoValue',this.vetoVotingPowerWidth);
		        this.$set(this.vetoVotingPowerStyleObj,'width',`${this.vetoVotingPowerWidth}%`)
	        },
	        getAbstainVotingPower(votTx){
		        let abstainVotingPowerArr = [];
		        votTx.forEach((item) => {
			        if(item.option === "Abstain"){
				        abstainVotingPowerArr.push((item.del_voting_power + item.val_voting_power))
			        }
		        });
		        let yesVotingPower;
		        if(abstainVotingPowerArr.length > 0){
			        yesVotingPower = abstainVotingPowerArr.reduce((a,b) => {
				        return a + b
			        });
		        }
		        this.abstainVotingPowerWidth = ((Number(yesVotingPower) / Number(this.totalVotedNumber)) * 100).toFixed(2);
		        this.$store.commit('currentAbstainValue',this.abstainVotingPowerWidth);
		        this.$set(this.abstainVotingPowerStyleObj,'width',`${this.abstainVotingPowerWidth}%`)
	        },
            setStyleFunc(){
	        	if(Number(this.totalVoted) > Number(this.participationThreshold)){
	        		this.flHighlightParticipation = true
		        }else if(Number(this.totalVoted) < Number(this.participationThreshold)){
			        this.flHighlightParticipation = false;
			        this.totalVotedGreaterThan = false;
                    this.minVotingPowerStyleNumber = ((Number(this.totalVoted) / Number(this.participationThreshold))* 100).toFixed(0);
                }
            }
        }
	}
</script>

<style scoped lang="scss">
    .voting_card_content{
        flex: 1;
        justify-content: space-between;
        min-height: 2rem;
        background: #fff;
        box-sizing: border-box;
        padding: 0.2rem;
        border: 0.01rem solid #d7d9e0;
        margin: 0 0 0.2rem 0;
        .title_container{
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
                    color: var(--bgColor);
                }
            }
            .view_all_content{
                color: var(--bgColor);
                span{
                    line-height: 1;
                    border-bottom: 0.01rem solid var(--bgColor);
                    a{
                        font-size: 0.14rem;
                        color: var(--bgColor) !important;
                    }
                }
            }

        }
        .voting_title_container{
            font-size: 0.12rem;
            display: flex;
            i{
                margin-right: 0.06rem;
                font-size: 0.14rem;
            }
            .iconVeto{
                color: #ff0000;
            }
            .iconHoursLeft{
                color: #787C99;
            }
            span{
                margin-right: 0.12rem;
                color: #22252A;
            }
            .voting_card_right_content{
                display: flex;
            }
        }
        .home_style{
            margin-top: 0.1rem;
        }
        .voting_content{
            margin-top: 0.5rem;
            display: flex;
            .voting_left_container{
                display: flex;
                flex-direction: column;
                align-items: flex-end;
                .delegator_voted_content{
                    color: #171D44;
                    font-size: 0.12rem;
                    line-height: 0.14rem;
                }
                .yes_content{
                    color: #171D44;
                    font-size: 0.12rem;
                    line-height: 0.14rem;
                    margin-top: 0.26rem;
                }
                .no_content{
                    font-size: 0.12rem;
                    color: #171D44;
                    line-height: 0.14rem;
                    margin-top: 0.1rem;
                }

            }
            .voting_center_container{
                flex: 1;
                width: 100%;
                margin: 0 0.1rem;
                display: flex;
                flex-direction: column;
                justify-content: space-between;
                .voting_progress_bar_content{
                    flex: 1;
                    width: 100%;
                    position: relative;
                    bottom: 0.07rem;
                    margin-top: 0.06rem;
                    .min_value_content{
                        color: #FF9942;
                        bottom: 0.5rem;
                        position: absolute;
                        .min_value_title{
                            position: relative;
                            display: inline-block;
                            right: 30%;
                            width: 100%;
                            white-space: nowrap;
                            font-size: 0.12rem;
                            margin-bottom: 0.08rem;
                            &:after{
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
                        z-index: 7;
                        background: #E5E9FB;
                    }
                    .burned_progress_bar_content{
                        width: 30%;
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                        background: #D25F59;
                        z-index: 1;
                    }
                    .min_deposit_bar_content{
                        background: var(--bgColor);
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                        z-index: 2;
                    }

                }
                .voting_bottom_progress_bar_content{
                    margin-bottom: 0.14rem;
                    display: flex;
                    align-items: center;
                    border-radius: 0.06rem;
                    overflow: hidden;
                    height: 0.12rem;
                    .voting_bottom_default_progress_bar_content{
                        height: 0.12rem;
                        background: #CCDCFF;
                        z-index: 1;
                    }
                    .voting_bottom_burned_progress_bar_content{
                        height: 0.12rem;
                        background: #FE8A8A;
                        z-index: 1;
                    }
                    .voting_bottom_min_deposit_bar_content{
                        background: var(--bgColor);
                        height: 0.12rem;
                        z-index: 2;
                    }
                    .voting_bottom_total_deposit_bar_content{
                        background: #FFCF65;
                        height: 0.12rem;
                        z-index: 1;
                    }
                }
            }
            .voting_right_container{
                display: flex;
                flex-direction: column;
                .participation_threshold_content{
                    color: #171D44;
                    font-size: 0.12rem;
                    line-height: 0.14rem;
                }
                .veto_content{
                    color: #171D44;
                    font-size: 0.12rem;
                    line-height: 0.14rem;
                    margin-top: 0.26rem;
                }
                 .abstain_content{
                    font-size: 0.12rem;
                    color: #171D44;
                    line-height: 0.14rem;
                    margin-top: 0.1rem;
                 }
            }
        }
    }
    @media screen and (max-width: 1280px){
        .voting_title_container{
            flex-direction: column;
            .voting_card_right_content{
                display: flex;
            }
        }
    }
    @media screen and (max-width: 910px){
        .voting_card_content{
            margin: 0 ;
        }

    }
</style>
