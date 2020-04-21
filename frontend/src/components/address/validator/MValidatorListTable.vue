<template>
  <div :style="{'min-width': minWidth + 'rem'}">
    <m-table v-table-head-fixed
             :columns="validatorFields"
             :data="validatorData"
             :sort-by.sync="sortBy"
             :sort-desc.sync="sortDesc">
      <template slot-scope="{ row }"
                slot="moniker">
        <div style="display: flex;align-items: center;position: relative">
          <span style="background:#E0E8FF;width: 0.3rem;height: 0.3rem;border-radius: 0.3rem;overflow: hidden;display: flex;align-items: center;justify-content: center">{{row.validatorIconSrc}}</span>
          <img v-if="row.url"
               style="width: 0.3rem;height: 0.3rem;border-radius: 0.3rem;overflow: hidden;position: absolute"
               :src="row.url ? row.url : ''" />
          <router-link style="margin-left: 0.2rem;"
                       :to="addressRoute(row.operatorAddress)"
                       class="link_style">{{row.moniker}}</router-link>
        </div>
      </template>
      <template slot-scope="{ row }"
                slot="operatorAddress">
        <span class="remove_default_style">
          <router-link :to="addressRoute(row.operatorAddress)"
                       class="link_style operator_address_style" style="font-family: Consolas,Menlo;color:#171d44;">{{formatAddress(row.operatorAddress)}}</router-link>
        </span>
      </template>
    </m-table>
    <no-data :fl-show-no-data="true" :no-data-doc="'no Validators'"></no-data>
  </div>
</template>

<script>
import Tools from '../../../util/Tools';
import NoData from "../../noDataComponent/NoData";

export default {
  name: 'MValidatorListTable',
  components: {NoData},
  props: {
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
      validatorFields: [],
      sortBy: 'votingPower',
      sortDesc: true,
      validatorData:[],
      activeValidatorFields: [
        {
          title: '#',
          key: 'index',
        },
        {
          title: 'Moniker',
          slot: 'moniker',
          sortable: true
        },
        {
          title: 'Operator',
          slot: 'operatorAddress',
          tooltip: true
        },
        {
          title: 'Commission',
          key: 'commission',
          sortable: true,
          sortMethod: this.sortMethodPer('commission'),
          className: 'text_right'
        },
        {
          title: 'Bonded_Tokens',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken'),
          className: 'text_right'
        },
        {
          title: 'Voting_Power',
          key: 'votingPower',
          sortable: true,
          sortMethod: this.sortMethodPer('votingPower'),
          className: 'text_right'
        },
        {
          title: 'Uptime',
          key: 'uptime',
          sortable: true,
          sortMethod: this.sortMethodPer('uptime'),
          className: 'text_right'
        },
        {
          title: 'Self-Bonded',
          key: 'selfBond',
          sortable: true,
          sortMethod: this.sortMethodSplit('selfBond'),
          className: 'text_right'
        },
        {
          title: 'Delegators',
          key: 'delegatorNum',
          sortable: true,
          sortMethod: this.sortMethodNumber('delegatorNum'),
          className: 'text_right'
        },
        {
          title: 'Bond_Height',
          key: 'bondHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('bondHeight'),
          className: 'text_right'
        }
      ],
      jailedValidatorFields: [
          {
              title: '#',
              key: 'index',
          },
        {
          title: 'Moniker',
          slot: 'moniker',
          sortable: true
        },
        {
          title: 'Operator',
          slot: 'operatorAddress',
          tooltip: true
        },
        {
          title: 'Commission',
          key: 'commission',
          sortable: true,
          sortMethod: this.sortMethodPer('commission'),
          className: 'text_right'
        },
        {
          title: 'Bonded_Tokens',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken'),
          className: 'text_right'
        },
        {
          title: 'Self_Bonded',
          key: 'selfBond',
          sortable: true,
          sortMethod: this.sortMethodSplit('selfBond'),
          className: 'text_right'
        },
        {
          title: 'Bond_Height',
          key: 'bondHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('bondHeight'),
          className: 'text_right'
        },
        {
          title: 'Unbonding_Height',
          key: 'unbondingHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('unbondingHeight'),
          className: 'text_right'
        }
      ],
      candidateValidatorFields: [
          {
              title: '#',
              key: 'index',
          },
        {
          title: 'Moniker',
          slot: 'moniker',
          sortable: true
        },
        {
          title: 'Operator',
          slot: 'operatorAddress',
          tooltip: true
        },
        {
          title: 'Commission',
          key: 'commission',
          sortable: true,
          sortMethod: this.sortMethodPer('commission'),
          className: 'text_right'
        },
        {
          title: 'Bonded_Tokens',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken'),
          className: 'text_right'
        },
        {
          title: 'Self-Bonded',
          key: 'selfBond',
          sortable: true,
          sortMethod: this.sortMethodSplit('selfBond'),
          className: 'text_right'
        },
        {
          title: 'Delegators',
          key: 'delegatorNum',
          sortable: true,
          sortMethod: this.sortMethodNumber('delegatorNum'),
          className: 'text_right'
        },
        {
          title: 'Bond_Height',
          key: 'bondHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('bondHeight'),
          className: 'text_right'
        },
        {
          title: 'Unbonding_Height',
          key: 'unbondingHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('unbondingHeight'),
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
    },
    setValidatorField (status) {
      switch (status) {
        case 'validator':
          this.validatorFields = this.activeValidatorFields;
          break;
        case 'candidate':
          this.validatorFields = this.candidateValidatorFields;
          break;
        case 'jailed':
          this.validatorFields = this.jailedValidatorFields;
          break;
      }
    }
  },
  watch:{
    items(items){
      if(items){
        this.items.map((item,i)=>{
          return item.index = i + 1
        })
      }
      this.validatorData = this.items;
    }
  },
  mounted () {
    this.validatorFields = this.activeValidatorFields;
  }
}
</script>

<style lang="scss">
.operator_address_style{
  font-family: "Consolas","Courier","Menlo","Arial",-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
}
</style>
