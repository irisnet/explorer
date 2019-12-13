<template>
    <div :class="echartsComponentWrapLine">
        <div class="echarts_title_wrap_line">
            <span>Uptime</span>
            <slot name="tabs"></slot>
        </div>
        <div id="echarts_uptime_line">
        </div>
    </div>
</template>

<script>
var echarts = require('echarts/lib/echarts')
require('echarts/lib/chart/line')
let line = null;
export default {
    name: 'echarts-uptime-line',
    watch: {
        informationUptimeLine (informationUptimeLine) {
            //根据设备大小显示饼图的大小
            let radius = this.deviceType === 1 ? '85%' : '65%';
            let option = {
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        axis: "x",
                        type: "line",
                        lineStyle: {
                            color: "var(--contentColor)",
                        },
                    },
                    formatter (params) {
                        let res;
                        if (params[0].name.indexOf(":") != -1) {
                            res = `<span style="display:block;">${params[0].name} +UTC</span>`;
                            res += `<span style="display:block;">Average Uptime: ${params[0].value}%</span>`;
                            return res;
                        } else {
                            res = `<span style="display:block;">${params[0].name} </span>`;
                            res += `<span style="display:block;">Average Uptime: ${params[0].value}%</span>`;
                        }
                        return res;
                    }
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: [],
                    axisLine: {
                        lineStyle: {
                            color: 'var(--contentColor)',
                        }
                    },
                    axisLabel: {
                        color: "#000",
                        interval: 5,
                        margin: 12,
                        rotate: 45,
                        formatter: (value) => {
                            if (value.split(":")[1] == "00") {
                                return value
                            } else {
                                return `${this.month[this.monthNum.findIndex(item => value.substr(0, 2) === item)]}${value.substr(3, 2)}`;
                            }
                        }
                    }
                },
                yAxis: {
                    type: 'value',
                    axisLine: {
                        lineStyle: {
                            color: '#a2a2ae',
                        }
                    },
                    axisLabel: {
                        color: "#000",
                    },
                    min: 0,
                    splitNumber: 5,
                    splitLine: {
                        lineStyle: {
                            // 使用深浅的间隔色
                            color: ['#eee']
                        }
                    }
                },
                grid: {
                    left: '3%',   //图表距边框的距离
                    right: '5%',
                    bottom: '3%',
                    top: '5%',
                    containLabel: true
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        areaStyle: {
                            normal: {
                                color: new echarts.graphic.LinearGradient(//设置渐变颜色
                                    0, 0, 0, 1,
                                    [
                                        { offset: 0, color: '#a4d7f4' },
                                        { offset: 1, color: '#dcf6ff' }
                                    ]
                                )
                            }
                        },
                        smooth: true,//曲线平滑
                        itemStyle: {
                            normal: {
                                color: '#3598db',
                                borderColor: '#3598db',  //拐点边框颜色
                                // opacity:0,拐点显示隐藏 0为隐藏
                            }
                        },
                    }
                ]
            };
            if (line) {
                if (informationUptimeLine.noDatayAxisDefaultMax === "100") {
                    option.yAxis.max = "100"
                } else {
                    option.yAxis.max = null
                }
                // 计算min坐标轴
                let min = Math.min.apply(null, informationUptimeLine.seriesData);
                min = Math.floor(min / 10) * 10;
                let max = option.yAxis.max || Math.max.apply(null, informationUptimeLine.seriesData);
                min = min == max ? min - 10 : min;
                min = min < 0 ? 0 : min;
                option.yAxis.min = min;

                option.xAxis.data = informationUptimeLine.xData;
                option.series[0].data = informationUptimeLine.seriesData;
                line.setOption(option)
            }
        }
    },
    data () {
        return {
            deviceType: window.innerWidth > 500 ? 1 : 0,
            echartsComponentWrapLine: window.innerWidth > 500 ? 'echarts_component_wrap_line_personal_computer' : 'echarts_component_wrap_line_mobile',
            month: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
            monthNum: ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12'],

        }
    },
    props: ['informationUptimeLine'],
    beforeMount () {

    },
    mounted () {
        line = echarts.init(document.getElementById('echarts_uptime_line'));
        window.addEventListener('resize', this.onWindowResize)
    },

    methods: {
        onWindowResize () {
            line.resize();
        }
    },
    beforeDestroy () {
        window.removeEventListener('resize', this.onWindowResize);
    }
}
</script>
<style lang="scss" scoped>
@import "../../style/mixin";
.echarts_component_wrap_line_personal_computer{
   background: #fff;
}
.echarts_component_wrap_line_personal_computer,
.echarts_component_wrap_line_mobile {
    padding: 0;
    margin: 0 !important;
    padding-top: 0.3rem;
    .echarts_title_wrap_line {
        height: 0.3rem;
        font-size: 0.18rem;
        @include fontWeight;
        line-height: 0.53rem;
        padding: 0 0.3rem 0 0.2rem;
        display: flex;
        justify-content: space-between;
        & > span {
            color: var(--titleColor);
            font-size: 16px;
            line-height: 0.3rem;
        }
    }
    #echarts_uptime_line {
        width: 100%;
        height: 2.9rem;
    }
}
.echarts_component_wrap_line_mobile {
    min-width: 4rem;
}
</style>
