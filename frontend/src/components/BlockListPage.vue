<template>
    <div
        class="blocks_list_page_container"
        :class="[$store.state.isMobile ? 'mobile_blocks_list_page_container' : 'blocks_list_page_container_fixed']"
    >
        <div class="block_list_container">
            <div class="block_list_content">
                <div class="page_nav_container">
                    <span>
                        Current Height:
                        <span v-if="currentHeight > 0" class="skip_route">
                            <router-link :to="`/block/${currentHeight}`">{{currentHeight}}</router-link>
                        </span>
                    </span>
                    <div v-if="!$store.state.isMobile" class="pagination_container">
                        <m-pagination
                            :ascending="false"
                            :page-size="pageSize"
                            :total="currentHeight"
                            :page="currentPageNum"
                            :page-change="pageChange"
                            :range="range"
                        ></m-pagination>
                    </div>
                </div>
                <div class="block_list_table_container">
                    <spin-component :showLoading="flShowLoading"></spin-component>
                    <m-block-list-page-table :items="items"></m-block-list-page-table>
                    <div v-show="showNoData" class="no_data_show">No Data</div>
                </div>
                <div class="pagination_footer_container">
                    <m-pagination
                        :ascending="false"
                        :page-size="pageSize"
                        :total="currentHeight"
                        :page="currentPageNum"
                        :page-change="pageChange"
                        :range="range"
                    ></m-pagination>
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
    data() {
        return {
            pageSize: 30,
            currentPageNum: this.forCurrentPageNum(),
            currentHeight: sessionStorage.getItem("blockListTotal")
                ? Number(sessionStorage.getItem("blockListTotal"))
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
    watch: {
        $route(newVal) {
            // 有时候 mounted 方法不起作用，为此添加该 watch
            this.currentPageNum = Number(this.$route.query.page || 1);
            this.isoMunted = false;
            this.getLatestheight();
            this.getBlockList();
        }
    },
    mounted() {
        this.getLatestheight();
        this.getBlockList();
    },
    methods: {
        forCurrentPageNum() {
            let currentPageNum = 1;
            let urlPageSize =
                this.$route.query.page && Number(this.$route.query.page);
            currentPageNum = urlPageSize ? urlPageSize : 1;
            return currentPageNum;
        },
        pageChange(pageNum) {
            this.currentPageNum = pageNum;
            if (this.currentPageNumCache === this.currentPageNum) {
                return;
            }
            this.currentPageNumCache = this.currentPageNum;
            if (this.isoMunted) {
                let path = "/blocks";
                history.pushState(null, null, `/#${path}?page=${pageNum}`);
                this.getLatestheight();
                this.getBlockList();
            }
        },
        getLatestheight() {
            Service.commonInterface({ blockListLatestheight: {} }, data => {
                try {
                    this.isoMunted = true;
                    this.currentHeight = data.data || 0;
                    sessionStorage.setItem(
                        "blockListTotal",
                        this.currentHeight
                    );
                } catch (e) {}
            });
        },
        getBlockList() {
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
                                    new Date().getTime() +
                                    that.diffMilliseconds;
                                return {
                                    height: item.height,
                                    transactions: item.txs_num,
                                    validators: `${item.precommit_validator_num}/${item.validator_num_for_height}`,
                                    votingPower:
                                        item.voting_power_for_height > 0
                                            ? Tools.subStrings(String((item.precommit_voting_power /
                                                      item.voting_power_for_height) *
                                                      100
                                              ), 4) + "%"
                                            : "--",
                                    moniker: item.proposer_moniker,
                                    proposerAddr:
                                        item.proposer_as_validator_addr,
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
                            this.range = [
                                Math.max.apply(null, heightArr),
                                Math.min.apply(null, heightArr)
                            ];
                            this.timer = setInterval(function() {
                                let currentServerTime =
                                    new Date().getTime() +
                                    that.diffMilliseconds;
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
                z-index: 5;
                .skip_route {
                    margin-left: 0.09rem;
                    a {
                        color: var(--bgColor) !important;
                        cursor: pointer;
                    }
                }
                span {
                    color: var(--contentColor);
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
.mobile_blocks_list_page_container {
    .block_list_container {
        padding: 0 0.1rem;
    }
}
.blocks_list_page_container_fixed {
    div.block_list_title_container {
        position: fixed;
    }
    .page_nav_container {
        position: fixed;
        width: 100%;
        max-width: 12.8rem;
        background-color: #ffffff;
    }
    .block_list_table_container {
        padding-top: 0.7rem;
    }
}
@media screen and (max-width: 910px) {
    .page_nav_container {
        padding-left: 0.1rem !important;
    }
}
</style>
