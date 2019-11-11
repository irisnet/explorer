<template>
    <div class="home_wrap">
        <div :class="pageClassName">
            <div class="information_preview">
                <ul class="current_net_status_list">
                    <li class="item_status">
                        <div class="img_container">
                            <div class="img_content">
                                <i class="iconfont iconBlockHeight"></i>
                            </div>
                            <span class="item_name">{{lang.home.blockHeight}}</span>
                        </div>
                        <p class="current_block proposer_content"><router-link :to="`/block/${currentBlockHeight}`">{{currentBlockHeight}}</router-link></p>
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
                        <p class="current_block transaction_link"><router-link :to="`/txs`">{{transactionValue}}</router-link></p>
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
            <div :class="module_item_wrap">
                <div class="home_module_item home_module_item_pie">
                    <echarts-pie :information="information"></echarts-pie>
                </div>
                <div class="home_module_item home_module_item_pie">
                    <echarts-line :informationLine="informationLine"></echarts-line>
                </div>
            </div>
            <div :class="module_item_wrap">
                <div class="home_module_item fixed_item_height">
                    <home-block-module :moduleName="'Blocks'" :information="blocksInformation"></home-block-module>
                </div>
                <div class="home_module_item fixed_item_height">
                    <home-block-module :moduleName="'Transactions'" :information="transactionInformation"></home-block-module>
                </div>
            </div>
            <div class="home_proposal_container">
                <div class="home_proposal_item_bar" v-for="item in votingBarArr" :key="item.proposal_id">
                    <m-voting-card :votingBarObj="item" :showTitle="true"></m-voting-card>
                </div>
                <div class="home_proposal_item_bar" v-for="v in depositorBarArr" :key="v.proposal_id">
                    <m-deposit-card :depositObj="v" :showTitle="true" :levelValue="v.levelValue"></m-deposit-card>
                </div>
            </div>
        </div>
    </div>

</template>

