<template>
    <div class="parameters_page_container">
        <div class="parameters_list_container">
            <page-title :title="pageTitle"
                        :content="contentDoc"
                        :number="totalNumber"
                        :reversal="false"></page-title>
            <div class="parameter_list_content">
                <div class="parameters_list_cards_content">
                    <div style="overflow-y: hidden;" v-show="!showNoData">
                        <div
                            :class="[$store.state.isMobile ? 'mobile_cards_layout' : 'pc_cards_layout']"
                            v-for="(value, index) in parametersList"
                            :key="index"
                        >
                            <div class="card_title" v-show="value && value.length > 0">
                                <span>{{index}}</span>
                                <div class="parameter_note">
                                    <div class="current">
                                        <i></i>
                                        <span>Current Value</span>
                                    </div>
                                    <div class="genesis">
                                        <i></i>
                                        <span>{{index === 'Asset' ? 'Initial Value' : 'Genesis Value'}}</span>
                                    </div>
                                </div>
                            </div>
                            <p class="cards" v-for="(v, i) in value" :key="i">
                                <m-parameters-card :class="index" :data="v"></m-parameters-card>
                            </p>
                        </div>
                    </div>
                </div>
                <div v-show="showNoData" class="no_data_show">
                    <img src="../../assets/no_data.svg" alt />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import VParameters from "./ParametersTable";
