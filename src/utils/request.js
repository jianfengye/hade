import axios from "axios";
import store from "@/stores";

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  timeout: 5000, // request timeout
});

// request interceptor
service.interceptors.request.use(
  (config) => {
    // 如果全局有token，jwt请求增加token
    if (store.getters.token) {
      config.headers["X-Token"] = store.getters.token; 
    }
    return config;
  },
  (error) => {
    console.log(error); // for debug
    return Promise.reject(error);
  }
);

// response interceptor
service.interceptors.response.use(

  (response) => {
    const res = response.data;

    if (res.code !== 200) {
      alert("请求失败")
    } else {
      return res;
    }
  },
  (error) => {
    console.log("err" + error); // for debug
    return Promise.reject(error);
  }
);
export default service;