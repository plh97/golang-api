<template lang="pug">
  .Login
    .title Log In
    el-form(label-width="120px" :model="loginForm")
      el-form-item(label="Name")
        el-input(v-model="loginForm.name")
      el-form-item(label="Password")
        el-input(v-model="loginForm.password")
      el-form-item
        el-button(type="primary" @click="handleLogin") Login
        span &nbsp;&nbsp;
        router-link(to="/register")
          el-button register

</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Account } from "../dataStore";

type LoginType = {
  name: string;
  password: string;
};

@Component({})
export default class Login extends Vue {
  public loginForm: LoginType = {
    name: "",
    password: ""
  };
  public async handleLogin(): Promise<void> {
    await this.$store.dispatch("user/login", this.loginForm);
    this.$message.success("success login in")
    this.$router.push("/")
  }
}
</script>

<style lang="scss">
.Login {
  padding: 40px 80px;
  .title {
    display: block;
    font-size: 20px;
    margin: 20px;
    text-align: center;
  }
}
</style>