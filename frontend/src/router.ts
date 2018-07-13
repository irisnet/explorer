import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import FaucetPage from './components/FaucetPage.vue'
import BlockPage from './components/BlockPage.vue'
import TransactionsDetail from './components/TransactionsDetail.vue'
import BlocksListPage from './components/BlocksListPage.vue'

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
      path: '/faucet', component: FaucetPage
    },
    {
      //BlocksListPage为一个组件，根据type类型不同相应不同页面
      //1 BlocksList页面
      path: '/block/:type/:param', component: BlocksListPage

    },
    {
      path: '/recent_transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/tx', component: TransactionsDetail
    },

  ]

})
