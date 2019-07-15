<template>
    <div>
        <div class="treemap_breadcrumb_container">
            <div
                @click="leafDepthTop"
                :class="breadcrumb ? 'breadcrumb_first' : ''"
            >{{validatorType}}</div>
            <div v-if="breadcrumb" @click="setData(items);" class="breadcrumb_last">
                <div>{{breadcrumb}}</div>
            </div>
        </div>
        <div class="treemap_container">
            <div ref="treemapEcharts"></div>
        </div>
    </div>
</template>

<script>
import echarts from "echarts";
import Constants from "../../constant/Constant";

export default {
    name: "MTree",
    props: {
        items: {
            type: Array,
            default: function() {
                return [];
            }
        },
        clickFunc: {
            type: Function,
            default: function() {}
        },
        validatorType: {
            type: String,
            default: ""
        }
    },
    watch: {
        items: {
            handler(newVal, oldVal) {
                this.$nextTick(() => {
                    this.setData(newVal);
                });
            },
            deep: true
        }
    },
    data() {
        return {
            timer: null,
            chart: null,
            breadcrumb: "",
            colors: [
                "hsla(216, 94%, 58%, 1)",
                "hsla(216, 89%, 55%, 1)",
                "hsla(216, 82%, 51%, 1)",
                "hsla(216, 84%, 47%, 1)",
                "hsla(216, 91%, 44%, 1)",
                "hsla(216, 91%, 41%, 1)",

                "hsla(246, 94%, 72%, 1)",
                "hsla(246, 81%, 67%, 1)",
                "hsla(246, 69%, 61%, 1)",
                "hsla(246, 62%, 56%, 1)",

                "hsla(201, 69%, 45%, 1)",
                "hsla(201, 71%, 42%, 1)",
                "hsla(201, 74%, 38%, 1)",
                "hsla(201, 79%, 35%, 1)",
                "hsla(201, 87%, 32%, 1)",
                "hsla(200, 87%, 30%, 1)",

                "hsla(192, 83%, 49%, 1)",
                "hsla(192, 86%, 44%, 1)",
                "hsla(192, 89%, 39%, 1)",
                "hsla(192, 87%, 37%, 1)",
                "hsla(192, 96%, 33%, 1)",
                "hsla(192, 95%, 31%, 1)",

                "hsla(181, 67%, 47%, 1)",
                "hsla(181, 67%, 43%, 1)",
                "hsla(181, 70%, 35%, 1)",
                "hsla(182, 69%, 33%, 1)",
                "hsla(182, 70%, 31%, 1)",
                "hsla(182, 75%, 27%, 1)"
            ],
            series: [
                {
                    name: "validators",
                    type: "treemap",
                    data: [],
                    leafDepth: 1,
                    drillDownIcon: "",
                    width: "100%",
                    height: "100%",
                    bottom: 0,
                    nodeClick: false,
                    breadcrumb: {
                        show: false
                    }
                }
            ],
            optionData: {
                tooltip: {},
                series: null
            }
        };
    },
    methods: {
        formatPerNumber(num) {
            if (typeof num === "number" && !Object.is(num, NaN)) {
                let afterPoint = String(num).split(".")[1];
                let afterPointLong = (afterPoint && afterPoint.length) || 0;
                if (afterPointLong > 2 && num !== 0) {
                    return num.toFixed(4);
                } else {
                    return num.toFixed(2);
                }
            }
            return num;
        },
        setData(newVal) {
            let data = [];
            for (let [i, v] of Object.entries(newVal)) {
                let arr = [];
                i = i % this.colors.length;
                let color = this.colors[i];
                let colorArr = color.split(",");
                let colorHS = `${colorArr[0]},${colorArr[1]}`;
                let colorL = +color.split(",")[2].replace("%", "");
                if (
                    Array.isArray(v.delegations) &&
                    v.delegations.length > 0 &&
                    v.delegationsActive
                ) {
                    let step = 20 / v.delegations.length;

                    arr = v.delegations.map(value => {
                        let selfShares = Number(value.self_shares);
                        // 判断是否是owner address
                        let labelValue =
                            v.ownerAddress === value.address
                                ? v.imageUrl
                                    ? `{b|} {a|${v.moniker ||
                                          v.operatorAddress}}`
                                    : `{a|${v.moniker || v.operatorAddress}}`
                                : value.address;
                        let name =
                            v.ownerAddress === value.address
                                ? v.moniker || v.operatorAddress
                                : value.address;
                        let o = {
                            name: name,
                            value: selfShares,
                            infoData: {
                                name: `${name}`,
                                amount: `Amount: ${this.$options.filters.amountFromat(
                                    value.amount,
                                    Constants.Denom.IRIS.toUpperCase()
                                )}`,
                                shares: `Self Shares: ${selfShares} (${this.formatPerNumber(
                                    (selfShares / Number(value.total_shares)) *
                                        100
                                )}%)`,
                                block: `Block: ${value.block}`
                            },
                            itemStyle: {
                                color: `${colorHS}, ${colorL}%)`
                            },
                            emphasis: {
                                itemStyle: {
                                    color: `${colorHS}, ${colorL + 4}%)`
                                }
                            },
                            label: {
                                normal: {
                                    show: true,
                                    formatter: [labelValue].join("\n"),
                                    verticalAlign: "middle",
                                    rich: {
                                        a: {
                                            color: "#ffffff"
                                        },
                                        b: {
                                            backgroundColor: {
                                                image: v.imageUrl
                                            },
                                            width: 16,
                                            height: 16,
                                            borderRadius: 8,
                                            lineHeight: 16
                                        }
                                    }
                                },
                                emphasis: {
                                    show: true,
                                    formatter: [labelValue].join("\n"),
                                    verticalAlign: "middle",
                                    rich: {
                                        a: {
                                            color: "#ffffff"
                                        },
                                        b: {
                                            backgroundColor: {
                                                image: v.imageUrl
                                            },
                                            width: 16,
                                            height: 16,
                                            borderRadius: 8,
                                            lineHeight: 16
                                        }
                                    }
                                }
                            }
                        };
                        colorL += step;
                        return o;
                    });
                    this.optionData.tooltip = {
                        show: true,
                        confine: true,
                        formatter: function(v) {
                            let info = v && v.data && v.data.infoData;
                            if (info) {
                                return `${info.name}<br/>${info.amount}<br/>${info.shares}<br/>${info.block}`;
                            }
                        }
                    };
                    this.breadcrumb = v.moniker || v.operatorAddress;
                    return this.setOption(arr);
                }
                this.breadcrumb = "";
                let labelValue = v.imageUrl
                    ? `{b|} {a|${v.moniker || v.operatorAddress}}`
                    : `{a|${v.moniker || v.operatorAddress}}`;
                let validator = {
                    name: v.moniker || v.operatorAddress,
                    value: +v.votingPower,
                    info: v,
                    itemStyle: {
                        color: color
                    },
                    emphasis: {
                        itemStyle: {
                            color: `${colorHS}, ${colorL + 4}%)`
                        }
                    },
                    label: {
                        normal: {
                            show: true,
                            formatter: [labelValue].join("\n"),
                            verticalAlign: "middle",
                            rich: {
                                a: {
                                    color: "#ffffff"
                                },
                                b: {
                                    backgroundColor: {
                                        image: v.imageUrl
                                    },
                                    width: 16,
                                    height: 16,
                                    borderRadius: 8,
                                    lineHeight: 16
                                }
                            }
                        },
                        emphasis: {
                            show: true,
                            formatter: [labelValue].join("\n"),
                            verticalAlign: "middle",
                            rich: {
                                a: {
                                    color: "#ffffff"
                                },
                                b: {
                                    backgroundColor: {
                                        image: v.imageUrl
                                    },
                                    width: 16,
                                    height: 16,
                                    borderRadius: 8,
                                    lineHeight: 16
                                }
                            }
                        }
                    }
                };
                data.push(validator);
            }
            this.optionData.tooltip = {
                show: true,
                confine: true,
                formatter: function(v) {
                    return `${v.name}<br/> 
                    voting_power: ${v.value}`;
                }
            };
            this.setOption(data);
        },
        setOption(data) {
            this.forSeries(data);
            this.optionData.series = this.series;
            this.chart.setOption(this.optionData);
        },
        forSeries(data) {
            this.series[0].data = data;
        },
        echartsClick(data) {
            let delegations =
                data &&
                data.data &&
                data.data.info &&
                data.data.info.delegations;
            if (!delegations) {
                this.clickFunc(data);
            } else {
                let operatorAddress =
                    data &&
                    data.data &&
                    data.data.info &&
                    data.data.info.operatorAddress;
                if (operatorAddress) {
                    let d = this.items.find(
                        v => v.operatorAddress === operatorAddress
                    );
                    d.delegationsActive = true;
                }
            }
            if (data && data.data && !data.data.info) {
                //delegations 跳转到address页面
                let routeUrl = this.$router.resolve({
                    path: `/address/${data.name}`
                });
                window.open(routeUrl.href, "_blank");
            }
        },
        leafDepthTop() {
            this.items.forEach(v => (v.delegationsActive = false));
            this.setData(this.items);
        },
        computedTreemapEchartsHeight(node) {
            if (this.$store.state.isMobile) {
                let h = Math.round(node.offsetWidth * 1.618) - 60;
                node.style.height = `${h}px`;
            } else {
                let h = Math.round(node.offsetWidth / 1.618) - 60;
                node.style.height = `${h}px`;
            }
        },
        windowResizeFunc() {
            clearTimeout(this.timer);
            this.timer = setTimeout(() => {
                let node = this.$refs.treemapEcharts.parentNode;
                this.computedTreemapEchartsHeight(node);
                this.chart.resize();
            }, 500);
        }
    },
    mounted() {
        let node = this.$refs.treemapEcharts.parentNode;
        this.computedTreemapEchartsHeight(node);
        this.chart = echarts.init(this.$refs.treemapEcharts);

        window.addEventListener("resize", this.windowResizeFunc);
        this.chart.on("click", this.echartsClick);
    },
    destroyed() {
        window.removeEventListener("resize", this.windowResizeFunc);
        this.chart.off("click", this.echartsClick);
    }
};
</script>

