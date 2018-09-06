import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './registerServiceWorker'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue);
Vue.config.productionTip = false;

let faucet_url;
let fuxi;

if(localStorage.getItem('Faucet_url') && localStorage.getItem('Faucet_url') !== "null"){
  faucet_url = localStorage.getItem('Faucet_url')
}else{
  faucet_url = ''
}

if(localStorage.getItem('Chain_id') && localStorage.getItem('Chain_id') !== "null"){
  fuxi = localStorage.getItem('Chain_id')
}else{
  fuxi = ''

}
Vue.prototype.faucet_url = faucet_url
Vue.prototype.fuxi = fuxi;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