<script>
    import Tools from '../util/Tools';
    import EchartsPie from "../components/EchartsPie";
    import EchartsLine from "../components/EchartsLine";
    import HomeBlockModule from "../components/HomeBlockModule";
    import Service from '../service/index';
    import Constant from "../constant/Constant";
    import lang from "../lang/index"
    import MHomeVotingCrad from "../components/commonComponents/MHomeVotingCrad";
    import MDepositCard from "../components/commonComponents/MDepositCard";
    import MVotingCard from "../components/commonComponents/MVotingCard";
    export default {
        name: 'app-header',
        components: {MVotingCard, MDepositCard, MHomeVotingCrad, EchartsPie, EchartsLine, HomeBlockModule},
        data() {
            return {
                devicesWidth: window.innerWidth,
                pageClassName: 'personal_computer_home_wrap',
                module_item_wrap: 'module_item_wrap_computer',
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
            if (!this.$store.state.isMobile) {
                this.pageClassName = 'personal_computer_home_wrap';
                this.module_item_wrap = 'module_item_wrap_computer';
                if(document.getElementsByClassName('fixed_item_height').length > 0){
                    document.getElementsByClassName('fixed_item_height')[0].style.height = '6.55rem';
                    document.getElementsByClassName('fixed_item_height')[1].style.height = '6.55rem';
                }
            } else {
                this.pageClassName = 'mobile_home_wrap';
                this.module_item_wrap = 'module_item_wrap_mobile';
            }

        },
        beforeDestroy () {
            window.removeEventListener('resize',this.onWindowResize);
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
            }
        },
        methods: {
        	setThemeStyle(){
		        if(this.$store.state.currentSkinStyle ===  `${Constant.ENVCONFIG.MAINNET}${Constant.CHAINID.IRISHUB}`){
			        this.themeStyleArray = this.mainnetThemeStyle;
		        }else if(this.$store.state.currentSkinStyle ===  `${Constant.ENVCONFIG.TESTNET}${Constant.CHAINID.FUXI}`){
			        this.themeStyleArray = this.testnetFuXiThemeStyle;
		        }else if(this.$store.state.currentSkinStyle ===  `${Constant.ENVCONFIG.TESTNET}${Constant.CHAINID.NYANCAT}`){
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
            onresize (isMobile) {
                this.innerWidth = window.innerWidth;
                if(!isMobile){
                    this.pageClassName = 'personal_computer_home_wrap';
                    this.module_item_wrap = 'module_item_wrap_computer';
                    if(document.getElementsByClassName('fixed_item_height').length > 0) {
                        document.getElementsByClassName('fixed_item_height')[0].style.height = '6.55rem';
                        document.getElementsByClassName('fixed_item_height')[1].style.height = '6.55rem';
                    }
                }else {
                    this.pageClassName = 'mobile_home_wrap';
                    this.module_item_wrap = 'module_item_wrap_mobile';
                    if(document.getElementsByClassName('fixed_item_height').length > 0) {
                        document.getElementsByClassName('fixed_item_height')[0].style.height = 'auto';
                        document.getElementsByClassName('fixed_item_height')[1].style.height = 'auto';
                    }
                }
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
    @import '../style/mixin.scss';
    .home_wrap {
        @include flex();
        @include pcContainer;
        //pc端和移动端公共样式
        .personal_computer_home_wrap, .mobile_home_wrap {
            padding-top: 0.3rem;
            .information_preview {
                @include flex;
                margin-bottom: 0.3rem;
                .information_preview_module {
                    border-right: 1px solid #d6d9e0;
                    @include flex;
                    flex-direction: column;
                    align-items: center;
                    span:nth-child(1){
                        height: 0.27rem;
                    }
                    &:last-child {
                        border-right: none;
                    }
                    span {
                        &:first-child {
                            font-size: 0.18rem;
                            @include fontWeight;
                        }
                    }
                    .information_module_key {
                        font-size: 0.14rem;
                        color: var(--contentColor);
                    }
                }
            }
            .home_proposal_container{
                display: flex;
                justify-content: space-between;
                flex-wrap: wrap;
                box-sizing: border-box;
                .home_proposal_item_bar{
                    min-width: 5.98rem;
                    flex: 1;
                    width: auto;
                    margin-right: 0.4rem;
               }
                .home_proposal_item_bar:nth-child(even){
                    margin-right: 0;
                }
                .home_proposal_item_bar:last-child{
                    margin-right: 0;
                }
            }
            .current_net_status_list{
                display: flex;
                width: 100%;
                .item_status{
                    flex: 1;
                    background: #fff;
                    border: 0.01rem solid rgba(215, 217, 224, 1);
                    border-radius: 0.01rem;
                    margin-right: 3%;
                    box-sizing: border-box;
                    padding: 0.14rem;
                    .img_container{
                        display: flex;
                        align-items: center;
                        .img_content{
                            display: flex;
                            align-items: center;
                            width: 0.15rem;
                            i{
                                font-size: 0.16rem;
                                color: var(--bgColor)
                            }
                            img{
                                width: 100%;
                            }
                        }
                        .item_name{
                            margin-left: 0.08rem;
                            font-size: 0.14rem;
                            color: var(--contentColor);
                        }
                    }
                    .current_block{
                        padding-top: 0.2rem;
                        font-size: 0.2rem;
                        line-height: 0.23rem;
                        color: var(--titleColor);
                    }
                    .transaction_link{
                        a{
                            color: var(--bgColor) !important;
                        }
                    }
                    .block_time{
                        color: var(--contentColor);
                        font-size: 0.14rem;
                        padding-top: 0.1rem;
                    }
                }
                .item_status:last-child{
                    margin-right: 0;
                }
            }
            //饼状图
            .home_module_item_pie {
                height: 3.5rem;
                background: #fff;
            }
        }

        .personal_computer_home_wrap {
            width: 100%!important;
            padding: 0.3rem 0.2rem 0 0.2rem;
            @include pcCenter;
            .information_preview {
                .information_preview_module {
                    flex: 1;
                }
            }

            .module_item_wrap_computer {
                width: 100%;
                @include flex();
                justify-content: space-between;
                .home_module_item {
                    flex:1;
                    border: 0.01rem solid #d6d9e0;
                    margin-bottom: 0.3rem;
                    height: 3.54rem;
                    background: #fff;
                    &:nth-child(2n+1){
                        margin-right:0.4rem;
                    }
                }
            }

        }
        .mobile_home_wrap {
            flex-direction: column;
            justify-content: space-between;
            width: 100%;
            padding: 0.1rem;
            .information_preview {
                overflow-x: auto;
                -webkit-overflow-scrolling:touch;
                .current_net_status_list{
                    display: flex;
                    flex-direction: column;
                    .item_status{
                        min-height: 1.24rem;
                        border-radius: 0.01rem;
                        margin-top: 0.1rem;
                        margin-right: 0;
                        box-sizing: border-box;
                        padding: 0.14rem;
                        width: 100%;
                        border: 0.01rem solid rgba(215, 217, 224, 1);
                    }
                }
            }
            .module_item_wrap_mobile {
                @include flex();
                flex-direction: column;
                align-items: center;
                .home_module_item {
                    width: 100%;
                    margin-bottom: 0.4rem;
                    border: 0.01rem solid #d6d9e0;
                }
                .home_module_item_pie {
                    overflow-x: auto;
                    -webkit-overflow-scrolling:touch;
                }
            }
        }
        .home_module_item_status {
            padding: 0.1rem;
            background: var(--bgColor);

            .current_block {
                @include fontWeight;
                color: #fff;
                display: inline-block;
                height: 0.28rem;
                line-height: 0.28rem;
                border-bottom: 0.01rem solid #ffffff;
                width: 100%;

            }
            .home_status_bottom {
                @include flex();

                .home_status_bottom_content {
                    flex: 1;
                    @include flex();
                    flex-direction: column;
                    span {
                        font-size: 0.14rem;
                        color: #ffffff;
                        font-weight: normal;
                    }
                }

            }
        }
    }
    .blocks_background_type{
        background: url('../assets/block.png') no-repeat 0 0.02rem;
        text-indent:0.2rem;
        color:var(--bgColor);
    }
    .home_module_transaction_title_wrap{
        @include flex;
        padding:0.2rem;
        height:0.64rem;
        justify-content: space-between;
        background: #efeff1;
        border-bottom:1px solid #e4e4e4;
        align-items: center;
        .home_module_transaction_title{
            font-size:0.18rem;
            @include fontWeight;
        }
        .blocks_background{
            background: url('../assets/blocks.png') no-repeat 0 0.02rem;
            text-indent:0.35rem;
        }
        .transactions_background{
            background: url('../assets/transactions.png') no-repeat 0 0.02rem;
            text-indent:0.3rem;
        }
        .view_all_btn{
            @include viewBtn;
        }
    }
    .transaction_title_name{
        padding-left: 0.1rem;
    }
    .transactions_background_type{
        background: url('../assets/transactions.png') no-repeat 0 0.02rem;
        text-indent:0.2rem;
    }
    pre{
        margin: 0!important;
    }
    .animation{
        @include fadeInAnimation
    }
    .latest_block_content:hover{
        cursor: pointer;
    }
    .proposer_content{
        cursor: pointer;
        a{
            color:var(--bgColor) !important;
        }
    }

    @media screen and (min-width: 910px) and (max-width: 1280px){
        .home_wrap{
            .personal_computer_home_wrap{
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
            .mobile_home_wrap{
                .home_proposal_container{
                    display: flex;
                    flex-direction: column;
                    margin-bottom: 0.2rem;
                    overflow-x: auto;
                    .home_proposal_item_bar{
                        margin-right: 0;
                        margin-bottom: 0.2rem;
                    }
                }
            }
        }
    }
</style>
