<template>
  <div :style="{opacity: show ? 1 : 0}">
    <div class="m-table-header">
      <table class="m_table" cellspacing="0"
             cellpadding="0"
             border="0" ref="tableHeader">
        <colgroup>
          <col v-for="(v, i) in colWidth" :key="i" :style="{width: (columns[i] && columns[i].width) ? columns[i].width+ 'px' : v + 'px', 
          minWidth: columns[i] && columns[i].minWidth + 'px', 
          maxWidth: columns[i] && columns[i].maxWidth + 'px'}"/>
        </colgroup>
        <thead>
          <tr>
            <th :class="v.sortable ? 'sorting' : ''" @click="sortDataByClick(v)" v-for="(v, i) in columns" :key="i">{{v.title}}
              <i class="sort" :class="{'desc': (v.key === sortAsBy || v.slot === sortAsBy) && sortAsDesc, 
              'asc': (v.key === sortAsBy || v.slot === sortAsBy) && !sortAsDesc}"></i>
            </th>
          </tr>
        </thead>
      </table>
    </div>
    <div class="m-table-body">
      <table class="m_table" cellspacing="0"
             cellpadding="0"
             border="0">
        <thead class="hidden_thead">
          <tr>
            <!-- <th v-for="(v, i) in columns" :key="i">{{v.title}}</th> -->
            <th :class="v.sortable ? 'sorting' : ''" @click="sortDataByClick(v)" v-for="(v, i) in columns" :key="i">{{v.title}}
              <i class="sort" :class="{'desc': (v.key === sortAsBy || v.slot === sortAsBy) && sortAsDesc, 
              'asc': (v.key === sortAsBy || v.slot === sortAsBy) && !sortAsDesc}"></i>
            </th>
          </tr>
        </thead>
        <tbody ref="table_body">
          <tr v-for="(v, i) in data" :key="i">
            <td v-for="(it, j) in columns" :width="it.width" :key="j">
              <template v-if="it.key">
                <div :class="{'tooltip_span_container': it.tooltip}">
                  {{v[it.key]}}
                  <span class="tooltip_span" v-if="it.tooltip">{{v[it.key || it.slot]}}</span>
                </div>
              </template>
              <template v-else>
                <div :class="{'tooltip_span_container': it.tooltip}">
                  <slot :name="it.slot" :row="v">
                  </slot>
                  <span class="tooltip_span" v-if="it.tooltip">{{v[it.key || it.slot]}}</span>
                </div>
              </template>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { setTimeout } from 'timers';
export default {
  name: 'MTable',
  props: {
    columns: {
      type: Array,
      default: []
    },
    data: {
      type: Array,
      default: []
    },
    sortBy: {
      type: String,
      default: ''
    },
    sortDesc: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return{
      show: false,
      colWidth: [],
      sorted: false,
      sortAsBy: '',
      sortAsDesc: true,
      columnsChange: true
    }
  },
  watch: {
    data(newVal, oldVal) {
      this.sortData();
      this.computedColWidth();
    },
    columns(newVal, oldVal) {
      this.columnsChange = true;
      this.sorted = false;
    }
  },
  methods: {
    computedColWidth() {
      this.$nextTick(() => {
        if (!this.columnsChange) {
          return;
        }
        let e = this.$refs.table_body && this.$refs.table_body.querySelector('tr');
        if (!e) {
          this.show = true;
          return;
        }
        let arr = [];
        e.childNodes.forEach(v => {
          arr.push(v.offsetWidth);
        });
        this.colWidth = arr;
        this.show = true;
        this.columnsChange = false;
      });
    },
    sortData() {
      if (!this.sortAsBy || this.sorted) {
        return;
      }
      let col = this.columns.find(v => (v.key === this.sortAsBy || v.slot === this.sortAsBy));
      let func = col && col.sortMethod;
      if (func) {
        this.sortAsDesc ? this.data.sort(func).reverse() : this.data.sort(func);
      } else {
        this.defaultSortFunc();
      }
      this.sorted = true;
    },
    defaultSortFunc() {
      this.data.sort((a, b) => {
        let v = a[this.sortAsBy] > b[this.sortAsBy] ?  1 : -1;
        v = this.sortAsDesc ? -v : v;
        return v;
      });
    },
    sortDataByClick(v) {
      if (!v.sortable) {
        return;
      }
      this.sortAsDesc = this.sortAsBy === (v.key || v.slot) ? !this.sortAsDesc : false;
      this.sortAsBy = (v.key || v.slot);
      this.sorted = false;
      this.sortData();
    }
  },
  mounted() {
    this.computedColWidth();
  },
  created() {
    this.sortAsBy = this.sortBy;
    this.sortAsDesc = this.sortDesc;
  },
  updated() {
  }
}
</script>

<style lang="scss">
  table.m_table {
    width: 100%;
    min-width: 12.8rem;
    .hidden_thead {
      opacity: 0;
      line-height: 0px;
      border-color: transparent; 
      th {
        white-space: nowrap;
      }
    }
  }
  .m-table-header {
    position: relative;
    z-index: 2;
    height: 50px;
    table{
      font-size: 14px;
      color: rgb(0, 0, 0);
      thead {
        tr {
          th {
            box-sizing: border-box;
            font-weight: 500!important;
            padding: 7.5px;
            vertical-align: middle;
            white-space: nowrap;
          }
          th.sorting {
            cursor: pointer;
            .sort {
              display: inline-block;
              vertical-align: middle;
              &::before {
                display: block;
                content: '';
                border-width: 0px 5px 6px;
                border-color: transparent;
                border-style: solid;
                border-bottom-color: #acacb6;
              }
              &::after {
                display: block;
                content: '';
                margin-top: 2px;
                border-width: 6px 5px 0px;
                border-color: transparent;
                border-style: solid;
                border-top-color: #acacb6;
              }
            }
            .desc {
              &::after {
                border-top-color: #3598db;
              }
            }
            .asc {
              &::before {
                border-bottom-color: #3598db;
              }
            }
          }
          height: 50px;
          border-bottom: .02rem solid #3598db;
        }
      }
    }
  }
  .m-table-body {
    margin-top: -5px;
    tr {
      td {
        padding:7.5px;
        box-sizing: border-box;
        line-height: 24px;
        color: #a2a2ae;
      }
      border-bottom: 1px solid #dee2e6;
      &:nth-of-type(2n) {
        td {
          background-color: #f6f6f6;
        }
      }
      &:nth-of-type(2n + 1) {
        td {
          background-color: #ffffff;
        }
      }
    }
    .tooltip_span_container {
      position: relative;
      display: inline;
      cursor: pointer;
      &:hover .tooltip_span{
        display: block;
      }
      .tooltip_span {
        display: none;
        position: absolute;
        z-index: 1000;
        bottom: calc(100% + 4px);
        left: 50%;
        transform: translateX(-50%);
        margin-top: -10px auto 0;
        padding: 0 15px;
        color: #ffffff;
        background-color: #000000;
        line-height: 35px;
        border-radius: .04rem;
        &::after {
          width: 0;
          height: 0;
          border: .04rem solid transparent;
          content: "";
          display: block;
          position: absolute;
          border-top-color: #000000;
          left: 50%;
          margin-left: -4px;
        }
      }
    }
  }
</style>