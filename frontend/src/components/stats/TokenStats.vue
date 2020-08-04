<template>
    <div class="blocks_list_page_wrap">
        <page-title :title="pageTitle"
                    :flShowPageLink="false"></page-title>
        <div
            class="list_page_container"
            :class="[$store.state.isMobile ? 'mobile_list_page_container' : 'pc_list_page_container']"
        >
            <div>
                <div class="page_title">IRIS Token Stats</div>
                <div class="table_container" v-show="!itemsNoData">
                    <div class="information_props_wrap" v-for="v in items" :key="v.label">
                        <span class="information_props">{{v.label}}</span>
                        <span class="information_value"
                              :class="v.value ? 'skip_route' : ''"
                              v-if="v.label === 'Burned'">
                            <router-link @click.native="$uMeng.push('IRIS Stats_Burned','click')"
                                         :to="burnedCoins">{{v.value || '--'}}</router-link>
                            </span>
                        <span class="information_value"
                              :class="v.value ? 'skip_route' : ''"
                              v-if="v.label === 'Community Tax'">
                            <router-link @click.native="$uMeng.push('IRIS Stats_Community Tax','click')"
                                         v-if="v.value && v.value !== '--'"
                                         :to="communityTaxCoins">{{v.value || '--'}}</router-link>
                            </span>
                        <span v-if="v.value && v.value === '--'">--</span>
                        <span class="information_value" v-if="v.label !== 'Burned' && v.label !== 'Community Tax'">{{v.value || '--'}}</span>
                    </div>
                </div>
                <div v-show="itemsNoData" class="no_data_show"><img src="../../assets/no_data.svg" alt=""></div>
                <div class="page_title">IRIS Token Distribution</div>
                <div class="echarts_container" v-show="!pieDatasNoData">
                    <m-token-stats-echart :data="pieDatas"></m-token-stats-echart>
                </div>
                <div v-show="pieDatasNoData" class="no_data_show">
                    <img src="../../assets/no_data.svg" alt="">
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import MTokenStatsEchart from "../commontables/MTokenStatsEchart";
import Service from "../../service";
import Tools from "../../util/Tools";
import PageTitle from "../pageTitle/PageTitle";
import pageTitleConfig from "../pageTitle/pageTitleConfig";
import bigNumber from "bignumber.js"
export default {
    components: {
        PageTitle,
        MTokenStatsEchart
    },
    data() {
        return {
            pageTitle:pageTitleConfig.StatsIRISStats,
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
                    label: "Community Tax",
                    value: ""
                },
                {
                    label: "Burned",
                    value: ""
                },
                {
                    label: "Bonded",
                    value: ""
                }
            ],
            burnedCoins:'',
            communityTaxCoins:'',
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
                                        value: data.totalsupply_tokens ? Tools.formatAmount2(data.totalsupply_tokens,4,) : '--'
                                    },
                                    {
                                        label: "Circulation",
                                        value: data.circulation_tokens ? Tools.formatAmount2(data.circulation_tokens,4,) : '--'
                                    },
                                    {
                                        label: "Community Tax",
                                        value: data.community_tax ? Tools.formatAmount2(data.community_tax, 4) : '--'
                                    },
                                    {
                                        label: "Burned",
                                        value: data.burned_tokens ? Tools.formatAmount2(data.burned_tokens, 4) : '--'
                                    },
                                    {
                                        label: "Bonded",
                                        value: data.delegated_tokens ? Tools.formatAmount2(data.delegated_tokens,4) : '--'
                                    }
                                ];
                                console.error(data)
                                console.error(obj)
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
                                    v[1].totalAmount = `${new bigNumber(v[1].totalAmount.split(' ')[0]).toFormat()} ${v[1].totalAmount.split(' ')[1]}`
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
			    let s = num + "",n;
			    let arr = s.split(".");
			    arr[1] = arr[1] || "";
			    if(arr[1].toString().length > 4){
				    n =`${arr[0]}.${arr[1].substring(0, 4)}`
			    }else {
				    let diffNum = 4 - arr[1].toString().length;
				    for(let i = 0; i < diffNum; i++){
					    arr[1] += '0'
				    }
				    n = `${arr[0]}.${arr[1]}`
			    }
			    // let n = `${arr[0]}.${arr[1].padEnd(4, "0").substring(0, 4)}`;
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
        let envInformation = JSON.parse(sessionStorage.getItem('skinEnvInformation'));
        if( envInformation && envInformation.cur_env === 'mainnet' && envInformation.chain_id === "irishub"){
            this.burnedCoins = `/address/iaa108a0ts008fphurftmsvj5p2q8ltq8qedy0jxd8`;
            this.communityTaxCoins = `/address/iaa18rtw90hxz4jsgydcusakz6q245jh59kfma3e5h`
        }else {
            this.burnedCoins = `/address/faa108a0ts008fphurftmsvj5p2q8ltq8qeduq57d6`;
            this.communityTaxCoins = `/address/faa18rtw90hxz4jsgydcusakz6q245jh59kfrjhp52`
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style/mixin";

.blocks_list_page_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    .list_page_container {
        width: 100%;
        padding-top: 0.54rem;
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
                color: var(--titleColor);
            }
            .table_container {
                display: flex;
                flex-wrap: wrap;
                width: 100%;
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
                    background: #fff;
                    .information_props {
                        color: var(--contentColor);
                        font-size: 0.14rem;
                    }
                    .information_value {
                        color: var(--titleColor);
                        font-size: 0.16rem;
                        margin-top: 0.12rem;
                        word-break: break-all;
                        word-wrap: break-word;
                        > a{
                            font-size: 0.16rem;
                        }
                    }
                    .skip_route {
                        a,
                        span {
                            cursor: pointer;
                            color: var(--bgColor) !important;
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
                background: #fff;
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
            box-sizing: border-box;
            padding: 0 0.1rem;
            .page_title {
                padding-left: 0.1rem;
            }
        }
        .table_container {
            margin-right: 0 !important;
            width: 100%;
            .information_props_wrap {
                width: 100% !important;
                flex: 0 0 calc(100% - 0.22rem) !important;
                padding: 0.1rem !important;
                margin-bottom: 0.1rem !important;
                .information_props {
                    width: 100% !important;
                    box-sizing: border-box;
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
    @media screen and (max-width: 910px){
       .blocks_list_page_wrap{
           .list_page_container{
               padding-top: 0;
           }
       }
    }
</style>
