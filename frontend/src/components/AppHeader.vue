<!--头部页面-->
<template>
  <div :class="appHeaderVar">
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
              @click="featureButtonClick('/transactions_detail')"
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
        <span class="feature_btn_mobile feature_nav feature_arrow" @click="transactionShow =! transactionShow">Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/recent_transactions')">Recent Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/transfer_transactions')">Transfer Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/delegate_transactions')">Delegate Transactions</span>
        <span class="feature_btn_mobile feature_nav feature_arrow" @click="validatorsShow =! validatorsShow">Validators</span>
        <span class="feature_btn_mobile feature_subNav" v-show="validatorsShow"
              @click="featureButtonClick('/validators')">Validators</span>
        <span class="feature_btn_mobile feature_subNav" v-show="validatorsShow"
              @click="featureButtonClick('/candidates')">Candidates</span>
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/faucet')">Faucet</span>
      </div>
      <div class="search_input_mobile">
        <input type="text" class="search_input"
               v-model="searchInputValue"
               @keyup.enter="onInputChange"
               placeholder="Search by Address / Txhash / Block">
        <i class="search_icon" @click="getData(searchInputValue)"></i>
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
        appHeaderVar:'person_computer_header_var',
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
        this.appHeaderVar = 'person_computer_header_var';
      } else {
        this.devicesShow = 0;
        this.appHeaderVar = 'mobile_header_var';
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

  .person_computer_header_var{
    height: 15rem;
  }
  .mobile_header_var{
    height:10rem;
  }
  .person_computer_header_var, .mobile_header_var {
    width: 100%;
    @include flex();
    flex-direction: column;
    align-items: center;
    .app_header_person_computer {
      width: 80%;
      min-width: 70rem;
      max-width: 100rem;
      justify-content: space-between;
      padding-top: 1rem;
      .header_top{
        @include flex();
        justify-content: space-between;
        .imageWrap {
          width: 10rem;
          height: 7.8rem;
          img {
            width: 10rem;
            height: 5rem;
          }
        }
        .navSearch {
          margin-bottom: 1rem;
          position: relative;
          input::-webkit-input-placeholder {
            text-align: left;
            font-size: 1.3rem;
            color: #cccccc;
          }

          .search_input {
            @include borderRadius(0.4rem);
            width: 30rem;
            height: 2.8rem;
            line-height: 2.8rem;
            text-indent: 1rem;
            outline: none;
            border: 0.1rem solid #dddddd;
            margin-top:2rem;
          }
          .search_icon {
            position: absolute;
            top: 2.7rem;
            right: 1rem;
            width: 1.5rem;
            height: 1.5rem;
            background: url('../assets/search.svg') no-repeat;
            cursor:pointer;
          }
        }
      }



    }
    .useFeature {
      width:100%;
      border-top:0.1rem solid #dddddd;
      border-bottom:0.1rem solid #ccc;
      @include flex();
      flex-direction: column;
      align-items: center;
      background: #EBEBEB;
      .navButton {
        width:80%;
        min-width: 70rem;
        max-width: 100rem;
        .nav_item {
          display: inline-block;
          height: 5.7rem;
          line-height: 5.4rem;
          padding: 0 4rem;
          text-align: center;
          font-size: 1.4rem;
          cursor: pointer;
          color:#cccccc;
        }
        .nav_item_active{
          border-bottom:0.2rem solid #3190e8;
          color:#3190e8;
        }
        .m-2 {
          margin: 0 !important;
          height:5.9rem;
          button {
            padding: 0 4rem;
            color:#555 !important;
          }
          .btn {
            background-color: transparent;
            color: #cccccc !important;
            border: none;
            font-weight: normal;
          }
        }
      }

    }
    .app_header_mobile {
      width: 100%;
      padding: 1rem 0;
      @include flex();
      flex-direction: column;
      position: relative;
      height: 10rem;
      border-bottom:0.1rem solid #cccccc;
      .feature_btn {
        position: absolute;
        width: 4rem;
        height: 4rem;
        top: 0;
        right: 0;
        background: url('../assets/menu.svg') no-repeat;
      }
      .image_wrap_mobile {
        width: 70vw;
        height: 4rem;
        img {
          width: 100%;
          height: 100%;
        }
      }
      .search_input_mobile {
        width: 100%;
        margin-top: 1rem;
        @include flex();
        flex-direction: column;
        align-items: center;
        position:relative;
        input::-webkit-input-placeholder {
          text-align: center;
          font-size: 1.4rem;
          color: #cccccc;
          line-height: 2.8rem;
        }
        input {
          width: 95%;
          @include borderRadius(0.5rem);
          border: 0.1rem solid #eee;
          font-size: 1.4rem;
          &:focus {
            border: 0.1rem solid #3190e8;
            outline: none;
          }
        }
        .search_icon{
          position: absolute;
          top: 0.5rem;
          right: 1.2rem;
          width: 1.5rem;
          height: 1.5rem;
          background: url('../assets/search.svg') no-repeat;
          cursor:pointer;
        }
      }
      .use_feature_mobile {
        position: absolute;
        width: 100%;
        top: 10rem;
        left: 0;
        background: #f2f2f2;
        @include flex();
        flex-direction: column;
        .feature_btn_mobile {
          border-bottom: 0.1rem solid #cccccc;
          height: 3.9rem;
          line-height: 3.9rem;
          padding-left: 1.5rem;
          color: #555;
        }
        .feature_arrow{
          position:relative;
          background: url('../assets/arrow-bottom.svg') no-repeat 97% 1.2rem;
        }
        .feature_subNav {
          padding-left: 3rem;
        }
      }
    }
  }


</style>
