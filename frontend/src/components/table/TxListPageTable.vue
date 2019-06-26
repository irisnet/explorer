<template>
    <div :class="showNoData?'show_no_data':''" style="min-width: 12.8rem">
      <b-table :fields='fields' :items='items' striped>
        <template slot='Tx_Hash' slot-scope='data'>
          <span class="skip_route" style="display: flex">
            <div class="hash_container">
              <span>
                <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.Tx_Hash ? `${formatTxHash(String(data.item.Tx_Hash))}` : ''}}</router-link>
              </span>
              <span class="hash_content">{{data.item.Tx_Hash}}</span>
            </div>
          </span>
        </template>
        <template slot='Proposal_ID' slot-scope='data'>
          <span class="skip_route" v-if="data.item.Proposal_ID && data.item.Proposal_ID !== '--'">
            <router-link :to="`/ProposalsDetail/${data.item.Proposal_ID}`">{{data.item.Proposal_ID}}</router-link>
          </span>
          <span v-if="data.item.Proposal_ID && data.item.Proposal_ID === '--'">--</span>
        </template>
        <template slot='Proposal_Title' slot-scope='data'>
          <span class="skip_route" v-if="data.item.Proposal_ID !== '--' && data.item.Proposal_Title && data.item.Proposal_Title !== '--'">
            <router-link :to="`/ProposalsDetail/${data.item.Proposal_ID}`">{{data.item.Proposal_Title}}</router-link>
          </span>
          <span v-if="data.item.Proposal_ID === '--' && data.item.Proposal_Title && data.item.Proposal_Title !== '--'">{{data.item.Proposal_Title}}</span>
          <span v-if="data.item.Proposal_Title && data.item.Proposal_Title === '--'">--</span>
        </template>
        <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
        </template>
        <template slot='From' slot-scope='data'>
        <span v-if="(/^[1-9]\d*$/).test(data.item.From)" class="skip_route">
           <router-link :to="`/tx?txHash=${data.item.Tx_Hash}`">{{data.item.From}} Validators</router-link>
        </span>
        <span class="skip_route" style="display: flex" v-if="!(/^[0-9]\d*$/).test(data.item.From) && data.item.From !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.From)" class="link_style justify">{{formatAddress(data.item.From)}}</router-link>
            </span>
            <span class="address">{{data.item.From}}</span>
          </div>
        </span>
            <span class="no_skip" v-show="(/^[0]\d*$/).test(data.item.From) || data.item.From === '--'">--</span>
        </template>
        <template slot='OperatorAddr' slot-scope='data'>
        <span class="skip_route" style="display: flex" v-if="data.item.OperatorAddr !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.OperatorAddr)" class="link_style justify">{{formatAddress(data.item.OperatorAddr)}}</router-link>
            </span>
            <span class="address">{{data.item.OperatorAddr}}</span>
          </div>
        </span>
          <span class="no_skip" v-show="data.item.OperatorAddr === '--'">--</span>
        </template>
        <template slot='To' slot-scope='data'>
        <span class="skip_route" style="display: flex" v-if="data.item.To !== '--'">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.To)" class="link_style">{{formatAddress(data.item.To)}}</router-link>
            </span>
            <span class="address">{{data.item.To}}</span>
          </div>
        </span>
          <span class="no_skip" v-show="data.item.To === '--'">
          --
        </span>
        </template>
        <template slot='Owner' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="addressRoute(data.item.Owner)">{{data.item.Owner?`${String(data.item.Owner).substr(0,16)}...`:''}}</router-link>
        </span>
        </template>
        <template slot='Moniker' slot-scope='data'>
        <span v-show="data.item.Moniker && data.item.Moniker !== '--' ">
            <router-link :to="addressRoute(data.item.OperatorAddr)" class="skip_route">
              <pre class="pre_global_style moniker_link_style">{{data.item.Moniker}}</pre>
            </router-link>
        </span>
          <span v-show="data.item.Moniker && data.item.Moniker === '--' ">--</span>
        </template>
        <template slot='Tx_Signer' slot-scope='data'>
        <span class="skip_route" style="display: flex" v-if="data.item.Tx_Signer">
          <div class="name_address">
            <span class="remove_default_style">
              <router-link :to="addressRoute(data.item.Tx_Signer)" class="link_style justify">{{formatAddress(data.item.Tx_Signer)}}</router-link>
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
      name: "TxListPageTable",
      props:['items','showNoData','minWidth'],
      data () {
          return {
            fields:null,
            transferFields:{
              'Tx_Hash':{
                label:'Tx_Hash'
              },
              'Block':{
                label:'Block'
              },
              'From':{
                label:'From'
              },
              'Amount':{
                label:'Amount'
              },
              'To':{
                label:'To'
              },
              'Tx_Type':{
                label:'Tx_Type'
              },
              'Tx_Fee':{
                label:'Tx_Fee'
              },
              'Tx_Signer':{
                label:'Tx_Signer'
              },
              'Tx_Status':{
                label:'Tx_Status'
              },
              'Timestamp':{
                label:'Timestamp'
              },
            },
            declarationFields:{
              'Tx_Hash':{
                label:'Tx_Hash'
              },
              'Block':{
                label:'Block'
              },
              'Moniker':{
                label:'Moniker'
              },
              'OperatorAddr':{
                label:'Operator_Address'
              },
              'Amount':{
                label:'Self_Bonded'
              },
              'Tx_Type':{
                label:'Tx_Type'
              },
              'Tx_Fee':{
                label:'Tx_Fee'
              },
              'Tx_Signer':{
                label:'Tx_Signer'
              },
              'Tx_Status':{
                label:'Tx_Status'
              },
              'Timestamp':{
                label:'Timestamp'
              },
            },
            stakeFields:{
              'Tx_Hash':{
                label:'Tx_Hash'
              },
              'Block':{
                label:'Block'
              },
              'From':{
                label:'From'
              },
              'Amount':{
                label:'Amount'
              },
              'To':{
                label:'To'
              },
              'Tx_Type':{
                label:'Tx_Type'
              },
              'Tx_Fee':{
                label:'Tx_Fee'
              },
              'Tx_Signer':{
                label:'Tx_Signer'
              },
              'Tx_Status':{
                label:'Tx_Status'
              },
              'Timestamp':{
                label:'Timestamp'
              },
            },
            govFields:{
              'Tx_Hash':{
                label:'Tx_Hash'
              },
              'Block':{
                label:'Block'
              },
              'Proposal_Type':{
                label:'Proposal_Type'
              },
              'Proposal_ID':{
                label:'Proposal_ID'
              },
              'Proposal_Title':{
                label:'Proposal_Title'
              },
              'Amount':{
                label:'Amount'
              },
              'Tx_Type':{
                label:'Tx_Type'
              },
              'Tx_Fee':{
                label:'Tx_Fee'
              },
              'Tx_Signer':{
                label:'Tx_Signer'
              },
              'Tx_Status':{
                label:'Tx_Status'
              },
              'Timestamp':{
                label:'Timestamp'
              },
            },
          }
      },
      watch:{
        showNoData(showNoData){
        },
        items(items){
          this.setTxFields(items)
        }
      },
      methods:{
        formatAddress(address){
          return Tools.formatValidatorAddress(address)
        },
        formatTxHash(TxHash){
          if(TxHash){
            return Tools.formatTxHash(TxHash)
          }
        },
        setTxFields(items){
          items.forEach( (tx) => {
            if(tx.listName === 'transfer'){
              this.fields = this.transferFields
            }else if(tx.listName === 'declarations') {
              this.fields = this.declarationFields
            }else  if(tx.listName === 'stakes'){
              this.fields = this.stakeFields
            }else if(tx.listName === 'gov'){
              this.fields = this.govFields
            }else {
              this.fields = []
            }
          })
        }
      }
	}
</script>

<style lang="scss">
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
  .moniker_link_style{
    color: #3598db !important;
  }
  .hash_container{
    display: flex;
    flex-direction: column;
    position: relative;
    font-family: "Consolas","Arial",-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
    .hash_content{
      display: none;
      position: absolute;
      left: 0;
      top: -0.38rem;
      color: #3598db;
      background: rgba(0,0,0,0.8);
      border-radius:0.04rem;
      z-index: 10;
    }
    &:hover{
      .hash_content{
        background: rgba(0,0,0,1);
        color: #fff;
        padding: 0.06rem 0.15rem 0 0.15rem;
        display: block;
        border-radius:0.04rem;
        font-size: 0.14rem;
        &::after{
          content: '';
          display: block;
          background: rgba(0,0,0,1);
          transform: rotate(45deg);
          width: 0;
          height: 0;
          border: 0.04rem solid transparent;
          position: relative;
          top: 0.03rem;
          z-index: 1;
          left: 0.16rem;
        }
      }
    }
  }
</style>
