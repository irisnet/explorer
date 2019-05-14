<template>
    <div :class="showNoData?'show_no_data':''" class="list_table_container">
      <b-table :fields='blockFields' :items='items' striped>
        <template slot='Height' slot-scope='data'>
        <span class="skip_route">
          <router-link :to="`/block/${data.item.Height}`">{{data.item.Height}}</router-link>
        </span>
        </template>
        <template slot='Txn' slot-scope='data'>
          <span>{{data.item.Txn}}</span>
        </template>
        <template slot='Age' slot-scope='data'>
          <span v-show="data.item.Age">{{data.item.Age}}</span>
        </template>
      </b-table>
    </div>
</template>

<script>
	export default {
		name: "BlockListPageTable",
    props:['items','showNoData','minWidth'],
    watch: {
      items(items){
      }
    },
    data() {
      return {
        blockFields:['Height','Txn','Age','Precommit Validators','Voting Power']
      }
    },
	}
</script>

<style lang="scss">
  .list_table_container{
    min-width: 8rem;
  }
  table tbody tr td{
    font-size: 0.14rem;
  }
  table thead tr th{
    font-size: 0.14rem;
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
</style>
