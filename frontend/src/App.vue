<template lang="pug">
  .container 
    .header
      el-button(type="primary" size="mini") Add Book
    el-table(:data="tableData" style="width: 100%" v-loading="loading.global")
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
          span {{scope.row.createAt | parseDate}}
      el-table-column(prop="updateAt" label="UpdateAt")
        template(slot-scope="scope")
          span {{scope.row.updateAt | parseDate}}

      el-table-column(label="Action" width="160")
        template(slot-scope="scope")
          el-button(@click="handleUpdateBook(scope.row)" v-if="currentEditLineId === scope.row.id" size="mini" type="success") Save
          el-button(@click="handleEditBook(scope.row.id)" v-else size="mini" type="primary") Edit
          el-button(@click="handleDeleteBook(scope.row.id)" size="mini" type="danger") Delete

    
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { getBooks, updateBook, deleteBook, addBook } from "./dataStore";
type Book = {
  id: string;
  name: string;
  author: string;
  createTime: number;
  updateTime: number;
};
type Loading = {
  global: boolean;
};
@Component({
  filters: {
    parseDate(date) {
      return (
        new Date(date).toLocaleDateString() +
        " " +
        new Date(date).toLocaleTimeString()
      );
    }
  }
})
export default class MyComponent extends Vue {
  public tableData: Book[] = [];
  public currentEditLineId: string = "";
  public loading: Loading = {
    global: false
  };
  async mounted() {
    this.loading.global = true;
    this.tableData = await getBooks().catch(err => {
      this.$message.error(err.message);
    });
    this.loading.global = false;
  }
  /**
   * handleEdit
   * @param {id}
   */
  public handleEditBook(id: string) {
    this.currentEditLineId = id;
  }
  /**
   * handleUpdateBook
   * @param {id}
   */
  public async handleUpdateBook(book: Book) {
    this.currentEditLineId = "";
    const res = await updateBook(book);
    debugger;
  }
  /**
   * handleSave
   * @param {id}
   */
  public async handleDeleteBook(id: string) {
    const res = await deleteBook(id);
    debugger;
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