<template>
    <div class="proposals_list_page_wrap">
        <page-title :title="pageTitle"
                    :content="contentDoc"
                    :number="count"
                    :reversal="false"></page-title>
        <div class="graph_containers">
            <div class="graph_container mobile_graph_container" v-if="$store.state.isMobile && (depositPeriodDatas.length > 0 || votingPeriodDatas.length > 0)">
                <div v-for="v in votingPeriodDatas" :key="v.proposal_id">
                    <m-proposals-echart :data="v" v-if="v"></m-proposals-echart>
                </div>
                <div v-for="v in depositPeriodDatas" :key="v.proposal_id">
                    <m-proposals-card :data="v" v-if="v"></m-proposals-card>
                </div>
            </div>

            <div class="graph_container votingPeriodDatas_one"
                 v-if="!$store.state.isMobile && (votingPeriodDatas.length === 1 && (depositPeriodDatas.length === 0 || depositPeriodDatas.length > 1))">
                <div v-for="v in votingPeriodDatas" :key="v.proposal_id">
                    <m-proposals-echart :data="v" v-if="v"></m-proposals-echart>
                </div>
                <div v-for="v in depositPeriodDatas" :key="v.proposal_id">
                    <m-proposals-card :data="v" v-if="v"></m-proposals-card>
                </div>
            </div>

            <div class="graph_container depositPeriodDatas_one"
                 v-if="!$store.state.isMobile && (votingPeriodDatas.length === 0 && depositPeriodDatas.length === 1)">
                <div v-for="v in depositPeriodDatas" :key="v.proposal_id">
                    <m-proposals-card :data="v" v-if="v"></m-proposals-card>
                </div>
            </div>

            <div class="graph_container votingPeriodDatas_one_depositPeriodDatas_one"
                 v-if="!$store.state.isMobile && (votingPeriodDatas.length === 1 && depositPeriodDatas.length === 1)">
                <div v-for="v in votingPeriodDatas" :key="v.proposal_id">
                    <m-proposals-echart :data="v" v-if="v"></m-proposals-echart>
                </div>
                <div v-for="v in depositPeriodDatas" :key="v.proposal_id">
                    <m-proposals-card :data="v" v-if="v"></m-proposals-card>
                </div>
            </div>

            <div class="graph_container votingPeriodDatas_depositPeriodDatas"
                 v-if="!$store.state.isMobile && votingPeriodDatas.length !== 1 && (depositPeriodDatas.length > 1 || votingPeriodDatas.length > 1)">
                <div>
                    <div v-for="v in votingPeriodDatas" :key="v.proposal_id">
                        <m-proposals-echart :data="v" v-if="v"></m-proposals-echart>
                    </div>
                </div>
                <div>
                    <div v-for="v in depositPeriodDatas" :key="v.proposal_id">
                        <m-proposals-card :data="v" v-if="v"></m-proposals-card>
                    </div>
                </div>
            </div>
        </div>


        <div :class="proposalsListPageWrap" :style="{'padding-top': depositPeriodDatas.length === 0 && votingPeriodDatas.length === 0 ? '0.54rem' : 0}">
            <div class="pagination total_num header_fixed_style" :style="{'position':flTableFixed ? 'static' : 'fixed'}" :class="[$store.state.isMobile ? 'mobile_graph_pagination_container' : '']">
                <div style="height: 70px; display: flex; align-items: center;">
                    <!--<span class="proposals_list_page_wrap_hash_var" :class="count ? 'count_show' : 'count_hidden' ">{{count}} Proposals</span>-->
                    <div class="icon_list">
                        <div>
                            <img src="../../assets/critical.png" style="margin-left: 0;" />
                            <span>Critical</span>
                        </div>
                        <div>
                            <img src="../../assets/important.png" />
                            <span>Important</span>
                        </div>
                        <div>
                            <img src="../../assets/normal.png" />
                            <span>Normal</span>
                        </div>
                    </div>
                </div>
                <div class="mobile_graph_pagination_last_node">
                    <div class="icon_list">
                        <div>
                            <i style="margin-left: 0;"></i>
                            <span>Yes</span>
                        </div>
                        <div>
                            <i style="background-color: #CCDCFF;"></i>
                            <span>Abstain</span>
                        </div>
                        <div>
                            <i style="background-color: #FFCF65;"></i>
                            <span>No</span>
                        </div>
                        <div>
                            <i style="background-color: #FE8A8A;"></i>
                            <span>NoWithVeto</span>
                        </div>
                    </div>
                </div>
            </div>
            <div class="mobile_style" style="overflow-x: auto;-webkit-overflow-scrolling:touch;" :style="{'padding-top':flTableFixed ? '' : '0.7rem'}">
                <m-proposals-list-table :items="items" :showFixedHeader="flTableFixed"></m-proposals-list-table>
                <div v-show="showNoData" class="no_data_show">
                    <img src="../../assets/no_data.svg" alt="">
                </div>
            </div>
            <div class="pagination" style='margin:0.2rem 0 0.4rem;'>
                <m-pagination :page-size="pageSize" :total="count" :page="currentPageNum"
                              :page-change="pageChange">
                </m-pagination>
            </div>
        </div>
    </div>
