<template>
  <div class="echarts_component_wrap">
    <div class="echarts_title_wrap">
      hello
    </div>
    <div id="echarts_pie">

    </div>
  </div>


</template>

<script>
  import echarts from 'echarts';

  let pie = null;
  export default {
    name: 'echarts-pie',
    watch:{
      information(information){
        console.log(information)

        //根据设备大小显示饼图的大小
        let radius = this.deviceType === 1 ? '85%' : '65%';


        let option = {

          tooltip : {
            trigger: 'item',
            formatter: "{a} <br/>{b} : {c} ({d}%)"
          },
          legend: {
            orient: 'vertical',
            right: '5%',
            data: ['直接访问','邮件营销','联盟广告','视频广告','搜索引擎','1','2','3','4','5','6']
          },

          series : [
            {
              name: '访问来源',
              type: 'pie',
              radius,
              center: ['35%', '50%'],
              label:{
                show:false
              },
              data:[
                {value:1, name:'直接访问',itemStyle:{color:'#3498db'}},
                {value:2, name:'邮件营销',itemStyle:{color:'#47a2df'}},
                {value:3, name:'联盟广告',itemStyle:{color:'#59ade3'}},
                {value:4, name:'视频广告',itemStyle:{color:'#6cb7e7'}},
                {value:5, name:'搜索引擎',itemStyle:{color:'#7fc2eb'}},
                {value:6, name:'1',itemStyle:{color:'#91ccef'}},
                {value:7, name:'2',itemStyle:{color:'#a4d7f3'}},
                {value:8, name:'3',itemStyle:{color:'#b7e1f7'}},
                {value:9, name:'4',itemStyle:{color:'#c9ecfb'}},
                {value:10, name:'5',itemStyle:{color:'#dcf6ff'}},
                {value:45, name:'6',itemStyle:{color:'#efeff1'}},

              ],
              itemStyle: {
                emphasis: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
              }
            }
          ]
        };
        if(pie){
          pie.setOption(option)
        }
      }
    },
    data() {
      return {
        deviceType:window.innerWidth > 500 ? 1 : 0,

      }
    },
    props:['information'],
    beforeMount() {

    },
    mounted() {
      pie = echarts.init(document.getElementById('echarts_pie'));



    },

    methods: {}
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .echarts_component_wrap{
    width:100%;
    height:100%;
    .echarts_title_wrap{
      height:10%;
    }
    #echarts_pie{
      width:100%;
      height:90%;
    }
  }



</style>
