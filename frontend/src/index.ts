import Vue from 'vue';
import App from './App.vue';
import * as element from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import locale from 'element-ui/lib/locale/lang/en';
import Router from 'vue-router';
import axios from 'axios';

Vue.prototype.$http = axios;

Vue.use(element, { locale });
Vue.use(Router);

(() => {
  new Vue({
    el: '#root',
    render: h => h(App)
  });
})();

var vm = new Vue();
console.log(vm.$http);
