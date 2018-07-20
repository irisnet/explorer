<template>
  <div class="echarts_component_wrap">
    <div class="echarts_title_wrap">
      <span class="validators_title">Validators Top10</span>
      <span class="validators_top" @click="viewAllClick()">View All</span>
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
        //根据设备大小显示饼图的大小
        let radius = this.deviceType === 1 ? '80%' : '65%';
        let option = {

          tooltip : {
            trigger: 'item',
            formatter(params){
              let res =  `<span style="display:block;color:#00f0ff;padding:0 0.5rem;">${params.name}</span>`;
              if(params.name !== 'others'){
                res += `<span style="display:block;padding:0 0.5rem;">Uptime: ${params.data.upTime}</span>`;
              }
              res += `<span style="display:block;padding:0 0.5rem;">Voting Power: ${params.value}</span>`;
              return res;
            }
          },
          legend: {
            orient: 'vertical',
            right: '5%',
            data: [],
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
              data:[],
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
          option.legend.data = information.legendData;
          option.series[0].data = information.seriesData;
          pie.setOption(option);
          pie.on('click',(param)=>{
            if(param.data.name !== 'others'){
              this.$router.push(`/address/1/${param.data.address}`)
            }

          })
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

    methods: {
      viewAllClick(){
          this.$router.push('/recent_transactions/3/0')
      }
    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .echarts_component_wrap{
    width:100%;
    height:100%;
    .echarts_title_wrap{
      height:15%;
      @include flex;
      padding:1.2rem 1rem 0 1rem;
      justify-content: space-between;

      .validators_title{
        font-size:1.8rem;
        text-indent:2.5rem;
        font-weight:600;
        background: url('../assets/people.svg') no-repeat 0 -0.2rem;
      }
      .validators_top{
        @include viewBtn;
      }
    }
    #echarts_pie{
      width:100%;
      height:85%;
    }
  }



</style>
