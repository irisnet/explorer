<template>
  <div style="min-width: 12.5rem;" ref="tableContainer">
    <i-table class="override_itable" :columns="validatorFields" :data="items">
      <template slot-scope="{ row }" slot="moniker">
        <div>
          <img v-if="row.url" style="width: 0.3rem;height: 0.3rem;border-radius: 0.3rem;overflow: hidden;" :src="row.url ? row.url : ''" />
          <router-link style="margin-left: 0.2rem;" :to="`/address/1/${row.operatorAddress}`" class="link_style">{{row.moniker}}</router-link>
        </div>
      </template>
      <template slot-scope="{ row }" slot="operatorAddress">
        <i-tooltip :content="row.operatorAddress" placement="top">
          <span class="remove_default_style">
            <router-link :to="`/address/1/${row.operatorAddress}`" class="link_style">{{formatAddress(row.operatorAddress)}}</router-link>
          </span>
        </i-tooltip>
      </template>
    </i-table>
  </div>
</template>

<script>
  import Tools from '../../util/Tools';

  export default {
    name: "IViewValidatorListTable",
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
    data() {
      return {
        resizeTime: null,
        validatorFields: [],
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
            width: 150
          },
          {
            title: 'Commission',
            key: 'commission',
            width: 120,
            sortable: true
          },
          {
            title: 'Bonded Tokens',
            key: 'bondedToken',
            minWidth: 140,
            sortable: true
          },
          {
            title: 'Voting Power',
            key: 'votingPower',
            width: 130,
            sortable: true
          },
          {
            title: 'Uptime',
            key: 'uptime',
            width: 90,
            sortable: true
          },
          {
            title: 'Self Bonded',
            key: 'selfBond',
            width: 130,
            sortable: true
          },
          {
            title: 'Delegator Number',
            key: 'delegatorNum',
            width: 160,
            sortable: true
          },
          {
            title: 'Bond Height',
            key: 'bondHeight',
            width: 130,
            sortable: true
          }
        ],
        jailedValidatorFields: [
          {
            title: 'Moniker',
            slot: 'moniker',
            sortable: true
          },
          {
            title: 'Operator_Address',
            slot: 'operatorAddress'
          },
          {
            title: 'Commission',
            key: 'commission',
            sortable: true
          },
          {
            title: 'Bonded Tokens',
            key: 'bondedToken',
            sortable: true
          },
          {
            title: 'Self Bonded',
            key: 'selfBond',
            sortable: true
          },
          {
            title: 'Bond Height',
            key: 'bondHeight',
            sortable: true
          }
        ],
        candidateValidatorFields: [
          {
            title: 'Moniker',
            slot: 'moniker',
            sortable: true
          },
          {
            title: 'Operator_Address',
            slot: 'operatorAddress'
          },
          {
            title: 'Commission',
            key: 'commission',
            sortable: true
          },
          {
            title: 'Bonded Tokens',
            key: 'bondedToken',
            sortable: true
          },
          {
            title: 'Self Bonded',
            key: 'selfBond',
            sortable: true
          },
          {
            title: 'Bonded Tokens',
            key: 'delegatorNum',
            sortable: true
          },
          {
            title: 'Delegator Number',
            key: 'bondHeight',
            sortable: true
          }
        ]
      }
    },
    methods: {
      formatAddress(address) {
        return Tools.formatValidatorAddress(address);
      },
      setValidatorFields(validatorList) {
        validatorList.forEach(item => {
          if (item && item.validatorStatus && item.validatorStatus === 'jailed') {
            this.validatorFields = this.jailedValidatorFields;
          } else if (item && item.validatorStatus && item.validatorStatus === 'validator') {
            this.validatorFields = this.activeValidatorFields;
          } else if (item && item.validatorStatus && item.validatorStatus === 'candidate') {
            this.validatorFields = this.candidateValidatorFields;
          }
        })
      },
      changeScroll(e) {
        let element = document.querySelector('.ivu-table-header');
        if (e.target.scrollLeft > 0) {
          element.style.left = -e.target.scrollLeft + 'px';
        } else {
          element.style.left = 'auto';
        }
      },
      windowRisize(node) {
        let element = document.querySelector('.ivu-table-body');
        if (node.offsetWidth === element.offsetWidth) {
          document.querySelector('.ivu-table-header').style.left = 'auto';
        }
      },
      bundEvents(method) {
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
      unbundEvents(method) {
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
    watch: {
      items(items) {
        this.setValidatorFields(items);
      }
    },
    mounted(){
      this.validatorFields = this.activeValidatorFields;
      this.bundEvents(this.changeScroll);
    },
    beforeDestroy() {
      this.unbundEvents(this.changeScroll);
    }
  }
</script>

<style lang="scss">
  .override_itable {
    border-left-width: 0!important; 
    & > .ivu-table {
      &:before {
        width: 0;
      }
      &:after {
        width: 0;
      }
    }
    .ivu-table-header {
      position: fixed;
      z-index: 1000;
    }
    .ivu-table-body {
      margin-top: 0.5rem;
    }
    table {
      thead {
        tr {
          height: 0.49rem;
          border-bottom: 0.02rem solid #3598db;
          th {
            padding: 0.075rem;
            border-bottom-width: 0;
            div.ivu-table-cell {
              padding: 0;
            }
          }
        }
      }
      tbody {
        tr {
          td {
            height: 0.45rem;
            padding: 0.075rem;
            div.ivu-table-cell {
              padding: 0;
            }
          }
          &:nth-of-type(2n) {
            td {
              background-color: #f6f6f6;
            }
          }
          &:nth-of-type(2n+1) {
            td {
              background-color: #ffffff;
            }
          }
        }
      }
    }
    .ivu-tooltip-popper {
      .ivu-tooltip-inner {
        background-color: #000000!important;
        max-width: 100%!important;
        padding-left: 0.15rem;
        padding-right: 0.15rem;
        font-size: 14px!important;
      }
      .ivu-tooltip-arrow {
        border-top-color: #000000!important;
      }
    }
  }
  @media screen and (max-width: 910px) {
    .ivu-table-header {
      position: static!important;
    }
    .ivu-table-body {
      margin-top: 0!important;
    }
  }
</style>
