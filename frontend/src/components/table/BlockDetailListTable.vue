<template>
  <div :class="showNoData?'show_no_data':''"
       style="min-width: 12rem; border-left: 1px solid #dee2e6; border-right: 1px solid #dee2e6;"
       class="validator_table">
    <b-table :borderless="true" :fields='listFields'
             :items='items'
             striped
             nodelabel
             :class="flIsValidatorTable ? 'validator_set_table_style' : ''">
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
      <template v-slot:cell(Consensus)="data">
        <div>{{data.item.Consensus}}</div>
      </template>
      <template v-slot:cell(Tx_Signer)="data">
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
      <template v-slot:cell(OperatorAddress)="data">
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
      <template v-slot:cell(moniker)="data">
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
      <template v-slot:cell(OperatorAddr)="data">
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
      <template v-slot:cell(index)="data">
        <span class="sequence_number_content">
          {{data.index + 1}}
        </span>
      </template>
      <template v-slot:cell(Proposal_ID)="data">
        <span class="skip_route"
              v-if="data.item.Proposal_ID && data.item.Proposal_ID !== '--'">
          <router-link :to="`/ProposalsDetail/${data.item.Proposal_ID}`">{{data.item.Proposal_ID}}</router-link>
        </span>
        <span v-if="data.item.Proposal_ID && data.item.Proposal_ID === '--'">--</span>
      </template>
      <template v-slot:cell(Proposal_Title)="data">
        <span class="skip_route"
              v-if="data.item.Proposal_ID !== '--' && data.item.Proposal_Title && data.item.Proposal_Title !== '--' ">
          <router-link :to="`/ProposalsDetail/${data.item.Proposal_ID}`">{{data.item.Proposal_Title}}</router-link>
        </span>
        <span v-if="data.item.Proposal_ID === '--' && data.item.Proposal_Title && data.item.Proposal_Title !== '--'">{{data.item.Proposal_Title}}</span>
        <span v-if="data.item.Proposal_Title && data.item.Proposal_Title === '--'">--</span>
      </template>
      <template v-slot:cell(Block)="data">
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template v-slot:cell(From)="data">
        <span v-if="(/^[1-9]\d*$/).test(data.item.From) && data.item.tokenId === 'IRIS'"
              class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.From}} Validators</router-link>
        </span>
        <span class="skip_route"
              style="display: flex"
              v-if="!(/^[0-9]\d*$/).test(data.item.From) && data.item.From !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.From)"
                           class="link_style justify">{{formatMoniker(data.item.fromMoniker) || formatAddress(data.item.From)}}</router-link>
            </span>
            <span class="address" v-if="!data.item.fromMoniker">{{data.item.From}}</span>
          </div>
        </span>
        <span class="no_skip"
              v-show="(/^[0]\d*$/).test(data.item.From) || data.item.From === '--'">--</span>
      </template>
      <template v-slot:cell(To)="data">
        <span class="skip_route"
              style="display: flex"
              v-if="data.item.To !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.To)"
                           class="link_style">{{formatMoniker(data.item.toMoniker) || formatAddress(data.item.To)}}</router-link>
            </span>
            <span class="address" v-if="!data.item.toMoniker">{{data.item.To}}</span>
          </div>
        </span>
        <span class="no_skip"
              v-show="data.item.To === '--'">
          --
        </span>
      </template>
      <template v-slot:cell(Owner)="data">
        <span class="skip_route">
          <router-link :to="addressRoute(data.item.Owner)">{{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}</router-link>
        </span>
      </template>
      <template v-slot:cell(Moniker)="data">
        <span v-show="data.item.Moniker && data.item.Moniker !== '--' ">
          <router-link :to="addressRoute(data.item.OperatorAddr)"
                       class="skip_route">
            <pre class="pre_global_style moniker_link_style">{{data.item.Moniker}}</pre>
          </router-link>
        </span>
        <span v-show="data.item.Moniker && data.item.Moniker === '--' ">--</span>
      </template>
      <template v-slot:cell(Tx_Signer)="data">
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
      transferFields: [
        {
          key:'Tx_Hash',
          label:'TxHash'
        },
        {
          key:'From',
          label:'From'
        },
        {
          key:'Amount',
          label:'Amount'
        },
        {
          key:'tokenId',
          label:'Token'
        },
        {
          key:'To',
          label:'To'
        },
        {
          key:'Tx_Type',
          label:'TxType'
        },
        {
          key:'transferFee',
          label:'Fee'
        },
        {
          key:'Tx_Signer',
          label:'Signer'
        },
        {
          key:'Tx_Status',
          label:'Status'
        },
      ],
      declarationFields: [
        {
          key:'Tx_Hash',
          label:'TxHash'
        },
        {
          key:'Moniker',
          label:'Moniker'
        },
        {
          key:'OperatorAddr',
          label:'Operator'
        },
        {
          key:'Amount',
          label:'Self-Bonded'
        },
        {
          key:'Tx_Type',
          label:'TxType'
        },
        {
          key:'Tx_Fee',
          label:'Fee'
        },
        {
          key:'Tx_Signer',
          label:'Signer'
        },
        {
          key:'Tx_Status',
          label:'Status'
        },
      ],
      stakeFields: [
        {
          key:'Tx_Hash',
          label:'TxHash'
        },
        {
          key:'From',
          label:'From'
        },
        {
          key:'Amount',
          label:'Amount'
        },
        {
          key:'To',
          label:'To'
        },
        {
          key:'Tx_Type',
          label:'TxType'
        },
        {
          key:'Tx_Fee',
          label:'Fee'
        },
        {
          key:'Tx_Signer',
          label:'Signer'
        },
        {
          key:'Tx_Status',
          label:'Status'
        },
      ],
      govFields: [
        {
          key:'Tx_Hash',
          label:'TxHash'
        },
        {
          key:'Proposal_Type',
          label:'Proposal_Type'
        },
        {
          key:'Proposal_ID',
          label:'Proposal_ID'
        },
        {
          key:'Proposal_Title',
          label:'Proposal_Title'
        },
        {
          key:'Amount',
          label:'Amount'
        },
        {
          key:'Tx_Type',
          label:'TxType'
        },
        {
          key:'Tx_Fee',
          label:'Fee'
        },
        {
          key:'Tx_Signer',
          label:'Signer'
        },
        {
          key:'Tx_Status',
          label:'Status'
        },
      ],
      validatorFields: [
        {
          key:'index',
          label: '#'
        },
        {
          key:'moniker',
          label: 'Moniker'
        },
        {
          key:'OperatorAddress',
          label: 'Operator'
        },
        {
          key:'Consensus',
          label: 'Consensus Address'
        },
        {
          key:'ProposerPriority',
          label: 'Proposer Priority'
        },
        {
          key:'VotingPower',
          label: 'Voting Power'
        },
      ],
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
    formatMoniker (moniker) {
      if (!moniker) {
        return '';
      }
      return Tools.formatString(moniker,13,"...");
    },
    formatListName (items) {
      Array.isArray(items) && items.forEach((tx) => {
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
    color:var(--titleColor)
  }
}
.validator_table .validator_set_table_style thead tr th:nth-child(2) {
  padding-left: 0.26rem !important;
}
.validator_table .validator_set_table_style thead {
  background: #fff;
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
</style>
