<template>
    <div :class="showNoData?'show_no_data':''" :style="`${minWidth?(`min-width:${minWidth}rem`):''}`">
      <b-table :fields='validatorFields' :status="status" :items='items' striped class="show_trim" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc">
        <template slot="index" slot-scope="data" small>
          {{data.index + 1}}
        </template>
        <template slot='moniker' slot-scope='data'>
        <span class="skip_route" style="display: flex">
          <div style="width: 0.3rem;height: 0.3rem;border-radius: 0.3rem;overflow: hidden" v-if="data.item.url">
            <img style="width: 100%;" :src="data.item.url ? data.item.url : ''">
          </div>
          <div class="name_address" style="margin-left:0.2rem;">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${data.item.operatorAddress}`" class="link_style">{{data.item.moniker}}</router-link>
            </span>
          </div>
        </span>
        </template>
        <template slot='operatorAddress' slot-scope='data'>
        <span class="skip_route" style="display: flex" v-if="data.item.operatorAddress">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${data.item.operatorAddress}`" class="link_style">{{formatAddress(data.item.operatorAddress)}}</router-link>
            </span>
            <span class="address">{{data.item.operatorAddress}}</span>
          </div>
        </span>
        </template>
      </b-table>
    </div>
</template>

<script>
    import Tools from '../../util/Tools';
    export default {
        name: "ValidatorListTable",
        watch: {
          items(items) {
            this.setValidatorFields(items);
          },
        },
        props: ['items', 'type','showNoData','minWidth','status'],
        data(){
          return {
            fields: [],
            sortBy: 'votingPower',
            sortDesc: true,
            validatorFields:null,
            activeValidatorFields: {

              // index:{
              //   label:'Moniker',
              // },
              moniker:{
                label:'Moniker',
                sortable:true,
              },
              operatorAddress:{
                label:'Operator_Address',
                sortable:false,
              },
              commission:{
                label:'Commission',
                sortable:true,
              },
              'bondedToken':{
                label:'Bonded Tokens',
                sortable:true,
              },
              'votingPower':{
                label:'Voting Power',
                sortable:true,
              },
              'uptime':{
                label:'Uptime',
                sortable:true,
              },
              'selfBond':{
                label:'Self Bonded',
                sortable:true,
              },
              'delegatorNum':{
                label:'Delegator Number',
                sortable:true,
              },
              'bondHeight':{
                label:'Bond Height',
                sortable:false,
              },
            },
            jailedValidatorFields: {
              moniker:{
                label:'Moniker',
                sortable:true,
              },
              operatorAddress:{
                label:'Operator_Address',
                sortable:false,
              },
              commission:{
                label:'Commission',
                sortable:true,
              },
              'bondedToken':{
                label:'Bonded Tokens',
                sortable:true,
              },
              'selfBond':{
                label:'Self Bonded',
                sortable:true,
              },
              'bondHeight':{
                label:'Bond Height',
                sortable:false,
              },
            },
            candidateValidatorFields: {
              moniker:{
                label:'Moniker',
                sortable:true,
              },
              operatorAddress:{
                label:'Operator_Address',
                sortable:false,
              },
              commission:{
                label:'Commission',
                sortable:true,
              },
              'bondedToken':{
                label:'Bonded Tokens',
                sortable:true,
              },
              'selfBond':{
                label:'Self Bonded',
                sortable:true,
              },
              'delegatorNum':{
                label:'Delegator Number',
                sortable:true,
              },
              'bondHeight':{
                label:'Bond Height',
                sortable:false,
              },
            },
          }
        },
        mounted(){
          this.validatorFields =  this.activeValidatorFields;
        },
        methods: {
            formatAddress(address){
                return Tools.formatValidatorAddress(address)
            },
            setValidatorFields(validatorList){
                validatorList.forEach(item => {
                    if(item && item.validatorStatus && item.validatorStatus === 'jailed'){
                        this.validatorFields = this.jailedValidatorFields
                    }else if(item && item.validatorStatus && item.validatorStatus === 'validator'){
                        this.validatorFields = this.activeValidatorFields
                    }else if(item && item.validatorStatus && item.validatorStatus === 'candidate'){
                        this.validatorFields = this.candidateValidatorFields
                    }
                })
        }
      }
    }
</script>

