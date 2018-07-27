<!--头部页面-->
<template>
  <div :class="appHeaderVar" v-show="showHeader">
    <header class="app_header_person_computer" v-show="devicesShow === 1">
      <div class="header_top">
        <div class="imageWrap" style="cursor:pointer;margin-top:0.1rem;" @click="featureButtonClick('/home')">
          <img src="../assets/logo.png" alt="失去网络了..."/>
          <div class="logo_title_wrap">
            <div class="logo_title_content">
              <span>IRIS</span>
              <span>Explorer</span>
            </div>
            <p>{{fuxi}}</p>
          </div>
        </div>
        <div class="navSearch">
          <input type="text" class="search_input"
                 placeholder="Search by Address / Txhash / Block"
                 v-model="searchInputValue"
                 @keyup.enter="onInputChange">
          <i class="search_icon" @click="getData(searchInputValue)"></i>
          <i class="clear_icon" @click="clearData" v-show="showClear"></i>
        </div>
      </div>

    </header>
    <div class="useFeature" v-show="devicesShow === 1">
      <div class="navButton">
        <span class="nav_item common_item_style" :class="activeClassName === '/home'?'nav_item_active':''"
              @click="featureButtonClick('/home')"
        >Home</span>
        <span class="nav_item common_item_style" :class="activeClassName === '/block'?'nav_item_active':''"
              @click="featureButtonClick('/block/1/0')"
        >Blocks</span>
        <div class="nav_item sub_btn_wrap common_item_style" :class="activeClassName === '/transaction'?'nav_item_active':''"
             @mouseover="transactionMouseOver" @mouseleave="transactionMouseLeave">

          <span class="nav_item common_item_style">
            Transactions
            <span class="bottom_arrow"></span>
          </span>
          <span class="sub_btn_item" @click="featureButtonClick('/recent_transactions/2/recent')"
                v-show="showSubTransaction">Recent Transactions</span>
          <span class="sub_btn_item" @click="featureButtonClick('/transfer_transactions/2/transfer')"
                v-show="showSubTransaction">Transfer Transactions</span>
          <span class="sub_btn_item" @click="featureButtonClick('/stake_transactions/2/stake')"
                v-show="showSubTransaction">Stake Transactions</span>
        </div>
        <div class="nav_item sub_btn_wrap common_item_style" :class="activeClassName === '/validators'?'nav_item_active':''"
             @mouseover="validatorsMouseOver" @mouseleave="validatorsMouseLeave">

          <span class="nav_item common_item_style">
            Validators
            <span class="bottom_arrow" style="right:0.22rem;"></span>
          </span>
          <span class="sub_btn_item" @click="featureButtonClick('/validators/3/0')"
                v-show="showSubValidators">Validators</span>
          <span class="sub_btn_item" @click="featureButtonClick('/candidates/4/0')"
                v-show="showSubValidators">Candidates</span>

        </div>
        <span class="nav_item common_item_style" :class="activeClassName === '/faucet'?'nav_item_active':''"
              @click="featureButtonClick('/faucet')"
        >Faucet</span>
      </div>
    </div>

    <div class="app_header_mobile" v-show="devicesShow === 0">
      <div class="feature_btn" @click="showFeature"></div>
      <div class="image_wrap_mobile" @click="featureButtonClick('/home',true)">
        <img src="../assets/logo.png" alt="失去网络了..." style="height:0.7rem;"/>
        <div class="logo_title_wrap">
          <div class="logo_title_content">
            <span>IRIS</span>
            <span>Explorer</span>
          </div>
          <p>{{fuxi}}</p>
        </div>
      </div>

      <div class="use_feature_mobile" v-show="featureShow">
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/home')">Home</span>
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/block/1/0')">Blocks</span>
        <span class="feature_btn_mobile feature_nav feature_arrow" @click="transactionShow =! transactionShow">Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/recent_transactions/2/recent')">Recent Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/transfer_transactions/2/transfer')">Transfer Transactions</span>
        <span class="feature_btn_mobile feature_subNav" v-show="transactionShow"
              @click="featureButtonClick('/stake_transactions/2/stake')">Stake Transactions</span>
        <span class="feature_btn_mobile feature_nav feature_arrow"
              @click="validatorsShow =! validatorsShow">Validators</span>
        <span class="feature_btn_mobile feature_subNav" v-show="validatorsShow"
              @click="featureButtonClick('/validators/3/0')">Validators</span>
        <span class="feature_btn_mobile feature_subNav" v-show="validatorsShow"
              @click="featureButtonClick('/candidates/4/0')">Candidates</span>
        <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/faucet')">Faucet</span>
      </div>
      <div class="search_input_mobile">
        <input type="text" class="search_input"
               v-model="searchInputValue"
               @keyup.enter="onInputChange"
               placeholder="Search by Address / Txhash / Block">
        <i class="search_icon" @click="getData(searchInputValue)"></i>
        <i class="clear_icon" @click="clearData" v-show="showClear"></i>
      </div>
    </div>
  </div>
