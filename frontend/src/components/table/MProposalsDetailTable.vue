<template>
  <div>
    <m-table :columns="columns"
             :data="items">
      <template slot-scope="{ row }"
                slot="Voter">
        <div class="skip_route">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="`/address/1/${row.Voter}`" class="link_style justify">{{row.moniker}}</router-link>
            </span>
          </div>
        </div>
      </template>
      <template slot='Tx_Hash' slot-scope="{ row }">
        <router-link class="skip_route" :to="`/tx?txHash=${row.Tx_Hash}`">{{row.Tx_Hash ? `${formatTxHash(String(row.Tx_Hash))}` : ''}}</router-link>
      </template>
      <template slot='Depositor' slot-scope="{ row }">
        <span v-if="(/^[1-9]\d*$/).test(row.Depositor)" class="skip_route">
          <router-link :to="`/tx?txHash=${row.Tx_Hash}`">{{row.From}} Validators</router-link>
        </span>
        <span class="skip_route" v-if="!(/^[0-9]\d*$/).test(row.Depositor) && row.Depositor !== '--'">
          <div class="name_address">
            <router-link :to="`/address/1/${row.Depositor}`" class="link_style justify">{{formatAddress(row.Depositor)}}</router-link>
          </div>
        </span>
        <span class="no_skip" v-show="(/^[0]\d*$/).test(row.Depositor) || row.Depositor === '--'">--</span>
      </template>
    </m-table>
  </div>
</template>
<script>
import Tools from '../../util/Tools';

export default {
  name: 'MProposalsDetailTable',
  props: {
    items: {
      type: Array,
      default: function() {return[]}
    },
    fields: {
      type: String,
      default: 'votersFields'
    }
  },
  data () {
    return {
      columns: [],
      votersFields: [
        {
          title: 'Voter',
          slot: 'Voter'
        },
        {
          title: 'Vote Option',
          key: 'Vote_Option'
        },
        {
          title: 'Tx Hash',
          slot: 'Tx_Hash',
          tooltip: true
        },
        {
          title: 'Time',
          key: 'Time',
          className: 'text_right'
        }
      ],
      depositorsFields: [
        {
          title: 'Depositor',
          slot: 'Depositor',
          tooltip: true,
          tooltipClassName: 'tooltip_left'
        },
        {
          title: 'Amount',
          key: 'Amount'
        },
        {
          title: 'Type',
          key: 'Type'
        },
        {
          title: 'Tx Hash',
          slot: 'Tx_Hash',
          className: 'text_right',
          tooltip: true
        },
        {
          title: 'Time',
          key: 'Time',
          className: 'text_right'
        }
      ]
    }
  },
  methods: {
    formatTxHash(TxHash){
      if(TxHash){
        return Tools.formatTxHash(TxHash)
      }
    },
    formatAddress(address){
      return Tools.formatValidatorAddress(address)
    }
  },
  mounted() {
    this.columns = this[this.fields];
  }
}
</script>

<style lang="scss" scoped>
.skip_route {
  color: #3598db!important;
}
.name_address {
  .remove_default_style {
    line-height: 16px;
    a {
      line-height: 16px;
    }
  }
}
</style>
