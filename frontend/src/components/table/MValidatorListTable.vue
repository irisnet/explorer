<template>
  <div style="min-width: 12.8rem;"
       ref="tableContainer">
    <m-table ref="mTable" class="override_mtable" :columns="validatorFields" :data="items" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc">
      <template slot-scope="{ row }"
                slot="moniker">
        <div>
          <img v-if="row.url"
                style="width: 0.3rem;height: 0.3rem;border-radius: 0.3rem;overflow: hidden;"
                :src="row.url ? row.url : ''" />
          <router-link style="margin-left: 0.2rem;"
                        :to="`/address/1/${row.operatorAddress}`"
                        class="link_style">{{row.moniker}}</router-link>
        </div>
      </template>
      <template slot-scope="{ row }"
                slot="operatorAddress">
        <span class="remove_default_style">
          <router-link :to="`/address/1/${row.operatorAddress}`"
            class="link_style">{{formatAddress(row.operatorAddress)}}</router-link>
        </span>
      </template>
    </m-table>
  </div>
</template>

<script>
import Tools from '../../util/Tools';

export default {
  name: 'MValidatorListTable',
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
      default: 12
    }
  },
  data () {
    return {
      resizeTime: null,
      validatorFields: [],
      sortBy: 'votingPower',
      sortDesc: true,
      activeValidatorFields: [
        {
          title: 'Moniker',
          slot: 'moniker',
          width: 190,
          sortable: true
        },
        {
          title: 'Operator_Address',
          slot: 'operatorAddress',
          tooltip: true
        },
        {
          title: 'Commission',
          key: 'commission',
          sortable: true,
          sortMethod: this.sortMethodPer('commission')
        },
        {
          title: 'Bonded Tokens',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken')
        },
        {
          title: 'Voting Power',
          key: 'votingPower',
          sortable: true,
          sortMethod: this.sortMethodPer('votingPower')
        },
        {
          title: 'Uptime',
          key: 'uptime',
          sortable: true,
          sortMethod: this.sortMethodPer('uptime')
        },
        {
          title: 'Self Bonded',
          key: 'selfBond',
          sortable: true,
          sortMethod: this.sortMethodSplit('selfBond')
        },
        {
          title: 'Delegator Number',
          key: 'delegatorNum',
          sortable: true,
          sortMethod: this.sortMethodNumber('delegatorNum')
        },
        {
          title: 'Bond Height',
          key: 'bondHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('bondHeight')
        }
      ],
      jailedValidatorFields: [
        {
          title: 'Moniker',
          slot: 'moniker',
          width: 190,
          sortable: true
        },
        {
          title: 'Operator_Address',
          slot: 'operatorAddress',
          tooltip: true
        },
        {
          title: 'Commission',
          key: 'commission',
          sortable: true,
          sortMethod: this.sortMethodPer('commission')
        },
        {
          title: 'Bonded Tokens',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken')
        },
        {
          title: 'Self Bonded',
          key: 'selfBond',
          sortable: true,
          sortMethod: this.sortMethodSplit('selfBond')
        },
        {
          title: 'Bond Height',
          key: 'bondHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('bondHeight')
        }
      ],
      candidateValidatorFields: [
        {
          title: 'Moniker',
          slot: 'moniker',
          width: 190,
          sortable: true
        },
        {
          title: 'Operator_Address',
          slot: 'operatorAddress',
          tooltip: true
        },
        {
          title: 'Commission',
          key: 'commission',
          sortable: true,
          sortMethod: this.sortMethodPer('commission')
        },
        {
          title: 'Bonded Tokens',
          key: 'bondedToken',
          sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken')
        },
        {
          title: 'Self Bonded',
          key: 'selfBond',
          sortable: true,
          sortMethod: this.sortMethodSplit('selfBond')
        },
        {
          title: 'Bonded Tokens',
          key: 'delegatorNum',
          sortable: true,
          sortMethod: this.sortMethodNumber('delegatorNum')
        },
        {
          title: 'Delegator Number',
          key: 'bondHeight',
          sortable: true,
          sortMethod: this.sortMethodNumber('bondHeight')
        }
      ]
    }
  },
  methods: {
    sortMethodPer(key) {
      return (a, b) => Number(a[key].split('%')[0]) - Number(b[key].split('%')[0]);
    },
    sortMethodNumber(key) {
      return (a, b) => Number(a[key]) - Number(b[key]);
    },
    sortMethodSplit(key) {
      return (a, b) => Number(a[key].split(' ')[0]) - Number(b[key].split(' ')[0]);
    },
    formatAddress (address) {
      return Tools.formatValidatorAddress(address);
    },
    setValidatorField(status) {
      switch(status) {
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
    },
    changeScroll (e) {
      let element = document.querySelector('.m-table-header');
      if (e.target.scrollLeft > 0) {
        element.style.left = -e.target.scrollLeft + 'px';
      } else {
        element.style.left = 'auto';
      }
    },
    windowRisize (node) {
      let element = document.querySelector('.m-table-body');
      if (node.offsetWidth === element.offsetWidth) {
        document.querySelector('.m-table-header').style.left = 'auto';
      }
    },
    bundEvents (method) {
      let node = this.$refs.tableContainer && this.$refs.tableContainer.parentNode;
      if (!node) { return };
      this.resizeFn = () => {
        clearTimeout(this.resizeTime);
        this.resizeTime = setTimeout(() => {
          this.windowRisize(node);
        }, 500);
      };
      if (node.addEventListener) {
        node.addEventListener('scroll', method, false);
        window.addEventListener('resize', this.resizeFn, false);
      } else if (node.attachEvent) {
        node.attachEvent('onscroll', method);
        window.attachEvent('onresize', this.resizeFn);
      }
    },
    unbundEvents (method) {
      let node = this.$refs.tableContainer && this.$refs.tableContainer.parentNode;
      if (!node) { return };
      if (node.removeEventListener) {
        node.removeEventListener('scroll', method, false);
        window.removeEventListener('resize', this.resizeFn, false);
      } else if (node.dettachEvent) {
        node.dettachEvent('onscroll', method);
        window.dettachEvent('onresize', this.resizeFn);
      }
    }
  },
  mounted () {
    this.validatorFields = this.activeValidatorFields;
    this.bundEvents(this.changeScroll);
  },
  beforeDestroy () {
    this.unbundEvents(this.changeScroll);
  }
}
</script>

<style lang="scss">
.override_mtable {
  .m-table-header {
    position: fixed;
    margin-top: -0.45rem;
    background-color: #ffffff;
  }
  .m-table-body{
    margin-top: 0.45rem;
  }
}
@media screen and (max-width: 910px) {
  .m-table-header {
    position: static !important;
    margin-top: 0rem!important;
  }
  .m-table-body{
    margin-top: -0.05rem!important;
  }
}
</style>