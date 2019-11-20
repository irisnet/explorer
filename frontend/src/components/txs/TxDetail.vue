<template>
    <div class="tx_detail_container">
        <div class="tx_detail_common_information_content_wrap">
            <div class="tx_detail_header_title">Transaction Information</div>
            <div class="tx_detail_common_information_wrap">
                <ul class="tx_detail_common_information_content">
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">TxHash :</span>
                        <div class="tx_detail_common_information_item_value">
                            <span>{{txHashValue}}</span>
                            <m-clip style="margin-left: 0.1rem" :text="txHashValue"></m-clip>
                        </div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Block :</span>
                        <div class="tx_detail_common_information_item_value skip_route">
                            <router-link :to="`/block/${blockValue}`">{{blockValue}}</router-link>
                        </div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Status :</span>
                        <div class="tx_detail_common_information_item_value" style="display: flex">
                            <span :style="{color: statusValue === 'Failed' ? '#fa8593' : ''}">{{statusValue}}</span>
                            <div v-if="statusValue === 'Failed'" class="log_content_container">
                                <i class="iconfont iconyiwen"></i>
                                <div class="tip_content" :class="failTipStyle ? 'overstep_style' :''" :style="{left:`${-((logContentWidth / 2) - 6) / 100}rem`}">
                                    <span class="log_content"  ref="logContent">{{log}}</span>
                                </div>
                            </div>
                        </div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Timestamp :</span>
                        <div class="tx_detail_common_information_item_value">{{timestampValue}}</div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Fee  :</span>
                        <div class="tx_detail_common_information_item_value">{{feeValue}}</div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Gas Used :</span>
                        <div class="tx_detail_common_information_item_value gas_container">
                            <span>{{gasUsedValue}}</span>
                            <div class="gas_content_container">
                                <i class="iconfont icontishi"></i>
                                <div class="gas_tip_content_wrap"  :style="{left:`${-((gasContentWidth / 2) - 7) / 100}rem`}">
                                    <div class="gas_tip_content">
                                        <p>Gas Price : {{gasPrice ? gasPrice + ' Nano' : '--'}}</p>
                                        <p>Gas Used : {{gasUsed || '--'}}</p>
                                        <p>Gas Wanted : {{gasWanted || '--'}}</p>
                                        <p>Gas Limit : {{gasLimit || '--'}}</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Signer :</span>
                        <div class="tx_detail_common_information_item_value skip_route iconProfiler_content">
                            <router-link :to="addressRoute(signerValue)">{{signerValue}}</router-link>
                            <span class="address_information_address_status_profiler" v-if="isProfiler">Profiler</span>
                        </div>
                    </li>
                    <li class="tx_detail_common_information_item">
                        <span class="tx_detail_common_information_item_name">Memo :</span>
                        <div class="tx_detail_common_information_item_value">{{memoValue}}</div>
                    </li>
                </ul>
            </div>
            <div class="tx_detail_message_content_title">Transaction Message</div>
            <div class="tx_detail_message_information_content">
                <ul class="tx_detail_message_information">
                    <li class="tx_detail_message_information_item" v-for="(item,key) in messageList">
                        <span class="tx_detail_message_information_name" v-if="key !=='tooltip'">{{key}}</span>
                        <div>
                            <div class="tx_detail_message_information_value" :class="key !== 'Commission Rate :'? 'hide_rate_tip' : ''" v-for="value in item">
                            <span v-if="key !== 'From :'
                            && key !== 'To :'
                            && key !== 'Owner :'
                            && key !== 'Operator Address :'
                            && key !== 'Owner Address :'
                            && key !== 'Original Address :'
                            && key !== 'New Address :'
                            && key !== 'Proposer :'
                            && key !== 'Depositor :'
                            && key !== 'Proposal ID :'
                            && key !== 'Voter :'
                            && key !== 'Original Owner :'
                            && key !== 'New Owner :'
                            && key !== 'Website :'
                            && key !== 'tooltip'
                            && key !== 'Identity :'
                            && key !== 'Software :'
                            && key !== 'Address :'
                            && key !== 'Added By :'
                            && key !== 'Deleted By :'
                            && key !== 'Hash Lock :'
                            && key !== 'DestAddress :'
                            && key !== 'Symbol :'">{{value}}</span>
                                <span class="skip_route"
                                      v-if="key === 'From :'
                                || key === 'To :'
                                || key === 'Owner :'
                                || key === 'Operator Address :'
                                || key === 'Owner Address :'
                                || key === 'Original Address :'
                                || key === 'New Address :'
                                || key === 'Proposer :'
                                || key === 'Depositor :'
                                || key === 'Proposal ID :'
                                || key === 'Voter :'
                                || key === 'Original Owner :'
                                || key === 'New Owner :'
                                || key === 'Website :'
                                || key === 'tooltip'
                                || key === 'Identity :'
                                || key === 'Software :'
                                || key === 'Address :'
                                || key === 'Added By :'
                                || key === 'Deleted By :'
                                || key === 'Hash Lock :'
                                || key === 'DestAddress :'
                                || key === 'Symbol :'">
                                <router-link v-if="key === 'Symbol :' && value !== '-'" :to="`asset/${value}`">{{value}}</router-link>
                                <router-link v-if="key === 'DestAddress :' && value !== '-'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'From :' && value !== '-'" :to="addressRoute(value)">{{fromMoniker || value}}</router-link>
                                <router-link v-if="key === 'To :' && value !== '-'" :to="addressRoute(value)">{{toMoniker || value}}</router-link>
                                <router-link v-if="key === 'Owner :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Operator Address :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Owner Address :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Original Address :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'New Address :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Proposer :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Depositor :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Proposal ID :'" :to="`ProposalsDetail/${value}`">{{value}}</router-link>
                                <router-link v-if="key === 'Voter :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Original Owner :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'New Owner :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Address :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Added By :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Deleted By :'" :to="addressRoute(value)">{{value}}</router-link>
                                <router-link v-if="key === 'Hash Lock :' && messageList['TxType :'][0] !== 'CreateHTLC'" :to="`/htlc/${value}`">{{value}}</router-link>
                                <router-link v-if="key === 'Hash Lock :' && messageList['TxType :'][0] === 'CreateHTLC' && statusValue === 'Success'" :to="`/htlc/${value}`">{{value}}</router-link>
                                <span v-if="key === 'Hash Lock :' && statusValue === 'Failed' && messageList['TxType :'][0] === 'CreateHTLC'">{{value}}</span>
                                <span v-if="key === 'Website :'" @click="openUrl(value)" style="color: var(--bgColor);cursor: pointer;">{{value}}</span>
                                <span v-if="key === 'Software :'" @click="openUrl(value)" style="color: var(--bgColor);cursor: pointer;">{{value}}</span>
                                <span v-if="key === 'Identity :'" @click="getKeyBaseName(value)" style="color: var(--bgColor);cursor: pointer;">{{value}}</span>
                                <span v-if="key === 'From :' && value === '-'">--</span>
                                <span v-if="key === 'To :' && value === '-'">--</span>
                                <span v-if="key === 'DestAddress :' && value === '-'">--</span>

                            </span>
                                <div class="commission_rate_container" v-if="key === 'Commission Rate :' && messageList['TxType :'][0] === 'CreateValidator' && value !== '--'">
                                    <i class="iconfont icontishi"></i>
                                    <div class="commission_rate_content_wrap" :style="{left:`${-((rateContentWidth / 2) - 7) / 100}rem`}">
                                        <div class="commission_rate_content">
                                            <p>Max Rate : {{messageList.tooltip.maxRate || '--'}}</p>
                                            <p>Max Change Rate : {{messageList.tooltip.maxChangeRate || '--'}}</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
    import Server from "../../service";
    import MClip from "../commonComponents/MClip";
    import Tools from "../../util/Tools";
    import formatMessage from "../../util/txdetail"
    import axios from "../../util/axios"
	export default {
		name: "TxDetail",
        components: {MClip},
        data () {
		    return {
                txHashValue: '',
                blockValue: '',
                statusValue: '',
                timestampValue: '',
                feeValue: '',
                gasUsedValue: '',
                signerValue: '',
                memoValue: '',
                messageList: null,
                fromMoniker:'',
                toMoniker:'',
                log:'',
                gasPrice:'',
                gasLimit:'',
                gasUsed:'',
                gasWanted:'',
                logContentWidth: '',
                gasContentWidth:'',
                rateContentWidth:'',
                flShowRateToolTip: false,
                isProfiler:false,
                failTipStyle:false,
            }
        },
        watch:{
            log(){
              this.$nextTick( () => {
                  let toolTipMaxWidth = 400;
                  this.logContentWidth = document.getElementsByClassName('tip_content')[0].clientWidth;
                  if(this.logContentWidth > toolTipMaxWidth){
                    this.failTipStyle = true;
                  }
              })
            },
            gasPrice(){
                this.$nextTick( () => {
                    this.gasContentWidth = document.getElementsByClassName('gas_tip_content_wrap')[0].clientWidth;
                })
            },
            messageList(){
                if(this.flShowRateToolTip){
                    this.$nextTick(() => {
                        this.rateContentWidth = document.getElementsByClassName('commission_rate_content_wrap')[0].clientWidth;
                    })
                }
            }
        },
        mounted(){
		    this.getTxDetailInformation()


        },
        methods:{
            openUrl(url) {
                url = url.trim();
                if (url) {
                    if (!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)) {
                        url = `http://${url}`;
                    }
                    window.open(url);
                }
            },
            getKeyBaseName(identity) {
                if(identity !== '--'){
                    let url = `https://keybase.io/_/api/1.0/user/lookup.json?fields=basics&key_suffix=${identity}`;
                    if (identity) {
                        axios.http(url).then(res => {
                            if (res.them && res.them[0].basics.username) {
                                window.open(`https://keybase.io/${res.them[0].basics.username}`)
                            }
                        });
                    }
                }
             },
            getTxDetailInformation(){
                Server.commonInterface( {txDetail: {txHash: this.$route.query.txHash} },(res) => {
                    try {
                        if(res){
                            this.gasPrice = Tools.convertScientificNotation2Number(
                                Tools.formaNumberAboutGasPrice(res.gas_price)
                            );
                            res.isProfiler = false;
                            if(res.isProfiler){
                                this.isProfiler = res.isProfiler;
                            }
                            this.gasLimit = res.gas_limit;
                            this.gasUsed = res.gas_used;
                            this.gasWanted = res.gas_wanted;
                            this.txHashValue = res.hash;
                            this.log = res.log;
                            this.blockValue = res.block_height;
                            this.statusValue = Tools.firstWordUpperCase(res.status);
                            this.timestampValue = Tools.format2UTC(res.timestamp);
                            this.feeValue = Tools.formatFee(res.fee);
                            this.gasUsedValue = res.gas_used;
                            this.signerValue = res.signer;
                            this.memoValue = res.memo ? res.memo : '--';
                            this.fromMoniker = res.from_moniker;
                            this.toMoniker = res.to_moniker;
                            this.messageList = formatMessage.switchTxType(res);
                            if(this.messageList.tooltip){
                                this.flShowRateToolTip = true
                            }
                        }

                    }catch (e) {
                        console.error(e)
                    }
                })
            }
        }
	}
