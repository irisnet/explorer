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
      innerWidth(innerWidth){
        //根据设备大小显示饼图的大小
        let radius = innerWidth > 1258 === 1 ? '90%' : '80%';
        let legend = innerWidth > 1258 ?
          {
            orient: 'vertical',
            right: '10%',
            data: [],
          } : {
            orient: 'horizontal',
            y : 'bottom',
            data: [],
            type: 'scroll',
          };
        let center = innerWidth > 1258 ? ['40%', '50%'] : ['50%', '45%'];
        let option = {

          tooltip : {
            trigger: 'item',
            formatter(params){
              let res =  `<span style="display:block;color:#00f0ff;padding:0 0.05rem;">${params.name}</span>`;
              if(params.name !== 'others'){
                res += `<span style="display:block;padding:0 0.05rem;">Uptime: ${params.data.upTime}</span>`;
              }
              res += `<span style="display:block;padding:0 0.05rem;">Voting Power: ${(params.value/params.data.totalCount*100).toFixed(2)}%</span>`;
              return res;
            }
          },
          legend,
          series : [
            {
              name: '访问来源',
              type: 'pie',
              radius,
              center,
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
          option.legend.data = this.information.legendData;
          option.series[0].data = this.information.seriesData;
          pie.setOption(option);
          pie.on('click',(param)=>{
            if(param.data.name !== 'others'){
              this.$router.push(`/address/1/${param.data.address}`)
            }

          })
        }
      },
      information(information){
        //根据设备大小显示饼图的大小
        let radius = this.innerWidth > 1258 ? '90%' : '80%';
        let legend = this.innerWidth > 1258 === 1 ?
          {
            orient: 'vertical',
            right: '10%',
            data: [],
        } : {
          orient: 'horizontal',
            y : 'bottom',
            data: [],
            type: 'scroll',
        };
        let center = this.innerWidth > 1258 ? ['40%', '50%'] : ['50%', '45%'];
        let option = {

          tooltip : {
            trigger: 'item',
            formatter(params){
              let res =  `<span style="display:block;color:#00f0ff;padding:0 0.05rem;">${params.name}</span>`;
              if(params.name !== 'others'){
                res += `<span style="display:block;padding:0 0.05rem;">Uptime: ${params.data.upTime}</span>`;
              }
              res += `<span style="display:block;padding:0 0.05rem;">Voting Power: ${(params.value/params.data.totalCount*100).toFixed(2)}%</span>`;
              return res;
            }
          },
          legend,
          series : [
            {
              name: '访问来源',
              type: 'pie',
              radius,
              center,
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
        innerWidth:window.innerWidth,

      }
    },
    props:['information'],
    beforeMount() {

    },
    mounted() {
      pie = echarts.init(document.getElementById('echarts_pie'));
      window.addEventListener('resize',this.onWindowResize)
    },

    methods: {
      viewAllClick(){
          this.$router.push('/validators/3/0')
      },
      onWindowResize(){
        pie.resize();
        this.innerWidth = window.innerWidth;
      },
    },
    beforeDestroy(){
      window.removeEventListener('resize',this.onWindowResize);
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
      padding:0.12rem 0.1rem 0 0.1rem;
      justify-content: space-between;

      .validators_title{
        font-size:0.18rem;
        text-indent:0.35rem;
        font-weight:600;
        background: url('../assets/people.svg') no-repeat 0 -0.02rem;
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
