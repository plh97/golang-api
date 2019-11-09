const state = {
  status: false
};

const mutations = {
  SET_STATUS: (state, status) => {
    state.status = status;
  }
};

const actions = {
  // get user info
  start({ commit, state }) {
    return new Promise(() => {
      commit("SET_STATUS", true);
    });
  },
  end({ commit, state }) {
    return new Promise(() => {
      commit("SET_STATUS", false);
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
