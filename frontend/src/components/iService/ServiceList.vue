<template>
	<div class="service_list_container_content">
		<div class="service_list_content_wrap">
			<div class="service_list_title">
                {{$t('ExplorerCN.service.services')}}
            </div>
            <div class="service_list_search_wrap">
                <el-input class="search_input" v-model="input" :placeholder="$t('ExplorerCN.service.searchService')"></el-input>
                <div class="tx_type_mobile_content">
                    <div class="search_btn" @click="searchServiceList">{{$t('ExplorerCN.common.search')}}</div>
                    <div class="reset_btn" @click="resetFilterCondition"><i class="iconfont iconzhongzhi"></i></div>
                </div>
            </div>
			<div class="service_list_content" v-for="service in serviceList">
                <div class="service_list_top">
                    <span class="service_list_service_name bold_name">
                        <router-link :to="`/service?serviceName=${service.serviceName}`">
                            {{ service.serviceName }}
                        </router-link>
                    </span>
                    <span class="service_list_service_name">
                        <router-link :to="`/service?serviceName=${service.serviceName}`">
                            {{ service.description }}
                        </router-link>
                    </span>
                </div>
				<el-table :data="service.bindList" :empty-text="$t('ExplorerCN.element.table.emptyDescription')">
					<el-table-column :min-width="ColumnMinWidth.address" :label="$t('ExplorerCN.service.provider')">
						<template slot-scope="scope">
							<span>
                                <el-tooltip placement="top" :content="scope.row.provider">
                                    <router-link :to="`/address/${scope.row.provider}`">
                                        {{Tools.formatValidatorAddress(scope.row.provider)}}
                                    </router-link>
                                </el-tooltip>
							</span>
						</template>
					</el-table-column>
					<el-table-column :min-width="180" :label="$t('ExplorerCN.service.respondTimes')">
						<template slot-scope="scope">
							<span>
								<router-link
                                        :to="`service/respond/${service.serviceName}/${scope.row.provider}`">
                                        {{`${scope.row.respondTimes} ${$t('ExplorerCN.unit.time')}`}}
                                    </router-link>
							</span>
						</template>
					</el-table-column>
					<el-table-column :min-width="ColumnMinWidth.available" :label="$t('ExplorerCN.service.isAvailable')">
                        <template slot-scope="scope">
                        <div class="service_information_available_container">
                            <img class="service_tx_status"
                                 v-if="scope.row.available"
                                 src="../../assets/iService/true.png"/>
                            <img class="service_tx_status"
                                 v-else
                                 src="../../assets/iService/false.png"/>
                            <span>
                                {{scope.row.isAvailable}}
                            </span>
                        </div>

                    </template>
                    </el-table-column>
                    <el-table-column :min-width="ColumnMinWidth.price" :label="$t('ExplorerCN.service.price')" prop="price"></el-table-column>
					<el-table-column :min-width="ColumnMinWidth.qos" :label="$t('ExplorerCN.service.minBlock')" prop="qos"></el-table-column>
					<el-table-column :min-width="ColumnMinWidth.time" :label="$t('ExplorerCN.service.time')" prop="bindTime"></el-table-column>
				</el-table>
			</div>
			<div class="pagination_content" v-if="txCount > pageSize">
				<m-pagination :page-size="pageSize"
				              :total="txCount"
				              :page="pageNum"
				              :page-change="pageChange">
					
				</m-pagination>
			</div>
            <div class="service_list_empty_container" v-if="serviceList.length === 0">
                <img src="../../assets/empty.png" alt="" class="service_list_empty">
                <span class="service_list_empty_description">
                    {{ $t('ExplorerCN.element.table.emptyDescription') }}
                </span>
            </div>
		</div>
	</div>
</template>

