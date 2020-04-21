declare var require: any;
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueRouter from 'vue-router';
import store from './store'
import './style/reset.css'
import '../icon/iconfont.css'
import Umeng from "./umeng/umeng"
/*引入自定义修饰器*/
import directives from './directives';
Vue.use(directives);
/*引入自定义过滤器*/
import filters from './filters';
Vue.use(filters);
/*引入自定义组件*/
import MTable from './components/commontables/MTable.vue';
Vue.component('MTable', MTable);
import Tools from './util/Tools';
import 'element-ui/lib/theme-chalk/index.css'
import {Select} from 'element-ui'
import {DatePicker} from 'element-ui'
import {Option} from 'element-ui'
import {Tooltip} from 'element-ui'
import {Cascader} from 'element-ui'
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
locale.use(lang);
Vue.component('el-select',Select);
Vue.component('el-option',Option);
Vue.component('el-date-picker',DatePicker);
Vue.component('el-tooltip',Tooltip);
Vue.component('el-cascader',Cascader);
Vue.use(locale);


import axios from "axios"
const  codec  = require("irisnet-crypto/util/codec.js") ;
const  crypto = require("irisnet-crypto");
Vue.prototype.$Crypto = crypto;
Vue.prototype.$Codec = codec;
Vue.prototype.$uMeng = Umeng;

Vue.config.productionTip = false;
axios.defaults.timeout = 120000;
axios.interceptors.request.use(function (config) {
  return config;
}, function (error) {
  return Promise.reject(error);
});
sessionStorage.setItem("Show_faucet",JSON.stringify(0));
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
