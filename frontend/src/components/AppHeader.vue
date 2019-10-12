<!--头部页面-->
<template>
    <div :class="appHeaderVar"
         v-show="showHeader"
         id="header">
        <header class="app_header_person_computer"
                v-show="devicesShow === 1">
            <div class="header_navigation_container">
                <div class="header_navigation_wrap">
                    <div class="header_left_content">
                        <div class="header_logo_content">
                            <router-link :to="`/home`"><img src="../assets/logo.png"></router-link>
                        </div>
                        <ul class="header_menu_content">
                            <li class="header_menu_item"
                                :class="activeBlockChain ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('blockChain')" @mouseleave="hideTwoMenu('blockChain')">Blockchain</li>
                            <li class="header_menu_item"
                                :class="activeStaking ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('staking')" @mouseleave="hideTwoMenu('staking')">Staking</li>
                            <li class="header_menu_item"
                                :class="activeTransfers ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('transfers')" @mouseleave="hideTwoMenu('transfers')">
                                <router-link :to="`/txs/transfers`">Transfer</router-link></li>
                            <li class="header_menu_item"
                                :class="activeAssets ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('assets')" @mouseleave="hideTwoMenu('assets')">Asset</li>
                            <li class="header_menu_item"
                                :class="activeGov ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('gov')" @mouseleave="hideTwoMenu('gov')">Gov</li>
                            <li class="header_menu_item"
                                :class="activeStats ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('stats')" @mouseleave="hideTwoMenu('stats')">Stats</li>
                            <li v-show="flShowFaucet" class="header_menu_item"
                                :class="activeFaucet ? 'nav_item_active' : ''"
                                @mouseenter="showTwoMenu('faucet')" @mouseleave="hideTwoMenu('faucet')"
                            ><router-link :to="`/faucet`">Faucet</router-link></li>
                        </ul>
                    </div>

                    <div class="header_right_content">
                        <div class="search_input_container">
                            <div class="search_input_wrap">
                                <input type="text"
                                       class="search_input"
                                       placeholder="Search by Address / Txhash / Block"
                                       v-model.trim="searchInputValue"
                                       @keyup.enter="onInputChange">
                                <span @click="getData(searchInputValue)" class="iconfont iconsousuo"></span>
                            </div>
                        </div>
                        <div class="network_container" @mouseenter="showNetWorkLogo()" @mouseleave="hideNetWorkLogo()">
                            <span style="color: #fff">
                                <i style="font-size: 0.24rem;padding-right: 0.02rem;" :class="currentNetworkClass"></i>
                                <i style="font-size: 0.08rem" class="iconfont iconwangluoqiehuanjiantou"></i>
                            </span>
                            <ul class="network_list_container" v-show="flShowNetworkLogo" @mouseenter="showNetWorkLogo()" @mouseleave="hideNetWorkLogo()">
                                <li class="network_list_item"
                                    v-for="item in netWorkArray"
                                    @click="windowOpenUrl(item.host)"><i :class="item.icon"></i>{{item.netWorkSelectOption}}</li>
                            </ul>
                        </div>
                    </div>

                </div>
            </div>
            <transition name="fade">
                <div class="header_submenu_container" @mouseenter="showTwoMenu(menuActiveName)" @mouseleave="hideTwoMenu(menuActiveName)">
                    <div class="header_submenu_content_wrap">
                        <ul class="header_submenu_content"  :style="{'left':offSetLeft,width:contentWidth}">
                            <li class="header_submenu_item" v-show="flShowChain"><router-link :to="`/home`">Overview</router-link></li>
                            <li class="header_submenu_item" v-show="flShowChain"><router-link :to="`/blocks`">Blocks</router-link></li>
                            <li class="header_submenu_item" v-show="flShowChain"><router-link :to="`/txs`">Transactions</router-link></li>
                            <!--<li class="header_submenu_item no_border_style" v-if="flShowChain"><router-link :to="`/validators`">Validators</router-link></li>-->
                            <!--               <li class="header_submenu_item" v-if="flShowChain">Assets</li>
                                           <li class="header_submenu_item" v-if="flShowChain">Gateways</li>-->
                            <li class="header_submenu_item" v-show="flShowStaking"><router-link :to="`/validators`">Validators</router-link></li>
                            <li class="header_submenu_item" v-if="flShowStaking"><router-link :to="`/txs/delegations`">Delegation Txs</router-link></li>
                            <li class="header_submenu_item" v-show="flShowStaking"><router-link :to="`/txs/validations`">Validation Txs</router-link></li>
                            <!--  <li class="header_submenu_item" v-if="flShowStaking">Validator Txs</li>-->
                            <!--<li class="header_submenu_item" v-if="flShowStaking"> <router-link :to="`/txs/stakes`">Delegation Txs</router-link></li>
                            <li class="header_submenu_item" v-if="flShowStaking">Reward Txs</li>-->
                            <!--<li class="header_submenu_item no_border_style" v-show="flShowTransfers"> <router-link :to="`/txs/transfers`">IRIS Transfers Txs</router-link></li>-->
                            <!--<li class="header_submenu_item" v-if="flShowTransfers">Asset Transfers Txs</li>-->
                            <!--<li class="header_submenu_item" v-if="flShowTransfers">Inter-chain Txs</li>-->
                            <!--<li class="header_submenu_item" v-if="flShowTransfers">IRIS Burn Txs</li>-->
                            <li class="header_submenu_item" v-show="flShowAssets"><router-link :to="`/assets/ntvassets`">Native Asset </router-link></li>
                            <li class="header_submenu_item" v-show="flShowAssets"><router-link :to="`/assets/ntvassetstxs`">Native Asset Txs</router-link></li>
                            <li class="header_submenu_item" v-show="flShowAssets && flShowGatewayMenu"><router-link :to="`/assets/gtwassets`">Gateway Asset</router-link></li>
                            <li class="header_submenu_item no_border_style" v-show="flShowAssets && flShowGatewayMenu"><router-link :to="`/assets/gtwassetstxs`">Gateway Asset Txs</router-link></li>
                            <!--<li class="header_submenu_item" v-if="flShowAssets">Assets Transfers</li>-->
                            <li class="header_submenu_item" v-show="flShowGov"><router-link :to="`/gov/parameters`">Parameters</router-link></li>
                            <li class="header_submenu_item" v-show="flShowGov"><router-link :to="`/gov/proposals`">Proposals</router-link></li>
                            <li class="header_submenu_item no_border_style" v-show="flShowGov"><router-link :to="`/txs/governance`">Gov Txs</router-link></li>
                            <!--<li class="header_submenu_item" v-if="flShowGov">Vote Tx</li>-->
                            <li class="header_submenu_item" v-show="flShowStats"><router-link :to="`/stats/irisrichlist`">IRIS Rich List</router-link></li>
                            <li class="header_submenu_item no_border_style" v-show="flShowStats"><router-link :to="`/stats/irisstats`">IRIS Stats</router-link></li>
                            <!--<li class="header_submenu_item" v-if="flShowStats">Public Address</li>-->
                        </ul>
                    </div>
                </div>
            </transition>
        </header>
        <div class="app_header_mobile"
             v-show="devicesShow === 0"
             ref="header_content">
            <div style="display: flex;flex-direction: row-reverse;justify-content: space-between;align-items: center;padding: 0 0.15rem;">
                <div class="feature_btn"
                     @click="showFeature">
                    <img src="../assets/menu.png">
                </div>
                <div class="image_wrap_mobile"
                     @click="featureButtonClick('/home',true)" :style="{visibility: flShowLogo ? 'visible' : 'hidden'}">
                    <img :src="explorerLogo" />
                    <i :class="currentNetworkClass"></i>
                </div>
            </div>
            <div class="search_input_mobile">
                <div style="width:100%;display: flex;justify-content: center;">
                    <input type="text"
                           class="search_input"
                           v-model.trim="searchInputValue"
                           @keyup.enter="onInputChange"
                           placeholder="Search by Address / Txhash / Block">
                    <i class="search_icon"
                       @click="getData(searchInputValue)"></i>
                    <i class="clear_icon"
                       @click="clearSearchContent"
                       v-show="showClear"></i>
                </div>
            </div>
            <div class="use_feature_mobile"
                 :style="{'top':absoluteTop}"
                 v-show="featureShow">
                <div class="mobile_menu_container" @click="flShowBlockchain('blockChain')">
                    <div class="mobile_menu_item_content">
                        <span>Blockchain</span>
                        <i class="iconfont iconwangluoqiehuanjiantou" :class="flShowBlockchainMenu ? 'up_style' : 'down_style'"> </i>
                    </div>
                    <ul class="blockchain_list_content" v-show="flShowBlockchainMenu">
                        <li class="blockchain_list_item" @click="featureButtonClick(`/home`)">Overview</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/blocks`)">Blocks</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/txs`)">Transactions</li>
                    </ul>
                </div>

                <div class="mobile_menu_container" @click="flShowBlockchain('staking')">
                    <div class="mobile_menu_item_content">
                        <span>Staking</span>
                        <i class="iconfont iconwangluoqiehuanjiantou" :class="flShowStakingMenu ? 'up_style' : 'down_style'"> </i>
                    </div>
                    <ul class="blockchain_list_content" v-show="flShowStakingMenu">
                        <li class="blockchain_list_item" @click="featureButtonClick(`/validators`)">Validators</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/txs/delegations`)">Delegation Txs</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/txs/validations`)">Validation Txs</li>
                    </ul>
                </div>

                <div class="mobile_menu_container" @click="flShowBlockchain('transfers')">
                    <div class="mobile_menu_item_content">
                        <span  @click="featureButtonClick(`/txs/transfers`)">Transfers</span>
                        <!--<i class="iconfont iconwangluoqiehuanjiantou" :class="flShowTransfersMenu ? 'up_style' : 'down_style'"> </i>-->
                    </div>
                    <!--<ul class="blockchain_list_content" v-show="flShowTransfersMenu">-->
                        <!--<li class="blockchain_list_item" @click="featureButtonClick(`/txs/transfers`)">IRIS Transfers Txs</li>-->
                    <!--</ul>-->
                </div>

                <div class="mobile_menu_container" @click="flShowBlockchain('assets')">
                    <div class="mobile_menu_item_content">
                        <span>Assets</span>
                        <i class="iconfont iconwangluoqiehuanjiantou" :class="flShowAssetsMenu ? 'up_style' : 'down_style'"> </i>
                    </div>
                    <ul class="blockchain_list_content" v-show="flShowAssetsMenu">
                        <li class="blockchain_list_item" @click="featureButtonClick(`/assets/ntvassets`)">Native Asset</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/assets/ntvassetstxs`)">Native Asset Txs</li>
                        <li class="blockchain_list_item" v-if="flShowGatewayMenu" @click="featureButtonClick(`/assets/gtwassets`)">Gateway Asset</li>
                        <li class="blockchain_list_item" v-if="flShowGatewayMenu" @click="featureButtonClick(`/assets/gtwassetstxs`)">Gateway Asset Txs</li>
                    </ul>
                </div>

                <div class="mobile_menu_container" @click="flShowBlockchain('gov')">
                    <div class="mobile_menu_item_content">
                        <span>Gov</span>
                        <i class="iconfont iconwangluoqiehuanjiantou" :class="flShowGovMenu ? 'up_style' : 'down_style'"> </i>
                    </div>
                    <ul class="blockchain_list_content" v-show="flShowGovMenu">
                        <li class="blockchain_list_item" @click="featureButtonClick(`/gov/parameters`)">Parameters</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/gov/proposals`)">Proposals</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/txs/governance`)">Gov Txs</li>
                    </ul>
                </div>


                <div class="mobile_menu_container" @click="flShowBlockchain('stats')">
                    <div class="mobile_menu_item_content">
                        <span>Stats</span>
                        <i class="iconfont iconwangluoqiehuanjiantou" :class="flShowStatsMenu ? 'up_style' : 'down_style'"> </i>
                    </div>
                    <ul class="blockchain_list_content" v-show="flShowStatsMenu">
                        <li class="blockchain_list_item" @click="featureButtonClick(`/stats/irisrichlist`)">IRIS Rich List</li>
                        <li class="blockchain_list_item" @click="featureButtonClick(`/stats/irisstats`)">IRIS Stats</li>
                    </ul>
                </div>

                <div class="mobile_menu_container" v-if="flShowFaucet">
                    <div class="mobile_menu_item_content">
                        <span @click="featureButtonClick(`/faucet`)">Faucet</span>
                    </div>
                </div>

                <div class="mobile_menu_container" @click="flShowBlockchain('network')">
                    <div class="mobile_menu_item_content">
                        <span>Network</span>
                        <i class="iconfont iconwangluoqiehuanjiantou" :class="flShowNetWorkMenu ? 'up_style' : 'down_style'"> </i>
                    </div>
                    <ul class="blockchain_list_content" v-show="flShowNetWorkMenu">
                        <li class="blockchain_list_item"  v-for="item in netWorkArray">
                            <a :href="item.host">{{item.netWorkSelectOption}}</a> <i :class="item.icon"></i>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
	import Tools from '../util/Tools';
	import Service from "../service";
	import constant from '../constant/Constant';
	import lang from "../lang/index";
	import skinStyle from '../skinStyle'
	export default {
		name: 'app-header',
		watch: {
			$route () {
				this.searchInputValue = "";
				this.flShowSubMenu = false;
				this.listenRouteForChangeActiveButton();
				this.showHeader = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
			},
			searchInputValue (searchInputValue) {
				if (searchInputValue) {
					this.showClear = true;
				} else {
					this.showClear = false;
				}
			},
		},
		data () {
			return {
				devicesWidth: window.innerWidth,
				devicesShow: 1,
				searchValue: '',
				appHeaderVar: 'person_computer_header_var',
				featureShow: false,
				transactionShow: false,
				validatorsShow: false,
				flShowGovernanceOption: false,
				flShowStatistics: false,
				searchInputValue: '',
				activeClassName: '/home',
				showHeader: !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
				flShowFaucet: false,
				showSubTransaction: false,
				showSubValidators: false,
				showClear: false,
				innerWidth: window.innerWidth,
				flShowTransactionsSelect: false,
				flShowAssetSelect: false,
				flShowValidatorsSelect: false,
				flShowNetworkSelect: false,
				flShowGovernanceSelect: false,
				flShowTopListSelection: false,
				flShowUpOrDown: false,
				flShowNetwork: false,
				flShowHeaderNetwork: false,
				flShowValidatorsUpOrDown: false,
				flShowNetworkUpOrDown: false,
				currentSelected: '',
				absoluteTop: '',
				lang: lang,
				chainId: '',
				upImg: require("../assets/caret-bottom.png"),
				downImg: require("../assets/caret-bottom.png"),
				netWorkArray: [],
				explorerLogo: '',
				flShowLogo: false,
				showAssetTransaction: false,
				flShowSubMenu: false,
                flShowChain: false,
                flShowStaking: false,
                flShowTransfers: false,
                flShowAssets: false,
                flShowGov: false,
                flShowStats: false,
                flShowNetworkLogo: false,
				flShowBlockchainMenu:false,
				flShowStakingMenu:false,
				flShowTransfersMenu:false,
				flShowAssetsMenu:false,
				flShowGovMenu:false,
				flShowStatsMenu:false,
				flShowNetWorkMenu:false,
				activeBlockChain:false,
				activeStaking:false,
				activeTransfers:false,
				activeAssets:false,
				activeGov:false,
				activeStats:false,
				activeFaucet:false,
                hoverBlockChainTag:false,
                menuActiveName: '',
				currentNetworkClass:'',
				offSetLeft:'1.6rem',
                contentWidth:'',
                flShowGatewayMenu:false
			}
		},
		beforeMount () {
			if (window.innerWidth > 910) {
				this.devicesShow = 1;
				this.appHeaderVar = 'person_computer_header_var';
			} else {
				this.devicesShow = 0;
				this.appHeaderVar = 'mobile_header_var';
			}
		},
		mounted () {
			document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
			this.listenRouteForChangeActiveButton();
			window.addEventListener('resize', this.onresize);
			this.getConfig();
			if(localStorage.getItem('currentEnv')){
				this.toggleTestnetLogo({
					cur_env:localStorage.getItem('currentEnv')
				});
            }
			this.listenRouteForChangeActiveButton();
		},
		beforeDestroy () {
			document.getElementById('router_wrap').removeEventListener('click', this.hideFeature);
			window.removeEventListener('resize', this.onWindowResize);
		},
		methods: {
			flShowBlockchain(v){
				switch (v) {
					case 'blockChain' :
						this.flShowBlockchainMenu = !this.flShowBlockchainMenu;
						break;
					case 'staking' :
						this.flShowStakingMenu = !this.flShowStakingMenu;
						break;
					case 'transfers' :
						this.flShowTransfersMenu = !this.flShowTransfersMenu;
						break;
					case 'assets'	:
						this.flShowAssetsMenu = !this.flShowAssetsMenu;
						break;
					case 'gov' :
						this.flShowGovMenu = !this.flShowGovMenu;
						break;
					case 'stats' :
						this.flShowStatsMenu = !this.flShowStatsMenu;
						this.flShowStats = true;
						break;
					case 'network' :
						this.flShowNetWorkMenu = !this.flShowNetWorkMenu;
						this.flShowStats = true;
						break;
					case 'faucet' :
						this.activeFaucet = true;
						break;
				}
            },
			showTwoMenu(v){
				this.flShowSubMenu = true;
				this.menuActiveName = v;
				this.hideActiveStyle();
				switch (v) {
                    case 'blockChain' :
	                    this.offSetLeft = `1.7rem`;
	                    this.contentWidth = '1.15rem';
	                    this.flShowChain = true;
	                    this.hoverBlockChainTag = false;
	                    this.activeBlockChain  = true;
                    	break;
                    case 'staking' :
                    	this.offSetLeft = `2.575rem`;
	                    this.flShowStaking = true;
	                    this.contentWidth = '1.25rem';
	                    this.hoverBlockChainTag = true;
	                    this.activeStaking  = true;
                    	break;
                    case 'transfers' :
	                    this.offSetLeft = `3.245rem`;
	                    this.contentWidth = '1.47rem';
	                    this.flShowTransfers = true;
	                    this.activeTransfers  = true;
                    	break;
                    case 'assets'	:
	                    this.offSetLeft = `3.96rem`;
	                    this.contentWidth = '1.55rem';
	                    this.flShowAssets = true;
	                    this.activeAssets = true;
                    	break;
                    case 'gov' :
	                    this.offSetLeft = `4.51rem`;
	                    this.contentWidth = '1.03rem';
	                    this.flShowGov = true;
	                    this.activeGov = true;
                    	break;
                    case 'stats' :
	                    this.offSetLeft = `4.965rem`;
	                    this.contentWidth = '1.15rem';
	                    this.flShowStats = true;
	                    this.activeStats  = true
				}
            },
			hideTwoMenu(v){
				this.flShowSubMenu = false;
				this.menuActiveName = v;
				switch (v) {
					case 'blockChain' :
						this.flShowChain = false;
						break;
					case 'staking' :
						this.flShowStaking = false;
						break;
					case 'transfers' :
						this.flShowTransfers = false;
						break;
					case 'assets'	:
						this.flShowAssets = false;
						break;
					case 'gov' :
						this.flShowGov = false;
						break;
					case 'stats' :
						this.flShowStats = false;

				}
            },
			hideFeature () {
				if (this.featureShow) {
					this.featureShow = false;
				}
			},
			onresize () {
				this.innerWidth = window.innerWidth;
				if (window.innerWidth > 910) {
					this.devicesShow = 1;
					this.appHeaderVar = 'person_computer_header_var';
				} else {
					this.devicesShow = 0;
					this.appHeaderVar = 'mobile_header_var';
				}
			},
			showFeature () {
				this.featureShow = !this.featureShow;
			},
			featureButtonClick (path, isLogoClick) {
				if (!isLogoClick) {
					this.featureShow = !this.featureShow;
				}
				this.showSubTransaction = false;
				this.showSubValidators = false;
				this.flShowGovernanceOption = false;
				this.flShowStatistics = false;
				this.listenRouteForChangeActiveButton();
				if (path !== 'network') {
					this.$router.push(path);
				}
			},
			searchTx () {
				Service.commonInterface({headerTx:{searchValue: this.searchInputValue}}, (tx) => {
					try {
						if (tx) {
							this.$router.push(`/tx?txHash=${tx.hash}`);
							this.clearSearchInputValue();
						} else {
							this.toSearchResultPage();
						}
					}catch (e) {
						console.error(e);
						this.toSearchResultPage();
					}
                });
			},
			searchDelegator () {
				Service.commonInterface({headerSearchAccount:{searchValue:this.searchInputValue}}, (delegatorAddress) => {
					try {
						if (delegatorAddress) {
              let url = this.addressRoute(delegatorAddress.address);
							this.$router.push(url);
							this.clearSearchInputValue();
						} else {
							this.toSearchResultPage()
						}
					}catch (e) {
						console.error(e);
						this.toSearchResultPage();
					}
                });
			},
			searchValidator () {
				Service.commonInterface({headerSearchCandidate:{searchValue:this.searchInputValue}}, (validatorAddress) => {
					try {
						if (validatorAddress) {
                            let url = this.addressRoute(validatorAddress.validator.address);
							this.$router.push(url);
							this.clearSearchInputValue();
						} else {
							this.toSearchResultPage()
						}
					}catch (e) {
						console.error(e);
						this.toSearchResultPage();
					}
                });
			},
			searchBlockAndProposal () {
				Service.commonInterface({headerSearchValue:{searchValue:this.searchInputValue}},(searchResult) => {
					try {
						if (searchResult) {
							let searchResBlockLength = Object.keys(searchResult.block).length,
                                searchResProposalLength = Object.keys(searchResult.proposal).length,
                                searchResultLength = 0;
                            if(searchResBlockLength !== searchResultLength && searchResProposalLength !== searchResultLength ){
	                            this.toSearchResultPage();
                            }else if(searchResBlockLength !== searchResultLength && searchResProposalLength === searchResultLength){
	                            this.$router.push(`/block/${searchResult.block.height}`);
	                            this.clearSearchInputValue();
                            }else if(searchResBlockLength === searchResultLength && searchResProposalLength !== searchResultLength){
	                            this.$router.push(`/ProposalsDetail/${searchResult.proposal.proposal_id}`);
	                            this.clearSearchInputValue();
                            }else if(searchResBlockLength === searchResultLength && searchResProposalLength === searchResultLength){
	                            this.toSearchResultPage();
                            }
						} else {
							this.toSearchResultPage();
						}
					}catch (e) {
						console.error(e);
						this.toSearchResultPage();
					}
                });
			},
			getData () {
				if (Tools.removeAllSpace(this.searchInputValue) === '') {
					this.clearSearchContent();
					return
				} else {
					if (/^[A-F0-9]{64}$/.test(this.searchInputValue)) {
						this.searchTx()
					} else if (this.$Codec.Bech32.isBech32(this.$Crypto.config.iris.bech32.accAddr, this.searchInputValue)) {
						this.searchDelegator();
					} else if (this.$Codec.Bech32.isBech32(this.$Crypto.config.iris.bech32.valAddr, this.searchInputValue)) {
						this.searchValidator();
					} else if (/^\+?[1-9][0-9]*$/.test(this.searchInputValue)) {
						this.searchBlockAndProposal();
					} else {
						this.toSearchResultPage();
					}
				}
			},
			toSearchResultPage () {
				this.$router.push(`/searchResult?${this.searchInputValue}`);
				this.searchInputValue = "";
			},
			onInputChange () {
				this.getData();
			},
			clearSearchInputValue () {
				this.searchInputValue = "";
			},
			hideActiveStyle(){
				this.activeStaking  = false;
				this.activeTransfers  = false;
				this.activeBlockChain  = false;
				this.activeAssets  = false;
				this.activeGov  = false;
				this.activeStats  = false;
				this.activeClassName  = false;
				this.activeFaucet = false;
            },
			listenRouteForChangeActiveButton () {
				//刷新的时候路由不变，active按钮不变
                this.hideActiveStyle();
				if (this.$route.fullPath === '/txs/validations' || this.$route.fullPath === '/txs/delegations') {
					this.activeStaking  = true
				}else if(this.$route.fullPath === '/txs/transfers'){
					this.activeTransfers  = true
                }else if(this.$route.fullPath === '/home' || this.$route.fullPath === '/blocks' || this.$route.fullPath === '/txs' ){
					this.activeBlockChain  = true
                }else if (this.$route.fullPath === '/assets/ntvassetstxs' || this.$route.fullPath === '/assets/gtwassetstxs'){
					this.activeAssets  = true
                }else if(this.$route.fullPath === '/gov/parameters' || this.$route.fullPath === '/gov/proposals' || this.$route.fullPath === '/txs/governance'){
					this.activeGov = true
                }else if(this.$route.fullPath === '/stats/irisrichlist' || this.$route.fullPath === '/stats/irisstats'){
					this.activeStats= true
                }else if(this.$route.fullPath === '/faucet'){
					this.activeFaucet = '/faucet';
				}else if(this.$route.fullPath === '/validators') {
					if(this.hoverBlockChainTag){
						this.activeStaking  = true
                    }else {
						this.activeBlockChain  = true
                    }
                }
			},
			clearSearchContent () {
				this.searchInputValue = '';
			},
			toHelp () {
				this.$router.push('/help')
			},
			getConfig () {
				Service.commonInterface({headerConfig:{}},(res) => {
					try {
						this.flShowLogo = true;
						if(!localStorage.getItem('currentEnv')){
							this.toggleTestnetLogo(res);
						}
						this.setCurrentSelectOption(res.cur_env, res.chain_id, res.configs);
						this.setNetWorkLogo(res.cur_env, res.chain_id);
						this.setEnvConfig(res);
						this.handleConfigs(res.configs);
						this.flShowHeaderNetwork = true;
						this.chainId = `${res.chain_id.toUpperCase()} ${res.cur_env.toUpperCase()}`;
						res.configs.forEach(item => {
							if (res.cur_env === item.env_nm && res.chain_id === item.chain_id) {
								if (item.show_faucet && item.show_faucet === 1) {
									this.flShowFaucet = true;
									sessionStorage.setItem("Show_faucet", JSON.stringify(1))
								} else {
									this.flShowFaucet = false;
									sessionStorage.setItem("Show_faucet", JSON.stringify(0))
								}
							}
						})
					}catch (e) {
						console.error(e);
						this.explorerLogo = require("../assets/logo.png")
					}
                });
			},
			handleConfigs (configs) {
				this.netWorkArray = configs.map(item => {
					if(item.env_nm === constant.ENVCONFIG.MAINNET && item.chain_id === constant.CHAINID.MAINNET){
						item.icon = 'iconfont iconiris'
                    }else if(item.env_nm === constant.ENVCONFIG.TESTNET && item.chain_id === constant.CHAINID.FUXI){
						item.icon = 'iconfont iconfuxi1'
                    }else if(item.env_nm === constant.ENVCONFIG.TESTNET && item.chain_id !== constant.CHAINID.FUXI){
						item.icon = 'iconfont iconcaihongmao'
                    }
					item.netWorkSelectOption = `${Tools.firstWordUpperCase(item.env_nm)} ${item.chain_id.toLocaleUpperCase()}`
					return item
				});
				this.netWorkArray = this.netWorkArray.filter(item => {
					return item.env_nm !== constant.ENVCONFIG.DEV && item.env_nm !== constant.ENVCONFIG.QA && item.env_nm !== constant.ENVCONFIG.STAGE
				});
			},
			setEnvConfig (currentEnv) {
				sessionStorage.setItem("currentEnv",currentEnv.cur_env);
				if (currentEnv.cur_env !== constant.ENVCONFIG.MAINNET) {
					this.$Crypto.getCrypto(constant.CHAINNAME, constant.ENVCONFIG.TESTNET);
					this.$store.commit('currentEnv', constant.ENVCONFIG.TESTNET)
				}
			},
			toggleTestnetLogo (currentEnv) {
                this.$store.commit('currentSkinStyle',`${currentEnv.cur_env}${currentEnv.chain_id}`);
				const root = document.documentElement;
				root.style.setProperty(skinStyle.skinStyle.TITLECOLORNAME,skinStyle.skinStyle.commonFontBlackColor);
				root.style.setProperty(skinStyle.skinStyle.CONTENTCOLORNAME,skinStyle.skinStyle.commonFontContentColor);
				root.style.setProperty(skinStyle.skinStyle.MODULEBLACKCOLOR,skinStyle.skinStyle.commonModuleBlackColor);
				if (currentEnv.cur_env === constant.ENVCONFIG.MAINNET && currentEnv.chain_id === constant.CHAINID.MAINNET) {
					root.style.setProperty(skinStyle.skinStyle.BGCOLORNAME,skinStyle.skinStyle.MAINNETBGCOLOR);
					root.style.setProperty(skinStyle.skinStyle.HOVERCOLORNAME,skinStyle.skinStyle.MAINNETHOVERCOLOR);
					root.style.setProperty(skinStyle.skinStyle.ACTIVECOLORNAME,skinStyle.skinStyle.MAINNETACTIVECOLOR);
				} else if(currentEnv.cur_env === constant.ENVCONFIG.TESTNET && currentEnv.chain_id === constant.CHAINID.FUXI) {
					root.style.setProperty(skinStyle.skinStyle.BGCOLORNAME,skinStyle.skinStyle.TESTNETBGCOLOR);
					root.style.setProperty(skinStyle.skinStyle.HOVERCOLORNAME,skinStyle.skinStyle.TESTNETHOVERCOLOR);
					root.style.setProperty(skinStyle.skinStyle.ACTIVECOLORNAME,skinStyle.skinStyle.TESTNETACTIVECOLOR);

				}else if(currentEnv.cur_env === constant.ENVCONFIG.TESTNET && currentEnv.chain_id !== constant.CHAINID.FUXI){
					root.style.setProperty(skinStyle.skinStyle.BGCOLORNAME,skinStyle.skinStyle.NYANCATTESTNETBGCOLOR);
					root.style.setProperty(skinStyle.skinStyle.HOVERCOLORNAME,skinStyle.skinStyle.NYANCATTESTNETHOVERCOLOR);
					root.style.setProperty(skinStyle.skinStyle.ACTIVECOLORNAME,skinStyle.skinStyle.NYANCATTESTNETACTIVECOLOR);
                }else{
					root.style.setProperty(skinStyle.skinStyle.BGCOLORNAME,skinStyle.skinStyle.DEFAULTBGCOLOR);
					root.style.setProperty(skinStyle.skinStyle.HOVERCOLORNAME,skinStyle.skinStyle.DEFAULTHOVERCOLOR);
					root.style.setProperty(skinStyle.skinStyle.ACTIVECOLORNAME,skinStyle.skinStyle.DEFAULTACTIVECOLOR);
                }
				localStorage.setItem('currentEnv',`${currentEnv.cur_env}${currentEnv.chain_id}`)
			},
			setCurrentSelectOption (currentEnv, currentChainId) {
				if (currentEnv === constant.ENVCONFIG.DEV || currentEnv === constant.ENVCONFIG.QA || currentEnv === constant.ENVCONFIG.STAGE) {
					this.currentSelected = lang.home.mainnet;
				} else {
					this.currentSelected = `${currentChainId.toLocaleUpperCase()} ${Tools.firstWordUpperCase(currentEnv)}`;
				}
			},
			setNetWorkLogo (currentEnv, currentChainId) {
				if (currentEnv === constant.ENVCONFIG.MAINNET && currentChainId === constant.CHAINID.MAINNET) {
					this.flShowGatewayMenu = false;
					this.explorerLogo = require("../assets/logo.png");
					this.currentNetworkClass = 'iconfont iconiris'
				} else if (currentEnv === constant.ENVCONFIG.TESTNET && currentChainId === constant.CHAINID.FUXI) {
					this.flShowGatewayMenu = true;
					this.explorerLogo = require("../assets/logo.png");
					this.currentNetworkClass = 'iconfont iconfuxi'
				} else if (currentEnv === constant.ENVCONFIG.TESTNET && currentChainId === constant.CHAINID.NYANCAT) {
					this.flShowGatewayMenu = true;
					this.explorerLogo = require("../assets/logo.png");
					this.currentNetworkClass = 'iconfont iconcaihongmao'
				} else {
					this.flShowGatewayMenu = true;
					this.currentNetworkClass = 'iconfont iconiris';
					this.explorerLogo = require("../assets/logo.png")
				}
			},
			showNetWorkLogo(){
				this.flShowNetworkLogo = true;
            },
			hideNetWorkLogo(){
				this.flShowNetworkLogo = false;
            },
			windowOpenUrl (url) {
				window.open(url)
			},
		},
		updated () {
			this.absoluteTop = `${this.$refs.header_content.clientHeight / 100}rem`
		}
	}
