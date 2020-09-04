<template>
	<div class="search_bar_container" :class="$store.state.flShowSearchBar === 'show'  ? 'show_search_bar' : 'hide_search_bar'" >
		<div class="search_bar_wrap" >
			<div class="search_bar_content">
				<p class="search_bar_title">Welcome to IRISplorer</p>
				<div class="search_bar_ipt_content" :class="$store.state.flShowSearchIpt === 'show'  ? 'show_ipt' : 'hide_ipt'" v-if="$store.state.flShowSearchIpt === 'show'">
					<input class="search_bar_ipt"
					       type="text"
					       placeholder="Search by Address / Txhash / Block / HashLock"
					       v-model.trim="searchInputValue"
					       @keyup.enter="onInputChange">
					<div class="iconfont iconsousuo" @click="getData(searchInputValue)"></div>
				</div>
				<ul class="search_bar_link_content">
					<li class="search_bar_link_item" v-for="(item,index) in searchBarLinkArray" :key="index">
						<a :href="item.href" @click="UmengPush(item.label)"  target="_blank">{{item.label}}</a>
					</li>
				</ul>
				<div class="search_bar_pack_up_content">
					<span class="iconfont iconshouqi" @click="hideSearchBar()"></span>
				</div>
			</div>
		</div>
		<ul class="mobile_search_bar_link_content">
			<li class="mobile_search_bar_link_item" v-for="(item,index) in searchBarLinkArray" :key="index">
				<a :href="item.href" target="_blank">{{item.label}}</a>
			</li>
		</ul>
	</div>
</template>

