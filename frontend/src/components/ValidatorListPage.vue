<template>
    <div class="blocks_list_page_wrap">
        <div class="blocks_list_title_wrap" :class="blocksListPageWrap === 'personal_computer_blocks_list_page_wrap' ? 'fixed_style' :''">
            <div :class="blocksListPageWrap" style="margin-bottom:0;">
                <div class="validators_status_tab">
                    <div style="height: 0.3rem; margin-top: 0.2rem;">
                        <m-tabs :data="validatorStatusTitleList" :chose="selectValidatorStatus"></m-tabs>
                    </div>
                </div>
            </div>
        </div>
        
        <div :class="blocksListPageWrap" :style="{'margin-top':`${blocksListPageWrap === 'personal_computer_blocks_list_page_wrap' ? '0.7rem' : '0'}`}">
            <div style="overflow-x: auto; overflow-y: hidden; -webkit-overflow-scrolling:touch;">
                <!-- <validator-list-table :items="items" :minWidth="tableMinWidth" :showNoData="showNoData"></validator-list-table> -->
                <m-validator-list-table ref="mtable"
                                        :items="items"
                                        :minWidth="tableMinWidth"
                                        :showNoData="showNoData"></m-validator-list-table>
                <div v-show="showNoData" class="no_data_show">
                    No Data
                </div>
            </div>
            <div class="pagination total_num" style='margin-bottom:0.2rem;'>
                <span class="blocks_list_page_wrap_hash_var">{{count}} Total</span>
                <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
                </b-pagination>
            </div>
        </div>
        <spin-component :showLoading="showLoading"/>
    </div>
</template>

<script>
	import Tools from '../util/Tools';
	import Constant from "../constant/Constant"
	import Service from "../service"
    import Http from "../util/axios"
	import SpinComponent from './commonComponents/SpinComponent';
	import ValidatorListTable from "./table/ValidatorListTable";
    import MValidatorListTable from "./table/MValidatorListTable";
    import MTabs from "./commonComponents/MTabs";

	export default {
		components:{
			ValidatorListTable,
			SpinComponent,
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
				showLoading:false,
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
		beforeDestroy() {
			window.removeEventListener('resize',this.onWindowResize);
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
								return {
									validatorStatus: status,
									url:require('../assets/header_img.png'),
									moniker: Tools.formatString(item.description.moniker,15,'...'),
									operatorAddress: item.operator_address,
									commission: `${(item.commission.rate * 100).toFixed(2)} %`,
									bondedToken: `${Tools.formatPriceToFixed(Number(item.tokens),2)} ${Constant.CHAINNAME.toLocaleUpperCase()}`,
									uptime: `${(item.uptime * 100).toFixed(2)}%`,
									votingPower: `${(item.voting_rate * 100).toFixed(4)}%`,
									selfBond: `${Tools.subStrings(Tools.formatPriceToFixed(Number(item.self_bond.match(/\d*(\.\d{0,4})?/)[0])), 2)} ${Constant.CHAINNAME.toLocaleUpperCase()}`,
									delegatorNum: item.delegator_num,
									bondHeight: Number(item.bond_height),
									unbondingHeight: item.unbonding_height && Number(item.unbonding_height) > 0 ? Number(item.unbonding_height) : '--',
									unbondingTime: (new Date(item.unbonding_time).getTime()) > 0 ? Tools.format2UTC(item.unbonding_time) : '--',
                                    identity: item.description.identity,
                                    url: item.icons || require('../assets/header_img.png')
								}
							});
							// this.items = this.getValidatorHeaderImg(this.items);
							this.showNoData = false;
						}else {
							this.showNoData = true;
							this.items = [];
						}
						this.showLoading = false;
						this.$refs.mtable && this.$refs.mtable.setValidatorField(status);
					}catch (e) {
						this.showLoading = false;
						this.showNoData = true;
						this.$refs.mtable && this.$refs.mtable.setValidatorField(status);
						console.error(e)
					}
				})
			},
			getValidatorHeaderImg(data){
				let url = 'https://keybase.io/_/api/1.0/user/lookup.json?fields=pictures&key_suffix=';
				for(let i = 0; i < data.length; i++){
					if(data[i].identity){
						Http.http(`${url}${data[i].identity}`).then(res =>{
							if(res && res.them && res.them[0].pictures && res.them[0].pictures.primary && res.them[0].pictures.primary.url){
                                data[i].url = res.them[0].pictures.primary.url;
							}else {
								data[i].url = require('../assets/header_img.png');
							}
						})
					}else {
						data[i].url = require('../assets/header_img.png');
					}
				}
				return data
			},
			getValidatorStatus(index){
				let validatorStatus;
				switch (index) {
					case 0 :
						validatorStatus = 'validator';
						break;
					case 1 :
						validatorStatus = 'candidate';
						break;
					case 2 :
						validatorStatus = 'jailed'
				}
				return validatorStatus
			},
			getDataList() {
				this.showLoading = true;
				this.selectValidatorStatus(Number(this.validatorTabIndex))
			},
		},
		beforeDestroy () {
			let defaultValidatorTabIndex = 0;
			localStorage.setItem('validatorTabIndex',defaultValidatorTabIndex);
		}
	}
</script>
<style lang="scss" scoped>
    @import '../style/mixin.scss';
    @import '../style/validator_list_page.scss';
</style>
