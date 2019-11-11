<template>
    <div :class="showNoData?'show_no_data':''" :style="{'min-width': minWidth + 'rem'}">
      <m-table :columns='fields' :data='items'  class="top_list_table_container" v-table-head-fixed>
        <template slot-scope="{ row }" slot="rank" style="text-align: center">
          <span style="padding-left: 0.2rem">{{row.rank}}</span>
        </template>
        <template slot-scope="{ row }" slot="Address" style="text-align: center">
          <span class="skip_route" style="display: flex" v-if="row.Address">
          <div class="name_address">
            <span class="remove_default_style" style="font-family: Consolas,Menlo ;">
              <router-link :to="addressRoute(row.Address)" class="link_style">{{row.Address}}</router-link>
            </span>
          </div>
        </span>
        </template>
        <template slot-scope="{ row }" slot="Percentage">
          <div  class="percentage_item" style="width: 100%;display: flex;height: 0.2rem;overflow: hidden">
            <!--<div :style="{'width':`${data.item.Percentage < 0.6 ? '0.3': data.item.Percentage}%`,'background':'var(--bgColor)'}"><span style="visibility: hidden;">{{data.item.Percentage}}</span></div>-->
            <!--<div style="padding-left: 0.2rem">{{data.item.Percentage}}%</div>-->
            <div>{{row.Percentage}}%</div>
          </div>
        </template>
      </m-table>
    </div>
</template>

<script>
  import Tools from '../../util/Tools'
    export default {
        name: "TopListTable",
        props:{
          items: {
            type: Array,
            default: []
          },
          showNoData: {
            type: Boolean,
            default: true
          },
          minWidth: {
            type: Number,
            default: 12.8
          }
        },
        data () {
            return {
              fields: [
                {
                  title: '#',
                  slot: 'rank',
                },
                {
                  title: 'Address',
                  slot: 'Address',
                },
                {
                  title: 'Amount (IRIS)',
                  key: 'Balance',
                  className: 'text_right'
                },
                {
                  title: 'Percentage',
                  slot: 'Percentage',
                },
              ]
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
  .top_list_table_container{
    table thead tr:nth-of-type(1){
      th:nth-of-type(1){
        padding-left: 0.28rem;
      }
    }
    .text_right{
      padding-right: 1.5rem !important;
    }
  }
</style>
