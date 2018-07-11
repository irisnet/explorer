<template>
  <div id="app">
      <app-header></app-header>
      <div id="router_wrap">
          <router-view/>
      </div>

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
    @import './style/mixin.scss';
    body,html{
        width:100%;
        height:100%;
    }
    #app{
        width:100%;
        height:100%;
        @include flex();
        flex-direction:column;
        #router_wrap{
            flex:1;
            overflow:auto;
        }
    }
</style>