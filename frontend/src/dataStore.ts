import axios, { AxiosError, AxiosResponse, AxiosRequestConfig } from "axios";
import store from "./store/index";
import { Message } from "element-ui";

type LoginType = {
  name: string;
  password: string;
};

type UserInfoType = {
  name: string;
}

type Book = {
  _id: string;
  name: string;
  author: string;
  createAt: Date;
  updateAt: Date;
};

const axiosApi = axios.create({
  baseURL: `//${document.domain}:8002`
});

// request interceptor
axiosApi.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    store.dispatch("loading/start");
    return config;
  },
  (error: AxiosError) => {
    // do something with request error
    return Promise.reject(error);
  }
);
// response interceptor
axiosApi.interceptors.response.use(
  (response: AxiosResponse) => {
    // const res = response.data;
    // 先根据statusCode判断
    store.dispatch("loading/end");
    if (response.status >= 400 || response.status < 200) {
      // 如果是非200的,则判断是哪种类型的错误,并handle住
      if (response.status === 401) {
        Message({
          message: "Unauthorized request",
          type: "error",
          duration: 2000
        });
      } else {
        Message({
          message: "Could not connect to server",
          type: "error",
          duration: 2000
        });
      }
    } else {
      // 处理200的逻辑
    }
    return response;
  },
  (error: AxiosError) => {
    if (error.response && error.response.status !== 401) {
      Message({
        message: error.response.data.message
          ? error.response.data.message
          : "Could not connect to server",
        type: "error",
        duration: 2000
      });
    }
    return Promise.reject(error);
  }
);

axiosApi.interceptors.response.use((r: AxiosResponse) => r.data);


export default {
  addBook(data: Book) {
    return axiosApi({
      url: "/api/book",
      method: "POST",
      data
    });
  },
  deleteBook(_id: string) {
    return axiosApi({
      url: "/api/book/" + _id,
      method: "DELETE"
    });
  },
  updateBook(data: Book) {
    return axiosApi({
      url: "/api/book/" + data._id,
      method: "PATCH",
      data: data
    });
  },
  getBooks(): Promise<Book[]> {
    return axiosApi.request<Book[], Book[]>({
      url: "/api/book",
      method: "GET"
    });
  },
  login(data: LoginType) {
    return axiosApi.request<UserInfoType, UserInfoType>({
      url: "/api/login",
      method: "POST",
      data
    });
  }
};
