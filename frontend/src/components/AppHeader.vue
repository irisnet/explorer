<!--头部页面-->
<template>
    <div :class="appHeaderVar"
         v-show="showHeader"
         id="header">
        <header class="app_header_person_computer"
                v-show="devicesShow === 1">
            <div class="header_top_container">
                <div class="header_top_content_wrap">
                    <div class="header_logo_content">
                        <router-link :to="`/home`"><img :src="explorerLogo"></router-link>
                    </div>
                    <div class="header_menu_content">
                        <ul class="header_menu_list_content">
                            <li :class="activeClassName === '/home'?'nav_item_active':''"
                                @click="featureButtonClick('/home')">Home</li>
                            <li :class="activeClassName === '/validators'?'nav_item_active':''"
                                @click="featureButtonClick('/validators')">Validators</li>
                            <li :class="activeClassName === '/blocks'?'nav_item_active':''"
                                @click="featureButtonClick('/blocks')">Blocks</li>
                            <li class="transaction_list_content"
                                :class="activeClassName === '/transaction'?'nav_item_active':''"
                                @mouseover="transactionMouseOver"
                                @mouseleave="transactionMouseLeave">
                                <div class="transaction_content"
                                     style="display: flex">Transactions<span class="bottom_arrow"></span></div>
                                <ul class="transaction_list_item_content"
                                    v-show="showSubTransaction">
                                    <router-link :to="`/txs/transfers`">
                                        <li class="transaction_list_item">Transfers</li>
                                    </router-link>
                                    <router-link :to="`/txs/declarations`">
                                        <li class="transaction_list_item">Declarations</li>
                                    </router-link>
                                    <router-link :to="`/txs/stakes`">
                                        <li class="transaction_list_item">Stakes</li>
                                    </router-link>
                                    <router-link :to="`/txs/governance`">
                                        <li class="transaction_list_item">Governance</li>
                                    </router-link>
                                </ul>
                            </li>
                            <li class="assets_list_content"
                                :class="activeClassName === '/Assets'?'nav_item_active':''"
                                @mouseover="nativeAssetMouseOver"
                                @mouseleave="nativeAssetMouseLeave">
                                <div class="assets_content"
                                     style="display: flex">Assets<span class="bottom_arrow"></span></div>
                                <ul class="assets_list_item_content"
                                    v-show="showAssetTransaction">
                                    <router-link :to="`/nativeasset`">
                                        <li class="assets_list_item">Native Asset</li>
                                    </router-link>
                                    <router-link :to="`/gatewayasset`">
                                        <li class="assets_list_item">Gateway Asset</li>
                                    </router-link>
                                </ul>
                            </li>
                            <li class="statics_list_content"
                                :class="activeClassName === '/statistics'?'nav_item_active':''"
                                @mouseover="statisticsMouseOver"
                                @mouseleave="statisticsMouseLeave">
                                <div class="statics_content"
                                     style="display: flex">Statistics<span class="bottom_arrow"></span></div>
                                <ul class="statics_list_item_content"
                                    style="color: #000;background: --bgColor"
                                    v-show="flShowStatistics">
                                    <router-link :to="`/statistics/richlist`">
                                        <li class="static_list_item">Rich List</li>
                                    </router-link>
                                    <router-link :to="`/statistics/tokenstats`">
                                        <li class="static_list_item">Tokens Stats</li>
                                    </router-link>
                                    <!-- <router-link :to="`/statistics/bondedTokens`">
                                        <li class="static_list_item">Bonded Tokens</li>
                                    </router-link> -->
                                </ul>
                            </li>
                            <li class="governance_list_content"
                                :class="activeClassName === '/governance'?'nav_item_active':''"
                                @mouseover="governanceMouseOver"
                                @mouseleave="governanceMouseLeave">
                                <div class="governance_content"
                                     style="display: flex">Governance<span class="bottom_arrow"></span></div>
                                <ul class="governance_list_item_content"
                                    style="color: #000;background: --bgColor"
                                    v-show="flShowGovernanceOption">
                                    <router-link :to="`/gov/parameters`">
                                        <li class="governance_list_item">Parameters</li>
                                    </router-link>
                                    <router-link :to="`/gov/proposals`">
                                        <li class="governance_list_item">Proposals</li>
                                    </router-link>
                                </ul>
                            </li>
                            <li v-if="flShowFaucet"
                                class="nav_item common_item_style faucet_content"
                                :class="activeClassName === '/faucet'?'nav_item_active':''"
                                @click="featureButtonClick('/faucet')">Faucet</li>
                        </ul>
                        <div class="header_network_content">
                            <div class="network_list_content">
                                <div class="network_list_title_content"
                                     style="display: flex">{{currentSelected}}<span class="bottom_arrow"></span></div>
                                <ul class="network_list_item_content"
                                    style="color: #000;background: --bgColor">
                                    <li class="network_list_item"
                                        v-for="item in netWorkArray"
                                        v-show="item.netWorkSelectOption.trim() !== currentSelected.trim()"
                                        @click="windowOpenUrl(item.host)">{{item.netWorkSelectOption}}</li>
                                </ul>
                            </div>
                            <div class="search_content"
                                 @click="toggleShowSearchIpt()">
                                <img src="../assets/search_icon.png"
                                     alt="">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="search_input_container"
                 v-show="flShowSearchIpt">
                <div class="search_input_wrap">
                    <input type="text"
                           class="search_input"
                           placeholder="Search by Address / Txhash / Block"
                           v-model.trim="searchInputValue"
                           @keyup.enter="onInputChange">
                    <span @click="getData(searchInputValue)">Search</span>
                </div>
            </div>
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
            <div class="mobile_chain_id_content"
                 v-if="flShowChainId">
                <span class="mobile_chain_content">{{chainId}}</span>
            </div>
            <div class="use_feature_mobile"
                 :style="{'top':absoluteTop}"
                 v-show="featureShow">
        <span class="feature_btn_mobile feature_nav"
            @click="featureButtonClick('/home')">Home</span>
                <span class="feature_btn_mobile feature_nav select_option_container"
                    @click="featureButtonClick('/validators')">
                    <span>Validators</span>
                </span>
                <span class="feature_btn_mobile feature_nav"
                    @click="featureButtonClick('/blocks')">Blocks</span>
                <span class="feature_btn_mobile feature_nav select_option_container"
                    @click="transactionsSelect(flShowTransactionsSelect)">
                    <span>Transactions</span>
                    <div :class="flShowUpOrDown ? 'upImg_content' : 'downImg_content'">
                       <img :src="flShowUpOrDown ? upImg : downImg ">
                    </div>
                </span>
                <div class="select_option" v-show="flShowTransactionsSelect">
                    <span class="feature_btn_mobile feature_nav"
                          @click="featureButtonClick('/txs/transfers')">Transfers</span>
                <span class="feature_btn_mobile feature_nav"
                      @click="featureButtonClick('/txs/declarations')">Declarations</span>
                <span class="feature_btn_mobile feature_nav"
                      @click="featureButtonClick('/txs/stakes')">Stakes</span>
                <span class="feature_btn_mobile feature_nav"
                      @click="featureButtonClick('/txs/governance')">Governance</span>
            </div>


                <span class="feature_btn_mobile feature_nav select_option_container"
                      @click="assetsSelect(flShowAssetSelect)">
                    <span>Assets</span>
                    <div :class="flShowUpOrDown ? 'upImg_content' : 'downImg_content'">
                       <img :src="flShowUpOrDown ? upImg : downImg ">
                    </div>
                </span>
                <div class="select_option" v-show="flShowAssetSelect">
                    <span class="feature_btn_mobile feature_nav"
                          @click="featureButtonClick('/nativeasset')">Native Asset</span>
                    <span class="feature_btn_mobile feature_nav"
                          @click="featureButtonClick('/gatewayasset')">Declarations</span>

                </div>





                <span class="feature_btn_mobile feature_nav select_option_container"
                    @click="topListSelect(flShowTopListSelection)">
                    <span>Statistics</span>
                    <div :class="flShowUpOrDown ? 'upImg_content' : 'downImg_content'">
                       <img :src="flShowUpOrDown ? upImg : downImg ">
                    </div>
                </span>
                <div class="select_option"
                     v-show="flShowTopListSelection">
                <span class="feature_btn_mobile feature_nav"
                    @click="featureButtonClick('/statistics/richlist')">Rich List</span>
                <span class="feature_btn_mobile feature_nav"
                    @click="featureButtonClick('/statistics/tokenstats')">Tokens Stats</span>
                <!-- <span class="feature_btn_mobile feature_nav"
                    @click="featureButtonClick('/statistics/bondedTokens')">Bonded Tokens</span> -->
                </div>
                <span class="feature_btn_mobile feature_nav select_option_container"
                      @click="governanceSelect(flShowGovernanceSelect)">
            <span>Governance</span>
            <div :class="flShowUpOrDown ? 'upImg_content' : 'downImg_content'">
               <img :src="flShowUpOrDown ? upImg : downImg ">
            </div>
            </span>
                <div class="select_option"
                     v-show="flShowGovernanceSelect">
            <span class="feature_btn_mobile feature_nav"
                @click="featureButtonClick('/gov/parameters')">Parameters</span>
                    <span class="feature_btn_mobile feature_nav"
                          @click="featureButtonClick('/gov/proposals')">Proposals</span>

                </div>
                <span v-if="flShowFaucet"
                      class="feature_btn_mobile feature_nav mobile_faucet_content"
                      @click="featureButtonClick('/faucet')">Faucet</span>
                <span class="feature_btn_mobile feature_nav select_option_container"
                      @click="netWorkSelect(flShowNetworkSelect)">
            <span>Network</span>
            <div :class="flShowNetworkUpOrDown ? 'upImg_content' : 'downImg_content'">
                <img :src="flShowNetworkUpOrDown ? upImg : downImg ">
            </div>
            </span>
                <div class="select_option"
                     v-show="flShowNetworkSelect">
            <span class="feature_btn_mobile feature_nav"
                v-for="item in netWorkArray">
            <a :href="item.host"
               target="_blank">{{item.netWorkSelectOption}}</a>
            </span>
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
				flShowChainId: false,
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
				flShowSearchIpt: false,
				flShowLogo: false,
				showAssetTransaction: false
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
		},
		beforeDestroy () {
			document.getElementById('router_wrap').removeEventListener('click', this.hideFeature);
			window.removeEventListener('resize', this.onWindowResize);
		},
		methods: {
			transactionsSelect (flShowTransactionsSelect) {
				this.flShowValidatorsSelect = false;
				if (!flShowTransactionsSelect) {
					this.flShowTransactionsSelect = true;
					this.flShowUpOrDown = true
				} else {
					this.flShowUpOrDown = false;
					this.flShowTransactionsSelect = false
				}
			},
            assetsSelect (flShowAssetSelect) {
                this.flShowAssetSelect = false;
                if (!flShowAssetSelect) {
                    this.flShowAssetSelect = true;
                    this.flShowUpOrDown = true
                } else {
                    this.flShowUpOrDown = false;
                    this.flShowAssetSelect = false
                }
            },
			netWorkSelect (flShowNetworkSelect) {
				this.flShowNetworkSelect = false;
				if (!flShowNetworkSelect) {
					this.flShowNetworkSelect = true;
					this.flShowNetworkUpOrDown = true
				} else {
					this.flShowNetworkSelect = false;
					this.flShowNetworkUpOrDown = false
				}
			},
			governanceSelect (flShowNetworkSelect) {
				this.flShowGovernanceSelect = false;
				if (!flShowNetworkSelect) {
					this.flShowGovernanceSelect = true;
					this.flShowNetworkUpOrDown = true
				} else {
					this.flShowGovernanceSelect = false;
					this.flShowNetworkUpOrDown = false
				}
			},
			topListSelect (flShowTopListSelection) {
				this.flShowTopListSelection = false;
				if (!flShowTopListSelection) {
					this.flShowTopListSelection = true;
					this.flShowNetworkUpOrDown = true
				} else {
					this.flShowTopListSelection = false;
					this.flShowNetworkUpOrDown = false
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
			transactionMouseOver () {
				this.showSubTransaction = true;
			},
			transactionMouseLeave () {
				this.showSubTransaction = false;
			},
			nativeAssetMouseOver(){
				this.showAssetTransaction = true;
            },
			nativeAssetMouseLeave(){
				this.showAssetTransaction = false;
            },
			governanceMouseOver () {
				this.flShowGovernanceOption = true
			},
			governanceMouseLeave () {
				this.flShowGovernanceOption = false
			},
			statisticsMouseOver () {
				this.flShowStatistics = true
			},
			statisticsMouseLeave () {
				this.flShowStatistics = false
			},
			searchTx () {
				Service.commonInterface({headerTx:{searchValue: this.searchInputValue}}, (tx) => {
					try {
						if (tx) {
							this.$router.push(`/tx?txHash=${tx.Hash}`);
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
			listenRouteForChangeActiveButton () {
				//刷新的时候路由不变，active按钮不变
				let path = window.location.href;
				if (path.includes('txs') || path.includes('tx?')) {
					this.activeClassName = '/transaction';
				} else if (path.includes('/validators')) {
					this.activeClassName = '/validators';
				} else if (path.includes('/blocks')) {
					this.activeClassName = '/blocks';
				} else if (path.includes('/home')) {
					this.activeClassName = '/home';
				} else if (path.includes('/faucet')) {
					this.activeClassName = '/faucet';
				} else if (path.includes('/parameters')) {
					this.activeClassName = '/governance';
				} else if (path.includes('/proposals')) {
					this.activeClassName = '/governance';
				} else if (path.includes('/statistics')) {
					this.activeClassName = '/Assets';
				} else if (path.includes('/Assets')) {
					this.activeClassName = '/statistics';
				} else {
					this.activeClassName = '';
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
						this.toggleTestnetLogo(res);
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
					item.netWorkSelectOption = `${item.chain_id.toLocaleUpperCase()} ${Tools.firstWordUpperCase(item.env_nm)}`
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
				const root = document.documentElement;
				if (currentEnv.cur_env === constant.ENVCONFIG.MAINNET) {
					this.flShowChainId = false;
					root.style.setProperty(skinStyle.skinStyle.BGCOLORNAME,skinStyle.skinStyle.MAINNETBGCOLOR);
					root.style.setProperty(skinStyle.skinStyle.HOVERCOLORNAME,skinStyle.skinStyle.MAINNETHOVERCOLOR);
					root.style.setProperty(skinStyle.skinStyle.ACTIVECOLORNAME,skinStyle.skinStyle.MAINNETACTIVECOLOR);
				} else {
					this.flShowChainId = true;
					root.style.setProperty(skinStyle.skinStyle.BGCOLORNAME,skinStyle.skinStyle.TESTNETBGCOLOR);
					root.style.setProperty(skinStyle.skinStyle.HOVERCOLORNAME,skinStyle.skinStyle.TESTNETHOVERCOLOR);
					root.style.setProperty(skinStyle.skinStyle.ACTIVECOLORNAME,skinStyle.skinStyle.TESTNETACTIVECOLOR);
				}

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
					this.explorerLogo = require("../assets/logo.png")
				} else if (currentEnv === constant.ENVCONFIG.TESTNET && currentChainId === constant.CHAINID.FUXI) {
					this.explorerLogo = require("../assets/fuxi_testnet_logo.png")
				} else if (currentEnv === constant.ENVCONFIG.TESTNET && currentChainId === constant.CHAINID.NYANCAT) {
					this.explorerLogo = require("../assets/nyancat_testnet.png")
				} else {
					this.explorerLogo = require("../assets/logo.png")
				}
			},
			windowOpenUrl (url) {
				window.open(url)
			},
			toggleShowSearchIpt () {
				this.flShowSearchIpt = !this.flShowSearchIpt;
				this.$store.commit('flShowIpt', this.flShowSearchIpt)
			}
		},
		updated () {
			this.absoluteTop = `${this.$refs.header_content.clientHeight / 100}rem`
		}
	}
</script>
<style lang="scss">
    @import "../style/mixin.scss";
    :root{
        --bgColor: #fff;
    }
    .person_computer_header_var {
        position: fixed;
        z-index: 10001;
        background: rgba(255, 255, 255, 1);
        height: 0.6rem;
        .app_header_person_computer {
            width: 100%;
            background: var(--bgColor);
            .header_top_container {
                max-width: 12.8rem;
                width: 100%;
                margin: 0 auto;
                padding: 0 0.2rem;
                .header_top_content_wrap {
                    display: flex;
                    align-items: center;
                    .header_logo_content {
                        width: 1rem;
                        img {
                            width: 100%;
                        }
                    }
                    .header_menu_content {
                        flex: 1;
                        color: #fff;
                        display: flex;
                        justify-content: space-between;
                        .header_menu_list_content {
                            display: flex;
                            margin-left: 0.4rem;
                            li {
                                font-size: 0.14rem;
                                padding: 0 0.23rem;
                                height: 0.6rem;
                                line-height: 0.6rem;
                                cursor: pointer;
                                &:hover {
                                    background: var(--activeColor);
                                }
                            }
                            .nav_item_active {
                                color: #ffffff;
                                background: var(--activeColor);
                            }
                            .transaction_list_content {
                                z-index: 10000;
                                display: block;
                                padding: 0;
                                .transaction_content {
                                    box-sizing: border-box;
                                    padding: 0 0.23rem;
                                }
                                .bottom_arrow {
                                    display: inline-block;
                                    height: 0.6rem;
                                    width: 0.11rem;
                                    background: url("../assets/caret-bottom.png") no-repeat 50% 50%;
                                }
                                .transaction_list_item_content {
                                    z-index: 100;
                                    background: var(--activeColor);
                                    a {
                                        .transaction_list_item {
                                            height: 0.34rem;
                                            line-height: 0.34rem;
                                            text-align: left;
                                            color: #c9eafd;
                                            cursor: pointer;
                                            padding: 0 0.24rem;
                                            &:hover {
                                                background: var(--hoverColor);
                                            }
                                        }
                                    }
                                }
                            }
                            .assets_list_content {
                                z-index: 10000;
                                display: block;
                                padding: 0;
                                min-width: 1.42rem;
                                .assets_content {
                                    justify-content: center;
                                    box-sizing: border-box;
                                }
                                .bottom_arrow {
                                    display: inline-block;
                                    height: 0.6rem;
                                    width: 0.11rem;
                                    background: url("../assets/caret-bottom.png") no-repeat 50% 50%;
                                }
                                .assets_list_item_content {
                                    z-index: 100;
                                    background: var(--activeColor);
                                    a {
                                        .assets_list_item {
                                            height: 0.34rem;
                                            line-height: 0.34rem;
                                            text-align: left;
                                            color: #c9eafd;
                                            cursor: pointer;
                                            &:hover {
                                                background: var(--hoverColor);
                                            }
                                        }
                                    }
                                }
                            }
                            .statics_list_content {
                                padding: 0;
                                display: block;
                                z-index: 10000;
                                .bottom_arrow {
                                    display: inline-block;
                                    height: 0.6rem;
                                    width: 0.11rem;
                                    background: url("../assets/caret-bottom.png") no-repeat 50% 50%;
                                }
                                .statics_content {
                                    box-sizing: border-box;
                                    padding: 0 0.23rem;
                                }
                                .statics_list_item_content {
                                    z-index: 100;
                                    a {
                                        .static_list_item {
                                            height: 0.34rem;
                                            line-height: 0.34rem;
                                            text-align: left;
                                            color: #c9eafd;
                                            cursor: pointer;
                                            padding: 0 0.24rem;
                                            background: var(--activeColor);
                                            &:hover {
                                                background: var(--hoverColor);
                                            }
                                        }
                                    }
                                }
                            }
                            .governance_list_content {
                                padding: 0;
                                display:  block;
                                z-index: 10000;
                                .bottom_arrow {
                                    display: inline-block;
                                    height: 0.6rem;
                                    width: 0.11rem;
                                    background: url("../assets/caret-bottom.png") no-repeat 50% 50%;
                                }
                                .governance_content {
                                    box-sizing: border-box;
                                    padding: 0 0.23rem;
                                }
                                .governance_list_item_content {
                                    z-index: 100;
                                    a {
                                        .governance_list_item {
                                            height: 0.34rem;
                                            line-height: 0.34rem;
                                            text-align: left;
                                            color: #c9eafd;
                                            cursor: pointer;
                                            padding: 0 0.25rem;
                                            background: var(--activeColor);
                                            &:hover {
                                                background: var(--hoverColor);
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        .header_network_content {
                            display: flex;
                            align-items: center;
                            .network_list_content {
                                font-size: 0.14rem;
                                height: 0.6rem;
                                margin-right: 0.1rem;
                                line-height: 0.6rem;
                                z-index: 100;
                                cursor: pointer;
                                &:hover {
                                    background: var(--hoverColor);
                                    .network_list_item_content {
                                        display: block;
                                    }
                                }
                                .network_list_title_content {
                                    box-sizing: border-box;
                                    padding: 0 0.2rem;
                                    min-width: 1.8rem;
                                    &:hover {
                                        background: var(--hoverColor);
                                    }
                                    .bottom_arrow {
                                        display: inline-block;
                                        height: 0.6rem;
                                        width: 0.11rem;
                                        background: url("../assets/caret-bottom.png") no-repeat 50%
                                        50%;
                                    }
                                }
                                .network_list_item_content {
                                    display: none;
                                    .network_list_item {
                                        height: 0.34rem;
                                        line-height: 0.34rem;
                                        text-align: left;
                                        color: #c9eafd;
                                        cursor: pointer;
                                        padding: 0 0.2rem;
                                        background: var(--activeColor);
                                        &:hover {
                                            background: var(--hoverColor);
                                        }
                                    }
                                }
                            }
                            .search_content {
                                width: 0.35rem;
                                height: 0.35rem;
                                border: 0.01rem solid #fff;
                                border-radius: 0.06rem;
                                box-sizing: border-box;
                                padding: 0.04rem 0.08rem 0.08rem 0.08rem;
                                cursor: pointer;
                                img {
                                    width: 100%;
                                }
                            }
                        }
                    }
                }
            }
            .search_input_container {
                width: 100%;
                background: var(--bgColor);
                z-index: 1;
                .search_input_wrap {
                    max-width: 12.8rem;
                    width: 100%;
                    margin: 0 auto;
                    padding: 0.13rem 0.2rem;
                    display: flex;
                    align-items: center;
                    input {
                        text-indent: 0.22rem;
                        width: 100%;
                        height: 0.35rem;
                        border-radius: 0.06rem 0 0 0.06rem;
                        box-shadow: none;
                        border: none;
                    }
                    span {
                        color: var(--bgColor);
                        padding-right: 0.2rem;
                        background: #fff;
                        height: 0.35rem;
                        line-height: 0.35rem;
                        border-radius: 0 0.06rem 0.06rem 0;
                        cursor: pointer;
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
                .nav_item_active {
                    color: #ffffff;
                    background: var(--bgColor);
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
                /*background: url('../assets/menu.png') no-repeat;*/
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
</style>
