/**
 * 工具类
 */
import BigNumber from 'bignumber.js';
import moveDecimal  from "move-decimal-point"
import Constant from "../constant/Constant"
export default class Tools{
  /**
   * 根据展示的需求拼接字符串展示成 > xxdxxhxxmxxs ago 或者 xxdxxhxxmxxs ago 或者 xxdxxhxxmxxs
   * param prefix string;
   * param suffix string;
   * return string
   * */
  static formatAge(currentServerTime,time,suffix,prefix){
    let dateBegin = new Date(time);
    let dateDiff = currentServerTime - dateBegin.getTime();
    let dayDiff = Math.floor(dateDiff / (24 * 3600 * 1000));
    let hourLevel = dateDiff % (24 * 3600 * 1000);
    let hours = Math.floor(hourLevel / (3600 * 1000));
    let minuteLevel = hourLevel % (3600 * 1000);
    let minutes = Math.floor(minuteLevel / (60 * 1000));
    let secondLevel = minuteLevel % (60 * 1000) ;
    let seconds = Math.round(secondLevel / 1000);

    let str = `${dayDiff?`${dayDiff}d`:''} ${hours ? `${hours}h` : ''} ${dayDiff && hours ? '' : (minutes ? `${minutes}m`:'')} ${dayDiff || hours || minutes? '' : (seconds ? `${seconds}s`:'')}`;
    if(prefix && suffix){
      return`${prefix} ${str} ${suffix}`
    }else if(suffix){
      return`${str} ${suffix}`
    }else {
      return`${str}`
    }
  }
  static getDiffMilliseconds(currentServerTime,time){
    let dateBegin = new Date(time);
    let dateDiff = currentServerTime - dateBegin.getTime();
    return dateDiff
  }
  static formatDuring(ms) {
    let s = ms/1000;
    let days = (s / (60 * 60 * 24));
    let hours = ((s % (60 * 60 * 24)) / (60 * 60));
    let minutes = ((s % (60 * 60)) / (60));
    let seconds = (s % 60);
    return {
      days,hours,minutes,seconds
    }
  }
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
  static format2UTC(originTime){
    return `${originTime.substr(0,4)}/${originTime.substr(5,2)}/${originTime.substr(8,2)} ${originTime.substr(11,8)}+UTC`;
  }

