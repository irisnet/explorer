<template>
  <div>
    <b-table :fields='fields' :items='items' striped v-if="type === '1'">
      <template slot='Height' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Height}`)">
          {{data.item.Height}}
        </span>
      </template>
      <template slot='Txn' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.Txn}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '2'">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.TxHash}}
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.Block}}
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.From}}
        </span>
      </template>
      <template slot='To' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.To}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '3' || type === '4'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.Address}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '5'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute()">
          {{data.item.Address}}
        </span>
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
    },
    methods:{
      skipRoute(path){
        console.log(path);
        this.$router.push(path);
      }
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
      .skip_route{
        color:#3190e8;
        cursor:pointer;
      }
    }
  }

</style>
