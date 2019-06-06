<template>
  <div class="proposals_list_page_wrap">
    <div class="proposals_list_title_wrap">
      <p :class="proposalsListPageWrap" style="margin-bottom:0;">
        <span class="proposals_list_title">Proposals</span>
      </p>
    </div>
    <div class="graph_container" :class="[$store.state.isMobile ? 'mobile_graph_container' : '']" 
      v-if="$store.state.isMobile || (votingPeriodDatas.length === 1 && depositPeriodDatas.length === 1) ">
      <div v-for="v in votingPeriodDatas" :key="v.proposal_id"
        :class="[votingPeriodDatas.length === 1 && depositPeriodDatas.length === 0 ? 'one_votingPeriodDatas' : '',
        votingPeriodDatas.length > 1 ? 'two_votingPeriodDatas' : '']">
        <m-proposals-echart :data="v" v-if="v"></m-proposals-echart>
      </div>
      <div v-for="v in depositPeriodDatas" :key="v.proposal_id" 
        :class="[depositPeriodDatas.length === 1 && votingPeriodDatas.length === 0 ? 'one_depositPeriodDatas' : '',
        depositPeriodDatas.length > 1 ? 'two_depositPeriodDatas' : '']">
        <m-proposals-card :data="v" v-if="v"></m-proposals-card>
      </div>
    </div>

    <div class="graph_container graph_container_one"
      v-if="!$store.state.isMobile && ((votingPeriodDatas.length === 1 && depositPeriodDatas.length === 0) || depositPeriodDatas.length === 1 && votingPeriodDatas.length === 0)">
      <div v-for="v in votingPeriodDatas" :key="v.proposal_id">
        <m-proposals-echart :data="v" v-if="v"></m-proposals-echart>
      </div>
      <div v-for="v in depositPeriodDatas" :key="v.proposal_id">
        <m-proposals-card :data="v" v-if="v"></m-proposals-card>
      </div>
    </div>

    <div class="graph_container graph_container_warp"
      v-if="!$store.state.isMobile && (depositPeriodDatas.length > 1 || votingPeriodDatas.length > 1)">
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

    <div :class="proposalsListPageWrap">
      <div class="pagination total_num">
        <div style="height: 100%; display: flex; align-items: center;">
          <span class="proposals_list_page_wrap_hash_var" :class="count ? 'count_show' : 'count_hidden' ">{{count}} Proposals</span>
          <div class="icon_list">
            <div>
              <img src="../assets/critical.png" />
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
        <div>
          <div class="icon_list">
            <div>
              <i></i>
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
        proposalListTimer: null,
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
        let opacity = 1000;
        return arr.map(v => {
          let obj =  {
            value: v.voting_power,
            info: v,
            nodeClick: false,
            itemStyle: {
              color: `rgba(${color},${opacity / 1000})`,
              borderColor: '#CCDCFF',
              borderWidth: 0
            }
          }
          opacity = opacity - 90;
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
          o.participation = item.level && item.level.gov_param && item.level.gov_param.participation && this.formatNumber(item.level.gov_param.participation);
          o.threshold = item.level && item.level.gov_param && item.level.gov_param.threshold && this.formatNumber(item.level.gov_param.threshold);
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
          let data = [
            {
              name: 'Participant',
              value: votes,
              itemStyle: {
                color: '#3598DB',
                borderColor: '#CCDCFF',
                borderWidth: 0
              },
              children: [
                {
                  name: 'Yes',
                  value: yes,
                  itemStyle: {
                    color: '#45B4FF',
                    borderColor: '#CCDCFF',
                    borderWidth: 0
                  },
                  children: this.formatGrahpChildren(yesArr, '69,180,255')
                },
                {
                  name: 'Abstain',
                  value: abstain,
                  itemStyle: {
                    color: '#CCDCFF',
                    borderColor: '#CCDCFF',
                    borderWidth: 0
                  },
                  children: this.formatGrahpChildren(abstainArr, '204,220,255')
                },
                {
                  name: 'No',
                  value: no,
                  itemStyle: {
                    color: '#FFCF65',
                    borderColor: '#CCDCFF',
                    borderWidth: 0
                  },
                  children: this.formatGrahpChildren(noArr, '255,207,101')
                },
                {
                  name: 'NoWithVeto',
                  value: noWithVeto,
                  itemStyle: {
                    color: '#FE8A8A',
                    borderColor: '#CCDCFF',
                    borderWidth: 0
                  },
                  children: this.formatGrahpChildren(noWithVetoArr, '254,138,138')
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
              children: [
                {
                  name: '',
                  value: all - votes,
                  nodeClick: false,
                  itemStyle: {
                    color: '#E5E9FB',
                    borderColor: '#CCDCFF',
                    borderWidth: 0
                  },
                  children: [
                    {
                      name: '',
                      value: all - votes,
                      nodeClick: false,
                      itemStyle: {
                        color: '#E5E9FB',
                        borderColor: '#CCDCFF',
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
        let url=`/api/gov/depositvotingproposals`;
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
            let that = this;
            clearInterval(this.proposalListTimer);
            // this.proposalListTimer = setInterval(function () {
              that.items = proposalList.Data.map(item =>{
                let proposalId = item.proposal_id === 0 ? "--" : item.proposal_id;
                let type = item.type;
                let status  = item.status;
                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                let submitTime = (new Date(item.submit_time).getTime()) > 0 ? Tools.format2UTC(item.submit_time) : '--';
                let depositEndTime = (new Date(item.deposit_end_time).getTime()) > 0 ? Tools.format2UTC(item.deposit_end_time) : '--';
                let votingEndTime = (new Date(item.voting_end_time).getTime()) > 0 ? Tools.format2UTC(item.voting_end_time) : '--';
                let title = Tools.formatString(item.title,20,"...");
                let final_votes = Object.keys(item.final_votes).length > 0 ? item.final_votes : null;
                let finalTotalVotes = 0;
                for (let k in item.final_votes) {
                  finalTotalVotes += Number(item.final_votes[k]);
                }
                if(final_votes) {
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
              })
            // },1000);
          }else {
            this.items = [];
            this.showNoData = true;
          }
          this.showLoading = false;
        }).catch(e => {
          console.log(e)
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
      height:0.3rem;
      .icon_list {
        display: flex;
        div {
          font-size: 14px;
          color: #22252A;
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
      height:0.7rem;
      align-items: center;
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
  //min-width: 8rem;
    a {
      text-decoration: none;
    }
  }
  .proposals_list_title_wrap {
    width: 100%;
    border-bottom: 1px solid #d6d9e0 !important;
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
    //margin-left:0.1rem;
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
    visibility: visible;
  }
  .count_hidden{
    visibility: hidden;
  }
  .graph_container {
    display: flex;
    margin-top: 0.3rem;
    min-width: 12.8rem;
    max-width: 12.8rem;
    flex-wrap: wrap;
    & > div {
      margin-right: 0.2rem;
      margin-top: 0.2rem;
      &:nth-of-type(2n) {
        margin-right: 0rem;
      }
    }
  }
  .graph_container_warp {
    display: block;
    & > div {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
      margin-right: 0rem!important;
      & > div {
        margin-right: 0.2rem;
        margin-top: 0.2rem;
        &:nth-of-type(2n) {
          margin-right: 0rem;
        }
      }
      &:nth-of-type(2) {
        margin-top: 0rem!important;
      }
    }
  }
  .graph_container_one {
    width: 100%;
    & > div {
      width: 100%;
      margin-right: 0rem!important;
      & > div {
        width: 100%;
      }
    }
  }
  .mobile_graph_container {
    display: flex;
    flex-direction: column;
    width: 100%;
    & > div {
      margin-left: 0rem!important;
      margin-right: 0rem!important;
      flex: 1;
      display: flex;
      & > div {
        width:100%!important;
        margin-left: 0.2rem!important;
        margin-right: 0.2rem!important;
      }
    }
  }
</style>
