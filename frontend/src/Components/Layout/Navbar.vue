<template lang="pug">
  .Navbar
    el-dropdown(trigger="click")
      .el-dropdown-link
        el-avatar(shape="square") users
        i.el-icon-caret-bottom
      el-dropdown-menu(slot="dropdown")
        el-dropdown-item(@click="handleLogout") log outs
</template>

<script lang="ts">
import Vue from 'vue'
import Component from "vue-class-component";

@Component({})
export default class Navbar extends Vue {
  public async handleLogout(){
    debugger
    // await this.$store.dispatch('user/logout')
    this.$router.push('/login')
    this.$message.success('Successfully logout')
  }
}
</script>

<style lang="scss" scope>
.Navbar {
  display: flex;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
  height: 50px;
  width: 100%;
  justify-content: flex-end;
  padding: 0 30px;
  box-sizing: border-box;
  align-items: center;
  .el-dropdown-link {
    display: flex;
    flex-direction: row;
    align-items: center;
    cursor: pointer;
  }
  
}
</style>