<template>
    <div class="blocks_list_page_container">
      <div class="block_list_title_container">
        <div class="block_list_title_content">
          <span class="block_list_title">Blocks</span>
        </div>
      </div>
      <div class="block_list_container">
        <div class="block_list_content">
          <div class="page_nav_container">
            <span>{{count}} Total</span>
            <div class="pagination_container">
              <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPageNum" v-model="currentPageNum" use-router></b-pagination-nav>
            </div>
          </div>
          <div class="block_list_table_container">
            <spin-component :showLoading="flShowLoading"></spin-component>
            <!--<blocks-list-table :items="items" :type="this.$route.params.type"-->
                               <!--:minWidth="tableMinWidth"-->
                               <!--:showNoData="showNoData" :status="this.$route.params.param"></blocks-list-table>-->
            <block-list-page-table :items="items" :showNoData="showNoData"></block-list-page-table>
            <div v-show="showNoData" class="no_data_show">
              No Data
            </div>
          </div>
          <div class="pagination_footer_container">
            <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPageNum" v-model="currentPageNum" use-router></b-pagination-nav>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
    import Service from "../util/axios"
    import Constant from "../constant/Constant"
    import Tools from '../util/Tools';
    import SpinComponent from './commonComponents/SpinComponent';
    import BlockListPageTable from "./table/BlockListPageTable";
    export default {
        name: "blockListPage",
        components: {SpinComponent, BlockListPageTable},
        watch: {
            currentPageNum(currentPageNum) {
                this.currentPageNum = currentPageNum;
                new Promise((resolve)=>{
                    this.getBlockList(currentPageNum, this.pageSize);
                resolve();
            }).then(()=>{
                document.getElementById('router_wrap').scrollTop = 0;
            })
            },
            $route() {
                clearInterval(this.timer);
                this.items = [];
                this.currentPage = 1;
                this.showNoData = false;
            }
        },
        data(){
            return {
                currentPage: 1,
                pageSize: 30,
                currentPageNum: this.$route.query.page ? Number(this.$route.query.page) : 1,
                totalPageNum: sessionStorage.getItem("blockListTotalPageNum") ? JSON.parse(sessionStorage.getItem("blockListTotalPageNum")) : 1,
                count: 0,
                items: [],
                showNoData: false,
                timer: null,
                flShowLoading: false
          }
        },
        mounted() {
            this.getBlockList(this.currentPageNum, this.pageSize)
        },
        methods:{
            linkGen(pageNum) {
              return pageNum === 1 ? '?' : `?page=${pageNum}`
            },
            getBlockList(currentPage, pageSize){
                this.flShowLoading = true;
                let url = `/api/blocks?page=${currentPage}&size=${pageSize}`;
                Service.http(url).then((data) => {
                    this.flShowLoading = false;
                    if(data){
                        let that = this;
                        clearInterval(this.timer);
                        this.count = data.Count;
                        this.totalPageNum =  Math.ceil(data.Count/this.pageSize);
                        sessionStorage.setItem('blockListTotalPageNum',JSON.stringify(this.totalPageNum))
                        this.items = data.Data.map(item => {
                            let txn = item.num_txs;
                            let precommit = item.last_commit && item.last_commit.length !== 0  ? item.last_commit.length : 0;
                            let [votingPower,denominator,numerator] = [0,0,0];
                            item.validators.forEach(listItem=>votingPower += listItem.voting_power);
                            item.validators.forEach(item=>denominator += item.voting_power);
                            if(item.last_commit && item.last_commit.length !== 0){
                                for(let i = 0; i < item.last_commit.length; i++){
                                    for (let j = 0; j < item.validators.length; j++){
                                        if(item.last_commit[i] === item.validators[j].address){
                                            numerator += item.validators[j].voting_power;
                                            break;
                                        }
                                    }
                                }
                            }
                            let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                            return {
                                Height: item.height,
                                Txn:txn,
                                Time:item.time,
                                Age: Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX),
                                'Precommit Validators':precommit,
                                'Voting Power': denominator !== 0? `${(numerator/denominator).toFixed(2)*100}%`:'',
                            };
                        })
                        this.timer = setInterval(function () {
                            that.items = that.items.map(item => {
                                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
                                item.Age = Tools.formatAge(currentServerTime,item.Time,Constant.SUFFIX,Constant.PREFIX)
                                return item
                            })
                        },1000)

                    }else{
                        this.items = [{Height:'',Txn:'',Fee:'',Age:'','Precommit Validators':'','Voting Power':''}];
                        this.showNoData = true;
                    }
                    this.showLoading = false;
                }).catch(e => {
                    this.flShowLoading = false;
                    console.log(e)
                })
            },
        }
    }
</script>

<style scoped lang="scss">
  .blocks_list_page_container{
    .block_list_title_container{
      width: 100%;
      background: #efeff1;
      border-bottom: 0.01rem solid #d6d9e0;
      .block_list_title_content{
        height: 0.61rem;
        max-width: 12.8rem;
        width: 100%;
        margin: 0 auto;
        display: flex;
        align-items: center;
        .block_list_title{
          color: #000;
          font-size: 0.22rem;
          font-weight: 500;
          padding-left: 0.2rem;
        }
      }
    }
    .block_list_container{
      max-width: 12.8rem;
      margin: 0 auto;
      .block_list_content{
        .page_nav_container{
          padding-left: 0.2rem;
          display: flex;
          justify-content: space-between;
          height: 0.7rem;
          align-items: center;
          span{
            color: #a2a2ae;
            font-size: 0.18rem;
          }
          .pagination_container{
            font-size: 0.14rem;
          }
        }
        .block_list_table_container{
          overflow-x: auto;
          -webkit-overflow-scrolling:touch;
        }
        .pagination_footer_container{
          display: flex;
          justify-content: flex-end;
          height: 0.7rem;
          align-items: center;
          margin-bottom: 0.2rem;
          font-size: 0.14rem;
        }
        .no_data_show{
          display: flex;
          justify-content: center;
          border-top:0.01rem solid #eee;
          border-bottom:0.01rem solid #eee;
          font-size:0.14rem;
          height:3rem;
          align-items: center;
        }
      }
    }
  }

</style>
