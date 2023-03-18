
import axios from 'axios'

var axios2 = axios.create({
  baseURL: 'http://localhost:8888/user/',  // 设置服务器地址
  timeout: 60 * 1000,                      // 超时设置
  changeOrigin: true                       // 允许跨域
})

export default axios2