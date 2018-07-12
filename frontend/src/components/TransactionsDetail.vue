<template>
  <div class="transactions_detail_wrap">
    <p>
      <span>Transaction</span>
      <span>{{hashValue}}</span>
    </p>
    <div>
      <p>Transaction Information</p>
      <div>

      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        devicesWidth: window.innerWidth,
        pageClassName: 'personal_computer_home_wrap',//1是显示pc端，0是移动端
        hashValue:'',


      }
    },
    beforeMount() {
      if (this.devicesWidth > 500) {
        this.pageClassName = 'personal_computer_home_wrap';
      } else {
        this.pageClassName = 'mobile_home_wrap';
      }
    },
    mounted() {

    },
    methods: {

    }
  }
</script>
<style lang="scss">
  @import '../style/mixin.scss';


</style>
