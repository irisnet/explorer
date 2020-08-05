<template>
  <div :class="echartsComponentWrapLine">
    <div class="echarts_title_wrap_line">
      <span class="echarts_title"><i class="iconfont iconTransactionHistory"></i>14 day Transaction History</span>
      <span class="view_all_content"><router-link @click.native="$uMeng.push('HomeOverview_14 day Transaction History','click')" :to="`/txs?txType=&status=&startTime=${startTime}&endTime=${endTime}&page=1`">View All</router-link></span>
    </div>
    <div id="echarts_line">

    </div>
  </div>


</template>

<script>
    var echarts = require('echarts/lib/echarts')
  require('echarts/lib/chart/line')
  import Constant from "../../constant/Constant"
  let line = null;
  export default {
    name: 'echarts-line',
    watch: {
      '$store.state.currentSkinStyle'(newVal){
        if(newVal !== 'default '){
          this.setThemeStyle();
          this.setEcartsLine(this.informationLine)
        }
      },
      informationLine(informationLine) {
        if(this.$store.state.currentSkinStyle !== 'default'){
          this.setThemeStyle();
          this.setEcartsLine(this.informationLine)
            this.startTime = this.formatDate(this.informationLine.xData[0])
            this.endTime = this.formatDate(this.informationLine.xData[this.informationLine.xData.length - 1])
        }
      },
    },
    data() {
      return {
        skinStyle:'',
        endTime:'',
        startTime:'',
        lineSkinBackgroundStyle:[

        ],
        deviceType: window.innerWidth > 500 ? 1 : 0,
        echartsComponentWrapLine:window.innerWidth > 500 ?'echarts_component_wrap_line_personal_computer':'echarts_component_wrap_line_mobile',
        month:['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'],
        monthNum:['01','02','03','04','05','06','07','08','09','10','11','12'],

      }
    },
    props: ['informationLine'],
    beforeMount() {

    },
    created(){

    },
    mounted() {
      line = echarts.init(document.getElementById('echarts_line'));
      window.addEventListener('resize',this.onWindowResize);
    },

    methods: {
      formatDate(time){
          return `${String(time).substr(6,4)}-${String(time).substr(0,2)}-${String(time).substr(3,2)}`
      },
      onWindowResize(){
        line.resize();
      },
      setEcartsLine(informationLine){
        //根据设备大小显示饼图的大小
        let radius = this.deviceType === 1 ? '85%' : '65%';
        let option = {
          tooltip : {
            trigger: 'axis',
            axisPointer:{
              axis:"x",
              type:"line",
              lineStyle:{
                color:"var(--contentColor)",
              },
            },
            formatter(params){
              let res =  `<span style="display:block;">${params[0].name.substr(6,12)}/${params[0].name.substr(0,5)}</span>`;
              res += `<span style="display:block;">Transactions: ${params[0].value}</span>`;
              return res;
            }
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: [],
            axisLine: {
              lineStyle: {
                color: 'var(--contentColor)'
              }
            },
            axisLabel:{
              interval:0,
              rotate:45,
              margin:12,
              formatter:(value)=>{
                return `${this.month[this.monthNum.findIndex(item=>value.substr(0,2) === item)]}${value.substr(3,2)}`;
              }
            }
          },
          yAxis: {
            type: 'value',
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },
            min:0,
            splitNumber:5,
            splitLine: {
              lineStyle: {
                // 使用深浅的间隔色
                color: ['#eee']
              }
            }
          },
          grid: {
            left: '3%',   //图表距边框的距离
            right: '5%',
            bottom: '3%',
            top:'5%',
            containLabel: true
          },
          series: [
            {
              data: [],
              type: 'line',
              areaStyle: {
                normal: {
                  color: new echarts.graphic.LinearGradient(//设置渐变颜色
                          0, 0, 0, 1,
                          this.lineSkinBackgroundStyle
                  )
                }
              },
              smooth:true,//曲线平滑
              smoothMonotone: 'x',
              itemStyle:{
                normal:{
                  color:this.skinStyle,
                  borderColor:this.skinStyle,  //拐点边框颜色
                }
              },
            }
          ]
        };
        if (line) {
          option.xAxis.data = informationLine.xData;
          option.series[0].data = informationLine.seriesData;
          line.setOption(option)
        }
      },
      setThemeStyle(){
        if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.IRISHUB){
          this.skinStyle = '#3264FD';
          this.lineSkinBackgroundStyle = [
            {offset: 0, color: 'rgba(50, 100, 253, 0.8)'},
            {offset: 1, color: 'rgba(50, 100, 253, 0)'}
          ]
        }else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.FUXI){
          this.skinStyle = '#0C4282';
          this.lineSkinBackgroundStyle = [
            {offset: 0, color: 'rgba(12, 66, 130, 0.8)'},
            {offset: 1, color: 'rgba(12, 66, 130, 0)'}
          ]
        }else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.NYANCAT){
          this.skinStyle = '#0D9388';
          this.lineSkinBackgroundStyle = [
            {offset: 0, color: 'rgba(13, 147, 136, 0.8)'},
            {offset: 1, color: 'rgba(245, 247, 253, 1)'}
          ]
        }else if(this.$store.state.currentSkinStyle ===  Constant.CHAINID.BIFROST){
          this.skinStyle = '#D6793A';
          this.lineSkinBackgroundStyle = [
            {offset: 0, color: 'rgba(214,121,58, 0.8)'},
            {offset: 1, color: 'rgba(245, 247, 253, 1)'}
          ]
        }else {
          this.skinStyle = '#0580D3';
          this.lineSkinBackgroundStyle = [
            {offset: 0, color: 'rgba(5,128,211, 0.8)'},
            {offset: 1, color: 'rgba(5,128,211, 0)'}
          ]
        }
      }
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.onWindowResize);
    }
  }
</script>
<style lang="scss">
  @import '../../style/mixin';

  .echarts_component_wrap_line_personal_computer, .echarts_component_wrap_line_mobile {
    padding:0.12rem 0.2rem 0 0.2rem;
    .echarts_title_wrap_line {
      height: 15%;
      font-weight:400;
      display: flex;
      justify-content: space-between;
      color:var(--bgColor);
     .echarts_title{
       font-size:0.18rem;
       color:#000 !important;
       i{
         color: #C8D1DA;
         padding-right: 0.1rem;
       }
     }
      .view_all_content{
        a{
          font-size: 0.14rem;
          border-bottom: 0.01rem solid var(--bgColor);
          height: 0.2rem;
          line-height: 0.2rem;
          color: var(--bgColor)!important;
          i{
            display: inline-block;
            font-size: 0.14rem;
            transform: rotate(-90deg);
          }
        }

      }
    }
    #echarts_line {
        padding-top: 0.3rem;
      height: 2.8rem;
    }
  }
  .echarts_component_wrap_line_mobile{
    min-width:4rem;
  }


</style>
