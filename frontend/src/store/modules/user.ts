import { ActionContext } from "vuex";
import {Account} from '../../dataStore'
import { getToken, setToken, removeToken } from "../../utils/auth";

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
        const {name} = await Account.login({
          name: '',
          password: '',
        })
        context.commit("SET_NAME", name);
        resolve(name);
      } catch (error) {
        reject(error)
      }
    })
  },
  logout(context: ActionContext<any, any>) {
    return new Promise(async (resolve, reject) => {
      try {
        context.commit("SET_NAME", "");
        context.commit("SET_TOKEN", "");
        removeToken();
        resolve();
      } catch (error) {
        reject(error);
      }
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
};
