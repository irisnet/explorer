<!--
  feature: loading,如果1000ms内没有结果返回则显示;
  use:     在父组件中引入，传入props,并且父元素设置position为relative.
  props:   showLoading:boolean 是否显示弹框
-->
<template>
  <div class="spin-component" v-show="showLoading">
    <div class="spin-content">
      <img style="with:48px;height: 48px;" src="../../assets/loading.gif">
      <p>Loading,please wait...</p>
    </div>
  </div>

</template>
<script>
  export default {
    name: "SpinComponent",
    // watch: {
    //   showLoading(showLoading) {
    //     if (showLoading) {//如果父组件传值需要显示loading框的时候就开始设置延时器
    //       if (!this.func) {
    //         this.func = setTimeout(this.timeoutShowLoading.bind(this), 1000);
    //       }

    //     } else {//控制不显示弹框
    //       //判断是否在规定时间内延时器已经生效
    //       try {
    //         clearTimeout(this.func);
    //       }catch (e) {
    //         console.log('延时器已经生效，不必清除')
    //       }

    //       this.func = '';
    //       if (this.loading) {
    //         this.loading = false;
    //       }
    //     }
    //   }
    // },
    data() {
      return {
        loading: false,
      }
    },
    props: ['showLoading'],
    //父组件传一个变量，从这个时候自组件开始延时1.5s显示loading框，如果在1.5s内返回了数据，就清除延时器
    mounted() {
    //   this.func = setTimeout(this.timeoutShowLoading.bind(this), 1000);

    },
    // methods: {
    //   timeoutShowLoading() {
    //     this.loading = true;
    //   }
    // }
  }
</script>

<style scoped lang="scss">
  @import '../../style/mixin.scss';

  .spin-component {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;

    z-index: 1000;
    @include flex;
    align-items: center;
    justify-content: center;
    .spin-content {
      background: rgba(255, 255, 255, 1);
      box-shadow:0px 0px 6px 0px rgba(10, 97, 142, 0.15);
      width: 2.54rem;
      height: 1.32rem;
      @include borderRadius(0.1rem);
      @include flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      img {

      }
      p {
        color: #000;
        font-size: 0.14rem;
        font-family: ArialMT;
        margin-top: 0.19rem;
      }
    }
  }

</style>
