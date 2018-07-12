<template>
  <div id="app">
    <app-header></app-header>
    <div id="router_wrap">
      <router-view class="router_view"/>
      <footer :class="footerClass">
        <div :class="footerClassName" style="height:100%;">
          <div class="footer_left" :class="footerLeftVar">
            <div class="footer_logo_wrap">
              <a href="https://www.baidu.com">logo</a>
            </div>
            <div class="footer_description_wrap">
              Inter-chain Service Infrastructure and Protocol Technology Foundation for a Distributed Business Ecosystem
            </div>
          </div>
          <div class="footer_right" :class="footerRightVar">
            <div class="footer_link_wrap">
              <a href="http://www.baidu.com" class="footer_link_wrap_item" style="padding-left:0;">Join IRISnet</a>
              <span class="footer_link_separate">|</span>
              <a href="#" class="footer_link_wrap_item">Forum</a>
              <span class="footer_link_separate">|</span>
              <a href="#" class="footer_link_wrap_item">Validators Chat</a>
              <span class="footer_link_separate">|</span>
              <a href="#" class="footer_link_wrap_item" style="padding-right:0;">Contact Us</a>
            </div>
            <div class="footer_copyright_wrap">
              ©️ IRIS Explorer 2018 all nights reserved
            </div>
          </div>
        </div>
      </footer>
    </div>

  </div>
</template>
<script>
  import AppHeader from './components/AppHeader';

  export default {
    components: {
      AppHeader,
    },
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
      if (this.devicesWidth > 500) {
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
      overflow: auto;
      .router_view{
        min-height:53.2rem;
      }
      .person_computer_wrap_footer{
        height: 10rem;
      }
      .mobile_wrap_footer{
        height:18rem;
      }
      footer {

        background: #222222;
        @include flex();
        flex-direction: column;
        align-items: center;
        padding:1rem 0;
        .person_computer_footer {
          width: 80%;
          min-width: 70rem;
          max-width: 100rem;
          @include flex();
          flex-direction:row;
          justify-content: space-between;
          .person_computer_footer_left{
            width:50%;
            @include flex();
            margin-right:1.5rem;
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
        .person_computer_footer, .mobile_footer{
          .footer_left{
            .footer_logo_wrap{
              width:18rem;
              height:6rem;
              border:0.1rem solid red;
              margin-right:1.5rem;
              a{
                display:inline-block;
                text-decoration: none;
                width:100%;
                height:100%;
              }
            }
            .footer_description_wrap{
              flex:1;
              color:#cccccc;
              font-size:1.3rem;
              padding:0.5rem;
            }
          }
          .footer_right{
            .footer_link_wrap{
              height:4rem;
              .footer_link_wrap_item{
                display:inline-block;
                padding:0.5rem 1rem;
                color:#fff;
                font-weight: 300;
                font-size:1.2rem;
                &:hover{
                  text-decoration: none;
                }
              }
              .footer_link_separate{
                color:#eee;
              }

            }
            .footer_copyright_wrap{
              height:4rem;
              color:#cccccc;
              font-size:1.3rem;
            }
          }
        }
      }
    }

  }
</style>
