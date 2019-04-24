<!--头部页面-->
<template>
  <div :class="appHeaderVar" v-show="showHeader" id="header">
    <header class="app_header_person_computer" v-show="devicesShow === 1">
      <div class="header_top">
        <div class="header_left">
          <div class="imageWrap" @click="featureButtonClick('/home')">
            <img src="../assets/logo.png"/>
          </div>
          <div class="network_select_container" v-show="flShowHeaderNetwork">
            <div class="network_select_option" @click.stop="toggleSelectOption()">
              <div class="select_content">
                <div class="current_select_content">{{currentSelected}} <i class="iconfont icon-arrowdown"></i></div>
                <div class="select_option_content" v-if="$store.state.flShowSelectOption">
                  <div class="common_option" v-if="flShowMainnet" @click="toMainnet(mainnetHref)">{{lang.home.mainnet}}</div>
                  <div class="common_option" v-if="flShowTestnet" @click="toTestnet(testnetHref)">{{lang.home.testnet}}</div>
                  <div class="common_option" @click="toNyanCats(nyanCatsHref)">{{lang.home.catsNet}}</div>
                </div>
              </div>
            </div>
          </div>
          <div class="chain_id_container" v-if="flShowChainId">
            <span>{{chainId}}</span>
          </div>
        </div>
        <div class="navSearch">
          <input type="text" class="search_input"
                 placeholder="Search by Address / Txhash / Block"
                 v-model.trim="searchInputValue"
                 @keyup.enter="onInputChange">
          <i class="search_icon" @click="getData(searchInputValue)"></i>
          <i class="clear_icon" @click="clearSearchContent" v-show="showClear"></i>
        </div>
      </div>

    </header>
    <div class="useFeature" v-show="devicesShow === 1">
      <div class="navButton">
        <span class="nav_item common_item_style" :class="activeClassName === '/home'?'nav_item_active':''"
              @click="featureButtonClick('/home')"
        >Home</span>
        <div class="nav_item sub_btn_wrap common_item_style" :class="activeClassName === '/validators'?'nav_item_active':''">
          <span class="nav_item common_item_style" @click="featureButtonClick('/validators')">
            Validators
          </span>
        </div>
        <span class="nav_item common_item_style" :class="activeClassName === '/block'?'nav_item_active':''"
              @click="featureButtonClick('/block/1/0')"
        >Blocks</span>
        <div class="nav_item sub_btn_wrap common_item_style" :class="activeClassName === '/transaction'?'nav_item_active':''"
             @mouseover="transactionMouseOver" @mouseleave="transactionMouseLeave">

          <span class="nav_item common_item_style">
            Transactions
            <span class="bottom_arrow"></span>
          </span>
          <span class="sub_btn_item" @click="featureButtonClick('/transactions/2/Transfers')"
                v-show="showSubTransaction">Transfers</span>
          <span class="sub_btn_item" @click="featureButtonClick('/transactions/2/Declarations')"
                v-show="showSubTransaction">Declarations</span>
          <span class="sub_btn_item" @click="featureButtonClick('/transactions/2/Stakes')"
                v-show="showSubTransaction">Stakes</span>
          <span class="sub_btn_item" @click="featureButtonClick('/transactions/2/Governance')"
                v-show="showSubTransaction">Governance</span>
        </div>
        <span class="nav_item common_item_style" :class="activeClassName === '/Proposals'?'nav_item_active':''"
              @click="featureButtonClick('/Proposals')"
        >Proposals</span>
        <span v-if="flShowFaucet" class="nav_item common_item_style faucet_content" :class="activeClassName === '/faucet'?'nav_item_active':''"
              @click="featureButtonClick('/faucet')"
        >Faucet</span>
      </div>
    </div>
    <div class="app_header_mobile" v-show="devicesShow === 0" ref="header_content">
      <div style="display: flex;flex-direction: row-reverse;justify-content: space-between;align-items: center;padding: 0 0.15rem;">
        <div class="feature_btn" @click="showFeature">
          <img src="../assets/menu.png">
        </div>
        <div class="image_wrap_mobile" @click="featureButtonClick('/home',true)">
          <img src="../assets/logo.png"/>
        </div>
      </div>
      <div class="search_input_mobile">
        <div style="width:100%;display: flex;justify-content: center;">
          <input type="text" class="search_input"
                 v-model.trim="searchInputValue"
                 @keyup.enter="onInputChange"
                 placeholder="Search by Address / Txhash / Block">
          <i class="search_icon" @click="getData(searchInputValue)"></i>
          <i class="clear_icon" @click="clearSearchContent" v-show="showClear"></i>
        </div>
      </div>
      <div class="mobile_chain_id_content" v-if="flShowChainId">
        <span class="mobile_chain_content">{{chainId}}</span>
      </div>
      <div class="use_feature_mobile" :style="{'top':absoluteTop}" v-show="featureShow">
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/home')">Home</span>
        <span class="feature_btn_mobile feature_nav select_option_container" @click="featureButtonClick('/validators')">
         <span>Validators</span>
        </span>
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/block/1/0')">Blocks</span>
        <span class="feature_btn_mobile feature_nav select_option_container" @click="transactionsSelect(flShowTransactionsSelect)">
         <span>Transactions</span>
          <div :class="flShowUpOrDown ? 'upImg_content' : 'downImg_content'">
            <img :src="flShowUpOrDown ? upImg : downImg ">
          </div>
        </span>
        <div class="select_option" v-show="flShowTransactionsSelect">
             <span class="feature_btn_mobile feature_nav"
                   @click="featureButtonClick('/transactions/2/Transfers')">Transfers</span>
          <span class="feature_btn_mobile feature_nav"
                @click="featureButtonClick('/transactions/2/Declarations')">Declarations</span>

          <span class="feature_btn_mobile feature_nav"
                @click="featureButtonClick('/transactions/2/Stakes')">Stakes</span>
          <span class="feature_btn_mobile feature_nav"
                @click="featureButtonClick('/transactions/2/Governance')">Governance</span>
        </div>

        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/Proposals')">Proposals</span>
        <span v-if="flShowFaucet" class="feature_btn_mobile feature_nav mobile_faucet_content" @click="featureButtonClick('/faucet')">Faucet</span>

        <span class="feature_btn_mobile feature_nav select_option_container" @click="netWorkSelect(flShowNetworkSelect)">
         <span>Network</span>
          <div :class="flShowNetworkUpOrDown ? 'upImg_content' : 'downImg_content'">
            <img :src="flShowNetworkUpOrDown ? upImg : downImg ">
          </div>
        </span>
        <div class="select_option" v-show="flShowNetworkSelect">
          <span class="feature_btn_mobile feature_nav"><a :href="mainnetHref" target="_blank">Mainnet</a></span>
          <span class="feature_btn_mobile feature_nav"><a :href="testnetHref" target="_blank">FUXI Testnet</a></span>
          <span class="feature_btn_mobile feature_nav"><a :href="nyanCatsHref" target="_blank">NYAN_CATS Testnet</a></span>
        </div>

      </div>
    </div>
  </div>
