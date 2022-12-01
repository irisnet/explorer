<template>
    <div class="validator_graph_container" :class="switchValue ? 'start_sky_img' : 'default_bg_img'">
        <div class="graph_content_wrap">
            <div class="graph_content_title">
                <div>
                    <span v-if="flShowNetwork">Interchain Network</span>
                    <span class="beat_content" v-if="!flShowTestTooltip">Beta </span>
                </div>
                <div class="tooltip" v-if="flShowNetwork">
                    <div class="graph_tooltip">
                        <p><span></span><span>Connection Opened</span></p>
                        <p><span></span><span>Connection Unopened</span></p>
                    </div>
                    <!--<div class="background_toggle_content">
						 <el-switch v-model="switchValue"
									active-color="rgba(55, 124, 248, 1)"
									inactive-color="rgba(16, 19, 55, 1)"
									@change="switchBgColor(switchValue)"></el-switch>
						<span style="color: rgba(115, 122, 174, 1);font-size: 0.14rem;margin-left:0.1rem;line-height: 0.16rem">Starry-sky Mode</span>
					</div>-->
                </div>
            </div>
            <div class="graph_charts_container" :class="flShowNetwork ? '' : 'show_error_content'">
                <div class="graph_content_container" v-if="flShowNetwork">
                    <span class="iconfont icon-fuwei" v-if="flShowRevertIcon" @click="revertGraph()"></span>
                    <div class="graph_container_content_wrap" :style="{ background: switchValue ? 'transparent' : '#2d325a' }">
                        <!--<p class="graph_charts_title" v-if="flShowNetwork && !flShowTestTooltip"></p>-->
                        <div id="validator_graph_content" ref="chart_content"></div>
                    </div>
                    <div class="graph_list_container" id="graph_list_container" :style="{ background: switchValue ? '' : '#2d325a' }" :class="[flShowDefaultPickUp ? 'hide_pick_up_style' : '', showDefaultAnimation ? 'show_pick_up_style' : '']" v-show="colorDataArray.length > 0">
                        <div class="sort_content">
                            <span :class="flConnectionActiveStyle ? 'active_style' : ''" @click="sortByConnection(sortByConnectionSwitchIcon)">Connections <i class="iconfont" v-show="flShowConnectionSortIcon" :class="sortByConnectionSwitchIcon ? 'iconxiangxia' : 'iconxiangshang'"></i></span>
                            <span :class="flLetterActiveStyle ? 'active_style' : ''" @click="sortByLetter(sortByLetterSwitchIcon)">Chains <i class="iconfont" v-show="flShowLetterSortIcon" :class="sortByLetterSwitchIcon ? 'iconxiangxia' : 'iconxiangshang'"></i></span>
                        </div>
                        <div class="sort_scroll_content" ref="graph_list_content">
                            <div class="graph_list_item_all" @click="selectAll()">
                                <div class="legend_all_block">
                                    <img v-show="flAllCheckout" src="../../assets/select_all.svg" alt="" />
                                    <img v-show="!flAllCheckout" src="../../assets/unselect_all.svg" alt="" />
                                </div>
                                <span class="legend_name" :class="flAllCheckout ? 'hide_style' : ''">All</span>
                            </div>
                            <ul class="graph_list_content">
                                <li class="graph_list_item" v-for="(item, index) in colorDataArray" @click="changeChart(item, index)">
                                    <el-image v-if="item.icon" :src="item.icon" :class="item.isActive ? '' : 'img_hide_style'">
                                        <div slot="placeholder" class="image-slot">
                                            <img class="legend_block_img_loading" :src="Constant.LOADING.LOADING_IMG" alt="">
                                        </div>
                                    </el-image>
                                    <div v-else class="legend_block_img" :class="item.isActive ? '' : 'img_hide_style'" :style="{ background: item.color }" ></div>
                                    <div class="legend_name_content" :class="item.isActive ? '' : 'hide_style_color'">
                                        <el-tooltip v-if="prettyNamePopover(item.name)" effect="dark" :content="formatPrettyName(item.name)">
                                            <span class="legend_item">{{formatPrettyName(item.name)}}</span>
                                        </el-tooltip>
                                        <span v-else class="legend_item">{{formatPrettyName(item.name)}}</span>
                                        <el-tooltip v-if="currentChainIdPopover(item.name)" effect="dark" :content="formatCurrentChainId(item.name)">
                                            <span class="legend_item">{{formatCurrentChainId(item.name)}}</span>
                                        </el-tooltip>
                                        <span v-else class="legend_item">{{formatCurrentChainId(item.name)}}</span>
                                        <span class="legend_item">{{`${item.connection} ${item.connection > 1 ? "Connections": "Connection"}`}}</span>
                                    </div>
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div class="pick_up_content" v-if="colorDataArray.length > 0">
                        <div class="pick_up_img_content" @click="flPickUp()">
                            <img v-show="!flShowDefaultPickUp" :src="switchValue ? starLeftImg : defaultLeftImg" alt="" />
                        </div>
                    </div>
                    <div class="pick_up_show_content">
                        <div class="pick_up_show_img_content" @click="flShowPickUp()">
                            <img v-show="flShowDefaultRightPickUp" :src="switchValue ? starRightImg : defaultRightImg" alt="" />
                        </div>
                    </div>
                </div>
                <div class="error_content" v-if="!flShowNetwork">
                    <error-component :error-text="Constant.ERROR.ERROR_TEXT" :error-btn-text="Constant.ERROR.ERROR_BTN_TEXT" @refresh="refreshPage()" />
                </div>
                <!-- <div class="hover_up_content" v-if="flShowRevertIcon" v-show="flShowHoverUp" @click="scrollBottom()">
					<img style="width: 0.16rem;height:0.22rem;" src="../../assets/hover_up.gif" alt="">
				</div> -->
            </div>
            <!--<app-download></app-download>-->
            <!--<validator-bianjie-information></validator-bianjie-information>-->
        </div>
    </div>