<style lang="scss">
  @import '../../style/mixin.scss';
  .name_address{
    display: flex;
    flex-direction: column;
    position: relative;
  .address{
    display: none;
    position: absolute;
    left: -1.05rem;
    top: -0.38rem;
    color: #3598db;
    background: rgba(0,0,0,0.8);
    border-radius:0.04rem;
    z-index: 10;
  }
    &:hover{
      .address{
        background: rgba(0,0,0,1);
        color: #fff;
        padding: 0.06rem 0.15rem 0 0.15rem;
        display: block;
        border-radius:0.04rem;
        font-size: 0.14rem;
      &::after{
         content: '';
         display: block;
         background: rgba(0,0,0,1);
         transform: rotate(45deg);
         width: 0;
         height: 0;
         border: 0.04rem solid transparent;
         position: relative;
         top: 0.03rem;
         z-index: 1;
         left: 1.45rem;
       }
      }
    }
  }
  //重置bootstrap-vue的表格样式
    table {
      td {
        max-width: 2.2rem !important;
        overflow-wrap: break-word !important;
        word-wrap: break-word !important;
      .skip_route {
      a{
        color: #3598db!important;
        cursor: pointer;
      }
    }
    .no_skip{
      color:#A2A2AE;
      cursor:default;
    .link_style{
      color:#a2a2ae !important;
    }
  }
  }
  }

  .page-link{
    padding:0.05rem 0.075rem !important;
    height:0.29rem !important;
    color:#3598db !important;
  }
  .active{
    .page-link{
      background-color: #3598db !important;
      border-color:#3598db !important;
      color:#fff !important;
    }
  }
  .table{
    th, td{
      padding:0.075rem !important;
      color:#A2A2AE;
    @include fontWeight;
    }
    margin-bottom:0 !important;
    thead{
      th{
        border-bottom:none !important;
      }
      tr{
        th{
          color:#000000;
          height:0.5rem;
          vertical-align:middle;
          &:first-child{
             padding-left:0.2rem !important;
           }
          }
        border-bottom:0.02rem solid #3598db;
      }
    }
    tbody{
      tr:nth-child(1){
        td{
          border-top:none;
        }
      }
      tr{
        &:nth-of-type(even){
           background-color: #f6f6f6 !important;
         }
        &:nth-of-type(odd){
           background-color: #fff !important;
         }
        &:last-child{
           border-bottom:1px solid #dee2e6;
         }
      }
    }
  }
  .show_no_data{
    .table{
      tbody{
        tr{
        &:first-child{
           display:none;
         }
        }
      }
    }
  }
  .proposal_detail_list tr{
    th:nth-child(1){
      width: 50% !important;
    }
    th:nth-child(2){
      width: 35% !important;
    }
  }
  //使用rem设置max-width不生效
  @media screen and (max-width: 910px) {
    .proposal_detail_list tr{
      th:nth-child(1){
        width: 50% !important;
      }
      th:nth-child(2){
        width: auto !important;
      }
    }
  }
  .proposals-list{
    color: #3598db;
    cursor: pointer;
    margin: 0!important;
    padding: 0!important;
  }
  .remove_default_style{
    margin: 0!important;
    padding: 0!important;
    color: #3598db;
  }
  .show_trim td{
    white-space: pre;
  }
  .show_trim td span{
    white-space: pre;
  }
  .show_trim thead tr th{
    outline: transparent;
    &::after{
       margin-bottom: 0.13rem;
     }
  }
  .show_trim tbody tr td:nth-child(1){
    padding: 0.075rem !important;
    width: 1.9rem;
  }
  .show_trim tbody tr td:nth-child(3){
    padding: 0.075rem !important;
  }
  .pre_global_style{
    font-size: 0.14rem;
    color: #a2a2ae;
    margin: 0;
  }
  pre{
    font-family: Arial !important;
    margin: 0;
  }
  .link_style{
    color: #3598db !important;
  }
  //覆盖bootstrap-vue默认排序样式
  .b-table.table > thead > tr > th[aria-sort][aria-sort="descending"]::after, .b-table.table > tfoot > tr > th[aria-sort][aria-sort="descending"]::after{
    background: url("../../assets/desc.png") no-repeat center;
    content:"" !important;
    width: 0.1rem;
    height: 0.20rem;
    background-size: 100% ;
    opacity: 1;
  }
  .b-table.table > thead > tr > th[aria-sort][aria-sort="ascending"]::after, .b-table.table > tfoot > tr > th[aria-sort][aria-sort="ascending"]::after{
    background: url("../../assets/asc.png") no-repeat center;
    content:"" !important;
    width: 0.1rem;
    height: 0.20rem;
    background-size: 100% ;
    opacity: 1;
  }
  .b-table.table > thead > tr > th[aria-sort]::after, .b-table.table > tfoot > tr > th[aria-sort]::after{
    background: url("../../assets/default.png") no-repeat center;
    background-size: 100% ;
    width: 0.1rem;
    height: 0.20rem;
    content:"" !important;
    opacity: 1;
  }
  table.b-table>tfoot>tr>th.sorting:before, table.b-table>thead>tr>th.sorting:before{
    content:"" !important;
  }
  table.b-table>tfoot>tr>th.sorting:after, table.b-table>thead>tr>th.sorting:after{
    background: url("../../assets/default.png") no-repeat center;
    background-size: 100% ;
    width: 0.1rem;
    height: 0.20rem;
    content:"" !important;
    opacity: 1;
  }
</style>
