import axios from "axios";

const axiosApi = axios.create({
  baseURL: "http://35.247.143.160:8002"
});


export function getBooks() {
  return axiosApi({
    url: "/api/book",
    method: "GET"
  }).then(res => res.data);
}

export function updateBook(data: any) {
  return axiosApi({
    url: "/api/book/" + data.id,
    method: "PATCH",
    data
  }).then(res => res.data);
}

export function addBook(data: any) {
  return axiosApi({
    url: "/api/book",
    method: "POST",
    data
  }).then(res => res.data);
}

export function deleteBook(id: string) {
  return axiosApi({
    url: "/api/book/" + id,
    method: "DELETE",
  }).then(res => res.data);
}
