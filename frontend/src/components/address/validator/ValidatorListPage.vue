<template>
    <div class="blocks_list_page_wrap">
        <page-title :title="pageTitle"
                    :content="`${count} ${validatorPageTitle}`"></page-title>
        <div class="blocks_list_title_wrap" :class="blocksListPageWrap === 'personal_computer_blocks_list_page_wrap' ? 'fixed_style' :''">
            <div :class="blocksListPageWrap" style="margin-bottom:0;">
                <div class="validators_status_tab">
                    <div style="height: 0.3rem; margin-top: 0.2rem;">
                        <m-tabs :data="validatorStatusTitleList" :chose="selectValidatorStatus"></m-tabs>
                    </div>
                </div>
            </div>
        </div>

        <div :class="blocksListPageWrap" :style="{'margin-top':`${blocksListPageWrap === 'personal_computer_blocks_list_page_wrap' ? '1.24rem' : '0'}`}">
            <div style="overflow-x: auto; overflow-y: hidden; -webkit-overflow-scrolling:touch;background: #fff">
                <!-- <validator-list-table :items="items" :minWidth="tableMinWidth" :showNoData="showNoData"></validator-list-table> -->
                <m-validator-list-table ref="mtable"
                                        :items="items"
                                        :minWidth="tableMinWidth"
                                        :showNoData="showNoData"></m-validator-list-table>
                <div v-show="showNoData" class="no_data_show" style="background: #fff;">
                    <img src="../../../assets/no_data.svg" alt="">
                </div>
            </div>
            <div class="pagination total_num" style='margin-bottom:0.2rem;justify-content: flex-end'>
                <m-pagination :page-size="pageSize" :total="count"></m-pagination>
            </div>
        </div>
    </div>
</template>

