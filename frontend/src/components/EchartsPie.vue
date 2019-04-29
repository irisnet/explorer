<template>
  <div class="echarts_component_wrap" :style="`min-width:${minWidth}rem;`">
    <div class="echarts_title_wrap">
      <span class="validators_title">Validators Top10</span>
      <router-link class="validators_top" :to="`/validators`">
        <span>View All</span>
      </router-link>
    </div>
    <div id="echarts_pie">

    </div>
  </div>


</template>

<script>
  import echarts from 'echarts';
  import Tools from "../util/Tools";

  let pie = null;
  export default {
    name: 'echarts-pie',
    watch:{
      innerWidth(innerWidth){
        //根据设备大小显示饼图的大小
        let radius = innerWidth > 1258 === 1 ? '91%' : '80%';
        let legend = innerWidth > 1258 ?
          {
            orient: 'vertical',
            right: '10%',
            data: [],
            top:30,
          } : {
            orient: 'horizontal',
            bottom:'5%',
            data: [],
            type: 'scroll',
          };
        let center = innerWidth > 1258 ? ['34.2%', '50%'] : ['50%', '45%'];
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
        let radius = this.innerWidth > 1258 ? '91%' : '80%';
        let legend = this.innerWidth > 1258 ?
          {
            orient: 'vertical',
            right: '5%',
            data: [],
            top:30,
        } : {
          orient: 'horizontal',
            data: [],
            bottom:'5%',
            type: 'scroll',
        };
        let center = this.innerWidth > 1258 ? ['34.2%', '50%'] : ['50%', '45%'];
        let option = {
          tooltip : {
            trigger: 'item',
            formatter(params){
              let res = `<span style="display:block;color:#00f0ff;padding:0 0.05rem;">${params.name}</span>`;
              if(params.name !== 'others'){
                res += `<span style="display:block;padding:0 0.05rem;">Uptime: ${params.data.upTime}</span>`;
              }
              res += `<span style="display:block;padding:0 0.05rem;">Voting Power: ${(params.value/params.data.powerAll*100).toFixed(2)}%</span>`;
              return res;
            }
          },
          legend,
          series : [
            {
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
        minWidth:2.9,

      }
    },
    props:['information'],
    beforeMount() {
      if(this.innerWidth <= 320){
        this.minWidth = 2.9;
      }else if(this.innerWidth <= 375){
        this.minWidth = 3.4;
      }else if(this.innerWidth <= 414){
        this.minWidth = 3.9;
      }
    },
    mounted() {
      pie = echarts.init(document.getElementById('echarts_pie'));
      window.addEventListener('resize',this.onWindowResize)
    },

    methods: {
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
    padding:0.12rem 0.2rem 0 0.2rem;
    .echarts_title_wrap{
      height:15%;
      @include flex;
      justify-content: space-between;

      .validators_title{
        font-size:0.18rem;
        text-indent:0.35rem;
        @include fontWeight;
        background: url('../assets/people.svg') no-repeat 0 -0.02rem;
      }
      .validators_top{
        span{
          color:#c9eafd;
        }
        @include viewBtn;
      }
    }
    #echarts_pie{
      width:100%;
      height:85%;
    }
  }



</style>
