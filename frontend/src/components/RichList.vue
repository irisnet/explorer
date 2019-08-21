<template>
    <div class="top_list_page">
        <div class="top_list_title_container">
            <div class="top_list_title_content">
                <div class="top_list_title_content_div">
                    <span class="top_list_title">Top 100 Addresses by IRIS</span>
                    <div class="tooltip_icon">
                        <span>?</span>
                        <div class="tooltip_span">
                            <div>The assets include the delegated tokens, unbonding tokens and remaining tokens on the address</div>
                        </div>
                    </div>
                </div>
                <div class="top_list_time_content">
                    <span v-show="latestTime">Updated ：{{latestTime}}+UTC</span>
                </div>
            </div>
        </div>
        <div class="top_list_container">
            <div class="top_list_content">
                <div class="top_list_table_wrap">
                    <div class="top_list_table_content">
                        <top-list-table :items="topList" :showNoData="showNoData"></top-list-table>
                        <div v-show="showNoData" class="no_data_show">No Data</div>
                    </div>
                </div>
            </div>
        </div>
        <spin-component :show-loading="showLoading"></spin-component>
    </div>
</template>

<script>
import SpinComponent from "./commonComponents/SpinComponent";
import TopListTable from "./table/TopListTable";
import Server from "../service";
import Tools from "../util/Tools";
export default {
    name: "TopList",
    components: { SpinComponent, TopListTable },
    data() {
        return {
            topList: [],
            showNoData: false,
            showLoading: false,
            latestTime: "",
            richListTimer: null
        };
    },
    mounted() {
        this.getTopList();
        clearInterval(this.richListTimer);
        let that = this;
        this.richListTimer = setInterval(function() {
            that.getTopList();
        }, 60000);
    },
    methods: {
        getTopList() {
            this.showLoading = true;
            Server.commonInterface({ richListAccounts: {} }, res => {
                try {
                    this.showLoading = false;
                    if (res) {
                        this.latestTime = this.getUpDatedTime(res);
                        this.showLoading = false;
                        this.topList = res.map(item => {
                            return {
                                rank: item.rank,
                                Address: item.address,
                                Balance: `${Tools.formatPrice(
                                    Tools.convertScientificNotation3Number(
                                        Tools.formatNumber(
                                            item.balance[0].amount
                                        )
                                    )
                                )}`,
                                Percentage: this.formatPercentage(item.percent)
                            };
                        });
                    } else {
                        this.showLoading = false;
                        this.showNoData = true;
                        this.topList = [
                            {
                                "#": "",
                                Address: "",
                                Balance: "",
                                Percentage: ""
                            }
                        ];
                    }
                } catch (e) {
                    this.showLoading = false;
                    this.showNoData = true;
                    this.topList = [
                        {
                            "#": "",
                            Address: "",
                            Balance: "",
                            Percentage: ""
                        }
                    ];
                    console.error(e);
                }
            });
        },
        formatPercentage(percentage) {
            let minToFixedNumber = 0.0001,
                maxSubStrLength = 8;
            //科学计数法转成数字
            let percentageNumber = percentage
                .toExponential()
                .match(/\d(?:\.(\d*))?e([+-]\d+)/);
            let maxToFixedNumber = 20;
            //toFixed function biggest parameters 20
            let FixedNumber =
                Math.max(
                    0,
                    (percentageNumber[1] || "").length - percentageNumber[2]
                ) > maxToFixedNumber
                    ? maxToFixedNumber
                    : Math.max(
                          0,
                          (percentageNumber[1] || "").length -
                              percentageNumber[2]
                      );
            let toFixedNumber = (
                percentage
                    .toFixed(FixedNumber)
                    .toString()
                    .substring(0, maxSubStrLength) * 100
            ).toFixed(4);
            if (toFixedNumber < minToFixedNumber) {
                return "< 0.0001";
            } else {
                return Tools.subStrings(toFixedNumber + '', 4);
            }
        },
        getUpDatedTime(upDatedTime) {
            let maxUpDatedTime = 0;
            upDatedTime.forEach(item => {
                if (item.update_at > maxUpDatedTime) {
                    maxUpDatedTime = item.update_at;
                }
            });
            return new Date(maxUpDatedTime * 1000)
                .toISOString()
                .split(".")[0]
                .replace("T", " ");
        }
    }
};
</script>

<style scoped lang="scss">
@import "../style/mixin.scss";
.top_list_page {
    width: 100%;
    .top_list_title_container {
        width: 100%;
        z-index: 10;
        position: fixed;
        .top_list_title_content {
            max-width: 12.8rem;
            margin: 0 auto;
            display: flex;
            background: #fff;
            align-items: center;
            & > div.top_list_title_content_div {
                display: flex;
                align-items: center;
                .top_list_title {
                    height: 0.7rem;
                    display: flex;
                    align-items: center;
                    font-size: 0.18rem;
                    color: var(--titleColor);
                    line-height: 0.26rem;
                }
                .tooltip_icon {
                    width: 0.16rem;
                    height: 0.16rem;
                    margin-left: 0.1rem;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    border: 1px solid var(--bgColor);
                    border-radius: 50%;
                    font-size: 0.1rem;
                    color: var(--bgColor);
                    position: relative;
                    cursor: pointer;
                    &:hover .tooltip_span {
                        display: block;
                    }
                    .tooltip_span {
                        width: 2.4rem;
                        display: none;
                        position: fixed;
                        z-index: 100000;
                        margin-top: -20px;
                        transform: translateY(-50%);
                        color: #ffffff;
                        background-color: #000000;
                        border-radius: 0.04rem;
                        word-wrap: break-word;
                        white-space: normal;
                        line-height: 16px;
                        div {
                            padding: 8px 15px;
                        }
                        &::after {
                            width: 0;
                            height: 0;
                            border: 0.06rem solid transparent;
                            content: "";
                            display: block;
                            position: absolute;
                            border-top-color: #000000;
                            left: 50%;
                            margin-left: -6px;
                        }
                    }
                }
            }

            .top_list_time_content {
                padding-left: 0.2rem;
                padding-right: 0.2rem;
                opacity: 0.6;
                color: #000;
                font-size: 0.14rem;
                flex: 1;
                display: flex;
                justify-content: flex-end;
            }
        }
    }
}
.top_list_container {
    width: 100%;
    padding-top: 0.7rem;
    margin-bottom: 0.4rem;
    .top_list_content {
        max-width: 12.8rem;
        width: 100%;
        margin: 0 auto;
        .top_list_table_wrap {
            overflow-x: auto;
            .top_list_table_content {
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
@media (max-width: 910px) {
    .top_list_container {
        padding: 0 0.1rem;
    }
    .top_list_title_content {
        flex-direction: column !important;
        align-items: flex-start !important;
        .top_list_title {
            padding-left: 0.2rem !important;
            height: 0.45rem !important;
        }
        .top_list_time_content {
            line-height: 0.25rem;
        }
    }
}
@media screen and (max-width: 910px){
    .top_list_page {
        .top_list_title_container{
            position: static;
        }
    }
}
</style>
