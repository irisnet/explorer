<template>
  <div type="light" class='facet_wrap' :style="showTitle?'':'padding-top:0.38rem;'">
    <h3 class='faucet_title' :style="`width:${innerWidth/100}rem;`" v-show="showTitle">
      <p class="title" :style="innerWidth<=500?'width:100%;padding-left:0.1rem;':''">
        IRISnet Testnet Faucet
      </p>
  </h3>
    <div class="faucet text-center" :style="innerWidth<=500?'padding-top:0;':''">
      <div class="coin" style="display:flex;justify-content: center;margin-bottom:10px;">
        <img src="../assets/coin.png" alt="">
      </div>
      <p style="font-size:0.14rem;color:#A2A2AE;padding:0 0.1rem;">Use this faucet to get tokens for the latest IRISnet testnet.</p>
      <p style="font-size:0.14rem;color:#A2A2AE;padding:0 0.1rem;">Please don't abuse this service — the number of available tokens is limited.</p>
      <br/>
      <form @submit.prevent="apply">
        <div class="faucet-form">
          <input type="text" class="form-control" id="token" name="token" hidden>
          <input type="text" class="form-control" id="session_id" name="session_id" hidden>
          <input type="text" class="form-control" id="sig" name="sig" hidden>
          <fieldset class="form-group">
            <input type="text" class="form-control" id="address" v-model="address" placeholder="Please enter the collection address">
            <span class="alert_information" :style="`visibility:${alertShow}`">Please enter a valid address</span>
          </fieldset>
          <fieldset class="form-group">
            <div id="sc" style="margin:0 auto;" class="text-left">
            </div>
          </fieldset>
          <!--<div class="text-left" v-if="errMsg">
            <small class="err-msg">* {{errMsg}}</small>
          </div>-->
          <button id="submit" type="submit" class="btn btn-primary" disabled>Send me IRIS</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import Tools from '../common/Tools';

  window.NVC_Opt = {
    appkey: 'FFFF0N000000000063E3',
    scene: 'ic_activity_h5',
    renderTo: '#captcha',
    trans: {"key1": "code0", "nvcCode": 200},
    elements: [
      '//img.alicdn.com/tfs/TB17cwllsLJ8KJjy0FnXXcFDpXa-50-74.png',
      '//img.alicdn.com/tfs/TB17cwllsLJ8KJjy0FnXXcFDpXa-50-74.png'
    ],
    bg_back_prepared: '//img.alicdn.com/tps/TB1skE5SFXXXXb3XXXXXXXXXXXX-100-80.png',
    bg_front: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABQCAMAAADY1yDdAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAADUExURefk5w+ruswAAAAfSURBVFjD7cExAQAAAMKg9U9tCU+gAAAAAAAAAIC3AR+QAAFPlUGoAAAAAElFTkSuQmCC',
    obj_ok: '//img.alicdn.com/tfs/TB1rmyTltfJ8KJjy0FeXXXKEXXa-50-74.png',
    bg_back_pass: '//img.alicdn.com/tfs/TB1KDxCSVXXXXasXFXXXXXXXXXX-100-80.png',
    obj_error: '//img.alicdn.com/tfs/TB1q9yTltfJ8KJjy0FeXXXKEXXa-50-74.png',
    bg_back_fail: '//img.alicdn.com/tfs/TB1w2oOSFXXXXb4XpXXXXXXXXXX-100-80.png',
    upLang: {
      "cn": {
        _ggk_guide: "请在屏幕上滑动，刮出两面盾牌",
        _ggk_success: "恭喜您成功刮出盾牌<br/>继续下一步操作吧",
        _ggk_loading: "加载中",
        _ggk_fail: ['呀，盾牌不见了<br/>请', "javascript:noCaptcha.reset()", '再来一次', '或', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", '反馈问题'],
        _ggk_action_timeout: ['我等得太久啦<br/>请', "javascript:noCaptcha.reset()", '再来一次', '或', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", '反馈问题'],
        _ggk_net_err: ['网络实在不给力<br/>请', "javascript:noCaptcha.reset()", '再来一次', '或', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", '反馈问题'],
        _ggk_too_fast: ['您刮得太快啦<br/>请', "javascript:noCaptcha.reset()", '再来一次', '或', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", '反馈问题']
      },
      "en": {
        _ggk_guide: "Please swipe on the screen and scrape off both shields",
        _ggk_success: "Congratulations on successfully scraping the shield<br/>Move on",
        _ggk_loading: "Loading...",
        _ggk_fail: ['Ah, the shield is gone<br/>Please', "javascript:noCaptcha.reset()", 'one more time', 'or', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", 'feedback problems'],
        _ggk_action_timeout: ['I\'ve been waiting too long<br/>Please', "javascript:noCaptcha.reset()", 'one more time', 'or', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", 'feedback problems'],
        _ggk_net_err: ['The Internet isn\'t working<br/>Please', "javascript:noCaptcha.reset()", 'one more time', 'or', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", 'feedback problems'],
        _ggk_too_fast: ['You\'re shaving too fast<br/>Please', "javascript:noCaptcha.reset()", 'one more time', 'or', "http://survey.taobao.com/survey/QgzQDdDd?token=%TOKEN", 'feedback problems']

      }
    }
  }

  window.onload = function () {

  }
  export default {
    name: "FaucetPage",
    $route() {
      this.showTitle = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
    },
    data() {
      return {
        faucet_url: this.faucet_url,
        address: "",
        errMsg: "",
        alertShow:'hidden',
        innerWidth : window.innerWidth,
        showTitle: !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
      }
    },
    created: function () {
      let nvc = document.createElement('script');
      nvc.setAttribute('src', "//g.alicdn.com/sd/nvc/1.1.112/guide.js");
      document.head.appendChild(nvc);

      nvc.onload = () => {
        let captcha = document.createElement('script');
        captcha.setAttribute('src', "//g.alicdn.com/sd/smartCaptcha/0.0.3/index.js");
        document.body.appendChild(captcha);
        captcha.onload = () => {
          var ic = new smartCaptcha({
            width: 300,
            renderTo: '#sc',
            default_txt: 'I\'m not a bot',
            success_txt: 'Authentication success!',
            fail_txt: 'Authentication failure! Clock again',
            scaning_txt: 'Intelligent authenticating...',
            success: function (data) {
              document.getElementById("token").value = NVC_Opt.token;
              document.getElementById("session_id").value = data.sessionId;
              document.getElementById("sig").value = data.sig;
              document.getElementById("submit").removeAttribute("disabled");
            },
          });
          ic.init();
        }
      }
      let addr = this.$route.query.address;
      if (addr && addr !== "") {
        this.address = addr
      }
    },
    methods: {
      apply() {
        /*if (document.getElementById("address").value === "") {
          alert("address is empty");
          return false;
        }*/
        if(!this.address && this.alertShow === 'hidden'){
          this.alertShow = 'visible';
          setTimeout(()=>{
            if(this.alertShow === 'visible'){
              this.alertShow= 'hidden';
            }
           },2000);
        }
        axios.post(this.faucet_url + '/apply', JSON.stringify({
          address: document.getElementById("address").value,
          token: document.getElementById("token").value,
          session_id: document.getElementById("session_id").value,
          sig: document.getElementById("sig").value
        })).then(result => {
          let data = result.data;
          if (data.err_code) {
            this.errMsg = data.err_msg
          } else {
            alert("apply successfully");
            location.reload();
          }
        })
      },
      resize(){
        this.innerWidth = window.innerWidth;
      }
    },
    mounted(){
      window.addEventListener('resize',this.resize)
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.resize);
    }
  }

</script>

<style lang="scss">
  @import '../style/mixin.scss';

  .faucet {
    background: white;
    padding:1rem 0;
  }

  .faucet-form {
    margin: 0 auto;
    width: 3rem;

    .form-group{
      text-align: left;
      margin-bottom:0 !important;
      .alert_information{
        display:inline-block;
        height:0.14rem;
        line-height:0.14rem;
        color:red;
        font-size:0.14rem;
        margin-top:0.05rem;
      }
      .form-control{
        height:0.36rem;
        line-height:0.36rem;
        font-size:0.14rem;
        @include borderRadius(0.04rem)
      }
    }
    .btn-primary{
      margin-top:0.2rem;
      box-shadow: 0 0 0 transparent;
      @include borderRadius(0.04rem);
      width:1.28rem;
      height:0.36rem;
      background:#3498DB;
      &:disabled{
        background: rgba(214,217,224,1);
        border-color:rgba(214,217,224,1);
      }
    }


  }

  .err-msg {
    color: red
  }
  .btn{
    padding:0.0375rem 0.075rem !important;
    font-size:0.14rem !important;
  }
  .form-control{
    padding:0.0375rem 0.075rem !important;
    width:3rem !important;
    @include borderRadius(0.025rem);
    font-size:0.14rem !important;
  }
  a, input{
    -webkit-box-shadow: 0 0 0 transparent !important;
    -moz-box-shadow: 0 0 0 transparent !important;
    box-shadow: 0 0 0 transparent !important;
  }

  .facet_wrap{

    .faucet_title{
      height:0.62rem;
      line-height:0.62rem;
      background:#efeff1;
      @include flex;
      justify-content:center;
      border-bottom:1px solid #d6d9e0;

      .title{
        width:80%;
        font-size:0.22rem;
        color:#000000;
        height:0.62rem;
      line-height:0.62rem;
      }
    }

  }
  #sc{
    height:0.52rem;

    #SM_BTN_WRAPPER_1{
      height:0.52rem;
      #SM_BTN_1{
        @include flex;
        align-items: center;
        height:0.52rem !important;
      }
      #sm-btn-bg{
        height:0.52rem !important;
      }
    }
  }
</style>
