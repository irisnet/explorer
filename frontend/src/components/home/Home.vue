<template>
    <div class="home_wrap">
        <validator-graph></validator-graph>
       <!-- <search-bar></search-bar>
        <statistical-bar></statistical-bar>
        <div class="home_container">
            <div class="echarts_content">
                <div class="home_module_item home_module_item_pie">
                    <echarts-pie :information="information"></echarts-pie>
                </div>
                <div class="home_module_item home_module_item_pie">
                    <echarts-line :informationLine="informationLine"></echarts-line>
                </div>
            </div>
            <div class="home_module_container">
                <div class="home_module_item fixed_item_height">
                    <home-block-module :moduleName="'Blocks'" :information="blocksInformation"></home-block-module>
                </div>
                <div class="home_module_item fixed_item_height">
                    <home-block-module :moduleName="'Transactions'" :information="transactionInformation"></home-block-module>
                </div>
            </div>
            <div class="home_proposal_container">
                <div class="home_proposal_container_content" v-for="item in votingBarArr" :key="item.proposal_id">
                    <div class="home_proposal_item_bar">
                        <m-voting-card :votingBarObj="item" :showTitle="true"></m-voting-card>
                    </div>
                </div>
                <div class="home_proposal_container_content" v-for="v in depositorBarArr" :key="v.proposal_id">
                    <div class="home_proposal_item_bar">
                        <m-deposit-card :depositObj="v" :showTitle="true" :levelValue="v.levelValue"></m-deposit-card>
                    </div>
                </div>
            </div>
        </div>
        <div class="test_skin_content">
            <div style="display: flex" v-if="$store.state.hideTestSkinStyle">
                <div class="mainnet_irishub" @click="changeStyleToMainnet">Mainnet irishub</div>
                <div class="testnet_fuxi"@click="changeStyleToFuxiTestnet">Testnet FuXi</div>
                <div class="testnet_nyancat"@click="changeStyleToNyancatTestnet">Testnet Nyancat</div>
            </div>
        </div>-->
    </div>
</template>

<script>
    import Tools from '../../util/Tools';
    import EchartsPie from "./EchartsPie";
    import EchartsLine from "./EchartsLine";
    import HomeBlockModule from "./HomeBlockModule";
    import Service from '../../service';
    import Constant from "../../constant/Constant";
    import lang from "../../lang"
    import MHomeVotingCrad from "../commontables/MHomeVotingCrad";
    import MDepositCard from "../commontables/MDepositCard";
    import MVotingCard from "../commontables/MVotingCard";
    import axios from "axios"

    import SearchBar from "./SearchBar";
    import StatisticalBar from "./StatisticalBar";
    import ValidatorGraph from "./ValidatorGraph";
    export default {
        name: 'app-header',
        components: {
            ValidatorGraph,
            StatisticalBar,
            SearchBar, MVotingCard, MDepositCard, MHomeVotingCrad, EchartsPie, EchartsLine, HomeBlockModule},
        data() {
            return {
                circulationBondedToken:'--',
                circulationBondedTokenValue:'--',
                currentBlockHeight: '',
                validatorValue: '--',
                transactionValue: '--',
                averageBlockTime: '--',
                ageTime: '--',
                votingPowerValue: '--',
                blockHeightValue: '--',
                timestampValue: '--',
                bondedValue:'--',
                bondedRatio:'--',
                blockTime : '--',
                information: {},
                informationLine: {},
                blocksInformation: [],
                transactionInformation: [],
                innerWidth : window.innerWidth,
                blocksTimer:null,
                latestBlockTimer:null,
                transfersTimer:null,
                diffSeconds: 0,
                timer: null,
                lang:lang,
                isMobile: false,
                moniker:'',
                proposerAddress:"",
                depositorBarArr: [],
                votingBarArr: [],
	            levelValue:'',
                mainnetThemeStyle:['#3264FD', '#4571FA', '#537CFD', '#6287FB', '#7092FD', '#85A3FF', '#92ACFF', '#9CB3FF', '#AABEFF', '#BDCDFF', '#E9EBFC',],
                testnetFuXiThemeStyle:['#0C4282', '#144B8D', '#1B5498', '#235CA1', '#2E65A8', '#386EAE', '#427ABC', '#5087C8', '#5992D5', '#69A1E2', '#E9EBFC',],
                testnetNyancatThemeStyle:['#0D9388', '#149A8F', '#1DA196', '#26ABA0', '#32B5AA', '#3FBDB2', '#4EC2B8', '#59C8BE', '#64D1C7', '#70DCD2', '#E9EBFC',],
                defaultThemeStyle:['#3498db', '#47a2df', '#59ade3', '#6cb7e7', '#7fc2eb', '#91ccef', '#a4d7f3', '#b7e1f7', '#c9ecfb', '#dcf6ff', '#f0f9ff',],
                themeStyleArray:'',
                sessionStorageCurrentEnv:sessionStorage.getItem('skinCurrentEnv')
            }
        },
    }
