<template>
      <div class="clear_storage_help_container" :class="clearStroageDetailWrap">
        <div class="clear_storage_title_wrap">
          <p><span>{{title}}</span></p>
        </div>
        <div class="content_wrap">
          <div v-for="(item,index) in clearBrowserCacheSteps" class="content">
            <div class="clear_storage_content">
              <div class="title_content">
                <div class="logo_img">
                  <img :src="item.logoImg">
                </div>
                <span>{{item.name}}</span>
              </div>
              <div :class="index == 1 ? '' : 'line' " class="step_container">
                <p class="step">
                  <span class="step_num">{{item.stepNum.one}}</span>
                  <span class="step_content">{{item.stepInfo.one}}</span>
                </p>
                <div class="step_img">
                  <img :src="item.stepImg.one">
                </div>
              </div>
              <div class="step_container" :class="index == 3 || index == 4 ? 'line' : ''">
                <p class="step" v-show="item.stepNum.two">
                  <span class="step_num">{{item.stepNum.two}}</span>
                  <span class="step_content">{{item.stepInfo.two}}</span>
                </p>
                <div class="step_img" v-show="item.stepImg.two">
                  <img :src="item.stepImg.two">
                </div>
              </div>
              <div class="step_container">
                <p class="step" v-show="item.stepNum.three">
                  <span class="step_num">{{item.stepNum.three}}</span>
                  <span class="step_content">{{item.stepInfo.three}}</span>
                </p>
                <div class="step_img" v-show="item.stepImg.three">
                  <img :src="item.stepImg.three">
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
</template>

<script>
  import Tools from "../../util/Tools";
    export default {
        name: "clearStorageHelpPage",
      data(){
          return{
            clearBrowserCacheSteps:[
              {
                logoImg:require('../../assets/google_logo.png'),
                name:"Google Chrome",
                stepNum:{
                  one: "1",
                  two: "2",
                },
                stepInfo:{
                  one: "Press F12 or right click on check on the page.",
                  two: "Find “Application” and then \"Service Workers\"；find the service worker.js under the testnet.irisplorer.cn or testnet.irisplorer.io. Click \"Unregister\" and revisit as shown below:"
                },
                stepImg:{
                  one: require("../../assets/google_step.jpg"),
                  two: require("../../assets/google_02.jpg")
                }
              },

              {
                logoImg:require('../../assets/safari_logo.png'),
                name:"Safari PC",
                stepNum:{
                  one: "1",
                },
                stepInfo:{
                  one: "Open Safari on your Mac; click the Safari menu at the top left of the screen；then click \"Preferences\" in the menu. Click \"Privacy\"on the top, and find \"manage your website data\".Open \"Manage Website Data\" and find the data under testnet.irisplorer.cn or testnet.irisplorer.io; delete it and click \"Finish\".",
                },
                stepImg:{
                  one: require("../../assets/safari_pc.jpg"),
                }
              },

              {
                logoImg:require('../../assets/safari_logo.png'),
                name:"Safari Mobile Client",
                stepNum:{
                  one: "1",
                  two: "2",
                },
                stepInfo:{
                  one: "First, click “Settings” and then click “Safari” .",
                  two: "Click \"clear historical data and website data\n\"Click \"yes\", the color of the words will change into gray，and then into blue. The catching has been cleared."
                },
                stepImg:{
                  one: require("../../assets/safari-mobile-client_01.jpg"),
                  two: require("../../assets/safari-mobile-client_02.jpg")
                }

              },
              {
                logoImg:require('../../assets/firefox.png'),
                name:"Firefox",
                stepNum:{
                  one: "1",
                  two: "2",
                  three: "3",
                },
                stepInfo:{
                  one: "Click \"menu\" on the top left and then click \"web developer\"",
                  two: "Click \"service work\"",
                  three: "Find \"Service Workers\" under the testnet.irisplorer.cn or testnet.irisplorer.io and click \"unregister\". Revisit as shown below:",
                },
                stepImg:{
                  one: require("../../assets/firefox_01.jpg"),
                  two: require("../../assets/firefox_02.jpg"),
                  three: require("../../assets/firefox_03.jpg"),
                }
              },
              {
                logoImg:require('../../assets/android.png'),
                name:"Mobile browser for Andriod",
                stepNum:{
                  one: "1",
                  two: "2",
                  three: "3",
                },
                stepInfo:{
                  one: "Open chrome://inspect/#devices with Google Chrome and check \"Discover USB devices\" (Note: The first time you need to use the Internet).",
                  two: "Click developer options and allow USB debugging. Connect to the computer with a data cable; open the browser and find the page under the testnet.irisplorer.cn or testnet.irisplorer.io, then open chrome://inspect/#devices through Google Chrome.You will see the testnet.irisplorer.cn or testnet.irisplorer.io on the chrome://inspect/#devices page on your phone, click \"inspect\" below.",
                  three: "In the newly opened page, find \"Application and then \"Service Workers\"; find \"service worker.js\" under the testnet.irisplorer.cn or testnet.irisplorer.io. Click \"Unregister\" and revisit as shown below:",
                },
                stepImg:{
                  one: require("../../assets/andriod_01.jpg"),
                  two: require("../../assets/andriod_02.jpg"),
                  three: require("../../assets/andriod_03.jpg"),
                }
              }
            ],
            title:"How could I clear the browser cache?",
            clearStroageDetailWrap:"",
          }
      },
      beforeMount() {
        Tools.scrollToTop();
        if (Tools.currentDeviceIsPersonComputer()) {
          this.clearStroageDetailWrap = 'computer_clear_storage_detail_wrap';
        } else {
          this.$set(this.clearBrowserCacheSteps[2],"name","Safari Mobile");
          this.$set(this.clearBrowserCacheSteps[4],"name","Android Browser");
          this.clearStroageDetailWrap = 'mobile_clear_storage_detail_wrap';
        }
      },
      watch:{
          $route(){
            Tools.scrollToTop();
          }
      }
    }