</template>
<script>
  import Tools from '../common/Tools';
  import axios from 'axios';

  export default {
    name: 'app-header',
    watch:{
      $route(){
        this.listenRouteForChangeActiveButton();
        this.showHeader = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
      },
      searchInputValue(searchInputValue){
        console.log(searchInputValue)
        if(searchInputValue){
          this.showClear = true;
        }else{
          this.showClear = false;
        }
      }
    },
    data() {
      return {
        devicesWidth: window.innerWidth,
        devicesShow: 1,//1是显示pc端，0是移动端
        searchValue: '',
        appHeaderVar: 'person_computer_header_var',
        featureShow: false,//是否显示功能菜单栏
        transactionShow: false,//点击显示Transactions菜单
        validatorsShow: false,//点击显示validators菜单
        searchInputValue: '',
        activeClassName: '/home',
        showHeader:!(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
        fuxi:this.fuxi,
        showSubTransaction:false,
        showSubValidators:false,
        showClear:false,
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.devicesShow = 1;
        this.appHeaderVar = 'person_computer_header_var';
      } else {
        this.devicesShow = 0;
        this.appHeaderVar = 'mobile_header_var';
      }
    },
    mounted() {
      document.getElementById('router_wrap').addEventListener('click', this.hideFeature);
      this.listenRouteForChangeActiveButton();
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
      featureButtonClick(path, isLogoClick) {
        if (!isLogoClick) {
          this.featureShow = !this.featureShow;
        }
        this.showSubTransaction = false;
        this.showSubValidators = false;
        this.listenRouteForChangeActiveButton();
        this.$router.push(path);
      },
      transactionMouseOver(){
        this.showSubTransaction = true;
      },
      transactionMouseLeave(){
        this.showSubTransaction = false;
      },
      validatorsMouseOver(){
        this.showSubValidators = true;
      },
      validatorsMouseLeave(){
        this.showSubValidators = false;
      },

      getData(data) {
        let urlBlock = `/api/block/${this.searchInputValue}`;
        let urlTransaction = `/api/tx/${this.searchInputValue}`;
        let urlAddress = `/api/account/${this.searchInputValue}`;
        axios.get(urlBlock).then((data)=>{
          if (data.status === 200) {
            return data.data;
          }
        }).then((data)=>{
          if(data){
            this.$router.push(`/blocks_detail/${this.searchInputValue}`)
          }
        });
        axios.get(urlTransaction).then((data)=>{
          if (data.status === 200) {
            return data.data;
          }
        }).then((data)=>{
          if(data){
            this.$router.push(`/tx?txHash=${this.searchInputValue}`)
          }
        });
        axios.get(urlAddress).then((data)=>{
          if (data.status === 200) {
            return data.data;
          }
        }).then((data)=>{
          if(data){
            this.$router.push(`/address/1/${this.searchInputValue}`)
          }
        });
      },
      onInputChange() {
        this.getData();
      },
      listenRouteForChangeActiveButton(){
        //刷新的时候路由不变，active按钮不变
        let path = window.location.href;
        if (path.includes('/recent_transactions/2')) {
          this.activeClassName = '/transaction';
        } else if (path.includes('/validators/3') || path.includes('/candidates/4')) {
          this.activeClassName = '/validators';
        } else if (path.includes('/block')) {
          this.activeClassName = '/block';
        } else if (path.includes('/home')) {
          this.activeClassName = '/home';
        } else if (path.includes('/faucet')) {
          this.activeClassName = '/faucet';
        }
      },
      clearData(){
        this.searchInputValue = '';
      }


    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .person_computer_header_var {
    height: 1.8rem;
  }

  .mobile_header_var {
    height: 1.2rem;
  }

  .person_computer_header_var, .mobile_header_var {
    @include flex();
    @include pcContainer;
    .app_header_person_computer {
      height: 0.96rem;
      @include pcCenter;
      justify-content: space-between;
      padding-top: 0.01rem;
      .header_top {
        @include flex();
        justify-content: space-between;
        .imageWrap{
          @include flex;
          .logo_title_wrap{
            margin-left:0.16rem;
            .logo_title_content{
              span{
                &:first-child{
                  font-size:0.24rem;
                  color:#3498db;
                }
                &:last-child{
                  font-size:0.24rem;
                  color:#000000;
                  display:inline-block;
                  margin-left:0.12rem;
                }

              }
            }
            p{
              font-size:0.14rem;
              color:#aeaeb9;
            }

          }
        }
        .navSearch {
          margin-bottom: 0.01rem;
          position: relative;
          input::-webkit-input-placeholder {
            text-align: left;
            font-size: 0.14rem;
            color: #cccccc;
          }

          .search_input {
            @include borderRadius(0.04rem);
            width: 5.04rem;
            height: 0.36rem;
            line-height: 0.36rem;
            text-indent: 0.1rem;
            outline: none;
            border: 0.01rem solid #dddddd;
            margin-top: 0.3rem;
            font-size: 0.14rem;
            padding-right:0.3rem;
            &::-webkit-input-placeholder{
              font-size:0.14rem;
              color:#AEAEB9;
            }
            &:-moz-placeholder {
              font-size:0.14rem;
              color:#AEAEB9;
            }
            &::-moz-placeholder {
              font-size:0.14rem;
              color:#AEAEB9;
            }
            &:-ms-input-placeholder {
              font-size:0.14rem;
              color:#AEAEB9;
            }
          }
          .search_icon {
            position: absolute;
            top: 0.4rem;
            right: 0.1rem;
            width: 0.15rem;
            height: 0.15rem;
            background: url('../assets/search.svg') no-repeat;
            cursor: pointer;
          }
          .clear_icon{
            position: absolute;
            top: 0.42rem;
            right: 0.3rem;
            width: 0.15rem;
            height: 0.15rem;
            background: url('../assets/clear.png') no-repeat;
            cursor: pointer;
          }
        }
      }

    }
    .useFeature {
      width: 100%;
      height:0.66rem;
      @include flex;
      flex-direction: column;
      align-items: center;
      background: #3598db;
      .navButton {
        height:0.66rem;
        @include pcCenter;
        @include flex;
        .common_item_style{
          &:hover{
            color: #ffffff;
            background: #005a98;
          }
        }
        .nav_item {
          display: inline-block;
          height: 0.66rem;
          line-height: 0.66rem;
          padding: 0 0.57rem;;
          text-align: center;
          font-size: 0.18rem;
          cursor: pointer;
          color: #c9eafd;
          font-weight:300;
          .bottom_arrow{
            display:inline-block;
            height:0.12rem;
            width:0.2rem;
            background: url('../assets/caret-bottom.png') no-repeat 0 0;
            top:0.27rem;
            right:0.1rem;
          }
        }
        .sub_btn_wrap{
          @include flex;
          flex-direction:column;
          z-index:100;
          padding:0;
          .sub_btn_item{
            height:0.4rem;
            line-height:0.4rem;
            font-size:0.14rem;
            background: #005a98;
            color: #c9eafd;
            &:hover{
              color: #00f0ff;
            }
          }
        }
        .nav_item_active {
          color: #ffffff;
          background: #005a98;
        }
        .btn-group, .btn-group-vertical{
          vertical-align: baseline;
        }
      }

    }
    .app_header_mobile {
      width: 100%;
      padding: 0.1rem;
      @include flex();
      flex-direction: column;
      position: relative;
      height: 1rem;
      border-bottom: 0.01rem solid #cccccc;
      .feature_btn {
        position: absolute;
        width: 0.4rem;
        height: 0.4rem;
        top: 0;
        right: 0;
        background: url('../assets/menu.svg') no-repeat;
      }
      .image_wrap_mobile {
        @include flex;
        .logo_title_wrap{
          margin-left:0.16rem;
          .logo_title_content{
            span{
              &:first-child{
                font-size:0.24rem;
                color:#3498db;
              }
              &:last-child{
                font-size:0.24rem;
                color:#000000;
                display:inline-block;
                margin-left:0.12rem;
              }

            }
          }
          p{
            font-size:0.14rem;
            color:#aeaeb9;
          }

        }
      }
      .search_input_mobile {
        width: 100%;
        margin-top: 0.1rem;
        @include flex();
        flex-direction: column;
        align-items: center;
        position: relative;
        input::-webkit-input-placeholder {
          text-align: center;
          font-size: 0.14rem;
          color: #cccccc;
          line-height: 0.28rem;
        }
        input {
          width: 95%;
          @include borderRadius(0.5rem);
          border: 0.01rem solid #eee;
          font-size: 0.14rem;
          &:focus {
            border: 0.01rem solid #3190e8;
            outline: none;
          }
        }
        .search_icon {
          position: absolute;
          top: 0.05rem;
          right: 0.12rem;
          width: 0.15rem;
          height: 0.15rem;
          background: url('../assets/search.svg') no-repeat;
          cursor: pointer;
        }
        .clear_icon{

        }
      }
      .use_feature_mobile {
        position: absolute;
        width: 100%;
        top: 1rem;
        left: 0;
        background: #f2f2f2;
        @include flex();
        z-index: 100;
        flex-direction: column;
        .feature_btn_mobile {
          border-bottom: 0.01rem solid #cccccc;
          height: 0.39rem;
          line-height: 0.39rem;
          padding-left: 0.15rem;
          background: #3598db;
          color: #c9eafd;
          font-size:0.14rem;
        }
        .feature_arrow {
          position: relative;
          color: #c9eafd;
          font-size:0.14rem;
          background: url('../assets/arrow-bottom.svg') no-repeat 97% 1.2rem,#3598db;
        }
        .feature_subNav {
          padding-left: 0.3rem;
        }
      }
    }
  }


</style>
