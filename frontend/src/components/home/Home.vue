<template>
    <div class="home_wrap">
        <advertising-space></advertising-space>
        <div class="home_container">
            <div class="information_preview">
                <ul class="current_net_status_list">
                    <li class="item_status">
                        <div class="img_container">
                            <div class="img_content">
                                <i class="iconfont iconBlockHeight"></i>
                            </div>
                            <span class="item_name">{{lang.home.blockHeight}}</span>
                        </div>
                        <p class="current_block proposer_content"><router-link @click.native="$uMeng.push('HomeOverview_Block Height','click')" :to="`/block/${currentBlockHeight}`">{{currentBlockHeight}}</router-link></p>
                        <p class="block_time proposer_container">
                            <span class="proposer_content"><router-link :to="addressRoute(proposerAddress)">{{moniker}}</router-link></span>
                        </p>
                    </li>
                    <li class="item_status">
                        <div class="img_container">
                            <div class="img_content">
                                <i class="iconfont iconTransactions"></i>
                            </div>
                            <span class="item_name">{{lang.home.transactions}}</span>
                        </div>
                        <p class="current_block transaction_link">
                            <router-link @click.native="$uMeng.push('HomeOverview_Transactions','click')" v-show="transactionValue !== '--'" :to="`/txs`">{{transactionValue}}</router-link>
                            <span v-show="transactionValue === '--'">--</span>
                        </p>
                        <p class="block_time">{{blockTime}}</p>
                    </li>
                    <li class="item_status">
                        <div class="img_container">
                            <div class="img_content">
                                <i class="iconfont iconAvgBlockTime"></i>
                            </div>
                            <span class="item_name">{{lang.home.ageTime}}</span>
                        </div>
                        <p class="current_block" :style="{color:diffSeconds > 120 ? '#ff001b' : ''}">{{averageBlockTime}}</p>
                        <p class="block_time">{{lang.home.latestBlocks}}</p>
                    </li>
                    <li class="item_status">
                        <div class="img_container">
                            <div class="img_content">
                                <i class="iconfont iconVotingPower"></i>
                            </div>
                            <span  class="item_name">{{lang.home.votingPower }}</span>
                        </div>
                        <p class="current_block">{{votingPowerValue}}</p>
                        <p class="block_time">{{validatorValue}}</p>
                    </li>
                    <li class="item_status">
                        <div class="img_container">
                            <div class="img_content">
                                <i class="iconfont iconBondedTokens"></i>
                            </div>
                            <span class="item_name">{{lang.home.bondedTokens}}</span>
                        </div>
                        <p class="current_block">{{bondedRatio}}</p>
                        <p class="block_time">{{bondedValue}}</p>
                    </li>
                </ul>
            </div>
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
        </div>
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
    import AdvertisingSpace from "../advertisingSpace/AdvertisingSpace";
    export default {
        name: 'app-header',
        components: {
            AdvertisingSpace,
            MVotingCard, MDepositCard, MHomeVotingCrad, EchartsPie, EchartsLine, HomeBlockModule},
        data() {
            return {
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
                navigationTimer:null,
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

        beforeMount () {
	        this.getBlocksList();
            this.getTransactionHistory();
            this.getTransactionList();
            this.getNavigation();
            this.getPorposakList();
        },
        mounted () {
            document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
            let that =this;
            clearInterval(this.timer);
            clearInterval(this.navigationTimer);
            if(this.$store.state.currentSkinStyle !=='default'){
	            this.setThemeStyle()
            }
            this.timer = setInterval(function () {
                that.getBlocksList();
                that.getTransactionHistory();
                that.getTransactionList();
                that.getValidatorsList();
            },10000);
            this.navigationTimer = setInterval(function() {
                that.getNavigation();
            },5000);
        },
        beforeDestroy () {
            clearInterval(this.timer)
        },
        watch: {
            '$store.state.isMobile'(newVal, oldVal) {
                this.onresize(newVal);
            },

            '$store.state.currentSkinStyle'(newVal){
                if(newVal !== 'default '){
                    this.setThemeStyle()
                }
            },
            '$store.state.hideTestSkinStyle'(newVal){
                this.$store.commit('hideTestSkinStyle',newVal)
            }
        },
        methods: {
            changeStyleToMainnet(){
                this.$store.commit('currentSkinStyle','irishub');
                this.$store.commit('testSkinStyle',true);
            },
            changeStyleToFuxiTestnet(){
                this.$store.commit('currentSkinStyle','fuxi');
                this.$store.commit('testSkinStyle',true);
            },
            changeStyleToNyancatTestnet(){
                this.$store.commit('currentSkinStyle','nyancat');
                this.$store.commit('testSkinStyle',true);
            },
       	    setThemeStyle(){
		        if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.IRISHUB){
			        this.themeStyleArray = this.mainnetThemeStyle;
		        }else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.FUXI){
			        this.themeStyleArray = this.testnetFuXiThemeStyle;
		        }else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.NYANCAT){
			        this.themeStyleArray = this.testnetNyancatThemeStyle;
		        }else {
			        this.themeStyleArray = this.defaultThemeStyle;
		        }
                this.getValidatorsList();
	        },
	        getPorposakList(){
	        	Service.commonInterface({homeProposalList:{
                    }},(res) => {
	        		try {
				        if(Array.isArray(res)){
					        res.forEach(item => {
						        if(item.status === "DepositPeriod"){
							        item.levelValue = item.level.name;
							        this.depositorBarArr.push(item);
							        this.depositorBarArr = this.depositorBarArr.sort((a,b) => {
                                        return b.proposal_id - a.proposal_id
                                    })
						        }
					        })
					        res.forEach(item => {
						        if(item.status === "VotingPeriod"){
							        this.votingBarArr.push(item)
							        this.votingBarArr = this.votingBarArr.sort((a,b) => {
								        return b.proposal_id - a.proposal_id
							        })
						        }
					        })
				        }
			        }catch (e) {
                        console.error(e)
			        }
                })
            },
            getValidatorsList () {
                Service.commonInterface({candidatesTop:{}},(data) => {
                	try {
		                if(data){
			                let [seriesData, legendData] = [[], []];
			                if (data.validators instanceof Array) {
				                let totalCount = 0;
				                data.validators.forEach(item=>totalCount += item.voting_power);
				                let others = data.power_all - totalCount;
				                let monikerReserveLength = 10;
				                let addressReserveLength = 6;
				                let powerAll = data.power_all;
				                for (let i = 0; i < data.validators.length; i++) {
					                seriesData.push({
						                value: data.validators[i].voting_power,
						                name: data.validators[i].description && data.validators[i].description.moniker ? `${Tools.formatString(data.validators[i].description.moniker,monikerReserveLength,"...")} (${Tools.formatString(data.validators[i].address,addressReserveLength,"...")})` : (data.validators[i].address ? data.validators[i].address : ''),
						                itemStyle: {color: this.themeStyleArray[i]},
						                emphasis : {itemStyle:{color: this.themeStyleArray[i]}},
						                upTime:`${data.validators[i].up_time}%`,
						                address:data.validators[i].address,
						                powerAll,
					                });
					                legendData.push(data.validators[i].description && data.validators[i].description.moniker ? `${Tools.formatString(data.validators[i].description.moniker,monikerReserveLength,"...")} (${Tools.formatString(data.validators[i].address,addressReserveLength,"...")})` : (data.validators[i].address ? data.validators[i].address : ''))
				                }
				                if(others > 0 ){
					                seriesData.push({
						                value: others,
						                name:'others',
						                powerAll,
						                itemStyle:{color:this.themeStyleArray[10]},
					                });
				                }

			                }
			                this.information = {legendData, seriesData}
		                }
	                }catch (e) {
                        console.error(e)
	                }

                });
            },
            getTransactionHistory () {
                Service.commonInterface({txsByDay:{}},(data) => {
                	try {
		                if(data){
			                let maxValue = 0;
			                if(data){
				                data.forEach(item=>{
					                if(item.num > maxValue){
						                maxValue = item.num;
					                }
				                });
				                let xData = data.map(item=>`${String(item.date).substr(5,2)}/${String(item.date).substr(8,2)}/${String(item.date).substr(0,4)}`);
				                let seriesData = data.map(item=>item.num);
				                this.informationLine = {maxValue,xData,seriesData};
			                }
		                }
	                }catch (e) {
                        console.error(e)
	                }

                })
            },
            showBlockFadeinAnimation (blockList) {
                let storedLastBlockHeight = localStorage.getItem('lastBlockHeight') ? localStorage.getItem('lastBlockHeight') : '';
                if(storedLastBlockHeight){
                    for(let index = 0; index < blockList.length; index++){
                        if(blockList[index].height > storedLastBlockHeight){
                            blockList.forEach(item => {
                                item.flShowTranslationalAnimation = true
                            });
                            blockList[index].showAnimation = "show";
                        }
                    }
                }
            },
            getBlocksList() {
                Service.commonInterface({blocksRecent:{}},(blockList)=>{
                	try {
		                if(blockList){
			                this.showBlockFadeinAnimation(blockList);
			                let that = this;
			                setTimeout(function () {
				                that.blocksInformation.map(item => {
					                return item.flShowTranslationalAnimation = false
				                })
			                },1000);
			                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
			                localStorage.setItem("lastBlockHeight",blockList[0].height);

			                this.blocksInformation = blockList.map(item => {
				                return {
					                flShowTranslationalAnimation :  item.flShowTranslationalAnimation ? item.flShowTranslationalAnimation : "",
					                showAnimation: item.showAnimation ? item.showAnimation : "",
					                Height: item.height,
					                Txn: item.num_txs,
					                Time: Tools.format2UTC(item.time),
					                time:item.time,
					                age: Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX)
				                };
			                });
			                this.blocksTimer = setInterval(function () {
				                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
				                that.lastBlockAge = Tools.formatAge(currentServerTime,blockList[0].time);
				                that.blocksInformation = that.blocksInformation.map(item => {
					                item.age = Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX);
					                return item
				                })
			                },1000);
		                }
	                }catch (e) {
                        console.log(e)
	                }
                })
            },
            getTransactionList() {
                Service.commonInterface({txsRecent:{}},(transactionList) => {
                	try {
		                if(transactionList){
			                let that = this;
			                for (let txIndex = 0; txIndex < transactionList.length; txIndex++){
				                if(new Date(transactionList[txIndex].time).getTime() > localStorage.getItem("lastTxTime")){
					                transactionList[txIndex].showAnimation = "show";
					                transactionList.forEach(item => {
						                item.flShowTranslationalAnimation = true
					                })
				                }
			                }
			                setTimeout(function () {
				                that.transactionInformation.map(item => {
					                return item.flShowTranslationalAnimation = false
				                })
			                },1000);
			                let lastTxTime = new Date(transactionList[0].time).getTime();
			                localStorage.setItem('lastTxTime',lastTxTime);
			                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
			                that.transactionInformation = transactionList.map(item => {
				                let [Fee] = ['--'];
				                if(item.actual_fee.amount && item.actual_fee.denom){
					                Fee =  `${Tools.formatFeeToFixedNumber(item.actual_fee.amount)} ${Tools.formatDenom(item.actual_fee.denom).toUpperCase()}`;
				                }
				                let currentServerTime = new Date().getTime() + that.diffMilliseconds;
				                return {
					                flShowTranslationalAnimation :  item.flShowTranslationalAnimation ? item.flShowTranslationalAnimation : "",
					                showAnimation: item.showAnimation ? item.showAnimation : '',
					                TxHash: item.tx_hash,
					                From: item.from,
					                To: item.to,
					                Type: item.type === 'coin'?'transfer':item.type,
					                Fee,
					                Time: item.time,
					                age: Tools.formatAge(currentServerTime,item.time,Constant.SUFFIX,Constant.PREFIX)
				                };
			                });
			                clearInterval(this.transfersTimer);
			                this.transfersTimer = setInterval(function () {
				                that.transactionInformation.map(item => {
					                currentServerTime = new Date().getTime() + that.diffMilliseconds;
					                lastTxTime = new Date(transactionList[0].time).getTime();
					                item.age = Tools.formatAge(currentServerTime,item.Time,Constant.SUFFIX,Constant.PREFIX);
					                return item
				                })
			                },1000)
		                }
	                }catch (e) {
                        console.error(e)
	                }
                })
            },
            getNavigation(){
                Service.commonInterface({navigation:{}},(res) => {
                	try {
		                if(res){
			                let reservedStringLength = 12;
			                this.moniker = Tools.formatString(res.moniker,reservedStringLength,'...');
			                this.proposerAddress = res.operator_addr;
			                this.diffSeconds = Number(res.avg_block_time);
			                this.currentBlockHeight = res.block_height;
			                this.transactionValue = this.formatTransactions(res.total_txs);
			                this.averageBlockTime = `${Number(res.avg_block_time).toFixed(2)} s`;
			                this.votingPowerValue = `${Number(res.voting_ratio * 100).toFixed(2)} %`;
			                this.bondedValue = this.formatBondedTokens(res.bonded_tokens,res.total_supply);
			                this.validatorValue = `${res.vote_val_num} / ${res.active_val_num} Validators`;
			                this.bondedRatio = `${(res.bonded_ratio * 100).toFixed(2)} %`;
			                this.blockTime = Tools.format2UTC(res.block_time);
		                }
	                }catch (e) {
                        console.error(e)
	                }
                });
            },
            formatTransactions(totalTxs){
                let num, thousand = 1000,million = 1000000,billion = 1000000000;
                if(totalTxs > billion){
                    num = `${(totalTxs/billion).toFixed(2)} B`;
                }else if(totalTxs > million){
                    num = `${(totalTxs/million).toFixed(2)} M`;
                }else if(totalTxs > thousand){
                    num = `${(totalTxs/thousand).toFixed(2)} K`;
                }else {
                    num = totalTxs
                }
                return num
            },
            formatBondedTokens(bondedTokens,totalTokens){
                let tokens,allTokens,thousand = 1000,million = 1000000,billion = 1000000000;
                if(bondedTokens >= billion){
	                tokens = `${(Number(bondedTokens) / billion).toFixed(2)}B`
                }else if(bondedTokens >= million){
	                tokens = `${(Number(bondedTokens) / million).toFixed(2)}M`
                }else if(bondedTokens >= thousand){
	                tokens = `${(Number(bondedTokens) / thousand).toFixed(2)}k`
                }else {
	                tokens = `${Number(bondedTokens).toFixed(2)}`
                }

	            if(totalTokens >= billion){
		            allTokens = `${(Number(totalTokens) / billion).toFixed(2)}B`
	            }else if(totalTokens >= million){
		            allTokens = `${(Number(totalTokens) / million).toFixed(2)}M`
	            }else if(totalTokens >= thousand){
		            allTokens = `${(Number(totalTokens) / thousand).toFixed(2)}k`
	            }else {
		            allTokens = `${Number(totalTokens).toFixed(2)}`
	            }
                return `${tokens} / ${allTokens}`
            },
        },
        destroyed () {
            clearInterval(this.timer);
            clearInterval(this.navigationTimer);
            }
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
            .information_preview {
                display: flex;
                margin-bottom: 0.3rem;
                .current_net_status_list {
                    display: flex;
                    width: 100%;
                    .item_status {
                        flex: 1;
                        background: #fff;
                        border: 0.01rem solid rgba(215, 217, 224, 1);
                        border-radius: 0.01rem;
                        margin-right: 3%;
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
            }
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