</script>

<style scoped lang="scss">
    .tx_detail_container{
        .tx_detail_common_information_content_wrap{
            max-width: 12.8rem;
            margin: 0 auto;
            .tx_detail_header_title{
                height: 0.7rem;
                line-height: 0.7rem;
                color:#515a6e;
                margin-left: 0.2rem;
                font-weight: bold;
                font-size: 0.18rem;
            }
            .tx_detail_common_information_wrap{
                background: #fff;
                border:0.01rem solid #d7d9e0;
                box-sizing: border-box;
                padding: 0.2rem;
                font-size: 0.14rem;
                .tx_detail_common_information_content{
                    .tx_detail_common_information_item{
                        padding: 0.05rem 0;
                        display: flex;
                        .tx_detail_common_information_item_name{
                            width: 1.5rem;
                            color: var(--contentColor);
                        }
                        .tx_detail_common_information_item_value{
                            display: flex;
                            color: var(--titleColor);
                            span{
                                color: var(--titleColor);
                                overflow-x: auto;
                                overflow-y: hidden;
                            }
                            .log_content_container{
                                position: relative;
                                .iconyiwen{
                                    padding-left: 0.05rem;
                                    font-size: 0.14rem;
                                    color: #fa8593;
                                }
                                &:hover{
                                    .tip_content{
                                        visibility: visible;
                                    }
                                }
                                .tip_content{
                                    position: absolute;
                                    top:-0.33rem;
                                    visibility: hidden;
                                    .log_content{
                                        white-space: nowrap;
                                        bottom: 0.25rem;
                                        word-break: normal;
                                        background: #000;
                                        color:#fff;
                                        padding: 0.08rem 0.15rem;
                                        border-radius: 0.04rem;
                                        &::after{
                                            position: absolute;
                                            width: 0;
                                            height: 0;
                                            border: 0.1rem solid transparent;
                                            content: "";
                                            display: block;
                                            border-top-color: #000000;
                                            left: 50%;
                                            margin-left: -0.06rem;
                                        }
                                    }

                                }
                                .overstep_style{
                                    left: -1.5rem!important;
                                    .log_content{
                                        &::after{
                                            left: 1.5rem !important;
                                            margin-left: 0 !important;
                                        }
                                    }
                                }
                            }
                            .address_information_address_status_profiler{
                                background: var(--bgColor);
                                min-width: 0.65rem;
                                font-size: 0.12rem;
                                color: #fff;
                                padding: 0.02rem 0.14rem;
                                border-radius: 0.22rem;
                                margin-left: 0.05rem;
                            }
                        }
                        .iconProfiler_content{
                            overflow-x: auto;
                            overflow-y: hidden;
                        }
                        .gas_container{
                            .gas_content_container{
                                position: relative;
                                &:hover{
                                    .gas_tip_content_wrap{
                                        visibility: visible;
                                    }
                                }
                                .icontishi{
                                    padding-left: 0.05rem;
                                    font-size: 0.14rem;
                                    color:var(--bgColor);

                                }
                                .gas_tip_content_wrap{
                                    position: absolute;
                                    bottom: 0.25rem;
                                    left: -1.2rem;
                                    visibility: hidden;
                                    .gas_tip_content{
                                        background: #000;
                                        color: #fff;
                                        padding: 0.09rem 0.15rem;
                                        border-radius: 0.04rem;
                                        p{
                                            display: flex;
                                            white-space: nowrap;
                                        }
                                        &::after{
                                            position: absolute;
                                            width: 0;
                                            height: 0;
                                            border: 0.14rem solid transparent;
                                            content: "";
                                            display: block;
                                            border-top-color: #000000;
                                            left: 50%;
                                            margin-left: -0.08rem;
                                        }
                                    }
                                }
                            }
                        }

                    }
                }
            }
            .tx_detail_message_content_title{
                font-size: 0.18rem;
                height: 0.7rem;
                line-height: 0.7rem;
                margin-left: 0.2rem;
                color: #171d44;
            }
            .tx_detail_message_information_content{
                padding-bottom: 0.4rem;
                .tx_detail_message_information{
                    box-sizing: border-box;
                    padding: 0.2rem;
                    background: #fff;
                    border: 0.01rem solid #d7d9e0;
                    .tx_detail_message_information_item{
                        padding: 0.05rem 0;
                        display: flex;
                        font-size: 0.14rem;
                        .tx_detail_message_information_name{
                            width: 1.5rem;
                            min-width: 1.5rem;
                            max-width: 1.5rem;
                            color: var(--contentColor);
                        }
                        .hide_rate_tip{
                            overflow-x: auto;
                            overflow-y: hidden;
                        }
                        .tx_detail_message_information_value{
                            display: flex;
                            color: var(--titleColor);

                            .commission_rate_container{
                                position: relative;
                                display: flex;
                                i{
                                    font-size: 0.14rem;
                                    padding-left: 0.05rem;
                                    color:var(--bgColor);
                                }
                                &:hover{
                                    .commission_rate_content_wrap{
                                        visibility: visible;
                                    }
                                }
                                .commission_rate_content_wrap{
                                    background: #000;
                                    color: #fff;
                                    padding: 0.08rem 0.15rem;
                                    border-radius: 0.04rem;
                                    position: absolute;
                                    bottom:0.3rem;
                                    visibility: hidden;
                                    .commission_rate_content{
                                        p{
                                            display: flex;
                                            white-space: nowrap;
                                        }
                                        &::after{
                                            position: absolute;
                                            width: 0;
                                            height: 0;
                                            border: 0.14rem solid transparent;
                                            content: "";
                                            display: block;
                                            border-top-color: #000000;
                                            left: 50%;
                                            margin-left: -0.08rem;
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
    @media screen and (max-width: 910px){
        .tx_detail_container{
            .tx_detail_common_information_content_wrap{
                .tx_detail_common_information_wrap{
                    margin: 0 0.1rem;
                    .tx_detail_common_information_content{
                        .tx_detail_common_information_item{
                            flex-direction: column;
                            .tx_detail_common_information_item_value{
                                display: flex;
                                .log_content_container{
                                    .tip_content{
                                        width: 3rem;
                                        left: -0.12rem !important;
                                        .log_content{
                                            display: inline-block;
                                            overflow-x: auto;
                                            overflow-y: hidden;
                                            padding: 0.04rem 0.1rem !important;
                                            width: 100%;
                                            &::after{
                                                left: 0.2rem;
                                            }
                                        }
                                    }
                                }
                            }
                            .gas_container{
                                .gas_content_container{
                                    .gas_tip_content_wrap{
                                        left: -0.2rem !important;
                                        .gas_tip_content{
                                            &::after{
                                                left: 0.25rem;
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
                .tx_detail_message_information_content{
                    margin: 0 0.1rem;
                    padding-bottom: 0.4rem;
                    .tx_detail_message_information{
                        .tx_detail_message_information_item{
                            flex-direction: column;
                            padding: 0;
                            .tx_detail_message_information_value{
                                .commission_rate_container{
                                    .commission_rate_content_wrap{
                                        left: -0.2rem !important;
                                        .commission_rate_content{
                                            &::after{
                                                left: 0.25rem;
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
