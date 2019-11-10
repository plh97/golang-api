import Vue from "vue";
import VueRouter from "vue-router";
import Layout from "./Components/Layout/index.vue";
import Home from "./Views/Home.vue";

Vue.use(VueRouter);

export default new VueRouter({
  routes: [
    {
      path: "/",
      redirect: "/home"
    },
    {
      path: "/home",
      component: Layout,
      name: "home",
      redirect: "/home/index",
      children: [
        {
          path: "/home/index",
          component: Home,
          name: "Index"
        }
      ]
    }
  ]
});
