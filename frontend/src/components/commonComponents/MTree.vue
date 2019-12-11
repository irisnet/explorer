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
var echarts = require('echarts/lib/echarts')
require("echarts/lib/chart/treemap")
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
            colors: ["hsla(204, 91%, 63%, 1)"],
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
                if (num < 0.0001) {
                    return "<0.0001";
                } else {
                    let s = num + "";
                    let arr = s.split(".");
                    let n = arr[1]
                        ? `${arr[0]}.${arr[1].substring(0, 4)}`
                        : arr[0];
                    return n;
                }
            }
        },
        setData(newVal) {
            let data = [];
            for (let [i, v] of Object.entries(newVal)) {
                let arr = [];
                i = i % this.colors.length;
                let color = this.colors[i];
                let colorArr = color.split(",");
                let colorHS = `${colorArr[0]},${colorArr[1]}`;
                let colorL = +color.split(",")[2].replace("%", "") - 10;
                if (
                    Array.isArray(v.delegations) &&
                    v.delegations.length > 0 &&
                    v.delegationsActive
                ) {
                    let step = 15 / v.delegations.length;

                    arr = v.delegations.map(value => {
                        let selfShares = value.selfSharesFormat;
                        // 判断是否是owner address
                        let labelValue =
                            v.ownerAddress === value.address
                                ? v.imageUrl
                                    ? `{b|}{c|} {a|${v.moniker ||
                                          v.operatorAddress}}`
                                    : `{a|${v.moniker || v.operatorAddress}}`
                                : value.address;
                        let name =
                            v.ownerAddress === value.address
                                ? v.moniker || v.operatorAddress
                                : value.address;
                        let selfSharesPer =
                            Number(value.self_shares) /
                            Number(value.total_shares);
                        let labelOffset =
                            v.ownerAddress === value.address ? [18, 0] : [0, 0];
                        let o = {
                            name: name,
                            value:
                                selfSharesPer < 0.002 ? 0.002 : selfSharesPer,
                            infoData: {
                                name: `${name}`,
                                amount: `Amount: ${this.$options.filters.amountFromat(
                                    value.amount,
                                    Constants.Denom.IRIS.toUpperCase()
                                )}`,
                                shares: `Shares: ${selfShares} (${this.formatPerNumber(
                                    selfSharesPer * 100
                                )}%)`,
                                ownerAddress:
                                    v.ownerAddress === value.address
                                        ? v.operatorAddress
                                        : "",
                                operatorAddress: value.address
                            },
                            itemStyle: {
                                color: `${colorHS}, ${colorL}%, 1)`
                            },
                            emphasis: {
                                itemStyle: {
                                    color: `${colorHS}, ${colorL + 4}%, 1)`
                                }
                            },
                            label: {
                                normal: {
                                    show: true,
                                    formatter: [labelValue].join("\n"),
                                    verticalAlign: "middle",
                                    offset: labelOffset,
                                    rich: {
                                        a: {
                                            color: "#ffffff"
                                        },
                                        b: {
                                            backgroundColor: {
                                                image: v.imageUrl
                                            },
                                            width: 12,
                                            height: 12,
                                            borderRadius: 8,
                                            lineHeight: 16,
                                            padding: [2, -20, 2, -13],
                                            verticalAlign: "middle"
                                        },
                                        c: {
                                            backgroundColor: "transparent",
                                            width: 22,
                                            height: 22,
                                            borderRadius: 10,
                                            borderWidth: 6,
                                            borderColor: `${colorHS}, ${colorL}%, 1)`,
                                            verticalAlign: "middle"
                                        }
                                    }
                                },
                                emphasis: {
                                    show: true,
                                    formatter: [labelValue].join("\n"),
                                    verticalAlign: "middle",
                                    offset: labelOffset,
                                    rich: {
                                        a: {
                                            color: "#ffffff"
                                        },
                                        b: {
                                            backgroundColor: {
                                                image: v.imageUrl
                                            },
                                            width: 12,
                                            height: 12,
                                            borderRadius: 8,
                                            lineHeight: 16,
                                            padding: [2, -20, 2, -13],
                                            verticalAlign: "middle"
                                        },
                                        c: {
                                            backgroundColor: "transparent",
                                            width: 22,
                                            height: 22,
                                            borderRadius: 10,
                                            borderWidth: 6,
                                            borderColor: `${colorHS}, ${colorL +
                                                4}%, 1)`,
                                            verticalAlign: "middle"
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
                                return `${info.name}<br/>${info.amount}<br/>${info.shares}`;
                            }
                        }
                    };
                    this.breadcrumb = v.moniker || v.operatorAddress;
                    return this.setOption(arr);
                }
                this.breadcrumb = "";
                let labelValue = v.imageUrl
                    ? `{b|}{c|} {a|${v.moniker || v.operatorAddress}}`
                    : `{a|${v.moniker || v.operatorAddress}}`;
                let votingPowerPer = +v.votingPower / v.allVotingPower;
                let validator = {
                    name: v.moniker || v.operatorAddress,
                    value: votingPowerPer,
                    info: v,
                    itemStyle: {
                        color: color,
                        borderWidth: 0.5
                    },
                    emphasis: {
                        itemStyle: {
                            color: `${colorHS}, ${colorL + 4}%, 1)`,
                            borderWidth: 0.5
                        }
                    },
                    label: {
                        normal: {
                            show: true,
                            formatter: [labelValue].join("\n"),
                            verticalAlign: "middle",
                            offset: [18, 0],
                            rich: {
                                a: {
                                    color: "#ffffff"
                                },
                                b: {
                                    backgroundColor: {
                                        image: v.imageUrl
                                    },
                                    width: 12,
                                    height: 12,
                                    borderRadius: 8,
                                    lineHeight: 16,
                                    padding: [2, -20, 2, -13],
                                    verticalAlign: "middle"
                                },
                                c: {
                                    backgroundColor: "transparent",
                                    width: 22,
                                    height: 22,
                                    borderRadius: 10,
                                    borderWidth: 6,
                                    borderColor: color,
                                    verticalAlign: "middle"
                                }
                            }
                        },
                        emphasis: {
                            show: true,
                            formatter: [labelValue].join("\n"),
                            verticalAlign: "middle",
                            offset: [18, 0],
                            rich: {
                                a: {
                                    color: "#ffffff"
                                },
                                b: {
                                    backgroundColor: {
                                        image: v.imageUrl
                                    },
                                    width: 12,
                                    height: 12,
                                    borderRadius: 8,
                                    lineHeight: 16,
                                    padding: [2, -20, 2, -13],
                                    verticalAlign: "middle"
                                },
                                c: {
                                    backgroundColor: "transparent",
                                    width: 22,
                                    height: 22,
                                    borderRadius: 10,
                                    borderWidth: 6,
                                    borderColor: `${colorHS}, ${colorL +
                                        4}%, 1)`,
                                    verticalAlign: "middle"
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
                    voting_power: ${v.data &&
                        v.data.info &&
                        v.data.info.votingPower}`;
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
            this.series[0].roam = !this.$store.state.isMobile;
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
            if (data && data.data && data.data.infoData) {
                //delegations 跳转到address页面
                let routeUrl = this.$router.resolve({
                    path: `/address/${data.data.infoData.operatorAddress}`
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
        padding-right: 0;
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
        }
    }
    .breadcrumb_last {
        padding-right: 0;
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
