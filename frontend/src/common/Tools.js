/**
 * 工具类
 */
import BigNumber from 'bignumber.js';
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
    return new BigNumber(num).div(1000000000000000000).toNumber();
  }

  static decimalPlace(num,val){
    if(val){
      return Tools.toFixedformatNumber(num ,val);
    }else{
      if(/^\+?[1-9][0-9]*$/.test(num)){
        return num + " "
      }else {
        if(num){
          num = Tools.scientificToNumber(num);
          let str = String(num).split(".")[1];
          if(str.length > 2){
            return Tools.toFixedformatNumber(Number(num) ,2)+ "...";
          }else {
            return (parseInt(String(num*100))/100)
          }
        }
      }
    }

  }
  static  toFixedformatNumber(num,val){
    return new BigNumber(num).toFixed(val,1);

  }
  static scientificToNumber(num){
    return new BigNumber(num).toFixed();
  }
  static formatFeeToFixedNumber(num){
    return  Tools.toFixedformatNumber(Tools.formatNumber(num) ,4) + "...";
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
  static formatAmount(num){
    return Tools.decimalPlace(Tools.formatNumber(num))
  }
  /**
   * 根据字节截取字符串
   */
  static formatString(string,cutOutlength,addSuffix){
    var stringLength = string.replace(/[^\x00-\xff]/g,"**").length;
    if(stringLength>cutOutlength){
      if(!addSuffix) {
        addSuffix="......";
      }
      var bytesLength = 0;
      var unitStringUnicodeMaxlength = 255;
      for(var index = 0;index < cutOutlength;index++){
        if(string.charCodeAt(index) > unitStringUnicodeMaxlength){
          bytesLength = bytesLength + 2;
        }else{
          bytesLength = bytesLength + 1;
        }
        if(bytesLength >= cutOutlength){
          string=string.slice(0,(index + 1))+addSuffix;
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
  static formatDenom(denom){
    if(denom === "iris-atto" || denom === "iris"){
      return "IRIS"
    }
  }
  static formatAccountCoinsAmount(coinsAmount){
    return coinsAmount = /[0-9]+[.]?[0-9]*/.exec(coinsAmount)
  }
  static formatAccountCoinsDenom(coinsDenom){
    return coinsDenom = /[A-Za-z\-]{2,15}/.exec(coinsDenom)
  }
}
