import Service from '../util/axios'
import URlApi from "../api"
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
        url = URlApi[Object.keys(params)[0]]
      }else {
        for(let key in params[Object.keys(params)[0]]){
          let rule =`{${key}}`;
          URlApi[Object.keys(params)[0]] = URlApi[Object.keys(params)[0]].replace(new RegExp(rule,"g"),params[Object.keys(params)[0]][key]);
          url = URlApi[Object.keys(params)[0]]
        }
      }
      Service.http(url).then( res => {
        callback(res);
      }).catch(err => {
        callback({error:err})
      })
    },
};
export default Server
