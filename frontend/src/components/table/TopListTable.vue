<template>
    <div :class="showNoData?'show_no_data':''" :style="`${minWidth?(`min-width:${minWidth}rem`):''}`">
      <b-table :fields='fields' :items='items' striped class="top_list_table_container">
        <template slot="rank" slot-scope="data" style="text-align: center">
          <span style="padding-left: 0.2rem">{{data.item.rank}}</span>
        </template>
        <template slot="Address" slot-scope="data" style="text-align: center">
          <span class="skip_route" style="display: flex" v-if="data.item.Address">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${data.item.Address}`" class="link_style">{{data.item.Address}}</router-link>
            </span>
          </div>
        </span>
        </template>
        <template slot="Percentage" slot-scope="data">
          <div  class="percentage_item" style="width: 100%;display: flex;height: 0.2rem;overflow: hidden">
            <!--<div :style="{'width':`${data.item.Percentage < 0.6 ? '0.3': data.item.Percentage}%`,'background':'#3598db'}"><span style="visibility: hidden;">{{data.item.Percentage}}</span></div>-->
            <!--<div style="padding-left: 0.2rem">{{data.item.Percentage}}%</div>-->
            <div>{{data.item.Percentage}}%</div>
          </div>
        </template>
      </b-table>
    </div>
</template>

<script>
  import Tools from '../../util/Tools'
    export default {
        name: "TopListTable",
        props:['items','showNoData','minWidth',],
        data () {
            return {
              fields: {
                rank:{
                  label:'#',
                },
                Address:{
                  label:'Address',
                },
                Balance:{
                  label:'Amount (IRIS)',
                },
                Percentage:{
                  label:'Percentage',
                }
              }
            }
        },
        methods:{
          formatAddress(address){
            return Tools.formatValidatorAddress(address)
          },
        }
    }
</script>

<style  lang="scss">
  .name_address{
    display: flex;
    flex-direction: column;
    position: relative;
    .address{
      display: none;
      position: absolute;
      left: -1.29rem;
      top: -0.38rem;
      color: #3598db;
      background: rgba(0,0,0,0.8);
      border-radius:0.04rem;
      z-index: 10;
    }
    .percentage_item{
      &:nth-of-type(even){
        background-color: #f6f6f6 !important;
      }
      &:nth-of-type(odd){
        background-color: #fff !important;
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
  .top_list_table_container.table thead tr th{
    font-size: 0.14rem;
  }

  .top_list_table_container.table thead tr th:nth-child(1){
    width: 1.5rem;
    text-align: left;
    padding-left: 0.28rem !important;
  }
  .top_list_table_container.table thead tr th:nth-child(2){
    width: 4rem;
    text-align: left;
  }
  .top_list_table_container.table thead tr th:nth-child(3){
    padding-left: 0.94rem !important;
  }
  .top_list_table_container.table thead tr th:nth-child(4){
    width: 3rem;
  }
  .top_list_table_container.table tbody tr td:nth-child(1){
    text-align: left;
  }
  .top_list_table_container.table tbody tr td:nth-child(2){
    width: 4rem;
    .skip_route{
     display: flex;
     justify-content: flex-start;
   }
  }
  .top_list_table_container.table tbody tr td:nth-child(3){
    text-align: right;
    padding-right: 1.6rem !important;
  }
  .top_list_table_container.table tbody tr td:nth-child(4){
    text-align: right;
    width: 3rem;
    margin-left: 0.5rem;
  }
</style>
