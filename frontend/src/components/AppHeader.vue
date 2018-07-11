<!--头部页面-->
<template>
  <div class="app_header">
    <header class="app_header_person_computer" v-show="devicesShow === 1">
      <div class="header_top">
        <div class="imageWrap">
          <img src="../assets/logo.png" alt="失去网络了..."/>
        </div>
        <div class="navSearch">
          <input type="text" class="search_input"
                 placeholder="Search by Address / Txhash / Block"
                 v-model="searchInputValue"
                 @keyup.enter="onInputChange">
          <i class="search_icon" @click="getData(searchInputValue)"></i>
        </div>
      </div>

    </header>
    <div class="useFeature" v-show="devicesShow === 1">
      <div class="navButton">
        <span class="nav_item" :class="activeClassName === '/home'?'nav_item_active':''"
              @click="featureButtonClick('/home')"
        >Home</span>
        <span class="nav_item" :class="activeClassName === '/block'?'nav_item_active':''"
              @click="featureButtonClick('/block')"
        >Block</span>
        <b-dropdown id="ddown-left" text="Transaction" variant="primary" class="m-2" :class="activeClassName === '/transaction'?'nav_item_active':''">
          <b-dropdown-item @click="featureButtonClick('/recent_transactions')">Recent Transactions</b-dropdown-item>
          <b-dropdown-item @click="featureButtonClick('/transfer_transactions')">Transfer Transactions</b-dropdown-item>
          <b-dropdown-item @click="featureButtonClick('/delegate_transactions')">Delegate Transactions</b-dropdown-item>
        </b-dropdown>
        <b-dropdown text="Left align" variant="Validators" class="m-2" :class="activeClassName === '/validators'?'nav_item_active':''">
          <b-dropdown-item @click="featureButtonClick('/validators')">Validators</b-dropdown-item>
          <b-dropdown-item @click="featureButtonClick('/candidates')">Candidates</b-dropdown-item>
        </b-dropdown>
        <span class="nav_item" :class="activeClassName === '/faucet'?'nav_item_active':''"
              @click="featureButtonClick('/faucet')"
        >Faucet</span>
      </div>
    </div>

    <div class="app_header_mobile" v-show="devicesShow === 0">
      <div class="feature_btn" @click="showFeature"></div>
      <div class="image_wrap_mobile">
        <!--<img src="../assets/logo.png" alt="失去网络了..."/>-->
        logo
      </div>

      <div class="use_feature_mobile" v-show="featureShow">
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/home')">Home</span>
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/blocks')">Blocks</span>
        <span class="feature_btn_mobile feature_nav" @click="transactionShow =! transactionShow">Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/recent_transactions')">Recent Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/transfer_transactions')">Transfer Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/delegate_transactions')">Delegate Transactions</span>
        <span class="feature_btn_mobile feature_nav" @click="validatorsShow =! validatorsShow">Validators</span>
        <span class="feature_btn_mobile feature_subNav" v-show="validatorsShow"
              @click="featureButtonClick('/validators')">Validators</span>
        <span class="feature_btn_mobile feature_subNav" v-show="validatorsShow"
              @click="featureButtonClick('/candidates')">Candidates</span>
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/faucet')">Faucet</span>
      </div>
      <div class="search_input_mobile">
        <input type="text" class="search_input"
               placeholder="Search by Address / Txhash / Block"
        >
      </div>
    </div>
  </div>
