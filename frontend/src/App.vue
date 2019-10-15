<template>
  <div id="app" @click.stop="closeSelectOption" @click="handleClick($event)">
    <app-header v-show="flShowHeader"></app-header>
    <div id="router_wrap" :class="$store.state.flShowIpt && !$store.state.isMobile ? 'input_style' : ''" :style="{'padding-top': $store.state.isMobile ? '' :`${headerHeightStyle}`}">
      <div class="qr_code"
           @click="hideQRCode"
           :style="`width:${vw}px;height:${vh}px`" v-show="weChatQRShow">
        <div class="qr_mask">
          <img src="./assets/wechat_qr.png" style="width:1.8rem;height:1.8rem;">
          <h2>Scan QR Code</h2>
          <p>to follow our SubScriptions</p>
        </div>
      </div>
      <div class="qr_code"
           @click="hideQRCode"
           :style="`width:${vw}px;height:${vh}px`" v-show="qqQRShow">
        <div class="qr_mask">
          <img src="./assets/qq_qr.png" style="width:1.8rem;height:1.8rem;">
          <h2>Join QQ group</h2>
        </div>
      </div>
      <router-view class="router_view" :style="`min-height:${vh/100-2.5}rem;`" :key="key"/>
    </div>
    <footer :class="footerClass" v-show="flShowFooter" id="footer">
      <div class="person_computer_footer">
        <div class="footer_logo_content">
          <a class="irisnet_link_content" href="https://www.irisnet.org/" target="_blank">
            <img class="irisnet_logo_img" src="./assets/IRISnet-Rebrand-Capital-Black.png"/>
          </a>
        </div>
        <div class="community_container">
          <h4 class="community_title">
            Community
          </h4>
          <div class="community_list_content">
            <a target="_blank" href='https://github.com/irisnet'><i class="iconfont icongithub"></i></a>
            <a target="_blank" href='https://t.me/irisnetwork'><i class="iconfont icontelegram"></i></a>
            <a target="_blank" href='https://twitter.com/irisnetwork'><i class="iconfont icontuite"></i></a>
            <a target="_blank" href='https://medium.com/irisnet-blog'><i class="iconfont iconmedium"></i></a>
            <span class="we_chat" @click="showQRCode"><i class="iconfont iconweixin"></i></span>
            <span class="qq" @click="showqqQRCode"><i class="iconfont iconQQ"></i></span>
          </div>
        </div>
        <div class="footer_right_content">
          <h3 class="resources_content">Resources</h3>
          <div class="footer_link_wrap">
            <a href="https://www.irisnetwork.cn/testnets" target="_blank">
              <span class="footer_link_contact">Use Testnet</span>
            </a>
            <span class="footer_link_join">|</span>
            <span class="footer_link_privacy"><router-link :to="`/privacy_policy`">Privacy Policy</router-link></span>
            <span class="footer_link_join">|</span>
            <span class="footer-faq"><router-link :to="`/help`">FAQ</router-link></span>
          </div>
        </div>
      </div>
      <p class="footer_copyright_wrap">
        ©️ IRISplorer 2019 all rights reserved
      </p>
    </footer>
  </div>
