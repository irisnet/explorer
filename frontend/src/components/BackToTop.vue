<template>
  <div>
    <div class="to_top_img_container" @click="toTop" v-show="flShowToTopBtn">
      <img src="../assets/backtop.png"/>
    </div>
  </div>
</template>

<script>
  import $ from "jquery"
    export default {
        name: "backToTop",
        data(){
            return {
              flShowToTopBtn: false
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
      }
    }
</script>

<style lang="scss" scoped>
  .to_top_img_container{
    width: 0.46rem;
    height: 0.48rem;
    background: rgba(0,0,0,0.7);
    border-radius: 0.05rem;
    position: fixed;
    right: 0.18rem;
    bottom: 1.52rem;
    &:hover{
      cursor: pointer;
    }
    span{
      color: #Fff;
    }
  }
</style>
