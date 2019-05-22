import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import FaucetPage from './components/FaucetPage.vue';
import TransactionsDetail from './components/TransactionsDetail.vue';
import BlocksDetail from './components/BlockInfo.vue';
import AddressPage from './components/AddressPage.vue';
import PrivacyPolicy from './components/PrivacyPolicy.vue';
import ProposalsPage from "./components/ProposalsPage.vue";
import ProposalsDetail from "./components/ProposalsDetail.vue";
import SearchResult from "./components/searchResult.vue";
import Help from "./components/clearStorageHelpPage.vue";
import Version from "./components/version.vue";
import ValidatorListPage from "./components/ValidatorListPage.vue";
import Parameters from "./components/ParametersPage.vue"
import RichList from "./components/RichList.vue"
import BlockList from "./components/BlockListPage.vue"
import TxList from "./components/TxListPage.vue"
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
      path: '/gov/proposals', component: ProposalsPage
    },
    {
      path: '/gov/parameters', component: Parameters
    },
    {
      //BlocksListPage为一个组件，根据type类型不同相应不同页面
      //1 BlocksList页面
      //2 Transactions页面
      //3 Validators页面
      //4 Candidates页面
      path: '/blocks', component: BlockList

    },
    {
      path: '/txs/:txType', component: TxList
    },
    {
      path: '/txs/:txType', component: TxList
    },
    {
      path: '/txs/:txType', component: TxList
    },
    {
      path: '/validators', component: ValidatorListPage
    },
    {
      path: '/tx', component: TransactionsDetail
    },
    {
      path: '/block/:height', component: BlocksDetail,
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
      path: '/searchResult', component: SearchResult,
    },
    {
      path: '/version', component: Version,
    },
    {
      path: '/help', component: Help,
    },
    {
      path: '/statistics/richlist', component: RichList,
    },
  ]

})
