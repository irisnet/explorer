<template>
  <div type="light" class='facet_wrap' :style="showTitle?'':'padding-top:0.38rem;'">
    <h3 class='faucet_title' :style="`width:${innerWidth/100}rem;`" v-show="showTitle">
      <p class="title" :style="innerWidth<=500?'width:100%;padding-left:0.1rem;':''">
        Faucet
      </p>
  </h3>
    <div class="faucet text-center" :style="innerWidth<=500?'padding-top:0;':''">
      <div class="coin" style="display:flex;justify-content: center;margin-bottom:10px;">
        <img src="../assets/coin.png" alt="">
      </div>
      <p style="font-size:0.14rem;color:#A2A2AE;padding:0 0.1rem;">Get IRIS from this faucet for the latest IRISnet Testnet.</p>
      <p style="font-size:0.14rem;color:#A2A2AE;padding:0 0.1rem;">This faucet will send 10 IRIS to any valid testnet address.</p>
      <p class="Balance_number" :class="errStyle ? 'err_red' : 'err_black ' ">Balance:{{faucetBalance}} {{tokenName}}</p>
      <br/>
      <form @submit.prevent="apply">
        <div class="faucet-form">
          <input type="text" class="form-control" id="token" name="token" hidden>
          <input type="text" class="form-control" id="session_id" name="session_id" hidden>
          <input type="text" class="form-control" id="sig" name="sig" hidden>
          <fieldset class="form-group">
            <input type="text" class="form-control" id="address" v-model="address" placeholder="Please enter the collection address">
            <div class="alert_information" :style="{visibility:alertShowErrMsg}">{{errMsg}}</div>

          </fieldset>
          <fieldset class="form-group">
            <div id="sc" style="margin:0 auto;" class="text-left">
            </div>
          </fieldset>
          <button id="submit" type="submit" class="btn btn-primary" :disabled="btnDisabled" :class="showSendingImg ? 'waitingStyle' : ''">
            {{btninfo}}
            <span v-show="showSendingImg" style="padding: 0 0.06rem">
              <img src="../assets/wait.gif">
            </span>
          </button>
        </div>
      </form>
    </div>
    <div class="alert_block" v-if="showAlert">
        <div class="img_font_container">
          <div class="img_container">
            <img v-if="showSuccess" src="../assets/success.png">
            <img v-if="!showSuccess" src="../assets/x.png" alt="">
          </div>
          <span class="font_style" v-if="showSuccess">Success !</span>
          <span class="font_style" v-if="!showSuccess">Failed,try again.</span>
        </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import Tools from '../common/Tools';

  let UserAgent = navigator.userAgent.toLowerCase();
  let scene = 'ic_other';
  if (/android/.test(UserAgent) || /iphone os/.test(UserAgent)) {
      scene = 'ic_activity_h5';
  }
  window.NVC_Opt = {
    appkey: 'FFFF0N000000000063E3',
    scene: scene,
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
    watch:{
      address(address){
        if(this.insufficientBalanceStatus === false){
          if(this.$Crypto.getCrypto("iris").isValidAddress(address)){
            if(this.verifyStatus === true ){
              this.btnDisabled = false;
            }
            this.verifyStatus = true;
            this.alertShowErrMsg = 'hidden';
          }else {
            this.btnDisabled = true;
            this.errMsg = "Please enter a valid address";
            this.alertShowErrMsg = 'visible';
          }
        }
      }
    },
    data() {
      return {
        faucet_url: this.faucet_url,
        address: "",
        errMsg: "",
        alertShowErrMsg:'hidden',
        innerWidth : window.innerWidth,
        showTitle: !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer()),
        faucetBalance: 0,
        insufficientBalanceStatus: false,
        tokenName:"",
        verifyStatus:false,
        errStyle: false,
        btnDisabled: true,
        btninfo:"Send me IRIS",
        showSendingImg: false,
        showAlert: false,
        showSuccess: false,
      }
    },
    beforeCreate(){

      let faucet_url = this.faucet_url + "/account";
      axios.get(faucet_url).then((data)=>{
        if(data.status === 200){
          return data.data;
        }
      }).then((data)=>{
        this.errStyle = false;
        this.btnDisabled = true;
        this.faucetBalance = Tools.formatBalance(Number(Tools.formatNumber(data.value.coins[0].amount)).toString().split(".")[0]);
        let faucetQuota = 20;
        if(this.faucetBalance < faucetQuota){
          this.errStyle = true;
          this.btnDisabled = true;
          this.btninfo = "Insufficient Balance";
          this.insufficientBalanceStatus = true
        }else {
          this.insufficientBalanceStatus = false;
          this.btninfo = "Send me IRIS"
        }
        this.tokenName = data.value.coins[0].denom.toUpperCase();
      }).catch(e =>{
        console.log(e);
        this.faucetBalance = "Error";
        this.errStyle = true;
        this.btnDisabled = true;
        this.insufficientBalanceStatus = true
      })
      },
    created: function () {
      let that = this;
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
              if(that.faucetBalance === "Error" || that.btninfo === "Insufficient Balance" || that.address === "" || that.verifyStatus === false){
                that.btnDisabled = true;
              }else {
                that.btnDisabled = false;
              }
              that.verifyStatus = true;
            },
          });
            ic.init();
        }
      };
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
        if(!this.errMsg && this.alertShowErrMsg === 'hidden'){
          this.alertShowErrMsg = 'visible';
          setTimeout(()=>{
            if(this.alertShowErrMsg === 'visible'){
              this.alertShowErrMsg= 'hidden';
              this.errMsg = "";
            }
          },2000);
        }
        this.btninfo = "Sending";
        this.btnDisabled = true;
        this.showSendingImg = true;
        axios.post(this.faucet_url + '/apply', JSON.stringify({
          address: document.getElementById("address").value,
          token: document.getElementById("token").value,
          session_id: document.getElementById("session_id").value,
          sig: document.getElementById("sig").value,
          "scene": scene
        })).then(result => {
          let data = result.data;
          let that = this;
          this.btninfo = "Send me IRIS";
          this.showSendingImg = false;
          this.showAlert = true;
          if (data.err_code) {
            this.showSuccess = false;
          } else {
            this.showSuccess = true;
          }
          setTimeout(function () {
            that.showAlert = false;
            that.showSuccess = false;
            location.reload();
          },2000)
        }).catch((e) => {
          console.log(e);
          this.btninfo = "Send me IRIS";
          this.showSendingImg = false;
          this.showAlert = true;
          this.showSuccess = false;
          let that = this;
          setTimeout(function () {
            that.showAlert = false;
            that.showSuccess = false;
            location.reload();
          },2000)
        });

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