</template>
<script>
  import AppHeader from './components/AppHeader';
  import Tools from './util/Tools';
  import testVersion from '../testVersion';
  import BackToTop from "./components/BackToTop";
  import $ from "jquery"
  export default {
    components: {
      AppHeader,
      BackToTop
    },
    watch: {
      $route() {
        this.showHeaderAndFooterByVersionPath();
        Tools.scrollToTop()
      }
    },
    computed: {
      key() {
        return this.$route.name !== undefined ? this.$route.name + new Date() : this.$route + new Date()
      }
    },
    data() {
      return {
        show: 1,
        devicesWidth: window.innerWidth,
        devicesShow: 1,//1是显示pc端，0是移动端
        footerClass: 'person_computer_wrap_footer',
        footerClassName: 'person_computer_footer',
        footerLeftVar: 'person_computer_footer_left',
        footerRightVar: 'person_computer_footer_right',
        flShowFooter: !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
        vh: window.innerHeight,
        vw: window.innerWidth,
        weChatQRShow: false,
        qqQRShow: false,
        build: testVersion.app.buildNumber,
        env: testVersion.app.env,
        version: testVersion.app.version,
        innerWidth: window.innerWidth,
        scrollHeight:0,
        flShowHeader : true,
        headerHeightStyle:'',
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        document.body.style.minWidth = '4rem';
      }

    },
    created() {
      this.showHeaderAndFooterByVersionPath();
      window.addEventListener('resize', this.onresize);
      if (window.innerWidth > 910) {
        this.$store.commit('isMobile',false);
        this.footerClass = 'person_computer_wrap';
        this.footerClassName = 'person_computer_footer';
        this.footerLeftVar = 'person_computer_footer_left';
        this.footerRightVar = 'person_computer_footer_right';
      } else {
        this.$store.commit('isMobile',true);
        this.footerClass = 'mobile_wrap_footer';
        this.footerClassName = 'mobile_footer';
        this.footerLeftVar = 'mobile_footer_left';
        this.footerRightVar = 'mobile_footer_right';
      }
      window.addEventListener("scroll",this.handleScroll,true);
    },
    mounted() {
      this.headerHeightStyle = `${document.getElementById('header').clientHeight}px`;
    },
    updated() {
        setTimeout(() => {
            this.headerHeightStyle = `${document.getElementById('header').clientHeight}px`;
        });
    },
    beforeDestroy() {
      window.removeEventListener('resize', this.onWindowResize);
    },
    methods: {
      handleScroll(e){
        if(e.target.scrollTop > 0){
          this.scrollHeight = e.target.scrollTop
        }
      },
      handleClick(e){
        let $i = $("<i/>").addClass('iconfont iconiris');
        let x = e.pageX, y = e.pageY;
        $i.css({ "z-index": 9999999, "top": y - 20, "left": x,"position": "absolute","color": "var(--bgColor)","font-size": "20px"});
        $("body").append($i);
        $i.animate({"top": y - 180,"opacity": 0},1500,function () {$i.remove();});
      },
      onresize() {
        this.innerWidth = window.innerWidth;
        this.vh = window.innerHeight;
        this.vw = window.innerWidth;
        if (window.innerWidth > 910) {
          this.$store.commit('isMobile',false);
          this.footerClass = 'person_computer_wrap';
          this.footerClassName = 'person_computer_footer';
          this.footerLeftVar = 'person_computer_footer_left';
          this.footerRightVar = 'person_computer_footer_right';
        } else {
          this.$store.commit('isMobile',true);
          this.footerClass = 'mobile_wrap_footer';
          this.footerClassName = 'mobile_footer';
          this.footerLeftVar = 'mobile_footer_left';
          this.footerRightVar = 'mobile_footer_right';
        }
      },
      showHeaderAndFooterByVersionPath(){
        if(this.$route.path === "/version" || this.$route.matched[0].path && this.$route.matched[0].path === "*"){
          this.flShowFooter = false;
          this.flShowHeader = false;
        }else {
          this.flShowHeader = !(this.$route.query.flShow && this.$route.query.flShow === 'false');
          this.flShowFooter = !(this.$route.query.flShow && this.$route.query.flShow === 'false');
        }
      },
      showQRCode() {
        this.weChatQRShow = true;
      },
      showqqQRCode(){
        this.qqQRShow =  true;
      },
      hideQRCode() {
        this.weChatQRShow = false;
        this.qqQRShow =  false;
      },
      closeSelectOption(){
        this.$store.commit('flShowSelectOption',false)
      }
    }
  }