</script>
<style lang="scss">
    @import "../style/mixin.scss";
    :root{
    }
    .person_computer_header_var {
        position: fixed;
        z-index: 11;
        background: rgba(255, 255, 255, 1);
        height: 0.6rem;
        .app_header_person_computer {
            width: 100%;
            background: var(--bgColor);
            .header_navigation_container{
                max-width: 12.8rem;
                margin: 0 auto;
                background: var(--bgColor);
                box-sizing: border-box;
                padding: 0 0.2rem;
                display: flex;
                flex-direction: column;
                align-items: center;
                position: relative;
                .header_navigation_wrap{
                    width: 100%;
                    height:0.6rem;
                    display: flex;
                    justify-content: space-between;
                    .header_left_content{
                        display: flex;
                        .header_logo_content{
                            height: 100%;
                            width: 1.5rem;
                            padding: 0.1rem 0;
                            a{
                                display: inline-block;
                                width: 100%;
                                padding-right: 0.1rem;
                                img {
                                    height: 100%;
                                    width: 100%;
                                    max-width: 1.5rem;
                                }
                            }

                        }
                        .header_menu_content{
                            display: flex;
                            max-width: 12.8rem;
                            .header_menu_item{
                                padding: 0 0.1rem;
                                color: #fff;
                                height: 0.6rem;
                                line-height: 0.6rem;
                                font-size: 0.14rem;
                                cursor: pointer;
                                a{
                                    color: #fff !important;
                                }
                                &:hover{
                                    background: var(--activeColor);
                                }
                            }
                            .nav_item_active {
                                color: #ffffff;
                                background: var(--activeColor);
                            }
                        }
                    }
                    .header_right_content{
                        display: flex;
                        flex: 1;
                        align-items: center;
                        .search_input_container {
                            flex: 1;
                            background: var(--bgColor);
                            z-index: 1;
                            .search_input_wrap {
                                max-width: 12.8rem;
                                width: 100%;
                                margin: 0 auto;
                                padding: 0.1rem 0 0.1rem 0.2rem;
                                display: flex;
                                align-items: center;
                                input {
                                    width: 100%;
                                    height: 0.35rem;
                                    border-radius: 0.06rem 0 0 0.06rem;
                                    box-shadow: none;
                                    background: var(--bgColor);
                                    border: 0.01rem solid rgba(255,255,255,0.5);
                                    color: #fff;
                                    font-size: 0.14rem;
                                    border-right: none;
                                    text-indent: 0.1rem;
                                }
                                input::placeholder{
                                    font-size: 0.14rem;
                                    color:rgba(255,255,255,0.5);
                                }
                                span {
                                    right: 0.3rem;
                                    height:0.35rem;
                                    font-size: 0.2rem;
                                    padding: 0 0.1rem;
                                    line-height: 0.33rem;
                                    color: rgba(255,255,255,0.5);
                                    border-top: 0.01rem solid rgba(255,255,255,0.5);
                                    border-right: 0.01rem solid rgba(255,255,255,0.5);
                                    border-bottom: 0.01rem solid rgba(255,255,255,0.5);
                                    background: var(--bgColor);
                                    border-radius: 0 0.06rem 0.06rem 0;
                                    cursor: pointer;
                                }
                            }
                        }
                        .network_container{
                            position: relative;
                            height:0.6rem;
                            line-height: 0.6rem;
                            padding-left: 0.2rem;
                            .network_list_container{
                                background: #fff;
                                box-shadow: 0 0.02rem 0.1rem 0 rgba(3,16,114,0.15);
                                width: auto;
                                position: absolute;
                                right: 0;
                                top: 0.6rem;
                                z-index: 2;
                                text-align: right;
                                .network_list_item{
                                    height: 0.4rem;
                                    line-height: 0.4rem;
                                    white-space: nowrap;
                                    padding: 0 0.2rem;
                                    cursor: pointer;
                                    font-size: 0.14rem;
                                    display: flex;
                                    &:hover{
                                        background: #F6F7FF;
                                    }
                                    i{
                                        font-size: 0.18rem;
                                        color: var(--titleColor);
                                        padding-right: 0.2rem;
                                    }
                                }
                                .network_list_item:last-child{
                                    padding-bottom: 0.05rem;
                                }
                            }
                        }
                    }
                }
            }
                .header_submenu_container{
                    background: transparent;
                    color: var(--contentColor);
                    margin: 0 auto;
                    position: absolute;
                    top:0.6rem;
                    width: 100%;
                    //box-shadow: 0 0.02rem 0.1rem 0 rgba(3,16,114,0.15);
                    z-index: 11;
                    animation: flShowMenu;
                    .header_submenu_content_wrap{
                        max-width: 12.8rem;
                        margin: 0 auto;
                        width: 100%;
                        position: relative;
                        .header_submenu_content{
                            display: flex;
                            align-items: flex-start;
                            position: absolute;
                            justify-content: flex-start;
                            flex-direction: column;
                            background: #fff;
                            box-shadow: 0 0.02rem 0.05rem 0.02rem rgba(0,0,0,0.15);
                            overflow:hidden;
                            .header_submenu_item{
                                box-sizing: border-box;
                                font-size: 0.14rem;
                                line-height: 1;
                                padding: 0.1rem  0.15rem 0.1rem 0.15rem;
                                width: 100%;
                                background: #fff;
                                &:hover{
                                    background: #F6F7FF;
                                }
                                a{
                                    &:hover{
                                        color:var(--bgColor) !important;
                                    }
                                }
                        }
                            .no_border_style{
                            border-bottom:none;
                        }
                    }
                }
            }
        }
    }
    .mobile_header_var {
        position: relative;
        z-index: 2;
    }
    .person_computer_header_var,
    .mobile_header_var {
        @include flex();
        @include pcContainer;
        .useFeature {
            width: 100%;
            height: 0.66rem;
            @include flex;
            flex-direction: column;
            align-items: center;
            background: var(--bgColor);
            .navButton {
                width: 100% !important;
                padding: 0 0.2rem;
                height: 0.66rem;
                @include pcCenter;
                @include flex;
                .common_item_style {
                    &:hover {
                        color: #ffffff;
                        background: var(--hoverColor);
                    }
                }
                .nav_item {
                    display: inline-block;
                    height: 0.65rem;
                    line-height: 0.66rem;
                    width: 1.6rem;
                    text-align: center;
                    font-size: 0.18rem;
                    cursor: pointer;
                    color: #c9eafd;
                    font-weight: 500 !important;
                    .bottom_arrow {
                        display: inline-block;
                        height: 0.11rem;
                        width: 0.11rem;
                        background: url("../assets/caret-bottom.png") no-repeat 0 0;
                        top: 0.27rem;
                        right: 0.1rem;
                    }
                }
                .sub_btn_wrap {
                    @include flex;
                    flex-direction: column;
                    z-index: 100;
                    padding: 0;
                    .sub_btn_item {
                        height: 0.36rem;
                        line-height: 0.36rem;
                        font-size: 0.14rem;
                        background: var(--bgColor);
                        color: #c9eafd;
                        width: 1.6rem;
                        text-align: left;
                        padding-left: 0.2rem;
                        a {
                            width: 100%;
                            display: inline-block;
                            padding-left: 0.19rem;
                            color: #c9eafd !important;
                            &:hover {
                                color: #00f0ff !important;
                            }
                        }
                        &:hover {
                            color: #00f0ff;
                        }
                    }
                    .top_list_option {
                        padding-left: 0.38rem;
                    }
                    .validators_btn_item {
                        padding-left: 0.35rem;
                    }
                }

                .btn-group,
                .btn-group-vertical {
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
            background: var(--bgColor);
            .feature_btn {
                width: 0.25rem;
                height: 0.25rem;
                top: 0.26rem;
                right: 0.1rem;
                img {
                    width: 100%;
                }
            }
            .image_wrap_mobile {
                @include flex;
                align-items: center;
                width: 1.5rem;
                height: 0.5rem;
                img {
                    width: 100%;
                    height: 100%;
                }
                i{
                    padding-left: 0.1rem;
                    font-size: 0.25rem;
                    color: #fff;
                }
                .logo_title_wrap {
                    margin-left: 0.16rem;
                    @include flex;
                    flex-direction: column;
                    justify-content: center;
                    .logo_title_content {
                        height: 0.27rem;
                        line-height: 0.27rem;
                        span {
                            &:first-child {
                                font-size: 0.24rem;
                                color: var(--bgColor);
                            }
                            &:last-child {
                                font-size: 0.24rem;
                                color: var(--bgColor);
                            }
                        }
                    }
                    p {
                        font-size: 0.14rem;
                        span {
                            color: #aeaeb9;
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
                    box-shadow: none;
                    width: 100%;
                    @include borderRadius(0.06rem);
                    border: 0.01rem solid #eee;
                    font-size: 0.14rem;
                    padding: 0 0.48rem 0 0.1rem;
                    height: 0.3rem;
                    &:focus {
                        border: 0.01rem solid var(--bgColor);
                        outline: none;
                    }
                }
                .search_icon {
                    position: absolute;
                    top: 0.07rem;
                    right: 0.15rem;
                    width: 0.15rem;
                    height: 0.15rem;
                    background: url("../assets/search.svg") no-repeat;
                    cursor: pointer;
                }
                .clear_icon {
                    position: absolute;
                    top: 0.08rem;
                    right: 0.32rem;
                    width: 0.15rem;
                    height: 0.15rem;
                    background: url("../assets/clear.png") no-repeat;
                    cursor: pointer;
                }
            }
            .mobile_chain_id_content {
                height: 0.3rem;
                margin: 0 0.15rem;
                background: rgba(234, 248, 254, 1);
                display: flex;
                justify-content: flex-start;
                line-height: 0.3rem;
                font-size: 0.14rem;
                margin-bottom: 0.1rem;
                border-radius: 0.06rem;
                .mobile_chain_content {
                    padding-left: 0.1rem;
                }
            }
            .use_feature_mobile {
                z-index: 1010;
                width: 100%;
                left: 0;
                background: var(--bgColor);
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
                .mobile_menu_container{
                    display: flex;
                    flex-direction: column;
                    .mobile_menu_item_content{
                        display: flex;
                        justify-content: space-between;
                        align-items: center;
                        box-sizing: border-box;
                        color: #fff;
                        padding: 0.05rem 0.15rem;
                        i{
                            font-size: 0.08rem;
                        }
                        .up_style{
                            transform: rotate(180deg);
                        }
                        .down_style{
                            transform: rotate(0deg);
                        }
                    }
                    .blockchain_list_content{
                        display: flex;
                        flex-direction:column;
                        .blockchain_list_item{
                            padding: 0.05rem 0.15rem;
                            font-size: 0.14rem;
                            background: var(--hoverColor);
                            color:#fff;
                            display: flex;
                            justify-content: space-between;
                            a{
                                color:#fff !important;
                            }
                        }
                    }
                }

                .feature_btn_mobile {
                    height: 0.39rem;
                    line-height: 0.39rem;
                    padding-left: 0.15rem;
                    background: var(--bgColor);
                    color: #fff;
                    font-size: 0.14rem;
                    a {
                        display: inline-block;
                        width: 100%;
                        color: #fff !important;
                    }
                }
                .feature_arrow {
                    position: relative;
                    color: #c9eafd;
                    font-size: 0.14rem;
                    background: url("../assets/caret-bottom.png") no-repeat 97% 0.12rem,
                    var(--bgColor);
                }
                .feature_subNav {
                    padding-left: 0.3rem;
                }
            }
        }
    }
    .chain_id {
        padding-right: 0.26rem;
        font-size: 0.16rem;
        color: #f2711c;
        @include fontWeight;
    }
    .select_option_container {
        display: flex;
        padding-right: 0.2rem;
        justify-content: space-between;
        .upImg_content {
            width: 0.28rem;
            padding-left: 0.2rem;
            img {
                width: 100%;
            }
        }
        .downImg_content {
            width: 0.28rem;
            padding-left: 0.2rem;
            img {
                width: 100%;
            }
        }
    }
    @media screen and (max-width: 1200px) {
        .person_computer_header_var {
            .app_header_person_computer {
                .header_top_container {
                    .header_top_content_wrap {
                        .header_menu_content {
                            .header_menu_list_content {
                                margin-left: 0.08rem;
                                li {
                                    padding: 0 0.08rem;
                                }
                                .nav_item_active {
                                    padding: 0 0.08rem;
                                }
                                .transaction_list_content {
                                    padding: 0;
                                    .transaction_content {
                                        padding: 0 0.08rem;
                                    }
                                    .transaction_list_item_content {
                                        a {
                                            .transaction_list_item {
                                                padding: 0 0.08rem;
                                            }
                                        }
                                    }
                                }
                                .statics_list_content {
                                    padding: 0;
                                    .statics_content {
                                        padding: 0 0.08rem;
                                    }
                                    .statics_list_item_content {
                                        a {
                                            .static_list_item {
                                                padding: 0 0.08rem;
                                            }
                                        }
                                    }
                                }
                                .governance_list_content {
                                    padding: 0;
                                    .governance_content {
                                        padding: 0 0.08rem;
                                    }
                                    .governance_list_item_content {
                                        a {
                                            .governance_list_item {
                                                padding: 0 0.08rem;
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    .fade-enter-active, .fade-leave-active {
        transition: all 0.2s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }
    .fade-enter, .fade-leave-to{
        opacity: 0;
    }
</style>
