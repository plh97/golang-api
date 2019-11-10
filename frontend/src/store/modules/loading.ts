import { ActionContext } from "vuex";
const state = {
  status: false
};

type StateType = {
  status: boolean;
};

const mutations = {
  SET_STATUS: (state: StateType, status: boolean) => {
    state.status = status;
  }
};

const actions = {
  // get user info
  start(context: ActionContext<any, any>) {
    return new Promise(() => {
      context.commit("SET_STATUS", true);
    });
  },
  end(context: ActionContext<any, any>) {
    return new Promise(() => {
      context.commit("SET_STATUS", false);
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
