import Vue, { CreateElement } from "vue";
import * as element from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";

Vue.use(element);
Vue.config.devtools = true;

const _ = new Vue({
  el: "#root",
  router,
  store,
  render: (h: CreateElement) => h(App)
});
