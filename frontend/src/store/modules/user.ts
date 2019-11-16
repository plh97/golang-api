import { ActionContext } from 'vuex';
import { Account } from '../../dataStore'
import { getToken, setToken, removeToken } from '../../utils/auth';

interface StateType {
  name: string;
  token: string;
}

const mutations = {
  SET_NAME: (state: StateType, name: string) => {
    state.name = name;
  },
  SET_TOKEN: (state: StateType, token: string) => {
    state.token = token;
  },
};

const actions = {
  // get user info
  login(context: ActionContext<any, any>) {
    return new Promise(async (resolve, reject) => {
      try {
        const { name } = await Account.login({
          name: '',
          password: '',
        })
        context.commit('SET_NAME', name);
        resolve(name);
      } catch (error) {
        reject(error)
      }
    })
  },
  logout(context: ActionContext<any, any>) {
    return new Promise(async (resolve, reject) => {
      try {
        await Account.logout({
          name: '',
          password: '',
        })
        context.commit('SET_NAME', '');
        context.commit('SET_TOKEN', '');
        removeToken();
        resolve();
      } catch (error) {
        reject(error);
      }
    });
  },
  register(context: ActionContext<any, any>) {
    return new Promise(async (resolve, reject) => {
      try {
        const { name, token } = await Account.register({
          name: '',
          password: '',
        })
        context.commit('SET_NAME', name);
        context.commit('SET_TOKEN', token);
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
  state: {
    name: ''
  },
  mutations,
  actions,
};
