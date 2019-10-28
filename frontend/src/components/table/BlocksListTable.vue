<template>
  <div :class="showNoData?'show_no_data':''"
       class="table_wrap"
       :style="`${minWidth?(`min-width:${minWidth}rem`):''}`">

    <b-table :borderless="true" :fields='blockFields'
             :items='items'
             striped
             v-if="type === '1'">
      <template slot='Height'
                slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Height}`">{{data.item.Height}}</router-link>
        </span>
      </template>
      <template slot='Txn'
                slot-scope='data'>
        <span>{{data.item.Txn}}</span>
      </template>
      <template slot='Age'
                slot-scope='data'>
        <span v-show="data.item.Age">{{data.item.Age}}</span>
      </template>
    </b-table>

    <b-table :borderless="true" :fields='fields'
             :items='items'
             striped
             v-if="type === 'Proposals'"
             class="show_trim">
      <template slot='Title'
                slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/ProposalsDetail/${data.item['Proposal ID']}`">
            <pre class="proposals-list">{{data.item['Title']}}</pre>
          </router-link>
        </span>
      </template>
      <template slot='Proposal ID'
                slot-scope='data'>
        <span class="skip_route"
              v-show="data.item['Proposal ID'] && data.item['Proposal ID'] !== '--'">
          <router-link :to="`/ProposalsDetail/${data.item['Proposal ID']}`">{{data.item['Proposal ID']}}</router-link>
        </span>
        <span v-show="data.item['Proposal ID'] && data.item['Proposal ID'] === '--'">--</span>
      </template>
      <template slot="SubmitTime"
                slot-scope='data'>
        <span v-show="data.item.SubmitTime">{{data.item.SubmitTime}}</span>
      </template>
    </b-table>

    <b-table :borderless="true" :fields='fields'
             :items='items'
             striped
             v-if="type === 'ProposalsDetail'"
             nodelabel
             class="proposal_detail_list">
      <template slot='Amount'
                slot-scope='data'>
        <div class="skip_route_gray"
             style="width: 200px;">
          {{data.item['Amount']}}
        </div>
      </template>
      <template slot='Voter'
                slot-scope='data'>
        <div class="skip_route">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.Voter.split(' ')[0])"
                           class="link_style justify">{{data.item.Voter.split(' ')[1]}}</router-link>
            </span>
          </div>
        </div>
      </template>
      <template slot='Vote_Option'
                slot-scope='data'>
        <div style="width: 100px;">
          {{data.item['Vote_Option']}}
        </div>
      </template>
      <template slot='Depositor'
                slot-scope='data'>
        <span v-if="(/^[1-9]\d*$/).test(data.item.Depositor)"
              class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.From}} Validators</router-link>
        </span>
        <span class="skip_route"
              style="display: flex"
              v-if="!(/^[0-9]\d*$/).test(data.item.Depositor) && data.item.Depositor !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.Depositor)"
                           class="link_style justify">{{formatAddress(data.item.Depositor)}}</router-link>
            </span>
            <span class="address proposals_address_content">{{data.item.Depositor}}</span>
          </div>
        </span>
        <span class="no_skip"
              v-show="(/^[0]\d*$/).test(data.item.Depositor) || data.item.Depositor === '--'">--</span>
      </template>
      <template slot='Type'
                slot-scope='data'>
        <div class="skip_route_gray"
             style="width: 240px;">
          {{data.item['Type']}}
        </div>
      </template>
      <template slot='Tx_Hash'
                slot-scope='data'>
        <div class="skip_route"
             style="display: flex; width: 240px;">
          <div class="hash_container">
            <span>
              <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.Tx_Hash ? `${formatTxHash(String(data.item.Tx_Hash))}` : ''}}</router-link>
            </span>
            <span class="hash_content proposals_hash_content">{{data.item.Tx_Hash}}</span>
          </div>
        </div>
      </template>
      <template slot='VoteTime'
                slot-scope='data'>
        <span v-show="data.item.VoteTime">{{data.item.VoteTime}}</span>
      </template>
    </b-table>

    <b-table :borderless="true" :fields='fields'
             :items='items'
             striped
             v-if="type === 'addressTxList'"
             nodelabel>
      <template v-slot:cell(Tx_Hash)="data">
        <div class="common_hover_parent"
             v-if="data.item.Tx_Hash">
          <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`"
                       class="link_style common_font_style">{{formatTxHash(data.item.Tx_Hash)}}
            <div class="hover_content">
              {{data.item.Tx_Hash}}
            </div>
          </router-link>
        </div>
      </template>
      <template v-slot:cell(Age)="data">
        <span v-show="data.item.Age">{{data.item.Age}}</span>
      </template>
      <template v-slot:cell(Block)="data">
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template v-slot:cell(From)="data">
        <span v-if="(/^[1-9]\d*$/).test(data.item.From)"
              class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.From}} Validators</router-link>
        </span>
        <div class="name_address"
             v-show="!(/^[0-9]\d*$/).test(data.item.From) && data.item.From && data.item.From !== '--'">
          <span class="remove_default_style"
                :class="data.item.From === $route.params.param?'no_skip':''">
            <router-link :to="addressRoute(data.item.From)"
                         class="link_style">{{formatMoniker(data.item.fromMoniker) || formatAddress(data.item.From)}}</router-link>
          </span>
          <span v-if="!data.item.fromMoniker" class="address">{{data.item.From ? data.item.From : ''}}</span>
        </div>
        <span class="no_skip"
              v-show="(/^[0]\d*$/).test(data.item.From) || data.item.From === '--'">--</span>
      </template>
      <template v-slot:cell(OperatorAddr)="data">
        <div class="name_address"
             v-show="data.item.OperatorAddr && data.item.OperatorAddr !== '--'">
          <span class="remove_default_style"
                :class="data.item.OperatorAddr === $route.params.param?'no_skip':''">
            <router-link :to="addressRoute(data.item.OperatorAddr)"
                         class="link_style">{{formatAddress(data.item.OperatorAddr)}}</router-link>
          </span>
          <span class="address">{{data.item.OperatorAddr ? data.item.OperatorAddr : ''}}</span>
        </div>
        <span class="no_skip"
              v-show="data.item.OperatorAddr === '--'">--</span>
      </template>
      <template v-slot:cell(To)="data">
        <div class="name_address"
             v-show="data.item.To && data.item.To !== '--'">
          <span class="remove_default_style"
                :class="data.item.To === $route.params.param?'no_skip':''">
            <router-link :to="addressRoute(data.item.To)"
                         class="link_style">{{formatMoniker(data.item.toMoniker) || formatAddress(data.item.To)}}</router-link>
          </span>
          <span v-if="!data.item.toMoniker" class="address">{{data.item.To ? data.item.To : ''}}</span>
        </div>
        <span class="no_skip"
              v-show="data.item.To == '--'">
          --
        </span>
      </template>
      <template v-slot:cell(Owner)="data">
        <span>
          {{data.item.Owner?`${formatAddress(data.item.Owner)}`:''}}
        </span>
      </template>
      <template v-slot:cell(Tx_Signer)="data">
        <span class="skip_route"
              style="display: flex"
              v-if="data.item.Tx_Signer">
          <div class="name_address">
            <span class="remove_default_style"
                  :class="data.item.Tx_Signer === $route.params.param?'no_skip':''">
              <router-link :to="addressRoute(data.item.Tx_Signer)"
                           class="link_style justify">{{formatAddress(data.item.Tx_Signer)}}</router-link>
            </span>
            <span class="address">{{data.item.Tx_Signer}}</span>
          </div>
        </span>
      </template>
    </b-table>
  </div>

</template>

<script>
import Tools from '../../util/Tools';
export default {
  watch: {
    items (items) {
      this.setTxFields(items)
    },
  },
  data () {
    return {
      blockFields: ['Height', 'Txn', 'Age', 'Precommit Validators', 'Voting Power'],
      fields: null,
      transferFields:[
        {
          key: 'Tx_Hash',
          label:'TxHash',
        },
        {
          key: 'Block',
          label: 'Block'
        },
        {
          key: 'From',
          label: 'From'
        },
        {
          key: 'Amount',
          label: 'Amount'
        },
        {
          key: 'tokenId',
          label: 'Token'
        },
        {
          key: 'To',
          label: 'To'
        },
        {
          key: 'Tx_Type',
          label: 'TxType'
        },
        {
          key: 'transferFee',
          label: 'Fee'
        },
        {
          key: 'Tx_Signer',
          label: 'Signer'
        },
        {
          key: 'Tx_Status',
          label: 'Status'
        },
        {
          key: 'Timestamp',
          label: 'Timestamp'
        },
      ],
      declarationFields: [
        {
          key: 'Tx_Hash',
          label:'TxHash',
        },
        {
          key: 'Block',
          label: 'Block'
        },
        {
          key: 'Moniker',
          label: 'Moniker'
        },
        {
          key: 'OperatorAddr',
          label: 'Operator'
        },
        {
          key: 'Amount',
          label: 'Self-Bonded'
        },
        {
          key: 'Tx_Type',
          label: 'TxType'
        },
        {
          key: 'Tx_Fee',
          label: 'Fee'
        },
        {
          key: 'Tx_Signer',
          label: 'Signer'
        },
        {
          key: 'Tx_Status',
          label: 'Status'
        },
        {
          key: 'Timestamp',
          label: 'Timestamp'
        },
      ],
      stakeFields: [
        {
          key: 'Tx_Hash',
          label:'TxHash',
        },
        {
          key: 'Block',
          label: 'Block'
        },
        {
          key: 'From',
          label: 'From'
        },
        {
          key: 'Amount',
          label: 'Amount'
        },
        {
          key: 'To',
          label: 'To'
        },
        {
          key: 'Tx_Type',
          label: 'TxType'
        },
        {
          key: 'Tx_Fee',
          label: 'Fee'
        },
        {
          key: 'Tx_Signer',
          label: 'Signer'
        },
        {
          key: 'Tx_Status',
          label: 'Status'
        },
        {
          key: 'Timestamp',
          label: 'Timestamp'
        },
      ],
      govFields: [
        {
          key: 'Tx_Hash',
          label:'TxHash',
        },
        {
          key: 'Block',
          label: 'Block'
        },
        {
          key: 'Proposal_Type',
          label: 'Proposal_Type'
        },
        {
          key: 'Proposal_ID',
          label: 'Proposal_ID'
        },
        {
          key: 'Proposal_Title',
          label: 'Proposal_Title'
        },
        {
          key: 'Amount',
          label: 'Amount'
        },
        {
          key: 'Tx_Type',
          label: 'TxType'
        },
        {
          key: 'Tx_Fee',
          label: 'Fee'
        },
        {
          key: 'Tx_Signer',
          label: 'Signer'
        },
        {
          key: 'Tx_Status',
          label: 'Status'
        },
        {
          key: 'Timestamp',
          label: 'Timestamp'
        },
      ],
    }
  },
  props: ['items', 'type', 'showNoData', 'minWidth', 'status'],
  methods: {
    formatAmount (amount) {
      return Tools.formatAmount(amount)
    },
    formatAddress (address) {
      return Tools.formatValidatorAddress(address)
    },
    formatTxHash (TxHash) {
      if (TxHash) {
        return Tools.formatTxHash(TxHash)
      }
    },
    formatMoniker (moniker) {
        if (!moniker) {
            return '';
        }
        return Tools.formatString(moniker,15,"...");
    },
    setTxFields (items) {
      if(Array.isArray(items)){
        items.forEach((tx) => {
          if (tx.listName === 'transfer') {
            this.fields = this.transferFields
          } else if (tx.listName === 'validations') {
            this.fields = this.declarationFields
          } else if (tx.listName === 'delegations') {
            this.fields = this.stakeFields
          } else if (tx.listName === 'gov') {
            this.fields = this.govFields
          } else {
            this.fields = []
          }
        })
      }
    }
  }
}
</script>
<style lang="scss">
@import "../../style/mixin.scss";
.name_address {
  display: inline-block;
  position: relative;
  font-family: "Consolas","Ayuthaya","Arial", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif,
    "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  .address {
    display: none;
    position: absolute;
    padding: 0rem 0.15rem;
    top: -0.36rem;
    left: 50%;
    transform: translateX(-50%);
    color: #fff;
    background: rgba(0, 0, 0, 1);
    border-radius: 0.04rem;
    z-index: 10;
    line-height: 32px;
    font-size: 0.14rem;
    &::after {
      width: 0;
      height: 0;
      border: 0.06rem solid transparent;
      content: "";
      display: block;
      position: absolute;
      border-top-color: #000000;
      left: 50%;
      margin-left: -6px;
    }
  }
  &:hover {
    .address {
      display: block;
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
      a {
        color: var(--bgColor) !important;
        cursor: pointer;
      }
    }
    .no_skip {
      color: var(--titleColor);
      cursor: default;
      .link_style {
        color: var(--titleColor) !important;
      }
    }
  }
}
.common_font_style {
  .hover_content {
    display: none;
    position: absolute;
    padding: 0rem 0.15rem;
    top: -0.36rem;
    color: #fff;
    background: rgba(0, 0, 0, 1);
    border-radius: 0.04rem;
    z-index: 1;
    line-height: 32px;
    font-size: 0.14rem;
  }
  &:hover {
    .hover_content {
      display: block;
      &::after {
        width: 0;
        height: 0;
        border: 0.06rem solid transparent;
        content: "";
        display: block;
        position: absolute;
        border-top-color: #000000;
        left: 30px;
        margin-left: -6px;
      }
    }
  }
}

.page-link {
  padding: 0.05rem 0.075rem !important;
  height: 0.29rem !important;
  color: var(--bgColor) !important;
}
.active {
  .page-link {
    background-color: var(--bgColor) !important;
    border-color: var(--bgColor) !important;
    color: #fff !important;
  }
}
.table {
  th,
  td {
    padding: 0.075rem !important;
    color: var(--titleColor);
    @include fontWeight;
  }
  margin-bottom: 0 !important;
  thead {
    th {
      border-bottom: none !important;
    }
    tr {
      th {
        color: var(--contentColor)!important;
        height: 0.5rem;
        vertical-align: middle;
        &:first-child {
          padding-left: 0.2rem !important;
        }
      }
      border-bottom: 0.01rem solid var(--bgColor) !important;
    }
  }
  tbody {
    tr:nth-child(1) {
      td {
        border-top: none;
      }
    }
    tr {
      &:nth-of-type(even) {
        background-color: #f6f6f6 !important;
      }
      &:nth-of-type(odd) {
        background-color: #fff !important;
      }
      &:last-child {
        border-bottom: 1px solid #dee2e6;
      }
    }
  }
}
.show_no_data {
  .table {
    tbody {
      tr {
        &:first-child {
          display: none;
        }
      }
    }
  }
}
.proposal_detail_list tr {
  th {
    white-space: nowrap !important;
  }
  th:nth-child(1) {
    width: 50% !important;
  }
  th:nth-child(2) {
    width: 35% !important;
  }
}
//使用rem设置max-width不生效
@media screen and (max-width: 910px) {
  .proposal_detail_list tr {
    th {
      white-space: nowrap !important;
    }
    th:nth-child(1) {
      width: 50% !important;
    }
    th:nth-child(2) {
      width: auto !important;
    }
  }
}
.proposals-list {
  color: var(--bgColor);
  cursor: pointer;
  margin: 0 !important;
  padding: 0 !important;
}
.remove_default_style {
  margin: 0 !important;
  padding: 0 !important;
  color: var(--bgColor);
}
.show_trim td {
  white-space: pre;
}
.show_trim td span {
  white-space: pre;
}
.show_trim thead tr th {
  outline: transparent;
  &::after {
    margin-bottom: 0.13rem;
  }
}
.show_trim tbody tr td:nth-child(1) {
  padding: 0.075rem !important;
  width: 1.9rem;
}
.block_style thead tr th:nth-child(1) {
  width: 16%;
}
.block_style thead tr th:nth-child(2) {
  width: 8%;
}
.pre_global_style {
  font-size: 0.14rem;
  color: var(--contentColor);
  margin: 0;
}
.address_tx_list_table{
  table{
    td{
      color: var(--titleColor)!important;
       .no_skip{
         .link_style{
          color: var(--titleColor) !important;
         }
       }
    }
  }
}
.proposals_detail_table_wrap {
  tbody tr td:last-child {
    min-width: 2rem;
  }
}
pre {
  font-family: Arial !important;
  margin: 0;
}
.link_style {
  color: var(--bgColor) !important;
}
.proposals_hash_content {
  transform: translateX(-50%);
  left: 50% !important;
  &::after {
    left: 50% !important;
    transform: translateX(-50%);
  }
}
.proposals_address_content {
  left: 0 !important;
  &::after {
    left: 7px !important;
  }
}
</style>