</template>

<script>
	import axios from "axios";
	import apiUrlConfig from "../../../config/config";
    import Constant from "../../constant/Constant";
    import urlApi from '../../api';
    import Service from '../../util/axios';
	import ValidatorBianjieInformation from "./ValidatorBianjieInformation";
    import AppDownload from "@/components/home/AppDownload";
    import ErrorComponent from "@/components/errorComponent/ErrorComponent";
	import Store from "@/store";
	var echarts = require("echarts/lib/echarts");
	require("echarts/lib/component/legend");
	require("echarts/lib/component/tooltip");
	require("echarts/lib/component/title");
	require("echarts/lib/chart/graph");
	require("echarts/extension/dataTool");
    require("echarts/lib/component/legendScroll");
    import Tools from '@/util/Tools';
export default {
    name: "ValidatorGraph",
    components: { AppDownload, ValidatorBianjieInformation, ErrorComponent },
    data() {
        return {
            Constant,
            Tools,
            graphContent: "",
            flConnectionActiveStyle: true,
            flLetterActiveStyle: false,
            sortByConnectionSwitchIcon: true,
            sortByLetterSwitchIcon: false,
            flShowConnectionSortIcon: true,
            flShowLetterSortIcon: false,
            flAllCheckout: true,
            flShowDefaultPickUp: false,
            flShowDefaultRightPickUp: false,
            showDefaultAnimation: false,
            defaultLeftImg: require("../../assets/default_pick_up_left.png"),
            defaultRightImg: require("../../assets/default_pick_up_right.png"),
            starLeftImg: require("../../assets/star_pick_up_left.png"),
            starRightImg: require("../../assets/star_pick_up_right.png"),
            flShowNetwork: true,
            data: null,
            flShowRevertIcon: false,
            colorDataArray: [],
            sortCopyData: [],
            copyData: "",
            flShowTestTooltip: true,
            dataTimer: null,
            flShowHoverUp: true,
            switchValue: sessionStorage.getItem("starSky") ? (sessionStorage.getItem("starSky") === "show" ? true : false) : true,
            colorUseCopyData: "",
            timer: null,
            dataIndex: 2,
            maxLinks: 0,
            graphEcharts: null,
            linearColor: ["#35B1EF", "#DA6457", "#125EEB"],
            color: ["#CAA4FF", "#35B1EF", "#DA6457", "#125EEB", "#FBB27A", "#FECCB3", "#CFC2EC", "#F2D4EE", "#44C8ED", "#FF7CAD", "#AECFEB", "#FA9084", "#DF94AC", "#F2B8A9", "#BAC3FB", "#F3EFE9", "#DEE7EC", "#AFD3FC", "#F3D3EE", "#B8F28C", "#F27FCB", "#F7F6F7", "#4BB5FD", "#D3FCFC", "#FDAB74", "#FBDC94", "#75C6FF", "#ADDFF6", "#E8BDAF", "#ECE1D8", "#90B5EF", "#B2A9FF", "#18D6E4", "#CED8E7", "#E6C8F4", "#D3A8F4", "#DBD9D4", "#C4F3C9", "#77E9D6", "#4FDBC1", "#F19A6C", "#BED7F3", "#FDB2AD", "#FD8F8C", "#55B6E9", "#97BCFE", "#CF82E5", "#6ED1E7", "#FF9689", "#B7C6DA", "#66E9CE", "#E6FFC4", "#98EEFF", "#FFBEEE", "#58E2FF", "#B0FFB9", "#60F5BF", "#BBDDFF", "#EA98FF", "#AACCFF", "#D69EFD", "#FFCF9E", "#FFAB9F", "#FF839B", "#57D6FF", "#8CC2FF", "#CDBBFF", "#EB89DF", "#E6FFFE", "#F7FEFF", "#F9EBF9", "#6FDAFF", "#B5C9FF", "#AEFFE9", "#71E7FF", "#D3E2FF", "#C4DFFF", "#88FFFF", "#59DCFF", "#D293FF", "#EBE1FF", "#ADFAE1", "#72F4D2", "#29B5FF", "#85F4DF", "#4CDDF8", "#729EFF", "#E0CFFF", "#FF94EF", "#FF79CC", "#8FBEFF", "#64FFEE", "#D0B2FF", "#FFAA93", "#B2B8FF", "#33D4E0", "#DBE2FF", "#F193F8", "#FACE29", "#FF9A53", "#FDB6B2", "#DBAEF9", "#FBD0D6", "#FFA7B1", "#FBB27A", "#FECCB3", "#CFC2EC", "#F2D4EE", "#44C8ED", "#FF7CAD", "#AECFEB", "#FA9084", "#DF94AC", "#F2B8A9", "#BAC3FB", "#F3EFE9", "#DEE7EC", "#AFD3FC", "#F3D3EE", "#B8F28C", "#F27FCB", "#F7F6F7", "#4BB5FD", "#D3FCFC", "#FDAB74", "#FBDC94", "#75C6FF", "#ADDFF6", "#E8BDAF", "#ECE1D8", "#90B5EF", "#B2A9FF", "#18D6E4", "#CED8E7", "#E6C8F4", "#D3A8F4", "#DBD9D4", "#C4F3C9", "#77E9D6", "#4FDBC1", "#F19A6C", "#BED7F3", "#FDB2AD", "#FD8F8C", "#55B6E9", "#97BCFE", "#CF82E5", "#6ED1E7", "#FF9689", "#B7C6DA", "#66E9CE", "#E6FFC4", "#98EEFF", "#FFBEEE", "#58E2FF", "#B0FFB9", "#60F5BF", "#BBDDFF", "#EA98FF", "#AACCFF", "#D69EFD", "#FFCF9E", "#FFAB9F", "#FF839B", "#57D6FF", "#8CC2FF", "#CDBBFF", "#EB89DF", "#E6FFFE", "#F7FEFF", "#F9EBF9", "#6FDAFF", "#B5C9FF", "#AEFFE9", "#71E7FF", "#D3E2FF", "#C4DFFF", "#88FFFF", "#59DCFF", "#D293FF", "#EBE1FF", "#ADFAE1", "#72F4D2", "#29B5FF", "#85F4DF", "#4CDDF8", "#729EFF", "#E0CFFF", "#FF94EF", "#FF79CC", "#8FBEFF", "#64FFEE", "#D0B2FF", "#FFAA93", "#B2B8FF", "#33D4E0", "#DBE2FF", "#F193F8", "#FACE29", "#FF9A53", "#FDB6B2", "#DBAEF9", "#FBD0D6", "#FFA7B1", "#FBB27A", "#FECCB3", "#CFC2EC", "#F2D4EE", "#44C8ED", "#FF7CAD", "#AECFEB", "#FA9084", "#DF94AC", "#F2B8A9", "#BAC3FB", "#F3EFE9", "#DEE7EC", "#AFD3FC", "#F3D3EE", "#B8F28C", "#F27FCB", "#F7F6F7", "#4BB5FD", "#D3FCFC", "#FDAB74", "#FBDC94", "#75C6FF", "#ADDFF6", "#E8BDAF", "#ECE1D8", "#90B5EF", "#B2A9FF", "#18D6E4", "#CED8E7", "#E6C8F4", "#D3A8F4", "#DBD9D4", "#C4F3C9", "#77E9D6", "#4FDBC1", "#F19A6C", "#BED7F3", "#FDB2AD", "#FD8F8C", "#55B6E9", "#97BCFE", "#CF82E5", "#6ED1E7", "#FF9689", "#B7C6DA", "#66E9CE", "#E6FFC4", "#98EEFF", "#FFBEEE", "#58E2FF", "#B0FFB9", "#60F5BF", "#BBDDFF", "#EA98FF", "#AACCFF", "#D69EFD", "#FFCF9E", "#FFAB9F", "#FF839B", "#57D6FF", "#8CC2FF", "#CDBBFF", "#EB89DF", "#E6FFFE", "#F7FEFF", "#F9EBF9", "#6FDAFF", "#B5C9FF", "#AEFFE9", "#71E7FF", "#D3E2FF", "#C4DFFF", "#88FFFF", "#59DCFF", "#D293FF", "#EBE1FF", "#ADFAE1", "#72F4D2", "#29B5FF", "#85F4DF", "#4CDDF8", "#729EFF", "#E0CFFF", "#FF94EF", "#FF79CC", "#8FBEFF", "#64FFEE", "#D0B2FF", "#FFAA93", "#B2B8FF", "#33D4E0", "#DBE2FF", "#F193F8", "#FACE29", "#FF9A53"],
        };
    },
    watch: {
        colorUseCopyData: {
            deep: true,
            handler: function () {
                //判断是否点击了选项
                let AllIsActive = this.colorDataArray.every((item) => {
                    return item.isActive !== false;
                });
                //控制all的样式
                if (!AllIsActive) {
                    this.flAllCheckout = false;
                } else {
                    this.flAllCheckout = true;
                }
                // 处理选中节点的显示
                let nameArray = [];
                this.colorDataArray.forEach((item) => {
                    if (item.isActive) {
                        nameArray.push({
                            name: item.name,
                        });
                    }
                });

                let needShowNodeArray = [];
                this.data.nodes.forEach((item) => {
                    nameArray.forEach((value) => {
                        if (item["chain-id"] === value.name) {
                            needShowNodeArray.push(item);
                        }
                    });
                });
                this.copyData.nodes = needShowNodeArray;
                this.initChartsGraph();
            },
        },
    },
    methods: {
        handleScroll(e) {
            let scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
            if (scrollTop > 5) {
                this.flShowHoverUp = false;
            } else {
                this.flShowHoverUp = true;
            }
        },
        scrollBottom() {
            window.pageYOffset = 10000000;
            document.documentElement.scrollTop = 10000000;
            document.body.scrollTop = 10000000;
        },
        revertGraph() {
            window.location.reload();
        },
        sortByConnection(flSwitchValue) {
            this.flShowLetterSortIcon = false;
            this.flShowConnectionSortIcon = true;
            let copyData = JSON.parse(JSON.stringify(this.colorDataArray));
            this.flConnectionActiveStyle = true;
            this.flLetterActiveStyle = false;
            this.sortByConnectionSwitchIcon = !this.sortByConnectionSwitchIcon;
            if (flSwitchValue) {
                this.colorDataArray = copyData.sort((a, b) => {
                    return a.connection - b.connection;
                });
            } else {
                this.colorDataArray = copyData.sort((a, b) => {
                    return b.connection - a.connection;
                });
            }
        },
        sortByLetter(sortRule) {
            this.flShowLetterSortIcon = true;
            this.flShowConnectionSortIcon = false;
            let copyData = JSON.parse(JSON.stringify(this.colorDataArray));
            this.flConnectionActiveStyle = false;
            this.flLetterActiveStyle = true;
            this.sortByLetterSwitchIcon = !this.sortByLetterSwitchIcon;
            this.colorDataArray = copyData.sort((a, b) => {
                let value1 = a.name.toLowerCase(),
                    value2 = b.name.toLowerCase();
                if (sortRule) {
                    if (value1 > value2) {
                        return -1;
                    } else {
                        return 1;
                    }
                } else {
                    if (value1 < value2) {
                        return -1;
                    } else {
                        return 1;
                    }
                }
            });
        },
        onresize() {
            this.$refs.chart_content.style.height = window.innerHeight - 353 + "px";
            this.$refs.graph_list_content.style.height = window.innerHeight - 350 + "px";
            this.graphEcharts.resize();
        },
        flPickUp() {
            this.showDefaultAnimation = false;
            this.flShowDefaultPickUp = true;
            setTimeout(() => {
                this.flShowDefaultRightPickUp = true;
            }, 1000);
        },
        flShowPickUp() {
            this.showDefaultAnimation = true;
            this.flShowDefaultRightPickUp = false;
            setTimeout(() => {
                this.flShowDefaultPickUp = false;
            }, 1000);
        },
        /*switchBgColor(switchValue){
            // true 为星空背景
            if(switchValue){
                sessionStorage.setItem('starSky','show')
                this.$uMeng.push('TurnOn_starry-skymode','click')
            }else {
                this.$uMeng.push('TurnOff_starry-skymode','click')
                sessionStorage.setItem('starSky','hide')
            }
        },*/
        async getData() {
            this.$store.commit("flShowLoading", true);
            let chainsConnectionUrl = `${process.env.VUE_APP_BASE_URL}${urlApi.chainsConnection}`;
            try {
                const {code, data, message} = await Service.http(chainsConnectionUrl);
                if(code === Constant.API_CODE.SUCCESS) {
                    if(data) {
                        if(data.items?.length) {
                            const chainListMap = new Map();
                            const production = false;
                            let nodes =  [];
                            let paths = [];
                            let path;
                            let itemIconPromiseAll = [];
                            data.items.forEach((item, index) => {
                                chainListMap.set(item.chain_name, {
                                    chain_name: item.chain_name,
                                    pretty_name: item.pretty_name,
                                    current_chain_id: item.current_chain_id
                                })
                                itemIconPromiseAll[index] = new Promise(resolve => {
                                    const img = new Image();
                                    img.src = item.icon;
                                    img.onload = function () {
                                        resolve(item.icon);
                                    };
                                    img.onerror = function () {
                                        resolve(Constant.ERROR_LOAD_IMG);
                                    };
                                })
                            })
                            try {
                                const imageSrcList = await Promise.all(itemIconPromiseAll);
                                data.items.forEach((item, index) => {
                                    item.icon = imageSrcList[index];
                                    let node = {
                                        "chain-id": `${item.pretty_name} (${item.current_chain_id})`,
                                        "connections": item.connection_chains.length,
                                        "icon": item.icon,
                                        "chain-name": item.chain_name
                                    }
                                    nodes.push(node);
                                    item.connection_chains.forEach(chain => {
                                        const matchChainInfo = chainListMap.has(chain.chain_name) && chainListMap.get(chain.chain_name);
                                        if(matchChainInfo) {
                                            path = {
                                                "src-chain-id": `${item.pretty_name} (${item.current_chain_id})`,
                                                "dst-chain-id": `${matchChainInfo.pretty_name} (${matchChainInfo.current_chain_id})`,
                                                "state": Constant.CHAIN_CONNECT_STATUS[chain.connection_status]
                                            }
                                            paths.push(path);
                                        }
                                    })
                                })
                                let res = {
                                    "production": production,
                                    "nodes": [...nodes],
                                    "paths": [...paths]
                                }
                                this.data = res;
                                this.flShowTestTooltip = res.production;
                                //数据先排序
                                this.data.nodes.sort((a, b) => {
                                    return b.connections - a.connections;
                                });
                                this.maxLinks = this.data.nodes[0].connections;
                                this.data.nodes.forEach((item, index) => {
                                    item.isDelete = false;
                                    item.color = this.color[index];
                                });
                                this.copyData = JSON.parse(JSON.stringify(this.data));
                                this.colorUseCopyData = JSON.parse(JSON.stringify(this.data));
                                this.initLegend();
                                this.initChartsGraph();
                            } catch (error) {
                                this.$store.commit('flShowLoading', false)
                                console.log(error);
                            }
                        }
                    }
                } else {
                    console.log(message);
                    this.$store.commit('flShowLoading', false)
                    this.flShowNetwork = false;
                }
            } catch (error) {
                console.log(error);
                this.$store.commit('flShowLoading', false)
                if (error) {
                    this.flShowNetwork = false
                }
            }
        },
        refreshPage() {
            window.location.reload();
        },
        selectAll() {
            this.flAllCheckout = !this.flAllCheckout;
            let flIsActive = this.colorDataArray.every((item) => {
                return item.isActive !== false;
            });
            if (flIsActive) {
                this.colorDataArray.forEach((item) => {
                    item.isActive = false;
                });
                this.copyData.nodes = [];
                this.initChartsGraph();
            } else {
                this.colorDataArray.forEach((item) => {
                    item.isActive = true;
                    this.copyData = JSON.parse(JSON.stringify(this.data));
                });
                this.initChartsGraph();
            }
        },
        initLegend() {
            this.colorDataArray = [];
            this.copyData.nodes.forEach((item) => {
                this.colorDataArray.push({
                    icon: item.icon,
                    name: item["chain-id"],
                    color: item.color,
                    isActive: true,
                    connection: item.connections,
                });
            });
            this.sortCopyData = JSON.parse(JSON.stringify(this.colorDataArray));
        },
        changeChart(value, index) {
            this.colorDataArray[index].isActive = !this.colorDataArray[index].isActive;
            this.colorUseCopyData.nodes.forEach((item) => {
                if (item["chain-id"] === value.name) {
                    item.isDelete = !item.isDelete;
                }
            });
            this.copyData.nodes = this.colorUseCopyData.nodes.filter((item) => {
                return item.isDelete === false;
            });
        },
		setNodeImage (nodeName) {
			return nodeName ? `image://${nodeName}` : '';
        },
        formatName(nodeName) {
            let formatNameArr = nodeName.split(" ");
            const name = formatNameArr
                .filter((item, index) => {
                    if (index >= 1) {
                        return index !== formatNameArr.length - 1;
                    }
                    return item;
                })
                .join(" ");
            return name;
        },
        initChartsGraph() {
            this.flShowRevertIcon = false;
            let nodeLinksArray = [],
                nodeArray = [];
            //最大像素点与最小像素点的差值66  最小的symbolSize 为 8 * 节点递增的比例
            let symbolSizeRule = 50;
            //数据结果展示
            let minSymbolSizeRule = Math.floor(20 / (Number(symbolSizeRule) / this.data.nodes.length));
            for (let i in this.copyData.nodes) {
                let connectionValue = this.copyData.nodes[i].connections;
                let symbolSize = 50;
                if (this.copyData.nodes[i].connections === 1) {
					symbolSize = 50;
                } else if (this.copyData.nodes[i]["chain-name"] === Constant.CHAIN_SIGNER.COSMOSHUB || this.copyData.nodes[i]["chain-name"] === Constant.CHAIN_SIGNER.IRISNET) {
					symbolSize = (this.maxLinks + 0.5) * (symbolSizeRule / this.maxLinks) + 20;
                } else {
					symbolSize = connectionValue * (symbolSizeRule / this.maxLinks) + 20;
                }

                const name = this.formatName(this.copyData.nodes[i]["chain-id"]);
                nodeArray.push({
                    name: name,
                    nodeName: this.copyData.nodes[i]["chain-id"],
                    symbolSize: symbolSize,
                    label: {
                        show: true, //默认展示信息
                        position: "bottom",
                    },
                    symbol: this.setNodeImage(this.copyData.nodes[i].icon),
                    itemStyle: {
                        color: this.copyData.nodes[i].color,
                    },
                    tooltip: {
                        backgroundColor: "rgba(34,33,78,0.5)",
                        borderColor: this.copyData.nodes[i].color,
                        borderWidth: 1,
                        borderRadius: 15,
                        padding: [5, 20],
                        formatter: function (value) {
                            let name = value.data.nodeName.split(" ");
                            return name[name.length - 1].replace(/[(|)]/g, "");
                        },
                    },
                });
            }
            let categories = [];
            this.copyData.nodes.forEach((item) => {
                categories.push({
                    name: item["chain-id"],
                });
            });

            this.copyData.paths.forEach((item, index) => {
                const sourceName = this.formatName(item["src-chain-id"]);
                const targetName = this.formatName(item["dst-chain-id"]);
                if (item.state === Constant.CHAIN_CONNECT_STATUS[1]) {
                    nodeLinksArray.push({
                        source: sourceName,
                        target: targetName,
                        lineStyle: {
                            color: "rgba(112, 198, 199, 1)",
                            width: 1,
                        },
                        emphasis: {
                            lineStyle: {
                                width: 1,
                                type: "solid",
                            },
                        },
                    });
                } else {
                    nodeLinksArray.push({
                        source: sourceName,
                        target: targetName,
                        //连接线的样式设置
                        lineStyle: {
                            color: "rgba(112, 198, 199, 1)",
                            type: "dashed",
                        },
                        //鼠标滑到连接线上的样式配置
                        emphasis: {
                            lineStyle: {
                                width: 1,
                                type: "dashed",
                            },
                        },
                    });
                }
            });
            let graphOption = {
                name: "",
                tooltip: {
                    trigger: "item",
                    formatter: function (data) {
                        return `<span>${data.name.toString().replace(">", "<->")}</span>`;
                    },
                },
                series: [
                    {
                        type: "graph",
                        layout: "force",
                        animation: false,
                        // animationDuration: 10000,
                        // animationThreshold: 1000,
                        animationEasingUpdate: "exponentialOut", //缓动效果
                        // animationDurationUpdate:300,
                        // edgeSymbol: ['','arrow'], //指示线的箭头
                        // edgeSymbolSize:5,
                        data: nodeArray,
                        // links: nodeLinksArray,
                        nodeScaleRatio: 0.6, //鼠标每次缩放的整体缩放比例
                        force: {
                            repulsion: [800, 1000], //斥力因子
                            gravity: 0.8, //是否向中心靠拢 值越大越接近于中心
                            edgeLength: [100, 200], //链接线的长度范围
                            layoutAnimation: true,
                        },
                        // zoom:0.1, //设置整体视图缩放的比例
                        zoom: this.dataIndex,
                        // symbol:'',//图形形状
                        roam: true, //平移和缩放  move 和 scale true 包含所有
                        draggable: true, //是否拖拽
                        hoverAnimation: true,
                        focusNodeAdjacency: true,
                        scaleLimit: {
                            min: 0.4,
                            max: 2,
                        },
                        /*itemStyle: {
								shadowColor: 'rgba(255, 255, 255, 0.8)',
								shadowOffsetY:2,
								shadowBlur: 3,
								opacity:1
							},*/
                        emphasis: {
                            lineStyle: {
                                width: 1,
                            },
                            label: {
                                show: false,
                            },
                        },
                    },
                ],
            };
            this.graphEcharts.setOption(graphOption);
            clearInterval(this.dataTimer);
            // 根据节点数使用不同的缩放规则
            let zoomRule = 1,
                zoomSpeedRule = 0.05,
                setTime = 950;
            if (nodeArray.length > 150) {
                zoomRule = 0.3;
                zoomSpeedRule = 0.2;
                setTime = 1450;
            } else if (nodeArray.length > 100 && nodeArray.length < 150) {
                zoomRule = 0.6;
                zoomSpeedRule = 0.2;
                setTime = 1250;
            } else if (nodeArray.length > 50 && nodeArray.length < 100) {
                zoomRule = 0.7;
                zoomSpeedRule = 0.04;
                setTime = 950;
            } else if (nodeArray.length < 50) {
                zoomRule = 1;
                zoomSpeedRule = 0.07;
                setTime = 950;
            }
            this.$store.commit("flShowLoading", false);
            // 渲染完成收缩图形让其在区域内全部展示

            this.dataTimer = setInterval(() => {
                graphOption.series[0].draggable = false;
                this.dataIndex = this.dataIndex - 0.1;
                if (this.dataIndex < zoomRule - 0.5) {
                    this.dataIndex = zoomRule - 0.28;
                    clearInterval(this.dataTimer);
                    graphOption.series[0].draggable = true;
                }
                graphOption.series[0].zoom = this.dataIndex;
                this.graphEcharts.setOption(graphOption);
            }, 30);

            //最后一次渲染
            setTimeout(() => {
                graphOption.series[0].zoom = zoomRule - 0.28;
                graphOption.series[0].links = nodeLinksArray;
                graphOption.series[0].force.gravity = 0.3;
                this.graphEcharts.setOption(graphOption);
                this.flShowRevertIcon = true;
            }, setTime);
        },
        formatPrettyName(name) {
            const prettyName = this.formatName(name);
            return prettyName;
        },
        formatCurrentChainId(name) {
            const nameSplit = name.split(" ");
            const currentChainId = nameSplit[nameSplit.length - 1];
            return currentChainId;
        },
        prettyNamePopover(name) {
            return Tools.getTextWidth(this.formatPrettyName(name)) > 139;
        },
        currentChainIdPopover(name) {
            return Tools.getTextWidth(this.formatCurrentChainId(name)) > 139;
        },
        displayImgSrc(image) {
            return image;
        }
    },
    async mounted() {
        /*let container = new PerfectScrollbar('#graph_list_container',{
            wheelSpeed: 2,
            wheelPropagation: true,
            minScrollbarLength: 20
        });
        container.update()*/
        this.$refs.chart_content.style.height = window.innerHeight - 353 + "px";
        this.$refs.graph_list_content.style.height = window.innerHeight - 350 + "px";
        this.graphEcharts = echarts.init(document.getElementById("validator_graph_content"));
        window.addEventListener("resize", this.onresize);
        clearTimeout(this.timer);
        await this.getData();
        this.timer = setInterval(async () => {
            await this.getData();
        }, 300000);
        window.addEventListener("scroll", this.handleScroll, true);
    },
    beforeDestory() {
        window.removeEventListener("resize", this.onresize);
        window.removeEventListener("scroll", this.handleScroll, true);
    }
};
</script>

