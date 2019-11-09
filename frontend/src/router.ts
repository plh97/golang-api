import Vue from "vue";
import VueRouter from "vue-router";
import Layout from "./Components/Layout/index.vue";
import Home from "./Views/Home.vue";

Vue.use(VueRouter);

export default new VueRouter({
  routes: [
    {
      path: "/",
      component: Layout,
      name: "home",
      children: [{ 
        path: "/", component: Home, name: "" }]
    }
  ]
});