<script>
	import Tools from '../../../util/Tools';
	import Constant from "../../../constant/Constant"
	import Service from "../../../service"
	import ValidatorListTable from "./ValidatorListTable";
    import MValidatorListTable from "./MValidatorListTable";
    import MTabs from "../../commontables/MTabs";
	import MPagination from "../../commontables/MPagination";
    import PageTitle from "../../pageTitle/PageTitle";
    import pageTitleConfig from "../../pageTitle/pageTitleConfig";

	export default {
		components:{
            PageTitle,
			MPagination,
			ValidatorListTable,
            MValidatorListTable,
            MTabs
		},
		watch: {
			currentPage(currentPage) {
				this.currentPage = currentPage;
				new Promise((resolve)=>{
					this.getDataList(currentPage, 30, this.$route.params.type);
					resolve();
				}).then(()=>{
					document.getElementById('router_wrap').scrollTop = 0;
				})
			},
			$route() {
				this.items = [];
				this.currentPage = 1;
				this.showNoData = false;
				this.computeMinWidth();
				this.getDataList();
			}
		},
		data() {
			return {
			    pageTitle:pageTitleConfig.StakingValidators,
                validatorPageTitle:'',
				devicesWidth: window.innerWidth,
				blocksListPageWrap: 'personal_computer_blocks_list_page',
				blocksValue: '',
				currentPage: 1,
				pageSize: 30,
				validatorPageSize: 100,
				defaultValidatorPageNumber:1,
				count: 0,
				fields: [],
				items: [],
				type: '1',
				showNoData:false,//是否显示列表的无数据
				innerWidth : window.innerWidth,
				tableMinWidth: 0,
				listTitleName:"",
				validatorTabIndex: localStorage.getItem('validatorTabIndex') ? localStorage.getItem('validatorTabIndex') : 0,
				validatorStatusTitleList:[
					{
						title:'Active',
						isActive: true,
					},
					{
						title:'Candidate',
						isActive: false,
					},
					{
						title:'Jailed',
						isActive: false,
					}
				]
			}
		},
		beforeMount() {
			this.type = this.$route.params.type;
			if (window.innerWidth > 910) {
				this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
			} else {
				this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
			}
		},
		mounted() {
			this.getDataList(this.$route.query.page ? this.$route.query.page : this.defaultValidatorPageNumber, this.pageSize, this.$route.params.type);
			window.addEventListener('resize',this.onresize);
			this.computeMinWidth();
		},
		methods: {
			selectValidatorStatus(index){
                index = index || 0;
				this.validatorStatusTitleList.forEach( item => {
					item.isActive = false
				});
				localStorage.setItem('validatorTabIndex',index);
				this.validatorStatusTitleList[index].isActive = true;
				this.getValidatorList(this.defaultValidatorPageNumber,this.validatorPageSize,this.getValidatorStatus(index))
			},
			onresize(){
				this.innerWidth = window.innerWidth;
				if(window.innerWidth > 910){
					this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
				} else {
					this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
				}
			},
			computeMinWidth(){
				this.tableMinWidth = 12.8;
			},
			getValidatorList(currentPage, pageSize,status){
				this.pageSize = this.validatorPageSize;
				let url;
				url = `/api/stake/validators?page=${currentPage}&size=${pageSize}&type=${status}&origin=browser`;
				Service.commonInterface({validatorList:{
						pageNumber: currentPage,
						pageSize: pageSize,
						validatorStatus: status,
					}},(result) => {
					try {
						this.count = result && result.Count ? result.Count : 0;
						result = result && result.Data ? result.Data : null;
						if(result){
							this.items = result.map((item) => {
                                let regex =  /[^\w\u4e00-\u9fa50-9a-zA-Z]/g;
                                let replaceMoniker = item.description.moniker.replace(regex,'');
                                let validatorIconSrc = replaceMoniker ? Tools.firstWordUpperCase(replaceMoniker.match(/^[0-9a-zA-Z\u4E00-\u9FA5]/g)[0]) : '';
								return {
									validatorStatus: status,
									moniker: Tools.formatString(item.description.moniker,15,'...'),
									operatorAddress: item.operator_address,
									commission: `${(item.commission.rate * 100).toFixed(2)} %`,
									bondedToken: `${Tools.formatPriceToFixed(Number(item.tokens),2)} ${this.$store.state.displayToken.toLocaleUpperCase()}`,
									uptime: Tools.FormatUptime(item.uptime),
									votingPower: `${(item.voting_rate * 100).toFixed(4)}%`,
									selfBond: `${Tools.subStrings(Tools.formatPriceToFixed(Number(item.self_bond.match(/\d*(\.\d{0,4})?/)[0])), 2)} ${this.$store.state.displayToken.toLocaleUpperCase()}`,
									delegatorNum: item.delegator_num,
									bondHeight: Number(item.bond_height),
									unbondingHeight: item.unbonding_height && Number(item.unbonding_height) > 0 ? Number(item.unbonding_height) : '--',
									unbondingTime: (new Date(item.unbonding_time).getTime()) > 0 ? Tools.format2UTC(item.unbonding_time) : '--',
                                    identity: item.description.identity,
                                    url: item.icons ? item.icons : replaceMoniker ? '' : require('../../../assets/default_validator_icon.svg'),
                                    validatorIconSrc: validatorIconSrc
								}
							});
							this.showNoData = false;
						}else {
							this.showNoData = true;
							this.items = [];
						}
						this.$refs.mtable && this.$refs.mtable.setValidatorField(status);
					}catch (e) {
						this.showNoData = true;
						this.$refs.mtable && this.$refs.mtable.setValidatorField(status);
						console.error(e)
					}
				})
			},
			getValidatorStatus(index){
				let validatorStatus;
				switch (index) {
					case 0 :
						validatorStatus = 'validator';
						this.validatorPageTitle = 'Active';
						break;
					case 1 :
						validatorStatus = 'candidate';
                        this.validatorPageTitle = 'Candidate';
						break;
					case 2 :
						validatorStatus = 'jailed'
                        this.validatorPageTitle = 'Jailed';
				}
				
				return validatorStatus
			},
			getDataList() {
				this.selectValidatorStatus(Number(this.validatorTabIndex))
			},
		},
		beforeDestroy () {
			let defaultValidatorTabIndex = 0;
			localStorage.setItem('validatorTabIndex',defaultValidatorTabIndex);
            window.removeEventListener('resize',this.onWindowResize);
		}
	}
</script>
<style lang="scss" scoped>
    @import '../../../style/mixin';
    @import '../../../style/validator_list_page';
</style>
