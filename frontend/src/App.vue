<template lang="pug">
  .container 
    .top
      el-button New Book
    el-table(:data="tableData")
    
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
type Book = {
  name: string
  age: number
}
@Component({})
export default class MyComponent extends Vue {
  public tableData: Book[] = [];
  created() {
    this.$http("http://35.247.143.160:8080").then(res=>{

    })
  }
}
</script>

<style lang="scss" scoped>
.container {
  color: red;
  display: flex;
  flex-direction: row;
  height: 100%;
  .row {
    flex: 1;
  }
}
</style>