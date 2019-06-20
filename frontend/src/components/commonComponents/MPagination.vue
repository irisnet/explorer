<template>
  <div>
    <span v-show="total > 0"
          class="info">{{range[0]}}-{{range[1]}} of {{total}}</span>
    <!-- <button @click="toFrist" :class="[currentPage === 1 ? 'no_disable' : '']">«</button> -->
    <button @click="prev"
            :class="[currentPage === 1 || total === 0 ? 'no_disable' : '']">‹</button>
    <button @click="after"
            :class="[range[1] === 1 || total === 0 ? 'no_disable' : '']">›</button>
    <!-- <button @click="toEnd" :class="[currentPage === totalPages ? 'no_disable' : '']">»</button> -->
  </div>
</template>

<script>
export default {
  name: 'MPagination',
  props: {
    page: {
      type: Number,
      default: 1
    },
    total: {
      type: Number,
      default: 100
    },
    pageSize: {
      type: Number,
      default: 30
    },
    pageChange: {
      type: Function,
      default: function () { }
    },
    range: {
      type: Array,
      default: function () { return [] }
    }
  },
  data () {
    return {
      currentPage: 1
    }
  },
  computed: {
    totalPages () {
      return Math.ceil(this.total / this.pageSize);
    },
    from () {
      return ((this.currentPage - 1) * this.pageSize + 1);
    },
    to () {
      let to = this.currentPage * this.pageSize;
      if (to < this.total) {
        return to;
      } else {
        return this.total;
      }
    }
  },
  methods: {
    toFrist () {
      this.currentPage = 1;
    },
    toEnd () {
      this.currentPage = this.totalPages;
    },
    prev () {
      if (this.currentPage > 1) {
        this.currentPage = this.currentPage - 1;
      }
    },
    after () {
      if (this.currentPage < this.totalPages && this.range[1] && this.range[1] > 1) {
        this.currentPage = this.currentPage + 1;
      }
    }
  },
  watch: {
    page (newVal) {
      this.currentPage = newVal;
    },
    currentPage (newVal) {
      this.pageChange(newVal);
    }
  },
  mounted () {
    this.currentPage = this.page;
  }
}
</script>

<style lang="scss" scoped>
.info {
  color: #a2a2ae;
  vertical-align: middle;
  margin-right: 0.1rem;
}
button {
  outline: none;
  border: 1px solid #dee2e6;
  border-right-width: 0;
  padding: 5px 7.5px;
  color: #3598db;
  line-height: 17px;
  background-color: #ffffff;
  cursor: pointer;
  &:hover {
    background-color: #e9ecef;
  }
  &:focus {
    outline: none;
  }
}
button:nth-child(1) {
  border-top-left-radius: 0.025rem;
  border-bottom-left-radius: 0.025rem;
}
button:nth-last-child(1) {
  border-right-width: 1px;
  border-top-right-radius: 0.025rem;
  border-bottom-right-radius: 0.025rem;
}
button.no_disable {
  cursor: auto;
  color: #c0c4cc;
  &:hover {
    background-color: #ffffff;
  }
}
</style>
