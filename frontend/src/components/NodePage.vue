<template>
  <div class="blocks_list_page_wrap">
    <div class="blocks_list_title_wrap">
      <p :class="blocksListPageWrap" style="margin-bottom:0;">
        <span class="blocks_list_title">{{titleVar}}</span><span class="blocks_list_page_wrap_hash_var" >{{count}} total</span>
        <!--<span class="blocks_list_page_wrap_hash_var">{{blocksValue}}</span>-->

        <span class="blocks_list_page_wrap_hash_var for_block"
              v-show="$route.query.block || $route.query.address">
          {{blockVar}}
        </span>
      </p>
    </div>

    <div :class="blocksListPageWrap">
      <div class="pagination total_num">
      </div>
      <div style="position:relative;overflow-x: auto;">
        <spin-component :showLoading="showLoading"/>
          <ul id="node-ul-container">
            <li class="list-items">
              <span>Moniker / NodeID</span>
              <span>IP / Region</span>
              <span>Start Date</span>
              <span>Neteork / Channel</span>
            </li>
            <li class="list-items-info" v-for="(item,index) in nodeList" :class="index % 2 == 0 ? 'background_white ' : ' background_gray'">
              <div>
                <p>{{item.node_info.moniker}}</p>
                <p>{{item.node_info.id}}</p>
              </div>
              <div>
                <p>{{item.node_info.listen_addr}}</p>
                <p>{{item.node_info.country}}</p>
              </div>
              <div>
                <p>{{item.connection_status.SendMonitor.Start}}</p>
              </div>
              <div>
                <p>{{item.node_info.network}}</p>
                <p>{{item.node_info.channels}}</p>
              </div>
            </li>
          </ul>
        <div v-show="showNoData" class="no_data_show">
          No Data
        </div>
      </div>
    </div>

  </div>
</template>