</template>

<script>
	import Tools from '../../util/Tools';
	import Service from "../../service";
	import MProposalsCard from '../commontables/MProposalsCard';
	import MProposalsEchart from '../commontables/MProposalsEchart';
	import MProposalsListTable from './MProposalsListTable';
	import MPagination from "../commontables/MPagination";
	import Constant from "../../constant/Constant";
	import skinStyle from "../../skinStyle/"
    import PageTitle from "../pageTitle/PageTitle";
    import pageTitleConfig from "../pageTitle/pageTitleConfig";
	export default {
		components:{
            PageTitle,
			MProposalsCard,
			MProposalsEchart,
			MProposalsListTable,
			MPagination
		},
		watch: {
			$route() {
				this.items = [];
				this.type = this.$route.params.type;
				this.currentPageNum = Number(this.$route.query.page || 1);
				this.getDataList();
				this.showNoData = false;
				this.computeMinWidth();
			}
		},
		data() {
			return {
                pageTitle:pageTitleConfig.GovProposals,
                contentDoc:'Proposals',
				devicesWidth: window.innerWidth,
				proposalsListPageWrap: 'personal_computer_proposals_list_page',
				currentPageNum: this.forCurrentPageNum(),
				currentPageNumCache: 0,
				pageSize: 30,
				count: 0,
				items: [],
				showNoData:false,
				innerWidth : window.innerWidth,
				tableMinWidth:"",
				votingPeriodDatas: [],
				depositPeriodDatas: [],
				flTableFixed:'',
			}
		},
		beforeMount() {
			this.computeMinWidth();
			if (window.innerWidth > 910) {
				this.proposalsListPageWrap = 'personal_computer_proposals_list_page_wrap';
			} else {
				this.proposalsListPageWrap = 'mobile_proposals_list_page_wrap';
			}
		},
		mounted() {
			this.getDataList();
			this.getGrahpData();
			window.addEventListener('resize',this.onresize);
		},
		beforeDestroy() {
			window.removeEventListener('resize',this.onWindowResize);
		},
		methods: {
			forCurrentPageNum() {
				let currentPageNum = 1;
				let urlPageSize = this.$route.query.page && Number(this.$route.query.page);
				currentPageNum = urlPageSize ? urlPageSize : 1;
				return currentPageNum;
			},
			pageChange(pageNum) {
				this.currentPageNum = pageNum;
				if (this.currentPageNumCache === this.currentPageNum) {
					return;
				}
				this.currentPageNumCache = this.currentPageNum;
                let path = "/gov/proposals";
                history.pushState(null, null, `/#${path}?page=${pageNum}`);
                this.getDataList()
			},
			onresize(){
				this.innerWidth = window.innerWidth;
				if(window.innerWidth > 910){
					this.proposalsListPageWrap = 'personal_computer_proposals_list_page_wrap';
				} else {
					this.proposalsListPageWrap = 'mobile_proposals_list_page_wrap';
				}
			},
			//根绝页面的不同展示最小宽度,不换行显示列表
			computeMinWidth(){
				if(this.$route.path === '/Proposals'){
					this.tableMinWidth = 8.8;
				}
			},
			formatGrahpChildren(arr, color) {
				let hStep = ((color.h[1] - color.h[0]) / 100);
				let sStep = ((color.s[1] - color.s[0]) / 100);
				let lStep = ((color.l[1] - color.l[0]) / 100);
				return arr.map((v, i) => {
					let h = color.h[0] + hStep * i;
					let s = color.s[0] + sStep * i;
					let l = color.l[0] + lStep * i;
					let obj =  {
						value: v.voting_power,
						info: v,
						nodeClick: false,
						itemStyle: {
							color: `hsla(${h},${s}%,${l}%, 1)`,
							borderColor: '#ECEFFF',
							borderWidth: 0
						}
					}
					return obj;
				});
			},
			formatNumber(str) {
				if (!str) {
					return '';
				}
				let r = /\.(\d{5})/g;
				let arr = str.match(r);
				if (arr && arr[0]) {
					let n = Math.round(Number(arr[0].replace('.', '')) / 10);
					n = n / 100;
					return n;
				} else {
					return '';
				}
			},
			formatGrahpData(data) {
				this.flTableFixed =  true;
				let votingPeriodDatas = data.filter(v => v.status === 'VotingPeriod');
				let depositPeriodDatas = data.filter(v => v.status === 'DepositPeriod');
				this.votingPeriodDatas = votingPeriodDatas.map(item => {
					let o = {};
					o.proposal_id = item.proposal_id;
					o.title = item.title;
					o.level = item.level && item.level.name;
					o.type = item.type;
					o.status = item.status;
					o.votingEndTime = item.voting_end_time;
					let all = item.voting_power_for_height;
					let yesArr = item.votes.filter(v => v.option === 'Yes');
					let yes = yesArr.reduce((init, v) => {return v.voting_power + init}, 0);
					let noArr = item.votes.filter(v => v.option === 'No');
					let no = noArr.reduce((init, v) => {return v.voting_power + init}, 0);
					let abstainArr = item.votes.filter(v => v.option === 'Abstain');
					let abstain = abstainArr.reduce((init, v) => {return v.voting_power + init}, 0);
					let noWithVetoArr = item.votes.filter(v => v.option === 'NoWithVeto');
					let noWithVeto = noWithVetoArr.reduce((init, v) => {return v.voting_power + init}, 0);
					let votes = yes + no + abstain + noWithVeto;
					o.participationNum = item.level && item.level.gov_param && item.level.gov_param.participation && this.formatNumber(item.level.gov_param.participation);
					o.passThresholdNum = item.level && item.level.gov_param && item.level.gov_param.pass_threshold && this.formatNumber(item.level.gov_param.pass_threshold);
					o.vetoThresholdNum = item.level && item.level.gov_param && item.level.gov_param.veto_threshold && this.formatNumber(item.level.gov_param.veto_threshold);
					o.participation = all ? (votes / all) * 100 : 0;
					o.passThreshold = votes ? (yes / votes) * 100 : 0;
					o.vetoThreshold = votes ? (noWithVeto / votes) * 100 : 0;
					let nonparticipantPer = Tools.formatDecimalNumberToFixedNumber((all - votes) / all * 100);
					let data = [
						{
							name: 'Participant',
							value: votes,
							perData: Tools.formatDecimalNumberToFixedNumber(votes / all * 100),
							itemStyle: {
								// color: '#3598DB',
								borderColor: '#ECEFFF',
								borderWidth: 0
							},
							children: [
								{
									name: 'Yes',
									value: yes,
									perData: Tools.formatDecimalNumberToFixedNumber(yes / votes * 100),
									itemStyle: {
										// color: '#45B4FF',
										borderColor: '#ECEFFF',
										borderWidth: 0
									},
									// children: this.formatGrahpChildren(yesArr, {h: [205, 204], s: [100, 100], l: [79, 35]})
								},
								{
									name: 'Abstain',
									value: abstain,
									perData: Tools.formatDecimalNumberToFixedNumber(abstain / votes * 100),
									itemStyle: {
										color: '#CCDCFF',
										borderColor: '#ECEFFF',
										borderWidth: 0
									},
									children: this.formatGrahpChildren(abstainArr, {h: [222, 221], s: [100, 44], l: [86, 58]})
								},
								{
									name: 'No',
									value: no,
									perData: Tools.formatDecimalNumberToFixedNumber(no / votes * 100),
									itemStyle: {
										color: '#FFCF65',
										borderColor: '#ECEFFF',
										borderWidth: 0
									},
									children: this.formatGrahpChildren(noArr, {h: [36, 36], s: [100, 100], l: [77, 48]})
								},
								{
									name: 'NoWithVeto',
									value: noWithVeto,
									perData: Tools.formatDecimalNumberToFixedNumber(noWithVeto / votes * 100),
									itemStyle: {
										color: '#FE8A8A',
										borderColor: '#ECEFFF',
										borderWidth: 0
									},
									children: this.formatGrahpChildren(noWithVetoArr, {h: [21, 21], s: [100, 100], l: [79, 50]})
								}
							]
						},
						{
							name: 'Nonparticipant',
							value: all - votes,
							perData: nonparticipantPer,
							nodeClick: false,
							itemStyle: {
								color: '#E5E9FB',
								borderColor: '#E5E9FB',
								borderWidth: 0
							},
							label: {
								color: '#51A9FF',
								textBorderWidth: 0,
								fontWeight: 600
							},
							children: [
								{
									name: '',
									tipName: 'Nonparticipant',
									value: all - votes,
									perData: nonparticipantPer,
									nodeClick: false,
									itemStyle: {
										color: '#E5E9FB',
										borderColor: '#E5E9FB',
										borderWidth: 0
									},
									children: [
										{
											name: '',
											tipName: 'Nonparticipant',
											value: all - votes,
											perData: nonparticipantPer,
											nodeClick: false,
											itemStyle: {
												color: '#E5E9FB',
												borderColor: '#E5E9FB',
												borderWidth: 0
											}
										}
									]
								}
							]
						}
					];
                    if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.IRISHUB){
                        data[0].itemStyle.color = skinStyle.skinStyle.MAINNETBGCOLOR;
                        data[0].children[0].itemStyle.color = '#4371FF';
                        data[0].children[0].children = this.formatGrahpChildren(yesArr, {h: [223, 222], s: [100, 100], l: [75, 35]})
                    }else if(this.$store.state.currentSkinStyle === Constant.CHAINID.FUXI){
                        data[0].children[0].itemStyle.color = '#004EAA';
                        data[0].itemStyle.color = skinStyle.skinStyle.TESTNETBGCOLOR;
                        data[0].children[0].children = this.formatGrahpChildren(yesArr, {h: [213, 212], s: [100, 100], l: [75, 35]})
                    }else if(this.$store.state.currentSkinStyle === Constant.CHAINID.NYANCAT){
                        data[0].children[0].itemStyle.color = '#06A79A';
                        data[0].itemStyle.color = skinStyle.skinStyle.NYANCATTESTNETBGCOLOR;
                        data[0].children[0].children = this.formatGrahpChildren(yesArr, {h: [175, 174], s: [100, 100], l: [75, 35]})
                    }else {
                        data[0].children[0].itemStyle.color = '#008CEA';
                        data[0].itemStyle.color = skinStyle.skinStyle.DEFAULTBGCOLOR;
                        data[0].children[0].children = this.formatGrahpChildren(yesArr, {h: [196, 195], s: [100, 100], l: [75, 35]})
                    }
					o.data = data;
					return o;
				});
				this.votingPeriodDatas = this.votingPeriodDatas.sort((a,b) =>{
					return b.proposal_id - a.proposal_id
                });
				depositPeriodDatas.forEach(v => {
					if (v.level && v.level.gov_param && v.level.gov_param.min_deposit && (typeof v.level.gov_param.min_deposit.amount === 'number')) {
						v.min_deposit_number = Number(v.level.gov_param.min_deposit.amount);
						let n = v.min_deposit_number === 0 ? v.min_deposit_number : Tools.formatAmount(v.level.gov_param.min_deposit.amount);
						v.min_deposit_format = `${n} IRIS`;
					}
					if (v.intial_deposit && (typeof v.intial_deposit.amount === 'number')) {
						v.intial_deposit_number = Number(v.intial_deposit.amount);
						let n = v.intial_deposit_number === 0 ? v.intial_deposit_number : Tools.formatAmount(v.intial_deposit.amount);
						v.intial_deposit_format = `${n} IRIS`;
					}
					if (v.total_deposit && (typeof v.total_deposit.amount === 'number')) {
						v.total_deposit_number = Number(v.total_deposit.amount);
						let n = v.total_deposit_number === 0 ? v.total_deposit_number : Tools.formatAmount(v.total_deposit.amount);
						v.total_deposit_format = `${n} IRIS`;
					}
					v.intial_deposit_number_per = this.isNumber(v.intial_deposit_number) && this.isNumber(v.min_deposit_number) ?
            (this.forLimitNumer(v.intial_deposit_number / v.min_deposit_number)) * 100 + '%' : 0;
					v.total_deposit_number_per = this.isNumber(v.total_deposit_number) && this.isNumber(v.min_deposit_number) ?
						(this.forLimitNumer(v.total_deposit_number / v.min_deposit_number))* 100 + '%' : 0;
					v.level = v.level && v.level.name;
				});
				this.depositPeriodDatas = depositPeriodDatas.sort((a,b) => {
					return b.proposal_id - a.proposal_id
                });
			},
			isNumber(n) {
				return typeof n === 'number'
            },
            forLimitNumer(number) {
                if (typeof number === "number") {
                  return Math.max(Math.min(number, 1), 0)
                }
            },
			getGrahpData() {
				Service.commonInterface({proposalListVotingAndDeposit:{}}, (data) => {
					if (data && Array.isArray(data) && data.length > 0) {
					    this.totalNumber = data.length;
						this.formatGrahpData(data);
					}else {
						this.flTableFixed =  false;
                    }
				});
			},
			getDataList() {
				Service.commonInterface({proposalList:{
						pageNumber: this.currentPageNum,
						pageSize: this.pageSize
					}}, (proposalList) => {
					try {
                        this.count = proposalList.Count;
						if(proposalList.Data){
							this.showNoData = false;
							this.items = proposalList.Data.map(item =>{
								let proposalId = item.proposal_id === 0 ? "--" : item.proposal_id;
								let type = item.type;
								let status  = item.status;
								let submitTime = (new Date(item.submit_time).getTime()) > 0 ? Tools.format2UTC(item.submit_time) : '--';
								let depositEndTime = (new Date(item.deposit_end_time).getTime()) > 0 ? Tools.format2UTC(item.deposit_end_time) : '--';
								let votingEndTime = (new Date(item.voting_end_time).getTime()) > 0 ? Tools.format2UTC(item.voting_end_time) : '--';
								let title = Tools.formatString(item.title,20,"...");
								let final_votes = {};
								let finalTotalVotes = 0;
								if (status === 'VotingPeriod') {
									let yesArr = item.votes.filter(v => v.option === 'Yes');
									let yes = yesArr.reduce((init, v) => {return v.voting_power + init}, 0);
									let noArr = item.votes.filter(v => v.option === 'No');
									let no = noArr.reduce((init, v) => {return v.voting_power + init}, 0);
									let abstainArr = item.votes.filter(v => v.option === 'Abstain');
									let abstain = abstainArr.reduce((init, v) => {return v.voting_power + init}, 0);
									let noWithVetoArr = item.votes.filter(v => v.option === 'NoWithVeto');
									let no_with_veto = noWithVetoArr.reduce((init, v) => {return v.voting_power + init}, 0);
									finalTotalVotes = yes + no + abstain + no_with_veto;
									final_votes.yes = (yes / finalTotalVotes) * 100;
									final_votes.no = (no / finalTotalVotes) * 100;
									final_votes.abstain = (abstain / finalTotalVotes) * 100;
									final_votes.no_with_veto = (no_with_veto / finalTotalVotes) * 100;
								} else {
									final_votes = Object.keys(item.final_votes).length > 0 ? item.final_votes : final_votes;
									for (let k in item.final_votes) {
										finalTotalVotes += Number(item.final_votes[k]);
									}
									for (let k in final_votes) {
										final_votes[k] = (Number(final_votes[k]) / finalTotalVotes) * 100;
									}
								}
								return {
									title : title,
									id : proposalId,
									type : type,
									status : status,
									submitTime : submitTime,
									depositEndTime: depositEndTime,
									votingEndTime: votingEndTime,
									finalTotalVotes: finalTotalVotes,
									finalVotes: final_votes,
									level: item.level && item.level.name
								}
							});
						}else {
							this.items = [];
							this.showNoData = true;
						}
					}catch (e) {
						console.error(e);
						this.items = [];
						this.showNoData = true;
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
    @import '../../style/mixin';

    .proposals_list_page_wrap {
        @include flex;
        @include pcContainer;
        font-size: 0.14rem;
        .pagination {
            @include flex;
            justify-content: flex-end;
            @include borderRadius(0.025rem);
            .icon_list {
                display: flex;
                div {
                    font-size: 14px;
                    color: var(--contentColor);
                    display: flex;
                    align-items: center;
                    i {
                        width: 12px;
                        height: 12px;
                        border-radius: 2px;
                        display: inline-block;
                        margin-left: 50px;
                        background-color: var(--bgColor);
                        vertical-align: middle;
                    }
                    img {
                        width: 14px;
                        height: 14px;
                        margin-left: 50px;
                        vertical-align: middle;
                    }
                    span {
                        margin-left: 10px;
                        vertical-align: middle;
                    }
                }
            }
            li{
                height:0.3rem !important;
                a{
                    box-shadow: none;
                }
                a:focus{
                    -webkit-box-shadow:0 0 0 .2rem rgba(255,255,255,.5);
                    box-shadow:0 0 0 .2rem rgba(255,255,255,.5)
                }
            }
        }
        .header_fixed_style{
            z-index: 1 !important;
        }
        @media screen and (max-width: 910px){
            .header_fixed_style{
                position: static !important;
            }
        }
        .total_num{
            @include flex;
            justify-content: space-between;
            align-items: center;
            flex-wrap: wrap;
            width: 100%;
            max-width: 12.8rem;
            z-index: 10 !important;
            background: #F5F7FD;
            & > div {
                padding: 4px 0;
            }
        }
        .fixed_style{
            position: fixed;
        }
        .no_data_show{
            @include flex;
            justify-content: center;
            border-top:0.01rem solid #eee;
            border-bottom:0.01rem solid #eee;
            font-size:0.14rem;
            height:3rem;
            align-items: center;
        }
        .b-table {
            a {
                text-decoration: none;
            }
        }
        .proposals_list_title_wrap {
            width: 100%;
            border-bottom: 1px solid #d6d9e0 !important;
            position: relative;
            z-index: 0;
            @include flex;
            @include pcContainer;
            height:0.62rem;
            min-height: 0 !important;
            background:#efeff1;
            p{
                height:0.62rem;
                span{
                    height:0.62rem;
                    line-height:0.62rem;
                }
            }
            .personal_computer_proposals_list_page_wrap {
                @include flex;
            }
            .mobile_proposals_list_page_wrap {
                @include flex;
                flex-direction: column;
                overflow-x: auto;
                -webkit-overflow-scrolling:touch;
                width:100%;
                box-sizing: border-box;
                .proposals_list_page_wrap_hash_var{
                    min-width:7rem;
                    color: #515a6e;
                    font-weight: bold;
                }
            }
        }
        .personal_computer_proposals_list_page_wrap {
            width: 100%!important;
            padding-top: 0.54rem;
            .transaction_information_content_title {
                height: 0.4rem;
                line-height: 0.4rem;
                font-size: 0.18rem;
                color: #515a6e;
                margin-bottom: 0;
                font-weight: bold;
            }
            @include pcCenter;
            .transactions_detail_information_wrap {
                .information_props_wrap {
                    @include flex;
                    .information_props {
                        width: 15rem;
                    }
                }
            }

            .proposals_list_title {
                height: 0.62rem;
                line-height: 0.62rem;
                font-size: 0.22rem;
                color: #000000;
                margin-right: 0.2rem;
                @include fontWeight;
                margin-left: 0.2rem;
            }
            .proposals_list_page_wrap_hash_var {
                height:  0.62rem;
                line-height: 0.62rem;
                font-size: 0.18rem;
                color: #515a6e;
                font-weight: bold;
                margin-left: 0.2rem;
            }
            .for_proposals{
                display:inline-block;
                margin-left:0.1rem;
            }
        }

        .mobile_proposals_list_page_wrap {
            width: 100%;
            @include flex;
            flex-direction: column;
            padding: 0 0.1rem;
            box-sizing: border-box;
            .transaction_information_content_title {
                height: 0.4rem;
                line-height: 0.4rem;
                font-size: 0.18rem;
                color: #000000;
                margin-bottom: 0;
                @include fontWeight;
            }
            .transactions_detail_information_wrap {

                .information_props_wrap {
                    @include flex;
                    flex-direction: column;
                    border-bottom: 0.01rem solid #eee;
                    margin-bottom: 0.05rem;
                    .information_value {
                        overflow-x: auto;
                        -webkit-overflow-scrolling:touch;
                    }
                }
            }
            .pagination{
                padding-left: 0.2rem;
            }
            .proposals_list_title {
                height: 0.6rem;
                line-height: 0.6rem;
                font-size: 0.18rem;
                color: #000000;
                margin-right: 0.2rem;
                padding-left: 0.1rem;
                @include fontWeight;
            }
            .proposals_list_page_wrap_hash_var {
                overflow-x: auto;
                -webkit-overflow-scrolling:touch;
                height: 0.3rem;
                line-height: 0.3rem;
                font-size: 0.18rem;
                color:#515a6e;
                font-weight: bold;
            }
            .for_proposals{
                display:inline-block;
            }
        }
    }

    //重置bootstrap-vue的表格样式
    table{
        td{
            max-width:2.2rem !important;
            overflow-wrap: break-word !important;
            word-wrap: break-word !important;
        }
    }
    .page-item{
        &:first-child, &:last-child{
            .page-link{
                @include borderRadius(0.025rem);
            }
        }
    }
    .count_show{
        display: block;
        margin-right: 50px;
    }
    .count_hidden{
        display: none !important;
    }
    .graph_containers {
        width: 100%;
        position: relative;
        z-index: 1;
        padding-top: 10rem;
        margin-top: -10rem;
        overflow-x: auto;
    }
    .graph_container {
        display: flex;
        width: 12.8rem;
        flex-wrap: wrap;
        margin: 0.74rem auto 0.1rem;
        &:nth-last-of-type(1) {
            margin-bottom: 0;
        }
    }
    .votingPeriodDatas_one {
        justify-content: space-between;
        & > div {
            margin-top: 0.2rem !important;
            width: calc(50% - 0.1rem);
            .propsals_card_container {
                height: 2.5rem;
            }
        }
        & > div:nth-child(1) {
            width: 100%;
            margin-top: 0rem !important;
            .propsals_echart_container {
                width: 100%;
                background: #fff;
            }
        }
    }
    .depositPeriodDatas_one {
        height: 2.2rem;
        & > div {
            width: 100%;
            display: flex;
            & > div {
                flex: 1;
            }
        }
    }
    .votingPeriodDatas_one_depositPeriodDatas_one {
        justify-content: space-between;
        & > div {
            width: calc(50% - 0.1rem);
        }
    }
    .votingPeriodDatas_depositPeriodDatas {
        flex-direction: column;
        display: flex;
        margin-top: 0.54rem;
        & > div {
            flex-direction: row;
            display: flex;
            justify-content: space-between;
            flex-wrap: wrap;
            & > div {
                width: calc(50% - 0.1rem);
                margin-top: 0.2rem;
                .propsals_card_container {
                    height: 2.5rem;
                }
            }
        }
    }
    .mobile_graph_container {
        display: flex;
        flex-direction: column;
        width: 100%;
        & > div:nth-child(1) {
            margin-top: 0!important;
        }
        & > div {
            margin-left: 0.1rem!important;
            margin-right: 0.1rem!important;
            margin-top: 0.1rem!important;
            flex: 1;
            display: flex;
            width: calc(100% - 0.2rem);
            & > div {
                width: 100%!important;
                margin: auto;
            }
            & > div.propsals_card_container {
                height: 2.5rem;
            }
        }
        .propsals_echart_container {
            flex-direction: column;
            height: 7rem;
        }
    }
    .mobile_graph_pagination_container {
        padding-left: 0.1rem !important;
        & > div:nth-child(1) {
            height: auto!important;
        }
        & > div {
            flex-wrap: wrap;
            width: 100%;
            .proposals_list_page_wrap_hash_var {
                white-space: nowrap;
                display: block;
                width: 100%;
            }
            padding: 10px 0!important;
            .icon_list {
                div {
                    img, i {
                        width: 14px !important;
                        height: 14px !important;
                        margin-left: 10px!important;
                    }
                    &:first-child {
                        img, i {
                            margin-left: 0px!important;
                        }
                    }
                }
            }
        }
        .mobile_graph_pagination_last_node {
            display: flex;
            padding: 0 0 10px 0!important;
        }
    }
    @media screen and (max-width: 910px){
        .mobile_style{
            padding-top: 0 !important;
        }
    }
</style>

<style lang="scss">
    .votingPeriodDatas_one {
        & > div:nth-child(1) {
            .propsals_echart_container {
                .content {
                    .content_div {
                        position: absolute;
                    }
                }
            }
        }
    }
</style>