<style scoped lang="scss">
@import "../../style/mixin";
.default_bg_img {
    background: rgba(32, 35, 66, 1);
}

.start_sky_img {
    background-image: url("../../assets/GOZ_bg.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
}

.validator_graph_container {
    width: 100%;
    height: 100%;
    //线上数据
    .graph_content_wrap {
        width: 100%;
        box-sizing: border-box;
        padding: 0 0.4rem;
        margin: 0 auto;
        margin-bottom: 0.2rem;
        .graph_content_title {
            height: 0.6rem;
            font-size: 0.18rem;
            color: rgba(255, 255, 255, 1);
            display: flex;
            align-items: center;
            .beat_content {
                margin-left: 0.18rem;
                background: rgba(55, 124, 248, 1);
                color: #fff;
                font-size: 0.14rem;
                border-radius: 0.1rem;
                padding: 0.02rem 0.1rem;
            }

            .tooltip {
                display: flex;
                justify-content: space-between;

                .graph_tooltip {
                    display: flex;

                    p:nth-of-type(1) {
                        display: flex;
                        align-items: center;

                        span:nth-of-type(1) {
                            margin-left: 0.2rem;
                            display: inline-block;
                            width: 0.2rem;
                            height: 0.02rem;
                            border-top: 0.02rem solid #70c6c7;
                        }

                        span:nth-of-type(2) {
                            margin-left: 0.1rem;
                            font-size: 0.14rem;
                            color: rgba(115, 122, 174, 1);
                        }
                    }

                    p:nth-of-type(2) {
                        display: flex;
                        align-items: center;

                        span:nth-of-type(1) {
                            width: 0.2rem;
                            margin-left: 0.41rem;
                            display: inline-block;
                            border-top: 0.02rem dashed #70c6c7;
                        }

                        span:nth-of-type(2) {
                            margin-left: 0.1rem;
                            font-size: 0.14rem;
                            color: rgba(115, 122, 174, 1);
                        }
                    }
                }

                .background_toggle_content {
                    margin-left: 0.35rem;
                }
            }
        }

        .graph_charts_container {
            height: 100%;

            .graph_charts_title {
                padding: 0.2rem 0 0.2rem 0.2rem;
                font-size: 0.12rem;
                line-height: 0.14rem;
                color: rgba(115, 122, 174, 1);
            }

            .graph_content_container {
                height: 100%;
                display: flex;
                position: relative;

                .icon-fuwei {
                    position: absolute;
                    left: 0;
                    top: 0;
                    cursor: pointer;
                    color: #377cf8;
                    font-size: 0.24rem;
                    z-index: 9;
                }

                .graph_container_content_wrap {
                    flex: 1;
                    background: #2d325a;
                    border-radius: 0.04rem;

                    #validator_graph_content {
                        /*max-width: 8rem;*/
                        width: 100%;
                        background: transparent;
                        //线上
                        /*background: #2D325A;*/
                    }
                }

                .sort_content {
                    box-sizing: border-box;
                    padding-left: 0.38rem;
                    margin-bottom: 0.22rem;

                    .active_style {
                        color: #377cf8;
                    }

                    .un_active_style {
                    }

                    span {
                        cursor: pointer;
                        margin-right: 0.2rem;
                        font-size: 0.14rem;
                        color: #868fd3;
                        white-space: nowrap;
                    }
                }

                .sort_scroll_content::-webkit-scrollbar {
                    width: 0.05rem;
                    height: 0.05rem;
                }

                .sort_scroll_content::-webkit-scrollbar-thumb {
                    background: rgba(63, 67, 105, 1);
                    border-radius: 0.05rem;
                }

                /*.sort_scroll_content::-webkit-scrollbar-thumb:hover{
						background: #333;
					}
					.sort_scroll_content::-webkit-scrollbar-corner{
						background: #179a16;
					}*/
                
                .graph_list_container {
                    margin-left: 0.02rem;
                    padding-top: 0.2rem;
                    border-radius: 0.04rem;
                    position: absolute;
                    overflow: hidden;
                    right: 0;
                    top: 0;
                    background: rgba(95, 115, 255, 0.15);
                    backdrop-filter: blur(2px);
                    scrollbar-width: none;
                    .sort_scroll_content {
                        overflow-x: hidden;
                        overflow-y: auto;
                        margin-top: 0.02rem;
                        @media (max-width: 1040px) {
                            height: 2rem !important;
                        }
                        .graph_list_item_all {
                            /*position: fixed;*/
                            display: flex;
                            justify-content: flex-start;
                            margin: 0 0.6rem 0.2rem 0.38rem;
                            background: transparent;

                            .legend_all_block {
                                width: 0.14rem;
                                height: 0.14rem;
                                border-radius: 0.07rem;
                                cursor: pointer;
                                background: transparent;
                                img {
                                    width: 0.14rem;
                                }
                            }

                            .legend_block {
                                box-sizing: border-box;
                                width: 0.14rem;
                                height: 0.14rem;
                                border-radius: 0.07rem;
                                cursor: pointer;
                            }

                            .hide_style {
                                color: rgba(134, 143, 211, 1) !important;
                            }

                            .legend_name {
                                cursor: pointer;
                                height: 0.14rem;
                                color: rgba(134, 143, 211, 0.5);
                                margin-left: 0.1rem;
                                width: 0.9rem;
                                white-space: nowrap;
                            }

                            .legend_name_content {
                                .legend_name {
                                    cursor: pointer;
                                    height: 0.14rem;
                                    color: rgba(134, 143, 211, 0.5);
                                    margin-left: 0.1rem;
                                    width: 0.9rem;
                                    white-space: nowrap;
                                }
                            }
                        }

                        .graph_list_content {
                            flex: 1;
                            max-width: 2.2rem;
                            /*margin-top: 0.34rem;*/
                            min-height: 0.14rem;
                            display: flex;
                            flex-wrap: wrap;
                            margin-left: 0.18rem;
                            padding-bottom: 0.4rem;

                            .graph_list_item {
                                width: 2rem;
                                align-items: center;
                                margin-left: 0.1rem;
                                display: flex;
                                padding: 0.05rem 0 0.1rem 0;
                                cursor: pointer;

                                .legend_all_block {
                                    width: 0.14rem;
                                    height: 0.14rem;
                                    border-radius: 0.07rem;
                                    cursor: pointer;

                                    img {
                                        width: 0.14rem;
                                    }
                                }

                                .legend_name_content {
                                    @include flex;
                                    flex-direction: column;
                                    align-items: flex-start;
                                    margin-left: 0.08rem;
                                    font-weight: 600;
                                    color: rgba(134, 143, 211, 1);
                                    .legend_item {
                                        margin-top: 0.03rem;
                                        max-width: 1.6rem;
                                        overflow: hidden;
                                        text-overflow: ellipsis;
                                        white-space: nowrap;
                                        &:first-child {
                                            margin-top: 0;
                                        }
                                    }
                                    .hide_style_color {
                                        color: rgba(134, 143, 211, 0.5) !important;
                                    }
                                }

                                .hide_style {
                                    background: transparent !important;
                                    border: 0.01rem solid rgba(134, 143, 211, 0.5);
                                }

                                .legend_block {
                                    box-sizing: border-box;
                                    width: 0.14rem;
                                    height: 0.14rem;
                                    border-radius: 0.07rem;
                                }
                                .legend_block_img_loading {
                                    width: 0.24rem;
                                    height: 0.24rem;
                                    opacity: 1;
                                }
                                .legend_block_img {
                                    box-sizing: border-box;
                                    width: 0.4rem;
                                    height: 0.4rem;
                                    border-radius: 50%;
                                    cursor: pointer;
                                    opacity: 1;
                                }
                                .img_hide_style {
                                    opacity: 0.4;
                                }
                                :deep(.el-image__inner) {
                                    box-sizing: border-box;
                                    width: 0.4rem;
                                    height: 0.4rem;
                                    border-radius: 50%;
                                    cursor: pointer;
                                    opacity: 1;
                                }
                            }
                        }
                    }
                }

                .pick_up_content {
                    position: absolute;
                    right: 2.29rem;
                    top: 50%;
                    z-index: 10;

                    .pick_up_img_content {
                        width: 0.14rem;
                        height: 0.6rem;
                        cursor: pointer;

                        img {
                            width: 100%;
                        }
                    }
                }

                .pick_up_show_content {
                    position: absolute;
                    top: 50%;
                    right: 0;

                    .pick_up_show_img_content {
                        width: 0.14rem;
                        height: 0.6rem;
                        cursor: pointer;

                        img {
                            width: 100%;
                        }
                    }
                }

                .hide_pick_up_style {
                    animation: hideGraphOptionList 1s ease-in;
                    animation-fill-mode: forwards;
                }

                .show_pick_up_style {
                    animation: showGraphOptionList 1s ease-in;
                    animation-fill-mode: forwards;
                }
            }
        }

        .hover_up_content {
            position: absolute;
            bottom: 0;
            left: 50%;
            cursor: pointer;
            display: flex;
            justify-content: center;
            margin-top: 0.05rem;
        }

        .show_error_content {
            min-height: 6rem;
            display: flex;
            align-items: center;
            justify-content: center;

            .error_content {
                text-align: center;

                p {
                    max-width: 2.05rem;
                    color: rgba(115, 122, 174, 1);
                    font-size: 0.14rem;
                    margin-bottom: 0.2rem;
                }

                span {
                    cursor: pointer;
                    text-align: center;
                    color: rgba(55, 124, 248, 1);
                    font-size: 0.14rem;
                    line-height: 0.16rem;
                    border-bottom: 0.01rem solid rgba(55, 124, 248, 1);
                }
            }
        }
    }
}

