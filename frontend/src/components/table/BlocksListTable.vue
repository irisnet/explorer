<template>
  <div :class="showNoData?'show_no_data':''"
       class="table_wrap"
       :style="`${minWidth?(`min-width:${minWidth}rem`):''}`">

    <b-table :fields='blockFields' :items='items' striped v-if="type === '1'">
      <template slot='Height' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/blocks_detail/${data.item.Height}`">{{data.item.Height}}</router-link>
        </span>
      </template>
      <template slot='Txn' slot-scope='data'>
        <span>{{data.item.Txn}}</span>
      </template>
      <template slot='Age' slot-scope='data'>
        <span v-show="data.item.Age">{{data.item.Age}}</span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === '2'" class="block_style">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item.TxHash?`${String(data.item.TxHash).substr(0,16)}...`:''}}</router-link>
        </span>
      </template>
      <template slot='Age' slot-scope='data'>
        <span v-show="data.item.Age">{{data.item.Age}}</span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/blocks_detail/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <span class="skip_route" style="display: flex" v-if="data.item.From">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${data.item.From}`" class="link_style justify">{{formatAddress(data.item.From)}}</router-link>
            </span>
            <span class="address">{{data.item.From}}</span>
          </div>
        </span>
      </template>
      <template slot='To' slot-scope='data'>
        <span class="skip_route" style="display: flex" v-if="data.item.To !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${data.item.To}`" class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
            <span class="address">{{data.item.To}}</span>
          </div>
        </span>
        <span class="no_skip" v-show="data.item.To === '--'">
          --
        </span>
      </template>
      <template slot='Owner' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/address/1/${data.item.Owner}`">{{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}</router-link>
        </span>
      </template>
      <template slot='Moniker' slot-scope='data'>
        <span>
          <pre class="pre_global_style">{{data.item.Moniker ? data.item.Moniker : ''}}</pre>
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '5'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" style="display: flex">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${data.item.Address}`" class="link_style">{{formatAddress(data.item.Address)}}</router-link>
            </span>
            <span class="address">{{data.item.Address?`${formatAddress(data.item.Address)}`:''}}</span>
          </div>
        </span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === '6'" style="margin-bottom:0;">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item.TxHash?`${String(data.item.TxHash).substr(0,16)}...`:''}}</router-link>
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/blocks_detail/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <span class="skip_route" style="display: flex">
          <div class="name_address">
            <span class="remove_default_style" :class="data.item.From === $route.params.param?'no_skip':''">
              <router-link :to="`/address/1/${data.item.Address}`" class="link_style">{{formatAddress(data.item.Address)}}</router-link>
            </span>
            <span class="address">{{data.item.Address?`${formatAddress(data.item.Address)}`:''}}</span>
          </div>
        </span>
      </template>
      <template slot='To' slot-scope='data'>
        <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="data.item.To === $route.params.param ? '' : `/address/1/${data.item.To}`" class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
          <span class="address">{{data.item.Address?`${formatAddress(data.item.To)}`:''}}</span>
        </div>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === '7'" style="margin-bottom:0;">
      <template slot='Block Height' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/blocks_detail/${data.item['Block Height']}`">{{data.item['Block Height']}}</router-link>
        </span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped  v-if="type === 'Proposals'" class="show_trim">
      <template slot='Title' slot-scope='data'>
        <span class="skip_route">
         <router-link :to="`/ProposalsDetail/${data.item['Proposal ID']}`"><pre class="proposals-list">{{data.item['Title']}}</pre></router-link>
        </span>
      </template>
      <template slot="SubmitTime" slot-scope='data'>
        <span v-show="data.item.SubmitTime">{{data.item.SubmitTime}}</span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === 'ProposalsDetail'" nodelabel  class="proposal_detail_list">
      <template slot='Voter' slot-scope='data'>
        <span class="skip_route_gray">
          {{data.item['Voter']}}
        </span>
      </template>
      <template slot='VoteTime' slot-scope='data'>
        <span v-show="data.item.VoteTime">{{data.item.VoteTime}}</span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === 'addressTxList'" nodelabel >
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item['TxHash'] ? `${String(data.item.TxHash).substr(0,16)}...` : ''}}</router-link>
        </span>
      </template>
      <template slot='Age' slot-scope='data'>
        <span v-show="data.item.Age">{{data.item.Age}}</span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/blocks_detail/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <div class="name_address" v-show="data.item.From">
            <span class="remove_default_style" :class="data.item.From === $route.params.param?'no_skip':''">
              <router-link :to="`/address/1/${data.item.From}`" class="link_style">{{formatAddress(data.item.From)}}</router-link>
            </span>
          <span class="address">{{data.item.From ? data.item.From : ''}}</span>
        </div>
      </template>
      <template slot='To' slot-scope='data'>
        <div class="name_address" v-show="data.item.To && data.item.To !== '--'">
            <span class="remove_default_style" :class="data.item.To === $route.params.param?'no_skip':''">
              <router-link :to="`/address/1/${data.item.To}`" class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
          <span class="address">{{data.item.To ? data.item.To : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.To == '--'">
          --
        </span>
      </template>
      <template slot='Owner' slot-scope='data'>
        <span>
          {{data.item.Owner?`${formatAddress(data.item.Owner)}`:''}}
        </span>
      </template>
    </b-table>

    <b-table :fields='fields' :items='items' striped v-if="type === 'blockTxList'" nodelabel class="block_style">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item['TxHash'] ? `${String(data.item.TxHash).substr(0,16)}...` : ''}}</router-link>
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/blocks_detail/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <div class="name_address" v-show="data.item.From && data.item.From !== '--'">
            <span class="remove_default_style" :class="data.item.From === $route.params.param?'no_skip':''">
              <router-link :to="data.item.From === $route.params.param ? '' : `/address/1/${data.item.From}`" class="link_style">{{formatAddress(data.item.From)}}</router-link>
            </span>
          <span class="address">{{data.item.From ? data.item.From : ''}}</span>
        </div>
      </template>
      <template slot='To' slot-scope='data'>
        <div class="name_address" v-show="data.item.To && data.item.To !== '--'">
            <span class="remove_default_style" :class="data.item.To === $route.params.param?'no_skip':''">
              <router-link :to="data.item.To === $route.params.param ? '' : `/address/1/${data.item.To}`" class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
          <span class="address">{{data.item.To ? data.item.To : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.To == '--'">
          --
        </span>
      </template>
      <template slot='Owner' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/address/1/${data.item.Owner}`">{{data.item.Owner?`${formatAddress(data.item.Owner)}`:''}}</router-link>
        </span>
      </template>
    </b-table>
  </div>

</template>

<script>
 import Tools from '../../util/Tools';
  export default {
    watch: {
      items(items) {

      },
    },
    data() {
      return {
        fields: [],
        blockFields:['Height','Txn','Age','Precommit Validators','Voting Power']
      }
    },
    props: ['items', 'type','showNoData','minWidth','status'],
    methods: {
      formatAddress(address){
        return Tools.formatValidatorAddress(address)
      },
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
  .block_style thead tr th:nth-child(1){
    width: 16%;
  }
  .block_style thead tr th:nth-child(2){
    width: 8%;
  }
  .pre_global_style{
    font-size: 0.14rem;
    color: #a2a2ae;
    margin: 0;
  }
  .proposals_detail_table_wrap{
    tbody tr td:last-child{
      min-width: 2rem;
    }
  }
  pre{
    font-family: Arial !important;
    margin: 0;
  }
  .link_style{
    color: #3598db !important;
  }
</style>