<script>
  import Tools from '../common/Tools';
  import axios from 'axios';
  import BlocksListTable from './table/BlocksListTable.vue';
  import SpinComponent from './commonComponents/SpinComponent';
  export default {
    components:{
      BlocksListTable,
      SpinComponent,
    },
    watch: {
      $route() {
        this.items = [];
        this.type = this.$route.params.type;
        this.currentPage = 1;
        this.getDataList(1, 30, this.$route.params.type);
        this.showNoData = false;
        switch (this.$route.params.type) {
          case '1': this.titleVar = 'Full Nodes';
            break;
        }


        this.computeMinWidth();
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        blocksListPageWrap: 'personal_computer_blocks_list_page',//1是显示pc端，0是移动端
        blocksValue: '',
        currentPage: 1,
        pageSize: 30,
        count: 0,
        fields: [],
        items: [],
        type: '1',
        titleVar: 'Full Nodes',
        showNoData:false,//是否显示列表的无数据
        showLoading:false,
        blockVar:'',
        innerWidth : window.innerWidth,
        tableMinWidth:'',
        nodeList:[]
      }
    },
    beforeMount() {
      this.type = this.$route.params.type;
      if (window.innerWidth > 910) {
        this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
      } else {
        this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
      }
    },
    mounted() {
      //组件页面根据路由类型判断是哪个页面
      this.getDataList();
      switch (this.$route.params.type) {
        case '1': this.titleVar = 'Full Nodes';
          break;
      }
      window.addEventListener('resize',this.onresize);
      this.computeMinWidth();
    },
    beforeDestroy() {
      window.removeEventListener('resize',this.onWindowResize);
    },
    methods: {
      onresize(){
        this.innerWidth = window.innerWidth;
        if(window.innerWidth > 910){
          this.blocksListPageWrap = 'personal_computer_blocks_list_page_wrap';
        } else {
          this.blocksListPageWrap = 'mobile_blocks_list_page_wrap';
        }
      },
      //根绝页面的不同展示最小宽度,不换行显示列表
      computeMinWidth(){
        this.tableMinWidth = 18;
      },
      getDataList() {
        this.showLoading = true;
        let url = `http://116.62.62.39:26657/net_info`;
        let searchIpUrl = `/api/ip/`;
        axios.get(url).then((data) => {
          if (data.status === 200) {
            return data.data;
          }
        }).then((data) => {
          this.count = data.result.peers.length;
          this.nodeList = data.result.peers;
          this.nodeList.forEach(item =>{
            item.connection_status.SendMonitor.Start = Tools.conversionTimeToUTC(item.connection_status.SendMonitor.Start);
            item.node_info.listen_addr = item.node_info.listen_addr.split(":")[0];
            axios.get(searchIpUrl + item.node_info.listen_addr).then((data) => {
              if (data.status === 200) {
                return data.data;
              }
            }).then((data) => {
              if(data && data.data){
                if(data.data.country_id === "xx"){
                  data.data.country_id = "local"
                }
                let tempNodeList = JSON.parse(JSON.stringify(this.nodeList));
                for(let i=0; i < tempNodeList.length;i++){
                  if(tempNodeList[i].node_info.listen_addr.split(":")[0] === data.data.ip){
                    tempNodeList[i].node_info.country = data.data.country_id;
                    break;
                  }
                }
                this.nodeList = tempNodeList;
              }
            })
          });
          this.showLoading = false;
          return this.nodeList
        })

      },
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';
  .blocks_list_page_wrap {
    @include flex;
    @include pcContainer;
    font-size: 0.14rem;
    .pagination {
      @include flex;
      justify-content: flex-end;
      @include borderRadius(0.025rem);
      height:0.3rem;
      li{
        height:0.3rem !important;
      }
    }
    .total_num{
      @include flex;
      justify-content: space-between;
      height:0.45rem;
      align-items: center;
    }
    .no_data_show{
      @include flex;
      justify-content: center;
      border-top:0.01rem solid #eee;
      border-bottom:0.01rem solid #eee;
      font-size:0.14rem;
      height:3rem;
      align-items: center;
    }
    .b-table {
      //min-width: 8rem;

      a {
        text-decoration: none;
      }
    }
    .background_white{
      background: #fff;
    }
    .background_gray{
      background: #f6f6f6 !important;
    }
    .blocks_list_title_wrap {
      width: 100%;
      border-bottom: 1px solid #d6d9e0 !important;
      @include flex;
      @include pcContainer;
      height:0.62rem;
      background:#efeff1;
      p{
        height:0.62rem;
        span{
          height:0.62rem;
          line-height:0.62rem;
        }
      }
      .personal_computer_blocks_list_page_wrap {
        @include flex;

      }
      .mobile_blocks_list_page_wrap {
        @include flex;
        flex-direction: column;
        overflow-x: auto;
        width:100%;
        .blocks_list_page_wrap_hash_var{
          min-width:7rem;
          font-size: 0.22rem;
          color: #a2a2ae;
        }
      }

    }
    .personal_computer_blocks_list_page_wrap {
      .transaction_information_content_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #000000;
        margin-bottom: 0;
      }
      @include pcCenter;
      min-height:4.6rem;
      .transactions_detail_information_wrap {
        .information_props_wrap {
          @include flex;
          .information_props {
            width: 15rem;
          }
        }
      }

      .blocks_list_title {
        height: 0.62rem;
        line-height: 0.62rem;
        font-size: 0.22rem;
        color: #000000;
        margin-right: 0.2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        height:  0.62rem;
        line-height: 0.62rem;
        font-size: 0.22rem;
        color: #a2a2ae;
      }
      .for_block{
        display:inline-block;
        margin-left:0.1rem;
      }
    }

    .mobile_blocks_list_page_wrap {
      width: 100%;
      @include flex;
      flex-direction: column;
      padding: 0 0.1rem;
      .transaction_information_content_title {
        height: 0.4rem;
        line-height: 0.4rem;
        font-size: 0.18rem;
        color: #000000;
        margin-bottom: 0;
      }
      .transactions_detail_information_wrap {

        .information_props_wrap {
          @include flex;
          flex-direction: column;
          border-bottom: 0.01rem solid #eee;
          margin-bottom: 0.05rem;
          .information_value {
            overflow-x: auto;
          }

        }
      }

      .blocks_list_title {
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #000000;
        margin-right: 0.2rem;
        font-weight: 500;
      }
      .blocks_list_page_wrap_hash_var {
        overflow-x: auto;
        height: 0.3rem;
        line-height: 0.3rem;
        font-size: 0.18rem;
        color: #a2a2ae;
      }
      .for_block{
        display:inline-block;
        //margin-left:0.1rem;
      }

    }
  }

  //重置bootstrap-vue的表格样式
  table{

    td{
      max-width:2.2rem !important;
      overflow-wrap: break-word !important;
    }
  }
  .page-item{
    &:first-child, &:last-child{
      .page-link{
        @include borderRadius(0.025rem);
      }
    }
  }
  #node-ul-container{
    min-width: 10rem;
    max-width: 12.8rem;
    min-height: 4.6rem;
  }
  .list-items{
    box-sizing: border-box;
    padding: 0 0.2rem;
    @include flex;
    font-weight: 600;
    font-size: 0.14rem;
    height: 0.5rem;
    line-height: 0.5rem;
    border-bottom: 0.02rem solid #3598db;
    border-top: 0.01rem solid #d6d9e0;
    span:nth-child(1){
      display: inline-block;
      flex: 4;
    };
    span:nth-child(2){
      display: inline-block;
      flex: 2;
    };
    span:nth-child(3){
      display: inline-block;
      flex: 2;
    };
    span:nth-child(4){
      display: inline-block;
      flex: 2;
    }
  }
  .list-items-info{
    box-sizing: border-box;
    padding: 0 0.2rem;
    padding-top: 0.05rem;
    @include flex;
    height: 0.51rem;
    line-height: 0.2rem;
    div:nth-child(1){
      flex: 4;
      p:nth-child(1){
        color: #000;
        font-weight: 400;
      }
      p:nth-child(2){
        color: #a2a2ae;
      }
    };
    div:nth-child(2){
      flex: 2;
      p:nth-child(1){
        color: #000;
        font-weight: 400;
      }
      p:nth-child(2){
        color: #a2a2ae;
      }
    };
    div:nth-child(3){
      flex: 2;
      padding-top: 0;
      p{
        line-height: 0.40rem;
        display: inline-block;
        color: #000;
        font-weight: 400;
      }
    };
    div:nth-child(4){
      flex: 2;
      p:nth-child(1){
        color: #000;
        font-weight: 400;
      }
      p:nth-child(2){
        color: #a2a2ae;
      }
    };
  }
</style>
