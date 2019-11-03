<template lang="pug">
  .container 
    .header
      el-button New Book
    el-table(:data="tableData" style="width: 100%")
      el-table-column(prop="id" label="ID" width="50")
      el-table-column(prop="name" label="Name")
        template(slot-scope="scope")
          el-input(v-if="currentEditLineId === scope.row.id" v-model="scope.row.name" size="mini")
          span(v-else) {{scope.row.name}}
      el-table-column(prop="author" label="Author")
        template(slot-scope="scope")
          el-input(v-if="currentEditLineId === scope.row.id" v-model="scope.row.author" size="mini")
          span(v-else) {{scope.row.author}}
      el-table-column(prop="createAt" label="CreateAt")
        template(slot-scope="scope")
          el-input(v-if="currentEditLineId === scope.row.id" v-model="scope.row.createAt" size="mini")
          span(v-else) {{scope.row.createAt}}
      el-table-column(prop="updateAt" label="UpdateAt")
        template(slot-scope="scope")
          el-input(v-if="currentEditLineId === scope.row.id" v-model="scope.row.updateAt" size="mini")
          span(v-else) {{scope.row.updateAt}}

      el-table-column(label="Action" width="200")
        template(slot-scope="scope")
          el-button(@click="handleSave(scope.row.id)" v-if="currentEditLineId === scope.row.id" size="mini" type="success") Save
          el-button(@click="handleEdit(scope.row.id)" v-else size="mini" type="primary") Edit
          el-button(size="mini" type="danger") Delete

    
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
  public currentEditLineId: string = ""
  async created() {
    
    this.$http("http://35.247.143.160:8002/api/book").then((res:AxiosResponse<Book[]>) => {
      this.tableData = res.data
    })
  }
  /**
   * handleEdit
   * @param {id}
   */
  public handleEdit(id:string) {
    this.currentEditLineId = id;
  }
  /**
   * handleSave
   * @param {id}
   */
  public handleSave(id:string) {
    this.currentEditLineId = "";
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