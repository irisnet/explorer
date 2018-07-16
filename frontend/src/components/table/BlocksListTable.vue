<template>
  <div>
    <b-table :fields='fields' :items='items' striped v-if="type === '1'">
      <template slot='Height' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.Height}}
        </a>
      </template>
      <template slot='Txn' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.Txn}}
        </a>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '2'">
      <template slot='TxHash' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.TxHash}}
        </a>
      </template>
      <template slot='Block' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.Block}}
        </a>
      </template>
      <template slot='From' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.From}}
        </a>
      </template>
      <template slot='To' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.To}}
        </a>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '3' || type === '4'">
      <template slot='Address' slot-scope='data'>
        <a href="http://www.baidu.com">
          {{data.item.Address}}
        </a>
      </template>
    </b-table>
  </div>

</template>

<script>

  export default {

    data() {
      return {
        fields: [],
      }
    },
    props:['items','type'],
    mounted(){
      console.log(this.type)
    }
  }
</script>
<style lang="scss">
  @import '../../style/mixin.scss';


  //重置bootstrap-vue的表格样式
  table{

    td{
      max-width:20rem !important;
      overflow-wrap: break-word !important;
    }
  }

</style>
