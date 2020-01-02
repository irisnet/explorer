<template>
   <div>
       <m-table :columns="fields" :data="items">
           <template slot-scope="{ row }" slot="txHash">
               <div class="skip_route" v-if="row.txHash">
                   <router-link :to="`/tx?txHash=${row.txHash}`" style="font-family: Consolas,Menlo ;" class="link_style common_font_style">{{formatTxHash(row.txHash)}}
                   </router-link>
               </div>
           </template>
           <template slot-scope="{ row }" slot="block">
                <span class="skip_route">
                    <router-link :to="`/block/${row.block}`" class="link_style">{{row.block}}</router-link>
                </span>
           </template>
           <template slot-scope="{ row }" slot="signer">
               <div class="skip_route" v-if="row.signer">
                   <router-link :to="addressRoute(row.signer)" style="font-family: Consolas,Menlo ;" class="link_style common_font_style">{{formatAddress(row.signer)}}
                   </router-link>
               </div>
           </template>
        <!--   <template slot-scope="{ row }" slot="from">
               <div class="skip_route" v-if="row.from">
                   <router-link v-if="row.from !== '&#45;&#45;'" :to="addressRoute(row.from)" style="font-family: Consolas,Menlo ;" class="link_style common_font_style">{{row.fromMoniker || formatAddress(row.from)}}
                   </router-link>
                   <span v-if="row.from === '&#45;&#45;'">&#45;&#45;</span>
               </div>
           </template>
           <template slot-scope="{ row }" slot="to">
               <div class="skip_route" v-if="row.to">
                   <router-link v-if="row.to !== '&#45;&#45;'" :to="addressRoute(row.to)" style="font-family: Consolas,Menlo ;" class="link_style common_font_style">{{row.toMoniker || formatAddress(row.to)}}
                   </router-link>
                   <span v-if="row.to === '&#45;&#45;'">&#45;&#45;</span>
               </div>
           </template>-->
       </m-table>
       <div class="no_data_img_content" v-if="showNoData">
           <img src="../../assets/no_data.svg" >
       </div>
   </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MHtlsTransferTableList",
        props:{
            items:{
                type: Array,
                default: function() {
                    return [];
                }
            },
            showNoData: {
                type: Boolean,
                default: true
            },
        },
        data () {
            return{
                fields:[
                    {
                        title:'TxHash',
                        slot:'txHash',
                        tooltip: true
                    },
                    {
                        title:'Block',
                        slot:'block'
                    },
               /*     {
                        title:'From',
                        slot:'from',
                        tooltip: true
                    },
                    {
                        title:'Amount',
                        key:'amount'
                    },
                    {
                        title:'To',
                        slot:'to',
                        tooltip: true
                    },*/
                    {
                        title:'Type',
                        key:'type'
                    },
                    {
                        title:'Fee',
                        key:'fee'
                    },
                    {
                        title:'Signer',
                        slot:'signer',
                        tooltip: true
                    },
                    {
                        title:'Status',
                        key:'status'
                    },
                    {
                        title:'Timestamp',
                        key:'timestamp'
                    },
                ]
            }
        },
        methods:{
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
    .no_data_img_content{
        display: flex;
        justify-content: center;
        font-size: 0.14rem;
        height: 1.8rem;
        align-items: center;
        img{
            width: 1.2rem;
        }
    }
</style>
