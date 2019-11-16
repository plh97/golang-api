import axios, { AxiosError, AxiosResponse, AxiosRequestConfig } from 'axios';
import store from './store/index';
import { Message } from 'element-ui';

interface LoginType {
  name: string;
  password: string;
}

interface UserInfoType {
  name: string;
  token: string;
}

interface Book {
  _id: string;
  name: string;
  author: string;
  createAt: Date;
  updateAt: Date;
}

const axiosApi = axios.create({
  baseURL: `//${document.domain}:8002`
});

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
        throw new Error(res)
      } else {
        return res
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

axiosApi.interceptors.response.use((r: AxiosResponse) => r.data);

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
  getBooks(): Promise<Book[]> {
    return axiosApi.request<Book[], Book[]>({
      url: '/api/book',
      method: 'GET'
    });
  }
};

export const Account = {
  login(data: LoginType) {
    return axiosApi.request<UserInfoType, UserInfoType>({
      url: '/api/login',
      method: 'POST',
      data
    });
  },
  logout(data: LoginType) {
    return axiosApi.request<UserInfoType, UserInfoType>({
      url: '/api/logout',
      method: 'POST',
    });
  },
  register(data: LoginType) {
    return axiosApi.request<UserInfoType, UserInfoType>({
      url: '/api/register',
      method: 'POST',
      data
    });
  },
};
