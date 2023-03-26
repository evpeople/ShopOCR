
import axiosUpload from '../utils/axios/axiosUpload';

import authHeader from './auth-header';

class UploadService {
  uploadImg(data) {
    return axiosUpload.post('/', data, {
      headers: {
        'Content-Type': 'multipart/form-data',
        // 设置Content-Type头部
      }
    })
  }
}

export default new UploadService();