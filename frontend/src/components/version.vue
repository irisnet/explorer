<template>
    <div>
      <p class="version_title">Version: 1.0.0</p>
      <button class="go_home_btn" @click="goHome">HOME</button>
    </div>
</template>

<script>
    export default {
      data(){
          return {

          }
      },
      methods:{
          goHome(){
            this.$router.push("/home")
          }
      }
    }
</script>

<style scoped lang="scss">
  .version_title{
    padding: 0.5rem 0 0.2rem 0.5rem;

  }
  .go_home_btn{
    margin-left: 0.5rem;
  }
</style>
