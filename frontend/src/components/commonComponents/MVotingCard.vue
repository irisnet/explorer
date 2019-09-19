<template>
    <div class="voting_card_content">
        <div class="voting_title_container">
            <span><i :style="{color:flShowPassThreshold ? '#44C190' : '#D7DCE0'}" class="iconfont iconPass"></i>Pass Threshold {{passThreshold}} %</span>
            <span><i :style="{color:flShowVoteThreshold ? '#FE8A8A' : '#D7DCE0'}" class="iconfont iconVeto"></i>Vote Threshold {{voteThreshold}}%</span>
            <span v-show="hourLeft > 1"><i class="iconfont iconHoursLeft"></i>{{hourLeft === 1 ? `${hourLeft} Hour Left` : `${hourLeft} Hours Left` }}</span>
            <span v-show="hourLeft < 1 && hourLeft > 0"><i class="iconfont iconHoursLeft"></i>{{ `${hourLeft} Min Left` }}</span>
        </div>
        <div class="voting_content">
            <div class="voting_left_container">
                <span class="delegator_voted_content">DelegatorVoted {{delegatorVoted}}%</span>
                <span class="yes_content">Yes {{yesVotingPowerWidth === 'NaN' ? '0.00' : yesVotingPowerWidth}}%</span>
                <span class="abstain_content">Abstain {{abstainVotingPowerWidth === 'NaN' ? '0.00' : abstainVotingPowerWidth}}%</span>
            </div>
            <div class="voting_center_container">
                <div class="voting_progress_bar_content">
                    <span class="min_value_content" :style="minTotalTipStyleNumber">
                        <span class="min_value_title">Total Voted {{totalVoted}} %</span>
                    </span>
                    <div class="default_progress_bar_content" :style="{background:totalVotedGreaterThan ? '#79C9FF' : ''}"></div>
                    <div class="min_deposit_bar_content" :style="minVotingPowerStyleObj"></div>
                </div>
                <div class="voting_bottom_progress_bar_content" :style="minVotingPowerStyleObj">
                    <div class="voting_bottom_min_deposit_bar_content" :style="yesVotingPowerStyleObj"></div>
                    <div class="voting_bottom_default_progress_bar_content" :style="abstainVotingPowerStyleObj"></div>
                    <div class="voting_bottom_total_deposit_bar_content" :style="noVotingPowerStyleObj"></div>
                    <div class="voting_bottom_burned_progress_bar_content" :style="vetoVotingPowerStyleObj"></div>
                </div>
            </div>
            <div class="voting_right_container">
                <span class="participation_threshold_content">{{participationThreshold}}% Participation Threshold</span>
                <span class="veto_content">{{vetoVotingPowerWidth === 'NaN' ? '0.00' : vetoVotingPowerWidth}}% Veto</span>
                <span class="no_content">{{noVotingPowerWidth === 'NaN' ? '0.00' : noVotingPowerWidth}}% No</span>
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
            }
        },
        watch:{
	        votingBarObj(votingBarObj){
	        	this.formatVotingBarObj(votingBarObj)
            }
        },
        data(){
			return {
                totalVoted: '',
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
                flShowPassThreshold:false,
                flShowVoteThreshold:false,
                totalVotedGreaterThan: false,
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
        methods:{
	        formatVotingBarObj(votingBarObj){
		        this.systemVotingPower = votingBarObj.voting_power_for_height;
		        this.participationThreshold = `${((Number(votingBarObj.level.gov_param.participation)) *100).toFixed(2)}`;
		        this.passThreshold = `${(Number(votingBarObj.level.gov_param.pass_threshold) * 100).toFixed(2)}`;
		        this.voteThreshold = `${(Number(votingBarObj.level.gov_param.veto_threshold) * 100).toFixed(2)}`;
		        this.getVotingEndTime(votingBarObj.voting_end_time);
		        if(Array.isArray(votingBarObj.votes) && votingBarObj.votes.length > 0){
			        this.getTotalVoted(votingBarObj.votes);
			        this.getYesVotingPower(votingBarObj.votes);
			        this.getNoVotingPower(votingBarObj.votes);
			        this.getVetoVotingPower(votingBarObj.votes);
			        this.getAbstainVotingPower(votingBarObj.votes);
                }
            },
	        getVotingEndTime(time){
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
            getTotalVoted(votTx){
                let totalDelegatorPower = votTx.map( item => {
                    return ++item.del_voting_power
                });
                this.delegatorVotedPower = totalDelegatorPower.reduce( (a,b) => {
                    return a + b
                });
                let totalValidatorPower = votTx.map( item => {
                    return ++item.val_voting_power
                });
               this.validatorVotedPower = totalValidatorPower.reduce( (a,b) => {
                    return a + b
                });
               this.totalVotedNumber =  this.validatorVotedPower + this.delegatorVotedPower;
               this.delegatorVoted = (this.validatorVotedPower / (this.delegatorVotedPower + this.validatorVotedPower) * 100).toFixed(2);
               this.totalVoted = (((this.delegatorVotedPower + this.validatorVotedPower) / this.systemVotingPower) *100).toFixed(2);
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
		        this.$set(this.abstainVotingPowerStyleObj,'width',`${this.abstainVotingPowerWidth}%`)
	        },
            setStyleFunc(){
	        	if(Number(this.totalVoted) > Number(this.participationThreshold)){
	        		this.flShowPassThreshold = true;
	        		this.totalVotedGreaterThan = true;
                    let diff = this.totalVoted - this.participationThreshold;
			        this.$set(this.minTotalTipStyleNumber,'left',`${(100 - diff)}%`);
			        this.$set(this.minVotingPowerStyleObj,'width',`${(100 - diff)}%`)
		        }else if(Number(this.totalVoted) < Number(this.participationThreshold)){
			        this.flShowPassThreshold = false;
			        this.totalVotedGreaterThan = false;
                    this.minVotingPowerStyleNumber = ((Number(this.totalVoted) / Number(this.participationThreshold))* 100).toFixed(0);
                    this.$set(this.minTotalTipStyleNumber,'left',`${((Number(this.totalVoted) / Number(this.participationThreshold))* 100).toFixed(0)}%`);
			        this.$set(this.minVotingPowerStyleObj,'width',`${((Number(this.totalVoted) / Number(this.participationThreshold))* 100).toFixed(0)}%`)
                }
            }
        }
	}
</script>

<style scoped lang="scss">
    .voting_card_content{
        flex: 1;
        justify-content: space-between;
        height: 2rem;
        background: #fff;
        box-sizing: border-box;
        padding: 0.2rem;
        border: 0.01rem solid #d7d9e0;
        margin: 0 0 0.2rem 0.1rem;
        .proposal_title_content{

        }
        .voting_title_container{
            font-size: 0.12rem;
            i{
                margin-right: 0.1rem;
                font-size: 0.14rem;
            }
            .iconVeto{
                color: #ff0000;
            }
            .iconHoursLeft{
                color: #787C99;
            }
            span{
                margin-right: 0.5rem;
                color: #22252A;
            }
        }
        .voting_content{
            margin-top: 0.4rem;
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
                .abstain_content{
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
                    .burned_progress_bar_content{
                        width: 30%;
                        height: 0.12rem;
                        border-radius: 0.06rem;
                        position: absolute;
                        background: #D25F59;
                        z-index: 1;
                    }
                    .min_deposit_bar_content{
                        background: #0580D3;
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
                        background: #0580D3;
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
                .no_content{
                    font-size: 0.12rem;
                    color: #171D44;
                    line-height: 0.14rem;
                    margin-top: 0.1rem;
                }
            }
        }
    }
    @media screen and (max-width: 910px){
        .voting_card_content{
            margin: 0 0.1rem;
        }
    }
</style>
