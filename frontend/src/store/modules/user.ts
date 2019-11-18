import { ActionContext } from 'vuex';
import { Account } from '../../dataStore'
import { getToken, setToken, removeToken } from '../../utils/auth';

interface StateType {
  name: string;
  token: string;
}

interface LoginType {
  name: string;
  password: string;
}


const mutations = {
  SET_NAME: (state: StateType, name: string) => {
    state.name = name;
  },
  SET_TOKEN: (state: StateType, token: string) => {
    setToken(token)
    state.token = token;
  },
  REMOVE_TOKEN: (state: StateType, token: string) => {
    removeToken()
    state.token = '';
  },
};

const actions = {
  getInfo(context: ActionContext<any, any>, form: LoginType) {
    return new Promise(async (resolve, reject) => {
        const res = await Account.userInfo(form)
    })
  },
  // get user info
  login(context: ActionContext<any, any>, form: LoginType) {
    return new Promise(async (resolve, reject) => {
      try {
        const res = await Account.login(form)
        const { errorCode, data } = res
        if (errorCode === 0) {
          const { name, token } = data
          context.commit('SET_NAME', name);
          context.commit('SET_TOKEN', token);
          resolve(res);
        } else {
          reject(res)
        }
      } catch (error) {
        reject(error)
      }
    })
  },
  logout(context: ActionContext<any, any>) {
    return new Promise(async (resolve, reject) => {
      try {
        context.commit('SET_NAME', '');
        context.commit('REMOVE_TOKEN');
        resolve();
      } catch (error) {
        reject(error);
      }
    });
  },
  register(context: ActionContext<any, any>, form: LoginType) {
    return new Promise(async (resolve, reject) => {
      try {
        const { errorCode, data, message } = await Account.register(form)
        if (errorCode === 0) {
          const { name, token } = data
          context.commit('SET_NAME', name);
          context.commit('SET_TOKEN', token);
          resolve(data);
        } else {
          reject(message);
        }
      } catch (error) {
        reject(error);
      }
    });
  }
};

export default {
  name: 'user',
  namespaced: true,
  state: {
    name: '',
    token: ''
  },
  mutations,
  actions,
};
