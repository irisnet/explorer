<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="blockListPageWrap">
        <span class="blocks_list_title">{{listTitleName}}</span>
      </p>
    </div>
    <div :class="blockListPageWrap">
      <div class="pagination total_num">
        <span class="blocks_list_page_wrap_hash_var" v-show="count">{{count}} total</span>
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>

      <div class="list_wrap">
        <spin-component :showLoading="showLoading"/>
        <blocks-list-table :items="items" :type="this.$route.params.type"
                           :minWidth="tableMinWidth"
                           :showNoData="showNoData"></blocks-list-table>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
      <div class="pagination">
        <b-pagination size="md" :total-rows="count" v-model="currentPage" :per-page="pageSize">
        </b-pagination>
      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../util/Tools';
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  import Service from "../util/axios"
  export default {
    name:"blockListPage",
    components:{
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      currentPage(currentPage) {
        this.currentPage = currentPage;
        this.$router.push({
          path: this.$route.path,
          query:{
            pagenum: this.currentPage
          }
        });
        new Promise((resolve)=>{
          this.getDataList(currentPage, this.pageSize);
          resolve();
        }).then(()=>{
          document.body.scrollTop = 0;
        })
      },
      $route() {
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = Number(this.$route.query.pagenum);
        this.getDataList(Number(this.$route.query.pagenum), this.pageSize);
        this.showNoData = false;
        this.listTitleName = 'Blocks';
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blockListPageWrap: 'personal_computer_blocks_list_page',
        currentPage: Number(this.$route.query.pagenum),
        pageSize: 30,
        count: localStorage.getItem("blockListCount") ? Number(localStorage.getItem("blockListCount")) : 0,
        items: [],
        type: 'list',
        showNoData: false,
        showLoading: false,
        tableMinWidth: '',
        listTitleName: "",
        minWidth: 8.5,
      }
    },
    beforeMount() {
      this.type = this.$route.params.type;
      this.blockListPageWrap = Tools.getClassByWindowInnerWidth(window.innerWidth);
    },
    mounted() {
      this.getDataList(Number(this.$route.query.pagenum), this.pageSize, this.$route.params.type);
      this.listTitleName = 'Blocks';
      this.tableMinWidth = this.minWidth;
    },
    methods: {
      getDataList(currentPage, pageSize) {
        this.showLoading = true;
        let uri = `/api/blocks/${currentPage}/${pageSize}`;
        Service.http(uri).then((blockList) => {
          if(blockList){
            this.count = blockList.Count;
            localStorage.setItem("blockListCount",this.count);
            if(blockList.Data && typeof blockList.data === "object"){
              this.items = blockList.Data.map(item => {
                let txn = item.NumTxs;
                let precommitValidatorsLength = item.Block.LastCommit.Precommits.length;
                let [votingPower,computeValidatorsVotingPowerDenominator,computeValidatorsVotingPowerNumerator] = [0,0,0];
                item.Validators.forEach(listItem=>votingPower += listItem.VotingPower);
                item.Validators.forEach(item=>computeValidatorsVotingPowerDenominator += item.VotingPower);
                for(let i = 0; i < item.Block.LastCommit.Precommits.length; i++){
                  for (let j = 0; j < item.Validators.length; j++){
                    if(item.Block.LastCommit.Precommits[i].ValidatorAddress === item.Validators[j].Address){
                      computeValidatorsVotingPowerNumerator += item.Validators[j].VotingPower;
                      break;
                    }
                  }
                }
                return {
                  Height: item.Height,
                  Txn: txn,
                  Timestamp: Tools.conversionTimeToUTCToYYMMDD(item.Time),
                  'Precommit Validators': precommitValidatorsLength,
                  'Voting Power': computeValidatorsVotingPowerDenominator !== 0 ? `${(computeValidatorsVotingPowerNumerator/computeValidatorsVotingPowerDenominator).toFixed(2)*100}%` : '',
                };
              })
            }else{
              this.items = [
                {
                  Height: '',
                  Txn: '',
                  Fee: '',
                  Timestamp: '',
                  'Precommit Validators': '',
                  'Voting Power': ''
                }
              ];
              this.showNoData = true;
            }
            this.showLoading = false;
          }else {
            console.log(blockList.msg)
          }
        }).catch(e => {
          this.showLoading = false;
          console.log(e)
        })
        }
      }
    }
</script>
<style lang="scss" scoped>
  @import "../style/tabListCommonStyle.scss";
</style>
