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

console.log(faucet_url)

Vue.prototype.faucet_url = faucet_url

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
