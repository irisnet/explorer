import axios from "axios"
export default class Service {
  static http(url){
   return axios.get(url).then(data => {
      if(data.status === 200){
        return data.data
      }
    }).then(data => {
      if(data && typeof data ==="object"){
        let  errorCodeArray =["EX-100000","EX-100001","EX-100002","EX-100003","EX-100004","EX-100005","EX-100006","EX-100007","EX-100008"];
        if(data.code &&errorCodeArray.some((item)=> {
          return item === data.code
        })){
          return null
        }else {
          return data
        }
      }else {
        return null;
      }
    });
  }
}

