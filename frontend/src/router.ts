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
import Nodespage from "./components/NodesPage.vue";
import SearchResult from "./components/searchResult.vue";
import Version from "./components/version.vue";

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
      path: '/nodespage', component: Nodespage
    },
    {
      path: '/faucet', component: FaucetPage
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

  ]

})
