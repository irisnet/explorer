import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './registerServiceWorker'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)
Vue.config.productionTip = false

let faucet_url
let fuxi

if(localStorage.getItem('Faucet_url')){
  faucet_url = localStorage.getItem('Faucet_url')
}else{
  faucet_url = 'http://dev.faucet.irisplorer.io'
}

if(localStorage.getItem('Chain_id')){
  fuxi = localStorage.getItem('Chain_id')
}else{
  fuxi = 'rainbow-dev'
}
Vue.prototype.faucet_url = faucet_url
Vue.prototype.fuxi = fuxi;
String.prototype.trim = function() {
  return  this.replace(/^\s+|\s+$/g, '');
}

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
