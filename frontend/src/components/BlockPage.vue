<template>
  <b-table striped hover :items="blocks" :fields="fields"></b-table>
</template>

<script>
  import axios from 'axios'

  export default {
    name: "BlockPage",
    data() {
      return {
        fields: ["Height", "Hash", "Time", {key: "NumTxs", label: "Txn"}],
        items: [],
        blocks: [],
        count: 0,
        page: 0,
        size: 20
      }
    },
    created: function () {
      this.load()
    },
    methods: {
      load() {
        axios.get("http://localhost:8080" + "/blocks/" + this.page + "/" + this.size).then(result => {
          let data = result.data;
          this.blocks = data.Data;
          this.blocks.forEach(block)
          this.count = data.Count;
        })
      }
    }
  }
</script>

<style scoped>

</style>
