<template>
    <div class="blocks_list_page_wrap">
        <div
            class="list_page_container"
            :class="[$store.state.isMobile ? 'mobile_list_page_container' : 'pc_list_page_container']"
        >
            <div>
                <div class="page_title">
                    <span>IRIS Bonded Tokens</span>
                    <div
                        class="select_container"
                        :class="validatorTypeSeleted ? 'active_select_container' : ''"
                        @click.stop="validatorTypeSeleted = !validatorTypeSeleted"
                    >
                        <span>{{validatorTypeLabel}}</span>
                        <ul v-show="validatorTypeSeleted">
                            <li
                                v-for="v in validatorTypeList"
                                :key="v.label"
                                @click="validatorTypeChange(v)"
                            >{{v.label}}</li>
                        </ul>
                    </div>
                </div>
                <div v-show="!showNoData" class="tree_map_container">
                    <m-tree
                        :items.sync="validatorList"
                        :validatorType="validatorTypeLabel"
                        :click-func="treeMapFunc"
                    ></m-tree>
                </div>
                <div v-show="showNoData" class="no_data_show">No Data</div>
            </div>
        </div>
        <spin-component :showLoading="showLoading" />
    </div>
</template>

<script>
import Tools from "../util/Tools";
import Constant from "../constant/Constant";
import Service from "../service";
import Http from "../util/axios";
import SpinComponent from "./commonComponents/SpinComponent";
import MTree from "./commonComponents/MTree";

export default {
    components: {
        SpinComponent,
        MTree
    },
    data() {
        return {
            items: [],
            validatorList: [],
            showNoData: false, //是否显示列表的无数据
            showLoading: false,
            validatorTypeSeleted: false,
            validatorType: "validator",
            validatorTypeLabel: "Active",
            validatorTypeList: [
                {
                    label: "All",
                    value: ""
                },
                {
                    label: "Active",
                    value: "validator"
                },
                {
                    label: "Candidate",
                    value: "candidate"
                },
                {
                    label: "Jailed",
                    value: "jailed"
                }
            ]
        };
    },
    mounted() {
        this.getData();
        this.bindEvents();
    },
    destroyed() {
        this.unbindEvents();
    },
    methods: {
        getData() {
            (async () => {
                try {
                    this.showLoading = true;
                    await this.getValidatorList();
                    // await this.getValidatorHeaderImg(this.items);
                    this.validatorList = this.items.sort((a, b) => {
                        return Number(b.votingPower) - Number(a.votingPower);
                    });
                    this.showLoading = false;
                } catch (err) {
                    this.showLoading = false;
                    console.log(err);
                }
            })();
        },
        validatorTypeChange(v) {
            this.validatorTypeLabel = v.label;
            this.validatorType = v.value;
            this.getData();
        },
        treeMapFunc(data) {
            let operatorAddress =
                data &&
                data.data &&
                data.data.info &&
                data.data.info.operatorAddress;
            if (operatorAddress) {
                this.getDelegations(operatorAddress);
            }
        },
        getDelegations(operatorAddress) {
            this.showLoading = true;
            Service.commonInterface(
                {
                    bondedtokensDelegations: {
                        address: operatorAddress
                    }
                },
                data => {
                    try {
                        let delegations = [];
                        if (Array.isArray(data.items)) {
                            for (let it of data.items) {
                                it.selfSharesFormat = Tools.formatPrice(
                                    Number(it.self_shares) + ""
                                );
                            }
                            delegations = data.items;
                        } else {
                            delegations = null;
                        }
                        let d = this.validatorList.find(
                            v => v.operatorAddress === operatorAddress
                        );
                        if (d && delegations) {
                            delegations.sort((a, b) => {
                                return (
                                    Number(a.self_shares) -
                                    Number(b.self_shares)
                                );
                            });
                            d.delegations = delegations;
                            d.delegationsActive = true;
                        }
                        this.showLoading = false;
                    } catch (e) {
                        this.showLoading = false;
                    }
                }
            );
        },
        getValidatorList() {
            return new Promise((resolve, reject) => {
                Service.commonInterface(
                    {
                        bondedtokensValidators: {
                            validatorStatus: this.validatorType
                        }
                    },
                    result => {
                        try {
                            result = result && result.data;
                            if (Array.isArray(result)) {
                                let allVotingPower = result.reduce(
                                    (init, v) => {
                                        return +v.voting_power + init;
                                    },
                                    0
                                );
                                this.items = result.map(item => {
                                    return {
                                        delegations: null,
                                        delegationsActive: false,
                                        moniker: item.moniker,
                                        operatorAddress: item.operator_address,
                                        ownerAddress: item.owner_address,
                                        votingPower: item.voting_power,
                                        identity: item.identity,
                                        imageUrl: item.icons,
                                        allVotingPower: allVotingPower
                                    };
                                });
                                this.showNoData = false;
                            } else {
                                this.showNoData = true;
                                this.items = [];
                            }
                            resolve();
                        } catch (e) {
                            this.showNoData = true;
                            reject();
                        }
                    }
                );
            });
        },
        getValidatorHeaderImg(data) {
            let url =
                "https://keybase.io/_/api/1.0/user/lookup.json?fields=pictures&key_suffix=";
            let requestArr = [];
            for (let i = 0; i < data.length; i++) {
                if (data[i].identity) {
                    requestArr.push(
                        new Promise((resolve, reject) => {
                            Http.http(`${url}${data[i].identity}`)
                                .then(res => {
                                    if (
                                        res &&
                                        res.them &&
                                        res.them[0].pictures &&
                                        res.them[0].pictures.primary &&
                                        res.them[0].pictures.primary.url
                                    ) {
                                        data[i].imageUrl =
                                            res.them[0].pictures.primary.url;
                                    }
                                    resolve();
                                })
                                .catch(err => {
                                    resolve();
                                });
                        })
                    );
                }
            }
            return Promise.all(requestArr);
        },
        bindEventsFunc() {
            this.validatorTypeSeleted = false;
        },
        bindEvents() {
            let node = document.querySelector("#app");
            node.addEventListener("click", this.bindEventsFunc, false);
        },
        unbindEvents() {
            let node = document.querySelector("#app");
            node.removeEventListener("click", this.bindEventsFunc, false);
        }
    }
};
</script>
<style lang="scss" scoped>
@import "../style/mixin.scss";

