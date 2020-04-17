import Vue from 'vue'
import Vuex from 'vuex'
import { mapState } from 'vuex'
Vue.use(Vuex)
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
    currentSkinStyle:'default',
    flShowLoading: false,
    flShowQR:false,
    setQrImg:'',
    testSkinStyle:false,
    hideTestSkinStyle: true,
    showHeaderUnfoldBtn:sessionStorage.getItem('flShowHeaderUnfoldBtn') ? sessionStorage.getItem('flShowHeaderUnfoldBtn') : 'hide',
    flShowSearchBar: sessionStorage.getItem('flShowSearchBar') ? sessionStorage.getItem('flShowSearchBar') : 'show',
    flShowSearchIpt: sessionStorage.getItem('flShowSearchIpt') ? sessionStorage.getItem('flShowSearchIpt') : 'show',
  },
  mutations:{
   /* ...mapState({
      currentNoValue:"currentNoValue",
      currentAbstainValue:"currentAbstainValue",
      currentNoWithVetoValue:"currentNoWithVetoValue",
      currentParticipationValue:"currentParticipationValue",
      currentYesValue:"currentYesValue",
      flShowSelectOption:"flShowSelectOption",
      currentEnv:"currentEnv",
      isMobile:"isMobile",
      validatorTabIndex:"validatorTabIndex",
      flShowIpt:"flShowIpt",
      currentSkinStyle:"currentSkinStyle",
      flShowLoading:"flShowLoading",
      flShowQR:"flShowQR",
      setQrImg:"setQrImg",
      testSkinStyle:"testSkinStyle",
      hideTestSkinStyle:"hideTestSkinStyle",
      showHeaderUnfoldBtn:"showHeaderUnfoldBtn",
      flShowSearchBar:"flShowSearchBar",
      flShowSearchIpt:"flShowSearchIpt",
    }),*/
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
