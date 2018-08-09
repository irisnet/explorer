<template>
  <div :class="showNoData?'show_no_data':''"
       class="table_wrap"
       :style="`${minWidth?(`min-width:${minWidth}rem`):''}`">
    <b-table :fields='fields' :items='items' striped v-if="type === '1'">
      <template slot='Height' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Height}`)">
          {{data.item.Height}}
        </span>
      </template>
      <template slot='Txn' slot-scope='data'>
        <span class="skip_route"
              v-show="data.item.Txn != 0"
              @click="skipRoute(`/recent_transactions/2/recent?block=${data.item.Height}`)">
          {{data.item.Txn}}
        </span>
        <span v-show="data.item.Txn == 0">{{data.item.Txn}}</span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '2'">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/tx?txHash=${data.item.TxHash}`)">
          {{data.item.TxHash?`${String(data.item.TxHash).substr(0,16)}...`:''}}
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Block}`)">
          {{data.item.Block}}
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.From}`)">
          {{data.item.From?`${String(data.item.From).substr(0,16)}...`:''}}
        </span>
      </template>
      <template slot='To' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.To}`)">
          {{data.item.To?`${String(data.item.To).substr(0,16)}...`:''}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '3' || type === '4'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.Address}`)">
          {{data.item.Address?`${String(data.item.Address).substr(0,16)}...`:''}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '5'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.Address}`)">
          {{data.item.Address?`${String(data.item.Address).substr(0,16)}...`:''}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '6'" style="margin-bottom:0;">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/tx?txHash=${data.item.TxHash}`)">
          {{data.item.TxHash?`${String(data.item.TxHash).substr(0,16)}...`:''}}
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Block}`)">
          {{data.item.Block}}
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <span class="skip_route"
              :class="data.item.From === $route.params.param?'no_skip':''"
              @click="skipRoute(`/address/1/${data.item.From}`)">
          {{data.item.From?`${String(data.item.From).substr(0,16)}...`:''}}
        </span>
      </template>
      <template slot='To' slot-scope='data'>
        <span class="skip_route"
              :class="data.item.To === $route.params.param?'no_skip':''"
              @click="skipRoute(`/address/1/${data.item.To}`)">
          {{data.item.To?`${String(data.item.To).substr(0,16)}...`:''}}
        </span>
      </template>


    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '7'" style="margin-bottom:0;">
      <template slot='Block Height' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item['Block Height']}`)">
          {{data.item['Block Height']}}
        </span>
      </template>
    </b-table>
  </div>

</template>

<script>

  export default {
    watch: {
      items(items) {

      },

    },
    data() {
      return {
        fields: [],
      }
    },
    props: ['items', 'type','showNoData','minWidth'],
    mounted() {

    },
    methods: {
      skipRoute(path) {
        this.$router.push(path);
      }
    }
  }
</script>
<style lang="scss">
  @import '../../style/mixin.scss';

  //重置bootstrap-vue的表格样式
  table {


    td {
      max-width: 2.2rem !important;
      overflow-wrap: break-word !important;
      .skip_route {
        color: #3598db;
        cursor: pointer;
      }
      .no_skip{
        color:#A2A2AE;
        cursor:default;
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
      tr{
        &:nth-of-type(odd){
          background-color: #f6f6f6 !important;
        }
        &:last-child{
          border-bottom:1px solid #dee2e6;
        }
        td{
          &:first-child{
            padding-left:0.2rem !important;
          }
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

</style>
