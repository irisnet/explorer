<template>
  <div class="proposals_list_page_wrap">
    <div class="proposals_list_title_wrap">
      <p :class="proposalsListPageWrap" style="margin-bottom:0;">
        <span class="proposals_list_title">Proposals</span>
      </p>
    </div>
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
   

    <div :class="proposalsListPageWrap">
      <div class="pagination total_num" :class="[$store.state.isMobile ? 'mobile_graph_pagination_container' : '']">
        <div style="height: 70px; display: flex; align-items: center;">
          <span class="proposals_list_page_wrap_hash_var" :class="count ? 'count_show' : 'count_hidden' ">{{count}} Proposals</span>
          <div class="icon_list">
            <div>
              <img src="../assets/critical.png" style="margin-left: 0;" />
              <span>Critical</span>
            </div>
            <div>
              <img src="../assets/important.png" />
              <span>Important</span>
            </div>
            <div>
              <img src="../assets/normal.png" />
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
      <div style="position:relative;overflow-x: auto; overflow-y: hidden;-webkit-overflow-scrolling:touch;">
        <m-proposals-list-table :items="items"></m-proposals-list-table>
        <!-- <blocks-list-table :items="items" :type="'Proposals'" :showNoData="showNoData" :minWidth="tableMinWidth"></blocks-list-table> -->
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination" style='margin:0.2rem 0;'>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
    </div>
    <spin-component :showLoading="showLoading"/>
  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  import Service from "../util/axios";
  import Constant from "../constant/Constant"
  import MProposalsCard from './commonComponents/MProposalsCard';
  import MProposalsEchart from './commonComponents/MProposalsEchart';
  import MProposalsListTable from './table/MProposalsListTable';

  export default {
    components:{
      BlocksListTable,
      SpinComponent,
      MProposalsCard,
      MProposalsEchart,
      MProposalsListTable
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        new Promise((resolve)=>{
          this.getDataList(currentPage, 30);
          resolve();
        }).then(()=>{
          Tools.scrollToTop()
        })
      },
      $route() {
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = 1;
        this.getDataList(1, 30);
        this.showNoData = false;
        this.computeMinWidth();
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        proposalsListPageWrap: 'personal_computer_proposals_list_page',
        currentPage: 1,
        pageSize: 30,
        count: 0,
        items: [],
        showNoData:false,
        showLoading:false,
        innerWidth : window.innerWidth,
        tableMinWidth:"",
        votingPeriodDatas: [],
        depositPeriodDatas: []
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
      this.getGrahpData();
      this.getDataList(1, 30);
      window.addEventListener('resize',this.onresize);
    },
    beforeDestroy() {
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
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
        let votingPeriodDatas = data.filter(v => v.status === 'VotingPeriod');
        let depositPeriodDatas = data.filter(v => v.status === 'DepositPeriod');
        this.votingPeriodDatas = votingPeriodDatas.map(item => {
          let o = {};
          o.proposal_id = item.proposal_id;
          o.title = item.title;
          o.level = item.level.name;
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
          o.participation_num = item.level && item.level.gov_param && item.level.gov_param.participation && this.formatNumber(item.level.gov_param.participation);
          o.threshold_num = item.level && item.level.gov_param && item.level.gov_param.threshold && this.formatNumber(item.level.gov_param.threshold);
          o.participation = all ? (votes / all) * 100 : 0;
          o.threshold = votes ? (yes / votes) * 100 : 0;
          let data = [
            {
              name: 'Participant',
              value: votes,
              itemStyle: {
                color: '#3598DB',
                borderColor: '#ECEFFF',
                borderWidth: 0
              },
              children: [
                {
                  name: 'Yes',
                  value: yes,
                  itemStyle: {
                    color: '#45B4FF',
                    borderColor: '#ECEFFF',
                    borderWidth: 0
                  },
                  children: this.formatGrahpChildren(yesArr, {h: [205, 204], s: [100, 100], l: [79, 35]})
                },
                {
                  name: 'Abstain',
                  value: abstain,
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
          o.data = data;
          return o;
        });
        depositPeriodDatas.forEach(v => {
          if (v.level && v.level.gov_param && v.level.gov_param.min_deposit && v.level.gov_param.min_deposit.amount) {
            v.min_deposit_number = Number(v.level.gov_param.min_deposit.amount);
            v.min_deposit_format = `${Tools.formatAmount(v.level.gov_param.min_deposit.amount)} IRIS`;
          }
          if (v.intial_deposit && v.intial_deposit.amount) {
            v.intial_deposit_number = Number(v.intial_deposit.amount);
            v.intial_deposit_format = `${Tools.formatAmount(v.intial_deposit.amount)} IRIS`;
          }
          if (v.total_deposit && v.total_deposit.amount) {
            v.total_deposit_number = Number(v.total_deposit.amount);
            v.total_deposit_format = `${Tools.formatAmount(v.total_deposit.amount)} IRIS`;
          }
          v.intial_deposit_number_per = this.isNumber(v.intial_deposit_number) && this.isNumber(v.min_deposit_number) ?
          (v.intial_deposit_number / v.min_deposit_number) * 100 + '%' : 0;
          v.total_deposit_number_per = this.isNumber(v.total_deposit_number) && this.isNumber(v.min_deposit_number) ?
          (v.total_deposit_number / v.min_deposit_number) * 100 + '%' : 0;
        });
        this.depositPeriodDatas = depositPeriodDatas;
      },
      isNumber(n) {
        return typeof n === 'number'
      },
      getGrahpData() {
        let url=`/api/gov/deposit_voting_proposals`;
        Service.http(url).then((data)=>{
          if (data && Array.isArray(data) && data.length > 0) {
            this.formatGrahpData(data);
          }
        }).catch(err => {
        });
      },
      getDataList(currentPage, pageSize) {
        this.showLoading = true;
        let url=`/api/gov/proposals?page=${currentPage}&size=${pageSize}`;
        Service.http(url).then((proposalList)=>{
          if(proposalList.Data){
            this.showNoData = false;
            this.count = proposalList.Count;
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
          this.showLoading = false;
        }).catch(e => {
          this.items = [];
          this.showNoData = true;
          this.showLoading = false;
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import '../style/mixin.scss';

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
          color: #22252A;
          display: flex;
          align-items: center;
          i {
            width: 12px;
            height: 12px;
            border-radius: 2px;
            display: inline-block;
            margin-left: 50px;
            background-color: #45B4FF;
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
  .total_num{
    @include flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    & > div {
      padding: 4px 0;
    }
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
      padding-bottom: 0.2rem;
    }
    .mobile_proposals_list_page_wrap {
      @include flex;
      flex-direction: column;
      overflow-x: auto;
      -webkit-overflow-scrolling:touch;
      width:100%;
      .proposals_list_page_wrap_hash_var{
        min-width:7rem;
      }
    }
  }
  .personal_computer_proposals_list_page_wrap {
    padding-bottom: 0.2rem;
    width: 100%!important;
    .transaction_information_content_title {
      height: 0.4rem;
      line-height: 0.4rem;
      font-size: 0.18rem;
      color: #000000;
      margin-bottom: 0;
      @include fontWeight;
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
      color: #a2a2ae;
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
    padding-left: 0.2rem;
    @include fontWeight;
  }
  .proposals_list_page_wrap_hash_var {
    overflow-x: auto;
    -webkit-overflow-scrolling:touch;
    height: 0.3rem;
    line-height: 0.3rem;
    font-size: 0.18rem;
    color: #a2a2ae;
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
    margin: 0.3rem auto 0.1rem;
  }
  .votingPeriodDatas_one {
    justify-content: space-between;
    & > div {
      margin-top: 0.2rem !important;
      width: calc(50% - 0.1rem);
      .propsals_card_container {
        height: 2.3rem;
      }
    }
    & > div:nth-child(1) {
      width: 100%;
      margin-top: 0rem !important;
      .propsals_echart_container {
        width: 100%;
      }
    }
  }
  .depositPeriodDatas_one {
    height: 2.3rem;
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
    margin-top: 0.1rem;
    & > div {
      flex-direction: row;
      display: flex;
      justify-content: space-between;
      flex-wrap: wrap;
      & > div {
        width: calc(50% - 0.1rem);
        margin-top: 0.2rem;
        .propsals_card_container {
          height: 2.3rem;
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
      & > div {
        width: 100%!important;
        margin: auto;
      }
      & > div.propsals_card_container {
        height: 2.3rem;
      }
    }
    .propsals_echart_container {
      flex-direction: column;
    }
  }
  .mobile_graph_pagination_container {
    padding-left: 0!important;
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