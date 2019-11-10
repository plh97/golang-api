<template lang="pug">
  .container
    .header
      el-button(type="primary" size="mini" @click="centerDialogVisible=true") Add Book
    el-table(:data="tableData" style="width: 100%" v-loading="loading")
      el-table-column(prop="_id" label="NO" width="70")
        template(slot-scope="scope")
          span {{ scope.$index+1 }}
      el-table-column(prop="name" label="Name")
        template(slot-scope="scope")
          el-input(v-if="currentEditLineId === scope.row._id" v-model="scope.row.name" size="mini")
          span(v-else) {{scope.row.name}}
      el-table-column(prop="author" label="Author")
        template(slot-scope="scope")
          el-input(v-if="currentEditLineId === scope.row._id" v-model="scope.row.author" size="mini")
          span(v-else) {{scope.row.author}}
      el-table-column(prop="createAt" label="CreateAt")
        template(slot-scope="scope")
          span {{scope.row.createAt | parseDate}}
      el-table-column(prop="updateAt" label="UpdateAt")
        template(slot-scope="scope")
          span {{scope.row.updateAt | parseDate}}
      el-table-column(label="Action" width="160")
        template(slot-scope="scope")
          el-button(@click="handleUpdateBook(scope.row)" v-if="currentEditLineId === scope.row._id" size="mini" type="success") Save
          el-button(@click="handleEditBook(scope.row._id)" v-else size="mini" type="primary") Edit
          el-button(@click="handleDeleteBook(scope.row._id)" size="mini" type="danger") Delete
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
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { getBooks, updateBook, deleteBook, addBook } from "../dataStore";
import { mapGetters } from 'vuex'
// import { Form } from "element-ui";
// export type VForm = Vue & { validate: (valid: boolean) => boolean };
type Book = {
  _id: string;
  name: string;
  author: string;
  createAt: Date;
  updateAt: Date;
};
@Component({
  computed: {
    ...mapGetters([
      'loading'
    ])
  },
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
    form: HTMLFormElement;
  };
  public centerDialogVisible: boolean = false;
  public tableData: Book[] = [];
  public currentEditLineId: string = String(Math.random());
  public dialogForm: Book = {
    _id: "234",
    name: "23456",
    author: "23456",
    createAt: new Date(),
    updateAt: new Date()
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
    await this.init();
  }
  public async init(){
    try {
      this.tableData = await getBooks()
    } catch (err) {
      this.$message.error(err.message);
    }
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
    await updateBook(book);
    this.init()
    this.currentEditLineId = "";
  }
  /**
   * handleSaveBook
   * @param {id}
   */
  public async handleDeleteBook(id: string) {
    await deleteBook(id);
    this.init()
  }
  /**
   * handleCreateBook
   * @param {id}
   */
  public async handleCreateBook(formName) {
    this.$refs.form.validate(async (valid: boolean) => {
      if (valid) {
        this.centerDialogVisible = false;
        await addBook(this.dialogForm);
        this.init()
        this.dialogForm.name = "";
        this.dialogForm.author = "";
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