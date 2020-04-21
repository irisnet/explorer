import Vue from 'vue';
import Router from 'vue-router';
Vue.use(Router);
import Home from "./components/home/Home.vue"
const router = new Router({
  routes: [
    // {
    //   path: '/',
    //   redirect: '/',
    // },
    {
      path: '/',
      component: Home,
    },

    {
      path: "*",
      redirect: '/',
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
