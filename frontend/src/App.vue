<template>
  <div id="app">
      <app-header></app-header>
    <router-view/>
  </div>
</template>
<script>
    import AppHeader from './components/AppHeader';
    export default {
        components:{
            AppHeader,
        },
        data(){
            return {
                show:1,
            }
        }
    }
</script>
<style lang="scss">
    body,html{
        width:100%;
        height:100%;
    }
    #app{
        width:100%;
        height:100%;
    }
</style>