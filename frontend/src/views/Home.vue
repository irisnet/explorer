<template>
  <div class="home">
    <Header/>
    <div class="content">
      <router-view/>
    </div>
  </div>
</template>

<script lang="ts">
  import {Component, Vue} from 'vue-property-decorator';
  import Header from '@/components/AppHeader.vue'; // @ is an alias to /src

  @Component({
    components: {
      Header,
    },
  })
  export default class Home extends Vue {
  }
</script>

<style>
  .content {
    background: darkgrey;
    width: 100%;
    height: 100%;
    position: absolute;
  }
</style>
