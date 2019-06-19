<template>
  <div id="app" @click.stop="closeSelectOption">
    <app-header v-show="flShowHeader"></app-header>
    <div id="router_wrap" :style="{'padding-top': $store.state.isMobile ? '' :`${headerHeightStyle}`}">
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
      <router-view class="router_view" :style="`min-height:${vh/100-2.72}rem;`" :key="key"/>
      <footer :class="footerClass" v-show="flShowFooter" id="footer">
        <div :class="footerClassName" style="height:100%;">
          <div class="footer_left" :class="footerLeftVar">
            <a target="_blank" href='https://github.com/irisnet' class="github"></a>
            <a target="_blank" href='https://t.me/irisnetwork' class="airplane"></a>
            <a target="_blank" href='https://twitter.com/irisnetwork' class="twitter"></a>
            <a target="_blank" href='https://medium.com/irisnet-blog' class="facebook"></a>
            <span class="we_chat" @click="showQRCode"></span>
            <span class="qq" @click="showqqQRCode"></span>
          </div>
          <div class="footer_center">
            <a href="https://www.irisnet.org/" target="_blank">
              <img src="./assets/irisnet.png">
            </a>
          </div>
          <div class="footer_right" :class="footerRightVar">
            <div class="footer_link_wrap">
              <a href="https://www.irisnetwork.cn/testnets" target="_blank">
                <span class="footer_link_contact">Use Testnet</span>
              </a>
              <span class="footer_link_join">|</span>
              <span class="footer_link_privacy"><router-link :to="`/privacy_policy`">Privacy Policy</router-link></span>
              <span class="footer_link_join">|</span>
              <span class="footer-faq"><router-link :to="`/help`">FAQ</router-link></span>
            </div>
            <p class="footer_copyright_wrap">
              ©️ IRISplorer 2019 all rights reserved
            </p>
          </div>
        </div>
      </footer>
      <backToTop :scrollHeight="scrollHeight"></backToTop>
    </div>

  </div>
</template>
<script>
  import AppHeader from './components/AppHeader';
  import Tools from './util/Tools';
  import testVersion from '../testVersion';
  import BackToTop from "./components/BackToTop";
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
        headerHeightStyle:''
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        document.body.style.minWidth = '4rem';
      }

    },
    mounted() {
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
    beforeDestroy() {
      window.removeEventListener('resize', this.onWindowResize);
    },
    methods: {
      handleScroll(e){
        if(e.target.scrollTop > 0){
          this.scrollHeight = e.target.scrollTop
        }
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
        if(this.$route.path === "/version"){
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
    },
    updated () {
      this.headerHeightStyle = `${document.getElementById('header').clientHeight}px`
    }
  }
</script>
<style lang="scss">
  @import './style/mixin.scss';

  html {
    font-size: 625%;
    -webkit-text-size-adjust: none;
    overflow-y: scroll;
    position: relative;
  }

  body, html {
    width: 100%;
    height: 100%;
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
    #router_wrap {
      .router_view {
        min-height: 4.5rem;
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
      footer {
        background: #363a3d;
        @include flex();
        @include pcContainer;
        padding: 0.3rem 0 0.2rem 0;
        .person_computer_footer { //分别对pc和移动端兼容
          @include pcCenter;
          flex-direction: row;
          justify-content: space-between;
          .footer_left {
            flex-direction: row;
          }
          .footer_right {
            .footer_copyright_wrap {
              text-align: end;
            }
          }

        }
        .mobile_footer {
          flex-direction: column;
          .footer_left {
            flex-direction: row;
            justify-content: center;
          }
          .footer_center {
            text-align: center;
            margin-bottom: 0.2rem;
          }
          .footer_right {
            .footer_link_wrap {
              align-items: center;
              justify-content: center;

            }
            .footer_copyright_wrap {
              text-align: center;
            }
          }
        }

        .person_computer_footer, .mobile_footer { //底部公用的样式
          @include flex;
          .footer_left {
            @include flex;
            margin-bottom: 0.2rem;
            span, a {
              width: 0.4rem;
              height: 0.4rem;
              @include borderRadius(0.3rem);
              margin-right: 0.1rem;
              cursor: pointer;
              &:last-child {
                margin-right: 0;
              }
            }
            a {
              text-decoration: none;
            }
            .github {
              @include linkBtn('./assets/github.png', './assets/github_h.png');
            }
            .airplane {
              @include linkBtn('./assets/airplane.png', './assets/airplane_h.png');
            }
            .twitter {
              @include linkBtn('./assets/twitter.png', './assets/twitter_h.png');
            }
            .facebook {
              @include linkBtn('./assets/facebook.png', './assets/facebook_h.png');
            }
            .we_chat {
              @include linkBtn('./assets/we_chat.png', './assets/we_chat_h.png');
            }
            .qq {
              @include linkBtn('./assets/qq.png', './assets/qq_h.png');
            }

          }
          .footer_right {
            @include flex;
            flex-direction: column;
            .footer_link_wrap {
              @include flex;
              span {
                font-size: 0.14rem;
                &:nth-child(2n-1) {
                  color: #3698db;
                  cursor: pointer;
                  margin-top: 0.02rem;
                }
                &:nth-child(2n) {
                  color: #a2a2ae;
                  margin: 0 0.1rem;
                }
              }
              .footer_link_privacy{
                a{
                  color: #3598db !important;
                }
              }
              .footer-faq{
                a{
                  color: #3598db !important;
                }
              }
            }
            .footer_copyright_wrap {
              color: #a2a2ae;
              margin-top: 0.1rem;
              font-size: 0.14rem;
            }

          }

        }
      }
    }

  }
  pre{
    font-family: Arial !important;
    font-size: 0.14rem !important;
  }
</style>
