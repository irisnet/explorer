import axios from "axios"
export default class Service {
  static http(url){
   return axios.get(url).then(data => {
      if(data.status === 200){
        return data.data
      }
    }).then(data => {
      if(data && typeof data ==="object"){
        if(data.code && data.code === "EX-100001"){
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

