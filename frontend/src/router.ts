import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import FaucetPage from './components/FaucetPage.vue';
import TransactionsDetail from './components/TransactionsDetail.vue';
import BlocksDetail from './components/BlocksDetail.vue';
import BlocksListPage from './components/BlocksListPage.vue';
import AddressPage from './components/AddressPage.vue';
import PrivacyPolicy from './components/PrivacyPolicy.vue';

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
      //2 Transactions页面
      //3 Validators页面
      //4 Candidates页面
      path: '/block/:type/:param', component: BlocksListPage

    },
    {
      path: '/recent_transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/transfer_transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/stake_transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/validators/:type/:param', component: BlocksListPage
    },
    {
      path: '/candidates/:type/:param', component: BlocksListPage
    },
    {
      path: '/tx', component: TransactionsDetail
    },
    {
      path: '/blocks_detail/:height', component: BlocksDetail,
    },

    {
      path:'/address/:type/:param',
      component:AddressPage,
    },
    {
      path:'/privacy_policy',
      component:PrivacyPolicy,
    },
  ]

})
