import Vue from 'vue'
// 引入axios
import axios from 'axios'

let Url = 'http://127.0.0.1:3000/api/v1/'

axios.defaults.baseURL = Url
// axios拦截器 发送请求的时候，携带token
axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('ginblog_token')}`
  return config
})

Vue.prototype.$http = axios

export {Url}