import Service from "../../service";
import Tools from "../../util/Tools";
import MParametersCard from "../commontables/MParametersCard";
import BigNumber from "bignumber.js";
import PageTitle from "../pageTitle/PageTitle";
import pageTitleConfig from "../pageTitle/pageTitleConfig";
export default {
    name: "Parameters",
    components: {PageTitle, VParameters, MParametersCard },
    data() {
        return {
            parametersList: {},
            proposalsListPageWrap: "",
            showNoData: false,
            totalNumber:0,
            pageTitle:pageTitleConfig.GovParameters,
            contentDoc:'Total'
        };
    },
    mounted() {
        this.getParametersList();
    },
    methods: {
        getParametersList() {
            this.showNoData = false;
            this.parametersList = {};
            Service.commonInterface({ govParams: {} }, res => {
                try {
                    if (Array.isArray(res)) {
                        this.totalNumber = res.length;
                        let arr = res.map(item => {
                            item.genesis_value =
                                item.genesis_value || item.initial_value;
                            this.handleParameterItem(item);
                            return item;
                        });
                        let o = {
                            Staking: null,
                            Slashing: null,
                            General: null,
                            Asset: null
                        };
                        o.Staking =
                            arr.filter(
                                v =>
                                    v.module === "mint" ||
                                    v.module === "stake" ||
                                    v.module === "distr"
                            ) || null;
                        o.Slashing =
                            arr.filter(v => v.module === "slashing") || null;
                        o.Asset = arr.filter(v => v.module === "asset") || null;
                        o.General =
                            arr.filter(
                                v =>
                                    v.module !== "stake" &&
                                    v.module !== "slashing" &&
                                    v.module !== "distr" &&
                                    v.module !== "mint" &&
                                    v.module !== "asset"
                            ) || null;
                        o.Staking = [
                            ...o.Staking.filter(v => v.module === "mint"),
                            ...o.Staking.filter(v => v.module === "stake"),
                            ...o.Staking.filter(v => v.module === "distr")
                        ];
                        this.parametersList = o;
                    } else {
                        this.parametersList = {};
                        this.showNoData = true;
                    }
                } catch (e) {
                    this.parametersList = {};
                    this.showNoData = true;
                    console.error(e);
                }
            });
        },
        formatPer(value, min, max) {
            min = Number(min) || 0;
            max = Number(max) || 0;
            value = Number(value);
            if (value >= min && max > min && max >= value) {
                let per = ((Number(value) - min) / (max - min)) * 100;
                per = Math.max(per, 0);
                per = Math.min(per, 100);
                return per;
            } else if (value <= min) {
                return 0;
            } else if (value > max) {
                return 100;
            } else {
                return 0;
            }
        },
        handleParameterItem(parameterItem) {
            let arr = parameterItem.range.split(",");
            if (arr.length !== 2) {
                return;
            }
            parameterItem.min = arr[0];
            parameterItem.max = arr[1];
            parameterItem.current_per = this.formatPer(
                parameterItem.current_value,
                arr[0],
                arr[1]
            );
            parameterItem.genesis_per = this.formatPer(
                parameterItem.genesis_value,
                arr[0],
                arr[1]
            );
            parameterItem.current = parameterItem.current_value;
            parameterItem.genesis = parameterItem.genesis_value;
            //百分比
            let perDataArr = [
                "inflation",
                "base_proposer_reward",
                "bonus_proposer_reward",
                "min_signed_per_window",
                "community_tax",
                "slash_fraction_double_sign",
                "slash_fraction_downtime",
                "slash_fraction_censorship",
                "asset_tax_rate",
                "mint_token_fee_ratio",
                "gateway_asset_fee_ratio"
            ];
            //days
            let dayDataArr = [
                "double_sign_jail_duration",
                "unbonding_time",
                "downtime_jail_duration",
                "censorship_jail_duration"
            ];
            //amount
            let amountArr = ["issue_token_base_fee", "create_gateway_base_fee"];
            let blockArr = ["max_evidence_age", "signed_blocks_window"];
            if (parameterItem.key === "max_validators") {
            } else if (amountArr.indexOf(parameterItem.key) > -1) {
                let amountStr = this.$options.filters.amountFromat(
                    parameterItem.current_value
                );
                let current_value =
                    parameterItem.current_value.amount ||
                    parameterItem.current_value;
                let max =
                    parameterItem.max ||
                    Math.max(
                        Number(current_value),
                        Number(parameterItem.genesis_value)
                    );
                parameterItem.current_per = this.formatPer(
                    current_value,
                    arr[0],
                    max
                );
                parameterItem.genesis_per = this.formatPer(
                    parameterItem.genesis_value,
                    arr[0],
                    max
                );
                parameterItem.current = amountStr;
                parameterItem.genesis =
                    Tools.formatPrice(parameterItem.genesis_value) + " IRIS";
                parameterItem.max = parameterItem.max || "+∞";
            } else if (dayDataArr.indexOf(parameterItem.key) > -1) {
                let max =
                    parameterItem.max ||
                    Math.max(
                        Number(parameterItem.current_value),
                        Number(parameterItem.genesis_value)
                    );
                parameterItem.current_per = this.formatPer(
                    parameterItem.current_value,
                    arr[0],
                    max
                );
                parameterItem.genesis_per = this.formatPer(
                    parameterItem.genesis_value,
                    arr[0],
                    max
                );
                parameterItem.min =
                    this.formatUnbondingTime(parameterItem.min) || "0";
                parameterItem.max =
                    this.formatUnbondingTime(parameterItem.max) || "+∞";
                parameterItem.current = this.formatUnbondingTime(
                    parameterItem.current
                );
                parameterItem.genesis = this.formatUnbondingTime(
                    parameterItem.genesis
                );
            } else if (perDataArr.indexOf(parameterItem.key) > -1) {
                parameterItem.min =
                    Number(arr[0]) === 0
                        ? arr[0]
                        : `${new BigNumber(arr[0]).multipliedBy(100)}`;
                parameterItem.max =
                    Number(arr[1]) === 0
                        ? arr[1]
                        : `${new BigNumber(arr[1]).multipliedBy(100)} %`;
                parameterItem.current = `${
                    parameterItem.current_value
                        ? new BigNumber(
                              parameterItem.current_value
                          ).multipliedBy(100)
                        : parameterItem.current_value * 100
                } %`;
                parameterItem.genesis = `${
                    parameterItem.genesis_value
                        ? new BigNumber(
                              parameterItem.genesis_value
                          ).multipliedBy(100)
                        : parameterItem.genesis_value * 100
                } %`;
            } else if (parameterItem.key === "gas_price_threshold") {
                let maxL = String(arr[1]).length - 1;
                parameterItem.max = `${Number(arr[1]) / 10 ** maxL}*10^${maxL -
                    9} Nano`;
                let current_value_arr = String(
                    Number(parameterItem.current_value) / 10 ** 9
                ).split(".");
                let genesis_value_arr = String(
                    Number(parameterItem.genesis_value) / 10 ** 9
                ).split(".");
                let current_value_arr_0 = Tools.formatPrice(
                    current_value_arr[0]
                ).split(".")[0];
                let genesis_value_arr_0 = Tools.formatPrice(
                    genesis_value_arr[0]
                ).split(".")[0];
                parameterItem.current = current_value_arr[1]
                    ? `${current_value_arr_0}.${current_value_arr[1]} Nano`
                    : `${current_value_arr_0} Nano`;
                parameterItem.genesis = genesis_value_arr[1]
                    ? `${genesis_value_arr_0}.${genesis_value_arr[1]} Nano`
                    : `${genesis_value_arr_0} Nano`;
            } else if (parameterItem.key === "tx_size") {
                parameterItem.max =
                    Number(arr[1]) === 0
                        ? arr[1]
                        : arr[1] > 1
                        ? `${arr[1]} Bytes`
                        : `${arr[1]} Byte`;
                parameterItem.current =
                    parameterItem.current_value > 1
                        ? `${parameterItem.current_value} Bytes`
                        : `${parameterItem.current_value} Byte`;
                parameterItem.genesis =
                    parameterItem.genesis_value > 1
                        ? `${parameterItem.genesis_value} Bytes`
                        : `${parameterItem.genesis_value} Byte`;
            } else if (blockArr.indexOf(parameterItem.key) > -1) {
                parameterItem.min = parameterItem.max
                    ? parameterItem.min
                    : parameterItem.min > 1
                    ? `${parameterItem.min} Blocks`
                    : `${parameterItem.min} Block`;
                parameterItem.max = parameterItem.max
                    ? parameterItem.max > 1
                        ? `${parameterItem.max} Blocks`
                        : `${parameterItem.max} Block`
                    : "+∞";
                parameterItem.current =
                    parameterItem.current > 1
                        ? `${parameterItem.current} Blocks`
                        : `${parameterItem.current} Block`;
                parameterItem.genesis =
                    parameterItem.genesis > 1
                        ? `${parameterItem.genesis} Blocks`
                        : `${parameterItem.genesis} Block`;
            }
        },
        formatUnbondingTime(time) {
            let nsToMSRatio = 1000000,
                dToHRatio = 24,
                HToMRatio = 60;
            let dateTime = Tools.formatDuring(Number(time) / nsToMSRatio),
                d,
                h,
                m;
            if (dateTime.days >= 1) {
                d =
                    Math.floor(dateTime.days) > 1
                        ? `${Math.floor(dateTime.days)} Days`
                        : `${Math.floor(dateTime.days)} Day`;
            } else {
                d = "";
            }
            if (dateTime.hours >= 1 && dateTime.hours < dToHRatio) {
                h =
                    Math.floor(dateTime.hours) > 1
                        ? `${Math.floor(dateTime.hours)} Hours`
                        : `${Math.floor(dateTime.hours)} Hour`;
            } else {
                h = "";
            }
            if (dateTime.minutes >= 1 && dateTime.minutes < HToMRatio) {
                m =
                    Math.floor(dateTime.minutes) > 1
                        ? `${Math.floor(dateTime.minutes)} Minutes`
                        : `${Math.floor(dateTime.minutes)} Minute`;
            } else {
                m = "";
            }
            return `${d ? d : ""}${h ? h : ""}${m ? m : ""}`;
        }
    }
};
</script>

