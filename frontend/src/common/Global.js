export default class Global {
   static formatTimeDiff(time){
      let dateBegin = new Date(time);//将-转化为/，使用new Date
      let dateEnd = new Date();
      let dateDiff = dateEnd.getTime() - dateBegin.getTime();//时间差的毫秒数
      let dayDiff = Math.floor(dateDiff / (24 * 3600 * 1000));//计算出相差天数

      let leave1=dateDiff%(24*3600*1000);   //计算天数后剩余的毫秒数

      let hours=Math.floor(leave1/(3600*1000)); //计算出小时数

      //计算相差分钟数
      let leave2=leave1%(3600*1000);   //计算小时数后剩余的毫秒数

      let minutes=Math.floor(leave2/(60*1000)); //计算相差分钟数

      //计算相差秒数
      let leave3=leave2%(60*1000) ;    //计算分钟数后剩余的毫秒数

      let seconds=Math.round(leave3/1000);

      return`${dayDiff?`${dayDiff}d`:''}
        ${hours ? `${hours}h` : ''}
        ${dayDiff ? '' : (minutes ? `${minutes}m`:'')}
        ${dayDiff || hours ? '' : (seconds ? `${seconds}s`:'')}`
    }
}
