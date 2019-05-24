<template>
  <div :class="showNoData?'show_no_data':''" style="min-width: 12rem;" class="validator_table">
    <b-table :fields='listFields' :items='items' striped nodelabel>
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item['TxHash'] ? `${formatTxHash(String(data.item.TxHash))}` : ''}}</router-link>
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='ProposalTitle' slot-scope='data'>
        <span class="skip_route" v-if="data.item.ProposalId !== 0">
          <router-link :to="`/ProposalsDetail/${data.item.ProposalId}`">{{data.item.ProposalTitle.length > 10 ?`${data.item.ProposalTitle.substring(0,10)}...` : `${data.item.ProposalTitle}`}}</router-link>
        </span>
        <span v-if="data.item.ProposalId === 0">{{data.item.ProposalTitle}}</span>
      </template>
      <template slot='From' slot-scope='data'>
        <div class="name_address" v-show="data.item.From && data.item.From !== '--'">
            <span class="remove_default_style" :class="data.item.From === $route.params.param?'no_skip':''">
              <router-link :to="data.item.From === $route.params.param ? '' : `/address/1/${data.item.From}`" class="link_style">{{formatAddress(data.item.From)}}</router-link>
            </span>
          <span class="address">{{data.item.From ? data.item.From : ''}}</span>
        </div>
      </template>
      <template slot='To' slot-scope='data'>
        <div class="name_address" v-show="data.item.To && data.item.To !== '--'">
            <span class="remove_default_style" :class="data.item.To === $route.params.param?'no_skip':''">
              <router-link :to="data.item.To === $route.params.param ? '' : `/address/1/${data.item.To}`" class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
          <span class="address">{{data.item.To ? data.item.To : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.To == '--'">
          --
        </span>
      </template>
      <template slot='Consensus' slot-scope='data'>
        <div>{{data.item.Consensus}}</div>
      </template>
      <template slot='Tx_Initiator' slot-scope='data'>
        <div class="name_address" v-show="data.item.Tx_Initiator && data.item.Tx_Initiator !== '--'">
            <span class="remove_default_style" :class="data.item.Tx_Initiator === $route.params.param?'no_skip':''">
              <router-link :to="data.item.Tx_Initiator === $route.params.param ? '' : `/address/1/${data.item.Tx_Initiator}`" class="link_style">{{formatAddress(data.item.Tx_Initiator)}}</router-link>
            </span>
          <span class="address">{{data.item.Tx_Initiator ? data.item.Tx_Initiator : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.Tx_Initiator == '--'">
          --
        </span>
      </template>
      <template slot='ProposalId' slot-scope='data'>
        <div>
          <span v-show="data.item.ProposalId !== 0">
            <router-link :to="`/ProposalsDetail/${data.item.ProposalId}`" style="color:#3598db !important;">{{data.item.ProposalId}}</router-link>
          </span>
          <span v-show="data.item.ProposalId === 0">--</span>
        </div>
      </template>
      <template slot='OperatorAddress' slot-scope='data'>
        <div class="name_address" v-show="data.item.OperatorAddress && data.item.OperatorAddress !== '--'">
            <span class="remove_default_style" :class="data.item.OperatorAddress === $route.params.param?'no_skip':''">
              <router-link :to="data.item.OperatorAddress === $route.params.param ? '' : `/address/1/${data.item.OperatorAddress}`" class="link_style">{{formatAddress(data.item.OperatorAddress)}}</router-link>
            </span>
          <span class="address">{{data.item.OperatorAddress ? data.item.OperatorAddress : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.OperatorAddress == '--'">
          --
        </span>
      </template>
      <template slot='TxSigner' slot-scope='data'>
        <div class="name_address" v-show="data.item.TxSigner && data.item.TxSigner !== '--'">
            <span class="remove_default_style" :class="data.item.TxSigner === $route.params.param?'no_skip':''">
              <router-link :to="data.item.TxSigner === $route.params.param ? '' : `/address/1/${data.item.TxSigner}`" class="link_style">{{formatAddress(data.item.TxSigner)}}</router-link>
            </span>
          <span class="address">{{data.item.TxSigner ? data.item.TxSigner : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.TxSigner == '--'">
          --
        </span>
      </template>
      <template slot='Moniker' slot-scope='data'>
        <div class="moniker_conent">
          <div class="proposer_img_content">
            <img :style="{visibility:data.item.flProposer ? 'visible' : 'hidden'}" src="../../assets/proposer_img.png">
          </div>
          <span class="skip_route">
          <router-link :to="`/address/1/${data.item.OperatorAddress}`">{{data.item.Moniker? data.item.Moniker :''}}</router-link>
        </span>
        </div>
      </template>
      <template slot='index' slot-scope='data'>
        <span class="sequence_number_content">
          {{data.index + 1}}
        </span>
      </template>
    </b-table>
  </div>
</template>

<script>
  import Tools from '../../util/Tools';
      export default {
          name: "BlockDetailListTable.vue",
          props: ['items','governanceList','validatorSetList', 'validatorType','type','showNoData','minWidth','status'],
          watch:{
            items(items){
              this.formatListName(items)
            },
            governanceList(governanceList){
              this.formatListName(governanceList)
            },
            validatorSetList(validatorSetList){
            }
          },
          data () {
              return {
                  listFields: null,
                  txFields:{
                      TxHash:{
                        label:'Tx_Hash',
                      },
                      From:{
                        label:'From',
                      },
                      To:{
                        label:'To',
                      },
                      Amount:{
                        label:'Amount',
                      },
                      Fee:{
                        label:'Tx_Fee',
                      },
                      Tx_Initiator:{
                        label:'Tx_Signer',
                      },
                      Tx_Type:{
                        label:'Tx_Type',
                      },
                      Status:{
                        label:'Status',
                      },
                  },
                  governanceFields:{
                      'TxHash':{
                          label:'Tx_Hash',
                      },
                      'ProposalTitle':{
                        label:'Proposal_Title',
                      },
                      'ProposalId':{
                        label:'Proposal_ID',
                      },
                      'ProposalType':{
                        label:'Proposal_Type',
                      },
                      'TxFee':{
                          label:'Tx_Fee',
                      },
                      'TxSigner':{
                          label:'Tx_Signer',
                      },
                      'TxType':{
                          label:'Tx_Type',
                      },
                      'TxStatus':{
                          label:'Tx_Status',
                      },

                  },
                  validatorFields:{
                      index:{
                          label:'#'
                      },
                      'Moniker':{
                          label:'Moniker'
                      },
                      'OperatorAddress':{
                          label:'Operator Address'
                      },
                      'Consensus':{
                          label:'Consensus Address'
                      },
                      'ProposerPriority':{
                          label:'Proposer Priority'
                      },
                      'VotingPower':{
                          label:'Voting Power'
                      }
                  }
              }
          },
          mounted(){
          },
          methods:{
              formatAddress(address){
                if(address){
                  return Tools.formatValidatorAddress(address)
                }
              },
              formatTxHash(TxHash){
                if(TxHash){
                  return Tools.formatTxHash(TxHash)
                }
              },
              formatListName(items){
                items.forEach( item => {
                      if(item.listName === 'tx'){
                          this.listFields = this.txFields;
                      }else if(item.listName === 'gov') {
                          this.listFields = this.governanceFields;
                      }else {
                          this.listFields = this.validatorFields;
                      }
                  })
              }
          }
    }
</script>

<style lang="scss">
  @import '../../style/mixin.scss';
  .validator_table table{
      td{
        max-width:none !important;
        overflow-wrap: break-word !important;
        word-wrap: break-word !important;
      }
  }
  .show_no_data{
    .table{
      tbody{
        tr{
          &:first-child{
             display:none;
          }
        }
      }
    }
  }
  .table tbody tr{
    height: 0.3rem !important;
  }
  .page-item{
    &:first-child, &:last-child{
      .page-link{
        @include borderRadius(0.025rem);
      }
    }
  }
  .moniker_conent{
    display: flex;
    .proposer_img_content{
      width: 0.13rem;
      margin-right: 0.06rem;
      img{
        width: 100%;
      }
    }
  }
  .sequence_number_content{
    padding-left: 0.1rem;
  }
</style>
