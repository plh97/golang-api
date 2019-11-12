// profile/index.ts
import Vuex, { Store } from "vuex";
import Vue from "vue";
import getters from "./getters";
import loading from "./modules/loading";
import user from "./modules/user";
import createLogger from "vuex/dist/logger";

Vue.use(Vuex);

export default new Store({
  getters,
  modules: {
    loading,
    user,
  },
  plugins: [createLogger()]
});
