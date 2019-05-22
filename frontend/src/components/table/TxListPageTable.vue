<template>
    <div :class="showNoData?'show_no_data':''" style="min-width: 12rem">
      <b-table :fields='fields' :items='items' striped class="block_style">
        <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item.TxHash ? `${formatTxHash(String(data.item.TxHash))}` : ''}}</router-link>
        </span>
        </template>
        <template slot='Age' slot-scope='data'>
          <span v-show="data.item.Age">{{data.item.Age}}</span>
        </template>
        <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
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
    </div>
</template>

<script>
  import Tools from '../../util/Tools';
	export default {
      name: "TxListPageTable",
      props:['items','showNoData','minWidth'],
      data () {
          return {
            fields:[]
          }
      },
      watch:{
        showNoData(showNoData){
        }
      },
      methods:{
        formatAddress(address){
          return Tools.formatValidatorAddress(address)
        },
        formatTxHash(TxHash){
          if(TxHash){
            return Tools.formatTxHash(TxHash)
          }
        },
      }
	}
</script>

<style lang="scss">
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