</script>
<style lang="scss">
  @import './style/mixin.scss';

  html {
    font-size: 625%;
    -webkit-text-size-adjust: none;
    overflow-y: auto;
    position: relative;
  }

  body, html {
    width: 100%;
    height: 100%;
    color: var(--moduleColor);
  }
  body {
    font-size: 16px !important;
    font-family:Arial !important;
    position: relative;
  }

  p {
    margin-bottom: 0 !important;
  }

  ul {
    margin-bottom: 0 !important;
  }

  #app {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    #router_wrap {
      flex: 1;
      background: #F5F7FD;
      .router_view {
        min-height: 4.5rem;
        background: #F5F7FD;
      }
      .person_computer_wrap_footer {
        height: 1rem;
      }
      .mobile_wrap_footer {
        /*height:21rem;*/
      }
      .qr_code { //微信二维码弹框
        position: fixed;
        top: 0;
        left: 0;
        background: rgba(0, 0, 0, .6);
        z-index: 10002;
        @include flex;
        justify-content: center;
        align-items: center;
        .qr_mask {
          width: 3.7rem;
          height: 3.7rem;
          background: #141426;
          @include borderRadius(0.05rem);
          @include flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
          h2 {
            color: #5c5c7c;
            font-size: 0.24rem;
            margin-top: 0.1rem;
            margin-bottom: 0;
          }
          p {
            color: #5c5c7c;
            margin-top: 0.1rem;
            font-size: 0.18rem;
          }
        }
      }
    }
    #footer{
      /*background: var(--bgColor);*/
      background: #363a3d;
      flex-direction:column;
      display: flex;
      height: 1.9rem;
      .person_computer_footer { //分别对pc和移动端兼容
        max-width: 12.8rem;
        width: 100%;
        height: 1.35rem;
        margin: 0 auto;
        display: flex;
        .footer_logo_content{
          flex: 1;
          height: 100%;
          .irisnet_link_content{
            display: flex;
            align-items: center;
            padding-left: 0.2rem;
            height: 100%;
            .irisnet_logo_img{
              height: 43%;
            }
          }
        }
        .community_container{
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: center;
          .community_title{
            font-size: 0.16rem;
            color: #fff;
            padding-bottom: 0.1rem;
            margin: 0;
          }
          .community_list_content{
            display: flex;
            a{
              margin-right: 0.25rem;
              i{
                font-size: 0.25rem;
                color: rgba(255,255,255,0.5);
                &:hover{
                  color: #fff;
                }
              }
            }
            span{
              margin-right: 0.25rem;
              cursor: pointer;
              i{
                font-size: 0.25rem;
                color: rgba(255,255,255,0.5);
                &:hover{
                  color: #fff;
                }
              }
            }
          }
        }
        .footer_right_content{
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: center;
          .resources_content{
            font-size: 0.16rem;
            color: #fff;
            margin-bottom: 0.1rem;
          }
          .footer_link_wrap{
            color: rgba(255,255,255,0.5);
            a{
              padding: 0 0.2rem;
              color: rgba(255,255,255,0.5) !important;
              &:hover{
                color: #fff !important;
              }
            }
            a:first-child{
              padding-left: 0;
            }
            .footer_link_privacy{
              padding-left: 0.2rem;
            }
            .footer-faq{
              padding-left: 0.2rem;
            }

          }
        }
      }
      .footer_copyright_wrap {
        /*background: var(--bgColor);*/
        background: #363a3d;
        border-top: 0.01rem  solid rgba(255,255,255,0.2);
        padding: 0.15rem 0;
        text-align: center;
        color: rgba(255,255,255,0.5);
      }
    }
    .input_style{
      padding-top: 1.21rem !important;
    }
  }
  pre{
    font-family: Arial !important;
    font-size: 0.14rem !important;
  }
  @media screen and (max-width: 910px) {
    #app{
      #footer {
        height: auto;
        .person_computer_footer{
          flex-direction: column;
          height: auto;
          .footer_logo_content{
            width: 100%;
            display: flex;
            justify-content: center;
            flex: 1;
            height: 100%;
            .irisnet_link_content{
              width: 40%;
              padding: 0.1rem 0;
              display: flex;
              align-items: center;
              .irisnet_logo_img{
                width: 100%;
              }
            }
          }
          .community_container{
            justify-content: center;
            .community_title{
              text-align: center;
            }
            .community_list_content{
              justify-content: center;
            }
          }
          .footer_right_content{
            padding: 0.2rem 0;
            .resources_content{
              text-align: center;
            }
            .footer_link_wrap{
              text-align: center;
            }
          }
        }
        .footer_copyright_wrap{}
      }
    }
    //解决在移动端，日期选择器会移出到窗口以外
    .ivu-date-picker:nth-of-type(2){
      .ivu-select-dropdown{
        left: 150px !important;
      }
    }
  }

</style>
