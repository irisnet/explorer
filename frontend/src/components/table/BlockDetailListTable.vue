<template>
  <div :class="showNoData?'show_no_data':''" style="min-width: 12rem;">
    <b-table :fields='listFields' :items='items' striped nodelabel class="block_style">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/tx?txHash=${data.item.TxHash}`">{{data.item['TxHash'] ? `${String(data.item.TxHash).substr(0,16)}...` : ''}}{{data.item.proposal_id}}</router-link>
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Block}`">{{data.item.Block}}</router-link>
        </span>
      </template>
      <template slot='Title' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/ProposalsDetail/${data.item.ProposalId}`">{{data.item.Title.length > 10 ?`${data.item.Title.substring(0,10)}...` : `${data.item.Title}`}}</router-link>
        </span>
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
      <template slot='consensus' slot-scope='data'>
        <div class="name_address" v-show="data.item.consensus && data.item.consensus !== '--'">
            <span class="remove_default_style" :class="data.item.consensus === $route.params.param?'no_skip':''">
              <router-link :to="data.item.consensus === $route.params.param ? '' : `/address/1/${data.item.consensus}`" class="link_style">{{formatAddress(data.item.consensus)}}</router-link>
            </span>
          <span class="address">{{data.item.consensus ? data.item.consensus : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.consensus == '--'">
          --
        </span>
      </template>
      <template slot='Tx_Initiator' slot-scope='data'>
        <div class="name_address" v-show="data.item.Tx_Initiator && data.item.Tx_Initiator !== '--'">
            <span class="remove_default_style" :class="data.item.Tx_Initiator === $route.params.param?'no_skip':''">
              <router-link :to="data.item.Tx_Initiator === $route.params.param ? '' : `/address/1/${data.item.Tx_Initiator}`" class="link_style">{{formatAddress(data.item.Tx_Initiator)}}</router-link>
            </span>
          <span class="address">{{data.item.To ? data.item.To : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.To == '--'">
          --
        </span>
      </template>
      <template slot='ProposalId' slot-scope='data'>
        <div>
          <span>
            <router-link :to="`/ProposalsDetail/${data.item.ProposalId}`" style="color:#3598db !important;">{{data.item.ProposalId}}</router-link>
          </span>
        </div>
      </template>
      <template slot='Owner' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/address/1/${data.item.Owner}`">{{data.item.Owner?`${formatAddress(data.item.Owner)}`:''}}</router-link>
        </span>
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
      <template slot='Consensus' slot-scope='data'>
      <div class="name_address" v-show="data.item.Consensus && data.item.Consensus !== '--'">
            <span class="remove_default_style" :class="data.item.Consensus === $route.params.param?'no_skip':''">
              <router-link :to="data.item.Consensus === $route.params.param ? '' : `/address/1/${data.item.Consensus}`" class="link_style">{{formatAddress(data.item.Consensus)}}</router-link>
            </span>
        <span class="address">{{data.item.Consensus ? data.item.Consensus : ''}}</span>
      </div>
      <span class="no_skip" v-show="data.item.Consensus == '--'">
          --
        </span>
    </template>
      <template slot='Proposer' slot-scope='data'>
        <div class="name_address" v-show="data.item.Proposer && data.item.Proposer !== '--'">
            <span class="remove_default_style" :class="data.item.Proposer === $route.params.param?'no_skip':''">
              <router-link :to="data.item.Proposer === $route.params.param ? '' : `/address/1/${data.item.Proposer}`" class="link_style">{{formatAddress(data.item.Proposer)}}</router-link>
            </span>
          <span class="address">{{data.item.Proposer ? data.item.Proposer : ''}}</span>
        </div>
        <span class="no_skip" v-show="data.item.Proposer == '--'">
          --
        </span>
      </template>
      <template slot='radio' slot-scope='data'>
        <div class="radio_container">
          <div class="radio_content" v-if="data.item.yesRadio && data.item.abstainRadio && data.item.noRadio && data.item.noWithVetoRadio">
            <div class="test_content" v-if="data.item.yesRadio !== NaN" :style="{'width':`${data.item.yesRadio !== NaN ? data.item.yesRadio : 0}%`,'background': '#3598DB',}">
              <span class="radio_toast">Yes: {{data.item.yesRadio}}%</span>
            </div>
            <div class="test_content" v-if="data.item.abstainRadio !== NaN" :style="{'width':`${data.item.abstainRadio !== NaN ? data.item.abstainRadio : 0}%`,'background': '#73D1FF',}">
              <span class="radio_toast">Abstain: {{100 - data.item.noRadio - data.item.noWithVetoRadio - data.item.yesRadio}}%</span>
            </div>
            <div class="test_content" v-if="data.item.noRadio !== NaN" :style="{'width':`${data.item.noRadio !== NaN ? data.item.noRadio : 0}%`,'background': '#F9CA18',}">
              <span class="radio_toast">No: {{data.item.noRadio}}%</span>
            </div>
            <div class="test_content" v-if="data.item.noWithVetoRadio !== NaN" :style="{'width':`${data.item.noWithVetoRadio !== NaN ? data.item.noWithVetoRadio : 0}%`,'background': '#F27777',}">
              <span class="radio_toast">NoWithVeto: {{data.item.noWithVetoRadio}}%</span>
            </div>
          </div>
        </div>
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
                txFields:['TxHash','From','To','Amount','Fee','Tx_Initiator','Tx_Type','Tags','Status','Timestamp'],
                governanceFields:{
                  radio:{
                    label:'',
                  },
                  Title:{
                    label:'Title',
                  },
                  'ProposalId':{
                    label:'Proposal ID',
                  },
                  'Type':{
                    label:'Type',
                  },
                  'Status':{
                    label:'Status',
                  },
                  'Proposer':{
                    label:'Proposer',
                  },
                  'Submit Time':{
                    label:'Submit Time',
                  },
                  'Voting Start Time':{
                    label:'Voting Start Time',
                  },
                  'Total Deposits':{
                    label:'Total Deposits',
                  },
                }
              }
          },
          mounted(){
          },
          methods:{
              formatAddress(address){
                  return Tools.formatValidatorAddress(address)
              },
              formatListName(items){
                  items.forEach( item => {
                      if(item.listName === 'tx'){
                          this.listFields = this.txFields;
                      }else if(item.listName === 'gov') {
                          this.listFields = this.governanceFields;
                      }else {
                          this.listFields = [];
                      }
                  })
              }
          }
    }
