
import axiosOCR from '../utils/axios/axiosOCR';

import authHeader from './auth-header';

class OCRService {
  postImageData(data) {
    return axiosOCR.post('/', data, {headers: authHeader()});
  }
  postImageURL(data) {
    console.log(authHeader())
    return axiosOCR.post('/', data, {headers: authHeader()});
  }
}

export default new OCRService();