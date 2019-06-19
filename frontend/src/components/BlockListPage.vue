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
          <span>Current Height {{count}}</span>
          <div class="pagination_container">
            <m-pagination :page-size="pageSize"
                          :total="count"
                          :page="currentPageNum"
                          :page-change="pageChange"
                          :range="range"></m-pagination>
          </div>
        </div>
        <div class="block_list_table_container">
          <spin-component :showLoading="flShowLoading"></spin-component>
          <m-block-list-page-table :items="items"></m-block-list-page-table>
          <div v-show="showNoData"
               class="no_data_show">No Data</div>
        </div>
        <div class="pagination_footer_container">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Service from "../service";
import Constant from "../constant/Constant";
import Tools from "../util/Tools";
import SpinComponent from "./commonComponents/SpinComponent";
import BlockListPageTable from "./table/BlockListPageTable";
import MBlockListPageTable from "./table/MBlockListPageTable";
import MPagination from "./commonComponents/MPagination";

export default {
  name: "blockListPage",
  components: {
    SpinComponent,
    BlockListPageTable,
    MBlockListPageTable,
    MPagination
  },
  data () {
    return {
      pageSize: 30,
      currentPageNum: this.forCurrentPageNum(),
      count: sessionStorage.getItem("blockListTotal")
        ? JSON.parse(sessionStorage.getItem("blockListTotal"))
        : 0,
      items: [],
      showNoData: false,
      timer: null,
      flShowLoading: false,
      range: [],
      currentPageNumCache: 0,
      isoMunted: false
    };
  },
  mounted () {
    this.getLatestheight();
    this.getBlockList();
  },
  methods: {
    forCurrentPageNum () {
      let currentPageNum = 1;
      let urlPageSize = this.$route.query.page && Number(this.$route.query.page);
      currentPageNum = urlPageSize ? urlPageSize : 1;
      return currentPageNum;
    },
    pageChange (pageNum) {
      this.currentPageNum = pageNum;
      if (this.currentPageNumCache === this.currentPageNum) {
        return;
      }
      this.currentPageNumCache = this.currentPageNum;
      if (this.isoMunted) {
        let path = '/blocks';
        history.pushState(null, null, `/#${path}?page=${pageNum}`);
        this.getLatestheight();
        this.getBlockList();
      }
    },
    getLatestheight () {
      Service.commonInterface({ blockListLatestheight: {} }, data => {
        try {
          this.isoMunted = true;
          this.count = data.data || 0;
          sessionStorage.setItem("blockListTotal", JSON.stringify(this.count));
        } catch (e) { }
      });
    },
    getBlockList () {
      this.flShowLoading = true;
      this.showNoData = false;
      Service.commonInterface(
        {
          blockList: {
            pageNumber: this.currentPageNum,
            pageSize: this.pageSize
          }
        },
        data => {
          this.flShowLoading = false;
          try {
            if (data) {
              let that = this;
              clearInterval(this.timer);
              this.items = data.map(item => {
                let currentServerTime =
                  new Date().getTime() + that.diffMilliseconds;
                return {
                  height: item.height,
                  transactions: item.txs_num,
                  validators: `${item.precommit_validator_num}/${
                    item.validator_num_for_height
                    }`,
                  votingPower:
                    item.voting_power_for_height > 0
                      ? Tools.formatDecimalNumberToFixedNumber(
                        (item.precommit_voting_power /
                          item.voting_power_for_height) *
                        100
                      ) + "%"
                      : "--",
                  moniker: item.proposer_moniker,
                  proposerAddr: item.proposer_as_validator_addr,
                  time: item.timestamp,
                  timestamp:
                    new Date(item.timestamp).getTime() > 0
                      ? Tools.format2UTC(item.timestamp)
                      : "--",
                  Age: Tools.formatAge(
                    currentServerTime,
                    item.timestamp,
                    Constant.SUFFIX,
                    Constant.PREFIX
                  )
                };
              });
              let heightArr = this.items.map(v => v.height);
              this.range = [Math.max.apply(null, heightArr), Math.min.apply(null, heightArr)];
              this.timer = setInterval(function () {
                let currentServerTime =
                  new Date().getTime() + that.diffMilliseconds;
                that.items = that.items.map(item => {
                  item.Age = Tools.formatAge(
                    currentServerTime,
                    item.time,
                    Constant.SUFFIX,
                    Constant.PREFIX
                  );
                  return item;
                });
              }, 1000);
            } else {
              this.items = [];
              this.showNoData = true;
            }
          } catch (e) {
            this.items = [];
            this.showNoData = true;
          }
        }
      );
    }
  }
};
</script>

<style scoped lang="scss">
.blocks_list_page_container {
  .block_list_title_container {
    width: 100%;
    background: #efeff1;
    border-bottom: 0.01rem solid #d6d9e0;
    .block_list_title_content {
      height: 0.61rem;
      max-width: 12.8rem;
      width: 100%;
      margin: 0 auto;
      display: flex;
      align-items: center;
      .block_list_title {
        color: #000;
        font-size: 0.22rem;
        font-weight: 500;
        padding-left: 0.2rem;
      }
    }
  }
  .block_list_container {
    max-width: 12.8rem;
    margin: 0 auto;
    .block_list_content {
      .page_nav_container {
        padding-left: 0.2rem;
        display: flex;
        justify-content: space-between;
        height: 0.7rem;
        align-items: center;
        span {
          color: #a2a2ae;
          font-size: 0.18rem;
        }
        .pagination_container {
          font-size: 0.14rem;
        }
      }
      .block_list_table_container {
        overflow-x: auto;
        overflow-y: hidden;
        -webkit-overflow-scrolling: touch;
      }
      .pagination_footer_container {
        display: flex;
        justify-content: flex-end;
        height: 0.7rem;
        align-items: center;
        margin-bottom: 0.2rem;
        font-size: 0.14rem;
      }
      .no_data_show {
        display: flex;
        justify-content: center;
        border-top: 0.01rem solid #eee;
        border-bottom: 0.01rem solid #eee;
        font-size: 0.14rem;
        height: 3rem;
        align-items: center;
      }
    }
  }
}
</style>
