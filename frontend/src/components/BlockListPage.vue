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
                        <span v-if="currentHeight > 0" class="skip_route current_height">
                            <router-link :to="`/block/${currentHeight}`">{{currentHeight}}</router-link>
                        </span>
                    </span>
                    <div v-if="!$store.state.isMobile" class="pagination_container">
                        <m-pagination
                            :ascending="false"
                            :page-size="pageSize"
                            :total="currentHeightByDb"
                            :page="currentPageNum"
                            :page-change="pageChange"
                            :range="range"
                        ></m-pagination>
                    </div>
                </div>
                <div class="block_list_table_container">
                    <m-block-list-page-table :items="items"></m-block-list-page-table>
                    <div v-show="showNoData" class="no_data_show">
                        <img src="../assets/no_data.svg">
                    </div>
                </div>
                <div class="pagination_footer_container">
                    <m-pagination
                        :ascending="false"
                        :page-size="pageSize"
                        :total="currentHeightByDb"
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
import SpinComponent from "../loadingComponent/SpinComponent";
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
            isoMunted: false,
	        currentHeightByDb:sessionStorage.getItem("blockListdbTotal")
		        ? Number(sessionStorage.getItem("blockListdbTotal"))
		        : 0,
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
                    this.currentHeight = data.block_height_lcd || 0;
                    this.currentHeightByDb = data.block_height_db || 0;
                    sessionStorage.setItem(
                        "blockListTotal",
                        this.currentHeight
                    );
	                sessionStorage.setItem(
		                "blockListdbTotal",
		                this.currentHeightByDb
	                );
                } catch (e) {}
            });
        },
        getBlockList() {
        	this.$store.commit('flShowLoading',true);
            this.showNoData = false;
            Service.commonInterface(
                {
                    blockList: {
                        pageNumber: this.currentPageNum,
                        pageSize: this.pageSize
                    }
                },
                data => {
	                this.$store.commit('flShowLoading',false);
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
        padding-bottom: 0.2rem;
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
                .current_height{
                    a{
                        font-size: 0.18rem !important;

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
        background-color: #F5F7FD;
    }
    .block_list_table_container {
        padding-top: 0.7rem;
    }
}
@media screen and (max-width: 910px) {
    .blocks_list_page_container_fixed {
        .page_nav_container {
            padding-left: 0.1rem !important;
            position: static;
        }
        .block_list_table_container{
            padding-top: 0;
            margin: 0 0.1rem;
        }
        .page_nav_container{
            .pagination_container{
                display: none;
            }
        }
        .pagination_footer_container{
            .common_pagination_content{
                margin-right: 0.1rem;
            }
        }
    }
}
</style>
