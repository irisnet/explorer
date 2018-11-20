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
  static formaNumberAboutGasPrice(num){
    return new BigNumber(num).div(1000000000).toNumber();
  }
  /**
   * 格式化数字类型是string的数字并让小数点左移18位
   * param string;
   * return string
   */

  static formatNumberTypeOfString(stringnum){
    let stringLength = stringnum.length;
    let splitSite = 18;
    let stringSplitSiteLength = 19;
    let completeNumberString;
    if(stringLength >= stringSplitSiteLength){
      let integePartLength = stringLength - splitSite;
      let integePart = stringnum.substr(0,integePartLength);
      let fractionalPart = stringnum.substr(integePartLength,stringLength);
          completeNumberString = integePart.concat(".",fractionalPart);
      return  Tools.formatContinuousNumberZero(completeNumberString).split(".")[1] == "" ?
              Tools.formatContinuousNumberZero(completeNumberString).split(".")[0] :
              Tools.formatContinuousNumberZero(completeNumberString);
    }else {
        let integePartLength = splitSite - stringLength;
        let srtingNumArray = stringnum.split("");
        for(let j = 0; j < integePartLength; j++){
          srtingNumArray.unshift("0")
        }
        completeNumberString = "0." + srtingNumArray.join("");
        return Tools.formatContinuousNumberZero(completeNumberString)
    }
  }
  /**
   * 去除数字的类型是string的尾部连续为 0 的数字
   * param string;
   * return string
   */
  static formatContinuousNumberZero(str){
    let i;
    for(i = str.length - 1;i >= 0;i--) {
      if(str.charAt(i) != "0" && str.charAt(i) != "0")break;
    }
    return str.substring(0,i+1);
  }
  /**
   * 格式化数字的类型是string的数字并展示小数点后面超过多少位加 ...
   * param string;
   * return string
   */
  static formatNumberTypeOfStringToFixed(str,splitNum){
    if(str.indexOf(".") !== -1) {
      let splitString = str.split(".")[1];
      if(splitString.length > splitNum){
        return str.split(".")[0] + '.' +  splitString.substr(0,splitNum) + "..."
      }else {
        return str.split(".")[0] + '.' + splitString
      }
    }else {
      return str
    }
  }





  static decimalPlace(num,val){
    if(val){
      return Tools.toFixedformatNumber(num ,val);
    }else{
      if(/^\+?[1-9][0-9]*$/.test(num)){
        return num + " "
      }else {
        if(num){
          num = Tools.convertScientificNotation2Number(num);
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
  static convertScientificNotation2Number(num){
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
  static flTxType(TxType){
    if(TxType === "CompleteUnbonding" || TxType === "BeginUnbonding" || TxType === "BeginRedelegate" || TxType === "CompleteRedelegation" ){
      return true
    }
  }
  static scrollToTop(){
    document.body.scrollTop = 0;
  }

  static commonTxListItem(list,txType){
    if(list !== null){
      return list.map(item => {
        let [Amount,Fee] = ['--','--'];
        let commonHeaderObjList,objList,commonFooterObjList;
        if(txType === 'Transfers' || txType === 'Stakes' || txType === 'governance'){
          if(item.Amount){
            if(item.Amount instanceof Array){
              if(item.Amount.length > 0){
                item.Amount[0].amount = Tools.formatAmount(item.Amount[0].amount);
                if(Tools.flTxType(item.Type)){
                  Amount = item.Amount.map(listItem => `${listItem.amount} SHARES`).join(',');
                }else {
                  Amount = item.Amount.map(listItem=>`${listItem.amount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',');
                }
              }
            }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
              Amount = `${item.Amount.amount}  ${Tools.formatDenom(item.Amount.denom).toUpperCase()}`;
              if(Tools.flTxType(item.Type)){
                Amount = `${item.Amount.amount} SHARES`;
              }
            }
          }
          if(item.Fee.amount && item.Fee.denom){
            Fee = item.Fee.amount = `${Tools.formatFeeToFixedNumber(item.Fee.amount)} ${Tools.formatDenom(item.Fee.denom).toUpperCase()}`;
          }
        }
        commonHeaderObjList = {
          TxHash : item.Hash,
          Block : item.BlockHeight
        };
        commonFooterObjList = {
          Status : item.Status,
          Timestamp: Tools.conversionTimeToUTCToYYMMDD(item.Timestamp)
        };
        if(txType === 'Transfers' ){
          objList = {
            From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
            To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
            Amount,
            Fee,
          };
        }else if(txType === 'Declarations'){
          objList = {
            Owner: item.Owner ? item.Owner : "--",
            Moniker: item.Moniker ? Tools.formatString(item.Moniker,20,"...") : "--",
            "Self-Bond": item.SelfBond && item.SelfBond.length > 0 ? `${Tools.formatAmount(item.SelfBond[0].amount)} ${Tools.formatDenom(item.SelfBond[0].denom).toUpperCase()}` : "--",
            Type: item.Type,
            Fee: `${Tools.formatFeeToFixedNumber(item.Fee.amount)} ${Tools.formatDenom(item.Fee.denom).toUpperCase()}`,
          }
        }else if(txType === 'Stakes'){
          objList = {
            TxHash: item.Hash,
            Block:item.BlockHeight,
            From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
            To:item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:''),
            Type:item.Type === 'coin'?'transfer':item.Type,
            Amount,
            Fee,
          }
        }else if(txType === 'Governance'){
          objList = {
            From:item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:''),
            "Proposal_ID": item.ProposalId === 0 ? "--" : item.ProposalId,
            Type: item.Type,
            Fee,
          }
        }
        let allObjList = Object.assign(commonHeaderObjList,objList,commonFooterObjList);
        return allObjList;
      })
    }else {
      let noObjList;
      if(txType === 'Transfers'){
        noObjList = [{
          TxHash: '',
          Block:'',
          From:'',
          To:'',
          Amount:'',
          Fee:'',
          Status: "",
          Timestamp:'',
        }];
      }else if(txType === 'Declarations'){
        noObjList = [{
          TxHash: '',
          Block:'',
          Owner:'',
          Moniker:'',
          "Self-Bond":'',
          Type:'',
          Fee:'',
          Status: "",
          Timestamp:'',
        }];
      }else if(txType === 'Stakes'){
        noObjList = [{
          TxHash: '',
          Block:'',
          From:'',
          To:'',
          Type:'',
          Amount:'',
          Fee:'',
          Status: "",
          Timestamp:'',
        }];
      }else if(txType === 'Governance'){
        noObjList = [{
          TxHash: '',
          Block:'',
          From:'',
          "Proposal_ID":'',
          Type:'',
          Fee:'',
          Status: "",
          Timestamp:'',
        }];
      }
      return noObjList;
    }

  }
}
