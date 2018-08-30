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

  static formatNumber(num){
    return num/1000000000000000000;
  }

  static scientificToNumber(num) {
    //将科学计数法转换成字符串
    var m = num.toExponential().match(/\d(?:\.(\d*))?e([+-]\d+)/);
    return num.toFixed(Math.max(0, (m[1] || '').length - m[2]));
  }
}
