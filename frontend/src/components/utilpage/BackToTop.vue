<template>
  <div :class="backToTopBtnWrap">
    <div class="to_top_img_container" @click="toTop" v-show="flShowToTopBtn">
      <img src="../../assets/backtop.png"/>
    </div>
  </div>
</template>

<script>
  import $ from "jquery"
  import Tools from "../../util/Tools"
    export default {
        name: "backToTop",
        data(){
            return {
              flShowToTopBtn: false,
              backToTopBtnWrap: "personal_computer_transactions_detail",
            }
        },
      watch:{
        scrollHeight(scrollHeight){
          if(scrollHeight > 500){
            this.flShowToTopBtn = true
          }else {
            this.flShowToTopBtn = false
          }
        }
      },
      props:['scrollHeight'],
      methods: {
          toTop(){
            $('body,html').animate({
                scrollTop: 0
              }, 500
            );
          }
      },
      beforeMount(){
        if (Tools.currentDeviceIsPersonComputer()) {
          this.backToTopBtnWrap = 'personal_computer_transactions_detail_wrap';
        } else {
          this.backToTopBtnWrap = 'mobile_transactions_detail_wrap';
        }
      }
    }
</script>

<style lang="scss" scoped>
  .personal_computer_transactions_detail_wrap{
    .to_top_img_container{
      width: 0.46rem;
      height: 0.48rem;
      background: rgba(0,0,0,0.3);
      border-radius: 0.05rem 0 0 0.05rem;
      position: fixed;
      right: 0.18rem;
      bottom: 1rem;
      img{
        width: 100%;
        height: 100%;
      }
      &:hover{
        cursor: pointer;
      }
      span{
        color: #Fff;
      }
    }
  }
  .mobile_transactions_detail_wrap{
    .to_top_img_container{
      width: 0.46rem;
      height: 0.48rem;
      background: rgba(0,0,0,0.7);
      border-radius: 0.05rem;
      position: fixed;
      right: 0.1rem;
      bottom: 0.8rem;
      &:hover{
        cursor: pointer;
      }
      span{
        color: #Fff;
      }
    }
  }
</style>
