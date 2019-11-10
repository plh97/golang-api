<template lang="pug">
  .container
    router-view
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
export default class MyComponent extends Vue {
}
</script>

<style lang="scss">
.container {
  height: 100vh;
}
</style>
<style>