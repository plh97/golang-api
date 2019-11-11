<template lang="pug">
  .App-Main
    router-view(v-loading="loading")
</template>

<script lang="ts">
import Vue from "vue";
import { mapGetters } from 'vuex'
import Component from "vue-class-component";

@Component({
  computed: {
    ...mapGetters([
      'loading'
    ])
  }
})
export default class Layout extends Vue {
  
}
</script>

<style lang="scss">
.App-Main {
  height: 100vh;
  overflow: scroll;
  padding: 10px;
  box-sizing: border-box;
  -webkit-overflow-scrolling: touch; /* Lets it scroll lazy */
}
</style>