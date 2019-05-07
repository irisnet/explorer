<template>
    <div class="top_list_page">
      <div class="top_list_title_container">
        <div class="top_list_title_content">
          <span class="top_list_title">Top 100 Rich List</span>
        </div>
      </div>
      <div class="top_list_container">
        <div class="top_list_content">
          <div class="top_list_title_wrap">
            <h5 class="top_list_time_content"><span v-show="latestTime">Updated ：{{latestTime}}+UTC</span></h5>
            <h5 class="top_list_assets_content"><span><i>*</i> The assets include the delegated, unbonding and balance.</span></h5>
          </div>
          <div class="top_list_table_wrap">
            <div class="top_list_table_content">
              <spin-component :show-loading="showLoading"></spin-component>
              <top-list-table :items="topList" :showNoData="showNoData"></top-list-table>
              <div v-show="showNoData" class="no_data_show">
                No Data
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
    import SpinComponent from './commonComponents/SpinComponent';
    import TopListTable from "./table/TopListTable";
    import Server from '../util/axios';
    import Tools from '../util/Tools'
    export default {
        name: "TopList",
        components: {SpinComponent, TopListTable},
        data () {
            return {
                topList:[],
                showNoData: false,
                showLoading: false,
                latestTime: "",
                richListTimer: null,
            }
        },
        mounted(){
            this.getTopList();
            clearInterval(this.richListTimer);
            let that = this;
            this.richListTimer = setInterval(function () {
              that.getTopList();
            },60000);
        },
        methods:{
            getTopList(){
                let url= `/api/accounts`;
                this.showLoading = true;
                Server.http(url).then(res => {
                    this.showLoading = false;
                    if(res) {
                      this.latestTime = this.getUpDatedTime(res);
                      this.showLoading = false;
                        this.topList = res.map(item => {
                            return {
                              rank: item.rank,
                              Address: item.address,
                              Balance: `${Tools.formatPrice(Tools.convertScientificNotation3Number(Tools.formatNumber(item.balance[0].amount)))}`,
                              Percentage: `${(item.percent * 100).toFixed(4)}`
                               }
                           })
                    }else {
                        this.showLoading = false;
                        this.showNoData = true;
                        this.topList = [{
                            '#':'',
                            Address: '',
                            Balance: '',
                            Percentage: ''
                        }];
                    }
                }).catch(err => {
                    this.showLoading = false;
                    this.showNoData = true;
                    console.error(err);
                    this.topList = [{
                        '#':'',
                        Address: '',
                        Balance: '',
                        Percentage: ''
                    }];
                   })
            },
            getUpDatedTime(upDatedTime){
              let maxUpDatedTime = 0;
              upDatedTime.forEach(item => {
                if(item.update_at > maxUpDatedTime){
                  maxUpDatedTime = item.update_at
                }
              });
              return new Date(maxUpDatedTime * 1000).toISOString().split(".")[0].replace("T", " ")
            }
        }
    }
</script>

<style scoped lang="scss">
  @import "../style/mixin.scss";
  .top_list_page{

    .top_list_title_container{
      background: rgba(239, 239, 241, 1);
      border-bottom: 0.01rem solid rgba(215, 217, 224, 1);
      .top_list_title_content{
        max-width: 12.8rem;
        margin: 0 auto;
        .top_list_title{
          height: 0.62rem;
          display: flex;
          align-items: center;
          font-size: 0.22rem;
          color: rgba(34, 37, 42, 1);
          line-height: 0.26rem;
          padding-left: 0.2rem;
        }
      }
    }
  }
  .top_list_container{
    position: relative;
    top: -0.01rem;
    width: 100%;
    margin-bottom: 0.2rem;
    .top_list_content{
      max-width: 12.8rem;
      width: 100%;
      margin: 0 auto;
      .top_list_time_content{
        padding-left: 0.2rem;
        opacity:0.6;
        color:#000;
        font-size: 0.12rem;
      }
      .top_list_assets_content{
        padding: 0 0.3rem 0 0;
        opacity:0.8;
        color:#000;
        font-size: 0.12rem;
        i{
          color: #3598db;
        }
      }
      .top_list_table_wrap{
        overflow-x: auto;
        .top_list_table_content{
          width: 12.5rem;
        }
      }
    }
  }
  .no_data_show {
    @include flex;
    justify-content: center;
    border-top: 0.01rem solid #eee;
    border-bottom: 0.01rem solid #eee;
    font-size: 0.14rem;
    height: 3rem;
    align-items: center;
  }
  @media (min-width: 960px){
    .top_list_title_wrap{
      padding: 0.15rem 0;
      display: flex;
      justify-content: space-between
    }
  }
  @media (max-width: 959px){
    .top_list_title_wrap{
      padding: 0.15rem 0;
      display: flex;
      flex-direction:column;
      align-items:flex-start;
      .top_list_assets_content{
        padding: 0.1rem 0.3rem 0 0.2rem !important;
      }
    }
  }
</style>
