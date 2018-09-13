<template>
  <div class="proposals_list_page_wrap">
    <div class="proposals_list_title_wrap">
      <p :class="proposalsListPageWrap" style="margin-bottom:0;">
        <span class="proposals_list_title">Proposals</span>
      </p>
    </div>

    <div :class="proposalsListPageWrap">
      <div class="pagination total_num">
        <span class="proposals_list_page_wrap_hash_var" :class="count ? 'count_show' : 'count_hidden' ">{{count}} total</span>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
      <div style="position:relative;overflow-x: auto;">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="'Proposals'"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination" style='margin:0.2rem 0;'>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';

  export default {
    components:{
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        new Promise((resolve)=>{
          this.getDataList(currentPage, 30);
          resolve();
        }).then(()=>{
          document.getElementById('router_wrap').scrollTop = 0;
        })

      },
      $route() {
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = 1;
        this.getDataList(1, 30);
        this.showNoData = false;
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        proposalsListPageWrap: 'personal_computer_proposals_list_page',//1是显示pc端，0是移动端
        currentPage: 1,
        pageSize: 30,
        count: 0,
        items: [],
        showNoData:false,
        showLoading:false,
        innerWidth : window.innerWidth,
      }
    },
    beforeMount() {
      if (window.innerWidth > 910) {
        this.proposalsListPageWrap = 'personal_computer_proposals_list_page_wrap';
      } else {
        this.proposalsListPageWrap = 'mobile_proposals_list_page_wrap';
      }
    },
    mounted() {
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
      getDataList(currentPage, pageSize) {
        this.showLoading = true;
        let url=`/api/proposals/${currentPage}/${pageSize}`;
        axios.get(url).then((data)=>{
          if(data.status === 200){
            return data.data
          }
        }).then((data)=>{
          if(data.Data && typeof data === "object"){
            this.count = data.Count;
            this.items = data.Data.map(item =>{
              let title = item.title;
              let proposalId = item.proposal_id;
              let type = item.type;
              let status  = item.status;
              let submitBlock = item.submit_block;
              let submitTime = Tools.conversionTimeToUTCToYYMMDD(item.submit_time);
              let votingStartBlock = item.voting_start_block ? item.voting_start_block : "" ;
              return {
                Title : title,
                'Proposal ID' : proposalId,
                Type : type,
                Status : status,
                'Submit Block' : submitBlock,
                'Submit Time' : submitTime,
                'Voting Start Block' : votingStartBlock,
              }
            })
          }else {
            this.items = [{Title:"",'Proposal ID':"",Type:"",Status:"",'Submit Block':"",'Submit Time':'','Voting Start Block':''}]
          }
          this.showLoading = false;
        })
      }
    }
  }
</script>

<style lang="scss">
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
    width:100%;
  .proposals_list_page_wrap_hash_var{
    min-width:7rem;
  }
  }

  }
  .personal_computer_proposals_list_page_wrap {
  .transaction_information_content_title {
    height: 0.4rem;
    line-height: 0.4rem;
    font-size: 0.18rem;
    color: #000000;
    margin-bottom: 0;
  }
  @include pcCenter;
  min-height:4.6rem;
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
    font-weight: 500;
  }
  .proposals_list_page_wrap_hash_var {
    height:  0.62rem;
    line-height: 0.62rem;
    font-size: 0.22rem;
    color: #a2a2ae;
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
  }
  .transactions_detail_information_wrap {

  .information_props_wrap {
  @include flex;
    flex-direction: column;
    border-bottom: 0.01rem solid #eee;
    margin-bottom: 0.05rem;
  .information_value {
    overflow-x: auto;
  }

  }
  }

  .proposals_list_title {
    height: 0.3rem;
    line-height: 0.3rem;
    font-size: 0.18rem;
    color: #000000;
    margin-right: 0.2rem;
    font-weight: 500;
  }
  .proposals_list_page_wrap_hash_var {
    overflow-x: auto;
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
</style>
