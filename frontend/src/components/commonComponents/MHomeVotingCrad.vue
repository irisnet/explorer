<template>
    <div class="home_voting_container">
        <div class="proposal_title_content">
            <span class="proposal_title">[High Importance] IRIS Hub Upgrade - v0.15.12</span>
            <span class="view_all">
                <router-link :to="`/proposal`">View All</router-link>
            </span>
        </div>
        <div class="home_voting_title_container">
            <span><i class="iconfont iconPass"></i>ParameterChange</span>
            <span><i class="iconfont iconVeto"></i>DepositPeriod</span>
            <span><i class="iconfont iconHoursLeft"></i>4 Hours Left</span>
        </div>
        <div class="home_voting_content">
            <div class="home_voting_left_content">
                <span class="voted_content">
                    <span>Voted</span>
                    <span>66.6%</span>
                </span>
                <span class="yes_content">
                    <span>Yes</span>
                    <span>66.6%</span>
                </span>
                <span class="vote_content">
                    <span>Vote</span>
                    <span>66.6%</span>
                </span>
            </div>
            <div class="home_voting_center_content">
                <div class="home_voting_voted_participation_content">
                    <div class="blue_line_block"></div>
                    <div class="gray_line_block"></div>
                    <div class="yellow_line_block"></div>
                    <div class="red_line_block"></div>
                </div>
                <div class="yes_pass_threshold_content">
                    <div class="yes_content_blue_link_block"></div>
                    <div class="yes_content_gray_link_block"></div>
                </div>
                <div class="no_vote_threshold_content">
                    <div class="no_content_red_link_block"></div>
                    <div class="no_content_gray_line_block"></div>
                </div>
            </div>
            <div class="home_voting_right_content">
                <span class="participation_content">
                    <span>88.8%</span>
                    <span>Participation</span>
                </span>
                <span class="pass_threshold">
                    <span>88.8%</span>
                    <span>Pass Threshold</span>
                </span>
                <span class="vote_threshold">
                    <span>88.8%</span>
                    <span>Veto Threshold</span>
                </span>
            </div>
        </div>
    </div>
</template>

<script>
	export default {
		name: "MHomeVotingCrad"
	}
</script>

<style scoped lang="scss">
    .home_voting_container{
        flex: 1;
        width: 50%;
        box-sizing: border-box;
        padding: 0.2rem;
        border: 0.01rem solid #d7d9e0;
        .proposal_title_content{
            font-size: 0.16rem;
            color: #0580D3;
            line-height: 0.18rem;
            display: flex;
            justify-content:space-between;
            span{
                line-height: 1;
                a{
                    color: #0580D3!important;
                    font-size: 0.16rem;
                }
            }
            span:last-child{
                border-bottom: 0.01rem solid #0580d3;
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
                    width: 100%;
                    border-radius: 0.06rem;
                    overflow: hidden;
                    position: relative;
                    height: 0.12rem;
                    top: 0.01rem;
                    .blue_line_block{
                        width: 100%;
                        height: 0.12rem;
                        background: #3598db;
                        position: absolute;
                        z-index: 1;
                    }
                    .gray_line_block{
                        position: absolute;
                        width: 70%;
                        height: 0.12rem;
                        background: #d7d9e0;
                        z-index: 2;
                    }
                    .yellow_line_block{
                        position: absolute;
                        width: 40%;
                        height: 0.12rem;
                        background: #f0ad4e;
                        z-index: 3;
                    }
                    .red_line_block{
                        position: absolute;
                        width: 10%;
                        height: 0.12rem;
                        background: #ff0000;
                        z-index: 4;
                    }
                }
                .yes_pass_threshold_content{
                    width: 100%;
                    border-radius: 0.06rem;
                    overflow: hidden;
                    position: relative;
                    height: 0.12rem;
                    bottom: 0.02rem;
                    .yes_content_blue_link_block{
                        width: 100%;
                        height: 0.12rem;
                        background: #3598db;
                        position: absolute;
                        z-index: 1;
                    }
                    .yes_content_gray_link_block{
                        position: absolute;
                        width: 70%;
                        height: 0.12rem;
                        background: #d7d9e0;
                        z-index: 2;
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
                        width: 100%;
                        height: 0.12rem;
                        background: #3598db;
                        position: absolute;
                        z-index: 1;
                    }
                    .no_content_gray_line_block{
                        position: absolute;
                        width: 70%;
                        height: 0.12rem;
                        background: #d7d9e0;
                        z-index: 2;
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
