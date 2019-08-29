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
				return str.split(".")[0] + '.' +  splitString.substr(0,splitNum)
			}else {
				return str.split(".")[0] + '.' + splitString.padEnd(4, "0")
			}
		}else {
			return str + '.0000'
		}
	}
	
	static decimalPlace(num,val){
		if(val){
			return Tools.toFixedformatNumber(num ,val);
		}else{
			if(/^\+?[1-9][0-9]*$/.test(num)){
				return Tools.formatPriceToFixed(num)
			}else {
				if(num){
					num = Tools.convertScientificNotation2Number(num);
					let str = String(num).split(".")[1];
					if(str.length > 2){
						return Tools.formatPriceToFixed(num ,2);
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
		return new BigNumber(num).toFixed(2);
	}
	static formatFeeToFixedNumber(num){
		return  Tools.toFixedformatNumber(Tools.formatNumber(num) ,4);
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
        let arr = value.split('.');
		let integer = arr[0];
		let decimals = arr[1];
		let formattedInteger = integer.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
		return decimals ? `${formattedInteger}.${decimals}` : `${formattedInteger}`;
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
		if(denom.toLowerCase() === "iris-atto" || denom.toLowerCase() === "iris"){
			return "IRIS"
		}else {
			return denom
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
		if (address && address.length > 11) {
			return `${address.substring(0,3)}...${address.substring(address.length - 8)}`
		}else if(address && address.length < 11){
			return address
		}
		return '';
	}
	/**
	 * format txHash
	 * param String
	 * return String
	 */
	static formatTxHash(txHash){
		return `${txHash.substring(0,3)}...${txHash.substring(txHash.length - 3)}`
    }
    static subStrings(value, afterPointLength) { //截取指定小数位长度字符串
        if (value) {
            let arr = value.split('.');
            arr[1] = arr[1] || '';
            value = `${arr[0]}.${arr[1].padEnd(afterPointLength, '0').substring(0, afterPointLength)}`;
        }
        return value;
    }
	static formatTxList(list,txType){
		if(list !== null){
			return list.map(item => {
				let [Amount,Fee,transferAmount,transferFee,tokenId] = ['--','--','--','--','--'];
				let commonHeaderObjList,objList,commonFooterObjList;
				if(item.amount){
					if(item.amount instanceof Array && item.amount.length > 0){
						if(item.amount[0].denom && item.amount[0].amount && item.amount[0].denom === Constant.Denom.IRISATTO){
							transferAmount = item.amount[0].formatAmount = item.amount[0].amount > 0 ? Tools.formatAmount(item.amount[0].amount) : item.amount[0].amount;
							tokenId = item.amount[0].tokenId = Constant.Denom.IRIS.toLocaleUpperCase();
							Amount = item.amount.map(listItem=>`${listItem.formatAmount} ${Tools.formatDenom(listItem.denom).toUpperCase()}`).join(',');
						}else if(item.amount[0].denom && item.amount[0].amount && item.amount[0].denom !== Constant.Denom.IRISATTO){
							transferAmount = item.amount[0].formatAmount = Tools.FormatScientificNotationToNumber(item.amount[0].amount);
							tokenId = item.amount[0].tokenId = item.amount[0].denom.toLocaleUpperCase();
						}else {
							transferAmount = item.amount[0].formatAmount = item.amount[0].amount;
							tokenId = item.amount[0].tokenId = item.amount[0].denom.toLocaleUpperCase();
							if(item.Type === 'BeginUnbonding' || item.Type === 'BeginRedelegate'){
								item.amount[0].formatAmount = item.amount[0].amount > 0 ? Tools.formatAmount(item.amount[0].amount) : item.amount[0].amount;
								Amount = item.amount.map(listItem => `${listItem.formatAmount} SHARES`).join(',');
							}
						}
					}else if(item.amount && Object.keys(item.amount).includes('amount') && Object.keys(item.amount).includes('denom')){
						if(item.amount.denom === Constant.Denom.IRISATTO){
							transferAmount = Tools.formatAmount(item.amount);
							tokenId = Constant.Denom.IRIS.toLocaleUpperCase();
							Amount = `${transferAmount}  ${Tools.formatDenom(item.amount.denom).toUpperCase()}`;
							
						}else if(!item.amount.denom){
							transferAmount = Tools.FormatScientificNotationToNumber(item.amount);
							tokenId = ''
						}else {
							transferAmount = item.amount;
							tokenId = item.denom.toLocaleUpperCase();
							if(item.Type === 'BeginUnbonding' || item.Type === 'BeginRedelegate'){
								Amount = item.amount.map(listItem => `${listItem.amount} SHARES`).join(',');
							}
						}
					}
				}
				if(item.fee.amount && item.fee.denom){
					let feeAmount = item.fee.amount;
					transferFee = `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(feeAmount)))}`;
					Fee = `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(feeAmount)),4)} ${Tools.formatDenom(item.fee.denom).toUpperCase()}`;
				}
				commonHeaderObjList = {
					'Tx_Hash' : item.hash,
					'Block' : item.block_height
				};
				commonFooterObjList = {
					'Tx_Type': item.type,
					'Tx_Fee': Fee,
					'Tx_Signer': item.signer ? item.signer : '',
					'Tx_Status': Tools.firstWordUpperCase(item.status),
					'Timestamp': Tools.format2UTC(item.timestamp),
				};
				if(txType === 'transfers' ){
					objList = {
						'From':item.from?item.from:(item.delegator_addr?item.delegator_addr:'--'),
						Amount:transferAmount,
						'transferFee': transferFee,
						tokenId,
						'To':item.to?item.to:(item.validator_addr?item.validator_addr:'--'),
						'listName':'transfer'
					};
				}else if(txType === 'validations'){
					let Moniker = item.moniker;
					objList = {
						'Moniker': item.moniker ? Tools.formatString(Moniker,15,"...") : "--",
						Amount,
						'OperatorAddr': item.operator_addr ? item.operator_addr : '--',
						'listName':'validations',
						'Self_Bonded': item.self_bonded
					}
				}else if(txType === 'delegations'){
					objList = {
						'From': `${item.from ? item.from : (item.delegator_addr?item.delegator_addr:'--')}`,
						Amount,
						'To': `${item.to ? item.to : (item.validator_addr?item.validator_addr:'--')}`,
						'listName':'delegations',
						fromMoniker: item.from_moniker,
						toMoniker: item.to_moniker
					}
				}else if(txType === 'governance'){
					objList = {
						'Proposal_Type': item.proposal_type ? item.proposal_type : '--',
						"Proposal_ID": item.proposal_id === 0 ? "--" : item.proposal_id,
						'Proposal_Title': item.title ?  Tools.formatString(item.title,15,"...") : '--',
						Amount,
						'Tx_Type': item.type,
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
				noObjList = [];
				// noObjList = [{
				// 	'Tx_Hash': '',
				// 	'Block': '',
				// 	'Proposal_Type': '',
				// 	"Proposal_ID": '',
				// 	'Proposal_Title': '',
				// 	'Amount': '',
				// 	'Tx_Type': '',
				// 	'Tx_Fee': '',
				// 	'Tx_Signer': '',
				// 	'Tx_Status': '',
				// 	'Timestamp':'',
				// 	'listName':'gov'
				// }];
			}
			return noObjList;
		}
	}
	
	static addressRoute(address) {
		if(address) {
			if (address.substring(0, 3) === this.$Crypto.config.iris.bech32.valAddr || address.substring(1, 3) === 'va') {
				return `/validators/${address}`;
			} else {
				return `/address/${address}`;
			}
		}
		return '';
	}
	/**
	 *   formatAmount
	 *   param Object or array
	 *   return string
	 * */
	static formatAmount2(param){
		let amount,amountDenom, amountNumber,amountRadixNumber;
		if(param instanceof Array){
			amount = param[0].amount;
			amountDenom = param[0].denom;
		}else if(param instanceof Object){
			amount = param.amount;
			amountDenom = param.denom;
		}
		
		if(amount.toString().indexOf('e') !== -1 || amount.toString().indexOf('E') !== -1){
			amountNumber = new BigNumber(amount).toFixed().toString() + ".";
			amountRadixNumber = Tools.amountRadix(amountDenom);
		}else {
			amountNumber = amount.toString() + ".";
			amountRadixNumber = Tools.amountRadix(amountDenom);
		}
		if(amountDenom){
			return `${new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber)* -1))).toFormat()} ${Constant.RADIXDENOM.IRIS.toLocaleUpperCase()}`
		}else {
			return `${new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber) * -1))).toFormat()} SHARES`
		}
	}
	/**
	*   Amount iris Radix
	 *   param string
	 *   return Radix length
	* */
	static amountRadix(param){
	let diffValue = 1,defaultValue = 18;
	if(param){
		//radix number length need to minus 1;
		switch (param) {
			case Constant.RADIXDENOM.IRISATTO:
				return Constant.RADIXDENOM.IRISATTONUMBER.length - diffValue;
				break;
			case Constant.RADIXDENOM.IRISMILLI:
				return Constant.RADIXDENOM.IRISMILLINUMBER.length - diffValue;
				break;
			case Constant.RADIXDENOM.IRISMICRO:
				return Constant.RADIXDENOM.IRISMICRONUMBER.length - diffValue;
				break;
			case Constant.RADIXDENOM.IRISNANO :
				return Constant.RADIXDENOM.IRISNANONUMBER.length - diffValue;
				break;
			case Constant.RADIXDENOM.IRISPICO :
				return Constant.RADIXDENOM.IRISPICONUMBER.length - diffValue;
				break;
			case Constant.RADIXDENOM.IRISFEMTO :
				return Constant.RADIXDENOM.IRISFEMTONUMBER.length - diffValue;
				break;
			case Constant.RADIXDENOM.IRIS :
				return Constant.RADIXDENOM.IRISNUMBER.length;
				break;
			default:
				return defaultValue;
			}
		}else {
			return defaultValue
		}
	}
	/**
	 *   科学计数法转成数字
	 *
	 * */
	static FormatScientificNotationToNumber(number){
		let formattedNumber;
		if(number.toString().indexOf('e') !== -1 || number.toString().indexOf('E') !== -1){
			formattedNumber = new BigNumber(number).toFixed().toString();
		}else {
			formattedNumber = number
		}
		return formattedNumber
	}
	
}
