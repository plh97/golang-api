import { ActionContext } from 'vuex';

interface StateType {
  status: boolean;
}

const mutations = {
  SET_STATUS: (state: StateType, status: boolean) => {
    state.status = status;
  }
};

const actions = {
  // get user info
  start(context: ActionContext<any, any>) {
    return new Promise(() => {
      context.commit('SET_STATUS', true);
    });
  },
  end(context: ActionContext<any, any>) {
    return new Promise(() => {
      context.commit('SET_STATUS', false);
    });
  }
};

export default {
  name: 'loading',
  namespaced: true,
  state: {
    status: false
  },
  mutations,
  actions,
  // getter: {
  //   loading: (state: StateType) => state.status,
  // }
};
