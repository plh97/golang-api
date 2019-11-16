import Vue from 'vue';
import VueRouter from 'vue-router';
import Layout from './Components/Layout/index';
import Home from './Views/Home';
import Login from './Views/Login';
import Register from './Views/Register';
import { getToken } from './utils/auth'; // get token from cookie
import store from './store'; // get token from cookie

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Layout,
      name: 'home',
      redirect: '/home',
      children: [
        {
          name: 'Home',
          path: '/home',
          component: Home
        }
      ]
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '*',
      redirect: '/login'
    }
  ]
});

router.beforeEach((to, from, next) => {
  console.log('hasToken', getToken(), store.getters.token);
  if (to.path === '/login' || to.path === '/register') {
    return next();
  } else if (store.getters.token) {
    return next();
  } else {
    // 如果没有 token
    if (getToken()) {
      store.commit('SET_TOKEN', getToken())
      next()
    } else {
      next({ path: '/login' });
    }
  }
});

export default router;