</script>

<style lang="scss">
  /*@import '../style/mixin.scss';*/
  @import '../../style/mixin.scss';

  //重置bootstrap-vue的表格样式
    table{
      td{
        max-width:2.2rem !important;
        overflow-wrap: break-word !important;
        word-wrap: break-word !important;
      }
  }
  .page-item{
    &:first-child, &:last-child{
      .page-link{
        @include borderRadius(0.025rem);
      }
    }
  }
  .count_show{
    visibility: visible;
  }
  .count_hidden{
    visibility: hidden;
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
  .radio_container{
    min-width: 2rem;
    display: flex;
    align-items: center;
    width: 100%;
    background: rgba(241, 242, 247, 1);
    .radio_content{
      display: flex;
      width: 100%;
      height: 0.2rem;
      align-items:center;
      background: rgba(241, 242, 247, 1);
    }
  }
  .test_content{
    height: 0.18rem;
    position: relative;
    .radio_toast{
      display: none;
      width: 1.3rem;
      background: #000;
      position: absolute;
      bottom: 0.25rem;
      height: 0.2rem;
      color: #fff;
      border-radius: 0.03rem;
      text-align: center;
      line-height: 0.2rem;
      font-size: 0.14rem;
      &::after{
        content: '';
        display: block;
        width: 0.06rem;
        height:0.06rem;
        position: relative;
        left: 0.03rem;
        bottom: 0.04rem;
        transform: rotate(45deg);
        background: rgba(0,0,0,1);
        z-index: -1;
      }
    }
  }
  .test_content:hover{
    height: 0.2rem;
    transform: scale(1.1,1.1);
    z-index: 10;
    .radio_toast{
      display: block;
    }
  }
  .table tbody tr{
    height: 0.3rem !important;
  }
</style>