<script>
	import Tools from "../../util/Tools"
    import MPagination from "../commontables/MPagination";
    import {getAllServiceTxList, getServiceBindingByServiceName} from "../../service/api";
    import { ColumnMinWidth } from '../../constant/Constant';
    export default {
		name: "ServiceList",
		components: {MPagination},
		data() {
			return {
                ColumnMinWidth,
				pageNum: 1,
				pageSize: 5,
				serviceList:[],
				txCount:0,
                Tools,
                input:'',
			}
		},
		mounted () {
			this.getServiceList();
		},
		methods:{
			async getServiceList(){
                try {
                    let serviceList = await getAllServiceTxList(this.input, this.pageNum,this.pageSize);
                    if(serviceList && serviceList.data){
                        console.log(serviceList)
                        for(let service of serviceList.data){
                            let bindings = [];
                            if (service.serviceName && service.serviceName.length) {
                                bindings = await getServiceBindingByServiceName(service.serviceName);
                            }
                            service.bindList.forEach((s)=>{
                                s.bindTime = Tools.getDisplayDate(s.bindTime);
                                s.price = '--';
                                s.qos = '--';
                                s.isAvailable = 'False';
                                bindings.forEach((b)=>{
                                    if(s.provider === b.provider){
                                        s.isAvailable = b.available ? 'True' : 'False';
                                        s.available = b.available;
                                        s.price = JSON.parse(b.pricing).price;
                                        s.qos = `${b.qos} ${this.$t('ExplorerCN.unit.blocks')}`;
                                    }
                                })
                            })
                        }
                        console.log(serviceList)
                        this.serviceList = serviceList.data;

                        this.txCount = serviceList.count;

                    }
                }catch (e) {
                    console.error(e);
                    this.$message.error(this.$t('ExplorerCN.message.serviceTxListFailed'));
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
			pageChange(pageNum) {
				this.pageNum = pageNum;
				this.getServiceList();
			},
            searchServiceList(){
                this.pageNum = 1;
                this.getServiceList()
            },
            resetFilterCondition(){
                this.input = '';
                this.pageNum = 1;
                this.getServiceList()
            },
		}
	}
</script>

<style scoped lang="scss">
	a{
		color: var(--bgColor) !important;
	}
	.service_list_container_content{
        @media screen and (min-width: 910px){
            .service_list_title{
                // padding-left: 0.27rem;
            }
            .service_list_content_wrap{
                max-width: 12rem;
            }

        }
        @media screen and (max-width: 910px){
            .service_list_content_wrap{
                width:100%;
                padding:0 0.15rem;
                box-sizing: border-box;

            }

        }
		.service_list_content_wrap{
			margin: 0 auto;
            .service_list_empty_container{
                display:flex;
                flex-direction:column;
                align-items: center;
                margin-top:1.32rem;
                .service_list_empty{
                    height:1.78rem;
                    width:1.74rem;
                    margin-bottom:0.18rem;
                }
                .service_list_empty_description{
                    color:#B8B8B8;
                    font-size:0.14rem;

                }
            }

			.service_list_title{
                text-align: left;
                margin: 1rem 0 0.15rem 0.28rem;
                width: 100%;
                box-sizing: border-box;
                font-size: 0.18rem;
                font-weight: 600;
                color: #171D44;
			}
            .service_list_search_wrap{
                display: flex;
                width: 100%;
                margin: 0.25rem 0 0.25rem 0.28rem;
                .search_input{
                    max-width: 3.5rem;
                    .el-input__inner{
                        padding-left: 0.07rem;
                        height: 0.32rem;
                        font-size: 0.14rem !important;
                        line-height: 0.32rem;
                    }
                }
                .tx_type_mobile_content{
                    display: flex;
                    align-items: center;
                    .reset_btn{
                        background: var(--bgColor) !important;
                        color: #fff;
                        border-radius: 0.04rem;
                        margin-left: 0.1rem;
                        cursor: pointer;
                        i{
                            padding: 0.08rem;
                            font-size: 0.14rem;
                            line-height: 1;
                            display: inline-block;
                        }
                    }
                    .search_btn{
                        min-width:0.4rem;
                        text-align:center;
                        cursor: pointer;
                        background: var(--bgColor) !important;
                        margin-left: 0.1rem;
                        color: #fff;
                        border-radius: 0.04rem;
                        padding: 0.05rem 0.18rem;
                        font-size: 0.14rem;
                        line-height: 0.2rem;
                    }
                }
            }
            
            .service_list_content{
                display:flex;
                flex-direction:column;
                margin-bottom:0.48rem;
                padding:0.28rem 0.28rem 0.18rem 0.28rem;
                background: #ffffff;
                border-radius:5px;
                border:1px solid rgba(215,215,215,1);
                .service_information_available_container{
                    display:flex;
                    align-items: center;
                }
                .service_list_top{
                    display:flex;
                    justify-content: flex-start;
                    margin-bottom:0.4rem;
                    .service_list_service_name{
                        color:#3264FD;
                        font-size:0.14rem;
                        cursor:pointer;
                        margin-right:0.15rem;
                    }
                    .bold_name{
                        font-weight:600;
                    }
                }

            }
		}
		.pagination_content{
			display: flex;
			justify-content: flex-end;
			margin: 0.1rem 0 0.2rem 0;
		}
	}
</style>