</script>

<style scoped lang="scss">
  @import "../../style/mixin";
  .mobile_clear_storage_detail_wrap{
    .clear_storage_title_wrap{
      p{
        font-size: 0.15rem !important;
        span{
          font-size: 0.18rem !important;
        }
      }
    }
    .step_num{
      width: 0.24rem;
      height: 0.24rem;
      border-radius: 0.12rem;
      font-size: 0.12rem;
      text-align: center;
      line-height: 0.24rem;
    }
    .step_content{
      margin-left: 0.1rem;
      line-height: 0.18rem;
    }
    .logo_img{
      padding-left: 0.1rem!important;
    }
    .step{
      display: flex;
      position: relative;
      left: -0.13rem!important;
    }
    .step_container{
      margin-left: 0.2rem;
    }
    .step_img{
      width: 100%!important;
      padding: 0.2rem 0.1rem 0.3rem 0.17rem;
      img{
        width: 100%;
      }
    }
  }
.clear_storage_help_container{
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  -webkit-box-align: center;
  .clear_storage_title_wrap{
    width: 100%;
    height: 0.62rem;
    line-height: 0.62rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    -webkit-box-align: center;
    background: #F5F7FD;
    p{
      width: 100%;
      max-width: 12.8rem;
      font-size: 0.22rem;
      color: #000000;
      font-weight: 500 !important;
      span{
        margin-left: 0.2rem;
      }
    }

  }
  .content_wrap{
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    .content{
      width: 100%;
      max-width: 12.8rem;
    }
  }

  .clear_storage_content{
    max-width: 12.8rem;
    width: 100%;
    .title_content{
      display: flex;
      height: 0.71rem;
      line-height: 0.71rem;
      border-bottom: 0.01rem solid #d6d9e0;
      margin-bottom: 0.2rem;
      .logo_img{
        padding-left: 0.2rem;
      }
      span{
        margin-left: 0.12rem;
        font-size: 0.18rem;
      }
    }
    .step{
      display: flex;
      position: relative;
      left: -0.16rem;
    }
  }
}
  .step_container{
    margin-left: 0.4rem;
    border-left: 0.01rem solid #fff;
  }
  .line{
    border-left: 0.01rem solid #d6d9e0;
  }
  .step_num{
    display: inline-block;
    min-width: 0.24rem;
    width: 0.3rem;
    height: 0.3rem;
    text-align: center;
    line-height: 0.3rem;
    border-radius: 0.15rem;
    background: var(--bgColor);
    color: #fff;
  }
  .step_content{
    display: inline-block;
    max-width: 8rem;
    margin-left: 0.18rem;
    line-height: 0.3rem;
    font-size: 0.14rem;
    color: var(--contentColor);
  }
  .step_img{
    padding: 0.2rem 0  0.6rem 0.33rem;

  }
</style>