</template>
<script>

  export default {
    name: 'app-header',
    data() {
      return {
        devicesWidth: window.innerWidth,
        devicesShow: 1,//1是显示pc端，0是移动端
        searchValue: '',
        featureShow: false,//是否显示功能菜单栏
        transactionShow: false,//点击显示Transactions菜单
        validatorsShow: false,//点击显示validators菜单
        searchInputValue: '',
        activeClassName:'/home',
      }
    },
    beforeMount() {
      if (this.devicesWidth > 500) {
        this.devicesShow = 1;
      } else {
        this.devicesShow = 0;
      }
    },
    mounted() {
      document.getElementById('router_wrap').addEventListener('click', this.hideFeature)
    },
    beforeDestroy() {
      document.getElementById('router_wrap').removeEventListener('click', this.hideFeature)
    },
    methods: {
      hideFeature() {
        if (this.featureShow) {
          this.featureShow = false;
        }
      },
      showFeature() {
        this.featureShow = !this.featureShow;
      },
      featureButtonClick(path) {
        this.featureShow = !this.featureShow;
        if(path === '/recent_transactions' || path === '/transfer_transactions' || path === '/delegate_transactions'){
          this.activeClassName = '/transaction';
        }else if(path === '/candidates'){
          this.activeClassName = '/validators';
        }else{
          this.activeClassName = path;
        }
        this.$router.push(path);

        console.log(path)

      },
      getData(data) {
        console.log(data)
      },
      onInputChange() {
        console.log(this.searchInputValue)
      }


    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .app_header {
    height: 116px;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    background: yellow;
    .app_header_person_computer {
      width: 80%;
      min-width: 600px;
      max-width: 1000px;
      border: 1px solid red;
      justify-content: space-between;
      padding-top: 10px;
      .header_top{
        @include flex();
        justify-content: space-between;
        .imageWrap {
          width: 100px;
          height: 50px;
          img {
            width: 100px;
            height: 50px;
          }
        }
        .navSearch {
          margin-bottom: 10px;
          position: relative;
          input::-webkit-input-placeholder {
            text-align: left;
            font-size: 13px;
            color: #cccccc;
          }

          .search_input {
            @include borderRadius(4px);
            width: 300px;
            height: 28px;
            line-height: 28px;
            text-indent: 10px;
            outline: none;
            border: 1px solid #dddddd;
          }
          .search_icon {
            position: absolute;
            top: 5px;
            right: 10px;
            width: 15px;
            height: 15px;
            background: lawngreen;
          }
        }
      }



    }
    .useFeature {
      width:100%;
      border-top:1px solid #dddddd;
      border-bottom:1px solid #ccc;
      @include flex();
      flex-direction: column;
      align-items: center;
      .navButton {
        width:80%;
        min-width: 600px;
        max-width: 1000px;
        .nav_item {
          display: inline-block;
          height: 51px;
          line-height: 50px;
          padding: 0 40px;
          text-align: center;
          font-size: 14px;
          cursor: pointer;
        }
        .nav_item_active{
          border-bottom:2px solid #3190e8;
          color:#3190e8;
        }
        .m-2 {
          margin: 0 !important;
          height:50px;
          button {
            padding: 0 40px;
            color:#555 !important;
          }
          .btn {
            background-color: transparent;
            color: #dddddd;
            border: none;
            font-weight: normal;
          }
        }
      }

    }
    .app_header_mobile {
      width: 100%;
      padding: 10px 0;
      @include flex();
      flex-direction: column;
      position: relative;
      height: 100px;
      .feature_btn {
        position: absolute;
        width: 40px;
        height: 40px;
        top: 0;
        right: 0;
        background: lawngreen;
      }
      .image_wrap_mobile {
        width: 70vw;
        height: 40px;
        img {
          width: 100%;
          height: 100%;
        }
      }
      .search_input_mobile {
        width: 100%;
        margin-top: 10px;
        @include flex();
        flex-direction: column;
        align-items: center;
        input::-webkit-input-placeholder {
          text-align: center;
          font-size: 14px;
          color: #cccccc;
          line-height: 28px;
        }
        input {
          width: 95%;
          @include borderRadius(5px);
          border: 1px solid #eee;
          font-size: 14px;
          &:focus {
            border: 1px solid #3190e8;
            outline: none;
          }
        }
      }
      .use_feature_mobile {
        position: absolute;
        width: 100%;
        top: 100px;
        left: 0;
        background: #3190e8;
        @include flex();
        flex-direction: column;
        .feature_btn_mobile {
          border-bottom: 1px solid #eee;
          height: 39px;
          line-height: 39px;
          padding-left: 15px;
          color: #555;
        }
        .feature_subNav {
          padding-left: 30px;
        }
      }
    }
  }


</style>
