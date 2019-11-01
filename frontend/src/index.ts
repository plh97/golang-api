import Vue, { CreateElement } from 'vue';
import App from './App.vue';
import * as element from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Router from 'vue-router';
import axios from 'axios';

Vue.prototype.$http = axios;

Vue.use(element);
Vue.use(Router);
Vue.config.devtools = true;

const _ = new Vue({
  el: '#root',
  render: (h: CreateElement) => h(App),
});
