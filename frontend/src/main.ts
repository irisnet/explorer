declare var require: any;
import Vue from 'vue'
import App from './App1.vue'
import router from './router'
import VueRouter from 'vue-router';
import store from './store'
// import BootstrapVue from 'bootstrap-vue'
import './style/reset.css'
import '../icon/iconfont.css'
// import 'bootstrap/dist/css/bootstrap.css'
// import 'bootstrap-vue/dist/bootstrap-vue.css'
/*引入自定义修饰器*/
import directives from './directives';
Vue.use(directives);
/*引入自定义过滤器*/
import filters from './filters';
Vue.use(filters);
/*引入自定义组件*/
import MTable from './components/commonComponents/MTable.vue';
import iView from 'iview';
import select from 'iview/src/components/select'
import datePicker from 'iview/src/components/date-picker'
import locale from 'iview/src/locale/lang/en-US'
import 'iview/dist/styles/iview.css'
Vue.use(iView,{locale},select,datePicker);
Vue.component('MTable', MTable);

import axios from "axios"
const  codec  = require("irisnet-crypto/util/codec.js") ;
const  crypto = require("irisnet-crypto");
// Vue.use(BootstrapVue);
Vue.prototype.$Crypto = crypto;
Vue.prototype.$Codec = codec;
Vue.config.productionTip = false;
axios.defaults.timeout = 120000;
axios.interceptors.request.use(function (config) {
  return config;
}, function (error) {
  return Promise.reject(error);
});
let currentServerTime = new Date().getTime();
axios.get(`/api/sysdate`).then(data => {
  if(data.status === 200){
    return data.data
  }
}).then(sysdate => {
    Vue.prototype.diffMilliseconds = sysdate.data*1000 - currentServerTime;


});
sessionStorage.setItem("Show_faucet",JSON.stringify(0));
import Tools from './util/Tools';
const routerPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return routerPush.call(this, location).catch(error=> error)
}


Vue.prototype.addressRoute = Tools.addressRoute;
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