<style lang="scss" scoped>
  @import '../style/mixin.scss';

  .faucet {
    background: white;
    padding: 0.38rem 0;
  }

  .faucet-form {
    margin: 0 auto;
    width: 3rem;

    .form-group{
      text-align: left;
      margin-bottom:0 !important;
      height: 63px;
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
      /*box-shadow: 0 0 0 transparent;*/
      @include borderRadius(0.04rem);
      padding: 0.11rem 0.19rem!important;
      background:#3498DB;
      @include fontSize;
      color: #fff;
      &:disabled{
        background: rgba(214,217,224,1);
        border-color:rgba(214,217,224,1);
      }
    }


  }
  .err_red{
    color: red !important;
  }
  .err_black{
    color: #000;
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
      width: 100%;
      height:0.62rem;
      line-height:0.62rem;
      background:#efeff1;
      @include flex;
      justify-content:center;
      border-bottom:1px solid #d6d9e0;
      .title{
        width:80%;
        max-width: 12.8rem;
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
  .Balance_number{
    font-size: 0.14rem;
    color: #000;
    padding-top: 0.04rem;
  }
  .faucet_address{
    font-size: 0.14rem;
    color: #000;
  }
  .btn-primary:focus{
    -webkit-box-shadow:0 0 0 .2rem rgba(255,255,255,.5);
    box-shadow:0 0 0 .2rem rgba(255,255,255,.5)
  }
  .alert_block{
    width: 2.54rem;
    height: 1.32rem;
    background:rgba(0,0,0,0.6);
    @include fixedCenter;
    border-radius: 0.1rem;
    z-index: 2000;
    .img_container{
      width: 36px;
      height:36px;
      margin: auto;
      margin-bottom: 0.1rem;
      img{
        width: 100%;
      }
    }
    .font_style{
      color: #fff;
      @include fontSize;
    }
    .img_font_container{
      @include center;
      padding-top: 0.02rem;
    }
  }
  .waitingStyle{
    color: #000!important;
  }
</style>
