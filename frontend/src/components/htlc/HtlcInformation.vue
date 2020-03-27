<template>
    <div class="htlc_information_container">
        <page-title :title="pageTitle" :flShowPageLink="false"></page-title>
        <div class="htlc_information_content">
         <!--   <div class="htlc_information_header_content">
                <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Completed'" src="../../assets/Complete.png">
                <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Open'" src="../../assets/Open.png">
                <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Expired'" src="../../assets/Expired.png">
                <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Refunded'" src="../../assets/Refunded.png">
            </div>-->
            <div class="htlc_information_content_container">
                <ul class="htlc_information_content_wrap">
                    <li class="htlc_information_item"
                        v-for="item in htlcInformationArr">
                        <span class="htlc_information_name">{{Object.keys(item)[0]}}</span>
                        <div class="img_content">
                            <div class="htlc_information_value"
                                 v-if="Object.keys(item)[0] !== 'HTLC Sender:'
                                 && Object.keys(item)[0] !== 'HTLC Receiver:'
                                ">{{Object.values(item)[0]}}
                                <span v-if=" Object.keys(item)[0] === 'State:'">
                                    <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Completed'" src="../../assets/Complete.png">
                                    <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Open'" src="../../assets/Open.png">
                                    <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Expired'" src="../../assets/Expired.png">
                                    <img v-show="htlcInformationArr[htlcInformationArr.length -1]['State:'] === 'Refunded'" src="../../assets/Refunded.png">
                                </span>
                            </div>
                            <div class="skip_route" v-if="Object.keys(item)[0] === 'HTLC Sender:'|| Object.keys(item)[0] === 'HTLC Receiver:'">
                                <router-link v-if="Object.keys(item)[0] === 'HTLC Sender:'" :to="addressRoute(Object.values(item)[0])">{{ monikerFrom !== '' ? monikerFrom : Object.values(item)[0]}}</router-link>
                                <router-link v-if="Object.keys(item)[0] === 'HTLC Receiver:'" :to="addressRoute(Object.values(item)[0])">{{ monikerTo !== '' ? monikerTo : Object.values(item)[0]}}</router-link>
                            </div>
                            <m-clip v-if="Object.keys(item)[0] === 'Hash Lock:'" :text="Object.values(item)[0]" style="margin-left: 0.1rem"></m-clip>
                        </div>
                    </li>
                </ul>
            </div>
            <div class="htlc_information_transfer_list_content">
                <div class="htlc_information_transfer_list_title_content">
                    <span class="htlc_transaction_title">HTLC Transaction</span>
                </div>
                <div class="htlc_information_transfer_table_list_content">
                    <div class="htlv_information_transfer_table_list_wrap">
                        <m-htls-transfer-table-list :items="HtlcTxListArray" :showNoData="false"></m-htls-transfer-table-list>
                    </div>
                    <div class="pagination_content">
                        <m-pagination
                            :page-size="pageSize"
                            :total="countNum"
                            :page="pageNumber"
                            :page-change="pageChange"
                        ></m-pagination>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
	import MHtlsTransferTableList from "./MHtlsTransferTableList";
	import Tools from "../../util/Tools"
    import Server from "../../service"
    import MClip from "../commonComponents/MClip";
    import MPagination from "../commonComponents/MPagination";
    import PageTitle from "../pageTitle/PageTitle";
    export default {
		name: "HtlcInformation",
        components: {PageTitle, MPagination, MClip, MHtlsTransferTableList},
        data () {
		    return {
                HtlcTxListArray: [],
                pageNumber:1,
                pageSize:10,
                countNum:0,
                monikerFrom:'',
                monikerTo:'',
                pageTitle:'HTLC Details',
                htlcInformationArr:[
                    {
                        "Hash Lock:" : '--',
                        key:'hash_lock',
                    },
                    {
                        "HTLC Sender:" : '--',
                        key:'from'
                    },
                    {
                        "Amount:" : '--',
                        key:'amount'
                    },
                    {
                        "HTLC Receiver:" : '--',
                        key:'to'
                    },
                    {
                        "Time Lock:" : '--',
                        key:'time_lock'
                    },
                    {
                        "Timestamp:" : '--',
                        key:'timestamp'
                    },
                    {
                        "Expiry Height:" : '--',
                        key:'expire_height'
                    },
                    {
                        "Cross-chain Receiver:" : '--',
                        key:'cross_chain_receiver'
                    },
                    {
                        "State:" : '--',
                        key:'state'
                    }
                ]
            }
        },
        mounted(){
		  this.getHtlcInformation();
		  this.getHtlcTxList();
        },
        methods:{
            getHtlcInformation(){
                Server.commonInterface({
                    htlcs:{
                        hashLock: this.$route.params.txHash
                    }
                },(res) => {
                    this.htlcInformationArr.forEach( item => {

                        if(Object.values(item)[1] === 'amount'){
                            item[Object.keys(item)[0]] = Tools.formatAmount2(res[item.key])
                        }else {
                            if(Object.keys(item)[0] ===  'State:'){
                                res[item.key] = Tools.firstWordUpperCase(res[item.key])
                            }
                            if(Object.keys(item)[0] ===  'Time Lock:'){
                                res[item.key] = `${res[item.key]} blocks`
                            }
                            if(Object.keys(item)[0] ===  'Timestamp:'){
                                res[item.key] = `${res[item.key]} secs`
                            }
                            item[Object.keys(item)[0]] = res[item.key] || res[item.key] === 0 ? res[item.key] : '--'
                        }
                        this.monikerFrom = res.from_moniker;
                        this.monikerTo = res.to_moniker
                    });
                })
            },
            getHtlcTxList(){
                Server.commonInterface({
                    htlcTxList:{
                        hashLock: this.$route.params.txHash,
                        pageNumber: this.pageNumber,
                        pageSize: this.pageSize
                    }
                },(res) => {
                    try {
                        if(res){
                            this.countNum = res.Count;
                            if(res.Data){
                                this.HtlcTxListArray = res.Data.map( item => {
                                    let fromInformation = Tools.formatListAmount(item).fromAddressAndMoniker,
                                        toInformation = Tools.formatListAmount(item).toAddressAndMoniker;
                                  return{
                                      txHash: item.hash,
                                      block: item.block_height,
                                      from: fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].address : '--',
                                      amount: item.amount.length > 0 ? Tools.formatAmount2(item.amount,2) : '--',
                                      to:toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].address : '--',
                                      type: item.type,
                                      fee: Tools.formatFee(item.fee),
                                      signer: item.signer,
                                      status: Tools.firstWordUpperCase(item.status),
                                      timestamp: Tools.format2UTC(item.timestamp),
                                      fromMoniker:  fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].moniker :'',
                                      toMoniker: toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].moniker :'',
                                  }
                                });
                            }
                        }

                    }catch (e) {
                        console.error(e)
                    }
                })
            },
            pageChange(pageNum){
                this.pageNumber = pageNum;
                this.getHtlcTxList()
            }
        }
    }
