/**
 * 工具类
 */
import BigNumber from 'bignumber.js';
import moveDecimal  from "move-decimal-point"
import Constant from "../constant/Constant"
import bech32 from 'bech32';
import moment from 'moment';
import store from "../store";
export default class Tools {
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
	/**
	 * 获取时间差值
	 * */
	static getDiffMilliseconds(currentServerTime,time){
		let dateBegin = new Date(time);
		let dateDiff = currentServerTime - dateBegin.getTime();
		return dateDiff
	}
	/**
	 * 根据毫秒数算出相差的天 小时 分钟 秒
	 * */
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
	 * 后端返回的数据转换成标准格式 (+UTC)
	 */
	static format2UTC(originTime){
		return `${originTime.substr(0,4)}/${originTime.substr(5,2)}/${originTime.substr(8,2)} ${originTime.substr(11,8)}+UTC`;
	}
	/**
	 * 后端返回的数据转换成标准格式 ()
	 */
	static conversionTimeToUTCByValidatorsLine(originTime){
		return `${originTime.substr(0,4)}/${originTime.substr(5,2)}/${originTime.substr(8,2)} ${originTime.substr(11,8)}`;
	}
	/**
	 * 格式化数字（除以10的18次方）
	 * */
	static formatNumber(num){
		return new BigNumber(num).div(1000000000000000000).toNumber();
	}
	/**
	 * 格式化数字并保留两位数(可以废除)
	 *
	 * */
	static formatRate(rate){
		let toFixedValue = 2;
		let rateNum = new BigNumber(rate).multipliedBy(100).toNumber();
		if(rateNum.toString().indexOf(".") !== -1 && rateNum.toString().split('.')[1].length > 2){
			return rateNum
		}else {
			return Tools.toFixedformatNumber(rateNum,toFixedValue)
		}
	}
	/**
	 * 格式化数字（除以10的18次方）
	 * */
	static formaNumberAboutGasPrice(num){
		return new BigNumber(num).div(1000000000).toNumber();
	}
	/**
	 * 格式化数字类型是string的数字并让小数点左移18位 (本质是移动小数点的位置)
	 *
	 */
	static numberMoveDecimal(number){
		let leftLength = -18;
		if(number.toString().indexOf('e') !== -1 || number.toString().indexOf('E') !== -1){
		    if(number.toString().indexOf('e')!== -1) {
                return moveDecimal(new BigNumber(number).toFixed().toString(),leftLength)

            }else {
                return moveDecimal(new BigNumber(number).toFixed().toString() + ".",leftLength)
            }
		}else {
		    if(number.toString().indexOf('e')!== -1){
                return moveDecimal(number.toString(),leftLength)

            }else {
                return moveDecimal(number.toString() + ".",leftLength)
            }
		}
	}

	/**
	 * 格式化数字（可废除）
	 */

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