@media (max-width: 1040px) {
    .validator_graph_container {
        .graph_content_wrap {
            box-sizing: border-box;
            padding: 0 0.3rem;

            .graph_content_title {
                height: auto;
                padding: 0.15rem 0;
                align-items: flex-start;
            }

            .tooltip {
                justify-content: space-between;

                .graph_tooltip {
                    flex: 1;
                    display: flex;
                    align-items: flex-start;
                }

                .background_toggle_content {
                    text-align: right;
                }
            }

            .graph_charts_container {
                .graph_content_container {
                    flex-direction: column;

                    .sort_content {
                        padding-left: 0.2rem;
                    }

                    .graph_container_content_wrap {
                        border-radius: 0;
                    }

                    .graph_list_container {
                        max-width: none;
                        height: 2rem !important;
                        margin-bottom: 0.2rem;
                        border-radius: 0;
                        margin-left: 0;
                        position: static;
                        .sort_scroll_content {
                            .graph_list_item_all {
                                width: 1.2rem;
                                margin-left: 0.2rem;
                                justify-content: flex-start;
                            }

                            .graph_list_content {
                                height: auto;
                                max-width: none;

                                .graph_list_item {
                                    margin-left: 0;
                                }
                            }
                        }
                    }

                    .pick_up_content {
                        display: none;
                    }
                }
            }
        }
    }
}