.blocks_list_page_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    width: 100% !important;
    .list_page_container {
        width: 100%;
        & > div {
            width: 100%;
            max-width: 12.8rem;
            min-width: 3rem;
            margin: 0 auto 0.4rem;
            .page_title {
                display: flex;
                align-items: center;
                & > span {
                    height: 0.7rem;
                    line-height: 0.7rem;
                    padding-left: 0.2rem;
                    font-size: 0.18rem;
                    margin-right: 0.12rem;
                }
                div.select_container {
                    width: 1rem;
                    height: 0.26rem;
                    border: 1px solid rgba(215, 217, 224, 1);
                    border-radius: 0.02rem;
                    font-size: 12px;
                    line-height: 0.26rem;
                    text-indent: 0.08rem;
                    color: #22252a;
                    position: relative;
                    z-index: 1;
                    cursor: pointer;
                    & > ul {
                        width: 1rem;
                        margin-left: -0.01rem;
                        border: 1px solid rgba(215, 217, 224, 1);
                        border-radius: 0.02rem;
                        background: #ffffff;
                        li {
                            &:hover {
                                background: rgba(235, 248, 255, 1);
                            }
                        }
                    }
                    &::after {
                        display: block;
                        content: "";
                        width: 0;
                        height: 0;
                        border-width: 6px 5px 6px 5px;
                        border-style: solid;
                        border-color: transparent;
                        border-top-color: #d7d9e0;
                        position: absolute;
                        top: 0.1rem;
                        right: 0.1rem;
                    }
                }
                div.active_select_container {
                    &::after {
                        top: 0.04rem;
                        border-top-color: transparent;
                        border-bottom-color: #d7d9e0;
                    }
                }
            }
            .tree_map_container {
                border: 1px solid #d7d9e0;
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
        }
    }
    .mobile_list_page_container {
        width: 100%;
        .tree_map_container {
            margin: 0 0.1rem;
        }
    }
}
</style>
