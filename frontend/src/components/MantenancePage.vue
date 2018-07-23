<template>
  <div class="mantence_wrap">
    <div class="photo_wrap">
      <img src="../assets/error.png" alt="">
    </div>
    <p>
      Sorry!
    </p>
    <h3>
      Our system is under mantenance, please visit later!
    </h3>
    <div class="footer_link_wrap">
      <span class="footer_link_contact">Join Testnet</span>
      <span class="footer_link_forum">|</span>
      <span class="footer_link_chat">Contact Us</span>
      <span class="footer_link_join">|</span>
      <span class="footer_link_privacy" @click="footerLinkClick('/privacy_policy')">Privacy Policy</span>
    </div>
  </div>
</template>
<script>

  import Tools from '../common/Tools';

  export default {

    watch:{
      $route(){
        this.showFooter = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
      }
    },
    data() {
      return {

      }
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {

      } else {

      }
    },
    methods: {

    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .mantence_wrap{
    @include flex;
    flex-direction:column;
    align-items:center;
    justify-content: center;

    p{
      font-size:18px;
      margin-top:1rem;
    }
    h3{
      font-size:14px;
      color: #a2a2ae;
    }

    .footer_link_wrap {
      justify-content: center;
      font-size:14px;
      @include flex;
      span{
        font-size:1.4rem;
        &:nth-child(2n-1){
          color:#3698db;
          cursor:pointer;
        }
        &:nth-child(2n){
          color:#a2a2ae;
          margin:0 1rem;
        }
      }
    }
  }


</style>
