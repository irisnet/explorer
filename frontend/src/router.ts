import Vue from 'vue';
import Router from 'vue-router';
Vue.use(Router);
import Home from "./components/home/Home.vue"
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
      component: () => import('@/components/utilpage/FaucetPage.vue'),
    },
    {
      path: '/gov/proposals', component: () => import('@/components/gov/ProposalsPage.vue')
    },
    {
      path: '/gov/parameters', component: () => import('@/components/gov/ParametersPage.vue')
    },
    {
      path: '/blocks', component: () => import('@/components/block/BlockListPage.vue')
    },
    {
      path: '/txs/:txType', component: () => import('@/components/txs/TxListPage.vue'),
    },
    {
      path: '/validators', component: () => import('@/components/address/validator/ValidatorListPage.vue')
    },
    {
      path: '/tx', component: () => import('@/components/txs/TxDetail.vue')
    },
    {
      path: '/block/:height', component: () => import('@/components/block/BlockInfo.vue'),
    },
    {
      path: '/ProposalsDetail/:proposal_id', component: () => import('@/components/gov/ProposalsDetail.vue')
    },
    {
      path: '/address/:param',
      component: () => import('@/components/address/delegator/AddressInfomation.vue'),
    },
    {
      path: '/validators/:param',
      component: () => import('@/components/address/validator/ValidatorDetail.vue'),
    },
    {
      path: '/privacy_policy',
      component: () => import('@/components/utilpage/PrivacyPolicy.vue'),
    },
    {
      path: '/searchResult/:searchContent', component: () => import('@/components/utilpage/searchResult.vue'),
    },
    {
      path: '/searchResult', component: () => import('@/components/utilpage/searchResult.vue'),
    },
    {
      path: '/version', component: () => import('@/components/utilpage/version.vue'),
    },
    {
      path: '/help', component: () => import('@/components/utilpage/clearStorageHelpPage.vue'),
    },
    {
      path: '/stats/irisrichlist', component: () => import('@/components/stats/RichList.vue'),
    },
    {
      path: '/stats/irisstats', component: () => import('@/components/stats/TokenStats.vue'),
    },
    // {
    //   path: '/statistics/bondedTokens', component: BondedTokens,
    // }
    {
      path: '/assets/ntvassetstxs', component: () => import('@/components/assets/NativeAsset.vue'),
    },
    {
      path: '/assets/gtwassetstxs', component: () => import('@/components/assets/GateWayAsset.vue'),
    },
    {
      path: '/assets/ntvassets', component: () => import('@/components/assets/NativeAssetList.vue'),
    },
    {
      path: '/assets/gtwassets', component: () => import('@/components/assets/GatewayAssetList.vue'),
    },
   /* {
      path: '/asset/:assetType/:tokenName', component: AssetInfo,
    },*/
    {
      path: '/asset/:assetType', component: () => import('@/components/assets/AssetInformation.vue'),
    },
    {
      path: '/error', component: () => import('@/components/utilpage/ErrorPage.vue'),
    },
    {
      path: '/txs', component: () => import('@/components/txs/AllTxTypeList.vue'),
    },
    {
      path: '/htlc/:txHash', component: () => import('@/components/htlc/HtlcInformation.vue'),
    },

    {
      path: "*",
      component: () => import('@/components/utilpage/ErrorPage.vue'),
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
