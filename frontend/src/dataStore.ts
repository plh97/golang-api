import axios from "axios";

export function getBooks() {
  return axios({
    url: "http://35.247.143.160:8002/api/book",
    method: "GET"
  }).then(res => res.data);
}

export function updateBook(data: any) {
  return axios({
    url: "http://35.247.143.160:8002/api/book/" + data.id,
    method: "PATCH",
    withCredentials: false,
    data
  }).then(res => res.data);
}

export function addBook(data: any) {
  return axios({
    url: "http://35.247.143.160:8002/api/book",
    method: "POST",
    data
  }).then(res => res.data);
}

export function deleteBook(id: string) {
  return axios({
    withCredentials: true,
    url: "http://35.247.143.160:8002/api/book/" + id,
    method: "DELETE",
    data: {
      id: id
    }
  }).then(res => res.data);
}
