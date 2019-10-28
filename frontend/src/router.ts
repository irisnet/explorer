import Vue from 'vue';
import Router from 'vue-router';
Vue.use(Router);
import Home from "./views/Home.vue"
const router = new Router({
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
      component: () => import('@/components/FaucetPage.vue'),
    },
    {
      path: '/gov/proposals', component: () => import('@/components/ProposalsPage.vue')
    },
    {
      path: '/gov/parameters', component: () => import('@/components/ParametersPage.vue')
    },
    {
      path: '/blocks', component: () => import('@/components/BlockListPage.vue')
    },
    {
      path: '/txs/:txType', component: () => import('@/components/TxListPage.vue'),
    },
    {
      path: '/validators', component: () => import('@/components/ValidatorListPage.vue')
    },
    {
      path: '/tx', component: () => import('@/components/TransactionsDetail.vue')
    },
    {
      path: '/block/:height', component: () => import('@/components/BlockInfo.vue'),
    },
    {
      path: '/ProposalsDetail/:proposal_id', component: () => import('@/components/ProposalsDetail.vue')
    },
    {
      path: '/address/:param',
      component: () => import('@/components/AddressInfomation.vue'),
    },
    {
      path: '/validators/:param',
      component: () => import('@/components/ValidatorDetail.vue'),
    },
    {
      path: '/privacy_policy',
      component: () => import('@/components/PrivacyPolicy.vue'),
    },
    {
      path: '/searchResult/:searchContent', component: () => import('@/components/searchResult.vue'),
    },
    {
      path: '/searchResult', component: () => import('@/components/searchResult.vue'),
    },
    {
      path: '/version', component: () => import('@/components/version.vue'),
    },
    {
      path: '/help', component: () => import('@/components/clearStorageHelpPage.vue'),
    },
    {
      path: '/stats/irisrichlist', component: () => import('@/components/RichList.vue'),
    },
    {
      path: '/stats/irisstats', component: () => import('@/components/TokenStats.vue'),
    },
    // {
    //   path: '/statistics/bondedTokens', component: BondedTokens,
    // }
    {
      path: '/assets/ntvassetstxs', component: () => import('@/components/NativeAsset.vue'),
    },
    {
      path: '/assets/gtwassetstxs', component: () => import('@/components/GateWayAsset.vue'),
    },
    {
      path: '/assets/ntvassets', component: () => import('@/components/NativeAssetList.vue'),
    },
    {
      path: '/assets/gtwassets', component: () => import('@/components/GatewayAssetList.vue'),
    },
   /* {
      path: '/asset/:assetType/:tokenName', component: AssetInfo,
    },*/
    {
      path: '/asset/:assetType', component: () => import('@/components/AssetInformation.vue'),
    },
    {
      path: '/error', component: () => import('@/components/ErrorPage.vue'),
    },
    {
      path: '/txs', component: () => import('@/components/AllTxTypeList.vue'),
    },
    {
      path: "*",
      component: () => import('@/components/ErrorPage.vue'),
    }
  ]

})
router.beforeEach((to, from, next) => {
    if (sessionStorage.getItem('Show_faucet') === '0') {
      if (to.path === '/faucet') {
        next('/')
      } else {
        next()
      }
    } else {
      next()
    }
})
export default router;
