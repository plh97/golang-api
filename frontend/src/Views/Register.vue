<template lang="pug">
  .Login {{name}}
    .title Register
    el-form(label-width="120px" :model="loginForm")
      el-form-item(label="Name")
        el-input(v-model="loginForm.name")
      el-form-item(label="Password")
        el-input(v-model="loginForm.password")
      el-form-item(label="Repeat Password")
        el-input(v-model="loginForm.repeatPassword")
      el-form-item
        router-link(to="/login")
          el-button Login
        span &nbsp;&nbsp;
          el-button(@click="handleRegister" type="primary") register

</template>

<script lang="ts">
import Vue from "vue";
import { mapGetters } from "vuex";
import Component from "vue-class-component";
import { Account } from "../dataStore";

type LoginType = {
  name: string;
  password: string;
};

@Component({
  computed: {
    ...mapGetters([
      "name"
    ])
  }
})
export default class Register extends Vue {
  public loginForm: LoginType = {
    name: "",
    password: ""
  };
  public async handleRegister(): Promise<void> {
    const name = await this.$store.dispatch('user/register', this.loginForm)
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