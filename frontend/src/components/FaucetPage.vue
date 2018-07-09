<template>
  <b-container type="light">
    <h3>Irisnet Testnet Faucet</h3>
    <div class="faucet text-center" style="">
      <p>Use this faucet to get tokens for the latest Irisnet testnet.</p>
      <p>Please don't abuse this service — the number of available tokens is limited.</p>
      <br/>
      <form @submit.prevent="apply">
        <div class="faucet-form">
          <input type="text" class="form-control" id="token" name="token" hidden>
          <input type="text" class="form-control" id="session_id" name="session_id" hidden>
          <input type="text" class="form-control" id="sig" name="sig" hidden>
          <fieldset class="form-group">
            <input type="text" class="form-control" id="address" v-model="address" placeholder="Testnet address"
                   required>
          </fieldset>
          <fieldset class="form-group">
            <div id="sc" style="margin:0 auto;" class="text-left">
            </div>
          </fieldset>
          <div class="text-left" v-if="errMsg">
            <small class="err-msg">* {{errMsg}}</small>
          </div>
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
      }
    }
  }

  window.onload = function () {

  }
  export default {
    name: "FaucetPage",
    data() {
      return {
        faucet_url: "http://192.168.150.199:4000",
        address: "",
        errMsg: "",
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
        if (document.getElementById("address").value === "") {
          alert("address is empty");
          return false;
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
            this.errMsg = "apply successfully";
            setTimeout(location.reload(), 1000);
          }
        })
      }
    }
  }

</script>

<style>
  .faucet {
    background: white;
    padding: 15% 0
  }

  .faucet-form {
    margin: 0 auto;
    width: 300px;
  }

  .err-msg {
    color: red
  }
</style>