</template>
<script>
  import Tools from '../util/Tools';
  import Service from "../util/axios";
  import constant from '../constant/Constant';
  import lang from "../lang/index"
  export default {
    name: 'app-header',
    watch:{
      $route(){
        this.searchInputValue = "";
        this.listenRouteForChangeActiveButton();
        this.showHeader = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
      },
      searchInputValue(searchInputValue){
        if(searchInputValue){
          this.showClear = true;
        }else{
          this.showClear = false;
        }
      },
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        devicesShow: 1,
        searchValue: '',
        appHeaderVar: 'person_computer_header_var',
        featureShow: false,
        transactionShow: false,
        validatorsShow: false,
        searchInputValue: '',
        activeClassName: '/home',
        showHeader:!(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
        flShowFaucet: false,
        showSubTransaction:false,
        showSubValidators:false,
        showClear:false,
        innerWidth : window.innerWidth,
        flShowTransactionsSelect: false,
        flShowValidatorsSelect: false,
        flShowNetworkSelect:false,
        flShowUpOrDown: false,
        flShowNetwork: false,
        flShowHeaderNetwork: false,
        flShowChainId: false,
        flShowValidatorsUpOrDown: false,
        flShowNetworkUpOrDown: false,
        flShowMainnet:false,
        flShowTestnet:false,
        currentSelected:'',
        absoluteTop:'',
        lang: lang,
        chainId: '',
        mainnetHref:'',
        testnetHref:'',
        nyanCatsHref:'',
        upImg: require("../assets/caret-bottom.png"),
        downImg: require("../assets/caret-bottom.png"),
    }
    },
    beforeMount() {
      if (window.innerWidth > 910) {
        this.devicesShow = 1;
        this.appHeaderVar = 'person_computer_header_var';
      } else {
        this.devicesShow = 0;
        this.appHeaderVar = 'mobile_header_var';
      }
    },
    mounted() {
      document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
      this.listenRouteForChangeActiveButton();
      window.addEventListener('resize',this.onresize);
      this.getConfig();
    },
    beforeDestroy() {
      document.getElementById('router_wrap').removeEventListener('click', this.hideFeature);
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
      transactionsSelect(flShowTransactionsSelect){
        this.flShowValidatorsSelect = false;
        if(!flShowTransactionsSelect){
          this.flShowTransactionsSelect = true;
          this.flShowUpOrDown = true
        }else {
          this.flShowUpOrDown = false;
          this.flShowTransactionsSelect = false
        }
      },
      netWorkSelect(flShowNetworkSelect){
        this.flShowNetworkSelect = false;
        if(!flShowNetworkSelect){
          this.flShowNetworkSelect = true;
          this.flShowNetworkUpOrDown = true
        }else {
          this.flShowNetworkSelect = false;
          this.flShowNetworkUpOrDown = false
        }
      },
      hideFeature() {
        if (this.featureShow) {
          this.featureShow = false;
        }
      },
      onresize(){
        this.innerWidth = window.innerWidth;
        if(window.innerWidth > 910){
          this.devicesShow = 1;
          this.appHeaderVar = 'person_computer_header_var';
        }else {
          this.devicesShow = 0;
          this.appHeaderVar = 'mobile_header_var';
        }
      },
      showFeature() {
        this.featureShow = !this.featureShow;
      },
      featureButtonClick(path, isLogoClick) {
        if (!isLogoClick) {
          this.featureShow = !this.featureShow;
        }
        this.showSubTransaction = false;
        this.showSubValidators = false;
        this.listenRouteForChangeActiveButton();
        if(path !== 'network'){
          this.$router.push(path);
        }
      },
      transactionMouseOver(){
        this.showSubTransaction = true;
      },
      transactionMouseLeave(){
        this.showSubTransaction = false;
      },
      searchTx(){
        let uri = `/api/tx/${this.searchInputValue}`;
        Service.http(uri).then((tx) => {
          if (tx) {
            this.$router.push(`/tx?txHash=${tx.Hash}`);
            this.clearSearchInputValue();
          }else {
            this.toSearchResultPage();
          }
        }).catch(e => {
          this.toSearchResultPage();
          console.error(e)
        });
      },
      searchDelegator(){
        let uri = `/api/account/${this.searchInputValue}`;
        Service.http(uri).then((delegatorAddress) => {
          if (delegatorAddress) {
            this.$router.push(`/address/1/${delegatorAddress.address}`);
            this.clearSearchInputValue();
          }else {
            this.toSearchResultPage()
          }
        }).catch(e => {
          this.toSearchResultPage();
          console.error(e)
        });
      },
      searchValidator(){
        let uri = `/api/stake/candidate/${this.searchInputValue}`;
        Service.http(uri).then((validatorAddress) => {
          if (validatorAddress) {
            this.$router.push(`/address/1/${validatorAddress.validator.address}`);
            this.clearSearchInputValue();
          }else {
            this.toSearchResultPage()
          }
        }).catch(e => {
          this.toSearchResultPage();
          console.error(e)
        });
      },
      searchBlockAndProposal(){
        let uri = `/api/search/${this.searchInputValue}`;
        Service.http(uri).then((searchResult) => {
          if(searchResult){
            //searchResult：[ {Type：block，Data:{}} ，{Type：proposal,Data:{}} ]
            let searchResultIsBlockOrProposalId = 1;
            let searchBlockAndProposalInResult = 2;
            if(searchResult.length === searchResultIsBlockOrProposalId){
              if(searchResult[0].Type === "block" && searchResult[0].Data.Height !== 0){
                this.$router.push(`/blocks_detail/${searchResult[0].Data.Height}`);
                this.clearSearchInputValue();
              }else if(searchResult[0].Type === "proposal" && searchResult[0].Data.ProposalID !== 0){
                this.$router.push(`/ProposalsDetail/${searchResult[0].Data.ProposalID}`);
                this.clearSearchInputValue();
              }
            }else if(searchResult.length === searchBlockAndProposalInResult){
              this.toSearchResultPage();
            }
          }else {
            this.toSearchResultPage();
          }
        }).catch((e) => {
          console.error(e);
          this.toSearchResultPage();
        });
      },
      getData() {
        if(Tools.removeAllSpace(this.searchInputValue) === ''){
          this.clearSearchContent();
          return
        }else{
          if(/^[A-F0-9]{64}$/.test(this.searchInputValue)){
            this.searchTx()
          }else if(this.$Codec.Bech32.isBech32(this.$Crypto.config.iris.bech32.accAddr,this.searchInputValue) ) {
            this.searchDelegator();
          }else if(this.$Codec.Bech32.isBech32(this.$Crypto.config.iris.bech32.valAddr,this.searchInputValue)){
            this.searchValidator();
          }else if (/^\+?[1-9][0-9]*$/.test(this.searchInputValue)){
            this.searchBlockAndProposal();
          }else {
            this.toSearchResultPage();
          }
        }
      },
      toSearchResultPage(){
        this.$router.push(`/searchResult/${this.searchInputValue}`);
        this.searchInputValue = "";
      },
      onInputChange() {
        this.getData();
      },
      clearSearchInputValue(){
        this.searchInputValue = "";
      },
      listenRouteForChangeActiveButton(){
        //刷新的时候路由不变，active按钮不变
        let path = window.location.href;
        if (path.includes('transactions/2') || path.includes('tx?')) {
          this.activeClassName = '/transaction';
        } else if (path.includes('/validators')) {
          this.activeClassName = '/validators';
        } else if (path.includes('/block')) {
          this.activeClassName = '/block';
        } else if (path.includes('/home')) {
          this.activeClassName = '/home';
        } else if (path.includes('/faucet')) {
          this.activeClassName = '/faucet';
        } else if(path.includes('/Proposals')){
          this.activeClassName = '/Proposals';
        }else if(path.includes('/nodespage')){
          this.activeClassName = '/nodespage';
        }else {
          this.activeClassName = '';
        }
      },
      clearSearchContent(){
        this.searchInputValue = '';
      },
      toHelp(){
        this.$router.push('/help')
      },
      getConfig(){
        Service.http('/api/config').then(res => {
          this.toggleTestnetLogo(res);
          this.explorerLink(res);
          this.setCurrentSelectOption(res.cur_env);
          this.setEnvConfig(res);
          this.flShowHeaderNetwork = true;
          res.configs.forEach( item => {
            if(res.cur_env === item.env_nm){
              this.chainId = `${item.chain_id.toUpperCase()} ${item.env_nm.toUpperCase()}`
              if(item.show_faucet && item.show_faucet === 1){
                this.flShowFaucet = true
              }else {
                this.flShowFaucet = false
              }
            }
          })
        });
      },
      setEnvConfig(currentEnv){
        if(currentEnv.cur_env !== constant.ENVCONFIG.MAINNET){
          this.$Crypto.getCrypto(constant.CHAINNAME,constant.ENVCONFIG.TESTNET);
          this.$store.commit('currentEnv',constant.ENVCONFIG.TESTNET)
        }
      },
      toggleTestnetLogo(currentEnv){
        if(currentEnv.cur_env === constant.ENVCONFIG.MAINNET){
          this.flShowChainId = false;
        }else {
          this.flShowChainId = true;
        }
      },
      explorerLink(envLink){
        envLink.configs.forEach( item => {
          if(item.env_nm === constant.ENVCONFIG.MAINNET){
            this.mainnetHref = item.host
          }else if(item.env_nm === constant.ENVCONFIG.TESTNET){
            this.testnetHref = item.host
          }else if(item.env_nm  === "NYAN_CATS"){
           this.nyanCatsHref = item.host
          }
        });
      },
      setCurrentSelectOption(currentEnv){
        if(currentEnv === constant.ENVCONFIG.MAINNET){
          this.currentSelected = lang.home.mainnet;
          this.flShowTestnet = true;
        }else if(currentEnv === constant.ENVCONFIG.TESTNET){
          this.currentSelected = lang.home.testnet;
          this.flShowMainnet = true;
        }else {
          this.flShowMainnet = false;
          this.flShowTestnet = true;
          this.currentSelected = lang.home.mainnet;
        }
      },
      toMainnet(MainnetHref){
        window.open(MainnetHref)
      },
      toTestnet(testnetHref){
        window.open(testnetHref)
      },
      toNyanCats(nyanCatsHref){
        window.open(nyanCatsHref)
      },
      toggleSelectOption(){
        this.$store.commit('flShowSelectOption',!this.$store.state.flShowSelectOption)
      }
    },
    updated(){
      this.absoluteTop = `${this.$refs.header_content.clientHeight/100}rem`
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';
  .person_computer_header_var {
    height: 1.62rem;
    position: fixed;
    z-index: 10001;
    background: rgba(255,255,255,1);
  }
  .person_computer_header_var, .mobile_header_var {
    @include flex();
    @include pcContainer;
    .app_header_person_computer {
      @include pcCenter;
      justify-content: space-between;
      padding-top: 0.01rem;
      width: 100%;
      .header_top {
        padding: 0 0.2rem;
        height: 0.96rem;
        @include flex;
        justify-content: space-between;
        .header_left{
          @include flex;
          .imageWrap{
            @include flex;
            cursor:pointer;
            margin-top:0.2rem;
            width: 1.7rem;
            height: 0.5rem;
            img{
              width: 100%;
            }
            .chain_content{
              font-size: 0.18rem;
              white-space: nowrap;
              display: flex;
              height: 0.36rem;
              margin-top: 0.09rem;
              line-height: 0.36rem;
              align-items: center;
              margin-left: 0.2rem;
              padding: 0 0.14rem;
              color: #F2711D;
              border: 0.01rem solid #FFD3B6;
              border-radius: 0.04rem;
              background: #FFF9F5;
            }
            .logo_title_wrap{
              margin-left:0.16rem;
              @include flex;
              flex-direction:column;
              justify-content: center;
              .logo_title_content{
                height:0.27rem;
                line-height:0.27rem;
                span{
                  &:first-child{
                    font-size:0.24rem;
                    color:#005a98;
                  }
                  &:last-child{
                    font-size:0.24rem;
                    color:#3598db;
                  }

                }
              }
              p{
                font-size:0.14rem;
                height:0.16rem;
                line-height:0.16rem;
                margin-top:0.04rem;
                span{
                  color:#a2a2ae;
                }
              }

            }
          }
        }
        .chain_id_container{
          display: flex;
          align-items: center;
          margin-left: 0.15rem;
          span{
            font-size: 0.16rem;
            height: 0.36rem;
            line-height: 0.36rem;
            border-radius: 0.06rem;
            padding: 0 0.13rem;
            background: rgba(234, 248, 254, 1);
            color: rgba(34, 37, 42, 1);
          }
        }
        .network_select_container{
          display: flex;
          align-items: center;
          margin-left: 0.1rem;
          .network_select_option{
            height: 0.36rem;
            position: relative;
            line-height: 0.36rem;
            text-align: center;
            z-index: 1000;
            border-radius: 0.06rem;
            border: 0.01rem solid rgba(215, 217, 224, 1);
            .select_content{
              width: 1.9rem;
              cursor: pointer;
              .current_select_content{
                color: rgba(53, 152, 219, 1);
                display: flex;
                justify-content: space-between;
                box-sizing: border-box;
                padding: 0 0.08rem;
                i{
                  color: rgba(53, 152, 219, 1);
                  font-size: 0.16rem;
                }
              }
              .select_option_content{
                background: rgba(255,255,255,1);
                color: rgba(102, 102, 102, 1);
                border-radius: 0.06rem;
                overflow: hidden;
                margin-top: 0.04rem;
                box-shadow:0 0 0.04rem 0 rgba(0,20,33,0.12);
                user-select: none;
              }
              .common_option{
                padding-left: 0.08rem;
                text-align: left;
                &:hover{
                  color: rgba(53, 152, 219, 1);
                  background: rgba(234, 248, 254, 1);
                }
              }
            }
          }
        }
        .navSearch {
          flex: 1;
          display: flex;
          justify-content: flex-end;
          margin-bottom: 0.01rem;
          position: relative;
          margin-left: 0.2rem;
          input::-webkit-input-placeholder {
            text-align: left;
            font-size: 0.14rem;
            color: #cccccc;
          }

          .search_input {
            @include inputBoxShadow;
            @include borderRadius(0.06rem);
            flex: 1;
            width: 100%;
            max-width: 4.5rem;
            height: 0.36rem;
            text-indent: 0.1rem;
            outline: none;
            border: 0.01rem solid #dddddd;
            margin-top: 0.3rem;
            font-size: 0.14rem;
            padding-right:0.48rem;
            &::-webkit-input-placeholder{
              font-size:0.14rem;
              color:#AEAEB9;
            }
            &:-moz-placeholder {
              font-size:0.14rem;
              color:#AEAEB9;
            }
            &::-moz-placeholder {
              font-size:0.14rem;
              color:#AEAEB9;
            }
            &:-ms-input-placeholder {
              font-size:0.14rem;
              color:#AEAEB9;
            }
          }
          .search_icon {
            position: absolute;
            top: 0.4rem;
            right: 0.1rem;
            width: 0.15rem;
            height: 0.15rem;
            background: url('../assets/search.svg') no-repeat;
            cursor: pointer;
          }
          .clear_icon{
            position: absolute;
            top: 0.42rem;
            right: 0.3rem;
            width: 0.15rem;
            height: 0.15rem;
            background: url('../assets/clear.png') no-repeat;
            cursor: pointer;
          }
        }
      }

    }
    .useFeature {
      width: 100%;
      height:0.66rem;
      @include flex;
      flex-direction: column;
      align-items: center;
      background: #3598db;
      .navButton {
        width: 100%!important;
        padding: 0 0.2rem;
        height:0.66rem;
        @include pcCenter;
        @include flex;
        .common_item_style{
          &:hover{
            color: #ffffff;
            background: #005a98;
          }
        }
        .nav_item {
          display: inline-block;
          height: 0.65rem;
          line-height: 0.66rem;
          width:1.6rem;
          text-align: center;
          font-size: 0.18rem;
          cursor: pointer;
          color: #c9eafd;
          font-weight: 500!important;
          .bottom_arrow{
            display:inline-block;
            height:0.11rem;
            width:0.11rem;
            background: url('../assets/caret-bottom.png') no-repeat 0 0;
            top:0.27rem;
            right:0.1rem;
          }
        }
        .sub_btn_wrap{
          @include flex;
          flex-direction:column;
          z-index:100;
          padding:0;
          .sub_btn_item{
            height:0.36rem;
            line-height:0.36rem;
            font-size:0.14rem;
            background: #005a98;
            color: #c9eafd;
            width:1.6rem;
            text-align: left;
            padding-left:0.2rem;
            a{
              width: 100%;
              display: inline-block;
              padding-left: 0.19rem;
              color: #c9eafd!important;
              &:hover{
                color: #00f0ff!important;
              }
            }
            &:hover{
              color: #00f0ff;
            }
          }
          .validators_btn_item{
            padding-left:0.35rem;
          }
        }
        .nav_item_active {
          color: #ffffff;
          background: #005a98;
        }
        .btn-group, .btn-group-vertical{
          vertical-align: baseline;
        }
      }

    }
    .app_header_mobile {
      width: 100%;
      padding: 0.15rem 0 0 0;
      @include flex();
      flex-direction: column;
      border-bottom: 0.01rem solid #cccccc;
      position: relative;
      .search_input {
        @include inputBoxShadow;
      }
      .feature_btn {
        width: 0.25rem;
        height: 0.25rem;
        top: 0.26rem;
        right: 0.1rem;
        img{
          width: 100%;
        }
        /*background: url('../assets/menu.png') no-repeat;*/
      }
      .image_wrap_mobile {
        @include flex;
        align-items: center;
        width: 1.5rem;
        height: 0.5rem;
        img{
          width: 100%;
          height: 100%;
        }
        .logo_title_wrap{
          margin-left:0.16rem;
          @include flex;
          flex-direction:column;
          justify-content: center;
          .logo_title_content{
            height:0.27rem;
            line-height:0.27rem;
            span{
              &:first-child{
                font-size:0.24rem;
                color:#005a98;
              }
              &:last-child{
                font-size:0.24rem;
                color:#3598db;
              }

            }
          }
          p{
            font-size:0.14rem;
            span{
              color:#aeaeb9;
            }
          }

        }
      }
      .search_input_mobile {
        margin: 0.1rem 0.15rem;
        @include flex();
        flex-direction: column;
        align-items: center;
        position: relative;
        input::-webkit-input-placeholder {
          text-align: center;
          font-size: 0.14rem;
          color: #cccccc;
        }
        input {
          width: 100%;
          @include borderRadius(0.06rem);
          border: 0.01rem solid #eee;
          font-size: 0.14rem;
          padding:0 0.48rem 0 0.1rem;
          height:0.3rem;
          &:focus {
            border: 0.01rem solid #3190e8;
            outline: none;
          }
        }
        .search_icon {
          position: absolute;
          top: 0.07rem;
          right: 0.15rem;
          width: 0.15rem;
          height: 0.15rem;
          background: url('../assets/search.svg') no-repeat;
          cursor: pointer;
        }
        .clear_icon{
          position: absolute;
          top: 0.08rem;
          right: 0.32rem;
          width: 0.15rem;
          height: 0.15rem;
          background: url('../assets/clear.png') no-repeat;
          cursor: pointer;
        }
      }
      .mobile_chain_id_content{
        height: 0.3rem;
        margin:  0 0.15rem;
        background: rgba(234, 248, 254, 1);
        display: flex;
        justify-content: flex-start;
        line-height: 0.3rem;
        font-size: 0.14rem;
        margin-bottom: 0.1rem;
        border-radius: 0.06rem;
        .mobile_chain_content{
          padding-left: 0.1rem;
        }
      }
      .use_feature_mobile {
        z-index: 1010;
        width: 100%;
        left: 0;
        background: #3598db;
        @include flex();
        flex-direction: column;
        position: absolute;
        .select_option {
          display: flex;
          flex-direction: column;
          .feature_btn_mobile {
            height: 0.39rem;
            line-height: 0.39rem;
            padding-left: 0.15rem;
            background: #005a98;
            color: #fff;
            font-size: 0.14rem;
          }
        }

        .feature_btn_mobile {
          height: 0.39rem;
          line-height: 0.39rem;
          padding-left: 0.15rem;
          background: #3598db;
          color: #fff;
          font-size:0.14rem;
          a{
            display: inline-block;
            width: 100%;
            color: #fff!important;
          }
        }
        .feature_arrow {
          position: relative;
          color: #c9eafd;
          font-size:0.14rem;
          background: url('../assets/caret-bottom.png') no-repeat 97% 0.12rem,#3598db;
        }
        .feature_subNav {
          padding-left: 0.3rem;
        }
      }
    }
  }
  .chain_id{
    padding-right: 0.26rem;
    font-size: 0.16rem;
    color: #F2711C;
    @include fontWeight;
  }
  .select_option_container{
    display: flex;
    padding-right: 0.2rem;
    justify-content: space-between;
    .upImg_content{
      width: 0.28rem;
      padding-left: 0.2rem;
      img{
        width: 100%;
      }
    }
    .downImg_content{
      width: 0.28rem;
      padding-left: 0.2rem;
      img{
        width: 100%;
      }
    }
  }
</style>
