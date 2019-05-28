declare var require: any;
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import axios from "axios"
const  codec  = require("irisnet-crypto/util/codec.js") ;
const  crypto = require("irisnet-crypto");
import {Table, Tooltip} from 'iview';
import 'iview/dist/styles/iview.css';    // 使用 CSS
Vue.component('iTable', Table);
Vue.component('iTooltip', Tooltip);

import MTable from './components/commonComponents/MTable.vue';
Vue.component('MTable', MTable);

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
sessionStorage.setItem("Show_faucet",JSON.stringify(0))
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
