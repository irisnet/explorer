import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import FaucetPage from './components/FaucetPage.vue'
import BlockPage from './components/BlockPage.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      children: [
        {
          path: 'faucet', component: FaucetPage
        },
        {
          path: 'block', component: BlockPage
        }
      ]
    },
  ]
})
