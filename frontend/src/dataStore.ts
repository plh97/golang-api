import axios, { AxiosError, AxiosResponse, AxiosRequestConfig } from 'axios';
import store from './store/index';
import { Message } from 'element-ui';
import { getToken } from './utils/auth';

interface LoginType {
  name: string;
  password: string;
}

interface UserInfoType {
  name: string;
  token: string;
}

interface ResponseType {
  errorCode: number;
  data: any;
  message: string;
}

interface Book {
  _id: string;
  name: string;
  author: string;
  createAt: Date;
  updateAt: Date;
}

const axiosApi = axios.create({
  baseURL: `//${document.domain}:8002`,
  // headers: {
  //   'Access-Control-Allow-Origin': '*',
  //   token: getToken()
  // },
});
axios.defaults.withCredentials = true
// request interceptor
axiosApi.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    store.dispatch('loading/start');
    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);
axiosApi.interceptors.response.use(
  (response: AxiosResponse) => {
    store.dispatch('loading/end');
    if (response.status >= 400 || response.status < 200) {
      if (response.status === 401) {
        Message({
          message: 'Unauthorized request',
          type: 'error',
          duration: 2000
        });
      } else {
        Message({
          message: 'Could not connect to server',
          type: 'error',
          duration: 2000
        });
      }
    } else {
      // 处理200的逻辑
      const res = response.data
      if (res.errorCode !== 0) {
        Message({
          message: res.message,
          type: 'error'
        });
      }
    }
    return response;
  },
  (error: AxiosError) => {
    if (error.response && error.response.status !== 401) {
      Message({
        message: error.response.data.message
          ? error.response.data.message
          : 'Could not connect to server',
        type: 'error',
        duration: 2000
      });
    }
    return Promise.reject(error);
  }
);

axiosApi.interceptors.response.use((r: AxiosResponse) => {
  return r.data
});

export const Book = {
  addBook(data: Book) {
    return axiosApi({
      url: '/api/book',
      method: 'POST',
      data
    });
  },
  deleteBook(id: string) {
    return axiosApi({
      url: '/api/book/' + id,
      method: 'DELETE'
    });
  },
  updateBook(data: Book) {
    return axiosApi({
      url: '/api/book/' + data._id,
      method: 'PATCH',
      data
    });
  },
  getBooks(): Promise<ResponseType> {
    return axiosApi.request<ResponseType, ResponseType>({
      url: '/api/book',
      method: 'GET'
    });
  }
};

export const Account = {
  userInfo(data: LoginType) {
    return axiosApi.request<ResponseType, ResponseType>({
      url: '/api/userInfo',
      method: 'GET',
      withCredentials: true,
      data
    });
  },
  login(data: LoginType) {
    return axiosApi.request<ResponseType, ResponseType>({
      url: '/api/login',
      method: 'POST',
      data
    });
  },
  logout(data: LoginType) {
    return axiosApi.request<UserInfoType, ResponseType>({
      url: '/api/logout',
      method: 'POST',
    });
  },
  register(data: LoginType) {
    return axiosApi.request<UserInfoType, ResponseType>({
      url: '/api/register',
      method: 'POST',
      data
    });
  },
};
