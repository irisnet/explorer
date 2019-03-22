declare var require: any;
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import axios from "axios"
import '../icon/iconfont.css'
const  codec  = require("irisnet-crypto/util/codec.js") ;
const  crypto = require("irisnet-crypto");
Vue.use(BootstrapVue);
Vue.prototype.$Crypto = crypto;
Vue.prototype.$Codec = codec;
Vue.config.productionTip = false;
let faucet_url;
axios.defaults.timeout = 120000;
axios.interceptors.request.use(function (config) {
  return config;
}, function (error) {
  return Promise.reject(error);
});
Vue.prototype.faucet_url = faucet_url;
let currentServerTime = new Date().getTime();
axios.get(`/api/sysdate`).then(data => {
  if(data.status === 200){
    return data.data
  }
}).then(sysdate => {
    Vue.prototype.diffMilliseconds = sysdate.data*1000 - currentServerTime;
});


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
