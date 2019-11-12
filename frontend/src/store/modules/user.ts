import { ActionContext } from "vuex";
import dataStore from '../../dataStore'

type StateType = {
  name: string;
};

const state = {
  name: ''
};

const mutations = {
  SET_NAME: (state: StateType, name: string) => {
    state.name = name;
  }
};

const actions = {
  // get user info
  login(context: ActionContext<any, any>) {
    return new Promise(async(resolve,reject) => {
      try {
        const {name} = await dataStore.login({
          name: '',
          password: '',
        })
        context.commit("SET_NAME", name);
        resolve(name);
      } catch (error) {
        reject(error)
      }
    })
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  // getter: {
  //   name: (state: StateType) => state.name
  // }
};
