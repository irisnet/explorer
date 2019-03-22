import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state:{
    flShowSelectOption: false
  },
  mutations:{
    flShowSelectOption(state,data){
      state.flShowSelectOption = data
    }
  }
})
export default store
