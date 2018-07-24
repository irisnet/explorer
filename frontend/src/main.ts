import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './registerServiceWorker'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)
Vue.config.productionTip = false

let faucet_url = process.env.VUE_APP_FAUCET_URL
if (!faucet_url || faucet_url === '') {
  faucet_url = 'http://dev.faucet.irisplorer.io'
}
Vue.prototype.faucet_url = faucet_url

let fuxi = process.env.VUE_APP_FUXI;
if(!fuxi || fuxi === ''){
  fuxi = 'fuxi-test'
}
Vue.prototype.fuxi = fuxi;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
