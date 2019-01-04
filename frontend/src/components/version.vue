<template>
    <div>
<<<<<<< HEAD
      <p class="version_title">Version: 0.7.5</p>
=======
      <p class="version_title">Version: 0.7.0</p>
>>>>>>> feature/0.7.2
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
