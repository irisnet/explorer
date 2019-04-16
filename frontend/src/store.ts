import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state:{
    flShowSelectOption: false,
    currentEnv:'',
    isMobile:false,
    tabIndex:0,
    flShowValidatorStatus: false
  },
  mutations:{
    flShowSelectOption(state,data){
      state.flShowSelectOption = data
    },
    currentEnv(state,data){
      state.currentEnv = data
    },
    isMobile(state,data){
      state.isMobile = data
    },
    tabIndex(state,data){
      state.tabIndex = data
    },
    flShowValidatorStatus(state,data){
      state.flShowValidatorStatus = data
    }
  }
})
export default store
