<template lang="pug">
  .Login {{name}}
    el-form(label-width="100px" :model="loginForm")
      el-form-item(label="Name")
        el-input(v-model="loginForm.name")
      el-form-item(label="Password")
        el-input(v-model="loginForm.password")
      el-form-item
        el-button(type="primary" @click="handleLogin") Login
        el-button(@click="handleRegister") register

</template>

<script lang="ts">
import Vue from "vue";
// import { mapGetters } from "vuex";
import Component from "vue-class-component";
import { Account } from "../dataStore";

type LoginType = {
  name: string;
  password: string;
};

@Component({
  // computed: {
  //   ...mapGetters([
  //     "name"
  //   ])
  // }
})
export default class Login extends Vue {
  public loginForm: LoginType = {
    name: "",
    password: ""
  };
  public name: string = "";
  public async handleLogin(): Promise<void> {
    const name = await this.$store.dispatch('user/login', this.loginForm)
    this.name = name;
  }
  public async handleRegister(): Promise<void> {
    const res = await Account.login(this.loginForm);
    console.log(res)
  }
}
</script>

<style lang="scss">
.App-Main {
  height: 100vh;
  overflow: scroll;
  padding: 10px;
  box-sizing: border-box;
  -webkit-overflow-scrolling: touch; /* Lets it scroll lazy */
}
</style>