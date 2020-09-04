<template>
    <div class="service_respond_record_container">
        <div class="service_respond_record_content_wrap">
            <p class="service_respond_record_title">
                {{$t('ExplorerCN.serviceDetail.respondRecord')}}
                {{$route.params.serviceName}}
                <span class="service_respond_record_spread">
                    |
                </span>
                <span class="service_respond_record_provider">
                    {{$t('ExplorerCN.serviceDetail.provider')}}
                </span>
                <span class="service_respond_record_provider_content">
                    <router-link :to="`/address/${$route.params.provider}`">
                        {{$route.params.provider}}
                    </router-link>
                </span>

            </p>
            <div class="service_respond_record_definition_content">
                <h3 class="service_respond_record_definition_title">
                    {{$t('ExplorerCN.serviceDetail.primary')}}
                </h3>
                <div class="service_respond_record_content">
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.available')}}:</span>
                        <span>{{isAvailable}}</span>
                    </p>
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.pricing')}}:</span>
                        <span>{{price}}</span>
                    </p>
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.qos')}}:</span>
                        <span>{{`${qos} ${$t('ExplorerCN.unit.blocks')}`}} </span>
                    </p>
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.deposit')}}:</span>
                        <span>{{deposit}}</span>
                    </p>
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.hash')}}:</span>
                        <span>
                             <router-link :to="`/tx?txHash=${hash}`">
                                {{hash}}
                            </router-link>
                        </span>
                    </p>
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.owner')}}:</span>
                        <span>
                            <router-link :to="`/address/${owner}`">
                                {{owner}}
                            </router-link>
                        </span>
                    </p>
                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.bindTime')}}:</span>
                        <span>{{bindTime}}</span>
                    </p>

                    <p class="service_respond_record_text_content">
                        <span>{{$t('ExplorerCN.serviceDetail.serviceBindings.disabledTime')}}:</span>
                        <span>{{disabledTime}}</span>
                    </p>

                </div>
            </div>
            <div class="service_respond_record_transaction_content">
                <h3 class="service_respond_record_transaction_title">
                    {{$t('ExplorerCN.serviceDetail.txRecord')}}
                </h3>
                <div class="service_respond_record_transaction_table_content">
                    <el-table :data="txList" :empty-text="$t('ExplorerCN.element.table.emptyDescription')">
                        <el-table-column :min-width="ColumnMinWidth.txHash" :label="$t('ExplorerCN.serviceDetail.respondHash')">
                            <template slot-scope="scope">
                                <img class="service_tx_status"
                                     v-if="scope.row.respondStatus === TX_STATUS.success"
                                     src="../../assets/iService/success.png"/>
                                <img class="service_tx_status"
                                     v-else
                                     src="../../assets/iService/failed.png"/>
                                <el-tooltip placement="top" :content="scope.row.respondHash">
                                    <router-link :to="`/tx?txHash=${scope.row.respondHash}`">
                                        {{formatTxHash(scope.row.respondHash)}}
                                    </router-link>
                                </el-tooltip>

                            </template>
                        </el-table-column>
                        <el-table-column :min-width="ColumnMinWidth.txType" :label="$t('ExplorerCN.transactions.type')"
                                         prop="type"></el-table-column>

                        <el-table-column :min-width="ColumnMinWidth.requestId" :label="$t('ExplorerCN.transactions.requestId')">
                            <template slot-scope="scope">
                                <el-tooltip placement="top" :content="scope.row.requestContextId"
                                            v-if="scope.row.requestContextId">
                                    <span>
                                        {{formatAddress(scope.row.requestContextId)}}
                                    </span>
                                </el-tooltip>
                                <span v-else>--</span>
                            </template>
                        </el-table-column>

                        <el-table-column :min-width="ColumnMinWidth.blockHeight" :label="$t('ExplorerCN.transactions.block')">
                            <template slot-scope="scope">
                                <router-link :to="`/block/${scope.row.height}`">
                                    {{scope.row.height}}
                                </router-link>
                            </template>
                        </el-table-column>

                        <el-table-column :min-width="ColumnMinWidth.time" :label="$t('ExplorerCN.transactions.timestamp')"
                                         prop="time"></el-table-column>
                        <el-table-column :min-width="ColumnMinWidth.address" :label="$t('ExplorerCN.serviceDetail.consumer')">
                            <template slot-scope="scope">
                                <el-tooltip placement="top" :content="scope.row.consumer">
                                    <router-link :to="`/address/${scope.row.consumer}`">{{formatAddress(scope.row.consumer)}}
                                    </router-link>
                                </el-tooltip>
                            </template>
                        </el-table-column>

                        <el-table-column :min-width="ColumnMinWidth.txHash"
                                         :label="$t('ExplorerCN.serviceDetail.requestHash')">
                                <template slot-scope="scope">
                                    <img class="service_tx_status"
                                         src="../../assets/iService/success.png"/>

                                    <el-tooltip placement="top" :content="scope.row.requestHash">
                                        <router-link :to="`/tx?txHash=${scope.row.requestHash}`">
                                            {{formatTxHash(scope.row.requestHash)}}
                                        </router-link>
                                    </el-tooltip>
                                </template>
                        </el-table-column>
                    </el-table>
                </div>
                <div class="pagination_content" v-show="txCount > txPageSize">
                    <keep-alive>
                        <m-pagination :page-size="txPageSize"
                                      :total="txCount"
                                      :page="txPageNum"
                                      :page-change="pageChange">
                        </m-pagination>
                    </keep-alive>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
    import MPagination from "../commontables/MPagination";
    import {
        getRespondServiceRecord,
        getServiceRespondInfo,
        getServiceBindingByServiceName,
    } from "../../service/api";
    import { TX_STATUS,ColumnMinWidth } from '../../constant/Constant';
    export default {
        name : "ServiceInformation",
        components : {MPagination},
        data(){
            return {
                TX_STATUS,
                ColumnMinWidth,
                txList : [],
                txPageSize : 10,
                txPageNum : 1,
                txCount : 0,
                isAvailable : '',
                price : '',
                qos : '',
                deposit : '',
                hash : '',
                owner : '',
                bindTime : '',
                disabledTime : '',
                providerPageSize : 10,
                providerCount : 0,
                providerPageNum : 1,
                Tools,

            }
        },
        mounted(){
            this.getServiceInfo();
            this.getRespondTxList()
        },
        methods : {
            pageChange(pageNum){
                this.txPageNum = pageNum;
                this.getRespondTxList();
            },
            async getServiceInfo(){
                try {
                    const serviceInfo = await getServiceRespondInfo(this.$route.params.serviceName, this.$route.params.provider);
                    const bindings = await getServiceBindingByServiceName(this.$route.params.serviceName, this.$route.params.provider);
                    if(serviceInfo){
                        this.hash = serviceInfo.hash;
                        this.owner = serviceInfo.owner;
                        this.bindTime = Tools.getDisplayDate(serviceInfo.time);
                    }
                    if(bindings && bindings.length){
                        const {available, qos, deposit, disabled_time} = bindings[0];
                        
                        const  copyPricing = {
                            amount: Tools.formatAccountCoinsAmount(JSON.parse(bindings[0].pricing || '{}').price)[0],
                            denom: Tools.formatAccountCoinsDenom(JSON.parse(bindings[0].pricing || '{}').price)[0]
                        }
                        this.isAvailable = available ? 'True' : 'False';
                        this.price = `${Tools.formatAmount3(copyPricing).amount} ${this.$store.state.displayToken}`;
                        this.qos = qos;
                        this.deposit = `${Tools.formatAmount2(deposit[0])}`;
                        this.disabledTime = available ? '--' : Tools.getFormatDate(disabled_time);
                    }

                } catch (e) {
                    console.error(e);
                    this.$message.error(this.$t('ExplorerCN.message.serviceInfoFailed'));
                }
            },

            async getRespondTxList(){
                try {
                    const res = await getRespondServiceRecord(
                        this.$route.params.serviceName,
                        this.$route.params.provider,
                        this.txPageNum,
                        this.txPageSize
                    );
                    console.log(res);
                    this.txList = res.data.map((item) =>{
                        return {
                            type : item.type,
                            respondHash : item.respondHash,
                            requestContextId : item.requestContextId,
                            height : item.height,
                            time : Tools.getDisplayDate(item.time),
                            consumer : item.consumer,
                            requestHash : item.requestHash,
                            respondStatus : item.respondStatus,
                        };
                    });
                    this.txCount = res.count;
                    this.txPageNum = Number(res.pageNum);
                    this.txPageSize = Number(res.pageSize);
                } catch (e) {
                    console.error(e);
                    this.$message.error(this.$t('ExplorerCN.message.txListFailed'));
                }

            },
            formatTxHash(TxHash){
                if(TxHash){
                    return Tools.formatTxHash(TxHash)
                }
            },
            formatAddress(address){
                return Tools.formatValidatorAddress(address)
            },
        }
    }