  static conversionTimeToUTCByValidatorsLine(originTime){
    return `${originTime.substr(0,4)}/${originTime.substr(5,2)}/${originTime.substr(8,2)} ${originTime.substr(11,8)}`;
  }
  static formatNumber(num){
    return new BigNumber(num).div(1000000000000000000).toNumber();
  }
  static formatRate(rate){
    let toFixedValue = 2;
    let rateNum = new BigNumber(rate).multipliedBy(100).toNumber();
    if(rateNum.toString().indexOf(".") !== -1 && rateNum.toString().split('.')[1].length > 2){
      return rateNum
    }else {
      return Tools.toFixedformatNumber(rateNum,toFixedValue)
    }
  }
  static formaNumberAboutGasPrice(num){
    return new BigNumber(num).div(1000000000).toNumber();
  }
  /**
   * 格式化数字类型是string的数字并让小数点左移18位
   * param string or number;
   * return string
   */
  static numberMoveDecimal(number){
    let leftLength = -18;
    if(number.toString().indexOf('e') !== -1 || number.toString().indexOf('E') !== -1){
      return moveDecimal(new BigNumber(number).toFixed().toString() + ".",leftLength)
    }else {
      return moveDecimal(number.toString() + ".",leftLength)
    }
  }
  static formatStringToNumber(number){
    if(number.toString().indexOf('e') !== -1 || number.toString().indexOf('E') !== -1){
      number = new BigNumber(number).toFixed().toString();
    }else {
      number = number.toString()
    }
    let stringLength = number.length;
    let splitSite = 18;
    let stringSplitSiteLength = 19;
    let completeNumberString;
    if(stringLength >= stringSplitSiteLength){
      let integePartLength = stringLength - splitSite;
      let integePart = number.substr(0,integePartLength);
      let fractionalPart = number.substr(integePartLength,stringLength);
          completeNumberString = integePart.concat(".",fractionalPart);
      return  Tools.formatContinuousNumberZero(completeNumberString).split(".")[1] == "" ?
              Tools.formatContinuousNumberZero(completeNumberString).split(".")[0] :
              Tools.formatContinuousNumberZero(completeNumberString);
    }else {
        let integePartLength = splitSite - stringLength;
        let srtingNumArray = number.split("");
        for(let j = 0; j < integePartLength; j++){
          srtingNumArray.unshift("0")
        }
        completeNumberString = "0." + srtingNumArray.join("");
        return Tools.formatContinuousNumberZero(completeNumberString)
    }
  }
  static formatToken (token) {
    let coin = {};
    let amount = '';
    if(token.amount && token.amount !== '' && token.amount !== 0 && token.denom === Constant.Denom.IRISATTO){
      amount  = Tools.convertScientificNotation2Number(Tools.formatNumber(token.amount));
    }else if(token.denom === Constant.Denom.IRIS){
      amount  = token.amount;
    }
    coin.amount = amount;
    coin.denom = Constant.Denom.IRIS.toUpperCase();
    return coin
  }
  /**
   * 去除数字的类型是string的尾部连续为 0 的数字
   * param string;
   * return string
   */
  static formatContinuousNumberZero(str){
    let i;
    for(i = str.length - 1;i >= 0;i--) {
      if(str.charAt(i) != "0")break;
    }
    //TODO(zhangjinbiao)判断逻辑后面要更改
    if(str.substring(0,i+1) === '0.'){
      return '0'
    }else {
      return str.substring(0,i+1);
    }
  }
  static formatPriceToFixed(value,fixedValue){
    return new BigNumber(value).toFormat(fixedValue)
  }
  /**
   * 格式化数字的类型是string的数字并在小数点后面超过多少位以后加 ...
   * param string;
   * return string
   */
  static formatStringToFixedNumber(str,splitNum){
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
        return Tools.formatPriceToFixed(num) + " "
      }else {
        if(num){
          num = Tools.convertScientificNotation2Number(num);
          let str = String(num).split(".")[1];
          if(str.length > 2){
            return Tools.formatPriceToFixed(num ,2)+ "...";
          }else {
            return Tools.formatPriceToFixed(num)
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
  static convertScientificNotation3Number(num){
    return new BigNumber(num).toFixed(6);
  }
  static formatFeeToFixedNumber(num){
    return  Tools.toFixedformatNumber(Tools.formatNumber(num) ,4) + "...";
  }
  static formatDecimalNumberToFixedNumber(num){
    if(typeof num === 'number' && !Object.is(num, NaN)) {
      if (num < 0.01 && num !== 0) {
        return num.toFixed(4);
      } else {
        return num.toFixed(2);
      }
    }
    return num;
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
    let stringLength = string.replace(/[^\x00-\xff]/g,"**").length;
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
  /**
   * 格式化货币价格
   */
  static formatPrice(value) {
    let integer = value.split('.')[0];
    let decimals = value.split('.')[1];
    let formattedInteger = integer.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    return`${formattedInteger}.${decimals}`
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
  /**
   * 获取水龙头Amount
   * param ['11.1111iris']
   * return ['11.1111']
   */
  static formatAccountCoinsAmount(coinsAmount){
    return coinsAmount = /[0-9]+[.]?[0-9]*/.exec(coinsAmount)
  }
  /**
   * 获取水龙头demon
   * param ['11.1111iris']
   * return ['iris']
   */
  static formatAccountCoinsDenom(coinsDenom){
    return coinsDenom = /[A-Za-z\-]{2,15}/.exec(coinsDenom)
  }
  static scrollToTop(){
    document.documentElement.scrollTop = 0;
  }

  static firstWordUpperCase (str){
    return str.toLowerCase().replace(/(\s|^)[a-z]/g, function(char){
      return char.toUpperCase();
    });
  }
    static firstWordLowerCase (str){
        return str.toLowerCase().replace(/(\s|^)[a-z]/g, function(char){
            return char.toLocaleLowerCase();
        });
    }
  /**
   * format address
   * param String
   * return String
   */
  static formatValidatorAddress(address){
    return `${address.substring(0,3)}...${address.substring(address.length - 8)}`
  }
  /**
   * format txHash
   * param String
   * return String
   */
  static formatTxHash(txHash){
    return `${txHash.substring(0,3)}...${txHash.substring(txHash.length - 3)}`
  }
  static formatTxList(list,txType){
    if(list !== null){
      return list.map(item => {
        let [Amount,Fee] = ['--','--'];
        let commonHeaderObjList,objList,commonFooterObjList;
          if(item.Amount){
            if(item.Amount instanceof Array){
              if(item.Amount.length > 0){
                item.Amount[0].formatAmount = Tools.formatAmount(item.Amount[0].amount);
                if(!item.Amount[0].denom){
                  Amount = item.Amount.map(listItem => `${listItem.formatAmount} SHARES`).join(',');
                }else {
                  Amount = item.Amount.map(listItem=>`${listItem.formatAmount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',');
                }
              }
            }else if(item.Amount && Object.keys(item.Amount).includes('amount') && Object.keys(item.Amount).includes('denom')){
              if(!item.Amount.denom){
                Amount = `${Tools.formatPriceToFixed(item.Amount.amount)} SHARES`;
              }else {
                Amount = `${Tools.formatPriceToFixed(item.Amount.amount)}  ${Tools.formatDenom(item.Amount.denom).toUpperCase()}`;
              }
            }
          }
          if(item.Fee.amount && item.Fee.denom){
            let feeAmount = item.Fee.amount;
            Fee = `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(feeAmount)),4)} ${Tools.formatDenom(item.Fee.denom).toUpperCase()}`;
          }
        commonHeaderObjList = {
          'Tx_Hash' : item.Hash,
          'Block' : item.BlockHeight
        };
        commonFooterObjList = {
          'Tx_Type': item.Type,
          'Tx_Fee': Fee,
          'Tx_Signer': item.Signer ? item.Signer : '',
          'Tx_Status': Tools.firstWordUpperCase(item.Status),
          'Timestamp': Tools.format2UTC(item.Timestamp),
        };
        if(txType === 'transfers' ){
          objList = {
            'From':item.From?item.From:(item.DelegatorAddr?item.DelegatorAddr:'--'),
            Amount,
            'To':item.To?item.To:(item.ValidatorAddr?item.ValidatorAddr:'--'),
            'listName':'transfer'
          };
        }else if(txType === 'declarations'){
          let Moniker = item.Moniker;
          objList = {
            'Moniker': item.Moniker ? Tools.formatString(Moniker,15,"...") : "--",
            Amount,
            'OperatorAddr': item.OperatorAddr ? item.OperatorAddr : '--',
            'listName':'declarations'
          }
        }else if(txType === 'stakes'){
          objList = {
            'From':item.From ? item.From : (item.DelegatorAddr?item.DelegatorAddr:'--'),
            Amount,
            'To':item.To ? item.To : (item.ValidatorAddr?item.ValidatorAddr:'--'),
            'listName':'stakes'
          }
        }else if(txType === 'governance'){
          objList = {
            'Proposal_Type': item.ProposalType ? item.ProposalType : '--',
            "Proposal_ID": item.ProposalId === 0 ? "--" : item.ProposalId,
            'Proposal_Title': item.Title ?  Tools.formatString(item.Title,15,"...") : '--',
            Amount,
            'Tx_Type': item.Type,
            'Tx_Fee': '',
            'listName':'gov'
          }
        }
        let allObjList = Object.assign(commonHeaderObjList,objList,commonFooterObjList);
        return allObjList;
      })
    }else {
      let noObjList;
      if(txType === 'transfers'){
        noObjList = [{
          'Tx_Hash': '',
          'Block':'',
          'From':'',
          'Amount':'',
          'To':'',
          'Tx_Type':'',
          'Tx_Fee':'',
          'Tx_Signer':'',
          'Tx_Status': '',
          'Timestamp':'',
          'listName':'transfer'
        }];
      }else if(txType === 'declarations'){
        noObjList = [{
          'Tx_Hash': '',
          'Block':'',
          'Moniker':'',
		  'OperatorAddr':'',
          'From':'',
          'Amount':'',
          "To":'',
          'Tx_Type':'',
          'Tx_Fee':'',
          'Tx_Signer':'',
          'Tx_Status': '',
          'Timestamp':'',
          'listName':'declarations'
        }];
      }else if(txType === 'stakes'){
        noObjList = [{
          'TxHash': '',
          'Block':'',
          'From':'',
          'Amount':'',
          'To':'',
          'Tx_Type':'',
          'Tx_Fee':'',
          'Tx_Signer': '',
          'Tx_Status': '',
          'Timestamp':'',
          'listName':'stakes'
        }];
      }else if(txType === 'governance'){
        noObjList = [{
          'Tx_Hash': '',
          'Block': '',
          'Proposal_Type': '',
          "Proposal_ID": '',
          'Proposal_Title': '',
          'Amount': '',
          'Tx_Type': '',
          'Tx_Fee': '',
          'Tx_Signer': '',
          'Tx_Status': '',
          'Timestamp':'',
          'listName':'gov'
        }];
      }
      return noObjList;
    }
  }

  static addressRoute(address) {
    if (address.substring(0, 3) === this.$Crypto.config.iris.bech32.valAddr || address.substring(1, 3) === 'va') {
      return `/validators/${address}`;
    } else {
      return `/address/${address}`;
    }
  }
}
