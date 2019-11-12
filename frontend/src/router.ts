import Vue from "vue";
import VueRouter from "vue-router";
import Layout from "./Components/Layout/index.vue";
import Home from "./Views/Home.vue";
import Login from "./Views/Login.vue";
import store from './store'

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: "/",
      component: Layout,
      name: "home",
      redirect: "/home",
      children: [
        {
          name: "Home",
          path: "/home",
          component: Home
        },
        {
          path: "/login",
          name: "Login",
          component: Login
        }
      ]
    },
    {
      path:"*",
      redirect: '/login'
    }
  ]
});

router.beforeEach((to, from, next) => {
  next()
})

export default router;
