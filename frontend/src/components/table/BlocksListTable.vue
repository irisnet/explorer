<template>
  <div>
    <b-table :fields='fields' :items='items' striped v-if="type === '1'">
      <template slot='Height' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Height}`)">
          {{data.item.Height}}
        </span>
      </template>
      <template slot='Txn' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/recent_transactions/2/0`)">
          {{data.item.Txn}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '2'">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/tx?txHash=${data.item.TxHash}`)">
          {{data.item.TxHash}}
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Block}`)">
          {{data.item.Block}}
        </span>
      </template>
      <template slot='From' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.From}`)">
          {{data.item.From}}
        </span>
      </template>
      <template slot='To' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.To}`)">
          {{data.item.To}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '3' || type === '4'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.Address}`)">
          {{data.item.Address}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '5'">
      <template slot='Address' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/address/1/${data.item.Address}`)">
          {{data.item.Address}}
        </span>
      </template>
    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '6'" style="margin-bottom:0;">
      <template slot='TxHash' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/tx?txHash=${data.item.TxHash}`)">
          {{data.item.TxHash}}
        </span>
      </template>
      <template slot='Block' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item.Block}`)">
          {{data.item.Block}}
        </span>
      </template>

    </b-table>
    <b-table :fields='fields' :items='items' striped v-if="type === '7'" style="margin-bottom:0;">
      <template slot='Block Height' slot-scope='data'>
        <span class="skip_route" @click="skipRoute(`/blocks_detail/${data.item['Block Height']}`)">
          {{data.item['Block Height']}}
        </span>
      </template>
    </b-table>
  </div>

</template>

<script>

  export default {
    watch: {
      items(items) {

      },

    },
    data() {
      return {
        fields: [],
      }
    },
    props: ['items', 'type'],
    mounted() {

    },
    methods: {
      skipRoute(path) {
        console.log(path);
        this.$router.push(path);
      }
    }
  }
</script>
<style lang="scss">
  @import '../../style/mixin.scss';

  //重置bootstrap-vue的表格样式
  table {

    td {
      max-width: 20rem !important;
      overflow-wrap: break-word !important;
      .skip_route {
        color: #3190e8;
        cursor: pointer;
      }
    }
  }

</style>
