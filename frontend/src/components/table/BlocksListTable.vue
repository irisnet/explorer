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
        <span>{{data.item.Txn}}</span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === '2'" class="block_style">
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
      <template slot='Owner' slot-scope='data'>
        <span class="skip_route"  @click="skipRoute(`/address/1/${data.item.Owner}`)">
          {{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}
        </span>
      </template>
      <template slot='Moniker' slot-scope='data'>
        <span>
          <pre class="pre_global_style">{{data.item.Moniker ? data.item.Moniker : ''}}</pre>
        </span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === '3' || type === '4'" class="show_trim">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.Address}`)">
          <span class="remove_default_style">{{data.item.Address?`${String(data.item.Address).substr(0,16)}...`:''}}</span>
        </span>
      </template>
      <template slot='Name' slot-scope='data'>
        <span>
          <pre class="pre_global_style">{{data.item['Name']}}</pre>
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

    <b-table :fields='fields' :items='items' striped  v-if="type === 'Proposals'" class="show_trim">
      <template slot='Title' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/ProposalsDetail/${data.item['Proposal ID']}`)">
          <pre class="proposals-list">{{data.item['Title']}}</pre>
        </span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === 'ProposalsDetail'" nodelabel  class="proposal_detail_list">
      <template slot='Voter' slot-scope='data'>
        <span class="skip_route_gray">
          {{data.item['Voter']}}
        </span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === 'addressTxList'" nodelabel >
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/tx?txHash=${data.item.TxHash}`)">
          {{data.item['TxHash'] ? `${String(data.item.TxHash).substr(0,16)}...` : ''}}
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
      <template slot='Owner' slot-scope='data'>
        <span>
          {{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === 'blockTxList'" nodelabel class="block_style">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/tx?txHash=${data.item.TxHash}`)">
          {{data.item['TxHash'] ? `${String(data.item.TxHash).substr(0,16)}...` : ''}}
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
      <template slot='Owner' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.Owner}`)">{{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}</span>
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
      word-wrap: break-word !important;
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
  .proposal_detail_list tr{
    th:nth-child(1){
      width: 50% !important;
    }
    th:nth-child(2){
      width: 35% !important;
    }
  }
  .proposals-list{
    color: #3598db;
    cursor: pointer;
    margin: 0;
    padding: 0;
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
  .block_style thead tr th:nth-child(1){
    width: 16%;
  }
  .block_style thead tr th:nth-child(2){
    width: 8%;
  }
  .pre_global_style{
    font-size: 0.14rem;
    color: #a2a2ae;
  }
  pre{
    font-family: Arial !important;
  }
</style>