</script>
<style lang="scss">
    @import '../../style/mixin';
    .home_wrap {
        display: flex;
        width: 100%;
        flex-direction: column;
        align-items: center;
        .test_skin_content{
            display: none;
            margin: 0.2rem 0;
            color: #fff;
            cursor: pointer;
            .mainnet_irishub{
                padding: 0.05rem 0.2rem;
                background: #3264FD;
                border-radius: 0.04rem;
                margin-right: 0.1rem;

            }
            .testnet_fuxi{
                padding: 0.05rem 0.2rem;
                background: #0C4282;
                border-radius: 0.04rem;
                margin-right: 0.1rem;
            }
            .testnet_nyancat{
                padding: 0.05rem 0.2rem;
                background: #0D9388;
                border-radius: 0.04rem;
            }
        }

        .home_container {
            display: flex;
            flex-direction: column;
            max-width: 12.8rem;
            width: 100%;
            box-sizing: border-box;
            padding: 0.2rem 0.2rem 0 0.2rem;
            /*.information_preview {
                display: flex;
                flex-direction: column;
                margin-bottom: 0.3rem;
                .current_net_status_list {
                    display: flex;
                    width: 100%;
                    .item_status {
                        flex: 1;
                        background: #fff;
                        border: 0.01rem solid rgba(215, 217, 224, 1);
                        border-radius: 0.01rem;
                        margin-right: 0.2rem;
                        box-sizing: border-box;
                        padding: 0.14rem;
                        .img_container {
                            display: flex;
                            align-items: center;
                            line-height: 0.24rem;
                            .img_content {
                                display: flex;
                                align-items: center;
                                width: 0.15rem;

                                i {
                                    font-size: 0.16rem;
                                    color: var(--bgColor)
                                }

                                img {
                                    width: 100%;
                                }
                            }

                            .item_name {
                                margin-left: 0.08rem;
                                font-size: 0.14rem;
                                color: var(--contentColor);
                            }
                        }
                        .current_block {
                            padding-top: 0.2rem;
                            font-size: 0.2rem;
                            line-height: 0.23rem;
                            color: var(--titleColor);
                            a {
                                color: var(--bgColor) !important;
                            }
                        }
                        .proposer_content{
                            cursor: pointer;
                            a{
                                color:var(--bgColor) !important;
                            }
                        }
                        .block_time {
                            color: var(--contentColor);
                            font-size: 0.14rem;
                            padding-top: 0.1rem;
                        }
                    }

                    .item_status:last-child {
                        margin-right: 0;
                    }
                }
            }*/
            .echarts_content{
                width: 100%;
                display: flex;
                justify-content: space-between;
                .home_module_item {
                    flex:1;
                    border: 0.01rem solid #d6d9e0;
                    margin-bottom: 0.4rem;
                    background: #fff;
                    height: 3.52rem;
                    &:nth-child(2n+1){
                        margin-right:0.4rem;
                    }
                }
            }
            .home_module_container{
                width: 100%;
                display: flex;
                justify-content: space-between;
                .home_module_item {
                    flex:1;
                    border: 0.01rem solid #d6d9e0;
                    margin-bottom: 0.4rem;
                    background: #fff;
                    &:nth-child(2n+1){
                        margin-right:0.4rem;
                    }
                }
            }
            .home_proposal_container {
                display: flex;
                justify-content: space-between;
                flex-wrap: wrap;
                box-sizing: border-box;
                .home_proposal_container_content{
                    .home_proposal_item_bar {
                        min-width: 5.98rem;
                        flex: 1;
                        width: auto;
                        margin-right: 0.4rem;
                    }

                    .home_proposal_item_bar:nth-child(even) {
                        margin-right: 0;
                    }

                    .home_proposal_item_bar:last-child {
                        margin-right: 0;
                    }
                }

            }
        }
    }
    @media screen and (min-width: 910px) and (max-width: 1280px){
        .home_wrap{
            .home_container{
                .home_proposal_container{
                    display: flex;
                    flex-direction: column;
                    .home_proposal_item_bar{
                        margin-right: 0;
                    }
                    .home_proposal_item_bar:nth-child(even){
                        margin-right: 0;
                    }
                    .home_proposal_item_bar:last-child{
                        margin-right: 0;
                    }
                }

            }
        }
    }
@media screen and (max-width: 910px){
    .home_wrap{
        .home_container{
            box-sizing: border-box;
            padding: 0.2rem 0.1rem 0 0.1rem;
            .information_preview{
                .current_net_status_list{
                    flex-direction: column;
                    .item_status{
                        margin-right: 0;
                        margin-bottom: 0.1rem;
                    }
                }
                .current_net_status_list:last-child{
                    margin-top: 0;
                }
            }
            .echarts_content{
                flex-direction: column;
                .home_module_item:nth-child(2n+1){
                    margin-right: 0;
                }
                .home_module_item_pie{
                    overflow-x: auto;
                    overflow-y: hidden;
                }
            }
            .home_module_container{
                flex-direction: column;
                .home_module_item:nth-child(2n+1){
                    margin-right: 0;
                }
            }
            .home_proposal_container{
                display: flex;
                flex-direction: column;
                margin-bottom: 0.2rem;
                .home_proposal_container_content{
                    width: 100%;
                    overflow-x: auto;
                    overflow-y: hidden;
                    .home_proposal_item_bar{
                        margin-right: 0;
                        margin-bottom: 0.2rem;
                        .deposit_card_content{
                            margin-bottom: 0;
                        }
                    }

                }

            }
        }
    }
}

</style>
