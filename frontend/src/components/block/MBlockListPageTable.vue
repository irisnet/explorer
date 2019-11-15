<template>
  <div>
    <m-table v-table-head-fixed :columns="columns"
             :data="items">
      <template slot='height' slot-scope='{ row }'>
        <span class="skip_route">
          <router-link :to="`/block/${row.height}`">{{row.height}}</router-link>
        </span>
      </template>
      <template slot='proposer' slot-scope='{ row }'>
        <span class="skip_route">
          <router-link v-if="row.moniker" :to="addressRoute(row.proposerAddr)">{{row.moniker}}</router-link>
          <span v-if="!row.moniker">--</span>
        </span>
      </template>
      <template slot='Age' slot-scope='{ row }'>
        <span v-show="row.Age">{{row.Age}}</span>
      </template>
    </m-table>
  </div>
</template>

<script>
export default {
  name: 'MBlockListPageTable',
  props: {
    items: {
      type: Array,
      default: function() {return[]}
    }
  },
  data() {
    return {
      columns: [
        {
          title: 'Block',
          slot: 'height'
        },
        {
          title: 'Proposer',
          slot: 'proposer'
        },
        {
          title: 'Transactions',
          key: 'transactions'
        },
        {
          title: 'Validators',
          key: 'validators'
        },
        {
          title: 'Voting_Power',
          key: 'votingPower'
        },
        {
          title: 'Timestamp',
          key: 'timestamp'
        },
        {
          title: 'Age',
          key: 'Age',
          width: 220
        }
      ]
    }
  }
}
</script>
<style lang="scss">
/*  .skip_route{
    a{
      color: var(--bgColor) !important;
    }
  }*/
</style>