<style scoped lang="scss">
@import "../../style/mixin";
@media screen and (max-width: 910px){
    .parameters_page_container{
        .parameters_list_container{
            .parameter_list_content{
                padding-top: 0.05rem !important;
            }
        }
    }
}
.parameters_page_container {
    .parameters_list_container {
        width: 100%;
        padding-bottom: 0.4rem;
        .parameter_list_content {
            max-width: 12.8rem;
            width: 100%;
            margin: 0 auto;
            padding-top: 0.54rem;
            .parameter_list_title_content {
                font-size: 0.18rem;
                height: 0.7rem;
                display: inline-block;
                align-items: center;
                padding-left: 0.2rem;
                line-height: 0.7rem;
                font-weight: 400;
                margin-right: 0.2rem;
            }
            .parameters_list_table_wrap {
                overflow-x: auto;
                .parameters_list_table_content {
                    min-width: 11.5rem;
                }
            }
            .parameters_list_cards_content {
                overflow-x: auto;
                .pc_cards_layout {
                    display: flex;
                    flex-wrap: wrap;
                    min-width: 1280px;
                    & > p {
                        margin-bottom: 0.2rem !important;
                        margin-right: 0.2rem;
                        &:nth-of-type(3n) {
                            margin-right: 0;
                        }
                    }
                    &:nth-last-of-type(1) {
                        margin-bottom: -0.2rem;
                    }
                    &:nth-of-type(2) {
                        margin-top: -0.2rem;
                    }
                    &:nth-of-type(3) {
                        margin-top: -0.2rem;
                    }
                    &:nth-of-type(4) {
                        margin-top: -0.2rem;
                    }
                }
                div.card_title {
                    width: 100%;
                    height: 0.7rem;
                    margin-left: 0.2rem;
                    display: flex;
                    align-items: center;
                    font-size: 18px;
                    color: #515a6e;
                    font-weight: bold;
                }
                .cards{
                    background: #fff;
                }
                @mixin mobile_cards_layout {
                    display: flex;
                    justify-content: space-between;
                    flex-wrap: wrap;
                    background: #F5F7FD;
                    margin: 0 0.1rem;
                    & > p {
                        width: 100%;
                        margin-bottom: 0.1rem !important;
                        margin-right: 0 !important;
                        & > div.m_parameters_container {
                            width: 100%;
                        }
                    }
                }
                .mobile_cards_layout {
                    @include mobile_cards_layout;
                    &:nth-last-child(1) {
                        .cards {
                            &:nth-last-child(1) {
                                padding-bottom: 0;
                            }
                        }
                    }
                    &:nth-of-type(2) {
                        margin-top: -0.1rem;
                    }
                    &:nth-of-type(3) {
                        margin-top: -0.1rem;
                    }
                    &:nth-of-type(4) {
                        margin-top: -0.1rem;
                    }
                }
                //使用rem设置max-width不生效
                @media screen and (max-width: 910px) {
                    .pc_cards_layout {
                        @include mobile_cards_layout;
                    }
                }
            }
            .parameter_note {
                display: inline-block;
                font-size: 0.12rem;
                margin-left: 0.2rem;
                & > div {
                    display: inline-block;
                    i {
                        width: 0.11rem;
                        height: 0.11rem;
                        display: inline-block;
                        background-color: #ffb779;
                    }
                    span {
                        margin-left: 0.06rem;
                        color: var(--contentColor);
                    }
                }
                div.genesis {
                    margin-left: 0.3rem;
                    i {
                        background-color: #3698db;
                    }
                }
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
</style>
