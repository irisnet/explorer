/**
 * 工具类
 */
export default class Tools{

  /**
   * 判断当前是移动端还是pc端
   * param void;
   * return boolean
   */
  static currentDeviceIsPersonComputer(){
    const userAgentInfo = navigator.userAgent;
    const Agents = ["Android", "iPhone", "SymbianOS", "Windows Phone", "iPad", "iPod"];
    let flag = true;
    for (let i = 0; i < Agents.length; i++) {
      if (userAgentInfo.includes(Agents[i]) > 0) {
        flag = false;
        break;
      }
    }
    return flag;
  }

  /**
   * 后端返回的数据转换成标准格式
   * param string;
   * return string
   */
  static conversionTimeToUTC(originTime){
    return `${originTime.substr(5,2)}/${originTime.substr(8,2)}/${originTime.substr(0,4)} ${originTime.substr(11,8)}+UTC`;
  }
  static conversionTimeToUTCToYYMMDD(originTime){
    return `${originTime.substr(0,4)}/${originTime.substr(5,2)}/${originTime.substr(8,2)} ${originTime.substr(11,8)}+UTC`;
  }
  static conversionTimeToUTCByValidatorsLine(originTime){
    return `${originTime.substr(0,4)}/${originTime.substr(5,2)}/${originTime.substr(8,2)} ${originTime.substr(11,8)}`;
  }

  static formatNumber(num){
    return num/1000000000000000000;
  }


  static decimalPlace(num,val){
    if(val){
      return (parseInt(String(num*100000000))/100000000).toFixed(8);
    }else{
      if(/^\+?[1-9][0-9]*$/.test(num)){
        return num + " "
      }else {
        if(num){
          num = Tools.scientificToNumber(num);
          let str = String(num).split(".")[1];
          if(str.length > 2){
            return (parseInt(String(num*100))/100).toFixed(2) + "...";
          }else {
            return (parseInt(String(num*100))/100)
          }
        }
      }
    }

  }
  static scientificToNumber(num){
    //处理非数字
    if(isNaN(num)){return num}
    //处理不需要转换的数字
    let str = ''+num;
    if(!/e/i.test(str)){return num;}
    return (num).toFixed(18).replace(/\.?0+$/, "");
  }

  static formatNumberToFixedNumber(num){
    return (parseInt(String(num*10000))/10000).toFixed(4);
  }
  static formatFeeToFixedNumber(num){
    return (Tools.scientificToNumber(parseInt(String(Tools.formatNumber(num)*10000))/10000)).toFixed(4) + "...";
  }
  /**
   * 格式化年月日
   * param string;
   * return string
   */
  static formatDateYearToDate(timestamp) {
    var date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '/';
    var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '/';
    var D = date.getDate() + ' ';
    var h = date.getHours() + ':';
    var m = date.getMinutes() + ':';
    var s = date.getSeconds();
    return Y+M+D;
  }
  /**
   * 格式化年月日时分秒
   * param string;
   * return string
   */
  static formatDateYearAndMinutesAndSeconds(timestamp) {
    var date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '/';
    var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '/';
    var D = date.getDate() < 10 ? '0' + (date.getDate()) + " " : (date.getDate()) + ' ';
    var h = date.getHours() < 10 ? '0' + (date.getHours()) : + (date.getHours()) ;
    var m = date.getMinutes() < 10 ? '0' + (date.getMinutes()) : + (date.getMinutes());
    var s = date.getSeconds() < 10 ? '0' + (date.getSeconds()) : + (date.getSeconds()) ;
    return Y + M + D + h + ':'+ m + ':'+s;
  }/**
   * 格式化amount
   * param string;
   * return string
   */
  static dealWithFees(num){
    return Tools.scientificToNumber(Tools.decimalPlace(Tools.formatNumber(num)))
  }
  /**
   * 根据字节截取字符串
   */
  static getShortForm(string,cutOutlength,addSuffix){
    var stringLength = string.replace(/[^\x00-\xff]/g,"**").length;
    if(stringLength>cutOutlength){
      if(!addSuffix) {
        addSuffix="......";
      }
      var i=0;
      for(var z=0;z<cutOutlength;z++){
        if(string.charCodeAt(z)>255){
          i=i+2;
        }else{
          i=i+1;
        }
        if(i>=cutOutlength){
          string=string.slice(0,(z + 1))+addSuffix;
          break;
        }
      }
      return string;
    }else{
      return string+"";
    }
  }

  /**
   * 格式化空格
   */
  static removeAllSpace(str) {
    return str.replace(/\s+/g, "");
  }

  static formatBalance(number, places, symbol, thousand, decimal) {
    number = number || 0;
    places = !isNaN(places = Math.abs(places)) ? places : 2;
    symbol = symbol !== undefined ? symbol : "";
    thousand = thousand || ",";
    decimal = decimal || "";
    var negative = number < 0 ? "-" : "",
      i = parseInt(number = Math.abs(+number || 0).toFixed(places), 10) + "",
      j = (j = i.length) > 3 ? j % 3 : 0;
    return symbol + negative + (j ? i.substr(0, j) + thousand : "") + i.substr(j).replace(/(\d{3})(?=\d)/g, "$1" + thousand) + (places ? decimal : "");
  }
}
