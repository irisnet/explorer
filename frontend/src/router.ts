import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import FaucetPage from './components/FaucetPage.vue';
import TransactionsDetail from './components/TransactionsDetail.vue';
import BlocksDetail from './components/BlocksDetail.vue';
import BlocksListPage from './components/BlocksListPage.vue';
import AddressPage from './components/AddressPage.vue';
import PrivacyPolicy from './components/PrivacyPolicy.vue';
import ProposalsPage from "./components/ProposalsPage.vue";
import ProposalsDetail from "./components/ProposalsDetail.vue";
import SearchResult from "./components/searchResult.vue";
import Help from "./components/clearStorageHelpPage.vue";
import Version from "./components/version.vue";
import ValidatorListPage from "./components/ValidatorListPage.vue"
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
      path: '/faucet',
      component: FaucetPage,
      beforeEnter (to, from, next) {
        if (localStorage.getItem('Show_faucet') === '0'){
          next(false)
        }else{
          next()
        }
      }
    },
    {
      path: '/Proposals', component: ProposalsPage
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
      path: '/transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/transactions/:type/:param', component: BlocksListPage
    },
    {
      path: '/validators', component: ValidatorListPage
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
      path: '/ProposalsDetail/:proposal_id', component: ProposalsDetail,
    },
    {
      path:'/address/:type/:param',
      component:AddressPage,
    },
    {
      path:'/privacy_policy',
      component:PrivacyPolicy,
    },
    {
      path: '/searchResult/:searchContent', component: SearchResult,
    },
    {
      path: '/searchResult/', component: SearchResult,
    },
    {
      path: '/version', component: Version,
    },
    {
      path: '/help', component: Help,
    },
  ]

})
