<template>
  <div id="app">
    <app-header></app-header>
    <div id="router_wrap">
      <router-view class="router_view"/>
      <footer :class="footerClass">
        <div :class="footerClassName" style="height:100%;">
          <div class="footer_left" :class="footerLeftVar">
            <div class="footer_logo_wrap">
              <a href="https://www.irisnet.org/">
                <img src="./assets/IRIS-logo.png" alt="">
              </a>
              <div class="footer_description_wrap">
                Inter-chain Service Infrastructure and Protocol Technology Foundation for a Distributed Business Ecosystem
              </div>
              <p class="footer_copyright_wrap">
                ©️ IRIS Explorer 2018 all nights reserved
              </p>
            </div>

          </div>
          <div class="footer_center">
            <div>
              <span>Latest Discussions</span>
              <span>More</span>
            </div>
            <div></div>
          </div>
          <div class="footer_right" :class="footerRightVar">
            <div class="footer_link_wrap">
              <span class="footer_link_contact">Contact Us</span>
              <span class="footer_link_forum">Forum</span>
              <span class="footer_link_chat">Validators Chat</span>
              <span class="footer_link_join">Join IRISnet</span>
              <span class="footer_link_privacy" @click="footerLinkClick('/privacy_policy')">Privacy Policy</span>
              <span class="footer_link_telegramy">Telegramy</span>
            </div>

          </div>
        </div>
      </footer>
    </div>

  </div>
</template>
<script>
  import AppHeader from './components/AppHeader';
  import Tools from './common/Tools';

  export default {
    components: {
      AppHeader,
    },
    /*watch:{
      $route(to,from){
        console.log(to);
        console.log(from)
      }
    },*/
    data() {
      return {
        show: 1,
        devicesWidth: window.innerWidth,
        devicesShow: 1,//1是显示pc端，0是移动端
        footerClass:'person_computer_wrap_footer',
        footerClassName: 'person_computer_footer',
        footerLeftVar:'person_computer_footer_left',
        footerRightVar:'person_computer_footer_right',
      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {
        this.devicesShow = 1;
        this.footerClass = 'person_computer_wrap';
        this.footerClassName = 'person_computer_footer';
        this.footerLeftVar = 'person_computer_footer_left';
        this.footerRightVar = 'person_computer_footer_right';
      } else {
        this.devicesShow = 0;
        this.footerClass = 'mobile_wrap_footer';
        this.footerClassName = 'mobile_footer';
        this.footerLeftVar = 'mobile_footer_left';
        this.footerRightVar = 'mobile_footer_right';
      }
    },
    methods:{
      footerLinkClick(path){
        this.$router.push(path);
      }
    }
  }
</script>
<style lang="scss">
  @import './style/mixin.scss';
  html{
    font-size:62.5%;
  }
  body, html {
    width: 100%;
    height: 100%;
  }

  #app {
    width: 100%;
    height: 100%;
    @include flex();
    flex-direction: column;
    #router_wrap {
      flex: 1;
      overflow-y: auto;
      overflow-x:hidden;
      .router_view{
        min-height:53.2rem;
      }
      .person_computer_wrap_footer{
        height: 10rem;
      }
      .mobile_wrap_footer{
        /*height:21rem;*/
      }
      footer {
        background: #363a3d;
        @include flex();
        @include pcContainer;
        padding-top:2rem;
        .person_computer_footer {
          @include pcCenter;
          @include flex();
          flex-direction:row;
          justify-content: space-between;
          .person_computer_footer_left{
            width:50%;
            @include flex();
          }
          .person_computer_footer_right{
            width:45%;
            @include flex();
            flex-direction: column;
            align-items:flex-end;
          }
        }
        .mobile_footer{
          @include flex();
          flex-direction:column;
          width:100%;
          .mobile_footer_left{
            @include flex();
            flex-direction:column;
            align-items:center;
          }
          .mobile_footer_right{
            @include flex();
            flex-direction:column;
            align-items:center;
          }
        }
        .person_computer_footer{//分别对pc和移动端兼容
          @include flex;
          .footer_left{
            flex:3;
          }
          .footer_center{
            flex:6;
            border:1px solid red;
          }
          .footer_right{
            flex:2;
          }
        }
        .mobile_footer{
          @include flex;
          flex-direction:column;
          align-items:center;
          .mobile_footer_left{
            .footer_logo_wrap{
              text-align: center;
            }
          }
        }
        .person_computer_footer, .mobile_footer{//底部公用的样式
          .footer_left{
            .footer_logo_wrap{
              .footer_description_wrap{
                flex:1;
                color:#a2a2ae;
                font-size:1.4rem;
                margin-top:2rem;
              }
              .footer_copyright_wrap{
                color:#fff;
                font-size:1.4rem;
                margin-top:2rem;
              }
              a{
                display:inline-block;
                text-decoration: none;
              }
            }

          }
          .footer_right{
            .footer_link_wrap{

              @include flex;
              flex-direction:column;
              span{
                color:#3598db;
                cursor:pointer;
                font-size:1.4rem;
                margin-bottom:1.4rem;
                text-indent:3rem;
              }
              .footer_link_contact{
                background: url(./assets/message.svg) no-repeat 0 0;
              }
              .footer_link_forum{
                background: url(./assets/forum.svg) no-repeat 0 0;
              }
              .footer_link_chat{
                background: url(./assets/chat.svg) no-repeat 0 0;
              }
              .footer_link_join{
                background: url(./assets/join.svg) no-repeat 0 0;
              }
              .footer_link_privacy{
                background: url(./assets/policy.svg) no-repeat 0 0;
              }
              .footer_link_telegramy{
                background: url(./assets/telegramy.svg) no-repeat 0 0;
              }


            }

          }

        }
      }
    }

  }
</style>
