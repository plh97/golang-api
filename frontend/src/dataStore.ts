import axios from 'axios';

export default {
  getBookList() {
    return axios('http://35.247.143.160:8002/api/book');
  },
  editBookList() {
    return axios('http://35.247.143.160:8002/api/book');
  },
};
