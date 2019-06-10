<template>
  <div :style="{'min-width': '12.8rem'}">
    <m-table :columns="validatorFields" class="top_border_table"
             :data="items"
             :sort-by.sync="sortBy"
             :sort-desc.sync="sortDesc">
      <template slot-scope="{ row }"
                slot="id">
        <router-link :to="`/ProposalsDetail/${row.id}`"
                      class="link_style">{{row.id}}</router-link>
      </template>
      <template slot-scope="{ row }"
                slot="status">
        <img class="status_icon" v-if="row.status === 'Passed'" src="../../assets/pass.png" />
        <img class="status_icon" v-if="row.status === 'Rejected'" src="../../assets/rejected.png" />
        <img class="status_icon" v-if="row.status === 'VotingPeriod'" src="../../assets/voting_period.png" />
        <img class="status_icon" v-if="row.status === 'DepositPeriod'" src="../../assets/deposit_period.png" />
        <span>{{row.status}}</span>
      </template>
      <template slot-scope="{ row }"
                slot="type">
          <img class="status_icon" v-if="row.level === 'Important'" src="../../assets/important.png" />
          <img class="status_icon" v-if="row.level === 'Normal'" src="../../assets/normal.png" />
          <img class="status_icon" v-if="row.level === 'Critical'" src="../../assets/critical.png" />
          <span>{{row.type}}</span>
      </template>
      <template slot-scope="{ row }"
                slot="votes">
        <div class="votes_per_content" v-if="row.finalVotes">
          <div class="votes_per" :style="{backgroundColor: '#45B4FF', width: `${row.finalVotes.yes}%`}">
            <div class="tooltip_span">
              Yes: {{row.finalVotes.yes}}%
            </div>
          </div>
          <div class="votes_per" :style="{backgroundColor: '#CCDCFF', width: `${row.finalVotes.abstain}%`}">
            <div class="tooltip_span">
              Abstain: {{row.finalVotes.abstain}}%
            </div>
          </div>
          <div class="votes_per" :style="{backgroundColor: '#FFCF65', width: `${row.finalVotes.no}%`}">
            <div class="tooltip_span">
            No: {{row.finalVotes.no}}%
            </div>
          </div>
          <div class="votes_per" :style="{backgroundColor: '#FE8A8A', width: `${row.finalVotes.no_with_veto}%`}">
            <div class="tooltip_span">
              NoWithVeto: {{row.finalVotes.no_with_veto}}%
            </div>
          </div>
        </div>
      </template>
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
          width: 80
        },
        {
          title: 'Title',
          key: 'title'
        },
        {
          title: 'Type',
          slot: 'type'
        },
        {
          title: 'Status',
          slot: 'status'
        },
        {
          title: '',
          slot: 'votes'
        },
        {
          title: 'Submit_Time',
          key: 'submitTime',
          // sortable: true,
          sortMethod: this.sortMethodSplit('bondedToken'),
          className: 'text_right'
        },
        {
          title: 'Deposit_Endtime',
          key: 'depositEndTime',
          // sortable: true,
          sortMethod: this.sortMethodPer('votingPower'),
          className: 'text_right'
        },
        {
          title: 'Voting_Endtime',
          key: 'votingEndTime',
          // sortable: true,
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
  .status_icon {
    width: 14px;
    height: 14px;
    margin: 8px 0;
    vertical-align: middle;
    margin-right: 10px;
  }
  .votes_per_content {
    width: 102px;
    height: 10px;
    display: flex;
    .votes_per {
      height: 100%;
      cursor: pointer;
      position: relative;
      &:hover .tooltip_span {
        display: block;
      }
      .tooltip_span {
        display: none;
        position: absolute;
        z-index: 1000;
        bottom: calc(100% + 4px);
        left: 50%;
        transform: translateX(-50%);
        margin-top: -10px auto 0;
        padding: 5px 15px;
        color: #ffffff;
        background-color: #000000;
        line-height: 35px;
        border-radius: 0.04rem;
        word-wrap: break-word;
        white-space: nowrap;
        line-height: 1.7;
        &::after {
          width: 0;
          height: 0;
          border: 0.04rem solid transparent;
          content: "";
          display: block;
          position: absolute;
          border-top-color: #000000;
          left: 50%;
          margin-left: -4px;
          bottom: -8px;
        }
      }
    }
  }
</style>
