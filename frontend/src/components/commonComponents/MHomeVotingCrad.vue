<template>
    <div class="home_voting_container">
        <div class="home_voting_title_container_wrap" v-if="true">
            <div class="title_content">
                <router-link :to="`/ProposalsDetail/${proposalId}`">
                    <span class="proposal_id_content">#{{proposalId}}</span>
                    <span class="proposal_title">{{proposalTile}}</span>
                </router-link>
            </div>
            <div class="view_all_content">
                <span>
                   <router-link :to="`/gov/proposals`">View All</router-link>
                </span>
            </div>
        </div>
        <div class="home_voting_title_container">
            <span><i :style="{color:flShowParameterChange ? '#44C190' : '#D7DCE0'}"  class="iconfont iconPass"></i>ParameterChange</span>
            <span><i :style="{color:flShowRedStyle ? '#3598db' : '#D7DCE0'}" class="iconfont iconDepositPeriod"></i>DepositPeriod</span>
            <span v-show="hourLeft > 1"><i class="iconfont iconHoursLeft"></i>{{hourLeft === 1 ? `${hourLeft} Hour Left` : `${hourLeft} Hours Left` }}</span>
            <span v-show="hourLeft < 1 && hourLeft > 0"><i class="iconfont iconHoursLeft"></i>{{ `${hourLeft} Min Left` }}</span>
        </div>
        <div class="home_voting_content">
            <div class="home_voting_left_content">
                <span class="voted_content">
                    <span>Voted</span>
                    <span>{{votedValue === 'NaN' ? '0.00' : votedValue}}%</span>
                </span>
                <span class="yes_content">
                    <span>Yes</span>
                    <span>{{yesVotingPowerWidth === 'NaN' ? '0.00' : yesVotingPowerWidth}}%</span>
                </span>
                <span class="vote_content">
                    <span>Vote</span>
                    <span>{{vetoVotingPowerWidth === 'NaN' ? '0.00' : vetoVotingPowerWidth}}%</span>
                </span>
            </div>
            <div class="home_voting_center_content">
                <div class="home_voting_voted_participation_content">
                    <div class="blue_line_block" :style="yesVotingPowerStyleObj"></div>
                    <div class="gray_line_block" :style="abstainVotingPowerStyleObj"></div>
                    <div class="yellow_line_block" :style="noVotingPowerStyleObj"></div>
                    <div class="red_line_block" :style="vetoVotingPowerStyleObj"></div>
                </div>
                <div class="yes_pass_threshold_content">
                    <div class="yes_content_blue_link_block" :class="flShowGrayStyle ? 'show_all_blue' : ''" :style="yesVotingPowerStyleObj"></div>
                    <div class="yes_content_gray_link_block" :style="PassThresholdDefaultStyleObj"></div>
                </div>
                <div class="no_vote_threshold_content">
                    <div class="no_content_red_link_block" :class="flShowRedStyle ? 'show_all_red' : ''" :style="vetoVotingPowerStyleObj"></div>
                    <div class="no_content_gray_line_block" :style="vetoThresholdDefaultStyleObj"></div>
                </div>
            </div>
            <div class="home_voting_right_content">
                <span class="participation_content">
                    <span>{{participationValue}}%</span>
                    <span>Participation</span>
                </span>
                <span class="pass_threshold">
                    <span>{{passThresholdValue}}%</span>
                    <span>Pass Threshold</span>
                </span>
                <span class="vote_threshold">
                    <span>{{voteThresholdValue}}%</span>
                    <span>Veto Threshold</span>
                </span>
            </div>
        </div>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MHomeVotingCrad",
        props:{
	        votingObj:{
	        	type:Object
            },
	        showTitle:{
	        	type: Boolean
            }
        },
        data () {
			return {
				proposalTile: '',
                proposalId: '',
				participationValue: '',
                passThresholdValue: '',
                voteThresholdValue: '',
				hourLeft: '',
				delegatorVotedPower: '',
				validatorVotedPower: '',
				totalVotedNumber:'',
				yesVotingPowerWidth:'',
				vetoVotingPowerWidth:'',
				noVotingPowerWidth:'',
				systemVotingPower:'',
				votedValue: '',
				flShowGrayStyle: false,
                flShowParameterChange:false,
                flShowRedStyle: false,
				PassThresholdDefaultStyleObj:{
					width: ""
                },
                vetoThresholdDefaultStyleObj:{
					width: ""
                },
                passThresholdStyleObj:{
				    width:""
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
        watch:{
	        votingObj(votingObj){
	        	this.formatVotingObj(votingObj);
		        this.getVotingEndTime(votingObj.voting_end_time);
	        	this.getTotalVoted(votingObj.votTx);
            }
        },
        mounted(){
	        this.formatVotingObj(this.votingObj);
        },
        methods:{

	        formatVotingObj(votingObj){
	        	this.proposalId = votingObj.proposal_id;
	        	this.proposalTile = votingObj.title;
		        this.systemVotingPower = votingObj.voting_power_for_height;
		        this.participationValue = (Number(votingObj.level.gov_param.participation) * 100).toFixed(2);
		        this.passThresholdValue = (Number(votingObj.level.gov_param.pass_threshold) * 100).toFixed(2);
		        this.voteThresholdValue = (Number(votingObj.level.gov_param.veto_threshold) * 100).toFixed(2);
		        if(Array.isArray(votingObj.votes) && votingObj.votes.length > 0){
			        this.getTotalVoted(votingObj.votes);
			        this.getYesVotingPower(votingObj.votes);
			        this.getNoVotingPower(votingObj.votes);
			        this.getVetoVotingPower(votingObj.votes);
			        this.getAbstainVotingPower(votingObj.votes);
			        this.setYesPassThresholdStyle();
			        this.setVetoThresholdDefaultStyle();
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
		        this.totalVotedNumber = this.delegatorVotedPower + this.validatorVotedPower;
                this.votedValue = (Number(this.totalVotedNumber) / Number(this.systemVotingPower) *100).toFixed(2);
                if(Number(this.votedValue) > Number(this.participationValue)){
                    this.flShowParameterChange = true;
                }
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
		        if(Number(this.yesVotingPowerWidth) > Number(this.passThresholdValue)){
		        	this.flShowGrayStyle = true
                }
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
		        if(Number(this.vetoVotingPowerWidth) > Number(this.voteThresholdValue)){
			        this.flShowRedStyle = true
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
            setYesPassThresholdStyle(){
	        	let yesPassThresholdWidth;
	        	if(this.yesVotingPowerWidth && this.yesVotingPowerWidth !== 'NaN'){
			        yesPassThresholdWidth = ((((Number(this.passThresholdValue) / 100) * Number(this.totalVotedNumber)) / Number(this.totalVotedNumber)) * 100).toFixed(2)
		        }else {
			        yesPassThresholdWidth = this.passThresholdValue
                }
	        	this.$set(this.PassThresholdDefaultStyleObj,"width",`${yesPassThresholdWidth}%`)
            },
            setVetoThresholdDefaultStyle () {
	        	let vetoThresholdStyleWidth;
	        	if(this.vetoVotingPowerWidth && this.vetoVotingPowerWidth !== 'NaN'){
			        vetoThresholdStyleWidth = ((((Number(this.voteThresholdValue) / 100) * Number(this.totalVotedNumber)) / Number(this.totalVotedNumber)) * 100).toFixed(2);
		        }else {
			        vetoThresholdStyleWidth = this.voteThresholdValue
                }
	            this.$set(this.vetoThresholdDefaultStyleObj,"width",`${vetoThresholdStyleWidth}%`)
            }
        }
	}
</script>

<style scoped lang="scss">
    .home_voting_container{
        flex: 1;
        box-sizing: border-box;
        padding: 0.2rem;
        border: 0.01rem solid #d7d9e0;
        margin-bottom: 0.3rem;
        .home_voting_title_container_wrap{
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
        .home_voting_title_container{
            font-size: 0.12rem;
            margin-top: 0.2rem;
            i{
                margin-right: 0.1rem;
                font-size: 0.14rem;
            }
            .iconPass{
                color: #44C190;
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
        .home_voting_content{
            display: flex;
            margin-top: 0.4rem;
            .home_voting_left_content{
                display: flex;
                flex-direction: column;
                color: #22252a;
                font-size: 0.12rem;
                .voted_content{
                    min-width: 0.8rem;
                    width: auto;
                    display: flex;
                    justify-content: space-between;
                }
                .yes_content{
                    margin: 0.18rem 0;
                    min-width: 0.8rem;
                    width: auto;
                    display: flex;
                    justify-content: space-between;
                }
                .vote_content{
                    min-width: 0.8rem;
                    width: auto;
                    display: flex;
                    justify-content: space-between;
                }
            }
            .home_voting_center_content{
                flex: 1;
                margin: 0 0.1rem;
                display: flex;
                flex-direction: column;
                justify-content: space-between;
                .home_voting_voted_participation_content{
                    border-radius: 0.06rem;
                    overflow: hidden;
                    height: 0.12rem;
                    top: 0.01rem;
                    display: flex;
                    .blue_line_block{
                        height: 0.12rem;
                        background: #3598db;
                        z-index: 1;
                    }
                    .gray_line_block{
                        height: 0.12rem;
                        background: #d7d9e0;
                        z-index: 1;
                    }
                    .yellow_line_block{
                        height: 0.12rem;
                        background: #f0ad4e;
                        z-index: 1;
                    }
                    .red_line_block{
                        height: 0.12rem;
                        background: #FE8A8A;
                        z-index: 1;
                    }
                }
                .yes_pass_threshold_content{
                    border-radius: 0.06rem;
                    overflow: hidden;
                    position: relative;
                    height: 0.12rem;
                    bottom: 0.02rem;
                    .yes_content_blue_link_block{
                        height: 0.12rem;
                        background: #3598db;
                        position: absolute;
                        z-index: 1;
                        border-radius:  0 0.06rem 0.06rem 0;
                    }
                    .show_all_blue{
                        z-index: 2 !important;
                    }
                    .yes_content_gray_link_block{
                        position: absolute;
                        height: 0.12rem;
                        background: #d7d9e0;
                        z-index: 1;
                        border-radius:  0 0.06rem 0.06rem 0;
                    }
                }
                .no_vote_threshold_content{
                    width: 100%;
                    border-radius: 0.06rem;
                    overflow: hidden;
                    position: relative;
                    height: 0.12rem;
                    bottom: 0.05rem;
                    .no_content_red_link_block{
                        height: 0.12rem;
                        background: #FE8A8A;
                        position: absolute;
                        z-index: 2;
                        border-radius:  0 0.06rem 0.06rem 0;
                    }
                    .show_all_red{
                        z-index: 4 !important;
                    }
                    .no_content_gray_line_block{
                        position: absolute;
                        height: 0.12rem;
                        background: #d7d9e0;
                        z-index: 1;
                        border-radius:  0 0.06rem 0.06rem 0;
                    }
                }
            }
            .home_voting_right_content{
                display: flex;
                flex-direction: column;
                color: #22252a;
                font-size: 0.12rem;
                .participation_content{
                    min-width: 1.3rem;
                    width: auto;
                    display: flex;
                    justify-content: space-between;
                }
                .pass_threshold{
                    margin: 0.18rem 0;
                    min-width: 1.3rem;
                    width: auto;
                    display: flex;
                    justify-content: space-between;
                }
                .vote_threshold{
                    min-width: 1.3rem;
                    width: auto;
                    display: flex;
                    justify-content: space-between;
                }
            }
        }
    }
</style>
