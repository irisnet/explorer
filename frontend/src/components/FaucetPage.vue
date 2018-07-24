<template>
  <b-container type="light">
    <h3 style="font-size:1.8rem;">Irisnet Testnet Faucet</h3>
    <div class="faucet text-center" style="">
      <div class="coin" style="display:flex;justify-content: center;margin-bottom:10px;">
        <img src="../assets/coin.png" alt="">
      </div>
      <p style="font-size:1.4rem;">Use this faucet to get tokens for the latest IRISnet testnet.</p>
      <p style="font-size:1.4rem;">Please don't abuse this service — the number of available tokens is limited.</p>
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
  </b-container>
</template>

<script>
  import axios from 'axios'

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
    data() {
      return {
        faucet_url: this.faucet_url,
        address: "",
        errMsg: "",
        alertShow:'hidden',
      }
    },
    created: function () {
      let nvc = document.createElement('script');
      nvc.setAttribute('src', "//g.alicdn.com/sd/nvc/1.1.112/guide.js");
      document.head.appendChild(nvc);

      nvc.onload = () => {
        let captcha = document.createElement('script');
        captcha.setAttribute('src', "//g.alicdn.com/sd/smartCaptcha/0.0.1/index.js");
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
      }
    }
  }

</script>

<style lang="scss">
  .faucet {
    background: white;
    padding: 15% 0
  }

  .faucet-form {
    margin: 0 auto;
    width: 300px;

    .form-group{
      text-align: left;
      .alert_information{
        display:inline-block;
        height:1.4rem;
        line-height:1.4rem;
        color:red;
        font-size:1.8rem;
        margin-top:0.5rem;
      }
      .form-control{
        height:2.8rem;
        line-height:2.8rem;
        font-size:1.4rem;
      }
    }

  }

  .err-msg {
    color: red
  }
</style>
