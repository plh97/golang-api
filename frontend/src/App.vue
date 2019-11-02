<template lang="pug">
  .container 
    .header
      el-button New Book
    el-table(:data="tableData" style="width: 100%")
      el-table-column(prop="id" label="id")
      el-table-column(prop="name" label="name")
      el-table-column(prop="author" label="author")
      el-table-column(prop="createTime" label="createAt")
      el-table-column(prop="updateTime" label="updateAt")

    
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import {AxiosResponse} from 'axios'
type Book = {
	id:     		  string
	name:  			  string
	author: 			string
	createTime: 	number
	updateTime: 	number
}
@Component({})
export default class MyComponent extends Vue {
  public tableData: Book[] = [];
  created() {
    this.$http("http://35.247.143.160:8002/api/book").then((res:AxiosResponse<Book[]>) => {
      this.tableData = res.data
    })
  }
}
</script>

<style lang="scss" scoped>
.container {
  color: red;
  display: flex;
  flex-direction: column;
  height: 100%;
  .header {
    display: flex;
    justify-content: flex-end;
  }
  .row {
    flex: 1;
  }
}
</style>