<style lang="scss" scoped>
.treemap_breadcrumb_container {
    margin: 0.2rem 0.2rem 0;
    display: flex;
    & > div {
        line-height: 0.2rem;
        background-color: #3598da;
        color: #ffffff;
        padding: 0 0.1rem;
        cursor: pointer;
        margin-right: 0.08rem;
        position: relative;
    }
    .breadcrumb_first {
        padding-right: 0.04rem;
        &::after {
            height: 0;
            width: 0;
            display: inline-block;
            content: "";
            border-width: 0.1rem 0.08rem 0.1rem;
            border-color: transparent;
            border-style: solid;
            border-left-color: #3598da;
            position: absolute;
            top: 0;
            right: -16px;
        }
    }
    .breadcrumb_last {
        padding-right: 0.04rem;
        max-width: calc(100% - 82px);
        & > div {
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
        }
        &::before {
            height: 0;
            width: 0;
            display: inline-block;
            content: "";
            border-width: 0.1rem 0.08rem 0.1rem;
            border-color: transparent;
            border-style: solid;
            border-left-color: #ffffff;
            position: absolute;
            top: 0;
            left: 0;
        }
        &::after {
            height: 0;
            width: 0;
            display: inline-block;
            content: "";
            border-width: 0.1rem 0.08rem 0.1rem;
            border-color: transparent;
            border-style: solid;
            border-left-color: #3598da;
            position: absolute;
            top: 0;
            right: -16px;
        }
    }
}
.treemap_container {
    width: 100%;
    height: 60vw;
    padding: 0.2rem;
    & > div {
        width: 100%;
        height: 100%;
    }
}
</style>
