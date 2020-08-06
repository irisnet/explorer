declare var require: any;
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueRouter from 'vue-router';
import store from './store';
import VueI18n from 'vue-i18n'
import './style/reset.css';
import '../icon/iconfont.css';
import Umeng from "./umeng/umeng";
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
// import ElementUI from "element-ui"
import {
    Select,
    DatePicker,
    Option,
    Tooltip,
    Cascader,
    Table,
    TableColumn,
    Message,
    Input} from 'element-ui'

import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
locale.use(lang);
Vue.use(locale);

Vue.component('el-select',Select);
Vue.component('el-option',Option);
Vue.component('el-date-picker',DatePicker);
Vue.component('el-tooltip',Tooltip);
Vue.component('el-cascader',Cascader);
Vue.use(Input);
Vue.use(Table);
Vue.use(TableColumn);
Vue.prototype.$message = Message;

// Vue.use(ElementUI)

Vue.use(VueI18n)
const i18n = new VueI18n({
  locale:'EN',
  messages: {
    "CN" : require('./lang/CN-Cindy'),
    "EN" : require('./lang/EN-Cindy')
  }
})

/*引入轮播*/
import VueAwesomeSwiper  from 'vue-awesome-swiper'
import 'swiper/css/swiper.css'
Vue.use(VueAwesomeSwiper);
import Swipe  from 'vue-awesome-swiper';
import SwipeItem  from 'vue-awesome-swiper';
Vue.component('swipe', Swipe);
Vue.component('swipe-item', SwipeItem);

import axios from "axios"
const  crypto = require("irisnet-crypto");
Vue.prototype.$Crypto = crypto;

Vue.prototype.$uMeng = Umeng;

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
axios.get(`/api/block/blockinfo/1`).then(data => {
  if(data.status === 200){
    return data.data
  }
}).then(blockinfo => {
  let firstBlockTime = new Date(blockinfo.timestamp.split('T')[0]).getTime() - 24 * 60 * 60 * 1000;
  sessionStorage.setItem('firstBlockTime',Tools.formatDateYearToDate(firstBlockTime))
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
  i18n,
  render: h => h(App)
}).$mount('#app')
