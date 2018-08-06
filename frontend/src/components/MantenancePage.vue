<template>
  <div class="mantence_wrap">
    <div class="photo_wrap">
      <img src="../assets/error.png" alt="">
    </div>
    <p>
      Sorry!
    </p>
    <h3 style="margin-bottom:0.2rem;margin-top:0.2rem;">
      Our system is under mantenance, please visit later.
    </h3>
    <div class="footer_link_wrap">

      <a class="footer_link_chat"
         target="_blank"
         href="https://twitter.com/irisnetwork">Follow us on Twitter</a>
      <span class="footer_link_join">|</span>
      <a class="footer_link_privacy"
         href="https://www.irisnet.org/"
         target="_blank"
         >About IRISnet</a>
    </div>
  </div>
</template>
<script>

  import Tools from '../common/Tools';

  export default {

    watch: {
      $route() {
        this.showFooter = !(this.$route.query.flShow && this.$route.query.flShow === 'false' && !Tools.currentDeviceIsPersonComputer());
      }
    },
    data() {
      return {}
    },
    beforeMount() {
      if (Tools.currentDeviceIsPersonComputer()) {

      } else {

      }
    },
    mounted() {
      document.getElementById('header').style.display = 'none';
      document.getElementById('footer').style.display = 'none';
    },
    methods: {}
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';

  .mantence_wrap {
    @include flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    p {
      font-size: 0.18rem;
      margin-top: 0.1rem;
    }
    h3 {
      font-size: 0.14rem;
      color: #a2a2ae;
    }

    .footer_link_wrap {
      justify-content: center;
      font-size: 0.14rem;
      @include flex;
      a{
        color:#3598db;
      }
      span {
        font-size: 0.14rem;
        &:nth-child(2n-1) {
          color: #3698db;
          cursor: pointer;
        }
        &:nth-child(2n) {
          color: #a2a2ae;
          margin: 0 0.1rem;
        }
      }
    }
  }


</style>