</script>

<style scoped lang="scss">
    a {
        color: var(--bgColor) !important;
    }
    @media screen and (min-width: 910px){
        .service_respond_record_content_wrap{
            max-width: 12rem;
            .service_respond_record_title {
                padding-left: 0.27rem;
                .service_respond_record_provider {

                    margin-right: 0.15rem;
                }
                .service_respond_record_spread {
                    margin: 0 0.15rem;
                }
            }

        }

    }
    @media screen and (max-width: 910px){
        .service_respond_record_content_wrap{
            width:100%;
            padding:0 0.15rem;
            box-sizing: border-box;
            .service_respond_record_title {
                display:flex;
                flex-direction:column;
                .service_respond_record_provider {
                    margin: 0.1rem 0;
                }
                .service_respond_record_spread {
                    display:none;
                }

            }

        }

    }

    .service_respond_record_container {
        padding: 0 0.15rem;
        .service_respond_record_content_wrap {
            max-width: 12rem;
            margin: 0 auto;
            display: flex;
            flex-direction: column;
            .service_respond_record_title {
                text-align: left;
                margin: 0.42rem 0 0.15rem 0;
                width: 100%;
                box-sizing: border-box;
                font-size: 0.18rem;
                font-weight: 600;
                color: #171D44;
                .service_respond_record_provider {
                    font-size: 0.14rem;
                    margin-right: 0.15rem;
                }
                .service_respond_record_spread {
                    margin: 0 0.15rem;
                }
                .service_respond_record_provider_content {
                    font-size: 0.14rem;

                }

            }
            .service_respond_record_definition_content {
                background: #ffffff;
                box-sizing: border-box;
                padding: 0.25rem 0.27rem 0.2rem 0.27rem;
                margin-bottom: 0.48rem;
                border-radius:5px;
                border:1px solid rgba(215,215,215,1);
                .service_respond_record_definition_title {
                    font-size: 0.18rem;
                    color: #22252A;
                    font-weight: 600;
                    margin-bottom: 0.36rem;
                    text-align: left;
                }
                .service_respond_record_content {
                    box-sizing: border-box;
                    background: #fff;
                    .service_respond_record_text_content {
                        display: flex;
                        justify-content: flex-start;
                        margin-bottom: 0.26rem;
                        span:nth-of-type(1) {
                            font-size: 0.14rem;
                            line-height: 0.16rem;
                            color: #787C99;
                            min-width: 1.5rem;
                            text-align: left;
                            font-weight: 600;
                        }
                        span:last-child {
                            font-size: 0.14rem;
                            line-height: 0.16rem;
                            color: #171D44;
                            flex: 1;
                            text-align: left;
                            word-break: break-all;
                        }
                    }
                    .service_respond_record_text_content:last-child {
                        margin-bottom: 0;
                    }

                }
            }
            .service_respond_record_transaction_content {
                background: #ffffff;
                box-sizing: border-box;
                padding: 0.25rem 0.27rem 0.2rem 0.27rem;
                border-radius:5px;
                border:1px solid rgba(215,215,215,1);
                .service_respond_record_transaction_title {
                    font-size: 0.18rem;
                    color: #22252A;
                    font-weight: 600;
                    margin-bottom: 0.36rem;
                    text-align: left;
                }
                .service_respond_record_transaction_condition_container {
                    width: 100%;
                    display: flex;
                    justify-content: flex-start;
                    margin-bottom: 0.4rem;
                    align-items: center;
                    .service_respond_record_transaction_condition_count {
                        font-size: 0.14rem;
                        margin-right: 0.42rem;
                        font-weight: 600;

                    }
                    /deep/ .el-select {
                        width: 1.3rem;
                        margin-right: 0.22rem;
                        .el-input {
                            .el-input__inner {
                                padding-left: 0.07rem;
                                height: 0.32rem;
                                font-size: 0.14rem !important;
                                line-height: 0.32rem;
                                &::-webkit-input-placeholder {
                                    font-size: 0.14rem !important;
                                }
                            }
                            .el-input__inner:focus {
                                border-color: #3264FD !important;
                            }
                            .el-input__suffix {
                                .el-input__suffix-inner {
                                    .el-input__icon {
                                        line-height: 0.32rem;
                                    }
                                }
                            }
                        }
                        .is-focus {
                            .el-input__inner {
                                border-color: #3264FD !important;
                            }
                        }

                    }
                    .search_btn {
                        cursor: pointer;
                        background: #3264FD;
                        color: #fff;
                        border-radius: 0.04rem;
                        padding: 0.05rem 0.18rem;
                        font-size: 0.14rem;
                        line-height: 0.2rem;
                    }
                }
                .service_respond_record_transaction_table_content {
                    background: #fff;
                    .service_tx_status {
                        position: relative;
                        top: 0.02rem;
                        width:0.13rem;
                        height:0.13rem;
                    }
                    .service_tx_to_container {

                        .service_tx_muti_to_container {
                            display: flex;
                            flex-direction: column;
                            align-items: flex-start;
                            .service_tx_muti_to_ellipsis {
                                color: #3264FD;
                            }
                        }
                    }
                }
            }
            .pagination_content {
                display: flex;
                justify-content: flex-end;
                margin-top: 0.25rem;
                //background: #ffffff;
            }
        }
    }

    @media screen and (max-width: 768px) {
        .service_respond_record_container {

            .service_respond_record_content_wrap {

                .service_respond_record_definition_content {
                    .service_respond_record_definition_title {

                    }
                    .service_respond_record_content {

                        .service_respond_record_text_content {

                            span:nth-of-type(1) {
                                min-width: 1rem;
                            }
                            span:last-child {

                            }
                        }
                        .service_respond_record_text_content:last-child {

                        }
                    }
                }
                .service_respond_record_transaction_content {
                    .service_respond_record_transaction_title {

                    }
                    .service_respond_record_transaction_table_content {

                    }
                }
                .pagination_content {

                }
            }
        }
    }
</style>
