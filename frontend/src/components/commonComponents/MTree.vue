<template>
    <div class="treemap_container">
        <div ref="treemapEcharts"></div>
        <div style="width: 50px; height: 50px;">
            <canvas ref="clipImage"></canvas>
        </div>
    </div>
</template>

<script>
import echarts from "echarts";

export default {
    name: "MTree",
    props: {
        items: {
            type: Array,
            default: function() {
                return [];
            }
        }
    },
    watch: {
        items: {
            handler(newVal, oldVal) {
                let data = newVal.map(v => {
                    let random = Math.ceil(Math.random() * 9 + 1);
                    let arr = [];
                    for (let i = 0; i < random; i++) {
                        arr.push({
                            name: i,
                            value: i
                        });
                    }
                    let formatter = v.imageUrl
                        ? `{b|''}  {a|${v.moniker || v.operatorAddress}}`
                        : `{a|${v.moniker || v.operatorAddress}}`;
                    return {
                        name: v.moniker || v.operatorAddress,
                        value: +v.votingPower.replace("%", ""),
                        label: {
                            show: true,
                            formatter: [formatter].join("\n"),
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
                                    borderRadius: 10
                                }
                            }
                        },
                        children: arr
                    };
                });
                this.setOption(data);
            },
            deep: true
        }
    },
    data() {
        return {
            chart: null
        };
    },
    methods: {
        setOption(data) {
            this.chart.setOption({
                title: {},
                tooltip: {},
                series: [
                    {
                        name: "validators",
                        type: "treemap",
                        data: data,
                        leafDepth: 1,
                        drillDownIcon: "",
                        width: "100%",
                        height: 520,
                        top: 0,
                        // roam: false,
                        breadcrumb: {
                            itemStyle: {
                                color: "#3598db",
                                borderColor: "#3598db",
                                borderWidth: 0
                            }
                        },
                        levels: [
                            {
                                itemStyle: {
                                    normal: {
                                        borderColor: "#555"
                                    }
                                },
                                emphasis: {
                                    itemStyle: {
                                        color: "red"
                                    }
                                }
                            },
                            {
                                colorSaturation: [0.3, 0.6],
                                itemStyle: {
                                    normal: {
                                        borderColorSaturation: 0.7
                                    }
                                }
                            }
                        ]
                    }
                ]
            });
        },
        clipImageFunc(url) {
            let image = new Image();
            image.src = url;
            image.onload = () => {
                let ctx = this.clipImageCanvasContext;
                ctx.save();
                ctx.beginPath();
                ctx.arc(25, 25, 25, 0, Math.PI * 2, false);
                ctx.clip(); //剪切路径
                ctx.drawImage(image, 0, 0, 50, 50);
                //恢复状态
                let dataURL = this.clipImage.toDataURL("image/jpg");
                ctx.restore();

                console.log(dataURL);
            };

            // let that = this;
            // var xhr = new XMLHttpRequest();
            // xhr.onload = function() {
            //     var url = URL.createObjectURL(this.response);
            //     var img = new Image();
            //     img.src = url;
            //     img.onload = function() {
            //         that.clipImageCanvasContext.drawImage(img, 0, 0, 50, 50);
            //         let dataURL = that.clipImage.toDataURL("image/jpeg");
            //         URL.revokeObjectURL(url); // 图片用完后记得释放内存
            //     };
            // };
            // xhr.open("GET", url, true);
            // xhr.responseType = "blob";
            // xhr.send();
        }
    },
    mounted() {
        this.clipImage = this.$refs.clipImage;
        this.clipImage.width = 50;
        this.clipImage.height = 50;
        this.clipImageCanvasContext = this.clipImage.getContext("2d");
        this.chart = echarts.init(this.$refs.treemapEcharts);
    }
};
</script>

<style lang="scss" scoped>
.treemap_container {
    width: 12.8rem;
    height: 6rem;
    padding: 0.2rem 0.2rem 0.2rem;
    border: 1px solid #d7d9e0;
    margin-bottom: 0.4rem;
    & > div {
        width: 100%;
        height: 100%;
    }
}
</style>
