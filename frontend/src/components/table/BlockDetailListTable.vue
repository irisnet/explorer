<template>
  <div :class="showNoData?'show_no_data':''"
       style="min-width: 12rem;"
       class="validator_table">
    <b-table :fields='listFields'
             :items='items'
             striped
             nodelabel
             :class="flIsValidatorTable ? 'validator_set_table_style' : ''">
      <template slot='Tx_Hash'
                slot-scope='data'>
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
      <template slot='Consensus'
                slot-scope='data'>
        <div>{{data.item.Consensus}}</div>
      </template>
      <template slot='Tx_Signer'
                slot-scope='data'>
        <div class="name_address"
             v-show="data.item.Tx_Signer && data.item.Tx_Signer !== '--'">
          <span class="remove_default_style"
                :class="data.item.Tx_Signer === $route.params.param?'no_skip':''">
            <router-link :to="data.item.Tx_Signer === $route.params.param ? '' : addressRoute(data.item.Tx_Signer)"
                         class="link_style">{{formatAddress(data.item.Tx_Signer)}}</router-link>
          </span>
          <span class="address">{{data.item.Tx_Signer ? data.item.Tx_Signer : ''}}</span>
        </div>
        <span class="no_skip"
              v-show="data.item.Tx_Initiator == '--'">
          --
        </span>
      </template>
      <template slot='OperatorAddress'
                slot-scope='data'>
        <div class="name_address"
             v-show="data.item.OperatorAddress && data.item.OperatorAddress !== '--'">
          <span class="remove_default_style"
                :class="data.item.OperatorAddress === $route.params.param?'no_skip':''">
            <router-link :to="data.item.OperatorAddress === $route.params.param ? '' : addressRoute(data.item.OperatorAddress)"
                         class="link_style">{{formatAddress(data.item.OperatorAddress)}}</router-link>
          </span>
          <span class="address">{{data.item.OperatorAddress ? data.item.OperatorAddress : ''}}</span>
        </div>
        <span class="no_skip"
              v-show="data.item.OperatorAddress == '--'">
          --
        </span>
      </template>
      <template slot='moniker'
                slot-scope='data'>
        <div class="moniker_conent">
          <div class="proposer_img_content">
            <img :style="{visibility:data.item.flProposer ? 'visible' : 'hidden'}"
                 src="../../assets/proposer_img.png">
          </div>
          <span class="skip_route">
            <router-link :to="addressRoute(data.item.OperatorAddress)">{{data.item.moniker? data.item.moniker :''}}</router-link>
          </span>
        </div>
      </template>
      <template slot='OperatorAddr'
                slot-scope='data'>
        <span class="skip_route"
              style="display: flex"
              v-if="data.item.OperatorAddr !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.OperatorAddr)"
                           class="link_style justify">{{formatAddress(data.item.OperatorAddr)}}</router-link>
            </span>
            <span class="address">{{data.item.OperatorAddr}}</span>
          </div>
        </span>
        <span class="no_skip"
              v-show="data.item.OperatorAddr === '--'">--</span>
      </template>
      <template slot='index'
                slot-scope='data'>
        <span class="sequence_number_content">
          {{data.index + 1}}
        </span>
      </template>
      <template slot='Proposal_ID'
                slot-scope='data'>
        <span class="skip_route"
              v-if="data.item.Proposal_ID && data.item.Proposal_ID !== '--'">
          <router-link :to="`/ProposalsDetail/${data.item.Proposal_ID}`">{{data.item.Proposal_ID}}</router-link>
        </span>
        <span v-if="data.item.Proposal_ID && data.item.Proposal_ID === '--'">--</span>
      </template>
      <template slot='Proposal_Title'
                slot-scope='data'>
        <span class="skip_route"
              v-if="data.item.Proposal_ID !== '--' && data.item.Proposal_Title && data.item.Proposal_Title !== '--' ">
          <router-link :to="`/ProposalsDetail/${data.item.Proposal_ID}`">{{data.item.Proposal_Title}}</router-link>
        </span>
        <span v-if="data.item.Proposal_ID === '--' && data.item.Proposal_Title && data.item.Proposal_Title !== '--'">{{data.item.Proposal_Title}}</span>
        <span v-if="data.item.Proposal_Title && data.item.Proposal_Title === '--'">--</span>
      </template>
      <template slot='Block'
                slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='From'
                slot-scope='data'>
        <span v-if="(/^[1-9]\d*$/).test(data.item.From)"
              class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.From}} Validators</router-link>
        </span>
        <span class="skip_route"
              style="display: flex"
              v-if="!(/^[0-9]\d*$/).test(data.item.From) && data.item.From !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.From)"
                           class="link_style justify">{{formatAddress(data.item.From)}}</router-link>
            </span>
            <span class="address">{{data.item.From}}</span>
          </div>
        </span>
        <span class="no_skip"
              v-show="(/^[0]\d*$/).test(data.item.From) || data.item.From === '--'">--</span>
      </template>
      <template slot='To'
                slot-scope='data'>
        <span class="skip_route"
              style="display: flex"
              v-if="data.item.To !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.To)"
                           class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
            <span class="address">{{data.item.To}}</span>
          </div>
        </span>
        <span class="no_skip"
              v-show="data.item.To === '--'">
          --
        </span>
      </template>
      <template slot='Owner'
                slot-scope='data'>
        <span class="skip_route">
          <router-link :to="addressRoute(data.item.Owner)">{{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}</router-link>
        </span>
      </template>
      <template slot='Moniker'
                slot-scope='data'>
        <span v-show="data.item.Moniker && data.item.Moniker !== '--' ">
          <router-link :to="addressRoute(data.item.OperatorAddr)"
                       class="skip_route">
            <pre class="pre_global_style moniker_link_style">{{data.item.Moniker}}</pre>
          </router-link>
        </span>
        <span v-show="data.item.Moniker && data.item.Moniker === '--' ">--</span>
      </template>
      <template slot='Tx_Signer'
                slot-scope='data'>
        <span class="skip_route"
              style="display: flex"
              v-if="data.item.Tx_Signer">
          <div class="name_address">
            <span class="remove_default_style">
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
  name: "BlockDetailListTable.vue",
  props: ['items', 'governanceList', 'validatorSetList', 'validatorType', 'type', 'showNoData', 'minWidth', 'status'],
  watch: {
    items (items) {
      this.formatListName(items)
    },
    governanceList (governanceList) {
      this.formatListName(governanceList)
    },
    validatorSetList (validatorSetList) {
    }
  },
  data () {
    return {
      listFields: null,
      transferFields: {
        'Tx_Hash': {
          label: 'Tx_Hash'
        },
        'From': {
          label: 'From'
        },
        'Amount': {
          label: 'Amount'
        },
        'To': {
          label: 'To'
        },
        'Tx_Type': {
          label: 'Tx_Type'
        },
        'Tx_Fee': {
          label: 'Tx_Fee'
        },
        'Tx_Signer': {
          label: 'Tx_Signer'
        },
        'Tx_Status': {
          label: 'Tx_Status'
        },
      },
      declarationFields: {
        'Tx_Hash': {
          label: 'Tx_Hash'
        },
        'Moniker': {
          label: 'Moniker'
        },
        'OperatorAddr': {
          label: 'Operator_Address'
        },
        'Amount': {
          label: 'Self_Bonded'
        },
        'Tx_Type': {
          label: 'Tx_Type'
        },
        'Tx_Fee': {
          label: 'Tx_Fee'
        },
        'Tx_Signer': {
          label: 'Tx_Signer'
        },
        'Tx_Status': {
          label: 'Tx_Status'
        },
      },
      stakeFields: {
        'Tx_Hash': {
          label: 'Tx_Hash'
        },
        'From': {
          label: 'From'
        },
        'Amount': {
          label: 'Amount'
        },
        'To': {
          label: 'To'
        },
        'Tx_Type': {
          label: 'Tx_Type'
        },
        'Tx_Fee': {
          label: 'Tx_Fee'
        },
        'Tx_Signer': {
          label: 'Tx_Signer'
        },
        'Tx_Status': {
          label: 'Tx_Status'
        },
      },
      govFields: {
        'Tx_Hash': {
          label: 'Tx_Hash'
        },
        'Proposal_Type': {
          label: 'Proposal_Type'
        },
        'Proposal_ID': {
          label: 'Proposal_ID'
        },
        'Proposal_Title': {
          label: 'Proposal_Title'
        },
        'Amount': {
          label: 'Amount'
        },
        'Tx_Type': {
          label: 'Tx_Type'
        },
        'Tx_Fee': {
          label: 'Tx_Fee'
        },
        'Tx_Signer': {
          label: 'Tx_Signer'
        },
        'Tx_Status': {
          label: 'Tx_Status'
        },
      },
      validatorFields: {
        index: {
          label: '#'
        },
        'moniker': {
          label: 'Moniker'
        },
        'OperatorAddress': {
          label: 'Operator Address'
        },
        'Consensus': {
          label: 'Consensus Address'
        },
        'ProposerPriority': {
          label: 'Proposer Priority'
        },
        'VotingPower': {
          label: 'Voting Power'
        }
      },
      flIsValidatorTable: false,
    }
  },
  mounted () {
  },
  methods: {
    formatAddress (address) {
      if (address) {
        return Tools.formatValidatorAddress(address)
      }
    },
    formatTxHash (TxHash) {
      if (TxHash) {
        return Tools.formatTxHash(TxHash)
      }
    },
    formatListName (items) {
      items.forEach((tx) => {
        if (tx.listName === 'transfer') {
          this.listFields = this.transferFields
        } else if (tx.listName === 'declarations') {
          this.listFields = this.declarationFields
        } else if (tx.listName === 'stakes') {
          this.listFields = this.stakeFields
        } else if (tx.listName === 'gov') {
          this.listFields = this.govFields
        } else {
          this.listFields = this.validatorFields;
          this.flIsValidatorTable = true
        }
      })
    }
  }
}
</script>

<style lang="scss">
@import "../../style/mixin.scss";
.validator_table table {
  td {
    max-width: none !important;
    overflow-wrap: break-word !important;
    word-wrap: break-word !important;
  }
}
.validator_table .validator_set_table_style thead tr th:nth-child(2) {
  padding-left: 0.26rem !important;
}
.validator_table {
  min-width: 12.8rem;
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
.table tbody tr {
  height: 0.3rem !important;
}
.page-item {
  &:first-child,
  &:last-child {
    .page-link {
      @include borderRadius(0.025rem);
    }
  }
}
.moniker_conent {
  display: flex;
  .proposer_img_content {
    width: 0.13rem;
    margin-right: 0.06rem;
    img {
      width: 100%;
    }
  }
}
.sequence_number_content {
  padding-left: 0.1rem;
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
    z-index: 10;
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
</style>