</script>

<style scoped lang="scss">
    .htlc_information_container{
       
        .htlc_information_content{
            max-width: 12.8rem;
            margin: 0 auto;
            padding-top: 0.8rem;
            .htlc_information_header_content{
                color: #515a6e;
                font-size: 0.18rem;
                height: 0.7rem;
                line-height: 0.7rem;
                font-weight: bold;
                margin-left: 0.2rem;
                display: flex;
                align-items: center;
                .htlc_information_header_title{
                    padding-right: 0.1rem
                }
                img{
                    width: 0.2rem;

                }
            }
            .htlc_information_content_container{
                background: #fff;
                box-sizing: border-box;
                padding: 0.2rem;
                border: 0.01rem solid #e7e9eb;
                .htlc_information_content_wrap{
                    .htlc_information_item{
                        display: flex;
                        align-items: center;
                        width: 100%;
                        padding: 0.05rem;
                        .htlc_information_name{
                            color: var(--contentColor);
                            font-size: 0.14rem;
                            min-width: 1.58rem;
                            display: inline-block;
                        }
                        .img_content{
                            width: 100%;
                            display: flex;
                            .htlc_information_value{
                                overflow-x: auto;
                                color: #171d44;
                                font-size: 0.14rem;
                                display: flex;
                                span{
                                    display: flex;
                                    align-items: center;
                                    img{
                                        margin-left: 0.1rem;
                                        width: 0.14rem;
                                    }
                                }
                            }
                        }
                    }
                }
            }
            .htlc_information_transfer_list_content{
                .htlc_information_transfer_list_title_content{
                    margin: 0.3rem 0 0.12rem 0.2rem;
                    .htlc_transaction_title{
                        font-size: 0.18rem;
                        color: #515a6e;
                    }
                }
                .htlc_information_transfer_table_list_content{
                    .htlv_information_transfer_table_list_wrap{
                        max-width: 12.8rem;
                        overflow-x: auto;
                        overflow-y: hidden;
                    }
                    .pagination_content{
                        max-width: 12.8rem;
                        display: flex;
                        margin: 0.2rem auto 0 auto;
                        padding-bottom: 0.4rem;
                        justify-content:flex-end;
                    }
                }
            }
        }
    }
    @media screen and (max-width: 910px){
        .htlc_information_container{
            .htlc_information_content{
                padding-top: 0.15rem;
                .htlc_information_header_content{
                    margin-left: 0.1rem;
                }
                .htlc_information_content_container{
                    margin: 0 0.1rem;
                    .htlc_information_content_wrap{
                        .htlc_information_item{
                            display: flex;
                            flex-direction: column;
                            align-items: flex-start;
                        }
                    }
                }
                .htlc_information_transfer_list_content{
                    .htlc_information_transfer_list_title_content{
                        margin: 0.3rem 0 0.12rem 0.1rem;
                    }
                    .htlc_information_transfer_table_list_content{

                        margin: 0 0.1rem;
                    }
                }
            }
        }
    }
</style>
