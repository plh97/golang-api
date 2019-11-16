<template lang="pug">
  .Login {{name}}
    .title Register
    el-form(label-width="120px" :model="form")
      el-form-item(label="Name")
        el-input(v-model="form.name")
      el-form-item(label="Password")
        el-input(v-model="form.password")
      el-form-item(label="Repeat Password")
        el-input(v-model="form.repeatPassword")
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

type FormType = {
  name: string;
  password: string;
  repeatPassword: string;
};

@Component({
  computed: {
    ...mapGetters([
      "name"
    ])
  }
})
export default class Register extends Vue {
  public form: FormType = {
    name: "admin",
    password: "12345678",
    repeatPassword: "12345678"
  };
  public async handleRegister(): Promise<void> {
    await this.$store.dispatch('user/register', {
      name: this.form.name,
      password: this.form.password,
    })
    this.$message.success("success register")
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