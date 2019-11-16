<template lang="pug">
  .Navbar
    el-dropdown(trigger="click")
      .el-dropdown-link
        el-avatar(shape="square") {{name}}
        i.el-icon-caret-bottom
      el-dropdown-menu(slot="dropdown")
        el-dropdown-item 
          span(@click="handleLogout") log out
</template>

<script lang="ts">
import Vue from 'vue'
import Component from "vue-class-component";
import { mapGetters } from "vuex";

@Component({
  computed: {
    ...mapGetters([
      "name"
    ])
  },
  methods: {
    async handleLogout(){
      await this.$store.dispatch('user/logout')
      this.$router.push('/login')
      this.$message.success('Successfully logout')
    }
  }
})
export default class Navbar extends Vue {
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