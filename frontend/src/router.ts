import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import FaucetPage from './components/FaucetPage.vue'
import BlockPage from './components/BlockPage.vue'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/home',
    },
    {
      path: '/home',
      component: Home,
    },
    {
      path: 'faucet', component: FaucetPage
    },
    {
      path: 'block', component: BlockPage
    }
  ]

})
