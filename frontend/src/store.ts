import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
import {NATIVE_TOKEN} from "@/config";

const store = new Vuex.Store({
  state:{
    flShowSelectOption: false,
    currentEnv:'',
    isMobile:false,
    validatorTabIndex:0,
    flShowIpt:false,
    currentParticipationValue:"",
    currentYesValue:'',
    currentNoWithVetoValue:'',
    currentAbstainValue: '',
    currentNoValue: '',
    currentSkinStyle:'bifrost',
    flShowLoading: false,
    flShowQR:false,
    setQrImg:'',
    testSkinStyle:false,
    hideTestSkinStyle: true,
    nativeToken: NATIVE_TOKEN,
    showHeaderUnfoldBtn:sessionStorage.getItem('flShowHeaderUnfoldBtn') ? sessionStorage.getItem('flShowHeaderUnfoldBtn') : 'show',
    flShowSearchBar: sessionStorage.getItem('flShowSearchBar') ? sessionStorage.getItem('flShowSearchBar') : 'hide',
    flShowSearchIpt: sessionStorage.getItem('flShowSearchIpt') ? sessionStorage.getItem('flShowSearchIpt') : 'hide',
  },
  mutations:{
    currentNoValue(state,data){
      state.currentNoValue = data
    },
    currentAbstainValue(state,data){
      state.currentAbstainValue = data
    },
    currentNoWithVetoValue(state,data){
      state.currentNoWithVetoValue = data
    },
    currentParticipationValue(state,data){
      state.currentParticipationValue = data
    },
    currentYesValue(state,data){
      state.currentYesValue = data
    },
    flShowSelectOption(state,data){
      state.flShowSelectOption = data
    },
    currentEnv(state,data){
      state.currentEnv = data
    },
    isMobile(state,data){
      state.isMobile = data
    },
    validatorTabIndex(state,data){
      state.validatorTabIndex = data
    },
    flShowIpt(state,data){
      state.flShowIpt = data
    },
    currentSkinStyle(state,data){
      state.currentSkinStyle = data
    },
    flShowLoading(state,data){
      state.flShowLoading = data
    },
    flShowQR(state,data){
      state.flShowQR = data
    },
    setQrImg(state,data){
      state.setQrImg = data
    },
    testSkinStyle(state,data){
      state.testSkinStyle = data
    },
    hideTestSkinStyle(state,data){
      state.hideTestSkinStyle = data
    },
    showHeaderUnfoldBtn(state,data){
      state.showHeaderUnfoldBtn = data
    },
    flShowSearchBar(state,data){
      state.flShowSearchBar = data
    },
    flShowSearchIpt(state,data){
      state.flShowSearchIpt = data
    }
  }
})
export default store
