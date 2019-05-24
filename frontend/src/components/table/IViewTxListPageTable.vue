<template>
  <div>
    <i-table :columns="columns" :data="data">
      <template slot-scope="{ row }" slot="TxHash">
        <router-link class="equare_width" :to="`/tx?txHash=${row.TxHash}`">{{row.TxHash?`${String(row.TxHash).substr(0,16)}...`:''}}</router-link>
      </template>
      <template slot-scope="{ row }" slot="Block">
        <router-link class="skip_route" :to="`/block/${row.Block}`">{{row.Block}}</router-link>
      </template>
      <template slot-scope="{ row }" slot="From">
        <router-link :to="`/address/1/${row.From}`" class="link_style justify">{{row.From && formatAddress(row.From)}}</router-link>
      </template>
      <template slot-scope="{ row }" slot="To">
        <router-link :to="`/address/1/${row.To}`" class="link_style">{{row.To && formatAddress(row.To)}}</router-link>
      </template>
    </i-table>
  </div>
</template>

<script>
  import Tools from '../../util/Tools';

	export default {
    name: "IViewTxListPageTable",
    props: {
      items: {
        type: Array,
        default: []
      }
    },
    data() {
      return {
        columns: [
          {
            title: 'Tx Hash',
            slot: 'TxHash',
            width: 200
          },
          {
            title: 'Block',
            slot: 'Block',
            width: 100
          },
          {
            title: 'From',
            slot: 'From',
            width: 200
          },
          {
            title: 'To',
            slot: 'To',
            width: 200
          },
          {
            title: 'Amount', 
            key: 'Amount',
            width: 170
          },
          {
            title: 'Fee',
            key: 'Fee',
            width: 140
          },
          {
            title: 'Status',
            key: 'Status'
          },
          {
            title: 'Age',
            key: 'Age',
            width: 140
          }
        ],
        data: []
      }
    },
    methods: {
      formatAddress(address){
        return Tools.formatValidatorAddress(address)
      }
    },
    watch: {
      items(newVal, oldVal) {
        this.data = newVal;
      }
    }
	}
</script>

<style lang="scss" scoped>
  .equare_width {
    // font-family:Consolas, monospace, Monaco;
    font-size: 0.14rem;
    color: #3598db!important;
  }
  .skip_route {
    color: #3598db!important;
  }
</style>
