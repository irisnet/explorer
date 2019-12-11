<template>
    <div>
        <div ref="TokenStatsEchartContainer" class="echart_container"></div>
    </div>
</template>

<script>
var echarts  = require('echarts/lib/echarts')
require('echarts/lib/chart/pie')
import skinStyle  from "../../skinStyle";
import Constant from "../../constant/Constant"
export default {
    name: "MTokenStatsEchart",
    props: {
        data: {
            type: Array,
            default: function() {
                return [];
            }
        }
    },
    data() {
        return {
            chart: null,
            timer: null
        };
    },
    watch: {
        data(newVal) {
            this.setOption(newVal);
        }
    },
    methods: {
        setOption(newVal) {
            let data = newVal.map(v => {
                return {
                    name: this.$store.state.isMobile
                        ? `No. ${v[0]}:{d|${v[1].totalAmount} (${
                              v[1].percentValue
                          }%)}`
                        : `No. ${v[0]}`,
                    value: v[1].percent,
                    info: v[1],
                    nameValue: `No. ${v[0]}`
                };
            });
            let option = {
                tooltip: {
                    show: this.$store.state.isMobile,
                    confine: true,
                    formatter: function(v) {
                        return `${v.marker}${v.data.nameValue}:<br/>
                        ${v.data.info.totalAmount} (${v.data.info.percentValue}%)`;
                    }
                },
               /* color: [
                    "#3598DB",
                    "#FFA85C",
                    "#C0CEFF",
                    "#FFCF3A",
                    "#A69DFF",
                    "#93ED56",
                    "#FF837E"
                ],*/
                legend: {
                    orient: "vertical",
                    x: "left",
                    selectedMode: false,
                    data: data.map(v => {
                        return {
                            name: v.name,
                            textStyle: {
                                rich: {
                                    d: {
                                        color: "var(--contentColor)",
                                        padding: [0, 0, 0, 6]
                                    }
                                }
                            }
                        };
                    })
                },
                series: [
                    {
                        type: "pie",
                        radius: ["30%", "60%"],
                        avoidLabelOverlap: true, //是否启用防止标签重叠策略
                        minAngle: 5,
                        center: [
                            "50%",
                            this.$store.state.isMobile
                                ? 250 + data.length * 10
                                : "55%"
                        ],
                        hoverOffset: this.$store.state.isMobile ? 10 : 0,
                        label: {
                            normal: {
                                show: !this.$store.state.isMobile,
                                formatter: function(params) {
                                    let text = `${params.data.info.totalAmount} (${params.data.info.percentValue}%)`;
                                    let minLen = 100;
                                    if (text.length > minLen) {
                                        let len = Math.ceil(
                                            text.length / minLen
                                        );
                                        let s = "";
                                        for (let i = 0; i < len; i++) {
                                            s +=
                                                text.substring(
                                                    i * minLen,
                                                    (i + 1) * minLen
                                                ) + "\n";
                                        }
                                        return `\n{b|${params.name}}\n{per|${s}}`;
                                    } else {
                                        return `\n{b|${params.name}}\n{per|${text}}`;
                                    }
                                },
                                position: "outside",
                                align: "left",
                                rich: {
                                    b: {
                                        fontSize: 14,
                                        lineHeight: 20,
                                        color: "var(--contentColor)",
                                        align: "left"
                                    },
                                    per: {
                                        fontSize: 14,
                                        lineHeight: 20,
                                        color: "var(--contentColor)",
                                        align: "left"
                                    }
                                }
                            },
                            emphasis: {
                                show: !this.$store.state.isMobile
                            }
                        },
                        labelLine: {
                            normal: {
                                show: !this.$store.state.isMobile,
                                length: 30,
                                length2: 30,
                                lineStyle: {
                                    color: "#979797"
                                }
                            },
                            emphasis: {
                                show: !this.$store.state.isMobile
                            }
                        },
                        data: data
                    }
                ]
            };
            option.color= [
                "#FFA85C",
                "#C0CEFF",
                "#FFCF3A",
                "#A69DFF",
                "#93ED56",
                "#FF837E"
            ];
            if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.IRISHUB){
                option.color.unshift(skinStyle.skinStyle.MAINNETBGCOLOR)
            }else if(this.$store.state.currentSkinStyle === Constant.CHAINID.FUXI){
                option.color.unshift(skinStyle.skinStyle.TESTNETBGCOLOR)
            }else if(this.$store.state.currentSkinStyle === Constant.CHAINID.NYANCAT){
                option.color.unshift(skinStyle.skinStyle.NYANCATTESTNETBGCOLOR)
            }else {
                option.color.unshift(skinStyle.skinStyle.DEFAULTBGCOLOR)
            }

            this.chart.setOption(option);
        },
        windowResizeFunc() {
            clearTimeout(this.timer);
            this.timer = setTimeout(() => {
                this.chart.resize();
                this.setOption(this.data);
            }, 500);
        }
    },
    mounted() {
        this.chart = echarts.init(this.$refs.TokenStatsEchartContainer);
        window.addEventListener("resize", this.windowResizeFunc);
    },
    destroyed() {
        window.removeEventListener("resize", this.windowResizeFunc);
    }
};
</script>

<style lang="scss" scoped>
.echart_container {
    height: 5rem;
}
</style>
