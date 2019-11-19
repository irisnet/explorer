import Service from '../util/axios'
import urlApi from "../api"
import Store from "../store"
const Server = {
    /**
    * params
    * {
    *     apiUrlObjectKey:{
    *         urlParamsName: urlParamsValue
    *     }
    * }
    * */
  commonInterface(params,callback){
      let url ;
      if(JSON.stringify(params[Object.keys(params)[0]]) === '{}'){
        url = urlApi[Object.keys(params)[0]]
      }else {
        url = urlApi[Object.keys(params)[0]];
        for(let key in params[Object.keys(params)[0]]){
          let rule =`{${key}}`;
          url = url.replace(new RegExp(rule,"g"),params[Object.keys(params)[0]][key]);
        }
      }
        if(Object.keys(params)[0] === 'candidatesTop' ||
            Object.keys(params)[0] === 'txsByDay' ||
            Object.keys(params)[0] === 'blocksRecent' ||
            Object.keys(params)[0] === 'txsRecent' ||
            Object.keys(params)[0] === 'navigation' ||
            Object.keys(params)[0] === 'homeProposalList'){
            Store.commit('flShowLoading',false)
        }else {
            Store.commit('flShowLoading',true)
        }
      Service.http(url).then( res => {
      Store.commit('flShowLoading',false)
          callback(res);
      }).catch(err => {
        callback({error:err})
      })
    },
};
export default Server
