<template>
  <div id="app">
    <app-header></app-header>
    <div id="router_wrap">
      <div class="weChat_qr_code"
           @click="hideQRCode"
           :style="`width:${vw}px;height:${vh}px`" v-show="qrShow">
        <div class="weChat_all">
          <img src="./assets/wechat_qr.png" alt="失去网络了..." style="width:1.8rem;height:1.8rem;">
          <h2>Scan QR Code</h2>
          <p>to follow our SubScriptions</p>
        </div>
      </div>
      <router-view class="router_view"/>
      <footer :class="footerClass" v-show="showFooter">
        <div :class="footerClassName" style="height:100%;">
          <div class="footer_left" :class="footerLeftVar">
            <a target="_blank" href='https://github.com/irisnet' class="github"></a>
            <a target="_blank" href='https://t.me/irisnetwork' class="airplane"></a>
            <a target="_blank" href='https://twitter.com/irisnetwork' class="twitter"></a>
            <a target="_blank" href='https://medium.com/irisnet-blog' class="facebook"></a>
            <span class="we_chat" @click="showQRCode"></span>
            <span class="qq"></span>
          </div>
          <div class="footer_center">
            <a href="https://www.irisnet.org/" target="_blank">
              <img src="./assets/IRIS-logo.png" alt="">
            </a>
          </div>
          <div class="footer_right" :class="footerRightVar">
            <div class="footer_link_wrap">
              <span class="footer_link_contact">Join Testnet</span>
              <span class="footer_link_forum">|</span>
              <span class="footer_link_chat">Contact Us</span>
              <span class="footer_link_join">|</span>
              <span class="footer_link_privacy" @click="footerLinkClick('/privacy_policy')">Privacy Policy</span>
            </div>
            <p class="footer_copyright_wrap">
              ©️ IRIS Explorer 2018 all rights reserved
            </p>
            <p class="footer_copyright_wrap">
              {{`${version}(${env}_${build})`}}
            </p>

          </div>
        </div>
      </footer>
    </div>

  </div>
</template>
<script>
  import AppHeader from './components/AppHeader';
  import Tools from './common/Tools';
  import testVersion from '../testVersion';

  export default {
    components: {
      AppHeader,
    },
    watch: {
      $route() {
        this.showFooter = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
        document.getElementById('router_wrap').scrollTop = 0;
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
        showFooter: !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
        vh: window.innerHeight,
        vw: window.innerWidth,
        qrShow: false,
        build: testVersion.app.buildNumber,
        env: testVersion.app.env,
        version: testVersion.app.version,
        innerWidth: window.innerWidth,
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        document.body.style.minWidth = '4rem';
      }

    },
    mounted() {
      window.addEventListener('resize', this.onresize);
      if (window.innerWidth > 960) {
        this.footerClass = 'person_computer_wrap';
        this.footerClassName = 'person_computer_footer';
        this.footerLeftVar = 'person_computer_footer_left';
        this.footerRightVar = 'person_computer_footer_right';
      } else {
        this.footerClass = 'mobile_wrap_footer';
        this.footerClassName = 'mobile_footer';
        this.footerLeftVar = 'mobile_footer_left';
        this.footerRightVar = 'mobile_footer_right';
      }
    },
    beforeDestroy() {
      window.removeEventListener('resize', this.onWindowResize);
    },
    methods: {
      footerLinkClick(path) {
        this.$router.push(path);
      },
      onresize() {
        this.innerWidth = window.innerWidth;
        this.vh = window.innerHeight;
        this.vw = window.innerWidth;
        if (window.innerWidth > 960) {
          this.footerClass = 'person_computer_wrap';
          this.footerClassName = 'person_computer_footer';
          this.footerLeftVar = 'person_computer_footer_left';
          this.footerRightVar = 'person_computer_footer_right';
        } else {
          this.footerClass = 'mobile_wrap_footer';
          this.footerClassName = 'mobile_footer';
          this.footerLeftVar = 'mobile_footer_left';
          this.footerRightVar = 'mobile_footer_right';
        }
      },
      showQRCode() {
        this.qrShow = true;
      },
      hideQRCode() {
        this.qrShow = false;
      }
    }
  }
</script>
<style lang="scss">
  @import './style/mixin.scss';

  html {
    font-size: 625%;
    -webkit-text-size-adjust: none;
  }

  body, html {
    width: 100%;
    height: 100%;
  }

  body {
    font-size: 16px !important;
    font-family:Arial;
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
    @include flex();
    flex-direction: column;
    #router_wrap {
      flex: 1;
      overflow-y: auto;
      overflow-x: auto;
      .router_view {
        min-height: 4.5rem;
      }
      .person_computer_wrap_footer {
        height: 1rem;
      }
      .mobile_wrap_footer {
        /*height:21rem;*/
      }
      .weChat_qr_code { //微信二维码弹框
        position: fixed;
        top: 0;
        left: 0;
        background: rgba(0, 0, 0, .6);
        z-index: 1000;
        @include flex;
        justify-content: center;
        align-items: center;
        .weChat_all {
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
          }
          .footer_center {
            text-align: center;
            margin-bottom: 0.1rem;
          }
          .footer_right {
            .footer_link_wrap {
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
            margin-bottom: 0.1rem;
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
                }
                &:nth-child(2n) {
                  color: #a2a2ae;
                  margin: 0 0.1rem;
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
</style>
