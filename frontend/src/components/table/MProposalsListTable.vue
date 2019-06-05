<template>
  <div :style="{'min-width': '12.8rem'}">
    <m-table :columns="validatorFields" class="top_border_table"
             :data="items"
             :sort-by.sync="sortBy"
             :sort-desc.sync="sortDesc">
      <!-- <template slot-scope="{ row }"
                slot="operatorAddress">
        <span class="remove_default_style">
          <router-link :to="`/address/1/${row.operatorAddress}`"
                       class="link_style">{{formatAddress(row.operatorAddress)}}</router-link>
        </span>
      </template> -->
    </m-table>
  </div>
</template>

<script>
import Tools from '../../util/Tools';

export default {
  name: 'MProposalsListTable',
  props: {
    items: {
      type: Array,
      default: function() {return[]}
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
      resizeTime: null,
      sortBy: '',
      sortDesc: true,
      validatorFields: [
        {
          title: 'ID',
          slot: 'id',
          width: 190,
          sortable: true
        },
        {
          title: 'Title',
          slot: 'title',
          tooltip: true
        },
        {
          title: 'Status',
          key: 'status',
          sortable: true,
          sortMethod: this.sortMethodPer('commission'),
          className: 'text_right'
        },
        {
          title: 'Submit_Time',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken'),
          className: 'text_right'
        },
        {
          title: 'Deposit_Endtime',
          key: 'votingPower',
          sortable: true,
          sortMethod: this.sortMethodPer('votingPower'),
          className: 'text_right'
        },
        {
          title: 'Voting_Endtime',
          key: 'uptime',
          sortable: true,
          sortMethod: this.sortMethodPer('uptime'),
          className: 'text_right'
        }
      ]
    }
  },
  methods: {
    sortMethodPer (key) {
      return (a, b) => Number(a[key].split('%')[0]) - Number(b[key].split('%')[0]);
    },
    sortMethodNumber (key) {
      return (a, b) => Number(a[key]) - Number(b[key]);
    },
    sortMethodSplit (key) {
      return (a, b) => Number(a[key].replace(/\,/g, '').split(' ')[0]) - Number(b[key].replace(/\,/g, '').split(' ')[0]);
    },
    formatAddress (address) {
      return Tools.formatValidatorAddress(address);
    }
  }
}
</script>

<style lang="scss">
  .top_border_table {
    .m-table-header {
      table {
        thead {
          tr {
            th {
              box-sizing: border-box;
              border-top: 1px solid rgb(222, 226, 230);
            }
          }
        }
      }
    }
  }
</style>
