// profile/index.ts
import Vuex, { Module, Store } from "vuex";
import Vue from "vue";
import getters from "./getters";
import { RootState, ProfileState } from "./types";
import loading from "./modules/loading";
const namespaced: boolean = true;
import createLogger from "vuex/dist/logger";

Vue.use(Vuex);

export default new Store({
  // namespaced,
  getters,
  modules: {
    loading
  },
  plugins: [createLogger()]
});
