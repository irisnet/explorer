<template>
    <div class="search_result_detail_wrap">
        <div class="search_result_title_wrap">
            <p :class="searchResultDetailWrap" style="margin-bottom:0;">
                <span class="search_result_detail_title">Search Results</span>
                <span class="search_result_detail_wrap_hash_var">
                    <span class="title_for" v-show="flshowTitle">for </span>
                    <span>{{Object.keys(this.$route.query)[0]}}</span></span>
            </p>
        </div>
        <div :class="searchResultDetailWrap">
            <div v-show="!flshowResult">
                <p class="transaction_information_content_title">Block</p>
                <div class="block_content_container">
                    <p  class="block_height_container">
                        <span>Height:</span>
                        <span><router-link :to="`/block/${blockHeight}`">{{blockHeight}}</router-link></span>
                    </p>
                    <p class="block_time_container">
                        <span>Timestamp</span>
                        <span>{{blockTime}}</span>
                    </p>
                    <p class="block_hash_container">
                        <span>Block Hash</span>
                        <span class="block_hash">
              <span>{{blockHash}}</span>
            </span>
                    </p>
                </div>
                <p class="transaction_information_content_title proposal_title">Proposal</p>
                <div class="proposal_content_container">
                    <p class="proposal_id_container">
                        <span>Proposal id :</span>
                        <span><router-link :to="`/ProposalsDetail/${proposalId}`" style="color: var(--bgColor) !important;">{{proposalId}}</router-link></span>
                    </p>
                    <p class="proposal_title_container">
                        <span>Title :</span>
                        <span>{{proposalTitle}}</span>
                    </p>
                    <p class="proposal_type_container">
                        <span>Type :</span>
                        <span>{{proposalType}}</span>
                    </p>
                    <p class="proposal_status_container">
                        <span>Status :</span>
                        <span>{{proposalStatus}}</span>
                    </p>
                    <p class="proposal_time_container">
                        <span>Submit Time :</span>
                        <span>{{proposalTime}}</span>
                    </p>
                </div>
            </div>
            <div class="result_container" v-show="flshowResult">
                <div class="result_content_container">
                    <div class="result_img">
                        <img src="../../assets/resultless.png">
                    </div>
                    <p class="result_title">There is no valid result.</p>
                    <p class="try_info">Try to search with Address, Transaction, Block Number, Proposal ID, HashLock.</p>
                    <div class="back_home_btn" @click="backHome">
                        <span>Back Home</span>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>

	import Tools from '../../util/Tools';
	import Service from "../../service"

	export default {
		name: "searchResult",
		data() {
			return {
				flshowResult: true,
				blockHeight: "",
				blockTime: "",
				blockHash: "",
				proposalId: "",
				proposalTitle: "",
				proposalType: "",
				proposalStatus: "",
				proposalTime: "",
				flshowTitle: false,
			}
		},
		watch:{
			$route(){
				if(this.$route.path === "/searchResult"){
					this.flshowTitle = false;
				}else {
					this.flshowTitle = true;
				}
				if(/^\+?[1-9][0-9]*$/.test(Object.keys(this.$route.query)[0])){
					this.flshowResult = false;
					this.searchResult(Object.keys(this.$route.query)[0])
				}else {
					this.flshowResult = true;
				}
			}
		},
		mounted(){
			if(this.$route.path === "/searchResult"){
				this.flshowTitle = true;
			}else {
				this.flshowTitle = false
			}
			if(/^\+?[1-9][0-9]*$/.test(Object.keys(this.$route.query)[0])){
				this.flshowResult = false;
				this.searchResult(Object.keys(this.$route.query)[0])
			}else {
				this.flshowResult = true;
			}
		},
		methods:{
			backHome(){
				this.$router.push("/home")
			},
			searchResult(searchValue){
				let that = this;
				Service.commonInterface({searchResult:{searchValue:searchValue}}, (searchResult) => {
					try {
						that.flshowResult = false;
						if(searchResult){
							Object.entries(searchResult).forEach((item) => {
								let [type, value] = item;
								if(type === "block" && Object.keys(value).length !== 0){
									that.blockHeight = value.height;
									that.blockTime = Tools.format2UTC(value.timestamp);
									that.blockHash = value.hash;
								}else if(type === "proposal" && Object.keys(value).length !== 0){
									that.proposalId = value.proposal_id;
									that.proposalTitle = value.title;
									that.proposalType = value.type;
									that.proposalStatus = value.status;
									that.proposalTime = value.submit_time && Tools.format2UTC(value.submit_time)
								}else {
									this.flshowResult = true;
                                }
							})
						}else {
							that.flshowResult = true;
						}
					}catch (e) {
                        console.error(e)
					}
				})
			}
		},
		beforeMount() {
			if (Tools.currentDeviceIsPersonComputer()) {
				this.searchResultDetailWrap = 'personal_computer_search_result_detail_wrap';
			} else {
				this.searchResultDetailWrap = 'mobile_search_result_detail_wrap';
			}
		},

	}
</script>

