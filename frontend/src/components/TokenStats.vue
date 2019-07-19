<template>
    <div class="blocks_list_page_wrap">
        <div
            class="list_page_container"
            :class="[$store.state.isMobile ? 'mobile_list_page_container' : 'pc_list_page_container']"
        >
            <div>
                <div class="page_title">IRIS Token Stats</div>
                <div class="table_contanier" v-show="!itemsNoData">
                    <div class="information_props_wrap" v-for="v in items" :key="v.label">
                        <span class="information_props">{{v.label}}</span>
                        <span class="information_value">{{v.value || '--'}}</span>
                    </div>
                </div>
                <div v-show="itemsNoData" class="no_data_show">No Data</div>
                <div class="page_title">IRIS Token Distribution</div>
                <div class="echarts_container" v-show="!pieDatasNoData">
                    <m-token-stats-echart :data="pieDatas"></m-token-stats-echart>
                </div>
                <div v-show="pieDatasNoData" class="no_data_show">No Data</div>
            </div>
        </div>
    </div>
</template>

<script>
import MTokenStatsEchart from "./commonComponents/MTokenStatsEchart";
import Service from "../service";
import Tools from "../util/Tools";

export default {
    components: {
        MTokenStatsEchart
    },
    data() {
        return {
            items: [
                {
                    label: "Total Supply",
                    value: ""
                },
                {
                    label: "Circulation",
                    value: ""
                },
                {
                    label: "Initial Supply",
                    value: ""
                },
                {
                    label: "Burned",
                    value: ""
                },
                {
                    label: "Bonded Tokens",
                    value: ""
                }
            ],
            pieDatas: [],
            itemsNoData: false,
            pieDatasNoData: false
        };
    },
    methods: {
        getTokenStats() {
            return new Promise((resolve, reject) => {
                Service.commonInterface(
                    {
                        tokenStats: {}
                    },
                    result => {
                        try {
                            let data = result.data;
                            if (data) {
                                let obj = [
                                    {
                                        label: "Total Supply",
                                        value: this.$options.filters.amountFromat(
                                            data.totalsupply_tokens,
                                            undefined,
                                            4,
                                            true
                                        )
                                    },
                                    {
                                        label: "Circulation",
                                        value: this.$options.filters.amountFromat(
                                            data.circulation_tokens,
                                            undefined,
                                            4,
                                            true
                                        )
                                    },
                                    {
                                        label: "Initial Supply",
                                        value: this.$options.filters.amountFromat(
                                            data.initsupply_tokens,
                                            undefined,
                                            4,
                                            true
                                        )
                                    },
                                    {
                                        label: "Burned",
                                        value: this.$options.filters.amountFromat(
                                            data.burned_tokens,
                                            undefined,
                                            4,
                                            true
                                        )
                                    },
                                    {
                                        label: "Bonded Tokens",
                                        value: this.$options.filters.amountFromat(
                                            data.delegated_tokens,
                                            undefined,
                                            4,
                                            true
                                        )
                                    }
                                ];
                                this.items = obj;
                            } else {
                                this.itemsNoData = true;
                            }
                            resolve();
                        } catch (err) {
                            this.itemsNoData = true;
                            resolve();
                            console.log(err);
                        }
                    }
                );
            });
        },
        getTokenStatsDistribution() {
            return new Promise((resolve, reject) => {
                Service.commonInterface(
                    {
                        tokenStatsDistribution: {}
                    },
                    result => {
                        try {
                            result = result && result.data;
                            if (result) {
                                let data = Object.entries(result);

                                for (let v of data) {
                                    v[1].totalAmount = this.$options.filters.amountFromat(
                                        v[1].total_amount,
                                        undefined,
                                        4,
                                        true
                                    );
                                    v[1].percentValue = this.formatDecimalNumberToFixedNumber(
                                        Number(v[1].percent) * 100
                                    );
                                }
                                data.sort((a, b) => {
                                    return (
                                        Number(a[0].split("-")[0]) -
                                        Number(b[0].split("-")[0])
                                    );
                                });
                                this.pieDatas = data;
                            } else {
                                this.pieDatasNoData = true;
                            }
                            resolve();
                        } catch (err) {
                            this.pieDatasNoData = true;
                            resolve();
                        }
                    }
                );
            });
        },
        formatDecimalNumberToFixedNumber(num) {
            if (Number(num) < 0.0001) {
                return "<0.0001";
            } else {
                let s = num + "";
                let arr = s.split(".");
                let n = arr[1] ? `${arr[0]}.${arr[1].substring(0, 4)}` : arr[0];
                return n;
            }
        }
    },
    mounted() {
        (async () => {
            try {
                await Promise.all([
                    this.getTokenStats(),
                    this.getTokenStatsDistribution()
                ]);
            } catch (err) {}
        })();
    }
};
</script>

<style lang="scss" scoped>
@import "../style/mixin.scss";

.blocks_list_page_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    .list_page_container {
        width: 100%;
        & > div {
            width: 100%;
            max-width: 12.8rem;
            min-width: 3rem;
            margin: 0 auto 0.4rem;
            .page_title {
                height: 0.7rem;
                line-height: 0.7rem;
                padding-left: 0.2rem;
                font-size: 0.18rem;
            }
            .table_contanier {
                display: flex;
                flex-wrap: wrap;
                .information_props_wrap {
                    font-size: 14px;
                    line-height: 20px;
                    margin-right: 0.2rem;
                    flex: 1 1 0px;
                    display: flex;
                    flex-direction: column;
                    border: 1px solid rgba(215, 217, 224, 1);
                    border-radius: 1px;
                    padding: 0.2rem;
                    .information_props {
                        color: #a2a2ae;
                        font-size: 0.14rem;
                    }
                    .information_value {
                        color: #22252a;
                        font-size: 0.16rem;
                        margin-top: 0.12rem;
                        word-break: break-all;
                        word-wrap: break-word;
                    }
                    .skip_route {
                        a,
                        span {
                            cursor: pointer;
                            color: #3598db !important;
                        }
                    }
                    &:nth-last-of-type(1) {
                        margin-right: 0;
                    }
                }
            }
            .echarts_container {
                padding: 0.2rem;
                border: 1px solid #d7d9e0;
            }
            .no_data_show {
                @include flex;
                justify-content: center;
                border-top: 0.01rem solid #eee;
                border-bottom: 0.01rem solid #eee;
                font-size: 0.14rem;
                height: 2rem;
                align-items: center;
            }
        }
    }
    .mobile_list_page_container {
        & > div {
            padding: 0 0.1rem;
            .page_title {
                padding-left: 0.1rem;
            }
        }
        .table_contanier {
            margin-right: 0 !important;
            .information_props_wrap {
                padding: 0.1rem !important;
                margin-bottom: 0.1rem !important;
                flex: 0 0 100% !important;
                .information_props {
                    width: 100% !important;
                }
                .information_value {
                    width: 100% !important;
                }
                &:nth-last-of-type(1) {
                    margin-bottom: 0 !important;
                }
            }
        }
    }
}
</style>
