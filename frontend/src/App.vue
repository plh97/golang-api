<template lang="pug">
  .container 
    el-dialog(
      title="Add Book"
      :visible.sync="centerDialogVisible"
      width="600px"
    )
      .dialog-container
        el-form(ref="form" :model="dialogForm" label-width="80px" :rules="dialogRules")
          el-form-item(label="Name" prop="name")
            el-input(v-model="dialogForm.name")
          el-form-item(label="Author" prop="author")
            el-input(v-model="dialogForm.author")
      span.dialog-footer(slot="footer")
        el-button(@click="centerDialogVisible = false") Cancel
        el-button(type="primary" @click="handleCreateBook") Comfirm
    .header
      el-button(type="primary" size="mini" @click="centerDialogVisible=true") Add Book
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
// import { Form } from "element-ui";
// export type VForm = Vue & { validate: (valid: boolean) => boolean };
type Book = {
  id: string;
  name: string;
  author: string;
  createTime: Date;
  updateTime: Date;
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
  $refs!: {
    form: HTMLFormElement
  }
  public centerDialogVisible: boolean = false;
  public tableData: Book[] = [];
  public currentEditLineId: string = String(Math.random());
  public loading: Loading = {
    global: false
  };
  public dialogForm: Book = {
    id: "",
    name: "",
    author: "",
    createTime: new Date(),
    updateTime: new Date()
  };
  public dialogRules = {
    name: [
      { required: true, message: "Require Book Name", trigger: "blur" },
      { min: 3, message: "At Least 3 Letter", trigger: "blur" },
      { max: 10, message: "No More Than 10 Letter", trigger: "blur" }
    ],
    author: [
      { required: true, message: "Require Book Author", trigger: "blur" },
      { min: 3, message: "At Least 3 Letter", trigger: "blur" },
      { max: 10, message: "No More Than 10 Letter", trigger: "blur" }
    ]
  };
  async mounted() {
    this.loading.global = true;
    this.tableData = await getBooks().catch(err => {
      this.$message.error(err.message);
    });
    this.loading.global = false;
  }
  /**
   * handleEditBook
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
    this.loading.global = true;
    this.tableData = await updateBook(book);
    this.currentEditLineId = "";
    this.loading.global = false;
  }
  /**
   * handleSaveBook
   * @param {id}
   */
  public async handleDeleteBook(id: string) {
    this.loading.global = true;
    this.tableData = await deleteBook(id);
    this.loading.global = false;
  }
  /**
   * handleCreateBook
   * @param {id}
   */
  public async handleCreateBook(formName) {
    this.$refs.form.validate(async(valid: boolean) => {
      if (valid) {
        this.centerDialogVisible = false
        this.tableData = await addBook(this.dialogForm)
        this.dialogForm.name = ''
        this.dialogForm.author = ''
      }
    });
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
  .dialog-container {
    padding-right: 30px;
  }
}
</style>