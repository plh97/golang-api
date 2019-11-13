<template lang="pug">
  .Layout
    Navbar
    .body
      router-view(v-loading="loading")
</template>

<script lang="ts">
import Vue from "vue";
import { mapGetters } from "vuex";
import Component from "vue-class-component";
import Navbar from "./Navbar.vue";

@Component({
  components: { Navbar },
  computed: {
    ...mapGetters(["loading"])
  }
})
export default class Layout extends Vue {}
</script>

<style lang="scss" scope>
.Layout {
  height: 100vh;
  overflow: scroll;
  // box-sizing: border-box;
  -webkit-overflow-scrolling: touch; /* Lets it scroll lazy */

  .body {
    padding: 10px;
    height: calc(100vh - 50px);
    box-sizing: border-box;
  }
}
</style>