@keyframes rotate {
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);
    }
}

@media (max-width: 900px) {
    .validator_graph_container {
        .graph_content_wrap {
            .graph_content_title {
                flex-direction: column;

                .beat_content {
                    ::v-deep.tooltip_content {
                        width: 3rem;
                    }
                }

                .tooltip {
                    margin-top: 0.1rem;
                    flex-direction: row;

                    .graph_tooltip {
                        flex-direction: row;

                        p:nth-of-type(1) {
                            span:nth-of-type(1) {
                                margin-left: 0;
                            }
                        }

                        p:nth-of-type(2) {
                            span:nth-of-type(1) {
                                margin-left: 0;
                            }
                        }
                    }

                    .background_toggle_content {
                        margin-right: 0.1rem;
                    }
                }
            }
        }
    }
}

@media (max-width: 610px) {
    .validator_graph_container {
        .graph_content_wrap {
            .graph_content_title {
                .tooltip {
                    flex-direction: column;

                    .graph_tooltip {
                        flex-direction: column;

                        p:nth-of-type(1) {
                            span:nth-of-type(1) {
                                margin-left: 0;
                            }
                        }

                        p:nth-of-type(2) {
                            span:nth-of-type(1) {
                                margin-left: 0;
                            }
                        }
                    }

                    .background_toggle_content {
                        margin-top: 0.1rem;
                        margin-left: 0;
                    }
                }
            }

            .graph_charts_container {
                .graph_charts_title {
                    padding-left: 0;
                    padding-top: 0;
                }

                .graph_content_container {
                    .iconfuwei {
                        padding-left: 0;
                    }

                    .sort_content {
                        padding-left: 0.2rem;
                    }

                    .graph_list_container {
                        height: 2rem !important;
                        margin-bottom: 0.2rem;
                    }
                }
            }
        }
    }
}

@keyframes hideGraphOptionList {
    from {
        width: 2.43rem;
    }
    to {
        width: 0;
    }
}

@keyframes showGraphOptionList {
    from {
        width: 0;
    }
    to {
        width: 2.43rem;
    }
}
</style>
