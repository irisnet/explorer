<template>
  <div :class="showNoData?'show_no_data':''" :style="`${minWidth?(`min-width:${minWidth}rem`):''}`">
    <b-table :fields='fields' :items='items' striped class="parameter_table_container">
      <template slot="Range" slot-scope="data">
        <div v-if="data.item.Range.minimum.data || data.item.Range.maximum.data">
          <span>{{formatMinRange(data.item.Range.minimum)}}{{data.item.Range.minimum.data}} : {{data.item.Range.maximum.data}}{{formatMaxRange(data.item.Range.maximum)}}</span>
        </div>
      </template>
    </b-table>
  </div>
</template>

<script>
    import Constant from "../../constant/Constant";

    export default {
        name: "vParametersTable",
        props:['items', 'type','showNoData','minWidth','status'],
        watch: {
            parametersList(parametersList){
            }
        },
        data(){
            return {
                fields: [],
            }
        },
        methods:{
            formatMinRange(range){
                if(range.sign === Constant.PARAMETER.EQUAL){
                  return `[ `
                }else if(range.sign === Constant.PARAMETER.UNEQUAL){
                  return '( '
                }
            },
            formatMaxRange(range){
              if(range.sign === Constant.PARAMETER.EQUAL){
                return ` ]`
              }else if(range.sign === Constant.PARAMETER.UNEQUAL){
                return ' )'
              }else{
                return '+∞ )'
              }
            },
        }
    }
</script>

<style lang="scss">
  @import "../../style/mixin.scss";
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
  .parameter_table_container td{
    white-space: pre;
  }
  .parameter_table_container td span{
    white-space: pre;
  }
  .parameter_table_container thead tr th{
    outline: transparent;
    &::after{
       margin-bottom: 0.13rem;
     }
  }
  .parameter_table_container tbody tr td:nth-child(1){
    padding: 0.075rem !important;
    width: 1.9rem;
  }
  .parameter_table_container tbody tr td{
    word-break: break-all;
    white-space: inherit;
  }
  .parameter_table_container tbody tr td:nth-child(1){
    width: 20%;
  }
</style>