<style lang="scss" scoped>
    @import "../../style/mixin";
    .search_result_detail_wrap{
        width: 100%;
        .search_result_title_wrap{
            width: 100%;
            background: #F5F7FD;
            box-sizing: border-box;
            padding: 0 0.2rem;

            .personal_computer_search_result_detail_wrap{
                max-width: 12.8rem;
                margin: 0 auto;
                display: flex;
                align-items: center;
                .search_result_detail_title{
                    font-size: 0.22rem;
                    padding-left: 0.2rem;
                    color: #000;
                    margin-right: 0.2rem;
                    white-space:nowrap;
                }
                .search_result_detail_wrap_hash_var{
                    font-size: 0.22rem;
                    color: var(--contentColor);
                    word-break: break-all;
                    word-wrap: break-word;
                    display: flex;
                    align-items: center;
                    span:first-child{
                        white-space:nowrap;
                    }
                }
            }
            .mobile_search_result_detail_wrap{
                max-width: 12.8rem;
                padding: 0.2rem 0;
                margin: 0 auto;
                display: flex;
                align-items: center;
                .search_result_detail_title{
                    white-space: nowrap;
                }
                .search_result_detail_wrap_hash_var{
                    white-space: nowrap;
                    span:last-child{
                        word-break: break-all;
                        word-wrap: break-word;
                        white-space: pre-wrap;
                    }
                }
            }
        }
        .result_content_container{
            width: 100%;
            text-align: center;
            margin-top: 1.1rem;
            .result_img{
                margin: 0 auto;
                width: 1rem;
                height: 1rem;
                img{
                    width: 100%;
                }
            }
            .result_title{
                margin-top: 0.2rem;
                color: #000;
                font-size: 0.18rem;
            }
            .try_info{
                font-size: 0.14rem;
                color: var(--contentColor);
                margin-bottom: 0.44rem !important;
            }
            .back_home_btn{
                width: 1.58rem;
                height: 0.36rem;
                margin: 0 auto;
                background: var(--bgColor);
                border-radius: 0.05rem;
                color: #fff;
                font-size: 0.14rem;
                line-height: 0.36rem;
            }
        }
        .personal_computer_search_result_detail_wrap{
            max-width: 12.8rem;
            margin: 0 auto;
            margin-top: 0.27rem;
            padding: 0.14rem 0;
            .transaction_information_content_title{
                padding-left: 0.2rem;
                padding-bottom: 0.2rem;
                border-bottom: 0.01rem solid #d7d9e0;
            }
        }
    }
    .block_content_container{
        margin-left: 0.2rem;
        .block_height_container{
            margin-top: 0.2rem;
            line-height: 1;
            span:nth-child(1){
                display: inline-block;
                text-align: left;
                width: 1.3rem;
                @include fontSize;
                color: #787c99;
            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
                cursor: pointer;
                a{
                    color: var(--bgColor)!important;
                }
            }
        }
        .block_time_container{
            margin-top: 0.08rem;
            span:nth-child(1){
                display: inline-block;
                line-height: 1;
                text-align: left;
                width: 1.3rem;
                @include fontSize;
                color: #787c99;
            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
            }
        }
        .block_hash_container{
            line-height: 1;
            margin-top: 0.08rem;
            span:nth-child(1){
                display: inline-block;
                text-align: left;
                width: 1.3rem;
                @include fontSize;
                color: #787c99;
            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
            }
        }
    }
    .proposal_content_container{
        margin-top: 0.21rem;
        margin-left: 0.2rem;
        .proposal_id_container{
            line-height: 1;
            span:nth-child(1){
                display: inline-block;
                width: 1.3rem;
                text-align: left;
                @include fontSize;
            }
            span:nth-child(2){
                color: var(--bgColor);
                cursor: pointer;
            }
        }
        .proposal_title_container{
            line-height: 1;
            margin-top: 0.08rem;
            span:nth-child(1){
                display: inline-block;
                line-height: 1;
                text-align: left;
                width: 1.3rem;
                @include fontSize;

            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
            }
        }
        .proposal_type_container{
            line-height: 1;
            margin-top: 0.08rem;
            span:nth-child(1){
                display: inline-block;
                line-height: 1;
                text-align: left;
                width: 1.3rem;
                @include fontSize;

            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
            }
        }
        .proposal_status_container{
            line-height: 1;
            margin-top: 0.08rem;
            span:nth-child(1){
                display: inline-block;
                line-height: 1;
                text-align: left;
                width: 1.3rem;
                @include fontSize;

            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
            }
        }
        .proposal_time_container{
            line-height: 1;
            margin-top: 0.08rem;
            span:nth-child(1){
                display: inline-block;
                line-height: 1;
                text-align: left;
                width: 1.3rem;
                @include fontSize;

            }
            span:nth-child(2){
                text-align: left;
                color: var(--titleColor);
            }
        }
    }
    .proposal_title{
        margin-top: 0.27rem;
    }
    .mobile_search_result_detail_wrap{
        .search_result_detail_title{
            margin-right: 0.1rem;
        }
        .transaction_information_content_title{
            padding-left: 0.2rem;
        }
        .proposal_content_container{
            margin-bottom: 0.2rem;
        }
        .result_content_container{
            margin-bottom: 1.6rem !important;
        }
        .block_content_container{
            .block_hash_container{
                .block_hash{
                    color: var(--titleColor);
                    width: 100%;
                    display: inline-block;
                    overflow-x: auto;
                    overflow-y: hidden;
                    webkit-overflow-scrolling: touch;
                    &::-webkit-scrollbar {
                        display: none;
                    }
                }
            }
        }
    }
    .title_for{
        padding-right: 0.05rem;
    }
</style>