	/**
	 *处理Token(可废除)
	 */
	static formatToken (token) {
		let coin = {};
		let amount = '';
		if(token.amount && token.amount !== '' && token.amount !== 0 && token.denom === store.state.nativeToken){
			amount  = Tools.convertScientificNotation2Number(Tools.formatNumber(token.amount));
		}else if(token.denom === store.state.displayToken){
			amount  = token.amount;
		}
		coin.amount = amount;
		coin.denom = store.state.displayToken.toUpperCase();
		return coin
	}
	/**
	 * 去除数字的类型是string的尾部连续为 0 的数字
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
	 * 格式化数字的类型是string的数字并在小数点后面补零
	 */
	static formatStringToFixedNumber(str,splitNum){
		if(str.toString().indexOf(".") !== -1) {
			let splitString = str.split(".")[1];
			if(splitString.length > splitNum){
				return str.split(".")[0] + '.' +  splitString.substr(0,splitNum)
			}else {
                let diffNum = 2 - splitString.length;
                for(let i = 0; i < diffNum; i++){
                    splitString += '0'
                }
				return str.split(".")[0] + '.' + splitString
			}
		}else {
			return str + '.00'
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
		//return  Tools.toFixedformatNumber(Tools.formatNumber(num) ,4);
		return  Tools.toFixedformatNumber(num ,4);
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
		var D = date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate() + ' ';
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
	static formatDenom(denom){
		if(denom.toLowerCase() === store.state.displayToken || denom.toLowerCase() === store.state.nativeToken){
			return store.state.nativeToken
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
		if (!str) {return ''};
		return str.toLowerCase().replace(/(\s|^)[a-z]/g, function(char){
			return char.toUpperCase();
		});
	}
	static onlyFirstWordUpperCase(str){
		return str.replace(/(\s|^)[a-z]/, function(char){
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
            if(arr[1].toString().length > afterPointLength){
                value =`${arr[0]}.${arr[1].substring(0, afterPointLength)}`
            }else {
                let diffNum = afterPointLength - arr[1].toString().length;
                for(let i = 0; i < diffNum; i++){
                    arr[1] += '0'
                }
                value = `${arr[0]}.${arr[1]}`
            }
            // value = `${arr[0]}.${arr[1].padEnd(afterPointLength, '0').substring(0, afterPointLength)}`;
        }
        return value;
    }
	static formatTxList(list,txType){
		if(list !== null){
			return list.map(item => {
				let [Amount,Fee,transferAmount,transferFee,tokenId] = ['--','--','--','--','--'];
				let commonHeaderObjList,objList,commonFooterObjList;
				let formatListAmount,fromInformation,toInformation;
				formatListAmount = Tools.formatListAmount(item).amount;
                fromInformation = Tools.formatListAmount(item).fromAddressAndMoniker;
                toInformation = Tools.formatListAmount(item).toAddressAndMoniker;
				Amount = formatListAmount.amountNumber === '--' || formatListAmount.tokenName === '--' ? '--' : `${Tools.formatStringToFixedNumber(formatListAmount.amountNumber,2)} ${formatListAmount.tokenName}`;
				transferAmount = formatListAmount.amountNumber === '--' ? '--' : Tools.formatStringToFixedNumber(formatListAmount.amountNumber,2);
				tokenId = formatListAmount.tokenName === '--' ? '--' : formatListAmount.tokenName;
				if(item.fee.amount && item.fee.denom){
					let feeAmount = Tools.formatAmount3(item.fee,6);
					transferFee = `${feeAmount.amount}`;
					Fee = `${feeAmount.amount} ${feeAmount.denom}`;
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
						'From': fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].address : '--',
						Amount:transferAmount,
						'transferFee': transferFee,
						tokenId,
                        fromMoniker: fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].moniker :'',
                        toMoniker: toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].moniker :'',
						'To':toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ?  toInformation[0].address : '--',
						'listName':'transfer'
					};
				}else if(txType === 'validations'){
                    let Moniker = [];
                    if(item.monikers && JSON.stringify(item.monikers) !== '{}'){
                       for(let value in item.monikers){
                           if(value === item.operator_addr){
                               Moniker.unshift(item.monikers[item.operator_addr])
                           }
                       }
                    }
					objList = {
						'Moniker': Moniker.length === 1 ? Tools.formatString(Moniker[0],15,"...") : "--",
						Amount,
						'OperatorAddr': item.operator_addr ? item.operator_addr : '--',
						'listName':'validations',
						'Self_Bonded': item.self_bonded
					}
				}else if(txType === 'delegations'){
					objList = {
						'From': fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ?  fromInformation[0].address : '--',
						Amount,
						'To': toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].address : '--',
						'listName':'delegations',
						fromMoniker: fromInformation.length > 1 ? fromInformation.length : fromInformation.length === 1 ? fromInformation[0].moniker :'',
						toMoniker: toInformation.length > 1 ? toInformation.length : toInformation.length === 1 ? toInformation[0].moniker :'',
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
	 *   fixedNumber: nonzero
	 * */
	static formatAmount2(param,fixedNumber){
		let amount,amountDenom, amountNumber,amountRadixNumber;
		if(param instanceof Array){
			amount = param[0].amount;
			amountDenom = param[0].denom;
		}else if(param instanceof Object){
			amount = param.amount;
			amountDenom = param.denom;
		}
		if(amount.toString().indexOf('e') !== -1 || amount.toString().indexOf('E') !== -1){
			if(amount.toString().indexOf('.') !== -1){
				amountNumber = new BigNumber(amount).toFixed().toString();
			}else {
				amountNumber = new BigNumber(amount).toFixed().toString() + ".";
			}
			amountRadixNumber = Tools.amountRadix(amountDenom);
		}else {
			if(amount.toString().indexOf('.') !== -1){
				amountNumber = amount.toString();
			}else {
				amountNumber = amount.toString() + ".";
			}
			amountRadixNumber = Tools.amountRadix(amountDenom);
		}
		if(amountRadixNumber > 0 ){
			return `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber)* -1))).toFormat(),fixedNumber)} ${store.state.displayToken.toLocaleUpperCase()}`
		}else {
			if(amountDenom && amountDenom === store.state.nativeToken){
				return `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber) * -1))).toFormat(),fixedNumber)} ${store.state.nativeToken.toLocaleUpperCase()}`
			}else if(amountDenom && amountDenom !== store.state.nativeToken){
				return `${Tools.formatStringToFixedNumber(new BigNumber(amountNumber).toFormat(),fixedNumber)} ${amountDenom.toLocaleUpperCase()}`
			}else if(amountDenom === '' && amountNumber=== '0.'){
				return '--'
			}else {
				return `${Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber) * -1))).toFormat(),fixedNumber)} SHARES`
			}
		}
	}
	/**
	*   Amount iris Radix
	 *   param string
	 *   return Radix length
	* */
	static amountRadix(param,defaultValue = 18){
	let diffValue = 1;
	if(param){
		//radix number length need to minus 1;
		if(param === store.state.nativeToken){
			return  store.state.scaleLength
		}else {
			return  0
		}
		/*switch (param) {
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
				return Constant.RADIXDENOM.IRISNUMBER.length - diffValue;
				break;
			default:
				return defaultValue;
			}*/
		}else {
			return 0
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

	// 转化uptime的方法
	static FormatUptime(number) {
	return `${(number * 100).toFixed(4)}%`
	}
	/**
	 * 格式化数字是字符串类型的百分比数字数字
	 *
	 */
	static formatPercent(percent){
		if (!percent) { return '';}
		percent = percent.toString();
		let formatNumberValue = (Tools.formatContinuousNumberZero(percent) * 100).toString(),number;
		if(formatNumberValue.indexOf('.') !== -1){
			number = Number(formatNumberValue).toFixed(2);
		}else {
			number = formatNumberValue
		}
		return number
	}
	/**
	* 交易列表Amount方法重构(data -> 后端返回数据的结构)
	* */
	static formatListAmount(data){
		let amount,fromAddressAndMoniker = [],toAddressAndMoniker = [];
		if(data && data.type){
			switch (data.type) {
				case Constant.TxType.TRANSFER :
					amount = Tools.formatListByAmount(data.amount);
					if(data.msgs){
	                    data.msgs.forEach( item => {
	                        if(item.msg){
	                            fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.from_address,data.monikers));
	                            toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.to_address,data.monikers))
	                            // item.msg.inputs.forEach( item => {
	                            //     fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.address,data.monikers));
	                            // })
	                            // item.msg.outputs.forEach( item => {
	                            //     toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.address,data.monikers))
	                            // })
	                        }
	                    });
	                }
				    break;
				case Constant.TxType.BURN:
					amount = Tools.formatListByAmount(data.amount);
					if(data.msgs){
					    data.msgs.forEach( item => {
					        if(item.msg){
	                            fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.owner,data.monikers))
	                        }
	                    })
	                }
				    break;
				case Constant.TxType.SETMEMOREGEXP:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.CREATEVALIDATOR:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.EDITVALIDATOR:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.UNJAIL:
					amount = Tools.formatListByAmount(data.amount);
	
					break;
				case Constant.TxType.DELEGATE:
					amount = Tools.formatListByAmount(data.amount);
					if(data.msgs){
					    data.msgs.forEach( item => {
					        if(item.msg){
	                            fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.delegator_addr,data.monikers));
	                            toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.validator_addr,data.monikers))
	                        }
	                    })
	                }
					break;
				case Constant.TxType.BEGINREDELEGATE:
					amount = Tools.formatListByTagsBalance(data.tags);
					if(data.status === 'success'){
					    if(data.tags){
	                        fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags['source-validator'],data.monikers));
	                        toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags['destination-validator'],data.monikers))
	                    }
	                }else {
					    if(data.msgs){
	                        data.msgs.forEach( item => {
	                             if(item.msg){
	                                 fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.validator_src_addr,data.monikers));
	                                 toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.validator_dst_addr,data.monikers))
	                             }
	                        })
	                    }
	                }
					break;
				case Constant.TxType.SETWITHDRAWADDRESS:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.BEGINUNBONDING:
					amount = Tools.formatListByTagsBalance(data.tags);
					if(data.status === 'success'){
					    if(data.tags){
	                        fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags['source-validator'],data.monikers));
	                        toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags['delegator'],data.monikers))
	                    }
	                }else {
					    if(data.msgs){
					        data.msgs.forEach( item => {
					            if(item.msg){
	                                fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.validator_addr,data.monikers));
	                                toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.delegator_addr,data.monikers))
	                            }
	                        })
	                    }
	                }
					break;
				case Constant.TxType.WITHDRAWDELEGATORREWARD:
					amount = Tools.formatListByAmount(data.amount);
					if(data.msgs){
	                    data.msgs.forEach( item => {
	                        if(item.msg){
	                            fromAddressAndMoniker.unshift( Tools.getFromAndToMoniker(item.msg.validator_addr,data.monikers))
	                            toAddressAndMoniker.unshift( Tools.getFromAndToMoniker(item.msg.delegator_addr,data.monikers))
	                        }
	                    })
	                }
	
					break;
				case Constant.TxType.WITHDRAWDELEGATORREWARDSALL:
					amount = Tools.formatListByAmount(data.amount);
	                let fromAddressArray = [],toAddressArray = [];
	                if(data.status === 'success'){
	                    if(data.tags){
	                        for(let item in data.tags){
	                            if(item.startsWith('withdraw-reward-from-validator')){
	                                fromAddressArray.push(item.split('-')[item.split('-').length - 1])
	                            }
	                            if(item === 'delegator'){
	                                toAddressArray.push(data.tags[item])
	                            }
	                        }
	                    }
	                }else {
	                    if(data.msgs){
	                        data.msgs.forEach( item => {
	                            if(item.msg){
	                                if(item.msg.validator_addr){
	                                    fromAddressArray.unshift(item.msg.validator_addr)
	                                }
	                                if(item.msg.delegator_addr){
	                                    toAddressArray.unshift(item.msg.delegator_addr)
	                                }
	                            }
	                        })
	                    }
	                }
	                fromAddressArray.forEach( item => {
	                    fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item,data.monikers))
	                })
	                toAddressArray.forEach( item => {
	                    toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item,data.monikers))
	                })
					break;
				case Constant.TxType.WITHDRAWVALIDATORREWARDSALL:
					amount = Tools.formatListByAmount(data.amount);
	                let withDrawValidatorFromAddressArray = [],withDrawValidatorToAddressArray = [];
	                if(data.tags){
	                    for(let item in data.tags){
	                        if(item.startsWith('withdraw-reward-from-validator')){
	                            withDrawValidatorFromAddressArray.push(item.split('-')[item.split('-').length - 1])
	                        }
	                        if(item === 'withdraw-address'){
	                            withDrawValidatorToAddressArray.push(data.tags[item])
	                        }
	                    }
	                }
	                withDrawValidatorFromAddressArray.forEach( item => {
	                    fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item,data.monikers))
	                })
	                withDrawValidatorToAddressArray.forEach( item => {
	                    toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item,data.monikers))
	                })
					break;
				case Constant.TxType.SUBMITPROPOSAL:
	                fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.signer,data.monikers))
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.DEPOSIT:
	                fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.signer,data.monikers))
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.VOTE:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.ISSUETOKEN:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.EDITTOKEN:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.MINTTOKEN:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.TRANSFERTOKENOWNER:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.CREATEGATEWAY:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.EDITGATEWAY:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.TRANSFERGATEWAYOWNER:
					amount = Tools.formatListByAmount(data.amount);
					break;
				case Constant.TxType.REQUESTRAND:
					amount = Tools.formatListByAmount(data.amount);
					break;
	            case Constant.TxType.CLAIMHTLC:
	                if(data.status === 'success'){
	                    if(data.tags){
	                       fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags.sender,data.monikers))
	                        toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags.receiver,data.monikers))
	                    }
	                }
	                break;
	            case Constant.TxType.REFUNDHTLC:
	                if(data.tags && data.tags.sender){
	                    toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(data.tags.sender,data.monikers))
	                }
	                break;
	            case Constant.TxType.CREATEHTLC:
	                if(data.msgs){
	                    data.msgs.forEach( item => {
	                        if(item.msg){
	                            fromAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.sender,data.monikers))
	                            toAddressAndMoniker.unshift(Tools.getFromAndToMoniker(item.msg.to,data.monikers))
	                        }
	                    })
	                }
	                break;
	            case Constant.TxType.FundCommunityPool:
	                amount = Tools.formatListByAmount(data.amount);
	                break;
	                //TODO(lvshenchao) this type of tx need to be configured
	            case Constant.TxType.WithdrawValidatorCommission:
	                amount = Tools.formatListByAmount(data.amount);
	                break;
	                //TODO(lvshenchao) this type of tx need to be configured
	
	
	
			}
		}
		return {
	        amount,
	        fromAddressAndMoniker,
	        toAddressAndMoniker
	    };
	
	}
	
	/**
	 * 从amount字段中获取amount
	* */
	static formatListByAmount(amount){
		let [amountNumber,tokenName] = ['--','--'];
		if(amount instanceof Array && amount.length > 0) {
	        if(amount.length > 1 ){
	            amountNumber = `${amount.length} Tokens`;
	            tokenName = '--'
	        }else {
	            if (amount[0].denom && amount[0].amount && amount[0].denom === store.state.nativeToken || amount[0].amount == 0) {
	                amountNumber = amount[0].amount > 0 ? Tools.formatStringToFixedNumber(String(Tools.formatAmount3(amount[0]).amount), 2) : Number(amount[0].amount).toFixed(2);
	                tokenName = store.state.displayToken.toLocaleUpperCase();
	            } else if (amount[0].denom && amount[0].amount && amount[0].denom !== store.state.nativeToken) {
	                amountNumber = amount[0].amount;
	                tokenName = amount[0].denom.toLocaleUpperCase();
	            } else {
	                amountNumber = Tools.formatStringToFixedNumber(Tools.FormatScientificNotationToNumber(amount[0].amount), 2);
	                if (amount[0].denom === store.state.nativeToken) {
	                    tokenName = store.state.displayToken.toLocaleUpperCase();
	                } else {
	                    tokenName = amount[0].denom ? amount[0].denom.toLocaleUpperCase() : '--';
	                }
	            }
	        }
		}else if(amount.amount && Object.keys(amount.amount).includes('amount') && Object.keys(amount.amount).includes('denom')){
			if(amount.denom === store.state.nativeToken){
				amountNumber = Tools.formatStringToFixedNumber(String(Tools.formatAmount3(amount.amount).amount),2);
				tokenName = store.state.displayToken.toLocaleUpperCase();
			}else if(amount.denom !== store.state.nativeToken){
				amountNumber = amount.amount
				tokenName = amount.denom.toLocaleUpperCase()
			}else if(!amount.denom){
				amountNumber = amount.amount
				tokenName = ''
			}
		}else if(amount&& Object.keys(amount).includes('amount') && Object.keys(amount).includes('denom')){
	        if(amount.denom === store.state.nativeToken){
	            amountNumber = Tools.formatStringToFixedNumber(String(Tools.formatAmount3(amount).amount),2);
	            tokenName = store.state.displayToken.toLocaleUpperCase();
	        }else if(amount.denom !== store.state.nativeToken){
	            amountNumber = amount.amount
	            tokenName = amount.denom.toLocaleUpperCase()
	        }else if(!amount.denom){
	            amountNumber = amount.amount
	            tokenName = ''
	        }
	    }
		// console.log(amountNumber,tokenName,"amount information")
		return {amountNumber,tokenName}
	}
	/**
	 * 从tags中的balance获取amount
	 * */
	static formatListByTagsBalance(tags,flSplitNum){
		let [amountNumber,tokenName] = ['--','--'];
		if(tags && tags.balance){
			let tokenValue = Tools.formatAccountCoinsAmount(tags.balance);
			let tokenStr = String(Tools.numberMoveDecimal(tokenValue[0],18));
			if(flSplitNum){
	            amountNumber =  tokenStr
	        }else {
	            amountNumber =  Tools.formatStringToFixedNumber(tokenStr,2);
	        }
			tokenName = store.state.displayToken.toLocaleUpperCase()
		}
		return {amountNumber,tokenName}
	
	}
	/**
	 * 格式化fee
	 * */
	static formatFee(Fee){
	    if(Fee.amount && Fee.denom){
	        return `${Tools.formatStringToFixedNumber(String(Tools.formatNumber(Fee.amount)),4)} ${Tools.formatDenom(Fee.denom).toUpperCase()}`;
	    }
	}
	/**
	 * 格式化交易详情页的amount
	 * */
	static formatAmountOfTxDetail(amount){
	    let [amountNumber,tokenName,moreAmountsNumber] = ['--','--',[]];
	    if(amount instanceof Array && amount.length > 0) {
	        if(amount.length !== 1){
	            moreAmountsNumber = amount.map( (item) => {
	                if(item.denom === store.state.nativeToken){
	                    return {
	                        denom : store.state.displayToken.toLocaleUpperCase(),
	                        amount: String(Tools.numberMoveDecimal(item.amount))
	                    }
	                }else {
	                    return {
	                        denom : item.denom.toLocaleUpperCase(),
	                        amount: Tools.FormatScientificNotationToNumber(item.amount)
	                    }
	                }
	            })
	        }else {
	            if (amount[0].denom && amount[0].amount && amount[0].denom === store.state.nativeToken || amount[0].amount == 0) {
	                amountNumber = amount[0].amount > 0 ? String(Tools.numberMoveDecimal(amount[0].amount)) : Number(amount[0].amount).toFixed(2);
	                tokenName = store.state.displayToken.toLocaleUpperCase();
	            } else if (amount[0].denom && amount[0].amount && amount[0].denom !== store.state.nativeToken) {
	                amountNumber = amount[0].amount;
	                tokenName = amount[0].denom.toLocaleUpperCase();
	            } else {
	                amountNumber = Tools.formatStringToFixedNumber(Tools.FormatScientificNotationToNumber(amount[0].amount), 2);
	                if (amount[0].denom === store.state.nativeToken) {
	                    tokenName = store.state.displayToken.toLocaleUpperCase();
	                } else {
	                    tokenName = amount[0].denom ? amount[0].denom.toLocaleUpperCase() : '--';
	
	                }
	            }
	        }
	    }else if(amount.amount && Object.keys(amount.amount).includes('amount') && Object.keys(amount.amount).includes('denom')){
	        if(amount.denom === store.state.nativeToken){
	            amountNumber =String(Tools.numberMoveDecimal(amount.amount));
	            tokenName = store.state.displayToken.toLocaleUpperCase();
	        }else if(amount.denom !== store.state.nativeToken){
	            amountNumber = amount.amount
	            tokenName = amount.denom.toLocaleUpperCase()
	        }else if(!amount.denom){
	            amountNumber = amount.amount
	            tokenName = ''
	        }
	    }else if(amount&& Object.keys(amount).includes('amount') && Object.keys(amount).includes('denom')){
	        if(amount.denom === store.state.nativeToken){
	            amountNumber = String(Tools.numberMoveDecimal(amount.amount));
	            tokenName = store.state.displayToken.toLocaleUpperCase();
	        }else if(amount.denom !== store.state.nativeToken){
	            amountNumber = amount.amount
	            tokenName = amount.denom.toLocaleUpperCase()
	        }else if(!amount.denom){
	            amountNumber = amount.amount
	            tokenName = ''
	        }
	    }
	    // console.log(amountNumber,tokenName,"amount information")
	    return {amountNumber,tokenName,moreAmountsNumber}
	}
	
	static formatPercentage(number){
	    return new BigNumber(number).multipliedBy(100)
	}
	    /**
	    * Form和To字段展示的问题
	    * */
	static getFromAndToMoniker(address,monikers){
	    let resData =
	        {
	            address: address,
	            moniker: monikers[address] ? monikers[address] : ''
	        };
	
	    return resData
	}
	
	static getDisplayDate(timestamp, format = "YYYY-MM-DD HH:mm:ss"){
	    return moment(timestamp).utcOffset(+480).format(format);
	}
	
	static getFormatDate(date, format = "YYYY-MM-DD HH:mm:ss"){
	    return moment(date).utcOffset(+480).format(format);
	}
	
	static isBech32(prefix, str) {
	    if (!prefix || prefix.length == 0) {
	        return false
	    }
	
	    let preReg = new RegExp('^' + prefix + '1');
	    if (!preReg.test(str) ){
	        return false
	    }
	    let allReg = new RegExp(/^[0-9a-zA-Z]*$/i);
	    if (!allReg.test(str)){
	        return false
	    }
	
	    try {
	        bech32.decode(str);
	        return true
	    }catch (e) {
	        return false
	    }
	}
	static formatAmount3(param,fixedNumber){
	    let amount,amountDenom, amountNumber,amountRadixNumber;
	    if(param instanceof Array){
		    amount = param[0].amount;
		    amountDenom = param[0].denom;
	    }else if(param instanceof Object){
		    amount = param.amount;
		    amountDenom = param.denom;
	    }
	    if(amount.toString().indexOf('e') !== -1 || amount.toString().indexOf('E') !== -1){
		    if(amount.toString().indexOf('.') !== -1){
			    amountNumber = new BigNumber(amount).toFixed().toString();
		    }else {
			    amountNumber = new BigNumber(amount).toFixed().toString() + ".";
		    }
		    amountRadixNumber = Tools.amountRadix(amountDenom);
	    }else {
		    if(amount.toString().indexOf('.') !== -1){
			    amountNumber = amount.toString();
		    }else {
			    amountNumber = amount.toString() + ".";
		    }
		    amountRadixNumber = Tools.amountRadix(amountDenom);
	    }
		
		if(amountRadixNumber > 0 ){
			return {
				amount: Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber)* -1))).toFormat(),fixedNumber),
				denom: amountDenom.toLocaleUpperCase()
			}
		}else {
			if(amountDenom && amountDenom === store.state.nativeToken){
				return {
					amount: Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber) * -1))).toFormat(),fixedNumber),
					denom: store.state.nativeToken.toLocaleUpperCase()
				}
			}else if(amountDenom && amountDenom !== store.state.nativeToken){
				return {
					amount: Tools.formatStringToFixedNumber(new BigNumber(amountNumber).toFormat(),fixedNumber),
					denom: amountDenom.toLocaleUpperCase()
				}
			}else if(amountDenom === '' && amountNumber=== '0.'){
				return '--'
			}else {
				return {
					amount: Tools.formatStringToFixedNumber(new BigNumber(moveDecimal(amountNumber,(Number(amountRadixNumber) * -1))).toFormat(),fixedNumber),
					denom: 'SHARES'
				}
			}
		}
	}
}