<script>
	import Service from "../../service"
	import Tools from "../../util/Tools"
	export default {
		name: "SearchBar",
		data () {
			return {
				searchBarLinkArray:[
					{
						label:'IRISnetDocs',
						href:'https://www.irisnet.org/docs/'
					},
					{
						label:'RainbowWallet',
						href:'https://www.rainbow.one/'
					},
					{
						label:'Whitepaper',
						href:'https://github.com/irisnet/irisnet/blob/master/WHITEPAPER.md'
					},
					{
						label:'GitHub',
						href:'https://github.com/irisnet'
					}
				],
				searchInputValue:'',
			}
		},
		methods:{
			UmengPush(data){
				switch (data) {
					case 'IRISnetDocs':
						this.$uMeng.push('Overview_IRISnetDocs','click');
						break;
					case 'RainbowWallet':
						this.$uMeng.push('Overview_RainbowWallet','click');
						break;
					case 'Whitepaper':
						this.$uMeng.push('Overview_Whitepaper','click');
						break;
					case 'Github':
						this.$uMeng.push('Overview_Github','click');
				}
			},
			onInputChange () {
				this.$uMeng.push('Overview_Search','click');
				this.getData();
			},
			hideSearchBar(){
				this.$store.commit('flShowSearchBar','hide');
				this.$store.commit('showHeaderUnfoldBtn','show');
				sessionStorage.setItem('flShowHeaderUnfoldBtn','show');
				setTimeout( () => {
					// this.$store.commit('showHeaderUnfoldBtn','show');
					// sessionStorage.setItem('flShowHeaderUnfoldBtn','show');
					
					sessionStorage.setItem('flShowSearchBar', 'hide');
					
					this.$store.commit('flShowSearchIpt','hide');
					sessionStorage.setItem('flShowSearchIpt', 'hide');
					
				},300);
				this.$uMeng.push('Overview_Collapse','click');
				
			},
			getData () {
				if (Tools.removeAllSpace(this.searchInputValue) === '') {
					this.clearSearchContent();
					return
				} else {
					if (/^[A-F0-9]{64}$/.test(this.searchInputValue)) {
						this.searchTx();
					} else if(/^[a-f0-9]{64}$/.test(this.searchInputValue)){
						this.searchHashLock()
					} else if (Tools.isBech32(localStorage['acc_addr_prefix'] || this.$Crypto.config.iris.bech32.accAddr, this.searchInputValue)) {
						this.searchDelegator();
					} else if (Tools.isBech32(localStorage['val_addr_prefix'] || this.$Crypto.config.iris.bech32.valAddr, this.searchInputValue)) {
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
			searchHashLock(){
				Service.commonInterface({headerSearchHtlcs:{searchValue: this.searchInputValue}}, (hashlockInformation) => {
					try {
						if(hashlockInformation){
							this.$router.push(`/htlc/${hashlockInformation.hash_lock}`);
						}else {
							this.toSearchResultPage();
						}
					}catch (e) {
						console.error(e);
						this.toSearchResultPage();
					}
				})
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
			clearSearchInputValue () {
				this.searchInputValue = "";
			},
			clearSearchContent () {
				this.searchInputValue = '';
			},
		}
	}
</script>

<style scoped lang="scss">
	.hide_search_bar{
		display: none !important;
		/*animation: hideSearchBar 0.3s ease-in;*/
		/*animation-fill-mode: forwards;*/
	}
	.show_search_bar{
		display: block;
		/*animation: showSearchBar 0.3s ease-in;*/
		/*animation-fill-mode: forwards;*/
	}
	@keyframes hideSearchBar {
		from{
			height: 2.42rem;
		}
		to{
			height: 0;
		}
	}
	@keyframes showSearchBar {
		from{
			height: 0;
		}
		to{
			height: 2.42rem;
		}
	}
	@keyframes hideIpt {
		0%{
			height: 0.5rem;
		}
		100%{
			height: 0;
		}
	}
	@keyframes showIpt {
		0%{
			height: 0;
		}
		100%{
			height: 0.5rem;
		}
	}
	.search_bar_container{
		background: var(--bgColor);
		width: 100%;
		display: block;
		position: relative;
		.search_bar_wrap{
			max-width: 12.4rem;
			margin: 0 auto;
			display: flex;
			justify-content: center;
			flex-direction: column;
			.search_bar_content{
				.search_bar_title{
					margin-top: 0.5rem;
					font-size: 0.28rem;
					color:rgba(255,255,255,1);
					line-height: 0.32rem;
					text-align: center;
				}
				.hide_ipt{
					display: none;
					/*animation: hideIpt ease-in 0.3s;*/
					/*animation-fill-mode: forwards;*/
				}
				.show_ipt{
					display: block;
					/*animation: showIpt ease-in 0.3s;*/
					/*animation-fill-mode: forwards;*/
				}
				.search_bar_ipt_content{
					max-width: 8rem;
					margin: 0.3rem auto 0 auto;
					display: flex;
					height: 0.5rem;
					border-style:none;
					.search_bar_ipt{
						box-sizing: border-box;
						padding-left: 0.25rem;
						width: 100%;
						border-radius: 0.25rem 0 0 0.25rem;
						border: none;
						color: rgba(120,124,153,0.5);
					}
					.search_bar_ipt::placeholder{
						color:rgba(0,0,0,0.5);
					}
					.iconsousuo{
						cursor: pointer;
						width: 0.8rem;
						font-size: 0.25rem;
						display: flex;
						justify-content: center;
						align-items: center;
						background: #fff;
						border-radius: 0 0.25rem 0.25rem 0;
						color: rgba(0,0,0,0.3);
						border: none;
					}
				}
				.search_bar_link_content{
					max-width: 8rem;
					margin: 0.3rem auto 0.32rem auto;
					display: grid;
					justify-items: center;
					grid-template-columns: repeat(auto-fill, 25%);
					.search_bar_link_item{
						a{
							color: rgba(255,255,255,0.75) !important;
							&:hover{
								color: rgba(255,255,255,1)!important;
								border-bottom: 0.01rem solid rgba(255,255,255,1);
							}
						}
					}
				}
				.search_bar_pack_up_content{
					display: flex;
					justify-content: flex-end;
					position: absolute;
					bottom: 0.2rem;
					right: 0.2rem;
					.iconshouqi{
						cursor: pointer;
						font-size: 0.32rem;
						color: rgba(255,255,225,0.8);
					}
				}
			}
		}
		.mobile_search_bar_link_content{
			display: none;
		}
	}
	@media screen and (max-width: 910px){
		.search_bar_container{
			.search_bar_wrap{
				display: none !important;
			}
			.mobile_search_bar_link_content{
				max-width: 8rem;
				margin: 0.2rem auto 0.14rem auto;
				display: flex;
				justify-content: space-between;
				.mobile_search_bar_link_item{
					flex: 1;
					font-size: 0.12rem;
					text-align: center;
					a{
						color: rgba(255,255,255,0.4) !important;
						&:hover{
							color: rgba(255,255,255,1)!important;
							border-bottom: 0.01rem solid rgba(255,255,255,1);
						}
					}
				}
			}
		}
	}
</style>
