<template>
  <div :class="echartsComponentWrapLine">
    <div class="echarts_title_wrap_line">
      14 day Transaction History
    </div>
    <div id="echarts_line">

    </div>
  </div>


</template>

<script>
  import echarts from 'echarts';

  let line = null;
  export default {
    name: 'echarts-line',
    watch: {
      informationLine(informationLine) {
        //根据设备大小显示饼图的大小
        let radius = this.deviceType === 1 ? '85%' : '65%';
        let option = {
          tooltip : {
            trigger: 'item',
            formatter(params){
              let res =  `<span style="display:block;">${params.name}</span>`;
              res += `<span style="display:block;">Transactions: ${params.value}</span>`;
              return res;
            }
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun','1', '2', '3', '4', '5', '6', '7',],
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },
          },
          yAxis: {
            type: 'value',
            axisLine: {
              lineStyle: {
                color: '#a2a2ae'
              }
            },
            min:0,
            max:1000000,
            splitNumber:2,
            axisLabel:{
              formatter(value) {
                let texts = [];
                if(value===0){
                  texts.push('0');
                }
                else if (value <=100000) {
                  texts.push('100k');
                }
                else if (value<= 500000) {
                  texts.push('500k');
                }
                else if(value<= 1000000){
                  texts.push('1000k');
                }
                return texts.join('/');

              }
            }
          },
          grid: {
            left: '3%',   //图表距边框的距离
            right: '4%',
            bottom: '3%',
            top:'5%',
            containLabel: true
          },
          series: [
            {
              data: [300000, 400000, 1000000, 100000, 600000, 1000000, 100000,700000, 1000000, 100000, 800000, 1000000, 100000, 300000,],
              type: 'line',
              areaStyle: {
                normal: {
                  color: new echarts.graphic.LinearGradient(//设置渐变颜色
                    0, 0, 0, 1,
                    [
                      {offset: 0, color: '#3498db'},
                      {offset: 0.5, color: '#91ccef'},
                      {offset: 1, color: '#dcf6ff'}
                    ]
                  )
                }
              },
              smooth:true,//曲线平滑
              itemStyle:{
                normal:{
                  color:'#3598db',
                  borderColor:'#3598db'  //拐点边框颜色
                }
              },
            }
          ]
        };
        console.log(line)
        if (line) {
          line.setOption(option)
        }
      }
    },
    data() {
      return {
        deviceType: window.innerWidth > 500 ? 1 : 0,
        echartsComponentWrapLine:window.innerWidth > 500 ?'echarts_component_wrap_line_personal_computer':'echarts_component_wrap_line_mobile'

      }
    },
    props: ['informationLine'],
    beforeMount() {

    },
    mounted() {
      line = echarts.init(document.getElementById('echarts_line'));


    },

    methods: {}
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .echarts_component_wrap_line_personal_computer, .echarts_component_wrap_line_mobile {
    width: 100%;
    height: 100%;

    .echarts_title_wrap_line {
      padding:1.2rem 1rem 0 1rem;
      height: 15%;
      font-size:1.8rem;
      font-weight:600;
    }
    #echarts_line {
      width: 100%;
      height: 85%;
    }
  }
  .echarts_component_wrap_line_mobile{
    min-width:40rem;
  }


</style>
